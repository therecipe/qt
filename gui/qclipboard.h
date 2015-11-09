#ifdef __cplusplus
extern "C" {
#endif

void QClipboard_Clear(void* ptr, int mode);
void* QClipboard_MimeData(void* ptr, int mode);
void QClipboard_SetMimeData(void* ptr, void* src, int mode);
void QClipboard_ConnectChanged(void* ptr);
void QClipboard_DisconnectChanged(void* ptr);
void QClipboard_ConnectDataChanged(void* ptr);
void QClipboard_DisconnectDataChanged(void* ptr);
void QClipboard_ConnectFindBufferChanged(void* ptr);
void QClipboard_DisconnectFindBufferChanged(void* ptr);
int QClipboard_OwnsClipboard(void* ptr);
int QClipboard_OwnsFindBuffer(void* ptr);
int QClipboard_OwnsSelection(void* ptr);
void QClipboard_ConnectSelectionChanged(void* ptr);
void QClipboard_DisconnectSelectionChanged(void* ptr);
void QClipboard_SetImage(void* ptr, void* image, int mode);
void QClipboard_SetPixmap(void* ptr, void* pixmap, int mode);
void QClipboard_SetText(void* ptr, char* text, int mode);
int QClipboard_SupportsFindBuffer(void* ptr);
int QClipboard_SupportsSelection(void* ptr);
char* QClipboard_Text(void* ptr, int mode);

#ifdef __cplusplus
}
#endif