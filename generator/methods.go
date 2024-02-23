package generator

import "strings"

func Init(config Config) (Generator, error) {
	return generator{config}, nil
}

func (gnrtr generator) Generate() string {
	return strings.Repeat("0", gnrtr.config.Length)
}
