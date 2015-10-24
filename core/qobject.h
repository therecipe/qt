#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QObject_InstallEventFilter(QtObjectPtr ptr, QtObjectPtr filterObj);
char* QObject_ObjectName(QtObjectPtr ptr);
void QObject_SetObjectName(QtObjectPtr ptr, char* name);
QtObjectPtr QObject_NewQObject(QtObjectPtr parent);
int QObject_BlockSignals(QtObjectPtr ptr, int block);
void QObject_DeleteLater(QtObjectPtr ptr);
void QObject_ConnectDestroyed(QtObjectPtr ptr);
void QObject_DisconnectDestroyed(QtObjectPtr ptr);
void QObject_DumpObjectInfo(QtObjectPtr ptr);
void QObject_DumpObjectTree(QtObjectPtr ptr);
int QObject_Event(QtObjectPtr ptr, QtObjectPtr e);
int QObject_EventFilter(QtObjectPtr ptr, QtObjectPtr watched, QtObjectPtr event);
int QObject_Inherits(QtObjectPtr ptr, char* className);
int QObject_IsWidgetType(QtObjectPtr ptr);
int QObject_IsWindowType(QtObjectPtr ptr);
void QObject_KillTimer(QtObjectPtr ptr, int id);
QtObjectPtr QObject_MetaObject(QtObjectPtr ptr);
void QObject_MoveToThread(QtObjectPtr ptr, QtObjectPtr targetThread);
void QObject_ConnectObjectNameChanged(QtObjectPtr ptr);
void QObject_DisconnectObjectNameChanged(QtObjectPtr ptr);
QtObjectPtr QObject_Parent(QtObjectPtr ptr);
char* QObject_Property(QtObjectPtr ptr, char* name);
void QObject_RemoveEventFilter(QtObjectPtr ptr, QtObjectPtr obj);
void QObject_SetParent(QtObjectPtr ptr, QtObjectPtr parent);
int QObject_SetProperty(QtObjectPtr ptr, char* name, char* value);
int QObject_StartTimer(QtObjectPtr ptr, int interval, int timerType);
int QObject_SignalsBlocked(QtObjectPtr ptr);
QtObjectPtr QObject_Thread(QtObjectPtr ptr);
char* QObject_QObject_Tr(char* sourceText, char* disambiguation, int n);
void QObject_DestroyQObject(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif