package megaCrawler

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

type errorResp struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"msg"`
}

func contain[T comparable](slice []T, check T) bool {
	for _, a := range slice {
		if a == check {
			return true
		}
	}
	return false
}

func mapToKeySlice[T any](m map[string]T) (slice []string) {
	for s, _ := range m {
		slice = append(slice, s)
	}
	return
}

func successResponse(msg string) (b []byte, err error) {
	errorJson := errorResp{
		StatusCode: 200,
		Message:    msg,
	}
	return json.Marshal(errorJson)
}

func errorResponse(w http.ResponseWriter, statusCode int, msg string) (err error) {
	errorJson := errorResp{
		StatusCode: statusCode,
		Message:    msg,
	}
	b, err := json.Marshal(errorJson)
	if err != nil {
		return err
	}
	w.WriteHeader(statusCode)
	_, err = w.Write(b)
	if err != nil {
		return err
	}
	err = Logger.Error(msg)
	if err != nil {
		return err
	}
	return nil
}

func shortDur(d time.Duration) string {
	s := d.String()
	if strings.HasSuffix(s, "m0s") {
		s = s[:len(s)-2]
	}
	if strings.HasSuffix(s, "h0m") {
		s = s[:len(s)-2]
	}
	return s
}
