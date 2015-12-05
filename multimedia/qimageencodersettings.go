package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QImageEncoderSettings struct {
	ptr unsafe.Pointer
}

type QImageEncoderSettings_ITF interface {
	QImageEncoderSettings_PTR() *QImageEncoderSettings
}

func (p *QImageEncoderSettings) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QImageEncoderSettings) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQImageEncoderSettings(ptr QImageEncoderSettings_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QImageEncoderSettings_PTR().Pointer()
	}
	return nil
}

func NewQImageEncoderSettingsFromPointer(ptr unsafe.Pointer) *QImageEncoderSettings {
	var n = new(QImageEncoderSettings)
	n.SetPointer(ptr)
	return n
}

func (ptr *QImageEncoderSettings) QImageEncoderSettings_PTR() *QImageEncoderSettings {
	return ptr
}

func NewQImageEncoderSettings() *QImageEncoderSettings {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QImageEncoderSettings::QImageEncoderSettings")
		}
	}()

	return NewQImageEncoderSettingsFromPointer(C.QImageEncoderSettings_NewQImageEncoderSettings())
}

func NewQImageEncoderSettings2(other QImageEncoderSettings_ITF) *QImageEncoderSettings {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QImageEncoderSettings::QImageEncoderSettings")
		}
	}()

	return NewQImageEncoderSettingsFromPointer(C.QImageEncoderSettings_NewQImageEncoderSettings2(PointerFromQImageEncoderSettings(other)))
}

func (ptr *QImageEncoderSettings) Codec() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QImageEncoderSettings::codec")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QImageEncoderSettings_Codec(ptr.Pointer()))
	}
	return ""
}

func (ptr *QImageEncoderSettings) EncodingOption(option string) *core.QVariant {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QImageEncoderSettings::encodingOption")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QImageEncoderSettings_EncodingOption(ptr.Pointer(), C.CString(option)))
	}
	return nil
}

func (ptr *QImageEncoderSettings) IsNull() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QImageEncoderSettings::isNull")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QImageEncoderSettings_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QImageEncoderSettings) SetCodec(codec string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QImageEncoderSettings::setCodec")
		}
	}()

	if ptr.Pointer() != nil {
		C.QImageEncoderSettings_SetCodec(ptr.Pointer(), C.CString(codec))
	}
}

func (ptr *QImageEncoderSettings) SetEncodingOption(option string, value core.QVariant_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QImageEncoderSettings::setEncodingOption")
		}
	}()

	if ptr.Pointer() != nil {
		C.QImageEncoderSettings_SetEncodingOption(ptr.Pointer(), C.CString(option), core.PointerFromQVariant(value))
	}
}

func (ptr *QImageEncoderSettings) SetResolution(resolution core.QSize_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QImageEncoderSettings::setResolution")
		}
	}()

	if ptr.Pointer() != nil {
		C.QImageEncoderSettings_SetResolution(ptr.Pointer(), core.PointerFromQSize(resolution))
	}
}

func (ptr *QImageEncoderSettings) SetResolution2(width int, height int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QImageEncoderSettings::setResolution")
		}
	}()

	if ptr.Pointer() != nil {
		C.QImageEncoderSettings_SetResolution2(ptr.Pointer(), C.int(width), C.int(height))
	}
}

func (ptr *QImageEncoderSettings) DestroyQImageEncoderSettings() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QImageEncoderSettings::~QImageEncoderSettings")
		}
	}()

	if ptr.Pointer() != nil {
		C.QImageEncoderSettings_DestroyQImageEncoderSettings(ptr.Pointer())
	}
}
