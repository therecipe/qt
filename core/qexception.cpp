#include "qexception.h"
#include <QUrl>
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QException>
#include "_cgo_export.h"

class MyQException: public QException {
public:
};

QtObjectPtr QException_Clone(QtObjectPtr ptr){
	return static_cast<QException*>(ptr)->clone();
}

void QException_Raise(QtObjectPtr ptr){
	static_cast<QException*>(ptr)->raise();
}

