package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Res struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-type", "application/json")
		res := Res{
			0, "http",
		}

		jRes, err := json.Marshal(res)
		if err != nil {
			log.Println(err)
		}

		_, err = writer.Write(jRes)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(err)
	})

	httpErr := http.ListenAndServe(":20080", mux)

	if httpErr != nil {
		log.Fatalln(httpErr)
	}
}
