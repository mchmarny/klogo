# klogo - logo recognision service using Knative and Google Vision API



### REST Service


```shell
curl -H "Content-Type: application/json" \
     -d '{ "id": "test", "url": "https://storage.googleapis.com/kdemo-logos/0.png" }' \
     -X POST https://klogo.demo.knative.tech/
```
