package core

import (
	"errors"
	"time"
)

type Item struct {
	data               any
	createdOn          time.Time
	expiredOn          time.Time
	expirationDuration time.Duration
}

func (i *Item) GetData() any {
	return i.data
}

func (i *Item) IsExpired(expirationTime time.Time) bool {
	if expirationTime.After(i.expiredOn) {
		return true
	}
	return false
}

func NewItemPtr(data any, expirationDuration time.Duration) (*Item, error) {
	if expirationDuration <= 0 {
		return nil, errors.New("expirationDuration should be > 0 ")
	}

	now := time.Now()

	return &Item{
			data:               data,
			expirationDuration: expirationDuration,
			createdOn:          now,
			expiredOn:          now.Add(expirationDuration),
		},
		nil
}

func NewItem(data any, expirationDuration time.Duration) (Item, error) {
	item, e := NewItemPtr(data, expirationDuration)
	if e != nil {
		return Item{}, e
	}
	return *item, e
}
