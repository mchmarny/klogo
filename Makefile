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
		--tag gcr.io/${GCP_PROJECT}/klogo

public-image:
	go mod tidy
	go mod vendor
	gcloud builds submit \
		--project cloudylabs-public \
		--tag gcr.io/cloudylabs-public/logo:1.0.1

service:
	kubectl apply -f service.yaml

