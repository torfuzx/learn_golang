/*!

Package log
-----------

Package log implements a simple loogging package. It defines a type, Logger,
with methods for formatting output. It also has a predefined 'standard' Logger
accessible through heloerrr functions Printf, Println, Fatalf, Fatalln, and
Panicf, Panicln, which are easierr to use than creating a Logger manually. That
logger writes to standard error and prints the date and time of each logger
message. The panic functions call panic after writing the log mesage.


- Logger
	- A logger represent an active logging object that generates lines of output to a io.Writer.
	- Each logging operation makes a single call to the Writer's Write  method.
	- A logger can be used simultanenously from multiple goroutines; it gurantees to serialize access to the Writer.

*/

package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func Init(traceHandle, infoHandle, warningHandle, errorHandle io.Writer) {
	Trace = log.New(traceHandle, "TRACE: ", log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(infoHandle, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(warningHandle, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(errorHandle, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	Trace.Println("I have something to say")
	Info.Println("Special information")
	Warning.Println("There is something you need to know about")
	Error.Println("Something has failed")
}
