// +build !minimal

#define protected public
#define private public

#include "sql.h"
#include "_cgo_export.h"

#include <QAbstractItemModel>
#include <QAbstractItemView>
#include <QByteArray>
#include <QChildEvent>
#include <QEvent>
#include <QHash>
#include <QHelpEvent>
#include <QList>
#include <QMap>
#include <QMetaMethod>
#include <QMetaObject>
#include <QMimeData>
#include <QModelIndex>
#include <QObject>
#include <QPainter>
#include <QPersistentModelIndex>
#include <QSize>
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
#include <QSqlRelationalDelegate>
#include <QSqlRelationalTableModel>
#include <QSqlResult>
#include <QSqlTableModel>
#include <QString>
#include <QStringList>
#include <QStyle>
#include <QStyleOption>
#include <QStyleOptionViewItem>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>
#include <QVariant>
#include <QVector>
#include <QWidget>

void* QSqlDatabase_QSqlDatabase_AddDatabase2(void* driver, char* connectionName)
{
	return new QSqlDatabase(QSqlDatabase::addDatabase(static_cast<QSqlDriver*>(driver), QString(connectionName)));
}

void* QSqlDatabase_QSqlDatabase_AddDatabase(char* ty, char* connectionName)
{
	return new QSqlDatabase(QSqlDatabase::addDatabase(QString(ty), QString(connectionName)));
}

void* QSqlDatabase_NewQSqlDatabase()
{
	return new QSqlDatabase();
}

void* QSqlDatabase_NewQSqlDatabase2(void* other)
{
	return new QSqlDatabase(*static_cast<QSqlDatabase*>(other));
}

void* QSqlDatabase_NewQSqlDatabase4(void* driver)
{
	return new QSqlDatabase(static_cast<QSqlDriver*>(driver));
}

void* QSqlDatabase_NewQSqlDatabase3(char* ty)
{
	return new QSqlDatabase(QString(ty));
}

void* QSqlDatabase_QSqlDatabase_CloneDatabase(void* other, char* connectionName)
{
	return new QSqlDatabase(QSqlDatabase::cloneDatabase(*static_cast<QSqlDatabase*>(other), QString(connectionName)));
}

void QSqlDatabase_Close(void* ptr)
{
	static_cast<QSqlDatabase*>(ptr)->close();
}

char QSqlDatabase_Commit(void* ptr)
{
	return static_cast<QSqlDatabase*>(ptr)->commit();
}

struct QtSql_PackedString QSqlDatabase_ConnectOptions(void* ptr)
{
	return ({ QByteArray t27c06b = static_cast<QSqlDatabase*>(ptr)->connectOptions().toUtf8(); QtSql_PackedString { const_cast<char*>(t27c06b.prepend("WHITESPACE").constData()+10), t27c06b.size()-10 }; });
}

struct QtSql_PackedString QSqlDatabase_ConnectionName(void* ptr)
{
	return ({ QByteArray tb2dc55 = static_cast<QSqlDatabase*>(ptr)->connectionName().toUtf8(); QtSql_PackedString { const_cast<char*>(tb2dc55.prepend("WHITESPACE").constData()+10), tb2dc55.size()-10 }; });
}

struct QtSql_PackedString QSqlDatabase_QSqlDatabase_ConnectionNames()
{
	return ({ QByteArray t34c9e9 = QSqlDatabase::connectionNames().join("|").toUtf8(); QtSql_PackedString { const_cast<char*>(t34c9e9.prepend("WHITESPACE").constData()+10), t34c9e9.size()-10 }; });
}

char QSqlDatabase_QSqlDatabase_Contains(char* connectionName)
{
	return QSqlDatabase::contains(QString(connectionName));
}

void* QSqlDatabase_QSqlDatabase_Database(char* connectionName, char open)
{
	return new QSqlDatabase(QSqlDatabase::database(QString(connectionName), open != 0));
}

struct QtSql_PackedString QSqlDatabase_DatabaseName(void* ptr)
{
	return ({ QByteArray t15dc5d = static_cast<QSqlDatabase*>(ptr)->databaseName().toUtf8(); QtSql_PackedString { const_cast<char*>(t15dc5d.prepend("WHITESPACE").constData()+10), t15dc5d.size()-10 }; });
}

void* QSqlDatabase_Driver(void* ptr)
{
	return static_cast<QSqlDatabase*>(ptr)->driver();
}

struct QtSql_PackedString QSqlDatabase_DriverName(void* ptr)
{
	return ({ QByteArray t58c861 = static_cast<QSqlDatabase*>(ptr)->driverName().toUtf8(); QtSql_PackedString { const_cast<char*>(t58c861.prepend("WHITESPACE").constData()+10), t58c861.size()-10 }; });
}

struct QtSql_PackedString QSqlDatabase_QSqlDatabase_Drivers()
{
	return ({ QByteArray tf3b332 = QSqlDatabase::drivers().join("|").toUtf8(); QtSql_PackedString { const_cast<char*>(tf3b332.prepend("WHITESPACE").constData()+10), tf3b332.size()-10 }; });
}

void* QSqlDatabase_Exec(void* ptr, char* query)
{
	return new QSqlQuery(static_cast<QSqlDatabase*>(ptr)->exec(QString(query)));
}

struct QtSql_PackedString QSqlDatabase_HostName(void* ptr)
{
	return ({ QByteArray taeb29f = static_cast<QSqlDatabase*>(ptr)->hostName().toUtf8(); QtSql_PackedString { const_cast<char*>(taeb29f.prepend("WHITESPACE").constData()+10), taeb29f.size()-10 }; });
}

char QSqlDatabase_QSqlDatabase_IsDriverAvailable(char* name)
{
	return QSqlDatabase::isDriverAvailable(QString(name));
}

char QSqlDatabase_IsOpen(void* ptr)
{
	return static_cast<QSqlDatabase*>(ptr)->isOpen();
}

char QSqlDatabase_IsOpenError(void* ptr)
{
	return static_cast<QSqlDatabase*>(ptr)->isOpenError();
}

char QSqlDatabase_IsValid(void* ptr)
{
	return static_cast<QSqlDatabase*>(ptr)->isValid();
}

void* QSqlDatabase_LastError(void* ptr)
{
	return new QSqlError(static_cast<QSqlDatabase*>(ptr)->lastError());
}

long long QSqlDatabase_NumericalPrecisionPolicy(void* ptr)
{
	return static_cast<QSqlDatabase*>(ptr)->numericalPrecisionPolicy();
}

char QSqlDatabase_Open(void* ptr)
{
	return static_cast<QSqlDatabase*>(ptr)->open();
}

char QSqlDatabase_Open2(void* ptr, char* user, char* password)
{
	return static_cast<QSqlDatabase*>(ptr)->open(QString(user), QString(password));
}

struct QtSql_PackedString QSqlDatabase_Password(void* ptr)
{
	return ({ QByteArray t136367 = static_cast<QSqlDatabase*>(ptr)->password().toUtf8(); QtSql_PackedString { const_cast<char*>(t136367.prepend("WHITESPACE").constData()+10), t136367.size()-10 }; });
}

int QSqlDatabase_Port(void* ptr)
{
	return static_cast<QSqlDatabase*>(ptr)->port();
}

void* QSqlDatabase_PrimaryIndex(void* ptr, char* tablename)
{
	return new QSqlIndex(static_cast<QSqlDatabase*>(ptr)->primaryIndex(QString(tablename)));
}

void* QSqlDatabase_Record(void* ptr, char* tablename)
{
	return new QSqlRecord(static_cast<QSqlDatabase*>(ptr)->record(QString(tablename)));
}

void QSqlDatabase_QSqlDatabase_RegisterSqlDriver(char* name, void* creator)
{
	QSqlDatabase::registerSqlDriver(QString(name), static_cast<QSqlDriverCreatorBase*>(creator));
}

void QSqlDatabase_QSqlDatabase_RemoveDatabase(char* connectionName)
{
	QSqlDatabase::removeDatabase(QString(connectionName));
}

char QSqlDatabase_Rollback(void* ptr)
{
	return static_cast<QSqlDatabase*>(ptr)->rollback();
}

void QSqlDatabase_SetConnectOptions(void* ptr, char* options)
{
	static_cast<QSqlDatabase*>(ptr)->setConnectOptions(QString(options));
}

void QSqlDatabase_SetDatabaseName(void* ptr, char* name)
{
	static_cast<QSqlDatabase*>(ptr)->setDatabaseName(QString(name));
}

void QSqlDatabase_SetHostName(void* ptr, char* host)
{
	static_cast<QSqlDatabase*>(ptr)->setHostName(QString(host));
}

void QSqlDatabase_SetNumericalPrecisionPolicy(void* ptr, long long precisionPolicy)
{
	static_cast<QSqlDatabase*>(ptr)->setNumericalPrecisionPolicy(static_cast<QSql::NumericalPrecisionPolicy>(precisionPolicy));
}

void QSqlDatabase_SetPassword(void* ptr, char* password)
{
	static_cast<QSqlDatabase*>(ptr)->setPassword(QString(password));
}

void QSqlDatabase_SetPort(void* ptr, int port)
{
	static_cast<QSqlDatabase*>(ptr)->setPort(port);
}

void QSqlDatabase_SetUserName(void* ptr, char* name)
{
	static_cast<QSqlDatabase*>(ptr)->setUserName(QString(name));
}

struct QtSql_PackedString QSqlDatabase_Tables(void* ptr, long long ty)
{
	return ({ QByteArray t302bb5 = static_cast<QSqlDatabase*>(ptr)->tables(static_cast<QSql::TableType>(ty)).join("|").toUtf8(); QtSql_PackedString { const_cast<char*>(t302bb5.prepend("WHITESPACE").constData()+10), t302bb5.size()-10 }; });
}

char QSqlDatabase_Transaction(void* ptr)
{
	return static_cast<QSqlDatabase*>(ptr)->transaction();
}

struct QtSql_PackedString QSqlDatabase_UserName(void* ptr)
{
	return ({ QByteArray t7a5cf4 = static_cast<QSqlDatabase*>(ptr)->userName().toUtf8(); QtSql_PackedString { const_cast<char*>(t7a5cf4.prepend("WHITESPACE").constData()+10), t7a5cf4.size()-10 }; });
}

void QSqlDatabase_DestroyQSqlDatabase(void* ptr)
{
	static_cast<QSqlDatabase*>(ptr)->~QSqlDatabase();
}

struct QtSql_PackedString QSqlDatabase_QSqlDatabase_DefaultConnection()
{
	return QtSql_PackedString { const_cast<char*>(QSqlDatabase::defaultConnection), -1 };
}

class MyQSqlDriver: public QSqlDriver
{
public:
	MyQSqlDriver(QObject *parent) : QSqlDriver(parent) {};
	bool beginTransaction() { return callbackQSqlDriver_BeginTransaction(this) != 0; };
	void close() { callbackQSqlDriver_Close(this); };
	bool commitTransaction() { return callbackQSqlDriver_CommitTransaction(this) != 0; };
	QSqlResult * createResult() const { return static_cast<QSqlResult*>(callbackQSqlDriver_CreateResult(const_cast<MyQSqlDriver*>(this))); };
	QString escapeIdentifier(const QString & identifier, QSqlDriver::IdentifierType ty) const { QByteArray tfae9fd = identifier.toUtf8(); QtSql_PackedString identifierPacked = { const_cast<char*>(tfae9fd.prepend("WHITESPACE").constData()+10), tfae9fd.size()-10 };return QString(callbackQSqlDriver_EscapeIdentifier(const_cast<MyQSqlDriver*>(this), identifierPacked, ty)); };
	QString formatValue(const QSqlField & field, bool trimStrings) const { return QString(callbackQSqlDriver_FormatValue(const_cast<MyQSqlDriver*>(this), const_cast<QSqlField*>(&field), trimStrings)); };
	QVariant handle() const { return *static_cast<QVariant*>(callbackQSqlDriver_Handle(const_cast<MyQSqlDriver*>(this))); };
	bool hasFeature(QSqlDriver::DriverFeature feature) const { return callbackQSqlDriver_HasFeature(const_cast<MyQSqlDriver*>(this), feature) != 0; };
	bool isIdentifierEscaped(const QString & identifier, QSqlDriver::IdentifierType ty) const { QByteArray tfae9fd = identifier.toUtf8(); QtSql_PackedString identifierPacked = { const_cast<char*>(tfae9fd.prepend("WHITESPACE").constData()+10), tfae9fd.size()-10 };return callbackQSqlDriver_IsIdentifierEscaped(const_cast<MyQSqlDriver*>(this), identifierPacked, ty) != 0; };
	bool isOpen() const { return callbackQSqlDriver_IsOpen(const_cast<MyQSqlDriver*>(this)) != 0; };
	void Signal_Notification(const QString & name) { QByteArray t6ae999 = name.toUtf8(); QtSql_PackedString namePacked = { const_cast<char*>(t6ae999.prepend("WHITESPACE").constData()+10), t6ae999.size()-10 };callbackQSqlDriver_Notification(this, namePacked); };
	void Signal_Notification2(const QString & name, QSqlDriver::NotificationSource source, const QVariant & payload) { QByteArray t6ae999 = name.toUtf8(); QtSql_PackedString namePacked = { const_cast<char*>(t6ae999.prepend("WHITESPACE").constData()+10), t6ae999.size()-10 };callbackQSqlDriver_Notification2(this, namePacked, source, const_cast<QVariant*>(&payload)); };
	bool open(const QString & db, const QString & user, const QString & password, const QString & host, int port, const QString & options) { QByteArray t0352a8 = db.toUtf8(); QtSql_PackedString dbPacked = { const_cast<char*>(t0352a8.prepend("WHITESPACE").constData()+10), t0352a8.size()-10 };QByteArray t12dea9 = user.toUtf8(); QtSql_PackedString userPacked = { const_cast<char*>(t12dea9.prepend("WHITESPACE").constData()+10), t12dea9.size()-10 };QByteArray t5baa61 = password.toUtf8(); QtSql_PackedString passwordPacked = { const_cast<char*>(t5baa61.prepend("WHITESPACE").constData()+10), t5baa61.size()-10 };QByteArray t86dd1c = host.toUtf8(); QtSql_PackedString hostPacked = { const_cast<char*>(t86dd1c.prepend("WHITESPACE").constData()+10), t86dd1c.size()-10 };QByteArray t513f8d = options.toUtf8(); QtSql_PackedString optionsPacked = { const_cast<char*>(t513f8d.prepend("WHITESPACE").constData()+10), t513f8d.size()-10 };return callbackQSqlDriver_Open(this, dbPacked, userPacked, passwordPacked, hostPacked, port, optionsPacked) != 0; };
	QSqlIndex primaryIndex(const QString & tableName) const { QByteArray t3e7060 = tableName.toUtf8(); QtSql_PackedString tableNamePacked = { const_cast<char*>(t3e7060.prepend("WHITESPACE").constData()+10), t3e7060.size()-10 };return *static_cast<QSqlIndex*>(callbackQSqlDriver_PrimaryIndex(const_cast<MyQSqlDriver*>(this), tableNamePacked)); };
	QSqlRecord record(const QString & tableName) const { QByteArray t3e7060 = tableName.toUtf8(); QtSql_PackedString tableNamePacked = { const_cast<char*>(t3e7060.prepend("WHITESPACE").constData()+10), t3e7060.size()-10 };return *static_cast<QSqlRecord*>(callbackQSqlDriver_Record(const_cast<MyQSqlDriver*>(this), tableNamePacked)); };
	bool rollbackTransaction() { return callbackQSqlDriver_RollbackTransaction(this) != 0; };
	void setLastError(const QSqlError & error) { callbackQSqlDriver_SetLastError(this, const_cast<QSqlError*>(&error)); };
	void setOpen(bool open) { callbackQSqlDriver_SetOpen(this, open); };
	void setOpenError(bool error) { callbackQSqlDriver_SetOpenError(this, error); };
	QString sqlStatement(QSqlDriver::StatementType ty, const QString & tableName, const QSqlRecord & rec, bool preparedStatement) const { QByteArray t3e7060 = tableName.toUtf8(); QtSql_PackedString tableNamePacked = { const_cast<char*>(t3e7060.prepend("WHITESPACE").constData()+10), t3e7060.size()-10 };return QString(callbackQSqlDriver_SqlStatement(const_cast<MyQSqlDriver*>(this), ty, tableNamePacked, const_cast<QSqlRecord*>(&rec), preparedStatement)); };
	QString stripDelimiters(const QString & identifier, QSqlDriver::IdentifierType ty) const { QByteArray tfae9fd = identifier.toUtf8(); QtSql_PackedString identifierPacked = { const_cast<char*>(tfae9fd.prepend("WHITESPACE").constData()+10), tfae9fd.size()-10 };return QString(callbackQSqlDriver_StripDelimiters(const_cast<MyQSqlDriver*>(this), identifierPacked, ty)); };
	bool subscribeToNotification(const QString & name) { QByteArray t6ae999 = name.toUtf8(); QtSql_PackedString namePacked = { const_cast<char*>(t6ae999.prepend("WHITESPACE").constData()+10), t6ae999.size()-10 };return callbackQSqlDriver_SubscribeToNotification(this, namePacked) != 0; };
	QStringList subscribedToNotifications() const { return QString(callbackQSqlDriver_SubscribedToNotifications(const_cast<MyQSqlDriver*>(this))).split("|", QString::SkipEmptyParts); };
	QStringList tables(QSql::TableType tableType) const { return QString(callbackQSqlDriver_Tables(const_cast<MyQSqlDriver*>(this), tableType)).split("|", QString::SkipEmptyParts); };
	bool unsubscribeFromNotification(const QString & name) { QByteArray t6ae999 = name.toUtf8(); QtSql_PackedString namePacked = { const_cast<char*>(t6ae999.prepend("WHITESPACE").constData()+10), t6ae999.size()-10 };return callbackQSqlDriver_UnsubscribeFromNotification(this, namePacked) != 0; };
	void timerEvent(QTimerEvent * event) { callbackQSqlDriver_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQSqlDriver_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSqlDriver_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQSqlDriver_CustomEvent(this, event); };
	void deleteLater() { callbackQSqlDriver_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSqlDriver_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQSqlDriver_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSqlDriver_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSqlDriver_MetaObject(const_cast<MyQSqlDriver*>(this))); };
};

void* QSqlDriver_NewQSqlDriver(void* parent)
{
	return new MyQSqlDriver(static_cast<QObject*>(parent));
}

char QSqlDriver_BeginTransaction(void* ptr)
{
	return static_cast<QSqlDriver*>(ptr)->beginTransaction();
}

char QSqlDriver_BeginTransactionDefault(void* ptr)
{
	return static_cast<QSqlDriver*>(ptr)->QSqlDriver::beginTransaction();
}

void QSqlDriver_Close(void* ptr)
{
	static_cast<QSqlDriver*>(ptr)->close();
}

char QSqlDriver_CommitTransaction(void* ptr)
{
	return static_cast<QSqlDriver*>(ptr)->commitTransaction();
}

char QSqlDriver_CommitTransactionDefault(void* ptr)
{
	return static_cast<QSqlDriver*>(ptr)->QSqlDriver::commitTransaction();
}

void* QSqlDriver_CreateResult(void* ptr)
{
	return static_cast<QSqlDriver*>(ptr)->createResult();
}

long long QSqlDriver_DbmsType(void* ptr)
{
	return static_cast<QSqlDriver*>(ptr)->dbmsType();
}

struct QtSql_PackedString QSqlDriver_EscapeIdentifier(void* ptr, char* identifier, long long ty)
{
	return ({ QByteArray t83477e = static_cast<QSqlDriver*>(ptr)->escapeIdentifier(QString(identifier), static_cast<QSqlDriver::IdentifierType>(ty)).toUtf8(); QtSql_PackedString { const_cast<char*>(t83477e.prepend("WHITESPACE").constData()+10), t83477e.size()-10 }; });
}

struct QtSql_PackedString QSqlDriver_EscapeIdentifierDefault(void* ptr, char* identifier, long long ty)
{
	return ({ QByteArray t7908f7 = static_cast<QSqlDriver*>(ptr)->QSqlDriver::escapeIdentifier(QString(identifier), static_cast<QSqlDriver::IdentifierType>(ty)).toUtf8(); QtSql_PackedString { const_cast<char*>(t7908f7.prepend("WHITESPACE").constData()+10), t7908f7.size()-10 }; });
}

struct QtSql_PackedString QSqlDriver_FormatValue(void* ptr, void* field, char trimStrings)
{
	return ({ QByteArray t09cd9b = static_cast<QSqlDriver*>(ptr)->formatValue(*static_cast<QSqlField*>(field), trimStrings != 0).toUtf8(); QtSql_PackedString { const_cast<char*>(t09cd9b.prepend("WHITESPACE").constData()+10), t09cd9b.size()-10 }; });
}

struct QtSql_PackedString QSqlDriver_FormatValueDefault(void* ptr, void* field, char trimStrings)
{
	return ({ QByteArray tcdc4d9 = static_cast<QSqlDriver*>(ptr)->QSqlDriver::formatValue(*static_cast<QSqlField*>(field), trimStrings != 0).toUtf8(); QtSql_PackedString { const_cast<char*>(tcdc4d9.prepend("WHITESPACE").constData()+10), tcdc4d9.size()-10 }; });
}

void* QSqlDriver_Handle(void* ptr)
{
	return new QVariant(static_cast<QSqlDriver*>(ptr)->handle());
}

void* QSqlDriver_HandleDefault(void* ptr)
{
	return new QVariant(static_cast<QSqlDriver*>(ptr)->QSqlDriver::handle());
}

char QSqlDriver_HasFeature(void* ptr, long long feature)
{
	return static_cast<QSqlDriver*>(ptr)->hasFeature(static_cast<QSqlDriver::DriverFeature>(feature));
}

char QSqlDriver_IsIdentifierEscaped(void* ptr, char* identifier, long long ty)
{
	return static_cast<QSqlDriver*>(ptr)->isIdentifierEscaped(QString(identifier), static_cast<QSqlDriver::IdentifierType>(ty));
}

char QSqlDriver_IsIdentifierEscapedDefault(void* ptr, char* identifier, long long ty)
{
	return static_cast<QSqlDriver*>(ptr)->QSqlDriver::isIdentifierEscaped(QString(identifier), static_cast<QSqlDriver::IdentifierType>(ty));
}

char QSqlDriver_IsOpen(void* ptr)
{
	return static_cast<QSqlDriver*>(ptr)->isOpen();
}

char QSqlDriver_IsOpenDefault(void* ptr)
{
	return static_cast<QSqlDriver*>(ptr)->QSqlDriver::isOpen();
}

char QSqlDriver_IsOpenError(void* ptr)
{
	return static_cast<QSqlDriver*>(ptr)->isOpenError();
}

void* QSqlDriver_LastError(void* ptr)
{
	return new QSqlError(static_cast<QSqlDriver*>(ptr)->lastError());
}

void QSqlDriver_ConnectNotification(void* ptr)
{
	QObject::connect(static_cast<QSqlDriver*>(ptr), static_cast<void (QSqlDriver::*)(const QString &)>(&QSqlDriver::notification), static_cast<MyQSqlDriver*>(ptr), static_cast<void (MyQSqlDriver::*)(const QString &)>(&MyQSqlDriver::Signal_Notification));
}

void QSqlDriver_DisconnectNotification(void* ptr)
{
	QObject::disconnect(static_cast<QSqlDriver*>(ptr), static_cast<void (QSqlDriver::*)(const QString &)>(&QSqlDriver::notification), static_cast<MyQSqlDriver*>(ptr), static_cast<void (MyQSqlDriver::*)(const QString &)>(&MyQSqlDriver::Signal_Notification));
}

void QSqlDriver_Notification(void* ptr, char* name)
{
	static_cast<QSqlDriver*>(ptr)->notification(QString(name));
}

void QSqlDriver_ConnectNotification2(void* ptr)
{
	QObject::connect(static_cast<QSqlDriver*>(ptr), static_cast<void (QSqlDriver::*)(const QString &, QSqlDriver::NotificationSource, const QVariant &)>(&QSqlDriver::notification), static_cast<MyQSqlDriver*>(ptr), static_cast<void (MyQSqlDriver::*)(const QString &, QSqlDriver::NotificationSource, const QVariant &)>(&MyQSqlDriver::Signal_Notification2));
}

void QSqlDriver_DisconnectNotification2(void* ptr)
{
	QObject::disconnect(static_cast<QSqlDriver*>(ptr), static_cast<void (QSqlDriver::*)(const QString &, QSqlDriver::NotificationSource, const QVariant &)>(&QSqlDriver::notification), static_cast<MyQSqlDriver*>(ptr), static_cast<void (MyQSqlDriver::*)(const QString &, QSqlDriver::NotificationSource, const QVariant &)>(&MyQSqlDriver::Signal_Notification2));
}

void QSqlDriver_Notification2(void* ptr, char* name, long long source, void* payload)
{
	static_cast<QSqlDriver*>(ptr)->notification(QString(name), static_cast<QSqlDriver::NotificationSource>(source), *static_cast<QVariant*>(payload));
}

long long QSqlDriver_NumericalPrecisionPolicy(void* ptr)
{
	return static_cast<QSqlDriver*>(ptr)->numericalPrecisionPolicy();
}

char QSqlDriver_Open(void* ptr, char* db, char* user, char* password, char* host, int port, char* options)
{
	return static_cast<QSqlDriver*>(ptr)->open(QString(db), QString(user), QString(password), QString(host), port, QString(options));
}

void* QSqlDriver_PrimaryIndex(void* ptr, char* tableName)
{
	return new QSqlIndex(static_cast<QSqlDriver*>(ptr)->primaryIndex(QString(tableName)));
}

void* QSqlDriver_PrimaryIndexDefault(void* ptr, char* tableName)
{
	return new QSqlIndex(static_cast<QSqlDriver*>(ptr)->QSqlDriver::primaryIndex(QString(tableName)));
}

void* QSqlDriver_Record(void* ptr, char* tableName)
{
	return new QSqlRecord(static_cast<QSqlDriver*>(ptr)->record(QString(tableName)));
}

void* QSqlDriver_RecordDefault(void* ptr, char* tableName)
{
	return new QSqlRecord(static_cast<QSqlDriver*>(ptr)->QSqlDriver::record(QString(tableName)));
}

char QSqlDriver_RollbackTransaction(void* ptr)
{
	return static_cast<QSqlDriver*>(ptr)->rollbackTransaction();
}

char QSqlDriver_RollbackTransactionDefault(void* ptr)
{
	return static_cast<QSqlDriver*>(ptr)->QSqlDriver::rollbackTransaction();
}

void QSqlDriver_SetLastError(void* ptr, void* error)
{
	static_cast<QSqlDriver*>(ptr)->setLastError(*static_cast<QSqlError*>(error));
}

void QSqlDriver_SetLastErrorDefault(void* ptr, void* error)
{
	static_cast<QSqlDriver*>(ptr)->QSqlDriver::setLastError(*static_cast<QSqlError*>(error));
}

void QSqlDriver_SetNumericalPrecisionPolicy(void* ptr, long long precisionPolicy)
{
	static_cast<QSqlDriver*>(ptr)->setNumericalPrecisionPolicy(static_cast<QSql::NumericalPrecisionPolicy>(precisionPolicy));
}

void QSqlDriver_SetOpen(void* ptr, char open)
{
	static_cast<QSqlDriver*>(ptr)->setOpen(open != 0);
}

void QSqlDriver_SetOpenDefault(void* ptr, char open)
{
	static_cast<QSqlDriver*>(ptr)->QSqlDriver::setOpen(open != 0);
}

void QSqlDriver_SetOpenError(void* ptr, char error)
{
	static_cast<QSqlDriver*>(ptr)->setOpenError(error != 0);
}

void QSqlDriver_SetOpenErrorDefault(void* ptr, char error)
{
	static_cast<QSqlDriver*>(ptr)->QSqlDriver::setOpenError(error != 0);
}

struct QtSql_PackedString QSqlDriver_SqlStatement(void* ptr, long long ty, char* tableName, void* rec, char preparedStatement)
{
	return ({ QByteArray t1bee39 = static_cast<QSqlDriver*>(ptr)->sqlStatement(static_cast<QSqlDriver::StatementType>(ty), QString(tableName), *static_cast<QSqlRecord*>(rec), preparedStatement != 0).toUtf8(); QtSql_PackedString { const_cast<char*>(t1bee39.prepend("WHITESPACE").constData()+10), t1bee39.size()-10 }; });
}

struct QtSql_PackedString QSqlDriver_SqlStatementDefault(void* ptr, long long ty, char* tableName, void* rec, char preparedStatement)
{
	return ({ QByteArray td11842 = static_cast<QSqlDriver*>(ptr)->QSqlDriver::sqlStatement(static_cast<QSqlDriver::StatementType>(ty), QString(tableName), *static_cast<QSqlRecord*>(rec), preparedStatement != 0).toUtf8(); QtSql_PackedString { const_cast<char*>(td11842.prepend("WHITESPACE").constData()+10), td11842.size()-10 }; });
}

struct QtSql_PackedString QSqlDriver_StripDelimiters(void* ptr, char* identifier, long long ty)
{
	return ({ QByteArray td6b7fd = static_cast<QSqlDriver*>(ptr)->stripDelimiters(QString(identifier), static_cast<QSqlDriver::IdentifierType>(ty)).toUtf8(); QtSql_PackedString { const_cast<char*>(td6b7fd.prepend("WHITESPACE").constData()+10), td6b7fd.size()-10 }; });
}

struct QtSql_PackedString QSqlDriver_StripDelimitersDefault(void* ptr, char* identifier, long long ty)
{
	return ({ QByteArray tdbe1ce = static_cast<QSqlDriver*>(ptr)->QSqlDriver::stripDelimiters(QString(identifier), static_cast<QSqlDriver::IdentifierType>(ty)).toUtf8(); QtSql_PackedString { const_cast<char*>(tdbe1ce.prepend("WHITESPACE").constData()+10), tdbe1ce.size()-10 }; });
}

char QSqlDriver_SubscribeToNotification(void* ptr, char* name)
{
	return static_cast<QSqlDriver*>(ptr)->subscribeToNotification(QString(name));
}

char QSqlDriver_SubscribeToNotificationDefault(void* ptr, char* name)
{
	return static_cast<QSqlDriver*>(ptr)->QSqlDriver::subscribeToNotification(QString(name));
}

struct QtSql_PackedString QSqlDriver_SubscribedToNotifications(void* ptr)
{
	return ({ QByteArray t545d09 = static_cast<QSqlDriver*>(ptr)->subscribedToNotifications().join("|").toUtf8(); QtSql_PackedString { const_cast<char*>(t545d09.prepend("WHITESPACE").constData()+10), t545d09.size()-10 }; });
}

struct QtSql_PackedString QSqlDriver_SubscribedToNotificationsDefault(void* ptr)
{
	return ({ QByteArray t063c3c = static_cast<QSqlDriver*>(ptr)->QSqlDriver::subscribedToNotifications().join("|").toUtf8(); QtSql_PackedString { const_cast<char*>(t063c3c.prepend("WHITESPACE").constData()+10), t063c3c.size()-10 }; });
}

struct QtSql_PackedString QSqlDriver_Tables(void* ptr, long long tableType)
{
	return ({ QByteArray tcfb904 = static_cast<QSqlDriver*>(ptr)->tables(static_cast<QSql::TableType>(tableType)).join("|").toUtf8(); QtSql_PackedString { const_cast<char*>(tcfb904.prepend("WHITESPACE").constData()+10), tcfb904.size()-10 }; });
}

struct QtSql_PackedString QSqlDriver_TablesDefault(void* ptr, long long tableType)
{
	return ({ QByteArray tf0f7df = static_cast<QSqlDriver*>(ptr)->QSqlDriver::tables(static_cast<QSql::TableType>(tableType)).join("|").toUtf8(); QtSql_PackedString { const_cast<char*>(tf0f7df.prepend("WHITESPACE").constData()+10), tf0f7df.size()-10 }; });
}

char QSqlDriver_UnsubscribeFromNotification(void* ptr, char* name)
{
	return static_cast<QSqlDriver*>(ptr)->unsubscribeFromNotification(QString(name));
}

char QSqlDriver_UnsubscribeFromNotificationDefault(void* ptr, char* name)
{
	return static_cast<QSqlDriver*>(ptr)->QSqlDriver::unsubscribeFromNotification(QString(name));
}

void QSqlDriver_DestroyQSqlDriver(void* ptr)
{
	static_cast<QSqlDriver*>(ptr)->~QSqlDriver();
}

void* QSqlDriver___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QSqlDriver___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSqlDriver___children_newList(void* ptr)
{
	return new QList<QObject *>;
}

void* QSqlDriver___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QSqlDriver___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QSqlDriver___dynamicPropertyNames_newList(void* ptr)
{
	return new QList<QByteArray>;
}

void* QSqlDriver___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QSqlDriver___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSqlDriver___findChildren_newList2(void* ptr)
{
	return new QList<QObject*>;
}

void* QSqlDriver___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QSqlDriver___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSqlDriver___findChildren_newList3(void* ptr)
{
	return new QList<QObject*>;
}

void* QSqlDriver___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QSqlDriver___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSqlDriver___findChildren_newList(void* ptr)
{
	return new QList<QObject*>;
}

void QSqlDriver_TimerEvent(void* ptr, void* event)
{
	static_cast<QSqlDriver*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QSqlDriver_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QSqlDriver*>(ptr)->QSqlDriver::timerEvent(static_cast<QTimerEvent*>(event));
}

void QSqlDriver_ChildEvent(void* ptr, void* event)
{
	static_cast<QSqlDriver*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QSqlDriver_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QSqlDriver*>(ptr)->QSqlDriver::childEvent(static_cast<QChildEvent*>(event));
}

void QSqlDriver_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QSqlDriver*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSqlDriver_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QSqlDriver*>(ptr)->QSqlDriver::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSqlDriver_CustomEvent(void* ptr, void* event)
{
	static_cast<QSqlDriver*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QSqlDriver_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QSqlDriver*>(ptr)->QSqlDriver::customEvent(static_cast<QEvent*>(event));
}

void QSqlDriver_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QSqlDriver*>(ptr), "deleteLater");
}

void QSqlDriver_DeleteLaterDefault(void* ptr)
{
	static_cast<QSqlDriver*>(ptr)->QSqlDriver::deleteLater();
}

void QSqlDriver_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QSqlDriver*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSqlDriver_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QSqlDriver*>(ptr)->QSqlDriver::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QSqlDriver_Event(void* ptr, void* e)
{
	return static_cast<QSqlDriver*>(ptr)->event(static_cast<QEvent*>(e));
}

char QSqlDriver_EventDefault(void* ptr, void* e)
{
	return static_cast<QSqlDriver*>(ptr)->QSqlDriver::event(static_cast<QEvent*>(e));
}

char QSqlDriver_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QSqlDriver*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QSqlDriver_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QSqlDriver*>(ptr)->QSqlDriver::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QSqlDriver_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QSqlDriver*>(ptr)->metaObject());
}

void* QSqlDriver_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QSqlDriver*>(ptr)->QSqlDriver::metaObject());
}

class MyQSqlDriverCreatorBase: public QSqlDriverCreatorBase
{
public:
	QSqlDriver * createObject() const { return static_cast<QSqlDriver*>(callbackQSqlDriverCreatorBase_CreateObject(const_cast<MyQSqlDriverCreatorBase*>(this))); };
	 ~MyQSqlDriverCreatorBase() { callbackQSqlDriverCreatorBase_DestroyQSqlDriverCreatorBase(this); };
};

void* QSqlDriverCreatorBase_CreateObject(void* ptr)
{
	return static_cast<QSqlDriverCreatorBase*>(ptr)->createObject();
}

void QSqlDriverCreatorBase_DestroyQSqlDriverCreatorBase(void* ptr)
{
	static_cast<QSqlDriverCreatorBase*>(ptr)->~QSqlDriverCreatorBase();
}

void QSqlDriverCreatorBase_DestroyQSqlDriverCreatorBaseDefault(void* ptr)
{

}

class MyQSqlDriverPlugin: public QSqlDriverPlugin
{
public:
	MyQSqlDriverPlugin(QObject *parent) : QSqlDriverPlugin(parent) {};
	QSqlDriver * create(const QString & key) { QByteArray ta62f22 = key.toUtf8(); QtSql_PackedString keyPacked = { const_cast<char*>(ta62f22.prepend("WHITESPACE").constData()+10), ta62f22.size()-10 };return static_cast<QSqlDriver*>(callbackQSqlDriverPlugin_Create(this, keyPacked)); };
	void timerEvent(QTimerEvent * event) { callbackQSqlDriverPlugin_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQSqlDriverPlugin_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSqlDriverPlugin_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQSqlDriverPlugin_CustomEvent(this, event); };
	void deleteLater() { callbackQSqlDriverPlugin_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSqlDriverPlugin_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQSqlDriverPlugin_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSqlDriverPlugin_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSqlDriverPlugin_MetaObject(const_cast<MyQSqlDriverPlugin*>(this))); };
};

void* QSqlDriverPlugin_NewQSqlDriverPlugin(void* parent)
{
	return new MyQSqlDriverPlugin(static_cast<QObject*>(parent));
}

void* QSqlDriverPlugin_Create(void* ptr, char* key)
{
	return static_cast<QSqlDriverPlugin*>(ptr)->create(QString(key));
}

void QSqlDriverPlugin_DestroyQSqlDriverPlugin(void* ptr)
{
	static_cast<QSqlDriverPlugin*>(ptr)->~QSqlDriverPlugin();
}

void* QSqlDriverPlugin___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QSqlDriverPlugin___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSqlDriverPlugin___children_newList(void* ptr)
{
	return new QList<QObject *>;
}

void* QSqlDriverPlugin___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QSqlDriverPlugin___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QSqlDriverPlugin___dynamicPropertyNames_newList(void* ptr)
{
	return new QList<QByteArray>;
}

void* QSqlDriverPlugin___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QSqlDriverPlugin___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSqlDriverPlugin___findChildren_newList2(void* ptr)
{
	return new QList<QObject*>;
}

void* QSqlDriverPlugin___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QSqlDriverPlugin___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSqlDriverPlugin___findChildren_newList3(void* ptr)
{
	return new QList<QObject*>;
}

void* QSqlDriverPlugin___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QSqlDriverPlugin___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSqlDriverPlugin___findChildren_newList(void* ptr)
{
	return new QList<QObject*>;
}

void QSqlDriverPlugin_TimerEvent(void* ptr, void* event)
{
	static_cast<QSqlDriverPlugin*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QSqlDriverPlugin_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QSqlDriverPlugin*>(ptr)->QSqlDriverPlugin::timerEvent(static_cast<QTimerEvent*>(event));
}

void QSqlDriverPlugin_ChildEvent(void* ptr, void* event)
{
	static_cast<QSqlDriverPlugin*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QSqlDriverPlugin_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QSqlDriverPlugin*>(ptr)->QSqlDriverPlugin::childEvent(static_cast<QChildEvent*>(event));
}

void QSqlDriverPlugin_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QSqlDriverPlugin*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSqlDriverPlugin_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QSqlDriverPlugin*>(ptr)->QSqlDriverPlugin::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSqlDriverPlugin_CustomEvent(void* ptr, void* event)
{
	static_cast<QSqlDriverPlugin*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QSqlDriverPlugin_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QSqlDriverPlugin*>(ptr)->QSqlDriverPlugin::customEvent(static_cast<QEvent*>(event));
}

void QSqlDriverPlugin_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QSqlDriverPlugin*>(ptr), "deleteLater");
}

void QSqlDriverPlugin_DeleteLaterDefault(void* ptr)
{
	static_cast<QSqlDriverPlugin*>(ptr)->QSqlDriverPlugin::deleteLater();
}

void QSqlDriverPlugin_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QSqlDriverPlugin*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSqlDriverPlugin_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QSqlDriverPlugin*>(ptr)->QSqlDriverPlugin::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QSqlDriverPlugin_Event(void* ptr, void* e)
{
	return static_cast<QSqlDriverPlugin*>(ptr)->event(static_cast<QEvent*>(e));
}

char QSqlDriverPlugin_EventDefault(void* ptr, void* e)
{
	return static_cast<QSqlDriverPlugin*>(ptr)->QSqlDriverPlugin::event(static_cast<QEvent*>(e));
}

char QSqlDriverPlugin_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QSqlDriverPlugin*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QSqlDriverPlugin_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QSqlDriverPlugin*>(ptr)->QSqlDriverPlugin::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QSqlDriverPlugin_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QSqlDriverPlugin*>(ptr)->metaObject());
}

void* QSqlDriverPlugin_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QSqlDriverPlugin*>(ptr)->QSqlDriverPlugin::metaObject());
}

void* QSqlError_NewQSqlError3(void* other)
{
	return new QSqlError(*static_cast<QSqlError*>(other));
}

void* QSqlError_NewQSqlError(char* driverText, char* databaseText, long long ty, char* code)
{
	return new QSqlError(QString(driverText), QString(databaseText), static_cast<QSqlError::ErrorType>(ty), QString(code));
}

struct QtSql_PackedString QSqlError_DatabaseText(void* ptr)
{
	return ({ QByteArray t7edaa1 = static_cast<QSqlError*>(ptr)->databaseText().toUtf8(); QtSql_PackedString { const_cast<char*>(t7edaa1.prepend("WHITESPACE").constData()+10), t7edaa1.size()-10 }; });
}

struct QtSql_PackedString QSqlError_DriverText(void* ptr)
{
	return ({ QByteArray tf41ece = static_cast<QSqlError*>(ptr)->driverText().toUtf8(); QtSql_PackedString { const_cast<char*>(tf41ece.prepend("WHITESPACE").constData()+10), tf41ece.size()-10 }; });
}

char QSqlError_IsValid(void* ptr)
{
	return static_cast<QSqlError*>(ptr)->isValid();
}

struct QtSql_PackedString QSqlError_NativeErrorCode(void* ptr)
{
	return ({ QByteArray t468473 = static_cast<QSqlError*>(ptr)->nativeErrorCode().toUtf8(); QtSql_PackedString { const_cast<char*>(t468473.prepend("WHITESPACE").constData()+10), t468473.size()-10 }; });
}

struct QtSql_PackedString QSqlError_Text(void* ptr)
{
	return ({ QByteArray t22e689 = static_cast<QSqlError*>(ptr)->text().toUtf8(); QtSql_PackedString { const_cast<char*>(t22e689.prepend("WHITESPACE").constData()+10), t22e689.size()-10 }; });
}

long long QSqlError_Type(void* ptr)
{
	return static_cast<QSqlError*>(ptr)->type();
}

void QSqlError_DestroyQSqlError(void* ptr)
{
	static_cast<QSqlError*>(ptr)->~QSqlError();
}

void* QSqlField_NewQSqlField2(void* other)
{
	return new QSqlField(*static_cast<QSqlField*>(other));
}

void QSqlField_Clear(void* ptr)
{
	static_cast<QSqlField*>(ptr)->clear();
}

void* QSqlField_DefaultValue(void* ptr)
{
	return new QVariant(static_cast<QSqlField*>(ptr)->defaultValue());
}

char QSqlField_IsAutoValue(void* ptr)
{
	return static_cast<QSqlField*>(ptr)->isAutoValue();
}

char QSqlField_IsGenerated(void* ptr)
{
	return static_cast<QSqlField*>(ptr)->isGenerated();
}

char QSqlField_IsNull(void* ptr)
{
	return static_cast<QSqlField*>(ptr)->isNull();
}

char QSqlField_IsReadOnly(void* ptr)
{
	return static_cast<QSqlField*>(ptr)->isReadOnly();
}

char QSqlField_IsValid(void* ptr)
{
	return static_cast<QSqlField*>(ptr)->isValid();
}

int QSqlField_Length(void* ptr)
{
	return static_cast<QSqlField*>(ptr)->length();
}

struct QtSql_PackedString QSqlField_Name(void* ptr)
{
	return ({ QByteArray t784037 = static_cast<QSqlField*>(ptr)->name().toUtf8(); QtSql_PackedString { const_cast<char*>(t784037.prepend("WHITESPACE").constData()+10), t784037.size()-10 }; });
}

int QSqlField_Precision(void* ptr)
{
	return static_cast<QSqlField*>(ptr)->precision();
}

long long QSqlField_RequiredStatus(void* ptr)
{
	return static_cast<QSqlField*>(ptr)->requiredStatus();
}

void QSqlField_SetAutoValue(void* ptr, char autoVal)
{
	static_cast<QSqlField*>(ptr)->setAutoValue(autoVal != 0);
}

void QSqlField_SetDefaultValue(void* ptr, void* value)
{
	static_cast<QSqlField*>(ptr)->setDefaultValue(*static_cast<QVariant*>(value));
}

void QSqlField_SetGenerated(void* ptr, char gen)
{
	static_cast<QSqlField*>(ptr)->setGenerated(gen != 0);
}

void QSqlField_SetLength(void* ptr, int fieldLength)
{
	static_cast<QSqlField*>(ptr)->setLength(fieldLength);
}

void QSqlField_SetName(void* ptr, char* name)
{
	static_cast<QSqlField*>(ptr)->setName(QString(name));
}

void QSqlField_SetPrecision(void* ptr, int precision)
{
	static_cast<QSqlField*>(ptr)->setPrecision(precision);
}

void QSqlField_SetReadOnly(void* ptr, char readOnly)
{
	static_cast<QSqlField*>(ptr)->setReadOnly(readOnly != 0);
}

void QSqlField_SetRequired(void* ptr, char required)
{
	static_cast<QSqlField*>(ptr)->setRequired(required != 0);
}

void QSqlField_SetRequiredStatus(void* ptr, long long required)
{
	static_cast<QSqlField*>(ptr)->setRequiredStatus(static_cast<QSqlField::RequiredStatus>(required));
}

void QSqlField_SetValue(void* ptr, void* value)
{
	static_cast<QSqlField*>(ptr)->setValue(*static_cast<QVariant*>(value));
}

void* QSqlField_Value(void* ptr)
{
	return new QVariant(static_cast<QSqlField*>(ptr)->value());
}

void QSqlField_DestroyQSqlField(void* ptr)
{
	static_cast<QSqlField*>(ptr)->~QSqlField();
}

void* QSqlIndex_NewQSqlIndex2(void* other)
{
	return new QSqlIndex(*static_cast<QSqlIndex*>(other));
}

void* QSqlIndex_NewQSqlIndex(char* cursorname, char* name)
{
	return new QSqlIndex(QString(cursorname), QString(name));
}

void QSqlIndex_Append(void* ptr, void* field)
{
	static_cast<QSqlIndex*>(ptr)->append(*static_cast<QSqlField*>(field));
}

void QSqlIndex_Append2(void* ptr, void* field, char desc)
{
	static_cast<QSqlIndex*>(ptr)->append(*static_cast<QSqlField*>(field), desc != 0);
}

struct QtSql_PackedString QSqlIndex_CursorName(void* ptr)
{
	return ({ QByteArray t1764a5 = static_cast<QSqlIndex*>(ptr)->cursorName().toUtf8(); QtSql_PackedString { const_cast<char*>(t1764a5.prepend("WHITESPACE").constData()+10), t1764a5.size()-10 }; });
}

char QSqlIndex_IsDescending(void* ptr, int i)
{
	return static_cast<QSqlIndex*>(ptr)->isDescending(i);
}

struct QtSql_PackedString QSqlIndex_Name(void* ptr)
{
	return ({ QByteArray tdb3ffc = static_cast<QSqlIndex*>(ptr)->name().toUtf8(); QtSql_PackedString { const_cast<char*>(tdb3ffc.prepend("WHITESPACE").constData()+10), tdb3ffc.size()-10 }; });
}

void QSqlIndex_SetCursorName(void* ptr, char* cursorName)
{
	static_cast<QSqlIndex*>(ptr)->setCursorName(QString(cursorName));
}

void QSqlIndex_SetDescending(void* ptr, int i, char desc)
{
	static_cast<QSqlIndex*>(ptr)->setDescending(i, desc != 0);
}

void QSqlIndex_SetName(void* ptr, char* name)
{
	static_cast<QSqlIndex*>(ptr)->setName(QString(name));
}

void QSqlIndex_DestroyQSqlIndex(void* ptr)
{
	static_cast<QSqlIndex*>(ptr)->~QSqlIndex();
}

void* QSqlQuery_NewQSqlQuery3(void* db)
{
	return new QSqlQuery(*static_cast<QSqlDatabase*>(db));
}

void* QSqlQuery_NewQSqlQuery(void* result)
{
	return new QSqlQuery(static_cast<QSqlResult*>(result));
}

void* QSqlQuery_NewQSqlQuery4(void* other)
{
	return new QSqlQuery(*static_cast<QSqlQuery*>(other));
}

void* QSqlQuery_NewQSqlQuery2(char* query, void* db)
{
	return new QSqlQuery(QString(query), *static_cast<QSqlDatabase*>(db));
}

void QSqlQuery_AddBindValue(void* ptr, void* val, long long paramType)
{
	static_cast<QSqlQuery*>(ptr)->addBindValue(*static_cast<QVariant*>(val), static_cast<QSql::ParamTypeFlag>(paramType));
}

int QSqlQuery_At(void* ptr)
{
	return static_cast<QSqlQuery*>(ptr)->at();
}

void QSqlQuery_BindValue(void* ptr, char* placeholder, void* val, long long paramType)
{
	static_cast<QSqlQuery*>(ptr)->bindValue(QString(placeholder), *static_cast<QVariant*>(val), static_cast<QSql::ParamTypeFlag>(paramType));
}

void QSqlQuery_BindValue2(void* ptr, int pos, void* val, long long paramType)
{
	static_cast<QSqlQuery*>(ptr)->bindValue(pos, *static_cast<QVariant*>(val), static_cast<QSql::ParamTypeFlag>(paramType));
}

void* QSqlQuery_BoundValue(void* ptr, char* placeholder)
{
	return new QVariant(static_cast<QSqlQuery*>(ptr)->boundValue(QString(placeholder)));
}

void* QSqlQuery_BoundValue2(void* ptr, int pos)
{
	return new QVariant(static_cast<QSqlQuery*>(ptr)->boundValue(pos));
}

struct QtSql_PackedList QSqlQuery_BoundValues(void* ptr)
{
	return ({ QMap<QString, QVariant>* tmpValue = new QMap<QString, QVariant>(static_cast<QSqlQuery*>(ptr)->boundValues()); QtSql_PackedList { tmpValue, tmpValue->size() }; });
}

void QSqlQuery_Clear(void* ptr)
{
	static_cast<QSqlQuery*>(ptr)->clear();
}

void* QSqlQuery_Driver(void* ptr)
{
	return const_cast<QSqlDriver*>(static_cast<QSqlQuery*>(ptr)->driver());
}

char QSqlQuery_Exec2(void* ptr)
{
	return static_cast<QSqlQuery*>(ptr)->exec();
}

char QSqlQuery_Exec(void* ptr, char* query)
{
	return static_cast<QSqlQuery*>(ptr)->exec(QString(query));
}

char QSqlQuery_ExecBatch(void* ptr, long long mode)
{
	return static_cast<QSqlQuery*>(ptr)->execBatch(static_cast<QSqlQuery::BatchExecutionMode>(mode));
}

struct QtSql_PackedString QSqlQuery_ExecutedQuery(void* ptr)
{
	return ({ QByteArray tc2ded8 = static_cast<QSqlQuery*>(ptr)->executedQuery().toUtf8(); QtSql_PackedString { const_cast<char*>(tc2ded8.prepend("WHITESPACE").constData()+10), tc2ded8.size()-10 }; });
}

void QSqlQuery_Finish(void* ptr)
{
	static_cast<QSqlQuery*>(ptr)->finish();
}

char QSqlQuery_First(void* ptr)
{
	return static_cast<QSqlQuery*>(ptr)->first();
}

char QSqlQuery_IsActive(void* ptr)
{
	return static_cast<QSqlQuery*>(ptr)->isActive();
}

char QSqlQuery_IsForwardOnly(void* ptr)
{
	return static_cast<QSqlQuery*>(ptr)->isForwardOnly();
}

char QSqlQuery_IsNull2(void* ptr, char* name)
{
	return static_cast<QSqlQuery*>(ptr)->isNull(QString(name));
}

char QSqlQuery_IsNull(void* ptr, int field)
{
	return static_cast<QSqlQuery*>(ptr)->isNull(field);
}

char QSqlQuery_IsSelect(void* ptr)
{
	return static_cast<QSqlQuery*>(ptr)->isSelect();
}

char QSqlQuery_IsValid(void* ptr)
{
	return static_cast<QSqlQuery*>(ptr)->isValid();
}

char QSqlQuery_Last(void* ptr)
{
	return static_cast<QSqlQuery*>(ptr)->last();
}

void* QSqlQuery_LastError(void* ptr)
{
	return new QSqlError(static_cast<QSqlQuery*>(ptr)->lastError());
}

void* QSqlQuery_LastInsertId(void* ptr)
{
	return new QVariant(static_cast<QSqlQuery*>(ptr)->lastInsertId());
}

struct QtSql_PackedString QSqlQuery_LastQuery(void* ptr)
{
	return ({ QByteArray t0d33b2 = static_cast<QSqlQuery*>(ptr)->lastQuery().toUtf8(); QtSql_PackedString { const_cast<char*>(t0d33b2.prepend("WHITESPACE").constData()+10), t0d33b2.size()-10 }; });
}

char QSqlQuery_Next(void* ptr)
{
	return static_cast<QSqlQuery*>(ptr)->next();
}

char QSqlQuery_NextResult(void* ptr)
{
	return static_cast<QSqlQuery*>(ptr)->nextResult();
}

int QSqlQuery_NumRowsAffected(void* ptr)
{
	return static_cast<QSqlQuery*>(ptr)->numRowsAffected();
}

long long QSqlQuery_NumericalPrecisionPolicy(void* ptr)
{
	return static_cast<QSqlQuery*>(ptr)->numericalPrecisionPolicy();
}

char QSqlQuery_Prepare(void* ptr, char* query)
{
	return static_cast<QSqlQuery*>(ptr)->prepare(QString(query));
}

char QSqlQuery_Previous(void* ptr)
{
	return static_cast<QSqlQuery*>(ptr)->previous();
}

void* QSqlQuery_Record(void* ptr)
{
	return new QSqlRecord(static_cast<QSqlQuery*>(ptr)->record());
}

void* QSqlQuery_Result(void* ptr)
{
	return const_cast<QSqlResult*>(static_cast<QSqlQuery*>(ptr)->result());
}

char QSqlQuery_Seek(void* ptr, int index, char relative)
{
	return static_cast<QSqlQuery*>(ptr)->seek(index, relative != 0);
}

void QSqlQuery_SetForwardOnly(void* ptr, char forward)
{
	static_cast<QSqlQuery*>(ptr)->setForwardOnly(forward != 0);
}

void QSqlQuery_SetNumericalPrecisionPolicy(void* ptr, long long precisionPolicy)
{
	static_cast<QSqlQuery*>(ptr)->setNumericalPrecisionPolicy(static_cast<QSql::NumericalPrecisionPolicy>(precisionPolicy));
}

int QSqlQuery_Size(void* ptr)
{
	return static_cast<QSqlQuery*>(ptr)->size();
}

void* QSqlQuery_Value2(void* ptr, char* name)
{
	return new QVariant(static_cast<QSqlQuery*>(ptr)->value(QString(name)));
}

void* QSqlQuery_Value(void* ptr, int index)
{
	return new QVariant(static_cast<QSqlQuery*>(ptr)->value(index));
}

void QSqlQuery_DestroyQSqlQuery(void* ptr)
{
	static_cast<QSqlQuery*>(ptr)->~QSqlQuery();
}

void* QSqlQuery___boundValues_atList(void* ptr, char* i)
{
	return new QVariant(static_cast<QMap<QString, QVariant>*>(ptr)->value(QString(i)));
}

void QSqlQuery___boundValues_setList(void* ptr, char* key, void* i)
{
	static_cast<QMap<QString, QVariant>*>(ptr)->insert(QString(key), *static_cast<QVariant*>(i));
}

void* QSqlQuery___boundValues_newList(void* ptr)
{
	return new QMap<QString, QVariant>;
}

struct QtSql_PackedList QSqlQuery___boundValues_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue = new QList<QString>(static_cast<QMap<QString, QVariant>*>(ptr)->keys()); QtSql_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtSql_PackedString QSqlQuery_____boundValues_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray t29def6 = static_cast<QList<QString>*>(ptr)->at(i).toUtf8(); QtSql_PackedString { const_cast<char*>(t29def6.prepend("WHITESPACE").constData()+10), t29def6.size()-10 }; });
}

void QSqlQuery_____boundValues_keyList_setList(void* ptr, char* i)
{
	static_cast<QList<QString>*>(ptr)->append(QString(i));
}

void* QSqlQuery_____boundValues_keyList_newList(void* ptr)
{
	return new QList<QString>;
}

class MyQSqlQueryModel: public QSqlQueryModel
{
public:
	MyQSqlQueryModel(QObject *parent) : QSqlQueryModel(parent) {};
	int rowCount(const QModelIndex & parent) const { return callbackQSqlQueryModel_RowCount(const_cast<MyQSqlQueryModel*>(this), const_cast<QModelIndex*>(&parent)); };
	QVariant data(const QModelIndex & item, int role) const { return *static_cast<QVariant*>(callbackQSqlQueryModel_Data(const_cast<MyQSqlQueryModel*>(this), const_cast<QModelIndex*>(&item), role)); };
	bool canFetchMore(const QModelIndex & parent) const { return callbackQSqlQueryModel_CanFetchMore(const_cast<MyQSqlQueryModel*>(this), const_cast<QModelIndex*>(&parent)) != 0; };
	void clear() { callbackQSqlQueryModel_Clear(this); };
	int columnCount(const QModelIndex & index) const { return callbackQSqlQueryModel_ColumnCount(const_cast<MyQSqlQueryModel*>(this), const_cast<QModelIndex*>(&index)); };
	void fetchMore(const QModelIndex & parent) { callbackQSqlQueryModel_FetchMore(this, const_cast<QModelIndex*>(&parent)); };
	QVariant headerData(int section, Qt::Orientation orientation, int role) const { return *static_cast<QVariant*>(callbackQSqlQueryModel_HeaderData(const_cast<MyQSqlQueryModel*>(this), section, orientation, role)); };
	QModelIndex indexInQuery(const QModelIndex & item) const { return *static_cast<QModelIndex*>(callbackQSqlQueryModel_IndexInQuery(const_cast<MyQSqlQueryModel*>(this), const_cast<QModelIndex*>(&item))); };
	bool insertColumns(int column, int count, const QModelIndex & parent) { return callbackQSqlQueryModel_InsertColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	void queryChange() { callbackQSqlQueryModel_QueryChange(this); };
	bool removeColumns(int column, int count, const QModelIndex & parent) { return callbackQSqlQueryModel_RemoveColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool setHeaderData(int section, Qt::Orientation orientation, const QVariant & value, int role) { return callbackQSqlQueryModel_SetHeaderData(this, section, orientation, const_cast<QVariant*>(&value), role) != 0; };
	 ~MyQSqlQueryModel() { callbackQSqlQueryModel_DestroyQSqlQueryModel(this); };
	QModelIndex index(int row, int column, const QModelIndex & parent) const { return *static_cast<QModelIndex*>(callbackQSqlQueryModel_Index(const_cast<MyQSqlQueryModel*>(this), row, column, const_cast<QModelIndex*>(&parent))); };
	bool dropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) { return callbackQSqlQueryModel_DropMimeData(this, const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	Qt::ItemFlags flags(const QModelIndex & index) const { return static_cast<Qt::ItemFlag>(callbackQSqlQueryModel_Flags(const_cast<MyQSqlQueryModel*>(this), const_cast<QModelIndex*>(&index))); };
	QModelIndex sibling(int row, int column, const QModelIndex & idx) const { return *static_cast<QModelIndex*>(callbackQSqlQueryModel_Sibling(const_cast<MyQSqlQueryModel*>(this), row, column, const_cast<QModelIndex*>(&idx))); };
	QModelIndex buddy(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackQSqlQueryModel_Buddy(const_cast<MyQSqlQueryModel*>(this), const_cast<QModelIndex*>(&index))); };
	bool canDropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) const { return callbackQSqlQueryModel_CanDropMimeData(const_cast<MyQSqlQueryModel*>(this), const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	bool hasChildren(const QModelIndex & parent) const { return callbackQSqlQueryModel_HasChildren(const_cast<MyQSqlQueryModel*>(this), const_cast<QModelIndex*>(&parent)) != 0; };
	bool insertRows(int row, int count, const QModelIndex & parent) { return callbackQSqlQueryModel_InsertRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	QMap<int, QVariant> itemData(const QModelIndex & index) const { return *static_cast<QMap<int, QVariant>*>(callbackQSqlQueryModel_ItemData(const_cast<MyQSqlQueryModel*>(this), const_cast<QModelIndex*>(&index))); };
	QList<QModelIndex> match(const QModelIndex & start, int role, const QVariant & value, int hits, Qt::MatchFlags flags) const { return *static_cast<QList<QModelIndex>*>(callbackQSqlQueryModel_Match(const_cast<MyQSqlQueryModel*>(this), const_cast<QModelIndex*>(&start), role, const_cast<QVariant*>(&value), hits, flags)); };
	QMimeData * mimeData(const QModelIndexList & indexes) const { return static_cast<QMimeData*>(callbackQSqlQueryModel_MimeData(const_cast<MyQSqlQueryModel*>(this), ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(indexes); QtSql_PackedList { tmpValue, tmpValue->size() }; }))); };
	QStringList mimeTypes() const { return QString(callbackQSqlQueryModel_MimeTypes(const_cast<MyQSqlQueryModel*>(this))).split("|", QString::SkipEmptyParts); };
	bool moveColumns(const QModelIndex & sourceParent, int sourceColumn, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackQSqlQueryModel_MoveColumns(this, const_cast<QModelIndex*>(&sourceParent), sourceColumn, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	bool moveRows(const QModelIndex & sourceParent, int sourceRow, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackQSqlQueryModel_MoveRows(this, const_cast<QModelIndex*>(&sourceParent), sourceRow, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	QModelIndex parent(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackQSqlQueryModel_Parent(const_cast<MyQSqlQueryModel*>(this), const_cast<QModelIndex*>(&index))); };
	bool removeRows(int row, int count, const QModelIndex & parent) { return callbackQSqlQueryModel_RemoveRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	void resetInternalData() { callbackQSqlQueryModel_ResetInternalData(this); };
	void revert() { callbackQSqlQueryModel_Revert(this); };
	QHash<int, QByteArray> roleNames() const { return *static_cast<QHash<int, QByteArray>*>(callbackQSqlQueryModel_RoleNames(const_cast<MyQSqlQueryModel*>(this))); };
	bool setData(const QModelIndex & index, const QVariant & value, int role) { return callbackQSqlQueryModel_SetData(this, const_cast<QModelIndex*>(&index), const_cast<QVariant*>(&value), role) != 0; };
	bool setItemData(const QModelIndex & index, const QMap<int, QVariant> & roles) { return callbackQSqlQueryModel_SetItemData(this, const_cast<QModelIndex*>(&index), ({ QMap<int, QVariant>* tmpValue = const_cast<QMap<int, QVariant>*>(&roles); QtSql_PackedList { tmpValue, tmpValue->size() }; })) != 0; };
	void sort(int column, Qt::SortOrder order) { callbackQSqlQueryModel_Sort(this, column, order); };
	QSize span(const QModelIndex & index) const { return *static_cast<QSize*>(callbackQSqlQueryModel_Span(const_cast<MyQSqlQueryModel*>(this), const_cast<QModelIndex*>(&index))); };
	bool submit() { return callbackQSqlQueryModel_Submit(this) != 0; };
	Qt::DropActions supportedDragActions() const { return static_cast<Qt::DropAction>(callbackQSqlQueryModel_SupportedDragActions(const_cast<MyQSqlQueryModel*>(this))); };
	Qt::DropActions supportedDropActions() const { return static_cast<Qt::DropAction>(callbackQSqlQueryModel_SupportedDropActions(const_cast<MyQSqlQueryModel*>(this))); };
	void timerEvent(QTimerEvent * event) { callbackQSqlQueryModel_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQSqlQueryModel_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSqlQueryModel_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQSqlQueryModel_CustomEvent(this, event); };
	void deleteLater() { callbackQSqlQueryModel_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSqlQueryModel_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQSqlQueryModel_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSqlQueryModel_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSqlQueryModel_MetaObject(const_cast<MyQSqlQueryModel*>(this))); };
};

int QSqlQueryModel_RowCount(void* ptr, void* parent)
{
	return static_cast<QSqlQueryModel*>(ptr)->rowCount(*static_cast<QModelIndex*>(parent));
}

int QSqlQueryModel_RowCountDefault(void* ptr, void* parent)
{
	return static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::rowCount(*static_cast<QModelIndex*>(parent));
}

void* QSqlQueryModel_NewQSqlQueryModel(void* parent)
{
	return new MyQSqlQueryModel(static_cast<QObject*>(parent));
}

void* QSqlQueryModel_Data(void* ptr, void* item, int role)
{
	return new QVariant(static_cast<QSqlQueryModel*>(ptr)->data(*static_cast<QModelIndex*>(item), role));
}

void* QSqlQueryModel_DataDefault(void* ptr, void* item, int role)
{
	return new QVariant(static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::data(*static_cast<QModelIndex*>(item), role));
}

char QSqlQueryModel_CanFetchMore(void* ptr, void* parent)
{
	return static_cast<QSqlQueryModel*>(ptr)->canFetchMore(*static_cast<QModelIndex*>(parent));
}

char QSqlQueryModel_CanFetchMoreDefault(void* ptr, void* parent)
{
	return static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::canFetchMore(*static_cast<QModelIndex*>(parent));
}

void QSqlQueryModel_Clear(void* ptr)
{
	static_cast<QSqlQueryModel*>(ptr)->clear();
}

void QSqlQueryModel_ClearDefault(void* ptr)
{
	static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::clear();
}

int QSqlQueryModel_ColumnCount(void* ptr, void* index)
{
	return static_cast<QSqlQueryModel*>(ptr)->columnCount(*static_cast<QModelIndex*>(index));
}

int QSqlQueryModel_ColumnCountDefault(void* ptr, void* index)
{
	return static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::columnCount(*static_cast<QModelIndex*>(index));
}

void QSqlQueryModel_FetchMore(void* ptr, void* parent)
{
	static_cast<QSqlQueryModel*>(ptr)->fetchMore(*static_cast<QModelIndex*>(parent));
}

void QSqlQueryModel_FetchMoreDefault(void* ptr, void* parent)
{
	static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::fetchMore(*static_cast<QModelIndex*>(parent));
}

void* QSqlQueryModel_HeaderData(void* ptr, int section, long long orientation, int role)
{
	return new QVariant(static_cast<QSqlQueryModel*>(ptr)->headerData(section, static_cast<Qt::Orientation>(orientation), role));
}

void* QSqlQueryModel_HeaderDataDefault(void* ptr, int section, long long orientation, int role)
{
	return new QVariant(static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::headerData(section, static_cast<Qt::Orientation>(orientation), role));
}

void* QSqlQueryModel_IndexInQuery(void* ptr, void* item)
{
	return new QModelIndex(static_cast<QSqlQueryModel*>(ptr)->indexInQuery(*static_cast<QModelIndex*>(item)));
}

void* QSqlQueryModel_IndexInQueryDefault(void* ptr, void* item)
{
	return new QModelIndex(static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::indexInQuery(*static_cast<QModelIndex*>(item)));
}

char QSqlQueryModel_InsertColumns(void* ptr, int column, int count, void* parent)
{
	return static_cast<QSqlQueryModel*>(ptr)->insertColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char QSqlQueryModel_InsertColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::insertColumns(column, count, *static_cast<QModelIndex*>(parent));
}

void* QSqlQueryModel_LastError(void* ptr)
{
	return new QSqlError(static_cast<QSqlQueryModel*>(ptr)->lastError());
}

void* QSqlQueryModel_Query(void* ptr)
{
	return new QSqlQuery(static_cast<QSqlQueryModel*>(ptr)->query());
}

void QSqlQueryModel_QueryChange(void* ptr)
{
	static_cast<QSqlQueryModel*>(ptr)->queryChange();
}

void QSqlQueryModel_QueryChangeDefault(void* ptr)
{
	static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::queryChange();
}

void* QSqlQueryModel_Record2(void* ptr)
{
	return new QSqlRecord(static_cast<QSqlQueryModel*>(ptr)->record());
}

void* QSqlQueryModel_Record(void* ptr, int row)
{
	return new QSqlRecord(static_cast<QSqlQueryModel*>(ptr)->record(row));
}

char QSqlQueryModel_RemoveColumns(void* ptr, int column, int count, void* parent)
{
	return static_cast<QSqlQueryModel*>(ptr)->removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char QSqlQueryModel_RemoveColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char QSqlQueryModel_SetHeaderData(void* ptr, int section, long long orientation, void* value, int role)
{
	return static_cast<QSqlQueryModel*>(ptr)->setHeaderData(section, static_cast<Qt::Orientation>(orientation), *static_cast<QVariant*>(value), role);
}

char QSqlQueryModel_SetHeaderDataDefault(void* ptr, int section, long long orientation, void* value, int role)
{
	return static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::setHeaderData(section, static_cast<Qt::Orientation>(orientation), *static_cast<QVariant*>(value), role);
}

void QSqlQueryModel_SetLastError(void* ptr, void* error)
{
	static_cast<QSqlQueryModel*>(ptr)->setLastError(*static_cast<QSqlError*>(error));
}

void QSqlQueryModel_SetQuery(void* ptr, void* query)
{
	static_cast<QSqlQueryModel*>(ptr)->setQuery(*static_cast<QSqlQuery*>(query));
}

void QSqlQueryModel_SetQuery2(void* ptr, char* query, void* db)
{
	static_cast<QSqlQueryModel*>(ptr)->setQuery(QString(query), *static_cast<QSqlDatabase*>(db));
}

void QSqlQueryModel_DestroyQSqlQueryModel(void* ptr)
{
	static_cast<QSqlQueryModel*>(ptr)->~QSqlQueryModel();
}

void QSqlQueryModel_DestroyQSqlQueryModelDefault(void* ptr)
{

}

int QSqlQueryModel_____itemData_keyList_atList(void* ptr, int i)
{
	return static_cast<QList<int>*>(ptr)->at(i);
}

void QSqlQueryModel_____itemData_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* QSqlQueryModel_____itemData_keyList_newList(void* ptr)
{
	return new QList<int>;
}

int QSqlQueryModel_____roleNames_keyList_atList(void* ptr, int i)
{
	return static_cast<QList<int>*>(ptr)->at(i);
}

void QSqlQueryModel_____roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* QSqlQueryModel_____roleNames_keyList_newList(void* ptr)
{
	return new QList<int>;
}

int QSqlQueryModel_____setItemData_keyList_atList(void* ptr, int i)
{
	return static_cast<QList<int>*>(ptr)->at(i);
}

void QSqlQueryModel_____setItemData_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* QSqlQueryModel_____setItemData_keyList_newList(void* ptr)
{
	return new QList<int>;
}

void* QSqlQueryModel___changePersistentIndexList_from_atList(void* ptr, int i)
{
	return new QModelIndex(static_cast<QList<QModelIndex>*>(ptr)->at(i));
}

void QSqlQueryModel___changePersistentIndexList_from_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* QSqlQueryModel___changePersistentIndexList_from_newList(void* ptr)
{
	return new QList<QModelIndex>;
}

void* QSqlQueryModel___changePersistentIndexList_to_atList(void* ptr, int i)
{
	return new QModelIndex(static_cast<QList<QModelIndex>*>(ptr)->at(i));
}

void QSqlQueryModel___changePersistentIndexList_to_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* QSqlQueryModel___changePersistentIndexList_to_newList(void* ptr)
{
	return new QList<QModelIndex>;
}

int QSqlQueryModel___dataChanged_roles_atList(void* ptr, int i)
{
	return static_cast<QVector<int>*>(ptr)->at(i);
}

void QSqlQueryModel___dataChanged_roles_setList(void* ptr, int i)
{
	static_cast<QVector<int>*>(ptr)->append(i);
}

void* QSqlQueryModel___dataChanged_roles_newList(void* ptr)
{
	return new QVector<int>;
}

void* QSqlQueryModel___itemData_atList(void* ptr, int i)
{
	return new QVariant(static_cast<QMap<int, QVariant>*>(ptr)->value(i));
}

void QSqlQueryModel___itemData_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* QSqlQueryModel___itemData_newList(void* ptr)
{
	return new QMap<int, QVariant>;
}

struct QtSql_PackedList QSqlQueryModel___itemData_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); QtSql_PackedList { tmpValue, tmpValue->size() }; });
}

void* QSqlQueryModel___layoutAboutToBeChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i));
}

void QSqlQueryModel___layoutAboutToBeChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* QSqlQueryModel___layoutAboutToBeChanged_parents_newList(void* ptr)
{
	return new QList<QPersistentModelIndex>;
}

void* QSqlQueryModel___layoutChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i));
}

void QSqlQueryModel___layoutChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* QSqlQueryModel___layoutChanged_parents_newList(void* ptr)
{
	return new QList<QPersistentModelIndex>;
}

void* QSqlQueryModel___match_atList(void* ptr, int i)
{
	return new QModelIndex(static_cast<QList<QModelIndex>*>(ptr)->at(i));
}

void QSqlQueryModel___match_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* QSqlQueryModel___match_newList(void* ptr)
{
	return new QList<QModelIndex>;
}

void* QSqlQueryModel___mimeData_indexes_atList(void* ptr, int i)
{
	return new QModelIndex(static_cast<QList<QModelIndex>*>(ptr)->at(i));
}

void QSqlQueryModel___mimeData_indexes_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* QSqlQueryModel___mimeData_indexes_newList(void* ptr)
{
	return new QList<QModelIndex>;
}

void* QSqlQueryModel___persistentIndexList_atList(void* ptr, int i)
{
	return new QModelIndex(static_cast<QList<QModelIndex>*>(ptr)->at(i));
}

void QSqlQueryModel___persistentIndexList_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* QSqlQueryModel___persistentIndexList_newList(void* ptr)
{
	return new QList<QModelIndex>;
}

void* QSqlQueryModel___roleNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QHash<int, QByteArray>*>(ptr)->value(i));
}

void QSqlQueryModel___roleNames_setList(void* ptr, int key, void* i)
{
	static_cast<QHash<int, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* QSqlQueryModel___roleNames_newList(void* ptr)
{
	return new QHash<int, QByteArray>;
}

struct QtSql_PackedList QSqlQueryModel___roleNames_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QHash<int, QByteArray>*>(ptr)->keys()); QtSql_PackedList { tmpValue, tmpValue->size() }; });
}

void* QSqlQueryModel___setItemData_roles_atList(void* ptr, int i)
{
	return new QVariant(static_cast<QMap<int, QVariant>*>(ptr)->value(i));
}

void QSqlQueryModel___setItemData_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* QSqlQueryModel___setItemData_roles_newList(void* ptr)
{
	return new QMap<int, QVariant>;
}

struct QtSql_PackedList QSqlQueryModel___setItemData_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); QtSql_PackedList { tmpValue, tmpValue->size() }; });
}

int QSqlQueryModel_____doSetRoleNames_keyList_atList(void* ptr, int i)
{
	return static_cast<QList<int>*>(ptr)->at(i);
}

void QSqlQueryModel_____doSetRoleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* QSqlQueryModel_____doSetRoleNames_keyList_newList(void* ptr)
{
	return new QList<int>;
}

int QSqlQueryModel_____setRoleNames_keyList_atList(void* ptr, int i)
{
	return static_cast<QList<int>*>(ptr)->at(i);
}

void QSqlQueryModel_____setRoleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* QSqlQueryModel_____setRoleNames_keyList_newList(void* ptr)
{
	return new QList<int>;
}

void* QSqlQueryModel___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QSqlQueryModel___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSqlQueryModel___children_newList(void* ptr)
{
	return new QList<QObject *>;
}

void* QSqlQueryModel___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QSqlQueryModel___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QSqlQueryModel___dynamicPropertyNames_newList(void* ptr)
{
	return new QList<QByteArray>;
}

void* QSqlQueryModel___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QSqlQueryModel___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSqlQueryModel___findChildren_newList2(void* ptr)
{
	return new QList<QObject*>;
}

void* QSqlQueryModel___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QSqlQueryModel___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSqlQueryModel___findChildren_newList3(void* ptr)
{
	return new QList<QObject*>;
}

void* QSqlQueryModel___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QSqlQueryModel___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSqlQueryModel___findChildren_newList(void* ptr)
{
	return new QList<QObject*>;
}

void* QSqlQueryModel_Index(void* ptr, int row, int column, void* parent)
{
	return new QModelIndex(static_cast<QSqlQueryModel*>(ptr)->index(row, column, *static_cast<QModelIndex*>(parent)));
}

void* QSqlQueryModel_IndexDefault(void* ptr, int row, int column, void* parent)
{
	return new QModelIndex(static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::index(row, column, *static_cast<QModelIndex*>(parent)));
}

char QSqlQueryModel_DropMimeData(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<QSqlQueryModel*>(ptr)->dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

char QSqlQueryModel_DropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

long long QSqlQueryModel_Flags(void* ptr, void* index)
{
	return static_cast<QSqlQueryModel*>(ptr)->flags(*static_cast<QModelIndex*>(index));
}

long long QSqlQueryModel_FlagsDefault(void* ptr, void* index)
{
	return static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::flags(*static_cast<QModelIndex*>(index));
}

void* QSqlQueryModel_Sibling(void* ptr, int row, int column, void* idx)
{
	return new QModelIndex(static_cast<QSqlQueryModel*>(ptr)->sibling(row, column, *static_cast<QModelIndex*>(idx)));
}

void* QSqlQueryModel_SiblingDefault(void* ptr, int row, int column, void* idx)
{
	return new QModelIndex(static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::sibling(row, column, *static_cast<QModelIndex*>(idx)));
}

void* QSqlQueryModel_Buddy(void* ptr, void* index)
{
	return new QModelIndex(static_cast<QSqlQueryModel*>(ptr)->buddy(*static_cast<QModelIndex*>(index)));
}

void* QSqlQueryModel_BuddyDefault(void* ptr, void* index)
{
	return new QModelIndex(static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::buddy(*static_cast<QModelIndex*>(index)));
}

char QSqlQueryModel_CanDropMimeData(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<QSqlQueryModel*>(ptr)->canDropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

char QSqlQueryModel_CanDropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::canDropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

char QSqlQueryModel_HasChildren(void* ptr, void* parent)
{
	return static_cast<QSqlQueryModel*>(ptr)->hasChildren(*static_cast<QModelIndex*>(parent));
}

char QSqlQueryModel_HasChildrenDefault(void* ptr, void* parent)
{
	return static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::hasChildren(*static_cast<QModelIndex*>(parent));
}

char QSqlQueryModel_InsertRows(void* ptr, int row, int count, void* parent)
{
	return static_cast<QSqlQueryModel*>(ptr)->insertRows(row, count, *static_cast<QModelIndex*>(parent));
}

char QSqlQueryModel_InsertRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::insertRows(row, count, *static_cast<QModelIndex*>(parent));
}

struct QtSql_PackedList QSqlQueryModel_ItemData(void* ptr, void* index)
{
	return ({ QMap<int, QVariant>* tmpValue = new QMap<int, QVariant>(static_cast<QSqlQueryModel*>(ptr)->itemData(*static_cast<QModelIndex*>(index))); QtSql_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtSql_PackedList QSqlQueryModel_ItemDataDefault(void* ptr, void* index)
{
	return ({ QMap<int, QVariant>* tmpValue = new QMap<int, QVariant>(static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::itemData(*static_cast<QModelIndex*>(index))); QtSql_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtSql_PackedList QSqlQueryModel_Match(void* ptr, void* start, int role, void* value, int hits, long long flags)
{
	return ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(static_cast<QSqlQueryModel*>(ptr)->match(*static_cast<QModelIndex*>(start), role, *static_cast<QVariant*>(value), hits, static_cast<Qt::MatchFlag>(flags))); QtSql_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtSql_PackedList QSqlQueryModel_MatchDefault(void* ptr, void* start, int role, void* value, int hits, long long flags)
{
	return ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::match(*static_cast<QModelIndex*>(start), role, *static_cast<QVariant*>(value), hits, static_cast<Qt::MatchFlag>(flags))); QtSql_PackedList { tmpValue, tmpValue->size() }; });
}

void* QSqlQueryModel_MimeData(void* ptr, void* indexes)
{
	return static_cast<QSqlQueryModel*>(ptr)->mimeData(*static_cast<QList<QModelIndex>*>(indexes));
}

void* QSqlQueryModel_MimeDataDefault(void* ptr, void* indexes)
{
	return static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::mimeData(*static_cast<QList<QModelIndex>*>(indexes));
}

struct QtSql_PackedString QSqlQueryModel_MimeTypes(void* ptr)
{
	return ({ QByteArray t1df073 = static_cast<QSqlQueryModel*>(ptr)->mimeTypes().join("|").toUtf8(); QtSql_PackedString { const_cast<char*>(t1df073.prepend("WHITESPACE").constData()+10), t1df073.size()-10 }; });
}

struct QtSql_PackedString QSqlQueryModel_MimeTypesDefault(void* ptr)
{
	return ({ QByteArray t026d85 = static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::mimeTypes().join("|").toUtf8(); QtSql_PackedString { const_cast<char*>(t026d85.prepend("WHITESPACE").constData()+10), t026d85.size()-10 }; });
}

char QSqlQueryModel_MoveColumns(void* ptr, void* sourceParent, int sourceColumn, int count, void* destinationParent, int destinationChild)
{
	return static_cast<QSqlQueryModel*>(ptr)->moveColumns(*static_cast<QModelIndex*>(sourceParent), sourceColumn, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

char QSqlQueryModel_MoveColumnsDefault(void* ptr, void* sourceParent, int sourceColumn, int count, void* destinationParent, int destinationChild)
{
	return static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::moveColumns(*static_cast<QModelIndex*>(sourceParent), sourceColumn, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

char QSqlQueryModel_MoveRows(void* ptr, void* sourceParent, int sourceRow, int count, void* destinationParent, int destinationChild)
{
	return static_cast<QSqlQueryModel*>(ptr)->moveRows(*static_cast<QModelIndex*>(sourceParent), sourceRow, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

char QSqlQueryModel_MoveRowsDefault(void* ptr, void* sourceParent, int sourceRow, int count, void* destinationParent, int destinationChild)
{
	return static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::moveRows(*static_cast<QModelIndex*>(sourceParent), sourceRow, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

void* QSqlQueryModel_Parent(void* ptr, void* index)
{
	return new QModelIndex(static_cast<QSqlQueryModel*>(ptr)->parent(*static_cast<QModelIndex*>(index)));
}

char QSqlQueryModel_RemoveRows(void* ptr, int row, int count, void* parent)
{
	return static_cast<QSqlQueryModel*>(ptr)->removeRows(row, count, *static_cast<QModelIndex*>(parent));
}

char QSqlQueryModel_RemoveRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::removeRows(row, count, *static_cast<QModelIndex*>(parent));
}

void QSqlQueryModel_ResetInternalData(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QSqlQueryModel*>(ptr), "resetInternalData");
}

void QSqlQueryModel_ResetInternalDataDefault(void* ptr)
{
	static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::resetInternalData();
}

void QSqlQueryModel_Revert(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QSqlQueryModel*>(ptr), "revert");
}

void QSqlQueryModel_RevertDefault(void* ptr)
{
	static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::revert();
}

struct QtSql_PackedList QSqlQueryModel_RoleNames(void* ptr)
{
	return ({ QHash<int, QByteArray>* tmpValue = new QHash<int, QByteArray>(static_cast<QSqlQueryModel*>(ptr)->roleNames()); QtSql_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtSql_PackedList QSqlQueryModel_RoleNamesDefault(void* ptr)
{
	return ({ QHash<int, QByteArray>* tmpValue = new QHash<int, QByteArray>(static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::roleNames()); QtSql_PackedList { tmpValue, tmpValue->size() }; });
}

char QSqlQueryModel_SetData(void* ptr, void* index, void* value, int role)
{
	return static_cast<QSqlQueryModel*>(ptr)->setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
}

char QSqlQueryModel_SetDataDefault(void* ptr, void* index, void* value, int role)
{
	return static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
}

char QSqlQueryModel_SetItemData(void* ptr, void* index, void* roles)
{
	return static_cast<QSqlQueryModel*>(ptr)->setItemData(*static_cast<QModelIndex*>(index), *static_cast<QMap<int, QVariant>*>(roles));
}

char QSqlQueryModel_SetItemDataDefault(void* ptr, void* index, void* roles)
{
	return static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::setItemData(*static_cast<QModelIndex*>(index), *static_cast<QMap<int, QVariant>*>(roles));
}

void QSqlQueryModel_Sort(void* ptr, int column, long long order)
{
	static_cast<QSqlQueryModel*>(ptr)->sort(column, static_cast<Qt::SortOrder>(order));
}

void QSqlQueryModel_SortDefault(void* ptr, int column, long long order)
{
	static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::sort(column, static_cast<Qt::SortOrder>(order));
}

void* QSqlQueryModel_Span(void* ptr, void* index)
{
	return ({ QSize tmpValue = static_cast<QSqlQueryModel*>(ptr)->span(*static_cast<QModelIndex*>(index)); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QSqlQueryModel_SpanDefault(void* ptr, void* index)
{
	return ({ QSize tmpValue = static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::span(*static_cast<QModelIndex*>(index)); new QSize(tmpValue.width(), tmpValue.height()); });
}

char QSqlQueryModel_Submit(void* ptr)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QSqlQueryModel*>(ptr), "submit", Q_RETURN_ARG(bool, returnArg));
	return returnArg;
}

char QSqlQueryModel_SubmitDefault(void* ptr)
{
	return static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::submit();
}

long long QSqlQueryModel_SupportedDragActions(void* ptr)
{
	return static_cast<QSqlQueryModel*>(ptr)->supportedDragActions();
}

long long QSqlQueryModel_SupportedDragActionsDefault(void* ptr)
{
	return static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::supportedDragActions();
}

long long QSqlQueryModel_SupportedDropActions(void* ptr)
{
	return static_cast<QSqlQueryModel*>(ptr)->supportedDropActions();
}

long long QSqlQueryModel_SupportedDropActionsDefault(void* ptr)
{
	return static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::supportedDropActions();
}

void QSqlQueryModel_TimerEvent(void* ptr, void* event)
{
	static_cast<QSqlQueryModel*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QSqlQueryModel_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::timerEvent(static_cast<QTimerEvent*>(event));
}

void QSqlQueryModel_ChildEvent(void* ptr, void* event)
{
	static_cast<QSqlQueryModel*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QSqlQueryModel_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::childEvent(static_cast<QChildEvent*>(event));
}

void QSqlQueryModel_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QSqlQueryModel*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSqlQueryModel_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSqlQueryModel_CustomEvent(void* ptr, void* event)
{
	static_cast<QSqlQueryModel*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QSqlQueryModel_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::customEvent(static_cast<QEvent*>(event));
}

void QSqlQueryModel_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QSqlQueryModel*>(ptr), "deleteLater");
}

void QSqlQueryModel_DeleteLaterDefault(void* ptr)
{
	static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::deleteLater();
}

void QSqlQueryModel_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QSqlQueryModel*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSqlQueryModel_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QSqlQueryModel_Event(void* ptr, void* e)
{
	return static_cast<QSqlQueryModel*>(ptr)->event(static_cast<QEvent*>(e));
}

char QSqlQueryModel_EventDefault(void* ptr, void* e)
{
	return static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::event(static_cast<QEvent*>(e));
}

char QSqlQueryModel_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QSqlQueryModel*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QSqlQueryModel_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QSqlQueryModel_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QSqlQueryModel*>(ptr)->metaObject());
}

void* QSqlQueryModel_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::metaObject());
}

void* QSqlRecord_NewQSqlRecord()
{
	return new QSqlRecord();
}

void* QSqlRecord_NewQSqlRecord2(void* other)
{
	return new QSqlRecord(*static_cast<QSqlRecord*>(other));
}

void QSqlRecord_Append(void* ptr, void* field)
{
	static_cast<QSqlRecord*>(ptr)->append(*static_cast<QSqlField*>(field));
}

void QSqlRecord_Clear(void* ptr)
{
	static_cast<QSqlRecord*>(ptr)->clear();
}

void QSqlRecord_ClearValues(void* ptr)
{
	static_cast<QSqlRecord*>(ptr)->clearValues();
}

char QSqlRecord_Contains(void* ptr, char* name)
{
	return static_cast<QSqlRecord*>(ptr)->contains(QString(name));
}

int QSqlRecord_Count(void* ptr)
{
	return static_cast<QSqlRecord*>(ptr)->count();
}

void* QSqlRecord_Field2(void* ptr, char* name)
{
	return new QSqlField(static_cast<QSqlRecord*>(ptr)->field(QString(name)));
}

void* QSqlRecord_Field(void* ptr, int index)
{
	return new QSqlField(static_cast<QSqlRecord*>(ptr)->field(index));
}

struct QtSql_PackedString QSqlRecord_FieldName(void* ptr, int index)
{
	return ({ QByteArray te132cb = static_cast<QSqlRecord*>(ptr)->fieldName(index).toUtf8(); QtSql_PackedString { const_cast<char*>(te132cb.prepend("WHITESPACE").constData()+10), te132cb.size()-10 }; });
}

int QSqlRecord_IndexOf(void* ptr, char* name)
{
	return static_cast<QSqlRecord*>(ptr)->indexOf(QString(name));
}

void QSqlRecord_Insert(void* ptr, int pos, void* field)
{
	static_cast<QSqlRecord*>(ptr)->insert(pos, *static_cast<QSqlField*>(field));
}

char QSqlRecord_IsEmpty(void* ptr)
{
	return static_cast<QSqlRecord*>(ptr)->isEmpty();
}

char QSqlRecord_IsGenerated(void* ptr, char* name)
{
	return static_cast<QSqlRecord*>(ptr)->isGenerated(QString(name));
}

char QSqlRecord_IsGenerated2(void* ptr, int index)
{
	return static_cast<QSqlRecord*>(ptr)->isGenerated(index);
}

char QSqlRecord_IsNull(void* ptr, char* name)
{
	return static_cast<QSqlRecord*>(ptr)->isNull(QString(name));
}

char QSqlRecord_IsNull2(void* ptr, int index)
{
	return static_cast<QSqlRecord*>(ptr)->isNull(index);
}

void* QSqlRecord_KeyValues(void* ptr, void* keyFields)
{
	return new QSqlRecord(static_cast<QSqlRecord*>(ptr)->keyValues(*static_cast<QSqlRecord*>(keyFields)));
}

void QSqlRecord_Remove(void* ptr, int pos)
{
	static_cast<QSqlRecord*>(ptr)->remove(pos);
}

void QSqlRecord_Replace(void* ptr, int pos, void* field)
{
	static_cast<QSqlRecord*>(ptr)->replace(pos, *static_cast<QSqlField*>(field));
}

void QSqlRecord_SetGenerated(void* ptr, char* name, char generated)
{
	static_cast<QSqlRecord*>(ptr)->setGenerated(QString(name), generated != 0);
}

void QSqlRecord_SetGenerated2(void* ptr, int index, char generated)
{
	static_cast<QSqlRecord*>(ptr)->setGenerated(index, generated != 0);
}

void QSqlRecord_SetNull2(void* ptr, char* name)
{
	static_cast<QSqlRecord*>(ptr)->setNull(QString(name));
}

void QSqlRecord_SetNull(void* ptr, int index)
{
	static_cast<QSqlRecord*>(ptr)->setNull(index);
}

void QSqlRecord_SetValue2(void* ptr, char* name, void* val)
{
	static_cast<QSqlRecord*>(ptr)->setValue(QString(name), *static_cast<QVariant*>(val));
}

void QSqlRecord_SetValue(void* ptr, int index, void* val)
{
	static_cast<QSqlRecord*>(ptr)->setValue(index, *static_cast<QVariant*>(val));
}

void* QSqlRecord_Value2(void* ptr, char* name)
{
	return new QVariant(static_cast<QSqlRecord*>(ptr)->value(QString(name)));
}

void* QSqlRecord_Value(void* ptr, int index)
{
	return new QVariant(static_cast<QSqlRecord*>(ptr)->value(index));
}

void QSqlRecord_DestroyQSqlRecord(void* ptr)
{
	static_cast<QSqlRecord*>(ptr)->~QSqlRecord();
}

void* QSqlRelation_NewQSqlRelation()
{
	return new QSqlRelation();
}

void* QSqlRelation_NewQSqlRelation2(char* tableName, char* indexColumn, char* displayColumn)
{
	return new QSqlRelation(QString(tableName), QString(indexColumn), QString(displayColumn));
}

struct QtSql_PackedString QSqlRelation_DisplayColumn(void* ptr)
{
	return ({ QByteArray t5efc39 = static_cast<QSqlRelation*>(ptr)->displayColumn().toUtf8(); QtSql_PackedString { const_cast<char*>(t5efc39.prepend("WHITESPACE").constData()+10), t5efc39.size()-10 }; });
}

struct QtSql_PackedString QSqlRelation_IndexColumn(void* ptr)
{
	return ({ QByteArray t3d4445 = static_cast<QSqlRelation*>(ptr)->indexColumn().toUtf8(); QtSql_PackedString { const_cast<char*>(t3d4445.prepend("WHITESPACE").constData()+10), t3d4445.size()-10 }; });
}

char QSqlRelation_IsValid(void* ptr)
{
	return static_cast<QSqlRelation*>(ptr)->isValid();
}

struct QtSql_PackedString QSqlRelation_TableName(void* ptr)
{
	return ({ QByteArray t33aa41 = static_cast<QSqlRelation*>(ptr)->tableName().toUtf8(); QtSql_PackedString { const_cast<char*>(t33aa41.prepend("WHITESPACE").constData()+10), t33aa41.size()-10 }; });
}

class MyQSqlRelationalDelegate: public QSqlRelationalDelegate
{
public:
	MyQSqlRelationalDelegate(QObject *parent) : QSqlRelationalDelegate(parent) {};
	QWidget * createEditor(QWidget * parent, const QStyleOptionViewItem & option, const QModelIndex & index) const { return static_cast<QWidget*>(callbackQSqlRelationalDelegate_CreateEditor(const_cast<MyQSqlRelationalDelegate*>(this), parent, const_cast<QStyleOptionViewItem*>(&option), const_cast<QModelIndex*>(&index))); };
	void setModelData(QWidget * editor, QAbstractItemModel * model, const QModelIndex & index) const { callbackQSqlRelationalDelegate_SetModelData(const_cast<MyQSqlRelationalDelegate*>(this), editor, model, const_cast<QModelIndex*>(&index)); };
	bool eventFilter(QObject * editor, QEvent * event) { return callbackQSqlRelationalDelegate_EventFilter(this, editor, event) != 0; };
	bool editorEvent(QEvent * event, QAbstractItemModel * model, const QStyleOptionViewItem & option, const QModelIndex & index) { return callbackQSqlRelationalDelegate_EditorEvent(this, event, model, const_cast<QStyleOptionViewItem*>(&option), const_cast<QModelIndex*>(&index)) != 0; };
	void paint(QPainter * painter, const QStyleOptionViewItem & option, const QModelIndex & index) const { callbackQSqlRelationalDelegate_Paint(const_cast<MyQSqlRelationalDelegate*>(this), painter, const_cast<QStyleOptionViewItem*>(&option), const_cast<QModelIndex*>(&index)); };
	void setEditorData(QWidget * editor, const QModelIndex & index) const { callbackQSqlRelationalDelegate_SetEditorData(const_cast<MyQSqlRelationalDelegate*>(this), editor, const_cast<QModelIndex*>(&index)); };
	QSize sizeHint(const QStyleOptionViewItem & option, const QModelIndex & index) const { return *static_cast<QSize*>(callbackQSqlRelationalDelegate_SizeHint(const_cast<MyQSqlRelationalDelegate*>(this), const_cast<QStyleOptionViewItem*>(&option), const_cast<QModelIndex*>(&index))); };
	void updateEditorGeometry(QWidget * editor, const QStyleOptionViewItem & option, const QModelIndex & index) const { callbackQSqlRelationalDelegate_UpdateEditorGeometry(const_cast<MyQSqlRelationalDelegate*>(this), editor, const_cast<QStyleOptionViewItem*>(&option), const_cast<QModelIndex*>(&index)); };
	void destroyEditor(QWidget * editor, const QModelIndex & index) const { callbackQSqlRelationalDelegate_DestroyEditor(const_cast<MyQSqlRelationalDelegate*>(this), editor, const_cast<QModelIndex*>(&index)); };
	bool helpEvent(QHelpEvent * event, QAbstractItemView * view, const QStyleOptionViewItem & option, const QModelIndex & index) { return callbackQSqlRelationalDelegate_HelpEvent(this, event, view, const_cast<QStyleOptionViewItem*>(&option), const_cast<QModelIndex*>(&index)) != 0; };
	void timerEvent(QTimerEvent * event) { callbackQSqlRelationalDelegate_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQSqlRelationalDelegate_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSqlRelationalDelegate_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQSqlRelationalDelegate_CustomEvent(this, event); };
	void deleteLater() { callbackQSqlRelationalDelegate_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSqlRelationalDelegate_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQSqlRelationalDelegate_Event(this, e) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSqlRelationalDelegate_MetaObject(const_cast<MyQSqlRelationalDelegate*>(this))); };
};

void* QSqlRelationalDelegate_NewQSqlRelationalDelegate(void* parent)
{
	return new MyQSqlRelationalDelegate(static_cast<QObject*>(parent));
}

void* QSqlRelationalDelegate_CreateEditor(void* ptr, void* parent, void* option, void* index)
{
	return static_cast<QSqlRelationalDelegate*>(ptr)->createEditor(static_cast<QWidget*>(parent), *static_cast<QStyleOptionViewItem*>(option), *static_cast<QModelIndex*>(index));
}

void* QSqlRelationalDelegate_CreateEditorDefault(void* ptr, void* parent, void* option, void* index)
{
	return static_cast<QSqlRelationalDelegate*>(ptr)->QSqlRelationalDelegate::createEditor(static_cast<QWidget*>(parent), *static_cast<QStyleOptionViewItem*>(option), *static_cast<QModelIndex*>(index));
}

void QSqlRelationalDelegate_SetModelData(void* ptr, void* editor, void* model, void* index)
{
	static_cast<QSqlRelationalDelegate*>(ptr)->setModelData(static_cast<QWidget*>(editor), static_cast<QAbstractItemModel*>(model), *static_cast<QModelIndex*>(index));
}

void QSqlRelationalDelegate_SetModelDataDefault(void* ptr, void* editor, void* model, void* index)
{
	static_cast<QSqlRelationalDelegate*>(ptr)->QSqlRelationalDelegate::setModelData(static_cast<QWidget*>(editor), static_cast<QAbstractItemModel*>(model), *static_cast<QModelIndex*>(index));
}

void QSqlRelationalDelegate_DestroyQSqlRelationalDelegate(void* ptr)
{
	static_cast<QSqlRelationalDelegate*>(ptr)->~QSqlRelationalDelegate();
}

void* QSqlRelationalDelegate___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QSqlRelationalDelegate___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSqlRelationalDelegate___children_newList(void* ptr)
{
	return new QList<QObject *>;
}

void* QSqlRelationalDelegate___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QSqlRelationalDelegate___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QSqlRelationalDelegate___dynamicPropertyNames_newList(void* ptr)
{
	return new QList<QByteArray>;
}

void* QSqlRelationalDelegate___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QSqlRelationalDelegate___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSqlRelationalDelegate___findChildren_newList2(void* ptr)
{
	return new QList<QObject*>;
}

void* QSqlRelationalDelegate___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QSqlRelationalDelegate___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSqlRelationalDelegate___findChildren_newList3(void* ptr)
{
	return new QList<QObject*>;
}

void* QSqlRelationalDelegate___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QSqlRelationalDelegate___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSqlRelationalDelegate___findChildren_newList(void* ptr)
{
	return new QList<QObject*>;
}

char QSqlRelationalDelegate_EventFilter(void* ptr, void* editor, void* event)
{
	return static_cast<QSqlRelationalDelegate*>(ptr)->eventFilter(static_cast<QObject*>(editor), static_cast<QEvent*>(event));
}

char QSqlRelationalDelegate_EventFilterDefault(void* ptr, void* editor, void* event)
{
	return static_cast<QSqlRelationalDelegate*>(ptr)->QSqlRelationalDelegate::eventFilter(static_cast<QObject*>(editor), static_cast<QEvent*>(event));
}

char QSqlRelationalDelegate_EditorEvent(void* ptr, void* event, void* model, void* option, void* index)
{
	return static_cast<QSqlRelationalDelegate*>(ptr)->editorEvent(static_cast<QEvent*>(event), static_cast<QAbstractItemModel*>(model), *static_cast<QStyleOptionViewItem*>(option), *static_cast<QModelIndex*>(index));
}

char QSqlRelationalDelegate_EditorEventDefault(void* ptr, void* event, void* model, void* option, void* index)
{
	return static_cast<QSqlRelationalDelegate*>(ptr)->QSqlRelationalDelegate::editorEvent(static_cast<QEvent*>(event), static_cast<QAbstractItemModel*>(model), *static_cast<QStyleOptionViewItem*>(option), *static_cast<QModelIndex*>(index));
}

void QSqlRelationalDelegate_Paint(void* ptr, void* painter, void* option, void* index)
{
	static_cast<QSqlRelationalDelegate*>(ptr)->paint(static_cast<QPainter*>(painter), *static_cast<QStyleOptionViewItem*>(option), *static_cast<QModelIndex*>(index));
}

void QSqlRelationalDelegate_PaintDefault(void* ptr, void* painter, void* option, void* index)
{
	static_cast<QSqlRelationalDelegate*>(ptr)->QSqlRelationalDelegate::paint(static_cast<QPainter*>(painter), *static_cast<QStyleOptionViewItem*>(option), *static_cast<QModelIndex*>(index));
}

void QSqlRelationalDelegate_SetEditorData(void* ptr, void* editor, void* index)
{
	static_cast<QSqlRelationalDelegate*>(ptr)->setEditorData(static_cast<QWidget*>(editor), *static_cast<QModelIndex*>(index));
}

void QSqlRelationalDelegate_SetEditorDataDefault(void* ptr, void* editor, void* index)
{
	static_cast<QSqlRelationalDelegate*>(ptr)->QSqlRelationalDelegate::setEditorData(static_cast<QWidget*>(editor), *static_cast<QModelIndex*>(index));
}

void* QSqlRelationalDelegate_SizeHint(void* ptr, void* option, void* index)
{
	return ({ QSize tmpValue = static_cast<QSqlRelationalDelegate*>(ptr)->sizeHint(*static_cast<QStyleOptionViewItem*>(option), *static_cast<QModelIndex*>(index)); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QSqlRelationalDelegate_SizeHintDefault(void* ptr, void* option, void* index)
{
	return ({ QSize tmpValue = static_cast<QSqlRelationalDelegate*>(ptr)->QSqlRelationalDelegate::sizeHint(*static_cast<QStyleOptionViewItem*>(option), *static_cast<QModelIndex*>(index)); new QSize(tmpValue.width(), tmpValue.height()); });
}

void QSqlRelationalDelegate_UpdateEditorGeometry(void* ptr, void* editor, void* option, void* index)
{
	static_cast<QSqlRelationalDelegate*>(ptr)->updateEditorGeometry(static_cast<QWidget*>(editor), *static_cast<QStyleOptionViewItem*>(option), *static_cast<QModelIndex*>(index));
}

void QSqlRelationalDelegate_UpdateEditorGeometryDefault(void* ptr, void* editor, void* option, void* index)
{
	static_cast<QSqlRelationalDelegate*>(ptr)->QSqlRelationalDelegate::updateEditorGeometry(static_cast<QWidget*>(editor), *static_cast<QStyleOptionViewItem*>(option), *static_cast<QModelIndex*>(index));
}

void QSqlRelationalDelegate_DestroyEditor(void* ptr, void* editor, void* index)
{
	static_cast<QSqlRelationalDelegate*>(ptr)->destroyEditor(static_cast<QWidget*>(editor), *static_cast<QModelIndex*>(index));
}

void QSqlRelationalDelegate_DestroyEditorDefault(void* ptr, void* editor, void* index)
{
	static_cast<QSqlRelationalDelegate*>(ptr)->QSqlRelationalDelegate::destroyEditor(static_cast<QWidget*>(editor), *static_cast<QModelIndex*>(index));
}

char QSqlRelationalDelegate_HelpEvent(void* ptr, void* event, void* view, void* option, void* index)
{
	return static_cast<QSqlRelationalDelegate*>(ptr)->helpEvent(static_cast<QHelpEvent*>(event), static_cast<QAbstractItemView*>(view), *static_cast<QStyleOptionViewItem*>(option), *static_cast<QModelIndex*>(index));
}

char QSqlRelationalDelegate_HelpEventDefault(void* ptr, void* event, void* view, void* option, void* index)
{
	return static_cast<QSqlRelationalDelegate*>(ptr)->QSqlRelationalDelegate::helpEvent(static_cast<QHelpEvent*>(event), static_cast<QAbstractItemView*>(view), *static_cast<QStyleOptionViewItem*>(option), *static_cast<QModelIndex*>(index));
}

void QSqlRelationalDelegate_TimerEvent(void* ptr, void* event)
{
	static_cast<QSqlRelationalDelegate*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QSqlRelationalDelegate_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QSqlRelationalDelegate*>(ptr)->QSqlRelationalDelegate::timerEvent(static_cast<QTimerEvent*>(event));
}

void QSqlRelationalDelegate_ChildEvent(void* ptr, void* event)
{
	static_cast<QSqlRelationalDelegate*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QSqlRelationalDelegate_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QSqlRelationalDelegate*>(ptr)->QSqlRelationalDelegate::childEvent(static_cast<QChildEvent*>(event));
}

void QSqlRelationalDelegate_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QSqlRelationalDelegate*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSqlRelationalDelegate_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QSqlRelationalDelegate*>(ptr)->QSqlRelationalDelegate::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSqlRelationalDelegate_CustomEvent(void* ptr, void* event)
{
	static_cast<QSqlRelationalDelegate*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QSqlRelationalDelegate_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QSqlRelationalDelegate*>(ptr)->QSqlRelationalDelegate::customEvent(static_cast<QEvent*>(event));
}

void QSqlRelationalDelegate_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QSqlRelationalDelegate*>(ptr), "deleteLater");
}

void QSqlRelationalDelegate_DeleteLaterDefault(void* ptr)
{
	static_cast<QSqlRelationalDelegate*>(ptr)->QSqlRelationalDelegate::deleteLater();
}

void QSqlRelationalDelegate_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QSqlRelationalDelegate*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSqlRelationalDelegate_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QSqlRelationalDelegate*>(ptr)->QSqlRelationalDelegate::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QSqlRelationalDelegate_Event(void* ptr, void* e)
{
	return static_cast<QSqlRelationalDelegate*>(ptr)->event(static_cast<QEvent*>(e));
}

char QSqlRelationalDelegate_EventDefault(void* ptr, void* e)
{
	return static_cast<QSqlRelationalDelegate*>(ptr)->QSqlRelationalDelegate::event(static_cast<QEvent*>(e));
}

void* QSqlRelationalDelegate_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QSqlRelationalDelegate*>(ptr)->metaObject());
}

void* QSqlRelationalDelegate_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QSqlRelationalDelegate*>(ptr)->QSqlRelationalDelegate::metaObject());
}

class MyQSqlRelationalTableModel: public QSqlRelationalTableModel
{
public:
	MyQSqlRelationalTableModel(QObject *parent, QSqlDatabase db) : QSqlRelationalTableModel(parent, db) {};
	void clear() { callbackQSqlRelationalTableModel_Clear(this); };
	QVariant data(const QModelIndex & index, int role) const { return *static_cast<QVariant*>(callbackQSqlRelationalTableModel_Data(const_cast<MyQSqlRelationalTableModel*>(this), const_cast<QModelIndex*>(&index), role)); };
	bool insertRowIntoTable(const QSqlRecord & values) { return callbackQSqlRelationalTableModel_InsertRowIntoTable(this, const_cast<QSqlRecord*>(&values)) != 0; };
	QString orderByClause() const { return QString(callbackQSqlRelationalTableModel_OrderByClause(const_cast<MyQSqlRelationalTableModel*>(this))); };
	QSqlTableModel * relationModel(int column) const { return static_cast<QSqlTableModel*>(callbackQSqlRelationalTableModel_RelationModel(const_cast<MyQSqlRelationalTableModel*>(this), column)); };
	bool removeColumns(int column, int count, const QModelIndex & parent) { return callbackQSqlRelationalTableModel_RemoveColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	void revertRow(int row) { callbackQSqlRelationalTableModel_RevertRow(this, row); };
	bool select() { return callbackQSqlRelationalTableModel_Select(this) != 0; };
	QString selectStatement() const { return QString(callbackQSqlRelationalTableModel_SelectStatement(const_cast<MyQSqlRelationalTableModel*>(this))); };
	bool setData(const QModelIndex & index, const QVariant & value, int role) { return callbackQSqlRelationalTableModel_SetData(this, const_cast<QModelIndex*>(&index), const_cast<QVariant*>(&value), role) != 0; };
	void setRelation(int column, const QSqlRelation & relation) { callbackQSqlRelationalTableModel_SetRelation(this, column, const_cast<QSqlRelation*>(&relation)); };
	void setTable(const QString & table) { QByteArray tc3ee13 = table.toUtf8(); QtSql_PackedString tablePacked = { const_cast<char*>(tc3ee13.prepend("WHITESPACE").constData()+10), tc3ee13.size()-10 };callbackQSqlRelationalTableModel_SetTable(this, tablePacked); };
	bool updateRowInTable(int row, const QSqlRecord & values) { return callbackQSqlRelationalTableModel_UpdateRowInTable(this, row, const_cast<QSqlRecord*>(&values)) != 0; };
	 ~MyQSqlRelationalTableModel() { callbackQSqlRelationalTableModel_DestroyQSqlRelationalTableModel(this); };
	bool deleteRowFromTable(int row) { return callbackQSqlRelationalTableModel_DeleteRowFromTable(this, row) != 0; };
	Qt::ItemFlags flags(const QModelIndex & index) const { return static_cast<Qt::ItemFlag>(callbackQSqlRelationalTableModel_Flags(const_cast<MyQSqlRelationalTableModel*>(this), const_cast<QModelIndex*>(&index))); };
	QVariant headerData(int section, Qt::Orientation orientation, int role) const { return *static_cast<QVariant*>(callbackQSqlRelationalTableModel_HeaderData(const_cast<MyQSqlRelationalTableModel*>(this), section, orientation, role)); };
	QModelIndex indexInQuery(const QModelIndex & item) const { return *static_cast<QModelIndex*>(callbackQSqlRelationalTableModel_IndexInQuery(const_cast<MyQSqlRelationalTableModel*>(this), const_cast<QModelIndex*>(&item))); };
	bool insertRows(int row, int count, const QModelIndex & parent) { return callbackQSqlRelationalTableModel_InsertRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool removeRows(int row, int count, const QModelIndex & parent) { return callbackQSqlRelationalTableModel_RemoveRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	void revert() { callbackQSqlRelationalTableModel_Revert(this); };
	void revertAll() { callbackQSqlRelationalTableModel_RevertAll(this); };
	int rowCount(const QModelIndex & parent) const { return callbackQSqlRelationalTableModel_RowCount(const_cast<MyQSqlRelationalTableModel*>(this), const_cast<QModelIndex*>(&parent)); };
	bool selectRow(int row) { return callbackQSqlRelationalTableModel_SelectRow(this, row) != 0; };
	void setEditStrategy(QSqlTableModel::EditStrategy strategy) { callbackQSqlRelationalTableModel_SetEditStrategy(this, strategy); };
	void setFilter(const QString & filter) { QByteArray t4bb4ca = filter.toUtf8(); QtSql_PackedString filterPacked = { const_cast<char*>(t4bb4ca.prepend("WHITESPACE").constData()+10), t4bb4ca.size()-10 };callbackQSqlRelationalTableModel_SetFilter(this, filterPacked); };
	void setSort(int column, Qt::SortOrder order) { callbackQSqlRelationalTableModel_SetSort(this, column, order); };
	void sort(int column, Qt::SortOrder order) { callbackQSqlRelationalTableModel_Sort(this, column, order); };
	bool submit() { return callbackQSqlRelationalTableModel_Submit(this) != 0; };
	bool submitAll() { return callbackQSqlRelationalTableModel_SubmitAll(this) != 0; };
	bool canFetchMore(const QModelIndex & parent) const { return callbackQSqlRelationalTableModel_CanFetchMore(const_cast<MyQSqlRelationalTableModel*>(this), const_cast<QModelIndex*>(&parent)) != 0; };
	int columnCount(const QModelIndex & index) const { return callbackQSqlRelationalTableModel_ColumnCount(const_cast<MyQSqlRelationalTableModel*>(this), const_cast<QModelIndex*>(&index)); };
	void fetchMore(const QModelIndex & parent) { callbackQSqlRelationalTableModel_FetchMore(this, const_cast<QModelIndex*>(&parent)); };
	bool insertColumns(int column, int count, const QModelIndex & parent) { return callbackQSqlRelationalTableModel_InsertColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	void queryChange() { callbackQSqlRelationalTableModel_QueryChange(this); };
	bool setHeaderData(int section, Qt::Orientation orientation, const QVariant & value, int role) { return callbackQSqlRelationalTableModel_SetHeaderData(this, section, orientation, const_cast<QVariant*>(&value), role) != 0; };
	QModelIndex index(int row, int column, const QModelIndex & parent) const { return *static_cast<QModelIndex*>(callbackQSqlRelationalTableModel_Index(const_cast<MyQSqlRelationalTableModel*>(this), row, column, const_cast<QModelIndex*>(&parent))); };
	bool dropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) { return callbackQSqlRelationalTableModel_DropMimeData(this, const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	QModelIndex sibling(int row, int column, const QModelIndex & idx) const { return *static_cast<QModelIndex*>(callbackQSqlRelationalTableModel_Sibling(const_cast<MyQSqlRelationalTableModel*>(this), row, column, const_cast<QModelIndex*>(&idx))); };
	QModelIndex buddy(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackQSqlRelationalTableModel_Buddy(const_cast<MyQSqlRelationalTableModel*>(this), const_cast<QModelIndex*>(&index))); };
	bool canDropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) const { return callbackQSqlRelationalTableModel_CanDropMimeData(const_cast<MyQSqlRelationalTableModel*>(this), const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	bool hasChildren(const QModelIndex & parent) const { return callbackQSqlRelationalTableModel_HasChildren(const_cast<MyQSqlRelationalTableModel*>(this), const_cast<QModelIndex*>(&parent)) != 0; };
	QMap<int, QVariant> itemData(const QModelIndex & index) const { return *static_cast<QMap<int, QVariant>*>(callbackQSqlRelationalTableModel_ItemData(const_cast<MyQSqlRelationalTableModel*>(this), const_cast<QModelIndex*>(&index))); };
	QList<QModelIndex> match(const QModelIndex & start, int role, const QVariant & value, int hits, Qt::MatchFlags flags) const { return *static_cast<QList<QModelIndex>*>(callbackQSqlRelationalTableModel_Match(const_cast<MyQSqlRelationalTableModel*>(this), const_cast<QModelIndex*>(&start), role, const_cast<QVariant*>(&value), hits, flags)); };
	QMimeData * mimeData(const QModelIndexList & indexes) const { return static_cast<QMimeData*>(callbackQSqlRelationalTableModel_MimeData(const_cast<MyQSqlRelationalTableModel*>(this), ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(indexes); QtSql_PackedList { tmpValue, tmpValue->size() }; }))); };
	QStringList mimeTypes() const { return QString(callbackQSqlRelationalTableModel_MimeTypes(const_cast<MyQSqlRelationalTableModel*>(this))).split("|", QString::SkipEmptyParts); };
	bool moveColumns(const QModelIndex & sourceParent, int sourceColumn, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackQSqlRelationalTableModel_MoveColumns(this, const_cast<QModelIndex*>(&sourceParent), sourceColumn, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	bool moveRows(const QModelIndex & sourceParent, int sourceRow, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackQSqlRelationalTableModel_MoveRows(this, const_cast<QModelIndex*>(&sourceParent), sourceRow, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	QModelIndex parent(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackQSqlRelationalTableModel_Parent(const_cast<MyQSqlRelationalTableModel*>(this), const_cast<QModelIndex*>(&index))); };
	void resetInternalData() { callbackQSqlRelationalTableModel_ResetInternalData(this); };
	QHash<int, QByteArray> roleNames() const { return *static_cast<QHash<int, QByteArray>*>(callbackQSqlRelationalTableModel_RoleNames(const_cast<MyQSqlRelationalTableModel*>(this))); };
	bool setItemData(const QModelIndex & index, const QMap<int, QVariant> & roles) { return callbackQSqlRelationalTableModel_SetItemData(this, const_cast<QModelIndex*>(&index), ({ QMap<int, QVariant>* tmpValue = const_cast<QMap<int, QVariant>*>(&roles); QtSql_PackedList { tmpValue, tmpValue->size() }; })) != 0; };
	QSize span(const QModelIndex & index) const { return *static_cast<QSize*>(callbackQSqlRelationalTableModel_Span(const_cast<MyQSqlRelationalTableModel*>(this), const_cast<QModelIndex*>(&index))); };
	Qt::DropActions supportedDragActions() const { return static_cast<Qt::DropAction>(callbackQSqlRelationalTableModel_SupportedDragActions(const_cast<MyQSqlRelationalTableModel*>(this))); };
	Qt::DropActions supportedDropActions() const { return static_cast<Qt::DropAction>(callbackQSqlRelationalTableModel_SupportedDropActions(const_cast<MyQSqlRelationalTableModel*>(this))); };
	void timerEvent(QTimerEvent * event) { callbackQSqlRelationalTableModel_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQSqlRelationalTableModel_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSqlRelationalTableModel_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQSqlRelationalTableModel_CustomEvent(this, event); };
	void deleteLater() { callbackQSqlRelationalTableModel_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSqlRelationalTableModel_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQSqlRelationalTableModel_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSqlRelationalTableModel_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSqlRelationalTableModel_MetaObject(const_cast<MyQSqlRelationalTableModel*>(this))); };
};

void* QSqlRelationalTableModel_NewQSqlRelationalTableModel(void* parent, void* db)
{
	return new MyQSqlRelationalTableModel(static_cast<QObject*>(parent), *static_cast<QSqlDatabase*>(db));
}

void QSqlRelationalTableModel_Clear(void* ptr)
{
	static_cast<QSqlRelationalTableModel*>(ptr)->clear();
}

void QSqlRelationalTableModel_ClearDefault(void* ptr)
{
	static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::clear();
}

void* QSqlRelationalTableModel_Data(void* ptr, void* index, int role)
{
	return new QVariant(static_cast<QSqlRelationalTableModel*>(ptr)->data(*static_cast<QModelIndex*>(index), role));
}

void* QSqlRelationalTableModel_DataDefault(void* ptr, void* index, int role)
{
	return new QVariant(static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::data(*static_cast<QModelIndex*>(index), role));
}

char QSqlRelationalTableModel_InsertRowIntoTable(void* ptr, void* values)
{
	return static_cast<QSqlRelationalTableModel*>(ptr)->insertRowIntoTable(*static_cast<QSqlRecord*>(values));
}

char QSqlRelationalTableModel_InsertRowIntoTableDefault(void* ptr, void* values)
{
	return static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::insertRowIntoTable(*static_cast<QSqlRecord*>(values));
}

struct QtSql_PackedString QSqlRelationalTableModel_OrderByClause(void* ptr)
{
	return ({ QByteArray tfcdde7 = static_cast<QSqlRelationalTableModel*>(ptr)->orderByClause().toUtf8(); QtSql_PackedString { const_cast<char*>(tfcdde7.prepend("WHITESPACE").constData()+10), tfcdde7.size()-10 }; });
}

struct QtSql_PackedString QSqlRelationalTableModel_OrderByClauseDefault(void* ptr)
{
	return ({ QByteArray tbe9e07 = static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::orderByClause().toUtf8(); QtSql_PackedString { const_cast<char*>(tbe9e07.prepend("WHITESPACE").constData()+10), tbe9e07.size()-10 }; });
}

void* QSqlRelationalTableModel_Relation(void* ptr, int column)
{
	return new QSqlRelation(static_cast<QSqlRelationalTableModel*>(ptr)->relation(column));
}

void* QSqlRelationalTableModel_RelationModel(void* ptr, int column)
{
	return static_cast<QSqlRelationalTableModel*>(ptr)->relationModel(column);
}

void* QSqlRelationalTableModel_RelationModelDefault(void* ptr, int column)
{
	return static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::relationModel(column);
}

char QSqlRelationalTableModel_RemoveColumns(void* ptr, int column, int count, void* parent)
{
	return static_cast<QSqlRelationalTableModel*>(ptr)->removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char QSqlRelationalTableModel_RemoveColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

void QSqlRelationalTableModel_RevertRow(void* ptr, int row)
{
	QMetaObject::invokeMethod(static_cast<QSqlRelationalTableModel*>(ptr), "revertRow", Q_ARG(int, row));
}

void QSqlRelationalTableModel_RevertRowDefault(void* ptr, int row)
{
	static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::revertRow(row);
}

char QSqlRelationalTableModel_Select(void* ptr)
{
	return static_cast<QSqlRelationalTableModel*>(ptr)->select();
}

char QSqlRelationalTableModel_SelectDefault(void* ptr)
{
	return static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::select();
}

struct QtSql_PackedString QSqlRelationalTableModel_SelectStatement(void* ptr)
{
	return ({ QByteArray t93d923 = static_cast<QSqlRelationalTableModel*>(ptr)->selectStatement().toUtf8(); QtSql_PackedString { const_cast<char*>(t93d923.prepend("WHITESPACE").constData()+10), t93d923.size()-10 }; });
}

struct QtSql_PackedString QSqlRelationalTableModel_SelectStatementDefault(void* ptr)
{
	return ({ QByteArray t089ec7 = static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::selectStatement().toUtf8(); QtSql_PackedString { const_cast<char*>(t089ec7.prepend("WHITESPACE").constData()+10), t089ec7.size()-10 }; });
}

char QSqlRelationalTableModel_SetData(void* ptr, void* index, void* value, int role)
{
	return static_cast<QSqlRelationalTableModel*>(ptr)->setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
}

char QSqlRelationalTableModel_SetDataDefault(void* ptr, void* index, void* value, int role)
{
	return static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
}

void QSqlRelationalTableModel_SetJoinMode(void* ptr, long long joinMode)
{
	static_cast<QSqlRelationalTableModel*>(ptr)->setJoinMode(static_cast<QSqlRelationalTableModel::JoinMode>(joinMode));
}

void QSqlRelationalTableModel_SetRelation(void* ptr, int column, void* relation)
{
	static_cast<QSqlRelationalTableModel*>(ptr)->setRelation(column, *static_cast<QSqlRelation*>(relation));
}

void QSqlRelationalTableModel_SetRelationDefault(void* ptr, int column, void* relation)
{
	static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::setRelation(column, *static_cast<QSqlRelation*>(relation));
}

void QSqlRelationalTableModel_SetTable(void* ptr, char* table)
{
	static_cast<QSqlRelationalTableModel*>(ptr)->setTable(QString(table));
}

void QSqlRelationalTableModel_SetTableDefault(void* ptr, char* table)
{
	static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::setTable(QString(table));
}

char QSqlRelationalTableModel_UpdateRowInTable(void* ptr, int row, void* values)
{
	return static_cast<QSqlRelationalTableModel*>(ptr)->updateRowInTable(row, *static_cast<QSqlRecord*>(values));
}

char QSqlRelationalTableModel_UpdateRowInTableDefault(void* ptr, int row, void* values)
{
	return static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::updateRowInTable(row, *static_cast<QSqlRecord*>(values));
}

void QSqlRelationalTableModel_DestroyQSqlRelationalTableModel(void* ptr)
{
	static_cast<QSqlRelationalTableModel*>(ptr)->~QSqlRelationalTableModel();
}

void QSqlRelationalTableModel_DestroyQSqlRelationalTableModelDefault(void* ptr)
{

}

int QSqlRelationalTableModel_____itemData_keyList_atList(void* ptr, int i)
{
	return static_cast<QList<int>*>(ptr)->at(i);
}

void QSqlRelationalTableModel_____itemData_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* QSqlRelationalTableModel_____itemData_keyList_newList(void* ptr)
{
	return new QList<int>;
}

int QSqlRelationalTableModel_____roleNames_keyList_atList(void* ptr, int i)
{
	return static_cast<QList<int>*>(ptr)->at(i);
}

void QSqlRelationalTableModel_____roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* QSqlRelationalTableModel_____roleNames_keyList_newList(void* ptr)
{
	return new QList<int>;
}

int QSqlRelationalTableModel_____setItemData_keyList_atList(void* ptr, int i)
{
	return static_cast<QList<int>*>(ptr)->at(i);
}

void QSqlRelationalTableModel_____setItemData_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* QSqlRelationalTableModel_____setItemData_keyList_newList(void* ptr)
{
	return new QList<int>;
}

void* QSqlRelationalTableModel___changePersistentIndexList_from_atList(void* ptr, int i)
{
	return new QModelIndex(static_cast<QList<QModelIndex>*>(ptr)->at(i));
}

void QSqlRelationalTableModel___changePersistentIndexList_from_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* QSqlRelationalTableModel___changePersistentIndexList_from_newList(void* ptr)
{
	return new QList<QModelIndex>;
}

void* QSqlRelationalTableModel___changePersistentIndexList_to_atList(void* ptr, int i)
{
	return new QModelIndex(static_cast<QList<QModelIndex>*>(ptr)->at(i));
}

void QSqlRelationalTableModel___changePersistentIndexList_to_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* QSqlRelationalTableModel___changePersistentIndexList_to_newList(void* ptr)
{
	return new QList<QModelIndex>;
}

int QSqlRelationalTableModel___dataChanged_roles_atList(void* ptr, int i)
{
	return static_cast<QVector<int>*>(ptr)->at(i);
}

void QSqlRelationalTableModel___dataChanged_roles_setList(void* ptr, int i)
{
	static_cast<QVector<int>*>(ptr)->append(i);
}

void* QSqlRelationalTableModel___dataChanged_roles_newList(void* ptr)
{
	return new QVector<int>;
}

void* QSqlRelationalTableModel___itemData_atList(void* ptr, int i)
{
	return new QVariant(static_cast<QMap<int, QVariant>*>(ptr)->value(i));
}

void QSqlRelationalTableModel___itemData_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* QSqlRelationalTableModel___itemData_newList(void* ptr)
{
	return new QMap<int, QVariant>;
}

struct QtSql_PackedList QSqlRelationalTableModel___itemData_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); QtSql_PackedList { tmpValue, tmpValue->size() }; });
}

void* QSqlRelationalTableModel___layoutAboutToBeChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i));
}

void QSqlRelationalTableModel___layoutAboutToBeChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* QSqlRelationalTableModel___layoutAboutToBeChanged_parents_newList(void* ptr)
{
	return new QList<QPersistentModelIndex>;
}

void* QSqlRelationalTableModel___layoutChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i));
}

void QSqlRelationalTableModel___layoutChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* QSqlRelationalTableModel___layoutChanged_parents_newList(void* ptr)
{
	return new QList<QPersistentModelIndex>;
}

void* QSqlRelationalTableModel___match_atList(void* ptr, int i)
{
	return new QModelIndex(static_cast<QList<QModelIndex>*>(ptr)->at(i));
}

void QSqlRelationalTableModel___match_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* QSqlRelationalTableModel___match_newList(void* ptr)
{
	return new QList<QModelIndex>;
}

void* QSqlRelationalTableModel___mimeData_indexes_atList(void* ptr, int i)
{
	return new QModelIndex(static_cast<QList<QModelIndex>*>(ptr)->at(i));
}

void QSqlRelationalTableModel___mimeData_indexes_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* QSqlRelationalTableModel___mimeData_indexes_newList(void* ptr)
{
	return new QList<QModelIndex>;
}

void* QSqlRelationalTableModel___persistentIndexList_atList(void* ptr, int i)
{
	return new QModelIndex(static_cast<QList<QModelIndex>*>(ptr)->at(i));
}

void QSqlRelationalTableModel___persistentIndexList_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* QSqlRelationalTableModel___persistentIndexList_newList(void* ptr)
{
	return new QList<QModelIndex>;
}

void* QSqlRelationalTableModel___roleNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QHash<int, QByteArray>*>(ptr)->value(i));
}

void QSqlRelationalTableModel___roleNames_setList(void* ptr, int key, void* i)
{
	static_cast<QHash<int, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* QSqlRelationalTableModel___roleNames_newList(void* ptr)
{
	return new QHash<int, QByteArray>;
}

struct QtSql_PackedList QSqlRelationalTableModel___roleNames_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QHash<int, QByteArray>*>(ptr)->keys()); QtSql_PackedList { tmpValue, tmpValue->size() }; });
}

void* QSqlRelationalTableModel___setItemData_roles_atList(void* ptr, int i)
{
	return new QVariant(static_cast<QMap<int, QVariant>*>(ptr)->value(i));
}

void QSqlRelationalTableModel___setItemData_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* QSqlRelationalTableModel___setItemData_roles_newList(void* ptr)
{
	return new QMap<int, QVariant>;
}

struct QtSql_PackedList QSqlRelationalTableModel___setItemData_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); QtSql_PackedList { tmpValue, tmpValue->size() }; });
}

int QSqlRelationalTableModel_____doSetRoleNames_keyList_atList(void* ptr, int i)
{
	return static_cast<QList<int>*>(ptr)->at(i);
}

void QSqlRelationalTableModel_____doSetRoleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* QSqlRelationalTableModel_____doSetRoleNames_keyList_newList(void* ptr)
{
	return new QList<int>;
}

int QSqlRelationalTableModel_____setRoleNames_keyList_atList(void* ptr, int i)
{
	return static_cast<QList<int>*>(ptr)->at(i);
}

void QSqlRelationalTableModel_____setRoleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* QSqlRelationalTableModel_____setRoleNames_keyList_newList(void* ptr)
{
	return new QList<int>;
}

void* QSqlRelationalTableModel___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QSqlRelationalTableModel___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSqlRelationalTableModel___children_newList(void* ptr)
{
	return new QList<QObject *>;
}

void* QSqlRelationalTableModel___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QSqlRelationalTableModel___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QSqlRelationalTableModel___dynamicPropertyNames_newList(void* ptr)
{
	return new QList<QByteArray>;
}

void* QSqlRelationalTableModel___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QSqlRelationalTableModel___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSqlRelationalTableModel___findChildren_newList2(void* ptr)
{
	return new QList<QObject*>;
}

void* QSqlRelationalTableModel___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QSqlRelationalTableModel___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSqlRelationalTableModel___findChildren_newList3(void* ptr)
{
	return new QList<QObject*>;
}

void* QSqlRelationalTableModel___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QSqlRelationalTableModel___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSqlRelationalTableModel___findChildren_newList(void* ptr)
{
	return new QList<QObject*>;
}

char QSqlRelationalTableModel_DeleteRowFromTable(void* ptr, int row)
{
	return static_cast<QSqlRelationalTableModel*>(ptr)->deleteRowFromTable(row);
}

char QSqlRelationalTableModel_DeleteRowFromTableDefault(void* ptr, int row)
{
	return static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::deleteRowFromTable(row);
}

long long QSqlRelationalTableModel_Flags(void* ptr, void* index)
{
	return static_cast<QSqlRelationalTableModel*>(ptr)->flags(*static_cast<QModelIndex*>(index));
}

long long QSqlRelationalTableModel_FlagsDefault(void* ptr, void* index)
{
	return static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::flags(*static_cast<QModelIndex*>(index));
}

void* QSqlRelationalTableModel_HeaderData(void* ptr, int section, long long orientation, int role)
{
	return new QVariant(static_cast<QSqlRelationalTableModel*>(ptr)->headerData(section, static_cast<Qt::Orientation>(orientation), role));
}

void* QSqlRelationalTableModel_HeaderDataDefault(void* ptr, int section, long long orientation, int role)
{
	return new QVariant(static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::headerData(section, static_cast<Qt::Orientation>(orientation), role));
}

void* QSqlRelationalTableModel_IndexInQuery(void* ptr, void* item)
{
	return new QModelIndex(static_cast<QSqlRelationalTableModel*>(ptr)->indexInQuery(*static_cast<QModelIndex*>(item)));
}

void* QSqlRelationalTableModel_IndexInQueryDefault(void* ptr, void* item)
{
	return new QModelIndex(static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::indexInQuery(*static_cast<QModelIndex*>(item)));
}

char QSqlRelationalTableModel_InsertRows(void* ptr, int row, int count, void* parent)
{
	return static_cast<QSqlRelationalTableModel*>(ptr)->insertRows(row, count, *static_cast<QModelIndex*>(parent));
}

char QSqlRelationalTableModel_InsertRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::insertRows(row, count, *static_cast<QModelIndex*>(parent));
}

char QSqlRelationalTableModel_RemoveRows(void* ptr, int row, int count, void* parent)
{
	return static_cast<QSqlRelationalTableModel*>(ptr)->removeRows(row, count, *static_cast<QModelIndex*>(parent));
}

char QSqlRelationalTableModel_RemoveRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::removeRows(row, count, *static_cast<QModelIndex*>(parent));
}

void QSqlRelationalTableModel_Revert(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QSqlRelationalTableModel*>(ptr), "revert");
}

void QSqlRelationalTableModel_RevertDefault(void* ptr)
{
	static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::revert();
}

void QSqlRelationalTableModel_RevertAll(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QSqlRelationalTableModel*>(ptr), "revertAll");
}

void QSqlRelationalTableModel_RevertAllDefault(void* ptr)
{
	static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::revertAll();
}

int QSqlRelationalTableModel_RowCount(void* ptr, void* parent)
{
	return static_cast<QSqlRelationalTableModel*>(ptr)->rowCount(*static_cast<QModelIndex*>(parent));
}

int QSqlRelationalTableModel_RowCountDefault(void* ptr, void* parent)
{
	return static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::rowCount(*static_cast<QModelIndex*>(parent));
}

char QSqlRelationalTableModel_SelectRow(void* ptr, int row)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QSqlRelationalTableModel*>(ptr), "selectRow", Q_RETURN_ARG(bool, returnArg), Q_ARG(int, row));
	return returnArg;
}

char QSqlRelationalTableModel_SelectRowDefault(void* ptr, int row)
{
	return static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::selectRow(row);
}

void QSqlRelationalTableModel_SetEditStrategy(void* ptr, long long strategy)
{
	static_cast<QSqlRelationalTableModel*>(ptr)->setEditStrategy(static_cast<QSqlTableModel::EditStrategy>(strategy));
}

void QSqlRelationalTableModel_SetEditStrategyDefault(void* ptr, long long strategy)
{
	static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::setEditStrategy(static_cast<QSqlTableModel::EditStrategy>(strategy));
}

void QSqlRelationalTableModel_SetFilter(void* ptr, char* filter)
{
	static_cast<QSqlRelationalTableModel*>(ptr)->setFilter(QString(filter));
}

void QSqlRelationalTableModel_SetFilterDefault(void* ptr, char* filter)
{
	static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::setFilter(QString(filter));
}

void QSqlRelationalTableModel_SetSort(void* ptr, int column, long long order)
{
	static_cast<QSqlRelationalTableModel*>(ptr)->setSort(column, static_cast<Qt::SortOrder>(order));
}

void QSqlRelationalTableModel_SetSortDefault(void* ptr, int column, long long order)
{
	static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::setSort(column, static_cast<Qt::SortOrder>(order));
}

void QSqlRelationalTableModel_Sort(void* ptr, int column, long long order)
{
	static_cast<QSqlRelationalTableModel*>(ptr)->sort(column, static_cast<Qt::SortOrder>(order));
}

void QSqlRelationalTableModel_SortDefault(void* ptr, int column, long long order)
{
	static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::sort(column, static_cast<Qt::SortOrder>(order));
}

char QSqlRelationalTableModel_Submit(void* ptr)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QSqlRelationalTableModel*>(ptr), "submit", Q_RETURN_ARG(bool, returnArg));
	return returnArg;
}

char QSqlRelationalTableModel_SubmitDefault(void* ptr)
{
	return static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::submit();
}

char QSqlRelationalTableModel_SubmitAll(void* ptr)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QSqlRelationalTableModel*>(ptr), "submitAll", Q_RETURN_ARG(bool, returnArg));
	return returnArg;
}

char QSqlRelationalTableModel_SubmitAllDefault(void* ptr)
{
	return static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::submitAll();
}

char QSqlRelationalTableModel_CanFetchMore(void* ptr, void* parent)
{
	return static_cast<QSqlRelationalTableModel*>(ptr)->canFetchMore(*static_cast<QModelIndex*>(parent));
}

char QSqlRelationalTableModel_CanFetchMoreDefault(void* ptr, void* parent)
{
	return static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::canFetchMore(*static_cast<QModelIndex*>(parent));
}

int QSqlRelationalTableModel_ColumnCount(void* ptr, void* index)
{
	return static_cast<QSqlRelationalTableModel*>(ptr)->columnCount(*static_cast<QModelIndex*>(index));
}

int QSqlRelationalTableModel_ColumnCountDefault(void* ptr, void* index)
{
	return static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::columnCount(*static_cast<QModelIndex*>(index));
}

void QSqlRelationalTableModel_FetchMore(void* ptr, void* parent)
{
	static_cast<QSqlRelationalTableModel*>(ptr)->fetchMore(*static_cast<QModelIndex*>(parent));
}

void QSqlRelationalTableModel_FetchMoreDefault(void* ptr, void* parent)
{
	static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::fetchMore(*static_cast<QModelIndex*>(parent));
}

char QSqlRelationalTableModel_InsertColumns(void* ptr, int column, int count, void* parent)
{
	return static_cast<QSqlRelationalTableModel*>(ptr)->insertColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char QSqlRelationalTableModel_InsertColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::insertColumns(column, count, *static_cast<QModelIndex*>(parent));
}

void QSqlRelationalTableModel_QueryChange(void* ptr)
{
	static_cast<QSqlRelationalTableModel*>(ptr)->queryChange();
}

void QSqlRelationalTableModel_QueryChangeDefault(void* ptr)
{
	static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::queryChange();
}

char QSqlRelationalTableModel_SetHeaderData(void* ptr, int section, long long orientation, void* value, int role)
{
	return static_cast<QSqlRelationalTableModel*>(ptr)->setHeaderData(section, static_cast<Qt::Orientation>(orientation), *static_cast<QVariant*>(value), role);
}

char QSqlRelationalTableModel_SetHeaderDataDefault(void* ptr, int section, long long orientation, void* value, int role)
{
	return static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::setHeaderData(section, static_cast<Qt::Orientation>(orientation), *static_cast<QVariant*>(value), role);
}

void* QSqlRelationalTableModel_Index(void* ptr, int row, int column, void* parent)
{
	return new QModelIndex(static_cast<QSqlRelationalTableModel*>(ptr)->index(row, column, *static_cast<QModelIndex*>(parent)));
}

void* QSqlRelationalTableModel_IndexDefault(void* ptr, int row, int column, void* parent)
{
	return new QModelIndex(static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::index(row, column, *static_cast<QModelIndex*>(parent)));
}

char QSqlRelationalTableModel_DropMimeData(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<QSqlRelationalTableModel*>(ptr)->dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

char QSqlRelationalTableModel_DropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

void* QSqlRelationalTableModel_Sibling(void* ptr, int row, int column, void* idx)
{
	return new QModelIndex(static_cast<QSqlRelationalTableModel*>(ptr)->sibling(row, column, *static_cast<QModelIndex*>(idx)));
}

void* QSqlRelationalTableModel_SiblingDefault(void* ptr, int row, int column, void* idx)
{
	return new QModelIndex(static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::sibling(row, column, *static_cast<QModelIndex*>(idx)));
}

void* QSqlRelationalTableModel_Buddy(void* ptr, void* index)
{
	return new QModelIndex(static_cast<QSqlRelationalTableModel*>(ptr)->buddy(*static_cast<QModelIndex*>(index)));
}

void* QSqlRelationalTableModel_BuddyDefault(void* ptr, void* index)
{
	return new QModelIndex(static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::buddy(*static_cast<QModelIndex*>(index)));
}

char QSqlRelationalTableModel_CanDropMimeData(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<QSqlRelationalTableModel*>(ptr)->canDropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

char QSqlRelationalTableModel_CanDropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::canDropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

char QSqlRelationalTableModel_HasChildren(void* ptr, void* parent)
{
	return static_cast<QSqlRelationalTableModel*>(ptr)->hasChildren(*static_cast<QModelIndex*>(parent));
}

char QSqlRelationalTableModel_HasChildrenDefault(void* ptr, void* parent)
{
	return static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::hasChildren(*static_cast<QModelIndex*>(parent));
}

struct QtSql_PackedList QSqlRelationalTableModel_ItemData(void* ptr, void* index)
{
	return ({ QMap<int, QVariant>* tmpValue = new QMap<int, QVariant>(static_cast<QSqlRelationalTableModel*>(ptr)->itemData(*static_cast<QModelIndex*>(index))); QtSql_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtSql_PackedList QSqlRelationalTableModel_ItemDataDefault(void* ptr, void* index)
{
	return ({ QMap<int, QVariant>* tmpValue = new QMap<int, QVariant>(static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::itemData(*static_cast<QModelIndex*>(index))); QtSql_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtSql_PackedList QSqlRelationalTableModel_Match(void* ptr, void* start, int role, void* value, int hits, long long flags)
{
	return ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(static_cast<QSqlRelationalTableModel*>(ptr)->match(*static_cast<QModelIndex*>(start), role, *static_cast<QVariant*>(value), hits, static_cast<Qt::MatchFlag>(flags))); QtSql_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtSql_PackedList QSqlRelationalTableModel_MatchDefault(void* ptr, void* start, int role, void* value, int hits, long long flags)
{
	return ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::match(*static_cast<QModelIndex*>(start), role, *static_cast<QVariant*>(value), hits, static_cast<Qt::MatchFlag>(flags))); QtSql_PackedList { tmpValue, tmpValue->size() }; });
}

void* QSqlRelationalTableModel_MimeData(void* ptr, void* indexes)
{
	return static_cast<QSqlRelationalTableModel*>(ptr)->mimeData(*static_cast<QList<QModelIndex>*>(indexes));
}

void* QSqlRelationalTableModel_MimeDataDefault(void* ptr, void* indexes)
{
	return static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::mimeData(*static_cast<QList<QModelIndex>*>(indexes));
}

struct QtSql_PackedString QSqlRelationalTableModel_MimeTypes(void* ptr)
{
	return ({ QByteArray te5a031 = static_cast<QSqlRelationalTableModel*>(ptr)->mimeTypes().join("|").toUtf8(); QtSql_PackedString { const_cast<char*>(te5a031.prepend("WHITESPACE").constData()+10), te5a031.size()-10 }; });
}

struct QtSql_PackedString QSqlRelationalTableModel_MimeTypesDefault(void* ptr)
{
	return ({ QByteArray t385a7c = static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::mimeTypes().join("|").toUtf8(); QtSql_PackedString { const_cast<char*>(t385a7c.prepend("WHITESPACE").constData()+10), t385a7c.size()-10 }; });
}

char QSqlRelationalTableModel_MoveColumns(void* ptr, void* sourceParent, int sourceColumn, int count, void* destinationParent, int destinationChild)
{
	return static_cast<QSqlRelationalTableModel*>(ptr)->moveColumns(*static_cast<QModelIndex*>(sourceParent), sourceColumn, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

char QSqlRelationalTableModel_MoveColumnsDefault(void* ptr, void* sourceParent, int sourceColumn, int count, void* destinationParent, int destinationChild)
{
	return static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::moveColumns(*static_cast<QModelIndex*>(sourceParent), sourceColumn, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

char QSqlRelationalTableModel_MoveRows(void* ptr, void* sourceParent, int sourceRow, int count, void* destinationParent, int destinationChild)
{
	return static_cast<QSqlRelationalTableModel*>(ptr)->moveRows(*static_cast<QModelIndex*>(sourceParent), sourceRow, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

char QSqlRelationalTableModel_MoveRowsDefault(void* ptr, void* sourceParent, int sourceRow, int count, void* destinationParent, int destinationChild)
{
	return static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::moveRows(*static_cast<QModelIndex*>(sourceParent), sourceRow, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

void* QSqlRelationalTableModel_Parent(void* ptr, void* index)
{
	return new QModelIndex(static_cast<QSqlRelationalTableModel*>(ptr)->parent(*static_cast<QModelIndex*>(index)));
}

void QSqlRelationalTableModel_ResetInternalData(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QSqlRelationalTableModel*>(ptr), "resetInternalData");
}

void QSqlRelationalTableModel_ResetInternalDataDefault(void* ptr)
{
	static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::resetInternalData();
}

struct QtSql_PackedList QSqlRelationalTableModel_RoleNames(void* ptr)
{
	return ({ QHash<int, QByteArray>* tmpValue = new QHash<int, QByteArray>(static_cast<QSqlRelationalTableModel*>(ptr)->roleNames()); QtSql_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtSql_PackedList QSqlRelationalTableModel_RoleNamesDefault(void* ptr)
{
	return ({ QHash<int, QByteArray>* tmpValue = new QHash<int, QByteArray>(static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::roleNames()); QtSql_PackedList { tmpValue, tmpValue->size() }; });
}

char QSqlRelationalTableModel_SetItemData(void* ptr, void* index, void* roles)
{
	return static_cast<QSqlRelationalTableModel*>(ptr)->setItemData(*static_cast<QModelIndex*>(index), *static_cast<QMap<int, QVariant>*>(roles));
}

char QSqlRelationalTableModel_SetItemDataDefault(void* ptr, void* index, void* roles)
{
	return static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::setItemData(*static_cast<QModelIndex*>(index), *static_cast<QMap<int, QVariant>*>(roles));
}

void* QSqlRelationalTableModel_Span(void* ptr, void* index)
{
	return ({ QSize tmpValue = static_cast<QSqlRelationalTableModel*>(ptr)->span(*static_cast<QModelIndex*>(index)); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QSqlRelationalTableModel_SpanDefault(void* ptr, void* index)
{
	return ({ QSize tmpValue = static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::span(*static_cast<QModelIndex*>(index)); new QSize(tmpValue.width(), tmpValue.height()); });
}

long long QSqlRelationalTableModel_SupportedDragActions(void* ptr)
{
	return static_cast<QSqlRelationalTableModel*>(ptr)->supportedDragActions();
}

long long QSqlRelationalTableModel_SupportedDragActionsDefault(void* ptr)
{
	return static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::supportedDragActions();
}

long long QSqlRelationalTableModel_SupportedDropActions(void* ptr)
{
	return static_cast<QSqlRelationalTableModel*>(ptr)->supportedDropActions();
}

long long QSqlRelationalTableModel_SupportedDropActionsDefault(void* ptr)
{
	return static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::supportedDropActions();
}

void QSqlRelationalTableModel_TimerEvent(void* ptr, void* event)
{
	static_cast<QSqlRelationalTableModel*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QSqlRelationalTableModel_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::timerEvent(static_cast<QTimerEvent*>(event));
}

void QSqlRelationalTableModel_ChildEvent(void* ptr, void* event)
{
	static_cast<QSqlRelationalTableModel*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QSqlRelationalTableModel_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::childEvent(static_cast<QChildEvent*>(event));
}

void QSqlRelationalTableModel_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QSqlRelationalTableModel*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSqlRelationalTableModel_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSqlRelationalTableModel_CustomEvent(void* ptr, void* event)
{
	static_cast<QSqlRelationalTableModel*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QSqlRelationalTableModel_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::customEvent(static_cast<QEvent*>(event));
}

void QSqlRelationalTableModel_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QSqlRelationalTableModel*>(ptr), "deleteLater");
}

void QSqlRelationalTableModel_DeleteLaterDefault(void* ptr)
{
	static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::deleteLater();
}

void QSqlRelationalTableModel_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QSqlRelationalTableModel*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSqlRelationalTableModel_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QSqlRelationalTableModel_Event(void* ptr, void* e)
{
	return static_cast<QSqlRelationalTableModel*>(ptr)->event(static_cast<QEvent*>(e));
}

char QSqlRelationalTableModel_EventDefault(void* ptr, void* e)
{
	return static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::event(static_cast<QEvent*>(e));
}

char QSqlRelationalTableModel_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QSqlRelationalTableModel*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QSqlRelationalTableModel_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QSqlRelationalTableModel_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QSqlRelationalTableModel*>(ptr)->metaObject());
}

void* QSqlRelationalTableModel_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::metaObject());
}

class MyQSqlResult: public QSqlResult
{
public:
	MyQSqlResult(const QSqlDriver *db) : QSqlResult(db) {};
	void bindValue(const QString & placeholder, const QVariant & val, QSql::ParamType paramType) { QByteArray tff5543 = placeholder.toUtf8(); QtSql_PackedString placeholderPacked = { const_cast<char*>(tff5543.prepend("WHITESPACE").constData()+10), tff5543.size()-10 };callbackQSqlResult_BindValue2(this, placeholderPacked, const_cast<QVariant*>(&val), paramType); };
	void bindValue(int index, const QVariant & val, QSql::ParamType paramType) { callbackQSqlResult_BindValue(this, index, const_cast<QVariant*>(&val), paramType); };
	QVariant data(int index) { return *static_cast<QVariant*>(callbackQSqlResult_Data(this, index)); };
	bool exec() { return callbackQSqlResult_Exec(this) != 0; };
	bool fetch(int index) { return callbackQSqlResult_Fetch(this, index) != 0; };
	bool fetchFirst() { return callbackQSqlResult_FetchFirst(this) != 0; };
	bool fetchLast() { return callbackQSqlResult_FetchLast(this) != 0; };
	bool fetchNext() { return callbackQSqlResult_FetchNext(this) != 0; };
	bool fetchPrevious() { return callbackQSqlResult_FetchPrevious(this) != 0; };
	QVariant handle() const { return *static_cast<QVariant*>(callbackQSqlResult_Handle(const_cast<MyQSqlResult*>(this))); };
	bool isNull(int index) { return callbackQSqlResult_IsNull(this, index) != 0; };
	QVariant lastInsertId() const { return *static_cast<QVariant*>(callbackQSqlResult_LastInsertId(const_cast<MyQSqlResult*>(this))); };
	int numRowsAffected() { return callbackQSqlResult_NumRowsAffected(this); };
	bool prepare(const QString & query) { QByteArray t7cd914 = query.toUtf8(); QtSql_PackedString queryPacked = { const_cast<char*>(t7cd914.prepend("WHITESPACE").constData()+10), t7cd914.size()-10 };return callbackQSqlResult_Prepare(this, queryPacked) != 0; };
	QSqlRecord record() const { return *static_cast<QSqlRecord*>(callbackQSqlResult_Record(const_cast<MyQSqlResult*>(this))); };
	bool reset(const QString & query) { QByteArray t7cd914 = query.toUtf8(); QtSql_PackedString queryPacked = { const_cast<char*>(t7cd914.prepend("WHITESPACE").constData()+10), t7cd914.size()-10 };return callbackQSqlResult_Reset(this, queryPacked) != 0; };
	bool savePrepare(const QString & query) { QByteArray t7cd914 = query.toUtf8(); QtSql_PackedString queryPacked = { const_cast<char*>(t7cd914.prepend("WHITESPACE").constData()+10), t7cd914.size()-10 };return callbackQSqlResult_SavePrepare(this, queryPacked) != 0; };
	void setActive(bool active) { callbackQSqlResult_SetActive(this, active); };
	void setAt(int index) { callbackQSqlResult_SetAt(this, index); };
	void setForwardOnly(bool forward) { callbackQSqlResult_SetForwardOnly(this, forward); };
	void setLastError(const QSqlError & error) { callbackQSqlResult_SetLastError(this, const_cast<QSqlError*>(&error)); };
	void setQuery(const QString & query) { QByteArray t7cd914 = query.toUtf8(); QtSql_PackedString queryPacked = { const_cast<char*>(t7cd914.prepend("WHITESPACE").constData()+10), t7cd914.size()-10 };callbackQSqlResult_SetQuery(this, queryPacked); };
	void setSelect(bool sele) { callbackQSqlResult_SetSelect(this, sele); };
	int size() { return callbackQSqlResult_Size(this); };
	 ~MyQSqlResult() { callbackQSqlResult_DestroyQSqlResult(this); };
};

void* QSqlResult_NewQSqlResult(void* db)
{
	return new MyQSqlResult(static_cast<QSqlDriver*>(db));
}

void QSqlResult_AddBindValue(void* ptr, void* val, long long paramType)
{
	static_cast<QSqlResult*>(ptr)->addBindValue(*static_cast<QVariant*>(val), static_cast<QSql::ParamTypeFlag>(paramType));
}

int QSqlResult_At(void* ptr)
{
	return static_cast<QSqlResult*>(ptr)->at();
}

void QSqlResult_BindValue2(void* ptr, char* placeholder, void* val, long long paramType)
{
	static_cast<QSqlResult*>(ptr)->bindValue(QString(placeholder), *static_cast<QVariant*>(val), static_cast<QSql::ParamTypeFlag>(paramType));
}

void QSqlResult_BindValue2Default(void* ptr, char* placeholder, void* val, long long paramType)
{
	static_cast<QSqlResult*>(ptr)->QSqlResult::bindValue(QString(placeholder), *static_cast<QVariant*>(val), static_cast<QSql::ParamTypeFlag>(paramType));
}

void QSqlResult_BindValue(void* ptr, int index, void* val, long long paramType)
{
	static_cast<QSqlResult*>(ptr)->bindValue(index, *static_cast<QVariant*>(val), static_cast<QSql::ParamTypeFlag>(paramType));
}

void QSqlResult_BindValueDefault(void* ptr, int index, void* val, long long paramType)
{
	static_cast<QSqlResult*>(ptr)->QSqlResult::bindValue(index, *static_cast<QVariant*>(val), static_cast<QSql::ParamTypeFlag>(paramType));
}

long long QSqlResult_BindValueType2(void* ptr, char* placeholder)
{
	return static_cast<QSqlResult*>(ptr)->bindValueType(QString(placeholder));
}

long long QSqlResult_BindValueType(void* ptr, int index)
{
	return static_cast<QSqlResult*>(ptr)->bindValueType(index);
}

long long QSqlResult_BindingSyntax(void* ptr)
{
	return static_cast<QSqlResult*>(ptr)->bindingSyntax();
}

void* QSqlResult_BoundValue2(void* ptr, char* placeholder)
{
	return new QVariant(static_cast<QSqlResult*>(ptr)->boundValue(QString(placeholder)));
}

void* QSqlResult_BoundValue(void* ptr, int index)
{
	return new QVariant(static_cast<QSqlResult*>(ptr)->boundValue(index));
}

int QSqlResult_BoundValueCount(void* ptr)
{
	return static_cast<QSqlResult*>(ptr)->boundValueCount();
}

struct QtSql_PackedString QSqlResult_BoundValueName(void* ptr, int index)
{
	return ({ QByteArray t669294 = static_cast<QSqlResult*>(ptr)->boundValueName(index).toUtf8(); QtSql_PackedString { const_cast<char*>(t669294.prepend("WHITESPACE").constData()+10), t669294.size()-10 }; });
}

struct QtSql_PackedList QSqlResult_BoundValues(void* ptr)
{
	return ({ QVector<QVariant>* tmpValue = new QVector<QVariant>(static_cast<QSqlResult*>(ptr)->boundValues()); QtSql_PackedList { tmpValue, tmpValue->size() }; });
}

void QSqlResult_Clear(void* ptr)
{
	static_cast<QSqlResult*>(ptr)->clear();
}

void* QSqlResult_Data(void* ptr, int index)
{
	return new QVariant(static_cast<QSqlResult*>(ptr)->data(index));
}

void* QSqlResult_Driver(void* ptr)
{
	return const_cast<QSqlDriver*>(static_cast<QSqlResult*>(ptr)->driver());
}

char QSqlResult_Exec(void* ptr)
{
	return static_cast<QSqlResult*>(ptr)->exec();
}

char QSqlResult_ExecDefault(void* ptr)
{
	return static_cast<QSqlResult*>(ptr)->QSqlResult::exec();
}

struct QtSql_PackedString QSqlResult_ExecutedQuery(void* ptr)
{
	return ({ QByteArray t35c395 = static_cast<QSqlResult*>(ptr)->executedQuery().toUtf8(); QtSql_PackedString { const_cast<char*>(t35c395.prepend("WHITESPACE").constData()+10), t35c395.size()-10 }; });
}

char QSqlResult_Fetch(void* ptr, int index)
{
	return static_cast<QSqlResult*>(ptr)->fetch(index);
}

char QSqlResult_FetchFirst(void* ptr)
{
	return static_cast<QSqlResult*>(ptr)->fetchFirst();
}

char QSqlResult_FetchLast(void* ptr)
{
	return static_cast<QSqlResult*>(ptr)->fetchLast();
}

char QSqlResult_FetchNext(void* ptr)
{
	return static_cast<QSqlResult*>(ptr)->fetchNext();
}

char QSqlResult_FetchNextDefault(void* ptr)
{
	return static_cast<QSqlResult*>(ptr)->QSqlResult::fetchNext();
}

char QSqlResult_FetchPrevious(void* ptr)
{
	return static_cast<QSqlResult*>(ptr)->fetchPrevious();
}

char QSqlResult_FetchPreviousDefault(void* ptr)
{
	return static_cast<QSqlResult*>(ptr)->QSqlResult::fetchPrevious();
}

void* QSqlResult_Handle(void* ptr)
{
	return new QVariant(static_cast<QSqlResult*>(ptr)->handle());
}

void* QSqlResult_HandleDefault(void* ptr)
{
	return new QVariant(static_cast<QSqlResult*>(ptr)->QSqlResult::handle());
}

char QSqlResult_HasOutValues(void* ptr)
{
	return static_cast<QSqlResult*>(ptr)->hasOutValues();
}

char QSqlResult_IsActive(void* ptr)
{
	return static_cast<QSqlResult*>(ptr)->isActive();
}

char QSqlResult_IsForwardOnly(void* ptr)
{
	return static_cast<QSqlResult*>(ptr)->isForwardOnly();
}

char QSqlResult_IsNull(void* ptr, int index)
{
	return static_cast<QSqlResult*>(ptr)->isNull(index);
}

char QSqlResult_IsSelect(void* ptr)
{
	return static_cast<QSqlResult*>(ptr)->isSelect();
}

char QSqlResult_IsValid(void* ptr)
{
	return static_cast<QSqlResult*>(ptr)->isValid();
}

void* QSqlResult_LastError(void* ptr)
{
	return new QSqlError(static_cast<QSqlResult*>(ptr)->lastError());
}

void* QSqlResult_LastInsertId(void* ptr)
{
	return new QVariant(static_cast<QSqlResult*>(ptr)->lastInsertId());
}

void* QSqlResult_LastInsertIdDefault(void* ptr)
{
	return new QVariant(static_cast<QSqlResult*>(ptr)->QSqlResult::lastInsertId());
}

struct QtSql_PackedString QSqlResult_LastQuery(void* ptr)
{
	return ({ QByteArray t9bd7f1 = static_cast<QSqlResult*>(ptr)->lastQuery().toUtf8(); QtSql_PackedString { const_cast<char*>(t9bd7f1.prepend("WHITESPACE").constData()+10), t9bd7f1.size()-10 }; });
}

int QSqlResult_NumRowsAffected(void* ptr)
{
	return static_cast<QSqlResult*>(ptr)->numRowsAffected();
}

char QSqlResult_Prepare(void* ptr, char* query)
{
	return static_cast<QSqlResult*>(ptr)->prepare(QString(query));
}

char QSqlResult_PrepareDefault(void* ptr, char* query)
{
	return static_cast<QSqlResult*>(ptr)->QSqlResult::prepare(QString(query));
}

void* QSqlResult_Record(void* ptr)
{
	return new QSqlRecord(static_cast<QSqlResult*>(ptr)->record());
}

void* QSqlResult_RecordDefault(void* ptr)
{
	return new QSqlRecord(static_cast<QSqlResult*>(ptr)->QSqlResult::record());
}

char QSqlResult_Reset(void* ptr, char* query)
{
	return static_cast<QSqlResult*>(ptr)->reset(QString(query));
}

void QSqlResult_ResetBindCount(void* ptr)
{
	static_cast<QSqlResult*>(ptr)->resetBindCount();
}

char QSqlResult_SavePrepare(void* ptr, char* query)
{
	return static_cast<QSqlResult*>(ptr)->savePrepare(QString(query));
}

char QSqlResult_SavePrepareDefault(void* ptr, char* query)
{
	return static_cast<QSqlResult*>(ptr)->QSqlResult::savePrepare(QString(query));
}

void QSqlResult_SetActive(void* ptr, char active)
{
	static_cast<QSqlResult*>(ptr)->setActive(active != 0);
}

void QSqlResult_SetActiveDefault(void* ptr, char active)
{
	static_cast<QSqlResult*>(ptr)->QSqlResult::setActive(active != 0);
}

void QSqlResult_SetAt(void* ptr, int index)
{
	static_cast<QSqlResult*>(ptr)->setAt(index);
}

void QSqlResult_SetAtDefault(void* ptr, int index)
{
	static_cast<QSqlResult*>(ptr)->QSqlResult::setAt(index);
}

void QSqlResult_SetForwardOnly(void* ptr, char forward)
{
	static_cast<QSqlResult*>(ptr)->setForwardOnly(forward != 0);
}

void QSqlResult_SetForwardOnlyDefault(void* ptr, char forward)
{
	static_cast<QSqlResult*>(ptr)->QSqlResult::setForwardOnly(forward != 0);
}

void QSqlResult_SetLastError(void* ptr, void* error)
{
	static_cast<QSqlResult*>(ptr)->setLastError(*static_cast<QSqlError*>(error));
}

void QSqlResult_SetLastErrorDefault(void* ptr, void* error)
{
	static_cast<QSqlResult*>(ptr)->QSqlResult::setLastError(*static_cast<QSqlError*>(error));
}

void QSqlResult_SetQuery(void* ptr, char* query)
{
	static_cast<QSqlResult*>(ptr)->setQuery(QString(query));
}

void QSqlResult_SetQueryDefault(void* ptr, char* query)
{
	static_cast<QSqlResult*>(ptr)->QSqlResult::setQuery(QString(query));
}

void QSqlResult_SetSelect(void* ptr, char sele)
{
	static_cast<QSqlResult*>(ptr)->setSelect(sele != 0);
}

void QSqlResult_SetSelectDefault(void* ptr, char sele)
{
	static_cast<QSqlResult*>(ptr)->QSqlResult::setSelect(sele != 0);
}

int QSqlResult_Size(void* ptr)
{
	return static_cast<QSqlResult*>(ptr)->size();
}

void QSqlResult_DestroyQSqlResult(void* ptr)
{
	static_cast<QSqlResult*>(ptr)->~QSqlResult();
}

void QSqlResult_DestroyQSqlResultDefault(void* ptr)
{

}

void* QSqlResult___boundValues_atList(void* ptr, int i)
{
	return new QVariant(static_cast<QVector<QVariant>*>(ptr)->at(i));
}

void QSqlResult___boundValues_setList(void* ptr, void* i)
{
	static_cast<QVector<QVariant>*>(ptr)->append(*static_cast<QVariant*>(i));
}

void* QSqlResult___boundValues_newList(void* ptr)
{
	return new QVector<QVariant>;
}

class MyQSqlTableModel: public QSqlTableModel
{
public:
	MyQSqlTableModel(QObject *parent, QSqlDatabase db) : QSqlTableModel(parent, db) {};
	void Signal_BeforeDelete(int row) { callbackQSqlTableModel_BeforeDelete(this, row); };
	void Signal_BeforeInsert(QSqlRecord & record) { callbackQSqlTableModel_BeforeInsert(this, static_cast<QSqlRecord*>(&record)); };
	void Signal_BeforeUpdate(int row, QSqlRecord & record) { callbackQSqlTableModel_BeforeUpdate(this, row, static_cast<QSqlRecord*>(&record)); };
	void clear() { callbackQSqlTableModel_Clear(this); };
	QVariant data(const QModelIndex & index, int role) const { return *static_cast<QVariant*>(callbackQSqlTableModel_Data(const_cast<MyQSqlTableModel*>(this), const_cast<QModelIndex*>(&index), role)); };
	bool deleteRowFromTable(int row) { return callbackQSqlTableModel_DeleteRowFromTable(this, row) != 0; };
	Qt::ItemFlags flags(const QModelIndex & index) const { return static_cast<Qt::ItemFlag>(callbackQSqlTableModel_Flags(const_cast<MyQSqlTableModel*>(this), const_cast<QModelIndex*>(&index))); };
	QVariant headerData(int section, Qt::Orientation orientation, int role) const { return *static_cast<QVariant*>(callbackQSqlTableModel_HeaderData(const_cast<MyQSqlTableModel*>(this), section, orientation, role)); };
	QModelIndex indexInQuery(const QModelIndex & item) const { return *static_cast<QModelIndex*>(callbackQSqlTableModel_IndexInQuery(const_cast<MyQSqlTableModel*>(this), const_cast<QModelIndex*>(&item))); };
	bool insertRowIntoTable(const QSqlRecord & values) { return callbackQSqlTableModel_InsertRowIntoTable(this, const_cast<QSqlRecord*>(&values)) != 0; };
	bool insertRows(int row, int count, const QModelIndex & parent) { return callbackQSqlTableModel_InsertRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	QString orderByClause() const { return QString(callbackQSqlTableModel_OrderByClause(const_cast<MyQSqlTableModel*>(this))); };
	void Signal_PrimeInsert(int row, QSqlRecord & record) { callbackQSqlTableModel_PrimeInsert(this, row, static_cast<QSqlRecord*>(&record)); };
	bool removeColumns(int column, int count, const QModelIndex & parent) { return callbackQSqlTableModel_RemoveColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool removeRows(int row, int count, const QModelIndex & parent) { return callbackQSqlTableModel_RemoveRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	void revert() { callbackQSqlTableModel_Revert(this); };
	void revertAll() { callbackQSqlTableModel_RevertAll(this); };
	void revertRow(int row) { callbackQSqlTableModel_RevertRow(this, row); };
	int rowCount(const QModelIndex & parent) const { return callbackQSqlTableModel_RowCount(const_cast<MyQSqlTableModel*>(this), const_cast<QModelIndex*>(&parent)); };
	bool select() { return callbackQSqlTableModel_Select(this) != 0; };
	bool selectRow(int row) { return callbackQSqlTableModel_SelectRow(this, row) != 0; };
	QString selectStatement() const { return QString(callbackQSqlTableModel_SelectStatement(const_cast<MyQSqlTableModel*>(this))); };
	bool setData(const QModelIndex & index, const QVariant & value, int role) { return callbackQSqlTableModel_SetData(this, const_cast<QModelIndex*>(&index), const_cast<QVariant*>(&value), role) != 0; };
	void setEditStrategy(QSqlTableModel::EditStrategy strategy) { callbackQSqlTableModel_SetEditStrategy(this, strategy); };
	void setFilter(const QString & filter) { QByteArray t4bb4ca = filter.toUtf8(); QtSql_PackedString filterPacked = { const_cast<char*>(t4bb4ca.prepend("WHITESPACE").constData()+10), t4bb4ca.size()-10 };callbackQSqlTableModel_SetFilter(this, filterPacked); };
	void setSort(int column, Qt::SortOrder order) { callbackQSqlTableModel_SetSort(this, column, order); };
	void setTable(const QString & tableName) { QByteArray t3e7060 = tableName.toUtf8(); QtSql_PackedString tableNamePacked = { const_cast<char*>(t3e7060.prepend("WHITESPACE").constData()+10), t3e7060.size()-10 };callbackQSqlTableModel_SetTable(this, tableNamePacked); };
	void sort(int column, Qt::SortOrder order) { callbackQSqlTableModel_Sort(this, column, order); };
	bool submit() { return callbackQSqlTableModel_Submit(this) != 0; };
	bool submitAll() { return callbackQSqlTableModel_SubmitAll(this) != 0; };
	bool updateRowInTable(int row, const QSqlRecord & values) { return callbackQSqlTableModel_UpdateRowInTable(this, row, const_cast<QSqlRecord*>(&values)) != 0; };
	 ~MyQSqlTableModel() { callbackQSqlTableModel_DestroyQSqlTableModel(this); };
	bool canFetchMore(const QModelIndex & parent) const { return callbackQSqlTableModel_CanFetchMore(const_cast<MyQSqlTableModel*>(this), const_cast<QModelIndex*>(&parent)) != 0; };
	int columnCount(const QModelIndex & index) const { return callbackQSqlTableModel_ColumnCount(const_cast<MyQSqlTableModel*>(this), const_cast<QModelIndex*>(&index)); };
	void fetchMore(const QModelIndex & parent) { callbackQSqlTableModel_FetchMore(this, const_cast<QModelIndex*>(&parent)); };
	bool insertColumns(int column, int count, const QModelIndex & parent) { return callbackQSqlTableModel_InsertColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	void queryChange() { callbackQSqlTableModel_QueryChange(this); };
	bool setHeaderData(int section, Qt::Orientation orientation, const QVariant & value, int role) { return callbackQSqlTableModel_SetHeaderData(this, section, orientation, const_cast<QVariant*>(&value), role) != 0; };
	QModelIndex index(int row, int column, const QModelIndex & parent) const { return *static_cast<QModelIndex*>(callbackQSqlTableModel_Index(const_cast<MyQSqlTableModel*>(this), row, column, const_cast<QModelIndex*>(&parent))); };
	bool dropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) { return callbackQSqlTableModel_DropMimeData(this, const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	QModelIndex sibling(int row, int column, const QModelIndex & idx) const { return *static_cast<QModelIndex*>(callbackQSqlTableModel_Sibling(const_cast<MyQSqlTableModel*>(this), row, column, const_cast<QModelIndex*>(&idx))); };
	QModelIndex buddy(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackQSqlTableModel_Buddy(const_cast<MyQSqlTableModel*>(this), const_cast<QModelIndex*>(&index))); };
	bool canDropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) const { return callbackQSqlTableModel_CanDropMimeData(const_cast<MyQSqlTableModel*>(this), const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	bool hasChildren(const QModelIndex & parent) const { return callbackQSqlTableModel_HasChildren(const_cast<MyQSqlTableModel*>(this), const_cast<QModelIndex*>(&parent)) != 0; };
	QMap<int, QVariant> itemData(const QModelIndex & index) const { return *static_cast<QMap<int, QVariant>*>(callbackQSqlTableModel_ItemData(const_cast<MyQSqlTableModel*>(this), const_cast<QModelIndex*>(&index))); };
	QList<QModelIndex> match(const QModelIndex & start, int role, const QVariant & value, int hits, Qt::MatchFlags flags) const { return *static_cast<QList<QModelIndex>*>(callbackQSqlTableModel_Match(const_cast<MyQSqlTableModel*>(this), const_cast<QModelIndex*>(&start), role, const_cast<QVariant*>(&value), hits, flags)); };
	QMimeData * mimeData(const QModelIndexList & indexes) const { return static_cast<QMimeData*>(callbackQSqlTableModel_MimeData(const_cast<MyQSqlTableModel*>(this), ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(indexes); QtSql_PackedList { tmpValue, tmpValue->size() }; }))); };
	QStringList mimeTypes() const { return QString(callbackQSqlTableModel_MimeTypes(const_cast<MyQSqlTableModel*>(this))).split("|", QString::SkipEmptyParts); };
	bool moveColumns(const QModelIndex & sourceParent, int sourceColumn, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackQSqlTableModel_MoveColumns(this, const_cast<QModelIndex*>(&sourceParent), sourceColumn, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	bool moveRows(const QModelIndex & sourceParent, int sourceRow, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackQSqlTableModel_MoveRows(this, const_cast<QModelIndex*>(&sourceParent), sourceRow, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	QModelIndex parent(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackQSqlTableModel_Parent(const_cast<MyQSqlTableModel*>(this), const_cast<QModelIndex*>(&index))); };
	void resetInternalData() { callbackQSqlTableModel_ResetInternalData(this); };
	QHash<int, QByteArray> roleNames() const { return *static_cast<QHash<int, QByteArray>*>(callbackQSqlTableModel_RoleNames(const_cast<MyQSqlTableModel*>(this))); };
	bool setItemData(const QModelIndex & index, const QMap<int, QVariant> & roles) { return callbackQSqlTableModel_SetItemData(this, const_cast<QModelIndex*>(&index), ({ QMap<int, QVariant>* tmpValue = const_cast<QMap<int, QVariant>*>(&roles); QtSql_PackedList { tmpValue, tmpValue->size() }; })) != 0; };
	QSize span(const QModelIndex & index) const { return *static_cast<QSize*>(callbackQSqlTableModel_Span(const_cast<MyQSqlTableModel*>(this), const_cast<QModelIndex*>(&index))); };
	Qt::DropActions supportedDragActions() const { return static_cast<Qt::DropAction>(callbackQSqlTableModel_SupportedDragActions(const_cast<MyQSqlTableModel*>(this))); };
	Qt::DropActions supportedDropActions() const { return static_cast<Qt::DropAction>(callbackQSqlTableModel_SupportedDropActions(const_cast<MyQSqlTableModel*>(this))); };
	void timerEvent(QTimerEvent * event) { callbackQSqlTableModel_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQSqlTableModel_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSqlTableModel_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQSqlTableModel_CustomEvent(this, event); };
	void deleteLater() { callbackQSqlTableModel_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSqlTableModel_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQSqlTableModel_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSqlTableModel_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSqlTableModel_MetaObject(const_cast<MyQSqlTableModel*>(this))); };
};

void* QSqlTableModel_NewQSqlTableModel(void* parent, void* db)
{
	return new MyQSqlTableModel(static_cast<QObject*>(parent), *static_cast<QSqlDatabase*>(db));
}

void QSqlTableModel_ConnectBeforeDelete(void* ptr)
{
	QObject::connect(static_cast<QSqlTableModel*>(ptr), static_cast<void (QSqlTableModel::*)(int)>(&QSqlTableModel::beforeDelete), static_cast<MyQSqlTableModel*>(ptr), static_cast<void (MyQSqlTableModel::*)(int)>(&MyQSqlTableModel::Signal_BeforeDelete));
}

void QSqlTableModel_DisconnectBeforeDelete(void* ptr)
{
	QObject::disconnect(static_cast<QSqlTableModel*>(ptr), static_cast<void (QSqlTableModel::*)(int)>(&QSqlTableModel::beforeDelete), static_cast<MyQSqlTableModel*>(ptr), static_cast<void (MyQSqlTableModel::*)(int)>(&MyQSqlTableModel::Signal_BeforeDelete));
}

void QSqlTableModel_BeforeDelete(void* ptr, int row)
{
	static_cast<QSqlTableModel*>(ptr)->beforeDelete(row);
}

void QSqlTableModel_ConnectBeforeInsert(void* ptr)
{
	QObject::connect(static_cast<QSqlTableModel*>(ptr), static_cast<void (QSqlTableModel::*)(QSqlRecord &)>(&QSqlTableModel::beforeInsert), static_cast<MyQSqlTableModel*>(ptr), static_cast<void (MyQSqlTableModel::*)(QSqlRecord &)>(&MyQSqlTableModel::Signal_BeforeInsert));
}

void QSqlTableModel_DisconnectBeforeInsert(void* ptr)
{
	QObject::disconnect(static_cast<QSqlTableModel*>(ptr), static_cast<void (QSqlTableModel::*)(QSqlRecord &)>(&QSqlTableModel::beforeInsert), static_cast<MyQSqlTableModel*>(ptr), static_cast<void (MyQSqlTableModel::*)(QSqlRecord &)>(&MyQSqlTableModel::Signal_BeforeInsert));
}

void QSqlTableModel_BeforeInsert(void* ptr, void* record)
{
	static_cast<QSqlTableModel*>(ptr)->beforeInsert(*static_cast<QSqlRecord*>(record));
}

void QSqlTableModel_ConnectBeforeUpdate(void* ptr)
{
	QObject::connect(static_cast<QSqlTableModel*>(ptr), static_cast<void (QSqlTableModel::*)(int, QSqlRecord &)>(&QSqlTableModel::beforeUpdate), static_cast<MyQSqlTableModel*>(ptr), static_cast<void (MyQSqlTableModel::*)(int, QSqlRecord &)>(&MyQSqlTableModel::Signal_BeforeUpdate));
}

void QSqlTableModel_DisconnectBeforeUpdate(void* ptr)
{
	QObject::disconnect(static_cast<QSqlTableModel*>(ptr), static_cast<void (QSqlTableModel::*)(int, QSqlRecord &)>(&QSqlTableModel::beforeUpdate), static_cast<MyQSqlTableModel*>(ptr), static_cast<void (MyQSqlTableModel::*)(int, QSqlRecord &)>(&MyQSqlTableModel::Signal_BeforeUpdate));
}

void QSqlTableModel_BeforeUpdate(void* ptr, int row, void* record)
{
	static_cast<QSqlTableModel*>(ptr)->beforeUpdate(row, *static_cast<QSqlRecord*>(record));
}

void QSqlTableModel_Clear(void* ptr)
{
	static_cast<QSqlTableModel*>(ptr)->clear();
}

void QSqlTableModel_ClearDefault(void* ptr)
{
	static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::clear();
}

void* QSqlTableModel_Data(void* ptr, void* index, int role)
{
	return new QVariant(static_cast<QSqlTableModel*>(ptr)->data(*static_cast<QModelIndex*>(index), role));
}

void* QSqlTableModel_DataDefault(void* ptr, void* index, int role)
{
	return new QVariant(static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::data(*static_cast<QModelIndex*>(index), role));
}

void* QSqlTableModel_Database(void* ptr)
{
	return new QSqlDatabase(static_cast<QSqlTableModel*>(ptr)->database());
}

char QSqlTableModel_DeleteRowFromTable(void* ptr, int row)
{
	return static_cast<QSqlTableModel*>(ptr)->deleteRowFromTable(row);
}

char QSqlTableModel_DeleteRowFromTableDefault(void* ptr, int row)
{
	return static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::deleteRowFromTable(row);
}

long long QSqlTableModel_EditStrategy(void* ptr)
{
	return static_cast<QSqlTableModel*>(ptr)->editStrategy();
}

int QSqlTableModel_FieldIndex(void* ptr, char* fieldName)
{
	return static_cast<QSqlTableModel*>(ptr)->fieldIndex(QString(fieldName));
}

struct QtSql_PackedString QSqlTableModel_Filter(void* ptr)
{
	return ({ QByteArray td0370f = static_cast<QSqlTableModel*>(ptr)->filter().toUtf8(); QtSql_PackedString { const_cast<char*>(td0370f.prepend("WHITESPACE").constData()+10), td0370f.size()-10 }; });
}

long long QSqlTableModel_Flags(void* ptr, void* index)
{
	return static_cast<QSqlTableModel*>(ptr)->flags(*static_cast<QModelIndex*>(index));
}

long long QSqlTableModel_FlagsDefault(void* ptr, void* index)
{
	return static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::flags(*static_cast<QModelIndex*>(index));
}

void* QSqlTableModel_HeaderData(void* ptr, int section, long long orientation, int role)
{
	return new QVariant(static_cast<QSqlTableModel*>(ptr)->headerData(section, static_cast<Qt::Orientation>(orientation), role));
}

void* QSqlTableModel_HeaderDataDefault(void* ptr, int section, long long orientation, int role)
{
	return new QVariant(static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::headerData(section, static_cast<Qt::Orientation>(orientation), role));
}

void* QSqlTableModel_IndexInQuery(void* ptr, void* item)
{
	return new QModelIndex(static_cast<QSqlTableModel*>(ptr)->indexInQuery(*static_cast<QModelIndex*>(item)));
}

void* QSqlTableModel_IndexInQueryDefault(void* ptr, void* item)
{
	return new QModelIndex(static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::indexInQuery(*static_cast<QModelIndex*>(item)));
}

char QSqlTableModel_InsertRecord(void* ptr, int row, void* record)
{
	return static_cast<QSqlTableModel*>(ptr)->insertRecord(row, *static_cast<QSqlRecord*>(record));
}

char QSqlTableModel_InsertRowIntoTable(void* ptr, void* values)
{
	return static_cast<QSqlTableModel*>(ptr)->insertRowIntoTable(*static_cast<QSqlRecord*>(values));
}

char QSqlTableModel_InsertRowIntoTableDefault(void* ptr, void* values)
{
	return static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::insertRowIntoTable(*static_cast<QSqlRecord*>(values));
}

char QSqlTableModel_InsertRows(void* ptr, int row, int count, void* parent)
{
	return static_cast<QSqlTableModel*>(ptr)->insertRows(row, count, *static_cast<QModelIndex*>(parent));
}

char QSqlTableModel_InsertRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::insertRows(row, count, *static_cast<QModelIndex*>(parent));
}

char QSqlTableModel_IsDirty2(void* ptr)
{
	return static_cast<QSqlTableModel*>(ptr)->isDirty();
}

char QSqlTableModel_IsDirty(void* ptr, void* index)
{
	return static_cast<QSqlTableModel*>(ptr)->isDirty(*static_cast<QModelIndex*>(index));
}

struct QtSql_PackedString QSqlTableModel_OrderByClause(void* ptr)
{
	return ({ QByteArray tca221a = static_cast<QSqlTableModel*>(ptr)->orderByClause().toUtf8(); QtSql_PackedString { const_cast<char*>(tca221a.prepend("WHITESPACE").constData()+10), tca221a.size()-10 }; });
}

struct QtSql_PackedString QSqlTableModel_OrderByClauseDefault(void* ptr)
{
	return ({ QByteArray t5a4e03 = static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::orderByClause().toUtf8(); QtSql_PackedString { const_cast<char*>(t5a4e03.prepend("WHITESPACE").constData()+10), t5a4e03.size()-10 }; });
}

void* QSqlTableModel_PrimaryKey(void* ptr)
{
	return new QSqlIndex(static_cast<QSqlTableModel*>(ptr)->primaryKey());
}

void* QSqlTableModel_PrimaryValues(void* ptr, int row)
{
	return new QSqlRecord(static_cast<QSqlTableModel*>(ptr)->primaryValues(row));
}

void QSqlTableModel_ConnectPrimeInsert(void* ptr)
{
	QObject::connect(static_cast<QSqlTableModel*>(ptr), static_cast<void (QSqlTableModel::*)(int, QSqlRecord &)>(&QSqlTableModel::primeInsert), static_cast<MyQSqlTableModel*>(ptr), static_cast<void (MyQSqlTableModel::*)(int, QSqlRecord &)>(&MyQSqlTableModel::Signal_PrimeInsert));
}

void QSqlTableModel_DisconnectPrimeInsert(void* ptr)
{
	QObject::disconnect(static_cast<QSqlTableModel*>(ptr), static_cast<void (QSqlTableModel::*)(int, QSqlRecord &)>(&QSqlTableModel::primeInsert), static_cast<MyQSqlTableModel*>(ptr), static_cast<void (MyQSqlTableModel::*)(int, QSqlRecord &)>(&MyQSqlTableModel::Signal_PrimeInsert));
}

void QSqlTableModel_PrimeInsert(void* ptr, int row, void* record)
{
	static_cast<QSqlTableModel*>(ptr)->primeInsert(row, *static_cast<QSqlRecord*>(record));
}

void* QSqlTableModel_Record(void* ptr)
{
	return new QSqlRecord(static_cast<QSqlTableModel*>(ptr)->record());
}

void* QSqlTableModel_Record2(void* ptr, int row)
{
	return new QSqlRecord(static_cast<QSqlTableModel*>(ptr)->record(row));
}

char QSqlTableModel_RemoveColumns(void* ptr, int column, int count, void* parent)
{
	return static_cast<QSqlTableModel*>(ptr)->removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char QSqlTableModel_RemoveColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char QSqlTableModel_RemoveRows(void* ptr, int row, int count, void* parent)
{
	return static_cast<QSqlTableModel*>(ptr)->removeRows(row, count, *static_cast<QModelIndex*>(parent));
}

char QSqlTableModel_RemoveRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::removeRows(row, count, *static_cast<QModelIndex*>(parent));
}

void QSqlTableModel_Revert(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QSqlTableModel*>(ptr), "revert");
}

void QSqlTableModel_RevertDefault(void* ptr)
{
	static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::revert();
}

void QSqlTableModel_RevertAll(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QSqlTableModel*>(ptr), "revertAll");
}

void QSqlTableModel_RevertAllDefault(void* ptr)
{
	static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::revertAll();
}

void QSqlTableModel_RevertRow(void* ptr, int row)
{
	static_cast<QSqlTableModel*>(ptr)->revertRow(row);
}

void QSqlTableModel_RevertRowDefault(void* ptr, int row)
{
	static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::revertRow(row);
}

int QSqlTableModel_RowCount(void* ptr, void* parent)
{
	return static_cast<QSqlTableModel*>(ptr)->rowCount(*static_cast<QModelIndex*>(parent));
}

int QSqlTableModel_RowCountDefault(void* ptr, void* parent)
{
	return static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::rowCount(*static_cast<QModelIndex*>(parent));
}

char QSqlTableModel_Select(void* ptr)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QSqlTableModel*>(ptr), "select", Q_RETURN_ARG(bool, returnArg));
	return returnArg;
}

char QSqlTableModel_SelectDefault(void* ptr)
{
	return static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::select();
}

char QSqlTableModel_SelectRow(void* ptr, int row)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QSqlTableModel*>(ptr), "selectRow", Q_RETURN_ARG(bool, returnArg), Q_ARG(int, row));
	return returnArg;
}

char QSqlTableModel_SelectRowDefault(void* ptr, int row)
{
	return static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::selectRow(row);
}

struct QtSql_PackedString QSqlTableModel_SelectStatement(void* ptr)
{
	return ({ QByteArray tf1b779 = static_cast<QSqlTableModel*>(ptr)->selectStatement().toUtf8(); QtSql_PackedString { const_cast<char*>(tf1b779.prepend("WHITESPACE").constData()+10), tf1b779.size()-10 }; });
}

struct QtSql_PackedString QSqlTableModel_SelectStatementDefault(void* ptr)
{
	return ({ QByteArray t5344b7 = static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::selectStatement().toUtf8(); QtSql_PackedString { const_cast<char*>(t5344b7.prepend("WHITESPACE").constData()+10), t5344b7.size()-10 }; });
}

char QSqlTableModel_SetData(void* ptr, void* index, void* value, int role)
{
	return static_cast<QSqlTableModel*>(ptr)->setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
}

char QSqlTableModel_SetDataDefault(void* ptr, void* index, void* value, int role)
{
	return static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
}

void QSqlTableModel_SetEditStrategy(void* ptr, long long strategy)
{
	static_cast<QSqlTableModel*>(ptr)->setEditStrategy(static_cast<QSqlTableModel::EditStrategy>(strategy));
}

void QSqlTableModel_SetEditStrategyDefault(void* ptr, long long strategy)
{
	static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::setEditStrategy(static_cast<QSqlTableModel::EditStrategy>(strategy));
}

void QSqlTableModel_SetFilter(void* ptr, char* filter)
{
	static_cast<QSqlTableModel*>(ptr)->setFilter(QString(filter));
}

void QSqlTableModel_SetFilterDefault(void* ptr, char* filter)
{
	static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::setFilter(QString(filter));
}

void QSqlTableModel_SetPrimaryKey(void* ptr, void* key)
{
	static_cast<QSqlTableModel*>(ptr)->setPrimaryKey(*static_cast<QSqlIndex*>(key));
}

void QSqlTableModel_SetQuery(void* ptr, void* query)
{
	static_cast<QSqlTableModel*>(ptr)->setQuery(*static_cast<QSqlQuery*>(query));
}

char QSqlTableModel_SetRecord(void* ptr, int row, void* values)
{
	return static_cast<QSqlTableModel*>(ptr)->setRecord(row, *static_cast<QSqlRecord*>(values));
}

void QSqlTableModel_SetSort(void* ptr, int column, long long order)
{
	static_cast<QSqlTableModel*>(ptr)->setSort(column, static_cast<Qt::SortOrder>(order));
}

void QSqlTableModel_SetSortDefault(void* ptr, int column, long long order)
{
	static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::setSort(column, static_cast<Qt::SortOrder>(order));
}

void QSqlTableModel_SetTable(void* ptr, char* tableName)
{
	static_cast<QSqlTableModel*>(ptr)->setTable(QString(tableName));
}

void QSqlTableModel_SetTableDefault(void* ptr, char* tableName)
{
	static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::setTable(QString(tableName));
}

void QSqlTableModel_Sort(void* ptr, int column, long long order)
{
	static_cast<QSqlTableModel*>(ptr)->sort(column, static_cast<Qt::SortOrder>(order));
}

void QSqlTableModel_SortDefault(void* ptr, int column, long long order)
{
	static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::sort(column, static_cast<Qt::SortOrder>(order));
}

char QSqlTableModel_Submit(void* ptr)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QSqlTableModel*>(ptr), "submit", Q_RETURN_ARG(bool, returnArg));
	return returnArg;
}

char QSqlTableModel_SubmitDefault(void* ptr)
{
	return static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::submit();
}

char QSqlTableModel_SubmitAll(void* ptr)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QSqlTableModel*>(ptr), "submitAll", Q_RETURN_ARG(bool, returnArg));
	return returnArg;
}

char QSqlTableModel_SubmitAllDefault(void* ptr)
{
	return static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::submitAll();
}

struct QtSql_PackedString QSqlTableModel_TableName(void* ptr)
{
	return ({ QByteArray tefb63d = static_cast<QSqlTableModel*>(ptr)->tableName().toUtf8(); QtSql_PackedString { const_cast<char*>(tefb63d.prepend("WHITESPACE").constData()+10), tefb63d.size()-10 }; });
}

char QSqlTableModel_UpdateRowInTable(void* ptr, int row, void* values)
{
	return static_cast<QSqlTableModel*>(ptr)->updateRowInTable(row, *static_cast<QSqlRecord*>(values));
}

char QSqlTableModel_UpdateRowInTableDefault(void* ptr, int row, void* values)
{
	return static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::updateRowInTable(row, *static_cast<QSqlRecord*>(values));
}

void QSqlTableModel_DestroyQSqlTableModel(void* ptr)
{
	static_cast<QSqlTableModel*>(ptr)->~QSqlTableModel();
}

void QSqlTableModel_DestroyQSqlTableModelDefault(void* ptr)
{

}

int QSqlTableModel_____itemData_keyList_atList(void* ptr, int i)
{
	return static_cast<QList<int>*>(ptr)->at(i);
}

void QSqlTableModel_____itemData_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* QSqlTableModel_____itemData_keyList_newList(void* ptr)
{
	return new QList<int>;
}

int QSqlTableModel_____roleNames_keyList_atList(void* ptr, int i)
{
	return static_cast<QList<int>*>(ptr)->at(i);
}

void QSqlTableModel_____roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* QSqlTableModel_____roleNames_keyList_newList(void* ptr)
{
	return new QList<int>;
}

int QSqlTableModel_____setItemData_keyList_atList(void* ptr, int i)
{
	return static_cast<QList<int>*>(ptr)->at(i);
}

void QSqlTableModel_____setItemData_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* QSqlTableModel_____setItemData_keyList_newList(void* ptr)
{
	return new QList<int>;
}

void* QSqlTableModel___changePersistentIndexList_from_atList(void* ptr, int i)
{
	return new QModelIndex(static_cast<QList<QModelIndex>*>(ptr)->at(i));
}

void QSqlTableModel___changePersistentIndexList_from_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* QSqlTableModel___changePersistentIndexList_from_newList(void* ptr)
{
	return new QList<QModelIndex>;
}

void* QSqlTableModel___changePersistentIndexList_to_atList(void* ptr, int i)
{
	return new QModelIndex(static_cast<QList<QModelIndex>*>(ptr)->at(i));
}

void QSqlTableModel___changePersistentIndexList_to_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* QSqlTableModel___changePersistentIndexList_to_newList(void* ptr)
{
	return new QList<QModelIndex>;
}

int QSqlTableModel___dataChanged_roles_atList(void* ptr, int i)
{
	return static_cast<QVector<int>*>(ptr)->at(i);
}

void QSqlTableModel___dataChanged_roles_setList(void* ptr, int i)
{
	static_cast<QVector<int>*>(ptr)->append(i);
}

void* QSqlTableModel___dataChanged_roles_newList(void* ptr)
{
	return new QVector<int>;
}

void* QSqlTableModel___itemData_atList(void* ptr, int i)
{
	return new QVariant(static_cast<QMap<int, QVariant>*>(ptr)->value(i));
}

void QSqlTableModel___itemData_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* QSqlTableModel___itemData_newList(void* ptr)
{
	return new QMap<int, QVariant>;
}

struct QtSql_PackedList QSqlTableModel___itemData_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); QtSql_PackedList { tmpValue, tmpValue->size() }; });
}

void* QSqlTableModel___layoutAboutToBeChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i));
}

void QSqlTableModel___layoutAboutToBeChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* QSqlTableModel___layoutAboutToBeChanged_parents_newList(void* ptr)
{
	return new QList<QPersistentModelIndex>;
}

void* QSqlTableModel___layoutChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i));
}

void QSqlTableModel___layoutChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* QSqlTableModel___layoutChanged_parents_newList(void* ptr)
{
	return new QList<QPersistentModelIndex>;
}

void* QSqlTableModel___match_atList(void* ptr, int i)
{
	return new QModelIndex(static_cast<QList<QModelIndex>*>(ptr)->at(i));
}

void QSqlTableModel___match_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* QSqlTableModel___match_newList(void* ptr)
{
	return new QList<QModelIndex>;
}

void* QSqlTableModel___mimeData_indexes_atList(void* ptr, int i)
{
	return new QModelIndex(static_cast<QList<QModelIndex>*>(ptr)->at(i));
}

void QSqlTableModel___mimeData_indexes_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* QSqlTableModel___mimeData_indexes_newList(void* ptr)
{
	return new QList<QModelIndex>;
}

void* QSqlTableModel___persistentIndexList_atList(void* ptr, int i)
{
	return new QModelIndex(static_cast<QList<QModelIndex>*>(ptr)->at(i));
}

void QSqlTableModel___persistentIndexList_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* QSqlTableModel___persistentIndexList_newList(void* ptr)
{
	return new QList<QModelIndex>;
}

void* QSqlTableModel___roleNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QHash<int, QByteArray>*>(ptr)->value(i));
}

void QSqlTableModel___roleNames_setList(void* ptr, int key, void* i)
{
	static_cast<QHash<int, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* QSqlTableModel___roleNames_newList(void* ptr)
{
	return new QHash<int, QByteArray>;
}

struct QtSql_PackedList QSqlTableModel___roleNames_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QHash<int, QByteArray>*>(ptr)->keys()); QtSql_PackedList { tmpValue, tmpValue->size() }; });
}

void* QSqlTableModel___setItemData_roles_atList(void* ptr, int i)
{
	return new QVariant(static_cast<QMap<int, QVariant>*>(ptr)->value(i));
}

void QSqlTableModel___setItemData_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* QSqlTableModel___setItemData_roles_newList(void* ptr)
{
	return new QMap<int, QVariant>;
}

struct QtSql_PackedList QSqlTableModel___setItemData_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); QtSql_PackedList { tmpValue, tmpValue->size() }; });
}

int QSqlTableModel_____doSetRoleNames_keyList_atList(void* ptr, int i)
{
	return static_cast<QList<int>*>(ptr)->at(i);
}

void QSqlTableModel_____doSetRoleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* QSqlTableModel_____doSetRoleNames_keyList_newList(void* ptr)
{
	return new QList<int>;
}

int QSqlTableModel_____setRoleNames_keyList_atList(void* ptr, int i)
{
	return static_cast<QList<int>*>(ptr)->at(i);
}

void QSqlTableModel_____setRoleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* QSqlTableModel_____setRoleNames_keyList_newList(void* ptr)
{
	return new QList<int>;
}

void* QSqlTableModel___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QSqlTableModel___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSqlTableModel___children_newList(void* ptr)
{
	return new QList<QObject *>;
}

void* QSqlTableModel___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QSqlTableModel___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QSqlTableModel___dynamicPropertyNames_newList(void* ptr)
{
	return new QList<QByteArray>;
}

void* QSqlTableModel___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QSqlTableModel___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSqlTableModel___findChildren_newList2(void* ptr)
{
	return new QList<QObject*>;
}

void* QSqlTableModel___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QSqlTableModel___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSqlTableModel___findChildren_newList3(void* ptr)
{
	return new QList<QObject*>;
}

void* QSqlTableModel___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QSqlTableModel___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSqlTableModel___findChildren_newList(void* ptr)
{
	return new QList<QObject*>;
}

char QSqlTableModel_CanFetchMore(void* ptr, void* parent)
{
	return static_cast<QSqlTableModel*>(ptr)->canFetchMore(*static_cast<QModelIndex*>(parent));
}

char QSqlTableModel_CanFetchMoreDefault(void* ptr, void* parent)
{
	return static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::canFetchMore(*static_cast<QModelIndex*>(parent));
}

int QSqlTableModel_ColumnCount(void* ptr, void* index)
{
	return static_cast<QSqlTableModel*>(ptr)->columnCount(*static_cast<QModelIndex*>(index));
}

int QSqlTableModel_ColumnCountDefault(void* ptr, void* index)
{
	return static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::columnCount(*static_cast<QModelIndex*>(index));
}

void QSqlTableModel_FetchMore(void* ptr, void* parent)
{
	static_cast<QSqlTableModel*>(ptr)->fetchMore(*static_cast<QModelIndex*>(parent));
}

void QSqlTableModel_FetchMoreDefault(void* ptr, void* parent)
{
	static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::fetchMore(*static_cast<QModelIndex*>(parent));
}

char QSqlTableModel_InsertColumns(void* ptr, int column, int count, void* parent)
{
	return static_cast<QSqlTableModel*>(ptr)->insertColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char QSqlTableModel_InsertColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::insertColumns(column, count, *static_cast<QModelIndex*>(parent));
}

void QSqlTableModel_QueryChange(void* ptr)
{
	static_cast<QSqlTableModel*>(ptr)->queryChange();
}

void QSqlTableModel_QueryChangeDefault(void* ptr)
{
	static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::queryChange();
}

char QSqlTableModel_SetHeaderData(void* ptr, int section, long long orientation, void* value, int role)
{
	return static_cast<QSqlTableModel*>(ptr)->setHeaderData(section, static_cast<Qt::Orientation>(orientation), *static_cast<QVariant*>(value), role);
}

char QSqlTableModel_SetHeaderDataDefault(void* ptr, int section, long long orientation, void* value, int role)
{
	return static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::setHeaderData(section, static_cast<Qt::Orientation>(orientation), *static_cast<QVariant*>(value), role);
}

void* QSqlTableModel_Index(void* ptr, int row, int column, void* parent)
{
	return new QModelIndex(static_cast<QSqlTableModel*>(ptr)->index(row, column, *static_cast<QModelIndex*>(parent)));
}

void* QSqlTableModel_IndexDefault(void* ptr, int row, int column, void* parent)
{
	return new QModelIndex(static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::index(row, column, *static_cast<QModelIndex*>(parent)));
}

char QSqlTableModel_DropMimeData(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<QSqlTableModel*>(ptr)->dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

char QSqlTableModel_DropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

void* QSqlTableModel_Sibling(void* ptr, int row, int column, void* idx)
{
	return new QModelIndex(static_cast<QSqlTableModel*>(ptr)->sibling(row, column, *static_cast<QModelIndex*>(idx)));
}

void* QSqlTableModel_SiblingDefault(void* ptr, int row, int column, void* idx)
{
	return new QModelIndex(static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::sibling(row, column, *static_cast<QModelIndex*>(idx)));
}

void* QSqlTableModel_Buddy(void* ptr, void* index)
{
	return new QModelIndex(static_cast<QSqlTableModel*>(ptr)->buddy(*static_cast<QModelIndex*>(index)));
}

void* QSqlTableModel_BuddyDefault(void* ptr, void* index)
{
	return new QModelIndex(static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::buddy(*static_cast<QModelIndex*>(index)));
}

char QSqlTableModel_CanDropMimeData(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<QSqlTableModel*>(ptr)->canDropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

char QSqlTableModel_CanDropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::canDropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

char QSqlTableModel_HasChildren(void* ptr, void* parent)
{
	return static_cast<QSqlTableModel*>(ptr)->hasChildren(*static_cast<QModelIndex*>(parent));
}

char QSqlTableModel_HasChildrenDefault(void* ptr, void* parent)
{
	return static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::hasChildren(*static_cast<QModelIndex*>(parent));
}

struct QtSql_PackedList QSqlTableModel_ItemData(void* ptr, void* index)
{
	return ({ QMap<int, QVariant>* tmpValue = new QMap<int, QVariant>(static_cast<QSqlTableModel*>(ptr)->itemData(*static_cast<QModelIndex*>(index))); QtSql_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtSql_PackedList QSqlTableModel_ItemDataDefault(void* ptr, void* index)
{
	return ({ QMap<int, QVariant>* tmpValue = new QMap<int, QVariant>(static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::itemData(*static_cast<QModelIndex*>(index))); QtSql_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtSql_PackedList QSqlTableModel_Match(void* ptr, void* start, int role, void* value, int hits, long long flags)
{
	return ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(static_cast<QSqlTableModel*>(ptr)->match(*static_cast<QModelIndex*>(start), role, *static_cast<QVariant*>(value), hits, static_cast<Qt::MatchFlag>(flags))); QtSql_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtSql_PackedList QSqlTableModel_MatchDefault(void* ptr, void* start, int role, void* value, int hits, long long flags)
{
	return ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::match(*static_cast<QModelIndex*>(start), role, *static_cast<QVariant*>(value), hits, static_cast<Qt::MatchFlag>(flags))); QtSql_PackedList { tmpValue, tmpValue->size() }; });
}

void* QSqlTableModel_MimeData(void* ptr, void* indexes)
{
	return static_cast<QSqlTableModel*>(ptr)->mimeData(*static_cast<QList<QModelIndex>*>(indexes));
}

void* QSqlTableModel_MimeDataDefault(void* ptr, void* indexes)
{
	return static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::mimeData(*static_cast<QList<QModelIndex>*>(indexes));
}

struct QtSql_PackedString QSqlTableModel_MimeTypes(void* ptr)
{
	return ({ QByteArray tabeef0 = static_cast<QSqlTableModel*>(ptr)->mimeTypes().join("|").toUtf8(); QtSql_PackedString { const_cast<char*>(tabeef0.prepend("WHITESPACE").constData()+10), tabeef0.size()-10 }; });
}

struct QtSql_PackedString QSqlTableModel_MimeTypesDefault(void* ptr)
{
	return ({ QByteArray t97b0b0 = static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::mimeTypes().join("|").toUtf8(); QtSql_PackedString { const_cast<char*>(t97b0b0.prepend("WHITESPACE").constData()+10), t97b0b0.size()-10 }; });
}

char QSqlTableModel_MoveColumns(void* ptr, void* sourceParent, int sourceColumn, int count, void* destinationParent, int destinationChild)
{
	return static_cast<QSqlTableModel*>(ptr)->moveColumns(*static_cast<QModelIndex*>(sourceParent), sourceColumn, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

char QSqlTableModel_MoveColumnsDefault(void* ptr, void* sourceParent, int sourceColumn, int count, void* destinationParent, int destinationChild)
{
	return static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::moveColumns(*static_cast<QModelIndex*>(sourceParent), sourceColumn, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

char QSqlTableModel_MoveRows(void* ptr, void* sourceParent, int sourceRow, int count, void* destinationParent, int destinationChild)
{
	return static_cast<QSqlTableModel*>(ptr)->moveRows(*static_cast<QModelIndex*>(sourceParent), sourceRow, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

char QSqlTableModel_MoveRowsDefault(void* ptr, void* sourceParent, int sourceRow, int count, void* destinationParent, int destinationChild)
{
	return static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::moveRows(*static_cast<QModelIndex*>(sourceParent), sourceRow, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

void* QSqlTableModel_Parent(void* ptr, void* index)
{
	return new QModelIndex(static_cast<QSqlTableModel*>(ptr)->parent(*static_cast<QModelIndex*>(index)));
}

void QSqlTableModel_ResetInternalData(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QSqlTableModel*>(ptr), "resetInternalData");
}

void QSqlTableModel_ResetInternalDataDefault(void* ptr)
{
	static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::resetInternalData();
}

struct QtSql_PackedList QSqlTableModel_RoleNames(void* ptr)
{
	return ({ QHash<int, QByteArray>* tmpValue = new QHash<int, QByteArray>(static_cast<QSqlTableModel*>(ptr)->roleNames()); QtSql_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtSql_PackedList QSqlTableModel_RoleNamesDefault(void* ptr)
{
	return ({ QHash<int, QByteArray>* tmpValue = new QHash<int, QByteArray>(static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::roleNames()); QtSql_PackedList { tmpValue, tmpValue->size() }; });
}

char QSqlTableModel_SetItemData(void* ptr, void* index, void* roles)
{
	return static_cast<QSqlTableModel*>(ptr)->setItemData(*static_cast<QModelIndex*>(index), *static_cast<QMap<int, QVariant>*>(roles));
}

char QSqlTableModel_SetItemDataDefault(void* ptr, void* index, void* roles)
{
	return static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::setItemData(*static_cast<QModelIndex*>(index), *static_cast<QMap<int, QVariant>*>(roles));
}

void* QSqlTableModel_Span(void* ptr, void* index)
{
	return ({ QSize tmpValue = static_cast<QSqlTableModel*>(ptr)->span(*static_cast<QModelIndex*>(index)); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QSqlTableModel_SpanDefault(void* ptr, void* index)
{
	return ({ QSize tmpValue = static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::span(*static_cast<QModelIndex*>(index)); new QSize(tmpValue.width(), tmpValue.height()); });
}

long long QSqlTableModel_SupportedDragActions(void* ptr)
{
	return static_cast<QSqlTableModel*>(ptr)->supportedDragActions();
}

long long QSqlTableModel_SupportedDragActionsDefault(void* ptr)
{
	return static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::supportedDragActions();
}

long long QSqlTableModel_SupportedDropActions(void* ptr)
{
	return static_cast<QSqlTableModel*>(ptr)->supportedDropActions();
}

long long QSqlTableModel_SupportedDropActionsDefault(void* ptr)
{
	return static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::supportedDropActions();
}

void QSqlTableModel_TimerEvent(void* ptr, void* event)
{
	static_cast<QSqlTableModel*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QSqlTableModel_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::timerEvent(static_cast<QTimerEvent*>(event));
}

void QSqlTableModel_ChildEvent(void* ptr, void* event)
{
	static_cast<QSqlTableModel*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QSqlTableModel_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::childEvent(static_cast<QChildEvent*>(event));
}

void QSqlTableModel_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QSqlTableModel*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSqlTableModel_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSqlTableModel_CustomEvent(void* ptr, void* event)
{
	static_cast<QSqlTableModel*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QSqlTableModel_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::customEvent(static_cast<QEvent*>(event));
}

void QSqlTableModel_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QSqlTableModel*>(ptr), "deleteLater");
}

void QSqlTableModel_DeleteLaterDefault(void* ptr)
{
	static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::deleteLater();
}

void QSqlTableModel_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QSqlTableModel*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSqlTableModel_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QSqlTableModel_Event(void* ptr, void* e)
{
	return static_cast<QSqlTableModel*>(ptr)->event(static_cast<QEvent*>(e));
}

char QSqlTableModel_EventDefault(void* ptr, void* e)
{
	return static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::event(static_cast<QEvent*>(e));
}

char QSqlTableModel_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QSqlTableModel*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QSqlTableModel_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QSqlTableModel_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QSqlTableModel*>(ptr)->metaObject());
}

void* QSqlTableModel_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::metaObject());
}

