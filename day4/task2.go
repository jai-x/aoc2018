package main

import (
	"os"
	"fmt"
	"bufio"
	"time"
	"strings"
	"sort"
)

func die(v ...interface{}) {
	fmt.Println(v...)
	os.Exit(1)
}

type Event struct {
	ts time.Time
	log string
}

func NewEvent(line string) Event {
	parts := strings.SplitAfter(line, "]")
	ts, _ := time.Parse("[2006-01-02 15:04]", parts[0]) // magic date <('-'<)
	log := parts[1][1:]
	return Event{ts, log}
}

func (e Event) Before(other Event) bool {
	return e.ts.Before(other.ts)
}

func (e Event) String() string {
	return e.ts.Format(time.Stamp) + " " + e.log
}

func (e Event) Guard() string {
	if !strings.Contains(e.log, "Guard") {
		return ""
	}

	return strings.Split(e.log, " ")[1]
}

func (e Event) Minute() int {
	return e.ts.Minute()
}

func (e Event) Slept() bool {
	return strings.Contains(e.log, "falls")
}

func (e Event) Woke() bool {
	return strings.Contains(e.log, "wakes")
}

func solve(events []Event) {
	sleepsPerMinPerGuard := map[string][60]int{}

	onDuty  := ""
	sleptAt := 0

	for i := range events {
		// New guard
		guard := events[i].Guard()
		if len(guard) > 0 {
			onDuty = guard
			continue
		}

		if events[i].Slept() {
			sleptAt = events[i].Minute()
			continue
		}

		if events[i].Woke() {
			wokeAt := events[i].Minute()
			// mark minutes asleep
			for m := sleptAt; m < wokeAt; m++ {
				mins := sleepsPerMinPerGuard[onDuty]
				mins[m]++
				sleepsPerMinPerGuard[onDuty] = mins
			}
			continue
		}
	}

	maxSleepsPerGuard := 0
	atMinute := 0
	withGuard := ""

	for guard := range sleepsPerMinPerGuard {
		maxSleeps       := 0
		sleepiestMinute := 0

		for minute, sleeps := range sleepsPerMinPerGuard[guard] {
			if sleeps > maxSleeps {
				sleepiestMinute, maxSleeps = minute, sleeps
			}
		}

		if maxSleeps > maxSleepsPerGuard {
			maxSleepsPerGuard = maxSleeps
			atMinute = sleepiestMinute
			withGuard = guard
		}
	}

	fmt.Printf("Guard %v @ minute %v\n", withGuard, atMinute)

}
func main() {
	if len(os.Args) != 2 {
		die("Please specify input text file as only argument")
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		die(err)
	}
	defer file.Close()

	events := []Event{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		events = append(events, NewEvent(scanner.Text()))
	}

	// Sort events in chronological order
	sort.Slice(events, func(i, j int) bool { return events[i].Before(events[j]) })

	solve(events)
}
