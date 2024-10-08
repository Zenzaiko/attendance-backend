// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: attendance/v1/attendance.proto

package attendancev1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/Zenzaiko/attendance-backend/attendance/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// AttendanceServiceName is the fully-qualified name of the AttendanceService service.
	AttendanceServiceName = "attendance.v1.AttendanceService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// AttendanceServiceRegisterClockInOutProcedure is the fully-qualified name of the
	// AttendanceService's RegisterClockInOut RPC.
	AttendanceServiceRegisterClockInOutProcedure = "/attendance.v1.AttendanceService/RegisterClockInOut"
	// AttendanceServiceGetAttendanceListProcedure is the fully-qualified name of the
	// AttendanceService's GetAttendanceList RPC.
	AttendanceServiceGetAttendanceListProcedure = "/attendance.v1.AttendanceService/GetAttendanceList"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	attendanceServiceServiceDescriptor                  = v1.File_attendance_v1_attendance_proto.Services().ByName("AttendanceService")
	attendanceServiceRegisterClockInOutMethodDescriptor = attendanceServiceServiceDescriptor.Methods().ByName("RegisterClockInOut")
	attendanceServiceGetAttendanceListMethodDescriptor  = attendanceServiceServiceDescriptor.Methods().ByName("GetAttendanceList")
)

// AttendanceServiceClient is a client for the attendance.v1.AttendanceService service.
type AttendanceServiceClient interface {
	RegisterClockInOut(context.Context, *connect.Request[v1.RegisterClockInOutRequest]) (*connect.Response[v1.RegisterClockInOutResponse], error)
	GetAttendanceList(context.Context, *connect.Request[v1.GetAttendanceListRequest]) (*connect.Response[v1.GetAttendanceListResponse], error)
}

// NewAttendanceServiceClient constructs a client for the attendance.v1.AttendanceService service.
// By default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped
// responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewAttendanceServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) AttendanceServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &attendanceServiceClient{
		registerClockInOut: connect.NewClient[v1.RegisterClockInOutRequest, v1.RegisterClockInOutResponse](
			httpClient,
			baseURL+AttendanceServiceRegisterClockInOutProcedure,
			connect.WithSchema(attendanceServiceRegisterClockInOutMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getAttendanceList: connect.NewClient[v1.GetAttendanceListRequest, v1.GetAttendanceListResponse](
			httpClient,
			baseURL+AttendanceServiceGetAttendanceListProcedure,
			connect.WithSchema(attendanceServiceGetAttendanceListMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// attendanceServiceClient implements AttendanceServiceClient.
type attendanceServiceClient struct {
	registerClockInOut *connect.Client[v1.RegisterClockInOutRequest, v1.RegisterClockInOutResponse]
	getAttendanceList  *connect.Client[v1.GetAttendanceListRequest, v1.GetAttendanceListResponse]
}

// RegisterClockInOut calls attendance.v1.AttendanceService.RegisterClockInOut.
func (c *attendanceServiceClient) RegisterClockInOut(ctx context.Context, req *connect.Request[v1.RegisterClockInOutRequest]) (*connect.Response[v1.RegisterClockInOutResponse], error) {
	return c.registerClockInOut.CallUnary(ctx, req)
}

// GetAttendanceList calls attendance.v1.AttendanceService.GetAttendanceList.
func (c *attendanceServiceClient) GetAttendanceList(ctx context.Context, req *connect.Request[v1.GetAttendanceListRequest]) (*connect.Response[v1.GetAttendanceListResponse], error) {
	return c.getAttendanceList.CallUnary(ctx, req)
}

// AttendanceServiceHandler is an implementation of the attendance.v1.AttendanceService service.
type AttendanceServiceHandler interface {
	RegisterClockInOut(context.Context, *connect.Request[v1.RegisterClockInOutRequest]) (*connect.Response[v1.RegisterClockInOutResponse], error)
	GetAttendanceList(context.Context, *connect.Request[v1.GetAttendanceListRequest]) (*connect.Response[v1.GetAttendanceListResponse], error)
}

// NewAttendanceServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewAttendanceServiceHandler(svc AttendanceServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	attendanceServiceRegisterClockInOutHandler := connect.NewUnaryHandler(
		AttendanceServiceRegisterClockInOutProcedure,
		svc.RegisterClockInOut,
		connect.WithSchema(attendanceServiceRegisterClockInOutMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	attendanceServiceGetAttendanceListHandler := connect.NewUnaryHandler(
		AttendanceServiceGetAttendanceListProcedure,
		svc.GetAttendanceList,
		connect.WithSchema(attendanceServiceGetAttendanceListMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/attendance.v1.AttendanceService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case AttendanceServiceRegisterClockInOutProcedure:
			attendanceServiceRegisterClockInOutHandler.ServeHTTP(w, r)
		case AttendanceServiceGetAttendanceListProcedure:
			attendanceServiceGetAttendanceListHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedAttendanceServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedAttendanceServiceHandler struct{}

func (UnimplementedAttendanceServiceHandler) RegisterClockInOut(context.Context, *connect.Request[v1.RegisterClockInOutRequest]) (*connect.Response[v1.RegisterClockInOutResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("attendance.v1.AttendanceService.RegisterClockInOut is not implemented"))
}

func (UnimplementedAttendanceServiceHandler) GetAttendanceList(context.Context, *connect.Request[v1.GetAttendanceListRequest]) (*connect.Response[v1.GetAttendanceListResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("attendance.v1.AttendanceService.GetAttendanceList is not implemented"))
}
