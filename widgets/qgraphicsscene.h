#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QGraphicsScene_BspTreeDepth(QtObjectPtr ptr);
int QGraphicsScene_IsSortCacheEnabled(QtObjectPtr ptr);
int QGraphicsScene_ItemIndexMethod(QtObjectPtr ptr);
void QGraphicsScene_SetBackgroundBrush(QtObjectPtr ptr, QtObjectPtr brush);
void QGraphicsScene_SetBspTreeDepth(QtObjectPtr ptr, int depth);
void QGraphicsScene_SetFont(QtObjectPtr ptr, QtObjectPtr font);
void QGraphicsScene_SetForegroundBrush(QtObjectPtr ptr, QtObjectPtr brush);
void QGraphicsScene_SetItemIndexMethod(QtObjectPtr ptr, int method);
void QGraphicsScene_SetPalette(QtObjectPtr ptr, QtObjectPtr palette);
void QGraphicsScene_SetSceneRect(QtObjectPtr ptr, QtObjectPtr rect);
void QGraphicsScene_SetSortCacheEnabled(QtObjectPtr ptr, int enabled);
void QGraphicsScene_SetStickyFocus(QtObjectPtr ptr, int enabled);
int QGraphicsScene_StickyFocus(QtObjectPtr ptr);
void QGraphicsScene_Update(QtObjectPtr ptr, QtObjectPtr rect);
QtObjectPtr QGraphicsScene_NewQGraphicsScene(QtObjectPtr parent);
QtObjectPtr QGraphicsScene_NewQGraphicsScene2(QtObjectPtr sceneRect, QtObjectPtr parent);
QtObjectPtr QGraphicsScene_ActivePanel(QtObjectPtr ptr);
QtObjectPtr QGraphicsScene_ActiveWindow(QtObjectPtr ptr);
QtObjectPtr QGraphicsScene_AddEllipse(QtObjectPtr ptr, QtObjectPtr rect, QtObjectPtr pen, QtObjectPtr brush);
void QGraphicsScene_AddItem(QtObjectPtr ptr, QtObjectPtr item);
QtObjectPtr QGraphicsScene_AddLine(QtObjectPtr ptr, QtObjectPtr line, QtObjectPtr pen);
QtObjectPtr QGraphicsScene_AddPath(QtObjectPtr ptr, QtObjectPtr path, QtObjectPtr pen, QtObjectPtr brush);
QtObjectPtr QGraphicsScene_AddPixmap(QtObjectPtr ptr, QtObjectPtr pixmap);
QtObjectPtr QGraphicsScene_AddPolygon(QtObjectPtr ptr, QtObjectPtr polygon, QtObjectPtr pen, QtObjectPtr brush);
QtObjectPtr QGraphicsScene_AddRect(QtObjectPtr ptr, QtObjectPtr rect, QtObjectPtr pen, QtObjectPtr brush);
QtObjectPtr QGraphicsScene_AddSimpleText(QtObjectPtr ptr, char* text, QtObjectPtr font);
QtObjectPtr QGraphicsScene_AddText(QtObjectPtr ptr, char* text, QtObjectPtr font);
QtObjectPtr QGraphicsScene_AddWidget(QtObjectPtr ptr, QtObjectPtr widget, int wFlags);
void QGraphicsScene_Advance(QtObjectPtr ptr);
void QGraphicsScene_Clear(QtObjectPtr ptr);
void QGraphicsScene_ClearFocus(QtObjectPtr ptr);
void QGraphicsScene_ClearSelection(QtObjectPtr ptr);
void QGraphicsScene_DestroyItemGroup(QtObjectPtr ptr, QtObjectPtr group);
QtObjectPtr QGraphicsScene_FocusItem(QtObjectPtr ptr);
void QGraphicsScene_ConnectFocusItemChanged(QtObjectPtr ptr);
void QGraphicsScene_DisconnectFocusItemChanged(QtObjectPtr ptr);
int QGraphicsScene_HasFocus(QtObjectPtr ptr);
char* QGraphicsScene_InputMethodQuery(QtObjectPtr ptr, int query);
void QGraphicsScene_Invalidate(QtObjectPtr ptr, QtObjectPtr rect, int layers);
int QGraphicsScene_IsActive(QtObjectPtr ptr);
QtObjectPtr QGraphicsScene_ItemAt(QtObjectPtr ptr, QtObjectPtr position, QtObjectPtr deviceTransform);
QtObjectPtr QGraphicsScene_MouseGrabberItem(QtObjectPtr ptr);
void QGraphicsScene_RemoveItem(QtObjectPtr ptr, QtObjectPtr item);
void QGraphicsScene_Render(QtObjectPtr ptr, QtObjectPtr painter, QtObjectPtr target, QtObjectPtr source, int aspectRatioMode);
void QGraphicsScene_ConnectSelectionChanged(QtObjectPtr ptr);
void QGraphicsScene_DisconnectSelectionChanged(QtObjectPtr ptr);
int QGraphicsScene_SendEvent(QtObjectPtr ptr, QtObjectPtr item, QtObjectPtr event);
void QGraphicsScene_SetActivePanel(QtObjectPtr ptr, QtObjectPtr item);
void QGraphicsScene_SetActiveWindow(QtObjectPtr ptr, QtObjectPtr widget);
void QGraphicsScene_SetFocus(QtObjectPtr ptr, int focusReason);
void QGraphicsScene_SetFocusItem(QtObjectPtr ptr, QtObjectPtr item, int focusReason);
void QGraphicsScene_SetSelectionArea2(QtObjectPtr ptr, QtObjectPtr path, int mode, QtObjectPtr deviceTransform);
void QGraphicsScene_SetSelectionArea3(QtObjectPtr ptr, QtObjectPtr path, int selectionOperation, int mode, QtObjectPtr deviceTransform);
void QGraphicsScene_SetSelectionArea(QtObjectPtr ptr, QtObjectPtr path, QtObjectPtr deviceTransform);
void QGraphicsScene_SetStyle(QtObjectPtr ptr, QtObjectPtr style);
QtObjectPtr QGraphicsScene_Style(QtObjectPtr ptr);
void QGraphicsScene_DestroyQGraphicsScene(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif