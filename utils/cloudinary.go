package utils

import (
	"context"
	"fmt"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func credentials() (*cloudinary.Cloudinary, context.Context) {
	// Add your Cloudinary credentials, set configuration parameter
	// Secure=true to return "https" URLs, and create a context
	//===================
	cld, _ := cloudinary.New()
	cld.Config.URL.Secure = true
	ctx := context.Background()
	return cld, ctx
}
func UploadImage(cld *cloudinary.Cloudinary, ctx context.Context, filePath string) (string, error) {

	// Upload the image.
	// Set the asset's public ID and allow overwriting the asset with new versions
	resp, err := cld.Upload.Upload(ctx, filePath, uploader.UploadParams{
		PublicID:       "go_file",
		UniqueFilename: api.Bool(false),
		Overwrite:      api.Bool(true),
	})

	if err != nil {
		fmt.Println("error")
		return "", err
	}

	// // Log the delivery URL
	// fmt.Println("****2. Upload an image****\nDelivery URL:", resp.SecureURL, "\n")
	return resp.SecureURL, nil
}
