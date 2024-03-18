package mosaic

type SingleWriter[T any] func(primaryID, typeID uint64, data T) error
type SingeReader[T any] func(primaryID, typeID uint64) (*T, error)

type ListWriter[T any] func(primaryID, typeID uint64, data []T) error
type ListReader[T any] func(primaryID, typeID uint64) ([]T, error)
