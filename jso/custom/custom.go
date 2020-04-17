package custom

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
)

type Item struct {
	ID       string `json:"id"`
	HashCode string
}

// https://klotzandrew.com/blog/object-fingerprinting-for-efficient-data-ingestion

func (i *Item) UnmarshalJSON(data []byte) error {
	type aliasType *Item // to avoid recursive unmarshalling

	if err := json.Unmarshal(data, aliasType(i)); err != nil {
		return err
	}

	i.HashCode = fmt.Sprintf("%x", sha1.Sum(data))

	return nil
}
