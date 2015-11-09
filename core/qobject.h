#ifdef __cplusplus
extern "C" {
#endif

void QObject_InstallEventFilter(void* ptr, void* filterObj);
char* QObject_ObjectName(void* ptr);
void QObject_SetObjectName(void* ptr, char* name);
void* QObject_NewQObject(void* parent);
int QObject_BlockSignals(void* ptr, int block);
void QObject_DeleteLater(void* ptr);
void QObject_ConnectDestroyed(void* ptr);
void QObject_DisconnectDestroyed(void* ptr);
void QObject_DumpObjectInfo(void* ptr);
void QObject_DumpObjectTree(void* ptr);
int QObject_Event(void* ptr, void* e);
int QObject_EventFilter(void* ptr, void* watched, void* event);
void* QObject_FindChild(void* ptr, char* name, int options);
int QObject_Inherits(void* ptr, char* className);
int QObject_IsWidgetType(void* ptr);
int QObject_IsWindowType(void* ptr);
void QObject_KillTimer(void* ptr, int id);
void* QObject_MetaObject(void* ptr);
void QObject_MoveToThread(void* ptr, void* targetThread);
void QObject_ConnectObjectNameChanged(void* ptr);
void QObject_DisconnectObjectNameChanged(void* ptr);
void* QObject_Parent(void* ptr);
void* QObject_Property(void* ptr, char* name);
void QObject_RemoveEventFilter(void* ptr, void* obj);
void QObject_SetParent(void* ptr, void* parent);
int QObject_SetProperty(void* ptr, char* name, void* value);
int QObject_StartTimer(void* ptr, int interval, int timerType);
int QObject_SignalsBlocked(void* ptr);
void* QObject_Thread(void* ptr);
char* QObject_QObject_Tr(char* sourceText, char* disambiguation, int n);
void QObject_DestroyQObject(void* ptr);

#ifdef __cplusplus
}
#endif