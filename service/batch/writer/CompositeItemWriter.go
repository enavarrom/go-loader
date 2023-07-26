package writer

type CompositeItemWriter struct {
	delegates []ItemWriter
}

func NewCompositeItemWriter() *CompositeItemWriter {
	return &CompositeItemWriter{}
}

func (c *CompositeItemWriter) Add(writer ItemWriter) {
	c.delegates = append(c.delegates, writer)
}

func (c *CompositeItemWriter) Write(items []interface{}) {
	for _, itemWriter := range c.delegates {
		itemWriter.Write(items)
	}
}
