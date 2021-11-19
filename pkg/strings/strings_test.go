package strings

import (
	"fmt"
	"testing"
)

func TestParseFileName(t *testing.T) {
	filename := "8:30/开播:*?.mp4"
	nf, err := ParseFileName(filename)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(nf)
}
