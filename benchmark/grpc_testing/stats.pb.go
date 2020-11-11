// Copyright 2016 gRPC authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.3.0
// source: benchmark/grpc_testing/stats.proto

package grpc_testing

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ServerStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// wall clock time change in seconds since last reset
	TimeElapsed float64 `protobuf:"fixed64,1,opt,name=time_elapsed,json=timeElapsed,proto3" json:"time_elapsed,omitempty"`
	// change in user time (in seconds) used by the server since last reset
	TimeUser float64 `protobuf:"fixed64,2,opt,name=time_user,json=timeUser,proto3" json:"time_user,omitempty"`
	// change in server time (in seconds) used by the server process and all
	// threads since last reset
	TimeSystem float64 `protobuf:"fixed64,3,opt,name=time_system,json=timeSystem,proto3" json:"time_system,omitempty"`
}

func (x *ServerStats) Reset() {
	*x = ServerStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_benchmark_grpc_testing_stats_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerStats) ProtoMessage() {}

func (x *ServerStats) ProtoReflect() protoreflect.Message {
	mi := &file_benchmark_grpc_testing_stats_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerStats.ProtoReflect.Descriptor instead.
func (*ServerStats) Descriptor() ([]byte, []int) {
	return file_benchmark_grpc_testing_stats_proto_rawDescGZIP(), []int{0}
}

func (x *ServerStats) GetTimeElapsed() float64 {
	if x != nil {
		return x.TimeElapsed
	}
	return 0
}

func (x *ServerStats) GetTimeUser() float64 {
	if x != nil {
		return x.TimeUser
	}
	return 0
}

func (x *ServerStats) GetTimeSystem() float64 {
	if x != nil {
		return x.TimeSystem
	}
	return 0
}

// Histogram params based on grpc/support/histogram.c
type HistogramParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resolution  float64 `protobuf:"fixed64,1,opt,name=resolution,proto3" json:"resolution,omitempty"`                      // first bucket is [0, 1 + resolution)
	MaxPossible float64 `protobuf:"fixed64,2,opt,name=max_possible,json=maxPossible,proto3" json:"max_possible,omitempty"` // use enough buckets to allow this value
}

func (x *HistogramParams) Reset() {
	*x = HistogramParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_benchmark_grpc_testing_stats_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HistogramParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HistogramParams) ProtoMessage() {}

func (x *HistogramParams) ProtoReflect() protoreflect.Message {
	mi := &file_benchmark_grpc_testing_stats_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HistogramParams.ProtoReflect.Descriptor instead.
func (*HistogramParams) Descriptor() ([]byte, []int) {
	return file_benchmark_grpc_testing_stats_proto_rawDescGZIP(), []int{1}
}

func (x *HistogramParams) GetResolution() float64 {
	if x != nil {
		return x.Resolution
	}
	return 0
}

func (x *HistogramParams) GetMaxPossible() float64 {
	if x != nil {
		return x.MaxPossible
	}
	return 0
}

// Histogram data based on grpc/support/histogram.c
type HistogramData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bucket       []uint32 `protobuf:"varint,1,rep,packed,name=bucket,proto3" json:"bucket,omitempty"`
	MinSeen      float64  `protobuf:"fixed64,2,opt,name=min_seen,json=minSeen,proto3" json:"min_seen,omitempty"`
	MaxSeen      float64  `protobuf:"fixed64,3,opt,name=max_seen,json=maxSeen,proto3" json:"max_seen,omitempty"`
	Sum          float64  `protobuf:"fixed64,4,opt,name=sum,proto3" json:"sum,omitempty"`
	SumOfSquares float64  `protobuf:"fixed64,5,opt,name=sum_of_squares,json=sumOfSquares,proto3" json:"sum_of_squares,omitempty"`
	Count        float64  `protobuf:"fixed64,6,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *HistogramData) Reset() {
	*x = HistogramData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_benchmark_grpc_testing_stats_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HistogramData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HistogramData) ProtoMessage() {}

func (x *HistogramData) ProtoReflect() protoreflect.Message {
	mi := &file_benchmark_grpc_testing_stats_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HistogramData.ProtoReflect.Descriptor instead.
func (*HistogramData) Descriptor() ([]byte, []int) {
	return file_benchmark_grpc_testing_stats_proto_rawDescGZIP(), []int{2}
}

func (x *HistogramData) GetBucket() []uint32 {
	if x != nil {
		return x.Bucket
	}
	return nil
}

func (x *HistogramData) GetMinSeen() float64 {
	if x != nil {
		return x.MinSeen
	}
	return 0
}

func (x *HistogramData) GetMaxSeen() float64 {
	if x != nil {
		return x.MaxSeen
	}
	return 0
}

func (x *HistogramData) GetSum() float64 {
	if x != nil {
		return x.Sum
	}
	return 0
}

func (x *HistogramData) GetSumOfSquares() float64 {
	if x != nil {
		return x.SumOfSquares
	}
	return 0
}

func (x *HistogramData) GetCount() float64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type ClientStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Latency histogram. Data points are in nanoseconds.
	Latencies *HistogramData `protobuf:"bytes,1,opt,name=latencies,proto3" json:"latencies,omitempty"`
	// See ServerStats for details.
	TimeElapsed float64 `protobuf:"fixed64,2,opt,name=time_elapsed,json=timeElapsed,proto3" json:"time_elapsed,omitempty"`
	TimeUser    float64 `protobuf:"fixed64,3,opt,name=time_user,json=timeUser,proto3" json:"time_user,omitempty"`
	TimeSystem  float64 `protobuf:"fixed64,4,opt,name=time_system,json=timeSystem,proto3" json:"time_system,omitempty"`
}

func (x *ClientStats) Reset() {
	*x = ClientStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_benchmark_grpc_testing_stats_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientStats) ProtoMessage() {}

func (x *ClientStats) ProtoReflect() protoreflect.Message {
	mi := &file_benchmark_grpc_testing_stats_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientStats.ProtoReflect.Descriptor instead.
func (*ClientStats) Descriptor() ([]byte, []int) {
	return file_benchmark_grpc_testing_stats_proto_rawDescGZIP(), []int{3}
}

func (x *ClientStats) GetLatencies() *HistogramData {
	if x != nil {
		return x.Latencies
	}
	return nil
}

func (x *ClientStats) GetTimeElapsed() float64 {
	if x != nil {
		return x.TimeElapsed
	}
	return 0
}

func (x *ClientStats) GetTimeUser() float64 {
	if x != nil {
		return x.TimeUser
	}
	return 0
}

func (x *ClientStats) GetTimeSystem() float64 {
	if x != nil {
		return x.TimeSystem
	}
	return 0
}

var File_benchmark_grpc_testing_stats_proto protoreflect.FileDescriptor

var file_benchmark_grpc_testing_stats_proto_rawDesc = []byte{
	0x0a, 0x22, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x5f, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x22, 0x6e, 0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x65, 0x6c, 0x61, 0x70, 0x73, 0x65,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x45, 0x6c, 0x61,
	0x70, 0x73, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x75, 0x73, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x22, 0x54, 0x0a, 0x0f, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x6c,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x61, 0x78, 0x5f, 0x70, 0x6f, 0x73,
	0x73, 0x69, 0x62, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x6d, 0x61, 0x78,
	0x50, 0x6f, 0x73, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x22, 0xab, 0x01, 0x0a, 0x0d, 0x48, 0x69, 0x73,
	0x74, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x06, 0x62, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x69, 0x6e, 0x5f, 0x73, 0x65, 0x65, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x6d, 0x69, 0x6e, 0x53, 0x65, 0x65, 0x6e, 0x12, 0x19, 0x0a,
	0x08, 0x6d, 0x61, 0x78, 0x5f, 0x73, 0x65, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x07, 0x6d, 0x61, 0x78, 0x53, 0x65, 0x65, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x75, 0x6d, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x73, 0x75, 0x6d, 0x12, 0x24, 0x0a, 0x0e, 0x73, 0x75,
	0x6d, 0x5f, 0x6f, 0x66, 0x5f, 0x73, 0x71, 0x75, 0x61, 0x72, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x0c, 0x73, 0x75, 0x6d, 0x4f, 0x66, 0x53, 0x71, 0x75, 0x61, 0x72, 0x65, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xa9, 0x01, 0x0a, 0x0b, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x39, 0x0a, 0x09, 0x6c, 0x61, 0x74, 0x65, 0x6e, 0x63,
	0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x67, 0x72,
	0x61, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x52, 0x09, 0x6c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x69, 0x65,
	0x73, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x65, 0x6c, 0x61, 0x70, 0x73, 0x65,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x45, 0x6c, 0x61,
	0x70, 0x73, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x75, 0x73, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x62, 0x65, 0x6e,
	0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x74, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_benchmark_grpc_testing_stats_proto_rawDescOnce sync.Once
	file_benchmark_grpc_testing_stats_proto_rawDescData = file_benchmark_grpc_testing_stats_proto_rawDesc
)

func file_benchmark_grpc_testing_stats_proto_rawDescGZIP() []byte {
	file_benchmark_grpc_testing_stats_proto_rawDescOnce.Do(func() {
		file_benchmark_grpc_testing_stats_proto_rawDescData = protoimpl.X.CompressGZIP(file_benchmark_grpc_testing_stats_proto_rawDescData)
	})
	return file_benchmark_grpc_testing_stats_proto_rawDescData
}

var file_benchmark_grpc_testing_stats_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_benchmark_grpc_testing_stats_proto_goTypes = []interface{}{
	(*ServerStats)(nil),     // 0: grpc.testing.ServerStats
	(*HistogramParams)(nil), // 1: grpc.testing.HistogramParams
	(*HistogramData)(nil),   // 2: grpc.testing.HistogramData
	(*ClientStats)(nil),     // 3: grpc.testing.ClientStats
}
var file_benchmark_grpc_testing_stats_proto_depIdxs = []int32{
	2, // 0: grpc.testing.ClientStats.latencies:type_name -> grpc.testing.HistogramData
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_benchmark_grpc_testing_stats_proto_init() }
func file_benchmark_grpc_testing_stats_proto_init() {
	if File_benchmark_grpc_testing_stats_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_benchmark_grpc_testing_stats_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerStats); i {
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
		file_benchmark_grpc_testing_stats_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HistogramParams); i {
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
		file_benchmark_grpc_testing_stats_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HistogramData); i {
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
		file_benchmark_grpc_testing_stats_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientStats); i {
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
			RawDescriptor: file_benchmark_grpc_testing_stats_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_benchmark_grpc_testing_stats_proto_goTypes,
		DependencyIndexes: file_benchmark_grpc_testing_stats_proto_depIdxs,
		MessageInfos:      file_benchmark_grpc_testing_stats_proto_msgTypes,
	}.Build()
	File_benchmark_grpc_testing_stats_proto = out.File
	file_benchmark_grpc_testing_stats_proto_rawDesc = nil
	file_benchmark_grpc_testing_stats_proto_goTypes = nil
	file_benchmark_grpc_testing_stats_proto_depIdxs = nil
}
