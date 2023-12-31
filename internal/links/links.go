package links

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Chahine-tech/chrashind/internal/users"
	"github.com/Chahine-tech/chrashind/utils/services/prisma"
)

type Link struct {
	ID      string
	Title   string
	Address string
	User    *users.User
}

func (link Link) Save(ctx context.Context) (int64, error) {
	// user := ctx.Value(middleware.RequestCtxKey("user")).(*users.User)
	// if user == nil {
	// 	return 0, errors.New("user not found")
	// }
	// client := prisma.PrismaClient()
	// createdLink, err := client.Links.CreateOne(
	// 	db.Links.Title.Set(link.Title),
	// 	db.Links.Address.Set(link.Address),
	// 	db.Links.User.Link(
	// 		db.Users.ID.Equals(user.ID),
	// 	),
	// 	db.Links.ID.Set("user"),
	// ).Exec(ctx)
	// if err != nil {
	// 	return 0, err
	// }
	// linkID, err := strconv.ParseInt(createdLink.ID, 10, 64)
	// if err != nil {
	// 	return 0, err
	// }
	// return linkID, nil
	return 0, nil
}

func GetAll(ctx context.Context) ([]Link, error) {
	client := prisma.PrismaClient()
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
