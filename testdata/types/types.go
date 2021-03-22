package types

type TypePrimaryCollection struct {
	Bool  bool
	Int   int
	Str   string
	Float float64
}

type TypePrimaryPtrCollection struct {
	PtrBool  *bool
	PtrInt   *int
	PtrStr   *string
	PtrFloat *float64
}

type TypePrimarySliceCollection struct {
	BoolSlice  []bool
	IntSlice   []int
	StrSlice   []string
	FloatSlice []float64
}

type TypePrimaryPtrSliceCollection struct {
	BoolPtrSlice  []*bool
	IntPtrSlice   []*int
	StrPtrSlice   []*string
	FloatPtrSlice []*float64
}

type TypePrimarySlicePtrCollection struct {
	BoolPtrSlicePtr  *[]*bool
	IntPtrSlicePtr   *[]*int
	StrPtrSlicePtr   *[]*string
	FloatPtrSlicePtr *[]*float64
}

type TypePrimaryPtrSlicePtrCollection struct {
	PtrBool  *[]*bool
	PtrInt   *[]*int
	PtrStr   *[]*string
	PtrFloat *[]*float64
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
}

type TypePrimaryAliasPtrCollection struct {
	PtrBool  *Bool
	PtrInt   *Int
	PtrStr   *Str
	PtrFloat *Float
}

type TypePrimaryAliasSliceCollection struct {
	BoolSlice  []Bool
	IntSlice   []Int
	StrSlice   []Str
	FloatSlice []Float
}

type TypePrimaryAliasPtrSliceCollection struct {
	BoolPtrSlice   []*Bool
	IntPtrSlice    []*Int
	StringPtrSlice []*Str
	FloatPtrSlice  []*Float
}

type TypePrimaryAliasSlicePtrCollection struct {
	BoolSlicePtr   *[]Bool
	IntSlicePtr    *[]Int
	StringSlicePtr *[]Str
	FloatSlicePtr  *[]Float
}

type TypePrimaryAliasPtrSlicePtrCollection struct {
	BoolPtrSlicePtr   *[]*Bool
	IntPtrSlicePtr    *[]*Int
	StringPtrSlicePtr *[]*Str
	FloatPtrSlicePtr  *[]*Float
}

type BoolPtr *bool
type IntPtr *int
type StrPtr *string
type FloatPtr *float64

type TypePrimaryPtrAliasCollection struct {
	BoolPtr  BoolPtr
	IntPtr   IntPtr
	StrPtr   StrPtr
	FloatPtr FloatPtr
}

type TypePrimaryPtrAliasSliceCollection struct {
	BoolPtrSlice  []BoolPtr
	IntPtrSlice   []IntPtr
	StrPtrSlice   []StrPtr
	FloatPtrSlice []FloatPtr
}

type TypePrimaryPtrAliasSlicePtrCollection struct {
	BoolPtrSlicePtr *[]BoolPtr
	IntPtrSlicePtr  *[]IntPtr
	StrPtrSlicePtr  *[]StrPtr
	FlatPtrSlicePtr *[]FloatPtr
}

type TypeS1 struct {
	I int
}

type TypeS2 struct {
	S string
}

type TypeNamedStructCollection struct {
	S1 TypeS1
	S2 TypeS2
}

type TypeNamedStructPointerCollection struct {
	S1 *TypeS1
	S2 *TypeS2
}

type TypeNamedStructSliceCollection struct {
	S1Slice       []TypeS1
	S1SlicePtr    *[]TypeS1
	S1PtrSlice    []*TypeS1
	S1PtrSlicePtr *[]*TypeS1
}

type TypeIntSliceAlias []int

type TypeIntPtrSliceAlias []*int

type TypeNamedBoolSliceAlias []Bool

type TypeNamedStructSliceAlias []TypeS1
