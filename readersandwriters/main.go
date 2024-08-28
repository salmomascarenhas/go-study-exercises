package main

// func main() {
// writer
// io.Writer is interface that abstracts the basic Write method that writes
// data to an underlying data stream.
// Write method takes a byte slice and returns the number of bytes written and
// an error.
//
// As an example, we can write a string to a file
//
// file, _ := os.Create("file.txt")
// defer file.Close()
// writer := bufio.NewWriter(file)
// writer.WriteString("Hello, World!")
// writer.Flush()
// or
// io.WriteString(file, "Hello, World!")
//

// reader
// io.Reader is interface that abstracts the basic Read method that reads
// data from an underlying data stream.
// Read method takes a byte slice and returns the number of bytes read and
// an error.
//
// As an example, we can read a file
//
// file, _ = os.Open("file.txt")
// reader := bufio.NewReader(file)
// content, _ := reader.ReadString('\n')
// fmt.Print(content)
// or
// var content []byte
// io.ReadAtLeast(reader, content, 12)
// fmt.Print(string(content))

// io.ReadAll reads all the data from an io.Reader until an error or EOF and
// returns the data it read.
//
// As an example, we can use io.ReadAll to read all the data from a file
//
// file, _ = os.Open("file.txt")
// contentBytes, _ := io.ReadAll(file)
// content := string(contentBytes)
// fmt.Print(string(content))
// }


