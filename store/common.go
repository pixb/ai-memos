package store

import "google.golang.org/protobuf/encoding/protojson"

var (
	protojsonUnmarshaler = protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: true,
	}
)

type RowStatus string

const (
	Normal   RowStatus = "NORMAL"
	Archived RowStatus = "ARCHIVED"
)

func (r RowStatus) String() string {
	return string(r)
}
