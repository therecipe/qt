package core

//#include "qprocess.h"
import "C"
import (
	"github.com/therecipe/qt"
	"strings"
	"unsafe"
)

type QProcess struct {
	QIODevice
}

type QProcessITF interface {
	QIODeviceITF
	QProcessPTR() *QProcess
}

func PointerFromQProcess(ptr QProcessITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QProcessPTR().Pointer()
	}
	return nil
}

func QProcessFromPointer(ptr unsafe.Pointer) *QProcess {
	var n = new(QProcess)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QProcess_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QProcess) QProcessPTR() *QProcess {
	return ptr
}

//QProcess::ExitStatus
type QProcess__ExitStatus int

var (
	QProcess__NormalExit = QProcess__ExitStatus(0)
	QProcess__CrashExit  = QProcess__ExitStatus(1)
)

//QProcess::InputChannelMode
type QProcess__InputChannelMode int

var (
	QProcess__ManagedInputChannel   = QProcess__InputChannelMode(0)
	QProcess__ForwardedInputChannel = QProcess__InputChannelMode(1)
)

//QProcess::ProcessChannel
type QProcess__ProcessChannel int

var (
	QProcess__StandardOutput = QProcess__ProcessChannel(0)
	QProcess__StandardError  = QProcess__ProcessChannel(1)
)

//QProcess::ProcessChannelMode
type QProcess__ProcessChannelMode int

var (
	QProcess__SeparateChannels       = QProcess__ProcessChannelMode(0)
	QProcess__MergedChannels         = QProcess__ProcessChannelMode(1)
	QProcess__ForwardedChannels      = QProcess__ProcessChannelMode(2)
	QProcess__ForwardedOutputChannel = QProcess__ProcessChannelMode(3)
	QProcess__ForwardedErrorChannel  = QProcess__ProcessChannelMode(4)
)

//QProcess::ProcessError
type QProcess__ProcessError int

var (
	QProcess__FailedToStart = QProcess__ProcessError(0)
	QProcess__Crashed       = QProcess__ProcessError(1)
	QProcess__Timedout      = QProcess__ProcessError(2)
	QProcess__ReadError     = QProcess__ProcessError(3)
	QProcess__WriteError    = QProcess__ProcessError(4)
	QProcess__UnknownError  = QProcess__ProcessError(5)
)

//QProcess::ProcessState
type QProcess__ProcessState int

var (
	QProcess__NotRunning = QProcess__ProcessState(0)
	QProcess__Starting   = QProcess__ProcessState(1)
	QProcess__Running    = QProcess__ProcessState(2)
)

func NewQProcess(parent QObjectITF) *QProcess {
	return QProcessFromPointer(unsafe.Pointer(C.QProcess_NewQProcess(C.QtObjectPtr(PointerFromQObject(parent)))))
}

func (ptr *QProcess) Arguments() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QProcess_Arguments(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func (ptr *QProcess) AtEnd() bool {
	if ptr.Pointer() != nil {
		return C.QProcess_AtEnd(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QProcess) CanReadLine() bool {
	if ptr.Pointer() != nil {
		return C.QProcess_CanReadLine(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QProcess) Close() {
	if ptr.Pointer() != nil {
		C.QProcess_Close(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QProcess) CloseReadChannel(channel QProcess__ProcessChannel) {
	if ptr.Pointer() != nil {
		C.QProcess_CloseReadChannel(C.QtObjectPtr(ptr.Pointer()), C.int(channel))
	}
}

func (ptr *QProcess) CloseWriteChannel() {
	if ptr.Pointer() != nil {
		C.QProcess_CloseWriteChannel(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QProcess) Error() QProcess__ProcessError {
	if ptr.Pointer() != nil {
		return QProcess__ProcessError(C.QProcess_Error(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func QProcess_Execute2(command string) int {
	return int(C.QProcess_QProcess_Execute2(C.CString(command)))
}

func QProcess_Execute(program string, arguments []string) int {
	return int(C.QProcess_QProcess_Execute(C.CString(program), C.CString(strings.Join(arguments, "|"))))
}

func (ptr *QProcess) ExitCode() int {
	if ptr.Pointer() != nil {
		return int(C.QProcess_ExitCode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QProcess) ExitStatus() QProcess__ExitStatus {
	if ptr.Pointer() != nil {
		return QProcess__ExitStatus(C.QProcess_ExitStatus(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QProcess) ConnectFinished(f func(exitCode int, exitStatus QProcess__ExitStatus)) {
	if ptr.Pointer() != nil {
		C.QProcess_ConnectFinished(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "finished", f)
	}
}

func (ptr *QProcess) DisconnectFinished() {
	if ptr.Pointer() != nil {
		C.QProcess_DisconnectFinished(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "finished")
	}
}

//export callbackQProcessFinished
func callbackQProcessFinished(ptrName *C.char, exitCode C.int, exitStatus C.int) {
	qt.GetSignal(C.GoString(ptrName), "finished").(func(int, QProcess__ExitStatus))(int(exitCode), QProcess__ExitStatus(exitStatus))
}

func (ptr *QProcess) InputChannelMode() QProcess__InputChannelMode {
	if ptr.Pointer() != nil {
		return QProcess__InputChannelMode(C.QProcess_InputChannelMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QProcess) IsSequential() bool {
	if ptr.Pointer() != nil {
		return C.QProcess_IsSequential(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QProcess) Kill() {
	if ptr.Pointer() != nil {
		C.QProcess_Kill(C.QtObjectPtr(ptr.Pointer()))
	}
}

func QProcess_NullDevice() string {
	return C.GoString(C.QProcess_QProcess_NullDevice())
}

func (ptr *QProcess) Open(mode QIODevice__OpenModeFlag) bool {
	if ptr.Pointer() != nil {
		return C.QProcess_Open(C.QtObjectPtr(ptr.Pointer()), C.int(mode)) != 0
	}
	return false
}

func (ptr *QProcess) ProcessChannelMode() QProcess__ProcessChannelMode {
	if ptr.Pointer() != nil {
		return QProcess__ProcessChannelMode(C.QProcess_ProcessChannelMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QProcess) Program() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QProcess_Program(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QProcess) ReadChannel() QProcess__ProcessChannel {
	if ptr.Pointer() != nil {
		return QProcess__ProcessChannel(C.QProcess_ReadChannel(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QProcess) ConnectReadyReadStandardError(f func()) {
	if ptr.Pointer() != nil {
		C.QProcess_ConnectReadyReadStandardError(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "readyReadStandardError", f)
	}
}

func (ptr *QProcess) DisconnectReadyReadStandardError() {
	if ptr.Pointer() != nil {
		C.QProcess_DisconnectReadyReadStandardError(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "readyReadStandardError")
	}
}

//export callbackQProcessReadyReadStandardError
func callbackQProcessReadyReadStandardError(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "readyReadStandardError").(func())()
}

func (ptr *QProcess) ConnectReadyReadStandardOutput(f func()) {
	if ptr.Pointer() != nil {
		C.QProcess_ConnectReadyReadStandardOutput(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "readyReadStandardOutput", f)
	}
}

func (ptr *QProcess) DisconnectReadyReadStandardOutput() {
	if ptr.Pointer() != nil {
		C.QProcess_DisconnectReadyReadStandardOutput(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "readyReadStandardOutput")
	}
}

//export callbackQProcessReadyReadStandardOutput
func callbackQProcessReadyReadStandardOutput(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "readyReadStandardOutput").(func())()
}

func (ptr *QProcess) SetArguments(arguments []string) {
	if ptr.Pointer() != nil {
		C.QProcess_SetArguments(C.QtObjectPtr(ptr.Pointer()), C.CString(strings.Join(arguments, "|")))
	}
}

func (ptr *QProcess) SetInputChannelMode(mode QProcess__InputChannelMode) {
	if ptr.Pointer() != nil {
		C.QProcess_SetInputChannelMode(C.QtObjectPtr(ptr.Pointer()), C.int(mode))
	}
}

func (ptr *QProcess) SetProcessChannelMode(mode QProcess__ProcessChannelMode) {
	if ptr.Pointer() != nil {
		C.QProcess_SetProcessChannelMode(C.QtObjectPtr(ptr.Pointer()), C.int(mode))
	}
}

func (ptr *QProcess) SetProcessEnvironment(environment QProcessEnvironmentITF) {
	if ptr.Pointer() != nil {
		C.QProcess_SetProcessEnvironment(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQProcessEnvironment(environment)))
	}
}

func (ptr *QProcess) SetProgram(program string) {
	if ptr.Pointer() != nil {
		C.QProcess_SetProgram(C.QtObjectPtr(ptr.Pointer()), C.CString(program))
	}
}

func (ptr *QProcess) SetReadChannel(channel QProcess__ProcessChannel) {
	if ptr.Pointer() != nil {
		C.QProcess_SetReadChannel(C.QtObjectPtr(ptr.Pointer()), C.int(channel))
	}
}

func (ptr *QProcess) SetStandardErrorFile(fileName string, mode QIODevice__OpenModeFlag) {
	if ptr.Pointer() != nil {
		C.QProcess_SetStandardErrorFile(C.QtObjectPtr(ptr.Pointer()), C.CString(fileName), C.int(mode))
	}
}

func (ptr *QProcess) SetStandardInputFile(fileName string) {
	if ptr.Pointer() != nil {
		C.QProcess_SetStandardInputFile(C.QtObjectPtr(ptr.Pointer()), C.CString(fileName))
	}
}

func (ptr *QProcess) SetStandardOutputFile(fileName string, mode QIODevice__OpenModeFlag) {
	if ptr.Pointer() != nil {
		C.QProcess_SetStandardOutputFile(C.QtObjectPtr(ptr.Pointer()), C.CString(fileName), C.int(mode))
	}
}

func (ptr *QProcess) SetStandardOutputProcess(destination QProcessITF) {
	if ptr.Pointer() != nil {
		C.QProcess_SetStandardOutputProcess(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQProcess(destination)))
	}
}

func (ptr *QProcess) SetWorkingDirectory(dir string) {
	if ptr.Pointer() != nil {
		C.QProcess_SetWorkingDirectory(C.QtObjectPtr(ptr.Pointer()), C.CString(dir))
	}
}

func (ptr *QProcess) Start2(mode QIODevice__OpenModeFlag) {
	if ptr.Pointer() != nil {
		C.QProcess_Start2(C.QtObjectPtr(ptr.Pointer()), C.int(mode))
	}
}

func (ptr *QProcess) Start3(command string, mode QIODevice__OpenModeFlag) {
	if ptr.Pointer() != nil {
		C.QProcess_Start3(C.QtObjectPtr(ptr.Pointer()), C.CString(command), C.int(mode))
	}
}

func (ptr *QProcess) Start(program string, arguments []string, mode QIODevice__OpenModeFlag) {
	if ptr.Pointer() != nil {
		C.QProcess_Start(C.QtObjectPtr(ptr.Pointer()), C.CString(program), C.CString(strings.Join(arguments, "|")), C.int(mode))
	}
}

func QProcess_StartDetached2(command string) bool {
	return C.QProcess_QProcess_StartDetached2(C.CString(command)) != 0
}

func (ptr *QProcess) ConnectStarted(f func()) {
	if ptr.Pointer() != nil {
		C.QProcess_ConnectStarted(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "started", f)
	}
}

func (ptr *QProcess) DisconnectStarted() {
	if ptr.Pointer() != nil {
		C.QProcess_DisconnectStarted(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "started")
	}
}

//export callbackQProcessStarted
func callbackQProcessStarted(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "started").(func())()
}

func (ptr *QProcess) ConnectStateChanged(f func(newState QProcess__ProcessState)) {
	if ptr.Pointer() != nil {
		C.QProcess_ConnectStateChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QProcess) DisconnectStateChanged() {
	if ptr.Pointer() != nil {
		C.QProcess_DisconnectStateChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQProcessStateChanged
func callbackQProcessStateChanged(ptrName *C.char, newState C.int) {
	qt.GetSignal(C.GoString(ptrName), "stateChanged").(func(QProcess__ProcessState))(QProcess__ProcessState(newState))
}

func QProcess_SystemEnvironment() []string {
	return strings.Split(C.GoString(C.QProcess_QProcess_SystemEnvironment()), "|")
}

func (ptr *QProcess) Terminate() {
	if ptr.Pointer() != nil {
		C.QProcess_Terminate(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QProcess) WaitForBytesWritten(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QProcess_WaitForBytesWritten(C.QtObjectPtr(ptr.Pointer()), C.int(msecs)) != 0
	}
	return false
}

func (ptr *QProcess) WaitForFinished(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QProcess_WaitForFinished(C.QtObjectPtr(ptr.Pointer()), C.int(msecs)) != 0
	}
	return false
}

func (ptr *QProcess) WaitForReadyRead(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QProcess_WaitForReadyRead(C.QtObjectPtr(ptr.Pointer()), C.int(msecs)) != 0
	}
	return false
}

func (ptr *QProcess) WaitForStarted(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QProcess_WaitForStarted(C.QtObjectPtr(ptr.Pointer()), C.int(msecs)) != 0
	}
	return false
}

func (ptr *QProcess) WorkingDirectory() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QProcess_WorkingDirectory(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QProcess) DestroyQProcess() {
	if ptr.Pointer() != nil {
		C.QProcess_DestroyQProcess(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
