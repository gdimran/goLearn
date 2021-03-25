package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	resp := []byte("\x00" + "imrangd@gmail.com" + "\x00" + "Ih377290499098")
	fmt.Println(string(resp), resp)
	sEnc := base64.StdEncoding.EncodeToString([]byte(resp))
	fmt.Println(sEnc)
}
