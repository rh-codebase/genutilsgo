// functions to aid in unit testing
package genutilsgo

import (
	"fmt"
	"math"
	"testing"
)

// Helper to check int validity
func CheckI(t *testing.T, got, expected int, msg string) {
	if got != expected {
		fmt.Println(msg, " Expected ", expected, " Got ", got)
		t.Fail()
	}
}

// Helper to check float64 validity
func CheckF(t *testing.T, got, expected float64, msg string) {
	if math.IsNaN(got) || got != expected {
		fmt.Println(msg, " Expected ", expected, " Got ", got)
		t.Fail()
	}
}

// Helper to check float64 validity within a +/- tolerance
func CheckFT(t *testing.T, got, expected, tol float64, msg string) {
	if math.IsNaN(got) || math.Abs(got-expected) > tol {
		fmt.Println(msg, " Expected ", expected, " Got ", got, " Tolerance ", tol)
		t.Fail()
	}
}

// Helper to check string validity
func CheckS(t *testing.T, got, expected string, msg string) {
	if got != expected {
		fmt.Println(msg, " Expected ", expected, " Got ", got)
		t.Fail()
	}
}

// Helper to check expected Error
func CheckError(t *testing.T, got, expected error, msg string) {
	if got != expected {
		fmt.Println(msg, " Expected: ", expected, " Got: ", got)
		t.Fail()
	}
}

// Helper to check if error should not be nil
func CheckErrorNil(t *testing.T, got error, msg string) {
	if got == nil {
		fmt.Println(msg, " Expected <nil> Got ", got)
		t.Fail()
	}
}
