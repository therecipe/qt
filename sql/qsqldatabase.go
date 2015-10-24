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

type QSqlDatabaseITF interface {
	QSqlDatabasePTR() *QSqlDatabase
}

func (p *QSqlDatabase) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSqlDatabase) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSqlDatabase(ptr QSqlDatabaseITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlDatabasePTR().Pointer()
	}
	return nil
}

func QSqlDatabaseFromPointer(ptr unsafe.Pointer) *QSqlDatabase {
	var n = new(QSqlDatabase)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSqlDatabase) QSqlDatabasePTR() *QSqlDatabase {
	return ptr
}

func NewQSqlDatabase() *QSqlDatabase {
	return QSqlDatabaseFromPointer(unsafe.Pointer(C.QSqlDatabase_NewQSqlDatabase()))
}

func NewQSqlDatabase2(other QSqlDatabaseITF) *QSqlDatabase {
	return QSqlDatabaseFromPointer(unsafe.Pointer(C.QSqlDatabase_NewQSqlDatabase2(C.QtObjectPtr(PointerFromQSqlDatabase(other)))))
}

func (ptr *QSqlDatabase) Close() {
	if ptr.Pointer() != nil {
		C.QSqlDatabase_Close(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QSqlDatabase) Commit() bool {
	if ptr.Pointer() != nil {
		return C.QSqlDatabase_Commit(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSqlDatabase) ConnectOptions() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlDatabase_ConnectOptions(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QSqlDatabase) ConnectionName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlDatabase_ConnectionName(C.QtObjectPtr(ptr.Pointer())))
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
		return C.GoString(C.QSqlDatabase_DatabaseName(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QSqlDatabase) Driver() *QSqlDriver {
	if ptr.Pointer() != nil {
		return QSqlDriverFromPointer(unsafe.Pointer(C.QSqlDatabase_Driver(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QSqlDatabase) DriverName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlDatabase_DriverName(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func QSqlDatabase_Drivers() []string {
	return strings.Split(C.GoString(C.QSqlDatabase_QSqlDatabase_Drivers()), "|")
}

func (ptr *QSqlDatabase) HostName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlDatabase_HostName(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func QSqlDatabase_IsDriverAvailable(name string) bool {
	return C.QSqlDatabase_QSqlDatabase_IsDriverAvailable(C.CString(name)) != 0
}

func (ptr *QSqlDatabase) IsOpen() bool {
	if ptr.Pointer() != nil {
		return C.QSqlDatabase_IsOpen(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSqlDatabase) IsOpenError() bool {
	if ptr.Pointer() != nil {
		return C.QSqlDatabase_IsOpenError(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSqlDatabase) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QSqlDatabase_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSqlDatabase) Open() bool {
	if ptr.Pointer() != nil {
		return C.QSqlDatabase_Open(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSqlDatabase) Open2(user string, password string) bool {
	if ptr.Pointer() != nil {
		return C.QSqlDatabase_Open2(C.QtObjectPtr(ptr.Pointer()), C.CString(user), C.CString(password)) != 0
	}
	return false
}

func (ptr *QSqlDatabase) Password() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlDatabase_Password(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QSqlDatabase) Port() int {
	if ptr.Pointer() != nil {
		return int(C.QSqlDatabase_Port(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func QSqlDatabase_RegisterSqlDriver(name string, creator QSqlDriverCreatorBaseITF) {
	C.QSqlDatabase_QSqlDatabase_RegisterSqlDriver(C.CString(name), C.QtObjectPtr(PointerFromQSqlDriverCreatorBase(creator)))
}

func QSqlDatabase_RemoveDatabase(connectionName string) {
	C.QSqlDatabase_QSqlDatabase_RemoveDatabase(C.CString(connectionName))
}

func (ptr *QSqlDatabase) Rollback() bool {
	if ptr.Pointer() != nil {
		return C.QSqlDatabase_Rollback(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSqlDatabase) SetConnectOptions(options string) {
	if ptr.Pointer() != nil {
		C.QSqlDatabase_SetConnectOptions(C.QtObjectPtr(ptr.Pointer()), C.CString(options))
	}
}

func (ptr *QSqlDatabase) SetDatabaseName(name string) {
	if ptr.Pointer() != nil {
		C.QSqlDatabase_SetDatabaseName(C.QtObjectPtr(ptr.Pointer()), C.CString(name))
	}
}

func (ptr *QSqlDatabase) SetHostName(host string) {
	if ptr.Pointer() != nil {
		C.QSqlDatabase_SetHostName(C.QtObjectPtr(ptr.Pointer()), C.CString(host))
	}
}

func (ptr *QSqlDatabase) SetPassword(password string) {
	if ptr.Pointer() != nil {
		C.QSqlDatabase_SetPassword(C.QtObjectPtr(ptr.Pointer()), C.CString(password))
	}
}

func (ptr *QSqlDatabase) SetPort(port int) {
	if ptr.Pointer() != nil {
		C.QSqlDatabase_SetPort(C.QtObjectPtr(ptr.Pointer()), C.int(port))
	}
}

func (ptr *QSqlDatabase) SetUserName(name string) {
	if ptr.Pointer() != nil {
		C.QSqlDatabase_SetUserName(C.QtObjectPtr(ptr.Pointer()), C.CString(name))
	}
}

func (ptr *QSqlDatabase) Transaction() bool {
	if ptr.Pointer() != nil {
		return C.QSqlDatabase_Transaction(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSqlDatabase) UserName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlDatabase_UserName(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QSqlDatabase) DestroyQSqlDatabase() {
	if ptr.Pointer() != nil {
		C.QSqlDatabase_DestroyQSqlDatabase(C.QtObjectPtr(ptr.Pointer()))
	}
}
