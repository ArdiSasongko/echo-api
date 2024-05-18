package main

func main() {
	server := InitializeServer()
	err := server.Start(":8080")
	if err != nil {
		return
	}
}
