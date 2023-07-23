package links

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Chahine-tech/chrashind/internal/users"
	"github.com/Chahine-tech/chrashind/prisma/db"
)

type Link struct {
	ID      string
	Title   string
	Address string
	User    *users.User
}

func (link Link) Save(ctx context.Context) (int64, error) {
	// client := prisma.PrismaClient()
	// createdLink, err := client.Links.CreateOne(
	// 	db.Links.Title.Set(link.Title),
	// 	db.Links.Address.Set(link.Address),
	// ).Exec(ctx)
	// if err != nil {
	// 	return 0, err
	// }
	// return createdLink.ID, nil

	return 0, nil

}

func GetAll(client *db.PrismaClient, ctx context.Context) ([]Link, error) {
	links, err := client.Links.FindMany().Exec(ctx)
	if err != nil {
		return nil, err
	}

	result, _ := json.MarshalIndent(links, "", "  ")
	fmt.Printf("links: %s\n", result)

	var linkList []Link
	for _, link := range links {
		linkList = append(linkList, Link{
			ID:      link.ID,
			Title:   link.Title,
			Address: link.Address,
		})
	}
	return linkList, nil
}
