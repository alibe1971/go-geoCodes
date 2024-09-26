package geoCodes

type TransGeneric map[string]TransGenericItem

type TransGenericItem struct {
	Name string `json:"name"`
}
