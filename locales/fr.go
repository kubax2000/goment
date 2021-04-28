package locales

import (
	"strconv"
	"strings"
)

// FrLocale is the French language locale.
var FrLocale = newLocale(
	"fr",
	strings.Split("dimanche_lundi_mardi_mercredi_jeudi_vendredi_samedi", "_"),
	strings.Split("dim._lun._mar._mer._jeu._ven._sam.", "_"),
	strings.Split("di_lu_ma_me_je_ve_sa", "_"),
	strings.Split("janvier_février_mars_avril_mai_juin_juillet_août_septembre_octobre_novembre_décembre", "_"),
	strings.Split("janv._févr._mars_avr._mai_juin_juil._août_sept._oct._nov._déc.", "_"),
	func(num int, period string) string {
		suffix := ""
		switch period {
		case "D":
			if num == 1 {
				suffix = "er"
			}
			return strconv.Itoa(num) + suffix
		case "M", "Q", "DDD", "d":
			suffix = "e"
			if num == 1 {
				suffix = "er"
			}
		case "w", "W":
			suffix = "e"
			if num == 1 {
				suffix = "re"
			}
		default:
			suffix = "e"
			if num == 1 {
				suffix = "er"
			}
		}
		return strconv.Itoa(num) + suffix
	},
	nil,
	week{Dow: 1, Doy: 4},
	longDateFormats{
		"LTS":  "HH:mm:ss",
		"LT":   "HH:mm",
		"L":    "DD/MM/YYYY",
		"LL":   "D MMMM YYYY",
		"LLL":  "D MMMM YYYY HH:mm",
		"LLLL": "dddd D MMMM YYYY HH:mm",
	},
	relativeTimeFunctions{
		"future": func(number int, withoutSuffix bool, past bool) string {
			return "dans %s"
		},
		"past": func(number int, withoutSuffix bool, past bool) string {
			return "il y a %s"
		},
		"s": func(number int, withoutSuffix bool, past bool) string {
			return "quelques secondes"
		},
		"ss": func(number int, withoutSuffix bool, past bool) string {
			return "%d secondes"
		},
		"m": func(number int, withoutSuffix bool, past bool) string {
			return "une minute"
		},
		"mm": func(number int, withoutSuffix bool, past bool) string {
			return "%d minutes"
		},
		"h": func(number int, withoutSuffix bool, past bool) string {
			return "une heure"
		},
		"hh": func(number int, withoutSuffix bool, past bool) string {
			return "%d heures"
		},
		"d": func(number int, withoutSuffix bool, past bool) string {
			return "un jour"
		},
		"dd": func(number int, withoutSuffix bool, past bool) string {
			return "%d jours"
		},
		"M": func(number int, withoutSuffix bool, past bool) string {
			return "un mois"
		},
		"MM": func(number int, withoutSuffix bool, past bool) string {
			return "%d mois"
		},
		"y": func(number int, withoutSuffix bool, past bool) string {
			return "un an"
		},
		"yy": func(number int, withoutSuffix bool, past bool) string {
			return "%d ans"
		},
	},
	calendarFunctions{
		"sameDay": func(hours int, day int) string {
			return "[Aujourd’hui à] LT"
		},
		"nextDay": func(hours int, day int) string {
			return "[Demain à] LT"
		},
		"nextWeek": func(hours int, day int) string {
			return "dddd [à] LT"
		},
		"lastDay": func(hours int, day int) string {
			return "[Hier à] LT"
		},
		"lastWeek": func(hours int, day int) string {
			return "dddd [dernier à] LT"
		},
		"sameElse": func(hours int, day int) string {
			return "L"
		},
	},
	`(?i)(janvier|février|mars|avril|mai|juin|juillet|août|septembre|octobre|novembre|décembre)`,
	`(?i)(janv\.?|févr\.?|mars|avr\.?|mai|juin|juil\.?|août|sept\.?|oct\.?|nov\.?|déc\.?)`,
	`(?i)(dimanche|lundi|mardi|mercredi|jeudi|vendredi|samedi)`,
	`(?i)(dim\.?|lun\.?|mar\.?|mer\.?|jeu\.?|ven\.?|sam\.?)`,
	`(?i)(di|lu|ma|me|je|ve|sa)`,
	`\d{1,2}(er|)`,
)
