package processor

import (
	"context"
	"fmt"

	vision "cloud.google.com/go/vision/apiv1"
)

// Process processes passed URL and returns map of attributes
func Process(ctx context.Context, url string) error {

	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		return err
	}

	image := vision.NewImageFromURI(url)
	annotations, err := client.DetectLogos(ctx, image, nil, 10)
	if err != nil {
		return err
	}

	if len(annotations) == 0 {
		fmt.Println("No logos found.")
	} else {
		fmt.Println("Logos:")
		for _, annotation := range annotations {
			fmt.Println(annotation.Description)
		}
	}

	return nil
}
