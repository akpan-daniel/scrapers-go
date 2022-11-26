package amazon

type Product struct {
	Title     string
	Price     string
	URL       string
	Seller    string
	Rating    string
	Questions string
	Reviews   string
}

func (p Product) GetHeaders() []string {
	return []string{"Title", "Price", "URL", "Seller", "Rating", "Questions", "Reviews"}
}

func (p Product) ToSlice() []string {
	return []string{p.Title, p.Price, p.URL, p.Seller, p.Rating, p.Questions, p.Reviews}
}
