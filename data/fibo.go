package data

import (
	"encoding/json"
	"io"
)

type FiboInterval struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type FiboList []int

// ToJSON serializes the contents of the collection to JSON
// NewEncoder provides better performance than json.Unmarshal as it does not
// have to buffer the output into an in memory slice of bytes
// this reduces allocations and the overheads of the service
//
// https://golang.org/pkg/encoding/json/#NewEncoder
func (l *FiboList) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(l)
}

func (fi *FiboInterval) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(fi)
}

func CalculateSlice(fi *FiboInterval) FiboList {
	return Fibonacci(fi.X, fi.Y)
}
