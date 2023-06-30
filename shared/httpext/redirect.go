package httpext

import (
	"net/http"
	"net/url"
)

func Redirect(res http.ResponseWriter, req *http.Request, uri *url.URL) {
	http.Redirect(res, req, uri.String(), http.StatusPermanentRedirect)
}
