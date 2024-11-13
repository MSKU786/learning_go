package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args;

	if (len(args) < 2) {
		os.Exit(1);
	}
		
	fileName := args[1];
		
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1);
	}

	os.Stdout.Write(data)

	f, err := os.Open(fileName);

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1);
	}

	io.Copy(os.Stdout, f);
}


/*

Method: os.Open opens the file and returns a file handle f. Then, io.Copy copies data from the file handle f to os.Stdout. Unlike os.ReadFile, io.Copy reads and writes data in chunks rather than loading the entire file into memory at once.
Memory Usage: This approach is more memory-efficient, as it reads and writes data in chunks, making it ideal for large files. It only uses a small buffer for reading parts of the file at a time.
Error Handling: Errors are still caught when attempting to open the file, but io.Copy handles the reading and writing in chunks, which may provide more graceful error handling for large files.
Performance: This approach is better for large files since it does not require loading the entire file into memory. It may be marginally slower for very small files due to the chunked read/write, but this difference is often negligible.

*/

/*

Method: os.ReadFile reads the entire content of the file at once and loads it into memory. It then writes this data to os.Stdout using os.Stdout.Write(data).
Memory Usage: Since os.ReadFile reads the entire file content at once, it may use more memory, especially if the file is large. The entire file is stored in memory as a single slice of bytes.
Error Handling: Any error in reading the file, such as file not found or permission issues, is immediately caught.
Performance: This approach is fast for small files, but for large files, it can be less efficient due to high memory usage.
2. io.Copy(os.Stdout, f)

*/