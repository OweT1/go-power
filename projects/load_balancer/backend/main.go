package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func handle_request(port string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		resp := fmt.Sprintf("Server on port %s received a request!", port)
		fmt.Print(resp)
		fmt.Fprint(w, resp)
	})

	fmt.Printf("Backend Server running on Port %s...\n", port)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}


func main() {
	port := flag.String("port", "8081", "Server port to run on") // default port value 8081
	flag.Parse()

	handle_request(*port)
}