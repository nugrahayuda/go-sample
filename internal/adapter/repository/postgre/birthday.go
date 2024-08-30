package repo

import (
	"errors"
	"time"
)

type RepoBirthdays interface {
	Create(birthday time.Time, userID uint32) (bool, error)
	Delete(userID uint32) (bool, error)
	GetBirthdayByID(userID uint32) (time.Time, error)
}

type repoBirthdays struct {
	// what should be the
	
}

func NewPhonebook() RepoBirthdays {
	return &repoBirthdays{}
}

func (r repoBirthdays) Create(birthday time.Time, userID uint32) (bool, error) {
	// Implementation
	return false, nil
}

func (r repoBirthdays) Delete(userID uint32) (bool, error) {
	// Implementation
	return false, nil
}

func (r repoBirthdays) GetBirthdayByID(userID uint32) (time.Time, error) {
	// Implementation
	if userID == 1 {
		return time.Time{}, errors.New("no birthday found")
	}
	return time.Time{}, nil
}

// create unit testing for this file

