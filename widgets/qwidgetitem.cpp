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

QtObjectPtr QWidgetItem_NewQWidgetItem(QtObjectPtr widget){
	return new QWidgetItem(static_cast<QWidget*>(widget));
}

int QWidgetItem_ControlTypes(QtObjectPtr ptr){
	return static_cast<QWidgetItem*>(ptr)->controlTypes();
}

int QWidgetItem_ExpandingDirections(QtObjectPtr ptr){
	return static_cast<QWidgetItem*>(ptr)->expandingDirections();
}

int QWidgetItem_HasHeightForWidth(QtObjectPtr ptr){
	return static_cast<QWidgetItem*>(ptr)->hasHeightForWidth();
}

int QWidgetItem_HeightForWidth(QtObjectPtr ptr, int w){
	return static_cast<QWidgetItem*>(ptr)->heightForWidth(w);
}

int QWidgetItem_IsEmpty(QtObjectPtr ptr){
	return static_cast<QWidgetItem*>(ptr)->isEmpty();
}

void QWidgetItem_SetGeometry(QtObjectPtr ptr, QtObjectPtr rect){
	static_cast<QWidgetItem*>(ptr)->setGeometry(*static_cast<QRect*>(rect));
}

QtObjectPtr QWidgetItem_Widget(QtObjectPtr ptr){
	return static_cast<QWidgetItem*>(ptr)->widget();
}

void QWidgetItem_DestroyQWidgetItem(QtObjectPtr ptr){
	static_cast<QWidgetItem*>(ptr)->~QWidgetItem();
}

