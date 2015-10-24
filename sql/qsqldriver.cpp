#include "qsqldriver.h"
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSqlField>
#include <QSqlRecord>
#include <QSqlDriver>
#include "_cgo_export.h"

class MyQSqlDriver: public QSqlDriver {
public:
void Signal_Notification(const QString & name){callbackQSqlDriverNotification(this->objectName().toUtf8().data(), name.toUtf8().data());};
};

int QSqlDriver_BeginTransaction(QtObjectPtr ptr){
	return static_cast<QSqlDriver*>(ptr)->beginTransaction();
}

void QSqlDriver_Close(QtObjectPtr ptr){
	static_cast<QSqlDriver*>(ptr)->close();
}

int QSqlDriver_CommitTransaction(QtObjectPtr ptr){
	return static_cast<QSqlDriver*>(ptr)->commitTransaction();
}

QtObjectPtr QSqlDriver_CreateResult(QtObjectPtr ptr){
	return static_cast<QSqlDriver*>(ptr)->createResult();
}

int QSqlDriver_DbmsType(QtObjectPtr ptr){
	return static_cast<QSqlDriver*>(ptr)->dbmsType();
}

char* QSqlDriver_EscapeIdentifier(QtObjectPtr ptr, char* identifier, int ty){
	return static_cast<QSqlDriver*>(ptr)->escapeIdentifier(QString(identifier), static_cast<QSqlDriver::IdentifierType>(ty)).toUtf8().data();
}

char* QSqlDriver_FormatValue(QtObjectPtr ptr, QtObjectPtr field, int trimStrings){
	return static_cast<QSqlDriver*>(ptr)->formatValue(*static_cast<QSqlField*>(field), trimStrings != 0).toUtf8().data();
}

char* QSqlDriver_Handle(QtObjectPtr ptr){
	return static_cast<QSqlDriver*>(ptr)->handle().toString().toUtf8().data();
}

int QSqlDriver_HasFeature(QtObjectPtr ptr, int feature){
	return static_cast<QSqlDriver*>(ptr)->hasFeature(static_cast<QSqlDriver::DriverFeature>(feature));
}

int QSqlDriver_IsIdentifierEscaped(QtObjectPtr ptr, char* identifier, int ty){
	return static_cast<QSqlDriver*>(ptr)->isIdentifierEscaped(QString(identifier), static_cast<QSqlDriver::IdentifierType>(ty));
}

int QSqlDriver_IsOpen(QtObjectPtr ptr){
	return static_cast<QSqlDriver*>(ptr)->isOpen();
}

int QSqlDriver_IsOpenError(QtObjectPtr ptr){
	return static_cast<QSqlDriver*>(ptr)->isOpenError();
}

void QSqlDriver_ConnectNotification(QtObjectPtr ptr){
	QObject::connect(static_cast<QSqlDriver*>(ptr), static_cast<void (QSqlDriver::*)(const QString &)>(&QSqlDriver::notification), static_cast<MyQSqlDriver*>(ptr), static_cast<void (MyQSqlDriver::*)(const QString &)>(&MyQSqlDriver::Signal_Notification));;
}

void QSqlDriver_DisconnectNotification(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QSqlDriver*>(ptr), static_cast<void (QSqlDriver::*)(const QString &)>(&QSqlDriver::notification), static_cast<MyQSqlDriver*>(ptr), static_cast<void (MyQSqlDriver::*)(const QString &)>(&MyQSqlDriver::Signal_Notification));;
}

int QSqlDriver_Open(QtObjectPtr ptr, char* db, char* user, char* password, char* host, int port, char* options){
	return static_cast<QSqlDriver*>(ptr)->open(QString(db), QString(user), QString(password), QString(host), port, QString(options));
}

int QSqlDriver_RollbackTransaction(QtObjectPtr ptr){
	return static_cast<QSqlDriver*>(ptr)->rollbackTransaction();
}

char* QSqlDriver_SqlStatement(QtObjectPtr ptr, int ty, char* tableName, QtObjectPtr rec, int preparedStatement){
	return static_cast<QSqlDriver*>(ptr)->sqlStatement(static_cast<QSqlDriver::StatementType>(ty), QString(tableName), *static_cast<QSqlRecord*>(rec), preparedStatement != 0).toUtf8().data();
}

char* QSqlDriver_StripDelimiters(QtObjectPtr ptr, char* identifier, int ty){
	return static_cast<QSqlDriver*>(ptr)->stripDelimiters(QString(identifier), static_cast<QSqlDriver::IdentifierType>(ty)).toUtf8().data();
}

int QSqlDriver_SubscribeToNotification(QtObjectPtr ptr, char* name){
	return static_cast<QSqlDriver*>(ptr)->subscribeToNotification(QString(name));
}

char* QSqlDriver_SubscribedToNotifications(QtObjectPtr ptr){
	return static_cast<QSqlDriver*>(ptr)->subscribedToNotifications().join("|").toUtf8().data();
}

int QSqlDriver_UnsubscribeFromNotification(QtObjectPtr ptr, char* name){
	return static_cast<QSqlDriver*>(ptr)->unsubscribeFromNotification(QString(name));
}

void QSqlDriver_DestroyQSqlDriver(QtObjectPtr ptr){
	static_cast<QSqlDriver*>(ptr)->~QSqlDriver();
}

