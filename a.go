package main

import (
	"net/http"
)

type Example struct {
	Int    int
	Uint   uint
	String string
	Array  []int
	Map    map[int]int
	Object http.Client
	Status Status `accessor:",type:enum"`

	PInt    *int
	PUint   *uint
	PString *string
	PArray  *[]int
	PMap    *map[int]int
	PObject *http.Client
	PStatus *Status `accessor:",type:*enum"`
}

type Status int32
