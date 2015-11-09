#ifdef __cplusplus
extern "C" {
#endif

int QTextBrowser_OpenExternalLinks(void* ptr);
int QTextBrowser_OpenLinks(void* ptr);
char* QTextBrowser_SearchPaths(void* ptr);
void QTextBrowser_SetOpenExternalLinks(void* ptr, int open);
void QTextBrowser_SetOpenLinks(void* ptr, int open);
void QTextBrowser_SetSearchPaths(void* ptr, char* paths);
void QTextBrowser_SetSource(void* ptr, void* name);
void* QTextBrowser_NewQTextBrowser(void* parent);
void QTextBrowser_Backward(void* ptr);
void QTextBrowser_ConnectBackwardAvailable(void* ptr);
void QTextBrowser_DisconnectBackwardAvailable(void* ptr);
int QTextBrowser_BackwardHistoryCount(void* ptr);
void QTextBrowser_ClearHistory(void* ptr);
void QTextBrowser_Forward(void* ptr);
void QTextBrowser_ConnectForwardAvailable(void* ptr);
void QTextBrowser_DisconnectForwardAvailable(void* ptr);
int QTextBrowser_ForwardHistoryCount(void* ptr);
void QTextBrowser_ConnectHistoryChanged(void* ptr);
void QTextBrowser_DisconnectHistoryChanged(void* ptr);
char* QTextBrowser_HistoryTitle(void* ptr, int i);
void QTextBrowser_Home(void* ptr);
int QTextBrowser_IsBackwardAvailable(void* ptr);
int QTextBrowser_IsForwardAvailable(void* ptr);
void* QTextBrowser_LoadResource(void* ptr, int ty, void* name);
void QTextBrowser_Reload(void* ptr);

#ifdef __cplusplus
}
#endif