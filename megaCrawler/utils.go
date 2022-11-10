package megaCrawler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
	"strings"
	"time"
)

type errorResp struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"msg"`
}

func Contain[T comparable](slice []T, check T) bool {
	for _, a := range slice {
		if a == check {
			return true
		}
	}
	return false
}

func mapToKeySlice[T any](m map[string]T) (slice []string) {
	for s := range m {
		slice = append(slice, s)
	}
	return
}

func combineSlice[T any](bottom []T, top []T) []T {
	for _, t := range top {
		bottom = append(bottom, t)
	}
	return bottom
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
	Sugar.Error(msg)
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

func StandardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func downloadFile(filepath string, url string) (err error) {
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func spread(args interface{}) (k []interface{}) {
	s := reflect.ValueOf(args)
	st := s.Type()

	for i := 0; i < st.NumField(); i++ {
		iField := s.Field(i)
		tField := st.Field(i)
		k = append(k, tField.Tag.Get("json"))
		k = append(k, iField.Interface())
	}
	return
}
