package models

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"cloud.google.com/go/storage"
)

// UploadFile takes an http request with an image.
// Parses the image and passes it to be sent to GCS.
func UploadFile(r *http.Request) (string, error) {
	r.ParseMultipartForm(AppConfig.MultiPartMemLimit << 20)

	file, handler, err := r.FormFile(AppConfig.FormFileName)
	if err != nil {
		return "", fmt.Errorf("Failed to read file from request: %v", err)
	}
	defer file.Close()

	imageURL, err := sendFileToGCS(file, handler.Filename)
	if err != nil {
		return "", err
	}

	return imageURL, nil
}

// GetFileByName retrieves a file from GCS by file name
func GetFileByName(fileName string) (string, error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return "", fmt.Errorf("Failed to create client: %v", err)
	}

	obj := client.Bucket(AppConfig.GCSBucket).Object(fileName)

	return getGCSImageURL(ctx, obj, fileName)
}

func getGCSImageURL(ctx context.Context, oh *storage.ObjectHandle, fileName string) (string, error) {
	if _, err := oh.Attrs(ctx); err != nil {
		return "", fmt.Errorf("readFile: unable to open file from bucket %q, file %q: %v", AppConfig.GCSBucket, fileName, err)
	}

	imageURL := fmt.Sprintf("https://storage.googleapis.com/%s/%s", AppConfig.GCSBucket, fileName)

	return imageURL, nil
}

func sendFileToGCS(f multipart.File, fileName string) (string, error) {
	// Create GCS connection
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return "", fmt.Errorf("Failed to create client: %v", err)
	}

	// Connect to bucket, create bucket object, create new writer
	bkt := client.Bucket(AppConfig.GCSBucket)
	obj := bkt.Object(fileName)
	w := obj.NewWriter(ctx)

	// Copy file into GCS
	if _, err := io.Copy(w, f); err != nil {
		return "", fmt.Errorf("Failed to upload to GCS bucket: %s - with error %v", AppConfig.GCSBucket, err)
	}

	// Close, just like writing a file. File appears in GCS after
	if err := w.Close(); err != nil {
		return "", fmt.Errorf("Failed to close file after upload: %v", err)
	}

	return getGCSImageURL(ctx, obj, fileName)
}
