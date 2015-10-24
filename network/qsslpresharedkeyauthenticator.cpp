#include "qsslpresharedkeyauthenticator.h"
#include <QByteArray>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSslPreSharedKeyAuthenticator>
#include "_cgo_export.h"

class MyQSslPreSharedKeyAuthenticator: public QSslPreSharedKeyAuthenticator {
public:
};

QtObjectPtr QSslPreSharedKeyAuthenticator_NewQSslPreSharedKeyAuthenticator(){
	return new QSslPreSharedKeyAuthenticator();
}

QtObjectPtr QSslPreSharedKeyAuthenticator_NewQSslPreSharedKeyAuthenticator2(QtObjectPtr authenticator){
	return new QSslPreSharedKeyAuthenticator(*static_cast<QSslPreSharedKeyAuthenticator*>(authenticator));
}

int QSslPreSharedKeyAuthenticator_MaximumIdentityLength(QtObjectPtr ptr){
	return static_cast<QSslPreSharedKeyAuthenticator*>(ptr)->maximumIdentityLength();
}

int QSslPreSharedKeyAuthenticator_MaximumPreSharedKeyLength(QtObjectPtr ptr){
	return static_cast<QSslPreSharedKeyAuthenticator*>(ptr)->maximumPreSharedKeyLength();
}

void QSslPreSharedKeyAuthenticator_SetIdentity(QtObjectPtr ptr, QtObjectPtr identity){
	static_cast<QSslPreSharedKeyAuthenticator*>(ptr)->setIdentity(*static_cast<QByteArray*>(identity));
}

void QSslPreSharedKeyAuthenticator_SetPreSharedKey(QtObjectPtr ptr, QtObjectPtr preSharedKey){
	static_cast<QSslPreSharedKeyAuthenticator*>(ptr)->setPreSharedKey(*static_cast<QByteArray*>(preSharedKey));
}

void QSslPreSharedKeyAuthenticator_Swap(QtObjectPtr ptr, QtObjectPtr authenticator){
	static_cast<QSslPreSharedKeyAuthenticator*>(ptr)->swap(*static_cast<QSslPreSharedKeyAuthenticator*>(authenticator));
}

void QSslPreSharedKeyAuthenticator_DestroyQSslPreSharedKeyAuthenticator(QtObjectPtr ptr){
	static_cast<QSslPreSharedKeyAuthenticator*>(ptr)->~QSslPreSharedKeyAuthenticator();
}

