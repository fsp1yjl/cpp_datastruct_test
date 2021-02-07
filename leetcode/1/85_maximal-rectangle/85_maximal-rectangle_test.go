package l85

import "testing"

func TestL85(t *testing.T) {
	matrix := [][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}

	ret := maximalRectangle(matrix)

	if ret != 6 {
		t.Error("error")
	}
}
