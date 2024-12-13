package usecase

type Usecase struct {
	defaultMsg    string
	errorMsg      string
	inputErrorMsg string
	p             Provider
}

func NewUsecase(defaultMsg string, errorMsg string, inputErrorMsg string, p Provider) *Usecase {
	return &Usecase{
		defaultMsg:    defaultMsg,
		errorMsg:      errorMsg,
		inputErrorMsg: inputErrorMsg,
		p:             p,
	}
}
