#include "qprocess.h"
#include <QModelIndex>
#include <QByteArray>
#include <QObject>
#include <QIODevice>
#include <QProcessEnvironment>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QMetaObject>
#include <QProcess>
#include "_cgo_export.h"

class MyQProcess: public QProcess {
public:
void Signal_Finished(int exitCode, QProcess::ExitStatus exitStatus){callbackQProcessFinished(this->objectName().toUtf8().data(), exitCode, exitStatus);};
void Signal_ReadyReadStandardError(){callbackQProcessReadyReadStandardError(this->objectName().toUtf8().data());};
void Signal_ReadyReadStandardOutput(){callbackQProcessReadyReadStandardOutput(this->objectName().toUtf8().data());};
void Signal_Started(){callbackQProcessStarted(this->objectName().toUtf8().data());};
void Signal_StateChanged(QProcess::ProcessState newState){callbackQProcessStateChanged(this->objectName().toUtf8().data(), newState);};
};

void* QProcess_NewQProcess(void* parent){
	return new QProcess(static_cast<QObject*>(parent));
}

char* QProcess_Arguments(void* ptr){
	return static_cast<QProcess*>(ptr)->arguments().join("|").toUtf8().data();
}

int QProcess_AtEnd(void* ptr){
	return static_cast<QProcess*>(ptr)->atEnd();
}

int QProcess_CanReadLine(void* ptr){
	return static_cast<QProcess*>(ptr)->canReadLine();
}

void QProcess_Close(void* ptr){
	static_cast<QProcess*>(ptr)->close();
}

void QProcess_CloseReadChannel(void* ptr, int channel){
	static_cast<QProcess*>(ptr)->closeReadChannel(static_cast<QProcess::ProcessChannel>(channel));
}

void QProcess_CloseWriteChannel(void* ptr){
	static_cast<QProcess*>(ptr)->closeWriteChannel();
}

int QProcess_Error(void* ptr){
	return static_cast<QProcess*>(ptr)->error();
}

int QProcess_QProcess_Execute2(char* command){
	return QProcess::execute(QString(command));
}

int QProcess_QProcess_Execute(char* program, char* arguments){
	return QProcess::execute(QString(program), QString(arguments).split("|", QString::SkipEmptyParts));
}

int QProcess_ExitCode(void* ptr){
	return static_cast<QProcess*>(ptr)->exitCode();
}

int QProcess_ExitStatus(void* ptr){
	return static_cast<QProcess*>(ptr)->exitStatus();
}

void QProcess_ConnectFinished(void* ptr){
	QObject::connect(static_cast<QProcess*>(ptr), static_cast<void (QProcess::*)(int, QProcess::ExitStatus)>(&QProcess::finished), static_cast<MyQProcess*>(ptr), static_cast<void (MyQProcess::*)(int, QProcess::ExitStatus)>(&MyQProcess::Signal_Finished));;
}

void QProcess_DisconnectFinished(void* ptr){
	QObject::disconnect(static_cast<QProcess*>(ptr), static_cast<void (QProcess::*)(int, QProcess::ExitStatus)>(&QProcess::finished), static_cast<MyQProcess*>(ptr), static_cast<void (MyQProcess::*)(int, QProcess::ExitStatus)>(&MyQProcess::Signal_Finished));;
}

int QProcess_InputChannelMode(void* ptr){
	return static_cast<QProcess*>(ptr)->inputChannelMode();
}

int QProcess_IsSequential(void* ptr){
	return static_cast<QProcess*>(ptr)->isSequential();
}

void QProcess_Kill(void* ptr){
	QMetaObject::invokeMethod(static_cast<QProcess*>(ptr), "kill");
}

char* QProcess_QProcess_NullDevice(){
	return QProcess::nullDevice().toUtf8().data();
}

int QProcess_Open(void* ptr, int mode){
	return static_cast<QProcess*>(ptr)->open(static_cast<QIODevice::OpenModeFlag>(mode));
}

int QProcess_ProcessChannelMode(void* ptr){
	return static_cast<QProcess*>(ptr)->processChannelMode();
}

char* QProcess_Program(void* ptr){
	return static_cast<QProcess*>(ptr)->program().toUtf8().data();
}

void* QProcess_ReadAllStandardError(void* ptr){
	return new QByteArray(static_cast<QProcess*>(ptr)->readAllStandardError());
}

void* QProcess_ReadAllStandardOutput(void* ptr){
	return new QByteArray(static_cast<QProcess*>(ptr)->readAllStandardOutput());
}

int QProcess_ReadChannel(void* ptr){
	return static_cast<QProcess*>(ptr)->readChannel();
}

void QProcess_ConnectReadyReadStandardError(void* ptr){
	QObject::connect(static_cast<QProcess*>(ptr), &QProcess::readyReadStandardError, static_cast<MyQProcess*>(ptr), static_cast<void (MyQProcess::*)()>(&MyQProcess::Signal_ReadyReadStandardError));;
}

void QProcess_DisconnectReadyReadStandardError(void* ptr){
	QObject::disconnect(static_cast<QProcess*>(ptr), &QProcess::readyReadStandardError, static_cast<MyQProcess*>(ptr), static_cast<void (MyQProcess::*)()>(&MyQProcess::Signal_ReadyReadStandardError));;
}

void QProcess_ConnectReadyReadStandardOutput(void* ptr){
	QObject::connect(static_cast<QProcess*>(ptr), &QProcess::readyReadStandardOutput, static_cast<MyQProcess*>(ptr), static_cast<void (MyQProcess::*)()>(&MyQProcess::Signal_ReadyReadStandardOutput));;
}

void QProcess_DisconnectReadyReadStandardOutput(void* ptr){
	QObject::disconnect(static_cast<QProcess*>(ptr), &QProcess::readyReadStandardOutput, static_cast<MyQProcess*>(ptr), static_cast<void (MyQProcess::*)()>(&MyQProcess::Signal_ReadyReadStandardOutput));;
}

void QProcess_SetArguments(void* ptr, char* arguments){
	static_cast<QProcess*>(ptr)->setArguments(QString(arguments).split("|", QString::SkipEmptyParts));
}

void QProcess_SetInputChannelMode(void* ptr, int mode){
	static_cast<QProcess*>(ptr)->setInputChannelMode(static_cast<QProcess::InputChannelMode>(mode));
}

void QProcess_SetProcessChannelMode(void* ptr, int mode){
	static_cast<QProcess*>(ptr)->setProcessChannelMode(static_cast<QProcess::ProcessChannelMode>(mode));
}

void QProcess_SetProcessEnvironment(void* ptr, void* environment){
	static_cast<QProcess*>(ptr)->setProcessEnvironment(*static_cast<QProcessEnvironment*>(environment));
}

void QProcess_SetProgram(void* ptr, char* program){
	static_cast<QProcess*>(ptr)->setProgram(QString(program));
}

void QProcess_SetReadChannel(void* ptr, int channel){
	static_cast<QProcess*>(ptr)->setReadChannel(static_cast<QProcess::ProcessChannel>(channel));
}

void QProcess_SetStandardErrorFile(void* ptr, char* fileName, int mode){
	static_cast<QProcess*>(ptr)->setStandardErrorFile(QString(fileName), static_cast<QIODevice::OpenModeFlag>(mode));
}

void QProcess_SetStandardInputFile(void* ptr, char* fileName){
	static_cast<QProcess*>(ptr)->setStandardInputFile(QString(fileName));
}

void QProcess_SetStandardOutputFile(void* ptr, char* fileName, int mode){
	static_cast<QProcess*>(ptr)->setStandardOutputFile(QString(fileName), static_cast<QIODevice::OpenModeFlag>(mode));
}

void QProcess_SetStandardOutputProcess(void* ptr, void* destination){
	static_cast<QProcess*>(ptr)->setStandardOutputProcess(static_cast<QProcess*>(destination));
}

void QProcess_SetWorkingDirectory(void* ptr, char* dir){
	static_cast<QProcess*>(ptr)->setWorkingDirectory(QString(dir));
}

void QProcess_Start2(void* ptr, int mode){
	static_cast<QProcess*>(ptr)->start(static_cast<QIODevice::OpenModeFlag>(mode));
}

void QProcess_Start3(void* ptr, char* command, int mode){
	static_cast<QProcess*>(ptr)->start(QString(command), static_cast<QIODevice::OpenModeFlag>(mode));
}

void QProcess_Start(void* ptr, char* program, char* arguments, int mode){
	static_cast<QProcess*>(ptr)->start(QString(program), QString(arguments).split("|", QString::SkipEmptyParts), static_cast<QIODevice::OpenModeFlag>(mode));
}

int QProcess_QProcess_StartDetached2(char* command){
	return QProcess::startDetached(QString(command));
}

void QProcess_ConnectStarted(void* ptr){
	QObject::connect(static_cast<QProcess*>(ptr), &QProcess::started, static_cast<MyQProcess*>(ptr), static_cast<void (MyQProcess::*)()>(&MyQProcess::Signal_Started));;
}

void QProcess_DisconnectStarted(void* ptr){
	QObject::disconnect(static_cast<QProcess*>(ptr), &QProcess::started, static_cast<MyQProcess*>(ptr), static_cast<void (MyQProcess::*)()>(&MyQProcess::Signal_Started));;
}

void QProcess_ConnectStateChanged(void* ptr){
	QObject::connect(static_cast<QProcess*>(ptr), &QProcess::stateChanged, static_cast<MyQProcess*>(ptr), static_cast<void (MyQProcess::*)(QProcess::ProcessState)>(&MyQProcess::Signal_StateChanged));;
}

void QProcess_DisconnectStateChanged(void* ptr){
	QObject::disconnect(static_cast<QProcess*>(ptr), &QProcess::stateChanged, static_cast<MyQProcess*>(ptr), static_cast<void (MyQProcess::*)(QProcess::ProcessState)>(&MyQProcess::Signal_StateChanged));;
}

char* QProcess_QProcess_SystemEnvironment(){
	return QProcess::systemEnvironment().join("|").toUtf8().data();
}

void QProcess_Terminate(void* ptr){
	QMetaObject::invokeMethod(static_cast<QProcess*>(ptr), "terminate");
}

int QProcess_WaitForBytesWritten(void* ptr, int msecs){
	return static_cast<QProcess*>(ptr)->waitForBytesWritten(msecs);
}

int QProcess_WaitForFinished(void* ptr, int msecs){
	return static_cast<QProcess*>(ptr)->waitForFinished(msecs);
}

int QProcess_WaitForReadyRead(void* ptr, int msecs){
	return static_cast<QProcess*>(ptr)->waitForReadyRead(msecs);
}

int QProcess_WaitForStarted(void* ptr, int msecs){
	return static_cast<QProcess*>(ptr)->waitForStarted(msecs);
}

char* QProcess_WorkingDirectory(void* ptr){
	return static_cast<QProcess*>(ptr)->workingDirectory().toUtf8().data();
}

void QProcess_DestroyQProcess(void* ptr){
	static_cast<QProcess*>(ptr)->~QProcess();
}

