package main

import "fmt"

type Product struct {
	NameProduct string
	Price       int
	Count       int
}

func (p Product) GetSumProduct() int {
	return p.Count * p.Price
}

func summProduct(p []Product) {

	for _, prod := range p {
		summ := Product.GetSumProduct(prod)
		fmt.Print(summ)
		fmt.Print("\n")
	}
}

func main() {

	slicep := []Product{{NameProduct: "product1", Price: 12, Count: 900},
		{NameProduct: "product2", Price: 10, Count: 200}}

	summProduct(slicep)

}
