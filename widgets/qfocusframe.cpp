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

void* QFocusFrame_NewQFocusFrame(void* parent){
	return new QFocusFrame(static_cast<QWidget*>(parent));
}

void QFocusFrame_SetWidget(void* ptr, void* widget){
	static_cast<QFocusFrame*>(ptr)->setWidget(static_cast<QWidget*>(widget));
}

void* QFocusFrame_Widget(void* ptr){
	return static_cast<QFocusFrame*>(ptr)->widget();
}

void QFocusFrame_DestroyQFocusFrame(void* ptr){
	static_cast<QFocusFrame*>(ptr)->~QFocusFrame();
}

