package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QPagedPaintDevice struct {
	QPaintDevice
}

type QPagedPaintDevice_ITF interface {
	QPaintDevice_ITF
	QPagedPaintDevice_PTR() *QPagedPaintDevice
}

func PointerFromQPagedPaintDevice(ptr QPagedPaintDevice_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPagedPaintDevice_PTR().Pointer()
	}
	return nil
}

func NewQPagedPaintDeviceFromPointer(ptr unsafe.Pointer) *QPagedPaintDevice {
	var n = new(QPagedPaintDevice)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPagedPaintDevice) QPagedPaintDevice_PTR() *QPagedPaintDevice {
	return ptr
}

//QPagedPaintDevice::PageSize
type QPagedPaintDevice__PageSize int64

const (
	QPagedPaintDevice__A4                 = QPagedPaintDevice__PageSize(0)
	QPagedPaintDevice__B5                 = QPagedPaintDevice__PageSize(1)
	QPagedPaintDevice__Letter             = QPagedPaintDevice__PageSize(2)
	QPagedPaintDevice__Legal              = QPagedPaintDevice__PageSize(3)
	QPagedPaintDevice__Executive          = QPagedPaintDevice__PageSize(4)
	QPagedPaintDevice__A0                 = QPagedPaintDevice__PageSize(5)
	QPagedPaintDevice__A1                 = QPagedPaintDevice__PageSize(6)
	QPagedPaintDevice__A2                 = QPagedPaintDevice__PageSize(7)
	QPagedPaintDevice__A3                 = QPagedPaintDevice__PageSize(8)
	QPagedPaintDevice__A5                 = QPagedPaintDevice__PageSize(9)
	QPagedPaintDevice__A6                 = QPagedPaintDevice__PageSize(10)
	QPagedPaintDevice__A7                 = QPagedPaintDevice__PageSize(11)
	QPagedPaintDevice__A8                 = QPagedPaintDevice__PageSize(12)
	QPagedPaintDevice__A9                 = QPagedPaintDevice__PageSize(13)
	QPagedPaintDevice__B0                 = QPagedPaintDevice__PageSize(14)
	QPagedPaintDevice__B1                 = QPagedPaintDevice__PageSize(15)
	QPagedPaintDevice__B10                = QPagedPaintDevice__PageSize(16)
	QPagedPaintDevice__B2                 = QPagedPaintDevice__PageSize(17)
	QPagedPaintDevice__B3                 = QPagedPaintDevice__PageSize(18)
	QPagedPaintDevice__B4                 = QPagedPaintDevice__PageSize(19)
	QPagedPaintDevice__B6                 = QPagedPaintDevice__PageSize(20)
	QPagedPaintDevice__B7                 = QPagedPaintDevice__PageSize(21)
	QPagedPaintDevice__B8                 = QPagedPaintDevice__PageSize(22)
	QPagedPaintDevice__B9                 = QPagedPaintDevice__PageSize(23)
	QPagedPaintDevice__C5E                = QPagedPaintDevice__PageSize(24)
	QPagedPaintDevice__Comm10E            = QPagedPaintDevice__PageSize(25)
	QPagedPaintDevice__DLE                = QPagedPaintDevice__PageSize(26)
	QPagedPaintDevice__Folio              = QPagedPaintDevice__PageSize(27)
	QPagedPaintDevice__Ledger             = QPagedPaintDevice__PageSize(28)
	QPagedPaintDevice__Tabloid            = QPagedPaintDevice__PageSize(29)
	QPagedPaintDevice__Custom             = QPagedPaintDevice__PageSize(30)
	QPagedPaintDevice__A10                = QPagedPaintDevice__PageSize(31)
	QPagedPaintDevice__A3Extra            = QPagedPaintDevice__PageSize(32)
	QPagedPaintDevice__A4Extra            = QPagedPaintDevice__PageSize(33)
	QPagedPaintDevice__A4Plus             = QPagedPaintDevice__PageSize(34)
	QPagedPaintDevice__A4Small            = QPagedPaintDevice__PageSize(35)
	QPagedPaintDevice__A5Extra            = QPagedPaintDevice__PageSize(36)
	QPagedPaintDevice__B5Extra            = QPagedPaintDevice__PageSize(37)
	QPagedPaintDevice__JisB0              = QPagedPaintDevice__PageSize(38)
	QPagedPaintDevice__JisB1              = QPagedPaintDevice__PageSize(39)
	QPagedPaintDevice__JisB2              = QPagedPaintDevice__PageSize(40)
	QPagedPaintDevice__JisB3              = QPagedPaintDevice__PageSize(41)
	QPagedPaintDevice__JisB4              = QPagedPaintDevice__PageSize(42)
	QPagedPaintDevice__JisB5              = QPagedPaintDevice__PageSize(43)
	QPagedPaintDevice__JisB6              = QPagedPaintDevice__PageSize(44)
	QPagedPaintDevice__JisB7              = QPagedPaintDevice__PageSize(45)
	QPagedPaintDevice__JisB8              = QPagedPaintDevice__PageSize(46)
	QPagedPaintDevice__JisB9              = QPagedPaintDevice__PageSize(47)
	QPagedPaintDevice__JisB10             = QPagedPaintDevice__PageSize(48)
	QPagedPaintDevice__AnsiC              = QPagedPaintDevice__PageSize(49)
	QPagedPaintDevice__AnsiD              = QPagedPaintDevice__PageSize(50)
	QPagedPaintDevice__AnsiE              = QPagedPaintDevice__PageSize(51)
	QPagedPaintDevice__LegalExtra         = QPagedPaintDevice__PageSize(52)
	QPagedPaintDevice__LetterExtra        = QPagedPaintDevice__PageSize(53)
	QPagedPaintDevice__LetterPlus         = QPagedPaintDevice__PageSize(54)
	QPagedPaintDevice__LetterSmall        = QPagedPaintDevice__PageSize(55)
	QPagedPaintDevice__TabloidExtra       = QPagedPaintDevice__PageSize(56)
	QPagedPaintDevice__ArchA              = QPagedPaintDevice__PageSize(57)
	QPagedPaintDevice__ArchB              = QPagedPaintDevice__PageSize(58)
	QPagedPaintDevice__ArchC              = QPagedPaintDevice__PageSize(59)
	QPagedPaintDevice__ArchD              = QPagedPaintDevice__PageSize(60)
	QPagedPaintDevice__ArchE              = QPagedPaintDevice__PageSize(61)
	QPagedPaintDevice__Imperial7x9        = QPagedPaintDevice__PageSize(62)
	QPagedPaintDevice__Imperial8x10       = QPagedPaintDevice__PageSize(63)
	QPagedPaintDevice__Imperial9x11       = QPagedPaintDevice__PageSize(64)
	QPagedPaintDevice__Imperial9x12       = QPagedPaintDevice__PageSize(65)
	QPagedPaintDevice__Imperial10x11      = QPagedPaintDevice__PageSize(66)
	QPagedPaintDevice__Imperial10x13      = QPagedPaintDevice__PageSize(67)
	QPagedPaintDevice__Imperial10x14      = QPagedPaintDevice__PageSize(68)
	QPagedPaintDevice__Imperial12x11      = QPagedPaintDevice__PageSize(69)
	QPagedPaintDevice__Imperial15x11      = QPagedPaintDevice__PageSize(70)
	QPagedPaintDevice__ExecutiveStandard  = QPagedPaintDevice__PageSize(71)
	QPagedPaintDevice__Note               = QPagedPaintDevice__PageSize(72)
	QPagedPaintDevice__Quarto             = QPagedPaintDevice__PageSize(73)
	QPagedPaintDevice__Statement          = QPagedPaintDevice__PageSize(74)
	QPagedPaintDevice__SuperA             = QPagedPaintDevice__PageSize(75)
	QPagedPaintDevice__SuperB             = QPagedPaintDevice__PageSize(76)
	QPagedPaintDevice__Postcard           = QPagedPaintDevice__PageSize(77)
	QPagedPaintDevice__DoublePostcard     = QPagedPaintDevice__PageSize(78)
	QPagedPaintDevice__Prc16K             = QPagedPaintDevice__PageSize(79)
	QPagedPaintDevice__Prc32K             = QPagedPaintDevice__PageSize(80)
	QPagedPaintDevice__Prc32KBig          = QPagedPaintDevice__PageSize(81)
	QPagedPaintDevice__FanFoldUS          = QPagedPaintDevice__PageSize(82)
	QPagedPaintDevice__FanFoldGerman      = QPagedPaintDevice__PageSize(83)
	QPagedPaintDevice__FanFoldGermanLegal = QPagedPaintDevice__PageSize(84)
	QPagedPaintDevice__EnvelopeB4         = QPagedPaintDevice__PageSize(85)
	QPagedPaintDevice__EnvelopeB5         = QPagedPaintDevice__PageSize(86)
	QPagedPaintDevice__EnvelopeB6         = QPagedPaintDevice__PageSize(87)
	QPagedPaintDevice__EnvelopeC0         = QPagedPaintDevice__PageSize(88)
	QPagedPaintDevice__EnvelopeC1         = QPagedPaintDevice__PageSize(89)
	QPagedPaintDevice__EnvelopeC2         = QPagedPaintDevice__PageSize(90)
	QPagedPaintDevice__EnvelopeC3         = QPagedPaintDevice__PageSize(91)
	QPagedPaintDevice__EnvelopeC4         = QPagedPaintDevice__PageSize(92)
	QPagedPaintDevice__EnvelopeC6         = QPagedPaintDevice__PageSize(93)
	QPagedPaintDevice__EnvelopeC65        = QPagedPaintDevice__PageSize(94)
	QPagedPaintDevice__EnvelopeC7         = QPagedPaintDevice__PageSize(95)
	QPagedPaintDevice__Envelope9          = QPagedPaintDevice__PageSize(96)
	QPagedPaintDevice__Envelope11         = QPagedPaintDevice__PageSize(97)
	QPagedPaintDevice__Envelope12         = QPagedPaintDevice__PageSize(98)
	QPagedPaintDevice__Envelope14         = QPagedPaintDevice__PageSize(99)
	QPagedPaintDevice__EnvelopeMonarch    = QPagedPaintDevice__PageSize(100)
	QPagedPaintDevice__EnvelopePersonal   = QPagedPaintDevice__PageSize(101)
	QPagedPaintDevice__EnvelopeChou3      = QPagedPaintDevice__PageSize(102)
	QPagedPaintDevice__EnvelopeChou4      = QPagedPaintDevice__PageSize(103)
	QPagedPaintDevice__EnvelopeInvite     = QPagedPaintDevice__PageSize(104)
	QPagedPaintDevice__EnvelopeItalian    = QPagedPaintDevice__PageSize(105)
	QPagedPaintDevice__EnvelopeKaku2      = QPagedPaintDevice__PageSize(106)
	QPagedPaintDevice__EnvelopeKaku3      = QPagedPaintDevice__PageSize(107)
	QPagedPaintDevice__EnvelopePrc1       = QPagedPaintDevice__PageSize(108)
	QPagedPaintDevice__EnvelopePrc2       = QPagedPaintDevice__PageSize(109)
	QPagedPaintDevice__EnvelopePrc3       = QPagedPaintDevice__PageSize(110)
	QPagedPaintDevice__EnvelopePrc4       = QPagedPaintDevice__PageSize(111)
	QPagedPaintDevice__EnvelopePrc5       = QPagedPaintDevice__PageSize(112)
	QPagedPaintDevice__EnvelopePrc6       = QPagedPaintDevice__PageSize(113)
	QPagedPaintDevice__EnvelopePrc7       = QPagedPaintDevice__PageSize(114)
	QPagedPaintDevice__EnvelopePrc8       = QPagedPaintDevice__PageSize(115)
	QPagedPaintDevice__EnvelopePrc9       = QPagedPaintDevice__PageSize(116)
	QPagedPaintDevice__EnvelopePrc10      = QPagedPaintDevice__PageSize(117)
	QPagedPaintDevice__EnvelopeYou4       = QPagedPaintDevice__PageSize(118)
	QPagedPaintDevice__LastPageSize       = QPagedPaintDevice__PageSize(QPagedPaintDevice__EnvelopeYou4)
	QPagedPaintDevice__NPageSize          = QPagedPaintDevice__PageSize(QPagedPaintDevice__LastPageSize)
	QPagedPaintDevice__NPaperSize         = QPagedPaintDevice__PageSize(QPagedPaintDevice__LastPageSize)
	QPagedPaintDevice__AnsiA              = QPagedPaintDevice__PageSize(QPagedPaintDevice__Letter)
	QPagedPaintDevice__AnsiB              = QPagedPaintDevice__PageSize(QPagedPaintDevice__Ledger)
	QPagedPaintDevice__EnvelopeC5         = QPagedPaintDevice__PageSize(QPagedPaintDevice__C5E)
	QPagedPaintDevice__EnvelopeDL         = QPagedPaintDevice__PageSize(QPagedPaintDevice__DLE)
	QPagedPaintDevice__Envelope10         = QPagedPaintDevice__PageSize(QPagedPaintDevice__Comm10E)
)

func (ptr *QPagedPaintDevice) NewPage() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPagedPaintDevice::newPage")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QPagedPaintDevice_NewPage(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPagedPaintDevice) PageSize() QPagedPaintDevice__PageSize {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPagedPaintDevice::pageSize")
		}
	}()

	if ptr.Pointer() != nil {
		return QPagedPaintDevice__PageSize(C.QPagedPaintDevice_PageSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPagedPaintDevice) SetPageLayout(newPageLayout QPageLayout_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPagedPaintDevice::setPageLayout")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QPagedPaintDevice_SetPageLayout(ptr.Pointer(), PointerFromQPageLayout(newPageLayout)) != 0
	}
	return false
}

func (ptr *QPagedPaintDevice) SetPageMargins(margins core.QMarginsF_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPagedPaintDevice::setPageMargins")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QPagedPaintDevice_SetPageMargins(ptr.Pointer(), core.PointerFromQMarginsF(margins)) != 0
	}
	return false
}

func (ptr *QPagedPaintDevice) SetPageMargins2(margins core.QMarginsF_ITF, units QPageLayout__Unit) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPagedPaintDevice::setPageMargins")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QPagedPaintDevice_SetPageMargins2(ptr.Pointer(), core.PointerFromQMarginsF(margins), C.int(units)) != 0
	}
	return false
}

func (ptr *QPagedPaintDevice) SetPageOrientation(orientation QPageLayout__Orientation) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPagedPaintDevice::setPageOrientation")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QPagedPaintDevice_SetPageOrientation(ptr.Pointer(), C.int(orientation)) != 0
	}
	return false
}

func (ptr *QPagedPaintDevice) SetPageSize(pageSize QPageSize_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPagedPaintDevice::setPageSize")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QPagedPaintDevice_SetPageSize(ptr.Pointer(), PointerFromQPageSize(pageSize)) != 0
	}
	return false
}

func (ptr *QPagedPaintDevice) SetPageSize2(size QPagedPaintDevice__PageSize) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPagedPaintDevice::setPageSize")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPagedPaintDevice_SetPageSize2(ptr.Pointer(), C.int(size))
	}
}

func (ptr *QPagedPaintDevice) SetPageSizeMM(size core.QSizeF_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPagedPaintDevice::setPageSizeMM")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPagedPaintDevice_SetPageSizeMM(ptr.Pointer(), core.PointerFromQSizeF(size))
	}
}

func (ptr *QPagedPaintDevice) DestroyQPagedPaintDevice() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPagedPaintDevice::~QPagedPaintDevice")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPagedPaintDevice_DestroyQPagedPaintDevice(ptr.Pointer())
	}
}
