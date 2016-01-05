#ifdef __cplusplus
extern "C" {
#endif

void* QUiLoader_NewQUiLoader(void* parent);
void QUiLoader_AddPluginPath(void* ptr, char* path);
char* QUiLoader_AvailableLayouts(void* ptr);
char* QUiLoader_AvailableWidgets(void* ptr);
void QUiLoader_ClearPluginPaths(void* ptr);
void* QUiLoader_CreateAction(void* ptr, void* parent, char* name);
void* QUiLoader_CreateActionDefault(void* ptr, void* parent, char* name);
void* QUiLoader_CreateActionGroup(void* ptr, void* parent, char* name);
void* QUiLoader_CreateActionGroupDefault(void* ptr, void* parent, char* name);
void* QUiLoader_CreateLayout(void* ptr, char* className, void* parent, char* name);
void* QUiLoader_CreateLayoutDefault(void* ptr, char* className, void* parent, char* name);
void* QUiLoader_CreateWidget(void* ptr, char* className, void* parent, char* name);
void* QUiLoader_CreateWidgetDefault(void* ptr, char* className, void* parent, char* name);
char* QUiLoader_ErrorString(void* ptr);
int QUiLoader_IsLanguageChangeEnabled(void* ptr);
void* QUiLoader_Load(void* ptr, void* device, void* parentWidget);
char* QUiLoader_PluginPaths(void* ptr);
void QUiLoader_SetLanguageChangeEnabled(void* ptr, int enabled);
void QUiLoader_SetWorkingDirectory(void* ptr, void* dir);
void* QUiLoader_WorkingDirectory(void* ptr);
void QUiLoader_DestroyQUiLoader(void* ptr);
void QUiLoader_TimerEvent(void* ptr, void* event);
void QUiLoader_TimerEventDefault(void* ptr, void* event);
void QUiLoader_ChildEvent(void* ptr, void* event);
void QUiLoader_ChildEventDefault(void* ptr, void* event);
void QUiLoader_CustomEvent(void* ptr, void* event);
void QUiLoader_CustomEventDefault(void* ptr, void* event);

#ifdef __cplusplus
}
#endif