package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func badRequest(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(w, "Bad Request")
	log.Print(message)
}

func systemhookReceive(w http.ResponseWriter, r *http.Request) {

	// Gitlabのシステムフック以外は400でリターン
	if r.Header.Get("X-Gitlab-Event") != "System Hook" {
		badRequest(w, "Not exsit X-Gitlab-Event header")
		return
	}

	// POST以外は400でリターン
	if r.Method != "POST" {
		badRequest(w, "Not POST Method")
		return
	}

	// JSONパース,デコードエラー時は400でリターン
	decoder := json.NewDecoder(r.Body)
	var j interface{}
	err := decoder.Decode(&j)
	if err != nil {
		badRequest(w, "decord error")
		return
	}

	// デバッグ用JSON出力部分
	log.Printf("%v\n", j)
	m := j.(map[string]interface{})
	log.Printf("%s\n", m["event_name"])
	log.Printf("%s\n", m["name"])
	log.Printf("%s\n", m["owner_name"])
}

func main() {
	// ルーティング設定
	http.HandleFunc("/", systemhookReceive)
	// httpサーバ起動
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
