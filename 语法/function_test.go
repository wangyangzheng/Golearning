package function_test

import "testing"

func TestFunction(t *testing.T) {
	return rand.Intn(10), rand.Intn(20)	
}