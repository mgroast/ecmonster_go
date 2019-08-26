package models

type Product struct {
	ID string `gorm:"primary_key"`
	UserID string
	Category string
	Price int
}

func (p Product) FindNewProducts() *[]Product {
	db := GormDB()
	products := &[]Product{}
	db.Find(products)
	return products
}
