package sql

//#include "sql.h"
import "C"
import (
	"log"
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
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlDatabase::QSqlDatabase")
		}
	}()

	return NewQSqlDatabaseFromPointer(C.QSqlDatabase_NewQSqlDatabase())
}

func NewQSqlDatabase2(other QSqlDatabase_ITF) *QSqlDatabase {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlDatabase::QSqlDatabase")
		}
	}()

	return NewQSqlDatabaseFromPointer(C.QSqlDatabase_NewQSqlDatabase2(PointerFromQSqlDatabase(other)))
}

func (ptr *QSqlDatabase) Close() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlDatabase::close")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSqlDatabase_Close(ptr.Pointer())
	}
}

func (ptr *QSqlDatabase) Commit() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlDatabase::commit")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSqlDatabase_Commit(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlDatabase) ConnectOptions() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlDatabase::connectOptions")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlDatabase_ConnectOptions(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlDatabase) ConnectionName() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlDatabase::connectionName")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlDatabase_ConnectionName(ptr.Pointer()))
	}
	return ""
}

func QSqlDatabase_ConnectionNames() []string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlDatabase::connectionNames")
		}
	}()

	return strings.Split(C.GoString(C.QSqlDatabase_QSqlDatabase_ConnectionNames()), ",,,")
}

func QSqlDatabase_Contains(connectionName string) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlDatabase::contains")
		}
	}()

	return C.QSqlDatabase_QSqlDatabase_Contains(C.CString(connectionName)) != 0
}

func (ptr *QSqlDatabase) DatabaseName() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlDatabase::databaseName")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlDatabase_DatabaseName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlDatabase) Driver() *QSqlDriver {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlDatabase::driver")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQSqlDriverFromPointer(C.QSqlDatabase_Driver(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSqlDatabase) DriverName() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlDatabase::driverName")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlDatabase_DriverName(ptr.Pointer()))
	}
	return ""
}

func QSqlDatabase_Drivers() []string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlDatabase::drivers")
		}
	}()

	return strings.Split(C.GoString(C.QSqlDatabase_QSqlDatabase_Drivers()), ",,,")
}

func (ptr *QSqlDatabase) HostName() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlDatabase::hostName")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlDatabase_HostName(ptr.Pointer()))
	}
	return ""
}

func QSqlDatabase_IsDriverAvailable(name string) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlDatabase::isDriverAvailable")
		}
	}()

	return C.QSqlDatabase_QSqlDatabase_IsDriverAvailable(C.CString(name)) != 0
}

func (ptr *QSqlDatabase) IsOpen() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlDatabase::isOpen")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSqlDatabase_IsOpen(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlDatabase) IsOpenError() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlDatabase::isOpenError")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSqlDatabase_IsOpenError(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlDatabase) IsValid() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlDatabase::isValid")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSqlDatabase_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlDatabase) Open() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlDatabase::open")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSqlDatabase_Open(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlDatabase) Open2(user string, password string) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlDatabase::open")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSqlDatabase_Open2(ptr.Pointer(), C.CString(user), C.CString(password)) != 0
	}
	return false
}

func (ptr *QSqlDatabase) Password() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlDatabase::password")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlDatabase_Password(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlDatabase) Port() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlDatabase::port")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QSqlDatabase_Port(ptr.Pointer()))
	}
	return 0
}

func QSqlDatabase_RegisterSqlDriver(name string, creator QSqlDriverCreatorBase_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlDatabase::registerSqlDriver")
		}
	}()

	C.QSqlDatabase_QSqlDatabase_RegisterSqlDriver(C.CString(name), PointerFromQSqlDriverCreatorBase(creator))
}

func QSqlDatabase_RemoveDatabase(connectionName string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlDatabase::removeDatabase")
		}
	}()

	C.QSqlDatabase_QSqlDatabase_RemoveDatabase(C.CString(connectionName))
}

func (ptr *QSqlDatabase) Rollback() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlDatabase::rollback")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSqlDatabase_Rollback(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlDatabase) SetConnectOptions(options string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlDatabase::setConnectOptions")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSqlDatabase_SetConnectOptions(ptr.Pointer(), C.CString(options))
	}
}

func (ptr *QSqlDatabase) SetDatabaseName(name string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlDatabase::setDatabaseName")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSqlDatabase_SetDatabaseName(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QSqlDatabase) SetHostName(host string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlDatabase::setHostName")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSqlDatabase_SetHostName(ptr.Pointer(), C.CString(host))
	}
}

func (ptr *QSqlDatabase) SetPassword(password string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlDatabase::setPassword")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSqlDatabase_SetPassword(ptr.Pointer(), C.CString(password))
	}
}

func (ptr *QSqlDatabase) SetPort(port int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlDatabase::setPort")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSqlDatabase_SetPort(ptr.Pointer(), C.int(port))
	}
}

func (ptr *QSqlDatabase) SetUserName(name string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlDatabase::setUserName")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSqlDatabase_SetUserName(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QSqlDatabase) Transaction() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlDatabase::transaction")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSqlDatabase_Transaction(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlDatabase) UserName() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlDatabase::userName")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlDatabase_UserName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlDatabase) DestroyQSqlDatabase() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlDatabase::~QSqlDatabase")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSqlDatabase_DestroyQSqlDatabase(ptr.Pointer())
	}
}
