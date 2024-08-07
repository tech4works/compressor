package compressor

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
)

// TODO: Adicionar exemplo

// ToGzip converts any given value into a byte array using GZIP compression. It first
// converts the given value into a byte array using the toBytes function and then applies
// compression using gzip. If the conversion to byte array fails, it will cause a panic.
//
// Parameters:
//   - a: Any value to be converted into a GZIP compressed byte array.
//
// Returns:
//   - []byte: The GZIP compressed byte array representation of the given value.
//
// Panic:
//   - If the conversion to byte array fails, it will cause a panic.
//
// Example:
//
//	 var x string = "Hello, world!"
//	 var y int = 12345
//		fmt.Println(ToGzip(x))  // returns gzip compressed byte array of the string
//		fmt.Println(ToGzip(y))  // returns gzip compressed byte array of the integer
//
// Note:
//
//	This function is a variation of ToGzipWithErr where it suppresses any occurring error by causing a panic.
func ToGzip(a any) []byte {
	bs, err := ToGzipWithErr(a)
	if err != nil {
		panic(err)
	}
	return bs
}

// ToGzipWithErr converts any given value into a GZIP compressed byte array. It first
// converts the given value into a byte array using the toBytes function and then applies
// GZIP compression. Any errors that occur during the conversion or compression are
// returned alongside the byte array.
//
// Parameters:
//   - a: Any value to be converted into a GZIP compressed byte array.
//
// Returns:
//   - []byte: The GZIP compressed byte array representation of the given value.
//   - error: Any errors that might occur during the conversion or compression.
//
// Example:
//
//	var x string = "Hello, world!"
//	var y map[string]int = map[string]int{"one": 1, "two": 2}
//	bs, err := ToGzipWithErr(x)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(bs)  // Prints the gzip compressed byte array of the string
//
//	bs, err = ToGzipWithErr(y)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(bs)  // Prints the gzip compressed byte array of the map
func ToGzipWithErr(a any) ([]byte, error) {
	bs, err := toBytes(a)
	if err != nil {
		return nil, err
	}

	var gzipBuffer bytes.Buffer
	gz := gzip.NewWriter(&gzipBuffer)

	_, err = gz.Write(bs)
	if err != nil {
		return nil, err
	}

	err = gz.Close()
	if err != nil {
		return nil, err
	}

	return gzipBuffer.Bytes(), nil
}

// ToGzipBase64 converts any given value into a GZIP compressed Base64 string. It first
// compresses the given value into a GZIP byte array using the ToGzipBase64WithErr function
// and then converts it to a Base64 string. If any error occurs during the conversion
// or compression, it would panic and throw the error.
//
// Parameters:
//   - a: Any value to be converted into a GZIP compressed Base64 string.
//
// Returns:
//   - string: The GZIP compressed Base64 string representation of the given value.
//
// Panics:
//   - error: Any errors that might occur during the conversion or compression.
//
// Example:
//
//	var x string = "Hello, world!"
//	b64 := ToGzipBase64(x)
//	fmt.Println(b64)  // Prints the gzip compressed Base64 string of the string
//
//	var y map[string]int = map[string]int{"one": 1, "two": 2}
//	b64 = ToGzipBase64(y)
//	fmt.Println(b64)  // Prints the gzip compressed Base64 string of the map
func ToGzipBase64(a any) string {
	b64, err := ToGzipBase64WithErr(a)
	if err != nil {
		panic(err)
	}
	return b64
}

// ToGzipBase64WithErr converts any given value into a GZIP compressed Base64 string. It first compresses the given
// value into a GZIP byte array using the ToGzipWithErr function and then converts it into a Base64 string. Any errors
// that occur during the conversion or compression are returned alongside with the Base64 string.
//
// Parameters:
//   - a: Any value to be converted into a GZIP compressed Base64 string.
//
// Returns:
//   - string: The GZIP compressed Base64 string representation of the given value.
//   - error: Any errors that occur during the conversion or compression.
//
// Example:
//
//	var x string = "Hello, world!"
//	var y map[string]int = map[string]int{"one": 1, "two": 2}
//	b64, err := ToGzipBase64WithErr(x)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(b64)  // Prints the gzip compressed Base64 string of the string
//
//	b64, err = ToGzipBase64WithErr(y)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(b64)  // Prints the gzip compressed Base64 string of the map
func ToGzipBase64WithErr(a any) (string, error) {
	bs, err := ToGzipWithErr(a)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(bs), nil
}
