import (
	"bufio"
	"fmt"
	"math"
	"os"
	"runtime"
)

type polar struct {
	radius float64
	θ      float64
}

type cartesian struct {
	x float64
	y float64
}

var prompt = "Enter a radius and an angle (in degrees), e.g., 12.5 90, " + "or %s to quit."

func init() {
	if runtime.GOOS == "windows" {
		prompt = fmt.Sprintf(prompt, "Ctrl + Z, Enter")
	} else { // Unix-like
		prompt = fmt.Sprintf(prompt, "Ctrl + D")
	}
}

func main() {
	// crate a channel used for transfering objects of polar struct
	questions := make(chan polar)
	// ensure the channel get destroyed/closed properly after use
	defer close(questions)

	// return a channel used for receiving messages
	answers := createSolver(questions)
	// ensure resource get properly destroyed
	defer close(answers)

	// pass in the two channels to where the user interaction takes place
	interact(questions, answers)
}

func createSolver(questions chan polar) chan cartesian {
	// create a channel, for sending the questions' answers - the cartesian coordinates
	answers := make(chan cartesian)
	go func() {
		for {
			// note: <- used as an unary operater, obtain a polar coordinate
			// from the 'questions' channel
			polarCoord := <-questions
			θ := polarCoord.θ * math.Pi / 180.0
			x := polarCoord.radius * math.Cos(θ)
			y := polarCoord.radius * math.Sin(θ)

			// note: <- used as a binary operator, left-hand operator being the
			// channel, right-hand operator being the object to send
			answers <- cartesian{x, y}
		}
	}()
	return answers
}

const result = "Polar radius=%.02f θ=%.02f -> Cartesian x=%.02f y=%.02f\n"

func interact(questions chan polar, answers chan cartesian) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf(prompt)

	for {
		fmt.Println("Radius and angle: ")
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		var radius, θ float64
		if _, err := fmt.Sscan(line, "%f %f", &radius, &θ); err != nil {
			fmt.Println(os.Stderr, "invalid input")
			continue
		}

		questions <- polar{radius, θ}
		coord := <-answers

		fmt.Printf(result, radius, θ, coord.x, coord.y)
	}

	fmt.Println()
}
