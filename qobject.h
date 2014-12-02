#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

//Public Functions
QtObjectPtr QObject_New_QObject(QtObjectPtr parent);
void QObject_Destroy(QtObjectPtr ptr);
int QObject_BlockSignals_Bool(QtObjectPtr ptr, int block);
int QObject_Disconnect_String_QObject_String(QtObjectPtr ptr, char* signal, QtObjectPtr receiver, char* method);
int QObject_Disconnect_QObject_String(QtObjectPtr ptr, QtObjectPtr receiver, char* method);
void QObject_DumpObjectInfo(QtObjectPtr ptr);
void QObject_DumpObjectTree(QtObjectPtr ptr);
int QObject_Inherits_String(QtObjectPtr ptr, char* className);
void QObject_InstallEventFilter_QObject(QtObjectPtr ptr, QtObjectPtr filterObj);
int QObject_IsWidgetType(QtObjectPtr ptr);
int QObject_IsWindowType(QtObjectPtr ptr);
void QObject_KillTimer_Int(QtObjectPtr ptr, int id);
void QObject_MoveToThread_QThread(QtObjectPtr ptr, QtObjectPtr targetThread);
char* QObject_ObjectName(QtObjectPtr ptr);
QtObjectPtr QObject_Parent(QtObjectPtr ptr);
void QObject_RemoveEventFilter_QObject(QtObjectPtr ptr, QtObjectPtr obj);
void QObject_SetObjectName_String(QtObjectPtr ptr, char* name);
void QObject_SetParent_QObject(QtObjectPtr ptr, QtObjectPtr parent);
int QObject_SignalsBlocked(QtObjectPtr ptr);
int QObject_StartTimer_Int_TimerType(QtObjectPtr ptr, int interval, int timerType);
QtObjectPtr QObject_Thread(QtObjectPtr ptr);
//Public Slots
void QObject_ConnectSlotDeleteLater(QtObjectPtr ptr);
void QObject_DisconnectSlotDeleteLater(QtObjectPtr ptr);
void QObject_DeleteLater(QtObjectPtr ptr);
//Signals
void QObject_ConnectSignalDestroyed(QtObjectPtr ptr);
void QObject_DisconnectSignalDestroyed(QtObjectPtr ptr);
void QObject_ConnectSignalObjectNameChanged(QtObjectPtr ptr);
void QObject_DisconnectSignalObjectNameChanged(QtObjectPtr ptr);
//Static Public Members
int QObject_Disconnect_QObject_String_QObject_String(QtObjectPtr sender, char* signal, QtObjectPtr receiver, char* method);
char* QObject_Tr_String_String_Int(char* sourceText, char* disambiguation, int n);

#ifdef __cplusplus
}
#endif
