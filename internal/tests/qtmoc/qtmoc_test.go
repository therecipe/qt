package qtmoc

import (
	"os"
	"strings"
	"testing"
	"unsafe"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"

	"github.com/therecipe/qt/internal/tests/qtmoc/sub"
)

//go:generate qtmoc
type testStruct struct {
	otherTestStruct
	testOther otherTestStruct

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

type subSubTestStruct struct { //Qt structs from other pkgs will be ignored
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
