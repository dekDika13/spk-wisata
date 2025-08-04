package utils

import (
	"context"
	"mime/multipart"
	"os"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

func UploadToCloudinary(file multipart.File, fileName string) (string, string, error) {
	cloudName := os.Getenv("CLOUDINARY_CLOUD_NAME")
	apiKey := os.Getenv("CLOUDINARY_API_KEY")
	apiSecret := os.Getenv("CLOUDINARY_API_SECRET")

	cld, err := cloudinary.NewFromParams(cloudName, apiKey, apiSecret)
    if err != nil {
        return "", "", err
    }

    uploadParams := uploader.UploadParams{
        PublicID: fileName,
        Folder:  "profile",
    }

    uploadResult, err := cld.Upload.Upload(context.Background(), file, uploadParams)
    if err != nil {
        return "", "", err
    }
    

    return uploadResult.SecureURL, uploadResult.PublicID, nil
}

func DeleteFromCloudinary(publicID string) error {
    cloudName := os.Getenv("CLOUDINARY_CLOUD_NAME")
	apiKey := os.Getenv("CLOUDINARY_API_KEY")
	apiSecret := os.Getenv("CLOUDINARY_API_SECRET")
    cld, err := cloudinary.NewFromParams(cloudName, apiKey, apiSecret)
    if err != nil {
        return err
    }

    _, err = cld.Upload.Destroy(context.Background(), uploader.DestroyParams{
        PublicID: publicID,
    })
    return err
}
