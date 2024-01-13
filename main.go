package main

import (
	"flag"
	"fmt"
	"sync"

	"getGo/funcs"
)

func main() {
	flag.Parse()
	args := flag.Args()

	funcs.DeleteLog()

	if *funcs.HelpMessage {
		flag.Usage()
		return
	}

	if *funcs.Mirror {
		htmlContent, err := funcs.FetchHTML(args[0])
		if err != nil {
			fmt.Println("Error fetching HTML:", err)
			return
		}

		// Parse HTML and extract resources
		resources, err := funcs.ExtractResources(htmlContent, args[0])
		if err != nil {
			fmt.Println("Error extracting resources:", err)
			return
		}

		// Download and save resources
		err = funcs.DownloadResources(resources, args[0])
		if err != nil {
			fmt.Println("Error downloading resources:", err)
			return
		}

		fmt.Println("Resources downloaded successfully!")
		return
	}

	if *funcs.BgMode {
		fmt.Println("Logs will be written to wget-log")
	}

	// Use a separate wait group for the goroutines
	var wg sync.WaitGroup

	if *funcs.InputFile != "" {
		funcs.MultiReqSend()
	} else {
		for _, arg := range args {
			currentArg := arg
			// Increment the new wait group for each goroutine
			wg.Add(1)
			go funcs.SendSingleRequest(currentArg, &wg)
		}
	}

	// Wait for all downloads to complete
	wg.Wait()
}
