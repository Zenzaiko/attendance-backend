package main

import (
	"context"
	"log"
	"net/http"

	"connectrpc.com/connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	attendancev1 "github.com/Zenzaiko/attendance-backend/gen/attendance/v1"        // generated by protoc-gen-go
	"github.com/Zenzaiko/attendance-backend/gen/attendance/v1/attendancev1connect" // generated by protoc-gen-connect-go
)

type AttendanceServer struct{}

func (s *AttendanceServer) GetAttendance(ctx context.Context, req *connect.Request[attendancev1.GetAttendanceListRequest]) (*connect.Response[attendancev1.GetAttendanceListResponse], error) {
	log.Println("Request Headers: ", req.Header())
	res := connect.NewResponse(&attendancev1.GetAttendanceListResponse{
		List: []*attendancev1.AttendanceInfo{
			&attendancev1.AttendanceInfo{
				UserId:       "test",
				UserName:     "test",
				ClockInTime:  "test",
				ClockOutTime: "test",
			},
		},
	})
	return res, nil
}

func main() {
	attendancer := &AttendanceServer{}
	mux := http.NewServeMux()
	path, handler := attendancev1connect.NewAttendanceServiceHandler(attendancer)
	mux.Handle(path, handler)
	http.ListenAndServe("localhost:8080", h2c.NewHandler(mux, &http2.Server{}))
}
