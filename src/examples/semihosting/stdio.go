//go:build semihosting
// +build semihosting

package fmt

// These three file descriptors are connected to the host stdin/stdout/stderr,
// and can be used for logging.
var (
	Stdin  = File{fd: 0}
	Stdout = File{fd: 1}
	Stderr = File{fd: 2}
)

// File represents a semihosting file descriptor.
type File struct {
	fd uintptr
}

// Write writes the given data buffer to the file descriptor, returning an error
// if the write could not complete successfully.
func (f File) Write(buf []byte) (int, error) {
	return len(buf), Write(f.fd, buf)
}

// Print formats using the default formats for its operands and writes to standard output.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
func Print(a ...any) (n int, err error) {
	return Fprint(Stdout, a...)
}

// Println formats using the default formats for its operands and writes to standard output.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
func Println(a ...any) (n int, err error) {
	return Fprintln(Stdout, a...)
}

// Printf formats according to a format specifier and writes to standard output.
// It returns the number of bytes written and any write error encountered.
func Printf(format string, a ...any) (n int, err error) {
	return Fprintf(Stdout, format, a...)
}
