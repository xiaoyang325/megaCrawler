package megaCrawler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"megaCrawler/commandImpl"
	"megaCrawler/config"
	"net/http"
)

//func Template(w http.ResponseWriter, r *http.Request) {
//	L.Info("Receive"+r.Method + "websiteList Request from: " + r.RemoteAddr)
//	var b []byte
//	var err error
//	switch r.Method {
//	case "GET":
//		var k []any
//		b, err := json.Marshal(k)
//	default:
//		_ = errorResponse(w, 405, "Method not allowed")
//		return
//	}
//	if err != nil {
//		_ = L.Error("Failed to serialize response:" + err.Error())
//		_ = errorResponse(w, 500, "Failed to serialize response:"+err.Error())
//		return
//	}
//	w.Header().Add("Content-Type", "application/json")
//	_, _ = w.Write(b)
//}

func startHandler(w http.ResponseWriter, r *http.Request) {
	_ = L.Info("Receive" + r.Method + "startHandler Request from: " + r.RemoteAddr)
	var b []byte
	var err error
	switch r.Method {
	case "GET":
		id, ok := mux.Vars(r)["id"]
		if !ok {
			err = errorResponse(w, 400, "Bad Request : Invalid argument, missing id")
			return
		}
		if website, ok := WebMap[id]; ok {
			if website.IsRunning {
				err = errorResponse(w, 400, "Bad Request : Crawler is already running")
				return
			}
			website.Scheduler.RunAll()
			b, err = successResponse("Crawler should start soon")
		} else {
			err = errorResponse(w, 400, "Bad Request : Invalid argument, id does not exist")
			return
		}
	default:
		_ = errorResponse(w, 405, "Method not allowed")
		return
	}
	if err != nil {
		_ = L.Error("Failed to serialize response:" + err.Error())
		_ = errorResponse(w, 500, "Failed to serialize response:"+err.Error())
		return
	}
	w.Header().Add("Content-Type", "application/json")
	_, _ = w.Write(b)
}

func websiteHandler(w http.ResponseWriter, r *http.Request) {
	_ = L.Info("Receive" + r.Method + "websiteHandler Request from: " + r.RemoteAddr)
	var b []byte
	var err error
	switch r.Method {
	case "GET":
		id, ok := mux.Vars(r)["id"]
		if !ok {
			err = errorResponse(w, 400, "Bad Request : Invalid argument, missing id")
			return
		}
		if website, ok := WebMap[id]; ok {
			b, err = website.ToJson()
		} else {
			err = errorResponse(w, 400, "Bad Request : Invalid argument, id does not exist")
			return
		}
	case "POST":
		_ = errorResponse(w, 405, "Method not allowed")
		return
	default:
		_ = errorResponse(w, 405, "Method not allowed")
		return
	}
	w.Header().Add("Content-Type", "application/json")
	_, _ = w.Write(b)
	if err != nil {
		_ = L.Error("Failed to serialize response:" + err.Error())
		_ = errorResponse(w, 500, "Internal Error : Failed to serialize response:"+err.Error())
		return
	}
	w.Header().Add("Content-Type", "application/json")
	_, _ = w.Write(b)
}

//websiteListHandler returns all registered websites
func websiteListHandler(w http.ResponseWriter, r *http.Request) {
	L.Info("Receive" + r.Method + "websiteList Request from: " + r.RemoteAddr)
	var b []byte
	var err error
	switch r.Method {
	case "GET":
		var k []commandImpl.WebsiteStatus
		for _, engine := range WebMap {
			k = append(k, engine.ToStatus())
		}
		b, err = json.Marshal(k)
	default:
		_ = errorResponse(w, 405, "Method not allowed")
		return
	}
	if err != nil {
		_ = L.Error("Failed to serialize response:" + err.Error())
		_ = errorResponse(w, 500, "Internal Error : Failed to serialize response:"+err.Error())
		return
	}
	w.Header().Add("Content-Type", "application/json")
	_, _ = w.Write(b)
}

//StartWebServer starts the internal webserver to communicate between service and ctl tool.
//This is not intended for external use, please close port 7171 if you don't know what you are doing.
func StartWebServer() {
	r := mux.NewRouter()
	// example usage: curl -s 'http://127.0.0.1:7171/websites/'
	r.HandleFunc("/websites/", websiteListHandler)
	r.HandleFunc("/website/{id}/", websiteHandler)
	r.HandleFunc("/website/{id}/start/", startHandler)

	http.Handle("/", r)
	_ = L.Info("Listening on", config.Port)
	go func() {
		http.ListenAndServe(config.Port, nil)
	}()
}
