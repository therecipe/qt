// +build !minimal

package sql

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/internal"
	"github.com/therecipe/qt/widgets"
	"strings"
	"unsafe"
)

func unpackStringList(s string) []string {
	if len(s) == 0 {
		return make([]string, 0)
	}
	return strings.Split(s, "¡¦!")
}

type QSql struct {
	internal.Internal
}

type QSql_ITF interface {
	QSql_PTR() *QSql
}

func (ptr *QSql) QSql_PTR() *QSql {
	return ptr
}

func (ptr *QSql) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QSql) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQSql(ptr QSql_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSql_PTR().Pointer()
	}
	return nil
}

func (n *QSql) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQSqlFromPointer(ptr unsafe.Pointer) (n *QSql) {
	n = new(QSql)
	n.InitFromInternal(uintptr(ptr), "sql.QSql")
	return
}

func (ptr *QSql) DestroyQSql() {
}

//go:generate stringer -type=QSql__Location
//QSql::Location
type QSql__Location int64

const (
	QSql__BeforeFirstRow QSql__Location = QSql__Location(-1)
	QSql__AfterLastRow   QSql__Location = QSql__Location(-2)
)

//go:generate stringer -type=QSql__NumericalPrecisionPolicy
//QSql::NumericalPrecisionPolicy
type QSql__NumericalPrecisionPolicy int64

const (
	QSql__LowPrecisionInt32  QSql__NumericalPrecisionPolicy = QSql__NumericalPrecisionPolicy(0x01)
	QSql__LowPrecisionInt64  QSql__NumericalPrecisionPolicy = QSql__NumericalPrecisionPolicy(0x02)
	QSql__LowPrecisionDouble QSql__NumericalPrecisionPolicy = QSql__NumericalPrecisionPolicy(0x04)
	QSql__HighPrecision      QSql__NumericalPrecisionPolicy = QSql__NumericalPrecisionPolicy(0)
)

//go:generate stringer -type=QSql__ParamTypeFlag
//QSql::ParamTypeFlag
type QSql__ParamTypeFlag int64

const (
	QSql__In     QSql__ParamTypeFlag = QSql__ParamTypeFlag(0x00000001)
	QSql__Out    QSql__ParamTypeFlag = QSql__ParamTypeFlag(0x00000002)
	QSql__InOut  QSql__ParamTypeFlag = QSql__ParamTypeFlag(QSql__In | QSql__Out)
	QSql__Binary QSql__ParamTypeFlag = QSql__ParamTypeFlag(0x00000004)
)

//go:generate stringer -type=QSql__TableType
//QSql::TableType
type QSql__TableType int64

const (
	QSql__Tables       QSql__TableType = QSql__TableType(0x01)
	QSql__SystemTables QSql__TableType = QSql__TableType(0x02)
	QSql__Views        QSql__TableType = QSql__TableType(0x04)
	QSql__AllTables    QSql__TableType = QSql__TableType(0xff)
)

type QSqlDatabase struct {
	internal.Internal
}

type QSqlDatabase_ITF interface {
	QSqlDatabase_PTR() *QSqlDatabase
}

func (ptr *QSqlDatabase) QSqlDatabase_PTR() *QSqlDatabase {
	return ptr
}

func (ptr *QSqlDatabase) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QSqlDatabase) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQSqlDatabase(ptr QSqlDatabase_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlDatabase_PTR().Pointer()
	}
	return nil
}

func (n *QSqlDatabase) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQSqlDatabaseFromPointer(ptr unsafe.Pointer) (n *QSqlDatabase) {
	n = new(QSqlDatabase)
	n.InitFromInternal(uintptr(ptr), "sql.QSqlDatabase")
	return
}
func NewQSqlDatabase() *QSqlDatabase {

	return internal.CallLocalFunction([]interface{}{"", "", "sql.NewQSqlDatabase", ""}).(*QSqlDatabase)
}

func NewQSqlDatabase2(other QSqlDatabase_ITF) *QSqlDatabase {

	return internal.CallLocalFunction([]interface{}{"", "", "sql.NewQSqlDatabase2", "", other}).(*QSqlDatabase)
}

func NewQSqlDatabase3(ty string) *QSqlDatabase {

	return internal.CallLocalFunction([]interface{}{"", "", "sql.NewQSqlDatabase3", "", ty}).(*QSqlDatabase)
}

func NewQSqlDatabase4(driver QSqlDriver_ITF) *QSqlDatabase {

	return internal.CallLocalFunction([]interface{}{"", "", "sql.NewQSqlDatabase4", "", driver}).(*QSqlDatabase)
}

func QSqlDatabase_AddDatabase(ty string, connectionName string) *QSqlDatabase {

	return internal.CallLocalFunction([]interface{}{"", "", "sql.QSqlDatabase_AddDatabase", "", ty, connectionName}).(*QSqlDatabase)
}

func (ptr *QSqlDatabase) AddDatabase(ty string, connectionName string) *QSqlDatabase {

	return internal.CallLocalFunction([]interface{}{"", "", "sql.QSqlDatabase_AddDatabase", "", ty, connectionName}).(*QSqlDatabase)
}

func QSqlDatabase_AddDatabase2(driver QSqlDriver_ITF, connectionName string) *QSqlDatabase {

	return internal.CallLocalFunction([]interface{}{"", "", "sql.QSqlDatabase_AddDatabase2", "", driver, connectionName}).(*QSqlDatabase)
}

func (ptr *QSqlDatabase) AddDatabase2(driver QSqlDriver_ITF, connectionName string) *QSqlDatabase {

	return internal.CallLocalFunction([]interface{}{"", "", "sql.QSqlDatabase_AddDatabase2", "", driver, connectionName}).(*QSqlDatabase)
}

func QSqlDatabase_CloneDatabase(other QSqlDatabase_ITF, connectionName string) *QSqlDatabase {

	return internal.CallLocalFunction([]interface{}{"", "", "sql.QSqlDatabase_CloneDatabase", "", other, connectionName}).(*QSqlDatabase)
}

func (ptr *QSqlDatabase) CloneDatabase(other QSqlDatabase_ITF, connectionName string) *QSqlDatabase {

	return internal.CallLocalFunction([]interface{}{"", "", "sql.QSqlDatabase_CloneDatabase", "", other, connectionName}).(*QSqlDatabase)
}

func QSqlDatabase_CloneDatabase2(other string, connectionName string) *QSqlDatabase {

	return internal.CallLocalFunction([]interface{}{"", "", "sql.QSqlDatabase_CloneDatabase2", "", other, connectionName}).(*QSqlDatabase)
}

func (ptr *QSqlDatabase) CloneDatabase2(other string, connectionName string) *QSqlDatabase {

	return internal.CallLocalFunction([]interface{}{"", "", "sql.QSqlDatabase_CloneDatabase2", "", other, connectionName}).(*QSqlDatabase)
}

func (ptr *QSqlDatabase) Close() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Close"})
}

func (ptr *QSqlDatabase) Commit() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Commit"}).(bool)
}

func (ptr *QSqlDatabase) ConnectOptions() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectOptions"}).(string)
}

func (ptr *QSqlDatabase) ConnectionName() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectionName"}).(string)
}

func QSqlDatabase_ConnectionNames() []string {

	return internal.CallLocalFunction([]interface{}{"", "", "sql.QSqlDatabase_ConnectionNames", ""}).([]string)
}

func (ptr *QSqlDatabase) ConnectionNames() []string {

	return internal.CallLocalFunction([]interface{}{"", "", "sql.QSqlDatabase_ConnectionNames", ""}).([]string)
}

func QSqlDatabase_Contains(connectionName string) bool {

	return internal.CallLocalFunction([]interface{}{"", "", "sql.QSqlDatabase_Contains", "", connectionName}).(bool)
}

func (ptr *QSqlDatabase) Contains(connectionName string) bool {

	return internal.CallLocalFunction([]interface{}{"", "", "sql.QSqlDatabase_Contains", "", connectionName}).(bool)
}

func QSqlDatabase_Database(connectionName string, open bool) *QSqlDatabase {

	return internal.CallLocalFunction([]interface{}{"", "", "sql.QSqlDatabase_Database", "", connectionName, open}).(*QSqlDatabase)
}

func (ptr *QSqlDatabase) Database(connectionName string, open bool) *QSqlDatabase {

	return internal.CallLocalFunction([]interface{}{"", "", "sql.QSqlDatabase_Database", "", connectionName, open}).(*QSqlDatabase)
}

func (ptr *QSqlDatabase) DatabaseName() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DatabaseName"}).(string)
}

func (ptr *QSqlDatabase) Driver() *QSqlDriver {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Driver"}).(*QSqlDriver)
}

func (ptr *QSqlDatabase) DriverName() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DriverName"}).(string)
}

func QSqlDatabase_Drivers() []string {

	return internal.CallLocalFunction([]interface{}{"", "", "sql.QSqlDatabase_Drivers", ""}).([]string)
}

func (ptr *QSqlDatabase) Drivers() []string {

	return internal.CallLocalFunction([]interface{}{"", "", "sql.QSqlDatabase_Drivers", ""}).([]string)
}

func (ptr *QSqlDatabase) Exec(query string) *QSqlQuery {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Exec", query}).(*QSqlQuery)
}

func (ptr *QSqlDatabase) HostName() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HostName"}).(string)
}

func QSqlDatabase_IsDriverAvailable(name string) bool {

	return internal.CallLocalFunction([]interface{}{"", "", "sql.QSqlDatabase_IsDriverAvailable", "", name}).(bool)
}

func (ptr *QSqlDatabase) IsDriverAvailable(name string) bool {

	return internal.CallLocalFunction([]interface{}{"", "", "sql.QSqlDatabase_IsDriverAvailable", "", name}).(bool)
}

func (ptr *QSqlDatabase) IsOpen() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsOpen"}).(bool)
}

func (ptr *QSqlDatabase) IsOpenError() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsOpenError"}).(bool)
}

func (ptr *QSqlDatabase) IsValid() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsValid"}).(bool)
}

func (ptr *QSqlDatabase) LastError() *QSqlError {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LastError"}).(*QSqlError)
}

func (ptr *QSqlDatabase) NumericalPrecisionPolicy() QSql__NumericalPrecisionPolicy {

	return QSql__NumericalPrecisionPolicy(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NumericalPrecisionPolicy"}).(float64))
}

func (ptr *QSqlDatabase) Open() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Open"}).(bool)
}

func (ptr *QSqlDatabase) Open2(user string, password string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Open2", user, password}).(bool)
}

func (ptr *QSqlDatabase) Password() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Password"}).(string)
}

func (ptr *QSqlDatabase) Port() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Port"}).(float64))
}

func (ptr *QSqlDatabase) PrimaryIndex(tablename string) *QSqlIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PrimaryIndex", tablename}).(*QSqlIndex)
}

func (ptr *QSqlDatabase) Record(tablename string) *QSqlRecord {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Record", tablename}).(*QSqlRecord)
}

func QSqlDatabase_RegisterSqlDriver(name string, creator QSqlDriverCreatorBase_ITF) {

	internal.CallLocalFunction([]interface{}{"", "", "sql.QSqlDatabase_RegisterSqlDriver", "", name, creator})
}

func (ptr *QSqlDatabase) RegisterSqlDriver(name string, creator QSqlDriverCreatorBase_ITF) {

	internal.CallLocalFunction([]interface{}{"", "", "sql.QSqlDatabase_RegisterSqlDriver", "", name, creator})
}

func QSqlDatabase_RemoveDatabase(connectionName string) {

	internal.CallLocalFunction([]interface{}{"", "", "sql.QSqlDatabase_RemoveDatabase", "", connectionName})
}

func (ptr *QSqlDatabase) RemoveDatabase(connectionName string) {

	internal.CallLocalFunction([]interface{}{"", "", "sql.QSqlDatabase_RemoveDatabase", "", connectionName})
}

func (ptr *QSqlDatabase) Rollback() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Rollback"}).(bool)
}

func (ptr *QSqlDatabase) SetConnectOptions(options string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetConnectOptions", options})
}

func (ptr *QSqlDatabase) SetDatabaseName(name string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDatabaseName", name})
}

func (ptr *QSqlDatabase) SetHostName(host string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetHostName", host})
}

func (ptr *QSqlDatabase) SetNumericalPrecisionPolicy(precisionPolicy QSql__NumericalPrecisionPolicy) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetNumericalPrecisionPolicy", precisionPolicy})
}

func (ptr *QSqlDatabase) SetPassword(password string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPassword", password})
}

func (ptr *QSqlDatabase) SetPort(port int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPort", port})
}

func (ptr *QSqlDatabase) SetUserName(name string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetUserName", name})
}

func (ptr *QSqlDatabase) Tables(ty QSql__TableType) []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Tables", ty}).([]string)
}

func (ptr *QSqlDatabase) Transaction() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Transaction"}).(bool)
}

func (ptr *QSqlDatabase) UserName() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UserName"}).(string)
}

func (ptr *QSqlDatabase) DestroyQSqlDatabase() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSqlDatabase"})
}

type QSqlDriver struct {
	core.QObject
}

type QSqlDriver_ITF interface {
	core.QObject_ITF
	QSqlDriver_PTR() *QSqlDriver
}

func (ptr *QSqlDriver) QSqlDriver_PTR() *QSqlDriver {
	return ptr
}

func (ptr *QSqlDriver) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QSqlDriver) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQSqlDriver(ptr QSqlDriver_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlDriver_PTR().Pointer()
	}
	return nil
}

func (n *QSqlDriver) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QSqlDriver) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQSqlDriverFromPointer(ptr unsafe.Pointer) (n *QSqlDriver) {
	n = new(QSqlDriver)
	n.InitFromInternal(uintptr(ptr), "sql.QSqlDriver")
	return
}

//go:generate stringer -type=QSqlDriver__DriverFeature
//QSqlDriver::DriverFeature
type QSqlDriver__DriverFeature int64

const (
	QSqlDriver__Transactions           QSqlDriver__DriverFeature = QSqlDriver__DriverFeature(0)
	QSqlDriver__QuerySize              QSqlDriver__DriverFeature = QSqlDriver__DriverFeature(1)
	QSqlDriver__BLOB                   QSqlDriver__DriverFeature = QSqlDriver__DriverFeature(2)
	QSqlDriver__Unicode                QSqlDriver__DriverFeature = QSqlDriver__DriverFeature(3)
	QSqlDriver__PreparedQueries        QSqlDriver__DriverFeature = QSqlDriver__DriverFeature(4)
	QSqlDriver__NamedPlaceholders      QSqlDriver__DriverFeature = QSqlDriver__DriverFeature(5)
	QSqlDriver__PositionalPlaceholders QSqlDriver__DriverFeature = QSqlDriver__DriverFeature(6)
	QSqlDriver__LastInsertId           QSqlDriver__DriverFeature = QSqlDriver__DriverFeature(7)
	QSqlDriver__BatchOperations        QSqlDriver__DriverFeature = QSqlDriver__DriverFeature(8)
	QSqlDriver__SimpleLocking          QSqlDriver__DriverFeature = QSqlDriver__DriverFeature(9)
	QSqlDriver__LowPrecisionNumbers    QSqlDriver__DriverFeature = QSqlDriver__DriverFeature(10)
	QSqlDriver__EventNotifications     QSqlDriver__DriverFeature = QSqlDriver__DriverFeature(11)
	QSqlDriver__FinishQuery            QSqlDriver__DriverFeature = QSqlDriver__DriverFeature(12)
	QSqlDriver__MultipleResultSets     QSqlDriver__DriverFeature = QSqlDriver__DriverFeature(13)
	QSqlDriver__CancelQuery            QSqlDriver__DriverFeature = QSqlDriver__DriverFeature(14)
)

//go:generate stringer -type=QSqlDriver__StatementType
//QSqlDriver::StatementType
type QSqlDriver__StatementType int64

const (
	QSqlDriver__WhereStatement  QSqlDriver__StatementType = QSqlDriver__StatementType(0)
	QSqlDriver__SelectStatement QSqlDriver__StatementType = QSqlDriver__StatementType(1)
	QSqlDriver__UpdateStatement QSqlDriver__StatementType = QSqlDriver__StatementType(2)
	QSqlDriver__InsertStatement QSqlDriver__StatementType = QSqlDriver__StatementType(3)
	QSqlDriver__DeleteStatement QSqlDriver__StatementType = QSqlDriver__StatementType(4)
)

//go:generate stringer -type=QSqlDriver__IdentifierType
//QSqlDriver::IdentifierType
type QSqlDriver__IdentifierType int64

const (
	QSqlDriver__FieldName QSqlDriver__IdentifierType = QSqlDriver__IdentifierType(0)
	QSqlDriver__TableName QSqlDriver__IdentifierType = QSqlDriver__IdentifierType(1)
)

//go:generate stringer -type=QSqlDriver__NotificationSource
//QSqlDriver::NotificationSource
type QSqlDriver__NotificationSource int64

const (
	QSqlDriver__UnknownSource QSqlDriver__NotificationSource = QSqlDriver__NotificationSource(0)
	QSqlDriver__SelfSource    QSqlDriver__NotificationSource = QSqlDriver__NotificationSource(1)
	QSqlDriver__OtherSource   QSqlDriver__NotificationSource = QSqlDriver__NotificationSource(2)
)

func NewQSqlDriver(parent core.QObject_ITF) *QSqlDriver {

	return internal.CallLocalFunction([]interface{}{"", "", "sql.NewQSqlDriver", "", parent}).(*QSqlDriver)
}

func (ptr *QSqlDriver) ConnectBeginTransaction(f func() bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectBeginTransaction", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlDriver) DisconnectBeginTransaction() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectBeginTransaction"})
}

func (ptr *QSqlDriver) BeginTransaction() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BeginTransaction"}).(bool)
}

func (ptr *QSqlDriver) BeginTransactionDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BeginTransactionDefault"}).(bool)
}

func (ptr *QSqlDriver) ConnectClose(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectClose", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlDriver) DisconnectClose() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectClose"})
}

func (ptr *QSqlDriver) Close() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Close"})
}

func (ptr *QSqlDriver) ConnectCommitTransaction(f func() bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCommitTransaction", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlDriver) DisconnectCommitTransaction() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCommitTransaction"})
}

func (ptr *QSqlDriver) CommitTransaction() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CommitTransaction"}).(bool)
}

func (ptr *QSqlDriver) CommitTransactionDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CommitTransactionDefault"}).(bool)
}

func (ptr *QSqlDriver) ConnectCreateResult(f func() *QSqlResult) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCreateResult", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlDriver) DisconnectCreateResult() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCreateResult"})
}

func (ptr *QSqlDriver) CreateResult() *QSqlResult {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreateResult"}).(*QSqlResult)
}

func (ptr *QSqlDriver) ConnectEscapeIdentifier(f func(identifier string, ty QSqlDriver__IdentifierType) string) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectEscapeIdentifier", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlDriver) DisconnectEscapeIdentifier() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectEscapeIdentifier"})
}

func (ptr *QSqlDriver) EscapeIdentifier(identifier string, ty QSqlDriver__IdentifierType) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EscapeIdentifier", identifier, ty}).(string)
}

func (ptr *QSqlDriver) EscapeIdentifierDefault(identifier string, ty QSqlDriver__IdentifierType) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EscapeIdentifierDefault", identifier, ty}).(string)
}

func (ptr *QSqlDriver) ConnectFormatValue(f func(field *QSqlField, trimStrings bool) string) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFormatValue", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlDriver) DisconnectFormatValue() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFormatValue"})
}

func (ptr *QSqlDriver) FormatValue(field QSqlField_ITF, trimStrings bool) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FormatValue", field, trimStrings}).(string)
}

func (ptr *QSqlDriver) FormatValueDefault(field QSqlField_ITF, trimStrings bool) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FormatValueDefault", field, trimStrings}).(string)
}

func (ptr *QSqlDriver) ConnectHandle(f func() *core.QVariant) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectHandle", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlDriver) DisconnectHandle() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectHandle"})
}

func (ptr *QSqlDriver) Handle() *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Handle"}).(*core.QVariant)
}

func (ptr *QSqlDriver) HandleDefault() *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HandleDefault"}).(*core.QVariant)
}

func (ptr *QSqlDriver) ConnectHasFeature(f func(feature QSqlDriver__DriverFeature) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectHasFeature", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlDriver) DisconnectHasFeature() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectHasFeature"})
}

func (ptr *QSqlDriver) HasFeature(feature QSqlDriver__DriverFeature) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasFeature", feature}).(bool)
}

func (ptr *QSqlDriver) ConnectIsIdentifierEscaped(f func(identifier string, ty QSqlDriver__IdentifierType) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectIsIdentifierEscaped", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlDriver) DisconnectIsIdentifierEscaped() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectIsIdentifierEscaped"})
}

func (ptr *QSqlDriver) IsIdentifierEscaped(identifier string, ty QSqlDriver__IdentifierType) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsIdentifierEscaped", identifier, ty}).(bool)
}

func (ptr *QSqlDriver) IsIdentifierEscapedDefault(identifier string, ty QSqlDriver__IdentifierType) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsIdentifierEscapedDefault", identifier, ty}).(bool)
}

func (ptr *QSqlDriver) ConnectIsOpen(f func() bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectIsOpen", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlDriver) DisconnectIsOpen() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectIsOpen"})
}

func (ptr *QSqlDriver) IsOpen() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsOpen"}).(bool)
}

func (ptr *QSqlDriver) IsOpenDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsOpenDefault"}).(bool)
}

func (ptr *QSqlDriver) IsOpenError() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsOpenError"}).(bool)
}

func (ptr *QSqlDriver) LastError() *QSqlError {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LastError"}).(*QSqlError)
}

func (ptr *QSqlDriver) ConnectNotification(f func(name string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotification", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlDriver) DisconnectNotification() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotification"})
}

func (ptr *QSqlDriver) Notification(name string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Notification", name})
}

func (ptr *QSqlDriver) ConnectNotification2(f func(name string, source QSqlDriver__NotificationSource, payload *core.QVariant)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotification2", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlDriver) DisconnectNotification2() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotification2"})
}

func (ptr *QSqlDriver) Notification2(name string, source QSqlDriver__NotificationSource, payload core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Notification2", name, source, payload})
}

func (ptr *QSqlDriver) NumericalPrecisionPolicy() QSql__NumericalPrecisionPolicy {

	return QSql__NumericalPrecisionPolicy(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NumericalPrecisionPolicy"}).(float64))
}

func (ptr *QSqlDriver) ConnectOpen(f func(db string, user string, password string, host string, port int, options string) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectOpen", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlDriver) DisconnectOpen() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectOpen"})
}

func (ptr *QSqlDriver) Open(db string, user string, password string, host string, port int, options string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Open", db, user, password, host, port, options}).(bool)
}

func (ptr *QSqlDriver) ConnectPrimaryIndex(f func(tableName string) *QSqlIndex) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPrimaryIndex", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlDriver) DisconnectPrimaryIndex() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPrimaryIndex"})
}

func (ptr *QSqlDriver) PrimaryIndex(tableName string) *QSqlIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PrimaryIndex", tableName}).(*QSqlIndex)
}

func (ptr *QSqlDriver) PrimaryIndexDefault(tableName string) *QSqlIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PrimaryIndexDefault", tableName}).(*QSqlIndex)
}

func (ptr *QSqlDriver) ConnectRecord(f func(tableName string) *QSqlRecord) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRecord", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlDriver) DisconnectRecord() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRecord"})
}

func (ptr *QSqlDriver) Record(tableName string) *QSqlRecord {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Record", tableName}).(*QSqlRecord)
}

func (ptr *QSqlDriver) RecordDefault(tableName string) *QSqlRecord {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RecordDefault", tableName}).(*QSqlRecord)
}

func (ptr *QSqlDriver) ConnectRollbackTransaction(f func() bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRollbackTransaction", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlDriver) DisconnectRollbackTransaction() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRollbackTransaction"})
}

func (ptr *QSqlDriver) RollbackTransaction() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RollbackTransaction"}).(bool)
}

func (ptr *QSqlDriver) RollbackTransactionDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RollbackTransactionDefault"}).(bool)
}

func (ptr *QSqlDriver) ConnectSetLastError(f func(error *QSqlError)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetLastError", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlDriver) DisconnectSetLastError() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetLastError"})
}

func (ptr *QSqlDriver) SetLastError(error QSqlError_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLastError", error})
}

func (ptr *QSqlDriver) SetLastErrorDefault(error QSqlError_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLastErrorDefault", error})
}

func (ptr *QSqlDriver) SetNumericalPrecisionPolicy(precisionPolicy QSql__NumericalPrecisionPolicy) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetNumericalPrecisionPolicy", precisionPolicy})
}

func (ptr *QSqlDriver) ConnectSetOpen(f func(open bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetOpen", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlDriver) DisconnectSetOpen() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetOpen"})
}

func (ptr *QSqlDriver) SetOpen(open bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetOpen", open})
}

func (ptr *QSqlDriver) SetOpenDefault(open bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetOpenDefault", open})
}

func (ptr *QSqlDriver) ConnectSetOpenError(f func(error bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetOpenError", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlDriver) DisconnectSetOpenError() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetOpenError"})
}

func (ptr *QSqlDriver) SetOpenError(error bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetOpenError", error})
}

func (ptr *QSqlDriver) SetOpenErrorDefault(error bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetOpenErrorDefault", error})
}

func (ptr *QSqlDriver) ConnectSqlStatement(f func(ty QSqlDriver__StatementType, tableName string, rec *QSqlRecord, preparedStatement bool) string) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSqlStatement", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlDriver) DisconnectSqlStatement() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSqlStatement"})
}

func (ptr *QSqlDriver) SqlStatement(ty QSqlDriver__StatementType, tableName string, rec QSqlRecord_ITF, preparedStatement bool) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SqlStatement", ty, tableName, rec, preparedStatement}).(string)
}

func (ptr *QSqlDriver) SqlStatementDefault(ty QSqlDriver__StatementType, tableName string, rec QSqlRecord_ITF, preparedStatement bool) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SqlStatementDefault", ty, tableName, rec, preparedStatement}).(string)
}

func (ptr *QSqlDriver) ConnectStripDelimiters(f func(identifier string, ty QSqlDriver__IdentifierType) string) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectStripDelimiters", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlDriver) DisconnectStripDelimiters() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectStripDelimiters"})
}

func (ptr *QSqlDriver) StripDelimiters(identifier string, ty QSqlDriver__IdentifierType) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StripDelimiters", identifier, ty}).(string)
}

func (ptr *QSqlDriver) StripDelimitersDefault(identifier string, ty QSqlDriver__IdentifierType) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StripDelimitersDefault", identifier, ty}).(string)
}

func (ptr *QSqlDriver) ConnectSubscribeToNotification(f func(name string) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSubscribeToNotification", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlDriver) DisconnectSubscribeToNotification() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSubscribeToNotification"})
}

func (ptr *QSqlDriver) SubscribeToNotification(name string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SubscribeToNotification", name}).(bool)
}

func (ptr *QSqlDriver) SubscribeToNotificationDefault(name string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SubscribeToNotificationDefault", name}).(bool)
}

func (ptr *QSqlDriver) ConnectSubscribedToNotifications(f func() []string) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSubscribedToNotifications", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlDriver) DisconnectSubscribedToNotifications() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSubscribedToNotifications"})
}

func (ptr *QSqlDriver) SubscribedToNotifications() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SubscribedToNotifications"}).([]string)
}

func (ptr *QSqlDriver) SubscribedToNotificationsDefault() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SubscribedToNotificationsDefault"}).([]string)
}

func (ptr *QSqlDriver) ConnectTables(f func(tableType QSql__TableType) []string) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectTables", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlDriver) DisconnectTables() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectTables"})
}

func (ptr *QSqlDriver) Tables(tableType QSql__TableType) []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Tables", tableType}).([]string)
}

func (ptr *QSqlDriver) TablesDefault(tableType QSql__TableType) []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TablesDefault", tableType}).([]string)
}

func (ptr *QSqlDriver) ConnectUnsubscribeFromNotification(f func(name string) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectUnsubscribeFromNotification", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlDriver) DisconnectUnsubscribeFromNotification() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectUnsubscribeFromNotification"})
}

func (ptr *QSqlDriver) UnsubscribeFromNotification(name string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UnsubscribeFromNotification", name}).(bool)
}

func (ptr *QSqlDriver) UnsubscribeFromNotificationDefault(name string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UnsubscribeFromNotificationDefault", name}).(bool)
}

func (ptr *QSqlDriver) ConnectDestroyQSqlDriver(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQSqlDriver", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlDriver) DisconnectDestroyQSqlDriver() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQSqlDriver"})
}

func (ptr *QSqlDriver) DestroyQSqlDriver() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSqlDriver"})
}

func (ptr *QSqlDriver) DestroyQSqlDriverDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSqlDriverDefault"})
}

func (ptr *QSqlDriver) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QSqlDriver) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QSqlDriver) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QSqlDriver) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QSqlDriver) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QSqlDriver) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QSqlDriver) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QSqlDriver) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QSqlDriver) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QSqlDriver) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QSqlDriver) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QSqlDriver) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QSqlDriver) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QSqlDriver) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QSqlDriver) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QSqlDriver) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QSqlDriver) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QSqlDriver) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QSqlDriver) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QSqlDriver) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QSqlDriver) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QSqlDriverCreator struct {
	QSqlDriverCreatorBase
}

type QSqlDriverCreator_ITF interface {
	QSqlDriverCreatorBase_ITF
	QSqlDriverCreator_PTR() *QSqlDriverCreator
}

func (ptr *QSqlDriverCreator) QSqlDriverCreator_PTR() *QSqlDriverCreator {
	return ptr
}

func (ptr *QSqlDriverCreator) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlDriverCreatorBase_PTR().Pointer()
	}
	return nil
}

func (ptr *QSqlDriverCreator) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSqlDriverCreatorBase_PTR().SetPointer(p)
	}
}

func PointerFromQSqlDriverCreator(ptr QSqlDriverCreator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlDriverCreator_PTR().Pointer()
	}
	return nil
}

func (n *QSqlDriverCreator) InitFromInternal(ptr uintptr, name string) {
	n.QSqlDriverCreatorBase_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QSqlDriverCreator) ClassNameInternalF() string {
	return n.QSqlDriverCreatorBase_PTR().ClassNameInternalF()
}

func NewQSqlDriverCreatorFromPointer(ptr unsafe.Pointer) (n *QSqlDriverCreator) {
	n = new(QSqlDriverCreator)
	n.InitFromInternal(uintptr(ptr), "sql.QSqlDriverCreator")
	return
}

func (ptr *QSqlDriverCreator) DestroyQSqlDriverCreator() {
}

type QSqlDriverCreatorBase struct {
	internal.Internal
}

type QSqlDriverCreatorBase_ITF interface {
	QSqlDriverCreatorBase_PTR() *QSqlDriverCreatorBase
}

func (ptr *QSqlDriverCreatorBase) QSqlDriverCreatorBase_PTR() *QSqlDriverCreatorBase {
	return ptr
}

func (ptr *QSqlDriverCreatorBase) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QSqlDriverCreatorBase) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQSqlDriverCreatorBase(ptr QSqlDriverCreatorBase_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlDriverCreatorBase_PTR().Pointer()
	}
	return nil
}

func (n *QSqlDriverCreatorBase) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQSqlDriverCreatorBaseFromPointer(ptr unsafe.Pointer) (n *QSqlDriverCreatorBase) {
	n = new(QSqlDriverCreatorBase)
	n.InitFromInternal(uintptr(ptr), "sql.QSqlDriverCreatorBase")
	return
}
func (ptr *QSqlDriverCreatorBase) ConnectCreateObject(f func() *QSqlDriver) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCreateObject", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlDriverCreatorBase) DisconnectCreateObject() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCreateObject"})
}

func (ptr *QSqlDriverCreatorBase) CreateObject() *QSqlDriver {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreateObject"}).(*QSqlDriver)
}

func (ptr *QSqlDriverCreatorBase) ConnectDestroyQSqlDriverCreatorBase(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQSqlDriverCreatorBase", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlDriverCreatorBase) DisconnectDestroyQSqlDriverCreatorBase() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQSqlDriverCreatorBase"})
}

func (ptr *QSqlDriverCreatorBase) DestroyQSqlDriverCreatorBase() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSqlDriverCreatorBase"})
}

func (ptr *QSqlDriverCreatorBase) DestroyQSqlDriverCreatorBaseDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSqlDriverCreatorBaseDefault"})
}

type QSqlDriverPlugin struct {
	core.QObject
}

type QSqlDriverPlugin_ITF interface {
	core.QObject_ITF
	QSqlDriverPlugin_PTR() *QSqlDriverPlugin
}

func (ptr *QSqlDriverPlugin) QSqlDriverPlugin_PTR() *QSqlDriverPlugin {
	return ptr
}

func (ptr *QSqlDriverPlugin) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QSqlDriverPlugin) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQSqlDriverPlugin(ptr QSqlDriverPlugin_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlDriverPlugin_PTR().Pointer()
	}
	return nil
}

func (n *QSqlDriverPlugin) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QSqlDriverPlugin) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQSqlDriverPluginFromPointer(ptr unsafe.Pointer) (n *QSqlDriverPlugin) {
	n = new(QSqlDriverPlugin)
	n.InitFromInternal(uintptr(ptr), "sql.QSqlDriverPlugin")
	return
}
func NewQSqlDriverPlugin(parent core.QObject_ITF) *QSqlDriverPlugin {

	return internal.CallLocalFunction([]interface{}{"", "", "sql.NewQSqlDriverPlugin", "", parent}).(*QSqlDriverPlugin)
}

func (ptr *QSqlDriverPlugin) ConnectCreate(f func(key string) *QSqlDriver) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCreate", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlDriverPlugin) DisconnectCreate() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCreate"})
}

func (ptr *QSqlDriverPlugin) Create(key string) *QSqlDriver {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Create", key}).(*QSqlDriver)
}

func (ptr *QSqlDriverPlugin) ConnectDestroyQSqlDriverPlugin(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQSqlDriverPlugin", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlDriverPlugin) DisconnectDestroyQSqlDriverPlugin() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQSqlDriverPlugin"})
}

func (ptr *QSqlDriverPlugin) DestroyQSqlDriverPlugin() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSqlDriverPlugin"})
}

func (ptr *QSqlDriverPlugin) DestroyQSqlDriverPluginDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSqlDriverPluginDefault"})
}

func (ptr *QSqlDriverPlugin) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QSqlDriverPlugin) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QSqlDriverPlugin) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QSqlDriverPlugin) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QSqlDriverPlugin) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QSqlDriverPlugin) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QSqlDriverPlugin) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QSqlDriverPlugin) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QSqlDriverPlugin) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QSqlDriverPlugin) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QSqlDriverPlugin) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QSqlDriverPlugin) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QSqlDriverPlugin) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QSqlDriverPlugin) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QSqlDriverPlugin) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QSqlDriverPlugin) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QSqlDriverPlugin) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QSqlDriverPlugin) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QSqlDriverPlugin) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QSqlDriverPlugin) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QSqlDriverPlugin) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QSqlError struct {
	internal.Internal
}

type QSqlError_ITF interface {
	QSqlError_PTR() *QSqlError
}

func (ptr *QSqlError) QSqlError_PTR() *QSqlError {
	return ptr
}

func (ptr *QSqlError) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QSqlError) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQSqlError(ptr QSqlError_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlError_PTR().Pointer()
	}
	return nil
}

func (n *QSqlError) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQSqlErrorFromPointer(ptr unsafe.Pointer) (n *QSqlError) {
	n = new(QSqlError)
	n.InitFromInternal(uintptr(ptr), "sql.QSqlError")
	return
}

//go:generate stringer -type=QSqlError__ErrorType
//QSqlError::ErrorType
type QSqlError__ErrorType int64

const (
	QSqlError__NoError          QSqlError__ErrorType = QSqlError__ErrorType(0)
	QSqlError__ConnectionError  QSqlError__ErrorType = QSqlError__ErrorType(1)
	QSqlError__StatementError   QSqlError__ErrorType = QSqlError__ErrorType(2)
	QSqlError__TransactionError QSqlError__ErrorType = QSqlError__ErrorType(3)
	QSqlError__UnknownError     QSqlError__ErrorType = QSqlError__ErrorType(4)
)

func NewQSqlError2(driverText string, databaseText string, ty QSqlError__ErrorType, code string) *QSqlError {

	return internal.CallLocalFunction([]interface{}{"", "", "sql.NewQSqlError2", "", driverText, databaseText, ty, code}).(*QSqlError)
}

func NewQSqlError3(other QSqlError_ITF) *QSqlError {

	return internal.CallLocalFunction([]interface{}{"", "", "sql.NewQSqlError3", "", other}).(*QSqlError)
}

func NewQSqlError4(other QSqlError_ITF) *QSqlError {

	return internal.CallLocalFunction([]interface{}{"", "", "sql.NewQSqlError4", "", other}).(*QSqlError)
}

func (ptr *QSqlError) DatabaseText() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DatabaseText"}).(string)
}

func (ptr *QSqlError) DriverText() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DriverText"}).(string)
}

func (ptr *QSqlError) IsValid() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsValid"}).(bool)
}

func (ptr *QSqlError) NativeErrorCode() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NativeErrorCode"}).(string)
}

func (ptr *QSqlError) Swap(other QSqlError_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Swap", other})
}

func (ptr *QSqlError) Text() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Text"}).(string)
}

func (ptr *QSqlError) Type() QSqlError__ErrorType {

	return QSqlError__ErrorType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Type"}).(float64))
}

func (ptr *QSqlError) DestroyQSqlError() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSqlError"})
}

type QSqlField struct {
	internal.Internal
}

type QSqlField_ITF interface {
	QSqlField_PTR() *QSqlField
}

func (ptr *QSqlField) QSqlField_PTR() *QSqlField {
	return ptr
}

func (ptr *QSqlField) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QSqlField) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQSqlField(ptr QSqlField_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlField_PTR().Pointer()
	}
	return nil
}

func (n *QSqlField) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQSqlFieldFromPointer(ptr unsafe.Pointer) (n *QSqlField) {
	n = new(QSqlField)
	n.InitFromInternal(uintptr(ptr), "sql.QSqlField")
	return
}

//go:generate stringer -type=QSqlField__RequiredStatus
//QSqlField::RequiredStatus
type QSqlField__RequiredStatus int64

const (
	QSqlField__Unknown  QSqlField__RequiredStatus = QSqlField__RequiredStatus(-1)
	QSqlField__Optional QSqlField__RequiredStatus = QSqlField__RequiredStatus(0)
	QSqlField__Required QSqlField__RequiredStatus = QSqlField__RequiredStatus(1)
)

func NewQSqlField(fieldName string, ty core.QVariant__Type) *QSqlField {

	return internal.CallLocalFunction([]interface{}{"", "", "sql.NewQSqlField", "", fieldName, ty}).(*QSqlField)
}

func NewQSqlField2(fieldName string, ty core.QVariant__Type, table string) *QSqlField {

	return internal.CallLocalFunction([]interface{}{"", "", "sql.NewQSqlField2", "", fieldName, ty, table}).(*QSqlField)
}

func NewQSqlField3(other QSqlField_ITF) *QSqlField {

	return internal.CallLocalFunction([]interface{}{"", "", "sql.NewQSqlField3", "", other}).(*QSqlField)
}

func (ptr *QSqlField) Clear() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Clear"})
}

func (ptr *QSqlField) DefaultValue() *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DefaultValue"}).(*core.QVariant)
}

func (ptr *QSqlField) IsAutoValue() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsAutoValue"}).(bool)
}

func (ptr *QSqlField) IsGenerated() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsGenerated"}).(bool)
}

func (ptr *QSqlField) IsNull() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsNull"}).(bool)
}

func (ptr *QSqlField) IsReadOnly() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsReadOnly"}).(bool)
}

func (ptr *QSqlField) IsValid() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsValid"}).(bool)
}

func (ptr *QSqlField) Length() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Length"}).(float64))
}

func (ptr *QSqlField) Name() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Name"}).(string)
}

func (ptr *QSqlField) Precision() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Precision"}).(float64))
}

func (ptr *QSqlField) RequiredStatus() QSqlField__RequiredStatus {

	return QSqlField__RequiredStatus(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RequiredStatus"}).(float64))
}

func (ptr *QSqlField) SetAutoValue(autoVal bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAutoValue", autoVal})
}

func (ptr *QSqlField) SetDefaultValue(value core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDefaultValue", value})
}

func (ptr *QSqlField) SetGenerated(gen bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetGenerated", gen})
}

func (ptr *QSqlField) SetLength(fieldLength int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLength", fieldLength})
}

func (ptr *QSqlField) SetName(name string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetName", name})
}

func (ptr *QSqlField) SetPrecision(precision int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPrecision", precision})
}

func (ptr *QSqlField) SetReadOnly(readOnly bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetReadOnly", readOnly})
}

func (ptr *QSqlField) SetRequired(required bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRequired", required})
}

func (ptr *QSqlField) SetRequiredStatus(required QSqlField__RequiredStatus) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRequiredStatus", required})
}

func (ptr *QSqlField) SetTableName(table string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetTableName", table})
}

func (ptr *QSqlField) SetType(ty core.QVariant__Type) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetType", ty})
}

func (ptr *QSqlField) SetValue(value core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetValue", value})
}

func (ptr *QSqlField) TableName() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TableName"}).(string)
}

func (ptr *QSqlField) Type() core.QVariant__Type {

	return core.QVariant__Type(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Type"}).(float64))
}

func (ptr *QSqlField) Value() *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Value"}).(*core.QVariant)
}

func (ptr *QSqlField) DestroyQSqlField() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSqlField"})
}

type QSqlIndex struct {
	QSqlRecord
}

type QSqlIndex_ITF interface {
	QSqlRecord_ITF
	QSqlIndex_PTR() *QSqlIndex
}

func (ptr *QSqlIndex) QSqlIndex_PTR() *QSqlIndex {
	return ptr
}

func (ptr *QSqlIndex) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlRecord_PTR().Pointer()
	}
	return nil
}

func (ptr *QSqlIndex) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSqlRecord_PTR().SetPointer(p)
	}
}

func PointerFromQSqlIndex(ptr QSqlIndex_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlIndex_PTR().Pointer()
	}
	return nil
}

func (n *QSqlIndex) InitFromInternal(ptr uintptr, name string) {
	n.QSqlRecord_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QSqlIndex) ClassNameInternalF() string {
	return n.QSqlRecord_PTR().ClassNameInternalF()
}

func NewQSqlIndexFromPointer(ptr unsafe.Pointer) (n *QSqlIndex) {
	n = new(QSqlIndex)
	n.InitFromInternal(uintptr(ptr), "sql.QSqlIndex")
	return
}
func NewQSqlIndex(cursorname string, name string) *QSqlIndex {

	return internal.CallLocalFunction([]interface{}{"", "", "sql.NewQSqlIndex", "", cursorname, name}).(*QSqlIndex)
}

func NewQSqlIndex2(other QSqlIndex_ITF) *QSqlIndex {

	return internal.CallLocalFunction([]interface{}{"", "", "sql.NewQSqlIndex2", "", other}).(*QSqlIndex)
}

func (ptr *QSqlIndex) Append2(field QSqlField_ITF, desc bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Append2", field, desc})
}

func (ptr *QSqlIndex) CursorName() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CursorName"}).(string)
}

func (ptr *QSqlIndex) IsDescending(i int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsDescending", i}).(bool)
}

func (ptr *QSqlIndex) Name() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Name"}).(string)
}

func (ptr *QSqlIndex) SetCursorName(cursorName string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetCursorName", cursorName})
}

func (ptr *QSqlIndex) SetDescending(i int, desc bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDescending", i, desc})
}

func (ptr *QSqlIndex) SetName(name string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetName", name})
}

func (ptr *QSqlIndex) DestroyQSqlIndex() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSqlIndex"})
}

func (ptr *QSqlIndex) __sorts_atList(i int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__sorts_atList", i}).(bool)
}

func (ptr *QSqlIndex) __sorts_setList(i bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__sorts_setList", i})
}

func (ptr *QSqlIndex) __sorts_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__sorts_newList"}).(unsafe.Pointer)
}

func (ptr *QSqlIndex) __setSorts__atList(i int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setSorts__atList", i}).(bool)
}

func (ptr *QSqlIndex) __setSorts__setList(i bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setSorts__setList", i})
}

func (ptr *QSqlIndex) __setSorts__newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setSorts__newList"}).(unsafe.Pointer)
}

type QSqlQuery struct {
	internal.Internal
}

type QSqlQuery_ITF interface {
	QSqlQuery_PTR() *QSqlQuery
}

func (ptr *QSqlQuery) QSqlQuery_PTR() *QSqlQuery {
	return ptr
}

func (ptr *QSqlQuery) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QSqlQuery) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQSqlQuery(ptr QSqlQuery_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlQuery_PTR().Pointer()
	}
	return nil
}

func (n *QSqlQuery) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQSqlQueryFromPointer(ptr unsafe.Pointer) (n *QSqlQuery) {
	n = new(QSqlQuery)
	n.InitFromInternal(uintptr(ptr), "sql.QSqlQuery")
	return
}

//go:generate stringer -type=QSqlQuery__BatchExecutionMode
//QSqlQuery::BatchExecutionMode
type QSqlQuery__BatchExecutionMode int64

const (
	QSqlQuery__ValuesAsRows    QSqlQuery__BatchExecutionMode = QSqlQuery__BatchExecutionMode(0)
	QSqlQuery__ValuesAsColumns QSqlQuery__BatchExecutionMode = QSqlQuery__BatchExecutionMode(1)
)

func NewQSqlQuery(result QSqlResult_ITF) *QSqlQuery {

	return internal.CallLocalFunction([]interface{}{"", "", "sql.NewQSqlQuery", "", result}).(*QSqlQuery)
}

func NewQSqlQuery2(query string, db QSqlDatabase_ITF) *QSqlQuery {

	return internal.CallLocalFunction([]interface{}{"", "", "sql.NewQSqlQuery2", "", query, db}).(*QSqlQuery)
}

func NewQSqlQuery3(db QSqlDatabase_ITF) *QSqlQuery {

	return internal.CallLocalFunction([]interface{}{"", "", "sql.NewQSqlQuery3", "", db}).(*QSqlQuery)
}

func NewQSqlQuery4(other QSqlQuery_ITF) *QSqlQuery {

	return internal.CallLocalFunction([]interface{}{"", "", "sql.NewQSqlQuery4", "", other}).(*QSqlQuery)
}

func (ptr *QSqlQuery) AddBindValue(val core.QVariant_ITF, paramType QSql__ParamTypeFlag) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AddBindValue", val, paramType})
}

func (ptr *QSqlQuery) At() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "At"}).(float64))
}

func (ptr *QSqlQuery) BindValue(placeholder string, val core.QVariant_ITF, paramType QSql__ParamTypeFlag) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BindValue", placeholder, val, paramType})
}

func (ptr *QSqlQuery) BindValue2(pos int, val core.QVariant_ITF, paramType QSql__ParamTypeFlag) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BindValue2", pos, val, paramType})
}

func (ptr *QSqlQuery) BoundValue(placeholder string) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BoundValue", placeholder}).(*core.QVariant)
}

func (ptr *QSqlQuery) BoundValue2(pos int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BoundValue2", pos}).(*core.QVariant)
}

func (ptr *QSqlQuery) BoundValues() map[string]*core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BoundValues"}).(map[string]*core.QVariant)
}

func (ptr *QSqlQuery) Clear() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Clear"})
}

func (ptr *QSqlQuery) Driver() *QSqlDriver {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Driver"}).(*QSqlDriver)
}

func (ptr *QSqlQuery) Exec(query string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Exec", query}).(bool)
}

func (ptr *QSqlQuery) Exec2() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Exec2"}).(bool)
}

func (ptr *QSqlQuery) ExecBatch(mode QSqlQuery__BatchExecutionMode) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ExecBatch", mode}).(bool)
}

func (ptr *QSqlQuery) ExecutedQuery() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ExecutedQuery"}).(string)
}

func (ptr *QSqlQuery) Finish() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Finish"})
}

func (ptr *QSqlQuery) First() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "First"}).(bool)
}

func (ptr *QSqlQuery) IsActive() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsActive"}).(bool)
}

func (ptr *QSqlQuery) IsForwardOnly() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsForwardOnly"}).(bool)
}

func (ptr *QSqlQuery) IsNull(field int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsNull", field}).(bool)
}

func (ptr *QSqlQuery) IsNull2(name string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsNull2", name}).(bool)
}

func (ptr *QSqlQuery) IsSelect() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsSelect"}).(bool)
}

func (ptr *QSqlQuery) IsValid() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsValid"}).(bool)
}

func (ptr *QSqlQuery) Last() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Last"}).(bool)
}

func (ptr *QSqlQuery) LastError() *QSqlError {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LastError"}).(*QSqlError)
}

func (ptr *QSqlQuery) LastInsertId() *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LastInsertId"}).(*core.QVariant)
}

func (ptr *QSqlQuery) LastQuery() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LastQuery"}).(string)
}

func (ptr *QSqlQuery) Next() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Next"}).(bool)
}

func (ptr *QSqlQuery) NextResult() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NextResult"}).(bool)
}

func (ptr *QSqlQuery) NumRowsAffected() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NumRowsAffected"}).(float64))
}

func (ptr *QSqlQuery) NumericalPrecisionPolicy() QSql__NumericalPrecisionPolicy {

	return QSql__NumericalPrecisionPolicy(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NumericalPrecisionPolicy"}).(float64))
}

func (ptr *QSqlQuery) Prepare(query string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Prepare", query}).(bool)
}

func (ptr *QSqlQuery) Previous() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Previous"}).(bool)
}

func (ptr *QSqlQuery) Record() *QSqlRecord {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Record"}).(*QSqlRecord)
}

func (ptr *QSqlQuery) Result() *QSqlResult {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Result"}).(*QSqlResult)
}

func (ptr *QSqlQuery) Seek(index int, relative bool) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Seek", index, relative}).(bool)
}

func (ptr *QSqlQuery) SetForwardOnly(forward bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetForwardOnly", forward})
}

func (ptr *QSqlQuery) SetNumericalPrecisionPolicy(precisionPolicy QSql__NumericalPrecisionPolicy) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetNumericalPrecisionPolicy", precisionPolicy})
}

func (ptr *QSqlQuery) Size() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Size"}).(float64))
}

func (ptr *QSqlQuery) Value(index int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Value", index}).(*core.QVariant)
}

func (ptr *QSqlQuery) Value2(name string) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Value2", name}).(*core.QVariant)
}

func (ptr *QSqlQuery) DestroyQSqlQuery() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSqlQuery"})
}

func (ptr *QSqlQuery) __boundValues_atList(v string, i int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__boundValues_atList", v, i}).(*core.QVariant)
}

func (ptr *QSqlQuery) __boundValues_setList(key string, i core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__boundValues_setList", key, i})
}

func (ptr *QSqlQuery) __boundValues_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__boundValues_newList"}).(unsafe.Pointer)
}

func (ptr *QSqlQuery) __boundValues_keyList() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__boundValues_keyList"}).([]string)
}

func (ptr *QSqlQuery) ____boundValues_keyList_atList(i int) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____boundValues_keyList_atList", i}).(string)
}

func (ptr *QSqlQuery) ____boundValues_keyList_setList(i string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____boundValues_keyList_setList", i})
}

func (ptr *QSqlQuery) ____boundValues_keyList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____boundValues_keyList_newList"}).(unsafe.Pointer)
}

type QSqlQueryModel struct {
	core.QAbstractTableModel
}

type QSqlQueryModel_ITF interface {
	core.QAbstractTableModel_ITF
	QSqlQueryModel_PTR() *QSqlQueryModel
}

func (ptr *QSqlQueryModel) QSqlQueryModel_PTR() *QSqlQueryModel {
	return ptr
}

func (ptr *QSqlQueryModel) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractTableModel_PTR().Pointer()
	}
	return nil
}

func (ptr *QSqlQueryModel) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAbstractTableModel_PTR().SetPointer(p)
	}
}

func PointerFromQSqlQueryModel(ptr QSqlQueryModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlQueryModel_PTR().Pointer()
	}
	return nil
}

func (n *QSqlQueryModel) InitFromInternal(ptr uintptr, name string) {
	n.QAbstractTableModel_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QSqlQueryModel) ClassNameInternalF() string {
	return n.QAbstractTableModel_PTR().ClassNameInternalF()
}

func NewQSqlQueryModelFromPointer(ptr unsafe.Pointer) (n *QSqlQueryModel) {
	n = new(QSqlQueryModel)
	n.InitFromInternal(uintptr(ptr), "sql.QSqlQueryModel")
	return
}
func NewQSqlQueryModel(parent core.QObject_ITF) *QSqlQueryModel {

	return internal.CallLocalFunction([]interface{}{"", "", "sql.NewQSqlQueryModel", "", parent}).(*QSqlQueryModel)
}

func (ptr *QSqlQueryModel) CanFetchMoreDefault(parent core.QModelIndex_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CanFetchMoreDefault", parent}).(bool)
}

func (ptr *QSqlQueryModel) ConnectClear(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectClear", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlQueryModel) DisconnectClear() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectClear"})
}

func (ptr *QSqlQueryModel) Clear() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Clear"})
}

func (ptr *QSqlQueryModel) ClearDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ClearDefault"})
}

func (ptr *QSqlQueryModel) ConnectColumnCount(f func(index *core.QModelIndex) int) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectColumnCount", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlQueryModel) DisconnectColumnCount() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectColumnCount"})
}

func (ptr *QSqlQueryModel) ColumnCount(index core.QModelIndex_ITF) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ColumnCount", index}).(float64))
}

func (ptr *QSqlQueryModel) ColumnCountDefault(index core.QModelIndex_ITF) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ColumnCountDefault", index}).(float64))
}

func (ptr *QSqlQueryModel) ConnectData(f func(item *core.QModelIndex, role int) *core.QVariant) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectData", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlQueryModel) DisconnectData() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectData"})
}

func (ptr *QSqlQueryModel) Data(item core.QModelIndex_ITF, role int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Data", item, role}).(*core.QVariant)
}

func (ptr *QSqlQueryModel) DataDefault(item core.QModelIndex_ITF, role int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DataDefault", item, role}).(*core.QVariant)
}

func (ptr *QSqlQueryModel) FetchMoreDefault(parent core.QModelIndex_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FetchMoreDefault", parent})
}

func (ptr *QSqlQueryModel) HeaderDataDefault(section int, orientation core.Qt__Orientation, role int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HeaderDataDefault", section, orientation, role}).(*core.QVariant)
}

func (ptr *QSqlQueryModel) ConnectIndexInQuery(f func(item *core.QModelIndex) *core.QModelIndex) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectIndexInQuery", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlQueryModel) DisconnectIndexInQuery() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectIndexInQuery"})
}

func (ptr *QSqlQueryModel) IndexInQuery(item core.QModelIndex_ITF) *core.QModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IndexInQuery", item}).(*core.QModelIndex)
}

func (ptr *QSqlQueryModel) IndexInQueryDefault(item core.QModelIndex_ITF) *core.QModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IndexInQueryDefault", item}).(*core.QModelIndex)
}

func (ptr *QSqlQueryModel) InsertColumnsDefault(column int, count int, parent core.QModelIndex_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InsertColumnsDefault", column, count, parent}).(bool)
}

func (ptr *QSqlQueryModel) LastError() *QSqlError {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LastError"}).(*QSqlError)
}

func (ptr *QSqlQueryModel) Query() *QSqlQuery {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Query"}).(*QSqlQuery)
}

func (ptr *QSqlQueryModel) ConnectQueryChange(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectQueryChange", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlQueryModel) DisconnectQueryChange() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectQueryChange"})
}

func (ptr *QSqlQueryModel) QueryChange() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "QueryChange"})
}

func (ptr *QSqlQueryModel) QueryChangeDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "QueryChangeDefault"})
}

func (ptr *QSqlQueryModel) Record(row int) *QSqlRecord {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Record", row}).(*QSqlRecord)
}

func (ptr *QSqlQueryModel) Record2() *QSqlRecord {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Record2"}).(*QSqlRecord)
}

func (ptr *QSqlQueryModel) RemoveColumnsDefault(column int, count int, parent core.QModelIndex_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RemoveColumnsDefault", column, count, parent}).(bool)
}

func (ptr *QSqlQueryModel) RoleNamesDefault() map[int]*core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RoleNamesDefault"}).(map[int]*core.QByteArray)
}

func (ptr *QSqlQueryModel) ConnectRowCount(f func(parent *core.QModelIndex) int) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRowCount", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlQueryModel) DisconnectRowCount() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRowCount"})
}

func (ptr *QSqlQueryModel) RowCount(parent core.QModelIndex_ITF) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RowCount", parent}).(float64))
}

func (ptr *QSqlQueryModel) RowCountDefault(parent core.QModelIndex_ITF) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RowCountDefault", parent}).(float64))
}

func (ptr *QSqlQueryModel) SetHeaderDataDefault(section int, orientation core.Qt__Orientation, value core.QVariant_ITF, role int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetHeaderDataDefault", section, orientation, value, role}).(bool)
}

func (ptr *QSqlQueryModel) SetLastError(error QSqlError_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLastError", error})
}

func (ptr *QSqlQueryModel) SetQuery(query QSqlQuery_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetQuery", query})
}

func (ptr *QSqlQueryModel) SetQuery2(query string, db QSqlDatabase_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetQuery2", query, db})
}

func (ptr *QSqlQueryModel) ConnectDestroyQSqlQueryModel(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQSqlQueryModel", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlQueryModel) DisconnectDestroyQSqlQueryModel() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQSqlQueryModel"})
}

func (ptr *QSqlQueryModel) DestroyQSqlQueryModel() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSqlQueryModel"})
}

func (ptr *QSqlQueryModel) DestroyQSqlQueryModelDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSqlQueryModelDefault"})
}

func (ptr *QSqlQueryModel) __roleNames_atList(v int, i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__roleNames_atList", v, i}).(*core.QByteArray)
}

func (ptr *QSqlQueryModel) __roleNames_setList(key int, i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__roleNames_setList", key, i})
}

func (ptr *QSqlQueryModel) __roleNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__roleNames_newList"}).(unsafe.Pointer)
}

func (ptr *QSqlQueryModel) __roleNames_keyList() []int {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__roleNames_keyList"}).([]int)
}

func (ptr *QSqlQueryModel) ____roleNames_keyList_atList(i int) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____roleNames_keyList_atList", i}).(float64))
}

func (ptr *QSqlQueryModel) ____roleNames_keyList_setList(i int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____roleNames_keyList_setList", i})
}

func (ptr *QSqlQueryModel) ____roleNames_keyList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____roleNames_keyList_newList"}).(unsafe.Pointer)
}

func (ptr *QSqlQueryModel) ____itemData_keyList_atList(i int) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____itemData_keyList_atList", i}).(float64))
}

func (ptr *QSqlQueryModel) ____itemData_keyList_setList(i int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____itemData_keyList_setList", i})
}

func (ptr *QSqlQueryModel) ____itemData_keyList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____itemData_keyList_newList"}).(unsafe.Pointer)
}

func (ptr *QSqlQueryModel) ____setItemData_roles_keyList_atList(i int) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____setItemData_roles_keyList_atList", i}).(float64))
}

func (ptr *QSqlQueryModel) ____setItemData_roles_keyList_setList(i int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____setItemData_roles_keyList_setList", i})
}

func (ptr *QSqlQueryModel) ____setItemData_roles_keyList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____setItemData_roles_keyList_newList"}).(unsafe.Pointer)
}

func (ptr *QSqlQueryModel) __changePersistentIndexList_from_atList(i int) *core.QModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__changePersistentIndexList_from_atList", i}).(*core.QModelIndex)
}

func (ptr *QSqlQueryModel) __changePersistentIndexList_from_setList(i core.QModelIndex_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__changePersistentIndexList_from_setList", i})
}

func (ptr *QSqlQueryModel) __changePersistentIndexList_from_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__changePersistentIndexList_from_newList"}).(unsafe.Pointer)
}

func (ptr *QSqlQueryModel) __changePersistentIndexList_to_atList(i int) *core.QModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__changePersistentIndexList_to_atList", i}).(*core.QModelIndex)
}

func (ptr *QSqlQueryModel) __changePersistentIndexList_to_setList(i core.QModelIndex_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__changePersistentIndexList_to_setList", i})
}

func (ptr *QSqlQueryModel) __changePersistentIndexList_to_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__changePersistentIndexList_to_newList"}).(unsafe.Pointer)
}

func (ptr *QSqlQueryModel) __dataChanged_roles_atList(i int) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dataChanged_roles_atList", i}).(float64))
}

func (ptr *QSqlQueryModel) __dataChanged_roles_setList(i int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dataChanged_roles_setList", i})
}

func (ptr *QSqlQueryModel) __dataChanged_roles_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dataChanged_roles_newList"}).(unsafe.Pointer)
}

func (ptr *QSqlQueryModel) __itemData_atList(v int, i int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__itemData_atList", v, i}).(*core.QVariant)
}

func (ptr *QSqlQueryModel) __itemData_setList(key int, i core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__itemData_setList", key, i})
}

func (ptr *QSqlQueryModel) __itemData_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__itemData_newList"}).(unsafe.Pointer)
}

func (ptr *QSqlQueryModel) __itemData_keyList() []int {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__itemData_keyList"}).([]int)
}

func (ptr *QSqlQueryModel) __layoutAboutToBeChanged_parents_atList(i int) *core.QPersistentModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__layoutAboutToBeChanged_parents_atList", i}).(*core.QPersistentModelIndex)
}

func (ptr *QSqlQueryModel) __layoutAboutToBeChanged_parents_setList(i core.QPersistentModelIndex_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__layoutAboutToBeChanged_parents_setList", i})
}

func (ptr *QSqlQueryModel) __layoutAboutToBeChanged_parents_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__layoutAboutToBeChanged_parents_newList"}).(unsafe.Pointer)
}

func (ptr *QSqlQueryModel) __layoutChanged_parents_atList(i int) *core.QPersistentModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__layoutChanged_parents_atList", i}).(*core.QPersistentModelIndex)
}

func (ptr *QSqlQueryModel) __layoutChanged_parents_setList(i core.QPersistentModelIndex_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__layoutChanged_parents_setList", i})
}

func (ptr *QSqlQueryModel) __layoutChanged_parents_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__layoutChanged_parents_newList"}).(unsafe.Pointer)
}

func (ptr *QSqlQueryModel) __match_atList(i int) *core.QModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__match_atList", i}).(*core.QModelIndex)
}

func (ptr *QSqlQueryModel) __match_setList(i core.QModelIndex_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__match_setList", i})
}

func (ptr *QSqlQueryModel) __match_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__match_newList"}).(unsafe.Pointer)
}

func (ptr *QSqlQueryModel) __mimeData_indexes_atList(i int) *core.QModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__mimeData_indexes_atList", i}).(*core.QModelIndex)
}

func (ptr *QSqlQueryModel) __mimeData_indexes_setList(i core.QModelIndex_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__mimeData_indexes_setList", i})
}

func (ptr *QSqlQueryModel) __mimeData_indexes_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__mimeData_indexes_newList"}).(unsafe.Pointer)
}

func (ptr *QSqlQueryModel) __persistentIndexList_atList(i int) *core.QModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__persistentIndexList_atList", i}).(*core.QModelIndex)
}

func (ptr *QSqlQueryModel) __persistentIndexList_setList(i core.QModelIndex_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__persistentIndexList_setList", i})
}

func (ptr *QSqlQueryModel) __persistentIndexList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__persistentIndexList_newList"}).(unsafe.Pointer)
}

func (ptr *QSqlQueryModel) __setItemData_roles_atList(v int, i int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setItemData_roles_atList", v, i}).(*core.QVariant)
}

func (ptr *QSqlQueryModel) __setItemData_roles_setList(key int, i core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setItemData_roles_setList", key, i})
}

func (ptr *QSqlQueryModel) __setItemData_roles_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setItemData_roles_newList"}).(unsafe.Pointer)
}

func (ptr *QSqlQueryModel) __setItemData_roles_keyList() []int {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setItemData_roles_keyList"}).([]int)
}

func (ptr *QSqlQueryModel) ____doSetRoleNames_roleNames_keyList_atList(i int) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____doSetRoleNames_roleNames_keyList_atList", i}).(float64))
}

func (ptr *QSqlQueryModel) ____doSetRoleNames_roleNames_keyList_setList(i int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____doSetRoleNames_roleNames_keyList_setList", i})
}

func (ptr *QSqlQueryModel) ____doSetRoleNames_roleNames_keyList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____doSetRoleNames_roleNames_keyList_newList"}).(unsafe.Pointer)
}

func (ptr *QSqlQueryModel) ____setRoleNames_roleNames_keyList_atList(i int) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____setRoleNames_roleNames_keyList_atList", i}).(float64))
}

func (ptr *QSqlQueryModel) ____setRoleNames_roleNames_keyList_setList(i int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____setRoleNames_roleNames_keyList_setList", i})
}

func (ptr *QSqlQueryModel) ____setRoleNames_roleNames_keyList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____setRoleNames_roleNames_keyList_newList"}).(unsafe.Pointer)
}

func (ptr *QSqlQueryModel) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QSqlQueryModel) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QSqlQueryModel) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QSqlQueryModel) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QSqlQueryModel) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QSqlQueryModel) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QSqlQueryModel) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QSqlQueryModel) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QSqlQueryModel) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QSqlQueryModel) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QSqlQueryModel) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QSqlQueryModel) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QSqlQueryModel) DropMimeDataDefault(data core.QMimeData_ITF, action core.Qt__DropAction, row int, column int, parent core.QModelIndex_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DropMimeDataDefault", data, action, row, column, parent}).(bool)
}

func (ptr *QSqlQueryModel) FlagsDefault(index core.QModelIndex_ITF) core.Qt__ItemFlag {

	return core.Qt__ItemFlag(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FlagsDefault", index}).(float64))
}

func (ptr *QSqlQueryModel) IndexDefault(row int, column int, parent core.QModelIndex_ITF) *core.QModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IndexDefault", row, column, parent}).(*core.QModelIndex)
}

func (ptr *QSqlQueryModel) SiblingDefault(row int, column int, idx core.QModelIndex_ITF) *core.QModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SiblingDefault", row, column, idx}).(*core.QModelIndex)
}

func (ptr *QSqlQueryModel) BuddyDefault(index core.QModelIndex_ITF) *core.QModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BuddyDefault", index}).(*core.QModelIndex)
}

func (ptr *QSqlQueryModel) CanDropMimeDataDefault(data core.QMimeData_ITF, action core.Qt__DropAction, row int, column int, parent core.QModelIndex_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CanDropMimeDataDefault", data, action, row, column, parent}).(bool)
}

func (ptr *QSqlQueryModel) HasChildrenDefault(parent core.QModelIndex_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasChildrenDefault", parent}).(bool)
}

func (ptr *QSqlQueryModel) InsertRowsDefault(row int, count int, parent core.QModelIndex_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InsertRowsDefault", row, count, parent}).(bool)
}

func (ptr *QSqlQueryModel) ItemDataDefault(index core.QModelIndex_ITF) map[int]*core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ItemDataDefault", index}).(map[int]*core.QVariant)
}

func (ptr *QSqlQueryModel) MatchDefault(start core.QModelIndex_ITF, role int, value core.QVariant_ITF, hits int, flags core.Qt__MatchFlag) []*core.QModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MatchDefault", start, role, value, hits, flags}).([]*core.QModelIndex)
}

func (ptr *QSqlQueryModel) MimeDataDefault(indexes []*core.QModelIndex) *core.QMimeData {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MimeDataDefault", indexes}).(*core.QMimeData)
}

func (ptr *QSqlQueryModel) MimeTypesDefault() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MimeTypesDefault"}).([]string)
}

func (ptr *QSqlQueryModel) MoveColumnsDefault(sourceParent core.QModelIndex_ITF, sourceColumn int, count int, destinationParent core.QModelIndex_ITF, destinationChild int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MoveColumnsDefault", sourceParent, sourceColumn, count, destinationParent, destinationChild}).(bool)
}

func (ptr *QSqlQueryModel) MoveRowsDefault(sourceParent core.QModelIndex_ITF, sourceRow int, count int, destinationParent core.QModelIndex_ITF, destinationChild int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MoveRowsDefault", sourceParent, sourceRow, count, destinationParent, destinationChild}).(bool)
}

func (ptr *QSqlQueryModel) ParentDefault(index core.QModelIndex_ITF) *core.QModelIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ParentDefault", index}).(*core.QModelIndex)
}

func (ptr *QSqlQueryModel) RemoveRowsDefault(row int, count int, parent core.QModelIndex_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RemoveRowsDefault", row, count, parent}).(bool)
}

func (ptr *QSqlQueryModel) ResetInternalDataDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ResetInternalDataDefault"})
}

func (ptr *QSqlQueryModel) RevertDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RevertDefault"})
}

func (ptr *QSqlQueryModel) SetDataDefault(index core.QModelIndex_ITF, value core.QVariant_ITF, role int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDataDefault", index, value, role}).(bool)
}

func (ptr *QSqlQueryModel) SetItemDataDefault(index core.QModelIndex_ITF, roles map[int]*core.QVariant) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetItemDataDefault", index, roles}).(bool)
}

func (ptr *QSqlQueryModel) SortDefault(column int, order core.Qt__SortOrder) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SortDefault", column, order})
}

func (ptr *QSqlQueryModel) SpanDefault(index core.QModelIndex_ITF) *core.QSize {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SpanDefault", index}).(*core.QSize)
}

func (ptr *QSqlQueryModel) SubmitDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SubmitDefault"}).(bool)
}

func (ptr *QSqlQueryModel) SupportedDragActionsDefault() core.Qt__DropAction {

	return core.Qt__DropAction(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SupportedDragActionsDefault"}).(float64))
}

func (ptr *QSqlQueryModel) SupportedDropActionsDefault() core.Qt__DropAction {

	return core.Qt__DropAction(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SupportedDropActionsDefault"}).(float64))
}

func (ptr *QSqlQueryModel) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QSqlQueryModel) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QSqlQueryModel) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QSqlQueryModel) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QSqlQueryModel) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QSqlQueryModel) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QSqlQueryModel) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QSqlQueryModel) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QSqlQueryModel) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QSqlRecord struct {
	internal.Internal
}

type QSqlRecord_ITF interface {
	QSqlRecord_PTR() *QSqlRecord
}

func (ptr *QSqlRecord) QSqlRecord_PTR() *QSqlRecord {
	return ptr
}

func (ptr *QSqlRecord) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QSqlRecord) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQSqlRecord(ptr QSqlRecord_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlRecord_PTR().Pointer()
	}
	return nil
}

func (n *QSqlRecord) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQSqlRecordFromPointer(ptr unsafe.Pointer) (n *QSqlRecord) {
	n = new(QSqlRecord)
	n.InitFromInternal(uintptr(ptr), "sql.QSqlRecord")
	return
}
func NewQSqlRecord() *QSqlRecord {

	return internal.CallLocalFunction([]interface{}{"", "", "sql.NewQSqlRecord", ""}).(*QSqlRecord)
}

func NewQSqlRecord2(other QSqlRecord_ITF) *QSqlRecord {

	return internal.CallLocalFunction([]interface{}{"", "", "sql.NewQSqlRecord2", "", other}).(*QSqlRecord)
}

func (ptr *QSqlRecord) Append(field QSqlField_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Append", field})
}

func (ptr *QSqlRecord) Clear() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Clear"})
}

func (ptr *QSqlRecord) ClearValues() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ClearValues"})
}

func (ptr *QSqlRecord) Contains(name string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Contains", name}).(bool)
}

func (ptr *QSqlRecord) Count() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Count"}).(float64))
}

func (ptr *QSqlRecord) Field(index int) *QSqlField {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Field", index}).(*QSqlField)
}

func (ptr *QSqlRecord) Field2(name string) *QSqlField {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Field2", name}).(*QSqlField)
}

func (ptr *QSqlRecord) FieldName(index int) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FieldName", index}).(string)
}

func (ptr *QSqlRecord) IndexOf(name string) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IndexOf", name}).(float64))
}

func (ptr *QSqlRecord) Insert(pos int, field QSqlField_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Insert", pos, field})
}

func (ptr *QSqlRecord) IsEmpty() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsEmpty"}).(bool)
}

func (ptr *QSqlRecord) IsGenerated(name string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsGenerated", name}).(bool)
}

func (ptr *QSqlRecord) IsGenerated2(index int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsGenerated2", index}).(bool)
}

func (ptr *QSqlRecord) IsNull(name string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsNull", name}).(bool)
}

func (ptr *QSqlRecord) IsNull2(index int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsNull2", index}).(bool)
}

func (ptr *QSqlRecord) KeyValues(keyFields QSqlRecord_ITF) *QSqlRecord {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "KeyValues", keyFields}).(*QSqlRecord)
}

func (ptr *QSqlRecord) Remove(pos int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Remove", pos})
}

func (ptr *QSqlRecord) Replace(pos int, field QSqlField_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Replace", pos, field})
}

func (ptr *QSqlRecord) SetGenerated(name string, generated bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetGenerated", name, generated})
}

func (ptr *QSqlRecord) SetGenerated2(index int, generated bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetGenerated2", index, generated})
}

func (ptr *QSqlRecord) SetNull(index int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetNull", index})
}

func (ptr *QSqlRecord) SetNull2(name string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetNull2", name})
}

func (ptr *QSqlRecord) SetValue(index int, val core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetValue", index, val})
}

func (ptr *QSqlRecord) SetValue2(name string, val core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetValue2", name, val})
}

func (ptr *QSqlRecord) Value(index int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Value", index}).(*core.QVariant)
}

func (ptr *QSqlRecord) Value2(name string) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Value2", name}).(*core.QVariant)
}

func (ptr *QSqlRecord) DestroyQSqlRecord() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSqlRecord"})
}

type QSqlRelation struct {
	internal.Internal
}

type QSqlRelation_ITF interface {
	QSqlRelation_PTR() *QSqlRelation
}

func (ptr *QSqlRelation) QSqlRelation_PTR() *QSqlRelation {
	return ptr
}

func (ptr *QSqlRelation) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QSqlRelation) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQSqlRelation(ptr QSqlRelation_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlRelation_PTR().Pointer()
	}
	return nil
}

func (n *QSqlRelation) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQSqlRelationFromPointer(ptr unsafe.Pointer) (n *QSqlRelation) {
	n = new(QSqlRelation)
	n.InitFromInternal(uintptr(ptr), "sql.QSqlRelation")
	return
}

func (ptr *QSqlRelation) DestroyQSqlRelation() {
}

func NewQSqlRelation() *QSqlRelation {

	return internal.CallLocalFunction([]interface{}{"", "", "sql.NewQSqlRelation", ""}).(*QSqlRelation)
}

func NewQSqlRelation2(tableName string, indexColumn string, displayColumn string) *QSqlRelation {

	return internal.CallLocalFunction([]interface{}{"", "", "sql.NewQSqlRelation2", "", tableName, indexColumn, displayColumn}).(*QSqlRelation)
}

func (ptr *QSqlRelation) DisplayColumn() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisplayColumn"}).(string)
}

func (ptr *QSqlRelation) IndexColumn() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IndexColumn"}).(string)
}

func (ptr *QSqlRelation) IsValid() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsValid"}).(bool)
}

func (ptr *QSqlRelation) Swap(other QSqlRelation_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Swap", other})
}

func (ptr *QSqlRelation) TableName() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TableName"}).(string)
}

type QSqlRelationalDelegate struct {
	widgets.QItemDelegate
}

type QSqlRelationalDelegate_ITF interface {
	widgets.QItemDelegate_ITF
	QSqlRelationalDelegate_PTR() *QSqlRelationalDelegate
}

func (ptr *QSqlRelationalDelegate) QSqlRelationalDelegate_PTR() *QSqlRelationalDelegate {
	return ptr
}

func (ptr *QSqlRelationalDelegate) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QItemDelegate_PTR().Pointer()
	}
	return nil
}

func (ptr *QSqlRelationalDelegate) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QItemDelegate_PTR().SetPointer(p)
	}
}

func PointerFromQSqlRelationalDelegate(ptr QSqlRelationalDelegate_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlRelationalDelegate_PTR().Pointer()
	}
	return nil
}

func (n *QSqlRelationalDelegate) InitFromInternal(ptr uintptr, name string) {
	n.QItemDelegate_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QSqlRelationalDelegate) ClassNameInternalF() string {
	return n.QItemDelegate_PTR().ClassNameInternalF()
}

func NewQSqlRelationalDelegateFromPointer(ptr unsafe.Pointer) (n *QSqlRelationalDelegate) {
	n = new(QSqlRelationalDelegate)
	n.InitFromInternal(uintptr(ptr), "sql.QSqlRelationalDelegate")
	return
}
func NewQSqlRelationalDelegate(parent core.QObject_ITF) *QSqlRelationalDelegate {

	return internal.CallLocalFunction([]interface{}{"", "", "sql.NewQSqlRelationalDelegate", "", parent}).(*QSqlRelationalDelegate)
}

func (ptr *QSqlRelationalDelegate) CreateEditorDefault(parent widgets.QWidget_ITF, option widgets.QStyleOptionViewItem_ITF, index core.QModelIndex_ITF) *widgets.QWidget {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreateEditorDefault", parent, option, index}).(*widgets.QWidget)
}

func (ptr *QSqlRelationalDelegate) SetModelDataDefault(editor widgets.QWidget_ITF, model core.QAbstractItemModel_ITF, index core.QModelIndex_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetModelDataDefault", editor, model, index})
}

func (ptr *QSqlRelationalDelegate) ConnectDestroyQSqlRelationalDelegate(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQSqlRelationalDelegate", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlRelationalDelegate) DisconnectDestroyQSqlRelationalDelegate() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQSqlRelationalDelegate"})
}

func (ptr *QSqlRelationalDelegate) DestroyQSqlRelationalDelegate() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSqlRelationalDelegate"})
}

func (ptr *QSqlRelationalDelegate) DestroyQSqlRelationalDelegateDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSqlRelationalDelegateDefault"})
}

func (ptr *QSqlRelationalDelegate) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QSqlRelationalDelegate) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QSqlRelationalDelegate) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QSqlRelationalDelegate) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QSqlRelationalDelegate) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QSqlRelationalDelegate) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QSqlRelationalDelegate) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QSqlRelationalDelegate) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QSqlRelationalDelegate) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QSqlRelationalDelegate) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QSqlRelationalDelegate) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QSqlRelationalDelegate) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QSqlRelationalDelegate) DrawCheckDefault(painter gui.QPainter_ITF, option widgets.QStyleOptionViewItem_ITF, rect core.QRect_ITF, state core.Qt__CheckState) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DrawCheckDefault", painter, option, rect, state})
}

func (ptr *QSqlRelationalDelegate) DrawDecorationDefault(painter gui.QPainter_ITF, option widgets.QStyleOptionViewItem_ITF, rect core.QRect_ITF, pixmap gui.QPixmap_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DrawDecorationDefault", painter, option, rect, pixmap})
}

func (ptr *QSqlRelationalDelegate) DrawDisplayDefault(painter gui.QPainter_ITF, option widgets.QStyleOptionViewItem_ITF, rect core.QRect_ITF, text string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DrawDisplayDefault", painter, option, rect, text})
}

func (ptr *QSqlRelationalDelegate) DrawFocusDefault(painter gui.QPainter_ITF, option widgets.QStyleOptionViewItem_ITF, rect core.QRect_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DrawFocusDefault", painter, option, rect})
}

func (ptr *QSqlRelationalDelegate) EditorEventDefault(event core.QEvent_ITF, model core.QAbstractItemModel_ITF, option widgets.QStyleOptionViewItem_ITF, index core.QModelIndex_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EditorEventDefault", event, model, option, index}).(bool)
}

func (ptr *QSqlRelationalDelegate) EventFilterDefault(editor core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", editor, event}).(bool)
}

func (ptr *QSqlRelationalDelegate) PaintDefault(painter gui.QPainter_ITF, option widgets.QStyleOptionViewItem_ITF, index core.QModelIndex_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PaintDefault", painter, option, index})
}

func (ptr *QSqlRelationalDelegate) SetEditorDataDefault(editor widgets.QWidget_ITF, index core.QModelIndex_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetEditorDataDefault", editor, index})
}

func (ptr *QSqlRelationalDelegate) SizeHintDefault(option widgets.QStyleOptionViewItem_ITF, index core.QModelIndex_ITF) *core.QSize {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SizeHintDefault", option, index}).(*core.QSize)
}

func (ptr *QSqlRelationalDelegate) UpdateEditorGeometryDefault(editor widgets.QWidget_ITF, option widgets.QStyleOptionViewItem_ITF, index core.QModelIndex_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpdateEditorGeometryDefault", editor, option, index})
}

func (ptr *QSqlRelationalDelegate) DestroyEditorDefault(editor widgets.QWidget_ITF, index core.QModelIndex_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyEditorDefault", editor, index})
}

func (ptr *QSqlRelationalDelegate) HelpEventDefault(event gui.QHelpEvent_ITF, view widgets.QAbstractItemView_ITF, option widgets.QStyleOptionViewItem_ITF, index core.QModelIndex_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HelpEventDefault", event, view, option, index}).(bool)
}

func (ptr *QSqlRelationalDelegate) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QSqlRelationalDelegate) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QSqlRelationalDelegate) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QSqlRelationalDelegate) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QSqlRelationalDelegate) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QSqlRelationalDelegate) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QSqlRelationalDelegate) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QSqlRelationalDelegate) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QSqlRelationalTableModel struct {
	QSqlTableModel
}

type QSqlRelationalTableModel_ITF interface {
	QSqlTableModel_ITF
	QSqlRelationalTableModel_PTR() *QSqlRelationalTableModel
}

func (ptr *QSqlRelationalTableModel) QSqlRelationalTableModel_PTR() *QSqlRelationalTableModel {
	return ptr
}

func (ptr *QSqlRelationalTableModel) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlTableModel_PTR().Pointer()
	}
	return nil
}

func (ptr *QSqlRelationalTableModel) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSqlTableModel_PTR().SetPointer(p)
	}
}

func PointerFromQSqlRelationalTableModel(ptr QSqlRelationalTableModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlRelationalTableModel_PTR().Pointer()
	}
	return nil
}

func (n *QSqlRelationalTableModel) InitFromInternal(ptr uintptr, name string) {
	n.QSqlTableModel_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QSqlRelationalTableModel) ClassNameInternalF() string {
	return n.QSqlTableModel_PTR().ClassNameInternalF()
}

func NewQSqlRelationalTableModelFromPointer(ptr unsafe.Pointer) (n *QSqlRelationalTableModel) {
	n = new(QSqlRelationalTableModel)
	n.InitFromInternal(uintptr(ptr), "sql.QSqlRelationalTableModel")
	return
}

//go:generate stringer -type=QSqlRelationalTableModel__JoinMode
//QSqlRelationalTableModel::JoinMode
type QSqlRelationalTableModel__JoinMode int64

const (
	QSqlRelationalTableModel__InnerJoin QSqlRelationalTableModel__JoinMode = QSqlRelationalTableModel__JoinMode(0)
	QSqlRelationalTableModel__LeftJoin  QSqlRelationalTableModel__JoinMode = QSqlRelationalTableModel__JoinMode(1)
)

func NewQSqlRelationalTableModel(parent core.QObject_ITF, db QSqlDatabase_ITF) *QSqlRelationalTableModel {

	return internal.CallLocalFunction([]interface{}{"", "", "sql.NewQSqlRelationalTableModel", "", parent, db}).(*QSqlRelationalTableModel)
}

func (ptr *QSqlRelationalTableModel) Relation(column int) *QSqlRelation {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Relation", column}).(*QSqlRelation)
}

func (ptr *QSqlRelationalTableModel) ConnectRelationModel(f func(column int) *QSqlTableModel) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRelationModel", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlRelationalTableModel) DisconnectRelationModel() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRelationModel"})
}

func (ptr *QSqlRelationalTableModel) RelationModel(column int) *QSqlTableModel {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RelationModel", column}).(*QSqlTableModel)
}

func (ptr *QSqlRelationalTableModel) RelationModelDefault(column int) *QSqlTableModel {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RelationModelDefault", column}).(*QSqlTableModel)
}

func (ptr *QSqlRelationalTableModel) ConnectRevertRow(f func(row int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRevertRow", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlRelationalTableModel) DisconnectRevertRow() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRevertRow"})
}

func (ptr *QSqlRelationalTableModel) RevertRow(row int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RevertRow", row})
}

func (ptr *QSqlRelationalTableModel) RevertRowDefault(row int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RevertRowDefault", row})
}

func (ptr *QSqlRelationalTableModel) ConnectSelect(f func() bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSelect", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlRelationalTableModel) DisconnectSelect() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSelect"})
}

func (ptr *QSqlRelationalTableModel) Select() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Select"}).(bool)
}

func (ptr *QSqlRelationalTableModel) SelectDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SelectDefault"}).(bool)
}

func (ptr *QSqlRelationalTableModel) SetJoinMode(joinMode QSqlRelationalTableModel__JoinMode) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetJoinMode", joinMode})
}

func (ptr *QSqlRelationalTableModel) ConnectSetRelation(f func(column int, relation *QSqlRelation)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetRelation", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlRelationalTableModel) DisconnectSetRelation() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetRelation"})
}

func (ptr *QSqlRelationalTableModel) SetRelation(column int, relation QSqlRelation_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRelation", column, relation})
}

func (ptr *QSqlRelationalTableModel) SetRelationDefault(column int, relation QSqlRelation_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRelationDefault", column, relation})
}

func (ptr *QSqlRelationalTableModel) ConnectDestroyQSqlRelationalTableModel(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQSqlRelationalTableModel", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlRelationalTableModel) DisconnectDestroyQSqlRelationalTableModel() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQSqlRelationalTableModel"})
}

func (ptr *QSqlRelationalTableModel) DestroyQSqlRelationalTableModel() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSqlRelationalTableModel"})
}

func (ptr *QSqlRelationalTableModel) DestroyQSqlRelationalTableModelDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSqlRelationalTableModelDefault"})
}

type QSqlResult struct {
	internal.Internal
}

type QSqlResult_ITF interface {
	QSqlResult_PTR() *QSqlResult
}

func (ptr *QSqlResult) QSqlResult_PTR() *QSqlResult {
	return ptr
}

func (ptr *QSqlResult) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QSqlResult) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQSqlResult(ptr QSqlResult_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlResult_PTR().Pointer()
	}
	return nil
}

func (n *QSqlResult) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQSqlResultFromPointer(ptr unsafe.Pointer) (n *QSqlResult) {
	n = new(QSqlResult)
	n.InitFromInternal(uintptr(ptr), "sql.QSqlResult")
	return
}

//go:generate stringer -type=QSqlResult__BindingSyntax
//QSqlResult::BindingSyntax
type QSqlResult__BindingSyntax int64

const (
	QSqlResult__PositionalBinding QSqlResult__BindingSyntax = QSqlResult__BindingSyntax(0)
	QSqlResult__NamedBinding      QSqlResult__BindingSyntax = QSqlResult__BindingSyntax(1)
)

func NewQSqlResult(db QSqlDriver_ITF) *QSqlResult {

	return internal.CallLocalFunction([]interface{}{"", "", "sql.NewQSqlResult", "", db}).(*QSqlResult)
}

func (ptr *QSqlResult) AddBindValue(val core.QVariant_ITF, paramType QSql__ParamTypeFlag) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AddBindValue", val, paramType})
}

func (ptr *QSqlResult) At() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "At"}).(float64))
}

func (ptr *QSqlResult) ConnectBindValue(f func(index int, val *core.QVariant, paramType QSql__ParamTypeFlag)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectBindValue", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlResult) DisconnectBindValue() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectBindValue"})
}

func (ptr *QSqlResult) BindValue(index int, val core.QVariant_ITF, paramType QSql__ParamTypeFlag) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BindValue", index, val, paramType})
}

func (ptr *QSqlResult) BindValueDefault(index int, val core.QVariant_ITF, paramType QSql__ParamTypeFlag) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BindValueDefault", index, val, paramType})
}

func (ptr *QSqlResult) ConnectBindValue2(f func(placeholder string, val *core.QVariant, paramType QSql__ParamTypeFlag)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectBindValue2", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlResult) DisconnectBindValue2() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectBindValue2"})
}

func (ptr *QSqlResult) BindValue2(placeholder string, val core.QVariant_ITF, paramType QSql__ParamTypeFlag) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BindValue2", placeholder, val, paramType})
}

func (ptr *QSqlResult) BindValue2Default(placeholder string, val core.QVariant_ITF, paramType QSql__ParamTypeFlag) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BindValue2Default", placeholder, val, paramType})
}

func (ptr *QSqlResult) BindValueType(index int) QSql__ParamTypeFlag {

	return QSql__ParamTypeFlag(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BindValueType", index}).(float64))
}

func (ptr *QSqlResult) BindValueType2(placeholder string) QSql__ParamTypeFlag {

	return QSql__ParamTypeFlag(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BindValueType2", placeholder}).(float64))
}

func (ptr *QSqlResult) BindingSyntax() QSqlResult__BindingSyntax {

	return QSqlResult__BindingSyntax(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BindingSyntax"}).(float64))
}

func (ptr *QSqlResult) BoundValue(index int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BoundValue", index}).(*core.QVariant)
}

func (ptr *QSqlResult) BoundValue2(placeholder string) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BoundValue2", placeholder}).(*core.QVariant)
}

func (ptr *QSqlResult) BoundValueCount() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BoundValueCount"}).(float64))
}

func (ptr *QSqlResult) BoundValueName(index int) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BoundValueName", index}).(string)
}

func (ptr *QSqlResult) BoundValues() []*core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BoundValues"}).([]*core.QVariant)
}

func (ptr *QSqlResult) Clear() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Clear"})
}

func (ptr *QSqlResult) ConnectData(f func(index int) *core.QVariant) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectData", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlResult) DisconnectData() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectData"})
}

func (ptr *QSqlResult) Data(index int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Data", index}).(*core.QVariant)
}

func (ptr *QSqlResult) Driver() *QSqlDriver {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Driver"}).(*QSqlDriver)
}

func (ptr *QSqlResult) ConnectExec(f func() bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectExec", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlResult) DisconnectExec() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectExec"})
}

func (ptr *QSqlResult) Exec() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Exec"}).(bool)
}

func (ptr *QSqlResult) ExecDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ExecDefault"}).(bool)
}

func (ptr *QSqlResult) ExecutedQuery() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ExecutedQuery"}).(string)
}

func (ptr *QSqlResult) ConnectFetch(f func(index int) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFetch", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlResult) DisconnectFetch() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFetch"})
}

func (ptr *QSqlResult) Fetch(index int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Fetch", index}).(bool)
}

func (ptr *QSqlResult) ConnectFetchFirst(f func() bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFetchFirst", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlResult) DisconnectFetchFirst() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFetchFirst"})
}

func (ptr *QSqlResult) FetchFirst() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FetchFirst"}).(bool)
}

func (ptr *QSqlResult) ConnectFetchLast(f func() bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFetchLast", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlResult) DisconnectFetchLast() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFetchLast"})
}

func (ptr *QSqlResult) FetchLast() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FetchLast"}).(bool)
}

func (ptr *QSqlResult) ConnectFetchNext(f func() bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFetchNext", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlResult) DisconnectFetchNext() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFetchNext"})
}

func (ptr *QSqlResult) FetchNext() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FetchNext"}).(bool)
}

func (ptr *QSqlResult) FetchNextDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FetchNextDefault"}).(bool)
}

func (ptr *QSqlResult) ConnectFetchPrevious(f func() bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFetchPrevious", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlResult) DisconnectFetchPrevious() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFetchPrevious"})
}

func (ptr *QSqlResult) FetchPrevious() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FetchPrevious"}).(bool)
}

func (ptr *QSqlResult) FetchPreviousDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FetchPreviousDefault"}).(bool)
}

func (ptr *QSqlResult) ConnectHandle(f func() *core.QVariant) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectHandle", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlResult) DisconnectHandle() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectHandle"})
}

func (ptr *QSqlResult) Handle() *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Handle"}).(*core.QVariant)
}

func (ptr *QSqlResult) HandleDefault() *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HandleDefault"}).(*core.QVariant)
}

func (ptr *QSqlResult) HasOutValues() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasOutValues"}).(bool)
}

func (ptr *QSqlResult) IsActive() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsActive"}).(bool)
}

func (ptr *QSqlResult) IsForwardOnly() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsForwardOnly"}).(bool)
}

func (ptr *QSqlResult) ConnectIsNull(f func(index int) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectIsNull", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlResult) DisconnectIsNull() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectIsNull"})
}

func (ptr *QSqlResult) IsNull(index int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsNull", index}).(bool)
}

func (ptr *QSqlResult) IsSelect() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsSelect"}).(bool)
}

func (ptr *QSqlResult) IsValid() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsValid"}).(bool)
}

func (ptr *QSqlResult) LastError() *QSqlError {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LastError"}).(*QSqlError)
}

func (ptr *QSqlResult) ConnectLastInsertId(f func() *core.QVariant) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLastInsertId", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlResult) DisconnectLastInsertId() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLastInsertId"})
}

func (ptr *QSqlResult) LastInsertId() *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LastInsertId"}).(*core.QVariant)
}

func (ptr *QSqlResult) LastInsertIdDefault() *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LastInsertIdDefault"}).(*core.QVariant)
}

func (ptr *QSqlResult) LastQuery() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LastQuery"}).(string)
}

func (ptr *QSqlResult) ConnectNumRowsAffected(f func() int) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNumRowsAffected", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlResult) DisconnectNumRowsAffected() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNumRowsAffected"})
}

func (ptr *QSqlResult) NumRowsAffected() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NumRowsAffected"}).(float64))
}

func (ptr *QSqlResult) ConnectPrepare(f func(query string) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPrepare", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlResult) DisconnectPrepare() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPrepare"})
}

func (ptr *QSqlResult) Prepare(query string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Prepare", query}).(bool)
}

func (ptr *QSqlResult) PrepareDefault(query string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PrepareDefault", query}).(bool)
}

func (ptr *QSqlResult) ConnectRecord(f func() *QSqlRecord) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRecord", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlResult) DisconnectRecord() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRecord"})
}

func (ptr *QSqlResult) Record() *QSqlRecord {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Record"}).(*QSqlRecord)
}

func (ptr *QSqlResult) RecordDefault() *QSqlRecord {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RecordDefault"}).(*QSqlRecord)
}

func (ptr *QSqlResult) ConnectReset(f func(query string) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectReset", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlResult) DisconnectReset() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectReset"})
}

func (ptr *QSqlResult) Reset(query string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Reset", query}).(bool)
}

func (ptr *QSqlResult) ResetBindCount() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ResetBindCount"})
}

func (ptr *QSqlResult) ConnectSavePrepare(f func(query string) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSavePrepare", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlResult) DisconnectSavePrepare() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSavePrepare"})
}

func (ptr *QSqlResult) SavePrepare(query string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SavePrepare", query}).(bool)
}

func (ptr *QSqlResult) SavePrepareDefault(query string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SavePrepareDefault", query}).(bool)
}

func (ptr *QSqlResult) ConnectSetActive(f func(active bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetActive", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlResult) DisconnectSetActive() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetActive"})
}

func (ptr *QSqlResult) SetActive(active bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetActive", active})
}

func (ptr *QSqlResult) SetActiveDefault(active bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetActiveDefault", active})
}

func (ptr *QSqlResult) ConnectSetAt(f func(index int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetAt", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlResult) DisconnectSetAt() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetAt"})
}

func (ptr *QSqlResult) SetAt(index int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAt", index})
}

func (ptr *QSqlResult) SetAtDefault(index int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAtDefault", index})
}

func (ptr *QSqlResult) ConnectSetForwardOnly(f func(forward bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetForwardOnly", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlResult) DisconnectSetForwardOnly() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetForwardOnly"})
}

func (ptr *QSqlResult) SetForwardOnly(forward bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetForwardOnly", forward})
}

func (ptr *QSqlResult) SetForwardOnlyDefault(forward bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetForwardOnlyDefault", forward})
}

func (ptr *QSqlResult) ConnectSetLastError(f func(error *QSqlError)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetLastError", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlResult) DisconnectSetLastError() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetLastError"})
}

func (ptr *QSqlResult) SetLastError(error QSqlError_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLastError", error})
}

func (ptr *QSqlResult) SetLastErrorDefault(error QSqlError_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLastErrorDefault", error})
}

func (ptr *QSqlResult) ConnectSetQuery(f func(query string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetQuery", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlResult) DisconnectSetQuery() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetQuery"})
}

func (ptr *QSqlResult) SetQuery(query string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetQuery", query})
}

func (ptr *QSqlResult) SetQueryDefault(query string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetQueryDefault", query})
}

func (ptr *QSqlResult) ConnectSetSelect(f func(sele bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetSelect", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlResult) DisconnectSetSelect() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetSelect"})
}

func (ptr *QSqlResult) SetSelect(sele bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSelect", sele})
}

func (ptr *QSqlResult) SetSelectDefault(sele bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSelectDefault", sele})
}

func (ptr *QSqlResult) ConnectSize(f func() int) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSize", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlResult) DisconnectSize() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSize"})
}

func (ptr *QSqlResult) Size() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Size"}).(float64))
}

func (ptr *QSqlResult) ConnectDestroyQSqlResult(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQSqlResult", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlResult) DisconnectDestroyQSqlResult() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQSqlResult"})
}

func (ptr *QSqlResult) DestroyQSqlResult() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSqlResult"})
}

func (ptr *QSqlResult) DestroyQSqlResultDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSqlResultDefault"})
}

func (ptr *QSqlResult) __boundValues_atList(i int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__boundValues_atList", i}).(*core.QVariant)
}

func (ptr *QSqlResult) __boundValues_setList(i core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__boundValues_setList", i})
}

func (ptr *QSqlResult) __boundValues_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__boundValues_newList"}).(unsafe.Pointer)
}

type QSqlTableModel struct {
	QSqlQueryModel
}

type QSqlTableModel_ITF interface {
	QSqlQueryModel_ITF
	QSqlTableModel_PTR() *QSqlTableModel
}

func (ptr *QSqlTableModel) QSqlTableModel_PTR() *QSqlTableModel {
	return ptr
}

func (ptr *QSqlTableModel) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlQueryModel_PTR().Pointer()
	}
	return nil
}

func (ptr *QSqlTableModel) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSqlQueryModel_PTR().SetPointer(p)
	}
}

func PointerFromQSqlTableModel(ptr QSqlTableModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlTableModel_PTR().Pointer()
	}
	return nil
}

func (n *QSqlTableModel) InitFromInternal(ptr uintptr, name string) {
	n.QSqlQueryModel_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QSqlTableModel) ClassNameInternalF() string {
	return n.QSqlQueryModel_PTR().ClassNameInternalF()
}

func NewQSqlTableModelFromPointer(ptr unsafe.Pointer) (n *QSqlTableModel) {
	n = new(QSqlTableModel)
	n.InitFromInternal(uintptr(ptr), "sql.QSqlTableModel")
	return
}

//go:generate stringer -type=QSqlTableModel__EditStrategy
//QSqlTableModel::EditStrategy
type QSqlTableModel__EditStrategy int64

const (
	QSqlTableModel__OnFieldChange  QSqlTableModel__EditStrategy = QSqlTableModel__EditStrategy(0)
	QSqlTableModel__OnRowChange    QSqlTableModel__EditStrategy = QSqlTableModel__EditStrategy(1)
	QSqlTableModel__OnManualSubmit QSqlTableModel__EditStrategy = QSqlTableModel__EditStrategy(2)
)

func NewQSqlTableModel(parent core.QObject_ITF, db QSqlDatabase_ITF) *QSqlTableModel {

	return internal.CallLocalFunction([]interface{}{"", "", "sql.NewQSqlTableModel", "", parent, db}).(*QSqlTableModel)
}

func (ptr *QSqlTableModel) ConnectBeforeDelete(f func(row int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectBeforeDelete", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlTableModel) DisconnectBeforeDelete() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectBeforeDelete"})
}

func (ptr *QSqlTableModel) BeforeDelete(row int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BeforeDelete", row})
}

func (ptr *QSqlTableModel) ConnectBeforeInsert(f func(record *QSqlRecord)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectBeforeInsert", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlTableModel) DisconnectBeforeInsert() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectBeforeInsert"})
}

func (ptr *QSqlTableModel) BeforeInsert(record QSqlRecord_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BeforeInsert", record})
}

func (ptr *QSqlTableModel) ConnectBeforeUpdate(f func(row int, record *QSqlRecord)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectBeforeUpdate", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlTableModel) DisconnectBeforeUpdate() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectBeforeUpdate"})
}

func (ptr *QSqlTableModel) BeforeUpdate(row int, record QSqlRecord_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BeforeUpdate", row, record})
}

func (ptr *QSqlTableModel) Database() *QSqlDatabase {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Database"}).(*QSqlDatabase)
}

func (ptr *QSqlTableModel) ConnectDeleteRowFromTable(f func(row int) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDeleteRowFromTable", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlTableModel) DisconnectDeleteRowFromTable() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDeleteRowFromTable"})
}

func (ptr *QSqlTableModel) DeleteRowFromTable(row int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteRowFromTable", row}).(bool)
}

func (ptr *QSqlTableModel) DeleteRowFromTableDefault(row int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteRowFromTableDefault", row}).(bool)
}

func (ptr *QSqlTableModel) EditStrategy() QSqlTableModel__EditStrategy {

	return QSqlTableModel__EditStrategy(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EditStrategy"}).(float64))
}

func (ptr *QSqlTableModel) FieldIndex(fieldName string) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FieldIndex", fieldName}).(float64))
}

func (ptr *QSqlTableModel) Filter() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Filter"}).(string)
}

func (ptr *QSqlTableModel) InsertRecord(row int, record QSqlRecord_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InsertRecord", row, record}).(bool)
}

func (ptr *QSqlTableModel) ConnectInsertRowIntoTable(f func(values *QSqlRecord) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectInsertRowIntoTable", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlTableModel) DisconnectInsertRowIntoTable() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectInsertRowIntoTable"})
}

func (ptr *QSqlTableModel) InsertRowIntoTable(values QSqlRecord_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InsertRowIntoTable", values}).(bool)
}

func (ptr *QSqlTableModel) InsertRowIntoTableDefault(values QSqlRecord_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InsertRowIntoTableDefault", values}).(bool)
}

func (ptr *QSqlTableModel) IsDirty(index core.QModelIndex_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsDirty", index}).(bool)
}

func (ptr *QSqlTableModel) IsDirty2() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsDirty2"}).(bool)
}

func (ptr *QSqlTableModel) ConnectOrderByClause(f func() string) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectOrderByClause", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlTableModel) DisconnectOrderByClause() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectOrderByClause"})
}

func (ptr *QSqlTableModel) OrderByClause() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "OrderByClause"}).(string)
}

func (ptr *QSqlTableModel) OrderByClauseDefault() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "OrderByClauseDefault"}).(string)
}

func (ptr *QSqlTableModel) PrimaryKey() *QSqlIndex {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PrimaryKey"}).(*QSqlIndex)
}

func (ptr *QSqlTableModel) PrimaryValues(row int) *QSqlRecord {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PrimaryValues", row}).(*QSqlRecord)
}

func (ptr *QSqlTableModel) ConnectPrimeInsert(f func(row int, record *QSqlRecord)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPrimeInsert", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlTableModel) DisconnectPrimeInsert() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPrimeInsert"})
}

func (ptr *QSqlTableModel) PrimeInsert(row int, record QSqlRecord_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PrimeInsert", row, record})
}

func (ptr *QSqlTableModel) ConnectRevertAll(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRevertAll", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlTableModel) DisconnectRevertAll() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRevertAll"})
}

func (ptr *QSqlTableModel) RevertAll() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RevertAll"})
}

func (ptr *QSqlTableModel) RevertAllDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RevertAllDefault"})
}

func (ptr *QSqlTableModel) ConnectRevertRow(f func(row int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRevertRow", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlTableModel) DisconnectRevertRow() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRevertRow"})
}

func (ptr *QSqlTableModel) RevertRow(row int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RevertRow", row})
}

func (ptr *QSqlTableModel) RevertRowDefault(row int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RevertRowDefault", row})
}

func (ptr *QSqlTableModel) ConnectSelect(f func() bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSelect", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlTableModel) DisconnectSelect() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSelect"})
}

func (ptr *QSqlTableModel) Select() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Select"}).(bool)
}

func (ptr *QSqlTableModel) SelectDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SelectDefault"}).(bool)
}

func (ptr *QSqlTableModel) ConnectSelectRow(f func(row int) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSelectRow", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlTableModel) DisconnectSelectRow() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSelectRow"})
}

func (ptr *QSqlTableModel) SelectRow(row int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SelectRow", row}).(bool)
}

func (ptr *QSqlTableModel) SelectRowDefault(row int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SelectRowDefault", row}).(bool)
}

func (ptr *QSqlTableModel) ConnectSelectStatement(f func() string) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSelectStatement", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlTableModel) DisconnectSelectStatement() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSelectStatement"})
}

func (ptr *QSqlTableModel) SelectStatement() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SelectStatement"}).(string)
}

func (ptr *QSqlTableModel) SelectStatementDefault() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SelectStatementDefault"}).(string)
}

func (ptr *QSqlTableModel) ConnectSetEditStrategy(f func(strategy QSqlTableModel__EditStrategy)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetEditStrategy", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlTableModel) DisconnectSetEditStrategy() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetEditStrategy"})
}

func (ptr *QSqlTableModel) SetEditStrategy(strategy QSqlTableModel__EditStrategy) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetEditStrategy", strategy})
}

func (ptr *QSqlTableModel) SetEditStrategyDefault(strategy QSqlTableModel__EditStrategy) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetEditStrategyDefault", strategy})
}

func (ptr *QSqlTableModel) ConnectSetFilter(f func(filter string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetFilter", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlTableModel) DisconnectSetFilter() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetFilter"})
}

func (ptr *QSqlTableModel) SetFilter(filter string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFilter", filter})
}

func (ptr *QSqlTableModel) SetFilterDefault(filter string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFilterDefault", filter})
}

func (ptr *QSqlTableModel) SetPrimaryKey(key QSqlIndex_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPrimaryKey", key})
}

func (ptr *QSqlTableModel) SetRecord(row int, values QSqlRecord_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRecord", row, values}).(bool)
}

func (ptr *QSqlTableModel) ConnectSetSort(f func(column int, order core.Qt__SortOrder)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetSort", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlTableModel) DisconnectSetSort() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetSort"})
}

func (ptr *QSqlTableModel) SetSort(column int, order core.Qt__SortOrder) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSort", column, order})
}

func (ptr *QSqlTableModel) SetSortDefault(column int, order core.Qt__SortOrder) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSortDefault", column, order})
}

func (ptr *QSqlTableModel) ConnectSetTable(f func(tableName string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetTable", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlTableModel) DisconnectSetTable() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetTable"})
}

func (ptr *QSqlTableModel) SetTable(tableName string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetTable", tableName})
}

func (ptr *QSqlTableModel) SetTableDefault(tableName string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetTableDefault", tableName})
}

func (ptr *QSqlTableModel) ConnectSubmitAll(f func() bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSubmitAll", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlTableModel) DisconnectSubmitAll() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSubmitAll"})
}

func (ptr *QSqlTableModel) SubmitAll() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SubmitAll"}).(bool)
}

func (ptr *QSqlTableModel) SubmitAllDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SubmitAllDefault"}).(bool)
}

func (ptr *QSqlTableModel) TableName() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TableName"}).(string)
}

func (ptr *QSqlTableModel) ConnectUpdateRowInTable(f func(row int, values *QSqlRecord) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectUpdateRowInTable", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlTableModel) DisconnectUpdateRowInTable() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectUpdateRowInTable"})
}

func (ptr *QSqlTableModel) UpdateRowInTable(row int, values QSqlRecord_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpdateRowInTable", row, values}).(bool)
}

func (ptr *QSqlTableModel) UpdateRowInTableDefault(row int, values QSqlRecord_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpdateRowInTableDefault", row, values}).(bool)
}

func (ptr *QSqlTableModel) ConnectDestroyQSqlTableModel(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQSqlTableModel", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSqlTableModel) DisconnectDestroyQSqlTableModel() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQSqlTableModel"})
}

func (ptr *QSqlTableModel) DestroyQSqlTableModel() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSqlTableModel"})
}

func (ptr *QSqlTableModel) DestroyQSqlTableModelDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSqlTableModelDefault"})
}

func init() {
	internal.ConstructorTable["sql.QSql"] = NewQSqlFromPointer
	internal.ConstructorTable["sql.QSqlDatabase"] = NewQSqlDatabaseFromPointer
	internal.ConstructorTable["sql.QSqlDriver"] = NewQSqlDriverFromPointer
	internal.ConstructorTable["sql.QSqlDriverCreatorBase"] = NewQSqlDriverCreatorBaseFromPointer
	internal.ConstructorTable["sql.QSqlDriverPlugin"] = NewQSqlDriverPluginFromPointer
	internal.ConstructorTable["sql.QSqlError"] = NewQSqlErrorFromPointer
	internal.ConstructorTable["sql.QSqlField"] = NewQSqlFieldFromPointer
	internal.ConstructorTable["sql.QSqlIndex"] = NewQSqlIndexFromPointer
	internal.ConstructorTable["sql.QSqlQuery"] = NewQSqlQueryFromPointer
	internal.ConstructorTable["sql.QSqlQueryModel"] = NewQSqlQueryModelFromPointer
	internal.ConstructorTable["sql.QSqlRecord"] = NewQSqlRecordFromPointer
	internal.ConstructorTable["sql.QSqlRelation"] = NewQSqlRelationFromPointer
	internal.ConstructorTable["sql.QSqlRelationalDelegate"] = NewQSqlRelationalDelegateFromPointer
	internal.ConstructorTable["sql.QSqlRelationalTableModel"] = NewQSqlRelationalTableModelFromPointer
	internal.ConstructorTable["sql.QSqlResult"] = NewQSqlResultFromPointer
	internal.ConstructorTable["sql.QSqlTableModel"] = NewQSqlTableModelFromPointer
}
