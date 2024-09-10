package main

func main() {
	cfg := &Config{
		Addr: ":3000",
	}
	server := GetServer(cfg)
	server.Start()

}
