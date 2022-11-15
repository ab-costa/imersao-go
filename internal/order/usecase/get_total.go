package usecase

import "github.com/ab-costa/imersao-go/internal/order/entity"

type GetTotalOutputDTO struct {
	Total int
}

type GetTotalUsecase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewGetTotalUsecase(orderRepository *entity.OrderRepositoryInterface) *GetTotalUsecase {
	return &GetTotalUsecase{
		OrderRepository: *orderRepository,
	}
}

func (c *GetTotalUsecase) Execute() (*GetTotalOutputDTO, error) {
	total, err := c.OrderRepository.GetTotal()

	if err != nil {
		return nil, err
	}

	return &GetTotalOutputDTO{
		Total: total,
	}, nil
}