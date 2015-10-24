#include "qscrollbar.h"
#include <QWidget>
#include <QEvent>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QScrollBar>
#include "_cgo_export.h"

class MyQScrollBar: public QScrollBar {
public:
};

QtObjectPtr QScrollBar_NewQScrollBar(QtObjectPtr parent){
	return new QScrollBar(static_cast<QWidget*>(parent));
}

QtObjectPtr QScrollBar_NewQScrollBar2(int orientation, QtObjectPtr parent){
	return new QScrollBar(static_cast<Qt::Orientation>(orientation), static_cast<QWidget*>(parent));
}

int QScrollBar_Event(QtObjectPtr ptr, QtObjectPtr event){
	return static_cast<QScrollBar*>(ptr)->event(static_cast<QEvent*>(event));
}

void QScrollBar_DestroyQScrollBar(QtObjectPtr ptr){
	static_cast<QScrollBar*>(ptr)->~QScrollBar();
}

