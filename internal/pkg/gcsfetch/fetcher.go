package gcsfetch

import (
	"cloud.google.com/go/storage"
	"context"
	"errors"
	"fmt"
	"io"
	"net/url"
	"os"
)

func Fetch(gcsObject string, outputFilename string) (int64, error) {
	gcsURL, err := url.Parse(gcsObject)
	if err != nil {return 0, err}
	if gcsURL.Scheme != "gs" {
		return 0, errors.New("URL should start with gs://")
	}

	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil { return 0, err }

	objectName := gcsURL.Path[1:]

	fmt.Printf("getting bucket %s and object %s\n", gcsURL.Host, objectName)

	reader, err := client.Bucket(gcsURL.Host).Object(objectName).NewReader(ctx)
	if err != nil { return 0, err }

	defer reader.Close()

	dest,err := os.Create(outputFilename)
	if err != nil { return 0, err }

	defer dest.Close()

	bytesCopied, err := io.Copy(dest, reader)
	if err != nil { return 0, err }
	return bytesCopied, nil
}
