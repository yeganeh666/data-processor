package errors

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"testing"
)

func TestCast(t *testing.T) {
	tests := []struct {
		ctx    context.Context
		err    error
		result *Error
	}{
		{context.Background(), nil, nil},
		{context.Background(), New(context.Background(), codes.NotFound), New(context.Background(), codes.NotFound)},
		{context.Background(), errors.New("its test"), New(context.Background(), codes.Unknown).AddDetails("its test")},
	}

	for i := range tests {
		res := Cast(tests[i].ctx, tests[i].err)
		assert.Equal(t, tests[i].result, res)
	}
}
