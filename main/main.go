package main

import "github.com/hubertme/go-mod-init/sslstrip"

func main() {
	params := sslstrip.Params{
		Port:        8081,
		Filename:    "",
		PostOnly:    false,
		LogResponse: false,
	}

	sslstrip.Start(params)
}
