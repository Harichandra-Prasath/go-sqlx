package main

type User struct {
	Firstname  string `json:"first_name" db:"first_name"`
	Lastname   string `json:"last_name" db:"last_name"`
	Email      string `json:"email" db:"email"`
	Age        int    `json:"age" db:"age"`
	Experience int    `json:"experience" db:"experience"`
}

type Job struct {
	Role string `json:"role"`
	Pay  int    `json:"pay"`
}
