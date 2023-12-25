package days

import (
	"bufio"
	"strconv"
	"strings"

	utils "github.com/chriscamarillo/advent-of-code-2023/common"
)

type Rule struct {
	destination_start int
	source_start      int
	length            int
}

func Day5Part1Solver() string {
	input := utils.FetchAOCInput(5)

	var seeds []int
	var seed_to_soil []Rule
	var soil_to_fertilizer []Rule
	var fertilizer_to_water []Rule
	var water_to_light []Rule
	var light_to_temperature []Rule
	var temperature_to_humidity []Rule
	var humidity_to_location []Rule

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "seeds:") {
			seeds = parseSeedLine(line)
		}
		if strings.HasPrefix(line, "seed-to-soil map:") {
			seed_to_soil = parseRules(scanner)
		}
		if strings.HasPrefix(line, "soil-to-fertilizer map:") {
			soil_to_fertilizer = parseRules(scanner)
		}
		if strings.HasPrefix(line, "fertilizer-to-water map:") {
			fertilizer_to_water = parseRules(scanner)
		}
		if strings.HasPrefix(line, "water-to-light map:") {
			water_to_light = parseRules(scanner)
		}
		if strings.HasPrefix(line, "light-to-temperature map:") {
			light_to_temperature = parseRules(scanner)
		}
		if strings.HasPrefix(line, "temperature-to-humidity map:") {
			temperature_to_humidity = parseRules(scanner)
		}
		if strings.HasPrefix(line, "humidity-to-location map:") {
			humidity_to_location = parseRules(scanner)
		}
	}

	smallest_location := 2147483647
	for _, seed := range seeds {
		soil := transform(seed, seed_to_soil)
		fertilizer := transform(soil, soil_to_fertilizer)
		water := transform(fertilizer, fertilizer_to_water)
		light := transform(water, water_to_light)
		temperature := transform(light, light_to_temperature)
		humidity := transform(temperature, temperature_to_humidity)
		location := transform(humidity, humidity_to_location)
		if location < smallest_location {
			smallest_location = location
		}
	}

	answer := strconv.Itoa(smallest_location)
	return answer
}

func transform(n int, rules []Rule) int {
	for _, rule := range rules {
		if withinBounds(n, rule) {
			return n - rule.source_start + rule.destination_start
		}
	}
	return n
}

func withinBounds(n int, rule Rule) bool {
	diff := n - rule.source_start
	return diff > 0 && diff < rule.length
}

func parseRules(scanner *bufio.Scanner) []Rule {
	var rules []Rule
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			return rules
		}
		fields := strings.Fields(line)
		destination_start, _ := strconv.Atoi(fields[0])
		source_start, _ := strconv.Atoi(fields[1])
		length, _ := strconv.Atoi(fields[2])

		rule := Rule {
			destination_start,
			source_start,
			length,
		}
		rules = append(rules, rule)
	}
	return rules
}

func parseSeedLine(line string) []int {
	seeds_strip := strings.TrimSpace(strings.Split(line, ":")[1])
	var seeds []int
	for _, seed := range strings.Split(seeds_strip, " ") {
		seed_value, _ := strconv.Atoi(seed)
		seeds = append(seeds, seed_value)
	}
	return seeds
}
