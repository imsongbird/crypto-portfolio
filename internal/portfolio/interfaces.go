package portfolio

type Storage interface {
	GetAll() []Coin
	AddCoin(coin Coin)
	RemoveCoin(id int32)
}

type ServiceInterface interface {
	GetAll() []Coin
	AddCoin(coin Coin)
	RemoveCoin(id int32)
}
type PriceInterface interface {
	GetPrice(symbol string) (float64, error)
}
