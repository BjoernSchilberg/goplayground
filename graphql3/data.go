package main

type Pirate struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name"`
	Life        string `json:"life"`
	YearsActive string `json:"yearsactive"`
	Country     string `json:"country"`
	Comments    string `json:"comments"`
	Wikipedia   string `json:"wikipedia"`
}

var pirates = []Pirate{
	Pirate{
		ID:          "0",
		Name:        "Anne Bonny",
		Life:        "1698-1782",
		YearsActive: "to 1725",
		Country:     "Ireland",
		Comments:    "Despite never commanding a ship herself, Anne Bonny is remembered as one of few female historical pirates.",
		Wikipedia:   "https://en.wikipedia.org/wiki/Anne_Bonny",
	},
	Pirate{
		ID:          "1",
		Name:        "Mary Critchett",
		Life:        "died 1729",
		YearsActive: "1729",
		Country:     "Colonial America",
		Comments:    "She is best known for being one of only four female pirates from the Golden Age of Piracy.",
		Wikipedia:   "https://en.wikipedia.org/wiki/Mary_Critchett",
	},
	Pirate{
		ID:          "2",
		Name:        "Mary Read",
		Life:        "1690-1720",
		YearsActive: "to 1720",
		Country:     "England",
		Comments:    "Along with Anne Bonny, one of few female historical pirates. When captured, Read escaped hanging by claiming she was pregnant, but died soon after of a fever while still in prison.",
		Wikipedia:   "https://en.wikipedia.org/wiki/Mary_Read",
	},
	Pirate{
		ID:          "3",
		Name:        "Flora Burn",
		Life:        "fl. 1741",
		YearsActive: "1740s-1750s",
		Country:     "England",
		Comments:    "Female pirate active mainly off the East coast of North America from 1741.",
		Wikipedia:   "https://en.wikipedia.org/wiki/Flora_Burn",
	},
	Pirate{
		ID:          "4",
		Name:        "Ching Shih",
		Life:        "d. 1844",
		YearsActive: "1807-1810",
		Country:     "China",
		Comments:    "A prominent female pirate in late Qing China.",
		Wikipedia:   "https://en.wikipedia.org/wiki/Ching_Shih",
	},
}
