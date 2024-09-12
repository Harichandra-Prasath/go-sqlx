package main

func (S *Server) SetRoutes() {
	S.Mux.HandleFunc("/get", GetData(S))
	S.Mux.HandleFunc("/post", PostData(S))
}
