package main

import (
    "fmt"
    "math"
    "os"
    "strings"
    "strconv"
)

type Map struct {
    Destination int
    Source int
    Length int
}

type Almanac struct {
    Seeds []int
    SeedToSoil []Map
    SoilToFertilizer []Map
    FertilizerToWater []Map
    WaterToLight []Map
    LightToTemperature []Map
    TemperatureToHumidity []Map
    HumidityToLocation []Map
}

func day5(part string, filename string) {
    content, err := os.ReadFile(fmt.Sprintf("data/day5/%s.txt", filename))
    if err != nil {
        panic(err)
    }

    input := string(content)
    almanac := parseDay5Input(input)

    switch part {
    case "1": day5Part1(&almanac)
    case "2": day5Part2(&almanac)
    default: panic(fmt.Sprintf("unknown part: %s", part))
    }
}

func day5Part1(almanac *Almanac) {
    maps := [][]Map{
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

func day5Part2(almanac *Almanac) {
    maps := [][]Map{
        almanac.SeedToSoil,
        almanac.SoilToFertilizer,
        almanac.FertilizerToWater,
        almanac.WaterToLight,
        almanac.LightToTemperature,
        almanac.TemperatureToHumidity,
        almanac.HumidityToLocation,
    }

    mapped := make([]int, 0)
    for i := 0; i < len(almanac.Seeds); i += 2 {
        start := almanac.Seeds[i]
        length := almanac.Seeds[i + 1]

        fmt.Printf("start: %d, length: %d\n", start, length)

        for seed := start; seed < (start + length); seed += 1 {
            value := seed

            for _, m := range maps {
                value = mapValue(value, m) 
            }

            mapped = append(mapped, value)
        }
    }

    location := math.MaxInt
    for _, value := range mapped {
        if value < location {
            location = value
        }
    }

    fmt.Printf("location: %d\n", location)
}

func mapValue(value int, maps []Map) int {
    for _, m := range maps {
        if inRange(value, m.Source, m.Length) {
            return m.Destination + (value - m.Source)
        }
    }

    return value
}

func inRange(value int, start int, length int) bool {
    return value >= start && value < start + length
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
            maps := make([]Map, 0)
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

                maps = append(maps, m)
            }

            switch idx {
            case 1: almanac.SeedToSoil = maps
            case 2: almanac.SoilToFertilizer = maps
            case 3: almanac.FertilizerToWater = maps
            case 4: almanac.WaterToLight = maps
            case 5: almanac.LightToTemperature = maps
            case 6: almanac.TemperatureToHumidity = maps
            case 7: almanac.HumidityToLocation = maps
            default: panic(fmt.Sprintf("unknown segment: %d", idx))
            }
        }
    }

    return almanac
}

