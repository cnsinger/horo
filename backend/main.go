package main

import (
	"horo/cmd"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	cmd.Execute()
}
