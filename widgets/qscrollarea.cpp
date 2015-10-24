#include "qscrollarea.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QWidget>
#include <QString>
#include <QScrollArea>
#include "_cgo_export.h"

class MyQScrollArea: public QScrollArea {
public:
};

int QScrollArea_Alignment(QtObjectPtr ptr){
	return static_cast<QScrollArea*>(ptr)->alignment();
}

void QScrollArea_SetAlignment(QtObjectPtr ptr, int v){
	static_cast<QScrollArea*>(ptr)->setAlignment(static_cast<Qt::AlignmentFlag>(v));
}

void QScrollArea_SetWidget(QtObjectPtr ptr, QtObjectPtr widget){
	static_cast<QScrollArea*>(ptr)->setWidget(static_cast<QWidget*>(widget));
}

void QScrollArea_SetWidgetResizable(QtObjectPtr ptr, int resizable){
	static_cast<QScrollArea*>(ptr)->setWidgetResizable(resizable != 0);
}

int QScrollArea_WidgetResizable(QtObjectPtr ptr){
	return static_cast<QScrollArea*>(ptr)->widgetResizable();
}

QtObjectPtr QScrollArea_NewQScrollArea(QtObjectPtr parent){
	return new QScrollArea(static_cast<QWidget*>(parent));
}

void QScrollArea_EnsureVisible(QtObjectPtr ptr, int x, int y, int xmargin, int ymargin){
	static_cast<QScrollArea*>(ptr)->ensureVisible(x, y, xmargin, ymargin);
}

void QScrollArea_EnsureWidgetVisible(QtObjectPtr ptr, QtObjectPtr childWidget, int xmargin, int ymargin){
	static_cast<QScrollArea*>(ptr)->ensureWidgetVisible(static_cast<QWidget*>(childWidget), xmargin, ymargin);
}

int QScrollArea_FocusNextPrevChild(QtObjectPtr ptr, int next){
	return static_cast<QScrollArea*>(ptr)->focusNextPrevChild(next != 0);
}

QtObjectPtr QScrollArea_TakeWidget(QtObjectPtr ptr){
	return static_cast<QScrollArea*>(ptr)->takeWidget();
}

QtObjectPtr QScrollArea_Widget(QtObjectPtr ptr){
	return static_cast<QScrollArea*>(ptr)->widget();
}

void QScrollArea_DestroyQScrollArea(QtObjectPtr ptr){
	static_cast<QScrollArea*>(ptr)->~QScrollArea();
}

