package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Unzip(src string, dest string) ([]string, error) {

	var filenames []string

	r, err := zip.OpenReader(src)
	if err != nil {
		return filenames, err
	}
	defer r.Close()

	for _, f := range r.File {

		// Store filename/path for returning and using later on
		fpath := filepath.Join(dest, f.Name)

		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return filenames, fmt.Errorf("%s: illegal file path", fpath)
		}

		filenames = append(filenames, fpath)

		if f.FileInfo().IsDir() {
			// Make Folder
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		// Make File
		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return filenames, err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return filenames, err
		}

		rc, err := f.Open()
		if err != nil {
			return filenames, err
		}

		_, err = io.Copy(outFile, rc)

		// Close the file without defer to close before next iteration of loop
		outFile.Close()
		rc.Close()

		if err != nil {
			return filenames, err
		}
	}
	return filenames, nil
}

func copyFile(f1 string, f2 string) {

	from, err := os.Open(f1)
	if err != nil {
		log.Fatal(err)
	}
	defer from.Close()

	to, err := os.OpenFile(f2, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer to.Close()

	_, err = io.Copy(to, from)
	if err != nil {
		log.Fatal(err)
	}
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func main() {

	/* Copy board files */
	if fileExists("FILE_NAME") {
		fmt.Println("FILE_NAME to FILE_LOCATION")
		copyFile("FILE_NAME", "C:/Program Files/FILE_LOCATION/FILE_NAME")
	} else {
		fmt.Println("FILE_NAME does not exist")
	}

	/* Copy AVI.zip for avi softare */
	if fileExists("ARCHIVE.zip") {
		fmt.Println("Transferring EXTRACTED to C:/Program Files/FILE_LOCATION")
		files, err := Unzip("ARCHIVE.zip", "C:/Program Files/FILE_LOCATION")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Unzipped:\n" + strings.Join(files, "\n"))
	} else {
		fmt.Println("ARCHIVE.zip does not exist")
	}

	/* Exit */
	fmt.Println("Press any key to exit")
	fmt.Scanf("test")
}
