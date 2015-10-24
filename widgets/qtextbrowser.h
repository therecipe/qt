#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QTextBrowser_OpenExternalLinks(QtObjectPtr ptr);
int QTextBrowser_OpenLinks(QtObjectPtr ptr);
char* QTextBrowser_SearchPaths(QtObjectPtr ptr);
void QTextBrowser_SetOpenExternalLinks(QtObjectPtr ptr, int open);
void QTextBrowser_SetOpenLinks(QtObjectPtr ptr, int open);
void QTextBrowser_SetSearchPaths(QtObjectPtr ptr, char* paths);
void QTextBrowser_SetSource(QtObjectPtr ptr, char* name);
char* QTextBrowser_Source(QtObjectPtr ptr);
QtObjectPtr QTextBrowser_NewQTextBrowser(QtObjectPtr parent);
void QTextBrowser_ConnectAnchorClicked(QtObjectPtr ptr);
void QTextBrowser_DisconnectAnchorClicked(QtObjectPtr ptr);
void QTextBrowser_Backward(QtObjectPtr ptr);
void QTextBrowser_ConnectBackwardAvailable(QtObjectPtr ptr);
void QTextBrowser_DisconnectBackwardAvailable(QtObjectPtr ptr);
int QTextBrowser_BackwardHistoryCount(QtObjectPtr ptr);
void QTextBrowser_ClearHistory(QtObjectPtr ptr);
void QTextBrowser_Forward(QtObjectPtr ptr);
void QTextBrowser_ConnectForwardAvailable(QtObjectPtr ptr);
void QTextBrowser_DisconnectForwardAvailable(QtObjectPtr ptr);
int QTextBrowser_ForwardHistoryCount(QtObjectPtr ptr);
void QTextBrowser_ConnectHighlighted(QtObjectPtr ptr);
void QTextBrowser_DisconnectHighlighted(QtObjectPtr ptr);
void QTextBrowser_ConnectHistoryChanged(QtObjectPtr ptr);
void QTextBrowser_DisconnectHistoryChanged(QtObjectPtr ptr);
char* QTextBrowser_HistoryTitle(QtObjectPtr ptr, int i);
char* QTextBrowser_HistoryUrl(QtObjectPtr ptr, int i);
void QTextBrowser_Home(QtObjectPtr ptr);
int QTextBrowser_IsBackwardAvailable(QtObjectPtr ptr);
int QTextBrowser_IsForwardAvailable(QtObjectPtr ptr);
char* QTextBrowser_LoadResource(QtObjectPtr ptr, int ty, char* name);
void QTextBrowser_Reload(QtObjectPtr ptr);
void QTextBrowser_ConnectSourceChanged(QtObjectPtr ptr);
void QTextBrowser_DisconnectSourceChanged(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif