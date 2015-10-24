#include "qdbusinterface.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QDBusConnection>
#include <QDBusInterface>
#include "_cgo_export.h"

class MyQDBusInterface: public QDBusInterface {
public:
};

QtObjectPtr QDBusInterface_NewQDBusInterface(char* service, char* path, char* interfa, QtObjectPtr connection, QtObjectPtr parent){
	return new QDBusInterface(QString(service), QString(path), QString(interfa), *static_cast<QDBusConnection*>(connection), static_cast<QObject*>(parent));
}

void QDBusInterface_DestroyQDBusInterface(QtObjectPtr ptr){
	static_cast<QDBusInterface*>(ptr)->~QDBusInterface();
}

