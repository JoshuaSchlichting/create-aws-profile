package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/mitchellh/mapstructure"
)

var credentialsFilePath = path.Join(getHomeDir(), ".aws", "credentials")

type credentials struct {
	AccessKeyId     string
	SecretAccessKey string
	SessionToken    string
	Expiration      string
}

func main() {
	profile := flag.String("profile", "", "Name of the profile to store in ~/.aws/credentials")
	flag.Parse()
	if *profile == "" {
		log.Fatal("--profile is required. See --help")
	}
	log.Printf("Creating profile '%s'", *profile)
	// check if there is somethinig to read on STDIN
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		creds := getCredentialsFromStdIn()

		fileContent, err := readLines(credentialsFilePath)
		handleErr(err)

		if fileContainsProfile(fileContent, *profile) {
			log.Printf("Profile '%s' already exists. Removing it...", *profile)
			newCredentialsPayloud := removeCredentials(*profile, fileContent)
			writeLines(newCredentialsPayloud, credentialsFilePath)
		}

		f, err := os.OpenFile(credentialsFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		handleErr(err)
		defer f.Close()
		appendCredentials(creds, *profile, f)
		log.Printf("Credentials stored in %s. Verify this via 'cat %s'", credentialsFilePath, credentialsFilePath)
	} else {
		log.Fatal("No input on STDIN")
	}
}

func removeCredentials(profile string, payload []string) []string {

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

func fileContainsProfile(lines []string, profile string) bool {
	// search for profile
	for _, line := range lines {
		if strings.HasPrefix(line, fmt.Sprintf("[%s]", profile)) {
			return true
		}
	}
	return false
}

func getCredentialsFromStdIn() credentials {
	var stdin []byte

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		stdin = append(stdin, scanner.Bytes()...)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	m := make(map[string]interface{})
	err := json.Unmarshal(stdin, &m)
	handleErr(err)

	var credentials credentials
	err = mapstructure.Decode(m["Credentials"], &credentials)
	handleErr(err)
	return credentials
}

func appendCredentials(credentials credentials, profile string, file *os.File) {
	file.WriteString(fmt.Sprintf("\n[%s]\n", profile))
	file.WriteString(fmt.Sprintf("aws_access_key_id = %s\n", credentials.AccessKeyId))
	file.WriteString(fmt.Sprintf("aws_secret_access_key = %s\n", credentials.SecretAccessKey))
	file.WriteString(fmt.Sprintf("aws_session_token = %s\n", credentials.SessionToken))
}

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
