package dbreader

type Ingredients struct {
	name  string
	count float64
	unit  string
}

type Cake struct {
	name        string
	time        string
	ingredients []Ingredients
}

type Recipes struct {
	foods []Cake
}
