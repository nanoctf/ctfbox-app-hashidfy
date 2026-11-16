package detector

import "unicode"

func IsAlpha(s string) bool {
    if s == "" {
        return false
    }

    for _, r := range s {
        if !unicode.IsLetter(r) {
            return false
        }
    }
    return true
}

func IsAlnum(s string) bool {
    if s == "" {
        return false
    }

    for _, r := range s {
        if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
            return false
        }
    }
    return true
}

func IsDigit(s string) bool {
    if s == "" {
        return false
    }

    for _, r := range s {
        if !unicode.IsDigit(r) {
            return false
        }
    }
    return true
}