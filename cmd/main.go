package main

import (
	"os"

	"github.com/sirupsen/logrus"

	"github.com/streamnative/function-mesh-workflow/commands"
)

func main() {
	if err := commands.Execute(); err != nil {
		logrus.Errorln(err)
		os.Exit(1)
	}
}
