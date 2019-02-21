package main

import (
	_ "github.com/lib/pq"
	"github.com/zloy2005/webshop/cmd"
)

//const (
//	host     = "localhost"
//	port     = 5432
//	user     = "zloy"
//	password = "zloy"
//	dbname   = "db_shop"
//)
//
//type Product struct {
//	gorm.Model
//	Code  string
//	Price uint
//}

//func main() {
//	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
//		"password=%s dbname=%s sslmode=disable",
//		host, port, user, password, dbname)
//	db, err := gorm.Open("postgres", psqlInfo)
//	if err != nil {
//		panic("failed to connect database")
//	}
//	defer db.Close()
//
//	// Migrate the schema
//	db.AutoMigrate(&Product{})
//
//	// Create
//	db.Create(&Product{Code: "L1212", Price: 1000})
//	// Read
//	var product Product
//	db.First(&product, 1)                   // find product with id 1
//	db.First(&product, "code = ?", "L1212") // find product with code l1212
//
//	// Update - update product's price to 2000
//	db.Model(&product).Update("Price", 2000)
//	fmt.Println(product.Code)
//	// Delete - delete product
//	db.Delete(&product)
//	fmt.Println("Successfully connected!")
//
//}
func main() {
	cmd.Execute()
}
