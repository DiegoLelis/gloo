// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo/projects/gloo/api/v1/load_balancer.proto

package v1

import (
	reflect "reflect"
	sync "sync"

	duration "github.com/golang/protobuf/ptypes/duration"
	empty "github.com/golang/protobuf/ptypes/empty"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// LoadBalancerConfig is the settings for the load balancer used to send request to the Upstream
// endpoints.
type LoadBalancerConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Configures envoy's panic threshold Percent between 0-100. Once the number of non health hosts
	// reaches this percentage, envoy disregards health information.
	// see more info [here](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/upstream/load_balancing/panic_threshold.html).
	HealthyPanicThreshold *wrappers.DoubleValue `protobuf:"bytes,1,opt,name=healthy_panic_threshold,json=healthyPanicThreshold,proto3" json:"healthy_panic_threshold,omitempty"`
	// This allows batch updates of endpoints health/weight/metadata that happen during a time window.
	// this help lower cpu usage when endpoint change rate is high. defaults to 1 second.
	// Set to 0 to disable and have changes applied immediately.
	UpdateMergeWindow *duration.Duration `protobuf:"bytes,2,opt,name=update_merge_window,json=updateMergeWindow,proto3" json:"update_merge_window,omitempty"`
	// Types that are assignable to Type:
	//	*LoadBalancerConfig_RoundRobin_
	//	*LoadBalancerConfig_LeastRequest_
	//	*LoadBalancerConfig_Random_
	//	*LoadBalancerConfig_RingHash_
	//	*LoadBalancerConfig_Maglev_
	Type isLoadBalancerConfig_Type `protobuf_oneof:"type"`
	// Types that are assignable to LocalityConfig:
	//	*LoadBalancerConfig_LocalityWeightedLbConfig
	LocalityConfig isLoadBalancerConfig_LocalityConfig `protobuf_oneof:"locality_config"`
}

func (x *LoadBalancerConfig) Reset() {
	*x = LoadBalancerConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoadBalancerConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadBalancerConfig) ProtoMessage() {}

func (x *LoadBalancerConfig) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadBalancerConfig.ProtoReflect.Descriptor instead.
func (*LoadBalancerConfig) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_rawDescGZIP(), []int{0}
}

func (x *LoadBalancerConfig) GetHealthyPanicThreshold() *wrappers.DoubleValue {
	if x != nil {
		return x.HealthyPanicThreshold
	}
	return nil
}

func (x *LoadBalancerConfig) GetUpdateMergeWindow() *duration.Duration {
	if x != nil {
		return x.UpdateMergeWindow
	}
	return nil
}

func (m *LoadBalancerConfig) GetType() isLoadBalancerConfig_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *LoadBalancerConfig) GetRoundRobin() *LoadBalancerConfig_RoundRobin {
	if x, ok := x.GetType().(*LoadBalancerConfig_RoundRobin_); ok {
		return x.RoundRobin
	}
	return nil
}

func (x *LoadBalancerConfig) GetLeastRequest() *LoadBalancerConfig_LeastRequest {
	if x, ok := x.GetType().(*LoadBalancerConfig_LeastRequest_); ok {
		return x.LeastRequest
	}
	return nil
}

func (x *LoadBalancerConfig) GetRandom() *LoadBalancerConfig_Random {
	if x, ok := x.GetType().(*LoadBalancerConfig_Random_); ok {
		return x.Random
	}
	return nil
}

func (x *LoadBalancerConfig) GetRingHash() *LoadBalancerConfig_RingHash {
	if x, ok := x.GetType().(*LoadBalancerConfig_RingHash_); ok {
		return x.RingHash
	}
	return nil
}

func (x *LoadBalancerConfig) GetMaglev() *LoadBalancerConfig_Maglev {
	if x, ok := x.GetType().(*LoadBalancerConfig_Maglev_); ok {
		return x.Maglev
	}
	return nil
}

func (m *LoadBalancerConfig) GetLocalityConfig() isLoadBalancerConfig_LocalityConfig {
	if m != nil {
		return m.LocalityConfig
	}
	return nil
}

func (x *LoadBalancerConfig) GetLocalityWeightedLbConfig() *empty.Empty {
	if x, ok := x.GetLocalityConfig().(*LoadBalancerConfig_LocalityWeightedLbConfig); ok {
		return x.LocalityWeightedLbConfig
	}
	return nil
}

type isLoadBalancerConfig_Type interface {
	isLoadBalancerConfig_Type()
}

type LoadBalancerConfig_RoundRobin_ struct {
	// Use round robin for load balancing.
	RoundRobin *LoadBalancerConfig_RoundRobin `protobuf:"bytes,3,opt,name=round_robin,json=roundRobin,proto3,oneof"`
}

type LoadBalancerConfig_LeastRequest_ struct {
	// Use least request for load balancing.
	LeastRequest *LoadBalancerConfig_LeastRequest `protobuf:"bytes,4,opt,name=least_request,json=leastRequest,proto3,oneof"`
}

type LoadBalancerConfig_Random_ struct {
	// Use random for load balancing.
	Random *LoadBalancerConfig_Random `protobuf:"bytes,5,opt,name=random,proto3,oneof"`
}

type LoadBalancerConfig_RingHash_ struct {
	// Use ring hash for load balancing.
	RingHash *LoadBalancerConfig_RingHash `protobuf:"bytes,6,opt,name=ring_hash,json=ringHash,proto3,oneof"`
}

type LoadBalancerConfig_Maglev_ struct {
	// Use maglev for load balancing.
	Maglev *LoadBalancerConfig_Maglev `protobuf:"bytes,7,opt,name=maglev,proto3,oneof"`
}

func (*LoadBalancerConfig_RoundRobin_) isLoadBalancerConfig_Type() {}

func (*LoadBalancerConfig_LeastRequest_) isLoadBalancerConfig_Type() {}

func (*LoadBalancerConfig_Random_) isLoadBalancerConfig_Type() {}

func (*LoadBalancerConfig_RingHash_) isLoadBalancerConfig_Type() {}

func (*LoadBalancerConfig_Maglev_) isLoadBalancerConfig_Type() {}

type isLoadBalancerConfig_LocalityConfig interface {
	isLoadBalancerConfig_LocalityConfig()
}

type LoadBalancerConfig_LocalityWeightedLbConfig struct {
	// (Enterprise Only)
	// https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/upstream/load_balancing/locality_weight#locality-weighted-load-balancing
	// Locality weighted load balancing enables weighting assignments across different zones and geographical locations by using explicit weights.
	// This field is required to enable locality weighted load balancing
	LocalityWeightedLbConfig *empty.Empty `protobuf:"bytes,8,opt,name=locality_weighted_lb_config,json=localityWeightedLbConfig,proto3,oneof"`
}

func (*LoadBalancerConfig_LocalityWeightedLbConfig) isLoadBalancerConfig_LocalityConfig() {}

type LoadBalancerConfig_RoundRobin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LoadBalancerConfig_RoundRobin) Reset() {
	*x = LoadBalancerConfig_RoundRobin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoadBalancerConfig_RoundRobin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadBalancerConfig_RoundRobin) ProtoMessage() {}

func (x *LoadBalancerConfig_RoundRobin) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadBalancerConfig_RoundRobin.ProtoReflect.Descriptor instead.
func (*LoadBalancerConfig_RoundRobin) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_rawDescGZIP(), []int{0, 0}
}

type LoadBalancerConfig_LeastRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// How many choices to take into account. defaults to 2.
	ChoiceCount uint32 `protobuf:"varint,1,opt,name=choice_count,json=choiceCount,proto3" json:"choice_count,omitempty"`
}

func (x *LoadBalancerConfig_LeastRequest) Reset() {
	*x = LoadBalancerConfig_LeastRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoadBalancerConfig_LeastRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadBalancerConfig_LeastRequest) ProtoMessage() {}

func (x *LoadBalancerConfig_LeastRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadBalancerConfig_LeastRequest.ProtoReflect.Descriptor instead.
func (*LoadBalancerConfig_LeastRequest) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_rawDescGZIP(), []int{0, 1}
}

func (x *LoadBalancerConfig_LeastRequest) GetChoiceCount() uint32 {
	if x != nil {
		return x.ChoiceCount
	}
	return 0
}

type LoadBalancerConfig_Random struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LoadBalancerConfig_Random) Reset() {
	*x = LoadBalancerConfig_Random{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoadBalancerConfig_Random) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadBalancerConfig_Random) ProtoMessage() {}

func (x *LoadBalancerConfig_Random) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadBalancerConfig_Random.ProtoReflect.Descriptor instead.
func (*LoadBalancerConfig_Random) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_rawDescGZIP(), []int{0, 2}
}

// Customizes the parameters used in the hashing algorithm to refine performance or resource usage.
type LoadBalancerConfig_RingHashConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Minimum hash ring size. The larger the ring is (that is, the more hashes there are for each provided host)
	// the better the request distribution will reflect the desired weights. Defaults to 1024 entries, and limited
	// to 8M entries.
	MinimumRingSize uint64 `protobuf:"varint,1,opt,name=minimum_ring_size,json=minimumRingSize,proto3" json:"minimum_ring_size,omitempty"`
	// Maximum hash ring size. Defaults to 8M entries, and limited to 8M entries, but can be lowered to further
	// constrain resource use.
	MaximumRingSize uint64 `protobuf:"varint,2,opt,name=maximum_ring_size,json=maximumRingSize,proto3" json:"maximum_ring_size,omitempty"`
}

func (x *LoadBalancerConfig_RingHashConfig) Reset() {
	*x = LoadBalancerConfig_RingHashConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoadBalancerConfig_RingHashConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadBalancerConfig_RingHashConfig) ProtoMessage() {}

func (x *LoadBalancerConfig_RingHashConfig) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadBalancerConfig_RingHashConfig.ProtoReflect.Descriptor instead.
func (*LoadBalancerConfig_RingHashConfig) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_rawDescGZIP(), []int{0, 3}
}

func (x *LoadBalancerConfig_RingHashConfig) GetMinimumRingSize() uint64 {
	if x != nil {
		return x.MinimumRingSize
	}
	return 0
}

func (x *LoadBalancerConfig_RingHashConfig) GetMaximumRingSize() uint64 {
	if x != nil {
		return x.MaximumRingSize
	}
	return 0
}

type LoadBalancerConfig_RingHash struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional, customizes the parameters used in the hashing algorithm
	RingHashConfig *LoadBalancerConfig_RingHashConfig `protobuf:"bytes,1,opt,name=ring_hash_config,json=ringHashConfig,proto3" json:"ring_hash_config,omitempty"`
}

func (x *LoadBalancerConfig_RingHash) Reset() {
	*x = LoadBalancerConfig_RingHash{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoadBalancerConfig_RingHash) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadBalancerConfig_RingHash) ProtoMessage() {}

func (x *LoadBalancerConfig_RingHash) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadBalancerConfig_RingHash.ProtoReflect.Descriptor instead.
func (*LoadBalancerConfig_RingHash) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_rawDescGZIP(), []int{0, 4}
}

func (x *LoadBalancerConfig_RingHash) GetRingHashConfig() *LoadBalancerConfig_RingHashConfig {
	if x != nil {
		return x.RingHashConfig
	}
	return nil
}

type LoadBalancerConfig_Maglev struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LoadBalancerConfig_Maglev) Reset() {
	*x = LoadBalancerConfig_Maglev{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoadBalancerConfig_Maglev) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadBalancerConfig_Maglev) ProtoMessage() {}

func (x *LoadBalancerConfig_Maglev) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadBalancerConfig_Maglev.ProtoReflect.Descriptor instead.
func (*LoadBalancerConfig_Maglev) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_rawDescGZIP(), []int{0, 5}
}

var File_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_rawDesc = []byte{
	0x0a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6c,
	0x6f, 0x61, 0x64, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0c, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x65,
	0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xc5, 0x07, 0x0a, 0x12, 0x4c, 0x6f, 0x61, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x54, 0x0a, 0x17, 0x68, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x79, 0x5f, 0x70, 0x61, 0x6e, 0x69, 0x63, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68,
	0x6f, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x6f, 0x75, 0x62,
	0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x15, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x79,
	0x50, 0x61, 0x6e, 0x69, 0x63, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x49,
	0x0a, 0x13, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x5f, 0x77,
	0x69, 0x6e, 0x64, 0x6f, 0x77, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x11, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65,
	0x72, 0x67, 0x65, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x12, 0x4e, 0x0a, 0x0b, 0x72, 0x6f, 0x75,
	0x6e, 0x64, 0x5f, 0x72, 0x6f, 0x62, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b,
	0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4c, 0x6f,
	0x61, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x52, 0x6f, 0x62, 0x69, 0x6e, 0x48, 0x00, 0x52, 0x0a, 0x72,
	0x6f, 0x75, 0x6e, 0x64, 0x52, 0x6f, 0x62, 0x69, 0x6e, 0x12, 0x54, 0x0a, 0x0d, 0x6c, 0x65, 0x61,
	0x73, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2d, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e,
	0x4c, 0x6f, 0x61, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x4c, 0x65, 0x61, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48,
	0x00, 0x52, 0x0c, 0x6c, 0x65, 0x61, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x41, 0x0a, 0x06, 0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x27, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4c,
	0x6f, 0x61, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x48, 0x00, 0x52, 0x06, 0x72, 0x61, 0x6e, 0x64,
	0x6f, 0x6d, 0x12, 0x48, 0x0a, 0x09, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c,
	0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4c, 0x6f, 0x61, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x52, 0x69, 0x6e, 0x67, 0x48, 0x61, 0x73, 0x68,
	0x48, 0x00, 0x52, 0x08, 0x72, 0x69, 0x6e, 0x67, 0x48, 0x61, 0x73, 0x68, 0x12, 0x41, 0x0a, 0x06,
	0x6d, 0x61, 0x67, 0x6c, 0x65, 0x76, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x67,
	0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4c, 0x6f, 0x61, 0x64,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4d,
	0x61, 0x67, 0x6c, 0x65, 0x76, 0x48, 0x00, 0x52, 0x06, 0x6d, 0x61, 0x67, 0x6c, 0x65, 0x76, 0x12,
	0x57, 0x0a, 0x1b, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x77, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x65, 0x64, 0x5f, 0x6c, 0x62, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x48, 0x01, 0x52, 0x18,
	0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x65, 0x64,
	0x4c, 0x62, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x0c, 0x0a, 0x0a, 0x52, 0x6f, 0x75, 0x6e,
	0x64, 0x52, 0x6f, 0x62, 0x69, 0x6e, 0x1a, 0x31, 0x0a, 0x0c, 0x4c, 0x65, 0x61, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x63, 0x68,
	0x6f, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x08, 0x0a, 0x06, 0x52, 0x61, 0x6e,
	0x64, 0x6f, 0x6d, 0x1a, 0x68, 0x0a, 0x0e, 0x52, 0x69, 0x6e, 0x67, 0x48, 0x61, 0x73, 0x68, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x2a, 0x0a, 0x11, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d,
	0x5f, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0f, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x52, 0x69, 0x6e, 0x67, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x2a, 0x0a, 0x11, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x5f, 0x72, 0x69, 0x6e,
	0x67, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x6d, 0x61,
	0x78, 0x69, 0x6d, 0x75, 0x6d, 0x52, 0x69, 0x6e, 0x67, 0x53, 0x69, 0x7a, 0x65, 0x1a, 0x65, 0x0a,
	0x08, 0x52, 0x69, 0x6e, 0x67, 0x48, 0x61, 0x73, 0x68, 0x12, 0x59, 0x0a, 0x10, 0x72, 0x69, 0x6e,
	0x67, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e,
	0x69, 0x6f, 0x2e, 0x4c, 0x6f, 0x61, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x52, 0x69, 0x6e, 0x67, 0x48, 0x61, 0x73, 0x68, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x0e, 0x72, 0x69, 0x6e, 0x67, 0x48, 0x61, 0x73, 0x68, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x1a, 0x08, 0x0a, 0x06, 0x4d, 0x61, 0x67, 0x6c, 0x65, 0x76, 0x42, 0x06,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x42, 0x11, 0x0a, 0x0f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69,
	0x74, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x3e, 0x5a, 0x30, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f,
	0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c,
	0x6f, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0xb8, 0xf5, 0x04,
	0x01, 0xc0, 0xf5, 0x04, 0x01, 0xd0, 0xf5, 0x04, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_rawDescData = file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_rawDesc
)

func file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_rawDescData
}

var file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_goTypes = []interface{}{
	(*LoadBalancerConfig)(nil),                // 0: gloo.solo.io.LoadBalancerConfig
	(*LoadBalancerConfig_RoundRobin)(nil),     // 1: gloo.solo.io.LoadBalancerConfig.RoundRobin
	(*LoadBalancerConfig_LeastRequest)(nil),   // 2: gloo.solo.io.LoadBalancerConfig.LeastRequest
	(*LoadBalancerConfig_Random)(nil),         // 3: gloo.solo.io.LoadBalancerConfig.Random
	(*LoadBalancerConfig_RingHashConfig)(nil), // 4: gloo.solo.io.LoadBalancerConfig.RingHashConfig
	(*LoadBalancerConfig_RingHash)(nil),       // 5: gloo.solo.io.LoadBalancerConfig.RingHash
	(*LoadBalancerConfig_Maglev)(nil),         // 6: gloo.solo.io.LoadBalancerConfig.Maglev
	(*wrappers.DoubleValue)(nil),              // 7: google.protobuf.DoubleValue
	(*duration.Duration)(nil),                 // 8: google.protobuf.Duration
	(*empty.Empty)(nil),                       // 9: google.protobuf.Empty
}
var file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_depIdxs = []int32{
	7, // 0: gloo.solo.io.LoadBalancerConfig.healthy_panic_threshold:type_name -> google.protobuf.DoubleValue
	8, // 1: gloo.solo.io.LoadBalancerConfig.update_merge_window:type_name -> google.protobuf.Duration
	1, // 2: gloo.solo.io.LoadBalancerConfig.round_robin:type_name -> gloo.solo.io.LoadBalancerConfig.RoundRobin
	2, // 3: gloo.solo.io.LoadBalancerConfig.least_request:type_name -> gloo.solo.io.LoadBalancerConfig.LeastRequest
	3, // 4: gloo.solo.io.LoadBalancerConfig.random:type_name -> gloo.solo.io.LoadBalancerConfig.Random
	5, // 5: gloo.solo.io.LoadBalancerConfig.ring_hash:type_name -> gloo.solo.io.LoadBalancerConfig.RingHash
	6, // 6: gloo.solo.io.LoadBalancerConfig.maglev:type_name -> gloo.solo.io.LoadBalancerConfig.Maglev
	9, // 7: gloo.solo.io.LoadBalancerConfig.locality_weighted_lb_config:type_name -> google.protobuf.Empty
	4, // 8: gloo.solo.io.LoadBalancerConfig.RingHash.ring_hash_config:type_name -> gloo.solo.io.LoadBalancerConfig.RingHashConfig
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_init() }
func file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_init() {
	if File_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoadBalancerConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoadBalancerConfig_RoundRobin); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoadBalancerConfig_LeastRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoadBalancerConfig_Random); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoadBalancerConfig_RingHashConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoadBalancerConfig_RingHash); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoadBalancerConfig_Maglev); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*LoadBalancerConfig_RoundRobin_)(nil),
		(*LoadBalancerConfig_LeastRequest_)(nil),
		(*LoadBalancerConfig_Random_)(nil),
		(*LoadBalancerConfig_RingHash_)(nil),
		(*LoadBalancerConfig_Maglev_)(nil),
		(*LoadBalancerConfig_LocalityWeightedLbConfig)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto = out.File
	file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_rawDesc = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_load_balancer_proto_depIdxs = nil
}
