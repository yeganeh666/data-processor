package errors

import (
	"context"
	"errors"
	"github.com/iancoleman/strcase"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"testing"
)

func TestIsHandledCode(t *testing.T) {
	tests := []struct {
		code   codes.Code
		result bool
	}{
		{codes.OK, true},
		{codes.Canceled, true},
		{codes.InvalidArgument, true},
		{codes.NotFound, true},
		{codes.AlreadyExists, true},
		{codes.PermissionDenied, true},
		{codes.ResourceExhausted, true},
		{codes.FailedPrecondition, true},
		{codes.Aborted, true},
		{codes.OutOfRange, true},
		{codes.Unauthenticated, true},
		{codes.Unknown, false},
		{codes.DeadlineExceeded, false},
		{codes.Unimplemented, false},
		{codes.Internal, false},
		{codes.Unavailable, false},
		{codes.DataLoss, false},
	}

	for i := range tests {
		assert.Equal(t, tests[i].result, IsHandledCode(tests[i].code))
	}
}

func TestCodeToSlug(t *testing.T) {
	tests := []struct {
		code   codes.Code
		result string
	}{
		{codes.OK, strcase.ToSnake(codes.OK.String())},
		{codes.Canceled, strcase.ToSnake(codes.Canceled.String())},
		{codes.InvalidArgument, strcase.ToSnake(codes.InvalidArgument.String())},
		{codes.NotFound, strcase.ToSnake(codes.NotFound.String())},
		{codes.AlreadyExists, strcase.ToSnake(codes.AlreadyExists.String())},
		{codes.PermissionDenied, strcase.ToSnake(codes.PermissionDenied.String())},
		{codes.ResourceExhausted, strcase.ToSnake(codes.ResourceExhausted.String())},
		{codes.FailedPrecondition, strcase.ToSnake(codes.FailedPrecondition.String())},
		{codes.Aborted, strcase.ToSnake(codes.Aborted.String())},
		{codes.OutOfRange, strcase.ToSnake(codes.OutOfRange.String())},
		{codes.Unauthenticated, strcase.ToSnake(codes.Unauthenticated.String())},
		{codes.Unknown, strcase.ToSnake(codes.Unknown.String())},
		{codes.DeadlineExceeded, strcase.ToSnake(codes.DeadlineExceeded.String())},
		{codes.Unimplemented, strcase.ToSnake(codes.Unimplemented.String())},
		{codes.Internal, strcase.ToSnake(codes.Internal.String())},
		{codes.Unavailable, strcase.ToSnake(codes.Unavailable.String())},
		{codes.DataLoss, strcase.ToSnake(codes.DataLoss.String())},
	}

	for i := range tests {
		assert.Equal(t, tests[i].result, CodeToSlug(tests[i].code))
	}
}

func TestCode(t *testing.T) {
	tests := []struct {
		ctx  context.Context
		err  error
		code codes.Code
	}{
		{context.Background(), nil, codes.OK},
		{context.Background(), errors.New("test"), codes.Unknown},
		{context.Background(), New(context.Background(), codes.NotFound), codes.NotFound},
	}

	for i := range tests {
		assert.Equal(t, tests[i].code, Code(tests[i].ctx, tests[i].err))
	}
}
