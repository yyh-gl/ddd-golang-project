package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	// ルーティング設定
	r := httprouter.New()

	fmt.Println("========================")
	fmt.Println("サーバ起動 >> http://localhost:3000")
	fmt.Println("========================")
	if err := http.ListenAndServe(":3000", r); err != nil {
		panic("サーバ起動エラー")
	}
}
