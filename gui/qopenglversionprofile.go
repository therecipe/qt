package gui

//#include "qopenglversionprofile.h"
import "C"
import (
	"unsafe"
)

type QOpenGLVersionProfile struct {
	ptr unsafe.Pointer
}

type QOpenGLVersionProfileITF interface {
	QOpenGLVersionProfilePTR() *QOpenGLVersionProfile
}

func (p *QOpenGLVersionProfile) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QOpenGLVersionProfile) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQOpenGLVersionProfile(ptr QOpenGLVersionProfileITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenGLVersionProfilePTR().Pointer()
	}
	return nil
}

func QOpenGLVersionProfileFromPointer(ptr unsafe.Pointer) *QOpenGLVersionProfile {
	var n = new(QOpenGLVersionProfile)
	n.SetPointer(ptr)
	return n
}

func (ptr *QOpenGLVersionProfile) QOpenGLVersionProfilePTR() *QOpenGLVersionProfile {
	return ptr
}

func NewQOpenGLVersionProfile() *QOpenGLVersionProfile {
	return QOpenGLVersionProfileFromPointer(unsafe.Pointer(C.QOpenGLVersionProfile_NewQOpenGLVersionProfile()))
}

func NewQOpenGLVersionProfile3(other QOpenGLVersionProfileITF) *QOpenGLVersionProfile {
	return QOpenGLVersionProfileFromPointer(unsafe.Pointer(C.QOpenGLVersionProfile_NewQOpenGLVersionProfile3(C.QtObjectPtr(PointerFromQOpenGLVersionProfile(other)))))
}

func NewQOpenGLVersionProfile2(format QSurfaceFormatITF) *QOpenGLVersionProfile {
	return QOpenGLVersionProfileFromPointer(unsafe.Pointer(C.QOpenGLVersionProfile_NewQOpenGLVersionProfile2(C.QtObjectPtr(PointerFromQSurfaceFormat(format)))))
}

func (ptr *QOpenGLVersionProfile) HasProfiles() bool {
	if ptr.Pointer() != nil {
		return C.QOpenGLVersionProfile_HasProfiles(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QOpenGLVersionProfile) IsLegacyVersion() bool {
	if ptr.Pointer() != nil {
		return C.QOpenGLVersionProfile_IsLegacyVersion(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QOpenGLVersionProfile) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QOpenGLVersionProfile_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QOpenGLVersionProfile) Profile() QSurfaceFormat__OpenGLContextProfile {
	if ptr.Pointer() != nil {
		return QSurfaceFormat__OpenGLContextProfile(C.QOpenGLVersionProfile_Profile(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QOpenGLVersionProfile) SetProfile(profile QSurfaceFormat__OpenGLContextProfile) {
	if ptr.Pointer() != nil {
		C.QOpenGLVersionProfile_SetProfile(C.QtObjectPtr(ptr.Pointer()), C.int(profile))
	}
}

func (ptr *QOpenGLVersionProfile) SetVersion(majorVersion int, minorVersion int) {
	if ptr.Pointer() != nil {
		C.QOpenGLVersionProfile_SetVersion(C.QtObjectPtr(ptr.Pointer()), C.int(majorVersion), C.int(minorVersion))
	}
}

func (ptr *QOpenGLVersionProfile) DestroyQOpenGLVersionProfile() {
	if ptr.Pointer() != nil {
		C.QOpenGLVersionProfile_DestroyQOpenGLVersionProfile(C.QtObjectPtr(ptr.Pointer()))
	}
}
