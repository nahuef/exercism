// Package proverb Given a list of inputs, generate the relevant proverb.
package proverb

import "fmt"

// Proverb Given a list of inputs, generate the relevant proverb.
func Proverb(rhyme []string) (proverb []string) {
	for i, s := range rhyme {
		if i == len(rhyme)-1 {
			proverb = append(proverb, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
		} else {
			proverb = append(proverb, fmt.Sprintf("For want of a %s the %s was lost.", s, rhyme[i+1]))
		}
	}

	return
}
