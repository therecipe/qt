#ifdef __cplusplus
extern "C" {
#endif

void* QProcess_NewQProcess(void* parent);
char* QProcess_Arguments(void* ptr);
int QProcess_AtEnd(void* ptr);
int QProcess_CanReadLine(void* ptr);
void QProcess_Close(void* ptr);
void QProcess_CloseReadChannel(void* ptr, int channel);
void QProcess_CloseWriteChannel(void* ptr);
int QProcess_Error(void* ptr);
int QProcess_QProcess_Execute2(char* command);
int QProcess_QProcess_Execute(char* program, char* arguments);
int QProcess_ExitCode(void* ptr);
int QProcess_ExitStatus(void* ptr);
void QProcess_ConnectFinished(void* ptr);
void QProcess_DisconnectFinished(void* ptr);
int QProcess_InputChannelMode(void* ptr);
int QProcess_IsSequential(void* ptr);
void QProcess_Kill(void* ptr);
char* QProcess_QProcess_NullDevice();
int QProcess_Open(void* ptr, int mode);
int QProcess_ProcessChannelMode(void* ptr);
char* QProcess_Program(void* ptr);
void* QProcess_ReadAllStandardError(void* ptr);
void* QProcess_ReadAllStandardOutput(void* ptr);
int QProcess_ReadChannel(void* ptr);
void QProcess_ConnectReadyReadStandardError(void* ptr);
void QProcess_DisconnectReadyReadStandardError(void* ptr);
void QProcess_ConnectReadyReadStandardOutput(void* ptr);
void QProcess_DisconnectReadyReadStandardOutput(void* ptr);
void QProcess_SetArguments(void* ptr, char* arguments);
void QProcess_SetInputChannelMode(void* ptr, int mode);
void QProcess_SetProcessChannelMode(void* ptr, int mode);
void QProcess_SetProcessEnvironment(void* ptr, void* environment);
void QProcess_SetProgram(void* ptr, char* program);
void QProcess_SetReadChannel(void* ptr, int channel);
void QProcess_SetStandardErrorFile(void* ptr, char* fileName, int mode);
void QProcess_SetStandardInputFile(void* ptr, char* fileName);
void QProcess_SetStandardOutputFile(void* ptr, char* fileName, int mode);
void QProcess_SetStandardOutputProcess(void* ptr, void* destination);
void QProcess_SetWorkingDirectory(void* ptr, char* dir);
void QProcess_Start2(void* ptr, int mode);
void QProcess_Start3(void* ptr, char* command, int mode);
void QProcess_Start(void* ptr, char* program, char* arguments, int mode);
int QProcess_QProcess_StartDetached2(char* command);
void QProcess_ConnectStarted(void* ptr);
void QProcess_DisconnectStarted(void* ptr);
void QProcess_ConnectStateChanged(void* ptr);
void QProcess_DisconnectStateChanged(void* ptr);
char* QProcess_QProcess_SystemEnvironment();
void QProcess_Terminate(void* ptr);
int QProcess_WaitForBytesWritten(void* ptr, int msecs);
int QProcess_WaitForFinished(void* ptr, int msecs);
int QProcess_WaitForReadyRead(void* ptr, int msecs);
int QProcess_WaitForStarted(void* ptr, int msecs);
char* QProcess_WorkingDirectory(void* ptr);
void QProcess_DestroyQProcess(void* ptr);

#ifdef __cplusplus
}
#endif