package main

import (
	"fmt"
)

var x map[string][]string

func initializeDictionary() {
	x = make(map[string][]string)

	x["a"] = []string{"அ", "ஆ"}
	x["A"] = []string{"ஆ"}
	x["aa"] = []string{"ஆ"}
	x["ai"] = []string{"ஐ"}
	x["ae"] = []string{"ஐ"}
	x["au"] = []string{"ஔ"}
	x["ak"] = []string{"ஃ"}

	x["b"] = []string{"பி", "ப", "ப்"}
	x["ba"] = []string{"ப", "பா"}
	x["baa"] = []string{"பா"}
	x["bA"] = []string{"பா"}
	x["bae"] = []string{"பே"}
	x["be"] = []string{"பெ", "பே"}
	x["bE"] = []string{"பே"}
	x["bee"] = []string{"பீ"}
	x["bi"] = []string{"பி", "பை", "பீ"}
	x["bo"] = []string{"பொ", "போ", "பௌ"}
	x["bO"] = []string{"போ"}
	x["boo"] = []string{"பூ"}
	x["bou"] = []string{"பௌ"}
	x["bu"] = []string{"பு"}
	x["buu"] = []string{"பூ"}
	x["bU"] = []string{"பூ"}
	x["by"] = []string{"பை"}

	x["p"] = []string{"ப்", "பி", "ப"}
	x["pa"] = []string{"ப", "பா"}
	x["paa"] = []string{"பா"}
	x["pA"] = []string{"பா"}
	x["pae"] = []string{"பே"}
	x["pe"] = []string{"பெ", "பே"}
	x["pE"] = []string{"பே"}
	x["pee"] = []string{"பீ"}
	x["pi"] = []string{"பி", "பை", "பீ"}
	x["po"] = []string{"பொ", "போ", "பௌ"}
	x["pO"] = []string{"போ"}
	x["poo"] = []string{"பூ"}
	x["pou"] = []string{"பௌ"}
	x["pu"] = []string{"பு"}
	x["puu"] = []string{"பூ"}
	x["pU"] = []string{"பூ"}
	x["py"] = []string{"பை"}

	x["c"] = []string{"ச"}
	x["ca"] = []string{"ச", "சா"}
	x["caa"] = []string{"சா"}
	x["cA"] = []string{"சா"}
	x["cae"] = []string{"சே"}
	x["ce"] = []string{"செ", "சே"}
	x["cE"] = []string{"சே"}
	x["ch"] = []string{"ச்", "ச"}
	x["cha"] = []string{"ச", "சா"}
	x["chaa"] = []string{"சா"}
	x["chai"] = []string{"சை"}
	x["chou"] = []string{"சௌ"}
	x["chA"] = []string{"சா"}
	x["chae"] = []string{"சே"}
	x["che"] = []string{"செ", "சே"}
	x["chE"] = []string{"சே"}
	x["chee"] = []string{"சீ"}
	x["chi"] = []string{"சி", "சீ"}
	x["chii"] = []string{"சீ"}
	x["cho"] = []string{"சொ", "சோ"}
	x["chO"] = []string{"சோ"}
	x["choo"] = []string{"சூ"}
	x["chu"] = []string{"சு", "சூ"}
	x["chU"] = []string{"சூ"}
	x["chy"] = []string{"சை"}
	x["cy"] = []string{"சை"}

	x["s"] = []string{"ச"}
	x["sa"] = []string{"ச", "சா"}
	x["saa"] = []string{"சா"}
	x["sA"] = []string{"சா"}
	x["sae"] = []string{"சே"}
	x["se"] = []string{"செ", "சே"}
	x["sE"] = []string{"சே"}
	x["see"] = []string{"சீ"}
	x["si"] = []string{"சி", "சீ"}
	x["so"] = []string{"சொ", "சோ"}
	x["sO"] = []string{"சோ"}
	x["soo"] = []string{"சூ"}
	x["su"] = []string{"சு", "சூ"}
	x["sU"] = []string{"சூ"}
	x["sy"] = []string{"சை"}

	x["d"] = []string{"ட"}
	x["da"] = []string{"ட", "த", "டா", "தா"}
	x["daa"] = []string{"டா", "தா"}
	x["dA"] = []string{"டா", "தா"}
	x["dai"] = []string{"டை"}
	x["de"] = []string{"தெ", "டெ", "தே", "டே"}
	x["dE"] = []string{"தே", "டே"}
	x["dh"] = []string{"த்"}
	x["dha"] = []string{"த", "தா"}
	x["dhaa"] = []string{"தா"}
	x["dhai"] = []string{"தை"}
	x["dhA"] = []string{"தா"}
	x["dhe"] = []string{"தெ", "தே"}
	x["dhee"] = []string{"தீ"}
	x["dhE"] = []string{"தே"}
	x["dhi"] = []string{"தி", "தீ", "தை"}
	x["dho"] = []string{"தோ", "தொ"}
	x["dhO"] = []string{"தோ"}
	x["dhoo"] = []string{"தூ"}
	x["dhou"] = []string{"தௌ"}
	x["dhu"] = []string{"து", "தூ"}
	x["dhU"] = []string{"தூ"}
	x["dhy"] = []string{"தை"}
	x["do"] = []string{"டு"}
	x["doo"] = []string{"டூ"}
	x["dO"] = []string{"டோ"}
	x["dou"] = []string{"டௌ", "தௌ"}
	x["du"] = []string{"டு", "டூ"}
	x["dU"] = []string{"டூ"}
	x["dy"] = []string{"டை"}

	x["e"] = []string{"எ", "ஏ"}
	x["E"] = []string{"ஏ"}
	x["ei"] = []string{"ஐ"}

	x["g"] = []string{"க", "கா"}
	x["ga"] = []string{"க", "கா"}
	x["gaa"] = []string{"கா"}
	x["gA"] = []string{"கா"}
	x["gae"] = []string{"கே"}
	x["ge"] = []string{"கே", "கெ"}
	x["gE"] = []string{"கே"}
	x["gee"] = []string{"கீ"}
	x["gi"] = []string{"கி", "கீ"}
	x["go"] = []string{"கோ", "கொ"}
	x["gO"] = []string{"கோ"}
	x["goo"] = []string{"கூ"}
	x["gou"] = []string{"கௌ"}
	x["gy"] = []string{"கை"}

	x["ha"] = []string{"க"} //Some may type muham instead of mugam

	x["i"] = []string{"இ"}
	x["ich"] = []string{"ச்"} // This is probably not needed

	x["k"] = []string{"க்", "க", "கா"}
	x["ka"] = []string{"க", "கா"}
	x["kaa"] = []string{"கா"}
	x["kA"] = []string{"கா"}
	x["kae"] = []string{"கே"}
	x["kai"] = []string{"கை"}
	x["ke"] = []string{"கெ", "கே"}
	x["kE"] = []string{"கே"}
	x["kee"] = []string{"கீ"}
	x["ki"] = []string{"கி", "கீ"}
	x["ko"] = []string{"கொ", "கோ"}
	x["kO"] = []string{"கோ"}
	x["koo"] = []string{"கூ"}
	x["kou"] = []string{"கௌ"}
	x["ku"] = []string{"கு", "கூ"}
	x["kU"] = []string{"கூ"}
	x["ky"] = []string{"கை"}

	/* TODO Add ள ழ varisaigal too as we did for n series */
	x["l"] = []string{"ல்"}
	x["la"] = []string{"ல", "லா"}
	x["laa"] = []string{"லா"}
	x["lA"] = []string{"லா"}
	x["lae"] = []string{"லே"}
	x["lai"] = []string{"லை"}
	x["le"] = []string{"லெ", "லே"}
	x["lE"] = []string{"லே"}
	x["lee"] = []string{"லீ"}
	x["li"] = []string{"லி", "லீ"}
	x["lo"] = []string{"லொ", "லோ"}
	x["lO"] = []string{"லோ"}
	x["loo"] = []string{"லூ"}
	x["lou"] = []string{"லௌ"}
	x["lu"] = []string{"லு", "லூ"}
	x["lU"] = []string{"லூ"}
	x["ly"] = []string{"லை"}

	x["L"] = []string{"ள்"}
	x["La"] = []string{"ள", "ளா"}
	x["Laa"] = []string{"ளா"}
	x["LA"] = []string{"ளா"}
	x["Lae"] = []string{"ளே"}
	x["Lai"] = []string{"ளை"}
	x["Le"] = []string{"ளெ", "ளே"}
	x["LE"] = []string{"ளே"}
	x["Lee"] = []string{"ளீ"}
	x["Li"] = []string{"ளி", "ளீ"}
	x["Lo"] = []string{"ளொ", "ளோ"}
	x["LO"] = []string{"ளோ"}
	x["Loo"] = []string{"ளூ"}
	x["Lou"] = []string{"ளௌ"}
	x["Lu"] = []string{"ளு", "ளூ"}
	x["LU"] = []string{"ளூ"}
	x["Ly"] = []string{"ளை"}

	x["m"] = []string{"ம்"}
	x["ma"] = []string{"ம", "மா"}
	x["maa"] = []string{"மா"}
	x["mA"] = []string{"மா"}
	x["mae"] = []string{"மே"}
	x["mai"] = []string{"மை"}
	x["me"] = []string{"மெ"}
	x["mE"] = []string{"மே"}
	x["mee"] = []string{"மீ"}
	x["mi"] = []string{"மி", "மீ"}
	x["mo"] = []string{"மொ", "மோ"}
	x["mO"] = []string{"மோ"}
	x["moo"] = []string{"மூ"}
	x["mou"] = []string{"மௌ"}
	x["mu"] = []string{"மு"}
	x["mU"] = []string{"மூ"}
	x["my"] = []string{"மை"}

	/* TODO Give precedence to ந varisai on 1st letter position in a word
	 * Sort the hashtable values based on the usage count */
	x["n"] = []string{"ன்", "ண்", "ந்", "ங்"}
	x["na"] = []string{"ன", "ந", "ண"}
	x["naa"] = []string{"னா", "நா", "ணா"}
	x["nA"] = []string{"னா", "நா", "ணா"}
	x["nae"] = []string{"னே", "ணே", "நே"}
	x["nai"] = []string{"னை", "ணை", "நை"}
	x["ne"] = []string{"னெ", "ணெ", "நெ"}
	x["nE"] = []string{"னே", "ணே", "நே"}
	x["nee"] = []string{"னீ", "ணீ", "நீ"}
	x["ni"] = []string{"னி", "ணி", "நி"}
	x["no"] = []string{"னொ", "னோ", "நொ", "ணொ", "ணோ", "நோ"}
	x["nO"] = []string{"னோ", "ணோ", "நோ"}
	x["noo"] = []string{"னூ", "ணூ", "நூ"}
	x["nou"] = []string{"னௌ", "நௌ", "ணௌ"}
	x["nu"] = []string{"னு", "ணு", "நு"}
	x["nU"] = []string{"னூ", "ணூ", "நூ"}
	x["ny"] = []string{"னை", "ணை", "நை"}

	x["nh"] = []string{"ந்"}
	x["nha"] = []string{"ந"}
	x["nhaa"] = []string{"நா"}
	x["nhA"] = []string{"நா"}
	x["nhae"] = []string{"நே"}
	x["nhai"] = []string{"நை"}
	x["nhe"] = []string{"நெ"}
	x["nhE"] = []string{"நே"}
	x["nhee"] = []string{"நீ"}
	x["nhi"] = []string{"நி"}
	x["nho"] = []string{"நொ", "நோ"}
	x["nhO"] = []string{"நோ"}
	x["nhoo"] = []string{"நூ"}
	x["nhou"] = []string{"நௌ"}
	x["nhu"] = []string{"நு"}
	x["nhU"] = []string{"நூ"}
	x["nhy"] = []string{"நை"}

	x["N"] = []string{"ண்"}
	x["Na"] = []string{"ண"}
	x["Naa"] = []string{"ணா"}
	x["NA"] = []string{"ணா"}
	x["Nae"] = []string{"ணே"}
	x["Nai"] = []string{"ணை"}
	x["Ne"] = []string{"ணெ"}
	x["NE"] = []string{"ணே"}
	x["Nee"] = []string{"ணீ"}
	x["Ni"] = []string{"ணி"}
	x["No"] = []string{"ணொ", "ணோ"}
	x["NO"] = []string{"ணோ"}
	x["Noo"] = []string{"ணூ"}
	x["Nou"] = []string{"ணௌ"}
	x["Nu"] = []string{"ணு"}
	x["NU"] = []string{"ணூ"}
	x["Ny"] = []string{"ணை"}

	x["o"] = []string{"ஒ", "ஓ"}
	x["O"] = []string{"ஓ"}
	x["oo"] = []string{"ஊ"}
	x["ou"] = []string{"ஔ"}
	x["ow"] = []string{"ஔ"}

	x["p"] = []string{"ப்", "பி", "ப"}
	x["pa"] = []string{"ப"}
	x["paa"] = []string{"பா"}
	x["pA"] = []string{"பா"}
	x["pae"] = []string{"பே"}
	x["pai"] = []string{"பை"}
	x["pe"] = []string{"பெ"}
	x["pE"] = []string{"பே"}
	x["pee"] = []string{"பீ"}
	x["pi"] = []string{"பி"}
	x["po"] = []string{"பொ", "போ"}
	x["pO"] = []string{"போ"}
	x["poo"] = []string{"பூ"}
	x["pou"] = []string{"பௌ"}
	x["pu"] = []string{"பு"}
	x["pU"] = []string{"பூ"}
	x["py"] = []string{"பை"}

	x["r"] = []string{"ர்"}
	x["ra"] = []string{"ர"}
	x["raa"] = []string{"ரா"}
	x["rA"] = []string{"ரா"}
	x["rae"] = []string{"ரே"}
	x["rai"] = []string{"ரை"}
	x["re"] = []string{"ரெ"}
	x["rE"] = []string{"ரே"}
	x["ree"] = []string{"ரீ"}
	x["ri"] = []string{"ரி"}
	x["ro"] = []string{"ரொ", "ரோ"}
	x["rO"] = []string{"ரோ"}
	x["roo"] = []string{"ரூ"}
	x["rou"] = []string{"ரௌ"}
	x["ru"] = []string{"ரு"}
	x["rU"] = []string{"ரூ"}
	x["ry"] = []string{"ரை"}

	x["R"] = []string{"ற்"}
	x["Ra"] = []string{"ற"}
	x["Raa"] = []string{"றா"}
	x["RA"] = []string{"றா"}
	x["Rae"] = []string{"றே"}
	x["Rai"] = []string{"றை"}
	x["Re"] = []string{"றெ"}
	x["RE"] = []string{"றே"}
	x["Ree"] = []string{"றீ"}
	x["Ri"] = []string{"றி"}
	x["Ro"] = []string{"றொ", "றோ"}
	x["RO"] = []string{"றோ"}
	x["Roo"] = []string{"றூ"}
	x["Rou"] = []string{"றௌ"}
	x["Ru"] = []string{"று"}
	x["RU"] = []string{"றூ"}
	x["Ry"] = []string{"றை"}

	x["t"] = []string{"ட்", "ற்", "ட", "டா"}
	x["ta"] = []string{"ட", "டா"}
	x["taa"] = []string{"டா"}
	x["tA"] = []string{"டா"}
	x["tae"] = []string{"டே"}
	x["tai"] = []string{"டை"}
	x["te"] = []string{"டெ"}
	x["tE"] = []string{"டே"}
	x["tee"] = []string{"டீ"}
	x["ti"] = []string{"டி"}
	x["to"] = []string{"டொ", "டோ"}
	x["tO"] = []string{"டோ"}
	x["too"] = []string{"டூ"}
	x["tou"] = []string{"டௌ"}
	x["tu"] = []string{"டு"}
	x["tU"] = []string{"டூ"}
	x["ty"] = []string{"டை"}

	x["u"] = []string{"உ", "ஊ"}
	x["uu"] = []string{"ஊ"}
	x["U"] = []string{"ஊ"}

	x["v"] = []string{"வ்"}
	x["va"] = []string{"வ", "வா"}
	x["vaa"] = []string{"வா"}
	x["vA"] = []string{"வா"}
	x["vae"] = []string{"வே"}
	x["vai"] = []string{"வை"}
	x["ve"] = []string{"வெ"}
	x["vE"] = []string{"வே"}
	x["vee"] = []string{"வீ"}
	x["vi"] = []string{"வி"}
	x["vo"] = []string{"வொ", "வோ"}
	x["vO"] = []string{"வோ"}
	x["voo"] = []string{"வூ"}
	x["vou"] = []string{"வௌ"}
	x["vu"] = []string{"வு"}
	x["vU"] = []string{"வூ"}
	x["vy"] = []string{"வை"}

	x["w"] = []string{"வ்"}
	x["wa"] = []string{"வ", "வா"}
	x["waa"] = []string{"வா"}
	x["wA"] = []string{"வா"}
	x["wae"] = []string{"வே"}
	x["wai"] = []string{"வை"}
	x["we"] = []string{"வெ"}
	x["wE"] = []string{"வே"}
	x["wee"] = []string{"வீ"}
	x["wi"] = []string{"வி"}
	x["wo"] = []string{"வொ", "வோ"}
	x["wO"] = []string{"வோ"}
	x["woo"] = []string{"வூ"}
	x["wou"] = []string{"வௌ"}
	x["wu"] = []string{"வு"}
	x["wU"] = []string{"வூ"}
	x["wy"] = []string{"வை"}

	x["y"] = []string{"ய்"}
	x["ya"] = []string{"ய", "யா"}
	x["yaa"] = []string{"யா"}
	x["yA"] = []string{"யா"}
	x["yae"] = []string{"யே"}
	x["yai"] = []string{"யை"}
	x["ye"] = []string{"யெ"}
	x["yE"] = []string{"யே"}
	x["yee"] = []string{"யீ"}
	x["yi"] = []string{"யி"}
	x["yo"] = []string{"யொ", "யோ"}
	x["yO"] = []string{"யோ"}
	x["yoo"] = []string{"யூ"}
	x["you"] = []string{"யௌ"}
	x["yu"] = []string{"யு"}
	x["yU"] = []string{"யூ"}
	x["yy"] = []string{"யை"}

	x["z"] = []string{"ழ்"}
	x["za"] = []string{"ழ", "ழா"}
	x["zaa"] = []string{"ழா"}
	x["zA"] = []string{"ழா"}
	x["zae"] = []string{"ழே"}
	x["zai"] = []string{"ழை"}
	x["ze"] = []string{"ழெ"}
	x["zE"] = []string{"ழே"}
	x["zee"] = []string{"ழீ"}
	x["zi"] = []string{"ழி"}
	x["zo"] = []string{"ழொ", "ழோ"}
	x["zO"] = []string{"ழோ"}
	x["zoo"] = []string{"ழூ"}
	x["zou"] = []string{"ழௌ"}
	x["zu"] = []string{"ழு"}
	x["zU"] = []string{"ழூ"}
	x["zy"] = []string{"ழை"}

	x["zh"] = []string{"ழ்"}
	x["zha"] = []string{"ழ", "ழா"}
	x["zhaa"] = []string{"ழா"}
	x["zhA"] = []string{"ழா"}
	x["zhae"] = []string{"ழே"}
	x["zhai"] = []string{"ழை"}
	x["zhe"] = []string{"ழெ"}
	x["zhE"] = []string{"ழே"}
	x["zhee"] = []string{"ழீ"}
	x["zhi"] = []string{"ழி"}
	x["zho"] = []string{"ழொ", "ழோ"}
	x["zhO"] = []string{"ழோ"}
	x["zhoo"] = []string{"ழூ"}
	x["zhou"] = []string{"ழௌ"}
	x["zhu"] = []string{"ழு"}
	x["zhU"] = []string{"ழூ"}
	x["zhy"] = []string{"ழை"}

}

func main() {

	initializeDictionary()

	input := "murugaa"
	prefixKeys := ""
	buffer := ""
	completedString := ""

	var clearPrevious bool
	var suggestion string

	for _, v := range input {
		key := string(v)
		fmt.Println(key)
		fmt.Println("----------------")

		suggestion, prefixKeys, clearPrevious = transliterate(prefixKeys, key)
		fmt.Println(suggestion)
		if !clearPrevious {
			fmt.Print("Adding")
			fmt.Println(buffer)
			completedString += buffer
		}
		buffer = suggestion

		fmt.Println("\n")
	}

	completedString += buffer
	fmt.Println(completedString)
}

func transliterate(prefixKeys, key string) (string, string, bool) {

	var results []string
	var newPrefixKeys string

	if len(prefixKeys) != 0 {
		newPrefixKeys = prefixKeys + key
	} else {
		newPrefixKeys = key
	}

	results = x[newPrefixKeys]
	if results != nil {
		return results[0], newPrefixKeys, true
	} else {
		newPrefixKeys = key
		results = x[newPrefixKeys]
		if results != nil {
			return results[0], newPrefixKeys, false
		}
		return "", newPrefixKeys, false
	}
}
