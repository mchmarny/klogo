package main

import (
	"context"
	"log"

	vision "cloud.google.com/go/vision/apiv1"
)

// Process processes passed URL and returns map of attributes
func getLogoFromUrl(url string) (desc string, err error) {

	ctx := context.Background()
	client, e := vision.NewImageAnnotatorClient(ctx)
	if e != nil {
		return "", e
	}

	image := vision.NewImageFromURI(url)
	annotations, e := client.DetectLogos(ctx, image, nil, 10)
	if e != nil {
		return "", e
	}

	if len(annotations) == 0 {
		log.Printf("No logo desc found for: %s", url)
		return "", nil
	}

	log.Println("Found logos:")
	return annotations[0].Description, nil
}