package main

import (
	"fmt"
	w "github.com/NetSepio/webscrape"
)

func main() {

	w.CheckDomain("https://www.tokopedia.com")
	
	hash, _ := w.AddFileToIpfs("www.tokopedia.com")
	fmt.Printf(hash)

	w.GetObjectFromIpfs(hash, "output.txt")
	
}