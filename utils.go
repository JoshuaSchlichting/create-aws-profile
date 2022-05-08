package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
)

func getHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// writeLines writes the lines to the given file.
func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

func formatCredentials(lines []string) (formattedLines []string) {
	firstProfileIsLoaded := false
	for _, line := range lines {
		if line == "" {
			continue
		}
		if !firstProfileIsLoaded {
			if strings.HasPrefix(line, "[") {
				firstProfileIsLoaded = true
				formattedLines = append(formattedLines, line)
				continue
			}
		}
		if firstProfileIsLoaded && strings.HasPrefix(line, "[") {
			formattedLines = append(formattedLines, "")
		}
		formattedLines = append(formattedLines, line)
	}
	return
}

func removeProfile(profile string, payload []string) []string {

	var lines []string
	var indexInProfile bool = false
	for _, line := range payload {
		if strings.HasPrefix(line, "[") && indexInProfile {
			indexInProfile = false
		}
		if !indexInProfile {
			if strings.HasPrefix(line, fmt.Sprintf("[%s]", profile)) {
				indexInProfile = true
				continue
			} else {
				lines = append(lines, line)
			}
		}
	}
	return lines
}

func fileContainsProfile(profile string, lines []string) bool {
	// search for profile
	for _, line := range lines {
		if strings.HasPrefix(line, fmt.Sprintf("[%s]", profile)) {
			return true
		}
	}
	return false
}
