package webscrape

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	ipfsGateway "github.com/ipfs/go-ipfs-api"
)

// CheckDomain if the domain exists, write its content to a file else return error
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
func AddFileToIpfs(filePath string) (string, error) {
	ig := ipfsGateway.NewShell("https://ipfs.infura.io:5001")
	// Create io reader from a local file
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
	}

	//Uploads file to ipfs and returns metahash
	hash, err := ig.Add(file)
	if err != nil {
		fmt.Println(err)
	}
	return hash, err
}

//GetObjectFromIpfs get object from ipfs and writes to the specified file
func GetObjectFromIpfs(Hash string, filePath string) error {

	ig := ipfsGateway.NewShell("https://ipfs.infura.io:5001")

	err := ig.Get(Hash, filePath)
	if err != nil {
		fmt.Println(err)
	}
	return err
}
