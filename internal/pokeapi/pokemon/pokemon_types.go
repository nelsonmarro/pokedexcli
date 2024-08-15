package pokemon

type Pokemon struct {
	Sprites        Sprites   `json:"sprites"`
	Species        Species   `json:"species"`
	Name           string    `json:"name"`
	Forms          []Species `json:"forms"`
	Types          []Type    `json:"types"`
	Stats          []Stat    `json:"stats"`
	Moves          []Move    `json:"moves"`
	Abilities      []Ability `json:"abilities"`
	Height         int64     `json:"height"`
	Weight         int64     `json:"weight"`
	Order          int64     `json:"order"`
	ID             int64     `json:"id"`
	BaseExperience int64     `json:"base_experience"`
	IsDefault      bool      `json:"is_default"`
}

type Ability struct {
	Ability  Species `json:"ability"`
	Slot     int64   `json:"slot"`
	IsHidden bool    `json:"is_hidden"`
}

type Species struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Move struct {
	Move Species `json:"move"`
}

type Type struct {
	Type Species `json:"type"`
	Slot int64   `json:"slot"`
}

type Sprites struct {
	BackFemale       *string  `json:"back_female"`
	BackShinyFemale  *string  `json:"back_shiny_female"`
	FrontFemale      *string  `json:"front_female"`
	FrontShinyFemale *string  `json:"front_shiny_female"`
	Animated         *Sprites `json:"animated,omitempty"`
	BackDefault      string   `json:"back_default"`
	BackShiny        string   `json:"back_shiny"`
	FrontDefault     string   `json:"front_default"`
	FrontShiny       string   `json:"front_shiny"`
}

type Stat struct {
	Stat     Species `json:"stat"`
	BaseStat int64   `json:"base_stat"`
	Effort   int64   `json:"effort"`
}
