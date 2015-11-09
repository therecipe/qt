#include "qsqldriver.h"
#include <QSqlRecord>
#include <QSqlField>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QSqlDriver>
#include "_cgo_export.h"

class MyQSqlDriver: public QSqlDriver {
public:
void Signal_Notification(const QString & name){callbackQSqlDriverNotification(this->objectName().toUtf8().data(), name.toUtf8().data());};
};

int QSqlDriver_BeginTransaction(void* ptr){
	return static_cast<QSqlDriver*>(ptr)->beginTransaction();
}

void QSqlDriver_Close(void* ptr){
	static_cast<QSqlDriver*>(ptr)->close();
}

int QSqlDriver_CommitTransaction(void* ptr){
	return static_cast<QSqlDriver*>(ptr)->commitTransaction();
}

void* QSqlDriver_CreateResult(void* ptr){
	return static_cast<QSqlDriver*>(ptr)->createResult();
}

int QSqlDriver_DbmsType(void* ptr){
	return static_cast<QSqlDriver*>(ptr)->dbmsType();
}

char* QSqlDriver_EscapeIdentifier(void* ptr, char* identifier, int ty){
	return static_cast<QSqlDriver*>(ptr)->escapeIdentifier(QString(identifier), static_cast<QSqlDriver::IdentifierType>(ty)).toUtf8().data();
}

char* QSqlDriver_FormatValue(void* ptr, void* field, int trimStrings){
	return static_cast<QSqlDriver*>(ptr)->formatValue(*static_cast<QSqlField*>(field), trimStrings != 0).toUtf8().data();
}

void* QSqlDriver_Handle(void* ptr){
	return new QVariant(static_cast<QSqlDriver*>(ptr)->handle());
}

int QSqlDriver_HasFeature(void* ptr, int feature){
	return static_cast<QSqlDriver*>(ptr)->hasFeature(static_cast<QSqlDriver::DriverFeature>(feature));
}

int QSqlDriver_IsIdentifierEscaped(void* ptr, char* identifier, int ty){
	return static_cast<QSqlDriver*>(ptr)->isIdentifierEscaped(QString(identifier), static_cast<QSqlDriver::IdentifierType>(ty));
}

int QSqlDriver_IsOpen(void* ptr){
	return static_cast<QSqlDriver*>(ptr)->isOpen();
}

int QSqlDriver_IsOpenError(void* ptr){
	return static_cast<QSqlDriver*>(ptr)->isOpenError();
}

void QSqlDriver_ConnectNotification(void* ptr){
	QObject::connect(static_cast<QSqlDriver*>(ptr), static_cast<void (QSqlDriver::*)(const QString &)>(&QSqlDriver::notification), static_cast<MyQSqlDriver*>(ptr), static_cast<void (MyQSqlDriver::*)(const QString &)>(&MyQSqlDriver::Signal_Notification));;
}

void QSqlDriver_DisconnectNotification(void* ptr){
	QObject::disconnect(static_cast<QSqlDriver*>(ptr), static_cast<void (QSqlDriver::*)(const QString &)>(&QSqlDriver::notification), static_cast<MyQSqlDriver*>(ptr), static_cast<void (MyQSqlDriver::*)(const QString &)>(&MyQSqlDriver::Signal_Notification));;
}

int QSqlDriver_Open(void* ptr, char* db, char* user, char* password, char* host, int port, char* options){
	return static_cast<QSqlDriver*>(ptr)->open(QString(db), QString(user), QString(password), QString(host), port, QString(options));
}

int QSqlDriver_RollbackTransaction(void* ptr){
	return static_cast<QSqlDriver*>(ptr)->rollbackTransaction();
}

char* QSqlDriver_SqlStatement(void* ptr, int ty, char* tableName, void* rec, int preparedStatement){
	return static_cast<QSqlDriver*>(ptr)->sqlStatement(static_cast<QSqlDriver::StatementType>(ty), QString(tableName), *static_cast<QSqlRecord*>(rec), preparedStatement != 0).toUtf8().data();
}

char* QSqlDriver_StripDelimiters(void* ptr, char* identifier, int ty){
	return static_cast<QSqlDriver*>(ptr)->stripDelimiters(QString(identifier), static_cast<QSqlDriver::IdentifierType>(ty)).toUtf8().data();
}

int QSqlDriver_SubscribeToNotification(void* ptr, char* name){
	return static_cast<QSqlDriver*>(ptr)->subscribeToNotification(QString(name));
}

char* QSqlDriver_SubscribedToNotifications(void* ptr){
	return static_cast<QSqlDriver*>(ptr)->subscribedToNotifications().join("|").toUtf8().data();
}

int QSqlDriver_UnsubscribeFromNotification(void* ptr, char* name){
	return static_cast<QSqlDriver*>(ptr)->unsubscribeFromNotification(QString(name));
}

void QSqlDriver_DestroyQSqlDriver(void* ptr){
	static_cast<QSqlDriver*>(ptr)->~QSqlDriver();
}

