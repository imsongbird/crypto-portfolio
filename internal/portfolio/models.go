package portfolio

type Coin struct {
	ID     int32
	Symbol string
	Amount float64
}

type Service struct {
	repo Storage
}

type Handler struct {
	service ServiceInterface
}

type Repository struct {
	coins []Coin
}
