#include "qgraphicsview.h"
#include <QModelIndex>
#include <QTransform>
#include <QString>
#include <QVariant>
#include <QRect>
#include <QRectF>
#include <QPointF>
#include <QGraphicsItem>
#include <QWidget>
#include <QGraphicsScene>
#include <QUrl>
#include <QMetaObject>
#include <QPainter>
#include <QBrush>
#include <QPoint>
#include <QGraphicsView>
#include "_cgo_export.h"

class MyQGraphicsView: public QGraphicsView {
public:
};

int QGraphicsView_Alignment(QtObjectPtr ptr){
	return static_cast<QGraphicsView*>(ptr)->alignment();
}

int QGraphicsView_CacheMode(QtObjectPtr ptr){
	return static_cast<QGraphicsView*>(ptr)->cacheMode();
}

int QGraphicsView_DragMode(QtObjectPtr ptr){
	return static_cast<QGraphicsView*>(ptr)->dragMode();
}

int QGraphicsView_IsInteractive(QtObjectPtr ptr){
	return static_cast<QGraphicsView*>(ptr)->isInteractive();
}

int QGraphicsView_OptimizationFlags(QtObjectPtr ptr){
	return static_cast<QGraphicsView*>(ptr)->optimizationFlags();
}

int QGraphicsView_RenderHints(QtObjectPtr ptr){
	return static_cast<QGraphicsView*>(ptr)->renderHints();
}

int QGraphicsView_ResizeAnchor(QtObjectPtr ptr){
	return static_cast<QGraphicsView*>(ptr)->resizeAnchor();
}

int QGraphicsView_RubberBandSelectionMode(QtObjectPtr ptr){
	return static_cast<QGraphicsView*>(ptr)->rubberBandSelectionMode();
}

void QGraphicsView_SetAlignment(QtObjectPtr ptr, int alignment){
	static_cast<QGraphicsView*>(ptr)->setAlignment(static_cast<Qt::AlignmentFlag>(alignment));
}

void QGraphicsView_SetBackgroundBrush(QtObjectPtr ptr, QtObjectPtr brush){
	static_cast<QGraphicsView*>(ptr)->setBackgroundBrush(*static_cast<QBrush*>(brush));
}

void QGraphicsView_SetCacheMode(QtObjectPtr ptr, int mode){
	static_cast<QGraphicsView*>(ptr)->setCacheMode(static_cast<QGraphicsView::CacheModeFlag>(mode));
}

void QGraphicsView_SetDragMode(QtObjectPtr ptr, int mode){
	static_cast<QGraphicsView*>(ptr)->setDragMode(static_cast<QGraphicsView::DragMode>(mode));
}

void QGraphicsView_SetForegroundBrush(QtObjectPtr ptr, QtObjectPtr brush){
	static_cast<QGraphicsView*>(ptr)->setForegroundBrush(*static_cast<QBrush*>(brush));
}

void QGraphicsView_SetInteractive(QtObjectPtr ptr, int allowed){
	static_cast<QGraphicsView*>(ptr)->setInteractive(allowed != 0);
}

void QGraphicsView_SetOptimizationFlags(QtObjectPtr ptr, int flags){
	static_cast<QGraphicsView*>(ptr)->setOptimizationFlags(static_cast<QGraphicsView::OptimizationFlag>(flags));
}

void QGraphicsView_SetRenderHints(QtObjectPtr ptr, int hints){
	static_cast<QGraphicsView*>(ptr)->setRenderHints(static_cast<QPainter::RenderHint>(hints));
}

void QGraphicsView_SetResizeAnchor(QtObjectPtr ptr, int anchor){
	static_cast<QGraphicsView*>(ptr)->setResizeAnchor(static_cast<QGraphicsView::ViewportAnchor>(anchor));
}

void QGraphicsView_SetRubberBandSelectionMode(QtObjectPtr ptr, int mode){
	static_cast<QGraphicsView*>(ptr)->setRubberBandSelectionMode(static_cast<Qt::ItemSelectionMode>(mode));
}

void QGraphicsView_SetSceneRect(QtObjectPtr ptr, QtObjectPtr rect){
	static_cast<QGraphicsView*>(ptr)->setSceneRect(*static_cast<QRectF*>(rect));
}

void QGraphicsView_SetTransformationAnchor(QtObjectPtr ptr, int anchor){
	static_cast<QGraphicsView*>(ptr)->setTransformationAnchor(static_cast<QGraphicsView::ViewportAnchor>(anchor));
}

void QGraphicsView_SetViewportUpdateMode(QtObjectPtr ptr, int mode){
	static_cast<QGraphicsView*>(ptr)->setViewportUpdateMode(static_cast<QGraphicsView::ViewportUpdateMode>(mode));
}

int QGraphicsView_TransformationAnchor(QtObjectPtr ptr){
	return static_cast<QGraphicsView*>(ptr)->transformationAnchor();
}

int QGraphicsView_ViewportUpdateMode(QtObjectPtr ptr){
	return static_cast<QGraphicsView*>(ptr)->viewportUpdateMode();
}

QtObjectPtr QGraphicsView_NewQGraphicsView2(QtObjectPtr scene, QtObjectPtr parent){
	return new QGraphicsView(static_cast<QGraphicsScene*>(scene), static_cast<QWidget*>(parent));
}

QtObjectPtr QGraphicsView_NewQGraphicsView(QtObjectPtr parent){
	return new QGraphicsView(static_cast<QWidget*>(parent));
}

void QGraphicsView_CenterOn3(QtObjectPtr ptr, QtObjectPtr item){
	static_cast<QGraphicsView*>(ptr)->centerOn(static_cast<QGraphicsItem*>(item));
}

void QGraphicsView_CenterOn(QtObjectPtr ptr, QtObjectPtr pos){
	static_cast<QGraphicsView*>(ptr)->centerOn(*static_cast<QPointF*>(pos));
}

void QGraphicsView_EnsureVisible3(QtObjectPtr ptr, QtObjectPtr item, int xmargin, int ymargin){
	static_cast<QGraphicsView*>(ptr)->ensureVisible(static_cast<QGraphicsItem*>(item), xmargin, ymargin);
}

void QGraphicsView_EnsureVisible(QtObjectPtr ptr, QtObjectPtr rect, int xmargin, int ymargin){
	static_cast<QGraphicsView*>(ptr)->ensureVisible(*static_cast<QRectF*>(rect), xmargin, ymargin);
}

void QGraphicsView_FitInView3(QtObjectPtr ptr, QtObjectPtr item, int aspectRatioMode){
	static_cast<QGraphicsView*>(ptr)->fitInView(static_cast<QGraphicsItem*>(item), static_cast<Qt::AspectRatioMode>(aspectRatioMode));
}

void QGraphicsView_FitInView(QtObjectPtr ptr, QtObjectPtr rect, int aspectRatioMode){
	static_cast<QGraphicsView*>(ptr)->fitInView(*static_cast<QRectF*>(rect), static_cast<Qt::AspectRatioMode>(aspectRatioMode));
}

char* QGraphicsView_InputMethodQuery(QtObjectPtr ptr, int query){
	return static_cast<QGraphicsView*>(ptr)->inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)).toString().toUtf8().data();
}

void QGraphicsView_InvalidateScene(QtObjectPtr ptr, QtObjectPtr rect, int layers){
	QMetaObject::invokeMethod(static_cast<QGraphicsView*>(ptr), "invalidateScene", Q_ARG(QRectF, *static_cast<QRectF*>(rect)), Q_ARG(QGraphicsScene::SceneLayer, static_cast<QGraphicsScene::SceneLayer>(layers)));
}

int QGraphicsView_IsTransformed(QtObjectPtr ptr){
	return static_cast<QGraphicsView*>(ptr)->isTransformed();
}

QtObjectPtr QGraphicsView_ItemAt(QtObjectPtr ptr, QtObjectPtr pos){
	return static_cast<QGraphicsView*>(ptr)->itemAt(*static_cast<QPoint*>(pos));
}

QtObjectPtr QGraphicsView_ItemAt2(QtObjectPtr ptr, int x, int y){
	return static_cast<QGraphicsView*>(ptr)->itemAt(x, y);
}

void QGraphicsView_Render(QtObjectPtr ptr, QtObjectPtr painter, QtObjectPtr target, QtObjectPtr source, int aspectRatioMode){
	static_cast<QGraphicsView*>(ptr)->render(static_cast<QPainter*>(painter), *static_cast<QRectF*>(target), *static_cast<QRect*>(source), static_cast<Qt::AspectRatioMode>(aspectRatioMode));
}

void QGraphicsView_ResetCachedContent(QtObjectPtr ptr){
	static_cast<QGraphicsView*>(ptr)->resetCachedContent();
}

void QGraphicsView_ResetMatrix(QtObjectPtr ptr){
	static_cast<QGraphicsView*>(ptr)->resetMatrix();
}

void QGraphicsView_ResetTransform(QtObjectPtr ptr){
	static_cast<QGraphicsView*>(ptr)->resetTransform();
}

QtObjectPtr QGraphicsView_Scene(QtObjectPtr ptr){
	return static_cast<QGraphicsView*>(ptr)->scene();
}

void QGraphicsView_SetOptimizationFlag(QtObjectPtr ptr, int flag, int enabled){
	static_cast<QGraphicsView*>(ptr)->setOptimizationFlag(static_cast<QGraphicsView::OptimizationFlag>(flag), enabled != 0);
}

void QGraphicsView_SetRenderHint(QtObjectPtr ptr, int hint, int enabled){
	static_cast<QGraphicsView*>(ptr)->setRenderHint(static_cast<QPainter::RenderHint>(hint), enabled != 0);
}

void QGraphicsView_SetScene(QtObjectPtr ptr, QtObjectPtr scene){
	static_cast<QGraphicsView*>(ptr)->setScene(static_cast<QGraphicsScene*>(scene));
}

void QGraphicsView_SetTransform(QtObjectPtr ptr, QtObjectPtr matrix, int combine){
	static_cast<QGraphicsView*>(ptr)->setTransform(*static_cast<QTransform*>(matrix), combine != 0);
}

void QGraphicsView_UpdateSceneRect(QtObjectPtr ptr, QtObjectPtr rect){
	QMetaObject::invokeMethod(static_cast<QGraphicsView*>(ptr), "updateSceneRect", Q_ARG(QRectF, *static_cast<QRectF*>(rect)));
}

void QGraphicsView_DestroyQGraphicsView(QtObjectPtr ptr){
	static_cast<QGraphicsView*>(ptr)->~QGraphicsView();
}

