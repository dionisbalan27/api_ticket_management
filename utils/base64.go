package utils

import "encoding/base64"

func Base64(data string, jenis_operasi string) string {
	var hasil string
	switch jenis_operasi {
	case "enc":
		hasil = base64.StdEncoding.EncodeToString([]byte(data))
	case "dec":
		h1, _ := base64.StdEncoding.DecodeString(data)
		hasil = string(h1)
	}
	return hasil

}
