package processor

type ItemProcessor interface {
	Process(item interface{}) interface{}
}
