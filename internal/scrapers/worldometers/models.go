package worldometers

type Population struct {
	Country    string
	Year       string
	Migrants   string
	Population string
}

func (p Population) GetHeaders() []string {
	return []string{"Year", "Country", "Population", "Migrants"}
}

func (p Population) ToSlice() []string {
	return []string{p.Year, p.Country, p.Population, p.Migrants}
}
