package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
)

type Result struct {
	filename string
	lino     int
	line     string
}

type Job struct {
	filename string
	results  chan<- Result
}

// this is where the hard work is taken, it's invoked by a doJob goroutine
// it will send the matched file name, line number, and the contents of the
// matched line to the Result channel. which is waited by main goroutine
func (job Job) Do(lineRx *regexp.Regexp) {
	file, err := os.Open(job.filename)
	if err != nil {
		log.Printf("error: %s\n", err)
		return
	}

	// defer invokes a function whose execution is deffered to the moment
	// the surrouding function returns
	defer file.Close()

	// the bufio.NewReader will return a new Reader with default buffer size
	reader := bufio.NewReader(file)
	for lino := 1; ; lino++ {
		line, err := reader.ReadBytes('\n')
		line = bytes.TrimRight(line, "\n\r")
		if lineRx.Match(line) {
			job.results <- Result{job.filename, lino, string(line)}
		}

		if err != nil {
			if err != io.EOF {
				log.Printf("error:%d: %s\n", lino, err)
			}
			break
		}
	}
}

var workers = runtime.NumCPU()

func main() {
	// use as many cores as possible
	runtime.GOMAXPROCS(runtime.NumCPU())

	// provide help message when asked by the user
	if len(os.Args) < 3 || os.Args[1] == "-h" || os.Args[1] == "-help" {
		fmt.Printf("usage: %s <regexp> <files>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	// the regexp.Compile() returns a *regexp.Regexp and nil if compiled ok
	if lineRx, err := regexp.Compile(os.Args[1]); err != nil {
		log.Fatal("invalid regexp: %s\n", err)
	} else {
		grep(lineRx, commandLineFiles(os.Args[2:]))
	}
}

func commandLineFiles(files []string) []string {
	if runtime.GOOS == "windows" {
		args := make([]string, 0, len(files))
		for _, name := range files {
			if matches, err := filepath.Glob(name); err != nil {
				args = append(args, name)
			} else if matches != nil {
				args = append(args, matches...) // an example of append one slice to another slice
			}
		}
		return args
	}
	return files
}

func grep(lineRx *regexp.Regexp, filenames []string) {
	jobs := make(chan Job, workers)
	results := make(chan Result, minimum(1000, len(filenames)))
	done := make(chan struct{}, workers)

	go addJobs(jobs, filenames, results)
	for i := 0; i < workers; i++ {
		go doJobs(done, lineRx, jobs)
	}
	waitAndProcessResults(done, results)
}

func minimum(x int, ys ...int) int {
	for _, y := range ys {
		if y < x {
			x = y
		}
	}
	return x
}

func addJobs(jobs chan<- Job, filenames []string, results chan<- Result) {
	for _, filename := range filenames {
		jobs <- Job{filename, results}
	}
	close(jobs)
}

func doJobs(done chan<- struct{}, lineRx *regexp.Regexp, jobs <-chan Job) {
	for job := range jobs {
		job.Do(lineRx)
	}
	done <- struct{}{}
}

func waitAndProcessResults(done <-chan struct{}, results <-chan Result) {
	for working := workers; working > 0; {
		select { // blocking
		case result := <-results:
			fmt.Printf("%s:%d:%s\n", result.filename, result.lino, result.line)
		case <-done:
			working--
		}
	}

DONE:
	for {
		// non-blocking
		select {
		// if there are unhandled results in the results channel, then it will be processed here,
		// the loop will continue until no more results in the results channel
		case result := <-results:
			fmt.Printf("%s:%d:%s\n", result.filename, result.lino, result.line)
		default:
			break DONE // this will break from both the select and and for statement. to the position of DONE label above
		}
	}
}
