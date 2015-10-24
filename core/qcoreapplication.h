#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

char* QCoreApplication_QCoreApplication_ApplicationName();
char* QCoreApplication_QCoreApplication_ApplicationVersion();
char* QCoreApplication_QCoreApplication_OrganizationDomain();
char* QCoreApplication_QCoreApplication_OrganizationName();
void QCoreApplication_QCoreApplication_SetApplicationName(char* application);
void QCoreApplication_QCoreApplication_SetApplicationVersion(char* version);
void QCoreApplication_QCoreApplication_SetOrganizationDomain(char* orgDomain);
void QCoreApplication_QCoreApplication_SetOrganizationName(char* orgName);
QtObjectPtr QCoreApplication_NewQCoreApplication(int argc, char* argv);
void QCoreApplication_ConnectAboutToQuit(QtObjectPtr ptr);
void QCoreApplication_DisconnectAboutToQuit(QtObjectPtr ptr);
void QCoreApplication_QCoreApplication_AddLibraryPath(char* path);
char* QCoreApplication_QCoreApplication_ApplicationDirPath();
char* QCoreApplication_QCoreApplication_ApplicationFilePath();
char* QCoreApplication_QCoreApplication_Arguments();
int QCoreApplication_QCoreApplication_ClosingDown();
QtObjectPtr QCoreApplication_QCoreApplication_EventDispatcher();
int QCoreApplication_QCoreApplication_Exec();
void QCoreApplication_QCoreApplication_Exit(int returnCode);
void QCoreApplication_QCoreApplication_Flush();
void QCoreApplication_InstallNativeEventFilter(QtObjectPtr ptr, QtObjectPtr filterObj);
int QCoreApplication_QCoreApplication_InstallTranslator(QtObjectPtr translationFile);
QtObjectPtr QCoreApplication_QCoreApplication_Instance();
int QCoreApplication_QCoreApplication_IsQuitLockEnabled();
int QCoreApplication_QCoreApplication_IsSetuidAllowed();
char* QCoreApplication_QCoreApplication_LibraryPaths();
int QCoreApplication_Notify(QtObjectPtr ptr, QtObjectPtr receiver, QtObjectPtr event);
void QCoreApplication_QCoreApplication_PostEvent(QtObjectPtr receiver, QtObjectPtr event, int priority);
void QCoreApplication_QCoreApplication_ProcessEvents(int flags);
void QCoreApplication_QCoreApplication_ProcessEvents2(int flags, int maxtime);
void QCoreApplication_QCoreApplication_Quit();
void QCoreApplication_QCoreApplication_RemoveLibraryPath(char* path);
void QCoreApplication_RemoveNativeEventFilter(QtObjectPtr ptr, QtObjectPtr filterObject);
void QCoreApplication_QCoreApplication_RemovePostedEvents(QtObjectPtr receiver, int eventType);
int QCoreApplication_QCoreApplication_RemoveTranslator(QtObjectPtr translationFile);
int QCoreApplication_QCoreApplication_SendEvent(QtObjectPtr receiver, QtObjectPtr event);
void QCoreApplication_QCoreApplication_SendPostedEvents(QtObjectPtr receiver, int event_type);
void QCoreApplication_QCoreApplication_SetAttribute(int attribute, int on);
void QCoreApplication_QCoreApplication_SetEventDispatcher(QtObjectPtr eventDispatcher);
void QCoreApplication_QCoreApplication_SetLibraryPaths(char* paths);
void QCoreApplication_QCoreApplication_SetQuitLockEnabled(int enabled);
void QCoreApplication_QCoreApplication_SetSetuidAllowed(int allow);
int QCoreApplication_QCoreApplication_StartingUp();
int QCoreApplication_QCoreApplication_TestAttribute(int attribute);
char* QCoreApplication_QCoreApplication_Translate(char* context, char* sourceText, char* disambiguation, int n);
void QCoreApplication_DestroyQCoreApplication(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif