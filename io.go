package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/template"
)

// PrintSink just prints an item.
type PrintSink struct {
}

func (sink *PrintSink) Next(item Item) {
	fmt.Println(item.Id, item.Data)
}

type TmplSink struct {
	tmpl *template.Template
}

func NewTmplStdout(tmpl string) (sink TmplSink) {
	parsed, err := template.New("sink").Parse(tmpl)
	if err != nil {
		panic(err)
	}
	sink.tmpl = parsed
	return
}

func (sink *TmplSink) Next(item Item) {
	sink.tmpl.Execute(os.Stdout, item)
}

// BufIoLines is a source from a file.
type BufIoLines struct {
	id uint64
	rd *bufio.Reader
}

// StdInLines makes a BufIoLines from stdin.
func StdInLines() (src BufIoLines) {
	src.id = 0
	src.rd = bufio.NewReader(os.Stdin)
	return
}

func atof(num string) float32 {
	result, err := strconv.ParseFloat(num, 32)
	if err != nil {
		panic(err)
	}
	return float32(result)
}

func (src *BufIoLines) Next() (item Item) {
	line, err := src.rd.ReadString('\n')
	if err != nil {
		item.Id = 0
	} else {
		parts := strings.Split(line[:len(line)-1], " ")
		src.id++
		item.Id = src.id
		item.Data = make([]float32, len(parts))
		for i, part := range parts {
			item.Data[i] = atof(part)
		}
	}
	return
}
