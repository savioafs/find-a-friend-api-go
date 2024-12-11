package main

import (
	"github.com/savioafs/find-a-friend-api-go/internal/config"
	"github.com/savioafs/find-a-friend-api-go/internal/infra/web"
)

func main() {
	dbConn, _, _, err := config.LoadConfigs()
	if err != nil {
		panic(err)
	}

	server := web.NewServer(dbConn)

	server.Run("8080")
}
