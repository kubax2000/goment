package locales

import (
	"strings"
)

func inlineIf(condition bool, trueVal, falseVal string) string {
	if condition {
		return trueVal
	}
	return falseVal
}

func plural(n int) bool {
	return n > 1 && n < 5
}

// CsLocale is the Czech language locale.
var CsLocale = newLocale(
	"en",
	strings.Split("Neděle_Pondělí_Úterý_Středa_Čtvrtek_Pátek_Sobota", "_"),
	strings.Split("Ne_Po_Út_St_Čt_Pá_So", "_"),
	strings.Split("Ne_Po_Út_St_Čt_Pá_So", "_"),
	strings.Split("Leden_Únor_Březen_Duben_Květen_Červen_Červenec_Srpen_Září_Říjen_Listopad_Prosinec", "_"),
	strings.Split("Led_Úno_Bře_Dub_Kvě_Čvn_Čvc_Srp_Zář_Říj_Lis_Pro", "_"),
	nil,
	nil,
	week{Dow: 1, Doy: 4},
	longDateFormats{
		"LTS":  "H:mm:ss",
		"LT":   "H:mm",
		"L":    "DD.MM.YYYY",
		"LL":   "D. MMMM YYYY",
		"LLL":  "D. MMMM YYYY hh:mm",
		"LLLL": "dddd D. MMMM YYYY hh:mm",
		"l":    "D. M. YYYY",
	},
	relativeTimeFunctions{
		"future": func(number int, withoutSuffix bool, past bool) string {
			return "za %s"
		},
		"past": func(number int, withoutSuffix bool, past bool) string {
			return "před %s"
		},
		"s": func(number int, withoutSuffix bool, past bool) string {
			return inlineIf(withoutSuffix || !past, "pár sekund", "pár sekundami")
		},
		"ss": func(number int, withoutSuffix bool, past bool) string {
			if withoutSuffix || !past {
				return inlineIf(plural(number), "%d sekundy", "%d sekund")
			}
			return "%d sekundami"
		},
		"m": func(number int, withoutSuffix bool, past bool) string {
			return inlineIf(withoutSuffix, "minuta", inlineIf(!past, "minutu", "minutou"))
		},
		"mm": func(number int, withoutSuffix bool, past bool) string {
			if withoutSuffix || !past {
				return inlineIf(plural(number), "%d minuty", "%d minut")
			}
			return "%d minutami"
		},
		"h": func(number int, withoutSuffix bool, past bool) string {
			return inlineIf(withoutSuffix, "hodina", inlineIf(!past, "hodinu", "hodinou"))
		},
		"hh": func(number int, withoutSuffix bool, past bool) string {
			if withoutSuffix || !past {
				return inlineIf(plural(number), "%d hodiny", "%d hodin")
			}
			return "%d hodinami"
		},
		"d": func(number int, withoutSuffix bool, past bool) string {
			return inlineIf(withoutSuffix || !past, "den", "dnem")
		},
		"dd": func(number int, withoutSuffix bool, past bool) string {
			if withoutSuffix || !past {
				return inlineIf(plural(number), "%d dny", "%d dní")
			}
			return "%d dny"
		},
		"M": func(number int, withoutSuffix bool, past bool) string {
			return inlineIf(withoutSuffix || !past, "měsíc", "měsícem")
		},
		"MM": func(number int, withoutSuffix bool, past bool) string {
			if withoutSuffix || !past {
				return inlineIf(plural(number), "%d měsíce", "%d měsíců")
			}
			return "%d měsíci"
		},
		"y": func(number int, withoutSuffix bool, past bool) string {
			return inlineIf(withoutSuffix || !past, "rok", "rokem")
		},
		"yy": func(number int, withoutSuffix bool, past bool) string {
			if withoutSuffix || !past {
				return inlineIf(plural(number), "%d roky", "%d let")
			}
			return "%d lety"
		},
	},
	calendarFunctions{
		"sameDay": func(hours int, day int) string {
			return "[Dnes v] LT"
		},
		"nextDay": func(hours int, day int) string {
			return "[Zítra v] LT"
		},
		"nextWeek": func(hours int, day int) string {
			switch day {
			case 0:
				return "[V neděli v] LT"
			case 1:
				return "[v] dddd [v] LT"
			case 2:
				return "[v] dddd [v] LT"
			case 3:
				return "[Ve středu v] LT"
			case 4:
				return "[Ve čtvrtek v] LT"
			case 5:
				return "[V pátek v] LT"
			case 6:
				return "[V sobotu v] LT"
			}
			return "[v] dddd [v] LT"
		},
		"lastDay": func(hours int, day int) string {
			return "[Včera v] LT"
		},
		"lastWeek": func(hours int, day int) string {
			switch day {
			case 0:
				return "[Minulou neděli v] LT"
			case 1:
				return "[Minulé] dddd [v] LT"
			case 2:
				return "[Minulé] dddd [v] LT"
			case 3:
				return "[Minulou středu v] LT"
			case 4:
				return "[Minulý] dddd [v] LT"
			case 5:
				return "[Minulý] dddd [v] LT"
			case 6:
				return "[Minulou sobotu v] LT"
			}
			return "[Minulé] dddd [v] LT"
		},
		"sameElse": func(hours int, day int) string {
			return "L"
		},
	},
	`(?i)(Leden|Ledna|Února|Únor|Březen|Března|Duben|Dubna|Květen|Května|Červenec|Července|Červen|Června|Srpen|Srpna|Září|Říjen|Října|Listopadu|Listopad|Prosinec|Prosince)`,
	`(?i)(Led|Úno|Bře|Dub|Kvě|Čvn|Čvc|Srp|Zář|Říj|Lis|Pro)`,
	`(?i)(Neděle|Pondělí|Úterý|Středa|Čtvrtek|Pátek|Sobota)`,
	`(?i)(Ne|Po|Út|St|Čt|Pá|So)`,
	`(?i)(Ne|Po|Út|St|Čt|Pá|So)`,
	`\d{1,2}`,
)
