package converter

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/shubhamdixit863azarc/generate/pkg/inputs"
	"github.com/shubhamdixit863azarc/generate/pkg/utils"
)

func Convert() {
	dir, outputDir := utils.ParseFlags()

	files, err := utils.ReadFilesFromDir(dir)
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range files {

		var inputFiles []string
		inputFiles = append(inputFiles, filepath.Join(utils.GetAbsolutePath(dir), file.Name()))
		schemas, err := inputs.ReadInputFiles(inputFiles, true)
		if err != nil {
			fmt.Println(err)

		}
		g := inputs.New(schemas...)
		err = g.CreateTypes()
		if err != nil {
			fmt.Println(err)
		}
		var w io.Writer = os.Stdout
		packageDirectory, packageName := utils.PackageFormat(outputDir, file)
		err = os.Mkdir(packageDirectory, 0755)
		if err != nil {
			fmt.Println(err)
		}
		w, err = os.Create(filepath.Join(packageDirectory, filepath.Base(utils.FileNameCreation(file.Name()))))

		if err != nil {
			fmt.Println(err)
		}
		inputs.Output(w, g, packageName)

	}

}
