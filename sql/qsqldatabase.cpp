#include "qsqldatabase.h"
#include <QSqlDriverCreator>
#include <QSqlDriverCreatorBase>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSqlDriver>
#include <QSqlDatabase>
#include "_cgo_export.h"

class MyQSqlDatabase: public QSqlDatabase {
public:
};

QtObjectPtr QSqlDatabase_NewQSqlDatabase(){
	return new QSqlDatabase();
}

QtObjectPtr QSqlDatabase_NewQSqlDatabase2(QtObjectPtr other){
	return new QSqlDatabase(*static_cast<QSqlDatabase*>(other));
}

void QSqlDatabase_Close(QtObjectPtr ptr){
	static_cast<QSqlDatabase*>(ptr)->close();
}

int QSqlDatabase_Commit(QtObjectPtr ptr){
	return static_cast<QSqlDatabase*>(ptr)->commit();
}

char* QSqlDatabase_ConnectOptions(QtObjectPtr ptr){
	return static_cast<QSqlDatabase*>(ptr)->connectOptions().toUtf8().data();
}

char* QSqlDatabase_ConnectionName(QtObjectPtr ptr){
	return static_cast<QSqlDatabase*>(ptr)->connectionName().toUtf8().data();
}

char* QSqlDatabase_QSqlDatabase_ConnectionNames(){
	return QSqlDatabase::connectionNames().join("|").toUtf8().data();
}

int QSqlDatabase_QSqlDatabase_Contains(char* connectionName){
	return QSqlDatabase::contains(QString(connectionName));
}

char* QSqlDatabase_DatabaseName(QtObjectPtr ptr){
	return static_cast<QSqlDatabase*>(ptr)->databaseName().toUtf8().data();
}

QtObjectPtr QSqlDatabase_Driver(QtObjectPtr ptr){
	return static_cast<QSqlDatabase*>(ptr)->driver();
}

char* QSqlDatabase_DriverName(QtObjectPtr ptr){
	return static_cast<QSqlDatabase*>(ptr)->driverName().toUtf8().data();
}

char* QSqlDatabase_QSqlDatabase_Drivers(){
	return QSqlDatabase::drivers().join("|").toUtf8().data();
}

char* QSqlDatabase_HostName(QtObjectPtr ptr){
	return static_cast<QSqlDatabase*>(ptr)->hostName().toUtf8().data();
}

int QSqlDatabase_QSqlDatabase_IsDriverAvailable(char* name){
	return QSqlDatabase::isDriverAvailable(QString(name));
}

int QSqlDatabase_IsOpen(QtObjectPtr ptr){
	return static_cast<QSqlDatabase*>(ptr)->isOpen();
}

int QSqlDatabase_IsOpenError(QtObjectPtr ptr){
	return static_cast<QSqlDatabase*>(ptr)->isOpenError();
}

int QSqlDatabase_IsValid(QtObjectPtr ptr){
	return static_cast<QSqlDatabase*>(ptr)->isValid();
}

int QSqlDatabase_Open(QtObjectPtr ptr){
	return static_cast<QSqlDatabase*>(ptr)->open();
}

int QSqlDatabase_Open2(QtObjectPtr ptr, char* user, char* password){
	return static_cast<QSqlDatabase*>(ptr)->open(QString(user), QString(password));
}

char* QSqlDatabase_Password(QtObjectPtr ptr){
	return static_cast<QSqlDatabase*>(ptr)->password().toUtf8().data();
}

int QSqlDatabase_Port(QtObjectPtr ptr){
	return static_cast<QSqlDatabase*>(ptr)->port();
}

void QSqlDatabase_QSqlDatabase_RegisterSqlDriver(char* name, QtObjectPtr creator){
	QSqlDatabase::registerSqlDriver(QString(name), static_cast<QSqlDriverCreatorBase*>(creator));
}

void QSqlDatabase_QSqlDatabase_RemoveDatabase(char* connectionName){
	QSqlDatabase::removeDatabase(QString(connectionName));
}

int QSqlDatabase_Rollback(QtObjectPtr ptr){
	return static_cast<QSqlDatabase*>(ptr)->rollback();
}

void QSqlDatabase_SetConnectOptions(QtObjectPtr ptr, char* options){
	static_cast<QSqlDatabase*>(ptr)->setConnectOptions(QString(options));
}

void QSqlDatabase_SetDatabaseName(QtObjectPtr ptr, char* name){
	static_cast<QSqlDatabase*>(ptr)->setDatabaseName(QString(name));
}

void QSqlDatabase_SetHostName(QtObjectPtr ptr, char* host){
	static_cast<QSqlDatabase*>(ptr)->setHostName(QString(host));
}

void QSqlDatabase_SetPassword(QtObjectPtr ptr, char* password){
	static_cast<QSqlDatabase*>(ptr)->setPassword(QString(password));
}

void QSqlDatabase_SetPort(QtObjectPtr ptr, int port){
	static_cast<QSqlDatabase*>(ptr)->setPort(port);
}

void QSqlDatabase_SetUserName(QtObjectPtr ptr, char* name){
	static_cast<QSqlDatabase*>(ptr)->setUserName(QString(name));
}

int QSqlDatabase_Transaction(QtObjectPtr ptr){
	return static_cast<QSqlDatabase*>(ptr)->transaction();
}

char* QSqlDatabase_UserName(QtObjectPtr ptr){
	return static_cast<QSqlDatabase*>(ptr)->userName().toUtf8().data();
}

void QSqlDatabase_DestroyQSqlDatabase(QtObjectPtr ptr){
	static_cast<QSqlDatabase*>(ptr)->~QSqlDatabase();
}

