#ifdef __cplusplus
extern "C" {
#endif

char* QAbstractTextDocumentLayout_AnchorAt(void* ptr, void* position);
void* QAbstractTextDocumentLayout_Document(void* ptr);
void* QAbstractTextDocumentLayout_HandlerForObject(void* ptr, int objectType);
int QAbstractTextDocumentLayout_PageCount(void* ptr);
void QAbstractTextDocumentLayout_ConnectPageCountChanged(void* ptr);
void QAbstractTextDocumentLayout_DisconnectPageCountChanged(void* ptr);
void* QAbstractTextDocumentLayout_PaintDevice(void* ptr);
void QAbstractTextDocumentLayout_RegisterHandler(void* ptr, int objectType, void* component);
void QAbstractTextDocumentLayout_SetPaintDevice(void* ptr, void* device);
void QAbstractTextDocumentLayout_UnregisterHandler(void* ptr, int objectType, void* component);

#ifdef __cplusplus
}
#endif