package portfolio

type Storage interface {
	GetAll() ([]Coin, error)
	AddCoin(coin Coin) error
	RemoveCoin(id int32) error
}

type ServiceInterface interface {
	GetAll() ([]Coin, error)
	AddCoin(coin Coin) error
	RemoveCoin(id int32) error
}
type PriceInterface interface {
	GetPrice(symbol string) (float64, error)
}
