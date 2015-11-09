#ifdef __cplusplus
extern "C" {
#endif

int QDesktopServices_QDesktopServices_OpenUrl(void* url);
void QDesktopServices_QDesktopServices_SetUrlHandler(char* scheme, void* receiver, char* method);
void QDesktopServices_QDesktopServices_UnsetUrlHandler(char* scheme);

#ifdef __cplusplus
}
#endif