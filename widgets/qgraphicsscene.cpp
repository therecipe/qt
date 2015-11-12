#include "qgraphicsscene.h"
#include <QEvent>
#include <QBrush>
#include <QGraphicsItem>
#include <QString>
#include <QUrl>
#include <QPixmap>
#include <QWidget>
#include <QObject>
#include <QPointF>
#include <QPoint>
#include <QPainter>
#include <QFont>
#include <QRect>
#include <QPalette>
#include <QLine>
#include <QPolygon>
#include <QPen>
#include <QTransform>
#include <QLineF>
#include <QPolygonF>
#include <QModelIndex>
#include <QGraphicsWidget>
#include <QGraphicsItemGroup>
#include <QRectF>
#include <QStyle>
#include <QVariant>
#include <QMetaObject>
#include <QPainterPath>
#include <QGraphicsScene>
#include "_cgo_export.h"

class MyQGraphicsScene: public QGraphicsScene {
public:
void Signal_FocusItemChanged(QGraphicsItem * newFocusItem, QGraphicsItem * oldFocusItem, Qt::FocusReason reason){callbackQGraphicsSceneFocusItemChanged(this->objectName().toUtf8().data(), newFocusItem, oldFocusItem, reason);};
void Signal_SelectionChanged(){callbackQGraphicsSceneSelectionChanged(this->objectName().toUtf8().data());};
};

void* QGraphicsScene_BackgroundBrush(void* ptr){
	return new QBrush(static_cast<QGraphicsScene*>(ptr)->backgroundBrush());
}

int QGraphicsScene_BspTreeDepth(void* ptr){
	return static_cast<QGraphicsScene*>(ptr)->bspTreeDepth();
}

void* QGraphicsScene_ForegroundBrush(void* ptr){
	return new QBrush(static_cast<QGraphicsScene*>(ptr)->foregroundBrush());
}

int QGraphicsScene_IsSortCacheEnabled(void* ptr){
	return static_cast<QGraphicsScene*>(ptr)->isSortCacheEnabled();
}

int QGraphicsScene_ItemIndexMethod(void* ptr){
	return static_cast<QGraphicsScene*>(ptr)->itemIndexMethod();
}

double QGraphicsScene_MinimumRenderSize(void* ptr){
	return static_cast<double>(static_cast<QGraphicsScene*>(ptr)->minimumRenderSize());
}

void QGraphicsScene_SetBackgroundBrush(void* ptr, void* brush){
	static_cast<QGraphicsScene*>(ptr)->setBackgroundBrush(*static_cast<QBrush*>(brush));
}

void QGraphicsScene_SetBspTreeDepth(void* ptr, int depth){
	static_cast<QGraphicsScene*>(ptr)->setBspTreeDepth(depth);
}

void QGraphicsScene_SetFont(void* ptr, void* font){
	static_cast<QGraphicsScene*>(ptr)->setFont(*static_cast<QFont*>(font));
}

void QGraphicsScene_SetForegroundBrush(void* ptr, void* brush){
	static_cast<QGraphicsScene*>(ptr)->setForegroundBrush(*static_cast<QBrush*>(brush));
}

void QGraphicsScene_SetItemIndexMethod(void* ptr, int method){
	static_cast<QGraphicsScene*>(ptr)->setItemIndexMethod(static_cast<QGraphicsScene::ItemIndexMethod>(method));
}

void QGraphicsScene_SetMinimumRenderSize(void* ptr, double minSize){
	static_cast<QGraphicsScene*>(ptr)->setMinimumRenderSize(static_cast<qreal>(minSize));
}

void QGraphicsScene_SetPalette(void* ptr, void* palette){
	static_cast<QGraphicsScene*>(ptr)->setPalette(*static_cast<QPalette*>(palette));
}

void QGraphicsScene_SetSceneRect(void* ptr, void* rect){
	static_cast<QGraphicsScene*>(ptr)->setSceneRect(*static_cast<QRectF*>(rect));
}

void QGraphicsScene_SetSortCacheEnabled(void* ptr, int enabled){
	static_cast<QGraphicsScene*>(ptr)->setSortCacheEnabled(enabled != 0);
}

void QGraphicsScene_SetStickyFocus(void* ptr, int enabled){
	static_cast<QGraphicsScene*>(ptr)->setStickyFocus(enabled != 0);
}

int QGraphicsScene_StickyFocus(void* ptr){
	return static_cast<QGraphicsScene*>(ptr)->stickyFocus();
}

void QGraphicsScene_Update(void* ptr, void* rect){
	QMetaObject::invokeMethod(static_cast<QGraphicsScene*>(ptr), "update", Q_ARG(QRectF, *static_cast<QRectF*>(rect)));
}

void* QGraphicsScene_NewQGraphicsScene(void* parent){
	return new QGraphicsScene(static_cast<QObject*>(parent));
}

void* QGraphicsScene_NewQGraphicsScene2(void* sceneRect, void* parent){
	return new QGraphicsScene(*static_cast<QRectF*>(sceneRect), static_cast<QObject*>(parent));
}

void* QGraphicsScene_NewQGraphicsScene3(double x, double y, double width, double height, void* parent){
	return new QGraphicsScene(static_cast<qreal>(x), static_cast<qreal>(y), static_cast<qreal>(width), static_cast<qreal>(height), static_cast<QObject*>(parent));
}

void* QGraphicsScene_ActivePanel(void* ptr){
	return static_cast<QGraphicsScene*>(ptr)->activePanel();
}

void* QGraphicsScene_ActiveWindow(void* ptr){
	return static_cast<QGraphicsScene*>(ptr)->activeWindow();
}

void* QGraphicsScene_AddEllipse(void* ptr, void* rect, void* pen, void* brush){
	return static_cast<QGraphicsScene*>(ptr)->addEllipse(*static_cast<QRectF*>(rect), *static_cast<QPen*>(pen), *static_cast<QBrush*>(brush));
}

void* QGraphicsScene_AddEllipse2(void* ptr, double x, double y, double w, double h, void* pen, void* brush){
	return static_cast<QGraphicsScene*>(ptr)->addEllipse(static_cast<qreal>(x), static_cast<qreal>(y), static_cast<qreal>(w), static_cast<qreal>(h), *static_cast<QPen*>(pen), *static_cast<QBrush*>(brush));
}

void QGraphicsScene_AddItem(void* ptr, void* item){
	static_cast<QGraphicsScene*>(ptr)->addItem(static_cast<QGraphicsItem*>(item));
}

void* QGraphicsScene_AddLine(void* ptr, void* line, void* pen){
	return static_cast<QGraphicsScene*>(ptr)->addLine(*static_cast<QLineF*>(line), *static_cast<QPen*>(pen));
}

void* QGraphicsScene_AddLine2(void* ptr, double x1, double y1, double x2, double y2, void* pen){
	return static_cast<QGraphicsScene*>(ptr)->addLine(static_cast<qreal>(x1), static_cast<qreal>(y1), static_cast<qreal>(x2), static_cast<qreal>(y2), *static_cast<QPen*>(pen));
}

void* QGraphicsScene_AddPath(void* ptr, void* path, void* pen, void* brush){
	return static_cast<QGraphicsScene*>(ptr)->addPath(*static_cast<QPainterPath*>(path), *static_cast<QPen*>(pen), *static_cast<QBrush*>(brush));
}

void* QGraphicsScene_AddPixmap(void* ptr, void* pixmap){
	return static_cast<QGraphicsScene*>(ptr)->addPixmap(*static_cast<QPixmap*>(pixmap));
}

void* QGraphicsScene_AddPolygon(void* ptr, void* polygon, void* pen, void* brush){
	return static_cast<QGraphicsScene*>(ptr)->addPolygon(*static_cast<QPolygonF*>(polygon), *static_cast<QPen*>(pen), *static_cast<QBrush*>(brush));
}

void* QGraphicsScene_AddRect(void* ptr, void* rect, void* pen, void* brush){
	return static_cast<QGraphicsScene*>(ptr)->addRect(*static_cast<QRectF*>(rect), *static_cast<QPen*>(pen), *static_cast<QBrush*>(brush));
}

void* QGraphicsScene_AddRect2(void* ptr, double x, double y, double w, double h, void* pen, void* brush){
	return static_cast<QGraphicsScene*>(ptr)->addRect(static_cast<qreal>(x), static_cast<qreal>(y), static_cast<qreal>(w), static_cast<qreal>(h), *static_cast<QPen*>(pen), *static_cast<QBrush*>(brush));
}

void* QGraphicsScene_AddSimpleText(void* ptr, char* text, void* font){
	return static_cast<QGraphicsScene*>(ptr)->addSimpleText(QString(text), *static_cast<QFont*>(font));
}

void* QGraphicsScene_AddText(void* ptr, char* text, void* font){
	return static_cast<QGraphicsScene*>(ptr)->addText(QString(text), *static_cast<QFont*>(font));
}

void* QGraphicsScene_AddWidget(void* ptr, void* widget, int wFlags){
	return static_cast<QGraphicsScene*>(ptr)->addWidget(static_cast<QWidget*>(widget), static_cast<Qt::WindowType>(wFlags));
}

void QGraphicsScene_Advance(void* ptr){
	QMetaObject::invokeMethod(static_cast<QGraphicsScene*>(ptr), "advance");
}

void QGraphicsScene_Clear(void* ptr){
	QMetaObject::invokeMethod(static_cast<QGraphicsScene*>(ptr), "clear");
}

void QGraphicsScene_ClearFocus(void* ptr){
	static_cast<QGraphicsScene*>(ptr)->clearFocus();
}

void QGraphicsScene_ClearSelection(void* ptr){
	QMetaObject::invokeMethod(static_cast<QGraphicsScene*>(ptr), "clearSelection");
}

void QGraphicsScene_DestroyItemGroup(void* ptr, void* group){
	static_cast<QGraphicsScene*>(ptr)->destroyItemGroup(static_cast<QGraphicsItemGroup*>(group));
}

void* QGraphicsScene_FocusItem(void* ptr){
	return static_cast<QGraphicsScene*>(ptr)->focusItem();
}

void QGraphicsScene_ConnectFocusItemChanged(void* ptr){
	QObject::connect(static_cast<QGraphicsScene*>(ptr), static_cast<void (QGraphicsScene::*)(QGraphicsItem *, QGraphicsItem *, Qt::FocusReason)>(&QGraphicsScene::focusItemChanged), static_cast<MyQGraphicsScene*>(ptr), static_cast<void (MyQGraphicsScene::*)(QGraphicsItem *, QGraphicsItem *, Qt::FocusReason)>(&MyQGraphicsScene::Signal_FocusItemChanged));;
}

void QGraphicsScene_DisconnectFocusItemChanged(void* ptr){
	QObject::disconnect(static_cast<QGraphicsScene*>(ptr), static_cast<void (QGraphicsScene::*)(QGraphicsItem *, QGraphicsItem *, Qt::FocusReason)>(&QGraphicsScene::focusItemChanged), static_cast<MyQGraphicsScene*>(ptr), static_cast<void (MyQGraphicsScene::*)(QGraphicsItem *, QGraphicsItem *, Qt::FocusReason)>(&MyQGraphicsScene::Signal_FocusItemChanged));;
}

int QGraphicsScene_HasFocus(void* ptr){
	return static_cast<QGraphicsScene*>(ptr)->hasFocus();
}

double QGraphicsScene_Height(void* ptr){
	return static_cast<double>(static_cast<QGraphicsScene*>(ptr)->height());
}

void* QGraphicsScene_InputMethodQuery(void* ptr, int query){
	return new QVariant(static_cast<QGraphicsScene*>(ptr)->inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

void QGraphicsScene_Invalidate(void* ptr, void* rect, int layers){
	QMetaObject::invokeMethod(static_cast<QGraphicsScene*>(ptr), "invalidate", Q_ARG(QRectF, *static_cast<QRectF*>(rect)), Q_ARG(QGraphicsScene::SceneLayer, static_cast<QGraphicsScene::SceneLayer>(layers)));
}

void QGraphicsScene_Invalidate2(void* ptr, double x, double y, double w, double h, int layers){
	static_cast<QGraphicsScene*>(ptr)->invalidate(static_cast<qreal>(x), static_cast<qreal>(y), static_cast<qreal>(w), static_cast<qreal>(h), static_cast<QGraphicsScene::SceneLayer>(layers));
}

int QGraphicsScene_IsActive(void* ptr){
	return static_cast<QGraphicsScene*>(ptr)->isActive();
}

void* QGraphicsScene_ItemAt(void* ptr, void* position, void* deviceTransform){
	return static_cast<QGraphicsScene*>(ptr)->itemAt(*static_cast<QPointF*>(position), *static_cast<QTransform*>(deviceTransform));
}

void* QGraphicsScene_ItemAt3(void* ptr, double x, double y, void* deviceTransform){
	return static_cast<QGraphicsScene*>(ptr)->itemAt(static_cast<qreal>(x), static_cast<qreal>(y), *static_cast<QTransform*>(deviceTransform));
}

void* QGraphicsScene_MouseGrabberItem(void* ptr){
	return static_cast<QGraphicsScene*>(ptr)->mouseGrabberItem();
}

void QGraphicsScene_RemoveItem(void* ptr, void* item){
	static_cast<QGraphicsScene*>(ptr)->removeItem(static_cast<QGraphicsItem*>(item));
}

void QGraphicsScene_Render(void* ptr, void* painter, void* target, void* source, int aspectRatioMode){
	static_cast<QGraphicsScene*>(ptr)->render(static_cast<QPainter*>(painter), *static_cast<QRectF*>(target), *static_cast<QRectF*>(source), static_cast<Qt::AspectRatioMode>(aspectRatioMode));
}

void QGraphicsScene_ConnectSelectionChanged(void* ptr){
	QObject::connect(static_cast<QGraphicsScene*>(ptr), static_cast<void (QGraphicsScene::*)()>(&QGraphicsScene::selectionChanged), static_cast<MyQGraphicsScene*>(ptr), static_cast<void (MyQGraphicsScene::*)()>(&MyQGraphicsScene::Signal_SelectionChanged));;
}

void QGraphicsScene_DisconnectSelectionChanged(void* ptr){
	QObject::disconnect(static_cast<QGraphicsScene*>(ptr), static_cast<void (QGraphicsScene::*)()>(&QGraphicsScene::selectionChanged), static_cast<MyQGraphicsScene*>(ptr), static_cast<void (MyQGraphicsScene::*)()>(&MyQGraphicsScene::Signal_SelectionChanged));;
}

int QGraphicsScene_SendEvent(void* ptr, void* item, void* event){
	return static_cast<QGraphicsScene*>(ptr)->sendEvent(static_cast<QGraphicsItem*>(item), static_cast<QEvent*>(event));
}

void QGraphicsScene_SetActivePanel(void* ptr, void* item){
	static_cast<QGraphicsScene*>(ptr)->setActivePanel(static_cast<QGraphicsItem*>(item));
}

void QGraphicsScene_SetActiveWindow(void* ptr, void* widget){
	static_cast<QGraphicsScene*>(ptr)->setActiveWindow(static_cast<QGraphicsWidget*>(widget));
}

void QGraphicsScene_SetFocus(void* ptr, int focusReason){
	static_cast<QGraphicsScene*>(ptr)->setFocus(static_cast<Qt::FocusReason>(focusReason));
}

void QGraphicsScene_SetFocusItem(void* ptr, void* item, int focusReason){
	static_cast<QGraphicsScene*>(ptr)->setFocusItem(static_cast<QGraphicsItem*>(item), static_cast<Qt::FocusReason>(focusReason));
}

void QGraphicsScene_SetSceneRect2(void* ptr, double x, double y, double w, double h){
	static_cast<QGraphicsScene*>(ptr)->setSceneRect(static_cast<qreal>(x), static_cast<qreal>(y), static_cast<qreal>(w), static_cast<qreal>(h));
}

void QGraphicsScene_SetSelectionArea2(void* ptr, void* path, int mode, void* deviceTransform){
	static_cast<QGraphicsScene*>(ptr)->setSelectionArea(*static_cast<QPainterPath*>(path), static_cast<Qt::ItemSelectionMode>(mode), *static_cast<QTransform*>(deviceTransform));
}

void QGraphicsScene_SetSelectionArea3(void* ptr, void* path, int selectionOperation, int mode, void* deviceTransform){
	static_cast<QGraphicsScene*>(ptr)->setSelectionArea(*static_cast<QPainterPath*>(path), static_cast<Qt::ItemSelectionOperation>(selectionOperation), static_cast<Qt::ItemSelectionMode>(mode), *static_cast<QTransform*>(deviceTransform));
}

void QGraphicsScene_SetSelectionArea(void* ptr, void* path, void* deviceTransform){
	static_cast<QGraphicsScene*>(ptr)->setSelectionArea(*static_cast<QPainterPath*>(path), *static_cast<QTransform*>(deviceTransform));
}

void QGraphicsScene_SetStyle(void* ptr, void* style){
	static_cast<QGraphicsScene*>(ptr)->setStyle(static_cast<QStyle*>(style));
}

void* QGraphicsScene_Style(void* ptr){
	return static_cast<QGraphicsScene*>(ptr)->style();
}

void QGraphicsScene_Update2(void* ptr, double x, double y, double w, double h){
	static_cast<QGraphicsScene*>(ptr)->update(static_cast<qreal>(x), static_cast<qreal>(y), static_cast<qreal>(w), static_cast<qreal>(h));
}

double QGraphicsScene_Width(void* ptr){
	return static_cast<double>(static_cast<QGraphicsScene*>(ptr)->width());
}

void QGraphicsScene_DestroyQGraphicsScene(void* ptr){
	static_cast<QGraphicsScene*>(ptr)->~QGraphicsScene();
}

