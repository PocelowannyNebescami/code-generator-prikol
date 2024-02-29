package generator

import "strings"

func (generator Generator) Generate() string {
	return strings.Repeat("0", generator.Length)
}
