package main

import (
	_ "go.uber.org/automaxprocs"

	"stash.lamoda.ru/neurocapsule/neurocapsule/cmd"
)

func main() {
	cmd.Execute()
}
