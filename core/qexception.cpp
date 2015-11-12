#include "qexception.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QException>
#include "_cgo_export.h"

class MyQException: public QException {
public:
};

void* QException_Clone(void* ptr){
	return static_cast<QException*>(ptr)->clone();
}

void QException_Raise(void* ptr){
	static_cast<QException*>(ptr)->raise();
}

