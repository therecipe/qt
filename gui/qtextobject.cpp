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

QtObjectPtr QTextObject_Document(QtObjectPtr ptr){
	return static_cast<QTextObject*>(ptr)->document();
}

int QTextObject_FormatIndex(QtObjectPtr ptr){
	return static_cast<QTextObject*>(ptr)->formatIndex();
}

int QTextObject_ObjectIndex(QtObjectPtr ptr){
	return static_cast<QTextObject*>(ptr)->objectIndex();
}

