#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

char* QPluginLoader_FileName(QtObjectPtr ptr);
int QPluginLoader_LoadHints(QtObjectPtr ptr);
void QPluginLoader_SetFileName(QtObjectPtr ptr, char* fileName);
void QPluginLoader_SetLoadHints(QtObjectPtr ptr, int loadHints);
QtObjectPtr QPluginLoader_NewQPluginLoader(QtObjectPtr parent);
QtObjectPtr QPluginLoader_NewQPluginLoader2(char* fileName, QtObjectPtr parent);
char* QPluginLoader_ErrorString(QtObjectPtr ptr);
QtObjectPtr QPluginLoader_Instance(QtObjectPtr ptr);
int QPluginLoader_IsLoaded(QtObjectPtr ptr);
int QPluginLoader_Load(QtObjectPtr ptr);
int QPluginLoader_Unload(QtObjectPtr ptr);
void QPluginLoader_DestroyQPluginLoader(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif