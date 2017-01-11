package main

func main() {
	source := StdInLines()
	first := source.Next()
	filters := []Filter{
		&Diff{Last: first},
		&Accum{Acc: first},
	}
	sink := NewTmplStdout("{{.Id}} {{.Data}}\n")
	Process(&source, filters, &sink)
}
