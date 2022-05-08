package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/mitchellh/mapstructure"
)

type credentials struct {
	AccessKeyId     string
	SecretAccessKey string
	SessionToken    string
	Expiration      string
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

func appendCredentials(credentials credentials, profile string, buffer io.Writer) {
	fmt.Fprintf(buffer, "[%s]\n", profile)
	fmt.Fprintf(buffer, "aws_access_key_id = %s\n", credentials.AccessKeyId)
	fmt.Fprintf(buffer, "aws_secret_access_key = %s\n", credentials.SecretAccessKey)
	fmt.Fprintf(buffer, "aws_session_token = %s\n", credentials.SessionToken)
}
