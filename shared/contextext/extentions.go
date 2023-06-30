package contextext

import (
	"context"
	"golang.org/x/text/language"
	"google.golang.org/grpc/metadata"
)

func SetValue(ctx context.Context, key, value string) context.Context {
	ctx = metadata.AppendToOutgoingContext(ctx, key, value)
	ctx = context.WithValue(ctx, key, value)
	return ctx
}
func GetValue(ctx context.Context, key string) (string, bool) {
	value, ok := ctx.Value(key).(string)
	if ok {
		return value, true
	}
	var md metadata.MD
	md, ok = metadata.FromIncomingContext(ctx)
	if !ok {
		return "", false
	}
	if len(md[key]) > 0 {
		return md[key][0], true
	}
	return "", false
}

func SetToken(ctx context.Context, token string) context.Context {
	if token == "" {
		return ctx
	}
	return SetValue(ctx, "token", token)
}
func GetToken(ctx context.Context) (string, bool) {
	value, ok := GetValue(ctx, "token")
	if !ok {
		return "", false
	}
	return value, true
}

func SetLang(ctx context.Context, lang language.Tag) context.Context {
	if lang == language.Und {
		return ctx
	}
	return SetValue(ctx, "lang", lang.String())
}
func GetLang(ctx context.Context) (language.Tag, bool) {
	value, ok := GetValue(ctx, "lang")
	if !ok {
		return language.Und, false
	}
	tag, err := language.Parse(value)
	if err != nil {
		return language.Und, false
	}
	return tag, true
}
func GetLangWithDefault(ctx context.Context, def language.Tag) language.Tag {
	lang, ok := GetLang(ctx)
	if ok {
		return lang
	}
	return def
}
