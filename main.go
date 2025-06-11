package main

import "fileCompression/cmd"


func main() {
	cmd.Execute()
}





// func validateArguments() {



// }

// // checks for error
// func ErrorChecker(err error)  {
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// func main() {

// 	println("Welcome to the File Compression Project!")

// 	// // Open the file
// 	// file_name := "data.txt"
// 	// file_path := "C:/Users/layto/Coding/Golang/FileCompression/data/"
// 	// output_path := "C:/Users/layto/Coding/Golang/FileCompression/output/"

// 	inputFile := flag.String("file", "", "Input file to compress")
// 	// outputFile := flag.String("out", "", "Output file name (optional)")
// 	flag.Parse()

// 	if *inputFile == "" {
// 		fmt.Println("Usage: go run main.go -in=input.txt [-out=output.huff]")
// 		return
// 	}

// 	// fmt.Println("Compressing:", *inputFile)
// 	// if *outputFile != "" {
// 	// 	fmt.Println("Output will be saved to:", *outputFile)
// 	// }



// }