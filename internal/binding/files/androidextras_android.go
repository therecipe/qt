package androidextras

import (
	"C"
	"strings"
	"unsafe"

	"github.com/therecipe/qt"
)

func assertion(key int, input ...interface{}) unsafe.Pointer {
	if len(input) > key {
		switch deduced := input[key].(type) {
		case string:
			{
				return QAndroidJniObject_FromString(deduced).Object()
			}

		case []string:
			{
				return QAndroidJniObject_FromString(strings.Join(deduced, ",,,")).CallObjectMethod2("split", "(Ljava/lang/String;)[Ljava/lang/String;", ",,,").Object()
			}

		case int:
			{
				return unsafe.Pointer(uintptr(C.int(deduced)))
			}

		case bool:
			{
				return unsafe.Pointer(uintptr(C.int(qt.GoBoolToInt(deduced))))
			}

		case unsafe.Pointer:
			{
				return deduced
			}

		case *QAndroidJniObject:
			{
				return deduced.Object()
			}
		}
	}
	return nil
}
