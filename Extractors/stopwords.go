package Extractors

import (
	_ "embed"
	"github.com/PuerkitoBio/goquery"
	"regexp"
	"strings"
)

//go:embed text/stopwords-ar.txt
var arStopWord string

//go:embed text/stopwords-be.txt
var beStopWord string

//go:embed text/stopwords-bg.txt
var bgStopWord string

//go:embed text/stopwords-da.txt
var daStopWord string

//go:embed text/stopwords-de.txt
var deStopWord string

//go:embed text/stopwords-el.txt
var elStopWord string

//go:embed text/stopwords-en.txt
var enStopWord string

//go:embed text/stopwords-es.txt
var esStopWord string

//go:embed text/stopwords-et.txt
var etStopWord string

//go:embed text/stopwords-fi.txt
var fiStopWord string

//go:embed text/stopwords-fr.txt
var frStopWord string

//go:embed text/stopwords-he.txt
var heStopWord string

//go:embed text/stopwords-hi.txt
var hiStopWord string

//go:embed text/stopwords-hr.txt
var hrStopWord string

//go:embed text/stopwords-hu.txt
var huStopWord string

//go:embed text/stopwords-id.txt
var idStopWord string

//go:embed text/stopwords-it.txt
var itStopWord string

//go:embed text/stopwords-ja.txt
var jaStopWord string

//go:embed text/stopwords-ko.txt
var koStopWord string

//go:embed text/stopwords-lt.txt
var ltStopWord string

//go:embed text/stopwords-mk.txt
var mkStopWord string

//go:embed text/stopwords-nb.txt
var nbStopWord string

//go:embed text/stopwords-nl.txt
var nlStopWord string

//go:embed text/stopwords-no.txt
var noStopWord string

//go:embed text/stopwords-pl.txt
var plStopWord string

//go:embed text/stopwords-pt.txt
var ptStopWord string

//go:embed text/stopwords-ro.txt
var roStopWord string

//go:embed text/stopwords-ru.txt
var ruStopWord string

//go:embed text/stopwords-sl.txt
var slStopWord string

//go:embed text/stopwords-sr.txt
var srStopWord string

//go:embed text/stopwords-sv.txt
var svStopWord string

//go:embed text/stopwords-sw.txt
var swStopWord string

//go:embed text/stopwords-th.txt
var thStopWord string

//go:embed text/stopwords-tr.txt
var trStopWord string

//go:embed text/stopwords-uk.txt
var ukStopWord string

//go:embed text/stopwords-vi.txt
var viStopWord string

//go:embed text/stopwords-zh.txt
var zhStopWord string

var stopWords = map[string]string{
	"ar": arStopWord,
	"be": beStopWord,
	"bg": bgStopWord,
	"da": daStopWord,
	"de": deStopWord,
	"el": elStopWord,
	"en": enStopWord,
	"es": esStopWord,
	"et": etStopWord,
	"fi": fiStopWord,
	"fr": frStopWord,
	"he": heStopWord,
	"hi": hiStopWord,
	"hr": hrStopWord,
	"hu": huStopWord,
	"id": idStopWord,
	"it": itStopWord,
	"ja": jaStopWord,
	"ko": koStopWord,
	"lt": ltStopWord,
	"mk": mkStopWord,
	"nb": nbStopWord,
	"nl": nlStopWord,
	"no": noStopWord,
	"pl": plStopWord,
	"pt": ptStopWord,
	"ro": roStopWord,
	"ru": ruStopWord,
	"sl": slStopWord,
	"sr": srStopWord,
	"sv": svStopWord,
	"sw": swStopWord,
	"th": thStopWord,
	"tr": trStopWord,
	"uk": ukStopWord,
	"vi": viStopWord,
	"zh": zhStopWord,
}

var punctuation = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
var removePunctuationRegex = regexp.MustCompile(`[` + punctuation + `]`)

type WordStats struct {
	StopWordCount int
	WordCount     int
	StopWords     []string
}

func getWordCount(content string, language string) WordStats {
	var wordStats WordStats
	stopWords := strings.Split(stopWords[language], "\r\n")
	strippedContent := removePunctuationRegex.ReplaceAllString(content, "")
	candidateWords := strings.Split(strippedContent, " ")

	overlappingWords := make(map[string]bool)
	for _, word := range candidateWords {
		wordStats.WordCount++
		for _, stopWord := range stopWords {
			if word == stopWord {
				wordStats.StopWordCount++
				overlappingWords[word] = true
			}
		}
	}

	for word := range overlappingWords {
		wordStats.StopWords = append(wordStats.StopWords, word)
	}
	wordStats.StopWordCount = len(wordStats.StopWords)
	return wordStats
}

func isHighLinkDensity(content *goquery.Selection, stats WordStats) bool {
	links := content.Find("a")
	if links.Length() == 0 {
		return false
	}

	var sb []string
	links.Each(func(i int, s *goquery.Selection) {
		sb = append(sb, s.Text())
	})

	linksText := strings.Join(sb, "")
	linkWords := strings.Split(linksText, " ")
	linkWordCount := len(linkWords)
	linkCount := links.Length()
	linkDivisor := float64(linkWordCount) / float64(stats.WordCount)
	score := float64(linkCount) * linkDivisor
	return score >= 1
}
