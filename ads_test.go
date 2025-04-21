package main

import (
	"fmt"
	"github.com/feiin/sensitivewords"
	"testing"
)

func TestSensitiveWordsFind(t *testing.T) {
	sensitive := sensitivewords.New()
	sensitive.LoadFromFile("conf/keywords.txt")

	keyword := sensitive.Filter("大麻")

	fmt.Print(keyword)

}
