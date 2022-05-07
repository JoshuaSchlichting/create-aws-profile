package main

import "fmt"

type credentials struct {
	AccessKeyId     string
	SecretAccessKey string
	SessionToken    string
	Expiration      string
}

func (c credentials) String() string {
	fmtStrings := []string{
		"aws_access_key_id = %s\n",
		"aws_secret_access_key = %s\n",
		"aws_session_token = %s\n",
	}
	var stringVal string
	stringVal += fmt.Sprintf(fmtStrings[0], c.AccessKeyId)
	stringVal += fmt.Sprintf(fmtStrings[1], c.SecretAccessKey)
	stringVal += fmt.Sprintf(fmtStrings[2], c.SessionToken)

	return stringVal
}
