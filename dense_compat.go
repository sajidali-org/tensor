// Code generated by genlib2. DO NOT EDIT.

package tensor

import (
	"fmt"
	"math"
	"math/cmplx"
	"reflect"

	arrow "github.com/apache/arrow/go/arrow"
	arrowArray "github.com/apache/arrow/go/arrow/array"
	arrowTensor "github.com/apache/arrow/go/arrow/tensor"
	"github.com/chewxy/math32"
	"github.com/pkg/errors"
	"gonum.org/v1/gonum/mat"
)

func convFromFloat64s(to Dtype, data []float64) interface{} {
	switch to {
	case Int:
		retVal := make([]int, len(data))
		for i, v := range data {
			switch {
			case math.IsNaN(v), math.IsInf(v, 0):
				retVal[i] = 0
			default:
				retVal[i] = int(v)
			}
		}
		return retVal
	case Int8:
		retVal := make([]int8, len(data))
		for i, v := range data {
			switch {
			case math.IsNaN(v), math.IsInf(v, 0):
				retVal[i] = 0
			default:
				retVal[i] = int8(v)
			}
		}
		return retVal
	case Int16:
		retVal := make([]int16, len(data))
		for i, v := range data {
			switch {
			case math.IsNaN(v), math.IsInf(v, 0):
				retVal[i] = 0
			default:
				retVal[i] = int16(v)
			}
		}
		return retVal
	case Int32:
		retVal := make([]int32, len(data))
		for i, v := range data {
			switch {
			case math.IsNaN(v), math.IsInf(v, 0):
				retVal[i] = 0
			default:
				retVal[i] = int32(v)
			}
		}
		return retVal
	case Int64:
		retVal := make([]int64, len(data))
		for i, v := range data {
			switch {
			case math.IsNaN(v), math.IsInf(v, 0):
				retVal[i] = 0
			default:
				retVal[i] = int64(v)
			}
		}
		return retVal
	case Uint:
		retVal := make([]uint, len(data))
		for i, v := range data {
			switch {
			case math.IsNaN(v), math.IsInf(v, 0):
				retVal[i] = 0
			default:
				retVal[i] = uint(v)
			}
		}
		return retVal
	case Uint8:
		retVal := make([]uint8, len(data))
		for i, v := range data {
			switch {
			case math.IsNaN(v), math.IsInf(v, 0):
				retVal[i] = 0
			default:
				retVal[i] = uint8(v)
			}
		}
		return retVal
	case Uint16:
		retVal := make([]uint16, len(data))
		for i, v := range data {
			switch {
			case math.IsNaN(v), math.IsInf(v, 0):
				retVal[i] = 0
			default:
				retVal[i] = uint16(v)
			}
		}
		return retVal
	case Uint32:
		retVal := make([]uint32, len(data))
		for i, v := range data {
			switch {
			case math.IsNaN(v), math.IsInf(v, 0):
				retVal[i] = 0
			default:
				retVal[i] = uint32(v)
			}
		}
		return retVal
	case Uint64:
		retVal := make([]uint64, len(data))
		for i, v := range data {
			switch {
			case math.IsNaN(v), math.IsInf(v, 0):
				retVal[i] = 0
			default:
				retVal[i] = uint64(v)
			}
		}
		return retVal
	case Float32:
		retVal := make([]float32, len(data))
		for i, v := range data {
			switch {
			case math.IsNaN(v):
				retVal[i] = math32.NaN()
			case math.IsInf(v, 1):
				retVal[i] = math32.Inf(1)
			case math.IsInf(v, -1):
				retVal[i] = math32.Inf(-1)
			default:
				retVal[i] = float32(v)
			}
		}
		return retVal
	case Float64:
		retVal := make([]float64, len(data))
		copy(retVal, data)
		return retVal
	case Complex64:
		retVal := make([]complex64, len(data))
		for i, v := range data {
			switch {
			case math.IsNaN(v):
				retVal[i] = complex64(cmplx.NaN())
			case math.IsInf(v, 0):
				retVal[i] = complex64(cmplx.Inf())
			default:
				retVal[i] = complex(float32(v), float32(0))
			}
		}
		return retVal
	case Complex128:
		retVal := make([]complex128, len(data))
		for i, v := range data {
			switch {
			case math.IsNaN(v):
				retVal[i] = cmplx.NaN()
			case math.IsInf(v, 0):
				retVal[i] = cmplx.Inf()
			default:
				retVal[i] = complex(v, float64(0))
			}
		}
		return retVal
	default:
		panic("Unsupported Dtype")
	}
}

func convToFloat64s(t *Dense) (retVal []float64) {
	retVal = make([]float64, t.len())
	switch t.t {
	case Int:
		for i, v := range t.Ints() {
			retVal[i] = float64(v)
		}
		return retVal
	case Int8:
		for i, v := range t.Int8s() {
			retVal[i] = float64(v)
		}
		return retVal
	case Int16:
		for i, v := range t.Int16s() {
			retVal[i] = float64(v)
		}
		return retVal
	case Int32:
		for i, v := range t.Int32s() {
			retVal[i] = float64(v)
		}
		return retVal
	case Int64:
		for i, v := range t.Int64s() {
			retVal[i] = float64(v)
		}
		return retVal
	case Uint:
		for i, v := range t.Uints() {
			retVal[i] = float64(v)
		}
		return retVal
	case Uint8:
		for i, v := range t.Uint8s() {
			retVal[i] = float64(v)
		}
		return retVal
	case Uint16:
		for i, v := range t.Uint16s() {
			retVal[i] = float64(v)
		}
		return retVal
	case Uint32:
		for i, v := range t.Uint32s() {
			retVal[i] = float64(v)
		}
		return retVal
	case Uint64:
		for i, v := range t.Uint64s() {
			retVal[i] = float64(v)
		}
		return retVal
	case Float32:
		for i, v := range t.Float32s() {
			switch {
			case math32.IsNaN(v):
				retVal[i] = math.NaN()
			case math32.IsInf(v, 1):
				retVal[i] = math.Inf(1)
			case math32.IsInf(v, -1):
				retVal[i] = math.Inf(-1)
			default:
				retVal[i] = float64(v)
			}
		}
		return retVal
	case Float64:
		return t.Float64s()
		return retVal
	case Complex64:
		for i, v := range t.Complex64s() {
			switch {
			case cmplx.IsNaN(complex128(v)):
				retVal[i] = math.NaN()
			case cmplx.IsInf(complex128(v)):
				retVal[i] = math.Inf(1)
			default:
				retVal[i] = float64(real(v))
			}
		}
		return retVal
	case Complex128:
		for i, v := range t.Complex128s() {
			switch {
			case cmplx.IsNaN(v):
				retVal[i] = math.NaN()
			case cmplx.IsInf(v):
				retVal[i] = math.Inf(1)
			default:
				retVal[i] = real(v)
			}
		}
		return retVal
	default:
		panic(fmt.Sprintf("Cannot convert *Dense of %v to []float64", t.t))
	}
}

func convToFloat64(x interface{}) float64 {
	switch xt := x.(type) {
	case int:
		return float64(xt)
	case int8:
		return float64(xt)
	case int16:
		return float64(xt)
	case int32:
		return float64(xt)
	case int64:
		return float64(xt)
	case uint:
		return float64(xt)
	case uint8:
		return float64(xt)
	case uint16:
		return float64(xt)
	case uint32:
		return float64(xt)
	case uint64:
		return float64(xt)
	case float32:
		return float64(xt)
	case float64:
		return float64(xt)
	case complex64:
		return float64(real(xt))
	case complex128:
		return real(xt)
	default:
		panic("Cannot convert to float64")
	}
}

// FromMat64 converts a *"gonum/matrix/mat64".Dense into a *tensorf64.Tensor.
func FromMat64(m *mat.Dense, opts ...FuncOpt) *Dense {
	r, c := m.Dims()
	fo := ParseFuncOpts(opts...)
	defer returnOpOpt(fo)
	toCopy := fo.Safe()
	as := fo.As()
	if as.Type == nil {
		as = Float64
	}

	switch as.Kind() {
	case reflect.Int:
		backing := convFromFloat64s(Int, m.RawMatrix().Data).([]int)
		retVal := New(WithBacking(backing), WithShape(r, c))
		return retVal
	case reflect.Int8:
		backing := convFromFloat64s(Int8, m.RawMatrix().Data).([]int8)
		retVal := New(WithBacking(backing), WithShape(r, c))
		return retVal
	case reflect.Int16:
		backing := convFromFloat64s(Int16, m.RawMatrix().Data).([]int16)
		retVal := New(WithBacking(backing), WithShape(r, c))
		return retVal
	case reflect.Int32:
		backing := convFromFloat64s(Int32, m.RawMatrix().Data).([]int32)
		retVal := New(WithBacking(backing), WithShape(r, c))
		return retVal
	case reflect.Int64:
		backing := convFromFloat64s(Int64, m.RawMatrix().Data).([]int64)
		retVal := New(WithBacking(backing), WithShape(r, c))
		return retVal
	case reflect.Uint:
		backing := convFromFloat64s(Uint, m.RawMatrix().Data).([]uint)
		retVal := New(WithBacking(backing), WithShape(r, c))
		return retVal
	case reflect.Uint8:
		backing := convFromFloat64s(Uint8, m.RawMatrix().Data).([]uint8)
		retVal := New(WithBacking(backing), WithShape(r, c))
		return retVal
	case reflect.Uint16:
		backing := convFromFloat64s(Uint16, m.RawMatrix().Data).([]uint16)
		retVal := New(WithBacking(backing), WithShape(r, c))
		return retVal
	case reflect.Uint32:
		backing := convFromFloat64s(Uint32, m.RawMatrix().Data).([]uint32)
		retVal := New(WithBacking(backing), WithShape(r, c))
		return retVal
	case reflect.Uint64:
		backing := convFromFloat64s(Uint64, m.RawMatrix().Data).([]uint64)
		retVal := New(WithBacking(backing), WithShape(r, c))
		return retVal
	case reflect.Float32:
		backing := convFromFloat64s(Float32, m.RawMatrix().Data).([]float32)
		retVal := New(WithBacking(backing), WithShape(r, c))
		return retVal
	case reflect.Float64:
		var backing []float64
		if toCopy {
			backing = make([]float64, len(m.RawMatrix().Data))
			copy(backing, m.RawMatrix().Data)
		} else {
			backing = m.RawMatrix().Data
		}
		retVal := New(WithBacking(backing), WithShape(r, c))
		return retVal
	case reflect.Complex64:
		backing := convFromFloat64s(Complex64, m.RawMatrix().Data).([]complex64)
		retVal := New(WithBacking(backing), WithShape(r, c))
		return retVal
	case reflect.Complex128:
		backing := convFromFloat64s(Complex128, m.RawMatrix().Data).([]complex128)
		retVal := New(WithBacking(backing), WithShape(r, c))
		return retVal
	default:
		panic(fmt.Sprintf("Unsupported Dtype - cannot convert float64 to %v", as))
	}
	panic("Unreachable")
}

// ToMat64 converts a *Dense to a *mat.Dense. All the values are converted into float64s.
// This function will only convert matrices. Anything *Dense with dimensions larger than 2 will cause an error.
func ToMat64(t *Dense, opts ...FuncOpt) (retVal *mat.Dense, err error) {
	// checks:
	if !t.IsNativelyAccessible() {
		return nil, errors.Errorf("Cannot convert *Dense to *mat.Dense. Data is inaccessible")
	}

	if !t.IsMatrix() {
		// error
		return nil, errors.Errorf("Cannot convert *Dense to *mat.Dense. Expected number of dimensions: <=2, T has got %d dimensions (Shape: %v)", t.Dims(), t.Shape())
	}

	fo := ParseFuncOpts(opts...)
	defer returnOpOpt(fo)
	toCopy := fo.Safe()

	// fix dims
	r := t.Shape()[0]
	c := t.Shape()[1]

	var data []float64
	switch {
	case t.t == Float64 && toCopy && !t.IsMaterializable():
		data = make([]float64, t.len())
		copy(data, t.Float64s())
	case !t.IsMaterializable():
		data = convToFloat64s(t)
	default:
		it := newFlatIterator(&t.AP)
		var next int
		for next, err = it.Next(); err == nil; next, err = it.Next() {
			if err = handleNoOp(err); err != nil {
				return
			}
			data = append(data, convToFloat64(t.Get(next)))
		}
		err = nil

	}

	retVal = mat.NewDense(r, c, data)
	return
}

// FromArrowArray converts an "arrow/array".Interface into a Tensor of matching DataType.
func FromArrowArray(a arrowArray.Interface) *Dense {
	a.Retain()
	defer a.Release()

	r := a.Len()

	// TODO(poopoothegorilla): instead of creating bool ValidMask maybe
	// bitmapBytes can be used from arrow API
	mask := make([]bool, r)
	for i := 0; i < r; i++ {
		mask[i] = a.IsNull(i)
	}

	switch a.DataType() {
	case arrow.BinaryTypes.String:
		backing := make([]string, r)
		for i := 0; i < r; i++ {
			backing[i] = a.(*arrowArray.String).Value(i)
		}
		retVal := New(WithBacking(backing, mask), WithShape(r, 1))
		return retVal
	case arrow.FixedWidthTypes.Boolean:
		backing := make([]bool, r)
		for i := 0; i < r; i++ {
			backing[i] = a.(*arrowArray.Boolean).Value(i)
		}
		retVal := New(WithBacking(backing, mask), WithShape(r, 1))
		return retVal
	case arrow.PrimitiveTypes.Int8:
		backing := a.(*arrowArray.Int8).Int8Values()
		retVal := New(WithBacking(backing, mask), WithShape(r, 1))
		return retVal
	case arrow.PrimitiveTypes.Int16:
		backing := a.(*arrowArray.Int16).Int16Values()
		retVal := New(WithBacking(backing, mask), WithShape(r, 1))
		return retVal
	case arrow.PrimitiveTypes.Int32:
		backing := a.(*arrowArray.Int32).Int32Values()
		retVal := New(WithBacking(backing, mask), WithShape(r, 1))
		return retVal
	case arrow.PrimitiveTypes.Int64:
		backing := a.(*arrowArray.Int64).Int64Values()
		retVal := New(WithBacking(backing, mask), WithShape(r, 1))
		return retVal
	case arrow.PrimitiveTypes.Uint8:
		backing := a.(*arrowArray.Uint8).Uint8Values()
		retVal := New(WithBacking(backing, mask), WithShape(r, 1))
		return retVal
	case arrow.PrimitiveTypes.Uint16:
		backing := a.(*arrowArray.Uint16).Uint16Values()
		retVal := New(WithBacking(backing, mask), WithShape(r, 1))
		return retVal
	case arrow.PrimitiveTypes.Uint32:
		backing := a.(*arrowArray.Uint32).Uint32Values()
		retVal := New(WithBacking(backing, mask), WithShape(r, 1))
		return retVal
	case arrow.PrimitiveTypes.Uint64:
		backing := a.(*arrowArray.Uint64).Uint64Values()
		retVal := New(WithBacking(backing, mask), WithShape(r, 1))
		return retVal
	case arrow.PrimitiveTypes.Float32:
		backing := a.(*arrowArray.Float32).Float32Values()
		retVal := New(WithBacking(backing, mask), WithShape(r, 1))
		return retVal
	case arrow.PrimitiveTypes.Float64:
		backing := a.(*arrowArray.Float64).Float64Values()
		retVal := New(WithBacking(backing, mask), WithShape(r, 1))
		return retVal
	default:
		panic(fmt.Sprintf("Unsupported Arrow DataType - %v", a.DataType()))
	}

	panic("Unreachable")
}

// FromArrowTensor converts an "arrow/tensor".Interface into a Tensor of matching DataType.
func FromArrowTensor(a arrowTensor.Interface) *Dense {
	a.Retain()
	defer a.Release()

	var shape []int
	for _, val := range a.Shape() {
		shape = append(shape, int(val))
	}

	switch a.DataType() {
	case arrow.PrimitiveTypes.Int8:
		if !a.IsContiguous() {
			panic("Non-contiguous data is Unsupported")
		}
		backing := a.(*arrowTensor.Int8).Int8Values()
		if a.IsColMajor() {
			return New(WithShape(shape...), AsFortran(backing))
		}

		return New(WithShape(shape...), WithBacking(backing))
	case arrow.PrimitiveTypes.Int16:
		if !a.IsContiguous() {
			panic("Non-contiguous data is Unsupported")
		}
		backing := a.(*arrowTensor.Int16).Int16Values()
		if a.IsColMajor() {
			return New(WithShape(shape...), AsFortran(backing))
		}

		return New(WithShape(shape...), WithBacking(backing))
	case arrow.PrimitiveTypes.Int32:
		if !a.IsContiguous() {
			panic("Non-contiguous data is Unsupported")
		}
		backing := a.(*arrowTensor.Int32).Int32Values()
		if a.IsColMajor() {
			return New(WithShape(shape...), AsFortran(backing))
		}

		return New(WithShape(shape...), WithBacking(backing))
	case arrow.PrimitiveTypes.Int64:
		if !a.IsContiguous() {
			panic("Non-contiguous data is Unsupported")
		}
		backing := a.(*arrowTensor.Int64).Int64Values()
		if a.IsColMajor() {
			return New(WithShape(shape...), AsFortran(backing))
		}

		return New(WithShape(shape...), WithBacking(backing))
	case arrow.PrimitiveTypes.Uint8:
		if !a.IsContiguous() {
			panic("Non-contiguous data is Unsupported")
		}
		backing := a.(*arrowTensor.Uint8).Uint8Values()
		if a.IsColMajor() {
			return New(WithShape(shape...), AsFortran(backing))
		}

		return New(WithShape(shape...), WithBacking(backing))
	case arrow.PrimitiveTypes.Uint16:
		if !a.IsContiguous() {
			panic("Non-contiguous data is Unsupported")
		}
		backing := a.(*arrowTensor.Uint16).Uint16Values()
		if a.IsColMajor() {
			return New(WithShape(shape...), AsFortran(backing))
		}

		return New(WithShape(shape...), WithBacking(backing))
	case arrow.PrimitiveTypes.Uint32:
		if !a.IsContiguous() {
			panic("Non-contiguous data is Unsupported")
		}
		backing := a.(*arrowTensor.Uint32).Uint32Values()
		if a.IsColMajor() {
			return New(WithShape(shape...), AsFortran(backing))
		}

		return New(WithShape(shape...), WithBacking(backing))
	case arrow.PrimitiveTypes.Uint64:
		if !a.IsContiguous() {
			panic("Non-contiguous data is Unsupported")
		}
		backing := a.(*arrowTensor.Uint64).Uint64Values()
		if a.IsColMajor() {
			return New(WithShape(shape...), AsFortran(backing))
		}

		return New(WithShape(shape...), WithBacking(backing))
	case arrow.PrimitiveTypes.Float32:
		if !a.IsContiguous() {
			panic("Non-contiguous data is Unsupported")
		}
		backing := a.(*arrowTensor.Float32).Float32Values()
		if a.IsColMajor() {
			return New(WithShape(shape...), AsFortran(backing))
		}

		return New(WithShape(shape...), WithBacking(backing))
	case arrow.PrimitiveTypes.Float64:
		if !a.IsContiguous() {
			panic("Non-contiguous data is Unsupported")
		}
		backing := a.(*arrowTensor.Float64).Float64Values()
		if a.IsColMajor() {
			return New(WithShape(shape...), AsFortran(backing))
		}

		return New(WithShape(shape...), WithBacking(backing))
	default:
		panic(fmt.Sprintf("Unsupported Arrow DataType - %v", a.DataType()))
	}

	panic("Unreachable")
}
