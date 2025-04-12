package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		str := os.Getenv("HELLO_STR")
		if str == "" {
			str = "Welcome to the custom app"
		}
		fmt.Fprintf(w, str)
	})
	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})
	http.HandleFunc("/log", func(w http.ResponseWriter, r *http.Request) {
		type RequestBody struct {
			Message string `json:"message"`
		}
		bodyBytes, _ := io.ReadAll(r.Body)
		var requestBody RequestBody
		_ = json.Unmarshal(bodyBytes, &requestBody)
		_ = os.Mkdir("logs", 0644)
		file, _ := os.OpenFile("/app/logs/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		defer file.Close()
		_, _ = fmt.Fprintln(file, requestBody.Message)
	})
	http.HandleFunc("/logs", func(w http.ResponseWriter, r *http.Request) {
		_ = os.Mkdir("logs", 0644)
		file, _ := os.OpenFile("/app/logs/app.log", os.O_APPEND|os.O_CREATE|os.O_RDONLY, 0644)
		defer file.Close()
		bytes_arr, _ := io.ReadAll(file)
		fmt.Fprintln(w, string(bytes_arr))
	})
	http.ListenAndServe(":80", nil)
}
