package main

type Performance struct {
	playID   string
	audience int
}

type Invoice struct {
	customer    string
	performance []Performance
}

type Play struct {
	name     string
	typeName string
}

func main() {
}
