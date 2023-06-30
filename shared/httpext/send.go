package httpext

import (
	"IofIPOS/shared/errors"
	"IofIPOS/shared/i18next"
	"IofIPOS/shared/jsonext"
	"context"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
	"regexp"
)

type ResponseError struct {
	Message string   `json:"message"`
	Details []string `json:"details"`
}

func SendError(res http.ResponseWriter, req *http.Request, err error) {
	errModel := errors.Cast(req.Context(), err)
	translatedErr := &errors.Model{
		Code:    uint32(errModel.Code()),
		Message: i18next.ByContext(context.Background(), errModel.Message()),
		Details: errModel.Details()}

	// check if the error has been translated
	if !regexp.MustCompile(`[^\x00-\x7F]`).MatchString(translatedErr.Message) {
		log.WithError(errors.Cast(req.Context(), err)).Error("can not translate error message")
		translatedErr.Message = i18next.ByContext(context.Background(), "internal_server_error")
	}

	SendModel(res, req, errModel.HttpStatus(), translatedErr)
}

func SendModel(res http.ResponseWriter, req *http.Request, code int, model interface{}) {
	bytes, err := json.Marshal(model)
	if err != nil {
		log.WithError(errors.Cast(req.Context(), err)).Error("can not marshal model")
	}
	SendData(res, req, code, jsonext.MIME, bytes)
}

func SendData(res http.ResponseWriter, req *http.Request, code int, mime string, data []byte) {
	res.Header().Set(ContentTypeHeader, mime)
	res.Header().Set(CharsetHeader, "utf-8")
	res.WriteHeader(code)
	_, err := res.Write(data)
	if err != nil {
		log.WithError(errors.Cast(req.Context(), err)).Error("can not write data on response")
	}
}

func SendCode(res http.ResponseWriter, req *http.Request, code int) {
	res.WriteHeader(code)
}
