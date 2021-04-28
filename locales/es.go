package locales

import (
	"fmt"
	"strings"
)

func getEsCalendarPronoun(hours int) string {
	if hours != 1 {
		return "las"
	}
	return "la"
}

// EsLocale is the Spanish language locale.
var EsLocale = newLocale(
	"es",
	strings.Split("domingo_lunes_martes_miércoles_jueves_viernes_sábado", "_"),
	strings.Split("dom._lun._mar._mié._jue._vie._sáb.", "_"),
	strings.Split("do_lu_ma_mi_ju_vi_sá", "_"),
	strings.Split("enero_febrero_marzo_abril_mayo_junio_julio_agosto_septiembre_octubre_noviembre_diciembre", "_"),
	strings.Split("ene_feb_mar_abr_may_jun_jul_ago_sep_oct_nov_dic", "_"),
	func(num int, period string) string {
		return fmt.Sprintf("%dº", num)
	},
	nil,
	week{Dow: 1, Doy: 4},
	longDateFormats{
		"LTS":  "H:mm:ss",
		"LT":   "H:mm",
		"L":    "DD/MM/YYYY",
		"LL":   "D [de] MMMM [de] YYYY",
		"LLL":  "D [de] MMMM [de] YYYY H:mm",
		"LLLL": "dddd, D [de] MMMM [de] YYYY H:mm",
	},
	relativeTimeFunctions{
		"future": func(number int, withoutSuffix bool, past bool) string {
			return "en %s"
		},
		"past": func(number int, withoutSuffix bool, past bool) string {
			return "hace %s"
		},
		"s": func(number int, withoutSuffix bool, past bool) string {
			return "unos segundos"
		},
		"ss": func(number int, withoutSuffix bool, past bool) string {
			return "%d segundos"
		},
		"m": func(number int, withoutSuffix bool, past bool) string {
			return "un minuto"
		},
		"mm": func(number int, withoutSuffix bool, past bool) string {
			return "%d minutos"
		},
		"h": func(number int, withoutSuffix bool, past bool) string {
			return "una hora"
		},
		"hh": func(number int, withoutSuffix bool, past bool) string {
			return "%d horas"
		},
		"d": func(number int, withoutSuffix bool, past bool) string {
			return "un día"
		},
		"dd": func(number int, withoutSuffix bool, past bool) string {
			return "%d días"
		},
		"M": func(number int, withoutSuffix bool, past bool) string {
			return "un mes"
		},
		"MM": func(number int, withoutSuffix bool, past bool) string {
			return "%d meses"
		},
		"y": func(number int, withoutSuffix bool, past bool) string {
			return "un año"
		},
		"yy": func(number int, withoutSuffix bool, past bool) string {
			return "%d años"
		},
	},
	calendarFunctions{
		"sameDay": func(hours int, day int) string {
			return "[hoy a " + getEsCalendarPronoun(hours) + "] LT"
		},
		"nextDay": func(hours int, day int) string {
			return "[mañana a " + getEsCalendarPronoun(hours) + "] LT"
		},
		"nextWeek": func(hours int, day int) string {
			return "dddd [a " + getEsCalendarPronoun(hours) + "] LT"
		},
		"lastDay": func(hours int, day int) string {
			return "[ayer a " + getEsCalendarPronoun(hours) + "] LT"
		},
		"lastWeek": func(hours int, day int) string {
			return "[el] dddd [pasado a " + getEsCalendarPronoun(hours) + "] LT"
		},
		"sameElse": func(hours int, day int) string {
			return "L"
		},
	},
	`(?i)(enero|febrero|marzo|abril|mayo|junio|julio|agosto|septiembre|octubre|noviembre|diciembre)`,
	`(?i)(ene\.?|feb\.?|mar\.?|abr\.?|may\.?|jun\.?|jul\.?|ago\.?|sep\.?|oct\.?|nov\.?|dic\.?)`,
	`(?i)(domingo|lunes|martes|miércoles|jueves|viernes|sábado)`,
	`(?i)(dom\.?|lun\.?|mar\.?|mié\.?|jue\.?|vie\.?|sáb\.?)`,
	`(?i)(do|lu|ma|mi|ju|vi|sá)`,
	`\d{1,2}º`,
)
