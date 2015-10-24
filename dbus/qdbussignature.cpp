#include "qdbussignature.h"
#include <QLatin1String>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QDBusSignature>
#include "_cgo_export.h"

class MyQDBusSignature: public QDBusSignature {
public:
};

QtObjectPtr QDBusSignature_NewQDBusSignature(){
	return new QDBusSignature();
}

QtObjectPtr QDBusSignature_NewQDBusSignature3(QtObjectPtr signature){
	return new QDBusSignature(*static_cast<QLatin1String*>(signature));
}

QtObjectPtr QDBusSignature_NewQDBusSignature4(char* signature){
	return new QDBusSignature(QString(signature));
}

QtObjectPtr QDBusSignature_NewQDBusSignature2(char* signature){
	return new QDBusSignature(const_cast<const char*>(signature));
}

void QDBusSignature_SetSignature(QtObjectPtr ptr, char* signature){
	static_cast<QDBusSignature*>(ptr)->setSignature(QString(signature));
}

char* QDBusSignature_Signature(QtObjectPtr ptr){
	return static_cast<QDBusSignature*>(ptr)->signature().toUtf8().data();
}

