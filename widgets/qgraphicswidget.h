#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QGraphicsWidget_AutoFillBackground(QtObjectPtr ptr);
int QGraphicsWidget_FocusPolicy(QtObjectPtr ptr);
int QGraphicsWidget_LayoutDirection(QtObjectPtr ptr);
void QGraphicsWidget_Resize(QtObjectPtr ptr, QtObjectPtr size);
void QGraphicsWidget_SetAutoFillBackground(QtObjectPtr ptr, int enabled);
void QGraphicsWidget_SetFocusPolicy(QtObjectPtr ptr, int policy);
void QGraphicsWidget_SetFont(QtObjectPtr ptr, QtObjectPtr font);
void QGraphicsWidget_SetGeometry(QtObjectPtr ptr, QtObjectPtr rect);
void QGraphicsWidget_SetLayout(QtObjectPtr ptr, QtObjectPtr layout);
void QGraphicsWidget_SetLayoutDirection(QtObjectPtr ptr, int direction);
void QGraphicsWidget_SetPalette(QtObjectPtr ptr, QtObjectPtr palette);
void QGraphicsWidget_SetWindowFlags(QtObjectPtr ptr, int wFlags);
void QGraphicsWidget_SetWindowTitle(QtObjectPtr ptr, char* title);
void QGraphicsWidget_UnsetLayoutDirection(QtObjectPtr ptr);
int QGraphicsWidget_WindowFlags(QtObjectPtr ptr);
char* QGraphicsWidget_WindowTitle(QtObjectPtr ptr);
QtObjectPtr QGraphicsWidget_NewQGraphicsWidget(QtObjectPtr parent, int wFlags);
void QGraphicsWidget_AddAction(QtObjectPtr ptr, QtObjectPtr action);
void QGraphicsWidget_AdjustSize(QtObjectPtr ptr);
int QGraphicsWidget_Close(QtObjectPtr ptr);
QtObjectPtr QGraphicsWidget_FocusWidget(QtObjectPtr ptr);
void QGraphicsWidget_ConnectGeometryChanged(QtObjectPtr ptr);
void QGraphicsWidget_DisconnectGeometryChanged(QtObjectPtr ptr);
int QGraphicsWidget_GrabShortcut(QtObjectPtr ptr, QtObjectPtr sequence, int context);
void QGraphicsWidget_InsertAction(QtObjectPtr ptr, QtObjectPtr before, QtObjectPtr action);
int QGraphicsWidget_IsActiveWindow(QtObjectPtr ptr);
QtObjectPtr QGraphicsWidget_Layout(QtObjectPtr ptr);
void QGraphicsWidget_Paint(QtObjectPtr ptr, QtObjectPtr painter, QtObjectPtr option, QtObjectPtr widget);
void QGraphicsWidget_PaintWindowFrame(QtObjectPtr ptr, QtObjectPtr painter, QtObjectPtr option, QtObjectPtr widget);
void QGraphicsWidget_ReleaseShortcut(QtObjectPtr ptr, int id);
void QGraphicsWidget_RemoveAction(QtObjectPtr ptr, QtObjectPtr action);
void QGraphicsWidget_SetAttribute(QtObjectPtr ptr, int attribute, int on);
void QGraphicsWidget_SetShortcutAutoRepeat(QtObjectPtr ptr, int id, int enabled);
void QGraphicsWidget_SetShortcutEnabled(QtObjectPtr ptr, int id, int enabled);
void QGraphicsWidget_SetStyle(QtObjectPtr ptr, QtObjectPtr style);
void QGraphicsWidget_QGraphicsWidget_SetTabOrder(QtObjectPtr first, QtObjectPtr second);
QtObjectPtr QGraphicsWidget_Style(QtObjectPtr ptr);
int QGraphicsWidget_TestAttribute(QtObjectPtr ptr, int attribute);
int QGraphicsWidget_Type(QtObjectPtr ptr);
void QGraphicsWidget_UnsetWindowFrameMargins(QtObjectPtr ptr);
int QGraphicsWidget_WindowType(QtObjectPtr ptr);
void QGraphicsWidget_DestroyQGraphicsWidget(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif