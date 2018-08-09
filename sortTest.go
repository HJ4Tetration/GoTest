package main

type arrayOfKey []Key

type Key struct {
	a, b int
}

func (k arrayOfKey) Len() int {
	return len(k)
}

func (k arrayOfKey) Less(i, j int) bool {
	return k[i].a < k[j].a
}

func (k arrayOfKey) Swap(i, j int) {
	k[i], k[j] = k[j], k[i]
}

/*func main() {
	arr := arrayOfKey{Key{a: 12, b: 2}, Key{a: 1, b: 3}}
	sort.Sort(arr)
	fmt.Printf("%v\n", arr)
}*/
