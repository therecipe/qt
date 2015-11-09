#ifdef __cplusplus
extern "C" {
#endif

char* QPluginLoader_FileName(void* ptr);
int QPluginLoader_LoadHints(void* ptr);
void QPluginLoader_SetFileName(void* ptr, char* fileName);
void QPluginLoader_SetLoadHints(void* ptr, int loadHints);
void* QPluginLoader_NewQPluginLoader(void* parent);
void* QPluginLoader_NewQPluginLoader2(char* fileName, void* parent);
char* QPluginLoader_ErrorString(void* ptr);
void* QPluginLoader_Instance(void* ptr);
int QPluginLoader_IsLoaded(void* ptr);
int QPluginLoader_Load(void* ptr);
void* QPluginLoader_MetaData(void* ptr);
int QPluginLoader_Unload(void* ptr);
void QPluginLoader_DestroyQPluginLoader(void* ptr);

#ifdef __cplusplus
}
#endif