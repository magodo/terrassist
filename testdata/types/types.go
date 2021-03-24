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
	BoolSlicePtr  *[]bool
	IntSlicePtr   *[]int
	StrSlicePtr   *[]string
	FloatSlicePtr *[]float64
}

type TypePrimaryPtrSlicePtrCollection struct {
	PtrBool  *[]*bool
	PtrInt   *[]*int
	PtrStr   *[]*string
	PtrFloat *[]*float64
}

type TypePrimaryMapCollection struct {
	BoolMap  map[string]bool
	IntMap   map[string]int
	StrMap   map[string]string
	FloatMap map[string]float64
}

type TypePrimaryPtrMapCollection struct {
	BoolPtrMap  map[string]*bool
	IntPtrMap   map[string]*int
	StrPtrMap   map[string]*string
	FloatPtrMap map[string]*float64
}

type TypePrimaryMapPtrCollection struct {
	BoolMapPtr  *map[string]bool
	IntMapPtr   *map[string]int
	StrMapPtr   *map[string]string
	FloatMapPtr *map[string]float64
}

type TypePrimaryPtrMapPtrCollection struct {
	BoolPtrMapPtr  *map[string]*bool
	IntPtrMapPtr   *map[string]*int
	StrPtrMapPtr   *map[string]*string
	FloatPtrMapPtr *map[string]*float64
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

type TypePrimaryAliasMapCollection struct {
	BoolMap  map[string]Bool
	IntMap   map[string]Int
	StrMap   map[string]Str
	FloatMap map[string]Float
}

type TypePrimaryAliasPtrMapCollection struct {
	BoolPtrMap   map[string]*Bool
	IntPtrMap    map[string]*Int
	StringPtrMap map[string]*Str
	FloatPtrMap  map[string]*Float
}

type TypePrimaryAliasMapPtrCollection struct {
	BoolMapPtr   *map[string]Bool
	IntMapPtr    *map[string]Int
	StringMapPtr *map[string]Str
	FloatMapPtr  *map[string]Float
}

type TypePrimaryAliasPtrMapPtrCollection struct {
	BoolPtrMapPtr   *map[string]*Bool
	IntPtrMapPtr    *map[string]*Int
	StringPtrMapPtr *map[string]*Str
	FloatPtrMapPtr  *map[string]*Float
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

type TypePrimaryPtrAliasMapCollection struct {
	BoolPtrMap  map[string]BoolPtr
	IntPtrMap   map[string]IntPtr
	StrPtrMap   map[string]StrPtr
	FloatPtrMap map[string]FloatPtr
}

type TypePrimaryPtrAliasMapPtrCollection struct {
	BoolPtrMapPtr *map[string]BoolPtr
	IntPtrMapPtr  *map[string]IntPtr
	StrPtrMapPtr  *map[string]StrPtr
	FlatPtrMapPtr *map[string]FloatPtr
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

type TypeNamedStructMapCollection struct {
	S1Map       map[string]TypeS1
	S1MapPtr    *map[string]TypeS1
	S1PtrMap    map[string]*TypeS1
	S1PtrMapPtr *map[string]*TypeS1
}

type TypeIntSliceAlias []int

type TypeIntPtrSliceAlias []*int

type TypeNamedBoolSliceAlias []Bool

type TypeNamedStructSliceAlias []TypeS1

type TypeIntMapAlias map[string]int

type TypeIntPtrMapAlias map[string]*int

type TypeNamedBoolMapAlias map[string]Bool

type TypeNamedStructMapAlias map[string]TypeS1

type TypeNamedStructAlias TypeS1

type TypeNamedStructPtrAlias *TypeS1

type TypeNamedStructWithJSONIgnore struct {
	I       int    `json:"i"`
	J       string `json:"j"`
	Ignored int    `json:"-"`
}

type TypeCyclicRefStruct struct {
	Self *TypeCyclicRefStruct
}

type TypeNamedStructWithGoReservedWord struct {
	Go *string
	If *string
}
