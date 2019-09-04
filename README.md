# jobnetes
Kubernetes上で稼働するワークフローエンジンを作ってみた。趣味プロジェクトです。

## ワークフローとは
ある特定のバッチ処理などの単一のジョブを連鎖させたものがワークフロー。
airflowやdigdagで扱うものと同じ。

### 他との違い
ジョブはすべてkubernetesのJobリソースとして起動されるようになっている。
すべてをkubernetes上で完結させることで、リソース管理や実行管理を統一的に管理できる。

（Argo workflowがこれの最終進化版なイメージ）

# overview
![architecture](https://github.com/Attsun1031/jobnetes/raw/master/docs/images/architecture.jpg "architecture")

## 各種コンポーネント説明
### webadmin
ワークフローの実行状況の可視化等を行うためのウェブアプリケーション。

### manager
ワークフロー管理を行うアプリケーション。
一定の間隔でワークフロー状況をポーリングし、各ジョブ・ワークフローの起動やステータス管理を行う。

### jobapi
ワークフローの起動リクエストを受け付けたり、各ジョブが後続ジョブのために結果を書き込むために利用する。
外部アプリケーションからの利用を想定したため、gRPCでやりとりを行うようにすることでクライアントライブラリ生成を簡素化している。

## ワークフロー定義
jsonでワークフローを定義する。
スキーマや例は以下参照。

https://github.com/Attsun1031/jobnetes/blob/master/schema/workflow-schema.json
https://github.com/Attsun1031/jobnetes/blob/master/schema/test_schema.py

## その他
RDBで管理しているワークフロー情報をカスタムリソースとして登録し、managerやjobapiをcontrollerとして実装すれば、よりKubertenesの恩恵が受けられそう。

# Develop
## Setup dev env
1. set GOROOT
1. set GOPATH
1. clone this repository in $GOPATH/src/github.com/Attsun1031/jobnetes
1. `go get -u github.com/golang/dep/cmd/dep`
1. cd to jobnetes dir
1. `dep ensure`
1. add config.yaml and kube-config to $HOME/.jobnetes
1. start local mysql container
  `docker run --name jobnetes-db -p3333:3306 -e MYSQL_ALLOW_EMPTY_PASSWORD=yes -v ~/jobnetes-db:/var/lib/mysql -d mysql --character-set-server=utf8 --collation-server=utf8_unicode_ci`
1. execute cmd/dbmigration/dbmigration.go
  
### Setup local k8s env
1. start local kubernetes
1. apply `setting/k8s/deploy-mysql.yaml`
1. apply `setting/k8s/cm-config.yaml`
1. apply `setting/k8s/job-migration.yaml`
