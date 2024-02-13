package main

import (
	"fmt" // 標準出力を行うためのパッケージ
	"html/template" // HTMLファイルをパースするためのパッケージ
	"log" // エラーログを出力するためのパッケージ
	"net/http" // HTTPリクエストを処理するためのパッケージ
)

// HTTPリクエストを処理
func viewHandler(w http.ResponseWriter, r *http.Request) {
	// view.htmlをパース
	html, err := template.ParseFiles("view.html")
	// エラーが発生した場合はエラーログを出力
	if err != nil {
		log.Fatal(err)
	}
	// view.htmlをレンダリング
	if err := html.Execute(w, nil); err != nil {
		log.Fatal(err)
	}
}

// サーバーを起動
func main() {
	// localhost:8080/viewにアクセスした際にviewHandlerを呼び出す
	http.HandleFunc("/view", viewHandler)
	// サーバーが起動したことをログに出力
	fmt.Println("Server is running at http://localhost:8080/")
	// localhost:8080でサーバーを起動
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}