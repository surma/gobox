package tr

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"gobox/common"
	"os"
	"regexp"
	"strings"
)

var (
	flagSet      = flag.NewFlagSet("tr", flag.PanicOnError)
	isDelete     = flagSet.Bool("d", false, "Delete characters in SET1, do not translate")
	isSqueeze    = flagSet.Bool("s", false, "replace each input sequence of a repeated character that is listed in SET1 with a single occurence of that character")
	isComplement = flagSet.Bool("c", false, "use the complement of SET1")
	helpFlag     = flagSet.Bool("help", false, "Show this help")
)

func Tr(call []string) error {
	e := flagSet.Parse(call[1:])
	if e != nil {
		return e
	}
	if flagSet.NArg() < 2 || *helpFlag {
		println("`Tr` [OPTION]... SET1 [SET2]")
		flagSet.PrintDefaults()
		return nil
	}

	sets := flagSet.Args()
	if len(sets) > 2 {
		return errors.New("Only two sets allowed.")
	}

	if len(sets) == 1 && *isDelete {
		return errors.New("Not enough args supplied")
	}

	translation, e := preprocess(sets)
	if e != nil {
		return e
	}
	return invoke(translation)
}

// Parses SET1 and SET2 into a map of translations
func preprocess(sets []string) (map[*regexp.Regexp]string, error) {
	translations := map[*regexp.Regexp]string{}
	set1 := sets[0]

	var set1Part, set2Part string
	var set2len int
	var err error
	if *isDelete {
		// processSet1 only
		for len(set1) > 0 {
			set1Part, set1, _ = nextPartSet1(set1)
			reg, err := toRegexp(set1Part)
			if err != nil {
				return nil, err
			}
			//replacement is empty-string
			translations[reg] = ""
		}
	} else {
		// process both sets together
		set2 := sets[1]
		for len(set1) > 0 {
			set1Part, set1, set2len = nextPartSet1(set1)
			var set2New string
			set2Part, set2New, err = nextPartSet2(set2, set2len)
			if len(set2New) > 0 { //incase set2 is shorter than set1, behave like BSD tr (rather than SystemV, which truncates set1 instead)
				set2 = set2New
			}
			if err != nil {
				return nil, err
			}
			reg, err := toRegexp(set1Part)
			if err != nil {
				return nil, err
			}
			translations[reg] = set2Part
		}
	}
	return translations, nil
}

func toRegexp(set1Part string) (*regexp.Regexp, error) {
	maybeSqueeze := ""
	maybeComplement := ""
	if *isSqueeze {
		maybeSqueeze = "+"
	}
	if *isComplement {
		maybeComplement = "^"
	}
	regString := "^[" + maybeComplement + set1Part + "]" + maybeSqueeze
	//fmt.Println(regString)
	reg, err := regexp.Compile(regString)
	return reg, err
}

// Parser for SET1. Supports single-chars, ranges, and regex-like character groups
func nextPartSet1(set1 string) (string, string, int) {
	if strings.HasPrefix(set1, "[") {
		//find matching
		if strings.Contains(set1, "]") {
			return set1[:strings.Index(set1, "]")+1], set1[strings.Index(set1, "]")+1:], 1

		} else {
			return set1[:1], set1[1:], 1
		}
	} else if len(set1) > 2 && set1[1] == '-' {
		return set1[:3], set1[3:], 1
	} else {
		return set1[:1], set1[1:], 1
	}
}

// Parser for SET2. Supports single and multiple chars
func nextPartSet2(set2 string, set2len int) (string, string, error) {
	if len(set2) < set2len {
		return "", "", errors.New(fmt.Sprintf("Error out of range (%d - %s)", set2len, set2))
	}
	return set2[:set2len], set2[set2len:], nil
}

// Invoke actually carries out the command
func invoke(translations map[*regexp.Regexp]string) error {
	in := common.NewBufferedReader(os.Stdin)
	var buffer bytes.Buffer
	var err error
	for {
		remainder, err := in.ReadWholeLine()
		if err != nil {
			break
		}
		for len(remainder) > 0 {
			trimLeft := 1
			nextPart := remainder[:trimLeft]
			for reg, v := range translations {
				match := reg.MatchString(remainder)
				if match {
					toReplace := reg.FindString(remainder)
					replacement := reg.ReplaceAllString(toReplace, v)
					nextPart = replacement
					//if squeezing has taken place, remove more leading chars accordingly
					trimLeft = len(toReplace)
					if !*isComplement {
						break
					}
				} else if *isComplement {
					// this is a double-negative - non-match of negative-regex.
					// This implies that set1 matches the current input character.
					// So, keep it as-is and break out of the loop.
					trimLeft = 1
					nextPart = remainder[:trimLeft]
					break
				}
			}

			remainder = remainder[trimLeft:]
			buffer.WriteString(nextPart)
		}
		buffer.WriteString("\n")
	}
	out := buffer.String()
	_, err = fmt.Fprintf(os.Stdout, out)
	return err
}
