#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QDrag_NewQDrag(QtObjectPtr dragSource);
void QDrag_ConnectActionChanged(QtObjectPtr ptr);
void QDrag_DisconnectActionChanged(QtObjectPtr ptr);
int QDrag_Exec(QtObjectPtr ptr, int supportedActions);
int QDrag_Exec2(QtObjectPtr ptr, int supportedActions, int defaultDropAction);
QtObjectPtr QDrag_MimeData(QtObjectPtr ptr);
void QDrag_SetDragCursor(QtObjectPtr ptr, QtObjectPtr cursor, int action);
void QDrag_SetHotSpot(QtObjectPtr ptr, QtObjectPtr hotspot);
void QDrag_SetMimeData(QtObjectPtr ptr, QtObjectPtr data);
void QDrag_SetPixmap(QtObjectPtr ptr, QtObjectPtr pixmap);
QtObjectPtr QDrag_Source(QtObjectPtr ptr);
int QDrag_SupportedActions(QtObjectPtr ptr);
QtObjectPtr QDrag_Target(QtObjectPtr ptr);
void QDrag_ConnectTargetChanged(QtObjectPtr ptr);
void QDrag_DisconnectTargetChanged(QtObjectPtr ptr);
void QDrag_DestroyQDrag(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif