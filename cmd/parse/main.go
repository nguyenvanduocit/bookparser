
// Sample vision-quickstart uses the Google Cloud Vision API to label an image.
package main

import (
	"context"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"

	vision "cloud.google.com/go/vision/apiv1"
)

func main() {
	ctx := context.Background()

	// Creates a client.
	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()
	for i := 1; i <= 190; i++ {
		 fileName := fmt.Sprintf("../images/page-%d.png", i)
		 text, err := getText(ctx, fileName, client)
		if err != nil {
			log.Fatal(err)
		}

		ioutil.WriteFile(fmt.Sprintf("../images/page-%d.txt", i), []byte(text), fs.ModePerm)
	}
}

func getText(ctx context.Context, filePath string, client *vision.ImageAnnotatorClient) (string, error) {

	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	image, err := vision.NewImageFromReader(file)
	if err != nil {
		return "", err
	}

	result, err := client.DetectDocumentText(ctx, image, nil)
	if err != nil {
		return "", err
	}

	return result.GetText(), nil
}
