package middlewares

type Contract interface {
}

type Middlewares struct{}

func New() Contract {
	return &Middlewares{}
}
