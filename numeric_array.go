// Code generated by erb. DO NOT EDIT.

package pgtype

import (
	"database/sql/driver"
	"encoding/binary"
	"reflect"

	"github.com/jackc/pgio"
	errors "golang.org/x/xerrors"
)

type NumericArray struct {
	Elements   []Numeric
	Dimensions []ArrayDimension
	Status     Status
}

func (dst *NumericArray) Set(src interface{}) error {
	// untyped nil and typed nil interfaces are different
	if src == nil {
		*dst = NumericArray{Status: Null}
		return nil
	}

	if value, ok := src.(interface{ Get() interface{} }); ok {
		value2 := value.Get()
		if value2 != value {
			return dst.Set(value2)
		}
	}

	// Attempt to match to select common types:
	switch value := src.(type) {

	case []float32:
		if value == nil {
			*dst = NumericArray{Status: Null}
		} else if len(value) == 0 {
			*dst = NumericArray{Status: Present}
		} else {
			elements := make([]Numeric, len(value))
			for i := range value {
				if err := elements[i].Set(value[i]); err != nil {
					return err
				}
			}
			*dst = NumericArray{
				Elements:   elements,
				Dimensions: []ArrayDimension{{Length: int32(len(elements)), LowerBound: 1}},
				Status:     Present,
			}
		}

	case []*float32:
		if value == nil {
			*dst = NumericArray{Status: Null}
		} else if len(value) == 0 {
			*dst = NumericArray{Status: Present}
		} else {
			elements := make([]Numeric, len(value))
			for i := range value {
				if err := elements[i].Set(value[i]); err != nil {
					return err
				}
			}
			*dst = NumericArray{
				Elements:   elements,
				Dimensions: []ArrayDimension{{Length: int32(len(elements)), LowerBound: 1}},
				Status:     Present,
			}
		}

	case []float64:
		if value == nil {
			*dst = NumericArray{Status: Null}
		} else if len(value) == 0 {
			*dst = NumericArray{Status: Present}
		} else {
			elements := make([]Numeric, len(value))
			for i := range value {
				if err := elements[i].Set(value[i]); err != nil {
					return err
				}
			}
			*dst = NumericArray{
				Elements:   elements,
				Dimensions: []ArrayDimension{{Length: int32(len(elements)), LowerBound: 1}},
				Status:     Present,
			}
		}

	case []*float64:
		if value == nil {
			*dst = NumericArray{Status: Null}
		} else if len(value) == 0 {
			*dst = NumericArray{Status: Present}
		} else {
			elements := make([]Numeric, len(value))
			for i := range value {
				if err := elements[i].Set(value[i]); err != nil {
					return err
				}
			}
			*dst = NumericArray{
				Elements:   elements,
				Dimensions: []ArrayDimension{{Length: int32(len(elements)), LowerBound: 1}},
				Status:     Present,
			}
		}

	case []int64:
		if value == nil {
			*dst = NumericArray{Status: Null}
		} else if len(value) == 0 {
			*dst = NumericArray{Status: Present}
		} else {
			elements := make([]Numeric, len(value))
			for i := range value {
				if err := elements[i].Set(value[i]); err != nil {
					return err
				}
			}
			*dst = NumericArray{
				Elements:   elements,
				Dimensions: []ArrayDimension{{Length: int32(len(elements)), LowerBound: 1}},
				Status:     Present,
			}
		}

	case []*int64:
		if value == nil {
			*dst = NumericArray{Status: Null}
		} else if len(value) == 0 {
			*dst = NumericArray{Status: Present}
		} else {
			elements := make([]Numeric, len(value))
			for i := range value {
				if err := elements[i].Set(value[i]); err != nil {
					return err
				}
			}
			*dst = NumericArray{
				Elements:   elements,
				Dimensions: []ArrayDimension{{Length: int32(len(elements)), LowerBound: 1}},
				Status:     Present,
			}
		}

	case []uint64:
		if value == nil {
			*dst = NumericArray{Status: Null}
		} else if len(value) == 0 {
			*dst = NumericArray{Status: Present}
		} else {
			elements := make([]Numeric, len(value))
			for i := range value {
				if err := elements[i].Set(value[i]); err != nil {
					return err
				}
			}
			*dst = NumericArray{
				Elements:   elements,
				Dimensions: []ArrayDimension{{Length: int32(len(elements)), LowerBound: 1}},
				Status:     Present,
			}
		}

	case []*uint64:
		if value == nil {
			*dst = NumericArray{Status: Null}
		} else if len(value) == 0 {
			*dst = NumericArray{Status: Present}
		} else {
			elements := make([]Numeric, len(value))
			for i := range value {
				if err := elements[i].Set(value[i]); err != nil {
					return err
				}
			}
			*dst = NumericArray{
				Elements:   elements,
				Dimensions: []ArrayDimension{{Length: int32(len(elements)), LowerBound: 1}},
				Status:     Present,
			}
		}

	case []Numeric:
		if value == nil {
			*dst = NumericArray{Status: Null}
		} else if len(value) == 0 {
			*dst = NumericArray{Status: Present}
		} else {
			*dst = NumericArray{
				Elements:   value,
				Dimensions: []ArrayDimension{{Length: int32(len(value)), LowerBound: 1}},
				Status:     Present,
			}
		}
	default:
		// Fallback to reflection if an optimised match was not found.
		// The reflection is necessary for arrays and multidimensional slices,
		// but it comes with a 20-50% performance penalty for large arrays/slices
		reflectedValue := reflect.ValueOf(src)
		if !reflectedValue.IsValid() || reflectedValue.IsZero() {
			*dst = NumericArray{Status: Null}
			return nil
		}

		dimensions, elementsLength, ok := findDimensionsFromValue(reflectedValue, nil, 0)
		if !ok {
			return errors.Errorf("cannot find dimensions of %v for NumericArray", src)
		}
		if elementsLength == 0 {
			*dst = NumericArray{Status: Present}
			return nil
		}
		if len(dimensions) == 0 {
			if originalSrc, ok := underlyingSliceType(src); ok {
				return dst.Set(originalSrc)
			}
			return errors.Errorf("cannot convert %v to NumericArray", src)
		}

		*dst = NumericArray{
			Elements:   make([]Numeric, elementsLength),
			Dimensions: dimensions,
			Status:     Present,
		}
		elementCount, err := dst.setRecursive(reflectedValue, 0, 0)
		if err != nil {
			// Maybe the target was one dimension too far, try again:
			if len(dst.Dimensions) > 1 {
				dst.Dimensions = dst.Dimensions[:len(dst.Dimensions)-1]
				elementsLength = 0
				for _, dim := range dst.Dimensions {
					if elementsLength == 0 {
						elementsLength = int(dim.Length)
					} else {
						elementsLength *= int(dim.Length)
					}
				}
				dst.Elements = make([]Numeric, elementsLength)
				elementCount, err = dst.setRecursive(reflectedValue, 0, 0)
				if err != nil {
					return err
				}
			} else {
				return err
			}
		}
		if elementCount != len(dst.Elements) {
			return errors.Errorf("cannot convert %v to NumericArray, expected %d dst.Elements, but got %d instead", src, len(dst.Elements), elementCount)
		}
	}

	return nil
}

func (dst *NumericArray) setRecursive(value reflect.Value, index, dimension int) (int, error) {
	switch value.Kind() {
	case reflect.Array:
		fallthrough
	case reflect.Slice:
		if len(dst.Dimensions) == dimension {
			break
		}

		valueLen := value.Len()
		if int32(valueLen) != dst.Dimensions[dimension].Length {
			return 0, errors.Errorf("multidimensional arrays must have array expressions with matching dimensions")
		}
		for i := 0; i < valueLen; i++ {
			var err error
			index, err = dst.setRecursive(value.Index(i), index, dimension+1)
			if err != nil {
				return 0, err
			}
		}

		return index, nil
	}
	if !value.CanInterface() {
		return 0, errors.Errorf("cannot convert all values to NumericArray")
	}
	if err := dst.Elements[index].Set(value.Interface()); err != nil {
		return 0, errors.Errorf("%v in NumericArray", err)
	}
	index++

	return index, nil
}

func (dst NumericArray) Get() interface{} {
	switch dst.Status {
	case Present:
		return dst
	case Null:
		return nil
	default:
		return dst.Status
	}
}

func (src *NumericArray) AssignTo(dst interface{}) error {
	switch src.Status {
	case Present:
		if len(src.Dimensions) == 1 {
			// Attempt to match to select common types:
			switch v := dst.(type) {

			case *[]float32:
				*v = make([]float32, len(src.Elements))
				for i := range src.Elements {
					if err := src.Elements[i].AssignTo(&((*v)[i])); err != nil {
						return err
					}
				}
				return nil

			case *[]*float32:
				*v = make([]*float32, len(src.Elements))
				for i := range src.Elements {
					if err := src.Elements[i].AssignTo(&((*v)[i])); err != nil {
						return err
					}
				}
				return nil

			case *[]float64:
				*v = make([]float64, len(src.Elements))
				for i := range src.Elements {
					if err := src.Elements[i].AssignTo(&((*v)[i])); err != nil {
						return err
					}
				}
				return nil

			case *[]*float64:
				*v = make([]*float64, len(src.Elements))
				for i := range src.Elements {
					if err := src.Elements[i].AssignTo(&((*v)[i])); err != nil {
						return err
					}
				}
				return nil

			case *[]int64:
				*v = make([]int64, len(src.Elements))
				for i := range src.Elements {
					if err := src.Elements[i].AssignTo(&((*v)[i])); err != nil {
						return err
					}
				}
				return nil

			case *[]*int64:
				*v = make([]*int64, len(src.Elements))
				for i := range src.Elements {
					if err := src.Elements[i].AssignTo(&((*v)[i])); err != nil {
						return err
					}
				}
				return nil

			case *[]uint64:
				*v = make([]uint64, len(src.Elements))
				for i := range src.Elements {
					if err := src.Elements[i].AssignTo(&((*v)[i])); err != nil {
						return err
					}
				}
				return nil

			case *[]*uint64:
				*v = make([]*uint64, len(src.Elements))
				for i := range src.Elements {
					if err := src.Elements[i].AssignTo(&((*v)[i])); err != nil {
						return err
					}
				}
				return nil

			}
		}

		// Fallback to reflection if an optimised match was not found.
		// The reflection is necessary for arrays and multidimensional slices,
		// but it comes with a 20-50% performance penalty for large arrays/slices
		value := reflect.ValueOf(dst)
		if value.Kind() == reflect.Ptr {
			value = value.Elem()
		}
		if !value.CanSet() {
			if nextDst, retry := GetAssignToDstType(dst); retry {
				return src.AssignTo(nextDst)
			}
			return errors.Errorf("unable to assign to %T", dst)
		}

		elementCount, err := src.assignToRecursive(value, 0, 0)
		if err != nil {
			return err
		}
		if elementCount != len(src.Elements) {
			return errors.Errorf("cannot assign %v, needed to assign %d elements, but only assigned %d", dst, len(src.Elements), elementCount)
		}

		return nil
	case Null:
		return NullAssignTo(dst)
	}

	return errors.Errorf("cannot decode %#v into %T", src, dst)
}

func (src *NumericArray) assignToRecursive(value reflect.Value, index, dimension int) (int, error) {
	switch kind := value.Kind(); kind {
	case reflect.Array:
		fallthrough
	case reflect.Slice:
		if len(src.Dimensions) == dimension {
			break
		}

		length := int(src.Dimensions[dimension].Length)
		if reflect.Array == kind {
			typ := value.Type()
			if typ.Len() != length {
				return 0, errors.Errorf("expected size %d array, but %s has size %d array", length, typ, typ.Len())
			}
			value.Set(reflect.New(typ).Elem())
		} else {
			value.Set(reflect.MakeSlice(value.Type(), length, length))
		}

		var err error
		for i := 0; i < length; i++ {
			index, err = src.assignToRecursive(value.Index(i), index, dimension+1)
			if err != nil {
				return 0, err
			}
		}

		return index, nil
	}
	if len(src.Dimensions) != dimension {
		return 0, errors.Errorf("incorrect dimensions, expected %d, found %d", len(src.Dimensions), dimension)
	}
	if !value.CanAddr() {
		return 0, errors.Errorf("cannot assign all values from NumericArray")
	}
	addr := value.Addr()
	if !addr.CanInterface() {
		return 0, errors.Errorf("cannot assign all values from NumericArray")
	}
	if err := src.Elements[index].AssignTo(addr.Interface()); err != nil {
		return 0, err
	}
	index++
	return index, nil
}

func (dst *NumericArray) DecodeText(ci *ConnInfo, src []byte) error {
	if src == nil {
		*dst = NumericArray{Status: Null}
		return nil
	}

	uta, err := ParseUntypedTextArray(string(src))
	if err != nil {
		return err
	}

	var elements []Numeric

	if len(uta.Elements) > 0 {
		elements = make([]Numeric, len(uta.Elements))

		for i, s := range uta.Elements {
			var elem Numeric
			var elemSrc []byte
			if s != "NULL" {
				elemSrc = []byte(s)
			}
			err = elem.DecodeText(ci, elemSrc)
			if err != nil {
				return err
			}

			elements[i] = elem
		}
	}

	*dst = NumericArray{Elements: elements, Dimensions: uta.Dimensions, Status: Present}

	return nil
}

func (dst *NumericArray) DecodeBinary(ci *ConnInfo, src []byte) error {
	if src == nil {
		*dst = NumericArray{Status: Null}
		return nil
	}

	var arrayHeader ArrayHeader
	rp, err := arrayHeader.DecodeBinary(ci, src)
	if err != nil {
		return err
	}

	if len(arrayHeader.Dimensions) == 0 {
		*dst = NumericArray{Dimensions: arrayHeader.Dimensions, Status: Present}
		return nil
	}

	elementCount := arrayHeader.Dimensions[0].Length
	for _, d := range arrayHeader.Dimensions[1:] {
		elementCount *= d.Length
	}

	elements := make([]Numeric, elementCount)

	for i := range elements {
		elemLen := int(int32(binary.BigEndian.Uint32(src[rp:])))
		rp += 4
		var elemSrc []byte
		if elemLen >= 0 {
			elemSrc = src[rp : rp+elemLen]
			rp += elemLen
		}
		err = elements[i].DecodeBinary(ci, elemSrc)
		if err != nil {
			return err
		}
	}

	*dst = NumericArray{Elements: elements, Dimensions: arrayHeader.Dimensions, Status: Present}
	return nil
}

func (src NumericArray) EncodeText(ci *ConnInfo, buf []byte) ([]byte, error) {
	switch src.Status {
	case Null:
		return nil, nil
	case Undefined:
		return nil, errUndefined
	}

	if len(src.Dimensions) == 0 {
		return append(buf, '{', '}'), nil
	}

	buf = EncodeTextArrayDimensions(buf, src.Dimensions)

	// dimElemCounts is the multiples of elements that each array lies on. For
	// example, a single dimension array of length 4 would have a dimElemCounts of
	// [4]. A multi-dimensional array of lengths [3,5,2] would have a
	// dimElemCounts of [30,10,2]. This is used to simplify when to render a '{'
	// or '}'.
	dimElemCounts := make([]int, len(src.Dimensions))
	dimElemCounts[len(src.Dimensions)-1] = int(src.Dimensions[len(src.Dimensions)-1].Length)
	for i := len(src.Dimensions) - 2; i > -1; i-- {
		dimElemCounts[i] = int(src.Dimensions[i].Length) * dimElemCounts[i+1]
	}

	inElemBuf := make([]byte, 0, 32)
	for i, elem := range src.Elements {
		if i > 0 {
			buf = append(buf, ',')
		}

		for _, dec := range dimElemCounts {
			if i%dec == 0 {
				buf = append(buf, '{')
			}
		}

		elemBuf, err := elem.EncodeText(ci, inElemBuf)
		if err != nil {
			return nil, err
		}
		if elemBuf == nil {
			buf = append(buf, `NULL`...)
		} else {
			buf = append(buf, QuoteArrayElementIfNeeded(string(elemBuf))...)
		}

		for _, dec := range dimElemCounts {
			if (i+1)%dec == 0 {
				buf = append(buf, '}')
			}
		}
	}

	return buf, nil
}

func (src NumericArray) EncodeBinary(ci *ConnInfo, buf []byte) ([]byte, error) {
	switch src.Status {
	case Null:
		return nil, nil
	case Undefined:
		return nil, errUndefined
	}

	arrayHeader := ArrayHeader{
		Dimensions: src.Dimensions,
	}

	if dt, ok := ci.DataTypeForName("numeric"); ok {
		arrayHeader.ElementOID = int32(dt.OID)
	} else {
		return nil, errors.Errorf("unable to find oid for type name %v", "numeric")
	}

	for i := range src.Elements {
		if src.Elements[i].Status == Null {
			arrayHeader.ContainsNull = true
			break
		}
	}

	buf = arrayHeader.EncodeBinary(ci, buf)

	for i := range src.Elements {
		sp := len(buf)
		buf = pgio.AppendInt32(buf, -1)

		elemBuf, err := src.Elements[i].EncodeBinary(ci, buf)
		if err != nil {
			return nil, err
		}
		if elemBuf != nil {
			buf = elemBuf
			pgio.SetInt32(buf[sp:], int32(len(buf[sp:])-4))
		}
	}

	return buf, nil
}

// Scan implements the database/sql Scanner interface.
func (dst *NumericArray) Scan(src interface{}) error {
	if src == nil {
		return dst.DecodeText(nil, nil)
	}

	switch src := src.(type) {
	case string:
		return dst.DecodeText(nil, []byte(src))
	case []byte:
		srcCopy := make([]byte, len(src))
		copy(srcCopy, src)
		return dst.DecodeText(nil, srcCopy)
	}

	return errors.Errorf("cannot scan %T", src)
}

// Value implements the database/sql/driver Valuer interface.
func (src NumericArray) Value() (driver.Value, error) {
	buf, err := src.EncodeText(nil, nil)
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}

	return string(buf), nil
}
