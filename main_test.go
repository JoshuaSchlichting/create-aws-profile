package main

import (
	"bufio"
	"bytes"
	"reflect"
	"strings"
	"testing"
)

var testFileContent = []string{
	"[test_profile]",
	"aws_access_key_id = test_access_key_id",
	"aws_secret_access_key = test_secret_access_key",
	"aws_session_token = test_session_token",
	"",
	"[test_profile2]",
	"aws_access_key_id = test_access_key_id2",
	"aws_secret_access_key = test_secret_access_key2",
	"aws_session_token = test_session_token2",
	"",
	"[test_profile3]",
	"aws_access_key_id = test_access_key_id3",
	"aws_secret_access_key = test_secret_access_key3",
	"aws_session_token = test_session_token3",
}

func TestRemoveProfile(t *testing.T) {
	expectedFileContent := []string{
		"[test_profile]",
		"aws_access_key_id = test_access_key_id",
		"aws_secret_access_key = test_secret_access_key",
		"aws_session_token = test_session_token",
		"",
		"[test_profile3]",
		"aws_access_key_id = test_access_key_id3",
		"aws_secret_access_key = test_secret_access_key3",
		"aws_session_token = test_session_token3",
	}

	newCredentials := removeProfile("test_profile2", testFileContent)
	if !reflect.DeepEqual(newCredentials, expectedFileContent) {
		t.Errorf("Expected %v, got %v", expectedFileContent, newCredentials)
	}
}

func TestFileContainsProfile(t *testing.T) {
	fileContent := []string{
		"[test_profile]",
		"aws_access_key_id = test_access_key_id",
		"[test_profile2]",
		"aws_access_key_id = test_access_key_id2",
		"[test_profile3]",
		"aws_access_key_id = test_access_key_id3",
	}
	buff := bytes.Buffer{}
	for _, line := range fileContent {
		buff.WriteString(line)
		buff.WriteString("\n")
	}
	expected := true

	actual := fileContainsProfile("test_profile2", strings.Split(buff.String(), "\n"))
	if expected != actual {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestAddProfile(t *testing.T) {
	buffer := bytes.Buffer{}
	for _, line := range testFileContent {
		buffer.WriteString(line)
		buffer.WriteString("\n")
	}
	credentials := credentials{
		AccessKeyId:     "test_access_key_id",
		SecretAccessKey: "test_secret_access_key",
		SessionToken:    "test_session_token",
	}
	profile := "test_profile4"
	w := bufio.NewWriter(&buffer)
	appendCredentials(credentials, profile, w)
	w.Flush()

	if !fileContainsProfile(profile, strings.Split(buffer.String(), "\n")) {
		t.Errorf("Expected to see %v, got %v", profile, buffer.String())
	}
}

func TestFormatCredentials(t *testing.T) {
	uglyCredentials := []string{
		"[test_profile]",
		"aws_access_key_id = test_access_key_id",
		"aws_secret_access_key = test_secret_access_key",
		"aws_session_token = test_session_token",
		"", "", "", "",
		"[test_profile2]",
		"aws_access_key_id = test_access_key_id2",
		"aws_secret_access_key = test_secret_access_key2",
		"aws_session_token = test_session_token2",
		"", "", "", "",
		"[test_profile3]",
		"aws_access_key_id = test_access_key_id3",
		"aws_secret_access_key = test_secret_access_key3",
		"aws_session_token = test_session_token3",
	}
	formattedCredentials := formatCredentials(uglyCredentials)

	if !reflect.DeepEqual(testFileContent, formattedCredentials) {
		t.Errorf("Expected %v, got %v", testFileContent, formattedCredentials)
	}
}
