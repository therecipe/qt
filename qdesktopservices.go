package qt

//#include "qdesktopservices.h"
import "C"

func QDesktopServices_OpenUrl_String(url string) bool {
	return C.QDesktopServices_OpenUrl_String(C.CString(url)) != 0
}

func QDesktopServices_SetUrlHandler_String_QObject_String(scheme string, receiver QObject, method string) {
	var receiverPtr C.QtObjectPtr = nil
	if receiver != nil {
		receiverPtr = receiver.Pointer()
	}
	C.QDesktopServices_SetUrlHandler_String_QObject_String(C.CString(scheme), receiverPtr, C.CString(method))
}

func QDesktopServices_UnsetUrlHandler_String(scheme string) {
	C.QDesktopServices_UnsetUrlHandler_String(C.CString(scheme))
}
