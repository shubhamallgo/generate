package utils

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func FileNameCreation(fileName string) string {
	return fmt.Sprintf("%s %s", fileName[:len(fileName)-len(filepath.Ext(fileName))], ".go")
}

func SuffixFileExtension(fileName string) string {

	return strings.ToUpper(strings.TrimSuffix(fileName, filepath.Ext(fileName)))
}

//Gets the absolute Path for the files

func GetAbsolutePath(dir string) string {
	path, _ := filepath.Abs(dir) //Get the absolute path of the files
	return path
}

//Reads the files From Directories
func ReadFilesFromDir(dir string) ([]fs.DirEntry, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	return files, nil
}

// Formating the package name and the Directory Name

func PackageFormat(outputDir string, file os.DirEntry) (string, string) {
	packageDirectory := fmt.Sprintf("%s/%s", outputDir, Sanitizestring(file.Name()))
	packageName := Sanitizestring(file.Name())

	return packageDirectory, packageName
}

func ParseFlags() (string, string) {
	inputDir := flag.String("input", "./schemas", "Please Enter The Input Directory")
	outputDir := flag.String("output", "./schemas", "Please Enter The Input Directory")
	flag.Parse()
	return *inputDir, *outputDir
}

// Removes extra special characters from string and the

func Sanitizestring(fileName string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	return reg.ReplaceAllString(strings.TrimSuffix(fileName, filepath.Ext(fileName)), "")

}
