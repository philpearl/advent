package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	if err := part1(input); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	if err := part2(input); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func part1(input string) error {
	w := parseInput(input)
	_ = w
	w.readPacket()

	fmt.Println("version sum:", w.versionSum)

	return nil
}

func part2(input string) error {
	// World is extended
	w := parseInput(input)
	fmt.Println("result:", w.readPacket())
	return nil
}

type bitreader interface {
	next(bits int) int
}

type limitedBitReader struct {
	underlying bitreader
	count      int
}

func (l *limitedBitReader) next(bits int) int {
	l.count -= bits
	if l.count < 0 {
		panic(fmt.Sprintf("not enough bits! asked %d had %d", bits, l.count+bits))
	}
	return l.underlying.next(bits)
}

type byteBitReader struct {
	arena  []uint8
	cursor int
	bit    int
}

type world struct {
	reader     bitreader
	versionSum int
}

func (w *world) readPacket() (val int) {
	ver := w.reader.next(3)
	w.versionSum += ver
	typ := w.reader.next(3)
	fmt.Printf("ver %d, typ %d\n", ver, typ)
	switch typ {
	case 0:
		// sum
		for _, v := range w.readOperator() {
			val += v
		}
	case 1:
		// product
		val = 1
		for _, v := range w.readOperator() {
			val *= v
		}

	case 2:
		// min
		val = math.MaxInt
		for _, v := range w.readOperator() {
			if v < val {
				val = v
			}
		}
	case 3:
		// max
		for _, v := range w.readOperator() {
			if v > val {
				val = v
			}
		}
	case 4:
		// literal
		return w.readLiteral()
	case 5:
		// greater than
		v := w.readOperator()
		if v[0] > v[1] {
			val = 1
		}
	case 6:
		// less than
		v := w.readOperator()
		if v[0] < v[1] {
			val = 1
		}

	case 7:
		// equal
		v := w.readOperator()
		if v[0] == v[1] {
			val = 1
		}

	default:
		panic("unexpected type")
	}
	return val
}

func (w *world) readLiteral() int {
	var val int
	for {
		cont := w.reader.next(1)
		val = val << 4
		val |= w.reader.next(4)
		if cont == 0 {
			break
		}
	}
	return val
}

func (w *world) readOperator() []int {
	var vals []int
	lt := w.reader.next(1)
	if lt == 0 {
		numBits := w.reader.next(15)
		fmt.Printf("Operator %d bits\n", numBits)
		var lr limitedBitReader
		lr.underlying = w.reader
		lr.count = numBits
		w.reader = &lr

		for lr.count > 0 {
			fmt.Println("op: bits: count", lr.count)
			vals = append(vals, w.readPacket())
		}
		w.reader = lr.underlying
		w.reader.next(lr.count)

	} else {
		numPackets := w.reader.next(11)
		fmt.Printf("Operator %d packets\n", numPackets)
		for i := 0; i < numPackets; i++ {
			vals = append(vals, w.readPacket())
		}
	}
	return vals
}

func (w *byteBitReader) next(bits int) int {
	var out int
	for bits > 0 {
		bits--
		out = out << 1
		out |= int((w.arena[w.cursor] & (1 << (7 - w.bit))) >> (7 - w.bit))
		w.bit++
		if w.bit >= 8 {
			w.bit = 0
			w.cursor++
		}
	}
	return out
}

func parseInput(input string) (w world) {
	var bytes byteBitReader
	bytes.arena = make([]uint8, len(input)/2)
	for i := 0; i < len(input); i += 2 {
		v, err := strconv.ParseUint(input[i:i+2], 16, 8)
		if err != nil {
			panic(err)
		}
		bytes.arena[i/2] = uint8(v)
	}

	w.reader = &bytes

	return w
}

var input = `020D74FCE27E600A78020200DC298F1070401C8EF1F21A4D6394F9F48F4C1C00E3003500C74602F0080B1720298C400B7002540095003DC00F601B98806351003D004F66011148039450025C00B2007024717AFB5FBC11A7E73AF60F660094E5793A4E811C0123CECED79104ECED791380069D2522B96A53A81286B18263F75A300526246F60094A6651429ADB3B0068937BCF31A009ADB4C289C9C66526014CB33CB81CB3649B849911803B2EB1327F3CFC60094B01CBB4B80351E66E26B2DD0530070401C82D182080803D1C627C330004320C43789C40192D002F93566A9AFE5967372B378001F525DDDCF0C010A00D440010E84D10A2D0803D1761045C9EA9D9802FE00ACF1448844E9C30078723101912594FEE9C9A548D57A5B8B04012F6002092845284D3301A8951C8C008973D30046136001B705A79BD400B9ECCFD30E3004E62BD56B004E465D911C8CBB2258B06009D802C00087C628C71C4001088C113E27C6B10064C01E86F042181002131EE26C5D20043E34C798246009E80293F9E530052A4910A7E87240195CC7C6340129A967EF9352CFDF0802059210972C977094281007664E206CD57292201349AA4943554D91C9CCBADB80232C6927DE5E92D7A10463005A4657D4597002BC9AF51A24A54B7B33A73E2CE005CBFB3B4A30052801F69DB4B08F3B6961024AD4B43E6B319AA020020F15E4B46E40282CCDBF8CA56802600084C788CB088401A8911C20ECC436C2401CED0048325CC7A7F8CAA912AC72B7024007F24B1F789C0F9EC8810090D801AB8803D11E34C3B00043E27C6989B2C52A01348E24B53531291C4FF4884C9C2C10401B8C9D2D875A0072E6FB75E92AC205CA0154CE7398FB0053DAC3F43295519C9AE080250E657410600BC9EAD9CA56001BF3CEF07A5194C013E00542462332DA4295680`
