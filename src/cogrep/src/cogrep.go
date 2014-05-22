/**
 * An example to demonstrate cocurrent programming with go.
 */

package main

import (
	"fmt"
	"runtime"
	"os"
	"path/filepath"
	"regexp"
	"log"
	"bufio"
	"io"
)

type Job struct {
	filename string
	results chan<- Result
}

func (job Job) Do(lineRx *regexp.Regexp) {
	file, err := os.Open(job.filename)
	if err != nil {
		log.Printf("error: %s\n", err)
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	for lino := 1; ;lino ++ {
		line, err := reader.ReadBytes('\n')
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


type Result struct {
	filename 	string
	lino	 	int
	line		string
}

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
		for _,name := range (files) {
			if matches, err := filepath.Glob(name); err != nil {
				args = append(args, name)
			} else if  matches != nil {
				args = append(args, matches...)	// an example of append one slice to another slice
			}
		}
		return args
	}
	return files
}

var workers = runtime.NumCPU()

func grep(lineRx *regexp.Regexp, filenames []string) {
	// the lineRx will be shared among the worker goroutines, although we
	// should always assume that the vlaue pointed by a shared pointer is not
	// thread-safe, the go officdoc says it's thread safe, so we could safely use
	// this pointer among all the goroitines
	jobs 	:= make(chan Job, workers)							// asym channel of type Job, buffer size is the number of cpu cores
	results := make(chan Result, minimum(1000, len(filenames)))	// asym channel of type Result
	done 	:= make(chan struct {}, workers)					// asym channel of type struct{}, buffer size is the number of cpu cores

	// use a worker goroutine for scheduling/enqueing the jobs
	go addJobs(jobs, filenames, results)

	// use separate worker goroutines to process the jobs
	for i := 0; i < workers; i++ {
		go doJobs(done, lineRx, jobs)
	}

	// another goroutine to wait for complted jobs
	go awaitCompletion(done, results)

	// process the results in the main go routine
	processResults(results)
}

// the jobs channel can be only to send Job
// the result channel can be used only to send Result

// Note: the <- operator specifies the channel direction:
// 1. chan<- TYPE means can be only used to send objects of type TYPE
// 2. <-chan TYPE means can be only used to recv objects of type TYPE
func addJobs(jobs chan<- Job, filenames []string, results chan<- Result) {
	for _, filename := range filenames {
		jobs <- Job{filename, results}
	}
	// after the jobs channel sending all the jobs(file), close the channel
	close(jobs)
}

// do job, will run in a separate go routine
func doJobs(done chan <- struct{}, lineRx *regexp.Regexp, jobs chan Job) {
	for job := range jobs {
		job.Do(lineRx)
	}
	done <- struct{}{}// this is an send statement, the done channel is a
					  // send-only channel, so we can send things into it
}

// the destination channel is put before the source channel
// here the done channel is set to receive only
func awaitCompletion(done <-chan struct{}, result chan Result) {
	for i := 0; i < workers; i++ {
		<-done
	}
	close(result)
}

func processResults(results <-chan Result) {
	for result := range results {
		fmt.Println("%s:%d:%s\n", result.filename, result.lino, result.line)
	}
}

func minimum(x int, ys ...int) int {
	for _, y := range ys {
		if y < x {
			x = y
		}
	}
	return x
}


