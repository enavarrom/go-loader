package utils

func PartitionSlice(slice []any, size int) [][]any {
	var result [][]any

	for i := 0; i < len(slice); i += size {
		end := i + size

		// Evita que el índice final supere la longitud del slice
		if end > len(slice) {
			end = len(slice)
		}

		// Particiona el slice y agrega el segmento a la variable result
		partition := slice[i:end]
		result = append(result, partition)
	}

	return result
}
