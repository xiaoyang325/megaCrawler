package Extractors

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"math"
	"megaCrawler/Crawler"
	"strconv"
	"strings"
)

func nodesToCheck(dom *goquery.Selection) *goquery.Selection {
	return dom.Find("p,pre,td")
}

func nodeText(n *goquery.Selection) string {
	n.Find("script").Remove()
	n.Find("style").Remove()
	return strings.TrimSpace(n.Text())
}

func updateScore(node *goquery.Selection, score float64) (err error) {
	currentScore, err := getScore(node)
	if err != nil {
		return err
	}
	newScore := currentScore + score
	node.SetAttr("gravityScore", fmt.Sprintf("%f", newScore))
	return nil
}

func getScore(node *goquery.Selection) (currentScore float64, err error) {
	str, ok := node.Attr("gravityScore")
	if !ok {
		node.SetAttr("gravityScore", "0.0")
		return 0.0, nil
	}
	currentScore, err = strconv.ParseFloat(str, 64)
	if err != nil {
		return 0.0, err
	}
	return currentScore, nil
}

func isBootable(node *goquery.Selection, language string) bool {
	PARA := "p"
	MinimumStopWordCount := 5
	MaxStepsAway := 3
	isBootable := false

	node.Prev().Each(func(i int, selection *goquery.Selection) {
		if isBootable {
			return
		}
		if i > MaxStepsAway {
			return
		}
		if goquery.NodeName(selection) != PARA {
			return
		}

		paraText := nodeText(selection)
		wordStats, err := getWordCount(paraText, language)
		if err != nil {
			return
		}
		if wordStats.StopWordCount > MinimumStopWordCount {
			isBootable = true
		}
	})
	return isBootable
}

func CalculateBestNode(dom *goquery.Selection, language string) *goquery.Selection {
	startingBoost := 1.0
	i := 0
	var nodeWithText []*goquery.Selection
	var parentNodes []*goquery.Selection

	nodesToCheck(dom).Each(func(i int, selection *goquery.Selection) {
		textNode := nodeText(selection)
		wordStats, err := getWordCount(textNode, language)
		if err != nil {
			return
		}
		highLinkDensity := isHighLinkDensity(selection, wordStats)
		if wordStats.StopWordCount > 2 && !highLinkDensity {
			nodeWithText = append(nodeWithText, selection)
		}
	})

	nodeCount := len(nodeWithText)
	negativeScore := 0.0
	bottomNegativeScore := float64(nodeCount) * 0.25

	for _, node := range nodeWithText {
		boostScore := 0.0

		if isBootable(node, language) {
			if i >= 0 {
				boostScore = 1.0 / startingBoost
				startingBoost += 1
			}
		}

		if nodeCount > 15 {
			if float64(nodeCount-i) <= bottomNegativeScore {
				booster := bottomNegativeScore - float64(nodeCount-i)
				boostScore = -math.Pow(booster, -2)
				negScore := math.Abs(boostScore) + negativeScore
				if negScore > 40 {
					boostScore = 5.0
				}
			}
		}

		textNode := nodeText(node)
		wordStats, err := getWordCount(textNode, language)
		if err != nil {
			return nil
		}
		upScore := boostScore + float64(wordStats.StopWordCount)

		node.Parent().Each(func(i int, selection *goquery.Selection) {
			err := updateScore(selection, upScore)
			if err != nil {
				fmt.Println(err)
			}
			if !Crawler.Contain(parentNodes, selection) {
				parentNodes = append(parentNodes, selection)
			}
		})
		i++
	}

	topNodeScore := 0.0
	var topNode *goquery.Selection
	for _, node := range parentNodes {
		score, err := getScore(node)
		if err != nil {
			fmt.Println(err)
		}
		if score > topNodeScore {
			topNodeScore = score
			topNode = node
		}
	}
	return topNode
}

//func getSiblingScore(node *goquery.Selection, language string) (base float64) {
//	base = 100000
//	paragraphCount := 0
//	paragraphScore := 0.0
//	node.Find("p").Each(func(i int, selection *goquery.Selection) {
//		textNode := nodeText(node)
//		wordStats := getWordCount(textNode, language)
//		highLinkDensity := isHighLinkDensity(node, wordStats)
//		if wordStats.StopWordCount > 2 && !highLinkDensity {
//			paragraphScore += float64(wordStats.StopWordCount)
//			paragraphCount++
//		}
//	})
//
//	if paragraphCount > 0 {
//		base = (paragraphScore / float64(paragraphCount)) * base
//	}
//	return base
//}
//
//func addSibling(node *goquery.Selection, language string) *goquery.Selection {
//	siblingScore := getSiblingScore(node, language)
//	node.Siblings().Each(func(i int, selection *goquery.Selection) {
//		ps := getSiblingsContent(selection, siblingScore)
//		for _, p := range ps {
//			selection.AddSelection(goquery.NewDocumentFromNode(p).Find("p")
//		}
//	})
//	return node
//}
//
//func getSiblingsContent(node *goquery.Selection, baselineScoreSiblingsPara float64) (content []*html.Node) {
//	if node.Type == html.ElementNode && node.Data == "p" && len(nodeText(node)) > 0 {
//		return []*html.Node{node}
//	}
//	potentialParagraphs := goquery.NewDocumentFromNode(node).Find("p").Nodes
//	if len(potentialParagraphs) == 0 {
//		return []*html.Node{}
//	}
//	var ps []*html.Node
//	for _, p := range potentialParagraphs {
//		text := nodeText(p)
//		if len(text) > 0 {
//			wordStats := getWordCount(text, "en")
//			paragraphScore := float64(wordStats.StopWordCount)
//			siblingBaselineScore := 0.3
//			highLinkDensity := isHighLinkDensity(p, wordStats)
//			score := baselineScoreSiblingsPara * siblingBaselineScore
//			if paragraphScore > score && !highLinkDensity {
//				p := html.Node{Type: html.ElementNode, Data: "p"}
//				p.AppendChild(&html.Node{Type: html.TextNode, Data: text})
//				ps = append(ps, &p)
//			}
//		}
//	}
//	return ps
//}

func postCleanup(node *goquery.Selection, language string) *goquery.Selection {
	//addSibling(node, "en")
	node.Children().Each(func(i int, selection *goquery.Selection) {
		wordCount, err := getWordCount(nodeText(selection), language)
		if err != nil {
			return
		}
		if goquery.NodeName(selection) != "p" && isHighLinkDensity(selection, wordCount) {
			selection.Remove()
		}
	})
	return node
}
