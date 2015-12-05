package androidextras

import (
	"C"
	"strings"
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
		case []string:
			{
				return QAndroidJniObject_FromString(strings.Join(input[key].([]string), ",,,")).CallObjectMethod2("split", "(Ljava/lang/String;)[Ljava/lang/String;", ",,,").Object()
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
