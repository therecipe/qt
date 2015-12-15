package sql

//#include "sql.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

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
func callbackQSqlDriverNotification(ptrName *C.char, name *C.char) {
	defer qt.Recovering("callback QSqlDriver::notification")

	var signal = qt.GetSignal(C.GoString(ptrName), "notification")
	if signal != nil {
		signal.(func(string))(C.GoString(name))
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
func callbackQSqlDriverSetOpen(ptrName *C.char, open C.int) bool {
	defer qt.Recovering("callback QSqlDriver::setOpen")

	var signal = qt.GetSignal(C.GoString(ptrName), "setOpen")
	if signal != nil {
		defer signal.(func(bool))(int(open) != 0)
		return true
	}
	return false

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
func callbackQSqlDriverSetOpenError(ptrName *C.char, error C.int) bool {
	defer qt.Recovering("callback QSqlDriver::setOpenError")

	var signal = qt.GetSignal(C.GoString(ptrName), "setOpenError")
	if signal != nil {
		defer signal.(func(bool))(int(error) != 0)
		return true
	}
	return false

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
		return strings.Split(C.GoString(C.QSqlDriver_SubscribedToNotifications(ptr.Pointer())), ",,,")
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
