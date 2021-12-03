package main

import (
	"fmt"
	"strings"
)

func main() {
	run(input)
}

type vec [3]int

func (v *vec) move(w vec) {
	for i := range *v {
		(*v)[i] += w[i]
	}
}

var step = map[string]vec{
	"e":  {1, -1, 0},
	"w":  {-1, 1, 0},
	"se": {0, -1, 1},
	"sw": {-1, 0, 1},
	"ne": {1, 0, -1},
	"nw": {0, 1, -1},
}

func run(input string) {
	world := make(map[vec]struct{})

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		var loc vec
		for i := 0; i < len(line); i++ {
			c := line[i]
			var s vec
			if c == 's' || c == 'n' {
				s = step[line[i:i+2]]
				i++
			} else {
				s = step[line[i:i+1]]
			}
			loc.move(s)
		}
		_, black := world[loc]
		if black {
			delete(world, loc)
		} else {
			world[loc] = struct{}{}
		}
	}
	fmt.Println(len(world))

	// Next run the life algo for 100 steps
	next := make(map[vec]struct{})
	for i := 0; i < 100; i++ {
		for loc := range world {
			if c := countAdjacent(world, loc); c == 1 || c == 2 {
				// This black tile remains black in the next world
				next[loc] = struct{}{}
			}
			// Consider adjacent white tiles. Note work will be
			// repeated
			for _, dir := range step {
				l := loc
				l.move(dir)
				if _, black := world[l]; !black {
					if c := countAdjacent(world, l); c == 2 {
						next[l] = struct{}{}
					}
				}
			}
		}

		world, next = next, world
		for k := range next {
			delete(next, k)
		}
	}
	fmt.Println(len(world))
}

func countAdjacent(w map[vec]struct{}, loc vec) int {
	var count int
	for _, dir := range step {
		l := loc
		l.move(dir)
		if _, black := w[l]; black {
			count++
		}
	}
	return count
}

var testInput = `sesenwnenenewseeswwswswwnenewsewsw
neeenesenwnwwswnenewnwwsewnenwseswesw
seswneswswsenwwnwse
nwnwneseeswswnenewneswwnewseswneseene
swweswneswnenwsewnwneneseenw
eesenwseswswnenwswnwnwsewwnwsene
sewnenenenesenwsewnenwwwse
wenwwweseeeweswwwnwwe
wsweesenenewnwwnwsenewsenwwsesesenwne
neeswseenwwswnwswswnw
nenwswwsewswnenenewsenwsenwnesesenew
enewnwewneswsewnwswenweswnenwsenwsw
sweneswneswneneenwnewenewwneswswnese
swwesenesewenwneswnwwneseswwne
enesenwswwswneneswsenwnewswseenwsese
wnwnesenesenenwwnenwsewesewsesesew
nenewswnwewswnenesenwnesewesw
eneswnwswnwsenenwnwnwwseeswneewsenese
neswnwewnwnwseenwseesewsenwsweewe
wseweeenwnesenwwwswnew`

var input = `neesewweeseeseesewseeenwneswe
neneswnenenwnenenwnwenesenenwnwnenenwne
neswnweswseswnesewsesweswwswswneswsew
neneeneneeneewnenweneneeeeswswe
seeeeseseseeeeweeeseneenweeee
swseswwswswswswwswswswswswswneswwswsw
seseswsenwseseseseseseseesenwswswsenwse
seswnesenwsesweseswesenwsesesesewsenwsese
enwwwsewwnwwwwnenw
eeewseeeeseseseeeeseseseeneese
eseneswwseseseseseswsesenwseseseswswwse
nesweeeneneneenenwnenenenwseseeneenew
eseeeeseeneewnwseseenwnweeeene
nwswnwsenwwnewnenwnwnwnwnwnwenwnenwnwnwse
swneseneseseswseswwseneswnwseseesesesesw
wswnwwwwewnwwwwnwwwwwwnwnww
ewwwswwwwnww
wwnwwwwwwwseswwewwwnewsww
swwwnewswwseswswswswswswswswsw
wswwwwnewwwsewwwswwwwwww
nenweseswsesesesenwsewsesewsesenenesesw
swswswnwsenwseseseswseseswsesw
seseewseeseseseesesesesesesesesesese
swesweeeseneeneeeeeeeswenwese
wnwnwnwwwnwnwnwwnwwwnenwnwse
neseeseseseswsenesewesenesewswnwnwse
sweeneswwswswswnwswswsw
eeeeeeweeeeeeeeswenwneeew
swwesewwnwswenenenwnwnewseswwswew
seswswswnwesenwwswswsweeswwenwnwenw
nwnewenwwswnwewsewswnesenwenwwsew
nwnwnwnwswenwnwswsenenwwnwnenwnwseese
wwnewwwwwwnwwwswnwnwnwnw
nwenwwnwnwnwnwnwnwwnwnwnw
sewnenewwnenenwswnwenwnwseneeswnenwne
enewneenwneneeenewswseseneewewne
neneneswnwswneewnenenenwneeneneeneene
nweeeneeneseneneneneneneneeneneee
swwnenwnwnwwnwsenwwnwsewnwnwnwnwnwnenw
wwseswwwwnwneesewwseswnwsenwwwnew
seseswseswneswseswswseswwswseswneswsw
nenenwswnenwswnwneneswswenewewneene
sweswnwwneswseseseneseswseseswneseswsesesw
nwnwnenwnwwsenwnwnwnenwwnwnwswsenwnww
wnewwsenwwsenwwswwwnwwnenwnenww
nwwwwswseswwswnesw
eseeeeeeeeeeeeeswenewee
neeneneeswneneeneneneeenewnwnenenene
wneeswnwwsewenwneneswnwneesewneesw
nwwnwneseswnwwenwwsenwenenwnwseeseww
swnewwswswswswwswswwnwswswswseswwsw
sesenwseswneswswswswswnwswnwseswseeeswsese
nwnwnwswnweseswnwnwnwnenwnwnwne
swsweeeeeneeseeneswneswnwnewsesw
seswseswseseswswseseswseneseweseswsesese
eeneenenenenenesweneneenene
seeseeeeeseeeeeeeeeenwese
nenwnewwswwwnwwswwnwnwneswwsewnw
seseseswsesesesenweeesenwseseseesesese
wwwnesenesweswswnwwwwnwnwnwsewne
swswweseseswswsweseswswswswswseswsww
eneeseseneenwwnwwneneeseeseswnwne
seswswswenesesweseswnwswsesenwwsesw
wwwwswwwwnwwswnwneesewswswsew
wnwseeneswseeeseseeseewwseenwnee
wseeseseseeseseseseseesesesese
seswswseeseseswseswswseswseswnw
wwwwwwwwwewswwswnwnwseswwwwsw
eeeenwsweeeesweeneeeswenene
nwneewnwnewnenwenweneneswsewsenwsw
neeeneneeeeeeeneeeneeeneew
nwnwnenwwnwnwnwnwnwwnwnwnwnwnwnwsenwnw
sweenwnwswnwnwnwnwnwnenewnwsenwneeswnenw
seswseneneswnwswnweseseswsesewsesewswwe
wswwnwsewnewnwwswwewwswwsesewnww
wseswseseswswsesewnenwnwseneseneswsew
swseswwsesenesenwnwswseeswneswswseseswnesw
swnenwneneneenenewnwnenwnwseeswnwwne
senewnwnewseseneeseswswnenwsewwsesese
nwsweswswswseswswsweswswswnwseswswwnwsw
nwnenwnenwswnwnenenenwnwnenwnenwswneneenw
nwnwnwnwnwnwnenesenwnwnenwnenwnwnwnenwnw
wwnewwnwwwwsewwwwnwwwwww
swseneseseswseswseswseseesesesesew
eswswwswswswswwwswswswnenesweswswsw
swneneweneneneneneseneenenweeenene
eseeeseseesesenweseneseesenwnwsewse
eeeeeeeeeeeweeeeeeee
swneesewneenwnenesenenesenenenwwnewne
seseswsewseneneseswseswsewsewseswesese
enweneeneneeneeswneeseeneeee
nenenenwsenwnenwenenwwneneneswnenenesw
seseseeseneneeseseeseesesesewewse
neswseswseseseswswseswwseenwseswswnwswsw
eneneeweeeeneseeeneeeeeewnene
nwenwswesweswenwe
seswswswseswswseswswswswwneswswswswneswsw
enenenenenwnenenenenenenenenesenenene
wseswnewneswwwswswwswewnenewwse
wwseneseweneswneeneeswnwne
wnwwwwwwwewnwwnenwwsewwsww
seneswneneenewswnenenwenweenenenesee
seseeseeswewseeeeenwenwsweene
eneeeneeeeeneneeeneeneswwesenene
newswnwsenwwswenwnewne
nwnwnwnwwnwsenwnwwnwwnwnwnwnwnwnwnwnw
neneeeeneeneneeeeneeeesweee
wnenenwswsewenenwnenwneneswnweenenwwe
nwwsenenwseswwsenwwnewnenwswwnwsenw
wnwwnwwwwewnewwwwwsewwwwswnw
nwnwsenwnweseswwswsenenenwnw
neenenenwnwnenwneneswnenesenenwnenwnenwne
nwwsenewnenwnesene
eeeeeeesewsweeeswenweenwe
swwnewwseseswwswwnwnenewwnwsenese
wwswswseeneswswnwwswswswwswswswswswsw
sewnwnwnenenwenenwnwnwnwsenwswnwnwnwnenwnw
nwnwnenenwneswnwnenwnwswnwneenwnwnenwnwnw
nesenenwnwnwnwenenwnewnwnwnwnenwnwnww
wnesesesweseeeswneseenew
nenwewesenwnewwwnwenewswnwnwseenw
weeseeseeseneseseseeseeseesesee
swwswwwswswswwswwwwswswwwew
nwnenenenenwneswnwnwsenenenenenenenwnwnwnwe
seseseseeseseseseseseeesenwseeswseee
swseseseswseseswseswswsenesesesesesw
nwnenwnesenwnwnenwnwnenenwnwwnwnenwnwnwne
swswswswswseseswsesweseseswnwswswsw
ewwnwwwnwwwswweww
seewsenenwswseswswnwseswseseswneseswsesw
wnewwnwwwwwwwwwwwsewwwww
sesesesesenwseseseseseswseseseesenwnwse
swswswswswswswswswswswswswswsweswswnwsww
swswewseswswwwsewswwnweswswswswnenw
nwsewwnwnwnwnwwnwnwwwwnwnwnwnwnww
wneneseneneneneneesewnenenenenene
eenwseeswswswnweseeseswneeenwnwne
neneneneneneneneneneswswnenenenenenenenene
swswswsesewswswnenwswswswseswswswswseesw
nesenenwseswswswnwswseseswseswsweswseswe
seeswenwsewneswneneswswnwswnwseswsene
nenwnenwnwnwnwnesenenenenenenwwnwnwnwnwnw
wneeneneneeneneneenenenesweneneenene
nwenenwnenenwnwnwnenwnwswseenwswwnesw
neseseneneseswswswwswsenesesw
swswswwnweswsewewswswswswswswswsww
nwswswseseswswseseeswsesesewseseswnesesw
wnenwneswnenenwseenwnwenewnwnenenw
eeeneneneswnenenenenene
swswwswswswswwswweswwswswwnwwwww
nwswnwwnwwnwswnweseswenenwesenwnwnew
seeseseseseeeseseeseseenweeweee
swseseseseseseseswseseneseseseseseswsenw
nenwnenenenenenesenenenwnenenewnene
nenenenweenewneneswnesenenwnesenewne
swswswswwswswswneswswswswswswswsw
nwsesenwnwseseseseweeseseswsesesesese
eeneneneneneseneenenenenwnenenenenene
sweswneswneneseswneenenenwneneswne
eswseeseseeseseseseneseeeseewnee
neneeeweeeneneneeneneneeneeese
sweswswneswwwwswswswnwswnwseswswswwsw
wnwseseseeswnwenwsesweseweseseswnwse
seseeeswsewseseseesesenee
nenenenwneneneneneneneneneseenenenenene
nwnwnwenwnwnwnwnenwwnwsenwnwnwnwnwnwnw
wsewwwwswnwnenwwneenwwwnwsenwwnwse
eseeeswenweneweseseeeeeeeee
nwnwwwnwswnwwwwewnwnwwwwwnwwnw
esesesenesesewseseswsesenee
seswenwnesenweswsenwnewsese
nwnwswnwnwenwenwnwnwnwwnwnwnwnwswnwnwnwnw
seswnwwsesewneseseseeeswseswswsenesenw
wseswnewsewwwseswnenenwneswswne
esweeeeewneeneneneeeeeeee
eeweseswwswswwwnwswnwesenwwnenww
wswswwneswwnwwwswswwsewe
enwswwnwnwnenwnwnwnwnwnwnwnw
nwnwneneneneneneenenenenwnwneseswnwnwnww
wnwnwwwnwnwnwnwsenwnwwwnwenwswnenw
neseseneseewnesesenwseeeseswwseese
swswswwswswswswnewswswswswswswswsww
wwwsewwwnwswwnwnwnwnewnwwwnwnw
enwwwneesesweseeseneswewsesenesesese
esesesesesesenwsenwseseesesesesesesesesww
eweeeeesweeeeeeeenweee
newswseswseseseswsweswswnwneswwswsesw
nenwswnwsenwseswnwnwwswweesesewnenwse
seseseseseswseeswseswseswsenwswseswseswse
seseenewesesesesesesesesesenesesesesw
sweneesewsenwswswseseeeseenwenee
seeweseswnweseseeseeweeeeese
esweseeeeseeeeeewneeeeese
wwwswnwswswswweewswswswswwswwsw
swseswsenesesesesesesesenee
swenwnwnwswnwneswnwnwnwnenwnwnwseeswe
swnwnenwnwnenwnenwnwnwnenenenwnweenwwnw
nwswswenwneseswseenwnesenenewwnwwse
esweewneneeseneswnenee
eeeseneweneneeneeseneneneneneenenew
swneswswwwwswswswswwwwswswswwsww
eeseeweneeeesesese
senenwsenwwnenenwnwwnwnwnwnenwsenwsenenene
eneeenewwsweeeneneseeswneewenw
senwnenewenwnwneswweneneneseeneenesw
eesenweeseesweewewneeeseee
wsenwewnwwwwnwse
nwnwnwsenwnwnwnwnwnwnwnwnwnwnwsenwnwnwnwnw
eeweseneeeeweeneneeeneeenw
eseeeeseeseeseseeseseseenweese
sweeswneseesenenenwswwneneswnweenwnw
nwneenwswnenenenenwnwnwnwnenwnenwnenwnene
wnwnwnwwnwwwnwnwnwwsewnwnwenwswnwww
wwwwnwwwwwwwwwwwwwwenw
wwwswwswsewwswnewnwwwewswwneww
senwswseseesenesewseseseneeswsesesesesese
nwnwswsewsewswnenenwew
newnenwnenwnwneneenenenenenenwnwnenene
wnwnwnwnwnwwnwsewnwwwwwnwwnwwe
swswswswneswswswswswseswswswswswswseswsw
eneneneeneneseswnenenenenenenenewnenene
nwnwnwnwnwnwenwnenenwnwnwnww
nenwwswnwnesewenwnesewwnwsenesewswsw
nwwnenenwneeneneseenewnenwnwnenwnenwne
swswswswswswswswswswwwwsweswneswswe
eneneneneneneneswneswnenenenenewnenenene
eeeneeeeeneewseneeeeeeene
eseseeseneseswwsesesesesewsesesenwsw
nwnenenesenewnenenwnenenenenwnenesenenenene
wnweewswnwswewneswnesenwwewswew
wswswswsenweeeweneeeneewneeesene
eeeeseeeseeeseseeenweeesesw
swswswnwnwenwnwswnenwnenwnene
swwswswswswwwwswwweswwwswswneswsw
wwswwswseneswwwswswswenwsweswsenew
swwswswseswseswswseswswneseswswswseseswnesw
eeneeeeesweeeee
swwswswwswwswwneswwneswswswswswwsw
seesesenwseseseswse
nenwnwnwesweneswnenenenenenenwswne
nwnwwnwnwwnwsenwnwnwwsenenwnw
swseswwnwswswswswswseswswswenew
enweeenwesweeeeeeeeeesenwsw
eswseenwswnwsesesesenwseewwseenesenw
wswseeswneeeneewneneneeenwneeswnw
newnwsenwnenweneeeneesweseeswnw
nwnwwswwnwwwnenwnwnw
wnesenenenwnenenenewnenwnenenesenenene
swswwswsenwswswswswswswsw
sewseswswneseseswseneswswseesesenesese
swwseswwswswswwnwswwwwnewwwwww
wwwnewsewwwwwwwwwwwwwww
newnenwwnwesewnwnwwenwwnwsenwwnwsw
nwwsewwwnweseneneswswwsenewwsenw
swswswwswswnwsweswswseswswswweswswswswne
wswseseswsenwnweseeseswswse
nwnwnwnwnwswnwwnwnwnenenwnwswwnwnwewnw
neswswswseswswsenwesweseneseswwwnwswnwsw
ewwswswwswswswwswswswswswnweswwwsww
eeseeseeseeeeseeeeewseeese
neeeenwneweeeeswseeeneeneee
neenenwswnenesenenenwneneneneneneewwne
neneeeneneneswnenenenenenewnenenenenee
neswnenenenwnwswnwnenenenenee
wwwwenwwnewwwwnwwwwwwwnwsw
nenewsenesweseneeneenwwnenwneswwsw
sewwwnewwnewswwwww
senwnwenwwnwsenenwnwswenenwnwswnwnwsenew
nwnwnwnwnenwnwnwswnwnwnenenwnwne
wswseeenwewnweeeewswseesesese
wwswwwwwwwseneeenewwwwwse
eeeseeeeeneesweeeeeeeese
neesewneneneneeneeeseeeenweneee
seseseseseswseseseseseseswsesenesesesese
senenwnenenenwnenenenenweneswnewnwnenene
sewswwesweewswwnwswswwwwswsww
eeneneeneneesweneeneneneswnenenwe
swneswswsesesesenwswesewwneseswsweswnw
nenesweneeeneeneneneneneneneee
nwwsewnwnwnwwnewnwnwnwwnwwwnwsenw
wnenenwseswneneneswenenwsenwneswnwne
eeeeesweeeeeeeeneeneeee
nwnwnwnwwwnenwswnwnwewnwwwwwnwnww
newnwnwnwnwsenwnwnwsenwsenenwnwwnwnwnwnww
eseseswseseseseseneseseseeswenwesew
wswwwwswwwwwwwwwnwwwsewsw
nwnwnwnwwnwnwnwnwnwwwwnwnwnwwnwnwe
nweneseneneseewnwnenwwseseneneeenee
nwnwnenwnwnwnwnwnwnwnwnwnenwsenwnwnwnwse
seseseseesesesewsenwsesese
sweenweesenwwewsenwswsesweenene
nwswnwnenwenwnwnwnwneeswenwnwnwwnwnwne
nenenwwnenesenesenenenenenenesenenewswnw
senweswnweseesweeesenwseenweenw
seswsweseswewswnwnwswswseswnwseswswne
enwnenwnwnwnwsewnwnwswwnwnewnwsew
eneeneweneeeeeeneeeeenweesw
seseseseenweseswwewneeew
neenwsesweeenesweeee
nweenenenwneneseswesweeneneneneene
eseeeeesenweeesweeseeseeseee
nwnwnenwnwnwsenwnwnwnwnwnwnwnenwnwnwnwwnw
seseseswswswswseswnwswsw
eeswenwneneeeeneneeneeeeeenene
wnwnwnwwwnwwnwnwwswe
esenenenenenenwneeswnwnwnenenenwnwwnw
seseseswsesenesesesesewse
nwnwswnwswwnwnwnwnwnwnwnwwnwenwnwwew
seseseeeeeesweeseeseene
seeswwwneeswswnesweswswwenwswnwsww
nenenwnwnwswenenenwnwnenwnenenwnwnwnwne
swseswswwswswswnwnwnwseswwwwewwsw
neenwseeneesesenewneenwnwneeenwsesw
wswswswseesenwneswnwswseseneseswwswswsw
swenesenesesenwswsesweswwnwswnwwsenwsw
nwnwnenenwnwnwnwnenenenwnwnese
swswseswseeswsweswswseseseswswswwnwswswse
swswsweswwwswswswsweswswseswswswswsw
nwnwsenwwnwnwnwsenwnwnwnenwnwnenwnwwnene
nwwnwnwnwwnwwwwwwwnwewwewww
swwswswseswswswswswswsweswswswswswswsw
nesesesesenwewswnwsesesesesenw
seseeeseseeseeseeeesenwseseeswwe
swseswswswswseseswsenwnwswsw
swswwwwwewswswewnwnwnwswewwsw
swswnewswswsenewswnenwneeesenwswwsw
swswswswswswswswswswswswnewswswswwswsw
eweseeenwseseeseeseseeseseseee
sewseseswesewseseneseswswseesenwnwsw
wswwswwwwwnwswswswsenewwswwswww
nwnwnenwwnwwwwwnwwnwwwnewswnwsw
neseswwsenwnwswwsenwneesenene
sewewwswwwswswwsewnwneswswwwnw
nenenewnewsweeeeeseswswseenwnee
wwwnwnwneseswewswwswseswwswenewsw
neswsewswswswwswwswswswwwswswsw
nwnwnwenwwnwenwsewenwnwswwwnwnwnwnw
nwnwnenenwnenwnwnenewnwseneswnwnwseenwnw
neweeseenwseeeeeeseseenwneeew
seswnwnweseweeswnwnewnesweswnenenw
eneewneneeneeneenenwnesweeseene
eenwsweeneneneneeeneeneeeneenene
neneswnesweneenewneseneswnenwnenwwse
swneswswsesewwseswseswswswseneswswwswne
nwnwnwnwnwnwnenwnwnwnwnenwnwswenwnwnw
wswwswwwnewswsenwwwswswsewwswsw
wsewwwwwwwwewwnwwwwwww
seswsweseswswnwswsesenwseseseswwseswse
wwwswwswwwswswswswswswwswswswew
swswseseswswsesesenwneswswswseseeseswwsese
enwswnwnwnwnwneneneeswnwneeswnewnwnenw
swswsewswswnwswnwneswswwwwnwwewse
newneeneneneneneenenenenenewnewenese
nwnwnewnwewenenw
swwnesenwseneneswnenenenwnw
eeeeeeeneeeeeseweeneeene
swswswwwsesewswswswwwnwwwswwswwne
ewsenwwewwewnwwwwwwnewwww
nenewneeneenenenenewneneeneneneseneene
swswsewseswswseswweswswswswneswswseswsw
swseswseswsesenwsweswwseswseswseswenese
swewwnewseswnwwwwswswwwsewnene
seenesesweseenwnwseseesesewsesesweese
nwswsenwenwnwnwswnwewnwnwnwnwnwne
wsewwnewwwwwnwswwwewwwnw
eeeeenwswnwswee
wswwwwwwwwwwsenewwwwwwww
nenenenwseseneeseneneneeswnwesenwnwnwse
eneswswsweweseswswsenwswswswwwnesese
nenwnwnwsenwnewnwnwnwewenwnenwnwnwnwnw
newnwswwwwwnewswsewwwwwwswswne
sewseneseseseseseseseseseeseewsesese
enenenweeneeeeneneneeeneeneeesw
esweeeneeseeeesweeeseswnwnwenw
seseswsweeswswwseseswswseswseswsesenw
nenwneneenwnenwnwnenwneneneeswneneneneswnw
nwnenenwnwswnwnwnwnenwnenwnwnwnwnwneswne
ewnewnwswswweeswwewnwenenewwswse
eeesweseenwenwesweweeeeseswne
nenwswnenewwnweneneswewneneeneneswnee
swseseseseeswseseswseseswsenwseswsesenw
eneseeeswseseesesweeenwse
enenenenwwnwneenewnenwnenenwswnwsesenw
wswwswswswswnwswsweswwswswswswswswswsw
nenwneswnenenenenenenenenesenenenenenenene
swnenwnenewsewneneneneneneseeneenenenenw
wwswwnewswswwwwwenwsw
seeeseeenwneseseseeeeeewsesese
seneseseseswseseseswnwswseseswsenenewsw
swseseeswseswswswseswswwswswswswswswsw
nwnenenewnenenwneenenwseseneneneswnewne
newneneneneneneneenenesenenenenwnenew
eeneeeeweeneneeeeeeeeswenwsw
nwesesweseseswsesewsesesesesesesesesese
wesweeweewswwnenweeneenesenene
swswnewswswwnweneseseseneseswwwseesw
wwswnewnenwswswswwewwwswswnewwse
neswseseseseseswseseseswswsesesw
neneneneneneswnenenenenenenenenwnene
wnwenwwnwnwnwnwwnwnwwnwnwwwnw
senwneswnenenwnwseneneswnenwnenewnwnenee
newwwnewwwwwwsewwswwsewneww
nwsenenenenwswnenenwsenenenenwnenenenene
eseeseeseseseeseweseseeseesesee
wswswwneewnwnwsewwwnenewseneswww
eeseeeeeeeeeeeneeeeeew
senwwneneswnesenwsewnenwnwnwnenwnwnew
sewneseseseneswswswswswsweseseswseswwsw
eseseeneseeewneenweeenwnwswenee
swnwneswweneeneewneswsenwesenwesee
neneswnenenwnwnenenenenenwswneneeswnene
eeseseeesewsesee
seseesenwswseseseswsese
neenwseeenwsewwwseesenesesesesesw
enweseneeeneneenewenewenwswese
wwwsewwwswewwswwwnwswwwswsw
swnwnwwseewnwnwwswnwwnwnwwnwwwnewse
seseseseeseseseseseseeswenwseesesesese
neneneneneneeneneenenenenewneeneee
weenwnenwwsweseseswnwsee
sesewwwnwwwnewwww
neewewneeneweeeeeeneenesenene
nwwnwwsewnwnwnwnwnwwnwnwwnwsenwnwnwnw
senwnwnwseseseseswwsenwsesenwseneesenw
wewswwswswswwswswwwwsww
nwseneseswwnwswnwseseenwewseeeenwsee
neneswnenenenenwnenesesenenesenenenwnenw
nwneneeeswnewsenweenenenweswne
eeenwneswseenewswnwneenwseneneeeee
neneneeeneseesenewneenwwneneneneswne
swswswswswswswswswswswswswswswswneswswsw
enweeeseeseeeenwswweeeeeee
wenwnewseenwnwneweenwnwnewnwnwnwnw
swswswsweswwwswswwnewswwswwwseww
nwwnwnweseswnwwswnenenwnwnwwswewnw
wnewwwwsewwwewwwwwwswww
swseswswswwswswswswseesweswswsesenwsw
sweswnwseeswswseswwswnwswswswswneswsw
wwswswseswwneneswseswswwswnwwswseww
sweswswswswseswswswswswswesweswwswnwwsw
sewnenwnenwswnwswenenewnwnenwseewnw
wnwwneenesewneeneenwnesewnewnwnenene
seswswswneneswswswseswswwswesenww
swswswwseneswswwswseswswswswesw
nwwnwnwwwnewwwsewnwnwnwnwnwsenenwnww
seseseswseswswseseseswseneeseseswswsewsw
eeneweseseneesesweseeeeseeseese
swseswswseseswswnenweseswswseswseswsenw
wwwwwwwwweswewwwwwwwnew
swseswswswswswswswseseseswswswswwneswswsw
wwwswwwwswwwwwwwwnewsesww
nenenewneneneneneneneeneswnenenene
nenenwneneneneneneewneneneneneswnwnene
nwneeneneneneeeneeeswneswenwswneswse
nesenwseeseswneseesenwsesesweneeesesw
swwswwswwwseswwwwnwwswsewneww
wwwwwwwsewwwnew
eenwesweeeeweneneeseneeewee
swsenwswnwwwswneenwwnwnwenwenenwnw
swnesweswnwswneenwsesweneneswnwnwnwnw
ewswwnwseswsenewnewwsweeneneneese
sweseneseswnewenenwnwwswseseseseswsene
nwnwnwnwwsenwnwwwnwnwnwnwwneewnwnw
nwnesenenwnwnenenwnwnenwwnenenenenwnene
nwnwswnwnwnwsenwnwnenwnenwnwnenwnenwnwnwnw
wsenwwenwnwwswwnenwwewnwnwnwwwse
wwswswnewwswwwwwwwswwweww
wswswwswswwswnwswswswswseswswswswsww
wwnwwewnwnwnwnwnwseeswwww
eeeeeeeeeesweeeeenwseeeew
eewneseenesenweneswseseeeweee
swwwswswwswswswswseswswwswnwswswwsw
nweeeeneneeeeseweweesweeeene
weewneweswnwwsenenenenenwesene
sesenwseseseseeweseeseswenwsenwsee
esweeeeeeswnenenwenweeeeee
wswswwnwswswwswewnwswswswswswswswewsw
sweswswswswswnwswseseswswseseswswswswsw
nwewseeseeeneswswseeweseeeneenw
eseswnwneneeneneeenwneneenee
swwswswwneswsewswnwnwwseeswwwsew
neesewweneseseswnwnewnweneswnwneswnee
nwseswseswswswswnwswsweswwswseseswswe
wnwwwsewnwenwwwswnwwnwnwwwwww
eseseseeneesesesenesewsesewsesesewse
swnwnwswnewwenwnwnwnwwwnwnenwnwwsww
nwnwnwnwnwnwnwenwnwnwnwwnenwnwsewsenew
neseswseswwseseswenwnwswseswseseesew
nwwnwnwnwnwswesweeenwnwswseeswnwnw
eeeeeeseeeeeesenweeeesese
nwenwwnenwnenenwnwnwnenwnwneneswnw
nwswswseswnwsesenwseseswswse
swseswneswsweseswseswwwswswswnwswswswsee
enwsesesesewseeseeeseseee
wenenesesesewnwneneswnwsweweneeswenw
sewwswswnewsewewswswwwswswnenesesww
seewseseeseseseseseneseseeeeeeee
nwnwnwnwnwenwswnwnwnwnwnwenwnwnwnwnww
swswswswswwsweneswswswnwswswswswswswswse
neneeneeneneneswnenenewnenenenenenene
nenenwewnenenwnene
sesesewseseseseseswseseseswnesesenesew
nwnwnwwnewseenwnenwswnwswnwneneneene
eseeeeseseenwesweeeenweeswee
neswwnewswsewswwseswswswswwwnwseswswsw
nwnwnwnwnenenwnenwnwneswneenwnwnwnwseswsw
wswwwwwswewswwwwnewsewswwnw
nwwneenenenwnwneenenwnenenesenenwnesw
wwewwnwwswewswwwneewwneeww
wswnweseseseseswseseswsesesenesesenese
wnwewnenwnwwnesenwnwnwnwsenwnwnwnwnwnwse
seenewswswswswsewswswswseseswswswswsese
eweseswseeeenweeeeeeeeeee
eswsenwneswswswswwnwswneswneseseenwnenw
nenwnwnenenwnenesenwnw
neneneneneneeswnenwneneseneneenwsw
nwwnwwwwnwwnwnwwnwwe
neeneswsesesesenesenesewseseswesenwsewse
seenwswsweeswswswswwwnwseswseswsee
sweswnewwwswnwnwewswwswwswswwww
eewneenwneeeeswneeseseeee
wewwwwwwwwwwewwwwwwww
nwsweneneneewnwnenwnenenwnwnwne
nenwneneeneswenenewneneseneswwnenenene
swnwseswseswswseswswseswswswseswswswsese
wnwwnwnwnwneenwnwnwnwnwnwnwenwnwnwnw
wwswnwnewenwwswswwnewwwwwwwne
neewnwnwwnenenesweneneeenwseneneswnwne
eneswnenwswnwnenwnwnwne
wswnenewneswwswwwseswneswswnewswnenw
eseeseeeseseenweeeenweeeeeswse
eswsenweneeeeeneeeweneeneneeene
nenewnesesenenenenewsenenwnwne
newwwnwwnwnwnwwwseswwwwnwnwnenwww
nwnwnwnwnwenwnenwnwnwwnwnwnwnw
swswswswswswswswnwseswswseswswswswswsw
nwnwwsenwnesenwnenwnwnwnwnwwswwnwww
senenwswswswwswswswswswswswswsewswswswsw
nwwwwwwwsewwwwwwnewwwe
wnenwswnenweswnewnwnwenwnwswnwnweesenw
nweseeeswenewseeeneeeweswneenwe
eswnwseswnwswswswewswswwswswswswsww
swsesesenwsenwswswsewneseseseenwsw
seseseswseseneseseseseswswseswseswneswse
nenenenenwsenwsenwswsewwwwwswnwenwse
neeneneneseneneneeeswnweneeee
eeneneeneneswneneenweneswnenenenewne
esenwwswwnwswewwwswswswswnwwwsw
wwwseseswnwnwneeenwwnewswnwswwwe
wswseswseswswswswnwswswsweneswswswsesw
nwswnwwwnwnwenwnw
swsesenwneseseseseseseseseseeseseesese
swsenewwsewnwwwwswewwneewwse
wwsweswnwnwnwsenesewswneneneseewne
nwnwnwnwnwnwnwswnwnwnwnwnwnenwnwsewnwnwne
swseswswsesewswsesweswswswswswswseswsw
nwnesewnenenenenesenenenenenenenene
nwnwnenwwnwnenwnwnwnwnwnwnwnenenwnwnwnwse
neeesweneneenee
eeseseseneswweeesesenweeesenweenw
wswwswswwwwwwwwewwwswweww
seseseesesewsesesewsesesesesesenesese
nwnenewseseneswneneweeswnenesenwewne
neseeneneneseenenenenenenenenwewnenee
seseseswseseseseseswesesenwswseseswsesw
swswsenenwewneseeeneeeneenesenwwe
sewwwnesesenwwnewwwwwwnww
nwseesweseswseswnwwsesweswsenwnesenenw
wwneseseeeseseeseseseseseneewseseee
nenenenwnenwnwnwneenwnewwneenenenwnw
sweeneeeeeeneeeeeeswneeee
wseeswwswwwwwsweswwnwswswswswww
eneneneeneswneneeneswneeneneneeee
enenenenenwnwwswnenenese
swnewnweseseewneneeswswswwnweww
seswsesesesesesesesesesesesenesesesesese
wswswswswseswnewswneswwwswswneswewsw
wnenenenenwnenwwenwnwsenwnwnwnenwnw
seseeeeeweeeeeeeenwseesesese
eeeesweeeeneeseeenweesewesene
swnewswneswwseneswwswneswswswenwswse
nweneeneneneneneneesenwswswenenesenww
seneneneeenenenenenenwneneneeeenene
swwswswswswswwwwswswswweswwswsww
nenwneswnenenwnenenenenesenenwneneenwnwne
swswswsweswswnwswswswswswswswswswswswsw
sesewnenwnenenwnenenwneenenenwnwnwswnenwnw
sesenwseseseseseseseseseeseswsesesesese
seseseesenwsenwseseswseseseseseswsewnese
wwswnwwswwwwwwwwwneewnwww
eeeseswsesesesenwnweseweese
neeeseeeweneeeeeneenwswnwnesw
wnwnwwwnwnwnwewswnenwwwwwswwnw
swnwseseswwswneseswneseswswswseseeswsw
eswnwnwswswwswswswswsweneswseswswneswswsw
nwseseeseseseseesesesesesee
seswseswswewwswnewnwswwwswnwswww
wwwwwnwnwnwnwnwewswnwwnwwnwnww
neseswnwswswseswseseseswseswseswseswswsesw`
