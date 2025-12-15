package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"sync/atomic"

	"load_balancer/constants"

	"github.com/joho/godotenv"
)

type Backend struct {
	url *url.URL
	ReverseProxy *httputil.ReverseProxy
}

type ServerPool struct {
	backends []*Backend
	current uint64
}

func (s *ServerPool) addBackend(serverUrl string) {
	u, _ := url.Parse(serverUrl)
	proxy := httputil.NewSingleHostReverseProxy(u)

	s.backends = append(s.backends, &Backend{
		url: u,
		ReverseProxy: proxy,
	})
}

func (s *ServerPool) getNextAvailable() *Backend {
	next := atomic.AddUint64(&s.current, 1)
	index := int(next % uint64(len(s.backends)))

	return s.backends[index]
}

func lbHandler(w http.ResponseWriter, r *http.Request) {
	availableBE := serverPool.getNextAvailable()

	fmt.Printf("Redirecting request to %s\n", availableBE.url)

	availableBE.ReverseProxy.ServeHTTP(w, r)
}

var serverPool ServerPool = ServerPool{
	backends: []*Backend{},
	current: 0,
}

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Println("No .env file found, using defaults")
	}

	for _, p := range constants.SERVER_PORTS {
		serverPool.addBackend(fmt.Sprintf("%s:%s", constants.BASE_URL, p))
	}

	LB_PORT := os.Getenv("LB_PORT")

	mainServer := http.Server{
		Addr: fmt.Sprintf(":%s", LB_PORT),
		Handler: http.HandlerFunc(lbHandler),
	}

	fmt.Printf("Load Balancer started on Port %s\n", LB_PORT)
	log.Fatal(mainServer.ListenAndServe())
}