package localdebugserver

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
)

var (
	invalidURLParamCharsRE = regexp.MustCompile(`["'<>]`)
)

func respondWithJSON(req *http.Request, w http.ResponseWriter, obj interface{}) {
	j, err := json.Marshal(obj)
	if err != nil {
		log.Printf("error: %v", err)
		return
	}
	responseWithJSONBytes(req, w, j)
}

func responseWithJSONBytes(req *http.Request, w http.ResponseWriter, j []byte) {
	w.Header().Set("Content-Type", "text/json; charset=utf-8")
	fmt.Fprint(w, string(j))
	fmt.Fprint(w, "\n")
	if debug := getBoolURLParam(req, "debug"); debug {
		log.Printf("respondWithJSON: %s", string(j))
	}
}

func getBoolURLParam(req *http.Request, key string) bool {
	key = sanitizeURLParam(key)
	vals := req.URL.Query()[key]
	if len(vals) > 0 {
		v := vals[0]
		if v == "0" || strings.ToLower(v) == "false" {
			return false
		}
		return true
	}
	return false
}

func sanitizeURLParam(s string) string {
	return invalidURLParamCharsRE.ReplaceAllString(s, "")
}

type errorResponse struct {
	Error string
}

func respondWithError(w http.ResponseWriter, req *http.Request, err error) {
	respondWithJSON(req, w, errorResponse{err.Error()})
}
