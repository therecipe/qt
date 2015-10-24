#include "qsslcertificate.h"
#include <QModelIndex>
#include <QByteArray>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QSslCertificate>
#include "_cgo_export.h"

class MyQSslCertificate: public QSslCertificate {
public:
};

QtObjectPtr QSslCertificate_NewQSslCertificate3(QtObjectPtr other){
	return new QSslCertificate(*static_cast<QSslCertificate*>(other));
}

void QSslCertificate_Clear(QtObjectPtr ptr){
	static_cast<QSslCertificate*>(ptr)->clear();
}

int QSslCertificate_IsBlacklisted(QtObjectPtr ptr){
	return static_cast<QSslCertificate*>(ptr)->isBlacklisted();
}

void QSslCertificate_Swap(QtObjectPtr ptr, QtObjectPtr other){
	static_cast<QSslCertificate*>(ptr)->swap(*static_cast<QSslCertificate*>(other));
}

void QSslCertificate_DestroyQSslCertificate(QtObjectPtr ptr){
	static_cast<QSslCertificate*>(ptr)->~QSslCertificate();
}

int QSslCertificate_IsNull(QtObjectPtr ptr){
	return static_cast<QSslCertificate*>(ptr)->isNull();
}

int QSslCertificate_IsSelfSigned(QtObjectPtr ptr){
	return static_cast<QSslCertificate*>(ptr)->isSelfSigned();
}

char* QSslCertificate_IssuerInfo(QtObjectPtr ptr, int subject){
	return static_cast<QSslCertificate*>(ptr)->issuerInfo(static_cast<QSslCertificate::SubjectInfo>(subject)).join("|").toUtf8().data();
}

char* QSslCertificate_IssuerInfo2(QtObjectPtr ptr, QtObjectPtr attribute){
	return static_cast<QSslCertificate*>(ptr)->issuerInfo(*static_cast<QByteArray*>(attribute)).join("|").toUtf8().data();
}

char* QSslCertificate_SubjectInfo(QtObjectPtr ptr, int subject){
	return static_cast<QSslCertificate*>(ptr)->subjectInfo(static_cast<QSslCertificate::SubjectInfo>(subject)).join("|").toUtf8().data();
}

char* QSslCertificate_SubjectInfo2(QtObjectPtr ptr, QtObjectPtr attribute){
	return static_cast<QSslCertificate*>(ptr)->subjectInfo(*static_cast<QByteArray*>(attribute)).join("|").toUtf8().data();
}

char* QSslCertificate_ToText(QtObjectPtr ptr){
	return static_cast<QSslCertificate*>(ptr)->toText().toUtf8().data();
}

