package main

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {

	input := "cxdnnyjw"
	h := md5.New()
	count := 0
	password := []byte("........")
	used := make([]bool, 8)

	for i := 0; count < 8; i++ {
		h.Reset()
		io.WriteString(h, input)
		io.WriteString(h, strconv.Itoa(i))

		chksum := h.Sum(nil)

		hcheck := hex.EncodeToString(chksum)

		if strings.HasPrefix(hcheck, "00000") {
			position := hcheck[5] - '0'
			if position > 7 {
				continue
			}
			if used[position] {
				continue
			}
			used[position] = true
			password[position] = hcheck[6]
			count++
			io.WriteString(os.Stdout, "\b\b\b\b\b\b\b\b")
			io.WriteString(os.Stdout, string(password))
		}
	}

	io.WriteString(os.Stdout, "\n")
}
