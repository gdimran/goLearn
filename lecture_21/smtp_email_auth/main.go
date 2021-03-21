package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	resp := []byte("\x00" + "imrangd@mcwlbd.com" + "\x00" + "Ih377290499098")
	sEnc := base64.StdEncoding.EncodeToString([]byte(resp))
	fmt.Println(sEnc)
}
