#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QGraphicsView_Alignment(QtObjectPtr ptr);
int QGraphicsView_CacheMode(QtObjectPtr ptr);
int QGraphicsView_DragMode(QtObjectPtr ptr);
int QGraphicsView_IsInteractive(QtObjectPtr ptr);
int QGraphicsView_OptimizationFlags(QtObjectPtr ptr);
int QGraphicsView_RenderHints(QtObjectPtr ptr);
int QGraphicsView_ResizeAnchor(QtObjectPtr ptr);
int QGraphicsView_RubberBandSelectionMode(QtObjectPtr ptr);
void QGraphicsView_SetAlignment(QtObjectPtr ptr, int alignment);
void QGraphicsView_SetBackgroundBrush(QtObjectPtr ptr, QtObjectPtr brush);
void QGraphicsView_SetCacheMode(QtObjectPtr ptr, int mode);
void QGraphicsView_SetDragMode(QtObjectPtr ptr, int mode);
void QGraphicsView_SetForegroundBrush(QtObjectPtr ptr, QtObjectPtr brush);
void QGraphicsView_SetInteractive(QtObjectPtr ptr, int allowed);
void QGraphicsView_SetOptimizationFlags(QtObjectPtr ptr, int flags);
void QGraphicsView_SetRenderHints(QtObjectPtr ptr, int hints);
void QGraphicsView_SetResizeAnchor(QtObjectPtr ptr, int anchor);
void QGraphicsView_SetRubberBandSelectionMode(QtObjectPtr ptr, int mode);
void QGraphicsView_SetSceneRect(QtObjectPtr ptr, QtObjectPtr rect);
void QGraphicsView_SetTransformationAnchor(QtObjectPtr ptr, int anchor);
void QGraphicsView_SetViewportUpdateMode(QtObjectPtr ptr, int mode);
int QGraphicsView_TransformationAnchor(QtObjectPtr ptr);
int QGraphicsView_ViewportUpdateMode(QtObjectPtr ptr);
QtObjectPtr QGraphicsView_NewQGraphicsView2(QtObjectPtr scene, QtObjectPtr parent);
QtObjectPtr QGraphicsView_NewQGraphicsView(QtObjectPtr parent);
void QGraphicsView_CenterOn3(QtObjectPtr ptr, QtObjectPtr item);
void QGraphicsView_CenterOn(QtObjectPtr ptr, QtObjectPtr pos);
void QGraphicsView_EnsureVisible3(QtObjectPtr ptr, QtObjectPtr item, int xmargin, int ymargin);
void QGraphicsView_EnsureVisible(QtObjectPtr ptr, QtObjectPtr rect, int xmargin, int ymargin);
void QGraphicsView_FitInView3(QtObjectPtr ptr, QtObjectPtr item, int aspectRatioMode);
void QGraphicsView_FitInView(QtObjectPtr ptr, QtObjectPtr rect, int aspectRatioMode);
char* QGraphicsView_InputMethodQuery(QtObjectPtr ptr, int query);
void QGraphicsView_InvalidateScene(QtObjectPtr ptr, QtObjectPtr rect, int layers);
int QGraphicsView_IsTransformed(QtObjectPtr ptr);
QtObjectPtr QGraphicsView_ItemAt(QtObjectPtr ptr, QtObjectPtr pos);
QtObjectPtr QGraphicsView_ItemAt2(QtObjectPtr ptr, int x, int y);
void QGraphicsView_Render(QtObjectPtr ptr, QtObjectPtr painter, QtObjectPtr target, QtObjectPtr source, int aspectRatioMode);
void QGraphicsView_ResetCachedContent(QtObjectPtr ptr);
void QGraphicsView_ResetMatrix(QtObjectPtr ptr);
void QGraphicsView_ResetTransform(QtObjectPtr ptr);
QtObjectPtr QGraphicsView_Scene(QtObjectPtr ptr);
void QGraphicsView_SetOptimizationFlag(QtObjectPtr ptr, int flag, int enabled);
void QGraphicsView_SetRenderHint(QtObjectPtr ptr, int hint, int enabled);
void QGraphicsView_SetScene(QtObjectPtr ptr, QtObjectPtr scene);
void QGraphicsView_SetTransform(QtObjectPtr ptr, QtObjectPtr matrix, int combine);
void QGraphicsView_UpdateSceneRect(QtObjectPtr ptr, QtObjectPtr rect);
void QGraphicsView_DestroyQGraphicsView(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif