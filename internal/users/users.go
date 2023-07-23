package users

import (
	"context"
	"log"
	"strconv"

	"github.com/Chahine-tech/chrashind/prisma/db"
	"github.com/Chahine-tech/chrashind/utils/services/bcrypt"
	"github.com/Chahine-tech/chrashind/utils/services/prisma"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"name"`
	Password string `json:"password"`
}

func (user *User) Create(ctx context.Context) {
	client := prisma.PrismaClient()
	log.Print(client, "connect ?????")
	createdUser, err := client.Users.CreateOne(
		db.Users.Username.Set(user.Username),
		db.Users.Password.Set(user.Password),
	).Exec(ctx)
	log.Print(createdUser, "connect ?????")
	log.Print(err, "????")
	if err != nil {
		return
	}
	hashedPassword, err := bcrypt.HashPassword(user.Password)
	if err != nil {
		log.Fatal(err)
	}
	_, err = client.Users.FindUnique(
		db.Users.ID.Equals(createdUser.ID),
	).Update(
		db.Users.Password.Set(hashedPassword),
	).Exec(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func GetUserIdByUsername(username string, client *db.PrismaClient, ctx context.Context) (int, error) {
	user, err := client.Users.FindUnique(db.Users.ID.Equals(username)).Exec(ctx)
	if err != nil {
		// Handle the error appropriately
		return 0, err
	}

	userID, err := strconv.Atoi(user.ID)
	if err != nil {
		// Handle the conversion error appropriately
		return 0, err
	}

	return userID, nil
}

// func (user *User) Authenticate() bool {
// 	statement, err := database.Db.Prepare("select Password from Users WHERE Username = ?")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	row := statement.QueryRow(user.Username)

// 	var hashedPassword string
// 	err = row.Scan(&hashedPassword)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return false
// 		} else {
// 			log.Fatal(err)
// 		}
// 	}

// 	return bcrypt.CheckPasswordHash(user.Password, hashedPassword)
// }
