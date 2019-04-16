.PHONY: all test query image service


test:
	go test ./...

query:
	curl -H "Content-Type: application/json" \
		-d '{ "id": "test", "url": "https://storage.googleapis.com/kdemo-logos/0.png" }' \
		-X POST https://klogo.demo.knative.tech/

image:
	go mod tidy
	go mod vendor
	gcloud builds submit \
		--project ${GCP_PROJECT} \
		--tag gcr.io/${GCP_PROJECT}/klogo:latest

image-sample:
	go mod tidy
	go mod vendor
	gcloud builds submit \
		--project knative-samples \
		--tag gcr.io/knative-samples/klogo:latest

service:
	kubectl apply -f service.yaml

