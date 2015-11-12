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

void* QAuthenticator_NewQAuthenticator(){
	return new QAuthenticator();
}

void* QAuthenticator_NewQAuthenticator2(void* other){
	return new QAuthenticator(*static_cast<QAuthenticator*>(other));
}

int QAuthenticator_IsNull(void* ptr){
	return static_cast<QAuthenticator*>(ptr)->isNull();
}

void* QAuthenticator_Option(void* ptr, char* opt){
	return new QVariant(static_cast<QAuthenticator*>(ptr)->option(QString(opt)));
}

char* QAuthenticator_Password(void* ptr){
	return static_cast<QAuthenticator*>(ptr)->password().toUtf8().data();
}

char* QAuthenticator_Realm(void* ptr){
	return static_cast<QAuthenticator*>(ptr)->realm().toUtf8().data();
}

void QAuthenticator_SetOption(void* ptr, char* opt, void* value){
	static_cast<QAuthenticator*>(ptr)->setOption(QString(opt), *static_cast<QVariant*>(value));
}

void QAuthenticator_SetPassword(void* ptr, char* password){
	static_cast<QAuthenticator*>(ptr)->setPassword(QString(password));
}

void QAuthenticator_SetUser(void* ptr, char* user){
	static_cast<QAuthenticator*>(ptr)->setUser(QString(user));
}

char* QAuthenticator_User(void* ptr){
	return static_cast<QAuthenticator*>(ptr)->user().toUtf8().data();
}

void QAuthenticator_DestroyQAuthenticator(void* ptr){
	static_cast<QAuthenticator*>(ptr)->~QAuthenticator();
}

