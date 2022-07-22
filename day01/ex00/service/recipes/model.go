package recipes

type Ingredients struct {
	name  string
	count float64
	unit  string
}

type Food struct {
	name        string
	time        string
	ingredients []Ingredients
}

type Recipes struct {
	foods []Food
}
