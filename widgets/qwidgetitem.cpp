#include "qwidgetitem.h"
#include <QModelIndex>
#include <QRect>
#include <QWidget>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QWidgetItem>
#include "_cgo_export.h"

class MyQWidgetItem: public QWidgetItem {
public:
};

void* QWidgetItem_NewQWidgetItem(void* widget){
	return new QWidgetItem(static_cast<QWidget*>(widget));
}

int QWidgetItem_ControlTypes(void* ptr){
	return static_cast<QWidgetItem*>(ptr)->controlTypes();
}

int QWidgetItem_ExpandingDirections(void* ptr){
	return static_cast<QWidgetItem*>(ptr)->expandingDirections();
}

int QWidgetItem_HasHeightForWidth(void* ptr){
	return static_cast<QWidgetItem*>(ptr)->hasHeightForWidth();
}

int QWidgetItem_HeightForWidth(void* ptr, int w){
	return static_cast<QWidgetItem*>(ptr)->heightForWidth(w);
}

int QWidgetItem_IsEmpty(void* ptr){
	return static_cast<QWidgetItem*>(ptr)->isEmpty();
}

void QWidgetItem_SetGeometry(void* ptr, void* rect){
	static_cast<QWidgetItem*>(ptr)->setGeometry(*static_cast<QRect*>(rect));
}

void* QWidgetItem_Widget(void* ptr){
	return static_cast<QWidgetItem*>(ptr)->widget();
}

void QWidgetItem_DestroyQWidgetItem(void* ptr){
	static_cast<QWidgetItem*>(ptr)->~QWidgetItem();
}

