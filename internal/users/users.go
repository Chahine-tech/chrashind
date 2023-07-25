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

func (user *User) Create(ctx context.Context) error {
	client := prisma.PrismaClient()
	createdUser, err := client.Users.CreateOne(
		db.Users.Username.Set(user.Username),
		db.Users.Password.Set(user.Password),
	).Exec(ctx)
	if err != nil {
		return err
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
	return nil
}

func GetUserIdByUsername(username string, ctx context.Context) (int, error) {
	client := prisma.PrismaClient()
	user, err := client.Users.FindUnique(db.Users.ID.Equals(username)).Exec(ctx)
	if err != nil {
		return 0, err
	}

	userID, err := strconv.Atoi(user.ID)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func (user *User) Authenticate(ctx context.Context) bool {
	client := prisma.PrismaClient()
	foundUser, err := client.Users.FindFirst(db.Users.Username.Equals(user.Username)).Exec(ctx)
	if err != nil {
		// Handle the error appropriately, log or return false

		return false
	}

	// If no user is found, return false
	if foundUser == nil {
		return false
	}

	// Compare the hashed password with the provided password
	isPasswordMatch := bcrypt.CheckPasswordHash(user.Password, foundUser.Password)

	return isPasswordMatch
}
