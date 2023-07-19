package prisma

import "github.com/Chahine-tech/chrashind/prisma/db"

// func EnableConnect() error {
// 	client := db.NewClient()
// 	if err := client.Prisma.Connect(); err != nil {
// 		return err
// 	}

// 	defer func() {
// 		if err := client.Prisma.Disconnect(); err != nil {
// 			panic(err)
// 		}
// 	}()
// 	return nil
// }

var Connect = db.NewClient().Prisma.Connect()
