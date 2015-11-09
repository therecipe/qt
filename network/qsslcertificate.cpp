#include "qsslcertificate.h"
#include <QDateTime>
#include <QCryptographicHash>
#include <QDate>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QByteArray>
#include <QSslCertificate>
#include "_cgo_export.h"

class MyQSslCertificate: public QSslCertificate {
public:
};

void* QSslCertificate_NewQSslCertificate3(void* other){
	return new QSslCertificate(*static_cast<QSslCertificate*>(other));
}

void QSslCertificate_Clear(void* ptr){
	static_cast<QSslCertificate*>(ptr)->clear();
}

void* QSslCertificate_Digest(void* ptr, int algorithm){
	return new QByteArray(static_cast<QSslCertificate*>(ptr)->digest(static_cast<QCryptographicHash::Algorithm>(algorithm)));
}

int QSslCertificate_IsBlacklisted(void* ptr){
	return static_cast<QSslCertificate*>(ptr)->isBlacklisted();
}

void QSslCertificate_Swap(void* ptr, void* other){
	static_cast<QSslCertificate*>(ptr)->swap(*static_cast<QSslCertificate*>(other));
}

void QSslCertificate_DestroyQSslCertificate(void* ptr){
	static_cast<QSslCertificate*>(ptr)->~QSslCertificate();
}

void* QSslCertificate_EffectiveDate(void* ptr){
	return new QDateTime(static_cast<QSslCertificate*>(ptr)->effectiveDate());
}

void* QSslCertificate_ExpiryDate(void* ptr){
	return new QDateTime(static_cast<QSslCertificate*>(ptr)->expiryDate());
}

int QSslCertificate_IsNull(void* ptr){
	return static_cast<QSslCertificate*>(ptr)->isNull();
}

int QSslCertificate_IsSelfSigned(void* ptr){
	return static_cast<QSslCertificate*>(ptr)->isSelfSigned();
}

char* QSslCertificate_IssuerInfo(void* ptr, int subject){
	return static_cast<QSslCertificate*>(ptr)->issuerInfo(static_cast<QSslCertificate::SubjectInfo>(subject)).join("|").toUtf8().data();
}

char* QSslCertificate_IssuerInfo2(void* ptr, void* attribute){
	return static_cast<QSslCertificate*>(ptr)->issuerInfo(*static_cast<QByteArray*>(attribute)).join("|").toUtf8().data();
}

void* QSslCertificate_SerialNumber(void* ptr){
	return new QByteArray(static_cast<QSslCertificate*>(ptr)->serialNumber());
}

char* QSslCertificate_SubjectInfo(void* ptr, int subject){
	return static_cast<QSslCertificate*>(ptr)->subjectInfo(static_cast<QSslCertificate::SubjectInfo>(subject)).join("|").toUtf8().data();
}

char* QSslCertificate_SubjectInfo2(void* ptr, void* attribute){
	return static_cast<QSslCertificate*>(ptr)->subjectInfo(*static_cast<QByteArray*>(attribute)).join("|").toUtf8().data();
}

void* QSslCertificate_ToDer(void* ptr){
	return new QByteArray(static_cast<QSslCertificate*>(ptr)->toDer());
}

void* QSslCertificate_ToPem(void* ptr){
	return new QByteArray(static_cast<QSslCertificate*>(ptr)->toPem());
}

char* QSslCertificate_ToText(void* ptr){
	return static_cast<QSslCertificate*>(ptr)->toText().toUtf8().data();
}

void* QSslCertificate_Version(void* ptr){
	return new QByteArray(static_cast<QSslCertificate*>(ptr)->version());
}

