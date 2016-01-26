package sql

//#include "sql.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

type QSqlDatabase struct {
	ptr unsafe.Pointer
}

type QSqlDatabase_ITF interface {
	QSqlDatabase_PTR() *QSqlDatabase
}

func (p *QSqlDatabase) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSqlDatabase) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSqlDatabase(ptr QSqlDatabase_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlDatabase_PTR().Pointer()
	}
	return nil
}

func NewQSqlDatabaseFromPointer(ptr unsafe.Pointer) *QSqlDatabase {
	var n = new(QSqlDatabase)
	n.SetPointer(ptr)
	return n
}

func newQSqlDatabaseFromPointer(ptr unsafe.Pointer) *QSqlDatabase {
	var n = NewQSqlDatabaseFromPointer(ptr)
	return n
}

func (ptr *QSqlDatabase) QSqlDatabase_PTR() *QSqlDatabase {
	return ptr
}

func NewQSqlDatabase() *QSqlDatabase {
	defer qt.Recovering("QSqlDatabase::QSqlDatabase")

	return newQSqlDatabaseFromPointer(C.QSqlDatabase_NewQSqlDatabase())
}

func NewQSqlDatabase2(other QSqlDatabase_ITF) *QSqlDatabase {
	defer qt.Recovering("QSqlDatabase::QSqlDatabase")

	return newQSqlDatabaseFromPointer(C.QSqlDatabase_NewQSqlDatabase2(PointerFromQSqlDatabase(other)))
}

func (ptr *QSqlDatabase) Close() {
	defer qt.Recovering("QSqlDatabase::close")

	if ptr.Pointer() != nil {
		C.QSqlDatabase_Close(ptr.Pointer())
	}
}

func (ptr *QSqlDatabase) Commit() bool {
	defer qt.Recovering("QSqlDatabase::commit")

	if ptr.Pointer() != nil {
		return C.QSqlDatabase_Commit(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlDatabase) ConnectOptions() string {
	defer qt.Recovering("QSqlDatabase::connectOptions")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlDatabase_ConnectOptions(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlDatabase) ConnectionName() string {
	defer qt.Recovering("QSqlDatabase::connectionName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlDatabase_ConnectionName(ptr.Pointer()))
	}
	return ""
}

func QSqlDatabase_ConnectionNames() []string {
	defer qt.Recovering("QSqlDatabase::connectionNames")

	return strings.Split(C.GoString(C.QSqlDatabase_QSqlDatabase_ConnectionNames()), "|")
}

func QSqlDatabase_Contains(connectionName string) bool {
	defer qt.Recovering("QSqlDatabase::contains")

	return C.QSqlDatabase_QSqlDatabase_Contains(C.CString(connectionName)) != 0
}

func (ptr *QSqlDatabase) DatabaseName() string {
	defer qt.Recovering("QSqlDatabase::databaseName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlDatabase_DatabaseName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlDatabase) Driver() *QSqlDriver {
	defer qt.Recovering("QSqlDatabase::driver")

	if ptr.Pointer() != nil {
		return NewQSqlDriverFromPointer(C.QSqlDatabase_Driver(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSqlDatabase) DriverName() string {
	defer qt.Recovering("QSqlDatabase::driverName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlDatabase_DriverName(ptr.Pointer()))
	}
	return ""
}

func QSqlDatabase_Drivers() []string {
	defer qt.Recovering("QSqlDatabase::drivers")

	return strings.Split(C.GoString(C.QSqlDatabase_QSqlDatabase_Drivers()), "|")
}

func (ptr *QSqlDatabase) HostName() string {
	defer qt.Recovering("QSqlDatabase::hostName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlDatabase_HostName(ptr.Pointer()))
	}
	return ""
}

func QSqlDatabase_IsDriverAvailable(name string) bool {
	defer qt.Recovering("QSqlDatabase::isDriverAvailable")

	return C.QSqlDatabase_QSqlDatabase_IsDriverAvailable(C.CString(name)) != 0
}

func (ptr *QSqlDatabase) IsOpen() bool {
	defer qt.Recovering("QSqlDatabase::isOpen")

	if ptr.Pointer() != nil {
		return C.QSqlDatabase_IsOpen(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlDatabase) IsOpenError() bool {
	defer qt.Recovering("QSqlDatabase::isOpenError")

	if ptr.Pointer() != nil {
		return C.QSqlDatabase_IsOpenError(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlDatabase) IsValid() bool {
	defer qt.Recovering("QSqlDatabase::isValid")

	if ptr.Pointer() != nil {
		return C.QSqlDatabase_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlDatabase) Open() bool {
	defer qt.Recovering("QSqlDatabase::open")

	if ptr.Pointer() != nil {
		return C.QSqlDatabase_Open(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlDatabase) Open2(user string, password string) bool {
	defer qt.Recovering("QSqlDatabase::open")

	if ptr.Pointer() != nil {
		return C.QSqlDatabase_Open2(ptr.Pointer(), C.CString(user), C.CString(password)) != 0
	}
	return false
}

func (ptr *QSqlDatabase) Password() string {
	defer qt.Recovering("QSqlDatabase::password")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlDatabase_Password(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlDatabase) Port() int {
	defer qt.Recovering("QSqlDatabase::port")

	if ptr.Pointer() != nil {
		return int(C.QSqlDatabase_Port(ptr.Pointer()))
	}
	return 0
}

func QSqlDatabase_RegisterSqlDriver(name string, creator QSqlDriverCreatorBase_ITF) {
	defer qt.Recovering("QSqlDatabase::registerSqlDriver")

	C.QSqlDatabase_QSqlDatabase_RegisterSqlDriver(C.CString(name), PointerFromQSqlDriverCreatorBase(creator))
}

func QSqlDatabase_RemoveDatabase(connectionName string) {
	defer qt.Recovering("QSqlDatabase::removeDatabase")

	C.QSqlDatabase_QSqlDatabase_RemoveDatabase(C.CString(connectionName))
}

func (ptr *QSqlDatabase) Rollback() bool {
	defer qt.Recovering("QSqlDatabase::rollback")

	if ptr.Pointer() != nil {
		return C.QSqlDatabase_Rollback(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlDatabase) SetConnectOptions(options string) {
	defer qt.Recovering("QSqlDatabase::setConnectOptions")

	if ptr.Pointer() != nil {
		C.QSqlDatabase_SetConnectOptions(ptr.Pointer(), C.CString(options))
	}
}

func (ptr *QSqlDatabase) SetDatabaseName(name string) {
	defer qt.Recovering("QSqlDatabase::setDatabaseName")

	if ptr.Pointer() != nil {
		C.QSqlDatabase_SetDatabaseName(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QSqlDatabase) SetHostName(host string) {
	defer qt.Recovering("QSqlDatabase::setHostName")

	if ptr.Pointer() != nil {
		C.QSqlDatabase_SetHostName(ptr.Pointer(), C.CString(host))
	}
}

func (ptr *QSqlDatabase) SetPassword(password string) {
	defer qt.Recovering("QSqlDatabase::setPassword")

	if ptr.Pointer() != nil {
		C.QSqlDatabase_SetPassword(ptr.Pointer(), C.CString(password))
	}
}

func (ptr *QSqlDatabase) SetPort(port int) {
	defer qt.Recovering("QSqlDatabase::setPort")

	if ptr.Pointer() != nil {
		C.QSqlDatabase_SetPort(ptr.Pointer(), C.int(port))
	}
}

func (ptr *QSqlDatabase) SetUserName(name string) {
	defer qt.Recovering("QSqlDatabase::setUserName")

	if ptr.Pointer() != nil {
		C.QSqlDatabase_SetUserName(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QSqlDatabase) Transaction() bool {
	defer qt.Recovering("QSqlDatabase::transaction")

	if ptr.Pointer() != nil {
		return C.QSqlDatabase_Transaction(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlDatabase) UserName() string {
	defer qt.Recovering("QSqlDatabase::userName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlDatabase_UserName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlDatabase) DestroyQSqlDatabase() {
	defer qt.Recovering("QSqlDatabase::~QSqlDatabase")

	if ptr.Pointer() != nil {
		C.QSqlDatabase_DestroyQSqlDatabase(ptr.Pointer())
	}
}

type QSqlDriver struct {
	core.QObject
}

type QSqlDriver_ITF interface {
	core.QObject_ITF
	QSqlDriver_PTR() *QSqlDriver
}

func PointerFromQSqlDriver(ptr QSqlDriver_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlDriver_PTR().Pointer()
	}
	return nil
}

func NewQSqlDriverFromPointer(ptr unsafe.Pointer) *QSqlDriver {
	var n = new(QSqlDriver)
	n.SetPointer(ptr)
	return n
}

func newQSqlDriverFromPointer(ptr unsafe.Pointer) *QSqlDriver {
	var n = NewQSqlDriverFromPointer(ptr)
	for len(n.ObjectName()) < len("QSqlDriver_") {
		n.SetObjectName("QSqlDriver_" + qt.Identifier())
	}
	return n
}

func (ptr *QSqlDriver) QSqlDriver_PTR() *QSqlDriver {
	return ptr
}

//QSqlDriver::DbmsType
type QSqlDriver__DbmsType int64

const (
	QSqlDriver__UnknownDbms = QSqlDriver__DbmsType(0)
	QSqlDriver__MSSqlServer = QSqlDriver__DbmsType(1)
	QSqlDriver__MySqlServer = QSqlDriver__DbmsType(2)
	QSqlDriver__PostgreSQL  = QSqlDriver__DbmsType(3)
	QSqlDriver__Oracle      = QSqlDriver__DbmsType(4)
	QSqlDriver__Sybase      = QSqlDriver__DbmsType(5)
	QSqlDriver__SQLite      = QSqlDriver__DbmsType(6)
	QSqlDriver__Interbase   = QSqlDriver__DbmsType(7)
	QSqlDriver__DB2         = QSqlDriver__DbmsType(8)
)

//QSqlDriver::DriverFeature
type QSqlDriver__DriverFeature int64

const (
	QSqlDriver__Transactions           = QSqlDriver__DriverFeature(0)
	QSqlDriver__QuerySize              = QSqlDriver__DriverFeature(1)
	QSqlDriver__BLOB                   = QSqlDriver__DriverFeature(2)
	QSqlDriver__Unicode                = QSqlDriver__DriverFeature(3)
	QSqlDriver__PreparedQueries        = QSqlDriver__DriverFeature(4)
	QSqlDriver__NamedPlaceholders      = QSqlDriver__DriverFeature(5)
	QSqlDriver__PositionalPlaceholders = QSqlDriver__DriverFeature(6)
	QSqlDriver__LastInsertId           = QSqlDriver__DriverFeature(7)
	QSqlDriver__BatchOperations        = QSqlDriver__DriverFeature(8)
	QSqlDriver__SimpleLocking          = QSqlDriver__DriverFeature(9)
	QSqlDriver__LowPrecisionNumbers    = QSqlDriver__DriverFeature(10)
	QSqlDriver__EventNotifications     = QSqlDriver__DriverFeature(11)
	QSqlDriver__FinishQuery            = QSqlDriver__DriverFeature(12)
	QSqlDriver__MultipleResultSets     = QSqlDriver__DriverFeature(13)
	QSqlDriver__CancelQuery            = QSqlDriver__DriverFeature(14)
)

//QSqlDriver::IdentifierType
type QSqlDriver__IdentifierType int64

const (
	QSqlDriver__FieldName = QSqlDriver__IdentifierType(0)
	QSqlDriver__TableName = QSqlDriver__IdentifierType(1)
)

//QSqlDriver::NotificationSource
type QSqlDriver__NotificationSource int64

const (
	QSqlDriver__UnknownSource = QSqlDriver__NotificationSource(0)
	QSqlDriver__SelfSource    = QSqlDriver__NotificationSource(1)
	QSqlDriver__OtherSource   = QSqlDriver__NotificationSource(2)
)

//QSqlDriver::StatementType
type QSqlDriver__StatementType int64

const (
	QSqlDriver__WhereStatement  = QSqlDriver__StatementType(0)
	QSqlDriver__SelectStatement = QSqlDriver__StatementType(1)
	QSqlDriver__UpdateStatement = QSqlDriver__StatementType(2)
	QSqlDriver__InsertStatement = QSqlDriver__StatementType(3)
	QSqlDriver__DeleteStatement = QSqlDriver__StatementType(4)
)

func (ptr *QSqlDriver) BeginTransaction() bool {
	defer qt.Recovering("QSqlDriver::beginTransaction")

	if ptr.Pointer() != nil {
		return C.QSqlDriver_BeginTransaction(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlDriver) Close() {
	defer qt.Recovering("QSqlDriver::close")

	if ptr.Pointer() != nil {
		C.QSqlDriver_Close(ptr.Pointer())
	}
}

func (ptr *QSqlDriver) CommitTransaction() bool {
	defer qt.Recovering("QSqlDriver::commitTransaction")

	if ptr.Pointer() != nil {
		return C.QSqlDriver_CommitTransaction(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlDriver) CreateResult() *QSqlResult {
	defer qt.Recovering("QSqlDriver::createResult")

	if ptr.Pointer() != nil {
		return NewQSqlResultFromPointer(C.QSqlDriver_CreateResult(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSqlDriver) DbmsType() QSqlDriver__DbmsType {
	defer qt.Recovering("QSqlDriver::dbmsType")

	if ptr.Pointer() != nil {
		return QSqlDriver__DbmsType(C.QSqlDriver_DbmsType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSqlDriver) EscapeIdentifier(identifier string, ty QSqlDriver__IdentifierType) string {
	defer qt.Recovering("QSqlDriver::escapeIdentifier")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlDriver_EscapeIdentifier(ptr.Pointer(), C.CString(identifier), C.int(ty)))
	}
	return ""
}

func (ptr *QSqlDriver) FormatValue(field QSqlField_ITF, trimStrings bool) string {
	defer qt.Recovering("QSqlDriver::formatValue")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlDriver_FormatValue(ptr.Pointer(), PointerFromQSqlField(field), C.int(qt.GoBoolToInt(trimStrings))))
	}
	return ""
}

func (ptr *QSqlDriver) Handle() *core.QVariant {
	defer qt.Recovering("QSqlDriver::handle")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QSqlDriver_Handle(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSqlDriver) HasFeature(feature QSqlDriver__DriverFeature) bool {
	defer qt.Recovering("QSqlDriver::hasFeature")

	if ptr.Pointer() != nil {
		return C.QSqlDriver_HasFeature(ptr.Pointer(), C.int(feature)) != 0
	}
	return false
}

func (ptr *QSqlDriver) IsIdentifierEscaped(identifier string, ty QSqlDriver__IdentifierType) bool {
	defer qt.Recovering("QSqlDriver::isIdentifierEscaped")

	if ptr.Pointer() != nil {
		return C.QSqlDriver_IsIdentifierEscaped(ptr.Pointer(), C.CString(identifier), C.int(ty)) != 0
	}
	return false
}

func (ptr *QSqlDriver) IsOpen() bool {
	defer qt.Recovering("QSqlDriver::isOpen")

	if ptr.Pointer() != nil {
		return C.QSqlDriver_IsOpen(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlDriver) IsOpenError() bool {
	defer qt.Recovering("QSqlDriver::isOpenError")

	if ptr.Pointer() != nil {
		return C.QSqlDriver_IsOpenError(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlDriver) ConnectNotification(f func(name string)) {
	defer qt.Recovering("connect QSqlDriver::notification")

	if ptr.Pointer() != nil {
		C.QSqlDriver_ConnectNotification(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "notification", f)
	}
}

func (ptr *QSqlDriver) DisconnectNotification() {
	defer qt.Recovering("disconnect QSqlDriver::notification")

	if ptr.Pointer() != nil {
		C.QSqlDriver_DisconnectNotification(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "notification")
	}
}

//export callbackQSqlDriverNotification
func callbackQSqlDriverNotification(ptr unsafe.Pointer, ptrName *C.char, name *C.char) {
	defer qt.Recovering("callback QSqlDriver::notification")

	if signal := qt.GetSignal(C.GoString(ptrName), "notification"); signal != nil {
		signal.(func(string))(C.GoString(name))
	}

}

func (ptr *QSqlDriver) Notification(name string) {
	defer qt.Recovering("QSqlDriver::notification")

	if ptr.Pointer() != nil {
		C.QSqlDriver_Notification(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QSqlDriver) ConnectNotification2(f func(name string, source QSqlDriver__NotificationSource, payload *core.QVariant)) {
	defer qt.Recovering("connect QSqlDriver::notification")

	if ptr.Pointer() != nil {
		C.QSqlDriver_ConnectNotification2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "notification2", f)
	}
}

func (ptr *QSqlDriver) DisconnectNotification2() {
	defer qt.Recovering("disconnect QSqlDriver::notification")

	if ptr.Pointer() != nil {
		C.QSqlDriver_DisconnectNotification2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "notification2")
	}
}

//export callbackQSqlDriverNotification2
func callbackQSqlDriverNotification2(ptr unsafe.Pointer, ptrName *C.char, name *C.char, source C.int, payload unsafe.Pointer) {
	defer qt.Recovering("callback QSqlDriver::notification")

	if signal := qt.GetSignal(C.GoString(ptrName), "notification2"); signal != nil {
		signal.(func(string, QSqlDriver__NotificationSource, *core.QVariant))(C.GoString(name), QSqlDriver__NotificationSource(source), core.NewQVariantFromPointer(payload))
	}

}

func (ptr *QSqlDriver) Notification2(name string, source QSqlDriver__NotificationSource, payload core.QVariant_ITF) {
	defer qt.Recovering("QSqlDriver::notification")

	if ptr.Pointer() != nil {
		C.QSqlDriver_Notification2(ptr.Pointer(), C.CString(name), C.int(source), core.PointerFromQVariant(payload))
	}
}

func (ptr *QSqlDriver) Open(db string, user string, password string, host string, port int, options string) bool {
	defer qt.Recovering("QSqlDriver::open")

	if ptr.Pointer() != nil {
		return C.QSqlDriver_Open(ptr.Pointer(), C.CString(db), C.CString(user), C.CString(password), C.CString(host), C.int(port), C.CString(options)) != 0
	}
	return false
}

func (ptr *QSqlDriver) RollbackTransaction() bool {
	defer qt.Recovering("QSqlDriver::rollbackTransaction")

	if ptr.Pointer() != nil {
		return C.QSqlDriver_RollbackTransaction(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlDriver) ConnectSetOpen(f func(open bool)) {
	defer qt.Recovering("connect QSqlDriver::setOpen")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setOpen", f)
	}
}

func (ptr *QSqlDriver) DisconnectSetOpen() {
	defer qt.Recovering("disconnect QSqlDriver::setOpen")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setOpen")
	}
}

//export callbackQSqlDriverSetOpen
func callbackQSqlDriverSetOpen(ptr unsafe.Pointer, ptrName *C.char, open C.int) {
	defer qt.Recovering("callback QSqlDriver::setOpen")

	if signal := qt.GetSignal(C.GoString(ptrName), "setOpen"); signal != nil {
		signal.(func(bool))(int(open) != 0)
	} else {
		NewQSqlDriverFromPointer(ptr).SetOpenDefault(int(open) != 0)
	}
}

func (ptr *QSqlDriver) SetOpen(open bool) {
	defer qt.Recovering("QSqlDriver::setOpen")

	if ptr.Pointer() != nil {
		C.QSqlDriver_SetOpen(ptr.Pointer(), C.int(qt.GoBoolToInt(open)))
	}
}

func (ptr *QSqlDriver) SetOpenDefault(open bool) {
	defer qt.Recovering("QSqlDriver::setOpen")

	if ptr.Pointer() != nil {
		C.QSqlDriver_SetOpenDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(open)))
	}
}

func (ptr *QSqlDriver) ConnectSetOpenError(f func(error bool)) {
	defer qt.Recovering("connect QSqlDriver::setOpenError")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setOpenError", f)
	}
}

func (ptr *QSqlDriver) DisconnectSetOpenError() {
	defer qt.Recovering("disconnect QSqlDriver::setOpenError")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setOpenError")
	}
}

//export callbackQSqlDriverSetOpenError
func callbackQSqlDriverSetOpenError(ptr unsafe.Pointer, ptrName *C.char, error C.int) {
	defer qt.Recovering("callback QSqlDriver::setOpenError")

	if signal := qt.GetSignal(C.GoString(ptrName), "setOpenError"); signal != nil {
		signal.(func(bool))(int(error) != 0)
	} else {
		NewQSqlDriverFromPointer(ptr).SetOpenErrorDefault(int(error) != 0)
	}
}

func (ptr *QSqlDriver) SetOpenError(error bool) {
	defer qt.Recovering("QSqlDriver::setOpenError")

	if ptr.Pointer() != nil {
		C.QSqlDriver_SetOpenError(ptr.Pointer(), C.int(qt.GoBoolToInt(error)))
	}
}

func (ptr *QSqlDriver) SetOpenErrorDefault(error bool) {
	defer qt.Recovering("QSqlDriver::setOpenError")

	if ptr.Pointer() != nil {
		C.QSqlDriver_SetOpenErrorDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(error)))
	}
}

func (ptr *QSqlDriver) SqlStatement(ty QSqlDriver__StatementType, tableName string, rec QSqlRecord_ITF, preparedStatement bool) string {
	defer qt.Recovering("QSqlDriver::sqlStatement")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlDriver_SqlStatement(ptr.Pointer(), C.int(ty), C.CString(tableName), PointerFromQSqlRecord(rec), C.int(qt.GoBoolToInt(preparedStatement))))
	}
	return ""
}

func (ptr *QSqlDriver) StripDelimiters(identifier string, ty QSqlDriver__IdentifierType) string {
	defer qt.Recovering("QSqlDriver::stripDelimiters")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlDriver_StripDelimiters(ptr.Pointer(), C.CString(identifier), C.int(ty)))
	}
	return ""
}

func (ptr *QSqlDriver) SubscribeToNotification(name string) bool {
	defer qt.Recovering("QSqlDriver::subscribeToNotification")

	if ptr.Pointer() != nil {
		return C.QSqlDriver_SubscribeToNotification(ptr.Pointer(), C.CString(name)) != 0
	}
	return false
}

func (ptr *QSqlDriver) SubscribedToNotifications() []string {
	defer qt.Recovering("QSqlDriver::subscribedToNotifications")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QSqlDriver_SubscribedToNotifications(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QSqlDriver) UnsubscribeFromNotification(name string) bool {
	defer qt.Recovering("QSqlDriver::unsubscribeFromNotification")

	if ptr.Pointer() != nil {
		return C.QSqlDriver_UnsubscribeFromNotification(ptr.Pointer(), C.CString(name)) != 0
	}
	return false
}

func (ptr *QSqlDriver) DestroyQSqlDriver() {
	defer qt.Recovering("QSqlDriver::~QSqlDriver")

	if ptr.Pointer() != nil {
		C.QSqlDriver_DestroyQSqlDriver(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSqlDriver) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QSqlDriver::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QSqlDriver) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QSqlDriver::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQSqlDriverTimerEvent
func callbackQSqlDriverTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSqlDriver::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSqlDriverFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSqlDriver) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSqlDriver::timerEvent")

	if ptr.Pointer() != nil {
		C.QSqlDriver_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSqlDriver) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSqlDriver::timerEvent")

	if ptr.Pointer() != nil {
		C.QSqlDriver_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSqlDriver) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QSqlDriver::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QSqlDriver) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QSqlDriver::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQSqlDriverChildEvent
func callbackQSqlDriverChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSqlDriver::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSqlDriverFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSqlDriver) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSqlDriver::childEvent")

	if ptr.Pointer() != nil {
		C.QSqlDriver_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSqlDriver) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSqlDriver::childEvent")

	if ptr.Pointer() != nil {
		C.QSqlDriver_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSqlDriver) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSqlDriver::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QSqlDriver) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QSqlDriver::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQSqlDriverCustomEvent
func callbackQSqlDriverCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSqlDriver::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSqlDriverFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSqlDriver) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSqlDriver::customEvent")

	if ptr.Pointer() != nil {
		C.QSqlDriver_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSqlDriver) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSqlDriver::customEvent")

	if ptr.Pointer() != nil {
		C.QSqlDriver_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QSqlDriverCreator struct {
	QSqlDriverCreatorBase
}

type QSqlDriverCreator_ITF interface {
	QSqlDriverCreatorBase_ITF
	QSqlDriverCreator_PTR() *QSqlDriverCreator
}

func PointerFromQSqlDriverCreator(ptr QSqlDriverCreator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlDriverCreator_PTR().Pointer()
	}
	return nil
}

func NewQSqlDriverCreatorFromPointer(ptr unsafe.Pointer) *QSqlDriverCreator {
	var n = new(QSqlDriverCreator)
	n.SetPointer(ptr)
	return n
}

func newQSqlDriverCreatorFromPointer(ptr unsafe.Pointer) *QSqlDriverCreator {
	var n = NewQSqlDriverCreatorFromPointer(ptr)
	return n
}

func (ptr *QSqlDriverCreator) QSqlDriverCreator_PTR() *QSqlDriverCreator {
	return ptr
}

type QSqlDriverCreatorBase struct {
	ptr unsafe.Pointer
}

type QSqlDriverCreatorBase_ITF interface {
	QSqlDriverCreatorBase_PTR() *QSqlDriverCreatorBase
}

func (p *QSqlDriverCreatorBase) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSqlDriverCreatorBase) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSqlDriverCreatorBase(ptr QSqlDriverCreatorBase_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlDriverCreatorBase_PTR().Pointer()
	}
	return nil
}

func NewQSqlDriverCreatorBaseFromPointer(ptr unsafe.Pointer) *QSqlDriverCreatorBase {
	var n = new(QSqlDriverCreatorBase)
	n.SetPointer(ptr)
	return n
}

func newQSqlDriverCreatorBaseFromPointer(ptr unsafe.Pointer) *QSqlDriverCreatorBase {
	var n = NewQSqlDriverCreatorBaseFromPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QSqlDriverCreatorBase_") {
		n.SetObjectNameAbs("QSqlDriverCreatorBase_" + qt.Identifier())
	}
	return n
}

func (ptr *QSqlDriverCreatorBase) QSqlDriverCreatorBase_PTR() *QSqlDriverCreatorBase {
	return ptr
}

func (ptr *QSqlDriverCreatorBase) CreateObject() *QSqlDriver {
	defer qt.Recovering("QSqlDriverCreatorBase::createObject")

	if ptr.Pointer() != nil {
		return NewQSqlDriverFromPointer(C.QSqlDriverCreatorBase_CreateObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSqlDriverCreatorBase) DestroyQSqlDriverCreatorBase() {
	defer qt.Recovering("QSqlDriverCreatorBase::~QSqlDriverCreatorBase")

	if ptr.Pointer() != nil {
		C.QSqlDriverCreatorBase_DestroyQSqlDriverCreatorBase(ptr.Pointer())
	}
}

func (ptr *QSqlDriverCreatorBase) ObjectNameAbs() string {
	defer qt.Recovering("QSqlDriverCreatorBase::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlDriverCreatorBase_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlDriverCreatorBase) SetObjectNameAbs(name string) {
	defer qt.Recovering("QSqlDriverCreatorBase::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QSqlDriverCreatorBase_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}

type QSqlDriverPlugin struct {
	core.QObject
}

type QSqlDriverPlugin_ITF interface {
	core.QObject_ITF
	QSqlDriverPlugin_PTR() *QSqlDriverPlugin
}

func PointerFromQSqlDriverPlugin(ptr QSqlDriverPlugin_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlDriverPlugin_PTR().Pointer()
	}
	return nil
}

func NewQSqlDriverPluginFromPointer(ptr unsafe.Pointer) *QSqlDriverPlugin {
	var n = new(QSqlDriverPlugin)
	n.SetPointer(ptr)
	return n
}

func newQSqlDriverPluginFromPointer(ptr unsafe.Pointer) *QSqlDriverPlugin {
	var n = NewQSqlDriverPluginFromPointer(ptr)
	for len(n.ObjectName()) < len("QSqlDriverPlugin_") {
		n.SetObjectName("QSqlDriverPlugin_" + qt.Identifier())
	}
	return n
}

func (ptr *QSqlDriverPlugin) QSqlDriverPlugin_PTR() *QSqlDriverPlugin {
	return ptr
}

func (ptr *QSqlDriverPlugin) Create(key string) *QSqlDriver {
	defer qt.Recovering("QSqlDriverPlugin::create")

	if ptr.Pointer() != nil {
		return NewQSqlDriverFromPointer(C.QSqlDriverPlugin_Create(ptr.Pointer(), C.CString(key)))
	}
	return nil
}

func (ptr *QSqlDriverPlugin) DestroyQSqlDriverPlugin() {
	defer qt.Recovering("QSqlDriverPlugin::~QSqlDriverPlugin")

	if ptr.Pointer() != nil {
		C.QSqlDriverPlugin_DestroyQSqlDriverPlugin(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSqlDriverPlugin) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QSqlDriverPlugin::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QSqlDriverPlugin) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QSqlDriverPlugin::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQSqlDriverPluginTimerEvent
func callbackQSqlDriverPluginTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSqlDriverPlugin::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSqlDriverPluginFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSqlDriverPlugin) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSqlDriverPlugin::timerEvent")

	if ptr.Pointer() != nil {
		C.QSqlDriverPlugin_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSqlDriverPlugin) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSqlDriverPlugin::timerEvent")

	if ptr.Pointer() != nil {
		C.QSqlDriverPlugin_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSqlDriverPlugin) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QSqlDriverPlugin::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QSqlDriverPlugin) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QSqlDriverPlugin::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQSqlDriverPluginChildEvent
func callbackQSqlDriverPluginChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSqlDriverPlugin::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSqlDriverPluginFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSqlDriverPlugin) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSqlDriverPlugin::childEvent")

	if ptr.Pointer() != nil {
		C.QSqlDriverPlugin_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSqlDriverPlugin) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSqlDriverPlugin::childEvent")

	if ptr.Pointer() != nil {
		C.QSqlDriverPlugin_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSqlDriverPlugin) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSqlDriverPlugin::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QSqlDriverPlugin) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QSqlDriverPlugin::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQSqlDriverPluginCustomEvent
func callbackQSqlDriverPluginCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSqlDriverPlugin::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSqlDriverPluginFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSqlDriverPlugin) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSqlDriverPlugin::customEvent")

	if ptr.Pointer() != nil {
		C.QSqlDriverPlugin_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSqlDriverPlugin) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSqlDriverPlugin::customEvent")

	if ptr.Pointer() != nil {
		C.QSqlDriverPlugin_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QSqlError struct {
	ptr unsafe.Pointer
}

type QSqlError_ITF interface {
	QSqlError_PTR() *QSqlError
}

func (p *QSqlError) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSqlError) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSqlError(ptr QSqlError_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlError_PTR().Pointer()
	}
	return nil
}

func NewQSqlErrorFromPointer(ptr unsafe.Pointer) *QSqlError {
	var n = new(QSqlError)
	n.SetPointer(ptr)
	return n
}

func newQSqlErrorFromPointer(ptr unsafe.Pointer) *QSqlError {
	var n = NewQSqlErrorFromPointer(ptr)
	return n
}

func (ptr *QSqlError) QSqlError_PTR() *QSqlError {
	return ptr
}

//QSqlError::ErrorType
type QSqlError__ErrorType int64

const (
	QSqlError__NoError          = QSqlError__ErrorType(0)
	QSqlError__ConnectionError  = QSqlError__ErrorType(1)
	QSqlError__StatementError   = QSqlError__ErrorType(2)
	QSqlError__TransactionError = QSqlError__ErrorType(3)
	QSqlError__UnknownError     = QSqlError__ErrorType(4)
)

func NewQSqlError3(other QSqlError_ITF) *QSqlError {
	defer qt.Recovering("QSqlError::QSqlError")

	return newQSqlErrorFromPointer(C.QSqlError_NewQSqlError3(PointerFromQSqlError(other)))
}

func NewQSqlError(driverText string, databaseText string, ty QSqlError__ErrorType, code string) *QSqlError {
	defer qt.Recovering("QSqlError::QSqlError")

	return newQSqlErrorFromPointer(C.QSqlError_NewQSqlError(C.CString(driverText), C.CString(databaseText), C.int(ty), C.CString(code)))
}

func (ptr *QSqlError) DatabaseText() string {
	defer qt.Recovering("QSqlError::databaseText")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlError_DatabaseText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlError) DriverText() string {
	defer qt.Recovering("QSqlError::driverText")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlError_DriverText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlError) IsValid() bool {
	defer qt.Recovering("QSqlError::isValid")

	if ptr.Pointer() != nil {
		return C.QSqlError_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlError) NativeErrorCode() string {
	defer qt.Recovering("QSqlError::nativeErrorCode")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlError_NativeErrorCode(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlError) Text() string {
	defer qt.Recovering("QSqlError::text")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlError_Text(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlError) Type() QSqlError__ErrorType {
	defer qt.Recovering("QSqlError::type")

	if ptr.Pointer() != nil {
		return QSqlError__ErrorType(C.QSqlError_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSqlError) DestroyQSqlError() {
	defer qt.Recovering("QSqlError::~QSqlError")

	if ptr.Pointer() != nil {
		C.QSqlError_DestroyQSqlError(ptr.Pointer())
	}
}

type QSqlField struct {
	ptr unsafe.Pointer
}

type QSqlField_ITF interface {
	QSqlField_PTR() *QSqlField
}

func (p *QSqlField) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSqlField) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSqlField(ptr QSqlField_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlField_PTR().Pointer()
	}
	return nil
}

func NewQSqlFieldFromPointer(ptr unsafe.Pointer) *QSqlField {
	var n = new(QSqlField)
	n.SetPointer(ptr)
	return n
}

func newQSqlFieldFromPointer(ptr unsafe.Pointer) *QSqlField {
	var n = NewQSqlFieldFromPointer(ptr)
	return n
}

func (ptr *QSqlField) QSqlField_PTR() *QSqlField {
	return ptr
}

//QSqlField::RequiredStatus
type QSqlField__RequiredStatus int64

const (
	QSqlField__Unknown  = QSqlField__RequiredStatus(-1)
	QSqlField__Optional = QSqlField__RequiredStatus(0)
	QSqlField__Required = QSqlField__RequiredStatus(1)
)

func NewQSqlField2(other QSqlField_ITF) *QSqlField {
	defer qt.Recovering("QSqlField::QSqlField")

	return newQSqlFieldFromPointer(C.QSqlField_NewQSqlField2(PointerFromQSqlField(other)))
}

func (ptr *QSqlField) Clear() {
	defer qt.Recovering("QSqlField::clear")

	if ptr.Pointer() != nil {
		C.QSqlField_Clear(ptr.Pointer())
	}
}

func (ptr *QSqlField) DefaultValue() *core.QVariant {
	defer qt.Recovering("QSqlField::defaultValue")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QSqlField_DefaultValue(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSqlField) IsAutoValue() bool {
	defer qt.Recovering("QSqlField::isAutoValue")

	if ptr.Pointer() != nil {
		return C.QSqlField_IsAutoValue(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlField) IsGenerated() bool {
	defer qt.Recovering("QSqlField::isGenerated")

	if ptr.Pointer() != nil {
		return C.QSqlField_IsGenerated(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlField) IsNull() bool {
	defer qt.Recovering("QSqlField::isNull")

	if ptr.Pointer() != nil {
		return C.QSqlField_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlField) IsReadOnly() bool {
	defer qt.Recovering("QSqlField::isReadOnly")

	if ptr.Pointer() != nil {
		return C.QSqlField_IsReadOnly(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlField) IsValid() bool {
	defer qt.Recovering("QSqlField::isValid")

	if ptr.Pointer() != nil {
		return C.QSqlField_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlField) Length() int {
	defer qt.Recovering("QSqlField::length")

	if ptr.Pointer() != nil {
		return int(C.QSqlField_Length(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSqlField) Name() string {
	defer qt.Recovering("QSqlField::name")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlField_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlField) Precision() int {
	defer qt.Recovering("QSqlField::precision")

	if ptr.Pointer() != nil {
		return int(C.QSqlField_Precision(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSqlField) RequiredStatus() QSqlField__RequiredStatus {
	defer qt.Recovering("QSqlField::requiredStatus")

	if ptr.Pointer() != nil {
		return QSqlField__RequiredStatus(C.QSqlField_RequiredStatus(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSqlField) SetAutoValue(autoVal bool) {
	defer qt.Recovering("QSqlField::setAutoValue")

	if ptr.Pointer() != nil {
		C.QSqlField_SetAutoValue(ptr.Pointer(), C.int(qt.GoBoolToInt(autoVal)))
	}
}

func (ptr *QSqlField) SetDefaultValue(value core.QVariant_ITF) {
	defer qt.Recovering("QSqlField::setDefaultValue")

	if ptr.Pointer() != nil {
		C.QSqlField_SetDefaultValue(ptr.Pointer(), core.PointerFromQVariant(value))
	}
}

func (ptr *QSqlField) SetGenerated(gen bool) {
	defer qt.Recovering("QSqlField::setGenerated")

	if ptr.Pointer() != nil {
		C.QSqlField_SetGenerated(ptr.Pointer(), C.int(qt.GoBoolToInt(gen)))
	}
}

func (ptr *QSqlField) SetLength(fieldLength int) {
	defer qt.Recovering("QSqlField::setLength")

	if ptr.Pointer() != nil {
		C.QSqlField_SetLength(ptr.Pointer(), C.int(fieldLength))
	}
}

func (ptr *QSqlField) SetName(name string) {
	defer qt.Recovering("QSqlField::setName")

	if ptr.Pointer() != nil {
		C.QSqlField_SetName(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QSqlField) SetPrecision(precision int) {
	defer qt.Recovering("QSqlField::setPrecision")

	if ptr.Pointer() != nil {
		C.QSqlField_SetPrecision(ptr.Pointer(), C.int(precision))
	}
}

func (ptr *QSqlField) SetReadOnly(readOnly bool) {
	defer qt.Recovering("QSqlField::setReadOnly")

	if ptr.Pointer() != nil {
		C.QSqlField_SetReadOnly(ptr.Pointer(), C.int(qt.GoBoolToInt(readOnly)))
	}
}

func (ptr *QSqlField) SetRequired(required bool) {
	defer qt.Recovering("QSqlField::setRequired")

	if ptr.Pointer() != nil {
		C.QSqlField_SetRequired(ptr.Pointer(), C.int(qt.GoBoolToInt(required)))
	}
}

func (ptr *QSqlField) SetRequiredStatus(required QSqlField__RequiredStatus) {
	defer qt.Recovering("QSqlField::setRequiredStatus")

	if ptr.Pointer() != nil {
		C.QSqlField_SetRequiredStatus(ptr.Pointer(), C.int(required))
	}
}

func (ptr *QSqlField) SetValue(value core.QVariant_ITF) {
	defer qt.Recovering("QSqlField::setValue")

	if ptr.Pointer() != nil {
		C.QSqlField_SetValue(ptr.Pointer(), core.PointerFromQVariant(value))
	}
}

func (ptr *QSqlField) Value() *core.QVariant {
	defer qt.Recovering("QSqlField::value")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QSqlField_Value(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSqlField) DestroyQSqlField() {
	defer qt.Recovering("QSqlField::~QSqlField")

	if ptr.Pointer() != nil {
		C.QSqlField_DestroyQSqlField(ptr.Pointer())
	}
}

type QSqlIndex struct {
	QSqlRecord
}

type QSqlIndex_ITF interface {
	QSqlRecord_ITF
	QSqlIndex_PTR() *QSqlIndex
}

func PointerFromQSqlIndex(ptr QSqlIndex_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlIndex_PTR().Pointer()
	}
	return nil
}

func NewQSqlIndexFromPointer(ptr unsafe.Pointer) *QSqlIndex {
	var n = new(QSqlIndex)
	n.SetPointer(ptr)
	return n
}

func newQSqlIndexFromPointer(ptr unsafe.Pointer) *QSqlIndex {
	var n = NewQSqlIndexFromPointer(ptr)
	return n
}

func (ptr *QSqlIndex) QSqlIndex_PTR() *QSqlIndex {
	return ptr
}

func NewQSqlIndex2(other QSqlIndex_ITF) *QSqlIndex {
	defer qt.Recovering("QSqlIndex::QSqlIndex")

	return newQSqlIndexFromPointer(C.QSqlIndex_NewQSqlIndex2(PointerFromQSqlIndex(other)))
}

func NewQSqlIndex(cursorname string, name string) *QSqlIndex {
	defer qt.Recovering("QSqlIndex::QSqlIndex")

	return newQSqlIndexFromPointer(C.QSqlIndex_NewQSqlIndex(C.CString(cursorname), C.CString(name)))
}

func (ptr *QSqlIndex) Append(field QSqlField_ITF) {
	defer qt.Recovering("QSqlIndex::append")

	if ptr.Pointer() != nil {
		C.QSqlIndex_Append(ptr.Pointer(), PointerFromQSqlField(field))
	}
}

func (ptr *QSqlIndex) Append2(field QSqlField_ITF, desc bool) {
	defer qt.Recovering("QSqlIndex::append")

	if ptr.Pointer() != nil {
		C.QSqlIndex_Append2(ptr.Pointer(), PointerFromQSqlField(field), C.int(qt.GoBoolToInt(desc)))
	}
}

func (ptr *QSqlIndex) CursorName() string {
	defer qt.Recovering("QSqlIndex::cursorName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlIndex_CursorName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlIndex) IsDescending(i int) bool {
	defer qt.Recovering("QSqlIndex::isDescending")

	if ptr.Pointer() != nil {
		return C.QSqlIndex_IsDescending(ptr.Pointer(), C.int(i)) != 0
	}
	return false
}

func (ptr *QSqlIndex) Name() string {
	defer qt.Recovering("QSqlIndex::name")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlIndex_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlIndex) SetCursorName(cursorName string) {
	defer qt.Recovering("QSqlIndex::setCursorName")

	if ptr.Pointer() != nil {
		C.QSqlIndex_SetCursorName(ptr.Pointer(), C.CString(cursorName))
	}
}

func (ptr *QSqlIndex) SetDescending(i int, desc bool) {
	defer qt.Recovering("QSqlIndex::setDescending")

	if ptr.Pointer() != nil {
		C.QSqlIndex_SetDescending(ptr.Pointer(), C.int(i), C.int(qt.GoBoolToInt(desc)))
	}
}

func (ptr *QSqlIndex) SetName(name string) {
	defer qt.Recovering("QSqlIndex::setName")

	if ptr.Pointer() != nil {
		C.QSqlIndex_SetName(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QSqlIndex) DestroyQSqlIndex() {
	defer qt.Recovering("QSqlIndex::~QSqlIndex")

	if ptr.Pointer() != nil {
		C.QSqlIndex_DestroyQSqlIndex(ptr.Pointer())
	}
}

type QSqlQuery struct {
	ptr unsafe.Pointer
}

type QSqlQuery_ITF interface {
	QSqlQuery_PTR() *QSqlQuery
}

func (p *QSqlQuery) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSqlQuery) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSqlQuery(ptr QSqlQuery_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlQuery_PTR().Pointer()
	}
	return nil
}

func NewQSqlQueryFromPointer(ptr unsafe.Pointer) *QSqlQuery {
	var n = new(QSqlQuery)
	n.SetPointer(ptr)
	return n
}

func newQSqlQueryFromPointer(ptr unsafe.Pointer) *QSqlQuery {
	var n = NewQSqlQueryFromPointer(ptr)
	return n
}

func (ptr *QSqlQuery) QSqlQuery_PTR() *QSqlQuery {
	return ptr
}

//QSqlQuery::BatchExecutionMode
type QSqlQuery__BatchExecutionMode int64

const (
	QSqlQuery__ValuesAsRows    = QSqlQuery__BatchExecutionMode(0)
	QSqlQuery__ValuesAsColumns = QSqlQuery__BatchExecutionMode(1)
)

func NewQSqlQuery3(db QSqlDatabase_ITF) *QSqlQuery {
	defer qt.Recovering("QSqlQuery::QSqlQuery")

	return newQSqlQueryFromPointer(C.QSqlQuery_NewQSqlQuery3(PointerFromQSqlDatabase(db)))
}

func NewQSqlQuery(result QSqlResult_ITF) *QSqlQuery {
	defer qt.Recovering("QSqlQuery::QSqlQuery")

	return newQSqlQueryFromPointer(C.QSqlQuery_NewQSqlQuery(PointerFromQSqlResult(result)))
}

func NewQSqlQuery4(other QSqlQuery_ITF) *QSqlQuery {
	defer qt.Recovering("QSqlQuery::QSqlQuery")

	return newQSqlQueryFromPointer(C.QSqlQuery_NewQSqlQuery4(PointerFromQSqlQuery(other)))
}

func NewQSqlQuery2(query string, db QSqlDatabase_ITF) *QSqlQuery {
	defer qt.Recovering("QSqlQuery::QSqlQuery")

	return newQSqlQueryFromPointer(C.QSqlQuery_NewQSqlQuery2(C.CString(query), PointerFromQSqlDatabase(db)))
}

func (ptr *QSqlQuery) At() int {
	defer qt.Recovering("QSqlQuery::at")

	if ptr.Pointer() != nil {
		return int(C.QSqlQuery_At(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSqlQuery) BoundValue(placeholder string) *core.QVariant {
	defer qt.Recovering("QSqlQuery::boundValue")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QSqlQuery_BoundValue(ptr.Pointer(), C.CString(placeholder)))
	}
	return nil
}

func (ptr *QSqlQuery) BoundValue2(pos int) *core.QVariant {
	defer qt.Recovering("QSqlQuery::boundValue")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QSqlQuery_BoundValue2(ptr.Pointer(), C.int(pos)))
	}
	return nil
}

func (ptr *QSqlQuery) Clear() {
	defer qt.Recovering("QSqlQuery::clear")

	if ptr.Pointer() != nil {
		C.QSqlQuery_Clear(ptr.Pointer())
	}
}

func (ptr *QSqlQuery) Driver() *QSqlDriver {
	defer qt.Recovering("QSqlQuery::driver")

	if ptr.Pointer() != nil {
		return NewQSqlDriverFromPointer(C.QSqlQuery_Driver(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSqlQuery) Exec2() bool {
	defer qt.Recovering("QSqlQuery::exec")

	if ptr.Pointer() != nil {
		return C.QSqlQuery_Exec2(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlQuery) Exec(query string) bool {
	defer qt.Recovering("QSqlQuery::exec")

	if ptr.Pointer() != nil {
		return C.QSqlQuery_Exec(ptr.Pointer(), C.CString(query)) != 0
	}
	return false
}

func (ptr *QSqlQuery) ExecBatch(mode QSqlQuery__BatchExecutionMode) bool {
	defer qt.Recovering("QSqlQuery::execBatch")

	if ptr.Pointer() != nil {
		return C.QSqlQuery_ExecBatch(ptr.Pointer(), C.int(mode)) != 0
	}
	return false
}

func (ptr *QSqlQuery) ExecutedQuery() string {
	defer qt.Recovering("QSqlQuery::executedQuery")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlQuery_ExecutedQuery(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlQuery) Finish() {
	defer qt.Recovering("QSqlQuery::finish")

	if ptr.Pointer() != nil {
		C.QSqlQuery_Finish(ptr.Pointer())
	}
}

func (ptr *QSqlQuery) First() bool {
	defer qt.Recovering("QSqlQuery::first")

	if ptr.Pointer() != nil {
		return C.QSqlQuery_First(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlQuery) IsActive() bool {
	defer qt.Recovering("QSqlQuery::isActive")

	if ptr.Pointer() != nil {
		return C.QSqlQuery_IsActive(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlQuery) IsForwardOnly() bool {
	defer qt.Recovering("QSqlQuery::isForwardOnly")

	if ptr.Pointer() != nil {
		return C.QSqlQuery_IsForwardOnly(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlQuery) IsNull2(name string) bool {
	defer qt.Recovering("QSqlQuery::isNull")

	if ptr.Pointer() != nil {
		return C.QSqlQuery_IsNull2(ptr.Pointer(), C.CString(name)) != 0
	}
	return false
}

func (ptr *QSqlQuery) IsNull(field int) bool {
	defer qt.Recovering("QSqlQuery::isNull")

	if ptr.Pointer() != nil {
		return C.QSqlQuery_IsNull(ptr.Pointer(), C.int(field)) != 0
	}
	return false
}

func (ptr *QSqlQuery) IsSelect() bool {
	defer qt.Recovering("QSqlQuery::isSelect")

	if ptr.Pointer() != nil {
		return C.QSqlQuery_IsSelect(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlQuery) IsValid() bool {
	defer qt.Recovering("QSqlQuery::isValid")

	if ptr.Pointer() != nil {
		return C.QSqlQuery_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlQuery) Last() bool {
	defer qt.Recovering("QSqlQuery::last")

	if ptr.Pointer() != nil {
		return C.QSqlQuery_Last(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlQuery) LastInsertId() *core.QVariant {
	defer qt.Recovering("QSqlQuery::lastInsertId")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QSqlQuery_LastInsertId(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSqlQuery) LastQuery() string {
	defer qt.Recovering("QSqlQuery::lastQuery")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlQuery_LastQuery(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlQuery) Next() bool {
	defer qt.Recovering("QSqlQuery::next")

	if ptr.Pointer() != nil {
		return C.QSqlQuery_Next(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlQuery) NextResult() bool {
	defer qt.Recovering("QSqlQuery::nextResult")

	if ptr.Pointer() != nil {
		return C.QSqlQuery_NextResult(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlQuery) NumRowsAffected() int {
	defer qt.Recovering("QSqlQuery::numRowsAffected")

	if ptr.Pointer() != nil {
		return int(C.QSqlQuery_NumRowsAffected(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSqlQuery) Prepare(query string) bool {
	defer qt.Recovering("QSqlQuery::prepare")

	if ptr.Pointer() != nil {
		return C.QSqlQuery_Prepare(ptr.Pointer(), C.CString(query)) != 0
	}
	return false
}

func (ptr *QSqlQuery) Previous() bool {
	defer qt.Recovering("QSqlQuery::previous")

	if ptr.Pointer() != nil {
		return C.QSqlQuery_Previous(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlQuery) Result() *QSqlResult {
	defer qt.Recovering("QSqlQuery::result")

	if ptr.Pointer() != nil {
		return NewQSqlResultFromPointer(C.QSqlQuery_Result(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSqlQuery) Seek(index int, relative bool) bool {
	defer qt.Recovering("QSqlQuery::seek")

	if ptr.Pointer() != nil {
		return C.QSqlQuery_Seek(ptr.Pointer(), C.int(index), C.int(qt.GoBoolToInt(relative))) != 0
	}
	return false
}

func (ptr *QSqlQuery) SetForwardOnly(forward bool) {
	defer qt.Recovering("QSqlQuery::setForwardOnly")

	if ptr.Pointer() != nil {
		C.QSqlQuery_SetForwardOnly(ptr.Pointer(), C.int(qt.GoBoolToInt(forward)))
	}
}

func (ptr *QSqlQuery) Size() int {
	defer qt.Recovering("QSqlQuery::size")

	if ptr.Pointer() != nil {
		return int(C.QSqlQuery_Size(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSqlQuery) Value2(name string) *core.QVariant {
	defer qt.Recovering("QSqlQuery::value")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QSqlQuery_Value2(ptr.Pointer(), C.CString(name)))
	}
	return nil
}

func (ptr *QSqlQuery) Value(index int) *core.QVariant {
	defer qt.Recovering("QSqlQuery::value")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QSqlQuery_Value(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QSqlQuery) DestroyQSqlQuery() {
	defer qt.Recovering("QSqlQuery::~QSqlQuery")

	if ptr.Pointer() != nil {
		C.QSqlQuery_DestroyQSqlQuery(ptr.Pointer())
	}
}

type QSqlQueryModel struct {
	core.QAbstractTableModel
}

type QSqlQueryModel_ITF interface {
	core.QAbstractTableModel_ITF
	QSqlQueryModel_PTR() *QSqlQueryModel
}

func PointerFromQSqlQueryModel(ptr QSqlQueryModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlQueryModel_PTR().Pointer()
	}
	return nil
}

func NewQSqlQueryModelFromPointer(ptr unsafe.Pointer) *QSqlQueryModel {
	var n = new(QSqlQueryModel)
	n.SetPointer(ptr)
	return n
}

func newQSqlQueryModelFromPointer(ptr unsafe.Pointer) *QSqlQueryModel {
	var n = NewQSqlQueryModelFromPointer(ptr)
	for len(n.ObjectName()) < len("QSqlQueryModel_") {
		n.SetObjectName("QSqlQueryModel_" + qt.Identifier())
	}
	return n
}

func (ptr *QSqlQueryModel) QSqlQueryModel_PTR() *QSqlQueryModel {
	return ptr
}

func (ptr *QSqlQueryModel) RowCount(parent core.QModelIndex_ITF) int {
	defer qt.Recovering("QSqlQueryModel::rowCount")

	if ptr.Pointer() != nil {
		return int(C.QSqlQueryModel_RowCount(ptr.Pointer(), core.PointerFromQModelIndex(parent)))
	}
	return 0
}

func (ptr *QSqlQueryModel) Data(item core.QModelIndex_ITF, role int) *core.QVariant {
	defer qt.Recovering("QSqlQueryModel::data")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QSqlQueryModel_Data(ptr.Pointer(), core.PointerFromQModelIndex(item), C.int(role)))
	}
	return nil
}

func (ptr *QSqlQueryModel) CanFetchMore(parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QSqlQueryModel::canFetchMore")

	if ptr.Pointer() != nil {
		return C.QSqlQueryModel_CanFetchMore(ptr.Pointer(), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QSqlQueryModel) ConnectClear(f func()) {
	defer qt.Recovering("connect QSqlQueryModel::clear")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "clear", f)
	}
}

func (ptr *QSqlQueryModel) DisconnectClear() {
	defer qt.Recovering("disconnect QSqlQueryModel::clear")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "clear")
	}
}

//export callbackQSqlQueryModelClear
func callbackQSqlQueryModelClear(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSqlQueryModel::clear")

	if signal := qt.GetSignal(C.GoString(ptrName), "clear"); signal != nil {
		signal.(func())()
	} else {
		NewQSqlQueryModelFromPointer(ptr).ClearDefault()
	}
}

func (ptr *QSqlQueryModel) Clear() {
	defer qt.Recovering("QSqlQueryModel::clear")

	if ptr.Pointer() != nil {
		C.QSqlQueryModel_Clear(ptr.Pointer())
	}
}

func (ptr *QSqlQueryModel) ClearDefault() {
	defer qt.Recovering("QSqlQueryModel::clear")

	if ptr.Pointer() != nil {
		C.QSqlQueryModel_ClearDefault(ptr.Pointer())
	}
}

func (ptr *QSqlQueryModel) ColumnCount(index core.QModelIndex_ITF) int {
	defer qt.Recovering("QSqlQueryModel::columnCount")

	if ptr.Pointer() != nil {
		return int(C.QSqlQueryModel_ColumnCount(ptr.Pointer(), core.PointerFromQModelIndex(index)))
	}
	return 0
}

func (ptr *QSqlQueryModel) ConnectFetchMore(f func(parent *core.QModelIndex)) {
	defer qt.Recovering("connect QSqlQueryModel::fetchMore")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "fetchMore", f)
	}
}

func (ptr *QSqlQueryModel) DisconnectFetchMore() {
	defer qt.Recovering("disconnect QSqlQueryModel::fetchMore")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "fetchMore")
	}
}

//export callbackQSqlQueryModelFetchMore
func callbackQSqlQueryModelFetchMore(ptr unsafe.Pointer, ptrName *C.char, parent unsafe.Pointer) {
	defer qt.Recovering("callback QSqlQueryModel::fetchMore")

	if signal := qt.GetSignal(C.GoString(ptrName), "fetchMore"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(parent))
	} else {
		NewQSqlQueryModelFromPointer(ptr).FetchMoreDefault(core.NewQModelIndexFromPointer(parent))
	}
}

func (ptr *QSqlQueryModel) FetchMore(parent core.QModelIndex_ITF) {
	defer qt.Recovering("QSqlQueryModel::fetchMore")

	if ptr.Pointer() != nil {
		C.QSqlQueryModel_FetchMore(ptr.Pointer(), core.PointerFromQModelIndex(parent))
	}
}

func (ptr *QSqlQueryModel) FetchMoreDefault(parent core.QModelIndex_ITF) {
	defer qt.Recovering("QSqlQueryModel::fetchMore")

	if ptr.Pointer() != nil {
		C.QSqlQueryModel_FetchMoreDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent))
	}
}

func (ptr *QSqlQueryModel) HeaderData(section int, orientation core.Qt__Orientation, role int) *core.QVariant {
	defer qt.Recovering("QSqlQueryModel::headerData")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QSqlQueryModel_HeaderData(ptr.Pointer(), C.int(section), C.int(orientation), C.int(role)))
	}
	return nil
}

func (ptr *QSqlQueryModel) IndexInQuery(item core.QModelIndex_ITF) *core.QModelIndex {
	defer qt.Recovering("QSqlQueryModel::indexInQuery")

	if ptr.Pointer() != nil {
		return core.NewQModelIndexFromPointer(C.QSqlQueryModel_IndexInQuery(ptr.Pointer(), core.PointerFromQModelIndex(item)))
	}
	return nil
}

func (ptr *QSqlQueryModel) InsertColumns(column int, count int, parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QSqlQueryModel::insertColumns")

	if ptr.Pointer() != nil {
		return C.QSqlQueryModel_InsertColumns(ptr.Pointer(), C.int(column), C.int(count), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QSqlQueryModel) ConnectQueryChange(f func()) {
	defer qt.Recovering("connect QSqlQueryModel::queryChange")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "queryChange", f)
	}
}

func (ptr *QSqlQueryModel) DisconnectQueryChange() {
	defer qt.Recovering("disconnect QSqlQueryModel::queryChange")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "queryChange")
	}
}

//export callbackQSqlQueryModelQueryChange
func callbackQSqlQueryModelQueryChange(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSqlQueryModel::queryChange")

	if signal := qt.GetSignal(C.GoString(ptrName), "queryChange"); signal != nil {
		signal.(func())()
	} else {
		NewQSqlQueryModelFromPointer(ptr).QueryChangeDefault()
	}
}

func (ptr *QSqlQueryModel) QueryChange() {
	defer qt.Recovering("QSqlQueryModel::queryChange")

	if ptr.Pointer() != nil {
		C.QSqlQueryModel_QueryChange(ptr.Pointer())
	}
}

func (ptr *QSqlQueryModel) QueryChangeDefault() {
	defer qt.Recovering("QSqlQueryModel::queryChange")

	if ptr.Pointer() != nil {
		C.QSqlQueryModel_QueryChangeDefault(ptr.Pointer())
	}
}

func (ptr *QSqlQueryModel) RemoveColumns(column int, count int, parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QSqlQueryModel::removeColumns")

	if ptr.Pointer() != nil {
		return C.QSqlQueryModel_RemoveColumns(ptr.Pointer(), C.int(column), C.int(count), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QSqlQueryModel) SetHeaderData(section int, orientation core.Qt__Orientation, value core.QVariant_ITF, role int) bool {
	defer qt.Recovering("QSqlQueryModel::setHeaderData")

	if ptr.Pointer() != nil {
		return C.QSqlQueryModel_SetHeaderData(ptr.Pointer(), C.int(section), C.int(orientation), core.PointerFromQVariant(value), C.int(role)) != 0
	}
	return false
}

func (ptr *QSqlQueryModel) SetQuery(query QSqlQuery_ITF) {
	defer qt.Recovering("QSqlQueryModel::setQuery")

	if ptr.Pointer() != nil {
		C.QSqlQueryModel_SetQuery(ptr.Pointer(), PointerFromQSqlQuery(query))
	}
}

func (ptr *QSqlQueryModel) SetQuery2(query string, db QSqlDatabase_ITF) {
	defer qt.Recovering("QSqlQueryModel::setQuery")

	if ptr.Pointer() != nil {
		C.QSqlQueryModel_SetQuery2(ptr.Pointer(), C.CString(query), PointerFromQSqlDatabase(db))
	}
}

func (ptr *QSqlQueryModel) DestroyQSqlQueryModel() {
	defer qt.Recovering("QSqlQueryModel::~QSqlQueryModel")

	if ptr.Pointer() != nil {
		C.QSqlQueryModel_DestroyQSqlQueryModel(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSqlQueryModel) ConnectRevert(f func()) {
	defer qt.Recovering("connect QSqlQueryModel::revert")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "revert", f)
	}
}

func (ptr *QSqlQueryModel) DisconnectRevert() {
	defer qt.Recovering("disconnect QSqlQueryModel::revert")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "revert")
	}
}

//export callbackQSqlQueryModelRevert
func callbackQSqlQueryModelRevert(ptr unsafe.Pointer, ptrName *C.char) bool {
	defer qt.Recovering("callback QSqlQueryModel::revert")

	if signal := qt.GetSignal(C.GoString(ptrName), "revert"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QSqlQueryModel) Revert() {
	defer qt.Recovering("QSqlQueryModel::revert")

	if ptr.Pointer() != nil {
		C.QSqlQueryModel_Revert(ptr.Pointer())
	}
}

func (ptr *QSqlQueryModel) RevertDefault() {
	defer qt.Recovering("QSqlQueryModel::revert")

	if ptr.Pointer() != nil {
		C.QSqlQueryModel_RevertDefault(ptr.Pointer())
	}
}

func (ptr *QSqlQueryModel) ConnectSort(f func(column int, order core.Qt__SortOrder)) {
	defer qt.Recovering("connect QSqlQueryModel::sort")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "sort", f)
	}
}

func (ptr *QSqlQueryModel) DisconnectSort() {
	defer qt.Recovering("disconnect QSqlQueryModel::sort")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "sort")
	}
}

//export callbackQSqlQueryModelSort
func callbackQSqlQueryModelSort(ptr unsafe.Pointer, ptrName *C.char, column C.int, order C.int) {
	defer qt.Recovering("callback QSqlQueryModel::sort")

	if signal := qt.GetSignal(C.GoString(ptrName), "sort"); signal != nil {
		signal.(func(int, core.Qt__SortOrder))(int(column), core.Qt__SortOrder(order))
	} else {
		NewQSqlQueryModelFromPointer(ptr).SortDefault(int(column), core.Qt__SortOrder(order))
	}
}

func (ptr *QSqlQueryModel) Sort(column int, order core.Qt__SortOrder) {
	defer qt.Recovering("QSqlQueryModel::sort")

	if ptr.Pointer() != nil {
		C.QSqlQueryModel_Sort(ptr.Pointer(), C.int(column), C.int(order))
	}
}

func (ptr *QSqlQueryModel) SortDefault(column int, order core.Qt__SortOrder) {
	defer qt.Recovering("QSqlQueryModel::sort")

	if ptr.Pointer() != nil {
		C.QSqlQueryModel_SortDefault(ptr.Pointer(), C.int(column), C.int(order))
	}
}

func (ptr *QSqlQueryModel) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QSqlQueryModel::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QSqlQueryModel) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QSqlQueryModel::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQSqlQueryModelTimerEvent
func callbackQSqlQueryModelTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSqlQueryModel::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSqlQueryModelFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSqlQueryModel) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSqlQueryModel::timerEvent")

	if ptr.Pointer() != nil {
		C.QSqlQueryModel_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSqlQueryModel) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSqlQueryModel::timerEvent")

	if ptr.Pointer() != nil {
		C.QSqlQueryModel_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSqlQueryModel) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QSqlQueryModel::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QSqlQueryModel) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QSqlQueryModel::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQSqlQueryModelChildEvent
func callbackQSqlQueryModelChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSqlQueryModel::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSqlQueryModelFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSqlQueryModel) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSqlQueryModel::childEvent")

	if ptr.Pointer() != nil {
		C.QSqlQueryModel_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSqlQueryModel) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSqlQueryModel::childEvent")

	if ptr.Pointer() != nil {
		C.QSqlQueryModel_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSqlQueryModel) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSqlQueryModel::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QSqlQueryModel) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QSqlQueryModel::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQSqlQueryModelCustomEvent
func callbackQSqlQueryModelCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSqlQueryModel::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSqlQueryModelFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSqlQueryModel) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSqlQueryModel::customEvent")

	if ptr.Pointer() != nil {
		C.QSqlQueryModel_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSqlQueryModel) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSqlQueryModel::customEvent")

	if ptr.Pointer() != nil {
		C.QSqlQueryModel_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QSqlRecord struct {
	ptr unsafe.Pointer
}

type QSqlRecord_ITF interface {
	QSqlRecord_PTR() *QSqlRecord
}

func (p *QSqlRecord) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSqlRecord) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSqlRecord(ptr QSqlRecord_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlRecord_PTR().Pointer()
	}
	return nil
}

func NewQSqlRecordFromPointer(ptr unsafe.Pointer) *QSqlRecord {
	var n = new(QSqlRecord)
	n.SetPointer(ptr)
	return n
}

func newQSqlRecordFromPointer(ptr unsafe.Pointer) *QSqlRecord {
	var n = NewQSqlRecordFromPointer(ptr)
	return n
}

func (ptr *QSqlRecord) QSqlRecord_PTR() *QSqlRecord {
	return ptr
}

func NewQSqlRecord() *QSqlRecord {
	defer qt.Recovering("QSqlRecord::QSqlRecord")

	return newQSqlRecordFromPointer(C.QSqlRecord_NewQSqlRecord())
}

func NewQSqlRecord2(other QSqlRecord_ITF) *QSqlRecord {
	defer qt.Recovering("QSqlRecord::QSqlRecord")

	return newQSqlRecordFromPointer(C.QSqlRecord_NewQSqlRecord2(PointerFromQSqlRecord(other)))
}

func (ptr *QSqlRecord) Append(field QSqlField_ITF) {
	defer qt.Recovering("QSqlRecord::append")

	if ptr.Pointer() != nil {
		C.QSqlRecord_Append(ptr.Pointer(), PointerFromQSqlField(field))
	}
}

func (ptr *QSqlRecord) Clear() {
	defer qt.Recovering("QSqlRecord::clear")

	if ptr.Pointer() != nil {
		C.QSqlRecord_Clear(ptr.Pointer())
	}
}

func (ptr *QSqlRecord) ClearValues() {
	defer qt.Recovering("QSqlRecord::clearValues")

	if ptr.Pointer() != nil {
		C.QSqlRecord_ClearValues(ptr.Pointer())
	}
}

func (ptr *QSqlRecord) Contains(name string) bool {
	defer qt.Recovering("QSqlRecord::contains")

	if ptr.Pointer() != nil {
		return C.QSqlRecord_Contains(ptr.Pointer(), C.CString(name)) != 0
	}
	return false
}

func (ptr *QSqlRecord) Count() int {
	defer qt.Recovering("QSqlRecord::count")

	if ptr.Pointer() != nil {
		return int(C.QSqlRecord_Count(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSqlRecord) FieldName(index int) string {
	defer qt.Recovering("QSqlRecord::fieldName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlRecord_FieldName(ptr.Pointer(), C.int(index)))
	}
	return ""
}

func (ptr *QSqlRecord) IndexOf(name string) int {
	defer qt.Recovering("QSqlRecord::indexOf")

	if ptr.Pointer() != nil {
		return int(C.QSqlRecord_IndexOf(ptr.Pointer(), C.CString(name)))
	}
	return 0
}

func (ptr *QSqlRecord) IsEmpty() bool {
	defer qt.Recovering("QSqlRecord::isEmpty")

	if ptr.Pointer() != nil {
		return C.QSqlRecord_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlRecord) IsGenerated(name string) bool {
	defer qt.Recovering("QSqlRecord::isGenerated")

	if ptr.Pointer() != nil {
		return C.QSqlRecord_IsGenerated(ptr.Pointer(), C.CString(name)) != 0
	}
	return false
}

func (ptr *QSqlRecord) IsGenerated2(index int) bool {
	defer qt.Recovering("QSqlRecord::isGenerated")

	if ptr.Pointer() != nil {
		return C.QSqlRecord_IsGenerated2(ptr.Pointer(), C.int(index)) != 0
	}
	return false
}

func (ptr *QSqlRecord) IsNull(name string) bool {
	defer qt.Recovering("QSqlRecord::isNull")

	if ptr.Pointer() != nil {
		return C.QSqlRecord_IsNull(ptr.Pointer(), C.CString(name)) != 0
	}
	return false
}

func (ptr *QSqlRecord) IsNull2(index int) bool {
	defer qt.Recovering("QSqlRecord::isNull")

	if ptr.Pointer() != nil {
		return C.QSqlRecord_IsNull2(ptr.Pointer(), C.int(index)) != 0
	}
	return false
}

func (ptr *QSqlRecord) SetGenerated(name string, generated bool) {
	defer qt.Recovering("QSqlRecord::setGenerated")

	if ptr.Pointer() != nil {
		C.QSqlRecord_SetGenerated(ptr.Pointer(), C.CString(name), C.int(qt.GoBoolToInt(generated)))
	}
}

func (ptr *QSqlRecord) SetGenerated2(index int, generated bool) {
	defer qt.Recovering("QSqlRecord::setGenerated")

	if ptr.Pointer() != nil {
		C.QSqlRecord_SetGenerated2(ptr.Pointer(), C.int(index), C.int(qt.GoBoolToInt(generated)))
	}
}

func (ptr *QSqlRecord) SetNull2(name string) {
	defer qt.Recovering("QSqlRecord::setNull")

	if ptr.Pointer() != nil {
		C.QSqlRecord_SetNull2(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QSqlRecord) SetNull(index int) {
	defer qt.Recovering("QSqlRecord::setNull")

	if ptr.Pointer() != nil {
		C.QSqlRecord_SetNull(ptr.Pointer(), C.int(index))
	}
}

func (ptr *QSqlRecord) SetValue2(name string, val core.QVariant_ITF) {
	defer qt.Recovering("QSqlRecord::setValue")

	if ptr.Pointer() != nil {
		C.QSqlRecord_SetValue2(ptr.Pointer(), C.CString(name), core.PointerFromQVariant(val))
	}
}

func (ptr *QSqlRecord) SetValue(index int, val core.QVariant_ITF) {
	defer qt.Recovering("QSqlRecord::setValue")

	if ptr.Pointer() != nil {
		C.QSqlRecord_SetValue(ptr.Pointer(), C.int(index), core.PointerFromQVariant(val))
	}
}

func (ptr *QSqlRecord) Value2(name string) *core.QVariant {
	defer qt.Recovering("QSqlRecord::value")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QSqlRecord_Value2(ptr.Pointer(), C.CString(name)))
	}
	return nil
}

func (ptr *QSqlRecord) Value(index int) *core.QVariant {
	defer qt.Recovering("QSqlRecord::value")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QSqlRecord_Value(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QSqlRecord) DestroyQSqlRecord() {
	defer qt.Recovering("QSqlRecord::~QSqlRecord")

	if ptr.Pointer() != nil {
		C.QSqlRecord_DestroyQSqlRecord(ptr.Pointer())
	}
}

type QSqlRelation struct {
	ptr unsafe.Pointer
}

type QSqlRelation_ITF interface {
	QSqlRelation_PTR() *QSqlRelation
}

func (p *QSqlRelation) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSqlRelation) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSqlRelation(ptr QSqlRelation_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlRelation_PTR().Pointer()
	}
	return nil
}

func NewQSqlRelationFromPointer(ptr unsafe.Pointer) *QSqlRelation {
	var n = new(QSqlRelation)
	n.SetPointer(ptr)
	return n
}

func newQSqlRelationFromPointer(ptr unsafe.Pointer) *QSqlRelation {
	var n = NewQSqlRelationFromPointer(ptr)
	return n
}

func (ptr *QSqlRelation) QSqlRelation_PTR() *QSqlRelation {
	return ptr
}

func NewQSqlRelation() *QSqlRelation {
	defer qt.Recovering("QSqlRelation::QSqlRelation")

	return newQSqlRelationFromPointer(C.QSqlRelation_NewQSqlRelation())
}

func NewQSqlRelation2(tableName string, indexColumn string, displayColumn string) *QSqlRelation {
	defer qt.Recovering("QSqlRelation::QSqlRelation")

	return newQSqlRelationFromPointer(C.QSqlRelation_NewQSqlRelation2(C.CString(tableName), C.CString(indexColumn), C.CString(displayColumn)))
}

func (ptr *QSqlRelation) DisplayColumn() string {
	defer qt.Recovering("QSqlRelation::displayColumn")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlRelation_DisplayColumn(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlRelation) IndexColumn() string {
	defer qt.Recovering("QSqlRelation::indexColumn")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlRelation_IndexColumn(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlRelation) IsValid() bool {
	defer qt.Recovering("QSqlRelation::isValid")

	if ptr.Pointer() != nil {
		return C.QSqlRelation_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlRelation) TableName() string {
	defer qt.Recovering("QSqlRelation::tableName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlRelation_TableName(ptr.Pointer()))
	}
	return ""
}

type QSqlRelationalDelegate struct {
	ptr unsafe.Pointer
}

type QSqlRelationalDelegate_ITF interface {
	QSqlRelationalDelegate_PTR() *QSqlRelationalDelegate
}

func (p *QSqlRelationalDelegate) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSqlRelationalDelegate) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSqlRelationalDelegate(ptr QSqlRelationalDelegate_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlRelationalDelegate_PTR().Pointer()
	}
	return nil
}

func NewQSqlRelationalDelegateFromPointer(ptr unsafe.Pointer) *QSqlRelationalDelegate {
	var n = new(QSqlRelationalDelegate)
	n.SetPointer(ptr)
	return n
}

func newQSqlRelationalDelegateFromPointer(ptr unsafe.Pointer) *QSqlRelationalDelegate {
	var n = NewQSqlRelationalDelegateFromPointer(ptr)
	return n
}

func (ptr *QSqlRelationalDelegate) QSqlRelationalDelegate_PTR() *QSqlRelationalDelegate {
	return ptr
}

type QSqlRelationalTableModel struct {
	QSqlTableModel
}

type QSqlRelationalTableModel_ITF interface {
	QSqlTableModel_ITF
	QSqlRelationalTableModel_PTR() *QSqlRelationalTableModel
}

func PointerFromQSqlRelationalTableModel(ptr QSqlRelationalTableModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlRelationalTableModel_PTR().Pointer()
	}
	return nil
}

func NewQSqlRelationalTableModelFromPointer(ptr unsafe.Pointer) *QSqlRelationalTableModel {
	var n = new(QSqlRelationalTableModel)
	n.SetPointer(ptr)
	return n
}

func newQSqlRelationalTableModelFromPointer(ptr unsafe.Pointer) *QSqlRelationalTableModel {
	var n = NewQSqlRelationalTableModelFromPointer(ptr)
	for len(n.ObjectName()) < len("QSqlRelationalTableModel_") {
		n.SetObjectName("QSqlRelationalTableModel_" + qt.Identifier())
	}
	return n
}

func (ptr *QSqlRelationalTableModel) QSqlRelationalTableModel_PTR() *QSqlRelationalTableModel {
	return ptr
}

//QSqlRelationalTableModel::JoinMode
type QSqlRelationalTableModel__JoinMode int64

const (
	QSqlRelationalTableModel__InnerJoin = QSqlRelationalTableModel__JoinMode(0)
	QSqlRelationalTableModel__LeftJoin  = QSqlRelationalTableModel__JoinMode(1)
)

func (ptr *QSqlRelationalTableModel) ConnectClear(f func()) {
	defer qt.Recovering("connect QSqlRelationalTableModel::clear")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "clear", f)
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectClear() {
	defer qt.Recovering("disconnect QSqlRelationalTableModel::clear")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "clear")
	}
}

//export callbackQSqlRelationalTableModelClear
func callbackQSqlRelationalTableModelClear(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSqlRelationalTableModel::clear")

	if signal := qt.GetSignal(C.GoString(ptrName), "clear"); signal != nil {
		signal.(func())()
	} else {
		NewQSqlRelationalTableModelFromPointer(ptr).ClearDefault()
	}
}

func (ptr *QSqlRelationalTableModel) Clear() {
	defer qt.Recovering("QSqlRelationalTableModel::clear")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_Clear(ptr.Pointer())
	}
}

func (ptr *QSqlRelationalTableModel) ClearDefault() {
	defer qt.Recovering("QSqlRelationalTableModel::clear")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_ClearDefault(ptr.Pointer())
	}
}

func (ptr *QSqlRelationalTableModel) Data(index core.QModelIndex_ITF, role int) *core.QVariant {
	defer qt.Recovering("QSqlRelationalTableModel::data")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QSqlRelationalTableModel_Data(ptr.Pointer(), core.PointerFromQModelIndex(index), C.int(role)))
	}
	return nil
}

func (ptr *QSqlRelationalTableModel) InsertRowIntoTable(values QSqlRecord_ITF) bool {
	defer qt.Recovering("QSqlRelationalTableModel::insertRowIntoTable")

	if ptr.Pointer() != nil {
		return C.QSqlRelationalTableModel_InsertRowIntoTable(ptr.Pointer(), PointerFromQSqlRecord(values)) != 0
	}
	return false
}

func (ptr *QSqlRelationalTableModel) OrderByClause() string {
	defer qt.Recovering("QSqlRelationalTableModel::orderByClause")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlRelationalTableModel_OrderByClause(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlRelationalTableModel) RelationModel(column int) *QSqlTableModel {
	defer qt.Recovering("QSqlRelationalTableModel::relationModel")

	if ptr.Pointer() != nil {
		return NewQSqlTableModelFromPointer(C.QSqlRelationalTableModel_RelationModel(ptr.Pointer(), C.int(column)))
	}
	return nil
}

func (ptr *QSqlRelationalTableModel) RemoveColumns(column int, count int, parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QSqlRelationalTableModel::removeColumns")

	if ptr.Pointer() != nil {
		return C.QSqlRelationalTableModel_RemoveColumns(ptr.Pointer(), C.int(column), C.int(count), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QSqlRelationalTableModel) ConnectRevertRow(f func(row int)) {
	defer qt.Recovering("connect QSqlRelationalTableModel::revertRow")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "revertRow", f)
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectRevertRow() {
	defer qt.Recovering("disconnect QSqlRelationalTableModel::revertRow")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "revertRow")
	}
}

//export callbackQSqlRelationalTableModelRevertRow
func callbackQSqlRelationalTableModelRevertRow(ptr unsafe.Pointer, ptrName *C.char, row C.int) bool {
	defer qt.Recovering("callback QSqlRelationalTableModel::revertRow")

	if signal := qt.GetSignal(C.GoString(ptrName), "revertRow"); signal != nil {
		signal.(func(int))(int(row))
		return true
	}
	return false

}

func (ptr *QSqlRelationalTableModel) RevertRow(row int) {
	defer qt.Recovering("QSqlRelationalTableModel::revertRow")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_RevertRow(ptr.Pointer(), C.int(row))
	}
}

func (ptr *QSqlRelationalTableModel) RevertRowDefault(row int) {
	defer qt.Recovering("QSqlRelationalTableModel::revertRow")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_RevertRowDefault(ptr.Pointer(), C.int(row))
	}
}

func (ptr *QSqlRelationalTableModel) Select() bool {
	defer qt.Recovering("QSqlRelationalTableModel::select")

	if ptr.Pointer() != nil {
		return C.QSqlRelationalTableModel_Select(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlRelationalTableModel) SelectStatement() string {
	defer qt.Recovering("QSqlRelationalTableModel::selectStatement")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlRelationalTableModel_SelectStatement(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlRelationalTableModel) SetData(index core.QModelIndex_ITF, value core.QVariant_ITF, role int) bool {
	defer qt.Recovering("QSqlRelationalTableModel::setData")

	if ptr.Pointer() != nil {
		return C.QSqlRelationalTableModel_SetData(ptr.Pointer(), core.PointerFromQModelIndex(index), core.PointerFromQVariant(value), C.int(role)) != 0
	}
	return false
}

func (ptr *QSqlRelationalTableModel) SetJoinMode(joinMode QSqlRelationalTableModel__JoinMode) {
	defer qt.Recovering("QSqlRelationalTableModel::setJoinMode")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_SetJoinMode(ptr.Pointer(), C.int(joinMode))
	}
}

func (ptr *QSqlRelationalTableModel) ConnectSetTable(f func(table string)) {
	defer qt.Recovering("connect QSqlRelationalTableModel::setTable")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setTable", f)
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectSetTable() {
	defer qt.Recovering("disconnect QSqlRelationalTableModel::setTable")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setTable")
	}
}

//export callbackQSqlRelationalTableModelSetTable
func callbackQSqlRelationalTableModelSetTable(ptr unsafe.Pointer, ptrName *C.char, table *C.char) {
	defer qt.Recovering("callback QSqlRelationalTableModel::setTable")

	if signal := qt.GetSignal(C.GoString(ptrName), "setTable"); signal != nil {
		signal.(func(string))(C.GoString(table))
	} else {
		NewQSqlRelationalTableModelFromPointer(ptr).SetTableDefault(C.GoString(table))
	}
}

func (ptr *QSqlRelationalTableModel) SetTable(table string) {
	defer qt.Recovering("QSqlRelationalTableModel::setTable")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_SetTable(ptr.Pointer(), C.CString(table))
	}
}

func (ptr *QSqlRelationalTableModel) SetTableDefault(table string) {
	defer qt.Recovering("QSqlRelationalTableModel::setTable")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_SetTableDefault(ptr.Pointer(), C.CString(table))
	}
}

func (ptr *QSqlRelationalTableModel) UpdateRowInTable(row int, values QSqlRecord_ITF) bool {
	defer qt.Recovering("QSqlRelationalTableModel::updateRowInTable")

	if ptr.Pointer() != nil {
		return C.QSqlRelationalTableModel_UpdateRowInTable(ptr.Pointer(), C.int(row), PointerFromQSqlRecord(values)) != 0
	}
	return false
}

func (ptr *QSqlRelationalTableModel) DestroyQSqlRelationalTableModel() {
	defer qt.Recovering("QSqlRelationalTableModel::~QSqlRelationalTableModel")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_DestroyQSqlRelationalTableModel(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSqlRelationalTableModel) ConnectRevert(f func()) {
	defer qt.Recovering("connect QSqlRelationalTableModel::revert")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "revert", f)
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectRevert() {
	defer qt.Recovering("disconnect QSqlRelationalTableModel::revert")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "revert")
	}
}

//export callbackQSqlRelationalTableModelRevert
func callbackQSqlRelationalTableModelRevert(ptr unsafe.Pointer, ptrName *C.char) bool {
	defer qt.Recovering("callback QSqlRelationalTableModel::revert")

	if signal := qt.GetSignal(C.GoString(ptrName), "revert"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QSqlRelationalTableModel) Revert() {
	defer qt.Recovering("QSqlRelationalTableModel::revert")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_Revert(ptr.Pointer())
	}
}

func (ptr *QSqlRelationalTableModel) RevertDefault() {
	defer qt.Recovering("QSqlRelationalTableModel::revert")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_RevertDefault(ptr.Pointer())
	}
}

func (ptr *QSqlRelationalTableModel) ConnectSetEditStrategy(f func(strategy QSqlTableModel__EditStrategy)) {
	defer qt.Recovering("connect QSqlRelationalTableModel::setEditStrategy")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setEditStrategy", f)
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectSetEditStrategy() {
	defer qt.Recovering("disconnect QSqlRelationalTableModel::setEditStrategy")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setEditStrategy")
	}
}

//export callbackQSqlRelationalTableModelSetEditStrategy
func callbackQSqlRelationalTableModelSetEditStrategy(ptr unsafe.Pointer, ptrName *C.char, strategy C.int) {
	defer qt.Recovering("callback QSqlRelationalTableModel::setEditStrategy")

	if signal := qt.GetSignal(C.GoString(ptrName), "setEditStrategy"); signal != nil {
		signal.(func(QSqlTableModel__EditStrategy))(QSqlTableModel__EditStrategy(strategy))
	} else {
		NewQSqlRelationalTableModelFromPointer(ptr).SetEditStrategyDefault(QSqlTableModel__EditStrategy(strategy))
	}
}

func (ptr *QSqlRelationalTableModel) SetEditStrategy(strategy QSqlTableModel__EditStrategy) {
	defer qt.Recovering("QSqlRelationalTableModel::setEditStrategy")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_SetEditStrategy(ptr.Pointer(), C.int(strategy))
	}
}

func (ptr *QSqlRelationalTableModel) SetEditStrategyDefault(strategy QSqlTableModel__EditStrategy) {
	defer qt.Recovering("QSqlRelationalTableModel::setEditStrategy")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_SetEditStrategyDefault(ptr.Pointer(), C.int(strategy))
	}
}

func (ptr *QSqlRelationalTableModel) ConnectSetFilter(f func(filter string)) {
	defer qt.Recovering("connect QSqlRelationalTableModel::setFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setFilter", f)
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectSetFilter() {
	defer qt.Recovering("disconnect QSqlRelationalTableModel::setFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setFilter")
	}
}

//export callbackQSqlRelationalTableModelSetFilter
func callbackQSqlRelationalTableModelSetFilter(ptr unsafe.Pointer, ptrName *C.char, filter *C.char) {
	defer qt.Recovering("callback QSqlRelationalTableModel::setFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "setFilter"); signal != nil {
		signal.(func(string))(C.GoString(filter))
	} else {
		NewQSqlRelationalTableModelFromPointer(ptr).SetFilterDefault(C.GoString(filter))
	}
}

func (ptr *QSqlRelationalTableModel) SetFilter(filter string) {
	defer qt.Recovering("QSqlRelationalTableModel::setFilter")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_SetFilter(ptr.Pointer(), C.CString(filter))
	}
}

func (ptr *QSqlRelationalTableModel) SetFilterDefault(filter string) {
	defer qt.Recovering("QSqlRelationalTableModel::setFilter")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_SetFilterDefault(ptr.Pointer(), C.CString(filter))
	}
}

func (ptr *QSqlRelationalTableModel) ConnectSetSort(f func(column int, order core.Qt__SortOrder)) {
	defer qt.Recovering("connect QSqlRelationalTableModel::setSort")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setSort", f)
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectSetSort() {
	defer qt.Recovering("disconnect QSqlRelationalTableModel::setSort")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setSort")
	}
}

//export callbackQSqlRelationalTableModelSetSort
func callbackQSqlRelationalTableModelSetSort(ptr unsafe.Pointer, ptrName *C.char, column C.int, order C.int) {
	defer qt.Recovering("callback QSqlRelationalTableModel::setSort")

	if signal := qt.GetSignal(C.GoString(ptrName), "setSort"); signal != nil {
		signal.(func(int, core.Qt__SortOrder))(int(column), core.Qt__SortOrder(order))
	} else {
		NewQSqlRelationalTableModelFromPointer(ptr).SetSortDefault(int(column), core.Qt__SortOrder(order))
	}
}

func (ptr *QSqlRelationalTableModel) SetSort(column int, order core.Qt__SortOrder) {
	defer qt.Recovering("QSqlRelationalTableModel::setSort")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_SetSort(ptr.Pointer(), C.int(column), C.int(order))
	}
}

func (ptr *QSqlRelationalTableModel) SetSortDefault(column int, order core.Qt__SortOrder) {
	defer qt.Recovering("QSqlRelationalTableModel::setSort")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_SetSortDefault(ptr.Pointer(), C.int(column), C.int(order))
	}
}

func (ptr *QSqlRelationalTableModel) ConnectSort(f func(column int, order core.Qt__SortOrder)) {
	defer qt.Recovering("connect QSqlRelationalTableModel::sort")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "sort", f)
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectSort() {
	defer qt.Recovering("disconnect QSqlRelationalTableModel::sort")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "sort")
	}
}

//export callbackQSqlRelationalTableModelSort
func callbackQSqlRelationalTableModelSort(ptr unsafe.Pointer, ptrName *C.char, column C.int, order C.int) {
	defer qt.Recovering("callback QSqlRelationalTableModel::sort")

	if signal := qt.GetSignal(C.GoString(ptrName), "sort"); signal != nil {
		signal.(func(int, core.Qt__SortOrder))(int(column), core.Qt__SortOrder(order))
	} else {
		NewQSqlRelationalTableModelFromPointer(ptr).SortDefault(int(column), core.Qt__SortOrder(order))
	}
}

func (ptr *QSqlRelationalTableModel) Sort(column int, order core.Qt__SortOrder) {
	defer qt.Recovering("QSqlRelationalTableModel::sort")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_Sort(ptr.Pointer(), C.int(column), C.int(order))
	}
}

func (ptr *QSqlRelationalTableModel) SortDefault(column int, order core.Qt__SortOrder) {
	defer qt.Recovering("QSqlRelationalTableModel::sort")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_SortDefault(ptr.Pointer(), C.int(column), C.int(order))
	}
}

func (ptr *QSqlRelationalTableModel) ConnectFetchMore(f func(parent *core.QModelIndex)) {
	defer qt.Recovering("connect QSqlRelationalTableModel::fetchMore")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "fetchMore", f)
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectFetchMore() {
	defer qt.Recovering("disconnect QSqlRelationalTableModel::fetchMore")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "fetchMore")
	}
}

//export callbackQSqlRelationalTableModelFetchMore
func callbackQSqlRelationalTableModelFetchMore(ptr unsafe.Pointer, ptrName *C.char, parent unsafe.Pointer) {
	defer qt.Recovering("callback QSqlRelationalTableModel::fetchMore")

	if signal := qt.GetSignal(C.GoString(ptrName), "fetchMore"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(parent))
	} else {
		NewQSqlRelationalTableModelFromPointer(ptr).FetchMoreDefault(core.NewQModelIndexFromPointer(parent))
	}
}

func (ptr *QSqlRelationalTableModel) FetchMore(parent core.QModelIndex_ITF) {
	defer qt.Recovering("QSqlRelationalTableModel::fetchMore")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_FetchMore(ptr.Pointer(), core.PointerFromQModelIndex(parent))
	}
}

func (ptr *QSqlRelationalTableModel) FetchMoreDefault(parent core.QModelIndex_ITF) {
	defer qt.Recovering("QSqlRelationalTableModel::fetchMore")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_FetchMoreDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent))
	}
}

func (ptr *QSqlRelationalTableModel) ConnectQueryChange(f func()) {
	defer qt.Recovering("connect QSqlRelationalTableModel::queryChange")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "queryChange", f)
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectQueryChange() {
	defer qt.Recovering("disconnect QSqlRelationalTableModel::queryChange")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "queryChange")
	}
}

//export callbackQSqlRelationalTableModelQueryChange
func callbackQSqlRelationalTableModelQueryChange(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSqlRelationalTableModel::queryChange")

	if signal := qt.GetSignal(C.GoString(ptrName), "queryChange"); signal != nil {
		signal.(func())()
	} else {
		NewQSqlRelationalTableModelFromPointer(ptr).QueryChangeDefault()
	}
}

func (ptr *QSqlRelationalTableModel) QueryChange() {
	defer qt.Recovering("QSqlRelationalTableModel::queryChange")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_QueryChange(ptr.Pointer())
	}
}

func (ptr *QSqlRelationalTableModel) QueryChangeDefault() {
	defer qt.Recovering("QSqlRelationalTableModel::queryChange")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_QueryChangeDefault(ptr.Pointer())
	}
}

func (ptr *QSqlRelationalTableModel) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QSqlRelationalTableModel::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QSqlRelationalTableModel::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQSqlRelationalTableModelTimerEvent
func callbackQSqlRelationalTableModelTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSqlRelationalTableModel::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSqlRelationalTableModelFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSqlRelationalTableModel) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSqlRelationalTableModel::timerEvent")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSqlRelationalTableModel) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSqlRelationalTableModel::timerEvent")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSqlRelationalTableModel) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QSqlRelationalTableModel::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QSqlRelationalTableModel::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQSqlRelationalTableModelChildEvent
func callbackQSqlRelationalTableModelChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSqlRelationalTableModel::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSqlRelationalTableModelFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSqlRelationalTableModel) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSqlRelationalTableModel::childEvent")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSqlRelationalTableModel) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSqlRelationalTableModel::childEvent")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSqlRelationalTableModel) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSqlRelationalTableModel::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QSqlRelationalTableModel::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQSqlRelationalTableModelCustomEvent
func callbackQSqlRelationalTableModelCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSqlRelationalTableModel::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSqlRelationalTableModelFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSqlRelationalTableModel) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSqlRelationalTableModel::customEvent")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSqlRelationalTableModel) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSqlRelationalTableModel::customEvent")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QSqlResult struct {
	ptr unsafe.Pointer
}

type QSqlResult_ITF interface {
	QSqlResult_PTR() *QSqlResult
}

func (p *QSqlResult) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSqlResult) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSqlResult(ptr QSqlResult_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlResult_PTR().Pointer()
	}
	return nil
}

func NewQSqlResultFromPointer(ptr unsafe.Pointer) *QSqlResult {
	var n = new(QSqlResult)
	n.SetPointer(ptr)
	return n
}

func newQSqlResultFromPointer(ptr unsafe.Pointer) *QSqlResult {
	var n = NewQSqlResultFromPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QSqlResult_") {
		n.SetObjectNameAbs("QSqlResult_" + qt.Identifier())
	}
	return n
}

func (ptr *QSqlResult) QSqlResult_PTR() *QSqlResult {
	return ptr
}

//QSqlResult::BindingSyntax
type QSqlResult__BindingSyntax int64

const (
	QSqlResult__PositionalBinding = QSqlResult__BindingSyntax(0)
	QSqlResult__NamedBinding      = QSqlResult__BindingSyntax(1)
)

func (ptr *QSqlResult) Exec() bool {
	defer qt.Recovering("QSqlResult::exec")

	if ptr.Pointer() != nil {
		return C.QSqlResult_Exec(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlResult) FetchNext() bool {
	defer qt.Recovering("QSqlResult::fetchNext")

	if ptr.Pointer() != nil {
		return C.QSqlResult_FetchNext(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlResult) FetchPrevious() bool {
	defer qt.Recovering("QSqlResult::fetchPrevious")

	if ptr.Pointer() != nil {
		return C.QSqlResult_FetchPrevious(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlResult) Handle() *core.QVariant {
	defer qt.Recovering("QSqlResult::handle")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QSqlResult_Handle(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSqlResult) LastInsertId() *core.QVariant {
	defer qt.Recovering("QSqlResult::lastInsertId")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QSqlResult_LastInsertId(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSqlResult) Prepare(query string) bool {
	defer qt.Recovering("QSqlResult::prepare")

	if ptr.Pointer() != nil {
		return C.QSqlResult_Prepare(ptr.Pointer(), C.CString(query)) != 0
	}
	return false
}

func (ptr *QSqlResult) SavePrepare(query string) bool {
	defer qt.Recovering("QSqlResult::savePrepare")

	if ptr.Pointer() != nil {
		return C.QSqlResult_SavePrepare(ptr.Pointer(), C.CString(query)) != 0
	}
	return false
}

func (ptr *QSqlResult) ConnectSetActive(f func(active bool)) {
	defer qt.Recovering("connect QSqlResult::setActive")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "setActive", f)
	}
}

func (ptr *QSqlResult) DisconnectSetActive() {
	defer qt.Recovering("disconnect QSqlResult::setActive")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "setActive")
	}
}

//export callbackQSqlResultSetActive
func callbackQSqlResultSetActive(ptr unsafe.Pointer, ptrName *C.char, active C.int) {
	defer qt.Recovering("callback QSqlResult::setActive")

	if signal := qt.GetSignal(C.GoString(ptrName), "setActive"); signal != nil {
		signal.(func(bool))(int(active) != 0)
	} else {
		NewQSqlResultFromPointer(ptr).SetActiveDefault(int(active) != 0)
	}
}

func (ptr *QSqlResult) SetActive(active bool) {
	defer qt.Recovering("QSqlResult::setActive")

	if ptr.Pointer() != nil {
		C.QSqlResult_SetActive(ptr.Pointer(), C.int(qt.GoBoolToInt(active)))
	}
}

func (ptr *QSqlResult) SetActiveDefault(active bool) {
	defer qt.Recovering("QSqlResult::setActive")

	if ptr.Pointer() != nil {
		C.QSqlResult_SetActiveDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(active)))
	}
}

func (ptr *QSqlResult) ConnectSetAt(f func(index int)) {
	defer qt.Recovering("connect QSqlResult::setAt")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "setAt", f)
	}
}

func (ptr *QSqlResult) DisconnectSetAt() {
	defer qt.Recovering("disconnect QSqlResult::setAt")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "setAt")
	}
}

//export callbackQSqlResultSetAt
func callbackQSqlResultSetAt(ptr unsafe.Pointer, ptrName *C.char, index C.int) {
	defer qt.Recovering("callback QSqlResult::setAt")

	if signal := qt.GetSignal(C.GoString(ptrName), "setAt"); signal != nil {
		signal.(func(int))(int(index))
	} else {
		NewQSqlResultFromPointer(ptr).SetAtDefault(int(index))
	}
}

func (ptr *QSqlResult) SetAt(index int) {
	defer qt.Recovering("QSqlResult::setAt")

	if ptr.Pointer() != nil {
		C.QSqlResult_SetAt(ptr.Pointer(), C.int(index))
	}
}

func (ptr *QSqlResult) SetAtDefault(index int) {
	defer qt.Recovering("QSqlResult::setAt")

	if ptr.Pointer() != nil {
		C.QSqlResult_SetAtDefault(ptr.Pointer(), C.int(index))
	}
}

func (ptr *QSqlResult) ConnectSetForwardOnly(f func(forward bool)) {
	defer qt.Recovering("connect QSqlResult::setForwardOnly")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "setForwardOnly", f)
	}
}

func (ptr *QSqlResult) DisconnectSetForwardOnly() {
	defer qt.Recovering("disconnect QSqlResult::setForwardOnly")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "setForwardOnly")
	}
}

//export callbackQSqlResultSetForwardOnly
func callbackQSqlResultSetForwardOnly(ptr unsafe.Pointer, ptrName *C.char, forward C.int) {
	defer qt.Recovering("callback QSqlResult::setForwardOnly")

	if signal := qt.GetSignal(C.GoString(ptrName), "setForwardOnly"); signal != nil {
		signal.(func(bool))(int(forward) != 0)
	} else {
		NewQSqlResultFromPointer(ptr).SetForwardOnlyDefault(int(forward) != 0)
	}
}

func (ptr *QSqlResult) SetForwardOnly(forward bool) {
	defer qt.Recovering("QSqlResult::setForwardOnly")

	if ptr.Pointer() != nil {
		C.QSqlResult_SetForwardOnly(ptr.Pointer(), C.int(qt.GoBoolToInt(forward)))
	}
}

func (ptr *QSqlResult) SetForwardOnlyDefault(forward bool) {
	defer qt.Recovering("QSqlResult::setForwardOnly")

	if ptr.Pointer() != nil {
		C.QSqlResult_SetForwardOnlyDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(forward)))
	}
}

func (ptr *QSqlResult) ConnectSetQuery(f func(query string)) {
	defer qt.Recovering("connect QSqlResult::setQuery")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "setQuery", f)
	}
}

func (ptr *QSqlResult) DisconnectSetQuery() {
	defer qt.Recovering("disconnect QSqlResult::setQuery")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "setQuery")
	}
}

//export callbackQSqlResultSetQuery
func callbackQSqlResultSetQuery(ptr unsafe.Pointer, ptrName *C.char, query *C.char) {
	defer qt.Recovering("callback QSqlResult::setQuery")

	if signal := qt.GetSignal(C.GoString(ptrName), "setQuery"); signal != nil {
		signal.(func(string))(C.GoString(query))
	} else {
		NewQSqlResultFromPointer(ptr).SetQueryDefault(C.GoString(query))
	}
}

func (ptr *QSqlResult) SetQuery(query string) {
	defer qt.Recovering("QSqlResult::setQuery")

	if ptr.Pointer() != nil {
		C.QSqlResult_SetQuery(ptr.Pointer(), C.CString(query))
	}
}

func (ptr *QSqlResult) SetQueryDefault(query string) {
	defer qt.Recovering("QSqlResult::setQuery")

	if ptr.Pointer() != nil {
		C.QSqlResult_SetQueryDefault(ptr.Pointer(), C.CString(query))
	}
}

func (ptr *QSqlResult) DestroyQSqlResult() {
	defer qt.Recovering("QSqlResult::~QSqlResult")

	if ptr.Pointer() != nil {
		C.QSqlResult_DestroyQSqlResult(ptr.Pointer())
	}
}

func (ptr *QSqlResult) ObjectNameAbs() string {
	defer qt.Recovering("QSqlResult::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlResult_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlResult) SetObjectNameAbs(name string) {
	defer qt.Recovering("QSqlResult::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QSqlResult_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}

type QSqlTableModel struct {
	QSqlQueryModel
}

type QSqlTableModel_ITF interface {
	QSqlQueryModel_ITF
	QSqlTableModel_PTR() *QSqlTableModel
}

func PointerFromQSqlTableModel(ptr QSqlTableModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlTableModel_PTR().Pointer()
	}
	return nil
}

func NewQSqlTableModelFromPointer(ptr unsafe.Pointer) *QSqlTableModel {
	var n = new(QSqlTableModel)
	n.SetPointer(ptr)
	return n
}

func newQSqlTableModelFromPointer(ptr unsafe.Pointer) *QSqlTableModel {
	var n = NewQSqlTableModelFromPointer(ptr)
	for len(n.ObjectName()) < len("QSqlTableModel_") {
		n.SetObjectName("QSqlTableModel_" + qt.Identifier())
	}
	return n
}

func (ptr *QSqlTableModel) QSqlTableModel_PTR() *QSqlTableModel {
	return ptr
}

//QSqlTableModel::EditStrategy
type QSqlTableModel__EditStrategy int64

const (
	QSqlTableModel__OnFieldChange  = QSqlTableModel__EditStrategy(0)
	QSqlTableModel__OnRowChange    = QSqlTableModel__EditStrategy(1)
	QSqlTableModel__OnManualSubmit = QSqlTableModel__EditStrategy(2)
)

func (ptr *QSqlTableModel) ConnectBeforeDelete(f func(row int)) {
	defer qt.Recovering("connect QSqlTableModel::beforeDelete")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_ConnectBeforeDelete(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "beforeDelete", f)
	}
}

func (ptr *QSqlTableModel) DisconnectBeforeDelete() {
	defer qt.Recovering("disconnect QSqlTableModel::beforeDelete")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_DisconnectBeforeDelete(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "beforeDelete")
	}
}

//export callbackQSqlTableModelBeforeDelete
func callbackQSqlTableModelBeforeDelete(ptr unsafe.Pointer, ptrName *C.char, row C.int) {
	defer qt.Recovering("callback QSqlTableModel::beforeDelete")

	if signal := qt.GetSignal(C.GoString(ptrName), "beforeDelete"); signal != nil {
		signal.(func(int))(int(row))
	}

}

func (ptr *QSqlTableModel) BeforeDelete(row int) {
	defer qt.Recovering("QSqlTableModel::beforeDelete")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_BeforeDelete(ptr.Pointer(), C.int(row))
	}
}

func (ptr *QSqlTableModel) ConnectClear(f func()) {
	defer qt.Recovering("connect QSqlTableModel::clear")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "clear", f)
	}
}

func (ptr *QSqlTableModel) DisconnectClear() {
	defer qt.Recovering("disconnect QSqlTableModel::clear")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "clear")
	}
}

//export callbackQSqlTableModelClear
func callbackQSqlTableModelClear(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSqlTableModel::clear")

	if signal := qt.GetSignal(C.GoString(ptrName), "clear"); signal != nil {
		signal.(func())()
	} else {
		NewQSqlTableModelFromPointer(ptr).ClearDefault()
	}
}

func (ptr *QSqlTableModel) Clear() {
	defer qt.Recovering("QSqlTableModel::clear")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_Clear(ptr.Pointer())
	}
}

func (ptr *QSqlTableModel) ClearDefault() {
	defer qt.Recovering("QSqlTableModel::clear")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_ClearDefault(ptr.Pointer())
	}
}

func (ptr *QSqlTableModel) Data(index core.QModelIndex_ITF, role int) *core.QVariant {
	defer qt.Recovering("QSqlTableModel::data")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QSqlTableModel_Data(ptr.Pointer(), core.PointerFromQModelIndex(index), C.int(role)))
	}
	return nil
}

func (ptr *QSqlTableModel) DeleteRowFromTable(row int) bool {
	defer qt.Recovering("QSqlTableModel::deleteRowFromTable")

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_DeleteRowFromTable(ptr.Pointer(), C.int(row)) != 0
	}
	return false
}

func (ptr *QSqlTableModel) EditStrategy() QSqlTableModel__EditStrategy {
	defer qt.Recovering("QSqlTableModel::editStrategy")

	if ptr.Pointer() != nil {
		return QSqlTableModel__EditStrategy(C.QSqlTableModel_EditStrategy(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSqlTableModel) FieldIndex(fieldName string) int {
	defer qt.Recovering("QSqlTableModel::fieldIndex")

	if ptr.Pointer() != nil {
		return int(C.QSqlTableModel_FieldIndex(ptr.Pointer(), C.CString(fieldName)))
	}
	return 0
}

func (ptr *QSqlTableModel) Filter() string {
	defer qt.Recovering("QSqlTableModel::filter")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlTableModel_Filter(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlTableModel) Flags(index core.QModelIndex_ITF) core.Qt__ItemFlag {
	defer qt.Recovering("QSqlTableModel::flags")

	if ptr.Pointer() != nil {
		return core.Qt__ItemFlag(C.QSqlTableModel_Flags(ptr.Pointer(), core.PointerFromQModelIndex(index)))
	}
	return 0
}

func (ptr *QSqlTableModel) HeaderData(section int, orientation core.Qt__Orientation, role int) *core.QVariant {
	defer qt.Recovering("QSqlTableModel::headerData")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QSqlTableModel_HeaderData(ptr.Pointer(), C.int(section), C.int(orientation), C.int(role)))
	}
	return nil
}

func (ptr *QSqlTableModel) IndexInQuery(item core.QModelIndex_ITF) *core.QModelIndex {
	defer qt.Recovering("QSqlTableModel::indexInQuery")

	if ptr.Pointer() != nil {
		return core.NewQModelIndexFromPointer(C.QSqlTableModel_IndexInQuery(ptr.Pointer(), core.PointerFromQModelIndex(item)))
	}
	return nil
}

func (ptr *QSqlTableModel) InsertRecord(row int, record QSqlRecord_ITF) bool {
	defer qt.Recovering("QSqlTableModel::insertRecord")

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_InsertRecord(ptr.Pointer(), C.int(row), PointerFromQSqlRecord(record)) != 0
	}
	return false
}

func (ptr *QSqlTableModel) InsertRowIntoTable(values QSqlRecord_ITF) bool {
	defer qt.Recovering("QSqlTableModel::insertRowIntoTable")

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_InsertRowIntoTable(ptr.Pointer(), PointerFromQSqlRecord(values)) != 0
	}
	return false
}

func (ptr *QSqlTableModel) InsertRows(row int, count int, parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QSqlTableModel::insertRows")

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_InsertRows(ptr.Pointer(), C.int(row), C.int(count), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QSqlTableModel) IsDirty2() bool {
	defer qt.Recovering("QSqlTableModel::isDirty")

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_IsDirty2(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlTableModel) IsDirty(index core.QModelIndex_ITF) bool {
	defer qt.Recovering("QSqlTableModel::isDirty")

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_IsDirty(ptr.Pointer(), core.PointerFromQModelIndex(index)) != 0
	}
	return false
}

func (ptr *QSqlTableModel) OrderByClause() string {
	defer qt.Recovering("QSqlTableModel::orderByClause")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlTableModel_OrderByClause(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlTableModel) RemoveColumns(column int, count int, parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QSqlTableModel::removeColumns")

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_RemoveColumns(ptr.Pointer(), C.int(column), C.int(count), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QSqlTableModel) RemoveRows(row int, count int, parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QSqlTableModel::removeRows")

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_RemoveRows(ptr.Pointer(), C.int(row), C.int(count), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QSqlTableModel) ConnectRevert(f func()) {
	defer qt.Recovering("connect QSqlTableModel::revert")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "revert", f)
	}
}

func (ptr *QSqlTableModel) DisconnectRevert() {
	defer qt.Recovering("disconnect QSqlTableModel::revert")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "revert")
	}
}

//export callbackQSqlTableModelRevert
func callbackQSqlTableModelRevert(ptr unsafe.Pointer, ptrName *C.char) bool {
	defer qt.Recovering("callback QSqlTableModel::revert")

	if signal := qt.GetSignal(C.GoString(ptrName), "revert"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QSqlTableModel) Revert() {
	defer qt.Recovering("QSqlTableModel::revert")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_Revert(ptr.Pointer())
	}
}

func (ptr *QSqlTableModel) RevertDefault() {
	defer qt.Recovering("QSqlTableModel::revert")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_RevertDefault(ptr.Pointer())
	}
}

func (ptr *QSqlTableModel) RevertAll() {
	defer qt.Recovering("QSqlTableModel::revertAll")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_RevertAll(ptr.Pointer())
	}
}

func (ptr *QSqlTableModel) ConnectRevertRow(f func(row int)) {
	defer qt.Recovering("connect QSqlTableModel::revertRow")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "revertRow", f)
	}
}

func (ptr *QSqlTableModel) DisconnectRevertRow() {
	defer qt.Recovering("disconnect QSqlTableModel::revertRow")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "revertRow")
	}
}

//export callbackQSqlTableModelRevertRow
func callbackQSqlTableModelRevertRow(ptr unsafe.Pointer, ptrName *C.char, row C.int) {
	defer qt.Recovering("callback QSqlTableModel::revertRow")

	if signal := qt.GetSignal(C.GoString(ptrName), "revertRow"); signal != nil {
		signal.(func(int))(int(row))
	} else {
		NewQSqlTableModelFromPointer(ptr).RevertRowDefault(int(row))
	}
}

func (ptr *QSqlTableModel) RevertRow(row int) {
	defer qt.Recovering("QSqlTableModel::revertRow")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_RevertRow(ptr.Pointer(), C.int(row))
	}
}

func (ptr *QSqlTableModel) RevertRowDefault(row int) {
	defer qt.Recovering("QSqlTableModel::revertRow")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_RevertRowDefault(ptr.Pointer(), C.int(row))
	}
}

func (ptr *QSqlTableModel) RowCount(parent core.QModelIndex_ITF) int {
	defer qt.Recovering("QSqlTableModel::rowCount")

	if ptr.Pointer() != nil {
		return int(C.QSqlTableModel_RowCount(ptr.Pointer(), core.PointerFromQModelIndex(parent)))
	}
	return 0
}

func (ptr *QSqlTableModel) Select() bool {
	defer qt.Recovering("QSqlTableModel::select")

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_Select(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlTableModel) SelectRow(row int) bool {
	defer qt.Recovering("QSqlTableModel::selectRow")

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_SelectRow(ptr.Pointer(), C.int(row)) != 0
	}
	return false
}

func (ptr *QSqlTableModel) SelectStatement() string {
	defer qt.Recovering("QSqlTableModel::selectStatement")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlTableModel_SelectStatement(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlTableModel) SetData(index core.QModelIndex_ITF, value core.QVariant_ITF, role int) bool {
	defer qt.Recovering("QSqlTableModel::setData")

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_SetData(ptr.Pointer(), core.PointerFromQModelIndex(index), core.PointerFromQVariant(value), C.int(role)) != 0
	}
	return false
}

func (ptr *QSqlTableModel) ConnectSetEditStrategy(f func(strategy QSqlTableModel__EditStrategy)) {
	defer qt.Recovering("connect QSqlTableModel::setEditStrategy")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setEditStrategy", f)
	}
}

func (ptr *QSqlTableModel) DisconnectSetEditStrategy() {
	defer qt.Recovering("disconnect QSqlTableModel::setEditStrategy")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setEditStrategy")
	}
}

//export callbackQSqlTableModelSetEditStrategy
func callbackQSqlTableModelSetEditStrategy(ptr unsafe.Pointer, ptrName *C.char, strategy C.int) {
	defer qt.Recovering("callback QSqlTableModel::setEditStrategy")

	if signal := qt.GetSignal(C.GoString(ptrName), "setEditStrategy"); signal != nil {
		signal.(func(QSqlTableModel__EditStrategy))(QSqlTableModel__EditStrategy(strategy))
	} else {
		NewQSqlTableModelFromPointer(ptr).SetEditStrategyDefault(QSqlTableModel__EditStrategy(strategy))
	}
}

func (ptr *QSqlTableModel) SetEditStrategy(strategy QSqlTableModel__EditStrategy) {
	defer qt.Recovering("QSqlTableModel::setEditStrategy")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_SetEditStrategy(ptr.Pointer(), C.int(strategy))
	}
}

func (ptr *QSqlTableModel) SetEditStrategyDefault(strategy QSqlTableModel__EditStrategy) {
	defer qt.Recovering("QSqlTableModel::setEditStrategy")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_SetEditStrategyDefault(ptr.Pointer(), C.int(strategy))
	}
}

func (ptr *QSqlTableModel) ConnectSetFilter(f func(filter string)) {
	defer qt.Recovering("connect QSqlTableModel::setFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setFilter", f)
	}
}

func (ptr *QSqlTableModel) DisconnectSetFilter() {
	defer qt.Recovering("disconnect QSqlTableModel::setFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setFilter")
	}
}

//export callbackQSqlTableModelSetFilter
func callbackQSqlTableModelSetFilter(ptr unsafe.Pointer, ptrName *C.char, filter *C.char) {
	defer qt.Recovering("callback QSqlTableModel::setFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "setFilter"); signal != nil {
		signal.(func(string))(C.GoString(filter))
	} else {
		NewQSqlTableModelFromPointer(ptr).SetFilterDefault(C.GoString(filter))
	}
}

func (ptr *QSqlTableModel) SetFilter(filter string) {
	defer qt.Recovering("QSqlTableModel::setFilter")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_SetFilter(ptr.Pointer(), C.CString(filter))
	}
}

func (ptr *QSqlTableModel) SetFilterDefault(filter string) {
	defer qt.Recovering("QSqlTableModel::setFilter")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_SetFilterDefault(ptr.Pointer(), C.CString(filter))
	}
}

func (ptr *QSqlTableModel) SetRecord(row int, values QSqlRecord_ITF) bool {
	defer qt.Recovering("QSqlTableModel::setRecord")

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_SetRecord(ptr.Pointer(), C.int(row), PointerFromQSqlRecord(values)) != 0
	}
	return false
}

func (ptr *QSqlTableModel) ConnectSetSort(f func(column int, order core.Qt__SortOrder)) {
	defer qt.Recovering("connect QSqlTableModel::setSort")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setSort", f)
	}
}

func (ptr *QSqlTableModel) DisconnectSetSort() {
	defer qt.Recovering("disconnect QSqlTableModel::setSort")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setSort")
	}
}

//export callbackQSqlTableModelSetSort
func callbackQSqlTableModelSetSort(ptr unsafe.Pointer, ptrName *C.char, column C.int, order C.int) {
	defer qt.Recovering("callback QSqlTableModel::setSort")

	if signal := qt.GetSignal(C.GoString(ptrName), "setSort"); signal != nil {
		signal.(func(int, core.Qt__SortOrder))(int(column), core.Qt__SortOrder(order))
	} else {
		NewQSqlTableModelFromPointer(ptr).SetSortDefault(int(column), core.Qt__SortOrder(order))
	}
}

func (ptr *QSqlTableModel) SetSort(column int, order core.Qt__SortOrder) {
	defer qt.Recovering("QSqlTableModel::setSort")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_SetSort(ptr.Pointer(), C.int(column), C.int(order))
	}
}

func (ptr *QSqlTableModel) SetSortDefault(column int, order core.Qt__SortOrder) {
	defer qt.Recovering("QSqlTableModel::setSort")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_SetSortDefault(ptr.Pointer(), C.int(column), C.int(order))
	}
}

func (ptr *QSqlTableModel) ConnectSetTable(f func(tableName string)) {
	defer qt.Recovering("connect QSqlTableModel::setTable")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setTable", f)
	}
}

func (ptr *QSqlTableModel) DisconnectSetTable() {
	defer qt.Recovering("disconnect QSqlTableModel::setTable")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setTable")
	}
}

//export callbackQSqlTableModelSetTable
func callbackQSqlTableModelSetTable(ptr unsafe.Pointer, ptrName *C.char, tableName *C.char) {
	defer qt.Recovering("callback QSqlTableModel::setTable")

	if signal := qt.GetSignal(C.GoString(ptrName), "setTable"); signal != nil {
		signal.(func(string))(C.GoString(tableName))
	} else {
		NewQSqlTableModelFromPointer(ptr).SetTableDefault(C.GoString(tableName))
	}
}

func (ptr *QSqlTableModel) SetTable(tableName string) {
	defer qt.Recovering("QSqlTableModel::setTable")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_SetTable(ptr.Pointer(), C.CString(tableName))
	}
}

func (ptr *QSqlTableModel) SetTableDefault(tableName string) {
	defer qt.Recovering("QSqlTableModel::setTable")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_SetTableDefault(ptr.Pointer(), C.CString(tableName))
	}
}

func (ptr *QSqlTableModel) ConnectSort(f func(column int, order core.Qt__SortOrder)) {
	defer qt.Recovering("connect QSqlTableModel::sort")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "sort", f)
	}
}

func (ptr *QSqlTableModel) DisconnectSort() {
	defer qt.Recovering("disconnect QSqlTableModel::sort")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "sort")
	}
}

//export callbackQSqlTableModelSort
func callbackQSqlTableModelSort(ptr unsafe.Pointer, ptrName *C.char, column C.int, order C.int) {
	defer qt.Recovering("callback QSqlTableModel::sort")

	if signal := qt.GetSignal(C.GoString(ptrName), "sort"); signal != nil {
		signal.(func(int, core.Qt__SortOrder))(int(column), core.Qt__SortOrder(order))
	} else {
		NewQSqlTableModelFromPointer(ptr).SortDefault(int(column), core.Qt__SortOrder(order))
	}
}

func (ptr *QSqlTableModel) Sort(column int, order core.Qt__SortOrder) {
	defer qt.Recovering("QSqlTableModel::sort")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_Sort(ptr.Pointer(), C.int(column), C.int(order))
	}
}

func (ptr *QSqlTableModel) SortDefault(column int, order core.Qt__SortOrder) {
	defer qt.Recovering("QSqlTableModel::sort")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_SortDefault(ptr.Pointer(), C.int(column), C.int(order))
	}
}

func (ptr *QSqlTableModel) Submit() bool {
	defer qt.Recovering("QSqlTableModel::submit")

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_Submit(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlTableModel) SubmitAll() bool {
	defer qt.Recovering("QSqlTableModel::submitAll")

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_SubmitAll(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlTableModel) TableName() string {
	defer qt.Recovering("QSqlTableModel::tableName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlTableModel_TableName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlTableModel) UpdateRowInTable(row int, values QSqlRecord_ITF) bool {
	defer qt.Recovering("QSqlTableModel::updateRowInTable")

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_UpdateRowInTable(ptr.Pointer(), C.int(row), PointerFromQSqlRecord(values)) != 0
	}
	return false
}

func (ptr *QSqlTableModel) DestroyQSqlTableModel() {
	defer qt.Recovering("QSqlTableModel::~QSqlTableModel")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_DestroyQSqlTableModel(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSqlTableModel) ConnectFetchMore(f func(parent *core.QModelIndex)) {
	defer qt.Recovering("connect QSqlTableModel::fetchMore")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "fetchMore", f)
	}
}

func (ptr *QSqlTableModel) DisconnectFetchMore() {
	defer qt.Recovering("disconnect QSqlTableModel::fetchMore")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "fetchMore")
	}
}

//export callbackQSqlTableModelFetchMore
func callbackQSqlTableModelFetchMore(ptr unsafe.Pointer, ptrName *C.char, parent unsafe.Pointer) {
	defer qt.Recovering("callback QSqlTableModel::fetchMore")

	if signal := qt.GetSignal(C.GoString(ptrName), "fetchMore"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(parent))
	} else {
		NewQSqlTableModelFromPointer(ptr).FetchMoreDefault(core.NewQModelIndexFromPointer(parent))
	}
}

func (ptr *QSqlTableModel) FetchMore(parent core.QModelIndex_ITF) {
	defer qt.Recovering("QSqlTableModel::fetchMore")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_FetchMore(ptr.Pointer(), core.PointerFromQModelIndex(parent))
	}
}

func (ptr *QSqlTableModel) FetchMoreDefault(parent core.QModelIndex_ITF) {
	defer qt.Recovering("QSqlTableModel::fetchMore")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_FetchMoreDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent))
	}
}

func (ptr *QSqlTableModel) ConnectQueryChange(f func()) {
	defer qt.Recovering("connect QSqlTableModel::queryChange")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "queryChange", f)
	}
}

func (ptr *QSqlTableModel) DisconnectQueryChange() {
	defer qt.Recovering("disconnect QSqlTableModel::queryChange")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "queryChange")
	}
}

//export callbackQSqlTableModelQueryChange
func callbackQSqlTableModelQueryChange(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSqlTableModel::queryChange")

	if signal := qt.GetSignal(C.GoString(ptrName), "queryChange"); signal != nil {
		signal.(func())()
	} else {
		NewQSqlTableModelFromPointer(ptr).QueryChangeDefault()
	}
}

func (ptr *QSqlTableModel) QueryChange() {
	defer qt.Recovering("QSqlTableModel::queryChange")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_QueryChange(ptr.Pointer())
	}
}

func (ptr *QSqlTableModel) QueryChangeDefault() {
	defer qt.Recovering("QSqlTableModel::queryChange")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_QueryChangeDefault(ptr.Pointer())
	}
}

func (ptr *QSqlTableModel) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QSqlTableModel::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QSqlTableModel) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QSqlTableModel::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQSqlTableModelTimerEvent
func callbackQSqlTableModelTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSqlTableModel::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSqlTableModelFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSqlTableModel) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSqlTableModel::timerEvent")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSqlTableModel) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSqlTableModel::timerEvent")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSqlTableModel) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QSqlTableModel::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QSqlTableModel) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QSqlTableModel::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQSqlTableModelChildEvent
func callbackQSqlTableModelChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSqlTableModel::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSqlTableModelFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSqlTableModel) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSqlTableModel::childEvent")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSqlTableModel) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSqlTableModel::childEvent")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSqlTableModel) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSqlTableModel::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QSqlTableModel) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QSqlTableModel::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQSqlTableModelCustomEvent
func callbackQSqlTableModelCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSqlTableModel::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSqlTableModelFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSqlTableModel) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSqlTableModel::customEvent")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSqlTableModel) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSqlTableModel::customEvent")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
