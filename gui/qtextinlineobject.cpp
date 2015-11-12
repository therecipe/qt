#include "qtextinlineobject.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QString>
#include <QTextInlineObject>
#include "_cgo_export.h"

class MyQTextInlineObject: public QTextInlineObject {
public:
};

double QTextInlineObject_Ascent(void* ptr){
	return static_cast<double>(static_cast<QTextInlineObject*>(ptr)->ascent());
}

double QTextInlineObject_Descent(void* ptr){
	return static_cast<double>(static_cast<QTextInlineObject*>(ptr)->descent());
}

int QTextInlineObject_FormatIndex(void* ptr){
	return static_cast<QTextInlineObject*>(ptr)->formatIndex();
}

double QTextInlineObject_Height(void* ptr){
	return static_cast<double>(static_cast<QTextInlineObject*>(ptr)->height());
}

int QTextInlineObject_IsValid(void* ptr){
	return static_cast<QTextInlineObject*>(ptr)->isValid();
}

void QTextInlineObject_SetAscent(void* ptr, double a){
	static_cast<QTextInlineObject*>(ptr)->setAscent(static_cast<qreal>(a));
}

void QTextInlineObject_SetDescent(void* ptr, double d){
	static_cast<QTextInlineObject*>(ptr)->setDescent(static_cast<qreal>(d));
}

void QTextInlineObject_SetWidth(void* ptr, double w){
	static_cast<QTextInlineObject*>(ptr)->setWidth(static_cast<qreal>(w));
}

int QTextInlineObject_TextDirection(void* ptr){
	return static_cast<QTextInlineObject*>(ptr)->textDirection();
}

int QTextInlineObject_TextPosition(void* ptr){
	return static_cast<QTextInlineObject*>(ptr)->textPosition();
}

double QTextInlineObject_Width(void* ptr){
	return static_cast<double>(static_cast<QTextInlineObject*>(ptr)->width());
}

