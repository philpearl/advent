package main

import "fmt"

func main() {

	run(input, 100)
	run2(input, 100, 10000)
}

var pattern = []int8{0, 1, 0, -1}

func run2(input string, phases int, amplify int) {
	in1 := inputToSequence(input)
	var offset int
	for _, v := range in1[:7] {
		offset = offset*10 + int(v)
	}
	fmt.Printf("offset %d\n", offset)
	// zero is repeated n-1 times when calculating the n-1th character

	in := make([]int8, len(in1)*amplify)
	for i := 0; i < amplify; i++ {
		copy(in[i*len(in1):], in1)
	}

	out := make([]int8, len(in))
	for ; phases > 0; phases-- {
		var prev int
		prev = applyPattern(in, offset-1)
		out[offset-1] = tens(prev)
		for i := offset; i < len(in); i++ {
			prev = applyPatternDiff(in, i, prev)
			out[i] = tens(prev)
		}
		// fmt.Printf("%d\n", out)
		in, out = out, in
	}

	fmt.Printf("%d\n", in[offset:offset+8])
}

func run(input string, phases int) {
	in := inputToSequence(input)
	out := make([]int8, len(in))
	for ; phases > 0; phases-- {
		for i := range in {
			out[i] = tens(applyPattern(in, i))
		}
		in, out = out, in
	}

	fmt.Printf("%d\n", in[:8])
}

func tens(val int) int8 {
	if val < 0 {
		val = -val
	}
	return int8(val % 10)
}

func applyPattern(seq []int8, elt int) int {
	var val int
	// we know pattern is going to be zero for the first elt multiplications
	// Pattern is
	// 0 repeated elt times
	// 1 repeated elt+1 times
	// 0 repeated elt+1 times
	// -1 repeated elt+1 times
	// 0 repeated elt+1 times
	var i int
	sign := 1
	for (2*i+1)*elt+i < len(seq) {
		val += sign * sum(seq, (2*i+1)*elt+2*i, 2*(i+1)*elt+(2*i)+1)
		sign = -sign
		i++
	}
	return val
}

func sum(seq []int8, start, end int) int {
	if start >= len(seq) {
		return 0
	}
	if end > len(seq) {
		end = len(seq)
	}
	var total int
	for _, s := range seq[start:end] {
		total += int(s)
	}
	return total
}

// For larger values of elt we calculate the difference between applyPattern for
// elt and elt-1. This is a smaller number of calculations for larger values of
// elt
func applyPatternDiff(seq []int8, elt int, prev int) int {
	val := prev
	var i int
	sign := 1
	for (2*i+1)*(elt-1)+2*i < len(seq) {
		// we're basically doing prev + sum(current params) - sum(prev params)
		this := -sum(seq, (2*i+1)*(elt-1)+2*i, (2*i+1)*elt+2*i) +
			sum(seq, 2*(i+1)*(elt-1)+(2*i)+1, 2*(i+1)*elt+(2*i)+1)
		val += sign * this
		sign = -sign
		i++
	}
	return val
}

func inputToSequence(in string) []int8 {
	out := make([]int8, len(in))
	for i, c := range in {
		out[i] = int8(c - '0')
	}
	return out
}

var input = `59766977873078199970107568349014384917072096886862753001181795467415574411535593439580118271423936468093569795214812464528265609129756216554981001419093454383882560114421882354033176205096303121974045739484366182044891267778931831792562035297585485658843180220796069147506364472390622739583789825303426921751073753670825259141712329027078263584903642919122991531729298497467435911779410970734568708255590755424253797639255236759229935298472380039602200033415155467240682533288468148414065641667678718893872482168857631352275667414965503393341925955626006552556064728352731985387163635634298416016700583512112158756656289482437803808487304460165855189`
