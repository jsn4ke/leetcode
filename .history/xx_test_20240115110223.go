package leatcode

import (
	"fmt"
	"os"
	"testing"
)

func TestXxx(t *testing.T) {
	dir, err := os.Getwd()
	fmt.Println(dir, err)
}
