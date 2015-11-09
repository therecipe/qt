#include "qxmllocator.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QXmlLocator>
#include "_cgo_export.h"

class MyQXmlLocator: public QXmlLocator {
public:
};

int QXmlLocator_ColumnNumber(void* ptr){
	return static_cast<QXmlLocator*>(ptr)->columnNumber();
}

int QXmlLocator_LineNumber(void* ptr){
	return static_cast<QXmlLocator*>(ptr)->lineNumber();
}

void QXmlLocator_DestroyQXmlLocator(void* ptr){
	static_cast<QXmlLocator*>(ptr)->~QXmlLocator();
}

