package allergies

const testVersion = 1

var allergy = map[string]uint{
	"eggs":         1,
	"peanuts":      2,
	"shellfish":    4,
	"strawberries": 8,
	"tomatoes":     16,
	"chocolate":    32,
	"pollen":       64,
	"cats":         128,
}

func Allergies(score uint) []string {
	foods := make([]string, 0)
	for food, _ := range allergy {
		if AllergicTo(score, food) {
			foods = append(foods, food)
		}
	}
	return foods
}

func AllergicTo(score uint, food string) bool {
	return (score & allergy[food]) == allergy[food]
}
