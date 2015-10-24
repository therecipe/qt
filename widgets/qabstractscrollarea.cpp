#include "qabstractscrollarea.h"
#include <QModelIndex>
#include <QScrollBar>
#include <QWidget>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QAbstractScrollArea>
#include "_cgo_export.h"

class MyQAbstractScrollArea: public QAbstractScrollArea {
public:
};

int QAbstractScrollArea_HorizontalScrollBarPolicy(QtObjectPtr ptr){
	return static_cast<QAbstractScrollArea*>(ptr)->horizontalScrollBarPolicy();
}

void QAbstractScrollArea_SetHorizontalScrollBarPolicy(QtObjectPtr ptr, int v){
	static_cast<QAbstractScrollArea*>(ptr)->setHorizontalScrollBarPolicy(static_cast<Qt::ScrollBarPolicy>(v));
}

void QAbstractScrollArea_SetSizeAdjustPolicy(QtObjectPtr ptr, int policy){
	static_cast<QAbstractScrollArea*>(ptr)->setSizeAdjustPolicy(static_cast<QAbstractScrollArea::SizeAdjustPolicy>(policy));
}

void QAbstractScrollArea_SetVerticalScrollBarPolicy(QtObjectPtr ptr, int v){
	static_cast<QAbstractScrollArea*>(ptr)->setVerticalScrollBarPolicy(static_cast<Qt::ScrollBarPolicy>(v));
}

int QAbstractScrollArea_SizeAdjustPolicy(QtObjectPtr ptr){
	return static_cast<QAbstractScrollArea*>(ptr)->sizeAdjustPolicy();
}

int QAbstractScrollArea_VerticalScrollBarPolicy(QtObjectPtr ptr){
	return static_cast<QAbstractScrollArea*>(ptr)->verticalScrollBarPolicy();
}

QtObjectPtr QAbstractScrollArea_NewQAbstractScrollArea(QtObjectPtr parent){
	return new QAbstractScrollArea(static_cast<QWidget*>(parent));
}

void QAbstractScrollArea_AddScrollBarWidget(QtObjectPtr ptr, QtObjectPtr widget, int alignment){
	static_cast<QAbstractScrollArea*>(ptr)->addScrollBarWidget(static_cast<QWidget*>(widget), static_cast<Qt::AlignmentFlag>(alignment));
}

QtObjectPtr QAbstractScrollArea_CornerWidget(QtObjectPtr ptr){
	return static_cast<QAbstractScrollArea*>(ptr)->cornerWidget();
}

QtObjectPtr QAbstractScrollArea_HorizontalScrollBar(QtObjectPtr ptr){
	return static_cast<QAbstractScrollArea*>(ptr)->horizontalScrollBar();
}

void QAbstractScrollArea_SetCornerWidget(QtObjectPtr ptr, QtObjectPtr widget){
	static_cast<QAbstractScrollArea*>(ptr)->setCornerWidget(static_cast<QWidget*>(widget));
}

void QAbstractScrollArea_SetHorizontalScrollBar(QtObjectPtr ptr, QtObjectPtr scrollBar){
	static_cast<QAbstractScrollArea*>(ptr)->setHorizontalScrollBar(static_cast<QScrollBar*>(scrollBar));
}

void QAbstractScrollArea_SetVerticalScrollBar(QtObjectPtr ptr, QtObjectPtr scrollBar){
	static_cast<QAbstractScrollArea*>(ptr)->setVerticalScrollBar(static_cast<QScrollBar*>(scrollBar));
}

void QAbstractScrollArea_SetViewport(QtObjectPtr ptr, QtObjectPtr widget){
	static_cast<QAbstractScrollArea*>(ptr)->setViewport(static_cast<QWidget*>(widget));
}

void QAbstractScrollArea_SetupViewport(QtObjectPtr ptr, QtObjectPtr viewport){
	static_cast<QAbstractScrollArea*>(ptr)->setupViewport(static_cast<QWidget*>(viewport));
}

QtObjectPtr QAbstractScrollArea_VerticalScrollBar(QtObjectPtr ptr){
	return static_cast<QAbstractScrollArea*>(ptr)->verticalScrollBar();
}

QtObjectPtr QAbstractScrollArea_Viewport(QtObjectPtr ptr){
	return static_cast<QAbstractScrollArea*>(ptr)->viewport();
}

void QAbstractScrollArea_DestroyQAbstractScrollArea(QtObjectPtr ptr){
	static_cast<QAbstractScrollArea*>(ptr)->~QAbstractScrollArea();
}

