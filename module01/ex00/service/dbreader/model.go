package dbreader

type Ingredients struct {
	Name  string  `json:"ingredient_name" `
	Count float64 `json:"ingredient_count,string" `
	Unit  string  `json:"ingredient_unit,omitempty" `
}

type Cake struct {
	Name        string        `json:"name" `
	Time        string        `json:"time" `
	Ingredients []Ingredients `json:"ingredients" `
}

type Recipes struct {
	Foods []Cake `json:"cake" `
}
