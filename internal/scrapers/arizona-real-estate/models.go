package arizona

type Property struct {
	Name        string
	Description string
	Price       string
	Agency      string
}

func (p Property) GetHeaders() []string {
	return []string{"Name", "Price", "Agency", "Description"}
}

func (p Property) ToSlice() []string {
	return []string{p.Name, p.Price, p.Agency, p.Description}
}
