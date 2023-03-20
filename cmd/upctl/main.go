package main

import (
	"github.com/uptime-com/uptime-client-go/v2/internal/upctl"
)

var version = "0.0.0"

func main() {
	upctl.Execute(version)
}
