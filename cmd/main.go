package main

import (
	"fmt"
	"tgBotEcho/internal/config"
	"tgBotEcho/internal/db"
)

func main() {
	fmt.Println(config.File)
	fmt.Println(db.Conn)
}
