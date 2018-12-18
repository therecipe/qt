// +build !minimal

package sql

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "sql.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"runtime"
	"strings"
	"unsafe"
)

func cGoUnpackString(s C.struct_QtSql_PackedString) string {
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}
func cGoUnpackBytes(s C.struct_QtSql_PackedString) []byte {
	if int(s.len) == -1 {
		return []byte(C.GoString(s.data))
	}
	return C.GoBytes(unsafe.Pointer(s.data), C.int(s.len))
}

type QSql struct {
	ptr unsafe.Pointer
}

type QSql_ITF interface {
	QSql_PTR() *QSql
}

func (ptr *QSql) QSql_PTR() *QSql {
	return ptr
}

func (ptr *QSql) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QSql) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQSql(ptr QSql_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSql_PTR().Pointer()
	}
	return nil
}

func NewQSqlFromPointer(ptr unsafe.Pointer) (n *QSql) {
	n = new(QSql)
	n.SetPointer(ptr)
	return
}

func (ptr *QSql) DestroyQSql() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
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
	ptr unsafe.Pointer
}

type QSqlDatabase_ITF interface {
	QSqlDatabase_PTR() *QSqlDatabase
}

func (ptr *QSqlDatabase) QSqlDatabase_PTR() *QSqlDatabase {
	return ptr
}

func (ptr *QSqlDatabase) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QSqlDatabase) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQSqlDatabase(ptr QSqlDatabase_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlDatabase_PTR().Pointer()
	}
	return nil
}

func NewQSqlDatabaseFromPointer(ptr unsafe.Pointer) (n *QSqlDatabase) {
	n = new(QSqlDatabase)
	n.SetPointer(ptr)
	return
}
func QSqlDatabase_AddDatabase2(driver QSqlDriver_ITF, connectionName string) *QSqlDatabase {
	var connectionNameC *C.char
	if connectionName != "" {
		connectionNameC = C.CString(connectionName)
		defer C.free(unsafe.Pointer(connectionNameC))
	}
	tmpValue := NewQSqlDatabaseFromPointer(C.QSqlDatabase_QSqlDatabase_AddDatabase2(PointerFromQSqlDriver(driver), C.struct_QtSql_PackedString{data: connectionNameC, len: C.longlong(len(connectionName))}))
	runtime.SetFinalizer(tmpValue, (*QSqlDatabase).DestroyQSqlDatabase)
	return tmpValue
}

func (ptr *QSqlDatabase) AddDatabase2(driver QSqlDriver_ITF, connectionName string) *QSqlDatabase {
	var connectionNameC *C.char
	if connectionName != "" {
		connectionNameC = C.CString(connectionName)
		defer C.free(unsafe.Pointer(connectionNameC))
	}
	tmpValue := NewQSqlDatabaseFromPointer(C.QSqlDatabase_QSqlDatabase_AddDatabase2(PointerFromQSqlDriver(driver), C.struct_QtSql_PackedString{data: connectionNameC, len: C.longlong(len(connectionName))}))
	runtime.SetFinalizer(tmpValue, (*QSqlDatabase).DestroyQSqlDatabase)
	return tmpValue
}

func QSqlDatabase_AddDatabase(ty string, connectionName string) *QSqlDatabase {
	var tyC *C.char
	if ty != "" {
		tyC = C.CString(ty)
		defer C.free(unsafe.Pointer(tyC))
	}
	var connectionNameC *C.char
	if connectionName != "" {
		connectionNameC = C.CString(connectionName)
		defer C.free(unsafe.Pointer(connectionNameC))
	}
	tmpValue := NewQSqlDatabaseFromPointer(C.QSqlDatabase_QSqlDatabase_AddDatabase(C.struct_QtSql_PackedString{data: tyC, len: C.longlong(len(ty))}, C.struct_QtSql_PackedString{data: connectionNameC, len: C.longlong(len(connectionName))}))
	runtime.SetFinalizer(tmpValue, (*QSqlDatabase).DestroyQSqlDatabase)
	return tmpValue
}

func (ptr *QSqlDatabase) AddDatabase(ty string, connectionName string) *QSqlDatabase {
	var tyC *C.char
	if ty != "" {
		tyC = C.CString(ty)
		defer C.free(unsafe.Pointer(tyC))
	}
	var connectionNameC *C.char
	if connectionName != "" {
		connectionNameC = C.CString(connectionName)
		defer C.free(unsafe.Pointer(connectionNameC))
	}
	tmpValue := NewQSqlDatabaseFromPointer(C.QSqlDatabase_QSqlDatabase_AddDatabase(C.struct_QtSql_PackedString{data: tyC, len: C.longlong(len(ty))}, C.struct_QtSql_PackedString{data: connectionNameC, len: C.longlong(len(connectionName))}))
	runtime.SetFinalizer(tmpValue, (*QSqlDatabase).DestroyQSqlDatabase)
	return tmpValue
}

func QSqlDatabase_CloneDatabase(other QSqlDatabase_ITF, connectionName string) *QSqlDatabase {
	var connectionNameC *C.char
	if connectionName != "" {
		connectionNameC = C.CString(connectionName)
		defer C.free(unsafe.Pointer(connectionNameC))
	}
	tmpValue := NewQSqlDatabaseFromPointer(C.QSqlDatabase_QSqlDatabase_CloneDatabase(PointerFromQSqlDatabase(other), C.struct_QtSql_PackedString{data: connectionNameC, len: C.longlong(len(connectionName))}))
	runtime.SetFinalizer(tmpValue, (*QSqlDatabase).DestroyQSqlDatabase)
	return tmpValue
}

func (ptr *QSqlDatabase) CloneDatabase(other QSqlDatabase_ITF, connectionName string) *QSqlDatabase {
	var connectionNameC *C.char
	if connectionName != "" {
		connectionNameC = C.CString(connectionName)
		defer C.free(unsafe.Pointer(connectionNameC))
	}
	tmpValue := NewQSqlDatabaseFromPointer(C.QSqlDatabase_QSqlDatabase_CloneDatabase(PointerFromQSqlDatabase(other), C.struct_QtSql_PackedString{data: connectionNameC, len: C.longlong(len(connectionName))}))
	runtime.SetFinalizer(tmpValue, (*QSqlDatabase).DestroyQSqlDatabase)
	return tmpValue
}

func QSqlDatabase_Database(connectionName string, open bool) *QSqlDatabase {
	var connectionNameC *C.char
	if connectionName != "" {
		connectionNameC = C.CString(connectionName)
		defer C.free(unsafe.Pointer(connectionNameC))
	}
	tmpValue := NewQSqlDatabaseFromPointer(C.QSqlDatabase_QSqlDatabase_Database(C.struct_QtSql_PackedString{data: connectionNameC, len: C.longlong(len(connectionName))}, C.char(int8(qt.GoBoolToInt(open)))))
	runtime.SetFinalizer(tmpValue, (*QSqlDatabase).DestroyQSqlDatabase)
	return tmpValue
}

func (ptr *QSqlDatabase) Database(connectionName string, open bool) *QSqlDatabase {
	var connectionNameC *C.char
	if connectionName != "" {
		connectionNameC = C.CString(connectionName)
		defer C.free(unsafe.Pointer(connectionNameC))
	}
	tmpValue := NewQSqlDatabaseFromPointer(C.QSqlDatabase_QSqlDatabase_Database(C.struct_QtSql_PackedString{data: connectionNameC, len: C.longlong(len(connectionName))}, C.char(int8(qt.GoBoolToInt(open)))))
	runtime.SetFinalizer(tmpValue, (*QSqlDatabase).DestroyQSqlDatabase)
	return tmpValue
}

func NewQSqlDatabase() *QSqlDatabase {
	tmpValue := NewQSqlDatabaseFromPointer(C.QSqlDatabase_NewQSqlDatabase())
	runtime.SetFinalizer(tmpValue, (*QSqlDatabase).DestroyQSqlDatabase)
	return tmpValue
}

func NewQSqlDatabase4(driver QSqlDriver_ITF) *QSqlDatabase {
	tmpValue := NewQSqlDatabaseFromPointer(C.QSqlDatabase_NewQSqlDatabase4(PointerFromQSqlDriver(driver)))
	runtime.SetFinalizer(tmpValue, (*QSqlDatabase).DestroyQSqlDatabase)
	return tmpValue
}

func NewQSqlDatabase2(other QSqlDatabase_ITF) *QSqlDatabase {
	tmpValue := NewQSqlDatabaseFromPointer(C.QSqlDatabase_NewQSqlDatabase2(PointerFromQSqlDatabase(other)))
	runtime.SetFinalizer(tmpValue, (*QSqlDatabase).DestroyQSqlDatabase)
	return tmpValue
}

func NewQSqlDatabase3(ty string) *QSqlDatabase {
	var tyC *C.char
	if ty != "" {
		tyC = C.CString(ty)
		defer C.free(unsafe.Pointer(tyC))
	}
	tmpValue := NewQSqlDatabaseFromPointer(C.QSqlDatabase_NewQSqlDatabase3(C.struct_QtSql_PackedString{data: tyC, len: C.longlong(len(ty))}))
	runtime.SetFinalizer(tmpValue, (*QSqlDatabase).DestroyQSqlDatabase)
	return tmpValue
}

func QSqlDatabase_ConnectionNames() []string {
	return strings.Split(cGoUnpackString(C.QSqlDatabase_QSqlDatabase_ConnectionNames()), "|")
}

func (ptr *QSqlDatabase) ConnectionNames() []string {
	return strings.Split(cGoUnpackString(C.QSqlDatabase_QSqlDatabase_ConnectionNames()), "|")
}

func QSqlDatabase_Drivers() []string {
	return strings.Split(cGoUnpackString(C.QSqlDatabase_QSqlDatabase_Drivers()), "|")
}

func (ptr *QSqlDatabase) Drivers() []string {
	return strings.Split(cGoUnpackString(C.QSqlDatabase_QSqlDatabase_Drivers()), "|")
}

func (ptr *QSqlDatabase) Commit() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlDatabase_Commit(ptr.Pointer())) != 0
	}
	return false
}

func QSqlDatabase_Contains(connectionName string) bool {
	var connectionNameC *C.char
	if connectionName != "" {
		connectionNameC = C.CString(connectionName)
		defer C.free(unsafe.Pointer(connectionNameC))
	}
	return int8(C.QSqlDatabase_QSqlDatabase_Contains(C.struct_QtSql_PackedString{data: connectionNameC, len: C.longlong(len(connectionName))})) != 0
}

func (ptr *QSqlDatabase) Contains(connectionName string) bool {
	var connectionNameC *C.char
	if connectionName != "" {
		connectionNameC = C.CString(connectionName)
		defer C.free(unsafe.Pointer(connectionNameC))
	}
	return int8(C.QSqlDatabase_QSqlDatabase_Contains(C.struct_QtSql_PackedString{data: connectionNameC, len: C.longlong(len(connectionName))})) != 0
}

func QSqlDatabase_IsDriverAvailable(name string) bool {
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	return int8(C.QSqlDatabase_QSqlDatabase_IsDriverAvailable(C.struct_QtSql_PackedString{data: nameC, len: C.longlong(len(name))})) != 0
}

func (ptr *QSqlDatabase) IsDriverAvailable(name string) bool {
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	return int8(C.QSqlDatabase_QSqlDatabase_IsDriverAvailable(C.struct_QtSql_PackedString{data: nameC, len: C.longlong(len(name))})) != 0
}

func (ptr *QSqlDatabase) Open() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlDatabase_Open(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSqlDatabase) Open2(user string, password string) bool {
	if ptr.Pointer() != nil {
		var userC *C.char
		if user != "" {
			userC = C.CString(user)
			defer C.free(unsafe.Pointer(userC))
		}
		var passwordC *C.char
		if password != "" {
			passwordC = C.CString(password)
			defer C.free(unsafe.Pointer(passwordC))
		}
		return int8(C.QSqlDatabase_Open2(ptr.Pointer(), C.struct_QtSql_PackedString{data: userC, len: C.longlong(len(user))}, C.struct_QtSql_PackedString{data: passwordC, len: C.longlong(len(password))})) != 0
	}
	return false
}

func (ptr *QSqlDatabase) Rollback() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlDatabase_Rollback(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSqlDatabase) Transaction() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlDatabase_Transaction(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSqlDatabase) Close() {
	if ptr.Pointer() != nil {
		C.QSqlDatabase_Close(ptr.Pointer())
	}
}

func QSqlDatabase_RegisterSqlDriver(name string, creator QSqlDriverCreatorBase_ITF) {
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	C.QSqlDatabase_QSqlDatabase_RegisterSqlDriver(C.struct_QtSql_PackedString{data: nameC, len: C.longlong(len(name))}, PointerFromQSqlDriverCreatorBase(creator))
}

func (ptr *QSqlDatabase) RegisterSqlDriver(name string, creator QSqlDriverCreatorBase_ITF) {
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	C.QSqlDatabase_QSqlDatabase_RegisterSqlDriver(C.struct_QtSql_PackedString{data: nameC, len: C.longlong(len(name))}, PointerFromQSqlDriverCreatorBase(creator))
}

func QSqlDatabase_RemoveDatabase(connectionName string) {
	var connectionNameC *C.char
	if connectionName != "" {
		connectionNameC = C.CString(connectionName)
		defer C.free(unsafe.Pointer(connectionNameC))
	}
	C.QSqlDatabase_QSqlDatabase_RemoveDatabase(C.struct_QtSql_PackedString{data: connectionNameC, len: C.longlong(len(connectionName))})
}

func (ptr *QSqlDatabase) RemoveDatabase(connectionName string) {
	var connectionNameC *C.char
	if connectionName != "" {
		connectionNameC = C.CString(connectionName)
		defer C.free(unsafe.Pointer(connectionNameC))
	}
	C.QSqlDatabase_QSqlDatabase_RemoveDatabase(C.struct_QtSql_PackedString{data: connectionNameC, len: C.longlong(len(connectionName))})
}

func (ptr *QSqlDatabase) SetConnectOptions(options string) {
	if ptr.Pointer() != nil {
		var optionsC *C.char
		if options != "" {
			optionsC = C.CString(options)
			defer C.free(unsafe.Pointer(optionsC))
		}
		C.QSqlDatabase_SetConnectOptions(ptr.Pointer(), C.struct_QtSql_PackedString{data: optionsC, len: C.longlong(len(options))})
	}
}

func (ptr *QSqlDatabase) SetDatabaseName(name string) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		C.QSqlDatabase_SetDatabaseName(ptr.Pointer(), C.struct_QtSql_PackedString{data: nameC, len: C.longlong(len(name))})
	}
}

func (ptr *QSqlDatabase) SetHostName(host string) {
	if ptr.Pointer() != nil {
		var hostC *C.char
		if host != "" {
			hostC = C.CString(host)
			defer C.free(unsafe.Pointer(hostC))
		}
		C.QSqlDatabase_SetHostName(ptr.Pointer(), C.struct_QtSql_PackedString{data: hostC, len: C.longlong(len(host))})
	}
}

func (ptr *QSqlDatabase) SetNumericalPrecisionPolicy(precisionPolicy QSql__NumericalPrecisionPolicy) {
	if ptr.Pointer() != nil {
		C.QSqlDatabase_SetNumericalPrecisionPolicy(ptr.Pointer(), C.longlong(precisionPolicy))
	}
}

func (ptr *QSqlDatabase) SetPassword(password string) {
	if ptr.Pointer() != nil {
		var passwordC *C.char
		if password != "" {
			passwordC = C.CString(password)
			defer C.free(unsafe.Pointer(passwordC))
		}
		C.QSqlDatabase_SetPassword(ptr.Pointer(), C.struct_QtSql_PackedString{data: passwordC, len: C.longlong(len(password))})
	}
}

func (ptr *QSqlDatabase) SetPort(port int) {
	if ptr.Pointer() != nil {
		C.QSqlDatabase_SetPort(ptr.Pointer(), C.int(int32(port)))
	}
}

func (ptr *QSqlDatabase) SetUserName(name string) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		C.QSqlDatabase_SetUserName(ptr.Pointer(), C.struct_QtSql_PackedString{data: nameC, len: C.longlong(len(name))})
	}
}

func (ptr *QSqlDatabase) DestroyQSqlDatabase() {
	if ptr.Pointer() != nil {
		C.QSqlDatabase_DestroyQSqlDatabase(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QSqlDatabase) NumericalPrecisionPolicy() QSql__NumericalPrecisionPolicy {
	if ptr.Pointer() != nil {
		return QSql__NumericalPrecisionPolicy(C.QSqlDatabase_NumericalPrecisionPolicy(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSqlDatabase) Driver() *QSqlDriver {
	if ptr.Pointer() != nil {
		tmpValue := NewQSqlDriverFromPointer(C.QSqlDatabase_Driver(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSqlDatabase) LastError() *QSqlError {
	if ptr.Pointer() != nil {
		tmpValue := NewQSqlErrorFromPointer(C.QSqlDatabase_LastError(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QSqlError).DestroyQSqlError)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlDatabase) PrimaryIndex(tablename string) *QSqlIndex {
	if ptr.Pointer() != nil {
		var tablenameC *C.char
		if tablename != "" {
			tablenameC = C.CString(tablename)
			defer C.free(unsafe.Pointer(tablenameC))
		}
		tmpValue := NewQSqlIndexFromPointer(C.QSqlDatabase_PrimaryIndex(ptr.Pointer(), C.struct_QtSql_PackedString{data: tablenameC, len: C.longlong(len(tablename))}))
		runtime.SetFinalizer(tmpValue, (*QSqlIndex).DestroyQSqlIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlDatabase) Exec(query string) *QSqlQuery {
	if ptr.Pointer() != nil {
		var queryC *C.char
		if query != "" {
			queryC = C.CString(query)
			defer C.free(unsafe.Pointer(queryC))
		}
		tmpValue := NewQSqlQueryFromPointer(C.QSqlDatabase_Exec(ptr.Pointer(), C.struct_QtSql_PackedString{data: queryC, len: C.longlong(len(query))}))
		runtime.SetFinalizer(tmpValue, (*QSqlQuery).DestroyQSqlQuery)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlDatabase) Record(tablename string) *QSqlRecord {
	if ptr.Pointer() != nil {
		var tablenameC *C.char
		if tablename != "" {
			tablenameC = C.CString(tablename)
			defer C.free(unsafe.Pointer(tablenameC))
		}
		tmpValue := NewQSqlRecordFromPointer(C.QSqlDatabase_Record(ptr.Pointer(), C.struct_QtSql_PackedString{data: tablenameC, len: C.longlong(len(tablename))}))
		runtime.SetFinalizer(tmpValue, (*QSqlRecord).DestroyQSqlRecord)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlDatabase) ConnectOptions() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSqlDatabase_ConnectOptions(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlDatabase) ConnectionName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSqlDatabase_ConnectionName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlDatabase) DatabaseName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSqlDatabase_DatabaseName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlDatabase) DriverName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSqlDatabase_DriverName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlDatabase) HostName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSqlDatabase_HostName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlDatabase) Password() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSqlDatabase_Password(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlDatabase) UserName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSqlDatabase_UserName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlDatabase) Tables(ty QSql__TableType) []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QSqlDatabase_Tables(ptr.Pointer(), C.longlong(ty))), "|")
	}
	return make([]string, 0)
}

func (ptr *QSqlDatabase) IsOpen() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlDatabase_IsOpen(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSqlDatabase) IsOpenError() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlDatabase_IsOpenError(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSqlDatabase) IsValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlDatabase_IsValid(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSqlDatabase) Port() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSqlDatabase_Port(ptr.Pointer())))
	}
	return 0
}

func QSqlDatabase_DefaultConnection() string {
	return cGoUnpackString(C.QSqlDatabase_QSqlDatabase_DefaultConnection())
}

func (ptr *QSqlDatabase) DefaultConnection() string {
	return cGoUnpackString(C.QSqlDatabase_QSqlDatabase_DefaultConnection())
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

func NewQSqlDriverFromPointer(ptr unsafe.Pointer) (n *QSqlDriver) {
	n = new(QSqlDriver)
	n.SetPointer(ptr)
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

func NewQSqlDriver(parent core.QObject_ITF) *QSqlDriver {
	tmpValue := NewQSqlDriverFromPointer(C.QSqlDriver_NewQSqlDriver(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func QSqlDriver_Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QSqlDriver_QSqlDriver_Tr(sC, cC, C.int(int32(n))))
}

func (ptr *QSqlDriver) Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QSqlDriver_QSqlDriver_Tr(sC, cC, C.int(int32(n))))
}

func QSqlDriver_TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QSqlDriver_QSqlDriver_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QSqlDriver) TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QSqlDriver_QSqlDriver_TrUtf8(sC, cC, C.int(int32(n))))
}

//export callbackQSqlDriver_BeginTransaction
func callbackQSqlDriver_BeginTransaction(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "beginTransaction"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlDriverFromPointer(ptr).BeginTransactionDefault())))
}

func (ptr *QSqlDriver) ConnectBeginTransaction(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "beginTransaction"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "beginTransaction", func() bool {
				signal.(func() bool)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "beginTransaction", f)
		}
	}
}

func (ptr *QSqlDriver) DisconnectBeginTransaction() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "beginTransaction")
	}
}

func (ptr *QSqlDriver) BeginTransaction() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlDriver_BeginTransaction(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSqlDriver) BeginTransactionDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlDriver_BeginTransactionDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQSqlDriver_CommitTransaction
func callbackQSqlDriver_CommitTransaction(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "commitTransaction"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlDriverFromPointer(ptr).CommitTransactionDefault())))
}

func (ptr *QSqlDriver) ConnectCommitTransaction(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "commitTransaction"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "commitTransaction", func() bool {
				signal.(func() bool)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "commitTransaction", f)
		}
	}
}

func (ptr *QSqlDriver) DisconnectCommitTransaction() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "commitTransaction")
	}
}

func (ptr *QSqlDriver) CommitTransaction() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlDriver_CommitTransaction(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSqlDriver) CommitTransactionDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlDriver_CommitTransactionDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQSqlDriver_Open
func callbackQSqlDriver_Open(ptr unsafe.Pointer, db C.struct_QtSql_PackedString, user C.struct_QtSql_PackedString, password C.struct_QtSql_PackedString, host C.struct_QtSql_PackedString, port C.int, options C.struct_QtSql_PackedString) C.char {
	if signal := qt.GetSignal(ptr, "open"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string, string, string, string, int, string) bool)(cGoUnpackString(db), cGoUnpackString(user), cGoUnpackString(password), cGoUnpackString(host), int(int32(port)), cGoUnpackString(options)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QSqlDriver) ConnectOpen(f func(db string, user string, password string, host string, port int, options string) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "open"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "open", func(db string, user string, password string, host string, port int, options string) bool {
				signal.(func(string, string, string, string, int, string) bool)(db, user, password, host, port, options)
				return f(db, user, password, host, port, options)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "open", f)
		}
	}
}

func (ptr *QSqlDriver) DisconnectOpen() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "open")
	}
}

func (ptr *QSqlDriver) Open(db string, user string, password string, host string, port int, options string) bool {
	if ptr.Pointer() != nil {
		var dbC *C.char
		if db != "" {
			dbC = C.CString(db)
			defer C.free(unsafe.Pointer(dbC))
		}
		var userC *C.char
		if user != "" {
			userC = C.CString(user)
			defer C.free(unsafe.Pointer(userC))
		}
		var passwordC *C.char
		if password != "" {
			passwordC = C.CString(password)
			defer C.free(unsafe.Pointer(passwordC))
		}
		var hostC *C.char
		if host != "" {
			hostC = C.CString(host)
			defer C.free(unsafe.Pointer(hostC))
		}
		var optionsC *C.char
		if options != "" {
			optionsC = C.CString(options)
			defer C.free(unsafe.Pointer(optionsC))
		}
		return int8(C.QSqlDriver_Open(ptr.Pointer(), C.struct_QtSql_PackedString{data: dbC, len: C.longlong(len(db))}, C.struct_QtSql_PackedString{data: userC, len: C.longlong(len(user))}, C.struct_QtSql_PackedString{data: passwordC, len: C.longlong(len(password))}, C.struct_QtSql_PackedString{data: hostC, len: C.longlong(len(host))}, C.int(int32(port)), C.struct_QtSql_PackedString{data: optionsC, len: C.longlong(len(options))})) != 0
	}
	return false
}

//export callbackQSqlDriver_RollbackTransaction
func callbackQSqlDriver_RollbackTransaction(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "rollbackTransaction"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlDriverFromPointer(ptr).RollbackTransactionDefault())))
}

func (ptr *QSqlDriver) ConnectRollbackTransaction(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "rollbackTransaction"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "rollbackTransaction", func() bool {
				signal.(func() bool)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "rollbackTransaction", f)
		}
	}
}

func (ptr *QSqlDriver) DisconnectRollbackTransaction() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "rollbackTransaction")
	}
}

func (ptr *QSqlDriver) RollbackTransaction() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlDriver_RollbackTransaction(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSqlDriver) RollbackTransactionDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlDriver_RollbackTransactionDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQSqlDriver_SubscribeToNotification
func callbackQSqlDriver_SubscribeToNotification(ptr unsafe.Pointer, name C.struct_QtSql_PackedString) C.char {
	if signal := qt.GetSignal(ptr, "subscribeToNotification"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string) bool)(cGoUnpackString(name)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlDriverFromPointer(ptr).SubscribeToNotificationDefault(cGoUnpackString(name)))))
}

func (ptr *QSqlDriver) ConnectSubscribeToNotification(f func(name string) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "subscribeToNotification"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "subscribeToNotification", func(name string) bool {
				signal.(func(string) bool)(name)
				return f(name)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "subscribeToNotification", f)
		}
	}
}

func (ptr *QSqlDriver) DisconnectSubscribeToNotification() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "subscribeToNotification")
	}
}

func (ptr *QSqlDriver) SubscribeToNotification(name string) bool {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		return int8(C.QSqlDriver_SubscribeToNotification(ptr.Pointer(), C.struct_QtSql_PackedString{data: nameC, len: C.longlong(len(name))})) != 0
	}
	return false
}

func (ptr *QSqlDriver) SubscribeToNotificationDefault(name string) bool {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		return int8(C.QSqlDriver_SubscribeToNotificationDefault(ptr.Pointer(), C.struct_QtSql_PackedString{data: nameC, len: C.longlong(len(name))})) != 0
	}
	return false
}

//export callbackQSqlDriver_UnsubscribeFromNotification
func callbackQSqlDriver_UnsubscribeFromNotification(ptr unsafe.Pointer, name C.struct_QtSql_PackedString) C.char {
	if signal := qt.GetSignal(ptr, "unsubscribeFromNotification"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string) bool)(cGoUnpackString(name)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlDriverFromPointer(ptr).UnsubscribeFromNotificationDefault(cGoUnpackString(name)))))
}

func (ptr *QSqlDriver) ConnectUnsubscribeFromNotification(f func(name string) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "unsubscribeFromNotification"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "unsubscribeFromNotification", func(name string) bool {
				signal.(func(string) bool)(name)
				return f(name)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "unsubscribeFromNotification", f)
		}
	}
}

func (ptr *QSqlDriver) DisconnectUnsubscribeFromNotification() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "unsubscribeFromNotification")
	}
}

func (ptr *QSqlDriver) UnsubscribeFromNotification(name string) bool {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		return int8(C.QSqlDriver_UnsubscribeFromNotification(ptr.Pointer(), C.struct_QtSql_PackedString{data: nameC, len: C.longlong(len(name))})) != 0
	}
	return false
}

func (ptr *QSqlDriver) UnsubscribeFromNotificationDefault(name string) bool {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		return int8(C.QSqlDriver_UnsubscribeFromNotificationDefault(ptr.Pointer(), C.struct_QtSql_PackedString{data: nameC, len: C.longlong(len(name))})) != 0
	}
	return false
}

//export callbackQSqlDriver_Close
func callbackQSqlDriver_Close(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "close"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QSqlDriver) ConnectClose(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "close"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "close", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "close", f)
		}
	}
}

func (ptr *QSqlDriver) DisconnectClose() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "close")
	}
}

func (ptr *QSqlDriver) Close() {
	if ptr.Pointer() != nil {
		C.QSqlDriver_Close(ptr.Pointer())
	}
}

//export callbackQSqlDriver_Notification
func callbackQSqlDriver_Notification(ptr unsafe.Pointer, name C.struct_QtSql_PackedString) {
	if signal := qt.GetSignal(ptr, "notification"); signal != nil {
		signal.(func(string))(cGoUnpackString(name))
	}

}

func (ptr *QSqlDriver) ConnectNotification(f func(name string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "notification") {
			C.QSqlDriver_ConnectNotification(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "notification"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "notification", func(name string) {
				signal.(func(string))(name)
				f(name)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "notification", f)
		}
	}
}

func (ptr *QSqlDriver) DisconnectNotification() {
	if ptr.Pointer() != nil {
		C.QSqlDriver_DisconnectNotification(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "notification")
	}
}

func (ptr *QSqlDriver) Notification(name string) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		C.QSqlDriver_Notification(ptr.Pointer(), C.struct_QtSql_PackedString{data: nameC, len: C.longlong(len(name))})
	}
}

//export callbackQSqlDriver_Notification2
func callbackQSqlDriver_Notification2(ptr unsafe.Pointer, name C.struct_QtSql_PackedString, source C.longlong, payload unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "notification2"); signal != nil {
		signal.(func(string, QSqlDriver__NotificationSource, *core.QVariant))(cGoUnpackString(name), QSqlDriver__NotificationSource(source), core.NewQVariantFromPointer(payload))
	}

}

func (ptr *QSqlDriver) ConnectNotification2(f func(name string, source QSqlDriver__NotificationSource, payload *core.QVariant)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "notification2") {
			C.QSqlDriver_ConnectNotification2(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "notification2"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "notification2", func(name string, source QSqlDriver__NotificationSource, payload *core.QVariant) {
				signal.(func(string, QSqlDriver__NotificationSource, *core.QVariant))(name, source, payload)
				f(name, source, payload)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "notification2", f)
		}
	}
}

func (ptr *QSqlDriver) DisconnectNotification2() {
	if ptr.Pointer() != nil {
		C.QSqlDriver_DisconnectNotification2(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "notification2")
	}
}

func (ptr *QSqlDriver) Notification2(name string, source QSqlDriver__NotificationSource, payload core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		C.QSqlDriver_Notification2(ptr.Pointer(), C.struct_QtSql_PackedString{data: nameC, len: C.longlong(len(name))}, C.longlong(source), core.PointerFromQVariant(payload))
	}
}

//export callbackQSqlDriver_SetLastError
func callbackQSqlDriver_SetLastError(ptr unsafe.Pointer, error unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setLastError"); signal != nil {
		signal.(func(*QSqlError))(NewQSqlErrorFromPointer(error))
	} else {
		NewQSqlDriverFromPointer(ptr).SetLastErrorDefault(NewQSqlErrorFromPointer(error))
	}
}

func (ptr *QSqlDriver) ConnectSetLastError(f func(error *QSqlError)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setLastError"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setLastError", func(error *QSqlError) {
				signal.(func(*QSqlError))(error)
				f(error)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setLastError", f)
		}
	}
}

func (ptr *QSqlDriver) DisconnectSetLastError() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setLastError")
	}
}

func (ptr *QSqlDriver) SetLastError(error QSqlError_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlDriver_SetLastError(ptr.Pointer(), PointerFromQSqlError(error))
	}
}

func (ptr *QSqlDriver) SetLastErrorDefault(error QSqlError_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlDriver_SetLastErrorDefault(ptr.Pointer(), PointerFromQSqlError(error))
	}
}

func (ptr *QSqlDriver) SetNumericalPrecisionPolicy(precisionPolicy QSql__NumericalPrecisionPolicy) {
	if ptr.Pointer() != nil {
		C.QSqlDriver_SetNumericalPrecisionPolicy(ptr.Pointer(), C.longlong(precisionPolicy))
	}
}

//export callbackQSqlDriver_SetOpen
func callbackQSqlDriver_SetOpen(ptr unsafe.Pointer, open C.char) {
	if signal := qt.GetSignal(ptr, "setOpen"); signal != nil {
		signal.(func(bool))(int8(open) != 0)
	} else {
		NewQSqlDriverFromPointer(ptr).SetOpenDefault(int8(open) != 0)
	}
}

func (ptr *QSqlDriver) ConnectSetOpen(f func(open bool)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setOpen"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setOpen", func(open bool) {
				signal.(func(bool))(open)
				f(open)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setOpen", f)
		}
	}
}

func (ptr *QSqlDriver) DisconnectSetOpen() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setOpen")
	}
}

func (ptr *QSqlDriver) SetOpen(open bool) {
	if ptr.Pointer() != nil {
		C.QSqlDriver_SetOpen(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(open))))
	}
}

func (ptr *QSqlDriver) SetOpenDefault(open bool) {
	if ptr.Pointer() != nil {
		C.QSqlDriver_SetOpenDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(open))))
	}
}

//export callbackQSqlDriver_SetOpenError
func callbackQSqlDriver_SetOpenError(ptr unsafe.Pointer, error C.char) {
	if signal := qt.GetSignal(ptr, "setOpenError"); signal != nil {
		signal.(func(bool))(int8(error) != 0)
	} else {
		NewQSqlDriverFromPointer(ptr).SetOpenErrorDefault(int8(error) != 0)
	}
}

func (ptr *QSqlDriver) ConnectSetOpenError(f func(error bool)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setOpenError"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setOpenError", func(error bool) {
				signal.(func(bool))(error)
				f(error)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setOpenError", f)
		}
	}
}

func (ptr *QSqlDriver) DisconnectSetOpenError() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setOpenError")
	}
}

func (ptr *QSqlDriver) SetOpenError(error bool) {
	if ptr.Pointer() != nil {
		C.QSqlDriver_SetOpenError(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(error))))
	}
}

func (ptr *QSqlDriver) SetOpenErrorDefault(error bool) {
	if ptr.Pointer() != nil {
		C.QSqlDriver_SetOpenErrorDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(error))))
	}
}

//export callbackQSqlDriver_DestroyQSqlDriver
func callbackQSqlDriver_DestroyQSqlDriver(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QSqlDriver"); signal != nil {
		signal.(func())()
	} else {
		NewQSqlDriverFromPointer(ptr).DestroyQSqlDriverDefault()
	}
}

func (ptr *QSqlDriver) ConnectDestroyQSqlDriver(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QSqlDriver"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QSqlDriver", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QSqlDriver", f)
		}
	}
}

func (ptr *QSqlDriver) DisconnectDestroyQSqlDriver() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QSqlDriver")
	}
}

func (ptr *QSqlDriver) DestroyQSqlDriver() {
	if ptr.Pointer() != nil {
		C.QSqlDriver_DestroyQSqlDriver(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QSqlDriver) DestroyQSqlDriverDefault() {
	if ptr.Pointer() != nil {
		C.QSqlDriver_DestroyQSqlDriverDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QSqlDriver) NumericalPrecisionPolicy() QSql__NumericalPrecisionPolicy {
	if ptr.Pointer() != nil {
		return QSql__NumericalPrecisionPolicy(C.QSqlDriver_NumericalPrecisionPolicy(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSqlDriver) LastError() *QSqlError {
	if ptr.Pointer() != nil {
		tmpValue := NewQSqlErrorFromPointer(C.QSqlDriver_LastError(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QSqlError).DestroyQSqlError)
		return tmpValue
	}
	return nil
}

//export callbackQSqlDriver_PrimaryIndex
func callbackQSqlDriver_PrimaryIndex(ptr unsafe.Pointer, tableName C.struct_QtSql_PackedString) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "primaryIndex"); signal != nil {
		return PointerFromQSqlIndex(signal.(func(string) *QSqlIndex)(cGoUnpackString(tableName)))
	}

	return PointerFromQSqlIndex(NewQSqlDriverFromPointer(ptr).PrimaryIndexDefault(cGoUnpackString(tableName)))
}

func (ptr *QSqlDriver) ConnectPrimaryIndex(f func(tableName string) *QSqlIndex) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "primaryIndex"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "primaryIndex", func(tableName string) *QSqlIndex {
				signal.(func(string) *QSqlIndex)(tableName)
				return f(tableName)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "primaryIndex", f)
		}
	}
}

func (ptr *QSqlDriver) DisconnectPrimaryIndex() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "primaryIndex")
	}
}

func (ptr *QSqlDriver) PrimaryIndex(tableName string) *QSqlIndex {
	if ptr.Pointer() != nil {
		var tableNameC *C.char
		if tableName != "" {
			tableNameC = C.CString(tableName)
			defer C.free(unsafe.Pointer(tableNameC))
		}
		tmpValue := NewQSqlIndexFromPointer(C.QSqlDriver_PrimaryIndex(ptr.Pointer(), C.struct_QtSql_PackedString{data: tableNameC, len: C.longlong(len(tableName))}))
		runtime.SetFinalizer(tmpValue, (*QSqlIndex).DestroyQSqlIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlDriver) PrimaryIndexDefault(tableName string) *QSqlIndex {
	if ptr.Pointer() != nil {
		var tableNameC *C.char
		if tableName != "" {
			tableNameC = C.CString(tableName)
			defer C.free(unsafe.Pointer(tableNameC))
		}
		tmpValue := NewQSqlIndexFromPointer(C.QSqlDriver_PrimaryIndexDefault(ptr.Pointer(), C.struct_QtSql_PackedString{data: tableNameC, len: C.longlong(len(tableName))}))
		runtime.SetFinalizer(tmpValue, (*QSqlIndex).DestroyQSqlIndex)
		return tmpValue
	}
	return nil
}

//export callbackQSqlDriver_Record
func callbackQSqlDriver_Record(ptr unsafe.Pointer, tableName C.struct_QtSql_PackedString) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "record"); signal != nil {
		return PointerFromQSqlRecord(signal.(func(string) *QSqlRecord)(cGoUnpackString(tableName)))
	}

	return PointerFromQSqlRecord(NewQSqlDriverFromPointer(ptr).RecordDefault(cGoUnpackString(tableName)))
}

func (ptr *QSqlDriver) ConnectRecord(f func(tableName string) *QSqlRecord) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "record"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "record", func(tableName string) *QSqlRecord {
				signal.(func(string) *QSqlRecord)(tableName)
				return f(tableName)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "record", f)
		}
	}
}

func (ptr *QSqlDriver) DisconnectRecord() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "record")
	}
}

func (ptr *QSqlDriver) Record(tableName string) *QSqlRecord {
	if ptr.Pointer() != nil {
		var tableNameC *C.char
		if tableName != "" {
			tableNameC = C.CString(tableName)
			defer C.free(unsafe.Pointer(tableNameC))
		}
		tmpValue := NewQSqlRecordFromPointer(C.QSqlDriver_Record(ptr.Pointer(), C.struct_QtSql_PackedString{data: tableNameC, len: C.longlong(len(tableName))}))
		runtime.SetFinalizer(tmpValue, (*QSqlRecord).DestroyQSqlRecord)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlDriver) RecordDefault(tableName string) *QSqlRecord {
	if ptr.Pointer() != nil {
		var tableNameC *C.char
		if tableName != "" {
			tableNameC = C.CString(tableName)
			defer C.free(unsafe.Pointer(tableNameC))
		}
		tmpValue := NewQSqlRecordFromPointer(C.QSqlDriver_RecordDefault(ptr.Pointer(), C.struct_QtSql_PackedString{data: tableNameC, len: C.longlong(len(tableName))}))
		runtime.SetFinalizer(tmpValue, (*QSqlRecord).DestroyQSqlRecord)
		return tmpValue
	}
	return nil
}

//export callbackQSqlDriver_CreateResult
func callbackQSqlDriver_CreateResult(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "createResult"); signal != nil {
		return PointerFromQSqlResult(signal.(func() *QSqlResult)())
	}

	return PointerFromQSqlResult(NewQSqlResult(nil))
}

func (ptr *QSqlDriver) ConnectCreateResult(f func() *QSqlResult) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "createResult"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "createResult", func() *QSqlResult {
				signal.(func() *QSqlResult)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "createResult", f)
		}
	}
}

func (ptr *QSqlDriver) DisconnectCreateResult() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "createResult")
	}
}

func (ptr *QSqlDriver) CreateResult() *QSqlResult {
	if ptr.Pointer() != nil {
		return NewQSqlResultFromPointer(C.QSqlDriver_CreateResult(ptr.Pointer()))
	}
	return nil
}

//export callbackQSqlDriver_EscapeIdentifier
func callbackQSqlDriver_EscapeIdentifier(ptr unsafe.Pointer, identifier C.struct_QtSql_PackedString, ty C.longlong) C.struct_QtSql_PackedString {
	if signal := qt.GetSignal(ptr, "escapeIdentifier"); signal != nil {
		tempVal := signal.(func(string, QSqlDriver__IdentifierType) string)(cGoUnpackString(identifier), QSqlDriver__IdentifierType(ty))
		return C.struct_QtSql_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewQSqlDriverFromPointer(ptr).EscapeIdentifierDefault(cGoUnpackString(identifier), QSqlDriver__IdentifierType(ty))
	return C.struct_QtSql_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QSqlDriver) ConnectEscapeIdentifier(f func(identifier string, ty QSqlDriver__IdentifierType) string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "escapeIdentifier"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "escapeIdentifier", func(identifier string, ty QSqlDriver__IdentifierType) string {
				signal.(func(string, QSqlDriver__IdentifierType) string)(identifier, ty)
				return f(identifier, ty)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "escapeIdentifier", f)
		}
	}
}

func (ptr *QSqlDriver) DisconnectEscapeIdentifier() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "escapeIdentifier")
	}
}

func (ptr *QSqlDriver) EscapeIdentifier(identifier string, ty QSqlDriver__IdentifierType) string {
	if ptr.Pointer() != nil {
		var identifierC *C.char
		if identifier != "" {
			identifierC = C.CString(identifier)
			defer C.free(unsafe.Pointer(identifierC))
		}
		return cGoUnpackString(C.QSqlDriver_EscapeIdentifier(ptr.Pointer(), C.struct_QtSql_PackedString{data: identifierC, len: C.longlong(len(identifier))}, C.longlong(ty)))
	}
	return ""
}

func (ptr *QSqlDriver) EscapeIdentifierDefault(identifier string, ty QSqlDriver__IdentifierType) string {
	if ptr.Pointer() != nil {
		var identifierC *C.char
		if identifier != "" {
			identifierC = C.CString(identifier)
			defer C.free(unsafe.Pointer(identifierC))
		}
		return cGoUnpackString(C.QSqlDriver_EscapeIdentifierDefault(ptr.Pointer(), C.struct_QtSql_PackedString{data: identifierC, len: C.longlong(len(identifier))}, C.longlong(ty)))
	}
	return ""
}

//export callbackQSqlDriver_FormatValue
func callbackQSqlDriver_FormatValue(ptr unsafe.Pointer, field unsafe.Pointer, trimStrings C.char) C.struct_QtSql_PackedString {
	if signal := qt.GetSignal(ptr, "formatValue"); signal != nil {
		tempVal := signal.(func(*QSqlField, bool) string)(NewQSqlFieldFromPointer(field), int8(trimStrings) != 0)
		return C.struct_QtSql_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewQSqlDriverFromPointer(ptr).FormatValueDefault(NewQSqlFieldFromPointer(field), int8(trimStrings) != 0)
	return C.struct_QtSql_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QSqlDriver) ConnectFormatValue(f func(field *QSqlField, trimStrings bool) string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "formatValue"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "formatValue", func(field *QSqlField, trimStrings bool) string {
				signal.(func(*QSqlField, bool) string)(field, trimStrings)
				return f(field, trimStrings)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "formatValue", f)
		}
	}
}

func (ptr *QSqlDriver) DisconnectFormatValue() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "formatValue")
	}
}

func (ptr *QSqlDriver) FormatValue(field QSqlField_ITF, trimStrings bool) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSqlDriver_FormatValue(ptr.Pointer(), PointerFromQSqlField(field), C.char(int8(qt.GoBoolToInt(trimStrings)))))
	}
	return ""
}

func (ptr *QSqlDriver) FormatValueDefault(field QSqlField_ITF, trimStrings bool) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSqlDriver_FormatValueDefault(ptr.Pointer(), PointerFromQSqlField(field), C.char(int8(qt.GoBoolToInt(trimStrings)))))
	}
	return ""
}

//export callbackQSqlDriver_SqlStatement
func callbackQSqlDriver_SqlStatement(ptr unsafe.Pointer, ty C.longlong, tableName C.struct_QtSql_PackedString, rec unsafe.Pointer, preparedStatement C.char) C.struct_QtSql_PackedString {
	if signal := qt.GetSignal(ptr, "sqlStatement"); signal != nil {
		tempVal := signal.(func(QSqlDriver__StatementType, string, *QSqlRecord, bool) string)(QSqlDriver__StatementType(ty), cGoUnpackString(tableName), NewQSqlRecordFromPointer(rec), int8(preparedStatement) != 0)
		return C.struct_QtSql_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewQSqlDriverFromPointer(ptr).SqlStatementDefault(QSqlDriver__StatementType(ty), cGoUnpackString(tableName), NewQSqlRecordFromPointer(rec), int8(preparedStatement) != 0)
	return C.struct_QtSql_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QSqlDriver) ConnectSqlStatement(f func(ty QSqlDriver__StatementType, tableName string, rec *QSqlRecord, preparedStatement bool) string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "sqlStatement"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "sqlStatement", func(ty QSqlDriver__StatementType, tableName string, rec *QSqlRecord, preparedStatement bool) string {
				signal.(func(QSqlDriver__StatementType, string, *QSqlRecord, bool) string)(ty, tableName, rec, preparedStatement)
				return f(ty, tableName, rec, preparedStatement)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "sqlStatement", f)
		}
	}
}

func (ptr *QSqlDriver) DisconnectSqlStatement() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "sqlStatement")
	}
}

func (ptr *QSqlDriver) SqlStatement(ty QSqlDriver__StatementType, tableName string, rec QSqlRecord_ITF, preparedStatement bool) string {
	if ptr.Pointer() != nil {
		var tableNameC *C.char
		if tableName != "" {
			tableNameC = C.CString(tableName)
			defer C.free(unsafe.Pointer(tableNameC))
		}
		return cGoUnpackString(C.QSqlDriver_SqlStatement(ptr.Pointer(), C.longlong(ty), C.struct_QtSql_PackedString{data: tableNameC, len: C.longlong(len(tableName))}, PointerFromQSqlRecord(rec), C.char(int8(qt.GoBoolToInt(preparedStatement)))))
	}
	return ""
}

func (ptr *QSqlDriver) SqlStatementDefault(ty QSqlDriver__StatementType, tableName string, rec QSqlRecord_ITF, preparedStatement bool) string {
	if ptr.Pointer() != nil {
		var tableNameC *C.char
		if tableName != "" {
			tableNameC = C.CString(tableName)
			defer C.free(unsafe.Pointer(tableNameC))
		}
		return cGoUnpackString(C.QSqlDriver_SqlStatementDefault(ptr.Pointer(), C.longlong(ty), C.struct_QtSql_PackedString{data: tableNameC, len: C.longlong(len(tableName))}, PointerFromQSqlRecord(rec), C.char(int8(qt.GoBoolToInt(preparedStatement)))))
	}
	return ""
}

//export callbackQSqlDriver_StripDelimiters
func callbackQSqlDriver_StripDelimiters(ptr unsafe.Pointer, identifier C.struct_QtSql_PackedString, ty C.longlong) C.struct_QtSql_PackedString {
	if signal := qt.GetSignal(ptr, "stripDelimiters"); signal != nil {
		tempVal := signal.(func(string, QSqlDriver__IdentifierType) string)(cGoUnpackString(identifier), QSqlDriver__IdentifierType(ty))
		return C.struct_QtSql_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewQSqlDriverFromPointer(ptr).StripDelimitersDefault(cGoUnpackString(identifier), QSqlDriver__IdentifierType(ty))
	return C.struct_QtSql_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QSqlDriver) ConnectStripDelimiters(f func(identifier string, ty QSqlDriver__IdentifierType) string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "stripDelimiters"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "stripDelimiters", func(identifier string, ty QSqlDriver__IdentifierType) string {
				signal.(func(string, QSqlDriver__IdentifierType) string)(identifier, ty)
				return f(identifier, ty)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "stripDelimiters", f)
		}
	}
}

func (ptr *QSqlDriver) DisconnectStripDelimiters() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "stripDelimiters")
	}
}

func (ptr *QSqlDriver) StripDelimiters(identifier string, ty QSqlDriver__IdentifierType) string {
	if ptr.Pointer() != nil {
		var identifierC *C.char
		if identifier != "" {
			identifierC = C.CString(identifier)
			defer C.free(unsafe.Pointer(identifierC))
		}
		return cGoUnpackString(C.QSqlDriver_StripDelimiters(ptr.Pointer(), C.struct_QtSql_PackedString{data: identifierC, len: C.longlong(len(identifier))}, C.longlong(ty)))
	}
	return ""
}

func (ptr *QSqlDriver) StripDelimitersDefault(identifier string, ty QSqlDriver__IdentifierType) string {
	if ptr.Pointer() != nil {
		var identifierC *C.char
		if identifier != "" {
			identifierC = C.CString(identifier)
			defer C.free(unsafe.Pointer(identifierC))
		}
		return cGoUnpackString(C.QSqlDriver_StripDelimitersDefault(ptr.Pointer(), C.struct_QtSql_PackedString{data: identifierC, len: C.longlong(len(identifier))}, C.longlong(ty)))
	}
	return ""
}

//export callbackQSqlDriver_SubscribedToNotifications
func callbackQSqlDriver_SubscribedToNotifications(ptr unsafe.Pointer) C.struct_QtSql_PackedString {
	if signal := qt.GetSignal(ptr, "subscribedToNotifications"); signal != nil {
		tempVal := signal.(func() []string)()
		return C.struct_QtSql_PackedString{data: C.CString(strings.Join(tempVal, "|")), len: C.longlong(len(strings.Join(tempVal, "|")))}
	}
	tempVal := NewQSqlDriverFromPointer(ptr).SubscribedToNotificationsDefault()
	return C.struct_QtSql_PackedString{data: C.CString(strings.Join(tempVal, "|")), len: C.longlong(len(strings.Join(tempVal, "|")))}
}

func (ptr *QSqlDriver) ConnectSubscribedToNotifications(f func() []string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "subscribedToNotifications"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "subscribedToNotifications", func() []string {
				signal.(func() []string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "subscribedToNotifications", f)
		}
	}
}

func (ptr *QSqlDriver) DisconnectSubscribedToNotifications() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "subscribedToNotifications")
	}
}

func (ptr *QSqlDriver) SubscribedToNotifications() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QSqlDriver_SubscribedToNotifications(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QSqlDriver) SubscribedToNotificationsDefault() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QSqlDriver_SubscribedToNotificationsDefault(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

//export callbackQSqlDriver_Tables
func callbackQSqlDriver_Tables(ptr unsafe.Pointer, tableType C.longlong) C.struct_QtSql_PackedString {
	if signal := qt.GetSignal(ptr, "tables"); signal != nil {
		tempVal := signal.(func(QSql__TableType) []string)(QSql__TableType(tableType))
		return C.struct_QtSql_PackedString{data: C.CString(strings.Join(tempVal, "|")), len: C.longlong(len(strings.Join(tempVal, "|")))}
	}
	tempVal := NewQSqlDriverFromPointer(ptr).TablesDefault(QSql__TableType(tableType))
	return C.struct_QtSql_PackedString{data: C.CString(strings.Join(tempVal, "|")), len: C.longlong(len(strings.Join(tempVal, "|")))}
}

func (ptr *QSqlDriver) ConnectTables(f func(tableType QSql__TableType) []string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "tables"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "tables", func(tableType QSql__TableType) []string {
				signal.(func(QSql__TableType) []string)(tableType)
				return f(tableType)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "tables", f)
		}
	}
}

func (ptr *QSqlDriver) DisconnectTables() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "tables")
	}
}

func (ptr *QSqlDriver) Tables(tableType QSql__TableType) []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QSqlDriver_Tables(ptr.Pointer(), C.longlong(tableType))), "|")
	}
	return make([]string, 0)
}

func (ptr *QSqlDriver) TablesDefault(tableType QSql__TableType) []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QSqlDriver_TablesDefault(ptr.Pointer(), C.longlong(tableType))), "|")
	}
	return make([]string, 0)
}

//export callbackQSqlDriver_Handle
func callbackQSqlDriver_Handle(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "handle"); signal != nil {
		return core.PointerFromQVariant(signal.(func() *core.QVariant)())
	}

	return core.PointerFromQVariant(NewQSqlDriverFromPointer(ptr).HandleDefault())
}

func (ptr *QSqlDriver) ConnectHandle(f func() *core.QVariant) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "handle"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "handle", func() *core.QVariant {
				signal.(func() *core.QVariant)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "handle", f)
		}
	}
}

func (ptr *QSqlDriver) DisconnectHandle() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "handle")
	}
}

func (ptr *QSqlDriver) Handle() *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QSqlDriver_Handle(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlDriver) HandleDefault() *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QSqlDriver_HandleDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQSqlDriver_HasFeature
func callbackQSqlDriver_HasFeature(ptr unsafe.Pointer, feature C.longlong) C.char {
	if signal := qt.GetSignal(ptr, "hasFeature"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(QSqlDriver__DriverFeature) bool)(QSqlDriver__DriverFeature(feature)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QSqlDriver) ConnectHasFeature(f func(feature QSqlDriver__DriverFeature) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "hasFeature"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "hasFeature", func(feature QSqlDriver__DriverFeature) bool {
				signal.(func(QSqlDriver__DriverFeature) bool)(feature)
				return f(feature)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "hasFeature", f)
		}
	}
}

func (ptr *QSqlDriver) DisconnectHasFeature() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "hasFeature")
	}
}

func (ptr *QSqlDriver) HasFeature(feature QSqlDriver__DriverFeature) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlDriver_HasFeature(ptr.Pointer(), C.longlong(feature))) != 0
	}
	return false
}

//export callbackQSqlDriver_IsIdentifierEscaped
func callbackQSqlDriver_IsIdentifierEscaped(ptr unsafe.Pointer, identifier C.struct_QtSql_PackedString, ty C.longlong) C.char {
	if signal := qt.GetSignal(ptr, "isIdentifierEscaped"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string, QSqlDriver__IdentifierType) bool)(cGoUnpackString(identifier), QSqlDriver__IdentifierType(ty)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlDriverFromPointer(ptr).IsIdentifierEscapedDefault(cGoUnpackString(identifier), QSqlDriver__IdentifierType(ty)))))
}

func (ptr *QSqlDriver) ConnectIsIdentifierEscaped(f func(identifier string, ty QSqlDriver__IdentifierType) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "isIdentifierEscaped"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "isIdentifierEscaped", func(identifier string, ty QSqlDriver__IdentifierType) bool {
				signal.(func(string, QSqlDriver__IdentifierType) bool)(identifier, ty)
				return f(identifier, ty)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "isIdentifierEscaped", f)
		}
	}
}

func (ptr *QSqlDriver) DisconnectIsIdentifierEscaped() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "isIdentifierEscaped")
	}
}

func (ptr *QSqlDriver) IsIdentifierEscaped(identifier string, ty QSqlDriver__IdentifierType) bool {
	if ptr.Pointer() != nil {
		var identifierC *C.char
		if identifier != "" {
			identifierC = C.CString(identifier)
			defer C.free(unsafe.Pointer(identifierC))
		}
		return int8(C.QSqlDriver_IsIdentifierEscaped(ptr.Pointer(), C.struct_QtSql_PackedString{data: identifierC, len: C.longlong(len(identifier))}, C.longlong(ty))) != 0
	}
	return false
}

func (ptr *QSqlDriver) IsIdentifierEscapedDefault(identifier string, ty QSqlDriver__IdentifierType) bool {
	if ptr.Pointer() != nil {
		var identifierC *C.char
		if identifier != "" {
			identifierC = C.CString(identifier)
			defer C.free(unsafe.Pointer(identifierC))
		}
		return int8(C.QSqlDriver_IsIdentifierEscapedDefault(ptr.Pointer(), C.struct_QtSql_PackedString{data: identifierC, len: C.longlong(len(identifier))}, C.longlong(ty))) != 0
	}
	return false
}

//export callbackQSqlDriver_IsOpen
func callbackQSqlDriver_IsOpen(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "isOpen"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlDriverFromPointer(ptr).IsOpenDefault())))
}

func (ptr *QSqlDriver) ConnectIsOpen(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "isOpen"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "isOpen", func() bool {
				signal.(func() bool)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "isOpen", f)
		}
	}
}

func (ptr *QSqlDriver) DisconnectIsOpen() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "isOpen")
	}
}

func (ptr *QSqlDriver) IsOpen() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlDriver_IsOpen(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSqlDriver) IsOpenDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlDriver_IsOpenDefault(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSqlDriver) IsOpenError() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlDriver_IsOpenError(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQSqlDriver_MetaObject
func callbackQSqlDriver_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQSqlDriverFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QSqlDriver) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSqlDriver_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSqlDriver) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QSqlDriver___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlDriver) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlDriver___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QSqlDriver) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QSqlDriver___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QSqlDriver) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QSqlDriver___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSqlDriver) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlDriver___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QSqlDriver) __findChildren_newList2() unsafe.Pointer {
	return C.QSqlDriver___findChildren_newList2(ptr.Pointer())
}

func (ptr *QSqlDriver) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QSqlDriver___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSqlDriver) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlDriver___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QSqlDriver) __findChildren_newList3() unsafe.Pointer {
	return C.QSqlDriver___findChildren_newList3(ptr.Pointer())
}

func (ptr *QSqlDriver) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QSqlDriver___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSqlDriver) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlDriver___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QSqlDriver) __findChildren_newList() unsafe.Pointer {
	return C.QSqlDriver___findChildren_newList(ptr.Pointer())
}

func (ptr *QSqlDriver) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QSqlDriver___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSqlDriver) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlDriver___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QSqlDriver) __children_newList() unsafe.Pointer {
	return C.QSqlDriver___children_newList(ptr.Pointer())
}

//export callbackQSqlDriver_Event
func callbackQSqlDriver_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlDriverFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QSqlDriver) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlDriver_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQSqlDriver_EventFilter
func callbackQSqlDriver_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlDriverFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QSqlDriver) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlDriver_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQSqlDriver_ChildEvent
func callbackQSqlDriver_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSqlDriverFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSqlDriver) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlDriver_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQSqlDriver_ConnectNotify
func callbackQSqlDriver_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSqlDriverFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSqlDriver) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlDriver_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSqlDriver_CustomEvent
func callbackQSqlDriver_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSqlDriverFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSqlDriver) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlDriver_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQSqlDriver_DeleteLater
func callbackQSqlDriver_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQSqlDriverFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QSqlDriver) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QSqlDriver_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQSqlDriver_Destroyed
func callbackQSqlDriver_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQSqlDriver_DisconnectNotify
func callbackQSqlDriver_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSqlDriverFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSqlDriver) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlDriver_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSqlDriver_ObjectNameChanged
func callbackQSqlDriver_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtSql_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQSqlDriver_TimerEvent
func callbackQSqlDriver_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSqlDriverFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSqlDriver) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlDriver_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
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

func NewQSqlDriverCreatorFromPointer(ptr unsafe.Pointer) (n *QSqlDriverCreator) {
	n = new(QSqlDriverCreator)
	n.SetPointer(ptr)
	return
}

func (ptr *QSqlDriverCreator) DestroyQSqlDriverCreator() {
	if ptr != nil {
		C.free(ptr.Pointer())
		qt.DisconnectAllSignals(ptr.Pointer(), "")
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

type QSqlDriverCreatorBase struct {
	ptr unsafe.Pointer
}

type QSqlDriverCreatorBase_ITF interface {
	QSqlDriverCreatorBase_PTR() *QSqlDriverCreatorBase
}

func (ptr *QSqlDriverCreatorBase) QSqlDriverCreatorBase_PTR() *QSqlDriverCreatorBase {
	return ptr
}

func (ptr *QSqlDriverCreatorBase) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QSqlDriverCreatorBase) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQSqlDriverCreatorBase(ptr QSqlDriverCreatorBase_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlDriverCreatorBase_PTR().Pointer()
	}
	return nil
}

func NewQSqlDriverCreatorBaseFromPointer(ptr unsafe.Pointer) (n *QSqlDriverCreatorBase) {
	n = new(QSqlDriverCreatorBase)
	n.SetPointer(ptr)
	return
}

//export callbackQSqlDriverCreatorBase_DestroyQSqlDriverCreatorBase
func callbackQSqlDriverCreatorBase_DestroyQSqlDriverCreatorBase(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QSqlDriverCreatorBase"); signal != nil {
		signal.(func())()
	} else {
		NewQSqlDriverCreatorBaseFromPointer(ptr).DestroyQSqlDriverCreatorBaseDefault()
	}
}

func (ptr *QSqlDriverCreatorBase) ConnectDestroyQSqlDriverCreatorBase(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QSqlDriverCreatorBase"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QSqlDriverCreatorBase", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QSqlDriverCreatorBase", f)
		}
	}
}

func (ptr *QSqlDriverCreatorBase) DisconnectDestroyQSqlDriverCreatorBase() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QSqlDriverCreatorBase")
	}
}

func (ptr *QSqlDriverCreatorBase) DestroyQSqlDriverCreatorBase() {
	if ptr.Pointer() != nil {
		C.QSqlDriverCreatorBase_DestroyQSqlDriverCreatorBase(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSqlDriverCreatorBase) DestroyQSqlDriverCreatorBaseDefault() {
	if ptr.Pointer() != nil {
		C.QSqlDriverCreatorBase_DestroyQSqlDriverCreatorBaseDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQSqlDriverCreatorBase_CreateObject
func callbackQSqlDriverCreatorBase_CreateObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "createObject"); signal != nil {
		return PointerFromQSqlDriver(signal.(func() *QSqlDriver)())
	}

	return PointerFromQSqlDriver(NewQSqlDriver(nil))
}

func (ptr *QSqlDriverCreatorBase) ConnectCreateObject(f func() *QSqlDriver) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "createObject"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "createObject", func() *QSqlDriver {
				signal.(func() *QSqlDriver)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "createObject", f)
		}
	}
}

func (ptr *QSqlDriverCreatorBase) DisconnectCreateObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "createObject")
	}
}

func (ptr *QSqlDriverCreatorBase) CreateObject() *QSqlDriver {
	if ptr.Pointer() != nil {
		tmpValue := NewQSqlDriverFromPointer(C.QSqlDriverCreatorBase_CreateObject(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
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

func NewQSqlDriverPluginFromPointer(ptr unsafe.Pointer) (n *QSqlDriverPlugin) {
	n = new(QSqlDriverPlugin)
	n.SetPointer(ptr)
	return
}

//export callbackQSqlDriverPlugin_Create
func callbackQSqlDriverPlugin_Create(ptr unsafe.Pointer, key C.struct_QtSql_PackedString) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "create"); signal != nil {
		return PointerFromQSqlDriver(signal.(func(string) *QSqlDriver)(cGoUnpackString(key)))
	}

	return PointerFromQSqlDriver(NewQSqlDriver(nil))
}

func (ptr *QSqlDriverPlugin) ConnectCreate(f func(key string) *QSqlDriver) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "create"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "create", func(key string) *QSqlDriver {
				signal.(func(string) *QSqlDriver)(key)
				return f(key)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "create", f)
		}
	}
}

func (ptr *QSqlDriverPlugin) DisconnectCreate() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "create")
	}
}

func (ptr *QSqlDriverPlugin) Create(key string) *QSqlDriver {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		tmpValue := NewQSqlDriverFromPointer(C.QSqlDriverPlugin_Create(ptr.Pointer(), C.struct_QtSql_PackedString{data: keyC, len: C.longlong(len(key))}))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func NewQSqlDriverPlugin(parent core.QObject_ITF) *QSqlDriverPlugin {
	tmpValue := NewQSqlDriverPluginFromPointer(C.QSqlDriverPlugin_NewQSqlDriverPlugin(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func QSqlDriverPlugin_Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QSqlDriverPlugin_QSqlDriverPlugin_Tr(sC, cC, C.int(int32(n))))
}

func (ptr *QSqlDriverPlugin) Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QSqlDriverPlugin_QSqlDriverPlugin_Tr(sC, cC, C.int(int32(n))))
}

func QSqlDriverPlugin_TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QSqlDriverPlugin_QSqlDriverPlugin_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QSqlDriverPlugin) TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QSqlDriverPlugin_QSqlDriverPlugin_TrUtf8(sC, cC, C.int(int32(n))))
}

//export callbackQSqlDriverPlugin_DestroyQSqlDriverPlugin
func callbackQSqlDriverPlugin_DestroyQSqlDriverPlugin(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QSqlDriverPlugin"); signal != nil {
		signal.(func())()
	} else {
		NewQSqlDriverPluginFromPointer(ptr).DestroyQSqlDriverPluginDefault()
	}
}

func (ptr *QSqlDriverPlugin) ConnectDestroyQSqlDriverPlugin(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QSqlDriverPlugin"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QSqlDriverPlugin", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QSqlDriverPlugin", f)
		}
	}
}

func (ptr *QSqlDriverPlugin) DisconnectDestroyQSqlDriverPlugin() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QSqlDriverPlugin")
	}
}

func (ptr *QSqlDriverPlugin) DestroyQSqlDriverPlugin() {
	if ptr.Pointer() != nil {
		C.QSqlDriverPlugin_DestroyQSqlDriverPlugin(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QSqlDriverPlugin) DestroyQSqlDriverPluginDefault() {
	if ptr.Pointer() != nil {
		C.QSqlDriverPlugin_DestroyQSqlDriverPluginDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQSqlDriverPlugin_MetaObject
func callbackQSqlDriverPlugin_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQSqlDriverPluginFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QSqlDriverPlugin) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSqlDriverPlugin_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSqlDriverPlugin) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QSqlDriverPlugin___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlDriverPlugin) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlDriverPlugin___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QSqlDriverPlugin) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QSqlDriverPlugin___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QSqlDriverPlugin) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QSqlDriverPlugin___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSqlDriverPlugin) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlDriverPlugin___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QSqlDriverPlugin) __findChildren_newList2() unsafe.Pointer {
	return C.QSqlDriverPlugin___findChildren_newList2(ptr.Pointer())
}

func (ptr *QSqlDriverPlugin) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QSqlDriverPlugin___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSqlDriverPlugin) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlDriverPlugin___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QSqlDriverPlugin) __findChildren_newList3() unsafe.Pointer {
	return C.QSqlDriverPlugin___findChildren_newList3(ptr.Pointer())
}

func (ptr *QSqlDriverPlugin) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QSqlDriverPlugin___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSqlDriverPlugin) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlDriverPlugin___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QSqlDriverPlugin) __findChildren_newList() unsafe.Pointer {
	return C.QSqlDriverPlugin___findChildren_newList(ptr.Pointer())
}

func (ptr *QSqlDriverPlugin) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QSqlDriverPlugin___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSqlDriverPlugin) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlDriverPlugin___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QSqlDriverPlugin) __children_newList() unsafe.Pointer {
	return C.QSqlDriverPlugin___children_newList(ptr.Pointer())
}

//export callbackQSqlDriverPlugin_Event
func callbackQSqlDriverPlugin_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlDriverPluginFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QSqlDriverPlugin) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlDriverPlugin_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQSqlDriverPlugin_EventFilter
func callbackQSqlDriverPlugin_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlDriverPluginFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QSqlDriverPlugin) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlDriverPlugin_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQSqlDriverPlugin_ChildEvent
func callbackQSqlDriverPlugin_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSqlDriverPluginFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSqlDriverPlugin) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlDriverPlugin_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQSqlDriverPlugin_ConnectNotify
func callbackQSqlDriverPlugin_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSqlDriverPluginFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSqlDriverPlugin) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlDriverPlugin_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSqlDriverPlugin_CustomEvent
func callbackQSqlDriverPlugin_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSqlDriverPluginFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSqlDriverPlugin) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlDriverPlugin_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQSqlDriverPlugin_DeleteLater
func callbackQSqlDriverPlugin_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQSqlDriverPluginFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QSqlDriverPlugin) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QSqlDriverPlugin_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQSqlDriverPlugin_Destroyed
func callbackQSqlDriverPlugin_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQSqlDriverPlugin_DisconnectNotify
func callbackQSqlDriverPlugin_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSqlDriverPluginFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSqlDriverPlugin) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlDriverPlugin_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSqlDriverPlugin_ObjectNameChanged
func callbackQSqlDriverPlugin_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtSql_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQSqlDriverPlugin_TimerEvent
func callbackQSqlDriverPlugin_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSqlDriverPluginFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSqlDriverPlugin) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlDriverPlugin_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QSqlError struct {
	ptr unsafe.Pointer
}

type QSqlError_ITF interface {
	QSqlError_PTR() *QSqlError
}

func (ptr *QSqlError) QSqlError_PTR() *QSqlError {
	return ptr
}

func (ptr *QSqlError) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QSqlError) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQSqlError(ptr QSqlError_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlError_PTR().Pointer()
	}
	return nil
}

func NewQSqlErrorFromPointer(ptr unsafe.Pointer) (n *QSqlError) {
	n = new(QSqlError)
	n.SetPointer(ptr)
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

func NewQSqlError4(other QSqlError_ITF) *QSqlError {
	tmpValue := NewQSqlErrorFromPointer(C.QSqlError_NewQSqlError4(PointerFromQSqlError(other)))
	runtime.SetFinalizer(tmpValue, (*QSqlError).DestroyQSqlError)
	return tmpValue
}

func NewQSqlError3(other QSqlError_ITF) *QSqlError {
	tmpValue := NewQSqlErrorFromPointer(C.QSqlError_NewQSqlError3(PointerFromQSqlError(other)))
	runtime.SetFinalizer(tmpValue, (*QSqlError).DestroyQSqlError)
	return tmpValue
}

func NewQSqlError(driverText string, databaseText string, ty QSqlError__ErrorType, code string) *QSqlError {
	var driverTextC *C.char
	if driverText != "" {
		driverTextC = C.CString(driverText)
		defer C.free(unsafe.Pointer(driverTextC))
	}
	var databaseTextC *C.char
	if databaseText != "" {
		databaseTextC = C.CString(databaseText)
		defer C.free(unsafe.Pointer(databaseTextC))
	}
	var codeC *C.char
	if code != "" {
		codeC = C.CString(code)
		defer C.free(unsafe.Pointer(codeC))
	}
	tmpValue := NewQSqlErrorFromPointer(C.QSqlError_NewQSqlError(C.struct_QtSql_PackedString{data: driverTextC, len: C.longlong(len(driverText))}, C.struct_QtSql_PackedString{data: databaseTextC, len: C.longlong(len(databaseText))}, C.longlong(ty), C.struct_QtSql_PackedString{data: codeC, len: C.longlong(len(code))}))
	runtime.SetFinalizer(tmpValue, (*QSqlError).DestroyQSqlError)
	return tmpValue
}

func (ptr *QSqlError) Swap(other QSqlError_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlError_Swap(ptr.Pointer(), PointerFromQSqlError(other))
	}
}

func (ptr *QSqlError) DestroyQSqlError() {
	if ptr.Pointer() != nil {
		C.QSqlError_DestroyQSqlError(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QSqlError) Type() QSqlError__ErrorType {
	if ptr.Pointer() != nil {
		return QSqlError__ErrorType(C.QSqlError_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSqlError) DatabaseText() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSqlError_DatabaseText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlError) DriverText() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSqlError_DriverText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlError) NativeErrorCode() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSqlError_NativeErrorCode(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlError) Text() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSqlError_Text(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlError) IsValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlError_IsValid(ptr.Pointer())) != 0
	}
	return false
}

type QSqlField struct {
	ptr unsafe.Pointer
}

type QSqlField_ITF interface {
	QSqlField_PTR() *QSqlField
}

func (ptr *QSqlField) QSqlField_PTR() *QSqlField {
	return ptr
}

func (ptr *QSqlField) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QSqlField) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQSqlField(ptr QSqlField_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlField_PTR().Pointer()
	}
	return nil
}

func NewQSqlFieldFromPointer(ptr unsafe.Pointer) (n *QSqlField) {
	n = new(QSqlField)
	n.SetPointer(ptr)
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

func NewQSqlField3(other QSqlField_ITF) *QSqlField {
	tmpValue := NewQSqlFieldFromPointer(C.QSqlField_NewQSqlField3(PointerFromQSqlField(other)))
	runtime.SetFinalizer(tmpValue, (*QSqlField).DestroyQSqlField)
	return tmpValue
}

func NewQSqlField(fieldName string, ty core.QVariant__Type) *QSqlField {
	var fieldNameC *C.char
	if fieldName != "" {
		fieldNameC = C.CString(fieldName)
		defer C.free(unsafe.Pointer(fieldNameC))
	}
	tmpValue := NewQSqlFieldFromPointer(C.QSqlField_NewQSqlField(C.struct_QtSql_PackedString{data: fieldNameC, len: C.longlong(len(fieldName))}, C.longlong(ty)))
	runtime.SetFinalizer(tmpValue, (*QSqlField).DestroyQSqlField)
	return tmpValue
}

func NewQSqlField2(fieldName string, ty core.QVariant__Type, table string) *QSqlField {
	var fieldNameC *C.char
	if fieldName != "" {
		fieldNameC = C.CString(fieldName)
		defer C.free(unsafe.Pointer(fieldNameC))
	}
	var tableC *C.char
	if table != "" {
		tableC = C.CString(table)
		defer C.free(unsafe.Pointer(tableC))
	}
	tmpValue := NewQSqlFieldFromPointer(C.QSqlField_NewQSqlField2(C.struct_QtSql_PackedString{data: fieldNameC, len: C.longlong(len(fieldName))}, C.longlong(ty), C.struct_QtSql_PackedString{data: tableC, len: C.longlong(len(table))}))
	runtime.SetFinalizer(tmpValue, (*QSqlField).DestroyQSqlField)
	return tmpValue
}

func (ptr *QSqlField) Clear() {
	if ptr.Pointer() != nil {
		C.QSqlField_Clear(ptr.Pointer())
	}
}

func (ptr *QSqlField) SetAutoValue(autoVal bool) {
	if ptr.Pointer() != nil {
		C.QSqlField_SetAutoValue(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(autoVal))))
	}
}

func (ptr *QSqlField) SetDefaultValue(value core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlField_SetDefaultValue(ptr.Pointer(), core.PointerFromQVariant(value))
	}
}

func (ptr *QSqlField) SetGenerated(gen bool) {
	if ptr.Pointer() != nil {
		C.QSqlField_SetGenerated(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(gen))))
	}
}

func (ptr *QSqlField) SetLength(fieldLength int) {
	if ptr.Pointer() != nil {
		C.QSqlField_SetLength(ptr.Pointer(), C.int(int32(fieldLength)))
	}
}

func (ptr *QSqlField) SetName(name string) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		C.QSqlField_SetName(ptr.Pointer(), C.struct_QtSql_PackedString{data: nameC, len: C.longlong(len(name))})
	}
}

func (ptr *QSqlField) SetPrecision(precision int) {
	if ptr.Pointer() != nil {
		C.QSqlField_SetPrecision(ptr.Pointer(), C.int(int32(precision)))
	}
}

func (ptr *QSqlField) SetReadOnly(readOnly bool) {
	if ptr.Pointer() != nil {
		C.QSqlField_SetReadOnly(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(readOnly))))
	}
}

func (ptr *QSqlField) SetRequired(required bool) {
	if ptr.Pointer() != nil {
		C.QSqlField_SetRequired(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(required))))
	}
}

func (ptr *QSqlField) SetRequiredStatus(required QSqlField__RequiredStatus) {
	if ptr.Pointer() != nil {
		C.QSqlField_SetRequiredStatus(ptr.Pointer(), C.longlong(required))
	}
}

func (ptr *QSqlField) SetTableName(table string) {
	if ptr.Pointer() != nil {
		var tableC *C.char
		if table != "" {
			tableC = C.CString(table)
			defer C.free(unsafe.Pointer(tableC))
		}
		C.QSqlField_SetTableName(ptr.Pointer(), C.struct_QtSql_PackedString{data: tableC, len: C.longlong(len(table))})
	}
}

func (ptr *QSqlField) SetType(ty core.QVariant__Type) {
	if ptr.Pointer() != nil {
		C.QSqlField_SetType(ptr.Pointer(), C.longlong(ty))
	}
}

func (ptr *QSqlField) SetValue(value core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlField_SetValue(ptr.Pointer(), core.PointerFromQVariant(value))
	}
}

func (ptr *QSqlField) DestroyQSqlField() {
	if ptr.Pointer() != nil {
		C.QSqlField_DestroyQSqlField(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QSqlField) RequiredStatus() QSqlField__RequiredStatus {
	if ptr.Pointer() != nil {
		return QSqlField__RequiredStatus(C.QSqlField_RequiredStatus(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSqlField) Name() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSqlField_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlField) TableName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSqlField_TableName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlField) DefaultValue() *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QSqlField_DefaultValue(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlField) Value() *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QSqlField_Value(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlField) Type() core.QVariant__Type {
	if ptr.Pointer() != nil {
		return core.QVariant__Type(C.QSqlField_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSqlField) IsAutoValue() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlField_IsAutoValue(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSqlField) IsGenerated() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlField_IsGenerated(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSqlField) IsNull() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlField_IsNull(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSqlField) IsReadOnly() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlField_IsReadOnly(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSqlField) IsValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlField_IsValid(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSqlField) Length() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSqlField_Length(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSqlField) Precision() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSqlField_Precision(ptr.Pointer())))
	}
	return 0
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

func NewQSqlIndexFromPointer(ptr unsafe.Pointer) (n *QSqlIndex) {
	n = new(QSqlIndex)
	n.SetPointer(ptr)
	return
}
func NewQSqlIndex2(other QSqlIndex_ITF) *QSqlIndex {
	tmpValue := NewQSqlIndexFromPointer(C.QSqlIndex_NewQSqlIndex2(PointerFromQSqlIndex(other)))
	runtime.SetFinalizer(tmpValue, (*QSqlIndex).DestroyQSqlIndex)
	return tmpValue
}

func NewQSqlIndex(cursorname string, name string) *QSqlIndex {
	var cursornameC *C.char
	if cursorname != "" {
		cursornameC = C.CString(cursorname)
		defer C.free(unsafe.Pointer(cursornameC))
	}
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	tmpValue := NewQSqlIndexFromPointer(C.QSqlIndex_NewQSqlIndex(C.struct_QtSql_PackedString{data: cursornameC, len: C.longlong(len(cursorname))}, C.struct_QtSql_PackedString{data: nameC, len: C.longlong(len(name))}))
	runtime.SetFinalizer(tmpValue, (*QSqlIndex).DestroyQSqlIndex)
	return tmpValue
}

func (ptr *QSqlIndex) Append2(field QSqlField_ITF, desc bool) {
	if ptr.Pointer() != nil {
		C.QSqlIndex_Append2(ptr.Pointer(), PointerFromQSqlField(field), C.char(int8(qt.GoBoolToInt(desc))))
	}
}

func (ptr *QSqlIndex) SetCursorName(cursorName string) {
	if ptr.Pointer() != nil {
		var cursorNameC *C.char
		if cursorName != "" {
			cursorNameC = C.CString(cursorName)
			defer C.free(unsafe.Pointer(cursorNameC))
		}
		C.QSqlIndex_SetCursorName(ptr.Pointer(), C.struct_QtSql_PackedString{data: cursorNameC, len: C.longlong(len(cursorName))})
	}
}

func (ptr *QSqlIndex) SetDescending(i int, desc bool) {
	if ptr.Pointer() != nil {
		C.QSqlIndex_SetDescending(ptr.Pointer(), C.int(int32(i)), C.char(int8(qt.GoBoolToInt(desc))))
	}
}

func (ptr *QSqlIndex) SetName(name string) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		C.QSqlIndex_SetName(ptr.Pointer(), C.struct_QtSql_PackedString{data: nameC, len: C.longlong(len(name))})
	}
}

func (ptr *QSqlIndex) DestroyQSqlIndex() {
	if ptr.Pointer() != nil {
		C.QSqlIndex_DestroyQSqlIndex(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QSqlIndex) CursorName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSqlIndex_CursorName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlIndex) Name() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSqlIndex_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlIndex) IsDescending(i int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlIndex_IsDescending(ptr.Pointer(), C.int(int32(i)))) != 0
	}
	return false
}

func (ptr *QSqlIndex) __sorts_atList(i int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlIndex___sorts_atList(ptr.Pointer(), C.int(int32(i)))) != 0
	}
	return false
}

func (ptr *QSqlIndex) __sorts_setList(i bool) {
	if ptr.Pointer() != nil {
		C.QSqlIndex___sorts_setList(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(i))))
	}
}

func (ptr *QSqlIndex) __sorts_newList() unsafe.Pointer {
	return C.QSqlIndex___sorts_newList(ptr.Pointer())
}

func (ptr *QSqlIndex) __setSorts__atList(i int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlIndex___setSorts__atList(ptr.Pointer(), C.int(int32(i)))) != 0
	}
	return false
}

func (ptr *QSqlIndex) __setSorts__setList(i bool) {
	if ptr.Pointer() != nil {
		C.QSqlIndex___setSorts__setList(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(i))))
	}
}

func (ptr *QSqlIndex) __setSorts__newList() unsafe.Pointer {
	return C.QSqlIndex___setSorts__newList(ptr.Pointer())
}

type QSqlQuery struct {
	ptr unsafe.Pointer
}

type QSqlQuery_ITF interface {
	QSqlQuery_PTR() *QSqlQuery
}

func (ptr *QSqlQuery) QSqlQuery_PTR() *QSqlQuery {
	return ptr
}

func (ptr *QSqlQuery) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QSqlQuery) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQSqlQuery(ptr QSqlQuery_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlQuery_PTR().Pointer()
	}
	return nil
}

func NewQSqlQueryFromPointer(ptr unsafe.Pointer) (n *QSqlQuery) {
	n = new(QSqlQuery)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QSqlQuery__BatchExecutionMode
//QSqlQuery::BatchExecutionMode
type QSqlQuery__BatchExecutionMode int64

const (
	QSqlQuery__ValuesAsRows    QSqlQuery__BatchExecutionMode = QSqlQuery__BatchExecutionMode(0)
	QSqlQuery__ValuesAsColumns QSqlQuery__BatchExecutionMode = QSqlQuery__BatchExecutionMode(1)
)

func NewQSqlQuery3(db QSqlDatabase_ITF) *QSqlQuery {
	tmpValue := NewQSqlQueryFromPointer(C.QSqlQuery_NewQSqlQuery3(PointerFromQSqlDatabase(db)))
	runtime.SetFinalizer(tmpValue, (*QSqlQuery).DestroyQSqlQuery)
	return tmpValue
}

func NewQSqlQuery(result QSqlResult_ITF) *QSqlQuery {
	tmpValue := NewQSqlQueryFromPointer(C.QSqlQuery_NewQSqlQuery(PointerFromQSqlResult(result)))
	runtime.SetFinalizer(tmpValue, (*QSqlQuery).DestroyQSqlQuery)
	return tmpValue
}

func NewQSqlQuery4(other QSqlQuery_ITF) *QSqlQuery {
	tmpValue := NewQSqlQueryFromPointer(C.QSqlQuery_NewQSqlQuery4(PointerFromQSqlQuery(other)))
	runtime.SetFinalizer(tmpValue, (*QSqlQuery).DestroyQSqlQuery)
	return tmpValue
}

func NewQSqlQuery2(query string, db QSqlDatabase_ITF) *QSqlQuery {
	var queryC *C.char
	if query != "" {
		queryC = C.CString(query)
		defer C.free(unsafe.Pointer(queryC))
	}
	tmpValue := NewQSqlQueryFromPointer(C.QSqlQuery_NewQSqlQuery2(C.struct_QtSql_PackedString{data: queryC, len: C.longlong(len(query))}, PointerFromQSqlDatabase(db)))
	runtime.SetFinalizer(tmpValue, (*QSqlQuery).DestroyQSqlQuery)
	return tmpValue
}

func (ptr *QSqlQuery) Exec2() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlQuery_Exec2(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSqlQuery) Exec(query string) bool {
	if ptr.Pointer() != nil {
		var queryC *C.char
		if query != "" {
			queryC = C.CString(query)
			defer C.free(unsafe.Pointer(queryC))
		}
		return int8(C.QSqlQuery_Exec(ptr.Pointer(), C.struct_QtSql_PackedString{data: queryC, len: C.longlong(len(query))})) != 0
	}
	return false
}

func (ptr *QSqlQuery) ExecBatch(mode QSqlQuery__BatchExecutionMode) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlQuery_ExecBatch(ptr.Pointer(), C.longlong(mode))) != 0
	}
	return false
}

func (ptr *QSqlQuery) First() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlQuery_First(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSqlQuery) Last() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlQuery_Last(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSqlQuery) Next() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlQuery_Next(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSqlQuery) NextResult() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlQuery_NextResult(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSqlQuery) Prepare(query string) bool {
	if ptr.Pointer() != nil {
		var queryC *C.char
		if query != "" {
			queryC = C.CString(query)
			defer C.free(unsafe.Pointer(queryC))
		}
		return int8(C.QSqlQuery_Prepare(ptr.Pointer(), C.struct_QtSql_PackedString{data: queryC, len: C.longlong(len(query))})) != 0
	}
	return false
}

func (ptr *QSqlQuery) Previous() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlQuery_Previous(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSqlQuery) Seek(index int, relative bool) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlQuery_Seek(ptr.Pointer(), C.int(int32(index)), C.char(int8(qt.GoBoolToInt(relative))))) != 0
	}
	return false
}

func (ptr *QSqlQuery) AddBindValue(val core.QVariant_ITF, paramType QSql__ParamTypeFlag) {
	if ptr.Pointer() != nil {
		C.QSqlQuery_AddBindValue(ptr.Pointer(), core.PointerFromQVariant(val), C.longlong(paramType))
	}
}

func (ptr *QSqlQuery) BindValue(placeholder string, val core.QVariant_ITF, paramType QSql__ParamTypeFlag) {
	if ptr.Pointer() != nil {
		var placeholderC *C.char
		if placeholder != "" {
			placeholderC = C.CString(placeholder)
			defer C.free(unsafe.Pointer(placeholderC))
		}
		C.QSqlQuery_BindValue(ptr.Pointer(), C.struct_QtSql_PackedString{data: placeholderC, len: C.longlong(len(placeholder))}, core.PointerFromQVariant(val), C.longlong(paramType))
	}
}

func (ptr *QSqlQuery) BindValue2(pos int, val core.QVariant_ITF, paramType QSql__ParamTypeFlag) {
	if ptr.Pointer() != nil {
		C.QSqlQuery_BindValue2(ptr.Pointer(), C.int(int32(pos)), core.PointerFromQVariant(val), C.longlong(paramType))
	}
}

func (ptr *QSqlQuery) Clear() {
	if ptr.Pointer() != nil {
		C.QSqlQuery_Clear(ptr.Pointer())
	}
}

func (ptr *QSqlQuery) Finish() {
	if ptr.Pointer() != nil {
		C.QSqlQuery_Finish(ptr.Pointer())
	}
}

func (ptr *QSqlQuery) SetForwardOnly(forward bool) {
	if ptr.Pointer() != nil {
		C.QSqlQuery_SetForwardOnly(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(forward))))
	}
}

func (ptr *QSqlQuery) SetNumericalPrecisionPolicy(precisionPolicy QSql__NumericalPrecisionPolicy) {
	if ptr.Pointer() != nil {
		C.QSqlQuery_SetNumericalPrecisionPolicy(ptr.Pointer(), C.longlong(precisionPolicy))
	}
}

func (ptr *QSqlQuery) DestroyQSqlQuery() {
	if ptr.Pointer() != nil {
		C.QSqlQuery_DestroyQSqlQuery(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QSqlQuery) BoundValues() map[string]*core.QVariant {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtSql_PackedList) map[string]*core.QVariant {
			out := make(map[string]*core.QVariant, int(l.len))
			tmpList := NewQSqlQueryFromPointer(l.data)
			for i, v := range tmpList.__boundValues_keyList() {
				out[v] = tmpList.__boundValues_atList(v, i)
			}
			return out
		}(C.QSqlQuery_BoundValues(ptr.Pointer()))
	}
	return make(map[string]*core.QVariant, 0)
}

func (ptr *QSqlQuery) NumericalPrecisionPolicy() QSql__NumericalPrecisionPolicy {
	if ptr.Pointer() != nil {
		return QSql__NumericalPrecisionPolicy(C.QSqlQuery_NumericalPrecisionPolicy(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSqlQuery) LastError() *QSqlError {
	if ptr.Pointer() != nil {
		tmpValue := NewQSqlErrorFromPointer(C.QSqlQuery_LastError(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QSqlError).DestroyQSqlError)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlQuery) Record() *QSqlRecord {
	if ptr.Pointer() != nil {
		tmpValue := NewQSqlRecordFromPointer(C.QSqlQuery_Record(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QSqlRecord).DestroyQSqlRecord)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlQuery) ExecutedQuery() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSqlQuery_ExecutedQuery(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlQuery) LastQuery() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSqlQuery_LastQuery(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlQuery) BoundValue(placeholder string) *core.QVariant {
	if ptr.Pointer() != nil {
		var placeholderC *C.char
		if placeholder != "" {
			placeholderC = C.CString(placeholder)
			defer C.free(unsafe.Pointer(placeholderC))
		}
		tmpValue := core.NewQVariantFromPointer(C.QSqlQuery_BoundValue(ptr.Pointer(), C.struct_QtSql_PackedString{data: placeholderC, len: C.longlong(len(placeholder))}))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlQuery) BoundValue2(pos int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QSqlQuery_BoundValue2(ptr.Pointer(), C.int(int32(pos))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlQuery) LastInsertId() *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QSqlQuery_LastInsertId(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlQuery) Value2(name string) *core.QVariant {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		tmpValue := core.NewQVariantFromPointer(C.QSqlQuery_Value2(ptr.Pointer(), C.struct_QtSql_PackedString{data: nameC, len: C.longlong(len(name))}))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlQuery) Value(index int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QSqlQuery_Value(ptr.Pointer(), C.int(int32(index))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlQuery) IsActive() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlQuery_IsActive(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSqlQuery) IsForwardOnly() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlQuery_IsForwardOnly(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSqlQuery) IsNull2(name string) bool {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		return int8(C.QSqlQuery_IsNull2(ptr.Pointer(), C.struct_QtSql_PackedString{data: nameC, len: C.longlong(len(name))})) != 0
	}
	return false
}

func (ptr *QSqlQuery) IsNull(field int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlQuery_IsNull(ptr.Pointer(), C.int(int32(field)))) != 0
	}
	return false
}

func (ptr *QSqlQuery) IsSelect() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlQuery_IsSelect(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSqlQuery) IsValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlQuery_IsValid(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSqlQuery) Driver() *QSqlDriver {
	if ptr.Pointer() != nil {
		tmpValue := NewQSqlDriverFromPointer(C.QSqlQuery_Driver(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSqlQuery) Result() *QSqlResult {
	if ptr.Pointer() != nil {
		return NewQSqlResultFromPointer(C.QSqlQuery_Result(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSqlQuery) At() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSqlQuery_At(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSqlQuery) NumRowsAffected() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSqlQuery_NumRowsAffected(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSqlQuery) Size() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSqlQuery_Size(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSqlQuery) __boundValues_atList(v string, i int) *core.QVariant {
	if ptr.Pointer() != nil {
		var vC *C.char
		if v != "" {
			vC = C.CString(v)
			defer C.free(unsafe.Pointer(vC))
		}
		tmpValue := core.NewQVariantFromPointer(C.QSqlQuery___boundValues_atList(ptr.Pointer(), C.struct_QtSql_PackedString{data: vC, len: C.longlong(len(v))}, C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlQuery) __boundValues_setList(key string, i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		C.QSqlQuery___boundValues_setList(ptr.Pointer(), C.struct_QtSql_PackedString{data: keyC, len: C.longlong(len(key))}, core.PointerFromQVariant(i))
	}
}

func (ptr *QSqlQuery) __boundValues_newList() unsafe.Pointer {
	return C.QSqlQuery___boundValues_newList(ptr.Pointer())
}

func (ptr *QSqlQuery) __boundValues_keyList() []string {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtSql_PackedList) []string {
			out := make([]string, int(l.len))
			tmpList := NewQSqlQueryFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____boundValues_keyList_atList(i)
			}
			return out
		}(C.QSqlQuery___boundValues_keyList(ptr.Pointer()))
	}
	return make([]string, 0)
}

func (ptr *QSqlQuery) ____boundValues_keyList_atList(i int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSqlQuery_____boundValues_keyList_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return ""
}

func (ptr *QSqlQuery) ____boundValues_keyList_setList(i string) {
	if ptr.Pointer() != nil {
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		C.QSqlQuery_____boundValues_keyList_setList(ptr.Pointer(), C.struct_QtSql_PackedString{data: iC, len: C.longlong(len(i))})
	}
}

func (ptr *QSqlQuery) ____boundValues_keyList_newList() unsafe.Pointer {
	return C.QSqlQuery_____boundValues_keyList_newList(ptr.Pointer())
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

func NewQSqlQueryModelFromPointer(ptr unsafe.Pointer) (n *QSqlQueryModel) {
	n = new(QSqlQueryModel)
	n.SetPointer(ptr)
	return
}
func NewQSqlQueryModel(parent core.QObject_ITF) *QSqlQueryModel {
	tmpValue := NewQSqlQueryModelFromPointer(C.QSqlQueryModel_NewQSqlQueryModel(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func QSqlQueryModel_Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QSqlQueryModel_QSqlQueryModel_Tr(sC, cC, C.int(int32(n))))
}

func (ptr *QSqlQueryModel) Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QSqlQueryModel_QSqlQueryModel_Tr(sC, cC, C.int(int32(n))))
}

func QSqlQueryModel_TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QSqlQueryModel_QSqlQueryModel_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QSqlQueryModel) TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QSqlQueryModel_QSqlQueryModel_TrUtf8(sC, cC, C.int(int32(n))))
}

//export callbackQSqlQueryModel_InsertColumns
func callbackQSqlQueryModel_InsertColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "insertColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *core.QModelIndex) bool)(int(int32(column)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlQueryModelFromPointer(ptr).InsertColumnsDefault(int(int32(column)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QSqlQueryModel) InsertColumnsDefault(column int, count int, parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlQueryModel_InsertColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackQSqlQueryModel_RemoveColumns
func callbackQSqlQueryModel_RemoveColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "removeColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *core.QModelIndex) bool)(int(int32(column)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlQueryModelFromPointer(ptr).RemoveColumnsDefault(int(int32(column)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QSqlQueryModel) RemoveColumnsDefault(column int, count int, parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlQueryModel_RemoveColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackQSqlQueryModel_SetHeaderData
func callbackQSqlQueryModel_SetHeaderData(ptr unsafe.Pointer, section C.int, orientation C.longlong, value unsafe.Pointer, role C.int) C.char {
	if signal := qt.GetSignal(ptr, "setHeaderData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, core.Qt__Orientation, *core.QVariant, int) bool)(int(int32(section)), core.Qt__Orientation(orientation), core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlQueryModelFromPointer(ptr).SetHeaderDataDefault(int(int32(section)), core.Qt__Orientation(orientation), core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *QSqlQueryModel) SetHeaderDataDefault(section int, orientation core.Qt__Orientation, value core.QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlQueryModel_SetHeaderDataDefault(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), core.PointerFromQVariant(value), C.int(int32(role)))) != 0
	}
	return false
}

//export callbackQSqlQueryModel_Clear
func callbackQSqlQueryModel_Clear(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "clear"); signal != nil {
		signal.(func())()
	} else {
		NewQSqlQueryModelFromPointer(ptr).ClearDefault()
	}
}

func (ptr *QSqlQueryModel) ConnectClear(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "clear"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "clear", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "clear", f)
		}
	}
}

func (ptr *QSqlQueryModel) DisconnectClear() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "clear")
	}
}

func (ptr *QSqlQueryModel) Clear() {
	if ptr.Pointer() != nil {
		C.QSqlQueryModel_Clear(ptr.Pointer())
	}
}

func (ptr *QSqlQueryModel) ClearDefault() {
	if ptr.Pointer() != nil {
		C.QSqlQueryModel_ClearDefault(ptr.Pointer())
	}
}

//export callbackQSqlQueryModel_FetchMore
func callbackQSqlQueryModel_FetchMore(ptr unsafe.Pointer, parent unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "fetchMore"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(parent))
	} else {
		NewQSqlQueryModelFromPointer(ptr).FetchMoreDefault(core.NewQModelIndexFromPointer(parent))
	}
}

func (ptr *QSqlQueryModel) FetchMoreDefault(parent core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlQueryModel_FetchMoreDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent))
	}
}

//export callbackQSqlQueryModel_QueryChange
func callbackQSqlQueryModel_QueryChange(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "queryChange"); signal != nil {
		signal.(func())()
	} else {
		NewQSqlQueryModelFromPointer(ptr).QueryChangeDefault()
	}
}

func (ptr *QSqlQueryModel) ConnectQueryChange(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "queryChange"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "queryChange", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "queryChange", f)
		}
	}
}

func (ptr *QSqlQueryModel) DisconnectQueryChange() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "queryChange")
	}
}

func (ptr *QSqlQueryModel) QueryChange() {
	if ptr.Pointer() != nil {
		C.QSqlQueryModel_QueryChange(ptr.Pointer())
	}
}

func (ptr *QSqlQueryModel) QueryChangeDefault() {
	if ptr.Pointer() != nil {
		C.QSqlQueryModel_QueryChangeDefault(ptr.Pointer())
	}
}

func (ptr *QSqlQueryModel) SetLastError(error QSqlError_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlQueryModel_SetLastError(ptr.Pointer(), PointerFromQSqlError(error))
	}
}

func (ptr *QSqlQueryModel) SetQuery(query QSqlQuery_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlQueryModel_SetQuery(ptr.Pointer(), PointerFromQSqlQuery(query))
	}
}

func (ptr *QSqlQueryModel) SetQuery2(query string, db QSqlDatabase_ITF) {
	if ptr.Pointer() != nil {
		var queryC *C.char
		if query != "" {
			queryC = C.CString(query)
			defer C.free(unsafe.Pointer(queryC))
		}
		C.QSqlQueryModel_SetQuery2(ptr.Pointer(), C.struct_QtSql_PackedString{data: queryC, len: C.longlong(len(query))}, PointerFromQSqlDatabase(db))
	}
}

//export callbackQSqlQueryModel_DestroyQSqlQueryModel
func callbackQSqlQueryModel_DestroyQSqlQueryModel(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QSqlQueryModel"); signal != nil {
		signal.(func())()
	} else {
		NewQSqlQueryModelFromPointer(ptr).DestroyQSqlQueryModelDefault()
	}
}

func (ptr *QSqlQueryModel) ConnectDestroyQSqlQueryModel(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QSqlQueryModel"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QSqlQueryModel", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QSqlQueryModel", f)
		}
	}
}

func (ptr *QSqlQueryModel) DisconnectDestroyQSqlQueryModel() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QSqlQueryModel")
	}
}

func (ptr *QSqlQueryModel) DestroyQSqlQueryModel() {
	if ptr.Pointer() != nil {
		C.QSqlQueryModel_DestroyQSqlQueryModel(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QSqlQueryModel) DestroyQSqlQueryModelDefault() {
	if ptr.Pointer() != nil {
		C.QSqlQueryModel_DestroyQSqlQueryModelDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQSqlQueryModel_RoleNames
func callbackQSqlQueryModel_RoleNames(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "roleNames"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewQSqlQueryModelFromPointer(NewQSqlQueryModelFromPointer(nil).__roleNames_newList())
			for k, v := range signal.(func() map[int]*core.QByteArray)() {
				tmpList.__roleNames_setList(k, v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewQSqlQueryModelFromPointer(NewQSqlQueryModelFromPointer(nil).__roleNames_newList())
		for k, v := range NewQSqlQueryModelFromPointer(ptr).RoleNamesDefault() {
			tmpList.__roleNames_setList(k, v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *QSqlQueryModel) RoleNamesDefault() map[int]*core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtSql_PackedList) map[int]*core.QByteArray {
			out := make(map[int]*core.QByteArray, int(l.len))
			tmpList := NewQSqlQueryModelFromPointer(l.data)
			for i, v := range tmpList.__roleNames_keyList() {
				out[v] = tmpList.__roleNames_atList(v, i)
			}
			return out
		}(C.QSqlQueryModel_RoleNamesDefault(ptr.Pointer()))
	}
	return make(map[int]*core.QByteArray, 0)
}

//export callbackQSqlQueryModel_IndexInQuery
func callbackQSqlQueryModel_IndexInQuery(ptr unsafe.Pointer, item unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "indexInQuery"); signal != nil {
		return core.PointerFromQModelIndex(signal.(func(*core.QModelIndex) *core.QModelIndex)(core.NewQModelIndexFromPointer(item)))
	}

	return core.PointerFromQModelIndex(NewQSqlQueryModelFromPointer(ptr).IndexInQueryDefault(core.NewQModelIndexFromPointer(item)))
}

func (ptr *QSqlQueryModel) ConnectIndexInQuery(f func(item *core.QModelIndex) *core.QModelIndex) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "indexInQuery"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "indexInQuery", func(item *core.QModelIndex) *core.QModelIndex {
				signal.(func(*core.QModelIndex) *core.QModelIndex)(item)
				return f(item)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "indexInQuery", f)
		}
	}
}

func (ptr *QSqlQueryModel) DisconnectIndexInQuery() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "indexInQuery")
	}
}

func (ptr *QSqlQueryModel) IndexInQuery(item core.QModelIndex_ITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QSqlQueryModel_IndexInQuery(ptr.Pointer(), core.PointerFromQModelIndex(item)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlQueryModel) IndexInQueryDefault(item core.QModelIndex_ITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QSqlQueryModel_IndexInQueryDefault(ptr.Pointer(), core.PointerFromQModelIndex(item)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlQueryModel) LastError() *QSqlError {
	if ptr.Pointer() != nil {
		tmpValue := NewQSqlErrorFromPointer(C.QSqlQueryModel_LastError(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QSqlError).DestroyQSqlError)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlQueryModel) Query() *QSqlQuery {
	if ptr.Pointer() != nil {
		tmpValue := NewQSqlQueryFromPointer(C.QSqlQueryModel_Query(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QSqlQuery).DestroyQSqlQuery)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlQueryModel) Record2() *QSqlRecord {
	if ptr.Pointer() != nil {
		tmpValue := NewQSqlRecordFromPointer(C.QSqlQueryModel_Record2(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QSqlRecord).DestroyQSqlRecord)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlQueryModel) Record(row int) *QSqlRecord {
	if ptr.Pointer() != nil {
		tmpValue := NewQSqlRecordFromPointer(C.QSqlQueryModel_Record(ptr.Pointer(), C.int(int32(row))))
		runtime.SetFinalizer(tmpValue, (*QSqlRecord).DestroyQSqlRecord)
		return tmpValue
	}
	return nil
}

//export callbackQSqlQueryModel_Data
func callbackQSqlQueryModel_Data(ptr unsafe.Pointer, item unsafe.Pointer, role C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "data"); signal != nil {
		return core.PointerFromQVariant(signal.(func(*core.QModelIndex, int) *core.QVariant)(core.NewQModelIndexFromPointer(item), int(int32(role))))
	}

	return core.PointerFromQVariant(NewQSqlQueryModelFromPointer(ptr).DataDefault(core.NewQModelIndexFromPointer(item), int(int32(role))))
}

func (ptr *QSqlQueryModel) ConnectData(f func(item *core.QModelIndex, role int) *core.QVariant) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "data"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "data", func(item *core.QModelIndex, role int) *core.QVariant {
				signal.(func(*core.QModelIndex, int) *core.QVariant)(item, role)
				return f(item, role)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "data", f)
		}
	}
}

func (ptr *QSqlQueryModel) DisconnectData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "data")
	}
}

func (ptr *QSqlQueryModel) Data(item core.QModelIndex_ITF, role int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QSqlQueryModel_Data(ptr.Pointer(), core.PointerFromQModelIndex(item), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlQueryModel) DataDefault(item core.QModelIndex_ITF, role int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QSqlQueryModel_DataDefault(ptr.Pointer(), core.PointerFromQModelIndex(item), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQSqlQueryModel_HeaderData
func callbackQSqlQueryModel_HeaderData(ptr unsafe.Pointer, section C.int, orientation C.longlong, role C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "headerData"); signal != nil {
		return core.PointerFromQVariant(signal.(func(int, core.Qt__Orientation, int) *core.QVariant)(int(int32(section)), core.Qt__Orientation(orientation), int(int32(role))))
	}

	return core.PointerFromQVariant(NewQSqlQueryModelFromPointer(ptr).HeaderDataDefault(int(int32(section)), core.Qt__Orientation(orientation), int(int32(role))))
}

func (ptr *QSqlQueryModel) HeaderDataDefault(section int, orientation core.Qt__Orientation, role int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QSqlQueryModel_HeaderDataDefault(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQSqlQueryModel_CanFetchMore
func callbackQSqlQueryModel_CanFetchMore(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "canFetchMore"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QModelIndex) bool)(core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlQueryModelFromPointer(ptr).CanFetchMoreDefault(core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QSqlQueryModel) CanFetchMoreDefault(parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlQueryModel_CanFetchMoreDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackQSqlQueryModel_MetaObject
func callbackQSqlQueryModel_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQSqlQueryModelFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QSqlQueryModel) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSqlQueryModel_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQSqlQueryModel_ColumnCount
func callbackQSqlQueryModel_ColumnCount(ptr unsafe.Pointer, index unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "columnCount"); signal != nil {
		return C.int(int32(signal.(func(*core.QModelIndex) int)(core.NewQModelIndexFromPointer(index))))
	}

	return C.int(int32(NewQSqlQueryModelFromPointer(ptr).ColumnCountDefault(core.NewQModelIndexFromPointer(index))))
}

func (ptr *QSqlQueryModel) ConnectColumnCount(f func(index *core.QModelIndex) int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "columnCount"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "columnCount", func(index *core.QModelIndex) int {
				signal.(func(*core.QModelIndex) int)(index)
				return f(index)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "columnCount", f)
		}
	}
}

func (ptr *QSqlQueryModel) DisconnectColumnCount() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "columnCount")
	}
}

func (ptr *QSqlQueryModel) ColumnCount(index core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSqlQueryModel_ColumnCount(ptr.Pointer(), core.PointerFromQModelIndex(index))))
	}
	return 0
}

func (ptr *QSqlQueryModel) ColumnCountDefault(index core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSqlQueryModel_ColumnCountDefault(ptr.Pointer(), core.PointerFromQModelIndex(index))))
	}
	return 0
}

//export callbackQSqlQueryModel_RowCount
func callbackQSqlQueryModel_RowCount(ptr unsafe.Pointer, parent unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "rowCount"); signal != nil {
		return C.int(int32(signal.(func(*core.QModelIndex) int)(core.NewQModelIndexFromPointer(parent))))
	}

	return C.int(int32(NewQSqlQueryModelFromPointer(ptr).RowCountDefault(core.NewQModelIndexFromPointer(parent))))
}

func (ptr *QSqlQueryModel) ConnectRowCount(f func(parent *core.QModelIndex) int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "rowCount"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "rowCount", func(parent *core.QModelIndex) int {
				signal.(func(*core.QModelIndex) int)(parent)
				return f(parent)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "rowCount", f)
		}
	}
}

func (ptr *QSqlQueryModel) DisconnectRowCount() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "rowCount")
	}
}

func (ptr *QSqlQueryModel) RowCount(parent core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSqlQueryModel_RowCount(ptr.Pointer(), core.PointerFromQModelIndex(parent))))
	}
	return 0
}

func (ptr *QSqlQueryModel) RowCountDefault(parent core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSqlQueryModel_RowCountDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent))))
	}
	return 0
}

func (ptr *QSqlQueryModel) __roleNames_atList(v int, i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QSqlQueryModel___roleNames_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlQueryModel) __roleNames_setList(key int, i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlQueryModel___roleNames_setList(ptr.Pointer(), C.int(int32(key)), core.PointerFromQByteArray(i))
	}
}

func (ptr *QSqlQueryModel) __roleNames_newList() unsafe.Pointer {
	return C.QSqlQueryModel___roleNames_newList(ptr.Pointer())
}

func (ptr *QSqlQueryModel) __roleNames_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtSql_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewQSqlQueryModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____roleNames_keyList_atList(i)
			}
			return out
		}(C.QSqlQueryModel___roleNames_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *QSqlQueryModel) ____roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSqlQueryModel_____roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *QSqlQueryModel) ____roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.QSqlQueryModel_____roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *QSqlQueryModel) ____roleNames_keyList_newList() unsafe.Pointer {
	return C.QSqlQueryModel_____roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *QSqlQueryModel) ____setItemData_roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSqlQueryModel_____setItemData_roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *QSqlQueryModel) ____setItemData_roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.QSqlQueryModel_____setItemData_roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *QSqlQueryModel) ____setItemData_roles_keyList_newList() unsafe.Pointer {
	return C.QSqlQueryModel_____setItemData_roles_keyList_newList(ptr.Pointer())
}

func (ptr *QSqlQueryModel) ____itemData_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSqlQueryModel_____itemData_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *QSqlQueryModel) ____itemData_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.QSqlQueryModel_____itemData_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *QSqlQueryModel) ____itemData_keyList_newList() unsafe.Pointer {
	return C.QSqlQueryModel_____itemData_keyList_newList(ptr.Pointer())
}

func (ptr *QSqlQueryModel) __setItemData_roles_atList(v int, i int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QSqlQueryModel___setItemData_roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlQueryModel) __setItemData_roles_setList(key int, i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlQueryModel___setItemData_roles_setList(ptr.Pointer(), C.int(int32(key)), core.PointerFromQVariant(i))
	}
}

func (ptr *QSqlQueryModel) __setItemData_roles_newList() unsafe.Pointer {
	return C.QSqlQueryModel___setItemData_roles_newList(ptr.Pointer())
}

func (ptr *QSqlQueryModel) __setItemData_roles_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtSql_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewQSqlQueryModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____setItemData_roles_keyList_atList(i)
			}
			return out
		}(C.QSqlQueryModel___setItemData_roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *QSqlQueryModel) __changePersistentIndexList_from_atList(i int) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QSqlQueryModel___changePersistentIndexList_from_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlQueryModel) __changePersistentIndexList_from_setList(i core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlQueryModel___changePersistentIndexList_from_setList(ptr.Pointer(), core.PointerFromQModelIndex(i))
	}
}

func (ptr *QSqlQueryModel) __changePersistentIndexList_from_newList() unsafe.Pointer {
	return C.QSqlQueryModel___changePersistentIndexList_from_newList(ptr.Pointer())
}

func (ptr *QSqlQueryModel) __changePersistentIndexList_to_atList(i int) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QSqlQueryModel___changePersistentIndexList_to_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlQueryModel) __changePersistentIndexList_to_setList(i core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlQueryModel___changePersistentIndexList_to_setList(ptr.Pointer(), core.PointerFromQModelIndex(i))
	}
}

func (ptr *QSqlQueryModel) __changePersistentIndexList_to_newList() unsafe.Pointer {
	return C.QSqlQueryModel___changePersistentIndexList_to_newList(ptr.Pointer())
}

func (ptr *QSqlQueryModel) __dataChanged_roles_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSqlQueryModel___dataChanged_roles_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *QSqlQueryModel) __dataChanged_roles_setList(i int) {
	if ptr.Pointer() != nil {
		C.QSqlQueryModel___dataChanged_roles_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *QSqlQueryModel) __dataChanged_roles_newList() unsafe.Pointer {
	return C.QSqlQueryModel___dataChanged_roles_newList(ptr.Pointer())
}

func (ptr *QSqlQueryModel) __layoutAboutToBeChanged_parents_atList(i int) *core.QPersistentModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQPersistentModelIndexFromPointer(C.QSqlQueryModel___layoutAboutToBeChanged_parents_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QPersistentModelIndex).DestroyQPersistentModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlQueryModel) __layoutAboutToBeChanged_parents_setList(i core.QPersistentModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlQueryModel___layoutAboutToBeChanged_parents_setList(ptr.Pointer(), core.PointerFromQPersistentModelIndex(i))
	}
}

func (ptr *QSqlQueryModel) __layoutAboutToBeChanged_parents_newList() unsafe.Pointer {
	return C.QSqlQueryModel___layoutAboutToBeChanged_parents_newList(ptr.Pointer())
}

func (ptr *QSqlQueryModel) __layoutChanged_parents_atList(i int) *core.QPersistentModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQPersistentModelIndexFromPointer(C.QSqlQueryModel___layoutChanged_parents_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QPersistentModelIndex).DestroyQPersistentModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlQueryModel) __layoutChanged_parents_setList(i core.QPersistentModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlQueryModel___layoutChanged_parents_setList(ptr.Pointer(), core.PointerFromQPersistentModelIndex(i))
	}
}

func (ptr *QSqlQueryModel) __layoutChanged_parents_newList() unsafe.Pointer {
	return C.QSqlQueryModel___layoutChanged_parents_newList(ptr.Pointer())
}

func (ptr *QSqlQueryModel) __itemData_atList(v int, i int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QSqlQueryModel___itemData_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlQueryModel) __itemData_setList(key int, i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlQueryModel___itemData_setList(ptr.Pointer(), C.int(int32(key)), core.PointerFromQVariant(i))
	}
}

func (ptr *QSqlQueryModel) __itemData_newList() unsafe.Pointer {
	return C.QSqlQueryModel___itemData_newList(ptr.Pointer())
}

func (ptr *QSqlQueryModel) __itemData_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtSql_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewQSqlQueryModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____itemData_keyList_atList(i)
			}
			return out
		}(C.QSqlQueryModel___itemData_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *QSqlQueryModel) __mimeData_indexes_atList(i int) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QSqlQueryModel___mimeData_indexes_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlQueryModel) __mimeData_indexes_setList(i core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlQueryModel___mimeData_indexes_setList(ptr.Pointer(), core.PointerFromQModelIndex(i))
	}
}

func (ptr *QSqlQueryModel) __mimeData_indexes_newList() unsafe.Pointer {
	return C.QSqlQueryModel___mimeData_indexes_newList(ptr.Pointer())
}

func (ptr *QSqlQueryModel) __match_atList(i int) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QSqlQueryModel___match_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlQueryModel) __match_setList(i core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlQueryModel___match_setList(ptr.Pointer(), core.PointerFromQModelIndex(i))
	}
}

func (ptr *QSqlQueryModel) __match_newList() unsafe.Pointer {
	return C.QSqlQueryModel___match_newList(ptr.Pointer())
}

func (ptr *QSqlQueryModel) __persistentIndexList_atList(i int) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QSqlQueryModel___persistentIndexList_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlQueryModel) __persistentIndexList_setList(i core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlQueryModel___persistentIndexList_setList(ptr.Pointer(), core.PointerFromQModelIndex(i))
	}
}

func (ptr *QSqlQueryModel) __persistentIndexList_newList() unsafe.Pointer {
	return C.QSqlQueryModel___persistentIndexList_newList(ptr.Pointer())
}

func (ptr *QSqlQueryModel) ____doSetRoleNames_roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSqlQueryModel_____doSetRoleNames_roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *QSqlQueryModel) ____doSetRoleNames_roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.QSqlQueryModel_____doSetRoleNames_roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *QSqlQueryModel) ____doSetRoleNames_roleNames_keyList_newList() unsafe.Pointer {
	return C.QSqlQueryModel_____doSetRoleNames_roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *QSqlQueryModel) ____setRoleNames_roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSqlQueryModel_____setRoleNames_roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *QSqlQueryModel) ____setRoleNames_roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.QSqlQueryModel_____setRoleNames_roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *QSqlQueryModel) ____setRoleNames_roleNames_keyList_newList() unsafe.Pointer {
	return C.QSqlQueryModel_____setRoleNames_roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *QSqlQueryModel) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QSqlQueryModel___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlQueryModel) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlQueryModel___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QSqlQueryModel) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QSqlQueryModel___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QSqlQueryModel) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QSqlQueryModel___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSqlQueryModel) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlQueryModel___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QSqlQueryModel) __findChildren_newList2() unsafe.Pointer {
	return C.QSqlQueryModel___findChildren_newList2(ptr.Pointer())
}

func (ptr *QSqlQueryModel) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QSqlQueryModel___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSqlQueryModel) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlQueryModel___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QSqlQueryModel) __findChildren_newList3() unsafe.Pointer {
	return C.QSqlQueryModel___findChildren_newList3(ptr.Pointer())
}

func (ptr *QSqlQueryModel) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QSqlQueryModel___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSqlQueryModel) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlQueryModel___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QSqlQueryModel) __findChildren_newList() unsafe.Pointer {
	return C.QSqlQueryModel___findChildren_newList(ptr.Pointer())
}

func (ptr *QSqlQueryModel) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QSqlQueryModel___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSqlQueryModel) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlQueryModel___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QSqlQueryModel) __children_newList() unsafe.Pointer {
	return C.QSqlQueryModel___children_newList(ptr.Pointer())
}

//export callbackQSqlQueryModel_DropMimeData
func callbackQSqlQueryModel_DropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "dropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QMimeData, core.Qt__DropAction, int, int, *core.QModelIndex) bool)(core.NewQMimeDataFromPointer(data), core.Qt__DropAction(action), int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlQueryModelFromPointer(ptr).DropMimeDataDefault(core.NewQMimeDataFromPointer(data), core.Qt__DropAction(action), int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QSqlQueryModel) DropMimeDataDefault(data core.QMimeData_ITF, action core.Qt__DropAction, row int, column int, parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlQueryModel_DropMimeDataDefault(ptr.Pointer(), core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackQSqlQueryModel_Index
func callbackQSqlQueryModel_Index(ptr unsafe.Pointer, row C.int, column C.int, parent unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "index"); signal != nil {
		return core.PointerFromQModelIndex(signal.(func(int, int, *core.QModelIndex) *core.QModelIndex)(int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(parent)))
	}

	return core.PointerFromQModelIndex(NewQSqlQueryModelFromPointer(ptr).IndexDefault(int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(parent)))
}

func (ptr *QSqlQueryModel) IndexDefault(row int, column int, parent core.QModelIndex_ITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QSqlQueryModel_IndexDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(parent)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQSqlQueryModel_Sibling
func callbackQSqlQueryModel_Sibling(ptr unsafe.Pointer, row C.int, column C.int, idx unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sibling"); signal != nil {
		return core.PointerFromQModelIndex(signal.(func(int, int, *core.QModelIndex) *core.QModelIndex)(int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(idx)))
	}

	return core.PointerFromQModelIndex(NewQSqlQueryModelFromPointer(ptr).SiblingDefault(int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(idx)))
}

func (ptr *QSqlQueryModel) SiblingDefault(row int, column int, idx core.QModelIndex_ITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QSqlQueryModel_SiblingDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(idx)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQSqlQueryModel_Flags
func callbackQSqlQueryModel_Flags(ptr unsafe.Pointer, index unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "flags"); signal != nil {
		return C.longlong(signal.(func(*core.QModelIndex) core.Qt__ItemFlag)(core.NewQModelIndexFromPointer(index)))
	}

	return C.longlong(NewQSqlQueryModelFromPointer(ptr).FlagsDefault(core.NewQModelIndexFromPointer(index)))
}

func (ptr *QSqlQueryModel) FlagsDefault(index core.QModelIndex_ITF) core.Qt__ItemFlag {
	if ptr.Pointer() != nil {
		return core.Qt__ItemFlag(C.QSqlQueryModel_FlagsDefault(ptr.Pointer(), core.PointerFromQModelIndex(index)))
	}
	return 0
}

//export callbackQSqlQueryModel_InsertRows
func callbackQSqlQueryModel_InsertRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "insertRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *core.QModelIndex) bool)(int(int32(row)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlQueryModelFromPointer(ptr).InsertRowsDefault(int(int32(row)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QSqlQueryModel) InsertRowsDefault(row int, count int, parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlQueryModel_InsertRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackQSqlQueryModel_MoveColumns
func callbackQSqlQueryModel_MoveColumns(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceColumn C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	if signal := qt.GetSignal(ptr, "moveColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QModelIndex, int, int, *core.QModelIndex, int) bool)(core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlQueryModelFromPointer(ptr).MoveColumnsDefault(core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *QSqlQueryModel) MoveColumnsDefault(sourceParent core.QModelIndex_ITF, sourceColumn int, count int, destinationParent core.QModelIndex_ITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlQueryModel_MoveColumnsDefault(ptr.Pointer(), core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceColumn)), C.int(int32(count)), core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild)))) != 0
	}
	return false
}

//export callbackQSqlQueryModel_MoveRows
func callbackQSqlQueryModel_MoveRows(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceRow C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	if signal := qt.GetSignal(ptr, "moveRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QModelIndex, int, int, *core.QModelIndex, int) bool)(core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlQueryModelFromPointer(ptr).MoveRowsDefault(core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *QSqlQueryModel) MoveRowsDefault(sourceParent core.QModelIndex_ITF, sourceRow int, count int, destinationParent core.QModelIndex_ITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlQueryModel_MoveRowsDefault(ptr.Pointer(), core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceRow)), C.int(int32(count)), core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild)))) != 0
	}
	return false
}

//export callbackQSqlQueryModel_RemoveRows
func callbackQSqlQueryModel_RemoveRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "removeRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *core.QModelIndex) bool)(int(int32(row)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlQueryModelFromPointer(ptr).RemoveRowsDefault(int(int32(row)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QSqlQueryModel) RemoveRowsDefault(row int, count int, parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlQueryModel_RemoveRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackQSqlQueryModel_SetData
func callbackQSqlQueryModel_SetData(ptr unsafe.Pointer, index unsafe.Pointer, value unsafe.Pointer, role C.int) C.char {
	if signal := qt.GetSignal(ptr, "setData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QModelIndex, *core.QVariant, int) bool)(core.NewQModelIndexFromPointer(index), core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlQueryModelFromPointer(ptr).SetDataDefault(core.NewQModelIndexFromPointer(index), core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *QSqlQueryModel) SetDataDefault(index core.QModelIndex_ITF, value core.QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlQueryModel_SetDataDefault(ptr.Pointer(), core.PointerFromQModelIndex(index), core.PointerFromQVariant(value), C.int(int32(role)))) != 0
	}
	return false
}

//export callbackQSqlQueryModel_SetItemData
func callbackQSqlQueryModel_SetItemData(ptr unsafe.Pointer, index unsafe.Pointer, roles C.struct_QtSql_PackedList) C.char {
	if signal := qt.GetSignal(ptr, "setItemData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QModelIndex, map[int]*core.QVariant) bool)(core.NewQModelIndexFromPointer(index), func(l C.struct_QtSql_PackedList) map[int]*core.QVariant {
			out := make(map[int]*core.QVariant, int(l.len))
			tmpList := NewQSqlQueryModelFromPointer(l.data)
			for i, v := range tmpList.__setItemData_roles_keyList() {
				out[v] = tmpList.__setItemData_roles_atList(v, i)
			}
			return out
		}(roles)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlQueryModelFromPointer(ptr).SetItemDataDefault(core.NewQModelIndexFromPointer(index), func(l C.struct_QtSql_PackedList) map[int]*core.QVariant {
		out := make(map[int]*core.QVariant, int(l.len))
		tmpList := NewQSqlQueryModelFromPointer(l.data)
		for i, v := range tmpList.__setItemData_roles_keyList() {
			out[v] = tmpList.__setItemData_roles_atList(v, i)
		}
		return out
	}(roles)))))
}

func (ptr *QSqlQueryModel) SetItemDataDefault(index core.QModelIndex_ITF, roles map[int]*core.QVariant) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlQueryModel_SetItemDataDefault(ptr.Pointer(), core.PointerFromQModelIndex(index), func() unsafe.Pointer {
			tmpList := NewQSqlQueryModelFromPointer(NewQSqlQueryModelFromPointer(nil).__setItemData_roles_newList())
			for k, v := range roles {
				tmpList.__setItemData_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())) != 0
	}
	return false
}

//export callbackQSqlQueryModel_Submit
func callbackQSqlQueryModel_Submit(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "submit"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlQueryModelFromPointer(ptr).SubmitDefault())))
}

func (ptr *QSqlQueryModel) SubmitDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlQueryModel_SubmitDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQSqlQueryModel_ColumnsAboutToBeInserted
func callbackQSqlQueryModel_ColumnsAboutToBeInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeInserted"); signal != nil {
		signal.(func(*core.QModelIndex, int, int))(core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackQSqlQueryModel_ColumnsAboutToBeMoved
func callbackQSqlQueryModel_ColumnsAboutToBeMoved(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceStart C.int, sourceEnd C.int, destinationParent unsafe.Pointer, destinationColumn C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeMoved"); signal != nil {
		signal.(func(*core.QModelIndex, int, int, *core.QModelIndex, int))(core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceStart)), int(int32(sourceEnd)), core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationColumn)))
	}

}

//export callbackQSqlQueryModel_ColumnsAboutToBeRemoved
func callbackQSqlQueryModel_ColumnsAboutToBeRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeRemoved"); signal != nil {
		signal.(func(*core.QModelIndex, int, int))(core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackQSqlQueryModel_ColumnsInserted
func callbackQSqlQueryModel_ColumnsInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsInserted"); signal != nil {
		signal.(func(*core.QModelIndex, int, int))(core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackQSqlQueryModel_ColumnsMoved
func callbackQSqlQueryModel_ColumnsMoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int, destination unsafe.Pointer, column C.int) {
	if signal := qt.GetSignal(ptr, "columnsMoved"); signal != nil {
		signal.(func(*core.QModelIndex, int, int, *core.QModelIndex, int))(core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)), core.NewQModelIndexFromPointer(destination), int(int32(column)))
	}

}

//export callbackQSqlQueryModel_ColumnsRemoved
func callbackQSqlQueryModel_ColumnsRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsRemoved"); signal != nil {
		signal.(func(*core.QModelIndex, int, int))(core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackQSqlQueryModel_DataChanged
func callbackQSqlQueryModel_DataChanged(ptr unsafe.Pointer, topLeft unsafe.Pointer, bottomRight unsafe.Pointer, roles C.struct_QtSql_PackedList) {
	if signal := qt.GetSignal(ptr, "dataChanged"); signal != nil {
		signal.(func(*core.QModelIndex, *core.QModelIndex, []int))(core.NewQModelIndexFromPointer(topLeft), core.NewQModelIndexFromPointer(bottomRight), func(l C.struct_QtSql_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewQSqlQueryModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__dataChanged_roles_atList(i)
			}
			return out
		}(roles))
	}

}

//export callbackQSqlQueryModel_HeaderDataChanged
func callbackQSqlQueryModel_HeaderDataChanged(ptr unsafe.Pointer, orientation C.longlong, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "headerDataChanged"); signal != nil {
		signal.(func(core.Qt__Orientation, int, int))(core.Qt__Orientation(orientation), int(int32(first)), int(int32(last)))
	}

}

//export callbackQSqlQueryModel_LayoutAboutToBeChanged
func callbackQSqlQueryModel_LayoutAboutToBeChanged(ptr unsafe.Pointer, parents C.struct_QtSql_PackedList, hint C.longlong) {
	if signal := qt.GetSignal(ptr, "layoutAboutToBeChanged"); signal != nil {
		signal.(func([]*core.QPersistentModelIndex, core.QAbstractItemModel__LayoutChangeHint))(func(l C.struct_QtSql_PackedList) []*core.QPersistentModelIndex {
			out := make([]*core.QPersistentModelIndex, int(l.len))
			tmpList := NewQSqlQueryModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__layoutAboutToBeChanged_parents_atList(i)
			}
			return out
		}(parents), core.QAbstractItemModel__LayoutChangeHint(hint))
	}

}

//export callbackQSqlQueryModel_LayoutChanged
func callbackQSqlQueryModel_LayoutChanged(ptr unsafe.Pointer, parents C.struct_QtSql_PackedList, hint C.longlong) {
	if signal := qt.GetSignal(ptr, "layoutChanged"); signal != nil {
		signal.(func([]*core.QPersistentModelIndex, core.QAbstractItemModel__LayoutChangeHint))(func(l C.struct_QtSql_PackedList) []*core.QPersistentModelIndex {
			out := make([]*core.QPersistentModelIndex, int(l.len))
			tmpList := NewQSqlQueryModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__layoutChanged_parents_atList(i)
			}
			return out
		}(parents), core.QAbstractItemModel__LayoutChangeHint(hint))
	}

}

//export callbackQSqlQueryModel_ModelAboutToBeReset
func callbackQSqlQueryModel_ModelAboutToBeReset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelAboutToBeReset"); signal != nil {
		signal.(func())()
	}

}

//export callbackQSqlQueryModel_ModelReset
func callbackQSqlQueryModel_ModelReset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelReset"); signal != nil {
		signal.(func())()
	}

}

//export callbackQSqlQueryModel_ResetInternalData
func callbackQSqlQueryModel_ResetInternalData(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resetInternalData"); signal != nil {
		signal.(func())()
	} else {
		NewQSqlQueryModelFromPointer(ptr).ResetInternalDataDefault()
	}
}

func (ptr *QSqlQueryModel) ResetInternalDataDefault() {
	if ptr.Pointer() != nil {
		C.QSqlQueryModel_ResetInternalDataDefault(ptr.Pointer())
	}
}

//export callbackQSqlQueryModel_Revert
func callbackQSqlQueryModel_Revert(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "revert"); signal != nil {
		signal.(func())()
	} else {
		NewQSqlQueryModelFromPointer(ptr).RevertDefault()
	}
}

func (ptr *QSqlQueryModel) RevertDefault() {
	if ptr.Pointer() != nil {
		C.QSqlQueryModel_RevertDefault(ptr.Pointer())
	}
}

//export callbackQSqlQueryModel_RowsAboutToBeInserted
func callbackQSqlQueryModel_RowsAboutToBeInserted(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeInserted"); signal != nil {
		signal.(func(*core.QModelIndex, int, int))(core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)))
	}

}

//export callbackQSqlQueryModel_RowsAboutToBeMoved
func callbackQSqlQueryModel_RowsAboutToBeMoved(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceStart C.int, sourceEnd C.int, destinationParent unsafe.Pointer, destinationRow C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeMoved"); signal != nil {
		signal.(func(*core.QModelIndex, int, int, *core.QModelIndex, int))(core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceStart)), int(int32(sourceEnd)), core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationRow)))
	}

}

//export callbackQSqlQueryModel_RowsAboutToBeRemoved
func callbackQSqlQueryModel_RowsAboutToBeRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeRemoved"); signal != nil {
		signal.(func(*core.QModelIndex, int, int))(core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackQSqlQueryModel_RowsInserted
func callbackQSqlQueryModel_RowsInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsInserted"); signal != nil {
		signal.(func(*core.QModelIndex, int, int))(core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackQSqlQueryModel_RowsMoved
func callbackQSqlQueryModel_RowsMoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int, destination unsafe.Pointer, row C.int) {
	if signal := qt.GetSignal(ptr, "rowsMoved"); signal != nil {
		signal.(func(*core.QModelIndex, int, int, *core.QModelIndex, int))(core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)), core.NewQModelIndexFromPointer(destination), int(int32(row)))
	}

}

//export callbackQSqlQueryModel_RowsRemoved
func callbackQSqlQueryModel_RowsRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsRemoved"); signal != nil {
		signal.(func(*core.QModelIndex, int, int))(core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackQSqlQueryModel_Sort
func callbackQSqlQueryModel_Sort(ptr unsafe.Pointer, column C.int, order C.longlong) {
	if signal := qt.GetSignal(ptr, "sort"); signal != nil {
		signal.(func(int, core.Qt__SortOrder))(int(int32(column)), core.Qt__SortOrder(order))
	} else {
		NewQSqlQueryModelFromPointer(ptr).SortDefault(int(int32(column)), core.Qt__SortOrder(order))
	}
}

func (ptr *QSqlQueryModel) SortDefault(column int, order core.Qt__SortOrder) {
	if ptr.Pointer() != nil {
		C.QSqlQueryModel_SortDefault(ptr.Pointer(), C.int(int32(column)), C.longlong(order))
	}
}

//export callbackQSqlQueryModel_ItemData
func callbackQSqlQueryModel_ItemData(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "itemData"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewQSqlQueryModelFromPointer(NewQSqlQueryModelFromPointer(nil).__itemData_newList())
			for k, v := range signal.(func(*core.QModelIndex) map[int]*core.QVariant)(core.NewQModelIndexFromPointer(index)) {
				tmpList.__itemData_setList(k, v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewQSqlQueryModelFromPointer(NewQSqlQueryModelFromPointer(nil).__itemData_newList())
		for k, v := range NewQSqlQueryModelFromPointer(ptr).ItemDataDefault(core.NewQModelIndexFromPointer(index)) {
			tmpList.__itemData_setList(k, v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *QSqlQueryModel) ItemDataDefault(index core.QModelIndex_ITF) map[int]*core.QVariant {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtSql_PackedList) map[int]*core.QVariant {
			out := make(map[int]*core.QVariant, int(l.len))
			tmpList := NewQSqlQueryModelFromPointer(l.data)
			for i, v := range tmpList.__itemData_keyList() {
				out[v] = tmpList.__itemData_atList(v, i)
			}
			return out
		}(C.QSqlQueryModel_ItemDataDefault(ptr.Pointer(), core.PointerFromQModelIndex(index)))
	}
	return make(map[int]*core.QVariant, 0)
}

//export callbackQSqlQueryModel_MimeData
func callbackQSqlQueryModel_MimeData(ptr unsafe.Pointer, indexes C.struct_QtSql_PackedList) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "mimeData"); signal != nil {
		return core.PointerFromQMimeData(signal.(func([]*core.QModelIndex) *core.QMimeData)(func(l C.struct_QtSql_PackedList) []*core.QModelIndex {
			out := make([]*core.QModelIndex, int(l.len))
			tmpList := NewQSqlQueryModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__mimeData_indexes_atList(i)
			}
			return out
		}(indexes)))
	}

	return core.PointerFromQMimeData(NewQSqlQueryModelFromPointer(ptr).MimeDataDefault(func(l C.struct_QtSql_PackedList) []*core.QModelIndex {
		out := make([]*core.QModelIndex, int(l.len))
		tmpList := NewQSqlQueryModelFromPointer(l.data)
		for i := 0; i < len(out); i++ {
			out[i] = tmpList.__mimeData_indexes_atList(i)
		}
		return out
	}(indexes)))
}

func (ptr *QSqlQueryModel) MimeDataDefault(indexes []*core.QModelIndex) *core.QMimeData {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQMimeDataFromPointer(C.QSqlQueryModel_MimeDataDefault(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQSqlQueryModelFromPointer(NewQSqlQueryModelFromPointer(nil).__mimeData_indexes_newList())
			for _, v := range indexes {
				tmpList.__mimeData_indexes_setList(v)
			}
			return tmpList.Pointer()
		}()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQSqlQueryModel_Buddy
func callbackQSqlQueryModel_Buddy(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "buddy"); signal != nil {
		return core.PointerFromQModelIndex(signal.(func(*core.QModelIndex) *core.QModelIndex)(core.NewQModelIndexFromPointer(index)))
	}

	return core.PointerFromQModelIndex(NewQSqlQueryModelFromPointer(ptr).BuddyDefault(core.NewQModelIndexFromPointer(index)))
}

func (ptr *QSqlQueryModel) BuddyDefault(index core.QModelIndex_ITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QSqlQueryModel_BuddyDefault(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQSqlQueryModel_Parent
func callbackQSqlQueryModel_Parent(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "parent"); signal != nil {
		return core.PointerFromQModelIndex(signal.(func(*core.QModelIndex) *core.QModelIndex)(core.NewQModelIndexFromPointer(index)))
	}

	return core.PointerFromQModelIndex(NewQSqlQueryModelFromPointer(ptr).ParentDefault(core.NewQModelIndexFromPointer(index)))
}

func (ptr *QSqlQueryModel) ParentDefault(index core.QModelIndex_ITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QSqlQueryModel_ParentDefault(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQSqlQueryModel_Match
func callbackQSqlQueryModel_Match(ptr unsafe.Pointer, start unsafe.Pointer, role C.int, value unsafe.Pointer, hits C.int, flags C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "match"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewQSqlQueryModelFromPointer(NewQSqlQueryModelFromPointer(nil).__match_newList())
			for _, v := range signal.(func(*core.QModelIndex, int, *core.QVariant, int, core.Qt__MatchFlag) []*core.QModelIndex)(core.NewQModelIndexFromPointer(start), int(int32(role)), core.NewQVariantFromPointer(value), int(int32(hits)), core.Qt__MatchFlag(flags)) {
				tmpList.__match_setList(v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewQSqlQueryModelFromPointer(NewQSqlQueryModelFromPointer(nil).__match_newList())
		for _, v := range NewQSqlQueryModelFromPointer(ptr).MatchDefault(core.NewQModelIndexFromPointer(start), int(int32(role)), core.NewQVariantFromPointer(value), int(int32(hits)), core.Qt__MatchFlag(flags)) {
			tmpList.__match_setList(v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *QSqlQueryModel) MatchDefault(start core.QModelIndex_ITF, role int, value core.QVariant_ITF, hits int, flags core.Qt__MatchFlag) []*core.QModelIndex {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtSql_PackedList) []*core.QModelIndex {
			out := make([]*core.QModelIndex, int(l.len))
			tmpList := NewQSqlQueryModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__match_atList(i)
			}
			return out
		}(C.QSqlQueryModel_MatchDefault(ptr.Pointer(), core.PointerFromQModelIndex(start), C.int(int32(role)), core.PointerFromQVariant(value), C.int(int32(hits)), C.longlong(flags)))
	}
	return make([]*core.QModelIndex, 0)
}

//export callbackQSqlQueryModel_Span
func callbackQSqlQueryModel_Span(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "span"); signal != nil {
		return core.PointerFromQSize(signal.(func(*core.QModelIndex) *core.QSize)(core.NewQModelIndexFromPointer(index)))
	}

	return core.PointerFromQSize(NewQSqlQueryModelFromPointer(ptr).SpanDefault(core.NewQModelIndexFromPointer(index)))
}

func (ptr *QSqlQueryModel) SpanDefault(index core.QModelIndex_ITF) *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QSqlQueryModel_SpanDefault(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQSqlQueryModel_MimeTypes
func callbackQSqlQueryModel_MimeTypes(ptr unsafe.Pointer) C.struct_QtSql_PackedString {
	if signal := qt.GetSignal(ptr, "mimeTypes"); signal != nil {
		tempVal := signal.(func() []string)()
		return C.struct_QtSql_PackedString{data: C.CString(strings.Join(tempVal, "|")), len: C.longlong(len(strings.Join(tempVal, "|")))}
	}
	tempVal := NewQSqlQueryModelFromPointer(ptr).MimeTypesDefault()
	return C.struct_QtSql_PackedString{data: C.CString(strings.Join(tempVal, "|")), len: C.longlong(len(strings.Join(tempVal, "|")))}
}

func (ptr *QSqlQueryModel) MimeTypesDefault() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QSqlQueryModel_MimeTypesDefault(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

//export callbackQSqlQueryModel_SupportedDragActions
func callbackQSqlQueryModel_SupportedDragActions(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "supportedDragActions"); signal != nil {
		return C.longlong(signal.(func() core.Qt__DropAction)())
	}

	return C.longlong(NewQSqlQueryModelFromPointer(ptr).SupportedDragActionsDefault())
}

func (ptr *QSqlQueryModel) SupportedDragActionsDefault() core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return core.Qt__DropAction(C.QSqlQueryModel_SupportedDragActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQSqlQueryModel_SupportedDropActions
func callbackQSqlQueryModel_SupportedDropActions(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "supportedDropActions"); signal != nil {
		return C.longlong(signal.(func() core.Qt__DropAction)())
	}

	return C.longlong(NewQSqlQueryModelFromPointer(ptr).SupportedDropActionsDefault())
}

func (ptr *QSqlQueryModel) SupportedDropActionsDefault() core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return core.Qt__DropAction(C.QSqlQueryModel_SupportedDropActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQSqlQueryModel_CanDropMimeData
func callbackQSqlQueryModel_CanDropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "canDropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QMimeData, core.Qt__DropAction, int, int, *core.QModelIndex) bool)(core.NewQMimeDataFromPointer(data), core.Qt__DropAction(action), int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlQueryModelFromPointer(ptr).CanDropMimeDataDefault(core.NewQMimeDataFromPointer(data), core.Qt__DropAction(action), int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QSqlQueryModel) CanDropMimeDataDefault(data core.QMimeData_ITF, action core.Qt__DropAction, row int, column int, parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlQueryModel_CanDropMimeDataDefault(ptr.Pointer(), core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackQSqlQueryModel_HasChildren
func callbackQSqlQueryModel_HasChildren(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasChildren"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QModelIndex) bool)(core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlQueryModelFromPointer(ptr).HasChildrenDefault(core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QSqlQueryModel) HasChildrenDefault(parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlQueryModel_HasChildrenDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackQSqlQueryModel_Event
func callbackQSqlQueryModel_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlQueryModelFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QSqlQueryModel) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlQueryModel_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQSqlQueryModel_EventFilter
func callbackQSqlQueryModel_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlQueryModelFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QSqlQueryModel) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlQueryModel_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQSqlQueryModel_ChildEvent
func callbackQSqlQueryModel_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSqlQueryModelFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSqlQueryModel) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlQueryModel_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQSqlQueryModel_ConnectNotify
func callbackQSqlQueryModel_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSqlQueryModelFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSqlQueryModel) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlQueryModel_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSqlQueryModel_CustomEvent
func callbackQSqlQueryModel_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSqlQueryModelFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSqlQueryModel) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlQueryModel_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQSqlQueryModel_DeleteLater
func callbackQSqlQueryModel_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQSqlQueryModelFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QSqlQueryModel) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QSqlQueryModel_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQSqlQueryModel_Destroyed
func callbackQSqlQueryModel_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQSqlQueryModel_DisconnectNotify
func callbackQSqlQueryModel_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSqlQueryModelFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSqlQueryModel) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlQueryModel_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSqlQueryModel_ObjectNameChanged
func callbackQSqlQueryModel_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtSql_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQSqlQueryModel_TimerEvent
func callbackQSqlQueryModel_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSqlQueryModelFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSqlQueryModel) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlQueryModel_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QSqlRecord struct {
	ptr unsafe.Pointer
}

type QSqlRecord_ITF interface {
	QSqlRecord_PTR() *QSqlRecord
}

func (ptr *QSqlRecord) QSqlRecord_PTR() *QSqlRecord {
	return ptr
}

func (ptr *QSqlRecord) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QSqlRecord) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQSqlRecord(ptr QSqlRecord_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlRecord_PTR().Pointer()
	}
	return nil
}

func NewQSqlRecordFromPointer(ptr unsafe.Pointer) (n *QSqlRecord) {
	n = new(QSqlRecord)
	n.SetPointer(ptr)
	return
}
func NewQSqlRecord() *QSqlRecord {
	tmpValue := NewQSqlRecordFromPointer(C.QSqlRecord_NewQSqlRecord())
	runtime.SetFinalizer(tmpValue, (*QSqlRecord).DestroyQSqlRecord)
	return tmpValue
}

func NewQSqlRecord2(other QSqlRecord_ITF) *QSqlRecord {
	tmpValue := NewQSqlRecordFromPointer(C.QSqlRecord_NewQSqlRecord2(PointerFromQSqlRecord(other)))
	runtime.SetFinalizer(tmpValue, (*QSqlRecord).DestroyQSqlRecord)
	return tmpValue
}

func (ptr *QSqlRecord) Append(field QSqlField_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlRecord_Append(ptr.Pointer(), PointerFromQSqlField(field))
	}
}

func (ptr *QSqlRecord) Clear() {
	if ptr.Pointer() != nil {
		C.QSqlRecord_Clear(ptr.Pointer())
	}
}

func (ptr *QSqlRecord) ClearValues() {
	if ptr.Pointer() != nil {
		C.QSqlRecord_ClearValues(ptr.Pointer())
	}
}

func (ptr *QSqlRecord) Insert(pos int, field QSqlField_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlRecord_Insert(ptr.Pointer(), C.int(int32(pos)), PointerFromQSqlField(field))
	}
}

func (ptr *QSqlRecord) Remove(pos int) {
	if ptr.Pointer() != nil {
		C.QSqlRecord_Remove(ptr.Pointer(), C.int(int32(pos)))
	}
}

func (ptr *QSqlRecord) Replace(pos int, field QSqlField_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlRecord_Replace(ptr.Pointer(), C.int(int32(pos)), PointerFromQSqlField(field))
	}
}

func (ptr *QSqlRecord) SetGenerated(name string, generated bool) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		C.QSqlRecord_SetGenerated(ptr.Pointer(), C.struct_QtSql_PackedString{data: nameC, len: C.longlong(len(name))}, C.char(int8(qt.GoBoolToInt(generated))))
	}
}

func (ptr *QSqlRecord) SetGenerated2(index int, generated bool) {
	if ptr.Pointer() != nil {
		C.QSqlRecord_SetGenerated2(ptr.Pointer(), C.int(int32(index)), C.char(int8(qt.GoBoolToInt(generated))))
	}
}

func (ptr *QSqlRecord) SetNull2(name string) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		C.QSqlRecord_SetNull2(ptr.Pointer(), C.struct_QtSql_PackedString{data: nameC, len: C.longlong(len(name))})
	}
}

func (ptr *QSqlRecord) SetNull(index int) {
	if ptr.Pointer() != nil {
		C.QSqlRecord_SetNull(ptr.Pointer(), C.int(int32(index)))
	}
}

func (ptr *QSqlRecord) SetValue2(name string, val core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		C.QSqlRecord_SetValue2(ptr.Pointer(), C.struct_QtSql_PackedString{data: nameC, len: C.longlong(len(name))}, core.PointerFromQVariant(val))
	}
}

func (ptr *QSqlRecord) SetValue(index int, val core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlRecord_SetValue(ptr.Pointer(), C.int(int32(index)), core.PointerFromQVariant(val))
	}
}

func (ptr *QSqlRecord) DestroyQSqlRecord() {
	if ptr.Pointer() != nil {
		C.QSqlRecord_DestroyQSqlRecord(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QSqlRecord) Field2(name string) *QSqlField {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		tmpValue := NewQSqlFieldFromPointer(C.QSqlRecord_Field2(ptr.Pointer(), C.struct_QtSql_PackedString{data: nameC, len: C.longlong(len(name))}))
		runtime.SetFinalizer(tmpValue, (*QSqlField).DestroyQSqlField)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlRecord) Field(index int) *QSqlField {
	if ptr.Pointer() != nil {
		tmpValue := NewQSqlFieldFromPointer(C.QSqlRecord_Field(ptr.Pointer(), C.int(int32(index))))
		runtime.SetFinalizer(tmpValue, (*QSqlField).DestroyQSqlField)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlRecord) KeyValues(keyFields QSqlRecord_ITF) *QSqlRecord {
	if ptr.Pointer() != nil {
		tmpValue := NewQSqlRecordFromPointer(C.QSqlRecord_KeyValues(ptr.Pointer(), PointerFromQSqlRecord(keyFields)))
		runtime.SetFinalizer(tmpValue, (*QSqlRecord).DestroyQSqlRecord)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlRecord) FieldName(index int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSqlRecord_FieldName(ptr.Pointer(), C.int(int32(index))))
	}
	return ""
}

func (ptr *QSqlRecord) Value2(name string) *core.QVariant {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		tmpValue := core.NewQVariantFromPointer(C.QSqlRecord_Value2(ptr.Pointer(), C.struct_QtSql_PackedString{data: nameC, len: C.longlong(len(name))}))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlRecord) Value(index int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QSqlRecord_Value(ptr.Pointer(), C.int(int32(index))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlRecord) Contains(name string) bool {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		return int8(C.QSqlRecord_Contains(ptr.Pointer(), C.struct_QtSql_PackedString{data: nameC, len: C.longlong(len(name))})) != 0
	}
	return false
}

func (ptr *QSqlRecord) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlRecord_IsEmpty(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSqlRecord) IsGenerated(name string) bool {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		return int8(C.QSqlRecord_IsGenerated(ptr.Pointer(), C.struct_QtSql_PackedString{data: nameC, len: C.longlong(len(name))})) != 0
	}
	return false
}

func (ptr *QSqlRecord) IsGenerated2(index int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlRecord_IsGenerated2(ptr.Pointer(), C.int(int32(index)))) != 0
	}
	return false
}

func (ptr *QSqlRecord) IsNull(name string) bool {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		return int8(C.QSqlRecord_IsNull(ptr.Pointer(), C.struct_QtSql_PackedString{data: nameC, len: C.longlong(len(name))})) != 0
	}
	return false
}

func (ptr *QSqlRecord) IsNull2(index int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlRecord_IsNull2(ptr.Pointer(), C.int(int32(index)))) != 0
	}
	return false
}

func (ptr *QSqlRecord) Count() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSqlRecord_Count(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSqlRecord) IndexOf(name string) int {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		return int(int32(C.QSqlRecord_IndexOf(ptr.Pointer(), C.struct_QtSql_PackedString{data: nameC, len: C.longlong(len(name))})))
	}
	return 0
}

type QSqlRelation struct {
	ptr unsafe.Pointer
}

type QSqlRelation_ITF interface {
	QSqlRelation_PTR() *QSqlRelation
}

func (ptr *QSqlRelation) QSqlRelation_PTR() *QSqlRelation {
	return ptr
}

func (ptr *QSqlRelation) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QSqlRelation) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQSqlRelation(ptr QSqlRelation_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlRelation_PTR().Pointer()
	}
	return nil
}

func NewQSqlRelationFromPointer(ptr unsafe.Pointer) (n *QSqlRelation) {
	n = new(QSqlRelation)
	n.SetPointer(ptr)
	return
}

func (ptr *QSqlRelation) DestroyQSqlRelation() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func NewQSqlRelation() *QSqlRelation {
	tmpValue := NewQSqlRelationFromPointer(C.QSqlRelation_NewQSqlRelation())
	runtime.SetFinalizer(tmpValue, (*QSqlRelation).DestroyQSqlRelation)
	return tmpValue
}

func NewQSqlRelation2(tableName string, indexColumn string, displayColumn string) *QSqlRelation {
	var tableNameC *C.char
	if tableName != "" {
		tableNameC = C.CString(tableName)
		defer C.free(unsafe.Pointer(tableNameC))
	}
	var indexColumnC *C.char
	if indexColumn != "" {
		indexColumnC = C.CString(indexColumn)
		defer C.free(unsafe.Pointer(indexColumnC))
	}
	var displayColumnC *C.char
	if displayColumn != "" {
		displayColumnC = C.CString(displayColumn)
		defer C.free(unsafe.Pointer(displayColumnC))
	}
	tmpValue := NewQSqlRelationFromPointer(C.QSqlRelation_NewQSqlRelation2(C.struct_QtSql_PackedString{data: tableNameC, len: C.longlong(len(tableName))}, C.struct_QtSql_PackedString{data: indexColumnC, len: C.longlong(len(indexColumn))}, C.struct_QtSql_PackedString{data: displayColumnC, len: C.longlong(len(displayColumn))}))
	runtime.SetFinalizer(tmpValue, (*QSqlRelation).DestroyQSqlRelation)
	return tmpValue
}

func (ptr *QSqlRelation) Swap(other QSqlRelation_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlRelation_Swap(ptr.Pointer(), PointerFromQSqlRelation(other))
	}
}

func (ptr *QSqlRelation) DisplayColumn() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSqlRelation_DisplayColumn(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlRelation) IndexColumn() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSqlRelation_IndexColumn(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlRelation) TableName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSqlRelation_TableName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlRelation) IsValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlRelation_IsValid(ptr.Pointer())) != 0
	}
	return false
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

func NewQSqlRelationalDelegateFromPointer(ptr unsafe.Pointer) (n *QSqlRelationalDelegate) {
	n = new(QSqlRelationalDelegate)
	n.SetPointer(ptr)
	return
}
func NewQSqlRelationalDelegate(parent core.QObject_ITF) *QSqlRelationalDelegate {
	tmpValue := NewQSqlRelationalDelegateFromPointer(C.QSqlRelationalDelegate_NewQSqlRelationalDelegate(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQSqlRelationalDelegate_DestroyQSqlRelationalDelegate
func callbackQSqlRelationalDelegate_DestroyQSqlRelationalDelegate(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QSqlRelationalDelegate"); signal != nil {
		signal.(func())()
	} else {
		NewQSqlRelationalDelegateFromPointer(ptr).DestroyQSqlRelationalDelegateDefault()
	}
}

func (ptr *QSqlRelationalDelegate) ConnectDestroyQSqlRelationalDelegate(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QSqlRelationalDelegate"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QSqlRelationalDelegate", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QSqlRelationalDelegate", f)
		}
	}
}

func (ptr *QSqlRelationalDelegate) DisconnectDestroyQSqlRelationalDelegate() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QSqlRelationalDelegate")
	}
}

func (ptr *QSqlRelationalDelegate) DestroyQSqlRelationalDelegate() {
	if ptr.Pointer() != nil {
		C.QSqlRelationalDelegate_DestroyQSqlRelationalDelegate(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QSqlRelationalDelegate) DestroyQSqlRelationalDelegateDefault() {
	if ptr.Pointer() != nil {
		C.QSqlRelationalDelegate_DestroyQSqlRelationalDelegateDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQSqlRelationalDelegate_CreateEditor
func callbackQSqlRelationalDelegate_CreateEditor(ptr unsafe.Pointer, parent unsafe.Pointer, option unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "createEditor"); signal != nil {
		return widgets.PointerFromQWidget(signal.(func(*widgets.QWidget, *widgets.QStyleOptionViewItem, *core.QModelIndex) *widgets.QWidget)(widgets.NewQWidgetFromPointer(parent), widgets.NewQStyleOptionViewItemFromPointer(option), core.NewQModelIndexFromPointer(index)))
	}

	return widgets.PointerFromQWidget(NewQSqlRelationalDelegateFromPointer(ptr).CreateEditorDefault(widgets.NewQWidgetFromPointer(parent), widgets.NewQStyleOptionViewItemFromPointer(option), core.NewQModelIndexFromPointer(index)))
}

func (ptr *QSqlRelationalDelegate) CreateEditorDefault(parent widgets.QWidget_ITF, option widgets.QStyleOptionViewItem_ITF, index core.QModelIndex_ITF) *widgets.QWidget {
	if ptr.Pointer() != nil {
		tmpValue := widgets.NewQWidgetFromPointer(C.QSqlRelationalDelegate_CreateEditorDefault(ptr.Pointer(), widgets.PointerFromQWidget(parent), widgets.PointerFromQStyleOptionViewItem(option), core.PointerFromQModelIndex(index)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQSqlRelationalDelegate_SetEditorData
func callbackQSqlRelationalDelegate_SetEditorData(ptr unsafe.Pointer, editor unsafe.Pointer, index unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setEditorData"); signal != nil {
		signal.(func(*widgets.QWidget, *core.QModelIndex))(widgets.NewQWidgetFromPointer(editor), core.NewQModelIndexFromPointer(index))
	} else {
		NewQSqlRelationalDelegateFromPointer(ptr).SetEditorDataDefault(widgets.NewQWidgetFromPointer(editor), core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *QSqlRelationalDelegate) SetEditorDataDefault(editor widgets.QWidget_ITF, index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlRelationalDelegate_SetEditorDataDefault(ptr.Pointer(), widgets.PointerFromQWidget(editor), core.PointerFromQModelIndex(index))
	}
}

//export callbackQSqlRelationalDelegate_SetModelData
func callbackQSqlRelationalDelegate_SetModelData(ptr unsafe.Pointer, editor unsafe.Pointer, model unsafe.Pointer, index unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setModelData"); signal != nil {
		signal.(func(*widgets.QWidget, *core.QAbstractItemModel, *core.QModelIndex))(widgets.NewQWidgetFromPointer(editor), core.NewQAbstractItemModelFromPointer(model), core.NewQModelIndexFromPointer(index))
	} else {
		NewQSqlRelationalDelegateFromPointer(ptr).SetModelDataDefault(widgets.NewQWidgetFromPointer(editor), core.NewQAbstractItemModelFromPointer(model), core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *QSqlRelationalDelegate) SetModelDataDefault(editor widgets.QWidget_ITF, model core.QAbstractItemModel_ITF, index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlRelationalDelegate_SetModelDataDefault(ptr.Pointer(), widgets.PointerFromQWidget(editor), core.PointerFromQAbstractItemModel(model), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QSqlRelationalDelegate) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QSqlRelationalDelegate___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlRelationalDelegate) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlRelationalDelegate___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QSqlRelationalDelegate) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QSqlRelationalDelegate___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QSqlRelationalDelegate) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QSqlRelationalDelegate___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSqlRelationalDelegate) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlRelationalDelegate___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QSqlRelationalDelegate) __findChildren_newList2() unsafe.Pointer {
	return C.QSqlRelationalDelegate___findChildren_newList2(ptr.Pointer())
}

func (ptr *QSqlRelationalDelegate) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QSqlRelationalDelegate___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSqlRelationalDelegate) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlRelationalDelegate___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QSqlRelationalDelegate) __findChildren_newList3() unsafe.Pointer {
	return C.QSqlRelationalDelegate___findChildren_newList3(ptr.Pointer())
}

func (ptr *QSqlRelationalDelegate) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QSqlRelationalDelegate___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSqlRelationalDelegate) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlRelationalDelegate___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QSqlRelationalDelegate) __findChildren_newList() unsafe.Pointer {
	return C.QSqlRelationalDelegate___findChildren_newList(ptr.Pointer())
}

func (ptr *QSqlRelationalDelegate) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QSqlRelationalDelegate___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSqlRelationalDelegate) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlRelationalDelegate___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QSqlRelationalDelegate) __children_newList() unsafe.Pointer {
	return C.QSqlRelationalDelegate___children_newList(ptr.Pointer())
}

//export callbackQSqlRelationalDelegate_EditorEvent
func callbackQSqlRelationalDelegate_EditorEvent(ptr unsafe.Pointer, event unsafe.Pointer, model unsafe.Pointer, option unsafe.Pointer, index unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "editorEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent, *core.QAbstractItemModel, *widgets.QStyleOptionViewItem, *core.QModelIndex) bool)(core.NewQEventFromPointer(event), core.NewQAbstractItemModelFromPointer(model), widgets.NewQStyleOptionViewItemFromPointer(option), core.NewQModelIndexFromPointer(index)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlRelationalDelegateFromPointer(ptr).EditorEventDefault(core.NewQEventFromPointer(event), core.NewQAbstractItemModelFromPointer(model), widgets.NewQStyleOptionViewItemFromPointer(option), core.NewQModelIndexFromPointer(index)))))
}

func (ptr *QSqlRelationalDelegate) EditorEventDefault(event core.QEvent_ITF, model core.QAbstractItemModel_ITF, option widgets.QStyleOptionViewItem_ITF, index core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlRelationalDelegate_EditorEventDefault(ptr.Pointer(), core.PointerFromQEvent(event), core.PointerFromQAbstractItemModel(model), widgets.PointerFromQStyleOptionViewItem(option), core.PointerFromQModelIndex(index))) != 0
	}
	return false
}

//export callbackQSqlRelationalDelegate_EventFilter
func callbackQSqlRelationalDelegate_EventFilter(ptr unsafe.Pointer, editor unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(editor), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlRelationalDelegateFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(editor), core.NewQEventFromPointer(event)))))
}

func (ptr *QSqlRelationalDelegate) EventFilterDefault(editor core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlRelationalDelegate_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(editor), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQSqlRelationalDelegate_SizeHint
func callbackQSqlRelationalDelegate_SizeHint(ptr unsafe.Pointer, option unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func(*widgets.QStyleOptionViewItem, *core.QModelIndex) *core.QSize)(widgets.NewQStyleOptionViewItemFromPointer(option), core.NewQModelIndexFromPointer(index)))
	}

	return core.PointerFromQSize(NewQSqlRelationalDelegateFromPointer(ptr).SizeHintDefault(widgets.NewQStyleOptionViewItemFromPointer(option), core.NewQModelIndexFromPointer(index)))
}

func (ptr *QSqlRelationalDelegate) SizeHintDefault(option widgets.QStyleOptionViewItem_ITF, index core.QModelIndex_ITF) *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QSqlRelationalDelegate_SizeHintDefault(ptr.Pointer(), widgets.PointerFromQStyleOptionViewItem(option), core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQSqlRelationalDelegate_MetaObject
func callbackQSqlRelationalDelegate_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQSqlRelationalDelegateFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QSqlRelationalDelegate) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSqlRelationalDelegate_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQSqlRelationalDelegate_DrawCheck
func callbackQSqlRelationalDelegate_DrawCheck(ptr unsafe.Pointer, painter unsafe.Pointer, option unsafe.Pointer, rect unsafe.Pointer, state C.longlong) {
	if signal := qt.GetSignal(ptr, "drawCheck"); signal != nil {
		signal.(func(*gui.QPainter, *widgets.QStyleOptionViewItem, *core.QRect, core.Qt__CheckState))(gui.NewQPainterFromPointer(painter), widgets.NewQStyleOptionViewItemFromPointer(option), core.NewQRectFromPointer(rect), core.Qt__CheckState(state))
	} else {
		NewQSqlRelationalDelegateFromPointer(ptr).DrawCheckDefault(gui.NewQPainterFromPointer(painter), widgets.NewQStyleOptionViewItemFromPointer(option), core.NewQRectFromPointer(rect), core.Qt__CheckState(state))
	}
}

func (ptr *QSqlRelationalDelegate) DrawCheckDefault(painter gui.QPainter_ITF, option widgets.QStyleOptionViewItem_ITF, rect core.QRect_ITF, state core.Qt__CheckState) {
	if ptr.Pointer() != nil {
		C.QSqlRelationalDelegate_DrawCheckDefault(ptr.Pointer(), gui.PointerFromQPainter(painter), widgets.PointerFromQStyleOptionViewItem(option), core.PointerFromQRect(rect), C.longlong(state))
	}
}

//export callbackQSqlRelationalDelegate_DrawDecoration
func callbackQSqlRelationalDelegate_DrawDecoration(ptr unsafe.Pointer, painter unsafe.Pointer, option unsafe.Pointer, rect unsafe.Pointer, pixmap unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "drawDecoration"); signal != nil {
		signal.(func(*gui.QPainter, *widgets.QStyleOptionViewItem, *core.QRect, *gui.QPixmap))(gui.NewQPainterFromPointer(painter), widgets.NewQStyleOptionViewItemFromPointer(option), core.NewQRectFromPointer(rect), gui.NewQPixmapFromPointer(pixmap))
	} else {
		NewQSqlRelationalDelegateFromPointer(ptr).DrawDecorationDefault(gui.NewQPainterFromPointer(painter), widgets.NewQStyleOptionViewItemFromPointer(option), core.NewQRectFromPointer(rect), gui.NewQPixmapFromPointer(pixmap))
	}
}

func (ptr *QSqlRelationalDelegate) DrawDecorationDefault(painter gui.QPainter_ITF, option widgets.QStyleOptionViewItem_ITF, rect core.QRect_ITF, pixmap gui.QPixmap_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlRelationalDelegate_DrawDecorationDefault(ptr.Pointer(), gui.PointerFromQPainter(painter), widgets.PointerFromQStyleOptionViewItem(option), core.PointerFromQRect(rect), gui.PointerFromQPixmap(pixmap))
	}
}

//export callbackQSqlRelationalDelegate_DrawDisplay
func callbackQSqlRelationalDelegate_DrawDisplay(ptr unsafe.Pointer, painter unsafe.Pointer, option unsafe.Pointer, rect unsafe.Pointer, text C.struct_QtSql_PackedString) {
	if signal := qt.GetSignal(ptr, "drawDisplay"); signal != nil {
		signal.(func(*gui.QPainter, *widgets.QStyleOptionViewItem, *core.QRect, string))(gui.NewQPainterFromPointer(painter), widgets.NewQStyleOptionViewItemFromPointer(option), core.NewQRectFromPointer(rect), cGoUnpackString(text))
	} else {
		NewQSqlRelationalDelegateFromPointer(ptr).DrawDisplayDefault(gui.NewQPainterFromPointer(painter), widgets.NewQStyleOptionViewItemFromPointer(option), core.NewQRectFromPointer(rect), cGoUnpackString(text))
	}
}

func (ptr *QSqlRelationalDelegate) DrawDisplayDefault(painter gui.QPainter_ITF, option widgets.QStyleOptionViewItem_ITF, rect core.QRect_ITF, text string) {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		C.QSqlRelationalDelegate_DrawDisplayDefault(ptr.Pointer(), gui.PointerFromQPainter(painter), widgets.PointerFromQStyleOptionViewItem(option), core.PointerFromQRect(rect), C.struct_QtSql_PackedString{data: textC, len: C.longlong(len(text))})
	}
}

//export callbackQSqlRelationalDelegate_DrawFocus
func callbackQSqlRelationalDelegate_DrawFocus(ptr unsafe.Pointer, painter unsafe.Pointer, option unsafe.Pointer, rect unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "drawFocus"); signal != nil {
		signal.(func(*gui.QPainter, *widgets.QStyleOptionViewItem, *core.QRect))(gui.NewQPainterFromPointer(painter), widgets.NewQStyleOptionViewItemFromPointer(option), core.NewQRectFromPointer(rect))
	} else {
		NewQSqlRelationalDelegateFromPointer(ptr).DrawFocusDefault(gui.NewQPainterFromPointer(painter), widgets.NewQStyleOptionViewItemFromPointer(option), core.NewQRectFromPointer(rect))
	}
}

func (ptr *QSqlRelationalDelegate) DrawFocusDefault(painter gui.QPainter_ITF, option widgets.QStyleOptionViewItem_ITF, rect core.QRect_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlRelationalDelegate_DrawFocusDefault(ptr.Pointer(), gui.PointerFromQPainter(painter), widgets.PointerFromQStyleOptionViewItem(option), core.PointerFromQRect(rect))
	}
}

//export callbackQSqlRelationalDelegate_Paint
func callbackQSqlRelationalDelegate_Paint(ptr unsafe.Pointer, painter unsafe.Pointer, option unsafe.Pointer, index unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "paint"); signal != nil {
		signal.(func(*gui.QPainter, *widgets.QStyleOptionViewItem, *core.QModelIndex))(gui.NewQPainterFromPointer(painter), widgets.NewQStyleOptionViewItemFromPointer(option), core.NewQModelIndexFromPointer(index))
	} else {
		NewQSqlRelationalDelegateFromPointer(ptr).PaintDefault(gui.NewQPainterFromPointer(painter), widgets.NewQStyleOptionViewItemFromPointer(option), core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *QSqlRelationalDelegate) PaintDefault(painter gui.QPainter_ITF, option widgets.QStyleOptionViewItem_ITF, index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlRelationalDelegate_PaintDefault(ptr.Pointer(), gui.PointerFromQPainter(painter), widgets.PointerFromQStyleOptionViewItem(option), core.PointerFromQModelIndex(index))
	}
}

//export callbackQSqlRelationalDelegate_UpdateEditorGeometry
func callbackQSqlRelationalDelegate_UpdateEditorGeometry(ptr unsafe.Pointer, editor unsafe.Pointer, option unsafe.Pointer, index unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "updateEditorGeometry"); signal != nil {
		signal.(func(*widgets.QWidget, *widgets.QStyleOptionViewItem, *core.QModelIndex))(widgets.NewQWidgetFromPointer(editor), widgets.NewQStyleOptionViewItemFromPointer(option), core.NewQModelIndexFromPointer(index))
	} else {
		NewQSqlRelationalDelegateFromPointer(ptr).UpdateEditorGeometryDefault(widgets.NewQWidgetFromPointer(editor), widgets.NewQStyleOptionViewItemFromPointer(option), core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *QSqlRelationalDelegate) UpdateEditorGeometryDefault(editor widgets.QWidget_ITF, option widgets.QStyleOptionViewItem_ITF, index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlRelationalDelegate_UpdateEditorGeometryDefault(ptr.Pointer(), widgets.PointerFromQWidget(editor), widgets.PointerFromQStyleOptionViewItem(option), core.PointerFromQModelIndex(index))
	}
}

//export callbackQSqlRelationalDelegate_HelpEvent
func callbackQSqlRelationalDelegate_HelpEvent(ptr unsafe.Pointer, event unsafe.Pointer, view unsafe.Pointer, option unsafe.Pointer, index unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "helpEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*gui.QHelpEvent, *widgets.QAbstractItemView, *widgets.QStyleOptionViewItem, *core.QModelIndex) bool)(gui.NewQHelpEventFromPointer(event), widgets.NewQAbstractItemViewFromPointer(view), widgets.NewQStyleOptionViewItemFromPointer(option), core.NewQModelIndexFromPointer(index)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlRelationalDelegateFromPointer(ptr).HelpEventDefault(gui.NewQHelpEventFromPointer(event), widgets.NewQAbstractItemViewFromPointer(view), widgets.NewQStyleOptionViewItemFromPointer(option), core.NewQModelIndexFromPointer(index)))))
}

func (ptr *QSqlRelationalDelegate) HelpEventDefault(event gui.QHelpEvent_ITF, view widgets.QAbstractItemView_ITF, option widgets.QStyleOptionViewItem_ITF, index core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlRelationalDelegate_HelpEventDefault(ptr.Pointer(), gui.PointerFromQHelpEvent(event), widgets.PointerFromQAbstractItemView(view), widgets.PointerFromQStyleOptionViewItem(option), core.PointerFromQModelIndex(index))) != 0
	}
	return false
}

//export callbackQSqlRelationalDelegate_CloseEditor
func callbackQSqlRelationalDelegate_CloseEditor(ptr unsafe.Pointer, editor unsafe.Pointer, hint C.longlong) {
	if signal := qt.GetSignal(ptr, "closeEditor"); signal != nil {
		signal.(func(*widgets.QWidget, widgets.QAbstractItemDelegate__EndEditHint))(widgets.NewQWidgetFromPointer(editor), widgets.QAbstractItemDelegate__EndEditHint(hint))
	}

}

//export callbackQSqlRelationalDelegate_CommitData
func callbackQSqlRelationalDelegate_CommitData(ptr unsafe.Pointer, editor unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "commitData"); signal != nil {
		signal.(func(*widgets.QWidget))(widgets.NewQWidgetFromPointer(editor))
	}

}

//export callbackQSqlRelationalDelegate_SizeHintChanged
func callbackQSqlRelationalDelegate_SizeHintChanged(ptr unsafe.Pointer, index unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "sizeHintChanged"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
	}

}

//export callbackQSqlRelationalDelegate_DestroyEditor
func callbackQSqlRelationalDelegate_DestroyEditor(ptr unsafe.Pointer, editor unsafe.Pointer, index unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyEditor"); signal != nil {
		signal.(func(*widgets.QWidget, *core.QModelIndex))(widgets.NewQWidgetFromPointer(editor), core.NewQModelIndexFromPointer(index))
	} else {
		NewQSqlRelationalDelegateFromPointer(ptr).DestroyEditorDefault(widgets.NewQWidgetFromPointer(editor), core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *QSqlRelationalDelegate) DestroyEditorDefault(editor widgets.QWidget_ITF, index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlRelationalDelegate_DestroyEditorDefault(ptr.Pointer(), widgets.PointerFromQWidget(editor), core.PointerFromQModelIndex(index))
	}
}

//export callbackQSqlRelationalDelegate_Event
func callbackQSqlRelationalDelegate_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlRelationalDelegateFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QSqlRelationalDelegate) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlRelationalDelegate_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQSqlRelationalDelegate_ChildEvent
func callbackQSqlRelationalDelegate_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSqlRelationalDelegateFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSqlRelationalDelegate) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlRelationalDelegate_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQSqlRelationalDelegate_ConnectNotify
func callbackQSqlRelationalDelegate_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSqlRelationalDelegateFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSqlRelationalDelegate) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlRelationalDelegate_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSqlRelationalDelegate_CustomEvent
func callbackQSqlRelationalDelegate_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSqlRelationalDelegateFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSqlRelationalDelegate) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlRelationalDelegate_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQSqlRelationalDelegate_DeleteLater
func callbackQSqlRelationalDelegate_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQSqlRelationalDelegateFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QSqlRelationalDelegate) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QSqlRelationalDelegate_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQSqlRelationalDelegate_Destroyed
func callbackQSqlRelationalDelegate_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQSqlRelationalDelegate_DisconnectNotify
func callbackQSqlRelationalDelegate_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSqlRelationalDelegateFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSqlRelationalDelegate) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlRelationalDelegate_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSqlRelationalDelegate_ObjectNameChanged
func callbackQSqlRelationalDelegate_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtSql_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQSqlRelationalDelegate_TimerEvent
func callbackQSqlRelationalDelegate_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSqlRelationalDelegateFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSqlRelationalDelegate) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlRelationalDelegate_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
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

func NewQSqlRelationalTableModelFromPointer(ptr unsafe.Pointer) (n *QSqlRelationalTableModel) {
	n = new(QSqlRelationalTableModel)
	n.SetPointer(ptr)
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
	tmpValue := NewQSqlRelationalTableModelFromPointer(C.QSqlRelationalTableModel_NewQSqlRelationalTableModel(core.PointerFromQObject(parent), PointerFromQSqlDatabase(db)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQSqlRelationalTableModel_Select
func callbackQSqlRelationalTableModel_Select(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "select"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlRelationalTableModelFromPointer(ptr).SelectDefault())))
}

func (ptr *QSqlRelationalTableModel) ConnectSelect(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "select"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "select", func() bool {
				signal.(func() bool)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "select", f)
		}
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectSelect() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "select")
	}
}

func (ptr *QSqlRelationalTableModel) Select() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlRelationalTableModel_Select(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSqlRelationalTableModel) SelectDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlRelationalTableModel_SelectDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQSqlRelationalTableModel_RevertRow
func callbackQSqlRelationalTableModel_RevertRow(ptr unsafe.Pointer, row C.int) {
	if signal := qt.GetSignal(ptr, "revertRow"); signal != nil {
		signal.(func(int))(int(int32(row)))
	} else {
		NewQSqlRelationalTableModelFromPointer(ptr).RevertRowDefault(int(int32(row)))
	}
}

func (ptr *QSqlRelationalTableModel) ConnectRevertRow(f func(row int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "revertRow"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "revertRow", func(row int) {
				signal.(func(int))(row)
				f(row)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "revertRow", f)
		}
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectRevertRow() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "revertRow")
	}
}

func (ptr *QSqlRelationalTableModel) RevertRow(row int) {
	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_RevertRow(ptr.Pointer(), C.int(int32(row)))
	}
}

func (ptr *QSqlRelationalTableModel) RevertRowDefault(row int) {
	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_RevertRowDefault(ptr.Pointer(), C.int(int32(row)))
	}
}

func (ptr *QSqlRelationalTableModel) SetJoinMode(joinMode QSqlRelationalTableModel__JoinMode) {
	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_SetJoinMode(ptr.Pointer(), C.longlong(joinMode))
	}
}

//export callbackQSqlRelationalTableModel_SetRelation
func callbackQSqlRelationalTableModel_SetRelation(ptr unsafe.Pointer, column C.int, relation unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setRelation"); signal != nil {
		signal.(func(int, *QSqlRelation))(int(int32(column)), NewQSqlRelationFromPointer(relation))
	} else {
		NewQSqlRelationalTableModelFromPointer(ptr).SetRelationDefault(int(int32(column)), NewQSqlRelationFromPointer(relation))
	}
}

func (ptr *QSqlRelationalTableModel) ConnectSetRelation(f func(column int, relation *QSqlRelation)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setRelation"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setRelation", func(column int, relation *QSqlRelation) {
				signal.(func(int, *QSqlRelation))(column, relation)
				f(column, relation)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setRelation", f)
		}
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectSetRelation() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setRelation")
	}
}

func (ptr *QSqlRelationalTableModel) SetRelation(column int, relation QSqlRelation_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_SetRelation(ptr.Pointer(), C.int(int32(column)), PointerFromQSqlRelation(relation))
	}
}

func (ptr *QSqlRelationalTableModel) SetRelationDefault(column int, relation QSqlRelation_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_SetRelationDefault(ptr.Pointer(), C.int(int32(column)), PointerFromQSqlRelation(relation))
	}
}

//export callbackQSqlRelationalTableModel_DestroyQSqlRelationalTableModel
func callbackQSqlRelationalTableModel_DestroyQSqlRelationalTableModel(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QSqlRelationalTableModel"); signal != nil {
		signal.(func())()
	} else {
		NewQSqlRelationalTableModelFromPointer(ptr).DestroyQSqlRelationalTableModelDefault()
	}
}

func (ptr *QSqlRelationalTableModel) ConnectDestroyQSqlRelationalTableModel(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QSqlRelationalTableModel"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QSqlRelationalTableModel", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QSqlRelationalTableModel", f)
		}
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectDestroyQSqlRelationalTableModel() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QSqlRelationalTableModel")
	}
}

func (ptr *QSqlRelationalTableModel) DestroyQSqlRelationalTableModel() {
	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_DestroyQSqlRelationalTableModel(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QSqlRelationalTableModel) DestroyQSqlRelationalTableModelDefault() {
	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_DestroyQSqlRelationalTableModelDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QSqlRelationalTableModel) Relation(column int) *QSqlRelation {
	if ptr.Pointer() != nil {
		tmpValue := NewQSqlRelationFromPointer(C.QSqlRelationalTableModel_Relation(ptr.Pointer(), C.int(int32(column))))
		runtime.SetFinalizer(tmpValue, (*QSqlRelation).DestroyQSqlRelation)
		return tmpValue
	}
	return nil
}

//export callbackQSqlRelationalTableModel_RelationModel
func callbackQSqlRelationalTableModel_RelationModel(ptr unsafe.Pointer, column C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "relationModel"); signal != nil {
		return PointerFromQSqlTableModel(signal.(func(int) *QSqlTableModel)(int(int32(column))))
	}

	return PointerFromQSqlTableModel(NewQSqlRelationalTableModelFromPointer(ptr).RelationModelDefault(int(int32(column))))
}

func (ptr *QSqlRelationalTableModel) ConnectRelationModel(f func(column int) *QSqlTableModel) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "relationModel"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "relationModel", func(column int) *QSqlTableModel {
				signal.(func(int) *QSqlTableModel)(column)
				return f(column)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "relationModel", f)
		}
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectRelationModel() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "relationModel")
	}
}

func (ptr *QSqlRelationalTableModel) RelationModel(column int) *QSqlTableModel {
	if ptr.Pointer() != nil {
		tmpValue := NewQSqlTableModelFromPointer(C.QSqlRelationalTableModel_RelationModel(ptr.Pointer(), C.int(int32(column))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSqlRelationalTableModel) RelationModelDefault(column int) *QSqlTableModel {
	if ptr.Pointer() != nil {
		tmpValue := NewQSqlTableModelFromPointer(C.QSqlRelationalTableModel_RelationModelDefault(ptr.Pointer(), C.int(int32(column))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

type QSqlResult struct {
	ptr unsafe.Pointer
}

type QSqlResult_ITF interface {
	QSqlResult_PTR() *QSqlResult
}

func (ptr *QSqlResult) QSqlResult_PTR() *QSqlResult {
	return ptr
}

func (ptr *QSqlResult) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QSqlResult) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQSqlResult(ptr QSqlResult_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlResult_PTR().Pointer()
	}
	return nil
}

func NewQSqlResultFromPointer(ptr unsafe.Pointer) (n *QSqlResult) {
	n = new(QSqlResult)
	n.SetPointer(ptr)
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
	return NewQSqlResultFromPointer(C.QSqlResult_NewQSqlResult(PointerFromQSqlDriver(db)))
}

//export callbackQSqlResult_Data
func callbackQSqlResult_Data(ptr unsafe.Pointer, index C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "data"); signal != nil {
		return core.PointerFromQVariant(signal.(func(int) *core.QVariant)(int(int32(index))))
	}

	return core.PointerFromQVariant(core.NewQVariant())
}

func (ptr *QSqlResult) ConnectData(f func(index int) *core.QVariant) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "data"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "data", func(index int) *core.QVariant {
				signal.(func(int) *core.QVariant)(index)
				return f(index)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "data", f)
		}
	}
}

func (ptr *QSqlResult) DisconnectData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "data")
	}
}

func (ptr *QSqlResult) Data(index int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QSqlResult_Data(ptr.Pointer(), C.int(int32(index))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQSqlResult_Exec
func callbackQSqlResult_Exec(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "exec"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlResultFromPointer(ptr).ExecDefault())))
}

func (ptr *QSqlResult) ConnectExec(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "exec"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "exec", func() bool {
				signal.(func() bool)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "exec", f)
		}
	}
}

func (ptr *QSqlResult) DisconnectExec() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "exec")
	}
}

func (ptr *QSqlResult) Exec() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlResult_Exec(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSqlResult) ExecDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlResult_ExecDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQSqlResult_Fetch
func callbackQSqlResult_Fetch(ptr unsafe.Pointer, index C.int) C.char {
	if signal := qt.GetSignal(ptr, "fetch"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(index))))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QSqlResult) ConnectFetch(f func(index int) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "fetch"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "fetch", func(index int) bool {
				signal.(func(int) bool)(index)
				return f(index)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "fetch", f)
		}
	}
}

func (ptr *QSqlResult) DisconnectFetch() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "fetch")
	}
}

func (ptr *QSqlResult) Fetch(index int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlResult_Fetch(ptr.Pointer(), C.int(int32(index)))) != 0
	}
	return false
}

//export callbackQSqlResult_FetchFirst
func callbackQSqlResult_FetchFirst(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "fetchFirst"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QSqlResult) ConnectFetchFirst(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "fetchFirst"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "fetchFirst", func() bool {
				signal.(func() bool)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "fetchFirst", f)
		}
	}
}

func (ptr *QSqlResult) DisconnectFetchFirst() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "fetchFirst")
	}
}

func (ptr *QSqlResult) FetchFirst() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlResult_FetchFirst(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQSqlResult_FetchLast
func callbackQSqlResult_FetchLast(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "fetchLast"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QSqlResult) ConnectFetchLast(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "fetchLast"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "fetchLast", func() bool {
				signal.(func() bool)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "fetchLast", f)
		}
	}
}

func (ptr *QSqlResult) DisconnectFetchLast() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "fetchLast")
	}
}

func (ptr *QSqlResult) FetchLast() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlResult_FetchLast(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQSqlResult_FetchNext
func callbackQSqlResult_FetchNext(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "fetchNext"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlResultFromPointer(ptr).FetchNextDefault())))
}

func (ptr *QSqlResult) ConnectFetchNext(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "fetchNext"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "fetchNext", func() bool {
				signal.(func() bool)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "fetchNext", f)
		}
	}
}

func (ptr *QSqlResult) DisconnectFetchNext() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "fetchNext")
	}
}

func (ptr *QSqlResult) FetchNext() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlResult_FetchNext(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSqlResult) FetchNextDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlResult_FetchNextDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQSqlResult_FetchPrevious
func callbackQSqlResult_FetchPrevious(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "fetchPrevious"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlResultFromPointer(ptr).FetchPreviousDefault())))
}

func (ptr *QSqlResult) ConnectFetchPrevious(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "fetchPrevious"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "fetchPrevious", func() bool {
				signal.(func() bool)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "fetchPrevious", f)
		}
	}
}

func (ptr *QSqlResult) DisconnectFetchPrevious() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "fetchPrevious")
	}
}

func (ptr *QSqlResult) FetchPrevious() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlResult_FetchPrevious(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSqlResult) FetchPreviousDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlResult_FetchPreviousDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQSqlResult_IsNull
func callbackQSqlResult_IsNull(ptr unsafe.Pointer, index C.int) C.char {
	if signal := qt.GetSignal(ptr, "isNull"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(index))))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QSqlResult) ConnectIsNull(f func(index int) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "isNull"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "isNull", func(index int) bool {
				signal.(func(int) bool)(index)
				return f(index)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "isNull", f)
		}
	}
}

func (ptr *QSqlResult) DisconnectIsNull() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "isNull")
	}
}

func (ptr *QSqlResult) IsNull(index int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlResult_IsNull(ptr.Pointer(), C.int(int32(index)))) != 0
	}
	return false
}

//export callbackQSqlResult_Prepare
func callbackQSqlResult_Prepare(ptr unsafe.Pointer, query C.struct_QtSql_PackedString) C.char {
	if signal := qt.GetSignal(ptr, "prepare"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string) bool)(cGoUnpackString(query)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlResultFromPointer(ptr).PrepareDefault(cGoUnpackString(query)))))
}

func (ptr *QSqlResult) ConnectPrepare(f func(query string) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "prepare"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "prepare", func(query string) bool {
				signal.(func(string) bool)(query)
				return f(query)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "prepare", f)
		}
	}
}

func (ptr *QSqlResult) DisconnectPrepare() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "prepare")
	}
}

func (ptr *QSqlResult) Prepare(query string) bool {
	if ptr.Pointer() != nil {
		var queryC *C.char
		if query != "" {
			queryC = C.CString(query)
			defer C.free(unsafe.Pointer(queryC))
		}
		return int8(C.QSqlResult_Prepare(ptr.Pointer(), C.struct_QtSql_PackedString{data: queryC, len: C.longlong(len(query))})) != 0
	}
	return false
}

func (ptr *QSqlResult) PrepareDefault(query string) bool {
	if ptr.Pointer() != nil {
		var queryC *C.char
		if query != "" {
			queryC = C.CString(query)
			defer C.free(unsafe.Pointer(queryC))
		}
		return int8(C.QSqlResult_PrepareDefault(ptr.Pointer(), C.struct_QtSql_PackedString{data: queryC, len: C.longlong(len(query))})) != 0
	}
	return false
}

//export callbackQSqlResult_Reset
func callbackQSqlResult_Reset(ptr unsafe.Pointer, query C.struct_QtSql_PackedString) C.char {
	if signal := qt.GetSignal(ptr, "reset"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string) bool)(cGoUnpackString(query)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QSqlResult) ConnectReset(f func(query string) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "reset"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "reset", func(query string) bool {
				signal.(func(string) bool)(query)
				return f(query)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "reset", f)
		}
	}
}

func (ptr *QSqlResult) DisconnectReset() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "reset")
	}
}

func (ptr *QSqlResult) Reset(query string) bool {
	if ptr.Pointer() != nil {
		var queryC *C.char
		if query != "" {
			queryC = C.CString(query)
			defer C.free(unsafe.Pointer(queryC))
		}
		return int8(C.QSqlResult_Reset(ptr.Pointer(), C.struct_QtSql_PackedString{data: queryC, len: C.longlong(len(query))})) != 0
	}
	return false
}

//export callbackQSqlResult_SavePrepare
func callbackQSqlResult_SavePrepare(ptr unsafe.Pointer, query C.struct_QtSql_PackedString) C.char {
	if signal := qt.GetSignal(ptr, "savePrepare"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string) bool)(cGoUnpackString(query)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlResultFromPointer(ptr).SavePrepareDefault(cGoUnpackString(query)))))
}

func (ptr *QSqlResult) ConnectSavePrepare(f func(query string) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "savePrepare"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "savePrepare", func(query string) bool {
				signal.(func(string) bool)(query)
				return f(query)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "savePrepare", f)
		}
	}
}

func (ptr *QSqlResult) DisconnectSavePrepare() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "savePrepare")
	}
}

func (ptr *QSqlResult) SavePrepare(query string) bool {
	if ptr.Pointer() != nil {
		var queryC *C.char
		if query != "" {
			queryC = C.CString(query)
			defer C.free(unsafe.Pointer(queryC))
		}
		return int8(C.QSqlResult_SavePrepare(ptr.Pointer(), C.struct_QtSql_PackedString{data: queryC, len: C.longlong(len(query))})) != 0
	}
	return false
}

func (ptr *QSqlResult) SavePrepareDefault(query string) bool {
	if ptr.Pointer() != nil {
		var queryC *C.char
		if query != "" {
			queryC = C.CString(query)
			defer C.free(unsafe.Pointer(queryC))
		}
		return int8(C.QSqlResult_SavePrepareDefault(ptr.Pointer(), C.struct_QtSql_PackedString{data: queryC, len: C.longlong(len(query))})) != 0
	}
	return false
}

//export callbackQSqlResult_NumRowsAffected
func callbackQSqlResult_NumRowsAffected(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "numRowsAffected"); signal != nil {
		return C.int(int32(signal.(func() int)()))
	}

	return C.int(int32(0))
}

func (ptr *QSqlResult) ConnectNumRowsAffected(f func() int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "numRowsAffected"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "numRowsAffected", func() int {
				signal.(func() int)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "numRowsAffected", f)
		}
	}
}

func (ptr *QSqlResult) DisconnectNumRowsAffected() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "numRowsAffected")
	}
}

func (ptr *QSqlResult) NumRowsAffected() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSqlResult_NumRowsAffected(ptr.Pointer())))
	}
	return 0
}

//export callbackQSqlResult_Size
func callbackQSqlResult_Size(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "size"); signal != nil {
		return C.int(int32(signal.(func() int)()))
	}

	return C.int(int32(0))
}

func (ptr *QSqlResult) ConnectSize(f func() int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "size"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "size", func() int {
				signal.(func() int)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "size", f)
		}
	}
}

func (ptr *QSqlResult) DisconnectSize() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "size")
	}
}

func (ptr *QSqlResult) Size() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSqlResult_Size(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSqlResult) AddBindValue(val core.QVariant_ITF, paramType QSql__ParamTypeFlag) {
	if ptr.Pointer() != nil {
		C.QSqlResult_AddBindValue(ptr.Pointer(), core.PointerFromQVariant(val), C.longlong(paramType))
	}
}

//export callbackQSqlResult_BindValue2
func callbackQSqlResult_BindValue2(ptr unsafe.Pointer, placeholder C.struct_QtSql_PackedString, val unsafe.Pointer, paramType C.longlong) {
	if signal := qt.GetSignal(ptr, "bindValue2"); signal != nil {
		signal.(func(string, *core.QVariant, QSql__ParamTypeFlag))(cGoUnpackString(placeholder), core.NewQVariantFromPointer(val), QSql__ParamTypeFlag(paramType))
	} else {
		NewQSqlResultFromPointer(ptr).BindValue2Default(cGoUnpackString(placeholder), core.NewQVariantFromPointer(val), QSql__ParamTypeFlag(paramType))
	}
}

func (ptr *QSqlResult) ConnectBindValue2(f func(placeholder string, val *core.QVariant, paramType QSql__ParamTypeFlag)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "bindValue2"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "bindValue2", func(placeholder string, val *core.QVariant, paramType QSql__ParamTypeFlag) {
				signal.(func(string, *core.QVariant, QSql__ParamTypeFlag))(placeholder, val, paramType)
				f(placeholder, val, paramType)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "bindValue2", f)
		}
	}
}

func (ptr *QSqlResult) DisconnectBindValue2() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "bindValue2")
	}
}

func (ptr *QSqlResult) BindValue2(placeholder string, val core.QVariant_ITF, paramType QSql__ParamTypeFlag) {
	if ptr.Pointer() != nil {
		var placeholderC *C.char
		if placeholder != "" {
			placeholderC = C.CString(placeholder)
			defer C.free(unsafe.Pointer(placeholderC))
		}
		C.QSqlResult_BindValue2(ptr.Pointer(), C.struct_QtSql_PackedString{data: placeholderC, len: C.longlong(len(placeholder))}, core.PointerFromQVariant(val), C.longlong(paramType))
	}
}

func (ptr *QSqlResult) BindValue2Default(placeholder string, val core.QVariant_ITF, paramType QSql__ParamTypeFlag) {
	if ptr.Pointer() != nil {
		var placeholderC *C.char
		if placeholder != "" {
			placeholderC = C.CString(placeholder)
			defer C.free(unsafe.Pointer(placeholderC))
		}
		C.QSqlResult_BindValue2Default(ptr.Pointer(), C.struct_QtSql_PackedString{data: placeholderC, len: C.longlong(len(placeholder))}, core.PointerFromQVariant(val), C.longlong(paramType))
	}
}

//export callbackQSqlResult_BindValue
func callbackQSqlResult_BindValue(ptr unsafe.Pointer, index C.int, val unsafe.Pointer, paramType C.longlong) {
	if signal := qt.GetSignal(ptr, "bindValue"); signal != nil {
		signal.(func(int, *core.QVariant, QSql__ParamTypeFlag))(int(int32(index)), core.NewQVariantFromPointer(val), QSql__ParamTypeFlag(paramType))
	} else {
		NewQSqlResultFromPointer(ptr).BindValueDefault(int(int32(index)), core.NewQVariantFromPointer(val), QSql__ParamTypeFlag(paramType))
	}
}

func (ptr *QSqlResult) ConnectBindValue(f func(index int, val *core.QVariant, paramType QSql__ParamTypeFlag)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "bindValue"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "bindValue", func(index int, val *core.QVariant, paramType QSql__ParamTypeFlag) {
				signal.(func(int, *core.QVariant, QSql__ParamTypeFlag))(index, val, paramType)
				f(index, val, paramType)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "bindValue", f)
		}
	}
}

func (ptr *QSqlResult) DisconnectBindValue() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "bindValue")
	}
}

func (ptr *QSqlResult) BindValue(index int, val core.QVariant_ITF, paramType QSql__ParamTypeFlag) {
	if ptr.Pointer() != nil {
		C.QSqlResult_BindValue(ptr.Pointer(), C.int(int32(index)), core.PointerFromQVariant(val), C.longlong(paramType))
	}
}

func (ptr *QSqlResult) BindValueDefault(index int, val core.QVariant_ITF, paramType QSql__ParamTypeFlag) {
	if ptr.Pointer() != nil {
		C.QSqlResult_BindValueDefault(ptr.Pointer(), C.int(int32(index)), core.PointerFromQVariant(val), C.longlong(paramType))
	}
}

func (ptr *QSqlResult) Clear() {
	if ptr.Pointer() != nil {
		C.QSqlResult_Clear(ptr.Pointer())
	}
}

func (ptr *QSqlResult) ResetBindCount() {
	if ptr.Pointer() != nil {
		C.QSqlResult_ResetBindCount(ptr.Pointer())
	}
}

//export callbackQSqlResult_SetActive
func callbackQSqlResult_SetActive(ptr unsafe.Pointer, active C.char) {
	if signal := qt.GetSignal(ptr, "setActive"); signal != nil {
		signal.(func(bool))(int8(active) != 0)
	} else {
		NewQSqlResultFromPointer(ptr).SetActiveDefault(int8(active) != 0)
	}
}

func (ptr *QSqlResult) ConnectSetActive(f func(active bool)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setActive"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setActive", func(active bool) {
				signal.(func(bool))(active)
				f(active)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setActive", f)
		}
	}
}

func (ptr *QSqlResult) DisconnectSetActive() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setActive")
	}
}

func (ptr *QSqlResult) SetActive(active bool) {
	if ptr.Pointer() != nil {
		C.QSqlResult_SetActive(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(active))))
	}
}

func (ptr *QSqlResult) SetActiveDefault(active bool) {
	if ptr.Pointer() != nil {
		C.QSqlResult_SetActiveDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(active))))
	}
}

//export callbackQSqlResult_SetAt
func callbackQSqlResult_SetAt(ptr unsafe.Pointer, index C.int) {
	if signal := qt.GetSignal(ptr, "setAt"); signal != nil {
		signal.(func(int))(int(int32(index)))
	} else {
		NewQSqlResultFromPointer(ptr).SetAtDefault(int(int32(index)))
	}
}

func (ptr *QSqlResult) ConnectSetAt(f func(index int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setAt"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setAt", func(index int) {
				signal.(func(int))(index)
				f(index)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setAt", f)
		}
	}
}

func (ptr *QSqlResult) DisconnectSetAt() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setAt")
	}
}

func (ptr *QSqlResult) SetAt(index int) {
	if ptr.Pointer() != nil {
		C.QSqlResult_SetAt(ptr.Pointer(), C.int(int32(index)))
	}
}

func (ptr *QSqlResult) SetAtDefault(index int) {
	if ptr.Pointer() != nil {
		C.QSqlResult_SetAtDefault(ptr.Pointer(), C.int(int32(index)))
	}
}

//export callbackQSqlResult_SetForwardOnly
func callbackQSqlResult_SetForwardOnly(ptr unsafe.Pointer, forward C.char) {
	if signal := qt.GetSignal(ptr, "setForwardOnly"); signal != nil {
		signal.(func(bool))(int8(forward) != 0)
	} else {
		NewQSqlResultFromPointer(ptr).SetForwardOnlyDefault(int8(forward) != 0)
	}
}

func (ptr *QSqlResult) ConnectSetForwardOnly(f func(forward bool)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setForwardOnly"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setForwardOnly", func(forward bool) {
				signal.(func(bool))(forward)
				f(forward)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setForwardOnly", f)
		}
	}
}

func (ptr *QSqlResult) DisconnectSetForwardOnly() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setForwardOnly")
	}
}

func (ptr *QSqlResult) SetForwardOnly(forward bool) {
	if ptr.Pointer() != nil {
		C.QSqlResult_SetForwardOnly(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(forward))))
	}
}

func (ptr *QSqlResult) SetForwardOnlyDefault(forward bool) {
	if ptr.Pointer() != nil {
		C.QSqlResult_SetForwardOnlyDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(forward))))
	}
}

//export callbackQSqlResult_SetLastError
func callbackQSqlResult_SetLastError(ptr unsafe.Pointer, error unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setLastError"); signal != nil {
		signal.(func(*QSqlError))(NewQSqlErrorFromPointer(error))
	} else {
		NewQSqlResultFromPointer(ptr).SetLastErrorDefault(NewQSqlErrorFromPointer(error))
	}
}

func (ptr *QSqlResult) ConnectSetLastError(f func(error *QSqlError)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setLastError"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setLastError", func(error *QSqlError) {
				signal.(func(*QSqlError))(error)
				f(error)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setLastError", f)
		}
	}
}

func (ptr *QSqlResult) DisconnectSetLastError() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setLastError")
	}
}

func (ptr *QSqlResult) SetLastError(error QSqlError_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlResult_SetLastError(ptr.Pointer(), PointerFromQSqlError(error))
	}
}

func (ptr *QSqlResult) SetLastErrorDefault(error QSqlError_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlResult_SetLastErrorDefault(ptr.Pointer(), PointerFromQSqlError(error))
	}
}

//export callbackQSqlResult_SetQuery
func callbackQSqlResult_SetQuery(ptr unsafe.Pointer, query C.struct_QtSql_PackedString) {
	if signal := qt.GetSignal(ptr, "setQuery"); signal != nil {
		signal.(func(string))(cGoUnpackString(query))
	} else {
		NewQSqlResultFromPointer(ptr).SetQueryDefault(cGoUnpackString(query))
	}
}

func (ptr *QSqlResult) ConnectSetQuery(f func(query string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setQuery"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setQuery", func(query string) {
				signal.(func(string))(query)
				f(query)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setQuery", f)
		}
	}
}

func (ptr *QSqlResult) DisconnectSetQuery() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setQuery")
	}
}

func (ptr *QSqlResult) SetQuery(query string) {
	if ptr.Pointer() != nil {
		var queryC *C.char
		if query != "" {
			queryC = C.CString(query)
			defer C.free(unsafe.Pointer(queryC))
		}
		C.QSqlResult_SetQuery(ptr.Pointer(), C.struct_QtSql_PackedString{data: queryC, len: C.longlong(len(query))})
	}
}

func (ptr *QSqlResult) SetQueryDefault(query string) {
	if ptr.Pointer() != nil {
		var queryC *C.char
		if query != "" {
			queryC = C.CString(query)
			defer C.free(unsafe.Pointer(queryC))
		}
		C.QSqlResult_SetQueryDefault(ptr.Pointer(), C.struct_QtSql_PackedString{data: queryC, len: C.longlong(len(query))})
	}
}

//export callbackQSqlResult_SetSelect
func callbackQSqlResult_SetSelect(ptr unsafe.Pointer, sele C.char) {
	if signal := qt.GetSignal(ptr, "setSelect"); signal != nil {
		signal.(func(bool))(int8(sele) != 0)
	} else {
		NewQSqlResultFromPointer(ptr).SetSelectDefault(int8(sele) != 0)
	}
}

func (ptr *QSqlResult) ConnectSetSelect(f func(sele bool)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setSelect"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setSelect", func(sele bool) {
				signal.(func(bool))(sele)
				f(sele)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setSelect", f)
		}
	}
}

func (ptr *QSqlResult) DisconnectSetSelect() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setSelect")
	}
}

func (ptr *QSqlResult) SetSelect(sele bool) {
	if ptr.Pointer() != nil {
		C.QSqlResult_SetSelect(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(sele))))
	}
}

func (ptr *QSqlResult) SetSelectDefault(sele bool) {
	if ptr.Pointer() != nil {
		C.QSqlResult_SetSelectDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(sele))))
	}
}

//export callbackQSqlResult_DestroyQSqlResult
func callbackQSqlResult_DestroyQSqlResult(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QSqlResult"); signal != nil {
		signal.(func())()
	} else {
		NewQSqlResultFromPointer(ptr).DestroyQSqlResultDefault()
	}
}

func (ptr *QSqlResult) ConnectDestroyQSqlResult(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QSqlResult"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QSqlResult", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QSqlResult", f)
		}
	}
}

func (ptr *QSqlResult) DisconnectDestroyQSqlResult() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QSqlResult")
	}
}

func (ptr *QSqlResult) DestroyQSqlResult() {
	if ptr.Pointer() != nil {
		C.QSqlResult_DestroyQSqlResult(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QSqlResult) DestroyQSqlResultDefault() {
	if ptr.Pointer() != nil {
		C.QSqlResult_DestroyQSqlResultDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QSqlResult) BindValueType2(placeholder string) QSql__ParamTypeFlag {
	if ptr.Pointer() != nil {
		var placeholderC *C.char
		if placeholder != "" {
			placeholderC = C.CString(placeholder)
			defer C.free(unsafe.Pointer(placeholderC))
		}
		return QSql__ParamTypeFlag(C.QSqlResult_BindValueType2(ptr.Pointer(), C.struct_QtSql_PackedString{data: placeholderC, len: C.longlong(len(placeholder))}))
	}
	return 0
}

func (ptr *QSqlResult) BindValueType(index int) QSql__ParamTypeFlag {
	if ptr.Pointer() != nil {
		return QSql__ParamTypeFlag(C.QSqlResult_BindValueType(ptr.Pointer(), C.int(int32(index))))
	}
	return 0
}

func (ptr *QSqlResult) LastError() *QSqlError {
	if ptr.Pointer() != nil {
		tmpValue := NewQSqlErrorFromPointer(C.QSqlResult_LastError(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QSqlError).DestroyQSqlError)
		return tmpValue
	}
	return nil
}

//export callbackQSqlResult_Record
func callbackQSqlResult_Record(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "record"); signal != nil {
		return PointerFromQSqlRecord(signal.(func() *QSqlRecord)())
	}

	return PointerFromQSqlRecord(NewQSqlResultFromPointer(ptr).RecordDefault())
}

func (ptr *QSqlResult) ConnectRecord(f func() *QSqlRecord) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "record"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "record", func() *QSqlRecord {
				signal.(func() *QSqlRecord)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "record", f)
		}
	}
}

func (ptr *QSqlResult) DisconnectRecord() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "record")
	}
}

func (ptr *QSqlResult) Record() *QSqlRecord {
	if ptr.Pointer() != nil {
		tmpValue := NewQSqlRecordFromPointer(C.QSqlResult_Record(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QSqlRecord).DestroyQSqlRecord)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlResult) RecordDefault() *QSqlRecord {
	if ptr.Pointer() != nil {
		tmpValue := NewQSqlRecordFromPointer(C.QSqlResult_RecordDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QSqlRecord).DestroyQSqlRecord)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlResult) BindingSyntax() QSqlResult__BindingSyntax {
	if ptr.Pointer() != nil {
		return QSqlResult__BindingSyntax(C.QSqlResult_BindingSyntax(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSqlResult) BoundValueName(index int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSqlResult_BoundValueName(ptr.Pointer(), C.int(int32(index))))
	}
	return ""
}

func (ptr *QSqlResult) ExecutedQuery() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSqlResult_ExecutedQuery(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlResult) LastQuery() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSqlResult_LastQuery(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlResult) BoundValue2(placeholder string) *core.QVariant {
	if ptr.Pointer() != nil {
		var placeholderC *C.char
		if placeholder != "" {
			placeholderC = C.CString(placeholder)
			defer C.free(unsafe.Pointer(placeholderC))
		}
		tmpValue := core.NewQVariantFromPointer(C.QSqlResult_BoundValue2(ptr.Pointer(), C.struct_QtSql_PackedString{data: placeholderC, len: C.longlong(len(placeholder))}))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlResult) BoundValue(index int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QSqlResult_BoundValue(ptr.Pointer(), C.int(int32(index))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQSqlResult_Handle
func callbackQSqlResult_Handle(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "handle"); signal != nil {
		return core.PointerFromQVariant(signal.(func() *core.QVariant)())
	}

	return core.PointerFromQVariant(NewQSqlResultFromPointer(ptr).HandleDefault())
}

func (ptr *QSqlResult) ConnectHandle(f func() *core.QVariant) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "handle"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "handle", func() *core.QVariant {
				signal.(func() *core.QVariant)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "handle", f)
		}
	}
}

func (ptr *QSqlResult) DisconnectHandle() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "handle")
	}
}

func (ptr *QSqlResult) Handle() *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QSqlResult_Handle(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlResult) HandleDefault() *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QSqlResult_HandleDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQSqlResult_LastInsertId
func callbackQSqlResult_LastInsertId(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "lastInsertId"); signal != nil {
		return core.PointerFromQVariant(signal.(func() *core.QVariant)())
	}

	return core.PointerFromQVariant(NewQSqlResultFromPointer(ptr).LastInsertIdDefault())
}

func (ptr *QSqlResult) ConnectLastInsertId(f func() *core.QVariant) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "lastInsertId"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "lastInsertId", func() *core.QVariant {
				signal.(func() *core.QVariant)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "lastInsertId", f)
		}
	}
}

func (ptr *QSqlResult) DisconnectLastInsertId() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "lastInsertId")
	}
}

func (ptr *QSqlResult) LastInsertId() *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QSqlResult_LastInsertId(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlResult) LastInsertIdDefault() *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QSqlResult_LastInsertIdDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlResult) BoundValues() []*core.QVariant {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtSql_PackedList) []*core.QVariant {
			out := make([]*core.QVariant, int(l.len))
			tmpList := NewQSqlResultFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__boundValues_atList(i)
			}
			return out
		}(C.QSqlResult_BoundValues(ptr.Pointer()))
	}
	return make([]*core.QVariant, 0)
}

func (ptr *QSqlResult) HasOutValues() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlResult_HasOutValues(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSqlResult) IsActive() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlResult_IsActive(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSqlResult) IsForwardOnly() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlResult_IsForwardOnly(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSqlResult) IsSelect() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlResult_IsSelect(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSqlResult) IsValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlResult_IsValid(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSqlResult) Driver() *QSqlDriver {
	if ptr.Pointer() != nil {
		tmpValue := NewQSqlDriverFromPointer(C.QSqlResult_Driver(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSqlResult) At() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSqlResult_At(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSqlResult) BoundValueCount() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSqlResult_BoundValueCount(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSqlResult) __boundValues_atList(i int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QSqlResult___boundValues_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlResult) __boundValues_setList(i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlResult___boundValues_setList(ptr.Pointer(), core.PointerFromQVariant(i))
	}
}

func (ptr *QSqlResult) __boundValues_newList() unsafe.Pointer {
	return C.QSqlResult___boundValues_newList(ptr.Pointer())
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

func NewQSqlTableModelFromPointer(ptr unsafe.Pointer) (n *QSqlTableModel) {
	n = new(QSqlTableModel)
	n.SetPointer(ptr)
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
	tmpValue := NewQSqlTableModelFromPointer(C.QSqlTableModel_NewQSqlTableModel(core.PointerFromQObject(parent), PointerFromQSqlDatabase(db)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQSqlTableModel_DeleteRowFromTable
func callbackQSqlTableModel_DeleteRowFromTable(ptr unsafe.Pointer, row C.int) C.char {
	if signal := qt.GetSignal(ptr, "deleteRowFromTable"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(row))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlTableModelFromPointer(ptr).DeleteRowFromTableDefault(int(int32(row))))))
}

func (ptr *QSqlTableModel) ConnectDeleteRowFromTable(f func(row int) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "deleteRowFromTable"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "deleteRowFromTable", func(row int) bool {
				signal.(func(int) bool)(row)
				return f(row)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "deleteRowFromTable", f)
		}
	}
}

func (ptr *QSqlTableModel) DisconnectDeleteRowFromTable() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "deleteRowFromTable")
	}
}

func (ptr *QSqlTableModel) DeleteRowFromTable(row int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlTableModel_DeleteRowFromTable(ptr.Pointer(), C.int(int32(row)))) != 0
	}
	return false
}

func (ptr *QSqlTableModel) DeleteRowFromTableDefault(row int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlTableModel_DeleteRowFromTableDefault(ptr.Pointer(), C.int(int32(row)))) != 0
	}
	return false
}

func (ptr *QSqlTableModel) InsertRecord(row int, record QSqlRecord_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlTableModel_InsertRecord(ptr.Pointer(), C.int(int32(row)), PointerFromQSqlRecord(record))) != 0
	}
	return false
}

//export callbackQSqlTableModel_InsertRowIntoTable
func callbackQSqlTableModel_InsertRowIntoTable(ptr unsafe.Pointer, values unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "insertRowIntoTable"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QSqlRecord) bool)(NewQSqlRecordFromPointer(values)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlTableModelFromPointer(ptr).InsertRowIntoTableDefault(NewQSqlRecordFromPointer(values)))))
}

func (ptr *QSqlTableModel) ConnectInsertRowIntoTable(f func(values *QSqlRecord) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "insertRowIntoTable"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "insertRowIntoTable", func(values *QSqlRecord) bool {
				signal.(func(*QSqlRecord) bool)(values)
				return f(values)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "insertRowIntoTable", f)
		}
	}
}

func (ptr *QSqlTableModel) DisconnectInsertRowIntoTable() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "insertRowIntoTable")
	}
}

func (ptr *QSqlTableModel) InsertRowIntoTable(values QSqlRecord_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlTableModel_InsertRowIntoTable(ptr.Pointer(), PointerFromQSqlRecord(values))) != 0
	}
	return false
}

func (ptr *QSqlTableModel) InsertRowIntoTableDefault(values QSqlRecord_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlTableModel_InsertRowIntoTableDefault(ptr.Pointer(), PointerFromQSqlRecord(values))) != 0
	}
	return false
}

//export callbackQSqlTableModel_Select
func callbackQSqlTableModel_Select(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "select"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlTableModelFromPointer(ptr).SelectDefault())))
}

func (ptr *QSqlTableModel) ConnectSelect(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "select"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "select", func() bool {
				signal.(func() bool)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "select", f)
		}
	}
}

func (ptr *QSqlTableModel) DisconnectSelect() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "select")
	}
}

func (ptr *QSqlTableModel) Select() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlTableModel_Select(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSqlTableModel) SelectDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlTableModel_SelectDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQSqlTableModel_SelectRow
func callbackQSqlTableModel_SelectRow(ptr unsafe.Pointer, row C.int) C.char {
	if signal := qt.GetSignal(ptr, "selectRow"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(row))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlTableModelFromPointer(ptr).SelectRowDefault(int(int32(row))))))
}

func (ptr *QSqlTableModel) ConnectSelectRow(f func(row int) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "selectRow"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "selectRow", func(row int) bool {
				signal.(func(int) bool)(row)
				return f(row)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "selectRow", f)
		}
	}
}

func (ptr *QSqlTableModel) DisconnectSelectRow() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "selectRow")
	}
}

func (ptr *QSqlTableModel) SelectRow(row int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlTableModel_SelectRow(ptr.Pointer(), C.int(int32(row)))) != 0
	}
	return false
}

func (ptr *QSqlTableModel) SelectRowDefault(row int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlTableModel_SelectRowDefault(ptr.Pointer(), C.int(int32(row)))) != 0
	}
	return false
}

func (ptr *QSqlTableModel) SetRecord(row int, values QSqlRecord_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlTableModel_SetRecord(ptr.Pointer(), C.int(int32(row)), PointerFromQSqlRecord(values))) != 0
	}
	return false
}

//export callbackQSqlTableModel_SubmitAll
func callbackQSqlTableModel_SubmitAll(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "submitAll"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlTableModelFromPointer(ptr).SubmitAllDefault())))
}

func (ptr *QSqlTableModel) ConnectSubmitAll(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "submitAll"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "submitAll", func() bool {
				signal.(func() bool)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "submitAll", f)
		}
	}
}

func (ptr *QSqlTableModel) DisconnectSubmitAll() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "submitAll")
	}
}

func (ptr *QSqlTableModel) SubmitAll() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlTableModel_SubmitAll(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSqlTableModel) SubmitAllDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlTableModel_SubmitAllDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQSqlTableModel_UpdateRowInTable
func callbackQSqlTableModel_UpdateRowInTable(ptr unsafe.Pointer, row C.int, values unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "updateRowInTable"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, *QSqlRecord) bool)(int(int32(row)), NewQSqlRecordFromPointer(values)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlTableModelFromPointer(ptr).UpdateRowInTableDefault(int(int32(row)), NewQSqlRecordFromPointer(values)))))
}

func (ptr *QSqlTableModel) ConnectUpdateRowInTable(f func(row int, values *QSqlRecord) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "updateRowInTable"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "updateRowInTable", func(row int, values *QSqlRecord) bool {
				signal.(func(int, *QSqlRecord) bool)(row, values)
				return f(row, values)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "updateRowInTable", f)
		}
	}
}

func (ptr *QSqlTableModel) DisconnectUpdateRowInTable() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "updateRowInTable")
	}
}

func (ptr *QSqlTableModel) UpdateRowInTable(row int, values QSqlRecord_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlTableModel_UpdateRowInTable(ptr.Pointer(), C.int(int32(row)), PointerFromQSqlRecord(values))) != 0
	}
	return false
}

func (ptr *QSqlTableModel) UpdateRowInTableDefault(row int, values QSqlRecord_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlTableModel_UpdateRowInTableDefault(ptr.Pointer(), C.int(int32(row)), PointerFromQSqlRecord(values))) != 0
	}
	return false
}

//export callbackQSqlTableModel_BeforeDelete
func callbackQSqlTableModel_BeforeDelete(ptr unsafe.Pointer, row C.int) {
	if signal := qt.GetSignal(ptr, "beforeDelete"); signal != nil {
		signal.(func(int))(int(int32(row)))
	}

}

func (ptr *QSqlTableModel) ConnectBeforeDelete(f func(row int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "beforeDelete") {
			C.QSqlTableModel_ConnectBeforeDelete(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "beforeDelete"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "beforeDelete", func(row int) {
				signal.(func(int))(row)
				f(row)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "beforeDelete", f)
		}
	}
}

func (ptr *QSqlTableModel) DisconnectBeforeDelete() {
	if ptr.Pointer() != nil {
		C.QSqlTableModel_DisconnectBeforeDelete(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "beforeDelete")
	}
}

func (ptr *QSqlTableModel) BeforeDelete(row int) {
	if ptr.Pointer() != nil {
		C.QSqlTableModel_BeforeDelete(ptr.Pointer(), C.int(int32(row)))
	}
}

//export callbackQSqlTableModel_BeforeInsert
func callbackQSqlTableModel_BeforeInsert(ptr unsafe.Pointer, record unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "beforeInsert"); signal != nil {
		signal.(func(*QSqlRecord))(NewQSqlRecordFromPointer(record))
	}

}

func (ptr *QSqlTableModel) ConnectBeforeInsert(f func(record *QSqlRecord)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "beforeInsert") {
			C.QSqlTableModel_ConnectBeforeInsert(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "beforeInsert"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "beforeInsert", func(record *QSqlRecord) {
				signal.(func(*QSqlRecord))(record)
				f(record)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "beforeInsert", f)
		}
	}
}

func (ptr *QSqlTableModel) DisconnectBeforeInsert() {
	if ptr.Pointer() != nil {
		C.QSqlTableModel_DisconnectBeforeInsert(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "beforeInsert")
	}
}

func (ptr *QSqlTableModel) BeforeInsert(record QSqlRecord_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlTableModel_BeforeInsert(ptr.Pointer(), PointerFromQSqlRecord(record))
	}
}

//export callbackQSqlTableModel_BeforeUpdate
func callbackQSqlTableModel_BeforeUpdate(ptr unsafe.Pointer, row C.int, record unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "beforeUpdate"); signal != nil {
		signal.(func(int, *QSqlRecord))(int(int32(row)), NewQSqlRecordFromPointer(record))
	}

}

func (ptr *QSqlTableModel) ConnectBeforeUpdate(f func(row int, record *QSqlRecord)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "beforeUpdate") {
			C.QSqlTableModel_ConnectBeforeUpdate(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "beforeUpdate"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "beforeUpdate", func(row int, record *QSqlRecord) {
				signal.(func(int, *QSqlRecord))(row, record)
				f(row, record)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "beforeUpdate", f)
		}
	}
}

func (ptr *QSqlTableModel) DisconnectBeforeUpdate() {
	if ptr.Pointer() != nil {
		C.QSqlTableModel_DisconnectBeforeUpdate(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "beforeUpdate")
	}
}

func (ptr *QSqlTableModel) BeforeUpdate(row int, record QSqlRecord_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlTableModel_BeforeUpdate(ptr.Pointer(), C.int(int32(row)), PointerFromQSqlRecord(record))
	}
}

//export callbackQSqlTableModel_PrimeInsert
func callbackQSqlTableModel_PrimeInsert(ptr unsafe.Pointer, row C.int, record unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "primeInsert"); signal != nil {
		signal.(func(int, *QSqlRecord))(int(int32(row)), NewQSqlRecordFromPointer(record))
	}

}

func (ptr *QSqlTableModel) ConnectPrimeInsert(f func(row int, record *QSqlRecord)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "primeInsert") {
			C.QSqlTableModel_ConnectPrimeInsert(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "primeInsert"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "primeInsert", func(row int, record *QSqlRecord) {
				signal.(func(int, *QSqlRecord))(row, record)
				f(row, record)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "primeInsert", f)
		}
	}
}

func (ptr *QSqlTableModel) DisconnectPrimeInsert() {
	if ptr.Pointer() != nil {
		C.QSqlTableModel_DisconnectPrimeInsert(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "primeInsert")
	}
}

func (ptr *QSqlTableModel) PrimeInsert(row int, record QSqlRecord_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlTableModel_PrimeInsert(ptr.Pointer(), C.int(int32(row)), PointerFromQSqlRecord(record))
	}
}

//export callbackQSqlTableModel_RevertAll
func callbackQSqlTableModel_RevertAll(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "revertAll"); signal != nil {
		signal.(func())()
	} else {
		NewQSqlTableModelFromPointer(ptr).RevertAllDefault()
	}
}

func (ptr *QSqlTableModel) ConnectRevertAll(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "revertAll"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "revertAll", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "revertAll", f)
		}
	}
}

func (ptr *QSqlTableModel) DisconnectRevertAll() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "revertAll")
	}
}

func (ptr *QSqlTableModel) RevertAll() {
	if ptr.Pointer() != nil {
		C.QSqlTableModel_RevertAll(ptr.Pointer())
	}
}

func (ptr *QSqlTableModel) RevertAllDefault() {
	if ptr.Pointer() != nil {
		C.QSqlTableModel_RevertAllDefault(ptr.Pointer())
	}
}

//export callbackQSqlTableModel_RevertRow
func callbackQSqlTableModel_RevertRow(ptr unsafe.Pointer, row C.int) {
	if signal := qt.GetSignal(ptr, "revertRow"); signal != nil {
		signal.(func(int))(int(int32(row)))
	} else {
		NewQSqlTableModelFromPointer(ptr).RevertRowDefault(int(int32(row)))
	}
}

func (ptr *QSqlTableModel) ConnectRevertRow(f func(row int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "revertRow"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "revertRow", func(row int) {
				signal.(func(int))(row)
				f(row)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "revertRow", f)
		}
	}
}

func (ptr *QSqlTableModel) DisconnectRevertRow() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "revertRow")
	}
}

func (ptr *QSqlTableModel) RevertRow(row int) {
	if ptr.Pointer() != nil {
		C.QSqlTableModel_RevertRow(ptr.Pointer(), C.int(int32(row)))
	}
}

func (ptr *QSqlTableModel) RevertRowDefault(row int) {
	if ptr.Pointer() != nil {
		C.QSqlTableModel_RevertRowDefault(ptr.Pointer(), C.int(int32(row)))
	}
}

//export callbackQSqlTableModel_SetEditStrategy
func callbackQSqlTableModel_SetEditStrategy(ptr unsafe.Pointer, strategy C.longlong) {
	if signal := qt.GetSignal(ptr, "setEditStrategy"); signal != nil {
		signal.(func(QSqlTableModel__EditStrategy))(QSqlTableModel__EditStrategy(strategy))
	} else {
		NewQSqlTableModelFromPointer(ptr).SetEditStrategyDefault(QSqlTableModel__EditStrategy(strategy))
	}
}

func (ptr *QSqlTableModel) ConnectSetEditStrategy(f func(strategy QSqlTableModel__EditStrategy)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setEditStrategy"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setEditStrategy", func(strategy QSqlTableModel__EditStrategy) {
				signal.(func(QSqlTableModel__EditStrategy))(strategy)
				f(strategy)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setEditStrategy", f)
		}
	}
}

func (ptr *QSqlTableModel) DisconnectSetEditStrategy() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setEditStrategy")
	}
}

func (ptr *QSqlTableModel) SetEditStrategy(strategy QSqlTableModel__EditStrategy) {
	if ptr.Pointer() != nil {
		C.QSqlTableModel_SetEditStrategy(ptr.Pointer(), C.longlong(strategy))
	}
}

func (ptr *QSqlTableModel) SetEditStrategyDefault(strategy QSqlTableModel__EditStrategy) {
	if ptr.Pointer() != nil {
		C.QSqlTableModel_SetEditStrategyDefault(ptr.Pointer(), C.longlong(strategy))
	}
}

//export callbackQSqlTableModel_SetFilter
func callbackQSqlTableModel_SetFilter(ptr unsafe.Pointer, filter C.struct_QtSql_PackedString) {
	if signal := qt.GetSignal(ptr, "setFilter"); signal != nil {
		signal.(func(string))(cGoUnpackString(filter))
	} else {
		NewQSqlTableModelFromPointer(ptr).SetFilterDefault(cGoUnpackString(filter))
	}
}

func (ptr *QSqlTableModel) ConnectSetFilter(f func(filter string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setFilter"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setFilter", func(filter string) {
				signal.(func(string))(filter)
				f(filter)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setFilter", f)
		}
	}
}

func (ptr *QSqlTableModel) DisconnectSetFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setFilter")
	}
}

func (ptr *QSqlTableModel) SetFilter(filter string) {
	if ptr.Pointer() != nil {
		var filterC *C.char
		if filter != "" {
			filterC = C.CString(filter)
			defer C.free(unsafe.Pointer(filterC))
		}
		C.QSqlTableModel_SetFilter(ptr.Pointer(), C.struct_QtSql_PackedString{data: filterC, len: C.longlong(len(filter))})
	}
}

func (ptr *QSqlTableModel) SetFilterDefault(filter string) {
	if ptr.Pointer() != nil {
		var filterC *C.char
		if filter != "" {
			filterC = C.CString(filter)
			defer C.free(unsafe.Pointer(filterC))
		}
		C.QSqlTableModel_SetFilterDefault(ptr.Pointer(), C.struct_QtSql_PackedString{data: filterC, len: C.longlong(len(filter))})
	}
}

func (ptr *QSqlTableModel) SetPrimaryKey(key QSqlIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlTableModel_SetPrimaryKey(ptr.Pointer(), PointerFromQSqlIndex(key))
	}
}

//export callbackQSqlTableModel_SetSort
func callbackQSqlTableModel_SetSort(ptr unsafe.Pointer, column C.int, order C.longlong) {
	if signal := qt.GetSignal(ptr, "setSort"); signal != nil {
		signal.(func(int, core.Qt__SortOrder))(int(int32(column)), core.Qt__SortOrder(order))
	} else {
		NewQSqlTableModelFromPointer(ptr).SetSortDefault(int(int32(column)), core.Qt__SortOrder(order))
	}
}

func (ptr *QSqlTableModel) ConnectSetSort(f func(column int, order core.Qt__SortOrder)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setSort"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setSort", func(column int, order core.Qt__SortOrder) {
				signal.(func(int, core.Qt__SortOrder))(column, order)
				f(column, order)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setSort", f)
		}
	}
}

func (ptr *QSqlTableModel) DisconnectSetSort() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setSort")
	}
}

func (ptr *QSqlTableModel) SetSort(column int, order core.Qt__SortOrder) {
	if ptr.Pointer() != nil {
		C.QSqlTableModel_SetSort(ptr.Pointer(), C.int(int32(column)), C.longlong(order))
	}
}

func (ptr *QSqlTableModel) SetSortDefault(column int, order core.Qt__SortOrder) {
	if ptr.Pointer() != nil {
		C.QSqlTableModel_SetSortDefault(ptr.Pointer(), C.int(int32(column)), C.longlong(order))
	}
}

//export callbackQSqlTableModel_SetTable
func callbackQSqlTableModel_SetTable(ptr unsafe.Pointer, tableName C.struct_QtSql_PackedString) {
	if signal := qt.GetSignal(ptr, "setTable"); signal != nil {
		signal.(func(string))(cGoUnpackString(tableName))
	} else {
		NewQSqlTableModelFromPointer(ptr).SetTableDefault(cGoUnpackString(tableName))
	}
}

func (ptr *QSqlTableModel) ConnectSetTable(f func(tableName string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setTable"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setTable", func(tableName string) {
				signal.(func(string))(tableName)
				f(tableName)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setTable", f)
		}
	}
}

func (ptr *QSqlTableModel) DisconnectSetTable() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setTable")
	}
}

func (ptr *QSqlTableModel) SetTable(tableName string) {
	if ptr.Pointer() != nil {
		var tableNameC *C.char
		if tableName != "" {
			tableNameC = C.CString(tableName)
			defer C.free(unsafe.Pointer(tableNameC))
		}
		C.QSqlTableModel_SetTable(ptr.Pointer(), C.struct_QtSql_PackedString{data: tableNameC, len: C.longlong(len(tableName))})
	}
}

func (ptr *QSqlTableModel) SetTableDefault(tableName string) {
	if ptr.Pointer() != nil {
		var tableNameC *C.char
		if tableName != "" {
			tableNameC = C.CString(tableName)
			defer C.free(unsafe.Pointer(tableNameC))
		}
		C.QSqlTableModel_SetTableDefault(ptr.Pointer(), C.struct_QtSql_PackedString{data: tableNameC, len: C.longlong(len(tableName))})
	}
}

//export callbackQSqlTableModel_DestroyQSqlTableModel
func callbackQSqlTableModel_DestroyQSqlTableModel(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QSqlTableModel"); signal != nil {
		signal.(func())()
	} else {
		NewQSqlTableModelFromPointer(ptr).DestroyQSqlTableModelDefault()
	}
}

func (ptr *QSqlTableModel) ConnectDestroyQSqlTableModel(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QSqlTableModel"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QSqlTableModel", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QSqlTableModel", f)
		}
	}
}

func (ptr *QSqlTableModel) DisconnectDestroyQSqlTableModel() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QSqlTableModel")
	}
}

func (ptr *QSqlTableModel) DestroyQSqlTableModel() {
	if ptr.Pointer() != nil {
		C.QSqlTableModel_DestroyQSqlTableModel(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QSqlTableModel) DestroyQSqlTableModelDefault() {
	if ptr.Pointer() != nil {
		C.QSqlTableModel_DestroyQSqlTableModelDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QSqlTableModel) Database() *QSqlDatabase {
	if ptr.Pointer() != nil {
		tmpValue := NewQSqlDatabaseFromPointer(C.QSqlTableModel_Database(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QSqlDatabase).DestroyQSqlDatabase)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlTableModel) PrimaryKey() *QSqlIndex {
	if ptr.Pointer() != nil {
		tmpValue := NewQSqlIndexFromPointer(C.QSqlTableModel_PrimaryKey(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QSqlIndex).DestroyQSqlIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlTableModel) PrimaryValues(row int) *QSqlRecord {
	if ptr.Pointer() != nil {
		tmpValue := NewQSqlRecordFromPointer(C.QSqlTableModel_PrimaryValues(ptr.Pointer(), C.int(int32(row))))
		runtime.SetFinalizer(tmpValue, (*QSqlRecord).DestroyQSqlRecord)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlTableModel) EditStrategy() QSqlTableModel__EditStrategy {
	if ptr.Pointer() != nil {
		return QSqlTableModel__EditStrategy(C.QSqlTableModel_EditStrategy(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSqlTableModel) Filter() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSqlTableModel_Filter(ptr.Pointer()))
	}
	return ""
}

//export callbackQSqlTableModel_OrderByClause
func callbackQSqlTableModel_OrderByClause(ptr unsafe.Pointer) C.struct_QtSql_PackedString {
	if signal := qt.GetSignal(ptr, "orderByClause"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_QtSql_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewQSqlTableModelFromPointer(ptr).OrderByClauseDefault()
	return C.struct_QtSql_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QSqlTableModel) ConnectOrderByClause(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "orderByClause"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "orderByClause", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "orderByClause", f)
		}
	}
}

func (ptr *QSqlTableModel) DisconnectOrderByClause() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "orderByClause")
	}
}

func (ptr *QSqlTableModel) OrderByClause() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSqlTableModel_OrderByClause(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlTableModel) OrderByClauseDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSqlTableModel_OrderByClauseDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackQSqlTableModel_SelectStatement
func callbackQSqlTableModel_SelectStatement(ptr unsafe.Pointer) C.struct_QtSql_PackedString {
	if signal := qt.GetSignal(ptr, "selectStatement"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_QtSql_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewQSqlTableModelFromPointer(ptr).SelectStatementDefault()
	return C.struct_QtSql_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QSqlTableModel) ConnectSelectStatement(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "selectStatement"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "selectStatement", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "selectStatement", f)
		}
	}
}

func (ptr *QSqlTableModel) DisconnectSelectStatement() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "selectStatement")
	}
}

func (ptr *QSqlTableModel) SelectStatement() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSqlTableModel_SelectStatement(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlTableModel) SelectStatementDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSqlTableModel_SelectStatementDefault(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlTableModel) TableName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSqlTableModel_TableName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlTableModel) IsDirty2() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlTableModel_IsDirty2(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSqlTableModel) IsDirty(index core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSqlTableModel_IsDirty(ptr.Pointer(), core.PointerFromQModelIndex(index))) != 0
	}
	return false
}

func (ptr *QSqlTableModel) FieldIndex(fieldName string) int {
	if ptr.Pointer() != nil {
		var fieldNameC *C.char
		if fieldName != "" {
			fieldNameC = C.CString(fieldName)
			defer C.free(unsafe.Pointer(fieldNameC))
		}
		return int(int32(C.QSqlTableModel_FieldIndex(ptr.Pointer(), C.struct_QtSql_PackedString{data: fieldNameC, len: C.longlong(len(fieldName))})))
	}
	return 0
}
