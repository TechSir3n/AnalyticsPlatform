package clickhouse 

import ( 
	"database/sql"
)


type TransactionDB struct {
	db *sql.DB
}

func (c *TransactionDB) InsertData(id, name, tType string, amount,price,revenue,benefit float64, date string) error {
	return nil
}

func (c *TransactionDB) UpdateData() error {
	return nil
}

func (c *TransactionDB) GetData(id int64) error {
	return nil
}

func (c *TransactionDB) DeleteData(id int64) error {
	return nil
}