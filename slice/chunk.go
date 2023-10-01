package slice

func Chunk[T any](original []T, chunkSize int) [][]T {
	if chunkSize <= 0 {
		return [][]T{}
	}

	chunks := [][]T{}

	var end int
	for i := 0; i < len(original); i += chunkSize {
		end = i + chunkSize

		if end > len(original) {
			chunks = append(chunks, original[i:])
			break
		}

		chunks = append(chunks, original[i:end])
	}

	return chunks
}
