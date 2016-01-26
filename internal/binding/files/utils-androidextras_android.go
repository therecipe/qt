package androidextras

import (
	"C"
	"strings"
	"unsafe"

	"github.com/therecipe/qt"
)

func assertion(key int, input ...interface{}) (unsafe.Pointer, func()) {
	if len(input) > key {
		switch deduced := input[key].(type) {
		case string:
			{
				var jObject = QAndroidJniObject_FromString(deduced)

				return jObject.Object(), func() { jObject.DestroyQAndroidJniObject() }
			}

		case []string:
			{
				var (
					jObject  = QAndroidJniObject_FromString(strings.Join(deduced, ",,,"))
					jObject2 = jObject.CallObjectMethod2("split", "(Ljava/lang/String;)[Ljava/lang/String;", ",,,")
				)

				jObject.DestroyQAndroidJniObject()

				return jObject2.Object(), func() { jObject2.DestroyQAndroidJniObject() }
			}

		case int:
			{
				return unsafe.Pointer(uintptr(C.int(deduced))), nil
			}

		case bool:
			{
				return unsafe.Pointer(uintptr(C.int(qt.GoBoolToInt(deduced)))), nil
			}

		case unsafe.Pointer:
			{
				return deduced, nil
			}

		case *QAndroidJniObject:
			{
				return deduced.Object(), nil
			}
		}
	}
	return nil, nil
}
