package logger

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func PrintLogFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		panic(fmt.Errorf("could not open log file at %s", path))
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		PrintLogLine(scanner.Text())
	}
}

func PrintLogLine(logLine string) {
	var unstructured map[string]string
	if err := json.Unmarshal([]byte(logLine), &unstructured); err != nil {
		panic(fmt.Errorf("could not unmarshal JSON into unstructured map"))
		return
	}

	fmt.Printf("%s [%s]: %s",
		strings.ToUpper(unstructured["level"]),
		unstructured["time"],
		unstructured["msg"],
	)
	for key, val := range unstructured {
		switch key {
		case "level":
			continue
		case "time":
			continue
		case "msg":
			continue
		default:
			fmt.Printf(" (%s => %s)", key, val)
		}
	}
	fmt.Println()
}
