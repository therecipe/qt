#include "qlayoutitem.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QRect>
#include <QLayout>
#include <QLayoutItem>
#include "_cgo_export.h"

class MyQLayoutItem: public QLayoutItem {
public:
};

int QLayoutItem_Alignment(void* ptr){
	return static_cast<QLayoutItem*>(ptr)->alignment();
}

int QLayoutItem_ControlTypes(void* ptr){
	return static_cast<QLayoutItem*>(ptr)->controlTypes();
}

int QLayoutItem_ExpandingDirections(void* ptr){
	return static_cast<QLayoutItem*>(ptr)->expandingDirections();
}

int QLayoutItem_HasHeightForWidth(void* ptr){
	return static_cast<QLayoutItem*>(ptr)->hasHeightForWidth();
}

int QLayoutItem_HeightForWidth(void* ptr, int w){
	return static_cast<QLayoutItem*>(ptr)->heightForWidth(w);
}

void QLayoutItem_Invalidate(void* ptr){
	static_cast<QLayoutItem*>(ptr)->invalidate();
}

int QLayoutItem_IsEmpty(void* ptr){
	return static_cast<QLayoutItem*>(ptr)->isEmpty();
}

void* QLayoutItem_Layout(void* ptr){
	return static_cast<QLayoutItem*>(ptr)->layout();
}

int QLayoutItem_MinimumHeightForWidth(void* ptr, int w){
	return static_cast<QLayoutItem*>(ptr)->minimumHeightForWidth(w);
}

void QLayoutItem_SetAlignment(void* ptr, int alignment){
	static_cast<QLayoutItem*>(ptr)->setAlignment(static_cast<Qt::AlignmentFlag>(alignment));
}

void QLayoutItem_SetGeometry(void* ptr, void* r){
	static_cast<QLayoutItem*>(ptr)->setGeometry(*static_cast<QRect*>(r));
}

void* QLayoutItem_SpacerItem(void* ptr){
	return static_cast<QLayoutItem*>(ptr)->spacerItem();
}

void* QLayoutItem_Widget(void* ptr){
	return static_cast<QLayoutItem*>(ptr)->widget();
}

void QLayoutItem_DestroyQLayoutItem(void* ptr){
	static_cast<QLayoutItem*>(ptr)->~QLayoutItem();
}

