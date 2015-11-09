#ifdef __cplusplus
extern "C" {
#endif

char* QFileSystemWatcher_Directories(void* ptr);
char* QFileSystemWatcher_Files(void* ptr);
void* QFileSystemWatcher_NewQFileSystemWatcher(void* parent);
void* QFileSystemWatcher_NewQFileSystemWatcher2(char* paths, void* parent);
int QFileSystemWatcher_AddPath(void* ptr, char* path);
char* QFileSystemWatcher_AddPaths(void* ptr, char* paths);
void QFileSystemWatcher_ConnectDirectoryChanged(void* ptr);
void QFileSystemWatcher_DisconnectDirectoryChanged(void* ptr);
void QFileSystemWatcher_ConnectFileChanged(void* ptr);
void QFileSystemWatcher_DisconnectFileChanged(void* ptr);
int QFileSystemWatcher_RemovePath(void* ptr, char* path);
char* QFileSystemWatcher_RemovePaths(void* ptr, char* paths);
void QFileSystemWatcher_DestroyQFileSystemWatcher(void* ptr);

#ifdef __cplusplus
}
#endif