package clickhouse

import (
	"database/sql"
	"strconv"
)

type ProductDB struct {
	db *sql.DB
}

func (c *ProductDB) InsertData(id, quantity uint64, name, tType string, amount, price, revenue, benefit float64, date string) error {
	_, err := c.db.Exec("INSERT INTO Product(product_id,product_name,price,quantity) VALUES(?,?,?,?)", id, name, price, quantity)
	if err != nil {
		return err
	}

	return nil
}

func (c *ProductDB) UpdateData() error {
	return nil
}

func (c *ProductDB) GetData(id int64) (string, string, float64, string, error) {
	rows, err := c.db.Query("SELECT product_name,price,quantity FROM Product WHERE product_id = $1", id)
	if err != nil {
		return "", "", 0.0, "", err
	}

	var product_name string
	var price, quantity uint64

	for rows.Next() {
		if err = rows.Scan(&product_name, &price, &quantity); err != nil {
			return "", "", 0.0, "", err
		}
	}

	if err = rows.Err(); err != nil {
		return "", "", 0.0, "", err
	}

	return product_name, strconv.FormatUint(price, 10), 0.0, strconv.FormatUint(quantity, 10), nil
}

func (c *ProductDB) GetAllData() error {

	return nil
}

func (c *ProductDB) DeleteData(id int64) error {
	_, err := c.db.Exec("DELETE * FROM Product WHERE product_id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
