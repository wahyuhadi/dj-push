package services

import (
	"bytes"
	"dj-push/models"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func Push(data *models.Data) bool {

	url := data.URI
	method := "POST"
	is_scan_type, _ := scan_type(data.ScanType)
	is_token := fmt.Sprintf("Token %s", data.Token)

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("minimum_severity", "Info")
	_ = writer.WriteField("active", "true")
	_ = writer.WriteField("verified", "true")
	_ = writer.WriteField("scan_type", is_scan_type)

	file, err := os.Open(data.File)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer file.Close()

	part5, _ := writer.CreateFormFile("file", filepath.Base(data.File))
	_, errFile := io.Copy(part5, file)
	if errFile != nil {
		fmt.Println(errFile)
		return false
	}
	_ = writer.WriteField("engagement", data.Engagement)
	_ = writer.WriteField("close_old_findings", "false")
	_ = writer.WriteField("push_to_jira", "false")
	err = writer.Close()
	if err != nil {
		fmt.Println(err)
		return false
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return false
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", is_token)

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer res.Body.Close()

	if res.StatusCode == 200 || res.StatusCode == 201 {
		return true
	}

	return false

}
