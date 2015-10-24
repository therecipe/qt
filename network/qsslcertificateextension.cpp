#include "qsslcertificateextension.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSslCertificate>
#include <QSslCertificateExtension>
#include "_cgo_export.h"

class MyQSslCertificateExtension: public QSslCertificateExtension {
public:
};

QtObjectPtr QSslCertificateExtension_NewQSslCertificateExtension(){
	return new QSslCertificateExtension();
}

QtObjectPtr QSslCertificateExtension_NewQSslCertificateExtension2(QtObjectPtr other){
	return new QSslCertificateExtension(*static_cast<QSslCertificateExtension*>(other));
}

int QSslCertificateExtension_IsCritical(QtObjectPtr ptr){
	return static_cast<QSslCertificateExtension*>(ptr)->isCritical();
}

int QSslCertificateExtension_IsSupported(QtObjectPtr ptr){
	return static_cast<QSslCertificateExtension*>(ptr)->isSupported();
}

char* QSslCertificateExtension_Name(QtObjectPtr ptr){
	return static_cast<QSslCertificateExtension*>(ptr)->name().toUtf8().data();
}

char* QSslCertificateExtension_Oid(QtObjectPtr ptr){
	return static_cast<QSslCertificateExtension*>(ptr)->oid().toUtf8().data();
}

void QSslCertificateExtension_Swap(QtObjectPtr ptr, QtObjectPtr other){
	static_cast<QSslCertificateExtension*>(ptr)->swap(*static_cast<QSslCertificateExtension*>(other));
}

char* QSslCertificateExtension_Value(QtObjectPtr ptr){
	return static_cast<QSslCertificateExtension*>(ptr)->value().toString().toUtf8().data();
}

void QSslCertificateExtension_DestroyQSslCertificateExtension(QtObjectPtr ptr){
	static_cast<QSslCertificateExtension*>(ptr)->~QSslCertificateExtension();
}

