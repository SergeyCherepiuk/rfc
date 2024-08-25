package spellcheck

type Checker interface {
	Check(word string) (correct bool, suggestions map[string]int, err error)
}
