## FMT Package for Golang

The fmt package is a built-in package in Golang that provides functions for formatted I/O. It is used to print formatted text to the console, write formatted data to files, and scan formatted input from the console or files.

### Printing Formatted Text

The fmt package provides a variety of functions for printing formatted text to the console. The most common functions are:

* `Println(a ...interface{})`: Prints the arguments to the console, separated by spaces and terminated by a newline.
* `Printf(format string, a ...interface{})`: Prints the formatted text to the console, using the specified format string and arguments.
* `Fprintf(w io.Writer, format string, a ...interface{})`: Writes the formatted text to the specified writer, using the specified format string and arguments.

### Writing Formatted Data to Files

The fmt package also provides functions for writing formatted data to files. The most common function is:

* `Fprintln(w io.Writer, a ...interface{})`: Writes the formatted text to the specified writer, separated by spaces and terminated by a newline.

### Scanning Formatted Input

The fmt package also provides functions for scanning formatted input from the console or files. The most common functions are:

* `Scan(a ...interface{})`: Scans formatted input from the console and stores the values in the specified arguments.
* `Fscan(r io.Reader, a ...interface{})`: Scans formatted input from the specified reader and stores the values in the specified arguments.
* `Sscan(str string, a ...interface{})`: Scans formatted input from the specified string and stores the values in the specified arguments.

### Format Verbs

The fmt package uses format verbs to specify how to format the output of the printing and writing functions. The most common format verbs are:

* `%v`: Prints the value in its default format.
* `%s`: Prints the value as a string.
* `%d`: Prints the value as a decimal integer.
* `%f`: Prints the value as a floating-point number.
* `%t`: Prints the value as a boolean.

### Flags

The fmt package also supports flags that can be used to modify the format of the output. The most common flags are:

* `-`: Prints the value with a negative sign.
* `+`: Prints the value with a sign, even if it is positive.
* `#`: Prints the value with a prefix that indicates its type.
* `0`: Pads the value with zeros to the specified width.
* `-`: Pads the value with spaces to the specified width, and aligns it to the right.

### Examples

Here are some examples of how to use the fmt package:
```Go

// Print a formatted string to the console
fmt.Println("Hello, World!")

// Print a formatted string to a file
f, err := os.OpenFile("output.txt", os.O_CREATE|os.O_WRONLY, 0644)
if err != nil {
log.Fatal(err)
}
defer f.Close()

fmt.Fprintln(f, "Hello, World!")

// Scan a formatted input from the console
var name string
fmt.Scanln("What is your name?", &name)
fmt.Println("Hello,", name)

```

### References

For more info visit: https://pkg.go.dev/fmt
