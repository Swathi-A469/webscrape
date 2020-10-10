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

//CheckDomain if the domain exists, write its content to a file else return error
func CheckDomain(domain string) (err error) {
	// Make HTTP request
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, domain, nil)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("User-Agent", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_5) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11`)
	req.Header.Add("Accept", "*/*")
	response, err := client.Do(req)
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

//AddFileToIpfs adds the specified file to IPFS and returns hash
func AddFileToIpfs(filePath string) string {
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

//GetObjectFromIpfs get object from ipfs and writes to the specified file
func GetObjectFromIpfs(Hash string, filePath string) {
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