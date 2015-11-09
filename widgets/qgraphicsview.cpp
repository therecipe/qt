#include "qgraphicsview.h"
#include <QUrl>
#include <QRect>
#include <QVariant>
#include <QRectF>
#include <QPoint>
#include <QPainter>
#include <QGraphicsItem>
#include <QString>
#include <QModelIndex>
#include <QWidget>
#include <QTransform>
#include <QGraphicsScene>
#include <QMetaObject>
#include <QPointF>
#include <QBrush>
#include <QGraphicsView>
#include "_cgo_export.h"

class MyQGraphicsView: public QGraphicsView {
public:
};

int QGraphicsView_Alignment(void* ptr){
	return static_cast<QGraphicsView*>(ptr)->alignment();
}

void* QGraphicsView_BackgroundBrush(void* ptr){
	return new QBrush(static_cast<QGraphicsView*>(ptr)->backgroundBrush());
}

int QGraphicsView_CacheMode(void* ptr){
	return static_cast<QGraphicsView*>(ptr)->cacheMode();
}

int QGraphicsView_DragMode(void* ptr){
	return static_cast<QGraphicsView*>(ptr)->dragMode();
}

void* QGraphicsView_ForegroundBrush(void* ptr){
	return new QBrush(static_cast<QGraphicsView*>(ptr)->foregroundBrush());
}

int QGraphicsView_IsInteractive(void* ptr){
	return static_cast<QGraphicsView*>(ptr)->isInteractive();
}

int QGraphicsView_OptimizationFlags(void* ptr){
	return static_cast<QGraphicsView*>(ptr)->optimizationFlags();
}

int QGraphicsView_RenderHints(void* ptr){
	return static_cast<QGraphicsView*>(ptr)->renderHints();
}

int QGraphicsView_ResizeAnchor(void* ptr){
	return static_cast<QGraphicsView*>(ptr)->resizeAnchor();
}

int QGraphicsView_RubberBandSelectionMode(void* ptr){
	return static_cast<QGraphicsView*>(ptr)->rubberBandSelectionMode();
}

void QGraphicsView_SetAlignment(void* ptr, int alignment){
	static_cast<QGraphicsView*>(ptr)->setAlignment(static_cast<Qt::AlignmentFlag>(alignment));
}

void QGraphicsView_SetBackgroundBrush(void* ptr, void* brush){
	static_cast<QGraphicsView*>(ptr)->setBackgroundBrush(*static_cast<QBrush*>(brush));
}

void QGraphicsView_SetCacheMode(void* ptr, int mode){
	static_cast<QGraphicsView*>(ptr)->setCacheMode(static_cast<QGraphicsView::CacheModeFlag>(mode));
}

void QGraphicsView_SetDragMode(void* ptr, int mode){
	static_cast<QGraphicsView*>(ptr)->setDragMode(static_cast<QGraphicsView::DragMode>(mode));
}

void QGraphicsView_SetForegroundBrush(void* ptr, void* brush){
	static_cast<QGraphicsView*>(ptr)->setForegroundBrush(*static_cast<QBrush*>(brush));
}

void QGraphicsView_SetInteractive(void* ptr, int allowed){
	static_cast<QGraphicsView*>(ptr)->setInteractive(allowed != 0);
}

void QGraphicsView_SetOptimizationFlags(void* ptr, int flags){
	static_cast<QGraphicsView*>(ptr)->setOptimizationFlags(static_cast<QGraphicsView::OptimizationFlag>(flags));
}

void QGraphicsView_SetRenderHints(void* ptr, int hints){
	static_cast<QGraphicsView*>(ptr)->setRenderHints(static_cast<QPainter::RenderHint>(hints));
}

void QGraphicsView_SetResizeAnchor(void* ptr, int anchor){
	static_cast<QGraphicsView*>(ptr)->setResizeAnchor(static_cast<QGraphicsView::ViewportAnchor>(anchor));
}

void QGraphicsView_SetRubberBandSelectionMode(void* ptr, int mode){
	static_cast<QGraphicsView*>(ptr)->setRubberBandSelectionMode(static_cast<Qt::ItemSelectionMode>(mode));
}

void QGraphicsView_SetSceneRect(void* ptr, void* rect){
	static_cast<QGraphicsView*>(ptr)->setSceneRect(*static_cast<QRectF*>(rect));
}

void QGraphicsView_SetTransformationAnchor(void* ptr, int anchor){
	static_cast<QGraphicsView*>(ptr)->setTransformationAnchor(static_cast<QGraphicsView::ViewportAnchor>(anchor));
}

void QGraphicsView_SetViewportUpdateMode(void* ptr, int mode){
	static_cast<QGraphicsView*>(ptr)->setViewportUpdateMode(static_cast<QGraphicsView::ViewportUpdateMode>(mode));
}

int QGraphicsView_TransformationAnchor(void* ptr){
	return static_cast<QGraphicsView*>(ptr)->transformationAnchor();
}

int QGraphicsView_ViewportUpdateMode(void* ptr){
	return static_cast<QGraphicsView*>(ptr)->viewportUpdateMode();
}

void* QGraphicsView_NewQGraphicsView2(void* scene, void* parent){
	return new QGraphicsView(static_cast<QGraphicsScene*>(scene), static_cast<QWidget*>(parent));
}

void* QGraphicsView_NewQGraphicsView(void* parent){
	return new QGraphicsView(static_cast<QWidget*>(parent));
}

void QGraphicsView_CenterOn3(void* ptr, void* item){
	static_cast<QGraphicsView*>(ptr)->centerOn(static_cast<QGraphicsItem*>(item));
}

void QGraphicsView_CenterOn(void* ptr, void* pos){
	static_cast<QGraphicsView*>(ptr)->centerOn(*static_cast<QPointF*>(pos));
}

void QGraphicsView_CenterOn2(void* ptr, double x, double y){
	static_cast<QGraphicsView*>(ptr)->centerOn(static_cast<qreal>(x), static_cast<qreal>(y));
}

void QGraphicsView_EnsureVisible3(void* ptr, void* item, int xmargin, int ymargin){
	static_cast<QGraphicsView*>(ptr)->ensureVisible(static_cast<QGraphicsItem*>(item), xmargin, ymargin);
}

void QGraphicsView_EnsureVisible(void* ptr, void* rect, int xmargin, int ymargin){
	static_cast<QGraphicsView*>(ptr)->ensureVisible(*static_cast<QRectF*>(rect), xmargin, ymargin);
}

void QGraphicsView_EnsureVisible2(void* ptr, double x, double y, double w, double h, int xmargin, int ymargin){
	static_cast<QGraphicsView*>(ptr)->ensureVisible(static_cast<qreal>(x), static_cast<qreal>(y), static_cast<qreal>(w), static_cast<qreal>(h), xmargin, ymargin);
}

void QGraphicsView_FitInView3(void* ptr, void* item, int aspectRatioMode){
	static_cast<QGraphicsView*>(ptr)->fitInView(static_cast<QGraphicsItem*>(item), static_cast<Qt::AspectRatioMode>(aspectRatioMode));
}

void QGraphicsView_FitInView(void* ptr, void* rect, int aspectRatioMode){
	static_cast<QGraphicsView*>(ptr)->fitInView(*static_cast<QRectF*>(rect), static_cast<Qt::AspectRatioMode>(aspectRatioMode));
}

void QGraphicsView_FitInView2(void* ptr, double x, double y, double w, double h, int aspectRatioMode){
	static_cast<QGraphicsView*>(ptr)->fitInView(static_cast<qreal>(x), static_cast<qreal>(y), static_cast<qreal>(w), static_cast<qreal>(h), static_cast<Qt::AspectRatioMode>(aspectRatioMode));
}

void* QGraphicsView_InputMethodQuery(void* ptr, int query){
	return new QVariant(static_cast<QGraphicsView*>(ptr)->inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

void QGraphicsView_InvalidateScene(void* ptr, void* rect, int layers){
	QMetaObject::invokeMethod(static_cast<QGraphicsView*>(ptr), "invalidateScene", Q_ARG(QRectF, *static_cast<QRectF*>(rect)), Q_ARG(QGraphicsScene::SceneLayer, static_cast<QGraphicsScene::SceneLayer>(layers)));
}

int QGraphicsView_IsTransformed(void* ptr){
	return static_cast<QGraphicsView*>(ptr)->isTransformed();
}

void* QGraphicsView_ItemAt(void* ptr, void* pos){
	return static_cast<QGraphicsView*>(ptr)->itemAt(*static_cast<QPoint*>(pos));
}

void* QGraphicsView_ItemAt2(void* ptr, int x, int y){
	return static_cast<QGraphicsView*>(ptr)->itemAt(x, y);
}

void QGraphicsView_Render(void* ptr, void* painter, void* target, void* source, int aspectRatioMode){
	static_cast<QGraphicsView*>(ptr)->render(static_cast<QPainter*>(painter), *static_cast<QRectF*>(target), *static_cast<QRect*>(source), static_cast<Qt::AspectRatioMode>(aspectRatioMode));
}

void QGraphicsView_ResetCachedContent(void* ptr){
	static_cast<QGraphicsView*>(ptr)->resetCachedContent();
}

void QGraphicsView_ResetMatrix(void* ptr){
	static_cast<QGraphicsView*>(ptr)->resetMatrix();
}

void QGraphicsView_ResetTransform(void* ptr){
	static_cast<QGraphicsView*>(ptr)->resetTransform();
}

void QGraphicsView_Rotate(void* ptr, double angle){
	static_cast<QGraphicsView*>(ptr)->rotate(static_cast<qreal>(angle));
}

void QGraphicsView_Scale(void* ptr, double sx, double sy){
	static_cast<QGraphicsView*>(ptr)->scale(static_cast<qreal>(sx), static_cast<qreal>(sy));
}

void* QGraphicsView_Scene(void* ptr){
	return static_cast<QGraphicsView*>(ptr)->scene();
}

void QGraphicsView_SetOptimizationFlag(void* ptr, int flag, int enabled){
	static_cast<QGraphicsView*>(ptr)->setOptimizationFlag(static_cast<QGraphicsView::OptimizationFlag>(flag), enabled != 0);
}

void QGraphicsView_SetRenderHint(void* ptr, int hint, int enabled){
	static_cast<QGraphicsView*>(ptr)->setRenderHint(static_cast<QPainter::RenderHint>(hint), enabled != 0);
}

void QGraphicsView_SetScene(void* ptr, void* scene){
	static_cast<QGraphicsView*>(ptr)->setScene(static_cast<QGraphicsScene*>(scene));
}

void QGraphicsView_SetSceneRect2(void* ptr, double x, double y, double w, double h){
	static_cast<QGraphicsView*>(ptr)->setSceneRect(static_cast<qreal>(x), static_cast<qreal>(y), static_cast<qreal>(w), static_cast<qreal>(h));
}

void QGraphicsView_SetTransform(void* ptr, void* matrix, int combine){
	static_cast<QGraphicsView*>(ptr)->setTransform(*static_cast<QTransform*>(matrix), combine != 0);
}

void QGraphicsView_Shear(void* ptr, double sh, double sv){
	static_cast<QGraphicsView*>(ptr)->shear(static_cast<qreal>(sh), static_cast<qreal>(sv));
}

void QGraphicsView_Translate(void* ptr, double dx, double dy){
	static_cast<QGraphicsView*>(ptr)->translate(static_cast<qreal>(dx), static_cast<qreal>(dy));
}

void QGraphicsView_UpdateSceneRect(void* ptr, void* rect){
	QMetaObject::invokeMethod(static_cast<QGraphicsView*>(ptr), "updateSceneRect", Q_ARG(QRectF, *static_cast<QRectF*>(rect)));
}

void QGraphicsView_DestroyQGraphicsView(void* ptr){
	static_cast<QGraphicsView*>(ptr)->~QGraphicsView();
}

