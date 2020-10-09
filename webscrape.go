package webscrape

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
)

func checkDomain(domain string) (err error) {
	// Make HTTP request
	response, err := http.Get(domain)
	if err != nil {
		return err
	}
	fmt.Printf(response.Status)
	defer response.Body.Close()

	parsedURL, err := url.Parse(domain)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Create output file
	outFile, err := os.Create(parsedURL.Host)
	if err != nil {
		fmt.Println(err)
	}
	defer outFile.Close()

	// Copy data from the response to standard output
	n, err := io.Copy(outFile, response.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Number of bytes copied:", n)
	return nil
}

func addFileToIpfs(filePath string) string {
	url := "https://ipfs.infura.io:5001/api/v0/add"
	method := "POST"
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, errFile2 := os.Open(filePath)
	defer file.Close()
	part2,
		errFile2 := writer.CreateFormFile("file", filePath)
	_, errFile2 = io.Copy(part2, file)
	if errFile2 != nil {

		fmt.Println(errFile2)
	}
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	var response map[string]string
	json.NewDecoder(res.Body).Decode(&response)
	fmt.Println(response)
	return response["Hash"]
}

func getObjectFromIpfs(Hash string, filePath string) {
	fmt.Println(Hash)
	url := "https://ipfs.infura.io:5001/api/v0/cat?arg="
	url += Hash
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	// Create output file
	outFile, err := os.Create(filePath)
	if err != nil {
		fmt.Println(err)
	}
	defer outFile.Close()

	// Copy data from the response to standard output
	n, err := io.Copy(outFile, res.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Number of bytes copied :", n)
}
