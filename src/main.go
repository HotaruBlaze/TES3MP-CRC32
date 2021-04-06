package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"hash/crc32"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// https://mrwaggel.be/post/generate-crc32-hash-of-a-file-in-golang-turorial/
func hashFileCRC32(filePath string, polynomial uint32) (string, error) {
	//Initialize an empty return string now in case an error has to be returned
	var returnCRC32String string

	//Open the fhe file located at the given path and check for errors
	file, err := os.Open(filePath)
	if err != nil {
		return returnCRC32String, err
	}

	//Tell the program to close the file when the function returns
	defer file.Close()

	//Create the table with the given polynomial
	tablePolynomial := crc32.MakeTable(polynomial)

	//Open a new hash interface to write the file to
	hash := crc32.New(tablePolynomial)

	//Copy the file in the interface
	if _, err := io.Copy(hash, file); err != nil {
		return returnCRC32String, err
	}

	//Generate the hash
	hashInBytes := hash.Sum(nil)[:]

	//Encode the hash to a string
	returnCRC32String = hex.EncodeToString(hashInBytes)

	//Return the output
	return returnCRC32String, nil
}

func main() {
	var validFiles []string
	dirname := "." + string(filepath.Separator) + "Data Files" + string(filepath.Separator)

	d, err := os.Open(dirname)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer d.Close()

	files, err := d.Readdir(-1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, file := range files {
		if file.Mode().IsRegular() {
			if filepath.Ext(file.Name()) == ".esm" || filepath.Ext(file.Name()) == ".esp" {
				validFiles = append(validFiles, file.Name())
			}
		}
	}

	for _, validFile := range validFiles {
		validFilePath := "." + string(filepath.Separator) + "Data Files" + string(filepath.Separator) + validFile
		hash, err := hashFileCRC32(validFilePath, 0xedb88320)
		if err == nil {
			fmt.Println(filepath.Base(validFile) + ": " + "0x" + strings.ToUpper(hash))
		}
	}

	buf := bufio.NewReader(os.Stdin)
	fmt.Print("Press enter to close")
	_, _ = buf.ReadBytes('\n')
}
