package main

type Dto struct {
	Data      string     `json:"data"`
	NestedDto *NestedDto `json:"nested_dto"`
}

type NestedDto struct {
	NestedData int64 `json:"nested_data"`
}
