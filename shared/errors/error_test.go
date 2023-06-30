package errors

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"strings"
	"testing"
)

func TestError_MarshalJSON(t *testing.T) {
	tests := []struct {
		err          Error
		result       []byte
		expectsError bool
	}{
		{*New(context.Background(), codes.NotFound), []byte(fmt.Sprintf(`{"code":%v,"message":"%v"}`, int(codes.NotFound), CodeToSlug(codes.NotFound))), false},
		{*New(context.Background(), codes.NotFound).AddDetails("test"), []byte(fmt.Sprintf(`{"code":%v,"message":"%v","details":["test"]}`, int(codes.NotFound), CodeToSlug(codes.NotFound))), false},
		{*New(context.Background(), codes.NotFound).AddDetails("test").AddDetails("test2"), []byte(fmt.Sprintf(`{"code":%v,"message":"%v","details":["test","test2"]}`, int(codes.NotFound), CodeToSlug(codes.NotFound))), false},
	}

	for i := range tests {
		result, err := tests[i].err.MarshalJSON()
		if tests[i].expectsError {
			assert.NotNil(t, err)
			continue
		}
		assert.Nil(t, err)
		assert.Equal(t, tests[i].result, result)
	}
}

func TestError_Message(t *testing.T) {
	tests := []struct {
		err     Error
		message string
	}{
		{*New(context.Background(), codes.OK), CodeToSlug(codes.OK)},
		{*New(context.Background(), codes.NotFound), CodeToSlug(codes.NotFound)},
	}

	for i := range tests {
		assert.Equal(t, tests[i].message, tests[i].err.Message())
	}
}

func TestError_Details(t *testing.T) {
	tests := []struct {
		err     Error
		details []string
	}{
		{*New(context.Background(), codes.OK), nil},
		{*New(context.Background(), codes.NotFound).AddDetails("test"), []string{"test"}},
		{*New(context.Background(), codes.NotFound).AddDetails("test").AddDetails("test2"), []string{"test", "test2"}},
	}

	for i := range tests {
		assert.Equal(t, tests[i].details, tests[i].err.Details())
	}
}

func TestError_Error(t *testing.T) {
	tests := []struct {
		err           Error
		errorMessages []string
	}{
		{*New(context.Background(), codes.OK), []string{CodeToSlug(codes.OK)}},
		{*New(context.Background(), codes.NotFound).AddDetails("test"), []string{CodeToSlug(codes.NotFound), "test"}},
		{*New(context.Background(), codes.NotFound).AddDetails("test").AddDetails("test2"), []string{CodeToSlug(codes.NotFound), "test", "test2"}},
	}

	for i := range tests {
		for j := range tests[i].errorMessages {
			assert.True(t, strings.Contains(tests[i].err.Error(), tests[i].errorMessages[j]))
		}
	}
}

func TestError_GRPCStatus(t *testing.T) {
	tests := []struct {
		err  Error
		code codes.Code
	}{
		{*New(context.Background(), codes.OK), codes.OK},
		{*New(context.Background(), codes.NotFound), codes.NotFound},
		{*New(context.Background(), codes.Unavailable), codes.Unavailable},
	}

	for i := range tests {
		assert.Equal(t, tests[i].code, tests[i].err.GRPCStatus().Code())
	}
}

func TestError_Code(t *testing.T) {
	tests := []struct {
		err  Error
		code codes.Code
	}{
		{*New(context.Background(), codes.OK), codes.OK},
		{*New(context.Background(), codes.NotFound), codes.NotFound},
		{*New(context.Background(), codes.NotFound).AddDetails("test"), codes.NotFound},
	}

	for i := range tests {
		assert.Equal(t, tests[i].err.Code(), tests[i].code)
	}
}

func TestError_HttpStatus(t *testing.T) {
	tests := []struct {
		err  Error
		code int
	}{
		{*New(context.Background(), codes.OK), runtime.HTTPStatusFromCode(codes.OK)},
		{*New(context.Background(), codes.NotFound), runtime.HTTPStatusFromCode(codes.NotFound)},
		{*New(context.Background(), codes.Aborted), runtime.HTTPStatusFromCode(codes.Aborted)},
	}

	for i := range tests {
		assert.Equal(t, tests[i].err.HttpStatus(), tests[i].code)
	}
}

func TestError_IsHandled(t *testing.T) {
	tests := []struct {
		err     Error
		handled bool
	}{
		{*New(context.Background(), codes.OK), IsHandledCode(codes.OK)},
		{*New(context.Background(), codes.NotFound), IsHandledCode(codes.NotFound)},
		{*New(context.Background(), codes.Aborted), IsHandledCode(codes.Aborted)},
		{*New(context.Background(), codes.Internal), IsHandledCode(codes.Internal)},
	}

	for i := range tests {
		assert.Equal(t, tests[i].err.IsHandled(), tests[i].handled)
	}
}

func TestError_AddDetails(t *testing.T) {
	tests := []struct {
		err     Error
		details []string
	}{
		{*New(context.Background(), codes.OK), nil},
		{*New(context.Background(), codes.NotFound).AddDetails("test"), []string{"test"}},
		{*New(context.Background(), codes.NotFound).AddDetails("test").AddDetails("test2"), []string{"test", "test2"}},
	}

	for i := range tests {
		assert.Equal(t, tests[i].details, tests[i].err.Details())
	}
}

func TestError_AddDetailF(t *testing.T) {
	tests := []struct {
		err     Error
		details []string
	}{
		{*New(context.Background(), codes.OK), nil},
		{*New(context.Background(), codes.NotFound).AddDetailF("test%v", 1), []string{"test1"}},
		{*New(context.Background(), codes.NotFound).AddDetails("test").AddDetailF("test%v", 2), []string{"test", "test2"}},
	}

	for i := range tests {
		assert.Equal(t, tests[i].details, tests[i].err.Details())
	}
}
