package GraphQL_Service

// A People resource is an individual person or character within the Star Wars universe.
type People struct {
	ID        string      `json:"id"`
	Name      string      `json:"name"`
	BirthYear *string     `json:"birth_year"`
	EyeColor  *string     `json:"eye_color"`
	Gender    *string     `json:"gender"`
	HairColor *string     `json:"hair_color"`
	Height    *string     `json:"height"`
	Mass      *string     `json:"mass"`
	SkinColor *string     `json:"skin_color"`
	Homeworld *Planet     `json:"homeworld"`
	Films     []*Film     `json:"films"`
	Species   []*Specie   `json:"species"`
	Starships []*Starship `json:"starships"`
	Vehicles  []*Vehicle  `json:"vehicles"`
}
