package main

// Import libraries
import (
    "bufio"
    "compress/gzip"
    "fmt"
    "io/ioutil"
    "os"
    "strings"
	"log"
)

// checks for error
func ErrorChecker(err error)  {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	println("Welcome to the File Compression Program!")

	// Open the file
	file_name := "data.txt"
	file_path := "C:/Users/layto/Coding/Golang/FileCompression/data/"
	output_path := "C:/Users/layto/Coding/Golang/FileCompression/output/"

	println("Opening file:", file_name)
	println("File path:", file_path)

	file, err := os.Open(file_path + file_name) 
	
	println(err)
	ErrorChecker(err)

	println("Please wait while the file is being compressed...")


	read := bufio.NewReader(file)
	
	data, err := ioutil.ReadAll(read)

	// Close the file after reading
	err = file.Close()

	ErrorChecker(err) // Check for errors when closing the file

	file_name = strings.Replace(file_name, ".txt", ".gz", -1)

	// Create the output file
	println("Creating output file:", file_name)
	file, err = os.Create(output_path + file_name)
	ErrorChecker(err)
	
	w := gzip.NewWriter(file) // Write to the gzip writer

	w.Write(data)

  	// gives a notification when file compression is done
	fmt.Println("File compressed successfully")

	w.Close()


}