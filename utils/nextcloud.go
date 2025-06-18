package utils

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

// UploadToNextcloud uploads a file to Nextcloud via WebDAV and returns the public URL
func UploadToNextcloud(file *multipart.FileHeader) (string, error) {
	webdavURL := os.Getenv("NEXTCLOUD_WEBDAV_URL") // e.g. https://nextcloud.example.com/remote.php/dav/files/username/
	username := os.Getenv("NEXTCLOUD_USERNAME")
	password := os.Getenv("NEXTCLOUD_PASSWORD")
	publicURLBase := os.Getenv("NEXTCLOUD_PUBLIC_URL") // base URL to generate public link, e.g. https://nextcloud.example.com/s/...

	src, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer src.Close()

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, src)
	if err != nil {
		return "", fmt.Errorf("failed to copy file data: %w", err)
	}

	// Generate file name
	filename := filepath.Base(file.Filename)
	uploadPath := webdavURL + filename

	req, err := http.NewRequest("PUT", uploadPath, buf)
	if err != nil {
		return "", fmt.Errorf("failed to create upload request: %w", err)
	}
	req.SetBasicAuth(username, password)
	req.Header.Set("Content-Type", "application/octet-stream")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("upload request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 201 && resp.StatusCode != 204 {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("upload failed with status %d: %s", resp.StatusCode, string(body))
	}

	// Return direct link (if using public folder or pre-shared folder)
	publicURL := fmt.Sprintf("%s/%s", publicURLBase, filename)
	return publicURL, nil
}
