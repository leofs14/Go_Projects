package product

type Product struct{
	product string
}

func New(product string)(Product){
	return Product{product: product}
}