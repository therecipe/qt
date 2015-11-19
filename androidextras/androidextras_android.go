package androidextras

import (
	"C"
	"unsafe"

	"github.com/therecipe/qt"
)

func assertion(key int, input ...interface{}) unsafe.Pointer {
	if len(input) > key {
		switch input[key].(type) {
		case string:
			{
				return QAndroidJniObject_FromString(input[key].(string)).Object()
			}
		case int:
			{
				return unsafe.Pointer(uintptr(C.int(input[key].(int))))
			}
		case bool:
			{
				return unsafe.Pointer(uintptr(C.int(qt.GoBoolToInt(input[key].(bool)))))
			}
		case unsafe.Pointer:
			{
				return input[key].(unsafe.Pointer)
			}
		case *QAndroidJniObject:
			{
				return input[key].(*QAndroidJniObject).Object()
			}
		}
	}
	return nil
}
