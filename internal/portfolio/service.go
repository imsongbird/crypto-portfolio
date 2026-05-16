package portfolio

func (s *Service) GetAll() []Coin {
	return s.repo.GetAll()
}

func (s *Service) AddCoin(coin Coin) {
	s.repo.AddCoin(coin)
}

func (s *Service) RemoveCoin(id int32) {
	s.repo.RemoveCoin(id)
}
