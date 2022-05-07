package main

import (
	"log"
	"reflect"
	"testing"
)

var testFileContent = []string{
	"[test_profile]",
	"aws_access_key_id = test_access_key_id",
	"aws_secret_access_key = test_secret_access_key",
	"aws_session_token = test_session_token",

	"[test_profile2]",
	"aws_access_key_id = test_access_key_id2",
	"aws_secret_access_key = test_secret_access_key2",
	"aws_session_token = test_session_token2",

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

		"[test_profile3]",
		"aws_access_key_id = test_access_key_id3",
		"aws_secret_access_key = test_secret_access_key3",
		"aws_session_token = test_session_token3",
	}

	newCredentials := removeCredentials("test_profile2", testFileContent)
	log.Println(newCredentials)
	if !reflect.DeepEqual(newCredentials, expectedFileContent) {
		t.Errorf("Expected %v, got %v", expectedFileContent, newCredentials)
	}
}

func TestFileContainsProfile(t *testing.T) {
	fileContent := []string{
		"[test_profile]",
		"aws_access_key_id = test_access_key_id",
	}
	expected := true
	actual := fileContainsProfile(fileContent, "test_profile")
	if expected != actual {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
