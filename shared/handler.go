package shared

type Handler func(value interface{}) (interface{}, error)
type Handlers []Handler
