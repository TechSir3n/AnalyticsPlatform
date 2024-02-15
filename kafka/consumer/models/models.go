package models 


type Transaction struct { 
	Order_id uint64 
	Customer_name string 
	Transaction_type string 
	Amount float64 
	Transaction_date string 
}

type Product struct { 
	Product_id uint64 
	Product_name string 
	Price float64 
	Quantity uint64 
}