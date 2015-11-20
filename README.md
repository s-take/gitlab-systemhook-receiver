# gitlab-systemhook-receiver

gitlabのリポジトリ作成とJenkins側のジョブ作成およびアクセス制御をシームレスにするためのツール  
GitlabとJenkinsには同じアカウントが存在している前提

## 仕様
* gitlabのsystem-hookを受け取ってJSONをパースする
 * 受け取るJSON情報は以下のような感じ
```json
{
          "created_at": "2012-07-21T07:30:54Z",
          "event_name": "project_create",
                "name": "StoreCloud",
         "owner_email": "johnsmith@gmail.com",
          "owner_name": "John Smith",
                "path": "storecloud",
 "path_with_namespace": "jsmith/storecloud",
          "project_id": 74,
  "project_visibility": "private",
}
```
 * gitlabから来るやつはヘッダーにX-Gitlab-Event: System Hookがつく
* パースした結果をもとにJenkins側でごにょごにょやる
 * ジョブ作成、アクセス制御
 * Jenkins側はCLIとWebAPIがあるみたい。とりあえず、javaコマンドを実行する?WebAPIの仕様次第
 * WebAPIってbasic認証かな？簡単かつCLIとやれること一緒ならWebAPIで実装する
 * エラー通知はログ出力＋管理者にメール程度で
* portは開発系のOSSとバッティングしないような適当なポート(8090とか)で待ち受ける

## 機能
1. Gitlabにリポジトリを作成するとJenkins上に対応するジョブを作成する
  * reponame-push-hook ・・・プッシュのフックを受け取って実行させるやつ
  * reponame-master-UT,reponame-develop-UT ・・・↑のジョブから実行される子ジョブ
  * reponame-master-IT ・・・結合テスト用のジョブ、環境依存なのでGit以外は空のジョブ

1. Gitlabのリポジトリにメンバを追加するとJenkins側のアクセス権も反映する
  * guest→read権限、master,developer→admin権限
  * reponame始まりのジョブに対して設定する

## 使い方イメージ

とりあえず動かす方法  
```bash
$ gitlab-systemhook-receiver -p 8090
```

## TODO
* デーモンで起動させる。(systemd前提でgitlab-systemhook-receiver.serviceでもいいかな)
