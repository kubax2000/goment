package locales

import (
	"strconv"
	"strings"
)

// EnLocale is the US English language locale.
var EnLocale = newLocale(
	"en",
	strings.Split("Sunday_Monday_Tuesday_Wednesday_Thursday_Friday_Saturday", "_"),
	strings.Split("Sun_Mon_Tue_Wed_Thu_Fri_Sat", "_"),
	strings.Split("Su_Mo_Tu_We_Th_Fr_Sa", "_"),
	strings.Split("January_February_March_April_May_June_July_August_September_October_November_December", "_"),
	strings.Split("Jan_Feb_Mar_Apr_May_Jun_Jul_Aug_Sep_Oct_Nov_Dec", "_"),
	func(num int, period string) string {
		suffix := "th"
		switch num % 10 {
		case 1:
			if num%100 != 11 {
				suffix = "st"
			}
		case 2:
			if num%100 != 12 {
				suffix = "nd"
			}
		case 3:
			if num%100 != 13 {
				suffix = "rd"
			}
		}
		return strconv.Itoa(num) + suffix
	},
	nil,
	week{Dow: 0, Doy: 6},
	longDateFormats{
		"LTS":  "h:mm:ss A",
		"LT":   "h:mm A",
		"L":    "MM/DD/YYYY",
		"LL":   "MMMM D, YYYY",
		"LLL":  "MMMM D, YYYY h:mm A",
		"LLLL": "dddd, MMMM D, YYYY h:mm A",
	},
	relativeTimeFunctions{
		"future": func(number int, withoutSuffix bool, past bool) string {
			return "in %s"
		},
		"past": func(number int, withoutSuffix bool, past bool) string {
			return "%s ago"
		},
		"s": func(number int, withoutSuffix bool, past bool) string {
			return "a few seconds"
		},
		"ss": func(number int, withoutSuffix bool, past bool) string {
			return "%d seconds"
		},
		"m": func(number int, withoutSuffix bool, past bool) string {
			return "a minute"
		},
		"mm": func(number int, withoutSuffix bool, past bool) string {
			return "%d minutes"
		},
		"h": func(number int, withoutSuffix bool, past bool) string {
			return "an hour"
		},
		"hh": func(number int, withoutSuffix bool, past bool) string {
			return "%d hours"
		},
		"d": func(number int, withoutSuffix bool, past bool) string {
			return "a day"
		},
		"dd": func(number int, withoutSuffix bool, past bool) string {
			return "%d days"
		},
		"M": func(number int, withoutSuffix bool, past bool) string {
			return "a month"
		},
		"MM": func(number int, withoutSuffix bool, past bool) string {
			return "%d months"
		},
		"y": func(number int, withoutSuffix bool, past bool) string {
			return "a year"
		},
		"yy": func(number int, withoutSuffix bool, past bool) string {
			return "%d years"
		},
	},
	calendarFunctions{
		"sameDay": func(hours int, day int) string {
			return "[Today at] LT"
		},
		"nextDay": func(hours int, day int) string {
			return "[Tomorrow at] LT"
		},
		"nextWeek": func(hours int, day int) string {
			return "dddd [at] LT"
		},
		"lastDay": func(hours int, day int) string {
			return "[Yesterday at] LT"
		},
		"lastWeek": func(hours int, day int) string {
			return "[Last] dddd [at] LT"
		},
		"sameElse": func(hours int, day int) string {
			return "L"
		},
	},
	`(?i)(January|February|March|April|May|June|July|August|September|October|November|December)`,
	`(?i)(Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Oct|Nov|Dec)`,
	`(?i)(Sunday|Monday|Tuesday|Wednesday|Thursday|Friday|Saturday)`,
	`(?i)(Sun|Mon|Tue|Wed|Thu|Fri|Sat)`,
	`(?i)(Su|Mo|Tu|We|Th|Fr|Sa)`,
	`\d{1,2}(th|st|nd|rd)`,
)
