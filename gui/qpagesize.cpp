#include "qpagesize.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSizeF>
#include <QSize>
#include <QPageSize>
#include "_cgo_export.h"

class MyQPageSize: public QPageSize {
public:
};

QtObjectPtr QPageSize_NewQPageSize(){
	return new QPageSize();
}

QtObjectPtr QPageSize_NewQPageSize2(int pageSize){
	return new QPageSize(static_cast<QPageSize::PageSizeId>(pageSize));
}

QtObjectPtr QPageSize_NewQPageSize5(QtObjectPtr other){
	return new QPageSize(*static_cast<QPageSize*>(other));
}

QtObjectPtr QPageSize_NewQPageSize3(QtObjectPtr pointSize, char* name, int matchPolicy){
	return new QPageSize(*static_cast<QSize*>(pointSize), QString(name), static_cast<QPageSize::SizeMatchPolicy>(matchPolicy));
}

QtObjectPtr QPageSize_NewQPageSize4(QtObjectPtr size, int units, char* name, int matchPolicy){
	return new QPageSize(*static_cast<QSizeF*>(size), static_cast<QPageSize::Unit>(units), QString(name), static_cast<QPageSize::SizeMatchPolicy>(matchPolicy));
}

int QPageSize_QPageSize_DefinitionUnits2(int pageSizeId){
	return QPageSize::definitionUnits(static_cast<QPageSize::PageSizeId>(pageSizeId));
}

int QPageSize_DefinitionUnits(QtObjectPtr ptr){
	return static_cast<QPageSize*>(ptr)->definitionUnits();
}

int QPageSize_QPageSize_Id2(QtObjectPtr pointSize, int matchPolicy){
	return QPageSize::id(*static_cast<QSize*>(pointSize), static_cast<QPageSize::SizeMatchPolicy>(matchPolicy));
}

int QPageSize_QPageSize_Id3(QtObjectPtr size, int units, int matchPolicy){
	return QPageSize::id(*static_cast<QSizeF*>(size), static_cast<QPageSize::Unit>(units), static_cast<QPageSize::SizeMatchPolicy>(matchPolicy));
}

int QPageSize_QPageSize_Id4(int windowsId){
	return QPageSize::id(windowsId);
}

int QPageSize_Id(QtObjectPtr ptr){
	return static_cast<QPageSize*>(ptr)->id();
}

int QPageSize_IsEquivalentTo(QtObjectPtr ptr, QtObjectPtr other){
	return static_cast<QPageSize*>(ptr)->isEquivalentTo(*static_cast<QPageSize*>(other));
}

int QPageSize_IsValid(QtObjectPtr ptr){
	return static_cast<QPageSize*>(ptr)->isValid();
}

char* QPageSize_QPageSize_Key2(int pageSizeId){
	return QPageSize::key(static_cast<QPageSize::PageSizeId>(pageSizeId)).toUtf8().data();
}

char* QPageSize_Key(QtObjectPtr ptr){
	return static_cast<QPageSize*>(ptr)->key().toUtf8().data();
}

char* QPageSize_QPageSize_Name2(int pageSizeId){
	return QPageSize::name(static_cast<QPageSize::PageSizeId>(pageSizeId)).toUtf8().data();
}

char* QPageSize_Name(QtObjectPtr ptr){
	return static_cast<QPageSize*>(ptr)->name().toUtf8().data();
}

void QPageSize_Swap(QtObjectPtr ptr, QtObjectPtr other){
	static_cast<QPageSize*>(ptr)->swap(*static_cast<QPageSize*>(other));
}

int QPageSize_QPageSize_WindowsId2(int pageSizeId){
	return QPageSize::windowsId(static_cast<QPageSize::PageSizeId>(pageSizeId));
}

int QPageSize_WindowsId(QtObjectPtr ptr){
	return static_cast<QPageSize*>(ptr)->windowsId();
}

void QPageSize_DestroyQPageSize(QtObjectPtr ptr){
	static_cast<QPageSize*>(ptr)->~QPageSize();
}

