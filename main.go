package main

import (
	"file-orgz/helper"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// categories := map[string]string{
	// 	".jpg":  "Images",
	// 	".png":  "Images",
	// 	".pdf":  "Documents",
	// 	".txt":  "TextFiles",
	// 	".docx": "Documents",
	// 	".doc":  "Documents",
	// 	".jpeg": "Images",
	// 	".iso":  "Installer",
	// 	".exe":  "Installer",
	// }

	crntPath, err := os.Getwd()
	if err != nil {
		fmt.Print("Error :\t", err)
	} else {
		fmt.Printf("Current Path: %s \n", crntPath)
	}

	fileLs, err := os.ReadDir(crntPath)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("\n Berhasil mengambil fileLs\n")
	}

	for _, file := range fileLs {
		if file.IsDir() || file.Name() == "organizer" || file.Name() == "organizer.exe" {
			continue
		}
		ext := filepath.Ext(file.Name())
		folder := helper.FindCategory(ext)

		targetFolder := filepath.Join(crntPath, folder)
		os.MkdirAll(targetFolder, os.ModePerm)

		oldPath := filepath.Join(crntPath, file.Name())

		newPath := filepath.Join(targetFolder, file.Name())
		err := os.Rename(oldPath, newPath)
		if err != nil {
			fmt.Printf("Error moving '%s' to '%s'. error %v\n", file.Name(), targetFolder, err)
		} else {
			fmt.Printf("Moved %s to %s \n", file.Name(), targetFolder)
		}
	}
}
