package funcs

import (
	"bufio"
	"log"
	"os"
	"sync"
)

func MultiReqSend() {
	dfile, err := os.Open(*InputFile)
	var wg sync.WaitGroup
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("[MULTIREQSEND ERROR]: File dont exist")
		} else {
			log.Fatal("MULTIREQSEND ERROR]:", err)
		}
	}
	defer dfile.Close()

	scanner := bufio.NewScanner(dfile)
	urlArray := []string{}

	for scanner.Scan() {
		urlArray = append(urlArray, scanner.Text())
	}

	for _, url := range urlArray {
		wg.Add(1)
		go SendSingleRequest(url, &wg)
	}
	wg.Wait()
}
