package prisma

import (
	"github.com/Chahine-tech/chrashind/prisma/db"
)

func PrismaClient() *db.PrismaClient {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()
	return client
}
