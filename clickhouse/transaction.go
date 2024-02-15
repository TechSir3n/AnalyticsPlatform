package clickhouse

import (
	"database/sql"
	"time"
)

type TransactionDB struct {
	db *sql.DB
}

func (c *TransactionDB) InsertData(id,quantity uint64, name, tType string, amount, price, revenue, benefit float64, date string) error {
	parsedDate, err := time.Parse("02.01.2006", date)
	if err != nil {
		return err
	}

	formattedDate := parsedDate.Format("2006-01-02")

	_, err = c.db.Exec("INSERT INTO Transaction (order_id,customer_name,transaction_type,amount,transaction_date) VALUES(?,?,?,?,?)", id, name, tType, amount, formattedDate)
	if err != nil {
		return err
	}

	return nil
}

func (c *TransactionDB) UpdateData() error {
	return nil
}

func (c *TransactionDB) GetData(id int64) (string, string, float64, string, error) {
	rows, err := c.db.Query("SELECT customer_name,transaction_type,amount,transaction_date FROM Transaction WHERE order_id = $1", id)
	if err != nil {
		return "", "", 0.0, "", err
	}

	var (
		customer_name    string
		transaction_type string
		transaction_date string
	)

	var amount float64

	for rows.Next() {
		if err = rows.Scan(&customer_name, &transaction_type, &amount, &transaction_date); err != nil {
			return "", "", 0.0, "", err
		}
	}

	if err = rows.Err(); err != nil {
		return "", "", 0.0, "", err
	}

	return customer_name, transaction_type, amount, transaction_date, nil
}

func (c *TransactionDB) GetAllData() error {

	return nil
}

func (c *TransactionDB) DeleteData(id int64) error {
	_, err := c.db.Exec("DELETE * FROM Transaction WHERE order_id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
