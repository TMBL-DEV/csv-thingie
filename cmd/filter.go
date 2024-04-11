package cmd

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "filter",
	Short: "Filter through csv rows",
	Long:  "Filters through csv rows searching for what matches the given regex and creating a new file",

	Run: filterCsvFile,
}

func filterCsvFile(cmd *cobra.Command, args []string) {
	path, _ := cmd.Flags().GetString("file-path")
	grabHeader, _ := cmd.Flags().GetBool("grab-header")
	regexPattern, _ := cmd.Flags().GetString("regex-pattern")

	if strings.HasPrefix(path, "~/") {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Error getting current user home directory, try using an absolute path:", err)
			return
		}

		path = strings.Replace(path, "~", homeDir, 1)
	}

	outputPath := strings.Replace(path, ".csv", ".filtered.csv", 1)

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	regex, _ := regexp.Compile(regexPattern)
	reader := bufio.NewReader(file)

	// Create a CSV reader to parse each line when needed
	csvReader := csv.NewReader(reader)

	// Open a new file for writing the records
	outputFile, err := os.Create(outputPath)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}

	defer outputFile.Close()

	csvWriter := csv.NewWriter(outputFile)
	defer csvWriter.Flush()

	if grabHeader {
		fmt.Println("Setting up header row from old -> new file")
		firstRow, err := csvReader.Read()
		if err != nil {
			fmt.Println("Error reading record first row:", err)
			return
		}

		if err := csvWriter.Write(firstRow); err != nil {
			fmt.Println("error writing away header row")
		}
	}

	for {
		record, err := csvReader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			fmt.Println("Error reading record:", err)
			return
		}

		for _, slice := range record {
			if regex.MatchString(slice) {
				if err := csvWriter.Write(record); err != nil {
					fmt.Println("Error writing record to output file:", err)
					return
				}

				break
			}
		}
	}

	fmt.Println("Csv has been filtered see output in: ", outputPath)
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringP("file-path", "p", "empty", "path to csv file")
	generateCmd.MarkFlagRequired("file-path")

	generateCmd.Flags().StringP("regex-pattern", "r", "empty", "regex pattern which the row should match")
	generateCmd.MarkFlagRequired("regex-pattern")

	generateCmd.Flags().BoolP("grab-header", "g", true, "grabs first line as header for new file")
}
