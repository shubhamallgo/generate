package main

import (
	"path/filepath"

	"github.com/shubhamdixit863azarc/generate/pkg/converter"
	"github.com/shubhamdixit863azarc/generate/pkg/utils"
)

func main() {
	inputDir, outputDir := utils.ParseFlags() // Parsing the cl flags
	files, err := utils.ReadFilesFromDir(inputDir)
	utils.CheckError(err)
	var inputFiles []string
	for _, file := range files {
		inputFiles = append(inputFiles, filepath.Join(utils.GetAbsolutePath(inputDir), file.Name()))
	}
	err = converter.Convert(inputFiles, outputDir)
	utils.CheckError(err)

}
