package main

import (
	"context"

	vision "cloud.google.com/go/vision/apiv1"
)

// getLogoFromURL processes passed URL and returns map of attributes
func getLogoFromURL(url string) (desc string, err error) {

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
		logger.Printf("No logo desc found for: %s", url)
		return "", nil
	}

	logger.Println("Found logos:")
	return annotations[0].Description, nil
}
