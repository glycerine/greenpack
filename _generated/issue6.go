package _generated

import (
	"time"
)

var _ = time.Time{}

//go:generate greenpack

// greenpack Issue 6: must be big enough to *NOT* inline.
type Event struct {
	ID          []byte    `zid:"0",msg:"id"`
	AggregateID []byte    `zid:"1",msg:"aggregate_id"`
	Type        string    `zid:"2",msg:"type"`
	Time        time.Time `zid:"3",msg:"time"`
}

type Commit struct {
	Events []*Event `zid:"0",msg:"events"`
}
