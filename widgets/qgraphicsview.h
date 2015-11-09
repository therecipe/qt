#ifdef __cplusplus
extern "C" {
#endif

int QGraphicsView_Alignment(void* ptr);
void* QGraphicsView_BackgroundBrush(void* ptr);
int QGraphicsView_CacheMode(void* ptr);
int QGraphicsView_DragMode(void* ptr);
void* QGraphicsView_ForegroundBrush(void* ptr);
int QGraphicsView_IsInteractive(void* ptr);
int QGraphicsView_OptimizationFlags(void* ptr);
int QGraphicsView_RenderHints(void* ptr);
int QGraphicsView_ResizeAnchor(void* ptr);
int QGraphicsView_RubberBandSelectionMode(void* ptr);
void QGraphicsView_SetAlignment(void* ptr, int alignment);
void QGraphicsView_SetBackgroundBrush(void* ptr, void* brush);
void QGraphicsView_SetCacheMode(void* ptr, int mode);
void QGraphicsView_SetDragMode(void* ptr, int mode);
void QGraphicsView_SetForegroundBrush(void* ptr, void* brush);
void QGraphicsView_SetInteractive(void* ptr, int allowed);
void QGraphicsView_SetOptimizationFlags(void* ptr, int flags);
void QGraphicsView_SetRenderHints(void* ptr, int hints);
void QGraphicsView_SetResizeAnchor(void* ptr, int anchor);
void QGraphicsView_SetRubberBandSelectionMode(void* ptr, int mode);
void QGraphicsView_SetSceneRect(void* ptr, void* rect);
void QGraphicsView_SetTransformationAnchor(void* ptr, int anchor);
void QGraphicsView_SetViewportUpdateMode(void* ptr, int mode);
int QGraphicsView_TransformationAnchor(void* ptr);
int QGraphicsView_ViewportUpdateMode(void* ptr);
void* QGraphicsView_NewQGraphicsView2(void* scene, void* parent);
void* QGraphicsView_NewQGraphicsView(void* parent);
void QGraphicsView_CenterOn3(void* ptr, void* item);
void QGraphicsView_CenterOn(void* ptr, void* pos);
void QGraphicsView_CenterOn2(void* ptr, double x, double y);
void QGraphicsView_EnsureVisible3(void* ptr, void* item, int xmargin, int ymargin);
void QGraphicsView_EnsureVisible(void* ptr, void* rect, int xmargin, int ymargin);
void QGraphicsView_EnsureVisible2(void* ptr, double x, double y, double w, double h, int xmargin, int ymargin);
void QGraphicsView_FitInView3(void* ptr, void* item, int aspectRatioMode);
void QGraphicsView_FitInView(void* ptr, void* rect, int aspectRatioMode);
void QGraphicsView_FitInView2(void* ptr, double x, double y, double w, double h, int aspectRatioMode);
void* QGraphicsView_InputMethodQuery(void* ptr, int query);
void QGraphicsView_InvalidateScene(void* ptr, void* rect, int layers);
int QGraphicsView_IsTransformed(void* ptr);
void* QGraphicsView_ItemAt(void* ptr, void* pos);
void* QGraphicsView_ItemAt2(void* ptr, int x, int y);
void QGraphicsView_Render(void* ptr, void* painter, void* target, void* source, int aspectRatioMode);
void QGraphicsView_ResetCachedContent(void* ptr);
void QGraphicsView_ResetMatrix(void* ptr);
void QGraphicsView_ResetTransform(void* ptr);
void QGraphicsView_Rotate(void* ptr, double angle);
void QGraphicsView_Scale(void* ptr, double sx, double sy);
void* QGraphicsView_Scene(void* ptr);
void QGraphicsView_SetOptimizationFlag(void* ptr, int flag, int enabled);
void QGraphicsView_SetRenderHint(void* ptr, int hint, int enabled);
void QGraphicsView_SetScene(void* ptr, void* scene);
void QGraphicsView_SetSceneRect2(void* ptr, double x, double y, double w, double h);
void QGraphicsView_SetTransform(void* ptr, void* matrix, int combine);
void QGraphicsView_Shear(void* ptr, double sh, double sv);
void QGraphicsView_Translate(void* ptr, double dx, double dy);
void QGraphicsView_UpdateSceneRect(void* ptr, void* rect);
void QGraphicsView_DestroyQGraphicsView(void* ptr);

#ifdef __cplusplus
}
#endif