package tests

import (
	"testing"

	"github.com/rabdavinci/fibo/data"
)

func TestBig(t *testing.T) {
	var fiboInterval data.FiboInterval
	fiboInterval.X = 1
	fiboInterval.Y = 99
	fiboList := data.CalculateSlice(&fiboInterval)
	// real result for x=1, y=99
	rr := "1344719667586153181419716641724567886890850696275767987106294472017884974410332069524504824747437757"
	if fiboList != rr {
		t.Errorf("Expected %s, but got %s", rr, fiboList)
	}

}
