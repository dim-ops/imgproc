package task

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Tasker interface {
	Process() error
}

type dirCtx struct {
	SrcDir string
	DstDir string
	files  []string
}

func buildFileList(srcDir string) []string { //Liste les img
	files := []string{} //create slice
	fmt.Println("Generate file list...")
	filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error { //explorer les dossiers
		if info.IsDir() || !strings.HasSuffix(path, ".jpg") { // on s'en fiche des fichiers et des fichiers ne se terminaux pas par .jpg -> on sort
			return nil
		}
		// si on arrive jusque la cad qu'on a une img
		files = append(files, path) //append chemin courrant, ajoute le chemin à la liste de fichier
		return nil
	})
	return files
}
