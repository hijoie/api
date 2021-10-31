package response

import (
	"api/pb/bash"
	"api/pkg/bizerr"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func ErrorHandle(err error) error {
	if err == nil {
		return err
	}
	var t = &bizerr.Error{}
	if errors.As(err, t) {
		return status.Error(codes.Code(t.Code), t.Msg)
	}
	return status.Error(codes.Internal, err.Error())
}

func Ok() error {
	return status.Error(codes.OK, "ok")
}

func Response(data proto.Message) *bash.Response {
	d, err := anypb.New(data)
	if err != nil {
		panic(err)
	}
	return &bash.Response{Data: d}
}
