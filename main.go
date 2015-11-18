package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func systemhookReceive(w http.ResponseWriter, r *http.Request) {

	// headerチェックもした方がいいかも
	// POST以外はリターン
	if r.Method != "POST" {
		fmt.Printf("Not POST method\n")
		return
	}

	// JSONパース
	decoder := json.NewDecoder(r.Body)
	var j interface{}
	err := decoder.Decode(&j)
	if err != nil {
		panic(err)
	}

	// JSON出力部分
	fmt.Printf("%v\n", j)
	m := j.(map[string]interface{})
	fmt.Printf("%s\n", m["event_name"])
	fmt.Printf("%s\n", m["name"])
	fmt.Printf("%s\n", m["owner_name"])
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
