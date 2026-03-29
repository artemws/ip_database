package main

// CodesMainlands — допустимые коды материков.
var CodesMainlands = map[string]struct{}{
	"AF": {}, "AN": {}, "AS": {}, "EU": {},
	"NA": {}, "OC": {}, "SA": {}, "ZZ": {},
}

// CodesCountries — допустимые коды стран.
var CodesCountries = map[string]struct{}{
	"AF": {}, "AL": {}, "DZ": {}, "AS": {}, "AD": {}, "AO": {}, "AI": {}, "AQ": {},
	"AG": {}, "AR": {}, "AM": {}, "AW": {}, "AU": {}, "AT": {}, "AZ": {}, "BS": {},
	"BH": {}, "BD": {}, "BB": {}, "BY": {}, "BE": {}, "BZ": {}, "BJ": {}, "BM": {},
	"BT": {}, "BO": {}, "BA": {}, "BW": {}, "BR": {}, "IO": {}, "VG": {}, "BN": {},
	"BG": {}, "BF": {}, "BI": {}, "KH": {}, "CM": {}, "CA": {}, "CV": {}, "KY": {},
	"CF": {}, "TD": {}, "CL": {}, "CN": {}, "CC": {}, "CO": {}, "KM": {}, "CK": {},
	"CR": {}, "HR": {}, "CU": {}, "CW": {}, "CY": {}, "CZ": {}, "CD": {}, "DK": {},
	"DJ": {}, "DM": {}, "DO": {}, "TL": {}, "EC": {}, "EG": {}, "SV": {}, "GQ": {},
	"ER": {}, "EE": {}, "ET": {}, "FK": {}, "FO": {}, "FJ": {}, "FI": {}, "FR": {},
	"PF": {}, "GA": {}, "GM": {}, "GE": {}, "DE": {}, "GH": {}, "GI": {}, "GR": {},
	"GL": {}, "GD": {}, "GU": {}, "GT": {}, "GG": {}, "GN": {}, "GW": {}, "GY": {},
	"HT": {}, "HN": {}, "HK": {}, "HU": {}, "IS": {}, "IN": {}, "ID": {}, "IR": {},
	"IQ": {}, "IE": {}, "IM": {}, "IL": {}, "IT": {}, "CI": {}, "JM": {}, "JP": {},
	"JE": {}, "JO": {}, "KZ": {}, "KE": {}, "KI": {}, "KW": {}, "KG": {}, "LA": {},
	"LV": {}, "LB": {}, "LS": {}, "LR": {}, "LY": {}, "LI": {}, "LT": {}, "LU": {},
	"MO": {}, "MK": {}, "MG": {}, "MW": {}, "MY": {}, "MV": {}, "ML": {}, "MT": {},
	"MH": {}, "MR": {}, "MU": {}, "YT": {}, "MX": {}, "FM": {}, "MD": {}, "MC": {},
	"MN": {}, "ME": {}, "MS": {}, "MA": {}, "MZ": {}, "MM": {}, "NA": {}, "NR": {},
	"NP": {}, "NL": {}, "NC": {}, "NZ": {}, "NI": {}, "NE": {}, "NG": {}, "NU": {},
	"KP": {}, "MP": {}, "NO": {}, "OM": {}, "PK": {}, "PW": {}, "PS": {}, "PA": {},
	"PG": {}, "PY": {}, "PE": {}, "PH": {}, "PL": {}, "PT": {}, "PR": {}, "QA": {},
	"CG": {}, "RE": {}, "RO": {}, "RU": {}, "RW": {}, "BL": {}, "KN": {}, "LC": {},
	"MF": {}, "PM": {}, "VC": {}, "WS": {}, "SM": {}, "ST": {}, "SA": {}, "SN": {},
	"RS": {}, "SC": {}, "SL": {}, "SG": {}, "SX": {}, "SK": {}, "SI": {}, "SB": {},
	"SO": {}, "ZA": {}, "KR": {}, "SS": {}, "ES": {}, "LK": {}, "SD": {}, "SR": {},
	"SZ": {}, "SE": {}, "CH": {}, "SY": {}, "TW": {}, "TJ": {}, "TZ": {}, "TH": {},
	"TG": {}, "TK": {}, "TO": {}, "TT": {}, "TN": {}, "TR": {}, "TM": {}, "TC": {},
	"TV": {}, "VI": {}, "UG": {}, "UA": {}, "AE": {}, "GB": {}, "US": {}, "UY": {},
	"UZ": {}, "VU": {}, "VA": {}, "VE": {}, "VN": {}, "WF": {}, "YE": {}, "ZM": {},
	"ZW": {},
}

// MainlandCodeData — пары [Название, Код] для материков.
var MainlandCodeData = []string{
	"Europe", "EU",
	"Africa", "AF",
	"Antarctica", "AN",
	"Asia", "AS",
	"North_America", "NA",
	"Oceania", "OC",
	"South_America", "SA",
	"Unknown_or_unspecified", "ZZ",
}

// CountryCodeData — пары [Название, Код] для стран.
var CountryCodeData = []string{
	"Afghanistan", "AF", "Albania", "AL", "Algeria", "DZ", "American_Samoa", "AS",
	"Andorra", "AD", "Angola", "AO", "Anguilla", "AI", "Antarctica", "AQ",
	"Antigua_and_Barbuda", "AG", "Argentina", "AR", "Armenia", "AM", "Aruba", "AW",
	"Australia", "AU", "Austria", "AT", "Azerbaijan", "AZ", "Bahamas", "BS",
	"Bahrain", "BH", "Bangladesh", "BD", "Barbados", "BB", "Belarus", "BY",
	"Belgium", "BE", "Belize", "BZ", "Benin", "BJ", "Bermuda", "BM",
	"Bhutan", "BT", "Bolivia", "BO", "Bosnia_and_Herzegovina", "BA", "Botswana", "BW",
	"Brazil", "BR", "British_Indian_Ocean_Territory", "IO", "British_Virgin_Islands", "VG",
	"Brunei", "BN", "Bulgaria", "BG", "Burkina_Faso", "BF", "Burundi", "BI",
	"Cambodia", "KH", "Cameroon", "CM", "Canada", "CA", "Cape_Verde", "CV",
	"Cayman_Islands", "KY", "Central_African_Republic", "CF", "Chad", "TD",
	"Chile", "CL", "China", "CN", "Cocos_Islands", "CC", "Colombia", "CO",
	"Comoros", "KM", "Cook_Islands", "CK", "Costa_Rica", "CR", "Croatia", "HR",
	"Cuba", "CU", "Curacao", "CW", "Cyprus", "CY", "Czech_Republic", "CZ",
	"Democratic_Republic_of_the_Congo", "CD", "Denmark", "DK", "Djibouti", "DJ",
	"Dominica", "DM", "Dominican_Republic", "DO", "East_Timor", "TL", "Ecuador", "EC",
	"Egypt", "EG", "El_Salvador", "SV", "Equatorial_Guinea", "GQ", "Eritrea", "ER",
	"Estonia", "EE", "Ethiopia", "ET", "Falkland_Islands", "FK", "Faroe_Islands", "FO",
	"Fiji", "FJ", "Finland", "FI", "France", "FR", "French_Polynesia", "PF",
	"Gabon", "GA", "Gambia", "GM", "Georgia", "GE", "Germany", "DE",
	"Ghana", "GH", "Gibraltar", "GI", "Greece", "GR", "Greenland", "GL",
	"Grenada", "GD", "Guam", "GU", "Guatemala", "GT", "Guernsey", "GG",
	"Guinea", "GN", "Guinea_Bissau", "GW", "Guyana", "GY", "Haiti", "HT",
	"Honduras", "HN", "Hong_Kong", "HK", "Hungary", "HU", "Iceland", "IS",
	"India", "IN", "Indonesia", "ID", "Iran", "IR", "Iraq", "IQ",
	"Ireland", "IE", "Isle_of_Man", "IM", "Israel", "IL", "Italy", "IT",
	"Ivory_Coast", "CI", "Jamaica", "JM", "Japan", "JP", "Jersey", "JE",
	"Jordan", "JO", "Kazakhstan", "KZ", "Kenya", "KE", "Kiribati", "KI",
	"Kuwait", "KW", "Kyrgyzstan", "KG", "Laos", "LA", "Latvia", "LV",
	"Lebanon", "LB", "Lesotho", "LS", "Liberia", "LR", "Libya", "LY",
	"Liechtenstein", "LI", "Lithuania", "LT", "Luxembourg", "LU", "Macau", "MO",
	"Macedonia", "MK", "Madagascar", "MG", "Malawi", "MW", "Malaysia", "MY",
	"Maldives", "MV", "Mali", "ML", "Malta", "MT", "Marshall_Islands", "MH",
	"Mauritania", "MR", "Mauritius", "MU", "Mayotte", "YT", "Mexico", "MX",
	"Micronesia", "FM", "Moldova", "MD", "Monaco", "MC", "Mongolia", "MN",
	"Montenegro", "ME", "Montserrat", "MS", "Morocco", "MA", "Mozambique", "MZ",
	"Myanmar", "MM", "Namibia", "NA", "Nauru", "NR", "Nepal", "NP",
	"Netherlands", "NL", "New_Caledonia", "NC", "New_Zealand", "NZ", "Nicaragua", "NI",
	"Niger", "NE", "Nigeria", "NG", "Niue", "NU", "North_Korea", "KP",
	"Northern_Mariana_Islands", "MP", "Norway", "NO", "Oman", "OM", "Pakistan", "PK",
	"Palau", "PW", "Palestine", "PS", "Panama", "PA", "Papua_New_Guinea", "PG",
	"Paraguay", "PY", "Peru", "PE", "Philippines", "PH", "Poland", "PL",
	"Portugal", "PT", "Puerto_Rico", "PR", "Qatar", "QA", "Republic_of_the_Congo", "CG",
	"Reunion", "RE", "Romania", "RO", "Russia", "RU", "Rwanda", "RW",
	"Saint_Barthelemy", "BL", "Saint_Kitts_and_Nevis", "KN", "Saint_Lucia", "LC",
	"Saint_Martin", "MF", "Saint_Pierre_and_Miquelon", "PM",
	"Saint_Vincent_and_the_Grenadines", "VC", "Samoa", "WS", "San_Marino", "SM",
	"Sao_Tome_and_Principe", "ST", "Saudi_Arabia", "SA", "Senegal", "SN",
	"Serbia", "RS", "Seychelles", "SC", "Sierra_Leone", "SL", "Singapore", "SG",
	"Sint_Maarten", "SX", "Slovakia", "SK", "Slovenia", "SI", "Solomon_Islands", "SB",
	"Somalia", "SO", "South_Africa", "ZA", "South_Korea", "KR", "South_Sudan", "SS",
	"Spain", "ES", "Sri_Lanka", "LK", "Sudan", "SD", "Suriname", "SR",
	"Swaziland", "SZ", "Sweden", "SE", "Switzerland", "CH", "Syria", "SY",
	"Taiwan", "TW", "Tajikistan", "TJ", "Tanzania", "TZ", "Thailand", "TH",
	"Togo", "TG", "Tokelau", "TK", "Tonga", "TO", "Trinidad_and_Tobago", "TT",
	"Tunisia", "TN", "Turkey", "TR", "Turkmenistan", "TM", "Turks_and_Caicos_Islands", "TC",
	"Tuvalu", "TV", "U.S._Virgin_Islands", "VI", "Uganda", "UG", "Ukraine", "UA",
	"United_Arab_Emirates", "AE", "United_Kingdom", "GB", "United_States", "US",
	"Uruguay", "UY", "Uzbekistan", "UZ", "Vanuatu", "VU", "Vatican", "VA",
	"Venezuela", "VE", "Vietnam", "VN", "Wallis_and_Futuna", "WF",
	"Yemen", "YE", "Zambia", "ZM", "Zimbabwe", "ZW",
}

// MapCode возвращает map[код]→название из плоского слайса пар [Название, Код].
func MapCode(data []string) map[string]string {
	m := make(map[string]string, len(data)/2)
	for i := 0; i+1 < len(data); i += 2 {
		name, code := data[i], data[i+1]
		m[code] = name
	}
	return m
}

// PrintMap выводит map в формате "Название --> Код".
func PrintMap(m map[string]string) {
	// Для воспроизводимого вывода сортируем ключи.
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	// Простая сортировка пузырьком (список небольшой, stdlib sort не нужен)
	for i := 0; i < len(keys); i++ {
		for j := i + 1; j < len(keys); j++ {
			if keys[i] > keys[j] {
				keys[i], keys[j] = keys[j], keys[i]
			}
		}
	}
	for _, k := range keys {
		println(m[k] + " --> " + k)
	}
}
