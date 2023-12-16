package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Boss struct {
	name          string
	time          time.Time
	bigBoss, done bool
}

func GetBossTimes() []Boss {
	f, _ := os.Open("times.txt")

	defer f.Close()

	boss := make([]Boss, 0)

	scanner := bufio.NewScanner(f)

	currentTime := time.Now()
	var timeString string
	layout := "15:04"

	var line string
	for scanner.Scan() {
		line = strings.TrimSpace(scanner.Text())
		splits := strings.Split(line, "   ")

		timeString = strings.Split(splits[0], " ")[0]
		bossTime, _ := time.Parse(layout, timeString)
		fullBossTime := time.Date(currentTime.Year(),
			currentTime.Month(),
			currentTime.Day(),
			bossTime.Hour(),
			bossTime.Minute(),
			0,
			0,
			currentTime.Location())

		for i := 1; i < len(splits); i++ {
			bigBoss := false
			if i > 1 {
				bigBoss = true
			}
			boss = append(boss, CreateBoss(splits[i], fullBossTime, bigBoss, false))
		}

	}
	return boss
}

func CreateBoss(name string, time time.Time, bigBoss, done bool) Boss {
	boss := Boss{
		name:    name,
		time:    time,
		done:    done,
		bigBoss: bigBoss,
	}
	return boss
}

func FilterForTime(input []Boss) []Boss {
	currentTime := time.Now()
	maxTime := currentTime.Add(15 * time.Minute)
	minTime := currentTime.Add(-15 * time.Minute)

	filtered := make([]Boss, 0)

	for _, v := range input {
		if minTime.Before(v.time) && maxTime.After(v.time) {
			fmt.Println(v, "is cool to put into filtered")
			fmt.Println(v.time)
			filtered = append(filtered, v)
		}
	}

	return filtered
}
