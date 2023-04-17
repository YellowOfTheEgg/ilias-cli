package util

import (
	"fmt"
	"ilias-cli/ilias_api"
	"golang.org/x/crypto/ssh/terminal"
	"log"
	"os"
	"strings"
	"syscall"
)

const (
	envUser = "ILIAS_USER"
	envPassword = "ILIAS_PASS"
)

func GetCredentials() *ilias_api.Credentials {

	// Get the username
	username, present := os.LookupEnv(envUser)
	if present == false {
		fmt.Fprint(os.Stderr, "Username for 'https://ilias.hhu.de': ")
		_, err := fmt.Scanln(&username)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Get the password
	password, present := os.LookupEnv(envPassword)
	if present == false {
		fmt.Fprint(os.Stderr, "Password for 'https://ilias.hhu.de': ")
		inputBytes, err := terminal.ReadPassword(int(syscall.Stdin))
		if err != nil {
			log.Fatal(err)
		}
		password = string(inputBytes)
		fmt.Fprintln(os.Stderr)
	}

	return &ilias_api.Credentials{
		Username: strings.ToLower(username),
		Password: password,
	}
}
