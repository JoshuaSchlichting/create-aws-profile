package main

import (
	"bufio"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path"
)

var defaultCredentialsFilePath = path.Join(getHomeDir(), ".aws", "credentials")

func main() {
	profile := flag.String("profile", "", "Name of the profile to store in ~/.aws/credentials")
	credentialsFilePath := flag.String("credentials-file", "", "Path to the credentials file")
	flag.Parse()
	if *credentialsFilePath == "" {
		credentialsFilePath = &defaultCredentialsFilePath
	}
	if *profile == "" {
		log.Fatal("--profile is required. See --help")
	}
	log.Printf("Creating profile '%s'", *profile)
	// check if there is somethinig to read on STDIN
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		creds := getCredentialsFromStdIn()
		// create file if it does not exist
		if _, err := os.Stat(*credentialsFilePath); os.IsNotExist(err) {
			log.Printf("Creating credentials file '%s'", *credentialsFilePath)
			err := os.MkdirAll(path.Dir(*credentialsFilePath), 0700)
			handleErr(err)
			err = ioutil.WriteFile(*credentialsFilePath, []byte{}, 0600)
			handleErr(err)
		}
		fileContent, err := readLines(*credentialsFilePath)
		handleErr(err)

		if fileContainsProfile(*profile, fileContent) {
			log.Printf("Profile '%s' already exists. Removing it...", *profile)
			newCredentialsPayloud := removeProfile(*profile, fileContent)
			writeLines(newCredentialsPayloud, *credentialsFilePath)
		}

		f, err := os.OpenFile(*credentialsFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		handleErr(err)
		defer f.Close()
		bufio.NewWriter(f)

		appendCredentials(creds, *profile, f)

		completedFileContents, err := readLines(*credentialsFilePath)
		handleErr(err)
		formattedLines := formatCredentials(completedFileContents)
		writeLines(formattedLines, *credentialsFilePath)
		log.Printf("Credentials stored in %s. Verify this via 'cat %s'", *credentialsFilePath, *credentialsFilePath)
	} else {
		log.Fatal("No input on STDIN")
	}
}
