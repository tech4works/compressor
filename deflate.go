package compressor

import (
	"bytes"
	"compress/flate"
	"encoding/base64"
)

// TODO: Adicionar exemplo

// ToDeflate compresses any given value into the Deflate format. This function internally
// uses the ToDeflateWithErr function to perform the compression and checks for any errors that may arise
// during the compression process.
//
// If an error is encountered during the compression, this function will panic and throw the error.
//
// Parameters:
//   - a: an interface value to be compressed into Deflate format.
//
// Returns:
//   - []byte: A byte slice containing Deflate compressed form of the input value.
//
// Panic:
//
//	This function will panic if an error is encountered during compression.
//
// Example:
//
//	str := "Hello World"
//	deflateBytes := ToDeflate(str)
//	//deflateBytes now contains the Deflate compressed form of "Hello World"
//
//	arr := []int{1, 2, 3}
//	deflateBytes = ToDeflate(arr)
//	// deflateBytes now contains the Deflate compressed form of the array [1,2,3]
func ToDeflate(a any) []byte {
	bs, err := ToDeflateWithErr(a)
	if err != nil {
		panic(err)
	}
	return bs
}

// ToDeflateWithErr function converts and compresses any given value into the Deflate format.
// This function delegates the conversion process to the 'toBytes' function.
// Also, it enables error propagation which gives the ability to handle the error in parent functions.
//
// Parameters:
//   - a: Any data type, provides the interface value to be compressed into Deflate format.
//
// Returns:
//   - [ ]byte: A byte slice exhibiting the Deflate compressed form of the value given as the 'a' parameter.
//   - error: An error object indicating any error encountered during the operation. Returns 'nil' if the operation is
//     successful.
//
// Example:
//
//	// Example 1: Simple String
//	s := "Hello, I am AI Assistant!"
//	bs, err := ToDeflateWithErr(s)
//	if err != nil {
//	     log.Fatal(err)
//	}
//	fmt.Println(bs) // Prints the Deflate compressed form of the string
//
//	// Example 2: Struct
//	type user struct {
//	    Name string
//	    Age  int
//	}
//	u := user{"Bob", 25}
//	bs, err = ToDeflateWithErr(u)
//	if err != nil {
//	     log.Fatal(err)
//	}
//	fmt.Println(bs) // Prints the Deflate compressed form of the struct
func ToDeflateWithErr(a any) ([]byte, error) {
	bs, err := toBytes(a)
	if err != nil {
		return nil, err
	}

	var buffer bytes.Buffer
	writer, err := flate.NewWriter(&buffer, 9)
	if err != nil {
		return nil, err
	}

	_, err = writer.Write(bs)
	if err != nil {
		return nil, err
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

// ToDeflateBase64 compresses a given value using Deflate method and then encodes the compressed data into a Base64 string.
// It uses the ToDeflateBase64WithErr function to perform the compression and encoding.
// If any error occurs during the process, the function will panic and terminate the program.
//
// Parameters:
//   - a: An interface value. This can be any data type and holds the data to be compressed and encoded.
//
// Returns:
//   - string: A string in Base64 format that represents the Deflate compressed form of the input data.
//
// Example:
//
//	// Example 1: Simple String
//	s := "Hello, I am AI Assistant!"
//	b64 := ToDeflateBase64(s)
//	fmt.Println(b64) // Prints the Base64-encoded, Deflate compressed form of the string
//
//	// Example 2: Struct
//	type user struct {
//	    Name string
//	    Age  int
//	}
//	u := user{"Bob", 25}
//	b64 = ToDeflateBase64(u)
//	fmt.Println(b64) // Prints the Base64-encoded, Deflate compressed form of the struct
//
// Panics:
//   - If there is an error during the compression or encoding process, the function will panic.
func ToDeflateBase64(a any) string {
	b64, err := ToDeflateBase64WithErr(a)
	if err != nil {
		panic(err)
	}
	return b64
}

// ToDeflateBase64WithErr compresses any given value into the Deflate format and encodes the result into a Base64 string.
// This function leverages the ToDeflateWithErr function for the compression and then uses the
// base64.StdEncoding.EncodeToString method for the encoding of the compressed data into Base64. Consequently,
// it can propagate any errors encountered in the compression or encoding.
//
// Parameters:
//   - a: Any value that will be compressed into the Deflate format and encoded into Base64. The actual value can be
//     of any data type.
//
// Returns:
//   - string: A Base64-encoded string that represents the Deflate-compressed form of the input data.
//   - error: An error object that indicates any errors that occurred during the compression or encoding operations.
//     If no errors occurred, it returns nil.
//
// Example:
//
//	// Example 1: Simple String
//	s := "Hello, I am AI Assistant!"
//	b64, err := ToDeflateBase64WithErr(s)
//	if err != nil {
//	     log.Fatal(err)
//	}
//	fmt.Println(b64) // Prints the Base64-encoded, Deflate compressed form of the string
//
//	// Example 2: Struct
//	type user struct {
//	    Name string
//	    Age  int
//	}
//	u := user{"Bob", 25}
//	b64, err = ToDeflateBase64WithErr(u)
//	if err != nil {
//	     log.Fatal(err)
//	}
//	fmt.Println(b64) // Prints the Base64-encoded, Deflate compressed form of the struct
func ToDeflateBase64WithErr(a any) (string, error) {
	bs, err := ToDeflateWithErr(a)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(bs), nil
}
