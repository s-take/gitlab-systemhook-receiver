# gitlab-systemhook-receiver

## 仕様
* gitlabのsystem-hookを受け取ってJSONをパースする
 * 受け取るJSON情報はこちらhttps://devops.nes.nec.co.jp/gitlab/help/system_hooks/system_hooks
 * ちゃんとgitlabから来るやつはヘッダーにX-Gitlab-Event: System Hookがつく
* パースした結果をもとにコマンド発行とかいろいろやる
 * とりあえず、javaコマンドを実行する
 * エラーは管理者にメールするとか
* portはバッティングしないような適当なポートに8090とかで待ち受ける

## 実装する機能
1. プロジェクトを作成すると対応するジョブをjenkins上に作成する
  * repo-hook,repo-IT-master,repo-IT-develop,repo-UT-develop
1. リポジトリにメンバを追加するとJenkins側のアクセス権も反映する
  * guest→read権限、master,developer→admin権限



## 使い方

```go
package main

import (
  "xxxx/gitlab-systemhook-receiver"
  "log"
  "net/http"
  "time"
)

func main(){
  server := xxxx.
}

```
