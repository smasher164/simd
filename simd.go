package simd

import (
	"math"
	"unsafe"
)

type V128 struct {
	b [16]byte
}

type i8x16 [16]int8
type i16x8 [8]int16
type i32x4 [4]int32
type i64x2 [2]int64
type u8x16 [16]uint8
type u16x8 [8]uint16
type u32x4 [4]uint32
type u64x2 [2]uint64
type f32x4 [4]float32
type f64x2 [2]float64

func I8x16(i1, i2, i3, i4, i5, i6, i7, i8, i9, i10, i11, i12, i13, i14, i15, i16 int8) (v V128) {
	*(*i8x16)(unsafe.Pointer(&v)) = i8x16{i1, i2, i3, i4, i5, i6, i7, i8, i9, i10, i11, i12, i13, i14, i15, i16}
	return
}

func I16x8(i1, i2, i3, i4, i5, i6, i7, i8 int16) (v V128) {
	*(*i16x8)(unsafe.Pointer(&v)) = i16x8{i1, i2, i3, i4, i5, i6, i7, i8}
	return
}

func I32x4(i1, i2, i3, i4 int32) (v V128) {
	*(*i32x4)(unsafe.Pointer(&v)) = i32x4{i1, i2, i3, i4}
	return
}

func I64x2(i1, i2 int64) (v V128) {
	*(*i64x2)(unsafe.Pointer(&v)) = i64x2{i1, i2}
	return
}

func U8x16(u1, u2, u3, u4, u5, u6, u7, u8, u9, u10, u11, u12, u13, u14, u15, u16 uint8) (v V128) {
	*(*u8x16)(unsafe.Pointer(&v)) = u8x16{u1, u2, u3, u4, u5, u6, u7, u8, u9, u10, u11, u12, u13, u14, u15, u16}
	return
}

func U16x8(u1, u2, u3, u4, u5, u6, u7, u8 uint16) (v V128) {
	*(*u16x8)(unsafe.Pointer(&v)) = u16x8{u1, u2, u3, u4, u5, u6, u7, u8}
	return
}

func U32x4(u1, u2, u3, u4 uint32) V128 {
	*(*u32x4)(unsafe.Pointer(&v)) = u32x4{u1, u2, u3, u4}
	return
}

func U64x2(u1, u2 uint64) (v V128) {
	*(*u64x2)(unsafe.Pointer(&v)) = u64x2{u1, u2}
	return
}

func F32x4(f1, f2, f3, f4 float32) (v V128) {
	*(*f32x4)(unsafe.Pointer(&v)) = f32x4{f1, f2, f3, f4}
	return
}

func F64x2(f1, f2 float64) (v V128) {
	*(*f64x2)(unsafe.Pointer(&v)) = f64x2{f1, f2}
	return
}

func SplatI8(x int8) V128 {
	var a i8x16
	for i := 0; i < len(a); i++ {
		a[i] = x
	}
	return *(*V128)(unsafe.Pointer(&a))
}

func SplatI16(x int16) V128 {
	var a i16x8
	for i := 0; i < len(a); i++ {
		a[i] = x
	}
	return *(*V128)(unsafe.Pointer(&a))
}

func SplatI32(x int32) V128 {
	var a i32x4
	for i := 0; i < len(a); i++ {
		a[i] = x
	}
	return *(*V128)(unsafe.Pointer(&a))
}

func SplatI64(x int64) V128 {
	var a i64x2
	for i := 0; i < len(a); i++ {
		a[i] = x
	}
	return *(*V128)(unsafe.Pointer(&a))
}

func SplatU8(x uint8) V128 {
	var a u8x16
	for i := 0; i < len(a); i++ {
		a[i] = x
	}
	return *(*V128)(unsafe.Pointer(&a))
}

func SplatU16(x uint16) V128 {
	var a u16x8
	for i := 0; i < len(a); i++ {
		a[i] = x
	}
	return *(*V128)(unsafe.Pointer(&a))
}

func SplatU32(x uint32) (v V128) {
	var a u32x4
	for i := 0; i < len(a); i++ {
		a[i] = x
	}
	return *(*V128)(unsafe.Pointer(&a))
}

func SplatU64(x uint64) V128 {
	var a u64x2
	for i := 0; i < len(a); i++ {
		a[i] = x
	}
	return *(*V128)(unsafe.Pointer(&a))
}

func SplatF32(x float32) V128 {
	var a f32x4
	for i := 0; i < len(a); i++ {
		a[i] = x
	}
	return *(*V128)(unsafe.Pointer(&a))
}

func SplatF64(x float64) V128 {
	var a f64x2
	for i := 0; i < len(a); i++ {
		a[i] = x
	}
	return *(*V128)(unsafe.Pointer(&a))
}

func ExtractLaneI8(v V128, i int) int8     { return (*(*i8x16)(unsafe.Pointer(&v)))[i] }
func ExtractLaneI16(v V128, i int) int16   { return (*(*i16x8)(unsafe.Pointer(&v)))[i] }
func ExtractLaneI32(v V128, i int) int32   { return (*(*i32x4)(unsafe.Pointer(&v)))[i] }
func ExtractLaneI64(v V128, i int) int64   { return (*(*i64x2)(unsafe.Pointer(&v)))[i] }
func ExtractLaneU8(v V128, i int) uint8    { return (*(*u8x16)(unsafe.Pointer(&v)))[i] }
func ExtractLaneU16(v V128, i int) uint16  { return (*(*u16x8)(unsafe.Pointer(&v)))[i] }
func ExtractLaneU32(v V128, i int) uint32  { return (*(*u32x4)(unsafe.Pointer(&v)))[i] }
func ExtractLaneU64(v V128, i int) uint64  { return (*(*u64x2)(unsafe.Pointer(&v)))[i] }
func ExtractLaneF32(v V128, i int) float32 { return (*(*f32x4)(unsafe.Pointer(&v)))[i] }
func ExtractLaneF64(v V128, i int) float64 { return (*(*f64x2)(unsafe.Pointer(&v)))[i] }

func ReplaceLaneI8(v V128, i int, x int8) V128 {
	*(*i8x16)(unsafe.Pointer(&v))[i] = x
	return v
}
func ReplaceLaneI16(v V128, i int, x int16) V128 {
	*(*i16x8)(unsafe.Pointer(&v))[i] = x
	return v
}
func ReplaceLaneI32(v V128, i int, x int32) V128 {
	*(*i32x4)(unsafe.Pointer(&v))[i] = x
	return v
}
func ReplaceLaneI64(v V128, i int, x int64) V128 {
	*(*i64x2)(unsafe.Pointer(&v))[i] = x
	return v
}
func ReplaceLaneU8(v V128, i int, x uint8) V128 {
	*(*u8x16)(unsafe.Pointer(&v))[i] = x
	return v
}
func ReplaceLaneU16(v V128, i int, x uint16) V128 {
	*(*u16x8)(unsafe.Pointer(&v))[i] = x
	return v
}
func ReplaceLaneU32(v V128, i int, x uint32) V128 {
	*(*u32x4)(unsafe.Pointer(&v))[i] = x
	return v
}
func ReplaceLaneU64(v V128, i int, x uint64) V128 {
	*(*u64x2)(unsafe.Pointer(&v))[i] = x
	return v
}
func ReplaceLaneF32(v V128, i int, x float32) V128 {
	*(*f32x4)(unsafe.Pointer(&v))[i] = x
	return v
}
func ReplaceLaneF64(v V128, i int, x float64) V128 {
	*(*f64x2)(unsafe.Pointer(&v))[i] = x
	return v
}

func Shuffle(a, b V128, s [16]int) (v V128) {
	for i := 0; i < 16; i++ {
		if s[i] < 16 {
			v.b[i] = a[s[i]]
		} else {
			v.b[i] = b[s[i]] - 16
		}
	}
	return v
}

func Swizzle(a, s V128) (v V128) {
	for i := 0; i < 16; i++ {
		if s.b[i] < 16 {
			v.b[i] = a[s[i]]
		}
	}
	return v
}

func AddI8(a, b V128) V128 {
	var r i8x16
	x := *(*i8x16)(unsafe.Pointer(&a))
	y := *(*i8x16)(unsafe.Pointer(&b))
	for i := range r {
		r[i] = x[i] + y[i]
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func AddI16(a, b V128) V128 {
	var r i16x8
	x := *(*i16x8)(unsafe.Pointer(&a))
	y := *(*i16x8)(unsafe.Pointer(&b))
	for i := range r {
		r[i] = x[i] + y[i]
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func AddI32(a, b V128) V128 {
	var r i32x4
	x := *(*i32x4)(unsafe.Pointer(&a))
	y := *(*i32x4)(unsafe.Pointer(&b))
	for i := range r {
		r[i] = x[i] + y[i]
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func AddI64(a, b V128) V128 {
	var r i64x2
	x := *(*i64x2)(unsafe.Pointer(&a))
	y := *(*i64x2)(unsafe.Pointer(&b))
	for i := range r {
		r[i] = x[i] + y[i]
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func AddU8(a, b V128) V128 {
	var r u8x16
	x := *(*u8x16)(unsafe.Pointer(&a))
	y := *(*u8x16)(unsafe.Pointer(&b))
	for i := range r {
		r[i] = x[i] + y[i]
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func AddU16(a, b V128) V128 {
	var r u16x8
	x := *(*u16x8)(unsafe.Pointer(&a))
	y := *(*u16x8)(unsafe.Pointer(&b))
	for i := range r {
		r[i] = x[i] + y[i]
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func AddU32(a, b V128) V128 {
	var r u32x4
	x := *(*u32x4)(unsafe.Pointer(&a))
	y := *(*u32x4)(unsafe.Pointer(&b))
	for i := range r {
		r[i] = x[i] + y[i]
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func AddU64(a, b V128) V128 {
	var r u64x2
	x := *(*u64x2)(unsafe.Pointer(&a))
	y := *(*u64x2)(unsafe.Pointer(&b))
	for i := range r {
		r[i] = x[i] + y[i]
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func AddF32(a, b V128) V128 {
	var r f32x4
	x := *(*f32x4)(unsafe.Pointer(&a))
	y := *(*f32x4)(unsafe.Pointer(&b))
	for i := range r {
		r[i] = x[i] + y[i]
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func AddF64(a, b V128) V128 {
	var r f64x2
	x := *(*f64x2)(unsafe.Pointer(&a))
	y := *(*f64x2)(unsafe.Pointer(&b))
	for i := range r {
		r[i] = x[i] + y[i]
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func SubI8(a, b V128) V128 {
	var r i8x16
	x := *(*i8x16)(unsafe.Pointer(&a))
	y := *(*i8x16)(unsafe.Pointer(&b))
	for i := range r {
		r[i] = x[i] - y[i]
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func SubI16(a, b V128) V128 {
	var r i16x8
	x := *(*i16x8)(unsafe.Pointer(&a))
	y := *(*i16x8)(unsafe.Pointer(&b))
	for i := range r {
		r[i] = x[i] - y[i]
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func SubI32(a, b V128) V128 {
	var r i32x4
	x := *(*i32x4)(unsafe.Pointer(&a))
	y := *(*i32x4)(unsafe.Pointer(&b))
	for i := range r {
		r[i] = x[i] - y[i]
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func SubI64(a, b V128) V128 {
	var r i64x2
	x := *(*i64x2)(unsafe.Pointer(&a))
	y := *(*i64x2)(unsafe.Pointer(&b))
	for i := range r {
		r[i] = x[i] - y[i]
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func SubU8(a, b V128) V128 {
	var r u8x16
	x := *(*u8x16)(unsafe.Pointer(&a))
	y := *(*u8x16)(unsafe.Pointer(&b))
	for i := range r {
		r[i] = x[i] - y[i]
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func SubU16(a, b V128) V128 {
	var r u16x8
	x := *(*u16x8)(unsafe.Pointer(&a))
	y := *(*u16x8)(unsafe.Pointer(&b))
	for i := range r {
		r[i] = x[i] - y[i]
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func SubU32(a, b V128) V128 {
	var r u32x4
	x := *(*u32x4)(unsafe.Pointer(&a))
	y := *(*u32x4)(unsafe.Pointer(&b))
	for i := range r {
		r[i] = x[i] - y[i]
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func SubU64(a, b V128) V128 {
	var r u64x2
	x := *(*u64x2)(unsafe.Pointer(&a))
	y := *(*u64x2)(unsafe.Pointer(&b))
	for i := range r {
		r[i] = x[i] - y[i]
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func SubF32(a, b V128) V128 {
	var r f32x4
	x := *(*f32x4)(unsafe.Pointer(&a))
	y := *(*f32x4)(unsafe.Pointer(&b))
	for i := range r {
		r[i] = x[i] - y[i]
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func SubF64(a, b V128) V128 {
	var r f64x2
	x := *(*f64x2)(unsafe.Pointer(&a))
	y := *(*f64x2)(unsafe.Pointer(&b))
	for i := range r {
		r[i] = x[i] - y[i]
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func MulI8(a, b V128) V128 {
	var r i8x16
	x := *(*i8x16)(unsafe.Pointer(&a))
	y := *(*i8x16)(unsafe.Pointer(&b))
	for i := range r {
		r[i] = x[i] + y[i]
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func MulI16(a, b V128) V128 {
	var r i16x8
	x := *(*i16x8)(unsafe.Pointer(&a))
	y := *(*i16x8)(unsafe.Pointer(&b))
	for i := range r {
		r[i] = x[i] + y[i]
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func MulI32(a, b V128) V128 {
	var r i32x4
	x := *(*i32x4)(unsafe.Pointer(&a))
	y := *(*i32x4)(unsafe.Pointer(&b))
	for i := range r {
		r[i] = x[i] + y[i]
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func MulI64(a, b V128) V128 {
	var r i64x2
	x := *(*i64x2)(unsafe.Pointer(&a))
	y := *(*i64x2)(unsafe.Pointer(&b))
	for i := range r {
		r[i] = x[i] + y[i]
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func MulU8(a, b V128) V128 {
	var r u8x16
	x := *(*u8x16)(unsafe.Pointer(&a))
	y := *(*u8x16)(unsafe.Pointer(&b))
	for i := range r {
		r[i] = x[i] * y[i]
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func MulU16(a, b V128) V128 {
	var r u16x8
	x := *(*u16x8)(unsafe.Pointer(&a))
	y := *(*u16x8)(unsafe.Pointer(&b))
	for i := range r {
		r[i] = x[i] * y[i]
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func MulU32(a, b V128) V128 {
	var r u32x4
	x := *(*u32x4)(unsafe.Pointer(&a))
	y := *(*u32x4)(unsafe.Pointer(&b))
	for i := range r {
		r[i] = x[i] * y[i]
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func MulU64(a, b V128) V128 {
	var r u64x2
	x := *(*u64x2)(unsafe.Pointer(&a))
	y := *(*u64x2)(unsafe.Pointer(&b))
	for i := range r {
		r[i] = x[i] * y[i]
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func MulF32(a, b V128) V128 {
	var r f32x4
	x := *(*f32x4)(unsafe.Pointer(&a))
	y := *(*f32x4)(unsafe.Pointer(&b))
	for i := range r {
		r[i] = x[i] * y[i]
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func MulF64(a, b V128) V128 {
	var r f64x2
	x := *(*f64x2)(unsafe.Pointer(&a))
	y := *(*f64x2)(unsafe.Pointer(&b))
	for i := range r {
		r[i] = x[i] * y[i]
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func DivF32(a, b V128) V128 {
	var r f32x4
	x := *(*f32x4)(unsafe.Pointer(&a))
	y := *(*f32x4)(unsafe.Pointer(&b))
	for i := range r {
		r[i] = x[i] / y[i]
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func DivF64(a, b V128) V128 {
	var r f64x2
	x := *(*f64x2)(unsafe.Pointer(&a))
	y := *(*f64x2)(unsafe.Pointer(&b))
	for i := range r {
		r[i] = x[i] / y[i]
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func SaturatedAddI8(a, b V128) V128 {
	var r i8x16
	x := *(*i8x16)(unsafe.Pointer(&a))
	y := *(*i8x16)(unsafe.Pointer(&b))
	for i := range r {
		sum := int16(x[i]) + int16(y[i])
		if sum < -128 {
			r[i] = -128
		} else if sum > 127 {
			r[i] = 127
		} else {
			r[i] = int8(sum)
		}
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func SaturatedAddI16(a, b V128) V128 {
	var r i16x8
	x := *(*i16x8)(unsafe.Pointer(&a))
	y := *(*i16x8)(unsafe.Pointer(&b))
	for i := range r {
		sum := int32(x[i]) + int32(y[i])
		if sum < -32768 {
			r[i] = -32768
		} else if sum > 32767 {
			r[i] = 32767
		} else {
			r[i] = int16(sum)
		}
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func SaturatedAddI32(a, b V128) V128 {
	var r i32x4
	x := *(*i32x4)(unsafe.Pointer(&a))
	y := *(*i32x4)(unsafe.Pointer(&b))
	for i := range r {
		sum := int64(x[i]) + int64(y[i])
		if sum < -2147483648 {
			r[i] = -2147483648
		} else if sum > 2147483647 {
			r[i] = 2147483647
		} else {
			r[i] = int32(sum)
		}
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func SaturatedAddI64(a, b V128) V128 {
	var r i64x2
	x := *(*i64x2)(unsafe.Pointer(&a))
	y := *(*i64x2)(unsafe.Pointer(&b))
	for i := range r {
		ux, uy := uint64(x[i]), uint64(y[i])
		usum := ux + uy
		ux = (usum >> 63) + 9223372036854775807
		if int64(ux^uy)|^(uy^usum) >= 0 {
			usum = ux
		}
		r[i] = int64(usum)
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func SaturatedAddU8(a, b V128) V128 {
	var r u8x16
	x := *(*u8x16)(unsafe.Pointer(&a))
	y := *(*u8x16)(unsafe.Pointer(&b))
	for i := range r {
		if sum := x[i] + y[i]; sum < x[i] {
			r[i] = 255
		} else {
			r[i] = sum
		}
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func SaturatedAddU16(a, b V128) V128 {
	var r u16x8
	x := *(*u16x8)(unsafe.Pointer(&a))
	y := *(*u16x8)(unsafe.Pointer(&b))
	for i := range r {
		if sum := x[i] + y[i]; sum < x[i] {
			r[i] = 65535
		} else {
			r[i] = sum
		}
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func SaturatedAddU32(a, b V128) V128 {
	var r u32x4
	x := *(*u32x4)(unsafe.Pointer(&a))
	y := *(*u32x4)(unsafe.Pointer(&b))
	for i := range r {
		if sum := x[i] + y[i]; sum < x[i] {
			r[i] = 4294967295
		} else {
			r[i] = sum
		}
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func SaturatedAddU64(a, b V128) V128 {
	var r u64x2
	x := *(*u64x2)(unsafe.Pointer(&a))
	y := *(*u64x2)(unsafe.Pointer(&b))
	for i := range r {
		if sum := x[i] + y[i]; sum < x[i] {
			r[i] = 18446744073709551615
		} else {
			r[i] = sum
		}
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func SaturatedSubI8(a, b V128) V128 {
	var r i8x16
	x := *(*i8x16)(unsafe.Pointer(&a))
	y := *(*i8x16)(unsafe.Pointer(&b))
	for i := range r {
		diff := int16(x[i]) - int16(y[i])
		if diff < -128 {
			r[i] = -128
		} else if diff > 127 {
			r[i] = 127
		} else {
			r[i] = int8(diff)
		}
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func SaturatedSubI16(a, b V128) V128 {
	var r i16x8
	x := *(*i16x8)(unsafe.Pointer(&a))
	y := *(*i16x8)(unsafe.Pointer(&b))
	for i := range r {
		diff := int32(x[i]) - int32(y[i])
		if diff < -32768 {
			r[i] = -32768
		} else if diff > 32767 {
			r[i] = 32767
		} else {
			r[i] = int16(diff)
		}
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func SaturatedSubI32(a, b V128) V128 {
	var r i32x4
	x := *(*i32x4)(unsafe.Pointer(&a))
	y := *(*i32x4)(unsafe.Pointer(&b))
	for i := range r {
		diff := int64(x[i]) - int64(y[i])
		if diff < -2147483648 {
			r[i] = -2147483648
		} else if diff > 2147483647 {
			r[i] = 2147483647
		} else {
			r[i] = int32(diff)
		}
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func SaturatedSubI64(a, b V128) V128 {
	var r i64x2
	x := *(*i64x2)(unsafe.Pointer(&a))
	y := *(*i64x2)(unsafe.Pointer(&b))
	for i := range r {
		ux, uy := uint64(x[i]), uint64(y[i])
		udiff := ux - uy
		ux = (udiff >> 63) + 9223372036854775807
		if int64(ux^uy)&(ux^udiff) < 0 {
			udiff = ux
		}
		r[i] = int64(udiff)
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func SaturatedSubU8(a, b V128) V128 {
	var r u8x16
	x := *(*u8x16)(unsafe.Pointer(&a))
	y := *(*u8x16)(unsafe.Pointer(&b))
	for i := range r {
		if diff := x[i] - y[i]; diff > x[i] {
			r[i] = 0
		} else {
			r[i] = diff
		}
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func SaturatedSubU16(a, b V128) V128 {
	var r u16x8
	x := *(*u16x8)(unsafe.Pointer(&a))
	y := *(*u16x8)(unsafe.Pointer(&b))
	for i := range r {
		if diff := x[i] - y[i]; diff > x[i] {
			r[i] = 0
		} else {
			r[i] = diff
		}
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func SaturatedSubU32(a, b V128) V128 {
	var r u32x4
	x := *(*u32x4)(unsafe.Pointer(&a))
	y := *(*u32x4)(unsafe.Pointer(&b))
	for i := range r {
		if diff := x[i] - y[i]; diff > x[i] {
			r[i] = 0
		} else {
			r[i] = diff
		}
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func SaturatedSubU64(a, b V128) V128 {
	var r u64x2
	x := *(*u64x2)(unsafe.Pointer(&a))
	y := *(*u64x2)(unsafe.Pointer(&b))
	for i := range r {
		if diff := x[i] - y[i]; diff > x[i] {
			r[i] = 0
		} else {
			r[i] = diff
		}
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func ShlI8(a V128, y uint) V128 {
	var r i8x16
	x := *(*i8x16)(unsafe.Pointer(&a))
	for i := range r {
		r[i] <<= y
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func ShlI16(a V128, y uint) V128 {
	var r i16x8
	x := *(*i16x8)(unsafe.Pointer(&a))
	for i := range r {
		r[i] <<= y
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func ShlI32(a V128, y uint) V128 {
	var r i32x4
	x := *(*i32x4)(unsafe.Pointer(&a))
	for i := range r {
		r[i] <<= y
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func ShlI64(a V128, y uint) V128 {
	var r i64x2
	x := *(*i64x2)(unsafe.Pointer(&a))
	for i := range r {
		r[i] <<= y
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func ShrI8(a V128, y uint) V128 {
	var r i8x16
	x := *(*i8x16)(unsafe.Pointer(&a))
	for i := range r {
		r[i] >>= y
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func ShrI16(a V128, y uint) V128 {
	var r i16x8
	x := *(*i16x8)(unsafe.Pointer(&a))
	for i := range r {
		r[i] >>= y
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func ShrI32(a V128, y uint) V128 {
	var r i32x4
	x := *(*i32x4)(unsafe.Pointer(&a))
	for i := range r {
		r[i] >>= y
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func ShrI64(a V128, y uint) V128 {
	var r i64x2
	x := *(*i64x2)(unsafe.Pointer(&a))
	for i := range r {
		r[i] >>= y
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func ShrU8(a V128, y uint) V128 {
	var r u8x16
	x := *(*u8x16)(unsafe.Pointer(&a))
	for i := range r {
		r[i] >>= y
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func ShrU16(a V128, y uint) V128 {
	var r u16x8
	x := *(*u16x8)(unsafe.Pointer(&a))
	for i := range r {
		r[i] >>= y
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func ShrU32(a V128, y uint) V128 {
	var r u32x4
	x := *(*u32x4)(unsafe.Pointer(&a))
	for i := range r {
		r[i] >>= y
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func ShrU64(a V128, y uint) V128 {
	var r u64x2
	x := *(*u64x2)(unsafe.Pointer(&a))
	for i := range r {
		r[i] >>= y
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func And(a, b V128) (v V128) {
	for i := range v.b {
		v.b[i] = a.b[i] & b.b[i]
	}
	return v
}

func Or(a, b V128) (v V128) {
	for i := range v.b {
		v.b[i] = a.b[i] | b.b[i]
	}
	return v
}

func Xor(a, b V128) (v V128) {
	for i := range v.b {
		v.b[i] = a.b[i] ^ b.b[i]
	}
	return v
}

func Not(a V128) V128 {
	for i := range a.b {
		a.b[i] = ^a.b[i]
	}
	return a
}

func Select(v1, v2, c V128) V128 {
	return Or(And(v1, c), And(v2, Not(c)))
}

func AnyTrue(a V128) int {
	for i := range a.b {
		if a.b[i] != 0 {
			return 1
		}
	}
	return 0
}

func AllTrue(a V128) (r int) {
	for i := range a.b {
		if a.b[i] == 0 {
			return 0
		}
	}
	return 1
}

func EqI8(a, b V128) V128 {
	var r u8x16
	x := *(*u8x16)(unsafe.Pointer(&a))
	y := *(*u8x16)(unsafe.Pointer(&b))
	for i := range r {
		if x[i] == y[i] {
			r[i] = 1<<8 - 1
		}
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func EqI16(a, b V128) V128 {
	var r u16x8
	x := *(*u16x8)(unsafe.Pointer(&a))
	y := *(*u16x8)(unsafe.Pointer(&b))
	for i := range r {
		if x[i] == y[i] {
			r[i] = 1<<16 - 1
		}
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func EqI32(a, b V128) V128 {
	var r u32x4
	x := *(*u32x4)(unsafe.Pointer(&a))
	y := *(*u32x4)(unsafe.Pointer(&b))
	for i := range r {
		if x[i] == y[i] {
			r[i] = 1<<32 - 1
		}
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func EqI64(a, b V128) V128 {
	var r u64x2
	x := *(*u64x2)(unsafe.Pointer(&a))
	y := *(*u64x2)(unsafe.Pointer(&b))
	for i := range r {
		if x[i] == y[i] {
			r[i] = 1<<64 - 1
		}
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func EqF32(a, b V128) V128 {
	var r u32x4
	x := *(*f32x4)(unsafe.Pointer(&a))
	y := *(*f32x4)(unsafe.Pointer(&b))
	for i := range r {
		if x[i] == y[i] {
			r[i] = 1<<32 - 1
		}
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func EqF64(a, b V128) V128 {
	var r u64x2
	x := *(*f64x2)(unsafe.Pointer(&a))
	y := *(*f64x2)(unsafe.Pointer(&b))
	for i := range r {
		if x[i] == y[i] {
			r[i] = 1<<64 - 1
		}
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func NeI8(a, b V128) V128 {
	var r u8x16
	x := *(*u8x16)(unsafe.Pointer(&a))
	y := *(*u8x16)(unsafe.Pointer(&b))
	for i := range r {
		if x[i] != y[i] {
			r[i] = 1<<8 - 1
		}
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func NeI16(a, b V128) V128 {
	var r u16x8
	x := *(*u16x8)(unsafe.Pointer(&a))
	y := *(*u16x8)(unsafe.Pointer(&b))
	for i := range r {
		if x[i] != y[i] {
			r[i] = 1<<16 - 1
		}
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func NeI32(a, b V128) V128 {
	var r u32x4
	x := *(*u32x4)(unsafe.Pointer(&a))
	y := *(*u32x4)(unsafe.Pointer(&b))
	for i := range r {
		if x[i] != y[i] {
			r[i] = 1<<32 - 1
		}
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func NeI64(a, b V128) V128 {
	var r u64x2
	x := *(*u64x2)(unsafe.Pointer(&a))
	y := *(*u64x2)(unsafe.Pointer(&b))
	for i := range r {
		if x[i] != y[i] {
			r[i] = 1<<64 - 1
		}
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func NeF32(a, b V128) V128 {
	var r u32x4
	x := *(*f32x4)(unsafe.Pointer(&a))
	y := *(*f32x4)(unsafe.Pointer(&b))
	for i := range r {
		if x[i] != y[i] {
			r[i] = 1<<32 - 1
		}
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func NeF64(a, b V128) V128 {
	var r u64x2
	x := *(*f64x2)(unsafe.Pointer(&a))
	y := *(*f64x2)(unsafe.Pointer(&b))
	for i := range r {
		if x[i] != y[i] {
			r[i] = 1<<64 - 1
		}
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func LtI8(a, b V128) V128 {
	var r u8x16
	x := *(*i8x16)(unsafe.Pointer(&a))
	y := *(*i8x16)(unsafe.Pointer(&b))
	for i := range r {
		if x[i] < y[i] {
			r[i] = 1<<8 - 1
		}
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func LtI16(a, b V128) V128 {
	var r u16x8
	x := *(*i16x8)(unsafe.Pointer(&a))
	y := *(*i16x8)(unsafe.Pointer(&b))
	for i := range r {
		if x[i] < y[i] {
			r[i] = 1<<16 - 1
		}
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func LtI32(a, b V128) V128 {
	var r u32x4
	x := *(*i32x4)(unsafe.Pointer(&a))
	y := *(*i32x4)(unsafe.Pointer(&b))
	for i := range r {
		if x[i] < y[i] {
			r[i] = 1<<32 - 1
		}
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func LtI64(a, b V128) V128 {
	var r u64x2
	x := *(*i64x2)(unsafe.Pointer(&a))
	y := *(*i64x2)(unsafe.Pointer(&b))
	for i := range r {
		if x[i] < y[i] {
			r[i] = 1<<64 - 1
		}
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func LtU8(a, b V128) V128 {
	var r u8x16
	x := *(*u8x16)(unsafe.Pointer(&a))
	y := *(*u8x16)(unsafe.Pointer(&b))
	for i := range r {
		if x[i] < y[i] {
			r[i] = 1<<8 - 1
		}
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func LtU16(a, b V128) V128 {
	var r u16x8
	x := *(*u16x8)(unsafe.Pointer(&a))
	y := *(*u16x8)(unsafe.Pointer(&b))
	for i := range r {
		if x[i] < y[i] {
			r[i] = 1<<16 - 1
		}
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func LtU32(a, b V128) V128 {
	var r u32x4
	x := *(*u32x4)(unsafe.Pointer(&a))
	y := *(*u32x4)(unsafe.Pointer(&b))
	for i := range r {
		if x[i] < y[i] {
			r[i] = 1<<32 - 1
		}
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func LtU64(a, b V128) V128 {
	var r u64x2
	x := *(*u64x2)(unsafe.Pointer(&a))
	y := *(*u64x2)(unsafe.Pointer(&b))
	for i := range r {
		if x[i] < y[i] {
			r[i] = 1<<64 - 1
		}
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func LtF32(a, b V128) V128 {
	var r u32x4
	x := *(*f32x4)(unsafe.Pointer(&a))
	y := *(*f32x4)(unsafe.Pointer(&b))
	for i := range r {
		if x[i] < y[i] {
			r[i] = 1<<32 - 1
		}
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func LtF64(a, b V128) V128 {
	var r u64x2
	x := *(*f64x2)(unsafe.Pointer(&a))
	y := *(*f64x2)(unsafe.Pointer(&b))
	for i := range r {
		if x[i] < y[i] {
			r[i] = 1<<64 - 1
		}
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func LeI8(a, b V128) V128 {
	var r u8x16
	x := *(*i8x16)(unsafe.Pointer(&a))
	y := *(*i8x16)(unsafe.Pointer(&b))
	for i := range r {
		if x[i] <= y[i] {
			r[i] = 1<<8 - 1
		}
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func LeI16(a, b V128) V128 {
	var r u16x8
	x := *(*i16x8)(unsafe.Pointer(&a))
	y := *(*i16x8)(unsafe.Pointer(&b))
	for i := range r {
		if x[i] <= y[i] {
			r[i] = 1<<16 - 1
		}
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func LeI32(a, b V128) V128 {
	var r u32x4
	x := *(*i32x4)(unsafe.Pointer(&a))
	y := *(*i32x4)(unsafe.Pointer(&b))
	for i := range r {
		if x[i] <= y[i] {
			r[i] = 1<<32 - 1
		}
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func LeI64(a, b V128) V128 {
	var r u64x2
	x := *(*i64x2)(unsafe.Pointer(&a))
	y := *(*i64x2)(unsafe.Pointer(&b))
	for i := range r {
		if x[i] <= y[i] {
			r[i] = 1<<64 - 1
		}
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func LeU8(a, b V128) V128 {
	var r u8x16
	x := *(*u8x16)(unsafe.Pointer(&a))
	y := *(*u8x16)(unsafe.Pointer(&b))
	for i := range r {
		if x[i] <= y[i] {
			r[i] = 1<<8 - 1
		}
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func LeU16(a, b V128) V128 {
	var r u16x8
	x := *(*u16x8)(unsafe.Pointer(&a))
	y := *(*u16x8)(unsafe.Pointer(&b))
	for i := range r {
		if x[i] <= y[i] {
			r[i] = 1<<16 - 1
		}
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func LeU32(a, b V128) V128 {
	var r u32x4
	x := *(*u32x4)(unsafe.Pointer(&a))
	y := *(*u32x4)(unsafe.Pointer(&b))
	for i := range r {
		if x[i] <= y[i] {
			r[i] = 1<<32 - 1
		}
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func LeU64(a, b V128) V128 {
	var r u64x2
	x := *(*u64x2)(unsafe.Pointer(&a))
	y := *(*u64x2)(unsafe.Pointer(&b))
	for i := range r {
		if x[i] <= y[i] {
			r[i] = 1<<64 - 1
		}
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func LeF32(a, b V128) V128 {
	var r u32x4
	x := *(*f32x4)(unsafe.Pointer(&a))
	y := *(*f32x4)(unsafe.Pointer(&b))
	for i := range r {
		if x[i] <= y[i] {
			r[i] = 1<<32 - 1
		}
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func LeF64(a, b V128) V128 {
	var r u64x2
	x := *(*f64x2)(unsafe.Pointer(&a))
	y := *(*f64x2)(unsafe.Pointer(&b))
	for i := range r {
		if x[i] <= y[i] {
			r[i] = 1<<64 - 1
		}
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func GtI8(a, b V128) V128 {
	var r u8x16
	x := *(*i8x16)(unsafe.Pointer(&a))
	y := *(*i8x16)(unsafe.Pointer(&b))
	for i := range r {
		if x[i] > y[i] {
			r[i] = 1<<8 - 1
		}
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func GtI16(a, b V128) V128 {
	var r u16x8
	x := *(*i16x8)(unsafe.Pointer(&a))
	y := *(*i16x8)(unsafe.Pointer(&b))
	for i := range r {
		if x[i] > y[i] {
			r[i] = 1<<16 - 1
		}
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func GtI32(a, b V128) V128 {
	var r u32x4
	x := *(*i32x4)(unsafe.Pointer(&a))
	y := *(*i32x4)(unsafe.Pointer(&b))
	for i := range r {
		if x[i] > y[i] {
			r[i] = 1<<32 - 1
		}
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func GtI64(a, b V128) V128 {
	var r u64x2
	x := *(*i64x2)(unsafe.Pointer(&a))
	y := *(*i64x2)(unsafe.Pointer(&b))
	for i := range r {
		if x[i] > y[i] {
			r[i] = 1<<64 - 1
		}
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func GtU8(a, b V128) V128 {
	var r u8x16
	x := *(*u8x16)(unsafe.Pointer(&a))
	y := *(*u8x16)(unsafe.Pointer(&b))
	for i := range r {
		if x[i] > y[i] {
			r[i] = 1<<8 - 1
		}
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func GtU16(a, b V128) V128 {
	var r u16x8
	x := *(*u16x8)(unsafe.Pointer(&a))
	y := *(*u16x8)(unsafe.Pointer(&b))
	for i := range r {
		if x[i] > y[i] {
			r[i] = 1<<16 - 1
		}
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func GtU32(a, b V128) V128 {
	var r u32x4
	x := *(*u32x4)(unsafe.Pointer(&a))
	y := *(*u32x4)(unsafe.Pointer(&b))
	for i := range r {
		if x[i] > y[i] {
			r[i] = 1<<32 - 1
		}
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func GtU64(a, b V128) V128 {
	var r u64x2
	x := *(*u64x2)(unsafe.Pointer(&a))
	y := *(*u64x2)(unsafe.Pointer(&b))
	for i := range r {
		if x[i] > y[i] {
			r[i] = 1<<64 - 1
		}
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func GtF32(a, b V128) V128 {
	var r u32x4
	x := *(*f32x4)(unsafe.Pointer(&a))
	y := *(*f32x4)(unsafe.Pointer(&b))
	for i := range r {
		if x[i] > y[i] {
			r[i] = 1<<32 - 1
		}
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func GtF64(a, b V128) V128 {
	var r u64x2
	x := *(*f64x2)(unsafe.Pointer(&a))
	y := *(*f64x2)(unsafe.Pointer(&b))
	for i := range r {
		if x[i] > y[i] {
			r[i] = 1<<64 - 1
		}
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func GeI8(a, b V128) V128 {
	var r u8x16
	x := *(*i8x16)(unsafe.Pointer(&a))
	y := *(*i8x16)(unsafe.Pointer(&b))
	for i := range r {
		if x[i] >= y[i] {
			r[i] = 1<<8 - 1
		}
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func GeI16(a, b V128) V128 {
	var r u16x8
	x := *(*i16x8)(unsafe.Pointer(&a))
	y := *(*i16x8)(unsafe.Pointer(&b))
	for i := range r {
		if x[i] >= y[i] {
			r[i] = 1<<16 - 1
		}
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func GeI32(a, b V128) V128 {
	var r u32x4
	x := *(*i32x4)(unsafe.Pointer(&a))
	y := *(*i32x4)(unsafe.Pointer(&b))
	for i := range r {
		if x[i] >= y[i] {
			r[i] = 1<<32 - 1
		}
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func GeI64(a, b V128) V128 {
	var r u64x2
	x := *(*i64x2)(unsafe.Pointer(&a))
	y := *(*i64x2)(unsafe.Pointer(&b))
	for i := range r {
		if x[i] >= y[i] {
			r[i] = 1<<64 - 1
		}
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func GeU8(a, b V128) V128 {
	var r u8x16
	x := *(*u8x16)(unsafe.Pointer(&a))
	y := *(*u8x16)(unsafe.Pointer(&b))
	for i := range r {
		if x[i] >= y[i] {
			r[i] = 1<<8 - 1
		}
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func GeU16(a, b V128) V128 {
	var r u16x8
	x := *(*u16x8)(unsafe.Pointer(&a))
	y := *(*u16x8)(unsafe.Pointer(&b))
	for i := range r {
		if x[i] >= y[i] {
			r[i] = 1<<16 - 1
		}
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func GeU32(a, b V128) V128 {
	var r u32x4
	x := *(*u32x4)(unsafe.Pointer(&a))
	y := *(*u32x4)(unsafe.Pointer(&b))
	for i := range r {
		if x[i] >= y[i] {
			r[i] = 1<<32 - 1
		}
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func GeU64(a, b V128) V128 {
	var r u64x2
	x := *(*u64x2)(unsafe.Pointer(&a))
	y := *(*u64x2)(unsafe.Pointer(&b))
	for i := range r {
		if x[i] >= y[i] {
			r[i] = 1<<64 - 1
		}
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func GeF32(a, b V128) V128 {
	var r u32x4
	x := *(*f32x4)(unsafe.Pointer(&a))
	y := *(*f32x4)(unsafe.Pointer(&b))
	for i := range r {
		if x[i] >= y[i] {
			r[i] = 1<<32 - 1
		}
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func GeF64(a, b V128) V128 {
	var r u64x2
	x := *(*f64x2)(unsafe.Pointer(&a))
	y := *(*f64x2)(unsafe.Pointer(&b))
	for i := range r {
		if x[i] >= y[i] {
			r[i] = 1<<64 - 1
		}
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func NegF32(a V128) V128 {
	x := *(*f32x4)(unsafe.Pointer(&a))
	for i := range x {
		x[i] = math.Float32frombits(math.Float32bits(x[i]) ^ 1<<31)
	}
	return x
}

func NegF64(a V128) V128 {
	x := *(*f64x2)(unsafe.Pointer(&a))
	for i := range x {
		x[i] = math.Float64frombits(math.Float64bits(x[i]) ^ 1<<63)
	}
	return x
}

func absf32(f float32) float32 {
	return math.Float32frombits(math.Float32bits(x) &^ (1 << 31))
}

func AbsF32(a V128) V128 {
	x := *(*f32x4)(unsafe.Pointer(&a))
	for i := range x {
		x[i] = absf32(x[i])
	}
	return x
}

func AbsF64(a V128) V128 {
	x := *(*f64x2)(unsafe.Pointer(&a))
	for i := range x {
		x[i] = math.Abs(x[i])
	}
	return x
}

func MinF32(a, b V128) V128 {
	var r f32x4
	x := *(*f32x4)(unsafe.Pointer(&a))
	y := *(*f32x4)(unsafe.Pointer(&b))
	for i := range r {
		r[i] = float32(math.Min(float64(x[i]), float64(y[i])))
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func MinF64(a, b V128) V128 {
	var r f64x2
	x := *(*f64x2)(unsafe.Pointer(&a))
	y := *(*f64x2)(unsafe.Pointer(&b))
	for i := range r {
		r[i] = math.Min(x[i], y[i])
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func MaxF32(a, b V128) V128 {
	var r f32x4
	x := *(*f32x4)(unsafe.Pointer(&a))
	y := *(*f32x4)(unsafe.Pointer(&b))
	for i := range r {
		r[i] = float32(math.Max(float64(x[i]), float64(y[i])))
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func MaxF64(a, b V128) V128 {
	var r f64x2
	x := *(*f64x2)(unsafe.Pointer(&a))
	y := *(*f64x2)(unsafe.Pointer(&b))
	for i := range r {
		r[i] = math.Max(x[i], y[i])
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func SqrtF32(a V128) V128 {
	var r f32x4
	x := *(*f32x4)(unsafe.Pointer(&a))
	for i := range r {
		r[i] = float32(math.Sqrt(float64(x)))
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func SqrtF64(a V128) V128 {
	var r f64x2
	x := *(*f64x2)(unsafe.Pointer(&a))
	for i := range r {
		r[i] = math.Sqrt(x[i])
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func ConvertI32toF32(a V128) V128 {
	var r f32x4
	x := *(*i32x4)(unsafe.Pointer(&a))
	for i := range r {
		r[i] = float32(x[i])
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func ConvertU32toF32(a V128) V128 {
	var r f32x4
	x := *(*u32x4)(unsafe.Pointer(&a))
	for i := range r {
		r[i] = float32(x[i])
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func ConvertI64toF64(a V128) V128 {
	var r f64x2
	x := *(*i64x2)(unsafe.Pointer(&a))
	for i := range r {
		r[i] = float64(x[i])
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func ConvertU64toF64(a V128) V128 {
	var r f64x2
	x := *(*u64x2)(unsafe.Pointer(&a))
	for i := range r {
		r[i] = float64(x[i])
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func isNaN32(f float32) bool {
	return f != f
}

func TruncateF32toI32(a V128) V128 {
	var r i32x4
	x := *(*f32x4)(unsafe.Pointer(&a))
	for i := range r {
		if !isNaN32(x[i]) {
			r[i] = int32(x[i])
		}
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func TruncateF32toU32(a V128) V128 {
	var r u32x4
	x := *(*f32x4)(unsafe.Pointer(&a))
	for i := range r {
		if !isNaN32(x[i]) {
			r[i] = uint32(x[i])
		}
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func TruncateF64toI64(a V128) V128 {
	var r i64x2
	x := *(*f64x2)(unsafe.Pointer(&a))
	for i := range r {
		if !math.IsNaN(x[i]) {
			r[i] = int64(x[i])
		}
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func TruncateF64toU64(a V128) V128 {
	var r u64x2
	x := *(*f64x2)(unsafe.Pointer(&a))
	for i := range r {
		if !math.IsNaN(x[i]) {
			r[i] = uint64(x[i])
		}
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func NarrowI16toI8(a, b V128) V128 {
	var r i8x16
	x := *(*i16x8)(unsafe.Pointer(&a))
	lx := len(x)
	y := *(*i16x8)(unsafe.Pointer(&b))
	for i := range x {
		r[i] = int8(x[i])
	}
	for i := range y {
		r[i+lx] = int8(y[i])
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func NarrowI32toI16(a, b V128) V128 {
	var r i16x8
	x := *(*i32x4)(unsafe.Pointer(&a))
	lx := len(x)
	y := *(*i32x4)(unsafe.Pointer(&b))
	for i := range x {
		r[i] = int16(x[i])
	}
	for i := range y {
		r[i+lx] = int16(y[i])
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func NarrowI64toI32(a, b V128) V128 {
	var r i32x4
	x := *(*i64x2)(unsafe.Pointer(&a))
	lx := len(x)
	y := *(*i64x2)(unsafe.Pointer(&b))
	for i := range x {
		r[i] = int32(x[i])
	}
	for i := range y {
		r[i+lx] = int32(y[i])
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func NarrowU16toU8(a, b V128) V128 {
	var r u8x16
	x := *(*u16x8)(unsafe.Pointer(&a))
	lx := len(x)
	y := *(*u16x8)(unsafe.Pointer(&b))
	for i := range x {
		r[i] = uint8(x[i])
	}
	for i := range y {
		r[i+lx] = uint8(y[i])
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func NarrowU32toU16(a, b V128) V128 {
	var r u16x8
	x := *(*u32x4)(unsafe.Pointer(&a))
	lx := len(x)
	y := *(*u32x4)(unsafe.Pointer(&b))
	for i := range x {
		r[i] = uint16(x[i])
	}
	for i := range y {
		r[i+lx] = uint16(y[i])
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func NarrowU64toU32(a, b V128) V128 {
	var r u32x4
	x := *(*u64x2)(unsafe.Pointer(&a))
	lx := len(x)
	y := *(*u64x2)(unsafe.Pointer(&b))
	for i := range x {
		r[i] = uint32(x[i])
	}
	for i := range y {
		r[i+lx] = uint32(y[i])
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func WidenLowI8toI16(a V128) V128 {
	var r i16x8
	x := *(*i8x16)(unsafe.Pointer(&a))
	for i := range r {
		r[i] = int16(x[i*2+1])
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func WidenHighI8toI16(a V128) V128 {
	var r i16x8
	x := *(*i8x16)(unsafe.Pointer(&a))
	for i := range r {
		r[i] = int16(x[i*2])
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func WidenLowI16toI32(a V128) V128 {
	var r i32x4
	x := *(*i16x8)(unsafe.Pointer(&a))
	for i := range r {
		r[i] = int32(x[i*2+1])
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func WidenHighI16toI32(a V128) V128 {
	var r i32x4
	x := *(*i16x8)(unsafe.Pointer(&a))
	for i := range r {
		r[i] = int32(x[i*2])
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func WidenLowI32toI64(a V128) V128 {
	var r i64x2
	x := *(*i32x4)(unsafe.Pointer(&a))
	for i := range r {
		r[i] = int64(x[i*2+1])
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func WidenHighI32toI64(a V128) V128 {
	var r i64x2
	x := *(*i32x4)(unsafe.Pointer(&a))
	for i := range r {
		r[i] = int64(x[i*2])
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func WidenLowU8toU16(a V128) V128 {
	var r u16x8
	x := *(*u8x16)(unsafe.Pointer(&a))
	for i := range r {
		r[i] = uint16(x[i*2+1])
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func WidenHighU8toU16(a V128) V128 {
	var r u16x8
	x := *(*u8x16)(unsafe.Pointer(&a))
	for i := range r {
		r[i] = uint16(x[i*2])
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func WidenLowU16toU32(a V128) V128 {
	var r u32x4
	x := *(*u16x8)(unsafe.Pointer(&a))
	for i := range r {
		r[i] = uint32(x[i*2+1])
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func WidenHighU16toU32(a V128) V128 {
	var r u32x4
	x := *(*u16x8)(unsafe.Pointer(&a))
	for i := range r {
		r[i] = uint32(x[i*2])
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func WidenLowU32toU64(a V128) V128 {
	var r u64x2
	x := *(*u32x4)(unsafe.Pointer(&a))
	for i := range r {
		r[i] = uint64(x[i*2+1])
	}
	return *(*V128)(unsafe.Pointer(&r))
}

func WidenHighU32toU64(a V128) V128 {
	var r u64x2
	x := *(*u32x4)(unsafe.Pointer(&a))
	for i := range r {
		r[i] = uint64(x[i*2])
	}
	return *(*V128)(unsafe.Pointer(&r))
}
