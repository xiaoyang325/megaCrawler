// Package commands handles client side command query to the local server
package commands

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

func getJson(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	if r.StatusCode != http.StatusOK {
		bodyB, _ := ioutil.ReadAll(r.Body)
		bodyStr := string(bytes.Replace(bodyB, []byte("\r"), []byte("\r\n"), -1))
		return errors.New(bodyStr)
	}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&target)
	return err
}

type WebsiteStatus struct {
	Id          string    `json:"id"`
	BaseUrl     string    `json:"baseUrl"`
	IsRunning   bool      `json:"isRunning"`
	NextIter    time.Time `json:"nextIter"`
	ProgressBar string    `json:"progressBar"`
	Bar         string    `json:"bar"`
	Name        string    `json:"name"`
	IterPerSec  float64   `json:"iterPerSec"`
}

type ErrorResp struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"msg"`
}

func GetWebsite(id string) (s WebsiteStatus, err error) {
	err = getJson("http://127.0.0.1:7171/website/"+id+"/", &s)
	return
}

func GetWebsites() (websites []WebsiteStatus, err error) {
	err = getJson("http://127.0.0.1:7171/websites/", &websites)
	return
}

func StartWebsite(id string) (s ErrorResp, err error) {
	err = getJson("http://127.0.0.1:7171/website/"+id+"/start/", &s)
	return
}
