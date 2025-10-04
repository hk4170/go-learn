package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	// 打开文件（可以是很大的文件）
	filename := "test.db"
	file, err := os.Open(filename) // 替换为你的大文件
	if err != nil {
		http.Error(w, "file not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	// 设置响应头
	w.Header().Set("Content-Type", "application/octet-stream")
    end := fmt.Sprintf("attachment; filename=%s",filename)
	w.Header().Set("Content-Disposition", end)

	// 将文件流式地写入 HTTP 响应（流式传输！）
	_, err = io.Copy(w, file)
	if err != nil {
		log.Println("Error streaming file:", err)
	}
}

func main() {
	http.HandleFunc("/", downloadHandler)
	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}