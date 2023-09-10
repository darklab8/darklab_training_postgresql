package utils

const (
	PostgresqlSerialMax = 2147483647
)

func GetNext(number *int) int {
	*number++
	return *number
}
