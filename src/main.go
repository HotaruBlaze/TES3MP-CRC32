package main

import (
	"bufio"
	"encoding/hex"
	"flag"
	"fmt"
	"hash/crc32"
	"io"
	"log"
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
	loadOrderFlag := flag.String("loadOrder", "", "Load order file")
	dataFileFlag := flag.String("dataFiles", "", "Data Files Directory")

	flag.Parse()

	flag.VisitAll(func(f *flag.Flag) {
		if f.Value.String() == "" {
			fmt.Println(os.Args[0] + " --loadOrder ./loadOrder.txt" + " --dataFiles  \"/home/tes3mp/Data Files\"")
			os.Exit(1)
		}
	})

	loadOrder, loadOrdererr := os.Open(*loadOrderFlag)

	if loadOrdererr != nil {
		log.Fatal(loadOrdererr)
	}
	defer loadOrder.Close()

	scanner := bufio.NewScanner(loadOrder)
	for scanner.Scan() {
		file := *dataFileFlag + scanner.Text()
		hash, err := hashFileCRC32(file, 0xedb88320)
		if err == nil {
			fmt.Println(filepath.Base(file) + ": " + "0x" + strings.ToUpper(hash))
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
