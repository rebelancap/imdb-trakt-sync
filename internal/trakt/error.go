package trakt

import (
	"fmt"
	"net/http"
)

type UnexpectedStatusCodeError struct {
	Got  int
	Want []int
}

func (e *UnexpectedStatusCodeError) Error() string {
	return fmt.Sprintf("unexpected status code: got %d, want one of %d", e.Got, e.Want)
}

func NewUnexpectedStatusCodeError(got int, want ...int) error {
	return &UnexpectedStatusCodeError{
		Got:  got,
		Want: want,
	}
}

type AccountLimitExceededError struct {
	accountLimit string
}

func (e *AccountLimitExceededError) Error() string {
	return fmt.Sprintf("trakt account limit (%s) exceeded, more info here: https://forums.trakt.tv/t/freemium-experience-more-features-for-all-with-usage-limits/41641", e.accountLimit)
}

func NewAccountLimitExceededError(headers http.Header) error {
	return &AccountLimitExceededError{
		accountLimit: headers.Get("X-Account-Limit"),
	}
}

type ListNotFoundError struct {
	Slug string
}

func (e *ListNotFoundError) Error() string {
	return fmt.Sprintf("list with slug %s could not be found", e.Slug)
}

func NewListNotFoundError(slug string) error {
	return &ListNotFoundError{
		Slug: slug,
	}
}
