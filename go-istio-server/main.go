package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, Server!")
	})

	http.HandleFunc("/hello-go-client", goHelloHandler)

	fmt.Println("Server is running on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func goHelloHandler(w http.ResponseWriter, r *http.Request) {
	// go-client-server 的 URL
	url := "http://go-istio-client.default.svc.cluster.local/hello-java"

	// 发送 GET 请求到 go-client-server 的接口
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error making GET request to go-client-client: %v", err)
		http.Error(w, "Failed to call go-client-client", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// 读取 go-client-server 的响应
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response from go-client-client: %v", err)
		http.Error(w, "Failed to read response from go-client-client", http.StatusInternalServerError)
		return
	}

	// 将 go-client-server 的响应返回给客户端
	fmt.Fprintf(w, "Response from go-client-client: %s", body, time.Now())
}
