package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
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
		n.SetObjectName("QProcess_" + qt.Identifier())
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
	defer qt.Recovering("QProcess::QProcess")

	return NewQProcessFromPointer(C.QProcess_NewQProcess(PointerFromQObject(parent)))
}

func (ptr *QProcess) Arguments() []string {
	defer qt.Recovering("QProcess::arguments")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QProcess_Arguments(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QProcess) AtEnd() bool {
	defer qt.Recovering("QProcess::atEnd")

	if ptr.Pointer() != nil {
		return C.QProcess_AtEnd(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QProcess) BytesAvailable() int64 {
	defer qt.Recovering("QProcess::bytesAvailable")

	if ptr.Pointer() != nil {
		return int64(C.QProcess_BytesAvailable(ptr.Pointer()))
	}
	return 0
}

func (ptr *QProcess) BytesToWrite() int64 {
	defer qt.Recovering("QProcess::bytesToWrite")

	if ptr.Pointer() != nil {
		return int64(C.QProcess_BytesToWrite(ptr.Pointer()))
	}
	return 0
}

func (ptr *QProcess) CanReadLine() bool {
	defer qt.Recovering("QProcess::canReadLine")

	if ptr.Pointer() != nil {
		return C.QProcess_CanReadLine(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QProcess) ConnectClose(f func()) {
	defer qt.Recovering("connect QProcess::close")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "close", f)
	}
}

func (ptr *QProcess) DisconnectClose() {
	defer qt.Recovering("disconnect QProcess::close")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "close")
	}
}

//export callbackQProcessClose
func callbackQProcessClose(ptrName *C.char) bool {
	defer qt.Recovering("callback QProcess::close")

	if signal := qt.GetSignal(C.GoString(ptrName), "close"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QProcess) CloseReadChannel(channel QProcess__ProcessChannel) {
	defer qt.Recovering("QProcess::closeReadChannel")

	if ptr.Pointer() != nil {
		C.QProcess_CloseReadChannel(ptr.Pointer(), C.int(channel))
	}
}

func (ptr *QProcess) CloseWriteChannel() {
	defer qt.Recovering("QProcess::closeWriteChannel")

	if ptr.Pointer() != nil {
		C.QProcess_CloseWriteChannel(ptr.Pointer())
	}
}

func (ptr *QProcess) ConnectError2(f func(error QProcess__ProcessError)) {
	defer qt.Recovering("connect QProcess::error")

	if ptr.Pointer() != nil {
		C.QProcess_ConnectError2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error2", f)
	}
}

func (ptr *QProcess) DisconnectError2() {
	defer qt.Recovering("disconnect QProcess::error")

	if ptr.Pointer() != nil {
		C.QProcess_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error2")
	}
}

//export callbackQProcessError2
func callbackQProcessError2(ptrName *C.char, error C.int) {
	defer qt.Recovering("callback QProcess::error")

	if signal := qt.GetSignal(C.GoString(ptrName), "error2"); signal != nil {
		signal.(func(QProcess__ProcessError))(QProcess__ProcessError(error))
	}

}

func (ptr *QProcess) Error() QProcess__ProcessError {
	defer qt.Recovering("QProcess::error")

	if ptr.Pointer() != nil {
		return QProcess__ProcessError(C.QProcess_Error(ptr.Pointer()))
	}
	return 0
}

func QProcess_Execute2(command string) int {
	defer qt.Recovering("QProcess::execute")

	return int(C.QProcess_QProcess_Execute2(C.CString(command)))
}

func QProcess_Execute(program string, arguments []string) int {
	defer qt.Recovering("QProcess::execute")

	return int(C.QProcess_QProcess_Execute(C.CString(program), C.CString(strings.Join(arguments, ",,,"))))
}

func (ptr *QProcess) ExitCode() int {
	defer qt.Recovering("QProcess::exitCode")

	if ptr.Pointer() != nil {
		return int(C.QProcess_ExitCode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QProcess) ExitStatus() QProcess__ExitStatus {
	defer qt.Recovering("QProcess::exitStatus")

	if ptr.Pointer() != nil {
		return QProcess__ExitStatus(C.QProcess_ExitStatus(ptr.Pointer()))
	}
	return 0
}

func (ptr *QProcess) ConnectFinished(f func(exitCode int, exitStatus QProcess__ExitStatus)) {
	defer qt.Recovering("connect QProcess::finished")

	if ptr.Pointer() != nil {
		C.QProcess_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "finished", f)
	}
}

func (ptr *QProcess) DisconnectFinished() {
	defer qt.Recovering("disconnect QProcess::finished")

	if ptr.Pointer() != nil {
		C.QProcess_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "finished")
	}
}

//export callbackQProcessFinished
func callbackQProcessFinished(ptrName *C.char, exitCode C.int, exitStatus C.int) {
	defer qt.Recovering("callback QProcess::finished")

	if signal := qt.GetSignal(C.GoString(ptrName), "finished"); signal != nil {
		signal.(func(int, QProcess__ExitStatus))(int(exitCode), QProcess__ExitStatus(exitStatus))
	}

}

func (ptr *QProcess) InputChannelMode() QProcess__InputChannelMode {
	defer qt.Recovering("QProcess::inputChannelMode")

	if ptr.Pointer() != nil {
		return QProcess__InputChannelMode(C.QProcess_InputChannelMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QProcess) IsSequential() bool {
	defer qt.Recovering("QProcess::isSequential")

	if ptr.Pointer() != nil {
		return C.QProcess_IsSequential(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QProcess) Kill() {
	defer qt.Recovering("QProcess::kill")

	if ptr.Pointer() != nil {
		C.QProcess_Kill(ptr.Pointer())
	}
}

func QProcess_NullDevice() string {
	defer qt.Recovering("QProcess::nullDevice")

	return C.GoString(C.QProcess_QProcess_NullDevice())
}

func (ptr *QProcess) Open(mode QIODevice__OpenModeFlag) bool {
	defer qt.Recovering("QProcess::open")

	if ptr.Pointer() != nil {
		return C.QProcess_Open(ptr.Pointer(), C.int(mode)) != 0
	}
	return false
}

func (ptr *QProcess) ProcessChannelMode() QProcess__ProcessChannelMode {
	defer qt.Recovering("QProcess::processChannelMode")

	if ptr.Pointer() != nil {
		return QProcess__ProcessChannelMode(C.QProcess_ProcessChannelMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QProcess) ProcessId() int64 {
	defer qt.Recovering("QProcess::processId")

	if ptr.Pointer() != nil {
		return int64(C.QProcess_ProcessId(ptr.Pointer()))
	}
	return 0
}

func (ptr *QProcess) Program() string {
	defer qt.Recovering("QProcess::program")

	if ptr.Pointer() != nil {
		return C.GoString(C.QProcess_Program(ptr.Pointer()))
	}
	return ""
}

func (ptr *QProcess) ReadAllStandardError() *QByteArray {
	defer qt.Recovering("QProcess::readAllStandardError")

	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QProcess_ReadAllStandardError(ptr.Pointer()))
	}
	return nil
}

func (ptr *QProcess) ReadAllStandardOutput() *QByteArray {
	defer qt.Recovering("QProcess::readAllStandardOutput")

	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QProcess_ReadAllStandardOutput(ptr.Pointer()))
	}
	return nil
}

func (ptr *QProcess) ReadChannel() QProcess__ProcessChannel {
	defer qt.Recovering("QProcess::readChannel")

	if ptr.Pointer() != nil {
		return QProcess__ProcessChannel(C.QProcess_ReadChannel(ptr.Pointer()))
	}
	return 0
}

func (ptr *QProcess) ConnectReadyReadStandardError(f func()) {
	defer qt.Recovering("connect QProcess::readyReadStandardError")

	if ptr.Pointer() != nil {
		C.QProcess_ConnectReadyReadStandardError(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "readyReadStandardError", f)
	}
}

func (ptr *QProcess) DisconnectReadyReadStandardError() {
	defer qt.Recovering("disconnect QProcess::readyReadStandardError")

	if ptr.Pointer() != nil {
		C.QProcess_DisconnectReadyReadStandardError(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "readyReadStandardError")
	}
}

//export callbackQProcessReadyReadStandardError
func callbackQProcessReadyReadStandardError(ptrName *C.char) {
	defer qt.Recovering("callback QProcess::readyReadStandardError")

	if signal := qt.GetSignal(C.GoString(ptrName), "readyReadStandardError"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QProcess) ConnectReadyReadStandardOutput(f func()) {
	defer qt.Recovering("connect QProcess::readyReadStandardOutput")

	if ptr.Pointer() != nil {
		C.QProcess_ConnectReadyReadStandardOutput(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "readyReadStandardOutput", f)
	}
}

func (ptr *QProcess) DisconnectReadyReadStandardOutput() {
	defer qt.Recovering("disconnect QProcess::readyReadStandardOutput")

	if ptr.Pointer() != nil {
		C.QProcess_DisconnectReadyReadStandardOutput(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "readyReadStandardOutput")
	}
}

//export callbackQProcessReadyReadStandardOutput
func callbackQProcessReadyReadStandardOutput(ptrName *C.char) {
	defer qt.Recovering("callback QProcess::readyReadStandardOutput")

	if signal := qt.GetSignal(C.GoString(ptrName), "readyReadStandardOutput"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QProcess) SetArguments(arguments []string) {
	defer qt.Recovering("QProcess::setArguments")

	if ptr.Pointer() != nil {
		C.QProcess_SetArguments(ptr.Pointer(), C.CString(strings.Join(arguments, ",,,")))
	}
}

func (ptr *QProcess) SetInputChannelMode(mode QProcess__InputChannelMode) {
	defer qt.Recovering("QProcess::setInputChannelMode")

	if ptr.Pointer() != nil {
		C.QProcess_SetInputChannelMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QProcess) SetProcessChannelMode(mode QProcess__ProcessChannelMode) {
	defer qt.Recovering("QProcess::setProcessChannelMode")

	if ptr.Pointer() != nil {
		C.QProcess_SetProcessChannelMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QProcess) SetProcessEnvironment(environment QProcessEnvironment_ITF) {
	defer qt.Recovering("QProcess::setProcessEnvironment")

	if ptr.Pointer() != nil {
		C.QProcess_SetProcessEnvironment(ptr.Pointer(), PointerFromQProcessEnvironment(environment))
	}
}

func (ptr *QProcess) SetProgram(program string) {
	defer qt.Recovering("QProcess::setProgram")

	if ptr.Pointer() != nil {
		C.QProcess_SetProgram(ptr.Pointer(), C.CString(program))
	}
}

func (ptr *QProcess) SetReadChannel(channel QProcess__ProcessChannel) {
	defer qt.Recovering("QProcess::setReadChannel")

	if ptr.Pointer() != nil {
		C.QProcess_SetReadChannel(ptr.Pointer(), C.int(channel))
	}
}

func (ptr *QProcess) SetStandardErrorFile(fileName string, mode QIODevice__OpenModeFlag) {
	defer qt.Recovering("QProcess::setStandardErrorFile")

	if ptr.Pointer() != nil {
		C.QProcess_SetStandardErrorFile(ptr.Pointer(), C.CString(fileName), C.int(mode))
	}
}

func (ptr *QProcess) SetStandardInputFile(fileName string) {
	defer qt.Recovering("QProcess::setStandardInputFile")

	if ptr.Pointer() != nil {
		C.QProcess_SetStandardInputFile(ptr.Pointer(), C.CString(fileName))
	}
}

func (ptr *QProcess) SetStandardOutputFile(fileName string, mode QIODevice__OpenModeFlag) {
	defer qt.Recovering("QProcess::setStandardOutputFile")

	if ptr.Pointer() != nil {
		C.QProcess_SetStandardOutputFile(ptr.Pointer(), C.CString(fileName), C.int(mode))
	}
}

func (ptr *QProcess) SetStandardOutputProcess(destination QProcess_ITF) {
	defer qt.Recovering("QProcess::setStandardOutputProcess")

	if ptr.Pointer() != nil {
		C.QProcess_SetStandardOutputProcess(ptr.Pointer(), PointerFromQProcess(destination))
	}
}

func (ptr *QProcess) SetWorkingDirectory(dir string) {
	defer qt.Recovering("QProcess::setWorkingDirectory")

	if ptr.Pointer() != nil {
		C.QProcess_SetWorkingDirectory(ptr.Pointer(), C.CString(dir))
	}
}

func (ptr *QProcess) ConnectSetupChildProcess(f func()) {
	defer qt.Recovering("connect QProcess::setupChildProcess")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setupChildProcess", f)
	}
}

func (ptr *QProcess) DisconnectSetupChildProcess() {
	defer qt.Recovering("disconnect QProcess::setupChildProcess")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setupChildProcess")
	}
}

//export callbackQProcessSetupChildProcess
func callbackQProcessSetupChildProcess(ptrName *C.char) bool {
	defer qt.Recovering("callback QProcess::setupChildProcess")

	if signal := qt.GetSignal(C.GoString(ptrName), "setupChildProcess"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QProcess) Start2(mode QIODevice__OpenModeFlag) {
	defer qt.Recovering("QProcess::start")

	if ptr.Pointer() != nil {
		C.QProcess_Start2(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QProcess) Start3(command string, mode QIODevice__OpenModeFlag) {
	defer qt.Recovering("QProcess::start")

	if ptr.Pointer() != nil {
		C.QProcess_Start3(ptr.Pointer(), C.CString(command), C.int(mode))
	}
}

func (ptr *QProcess) Start(program string, arguments []string, mode QIODevice__OpenModeFlag) {
	defer qt.Recovering("QProcess::start")

	if ptr.Pointer() != nil {
		C.QProcess_Start(ptr.Pointer(), C.CString(program), C.CString(strings.Join(arguments, ",,,")), C.int(mode))
	}
}

func QProcess_StartDetached2(command string) bool {
	defer qt.Recovering("QProcess::startDetached")

	return C.QProcess_QProcess_StartDetached2(C.CString(command)) != 0
}

func (ptr *QProcess) ConnectStarted(f func()) {
	defer qt.Recovering("connect QProcess::started")

	if ptr.Pointer() != nil {
		C.QProcess_ConnectStarted(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "started", f)
	}
}

func (ptr *QProcess) DisconnectStarted() {
	defer qt.Recovering("disconnect QProcess::started")

	if ptr.Pointer() != nil {
		C.QProcess_DisconnectStarted(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "started")
	}
}

//export callbackQProcessStarted
func callbackQProcessStarted(ptrName *C.char) {
	defer qt.Recovering("callback QProcess::started")

	if signal := qt.GetSignal(C.GoString(ptrName), "started"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QProcess) State() QProcess__ProcessState {
	defer qt.Recovering("QProcess::state")

	if ptr.Pointer() != nil {
		return QProcess__ProcessState(C.QProcess_State(ptr.Pointer()))
	}
	return 0
}

func (ptr *QProcess) ConnectStateChanged(f func(newState QProcess__ProcessState)) {
	defer qt.Recovering("connect QProcess::stateChanged")

	if ptr.Pointer() != nil {
		C.QProcess_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QProcess) DisconnectStateChanged() {
	defer qt.Recovering("disconnect QProcess::stateChanged")

	if ptr.Pointer() != nil {
		C.QProcess_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQProcessStateChanged
func callbackQProcessStateChanged(ptrName *C.char, newState C.int) {
	defer qt.Recovering("callback QProcess::stateChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "stateChanged"); signal != nil {
		signal.(func(QProcess__ProcessState))(QProcess__ProcessState(newState))
	}

}

func QProcess_SystemEnvironment() []string {
	defer qt.Recovering("QProcess::systemEnvironment")

	return strings.Split(C.GoString(C.QProcess_QProcess_SystemEnvironment()), ",,,")
}

func (ptr *QProcess) Terminate() {
	defer qt.Recovering("QProcess::terminate")

	if ptr.Pointer() != nil {
		C.QProcess_Terminate(ptr.Pointer())
	}
}

func (ptr *QProcess) WaitForBytesWritten(msecs int) bool {
	defer qt.Recovering("QProcess::waitForBytesWritten")

	if ptr.Pointer() != nil {
		return C.QProcess_WaitForBytesWritten(ptr.Pointer(), C.int(msecs)) != 0
	}
	return false
}

func (ptr *QProcess) WaitForFinished(msecs int) bool {
	defer qt.Recovering("QProcess::waitForFinished")

	if ptr.Pointer() != nil {
		return C.QProcess_WaitForFinished(ptr.Pointer(), C.int(msecs)) != 0
	}
	return false
}

func (ptr *QProcess) WaitForReadyRead(msecs int) bool {
	defer qt.Recovering("QProcess::waitForReadyRead")

	if ptr.Pointer() != nil {
		return C.QProcess_WaitForReadyRead(ptr.Pointer(), C.int(msecs)) != 0
	}
	return false
}

func (ptr *QProcess) WaitForStarted(msecs int) bool {
	defer qt.Recovering("QProcess::waitForStarted")

	if ptr.Pointer() != nil {
		return C.QProcess_WaitForStarted(ptr.Pointer(), C.int(msecs)) != 0
	}
	return false
}

func (ptr *QProcess) WorkingDirectory() string {
	defer qt.Recovering("QProcess::workingDirectory")

	if ptr.Pointer() != nil {
		return C.GoString(C.QProcess_WorkingDirectory(ptr.Pointer()))
	}
	return ""
}

func (ptr *QProcess) DestroyQProcess() {
	defer qt.Recovering("QProcess::~QProcess")

	if ptr.Pointer() != nil {
		C.QProcess_DestroyQProcess(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QProcess) ConnectTimerEvent(f func(event *QTimerEvent)) {
	defer qt.Recovering("connect QProcess::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QProcess) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QProcess::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQProcessTimerEvent
func callbackQProcessTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QProcess::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*QTimerEvent))(NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QProcess) ConnectChildEvent(f func(event *QChildEvent)) {
	defer qt.Recovering("connect QProcess::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QProcess) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QProcess::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQProcessChildEvent
func callbackQProcessChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QProcess::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*QChildEvent))(NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QProcess) ConnectCustomEvent(f func(event *QEvent)) {
	defer qt.Recovering("connect QProcess::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QProcess) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QProcess::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQProcessCustomEvent
func callbackQProcessCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QProcess::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*QEvent))(NewQEventFromPointer(event))
		return true
	}
	return false

}
