package models

type Name struct {
	NConst    string `json:"nconst"`
	Name      string `json:"name"`
	BirthYear string `json:"birthYear"`
	DeathYear string `json:"deathYear"`
}

type Error struct {
	Message string `json:"error"`
}
