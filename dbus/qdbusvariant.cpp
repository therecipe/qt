#include "qdbusvariant.h"
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QDBusVariant>
#include "_cgo_export.h"

class MyQDBusVariant: public QDBusVariant {
public:
};

void* QDBusVariant_NewQDBusVariant(){
	return new QDBusVariant();
}

void* QDBusVariant_NewQDBusVariant2(void* variant){
	return new QDBusVariant(*static_cast<QVariant*>(variant));
}

void QDBusVariant_SetVariant(void* ptr, void* variant){
	static_cast<QDBusVariant*>(ptr)->setVariant(*static_cast<QVariant*>(variant));
}

void* QDBusVariant_Variant(void* ptr){
	return new QVariant(static_cast<QDBusVariant*>(ptr)->variant());
}

