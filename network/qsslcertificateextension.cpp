#include "qsslcertificateextension.h"
#include <QModelIndex>
#include <QSslCertificate>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QSslCertificateExtension>
#include "_cgo_export.h"

class MyQSslCertificateExtension: public QSslCertificateExtension {
public:
};

void* QSslCertificateExtension_NewQSslCertificateExtension(){
	return new QSslCertificateExtension();
}

void* QSslCertificateExtension_NewQSslCertificateExtension2(void* other){
	return new QSslCertificateExtension(*static_cast<QSslCertificateExtension*>(other));
}

int QSslCertificateExtension_IsCritical(void* ptr){
	return static_cast<QSslCertificateExtension*>(ptr)->isCritical();
}

int QSslCertificateExtension_IsSupported(void* ptr){
	return static_cast<QSslCertificateExtension*>(ptr)->isSupported();
}

char* QSslCertificateExtension_Name(void* ptr){
	return static_cast<QSslCertificateExtension*>(ptr)->name().toUtf8().data();
}

char* QSslCertificateExtension_Oid(void* ptr){
	return static_cast<QSslCertificateExtension*>(ptr)->oid().toUtf8().data();
}

void QSslCertificateExtension_Swap(void* ptr, void* other){
	static_cast<QSslCertificateExtension*>(ptr)->swap(*static_cast<QSslCertificateExtension*>(other));
}

void* QSslCertificateExtension_Value(void* ptr){
	return new QVariant(static_cast<QSslCertificateExtension*>(ptr)->value());
}

void QSslCertificateExtension_DestroyQSslCertificateExtension(void* ptr){
	static_cast<QSslCertificateExtension*>(ptr)->~QSslCertificateExtension();
}

