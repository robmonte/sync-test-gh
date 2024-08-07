package main

import (
	"encoding/json"
	"log"
	"os"
	"strings"
)

func main() {
	testType, exists := os.LookupEnv("STOREGH_TEST_TYPE")
	if !exists {
		log.Fatalf("missing required \"STOREGH_TEST_TYPE\" field to determine test type")
	}

	log.Printf("running %q test type\n", strings.ToLower(testType))

	switch testType {
	case "WRITE_TEST_TOKEN":
		writeTypeTokenTestCheck()

	case "WRITE_TEST_APP":
		writeTypeAppTestCheck()

	case "UPDATE_TEST_TOKEN":
		updateTypeTokenTestCheck()

	case "UPDATE_TEST_APP":
		updateTypeAppTestCheck()

	default:
		log.Fatalf("received unsupported test type")
	}

	log.Println("Success!")
}

// Token authenticated Write checks
func writeTypeTokenTestCheck() {
	// Token Write Test 1
	envVar1 := "STOREGH_TEST_TOKEN_WRITE_KEY_1_BAD_SYMBOLS________"
	secret, found := os.LookupEnv(envVar1)
	if !found || secret == "" {
		log.Fatalf("failed to find %q value\n", envVar1)
	}
	if secret != "I am secret one!" {
		log.Fatalf("received %q value but did not match expected value\n", envVar1)
	} else {
		log.Printf("the value of %q matched the expected value\n", envVar1)
	}

	// Token Write Test 2
	envVar2 := "STOREGH_TEST_TOKEN_WRITE_KEY_2"
	secret, found = os.LookupEnv(envVar2)
	if !found || secret == "" {
		log.Fatalf("failed to find %q value\n", envVar2)
	}
	if secret != "I am secret two!" {
		log.Fatalf("received %q value but did not match expected value\n", envVar2)
	} else {
		log.Printf("the value of %q matched the expected value\n", envVar2)
	}

	// Token Write Test 3
	envVar3 := "STOREGH_TEST_TOKEN_WRITE_KEY_3"
	secret, found = os.LookupEnv(envVar3)
	if !found || secret == "" {
		log.Fatalf("failed to find %q value\n", envVar3)
	}
	if secret != "1234" {
		log.Fatalf("received %q value but did not match expected value\n", envVar3)
	} else {
		log.Printf("the value of %q matched the expected value\n", envVar3)
	}

	// Token Write Test 4
	envVar4 := "STOREGH_TEST_TOKEN_WRITE_KEY_4"
	secret, found = os.LookupEnv(envVar4)
	if !found || secret == "" {
		log.Fatalf("failed to find %q value\n", envVar4)
	}
	if secret != "{\"0\":\"my\",\"1\":\"map\",\"2\":\"value\"}" {
		log.Fatalf("received %q value but did not match expected value\n", envVar4)
	} else {
		var val any
		if err := json.Unmarshal([]byte(secret), &val); err != nil {
			log.Fatalf("received %q value but it could not be unmarshalled into valid JSON", envVar4)
		}

		log.Printf("the value of %q matched the expected value\n", envVar4)
	}

	// Token Write Test 5
	envVar5 := "STOREGH_TEST_TOKEN_WRITE_KEY_5"
	secret, found = os.LookupEnv(envVar5)
	if !found || secret == "" {
		log.Fatalf("failed to find %q value\n", envVar5)
	}
	if secret != "{\"number\":1,\"true\":false}" {
		log.Fatalf("received %q value but did not match expected value\n", envVar5)
	} else {
		var val any
		if err := json.Unmarshal([]byte(secret), &val); err != nil {
			log.Fatalf("received %q value but it could not be unmarshalled into valid JSON", envVar5)
		}

		log.Printf("the value of %q matched the expected value\n", envVar5)
	}
}

// App authenticated Write checks
func writeTypeAppTestCheck() {
	// App Write Tets 1
	envVar5 := "STOREGH_TEST_APP_WRITE_KEY_1"
	secret, found := os.LookupEnv(envVar5)
	if !found || secret == "" {
		log.Fatalf("failed to find %q value\n", envVar5)
	}
	if secret != "12345" {
		log.Fatalf("received %q value but did not match expected value\n", envVar5)
	} else {
		log.Printf("the value of %q matched the expected value\n", envVar5)
	}
}

// Token authenticated Update checks
func updateTypeTokenTestCheck() {
	// Token Update Test 1
	envVar1 := "STOREGH_TEST_TOKEN_UPDATE_KEY_1"
	secret, found := os.LookupEnv(envVar1)
	if !found || secret == "" {
		log.Fatalf("failed to find %q value\n", envVar1)
	}
	if secret != "value-1-updated" {
		log.Fatalf("received %q value but did not match expected value\n", envVar1)
	} else {
		log.Printf("the value of %q matched the expected value\n", envVar1)
	}

	// Token Update Test 2
	envVar2 := "STOREGH_TEST_TOKEN_UPDATE_KEY_2"
	secret, found = os.LookupEnv(envVar2)
	if !found || secret == "" {
		log.Fatalf("failed to find %q value\n", envVar2)
	}
	if secret != "{\"0\":\"value-2-updated\"}" {
		log.Fatalf("received %q value but did not match expected value\n", envVar2)
	} else {
		log.Printf("the value of %q matched the expected value\n", envVar2)
	}
}

// App authenticated Update checks
func updateTypeAppTestCheck() {
	// App Update Test 1
	envVar3 := "STOREGH_TEST_APP_UPDATE_KEY_1"
	secret, found := os.LookupEnv(envVar3)
	if !found || secret == "" {
		log.Fatalf("failed to find %q value\n", envVar3)
	}
	if secret != "value-1-updated" {
		log.Fatalf("received %q value but did not match expected value\n", envVar3)
	} else {
		log.Printf("the value of %q matched the expected value\n", envVar3)
	}
}
