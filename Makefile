.PHONY: all test secret image query deps run service-new


test:
	go test ./...

query:
	curl -H "Content-Type: application/json" \
		 -d '{ "id": "test", "url": "https://storage.googleapis.com/kdemo-logos/0.png" }' \
		 -X POST http://localhost:8080/

image:
	gcloud builds submit \
		--project ${GCP_PROJECT} \
		--tag gcr.io/${GCP_PROJECT}/klogo:latest

service:
	kubectl apply -f service.yaml

