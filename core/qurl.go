package core

//#include "qurl.h"
import "C"
import (
	"unsafe"
)

type QUrl struct {
	ptr unsafe.Pointer
}

type QUrlITF interface {
	QUrlPTR() *QUrl
}

func (p *QUrl) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QUrl) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQUrl(ptr QUrlITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QUrlPTR().Pointer()
	}
	return nil
}

func QUrlFromPointer(ptr unsafe.Pointer) *QUrl {
	var n = new(QUrl)
	n.SetPointer(ptr)
	return n
}

func (ptr *QUrl) QUrlPTR() *QUrl {
	return ptr
}

//QUrl::ComponentFormattingOption
type QUrl__ComponentFormattingOption int

var (
	QUrl__PrettyDecoded    = QUrl__ComponentFormattingOption(0x000000)
	QUrl__EncodeSpaces     = QUrl__ComponentFormattingOption(0x100000)
	QUrl__EncodeUnicode    = QUrl__ComponentFormattingOption(0x200000)
	QUrl__EncodeDelimiters = QUrl__ComponentFormattingOption(0x400000 | 0x800000)
	QUrl__EncodeReserved   = QUrl__ComponentFormattingOption(0x1000000)
	QUrl__DecodeReserved   = QUrl__ComponentFormattingOption(0x2000000)
	QUrl__FullyEncoded     = QUrl__ComponentFormattingOption(QUrl__EncodeSpaces | QUrl__EncodeUnicode | QUrl__EncodeDelimiters | QUrl__EncodeReserved)
	QUrl__FullyDecoded     = QUrl__ComponentFormattingOption(QUrl__FullyEncoded | QUrl__DecodeReserved | 0x4000000)
)

//QUrl::ParsingMode
type QUrl__ParsingMode int

var (
	QUrl__TolerantMode = QUrl__ParsingMode(0)
	QUrl__StrictMode   = QUrl__ParsingMode(1)
	QUrl__DecodedMode  = QUrl__ParsingMode(2)
)

//QUrl::UrlFormattingOption
type QUrl__UrlFormattingOption int

var (
	QUrl__None                  = QUrl__UrlFormattingOption(0x0)
	QUrl__RemoveScheme          = QUrl__UrlFormattingOption(0x1)
	QUrl__RemovePassword        = QUrl__UrlFormattingOption(0x2)
	QUrl__RemoveUserInfo        = QUrl__UrlFormattingOption(QUrl__RemovePassword | 0x4)
	QUrl__RemovePort            = QUrl__UrlFormattingOption(0x8)
	QUrl__RemoveAuthority       = QUrl__UrlFormattingOption(QUrl__RemoveUserInfo | QUrl__RemovePort | 0x10)
	QUrl__RemovePath            = QUrl__UrlFormattingOption(0x20)
	QUrl__RemoveQuery           = QUrl__UrlFormattingOption(0x40)
	QUrl__RemoveFragment        = QUrl__UrlFormattingOption(0x80)
	QUrl__PreferLocalFile       = QUrl__UrlFormattingOption(0x200)
	QUrl__StripTrailingSlash    = QUrl__UrlFormattingOption(0x400)
	QUrl__RemoveFilename        = QUrl__UrlFormattingOption(0x800)
	QUrl__NormalizePathSegments = QUrl__UrlFormattingOption(0x1000)
)

//QUrl::UserInputResolutionOption
type QUrl__UserInputResolutionOption int

var (
	QUrl__DefaultResolution = QUrl__UserInputResolutionOption(0)
	QUrl__AssumeLocalFile   = QUrl__UserInputResolutionOption(1)
)
