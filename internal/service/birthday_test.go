package service

import (
	"reflect"
	"sort"
	"testing"
	"time"

	mock "integrationtests/internal/service/mock"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const (
	phone         = "+1 234567"
	userID uint32 = 1
)

func TestOurService_GetBirthdayByPhone(t *testing.T) {
	ctrl := gomock.NewController(t)

	birthdaysRepoMock := mock.NewMockbirthdaysRepo(ctrl)
	phonebookMock := mock.NewMockphonebook(ctrl)

	ourService := NewService(birthdaysRepoMock, phonebookMock)

	var birthday = time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)

	testCases := []struct {
		name     string
		mockFunc func()
		err      error
		birthday time.Time
	}{
		{
			"No phone in phonebook",
			func() {
				phonebookMock.EXPECT().GetUserIDByPhone(phone).Return(uint32(0), errNoPhone).Times(1)
			},
			errNoPhone,
			time.Time{},
		},
		{
			"No phone in phonebook",
			func() {
				phonebookMock.EXPECT().GetUserIDByPhone(phone).Return(userID, nil).Times(1)
				birthdaysRepoMock.EXPECT().GetBirthdayByID(userID).Return(time.Time{}, errNoBirthday).Times(1)
			},
			errNoBirthday,
			time.Time{},
		},
		{
			"Success",
			func() {
				phonebookMock.EXPECT().GetUserIDByPhone(phone).Return(userID, nil).MaxTimes(1)
				birthdaysRepoMock.EXPECT().GetBirthdayByID(userID).Return(birthday, nil).MaxTimes(1)
			},
			nil,
			birthday,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockFunc()
			bd, err := ourService.GetBirthdayByPhone(phone)
			assert.Equal(t, tc.birthday, bd)
			assert.Equal(t, tc.err, err)
		})
	}
}

func TestMergeSortInputSorted(t *testing.T) {
	tests := []struct {
		name string
		a    []string
		b    []string
		want []string
	}{
		{
			name: "Test Case 1: Empty slices",
			a:    []string{},
			b:    []string{},
			want: []string{},
		},
		{
			name: "Test Case 2: One empty slice",
			a:    []string{},
			b:    []string{"apple", "banana", "cherry"},
			want: []string{"apple", "banana", "cherry"},
		},
		{
			name: "Test Case 3: Strings with uppercase and lowercase letters",
			a:    []string{"apple", "Banana", "Cherry"},
			b:    []string{"grape", "kiwi", "Mango"},
			want: []string{"apple", "Banana", "Cherry", "grape", "kiwi", "Mango"},
		},
		{
			name: "Test Case 4: Strings with mixed order",
			a:    []string{"apple", "banana", "zebra"},
			b:    []string{"cherry", "kiwi", "mango"},
			want: []string{"apple", "banana", "cherry", "kiwi", "mango", "zebra"},
		},
		{
			name: "Test Case 5: Strings with duplicates",
			a:    []string{"apple", "banana", "banana"},
			b:    []string{"apple", "banana", "cherry"},
			want: []string{"apple", "apple", "banana", "banana", "banana", "cherry"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeSortWithInputSorted(tt.a, tt.b); !equalSlices(got, tt.want) {
				t.Errorf("mergeSort() = %v, want %v", got, tt.want)
			} else {
				t.Logf("mergeSort() = %v, want %v passed", got, tt.want)
			}
		})
	}
}

func equalSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestFactorial(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{name: "factorial of 0", n: 0, want: 1},
		{name: "factorial of 1", n: 1, want: 1},
		{name: "factorial of 2", n: 2, want: 2},
		{name: "factorial of 5", n: 5, want: 120},
		{name: "factorial of 10", n: 10, want: 3628800},
		// add more tests
		{name: "factorial of negative number", n: -1, want: 0},
		{name: "factorial of large number", n: 20, want: 2432902008176640000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := factorial(tt.n); got != tt.want {
				t.Errorf("factorial(%d) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}

func TestMerge(t *testing.T) {
	tests := []struct {
		name  string
		left  []int
		right []int
		want  []int
	}{
		{
			name:  "merge two sorted arrays",
			left:  []int{1, 3, 5, 7},
			right: []int{2, 4, 6, 8},
			want:  []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			name:  "merge two arrays with negative numbers",
			left:  []int{-5, -3, -1},
			right: []int{-4, -2, 0},
			want:  []int{-5, -4, -3, -2, -1, 0},
		},
		{
			name:  "merge two arrays with large numbers",
			left:  []int{1000000, 2000000, 3000000},
			right: []int{1500000, 2500000, 3500000},
			want:  []int{1000000, 1500000, 2000000, 2500000, 3000000, 3500000},
		},
		{
			name:  "merge two arrays with duplicate numbers",
			left:  []int{1, 2, 2, 3},
			right: []int{2, 3, 4, 4},
			want:  []int{1, 2, 2, 2, 3, 3, 4, 4},
		},
		{
			name:  "merge two empty arrays",
			left:  []int{},
			right: []int{},
			want:  []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := merge(tt.left, tt.right)
			sort.Ints(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func References(a *int, b int) (x, y int) {
	a = &b
    return *a, b
}

func TestReferences(t *testing.T) {
	var a int
    x, y := References(&a, 10)
	assert.Equal(t, 10, x)
    assert.Equal(t, 10, y)
}