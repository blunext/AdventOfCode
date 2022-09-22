package Day04

import (
	"regexp"
	"strconv"
)

type check func(string) bool
type checker func(string) bool

func NewChecker(ch ...check) checker {
	return func(s string) bool {
		for _, checkFunc := range ch {
			if !checkFunc(s) {
				return false
			}
		}
		return true
	}

}

func digit() check {
	return func(s string) bool {
		_, err := strconv.Atoi(s)
		if err != nil {
			return false
		}
		return true
	}
}

func length(v int) check {
	return func(s string) bool {
		return len(s) == v
	}
}

func span(min, max int) check {
	return func(s string) bool {
		d, _ := strconv.Atoi(s)
		return d >= min && d <= max
	}
}

// hgt (Height) - a number followed by either cm or in:
// If cm, the number must be at least 150 and at most 193.
// If in, the number must be at least 59 and at most 76.
func height() check {
	return func(s string) bool {

		prefix := s[:len(s)-2]
		height, err := strconv.Atoi(prefix)
		if err != nil {
			return false
		}

		switch suffix := s[len(s)-2:]; suffix {
		case "cm":
			if height < 150 || height > 193 {
				return false
			}
		case "in":
			if height < 59 || height > 76 {
				return false
			}
		default:
			return false
		}
		return true
	}
}

// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
func hairColor() check {
	return func(s string) bool {
		if len(s) != 7 {
			return false
		}
		if s[0:1] != "#" {
			return false
		}

		isValid := regexp.MustCompile(`^[a-f0-9]*$`).MatchString
		return isValid(s[1:])
	}
}

// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
func eyeColor() check {
	return func(s string) bool {
		switch s {
		case "amb":
		case "blu":
		case "brn":
		case "gry":
		case "grn":
		case "hzl":
		case "oth":
		default:
			return false
		}
		return true
	}
}
