package gloo_test

import (
	"context"
	"net"
	"os"
	"sync"
	"time"

	"github.com/solo-io/gloo/pkg/utils/setuputils"

	"github.com/solo-io/gloo/projects/gloo/pkg/plugins/registry"

	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/graphql/v1beta1"

	"github.com/solo-io/gloo/projects/gloo/pkg/defaults"
	"github.com/solo-io/solo-kit/pkg/utils/statusutils"

	gatewayv1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1"
	v1alpha1 "github.com/solo-io/gloo/projects/gloo/pkg/api/external/solo/ratelimit"
	extauthv1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/extauth/v1"
	"github.com/solo-io/k8s-utils/kubeutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/crd"
	"github.com/solo-io/solo-kit/test/helpers"
	apiext "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"

	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube"
	"github.com/solo-io/solo-kit/pkg/utils/prototime"
	"google.golang.org/grpc"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/solo-io/gloo/pkg/utils/settingsutil"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/memory"
	reflectpb "google.golang.org/grpc/reflection/grpc_reflection_v1alpha"

	. "github.com/solo-io/gloo/projects/gloo/pkg/syncer/setup"
)

var _ = Describe("SetupSyncer", func() {

	var (
		settings  *v1.Settings
		ctx       context.Context
		cancel    context.CancelFunc
		memcache  memory.InMemoryResourceCache
		setupLock sync.RWMutex
	)

	newContext := func() {
		if cancel != nil {
			cancel()
		}
		ctx, cancel = context.WithCancel(context.Background())
		ctx = settingsutil.WithSettings(ctx, settings)
	}

	// SetupFunc is used to configure Gloo with appropriate configuration
	// It is assumed to run once at construction time, and therefore it executes directives that
	// are also assumed to only run at construction time.
	// One of those, is the construction of schemes: https://github.com/kubernetes/kubernetes/pull/89019#issuecomment-600278461
	// In our tests we do not follow this pattern, and to avoid data races (that cause test failures)
	// we ensure that only 1 SetupFunc is ever called at a time
	newSynchronizedSetupFunc := func() setuputils.SetupFunc {
		setupFunc := NewSetupFunc()

		var synchronizedSetupFunc setuputils.SetupFunc
		synchronizedSetupFunc = func(ctx context.Context, kubeCache kube.SharedCache, inMemoryCache memory.InMemoryResourceCache, settings *v1.Settings) error {
			setupLock.Lock()
			defer setupLock.Unlock()
			return setupFunc(ctx, kubeCache, inMemoryCache, settings)
		}

		return synchronizedSetupFunc
	}

	BeforeEach(func() {
		settings = &v1.Settings{
			RefreshRate: prototime.DurationToProto(time.Hour),
			Gloo: &v1.GlooOptions{
				XdsBindAddr:        getRandomAddr(),
				ValidationBindAddr: getRandomAddr(),
			},
			DiscoveryNamespace: "non-existent-namespace",
			WatchNamespaces:    []string{"non-existent-namespace"},
		}
		memcache = memory.NewInMemoryResourceCache()
		newContext()
	})

	AfterEach(func() {
		cancel()
	})

	Context("Setup", func() {
		setupTestGrpcClient := func() func() error {
			cc, err := grpc.DialContext(ctx, settings.Gloo.XdsBindAddr, grpc.WithInsecure(), grpc.FailOnNonTempDialError(true))
			Expect(err).NotTo(HaveOccurred())
			// setup a gRPC client to make sure connection is persistent across invocations
			client := reflectpb.NewServerReflectionClient(cc)
			req := &reflectpb.ServerReflectionRequest{
				MessageRequest: &reflectpb.ServerReflectionRequest_ListServices{
					ListServices: "*",
				},
			}
			clientstream, err := client.ServerReflectionInfo(context.Background())
			Expect(err).NotTo(HaveOccurred())
			err = clientstream.Send(req)
			go func() {
				for {
					_, err := clientstream.Recv()
					if err != nil {
						return
					}
				}
			}()
			Expect(err).NotTo(HaveOccurred())
			return func() error { return clientstream.Send(req) }
		}

		Context("XDS tests", func() {

			It("setup can be called twice", func() {
				setup := newSynchronizedSetupFunc()

				err := setup(ctx, nil, memcache, settings)
				Expect(err).NotTo(HaveOccurred())

				testFunc := setupTestGrpcClient()

				newContext()
				err = setup(ctx, nil, memcache, settings)
				Expect(err).NotTo(HaveOccurred())

				// give things a chance to react
				time.Sleep(time.Second)

				// make sure that xds snapshot was not restarted
				err = testFunc()
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("Extensions tests", func() {

			var (
				plugin1 = &dummyPlugin{}
				plugin2 = &dummyPlugin{}
			)

			It("should return plugins", func() {
				extensions := Extensions{
					PluginRegistryFactory: func(ctx context.Context) plugins.PluginRegistry {
						return registry.NewPluginRegistry([]plugins.Plugin{
							plugin1,
							plugin2,
						})
					},
				}

				pluginRegistry := extensions.PluginRegistryFactory(context.TODO())
				plugins := pluginRegistry.GetPlugins()
				Expect(plugins).To(ContainElement(plugin1))
				Expect(plugins).To(ContainElement(plugin2))
			})

		})

		Context("Kube tests", func() {

			var (
				kubeCoreCache    kube.SharedCache
				registerCrdsOnce sync.Once
			)

			registerCRDs := func() {
				cfg, err := kubeutils.GetConfig("", "")
				ExpectWithOffset(1, err).NotTo(HaveOccurred())

				apiExts, err := apiext.NewForConfig(cfg)
				ExpectWithOffset(1, err).NotTo(HaveOccurred())

				crdsToRegister := []crd.Crd{
					v1.UpstreamCrd,
					v1.UpstreamGroupCrd,
					v1.ProxyCrd,
					gatewayv1.GatewayCrd,
					extauthv1.AuthConfigCrd,
					v1alpha1.RateLimitConfigCrd,
					v1beta1.GraphQLApiCrd,
					gatewayv1.VirtualServiceCrd,
					gatewayv1.RouteOptionCrd,
					gatewayv1.VirtualHostOptionCrd,
					gatewayv1.RouteTableCrd,
				}

				for _, crdToRegister := range crdsToRegister {
					err = helpers.AddAndRegisterCrd(ctx, crdToRegister, apiExts)
					ExpectWithOffset(1, err).NotTo(HaveOccurred())
				}
			}

			BeforeEach(func() {
				if os.Getenv("RUN_KUBE_TESTS") != "1" {
					Skip("This test creates kubernetes resources and is disabled by default. To enable, set RUN_KUBE_TESTS=1 in your env.")
				}
				settings.ConfigSource = &v1.Settings_KubernetesConfigSource{KubernetesConfigSource: &v1.Settings_KubernetesCrds{}}
				settings.SecretSource = &v1.Settings_KubernetesSecretSource{KubernetesSecretSource: &v1.Settings_KubernetesSecrets{}}
				settings.ArtifactSource = &v1.Settings_KubernetesArtifactSource{KubernetesArtifactSource: &v1.Settings_KubernetesConfigmaps{}}
				kubeCoreCache = kube.NewKubeCache(ctx)

				// Gloo SetupFunc is no longer responsible for registering CRDs. This was not used in production, and
				// required Gloo having RBAC permissions that it should not have. CRD registration is now only supported
				// by Helm. Therefore, this test needs to manually register CRDs to test setup.
				registerCrdsOnce.Do(registerCRDs)

				err := os.Setenv(statusutils.PodNamespaceEnvName, defaults.GlooSystem)
				Expect(err).NotTo(HaveOccurred())
			})

			AfterEach(func() {
				err := os.Unsetenv(statusutils.PodNamespaceEnvName)
				Expect(err).NotTo(HaveOccurred())
			})

			It("can be called with core cache", func() {
				setup := newSynchronizedSetupFunc()
				err := setup(ctx, kubeCoreCache, memcache, settings)
				Expect(err).NotTo(HaveOccurred())
			})

			It("can be called with core cache warming endpoints", func() {
				settings.Gloo.EndpointsWarmingTimeout = prototime.DurationToProto(time.Minute)
				setup := newSynchronizedSetupFunc()
				err := setup(ctx, kubeCoreCache, memcache, settings)
				Expect(err).NotTo(HaveOccurred())
			})

			It("panics when endpoints don't arrive in a timely manner", func() {
				settings.Gloo.EndpointsWarmingTimeout = prototime.DurationToProto(1 * time.Nanosecond)
				setup := newSynchronizedSetupFunc()
				Expect(func() { setup(ctx, kubeCoreCache, memcache, settings) }).To(Panic())
			})

			It("doesn't panic when endpoints don't arrive in a timely manner if set to zero", func() {
				settings.Gloo.EndpointsWarmingTimeout = prototime.DurationToProto(0)
				setup := newSynchronizedSetupFunc()
				Expect(func() { setup(ctx, kubeCoreCache, memcache, settings) }).NotTo(Panic())
			})

		})
	})
})

func getRandomAddr() string {
	listener, err := net.Listen("tcp", "localhost:0")
	Expect(err).NotTo(HaveOccurred())
	addr := listener.Addr().String()
	listener.Close()
	return addr
}

type dummyPlugin struct{}

func (*dummyPlugin) Name() string { return "dummy_plugin" }

func (*dummyPlugin) Init(params plugins.InitParams) error { return nil }
