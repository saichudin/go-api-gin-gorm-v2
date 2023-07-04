package product

import (
	"saichudin/go-api-gin/config"
)

func GetProducts() *[]Product {
	// product1 := Product{1, "Bantal", "btl-01", 2000}
	// product2 := Product{2, "Sandal", "sdl-01", 14000}

	// products := []Product{product1, product2}
	var products []Product
	err := config.Db.Unscoped().Table("products").Find(&products).Error
	if err != nil {

	}

	return &products
}

func GetProductDetail(id int64) (*Product, error) {
	var product Product

	// product.Id = id
	// product.Name = "Meja"
	// product.Sku = "mj-01"
	// product.Price = 10000
	err := config.Db.Unscoped().Table("products").Where("id = ?", id).First(&product).Error
	if err != nil {

	}

	return &product, err
}
