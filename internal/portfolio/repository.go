package portfolio

func (r *Repository) GetAll() []Coin {
	return r.coins
}

func (r *Repository) AddCoin(coin Coin) {
	r.coins = append(r.coins, coin)
}

func (r *Repository) RemoveCoin(id int32) {
	for i, coin := range r.coins {
		if coin.ID == id {
			r.coins = append(r.coins[:i], r.coins[i+1:]...)
			break
		}
	}
}
