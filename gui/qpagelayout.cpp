#include "qpagelayout.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMargins>
#include <QMarginsF>
#include <QPageSize>
#include <QString>
#include <QPageLayout>
#include "_cgo_export.h"

class MyQPageLayout: public QPageLayout {
public:
};

QtObjectPtr QPageLayout_NewQPageLayout(){
	return new QPageLayout();
}

QtObjectPtr QPageLayout_NewQPageLayout3(QtObjectPtr other){
	return new QPageLayout(*static_cast<QPageLayout*>(other));
}

QtObjectPtr QPageLayout_NewQPageLayout2(QtObjectPtr pageSize, int orientation, QtObjectPtr margins, int units, QtObjectPtr minMargins){
	return new QPageLayout(*static_cast<QPageSize*>(pageSize), static_cast<QPageLayout::Orientation>(orientation), *static_cast<QMarginsF*>(margins), static_cast<QPageLayout::Unit>(units), *static_cast<QMarginsF*>(minMargins));
}

int QPageLayout_IsEquivalentTo(QtObjectPtr ptr, QtObjectPtr other){
	return static_cast<QPageLayout*>(ptr)->isEquivalentTo(*static_cast<QPageLayout*>(other));
}

int QPageLayout_IsValid(QtObjectPtr ptr){
	return static_cast<QPageLayout*>(ptr)->isValid();
}

int QPageLayout_Mode(QtObjectPtr ptr){
	return static_cast<QPageLayout*>(ptr)->mode();
}

int QPageLayout_Orientation(QtObjectPtr ptr){
	return static_cast<QPageLayout*>(ptr)->orientation();
}

int QPageLayout_SetMargins(QtObjectPtr ptr, QtObjectPtr margins){
	return static_cast<QPageLayout*>(ptr)->setMargins(*static_cast<QMarginsF*>(margins));
}

void QPageLayout_SetMinimumMargins(QtObjectPtr ptr, QtObjectPtr minMargins){
	static_cast<QPageLayout*>(ptr)->setMinimumMargins(*static_cast<QMarginsF*>(minMargins));
}

void QPageLayout_SetMode(QtObjectPtr ptr, int mode){
	static_cast<QPageLayout*>(ptr)->setMode(static_cast<QPageLayout::Mode>(mode));
}

void QPageLayout_SetOrientation(QtObjectPtr ptr, int orientation){
	static_cast<QPageLayout*>(ptr)->setOrientation(static_cast<QPageLayout::Orientation>(orientation));
}

void QPageLayout_SetPageSize(QtObjectPtr ptr, QtObjectPtr pageSize, QtObjectPtr minMargins){
	static_cast<QPageLayout*>(ptr)->setPageSize(*static_cast<QPageSize*>(pageSize), *static_cast<QMarginsF*>(minMargins));
}

void QPageLayout_SetUnits(QtObjectPtr ptr, int units){
	static_cast<QPageLayout*>(ptr)->setUnits(static_cast<QPageLayout::Unit>(units));
}

void QPageLayout_Swap(QtObjectPtr ptr, QtObjectPtr other){
	static_cast<QPageLayout*>(ptr)->swap(*static_cast<QPageLayout*>(other));
}

int QPageLayout_Units(QtObjectPtr ptr){
	return static_cast<QPageLayout*>(ptr)->units();
}

void QPageLayout_DestroyQPageLayout(QtObjectPtr ptr){
	static_cast<QPageLayout*>(ptr)->~QPageLayout();
}

