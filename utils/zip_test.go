package utils

import (
	"fmt"
	"path"
	"strings"
	"testing"
)

func TestCompress(t *testing.T) {
	err := ZipUtil.CompressEncryption(path.Join("../storage", "Algorithm", "11"), path.Join("../storage", "bb.zip"), "")
	if err != nil {
		t.Error(err)
	}
}

func TestToTitle(t *testing.T) {
	str := strings.Title("HelloWorld")
	fmt.Printf("str: %v\n", str)
}
