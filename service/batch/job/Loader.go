package job

import (
	"fileLoader/service/batch/processor"
	"fileLoader/service/batch/reader"
	"fileLoader/service/batch/writer"
	"fileLoader/utils"
	"fmt"
)

type Loader struct {
	name          string
	itemReader    reader.ItemReader
	itemProcessor processor.ItemProcessor
	itemWriter    writer.ItemWriter
	fetchSize     int
	chunkSize     int
	filePath      string
}

func NewLoader(name string, itemReader reader.ItemReader, itemProcessor processor.ItemProcessor, itemWriter writer.ItemWriter) *Loader {
	return &Loader{name: name, itemReader: itemReader, itemProcessor: itemProcessor, itemWriter: itemWriter, fetchSize: 100, chunkSize: 100}
}

func (j *Loader) SetFetchSize(fetchSize int) {
	j.fetchSize = fetchSize
}

func (j *Loader) SetChunkSize(chunkSize int) {
	j.chunkSize = chunkSize
}

func (j *Loader) SetFilePath(filePath string) error {
	j.filePath = filePath
	if j.itemReader != nil {
		err := j.itemReader.SetFilePath(filePath)
		if err != nil {
			return err
		}
	}
	return nil
}

func (j *Loader) Launch() {
	fmt.Println("Start Loader: " + j.name)
	var itemsReaded []interface{}
	var itemsToWrite []interface{}
	var finishProcess bool

	for {

		if finishProcess {
			break
		}

		for i := 0; i < j.fetchSize; i++ {
			var item, err = j.itemReader.Read()
			if err != nil {
				fmt.Printf("No se pudo leer el registro #%d error: %s", i, err)
				continue
			}

			if item == nil {
				finishProcess = true
				break
			}
			itemsReaded = append(itemsReaded, item)
		}

		itemsToWrite = append(itemsToWrite, itemsReaded...)
		itemsReaded = make([]interface{}, 0)

		if len(itemsToWrite) >= j.chunkSize {
			var partitionedItems = utils.PartitionSlice(itemsToWrite, j.chunkSize)
			for i := 0; i < len(partitionedItems); i++ {
				j.write(partitionedItems[i])
			}
			itemsToWrite = make([]interface{}, 0)
		}

	}

	if len(itemsToWrite) > 0 {
		j.write(itemsToWrite)
		itemsToWrite = make([]interface{}, 0)
	}

	fmt.Println("Finish Loader: " + j.name)

}

func (j *Loader) write(itemsToWrite []interface{}) {
	items := make([]interface{}, len(itemsToWrite))

	for i, val := range itemsToWrite {
		items[i] = j.itemProcessor.Process(val)
	}

	j.itemWriter.Write(items)
}
