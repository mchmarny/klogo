.PHONY: all test query image service


test:
	go test ./...

run:
	bin/build
	./klogo

query:
	curl -H "Content-Type: application/json" \
		-d '{ "id": "test", "url": "https://storage.googleapis.com/kdemo-logos/youtube.jpg" }' \
		-X POST https://logo.demo.knative.tech/

image:
	bin/image

public-image:
	go mod tidy
	go mod vendor
	gcloud builds submit \
		--project cloudylabs-public \
		--tag gcr.io/cloudylabs-public/logo:1.0.1

