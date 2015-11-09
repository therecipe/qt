package sql

//#include "qsqldriver.h"
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
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QSqlDriver_" + qt.RandomIdentifier())
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
	if ptr.Pointer() != nil {
		return C.QSqlDriver_BeginTransaction(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlDriver) Close() {
	if ptr.Pointer() != nil {
		C.QSqlDriver_Close(ptr.Pointer())
	}
}

func (ptr *QSqlDriver) CommitTransaction() bool {
	if ptr.Pointer() != nil {
		return C.QSqlDriver_CommitTransaction(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlDriver) CreateResult() *QSqlResult {
	if ptr.Pointer() != nil {
		return NewQSqlResultFromPointer(C.QSqlDriver_CreateResult(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSqlDriver) DbmsType() QSqlDriver__DbmsType {
	if ptr.Pointer() != nil {
		return QSqlDriver__DbmsType(C.QSqlDriver_DbmsType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSqlDriver) EscapeIdentifier(identifier string, ty QSqlDriver__IdentifierType) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlDriver_EscapeIdentifier(ptr.Pointer(), C.CString(identifier), C.int(ty)))
	}
	return ""
}

func (ptr *QSqlDriver) FormatValue(field QSqlField_ITF, trimStrings bool) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlDriver_FormatValue(ptr.Pointer(), PointerFromQSqlField(field), C.int(qt.GoBoolToInt(trimStrings))))
	}
	return ""
}

func (ptr *QSqlDriver) Handle() *core.QVariant {
	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QSqlDriver_Handle(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSqlDriver) HasFeature(feature QSqlDriver__DriverFeature) bool {
	if ptr.Pointer() != nil {
		return C.QSqlDriver_HasFeature(ptr.Pointer(), C.int(feature)) != 0
	}
	return false
}

func (ptr *QSqlDriver) IsIdentifierEscaped(identifier string, ty QSqlDriver__IdentifierType) bool {
	if ptr.Pointer() != nil {
		return C.QSqlDriver_IsIdentifierEscaped(ptr.Pointer(), C.CString(identifier), C.int(ty)) != 0
	}
	return false
}

func (ptr *QSqlDriver) IsOpen() bool {
	if ptr.Pointer() != nil {
		return C.QSqlDriver_IsOpen(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlDriver) IsOpenError() bool {
	if ptr.Pointer() != nil {
		return C.QSqlDriver_IsOpenError(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlDriver) ConnectNotification(f func(name string)) {
	if ptr.Pointer() != nil {
		C.QSqlDriver_ConnectNotification(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "notification", f)
	}
}

func (ptr *QSqlDriver) DisconnectNotification() {
	if ptr.Pointer() != nil {
		C.QSqlDriver_DisconnectNotification(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "notification")
	}
}

//export callbackQSqlDriverNotification
func callbackQSqlDriverNotification(ptrName *C.char, name *C.char) {
	qt.GetSignal(C.GoString(ptrName), "notification").(func(string))(C.GoString(name))
}

func (ptr *QSqlDriver) Open(db string, user string, password string, host string, port int, options string) bool {
	if ptr.Pointer() != nil {
		return C.QSqlDriver_Open(ptr.Pointer(), C.CString(db), C.CString(user), C.CString(password), C.CString(host), C.int(port), C.CString(options)) != 0
	}
	return false
}

func (ptr *QSqlDriver) RollbackTransaction() bool {
	if ptr.Pointer() != nil {
		return C.QSqlDriver_RollbackTransaction(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlDriver) SqlStatement(ty QSqlDriver__StatementType, tableName string, rec QSqlRecord_ITF, preparedStatement bool) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlDriver_SqlStatement(ptr.Pointer(), C.int(ty), C.CString(tableName), PointerFromQSqlRecord(rec), C.int(qt.GoBoolToInt(preparedStatement))))
	}
	return ""
}

func (ptr *QSqlDriver) StripDelimiters(identifier string, ty QSqlDriver__IdentifierType) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlDriver_StripDelimiters(ptr.Pointer(), C.CString(identifier), C.int(ty)))
	}
	return ""
}

func (ptr *QSqlDriver) SubscribeToNotification(name string) bool {
	if ptr.Pointer() != nil {
		return C.QSqlDriver_SubscribeToNotification(ptr.Pointer(), C.CString(name)) != 0
	}
	return false
}

func (ptr *QSqlDriver) SubscribedToNotifications() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QSqlDriver_SubscribedToNotifications(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QSqlDriver) UnsubscribeFromNotification(name string) bool {
	if ptr.Pointer() != nil {
		return C.QSqlDriver_UnsubscribeFromNotification(ptr.Pointer(), C.CString(name)) != 0
	}
	return false
}

func (ptr *QSqlDriver) DestroyQSqlDriver() {
	if ptr.Pointer() != nil {
		C.QSqlDriver_DestroyQSqlDriver(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
