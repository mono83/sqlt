package sqlt

// GenericReader defines accessor to database providing
// ability to read single entity of type [T] using
// identifier with type [I]
type GenericReader[I, T any] func(ID I) (*T, error)

// GenericListReader defines accessor to database providing
// ability to read list(slice) of records of type [T]
// using slice of identifiers of type [I]
type GenericListReader[I, T any] func(ID ...I) ([]T, error)
