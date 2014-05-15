package main

import (
	"fmt"
	"math/big"
	"path/filepath"
	"strconv"
	"os"
)

func main () {
	places := handleCommandLine(1000)
	scaledPi := fmt.Sprint(π(places))
	fmt.Printf("3.%s\n", scaledPi[1:])
}

/**
 * @see: http://en.literateprograms.org/Pi_with_Machin's_formula_(Python)
 */
func π(places int) *big.Int {
	digits   := big.NewInt(int64(places))
	unity    := big.NewInt(0)
	ten      := big.NewInt(10)
	exponent := big.NewInt(0)
	pi       := big.NewInt(4)

	unity.Exp(ten, exponent.Add(digits, ten), nil)		// unity = 10**(digits + 10)

	left := arccot(big.NewInt(5), unity)				// pi = 4 * (4*arccot(5, unity) - arccot(239, unity))
	left.Mul(left, big.NewInt(4))
	right := arccot(big.NewInt(239), unity)
	left.Sub(left, right)
	pi.Mul(pi, left)
	return pi.Div(pi, big.NewInt(0).Exp(ten, ten, nil))	 // pi / 10**10
}

func handleCommandLine(defaultValue int) int {
	if len(os.Args) > 1 {
		if os.Args[1] == "-h" || os.Args[1] == "--help" {
			usage := "usage: %s [digits]\n e.g.: %s 1000"
			app := filepath.Base(os.Args[0])
			fmt.Fprintln(os.Stderr, fmt.Sprintf(usage, app, app))
			os.Exit(1)
		}

		if x, err := strconv.Atoi(os.Args[1]); err != nil {
			fmt.Fprintf(os.Stderr, "ignoring invalid number of " +
				"digits: will display %d\n", defaultValue)
		} else {
			return x
		}
	}
	return defaultValue
}

/**
 * Gives the arc cotangent cot⁻¹(z) of the complex number z.
 * Note: cot(x) = cos(x) / sin(x)
 */
func arccot(x, unity *big.Int) *big.Int {
	sum := big.NewInt(0)		// allocates and returns a new int set to sum
	sum.Div(unity, x)			// set sum as the quotient unity / x
	xpower := big.NewInt(0)		// same as above
	xpower.Div(unity, x)

	n       := big.NewInt(3)
	sign    := big.NewInt(-1)
	zero    := big.NewInt(0)
	xsquare := big.NewInt(0)
	xsquare.Mul(x, x)

	for {
		xpower.Div(xpower, xsquare)		// xpower = xpower / (x*x)

		term := big.NewInt(0)			// term = xpower/n
		term.Div(xpower, n)

		if term.Cmp(zero) == 0 {
			break
		}

		added := big.NewInt(0)				// added = sign * term
		sum.Add(sum, added.Mul(sign, term))	// sum += added
		sum.Neg(sign)						// sign = -sign
		n.Add(n, big.NewInt(2))
	}
	return sum
}
