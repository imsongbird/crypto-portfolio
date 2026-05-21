package portfolio

type Service struct {
	repo   Storage
	prices PriceInterface
}

func (s *Service) GetAll() []Coin {
	coins := s.repo.GetAll()
	for i, coin := range coins {
		price, err := s.prices.GetPrice(coin.Symbol)
		if err == nil {
			coins[i].Price = price
		}
	}
	return coins
}

func (s *Service) AddCoin(coin Coin) {
	s.repo.AddCoin(coin)
}

func (s *Service) RemoveCoin(id int32) {
	s.repo.RemoveCoin(id)
}

func NewService(repo Storage, prices PriceInterface) *Service {
	return &Service{repo: repo, prices: prices}
}
