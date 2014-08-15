package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main () {
	// get the file names from the command line for input/output
	inFilename, outFilename, err := filenamesFromCommandLine()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	inFile, outFile := os.Stdin, os.Stdout
	if inFilename != "" {
		if inFile, err = os.Open(inFilename); err != nil {
			log.Fatal(err)
		}
		defer inFile.Close()
	}

	if outFilename != "" {
		if outFile, err = os.Create(outFilename); err != nil {
			log.Fatal(err)
		}
		defer outFile.Close()
	}

	if err = americanize(inFile, outFile); err != nil {
		log.Fatal(err)
	}
}

/**
 * Get the file names from the command line
 */
func filenamesFromCommandLine () (inFilename, outFilename string, err error) {
	if len(os.Args)	> 1 && (os.Args[1] == "-h" || os.Args[1] == "--help") {
		err = fmt.Errorf("usage: %s [<] infile.txt [>] outfile.txt", filepath.Base(os.Args[0]))
		return "", "", err
	}

	if len(os.Args) > 1 {
		inFilename = os.Args[1]
		if len(os.Args) > 2 {
			outFilename = os.Args[2]
		}
	}

	if inFilename != "" && inFilename == outFilename {
		log.Fatal("won't overwrite the infile")
	}

	return inFilename, outFilename, nil
}

var britshAmerican = "british-american.txt"

func americanize (inFile io.Reader, outFile io.Writer) (err error) {
	reader := bufio.NewReader(inFile)
	writer := bufio.NewWriter(outFile)

	defer func() {
		if err == nil {
			err = writer.Flush()
		}
	}()

	var replacer func(string) string
	if replacer, err = makeReplacerFunction(britshAmerican); err != nil {
		return err
	}

	wordRx := regexp.MustCompile("[A-Za-z]+")
	eof := false

	for !eof {
		var line string
		line, err = reader.ReadString('\n')

		if err == io.EOF  {
			err = nil
			eof = true
		} else if err != nil {
			return err
		}

		line = wordRx.ReplaceAllStringFunc(line, replacer)
		if _, err = writer.WriteString(line); err != nil {
			return err
		}
	}
	return nil
}

/**
 * Given a string and a replacer function, for every match replace the match
 * using the replacer
 */
func makeReplacerFunction (file string) (func(string) string, error) {
	rawBytes, err :=  ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	text := string(rawBytes)

	usForBritish :=  make(map[string]string)
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		fields  := strings.Fields(line)
		if len(fields) == 2 {
			usForBritish[fields[0]] = fields[1]
		}
	}

	return func (word string) string {
		if usWord, found := usForBritish[word]; found {
			return usWord
		}
		return word
	}, nil
}