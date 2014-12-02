package qt

//#include "qmodelindex.h"
import "C"

type qmodelindex struct {
	ptr C.QtObjectPtr
}

type QModelIndex interface {
	Pointer() (ptr C.QtObjectPtr)
	SetPointer(ptr C.QtObjectPtr)
	Column() int
	Data_Int(role int) string
	Flags() ItemFlag
	IsValid() bool
	Row() int
}

func (p *qmodelindex) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qmodelindex) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

func (p *qmodelindex) Column() int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QModelIndex_Column(p.Pointer()))
	}
}

func (p *qmodelindex) Data_Int(role int) string {
	if p.Pointer() == nil {
		return ""
	} else {
		return C.GoString(C.QModelIndex_Data_Int(p.Pointer(), C.int(role)))
	}
}

func (p *qmodelindex) Flags() ItemFlag {
	if p.Pointer() == nil {
		return 0
	} else {
		return ItemFlag(C.QModelIndex_Flags(p.Pointer()))
	}
}

func (p *qmodelindex) IsValid() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QModelIndex_IsValid(p.Pointer()) != 0
	}
}

func (p *qmodelindex) Row() int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QModelIndex_Row(p.Pointer()))
	}
}
