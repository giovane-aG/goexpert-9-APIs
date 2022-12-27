package database

import (
	"testing"

	"github.com/giovane-aG/goexpert/9-APIs/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.User{})
	u := NewUser(db)
	newUser, err := entity.NewUser("Giovane", "giovane@email.com", "123456")

	u.DB.Create(newUser)

	var userCreated *entity.User
	u.DB.Last(&userCreated, "id = ?", newUser.ID.String())

	assert.NotNil(t, userCreated)
	assert.Equal(t, userCreated.ID, newUser.ID)
	assert.Equal(t, userCreated.Name, newUser.Name)
	assert.Equal(t, userCreated.Email, newUser.Email)
	assert.Equal(t, userCreated.Password, newUser.Password)
}
