package main

import (
	"bytes"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var (
	listenAddr = flag.String("listen", "localhost:8001", "HTTP listen address")
)

func main() {
	flag.Parse()

	g := new(Groundcontrol)
	go g.run()

	http.Handle("/tmp/data.png", g)
	http.Handle("/", http.FileServer(http.Dir("./web")))

	log.Fatal(http.ListenAndServe(*listenAddr, nil))
}

type Groundcontrol struct {
	datapoints []string
}

func (g *Groundcontrol) run() {

	// --------------------
	// START HACKING HERE
	// --------------------

	// 1. Create a new channel where we receive the datapoint information (string) using make
	// 2. Fire up 50 workers to connect to the spaceship (using 50 separate go routines).
	// 3. Use a endless for loop to append datapoint from each channel, to the groundcontrol datapoints (g.datapoints).
	// 4. Execute: go run groundcontrol and see wat happen on localhost:8001
}

// Worker continues try to reach the spaceship
func worker(id int, data chan string) {
	for {
		message, err := reachout()
		if err != nil {
			log.Printf("Error reading spaceship response: %v (worker %v)", err, id)
			continue
		}
		data <- message
	}
}

// reachout to the spaceship.
func reachout() (string, error) {

	// 1. Use the http.Get to access the spaceship
	resp, err := http.Get("http://localhost:8000/")
	if err != nil {
		return "", err
	}

	// 2. Don't forget closing the response body using 'defer'
	defer resp.Body.Close()

	// 3. When the response statuscode isn't a 200 return an error using fmt.Errof()
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("Non-200 status code %d", resp.StatusCode)
	}

	// 4. Next step is to read the body from the response using ioutil.ReadAll()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// 5. Cast the body to a string and return this function
	return string(body), nil
	// 6. That's it, use go run grouncontrol to test this.
}

func (g Groundcontrol) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m := image.NewRGBA(image.Rect(0, 0, 1261, 732))
	for _, datapoint := range g.datapoints {
		data := strings.Split(datapoint, ";")
		col, err := convertDataToColor(data)
		if err != nil {
			log.Printf("error convert datapoint %v", data)
			continue
		}

		x, err := strconv.Atoi(data[0])
		if err != nil {
			log.Printf("Unsupported x value %v", data[0])
			continue
		}

		y, err := strconv.Atoi(data[1])
		if err != nil {
			log.Printf("Unsupported y value %v", data[1])
			continue
		}

		m.Set(y, x, col)
	}

	var img image.Image = m
	writeImage(w, &img)
}

func convertDataToColor(data []string) (color.RGBA, error) {

	var c color.RGBA

	i, err := strconv.Atoi(data[2])
	if err != nil {
		return c, fmt.Errorf("Unsupported red value %v", data[2])
	}
	r := uint8(i >> 8)

	i, err = strconv.Atoi(data[3])
	if err != nil {
		return c, fmt.Errorf("Unsupported green value %v", data[3])
	}
	g := uint8(i >> 8)

	i, err = strconv.Atoi(data[4])
	if err != nil {
		return c, fmt.Errorf("Unsupported blue value %v", data[4])
	}
	b := uint8(i >> 8)

	i, err = strconv.Atoi(data[5])
	if err != nil {
		return c, fmt.Errorf("Unsupported alpha value %v", data[5])
	}
	a := uint8(i >> 8)

	c = color.RGBA{r, g, b, a}

	return c, nil
}

// writeImage encodes an image 'img' in jpeg format and writes it into ResponseWriter.
func writeImage(w http.ResponseWriter, img *image.Image) {

	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, *img, nil); err != nil {
		log.Println("unable to encode image.")
	}

	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	if _, err := w.Write(buffer.Bytes()); err != nil {
		log.Println("unable to write image.")
	}
}
