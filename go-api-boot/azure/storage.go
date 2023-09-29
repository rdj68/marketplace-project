package azure

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/url"
	"os"

	"github.com/Azure/azure-storage-blob-go/azblob"
	"github.com/SaiNageswarS/go-api-boot/logger"
)

var Storage storageWrapper = storageWrapper{}

type storageWrapper struct{}

// Uploads a stream to Azure storage.
// containerName - Azure Container Name.
// path - Azure path for the object like profile-photos/photo.jpg
func (s storageWrapper) UploadStream(containerName, path string, imageData bytes.Buffer) (chan string, chan error) {
	resultChan := make(chan string)
	errorChan := make(chan error)

	go func() {
		accountName, accountKey := os.Getenv("AZURE-STORAGE-ACCOUNT"), os.Getenv("AZURE-STORAGE-ACCESS-KEY")
		if len(accountName) == 0 || len(accountKey) == 0 {
			logger.Error("Either the AZURE_STORAGE_ACCOUNT or AZURE_STORAGE_ACCESS_KEY environment variable is not set")
			err := errors.New("missing azure account or access key")
			errorChan <- err
			return
		}

		// Create a default request pipeline using your storage account name and account key.
		credential, err := azblob.NewSharedKeyCredential(accountName, accountKey)
		if err != nil {
			logger.Error("Invalid credentials with error: " + err.Error())
			errorChan <- err
			return
		}
		p := azblob.NewPipeline(credential, azblob.PipelineOptions{})

		URL, _ := url.Parse(
			fmt.Sprintf("https://%s.blob.core.windows.net/%s", accountName, containerName))

		containerURL := azblob.NewContainerURL(*URL, p)
		blobURL := containerURL.NewBlockBlobURL(path)
		_, err = azblob.UploadBufferToBlockBlob(context.Background(), imageData.Bytes(), blobURL, azblob.UploadToBlockBlobOptions{
			BlockSize:   4 * 1024 * 1024,
			Parallelism: 16})

		if err != nil {
			logger.Error("Failed uploading image.")
			errorChan <- err
			return
		}

		uploadPath := fmt.Sprintf("https://%s.blob.core.windows.net/%s/%s", accountName, containerName, path)
		resultChan <- uploadPath
	}()

	return resultChan, errorChan
}
