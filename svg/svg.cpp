// +build !minimal

#define protected public
#define private public

#include "svg.h"
#include "_cgo_export.h"

#include <QAction>
#include <QActionEvent>
#include <QByteArray>
#include <QChildEvent>
#include <QCloseEvent>
#include <QContextMenuEvent>
#include <QDrag>
#include <QDragEnterEvent>
#include <QDragLeaveEvent>
#include <QDragMoveEvent>
#include <QDropEvent>
#include <QEvent>
#include <QFocusEvent>
#include <QGraphicsItem>
#include <QGraphicsScene>
#include <QGraphicsSceneContextMenuEvent>
#include <QGraphicsSceneDragDropEvent>
#include <QGraphicsSceneHoverEvent>
#include <QGraphicsSceneMouseEvent>
#include <QGraphicsSceneWheelEvent>
#include <QGraphicsSvgItem>
#include <QHideEvent>
#include <QIODevice>
#include <QInputMethod>
#include <QInputMethodEvent>
#include <QKeyEvent>
#include <QMetaMethod>
#include <QMetaObject>
#include <QMouseEvent>
#include <QMoveEvent>
#include <QObject>
#include <QPaintDevice>
#include <QPaintEvent>
#include <QPainter>
#include <QPainterPath>
#include <QPoint>
#include <QPointF>
#include <QRect>
#include <QRectF>
#include <QResizeEvent>
#include <QShowEvent>
#include <QSize>
#include <QString>
#include <QStyle>
#include <QStyleOption>
#include <QStyleOptionGraphicsItem>
#include <QSvgGenerator>
#include <QSvgRenderer>
#include <QSvgWidget>
#include <QTabletEvent>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>
#include <QVariant>
#include <QWheelEvent>
#include <QWidget>
#include <QXmlStreamReader>

class MyQGraphicsSvgItem: public QGraphicsSvgItem
{
public:
	MyQGraphicsSvgItem(QGraphicsItem *parent) : QGraphicsSvgItem(parent) {};
	MyQGraphicsSvgItem(const QString &fileName, QGraphicsItem *parent) : QGraphicsSvgItem(fileName, parent) {};
	QRectF boundingRect() const { return *static_cast<QRectF*>(callbackQGraphicsSvgItem_BoundingRect(const_cast<MyQGraphicsSvgItem*>(this))); };
	void paint(QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget) { callbackQGraphicsSvgItem_Paint(this, painter, const_cast<QStyleOptionGraphicsItem*>(option), widget); };
	int type() const { return callbackQGraphicsSvgItem_Type(const_cast<MyQGraphicsSvgItem*>(this)); };
	void updateMicroFocus() { callbackQGraphicsSvgItem_UpdateMicroFocus(this); };
	void timerEvent(QTimerEvent * event) { callbackQGraphicsSvgItem_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQGraphicsSvgItem_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQGraphicsSvgItem_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQGraphicsSvgItem_CustomEvent(this, event); };
	void deleteLater() { callbackQGraphicsSvgItem_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQGraphicsSvgItem_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQGraphicsSvgItem_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQGraphicsSvgItem_MetaObject(const_cast<MyQGraphicsSvgItem*>(this))); };
	void advance(int phase) { callbackQGraphicsSvgItem_Advance(this, phase); };
	bool collidesWithItem(const QGraphicsItem * other, Qt::ItemSelectionMode mode) const { return callbackQGraphicsSvgItem_CollidesWithItem(const_cast<MyQGraphicsSvgItem*>(this), const_cast<QGraphicsItem*>(other), mode) != 0; };
	bool collidesWithPath(const QPainterPath & path, Qt::ItemSelectionMode mode) const { return callbackQGraphicsSvgItem_CollidesWithPath(const_cast<MyQGraphicsSvgItem*>(this), const_cast<QPainterPath*>(&path), mode) != 0; };
	bool contains(const QPointF & point) const { return callbackQGraphicsSvgItem_Contains(const_cast<MyQGraphicsSvgItem*>(this), const_cast<QPointF*>(&point)) != 0; };
	void contextMenuEvent(QGraphicsSceneContextMenuEvent * event) { callbackQGraphicsSvgItem_ContextMenuEvent(this, event); };
	void dragEnterEvent(QGraphicsSceneDragDropEvent * event) { callbackQGraphicsSvgItem_DragEnterEvent(this, event); };
	void dragLeaveEvent(QGraphicsSceneDragDropEvent * event) { callbackQGraphicsSvgItem_DragLeaveEvent(this, event); };
	void dragMoveEvent(QGraphicsSceneDragDropEvent * event) { callbackQGraphicsSvgItem_DragMoveEvent(this, event); };
	void dropEvent(QGraphicsSceneDragDropEvent * event) { callbackQGraphicsSvgItem_DropEvent(this, event); };
	void focusInEvent(QFocusEvent * event) { callbackQGraphicsSvgItem_FocusInEvent(this, event); };
	void focusOutEvent(QFocusEvent * event) { callbackQGraphicsSvgItem_FocusOutEvent(this, event); };
	void hoverEnterEvent(QGraphicsSceneHoverEvent * event) { callbackQGraphicsSvgItem_HoverEnterEvent(this, event); };
	void hoverLeaveEvent(QGraphicsSceneHoverEvent * event) { callbackQGraphicsSvgItem_HoverLeaveEvent(this, event); };
	void hoverMoveEvent(QGraphicsSceneHoverEvent * event) { callbackQGraphicsSvgItem_HoverMoveEvent(this, event); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackQGraphicsSvgItem_InputMethodEvent(this, event); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackQGraphicsSvgItem_InputMethodQuery(const_cast<MyQGraphicsSvgItem*>(this), query)); };
	bool isObscuredBy(const QGraphicsItem * item) const { return callbackQGraphicsSvgItem_IsObscuredBy(const_cast<MyQGraphicsSvgItem*>(this), const_cast<QGraphicsItem*>(item)) != 0; };
	QVariant itemChange(QGraphicsItem::GraphicsItemChange change, const QVariant & value) { return *static_cast<QVariant*>(callbackQGraphicsSvgItem_ItemChange(this, change, const_cast<QVariant*>(&value))); };
	void keyPressEvent(QKeyEvent * event) { callbackQGraphicsSvgItem_KeyPressEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQGraphicsSvgItem_KeyReleaseEvent(this, event); };
	void mouseDoubleClickEvent(QGraphicsSceneMouseEvent * event) { callbackQGraphicsSvgItem_MouseDoubleClickEvent(this, event); };
	void mouseMoveEvent(QGraphicsSceneMouseEvent * event) { callbackQGraphicsSvgItem_MouseMoveEvent(this, event); };
	void mousePressEvent(QGraphicsSceneMouseEvent * event) { callbackQGraphicsSvgItem_MousePressEvent(this, event); };
	void mouseReleaseEvent(QGraphicsSceneMouseEvent * event) { callbackQGraphicsSvgItem_MouseReleaseEvent(this, event); };
	QPainterPath opaqueArea() const { return *static_cast<QPainterPath*>(callbackQGraphicsSvgItem_OpaqueArea(const_cast<MyQGraphicsSvgItem*>(this))); };
	bool sceneEvent(QEvent * event) { return callbackQGraphicsSvgItem_SceneEvent(this, event) != 0; };
	bool sceneEventFilter(QGraphicsItem * watched, QEvent * event) { return callbackQGraphicsSvgItem_SceneEventFilter(this, watched, event) != 0; };
	QPainterPath shape() const { return *static_cast<QPainterPath*>(callbackQGraphicsSvgItem_Shape(const_cast<MyQGraphicsSvgItem*>(this))); };
	void wheelEvent(QGraphicsSceneWheelEvent * event) { callbackQGraphicsSvgItem_WheelEvent(this, event); };
};

void* QGraphicsSvgItem_NewQGraphicsSvgItem(void* parent)
{
	return new MyQGraphicsSvgItem(static_cast<QGraphicsItem*>(parent));
}

void* QGraphicsSvgItem_NewQGraphicsSvgItem2(char* fileName, void* parent)
{
	return new MyQGraphicsSvgItem(QString(fileName), static_cast<QGraphicsItem*>(parent));
}

void* QGraphicsSvgItem_BoundingRect(void* ptr)
{
	return ({ QRectF tmpValue = static_cast<QGraphicsSvgItem*>(ptr)->boundingRect(); new QRectF(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void* QGraphicsSvgItem_BoundingRectDefault(void* ptr)
{
	return ({ QRectF tmpValue = static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::boundingRect(); new QRectF(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

struct QtSvg_PackedString QGraphicsSvgItem_ElementId(void* ptr)
{
	return ({ QByteArray t96b361 = static_cast<QGraphicsSvgItem*>(ptr)->elementId().toUtf8(); QtSvg_PackedString { const_cast<char*>(t96b361.prepend("WHITESPACE").constData()+10), t96b361.size()-10 }; });
}

void* QGraphicsSvgItem_MaximumCacheSize(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QGraphicsSvgItem*>(ptr)->maximumCacheSize(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void QGraphicsSvgItem_Paint(void* ptr, void* painter, void* option, void* widget)
{
	static_cast<QGraphicsSvgItem*>(ptr)->paint(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
}

void QGraphicsSvgItem_PaintDefault(void* ptr, void* painter, void* option, void* widget)
{
	static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::paint(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
}

void* QGraphicsSvgItem_Renderer(void* ptr)
{
	return static_cast<QGraphicsSvgItem*>(ptr)->renderer();
}

void QGraphicsSvgItem_SetElementId(void* ptr, char* id)
{
	static_cast<QGraphicsSvgItem*>(ptr)->setElementId(QString(id));
}

void QGraphicsSvgItem_SetMaximumCacheSize(void* ptr, void* size)
{
	static_cast<QGraphicsSvgItem*>(ptr)->setMaximumCacheSize(*static_cast<QSize*>(size));
}

void QGraphicsSvgItem_SetSharedRenderer(void* ptr, void* renderer)
{
	static_cast<QGraphicsSvgItem*>(ptr)->setSharedRenderer(static_cast<QSvgRenderer*>(renderer));
}

int QGraphicsSvgItem_Type(void* ptr)
{
	return static_cast<QGraphicsSvgItem*>(ptr)->type();
}

int QGraphicsSvgItem_TypeDefault(void* ptr)
{
	return static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::type();
}

void QGraphicsSvgItem_UpdateMicroFocus(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QGraphicsSvgItem*>(ptr), "updateMicroFocus");
}

void QGraphicsSvgItem_UpdateMicroFocusDefault(void* ptr)
{
	static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::updateMicroFocus();
}

void QGraphicsSvgItem_TimerEvent(void* ptr, void* event)
{
	static_cast<QGraphicsSvgItem*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QGraphicsSvgItem_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::timerEvent(static_cast<QTimerEvent*>(event));
}

void QGraphicsSvgItem_ChildEvent(void* ptr, void* event)
{
	static_cast<QGraphicsSvgItem*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QGraphicsSvgItem_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::childEvent(static_cast<QChildEvent*>(event));
}

void QGraphicsSvgItem_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QGraphicsSvgItem*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGraphicsSvgItem_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGraphicsSvgItem_CustomEvent(void* ptr, void* event)
{
	static_cast<QGraphicsSvgItem*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QGraphicsSvgItem_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::customEvent(static_cast<QEvent*>(event));
}

void QGraphicsSvgItem_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QGraphicsSvgItem*>(ptr), "deleteLater");
}

void QGraphicsSvgItem_DeleteLaterDefault(void* ptr)
{
	static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::deleteLater();
}

void QGraphicsSvgItem_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QGraphicsSvgItem*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGraphicsSvgItem_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QGraphicsSvgItem_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QGraphicsSvgItem*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QGraphicsSvgItem_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QGraphicsSvgItem_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QGraphicsSvgItem*>(ptr)->metaObject());
}

void* QGraphicsSvgItem_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::metaObject());
}

void QGraphicsSvgItem_Advance(void* ptr, int phase)
{
	static_cast<QGraphicsSvgItem*>(ptr)->advance(phase);
}

void QGraphicsSvgItem_AdvanceDefault(void* ptr, int phase)
{
	static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::advance(phase);
}

char QGraphicsSvgItem_CollidesWithItem(void* ptr, void* other, long long mode)
{
	return static_cast<QGraphicsSvgItem*>(ptr)->collidesWithItem(static_cast<QGraphicsItem*>(other), static_cast<Qt::ItemSelectionMode>(mode));
}

char QGraphicsSvgItem_CollidesWithItemDefault(void* ptr, void* other, long long mode)
{
	return static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::collidesWithItem(static_cast<QGraphicsItem*>(other), static_cast<Qt::ItemSelectionMode>(mode));
}

char QGraphicsSvgItem_CollidesWithPath(void* ptr, void* path, long long mode)
{
	return static_cast<QGraphicsSvgItem*>(ptr)->collidesWithPath(*static_cast<QPainterPath*>(path), static_cast<Qt::ItemSelectionMode>(mode));
}

char QGraphicsSvgItem_CollidesWithPathDefault(void* ptr, void* path, long long mode)
{
	return static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::collidesWithPath(*static_cast<QPainterPath*>(path), static_cast<Qt::ItemSelectionMode>(mode));
}

char QGraphicsSvgItem_Contains(void* ptr, void* point)
{
	return static_cast<QGraphicsSvgItem*>(ptr)->contains(*static_cast<QPointF*>(point));
}

char QGraphicsSvgItem_ContainsDefault(void* ptr, void* point)
{
	return static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::contains(*static_cast<QPointF*>(point));
}

void QGraphicsSvgItem_ContextMenuEvent(void* ptr, void* event)
{
	static_cast<QGraphicsSvgItem*>(ptr)->contextMenuEvent(static_cast<QGraphicsSceneContextMenuEvent*>(event));
}

void QGraphicsSvgItem_ContextMenuEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::contextMenuEvent(static_cast<QGraphicsSceneContextMenuEvent*>(event));
}

void QGraphicsSvgItem_DragEnterEvent(void* ptr, void* event)
{
	static_cast<QGraphicsSvgItem*>(ptr)->dragEnterEvent(static_cast<QGraphicsSceneDragDropEvent*>(event));
}

void QGraphicsSvgItem_DragEnterEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::dragEnterEvent(static_cast<QGraphicsSceneDragDropEvent*>(event));
}

void QGraphicsSvgItem_DragLeaveEvent(void* ptr, void* event)
{
	static_cast<QGraphicsSvgItem*>(ptr)->dragLeaveEvent(static_cast<QGraphicsSceneDragDropEvent*>(event));
}

void QGraphicsSvgItem_DragLeaveEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::dragLeaveEvent(static_cast<QGraphicsSceneDragDropEvent*>(event));
}

void QGraphicsSvgItem_DragMoveEvent(void* ptr, void* event)
{
	static_cast<QGraphicsSvgItem*>(ptr)->dragMoveEvent(static_cast<QGraphicsSceneDragDropEvent*>(event));
}

void QGraphicsSvgItem_DragMoveEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::dragMoveEvent(static_cast<QGraphicsSceneDragDropEvent*>(event));
}

void QGraphicsSvgItem_DropEvent(void* ptr, void* event)
{
	static_cast<QGraphicsSvgItem*>(ptr)->dropEvent(static_cast<QGraphicsSceneDragDropEvent*>(event));
}

void QGraphicsSvgItem_DropEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::dropEvent(static_cast<QGraphicsSceneDragDropEvent*>(event));
}

void QGraphicsSvgItem_FocusInEvent(void* ptr, void* event)
{
	static_cast<QGraphicsSvgItem*>(ptr)->focusInEvent(static_cast<QFocusEvent*>(event));
}

void QGraphicsSvgItem_FocusInEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::focusInEvent(static_cast<QFocusEvent*>(event));
}

void QGraphicsSvgItem_FocusOutEvent(void* ptr, void* event)
{
	static_cast<QGraphicsSvgItem*>(ptr)->focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QGraphicsSvgItem_FocusOutEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QGraphicsSvgItem_HoverEnterEvent(void* ptr, void* event)
{
	static_cast<QGraphicsSvgItem*>(ptr)->hoverEnterEvent(static_cast<QGraphicsSceneHoverEvent*>(event));
}

void QGraphicsSvgItem_HoverEnterEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::hoverEnterEvent(static_cast<QGraphicsSceneHoverEvent*>(event));
}

void QGraphicsSvgItem_HoverLeaveEvent(void* ptr, void* event)
{
	static_cast<QGraphicsSvgItem*>(ptr)->hoverLeaveEvent(static_cast<QGraphicsSceneHoverEvent*>(event));
}

void QGraphicsSvgItem_HoverLeaveEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::hoverLeaveEvent(static_cast<QGraphicsSceneHoverEvent*>(event));
}

void QGraphicsSvgItem_HoverMoveEvent(void* ptr, void* event)
{
	static_cast<QGraphicsSvgItem*>(ptr)->hoverMoveEvent(static_cast<QGraphicsSceneHoverEvent*>(event));
}

void QGraphicsSvgItem_HoverMoveEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::hoverMoveEvent(static_cast<QGraphicsSceneHoverEvent*>(event));
}

void QGraphicsSvgItem_InputMethodEvent(void* ptr, void* event)
{
	static_cast<QGraphicsSvgItem*>(ptr)->inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void QGraphicsSvgItem_InputMethodEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void* QGraphicsSvgItem_InputMethodQuery(void* ptr, long long query)
{
	return new QVariant(static_cast<QGraphicsSvgItem*>(ptr)->inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

void* QGraphicsSvgItem_InputMethodQueryDefault(void* ptr, long long query)
{
	return new QVariant(static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

char QGraphicsSvgItem_IsObscuredBy(void* ptr, void* item)
{
	return static_cast<QGraphicsSvgItem*>(ptr)->isObscuredBy(static_cast<QGraphicsItem*>(item));
}

char QGraphicsSvgItem_IsObscuredByDefault(void* ptr, void* item)
{
	return static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::isObscuredBy(static_cast<QGraphicsItem*>(item));
}

void* QGraphicsSvgItem_ItemChange(void* ptr, long long change, void* value)
{
	return new QVariant(static_cast<QGraphicsSvgItem*>(ptr)->itemChange(static_cast<QGraphicsItem::GraphicsItemChange>(change), *static_cast<QVariant*>(value)));
}

void* QGraphicsSvgItem_ItemChangeDefault(void* ptr, long long change, void* value)
{
	return new QVariant(static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::itemChange(static_cast<QGraphicsItem::GraphicsItemChange>(change), *static_cast<QVariant*>(value)));
}

void QGraphicsSvgItem_KeyPressEvent(void* ptr, void* event)
{
	static_cast<QGraphicsSvgItem*>(ptr)->keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QGraphicsSvgItem_KeyPressEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QGraphicsSvgItem_KeyReleaseEvent(void* ptr, void* event)
{
	static_cast<QGraphicsSvgItem*>(ptr)->keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QGraphicsSvgItem_KeyReleaseEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QGraphicsSvgItem_MouseDoubleClickEvent(void* ptr, void* event)
{
	static_cast<QGraphicsSvgItem*>(ptr)->mouseDoubleClickEvent(static_cast<QGraphicsSceneMouseEvent*>(event));
}

void QGraphicsSvgItem_MouseDoubleClickEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::mouseDoubleClickEvent(static_cast<QGraphicsSceneMouseEvent*>(event));
}

void QGraphicsSvgItem_MouseMoveEvent(void* ptr, void* event)
{
	static_cast<QGraphicsSvgItem*>(ptr)->mouseMoveEvent(static_cast<QGraphicsSceneMouseEvent*>(event));
}

void QGraphicsSvgItem_MouseMoveEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::mouseMoveEvent(static_cast<QGraphicsSceneMouseEvent*>(event));
}

void QGraphicsSvgItem_MousePressEvent(void* ptr, void* event)
{
	static_cast<QGraphicsSvgItem*>(ptr)->mousePressEvent(static_cast<QGraphicsSceneMouseEvent*>(event));
}

void QGraphicsSvgItem_MousePressEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::mousePressEvent(static_cast<QGraphicsSceneMouseEvent*>(event));
}

void QGraphicsSvgItem_MouseReleaseEvent(void* ptr, void* event)
{
	static_cast<QGraphicsSvgItem*>(ptr)->mouseReleaseEvent(static_cast<QGraphicsSceneMouseEvent*>(event));
}

void QGraphicsSvgItem_MouseReleaseEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::mouseReleaseEvent(static_cast<QGraphicsSceneMouseEvent*>(event));
}

void* QGraphicsSvgItem_OpaqueArea(void* ptr)
{
	return new QPainterPath(static_cast<QGraphicsSvgItem*>(ptr)->opaqueArea());
}

void* QGraphicsSvgItem_OpaqueAreaDefault(void* ptr)
{
	return new QPainterPath(static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::opaqueArea());
}

char QGraphicsSvgItem_SceneEvent(void* ptr, void* event)
{
	return static_cast<QGraphicsSvgItem*>(ptr)->sceneEvent(static_cast<QEvent*>(event));
}

char QGraphicsSvgItem_SceneEventDefault(void* ptr, void* event)
{
	return static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::sceneEvent(static_cast<QEvent*>(event));
}

char QGraphicsSvgItem_SceneEventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QGraphicsSvgItem*>(ptr)->sceneEventFilter(static_cast<QGraphicsItem*>(watched), static_cast<QEvent*>(event));
}

char QGraphicsSvgItem_SceneEventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::sceneEventFilter(static_cast<QGraphicsItem*>(watched), static_cast<QEvent*>(event));
}

void* QGraphicsSvgItem_Shape(void* ptr)
{
	return new QPainterPath(static_cast<QGraphicsSvgItem*>(ptr)->shape());
}

void* QGraphicsSvgItem_ShapeDefault(void* ptr)
{
	return new QPainterPath(static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::shape());
}

void QGraphicsSvgItem_WheelEvent(void* ptr, void* event)
{
	static_cast<QGraphicsSvgItem*>(ptr)->wheelEvent(static_cast<QGraphicsSceneWheelEvent*>(event));
}

void QGraphicsSvgItem_WheelEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::wheelEvent(static_cast<QGraphicsSceneWheelEvent*>(event));
}

struct QtSvg_PackedString QSvgGenerator_Description(void* ptr)
{
	return ({ QByteArray t85d23d = static_cast<QSvgGenerator*>(ptr)->description().toUtf8(); QtSvg_PackedString { const_cast<char*>(t85d23d.prepend("WHITESPACE").constData()+10), t85d23d.size()-10 }; });
}

struct QtSvg_PackedString QSvgGenerator_FileName(void* ptr)
{
	return ({ QByteArray ta251c4 = static_cast<QSvgGenerator*>(ptr)->fileName().toUtf8(); QtSvg_PackedString { const_cast<char*>(ta251c4.prepend("WHITESPACE").constData()+10), ta251c4.size()-10 }; });
}

void* QSvgGenerator_OutputDevice(void* ptr)
{
	return static_cast<QSvgGenerator*>(ptr)->outputDevice();
}

int QSvgGenerator_Resolution(void* ptr)
{
	return static_cast<QSvgGenerator*>(ptr)->resolution();
}

void QSvgGenerator_SetDescription(void* ptr, char* description)
{
	static_cast<QSvgGenerator*>(ptr)->setDescription(QString(description));
}

void QSvgGenerator_SetFileName(void* ptr, char* fileName)
{
	static_cast<QSvgGenerator*>(ptr)->setFileName(QString(fileName));
}

void QSvgGenerator_SetOutputDevice(void* ptr, void* outputDevice)
{
	static_cast<QSvgGenerator*>(ptr)->setOutputDevice(static_cast<QIODevice*>(outputDevice));
}

void QSvgGenerator_SetResolution(void* ptr, int dpi)
{
	static_cast<QSvgGenerator*>(ptr)->setResolution(dpi);
}

void QSvgGenerator_SetSize(void* ptr, void* size)
{
	static_cast<QSvgGenerator*>(ptr)->setSize(*static_cast<QSize*>(size));
}

void QSvgGenerator_SetTitle(void* ptr, char* title)
{
	static_cast<QSvgGenerator*>(ptr)->setTitle(QString(title));
}

void QSvgGenerator_SetViewBox(void* ptr, void* viewBox)
{
	static_cast<QSvgGenerator*>(ptr)->setViewBox(*static_cast<QRect*>(viewBox));
}

void QSvgGenerator_SetViewBox2(void* ptr, void* viewBox)
{
	static_cast<QSvgGenerator*>(ptr)->setViewBox(*static_cast<QRectF*>(viewBox));
}

void* QSvgGenerator_Size(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QSvgGenerator*>(ptr)->size(); new QSize(tmpValue.width(), tmpValue.height()); });
}

struct QtSvg_PackedString QSvgGenerator_Title(void* ptr)
{
	return ({ QByteArray tb87794 = static_cast<QSvgGenerator*>(ptr)->title().toUtf8(); QtSvg_PackedString { const_cast<char*>(tb87794.prepend("WHITESPACE").constData()+10), tb87794.size()-10 }; });
}

void* QSvgGenerator_ViewBoxF(void* ptr)
{
	return ({ QRectF tmpValue = static_cast<QSvgGenerator*>(ptr)->viewBoxF(); new QRectF(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void* QSvgGenerator_NewQSvgGenerator()
{
	return new QSvgGenerator();
}

int QSvgGenerator_Metric(void* ptr, long long metric)
{
	return static_cast<QSvgGenerator*>(ptr)->metric(static_cast<QPaintDevice::PaintDeviceMetric>(metric));
}

void* QSvgGenerator_PaintEngine(void* ptr)
{
	return static_cast<QSvgGenerator*>(ptr)->paintEngine();
}

void* QSvgGenerator_ViewBox(void* ptr)
{
	return ({ QRect tmpValue = static_cast<QSvgGenerator*>(ptr)->viewBox(); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void QSvgGenerator_DestroyQSvgGenerator(void* ptr)
{
	static_cast<QSvgGenerator*>(ptr)->~QSvgGenerator();
}

class MyQSvgRenderer: public QSvgRenderer
{
public:
	MyQSvgRenderer(QObject *parent) : QSvgRenderer(parent) {};
	MyQSvgRenderer(QXmlStreamReader *contents, QObject *parent) : QSvgRenderer(contents, parent) {};
	MyQSvgRenderer(const QByteArray &contents, QObject *parent) : QSvgRenderer(contents, parent) {};
	MyQSvgRenderer(const QString &filename, QObject *parent) : QSvgRenderer(filename, parent) {};
	bool load(QXmlStreamReader * contents) { return callbackQSvgRenderer_Load3(this, contents) != 0; };
	bool load(const QByteArray & contents) { return callbackQSvgRenderer_Load2(this, const_cast<QByteArray*>(&contents)) != 0; };
	bool load(const QString & filename) { QByteArray t08deae = filename.toUtf8(); QtSvg_PackedString filenamePacked = { const_cast<char*>(t08deae.prepend("WHITESPACE").constData()+10), t08deae.size()-10 };return callbackQSvgRenderer_Load(this, filenamePacked) != 0; };
	void render(QPainter * painter) { callbackQSvgRenderer_Render(this, painter); };
	void render(QPainter * painter, const QRectF & bounds) { callbackQSvgRenderer_Render2(this, painter, const_cast<QRectF*>(&bounds)); };
	void render(QPainter * painter, const QString & elementId, const QRectF & bounds) { QByteArray t5b5420 = elementId.toUtf8(); QtSvg_PackedString elementIdPacked = { const_cast<char*>(t5b5420.prepend("WHITESPACE").constData()+10), t5b5420.size()-10 };callbackQSvgRenderer_Render3(this, painter, elementIdPacked, const_cast<QRectF*>(&bounds)); };
	void Signal_RepaintNeeded() { callbackQSvgRenderer_RepaintNeeded(this); };
	void timerEvent(QTimerEvent * event) { callbackQSvgRenderer_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQSvgRenderer_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSvgRenderer_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQSvgRenderer_CustomEvent(this, event); };
	void deleteLater() { callbackQSvgRenderer_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSvgRenderer_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQSvgRenderer_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSvgRenderer_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSvgRenderer_MetaObject(const_cast<MyQSvgRenderer*>(this))); };
};

int QSvgRenderer_FramesPerSecond(void* ptr)
{
	return static_cast<QSvgRenderer*>(ptr)->framesPerSecond();
}

void QSvgRenderer_SetFramesPerSecond(void* ptr, int num)
{
	static_cast<QSvgRenderer*>(ptr)->setFramesPerSecond(num);
}

void QSvgRenderer_SetViewBox(void* ptr, void* viewbox)
{
	static_cast<QSvgRenderer*>(ptr)->setViewBox(*static_cast<QRect*>(viewbox));
}

void QSvgRenderer_SetViewBox2(void* ptr, void* viewbox)
{
	static_cast<QSvgRenderer*>(ptr)->setViewBox(*static_cast<QRectF*>(viewbox));
}

void* QSvgRenderer_ViewBoxF(void* ptr)
{
	return ({ QRectF tmpValue = static_cast<QSvgRenderer*>(ptr)->viewBoxF(); new QRectF(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void* QSvgRenderer_NewQSvgRenderer(void* parent)
{
	return new MyQSvgRenderer(static_cast<QObject*>(parent));
}

void* QSvgRenderer_NewQSvgRenderer4(void* contents, void* parent)
{
	return new MyQSvgRenderer(static_cast<QXmlStreamReader*>(contents), static_cast<QObject*>(parent));
}

void* QSvgRenderer_NewQSvgRenderer3(void* contents, void* parent)
{
	return new MyQSvgRenderer(*static_cast<QByteArray*>(contents), static_cast<QObject*>(parent));
}

void* QSvgRenderer_NewQSvgRenderer2(char* filename, void* parent)
{
	return new MyQSvgRenderer(QString(filename), static_cast<QObject*>(parent));
}

char QSvgRenderer_Animated(void* ptr)
{
	return static_cast<QSvgRenderer*>(ptr)->animated();
}

void* QSvgRenderer_BoundsOnElement(void* ptr, char* id)
{
	return ({ QRectF tmpValue = static_cast<QSvgRenderer*>(ptr)->boundsOnElement(QString(id)); new QRectF(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void* QSvgRenderer_DefaultSize(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QSvgRenderer*>(ptr)->defaultSize(); new QSize(tmpValue.width(), tmpValue.height()); });
}

char QSvgRenderer_ElementExists(void* ptr, char* id)
{
	return static_cast<QSvgRenderer*>(ptr)->elementExists(QString(id));
}

char QSvgRenderer_IsValid(void* ptr)
{
	return static_cast<QSvgRenderer*>(ptr)->isValid();
}

char QSvgRenderer_Load3(void* ptr, void* contents)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QSvgRenderer*>(ptr), "load", Q_RETURN_ARG(bool, returnArg), Q_ARG(QXmlStreamReader*, static_cast<QXmlStreamReader*>(contents)));
	return returnArg;
}

char QSvgRenderer_Load2(void* ptr, void* contents)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QSvgRenderer*>(ptr), "load", Q_RETURN_ARG(bool, returnArg), Q_ARG(QByteArray, *static_cast<QByteArray*>(contents)));
	return returnArg;
}

char QSvgRenderer_Load(void* ptr, char* filename)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QSvgRenderer*>(ptr), "load", Q_RETURN_ARG(bool, returnArg), Q_ARG(QString, QString(filename)));
	return returnArg;
}

void QSvgRenderer_Render(void* ptr, void* painter)
{
	QMetaObject::invokeMethod(static_cast<QSvgRenderer*>(ptr), "render", Q_ARG(QPainter*, static_cast<QPainter*>(painter)));
}

void QSvgRenderer_Render2(void* ptr, void* painter, void* bounds)
{
	QMetaObject::invokeMethod(static_cast<QSvgRenderer*>(ptr), "render", Q_ARG(QPainter*, static_cast<QPainter*>(painter)), Q_ARG(QRectF, *static_cast<QRectF*>(bounds)));
}

void QSvgRenderer_Render3(void* ptr, void* painter, char* elementId, void* bounds)
{
	QMetaObject::invokeMethod(static_cast<QSvgRenderer*>(ptr), "render", Q_ARG(QPainter*, static_cast<QPainter*>(painter)), Q_ARG(QString, QString(elementId)), Q_ARG(QRectF, *static_cast<QRectF*>(bounds)));
}

void QSvgRenderer_ConnectRepaintNeeded(void* ptr)
{
	QObject::connect(static_cast<QSvgRenderer*>(ptr), static_cast<void (QSvgRenderer::*)()>(&QSvgRenderer::repaintNeeded), static_cast<MyQSvgRenderer*>(ptr), static_cast<void (MyQSvgRenderer::*)()>(&MyQSvgRenderer::Signal_RepaintNeeded));
}

void QSvgRenderer_DisconnectRepaintNeeded(void* ptr)
{
	QObject::disconnect(static_cast<QSvgRenderer*>(ptr), static_cast<void (QSvgRenderer::*)()>(&QSvgRenderer::repaintNeeded), static_cast<MyQSvgRenderer*>(ptr), static_cast<void (MyQSvgRenderer::*)()>(&MyQSvgRenderer::Signal_RepaintNeeded));
}

void QSvgRenderer_RepaintNeeded(void* ptr)
{
	static_cast<QSvgRenderer*>(ptr)->repaintNeeded();
}

void* QSvgRenderer_ViewBox(void* ptr)
{
	return ({ QRect tmpValue = static_cast<QSvgRenderer*>(ptr)->viewBox(); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void QSvgRenderer_DestroyQSvgRenderer(void* ptr)
{
	static_cast<QSvgRenderer*>(ptr)->~QSvgRenderer();
}

void QSvgRenderer_TimerEvent(void* ptr, void* event)
{
	static_cast<QSvgRenderer*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QSvgRenderer_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QSvgRenderer*>(ptr)->QSvgRenderer::timerEvent(static_cast<QTimerEvent*>(event));
}

void QSvgRenderer_ChildEvent(void* ptr, void* event)
{
	static_cast<QSvgRenderer*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QSvgRenderer_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QSvgRenderer*>(ptr)->QSvgRenderer::childEvent(static_cast<QChildEvent*>(event));
}

void QSvgRenderer_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QSvgRenderer*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSvgRenderer_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QSvgRenderer*>(ptr)->QSvgRenderer::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSvgRenderer_CustomEvent(void* ptr, void* event)
{
	static_cast<QSvgRenderer*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QSvgRenderer_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QSvgRenderer*>(ptr)->QSvgRenderer::customEvent(static_cast<QEvent*>(event));
}

void QSvgRenderer_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QSvgRenderer*>(ptr), "deleteLater");
}

void QSvgRenderer_DeleteLaterDefault(void* ptr)
{
	static_cast<QSvgRenderer*>(ptr)->QSvgRenderer::deleteLater();
}

void QSvgRenderer_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QSvgRenderer*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSvgRenderer_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QSvgRenderer*>(ptr)->QSvgRenderer::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QSvgRenderer_Event(void* ptr, void* e)
{
	return static_cast<QSvgRenderer*>(ptr)->event(static_cast<QEvent*>(e));
}

char QSvgRenderer_EventDefault(void* ptr, void* e)
{
	return static_cast<QSvgRenderer*>(ptr)->QSvgRenderer::event(static_cast<QEvent*>(e));
}

char QSvgRenderer_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QSvgRenderer*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QSvgRenderer_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QSvgRenderer*>(ptr)->QSvgRenderer::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QSvgRenderer_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QSvgRenderer*>(ptr)->metaObject());
}

void* QSvgRenderer_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QSvgRenderer*>(ptr)->QSvgRenderer::metaObject());
}

class MyQSvgWidget: public QSvgWidget
{
public:
	MyQSvgWidget(QWidget *parent) : QSvgWidget(parent) {};
	MyQSvgWidget(const QString &file, QWidget *parent) : QSvgWidget(file, parent) {};
	void load(const QByteArray & contents) { callbackQSvgWidget_Load2(this, const_cast<QByteArray*>(&contents)); };
	void load(const QString & file) { QByteArray t971c41 = file.toUtf8(); QtSvg_PackedString filePacked = { const_cast<char*>(t971c41.prepend("WHITESPACE").constData()+10), t971c41.size()-10 };callbackQSvgWidget_Load(this, filePacked); };
	void actionEvent(QActionEvent * event) { callbackQSvgWidget_ActionEvent(this, event); };
	void dragEnterEvent(QDragEnterEvent * event) { callbackQSvgWidget_DragEnterEvent(this, event); };
	void dragLeaveEvent(QDragLeaveEvent * event) { callbackQSvgWidget_DragLeaveEvent(this, event); };
	void dragMoveEvent(QDragMoveEvent * event) { callbackQSvgWidget_DragMoveEvent(this, event); };
	void dropEvent(QDropEvent * event) { callbackQSvgWidget_DropEvent(this, event); };
	void enterEvent(QEvent * event) { callbackQSvgWidget_EnterEvent(this, event); };
	void focusInEvent(QFocusEvent * event) { callbackQSvgWidget_FocusInEvent(this, event); };
	void focusOutEvent(QFocusEvent * event) { callbackQSvgWidget_FocusOutEvent(this, event); };
	void hideEvent(QHideEvent * event) { callbackQSvgWidget_HideEvent(this, event); };
	void leaveEvent(QEvent * event) { callbackQSvgWidget_LeaveEvent(this, event); };
	QSize minimumSizeHint() const { return *static_cast<QSize*>(callbackQSvgWidget_MinimumSizeHint(const_cast<MyQSvgWidget*>(this))); };
	void moveEvent(QMoveEvent * event) { callbackQSvgWidget_MoveEvent(this, event); };
	void setEnabled(bool vbo) { callbackQSvgWidget_SetEnabled(this, vbo); };
	void setStyleSheet(const QString & styleSheet) { QByteArray t728ae7 = styleSheet.toUtf8(); QtSvg_PackedString styleSheetPacked = { const_cast<char*>(t728ae7.prepend("WHITESPACE").constData()+10), t728ae7.size()-10 };callbackQSvgWidget_SetStyleSheet(this, styleSheetPacked); };
	void setVisible(bool visible) { callbackQSvgWidget_SetVisible(this, visible); };
	void setWindowModified(bool vbo) { callbackQSvgWidget_SetWindowModified(this, vbo); };
	void setWindowTitle(const QString & vqs) { QByteArray tda39a3 = vqs.toUtf8(); QtSvg_PackedString vqsPacked = { const_cast<char*>(tda39a3.prepend("WHITESPACE").constData()+10), tda39a3.size()-10 };callbackQSvgWidget_SetWindowTitle(this, vqsPacked); };
	void showEvent(QShowEvent * event) { callbackQSvgWidget_ShowEvent(this, event); };
	void changeEvent(QEvent * event) { callbackQSvgWidget_ChangeEvent(this, event); };
	bool close() { return callbackQSvgWidget_Close(this) != 0; };
	void closeEvent(QCloseEvent * event) { callbackQSvgWidget_CloseEvent(this, event); };
	void contextMenuEvent(QContextMenuEvent * event) { callbackQSvgWidget_ContextMenuEvent(this, event); };
	bool focusNextPrevChild(bool next) { return callbackQSvgWidget_FocusNextPrevChild(this, next) != 0; };
	bool hasHeightForWidth() const { return callbackQSvgWidget_HasHeightForWidth(const_cast<MyQSvgWidget*>(this)) != 0; };
	int heightForWidth(int w) const { return callbackQSvgWidget_HeightForWidth(const_cast<MyQSvgWidget*>(this), w); };
	void hide() { callbackQSvgWidget_Hide(this); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackQSvgWidget_InputMethodEvent(this, event); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackQSvgWidget_InputMethodQuery(const_cast<MyQSvgWidget*>(this), query)); };
	void keyPressEvent(QKeyEvent * event) { callbackQSvgWidget_KeyPressEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQSvgWidget_KeyReleaseEvent(this, event); };
	void lower() { callbackQSvgWidget_Lower(this); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQSvgWidget_MouseDoubleClickEvent(this, event); };
	void mouseMoveEvent(QMouseEvent * event) { callbackQSvgWidget_MouseMoveEvent(this, event); };
	void mousePressEvent(QMouseEvent * event) { callbackQSvgWidget_MousePressEvent(this, event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackQSvgWidget_MouseReleaseEvent(this, event); };
	bool nativeEvent(const QByteArray & eventType, void * message, long * result) { return callbackQSvgWidget_NativeEvent(this, const_cast<QByteArray*>(&eventType), message, *result) != 0; };
	void raise() { callbackQSvgWidget_Raise(this); };
	void repaint() { callbackQSvgWidget_Repaint(this); };
	void resizeEvent(QResizeEvent * event) { callbackQSvgWidget_ResizeEvent(this, event); };
	void setDisabled(bool disable) { callbackQSvgWidget_SetDisabled(this, disable); };
	void setFocus() { callbackQSvgWidget_SetFocus2(this); };
	void setHidden(bool hidden) { callbackQSvgWidget_SetHidden(this, hidden); };
	void show() { callbackQSvgWidget_Show(this); };
	void showFullScreen() { callbackQSvgWidget_ShowFullScreen(this); };
	void showMaximized() { callbackQSvgWidget_ShowMaximized(this); };
	void showMinimized() { callbackQSvgWidget_ShowMinimized(this); };
	void showNormal() { callbackQSvgWidget_ShowNormal(this); };
	void tabletEvent(QTabletEvent * event) { callbackQSvgWidget_TabletEvent(this, event); };
	void update() { callbackQSvgWidget_Update(this); };
	void updateMicroFocus() { callbackQSvgWidget_UpdateMicroFocus(this); };
	void wheelEvent(QWheelEvent * event) { callbackQSvgWidget_WheelEvent(this, event); };
	void timerEvent(QTimerEvent * event) { callbackQSvgWidget_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQSvgWidget_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSvgWidget_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQSvgWidget_CustomEvent(this, event); };
	void deleteLater() { callbackQSvgWidget_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSvgWidget_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSvgWidget_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSvgWidget_MetaObject(const_cast<MyQSvgWidget*>(this))); };
};

void* QSvgWidget_NewQSvgWidget(void* parent)
{
	return new MyQSvgWidget(static_cast<QWidget*>(parent));
}

void* QSvgWidget_NewQSvgWidget2(char* file, void* parent)
{
	return new MyQSvgWidget(QString(file), static_cast<QWidget*>(parent));
}

void QSvgWidget_Load2(void* ptr, void* contents)
{
	QMetaObject::invokeMethod(static_cast<QSvgWidget*>(ptr), "load", Q_ARG(QByteArray, *static_cast<QByteArray*>(contents)));
}

void QSvgWidget_Load(void* ptr, char* file)
{
	QMetaObject::invokeMethod(static_cast<QSvgWidget*>(ptr), "load", Q_ARG(QString, QString(file)));
}

void QSvgWidget_PaintEvent(void* ptr, void* event)
{
	static_cast<QSvgWidget*>(ptr)->paintEvent(static_cast<QPaintEvent*>(event));
}

void* QSvgWidget_Renderer(void* ptr)
{
	return static_cast<QSvgWidget*>(ptr)->renderer();
}

void* QSvgWidget_SizeHint(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QSvgWidget*>(ptr)->sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void QSvgWidget_DestroyQSvgWidget(void* ptr)
{
	static_cast<QSvgWidget*>(ptr)->~QSvgWidget();
}

void QSvgWidget_ActionEvent(void* ptr, void* event)
{
	static_cast<QSvgWidget*>(ptr)->actionEvent(static_cast<QActionEvent*>(event));
}

void QSvgWidget_ActionEventDefault(void* ptr, void* event)
{
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::actionEvent(static_cast<QActionEvent*>(event));
}

void QSvgWidget_DragEnterEvent(void* ptr, void* event)
{
	static_cast<QSvgWidget*>(ptr)->dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QSvgWidget_DragEnterEventDefault(void* ptr, void* event)
{
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QSvgWidget_DragLeaveEvent(void* ptr, void* event)
{
	static_cast<QSvgWidget*>(ptr)->dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QSvgWidget_DragLeaveEventDefault(void* ptr, void* event)
{
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QSvgWidget_DragMoveEvent(void* ptr, void* event)
{
	static_cast<QSvgWidget*>(ptr)->dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QSvgWidget_DragMoveEventDefault(void* ptr, void* event)
{
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QSvgWidget_DropEvent(void* ptr, void* event)
{
	static_cast<QSvgWidget*>(ptr)->dropEvent(static_cast<QDropEvent*>(event));
}

void QSvgWidget_DropEventDefault(void* ptr, void* event)
{
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::dropEvent(static_cast<QDropEvent*>(event));
}

void QSvgWidget_EnterEvent(void* ptr, void* event)
{
	static_cast<QSvgWidget*>(ptr)->enterEvent(static_cast<QEvent*>(event));
}

void QSvgWidget_EnterEventDefault(void* ptr, void* event)
{
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::enterEvent(static_cast<QEvent*>(event));
}

void QSvgWidget_FocusInEvent(void* ptr, void* event)
{
	static_cast<QSvgWidget*>(ptr)->focusInEvent(static_cast<QFocusEvent*>(event));
}

void QSvgWidget_FocusInEventDefault(void* ptr, void* event)
{
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::focusInEvent(static_cast<QFocusEvent*>(event));
}

void QSvgWidget_FocusOutEvent(void* ptr, void* event)
{
	static_cast<QSvgWidget*>(ptr)->focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QSvgWidget_FocusOutEventDefault(void* ptr, void* event)
{
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QSvgWidget_HideEvent(void* ptr, void* event)
{
	static_cast<QSvgWidget*>(ptr)->hideEvent(static_cast<QHideEvent*>(event));
}

void QSvgWidget_HideEventDefault(void* ptr, void* event)
{
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::hideEvent(static_cast<QHideEvent*>(event));
}

void QSvgWidget_LeaveEvent(void* ptr, void* event)
{
	static_cast<QSvgWidget*>(ptr)->leaveEvent(static_cast<QEvent*>(event));
}

void QSvgWidget_LeaveEventDefault(void* ptr, void* event)
{
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::leaveEvent(static_cast<QEvent*>(event));
}

void* QSvgWidget_MinimumSizeHint(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QSvgWidget*>(ptr)->minimumSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QSvgWidget_MinimumSizeHintDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QSvgWidget*>(ptr)->QSvgWidget::minimumSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void QSvgWidget_MoveEvent(void* ptr, void* event)
{
	static_cast<QSvgWidget*>(ptr)->moveEvent(static_cast<QMoveEvent*>(event));
}

void QSvgWidget_MoveEventDefault(void* ptr, void* event)
{
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::moveEvent(static_cast<QMoveEvent*>(event));
}

void QSvgWidget_SetEnabled(void* ptr, char vbo)
{
	QMetaObject::invokeMethod(static_cast<QSvgWidget*>(ptr), "setEnabled", Q_ARG(bool, vbo != 0));
}

void QSvgWidget_SetEnabledDefault(void* ptr, char vbo)
{
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::setEnabled(vbo != 0);
}

void QSvgWidget_SetStyleSheet(void* ptr, char* styleSheet)
{
	QMetaObject::invokeMethod(static_cast<QSvgWidget*>(ptr), "setStyleSheet", Q_ARG(QString, QString(styleSheet)));
}

void QSvgWidget_SetStyleSheetDefault(void* ptr, char* styleSheet)
{
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::setStyleSheet(QString(styleSheet));
}

void QSvgWidget_SetVisible(void* ptr, char visible)
{
	QMetaObject::invokeMethod(static_cast<QSvgWidget*>(ptr), "setVisible", Q_ARG(bool, visible != 0));
}

void QSvgWidget_SetVisibleDefault(void* ptr, char visible)
{
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::setVisible(visible != 0);
}

void QSvgWidget_SetWindowModified(void* ptr, char vbo)
{
	QMetaObject::invokeMethod(static_cast<QSvgWidget*>(ptr), "setWindowModified", Q_ARG(bool, vbo != 0));
}

void QSvgWidget_SetWindowModifiedDefault(void* ptr, char vbo)
{
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::setWindowModified(vbo != 0);
}

void QSvgWidget_SetWindowTitle(void* ptr, char* vqs)
{
	QMetaObject::invokeMethod(static_cast<QSvgWidget*>(ptr), "setWindowTitle", Q_ARG(QString, QString(vqs)));
}

void QSvgWidget_SetWindowTitleDefault(void* ptr, char* vqs)
{
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::setWindowTitle(QString(vqs));
}

void QSvgWidget_ShowEvent(void* ptr, void* event)
{
	static_cast<QSvgWidget*>(ptr)->showEvent(static_cast<QShowEvent*>(event));
}

void QSvgWidget_ShowEventDefault(void* ptr, void* event)
{
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::showEvent(static_cast<QShowEvent*>(event));
}

void QSvgWidget_ChangeEvent(void* ptr, void* event)
{
	static_cast<QSvgWidget*>(ptr)->changeEvent(static_cast<QEvent*>(event));
}

void QSvgWidget_ChangeEventDefault(void* ptr, void* event)
{
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::changeEvent(static_cast<QEvent*>(event));
}

char QSvgWidget_Close(void* ptr)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QSvgWidget*>(ptr), "close", Q_RETURN_ARG(bool, returnArg));
	return returnArg;
}

char QSvgWidget_CloseDefault(void* ptr)
{
	return static_cast<QSvgWidget*>(ptr)->QSvgWidget::close();
}

void QSvgWidget_CloseEvent(void* ptr, void* event)
{
	static_cast<QSvgWidget*>(ptr)->closeEvent(static_cast<QCloseEvent*>(event));
}

void QSvgWidget_CloseEventDefault(void* ptr, void* event)
{
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::closeEvent(static_cast<QCloseEvent*>(event));
}

void QSvgWidget_ContextMenuEvent(void* ptr, void* event)
{
	static_cast<QSvgWidget*>(ptr)->contextMenuEvent(static_cast<QContextMenuEvent*>(event));
}

void QSvgWidget_ContextMenuEventDefault(void* ptr, void* event)
{
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
}

char QSvgWidget_FocusNextPrevChild(void* ptr, char next)
{
	return static_cast<QSvgWidget*>(ptr)->focusNextPrevChild(next != 0);
}

char QSvgWidget_FocusNextPrevChildDefault(void* ptr, char next)
{
	return static_cast<QSvgWidget*>(ptr)->QSvgWidget::focusNextPrevChild(next != 0);
}

char QSvgWidget_HasHeightForWidth(void* ptr)
{
	return static_cast<QSvgWidget*>(ptr)->hasHeightForWidth();
}

char QSvgWidget_HasHeightForWidthDefault(void* ptr)
{
	return static_cast<QSvgWidget*>(ptr)->QSvgWidget::hasHeightForWidth();
}

int QSvgWidget_HeightForWidth(void* ptr, int w)
{
	return static_cast<QSvgWidget*>(ptr)->heightForWidth(w);
}

int QSvgWidget_HeightForWidthDefault(void* ptr, int w)
{
	return static_cast<QSvgWidget*>(ptr)->QSvgWidget::heightForWidth(w);
}

void QSvgWidget_Hide(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QSvgWidget*>(ptr), "hide");
}

void QSvgWidget_HideDefault(void* ptr)
{
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::hide();
}

void QSvgWidget_InputMethodEvent(void* ptr, void* event)
{
	static_cast<QSvgWidget*>(ptr)->inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void QSvgWidget_InputMethodEventDefault(void* ptr, void* event)
{
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void* QSvgWidget_InputMethodQuery(void* ptr, long long query)
{
	return new QVariant(static_cast<QSvgWidget*>(ptr)->inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

void* QSvgWidget_InputMethodQueryDefault(void* ptr, long long query)
{
	return new QVariant(static_cast<QSvgWidget*>(ptr)->QSvgWidget::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

void QSvgWidget_KeyPressEvent(void* ptr, void* event)
{
	static_cast<QSvgWidget*>(ptr)->keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QSvgWidget_KeyPressEventDefault(void* ptr, void* event)
{
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QSvgWidget_KeyReleaseEvent(void* ptr, void* event)
{
	static_cast<QSvgWidget*>(ptr)->keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QSvgWidget_KeyReleaseEventDefault(void* ptr, void* event)
{
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QSvgWidget_Lower(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QSvgWidget*>(ptr), "lower");
}

void QSvgWidget_LowerDefault(void* ptr)
{
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::lower();
}

void QSvgWidget_MouseDoubleClickEvent(void* ptr, void* event)
{
	static_cast<QSvgWidget*>(ptr)->mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QSvgWidget_MouseDoubleClickEventDefault(void* ptr, void* event)
{
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QSvgWidget_MouseMoveEvent(void* ptr, void* event)
{
	static_cast<QSvgWidget*>(ptr)->mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QSvgWidget_MouseMoveEventDefault(void* ptr, void* event)
{
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QSvgWidget_MousePressEvent(void* ptr, void* event)
{
	static_cast<QSvgWidget*>(ptr)->mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QSvgWidget_MousePressEventDefault(void* ptr, void* event)
{
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QSvgWidget_MouseReleaseEvent(void* ptr, void* event)
{
	static_cast<QSvgWidget*>(ptr)->mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void QSvgWidget_MouseReleaseEventDefault(void* ptr, void* event)
{
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

char QSvgWidget_NativeEvent(void* ptr, void* eventType, void* message, long result)
{
	return static_cast<QSvgWidget*>(ptr)->nativeEvent(*static_cast<QByteArray*>(eventType), message, &result);
}

char QSvgWidget_NativeEventDefault(void* ptr, void* eventType, void* message, long result)
{
	return static_cast<QSvgWidget*>(ptr)->QSvgWidget::nativeEvent(*static_cast<QByteArray*>(eventType), message, &result);
}

void QSvgWidget_Raise(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QSvgWidget*>(ptr), "raise");
}

void QSvgWidget_RaiseDefault(void* ptr)
{
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::raise();
}

void QSvgWidget_Repaint(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QSvgWidget*>(ptr), "repaint");
}

void QSvgWidget_RepaintDefault(void* ptr)
{
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::repaint();
}

void QSvgWidget_ResizeEvent(void* ptr, void* event)
{
	static_cast<QSvgWidget*>(ptr)->resizeEvent(static_cast<QResizeEvent*>(event));
}

void QSvgWidget_ResizeEventDefault(void* ptr, void* event)
{
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::resizeEvent(static_cast<QResizeEvent*>(event));
}

void QSvgWidget_SetDisabled(void* ptr, char disable)
{
	QMetaObject::invokeMethod(static_cast<QSvgWidget*>(ptr), "setDisabled", Q_ARG(bool, disable != 0));
}

void QSvgWidget_SetDisabledDefault(void* ptr, char disable)
{
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::setDisabled(disable != 0);
}

void QSvgWidget_SetFocus2(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QSvgWidget*>(ptr), "setFocus");
}

void QSvgWidget_SetFocus2Default(void* ptr)
{
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::setFocus();
}

void QSvgWidget_SetHidden(void* ptr, char hidden)
{
	QMetaObject::invokeMethod(static_cast<QSvgWidget*>(ptr), "setHidden", Q_ARG(bool, hidden != 0));
}

void QSvgWidget_SetHiddenDefault(void* ptr, char hidden)
{
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::setHidden(hidden != 0);
}

void QSvgWidget_Show(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QSvgWidget*>(ptr), "show");
}

void QSvgWidget_ShowDefault(void* ptr)
{
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::show();
}

void QSvgWidget_ShowFullScreen(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QSvgWidget*>(ptr), "showFullScreen");
}

void QSvgWidget_ShowFullScreenDefault(void* ptr)
{
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::showFullScreen();
}

void QSvgWidget_ShowMaximized(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QSvgWidget*>(ptr), "showMaximized");
}

void QSvgWidget_ShowMaximizedDefault(void* ptr)
{
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::showMaximized();
}

void QSvgWidget_ShowMinimized(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QSvgWidget*>(ptr), "showMinimized");
}

void QSvgWidget_ShowMinimizedDefault(void* ptr)
{
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::showMinimized();
}

void QSvgWidget_ShowNormal(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QSvgWidget*>(ptr), "showNormal");
}

void QSvgWidget_ShowNormalDefault(void* ptr)
{
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::showNormal();
}

void QSvgWidget_TabletEvent(void* ptr, void* event)
{
	static_cast<QSvgWidget*>(ptr)->tabletEvent(static_cast<QTabletEvent*>(event));
}

void QSvgWidget_TabletEventDefault(void* ptr, void* event)
{
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::tabletEvent(static_cast<QTabletEvent*>(event));
}

void QSvgWidget_Update(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QSvgWidget*>(ptr), "update");
}

void QSvgWidget_UpdateDefault(void* ptr)
{
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::update();
}

void QSvgWidget_UpdateMicroFocus(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QSvgWidget*>(ptr), "updateMicroFocus");
}

void QSvgWidget_UpdateMicroFocusDefault(void* ptr)
{
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::updateMicroFocus();
}

void QSvgWidget_WheelEvent(void* ptr, void* event)
{
	static_cast<QSvgWidget*>(ptr)->wheelEvent(static_cast<QWheelEvent*>(event));
}

void QSvgWidget_WheelEventDefault(void* ptr, void* event)
{
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::wheelEvent(static_cast<QWheelEvent*>(event));
}

void QSvgWidget_TimerEvent(void* ptr, void* event)
{
	static_cast<QSvgWidget*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QSvgWidget_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::timerEvent(static_cast<QTimerEvent*>(event));
}

void QSvgWidget_ChildEvent(void* ptr, void* event)
{
	static_cast<QSvgWidget*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QSvgWidget_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::childEvent(static_cast<QChildEvent*>(event));
}

void QSvgWidget_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QSvgWidget*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSvgWidget_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSvgWidget_CustomEvent(void* ptr, void* event)
{
	static_cast<QSvgWidget*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QSvgWidget_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::customEvent(static_cast<QEvent*>(event));
}

void QSvgWidget_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QSvgWidget*>(ptr), "deleteLater");
}

void QSvgWidget_DeleteLaterDefault(void* ptr)
{
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::deleteLater();
}

void QSvgWidget_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QSvgWidget*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSvgWidget_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QSvgWidget*>(ptr)->QSvgWidget::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QSvgWidget_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QSvgWidget*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QSvgWidget_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QSvgWidget*>(ptr)->QSvgWidget::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QSvgWidget_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QSvgWidget*>(ptr)->metaObject());
}

void* QSvgWidget_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QSvgWidget*>(ptr)->QSvgWidget::metaObject());
}

