#ifdef __cplusplus
extern "C" {
#endif

void* QGraphicsScene_BackgroundBrush(void* ptr);
int QGraphicsScene_BspTreeDepth(void* ptr);
void* QGraphicsScene_ForegroundBrush(void* ptr);
int QGraphicsScene_IsSortCacheEnabled(void* ptr);
int QGraphicsScene_ItemIndexMethod(void* ptr);
double QGraphicsScene_MinimumRenderSize(void* ptr);
void QGraphicsScene_SetBackgroundBrush(void* ptr, void* brush);
void QGraphicsScene_SetBspTreeDepth(void* ptr, int depth);
void QGraphicsScene_SetFont(void* ptr, void* font);
void QGraphicsScene_SetForegroundBrush(void* ptr, void* brush);
void QGraphicsScene_SetItemIndexMethod(void* ptr, int method);
void QGraphicsScene_SetMinimumRenderSize(void* ptr, double minSize);
void QGraphicsScene_SetPalette(void* ptr, void* palette);
void QGraphicsScene_SetSceneRect(void* ptr, void* rect);
void QGraphicsScene_SetSortCacheEnabled(void* ptr, int enabled);
void QGraphicsScene_SetStickyFocus(void* ptr, int enabled);
int QGraphicsScene_StickyFocus(void* ptr);
void QGraphicsScene_Update(void* ptr, void* rect);
void* QGraphicsScene_NewQGraphicsScene(void* parent);
void* QGraphicsScene_NewQGraphicsScene2(void* sceneRect, void* parent);
void* QGraphicsScene_NewQGraphicsScene3(double x, double y, double width, double height, void* parent);
void* QGraphicsScene_ActivePanel(void* ptr);
void* QGraphicsScene_ActiveWindow(void* ptr);
void* QGraphicsScene_AddEllipse(void* ptr, void* rect, void* pen, void* brush);
void* QGraphicsScene_AddEllipse2(void* ptr, double x, double y, double w, double h, void* pen, void* brush);
void QGraphicsScene_AddItem(void* ptr, void* item);
void* QGraphicsScene_AddLine(void* ptr, void* line, void* pen);
void* QGraphicsScene_AddLine2(void* ptr, double x1, double y1, double x2, double y2, void* pen);
void* QGraphicsScene_AddPath(void* ptr, void* path, void* pen, void* brush);
void* QGraphicsScene_AddPixmap(void* ptr, void* pixmap);
void* QGraphicsScene_AddPolygon(void* ptr, void* polygon, void* pen, void* brush);
void* QGraphicsScene_AddRect(void* ptr, void* rect, void* pen, void* brush);
void* QGraphicsScene_AddRect2(void* ptr, double x, double y, double w, double h, void* pen, void* brush);
void* QGraphicsScene_AddSimpleText(void* ptr, char* text, void* font);
void* QGraphicsScene_AddText(void* ptr, char* text, void* font);
void* QGraphicsScene_AddWidget(void* ptr, void* widget, int wFlags);
void QGraphicsScene_Advance(void* ptr);
void QGraphicsScene_Clear(void* ptr);
void QGraphicsScene_ClearFocus(void* ptr);
void QGraphicsScene_ClearSelection(void* ptr);
void QGraphicsScene_DestroyItemGroup(void* ptr, void* group);
void* QGraphicsScene_FocusItem(void* ptr);
void QGraphicsScene_ConnectFocusItemChanged(void* ptr);
void QGraphicsScene_DisconnectFocusItemChanged(void* ptr);
int QGraphicsScene_HasFocus(void* ptr);
double QGraphicsScene_Height(void* ptr);
void* QGraphicsScene_InputMethodQuery(void* ptr, int query);
void QGraphicsScene_Invalidate(void* ptr, void* rect, int layers);
void QGraphicsScene_Invalidate2(void* ptr, double x, double y, double w, double h, int layers);
int QGraphicsScene_IsActive(void* ptr);
void* QGraphicsScene_ItemAt(void* ptr, void* position, void* deviceTransform);
void* QGraphicsScene_ItemAt3(void* ptr, double x, double y, void* deviceTransform);
void* QGraphicsScene_MouseGrabberItem(void* ptr);
void QGraphicsScene_RemoveItem(void* ptr, void* item);
void QGraphicsScene_Render(void* ptr, void* painter, void* target, void* source, int aspectRatioMode);
void QGraphicsScene_ConnectSelectionChanged(void* ptr);
void QGraphicsScene_DisconnectSelectionChanged(void* ptr);
int QGraphicsScene_SendEvent(void* ptr, void* item, void* event);
void QGraphicsScene_SetActivePanel(void* ptr, void* item);
void QGraphicsScene_SetActiveWindow(void* ptr, void* widget);
void QGraphicsScene_SetFocus(void* ptr, int focusReason);
void QGraphicsScene_SetFocusItem(void* ptr, void* item, int focusReason);
void QGraphicsScene_SetSceneRect2(void* ptr, double x, double y, double w, double h);
void QGraphicsScene_SetSelectionArea2(void* ptr, void* path, int mode, void* deviceTransform);
void QGraphicsScene_SetSelectionArea3(void* ptr, void* path, int selectionOperation, int mode, void* deviceTransform);
void QGraphicsScene_SetSelectionArea(void* ptr, void* path, void* deviceTransform);
void QGraphicsScene_SetStyle(void* ptr, void* style);
void* QGraphicsScene_Style(void* ptr);
void QGraphicsScene_Update2(void* ptr, double x, double y, double w, double h);
double QGraphicsScene_Width(void* ptr);
void QGraphicsScene_DestroyQGraphicsScene(void* ptr);

#ifdef __cplusplus
}
#endif