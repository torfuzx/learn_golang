/**
 *
 * An example to demonstrate cocurrent programming with go.
 *
 *
 */

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
)

var (
	pattern   string // the pattern to grep
	path      string // the path, either file or directory
	workerNum = runtime.NumCPU()
)

type Job struct {
	filename string
	results  chan<- Result
}

// Type used to represent the yield final type
type Result struct {
	filename string // the name of file that the pattern appears in
	lino     int    // the line number of the file
	line     string // the content of that file
}

func init() {
	parseCommandLine()
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Set a logger file
	file, err := os.OpenFile("cogrep.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error open log file: %v", err)
	}
	defer file.Close()

	log.SetOutput(file)
}

func main() {
	// buffered channel of type Job
	chanJob := make(chan Job, workerNum)
	// buffered channel of type Result
	chanResult := make(chan Result, 1000)
	// buffered channel of type struct{}, buffer size is the number of cpu cores
	chanDone := make(chan struct{}, workerNum)

	// use a worker goroutine for generating jobs
	// it acts as the producer in the producer-consumer modle
	// here a job represents a file's path, that will iterated under the path
	go produceJobs(path, chanJob, chanResult)

	// use separate worker goroutine to process the jobs
	// these series of goroutines act as the consumers in the producer-consumer
	// model
	for i := 0; i < workerNum; i++ {
		go consumeJobs(pattern, chanJob, chanDone)
	}

	// another goroutine to wait for completed jobs
	// here another goroutine for receiving the consumers' outcome, like a boss

	// this goroutine wait the completion of each doJobs goroutine,
	// the destination channel is put before the source channel
	// here the done channel is set to receive only
	go func(done <-chan struct{}, result chan Result) {
		for i := 0; i < workerNum; i++ {
			<-done
		}
		close(result)
	}(chanDone, chanResult)

	// process the results in the main goroutine
	for result := range chanResult {
		fmt.Printf("%s:%d:%s", result.filename, result.lino, result.line)
	}
}

func parseCommandLine() {
	flag.StringVar(&pattern, "e", "", "The search pattern")
	flag.StringVar(&path, "p", "", "The file to grep against")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s\n", filepath.Base(os.Args[0]))
		flag.PrintDefaults()
		os.Exit(2)
	}

	flag.Parse()

	if flag.NFlag() < 2 {
		flag.Usage()
	}

	if _, err := regexp.Compile(pattern); err != nil {
		log.Fatal("invalid pattern: %s\n", err)
	}
}

// this is where the hard work is taken, it's invoked by a doJob goroutine
// it will send the matched file name, line number, and the contents of the
// matched line to the Result channel. which is waited by main goroutine
func (job Job) Do(pattern string) {

	file, err := os.Open(job.filename)
	if err != nil {
		log.Printf("error: %s\n", err)
		return
	}

	// defer invokes a function whose execution is deffered to the moment
	// the surrouding function returns
	defer file.Close()

	regx, err := regexp.Compile(pattern)

	// the bufio.NewReader will return a new Reader with default buffer size
	reader := bufio.NewReader(file)
	for lino := 1; ; lino++ {
		line, err := reader.ReadBytes('\n')
		if regx.Match(line) {
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

// this goroutine will sending all the jobs to the jobs channel and close it
// after sending.
// - the jobs channel can be only to send Job
// - the result channel can be used only to send Result
func produceJobs(path string, chJob chan<- Job, chRes chan<- Result) {
	// iterate the directory, and sending each
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		log.Fatal(err)
	}

	switch mode := fi.Mode(); {
	case mode.IsDir():
		// iterate the directory
		filepath.Walk(path, func(p string, f os.FileInfo, err error) error {
			chJob <- Job{p, chRes}
			return nil
		})

	case mode.IsRegular():
		chJob <- Job{path, chRes}
	}

	// close channel after sending
	close(chJob)
}

// This goroutine will receive the jobs from the jobs channel and invoke the
// the job's do method to process method.
// this goroutine is blocking and after the it handles all the jobs dispatched
// this will send a succuess signal to the done channel.
// do job, will run in a separate go routine
func consumeJobs(pattern string, chJob <-chan Job, chDone chan<- struct{}) {
	for job := range chJob {
		job.Do(pattern)
	}
	chDone <- struct{}{}
}
