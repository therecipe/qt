#include "qframe.h"
#include <QRect>
#include <QWidget>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QFrame>
#include "_cgo_export.h"

class MyQFrame: public QFrame {
public:
};

int QFrame_FrameShadow(void* ptr){
	return static_cast<QFrame*>(ptr)->frameShadow();
}

int QFrame_FrameShape(void* ptr){
	return static_cast<QFrame*>(ptr)->frameShape();
}

int QFrame_FrameWidth(void* ptr){
	return static_cast<QFrame*>(ptr)->frameWidth();
}

int QFrame_LineWidth(void* ptr){
	return static_cast<QFrame*>(ptr)->lineWidth();
}

int QFrame_MidLineWidth(void* ptr){
	return static_cast<QFrame*>(ptr)->midLineWidth();
}

void QFrame_SetFrameRect(void* ptr, void* v){
	static_cast<QFrame*>(ptr)->setFrameRect(*static_cast<QRect*>(v));
}

void QFrame_SetFrameShadow(void* ptr, int v){
	static_cast<QFrame*>(ptr)->setFrameShadow(static_cast<QFrame::Shadow>(v));
}

void QFrame_SetFrameShape(void* ptr, int v){
	static_cast<QFrame*>(ptr)->setFrameShape(static_cast<QFrame::Shape>(v));
}

void QFrame_SetLineWidth(void* ptr, int v){
	static_cast<QFrame*>(ptr)->setLineWidth(v);
}

void QFrame_SetMidLineWidth(void* ptr, int v){
	static_cast<QFrame*>(ptr)->setMidLineWidth(v);
}

void* QFrame_NewQFrame(void* parent, int f){
	return new QFrame(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(f));
}

int QFrame_FrameStyle(void* ptr){
	return static_cast<QFrame*>(ptr)->frameStyle();
}

void QFrame_SetFrameStyle(void* ptr, int style){
	static_cast<QFrame*>(ptr)->setFrameStyle(style);
}

void QFrame_DestroyQFrame(void* ptr){
	static_cast<QFrame*>(ptr)->~QFrame();
}

