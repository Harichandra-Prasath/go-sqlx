package main

type User struct {
	Firstname  string `json:"first_name"`
	Lastname   string `json:"last_name"`
	Email      string `json:"email"`
	Age        int    `json:"age"`
	Experience int    `json:"experience"`
}

type Job struct {
	Role string `json:"role"`
	Pay  int    `json:"pay"`
}
