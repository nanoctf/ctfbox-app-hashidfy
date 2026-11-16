// Code generated from the supplied Python detector list.
// These are shape-based candidates, not cryptographic verification.
package detector

type hashRule struct {
	name   string
	length int
	id     string
	rule   string
}

var hashRules = []hashRule{
	{name: "CRC16", length: 4, id: "101020", rule: "notalpha"},
	{name: "CRC16CCITT", length: 4, id: "101040", rule: "mixed"},
	{name: "FCS16", length: 4, id: "101060", rule: "mixed"},
	{name: "CRC32", length: 8, id: "102040", rule: "mixed"},
	{name: "ADLER32", length: 8, id: "102020", rule: "mixed"},
	{name: "CRC32B", length: 8, id: "102060", rule: "mixed"},
	{name: "XOR32", length: 8, id: "102080", rule: "mixed"},
	{name: "GHash323", length: 8, id: "103040", rule: "digits"},
	{name: "GHash325", length: 8, id: "103020", rule: "digits"},
	{name: "DESUnix", length: 13, id: "104020", rule: "nondigit_nonalpha"},
	{name: "MD5Half", length: 16, id: "105060", rule: "mixed"},
	{name: "MD5Middle", length: 16, id: "105040", rule: "mixed"},
	{name: "MySQL", length: 16, id: "105020", rule: "mixed"},
	{name: "DomainCachedCredentials", length: 32, id: "106025", rule: "mixed"},
	{name: "Haval128", length: 32, id: "106160", rule: "mixed"},
	{name: "Haval128HMAC", length: 32, id: "106165", rule: "mixed"},
	{name: "MD2", length: 32, id: "106060", rule: "mixed"},
	{name: "MD2HMAC", length: 32, id: "106120", rule: "mixed"},
	{name: "MD4", length: 32, id: "106040", rule: "mixed"},
	{name: "MD4HMAC", length: 32, id: "106100", rule: "mixed"},
	{name: "MD5", length: 32, id: "106020", rule: "mixed"},
	{name: "MD5HMAC", length: 32, id: "106080", rule: "mixed"},
	{name: "MD5HMACWordpress", length: 32, id: "106140", rule: "mixed"},
	{name: "NTLM", length: 32, id: "106029", rule: "mixed"},
	{name: "RAdminv2x", length: 32, id: "106027", rule: "mixed"},
	{name: "RipeMD128", length: 32, id: "106180", rule: "mixed"},
	{name: "RipeMD128HMAC", length: 32, id: "106185", rule: "mixed"},
	{name: "SNEFRU128", length: 32, id: "106200", rule: "mixed"},
	{name: "SNEFRU128HMAC", length: 32, id: "106205", rule: "mixed"},
	{name: "Tiger128", length: 32, id: "106220", rule: "mixed"},
	{name: "Tiger128HMAC", length: 32, id: "106225", rule: "mixed"},
	{name: "md5passsalt", length: 32, id: "106240", rule: "mixed"},
	{name: "md5saltmd5pass", length: 32, id: "106260", rule: "mixed"},
	{name: "md5saltpass", length: 32, id: "106280", rule: "mixed"},
	{name: "md5saltpasssalt", length: 32, id: "106300", rule: "mixed"},
	{name: "md5saltpassusername", length: 32, id: "106320", rule: "mixed"},
	{name: "md5saltmd5pass", length: 32, id: "106340", rule: "mixed"},
	{name: "md5saltmd5passsalt", length: 32, id: "106360", rule: "mixed"},
	{name: "md5saltmd5passsalt", length: 32, id: "106380", rule: "mixed"},
	{name: "md5saltmd5saltpass", length: 32, id: "106400", rule: "mixed"},
	{name: "md5saltmd5md5passsalt", length: 32, id: "106420", rule: "mixed"},
	{name: "md5username0pass", length: 32, id: "106440", rule: "mixed"},
	{name: "md5usernameLFpass", length: 32, id: "106460", rule: "mixed"},
	{name: "md5usernamemd5passsalt", length: 32, id: "106480", rule: "mixed"},
	{name: "md5md5pass", length: 32, id: "106500", rule: "mixed"},
	{name: "md5md5passsalt", length: 32, id: "106520", rule: "mixed"},
	{name: "md5md5passmd5salt", length: 32, id: "106540", rule: "mixed"},
	{name: "md5md5saltpass", length: 32, id: "106560", rule: "mixed"},
	{name: "md5md5saltmd5pass", length: 32, id: "106580", rule: "mixed"},
	{name: "md5md5usernamepasssalt", length: 32, id: "106600", rule: "mixed"},
	{name: "md5md5md5pass", length: 32, id: "106620", rule: "mixed"},
	{name: "md5md5md5md5pass", length: 32, id: "106640", rule: "mixed"},
	{name: "md5md5md5md5md5pass", length: 32, id: "106660", rule: "mixed"},
	{name: "md5sha1pass", length: 32, id: "106680", rule: "mixed"},
	{name: "md5sha1md5pass", length: 32, id: "106700", rule: "mixed"},
	{name: "md5sha1md5sha1pass", length: 32, id: "106720", rule: "mixed"},
	{name: "md5strtouppermd5pass", length: 32, id: "106740", rule: "mixed"},
	{name: "LineageIIC4", length: 34, id: "107080", rule: "mixed_prefix:0x"},
	{name: "MD5phpBB3", length: 34, id: "107040", rule: "notalnum_prefix:$H$"},
	{name: "MD5Unix", length: 34, id: "107060", rule: "notalnum_prefix:$1$"},
	{name: "MD5Wordpress", length: 34, id: "107020", rule: "notalnum_prefix:$P$"},
	{name: "MD5APR", length: 37, id: "108020", rule: "nondigit_nonalpha_prefix:$apr"},
	{name: "Haval160", length: 40, id: "109100", rule: "mixed"},
	{name: "Haval160HMAC", length: 40, id: "109200", rule: "mixed"},
	{name: "MySQL5", length: 40, id: "109040", rule: "mixed"},
	{name: "MySQL160bit", length: 41, id: "109060", rule: "notalnum_prefix:*"},
	{name: "RipeMD160", length: 40, id: "109120", rule: "mixed"},
	{name: "RipeMD160HMAC", length: 40, id: "109180", rule: "mixed"},
	{name: "SHA1", length: 40, id: "109020", rule: "mixed"},
	{name: "SHA1HMAC", length: 40, id: "109140", rule: "mixed"},
	{name: "SHA1MaNGOS", length: 40, id: "109220", rule: "mixed"},
	{name: "SHA1MaNGOS2", length: 40, id: "109240", rule: "mixed"},
	{name: "Tiger160", length: 40, id: "109080", rule: "mixed"},
	{name: "Tiger160HMAC", length: 40, id: "109160", rule: "mixed"},
	{name: "sha1passsalt", length: 40, id: "109260", rule: "mixed"},
	{name: "sha1saltpass", length: 40, id: "109280", rule: "mixed"},
	{name: "sha1saltmd5pass", length: 40, id: "109300", rule: "mixed"},
	{name: "sha1saltmd5passsalt", length: 40, id: "109320", rule: "mixed"},
	{name: "sha1saltsha1pass", length: 40, id: "109340", rule: "mixed"},
	{name: "sha1saltsha1saltsha1pass", length: 40, id: "109360", rule: "mixed"},
	{name: "sha1usernamepass", length: 40, id: "109380", rule: "mixed"},
	{name: "sha1usernamepasssalt", length: 40, id: "109400", rule: "mixed"},
	{name: "sha1md5pass", length: 40, id: "1094202", rule: "mixed"},
	{name: "sha1md5passsalt", length: 40, id: "109440", rule: "mixed"},
	{name: "sha1md5sha1pass", length: 40, id: "109460", rule: "mixed"},
	{name: "sha1sha1pass", length: 40, id: "109480", rule: "mixed"},
	{name: "sha1sha1passsalt", length: 40, id: "109500", rule: "mixed"},
	{name: "sha1sha1passsubstrpass03", length: 40, id: "109520", rule: "mixed"},
	{name: "sha1sha1saltpass", length: 40, id: "109540", rule: "mixed"},
	{name: "sha1sha1sha1pass", length: 40, id: "109560", rule: "mixed"},
	{name: "sha1strtolowerusernamepass", length: 40, id: "109580", rule: "mixed"},
	{name: "Haval192", length: 48, id: "110040", rule: "mixed"},
	{name: "Haval192HMAC", length: 48, id: "110080", rule: "mixed"},
	{name: "Tiger192", length: 48, id: "110020", rule: "mixed"},
	{name: "Tiger192HMAC", length: 48, id: "110060", rule: "mixed"},
	{name: "MD5passsaltjoomla1", length: 54, id: "112020", rule: "colon32"},
	{name: "SHA1Django", length: 54, id: "113020", rule: "notalnum_prefix:sha1$"},
	{name: "Haval224", length: 56, id: "114040", rule: "mixed"},
	{name: "Haval224HMAC", length: 56, id: "114080", rule: "mixed"},
	{name: "SHA224", length: 56, id: "114020", rule: "mixed"},
	{name: "SHA224HMAC", length: 56, id: "114060", rule: "mixed"},
	{name: "SHA256", length: 64, id: "115020", rule: "mixed"},
	{name: "SHA256HMAC", length: 64, id: "115120", rule: "mixed"},
	{name: "Haval256", length: 64, id: "115040", rule: "mixed"},
	{name: "Haval256HMAC", length: 64, id: "115140", rule: "mixed"},
	{name: "GOSTR341194", length: 64, id: "115060", rule: "mixed"},
	{name: "RipeMD256", length: 64, id: "115080", rule: "mixed"},
	{name: "RipeMD256HMAC", length: 64, id: "115160", rule: "mixed"},
	{name: "SNEFRU256", length: 64, id: "115100", rule: "mixed"},
	{name: "SNEFRU256HMAC", length: 64, id: "115180", rule: "mixed"},
	{name: "SHA256md5pass", length: 64, id: "115200", rule: "mixed"},
	{name: "SHA256sha1pass", length: 64, id: "115220", rule: "mixed"},
	{name: "MD5passsaltjoomla2", length: 65, id: "116020", rule: "colon32"},
	{name: "SAM", length: 65, id: "116040", rule: "sam"},
	{name: "SHA256Django", length: 71, id: "117020", rule: "notalnum_prefix:sha256"},
	{name: "RipeMD320", length: 80, id: "118020", rule: "mixed"},
	{name: "RipeMD320HMAC", length: 80, id: "118040", rule: "mixed"},
	{name: "SHA384", length: 96, id: "119020", rule: "mixed"},
	{name: "SHA384HMAC", length: 96, id: "119040", rule: "mixed"},
	{name: "SHA256s", length: 116, id: "120020", rule: "notalnum_prefix:$6$"},
	{name: "SHA384Django", length: 112, id: "121020", rule: "notalnum_prefix:sha384$"},
	{name: "SHA512", length: 128, id: "122020", rule: "mixed"},
	{name: "SHA512HMAC", length: 128, id: "122060", rule: "mixed"},
	{name: "Whirlpool", length: 128, id: "122040", rule: "mixed"},
	{name: "WhirlpoolHMAC", length: 128, id: "122080", rule: "mixed"},
}

// Detect appends every matching detector ID to jerar. Multiple formats can
// match the same input because these checks inspect shape, not hash validity.
func Detect(hash string, jerar *[]string) {
	for _, r := range hashRules {
		if len(hash) != r.length {
			continue
		}
		if matchesRule(hash, r.rule) {
			*jerar = append(*jerar, r.id)
		}
	}
}

func matchesRule(s, rule string) bool {
	switch {
	case rule == "mixed":
		return isAlnumASCII(s) && !isAlphaASCII(s) && !isDigitASCII(s)
	case rule == "notalpha":
		return isAlnumASCII(s) && !isAlphaASCII(s)
	case rule == "digits":
		return isDigitASCII(s)
	case rule == "nondigit_nonalpha":
		return !isDigitASCII(s) && !isAlphaASCII(s)
	case rule == "colon32":
		return !isDigitASCII(s) && !isAlphaASCII(s) && !isAlnumASCII(s) && len(s) > 32 && s[32] == ':'
	case rule == "sam":
		return !isDigitASCII(s) && !isAlphaASCII(s) && !isAlnumASCII(s) && !isLowerASCII(s) && len(s) > 32 && s[32] == ':'
	case hasPrefixRule(rule, "mixed_prefix:"):
		p := rule[len("mixed_prefix:"):]
		return isAlnumASCII(s) && !isAlphaASCII(s) && !isDigitASCII(s) && hasPrefix(s, p)
	case hasPrefixRule(rule, "notalnum_prefix:"):
		p := rule[len("notalnum_prefix:"):]
		return !isDigitASCII(s) && !isAlphaASCII(s) && !isAlnumASCII(s) && hasPrefix(s, p)
	case hasPrefixRule(rule, "nondigit_nonalpha_prefix:"):
		p := rule[len("nondigit_nonalpha_prefix:"):]
		return !isDigitASCII(s) && !isAlphaASCII(s) && hasPrefix(s, p)
	default:
		return false
	}
}

func hasPrefixRule(rule, prefix string) bool {
	return len(rule) >= len(prefix) && rule[:len(prefix)] == prefix
}
func hasPrefix(s, prefix string) bool { return len(s) >= len(prefix) && s[:len(prefix)] == prefix }

func isAlphaASCII(s string) bool {
	if s == "" {
		return false
	}
	for i := 0; i < len(s); i++ {
		c := s[i]
		if !((c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')) {
			return false
		}
	}
	return true
}
func isDigitASCII(s string) bool {
	if s == "" {
		return false
	}
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return true
}
func isAlnumASCII(s string) bool {
	if s == "" {
		return false
	}
	for i := 0; i < len(s); i++ {
		c := s[i]
		if !((c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')) {
			return false
		}
	}
	return true
}
func isLowerASCII(s string) bool {
	hasLower := false
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c >= 'A' && c <= 'Z' {
			return false
		}
		if c >= 'a' && c <= 'z' {
			hasLower = true
		}
	}
	return hasLower
}

// NameForID returns the first detector name associated with an ID.
func NameForID(id string) string {
	for _, r := range hashRules {
		if r.id == id {
			return r.name
		}
	}
	return "Unknown (" + id + ")"
}
