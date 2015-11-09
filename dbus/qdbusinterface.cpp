#include "qdbusinterface.h"
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QDBusConnection>
#include <QDBusInterface>
#include "_cgo_export.h"

class MyQDBusInterface: public QDBusInterface {
public:
};

void* QDBusInterface_NewQDBusInterface(char* service, char* path, char* interfa, void* connection, void* parent){
	return new QDBusInterface(QString(service), QString(path), QString(interfa), *static_cast<QDBusConnection*>(connection), static_cast<QObject*>(parent));
}

void QDBusInterface_DestroyQDBusInterface(void* ptr){
	static_cast<QDBusInterface*>(ptr)->~QDBusInterface();
}

