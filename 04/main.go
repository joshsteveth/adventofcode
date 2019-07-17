package main

import (
	"fmt"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/joshsteveth/adventofcode/util"
)

const filename string = "input.txt"

func main() {
	lines, err := util.ReadLines(filename)
	if err != nil {
		panic(err)
	}

	var activities []activity

	for _, l := range lines {
		var a activity
		if err := a.parseString(l, adjustTime); err != nil {
			panic(err)
		}
		activities = append(activities, a)
	}

	sort.Slice(activities, func(i, j int) bool {
		if activities[i].timestamp.Equal(activities[j].timestamp) {
			typI, typJ := activities[i].typ, activities[j].typ

			//first priority is beginShift
			if typI == beginShift {
				return true
			}

			if typI == fallsAsleep && typJ == wakesUp {
				return true
			}

			return false
		}

		return activities[i].timestamp.Before(activities[j].timestamp)
	})

	// guardMap represents how many hours a guard is sleeping
	guardMap := make(map[int]int)
	// minuteMap represents which minute which guard is sleeping the most
	minuteMap := make(map[int]map[int]int)
	for i := 0; i < 60; i++ {
		minuteMap[i] = make(map[int]int)
	}

	currentActivities := []activity{}
	for i, a := range activities {
		currentActivities = append(currentActivities, a)

		// if next activity is a begin shift, or it is the last activity..
		// then we need to calculate current activities, update guardMap, and set it back to empty
		if i == len(activities)-1 || activities[i+1].typ == beginShift {
			gid, sleeptime, listMinute, err := calculateActivities(currentActivities)
			if err != nil {
				log.Fatalf("error while processing activity %+v: %v", a, err)
			}

			guardMap[gid] += sleeptime
			for _, min := range listMinute {
				minuteMap[min][gid]++
			}

			currentActivities = []activity{}
			continue
		}
	}
	var (
		sleepiestGuardID int
		maxSleepTime     int
	)

	for id, tm := range guardMap {
		if tm > maxSleepTime {
			sleepiestGuardID = id
			maxSleepTime = tm
		}
	}

	fmt.Printf("sleepiest guard is %d with astonishing sleeping time %d\n",
		sleepiestGuardID, maxSleepTime)

	// PART 1
	// Strategy 1: Find the guard that has the most minutes asleep.
	// What minute does that guard spend asleep the most?

	// find out on which minute this guard sleeps the most
	maxMin := 0
	totalSleepTime := 0
	for min, gmap := range minuteMap {
		val := gmap[sleepiestGuardID]
		if val > totalSleepTime {
			maxMin = min
			totalSleepTime = val
		}
	}

	fmt.Printf("1st Part: sleepiest minute for this guard is %d (%d times), which makes the score: %d\n",
		maxMin, totalSleepTime, maxMin*sleepiestGuardID)

	// PART 2
	// Strategy 2: Of all guards, which guard is most frequently asleep on the same minute?
	// What is the ID of the guard you chose multiplied by the minute you chose?

	maxGuard, maxMin := 0, 0
	maxTimes := 0

	for min, gmap := range minuteMap {
		for gid, times := range gmap {
			if times > maxTimes {
				maxGuard, maxMin = gid, min
				maxTimes = times
			}
		}
	}

	fmt.Printf("2nd Part: Guard #%d sleeps the most on minute %d (%d times). score: %d\n",
		maxGuard, maxMin, maxTimes, maxGuard*maxMin)
}

// 3 types of activities
// [1518-11-01 00:00] Guard #10 begins shift
// [1518-11-01 00:05] falls asleep
// [1518-11-01 00:25] wakes up
type activityType int

type activity struct {
	timestamp time.Time
	typ       activityType
	guardID   int
}

const timeFormat string = "[2006-01-02 15:04]"

const (
	beginShift activityType = iota
	fallsAsleep
	wakesUp

	guard string = "Guard"
	falls        = "falls"
	wakes        = "wakes"
)

func (a activity) String() string {
	str := a.timestamp.Format(timeFormat)

	switch a.typ {
	case beginShift:
		str = fmt.Sprintf("%s Guard #%d begins shift", str, a.guardID)
	case fallsAsleep:
		str = fmt.Sprintf("%s falls asleep", str)
	case wakesUp:
		str = fmt.Sprintf("%s wakes up", str)
	}

	return str
}

var (
	dateRegex     = regexp.MustCompile(`\[(.*?)\]`)
	guardNumRegex = regexp.MustCompile(`#([0-9])*`)
)

func (a *activity) parseString(s string, options ...func(time.Time) time.Time) error {
	if err := a.parseTime(s, options...); err != nil {
		return err
	}

	if err := a.parseActivityType(s); err != nil {
		return err
	}

	return nil
}

func (a *activity) parseTime(s string, options ...func(time.Time) time.Time) error {
	t, err := time.Parse(timeFormat, dateRegex.FindString(s))
	if err != nil {
		return fmt.Errorf("Unable to parse time from %s: %v", s, err)
	}

	for _, opt := range options {
		t = opt(t)
	}

	a.timestamp = t
	return nil
}

// adjustTime adjusts time in following rule:
// since we only considers midnight, if the hour is not 00 then we add duration until it's 00 on the next day
func adjustTime(t time.Time) time.Time {
	if t.Hour() == 0 {
		return t
	}

	// let's calculate how many minute we need to go to the next day
	min := 60 - t.Minute()
	min += (23 - t.Hour()) * 60

	return t.Add(time.Duration(min) * time.Minute)
}

func (a *activity) parseActivityType(s string) error {
	if strings.Contains(s, guard) {
		guardNum, err := strconv.Atoi(strings.TrimPrefix(guardNumRegex.FindString(s), "#"))
		if err != nil {
			return fmt.Errorf("Unable to parse guard num from %s: %v", s, err)
		}
		a.guardID = guardNum
		a.typ = beginShift
		return nil
	}

	if strings.Contains(s, falls) {
		a.typ = fallsAsleep
		return nil
	}

	if strings.Contains(s, wakes) {
		a.typ = wakesUp
		return nil
	}

	return fmt.Errorf("Unable to determine activity type from %s", s)
}

func calculateActivities(as []activity) (guardID int, sleepTime int, listMinute []int, err error) {
	var startTime int

	for i, a := range as {
		// first one should be beginShift type
		if i == 0 {
			if a.typ != beginShift {
				err = fmt.Errorf("first activity should be a start shift type")
				return
			}

			guardID = a.guardID
			continue
		}

		// and the last one should be wakesUp type
		/*if i == len(as)-1 {
			if a.typ != wakesUp {
				err = fmt.Errorf("last activity should be waking up")
				return
			}
		}*/

		switch a.typ {
		case fallsAsleep:
			startTime = a.timestamp.Minute()
		case wakesUp:
			sleepTime += a.timestamp.Minute() - startTime

			for i := startTime; i < a.timestamp.Minute(); i++ {
				listMinute = append(listMinute, i)
			}
		default:
			err = fmt.Errorf("invalid acvitiy type %v for activity number %d", a.typ, i)
			return
		}
	}

	return guardID, sleepTime, listMinute, nil
}
