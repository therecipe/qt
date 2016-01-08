#define protected public

#include "sql.h"
#include "_cgo_export.h"

#include <QChildEvent>
#include <QEvent>
#include <QMetaObject>
#include <QModelIndex>
#include <QObject>
#include <QSqlDatabase>
#include <QSqlDriver>
#include <QSqlDriverCreator>
#include <QSqlDriverCreatorBase>
#include <QSqlDriverPlugin>
#include <QSqlError>
#include <QSqlField>
#include <QSqlIndex>
#include <QSqlQuery>
#include <QSqlQueryModel>
#include <QSqlRecord>
#include <QSqlRelation>
#include <QSqlRelationalTableModel>
#include <QSqlResult>
#include <QSqlTableModel>
#include <QString>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>
#include <QVariant>

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

class MyQSqlDriver: public QSqlDriver {
public:
	void Signal_Notification(const QString & name) { callbackQSqlDriverNotification(this, this->objectName().toUtf8().data(), name.toUtf8().data()); };
	void Signal_Notification2(const QString & name, QSqlDriver::NotificationSource source, const QVariant & payload) { callbackQSqlDriverNotification2(this, this->objectName().toUtf8().data(), name.toUtf8().data(), source, new QVariant(payload)); };
	void setOpen(bool open) { callbackQSqlDriverSetOpen(this, this->objectName().toUtf8().data(), open); };
	void setOpenError(bool error) { callbackQSqlDriverSetOpenError(this, this->objectName().toUtf8().data(), error); };
	void timerEvent(QTimerEvent * event) { callbackQSqlDriverTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQSqlDriverChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQSqlDriverCustomEvent(this, this->objectName().toUtf8().data(), event); };
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

void QSqlDriver_Notification(void* ptr, char* name){
	static_cast<QSqlDriver*>(ptr)->notification(QString(name));
}

void QSqlDriver_ConnectNotification2(void* ptr){
	QObject::connect(static_cast<QSqlDriver*>(ptr), static_cast<void (QSqlDriver::*)(const QString &, QSqlDriver::NotificationSource, const QVariant &)>(&QSqlDriver::notification), static_cast<MyQSqlDriver*>(ptr), static_cast<void (MyQSqlDriver::*)(const QString &, QSqlDriver::NotificationSource, const QVariant &)>(&MyQSqlDriver::Signal_Notification2));;
}

void QSqlDriver_DisconnectNotification2(void* ptr){
	QObject::disconnect(static_cast<QSqlDriver*>(ptr), static_cast<void (QSqlDriver::*)(const QString &, QSqlDriver::NotificationSource, const QVariant &)>(&QSqlDriver::notification), static_cast<MyQSqlDriver*>(ptr), static_cast<void (MyQSqlDriver::*)(const QString &, QSqlDriver::NotificationSource, const QVariant &)>(&MyQSqlDriver::Signal_Notification2));;
}

void QSqlDriver_Notification2(void* ptr, char* name, int source, void* payload){
	static_cast<QSqlDriver*>(ptr)->notification(QString(name), static_cast<QSqlDriver::NotificationSource>(source), *static_cast<QVariant*>(payload));
}

int QSqlDriver_Open(void* ptr, char* db, char* user, char* password, char* host, int port, char* options){
	return static_cast<QSqlDriver*>(ptr)->open(QString(db), QString(user), QString(password), QString(host), port, QString(options));
}

int QSqlDriver_RollbackTransaction(void* ptr){
	return static_cast<QSqlDriver*>(ptr)->rollbackTransaction();
}

void QSqlDriver_SetOpen(void* ptr, int open){
	static_cast<MyQSqlDriver*>(ptr)->setOpen(open != 0);
}

void QSqlDriver_SetOpenDefault(void* ptr, int open){
	static_cast<QSqlDriver*>(ptr)->QSqlDriver::setOpen(open != 0);
}

void QSqlDriver_SetOpenError(void* ptr, int error){
	static_cast<MyQSqlDriver*>(ptr)->setOpenError(error != 0);
}

void QSqlDriver_SetOpenErrorDefault(void* ptr, int error){
	static_cast<QSqlDriver*>(ptr)->QSqlDriver::setOpenError(error != 0);
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

void QSqlDriver_TimerEvent(void* ptr, void* event){
	static_cast<MyQSqlDriver*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QSqlDriver_TimerEventDefault(void* ptr, void* event){
	static_cast<QSqlDriver*>(ptr)->QSqlDriver::timerEvent(static_cast<QTimerEvent*>(event));
}

void QSqlDriver_ChildEvent(void* ptr, void* event){
	static_cast<MyQSqlDriver*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QSqlDriver_ChildEventDefault(void* ptr, void* event){
	static_cast<QSqlDriver*>(ptr)->QSqlDriver::childEvent(static_cast<QChildEvent*>(event));
}

void QSqlDriver_CustomEvent(void* ptr, void* event){
	static_cast<MyQSqlDriver*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QSqlDriver_CustomEventDefault(void* ptr, void* event){
	static_cast<QSqlDriver*>(ptr)->QSqlDriver::customEvent(static_cast<QEvent*>(event));
}

class MyQSqlDriverCreatorBase: public QSqlDriverCreatorBase {
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
};

void* QSqlDriverCreatorBase_CreateObject(void* ptr){
	return static_cast<QSqlDriverCreatorBase*>(ptr)->createObject();
}

void QSqlDriverCreatorBase_DestroyQSqlDriverCreatorBase(void* ptr){
	static_cast<QSqlDriverCreatorBase*>(ptr)->~QSqlDriverCreatorBase();
}

char* QSqlDriverCreatorBase_ObjectNameAbs(void* ptr){
	if (dynamic_cast<MyQSqlDriverCreatorBase*>(static_cast<QSqlDriverCreatorBase*>(ptr))) {
		return static_cast<MyQSqlDriverCreatorBase*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QSqlDriverCreatorBase_BASE").toUtf8().data();
}

void QSqlDriverCreatorBase_SetObjectNameAbs(void* ptr, char* name){
	if (dynamic_cast<MyQSqlDriverCreatorBase*>(static_cast<QSqlDriverCreatorBase*>(ptr))) {
		static_cast<MyQSqlDriverCreatorBase*>(ptr)->setObjectNameAbs(QString(name));
	}
}

void* QSqlDriverPlugin_Create(void* ptr, char* key){
	return static_cast<QSqlDriverPlugin*>(ptr)->create(QString(key));
}

void QSqlDriverPlugin_DestroyQSqlDriverPlugin(void* ptr){
	static_cast<QSqlDriverPlugin*>(ptr)->~QSqlDriverPlugin();
}

void QSqlDriverPlugin_TimerEvent(void* ptr, void* event){
	static_cast<QSqlDriverPlugin*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QSqlDriverPlugin_TimerEventDefault(void* ptr, void* event){
	static_cast<QSqlDriverPlugin*>(ptr)->QSqlDriverPlugin::timerEvent(static_cast<QTimerEvent*>(event));
}

void QSqlDriverPlugin_ChildEvent(void* ptr, void* event){
	static_cast<QSqlDriverPlugin*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QSqlDriverPlugin_ChildEventDefault(void* ptr, void* event){
	static_cast<QSqlDriverPlugin*>(ptr)->QSqlDriverPlugin::childEvent(static_cast<QChildEvent*>(event));
}

void QSqlDriverPlugin_CustomEvent(void* ptr, void* event){
	static_cast<QSqlDriverPlugin*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QSqlDriverPlugin_CustomEventDefault(void* ptr, void* event){
	static_cast<QSqlDriverPlugin*>(ptr)->QSqlDriverPlugin::customEvent(static_cast<QEvent*>(event));
}

void* QSqlError_NewQSqlError3(void* other){
	return new QSqlError(*static_cast<QSqlError*>(other));
}

void* QSqlError_NewQSqlError(char* driverText, char* databaseText, int ty, char* code){
	return new QSqlError(QString(driverText), QString(databaseText), static_cast<QSqlError::ErrorType>(ty), QString(code));
}

char* QSqlError_DatabaseText(void* ptr){
	return static_cast<QSqlError*>(ptr)->databaseText().toUtf8().data();
}

char* QSqlError_DriverText(void* ptr){
	return static_cast<QSqlError*>(ptr)->driverText().toUtf8().data();
}

int QSqlError_IsValid(void* ptr){
	return static_cast<QSqlError*>(ptr)->isValid();
}

char* QSqlError_NativeErrorCode(void* ptr){
	return static_cast<QSqlError*>(ptr)->nativeErrorCode().toUtf8().data();
}

char* QSqlError_Text(void* ptr){
	return static_cast<QSqlError*>(ptr)->text().toUtf8().data();
}

int QSqlError_Type(void* ptr){
	return static_cast<QSqlError*>(ptr)->type();
}

void QSqlError_DestroyQSqlError(void* ptr){
	static_cast<QSqlError*>(ptr)->~QSqlError();
}

void* QSqlField_NewQSqlField2(void* other){
	return new QSqlField(*static_cast<QSqlField*>(other));
}

void QSqlField_Clear(void* ptr){
	static_cast<QSqlField*>(ptr)->clear();
}

void* QSqlField_DefaultValue(void* ptr){
	return new QVariant(static_cast<QSqlField*>(ptr)->defaultValue());
}

int QSqlField_IsAutoValue(void* ptr){
	return static_cast<QSqlField*>(ptr)->isAutoValue();
}

int QSqlField_IsGenerated(void* ptr){
	return static_cast<QSqlField*>(ptr)->isGenerated();
}

int QSqlField_IsNull(void* ptr){
	return static_cast<QSqlField*>(ptr)->isNull();
}

int QSqlField_IsReadOnly(void* ptr){
	return static_cast<QSqlField*>(ptr)->isReadOnly();
}

int QSqlField_IsValid(void* ptr){
	return static_cast<QSqlField*>(ptr)->isValid();
}

int QSqlField_Length(void* ptr){
	return static_cast<QSqlField*>(ptr)->length();
}

char* QSqlField_Name(void* ptr){
	return static_cast<QSqlField*>(ptr)->name().toUtf8().data();
}

int QSqlField_Precision(void* ptr){
	return static_cast<QSqlField*>(ptr)->precision();
}

int QSqlField_RequiredStatus(void* ptr){
	return static_cast<QSqlField*>(ptr)->requiredStatus();
}

void QSqlField_SetAutoValue(void* ptr, int autoVal){
	static_cast<QSqlField*>(ptr)->setAutoValue(autoVal != 0);
}

void QSqlField_SetDefaultValue(void* ptr, void* value){
	static_cast<QSqlField*>(ptr)->setDefaultValue(*static_cast<QVariant*>(value));
}

void QSqlField_SetGenerated(void* ptr, int gen){
	static_cast<QSqlField*>(ptr)->setGenerated(gen != 0);
}

void QSqlField_SetLength(void* ptr, int fieldLength){
	static_cast<QSqlField*>(ptr)->setLength(fieldLength);
}

void QSqlField_SetName(void* ptr, char* name){
	static_cast<QSqlField*>(ptr)->setName(QString(name));
}

void QSqlField_SetPrecision(void* ptr, int precision){
	static_cast<QSqlField*>(ptr)->setPrecision(precision);
}

void QSqlField_SetReadOnly(void* ptr, int readOnly){
	static_cast<QSqlField*>(ptr)->setReadOnly(readOnly != 0);
}

void QSqlField_SetRequired(void* ptr, int required){
	static_cast<QSqlField*>(ptr)->setRequired(required != 0);
}

void QSqlField_SetRequiredStatus(void* ptr, int required){
	static_cast<QSqlField*>(ptr)->setRequiredStatus(static_cast<QSqlField::RequiredStatus>(required));
}

void QSqlField_SetValue(void* ptr, void* value){
	static_cast<QSqlField*>(ptr)->setValue(*static_cast<QVariant*>(value));
}

void* QSqlField_Value(void* ptr){
	return new QVariant(static_cast<QSqlField*>(ptr)->value());
}

void QSqlField_DestroyQSqlField(void* ptr){
	static_cast<QSqlField*>(ptr)->~QSqlField();
}

void* QSqlIndex_NewQSqlIndex2(void* other){
	return new QSqlIndex(*static_cast<QSqlIndex*>(other));
}

void* QSqlIndex_NewQSqlIndex(char* cursorname, char* name){
	return new QSqlIndex(QString(cursorname), QString(name));
}

void QSqlIndex_Append(void* ptr, void* field){
	static_cast<QSqlIndex*>(ptr)->append(*static_cast<QSqlField*>(field));
}

void QSqlIndex_Append2(void* ptr, void* field, int desc){
	static_cast<QSqlIndex*>(ptr)->append(*static_cast<QSqlField*>(field), desc != 0);
}

char* QSqlIndex_CursorName(void* ptr){
	return static_cast<QSqlIndex*>(ptr)->cursorName().toUtf8().data();
}

int QSqlIndex_IsDescending(void* ptr, int i){
	return static_cast<QSqlIndex*>(ptr)->isDescending(i);
}

char* QSqlIndex_Name(void* ptr){
	return static_cast<QSqlIndex*>(ptr)->name().toUtf8().data();
}

void QSqlIndex_SetCursorName(void* ptr, char* cursorName){
	static_cast<QSqlIndex*>(ptr)->setCursorName(QString(cursorName));
}

void QSqlIndex_SetDescending(void* ptr, int i, int desc){
	static_cast<QSqlIndex*>(ptr)->setDescending(i, desc != 0);
}

void QSqlIndex_SetName(void* ptr, char* name){
	static_cast<QSqlIndex*>(ptr)->setName(QString(name));
}

void QSqlIndex_DestroyQSqlIndex(void* ptr){
	static_cast<QSqlIndex*>(ptr)->~QSqlIndex();
}

void* QSqlQuery_NewQSqlQuery3(void* db){
	return new QSqlQuery(*static_cast<QSqlDatabase*>(db));
}

void* QSqlQuery_NewQSqlQuery(void* result){
	return new QSqlQuery(static_cast<QSqlResult*>(result));
}

void* QSqlQuery_NewQSqlQuery4(void* other){
	return new QSqlQuery(*static_cast<QSqlQuery*>(other));
}

void* QSqlQuery_NewQSqlQuery2(char* query, void* db){
	return new QSqlQuery(QString(query), *static_cast<QSqlDatabase*>(db));
}

int QSqlQuery_At(void* ptr){
	return static_cast<QSqlQuery*>(ptr)->at();
}

void* QSqlQuery_BoundValue(void* ptr, char* placeholder){
	return new QVariant(static_cast<QSqlQuery*>(ptr)->boundValue(QString(placeholder)));
}

void* QSqlQuery_BoundValue2(void* ptr, int pos){
	return new QVariant(static_cast<QSqlQuery*>(ptr)->boundValue(pos));
}

void QSqlQuery_Clear(void* ptr){
	static_cast<QSqlQuery*>(ptr)->clear();
}

void* QSqlQuery_Driver(void* ptr){
	return const_cast<QSqlDriver*>(static_cast<QSqlQuery*>(ptr)->driver());
}

int QSqlQuery_Exec2(void* ptr){
	return static_cast<QSqlQuery*>(ptr)->exec();
}

int QSqlQuery_Exec(void* ptr, char* query){
	return static_cast<QSqlQuery*>(ptr)->exec(QString(query));
}

int QSqlQuery_ExecBatch(void* ptr, int mode){
	return static_cast<QSqlQuery*>(ptr)->execBatch(static_cast<QSqlQuery::BatchExecutionMode>(mode));
}

char* QSqlQuery_ExecutedQuery(void* ptr){
	return static_cast<QSqlQuery*>(ptr)->executedQuery().toUtf8().data();
}

void QSqlQuery_Finish(void* ptr){
	static_cast<QSqlQuery*>(ptr)->finish();
}

int QSqlQuery_First(void* ptr){
	return static_cast<QSqlQuery*>(ptr)->first();
}

int QSqlQuery_IsActive(void* ptr){
	return static_cast<QSqlQuery*>(ptr)->isActive();
}

int QSqlQuery_IsForwardOnly(void* ptr){
	return static_cast<QSqlQuery*>(ptr)->isForwardOnly();
}

int QSqlQuery_IsNull2(void* ptr, char* name){
	return static_cast<QSqlQuery*>(ptr)->isNull(QString(name));
}

int QSqlQuery_IsNull(void* ptr, int field){
	return static_cast<QSqlQuery*>(ptr)->isNull(field);
}

int QSqlQuery_IsSelect(void* ptr){
	return static_cast<QSqlQuery*>(ptr)->isSelect();
}

int QSqlQuery_IsValid(void* ptr){
	return static_cast<QSqlQuery*>(ptr)->isValid();
}

int QSqlQuery_Last(void* ptr){
	return static_cast<QSqlQuery*>(ptr)->last();
}

void* QSqlQuery_LastInsertId(void* ptr){
	return new QVariant(static_cast<QSqlQuery*>(ptr)->lastInsertId());
}

char* QSqlQuery_LastQuery(void* ptr){
	return static_cast<QSqlQuery*>(ptr)->lastQuery().toUtf8().data();
}

int QSqlQuery_Next(void* ptr){
	return static_cast<QSqlQuery*>(ptr)->next();
}

int QSqlQuery_NextResult(void* ptr){
	return static_cast<QSqlQuery*>(ptr)->nextResult();
}

int QSqlQuery_NumRowsAffected(void* ptr){
	return static_cast<QSqlQuery*>(ptr)->numRowsAffected();
}

int QSqlQuery_Prepare(void* ptr, char* query){
	return static_cast<QSqlQuery*>(ptr)->prepare(QString(query));
}

int QSqlQuery_Previous(void* ptr){
	return static_cast<QSqlQuery*>(ptr)->previous();
}

void* QSqlQuery_Result(void* ptr){
	return const_cast<QSqlResult*>(static_cast<QSqlQuery*>(ptr)->result());
}

int QSqlQuery_Seek(void* ptr, int index, int relative){
	return static_cast<QSqlQuery*>(ptr)->seek(index, relative != 0);
}

void QSqlQuery_SetForwardOnly(void* ptr, int forward){
	static_cast<QSqlQuery*>(ptr)->setForwardOnly(forward != 0);
}

int QSqlQuery_Size(void* ptr){
	return static_cast<QSqlQuery*>(ptr)->size();
}

void* QSqlQuery_Value2(void* ptr, char* name){
	return new QVariant(static_cast<QSqlQuery*>(ptr)->value(QString(name)));
}

void* QSqlQuery_Value(void* ptr, int index){
	return new QVariant(static_cast<QSqlQuery*>(ptr)->value(index));
}

void QSqlQuery_DestroyQSqlQuery(void* ptr){
	static_cast<QSqlQuery*>(ptr)->~QSqlQuery();
}

class MyQSqlQueryModel: public QSqlQueryModel {
public:
	void clear() { callbackQSqlQueryModelClear(this, this->objectName().toUtf8().data()); };
	void fetchMore(const QModelIndex & parent) { callbackQSqlQueryModelFetchMore(this, this->objectName().toUtf8().data(), new QModelIndex(parent)); };
	void queryChange() { callbackQSqlQueryModelQueryChange(this, this->objectName().toUtf8().data()); };
	void revert() { if (!callbackQSqlQueryModelRevert(this, this->objectName().toUtf8().data())) { QSqlQueryModel::revert(); }; };
	void sort(int column, Qt::SortOrder order) { callbackQSqlQueryModelSort(this, this->objectName().toUtf8().data(), column, order); };
	void timerEvent(QTimerEvent * event) { callbackQSqlQueryModelTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQSqlQueryModelChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQSqlQueryModelCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

int QSqlQueryModel_RowCount(void* ptr, void* parent){
	return static_cast<QSqlQueryModel*>(ptr)->rowCount(*static_cast<QModelIndex*>(parent));
}

void* QSqlQueryModel_Data(void* ptr, void* item, int role){
	return new QVariant(static_cast<QSqlQueryModel*>(ptr)->data(*static_cast<QModelIndex*>(item), role));
}

int QSqlQueryModel_CanFetchMore(void* ptr, void* parent){
	return static_cast<QSqlQueryModel*>(ptr)->canFetchMore(*static_cast<QModelIndex*>(parent));
}

void QSqlQueryModel_Clear(void* ptr){
	static_cast<MyQSqlQueryModel*>(ptr)->clear();
}

void QSqlQueryModel_ClearDefault(void* ptr){
	static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::clear();
}

int QSqlQueryModel_ColumnCount(void* ptr, void* index){
	return static_cast<QSqlQueryModel*>(ptr)->columnCount(*static_cast<QModelIndex*>(index));
}

void QSqlQueryModel_FetchMore(void* ptr, void* parent){
	static_cast<MyQSqlQueryModel*>(ptr)->fetchMore(*static_cast<QModelIndex*>(parent));
}

void QSqlQueryModel_FetchMoreDefault(void* ptr, void* parent){
	static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::fetchMore(*static_cast<QModelIndex*>(parent));
}

void* QSqlQueryModel_HeaderData(void* ptr, int section, int orientation, int role){
	return new QVariant(static_cast<QSqlQueryModel*>(ptr)->headerData(section, static_cast<Qt::Orientation>(orientation), role));
}

void* QSqlQueryModel_IndexInQuery(void* ptr, void* item){
	return new QModelIndex(static_cast<QSqlQueryModel*>(ptr)->indexInQuery(*static_cast<QModelIndex*>(item)));
}

int QSqlQueryModel_InsertColumns(void* ptr, int column, int count, void* parent){
	return static_cast<QSqlQueryModel*>(ptr)->insertColumns(column, count, *static_cast<QModelIndex*>(parent));
}

void QSqlQueryModel_QueryChange(void* ptr){
	static_cast<MyQSqlQueryModel*>(ptr)->queryChange();
}

void QSqlQueryModel_QueryChangeDefault(void* ptr){
	static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::queryChange();
}

int QSqlQueryModel_RemoveColumns(void* ptr, int column, int count, void* parent){
	return static_cast<QSqlQueryModel*>(ptr)->removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

int QSqlQueryModel_SetHeaderData(void* ptr, int section, int orientation, void* value, int role){
	return static_cast<QSqlQueryModel*>(ptr)->setHeaderData(section, static_cast<Qt::Orientation>(orientation), *static_cast<QVariant*>(value), role);
}

void QSqlQueryModel_SetQuery(void* ptr, void* query){
	static_cast<QSqlQueryModel*>(ptr)->setQuery(*static_cast<QSqlQuery*>(query));
}

void QSqlQueryModel_SetQuery2(void* ptr, char* query, void* db){
	static_cast<QSqlQueryModel*>(ptr)->setQuery(QString(query), *static_cast<QSqlDatabase*>(db));
}

void QSqlQueryModel_DestroyQSqlQueryModel(void* ptr){
	static_cast<QSqlQueryModel*>(ptr)->~QSqlQueryModel();
}

void QSqlQueryModel_Revert(void* ptr){
	QMetaObject::invokeMethod(static_cast<MyQSqlQueryModel*>(ptr), "revert");
}

void QSqlQueryModel_RevertDefault(void* ptr){
	QMetaObject::invokeMethod(static_cast<QSqlQueryModel*>(ptr), "revert");
}

void QSqlQueryModel_Sort(void* ptr, int column, int order){
	static_cast<MyQSqlQueryModel*>(ptr)->sort(column, static_cast<Qt::SortOrder>(order));
}

void QSqlQueryModel_SortDefault(void* ptr, int column, int order){
	static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::sort(column, static_cast<Qt::SortOrder>(order));
}

void QSqlQueryModel_TimerEvent(void* ptr, void* event){
	static_cast<MyQSqlQueryModel*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QSqlQueryModel_TimerEventDefault(void* ptr, void* event){
	static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::timerEvent(static_cast<QTimerEvent*>(event));
}

void QSqlQueryModel_ChildEvent(void* ptr, void* event){
	static_cast<MyQSqlQueryModel*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QSqlQueryModel_ChildEventDefault(void* ptr, void* event){
	static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::childEvent(static_cast<QChildEvent*>(event));
}

void QSqlQueryModel_CustomEvent(void* ptr, void* event){
	static_cast<MyQSqlQueryModel*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QSqlQueryModel_CustomEventDefault(void* ptr, void* event){
	static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::customEvent(static_cast<QEvent*>(event));
}

void* QSqlRecord_NewQSqlRecord(){
	return new QSqlRecord();
}

void* QSqlRecord_NewQSqlRecord2(void* other){
	return new QSqlRecord(*static_cast<QSqlRecord*>(other));
}

void QSqlRecord_Append(void* ptr, void* field){
	static_cast<QSqlRecord*>(ptr)->append(*static_cast<QSqlField*>(field));
}

void QSqlRecord_Clear(void* ptr){
	static_cast<QSqlRecord*>(ptr)->clear();
}

void QSqlRecord_ClearValues(void* ptr){
	static_cast<QSqlRecord*>(ptr)->clearValues();
}

int QSqlRecord_Contains(void* ptr, char* name){
	return static_cast<QSqlRecord*>(ptr)->contains(QString(name));
}

int QSqlRecord_Count(void* ptr){
	return static_cast<QSqlRecord*>(ptr)->count();
}

char* QSqlRecord_FieldName(void* ptr, int index){
	return static_cast<QSqlRecord*>(ptr)->fieldName(index).toUtf8().data();
}

int QSqlRecord_IndexOf(void* ptr, char* name){
	return static_cast<QSqlRecord*>(ptr)->indexOf(QString(name));
}

int QSqlRecord_IsEmpty(void* ptr){
	return static_cast<QSqlRecord*>(ptr)->isEmpty();
}

int QSqlRecord_IsGenerated(void* ptr, char* name){
	return static_cast<QSqlRecord*>(ptr)->isGenerated(QString(name));
}

int QSqlRecord_IsGenerated2(void* ptr, int index){
	return static_cast<QSqlRecord*>(ptr)->isGenerated(index);
}

int QSqlRecord_IsNull(void* ptr, char* name){
	return static_cast<QSqlRecord*>(ptr)->isNull(QString(name));
}

int QSqlRecord_IsNull2(void* ptr, int index){
	return static_cast<QSqlRecord*>(ptr)->isNull(index);
}

void QSqlRecord_SetGenerated(void* ptr, char* name, int generated){
	static_cast<QSqlRecord*>(ptr)->setGenerated(QString(name), generated != 0);
}

void QSqlRecord_SetGenerated2(void* ptr, int index, int generated){
	static_cast<QSqlRecord*>(ptr)->setGenerated(index, generated != 0);
}

void QSqlRecord_SetNull2(void* ptr, char* name){
	static_cast<QSqlRecord*>(ptr)->setNull(QString(name));
}

void QSqlRecord_SetNull(void* ptr, int index){
	static_cast<QSqlRecord*>(ptr)->setNull(index);
}

void QSqlRecord_SetValue2(void* ptr, char* name, void* val){
	static_cast<QSqlRecord*>(ptr)->setValue(QString(name), *static_cast<QVariant*>(val));
}

void QSqlRecord_SetValue(void* ptr, int index, void* val){
	static_cast<QSqlRecord*>(ptr)->setValue(index, *static_cast<QVariant*>(val));
}

void* QSqlRecord_Value2(void* ptr, char* name){
	return new QVariant(static_cast<QSqlRecord*>(ptr)->value(QString(name)));
}

void* QSqlRecord_Value(void* ptr, int index){
	return new QVariant(static_cast<QSqlRecord*>(ptr)->value(index));
}

void QSqlRecord_DestroyQSqlRecord(void* ptr){
	static_cast<QSqlRecord*>(ptr)->~QSqlRecord();
}

void* QSqlRelation_NewQSqlRelation(){
	return new QSqlRelation();
}

void* QSqlRelation_NewQSqlRelation2(char* tableName, char* indexColumn, char* displayColumn){
	return new QSqlRelation(QString(tableName), QString(indexColumn), QString(displayColumn));
}

char* QSqlRelation_DisplayColumn(void* ptr){
	return static_cast<QSqlRelation*>(ptr)->displayColumn().toUtf8().data();
}

char* QSqlRelation_IndexColumn(void* ptr){
	return static_cast<QSqlRelation*>(ptr)->indexColumn().toUtf8().data();
}

int QSqlRelation_IsValid(void* ptr){
	return static_cast<QSqlRelation*>(ptr)->isValid();
}

char* QSqlRelation_TableName(void* ptr){
	return static_cast<QSqlRelation*>(ptr)->tableName().toUtf8().data();
}

class MyQSqlRelationalTableModel: public QSqlRelationalTableModel {
public:
	void clear() { callbackQSqlRelationalTableModelClear(this, this->objectName().toUtf8().data()); };
	void revertRow(int row) { if (!callbackQSqlRelationalTableModelRevertRow(this, this->objectName().toUtf8().data(), row)) { QSqlRelationalTableModel::revertRow(row); }; };
	void setTable(const QString & table) { callbackQSqlRelationalTableModelSetTable(this, this->objectName().toUtf8().data(), table.toUtf8().data()); };
	void revert() { if (!callbackQSqlRelationalTableModelRevert(this, this->objectName().toUtf8().data())) { QSqlRelationalTableModel::revert(); }; };
	void setEditStrategy(QSqlRelationalTableModel::EditStrategy strategy) { callbackQSqlRelationalTableModelSetEditStrategy(this, this->objectName().toUtf8().data(), strategy); };
	void setFilter(const QString & filter) { callbackQSqlRelationalTableModelSetFilter(this, this->objectName().toUtf8().data(), filter.toUtf8().data()); };
	void setSort(int column, Qt::SortOrder order) { callbackQSqlRelationalTableModelSetSort(this, this->objectName().toUtf8().data(), column, order); };
	void sort(int column, Qt::SortOrder order) { callbackQSqlRelationalTableModelSort(this, this->objectName().toUtf8().data(), column, order); };
	void fetchMore(const QModelIndex & parent) { callbackQSqlRelationalTableModelFetchMore(this, this->objectName().toUtf8().data(), new QModelIndex(parent)); };
	void queryChange() { callbackQSqlRelationalTableModelQueryChange(this, this->objectName().toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQSqlRelationalTableModelTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQSqlRelationalTableModelChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQSqlRelationalTableModelCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void QSqlRelationalTableModel_Clear(void* ptr){
	static_cast<MyQSqlRelationalTableModel*>(ptr)->clear();
}

void QSqlRelationalTableModel_ClearDefault(void* ptr){
	static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::clear();
}

void* QSqlRelationalTableModel_Data(void* ptr, void* index, int role){
	return new QVariant(static_cast<QSqlRelationalTableModel*>(ptr)->data(*static_cast<QModelIndex*>(index), role));
}

int QSqlRelationalTableModel_InsertRowIntoTable(void* ptr, void* values){
	return static_cast<QSqlRelationalTableModel*>(ptr)->insertRowIntoTable(*static_cast<QSqlRecord*>(values));
}

char* QSqlRelationalTableModel_OrderByClause(void* ptr){
	return static_cast<QSqlRelationalTableModel*>(ptr)->orderByClause().toUtf8().data();
}

void* QSqlRelationalTableModel_RelationModel(void* ptr, int column){
	return static_cast<QSqlRelationalTableModel*>(ptr)->relationModel(column);
}

int QSqlRelationalTableModel_RemoveColumns(void* ptr, int column, int count, void* parent){
	return static_cast<QSqlRelationalTableModel*>(ptr)->removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

void QSqlRelationalTableModel_RevertRow(void* ptr, int row){
	QMetaObject::invokeMethod(static_cast<MyQSqlRelationalTableModel*>(ptr), "revertRow", Q_ARG(int, row));
}

void QSqlRelationalTableModel_RevertRowDefault(void* ptr, int row){
	QMetaObject::invokeMethod(static_cast<QSqlRelationalTableModel*>(ptr), "revertRow", Q_ARG(int, row));
}

int QSqlRelationalTableModel_Select(void* ptr){
	return static_cast<QSqlRelationalTableModel*>(ptr)->select();
}

char* QSqlRelationalTableModel_SelectStatement(void* ptr){
	return static_cast<QSqlRelationalTableModel*>(ptr)->selectStatement().toUtf8().data();
}

int QSqlRelationalTableModel_SetData(void* ptr, void* index, void* value, int role){
	return static_cast<QSqlRelationalTableModel*>(ptr)->setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
}

void QSqlRelationalTableModel_SetJoinMode(void* ptr, int joinMode){
	static_cast<QSqlRelationalTableModel*>(ptr)->setJoinMode(static_cast<QSqlRelationalTableModel::JoinMode>(joinMode));
}

void QSqlRelationalTableModel_SetTable(void* ptr, char* table){
	static_cast<MyQSqlRelationalTableModel*>(ptr)->setTable(QString(table));
}

void QSqlRelationalTableModel_SetTableDefault(void* ptr, char* table){
	static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::setTable(QString(table));
}

int QSqlRelationalTableModel_UpdateRowInTable(void* ptr, int row, void* values){
	return static_cast<QSqlRelationalTableModel*>(ptr)->updateRowInTable(row, *static_cast<QSqlRecord*>(values));
}

void QSqlRelationalTableModel_DestroyQSqlRelationalTableModel(void* ptr){
	static_cast<QSqlRelationalTableModel*>(ptr)->~QSqlRelationalTableModel();
}

void QSqlRelationalTableModel_Revert(void* ptr){
	QMetaObject::invokeMethod(static_cast<MyQSqlRelationalTableModel*>(ptr), "revert");
}

void QSqlRelationalTableModel_RevertDefault(void* ptr){
	QMetaObject::invokeMethod(static_cast<QSqlRelationalTableModel*>(ptr), "revert");
}

void QSqlRelationalTableModel_SetEditStrategy(void* ptr, int strategy){
	static_cast<MyQSqlRelationalTableModel*>(ptr)->setEditStrategy(static_cast<QSqlRelationalTableModel::EditStrategy>(strategy));
}

void QSqlRelationalTableModel_SetEditStrategyDefault(void* ptr, int strategy){
	static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::setEditStrategy(static_cast<QSqlRelationalTableModel::EditStrategy>(strategy));
}

void QSqlRelationalTableModel_SetFilter(void* ptr, char* filter){
	static_cast<MyQSqlRelationalTableModel*>(ptr)->setFilter(QString(filter));
}

void QSqlRelationalTableModel_SetFilterDefault(void* ptr, char* filter){
	static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::setFilter(QString(filter));
}

void QSqlRelationalTableModel_SetSort(void* ptr, int column, int order){
	static_cast<MyQSqlRelationalTableModel*>(ptr)->setSort(column, static_cast<Qt::SortOrder>(order));
}

void QSqlRelationalTableModel_SetSortDefault(void* ptr, int column, int order){
	static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::setSort(column, static_cast<Qt::SortOrder>(order));
}

void QSqlRelationalTableModel_Sort(void* ptr, int column, int order){
	static_cast<MyQSqlRelationalTableModel*>(ptr)->sort(column, static_cast<Qt::SortOrder>(order));
}

void QSqlRelationalTableModel_SortDefault(void* ptr, int column, int order){
	static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::sort(column, static_cast<Qt::SortOrder>(order));
}

void QSqlRelationalTableModel_FetchMore(void* ptr, void* parent){
	static_cast<MyQSqlRelationalTableModel*>(ptr)->fetchMore(*static_cast<QModelIndex*>(parent));
}

void QSqlRelationalTableModel_FetchMoreDefault(void* ptr, void* parent){
	static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::fetchMore(*static_cast<QModelIndex*>(parent));
}

void QSqlRelationalTableModel_QueryChange(void* ptr){
	static_cast<MyQSqlRelationalTableModel*>(ptr)->queryChange();
}

void QSqlRelationalTableModel_QueryChangeDefault(void* ptr){
	static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::queryChange();
}

void QSqlRelationalTableModel_TimerEvent(void* ptr, void* event){
	static_cast<MyQSqlRelationalTableModel*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QSqlRelationalTableModel_TimerEventDefault(void* ptr, void* event){
	static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::timerEvent(static_cast<QTimerEvent*>(event));
}

void QSqlRelationalTableModel_ChildEvent(void* ptr, void* event){
	static_cast<MyQSqlRelationalTableModel*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QSqlRelationalTableModel_ChildEventDefault(void* ptr, void* event){
	static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::childEvent(static_cast<QChildEvent*>(event));
}

void QSqlRelationalTableModel_CustomEvent(void* ptr, void* event){
	static_cast<MyQSqlRelationalTableModel*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QSqlRelationalTableModel_CustomEventDefault(void* ptr, void* event){
	static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::customEvent(static_cast<QEvent*>(event));
}

class MyQSqlResult: public QSqlResult {
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	void setActive(bool active) { callbackQSqlResultSetActive(this, this->objectNameAbs().toUtf8().data(), active); };
	void setAt(int index) { callbackQSqlResultSetAt(this, this->objectNameAbs().toUtf8().data(), index); };
	void setForwardOnly(bool forward) { callbackQSqlResultSetForwardOnly(this, this->objectNameAbs().toUtf8().data(), forward); };
	void setQuery(const QString & query) { callbackQSqlResultSetQuery(this, this->objectNameAbs().toUtf8().data(), query.toUtf8().data()); };
};

int QSqlResult_Exec(void* ptr){
	return static_cast<QSqlResult*>(ptr)->exec();
}

int QSqlResult_FetchNext(void* ptr){
	return static_cast<QSqlResult*>(ptr)->fetchNext();
}

int QSqlResult_FetchPrevious(void* ptr){
	return static_cast<QSqlResult*>(ptr)->fetchPrevious();
}

void* QSqlResult_Handle(void* ptr){
	return new QVariant(static_cast<QSqlResult*>(ptr)->handle());
}

void* QSqlResult_LastInsertId(void* ptr){
	return new QVariant(static_cast<QSqlResult*>(ptr)->lastInsertId());
}

int QSqlResult_Prepare(void* ptr, char* query){
	return static_cast<QSqlResult*>(ptr)->prepare(QString(query));
}

int QSqlResult_SavePrepare(void* ptr, char* query){
	return static_cast<QSqlResult*>(ptr)->savePrepare(QString(query));
}

void QSqlResult_SetActive(void* ptr, int active){
	static_cast<MyQSqlResult*>(ptr)->setActive(active != 0);
}

void QSqlResult_SetActiveDefault(void* ptr, int active){
	static_cast<QSqlResult*>(ptr)->QSqlResult::setActive(active != 0);
}

void QSqlResult_SetAt(void* ptr, int index){
	static_cast<MyQSqlResult*>(ptr)->setAt(index);
}

void QSqlResult_SetAtDefault(void* ptr, int index){
	static_cast<QSqlResult*>(ptr)->QSqlResult::setAt(index);
}

void QSqlResult_SetForwardOnly(void* ptr, int forward){
	static_cast<MyQSqlResult*>(ptr)->setForwardOnly(forward != 0);
}

void QSqlResult_SetForwardOnlyDefault(void* ptr, int forward){
	static_cast<QSqlResult*>(ptr)->QSqlResult::setForwardOnly(forward != 0);
}

void QSqlResult_SetQuery(void* ptr, char* query){
	static_cast<MyQSqlResult*>(ptr)->setQuery(QString(query));
}

void QSqlResult_SetQueryDefault(void* ptr, char* query){
	static_cast<QSqlResult*>(ptr)->QSqlResult::setQuery(QString(query));
}

void QSqlResult_DestroyQSqlResult(void* ptr){
	static_cast<QSqlResult*>(ptr)->~QSqlResult();
}

char* QSqlResult_ObjectNameAbs(void* ptr){
	if (dynamic_cast<MyQSqlResult*>(static_cast<QSqlResult*>(ptr))) {
		return static_cast<MyQSqlResult*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QSqlResult_BASE").toUtf8().data();
}

void QSqlResult_SetObjectNameAbs(void* ptr, char* name){
	if (dynamic_cast<MyQSqlResult*>(static_cast<QSqlResult*>(ptr))) {
		static_cast<MyQSqlResult*>(ptr)->setObjectNameAbs(QString(name));
	}
}

class MyQSqlTableModel: public QSqlTableModel {
public:
	void Signal_BeforeDelete(int row) { callbackQSqlTableModelBeforeDelete(this, this->objectName().toUtf8().data(), row); };
	void clear() { callbackQSqlTableModelClear(this, this->objectName().toUtf8().data()); };
	void revert() { if (!callbackQSqlTableModelRevert(this, this->objectName().toUtf8().data())) { QSqlTableModel::revert(); }; };
	void revertRow(int row) { callbackQSqlTableModelRevertRow(this, this->objectName().toUtf8().data(), row); };
	void setEditStrategy(QSqlTableModel::EditStrategy strategy) { callbackQSqlTableModelSetEditStrategy(this, this->objectName().toUtf8().data(), strategy); };
	void setFilter(const QString & filter) { callbackQSqlTableModelSetFilter(this, this->objectName().toUtf8().data(), filter.toUtf8().data()); };
	void setSort(int column, Qt::SortOrder order) { callbackQSqlTableModelSetSort(this, this->objectName().toUtf8().data(), column, order); };
	void setTable(const QString & tableName) { callbackQSqlTableModelSetTable(this, this->objectName().toUtf8().data(), tableName.toUtf8().data()); };
	void sort(int column, Qt::SortOrder order) { callbackQSqlTableModelSort(this, this->objectName().toUtf8().data(), column, order); };
	void fetchMore(const QModelIndex & parent) { callbackQSqlTableModelFetchMore(this, this->objectName().toUtf8().data(), new QModelIndex(parent)); };
	void queryChange() { callbackQSqlTableModelQueryChange(this, this->objectName().toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQSqlTableModelTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQSqlTableModelChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQSqlTableModelCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void QSqlTableModel_ConnectBeforeDelete(void* ptr){
	QObject::connect(static_cast<QSqlTableModel*>(ptr), static_cast<void (QSqlTableModel::*)(int)>(&QSqlTableModel::beforeDelete), static_cast<MyQSqlTableModel*>(ptr), static_cast<void (MyQSqlTableModel::*)(int)>(&MyQSqlTableModel::Signal_BeforeDelete));;
}

void QSqlTableModel_DisconnectBeforeDelete(void* ptr){
	QObject::disconnect(static_cast<QSqlTableModel*>(ptr), static_cast<void (QSqlTableModel::*)(int)>(&QSqlTableModel::beforeDelete), static_cast<MyQSqlTableModel*>(ptr), static_cast<void (MyQSqlTableModel::*)(int)>(&MyQSqlTableModel::Signal_BeforeDelete));;
}

void QSqlTableModel_BeforeDelete(void* ptr, int row){
	static_cast<QSqlTableModel*>(ptr)->beforeDelete(row);
}

void QSqlTableModel_Clear(void* ptr){
	static_cast<MyQSqlTableModel*>(ptr)->clear();
}

void QSqlTableModel_ClearDefault(void* ptr){
	static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::clear();
}

void* QSqlTableModel_Data(void* ptr, void* index, int role){
	return new QVariant(static_cast<QSqlTableModel*>(ptr)->data(*static_cast<QModelIndex*>(index), role));
}

int QSqlTableModel_DeleteRowFromTable(void* ptr, int row){
	return static_cast<QSqlTableModel*>(ptr)->deleteRowFromTable(row);
}

int QSqlTableModel_EditStrategy(void* ptr){
	return static_cast<QSqlTableModel*>(ptr)->editStrategy();
}

int QSqlTableModel_FieldIndex(void* ptr, char* fieldName){
	return static_cast<QSqlTableModel*>(ptr)->fieldIndex(QString(fieldName));
}

char* QSqlTableModel_Filter(void* ptr){
	return static_cast<QSqlTableModel*>(ptr)->filter().toUtf8().data();
}

int QSqlTableModel_Flags(void* ptr, void* index){
	return static_cast<QSqlTableModel*>(ptr)->flags(*static_cast<QModelIndex*>(index));
}

void* QSqlTableModel_HeaderData(void* ptr, int section, int orientation, int role){
	return new QVariant(static_cast<QSqlTableModel*>(ptr)->headerData(section, static_cast<Qt::Orientation>(orientation), role));
}

void* QSqlTableModel_IndexInQuery(void* ptr, void* item){
	return new QModelIndex(static_cast<QSqlTableModel*>(ptr)->indexInQuery(*static_cast<QModelIndex*>(item)));
}

int QSqlTableModel_InsertRecord(void* ptr, int row, void* record){
	return static_cast<QSqlTableModel*>(ptr)->insertRecord(row, *static_cast<QSqlRecord*>(record));
}

int QSqlTableModel_InsertRowIntoTable(void* ptr, void* values){
	return static_cast<QSqlTableModel*>(ptr)->insertRowIntoTable(*static_cast<QSqlRecord*>(values));
}

int QSqlTableModel_InsertRows(void* ptr, int row, int count, void* parent){
	return static_cast<QSqlTableModel*>(ptr)->insertRows(row, count, *static_cast<QModelIndex*>(parent));
}

int QSqlTableModel_IsDirty2(void* ptr){
	return static_cast<QSqlTableModel*>(ptr)->isDirty();
}

int QSqlTableModel_IsDirty(void* ptr, void* index){
	return static_cast<QSqlTableModel*>(ptr)->isDirty(*static_cast<QModelIndex*>(index));
}

char* QSqlTableModel_OrderByClause(void* ptr){
	return static_cast<QSqlTableModel*>(ptr)->orderByClause().toUtf8().data();
}

int QSqlTableModel_RemoveColumns(void* ptr, int column, int count, void* parent){
	return static_cast<QSqlTableModel*>(ptr)->removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

int QSqlTableModel_RemoveRows(void* ptr, int row, int count, void* parent){
	return static_cast<QSqlTableModel*>(ptr)->removeRows(row, count, *static_cast<QModelIndex*>(parent));
}

void QSqlTableModel_Revert(void* ptr){
	QMetaObject::invokeMethod(static_cast<MyQSqlTableModel*>(ptr), "revert");
}

void QSqlTableModel_RevertDefault(void* ptr){
	QMetaObject::invokeMethod(static_cast<QSqlTableModel*>(ptr), "revert");
}

void QSqlTableModel_RevertAll(void* ptr){
	QMetaObject::invokeMethod(static_cast<QSqlTableModel*>(ptr), "revertAll");
}

void QSqlTableModel_RevertRow(void* ptr, int row){
	static_cast<MyQSqlTableModel*>(ptr)->revertRow(row);
}

void QSqlTableModel_RevertRowDefault(void* ptr, int row){
	static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::revertRow(row);
}

int QSqlTableModel_RowCount(void* ptr, void* parent){
	return static_cast<QSqlTableModel*>(ptr)->rowCount(*static_cast<QModelIndex*>(parent));
}

int QSqlTableModel_Select(void* ptr){
	return QMetaObject::invokeMethod(static_cast<QSqlTableModel*>(ptr), "select");
}

int QSqlTableModel_SelectRow(void* ptr, int row){
	return QMetaObject::invokeMethod(static_cast<QSqlTableModel*>(ptr), "selectRow", Q_ARG(int, row));
}

char* QSqlTableModel_SelectStatement(void* ptr){
	return static_cast<QSqlTableModel*>(ptr)->selectStatement().toUtf8().data();
}

int QSqlTableModel_SetData(void* ptr, void* index, void* value, int role){
	return static_cast<QSqlTableModel*>(ptr)->setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
}

void QSqlTableModel_SetEditStrategy(void* ptr, int strategy){
	static_cast<MyQSqlTableModel*>(ptr)->setEditStrategy(static_cast<QSqlTableModel::EditStrategy>(strategy));
}

void QSqlTableModel_SetEditStrategyDefault(void* ptr, int strategy){
	static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::setEditStrategy(static_cast<QSqlTableModel::EditStrategy>(strategy));
}

void QSqlTableModel_SetFilter(void* ptr, char* filter){
	static_cast<MyQSqlTableModel*>(ptr)->setFilter(QString(filter));
}

void QSqlTableModel_SetFilterDefault(void* ptr, char* filter){
	static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::setFilter(QString(filter));
}

int QSqlTableModel_SetRecord(void* ptr, int row, void* values){
	return static_cast<QSqlTableModel*>(ptr)->setRecord(row, *static_cast<QSqlRecord*>(values));
}

void QSqlTableModel_SetSort(void* ptr, int column, int order){
	static_cast<MyQSqlTableModel*>(ptr)->setSort(column, static_cast<Qt::SortOrder>(order));
}

void QSqlTableModel_SetSortDefault(void* ptr, int column, int order){
	static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::setSort(column, static_cast<Qt::SortOrder>(order));
}

void QSqlTableModel_SetTable(void* ptr, char* tableName){
	static_cast<MyQSqlTableModel*>(ptr)->setTable(QString(tableName));
}

void QSqlTableModel_SetTableDefault(void* ptr, char* tableName){
	static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::setTable(QString(tableName));
}

void QSqlTableModel_Sort(void* ptr, int column, int order){
	static_cast<MyQSqlTableModel*>(ptr)->sort(column, static_cast<Qt::SortOrder>(order));
}

void QSqlTableModel_SortDefault(void* ptr, int column, int order){
	static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::sort(column, static_cast<Qt::SortOrder>(order));
}

int QSqlTableModel_Submit(void* ptr){
	return QMetaObject::invokeMethod(static_cast<QSqlTableModel*>(ptr), "submit");
}

int QSqlTableModel_SubmitAll(void* ptr){
	return QMetaObject::invokeMethod(static_cast<QSqlTableModel*>(ptr), "submitAll");
}

char* QSqlTableModel_TableName(void* ptr){
	return static_cast<QSqlTableModel*>(ptr)->tableName().toUtf8().data();
}

int QSqlTableModel_UpdateRowInTable(void* ptr, int row, void* values){
	return static_cast<QSqlTableModel*>(ptr)->updateRowInTable(row, *static_cast<QSqlRecord*>(values));
}

void QSqlTableModel_DestroyQSqlTableModel(void* ptr){
	static_cast<QSqlTableModel*>(ptr)->~QSqlTableModel();
}

void QSqlTableModel_FetchMore(void* ptr, void* parent){
	static_cast<MyQSqlTableModel*>(ptr)->fetchMore(*static_cast<QModelIndex*>(parent));
}

void QSqlTableModel_FetchMoreDefault(void* ptr, void* parent){
	static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::fetchMore(*static_cast<QModelIndex*>(parent));
}

void QSqlTableModel_QueryChange(void* ptr){
	static_cast<MyQSqlTableModel*>(ptr)->queryChange();
}

void QSqlTableModel_QueryChangeDefault(void* ptr){
	static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::queryChange();
}

void QSqlTableModel_TimerEvent(void* ptr, void* event){
	static_cast<MyQSqlTableModel*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QSqlTableModel_TimerEventDefault(void* ptr, void* event){
	static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::timerEvent(static_cast<QTimerEvent*>(event));
}

void QSqlTableModel_ChildEvent(void* ptr, void* event){
	static_cast<MyQSqlTableModel*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QSqlTableModel_ChildEventDefault(void* ptr, void* event){
	static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::childEvent(static_cast<QChildEvent*>(event));
}

void QSqlTableModel_CustomEvent(void* ptr, void* event){
	static_cast<MyQSqlTableModel*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QSqlTableModel_CustomEventDefault(void* ptr, void* event){
	static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::customEvent(static_cast<QEvent*>(event));
}

