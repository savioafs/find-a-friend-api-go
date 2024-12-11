package entity

import (
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrOwnerIsRequired    = errors.New("owner is required")
	ErrEmailIsRequired    = errors.New("email is required")
	ErrCepIsRequired      = errors.New("cep is required")
	ErrCityIsRequired     = errors.New("city is required")
	ErrWhatsappIsRequired = errors.New("whatsapp is required")
	ErrPasswordIsRequired = errors.New("password is required")
)

type Organization struct {
	ID       string `db:"id"`
	Name     string `db:"name"`
	Owner    string `db:"owner"`
	Email    string `db:"email"`
	CEP      string `db:"cep"`
	City     string `db:"city"`
	Whatsapp string `db:"whatsapp"`
	Password string `db:"password"`
}

func NewOrganization(name, owner, email, cep, city, whatsapp, password string) (*Organization, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	org := &Organization{
		ID:       uuid.New().String(),
		Name:     name,
		Owner:    owner,
		Email:    email,
		CEP:      cep,
		City:     city,
		Whatsapp: whatsapp,
		Password: string(hashPassword),
	}

	err = org.Validate()
	if err != nil {
		return nil, err
	}

	return org, nil
}

func (o *Organization) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(o.Password), []byte(password))
	return err == nil
}

func (o *Organization) Validate() error {
	if o.Name == "" {
		return ErrNameIsRequired
	}
	if o.Owner == "" {
		return ErrOwnerIsRequired
	}

	if o.Email == "" {
		return ErrEmailIsRequired
	}

	if o.CEP == "" {
		return ErrCepIsRequired
	}

	if o.City == "" {
		return ErrCityIsRequired
	}

	if o.Whatsapp == "" {
		return ErrWhatsappIsRequired
	}

	if o.Password == "" {
		return ErrPasswordIsRequired
	}

	return nil
}
