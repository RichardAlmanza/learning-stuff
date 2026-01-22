package vector

const (
	Float32 NumberType = 1 << iota
	Float64
	Int
	Int8
	Int16
	Int32
	Int64
	RealNumber
	IntegerNumber
)

type NumberType int

// Real represents the set of real numbers.
type Real interface {
	~float32 | ~float64
}

// Integer represents the set of integers.
type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

func AssertNumber(n any) NumberType {
	var result NumberType

	switch n.(type) {
	case float32, float64:
		result |= RealNumber
	case int, int8, int16, int32, int64:
		result |= IntegerNumber
	default:
		return 0
	}

	if result&RealNumber != 0 {
		switch n.(type) {
		case float32:
			result |= Float32
		case float64:
			result |= Float64
		}
	} else {
		switch n.(type) {
		case int8:
			result |= Int8
		case int16:
			result |= Int16
		case int32:
			result |= Int32
		case int64:
			result |= Int64
		case int:
			result |= Int
		}
	}

	return result
}

// Number represents the set of real and integer numbers.
type Number interface {
	Real | Integer
}

// Shape represents the dimensions of a tensor or vector.
type Shape []int

func (s Shape) TotalSize() int {
	// Multiply all dimensions together to get the total size.
	return Reduce(s, 1, func(_, n, p int) int { return n * p })
}

func (s Shape) Copy() Shape {
	return NewCopy(s)
}

func (s Shape) Dims() int {
	return len(s)
}

type Dense interface {
	Shape() Shape
	Type() NumberType
}
