# klogo - logo recognition service using Cloud Run on GKE and Google Vision API

Simple `go` app demonstrating rest service on Cloud Run

## Deploy

```shell
gcloud beta run deploy klogo \
    --image=gcr.io/knative-samples/klogo:latest \
    --set-env-vars=RELEASE=v0.0.3,GIN_MODE=release
```

## REST Service


```shell
curl -H "Content-Type: application/json" \
     -d '{ "id": "test", "url": "https://storage.googleapis.com/kdemo-logos/k8s.png" }' \
     -X POST https://klogo.demo.knative.tech/ | jq "."
```
