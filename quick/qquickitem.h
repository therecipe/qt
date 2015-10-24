#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QQuickItem_NewQQuickItem(QtObjectPtr parent);
int QQuickItem_ActiveFocusOnTab(QtObjectPtr ptr);
int QQuickItem_Antialiasing(QtObjectPtr ptr);
int QQuickItem_Clip(QtObjectPtr ptr);
int QQuickItem_HasActiveFocus(QtObjectPtr ptr);
int QQuickItem_HasFocus(QtObjectPtr ptr);
int QQuickItem_IsEnabled(QtObjectPtr ptr);
int QQuickItem_IsTextureProvider(QtObjectPtr ptr);
int QQuickItem_IsVisible(QtObjectPtr ptr);
QtObjectPtr QQuickItem_ParentItem(QtObjectPtr ptr);
void QQuickItem_ResetAntialiasing(QtObjectPtr ptr);
void QQuickItem_ResetHeight(QtObjectPtr ptr);
void QQuickItem_ResetWidth(QtObjectPtr ptr);
void QQuickItem_SetActiveFocusOnTab(QtObjectPtr ptr, int v);
void QQuickItem_SetAntialiasing(QtObjectPtr ptr, int v);
void QQuickItem_SetClip(QtObjectPtr ptr, int v);
void QQuickItem_SetEnabled(QtObjectPtr ptr, int v);
void QQuickItem_SetFocus(QtObjectPtr ptr, int v);
void QQuickItem_SetFocus2(QtObjectPtr ptr, int focus, int reason);
void QQuickItem_SetParentItem(QtObjectPtr ptr, QtObjectPtr parent);
void QQuickItem_SetSmooth(QtObjectPtr ptr, int v);
void QQuickItem_SetState(QtObjectPtr ptr, char* v);
void QQuickItem_SetTransformOrigin(QtObjectPtr ptr, int v);
void QQuickItem_SetVisible(QtObjectPtr ptr, int v);
int QQuickItem_Smooth(QtObjectPtr ptr);
QtObjectPtr QQuickItem_TextureProvider(QtObjectPtr ptr);
int QQuickItem_TransformOrigin(QtObjectPtr ptr);
int QQuickItem_AcceptHoverEvents(QtObjectPtr ptr);
int QQuickItem_AcceptedMouseButtons(QtObjectPtr ptr);
int QQuickItem_Contains(QtObjectPtr ptr, QtObjectPtr point);
int QQuickItem_FiltersChildMouseEvents(QtObjectPtr ptr);
int QQuickItem_Flags(QtObjectPtr ptr);
void QQuickItem_ForceActiveFocus(QtObjectPtr ptr);
void QQuickItem_ForceActiveFocus2(QtObjectPtr ptr, int reason);
void QQuickItem_GrabMouse(QtObjectPtr ptr);
char* QQuickItem_InputMethodQuery(QtObjectPtr ptr, int query);
int QQuickItem_IsFocusScope(QtObjectPtr ptr);
int QQuickItem_KeepMouseGrab(QtObjectPtr ptr);
int QQuickItem_KeepTouchGrab(QtObjectPtr ptr);
QtObjectPtr QQuickItem_NextItemInFocusChain(QtObjectPtr ptr, int forward);
void QQuickItem_Polish(QtObjectPtr ptr);
QtObjectPtr QQuickItem_ScopedFocusItem(QtObjectPtr ptr);
void QQuickItem_SetAcceptHoverEvents(QtObjectPtr ptr, int enabled);
void QQuickItem_SetAcceptedMouseButtons(QtObjectPtr ptr, int buttons);
void QQuickItem_SetCursor(QtObjectPtr ptr, QtObjectPtr cursor);
void QQuickItem_SetFiltersChildMouseEvents(QtObjectPtr ptr, int filter);
void QQuickItem_SetFlag(QtObjectPtr ptr, int flag, int enabled);
void QQuickItem_SetFlags(QtObjectPtr ptr, int flags);
void QQuickItem_SetKeepMouseGrab(QtObjectPtr ptr, int keep);
void QQuickItem_SetKeepTouchGrab(QtObjectPtr ptr, int keep);
void QQuickItem_StackAfter(QtObjectPtr ptr, QtObjectPtr sibling);
void QQuickItem_StackBefore(QtObjectPtr ptr, QtObjectPtr sibling);
void QQuickItem_UngrabMouse(QtObjectPtr ptr);
void QQuickItem_UngrabTouchPoints(QtObjectPtr ptr);
void QQuickItem_UnsetCursor(QtObjectPtr ptr);
void QQuickItem_Update(QtObjectPtr ptr);
QtObjectPtr QQuickItem_Window(QtObjectPtr ptr);
void QQuickItem_ConnectWindowChanged(QtObjectPtr ptr);
void QQuickItem_DisconnectWindowChanged(QtObjectPtr ptr);
void QQuickItem_DestroyQQuickItem(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif