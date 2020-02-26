/*
Package proverb implements

Given a list of inputs, generate the relevant proverb.
For example, given the list ["nail", "shoe", "horse", "rider", "message", "battle", "kingdom"],
you will output the full text of this proverbial rhyme:

For want of a nail the shoe was lost.
For want of a shoe the horse was lost.
For want of a horse the rider was lost.
For want of a rider the message was lost.
For want of a message the battle was lost.
For want of a battle the kingdom was lost.
And all for the want of a nail.

*/
package proverb

// Proverb - Given a list of inputs, generate the relevant proverb
func Proverb(rhyme []string) []string {
	proverb := []string{}

	if len(rhyme) != 0 {
		for i, v := range rhyme {
			if i < len(rhyme)-1 {
				proverb = append(proverb, "For want of a "+v+" the "+string(rhyme[i+1])+" was lost.")
			} else {
				proverb = append(proverb, "And all for the want of a "+string(rhyme[0])+".")
			}
		}
	}

	return proverb
}
