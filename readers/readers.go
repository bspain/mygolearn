package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"golang.org/x/tour/reader"
)

func main() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)	// An 8-byte slice

	for {
		n, err := r.Read(b)	// n is number of bytes read, err is any error (will be io.OEF when done) and b contains data.

		fmt.Printf("n = %v err = %v b = %v\n", n, err, b) // 'n = 8 err = <nil> b = [72 101 108 108 111 44 32 82]'
		fmt.Printf("b[:n] = %q\n", b[:n]) // 'b[:n] = "Hello, R"' - b[:n] represents a slice of b n-long
		
		if err == io.EOF {
			break
		}

		/** Full output
		
		// Iter 1 - note that 8 bytes (indicated by n) was put into the slice 
		n = 8 err = <nil> b = [72 101 108 108 111 44 32 82]
		b[:n] = "Hello, R"

		// Iter 2 - only 6 bytes was put into the slice (note that '32'( ) and '82' (R) are still present in the slice from Iter 1 )
		n = 6 err = <nil> b = [101 97 100 101 114 33 32 82]
		b[:n] = "eader!"
		
		// Iter 3 - The slice is effectively unchanged from Iter 2, but n = 0, and EOF was returned.
		n = 0 err = EOF b = [101 97 100 101 114 33 32 82]
		b[:n] = ""

		**/
	}

	fmt.Println("\nExercise: Readers")
	{
		reader.Validate(MyReader{})
	}

	fmt.Println("\nExercise: rot13Reader")
	{
		s := strings.NewReader("Lbh penpxrq gur pbqr!\n")
		r := rot13Reader{s}  // rot13Reader struct has a single property, which is an io.Reader
		io.Copy(os.Stdout, &r) // Copy will call .Read on rot13Reader, which will read from it's internal reader and rot13 the contents
	}
}

// Exercise: Readers - Implement a Reader type that emits an infinite stream of the ASCII charecter 'A'

// MyReader is a struct
type MyReader struct{}

func (r MyReader) Read(b []byte) ( int, error ) {

	for i := range b {
		b[i] = 'A'
	}

	return len(b), nil
}

type rot13Reader struct {
	r io.Reader
}

func (rot13 *rot13Reader) Read(b []byte) ( int, error) {

	return Rot13Convert(rot13.r, b)
}

// Rot13Convert takes b length of data from the reader and rot13 encodes it.
func Rot13Convert(r io.Reader, b []byte) (n int, err error) {
	n, err = r.Read(b)

	if (err != nil) { return }

	for i := range b {
		c := b[i]

		if (c >= 'A' && c <= 'M') || ( c >= 'a' && c <= 'm') {
			b[i] += 13
		} else if (c >= 'N' && c <= 'Z') || (c >= 'n' && c <= 'z') {
			b[i] -=13
		}
	} 

	return
}