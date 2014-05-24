package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	port := getPort()
	http.HandleFunc("/", proxyHandler)
	http.ListenAndServe(port, nil)
}

func getPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "8181"
		fmt.Println("INFO: defaulting to " + port)
		fmt.Println("set a port as an ENV var if you wish")
	}

	return ":" + port
}

func proxyHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	content := imgReader(params["img"][0])
	w.Write(content)
}

func imgReader(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		panic("image URL not read")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("the response could not be read into memory")
	}

	return body
}
