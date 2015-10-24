#include "qdbusvariant.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QString>
#include <QDBusVariant>
#include "_cgo_export.h"

class MyQDBusVariant: public QDBusVariant {
public:
};

QtObjectPtr QDBusVariant_NewQDBusVariant(){
	return new QDBusVariant();
}

QtObjectPtr QDBusVariant_NewQDBusVariant2(char* variant){
	return new QDBusVariant(QVariant(variant));
}

void QDBusVariant_SetVariant(QtObjectPtr ptr, char* variant){
	static_cast<QDBusVariant*>(ptr)->setVariant(QVariant(variant));
}

char* QDBusVariant_Variant(QtObjectPtr ptr){
	return static_cast<QDBusVariant*>(ptr)->variant().toString().toUtf8().data();
}

