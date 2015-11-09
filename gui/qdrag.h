#ifdef __cplusplus
extern "C" {
#endif

void* QDrag_NewQDrag(void* dragSource);
void QDrag_ConnectActionChanged(void* ptr);
void QDrag_DisconnectActionChanged(void* ptr);
int QDrag_Exec(void* ptr, int supportedActions);
int QDrag_Exec2(void* ptr, int supportedActions, int defaultDropAction);
void* QDrag_MimeData(void* ptr);
void QDrag_SetDragCursor(void* ptr, void* cursor, int action);
void QDrag_SetHotSpot(void* ptr, void* hotspot);
void QDrag_SetMimeData(void* ptr, void* data);
void QDrag_SetPixmap(void* ptr, void* pixmap);
void* QDrag_Source(void* ptr);
int QDrag_SupportedActions(void* ptr);
void* QDrag_Target(void* ptr);
void QDrag_ConnectTargetChanged(void* ptr);
void QDrag_DisconnectTargetChanged(void* ptr);
void QDrag_DestroyQDrag(void* ptr);

#ifdef __cplusplus
}
#endif