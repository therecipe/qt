#include "qscrollbar.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QEvent>
#include <QWidget>
#include <QScrollBar>
#include "_cgo_export.h"

class MyQScrollBar: public QScrollBar {
public:
};

void* QScrollBar_NewQScrollBar(void* parent){
	return new QScrollBar(static_cast<QWidget*>(parent));
}

void* QScrollBar_NewQScrollBar2(int orientation, void* parent){
	return new QScrollBar(static_cast<Qt::Orientation>(orientation), static_cast<QWidget*>(parent));
}

int QScrollBar_Event(void* ptr, void* event){
	return static_cast<QScrollBar*>(ptr)->event(static_cast<QEvent*>(event));
}

void QScrollBar_DestroyQScrollBar(void* ptr){
	static_cast<QScrollBar*>(ptr)->~QScrollBar();
}

