#include "qlayoutitem.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QRect>
#include <QLayout>
#include <QString>
#include <QLayoutItem>
#include "_cgo_export.h"

class MyQLayoutItem: public QLayoutItem {
public:
};

int QLayoutItem_Alignment(QtObjectPtr ptr){
	return static_cast<QLayoutItem*>(ptr)->alignment();
}

int QLayoutItem_ControlTypes(QtObjectPtr ptr){
	return static_cast<QLayoutItem*>(ptr)->controlTypes();
}

int QLayoutItem_ExpandingDirections(QtObjectPtr ptr){
	return static_cast<QLayoutItem*>(ptr)->expandingDirections();
}

int QLayoutItem_HasHeightForWidth(QtObjectPtr ptr){
	return static_cast<QLayoutItem*>(ptr)->hasHeightForWidth();
}

int QLayoutItem_HeightForWidth(QtObjectPtr ptr, int w){
	return static_cast<QLayoutItem*>(ptr)->heightForWidth(w);
}

void QLayoutItem_Invalidate(QtObjectPtr ptr){
	static_cast<QLayoutItem*>(ptr)->invalidate();
}

int QLayoutItem_IsEmpty(QtObjectPtr ptr){
	return static_cast<QLayoutItem*>(ptr)->isEmpty();
}

QtObjectPtr QLayoutItem_Layout(QtObjectPtr ptr){
	return static_cast<QLayoutItem*>(ptr)->layout();
}

int QLayoutItem_MinimumHeightForWidth(QtObjectPtr ptr, int w){
	return static_cast<QLayoutItem*>(ptr)->minimumHeightForWidth(w);
}

void QLayoutItem_SetAlignment(QtObjectPtr ptr, int alignment){
	static_cast<QLayoutItem*>(ptr)->setAlignment(static_cast<Qt::AlignmentFlag>(alignment));
}

void QLayoutItem_SetGeometry(QtObjectPtr ptr, QtObjectPtr r){
	static_cast<QLayoutItem*>(ptr)->setGeometry(*static_cast<QRect*>(r));
}

QtObjectPtr QLayoutItem_SpacerItem(QtObjectPtr ptr){
	return static_cast<QLayoutItem*>(ptr)->spacerItem();
}

QtObjectPtr QLayoutItem_Widget(QtObjectPtr ptr){
	return static_cast<QLayoutItem*>(ptr)->widget();
}

void QLayoutItem_DestroyQLayoutItem(QtObjectPtr ptr){
	static_cast<QLayoutItem*>(ptr)->~QLayoutItem();
}

