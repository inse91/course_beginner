package usecase

import "errors"

type UseCase struct {
}

func New() *UseCase {
	return &UseCase{}
}

func (uc *UseCase) Sum(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, errors.New("numbers must be positive")
	}
	return a + b, nil
}
