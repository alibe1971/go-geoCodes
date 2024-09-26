package geoCodes

type TransCountries map[string]TransCountry

type TransCountry struct {
	Keywords []string `json:"keywords"`
	Name     string   `json:"name"`
	FullName string   `json:"fullName"`
	Demonyms []string `json:"demonyms"`
}
