package unzip

import (
    "compress/gzip"
    "encoding/binary"
    "fmt"
    "io"
    "os"
    "fileCompression/utils"
)

// Decompress takes a slice of int codes and returns the decompressed string
func lzw_decompression(compressed []int) string {
	dictSize := 256
	// Initialize dictionary with single character strings for codes 0-255
	dictionary := make(map[int]string)
	for i := 0; i < dictSize; i++ {
		dictionary[i] = string(rune(i))
	}

	if len(compressed) == 0 {
		return ""
	}

	// Initialize w to first character of the first code
	w := dictionary[compressed[0]]
	result := w

	for _, k := range compressed[1:] {
		var entry string
		if val, ok := dictionary[k]; ok {
			entry = val
		} else if k == dictSize {
			// Special case: entry = w + first char of w
			entry = w + string(w[0])
		} else {
			panic(fmt.Sprintf("Bad compressed k: %d", k))
		}

		result += entry

		// Add w + first char of entry to the dictionary
		dictionary[dictSize] = w + string(entry[0])
		dictSize++

		w = entry
	}

	return result
}

func Unzip(file_path string, output_path string) {
    println("Opening compressed file:", file_path)
    fmt.Println("Please wait while the file is being unzipped...")

    file, err := os.Open(file_path)
    utils.ErrorChecker(err)
    defer file.Close()

    reader, err := gzip.NewReader(file)
    utils.ErrorChecker(err)
    defer reader.Close()

    // Read all decompressed bytes (which are the compressed LZW codes in bytes)
    compressedBytes, err := io.ReadAll(reader)
    utils.ErrorChecker(err)

    // Convert bytes to []int (uint16 little endian)
    var codes []int
    for i := 0; i < len(compressedBytes); i += 2 {
        if i+1 >= len(compressedBytes) {
            panic("incomplete code found in compressed data")
        }
        code := binary.LittleEndian.Uint16(compressedBytes[i : i+2])
        codes = append(codes, int(code))
    }

    // Decompress codes to original string
    decompressedString := lzw_decompression(codes)

    // Write decompressed string to output file
    writer, err := os.Create(output_path)
    utils.ErrorChecker(err)
    defer writer.Close()

    _, err = writer.Write([]byte(decompressedString))
    utils.ErrorChecker(err)

    fmt.Println("File decompressed successfully to:", output_path)
    fmt.Println("Unzip operation completed successfully.")
}


// package unzip

// import (
//     "fmt"
// 	"fileCompression/utils" // replace with your module path
// 	"os"
// 	"io"
// 	"compress/gzip"
// )


// func Unzip(file_path string, output_path string) {
// 	println("Opening compressed file:", file_path)
// 	fmt.Println("Please wait while the file is being unzipped...")

// 	file, err := os.Open(file_path)
// 	utils.ErrorChecker(err)
// 	defer file.Close()

// 	fmt.Println("File opened successfully:", file_path)
// 	reader, err := gzip.NewReader(file)
// 	if err != nil {
// 		fmt.Println("Error creating gzip reader:", err)
// 		return
// 	}
// 	defer reader.Close()

	
// 	fmt.Println("Creating output file:", output_path)
// 	writer, err := os.Create(output_path)

// 	utils.ErrorChecker(err)
// 	defer writer.Close()

// 	_, err = io.Copy(writer, reader)
// 	utils.ErrorChecker(err)

	
// 	fmt.Println("File unzipped successfully to:", output_path)
// 	fmt.Println("Unzip operation completed successfully.")


// }