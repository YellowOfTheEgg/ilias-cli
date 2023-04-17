package util

import (
	"fmt"
	"os"

	"ilias-cli/ilias_api"
)

func NewIliasClient() (*ilias_api.Client) {
	credentials := GetCredentials()
	spin := StartSpinner("Logging in at 'https://ilias.hhu.de'")
	client, err := ilias_api.NewClient(nil, credentials)
	if err != nil {
		spin.StopError(err)
		os.Exit(1)
	}

	spin.StopSuccess(fmt.Sprintf("Logged in as %s!", client.User.Username))
	WriteUserCache(client.User.Username)
	return client
}
