#include "qabstractscrollarea.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QWidget>
#include <QScrollBar>
#include <QAbstractScrollArea>
#include "_cgo_export.h"

class MyQAbstractScrollArea: public QAbstractScrollArea {
public:
};

int QAbstractScrollArea_HorizontalScrollBarPolicy(void* ptr){
	return static_cast<QAbstractScrollArea*>(ptr)->horizontalScrollBarPolicy();
}

void QAbstractScrollArea_SetHorizontalScrollBarPolicy(void* ptr, int v){
	static_cast<QAbstractScrollArea*>(ptr)->setHorizontalScrollBarPolicy(static_cast<Qt::ScrollBarPolicy>(v));
}

void QAbstractScrollArea_SetSizeAdjustPolicy(void* ptr, int policy){
	static_cast<QAbstractScrollArea*>(ptr)->setSizeAdjustPolicy(static_cast<QAbstractScrollArea::SizeAdjustPolicy>(policy));
}

void QAbstractScrollArea_SetVerticalScrollBarPolicy(void* ptr, int v){
	static_cast<QAbstractScrollArea*>(ptr)->setVerticalScrollBarPolicy(static_cast<Qt::ScrollBarPolicy>(v));
}

int QAbstractScrollArea_SizeAdjustPolicy(void* ptr){
	return static_cast<QAbstractScrollArea*>(ptr)->sizeAdjustPolicy();
}

int QAbstractScrollArea_VerticalScrollBarPolicy(void* ptr){
	return static_cast<QAbstractScrollArea*>(ptr)->verticalScrollBarPolicy();
}

void* QAbstractScrollArea_NewQAbstractScrollArea(void* parent){
	return new QAbstractScrollArea(static_cast<QWidget*>(parent));
}

void QAbstractScrollArea_AddScrollBarWidget(void* ptr, void* widget, int alignment){
	static_cast<QAbstractScrollArea*>(ptr)->addScrollBarWidget(static_cast<QWidget*>(widget), static_cast<Qt::AlignmentFlag>(alignment));
}

void* QAbstractScrollArea_CornerWidget(void* ptr){
	return static_cast<QAbstractScrollArea*>(ptr)->cornerWidget();
}

void* QAbstractScrollArea_HorizontalScrollBar(void* ptr){
	return static_cast<QAbstractScrollArea*>(ptr)->horizontalScrollBar();
}

void QAbstractScrollArea_SetCornerWidget(void* ptr, void* widget){
	static_cast<QAbstractScrollArea*>(ptr)->setCornerWidget(static_cast<QWidget*>(widget));
}

void QAbstractScrollArea_SetHorizontalScrollBar(void* ptr, void* scrollBar){
	static_cast<QAbstractScrollArea*>(ptr)->setHorizontalScrollBar(static_cast<QScrollBar*>(scrollBar));
}

void QAbstractScrollArea_SetVerticalScrollBar(void* ptr, void* scrollBar){
	static_cast<QAbstractScrollArea*>(ptr)->setVerticalScrollBar(static_cast<QScrollBar*>(scrollBar));
}

void QAbstractScrollArea_SetViewport(void* ptr, void* widget){
	static_cast<QAbstractScrollArea*>(ptr)->setViewport(static_cast<QWidget*>(widget));
}

void QAbstractScrollArea_SetupViewport(void* ptr, void* viewport){
	static_cast<QAbstractScrollArea*>(ptr)->setupViewport(static_cast<QWidget*>(viewport));
}

void* QAbstractScrollArea_VerticalScrollBar(void* ptr){
	return static_cast<QAbstractScrollArea*>(ptr)->verticalScrollBar();
}

void* QAbstractScrollArea_Viewport(void* ptr){
	return static_cast<QAbstractScrollArea*>(ptr)->viewport();
}

void QAbstractScrollArea_DestroyQAbstractScrollArea(void* ptr){
	static_cast<QAbstractScrollArea*>(ptr)->~QAbstractScrollArea();
}

