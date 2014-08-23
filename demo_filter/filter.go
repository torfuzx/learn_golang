/*!



*/
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type result struct {
	name   string
	size   int64
	suffix string
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// set the output flags for the standard logger, this will remove the formatting
	// as zero doesn't represent any predefined formatting
	log.SetFlags(0)

	algorithm, minSize, maxSize, suffixes, files := parseCommandLine()

	if algorithm == 1 {
		sink(filterSize(minSize, maxSize, filterSuffixes(suffixes, sources(files))))
	} else {
		// goroutine #1
		channel1 := sources(files)

		// goroutine #2
		channel2 := filterSuffixes(suffixes, channel1)

		// goroutine #3
		channel3 := filterSize(minSize, maxSize, channel2)

		sink(channel3)
	}
}

// Parse the command line, and adjust argument values.
func parseCommandLine() (algorithm int, minSize, maxSize int64, suffixes, files []string) {
	// defines a integer flag -algorithm, stored in the pointer algorithm
	flag.IntVar(&algorithm, "algorithm", 1, "1 [pipe style] or 2 [goroutine style]")
	// defines a int64 flag -min stored in the pointer minSize
	flag.Int64Var(&minSize, "min", -1, "minimum file size, default: -1 for no limit)")
	// defines a int64 flag -max stored in the pointer maxSize
	flag.Int64Var(&maxSize, "max", -1, "maximum file size, default: -1 for no limit)")
	// defines a string flag -suffixes stored in the pointer suffixesOpt
	var suffixesOpt *string = flag.String("suffixes", "", "comma separated list of file suffixes")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: %s -algorithm <int> -min <int64> -max <int64> -suffixes <string>\n", filepath.Base(os.Args[0]))
		flag.PrintDefaults()
		os.Exit(2)
	}

	// parse the command line into defined vars
	flag.Parse()
	if flag.NFlag() == 0 {
		flag.Usage()
	}

	if algorithm != 1 && algorithm != 2 {
		algorithm = 1
	}
	if minSize > maxSize && maxSize != -1 {
		log.Fatalln("minimum size must be less that maximum size")
	}

	suffixes = []string{}
	if *suffixesOpt != "" {
		suffixes = strings.Split(*suffixesOpt, ",")
	}

	// return non-flag command line arguments
	files = flag.Args()

	return algorithm, minSize, maxSize, suffixes, files
}

// Receives a slice of string and returns an receive-only channel for string.
// The contents of the files slice will be send to the channel sequentially.
// Channel #1: transfering all the original file names.
func sources(files []string) <-chan string {
	// create a buffered channel
	out := make(chan string, 1000)
	// invoking a temp anonymous goroutine
	go func() {
		for _, filename := range files {
			out <- filename
		}
		close(out)
	}()
	return out
}

// Receives the suffix slice as constraint and returns a string channel
// Channel #2: Tranfering all file names whose suffixes are not on blacklist.
func filterSuffixes(suffixes []string, in <-chan string) <-chan string {
	// create a channel with a pre-configured capacity, which makes the channel
	// works in a non-blocking asymchronous manner
	// make the buffer the same size as for files to maximize throughput, win time via space
	out := make(chan string, cap(in))

	// create a go routine by invoking a temp created anonymous function
	go func() {
		for filename := range in {
			if len(suffixes) == 0 {
				// send the file names to the channel, blocking manner, if none
				// suffix rule is set, then pass the suffix checking and
				// directly send the file
				out <- filename
				continue
			}

			// check the suffix
			ext := strings.ToLower(filepath.Ext(filename))
			for _, suffix := range suffixes {
				if ext == suffix {
					out <- filename // send the filename that meet the suffix constraint
					break
				}
			}
		}
		close(out)
	}()

	return out
}

// Receives the filter constraints and a chan string channel, and return its own channel
// Channel #3: transfering result after applying the size limits.
func filterSize(minimum, maximum int64, in <-chan string) <-chan string {
	// make a channel with a specified capacity, specify the capacity makes the
	// channel works in a asymchronous way, if there are used space in the buffer
	// that can be used for sending data, or still contains receivable data, them
	// the commumnication can work in the non-blocking way
	out := make(chan string, cap(in))

	// creat a goroutine by invoking a temporary anonymous function
	go func() {
		for filename := range in {
			if minimum == -1 && maximum == -1 {
				out <- filename // send method(blocking)
				continue
			}

			finfo, err := os.Stat(filename)
			if err != nil {
				continue
			}

			size := finfo.Size()
			if (minimum == -1 || minimum > -1 && minimum <= size) && (maximum == -1 || maximum > -1) {
				out <- filename // send method(blocking)
			}
		}
		close(out)
	}()
	return out
}

// Print the result
func sink(in <-chan string) {
	for filename := range in {
		fmt.Println(filename)
	}
}
