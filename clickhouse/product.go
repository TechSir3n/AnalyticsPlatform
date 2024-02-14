package clickhouse

import (
	"database/sql"
)

type ProductDB struct {
	db *sql.DB
}

func (c *ProductDB) InsertData(id, name, tType string, amount, price, revenue, benefit float64, date string) error {

	return nil
}

func (c *ProductDB) UpdateData() error {
	return nil
}

func (c *ProductDB) GetData(id int64) error {
	return nil
}
func (c *ProductDB) DeleteData(id int64) error {
	return nil
}
