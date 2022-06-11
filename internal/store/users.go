package store

import (
	"crypto/rand"
	"time"

	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID             int
	Username       string `binding:"required,min=3,max=30"`
	Password       string `pg:"-" binding:"required,min=7,max=30"`
	HashedPassword []byte `json:"-"`
	Salt           []byte `json:"-"`
	Cookies        int
	LastClaimed    time.Time
	CreatedAt      time.Time
	ModifiedAt     time.Time
}

type UpdateCookies struct {
	UserID     int
}

func AddUser(user *User) error {
	salt, err := GenerateSalt()
	if err != nil {
		return err
	}
	toHash := append([]byte(user.Password), salt...)
	hashedPassword, err := bcrypt.GenerateFromPassword(toHash, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Salt = salt
	user.HashedPassword = hashedPassword
	user.Cookies = 0

	_, err = db.Model(user).Returning("*").Insert()
	if err != nil {
		log.Error().Err(err).Msg("error inserting new user")
		return dbError(err)
	}
	return nil
}

func Authenticate(username, password string) (*User, error) {
	user := new(User)
	if err := db.Model(user).Where("username = ?", username).Select(); err != nil {
		return nil, err
	}

	salted := append([]byte(password), user.Salt...)
	if err := bcrypt.CompareHashAndPassword(user.HashedPassword, salted); err != nil {
		return nil, err
	}
	return user, nil
}

func GenerateSalt() ([]byte, error) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		log.Error().Err(err).Msg("unable to create salt")
		return nil, err
	}
	return salt, nil
}

func FetchUser(id int) (*User, error) {
	user := new(User)
	user.ID = id
	err := db.Model(user).Returning("*").WherePK().Select()
	if err != nil {
		log.Error().Err(err).Msg("error fetching user")
		return nil, err
	}
	return user, nil
}

func AddCookies(id int, newCookies int) (*User, error) {
	user := new(User)
	user.Cookies = newCookies
	user.LastClaimed = time.Now()
	_, err := db.Model(user).Column("cookies").Where("ID = ?", id).Update()
	if err != nil {
		log.Error().Err(err).Msg("error adding cookies")
		return nil, err
	}

	_, err = db.Model(user).Column("last_claimed").Where("ID = ?", id).Update()
	if err != nil {
		log.Error().Err(err).Msg("error updating lastclaimed")
		return nil, err
	}
	return user, nil
}
