package setup

import (
	"github.com/jjisolo/lastdisco-backend/config"
	"github.com/jjisolo/lastdisco-backend/internal/storage"
	"github.com/jjisolo/lastdisco-backend/internal/types"
	"log"
)

// SeedDatabase seeds some basic database entries such as super frontend user, initial category, etc.
func SeedDatabase(s storage.Storage) {
	createInitialUsers(s)
	createInitialCategories(s)
}

// createInitialUsers creates the initial users, such as for instance super frontend user, first moderator, etc.
func createInitialUsers(s storage.Storage) {
	user, err := types.NewUser(
		config.SUPER_ADMIN_FIRST_NAME,
		config.SUPER_ADMIN_LAST_NAME,
		config.SUPER_ADMIN_PHONE_NUMBER,
		config.SUPER_ADMIN_EMAIL_ADDRESS,
		config.SUPER_ADMIN_PASSWORD,
	)
	if err != nil {
		log.Fatal(err)
		return
	}

	user.Role = "ADMIN"

	err = s.CreateUser(user)
	if err != nil {
		log.Fatal(err)
		return
	}
}

// createInitialCategories creates the main categories, such as for instance 'uncategorized'
func createInitialCategories(s storage.Storage) {
	category, err := types.NewProductCategory(
		"uncategorized",
		"This category is for the products does not belong to any category",
	)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = s.CreateCategory(category)
	if err != nil {
		log.Fatal(err)
		return
	}
}
