package main

type TypePrimaryCollection struct {
	Bool  bool
	Int   int
	Str   string
	Float float64

	PtrBool  *bool
	PtrInt   *int
	PtrStr   *string
	PtrFloat *float64

	//SliceBool  []bool
	//SliceInt   []int
	//SliceStr   []string
	//SliceFloat []float64
	//
	//SlicePtrBool  []*bool
	//SlicePtrInt   []*int
	//SlicePtrStr   []*string
	//SlicePtrFloat []*float64
	//
	//PtrSliceBool  *[]bool
	//PtrSliceInt   *[]int
	//PtrSliceStr   *[]string
	//PtrSliceFloat *[]float64
	//
	//PtrSlicePtrBool  *[]*bool
	//PtrSlicePtrInt   *[]*int
	//PtrSlicePtrStr   *[]*string
	//PtrSlicePtrFloat *[]*float64
}

type Bool bool
type Int int
type Str string
type Float float64

type TypePrimaryAliasCollection struct {
	Bool  Bool
	Int   Int
	Str   Str
	Float Float

	PtrBool  *Bool
	PtrInt   *Int
	PtrStr   *Str
	PtrFloat *Float

	//SliceBool  []Bool
	//SliceInt   []Int
	//SliceStr   []Str
	//SliceFloat []Float
	//
	//SlicePtrBool  []*Bool
	//SlicePtrInt   []*Int
	//SlicePtrStr   []*Str
	//SlicePtrFloat []*Float
	//
	//PtrSliceBool  *[]Bool
	//PtrSliceInt   *[]Int
	//PtrSliceStr   *[]Str
	//PtrSliceFloat *[]Float
	//
	//PtrSlicePtrBool  *[]*Bool
	//PtrSlicePtrInt   *[]*Int
	//PtrSlicePtrStr   *[]*Str
	//PtrSlicePtrFloat *[]*Float
}

type PtrBool *bool
type PtrInt *int
type PtrStr *string
type PtrFloat *float64

type TypePrimaryPointerAliasCollection struct {
	PtrBool  PtrBool
	//PtrInt   PtrInt
	//PtrStr   PtrStr
	//PtrFloat PtrFloat

	//SlicePtrBool  []PtrBool
	//SlicePtrInt   []PtrInt
	//SlicePtrStr   []PtrStr
	//SlicePtrFloat []PtrFloat
	//
	//PtrSlicePtrBool  *[]PtrBool
	//PtrSlicePtrInt   *[]PtrInt
	//PtrSlicePtrStr   *[]PtrStr
	//PtrSlicePtrFloat *[]PtrFloat
}

func main() {}
