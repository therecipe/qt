#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

char* QFileSystemWatcher_Directories(QtObjectPtr ptr);
char* QFileSystemWatcher_Files(QtObjectPtr ptr);
QtObjectPtr QFileSystemWatcher_NewQFileSystemWatcher(QtObjectPtr parent);
QtObjectPtr QFileSystemWatcher_NewQFileSystemWatcher2(char* paths, QtObjectPtr parent);
int QFileSystemWatcher_AddPath(QtObjectPtr ptr, char* path);
char* QFileSystemWatcher_AddPaths(QtObjectPtr ptr, char* paths);
void QFileSystemWatcher_ConnectDirectoryChanged(QtObjectPtr ptr);
void QFileSystemWatcher_DisconnectDirectoryChanged(QtObjectPtr ptr);
void QFileSystemWatcher_ConnectFileChanged(QtObjectPtr ptr);
void QFileSystemWatcher_DisconnectFileChanged(QtObjectPtr ptr);
int QFileSystemWatcher_RemovePath(QtObjectPtr ptr, char* path);
char* QFileSystemWatcher_RemovePaths(QtObjectPtr ptr, char* paths);
void QFileSystemWatcher_DestroyQFileSystemWatcher(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif