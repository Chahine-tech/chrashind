package links

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/Chahine-tech/chrashind/internal/users"
	"github.com/Chahine-tech/chrashind/prisma/db"
	"github.com/Chahine-tech/chrashind/utils/services/prisma"
)

type Link struct {
	ID      string
	Title   string
	Address string
	User    *users.User
}

func (link Link) Save() (int64, error) {
	linkID, err := prisma.Connect.Links.CreateOne(
		db.Links.Title.Set(link.Title),
		db.Links.Address.Set(link.Address),
		db.Links.UserID.Equals(link.User.ID),
	).Exec(context.Background())
	if err != nil {
		return 0, err
	}

	linkIDInt, err := strconv.ParseInt(linkID.ID, 10, 64)
	if err != nil {
		return 0, err
	}

	return linkIDInt, nil
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
