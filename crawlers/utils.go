package crawlers

import (
	"encoding/json"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

type errorResp struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"msg"`
}

func Unique[T comparable](intSlice []T) []T {
	keys := make(map[T]bool)
	var list []T
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func Contain[T comparable](slice []T, check T) bool {
	for _, a := range slice {
		if a == check {
			return true
		}
	}
	return false
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

func StandardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func GetNextIndexURL(currentUrl string, currentPageNum string, paramName string) string {
	thisUrl, _ := url.Parse(currentUrl)
	paramList := thisUrl.Query()

	currentNum, _ := strconv.Atoi(currentPageNum)
	currentNum++

	paramList.Set(paramName, strconv.Itoa(currentNum))
	thisUrl.RawQuery = paramList.Encode()

	return thisUrl.String()
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

func SplitDelimiters(s string, delimiters []string) (result []string) {
	for _, delimiter := range delimiters {
		s = strings.ReplaceAll(s, delimiter, ",")
	}
	return strings.Split(s, ",")
}
