package qtmoc

import (
	"errors"
	"os"
	"strings"
	"testing"
	"unsafe"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/sql"
	"github.com/therecipe/qt/widgets"

	"github.com/therecipe/qt/internal/cmd/moc/test/sub"
)

//go:generate qtmoc
type testStruct struct {
	otherTestStruct
	testOther otherTestStruct

	_ bool             `prop:"propBool"`
	_ int8             `prop:"propInt8"`  //-> string
	_ uint8            `prop:"propInt82"` //-> string
	_ int16            `prop:"propInt16"`
	_ uint16           `prop:"propInt162"`
	_ int32            `prop:"propInt32"`  // -> int
	_ uint32           `prop:"propInt322"` // -> int
	_ int              `prop:"propInt"`
	_ uint             `prop:"propInt2"`
	_ int64            `prop:"propInt64"`
	_ uint64           `prop:"propInt642"`
	_ float32          `prop:"propFloat"`
	_ float64          `prop:"propFloat2"`
	_ string           `prop:"propString"`
	_ []string         `prop:"propString2"`
	_ uintptr          `prop:"propPointer"`
	_ unsafe.Pointer   `prop:"propPointer2"`
	_ core.QVariant    `prop:"propObject"`  // -> T (c++)
	_ *core.QObject    `prop:"propObject2"` // -> *T
	_ *core.QVariant   `prop:"propObject3"` // -> *T //TODO:
	_ core.Qt__Key     `prop:"propEnum"`
	_ error            `prop:"propError"`
	_ otherTestStruct  `prop:"propReturnTest"`  // -> *T
	_ *otherTestStruct `prop:"propReturnTest2"` // -> *T

	a, b bool
	ab   func(bool) bool
	abc  func(bool) bool `test:"test"`
	abcd struct {
		a, b bool
		ab   func(bool) bool
		abc  func(bool) bool `test:"test"`
	}

	test widgets.QWidget
	core.QObject
	widgets.QWidget

	_ func(bool, bool)                   `signal:"BoolSignalInput"`
	_ func(int8, uint8)                  `signal:"Int8SignalInput"` // -> string
	_ func(int16, uint16)                `signal:"Int16SignalInput"`
	_ func(int32, uint32)                `signal:"Int32SignalInput"` // -> int
	_ func(int, uint)                    `signal:"IntSignalInput"`
	_ func(int64, uint64)                `signal:"Int64SignalInput"`
	_ func(float32, float64)             `signal:"FloatSignalInput"`
	_ func(string, []string)             `signal:"StringSignalInput"`
	_ func(uintptr, unsafe.Pointer)      `signal:"PointerSignalInput"`
	_ func(core.QVariant, *core.QObject) `signal:"ObjectSignalInput"` // -> T (c++) *T
	_ func(core.Qt__Key)                 `signal:"EnumSignalInput"`
	_ func(error)                        `signal:"ErrorSignalInput"`

	_ func(bool, bool)                   `slot:"BoolSlotInput"`
	_ func(int8, uint8)                  `slot:"Int8SlotInput"` // -> string
	_ func(int16, uint16)                `slot:"Int16SlotInput"`
	_ func(int32, uint32)                `slot:"Int32SlotInput"` // -> int
	_ func(int, uint)                    `slot:"IntSlotInput"`
	_ func(int64, uint64)                `slot:"Int64SlotInput"`
	_ func(float32, float64)             `slot:"FloatSlotInput"`
	_ func(string, []string)             `slot:"StringSlotInput"`
	_ func(uintptr, unsafe.Pointer)      `slot:"PointerSlotInput"`
	_ func(core.QVariant, *core.QObject) `slot:"ObjectSlotInput"` // -> T (c++) *T
	_ func(core.Qt__Key)                 `slot:"EnumSlotInput"`
	_ func(error)                        `slot:"ErrorSlotInput"`

	_ func(bool) bool                     `slot:"BoolSlotOutput"`
	_ func(bool) bool                     `slot:"BoolSlotOutput2"`
	_ func(int8) int8                     `slot:"Int8SlotOutput"`  //-> string
	_ func(uint8) uint8                   `slot:"Int8SlotOutput2"` //-> string
	_ func(int16) int16                   `slot:"Int16SlotOutput"`
	_ func(uint16) uint16                 `slot:"Int16SlotOutput2"`
	_ func(int32) int32                   `slot:"Int32SlotOutput"`  // -> int
	_ func(uint32) uint32                 `slot:"Int32SlotOutput2"` // -> int
	_ func(int) int                       `slot:"IntSlotOutput"`
	_ func(uint) uint                     `slot:"IntSlotOutput2"`
	_ func(int64) int64                   `slot:"Int64SlotOutput"`
	_ func(uint64) uint64                 `slot:"Int64SlotOutput2"`
	_ func(float32) float32               `slot:"FloatSlotOutput"`
	_ func(float64) float64               `slot:"FloatSlotOutput2"`
	_ func(string) string                 `slot:"StringSlotOutput"`
	_ func([]string) []string             `slot:"StringSlotOutput2"`
	_ func(uintptr) uintptr               `slot:"PointerSlotOutput"`
	_ func(unsafe.Pointer) unsafe.Pointer `slot:"PointerSlotOutput2"`
	_ func(core.QVariant) core.QVariant   `slot:"ObjectSlotOutput"`  // -> T (c++)
	_ func(*core.QObject) *core.QObject   `slot:"ObjectSlotOutput2"` // -> *T
	_ func(core.Qt__Key) core.Qt__Key     `slot:"EnumSlotOutput"`
	_ func(error) error                   `slot:"ErrorSlotOutput"`
	_ func(testStruct) testStruct         `slot:"returnTest"`  // -> *T
	_ func(*testStruct) *testStruct       `slot:"returnTest2"` // -> *T
	_ func(a0 string) (a1 string)         `slot:"returnName"`
	_ func(a0, a1 int) (a2 int)           `slot:"returnName2"`
	_ func()                              `slot:"other"`
	_ func() bool                         `slot:"other2"`
}

type subTestStruct struct {
	testStruct
	_ func(*subTestStruct) *subTestStruct `slot:"returnTest3"`
}

type subSubTestStruct struct { //Qt structs from other pkgs will be ignored for now
	sub.SubTestStruct
	_ func(*sub.SubTestStruct) *sub.SubTestStruct `slot:"returnTest4"`
}

type otherTestStruct struct {
	core.QObject

	a, b bool
	ab   func(bool) bool
	abc  func(bool) bool `test:"test"`
	abcd struct {
		a, b bool
		ab   func(bool) bool
		abc  func(bool) bool `test:"test"`
	}

	_ bool `prop:"propBoolSub"`
}

type abstractTestStruct1 struct {
	core.QAbstractItemModel
}

type abstractTestStruct2 struct {
	core.QAbstractListModel
}

type abstractTestStruct3 struct {
	core.QStringListModel
}

type abstractTestStruct4 struct {
	core.QAbstractProxyModel
}

type abstractTestStruct5 struct {
	core.QAbstractTableModel
}

type abstractTestStruct6 struct {
	sql.QSqlQueryModel
}

type abstractTestStruct7 struct {
	gui.QStandardItemModel
}

type abstractTestStruct8 struct {
	widgets.QFileSystemModel
}

var (
	b0, b1 bool = false, true

	i0 int16  = 1
	i1 uint16 = 2

	i2 int  = 3
	i3 uint = 4

	i4 int  = 5
	i5 uint = 6

	i6 int64  = 7
	i7 uint64 = 8

	f0 float32 = 123.45
	f1 float64 = 678.91

	s0 string   = "test"
	s1 []string = []string{"t", "e", "s", "t"}

	p0 uintptr        = uintptr(12345)
	p1 unsafe.Pointer = unsafe.Pointer(uintptr(67891))

	o0 *core.QVariant
	o1 *core.QObject

	e0 core.Qt__Key = core.Qt__Key_Z
	e1 error        = errors.New("test")
)

func init() { core.NewQCoreApplication(len(os.Args), os.Args) }

func TestGeneral(t *testing.T) {
	if false {
		NewTestStruct(nil)
		NewOtherTestStruct(nil)
		NewSubTestStruct(nil)
		sub.NewSubTestStruct(nil)
		//NewSubSubTestStruct(nil)
	}
}

func TestProperties(t *testing.T) {
	var test = NewTestStruct(nil)

	test.ConnectPropBoolChanged(func(propBool bool) {
		if propBool != b1 {
			t.Fatal(propBool, b1)
		}
	})
	test.SetPropBool(b1)
	test.PropBoolChanged(b1)
	if test.IsPropBool() != b1 {
		t.Fatal("IsPropBool")
	}

	//_ int8             `prop:"propInt8"`
	//_ uint8            `prop:"propInt82"`

	test.ConnectPropInt16Changed(func(propInt16 int16) {
		if propInt16 != i0 {
			t.Fatal(propInt16, i0)
		}
	})
	test.SetPropInt16(i0)
	test.PropInt16Changed(i0)
	if test.PropInt16() != i0 {
		t.Fatal("PropInt16")
	}

	test.ConnectPropInt162Changed(func(propInt162 uint16) {
		if propInt162 != i1 {
			t.Fatal(propInt162, i1)
		}
	})
	test.SetPropInt162(i1)
	test.PropInt162Changed(i1)
	if test.PropInt162() != i1 {
		t.Fatal("PropInt162")
	}

	test.ConnectPropInt32Changed(func(propInt32 int) {
		if propInt32 != i2 {
			t.Fatal(propInt32, i2)
		}
	})
	test.SetPropInt32(i2)
	test.PropInt32Changed(i2)
	if test.PropInt32() != i2 {
		t.Fatal("PropInt32")
	}

	test.ConnectPropInt322Changed(func(propInt322 uint) {
		if propInt322 != i3 {
			t.Fatal(propInt322, i3)
		}
	})
	test.SetPropInt322(i3)
	test.PropInt322Changed(i3)
	if test.PropInt322() != i3 {
		t.Fatal("PropInt322")
	}

	test.ConnectPropIntChanged(func(propInt int) {
		if propInt != i2 {
			t.Fatal(propInt, i2)
		}
	})
	test.SetPropInt(i2)
	test.PropIntChanged(i2)
	if test.PropInt() != i2 {
		t.Fatal("PropInt")
	}

	test.ConnectPropInt2Changed(func(propInt2 uint) {
		if propInt2 != i3 {
			t.Fatal(propInt2, i3)
		}
	})
	test.SetPropInt2(i3)
	test.PropInt2Changed(i3)
	if test.PropInt2() != i3 {
		t.Fatal("PropInt2")
	}

	test.ConnectPropInt64Changed(func(propInt64 int64) {
		if propInt64 != i6 {
			t.Fatal(propInt64, i6)
		}
	})
	test.SetPropInt64(i6)
	test.PropInt64Changed(i6)
	if test.PropInt64() != i6 {
		t.Fatal("PropInt64")
	}

	test.ConnectPropInt642Changed(func(propInt642 uint64) {
		if propInt642 != i7 {
			t.Fatal(propInt642, i7)
		}
	})
	test.SetPropInt642(i7)
	test.PropInt642Changed(i7)
	if test.PropInt642() != i7 {
		t.Fatal("PropInt642")
	}

	test.ConnectPropFloatChanged(func(propFloat float32) {
		if propFloat != f0 {
			t.Fatal(propFloat, f0)
		}
	})
	test.SetPropFloat(f0)
	test.PropFloatChanged(f0)
	if test.PropFloat() != f0 {
		t.Fatal("PropFloat")
	}

	test.ConnectPropFloat2Changed(func(propFloat2 float64) {
		if propFloat2 != f1 {
			t.Fatal(propFloat2, f1)
		}
	})
	test.SetPropFloat2(f1)
	test.PropFloat2Changed(f1)
	if test.PropFloat2() != f1 {
		t.Fatal("PropFloat2")
	}

	test.ConnectPropStringChanged(func(propString string) {
		if propString != s0 {
			t.Fatal(propString, s0)
		}
	})
	test.SetPropString(s0)
	test.PropStringChanged(s0)
	if test.PropString() != s0 {
		t.Fatal("PropString")
	}

	test.ConnectPropString2Changed(func(propString2 []string) {
		if strings.Join(propString2, "") != strings.Join(s1, "") {
			t.Fatal(propString2, s1)
		}
	})
	test.SetPropString2(s1)
	test.PropString2Changed(s1)
	if strings.Join(test.PropString2(), "") != strings.Join(s1, "") {
		t.Fatal("PropString2")
	}

	test.ConnectPropPointerChanged(func(propPointer uintptr) {
		if int(propPointer) != int(p0) {
			t.Fatal(propPointer, p0, int(propPointer), int(p0))
		}
	})
	test.SetPropPointer(p0)
	test.PropPointerChanged(p0)
	if int(test.PropPointer()) != int(p0) {
		t.Fatal("PropPointer")
	}

	test.ConnectPropPointer2Changed(func(propPointer2 unsafe.Pointer) {
		if int(uintptr(propPointer2)) != int(uintptr(p1)) {
			t.Fatal(propPointer2, p1, int(uintptr(propPointer2)), int(uintptr(p1)))
		}
	})
	test.SetPropPointer2(p1)
	test.PropPointer2Changed(p1)
	if int(uintptr(test.PropPointer2())) != int(uintptr(p1)) {
		t.Fatal("PropPointer2")
	}

	test.ConnectPropObjectChanged(func(propObject *core.QVariant) {
		if propObject.ToString() != o0.ToString() { //TODO:
			t.Fatal(propObject, o0, propObject.ToString(), o0.ToString())
		}
	})
	o0 = core.NewQVariant14("test")
	test.SetPropObject(o0)
	test.PropObjectChanged(o0)
	if test.PropObject().ToString() != o0.ToString() {
		t.Fatal("PropObject")
	}

	test.ConnectPropObject2Changed(func(propObject2 *core.QObject) {
		if propObject2.ObjectName() != o1.ObjectName() {
			t.Fatal(propObject2, o1, propObject2.ObjectName(), o1.ObjectName())
		}
	})
	o1 = core.NewQObject(nil)
	o1.SetObjectName("test")
	test.SetPropObject2(o1)
	test.PropObject2Changed(o1)
	if test.PropObject2().ObjectName() != o1.ObjectName() {
		t.Fatal("PropObject2")
	}

	//TODO: ConnectPropObject3Changed *QVariant

	test.ConnectPropEnumChanged(func(propEnum core.Qt__Key) {
		if propEnum != e0 {
			t.Fatal(propEnum, e0)
		}
	})
	test.SetPropEnum(e0)
	test.PropEnumChanged(e0)
	if test.PropEnum() != e0 {
		t.Fatal("PropEnum")
	}

	test.ConnectPropErrorChanged(func(propError error) {
		if propError.Error() != e1.Error() {
			t.Fatal(propError, e1, propError.Error(), e1.Error())
		}
	})
	test.SetPropError(e1)
	test.PropErrorChanged(e1)
	if test.PropError().Error() != e1.Error() {
		t.Fatal("PropError")
	}

	var sTest = NewOtherTestStruct(nil)
	sTest.SetObjectName("test")
	test.ConnectPropReturnTestChanged(func(propReturnTest *otherTestStruct) {
		if propReturnTest.ObjectName() != sTest.ObjectName() {
			t.Fatal(propReturnTest, sTest, propReturnTest.ObjectName(), sTest.ObjectName())
		}
	})
	test.SetPropReturnTest(sTest)
	test.PropReturnTestChanged(sTest)
	if test.PropReturnTest().ObjectName() != sTest.ObjectName() {
		t.Fatal("PropReturnTest")
	}

	test.ConnectPropReturnTest2Changed(func(propReturnTest2 *otherTestStruct) {
		if propReturnTest2.ObjectName() != sTest.ObjectName() {
			t.Fatal(propReturnTest2, sTest, propReturnTest2.ObjectName(), sTest.ObjectName())
		}
	})
	test.SetPropReturnTest2(sTest)
	test.PropReturnTest2Changed(sTest)
	if test.PropReturnTest2().ObjectName() != sTest.ObjectName() {
		t.Fatal("PropReturnTest2")
	}

	test.ConnectPropBoolSubChanged(func(propBoolSub bool) {
		if propBoolSub != b1 {
			t.Fatal(propBoolSub, b1)
		}
	})
	test.SetPropBoolSub(b1)
	test.PropBoolSubChanged(b1)
	if test.IsPropBoolSub() != b1 {
		t.Fatal("IsPropBoolSub")
	}
}

func TestSignalInput(t *testing.T) {
	var test = NewTestStruct(nil)

	test.ConnectBoolSignalInput(func(v0, v1 bool) {
		if v0 != b0 || v1 != b1 {
			t.Fatal(v0, b0, v1, b1)
		}
	})

	test.ConnectInt8SignalInput(func(v0 string, v1 string) {
		if v0 != s0 || v1 != s0 {
			t.Fatal(v0, s0, v1, s0)
		}
	})

	test.ConnectInt16SignalInput(func(v0 int16, v1 uint16) {
		if v0 != i0 || v1 != i1 {
			t.Fatal(v0, i0, v1, i1)
		}
	})

	test.ConnectInt32SignalInput(func(v0 int, v1 uint) {
		if v0 != i2 || v1 != i3 {
			t.Fatal(v0, i2, v1, i3)
		}
	})

	test.ConnectIntSignalInput(func(v0 int, v1 uint) {
		if v0 != i4 || v1 != i5 {
			t.Fatal(v0, i4, v1, i5)
		}
	})

	test.ConnectInt64SignalInput(func(v0 int64, v1 uint64) {
		if v0 != i6 || v1 != i7 {
			t.Fatal(v0, i6, v1, i7)
		}
	})

	test.ConnectFloatSignalInput(func(v0 float32, v1 float64) {
		if v0 != f0 || v1 != f1 {
			t.Fatal(v0, f0, v1, f1)
		}
	})

	test.ConnectStringSignalInput(func(v0 string, v1 []string) {
		if v0 != s0 || strings.Join(v1, "") != strings.Join(s1, "") {
			t.Fatal(v0, s0, v1, s1)
		}
	})

	test.ConnectPointerSignalInput(func(v0 uintptr, v1 unsafe.Pointer) {
		if int(v0) != int(p0) || int(uintptr(v1)) != int(uintptr(p1)) {
			t.Fatal(v0, p0, v1, p1, int(v0), int(p0), int(uintptr(v1)), int(uintptr(p1)))
		}
	})

	test.ConnectObjectSignalInput(func(v0 *core.QVariant, v1 *core.QObject) {
		if v0.ToString() != o0.ToString() || v1.ObjectName() != o1.ObjectName() {
			t.Fatal(v0, o0, v1, o1)
		}
	})

	test.ConnectEnumSignalInput(func(v0 core.Qt__Key) {
		if v0 != e0 {
			t.Fatal(v0, e0)
		}
	})

	test.ConnectErrorSignalInput(func(v0 error) {
		if v0.Error() != e1.Error() {
			t.Fatal(v0, e1)
		}
	})

	test.BoolSignalInput(b0, b1)
	//test.Int8SignalInput(s0, s0)
	test.Int16SignalInput(i0, i1)
	test.Int32SignalInput(i2, i3)
	test.IntSignalInput(i4, i5)
	test.Int64SignalInput(i6, i7)
	test.FloatSignalInput(f0, f1)
	test.StringSignalInput(s0, s1)
	test.PointerSignalInput(p0, p1)

	o0 = core.NewQVariant14("test")
	o1 = core.NewQObject(nil)
	o1.SetObjectName("test")
	test.ObjectSignalInput(o0, o1)

	test.EnumSignalInput(e0)
	test.ErrorSignalInput(e1)
}

func TestSlotInput(t *testing.T) {
	var test = NewTestStruct(nil)

	test.ConnectBoolSlotInput(func(v0, v1 bool) {
		if v0 != b0 || v1 != b1 {
			t.Fatal(v0, b0, v1, b1)
		}
	})

	test.ConnectInt8SlotInput(func(v0 string, v1 string) {
		if v0 != s0 || v1 != s0 {
			t.Fatal(v0, s0, v1, s0)
		}
	})

	test.ConnectInt16SlotInput(func(v0 int16, v1 uint16) {
		if v0 != i0 || v1 != i1 {
			t.Fatal(v0, i0, v1, i1)
		}
	})

	test.ConnectInt32SlotInput(func(v0 int, v1 uint) {
		if v0 != i2 || v1 != i3 {
			t.Fatal(v0, i2, v1, i3)
		}
	})

	test.ConnectIntSlotInput(func(v0 int, v1 uint) {
		if v0 != i4 || v1 != i5 {
			t.Fatal(v0, i4, v1, i5)
		}
	})

	test.ConnectInt64SlotInput(func(v0 int64, v1 uint64) {
		if v0 != i6 || v1 != i7 {
			t.Fatal(v0, i6, v1, i7)
		}
	})

	test.ConnectFloatSlotInput(func(v0 float32, v1 float64) {
		if v0 != f0 || v1 != f1 {
			t.Fatal(v0, f0, v1, f1)
		}
	})

	test.ConnectStringSlotInput(func(v0 string, v1 []string) {
		if v0 != s0 || strings.Join(v1, "") != strings.Join(s1, "") {
			t.Fatal(v0, s0, v1, s1)
		}
	})

	test.ConnectPointerSlotInput(func(v0 uintptr, v1 unsafe.Pointer) {
		if int(v0) != int(p0) || int(uintptr(v1)) != int(uintptr(p1)) {
			t.Fatal(v0, p0, v1, p1, int(v0), int(p0), int(uintptr(v1)), int(uintptr(p1)))
		}
	})

	test.ConnectObjectSlotInput(func(v0 *core.QVariant, v1 *core.QObject) {
		if v0.ToString() != o0.ToString() || v1.ObjectName() != o1.ObjectName() {
			t.Fatal(v0, o0, v1, o1)
		}
	})

	test.ConnectEnumSlotInput(func(v0 core.Qt__Key) {
		if v0 != e0 {
			t.Fatal(v0, e0)
		}
	})

	test.ConnectErrorSlotInput(func(v0 error) {
		if v0.Error() != e1.Error() {
			t.Fatal(v0, e1)
		}
	})

	test.BoolSlotInput(b0, b1)
	//test.Int8SlotInput(s0, s0)
	test.Int16SlotInput(i0, i1)
	test.Int32SlotInput(i2, i3)
	test.IntSlotInput(i4, i5)
	test.Int64SlotInput(i6, i7)
	test.FloatSlotInput(f0, f1)
	test.StringSlotInput(s0, s1)
	test.PointerSlotInput(p0, p1)

	o0 = core.NewQVariant14("test")
	o1 = core.NewQObject(nil)
	o1.SetObjectName("test")
	test.ObjectSlotInput(o0, o1)

	test.EnumSlotInput(e0)
	test.ErrorSlotInput(e1)
}

func TestSlotOutput(t *testing.T) {
	var test = NewTestStruct(nil)

	test.ConnectBoolSlotOutput(func(v0 bool) bool { return v0 })
	test.ConnectBoolSlotOutput2(func(v0 bool) bool { return v0 })

	test.ConnectInt8SlotOutput(func(v0 string) string { return v0 })
	test.ConnectInt8SlotOutput2(func(v0 string) string { return v0 })

	test.ConnectInt16SlotOutput(func(v0 int16) int16 { return v0 })
	test.ConnectInt16SlotOutput2(func(v0 uint16) uint16 { return v0 })

	test.ConnectInt32SlotOutput(func(v0 int) int { return v0 })
	test.ConnectInt32SlotOutput2(func(v0 uint) uint { return v0 })

	test.ConnectIntSlotOutput(func(v0 int) int { return v0 })
	test.ConnectIntSlotOutput2(func(v0 uint) uint { return v0 })

	test.ConnectInt64SlotOutput(func(v0 int64) int64 { return v0 })
	test.ConnectInt64SlotOutput2(func(v0 uint64) uint64 { return v0 })

	test.ConnectFloatSlotOutput(func(v0 float32) float32 { return v0 })
	test.ConnectFloatSlotOutput2(func(v0 float64) float64 { return v0 })

	test.ConnectStringSlotOutput(func(v0 string) string { return v0 })
	test.ConnectStringSlotOutput2(func(v0 []string) []string { return v0 })

	test.ConnectPointerSlotOutput(func(v0 uintptr) uintptr { return v0 })
	test.ConnectPointerSlotOutput2(func(v0 unsafe.Pointer) unsafe.Pointer { return v0 })

	test.ConnectObjectSlotOutput(func(v0 *core.QVariant) *core.QVariant { return v0 })
	test.ConnectObjectSlotOutput2(func(v0 *core.QObject) *core.QObject { return v0 })

	test.ConnectEnumSlotOutput(func(v0 core.Qt__Key) core.Qt__Key { return v0 })
	test.ConnectErrorSlotOutput(func(v0 error) error { return v0 })

	test.ConnectReturnTest(func(v0 *testStruct) *testStruct { return v0 })
	test.ConnectReturnTest2(func(v0 *testStruct) *testStruct { return v0 })

	test.ConnectReturnName(func(a0 string) string { return a0 })
	test.ConnectReturnName2(func(a0 int, a1 int) int { return a0 + a1 })

	if test.BoolSlotOutput(b0) != b0 {
		t.Fatal("BoolSlotOutput")
	}
	if test.BoolSlotOutput2(b1) != b1 {
		t.Fatal("BoolSlotOutput2")
	}

	/*
		if test.Int8SlotOutput(s0) != s0 {
			t.Fatal("Int8SlotOutput")
		}
		if test.Int8SlotOutput2(s0) != s0 {
			t.Fatal("Int8SlotOutput2")
		}
	*/

	if test.Int16SlotOutput(i0) != i0 {
		t.Fatal("Int16SlotOutput")
	}
	if test.Int16SlotOutput2(i1) != i1 {
		t.Fatal("Int16SlotOutput2")
	}

	if test.Int32SlotOutput(i2) != i2 {
		t.Fatal("Int32SlotOutput")
	}
	if test.Int32SlotOutput2(i3) != i3 {
		t.Fatal("Int32SlotOutput2")
	}

	if test.IntSlotOutput(i4) != i4 {
		t.Fatal("IntSlotOutput")
	}
	if test.IntSlotOutput2(i5) != i5 {
		t.Fatal("IntSlotOutput2")
	}

	if test.Int64SlotOutput(i6) != i6 {
		t.Fatal("Int64SlotOutput")
	}
	if test.Int64SlotOutput2(i7) != i7 {
		t.Fatal("Int64SlotOutput2")
	}

	if test.FloatSlotOutput(f0) != f0 {
		t.Fatal("FloatSlotOutput")
	}
	if test.FloatSlotOutput2(f1) != f1 {
		t.Fatal("FloatSlotOutput2")
	}

	if test.StringSlotOutput(s0) != s0 {
		t.Fatal("StringSlotOutput")
	}
	if strings.Join(test.StringSlotOutput2(s1), "") != strings.Join(s1, "") {
		t.Fatal("StringSlotOutput2")
	}

	if int(test.PointerSlotOutput(p0)) != int(p0) {
		t.Fatal("PointerSlotOutput")
	}
	if int(uintptr(test.PointerSlotOutput2(p1))) != int(uintptr(p1)) {
		t.Fatal("PointerSlotOutput2")
	}

	o0 = core.NewQVariant14("test")
	o1 = core.NewQObject(nil)
	o1.SetObjectName("test")

	if test.ObjectSlotOutput(o0).ToString() != o0.ToString() {
		t.Fatal("ObjectSlotOutput")
	}
	if test.ObjectSlotOutput2(o1).ObjectName() != o1.ObjectName() {
		t.Fatal("ObjectSlotOutput2")
	}

	if test.EnumSlotOutput(e0) != e0 {
		t.Fatal("EnumSlotOutput")
	}

	if test.ErrorSlotOutput(e1).Error() != e1.Error() {
		t.Fatal("ErrorSlotOutput")
	}

	test.SetObjectName("testName")
	if test.ReturnTest(test.ReturnTest(test)).ReturnTest(test).ObjectName() != test.ObjectName() {
		t.Fatal("ReturnTest")
	}
	if test.ReturnTest2(test.ReturnTest2(test)).ReturnTest2(test).ObjectName() != test.ObjectName() {
		t.Fatal("ReturnTest2")
	}

	if test.ReturnName(s0) != s0 {
		t.Fatal("ReturnName")
	}
	if test.ReturnName2(i2, i2) != i2*2 {
		t.Fatal("ReturnName2")
	}

	var sTest = NewSubTestStruct(nil)
	sTest.ConnectReturnTest3(func(v0 *subTestStruct) *subTestStruct { return v0 })
	sTest.SetObjectName("testSubName")
	if sTest.ReturnTest3(sTest.ReturnTest3(sTest)).ReturnTest3(sTest).ObjectName() != sTest.ObjectName() {
		t.Fatal("ReturnName3")
	}
}
