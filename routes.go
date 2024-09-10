package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func SetRoutes(m *http.ServeMux) {
	m.HandleFunc("/get", GetData)
	m.HandleFunc("/post", PostData)
}

func GetData(w http.ResponseWriter, r *http.Request) {

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

func PostData(w http.ResponseWriter, r *http.Request) {

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

	time.Sleep(3 * time.Second)

	w.WriteHeader(201)

	w.Write([]byte("Successfully Posted..\n"))

}
