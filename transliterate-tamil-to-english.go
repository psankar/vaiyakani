package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"unicode"
)

var x map[string][]string

func initializeDictionary() {
	x = make(map[string][]string)

	x["அ"] = []string{"a"}
	x["ஆ"] = []string{"a", "aa"}
	x["இ"] = []string{"i"}
	x["ஈ"] = []string{"i"}
	x["உ"] = []string{"u"}
	x["ஊ"] = []string{"oo", "u"}
	x["எ"] = []string{"e"}
	x["ஏ"] = []string{"e"}
	x["ஐ"] = []string{"ai"}
	x["ஒ"] = []string{"o"}
	x["ஓ"] = []string{"o"}
	x["ஔ"] = []string{"au"}

	x["க"] = []string{"ka", "ga"}
	x["கா"] = []string{"ka", "ga"}
	x["கி"] = []string{"ki", "gi"}
	x["கீ"] = []string{"ki", "gi"}
	x["கு"] = []string{"ku", "gu"}
	x["கூ"] = []string{"koo", "ku", "goo", "gu"}
	x["கெ"] = []string{"ke", "ge"}
	x["கே"] = []string{"ke", "ge"}
	x["கை"] = []string{"kai", "gai"}
	x["கொ"] = []string{"ko", "go"}
	x["கோ"] = []string{"ko", "go"}
	x["கௌ"] = []string{"kow", "gow"}
	x["க்"] = []string{"k", "g"}

	x["ங"] = []string{"nga"}
	x["ஙா"] = []string{"nga"}
	x["ஙி"] = []string{"ngi"}
	x["ஙீ"] = []string{"ngi"}
	x["ஙு"] = []string{"ngu"}
	x["ஙூ"] = []string{"ngu"}
	x["ஙெ"] = []string{"nge"}
	x["ஙே"] = []string{"nge"}
	x["ஙை"] = []string{"ngai"}
	x["ஙொ"] = []string{"ngo"}
	x["ஙோ"] = []string{"ngo"}
	x["ஙௌ"] = []string{"ngow"}
	x["ங்"] = []string{"ng"}

	x["ச"] = []string{"sa", "cha"}
	x["சா"] = []string{"sa", "cha"}
	x["சி"] = []string{"si", "chi"}
	x["சீ"] = []string{"si", "chi"}
	x["சு"] = []string{"su", "chu"}
	x["சூ"] = []string{"soo", "su", "choo", "chu"}
	x["செ"] = []string{"se", "che"}
	x["சே"] = []string{"se", "che"}
	x["சை"] = []string{"sai", "chai"}
	x["சொ"] = []string{"so", "cho"}
	x["சோ"] = []string{"so", "cho"}
	x["சௌ"] = []string{"sow", "chow"}
	x["ச்"] = []string{"s", "ch"}

	x["ஞ"] = []string{"gna"}
	x["ஞா"] = []string{"gna"}
	x["ஞி"] = []string{"gni"}
	x["ஞீ"] = []string{"gni"}
	x["ஞு"] = []string{"gnu"}
	x["ஞூ"] = []string{"gnu"}
	x["ஞெ"] = []string{"gne"}
	x["ஞே"] = []string{"gne"}
	x["ஞை"] = []string{"gnai"}
	x["ஞொ"] = []string{"gno"}
	x["ஞோ"] = []string{"gno"}
	x["ஞௌ"] = []string{"gnow"}
	x["ஞ்"] = []string{"gn"}

	x["ட"] = []string{"ta", "da"}
	x["டா"] = []string{"ta", "da"}
	x["டி"] = []string{"ti", "di"}
	x["டீ"] = []string{"ti", "di"}
	x["டு"] = []string{"tu", "du"}
	x["டூ"] = []string{"too", "tu", "doo", "du"}
	x["டெ"] = []string{"te", "de"}
	x["டே"] = []string{"te", "de"}
	x["டை"] = []string{"tai", "dai"}
	x["டொ"] = []string{"to", "do"}
	x["டோ"] = []string{"to", "do"}
	x["டௌ"] = []string{"tow", "dow"}
	x["ட்"] = []string{"t", "d"}

	x["ண"] = []string{"na"}
	x["ணா"] = []string{"na"}
	x["ணி"] = []string{"ni"}
	x["ணீ"] = []string{"ni"}
	x["ணு"] = []string{"nu"}
	x["ணூ"] = []string{"noo", "nu"}
	x["ணெ"] = []string{"ne"}
	x["ணே"] = []string{"ne"}
	x["ணை"] = []string{"nai"}
	x["ணொ"] = []string{"no"}
	x["ணோ"] = []string{"no"}
	x["ணௌ"] = []string{"now"}
	x["ண்"] = []string{"n"}

	x["த"] = []string{"tha", "dha"}
	x["தா"] = []string{"tha", "dha"}
	x["தி"] = []string{"thi", "dhi"}
	x["தீ"] = []string{"thi", "dhi"}
	x["து"] = []string{"thu", "dhu"}
	x["தூ"] = []string{"thoo", "thu", "dhoo", "dhu"}
	x["தெ"] = []string{"the", "dhe"}
	x["தே"] = []string{"the", "dhe"}
	x["தை"] = []string{"thai", "dhai"}
	x["தொ"] = []string{"tho", "dho"}
	x["தோ"] = []string{"tho", "dho"}
	x["தௌ"] = []string{"thow", "dhow"}
	x["த்"] = []string{"th", "dh"}

	x["ந"] = []string{"na"}
	x["நா"] = []string{"na", "naa"}
	x["நி"] = []string{"ni"}
	x["நீ"] = []string{"ni", "nee"}
	x["நு"] = []string{"nu"}
	x["நூ"] = []string{"noo", "nu"}
	x["நெ"] = []string{"ne"}
	x["நே"] = []string{"ne"}
	x["நை"] = []string{"nai"}
	x["நொ"] = []string{"no"}
	x["நோ"] = []string{"no"}
	x["நௌ"] = []string{"now"}
	x["ந்"] = []string{"n"}

	x["ப"] = []string{"pa", "ba"}
	x["பா"] = []string{"pa", "ba"}
	x["பி"] = []string{"pi", "bi"}
	x["பீ"] = []string{"pi", "bi"}
	x["பு"] = []string{"pu", "bu"}
	x["பூ"] = []string{"poo", "pu", "boo", "bu"}
	x["பெ"] = []string{"pe", "be"}
	x["பே"] = []string{"pe", "be"}
	x["பை"] = []string{"pai", "bai"}
	x["பொ"] = []string{"po", "bo"}
	x["போ"] = []string{"po", "bo"}
	x["பௌ"] = []string{"pow", "bow"}
	x["ப்"] = []string{"p", "b"}

	x["ம"] = []string{"ma"}
	x["மா"] = []string{"ma"}
	x["மி"] = []string{"mi"}
	x["மீ"] = []string{"mi"}
	x["மு"] = []string{"mu"}
	x["மூ"] = []string{"moo", "mu"}
	x["மெ"] = []string{"me"}
	x["மே"] = []string{"me"}
	x["மை"] = []string{"mai"}
	x["மொ"] = []string{"mo"}
	x["மோ"] = []string{"mo"}
	x["மௌ"] = []string{"mow"}
	x["ம்"] = []string{"m"}

	x["ய"] = []string{"ya"}
	x["யா"] = []string{"ya"}
	x["யி"] = []string{"yi"}
	x["யீ"] = []string{"yi"}
	x["யு"] = []string{"yu"}
	x["யூ"] = []string{"yoo", "yu"}
	x["யெ"] = []string{"ye"}
	x["யே"] = []string{"ye"}
	x["யை"] = []string{"yai"}
	x["யொ"] = []string{"yo"}
	x["யோ"] = []string{"yo"}
	x["யௌ"] = []string{"yow"}
	x["ய்"] = []string{"y"}

	x["ர"] = []string{"ra"}
	x["ரா"] = []string{"ra"}
	x["ரி"] = []string{"ri"}
	x["ரீ"] = []string{"ri"}
	x["ரு"] = []string{"ru"}
	x["ரூ"] = []string{"roo", "ru"}
	x["ரெ"] = []string{"re"}
	x["ரே"] = []string{"re"}
	x["ரை"] = []string{"rai"}
	x["ரொ"] = []string{"ro"}
	x["ரோ"] = []string{"ro"}
	x["ரௌ"] = []string{"row"}
	x["ர்"] = []string{"r"}

	x["ல"] = []string{"la"}
	x["லா"] = []string{"la"}
	x["லி"] = []string{"li"}
	x["லீ"] = []string{"li"}
	x["லு"] = []string{"lu"}
	x["லூ"] = []string{"loo", "lu"}
	x["லெ"] = []string{"le"}
	x["லே"] = []string{"le"}
	x["லை"] = []string{"lai"}
	x["லொ"] = []string{"lo"}
	x["லோ"] = []string{"lo"}
	x["லௌ"] = []string{"low"}
	x["ல்"] = []string{"l"}

	x["வ"] = []string{"va"}
	x["வா"] = []string{"va"}
	x["வி"] = []string{"vi"}
	x["வீ"] = []string{"vi"}
	x["வு"] = []string{"vu"}
	x["வூ"] = []string{"voo", "vu"}
	x["வெ"] = []string{"ve"}
	x["வே"] = []string{"ve"}
	x["வை"] = []string{"vai"}
	x["வொ"] = []string{"vo"}
	x["வோ"] = []string{"vo"}
	x["வௌ"] = []string{"vow"}
	x["வ்"] = []string{"v"}

	x["ழ"] = []string{"za", "zha"}
	x["ழா"] = []string{"za", "zha"}
	x["ழி"] = []string{"zi", "zhi"}
	x["ழீ"] = []string{"zi", "zhi"}
	x["ழு"] = []string{"zu", "zhu"}
	x["ழூ"] = []string{"zoo", "zu", "zhoo", "zhu"}
	x["ழெ"] = []string{"ze", "zhe"}
	x["ழே"] = []string{"ze", "zhe"}
	x["ழை"] = []string{"zai", "zhai"}
	x["ழொ"] = []string{"zo", "zho"}
	x["ழோ"] = []string{"zo", "zho"}
	x["ழௌ"] = []string{"zow", "zhow"}
	x["ழ்"] = []string{"z", "zh"}

	x["ள"] = []string{"la"}
	x["ளா"] = []string{"la"}
	x["ளி"] = []string{"li"}
	x["ளீ"] = []string{"li"}
	x["ளு"] = []string{"lu"}
	x["ளூ"] = []string{"loo", "lu"}
	x["ளெ"] = []string{"le"}
	x["ளே"] = []string{"le"}
	x["ளை"] = []string{"lai"}
	x["ளொ"] = []string{"lo"}
	x["ளோ"] = []string{"lo"}
	x["ளௌ"] = []string{"low"}
	x["ள்"] = []string{"l"}

	x["ற"] = []string{"ra"}
	x["றா"] = []string{"ra"}
	x["றி"] = []string{"ri"}
	x["றீ"] = []string{"ri"}
	x["று"] = []string{"ru"}
	x["றூ"] = []string{"roo", "ru"}
	x["றெ"] = []string{"re"}
	x["றே"] = []string{"re"}
	x["றை"] = []string{"rai"}
	x["றொ"] = []string{"ro"}
	x["றோ"] = []string{"ro"}
	x["றௌ"] = []string{"row"}
	x["ற்"] = []string{"r", "t"}

	x["ன"] = []string{"na"}
	x["னா"] = []string{"na"}
	x["னி"] = []string{"ni"}
	x["னீ"] = []string{"ni"}
	x["னு"] = []string{"nu"}
	x["னூ"] = []string{"noo", "nu"}
	x["னெ"] = []string{"ne"}
	x["னே"] = []string{"ne"}
	x["னை"] = []string{"nai"}
	x["னொ"] = []string{"no"}
	x["னோ"] = []string{"no"}
	x["னௌ"] = []string{"now"}
	x["ன்"] = []string{"n"}
}

func transliterateTamilToEnglish(tamilWord string) []string {

	var results []string

	p := []rune(tamilWord)
	for i := 0; i < len(p); {
		j := i + 1
		for j < len(p) && (unicode.Is(unicode.Mn, p[j]) ||
			unicode.Is(unicode.Me, p[j]) ||
			unicode.Is(unicode.Mc, p[j])) {
			j++
		}

		tamilChar := string(p[i:j])
		englishChars := x[tamilChar]
		if englishChars != nil {

			if results == nil {
				/* english character(s) for the first tamil
				 * character in the tamilWord */
				for _, englishChar := range englishChars {
					results = append(results, englishChar)
				}
			} else {
				/* Copy the results so far into a temporary
				 * results and clear the current list */
				var tempResults []string
				tempResults = append(tempResults, results...)
				results = results[:0]

				/* For each item in the old result, append
				 * the results for the new character */
				for _, result := range tempResults {
					for _, englishChar := range englishChars {
						temp := result
						temp += englishChar
						results = append(results, temp)
					}
				}
			}
		} else {
			//fmt.Println("Character mapping not available for [" + tamilChar + "]")
			return nil
		}

		i = j
	}

	/* Delete impossible words and insert type optimized words */
	impossibleString := false
	for i := 0; i < len(results); i++ {
		result := results[i]

		/* Detects impossible strings
		 * TODO Knuth Morris Pratt algorithm could come in handy here */
		if strings.Contains(result, "bb") {
			impossibleString = true
		} else if strings.Contains(result, "gg") {
			impossibleString = true
		} else if strings.Contains(result, "thdh") {
			impossibleString = true
		} else if strings.Contains(result, "dhth") {
			impossibleString = true
		} else if strings.Contains(result, "dhdh") {
			impossibleString = true
		} else if strings.Contains(result, "kg") {
			impossibleString = true
		} else if strings.Contains(result, "pb") {
			impossibleString = true
		} else if strings.Contains(result, "bp") {
			impossibleString = true
		} else if strings.Contains(result, "gk") {
			impossibleString = true
		} else if strings.Contains(result, "dd") {
			impossibleString = true
		} else if strings.Contains(result, "td") {
			impossibleString = true
		} else if strings.Contains(result, "dt") {
			impossibleString = true
		} else if strings.HasSuffix(result, "b") {
			/* This else block is different Note "HasSuffix" */
			impossibleString = true
		}

		if impossibleString {
			/* Delete element at i */
			copy(results[i:], results[i+1:])
			results[len(results)-1] = ""
			results = results[:len(results)-1]

			i--
			impossibleString = false
			continue
		}

		/* If a word contains thth as well as chch, the following
		 * code segment will not produce a string where both the
		 * tunables are accomodated. But is there such a word ? */
		for _, tunable := range []string{"thth", "chch"} {
			if strings.Contains(result, tunable) {

				/* Insert new tuned string at position i */
				results = append(results, "")
				copy(results[i+1:], results[i:])
				results[i] = strings.Replace(result, tunable,
					tunable[:2], -1)

				/* This i++ and the for loop one helps
				 * in not rechecking the same string again
				 * and again */
				i++
				continue
			}
		}
	}
	return results
}

func main() {

	initializeDictionary()

	for _, corpusFile := range os.Args[1:] {
		//fmt.Println("Analyzing " + corpusFile)

		fileBytes, err := ioutil.ReadFile(corpusFile)
		if err != nil {
			fmt.Println("Error opening file:")
			fmt.Println(err)
			continue
		}

		lines := strings.Split(string(fileBytes), "\n")

		/* note that this will work with only
		 * linux style file line endings */
		for _, line := range lines {
			if line != "" {
				tamilWord := strings.Split(line, ",")[1:][0]
				englishWords := transliterateTamilToEnglish(tamilWord)

				if englishWords == nil {
					/* grantham thavir */
					continue
				}

				fmt.Print(tamilWord)
				for _, englishWord := range englishWords {
					fmt.Print(" " + englishWord)
				}
				fmt.Println()
			}
		}

	}
}
