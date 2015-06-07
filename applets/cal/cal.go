// cal - print calendar
package cal

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

func parseMonth(month string) time.Month {
	switch month {
	case "jan", "january":
		return time.January
	case "feb", "february":
		return time.February
	case "mar", "march":
		return time.March
	case "apr", "april":
		return time.April
	case "may":
		return time.May
	case "jun", "june":
		return time.June
	case "jul", "july":
		return time.July
	case "aug", "august":
		return time.August
	case "sep", "september":
		return time.September
	case "oct", "october":
		return time.October
	case "nov", "november":
		return time.November
	case "dec", "december":
		return time.December
	}
	return time.Month(0)
}

func parseYear(year string) int {
	yr, err := strconv.ParseInt(year, 10, 32)

	if err != nil {
		fmt.Fprintln(os.Stderr, "usage: cal [ month ] [ year ]")
		return 0
	}

	return int(yr)
}

func januaryFirst(year int) int {
	d := 4 + year + (year+3)/4

	if year > 1800 {
		d -= (year - 1701) / 100
		d += (year - 1601) / 400
	}

	if year > 1752 {
		d += 3
	}

	return int(d % 7)
}

var (
	flagSet     = flag.NewFlagSet("cal", flag.PanicOnError)
	nonleapyear = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	leapyear    = []int{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
)

func months(year int) []int {
	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		return leapyear
	}
	return nonleapyear
}

func printCal(month time.Month, year int) {
	dayw := "Su Mo Tu We Th Fr Sa"

	s := month.String() + " " + strconv.FormatInt(int64(year), 10)

	off := (20 - len(s)) / 2

	for i := 0; i < off; i++ {
		s = " " + s
	}

	s += "\n" + dayw + "\n"

	mth := months(year)

	day := januaryFirst(year)

	for i := 1; i < int(month); i++ {
		day += mth[i-1]
	}

	for i := 0; i < day%7; i++ {
		s = s + "   "
	}

	for i := 1; i <= mth[month-1]; i++ {
		s = s + " "
		if i < 10 {
			s = s + " "
		}
		s = s + strconv.Itoa(i)
		day += 1
		if day%7 == 0 {
			s = s + "\n"
		}
	}

	s = s + "\n\n"

	fmt.Print(s)
}

func Cal(call []string) error {
	flagSet.Parse(call[1:])

	local := time.Now()

	var month time.Month
	var year int

	if flagSet.NArg() > 2 {
		fmt.Fprintln(os.Stderr, "usage: cal [ month ] [ year ]")
		return nil
	}

	if flagSet.NArg() == 2 {
		month = parseMonth(flagSet.Arg(0))
		year = parseYear(flagSet.Arg(1))

		if month == 0 {
			fmt.Fprintln(os.Stderr, "usage: cal [ month ] [ year ]")
			return nil
		}
	} else {
		if flagSet.NArg() == 1 {
			year = local.Year()
			month = parseMonth(flagSet.Arg(0))

			if month == time.Month(0) {
				year = parseYear(flagSet.Arg(0))
				month = local.Month()
			}
		} else {
			year = local.Year()
			month = local.Month()
		}
	}

	printCal(month, year)

	return nil
}
