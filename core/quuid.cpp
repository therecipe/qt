#include "quuid.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QByteArray>
#include <QUuid>
#include "_cgo_export.h"

class MyQUuid: public QUuid {
public:
};

int QUuid_Variant(QtObjectPtr ptr){
	return static_cast<QUuid*>(ptr)->variant();
}

int QUuid_Version(QtObjectPtr ptr){
	return static_cast<QUuid*>(ptr)->version();
}

QtObjectPtr QUuid_NewQUuid(){
	return new QUuid();
}

QtObjectPtr QUuid_NewQUuid5(QtObjectPtr text){
	return new QUuid(*static_cast<QByteArray*>(text));
}

QtObjectPtr QUuid_NewQUuid3(char* text){
	return new QUuid(QString(text));
}

int QUuid_IsNull(QtObjectPtr ptr){
	return static_cast<QUuid*>(ptr)->isNull();
}

char* QUuid_ToString(QtObjectPtr ptr){
	return static_cast<QUuid*>(ptr)->toString().toUtf8().data();
}

