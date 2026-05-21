package portfolio

import "database/sql"

type Repository struct {
	db *sql.DB
}

func (r *Repository) GetAll() ([]Coin, error) {
	rows, err := r.db.Query("SELECT id, symbol, amount FROM coins")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var coins []Coin
	for rows.Next() {
		var coin Coin
		err := rows.Scan(&coin.ID, &coin.Symbol, &coin.Amount)
		if err != nil {
			return nil, err
		}
		coins = append(coins, coin)
	}

	return coins, nil
}

func (r *Repository) AddCoin(coin Coin) error {
	_, err := r.db.Exec("INSERT INTO coins (symbol, amount) VALUES ($1, $2)", coin.Symbol, coin.Amount)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) RemoveCoin(id int32) error {
	_, err := r.db.Exec("DELETE FROM coins WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}
