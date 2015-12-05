package core

//#include "core.h"
import "C"
import (
	"log"
	"unsafe"
)

type QSysInfo struct {
	ptr unsafe.Pointer
}

type QSysInfo_ITF interface {
	QSysInfo_PTR() *QSysInfo
}

func (p *QSysInfo) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSysInfo) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSysInfo(ptr QSysInfo_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSysInfo_PTR().Pointer()
	}
	return nil
}

func NewQSysInfoFromPointer(ptr unsafe.Pointer) *QSysInfo {
	var n = new(QSysInfo)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSysInfo) QSysInfo_PTR() *QSysInfo {
	return ptr
}

//QSysInfo::Endian
type QSysInfo__Endian int64

const (
	QSysInfo__BigEndian    = QSysInfo__Endian(0)
	QSysInfo__LittleEndian = QSysInfo__Endian(1)
)

//QSysInfo::MacVersion
type QSysInfo__MacVersion int64

var (
	QSysInfo__MV_None         = QSysInfo__MacVersion(0xffff)
	QSysInfo__MV_Unknown      = QSysInfo__MacVersion(0x0000)
	QSysInfo__MV_9            = QSysInfo__MacVersion(0x0001)
	QSysInfo__MV_10_0         = QSysInfo__MacVersion(0x0002)
	QSysInfo__MV_10_1         = QSysInfo__MacVersion(0x0003)
	QSysInfo__MV_10_2         = QSysInfo__MacVersion(0x0004)
	QSysInfo__MV_10_3         = QSysInfo__MacVersion(0x0005)
	QSysInfo__MV_10_4         = QSysInfo__MacVersion(0x0006)
	QSysInfo__MV_10_5         = QSysInfo__MacVersion(0x0007)
	QSysInfo__MV_10_6         = QSysInfo__MacVersion(0x0008)
	QSysInfo__MV_10_7         = QSysInfo__MacVersion(0x0009)
	QSysInfo__MV_10_8         = QSysInfo__MacVersion(0x000A)
	QSysInfo__MV_10_9         = QSysInfo__MacVersion(0x000B)
	QSysInfo__MV_10_10        = QSysInfo__MacVersion(0x000C)
	QSysInfo__MV_10_11        = QSysInfo__MacVersion(0x000D)
	QSysInfo__MV_CHEETAH      = QSysInfo__MacVersion(QSysInfo__MV_10_0)
	QSysInfo__MV_PUMA         = QSysInfo__MacVersion(QSysInfo__MV_10_1)
	QSysInfo__MV_JAGUAR       = QSysInfo__MacVersion(QSysInfo__MV_10_2)
	QSysInfo__MV_PANTHER      = QSysInfo__MacVersion(QSysInfo__MV_10_3)
	QSysInfo__MV_TIGER        = QSysInfo__MacVersion(QSysInfo__MV_10_4)
	QSysInfo__MV_LEOPARD      = QSysInfo__MacVersion(QSysInfo__MV_10_5)
	QSysInfo__MV_SNOWLEOPARD  = QSysInfo__MacVersion(QSysInfo__MV_10_6)
	QSysInfo__MV_LION         = QSysInfo__MacVersion(QSysInfo__MV_10_7)
	QSysInfo__MV_MOUNTAINLION = QSysInfo__MacVersion(QSysInfo__MV_10_8)
	QSysInfo__MV_MAVERICKS    = QSysInfo__MacVersion(QSysInfo__MV_10_9)
	QSysInfo__MV_YOSEMITE     = QSysInfo__MacVersion(QSysInfo__MV_10_10)
	QSysInfo__MV_ELCAPITAN    = QSysInfo__MacVersion(QSysInfo__MV_10_11)
	QSysInfo__MV_IOS          = QSysInfo__MacVersion(C.QSysInfo_MV_IOS_Type())
	QSysInfo__MV_IOS_4_3      = QSysInfo__MacVersion(C.QSysInfo_MV_IOS_4_3_Type())
	QSysInfo__MV_IOS_5_0      = QSysInfo__MacVersion(C.QSysInfo_MV_IOS_5_0_Type())
	QSysInfo__MV_IOS_5_1      = QSysInfo__MacVersion(C.QSysInfo_MV_IOS_5_1_Type())
	QSysInfo__MV_IOS_6_0      = QSysInfo__MacVersion(C.QSysInfo_MV_IOS_6_0_Type())
	QSysInfo__MV_IOS_6_1      = QSysInfo__MacVersion(C.QSysInfo_MV_IOS_6_1_Type())
	QSysInfo__MV_IOS_7_0      = QSysInfo__MacVersion(C.QSysInfo_MV_IOS_7_0_Type())
	QSysInfo__MV_IOS_7_1      = QSysInfo__MacVersion(C.QSysInfo_MV_IOS_7_1_Type())
	QSysInfo__MV_IOS_8_0      = QSysInfo__MacVersion(C.QSysInfo_MV_IOS_8_0_Type())
	QSysInfo__MV_IOS_8_1      = QSysInfo__MacVersion(C.QSysInfo_MV_IOS_8_1_Type())
	QSysInfo__MV_IOS_8_2      = QSysInfo__MacVersion(C.QSysInfo_MV_IOS_8_2_Type())
	QSysInfo__MV_IOS_8_3      = QSysInfo__MacVersion(C.QSysInfo_MV_IOS_8_3_Type())
	QSysInfo__MV_IOS_8_4      = QSysInfo__MacVersion(C.QSysInfo_MV_IOS_8_4_Type())
	QSysInfo__MV_IOS_9_0      = QSysInfo__MacVersion(C.QSysInfo_MV_IOS_9_0_Type())
)

//QSysInfo::Sizes
type QSysInfo__Sizes int64

var (
	QSysInfo__WordSize = QSysInfo__Sizes(C.QSysInfo_WordSize_Type())
)

//QSysInfo::WinVersion
type QSysInfo__WinVersion int64

const (
	QSysInfo__WV_None       = QSysInfo__WinVersion(0x0000)
	QSysInfo__WV_32s        = QSysInfo__WinVersion(0x0001)
	QSysInfo__WV_95         = QSysInfo__WinVersion(0x0002)
	QSysInfo__WV_98         = QSysInfo__WinVersion(0x0003)
	QSysInfo__WV_Me         = QSysInfo__WinVersion(0x0004)
	QSysInfo__WV_DOS_based  = QSysInfo__WinVersion(0x000f)
	QSysInfo__WV_NT         = QSysInfo__WinVersion(0x0010)
	QSysInfo__WV_2000       = QSysInfo__WinVersion(0x0020)
	QSysInfo__WV_XP         = QSysInfo__WinVersion(0x0030)
	QSysInfo__WV_2003       = QSysInfo__WinVersion(0x0040)
	QSysInfo__WV_VISTA      = QSysInfo__WinVersion(0x0080)
	QSysInfo__WV_WINDOWS7   = QSysInfo__WinVersion(0x0090)
	QSysInfo__WV_WINDOWS8   = QSysInfo__WinVersion(0x00a0)
	QSysInfo__WV_WINDOWS8_1 = QSysInfo__WinVersion(0x00b0)
	QSysInfo__WV_WINDOWS10  = QSysInfo__WinVersion(0x00c0)
	QSysInfo__WV_NT_based   = QSysInfo__WinVersion(0x00f0)
	QSysInfo__WV_4_0        = QSysInfo__WinVersion(QSysInfo__WV_NT)
	QSysInfo__WV_5_0        = QSysInfo__WinVersion(QSysInfo__WV_2000)
	QSysInfo__WV_5_1        = QSysInfo__WinVersion(QSysInfo__WV_XP)
	QSysInfo__WV_5_2        = QSysInfo__WinVersion(QSysInfo__WV_2003)
	QSysInfo__WV_6_0        = QSysInfo__WinVersion(QSysInfo__WV_VISTA)
	QSysInfo__WV_6_1        = QSysInfo__WinVersion(QSysInfo__WV_WINDOWS7)
	QSysInfo__WV_6_2        = QSysInfo__WinVersion(QSysInfo__WV_WINDOWS8)
	QSysInfo__WV_6_3        = QSysInfo__WinVersion(QSysInfo__WV_WINDOWS8_1)
	QSysInfo__WV_10_0       = QSysInfo__WinVersion(QSysInfo__WV_WINDOWS10)
	QSysInfo__WV_CE         = QSysInfo__WinVersion(0x0100)
	QSysInfo__WV_CENET      = QSysInfo__WinVersion(0x0200)
	QSysInfo__WV_CE_5       = QSysInfo__WinVersion(0x0300)
	QSysInfo__WV_CE_6       = QSysInfo__WinVersion(0x0400)
	QSysInfo__WV_CE_based   = QSysInfo__WinVersion(0x0f00)
)

func QSysInfo_MacVersion() QSysInfo__MacVersion {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSysInfo::macVersion")
		}
	}()

	return QSysInfo__MacVersion(C.QSysInfo_QSysInfo_MacVersion())
}

func QSysInfo_BuildAbi() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSysInfo::buildAbi")
		}
	}()

	return C.GoString(C.QSysInfo_QSysInfo_BuildAbi())
}

func QSysInfo_BuildCpuArchitecture() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSysInfo::buildCpuArchitecture")
		}
	}()

	return C.GoString(C.QSysInfo_QSysInfo_BuildCpuArchitecture())
}

func QSysInfo_CurrentCpuArchitecture() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSysInfo::currentCpuArchitecture")
		}
	}()

	return C.GoString(C.QSysInfo_QSysInfo_CurrentCpuArchitecture())
}

func QSysInfo_KernelType() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSysInfo::kernelType")
		}
	}()

	return C.GoString(C.QSysInfo_QSysInfo_KernelType())
}

func QSysInfo_KernelVersion() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSysInfo::kernelVersion")
		}
	}()

	return C.GoString(C.QSysInfo_QSysInfo_KernelVersion())
}

func QSysInfo_PrettyProductName() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSysInfo::prettyProductName")
		}
	}()

	return C.GoString(C.QSysInfo_QSysInfo_PrettyProductName())
}

func QSysInfo_ProductType() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSysInfo::productType")
		}
	}()

	return C.GoString(C.QSysInfo_QSysInfo_ProductType())
}

func QSysInfo_ProductVersion() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSysInfo::productVersion")
		}
	}()

	return C.GoString(C.QSysInfo_QSysInfo_ProductVersion())
}

func QSysInfo_WindowsVersion() QSysInfo__WinVersion {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSysInfo::windowsVersion")
		}
	}()

	return QSysInfo__WinVersion(C.QSysInfo_QSysInfo_WindowsVersion())
}
