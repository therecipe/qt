#include "sql.h"
#include "_cgo_export.h"

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
#include <QVariant>

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
	return QSqlDatabase::connectionNames().join(",,,").toUtf8().data();
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
	return QSqlDatabase::drivers().join(",,,").toUtf8().data();
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
	return static_cast<QSqlDriver*>(ptr)->subscribedToNotifications().join(",,,").toUtf8().data();
}

int QSqlDriver_UnsubscribeFromNotification(void* ptr, char* name){
	return static_cast<QSqlDriver*>(ptr)->unsubscribeFromNotification(QString(name));
}

void QSqlDriver_DestroyQSqlDriver(void* ptr){
	static_cast<QSqlDriver*>(ptr)->~QSqlDriver();
}

class MyQSqlDriverCreatorBase: public QSqlDriverCreatorBase {
public:
};

void* QSqlDriverCreatorBase_CreateObject(void* ptr){
	return static_cast<QSqlDriverCreatorBase*>(ptr)->createObject();
}

void QSqlDriverCreatorBase_DestroyQSqlDriverCreatorBase(void* ptr){
	static_cast<QSqlDriverCreatorBase*>(ptr)->~QSqlDriverCreatorBase();
}

class MyQSqlDriverPlugin: public QSqlDriverPlugin {
public:
};

void* QSqlDriverPlugin_Create(void* ptr, char* key){
	return static_cast<QSqlDriverPlugin*>(ptr)->create(QString(key));
}

void QSqlDriverPlugin_DestroyQSqlDriverPlugin(void* ptr){
	static_cast<QSqlDriverPlugin*>(ptr)->~QSqlDriverPlugin();
}

class MyQSqlError: public QSqlError {
public:
};

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

class MyQSqlField: public QSqlField {
public:
};

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

class MyQSqlIndex: public QSqlIndex {
public:
};

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

class MyQSqlQuery: public QSqlQuery {
public:
};

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
	static_cast<QSqlQueryModel*>(ptr)->clear();
}

int QSqlQueryModel_ColumnCount(void* ptr, void* index){
	return static_cast<QSqlQueryModel*>(ptr)->columnCount(*static_cast<QModelIndex*>(index));
}

void QSqlQueryModel_FetchMore(void* ptr, void* parent){
	static_cast<QSqlQueryModel*>(ptr)->fetchMore(*static_cast<QModelIndex*>(parent));
}

void* QSqlQueryModel_HeaderData(void* ptr, int section, int orientation, int role){
	return new QVariant(static_cast<QSqlQueryModel*>(ptr)->headerData(section, static_cast<Qt::Orientation>(orientation), role));
}

int QSqlQueryModel_InsertColumns(void* ptr, int column, int count, void* parent){
	return static_cast<QSqlQueryModel*>(ptr)->insertColumns(column, count, *static_cast<QModelIndex*>(parent));
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

class MyQSqlRecord: public QSqlRecord {
public:
};

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

class MyQSqlRelation: public QSqlRelation {
public:
};

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
};

void QSqlRelationalTableModel_Clear(void* ptr){
	static_cast<QSqlRelationalTableModel*>(ptr)->clear();
}

void* QSqlRelationalTableModel_Data(void* ptr, void* index, int role){
	return new QVariant(static_cast<QSqlRelationalTableModel*>(ptr)->data(*static_cast<QModelIndex*>(index), role));
}

void* QSqlRelationalTableModel_RelationModel(void* ptr, int column){
	return static_cast<QSqlRelationalTableModel*>(ptr)->relationModel(column);
}

int QSqlRelationalTableModel_RemoveColumns(void* ptr, int column, int count, void* parent){
	return static_cast<QSqlRelationalTableModel*>(ptr)->removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

void QSqlRelationalTableModel_RevertRow(void* ptr, int row){
	QMetaObject::invokeMethod(static_cast<QSqlRelationalTableModel*>(ptr), "revertRow", Q_ARG(int, row));
}

int QSqlRelationalTableModel_Select(void* ptr){
	return static_cast<QSqlRelationalTableModel*>(ptr)->select();
}

int QSqlRelationalTableModel_SetData(void* ptr, void* index, void* value, int role){
	return static_cast<QSqlRelationalTableModel*>(ptr)->setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
}

void QSqlRelationalTableModel_SetJoinMode(void* ptr, int joinMode){
	static_cast<QSqlRelationalTableModel*>(ptr)->setJoinMode(static_cast<QSqlRelationalTableModel::JoinMode>(joinMode));
}

void QSqlRelationalTableModel_SetRelation(void* ptr, int column, void* relation){
	static_cast<QSqlRelationalTableModel*>(ptr)->setRelation(column, *static_cast<QSqlRelation*>(relation));
}

void QSqlRelationalTableModel_SetTable(void* ptr, char* table){
	static_cast<QSqlRelationalTableModel*>(ptr)->setTable(QString(table));
}

void QSqlRelationalTableModel_DestroyQSqlRelationalTableModel(void* ptr){
	static_cast<QSqlRelationalTableModel*>(ptr)->~QSqlRelationalTableModel();
}

class MyQSqlResult: public QSqlResult {
public:
};

void* QSqlResult_Handle(void* ptr){
	return new QVariant(static_cast<QSqlResult*>(ptr)->handle());
}

void QSqlResult_DestroyQSqlResult(void* ptr){
	static_cast<QSqlResult*>(ptr)->~QSqlResult();
}

class MyQSqlTableModel: public QSqlTableModel {
public:
void Signal_BeforeDelete(int row){callbackQSqlTableModelBeforeDelete(this->objectName().toUtf8().data(), row);};
};

void QSqlTableModel_ConnectBeforeDelete(void* ptr){
	QObject::connect(static_cast<QSqlTableModel*>(ptr), static_cast<void (QSqlTableModel::*)(int)>(&QSqlTableModel::beforeDelete), static_cast<MyQSqlTableModel*>(ptr), static_cast<void (MyQSqlTableModel::*)(int)>(&MyQSqlTableModel::Signal_BeforeDelete));;
}

void QSqlTableModel_DisconnectBeforeDelete(void* ptr){
	QObject::disconnect(static_cast<QSqlTableModel*>(ptr), static_cast<void (QSqlTableModel::*)(int)>(&QSqlTableModel::beforeDelete), static_cast<MyQSqlTableModel*>(ptr), static_cast<void (MyQSqlTableModel::*)(int)>(&MyQSqlTableModel::Signal_BeforeDelete));;
}

void QSqlTableModel_Clear(void* ptr){
	static_cast<QSqlTableModel*>(ptr)->clear();
}

void* QSqlTableModel_Data(void* ptr, void* index, int role){
	return new QVariant(static_cast<QSqlTableModel*>(ptr)->data(*static_cast<QModelIndex*>(index), role));
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

int QSqlTableModel_InsertRecord(void* ptr, int row, void* record){
	return static_cast<QSqlTableModel*>(ptr)->insertRecord(row, *static_cast<QSqlRecord*>(record));
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

int QSqlTableModel_RemoveColumns(void* ptr, int column, int count, void* parent){
	return static_cast<QSqlTableModel*>(ptr)->removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

int QSqlTableModel_RemoveRows(void* ptr, int row, int count, void* parent){
	return static_cast<QSqlTableModel*>(ptr)->removeRows(row, count, *static_cast<QModelIndex*>(parent));
}

void QSqlTableModel_Revert(void* ptr){
	QMetaObject::invokeMethod(static_cast<QSqlTableModel*>(ptr), "revert");
}

void QSqlTableModel_RevertAll(void* ptr){
	QMetaObject::invokeMethod(static_cast<QSqlTableModel*>(ptr), "revertAll");
}

void QSqlTableModel_RevertRow(void* ptr, int row){
	static_cast<QSqlTableModel*>(ptr)->revertRow(row);
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

int QSqlTableModel_SetData(void* ptr, void* index, void* value, int role){
	return static_cast<QSqlTableModel*>(ptr)->setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
}

void QSqlTableModel_SetEditStrategy(void* ptr, int strategy){
	static_cast<QSqlTableModel*>(ptr)->setEditStrategy(static_cast<QSqlTableModel::EditStrategy>(strategy));
}

void QSqlTableModel_SetFilter(void* ptr, char* filter){
	static_cast<QSqlTableModel*>(ptr)->setFilter(QString(filter));
}

int QSqlTableModel_SetRecord(void* ptr, int row, void* values){
	return static_cast<QSqlTableModel*>(ptr)->setRecord(row, *static_cast<QSqlRecord*>(values));
}

void QSqlTableModel_SetSort(void* ptr, int column, int order){
	static_cast<QSqlTableModel*>(ptr)->setSort(column, static_cast<Qt::SortOrder>(order));
}

void QSqlTableModel_SetTable(void* ptr, char* tableName){
	static_cast<QSqlTableModel*>(ptr)->setTable(QString(tableName));
}

void QSqlTableModel_Sort(void* ptr, int column, int order){
	static_cast<QSqlTableModel*>(ptr)->sort(column, static_cast<Qt::SortOrder>(order));
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

void QSqlTableModel_DestroyQSqlTableModel(void* ptr){
	static_cast<QSqlTableModel*>(ptr)->~QSqlTableModel();
}

