package writer

import "fmt"

type ConsoleItemWriter struct{}

func NewConsoleItemWriter() *ConsoleItemWriter {
	return &ConsoleItemWriter{}
}

func (c *ConsoleItemWriter) Write(items []interface{}) {
	for _, val := range items {
		fmt.Println(string(val.([]byte)))
	}
}
