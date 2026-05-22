package portfolio

type Service struct {
	repo   Storage
	prices PriceInterface
}

func (s *Service) GetAll() ([]Coin, error) {
	coins, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	for i, coin := range coins {
		price, err := s.prices.GetPrice(coin.Symbol)
		if err == nil {
			coins[i].Price = price
		}
	}
	return coins, nil
}

func (s *Service) AddCoin(coin Coin) error {
	return s.repo.AddCoin(coin)
}

func (s *Service) RemoveCoin(id int32) error {
	return s.repo.RemoveCoin(id)
}

func NewService(repo Storage, prices PriceInterface) *Service {
	return &Service{repo: repo, prices: prices}
}
