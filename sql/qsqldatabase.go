package sql

//#include "qsqldatabase.h"
import "C"
import (
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

func (ptr *QSqlDatabase) QSqlDatabase_PTR() *QSqlDatabase {
	return ptr
}

func NewQSqlDatabase() *QSqlDatabase {
	return NewQSqlDatabaseFromPointer(C.QSqlDatabase_NewQSqlDatabase())
}

func NewQSqlDatabase2(other QSqlDatabase_ITF) *QSqlDatabase {
	return NewQSqlDatabaseFromPointer(C.QSqlDatabase_NewQSqlDatabase2(PointerFromQSqlDatabase(other)))
}

func (ptr *QSqlDatabase) Close() {
	if ptr.Pointer() != nil {
		C.QSqlDatabase_Close(ptr.Pointer())
	}
}

func (ptr *QSqlDatabase) Commit() bool {
	if ptr.Pointer() != nil {
		return C.QSqlDatabase_Commit(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlDatabase) ConnectOptions() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlDatabase_ConnectOptions(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlDatabase) ConnectionName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlDatabase_ConnectionName(ptr.Pointer()))
	}
	return ""
}

func QSqlDatabase_ConnectionNames() []string {
	return strings.Split(C.GoString(C.QSqlDatabase_QSqlDatabase_ConnectionNames()), "|")
}

func QSqlDatabase_Contains(connectionName string) bool {
	return C.QSqlDatabase_QSqlDatabase_Contains(C.CString(connectionName)) != 0
}

func (ptr *QSqlDatabase) DatabaseName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlDatabase_DatabaseName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlDatabase) Driver() *QSqlDriver {
	if ptr.Pointer() != nil {
		return NewQSqlDriverFromPointer(C.QSqlDatabase_Driver(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSqlDatabase) DriverName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlDatabase_DriverName(ptr.Pointer()))
	}
	return ""
}

func QSqlDatabase_Drivers() []string {
	return strings.Split(C.GoString(C.QSqlDatabase_QSqlDatabase_Drivers()), "|")
}

func (ptr *QSqlDatabase) HostName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlDatabase_HostName(ptr.Pointer()))
	}
	return ""
}

func QSqlDatabase_IsDriverAvailable(name string) bool {
	return C.QSqlDatabase_QSqlDatabase_IsDriverAvailable(C.CString(name)) != 0
}

func (ptr *QSqlDatabase) IsOpen() bool {
	if ptr.Pointer() != nil {
		return C.QSqlDatabase_IsOpen(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlDatabase) IsOpenError() bool {
	if ptr.Pointer() != nil {
		return C.QSqlDatabase_IsOpenError(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlDatabase) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QSqlDatabase_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlDatabase) Open() bool {
	if ptr.Pointer() != nil {
		return C.QSqlDatabase_Open(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlDatabase) Open2(user string, password string) bool {
	if ptr.Pointer() != nil {
		return C.QSqlDatabase_Open2(ptr.Pointer(), C.CString(user), C.CString(password)) != 0
	}
	return false
}

func (ptr *QSqlDatabase) Password() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlDatabase_Password(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlDatabase) Port() int {
	if ptr.Pointer() != nil {
		return int(C.QSqlDatabase_Port(ptr.Pointer()))
	}
	return 0
}

func QSqlDatabase_RegisterSqlDriver(name string, creator QSqlDriverCreatorBase_ITF) {
	C.QSqlDatabase_QSqlDatabase_RegisterSqlDriver(C.CString(name), PointerFromQSqlDriverCreatorBase(creator))
}

func QSqlDatabase_RemoveDatabase(connectionName string) {
	C.QSqlDatabase_QSqlDatabase_RemoveDatabase(C.CString(connectionName))
}

func (ptr *QSqlDatabase) Rollback() bool {
	if ptr.Pointer() != nil {
		return C.QSqlDatabase_Rollback(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlDatabase) SetConnectOptions(options string) {
	if ptr.Pointer() != nil {
		C.QSqlDatabase_SetConnectOptions(ptr.Pointer(), C.CString(options))
	}
}

func (ptr *QSqlDatabase) SetDatabaseName(name string) {
	if ptr.Pointer() != nil {
		C.QSqlDatabase_SetDatabaseName(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QSqlDatabase) SetHostName(host string) {
	if ptr.Pointer() != nil {
		C.QSqlDatabase_SetHostName(ptr.Pointer(), C.CString(host))
	}
}

func (ptr *QSqlDatabase) SetPassword(password string) {
	if ptr.Pointer() != nil {
		C.QSqlDatabase_SetPassword(ptr.Pointer(), C.CString(password))
	}
}

func (ptr *QSqlDatabase) SetPort(port int) {
	if ptr.Pointer() != nil {
		C.QSqlDatabase_SetPort(ptr.Pointer(), C.int(port))
	}
}

func (ptr *QSqlDatabase) SetUserName(name string) {
	if ptr.Pointer() != nil {
		C.QSqlDatabase_SetUserName(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QSqlDatabase) Transaction() bool {
	if ptr.Pointer() != nil {
		return C.QSqlDatabase_Transaction(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlDatabase) UserName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlDatabase_UserName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlDatabase) DestroyQSqlDatabase() {
	if ptr.Pointer() != nil {
		C.QSqlDatabase_DestroyQSqlDatabase(ptr.Pointer())
	}
}
