#include "qframe.h"
#include <QUrl>
#include <QModelIndex>
#include <QRect>
#include <QWidget>
#include <QString>
#include <QVariant>
#include <QFrame>
#include "_cgo_export.h"

class MyQFrame: public QFrame {
public:
};

int QFrame_FrameShadow(QtObjectPtr ptr){
	return static_cast<QFrame*>(ptr)->frameShadow();
}

int QFrame_FrameShape(QtObjectPtr ptr){
	return static_cast<QFrame*>(ptr)->frameShape();
}

int QFrame_FrameWidth(QtObjectPtr ptr){
	return static_cast<QFrame*>(ptr)->frameWidth();
}

int QFrame_LineWidth(QtObjectPtr ptr){
	return static_cast<QFrame*>(ptr)->lineWidth();
}

int QFrame_MidLineWidth(QtObjectPtr ptr){
	return static_cast<QFrame*>(ptr)->midLineWidth();
}

void QFrame_SetFrameRect(QtObjectPtr ptr, QtObjectPtr v){
	static_cast<QFrame*>(ptr)->setFrameRect(*static_cast<QRect*>(v));
}

void QFrame_SetFrameShadow(QtObjectPtr ptr, int v){
	static_cast<QFrame*>(ptr)->setFrameShadow(static_cast<QFrame::Shadow>(v));
}

void QFrame_SetFrameShape(QtObjectPtr ptr, int v){
	static_cast<QFrame*>(ptr)->setFrameShape(static_cast<QFrame::Shape>(v));
}

void QFrame_SetLineWidth(QtObjectPtr ptr, int v){
	static_cast<QFrame*>(ptr)->setLineWidth(v);
}

void QFrame_SetMidLineWidth(QtObjectPtr ptr, int v){
	static_cast<QFrame*>(ptr)->setMidLineWidth(v);
}

QtObjectPtr QFrame_NewQFrame(QtObjectPtr parent, int f){
	return new QFrame(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(f));
}

int QFrame_FrameStyle(QtObjectPtr ptr){
	return static_cast<QFrame*>(ptr)->frameStyle();
}

void QFrame_SetFrameStyle(QtObjectPtr ptr, int style){
	static_cast<QFrame*>(ptr)->setFrameStyle(style);
}

void QFrame_DestroyQFrame(QtObjectPtr ptr){
	static_cast<QFrame*>(ptr)->~QFrame();
}

