#include "qauthenticator.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QAuthenticator>
#include "_cgo_export.h"

class MyQAuthenticator: public QAuthenticator {
public:
};

QtObjectPtr QAuthenticator_NewQAuthenticator(){
	return new QAuthenticator();
}

QtObjectPtr QAuthenticator_NewQAuthenticator2(QtObjectPtr other){
	return new QAuthenticator(*static_cast<QAuthenticator*>(other));
}

int QAuthenticator_IsNull(QtObjectPtr ptr){
	return static_cast<QAuthenticator*>(ptr)->isNull();
}

char* QAuthenticator_Option(QtObjectPtr ptr, char* opt){
	return static_cast<QAuthenticator*>(ptr)->option(QString(opt)).toString().toUtf8().data();
}

char* QAuthenticator_Password(QtObjectPtr ptr){
	return static_cast<QAuthenticator*>(ptr)->password().toUtf8().data();
}

char* QAuthenticator_Realm(QtObjectPtr ptr){
	return static_cast<QAuthenticator*>(ptr)->realm().toUtf8().data();
}

void QAuthenticator_SetOption(QtObjectPtr ptr, char* opt, char* value){
	static_cast<QAuthenticator*>(ptr)->setOption(QString(opt), QVariant(value));
}

void QAuthenticator_SetPassword(QtObjectPtr ptr, char* password){
	static_cast<QAuthenticator*>(ptr)->setPassword(QString(password));
}

void QAuthenticator_SetUser(QtObjectPtr ptr, char* user){
	static_cast<QAuthenticator*>(ptr)->setUser(QString(user));
}

char* QAuthenticator_User(QtObjectPtr ptr){
	return static_cast<QAuthenticator*>(ptr)->user().toUtf8().data();
}

void QAuthenticator_DestroyQAuthenticator(QtObjectPtr ptr){
	static_cast<QAuthenticator*>(ptr)->~QAuthenticator();
}

