package main

// Diff is a difference filter.
type Diff struct {
	Last Item
}

func (d *Diff) Next(item Item) Item {
	newData := make([]float32, len(item.Data))
	for i := range newData {
		newData[i] = item.Data[i] - d.Last.Data[i]
	}
	d.Last = item
	return Item{item.Id, newData}
}

// Accum is a accumulative sum filter.
type Accum struct {
	Acc Item
}

func (f *Accum) Next(in Item) (out Item) {
	for i, el := range in.Data {
		f.Acc.Data[i] += el
	}
	out.Id = in.Id
	out.Data = make([]float32, len(in.Data))
	copy(out.Data, f.Acc.Data)
	return
}
