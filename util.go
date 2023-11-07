package util

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"log"
	"net"
)

// GetOutboundIP will get preferred outbound ip of this machine
func GetOutboundIP(site string) string {
	conn, err := net.Dial("udp", site+":80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	if localAddr == nil {
		return ""
	}
	if localAddr.IP == nil {
		return ""
	}
	return localAddr.IP.String()
}

// RandHexString will get hex encoded random string
func RandHexString(len int) string {
	b := make([]byte, len/2)
	rand.Read(b)
	return hex.EncodeToString(b)
}

// RandString will get the random base64 string
func RandString(len int) string {
	b := make([]byte, len)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}
