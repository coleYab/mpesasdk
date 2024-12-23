package auth

import (
	"log"
	"testing"
)

func TestGetBase64EncodedUserCredential(t *testing.T) {
	testCases := [][]string{
		{"userCredential", "userSecret"},
		{"userCredential1", "userSecret1"},
		{"sdjhfjksdhfjkds", "klsdjflksdjflksdj"},
		{"sdfkjhsdkj", "sdfjsd"},
		{"dsjfhsdjkfh", "dsfsdfsd"},
	}

	expecteds := []string{
		"dXNlckNyZWRlbnRpYWw6dXNlclNlY3JldA==",
		"dXNlckNyZWRlbnRpYWwxOnVzZXJTZWNyZXQx",
		"c2RqaGZqa3NkZmhqazM6a2xzZGpmbGtqZGZsazM=",
		"c2Rma2poc2RrajpzZGZqc2Q=",
		"ZHNqZmhzZGtmaDpkc2ZzZGY=",
	}

	for idx, testCase := range testCases {
		got := GetBase64EncodedUserCredential(testCase[0], testCase[1])
		expected := expecteds[idx]

		if got != expected {
			log.Fatalf("ERROR at test %d: while encoding to base 64 for userCredential - [%s], userSecret - [%s]\nExpected: %s, Got: %s\n",
				idx+1, testCase[0], testCase[1], expected, got,
			)
		}
	}
}
