package client

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func SendFileToFlask(fileID int64, filePath string) ([]byte, error) {
	// 如果filePath以"/uploads/"开头，需要转换为本地路径格式"./uploads/"
	localPath := filePath
	if strings.HasPrefix(filePath, "/uploads/") {
		localPath = "." + filePath
	} else if strings.HasPrefix(filePath, "uploads/") {
		localPath = "./" + filePath
	} else if strings.HasPrefix(filePath, "./uploads/") {
		localPath = filePath  // 已经是正确的本地路径格式
	}
	
	file, err := os.Open(localPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// 添加 fileID 作为表单字段
	err = writer.WriteField("file_id", strconv.FormatInt(fileID, 10))
	if err != nil {
		return nil, err
	}

	// Detect the content type based on file extension
	contentType := "application/octet-stream" // default
	fileExtension := strings.ToLower(filepath.Ext(filePath))
	switch fileExtension {
	case ".pdf":
		contentType = "application/pdf"
	case ".png":
		contentType = "image/png"
	case ".jpg", ".jpeg":
		contentType = "image/jpeg"
	case ".gif":
		contentType = "image/gif"
	case ".bmp":
		contentType = "image/bmp"
	case ".webp":
		contentType = "image/webp"
	case ".tiff", ".tif":
		contentType = "image/tiff"
	}

	// Create the form file part manually to set the content type
	header := make(map[string][]string)
	header["Content-Disposition"] = []string{`form-data; name="image"; filename="` + filepath.Base(filePath) + `"`}
	header["Content-Type"] = []string{contentType}

	part, err := writer.CreatePart(header)
	if err != nil {
		return nil, err
	}

	// Copy the file content to the form field
	_, err = io.Copy(part, file)
	if err != nil {
		return nil, err
	}

	// Close the writer to finalize the multipart form data
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	// Create the HTTP request
	req, err := http.NewRequest(
		"POST",
		"http://localhost:5000/process_image",
		body,
	)
	if err != nil {
		return nil, err
	}

	// Set the content type header
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{Timeout: 60 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
