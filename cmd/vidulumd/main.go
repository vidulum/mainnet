package main

import (
	"os"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/vidulum/vidulum/app"
)

func main() {
	setAddressPrefixes(app.AccountAddressPrefix)
	rootCmd := NewRootCmd(app.MakeEncodingConfig())
	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
