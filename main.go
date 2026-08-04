package main

import (
	"crypto/rand"
	"embed"
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
)

//go:embed public
var public embed.FS

func main() {
	pub, err := fs.Sub(public, "public")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	http.Handle("/", http.FileServerFS(pub))

	http.HandleFunc("/download", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Cache-Control", "no-store")

		chunk := make([]byte, 256*1024)
		if _, err := rand.Read(chunk); err != nil {
			http.Error(w, "Failed to generate data", http.StatusInternalServerError)
			return
		}

		ctx := r.Context()

		for {
			select {
			case <-ctx.Done():
				return
			default:
				_, err := w.Write(chunk)
				if err != nil {
					return
				}
				if flusher, ok := w.(http.Flusher); ok {
					flusher.Flush()
				}
			}
		}
	})

	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		bytesReceived, err := io.Copy(io.Discard, r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"ok":       true,
			"received": bytesReceived,
		})
	})

	port := 8083
	fmt.Printf("LAN Speedtest running on http://0.0.0.0:%d\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
