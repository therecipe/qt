#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QClipboard_Clear(QtObjectPtr ptr, int mode);
QtObjectPtr QClipboard_MimeData(QtObjectPtr ptr, int mode);
void QClipboard_SetMimeData(QtObjectPtr ptr, QtObjectPtr src, int mode);
void QClipboard_ConnectChanged(QtObjectPtr ptr);
void QClipboard_DisconnectChanged(QtObjectPtr ptr);
void QClipboard_ConnectDataChanged(QtObjectPtr ptr);
void QClipboard_DisconnectDataChanged(QtObjectPtr ptr);
void QClipboard_ConnectFindBufferChanged(QtObjectPtr ptr);
void QClipboard_DisconnectFindBufferChanged(QtObjectPtr ptr);
int QClipboard_OwnsClipboard(QtObjectPtr ptr);
int QClipboard_OwnsFindBuffer(QtObjectPtr ptr);
int QClipboard_OwnsSelection(QtObjectPtr ptr);
void QClipboard_ConnectSelectionChanged(QtObjectPtr ptr);
void QClipboard_DisconnectSelectionChanged(QtObjectPtr ptr);
void QClipboard_SetImage(QtObjectPtr ptr, QtObjectPtr image, int mode);
void QClipboard_SetPixmap(QtObjectPtr ptr, QtObjectPtr pixmap, int mode);
void QClipboard_SetText(QtObjectPtr ptr, char* text, int mode);
int QClipboard_SupportsFindBuffer(QtObjectPtr ptr);
int QClipboard_SupportsSelection(QtObjectPtr ptr);
char* QClipboard_Text(QtObjectPtr ptr, int mode);

#ifdef __cplusplus
}
#endif