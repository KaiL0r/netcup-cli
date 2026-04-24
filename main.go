package main

import (
	"net/http"
	"os"

	"github.com/KaiL0r/netcup-cli/api"
	"github.com/KaiL0r/netcup-cli/auth"
	"github.com/KaiL0r/netcup-cli/cmd"
)

func main() {
	apiClient := api.NewClient(
		"",
		auth.NewAuthService(
			auth.NewHTTPOAuthClient(http.DefaultClient),
			auth.NewFileStorage(os.Getenv("NETCUP_TOKEN_PATH")),
			auth.RealClock{},
		),
		http.DefaultClient,
	)

	cmd.Execute(apiClient)
}
