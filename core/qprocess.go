package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"log"
	"strings"
	"unsafe"
)

type QProcess struct {
	QIODevice
}

type QProcess_ITF interface {
	QIODevice_ITF
	QProcess_PTR() *QProcess
}

func PointerFromQProcess(ptr QProcess_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QProcess_PTR().Pointer()
	}
	return nil
}

func NewQProcessFromPointer(ptr unsafe.Pointer) *QProcess {
	var n = new(QProcess)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QProcess_") {
		n.SetObjectName("QProcess_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QProcess) QProcess_PTR() *QProcess {
	return ptr
}

//QProcess::ExitStatus
type QProcess__ExitStatus int64

const (
	QProcess__NormalExit = QProcess__ExitStatus(0)
	QProcess__CrashExit  = QProcess__ExitStatus(1)
)

//QProcess::InputChannelMode
type QProcess__InputChannelMode int64

const (
	QProcess__ManagedInputChannel   = QProcess__InputChannelMode(0)
	QProcess__ForwardedInputChannel = QProcess__InputChannelMode(1)
)

//QProcess::ProcessChannel
type QProcess__ProcessChannel int64

const (
	QProcess__StandardOutput = QProcess__ProcessChannel(0)
	QProcess__StandardError  = QProcess__ProcessChannel(1)
)

//QProcess::ProcessChannelMode
type QProcess__ProcessChannelMode int64

const (
	QProcess__SeparateChannels       = QProcess__ProcessChannelMode(0)
	QProcess__MergedChannels         = QProcess__ProcessChannelMode(1)
	QProcess__ForwardedChannels      = QProcess__ProcessChannelMode(2)
	QProcess__ForwardedOutputChannel = QProcess__ProcessChannelMode(3)
	QProcess__ForwardedErrorChannel  = QProcess__ProcessChannelMode(4)
)

//QProcess::ProcessError
type QProcess__ProcessError int64

const (
	QProcess__FailedToStart = QProcess__ProcessError(0)
	QProcess__Crashed       = QProcess__ProcessError(1)
	QProcess__Timedout      = QProcess__ProcessError(2)
	QProcess__ReadError     = QProcess__ProcessError(3)
	QProcess__WriteError    = QProcess__ProcessError(4)
	QProcess__UnknownError  = QProcess__ProcessError(5)
)

//QProcess::ProcessState
type QProcess__ProcessState int64

const (
	QProcess__NotRunning = QProcess__ProcessState(0)
	QProcess__Starting   = QProcess__ProcessState(1)
	QProcess__Running    = QProcess__ProcessState(2)
)

func NewQProcess(parent QObject_ITF) *QProcess {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProcess::QProcess")
		}
	}()

	return NewQProcessFromPointer(C.QProcess_NewQProcess(PointerFromQObject(parent)))
}

func (ptr *QProcess) Arguments() []string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProcess::arguments")
		}
	}()

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QProcess_Arguments(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QProcess) AtEnd() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProcess::atEnd")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QProcess_AtEnd(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QProcess) CanReadLine() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProcess::canReadLine")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QProcess_CanReadLine(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QProcess) Close() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProcess::close")
		}
	}()

	if ptr.Pointer() != nil {
		C.QProcess_Close(ptr.Pointer())
	}
}

func (ptr *QProcess) CloseReadChannel(channel QProcess__ProcessChannel) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProcess::closeReadChannel")
		}
	}()

	if ptr.Pointer() != nil {
		C.QProcess_CloseReadChannel(ptr.Pointer(), C.int(channel))
	}
}

func (ptr *QProcess) CloseWriteChannel() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProcess::closeWriteChannel")
		}
	}()

	if ptr.Pointer() != nil {
		C.QProcess_CloseWriteChannel(ptr.Pointer())
	}
}

func (ptr *QProcess) Error() QProcess__ProcessError {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProcess::error")
		}
	}()

	if ptr.Pointer() != nil {
		return QProcess__ProcessError(C.QProcess_Error(ptr.Pointer()))
	}
	return 0
}

func QProcess_Execute2(command string) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProcess::execute")
		}
	}()

	return int(C.QProcess_QProcess_Execute2(C.CString(command)))
}

func QProcess_Execute(program string, arguments []string) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProcess::execute")
		}
	}()

	return int(C.QProcess_QProcess_Execute(C.CString(program), C.CString(strings.Join(arguments, ",,,"))))
}

func (ptr *QProcess) ExitCode() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProcess::exitCode")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QProcess_ExitCode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QProcess) ExitStatus() QProcess__ExitStatus {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProcess::exitStatus")
		}
	}()

	if ptr.Pointer() != nil {
		return QProcess__ExitStatus(C.QProcess_ExitStatus(ptr.Pointer()))
	}
	return 0
}

func (ptr *QProcess) ConnectFinished(f func(exitCode int, exitStatus QProcess__ExitStatus)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProcess::finished")
		}
	}()

	if ptr.Pointer() != nil {
		C.QProcess_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "finished", f)
	}
}

func (ptr *QProcess) DisconnectFinished() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProcess::finished")
		}
	}()

	if ptr.Pointer() != nil {
		C.QProcess_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "finished")
	}
}

//export callbackQProcessFinished
func callbackQProcessFinished(ptrName *C.char, exitCode C.int, exitStatus C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProcess::finished")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "finished").(func(int, QProcess__ExitStatus))(int(exitCode), QProcess__ExitStatus(exitStatus))
}

func (ptr *QProcess) InputChannelMode() QProcess__InputChannelMode {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProcess::inputChannelMode")
		}
	}()

	if ptr.Pointer() != nil {
		return QProcess__InputChannelMode(C.QProcess_InputChannelMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QProcess) IsSequential() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProcess::isSequential")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QProcess_IsSequential(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QProcess) Kill() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProcess::kill")
		}
	}()

	if ptr.Pointer() != nil {
		C.QProcess_Kill(ptr.Pointer())
	}
}

func QProcess_NullDevice() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProcess::nullDevice")
		}
	}()

	return C.GoString(C.QProcess_QProcess_NullDevice())
}

func (ptr *QProcess) Open(mode QIODevice__OpenModeFlag) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProcess::open")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QProcess_Open(ptr.Pointer(), C.int(mode)) != 0
	}
	return false
}

func (ptr *QProcess) ProcessChannelMode() QProcess__ProcessChannelMode {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProcess::processChannelMode")
		}
	}()

	if ptr.Pointer() != nil {
		return QProcess__ProcessChannelMode(C.QProcess_ProcessChannelMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QProcess) Program() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProcess::program")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QProcess_Program(ptr.Pointer()))
	}
	return ""
}

func (ptr *QProcess) ReadAllStandardError() *QByteArray {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProcess::readAllStandardError")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QProcess_ReadAllStandardError(ptr.Pointer()))
	}
	return nil
}

func (ptr *QProcess) ReadAllStandardOutput() *QByteArray {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProcess::readAllStandardOutput")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QProcess_ReadAllStandardOutput(ptr.Pointer()))
	}
	return nil
}

func (ptr *QProcess) ReadChannel() QProcess__ProcessChannel {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProcess::readChannel")
		}
	}()

	if ptr.Pointer() != nil {
		return QProcess__ProcessChannel(C.QProcess_ReadChannel(ptr.Pointer()))
	}
	return 0
}

func (ptr *QProcess) ConnectReadyReadStandardError(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProcess::readyReadStandardError")
		}
	}()

	if ptr.Pointer() != nil {
		C.QProcess_ConnectReadyReadStandardError(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "readyReadStandardError", f)
	}
}

func (ptr *QProcess) DisconnectReadyReadStandardError() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProcess::readyReadStandardError")
		}
	}()

	if ptr.Pointer() != nil {
		C.QProcess_DisconnectReadyReadStandardError(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "readyReadStandardError")
	}
}

//export callbackQProcessReadyReadStandardError
func callbackQProcessReadyReadStandardError(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProcess::readyReadStandardError")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "readyReadStandardError").(func())()
}

func (ptr *QProcess) ConnectReadyReadStandardOutput(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProcess::readyReadStandardOutput")
		}
	}()

	if ptr.Pointer() != nil {
		C.QProcess_ConnectReadyReadStandardOutput(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "readyReadStandardOutput", f)
	}
}

func (ptr *QProcess) DisconnectReadyReadStandardOutput() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProcess::readyReadStandardOutput")
		}
	}()

	if ptr.Pointer() != nil {
		C.QProcess_DisconnectReadyReadStandardOutput(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "readyReadStandardOutput")
	}
}

//export callbackQProcessReadyReadStandardOutput
func callbackQProcessReadyReadStandardOutput(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProcess::readyReadStandardOutput")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "readyReadStandardOutput").(func())()
}

func (ptr *QProcess) SetArguments(arguments []string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProcess::setArguments")
		}
	}()

	if ptr.Pointer() != nil {
		C.QProcess_SetArguments(ptr.Pointer(), C.CString(strings.Join(arguments, ",,,")))
	}
}

func (ptr *QProcess) SetInputChannelMode(mode QProcess__InputChannelMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProcess::setInputChannelMode")
		}
	}()

	if ptr.Pointer() != nil {
		C.QProcess_SetInputChannelMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QProcess) SetProcessChannelMode(mode QProcess__ProcessChannelMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProcess::setProcessChannelMode")
		}
	}()

	if ptr.Pointer() != nil {
		C.QProcess_SetProcessChannelMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QProcess) SetProcessEnvironment(environment QProcessEnvironment_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProcess::setProcessEnvironment")
		}
	}()

	if ptr.Pointer() != nil {
		C.QProcess_SetProcessEnvironment(ptr.Pointer(), PointerFromQProcessEnvironment(environment))
	}
}

func (ptr *QProcess) SetProgram(program string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProcess::setProgram")
		}
	}()

	if ptr.Pointer() != nil {
		C.QProcess_SetProgram(ptr.Pointer(), C.CString(program))
	}
}

func (ptr *QProcess) SetReadChannel(channel QProcess__ProcessChannel) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProcess::setReadChannel")
		}
	}()

	if ptr.Pointer() != nil {
		C.QProcess_SetReadChannel(ptr.Pointer(), C.int(channel))
	}
}

func (ptr *QProcess) SetStandardErrorFile(fileName string, mode QIODevice__OpenModeFlag) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProcess::setStandardErrorFile")
		}
	}()

	if ptr.Pointer() != nil {
		C.QProcess_SetStandardErrorFile(ptr.Pointer(), C.CString(fileName), C.int(mode))
	}
}

func (ptr *QProcess) SetStandardInputFile(fileName string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProcess::setStandardInputFile")
		}
	}()

	if ptr.Pointer() != nil {
		C.QProcess_SetStandardInputFile(ptr.Pointer(), C.CString(fileName))
	}
}

func (ptr *QProcess) SetStandardOutputFile(fileName string, mode QIODevice__OpenModeFlag) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProcess::setStandardOutputFile")
		}
	}()

	if ptr.Pointer() != nil {
		C.QProcess_SetStandardOutputFile(ptr.Pointer(), C.CString(fileName), C.int(mode))
	}
}

func (ptr *QProcess) SetStandardOutputProcess(destination QProcess_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProcess::setStandardOutputProcess")
		}
	}()

	if ptr.Pointer() != nil {
		C.QProcess_SetStandardOutputProcess(ptr.Pointer(), PointerFromQProcess(destination))
	}
}

func (ptr *QProcess) SetWorkingDirectory(dir string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProcess::setWorkingDirectory")
		}
	}()

	if ptr.Pointer() != nil {
		C.QProcess_SetWorkingDirectory(ptr.Pointer(), C.CString(dir))
	}
}

func (ptr *QProcess) Start2(mode QIODevice__OpenModeFlag) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProcess::start")
		}
	}()

	if ptr.Pointer() != nil {
		C.QProcess_Start2(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QProcess) Start3(command string, mode QIODevice__OpenModeFlag) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProcess::start")
		}
	}()

	if ptr.Pointer() != nil {
		C.QProcess_Start3(ptr.Pointer(), C.CString(command), C.int(mode))
	}
}

func (ptr *QProcess) Start(program string, arguments []string, mode QIODevice__OpenModeFlag) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProcess::start")
		}
	}()

	if ptr.Pointer() != nil {
		C.QProcess_Start(ptr.Pointer(), C.CString(program), C.CString(strings.Join(arguments, ",,,")), C.int(mode))
	}
}

func QProcess_StartDetached2(command string) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProcess::startDetached")
		}
	}()

	return C.QProcess_QProcess_StartDetached2(C.CString(command)) != 0
}

func (ptr *QProcess) ConnectStarted(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProcess::started")
		}
	}()

	if ptr.Pointer() != nil {
		C.QProcess_ConnectStarted(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "started", f)
	}
}

func (ptr *QProcess) DisconnectStarted() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProcess::started")
		}
	}()

	if ptr.Pointer() != nil {
		C.QProcess_DisconnectStarted(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "started")
	}
}

//export callbackQProcessStarted
func callbackQProcessStarted(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProcess::started")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "started").(func())()
}

func (ptr *QProcess) ConnectStateChanged(f func(newState QProcess__ProcessState)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProcess::stateChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QProcess_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QProcess) DisconnectStateChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProcess::stateChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QProcess_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQProcessStateChanged
func callbackQProcessStateChanged(ptrName *C.char, newState C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProcess::stateChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "stateChanged").(func(QProcess__ProcessState))(QProcess__ProcessState(newState))
}

func QProcess_SystemEnvironment() []string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProcess::systemEnvironment")
		}
	}()

	return strings.Split(C.GoString(C.QProcess_QProcess_SystemEnvironment()), ",,,")
}

func (ptr *QProcess) Terminate() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProcess::terminate")
		}
	}()

	if ptr.Pointer() != nil {
		C.QProcess_Terminate(ptr.Pointer())
	}
}

func (ptr *QProcess) WaitForBytesWritten(msecs int) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProcess::waitForBytesWritten")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QProcess_WaitForBytesWritten(ptr.Pointer(), C.int(msecs)) != 0
	}
	return false
}

func (ptr *QProcess) WaitForFinished(msecs int) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProcess::waitForFinished")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QProcess_WaitForFinished(ptr.Pointer(), C.int(msecs)) != 0
	}
	return false
}

func (ptr *QProcess) WaitForReadyRead(msecs int) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProcess::waitForReadyRead")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QProcess_WaitForReadyRead(ptr.Pointer(), C.int(msecs)) != 0
	}
	return false
}

func (ptr *QProcess) WaitForStarted(msecs int) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProcess::waitForStarted")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QProcess_WaitForStarted(ptr.Pointer(), C.int(msecs)) != 0
	}
	return false
}

func (ptr *QProcess) WorkingDirectory() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProcess::workingDirectory")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QProcess_WorkingDirectory(ptr.Pointer()))
	}
	return ""
}

func (ptr *QProcess) DestroyQProcess() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProcess::~QProcess")
		}
	}()

	if ptr.Pointer() != nil {
		C.QProcess_DestroyQProcess(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
