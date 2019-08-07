# Telepresence Demo

## バージョン情報

* telepresence: 0.101
* macOS: 10.14.5
* Docker Desktop: 2.1.0.0
* kubectl: v1.15.2
* Kubernetes: GKE, v1.13.7-gke.8
* go: 1.12.7

## 事前準備

### Telepresence CLIのインストール

[Telepresence公式](https://www.telepresence.io/reference/install)の手順に沿ってインストールする

### Kubernetes リソースのデプロイ

```sh
kubectl apply -f ./k8s-manifest/telepresence-demo.yaml

# 動作確認
CURL_POD=$(kubectl get po -l app=curl -ojsonpath='{.items[0].metadata.name}')
kubectl exec $CURL_POD -- curl -v devapp:8080
# nginx のデフォルトページが出ればOK
```

## Dockerイメージの差し替え

```sh
# go-demo/facade.go の内容を変更
# 注: 本来は接続先は環境変数に設定して再ビルドを避ける方が良いが、デモ用のサンプルコードなので……
sed -i -e "s/backend-1/backend-2/" ./go-demo/facade.go

# Dockerイメージをビルド
docker build -t go-demo:v2 ./go-demo/

# 差し替え
telepresence --swap-deployment devapp --docker-run go-demo:v2 /demoapp
# 下記のログが表示されれば準備OK
# T: Setup complete. Launching your container.

# 動作確認
CURL_POD=$(kubectl get po -l app=curl -ojsonpath='{.items[0].metadata.name}')
kubectl exec $CURL_POD -- curl -v devapp:8080
# Apache httpdのデフォルトページ (It works!) が出ればOK
```
