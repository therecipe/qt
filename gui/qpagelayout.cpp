#include "qpagelayout.h"
#include <QMarginsF>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QPageSize>
#include <QMargins>
#include <QPageLayout>
#include "_cgo_export.h"

class MyQPageLayout: public QPageLayout {
public:
};

void* QPageLayout_NewQPageLayout(){
	return new QPageLayout();
}

void* QPageLayout_NewQPageLayout3(void* other){
	return new QPageLayout(*static_cast<QPageLayout*>(other));
}

void* QPageLayout_NewQPageLayout2(void* pageSize, int orientation, void* margins, int units, void* minMargins){
	return new QPageLayout(*static_cast<QPageSize*>(pageSize), static_cast<QPageLayout::Orientation>(orientation), *static_cast<QMarginsF*>(margins), static_cast<QPageLayout::Unit>(units), *static_cast<QMarginsF*>(minMargins));
}

int QPageLayout_IsEquivalentTo(void* ptr, void* other){
	return static_cast<QPageLayout*>(ptr)->isEquivalentTo(*static_cast<QPageLayout*>(other));
}

int QPageLayout_IsValid(void* ptr){
	return static_cast<QPageLayout*>(ptr)->isValid();
}

int QPageLayout_Mode(void* ptr){
	return static_cast<QPageLayout*>(ptr)->mode();
}

int QPageLayout_Orientation(void* ptr){
	return static_cast<QPageLayout*>(ptr)->orientation();
}

int QPageLayout_SetBottomMargin(void* ptr, double bottomMargin){
	return static_cast<QPageLayout*>(ptr)->setBottomMargin(static_cast<qreal>(bottomMargin));
}

int QPageLayout_SetLeftMargin(void* ptr, double leftMargin){
	return static_cast<QPageLayout*>(ptr)->setLeftMargin(static_cast<qreal>(leftMargin));
}

int QPageLayout_SetMargins(void* ptr, void* margins){
	return static_cast<QPageLayout*>(ptr)->setMargins(*static_cast<QMarginsF*>(margins));
}

void QPageLayout_SetMinimumMargins(void* ptr, void* minMargins){
	static_cast<QPageLayout*>(ptr)->setMinimumMargins(*static_cast<QMarginsF*>(minMargins));
}

void QPageLayout_SetMode(void* ptr, int mode){
	static_cast<QPageLayout*>(ptr)->setMode(static_cast<QPageLayout::Mode>(mode));
}

void QPageLayout_SetOrientation(void* ptr, int orientation){
	static_cast<QPageLayout*>(ptr)->setOrientation(static_cast<QPageLayout::Orientation>(orientation));
}

void QPageLayout_SetPageSize(void* ptr, void* pageSize, void* minMargins){
	static_cast<QPageLayout*>(ptr)->setPageSize(*static_cast<QPageSize*>(pageSize), *static_cast<QMarginsF*>(minMargins));
}

int QPageLayout_SetRightMargin(void* ptr, double rightMargin){
	return static_cast<QPageLayout*>(ptr)->setRightMargin(static_cast<qreal>(rightMargin));
}

int QPageLayout_SetTopMargin(void* ptr, double topMargin){
	return static_cast<QPageLayout*>(ptr)->setTopMargin(static_cast<qreal>(topMargin));
}

void QPageLayout_SetUnits(void* ptr, int units){
	static_cast<QPageLayout*>(ptr)->setUnits(static_cast<QPageLayout::Unit>(units));
}

void QPageLayout_Swap(void* ptr, void* other){
	static_cast<QPageLayout*>(ptr)->swap(*static_cast<QPageLayout*>(other));
}

int QPageLayout_Units(void* ptr){
	return static_cast<QPageLayout*>(ptr)->units();
}

void QPageLayout_DestroyQPageLayout(void* ptr){
	static_cast<QPageLayout*>(ptr)->~QPageLayout();
}

