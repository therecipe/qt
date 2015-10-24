#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QProcess_NewQProcess(QtObjectPtr parent);
char* QProcess_Arguments(QtObjectPtr ptr);
int QProcess_AtEnd(QtObjectPtr ptr);
int QProcess_CanReadLine(QtObjectPtr ptr);
void QProcess_Close(QtObjectPtr ptr);
void QProcess_CloseReadChannel(QtObjectPtr ptr, int channel);
void QProcess_CloseWriteChannel(QtObjectPtr ptr);
int QProcess_Error(QtObjectPtr ptr);
int QProcess_QProcess_Execute2(char* command);
int QProcess_QProcess_Execute(char* program, char* arguments);
int QProcess_ExitCode(QtObjectPtr ptr);
int QProcess_ExitStatus(QtObjectPtr ptr);
void QProcess_ConnectFinished(QtObjectPtr ptr);
void QProcess_DisconnectFinished(QtObjectPtr ptr);
int QProcess_InputChannelMode(QtObjectPtr ptr);
int QProcess_IsSequential(QtObjectPtr ptr);
void QProcess_Kill(QtObjectPtr ptr);
char* QProcess_QProcess_NullDevice();
int QProcess_Open(QtObjectPtr ptr, int mode);
int QProcess_ProcessChannelMode(QtObjectPtr ptr);
char* QProcess_Program(QtObjectPtr ptr);
int QProcess_ReadChannel(QtObjectPtr ptr);
void QProcess_ConnectReadyReadStandardError(QtObjectPtr ptr);
void QProcess_DisconnectReadyReadStandardError(QtObjectPtr ptr);
void QProcess_ConnectReadyReadStandardOutput(QtObjectPtr ptr);
void QProcess_DisconnectReadyReadStandardOutput(QtObjectPtr ptr);
void QProcess_SetArguments(QtObjectPtr ptr, char* arguments);
void QProcess_SetInputChannelMode(QtObjectPtr ptr, int mode);
void QProcess_SetProcessChannelMode(QtObjectPtr ptr, int mode);
void QProcess_SetProcessEnvironment(QtObjectPtr ptr, QtObjectPtr environment);
void QProcess_SetProgram(QtObjectPtr ptr, char* program);
void QProcess_SetReadChannel(QtObjectPtr ptr, int channel);
void QProcess_SetStandardErrorFile(QtObjectPtr ptr, char* fileName, int mode);
void QProcess_SetStandardInputFile(QtObjectPtr ptr, char* fileName);
void QProcess_SetStandardOutputFile(QtObjectPtr ptr, char* fileName, int mode);
void QProcess_SetStandardOutputProcess(QtObjectPtr ptr, QtObjectPtr destination);
void QProcess_SetWorkingDirectory(QtObjectPtr ptr, char* dir);
void QProcess_Start2(QtObjectPtr ptr, int mode);
void QProcess_Start3(QtObjectPtr ptr, char* command, int mode);
void QProcess_Start(QtObjectPtr ptr, char* program, char* arguments, int mode);
int QProcess_QProcess_StartDetached2(char* command);
void QProcess_ConnectStarted(QtObjectPtr ptr);
void QProcess_DisconnectStarted(QtObjectPtr ptr);
void QProcess_ConnectStateChanged(QtObjectPtr ptr);
void QProcess_DisconnectStateChanged(QtObjectPtr ptr);
char* QProcess_QProcess_SystemEnvironment();
void QProcess_Terminate(QtObjectPtr ptr);
int QProcess_WaitForBytesWritten(QtObjectPtr ptr, int msecs);
int QProcess_WaitForFinished(QtObjectPtr ptr, int msecs);
int QProcess_WaitForReadyRead(QtObjectPtr ptr, int msecs);
int QProcess_WaitForStarted(QtObjectPtr ptr, int msecs);
char* QProcess_WorkingDirectory(QtObjectPtr ptr);
void QProcess_DestroyQProcess(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif