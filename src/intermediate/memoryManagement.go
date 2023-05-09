/*
In Go, memory management is handled automatically by the garbage collector. The garbage collector is responsible for tracking memory usage and freeing memory that is no longer in use.

Go uses a concurrent garbage collector that runs in the background while the program is running. The garbage collector periodically scans the heap to identify objects that are no longer being used and frees up the memory they occupy.

Here are some key points to keep in mind about memory management in Go:

Go uses a heap to manage dynamically allocated memory.
Memory allocation in Go is fast and efficient because it uses a pool of pre-allocated memory blocks instead of requesting memory from the operating system every time an allocation is needed.
Go supports pointers and reference types, which means that you can allocate and manage memory manually if needed. However, this is not recommended in most cases because it can lead to memory leaks and other problems.
Go provides some tools to help diagnose memory-related issues, including the runtime.MemStats struct and the runtime.GC function.
Here's an example of using runtime.MemStats to gather information about memory usage in a Go program:
*/

package main

import (
    "fmt"
    "runtime"
)

func main() {
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    fmt.Printf("Alloc = %v MiB", m.Alloc/1024/1024)
}

/*
This program uses the runtime.MemStats struct to gather information about memory usage, and the runtime.ReadMemStats function to populate the struct.

It then prints out the amount of memory currently allocated in the heap, in megabytes.

Keep in mind that while the garbage collector automates memory management in Go, it's still important to write efficient code and avoid unnecessary memory allocations. In particular, be careful when working with large data structures, such as slices and maps, and avoid creating unnecessary copies of data.
*/