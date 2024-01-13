package funcs

import (
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"sync"
	"time"

	"github.com/schollz/progressbar/v3"
)

func SendSingleRequest(url string, wg *sync.WaitGroup) error {
	defer wg.Done()
	var mainname string
	var outputtingFunc func(string) (n int, e error)
	limit, _ := GetRateLimitInBytes()
	if *NameOfOutput == "" {
		mainname = FormatfileName(url)
	} else {
		mainname = *NameOfOutput
	}

	if limit == -1 {
		limit = 100000000
	}

	// check if silent Mode is on, and switch the outputting function
	// accordingly
	if *BgMode {
		outputtingFunc = WriteTextToWgetLog
	} else {
		outputtingFunc = os.Stdout.WriteString
	}

	if !(*PathVar == "") {
		_, err := os.Stat(*PathVar)
		if err != nil {
			if os.IsNotExist(err) {
				OutputString("Specified Path Doesnt Exist, saving to the current Directory", outputtingFunc)
			} else {
				log.Fatal("[reqsend error]:", err)
			}
		} else {
			mainname = path.Join(*PathVar, mainname)
		}
	}

	OutputString("Download started at "+time.Now().String()+"\n", outputtingFunc)

	file, err := os.Create(mainname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	OutputString("sent request, awaiting response...", outputtingFunc)
	serverResponse, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer serverResponse.Body.Close()

	if serverResponse.StatusCode != http.StatusOK {
		log.Fatal("SERVER RETURNED A NOT-OK RESPONSE " + strconv.Itoa(serverResponse.StatusCode))
	} else {
		OutputString("\033[1;32m"+"200 OK"+"\033[0m"+"\n", outputtingFunc)
	}

	OutputString("\033[1;33m"+"file is to be saved in "+mainname+"\033[0m"+"\n", outputtingFunc)

	buffer := make([]byte, limit)

	bar := progressbar.DefaultBytes(serverResponse.ContentLength, "Downloading.. ")

	if !*BgMode {
		for {
			n, err := serverResponse.Body.Read(buffer)
			if err == io.EOF {
				break
			} else if err != nil {
				return err
			}
			_, err = file.Write(buffer[:n])
			if err != nil {
				log.Fatal("[IOWRITE ERROR],", err)
			}
			bar.Add(n)
			time.Sleep(time.Millisecond)
		}
	} else {
		_, err := io.CopyBuffer(file, serverResponse.Body, buffer)
		if err != nil {
			log.Fatal(err)
		}
	}

	OutputString("Downloaded "+url+"\n", outputtingFunc)
	OutputString("Download Finished at "+time.Now().String()+"\n", outputtingFunc)

	return nil
}
