#include "qsslkey.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSslKey>
#include "_cgo_export.h"

class MyQSslKey: public QSslKey {
public:
};

QtObjectPtr QSslKey_NewQSslKey(){
	return new QSslKey();
}

QtObjectPtr QSslKey_NewQSslKey5(QtObjectPtr other){
	return new QSslKey(*static_cast<QSslKey*>(other));
}

void QSslKey_Clear(QtObjectPtr ptr){
	static_cast<QSslKey*>(ptr)->clear();
}

int QSslKey_IsNull(QtObjectPtr ptr){
	return static_cast<QSslKey*>(ptr)->isNull();
}

int QSslKey_Length(QtObjectPtr ptr){
	return static_cast<QSslKey*>(ptr)->length();
}

void QSslKey_Swap(QtObjectPtr ptr, QtObjectPtr other){
	static_cast<QSslKey*>(ptr)->swap(*static_cast<QSslKey*>(other));
}

void QSslKey_DestroyQSslKey(QtObjectPtr ptr){
	static_cast<QSslKey*>(ptr)->~QSslKey();
}

