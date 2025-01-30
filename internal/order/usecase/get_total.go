package usecase

import (
	"tax-service/internal/order/entity"
)

type GetTotalOutputDTO struct {
	Total float64
}

type GetTotalUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewGetTotalUseCase(orderRepository entity.OrderRepositoryInterface) *GetTotalUseCase {
	return &GetTotalUseCase{OrderRepository: orderRepository}
}

func (uc *GetTotalUseCase) Execute() (*GetTotalOutputDTO, error) {
	total, err := uc.OrderRepository.GetTotal()

	if err != nil {
		return nil, err
	}

	output := &GetTotalOutputDTO{
		Total: total,
	}

	return output, nil
}
