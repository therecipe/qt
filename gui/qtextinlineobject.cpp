#include "qtextinlineobject.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QTextInlineObject>
#include "_cgo_export.h"

class MyQTextInlineObject: public QTextInlineObject {
public:
};

int QTextInlineObject_FormatIndex(QtObjectPtr ptr){
	return static_cast<QTextInlineObject*>(ptr)->formatIndex();
}

int QTextInlineObject_IsValid(QtObjectPtr ptr){
	return static_cast<QTextInlineObject*>(ptr)->isValid();
}

int QTextInlineObject_TextDirection(QtObjectPtr ptr){
	return static_cast<QTextInlineObject*>(ptr)->textDirection();
}

int QTextInlineObject_TextPosition(QtObjectPtr ptr){
	return static_cast<QTextInlineObject*>(ptr)->textPosition();
}

