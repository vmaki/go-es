package requestx

type IRequest interface {
	Generate(data interface{}) error
}
