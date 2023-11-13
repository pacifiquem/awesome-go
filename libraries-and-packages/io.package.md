## IO Package for Golang

The io package provides basic interfaces to I/O primitives. It is used to read and write data from and to a variety of sources, including files, sockets, and pipes.

### Interfaces

The io package defines several interfaces that represent different types of I/O operations. The most common interfaces are:

* `io.Reader`: Represents a source of data that can be read from.
* `io.Writer`: Represents a destination of data that can be written to.
* `io.Seeker`: Represents a source or destination of data that can be seeked within.
* `io.Closer`: Represents an I/O resource that can be closed.

### Functions

The io package also provides a number of functions for working with I/O operations. The most common functions are:

* `Copy(dst io.Writer, src io.Reader)`: Copies data from the src Reader to the dst Writer.
* `ReadAt(r io.ReaderAt, p []byte, off int64)`: Reads len(p) bytes from the ReaderAt r at the specified offset off.
* `WriteAt(w io.WriterAt, p []byte, off int64)`: Writes len(p) bytes to the WriterAt w at the specified offset off.
* `Seek(r io.Seeker, offset int64, whence int)`: Seeks the specified offset within the Seeker r.

### Examples

Here are some examples of how to use the io package:

```Go

// Read the contents of a file
f, err := os.Open("myfile.txt")
if err != nil {
log.Fatal(err)
}
defer f.Close()

data, err := ioutil.ReadAll(f)
if err != nil {
log.Fatal(err)
}

fmt.Println(string(data))

// Write a string to a file
f, err := os.Create("myfile.txt")
if err != nil {
log.Fatal(err)
}
defer f.Close()

err = f.WriteString("Hello, World!")
if err != nil {
log.Fatal(err)
}

```

## More Information

For more information about the io package, please refer to the documentation: https://pkg.go.dev/io
