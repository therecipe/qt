#include "qpagesize.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSize>
#include <QSizeF>
#include <QString>
#include <QPageSize>
#include "_cgo_export.h"

class MyQPageSize: public QPageSize {
public:
};

void* QPageSize_NewQPageSize(){
	return new QPageSize();
}

void* QPageSize_NewQPageSize2(int pageSize){
	return new QPageSize(static_cast<QPageSize::PageSizeId>(pageSize));
}

void* QPageSize_NewQPageSize5(void* other){
	return new QPageSize(*static_cast<QPageSize*>(other));
}

void* QPageSize_NewQPageSize3(void* pointSize, char* name, int matchPolicy){
	return new QPageSize(*static_cast<QSize*>(pointSize), QString(name), static_cast<QPageSize::SizeMatchPolicy>(matchPolicy));
}

void* QPageSize_NewQPageSize4(void* size, int units, char* name, int matchPolicy){
	return new QPageSize(*static_cast<QSizeF*>(size), static_cast<QPageSize::Unit>(units), QString(name), static_cast<QPageSize::SizeMatchPolicy>(matchPolicy));
}

int QPageSize_QPageSize_DefinitionUnits2(int pageSizeId){
	return QPageSize::definitionUnits(static_cast<QPageSize::PageSizeId>(pageSizeId));
}

int QPageSize_DefinitionUnits(void* ptr){
	return static_cast<QPageSize*>(ptr)->definitionUnits();
}

int QPageSize_QPageSize_Id2(void* pointSize, int matchPolicy){
	return QPageSize::id(*static_cast<QSize*>(pointSize), static_cast<QPageSize::SizeMatchPolicy>(matchPolicy));
}

int QPageSize_QPageSize_Id3(void* size, int units, int matchPolicy){
	return QPageSize::id(*static_cast<QSizeF*>(size), static_cast<QPageSize::Unit>(units), static_cast<QPageSize::SizeMatchPolicy>(matchPolicy));
}

int QPageSize_QPageSize_Id4(int windowsId){
	return QPageSize::id(windowsId);
}

int QPageSize_Id(void* ptr){
	return static_cast<QPageSize*>(ptr)->id();
}

int QPageSize_IsEquivalentTo(void* ptr, void* other){
	return static_cast<QPageSize*>(ptr)->isEquivalentTo(*static_cast<QPageSize*>(other));
}

int QPageSize_IsValid(void* ptr){
	return static_cast<QPageSize*>(ptr)->isValid();
}

char* QPageSize_QPageSize_Key2(int pageSizeId){
	return QPageSize::key(static_cast<QPageSize::PageSizeId>(pageSizeId)).toUtf8().data();
}

char* QPageSize_Key(void* ptr){
	return static_cast<QPageSize*>(ptr)->key().toUtf8().data();
}

char* QPageSize_QPageSize_Name2(int pageSizeId){
	return QPageSize::name(static_cast<QPageSize::PageSizeId>(pageSizeId)).toUtf8().data();
}

char* QPageSize_Name(void* ptr){
	return static_cast<QPageSize*>(ptr)->name().toUtf8().data();
}

void QPageSize_Swap(void* ptr, void* other){
	static_cast<QPageSize*>(ptr)->swap(*static_cast<QPageSize*>(other));
}

int QPageSize_QPageSize_WindowsId2(int pageSizeId){
	return QPageSize::windowsId(static_cast<QPageSize::PageSizeId>(pageSizeId));
}

int QPageSize_WindowsId(void* ptr){
	return static_cast<QPageSize*>(ptr)->windowsId();
}

void QPageSize_DestroyQPageSize(void* ptr){
	static_cast<QPageSize*>(ptr)->~QPageSize();
}

