package vd
// ref: https://github.com/go-playground/validator/blob/059510c3bff4688135f150f0ed6e87d18539ca94/regexes.go

const (
	// alphaRegexString                 = "^[a-zA-Z]+$"
	// alphaNumericRegexString          = "^[a-zA-Z0-9]+$"
	// alphaUnicodeRegexString          = "^[\\p{L}]+$"
	// alphaUnicodeNumericRegexString   = "^[\\p{L}\\p{N}]+$"
	// numericRegexString               = "^[-+]?[0-9]+(?:\\.[0-9]+)?$"
	// numberRegexString                = "^[0-9]+$"
	hexRegexp          = "^[0-9a-fA-F]+$"
	hexcolorRegex             = "^#(?:[0-9a-fA-F]{3}|[0-9a-fA-F]{6})$"
	rgbRegex                  = "^rgb\\(\\s*(?:(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])|(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%)\\s*\\)$"
	rgbaRegex                 = "^rgba\\(\\s*(?:(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])|(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%)\\s*,\\s*(?:(?:0.[1-9]*)|[01])\\s*\\)$"
	hslRegex                  = "^hsl\\(\\s*(?:0|[1-9]\\d?|[12]\\d\\d|3[0-5]\\d|360)\\s*,\\s*(?:(?:0|[1-9]\\d?|100)%)\\s*,\\s*(?:(?:0|[1-9]\\d?|100)%)\\s*\\)$"
	hslaRegex                 = "^hsla\\(\\s*(?:0|[1-9]\\d?|[12]\\d\\d|3[0-5]\\d|360)\\s*,\\s*(?:(?:0|[1-9]\\d?|100)%)\\s*,\\s*(?:(?:0|[1-9]\\d?|100)%)\\s*,\\s*(?:(?:0.[1-9]*)|[01])\\s*\\)$"
	emailRegex                = "^(?:(?:(?:(?:[a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+(?:\\.([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+)*)|(?:(?:\\x22)(?:(?:(?:(?:\\x20|\\x09)*(?:\\x0d\\x0a))?(?:\\x20|\\x09)+)?(?:(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x7f]|\\x21|[\\x23-\\x5b]|[\\x5d-\\x7e]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[\\x01-\\x09\\x0b\\x0c\\x0d-\\x7f]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}]))))*(?:(?:(?:\\x20|\\x09)*(?:\\x0d\\x0a))?(\\x20|\\x09)+)?(?:\\x22))))@(?:(?:(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])(?:[a-zA-Z]|\\d|-|\\.|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.)+(?:(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])(?:[a-zA-Z]|\\d|-|\\.|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.?$"
	// e164Regex                 = "^\\+[1-9]?[0-9]{7,14}$"
	base64Regex               = "^(?:[A-Za-z0-9+\\/]{4})*(?:[A-Za-z0-9+\\/]{2}==|[A-Za-z0-9+\\/]{3}=|[A-Za-z0-9+\\/]{4})$"
	base64URLRegex            = "^(?:[A-Za-z0-9-_]{4})*(?:[A-Za-z0-9-_]{2}==|[A-Za-z0-9-_]{3}=|[A-Za-z0-9-_]{4})$"
	// iSBN10Regex               = "^(?:[0-9]{9}X|[0-9]{10})$"
	// iSBN13Regex               = "^(?:(?:97(?:8|9))[0-9]{10})$"
	// uUID3Regex                = "^[0-9a-f]{8}-[0-9a-f]{4}-3[0-9a-f]{3}-[0-9a-f]{4}-[0-9a-f]{12}$"
	// uUID4Regex                = "^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$"
	// uUID5Regex                = "^[0-9a-f]{8}-[0-9a-f]{4}-5[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$"
	uUIDRegex                 = "^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$"
	// uUID3RFC4122Regex         = "^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-3[0-9a-fA-F]{3}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$"
	// uUID4RFC4122Regex         = "^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$"
	// uUID5RFC4122Regex         = "^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-5[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$"
	// uUIDRFC4122Regex          = "^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$"
	aSCIIRegex                = "^[\x00-\x7F]*$"
	printableASCIIRegex       = "^[\x20-\x7E]*$"
	// multibyteRegex            = "[^\x00-\x7F]"
	dataURIRegex              = `^data:((?:\w+\/(?:([^;]|;[^;]).)+)?)`
	latitudeRegex             = "^[-+]?([1-8]?\\d(\\.\\d+)?|90(\\.0+)?)$"
	longitudeRegex            = "^[-+]?(180(\\.0+)?|((1[0-7]\\d)|([1-9]?\\d))(\\.\\d+)?)$"
	// sSNRegex                  = `^[0-9]{3}[ -]?(0[1-9]|[1-9][0-9])[ -]?([1-9][0-9]{3}|[0-9][1-9][0-9]{2}|[0-9]{2}[1-9][0-9]|[0-9]{3}[1-9])$`
	// hostnameRegexStringRFC952        = `^[a-zA-Z]([a-zA-Z0-9\-]+[\.]?)*[a-zA-Z0-9]$`                                                                      // https://tools.ietf.org/html/rfc952
	// hostnameRegexStringRFC1123       = `^([a-zA-Z0-9]{1}[a-zA-Z0-9_-]{0,62}){1}(\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*?$`                                 // accepts hostname starting with a digit https://tools.ietf.org/html/rfc1123
	// fqdnRegexStringRFC1123           = `^([a-zA-Z0-9]{1}[a-zA-Z0-9_-]{0,62})(\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*?(\.[a-zA-Z]{1}[a-zA-Z0-9]{0,62})\.?$` // same as hostnameRegexStringRFC1123 but must contain a non numerical TLD (possibly ending with '.')
	// btcAddressRegex           = `^[13][a-km-zA-HJ-NP-Z1-9]{25,34}$`                                                                                // bitcoin address
	// btcAddressUpperRegexStringBech32 = `^BC1[02-9AC-HJ-NP-Z]{7,76}$`                                                                                      // bitcoin bech32 address https://en.bitcoin.it/wiki/Bech32
	// btcAddressLowerRegexStringBech32 = `^bc1[02-9ac-hj-np-z]{7,76}$`                                                                                      // bitcoin bech32 address https://en.bitcoin.it/wiki/Bech32
	// ethAddressRegex           = `^0x[0-9a-fA-F]{40}$`
	// ethAddressUpperRegex      = `^0x[0-9A-F]{40}$`
	// ethAddressLowerRegex      = `^0x[0-9a-f]{40}$`
	// uRLEncodedRegex           = `(%[A-Fa-f0-9]{2})`
	// hTMLEncodedRegex          = `&#[x]?([0-9a-fA-F]{2})|(&gt)|(&lt)|(&quot)|(&amp)+[;]?`
	// hTMLRegex                 = `<[/]?([a-zA-Z]+).*?>`
	// splitParamsRegex          = `'[^']*'|\S+`
)


func Email() StringSpec {
	return StringSpec{
		Pattern: []string{emailRegex},
	}
}
func ChinaMobile() StringSpec {
	return StringSpec{
		MinRuneLen: 11,
		MaxRuneLen: 11,
	}
}
func UUID() StringSpec {
	return StringSpec{
		Pattern: []string{uUIDRegex},
	}
}
func Base64() StringSpec {
	return StringSpec{
		Pattern: []string{base64Regex},
	}
}
func Base64URL() StringSpec {
	return StringSpec{
		Pattern: []string{base64URLRegex},
	}
}
func DataURI() StringSpec {
	return StringSpec{
		Pattern: []string{dataURIRegex},
	}
}
func Latitude() StringSpec {
	return StringSpec{
		Pattern: []string{latitudeRegex},
	}
}
func Longitude() StringSpec {
	return StringSpec{
		Pattern: []string{longitudeRegex},
	}
}

func Hex() StringSpec {
	return StringSpec{
		Pattern: []string{hexRegexp},
	}
}
func HexColor() StringSpec {
	return StringSpec{
		Pattern: []string{hexcolorRegex},
	}
}
func RGB() StringSpec {
	return StringSpec{
		Pattern: []string{rgbRegex},
	}
}
func RGBA() StringSpec {
	return StringSpec{
		Pattern: []string{rgbaRegex},
	}
}
func HSL() StringSpec {
	return StringSpec{
		Pattern: []string{hslRegex},
	}
}
func HSLA() StringSpec {
	return StringSpec{
		Pattern: []string{hslaRegex},
	}
}
func ASCII() StringSpec {
	return StringSpec{
		Pattern: []string{aSCIIRegex},
	}
}
func PrintableASCII() StringSpec {
	return StringSpec{
		Pattern: []string{printableASCIIRegex},
	}
}
