package main

import (
	"fmt"
	"net/http"

	"go.x2ox.com/hitokoto/api/hitokoto"
)

func main() {
	http.HandleFunc("/api/hitokoto", hitokoto.Handler)

	if err := http.ListenAndServe(":8088", nil); err != nil {
		fmt.Println("http listen failed:", err)
	}
}
