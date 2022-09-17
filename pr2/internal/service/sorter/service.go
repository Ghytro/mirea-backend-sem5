package sorter

import "context"

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func insertionsort(items []int) {
	for i := 1; i < len(items); i++ {
		for j := i; j > 0; j-- {
			if items[j-1] > items[j] {
				items[j-1], items[j] = items[j], items[j-1]
			}
		}
	}
}

func (s *Service) SortArray(ctx context.Context, array []int) {
	insertionsort(array)
}
