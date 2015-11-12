#include "qsslpresharedkeyauthenticator.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QByteArray>
#include <QSslPreSharedKeyAuthenticator>
#include "_cgo_export.h"

class MyQSslPreSharedKeyAuthenticator: public QSslPreSharedKeyAuthenticator {
public:
};

void* QSslPreSharedKeyAuthenticator_NewQSslPreSharedKeyAuthenticator(){
	return new QSslPreSharedKeyAuthenticator();
}

void* QSslPreSharedKeyAuthenticator_NewQSslPreSharedKeyAuthenticator2(void* authenticator){
	return new QSslPreSharedKeyAuthenticator(*static_cast<QSslPreSharedKeyAuthenticator*>(authenticator));
}

void* QSslPreSharedKeyAuthenticator_Identity(void* ptr){
	return new QByteArray(static_cast<QSslPreSharedKeyAuthenticator*>(ptr)->identity());
}

void* QSslPreSharedKeyAuthenticator_IdentityHint(void* ptr){
	return new QByteArray(static_cast<QSslPreSharedKeyAuthenticator*>(ptr)->identityHint());
}

int QSslPreSharedKeyAuthenticator_MaximumIdentityLength(void* ptr){
	return static_cast<QSslPreSharedKeyAuthenticator*>(ptr)->maximumIdentityLength();
}

int QSslPreSharedKeyAuthenticator_MaximumPreSharedKeyLength(void* ptr){
	return static_cast<QSslPreSharedKeyAuthenticator*>(ptr)->maximumPreSharedKeyLength();
}

void* QSslPreSharedKeyAuthenticator_PreSharedKey(void* ptr){
	return new QByteArray(static_cast<QSslPreSharedKeyAuthenticator*>(ptr)->preSharedKey());
}

void QSslPreSharedKeyAuthenticator_SetIdentity(void* ptr, void* identity){
	static_cast<QSslPreSharedKeyAuthenticator*>(ptr)->setIdentity(*static_cast<QByteArray*>(identity));
}

void QSslPreSharedKeyAuthenticator_SetPreSharedKey(void* ptr, void* preSharedKey){
	static_cast<QSslPreSharedKeyAuthenticator*>(ptr)->setPreSharedKey(*static_cast<QByteArray*>(preSharedKey));
}

void QSslPreSharedKeyAuthenticator_Swap(void* ptr, void* authenticator){
	static_cast<QSslPreSharedKeyAuthenticator*>(ptr)->swap(*static_cast<QSslPreSharedKeyAuthenticator*>(authenticator));
}

void QSslPreSharedKeyAuthenticator_DestroyQSslPreSharedKeyAuthenticator(void* ptr){
	static_cast<QSslPreSharedKeyAuthenticator*>(ptr)->~QSslPreSharedKeyAuthenticator();
}

