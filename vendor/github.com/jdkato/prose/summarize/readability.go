package summarize

import (
	"math"

	"github.com/jdkato/prose/internal/util"
)

// FleschKincaid computes the Flesch–Kincaid grade level of the Document d.
// https://en.wikipedia.org/wiki/Flesch%E2%80%93Kincaid_readability_tests
func (d *Document) FleschKincaid() float64 {
	x := 0.39 * d.NumWords / d.NumSentences
	y := 11.8 * d.NumSyllables / d.NumWords
	return x + y - 15.59
}

// FleschReadingEase computes the Flesch reading-ease score of the Document d.
// https://en.wikipedia.org/wiki/Flesch%E2%80%93Kincaid_readability_tests
func (d *Document) FleschReadingEase() float64 {
	x := 1.015 * d.NumWords / d.NumSentences
	y := 84.6 * d.NumSyllables / d.NumWords
	return 206.835 - x - y
}

// GunningFog computes the Gunning Fog index score of the Document d.
// https://en.wikipedia.org/wiki/Gunning_fog_index
func (d *Document) GunningFog() float64 {
	x := d.NumWords / d.NumSentences
	y := d.NumComplexWords / d.NumWords
	return 0.4 * (x + 100.0*y)
}

// SMOG computes the SMOG grade of the Document d.
// https://en.wikipedia.org/wiki/SMOG
func (d *Document) SMOG() float64 {
	return 1.0430*math.Sqrt(d.NumPolysylWords*30.0/d.NumSentences) + 3.1291
}

// AutomatedReadability computes the automated readability index score of the
// Document d.
// https://en.wikipedia.org/wiki/Automated_readability_index
func (d *Document) AutomatedReadability() float64 {
	x := 4.71 * (d.NumCharacters / d.NumWords)
	y := 0.5 * (d.NumWords / d.NumSentences)
	return x + y - 21.43
}

// ColemanLiau computes the Coleman–Liau index score of the Document d.
// https://en.wikipedia.org/wiki/Coleman%E2%80%93Liau_index
func (d *Document) ColemanLiau() float64 {
	x := 0.0588 * (d.NumCharacters / d.NumWords) * 100
	y := 0.296 * (d.NumSentences / d.NumWords) * 100
	return x - y - 15.8
}

// DaleChall computes the Dale–Chall score of the Document d.
// https://en.wikipedia.org/wiki/Dale%E2%80%93Chall_readability_formula
func (d *Document) DaleChall() float64 {
	easy := 0.0
	for word := range d.WordFrequency {
		// TODO: look into more efficient lookup techniques.
		if util.StringInSlice(word, easyWords) {
			easy++
		}
	}
	hard := d.NumWords - easy
	x := (hard / d.NumWords) * 100
	y := (d.NumWords / d.NumSentences)
	return 0.1579*x + 0.0496*y
}
