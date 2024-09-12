package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetData(S *Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodGet {
			w.Write([]byte(fmt.Sprintf("%s Method not allowed\n", r.Method)))
			return
		}

		var res []User
		S.DB.Select(&res, "SELECT * FROM users")

		data, err := json.Marshal(res)
		if err != nil {
			fmt.Println("Error in Marshaling:", err)
			w.WriteHeader(500)
			return
		}
		w.Write(data)

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
			return
		}

		_, err = S.DB.NamedExec("INSERT INTO users (first_name, last_name, email, age, experience) VALUES (:first_name, :last_name, :email, :age, :experience)", &user)
		if err != nil {
			fmt.Println("Error in writing into db:", err)
			w.WriteHeader(500)
			return
		}
		w.WriteHeader(201)

		w.Write([]byte("Successfully Posted..\n"))

	}
}
