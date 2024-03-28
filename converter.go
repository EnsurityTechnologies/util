package util

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"math/big"
	"net"
	"time"
	"unicode/utf16"
)

func Assert(val bool, message string) {
	if !val {
		panic(message)
	}
}

func CheckEOF(conn *net.Conn) {
	_, err := (*conn).Read([]byte{})
	if err != nil {
		fmt.Println("Getting err from connection:", err)
	}
}

func Pad[T any](src []T, size int) []T {
	destination := make([]T, size)
	copy(destination, src)
	return destination
}

func ReadBE[T any](reader io.Reader) (T, error) {
	var value T
	err := binary.Read(reader, binary.BigEndian, &value)
	return value, err
}

func ReadLE[T any](reader io.Reader) (T, error) {
	var value T
	err := binary.Read(reader, binary.LittleEndian, &value)
	return value, err
}

func ToLE[T any](val T) []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, val)
	return buffer.Bytes()
}

func ToBE[T any](val T) []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.BigEndian, val)
	return buffer.Bytes()
}

func Write(writer io.Writer, data []byte) error {
	//fmt.Printf("\tWRITE: [%d]byte{%v}\n", len(data), data)
	_, err := writer.Write(data)
	return err
}

func Read(reader io.Reader, length uint) ([]byte, error) {
	output := make([]byte, length)
	_, err := reader.Read(output)
	return output, err
}

func Fill(buffer *bytes.Buffer, length int) {
	if buffer.Len() < length {
		zeroes := make([]byte, length-buffer.Len())
		Write(buffer, zeroes)
	}
}

func Utf16encode(message string) []byte {
	buffer := new(bytes.Buffer)
	for _, val := range utf16.Encode([]rune(message)) {
		binary.Write(buffer, binary.LittleEndian, val)
	}
	return buffer.Bytes()
}

func SizeOf[T any]() uint8 {
	var val T
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.BigEndian, &val)
	return uint8(buffer.Len())
}

func Flatten[T any](arrays [][]T) []T {
	output := make([]T, 0)
	for _, arr := range arrays {
		output = append(output, arr...)
	}
	return output
}

func StartRecurringFunction(f func(), interval int64) chan interface{} {
	stopSignal := make(chan interface{}, 1)
	trigger := make(chan interface{}, 1)
	wait := func() {
		time.Sleep(time.Millisecond * time.Duration(interval))
		trigger <- nil
	}
	go func() {
		for {
			go wait()
			select {
			case <-trigger:
				f()
			case <-stopSignal:
				return
			}
		}
	}()
	return stopSignal
}

func Delay(f func(), interval int64) {
	go func() {
		time.Sleep(time.Millisecond * time.Duration(interval))
		f()
	}()
}

func BytesToBigInt(b []byte) *big.Int {
	return big.NewInt(0).SetBytes(b)
}
