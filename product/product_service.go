package product

type ProductService struct {
	productRepo ProductRepository
}

func provideProductService(pr ProductRepository) ProductService {
	return ProductService{productRepo: pr}
}

func (ps *ProductService) findAll() []Product {
	return ps.productRepo.findAll()
}

func (ps *ProductService) findByID(id int) Product {
	return ps.productRepo.findByID(id)
}

func (ps *ProductService) save(product Product) Product {
	return ps.productRepo.save(product)
}

func (ps *ProductService) delete(product Product) {
	ps.productRepo.delete(product)
}
