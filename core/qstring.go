package core

//#include "qstring.h"
import "C"
import (
	"unsafe"
)

type QString struct {
	ptr unsafe.Pointer
}

type QStringITF interface {
	QStringPTR() *QString
}

func (p *QString) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QString) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQString(ptr QStringITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStringPTR().Pointer()
	}
	return nil
}

func QStringFromPointer(ptr unsafe.Pointer) *QString {
	var n = new(QString)
	n.SetPointer(ptr)
	return n
}

func (ptr *QString) QStringPTR() *QString {
	return ptr
}

//QString::NormalizationForm
type QString__NormalizationForm int

var (
	QString__NormalizationForm_D  = QString__NormalizationForm(0)
	QString__NormalizationForm_C  = QString__NormalizationForm(1)
	QString__NormalizationForm_KD = QString__NormalizationForm(2)
	QString__NormalizationForm_KC = QString__NormalizationForm(3)
)

//QString::SectionFlag
type QString__SectionFlag int

var (
	QString__SectionDefault             = QString__SectionFlag(0x00)
	QString__SectionSkipEmpty           = QString__SectionFlag(0x01)
	QString__SectionIncludeLeadingSep   = QString__SectionFlag(0x02)
	QString__SectionIncludeTrailingSep  = QString__SectionFlag(0x04)
	QString__SectionCaseInsensitiveSeps = QString__SectionFlag(0x08)
)

//QString::SplitBehavior
type QString__SplitBehavior int

var (
	QString__KeepEmptyParts = QString__SplitBehavior(0)
	QString__SkipEmptyParts = QString__SplitBehavior(1)
)
