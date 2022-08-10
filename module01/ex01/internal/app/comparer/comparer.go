package comparer

import (
	"dbreader"
	"fmt"
	"strings"
)

const (
	more = iota
	less
	equally
)

func Comparer(oldRecipes *dbreader.Recipes, newRecipes *dbreader.Recipes) (string, bool) {

	f := func(lenOld, lenNew int) (int, int) {
		if lenOld > lenNew {
			return lenNew, more
		} else if lenOld < lenNew {
			fmt.Printf("lenOld: %v\tlenNew: %v\n", lenOld, lenNew)
			return lenOld, less
		}
		return lenOld, equally
	}

	countCake, whoCakeIsBigger := f(len((*oldRecipes).Foods), len((*newRecipes).Foods))
	var sb strings.Builder

	for i := 0; i < countCake; i++ {

		if (*oldRecipes).Foods[i].Name != (*newRecipes).Foods[i].Name {
			sb.WriteString(
				fmt.Sprintf("ADDED cake \"%s\"\nREMOVED cake \"%s\"\n",
					(*newRecipes).Foods[i].Name,
					(*oldRecipes).Foods[i].Name),
			)
			continue
		}
		if (*oldRecipes).Foods[i].Time != (*newRecipes).Foods[i].Time {
			sb.WriteString(
				fmt.Sprintf("CHANGED cooking time for cake \"%s\" - \"%s\" instead of \"%s\"\n",
					(*oldRecipes).Foods[i].Name,
					(*newRecipes).Foods[i].Time,
					(*oldRecipes).Foods[i].Time),
			)
		}

		countIngr, whoIngrIsBigger := f(len((*oldRecipes).Foods[i].Ingredients), len((*newRecipes).Foods[i].Ingredients))
		for iIng := 0; iIng < countIngr; iIng++ {
			if (*oldRecipes).Foods[i].Ingredients[iIng].Name != (*newRecipes).Foods[i].Ingredients[iIng].Name {
				sb.WriteString(
					fmt.Sprintf("ADDED ingredient \"%s\" for cake \"%s\"\nREMOVED ingredient \"%s\" for cake \"%s\"\n",
						(*newRecipes).Foods[i].Ingredients[iIng].Name,
						(*oldRecipes).Foods[i].Name,
						(*oldRecipes).Foods[i].Ingredients[iIng].Name,
						(*oldRecipes).Foods[i].Name),
				)
				continue
			}
			if (*oldRecipes).Foods[i].Ingredients[iIng].Count != (*newRecipes).Foods[i].Ingredients[iIng].Count {
				sb.WriteString(
					fmt.Sprintf("CHANGED count for ingredient \"%s\" for cake \"%s\" - \"%v\" instead of \"%v\"\n",
						(*oldRecipes).Foods[i].Ingredients[iIng].Name,
						(*oldRecipes).Foods[i].Name,
						(*newRecipes).Foods[i].Ingredients[iIng].Count,
						(*oldRecipes).Foods[i].Ingredients[iIng].Count),
				)
			}
			if (*newRecipes).Foods[i].Ingredients[iIng].Unit == "" {
				sb.WriteString(
					fmt.Sprintf("REMOVED unit \"%s\" for ingredient \"%s\" for cake  \"%s\"\n",
						(*oldRecipes).Foods[i].Ingredients[iIng].Unit,
						(*oldRecipes).Foods[i].Ingredients[iIng].Name,
						(*oldRecipes).Foods[i].Name),
				)
			}
		}
		switch whoIngrIsBigger {
		case more:
			for iMore := len((*newRecipes).Foods[i].Ingredients); iMore < len((*oldRecipes).Foods[i].Ingredients); iMore++ {
				sb.WriteString(
					fmt.Sprintf("REMOVED ingredient \"%s\" for cake \"%s\"\n",
						(*oldRecipes).Foods[i].Ingredients[iMore].Name,
						(*oldRecipes).Foods[i].Name),
				)
			}
		case less:

			for iLess := len((*oldRecipes).Foods[i].Ingredients); iLess < len((*newRecipes).Foods[i].Ingredients); iLess++ {
				sb.WriteString(
					fmt.Sprintf("ADDED ingredient \"%s\" for cake \"%s\"\n",
						(*oldRecipes).Foods[i].Ingredients[iLess].Name,
						(*oldRecipes).Foods[i].Name),
				)
			}
		case equally:
		default:
		}
	}

	switch whoCakeIsBigger {
	case more:
		for iMore := len((*newRecipes).Foods); iMore < countCake; iMore++ {
			sb.WriteString(
				fmt.Sprintf("REMOVED cake \"%s\"\n",
					(*oldRecipes).Foods[iMore].Name),
			)
		}
	case less:
		for iLess := len((*oldRecipes).Foods); iLess < len((*newRecipes).Foods); iLess++ {
			sb.WriteString(
				fmt.Sprintf("ADDED cake \"%s\"\n",
					(*oldRecipes).Foods[iLess].Name),
			)
		}
	case equally:
	default:
	}

	if sb.Len() > 0 {
		return sb.String(), false
	}

	return "", true

}
