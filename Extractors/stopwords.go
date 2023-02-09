package Extractors

import (
	_ "embed"
	"fmt"
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

var stopWords = map[string][]string{
	"ar": strings.Split(arStopWord, "\r\n"),
	"be": strings.Split(beStopWord, "\r\n"),
	"bg": strings.Split(bgStopWord, "\r\n"),
	"da": strings.Split(daStopWord, "\r\n"),
	"de": strings.Split(deStopWord, "\r\n"),
	"el": strings.Split(elStopWord, "\r\n"),
	"en": strings.Split(enStopWord, "\r\n"),
	"es": strings.Split(esStopWord, "\r\n"),
	"et": strings.Split(etStopWord, "\r\n"),
	"fi": strings.Split(fiStopWord, "\r\n"),
	"fr": strings.Split(frStopWord, "\r\n"),
	"he": strings.Split(heStopWord, "\r\n"),
	"hi": strings.Split(hiStopWord, "\r\n"),
	"hr": strings.Split(hrStopWord, "\r\n"),
	"hu": strings.Split(huStopWord, "\r\n"),
	"id": strings.Split(idStopWord, "\r\n"),
	"it": strings.Split(itStopWord, "\r\n"),
	"ja": strings.Split(jaStopWord, "\r\n"),
	"ko": strings.Split(koStopWord, "\r\n"),
	"lt": strings.Split(ltStopWord, "\r\n"),
	"mk": strings.Split(mkStopWord, "\r\n"),
	"nb": strings.Split(nbStopWord, "\r\n"),
	"nl": strings.Split(nlStopWord, "\r\n"),
	"no": strings.Split(noStopWord, "\r\n"),
	"pl": strings.Split(plStopWord, "\r\n"),
	"pt": strings.Split(ptStopWord, "\r\n"),
	"ro": strings.Split(roStopWord, "\r\n"),
	"ru": strings.Split(ruStopWord, "\r\n"),
	"sl": strings.Split(slStopWord, "\r\n"),
	"sr": strings.Split(srStopWord, "\r\n"),
	"sv": strings.Split(svStopWord, "\r\n"),
	"sw": strings.Split(swStopWord, "\r\n"),
	"th": strings.Split(thStopWord, "\r\n"),
	"tr": strings.Split(trStopWord, "\r\n"),
	"uk": strings.Split(ukStopWord, "\r\n"),
	"vi": strings.Split(viStopWord, "\r\n"),
	"zh": strings.Split(zhStopWord, "\r\n"),
}

var punctuation = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
var removePunctuationRegex = regexp.MustCompile(`[` + punctuation + `]`)

type WordStats struct {
	StopWordCount int
	WordCount     int
	StopWords     []string
}

func getWordCount(content string, language string) (WordStats, error) {
	var wordStats WordStats
	stopWords, ok := stopWords[language]
	if !ok {
		return wordStats, fmt.Errorf("language %s not supported", language)
	}
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
	return wordStats, nil
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
