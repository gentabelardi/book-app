package models
import (
	"time"
)

type Book struct {
    ID          	int			`gorm:"primaryKey" json:"id"`
    Title       	string		`json:"title"`
	Description   	string 		`json:"description"`
	Genre       	string	  	`json:"genre"`
	Author    		string    	`json:"author"`
	Price       	float64		`json:"price"`
	StockQuantity 	int			`json:"stock_quantity"`
	Published 		time.Time 	`json:"published"`
}