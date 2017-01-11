package main

import "sync"

// Item is a single item in a stream of data.
type Item struct {
	Id   uint64
	Data []float32
}

// Source generates a stream of data.
type Source interface {
	Next() Item
}

// Sink consumes data.
type Sink interface {
	Next(item Item)
}

// Filter is a modification of stream of data.
type Filter interface {
	Next(Item Item) Item
}

// Process runs everything.
func Process(source Source, filters []Filter, sink Sink) {
	item := source.Next()
	for item.Id > 0 {
		for _, filter := range filters {
			item = filter.Next(item)
		}
		sink.Next(item)
		item = source.Next()
	}
}

func doSource(source Source, c chan Item, wg *sync.WaitGroup) {
	item := source.Next()
	for item.Id > 0 {
		c <- item
		item = source.Next()
	}
	close(c)
	wg.Done()
}

func doFilter(filter Filter, cin chan Item, cout chan Item, wg *sync.WaitGroup) {
	for item := range cin {
		cout <- filter.Next(item)
	}
	close(cout)
	wg.Done()
}

func doSink(sink Sink, c chan Item, wg *sync.WaitGroup) {
	for item := range c {
		sink.Next(item)
	}
	wg.Done()
}

// ProcessChans same as Process but with channels.
func ProcessChans(source Source, filters []Filter, sink Sink) {
	n := len(filters)
	var wg sync.WaitGroup
	wg.Add(n + 2)
	chans := make([]chan Item, n+1)

	for j, _ := range chans {
		chans[j] = make(chan Item)
	}

	go doSource(source, chans[0], &wg)

	for i, filter := range filters {
		go doFilter(filter, chans[i], chans[i+1], &wg)
	}

	go doSink(sink, chans[n], &wg)

	wg.Wait()
}
