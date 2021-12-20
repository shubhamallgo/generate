package converter

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/shubhamdixit863azarc/generate/pkg/inputs"
	"github.com/shubhamdixit863azarc/generate/pkg/utils"
)

func Convert(inputFiles []string, outputDir string) error {
	schemas, err := inputs.ReadInputFiles(inputFiles, true)
	if err != nil {
		return fmt.Errorf("error while reading input file %s", err)

	}
	generatorInstance := inputs.New(schemas...) // instance of generator which will produce structs
	err = generatorInstance.CreateTypes()
	if err != nil {
		return fmt.Errorf("error while generating instance for  proudcing structs %s", err)
	}

	for _, file := range inputFiles {
		var w io.Writer = os.Stdout
		packageDirectory, packageName := utils.PackageFormat(outputDir, file)
		err = os.Mkdir(packageDirectory, 0755)
		if err != nil {
			return fmt.Errorf("error while creating directory %s", err)

		}
		w, err = os.Create(filepath.Join(packageDirectory, filepath.Base(utils.FileNameCreation(file))))

		if err != nil {
			return fmt.Errorf("error while generating Files %s", err)
		}
		// Model Generation Method Called
		inputs.Output(w, generatorInstance, packageName)

	}
	return nil
}
