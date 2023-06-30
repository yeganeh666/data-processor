package contextext

import (
	"context"
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"
	"testing"
)

func TestSetValue(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		key   string
		value string
	}{
		{"key", "value"},
		{"key", "new-value"},
		{"new-key", "value"},
	}

	for i := range tests {
		ctx = SetValue(ctx, tests[i].key, tests[i].value)
		assert.Equal(t, tests[i].value, ctx.Value(tests[i].key).(string))
	}
}

func TestGetValue(t *testing.T) {
	tests := []struct {
		key    string
		value  string
		exists bool
		ctx    context.Context
	}{
		{"key", "value", true, context.WithValue(context.Background(), "key", "value")},
		{"key", "", true, context.WithValue(context.Background(), "key", "")},
		{"not-exists", "", false, context.Background()},
	}

	for i := range tests {
		value, exists := GetValue(tests[i].ctx, tests[i].key)
		assert.Equal(t, tests[i].value, value)
		assert.Equal(t, tests[i].exists, exists)
	}
}

func TestSetToken(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		token string
	}{
		{"token"},
		{"new-token"},
	}

	for i := range tests {
		ctx = SetToken(ctx, tests[i].token)
		token, exists := GetValue(ctx, "token")
		assert.Equal(t, tests[i].token, token)
		assert.True(t, exists)
	}
}

func TestGetToken(t *testing.T) {
	tests := []struct {
		token  string
		exists bool
		ctx    context.Context
	}{
		{"token", true, SetToken(context.Background(), "token")},
		{"", false, context.Background()},
	}

	for i := range tests {
		token, exists := GetToken(tests[i].ctx)
		assert.Equal(t, tests[i].token, token)
		assert.Equal(t, tests[i].exists, exists)
	}
}

func TestSetLang(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		lang language.Tag
	}{
		{language.Make("en")},
		{language.Make("fa")},
	}

	for i := range tests {
		ctx = SetLang(ctx, tests[i].lang)
		lang, exists := GetValue(ctx, "lang")
		assert.Equal(t, tests[i].lang.String(), lang)
		assert.True(t, exists)
	}
}

func TestGetLang(t *testing.T) {
	tests := []struct {
		lang   language.Tag
		exists bool
		ctx    context.Context
	}{
		{language.Make("en"), true, SetLang(context.Background(), language.Make("en"))},
		{language.Make(""), false, context.Background()},
	}

	for i := range tests {
		lang, exists := GetLang(tests[i].ctx)
		assert.Equal(t, tests[i].lang, lang)
		assert.Equal(t, tests[i].exists, exists)
	}
}

func TestGetLangWithDefault(t *testing.T) {
	tests := []struct {
		lang language.Tag
		ctx  context.Context
	}{
		{language.Make("en"), SetLang(context.Background(), language.Make("en"))},
		{language.Make("en"), context.Background()},
	}

	for i := range tests {
		lang := GetLangWithDefault(tests[i].ctx, tests[i].lang)
		assert.Equal(t, tests[i].lang, lang)
	}
}
