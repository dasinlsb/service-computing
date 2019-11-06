package  main

import (
	"fmt"
	"net/http"
)

func SayHello(w http.ResponseWriter, req *http.Request) {
	_, _ = w.Write([]byte("hello"))
}

func main() {
	http.HandleFunc("/", SayHello)
	fmt.Printf("Listening on port %v...\n", 8080)
	_ = http.ListenAndServe(":8080", nil)
}