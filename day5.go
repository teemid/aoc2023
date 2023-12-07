package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Map struct {
	Destination int
	Source      int
	Length      int
}

type Almanac struct {
	Seeds                 []int
	SeedToSoil            []*Map
	SoilToFertilizer      []*Map
	FertilizerToWater     []*Map
	WaterToLight          []*Map
	LightToTemperature    []*Map
	TemperatureToHumidity []*Map
	HumidityToLocation    []*Map
}

func day5(part string, filename string) {
	content, err := os.ReadFile(fmt.Sprintf("data/day5/%s.txt", filename))
	if err != nil {
		panic(err)
	}

	input := string(content)
	almanac := parseDay5Input(input)

	switch part {
	case "1":
		day5Part1(&almanac)
	case "2":
		day5Part2(&almanac)
	default:
		panic(fmt.Sprintf("unknown part: %s", part))
	}
}

func day5Part1(almanac *Almanac) {
	maps := [][]*Map{
		almanac.SeedToSoil,
		almanac.SoilToFertilizer,
		almanac.FertilizerToWater,
		almanac.WaterToLight,
		almanac.LightToTemperature,
		almanac.TemperatureToHumidity,
		almanac.HumidityToLocation,
	}

	mapped := make([]int, 0)

	for _, value := range almanac.Seeds {
		for _, m := range maps {
			value = mapValue(value, m)
		}

		mapped = append(mapped, value)
	}

	location := math.MaxInt
	for _, value := range mapped {
		if value < location {
			location = value
		}
	}

	fmt.Printf("location: %d\n", location)
}

type ByDestination []*Map

func (m ByDestination) Len() int           { return len(m) }
func (m ByDestination) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m ByDestination) Less(i, j int) bool { return m[i].Destination < m[j].Destination }

type BySource []*Map

func (m BySource) Len() int           { return len(m) }
func (m BySource) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m BySource) Less(i, j int) bool { return m[i].Source < m[j].Source }

func day5Part2(almanac *Almanac) {
	maps := [][]*Map{
		almanac.SeedToSoil,
		/*
		   almanac.SoilToFertilizer,
		   almanac.FertilizerToWater,
		   almanac.WaterToLight,
		   almanac.LightToTemperature,
		   almanac.TemperatureToHumidity,
		   almanac.HumidityToLocation,
		*/
	}

	sort.Sort(BySource(maps[0]))

	mapping := []*Map{{0, 0, math.MaxInt}}

	// First
	// 0, 0, 50
	// 50, 52, 48
	// 50, 98, 2
	// 100, 100, (MaxInt - 100)
	// Second
	// 39, 0, 15
	// 0, 15, 35
	// 50, 37, 2
	// 52, 54, 2
	// 98, 35, 2
	// 100, 100, (MaxInt - 100)

	newMapping := make([]*Map, 0)
	for _, layer := range maps {
		sort.Sort(ByDestination(mapping))
		sort.Sort(BySource(layer))

		fmt.Print("newMapping: ")
		printMapping(newMapping)
		fmt.Print("mapping: ")
		printMapping(mapping)
		fmt.Print("layer: ")
		printMapping(layer)

		var e *Map = mapping[0]
		mapping = mapping[1:]
		var m *Map = layer[0]
		layer = layer[1:]

		for e != nil && m != nil {
			fmt.Printf("newMapping: ")
			printMapping(newMapping)

			fmt.Printf("mapping: ")
			printMapping(mapping)

			fmt.Printf("layer: ")
			printMapping(layer)

			if e != nil && m != nil {
				if e.Destination < m.Source {
					left := &Map{
						Destination: e.Destination,
						Source:      e.Source,
						Length:      m.Source - e.Destination,
					}

					mid := &Map{
						Destination: m.Destination,
						Source:      m.Source,
						Length:      m.Length,
					}

					length := left.Length + mid.Length

					right := &Map{
						Destination: m.Destination + m.Length,
						Source:      e.Source + length,
						Length:      e.Length - length,
					}

					newMapping = append(newMapping, left)
					newMapping = append(newMapping, mid)

					if right.Length > 0 {
						mapping = append(mapping, right)
					}

					fmt.Printf("left: %+v\n", left)
					fmt.Printf("mid: %+v\n", mid)
					fmt.Printf("right: %+v\n", right)

					if len(mapping) > 0 {
						e = mapping[0]
						mapping = mapping[1:]
					} else {
						e = nil
					}

					if len(layer) > 0 {
						m = layer[0]
						layer = layer[1:]
					} else {
						m = nil
					}
				}
			} else if e != nil {

			} else if m != nil {

			} else {
				panic("Unreachable")
			}

			fmt.Print("newMapping: ")
			printMapping(newMapping)

			/*
			   // Line up source range with destination range
			   // If the destination range is within the source range
			   //    The destination range takes that part of the source range, and the source range is split
			   // If the destination range is outside the source range
			   //    The source range is kept as is
			   // If the destination range is partially within the source range
			   //    The source range is split

			   if isWithin(m, e) {
			       left := &Map{
			           Destination: e.Destination,
			           Source: e.Source,
			           Length: m.Source - e.Destination,
			       }

			       fmt.Printf("left: %+v\n", left)

			       combined := &Map{
			           Destination: m.Destination,
			           Source: m.Source,
			           Length: m.Length,
			       }

			       fmt.Printf("combined: %+v\n", combined)

			       length := left.Length + combined.Length

			       right := &Map{
			           Destination: m.Source + m.Length,
			           Source: e.Source + length,
			           Length: e.Length - length,
			       }

			       fmt.Printf("right: %+v\n", right)

			       if left.Length > 0 {
			           newMapping = append(newMapping, left)
			       }

			       newMapping = append(newMapping, combined)

			       mapping = append(mapping, right)
			   } else {
			   }
			*/
		}
	}

	fmt.Print("mapping: ")
	printMapping(mapping)
}

func isWithin(inner *Map, outer *Map) bool {
	return (inner.Source >= outer.Destination) && (inner.Source+inner.Length <= outer.Destination+outer.Length)
}

func split(a *Map, b *Map) {
	if a.Destination < b.Source {

	}
}

func isOverlapped(a *Map, b *Map) bool {
	return a.Destination >= b.Source
}

func printMapping(maps []*Map) {
	ms := make([]string, 0)
	for _, m := range maps {
		ms = append(ms, fmt.Sprintf("%+v", m))
	}

	fmt.Printf("[%s]\n", strings.Join(ms, ", "))
}

func mapValue(value int, maps []*Map) int {
	for _, m := range maps {
		if inRange(value, m.Source, m.Length) {
			return m.Destination + (value - m.Source)
		}
	}

	return value
}

func inRange(value int, start int, length int) bool {
	return value >= start && value < start+length
}

func parseDay5Input(input string) Almanac {
	almanac := Almanac{}

	segments := strings.Split(input, "\r\n\r\n")
	for idx, segment := range segments {
		if segment == "" {
			continue
		}

		switch idx {
		case 0: // seeds
			seeds := make([]int, 0)
			parts := strings.Split(segment, ": ")
			seedStrs := strings.Split(parts[1], " ")
			for _, seedStr := range seedStrs {
				seed, err := strconv.Atoi(seedStr)
				if err != nil {
					panic(err)
				}

				seeds = append(seeds, seed)
			}

			almanac.Seeds = seeds
		default: // maps
			maps := make([]*Map, 0)
			lines := strings.Split(segment, "\r\n")
			for _, line := range lines[1:] {
				if line == "" {
					continue
				}

				parts := strings.Split(line, " ")
				if len(parts) != 3 {
					panic(fmt.Sprintf("invalid map: %s", line))
				}

				var err error

				m := Map{}
				m.Destination, err = strconv.Atoi(parts[0])
				if err != nil {
					panic(err)
				}

				m.Source, err = strconv.Atoi(parts[1])
				if err != nil {
					panic(err)
				}

				m.Length, err = strconv.Atoi(parts[2])
				if err != nil {
					panic(err)
				}

				maps = append(maps, &m)
			}

			switch idx {
			case 1:
				almanac.SeedToSoil = maps
			case 2:
				almanac.SoilToFertilizer = maps
			case 3:
				almanac.FertilizerToWater = maps
			case 4:
				almanac.WaterToLight = maps
			case 5:
				almanac.LightToTemperature = maps
			case 6:
				almanac.TemperatureToHumidity = maps
			case 7:
				almanac.HumidityToLocation = maps
			default:
				panic(fmt.Sprintf("unknown segment: %d", idx))
			}
		}
	}

	return almanac
}
