#include "qscrollarea.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QWidget>
#include <QScrollArea>
#include "_cgo_export.h"

class MyQScrollArea: public QScrollArea {
public:
};

int QScrollArea_Alignment(void* ptr){
	return static_cast<QScrollArea*>(ptr)->alignment();
}

void QScrollArea_SetAlignment(void* ptr, int v){
	static_cast<QScrollArea*>(ptr)->setAlignment(static_cast<Qt::AlignmentFlag>(v));
}

void QScrollArea_SetWidget(void* ptr, void* widget){
	static_cast<QScrollArea*>(ptr)->setWidget(static_cast<QWidget*>(widget));
}

void QScrollArea_SetWidgetResizable(void* ptr, int resizable){
	static_cast<QScrollArea*>(ptr)->setWidgetResizable(resizable != 0);
}

int QScrollArea_WidgetResizable(void* ptr){
	return static_cast<QScrollArea*>(ptr)->widgetResizable();
}

void* QScrollArea_NewQScrollArea(void* parent){
	return new QScrollArea(static_cast<QWidget*>(parent));
}

void QScrollArea_EnsureVisible(void* ptr, int x, int y, int xmargin, int ymargin){
	static_cast<QScrollArea*>(ptr)->ensureVisible(x, y, xmargin, ymargin);
}

void QScrollArea_EnsureWidgetVisible(void* ptr, void* childWidget, int xmargin, int ymargin){
	static_cast<QScrollArea*>(ptr)->ensureWidgetVisible(static_cast<QWidget*>(childWidget), xmargin, ymargin);
}

int QScrollArea_FocusNextPrevChild(void* ptr, int next){
	return static_cast<QScrollArea*>(ptr)->focusNextPrevChild(next != 0);
}

void* QScrollArea_TakeWidget(void* ptr){
	return static_cast<QScrollArea*>(ptr)->takeWidget();
}

void* QScrollArea_Widget(void* ptr){
	return static_cast<QScrollArea*>(ptr)->widget();
}

void QScrollArea_DestroyQScrollArea(void* ptr){
	static_cast<QScrollArea*>(ptr)->~QScrollArea();
}

