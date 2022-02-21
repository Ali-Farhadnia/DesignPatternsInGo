package factories

func NewCardFactory() func(string, int, string, int) *Card {
	return func(name string, age int, as string, experience_level int) *Card {
		company_logo := []byte{1, 2, 3}      // Meaningless data for testing/string:☺☻♥
		manager_signature := []byte{4, 5, 6} // Meaningless data for testing/string:♦♣♠

		switch as {
		case "manager":
			return &Card{
				person: Employee{
					name:             name,
					age:              age,
					as:               as,
					experience_level: experience_level,
					salary:           float64(experience_level * 5000),
				},
				company_logo:      company_logo,
				manager_signature: manager_signature,
			}

		case "assistant":
			return &Card{
				person: Employee{
					name:             name,
					age:              age,
					as:               as,
					experience_level: experience_level,
					salary:           float64(experience_level * 4000),
				},
				company_logo:      company_logo,
				manager_signature: manager_signature,
			}

		case "employee":
			return &Card{
				person: Employee{
					name:             name,
					age:              age,
					as:               as,
					experience_level: experience_level,
					salary:           float64(experience_level * 3000),
				},
				company_logo:      company_logo,
				manager_signature: manager_signature,
			}
		default:
			return nil
		}
	}
}

/*
func main() {
	name := "Ali"
	age := 20
	as := "employee"
	experience_level := 2

	cardFactory := NewCardFactory()
	card := *cardFactory(name, age, as, experience_level)

	fmt.Println("Card:", card)
	fmt.Println("Employee:", card.person)
	fmt.Println("Company logo:", card.company_logo,
		" As string:", string(card.company_logo))
	fmt.Println("Manager signature:", card.manager_signature,
		" As string:", string(card.manager_signature))
}
*/
