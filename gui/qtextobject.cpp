#include "qtextobject.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QTextObject>
#include "_cgo_export.h"

class MyQTextObject: public QTextObject {
public:
};

void* QTextObject_Document(void* ptr){
	return static_cast<QTextObject*>(ptr)->document();
}

int QTextObject_FormatIndex(void* ptr){
	return static_cast<QTextObject*>(ptr)->formatIndex();
}

int QTextObject_ObjectIndex(void* ptr){
	return static_cast<QTextObject*>(ptr)->objectIndex();
}

