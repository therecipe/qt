#ifdef __cplusplus
extern "C" {
#endif

int QGraphicsWidget_AutoFillBackground(void* ptr);
int QGraphicsWidget_FocusPolicy(void* ptr);
int QGraphicsWidget_LayoutDirection(void* ptr);
void QGraphicsWidget_Resize(void* ptr, void* size);
void QGraphicsWidget_SetAutoFillBackground(void* ptr, int enabled);
void QGraphicsWidget_SetFocusPolicy(void* ptr, int policy);
void QGraphicsWidget_SetFont(void* ptr, void* font);
void QGraphicsWidget_SetGeometry(void* ptr, void* rect);
void QGraphicsWidget_SetLayout(void* ptr, void* layout);
void QGraphicsWidget_SetLayoutDirection(void* ptr, int direction);
void QGraphicsWidget_SetPalette(void* ptr, void* palette);
void QGraphicsWidget_SetWindowFlags(void* ptr, int wFlags);
void QGraphicsWidget_SetWindowTitle(void* ptr, char* title);
void QGraphicsWidget_UnsetLayoutDirection(void* ptr);
int QGraphicsWidget_WindowFlags(void* ptr);
char* QGraphicsWidget_WindowTitle(void* ptr);
void* QGraphicsWidget_NewQGraphicsWidget(void* parent, int wFlags);
void QGraphicsWidget_AddAction(void* ptr, void* action);
void QGraphicsWidget_AdjustSize(void* ptr);
int QGraphicsWidget_Close(void* ptr);
void* QGraphicsWidget_FocusWidget(void* ptr);
void QGraphicsWidget_ConnectGeometryChanged(void* ptr);
void QGraphicsWidget_DisconnectGeometryChanged(void* ptr);
int QGraphicsWidget_GrabShortcut(void* ptr, void* sequence, int context);
void QGraphicsWidget_InsertAction(void* ptr, void* before, void* action);
int QGraphicsWidget_IsActiveWindow(void* ptr);
void* QGraphicsWidget_Layout(void* ptr);
void QGraphicsWidget_Paint(void* ptr, void* painter, void* option, void* widget);
void QGraphicsWidget_PaintWindowFrame(void* ptr, void* painter, void* option, void* widget);
void QGraphicsWidget_ReleaseShortcut(void* ptr, int id);
void QGraphicsWidget_RemoveAction(void* ptr, void* action);
void QGraphicsWidget_Resize2(void* ptr, double w, double h);
void QGraphicsWidget_SetAttribute(void* ptr, int attribute, int on);
void QGraphicsWidget_SetContentsMargins(void* ptr, double left, double top, double right, double bottom);
void QGraphicsWidget_SetGeometry2(void* ptr, double x, double y, double w, double h);
void QGraphicsWidget_SetShortcutAutoRepeat(void* ptr, int id, int enabled);
void QGraphicsWidget_SetShortcutEnabled(void* ptr, int id, int enabled);
void QGraphicsWidget_SetStyle(void* ptr, void* style);
void QGraphicsWidget_QGraphicsWidget_SetTabOrder(void* first, void* second);
void QGraphicsWidget_SetWindowFrameMargins(void* ptr, double left, double top, double right, double bottom);
void* QGraphicsWidget_Style(void* ptr);
int QGraphicsWidget_TestAttribute(void* ptr, int attribute);
int QGraphicsWidget_Type(void* ptr);
void QGraphicsWidget_UnsetWindowFrameMargins(void* ptr);
int QGraphicsWidget_WindowType(void* ptr);
void QGraphicsWidget_DestroyQGraphicsWidget(void* ptr);

#ifdef __cplusplus
}
#endif