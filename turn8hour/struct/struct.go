package main

type Hero struct {
	Name  string
	Ad    int
	Level int
}

func NewHero(name string, ad int, level int) *Hero {
	return &Hero{Name: name, Ad: ad, Level: level}
}

func main() {

}
