/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
    "bufio"
    "compress/gzip"
    "io/ioutil"
    "os"
	// "log"
	"github.com/spf13/cobra"
	"fileCompression/utils"
)


func parseCompressFlags(cmd *cobra.Command) (string, string, error) {
	fileName, err := cmd.Flags().GetString("file")
	if err != nil {
		return "", "", err
	}

	outputName, err := cmd.Flags().GetString("output")
	if err != nil {
		return "", "", err
	}

	if outputName == "" {
		outputName = "output.gz" // Default output name if not provided
	}

	return fileName, outputName, nil
}


// compressCmd represents the compress command
var compressCmd = &cobra.Command{
	Use:   "compress",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Compress called")

		file_path, output_path, err := parseCompressFlags(cmd)

		if err != nil {
			fmt.Println("Error:", err)
			return // just exit the function early
		}
		fmt.Println("Compressing:", file_path)
		if output_path != "" {
			fmt.Println("Output:", output_path)
		}

		println("Opening file:", file_path)

		file, err := os.Open(file_path) // Open the file for reading
		
		println(err)
		utils.ErrorChecker(err)

		println("Please wait while the file is being compressed...")


		read := bufio.NewReader(file)
		
		data, err := ioutil.ReadAll(read)

		// Close the file after reading
		err = file.Close()

		utils.ErrorChecker(err) // Check for errors when closing the file

		// Create the output file
		println("Creating output file:", output_path)
		file, err = os.Create(output_path) // Create the output file for writing
		utils.ErrorChecker(err)
		
		w := gzip.NewWriter(file) // Write to the gzip writer

		w.Write(data)

		// gives a notification when file compression is done
		fmt.Println("File compressed successfully")

		w.Close()
	},
}

func init() {
	rootCmd.AddCommand(compressCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// compressCmd.PersistentFlags().String("foo", "", "A help for foo")


	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// compressCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	compressCmd.Flags().StringP("file", "f", "", "Input file to compress")
	compressCmd.Flags().StringP("output", "o", "", "Output file name (optional)")

	compressCmd.MarkFlagRequired("file") // Mark the file flag as required
	compressCmd.MarkFlagFilename("file", "txt", "gz") // Ensure the file is a text file and output is gzipped
	compressCmd.MarkFlagFilename("output", "gz") // Ensure the output is a gzipped file

	
}
