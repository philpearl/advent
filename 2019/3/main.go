package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	if err := run1(input); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	if err := run2(input); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func run1(input string) error {
	path1, path2 := parseInput(input)

	// mark up our arena with the positions on the first path
	//      16 15 14 13 12
	//      17  4  3  2 11
	//      18  5  0  1 10
	//      19  6  7  8  9
	//      20 21 22 23 24 25
	//
	// I'm sure there's some kind of easy transformation from x,y to this system

	// Note all the positions that path1 reaches
	points := make(map[position]struct{})
	var pos position
	for _, v := range path1 {
		switch v.dirn {
		case 'U':
			for i := 0; i < v.dist; i++ {
				pos.y--
				points[pos] = struct{}{}
			}
		case 'D':
			for i := 0; i < v.dist; i++ {
				pos.y++
				points[pos] = struct{}{}
			}
		case 'L':
			for i := 0; i < v.dist; i++ {
				pos.x--
				points[pos] = struct{}{}
			}
		case 'R':
			for i := 0; i < v.dist; i++ {
				pos.x++
				points[pos] = struct{}{}
			}
		}
	}

	// find all intersections
	pos = position{}
	var minD int = math.MaxInt64
	var minPos position
	for _, v := range path2 {
		switch v.dirn {
		case 'U':
			for i := 0; i < v.dist; i++ {
				pos.y--
				if _, hit := points[pos]; hit {
					if d := pos.manhat(); d < minD {
						minD = d
						minPos = pos
					}
				}
			}
		case 'D':
			for i := 0; i < v.dist; i++ {
				pos.y++
				if _, hit := points[pos]; hit {
					if d := pos.manhat(); d < minD {
						minD = d
						minPos = pos
					}
				}
			}
		case 'L':
			for i := 0; i < v.dist; i++ {
				pos.x--
				if _, hit := points[pos]; hit {
					if d := pos.manhat(); d < minD {
						minD = d
						minPos = pos
					}
				}
			}
		case 'R':
			for i := 0; i < v.dist; i++ {
				pos.x++
				if _, hit := points[pos]; hit {
					if d := pos.manhat(); d < minD {
						minD = d
						minPos = pos
					}
				}
			}
		}
	}
	fmt.Println(minPos, minD)
	return nil
}

func run2(input string) error {
	path1, path2 := parseInput(input)

	// mark up our arena with the positions on the first path
	//      16 15 14 13 12
	//      17  4  3  2 11
	//      18  5  0  1 10
	//      19  6  7  8  9
	//      20 21 22 23 24 25
	//
	// I'm sure there's some kind of easy transformation from x,y to this system

	// Note all the positions that path1 reaches
	points := make(map[position]int)
	var pos position
	var steps int
	save := func() {
		steps++
		if _, ok := points[pos]; !ok {
			points[pos] = steps
		}
	}

	for _, v := range path1 {
		switch v.dirn {
		case 'U':
			for i := 0; i < v.dist; i++ {
				pos.y--
				save()
			}
		case 'D':
			for i := 0; i < v.dist; i++ {
				pos.y++
				save()
			}
		case 'L':
			for i := 0; i < v.dist; i++ {
				pos.x--
				save()
			}
		case 'R':
			for i := 0; i < v.dist; i++ {
				pos.x++
				save()
			}
		}
	}

	// find all intersections
	pos = position{}
	var minSteps int = math.MaxInt64
	steps = 0
	save2 := func() {
		steps++
		if steps1, hit := points[pos]; hit {
			if s := steps1 + steps; s < minSteps {
				minSteps = s
			}
		}

	}

	for _, v := range path2 {
		switch v.dirn {
		case 'U':
			for i := 0; i < v.dist; i++ {
				pos.y--
				save2()
			}
		case 'D':
			for i := 0; i < v.dist; i++ {
				pos.y++
				save2()
			}
		case 'L':
			for i := 0; i < v.dist; i++ {
				pos.x--
				save2()
			}
		case 'R':
			for i := 0; i < v.dist; i++ {
				pos.x++
				save2()
			}
		}
	}
	fmt.Println(minSteps)
	return nil
}

type position struct{ x, y int }

func (p position) manhat() int {
	if p.x < 0 {
		p.x = -p.x
	}
	if p.y < 0 {
		p.y = -p.y
	}

	return p.x + p.y
}

type vector struct {
	dist int
	dirn byte
}

func parseInput(in string) (path1, path2 []vector) {
	paths := strings.Split(in, "\n")

	return parsePath(paths[0]), parsePath(paths[1])
}

func parsePath(path string) []vector {
	steps := strings.Split(path, ",")
	vv := make([]vector, 0, len(steps))
	for _, step := range steps {
		dirn := step[0]
		dist, _ := strconv.Atoi(step[1:])
		vv = append(vv, vector{dist: dist, dirn: dirn})
	}
	return vv
}

var input = `R1009,U286,L371,U985,R372,D887,R311,U609,L180,D986,L901,D592,R298,U955,R681,D68,R453,U654,L898,U498,R365,D863,L974,U333,L267,D230,R706,D67,L814,D280,R931,D539,R217,U384,L314,D162,L280,U484,L915,D512,L974,D220,R292,U465,L976,U837,R28,U68,L98,D177,L780,U732,R696,D412,L715,U993,L617,U999,R304,D277,R889,D604,R199,U498,R302,U958,R443,U957,R453,U362,R704,U301,R813,U404,L150,D673,L407,D233,L901,D965,R602,U615,R496,U467,R849,U530,L205,D43,R709,U127,L35,U801,L565,D890,R90,D763,R95,D542,R84,D421,L298,D58,R794,U722,R205,U830,L149,D759,L950,D708,L727,U401,L187,D598,L390,D469,R375,U985,L723,U63,L983,D39,L160,U276,R822,D504,L298,D484,L425,U228,L984,D623,L936,U624,L851,D748,L266,D576,L898,U783,L374,D276,R757,U89,L649,U73,L447,D11,L539,U291,L507,U208,R167,D874,L596,D235,R334,U328,R41,D212,L544,D72,L972,D790,L282,U662,R452,U892,L830,D86,L252,U701,L215,U179,L480,U963,L897,U489,R223,U757,R804,U373,R844,D518,R145,U304,L24,D988,R605,D644,R415,U34,L889,D827,R854,U836,R837,D334,L664,D883,L900,U448,R152,U473,R243,D147,L711,U642,R757,U272,R192,U741,L522,U785,L872,D128,L161,D347,L967,D295,R831,U535,R329,D752,R720,D806,R897,D320,R391,D737,L719,U652,L54,D271,L855,D112,R382,U959,R909,D687,L699,U892,L96,D537,L365,D182,R886,U566,R929,U532,L255,U823,R833,U542,R234,D339,R409,U100,L466,U572,L162,U843,L635,D153,L704,D317,L534,U205,R611,D672,L462,D506,L243,U509,L819,D787,R448,D353,R162,U108,R850,D919,R259,U877,R50,D733,L875,U106,L890,D275,L904,U849,L855,U314,L291,U170,L627,U608,R783,U404,R294
L1010,D347,R554,U465,L30,D816,R891,D778,R184,U253,R694,U346,L743,D298,L956,U703,R528,D16,L404,D818,L640,D50,R534,D99,L555,U974,L779,D774,L690,U19,R973,D588,L631,U35,L410,D332,L74,D858,R213,U889,R977,U803,L624,U627,R601,U499,L213,U692,L234,U401,L894,U733,R414,D431,R712,D284,R965,D624,R848,D17,R86,D285,R502,U516,L709,U343,L558,D615,L150,D590,R113,D887,R469,U584,L434,D9,L994,D704,R740,D541,R95,U219,L634,D184,R714,U81,L426,D437,R927,U232,L361,D756,R685,D206,R116,U844,R807,U811,L382,D338,L660,D997,L551,D294,L895,D208,R37,D90,R44,D131,R77,U883,R449,D24,R441,U659,R826,U259,R98,D548,R118,D470,L259,U170,R518,U731,L287,U191,L45,D672,L691,U117,R156,U308,R230,U112,L938,U644,R911,U110,L1,U162,R943,U433,R98,U610,R428,U231,R35,U590,R554,U612,R191,U261,R793,U3,R507,U632,L571,D535,R30,U281,L613,U199,R168,D948,R486,U913,R534,U131,R974,U399,L525,D174,L595,D567,L394,D969,L779,U346,L969,D943,L845,D727,R128,U241,L616,U117,R791,D419,L913,D949,R628,D738,R776,D294,L175,D708,R568,U484,R589,D930,L416,D114,L823,U16,R260,U450,R534,D94,R695,D982,R186,D422,L789,D886,L761,U30,R182,U930,L483,U863,L318,U343,L380,U650,R542,U92,L339,D390,L55,U343,L641,D556,R616,U936,R118,D997,R936,D979,L594,U326,L975,U52,L89,U679,L91,D969,R878,D798,R193,D858,R95,D989,R389,U960,R106,D564,R48,D151,L121,D241,L369,D476,L24,D229,R601,U849,L632,U894,R27,U200,L698,U788,L330,D73,R405,D526,L154,U942,L504,D579,L815,D643,L81,U172,R879,U28,R715,U367,L366,D964,R16,D415,L501,D176,R641,U523,L979,D556,R831`
