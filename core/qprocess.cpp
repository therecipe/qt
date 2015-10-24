#include "qprocess.h"
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QProcessEnvironment>
#include <QIODevice>
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

QtObjectPtr QProcess_NewQProcess(QtObjectPtr parent){
	return new QProcess(static_cast<QObject*>(parent));
}

char* QProcess_Arguments(QtObjectPtr ptr){
	return static_cast<QProcess*>(ptr)->arguments().join("|").toUtf8().data();
}

int QProcess_AtEnd(QtObjectPtr ptr){
	return static_cast<QProcess*>(ptr)->atEnd();
}

int QProcess_CanReadLine(QtObjectPtr ptr){
	return static_cast<QProcess*>(ptr)->canReadLine();
}

void QProcess_Close(QtObjectPtr ptr){
	static_cast<QProcess*>(ptr)->close();
}

void QProcess_CloseReadChannel(QtObjectPtr ptr, int channel){
	static_cast<QProcess*>(ptr)->closeReadChannel(static_cast<QProcess::ProcessChannel>(channel));
}

void QProcess_CloseWriteChannel(QtObjectPtr ptr){
	static_cast<QProcess*>(ptr)->closeWriteChannel();
}

int QProcess_Error(QtObjectPtr ptr){
	return static_cast<QProcess*>(ptr)->error();
}

int QProcess_QProcess_Execute2(char* command){
	return QProcess::execute(QString(command));
}

int QProcess_QProcess_Execute(char* program, char* arguments){
	return QProcess::execute(QString(program), QString(arguments).split("|", QString::SkipEmptyParts));
}

int QProcess_ExitCode(QtObjectPtr ptr){
	return static_cast<QProcess*>(ptr)->exitCode();
}

int QProcess_ExitStatus(QtObjectPtr ptr){
	return static_cast<QProcess*>(ptr)->exitStatus();
}

void QProcess_ConnectFinished(QtObjectPtr ptr){
	QObject::connect(static_cast<QProcess*>(ptr), static_cast<void (QProcess::*)(int, QProcess::ExitStatus)>(&QProcess::finished), static_cast<MyQProcess*>(ptr), static_cast<void (MyQProcess::*)(int, QProcess::ExitStatus)>(&MyQProcess::Signal_Finished));;
}

void QProcess_DisconnectFinished(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QProcess*>(ptr), static_cast<void (QProcess::*)(int, QProcess::ExitStatus)>(&QProcess::finished), static_cast<MyQProcess*>(ptr), static_cast<void (MyQProcess::*)(int, QProcess::ExitStatus)>(&MyQProcess::Signal_Finished));;
}

int QProcess_InputChannelMode(QtObjectPtr ptr){
	return static_cast<QProcess*>(ptr)->inputChannelMode();
}

int QProcess_IsSequential(QtObjectPtr ptr){
	return static_cast<QProcess*>(ptr)->isSequential();
}

void QProcess_Kill(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QProcess*>(ptr), "kill");
}

char* QProcess_QProcess_NullDevice(){
	return QProcess::nullDevice().toUtf8().data();
}

int QProcess_Open(QtObjectPtr ptr, int mode){
	return static_cast<QProcess*>(ptr)->open(static_cast<QIODevice::OpenModeFlag>(mode));
}

int QProcess_ProcessChannelMode(QtObjectPtr ptr){
	return static_cast<QProcess*>(ptr)->processChannelMode();
}

char* QProcess_Program(QtObjectPtr ptr){
	return static_cast<QProcess*>(ptr)->program().toUtf8().data();
}

int QProcess_ReadChannel(QtObjectPtr ptr){
	return static_cast<QProcess*>(ptr)->readChannel();
}

void QProcess_ConnectReadyReadStandardError(QtObjectPtr ptr){
	QObject::connect(static_cast<QProcess*>(ptr), &QProcess::readyReadStandardError, static_cast<MyQProcess*>(ptr), static_cast<void (MyQProcess::*)()>(&MyQProcess::Signal_ReadyReadStandardError));;
}

void QProcess_DisconnectReadyReadStandardError(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QProcess*>(ptr), &QProcess::readyReadStandardError, static_cast<MyQProcess*>(ptr), static_cast<void (MyQProcess::*)()>(&MyQProcess::Signal_ReadyReadStandardError));;
}

void QProcess_ConnectReadyReadStandardOutput(QtObjectPtr ptr){
	QObject::connect(static_cast<QProcess*>(ptr), &QProcess::readyReadStandardOutput, static_cast<MyQProcess*>(ptr), static_cast<void (MyQProcess::*)()>(&MyQProcess::Signal_ReadyReadStandardOutput));;
}

void QProcess_DisconnectReadyReadStandardOutput(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QProcess*>(ptr), &QProcess::readyReadStandardOutput, static_cast<MyQProcess*>(ptr), static_cast<void (MyQProcess::*)()>(&MyQProcess::Signal_ReadyReadStandardOutput));;
}

void QProcess_SetArguments(QtObjectPtr ptr, char* arguments){
	static_cast<QProcess*>(ptr)->setArguments(QString(arguments).split("|", QString::SkipEmptyParts));
}

void QProcess_SetInputChannelMode(QtObjectPtr ptr, int mode){
	static_cast<QProcess*>(ptr)->setInputChannelMode(static_cast<QProcess::InputChannelMode>(mode));
}

void QProcess_SetProcessChannelMode(QtObjectPtr ptr, int mode){
	static_cast<QProcess*>(ptr)->setProcessChannelMode(static_cast<QProcess::ProcessChannelMode>(mode));
}

void QProcess_SetProcessEnvironment(QtObjectPtr ptr, QtObjectPtr environment){
	static_cast<QProcess*>(ptr)->setProcessEnvironment(*static_cast<QProcessEnvironment*>(environment));
}

void QProcess_SetProgram(QtObjectPtr ptr, char* program){
	static_cast<QProcess*>(ptr)->setProgram(QString(program));
}

void QProcess_SetReadChannel(QtObjectPtr ptr, int channel){
	static_cast<QProcess*>(ptr)->setReadChannel(static_cast<QProcess::ProcessChannel>(channel));
}

void QProcess_SetStandardErrorFile(QtObjectPtr ptr, char* fileName, int mode){
	static_cast<QProcess*>(ptr)->setStandardErrorFile(QString(fileName), static_cast<QIODevice::OpenModeFlag>(mode));
}

void QProcess_SetStandardInputFile(QtObjectPtr ptr, char* fileName){
	static_cast<QProcess*>(ptr)->setStandardInputFile(QString(fileName));
}

void QProcess_SetStandardOutputFile(QtObjectPtr ptr, char* fileName, int mode){
	static_cast<QProcess*>(ptr)->setStandardOutputFile(QString(fileName), static_cast<QIODevice::OpenModeFlag>(mode));
}

void QProcess_SetStandardOutputProcess(QtObjectPtr ptr, QtObjectPtr destination){
	static_cast<QProcess*>(ptr)->setStandardOutputProcess(static_cast<QProcess*>(destination));
}

void QProcess_SetWorkingDirectory(QtObjectPtr ptr, char* dir){
	static_cast<QProcess*>(ptr)->setWorkingDirectory(QString(dir));
}

void QProcess_Start2(QtObjectPtr ptr, int mode){
	static_cast<QProcess*>(ptr)->start(static_cast<QIODevice::OpenModeFlag>(mode));
}

void QProcess_Start3(QtObjectPtr ptr, char* command, int mode){
	static_cast<QProcess*>(ptr)->start(QString(command), static_cast<QIODevice::OpenModeFlag>(mode));
}

void QProcess_Start(QtObjectPtr ptr, char* program, char* arguments, int mode){
	static_cast<QProcess*>(ptr)->start(QString(program), QString(arguments).split("|", QString::SkipEmptyParts), static_cast<QIODevice::OpenModeFlag>(mode));
}

int QProcess_QProcess_StartDetached2(char* command){
	return QProcess::startDetached(QString(command));
}

void QProcess_ConnectStarted(QtObjectPtr ptr){
	QObject::connect(static_cast<QProcess*>(ptr), &QProcess::started, static_cast<MyQProcess*>(ptr), static_cast<void (MyQProcess::*)()>(&MyQProcess::Signal_Started));;
}

void QProcess_DisconnectStarted(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QProcess*>(ptr), &QProcess::started, static_cast<MyQProcess*>(ptr), static_cast<void (MyQProcess::*)()>(&MyQProcess::Signal_Started));;
}

void QProcess_ConnectStateChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QProcess*>(ptr), &QProcess::stateChanged, static_cast<MyQProcess*>(ptr), static_cast<void (MyQProcess::*)(QProcess::ProcessState)>(&MyQProcess::Signal_StateChanged));;
}

void QProcess_DisconnectStateChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QProcess*>(ptr), &QProcess::stateChanged, static_cast<MyQProcess*>(ptr), static_cast<void (MyQProcess::*)(QProcess::ProcessState)>(&MyQProcess::Signal_StateChanged));;
}

char* QProcess_QProcess_SystemEnvironment(){
	return QProcess::systemEnvironment().join("|").toUtf8().data();
}

void QProcess_Terminate(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QProcess*>(ptr), "terminate");
}

int QProcess_WaitForBytesWritten(QtObjectPtr ptr, int msecs){
	return static_cast<QProcess*>(ptr)->waitForBytesWritten(msecs);
}

int QProcess_WaitForFinished(QtObjectPtr ptr, int msecs){
	return static_cast<QProcess*>(ptr)->waitForFinished(msecs);
}

int QProcess_WaitForReadyRead(QtObjectPtr ptr, int msecs){
	return static_cast<QProcess*>(ptr)->waitForReadyRead(msecs);
}

int QProcess_WaitForStarted(QtObjectPtr ptr, int msecs){
	return static_cast<QProcess*>(ptr)->waitForStarted(msecs);
}

char* QProcess_WorkingDirectory(QtObjectPtr ptr){
	return static_cast<QProcess*>(ptr)->workingDirectory().toUtf8().data();
}

void QProcess_DestroyQProcess(QtObjectPtr ptr){
	static_cast<QProcess*>(ptr)->~QProcess();
}

