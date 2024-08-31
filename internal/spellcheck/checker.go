package spellcheck

import "iter"

type CheckResult struct {
	Word        string
	Correct     bool
	Suggestions map[string]int
}

type Checker interface {
	Check(word string) (CheckResult, error)
	IncorrectWords(words []string) iter.Seq2[CheckResult, error]
}
