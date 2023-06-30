package httpext

import (
	"IofIPOS/shared/errors"
	"bytes"
	"encoding/json"
	"google.golang.org/grpc/codes"
	"net/http"
)

func BindModel(req *http.Request, model interface{}) (err error) {
	defer func() {
		_ = req.Body.Close()
		if err != nil {
			err = errors.New(req.Context(), codes.InvalidArgument).
				AddDetails(err.Error())
		}
	}()
	var obj map[string]interface{}
	if err = json.NewDecoder(req.Body).Decode(&obj); err != nil {
		return err
	}
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false)
	if err = encoder.Encode(obj); err != nil {
		return err
	}
	return json.NewDecoder(&buf).Decode(model)
}
