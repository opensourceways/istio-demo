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
		fmt.Fprintln(w, "Hello, Client!")
	})

	http.HandleFunc("/hello-go-server", goHelloHandler)

	http.HandleFunc("/hello-java", javaHelloHandler)

	fmt.Println("Server is running on port 8081...")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func javaHelloHandler(w http.ResponseWriter, r *http.Request) {
	// go-client-server 的 URL
	url := "http://java-istio-server.default.svc.cluster.local/hello-go-server"

	// 发送 GET 请求到 go-client-server 的接口
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error making GET request to java-client-server: %v", err)
		http.Error(w, "Failed to call java-client-server", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// 读取 go-client-server 的响应
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response from java-client-server: %v", err)
		http.Error(w, "Failed to read response from java-client-server", http.StatusInternalServerError)
		return
	}

	// 将 go-client-server 的响应返回给客户端
	fmt.Fprintf(w, "Response from java-client-server: %s", body, time.Now())
}

func goHelloHandler(w http.ResponseWriter, r *http.Request) {
	// go-client-server 的 URL
	url := "http://go-istio-server.default.svc.cluster.local/hello"

	// 发送 GET 请求到 go-client-server 的接口
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error making GET request to go-client-server: %v", err)
		http.Error(w, "Failed to call go-client-server", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// 读取 go-client-server 的响应
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response from go-client-server: %v", err)
		http.Error(w, "Failed to read response from go-client-server", http.StatusInternalServerError)
		return
	}

	// 将 go-client-server 的响应返回给客户端
	fmt.Fprintf(w, "Response from go-client-server: %s", body, time.Now())
}
