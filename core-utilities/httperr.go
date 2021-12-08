package core_utilities

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// HandleError checks if the error is not nil, writes it to the output
// with the specified status code, and returns true. If error is nil it returns false.
func HandleError(w http.ResponseWriter, err error, statusCode int) bool {
	if err == nil {
		return false
	}
	http.Error(w, string(err.Error()), statusCode)
	return true
}

func JSONResponse(w http.ResponseWriter, result interface{}) {
	body, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusOK)
	w.Write(prettyJSON(body))
}

func JSONResponseCode(w http.ResponseWriter, result interface{}, responseCode int) {
	body, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(responseCode)
	w.Write(prettyJSON(body))
}

func ErrorResponse(w http.ResponseWriter, error string, code int) {
	data := struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}{
		Code:    code,
		Message: error,
	}

	body, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusOK)
	w.Write(prettyJSON(body))
}

func prettyJSON(b []byte) []byte {
	var out bytes.Buffer
	json.Indent(&out, b, "", "  ")
	return out.Bytes()
}
