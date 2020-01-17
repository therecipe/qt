// +build !minimal

#define protected public
#define private public

#include "sql.h"
#include "_cgo_export.h"

#include <QAbstractItemDelegate>
#include <QAbstractItemModel>
#include <QAbstractItemView>
#include <QAudioSystemPlugin>
#include <QByteArray>
#include <QCameraImageCapture>
#include <QChildEvent>
#include <QDBusPendingCallWatcher>
#include <QEvent>
#include <QExtensionFactory>
#include <QExtensionManager>
#include <QGraphicsObject>
#include <QGraphicsWidget>
#include <QHash>
#include <QHelpEvent>
#include <QLayout>
#include <QMap>
#include <QMediaPlaylist>
#include <QMediaRecorder>
#include <QMediaServiceProviderPlugin>
#include <QMetaMethod>
#include <QMetaObject>
#include <QMimeData>
#include <QModelIndex>
#include <QObject>
#include <QOffscreenSurface>
#include <QPaintDeviceWindow>
#include <QPainter>
#include <QPdfWriter>
#include <QPersistentModelIndex>
#include <QPixmap>
#include <QQuickItem>
#include <QRadioData>
#include <QRect>
#include <QRemoteObjectPendingCallWatcher>
#include <QScriptExtensionPlugin>
#include <QSize>
#include <QSqlDatabase>
#include <QSqlDriver>
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
#include <QStyleOptionViewItem>
#include <QTimerEvent>
#include <QVariant>
#include <QVector>
#include <QWidget>
#include <QWindow>

Q_DECLARE_METATYPE(QSqlDatabase)
Q_DECLARE_METATYPE(QSqlDatabase*)
void* QSqlDatabase_NewQSqlDatabase()
{
	return new QSqlDatabase();
}

void* QSqlDatabase_NewQSqlDatabase2(void* other)
{
	return new QSqlDatabase(*static_cast<QSqlDatabase*>(other));
}

void* QSqlDatabase_NewQSqlDatabase3(struct QtSql_PackedString ty)
{
	return new QSqlDatabase(QString::fromUtf8(ty.data, ty.len));
}

void* QSqlDatabase_NewQSqlDatabase4(void* driver)
{
	return new QSqlDatabase(static_cast<QSqlDriver*>(driver));
}

void* QSqlDatabase_QSqlDatabase_AddDatabase(struct QtSql_PackedString ty, struct QtSql_PackedString connectionName)
{
	return new QSqlDatabase(QSqlDatabase::addDatabase(QString::fromUtf8(ty.data, ty.len), QString::fromUtf8(connectionName.data, connectionName.len)));
}

void* QSqlDatabase_QSqlDatabase_AddDatabase2(void* driver, struct QtSql_PackedString connectionName)
{
	return new QSqlDatabase(QSqlDatabase::addDatabase(static_cast<QSqlDriver*>(driver), QString::fromUtf8(connectionName.data, connectionName.len)));
}

void* QSqlDatabase_QSqlDatabase_CloneDatabase(void* other, struct QtSql_PackedString connectionName)
{
	return new QSqlDatabase(QSqlDatabase::cloneDatabase(*static_cast<QSqlDatabase*>(other), QString::fromUtf8(connectionName.data, connectionName.len)));
}

void* QSqlDatabase_QSqlDatabase_CloneDatabase2(struct QtSql_PackedString other, struct QtSql_PackedString connectionName)
{
	return new QSqlDatabase(QSqlDatabase::cloneDatabase(QString::fromUtf8(other.data, other.len), QString::fromUtf8(connectionName.data, connectionName.len)));
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
	return ({ QByteArray* t27c06b = new QByteArray(static_cast<QSqlDatabase*>(ptr)->connectOptions().toUtf8()); QtSql_PackedString { const_cast<char*>(t27c06b->prepend("WHITESPACE").constData()+10), t27c06b->size()-10, t27c06b }; });
}

struct QtSql_PackedString QSqlDatabase_ConnectionName(void* ptr)
{
	return ({ QByteArray* tb2dc55 = new QByteArray(static_cast<QSqlDatabase*>(ptr)->connectionName().toUtf8()); QtSql_PackedString { const_cast<char*>(tb2dc55->prepend("WHITESPACE").constData()+10), tb2dc55->size()-10, tb2dc55 }; });
}

struct QtSql_PackedString QSqlDatabase_QSqlDatabase_ConnectionNames()
{
	return ({ QByteArray* t34c9e9 = new QByteArray(QSqlDatabase::connectionNames().join("¡¦!").toUtf8()); QtSql_PackedString { const_cast<char*>(t34c9e9->prepend("WHITESPACE").constData()+10), t34c9e9->size()-10, t34c9e9 }; });
}

char QSqlDatabase_QSqlDatabase_Contains(struct QtSql_PackedString connectionName)
{
	return QSqlDatabase::contains(QString::fromUtf8(connectionName.data, connectionName.len));
}

void* QSqlDatabase_QSqlDatabase_Database(struct QtSql_PackedString connectionName, char open)
{
	return new QSqlDatabase(QSqlDatabase::database(QString::fromUtf8(connectionName.data, connectionName.len), open != 0));
}

struct QtSql_PackedString QSqlDatabase_DatabaseName(void* ptr)
{
	return ({ QByteArray* t15dc5d = new QByteArray(static_cast<QSqlDatabase*>(ptr)->databaseName().toUtf8()); QtSql_PackedString { const_cast<char*>(t15dc5d->prepend("WHITESPACE").constData()+10), t15dc5d->size()-10, t15dc5d }; });
}

void* QSqlDatabase_Driver(void* ptr)
{
	return static_cast<QSqlDatabase*>(ptr)->driver();
}

struct QtSql_PackedString QSqlDatabase_DriverName(void* ptr)
{
	return ({ QByteArray* t58c861 = new QByteArray(static_cast<QSqlDatabase*>(ptr)->driverName().toUtf8()); QtSql_PackedString { const_cast<char*>(t58c861->prepend("WHITESPACE").constData()+10), t58c861->size()-10, t58c861 }; });
}

struct QtSql_PackedString QSqlDatabase_QSqlDatabase_Drivers()
{
	return ({ QByteArray* tf3b332 = new QByteArray(QSqlDatabase::drivers().join("¡¦!").toUtf8()); QtSql_PackedString { const_cast<char*>(tf3b332->prepend("WHITESPACE").constData()+10), tf3b332->size()-10, tf3b332 }; });
}

void* QSqlDatabase_Exec(void* ptr, struct QtSql_PackedString query)
{
	return new QSqlQuery(static_cast<QSqlDatabase*>(ptr)->exec(QString::fromUtf8(query.data, query.len)));
}

struct QtSql_PackedString QSqlDatabase_HostName(void* ptr)
{
	return ({ QByteArray* taeb29f = new QByteArray(static_cast<QSqlDatabase*>(ptr)->hostName().toUtf8()); QtSql_PackedString { const_cast<char*>(taeb29f->prepend("WHITESPACE").constData()+10), taeb29f->size()-10, taeb29f }; });
}

char QSqlDatabase_QSqlDatabase_IsDriverAvailable(struct QtSql_PackedString name)
{
	return QSqlDatabase::isDriverAvailable(QString::fromUtf8(name.data, name.len));
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

char QSqlDatabase_Open2(void* ptr, struct QtSql_PackedString user, struct QtSql_PackedString password)
{
	return static_cast<QSqlDatabase*>(ptr)->open(QString::fromUtf8(user.data, user.len), QString::fromUtf8(password.data, password.len));
}

struct QtSql_PackedString QSqlDatabase_Password(void* ptr)
{
	return ({ QByteArray* t136367 = new QByteArray(static_cast<QSqlDatabase*>(ptr)->password().toUtf8()); QtSql_PackedString { const_cast<char*>(t136367->prepend("WHITESPACE").constData()+10), t136367->size()-10, t136367 }; });
}

int QSqlDatabase_Port(void* ptr)
{
	return static_cast<QSqlDatabase*>(ptr)->port();
}

void* QSqlDatabase_PrimaryIndex(void* ptr, struct QtSql_PackedString tablename)
{
	return new QSqlIndex(static_cast<QSqlDatabase*>(ptr)->primaryIndex(QString::fromUtf8(tablename.data, tablename.len)));
}

void* QSqlDatabase_Record(void* ptr, struct QtSql_PackedString tablename)
{
	return new QSqlRecord(static_cast<QSqlDatabase*>(ptr)->record(QString::fromUtf8(tablename.data, tablename.len)));
}

void QSqlDatabase_QSqlDatabase_RegisterSqlDriver(struct QtSql_PackedString name, void* creator)
{
	QSqlDatabase::registerSqlDriver(QString::fromUtf8(name.data, name.len), static_cast<QSqlDriverCreatorBase*>(creator));
}

void QSqlDatabase_QSqlDatabase_RemoveDatabase(struct QtSql_PackedString connectionName)
{
	QSqlDatabase::removeDatabase(QString::fromUtf8(connectionName.data, connectionName.len));
}

char QSqlDatabase_Rollback(void* ptr)
{
	return static_cast<QSqlDatabase*>(ptr)->rollback();
}

void QSqlDatabase_SetConnectOptions(void* ptr, struct QtSql_PackedString options)
{
	static_cast<QSqlDatabase*>(ptr)->setConnectOptions(QString::fromUtf8(options.data, options.len));
}

void QSqlDatabase_SetDatabaseName(void* ptr, struct QtSql_PackedString name)
{
	static_cast<QSqlDatabase*>(ptr)->setDatabaseName(QString::fromUtf8(name.data, name.len));
}

void QSqlDatabase_SetHostName(void* ptr, struct QtSql_PackedString host)
{
	static_cast<QSqlDatabase*>(ptr)->setHostName(QString::fromUtf8(host.data, host.len));
}

void QSqlDatabase_SetNumericalPrecisionPolicy(void* ptr, long long precisionPolicy)
{
	static_cast<QSqlDatabase*>(ptr)->setNumericalPrecisionPolicy(static_cast<QSql::NumericalPrecisionPolicy>(precisionPolicy));
}

void QSqlDatabase_SetPassword(void* ptr, struct QtSql_PackedString password)
{
	static_cast<QSqlDatabase*>(ptr)->setPassword(QString::fromUtf8(password.data, password.len));
}

void QSqlDatabase_SetPort(void* ptr, int port)
{
	static_cast<QSqlDatabase*>(ptr)->setPort(port);
}

void QSqlDatabase_SetUserName(void* ptr, struct QtSql_PackedString name)
{
	static_cast<QSqlDatabase*>(ptr)->setUserName(QString::fromUtf8(name.data, name.len));
}

struct QtSql_PackedString QSqlDatabase_Tables(void* ptr, long long ty)
{
	return ({ QByteArray* t302bb5 = new QByteArray(static_cast<QSqlDatabase*>(ptr)->tables(static_cast<QSql::TableType>(ty)).join("¡¦!").toUtf8()); QtSql_PackedString { const_cast<char*>(t302bb5->prepend("WHITESPACE").constData()+10), t302bb5->size()-10, t302bb5 }; });
}

char QSqlDatabase_Transaction(void* ptr)
{
	return static_cast<QSqlDatabase*>(ptr)->transaction();
}

struct QtSql_PackedString QSqlDatabase_UserName(void* ptr)
{
	return ({ QByteArray* t7a5cf4 = new QByteArray(static_cast<QSqlDatabase*>(ptr)->userName().toUtf8()); QtSql_PackedString { const_cast<char*>(t7a5cf4->prepend("WHITESPACE").constData()+10), t7a5cf4->size()-10, t7a5cf4 }; });
}

void QSqlDatabase_DestroyQSqlDatabase(void* ptr)
{
	static_cast<QSqlDatabase*>(ptr)->~QSqlDatabase();
}

class MyQSqlDriver: public QSqlDriver
{
public:
	MyQSqlDriver(QObject *parent = Q_NULLPTR) : QSqlDriver(parent) {QSqlDriver_QSqlDriver_QRegisterMetaType();};
	bool beginTransaction() { return callbackQSqlDriver_BeginTransaction(this) != 0; };
	void close() { callbackQSqlDriver_Close(this); };
	bool commitTransaction() { return callbackQSqlDriver_CommitTransaction(this) != 0; };
	QSqlResult * createResult() const { return static_cast<QSqlResult*>(callbackQSqlDriver_CreateResult(const_cast<void*>(static_cast<const void*>(this)))); };
	QString escapeIdentifier(const QString & identifier, QSqlDriver::IdentifierType ty) const { QByteArray* tfae9fd = new QByteArray(identifier.toUtf8()); QtSql_PackedString identifierPacked = { const_cast<char*>(tfae9fd->prepend("WHITESPACE").constData()+10), tfae9fd->size()-10, tfae9fd };return ({ QtSql_PackedString tempVal = callbackQSqlDriver_EscapeIdentifier(const_cast<void*>(static_cast<const void*>(this)), identifierPacked, ty); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	QString formatValue(const QSqlField & field, bool trimStrings) const { return ({ QtSql_PackedString tempVal = callbackQSqlDriver_FormatValue(const_cast<void*>(static_cast<const void*>(this)), const_cast<QSqlField*>(&field), trimStrings); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	QVariant handle() const { return *static_cast<QVariant*>(callbackQSqlDriver_Handle(const_cast<void*>(static_cast<const void*>(this)))); };
	bool hasFeature(QSqlDriver::DriverFeature feature) const { return callbackQSqlDriver_HasFeature(const_cast<void*>(static_cast<const void*>(this)), feature) != 0; };
	bool isIdentifierEscaped(const QString & identifier, QSqlDriver::IdentifierType ty) const { QByteArray* tfae9fd = new QByteArray(identifier.toUtf8()); QtSql_PackedString identifierPacked = { const_cast<char*>(tfae9fd->prepend("WHITESPACE").constData()+10), tfae9fd->size()-10, tfae9fd };return callbackQSqlDriver_IsIdentifierEscaped(const_cast<void*>(static_cast<const void*>(this)), identifierPacked, ty) != 0; };
	bool isOpen() const { return callbackQSqlDriver_IsOpen(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	void Signal_Notification(const QString & name) { QByteArray* t6ae999 = new QByteArray(name.toUtf8()); QtSql_PackedString namePacked = { const_cast<char*>(t6ae999->prepend("WHITESPACE").constData()+10), t6ae999->size()-10, t6ae999 };callbackQSqlDriver_Notification(this, namePacked); };
	void Signal_Notification2(const QString & name, QSqlDriver::NotificationSource source, const QVariant & payload) { QByteArray* t6ae999 = new QByteArray(name.toUtf8()); QtSql_PackedString namePacked = { const_cast<char*>(t6ae999->prepend("WHITESPACE").constData()+10), t6ae999->size()-10, t6ae999 };callbackQSqlDriver_Notification2(this, namePacked, source, const_cast<QVariant*>(&payload)); };
	bool open(const QString & db, const QString & user, const QString & password, const QString & host, int port, const QString & options) { QByteArray* t0352a8 = new QByteArray(db.toUtf8()); QtSql_PackedString dbPacked = { const_cast<char*>(t0352a8->prepend("WHITESPACE").constData()+10), t0352a8->size()-10, t0352a8 };QByteArray* t12dea9 = new QByteArray(user.toUtf8()); QtSql_PackedString userPacked = { const_cast<char*>(t12dea9->prepend("WHITESPACE").constData()+10), t12dea9->size()-10, t12dea9 };QByteArray* t5baa61 = new QByteArray(password.toUtf8()); QtSql_PackedString passwordPacked = { const_cast<char*>(t5baa61->prepend("WHITESPACE").constData()+10), t5baa61->size()-10, t5baa61 };QByteArray* t86dd1c = new QByteArray(host.toUtf8()); QtSql_PackedString hostPacked = { const_cast<char*>(t86dd1c->prepend("WHITESPACE").constData()+10), t86dd1c->size()-10, t86dd1c };QByteArray* t513f8d = new QByteArray(options.toUtf8()); QtSql_PackedString optionsPacked = { const_cast<char*>(t513f8d->prepend("WHITESPACE").constData()+10), t513f8d->size()-10, t513f8d };return callbackQSqlDriver_Open(this, dbPacked, userPacked, passwordPacked, hostPacked, port, optionsPacked) != 0; };
	QSqlIndex primaryIndex(const QString & tableName) const { QByteArray* t3e7060 = new QByteArray(tableName.toUtf8()); QtSql_PackedString tableNamePacked = { const_cast<char*>(t3e7060->prepend("WHITESPACE").constData()+10), t3e7060->size()-10, t3e7060 };return *static_cast<QSqlIndex*>(callbackQSqlDriver_PrimaryIndex(const_cast<void*>(static_cast<const void*>(this)), tableNamePacked)); };
	QSqlRecord record(const QString & tableName) const { QByteArray* t3e7060 = new QByteArray(tableName.toUtf8()); QtSql_PackedString tableNamePacked = { const_cast<char*>(t3e7060->prepend("WHITESPACE").constData()+10), t3e7060->size()-10, t3e7060 };return *static_cast<QSqlRecord*>(callbackQSqlDriver_Record(const_cast<void*>(static_cast<const void*>(this)), tableNamePacked)); };
	bool rollbackTransaction() { return callbackQSqlDriver_RollbackTransaction(this) != 0; };
	void setLastError(const QSqlError & error) { callbackQSqlDriver_SetLastError(this, const_cast<QSqlError*>(&error)); };
	void setOpen(bool open) { callbackQSqlDriver_SetOpen(this, open); };
	void setOpenError(bool error) { callbackQSqlDriver_SetOpenError(this, error); };
	QString sqlStatement(QSqlDriver::StatementType ty, const QString & tableName, const QSqlRecord & rec, bool preparedStatement) const { QByteArray* t3e7060 = new QByteArray(tableName.toUtf8()); QtSql_PackedString tableNamePacked = { const_cast<char*>(t3e7060->prepend("WHITESPACE").constData()+10), t3e7060->size()-10, t3e7060 };return ({ QtSql_PackedString tempVal = callbackQSqlDriver_SqlStatement(const_cast<void*>(static_cast<const void*>(this)), ty, tableNamePacked, const_cast<QSqlRecord*>(&rec), preparedStatement); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	QString stripDelimiters(const QString & identifier, QSqlDriver::IdentifierType ty) const { QByteArray* tfae9fd = new QByteArray(identifier.toUtf8()); QtSql_PackedString identifierPacked = { const_cast<char*>(tfae9fd->prepend("WHITESPACE").constData()+10), tfae9fd->size()-10, tfae9fd };return ({ QtSql_PackedString tempVal = callbackQSqlDriver_StripDelimiters(const_cast<void*>(static_cast<const void*>(this)), identifierPacked, ty); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	bool subscribeToNotification(const QString & name) { QByteArray* t6ae999 = new QByteArray(name.toUtf8()); QtSql_PackedString namePacked = { const_cast<char*>(t6ae999->prepend("WHITESPACE").constData()+10), t6ae999->size()-10, t6ae999 };return callbackQSqlDriver_SubscribeToNotification(this, namePacked) != 0; };
	QStringList subscribedToNotifications() const { return ({ QtSql_PackedString tempVal = callbackQSqlDriver_SubscribedToNotifications(const_cast<void*>(static_cast<const void*>(this))); QStringList ret = QString::fromUtf8(tempVal.data, tempVal.len).split("¡¦!", QString::SkipEmptyParts); free(tempVal.data); ret; }); };
	QStringList tables(QSql::TableType tableType) const { return ({ QtSql_PackedString tempVal = callbackQSqlDriver_Tables(const_cast<void*>(static_cast<const void*>(this)), tableType); QStringList ret = QString::fromUtf8(tempVal.data, tempVal.len).split("¡¦!", QString::SkipEmptyParts); free(tempVal.data); ret; }); };
	bool unsubscribeFromNotification(const QString & name) { QByteArray* t6ae999 = new QByteArray(name.toUtf8()); QtSql_PackedString namePacked = { const_cast<char*>(t6ae999->prepend("WHITESPACE").constData()+10), t6ae999->size()-10, t6ae999 };return callbackQSqlDriver_UnsubscribeFromNotification(this, namePacked) != 0; };
	 ~MyQSqlDriver() { callbackQSqlDriver_DestroyQSqlDriver(this); };
	void childEvent(QChildEvent * event) { callbackQSqlDriver_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSqlDriver_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQSqlDriver_CustomEvent(this, event); };
	void deleteLater() { callbackQSqlDriver_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQSqlDriver_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSqlDriver_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQSqlDriver_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSqlDriver_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSqlDriver_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtSql_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQSqlDriver_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQSqlDriver_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QSqlDriver*)
Q_DECLARE_METATYPE(MyQSqlDriver*)

int QSqlDriver_QSqlDriver_QRegisterMetaType(){qRegisterMetaType<QSqlDriver*>(); return qRegisterMetaType<MyQSqlDriver*>();}

void* QSqlDriver_NewQSqlDriver(void* parent)
{
	if (dynamic_cast<QAudioSystemPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQSqlDriver(static_cast<QAudioSystemPlugin*>(parent));
	} else if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQSqlDriver(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQSqlDriver(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQSqlDriver(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQSqlDriver(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQSqlDriver(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQSqlDriver(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQSqlDriver(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQSqlDriver(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQSqlDriver(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QMediaServiceProviderPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQSqlDriver(static_cast<QMediaServiceProviderPlugin*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQSqlDriver(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQSqlDriver(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQSqlDriver(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQSqlDriver(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQSqlDriver(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQSqlDriver(static_cast<QRemoteObjectPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QScriptExtensionPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQSqlDriver(static_cast<QScriptExtensionPlugin*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQSqlDriver(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQSqlDriver(static_cast<QWindow*>(parent));
	} else {
		return new MyQSqlDriver(static_cast<QObject*>(parent));
	}
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

struct QtSql_PackedString QSqlDriver_EscapeIdentifier(void* ptr, struct QtSql_PackedString identifier, long long ty)
{
	return ({ QByteArray* t179a9d = new QByteArray(static_cast<QSqlDriver*>(ptr)->escapeIdentifier(QString::fromUtf8(identifier.data, identifier.len), static_cast<QSqlDriver::IdentifierType>(ty)).toUtf8()); QtSql_PackedString { const_cast<char*>(t179a9d->prepend("WHITESPACE").constData()+10), t179a9d->size()-10, t179a9d }; });
}

struct QtSql_PackedString QSqlDriver_EscapeIdentifierDefault(void* ptr, struct QtSql_PackedString identifier, long long ty)
{
		return ({ QByteArray* t42081b = new QByteArray(static_cast<QSqlDriver*>(ptr)->QSqlDriver::escapeIdentifier(QString::fromUtf8(identifier.data, identifier.len), static_cast<QSqlDriver::IdentifierType>(ty)).toUtf8()); QtSql_PackedString { const_cast<char*>(t42081b->prepend("WHITESPACE").constData()+10), t42081b->size()-10, t42081b }; });
}

struct QtSql_PackedString QSqlDriver_FormatValue(void* ptr, void* field, char trimStrings)
{
	return ({ QByteArray* t09cd9b = new QByteArray(static_cast<QSqlDriver*>(ptr)->formatValue(*static_cast<QSqlField*>(field), trimStrings != 0).toUtf8()); QtSql_PackedString { const_cast<char*>(t09cd9b->prepend("WHITESPACE").constData()+10), t09cd9b->size()-10, t09cd9b }; });
}

struct QtSql_PackedString QSqlDriver_FormatValueDefault(void* ptr, void* field, char trimStrings)
{
		return ({ QByteArray* tcdc4d9 = new QByteArray(static_cast<QSqlDriver*>(ptr)->QSqlDriver::formatValue(*static_cast<QSqlField*>(field), trimStrings != 0).toUtf8()); QtSql_PackedString { const_cast<char*>(tcdc4d9->prepend("WHITESPACE").constData()+10), tcdc4d9->size()-10, tcdc4d9 }; });
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

char QSqlDriver_IsIdentifierEscaped(void* ptr, struct QtSql_PackedString identifier, long long ty)
{
	return static_cast<QSqlDriver*>(ptr)->isIdentifierEscaped(QString::fromUtf8(identifier.data, identifier.len), static_cast<QSqlDriver::IdentifierType>(ty));
}

char QSqlDriver_IsIdentifierEscapedDefault(void* ptr, struct QtSql_PackedString identifier, long long ty)
{
		return static_cast<QSqlDriver*>(ptr)->QSqlDriver::isIdentifierEscaped(QString::fromUtf8(identifier.data, identifier.len), static_cast<QSqlDriver::IdentifierType>(ty));
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

void QSqlDriver_ConnectNotification(void* ptr, long long t)
{
	QObject::connect(static_cast<QSqlDriver*>(ptr), static_cast<void (QSqlDriver::*)(const QString &)>(&QSqlDriver::notification), static_cast<MyQSqlDriver*>(ptr), static_cast<void (MyQSqlDriver::*)(const QString &)>(&MyQSqlDriver::Signal_Notification), static_cast<Qt::ConnectionType>(t));
}

void QSqlDriver_DisconnectNotification(void* ptr)
{
	QObject::disconnect(static_cast<QSqlDriver*>(ptr), static_cast<void (QSqlDriver::*)(const QString &)>(&QSqlDriver::notification), static_cast<MyQSqlDriver*>(ptr), static_cast<void (MyQSqlDriver::*)(const QString &)>(&MyQSqlDriver::Signal_Notification));
}

void QSqlDriver_Notification(void* ptr, struct QtSql_PackedString name)
{
	static_cast<QSqlDriver*>(ptr)->notification(QString::fromUtf8(name.data, name.len));
}

void QSqlDriver_ConnectNotification2(void* ptr, long long t)
{
	QObject::connect(static_cast<QSqlDriver*>(ptr), static_cast<void (QSqlDriver::*)(const QString &, QSqlDriver::NotificationSource, const QVariant &)>(&QSqlDriver::notification), static_cast<MyQSqlDriver*>(ptr), static_cast<void (MyQSqlDriver::*)(const QString &, QSqlDriver::NotificationSource, const QVariant &)>(&MyQSqlDriver::Signal_Notification2), static_cast<Qt::ConnectionType>(t));
}

void QSqlDriver_DisconnectNotification2(void* ptr)
{
	QObject::disconnect(static_cast<QSqlDriver*>(ptr), static_cast<void (QSqlDriver::*)(const QString &, QSqlDriver::NotificationSource, const QVariant &)>(&QSqlDriver::notification), static_cast<MyQSqlDriver*>(ptr), static_cast<void (MyQSqlDriver::*)(const QString &, QSqlDriver::NotificationSource, const QVariant &)>(&MyQSqlDriver::Signal_Notification2));
}

void QSqlDriver_Notification2(void* ptr, struct QtSql_PackedString name, long long source, void* payload)
{
	static_cast<QSqlDriver*>(ptr)->notification(QString::fromUtf8(name.data, name.len), static_cast<QSqlDriver::NotificationSource>(source), *static_cast<QVariant*>(payload));
}

long long QSqlDriver_NumericalPrecisionPolicy(void* ptr)
{
	return static_cast<QSqlDriver*>(ptr)->numericalPrecisionPolicy();
}

char QSqlDriver_Open(void* ptr, struct QtSql_PackedString db, struct QtSql_PackedString user, struct QtSql_PackedString password, struct QtSql_PackedString host, int port, struct QtSql_PackedString options)
{
	return static_cast<QSqlDriver*>(ptr)->open(QString::fromUtf8(db.data, db.len), QString::fromUtf8(user.data, user.len), QString::fromUtf8(password.data, password.len), QString::fromUtf8(host.data, host.len), port, QString::fromUtf8(options.data, options.len));
}

void* QSqlDriver_PrimaryIndex(void* ptr, struct QtSql_PackedString tableName)
{
	return new QSqlIndex(static_cast<QSqlDriver*>(ptr)->primaryIndex(QString::fromUtf8(tableName.data, tableName.len)));
}

void* QSqlDriver_PrimaryIndexDefault(void* ptr, struct QtSql_PackedString tableName)
{
		return new QSqlIndex(static_cast<QSqlDriver*>(ptr)->QSqlDriver::primaryIndex(QString::fromUtf8(tableName.data, tableName.len)));
}

void* QSqlDriver_Record(void* ptr, struct QtSql_PackedString tableName)
{
	return new QSqlRecord(static_cast<QSqlDriver*>(ptr)->record(QString::fromUtf8(tableName.data, tableName.len)));
}

void* QSqlDriver_RecordDefault(void* ptr, struct QtSql_PackedString tableName)
{
		return new QSqlRecord(static_cast<QSqlDriver*>(ptr)->QSqlDriver::record(QString::fromUtf8(tableName.data, tableName.len)));
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

struct QtSql_PackedString QSqlDriver_SqlStatement(void* ptr, long long ty, struct QtSql_PackedString tableName, void* rec, char preparedStatement)
{
	return ({ QByteArray* t404912 = new QByteArray(static_cast<QSqlDriver*>(ptr)->sqlStatement(static_cast<QSqlDriver::StatementType>(ty), QString::fromUtf8(tableName.data, tableName.len), *static_cast<QSqlRecord*>(rec), preparedStatement != 0).toUtf8()); QtSql_PackedString { const_cast<char*>(t404912->prepend("WHITESPACE").constData()+10), t404912->size()-10, t404912 }; });
}

struct QtSql_PackedString QSqlDriver_SqlStatementDefault(void* ptr, long long ty, struct QtSql_PackedString tableName, void* rec, char preparedStatement)
{
		return ({ QByteArray* t50868f = new QByteArray(static_cast<QSqlDriver*>(ptr)->QSqlDriver::sqlStatement(static_cast<QSqlDriver::StatementType>(ty), QString::fromUtf8(tableName.data, tableName.len), *static_cast<QSqlRecord*>(rec), preparedStatement != 0).toUtf8()); QtSql_PackedString { const_cast<char*>(t50868f->prepend("WHITESPACE").constData()+10), t50868f->size()-10, t50868f }; });
}

struct QtSql_PackedString QSqlDriver_StripDelimiters(void* ptr, struct QtSql_PackedString identifier, long long ty)
{
	return ({ QByteArray* t8be107 = new QByteArray(static_cast<QSqlDriver*>(ptr)->stripDelimiters(QString::fromUtf8(identifier.data, identifier.len), static_cast<QSqlDriver::IdentifierType>(ty)).toUtf8()); QtSql_PackedString { const_cast<char*>(t8be107->prepend("WHITESPACE").constData()+10), t8be107->size()-10, t8be107 }; });
}

struct QtSql_PackedString QSqlDriver_StripDelimitersDefault(void* ptr, struct QtSql_PackedString identifier, long long ty)
{
		return ({ QByteArray* tb71792 = new QByteArray(static_cast<QSqlDriver*>(ptr)->QSqlDriver::stripDelimiters(QString::fromUtf8(identifier.data, identifier.len), static_cast<QSqlDriver::IdentifierType>(ty)).toUtf8()); QtSql_PackedString { const_cast<char*>(tb71792->prepend("WHITESPACE").constData()+10), tb71792->size()-10, tb71792 }; });
}

char QSqlDriver_SubscribeToNotification(void* ptr, struct QtSql_PackedString name)
{
	return static_cast<QSqlDriver*>(ptr)->subscribeToNotification(QString::fromUtf8(name.data, name.len));
}

char QSqlDriver_SubscribeToNotificationDefault(void* ptr, struct QtSql_PackedString name)
{
		return static_cast<QSqlDriver*>(ptr)->QSqlDriver::subscribeToNotification(QString::fromUtf8(name.data, name.len));
}

struct QtSql_PackedString QSqlDriver_SubscribedToNotifications(void* ptr)
{
	return ({ QByteArray* t545d09 = new QByteArray(static_cast<QSqlDriver*>(ptr)->subscribedToNotifications().join("¡¦!").toUtf8()); QtSql_PackedString { const_cast<char*>(t545d09->prepend("WHITESPACE").constData()+10), t545d09->size()-10, t545d09 }; });
}

struct QtSql_PackedString QSqlDriver_SubscribedToNotificationsDefault(void* ptr)
{
		return ({ QByteArray* t063c3c = new QByteArray(static_cast<QSqlDriver*>(ptr)->QSqlDriver::subscribedToNotifications().join("¡¦!").toUtf8()); QtSql_PackedString { const_cast<char*>(t063c3c->prepend("WHITESPACE").constData()+10), t063c3c->size()-10, t063c3c }; });
}

struct QtSql_PackedString QSqlDriver_Tables(void* ptr, long long tableType)
{
	return ({ QByteArray* tcfb904 = new QByteArray(static_cast<QSqlDriver*>(ptr)->tables(static_cast<QSql::TableType>(tableType)).join("¡¦!").toUtf8()); QtSql_PackedString { const_cast<char*>(tcfb904->prepend("WHITESPACE").constData()+10), tcfb904->size()-10, tcfb904 }; });
}

struct QtSql_PackedString QSqlDriver_TablesDefault(void* ptr, long long tableType)
{
		return ({ QByteArray* tf0f7df = new QByteArray(static_cast<QSqlDriver*>(ptr)->QSqlDriver::tables(static_cast<QSql::TableType>(tableType)).join("¡¦!").toUtf8()); QtSql_PackedString { const_cast<char*>(tf0f7df->prepend("WHITESPACE").constData()+10), tf0f7df->size()-10, tf0f7df }; });
}

char QSqlDriver_UnsubscribeFromNotification(void* ptr, struct QtSql_PackedString name)
{
	return static_cast<QSqlDriver*>(ptr)->unsubscribeFromNotification(QString::fromUtf8(name.data, name.len));
}

char QSqlDriver_UnsubscribeFromNotificationDefault(void* ptr, struct QtSql_PackedString name)
{
		return static_cast<QSqlDriver*>(ptr)->QSqlDriver::unsubscribeFromNotification(QString::fromUtf8(name.data, name.len));
}

void QSqlDriver_DestroyQSqlDriver(void* ptr)
{
	static_cast<QSqlDriver*>(ptr)->~QSqlDriver();
}

void QSqlDriver_DestroyQSqlDriverDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QSqlDriver___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QSqlDriver___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSqlDriver___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QSqlDriver___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QSqlDriver___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QSqlDriver___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QSqlDriver___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QSqlDriver___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSqlDriver___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QSqlDriver___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QSqlDriver___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSqlDriver___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void QSqlDriver_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QSqlDriver*>(ptr)->QSqlDriver::childEvent(static_cast<QChildEvent*>(event));
}

void QSqlDriver_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QSqlDriver*>(ptr)->QSqlDriver::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSqlDriver_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QSqlDriver*>(ptr)->QSqlDriver::customEvent(static_cast<QEvent*>(event));
}

void QSqlDriver_DeleteLaterDefault(void* ptr)
{
		static_cast<QSqlDriver*>(ptr)->QSqlDriver::deleteLater();
}

void QSqlDriver_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QSqlDriver*>(ptr)->QSqlDriver::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QSqlDriver_EventDefault(void* ptr, void* e)
{
		return static_cast<QSqlDriver*>(ptr)->QSqlDriver::event(static_cast<QEvent*>(e));
}

char QSqlDriver_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QSqlDriver*>(ptr)->QSqlDriver::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QSqlDriver_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QSqlDriver*>(ptr)->QSqlDriver::metaObject());
}

void QSqlDriver_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QSqlDriver*>(ptr)->QSqlDriver::timerEvent(static_cast<QTimerEvent*>(event));
}

class MyQSqlDriverCreatorBase: public QSqlDriverCreatorBase
{
public:
	QSqlDriver * createObject() const { return static_cast<QSqlDriver*>(callbackQSqlDriverCreatorBase_CreateObject(const_cast<void*>(static_cast<const void*>(this)))); };
	 ~MyQSqlDriverCreatorBase() { callbackQSqlDriverCreatorBase_DestroyQSqlDriverCreatorBase(this); };
};

Q_DECLARE_METATYPE(QSqlDriverCreatorBase*)
Q_DECLARE_METATYPE(MyQSqlDriverCreatorBase*)

int QSqlDriverCreatorBase_QSqlDriverCreatorBase_QRegisterMetaType(){qRegisterMetaType<QSqlDriverCreatorBase*>(); return qRegisterMetaType<MyQSqlDriverCreatorBase*>();}

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
	Q_UNUSED(ptr);

}

class MyQSqlDriverPlugin: public QSqlDriverPlugin
{
public:
	MyQSqlDriverPlugin(QObject *parent = Q_NULLPTR) : QSqlDriverPlugin(parent) {QSqlDriverPlugin_QSqlDriverPlugin_QRegisterMetaType();};
	QSqlDriver * create(const QString & key) { QByteArray* ta62f22 = new QByteArray(key.toUtf8()); QtSql_PackedString keyPacked = { const_cast<char*>(ta62f22->prepend("WHITESPACE").constData()+10), ta62f22->size()-10, ta62f22 };return static_cast<QSqlDriver*>(callbackQSqlDriverPlugin_Create(this, keyPacked)); };
	 ~MyQSqlDriverPlugin() { callbackQSqlDriverPlugin_DestroyQSqlDriverPlugin(this); };
	void childEvent(QChildEvent * event) { callbackQSqlDriverPlugin_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSqlDriverPlugin_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQSqlDriverPlugin_CustomEvent(this, event); };
	void deleteLater() { callbackQSqlDriverPlugin_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQSqlDriverPlugin_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSqlDriverPlugin_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQSqlDriverPlugin_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSqlDriverPlugin_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSqlDriverPlugin_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtSql_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQSqlDriverPlugin_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQSqlDriverPlugin_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QSqlDriverPlugin*)
Q_DECLARE_METATYPE(MyQSqlDriverPlugin*)

int QSqlDriverPlugin_QSqlDriverPlugin_QRegisterMetaType(){qRegisterMetaType<QSqlDriverPlugin*>(); return qRegisterMetaType<MyQSqlDriverPlugin*>();}

void* QSqlDriverPlugin_NewQSqlDriverPlugin(void* parent)
{
	if (dynamic_cast<QAudioSystemPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQSqlDriverPlugin(static_cast<QAudioSystemPlugin*>(parent));
	} else if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQSqlDriverPlugin(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQSqlDriverPlugin(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQSqlDriverPlugin(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQSqlDriverPlugin(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQSqlDriverPlugin(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQSqlDriverPlugin(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQSqlDriverPlugin(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQSqlDriverPlugin(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQSqlDriverPlugin(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QMediaServiceProviderPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQSqlDriverPlugin(static_cast<QMediaServiceProviderPlugin*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQSqlDriverPlugin(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQSqlDriverPlugin(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQSqlDriverPlugin(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQSqlDriverPlugin(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQSqlDriverPlugin(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQSqlDriverPlugin(static_cast<QRemoteObjectPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QScriptExtensionPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQSqlDriverPlugin(static_cast<QScriptExtensionPlugin*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQSqlDriverPlugin(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQSqlDriverPlugin(static_cast<QWindow*>(parent));
	} else {
		return new MyQSqlDriverPlugin(static_cast<QObject*>(parent));
	}
}

void* QSqlDriverPlugin_Create(void* ptr, struct QtSql_PackedString key)
{
	return static_cast<QSqlDriverPlugin*>(ptr)->create(QString::fromUtf8(key.data, key.len));
}

void QSqlDriverPlugin_DestroyQSqlDriverPlugin(void* ptr)
{
	static_cast<QSqlDriverPlugin*>(ptr)->~QSqlDriverPlugin();
}

void QSqlDriverPlugin_DestroyQSqlDriverPluginDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QSqlDriverPlugin___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QSqlDriverPlugin___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSqlDriverPlugin___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QSqlDriverPlugin___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QSqlDriverPlugin___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QSqlDriverPlugin___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QSqlDriverPlugin___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QSqlDriverPlugin___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSqlDriverPlugin___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QSqlDriverPlugin___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QSqlDriverPlugin___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSqlDriverPlugin___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void QSqlDriverPlugin_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QSqlDriverPlugin*>(ptr)->QSqlDriverPlugin::childEvent(static_cast<QChildEvent*>(event));
}

void QSqlDriverPlugin_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QSqlDriverPlugin*>(ptr)->QSqlDriverPlugin::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSqlDriverPlugin_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QSqlDriverPlugin*>(ptr)->QSqlDriverPlugin::customEvent(static_cast<QEvent*>(event));
}

void QSqlDriverPlugin_DeleteLaterDefault(void* ptr)
{
		static_cast<QSqlDriverPlugin*>(ptr)->QSqlDriverPlugin::deleteLater();
}

void QSqlDriverPlugin_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QSqlDriverPlugin*>(ptr)->QSqlDriverPlugin::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QSqlDriverPlugin_EventDefault(void* ptr, void* e)
{
		return static_cast<QSqlDriverPlugin*>(ptr)->QSqlDriverPlugin::event(static_cast<QEvent*>(e));
}

char QSqlDriverPlugin_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QSqlDriverPlugin*>(ptr)->QSqlDriverPlugin::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QSqlDriverPlugin_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QSqlDriverPlugin*>(ptr)->QSqlDriverPlugin::metaObject());
}

void QSqlDriverPlugin_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QSqlDriverPlugin*>(ptr)->QSqlDriverPlugin::timerEvent(static_cast<QTimerEvent*>(event));
}

Q_DECLARE_METATYPE(QSqlError*)
void* QSqlError_NewQSqlError2(struct QtSql_PackedString driverText, struct QtSql_PackedString databaseText, long long ty, struct QtSql_PackedString code)
{
	return new QSqlError(QString::fromUtf8(driverText.data, driverText.len), QString::fromUtf8(databaseText.data, databaseText.len), static_cast<QSqlError::ErrorType>(ty), QString::fromUtf8(code.data, code.len));
}

void* QSqlError_NewQSqlError3(void* other)
{
	return new QSqlError(*static_cast<QSqlError*>(other));
}

void* QSqlError_NewQSqlError4(void* other)
{
	return new QSqlError(*static_cast<QSqlError*>(other));
}

struct QtSql_PackedString QSqlError_DatabaseText(void* ptr)
{
	return ({ QByteArray* t7edaa1 = new QByteArray(static_cast<QSqlError*>(ptr)->databaseText().toUtf8()); QtSql_PackedString { const_cast<char*>(t7edaa1->prepend("WHITESPACE").constData()+10), t7edaa1->size()-10, t7edaa1 }; });
}

struct QtSql_PackedString QSqlError_DriverText(void* ptr)
{
	return ({ QByteArray* tf41ece = new QByteArray(static_cast<QSqlError*>(ptr)->driverText().toUtf8()); QtSql_PackedString { const_cast<char*>(tf41ece->prepend("WHITESPACE").constData()+10), tf41ece->size()-10, tf41ece }; });
}

char QSqlError_IsValid(void* ptr)
{
	return static_cast<QSqlError*>(ptr)->isValid();
}

struct QtSql_PackedString QSqlError_NativeErrorCode(void* ptr)
{
	return ({ QByteArray* t468473 = new QByteArray(static_cast<QSqlError*>(ptr)->nativeErrorCode().toUtf8()); QtSql_PackedString { const_cast<char*>(t468473->prepend("WHITESPACE").constData()+10), t468473->size()-10, t468473 }; });
}

void QSqlError_Swap(void* ptr, void* other)
{
	static_cast<QSqlError*>(ptr)->swap(*static_cast<QSqlError*>(other));
}

struct QtSql_PackedString QSqlError_Text(void* ptr)
{
	return ({ QByteArray* t22e689 = new QByteArray(static_cast<QSqlError*>(ptr)->text().toUtf8()); QtSql_PackedString { const_cast<char*>(t22e689->prepend("WHITESPACE").constData()+10), t22e689->size()-10, t22e689 }; });
}

long long QSqlError_Type(void* ptr)
{
	return static_cast<QSqlError*>(ptr)->type();
}

void QSqlError_DestroyQSqlError(void* ptr)
{
	static_cast<QSqlError*>(ptr)->~QSqlError();
}

Q_DECLARE_METATYPE(QSqlField*)
void* QSqlField_NewQSqlField(struct QtSql_PackedString fieldName, long long ty)
{
	return new QSqlField(QString::fromUtf8(fieldName.data, fieldName.len), static_cast<QVariant::Type>(ty));
}

void* QSqlField_NewQSqlField2(struct QtSql_PackedString fieldName, long long ty, struct QtSql_PackedString table)
{
	return new QSqlField(QString::fromUtf8(fieldName.data, fieldName.len), static_cast<QVariant::Type>(ty), QString::fromUtf8(table.data, table.len));
}

void* QSqlField_NewQSqlField3(void* other)
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
	return ({ QByteArray* t784037 = new QByteArray(static_cast<QSqlField*>(ptr)->name().toUtf8()); QtSql_PackedString { const_cast<char*>(t784037->prepend("WHITESPACE").constData()+10), t784037->size()-10, t784037 }; });
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

void QSqlField_SetName(void* ptr, struct QtSql_PackedString name)
{
	static_cast<QSqlField*>(ptr)->setName(QString::fromUtf8(name.data, name.len));
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

void QSqlField_SetTableName(void* ptr, struct QtSql_PackedString table)
{
	static_cast<QSqlField*>(ptr)->setTableName(QString::fromUtf8(table.data, table.len));
}

void QSqlField_SetType(void* ptr, long long ty)
{
	static_cast<QSqlField*>(ptr)->setType(static_cast<QVariant::Type>(ty));
}

void QSqlField_SetValue(void* ptr, void* value)
{
	static_cast<QSqlField*>(ptr)->setValue(*static_cast<QVariant*>(value));
}

struct QtSql_PackedString QSqlField_TableName(void* ptr)
{
	return ({ QByteArray* tbfea2f = new QByteArray(static_cast<QSqlField*>(ptr)->tableName().toUtf8()); QtSql_PackedString { const_cast<char*>(tbfea2f->prepend("WHITESPACE").constData()+10), tbfea2f->size()-10, tbfea2f }; });
}

long long QSqlField_Type(void* ptr)
{
	return static_cast<QSqlField*>(ptr)->type();
}

void* QSqlField_Value(void* ptr)
{
	return new QVariant(static_cast<QSqlField*>(ptr)->value());
}

void QSqlField_DestroyQSqlField(void* ptr)
{
	static_cast<QSqlField*>(ptr)->~QSqlField();
}

Q_DECLARE_METATYPE(QSqlIndex*)
void* QSqlIndex_NewQSqlIndex(struct QtSql_PackedString cursorname, struct QtSql_PackedString name)
{
	return new QSqlIndex(QString::fromUtf8(cursorname.data, cursorname.len), QString::fromUtf8(name.data, name.len));
}

void* QSqlIndex_NewQSqlIndex2(void* other)
{
	return new QSqlIndex(*static_cast<QSqlIndex*>(other));
}

void QSqlIndex_Append2(void* ptr, void* field, char desc)
{
	static_cast<QSqlIndex*>(ptr)->append(*static_cast<QSqlField*>(field), desc != 0);
}

struct QtSql_PackedString QSqlIndex_CursorName(void* ptr)
{
	return ({ QByteArray* t1764a5 = new QByteArray(static_cast<QSqlIndex*>(ptr)->cursorName().toUtf8()); QtSql_PackedString { const_cast<char*>(t1764a5->prepend("WHITESPACE").constData()+10), t1764a5->size()-10, t1764a5 }; });
}

char QSqlIndex_IsDescending(void* ptr, int i)
{
	return static_cast<QSqlIndex*>(ptr)->isDescending(i);
}

struct QtSql_PackedString QSqlIndex_Name(void* ptr)
{
	return ({ QByteArray* tdb3ffc = new QByteArray(static_cast<QSqlIndex*>(ptr)->name().toUtf8()); QtSql_PackedString { const_cast<char*>(tdb3ffc->prepend("WHITESPACE").constData()+10), tdb3ffc->size()-10, tdb3ffc }; });
}

void QSqlIndex_SetCursorName(void* ptr, struct QtSql_PackedString cursorName)
{
	static_cast<QSqlIndex*>(ptr)->setCursorName(QString::fromUtf8(cursorName.data, cursorName.len));
}

void QSqlIndex_SetDescending(void* ptr, int i, char desc)
{
	static_cast<QSqlIndex*>(ptr)->setDescending(i, desc != 0);
}

void QSqlIndex_SetName(void* ptr, struct QtSql_PackedString name)
{
	static_cast<QSqlIndex*>(ptr)->setName(QString::fromUtf8(name.data, name.len));
}

void QSqlIndex_DestroyQSqlIndex(void* ptr)
{
	static_cast<QSqlIndex*>(ptr)->~QSqlIndex();
}

char QSqlIndex___sorts_atList(void* ptr, int i)
{
	return ({bool tmp = static_cast<QVector<bool>*>(ptr)->at(i); if (i == static_cast<QVector<bool>*>(ptr)->size()-1) { static_cast<QVector<bool>*>(ptr)->~QVector(); free(ptr); }; tmp; });
}

void QSqlIndex___sorts_setList(void* ptr, char i)
{
	static_cast<QVector<bool>*>(ptr)->append(i != 0);
}

void* QSqlIndex___sorts_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<bool>();
}

char QSqlIndex___setSorts__atList(void* ptr, int i)
{
	return ({bool tmp = static_cast<QVector<bool>*>(ptr)->at(i); if (i == static_cast<QVector<bool>*>(ptr)->size()-1) { static_cast<QVector<bool>*>(ptr)->~QVector(); free(ptr); }; tmp; });
}

void QSqlIndex___setSorts__setList(void* ptr, char i)
{
	static_cast<QVector<bool>*>(ptr)->append(i != 0);
}

void* QSqlIndex___setSorts__newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<bool>();
}

Q_DECLARE_METATYPE(QSqlQuery*)
void* QSqlQuery_NewQSqlQuery(void* result)
{
	return new QSqlQuery(static_cast<QSqlResult*>(result));
}

void* QSqlQuery_NewQSqlQuery2(struct QtSql_PackedString query, void* db)
{
	return new QSqlQuery(QString::fromUtf8(query.data, query.len), *static_cast<QSqlDatabase*>(db));
}

void* QSqlQuery_NewQSqlQuery3(void* db)
{
	return new QSqlQuery(*static_cast<QSqlDatabase*>(db));
}

void* QSqlQuery_NewQSqlQuery4(void* other)
{
	return new QSqlQuery(*static_cast<QSqlQuery*>(other));
}

void QSqlQuery_AddBindValue(void* ptr, void* val, long long paramType)
{
	static_cast<QSqlQuery*>(ptr)->addBindValue(*static_cast<QVariant*>(val), static_cast<QSql::ParamTypeFlag>(paramType));
}

int QSqlQuery_At(void* ptr)
{
	return static_cast<QSqlQuery*>(ptr)->at();
}

void QSqlQuery_BindValue(void* ptr, struct QtSql_PackedString placeholder, void* val, long long paramType)
{
	static_cast<QSqlQuery*>(ptr)->bindValue(QString::fromUtf8(placeholder.data, placeholder.len), *static_cast<QVariant*>(val), static_cast<QSql::ParamTypeFlag>(paramType));
}

void QSqlQuery_BindValue2(void* ptr, int pos, void* val, long long paramType)
{
	static_cast<QSqlQuery*>(ptr)->bindValue(pos, *static_cast<QVariant*>(val), static_cast<QSql::ParamTypeFlag>(paramType));
}

void* QSqlQuery_BoundValue(void* ptr, struct QtSql_PackedString placeholder)
{
	return new QVariant(static_cast<QSqlQuery*>(ptr)->boundValue(QString::fromUtf8(placeholder.data, placeholder.len)));
}

void* QSqlQuery_BoundValue2(void* ptr, int pos)
{
	return new QVariant(static_cast<QSqlQuery*>(ptr)->boundValue(pos));
}

struct QtSql_PackedList QSqlQuery_BoundValues(void* ptr)
{
	return ({ QMap<QString, QVariant>* tmpValue76e935 = new QMap<QString, QVariant>(static_cast<QSqlQuery*>(ptr)->boundValues()); QtSql_PackedList { tmpValue76e935, tmpValue76e935->size() }; });
}

void QSqlQuery_Clear(void* ptr)
{
	static_cast<QSqlQuery*>(ptr)->clear();
}

void* QSqlQuery_Driver(void* ptr)
{
	return const_cast<QSqlDriver*>(static_cast<QSqlQuery*>(ptr)->driver());
}

char QSqlQuery_Exec(void* ptr, struct QtSql_PackedString query)
{
	return static_cast<QSqlQuery*>(ptr)->exec(QString::fromUtf8(query.data, query.len));
}

char QSqlQuery_Exec2(void* ptr)
{
	return static_cast<QSqlQuery*>(ptr)->exec();
}

char QSqlQuery_ExecBatch(void* ptr, long long mode)
{
	return static_cast<QSqlQuery*>(ptr)->execBatch(static_cast<QSqlQuery::BatchExecutionMode>(mode));
}

struct QtSql_PackedString QSqlQuery_ExecutedQuery(void* ptr)
{
	return ({ QByteArray* tc2ded8 = new QByteArray(static_cast<QSqlQuery*>(ptr)->executedQuery().toUtf8()); QtSql_PackedString { const_cast<char*>(tc2ded8->prepend("WHITESPACE").constData()+10), tc2ded8->size()-10, tc2ded8 }; });
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

char QSqlQuery_IsNull(void* ptr, int field)
{
	return static_cast<QSqlQuery*>(ptr)->isNull(field);
}

char QSqlQuery_IsNull2(void* ptr, struct QtSql_PackedString name)
{
	return static_cast<QSqlQuery*>(ptr)->isNull(QString::fromUtf8(name.data, name.len));
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
	return ({ QByteArray* t0d33b2 = new QByteArray(static_cast<QSqlQuery*>(ptr)->lastQuery().toUtf8()); QtSql_PackedString { const_cast<char*>(t0d33b2->prepend("WHITESPACE").constData()+10), t0d33b2->size()-10, t0d33b2 }; });
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

char QSqlQuery_Prepare(void* ptr, struct QtSql_PackedString query)
{
	return static_cast<QSqlQuery*>(ptr)->prepare(QString::fromUtf8(query.data, query.len));
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

void* QSqlQuery_Value(void* ptr, int index)
{
	return new QVariant(static_cast<QSqlQuery*>(ptr)->value(index));
}

void* QSqlQuery_Value2(void* ptr, struct QtSql_PackedString name)
{
	return new QVariant(static_cast<QSqlQuery*>(ptr)->value(QString::fromUtf8(name.data, name.len)));
}

void QSqlQuery_DestroyQSqlQuery(void* ptr)
{
	static_cast<QSqlQuery*>(ptr)->~QSqlQuery();
}

void* QSqlQuery___boundValues_atList(void* ptr, struct QtSql_PackedString v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<QString, QVariant>*>(ptr)->value(QString::fromUtf8(v.data, v.len)); if (i == static_cast<QMap<QString, QVariant>*>(ptr)->size()-1) { static_cast<QMap<QString, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void QSqlQuery___boundValues_setList(void* ptr, struct QtSql_PackedString key, void* i)
{
	static_cast<QMap<QString, QVariant>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(i));
}

void* QSqlQuery___boundValues_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<QString, QVariant>();
}

struct QtSql_PackedList QSqlQuery___boundValues_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue1ab909 = new QList<QString>(static_cast<QMap<QString, QVariant>*>(ptr)->keys()); QtSql_PackedList { tmpValue1ab909, tmpValue1ab909->size() }; });
}

struct QtSql_PackedString QSqlQuery_____boundValues_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray* t94aa5e = new QByteArray(({QString tmp = static_cast<QList<QString>*>(ptr)->at(i); if (i == static_cast<QList<QString>*>(ptr)->size()-1) { static_cast<QList<QString>*>(ptr)->~QList(); free(ptr); }; tmp; }).toUtf8()); QtSql_PackedString { const_cast<char*>(t94aa5e->prepend("WHITESPACE").constData()+10), t94aa5e->size()-10, t94aa5e }; });
}

void QSqlQuery_____boundValues_keyList_setList(void* ptr, struct QtSql_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QSqlQuery_____boundValues_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>();
}

class MyQSqlQueryModel: public QSqlQueryModel
{
public:
	MyQSqlQueryModel(QObject *parent = Q_NULLPTR) : QSqlQueryModel(parent) {QSqlQueryModel_QSqlQueryModel_QRegisterMetaType();};
	bool canFetchMore(const QModelIndex & parent) const { return callbackQSqlQueryModel_CanFetchMore(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	void clear() { callbackQSqlQueryModel_Clear(this); };
	int columnCount(const QModelIndex & index) const { return callbackQSqlQueryModel_ColumnCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index)); };
	QVariant data(const QModelIndex & item, int role) const { return *static_cast<QVariant*>(callbackQSqlQueryModel_Data(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&item), role)); };
	void fetchMore(const QModelIndex & parent) { callbackQSqlQueryModel_FetchMore(this, const_cast<QModelIndex*>(&parent)); };
	QVariant headerData(int section, Qt::Orientation orientation, int role) const { return *static_cast<QVariant*>(callbackQSqlQueryModel_HeaderData(const_cast<void*>(static_cast<const void*>(this)), section, orientation, role)); };
	QModelIndex indexInQuery(const QModelIndex & item) const { return *static_cast<QModelIndex*>(callbackQSqlQueryModel_IndexInQuery(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&item))); };
	bool insertColumns(int column, int count, const QModelIndex & parent) { return callbackQSqlQueryModel_InsertColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	void queryChange() { callbackQSqlQueryModel_QueryChange(this); };
	bool removeColumns(int column, int count, const QModelIndex & parent) { return callbackQSqlQueryModel_RemoveColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	QHash<int, QByteArray> roleNames() const { return ({ QHash<int, QByteArray>* tmpP = static_cast<QHash<int, QByteArray>*>(callbackQSqlQueryModel_RoleNames(const_cast<void*>(static_cast<const void*>(this)))); QHash<int, QByteArray> tmpV = *tmpP; tmpP->~QHash(); free(tmpP); tmpV; }); };
	int rowCount(const QModelIndex & parent) const { return callbackQSqlQueryModel_RowCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	bool setHeaderData(int section, Qt::Orientation orientation, const QVariant & value, int role) { return callbackQSqlQueryModel_SetHeaderData(this, section, orientation, const_cast<QVariant*>(&value), role) != 0; };
	 ~MyQSqlQueryModel() { callbackQSqlQueryModel_DestroyQSqlQueryModel(this); };
	bool dropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) { return callbackQSqlQueryModel_DropMimeData(this, const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	Qt::ItemFlags flags(const QModelIndex & index) const { return static_cast<Qt::ItemFlag>(callbackQSqlQueryModel_Flags(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QModelIndex index(int row, int column, const QModelIndex & parent) const { return *static_cast<QModelIndex*>(callbackQSqlQueryModel_Index(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&parent))); };
	QModelIndex sibling(int row, int column, const QModelIndex & idx) const { return *static_cast<QModelIndex*>(callbackQSqlQueryModel_Sibling(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&idx))); };
	QModelIndex buddy(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackQSqlQueryModel_Buddy(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool canDropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) const { return callbackQSqlQueryModel_CanDropMimeData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	void Signal_ColumnsAboutToBeInserted(const QModelIndex & parent, int first, int last) { callbackQSqlQueryModel_ColumnsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationColumn) { callbackQSqlQueryModel_ColumnsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationColumn); };
	void Signal_ColumnsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackQSqlQueryModel_ColumnsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsInserted(const QModelIndex & parent, int first, int last) { callbackQSqlQueryModel_ColumnsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int column) { callbackQSqlQueryModel_ColumnsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), column); };
	void Signal_ColumnsRemoved(const QModelIndex & parent, int first, int last) { callbackQSqlQueryModel_ColumnsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_DataChanged(const QModelIndex & topLeft, const QModelIndex & bottomRight, const QVector<int> & roles) { callbackQSqlQueryModel_DataChanged(this, const_cast<QModelIndex*>(&topLeft), const_cast<QModelIndex*>(&bottomRight), ({ QVector<int>* tmpValue037c88 = new QVector<int>(roles); QtSql_PackedList { tmpValue037c88, tmpValue037c88->size() }; })); };
	bool hasChildren(const QModelIndex & parent) const { return callbackQSqlQueryModel_HasChildren(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	void Signal_HeaderDataChanged(Qt::Orientation orientation, int first, int last) { callbackQSqlQueryModel_HeaderDataChanged(this, orientation, first, last); };
	bool insertRows(int row, int count, const QModelIndex & parent) { return callbackQSqlQueryModel_InsertRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	QMap<int, QVariant> itemData(const QModelIndex & index) const { return ({ QMap<int, QVariant>* tmpP = static_cast<QMap<int, QVariant>*>(callbackQSqlQueryModel_ItemData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); QMap<int, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }); };
	void Signal_LayoutAboutToBeChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackQSqlQueryModel_LayoutAboutToBeChanged(this, ({ QList<QPersistentModelIndex>* tmpValuea664f1 = new QList<QPersistentModelIndex>(parents); QtSql_PackedList { tmpValuea664f1, tmpValuea664f1->size() }; }), hint); };
	void Signal_LayoutChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackQSqlQueryModel_LayoutChanged(this, ({ QList<QPersistentModelIndex>* tmpValuea664f1 = new QList<QPersistentModelIndex>(parents); QtSql_PackedList { tmpValuea664f1, tmpValuea664f1->size() }; }), hint); };
	QList<QModelIndex> match(const QModelIndex & start, int role, const QVariant & value, int hits, Qt::MatchFlags flags) const { return ({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(callbackQSqlQueryModel_Match(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&start), role, const_cast<QVariant*>(&value), hits, flags)); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	QMimeData * mimeData(const QModelIndexList & indexes) const { return static_cast<QMimeData*>(callbackQSqlQueryModel_MimeData(const_cast<void*>(static_cast<const void*>(this)), ({ QList<QModelIndex>* tmpValuee0adf2 = new QList<QModelIndex>(indexes); QtSql_PackedList { tmpValuee0adf2, tmpValuee0adf2->size() }; }))); };
	QStringList mimeTypes() const { return ({ QtSql_PackedString tempVal = callbackQSqlQueryModel_MimeTypes(const_cast<void*>(static_cast<const void*>(this))); QStringList ret = QString::fromUtf8(tempVal.data, tempVal.len).split("¡¦!", QString::SkipEmptyParts); free(tempVal.data); ret; }); };
	void Signal_ModelAboutToBeReset() { callbackQSqlQueryModel_ModelAboutToBeReset(this); };
	void Signal_ModelReset() { callbackQSqlQueryModel_ModelReset(this); };
	bool moveColumns(const QModelIndex & sourceParent, int sourceColumn, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackQSqlQueryModel_MoveColumns(this, const_cast<QModelIndex*>(&sourceParent), sourceColumn, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	bool moveRows(const QModelIndex & sourceParent, int sourceRow, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackQSqlQueryModel_MoveRows(this, const_cast<QModelIndex*>(&sourceParent), sourceRow, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	QModelIndex parent(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackQSqlQueryModel_Parent(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool removeRows(int row, int count, const QModelIndex & parent) { return callbackQSqlQueryModel_RemoveRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	void resetInternalData() { callbackQSqlQueryModel_ResetInternalData(this); };
	void revert() { callbackQSqlQueryModel_Revert(this); };
	void Signal_RowsAboutToBeInserted(const QModelIndex & parent, int start, int end) { callbackQSqlQueryModel_RowsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), start, end); };
	void Signal_RowsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationRow) { callbackQSqlQueryModel_RowsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationRow); };
	void Signal_RowsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackQSqlQueryModel_RowsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsInserted(const QModelIndex & parent, int first, int last) { callbackQSqlQueryModel_RowsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int row) { callbackQSqlQueryModel_RowsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), row); };
	void Signal_RowsRemoved(const QModelIndex & parent, int first, int last) { callbackQSqlQueryModel_RowsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	bool setData(const QModelIndex & index, const QVariant & value, int role) { return callbackQSqlQueryModel_SetData(this, const_cast<QModelIndex*>(&index), const_cast<QVariant*>(&value), role) != 0; };
	bool setItemData(const QModelIndex & index, const QMap<int, QVariant> & roles) { return callbackQSqlQueryModel_SetItemData(this, const_cast<QModelIndex*>(&index), ({ QMap<int, QVariant>* tmpValue037c88 = new QMap<int, QVariant>(roles); QtSql_PackedList { tmpValue037c88, tmpValue037c88->size() }; })) != 0; };
	void sort(int column, Qt::SortOrder order) { callbackQSqlQueryModel_Sort(this, column, order); };
	QSize span(const QModelIndex & index) const { return *static_cast<QSize*>(callbackQSqlQueryModel_Span(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool submit() { return callbackQSqlQueryModel_Submit(this) != 0; };
	Qt::DropActions supportedDragActions() const { return static_cast<Qt::DropAction>(callbackQSqlQueryModel_SupportedDragActions(const_cast<void*>(static_cast<const void*>(this)))); };
	Qt::DropActions supportedDropActions() const { return static_cast<Qt::DropAction>(callbackQSqlQueryModel_SupportedDropActions(const_cast<void*>(static_cast<const void*>(this)))); };
	void childEvent(QChildEvent * event) { callbackQSqlQueryModel_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSqlQueryModel_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQSqlQueryModel_CustomEvent(this, event); };
	void deleteLater() { callbackQSqlQueryModel_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQSqlQueryModel_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSqlQueryModel_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQSqlQueryModel_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSqlQueryModel_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSqlQueryModel_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtSql_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQSqlQueryModel_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQSqlQueryModel_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QSqlQueryModel*)
Q_DECLARE_METATYPE(MyQSqlQueryModel*)

int QSqlQueryModel_QSqlQueryModel_QRegisterMetaType(){qRegisterMetaType<QSqlQueryModel*>(); return qRegisterMetaType<MyQSqlQueryModel*>();}

void* QSqlQueryModel_NewQSqlQueryModel(void* parent)
{
	if (dynamic_cast<QAudioSystemPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQSqlQueryModel(static_cast<QAudioSystemPlugin*>(parent));
	} else if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQSqlQueryModel(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQSqlQueryModel(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQSqlQueryModel(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQSqlQueryModel(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQSqlQueryModel(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQSqlQueryModel(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQSqlQueryModel(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQSqlQueryModel(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQSqlQueryModel(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QMediaServiceProviderPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQSqlQueryModel(static_cast<QMediaServiceProviderPlugin*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQSqlQueryModel(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQSqlQueryModel(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQSqlQueryModel(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQSqlQueryModel(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQSqlQueryModel(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQSqlQueryModel(static_cast<QRemoteObjectPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QScriptExtensionPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQSqlQueryModel(static_cast<QScriptExtensionPlugin*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQSqlQueryModel(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQSqlQueryModel(static_cast<QWindow*>(parent));
	} else {
		return new MyQSqlQueryModel(static_cast<QObject*>(parent));
	}
}

char QSqlQueryModel_CanFetchMoreDefault(void* ptr, void* parent)
{
	if (dynamic_cast<QSqlRelationalTableModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::canFetchMore(*static_cast<QModelIndex*>(parent));
	} else if (dynamic_cast<QSqlTableModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::canFetchMore(*static_cast<QModelIndex*>(parent));
	} else {
		return static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::canFetchMore(*static_cast<QModelIndex*>(parent));
	}
}

void QSqlQueryModel_Clear(void* ptr)
{
	static_cast<QSqlQueryModel*>(ptr)->clear();
}

void QSqlQueryModel_ClearDefault(void* ptr)
{
	if (dynamic_cast<QSqlRelationalTableModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::clear();
	} else if (dynamic_cast<QSqlTableModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::clear();
	} else {
		static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::clear();
	}
}

int QSqlQueryModel_ColumnCount(void* ptr, void* index)
{
	return static_cast<QSqlQueryModel*>(ptr)->columnCount(*static_cast<QModelIndex*>(index));
}

int QSqlQueryModel_ColumnCountDefault(void* ptr, void* index)
{
	if (dynamic_cast<QSqlRelationalTableModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::columnCount(*static_cast<QModelIndex*>(index));
	} else if (dynamic_cast<QSqlTableModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::columnCount(*static_cast<QModelIndex*>(index));
	} else {
		return static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::columnCount(*static_cast<QModelIndex*>(index));
	}
}

void* QSqlQueryModel_Data(void* ptr, void* item, int role)
{
	return new QVariant(static_cast<QSqlQueryModel*>(ptr)->data(*static_cast<QModelIndex*>(item), role));
}

void* QSqlQueryModel_DataDefault(void* ptr, void* item, int role)
{
	if (dynamic_cast<QSqlRelationalTableModel*>(static_cast<QObject*>(ptr))) {
		return new QVariant(static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::data(*static_cast<QModelIndex*>(item), role));
	} else if (dynamic_cast<QSqlTableModel*>(static_cast<QObject*>(ptr))) {
		return new QVariant(static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::data(*static_cast<QModelIndex*>(item), role));
	} else {
		return new QVariant(static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::data(*static_cast<QModelIndex*>(item), role));
	}
}

void QSqlQueryModel_FetchMoreDefault(void* ptr, void* parent)
{
	if (dynamic_cast<QSqlRelationalTableModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::fetchMore(*static_cast<QModelIndex*>(parent));
	} else if (dynamic_cast<QSqlTableModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::fetchMore(*static_cast<QModelIndex*>(parent));
	} else {
		static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::fetchMore(*static_cast<QModelIndex*>(parent));
	}
}

void* QSqlQueryModel_HeaderDataDefault(void* ptr, int section, long long orientation, int role)
{
	if (dynamic_cast<QSqlRelationalTableModel*>(static_cast<QObject*>(ptr))) {
		return new QVariant(static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::headerData(section, static_cast<Qt::Orientation>(orientation), role));
	} else if (dynamic_cast<QSqlTableModel*>(static_cast<QObject*>(ptr))) {
		return new QVariant(static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::headerData(section, static_cast<Qt::Orientation>(orientation), role));
	} else {
		return new QVariant(static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::headerData(section, static_cast<Qt::Orientation>(orientation), role));
	}
}

void* QSqlQueryModel_IndexInQuery(void* ptr, void* item)
{
	return new QModelIndex(static_cast<QSqlQueryModel*>(ptr)->indexInQuery(*static_cast<QModelIndex*>(item)));
}

void* QSqlQueryModel_IndexInQueryDefault(void* ptr, void* item)
{
	if (dynamic_cast<QSqlRelationalTableModel*>(static_cast<QObject*>(ptr))) {
		return new QModelIndex(static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::indexInQuery(*static_cast<QModelIndex*>(item)));
	} else if (dynamic_cast<QSqlTableModel*>(static_cast<QObject*>(ptr))) {
		return new QModelIndex(static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::indexInQuery(*static_cast<QModelIndex*>(item)));
	} else {
		return new QModelIndex(static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::indexInQuery(*static_cast<QModelIndex*>(item)));
	}
}

char QSqlQueryModel_InsertColumnsDefault(void* ptr, int column, int count, void* parent)
{
	if (dynamic_cast<QSqlRelationalTableModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::insertColumns(column, count, *static_cast<QModelIndex*>(parent));
	} else if (dynamic_cast<QSqlTableModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::insertColumns(column, count, *static_cast<QModelIndex*>(parent));
	} else {
		return static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::insertColumns(column, count, *static_cast<QModelIndex*>(parent));
	}
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
	if (dynamic_cast<QSqlRelationalTableModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::queryChange();
	} else if (dynamic_cast<QSqlTableModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::queryChange();
	} else {
		static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::queryChange();
	}
}

void* QSqlQueryModel_Record(void* ptr, int row)
{
	return new QSqlRecord(static_cast<QSqlQueryModel*>(ptr)->record(row));
}

void* QSqlQueryModel_Record2(void* ptr)
{
	return new QSqlRecord(static_cast<QSqlQueryModel*>(ptr)->record());
}

char QSqlQueryModel_RemoveColumnsDefault(void* ptr, int column, int count, void* parent)
{
	if (dynamic_cast<QSqlRelationalTableModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::removeColumns(column, count, *static_cast<QModelIndex*>(parent));
	} else if (dynamic_cast<QSqlTableModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::removeColumns(column, count, *static_cast<QModelIndex*>(parent));
	} else {
		return static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::removeColumns(column, count, *static_cast<QModelIndex*>(parent));
	}
}

struct QtSql_PackedList QSqlQueryModel_RoleNamesDefault(void* ptr)
{
	if (dynamic_cast<QSqlRelationalTableModel*>(static_cast<QObject*>(ptr))) {
		return ({ QHash<int, QByteArray>* tmpValue7d28be = new QHash<int, QByteArray>(static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::roleNames()); QtSql_PackedList { tmpValue7d28be, tmpValue7d28be->size() }; });
	} else if (dynamic_cast<QSqlTableModel*>(static_cast<QObject*>(ptr))) {
		return ({ QHash<int, QByteArray>* tmpValue7d28be = new QHash<int, QByteArray>(static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::roleNames()); QtSql_PackedList { tmpValue7d28be, tmpValue7d28be->size() }; });
	} else {
		return ({ QHash<int, QByteArray>* tmpValue7d28be = new QHash<int, QByteArray>(static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::roleNames()); QtSql_PackedList { tmpValue7d28be, tmpValue7d28be->size() }; });
	}
}

int QSqlQueryModel_RowCount(void* ptr, void* parent)
{
	return static_cast<QSqlQueryModel*>(ptr)->rowCount(*static_cast<QModelIndex*>(parent));
}

int QSqlQueryModel_RowCountDefault(void* ptr, void* parent)
{
	if (dynamic_cast<QSqlRelationalTableModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::rowCount(*static_cast<QModelIndex*>(parent));
	} else if (dynamic_cast<QSqlTableModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::rowCount(*static_cast<QModelIndex*>(parent));
	} else {
		return static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::rowCount(*static_cast<QModelIndex*>(parent));
	}
}

char QSqlQueryModel_SetHeaderDataDefault(void* ptr, int section, long long orientation, void* value, int role)
{
	if (dynamic_cast<QSqlRelationalTableModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::setHeaderData(section, static_cast<Qt::Orientation>(orientation), *static_cast<QVariant*>(value), role);
	} else if (dynamic_cast<QSqlTableModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::setHeaderData(section, static_cast<Qt::Orientation>(orientation), *static_cast<QVariant*>(value), role);
	} else {
		return static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::setHeaderData(section, static_cast<Qt::Orientation>(orientation), *static_cast<QVariant*>(value), role);
	}
}

void QSqlQueryModel_SetLastError(void* ptr, void* error)
{
	static_cast<QSqlQueryModel*>(ptr)->setLastError(*static_cast<QSqlError*>(error));
}

void QSqlQueryModel_SetQuery(void* ptr, void* query)
{
	static_cast<QSqlQueryModel*>(ptr)->setQuery(*static_cast<QSqlQuery*>(query));
}

void QSqlQueryModel_SetQuery2(void* ptr, struct QtSql_PackedString query, void* db)
{
	static_cast<QSqlQueryModel*>(ptr)->setQuery(QString::fromUtf8(query.data, query.len), *static_cast<QSqlDatabase*>(db));
}

void QSqlQueryModel_DestroyQSqlQueryModel(void* ptr)
{
	static_cast<QSqlQueryModel*>(ptr)->~QSqlQueryModel();
}

void QSqlQueryModel_DestroyQSqlQueryModelDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QSqlQueryModel___roleNames_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QHash<int, QByteArray>*>(ptr)->value(v); if (i == static_cast<QHash<int, QByteArray>*>(ptr)->size()-1) { static_cast<QHash<int, QByteArray>*>(ptr)->~QHash(); free(ptr); }; tmp; }));
}

void QSqlQueryModel___roleNames_setList(void* ptr, int key, void* i)
{
	static_cast<QHash<int, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* QSqlQueryModel___roleNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<int, QByteArray>();
}

struct QtSql_PackedList QSqlQueryModel___roleNames_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue7fc3bb = new QList<int>(static_cast<QHash<int, QByteArray>*>(ptr)->keys()); QtSql_PackedList { tmpValue7fc3bb, tmpValue7fc3bb->size() }; });
}

int QSqlQueryModel_____roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QSqlQueryModel_____roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* QSqlQueryModel_____roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int QSqlQueryModel_____itemData_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QSqlQueryModel_____itemData_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* QSqlQueryModel_____itemData_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int QSqlQueryModel_____setItemData_roles_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QSqlQueryModel_____setItemData_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* QSqlQueryModel_____setItemData_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* QSqlQueryModel___changePersistentIndexList_from_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QSqlQueryModel___changePersistentIndexList_from_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* QSqlQueryModel___changePersistentIndexList_from_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* QSqlQueryModel___changePersistentIndexList_to_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QSqlQueryModel___changePersistentIndexList_to_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* QSqlQueryModel___changePersistentIndexList_to_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

int QSqlQueryModel___dataChanged_roles_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QVector<int>*>(ptr)->at(i); if (i == static_cast<QVector<int>*>(ptr)->size()-1) { static_cast<QVector<int>*>(ptr)->~QVector(); free(ptr); }; tmp; });
}

void QSqlQueryModel___dataChanged_roles_setList(void* ptr, int i)
{
	static_cast<QVector<int>*>(ptr)->append(i);
}

void* QSqlQueryModel___dataChanged_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<int>();
}

void* QSqlQueryModel___itemData_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void QSqlQueryModel___itemData_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* QSqlQueryModel___itemData_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct QtSql_PackedList QSqlQueryModel___itemData_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue249128 = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); QtSql_PackedList { tmpValue249128, tmpValue249128->size() }; });
}

void* QSqlQueryModel___layoutAboutToBeChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QSqlQueryModel___layoutAboutToBeChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* QSqlQueryModel___layoutAboutToBeChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* QSqlQueryModel___layoutChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QSqlQueryModel___layoutChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* QSqlQueryModel___layoutChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* QSqlQueryModel___match_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QSqlQueryModel___match_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* QSqlQueryModel___match_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* QSqlQueryModel___mimeData_indexes_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QSqlQueryModel___mimeData_indexes_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* QSqlQueryModel___mimeData_indexes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* QSqlQueryModel___persistentIndexList_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QSqlQueryModel___persistentIndexList_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* QSqlQueryModel___persistentIndexList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* QSqlQueryModel___setItemData_roles_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void QSqlQueryModel___setItemData_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* QSqlQueryModel___setItemData_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct QtSql_PackedList QSqlQueryModel___setItemData_roles_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue249128 = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); QtSql_PackedList { tmpValue249128, tmpValue249128->size() }; });
}

int QSqlQueryModel_____doSetRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QSqlQueryModel_____doSetRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* QSqlQueryModel_____doSetRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int QSqlQueryModel_____setRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QSqlQueryModel_____setRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* QSqlQueryModel_____setRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* QSqlQueryModel___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QSqlQueryModel___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSqlQueryModel___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QSqlQueryModel___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QSqlQueryModel___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QSqlQueryModel___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QSqlQueryModel___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QSqlQueryModel___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSqlQueryModel___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QSqlQueryModel___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QSqlQueryModel___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSqlQueryModel___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

char QSqlQueryModel_DropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	if (dynamic_cast<QSqlRelationalTableModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
	} else if (dynamic_cast<QSqlTableModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
	} else {
		return static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
	}
}

long long QSqlQueryModel_FlagsDefault(void* ptr, void* index)
{
	if (dynamic_cast<QSqlRelationalTableModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::flags(*static_cast<QModelIndex*>(index));
	} else if (dynamic_cast<QSqlTableModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::flags(*static_cast<QModelIndex*>(index));
	} else {
		return static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::flags(*static_cast<QModelIndex*>(index));
	}
}

void* QSqlQueryModel_IndexDefault(void* ptr, int row, int column, void* parent)
{
	if (dynamic_cast<QSqlRelationalTableModel*>(static_cast<QObject*>(ptr))) {
		return new QModelIndex(static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::index(row, column, *static_cast<QModelIndex*>(parent)));
	} else if (dynamic_cast<QSqlTableModel*>(static_cast<QObject*>(ptr))) {
		return new QModelIndex(static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::index(row, column, *static_cast<QModelIndex*>(parent)));
	} else {
		return new QModelIndex(static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::index(row, column, *static_cast<QModelIndex*>(parent)));
	}
}

void* QSqlQueryModel_SiblingDefault(void* ptr, int row, int column, void* idx)
{
	if (dynamic_cast<QSqlRelationalTableModel*>(static_cast<QObject*>(ptr))) {
		return new QModelIndex(static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::sibling(row, column, *static_cast<QModelIndex*>(idx)));
	} else if (dynamic_cast<QSqlTableModel*>(static_cast<QObject*>(ptr))) {
		return new QModelIndex(static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::sibling(row, column, *static_cast<QModelIndex*>(idx)));
	} else {
		return new QModelIndex(static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::sibling(row, column, *static_cast<QModelIndex*>(idx)));
	}
}

void* QSqlQueryModel_BuddyDefault(void* ptr, void* index)
{
	if (dynamic_cast<QSqlRelationalTableModel*>(static_cast<QObject*>(ptr))) {
		return new QModelIndex(static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::buddy(*static_cast<QModelIndex*>(index)));
	} else if (dynamic_cast<QSqlTableModel*>(static_cast<QObject*>(ptr))) {
		return new QModelIndex(static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::buddy(*static_cast<QModelIndex*>(index)));
	} else {
		return new QModelIndex(static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::buddy(*static_cast<QModelIndex*>(index)));
	}
}

char QSqlQueryModel_CanDropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	if (dynamic_cast<QSqlRelationalTableModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::canDropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
	} else if (dynamic_cast<QSqlTableModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::canDropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
	} else {
		return static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::canDropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
	}
}

char QSqlQueryModel_HasChildrenDefault(void* ptr, void* parent)
{
	if (dynamic_cast<QSqlRelationalTableModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::hasChildren(*static_cast<QModelIndex*>(parent));
	} else if (dynamic_cast<QSqlTableModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::hasChildren(*static_cast<QModelIndex*>(parent));
	} else {
		return static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::hasChildren(*static_cast<QModelIndex*>(parent));
	}
}

char QSqlQueryModel_InsertRowsDefault(void* ptr, int row, int count, void* parent)
{
	if (dynamic_cast<QSqlRelationalTableModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::insertRows(row, count, *static_cast<QModelIndex*>(parent));
	} else if (dynamic_cast<QSqlTableModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::insertRows(row, count, *static_cast<QModelIndex*>(parent));
	} else {
		return static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::insertRows(row, count, *static_cast<QModelIndex*>(parent));
	}
}

struct QtSql_PackedList QSqlQueryModel_ItemDataDefault(void* ptr, void* index)
{
	if (dynamic_cast<QSqlRelationalTableModel*>(static_cast<QObject*>(ptr))) {
		return ({ QMap<int, QVariant>* tmpValue8724bd = new QMap<int, QVariant>(static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::itemData(*static_cast<QModelIndex*>(index))); QtSql_PackedList { tmpValue8724bd, tmpValue8724bd->size() }; });
	} else if (dynamic_cast<QSqlTableModel*>(static_cast<QObject*>(ptr))) {
		return ({ QMap<int, QVariant>* tmpValue8724bd = new QMap<int, QVariant>(static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::itemData(*static_cast<QModelIndex*>(index))); QtSql_PackedList { tmpValue8724bd, tmpValue8724bd->size() }; });
	} else {
		return ({ QMap<int, QVariant>* tmpValue8724bd = new QMap<int, QVariant>(static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::itemData(*static_cast<QModelIndex*>(index))); QtSql_PackedList { tmpValue8724bd, tmpValue8724bd->size() }; });
	}
}

struct QtSql_PackedList QSqlQueryModel_MatchDefault(void* ptr, void* start, int role, void* value, int hits, long long flags)
{
	if (dynamic_cast<QSqlRelationalTableModel*>(static_cast<QObject*>(ptr))) {
		return ({ QList<QModelIndex>* tmpValue4356c2 = new QList<QModelIndex>(static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::match(*static_cast<QModelIndex*>(start), role, *static_cast<QVariant*>(value), hits, static_cast<Qt::MatchFlag>(flags))); QtSql_PackedList { tmpValue4356c2, tmpValue4356c2->size() }; });
	} else if (dynamic_cast<QSqlTableModel*>(static_cast<QObject*>(ptr))) {
		return ({ QList<QModelIndex>* tmpValue4356c2 = new QList<QModelIndex>(static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::match(*static_cast<QModelIndex*>(start), role, *static_cast<QVariant*>(value), hits, static_cast<Qt::MatchFlag>(flags))); QtSql_PackedList { tmpValue4356c2, tmpValue4356c2->size() }; });
	} else {
		return ({ QList<QModelIndex>* tmpValue4356c2 = new QList<QModelIndex>(static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::match(*static_cast<QModelIndex*>(start), role, *static_cast<QVariant*>(value), hits, static_cast<Qt::MatchFlag>(flags))); QtSql_PackedList { tmpValue4356c2, tmpValue4356c2->size() }; });
	}
}

void* QSqlQueryModel_MimeDataDefault(void* ptr, void* indexes)
{
	if (dynamic_cast<QSqlRelationalTableModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::mimeData(({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(indexes); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
	} else if (dynamic_cast<QSqlTableModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::mimeData(({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(indexes); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
	} else {
		return static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::mimeData(({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(indexes); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
	}
}

struct QtSql_PackedString QSqlQueryModel_MimeTypesDefault(void* ptr)
{
	if (dynamic_cast<QSqlRelationalTableModel*>(static_cast<QObject*>(ptr))) {
		return ({ QByteArray* t026d85 = new QByteArray(static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::mimeTypes().join("¡¦!").toUtf8()); QtSql_PackedString { const_cast<char*>(t026d85->prepend("WHITESPACE").constData()+10), t026d85->size()-10, t026d85 }; });
	} else if (dynamic_cast<QSqlTableModel*>(static_cast<QObject*>(ptr))) {
		return ({ QByteArray* t026d85 = new QByteArray(static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::mimeTypes().join("¡¦!").toUtf8()); QtSql_PackedString { const_cast<char*>(t026d85->prepend("WHITESPACE").constData()+10), t026d85->size()-10, t026d85 }; });
	} else {
		return ({ QByteArray* t026d85 = new QByteArray(static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::mimeTypes().join("¡¦!").toUtf8()); QtSql_PackedString { const_cast<char*>(t026d85->prepend("WHITESPACE").constData()+10), t026d85->size()-10, t026d85 }; });
	}
}

char QSqlQueryModel_MoveColumnsDefault(void* ptr, void* sourceParent, int sourceColumn, int count, void* destinationParent, int destinationChild)
{
	if (dynamic_cast<QSqlRelationalTableModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::moveColumns(*static_cast<QModelIndex*>(sourceParent), sourceColumn, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
	} else if (dynamic_cast<QSqlTableModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::moveColumns(*static_cast<QModelIndex*>(sourceParent), sourceColumn, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
	} else {
		return static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::moveColumns(*static_cast<QModelIndex*>(sourceParent), sourceColumn, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
	}
}

char QSqlQueryModel_MoveRowsDefault(void* ptr, void* sourceParent, int sourceRow, int count, void* destinationParent, int destinationChild)
{
	if (dynamic_cast<QSqlRelationalTableModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::moveRows(*static_cast<QModelIndex*>(sourceParent), sourceRow, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
	} else if (dynamic_cast<QSqlTableModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::moveRows(*static_cast<QModelIndex*>(sourceParent), sourceRow, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
	} else {
		return static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::moveRows(*static_cast<QModelIndex*>(sourceParent), sourceRow, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
	}
}

void* QSqlQueryModel_ParentDefault(void* ptr, void* index)
{
	if (dynamic_cast<QSqlRelationalTableModel*>(static_cast<QObject*>(ptr))) {
		return new QModelIndex(static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::parent(*static_cast<QModelIndex*>(index)));
	} else if (dynamic_cast<QSqlTableModel*>(static_cast<QObject*>(ptr))) {
		return new QModelIndex(static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::parent(*static_cast<QModelIndex*>(index)));
	} else {
		return new QModelIndex(static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::parent(*static_cast<QModelIndex*>(index)));
	}
}

char QSqlQueryModel_RemoveRowsDefault(void* ptr, int row, int count, void* parent)
{
	if (dynamic_cast<QSqlRelationalTableModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::removeRows(row, count, *static_cast<QModelIndex*>(parent));
	} else if (dynamic_cast<QSqlTableModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::removeRows(row, count, *static_cast<QModelIndex*>(parent));
	} else {
		return static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::removeRows(row, count, *static_cast<QModelIndex*>(parent));
	}
}

void QSqlQueryModel_ResetInternalDataDefault(void* ptr)
{
	if (dynamic_cast<QSqlRelationalTableModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::resetInternalData();
	} else if (dynamic_cast<QSqlTableModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::resetInternalData();
	} else {
		static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::resetInternalData();
	}
}

void QSqlQueryModel_RevertDefault(void* ptr)
{
	if (dynamic_cast<QSqlRelationalTableModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::revert();
	} else if (dynamic_cast<QSqlTableModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::revert();
	} else {
		static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::revert();
	}
}

char QSqlQueryModel_SetDataDefault(void* ptr, void* index, void* value, int role)
{
	if (dynamic_cast<QSqlRelationalTableModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
	} else if (dynamic_cast<QSqlTableModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
	} else {
		return static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
	}
}

char QSqlQueryModel_SetItemDataDefault(void* ptr, void* index, void* roles)
{
	if (dynamic_cast<QSqlRelationalTableModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::setItemData(*static_cast<QModelIndex*>(index), *static_cast<QMap<int, QVariant>*>(roles));
	} else if (dynamic_cast<QSqlTableModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::setItemData(*static_cast<QModelIndex*>(index), *static_cast<QMap<int, QVariant>*>(roles));
	} else {
		return static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::setItemData(*static_cast<QModelIndex*>(index), *static_cast<QMap<int, QVariant>*>(roles));
	}
}

void QSqlQueryModel_SortDefault(void* ptr, int column, long long order)
{
	if (dynamic_cast<QSqlRelationalTableModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::sort(column, static_cast<Qt::SortOrder>(order));
	} else if (dynamic_cast<QSqlTableModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::sort(column, static_cast<Qt::SortOrder>(order));
	} else {
		static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::sort(column, static_cast<Qt::SortOrder>(order));
	}
}

void* QSqlQueryModel_SpanDefault(void* ptr, void* index)
{
	if (dynamic_cast<QSqlRelationalTableModel*>(static_cast<QObject*>(ptr))) {
		return ({ QSize tmpValue = static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::span(*static_cast<QModelIndex*>(index)); new QSize(tmpValue.width(), tmpValue.height()); });
	} else if (dynamic_cast<QSqlTableModel*>(static_cast<QObject*>(ptr))) {
		return ({ QSize tmpValue = static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::span(*static_cast<QModelIndex*>(index)); new QSize(tmpValue.width(), tmpValue.height()); });
	} else {
		return ({ QSize tmpValue = static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::span(*static_cast<QModelIndex*>(index)); new QSize(tmpValue.width(), tmpValue.height()); });
	}
}

char QSqlQueryModel_SubmitDefault(void* ptr)
{
	if (dynamic_cast<QSqlRelationalTableModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::submit();
	} else if (dynamic_cast<QSqlTableModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::submit();
	} else {
		return static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::submit();
	}
}

long long QSqlQueryModel_SupportedDragActionsDefault(void* ptr)
{
	if (dynamic_cast<QSqlRelationalTableModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::supportedDragActions();
	} else if (dynamic_cast<QSqlTableModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::supportedDragActions();
	} else {
		return static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::supportedDragActions();
	}
}

long long QSqlQueryModel_SupportedDropActionsDefault(void* ptr)
{
	if (dynamic_cast<QSqlRelationalTableModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::supportedDropActions();
	} else if (dynamic_cast<QSqlTableModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::supportedDropActions();
	} else {
		return static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::supportedDropActions();
	}
}

void QSqlQueryModel_ChildEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QSqlRelationalTableModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QSqlTableModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::childEvent(static_cast<QChildEvent*>(event));
	} else {
		static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::childEvent(static_cast<QChildEvent*>(event));
	}
}

void QSqlQueryModel_ConnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QSqlRelationalTableModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QSqlTableModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::connectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

void QSqlQueryModel_CustomEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QSqlRelationalTableModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QSqlTableModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::customEvent(static_cast<QEvent*>(event));
	} else {
		static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::customEvent(static_cast<QEvent*>(event));
	}
}

void QSqlQueryModel_DeleteLaterDefault(void* ptr)
{
	if (dynamic_cast<QSqlRelationalTableModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::deleteLater();
	} else if (dynamic_cast<QSqlTableModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::deleteLater();
	} else {
		static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::deleteLater();
	}
}

void QSqlQueryModel_DisconnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QSqlRelationalTableModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QSqlTableModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

char QSqlQueryModel_EventDefault(void* ptr, void* e)
{
	if (dynamic_cast<QSqlRelationalTableModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QSqlTableModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::event(static_cast<QEvent*>(e));
	} else {
		return static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::event(static_cast<QEvent*>(e));
	}
}

char QSqlQueryModel_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QSqlRelationalTableModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QSqlTableModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

void* QSqlQueryModel_MetaObjectDefault(void* ptr)
{
	if (dynamic_cast<QSqlRelationalTableModel*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::metaObject());
	} else if (dynamic_cast<QSqlTableModel*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::metaObject());
	} else {
		return const_cast<QMetaObject*>(static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::metaObject());
	}
}

void QSqlQueryModel_TimerEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QSqlRelationalTableModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QSqlTableModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::timerEvent(static_cast<QTimerEvent*>(event));
	} else {
		static_cast<QSqlQueryModel*>(ptr)->QSqlQueryModel::timerEvent(static_cast<QTimerEvent*>(event));
	}
}

Q_DECLARE_METATYPE(QSqlRecord)
Q_DECLARE_METATYPE(QSqlRecord*)
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

char QSqlRecord_Contains(void* ptr, struct QtSql_PackedString name)
{
	return static_cast<QSqlRecord*>(ptr)->contains(QString::fromUtf8(name.data, name.len));
}

int QSqlRecord_Count(void* ptr)
{
	return static_cast<QSqlRecord*>(ptr)->count();
}

void* QSqlRecord_Field(void* ptr, int index)
{
	return new QSqlField(static_cast<QSqlRecord*>(ptr)->field(index));
}

void* QSqlRecord_Field2(void* ptr, struct QtSql_PackedString name)
{
	return new QSqlField(static_cast<QSqlRecord*>(ptr)->field(QString::fromUtf8(name.data, name.len)));
}

struct QtSql_PackedString QSqlRecord_FieldName(void* ptr, int index)
{
	return ({ QByteArray* te132cb = new QByteArray(static_cast<QSqlRecord*>(ptr)->fieldName(index).toUtf8()); QtSql_PackedString { const_cast<char*>(te132cb->prepend("WHITESPACE").constData()+10), te132cb->size()-10, te132cb }; });
}

int QSqlRecord_IndexOf(void* ptr, struct QtSql_PackedString name)
{
	return static_cast<QSqlRecord*>(ptr)->indexOf(QString::fromUtf8(name.data, name.len));
}

void QSqlRecord_Insert(void* ptr, int pos, void* field)
{
	static_cast<QSqlRecord*>(ptr)->insert(pos, *static_cast<QSqlField*>(field));
}

char QSqlRecord_IsEmpty(void* ptr)
{
	return static_cast<QSqlRecord*>(ptr)->isEmpty();
}

char QSqlRecord_IsGenerated(void* ptr, struct QtSql_PackedString name)
{
	return static_cast<QSqlRecord*>(ptr)->isGenerated(QString::fromUtf8(name.data, name.len));
}

char QSqlRecord_IsGenerated2(void* ptr, int index)
{
	return static_cast<QSqlRecord*>(ptr)->isGenerated(index);
}

char QSqlRecord_IsNull(void* ptr, struct QtSql_PackedString name)
{
	return static_cast<QSqlRecord*>(ptr)->isNull(QString::fromUtf8(name.data, name.len));
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

void QSqlRecord_SetGenerated(void* ptr, struct QtSql_PackedString name, char generated)
{
	static_cast<QSqlRecord*>(ptr)->setGenerated(QString::fromUtf8(name.data, name.len), generated != 0);
}

void QSqlRecord_SetGenerated2(void* ptr, int index, char generated)
{
	static_cast<QSqlRecord*>(ptr)->setGenerated(index, generated != 0);
}

void QSqlRecord_SetNull(void* ptr, int index)
{
	static_cast<QSqlRecord*>(ptr)->setNull(index);
}

void QSqlRecord_SetNull2(void* ptr, struct QtSql_PackedString name)
{
	static_cast<QSqlRecord*>(ptr)->setNull(QString::fromUtf8(name.data, name.len));
}

void QSqlRecord_SetValue(void* ptr, int index, void* val)
{
	static_cast<QSqlRecord*>(ptr)->setValue(index, *static_cast<QVariant*>(val));
}

void QSqlRecord_SetValue2(void* ptr, struct QtSql_PackedString name, void* val)
{
	static_cast<QSqlRecord*>(ptr)->setValue(QString::fromUtf8(name.data, name.len), *static_cast<QVariant*>(val));
}

void* QSqlRecord_Value(void* ptr, int index)
{
	return new QVariant(static_cast<QSqlRecord*>(ptr)->value(index));
}

void* QSqlRecord_Value2(void* ptr, struct QtSql_PackedString name)
{
	return new QVariant(static_cast<QSqlRecord*>(ptr)->value(QString::fromUtf8(name.data, name.len)));
}

void QSqlRecord_DestroyQSqlRecord(void* ptr)
{
	static_cast<QSqlRecord*>(ptr)->~QSqlRecord();
}

Q_DECLARE_METATYPE(QSqlRelation)
Q_DECLARE_METATYPE(QSqlRelation*)
void* QSqlRelation_NewQSqlRelation()
{
	return new QSqlRelation();
}

void* QSqlRelation_NewQSqlRelation2(struct QtSql_PackedString tableName, struct QtSql_PackedString indexColumn, struct QtSql_PackedString displayColumn)
{
	return new QSqlRelation(QString::fromUtf8(tableName.data, tableName.len), QString::fromUtf8(indexColumn.data, indexColumn.len), QString::fromUtf8(displayColumn.data, displayColumn.len));
}

struct QtSql_PackedString QSqlRelation_DisplayColumn(void* ptr)
{
	return ({ QByteArray* t5efc39 = new QByteArray(static_cast<QSqlRelation*>(ptr)->displayColumn().toUtf8()); QtSql_PackedString { const_cast<char*>(t5efc39->prepend("WHITESPACE").constData()+10), t5efc39->size()-10, t5efc39 }; });
}

struct QtSql_PackedString QSqlRelation_IndexColumn(void* ptr)
{
	return ({ QByteArray* t3d4445 = new QByteArray(static_cast<QSqlRelation*>(ptr)->indexColumn().toUtf8()); QtSql_PackedString { const_cast<char*>(t3d4445->prepend("WHITESPACE").constData()+10), t3d4445->size()-10, t3d4445 }; });
}

char QSqlRelation_IsValid(void* ptr)
{
	return static_cast<QSqlRelation*>(ptr)->isValid();
}

void QSqlRelation_Swap(void* ptr, void* other)
{
	static_cast<QSqlRelation*>(ptr)->swap(*static_cast<QSqlRelation*>(other));
}

struct QtSql_PackedString QSqlRelation_TableName(void* ptr)
{
	return ({ QByteArray* t33aa41 = new QByteArray(static_cast<QSqlRelation*>(ptr)->tableName().toUtf8()); QtSql_PackedString { const_cast<char*>(t33aa41->prepend("WHITESPACE").constData()+10), t33aa41->size()-10, t33aa41 }; });
}

class MyQSqlRelationalDelegate: public QSqlRelationalDelegate
{
public:
	MyQSqlRelationalDelegate(QObject *parent = Q_NULLPTR) : QSqlRelationalDelegate(parent) {QSqlRelationalDelegate_QSqlRelationalDelegate_QRegisterMetaType();};
	QWidget * createEditor(QWidget * parent, const QStyleOptionViewItem & option, const QModelIndex & index) const { return static_cast<QWidget*>(callbackQSqlRelationalDelegate_CreateEditor(const_cast<void*>(static_cast<const void*>(this)), parent, const_cast<QStyleOptionViewItem*>(&option), const_cast<QModelIndex*>(&index))); };
	void setModelData(QWidget * editor, QAbstractItemModel * model, const QModelIndex & index) const { callbackQSqlRelationalDelegate_SetModelData(const_cast<void*>(static_cast<const void*>(this)), editor, model, const_cast<QModelIndex*>(&index)); };
	 ~MyQSqlRelationalDelegate() { callbackQSqlRelationalDelegate_DestroyQSqlRelationalDelegate(this); };
	void drawCheck(QPainter * painter, const QStyleOptionViewItem & option, const QRect & rect, Qt::CheckState state) const { callbackQSqlRelationalDelegate_DrawCheck(const_cast<void*>(static_cast<const void*>(this)), painter, const_cast<QStyleOptionViewItem*>(&option), const_cast<QRect*>(&rect), state); };
	void drawDecoration(QPainter * painter, const QStyleOptionViewItem & option, const QRect & rect, const QPixmap & pixmap) const { callbackQSqlRelationalDelegate_DrawDecoration(const_cast<void*>(static_cast<const void*>(this)), painter, const_cast<QStyleOptionViewItem*>(&option), const_cast<QRect*>(&rect), const_cast<QPixmap*>(&pixmap)); };
	void drawDisplay(QPainter * painter, const QStyleOptionViewItem & option, const QRect & rect, const QString & text) const { QByteArray* t372ea0 = new QByteArray(text.toUtf8()); QtSql_PackedString textPacked = { const_cast<char*>(t372ea0->prepend("WHITESPACE").constData()+10), t372ea0->size()-10, t372ea0 };callbackQSqlRelationalDelegate_DrawDisplay(const_cast<void*>(static_cast<const void*>(this)), painter, const_cast<QStyleOptionViewItem*>(&option), const_cast<QRect*>(&rect), textPacked); };
	void drawFocus(QPainter * painter, const QStyleOptionViewItem & option, const QRect & rect) const { callbackQSqlRelationalDelegate_DrawFocus(const_cast<void*>(static_cast<const void*>(this)), painter, const_cast<QStyleOptionViewItem*>(&option), const_cast<QRect*>(&rect)); };
	bool editorEvent(QEvent * event, QAbstractItemModel * model, const QStyleOptionViewItem & option, const QModelIndex & index) { return callbackQSqlRelationalDelegate_EditorEvent(this, event, model, const_cast<QStyleOptionViewItem*>(&option), const_cast<QModelIndex*>(&index)) != 0; };
	bool eventFilter(QObject * editor, QEvent * event) { return callbackQSqlRelationalDelegate_EventFilter(this, editor, event) != 0; };
	void paint(QPainter * painter, const QStyleOptionViewItem & option, const QModelIndex & index) const { callbackQSqlRelationalDelegate_Paint(const_cast<void*>(static_cast<const void*>(this)), painter, const_cast<QStyleOptionViewItem*>(&option), const_cast<QModelIndex*>(&index)); };
	void setEditorData(QWidget * editor, const QModelIndex & index) const { callbackQSqlRelationalDelegate_SetEditorData(const_cast<void*>(static_cast<const void*>(this)), editor, const_cast<QModelIndex*>(&index)); };
	QSize sizeHint(const QStyleOptionViewItem & option, const QModelIndex & index) const { return *static_cast<QSize*>(callbackQSqlRelationalDelegate_SizeHint(const_cast<void*>(static_cast<const void*>(this)), const_cast<QStyleOptionViewItem*>(&option), const_cast<QModelIndex*>(&index))); };
	void updateEditorGeometry(QWidget * editor, const QStyleOptionViewItem & option, const QModelIndex & index) const { callbackQSqlRelationalDelegate_UpdateEditorGeometry(const_cast<void*>(static_cast<const void*>(this)), editor, const_cast<QStyleOptionViewItem*>(&option), const_cast<QModelIndex*>(&index)); };
	void Signal_CloseEditor(QWidget * editor, QAbstractItemDelegate::EndEditHint hint) { callbackQSqlRelationalDelegate_CloseEditor(this, editor, hint); };
	void Signal_CommitData(QWidget * editor) { callbackQSqlRelationalDelegate_CommitData(this, editor); };
	void destroyEditor(QWidget * editor, const QModelIndex & index) const { callbackQSqlRelationalDelegate_DestroyEditor(const_cast<void*>(static_cast<const void*>(this)), editor, const_cast<QModelIndex*>(&index)); };
	bool helpEvent(QHelpEvent * event, QAbstractItemView * view, const QStyleOptionViewItem & option, const QModelIndex & index) { return callbackQSqlRelationalDelegate_HelpEvent(this, event, view, const_cast<QStyleOptionViewItem*>(&option), const_cast<QModelIndex*>(&index)) != 0; };
	void Signal_SizeHintChanged(const QModelIndex & index) { callbackQSqlRelationalDelegate_SizeHintChanged(this, const_cast<QModelIndex*>(&index)); };
	void childEvent(QChildEvent * event) { callbackQSqlRelationalDelegate_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSqlRelationalDelegate_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQSqlRelationalDelegate_CustomEvent(this, event); };
	void deleteLater() { callbackQSqlRelationalDelegate_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQSqlRelationalDelegate_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSqlRelationalDelegate_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQSqlRelationalDelegate_Event(this, e) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSqlRelationalDelegate_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtSql_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQSqlRelationalDelegate_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQSqlRelationalDelegate_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QSqlRelationalDelegate*)
Q_DECLARE_METATYPE(MyQSqlRelationalDelegate*)

int QSqlRelationalDelegate_QSqlRelationalDelegate_QRegisterMetaType(){qRegisterMetaType<QSqlRelationalDelegate*>(); return qRegisterMetaType<MyQSqlRelationalDelegate*>();}

void* QSqlRelationalDelegate_NewQSqlRelationalDelegate(void* parent)
{
	if (dynamic_cast<QAudioSystemPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQSqlRelationalDelegate(static_cast<QAudioSystemPlugin*>(parent));
	} else if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQSqlRelationalDelegate(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQSqlRelationalDelegate(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQSqlRelationalDelegate(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQSqlRelationalDelegate(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQSqlRelationalDelegate(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQSqlRelationalDelegate(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQSqlRelationalDelegate(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQSqlRelationalDelegate(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQSqlRelationalDelegate(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QMediaServiceProviderPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQSqlRelationalDelegate(static_cast<QMediaServiceProviderPlugin*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQSqlRelationalDelegate(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQSqlRelationalDelegate(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQSqlRelationalDelegate(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQSqlRelationalDelegate(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQSqlRelationalDelegate(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQSqlRelationalDelegate(static_cast<QRemoteObjectPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QScriptExtensionPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQSqlRelationalDelegate(static_cast<QScriptExtensionPlugin*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQSqlRelationalDelegate(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQSqlRelationalDelegate(static_cast<QWindow*>(parent));
	} else {
		return new MyQSqlRelationalDelegate(static_cast<QObject*>(parent));
	}
}

void* QSqlRelationalDelegate_CreateEditorDefault(void* ptr, void* parent, void* option, void* index)
{
		return static_cast<QSqlRelationalDelegate*>(ptr)->QSqlRelationalDelegate::createEditor(static_cast<QWidget*>(parent), *static_cast<QStyleOptionViewItem*>(option), *static_cast<QModelIndex*>(index));
}

void QSqlRelationalDelegate_SetModelDataDefault(void* ptr, void* editor, void* model, void* index)
{
		static_cast<QSqlRelationalDelegate*>(ptr)->QSqlRelationalDelegate::setModelData(static_cast<QWidget*>(editor), static_cast<QAbstractItemModel*>(model), *static_cast<QModelIndex*>(index));
}

void QSqlRelationalDelegate_DestroyQSqlRelationalDelegate(void* ptr)
{
	static_cast<QSqlRelationalDelegate*>(ptr)->~QSqlRelationalDelegate();
}

void QSqlRelationalDelegate_DestroyQSqlRelationalDelegateDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QSqlRelationalDelegate___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QSqlRelationalDelegate___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSqlRelationalDelegate___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QSqlRelationalDelegate___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QSqlRelationalDelegate___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QSqlRelationalDelegate___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QSqlRelationalDelegate___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QSqlRelationalDelegate___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSqlRelationalDelegate___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QSqlRelationalDelegate___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QSqlRelationalDelegate___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSqlRelationalDelegate___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void QSqlRelationalDelegate_DrawCheckDefault(void* ptr, void* painter, void* option, void* rect, long long state)
{
		static_cast<QSqlRelationalDelegate*>(ptr)->QSqlRelationalDelegate::drawCheck(static_cast<QPainter*>(painter), *static_cast<QStyleOptionViewItem*>(option), *static_cast<QRect*>(rect), static_cast<Qt::CheckState>(state));
}

void QSqlRelationalDelegate_DrawDecorationDefault(void* ptr, void* painter, void* option, void* rect, void* pixmap)
{
		static_cast<QSqlRelationalDelegate*>(ptr)->QSqlRelationalDelegate::drawDecoration(static_cast<QPainter*>(painter), *static_cast<QStyleOptionViewItem*>(option), *static_cast<QRect*>(rect), *static_cast<QPixmap*>(pixmap));
}

void QSqlRelationalDelegate_DrawDisplayDefault(void* ptr, void* painter, void* option, void* rect, struct QtSql_PackedString text)
{
		static_cast<QSqlRelationalDelegate*>(ptr)->QSqlRelationalDelegate::drawDisplay(static_cast<QPainter*>(painter), *static_cast<QStyleOptionViewItem*>(option), *static_cast<QRect*>(rect), QString::fromUtf8(text.data, text.len));
}

void QSqlRelationalDelegate_DrawFocusDefault(void* ptr, void* painter, void* option, void* rect)
{
		static_cast<QSqlRelationalDelegate*>(ptr)->QSqlRelationalDelegate::drawFocus(static_cast<QPainter*>(painter), *static_cast<QStyleOptionViewItem*>(option), *static_cast<QRect*>(rect));
}

char QSqlRelationalDelegate_EditorEventDefault(void* ptr, void* event, void* model, void* option, void* index)
{
		return static_cast<QSqlRelationalDelegate*>(ptr)->QSqlRelationalDelegate::editorEvent(static_cast<QEvent*>(event), static_cast<QAbstractItemModel*>(model), *static_cast<QStyleOptionViewItem*>(option), *static_cast<QModelIndex*>(index));
}

char QSqlRelationalDelegate_EventFilterDefault(void* ptr, void* editor, void* event)
{
		return static_cast<QSqlRelationalDelegate*>(ptr)->QSqlRelationalDelegate::eventFilter(static_cast<QObject*>(editor), static_cast<QEvent*>(event));
}

void QSqlRelationalDelegate_PaintDefault(void* ptr, void* painter, void* option, void* index)
{
		static_cast<QSqlRelationalDelegate*>(ptr)->QSqlRelationalDelegate::paint(static_cast<QPainter*>(painter), *static_cast<QStyleOptionViewItem*>(option), *static_cast<QModelIndex*>(index));
}

void QSqlRelationalDelegate_SetEditorDataDefault(void* ptr, void* editor, void* index)
{
		static_cast<QSqlRelationalDelegate*>(ptr)->QSqlRelationalDelegate::setEditorData(static_cast<QWidget*>(editor), *static_cast<QModelIndex*>(index));
}

void* QSqlRelationalDelegate_SizeHintDefault(void* ptr, void* option, void* index)
{
		return ({ QSize tmpValue = static_cast<QSqlRelationalDelegate*>(ptr)->QSqlRelationalDelegate::sizeHint(*static_cast<QStyleOptionViewItem*>(option), *static_cast<QModelIndex*>(index)); new QSize(tmpValue.width(), tmpValue.height()); });
}

void QSqlRelationalDelegate_UpdateEditorGeometryDefault(void* ptr, void* editor, void* option, void* index)
{
		static_cast<QSqlRelationalDelegate*>(ptr)->QSqlRelationalDelegate::updateEditorGeometry(static_cast<QWidget*>(editor), *static_cast<QStyleOptionViewItem*>(option), *static_cast<QModelIndex*>(index));
}

void QSqlRelationalDelegate_DestroyEditorDefault(void* ptr, void* editor, void* index)
{
		static_cast<QSqlRelationalDelegate*>(ptr)->QSqlRelationalDelegate::destroyEditor(static_cast<QWidget*>(editor), *static_cast<QModelIndex*>(index));
}

char QSqlRelationalDelegate_HelpEventDefault(void* ptr, void* event, void* view, void* option, void* index)
{
		return static_cast<QSqlRelationalDelegate*>(ptr)->QSqlRelationalDelegate::helpEvent(static_cast<QHelpEvent*>(event), static_cast<QAbstractItemView*>(view), *static_cast<QStyleOptionViewItem*>(option), *static_cast<QModelIndex*>(index));
}

void QSqlRelationalDelegate_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QSqlRelationalDelegate*>(ptr)->QSqlRelationalDelegate::childEvent(static_cast<QChildEvent*>(event));
}

void QSqlRelationalDelegate_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QSqlRelationalDelegate*>(ptr)->QSqlRelationalDelegate::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSqlRelationalDelegate_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QSqlRelationalDelegate*>(ptr)->QSqlRelationalDelegate::customEvent(static_cast<QEvent*>(event));
}

void QSqlRelationalDelegate_DeleteLaterDefault(void* ptr)
{
		static_cast<QSqlRelationalDelegate*>(ptr)->QSqlRelationalDelegate::deleteLater();
}

void QSqlRelationalDelegate_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QSqlRelationalDelegate*>(ptr)->QSqlRelationalDelegate::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QSqlRelationalDelegate_EventDefault(void* ptr, void* e)
{
		return static_cast<QSqlRelationalDelegate*>(ptr)->QSqlRelationalDelegate::event(static_cast<QEvent*>(e));
}

void* QSqlRelationalDelegate_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QSqlRelationalDelegate*>(ptr)->QSqlRelationalDelegate::metaObject());
}

void QSqlRelationalDelegate_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QSqlRelationalDelegate*>(ptr)->QSqlRelationalDelegate::timerEvent(static_cast<QTimerEvent*>(event));
}

class MyQSqlRelationalTableModel: public QSqlRelationalTableModel
{
public:
	MyQSqlRelationalTableModel(QObject *parent = Q_NULLPTR, QSqlDatabase db = QSqlDatabase()) : QSqlRelationalTableModel(parent, db) {QSqlRelationalTableModel_QSqlRelationalTableModel_QRegisterMetaType();};
	void clear() { callbackQSqlQueryModel_Clear(this); };
	QVariant data(const QModelIndex & index, int role) const { return *static_cast<QVariant*>(callbackQSqlQueryModel_Data(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index), role)); };
	bool insertRowIntoTable(const QSqlRecord & values) { return callbackQSqlTableModel_InsertRowIntoTable(this, const_cast<QSqlRecord*>(&values)) != 0; };
	QString orderByClause() const { return ({ QtSql_PackedString tempVal = callbackQSqlTableModel_OrderByClause(const_cast<void*>(static_cast<const void*>(this))); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	QSqlTableModel * relationModel(int column) const { return static_cast<QSqlTableModel*>(callbackQSqlRelationalTableModel_RelationModel(const_cast<void*>(static_cast<const void*>(this)), column)); };
	bool removeColumns(int column, int count, const QModelIndex & parent) { return callbackQSqlQueryModel_RemoveColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	void revertRow(int row) { callbackQSqlRelationalTableModel_RevertRow(this, row); };
	bool select() { return callbackQSqlRelationalTableModel_Select(this) != 0; };
	QString selectStatement() const { return ({ QtSql_PackedString tempVal = callbackQSqlTableModel_SelectStatement(const_cast<void*>(static_cast<const void*>(this))); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	bool setData(const QModelIndex & index, const QVariant & value, int role) { return callbackQSqlQueryModel_SetData(this, const_cast<QModelIndex*>(&index), const_cast<QVariant*>(&value), role) != 0; };
	void setRelation(int column, const QSqlRelation & relation) { callbackQSqlRelationalTableModel_SetRelation(this, column, const_cast<QSqlRelation*>(&relation)); };
	void setTable(const QString & table) { QByteArray* tc3ee13 = new QByteArray(table.toUtf8()); QtSql_PackedString tablePacked = { const_cast<char*>(tc3ee13->prepend("WHITESPACE").constData()+10), tc3ee13->size()-10, tc3ee13 };callbackQSqlTableModel_SetTable(this, tablePacked); };
	bool updateRowInTable(int row, const QSqlRecord & values) { return callbackQSqlTableModel_UpdateRowInTable(this, row, const_cast<QSqlRecord*>(&values)) != 0; };
	 ~MyQSqlRelationalTableModel() { callbackQSqlRelationalTableModel_DestroyQSqlRelationalTableModel(this); };
	void Signal_BeforeDelete(int row) { callbackQSqlTableModel_BeforeDelete(this, row); };
	void Signal_BeforeInsert(QSqlRecord & record) { callbackQSqlTableModel_BeforeInsert(this, static_cast<QSqlRecord*>(&record)); };
	void Signal_BeforeUpdate(int row, QSqlRecord & record) { callbackQSqlTableModel_BeforeUpdate(this, row, static_cast<QSqlRecord*>(&record)); };
	bool deleteRowFromTable(int row) { return callbackQSqlTableModel_DeleteRowFromTable(this, row) != 0; };
	Qt::ItemFlags flags(const QModelIndex & index) const { return static_cast<Qt::ItemFlag>(callbackQSqlQueryModel_Flags(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QVariant headerData(int section, Qt::Orientation orientation, int role) const { return *static_cast<QVariant*>(callbackQSqlQueryModel_HeaderData(const_cast<void*>(static_cast<const void*>(this)), section, orientation, role)); };
	QModelIndex indexInQuery(const QModelIndex & item) const { return *static_cast<QModelIndex*>(callbackQSqlQueryModel_IndexInQuery(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&item))); };
	bool insertRows(int row, int count, const QModelIndex & parent) { return callbackQSqlQueryModel_InsertRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	void Signal_PrimeInsert(int row, QSqlRecord & record) { callbackQSqlTableModel_PrimeInsert(this, row, static_cast<QSqlRecord*>(&record)); };
	bool removeRows(int row, int count, const QModelIndex & parent) { return callbackQSqlQueryModel_RemoveRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	void revertAll() { callbackQSqlTableModel_RevertAll(this); };
	int rowCount(const QModelIndex & parent) const { return callbackQSqlQueryModel_RowCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	bool selectRow(int row) { return callbackQSqlTableModel_SelectRow(this, row) != 0; };
	void setEditStrategy(QSqlTableModel::EditStrategy strategy) { callbackQSqlTableModel_SetEditStrategy(this, strategy); };
	void setFilter(const QString & filter) { QByteArray* t4bb4ca = new QByteArray(filter.toUtf8()); QtSql_PackedString filterPacked = { const_cast<char*>(t4bb4ca->prepend("WHITESPACE").constData()+10), t4bb4ca->size()-10, t4bb4ca };callbackQSqlTableModel_SetFilter(this, filterPacked); };
	void setSort(int column, Qt::SortOrder order) { callbackQSqlTableModel_SetSort(this, column, order); };
	void sort(int column, Qt::SortOrder order) { callbackQSqlQueryModel_Sort(this, column, order); };
	bool submitAll() { return callbackQSqlTableModel_SubmitAll(this) != 0; };
	bool canFetchMore(const QModelIndex & parent) const { return callbackQSqlQueryModel_CanFetchMore(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	int columnCount(const QModelIndex & index) const { return callbackQSqlQueryModel_ColumnCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index)); };
	void fetchMore(const QModelIndex & parent) { callbackQSqlQueryModel_FetchMore(this, const_cast<QModelIndex*>(&parent)); };
	bool insertColumns(int column, int count, const QModelIndex & parent) { return callbackQSqlQueryModel_InsertColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	void queryChange() { callbackQSqlQueryModel_QueryChange(this); };
	QHash<int, QByteArray> roleNames() const { return ({ QHash<int, QByteArray>* tmpP = static_cast<QHash<int, QByteArray>*>(callbackQSqlQueryModel_RoleNames(const_cast<void*>(static_cast<const void*>(this)))); QHash<int, QByteArray> tmpV = *tmpP; tmpP->~QHash(); free(tmpP); tmpV; }); };
	bool setHeaderData(int section, Qt::Orientation orientation, const QVariant & value, int role) { return callbackQSqlQueryModel_SetHeaderData(this, section, orientation, const_cast<QVariant*>(&value), role) != 0; };
	bool dropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) { return callbackQSqlQueryModel_DropMimeData(this, const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	QModelIndex index(int row, int column, const QModelIndex & parent) const { return *static_cast<QModelIndex*>(callbackQSqlQueryModel_Index(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&parent))); };
	QModelIndex sibling(int row, int column, const QModelIndex & idx) const { return *static_cast<QModelIndex*>(callbackQSqlQueryModel_Sibling(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&idx))); };
	QModelIndex buddy(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackQSqlQueryModel_Buddy(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool canDropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) const { return callbackQSqlQueryModel_CanDropMimeData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	void Signal_ColumnsAboutToBeInserted(const QModelIndex & parent, int first, int last) { callbackQSqlQueryModel_ColumnsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationColumn) { callbackQSqlQueryModel_ColumnsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationColumn); };
	void Signal_ColumnsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackQSqlQueryModel_ColumnsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsInserted(const QModelIndex & parent, int first, int last) { callbackQSqlQueryModel_ColumnsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int column) { callbackQSqlQueryModel_ColumnsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), column); };
	void Signal_ColumnsRemoved(const QModelIndex & parent, int first, int last) { callbackQSqlQueryModel_ColumnsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_DataChanged(const QModelIndex & topLeft, const QModelIndex & bottomRight, const QVector<int> & roles) { callbackQSqlQueryModel_DataChanged(this, const_cast<QModelIndex*>(&topLeft), const_cast<QModelIndex*>(&bottomRight), ({ QVector<int>* tmpValue037c88 = new QVector<int>(roles); QtSql_PackedList { tmpValue037c88, tmpValue037c88->size() }; })); };
	bool hasChildren(const QModelIndex & parent) const { return callbackQSqlQueryModel_HasChildren(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	void Signal_HeaderDataChanged(Qt::Orientation orientation, int first, int last) { callbackQSqlQueryModel_HeaderDataChanged(this, orientation, first, last); };
	QMap<int, QVariant> itemData(const QModelIndex & index) const { return ({ QMap<int, QVariant>* tmpP = static_cast<QMap<int, QVariant>*>(callbackQSqlQueryModel_ItemData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); QMap<int, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }); };
	void Signal_LayoutAboutToBeChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackQSqlQueryModel_LayoutAboutToBeChanged(this, ({ QList<QPersistentModelIndex>* tmpValuea664f1 = new QList<QPersistentModelIndex>(parents); QtSql_PackedList { tmpValuea664f1, tmpValuea664f1->size() }; }), hint); };
	void Signal_LayoutChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackQSqlQueryModel_LayoutChanged(this, ({ QList<QPersistentModelIndex>* tmpValuea664f1 = new QList<QPersistentModelIndex>(parents); QtSql_PackedList { tmpValuea664f1, tmpValuea664f1->size() }; }), hint); };
	QList<QModelIndex> match(const QModelIndex & start, int role, const QVariant & value, int hits, Qt::MatchFlags flags) const { return ({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(callbackQSqlQueryModel_Match(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&start), role, const_cast<QVariant*>(&value), hits, flags)); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	QMimeData * mimeData(const QModelIndexList & indexes) const { return static_cast<QMimeData*>(callbackQSqlQueryModel_MimeData(const_cast<void*>(static_cast<const void*>(this)), ({ QList<QModelIndex>* tmpValuee0adf2 = new QList<QModelIndex>(indexes); QtSql_PackedList { tmpValuee0adf2, tmpValuee0adf2->size() }; }))); };
	QStringList mimeTypes() const { return ({ QtSql_PackedString tempVal = callbackQSqlQueryModel_MimeTypes(const_cast<void*>(static_cast<const void*>(this))); QStringList ret = QString::fromUtf8(tempVal.data, tempVal.len).split("¡¦!", QString::SkipEmptyParts); free(tempVal.data); ret; }); };
	void Signal_ModelAboutToBeReset() { callbackQSqlQueryModel_ModelAboutToBeReset(this); };
	void Signal_ModelReset() { callbackQSqlQueryModel_ModelReset(this); };
	bool moveColumns(const QModelIndex & sourceParent, int sourceColumn, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackQSqlQueryModel_MoveColumns(this, const_cast<QModelIndex*>(&sourceParent), sourceColumn, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	bool moveRows(const QModelIndex & sourceParent, int sourceRow, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackQSqlQueryModel_MoveRows(this, const_cast<QModelIndex*>(&sourceParent), sourceRow, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	QModelIndex parent(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackQSqlQueryModel_Parent(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	void resetInternalData() { callbackQSqlQueryModel_ResetInternalData(this); };
	void revert() { callbackQSqlQueryModel_Revert(this); };
	void Signal_RowsAboutToBeInserted(const QModelIndex & parent, int start, int end) { callbackQSqlQueryModel_RowsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), start, end); };
	void Signal_RowsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationRow) { callbackQSqlQueryModel_RowsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationRow); };
	void Signal_RowsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackQSqlQueryModel_RowsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsInserted(const QModelIndex & parent, int first, int last) { callbackQSqlQueryModel_RowsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int row) { callbackQSqlQueryModel_RowsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), row); };
	void Signal_RowsRemoved(const QModelIndex & parent, int first, int last) { callbackQSqlQueryModel_RowsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	bool setItemData(const QModelIndex & index, const QMap<int, QVariant> & roles) { return callbackQSqlQueryModel_SetItemData(this, const_cast<QModelIndex*>(&index), ({ QMap<int, QVariant>* tmpValue037c88 = new QMap<int, QVariant>(roles); QtSql_PackedList { tmpValue037c88, tmpValue037c88->size() }; })) != 0; };
	QSize span(const QModelIndex & index) const { return *static_cast<QSize*>(callbackQSqlQueryModel_Span(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool submit() { return callbackQSqlQueryModel_Submit(this) != 0; };
	Qt::DropActions supportedDragActions() const { return static_cast<Qt::DropAction>(callbackQSqlQueryModel_SupportedDragActions(const_cast<void*>(static_cast<const void*>(this)))); };
	Qt::DropActions supportedDropActions() const { return static_cast<Qt::DropAction>(callbackQSqlQueryModel_SupportedDropActions(const_cast<void*>(static_cast<const void*>(this)))); };
	void childEvent(QChildEvent * event) { callbackQSqlQueryModel_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSqlQueryModel_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQSqlQueryModel_CustomEvent(this, event); };
	void deleteLater() { callbackQSqlQueryModel_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQSqlQueryModel_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSqlQueryModel_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQSqlQueryModel_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSqlQueryModel_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSqlQueryModel_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtSql_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQSqlQueryModel_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQSqlQueryModel_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QSqlRelationalTableModel*)
Q_DECLARE_METATYPE(MyQSqlRelationalTableModel*)

int QSqlRelationalTableModel_QSqlRelationalTableModel_QRegisterMetaType(){qRegisterMetaType<QSqlRelationalTableModel*>(); return qRegisterMetaType<MyQSqlRelationalTableModel*>();}

void* QSqlRelationalTableModel_NewQSqlRelationalTableModel(void* parent, void* db)
{
	if (dynamic_cast<QAudioSystemPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQSqlRelationalTableModel(static_cast<QAudioSystemPlugin*>(parent), *static_cast<QSqlDatabase*>(db));
	} else if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQSqlRelationalTableModel(static_cast<QCameraImageCapture*>(parent), *static_cast<QSqlDatabase*>(db));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQSqlRelationalTableModel(static_cast<QDBusPendingCallWatcher*>(parent), *static_cast<QSqlDatabase*>(db));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQSqlRelationalTableModel(static_cast<QExtensionFactory*>(parent), *static_cast<QSqlDatabase*>(db));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQSqlRelationalTableModel(static_cast<QExtensionManager*>(parent), *static_cast<QSqlDatabase*>(db));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQSqlRelationalTableModel(static_cast<QGraphicsObject*>(parent), *static_cast<QSqlDatabase*>(db));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQSqlRelationalTableModel(static_cast<QGraphicsWidget*>(parent), *static_cast<QSqlDatabase*>(db));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQSqlRelationalTableModel(static_cast<QLayout*>(parent), *static_cast<QSqlDatabase*>(db));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQSqlRelationalTableModel(static_cast<QMediaPlaylist*>(parent), *static_cast<QSqlDatabase*>(db));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQSqlRelationalTableModel(static_cast<QMediaRecorder*>(parent), *static_cast<QSqlDatabase*>(db));
	} else if (dynamic_cast<QMediaServiceProviderPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQSqlRelationalTableModel(static_cast<QMediaServiceProviderPlugin*>(parent), *static_cast<QSqlDatabase*>(db));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQSqlRelationalTableModel(static_cast<QOffscreenSurface*>(parent), *static_cast<QSqlDatabase*>(db));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQSqlRelationalTableModel(static_cast<QPaintDeviceWindow*>(parent), *static_cast<QSqlDatabase*>(db));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQSqlRelationalTableModel(static_cast<QPdfWriter*>(parent), *static_cast<QSqlDatabase*>(db));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQSqlRelationalTableModel(static_cast<QQuickItem*>(parent), *static_cast<QSqlDatabase*>(db));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQSqlRelationalTableModel(static_cast<QRadioData*>(parent), *static_cast<QSqlDatabase*>(db));
	} else if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQSqlRelationalTableModel(static_cast<QRemoteObjectPendingCallWatcher*>(parent), *static_cast<QSqlDatabase*>(db));
	} else if (dynamic_cast<QScriptExtensionPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQSqlRelationalTableModel(static_cast<QScriptExtensionPlugin*>(parent), *static_cast<QSqlDatabase*>(db));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQSqlRelationalTableModel(static_cast<QWidget*>(parent), *static_cast<QSqlDatabase*>(db));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQSqlRelationalTableModel(static_cast<QWindow*>(parent), *static_cast<QSqlDatabase*>(db));
	} else {
		return new MyQSqlRelationalTableModel(static_cast<QObject*>(parent), *static_cast<QSqlDatabase*>(db));
	}
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

void QSqlRelationalTableModel_DestroyQSqlRelationalTableModel(void* ptr)
{
	static_cast<QSqlRelationalTableModel*>(ptr)->~QSqlRelationalTableModel();
}

void QSqlRelationalTableModel_DestroyQSqlRelationalTableModelDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

class MyQSqlResult: public QSqlResult
{
public:
	MyQSqlResult(const QSqlDriver *db) : QSqlResult(db) {QSqlResult_QSqlResult_QRegisterMetaType();};
	void bindValue(int index, const QVariant & val, QSql::ParamType paramType) { callbackQSqlResult_BindValue(this, index, const_cast<QVariant*>(&val), paramType); };
	void bindValue(const QString & placeholder, const QVariant & val, QSql::ParamType paramType) { QByteArray* tff5543 = new QByteArray(placeholder.toUtf8()); QtSql_PackedString placeholderPacked = { const_cast<char*>(tff5543->prepend("WHITESPACE").constData()+10), tff5543->size()-10, tff5543 };callbackQSqlResult_BindValue2(this, placeholderPacked, const_cast<QVariant*>(&val), paramType); };
	QVariant data(int index) { return *static_cast<QVariant*>(callbackQSqlResult_Data(this, index)); };
	bool exec() { return callbackQSqlResult_Exec(this) != 0; };
	bool fetch(int index) { return callbackQSqlResult_Fetch(this, index) != 0; };
	bool fetchFirst() { return callbackQSqlResult_FetchFirst(this) != 0; };
	bool fetchLast() { return callbackQSqlResult_FetchLast(this) != 0; };
	bool fetchNext() { return callbackQSqlResult_FetchNext(this) != 0; };
	bool fetchPrevious() { return callbackQSqlResult_FetchPrevious(this) != 0; };
	QVariant handle() const { return *static_cast<QVariant*>(callbackQSqlResult_Handle(const_cast<void*>(static_cast<const void*>(this)))); };
	bool isNull(int index) { return callbackQSqlResult_IsNull(this, index) != 0; };
	QVariant lastInsertId() const { return *static_cast<QVariant*>(callbackQSqlResult_LastInsertId(const_cast<void*>(static_cast<const void*>(this)))); };
	int numRowsAffected() { return callbackQSqlResult_NumRowsAffected(this); };
	bool prepare(const QString & query) { QByteArray* t7cd914 = new QByteArray(query.toUtf8()); QtSql_PackedString queryPacked = { const_cast<char*>(t7cd914->prepend("WHITESPACE").constData()+10), t7cd914->size()-10, t7cd914 };return callbackQSqlResult_Prepare(this, queryPacked) != 0; };
	QSqlRecord record() const { return *static_cast<QSqlRecord*>(callbackQSqlResult_Record(const_cast<void*>(static_cast<const void*>(this)))); };
	bool reset(const QString & query) { QByteArray* t7cd914 = new QByteArray(query.toUtf8()); QtSql_PackedString queryPacked = { const_cast<char*>(t7cd914->prepend("WHITESPACE").constData()+10), t7cd914->size()-10, t7cd914 };return callbackQSqlResult_Reset(this, queryPacked) != 0; };
	bool savePrepare(const QString & query) { QByteArray* t7cd914 = new QByteArray(query.toUtf8()); QtSql_PackedString queryPacked = { const_cast<char*>(t7cd914->prepend("WHITESPACE").constData()+10), t7cd914->size()-10, t7cd914 };return callbackQSqlResult_SavePrepare(this, queryPacked) != 0; };
	void setActive(bool active) { callbackQSqlResult_SetActive(this, active); };
	void setAt(int index) { callbackQSqlResult_SetAt(this, index); };
	void setForwardOnly(bool forward) { callbackQSqlResult_SetForwardOnly(this, forward); };
	void setLastError(const QSqlError & error) { callbackQSqlResult_SetLastError(this, const_cast<QSqlError*>(&error)); };
	void setQuery(const QString & query) { QByteArray* t7cd914 = new QByteArray(query.toUtf8()); QtSql_PackedString queryPacked = { const_cast<char*>(t7cd914->prepend("WHITESPACE").constData()+10), t7cd914->size()-10, t7cd914 };callbackQSqlResult_SetQuery(this, queryPacked); };
	void setSelect(bool sele) { callbackQSqlResult_SetSelect(this, sele); };
	int size() { return callbackQSqlResult_Size(this); };
	 ~MyQSqlResult() { callbackQSqlResult_DestroyQSqlResult(this); };
};

Q_DECLARE_METATYPE(QSqlResult*)
Q_DECLARE_METATYPE(MyQSqlResult*)

int QSqlResult_QSqlResult_QRegisterMetaType(){qRegisterMetaType<QSqlResult*>(); return qRegisterMetaType<MyQSqlResult*>();}

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

void QSqlResult_BindValue(void* ptr, int index, void* val, long long paramType)
{
	static_cast<QSqlResult*>(ptr)->bindValue(index, *static_cast<QVariant*>(val), static_cast<QSql::ParamTypeFlag>(paramType));
}

void QSqlResult_BindValueDefault(void* ptr, int index, void* val, long long paramType)
{
		static_cast<QSqlResult*>(ptr)->QSqlResult::bindValue(index, *static_cast<QVariant*>(val), static_cast<QSql::ParamTypeFlag>(paramType));
}

void QSqlResult_BindValue2(void* ptr, struct QtSql_PackedString placeholder, void* val, long long paramType)
{
	static_cast<QSqlResult*>(ptr)->bindValue(QString::fromUtf8(placeholder.data, placeholder.len), *static_cast<QVariant*>(val), static_cast<QSql::ParamTypeFlag>(paramType));
}

void QSqlResult_BindValue2Default(void* ptr, struct QtSql_PackedString placeholder, void* val, long long paramType)
{
		static_cast<QSqlResult*>(ptr)->QSqlResult::bindValue(QString::fromUtf8(placeholder.data, placeholder.len), *static_cast<QVariant*>(val), static_cast<QSql::ParamTypeFlag>(paramType));
}

long long QSqlResult_BindValueType(void* ptr, int index)
{
	return static_cast<QSqlResult*>(ptr)->bindValueType(index);
}

long long QSqlResult_BindValueType2(void* ptr, struct QtSql_PackedString placeholder)
{
	return static_cast<QSqlResult*>(ptr)->bindValueType(QString::fromUtf8(placeholder.data, placeholder.len));
}

long long QSqlResult_BindingSyntax(void* ptr)
{
	return static_cast<QSqlResult*>(ptr)->bindingSyntax();
}

void* QSqlResult_BoundValue(void* ptr, int index)
{
	return new QVariant(static_cast<QSqlResult*>(ptr)->boundValue(index));
}

void* QSqlResult_BoundValue2(void* ptr, struct QtSql_PackedString placeholder)
{
	return new QVariant(static_cast<QSqlResult*>(ptr)->boundValue(QString::fromUtf8(placeholder.data, placeholder.len)));
}

int QSqlResult_BoundValueCount(void* ptr)
{
	return static_cast<QSqlResult*>(ptr)->boundValueCount();
}

struct QtSql_PackedString QSqlResult_BoundValueName(void* ptr, int index)
{
	return ({ QByteArray* t669294 = new QByteArray(static_cast<QSqlResult*>(ptr)->boundValueName(index).toUtf8()); QtSql_PackedString { const_cast<char*>(t669294->prepend("WHITESPACE").constData()+10), t669294->size()-10, t669294 }; });
}

struct QtSql_PackedList QSqlResult_BoundValues(void* ptr)
{
	return ({ QVector<QVariant>* tmpValue6d4565 = new QVector<QVariant>(static_cast<QSqlResult*>(ptr)->boundValues()); QtSql_PackedList { tmpValue6d4565, tmpValue6d4565->size() }; });
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
	return ({ QByteArray* t35c395 = new QByteArray(static_cast<QSqlResult*>(ptr)->executedQuery().toUtf8()); QtSql_PackedString { const_cast<char*>(t35c395->prepend("WHITESPACE").constData()+10), t35c395->size()-10, t35c395 }; });
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
	return ({ QByteArray* t9bd7f1 = new QByteArray(static_cast<QSqlResult*>(ptr)->lastQuery().toUtf8()); QtSql_PackedString { const_cast<char*>(t9bd7f1->prepend("WHITESPACE").constData()+10), t9bd7f1->size()-10, t9bd7f1 }; });
}

int QSqlResult_NumRowsAffected(void* ptr)
{
	return static_cast<QSqlResult*>(ptr)->numRowsAffected();
}

char QSqlResult_Prepare(void* ptr, struct QtSql_PackedString query)
{
	return static_cast<QSqlResult*>(ptr)->prepare(QString::fromUtf8(query.data, query.len));
}

char QSqlResult_PrepareDefault(void* ptr, struct QtSql_PackedString query)
{
		return static_cast<QSqlResult*>(ptr)->QSqlResult::prepare(QString::fromUtf8(query.data, query.len));
}

void* QSqlResult_Record(void* ptr)
{
	return new QSqlRecord(static_cast<QSqlResult*>(ptr)->record());
}

void* QSqlResult_RecordDefault(void* ptr)
{
		return new QSqlRecord(static_cast<QSqlResult*>(ptr)->QSqlResult::record());
}

char QSqlResult_Reset(void* ptr, struct QtSql_PackedString query)
{
	return static_cast<QSqlResult*>(ptr)->reset(QString::fromUtf8(query.data, query.len));
}

void QSqlResult_ResetBindCount(void* ptr)
{
	static_cast<QSqlResult*>(ptr)->resetBindCount();
}

char QSqlResult_SavePrepare(void* ptr, struct QtSql_PackedString query)
{
	return static_cast<QSqlResult*>(ptr)->savePrepare(QString::fromUtf8(query.data, query.len));
}

char QSqlResult_SavePrepareDefault(void* ptr, struct QtSql_PackedString query)
{
		return static_cast<QSqlResult*>(ptr)->QSqlResult::savePrepare(QString::fromUtf8(query.data, query.len));
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

void QSqlResult_SetQuery(void* ptr, struct QtSql_PackedString query)
{
	static_cast<QSqlResult*>(ptr)->setQuery(QString::fromUtf8(query.data, query.len));
}

void QSqlResult_SetQueryDefault(void* ptr, struct QtSql_PackedString query)
{
		static_cast<QSqlResult*>(ptr)->QSqlResult::setQuery(QString::fromUtf8(query.data, query.len));
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
	Q_UNUSED(ptr);

}

void* QSqlResult___boundValues_atList(void* ptr, int i)
{
	return new QVariant(({QVariant tmp = static_cast<QVector<QVariant>*>(ptr)->at(i); if (i == static_cast<QVector<QVariant>*>(ptr)->size()-1) { static_cast<QVector<QVariant>*>(ptr)->~QVector(); free(ptr); }; tmp; }));
}

void QSqlResult___boundValues_setList(void* ptr, void* i)
{
	static_cast<QVector<QVariant>*>(ptr)->append(*static_cast<QVariant*>(i));
}

void* QSqlResult___boundValues_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QVariant>();
}

class MyQSqlTableModel: public QSqlTableModel
{
public:
	MyQSqlTableModel(QObject *parent = Q_NULLPTR, QSqlDatabase db = QSqlDatabase()) : QSqlTableModel(parent, db) {QSqlTableModel_QSqlTableModel_QRegisterMetaType();};
	void Signal_BeforeDelete(int row) { callbackQSqlTableModel_BeforeDelete(this, row); };
	void Signal_BeforeInsert(QSqlRecord & record) { callbackQSqlTableModel_BeforeInsert(this, static_cast<QSqlRecord*>(&record)); };
	void Signal_BeforeUpdate(int row, QSqlRecord & record) { callbackQSqlTableModel_BeforeUpdate(this, row, static_cast<QSqlRecord*>(&record)); };
	void clear() { callbackQSqlQueryModel_Clear(this); };
	QVariant data(const QModelIndex & index, int role) const { return *static_cast<QVariant*>(callbackQSqlQueryModel_Data(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index), role)); };
	bool deleteRowFromTable(int row) { return callbackQSqlTableModel_DeleteRowFromTable(this, row) != 0; };
	Qt::ItemFlags flags(const QModelIndex & index) const { return static_cast<Qt::ItemFlag>(callbackQSqlQueryModel_Flags(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QVariant headerData(int section, Qt::Orientation orientation, int role) const { return *static_cast<QVariant*>(callbackQSqlQueryModel_HeaderData(const_cast<void*>(static_cast<const void*>(this)), section, orientation, role)); };
	QModelIndex indexInQuery(const QModelIndex & item) const { return *static_cast<QModelIndex*>(callbackQSqlQueryModel_IndexInQuery(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&item))); };
	bool insertRowIntoTable(const QSqlRecord & values) { return callbackQSqlTableModel_InsertRowIntoTable(this, const_cast<QSqlRecord*>(&values)) != 0; };
	bool insertRows(int row, int count, const QModelIndex & parent) { return callbackQSqlQueryModel_InsertRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	QString orderByClause() const { return ({ QtSql_PackedString tempVal = callbackQSqlTableModel_OrderByClause(const_cast<void*>(static_cast<const void*>(this))); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void Signal_PrimeInsert(int row, QSqlRecord & record) { callbackQSqlTableModel_PrimeInsert(this, row, static_cast<QSqlRecord*>(&record)); };
	bool removeColumns(int column, int count, const QModelIndex & parent) { return callbackQSqlQueryModel_RemoveColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool removeRows(int row, int count, const QModelIndex & parent) { return callbackQSqlQueryModel_RemoveRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	void revertAll() { callbackQSqlTableModel_RevertAll(this); };
	void revertRow(int row) { callbackQSqlTableModel_RevertRow(this, row); };
	int rowCount(const QModelIndex & parent) const { return callbackQSqlQueryModel_RowCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	bool select() { return callbackQSqlTableModel_Select(this) != 0; };
	bool selectRow(int row) { return callbackQSqlTableModel_SelectRow(this, row) != 0; };
	QString selectStatement() const { return ({ QtSql_PackedString tempVal = callbackQSqlTableModel_SelectStatement(const_cast<void*>(static_cast<const void*>(this))); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	bool setData(const QModelIndex & index, const QVariant & value, int role) { return callbackQSqlQueryModel_SetData(this, const_cast<QModelIndex*>(&index), const_cast<QVariant*>(&value), role) != 0; };
	void setEditStrategy(QSqlTableModel::EditStrategy strategy) { callbackQSqlTableModel_SetEditStrategy(this, strategy); };
	void setFilter(const QString & filter) { QByteArray* t4bb4ca = new QByteArray(filter.toUtf8()); QtSql_PackedString filterPacked = { const_cast<char*>(t4bb4ca->prepend("WHITESPACE").constData()+10), t4bb4ca->size()-10, t4bb4ca };callbackQSqlTableModel_SetFilter(this, filterPacked); };
	void setSort(int column, Qt::SortOrder order) { callbackQSqlTableModel_SetSort(this, column, order); };
	void setTable(const QString & tableName) { QByteArray* t3e7060 = new QByteArray(tableName.toUtf8()); QtSql_PackedString tableNamePacked = { const_cast<char*>(t3e7060->prepend("WHITESPACE").constData()+10), t3e7060->size()-10, t3e7060 };callbackQSqlTableModel_SetTable(this, tableNamePacked); };
	void sort(int column, Qt::SortOrder order) { callbackQSqlQueryModel_Sort(this, column, order); };
	bool submitAll() { return callbackQSqlTableModel_SubmitAll(this) != 0; };
	bool updateRowInTable(int row, const QSqlRecord & values) { return callbackQSqlTableModel_UpdateRowInTable(this, row, const_cast<QSqlRecord*>(&values)) != 0; };
	 ~MyQSqlTableModel() { callbackQSqlTableModel_DestroyQSqlTableModel(this); };
	bool canFetchMore(const QModelIndex & parent) const { return callbackQSqlQueryModel_CanFetchMore(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	int columnCount(const QModelIndex & index) const { return callbackQSqlQueryModel_ColumnCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index)); };
	void fetchMore(const QModelIndex & parent) { callbackQSqlQueryModel_FetchMore(this, const_cast<QModelIndex*>(&parent)); };
	bool insertColumns(int column, int count, const QModelIndex & parent) { return callbackQSqlQueryModel_InsertColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	void queryChange() { callbackQSqlQueryModel_QueryChange(this); };
	QHash<int, QByteArray> roleNames() const { return ({ QHash<int, QByteArray>* tmpP = static_cast<QHash<int, QByteArray>*>(callbackQSqlQueryModel_RoleNames(const_cast<void*>(static_cast<const void*>(this)))); QHash<int, QByteArray> tmpV = *tmpP; tmpP->~QHash(); free(tmpP); tmpV; }); };
	bool setHeaderData(int section, Qt::Orientation orientation, const QVariant & value, int role) { return callbackQSqlQueryModel_SetHeaderData(this, section, orientation, const_cast<QVariant*>(&value), role) != 0; };
	bool dropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) { return callbackQSqlQueryModel_DropMimeData(this, const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	QModelIndex index(int row, int column, const QModelIndex & parent) const { return *static_cast<QModelIndex*>(callbackQSqlQueryModel_Index(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&parent))); };
	QModelIndex sibling(int row, int column, const QModelIndex & idx) const { return *static_cast<QModelIndex*>(callbackQSqlQueryModel_Sibling(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&idx))); };
	QModelIndex buddy(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackQSqlQueryModel_Buddy(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool canDropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) const { return callbackQSqlQueryModel_CanDropMimeData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	void Signal_ColumnsAboutToBeInserted(const QModelIndex & parent, int first, int last) { callbackQSqlQueryModel_ColumnsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationColumn) { callbackQSqlQueryModel_ColumnsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationColumn); };
	void Signal_ColumnsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackQSqlQueryModel_ColumnsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsInserted(const QModelIndex & parent, int first, int last) { callbackQSqlQueryModel_ColumnsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int column) { callbackQSqlQueryModel_ColumnsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), column); };
	void Signal_ColumnsRemoved(const QModelIndex & parent, int first, int last) { callbackQSqlQueryModel_ColumnsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_DataChanged(const QModelIndex & topLeft, const QModelIndex & bottomRight, const QVector<int> & roles) { callbackQSqlQueryModel_DataChanged(this, const_cast<QModelIndex*>(&topLeft), const_cast<QModelIndex*>(&bottomRight), ({ QVector<int>* tmpValue037c88 = new QVector<int>(roles); QtSql_PackedList { tmpValue037c88, tmpValue037c88->size() }; })); };
	bool hasChildren(const QModelIndex & parent) const { return callbackQSqlQueryModel_HasChildren(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	void Signal_HeaderDataChanged(Qt::Orientation orientation, int first, int last) { callbackQSqlQueryModel_HeaderDataChanged(this, orientation, first, last); };
	QMap<int, QVariant> itemData(const QModelIndex & index) const { return ({ QMap<int, QVariant>* tmpP = static_cast<QMap<int, QVariant>*>(callbackQSqlQueryModel_ItemData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); QMap<int, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }); };
	void Signal_LayoutAboutToBeChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackQSqlQueryModel_LayoutAboutToBeChanged(this, ({ QList<QPersistentModelIndex>* tmpValuea664f1 = new QList<QPersistentModelIndex>(parents); QtSql_PackedList { tmpValuea664f1, tmpValuea664f1->size() }; }), hint); };
	void Signal_LayoutChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackQSqlQueryModel_LayoutChanged(this, ({ QList<QPersistentModelIndex>* tmpValuea664f1 = new QList<QPersistentModelIndex>(parents); QtSql_PackedList { tmpValuea664f1, tmpValuea664f1->size() }; }), hint); };
	QList<QModelIndex> match(const QModelIndex & start, int role, const QVariant & value, int hits, Qt::MatchFlags flags) const { return ({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(callbackQSqlQueryModel_Match(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&start), role, const_cast<QVariant*>(&value), hits, flags)); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	QMimeData * mimeData(const QModelIndexList & indexes) const { return static_cast<QMimeData*>(callbackQSqlQueryModel_MimeData(const_cast<void*>(static_cast<const void*>(this)), ({ QList<QModelIndex>* tmpValuee0adf2 = new QList<QModelIndex>(indexes); QtSql_PackedList { tmpValuee0adf2, tmpValuee0adf2->size() }; }))); };
	QStringList mimeTypes() const { return ({ QtSql_PackedString tempVal = callbackQSqlQueryModel_MimeTypes(const_cast<void*>(static_cast<const void*>(this))); QStringList ret = QString::fromUtf8(tempVal.data, tempVal.len).split("¡¦!", QString::SkipEmptyParts); free(tempVal.data); ret; }); };
	void Signal_ModelAboutToBeReset() { callbackQSqlQueryModel_ModelAboutToBeReset(this); };
	void Signal_ModelReset() { callbackQSqlQueryModel_ModelReset(this); };
	bool moveColumns(const QModelIndex & sourceParent, int sourceColumn, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackQSqlQueryModel_MoveColumns(this, const_cast<QModelIndex*>(&sourceParent), sourceColumn, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	bool moveRows(const QModelIndex & sourceParent, int sourceRow, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackQSqlQueryModel_MoveRows(this, const_cast<QModelIndex*>(&sourceParent), sourceRow, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	QModelIndex parent(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackQSqlQueryModel_Parent(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	void resetInternalData() { callbackQSqlQueryModel_ResetInternalData(this); };
	void revert() { callbackQSqlQueryModel_Revert(this); };
	void Signal_RowsAboutToBeInserted(const QModelIndex & parent, int start, int end) { callbackQSqlQueryModel_RowsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), start, end); };
	void Signal_RowsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationRow) { callbackQSqlQueryModel_RowsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationRow); };
	void Signal_RowsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackQSqlQueryModel_RowsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsInserted(const QModelIndex & parent, int first, int last) { callbackQSqlQueryModel_RowsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int row) { callbackQSqlQueryModel_RowsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), row); };
	void Signal_RowsRemoved(const QModelIndex & parent, int first, int last) { callbackQSqlQueryModel_RowsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	bool setItemData(const QModelIndex & index, const QMap<int, QVariant> & roles) { return callbackQSqlQueryModel_SetItemData(this, const_cast<QModelIndex*>(&index), ({ QMap<int, QVariant>* tmpValue037c88 = new QMap<int, QVariant>(roles); QtSql_PackedList { tmpValue037c88, tmpValue037c88->size() }; })) != 0; };
	QSize span(const QModelIndex & index) const { return *static_cast<QSize*>(callbackQSqlQueryModel_Span(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool submit() { return callbackQSqlQueryModel_Submit(this) != 0; };
	Qt::DropActions supportedDragActions() const { return static_cast<Qt::DropAction>(callbackQSqlQueryModel_SupportedDragActions(const_cast<void*>(static_cast<const void*>(this)))); };
	Qt::DropActions supportedDropActions() const { return static_cast<Qt::DropAction>(callbackQSqlQueryModel_SupportedDropActions(const_cast<void*>(static_cast<const void*>(this)))); };
	void childEvent(QChildEvent * event) { callbackQSqlQueryModel_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSqlQueryModel_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQSqlQueryModel_CustomEvent(this, event); };
	void deleteLater() { callbackQSqlQueryModel_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQSqlQueryModel_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSqlQueryModel_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQSqlQueryModel_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSqlQueryModel_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSqlQueryModel_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtSql_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQSqlQueryModel_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQSqlQueryModel_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QSqlTableModel*)
Q_DECLARE_METATYPE(MyQSqlTableModel*)

int QSqlTableModel_QSqlTableModel_QRegisterMetaType(){qRegisterMetaType<QSqlTableModel*>(); return qRegisterMetaType<MyQSqlTableModel*>();}

void* QSqlTableModel_NewQSqlTableModel(void* parent, void* db)
{
	if (dynamic_cast<QAudioSystemPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQSqlTableModel(static_cast<QAudioSystemPlugin*>(parent), *static_cast<QSqlDatabase*>(db));
	} else if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQSqlTableModel(static_cast<QCameraImageCapture*>(parent), *static_cast<QSqlDatabase*>(db));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQSqlTableModel(static_cast<QDBusPendingCallWatcher*>(parent), *static_cast<QSqlDatabase*>(db));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQSqlTableModel(static_cast<QExtensionFactory*>(parent), *static_cast<QSqlDatabase*>(db));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQSqlTableModel(static_cast<QExtensionManager*>(parent), *static_cast<QSqlDatabase*>(db));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQSqlTableModel(static_cast<QGraphicsObject*>(parent), *static_cast<QSqlDatabase*>(db));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQSqlTableModel(static_cast<QGraphicsWidget*>(parent), *static_cast<QSqlDatabase*>(db));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQSqlTableModel(static_cast<QLayout*>(parent), *static_cast<QSqlDatabase*>(db));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQSqlTableModel(static_cast<QMediaPlaylist*>(parent), *static_cast<QSqlDatabase*>(db));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQSqlTableModel(static_cast<QMediaRecorder*>(parent), *static_cast<QSqlDatabase*>(db));
	} else if (dynamic_cast<QMediaServiceProviderPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQSqlTableModel(static_cast<QMediaServiceProviderPlugin*>(parent), *static_cast<QSqlDatabase*>(db));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQSqlTableModel(static_cast<QOffscreenSurface*>(parent), *static_cast<QSqlDatabase*>(db));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQSqlTableModel(static_cast<QPaintDeviceWindow*>(parent), *static_cast<QSqlDatabase*>(db));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQSqlTableModel(static_cast<QPdfWriter*>(parent), *static_cast<QSqlDatabase*>(db));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQSqlTableModel(static_cast<QQuickItem*>(parent), *static_cast<QSqlDatabase*>(db));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQSqlTableModel(static_cast<QRadioData*>(parent), *static_cast<QSqlDatabase*>(db));
	} else if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQSqlTableModel(static_cast<QRemoteObjectPendingCallWatcher*>(parent), *static_cast<QSqlDatabase*>(db));
	} else if (dynamic_cast<QScriptExtensionPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQSqlTableModel(static_cast<QScriptExtensionPlugin*>(parent), *static_cast<QSqlDatabase*>(db));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQSqlTableModel(static_cast<QWidget*>(parent), *static_cast<QSqlDatabase*>(db));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQSqlTableModel(static_cast<QWindow*>(parent), *static_cast<QSqlDatabase*>(db));
	} else {
		return new MyQSqlTableModel(static_cast<QObject*>(parent), *static_cast<QSqlDatabase*>(db));
	}
}

void QSqlTableModel_ConnectBeforeDelete(void* ptr, long long t)
{
	QObject::connect(static_cast<QSqlTableModel*>(ptr), static_cast<void (QSqlTableModel::*)(int)>(&QSqlTableModel::beforeDelete), static_cast<MyQSqlTableModel*>(ptr), static_cast<void (MyQSqlTableModel::*)(int)>(&MyQSqlTableModel::Signal_BeforeDelete), static_cast<Qt::ConnectionType>(t));
}

void QSqlTableModel_DisconnectBeforeDelete(void* ptr)
{
	QObject::disconnect(static_cast<QSqlTableModel*>(ptr), static_cast<void (QSqlTableModel::*)(int)>(&QSqlTableModel::beforeDelete), static_cast<MyQSqlTableModel*>(ptr), static_cast<void (MyQSqlTableModel::*)(int)>(&MyQSqlTableModel::Signal_BeforeDelete));
}

void QSqlTableModel_BeforeDelete(void* ptr, int row)
{
	static_cast<QSqlTableModel*>(ptr)->beforeDelete(row);
}

void QSqlTableModel_ConnectBeforeInsert(void* ptr, long long t)
{
	QObject::connect(static_cast<QSqlTableModel*>(ptr), static_cast<void (QSqlTableModel::*)(QSqlRecord &)>(&QSqlTableModel::beforeInsert), static_cast<MyQSqlTableModel*>(ptr), static_cast<void (MyQSqlTableModel::*)(QSqlRecord &)>(&MyQSqlTableModel::Signal_BeforeInsert), static_cast<Qt::ConnectionType>(t));
}

void QSqlTableModel_DisconnectBeforeInsert(void* ptr)
{
	QObject::disconnect(static_cast<QSqlTableModel*>(ptr), static_cast<void (QSqlTableModel::*)(QSqlRecord &)>(&QSqlTableModel::beforeInsert), static_cast<MyQSqlTableModel*>(ptr), static_cast<void (MyQSqlTableModel::*)(QSqlRecord &)>(&MyQSqlTableModel::Signal_BeforeInsert));
}

void QSqlTableModel_BeforeInsert(void* ptr, void* record)
{
	static_cast<QSqlTableModel*>(ptr)->beforeInsert(*static_cast<QSqlRecord*>(record));
}

void QSqlTableModel_ConnectBeforeUpdate(void* ptr, long long t)
{
	QObject::connect(static_cast<QSqlTableModel*>(ptr), static_cast<void (QSqlTableModel::*)(int, QSqlRecord &)>(&QSqlTableModel::beforeUpdate), static_cast<MyQSqlTableModel*>(ptr), static_cast<void (MyQSqlTableModel::*)(int, QSqlRecord &)>(&MyQSqlTableModel::Signal_BeforeUpdate), static_cast<Qt::ConnectionType>(t));
}

void QSqlTableModel_DisconnectBeforeUpdate(void* ptr)
{
	QObject::disconnect(static_cast<QSqlTableModel*>(ptr), static_cast<void (QSqlTableModel::*)(int, QSqlRecord &)>(&QSqlTableModel::beforeUpdate), static_cast<MyQSqlTableModel*>(ptr), static_cast<void (MyQSqlTableModel::*)(int, QSqlRecord &)>(&MyQSqlTableModel::Signal_BeforeUpdate));
}

void QSqlTableModel_BeforeUpdate(void* ptr, int row, void* record)
{
	static_cast<QSqlTableModel*>(ptr)->beforeUpdate(row, *static_cast<QSqlRecord*>(record));
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
	if (dynamic_cast<QSqlRelationalTableModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::deleteRowFromTable(row);
	} else {
		return static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::deleteRowFromTable(row);
	}
}

long long QSqlTableModel_EditStrategy(void* ptr)
{
	return static_cast<QSqlTableModel*>(ptr)->editStrategy();
}

int QSqlTableModel_FieldIndex(void* ptr, struct QtSql_PackedString fieldName)
{
	return static_cast<QSqlTableModel*>(ptr)->fieldIndex(QString::fromUtf8(fieldName.data, fieldName.len));
}

struct QtSql_PackedString QSqlTableModel_Filter(void* ptr)
{
	return ({ QByteArray* td0370f = new QByteArray(static_cast<QSqlTableModel*>(ptr)->filter().toUtf8()); QtSql_PackedString { const_cast<char*>(td0370f->prepend("WHITESPACE").constData()+10), td0370f->size()-10, td0370f }; });
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
	if (dynamic_cast<QSqlRelationalTableModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::insertRowIntoTable(*static_cast<QSqlRecord*>(values));
	} else {
		return static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::insertRowIntoTable(*static_cast<QSqlRecord*>(values));
	}
}

char QSqlTableModel_IsDirty(void* ptr, void* index)
{
	return static_cast<QSqlTableModel*>(ptr)->isDirty(*static_cast<QModelIndex*>(index));
}

char QSqlTableModel_IsDirty2(void* ptr)
{
	return static_cast<QSqlTableModel*>(ptr)->isDirty();
}

struct QtSql_PackedString QSqlTableModel_OrderByClause(void* ptr)
{
	return ({ QByteArray* tca221a = new QByteArray(static_cast<QSqlTableModel*>(ptr)->orderByClause().toUtf8()); QtSql_PackedString { const_cast<char*>(tca221a->prepend("WHITESPACE").constData()+10), tca221a->size()-10, tca221a }; });
}

struct QtSql_PackedString QSqlTableModel_OrderByClauseDefault(void* ptr)
{
	if (dynamic_cast<QSqlRelationalTableModel*>(static_cast<QObject*>(ptr))) {
		return ({ QByteArray* t5a4e03 = new QByteArray(static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::orderByClause().toUtf8()); QtSql_PackedString { const_cast<char*>(t5a4e03->prepend("WHITESPACE").constData()+10), t5a4e03->size()-10, t5a4e03 }; });
	} else {
		return ({ QByteArray* t5a4e03 = new QByteArray(static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::orderByClause().toUtf8()); QtSql_PackedString { const_cast<char*>(t5a4e03->prepend("WHITESPACE").constData()+10), t5a4e03->size()-10, t5a4e03 }; });
	}
}

void* QSqlTableModel_PrimaryKey(void* ptr)
{
	return new QSqlIndex(static_cast<QSqlTableModel*>(ptr)->primaryKey());
}

void* QSqlTableModel_PrimaryValues(void* ptr, int row)
{
	return new QSqlRecord(static_cast<QSqlTableModel*>(ptr)->primaryValues(row));
}

void QSqlTableModel_ConnectPrimeInsert(void* ptr, long long t)
{
	QObject::connect(static_cast<QSqlTableModel*>(ptr), static_cast<void (QSqlTableModel::*)(int, QSqlRecord &)>(&QSqlTableModel::primeInsert), static_cast<MyQSqlTableModel*>(ptr), static_cast<void (MyQSqlTableModel::*)(int, QSqlRecord &)>(&MyQSqlTableModel::Signal_PrimeInsert), static_cast<Qt::ConnectionType>(t));
}

void QSqlTableModel_DisconnectPrimeInsert(void* ptr)
{
	QObject::disconnect(static_cast<QSqlTableModel*>(ptr), static_cast<void (QSqlTableModel::*)(int, QSqlRecord &)>(&QSqlTableModel::primeInsert), static_cast<MyQSqlTableModel*>(ptr), static_cast<void (MyQSqlTableModel::*)(int, QSqlRecord &)>(&MyQSqlTableModel::Signal_PrimeInsert));
}

void QSqlTableModel_PrimeInsert(void* ptr, int row, void* record)
{
	static_cast<QSqlTableModel*>(ptr)->primeInsert(row, *static_cast<QSqlRecord*>(record));
}

void QSqlTableModel_RevertAll(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QSqlTableModel*>(ptr), "revertAll");
}

void QSqlTableModel_RevertAllDefault(void* ptr)
{
	if (dynamic_cast<QSqlRelationalTableModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::revertAll();
	} else {
		static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::revertAll();
	}
}

void QSqlTableModel_RevertRow(void* ptr, int row)
{
	static_cast<QSqlTableModel*>(ptr)->revertRow(row);
}

void QSqlTableModel_RevertRowDefault(void* ptr, int row)
{
	if (dynamic_cast<QSqlRelationalTableModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::revertRow(row);
	} else {
		static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::revertRow(row);
	}
}

char QSqlTableModel_Select(void* ptr)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QSqlTableModel*>(ptr), "select", Q_RETURN_ARG(bool, returnArg));
	return returnArg;
}

char QSqlTableModel_SelectDefault(void* ptr)
{
	if (dynamic_cast<QSqlRelationalTableModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::select();
	} else {
		return static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::select();
	}
}

char QSqlTableModel_SelectRow(void* ptr, int row)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QSqlTableModel*>(ptr), "selectRow", Q_RETURN_ARG(bool, returnArg), Q_ARG(int, row));
	return returnArg;
}

char QSqlTableModel_SelectRowDefault(void* ptr, int row)
{
	if (dynamic_cast<QSqlRelationalTableModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::selectRow(row);
	} else {
		return static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::selectRow(row);
	}
}

struct QtSql_PackedString QSqlTableModel_SelectStatement(void* ptr)
{
	return ({ QByteArray* tf1b779 = new QByteArray(static_cast<QSqlTableModel*>(ptr)->selectStatement().toUtf8()); QtSql_PackedString { const_cast<char*>(tf1b779->prepend("WHITESPACE").constData()+10), tf1b779->size()-10, tf1b779 }; });
}

struct QtSql_PackedString QSqlTableModel_SelectStatementDefault(void* ptr)
{
	if (dynamic_cast<QSqlRelationalTableModel*>(static_cast<QObject*>(ptr))) {
		return ({ QByteArray* t5344b7 = new QByteArray(static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::selectStatement().toUtf8()); QtSql_PackedString { const_cast<char*>(t5344b7->prepend("WHITESPACE").constData()+10), t5344b7->size()-10, t5344b7 }; });
	} else {
		return ({ QByteArray* t5344b7 = new QByteArray(static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::selectStatement().toUtf8()); QtSql_PackedString { const_cast<char*>(t5344b7->prepend("WHITESPACE").constData()+10), t5344b7->size()-10, t5344b7 }; });
	}
}

void QSqlTableModel_SetEditStrategy(void* ptr, long long strategy)
{
	static_cast<QSqlTableModel*>(ptr)->setEditStrategy(static_cast<QSqlTableModel::EditStrategy>(strategy));
}

void QSqlTableModel_SetEditStrategyDefault(void* ptr, long long strategy)
{
	if (dynamic_cast<QSqlRelationalTableModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::setEditStrategy(static_cast<QSqlTableModel::EditStrategy>(strategy));
	} else {
		static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::setEditStrategy(static_cast<QSqlTableModel::EditStrategy>(strategy));
	}
}

void QSqlTableModel_SetFilter(void* ptr, struct QtSql_PackedString filter)
{
	static_cast<QSqlTableModel*>(ptr)->setFilter(QString::fromUtf8(filter.data, filter.len));
}

void QSqlTableModel_SetFilterDefault(void* ptr, struct QtSql_PackedString filter)
{
	if (dynamic_cast<QSqlRelationalTableModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::setFilter(QString::fromUtf8(filter.data, filter.len));
	} else {
		static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::setFilter(QString::fromUtf8(filter.data, filter.len));
	}
}

void QSqlTableModel_SetPrimaryKey(void* ptr, void* key)
{
	static_cast<QSqlTableModel*>(ptr)->setPrimaryKey(*static_cast<QSqlIndex*>(key));
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
	if (dynamic_cast<QSqlRelationalTableModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::setSort(column, static_cast<Qt::SortOrder>(order));
	} else {
		static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::setSort(column, static_cast<Qt::SortOrder>(order));
	}
}

void QSqlTableModel_SetTable(void* ptr, struct QtSql_PackedString tableName)
{
	static_cast<QSqlTableModel*>(ptr)->setTable(QString::fromUtf8(tableName.data, tableName.len));
}

void QSqlTableModel_SetTableDefault(void* ptr, struct QtSql_PackedString tableName)
{
	if (dynamic_cast<QSqlRelationalTableModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::setTable(QString::fromUtf8(tableName.data, tableName.len));
	} else {
		static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::setTable(QString::fromUtf8(tableName.data, tableName.len));
	}
}

char QSqlTableModel_SubmitAll(void* ptr)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QSqlTableModel*>(ptr), "submitAll", Q_RETURN_ARG(bool, returnArg));
	return returnArg;
}

char QSqlTableModel_SubmitAllDefault(void* ptr)
{
	if (dynamic_cast<QSqlRelationalTableModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::submitAll();
	} else {
		return static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::submitAll();
	}
}

struct QtSql_PackedString QSqlTableModel_TableName(void* ptr)
{
	return ({ QByteArray* tefb63d = new QByteArray(static_cast<QSqlTableModel*>(ptr)->tableName().toUtf8()); QtSql_PackedString { const_cast<char*>(tefb63d->prepend("WHITESPACE").constData()+10), tefb63d->size()-10, tefb63d }; });
}

char QSqlTableModel_UpdateRowInTable(void* ptr, int row, void* values)
{
	return static_cast<QSqlTableModel*>(ptr)->updateRowInTable(row, *static_cast<QSqlRecord*>(values));
}

char QSqlTableModel_UpdateRowInTableDefault(void* ptr, int row, void* values)
{
	if (dynamic_cast<QSqlRelationalTableModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSqlRelationalTableModel*>(ptr)->QSqlRelationalTableModel::updateRowInTable(row, *static_cast<QSqlRecord*>(values));
	} else {
		return static_cast<QSqlTableModel*>(ptr)->QSqlTableModel::updateRowInTable(row, *static_cast<QSqlRecord*>(values));
	}
}

void QSqlTableModel_DestroyQSqlTableModel(void* ptr)
{
	static_cast<QSqlTableModel*>(ptr)->~QSqlTableModel();
}

void QSqlTableModel_DestroyQSqlTableModelDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

