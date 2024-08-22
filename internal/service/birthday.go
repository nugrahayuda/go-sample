package service

import (
	"errors"
	"math"
	"slices"
	"strings"
	"time"
)

//go:generate mockgen -source=service.go -destination=mock/service.go

var errNoPhone = errors.New("no phone")
var errNoBirthday = errors.New("no birthday")

type birthdaysRepo interface {
	GetBirthdayByID(userID uint32) (time.Time, error)
}

type phonebook interface {
	GetUserIDByPhone(phone string) (uint32, error)
}

type ourService struct {
	repo birthdaysRepo
	pb   phonebook
}

type OurService interface {
	GetBirthdayByPhone(phone string) (time.Time, error)
	GetBirthdayByID(userID uint32) (time.Time, error)
}

func NewService(repo birthdaysRepo, pb phonebook) OurService {
	return &ourService{
		repo: repo,
		pb:   pb,
	}
}

func (s *ourService) GetBirthdayByPhone(phone string) (time.Time, error) {
	userID, err := s.pb.GetUserIDByPhone(phone)
	if err != nil {
		return time.Time{}, errNoPhone
	}
	birthday, err := s.repo.GetBirthdayByID(userID)
	if err != nil {
		return time.Time{}, errNoBirthday
	}
	return birthday, nil
}

func (s *ourService) GetBirthdayByID(userID uint32) (time.Time, error) {
	birthday, err := s.repo.GetBirthdayByID(userID)

	if err != nil {
		return time.Time{}, err
	}
	return birthday, nil
}

func sortAlphabetical(a []string) []string {
	for i := range a {
		a[i] = strings.ToLower(a[i])
		for j := range a {
			a[i] = strings.ToLower(a[j])
			if a[i] < a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	return a
}
func mergeSortWithInputSorted(left, right []string) []string {
	result := make([]string, 0, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return sortAlphabetical(result)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {

		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	slices.Sort(result)

	return result
}

func factorial(n int) int {
	if n < 0 {
		return 0
	}
	if n == 0 || n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

// add test cases factorial

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
