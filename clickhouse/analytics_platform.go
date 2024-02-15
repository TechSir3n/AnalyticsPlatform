package clickhouse

import (
	"database/sql"
	"strconv"
	"time"
)

type AnalyticsPlatformDB struct {
	db *sql.DB
}

func (c *AnalyticsPlatformDB) InsertData(id, quantity uint64, name, tType string, amount, price, revenue, benefit float64, date string) error {
	parsedDate, err := time.Parse("02.01.2006", date)
	if err != nil {
		return err
	}

	formattedDate := parsedDate.Format("2006-01-02")

	_, err = c.db.Exec("INSERT INTO AnalyticsPlatform(id,revenue,benefit,date) VALUES(?,?,?,?)", id, revenue, benefit, formattedDate)
	if err != nil {
		return err
	}
	return nil
}

func (c *AnalyticsPlatformDB) UpdateData() error {
	return nil
}

func (c *AnalyticsPlatformDB) GetData(id int64) (string, string, float64, string, error) {
	rows, err := c.db.Query("SELECT revenue,benefit,date FROM AnalyticsPlatform WHERE id = $1", id)
	if err != nil {
		return "", "", 0.0, "", err
	}

	var date string
	var benefit, revenue float64

	for rows.Next() {
		if err = rows.Scan(&revenue, &benefit, &date); err != nil {
			return "", "", 0.0, "", err
		}
	}

	if err = rows.Err(); err != nil {
		return "", "", 0.0, "", err
	}

	return date, strconv.FormatFloat(revenue, 'f', -1, 64), benefit, "", nil
}

func (c *AnalyticsPlatformDB) GetAllData() error {

	return nil
}

func (c *AnalyticsPlatformDB) DeleteData(id int64) error {
	_, err := c.db.Exec("DELETE * FROM AnalyticsPlatform WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
