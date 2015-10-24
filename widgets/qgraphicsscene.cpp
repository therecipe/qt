#include "qgraphicsscene.h"
#include <QPixmap>
#include <QTransform>
#include <QMetaObject>
#include <QFont>
#include <QPolygon>
#include <QVariant>
#include <QPainter>
#include <QGraphicsItem>
#include <QLineF>
#include <QPointF>
#include <QLine>
#include <QEvent>
#include <QUrl>
#include <QModelIndex>
#include <QStyle>
#include <QPen>
#include <QString>
#include <QBrush>
#include <QPalette>
#include <QGraphicsWidget>
#include <QPainterPath>
#include <QGraphicsItemGroup>
#include <QObject>
#include <QRect>
#include <QPoint>
#include <QRectF>
#include <QWidget>
#include <QPolygonF>
#include <QGraphicsScene>
#include "_cgo_export.h"

class MyQGraphicsScene: public QGraphicsScene {
public:
void Signal_FocusItemChanged(QGraphicsItem * newFocusItem, QGraphicsItem * oldFocusItem, Qt::FocusReason reason){callbackQGraphicsSceneFocusItemChanged(this->objectName().toUtf8().data(), newFocusItem, oldFocusItem, reason);};
void Signal_SelectionChanged(){callbackQGraphicsSceneSelectionChanged(this->objectName().toUtf8().data());};
};

int QGraphicsScene_BspTreeDepth(QtObjectPtr ptr){
	return static_cast<QGraphicsScene*>(ptr)->bspTreeDepth();
}

int QGraphicsScene_IsSortCacheEnabled(QtObjectPtr ptr){
	return static_cast<QGraphicsScene*>(ptr)->isSortCacheEnabled();
}

int QGraphicsScene_ItemIndexMethod(QtObjectPtr ptr){
	return static_cast<QGraphicsScene*>(ptr)->itemIndexMethod();
}

void QGraphicsScene_SetBackgroundBrush(QtObjectPtr ptr, QtObjectPtr brush){
	static_cast<QGraphicsScene*>(ptr)->setBackgroundBrush(*static_cast<QBrush*>(brush));
}

void QGraphicsScene_SetBspTreeDepth(QtObjectPtr ptr, int depth){
	static_cast<QGraphicsScene*>(ptr)->setBspTreeDepth(depth);
}

void QGraphicsScene_SetFont(QtObjectPtr ptr, QtObjectPtr font){
	static_cast<QGraphicsScene*>(ptr)->setFont(*static_cast<QFont*>(font));
}

void QGraphicsScene_SetForegroundBrush(QtObjectPtr ptr, QtObjectPtr brush){
	static_cast<QGraphicsScene*>(ptr)->setForegroundBrush(*static_cast<QBrush*>(brush));
}

void QGraphicsScene_SetItemIndexMethod(QtObjectPtr ptr, int method){
	static_cast<QGraphicsScene*>(ptr)->setItemIndexMethod(static_cast<QGraphicsScene::ItemIndexMethod>(method));
}

void QGraphicsScene_SetPalette(QtObjectPtr ptr, QtObjectPtr palette){
	static_cast<QGraphicsScene*>(ptr)->setPalette(*static_cast<QPalette*>(palette));
}

void QGraphicsScene_SetSceneRect(QtObjectPtr ptr, QtObjectPtr rect){
	static_cast<QGraphicsScene*>(ptr)->setSceneRect(*static_cast<QRectF*>(rect));
}

void QGraphicsScene_SetSortCacheEnabled(QtObjectPtr ptr, int enabled){
	static_cast<QGraphicsScene*>(ptr)->setSortCacheEnabled(enabled != 0);
}

void QGraphicsScene_SetStickyFocus(QtObjectPtr ptr, int enabled){
	static_cast<QGraphicsScene*>(ptr)->setStickyFocus(enabled != 0);
}

int QGraphicsScene_StickyFocus(QtObjectPtr ptr){
	return static_cast<QGraphicsScene*>(ptr)->stickyFocus();
}

void QGraphicsScene_Update(QtObjectPtr ptr, QtObjectPtr rect){
	QMetaObject::invokeMethod(static_cast<QGraphicsScene*>(ptr), "update", Q_ARG(QRectF, *static_cast<QRectF*>(rect)));
}

QtObjectPtr QGraphicsScene_NewQGraphicsScene(QtObjectPtr parent){
	return new QGraphicsScene(static_cast<QObject*>(parent));
}

QtObjectPtr QGraphicsScene_NewQGraphicsScene2(QtObjectPtr sceneRect, QtObjectPtr parent){
	return new QGraphicsScene(*static_cast<QRectF*>(sceneRect), static_cast<QObject*>(parent));
}

QtObjectPtr QGraphicsScene_ActivePanel(QtObjectPtr ptr){
	return static_cast<QGraphicsScene*>(ptr)->activePanel();
}

QtObjectPtr QGraphicsScene_ActiveWindow(QtObjectPtr ptr){
	return static_cast<QGraphicsScene*>(ptr)->activeWindow();
}

QtObjectPtr QGraphicsScene_AddEllipse(QtObjectPtr ptr, QtObjectPtr rect, QtObjectPtr pen, QtObjectPtr brush){
	return static_cast<QGraphicsScene*>(ptr)->addEllipse(*static_cast<QRectF*>(rect), *static_cast<QPen*>(pen), *static_cast<QBrush*>(brush));
}

void QGraphicsScene_AddItem(QtObjectPtr ptr, QtObjectPtr item){
	static_cast<QGraphicsScene*>(ptr)->addItem(static_cast<QGraphicsItem*>(item));
}

QtObjectPtr QGraphicsScene_AddLine(QtObjectPtr ptr, QtObjectPtr line, QtObjectPtr pen){
	return static_cast<QGraphicsScene*>(ptr)->addLine(*static_cast<QLineF*>(line), *static_cast<QPen*>(pen));
}

QtObjectPtr QGraphicsScene_AddPath(QtObjectPtr ptr, QtObjectPtr path, QtObjectPtr pen, QtObjectPtr brush){
	return static_cast<QGraphicsScene*>(ptr)->addPath(*static_cast<QPainterPath*>(path), *static_cast<QPen*>(pen), *static_cast<QBrush*>(brush));
}

QtObjectPtr QGraphicsScene_AddPixmap(QtObjectPtr ptr, QtObjectPtr pixmap){
	return static_cast<QGraphicsScene*>(ptr)->addPixmap(*static_cast<QPixmap*>(pixmap));
}

QtObjectPtr QGraphicsScene_AddPolygon(QtObjectPtr ptr, QtObjectPtr polygon, QtObjectPtr pen, QtObjectPtr brush){
	return static_cast<QGraphicsScene*>(ptr)->addPolygon(*static_cast<QPolygonF*>(polygon), *static_cast<QPen*>(pen), *static_cast<QBrush*>(brush));
}

QtObjectPtr QGraphicsScene_AddRect(QtObjectPtr ptr, QtObjectPtr rect, QtObjectPtr pen, QtObjectPtr brush){
	return static_cast<QGraphicsScene*>(ptr)->addRect(*static_cast<QRectF*>(rect), *static_cast<QPen*>(pen), *static_cast<QBrush*>(brush));
}

QtObjectPtr QGraphicsScene_AddSimpleText(QtObjectPtr ptr, char* text, QtObjectPtr font){
	return static_cast<QGraphicsScene*>(ptr)->addSimpleText(QString(text), *static_cast<QFont*>(font));
}

QtObjectPtr QGraphicsScene_AddText(QtObjectPtr ptr, char* text, QtObjectPtr font){
	return static_cast<QGraphicsScene*>(ptr)->addText(QString(text), *static_cast<QFont*>(font));
}

QtObjectPtr QGraphicsScene_AddWidget(QtObjectPtr ptr, QtObjectPtr widget, int wFlags){
	return static_cast<QGraphicsScene*>(ptr)->addWidget(static_cast<QWidget*>(widget), static_cast<Qt::WindowType>(wFlags));
}

void QGraphicsScene_Advance(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QGraphicsScene*>(ptr), "advance");
}

void QGraphicsScene_Clear(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QGraphicsScene*>(ptr), "clear");
}

void QGraphicsScene_ClearFocus(QtObjectPtr ptr){
	static_cast<QGraphicsScene*>(ptr)->clearFocus();
}

void QGraphicsScene_ClearSelection(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QGraphicsScene*>(ptr), "clearSelection");
}

void QGraphicsScene_DestroyItemGroup(QtObjectPtr ptr, QtObjectPtr group){
	static_cast<QGraphicsScene*>(ptr)->destroyItemGroup(static_cast<QGraphicsItemGroup*>(group));
}

QtObjectPtr QGraphicsScene_FocusItem(QtObjectPtr ptr){
	return static_cast<QGraphicsScene*>(ptr)->focusItem();
}

void QGraphicsScene_ConnectFocusItemChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QGraphicsScene*>(ptr), static_cast<void (QGraphicsScene::*)(QGraphicsItem *, QGraphicsItem *, Qt::FocusReason)>(&QGraphicsScene::focusItemChanged), static_cast<MyQGraphicsScene*>(ptr), static_cast<void (MyQGraphicsScene::*)(QGraphicsItem *, QGraphicsItem *, Qt::FocusReason)>(&MyQGraphicsScene::Signal_FocusItemChanged));;
}

void QGraphicsScene_DisconnectFocusItemChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QGraphicsScene*>(ptr), static_cast<void (QGraphicsScene::*)(QGraphicsItem *, QGraphicsItem *, Qt::FocusReason)>(&QGraphicsScene::focusItemChanged), static_cast<MyQGraphicsScene*>(ptr), static_cast<void (MyQGraphicsScene::*)(QGraphicsItem *, QGraphicsItem *, Qt::FocusReason)>(&MyQGraphicsScene::Signal_FocusItemChanged));;
}

int QGraphicsScene_HasFocus(QtObjectPtr ptr){
	return static_cast<QGraphicsScene*>(ptr)->hasFocus();
}

char* QGraphicsScene_InputMethodQuery(QtObjectPtr ptr, int query){
	return static_cast<QGraphicsScene*>(ptr)->inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)).toString().toUtf8().data();
}

void QGraphicsScene_Invalidate(QtObjectPtr ptr, QtObjectPtr rect, int layers){
	QMetaObject::invokeMethod(static_cast<QGraphicsScene*>(ptr), "invalidate", Q_ARG(QRectF, *static_cast<QRectF*>(rect)), Q_ARG(QGraphicsScene::SceneLayer, static_cast<QGraphicsScene::SceneLayer>(layers)));
}

int QGraphicsScene_IsActive(QtObjectPtr ptr){
	return static_cast<QGraphicsScene*>(ptr)->isActive();
}

QtObjectPtr QGraphicsScene_ItemAt(QtObjectPtr ptr, QtObjectPtr position, QtObjectPtr deviceTransform){
	return static_cast<QGraphicsScene*>(ptr)->itemAt(*static_cast<QPointF*>(position), *static_cast<QTransform*>(deviceTransform));
}

QtObjectPtr QGraphicsScene_MouseGrabberItem(QtObjectPtr ptr){
	return static_cast<QGraphicsScene*>(ptr)->mouseGrabberItem();
}

void QGraphicsScene_RemoveItem(QtObjectPtr ptr, QtObjectPtr item){
	static_cast<QGraphicsScene*>(ptr)->removeItem(static_cast<QGraphicsItem*>(item));
}

void QGraphicsScene_Render(QtObjectPtr ptr, QtObjectPtr painter, QtObjectPtr target, QtObjectPtr source, int aspectRatioMode){
	static_cast<QGraphicsScene*>(ptr)->render(static_cast<QPainter*>(painter), *static_cast<QRectF*>(target), *static_cast<QRectF*>(source), static_cast<Qt::AspectRatioMode>(aspectRatioMode));
}

void QGraphicsScene_ConnectSelectionChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QGraphicsScene*>(ptr), static_cast<void (QGraphicsScene::*)()>(&QGraphicsScene::selectionChanged), static_cast<MyQGraphicsScene*>(ptr), static_cast<void (MyQGraphicsScene::*)()>(&MyQGraphicsScene::Signal_SelectionChanged));;
}

void QGraphicsScene_DisconnectSelectionChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QGraphicsScene*>(ptr), static_cast<void (QGraphicsScene::*)()>(&QGraphicsScene::selectionChanged), static_cast<MyQGraphicsScene*>(ptr), static_cast<void (MyQGraphicsScene::*)()>(&MyQGraphicsScene::Signal_SelectionChanged));;
}

int QGraphicsScene_SendEvent(QtObjectPtr ptr, QtObjectPtr item, QtObjectPtr event){
	return static_cast<QGraphicsScene*>(ptr)->sendEvent(static_cast<QGraphicsItem*>(item), static_cast<QEvent*>(event));
}

void QGraphicsScene_SetActivePanel(QtObjectPtr ptr, QtObjectPtr item){
	static_cast<QGraphicsScene*>(ptr)->setActivePanel(static_cast<QGraphicsItem*>(item));
}

void QGraphicsScene_SetActiveWindow(QtObjectPtr ptr, QtObjectPtr widget){
	static_cast<QGraphicsScene*>(ptr)->setActiveWindow(static_cast<QGraphicsWidget*>(widget));
}

void QGraphicsScene_SetFocus(QtObjectPtr ptr, int focusReason){
	static_cast<QGraphicsScene*>(ptr)->setFocus(static_cast<Qt::FocusReason>(focusReason));
}

void QGraphicsScene_SetFocusItem(QtObjectPtr ptr, QtObjectPtr item, int focusReason){
	static_cast<QGraphicsScene*>(ptr)->setFocusItem(static_cast<QGraphicsItem*>(item), static_cast<Qt::FocusReason>(focusReason));
}

void QGraphicsScene_SetSelectionArea2(QtObjectPtr ptr, QtObjectPtr path, int mode, QtObjectPtr deviceTransform){
	static_cast<QGraphicsScene*>(ptr)->setSelectionArea(*static_cast<QPainterPath*>(path), static_cast<Qt::ItemSelectionMode>(mode), *static_cast<QTransform*>(deviceTransform));
}

void QGraphicsScene_SetSelectionArea3(QtObjectPtr ptr, QtObjectPtr path, int selectionOperation, int mode, QtObjectPtr deviceTransform){
	static_cast<QGraphicsScene*>(ptr)->setSelectionArea(*static_cast<QPainterPath*>(path), static_cast<Qt::ItemSelectionOperation>(selectionOperation), static_cast<Qt::ItemSelectionMode>(mode), *static_cast<QTransform*>(deviceTransform));
}

void QGraphicsScene_SetSelectionArea(QtObjectPtr ptr, QtObjectPtr path, QtObjectPtr deviceTransform){
	static_cast<QGraphicsScene*>(ptr)->setSelectionArea(*static_cast<QPainterPath*>(path), *static_cast<QTransform*>(deviceTransform));
}

void QGraphicsScene_SetStyle(QtObjectPtr ptr, QtObjectPtr style){
	static_cast<QGraphicsScene*>(ptr)->setStyle(static_cast<QStyle*>(style));
}

QtObjectPtr QGraphicsScene_Style(QtObjectPtr ptr){
	return static_cast<QGraphicsScene*>(ptr)->style();
}

void QGraphicsScene_DestroyQGraphicsScene(QtObjectPtr ptr){
	static_cast<QGraphicsScene*>(ptr)->~QGraphicsScene();
}

