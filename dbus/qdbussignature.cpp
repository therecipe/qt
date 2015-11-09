#include "qdbussignature.h"
#include <QModelIndex>
#include <QLatin1String>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QDBusSignature>
#include "_cgo_export.h"

class MyQDBusSignature: public QDBusSignature {
public:
};

void* QDBusSignature_NewQDBusSignature(){
	return new QDBusSignature();
}

void* QDBusSignature_NewQDBusSignature3(void* signature){
	return new QDBusSignature(*static_cast<QLatin1String*>(signature));
}

void* QDBusSignature_NewQDBusSignature4(char* signature){
	return new QDBusSignature(QString(signature));
}

void* QDBusSignature_NewQDBusSignature2(char* signature){
	return new QDBusSignature(const_cast<const char*>(signature));
}

void QDBusSignature_SetSignature(void* ptr, char* signature){
	static_cast<QDBusSignature*>(ptr)->setSignature(QString(signature));
}

char* QDBusSignature_Signature(void* ptr){
	return static_cast<QDBusSignature*>(ptr)->signature().toUtf8().data();
}

