#include "qfocusframe.h"
#include <QWidget>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QFocusFrame>
#include "_cgo_export.h"

class MyQFocusFrame: public QFocusFrame {
public:
};

QtObjectPtr QFocusFrame_NewQFocusFrame(QtObjectPtr parent){
	return new QFocusFrame(static_cast<QWidget*>(parent));
}

void QFocusFrame_SetWidget(QtObjectPtr ptr, QtObjectPtr widget){
	static_cast<QFocusFrame*>(ptr)->setWidget(static_cast<QWidget*>(widget));
}

QtObjectPtr QFocusFrame_Widget(QtObjectPtr ptr){
	return static_cast<QFocusFrame*>(ptr)->widget();
}

void QFocusFrame_DestroyQFocusFrame(QtObjectPtr ptr){
	static_cast<QFocusFrame*>(ptr)->~QFocusFrame();
}

