package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (S *Server) SetRoutes() {
	S.Mux.HandleFunc("/get", GetData(S))
	S.Mux.HandleFunc("/post", PostData(S))
}

func GetData(S *Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodGet {
			w.Write([]byte(fmt.Sprintf("%s Method not allowed\n", r.Method)))
			return
		}

		var res string

		for i := 0; i < 5; i++ {
			res += fmt.Sprintf("%d Got it\n", i)
		}

		w.Write([]byte(res))

	}
}

func PostData(S *Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		user := User{}

		if r.Method != http.MethodPost {
			w.Write([]byte(fmt.Sprintf("%s Method not allowed\n", r.Method)))
			return
		}

		data, err := io.ReadAll(r.Body)

		if err != nil {
			fmt.Println("Error in Reading the data:", err)
			return
		}

		fmt.Println("Recieved Data:", string(data))

		err = json.Unmarshal(data, &user)
		if err != nil {
			fmt.Println("Error in umarshaling", err)
			w.WriteHeader(400)
			w.Write([]byte("Data failed"))
		}

		w.WriteHeader(201)

		w.Write([]byte("Successfully Posted..\n"))

	}
}
