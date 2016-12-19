package main

import (
	"flag"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

var (
	listenAddr = flag.String("listen", "localhost:8000", "HTTP listen address")
)

func main() {
	flag.Parse()

	s := new(spaceship)
	go s.liftoff()
	http.Handle("/", s)
	log.Fatal(http.ListenAndServe(*listenAddr, nil))
}

type spaceship struct {
	datapoints []string
	maxpoints  int
}

func (s *spaceship) liftoff() {
	dump, err := ioutil.ReadFile("dump.dat")
	if err != nil {
		log.Fatal(err)
	}

	s.datapoints = strings.Split(string(dump), "\n")
	s.maxpoints = len(s.datapoints)
	log.Println("Lift-off!")
}

func (s *spaceship) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	randNumber := rand.Intn(100)
	if randNumber == 5 {
		log.Println("slow connection")
		time.Sleep(10 * time.Second)
	}
	if randNumber == 2 {
		log.Println("404")
		http.NotFound(w, r)
		return
	}

	randDataPoint := rand.Intn(s.maxpoints - 1)
	_, err := io.WriteString(w, s.datapoints[randDataPoint])
	if err != nil {
		log.Fatal(err)
	}
	log.Println("datapoint: ", s.datapoints[randDataPoint])
}
