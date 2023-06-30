package httpext

import (
	"IofIPOS/shared/contextext"
	"net/http"
	"strings"
)

func AuthBearer(req *http.Request) string {
	return strings.TrimSpace(strings.Replace(req.Header.Get("Authorization"), "Bearer ", "", 1))
}

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if token := AuthBearer(req); token != "" {
			req = req.WithContext(contextext.SetToken(req.Context(), token))
		}
		next.ServeHTTP(res, req)
	})
}
