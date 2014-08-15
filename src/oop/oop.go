/**
 * Point:
 * Method set: the set of all methods that can be called on the value of that type
 * The mothod set on the pointer to a  value of a custom type contains all the methods defined for that type, no matter whether they accept a value or a pointer
 * The method set on a coustom defined type contains all the methods defined for that type that accept a value receiver, but not the ones that receive a pointer receiver
 */

package main

import (
	"image/color"
	"fmt"
	"strings"
	"unicode"
	"io"
)

// -----------------------------------------------------------------------------
type ColoredPoint struct {
	color.Color		// anoymous field (embeded)
	x, y int		// named field (aggregated)
}

// -----------------------------------------------------------------------------
type RuneForRuneFunc func (rune) rune

// -----------------------------------------------------------------------------
type Count int
func (count *Count) Increment() {*count++}
func (count *Count) Decrement() {*count--}
func (count Count) IsZero() bool {return count == 0}


// -----------------------------------------------------------------------------
type Part struct {
	Id 		int
	Name	string
}

func (part *Part) LowerCase() {
	part.Name = strings.ToLower(part.Name)
}

func (part *Part) UpperCase() {
	part.Name = strings.ToUpper(part.Name)
}

func (part Part) String() string {
	return fmt.Sprintf("<<%d %q>>", part.Id, part.Name)
}

func (part Part) HasPrefix(prefix string) bool{
	return strings.HasPrefix(part.Name, prefix)
}

// -----------------------------------------------------------------------------
type Item struct {
	id 			string
	price 		float64
	quantity 	int
}

func (item *Item) Cost() float64{
	return item.price * float64(item.quantity)
}

type SpeicialItem struct {
	Item		 		// anoymous (embeded)
	catalogId	int		// named field (aggregation)
}

type LuxuryItem struct {
	Item
	markup float64
}

func (item *LuxuryItem) Cost() float64 {
	// return item.Item.price * float64(item.Item.quantity) * item.markup
	// return item.price * float64(item.quantity) * item.markup
	return item.Item.Cost() * item.markup
}

// -----------------------------------------------------------------------------

type Place struct {
	latitude, longitude float64
	Name				string
}

// constuctor, to be called explicitly
func New(latitude, longitude float64, name string) *Place {
	return &Place{saneAngle(0, latitude), saneAngle(0, longitude), name}
}

// getter of 'latitude' property, return the latitude
func (place *Place) Latitude() float64 {
	return place.latitude
}

// setter of latitude, no return value
func (place *Place) SetLatitude(latitude float64) {
	place.latitude = latitude
}

// getter of longitude, return the longitude
func (place *Place) Longitude() float64 {
	return place.longitude
}

// setter of longitude, no return value
func (place* Place) SetLongitude (longitude float64) {
	place.longitude = longitude
}

// print method
func (place *Place) String() string{
	return fmt.Sprintf("(%.3f°, %.3f°) %q", place.latitude, place.longitude, place.Name)
}

// a copy contructor, return a pointer to the new copy
func (original *Place) Copy() *Place {
	return &Place{original.latitude, original.longitude, original.Name}
}

func saneAngle(oldAngle, newAngle float64) float64 {
	return oldAngle
}

// -----------------------------------------------------------------------------

type Exchanger interface {
	Exchange()
}

type StringPair struct {
	first, second string
}
func (pair *StringPair) Exchange() {
	pair.first, pair.second = pair.second, pair.first
}

type Point [2]int
func (point *Point) Exchange() {
	point[0], point[1] = point[1], point[0]
}

func (pair StringPair) String() string {
	return fmt.Sprintf("%q+%q", pair.first, pair.second)
}

func exchangeThese(exchangers...Exchanger) {
	for _, exchanger := range exchangers {
		exchanger.Exchange()
	}
}

// implements the interface from the standard library: io.Reader, which has
// 'Read([]byte)(int, error)' as signature
func (pair *StringPair) Read(data []byte) (n int, err error) {
	if pair.first == "" && pair.second == "" {
		return 0, io.EOF
	}

	if pair.first != "" {
		n = copy(data, pair.first)
		pair.first = pair.first[n:]
	}

	if n < len(data) && pair.second != "" {
		m := copy(data[n:], pair.second)
		pair.second = pair.second[m:]
		n += m
	}
	return n, nil
}

func ToBytes(reader io.Reader, size int) ([]byte, error) {
	data := make([]byte, size)
	n, err := reader.Read(data)
	if err != nil {
		return data, nil
	}
	return data[:n], nil // remove useless bytes
}


// -----------------------------------------------------------------------------
type LowerCaser interface {
	LowerCase()
}

type UpperCaser interface {
	UpperCase()
}

type LowerUpperCase interface {
	LowerCaser
	UpperCaser
}

type FixCaser interface {
	FixCase()
}

type ChangeCaser interface {
	LowerUpperCase
	FixCaser
}

func (part *Part) FixCase() {
	part.Name = fixCase(part.Name)
}

func fixCase(s string) string {
	var chars []rune
	upper := true
	for _, char := range s {
		if upper {
			char = unicode.ToUpper(char)
		} else {
			char = unicode.ToLower(char)
		}
		chars = append(chars, char)
		upper = unicode.IsSpace(char) || unicode.Is(unicode.Hyphen, char)
	}
	return string(chars)
}


// implements the interface UpperCaser
func (pair *StringPair) UpperCase() {
	pair.first 	= strings.ToUpper(pair.first)
	pair.second = strings.ToUpper(pair.second)
}

func (pair *StringPair) LowerCase() {
	pair.first 	= strings.ToLower(pair.first)
	pair.second = strings.ToLower(pair.second)
}

// implements the interface FixCaser
func (pair *StringPair) FixCase() {
	pair.first  = fixCase(pair.first)
	pair.second = fixCase(pair.second)
}

// -----------------------------------------------------------------------------

func processPhrases(phrases []string, function RuneForRuneFunc) {
	for _, phrase := range phrases {
		fmt.Println(strings.Map(function, phrase))
	}
}

func main() {
	point := ColoredPoint{}
	fmt.Println(point)
	fmt.Printf("%#v\n", point)

	{
		type Count int
		type StringMap map[string]string
		type FloatChan chan float64

		var i Count = 7
		i++
		fmt.Println(i)

		sm := make(StringMap)
		sm["key1"] = "value1"
		sm["key2"] = "value2"
		fmt.Println(sm)

		fc := make(FloatChan, 1)
		fc <- -2.29558714939
		fmt.Println(<-fc)
	}

	{
		// declare a signature that receive a rune as parameter and returns a rune
		var removePunctuation RuneForRuneFunc
		phrases := []string {"Day; dusk, and night.", "All day long"}
		removePunctuation = func(char rune) rune {
			if unicode.Is(unicode.Terminal_Punctuation, char) {
				return -1
			}
			return char
		}
		processPhrases(phrases, removePunctuation)
	}

	/**
	 * ADDING METHODS
	 * --------------
	 */
	{
		var count Count
		i := int(count)
		count.Increment()

		j := int(count)
		count.Decrement()

		k := int(count)
		fmt.Println(count, i, j, k, count.IsZero())
	}

	// testing the custom Part type and some comstom methods
	{
		part := Part{5, "wrench"}
		part.UpperCase()

		part.Id += 11
		fmt.Println(part, part.HasPrefix("w"))
	}


	// test the Item and SpicalItem types
	// NOTE: any methods on the emebeded type can be called on the custom struct
	// as if they were the struct's own methods
	//
	{
		special := SpeicialItem{Item{"Green", 3, 5}, 207}
		fmt.Println(special.id, special.price, special.quantity, special.catalogId)
		fmt.Println(special.Cost())

		luxury := LuxuryItem{Item{"Swatch", 800, 10}, 2.5}
		fmt.Println(luxury.id, luxury.price, luxury.quantity, luxury.markup)
		fmt.Println(luxury.Cost())
	}

	// method expressions
	{
		part := Part{10001, "WRENCH"}

		asStringV := Part.String		// effective signature: func(Part) string
		asStringP := (*Part).String		// effective signature: func(*Part) string
		hasPrefix := Part.HasPrefix		// effective signature: func(Part, string) bool
		lower 	  := (*Part).LowerCase	// effective signature: func(*Part)

		sv := asStringV(part)
		sp := asStringP(&part)
		lower(&part)
		fmt.Println(sv, sp, hasPrefix(part, "W"), part)
	}

	// type validation
	{

	}

	// test interface
	{
		jekyll	:= StringPair{"Henry", "Jekyll"}
		hyde	:= StringPair{"Edward", "Hyde"}
		point 	:= Point{5, -3}
		fmt.Println("Before: ", jekyll, hyde, point)

		jekyll.Exchange()	// treated as: (&jekyll).Exchange()
		hyde.Exchange()		// treated as: (&hyde).Exchange()
		point.Exchange()	// treated as: (&point).Exchange()
		fmt.Println("After #1: ", jekyll, hyde, point)

		exchangeThese(&jekyll, &hyde, &point)
		fmt.Println("After #2: ", jekyll, hyde, point)
	}

	{
		const size = 16
		robert := &StringPair{"Robert L.", "Stevenson"}
		david  := StringPair{"David", "Balfour"}
		longname  := StringPair{"Thisisalongstringandwhatever", "hehe"}
		for _, reader := range []io.Reader{robert, &david, &longname} {
			raw, err := ToBytes(reader, size)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("%q\n", raw)
		}
	}

	// interface embedding
	{
		toastRack := Part{8427, "TOAST RACK"}
		toastRacl.LowerCase()

		lobelia := StringPair{"LOBELIA", "STACKVILLE-BAGGINS"}
		lobelia.FixCase()

		fmt.Println(toastRack, lobelia)



	}
}

