package qt

//#include "qdesktopservices.h"
import "C"

func QDesktopServices_OpenUrl(url string) bool {
	return C.QDesktopServices_OpenUrl_String(C.CString(url)) != 0
}

func QDesktopServices_SetUrlHandler(scheme string, receiver QObject, method string) {
	var receiverPtr C.QtObjectPtr
	if receiver != nil {
		receiverPtr = receiver.Pointer()
	}
	C.QDesktopServices_SetUrlHandler_String_QObject_String(C.CString(scheme), receiverPtr, C.CString(method))
}

func QDesktopServices_UnsetUrlHandler(scheme string) {
	C.QDesktopServices_UnsetUrlHandler_String(C.CString(scheme))
}
