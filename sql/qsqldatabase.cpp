#include "qsqldatabase.h"
#include <QModelIndex>
#include <QSqlDriver>
#include <QSqlDriverCreatorBase>
#include <QSqlDriverCreator>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QSqlDatabase>
#include "_cgo_export.h"

class MyQSqlDatabase: public QSqlDatabase {
public:
};

void* QSqlDatabase_NewQSqlDatabase(){
	return new QSqlDatabase();
}

void* QSqlDatabase_NewQSqlDatabase2(void* other){
	return new QSqlDatabase(*static_cast<QSqlDatabase*>(other));
}

void QSqlDatabase_Close(void* ptr){
	static_cast<QSqlDatabase*>(ptr)->close();
}

int QSqlDatabase_Commit(void* ptr){
	return static_cast<QSqlDatabase*>(ptr)->commit();
}

char* QSqlDatabase_ConnectOptions(void* ptr){
	return static_cast<QSqlDatabase*>(ptr)->connectOptions().toUtf8().data();
}

char* QSqlDatabase_ConnectionName(void* ptr){
	return static_cast<QSqlDatabase*>(ptr)->connectionName().toUtf8().data();
}

char* QSqlDatabase_QSqlDatabase_ConnectionNames(){
	return QSqlDatabase::connectionNames().join("|").toUtf8().data();
}

int QSqlDatabase_QSqlDatabase_Contains(char* connectionName){
	return QSqlDatabase::contains(QString(connectionName));
}

char* QSqlDatabase_DatabaseName(void* ptr){
	return static_cast<QSqlDatabase*>(ptr)->databaseName().toUtf8().data();
}

void* QSqlDatabase_Driver(void* ptr){
	return static_cast<QSqlDatabase*>(ptr)->driver();
}

char* QSqlDatabase_DriverName(void* ptr){
	return static_cast<QSqlDatabase*>(ptr)->driverName().toUtf8().data();
}

char* QSqlDatabase_QSqlDatabase_Drivers(){
	return QSqlDatabase::drivers().join("|").toUtf8().data();
}

char* QSqlDatabase_HostName(void* ptr){
	return static_cast<QSqlDatabase*>(ptr)->hostName().toUtf8().data();
}

int QSqlDatabase_QSqlDatabase_IsDriverAvailable(char* name){
	return QSqlDatabase::isDriverAvailable(QString(name));
}

int QSqlDatabase_IsOpen(void* ptr){
	return static_cast<QSqlDatabase*>(ptr)->isOpen();
}

int QSqlDatabase_IsOpenError(void* ptr){
	return static_cast<QSqlDatabase*>(ptr)->isOpenError();
}

int QSqlDatabase_IsValid(void* ptr){
	return static_cast<QSqlDatabase*>(ptr)->isValid();
}

int QSqlDatabase_Open(void* ptr){
	return static_cast<QSqlDatabase*>(ptr)->open();
}

int QSqlDatabase_Open2(void* ptr, char* user, char* password){
	return static_cast<QSqlDatabase*>(ptr)->open(QString(user), QString(password));
}

char* QSqlDatabase_Password(void* ptr){
	return static_cast<QSqlDatabase*>(ptr)->password().toUtf8().data();
}

int QSqlDatabase_Port(void* ptr){
	return static_cast<QSqlDatabase*>(ptr)->port();
}

void QSqlDatabase_QSqlDatabase_RegisterSqlDriver(char* name, void* creator){
	QSqlDatabase::registerSqlDriver(QString(name), static_cast<QSqlDriverCreatorBase*>(creator));
}

void QSqlDatabase_QSqlDatabase_RemoveDatabase(char* connectionName){
	QSqlDatabase::removeDatabase(QString(connectionName));
}

int QSqlDatabase_Rollback(void* ptr){
	return static_cast<QSqlDatabase*>(ptr)->rollback();
}

void QSqlDatabase_SetConnectOptions(void* ptr, char* options){
	static_cast<QSqlDatabase*>(ptr)->setConnectOptions(QString(options));
}

void QSqlDatabase_SetDatabaseName(void* ptr, char* name){
	static_cast<QSqlDatabase*>(ptr)->setDatabaseName(QString(name));
}

void QSqlDatabase_SetHostName(void* ptr, char* host){
	static_cast<QSqlDatabase*>(ptr)->setHostName(QString(host));
}

void QSqlDatabase_SetPassword(void* ptr, char* password){
	static_cast<QSqlDatabase*>(ptr)->setPassword(QString(password));
}

void QSqlDatabase_SetPort(void* ptr, int port){
	static_cast<QSqlDatabase*>(ptr)->setPort(port);
}

void QSqlDatabase_SetUserName(void* ptr, char* name){
	static_cast<QSqlDatabase*>(ptr)->setUserName(QString(name));
}

int QSqlDatabase_Transaction(void* ptr){
	return static_cast<QSqlDatabase*>(ptr)->transaction();
}

char* QSqlDatabase_UserName(void* ptr){
	return static_cast<QSqlDatabase*>(ptr)->userName().toUtf8().data();
}

void QSqlDatabase_DestroyQSqlDatabase(void* ptr){
	static_cast<QSqlDatabase*>(ptr)->~QSqlDatabase();
}

