// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo/projects/gloo/api/v1/connection.proto

package v1

import (
	duration "github.com/golang/protobuf/ptypes/duration"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Action to take when Envoy receives client request with header names containing underscore
// characters.
// Underscore character is allowed in header names by the RFC-7230 and this behavior is implemented
// as a security measure due to systems that treat '_' and '-' as interchangeable. Envoy by default allows client request headers with underscore
// characters.
type ConnectionConfig_HttpProtocolOptions_HeadersWithUnderscoresAction int32

const (
	// Allow headers with underscores. This is the default behavior.
	ConnectionConfig_HttpProtocolOptions_ALLOW ConnectionConfig_HttpProtocolOptions_HeadersWithUnderscoresAction = 0
	// Reject client request. HTTP/1 requests are rejected with the 400 status. HTTP/2 requests
	// end with the stream reset. The "httpN.requests_rejected_with_underscores_in_headers" counter
	// is incremented for each rejected request.
	ConnectionConfig_HttpProtocolOptions_REJECT_REQUEST ConnectionConfig_HttpProtocolOptions_HeadersWithUnderscoresAction = 1
	// Drop the header with name containing underscores. The header is dropped before the filter chain is
	// invoked and as such filters will not see dropped headers. The
	// "httpN.dropped_headers_with_underscores" is incremented for each dropped header.
	ConnectionConfig_HttpProtocolOptions_DROP_HEADER ConnectionConfig_HttpProtocolOptions_HeadersWithUnderscoresAction = 2
)

// Enum value maps for ConnectionConfig_HttpProtocolOptions_HeadersWithUnderscoresAction.
var (
	ConnectionConfig_HttpProtocolOptions_HeadersWithUnderscoresAction_name = map[int32]string{
		0: "ALLOW",
		1: "REJECT_REQUEST",
		2: "DROP_HEADER",
	}
	ConnectionConfig_HttpProtocolOptions_HeadersWithUnderscoresAction_value = map[string]int32{
		"ALLOW":          0,
		"REJECT_REQUEST": 1,
		"DROP_HEADER":    2,
	}
)

func (x ConnectionConfig_HttpProtocolOptions_HeadersWithUnderscoresAction) Enum() *ConnectionConfig_HttpProtocolOptions_HeadersWithUnderscoresAction {
	p := new(ConnectionConfig_HttpProtocolOptions_HeadersWithUnderscoresAction)
	*p = x
	return p
}

func (x ConnectionConfig_HttpProtocolOptions_HeadersWithUnderscoresAction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConnectionConfig_HttpProtocolOptions_HeadersWithUnderscoresAction) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_connection_proto_enumTypes[0].Descriptor()
}

func (ConnectionConfig_HttpProtocolOptions_HeadersWithUnderscoresAction) Type() protoreflect.EnumType {
	return &file_github_com_solo_io_gloo_projects_gloo_api_v1_connection_proto_enumTypes[0]
}

func (x ConnectionConfig_HttpProtocolOptions_HeadersWithUnderscoresAction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConnectionConfig_HttpProtocolOptions_HeadersWithUnderscoresAction.Descriptor instead.
func (ConnectionConfig_HttpProtocolOptions_HeadersWithUnderscoresAction) EnumDescriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_connection_proto_rawDescGZIP(), []int{0, 1, 0}
}

// Fine tune the settings for connections to an upstream
type ConnectionConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Maximum requests for a single upstream connection (unspecified or zero = no limit)
	MaxRequestsPerConnection uint32 `protobuf:"varint,1,opt,name=max_requests_per_connection,json=maxRequestsPerConnection,proto3" json:"max_requests_per_connection,omitempty"`
	// The timeout for new network connections to hosts in the cluster
	ConnectTimeout *duration.Duration `protobuf:"bytes,2,opt,name=connect_timeout,json=connectTimeout,proto3" json:"connect_timeout,omitempty"`
	// Configure OS-level tcp keepalive checks
	TcpKeepalive *ConnectionConfig_TcpKeepAlive `protobuf:"bytes,3,opt,name=tcp_keepalive,json=tcpKeepalive,proto3" json:"tcp_keepalive,omitempty"`
	// Soft limit on size of the cluster’s connections read and write buffers. If unspecified, an implementation defined default is applied (1MiB).
	// For more info, see the [envoy docs](https://www.envoyproxy.io/docs/envoy/v1.14.1/api-v2/api/v2/cluster.proto#cluster)
	PerConnectionBufferLimitBytes *wrappers.UInt32Value `protobuf:"bytes,4,opt,name=per_connection_buffer_limit_bytes,json=perConnectionBufferLimitBytes,proto3" json:"per_connection_buffer_limit_bytes,omitempty"`
	// Additional options when handling HTTP requests upstream. These options will be applicable to
	// both HTTP1 and HTTP2 requests.
	CommonHttpProtocolOptions *ConnectionConfig_HttpProtocolOptions `protobuf:"bytes,5,opt,name=common_http_protocol_options,json=commonHttpProtocolOptions,proto3" json:"common_http_protocol_options,omitempty"`
}

func (x *ConnectionConfig) Reset() {
	*x = ConnectionConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_connection_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectionConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionConfig) ProtoMessage() {}

func (x *ConnectionConfig) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_connection_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionConfig.ProtoReflect.Descriptor instead.
func (*ConnectionConfig) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_connection_proto_rawDescGZIP(), []int{0}
}

func (x *ConnectionConfig) GetMaxRequestsPerConnection() uint32 {
	if x != nil {
		return x.MaxRequestsPerConnection
	}
	return 0
}

func (x *ConnectionConfig) GetConnectTimeout() *duration.Duration {
	if x != nil {
		return x.ConnectTimeout
	}
	return nil
}

func (x *ConnectionConfig) GetTcpKeepalive() *ConnectionConfig_TcpKeepAlive {
	if x != nil {
		return x.TcpKeepalive
	}
	return nil
}

func (x *ConnectionConfig) GetPerConnectionBufferLimitBytes() *wrappers.UInt32Value {
	if x != nil {
		return x.PerConnectionBufferLimitBytes
	}
	return nil
}

func (x *ConnectionConfig) GetCommonHttpProtocolOptions() *ConnectionConfig_HttpProtocolOptions {
	if x != nil {
		return x.CommonHttpProtocolOptions
	}
	return nil
}

// If set then set SO_KEEPALIVE on the socket to enable TCP Keepalives.
// see more info here: https://www.envoyproxy.io/docs/envoy/latest/api-v2/api/v2/core/address.proto#envoy-api-msg-core-tcpkeepalive
type ConnectionConfig_TcpKeepAlive struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Maximum number of keepalive probes to send without response before deciding the connection is dead.
	KeepaliveProbes uint32 `protobuf:"varint,1,opt,name=keepalive_probes,json=keepaliveProbes,proto3" json:"keepalive_probes,omitempty"`
	// The number of seconds a connection needs to be idle before keep-alive probes start being sent. This is rounded up to the second.
	KeepaliveTime *duration.Duration `protobuf:"bytes,2,opt,name=keepalive_time,json=keepaliveTime,proto3" json:"keepalive_time,omitempty"`
	// The number of seconds between keep-alive probes. This is rounded up to the second.
	KeepaliveInterval *duration.Duration `protobuf:"bytes,3,opt,name=keepalive_interval,json=keepaliveInterval,proto3" json:"keepalive_interval,omitempty"`
}

func (x *ConnectionConfig_TcpKeepAlive) Reset() {
	*x = ConnectionConfig_TcpKeepAlive{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_connection_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectionConfig_TcpKeepAlive) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionConfig_TcpKeepAlive) ProtoMessage() {}

func (x *ConnectionConfig_TcpKeepAlive) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_connection_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionConfig_TcpKeepAlive.ProtoReflect.Descriptor instead.
func (*ConnectionConfig_TcpKeepAlive) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_connection_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ConnectionConfig_TcpKeepAlive) GetKeepaliveProbes() uint32 {
	if x != nil {
		return x.KeepaliveProbes
	}
	return 0
}

func (x *ConnectionConfig_TcpKeepAlive) GetKeepaliveTime() *duration.Duration {
	if x != nil {
		return x.KeepaliveTime
	}
	return nil
}

func (x *ConnectionConfig_TcpKeepAlive) GetKeepaliveInterval() *duration.Duration {
	if x != nil {
		return x.KeepaliveInterval
	}
	return nil
}

type ConnectionConfig_HttpProtocolOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The idle timeout for connections. The idle timeout is defined as the
	// period in which there are no active requests. When the
	// idle timeout is reached the connection will be closed. If the connection is an HTTP/2
	// downstream connection a drain sequence will occur prior to closing the connection, see
	// :ref:`drain_timeout
	// <envoy_api_field_extensions.filters.network.http_connection_manager.v3.HttpConnectionManager.drain_timeout>`.
	// Note that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive.
	// If not specified, this defaults to 1 hour. To disable idle timeouts explicitly set this to 0.
	//
	// .. warning::
	//   Disabling this timeout has a highly likelihood of yielding connection leaks due to lost TCP
	//   FIN packets, etc.
	IdleTimeout *duration.Duration `protobuf:"bytes,1,opt,name=idle_timeout,json=idleTimeout,proto3" json:"idle_timeout,omitempty"`
	// The maximum number of headers. If unconfigured, the default
	// maximum number of request headers allowed is 100. Requests that exceed this limit will receive
	// a 431 response for HTTP/1.x and cause a stream reset for HTTP/2.
	MaxHeadersCount uint32 `protobuf:"varint,2,opt,name=max_headers_count,json=maxHeadersCount,proto3" json:"max_headers_count,omitempty"`
	// Total duration to keep alive an HTTP request/response stream. If the time limit is reached the stream will be
	// reset independent of any other timeouts. If not specified, this value is not set.
	MaxStreamDuration *duration.Duration `protobuf:"bytes,3,opt,name=max_stream_duration,json=maxStreamDuration,proto3" json:"max_stream_duration,omitempty"`
	// Action to take when a client request with a header name containing underscore characters is received.
	// If this setting is not specified, the value defaults to ALLOW.
	// Note: upstream responses are not affected by this setting.
	HeadersWithUnderscoresAction ConnectionConfig_HttpProtocolOptions_HeadersWithUnderscoresAction `protobuf:"varint,4,opt,name=headers_with_underscores_action,json=headersWithUnderscoresAction,proto3,enum=gloo.solo.io.ConnectionConfig_HttpProtocolOptions_HeadersWithUnderscoresAction" json:"headers_with_underscores_action,omitempty"`
}

func (x *ConnectionConfig_HttpProtocolOptions) Reset() {
	*x = ConnectionConfig_HttpProtocolOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_connection_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectionConfig_HttpProtocolOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionConfig_HttpProtocolOptions) ProtoMessage() {}

func (x *ConnectionConfig_HttpProtocolOptions) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_connection_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionConfig_HttpProtocolOptions.ProtoReflect.Descriptor instead.
func (*ConnectionConfig_HttpProtocolOptions) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_connection_proto_rawDescGZIP(), []int{0, 1}
}

func (x *ConnectionConfig_HttpProtocolOptions) GetIdleTimeout() *duration.Duration {
	if x != nil {
		return x.IdleTimeout
	}
	return nil
}

func (x *ConnectionConfig_HttpProtocolOptions) GetMaxHeadersCount() uint32 {
	if x != nil {
		return x.MaxHeadersCount
	}
	return 0
}

func (x *ConnectionConfig_HttpProtocolOptions) GetMaxStreamDuration() *duration.Duration {
	if x != nil {
		return x.MaxStreamDuration
	}
	return nil
}

func (x *ConnectionConfig_HttpProtocolOptions) GetHeadersWithUnderscoresAction() ConnectionConfig_HttpProtocolOptions_HeadersWithUnderscoresAction {
	if x != nil {
		return x.HeadersWithUnderscoresAction
	}
	return ConnectionConfig_HttpProtocolOptions_ALLOW
}

var File_github_com_solo_io_gloo_projects_gloo_api_v1_connection_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_projects_gloo_api_v1_connection_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0c, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77,
	0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x65,
	0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xc2, 0x08, 0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x3d, 0x0a, 0x1b, 0x6d, 0x61, 0x78, 0x5f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x18, 0x6d, 0x61, 0x78,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x50, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x42, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x50, 0x0a, 0x0d, 0x74, 0x63, 0x70,
	0x5f, 0x6b, 0x65, 0x65, 0x70, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2b, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x54, 0x63, 0x70, 0x4b, 0x65, 0x65, 0x70, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x52, 0x0c, 0x74,
	0x63, 0x70, 0x4b, 0x65, 0x65, 0x70, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x12, 0x66, 0x0a, 0x21, 0x70,
	0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x62, 0x75,
	0x66, 0x66, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x1d, 0x70, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x42, 0x75, 0x66, 0x66, 0x65, 0x72, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x42, 0x79,
	0x74, 0x65, 0x73, 0x12, 0x73, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x68, 0x74,
	0x74, 0x70, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5f, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x67, 0x6c, 0x6f, 0x6f,
	0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x19, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x48, 0x74, 0x74, 0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0xc5, 0x01, 0x0a, 0x0c, 0x54, 0x63, 0x70,
	0x4b, 0x65, 0x65, 0x70, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x6b, 0x65, 0x65,
	0x70, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0f, 0x6b, 0x65, 0x65, 0x70, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x50, 0x72,
	0x6f, 0x62, 0x65, 0x73, 0x12, 0x40, 0x0a, 0x0e, 0x6b, 0x65, 0x65, 0x70, 0x61, 0x6c, 0x69, 0x76,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x6b, 0x65, 0x65, 0x70, 0x61, 0x6c, 0x69,
	0x76, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x48, 0x0a, 0x12, 0x6b, 0x65, 0x65, 0x70, 0x61, 0x6c,
	0x69, 0x76, 0x65, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x11, 0x6b,
	0x65, 0x65, 0x70, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c,
	0x1a, 0xb3, 0x03, 0x0a, 0x13, 0x48, 0x74, 0x74, 0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x3c, 0x0a, 0x0c, 0x69, 0x64, 0x6c, 0x65,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x69, 0x64, 0x6c, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x2a, 0x0a, 0x11, 0x6d, 0x61, 0x78, 0x5f, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0f, 0x6d, 0x61, 0x78, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x49, 0x0a, 0x13, 0x6d, 0x61, 0x78, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x11, 0x6d, 0x61, 0x78, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x96, 0x01,
	0x0a, 0x1f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x75,
	0x6e, 0x64, 0x65, 0x72, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x4f, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x73, 0x57, 0x69, 0x74, 0x68, 0x55, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x73, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x1c, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x73, 0x57, 0x69, 0x74, 0x68, 0x55, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x73,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x4e, 0x0a, 0x1c, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x73, 0x57, 0x69, 0x74, 0x68, 0x55, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x73,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x10,
	0x00, 0x12, 0x12, 0x0a, 0x0e, 0x52, 0x45, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x52, 0x45, 0x51, 0x55,
	0x45, 0x53, 0x54, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x44, 0x52, 0x4f, 0x50, 0x5f, 0x48, 0x45,
	0x41, 0x44, 0x45, 0x52, 0x10, 0x02, 0x42, 0x3a, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f,
	0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0xc0, 0xf5, 0x04, 0x01, 0xb8, 0xf5,
	0x04, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_projects_gloo_api_v1_connection_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_projects_gloo_api_v1_connection_proto_rawDescData = file_github_com_solo_io_gloo_projects_gloo_api_v1_connection_proto_rawDesc
)

func file_github_com_solo_io_gloo_projects_gloo_api_v1_connection_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_projects_gloo_api_v1_connection_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_connection_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_projects_gloo_api_v1_connection_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_connection_proto_rawDescData
}

var file_github_com_solo_io_gloo_projects_gloo_api_v1_connection_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_solo_io_gloo_projects_gloo_api_v1_connection_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_github_com_solo_io_gloo_projects_gloo_api_v1_connection_proto_goTypes = []interface{}{
	(ConnectionConfig_HttpProtocolOptions_HeadersWithUnderscoresAction)(0), // 0: gloo.solo.io.ConnectionConfig.HttpProtocolOptions.HeadersWithUnderscoresAction
	(*ConnectionConfig)(nil),                     // 1: gloo.solo.io.ConnectionConfig
	(*ConnectionConfig_TcpKeepAlive)(nil),        // 2: gloo.solo.io.ConnectionConfig.TcpKeepAlive
	(*ConnectionConfig_HttpProtocolOptions)(nil), // 3: gloo.solo.io.ConnectionConfig.HttpProtocolOptions
	(*duration.Duration)(nil),                    // 4: google.protobuf.Duration
	(*wrappers.UInt32Value)(nil),                 // 5: google.protobuf.UInt32Value
}
var file_github_com_solo_io_gloo_projects_gloo_api_v1_connection_proto_depIdxs = []int32{
	4, // 0: gloo.solo.io.ConnectionConfig.connect_timeout:type_name -> google.protobuf.Duration
	2, // 1: gloo.solo.io.ConnectionConfig.tcp_keepalive:type_name -> gloo.solo.io.ConnectionConfig.TcpKeepAlive
	5, // 2: gloo.solo.io.ConnectionConfig.per_connection_buffer_limit_bytes:type_name -> google.protobuf.UInt32Value
	3, // 3: gloo.solo.io.ConnectionConfig.common_http_protocol_options:type_name -> gloo.solo.io.ConnectionConfig.HttpProtocolOptions
	4, // 4: gloo.solo.io.ConnectionConfig.TcpKeepAlive.keepalive_time:type_name -> google.protobuf.Duration
	4, // 5: gloo.solo.io.ConnectionConfig.TcpKeepAlive.keepalive_interval:type_name -> google.protobuf.Duration
	4, // 6: gloo.solo.io.ConnectionConfig.HttpProtocolOptions.idle_timeout:type_name -> google.protobuf.Duration
	4, // 7: gloo.solo.io.ConnectionConfig.HttpProtocolOptions.max_stream_duration:type_name -> google.protobuf.Duration
	0, // 8: gloo.solo.io.ConnectionConfig.HttpProtocolOptions.headers_with_underscores_action:type_name -> gloo.solo.io.ConnectionConfig.HttpProtocolOptions.HeadersWithUnderscoresAction
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_github_com_solo_io_gloo_projects_gloo_api_v1_connection_proto_init() }
func file_github_com_solo_io_gloo_projects_gloo_api_v1_connection_proto_init() {
	if File_github_com_solo_io_gloo_projects_gloo_api_v1_connection_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_connection_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectionConfig); i {
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
		file_github_com_solo_io_gloo_projects_gloo_api_v1_connection_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectionConfig_TcpKeepAlive); i {
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
		file_github_com_solo_io_gloo_projects_gloo_api_v1_connection_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectionConfig_HttpProtocolOptions); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_projects_gloo_api_v1_connection_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_gloo_api_v1_connection_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_gloo_api_v1_connection_proto_depIdxs,
		EnumInfos:         file_github_com_solo_io_gloo_projects_gloo_api_v1_connection_proto_enumTypes,
		MessageInfos:      file_github_com_solo_io_gloo_projects_gloo_api_v1_connection_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_gloo_api_v1_connection_proto = out.File
	file_github_com_solo_io_gloo_projects_gloo_api_v1_connection_proto_rawDesc = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_connection_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_connection_proto_depIdxs = nil
}
