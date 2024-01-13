package funcs

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func DeleteLog() {
	filename := "wget-log"

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return
	}

	err := os.Remove(filename)
	if err != nil {
		log.Fatal("ERROR RESETING WGET-LOG", err)
	}
}

func WriteTextToWgetLog(text string) (n int, e error) {
	file, err := os.OpenFile("wget-log.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0o644)
	if err != nil {
		log.Fatal("ERROR WRITING TO wget-log:", err)
	}
	defer file.Close()

	_, err = file.WriteString(text)
	if err != nil {
		log.Fatal("ERROR WRITING TO wget-log:", err)
	}
	return 0, nil
}

func OutputString(input string, processFunc func(a string) (n int, e error)) {
	processFunc(input)
}

func GetRateLimitInBytes() (int, error) {
	s := *RateLimit

	if s == "" {
		return -1, nil
	}

	index := len(s) - 1
	for i := len(s) - 1; i >= 0; i-- {
		if !unicode.IsDigit(rune(s[i])) {
			break
		}
		index = i
	}

	numericPart, err := strconv.Atoi(s[:index])
	if err != nil {
		return 0, err
	}

	letterPart := s[index:]

	if strings.EqualFold(letterPart, "k") {
		numericPart *= 1024
	} else if strings.EqualFold(letterPart, "m") {
		numericPart *= 1000000
	} else if strings.EqualFold(letterPart, "g") {
		numericPart *= 1000000000
	} else {
		fmt.Println("weird limit, wth are you downloading")
		os.Exit(1)
	}

	return numericPart, nil
}
