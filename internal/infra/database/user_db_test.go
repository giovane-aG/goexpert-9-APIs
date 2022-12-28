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

func TestFindByEmail(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.User{})
	u := NewUser(db)
	newUser, err := entity.NewUser("Giovane", "findbyemail@email.com", "123456")

	u.DB.Create(newUser)

	foundUser, err := u.FindByEmail(newUser.Email)
	assert.Nil(t, err)
	assert.NotNil(t, foundUser)
	assert.Equal(t, foundUser.ID, newUser.ID)
	assert.Equal(t, foundUser.Name, newUser.Name)
	assert.Equal(t, foundUser.Email, newUser.Email)
	assert.Equal(t, foundUser.Password, newUser.Password)
}

func TestFindUserById(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.User{})
	u := NewUser(db)

	newUser, err := entity.NewUser("Giovane", "findbyid@email.com", "123456")

	u.DB.Create(newUser)

	foundUser, err := u.FindById(newUser.ID.String())
	assert.Nil(t, err)
	assert.NotNil(t, foundUser)
	assert.Equal(t, foundUser.ID, newUser.ID)
	assert.Equal(t, foundUser.Name, newUser.Name)
	assert.Equal(t, foundUser.Email, newUser.Email)
	assert.Equal(t, foundUser.Password, newUser.Password)
}

func TestUpdateUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.User{})
	u := NewUser(db)

	newUser, err := entity.NewUser("Giovane", "update@email.com", "123456")

	u.DB.Create(newUser)

	newUser.Name = "Giovane Updated"
	u.Update(newUser)

	var userUpdated *entity.User
	u.DB.Last(&userUpdated, "id = ?", newUser.ID.String())

	assert.Nil(t, err)
	assert.NotNil(t, userUpdated)
	assert.Equal(t, userUpdated.ID, newUser.ID)
	assert.Equal(t, userUpdated.Name, newUser.Name)
	assert.Equal(t, userUpdated.Email, newUser.Email)
	assert.Equal(t, userUpdated.Password, newUser.Password)
}
