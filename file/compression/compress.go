package compression

import (
	"fmt"
    "compress/gzip"
    "os"
	"fileCompression/utils"
	"bytes"
    "encoding/binary"
	// "io/ioutil"
	// "log"
    // "bufio"
)

/*
* LZW compression algorithm implementation
* This is a placeholder function for LZW compression.
* You can implement the LZW compression logic here.
*/
func lzw_compression(uncompressed string) []int {
	dictSize := 256
	dictionary := make(map[string]int)

	// Initialize dictionary with single-byte characters
	for i := 0; i < 256; i++ {
		dictionary[string(rune(i))] = i
	}

	w := ""
	result := []int{}

	for _, c := range uncompressed {
		wc := w + string(c)
		if _, exists := dictionary[wc]; exists {
			w = wc
		} else {
			result = append(result, dictionary[w])
			// Add wc to the dictionary
			dictionary[wc] = dictSize
			dictSize++
			w = string(c)
		}
	}

	// Output the code for w
	if w != "" {
		result = append(result, dictionary[w])
	}

	return result
}


// Convert []int to []byte using uint16 encoding (little-endian)
func IntsToBytes(ints []int) []byte {
	buf := new(bytes.Buffer)
	for _, n := range ints {
		_ = binary.Write(buf, binary.LittleEndian, uint16(n))
	}
	return buf.Bytes()
}

// Write to .gz
func writeToGzip(data []byte, outputPath string) error {
	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := gzip.NewWriter(file)
	defer writer.Close()

	_, err = writer.Write(data)
	return err
}




func Compress(file_path string, output_path string) {
	

	fmt.Println("Compressing:", file_path)

	println("Opening file:", file_path)
	


	println("Please wait while the file is being compressed...")

	d, e := os.ReadFile(file_path)
	utils.ErrorChecker(e) // Check for errors when reading the file
	println("File read successfully, data length:", len(d))
	println("Data:", string(d))

	var lzwData []int = lzw_compression(string(d))
	println("LZW Compression done, data length:", len(lzwData))
	println("LZW Data:", lzwData)

	var byteData []byte = IntsToBytes(lzwData)
	println("Converted LZW data to bytes, length:", len(byteData))
	println("Byte Data:", byteData)

	writeToGzip(byteData, output_path)
	println("Gzip file created successfully")
	
	fmt.Println("Compression completed successfully!")
}