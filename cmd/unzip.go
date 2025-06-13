/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
    "fmt"
	"github.com/spf13/cobra"
	"fileCompression/file/unzip"
)




func parseUnzipFlags(cmd *cobra.Command) (string, string, error) {
	fileName, err := cmd.Flags().GetString("file")
	if err != nil {
		return "", "", err
	}

	outputName, err := cmd.Flags().GetString("output")
	if err != nil {
		return "", "", err
	}

	if outputName == "" {
		outputName = "output.txt" // Default output name if not provided
	}

	return fileName, outputName, nil
}

// unzipCmd represents the unzip command
var unzipCmd = &cobra.Command{
	Use:   "unzip",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("unzip called")

		file_path, output_path, err := parseUnzipFlags(cmd)

		if err != nil {
			fmt.Println("Error:", err)
			return // just exit the function early
		}
		fmt.Println("Unzipping:", file_path)
		if output_path != "" {
			fmt.Println("Output:", output_path)
		} else {
			output_path = "./data.txt"
		}

		unzip.Unzip(file_path, output_path)


	},
}

func init() {
	rootCmd.AddCommand(unzipCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// unzipCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// unzipCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	unzipCmd.Flags().StringP("file", "f", "", "File to unzip")
	unzipCmd.Flags().StringP("output", "o", "", "Output file name (default is 'output.txt')")
	unzipCmd.MarkFlagRequired("file") // Make the file flag required
}
