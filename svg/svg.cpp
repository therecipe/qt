// +build !minimal

#define protected public
#define private public

#include "svg.h"
#include "_cgo_export.h"

#include <QAction>
#include <QActionEvent>
#include <QByteArray>
#include <QCamera>
#include <QCameraImageCapture>
#include <QChildEvent>
#include <QCloseEvent>
#include <QContextMenuEvent>
#include <QDBusPendingCall>
#include <QDBusPendingCallWatcher>
#include <QDrag>
#include <QDragEnterEvent>
#include <QDragLeaveEvent>
#include <QDragMoveEvent>
#include <QDropEvent>
#include <QEvent>
#include <QExtensionFactory>
#include <QExtensionManager>
#include <QFocusEvent>
#include <QGraphicsItem>
#include <QGraphicsObject>
#include <QGraphicsScene>
#include <QGraphicsSceneContextMenuEvent>
#include <QGraphicsSceneDragDropEvent>
#include <QGraphicsSceneHoverEvent>
#include <QGraphicsSceneMouseEvent>
#include <QGraphicsSceneWheelEvent>
#include <QGraphicsSvgItem>
#include <QGraphicsTransform>
#include <QGraphicsWidget>
#include <QHideEvent>
#include <QIODevice>
#include <QIcon>
#include <QInputMethod>
#include <QInputMethodEvent>
#include <QKeyEvent>
#include <QLayout>
#include <QList>
#include <QMediaPlaylist>
#include <QMediaRecorder>
#include <QMetaMethod>
#include <QMetaObject>
#include <QMouseEvent>
#include <QMoveEvent>
#include <QObject>
#include <QOffscreenSurface>
#include <QPaintDevice>
#include <QPaintDeviceWindow>
#include <QPaintEngine>
#include <QPaintEvent>
#include <QPainter>
#include <QPainterPath>
#include <QPdfWriter>
#include <QPoint>
#include <QPointF>
#include <QQuickItem>
#include <QRadioData>
#include <QRect>
#include <QRectF>
#include <QResizeEvent>
#include <QShowEvent>
#include <QSignalSpy>
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
#include <QWindow>
#include <QXmlStreamReader>

class MyQGraphicsSvgItem: public QGraphicsSvgItem
{
public:
	MyQGraphicsSvgItem(QGraphicsItem *parent = Q_NULLPTR) : QGraphicsSvgItem(parent) {QGraphicsSvgItem_QGraphicsSvgItem_QRegisterMetaType();};
	MyQGraphicsSvgItem(const QString &fileName, QGraphicsItem *parent = Q_NULLPTR) : QGraphicsSvgItem(fileName, parent) {QGraphicsSvgItem_QGraphicsSvgItem_QRegisterMetaType();};
	void paint(QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget) { callbackQGraphicsSvgItem_Paint(this, painter, const_cast<QStyleOptionGraphicsItem*>(option), widget); };
	QRectF boundingRect() const { return *static_cast<QRectF*>(callbackQGraphicsSvgItem_BoundingRect(const_cast<void*>(static_cast<const void*>(this)))); };
	int type() const { return callbackQGraphicsSvgItem_Type(const_cast<void*>(static_cast<const void*>(this))); };
	bool event(QEvent * ev) { return callbackQGraphicsSvgItem_Event(this, ev) != 0; };
	void Signal_EnabledChanged() { callbackQGraphicsSvgItem_EnabledChanged(this); };
	void Signal_OpacityChanged() { callbackQGraphicsSvgItem_OpacityChanged(this); };
	void Signal_ParentChanged() { callbackQGraphicsSvgItem_ParentChanged(this); };
	void Signal_RotationChanged() { callbackQGraphicsSvgItem_RotationChanged(this); };
	void Signal_ScaleChanged() { callbackQGraphicsSvgItem_ScaleChanged(this); };
	void updateMicroFocus() { callbackQGraphicsSvgItem_UpdateMicroFocus(this); };
	void Signal_VisibleChanged() { callbackQGraphicsSvgItem_VisibleChanged(this); };
	void Signal_XChanged() { callbackQGraphicsSvgItem_XChanged(this); };
	void Signal_YChanged() { callbackQGraphicsSvgItem_YChanged(this); };
	void Signal_ZChanged() { callbackQGraphicsSvgItem_ZChanged(this); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQGraphicsSvgItem_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQGraphicsSvgItem_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQGraphicsSvgItem_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQGraphicsSvgItem_CustomEvent(this, event); };
	void deleteLater() { callbackQGraphicsSvgItem_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQGraphicsSvgItem_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQGraphicsSvgItem_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtSvg_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQGraphicsSvgItem_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQGraphicsSvgItem_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQGraphicsSvgItem_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	QVariant itemChange(QGraphicsItem::GraphicsItemChange change, const QVariant & value) { return *static_cast<QVariant*>(callbackQGraphicsSvgItem_ItemChange(this, change, const_cast<QVariant*>(&value))); };
	bool sceneEvent(QEvent * event) { return callbackQGraphicsSvgItem_SceneEvent(this, event) != 0; };
	bool sceneEventFilter(QGraphicsItem * watched, QEvent * event) { return callbackQGraphicsSvgItem_SceneEventFilter(this, watched, event) != 0; };
	void advance(int phase) { callbackQGraphicsSvgItem_Advance(this, phase); };
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
	void keyPressEvent(QKeyEvent * event) { callbackQGraphicsSvgItem_KeyPressEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQGraphicsSvgItem_KeyReleaseEvent(this, event); };
	void mouseDoubleClickEvent(QGraphicsSceneMouseEvent * event) { callbackQGraphicsSvgItem_MouseDoubleClickEvent(this, event); };
	void mouseMoveEvent(QGraphicsSceneMouseEvent * event) { callbackQGraphicsSvgItem_MouseMoveEvent(this, event); };
	void mousePressEvent(QGraphicsSceneMouseEvent * event) { callbackQGraphicsSvgItem_MousePressEvent(this, event); };
	void mouseReleaseEvent(QGraphicsSceneMouseEvent * event) { callbackQGraphicsSvgItem_MouseReleaseEvent(this, event); };
	void wheelEvent(QGraphicsSceneWheelEvent * event) { callbackQGraphicsSvgItem_WheelEvent(this, event); };
	QPainterPath opaqueArea() const { return *static_cast<QPainterPath*>(callbackQGraphicsSvgItem_OpaqueArea(const_cast<void*>(static_cast<const void*>(this)))); };
	QPainterPath shape() const { return *static_cast<QPainterPath*>(callbackQGraphicsSvgItem_Shape(const_cast<void*>(static_cast<const void*>(this)))); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackQGraphicsSvgItem_InputMethodQuery(const_cast<void*>(static_cast<const void*>(this)), query)); };
	bool collidesWithItem(const QGraphicsItem * other, Qt::ItemSelectionMode mode) const { return callbackQGraphicsSvgItem_CollidesWithItem(const_cast<void*>(static_cast<const void*>(this)), const_cast<QGraphicsItem*>(other), mode) != 0; };
	bool collidesWithPath(const QPainterPath & path, Qt::ItemSelectionMode mode) const { return callbackQGraphicsSvgItem_CollidesWithPath(const_cast<void*>(static_cast<const void*>(this)), const_cast<QPainterPath*>(&path), mode) != 0; };
	bool contains(const QPointF & point) const { return callbackQGraphicsSvgItem_Contains(const_cast<void*>(static_cast<const void*>(this)), const_cast<QPointF*>(&point)) != 0; };
	bool isObscuredBy(const QGraphicsItem * item) const { return callbackQGraphicsSvgItem_IsObscuredBy(const_cast<void*>(static_cast<const void*>(this)), const_cast<QGraphicsItem*>(item)) != 0; };
};

Q_DECLARE_METATYPE(MyQGraphicsSvgItem*)

int QGraphicsSvgItem_QGraphicsSvgItem_QRegisterMetaType(){qRegisterMetaType<QGraphicsSvgItem*>(); return qRegisterMetaType<MyQGraphicsSvgItem*>();}

void* QGraphicsSvgItem_NewQGraphicsSvgItem(void* parent)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQGraphicsSvgItem(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQGraphicsSvgItem(static_cast<QGraphicsWidget*>(parent));
	} else {
		return new MyQGraphicsSvgItem(static_cast<QGraphicsItem*>(parent));
	}
}

void* QGraphicsSvgItem_NewQGraphicsSvgItem2(struct QtSvg_PackedString fileName, void* parent)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQGraphicsSvgItem(QString::fromUtf8(fileName.data, fileName.len), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQGraphicsSvgItem(QString::fromUtf8(fileName.data, fileName.len), static_cast<QGraphicsWidget*>(parent));
	} else {
		return new MyQGraphicsSvgItem(QString::fromUtf8(fileName.data, fileName.len), static_cast<QGraphicsItem*>(parent));
	}
}

void QGraphicsSvgItem_Paint(void* ptr, void* painter, void* option, void* widget)
{
	static_cast<QGraphicsSvgItem*>(ptr)->paint(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
}

void QGraphicsSvgItem_PaintDefault(void* ptr, void* painter, void* option, void* widget)
{
		static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::paint(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
}

void QGraphicsSvgItem_SetElementId(void* ptr, struct QtSvg_PackedString id)
{
	static_cast<QGraphicsSvgItem*>(ptr)->setElementId(QString::fromUtf8(id.data, id.len));
}

void QGraphicsSvgItem_SetMaximumCacheSize(void* ptr, void* size)
{
	static_cast<QGraphicsSvgItem*>(ptr)->setMaximumCacheSize(*static_cast<QSize*>(size));
}

void QGraphicsSvgItem_SetSharedRenderer(void* ptr, void* renderer)
{
	static_cast<QGraphicsSvgItem*>(ptr)->setSharedRenderer(static_cast<QSvgRenderer*>(renderer));
}

void* QGraphicsSvgItem_BoundingRect(void* ptr)
{
	return ({ QRectF tmpValue = static_cast<QGraphicsSvgItem*>(ptr)->boundingRect(); new QRectF(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void* QGraphicsSvgItem_BoundingRectDefault(void* ptr)
{
		return ({ QRectF tmpValue = static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::boundingRect(); new QRectF(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void* QGraphicsSvgItem_MaximumCacheSize(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QGraphicsSvgItem*>(ptr)->maximumCacheSize(); new QSize(tmpValue.width(), tmpValue.height()); });
}

struct QtSvg_PackedString QGraphicsSvgItem_ElementId(void* ptr)
{
	return ({ QByteArray t96b361 = static_cast<QGraphicsSvgItem*>(ptr)->elementId().toUtf8(); QtSvg_PackedString { const_cast<char*>(t96b361.prepend("WHITESPACE").constData()+10), t96b361.size()-10 }; });
}

void* QGraphicsSvgItem_Renderer(void* ptr)
{
	return static_cast<QGraphicsSvgItem*>(ptr)->renderer();
}

int QGraphicsSvgItem_TypeDefault(void* ptr)
{
		return static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::type();
}

void* QGraphicsSvgItem___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QGraphicsSvgItem___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QGraphicsSvgItem___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QGraphicsSvgItem___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QGraphicsSvgItem___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGraphicsSvgItem___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QGraphicsSvgItem___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QGraphicsSvgItem___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGraphicsSvgItem___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QGraphicsSvgItem___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QGraphicsSvgItem___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGraphicsSvgItem___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QGraphicsSvgItem___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QGraphicsSvgItem___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGraphicsSvgItem___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

void* QGraphicsSvgItem___setTransformations_transformations_atList(void* ptr, int i)
{
	return const_cast<QGraphicsTransform*>(static_cast<QList<QGraphicsTransform *>*>(ptr)->at(i));
}

void QGraphicsSvgItem___setTransformations_transformations_setList(void* ptr, void* i)
{
	static_cast<QList<QGraphicsTransform *>*>(ptr)->append(static_cast<QGraphicsTransform*>(i));
}

void* QGraphicsSvgItem___setTransformations_transformations_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QGraphicsTransform *>;
}

void* QGraphicsSvgItem___childItems_atList(void* ptr, int i)
{
	return const_cast<QGraphicsItem*>(static_cast<QList<QGraphicsItem *>*>(ptr)->at(i));
}

void QGraphicsSvgItem___childItems_setList(void* ptr, void* i)
{
	static_cast<QList<QGraphicsItem *>*>(ptr)->append(static_cast<QGraphicsItem*>(i));
}

void* QGraphicsSvgItem___childItems_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QGraphicsItem *>;
}

void* QGraphicsSvgItem___collidingItems_atList(void* ptr, int i)
{
	return const_cast<QGraphicsItem*>(static_cast<QList<QGraphicsItem *>*>(ptr)->at(i));
}

void QGraphicsSvgItem___collidingItems_setList(void* ptr, void* i)
{
	static_cast<QList<QGraphicsItem *>*>(ptr)->append(static_cast<QGraphicsItem*>(i));
}

void* QGraphicsSvgItem___collidingItems_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QGraphicsItem *>;
}

void* QGraphicsSvgItem___transformations_atList(void* ptr, int i)
{
	return const_cast<QGraphicsTransform*>(static_cast<QList<QGraphicsTransform *>*>(ptr)->at(i));
}

void QGraphicsSvgItem___transformations_setList(void* ptr, void* i)
{
	static_cast<QList<QGraphicsTransform *>*>(ptr)->append(static_cast<QGraphicsTransform*>(i));
}

void* QGraphicsSvgItem___transformations_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QGraphicsTransform *>;
}

char QGraphicsSvgItem_EventDefault(void* ptr, void* ev)
{
		return static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::event(static_cast<QEvent*>(ev));
}

void QGraphicsSvgItem_UpdateMicroFocusDefault(void* ptr)
{
		static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::updateMicroFocus();
}

char QGraphicsSvgItem_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QGraphicsSvgItem_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::childEvent(static_cast<QChildEvent*>(event));
}

void QGraphicsSvgItem_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGraphicsSvgItem_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::customEvent(static_cast<QEvent*>(event));
}

void QGraphicsSvgItem_DeleteLaterDefault(void* ptr)
{
		static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::deleteLater();
}

void QGraphicsSvgItem_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGraphicsSvgItem_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QGraphicsSvgItem_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::metaObject());
}

void* QGraphicsSvgItem_ItemChangeDefault(void* ptr, long long change, void* value)
{
		return new QVariant(static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::itemChange(static_cast<QGraphicsItem::GraphicsItemChange>(change), *static_cast<QVariant*>(value)));
}

char QGraphicsSvgItem_SceneEventDefault(void* ptr, void* event)
{
		return static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::sceneEvent(static_cast<QEvent*>(event));
}

char QGraphicsSvgItem_SceneEventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::sceneEventFilter(static_cast<QGraphicsItem*>(watched), static_cast<QEvent*>(event));
}

void QGraphicsSvgItem_AdvanceDefault(void* ptr, int phase)
{
		static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::advance(phase);
}

void QGraphicsSvgItem_ContextMenuEventDefault(void* ptr, void* event)
{
		static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::contextMenuEvent(static_cast<QGraphicsSceneContextMenuEvent*>(event));
}

void QGraphicsSvgItem_DragEnterEventDefault(void* ptr, void* event)
{
		static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::dragEnterEvent(static_cast<QGraphicsSceneDragDropEvent*>(event));
}

void QGraphicsSvgItem_DragLeaveEventDefault(void* ptr, void* event)
{
		static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::dragLeaveEvent(static_cast<QGraphicsSceneDragDropEvent*>(event));
}

void QGraphicsSvgItem_DragMoveEventDefault(void* ptr, void* event)
{
		static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::dragMoveEvent(static_cast<QGraphicsSceneDragDropEvent*>(event));
}

void QGraphicsSvgItem_DropEventDefault(void* ptr, void* event)
{
		static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::dropEvent(static_cast<QGraphicsSceneDragDropEvent*>(event));
}

void QGraphicsSvgItem_FocusInEventDefault(void* ptr, void* event)
{
		static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::focusInEvent(static_cast<QFocusEvent*>(event));
}

void QGraphicsSvgItem_FocusOutEventDefault(void* ptr, void* event)
{
		static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QGraphicsSvgItem_HoverEnterEventDefault(void* ptr, void* event)
{
		static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::hoverEnterEvent(static_cast<QGraphicsSceneHoverEvent*>(event));
}

void QGraphicsSvgItem_HoverLeaveEventDefault(void* ptr, void* event)
{
		static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::hoverLeaveEvent(static_cast<QGraphicsSceneHoverEvent*>(event));
}

void QGraphicsSvgItem_HoverMoveEventDefault(void* ptr, void* event)
{
		static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::hoverMoveEvent(static_cast<QGraphicsSceneHoverEvent*>(event));
}

void QGraphicsSvgItem_InputMethodEventDefault(void* ptr, void* event)
{
		static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void QGraphicsSvgItem_KeyPressEventDefault(void* ptr, void* event)
{
		static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QGraphicsSvgItem_KeyReleaseEventDefault(void* ptr, void* event)
{
		static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QGraphicsSvgItem_MouseDoubleClickEventDefault(void* ptr, void* event)
{
		static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::mouseDoubleClickEvent(static_cast<QGraphicsSceneMouseEvent*>(event));
}

void QGraphicsSvgItem_MouseMoveEventDefault(void* ptr, void* event)
{
		static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::mouseMoveEvent(static_cast<QGraphicsSceneMouseEvent*>(event));
}

void QGraphicsSvgItem_MousePressEventDefault(void* ptr, void* event)
{
		static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::mousePressEvent(static_cast<QGraphicsSceneMouseEvent*>(event));
}

void QGraphicsSvgItem_MouseReleaseEventDefault(void* ptr, void* event)
{
		static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::mouseReleaseEvent(static_cast<QGraphicsSceneMouseEvent*>(event));
}

void QGraphicsSvgItem_WheelEventDefault(void* ptr, void* event)
{
		static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::wheelEvent(static_cast<QGraphicsSceneWheelEvent*>(event));
}

void* QGraphicsSvgItem_OpaqueAreaDefault(void* ptr)
{
		return new QPainterPath(static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::opaqueArea());
}

void* QGraphicsSvgItem_ShapeDefault(void* ptr)
{
		return new QPainterPath(static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::shape());
}

void* QGraphicsSvgItem_InputMethodQueryDefault(void* ptr, long long query)
{
		return new QVariant(static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

char QGraphicsSvgItem_CollidesWithItemDefault(void* ptr, void* other, long long mode)
{
		return static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::collidesWithItem(static_cast<QGraphicsItem*>(other), static_cast<Qt::ItemSelectionMode>(mode));
}

char QGraphicsSvgItem_CollidesWithPathDefault(void* ptr, void* path, long long mode)
{
		return static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::collidesWithPath(*static_cast<QPainterPath*>(path), static_cast<Qt::ItemSelectionMode>(mode));
}

char QGraphicsSvgItem_ContainsDefault(void* ptr, void* point)
{
		return static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::contains(*static_cast<QPointF*>(point));
}

char QGraphicsSvgItem_IsObscuredByDefault(void* ptr, void* item)
{
		return static_cast<QGraphicsSvgItem*>(ptr)->QGraphicsSvgItem::isObscuredBy(static_cast<QGraphicsItem*>(item));
}

class MyQSvgGenerator: public QSvgGenerator
{
public:
	MyQSvgGenerator() : QSvgGenerator() {};
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQSvgGenerator_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
	
};

void* QSvgGenerator_NewQSvgGenerator()
{
	return new MyQSvgGenerator();
}

void QSvgGenerator_SetDescription(void* ptr, struct QtSvg_PackedString description)
{
	static_cast<QSvgGenerator*>(ptr)->setDescription(QString::fromUtf8(description.data, description.len));
}

void QSvgGenerator_SetFileName(void* ptr, struct QtSvg_PackedString fileName)
{
	static_cast<QSvgGenerator*>(ptr)->setFileName(QString::fromUtf8(fileName.data, fileName.len));
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

void QSvgGenerator_SetTitle(void* ptr, struct QtSvg_PackedString title)
{
	static_cast<QSvgGenerator*>(ptr)->setTitle(QString::fromUtf8(title.data, title.len));
}

void QSvgGenerator_SetViewBox(void* ptr, void* viewBox)
{
	static_cast<QSvgGenerator*>(ptr)->setViewBox(*static_cast<QRect*>(viewBox));
}

void QSvgGenerator_SetViewBox2(void* ptr, void* viewBox)
{
	static_cast<QSvgGenerator*>(ptr)->setViewBox(*static_cast<QRectF*>(viewBox));
}

void QSvgGenerator_DestroyQSvgGenerator(void* ptr)
{
	static_cast<QSvgGenerator*>(ptr)->~QSvgGenerator();
}

void* QSvgGenerator_OutputDevice(void* ptr)
{
	return static_cast<QSvgGenerator*>(ptr)->outputDevice();
}

void* QSvgGenerator_PaintEngine(void* ptr)
{
	return static_cast<QSvgGenerator*>(ptr)->paintEngine();
}

void* QSvgGenerator_PaintEngineDefault(void* ptr)
{
		return static_cast<QSvgGenerator*>(ptr)->QSvgGenerator::paintEngine();
}

void* QSvgGenerator_ViewBox(void* ptr)
{
	return ({ QRect tmpValue = static_cast<QSvgGenerator*>(ptr)->viewBox(); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void* QSvgGenerator_ViewBoxF(void* ptr)
{
	return ({ QRectF tmpValue = static_cast<QSvgGenerator*>(ptr)->viewBoxF(); new QRectF(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void* QSvgGenerator_Size(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QSvgGenerator*>(ptr)->size(); new QSize(tmpValue.width(), tmpValue.height()); });
}

struct QtSvg_PackedString QSvgGenerator_Description(void* ptr)
{
	return ({ QByteArray t85d23d = static_cast<QSvgGenerator*>(ptr)->description().toUtf8(); QtSvg_PackedString { const_cast<char*>(t85d23d.prepend("WHITESPACE").constData()+10), t85d23d.size()-10 }; });
}

struct QtSvg_PackedString QSvgGenerator_FileName(void* ptr)
{
	return ({ QByteArray ta251c4 = static_cast<QSvgGenerator*>(ptr)->fileName().toUtf8(); QtSvg_PackedString { const_cast<char*>(ta251c4.prepend("WHITESPACE").constData()+10), ta251c4.size()-10 }; });
}

struct QtSvg_PackedString QSvgGenerator_Title(void* ptr)
{
	return ({ QByteArray tb87794 = static_cast<QSvgGenerator*>(ptr)->title().toUtf8(); QtSvg_PackedString { const_cast<char*>(tb87794.prepend("WHITESPACE").constData()+10), tb87794.size()-10 }; });
}

int QSvgGenerator_Resolution(void* ptr)
{
	return static_cast<QSvgGenerator*>(ptr)->resolution();
}



class MyQSvgRenderer: public QSvgRenderer
{
public:
	MyQSvgRenderer(QObject *parent = Q_NULLPTR) : QSvgRenderer(parent) {QSvgRenderer_QSvgRenderer_QRegisterMetaType();};
	MyQSvgRenderer(QXmlStreamReader *contents, QObject *parent = Q_NULLPTR) : QSvgRenderer(contents, parent) {QSvgRenderer_QSvgRenderer_QRegisterMetaType();};
	MyQSvgRenderer(const QByteArray &contents, QObject *parent = Q_NULLPTR) : QSvgRenderer(contents, parent) {QSvgRenderer_QSvgRenderer_QRegisterMetaType();};
	MyQSvgRenderer(const QString &filename, QObject *parent = Q_NULLPTR) : QSvgRenderer(filename, parent) {QSvgRenderer_QSvgRenderer_QRegisterMetaType();};
	bool load(QXmlStreamReader * contents) { return callbackQSvgRenderer_Load3(this, contents) != 0; };
	bool load(const QByteArray & contents) { return callbackQSvgRenderer_Load2(this, const_cast<QByteArray*>(&contents)) != 0; };
	bool load(const QString & filename) { QByteArray t08deae = filename.toUtf8(); QtSvg_PackedString filenamePacked = { const_cast<char*>(t08deae.prepend("WHITESPACE").constData()+10), t08deae.size()-10 };return callbackQSvgRenderer_Load(this, filenamePacked) != 0; };
	void render(QPainter * painter) { callbackQSvgRenderer_Render(this, painter); };
	void render(QPainter * painter, const QRectF & bounds) { callbackQSvgRenderer_Render2(this, painter, const_cast<QRectF*>(&bounds)); };
	void render(QPainter * painter, const QString & elementId, const QRectF & bounds) { QByteArray t5b5420 = elementId.toUtf8(); QtSvg_PackedString elementIdPacked = { const_cast<char*>(t5b5420.prepend("WHITESPACE").constData()+10), t5b5420.size()-10 };callbackQSvgRenderer_Render3(this, painter, elementIdPacked, const_cast<QRectF*>(&bounds)); };
	void Signal_RepaintNeeded() { callbackQSvgRenderer_RepaintNeeded(this); };
	bool event(QEvent * e) { return callbackQSvgRenderer_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSvgRenderer_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQSvgRenderer_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSvgRenderer_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQSvgRenderer_CustomEvent(this, event); };
	void deleteLater() { callbackQSvgRenderer_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQSvgRenderer_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSvgRenderer_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtSvg_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQSvgRenderer_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQSvgRenderer_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSvgRenderer_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQSvgRenderer*)

int QSvgRenderer_QSvgRenderer_QRegisterMetaType(){qRegisterMetaType<QSvgRenderer*>(); return qRegisterMetaType<MyQSvgRenderer*>();}

void* QSvgRenderer_NewQSvgRenderer(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQSvgRenderer(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQSvgRenderer(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQSvgRenderer(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQSvgRenderer(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQSvgRenderer(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQSvgRenderer(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQSvgRenderer(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQSvgRenderer(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQSvgRenderer(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQSvgRenderer(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQSvgRenderer(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQSvgRenderer(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQSvgRenderer(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQSvgRenderer(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQSvgRenderer(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQSvgRenderer(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQSvgRenderer(static_cast<QWindow*>(parent));
	} else {
		return new MyQSvgRenderer(static_cast<QObject*>(parent));
	}
}

void* QSvgRenderer_NewQSvgRenderer4(void* contents, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQSvgRenderer(static_cast<QXmlStreamReader*>(contents), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQSvgRenderer(static_cast<QXmlStreamReader*>(contents), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQSvgRenderer(static_cast<QXmlStreamReader*>(contents), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQSvgRenderer(static_cast<QXmlStreamReader*>(contents), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQSvgRenderer(static_cast<QXmlStreamReader*>(contents), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQSvgRenderer(static_cast<QXmlStreamReader*>(contents), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQSvgRenderer(static_cast<QXmlStreamReader*>(contents), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQSvgRenderer(static_cast<QXmlStreamReader*>(contents), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQSvgRenderer(static_cast<QXmlStreamReader*>(contents), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQSvgRenderer(static_cast<QXmlStreamReader*>(contents), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQSvgRenderer(static_cast<QXmlStreamReader*>(contents), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQSvgRenderer(static_cast<QXmlStreamReader*>(contents), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQSvgRenderer(static_cast<QXmlStreamReader*>(contents), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQSvgRenderer(static_cast<QXmlStreamReader*>(contents), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQSvgRenderer(static_cast<QXmlStreamReader*>(contents), static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQSvgRenderer(static_cast<QXmlStreamReader*>(contents), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQSvgRenderer(static_cast<QXmlStreamReader*>(contents), static_cast<QWindow*>(parent));
	} else {
		return new MyQSvgRenderer(static_cast<QXmlStreamReader*>(contents), static_cast<QObject*>(parent));
	}
}

void* QSvgRenderer_NewQSvgRenderer3(void* contents, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQSvgRenderer(*static_cast<QByteArray*>(contents), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQSvgRenderer(*static_cast<QByteArray*>(contents), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQSvgRenderer(*static_cast<QByteArray*>(contents), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQSvgRenderer(*static_cast<QByteArray*>(contents), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQSvgRenderer(*static_cast<QByteArray*>(contents), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQSvgRenderer(*static_cast<QByteArray*>(contents), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQSvgRenderer(*static_cast<QByteArray*>(contents), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQSvgRenderer(*static_cast<QByteArray*>(contents), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQSvgRenderer(*static_cast<QByteArray*>(contents), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQSvgRenderer(*static_cast<QByteArray*>(contents), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQSvgRenderer(*static_cast<QByteArray*>(contents), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQSvgRenderer(*static_cast<QByteArray*>(contents), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQSvgRenderer(*static_cast<QByteArray*>(contents), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQSvgRenderer(*static_cast<QByteArray*>(contents), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQSvgRenderer(*static_cast<QByteArray*>(contents), static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQSvgRenderer(*static_cast<QByteArray*>(contents), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQSvgRenderer(*static_cast<QByteArray*>(contents), static_cast<QWindow*>(parent));
	} else {
		return new MyQSvgRenderer(*static_cast<QByteArray*>(contents), static_cast<QObject*>(parent));
	}
}

void* QSvgRenderer_NewQSvgRenderer2(struct QtSvg_PackedString filename, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQSvgRenderer(QString::fromUtf8(filename.data, filename.len), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQSvgRenderer(QString::fromUtf8(filename.data, filename.len), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQSvgRenderer(QString::fromUtf8(filename.data, filename.len), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQSvgRenderer(QString::fromUtf8(filename.data, filename.len), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQSvgRenderer(QString::fromUtf8(filename.data, filename.len), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQSvgRenderer(QString::fromUtf8(filename.data, filename.len), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQSvgRenderer(QString::fromUtf8(filename.data, filename.len), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQSvgRenderer(QString::fromUtf8(filename.data, filename.len), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQSvgRenderer(QString::fromUtf8(filename.data, filename.len), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQSvgRenderer(QString::fromUtf8(filename.data, filename.len), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQSvgRenderer(QString::fromUtf8(filename.data, filename.len), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQSvgRenderer(QString::fromUtf8(filename.data, filename.len), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQSvgRenderer(QString::fromUtf8(filename.data, filename.len), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQSvgRenderer(QString::fromUtf8(filename.data, filename.len), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQSvgRenderer(QString::fromUtf8(filename.data, filename.len), static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQSvgRenderer(QString::fromUtf8(filename.data, filename.len), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQSvgRenderer(QString::fromUtf8(filename.data, filename.len), static_cast<QWindow*>(parent));
	} else {
		return new MyQSvgRenderer(QString::fromUtf8(filename.data, filename.len), static_cast<QObject*>(parent));
	}
}

char QSvgRenderer_Load3(void* ptr, void* contents)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QSvgRenderer*>(ptr), "load", Q_RETURN_ARG(bool, returnArg), Q_ARG(QXmlStreamReader*, static_cast<QXmlStreamReader*>(contents)));
	return returnArg;
}

char QSvgRenderer_Load3Default(void* ptr, void* contents)
{
		return static_cast<QSvgRenderer*>(ptr)->QSvgRenderer::load(static_cast<QXmlStreamReader*>(contents));
}

char QSvgRenderer_Load2(void* ptr, void* contents)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QSvgRenderer*>(ptr), "load", Q_RETURN_ARG(bool, returnArg), Q_ARG(const QByteArray, *static_cast<QByteArray*>(contents)));
	return returnArg;
}

char QSvgRenderer_Load2Default(void* ptr, void* contents)
{
		return static_cast<QSvgRenderer*>(ptr)->QSvgRenderer::load(*static_cast<QByteArray*>(contents));
}

char QSvgRenderer_Load(void* ptr, struct QtSvg_PackedString filename)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QSvgRenderer*>(ptr), "load", Q_RETURN_ARG(bool, returnArg), Q_ARG(const QString, QString::fromUtf8(filename.data, filename.len)));
	return returnArg;
}

char QSvgRenderer_LoadDefault(void* ptr, struct QtSvg_PackedString filename)
{
		return static_cast<QSvgRenderer*>(ptr)->QSvgRenderer::load(QString::fromUtf8(filename.data, filename.len));
}

void QSvgRenderer_Render(void* ptr, void* painter)
{
	QMetaObject::invokeMethod(static_cast<QSvgRenderer*>(ptr), "render", Q_ARG(QPainter*, static_cast<QPainter*>(painter)));
}

void QSvgRenderer_RenderDefault(void* ptr, void* painter)
{
		static_cast<QSvgRenderer*>(ptr)->QSvgRenderer::render(static_cast<QPainter*>(painter));
}

void QSvgRenderer_Render2(void* ptr, void* painter, void* bounds)
{
	QMetaObject::invokeMethod(static_cast<QSvgRenderer*>(ptr), "render", Q_ARG(QPainter*, static_cast<QPainter*>(painter)), Q_ARG(const QRectF, *static_cast<QRectF*>(bounds)));
}

void QSvgRenderer_Render2Default(void* ptr, void* painter, void* bounds)
{
		static_cast<QSvgRenderer*>(ptr)->QSvgRenderer::render(static_cast<QPainter*>(painter), *static_cast<QRectF*>(bounds));
}

void QSvgRenderer_Render3(void* ptr, void* painter, struct QtSvg_PackedString elementId, void* bounds)
{
	QMetaObject::invokeMethod(static_cast<QSvgRenderer*>(ptr), "render", Q_ARG(QPainter*, static_cast<QPainter*>(painter)), Q_ARG(const QString, QString::fromUtf8(elementId.data, elementId.len)), Q_ARG(const QRectF, *static_cast<QRectF*>(bounds)));
}

void QSvgRenderer_Render3Default(void* ptr, void* painter, struct QtSvg_PackedString elementId, void* bounds)
{
		static_cast<QSvgRenderer*>(ptr)->QSvgRenderer::render(static_cast<QPainter*>(painter), QString::fromUtf8(elementId.data, elementId.len), *static_cast<QRectF*>(bounds));
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

void QSvgRenderer_DestroyQSvgRenderer(void* ptr)
{
	static_cast<QSvgRenderer*>(ptr)->~QSvgRenderer();
}

void* QSvgRenderer_ViewBox(void* ptr)
{
	return ({ QRect tmpValue = static_cast<QSvgRenderer*>(ptr)->viewBox(); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void* QSvgRenderer_BoundsOnElement(void* ptr, struct QtSvg_PackedString id)
{
	return ({ QRectF tmpValue = static_cast<QSvgRenderer*>(ptr)->boundsOnElement(QString::fromUtf8(id.data, id.len)); new QRectF(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void* QSvgRenderer_ViewBoxF(void* ptr)
{
	return ({ QRectF tmpValue = static_cast<QSvgRenderer*>(ptr)->viewBoxF(); new QRectF(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void* QSvgRenderer_DefaultSize(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QSvgRenderer*>(ptr)->defaultSize(); new QSize(tmpValue.width(), tmpValue.height()); });
}

char QSvgRenderer_Animated(void* ptr)
{
	return static_cast<QSvgRenderer*>(ptr)->animated();
}

char QSvgRenderer_ElementExists(void* ptr, struct QtSvg_PackedString id)
{
	return static_cast<QSvgRenderer*>(ptr)->elementExists(QString::fromUtf8(id.data, id.len));
}

char QSvgRenderer_IsValid(void* ptr)
{
	return static_cast<QSvgRenderer*>(ptr)->isValid();
}

int QSvgRenderer_FramesPerSecond(void* ptr)
{
	return static_cast<QSvgRenderer*>(ptr)->framesPerSecond();
}

void* QSvgRenderer___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QSvgRenderer___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QSvgRenderer___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QSvgRenderer___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QSvgRenderer___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSvgRenderer___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QSvgRenderer___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QSvgRenderer___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSvgRenderer___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QSvgRenderer___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QSvgRenderer___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSvgRenderer___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QSvgRenderer___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QSvgRenderer___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSvgRenderer___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

char QSvgRenderer_EventDefault(void* ptr, void* e)
{
		return static_cast<QSvgRenderer*>(ptr)->QSvgRenderer::event(static_cast<QEvent*>(e));
}

char QSvgRenderer_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QSvgRenderer*>(ptr)->QSvgRenderer::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QSvgRenderer_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QSvgRenderer*>(ptr)->QSvgRenderer::childEvent(static_cast<QChildEvent*>(event));
}

void QSvgRenderer_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QSvgRenderer*>(ptr)->QSvgRenderer::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSvgRenderer_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QSvgRenderer*>(ptr)->QSvgRenderer::customEvent(static_cast<QEvent*>(event));
}

void QSvgRenderer_DeleteLaterDefault(void* ptr)
{
		static_cast<QSvgRenderer*>(ptr)->QSvgRenderer::deleteLater();
}

void QSvgRenderer_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QSvgRenderer*>(ptr)->QSvgRenderer::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSvgRenderer_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QSvgRenderer*>(ptr)->QSvgRenderer::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QSvgRenderer_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QSvgRenderer*>(ptr)->QSvgRenderer::metaObject());
}

class MyQSvgWidget: public QSvgWidget
{
public:
	MyQSvgWidget(QWidget *parent = Q_NULLPTR) : QSvgWidget(parent) {QSvgWidget_QSvgWidget_QRegisterMetaType();};
	MyQSvgWidget(const QString &file, QWidget *parent = Q_NULLPTR) : QSvgWidget(file, parent) {QSvgWidget_QSvgWidget_QRegisterMetaType();};
	void load(const QByteArray & contents) { callbackQSvgWidget_Load2(this, const_cast<QByteArray*>(&contents)); };
	void load(const QString & file) { QByteArray t971c41 = file.toUtf8(); QtSvg_PackedString filePacked = { const_cast<char*>(t971c41.prepend("WHITESPACE").constData()+10), t971c41.size()-10 };callbackQSvgWidget_Load(this, filePacked); };
	void paintEvent(QPaintEvent * event) { callbackQSvgWidget_PaintEvent(this, event); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQSvgWidget_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	bool close() { return callbackQSvgWidget_Close(this) != 0; };
	bool event(QEvent * event) { return callbackQSvgWidget_Event(this, event) != 0; };
	bool focusNextPrevChild(bool next) { return callbackQSvgWidget_FocusNextPrevChild(this, next) != 0; };
	bool nativeEvent(const QByteArray & eventType, void * message, long * result) { return callbackQSvgWidget_NativeEvent(this, const_cast<QByteArray*>(&eventType), message, *result) != 0; };
	void actionEvent(QActionEvent * event) { callbackQSvgWidget_ActionEvent(this, event); };
	void changeEvent(QEvent * event) { callbackQSvgWidget_ChangeEvent(this, event); };
	void closeEvent(QCloseEvent * event) { callbackQSvgWidget_CloseEvent(this, event); };
	void contextMenuEvent(QContextMenuEvent * event) { callbackQSvgWidget_ContextMenuEvent(this, event); };
	void Signal_CustomContextMenuRequested(const QPoint & pos) { callbackQSvgWidget_CustomContextMenuRequested(this, const_cast<QPoint*>(&pos)); };
	void dragEnterEvent(QDragEnterEvent * event) { callbackQSvgWidget_DragEnterEvent(this, event); };
	void dragLeaveEvent(QDragLeaveEvent * event) { callbackQSvgWidget_DragLeaveEvent(this, event); };
	void dragMoveEvent(QDragMoveEvent * event) { callbackQSvgWidget_DragMoveEvent(this, event); };
	void dropEvent(QDropEvent * event) { callbackQSvgWidget_DropEvent(this, event); };
	void enterEvent(QEvent * event) { callbackQSvgWidget_EnterEvent(this, event); };
	void focusInEvent(QFocusEvent * event) { callbackQSvgWidget_FocusInEvent(this, event); };
	void focusOutEvent(QFocusEvent * event) { callbackQSvgWidget_FocusOutEvent(this, event); };
	void hide() { callbackQSvgWidget_Hide(this); };
	void hideEvent(QHideEvent * event) { callbackQSvgWidget_HideEvent(this, event); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackQSvgWidget_InputMethodEvent(this, event); };
	void keyPressEvent(QKeyEvent * event) { callbackQSvgWidget_KeyPressEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQSvgWidget_KeyReleaseEvent(this, event); };
	void leaveEvent(QEvent * event) { callbackQSvgWidget_LeaveEvent(this, event); };
	void lower() { callbackQSvgWidget_Lower(this); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQSvgWidget_MouseDoubleClickEvent(this, event); };
	void mouseMoveEvent(QMouseEvent * event) { callbackQSvgWidget_MouseMoveEvent(this, event); };
	void mousePressEvent(QMouseEvent * event) { callbackQSvgWidget_MousePressEvent(this, event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackQSvgWidget_MouseReleaseEvent(this, event); };
	void moveEvent(QMoveEvent * event) { callbackQSvgWidget_MoveEvent(this, event); };
	void raise() { callbackQSvgWidget_Raise(this); };
	void repaint() { callbackQSvgWidget_Repaint(this); };
	void resizeEvent(QResizeEvent * event) { callbackQSvgWidget_ResizeEvent(this, event); };
	void setDisabled(bool disable) { callbackQSvgWidget_SetDisabled(this, disable); };
	void setEnabled(bool vbo) { callbackQSvgWidget_SetEnabled(this, vbo); };
	void setFocus() { callbackQSvgWidget_SetFocus2(this); };
	void setHidden(bool hidden) { callbackQSvgWidget_SetHidden(this, hidden); };
	void setStyleSheet(const QString & styleSheet) { QByteArray t728ae7 = styleSheet.toUtf8(); QtSvg_PackedString styleSheetPacked = { const_cast<char*>(t728ae7.prepend("WHITESPACE").constData()+10), t728ae7.size()-10 };callbackQSvgWidget_SetStyleSheet(this, styleSheetPacked); };
	void setVisible(bool visible) { callbackQSvgWidget_SetVisible(this, visible); };
	void setWindowModified(bool vbo) { callbackQSvgWidget_SetWindowModified(this, vbo); };
	void setWindowTitle(const QString & vqs) { QByteArray tda39a3 = vqs.toUtf8(); QtSvg_PackedString vqsPacked = { const_cast<char*>(tda39a3.prepend("WHITESPACE").constData()+10), tda39a3.size()-10 };callbackQSvgWidget_SetWindowTitle(this, vqsPacked); };
	void show() { callbackQSvgWidget_Show(this); };
	void showEvent(QShowEvent * event) { callbackQSvgWidget_ShowEvent(this, event); };
	void showFullScreen() { callbackQSvgWidget_ShowFullScreen(this); };
	void showMaximized() { callbackQSvgWidget_ShowMaximized(this); };
	void showMinimized() { callbackQSvgWidget_ShowMinimized(this); };
	void showNormal() { callbackQSvgWidget_ShowNormal(this); };
	void tabletEvent(QTabletEvent * event) { callbackQSvgWidget_TabletEvent(this, event); };
	void update() { callbackQSvgWidget_Update(this); };
	void updateMicroFocus() { callbackQSvgWidget_UpdateMicroFocus(this); };
	void wheelEvent(QWheelEvent * event) { callbackQSvgWidget_WheelEvent(this, event); };
	void Signal_WindowIconChanged(const QIcon & icon) { callbackQSvgWidget_WindowIconChanged(this, const_cast<QIcon*>(&icon)); };
	void Signal_WindowTitleChanged(const QString & title) { QByteArray t3c6de1 = title.toUtf8(); QtSvg_PackedString titlePacked = { const_cast<char*>(t3c6de1.prepend("WHITESPACE").constData()+10), t3c6de1.size()-10 };callbackQSvgWidget_WindowTitleChanged(this, titlePacked); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQSvgWidget_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize minimumSizeHint() const { return *static_cast<QSize*>(callbackQSvgWidget_MinimumSizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackQSvgWidget_InputMethodQuery(const_cast<void*>(static_cast<const void*>(this)), query)); };
	bool hasHeightForWidth() const { return callbackQSvgWidget_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int heightForWidth(int w) const { return callbackQSvgWidget_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackQSvgWidget_Metric(const_cast<void*>(static_cast<const void*>(this)), m); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSvgWidget_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQSvgWidget_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSvgWidget_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQSvgWidget_CustomEvent(this, event); };
	void deleteLater() { callbackQSvgWidget_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQSvgWidget_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSvgWidget_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtSvg_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQSvgWidget_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQSvgWidget_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSvgWidget_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQSvgWidget*)

int QSvgWidget_QSvgWidget_QRegisterMetaType(){qRegisterMetaType<QSvgWidget*>(); return qRegisterMetaType<MyQSvgWidget*>();}

void* QSvgWidget_NewQSvgWidget(void* parent)
{
		return new MyQSvgWidget(static_cast<QWidget*>(parent));
}

void* QSvgWidget_NewQSvgWidget2(struct QtSvg_PackedString file, void* parent)
{
		return new MyQSvgWidget(QString::fromUtf8(file.data, file.len), static_cast<QWidget*>(parent));
}

void QSvgWidget_Load2(void* ptr, void* contents)
{
	QMetaObject::invokeMethod(static_cast<QSvgWidget*>(ptr), "load", Q_ARG(const QByteArray, *static_cast<QByteArray*>(contents)));
}

void QSvgWidget_Load2Default(void* ptr, void* contents)
{
		static_cast<QSvgWidget*>(ptr)->QSvgWidget::load(*static_cast<QByteArray*>(contents));
}

void QSvgWidget_Load(void* ptr, struct QtSvg_PackedString file)
{
	QMetaObject::invokeMethod(static_cast<QSvgWidget*>(ptr), "load", Q_ARG(const QString, QString::fromUtf8(file.data, file.len)));
}

void QSvgWidget_LoadDefault(void* ptr, struct QtSvg_PackedString file)
{
		static_cast<QSvgWidget*>(ptr)->QSvgWidget::load(QString::fromUtf8(file.data, file.len));
}

void QSvgWidget_PaintEventDefault(void* ptr, void* event)
{
		static_cast<QSvgWidget*>(ptr)->QSvgWidget::paintEvent(static_cast<QPaintEvent*>(event));
}

void QSvgWidget_DestroyQSvgWidget(void* ptr)
{
	static_cast<QSvgWidget*>(ptr)->~QSvgWidget();
}

void* QSvgWidget_SizeHintDefault(void* ptr)
{
		return ({ QSize tmpValue = static_cast<QSvgWidget*>(ptr)->QSvgWidget::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QSvgWidget_Renderer(void* ptr)
{
	return static_cast<QSvgWidget*>(ptr)->renderer();
}

void* QSvgWidget___addActions_actions_atList(void* ptr, int i)
{
	return const_cast<QAction*>(static_cast<QList<QAction *>*>(ptr)->at(i));
}

void QSvgWidget___addActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QSvgWidget___addActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>;
}

void* QSvgWidget___insertActions_actions_atList(void* ptr, int i)
{
	return const_cast<QAction*>(static_cast<QList<QAction *>*>(ptr)->at(i));
}

void QSvgWidget___insertActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QSvgWidget___insertActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>;
}

void* QSvgWidget___actions_atList(void* ptr, int i)
{
	return const_cast<QAction*>(static_cast<QList<QAction *>*>(ptr)->at(i));
}

void QSvgWidget___actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QSvgWidget___actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>;
}

void* QSvgWidget___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QSvgWidget___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QSvgWidget___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QSvgWidget___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QSvgWidget___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSvgWidget___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QSvgWidget___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QSvgWidget___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSvgWidget___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QSvgWidget___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QSvgWidget___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSvgWidget___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QSvgWidget___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QSvgWidget___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSvgWidget___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

char QSvgWidget_CloseDefault(void* ptr)
{
		return static_cast<QSvgWidget*>(ptr)->QSvgWidget::close();
}

char QSvgWidget_EventDefault(void* ptr, void* event)
{
		return static_cast<QSvgWidget*>(ptr)->QSvgWidget::event(static_cast<QEvent*>(event));
}

char QSvgWidget_FocusNextPrevChildDefault(void* ptr, char next)
{
		return static_cast<QSvgWidget*>(ptr)->QSvgWidget::focusNextPrevChild(next != 0);
}

char QSvgWidget_NativeEventDefault(void* ptr, void* eventType, void* message, long result)
{
		return static_cast<QSvgWidget*>(ptr)->QSvgWidget::nativeEvent(*static_cast<QByteArray*>(eventType), message, &result);
}

void QSvgWidget_ActionEventDefault(void* ptr, void* event)
{
		static_cast<QSvgWidget*>(ptr)->QSvgWidget::actionEvent(static_cast<QActionEvent*>(event));
}

void QSvgWidget_ChangeEventDefault(void* ptr, void* event)
{
		static_cast<QSvgWidget*>(ptr)->QSvgWidget::changeEvent(static_cast<QEvent*>(event));
}

void QSvgWidget_CloseEventDefault(void* ptr, void* event)
{
		static_cast<QSvgWidget*>(ptr)->QSvgWidget::closeEvent(static_cast<QCloseEvent*>(event));
}

void QSvgWidget_ContextMenuEventDefault(void* ptr, void* event)
{
		static_cast<QSvgWidget*>(ptr)->QSvgWidget::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
}

void QSvgWidget_DragEnterEventDefault(void* ptr, void* event)
{
		static_cast<QSvgWidget*>(ptr)->QSvgWidget::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QSvgWidget_DragLeaveEventDefault(void* ptr, void* event)
{
		static_cast<QSvgWidget*>(ptr)->QSvgWidget::dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QSvgWidget_DragMoveEventDefault(void* ptr, void* event)
{
		static_cast<QSvgWidget*>(ptr)->QSvgWidget::dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QSvgWidget_DropEventDefault(void* ptr, void* event)
{
		static_cast<QSvgWidget*>(ptr)->QSvgWidget::dropEvent(static_cast<QDropEvent*>(event));
}

void QSvgWidget_EnterEventDefault(void* ptr, void* event)
{
		static_cast<QSvgWidget*>(ptr)->QSvgWidget::enterEvent(static_cast<QEvent*>(event));
}

void QSvgWidget_FocusInEventDefault(void* ptr, void* event)
{
		static_cast<QSvgWidget*>(ptr)->QSvgWidget::focusInEvent(static_cast<QFocusEvent*>(event));
}

void QSvgWidget_FocusOutEventDefault(void* ptr, void* event)
{
		static_cast<QSvgWidget*>(ptr)->QSvgWidget::focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QSvgWidget_HideDefault(void* ptr)
{
		static_cast<QSvgWidget*>(ptr)->QSvgWidget::hide();
}

void QSvgWidget_HideEventDefault(void* ptr, void* event)
{
		static_cast<QSvgWidget*>(ptr)->QSvgWidget::hideEvent(static_cast<QHideEvent*>(event));
}

void QSvgWidget_InputMethodEventDefault(void* ptr, void* event)
{
		static_cast<QSvgWidget*>(ptr)->QSvgWidget::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void QSvgWidget_KeyPressEventDefault(void* ptr, void* event)
{
		static_cast<QSvgWidget*>(ptr)->QSvgWidget::keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QSvgWidget_KeyReleaseEventDefault(void* ptr, void* event)
{
		static_cast<QSvgWidget*>(ptr)->QSvgWidget::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QSvgWidget_LeaveEventDefault(void* ptr, void* event)
{
		static_cast<QSvgWidget*>(ptr)->QSvgWidget::leaveEvent(static_cast<QEvent*>(event));
}

void QSvgWidget_LowerDefault(void* ptr)
{
		static_cast<QSvgWidget*>(ptr)->QSvgWidget::lower();
}

void QSvgWidget_MouseDoubleClickEventDefault(void* ptr, void* event)
{
		static_cast<QSvgWidget*>(ptr)->QSvgWidget::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QSvgWidget_MouseMoveEventDefault(void* ptr, void* event)
{
		static_cast<QSvgWidget*>(ptr)->QSvgWidget::mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QSvgWidget_MousePressEventDefault(void* ptr, void* event)
{
		static_cast<QSvgWidget*>(ptr)->QSvgWidget::mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QSvgWidget_MouseReleaseEventDefault(void* ptr, void* event)
{
		static_cast<QSvgWidget*>(ptr)->QSvgWidget::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void QSvgWidget_MoveEventDefault(void* ptr, void* event)
{
		static_cast<QSvgWidget*>(ptr)->QSvgWidget::moveEvent(static_cast<QMoveEvent*>(event));
}

void QSvgWidget_RaiseDefault(void* ptr)
{
		static_cast<QSvgWidget*>(ptr)->QSvgWidget::raise();
}

void QSvgWidget_RepaintDefault(void* ptr)
{
		static_cast<QSvgWidget*>(ptr)->QSvgWidget::repaint();
}

void QSvgWidget_ResizeEventDefault(void* ptr, void* event)
{
		static_cast<QSvgWidget*>(ptr)->QSvgWidget::resizeEvent(static_cast<QResizeEvent*>(event));
}

void QSvgWidget_SetDisabledDefault(void* ptr, char disable)
{
		static_cast<QSvgWidget*>(ptr)->QSvgWidget::setDisabled(disable != 0);
}

void QSvgWidget_SetEnabledDefault(void* ptr, char vbo)
{
		static_cast<QSvgWidget*>(ptr)->QSvgWidget::setEnabled(vbo != 0);
}

void QSvgWidget_SetFocus2Default(void* ptr)
{
		static_cast<QSvgWidget*>(ptr)->QSvgWidget::setFocus();
}

void QSvgWidget_SetHiddenDefault(void* ptr, char hidden)
{
		static_cast<QSvgWidget*>(ptr)->QSvgWidget::setHidden(hidden != 0);
}

void QSvgWidget_SetStyleSheetDefault(void* ptr, struct QtSvg_PackedString styleSheet)
{
		static_cast<QSvgWidget*>(ptr)->QSvgWidget::setStyleSheet(QString::fromUtf8(styleSheet.data, styleSheet.len));
}

void QSvgWidget_SetVisibleDefault(void* ptr, char visible)
{
		static_cast<QSvgWidget*>(ptr)->QSvgWidget::setVisible(visible != 0);
}

void QSvgWidget_SetWindowModifiedDefault(void* ptr, char vbo)
{
		static_cast<QSvgWidget*>(ptr)->QSvgWidget::setWindowModified(vbo != 0);
}

void QSvgWidget_SetWindowTitleDefault(void* ptr, struct QtSvg_PackedString vqs)
{
		static_cast<QSvgWidget*>(ptr)->QSvgWidget::setWindowTitle(QString::fromUtf8(vqs.data, vqs.len));
}

void QSvgWidget_ShowDefault(void* ptr)
{
		static_cast<QSvgWidget*>(ptr)->QSvgWidget::show();
}

void QSvgWidget_ShowEventDefault(void* ptr, void* event)
{
		static_cast<QSvgWidget*>(ptr)->QSvgWidget::showEvent(static_cast<QShowEvent*>(event));
}

void QSvgWidget_ShowFullScreenDefault(void* ptr)
{
		static_cast<QSvgWidget*>(ptr)->QSvgWidget::showFullScreen();
}

void QSvgWidget_ShowMaximizedDefault(void* ptr)
{
		static_cast<QSvgWidget*>(ptr)->QSvgWidget::showMaximized();
}

void QSvgWidget_ShowMinimizedDefault(void* ptr)
{
		static_cast<QSvgWidget*>(ptr)->QSvgWidget::showMinimized();
}

void QSvgWidget_ShowNormalDefault(void* ptr)
{
		static_cast<QSvgWidget*>(ptr)->QSvgWidget::showNormal();
}

void QSvgWidget_TabletEventDefault(void* ptr, void* event)
{
		static_cast<QSvgWidget*>(ptr)->QSvgWidget::tabletEvent(static_cast<QTabletEvent*>(event));
}

void QSvgWidget_UpdateDefault(void* ptr)
{
		static_cast<QSvgWidget*>(ptr)->QSvgWidget::update();
}

void QSvgWidget_UpdateMicroFocusDefault(void* ptr)
{
		static_cast<QSvgWidget*>(ptr)->QSvgWidget::updateMicroFocus();
}

void QSvgWidget_WheelEventDefault(void* ptr, void* event)
{
		static_cast<QSvgWidget*>(ptr)->QSvgWidget::wheelEvent(static_cast<QWheelEvent*>(event));
}

void* QSvgWidget_PaintEngineDefault(void* ptr)
{
		return static_cast<QSvgWidget*>(ptr)->QSvgWidget::paintEngine();
}

void* QSvgWidget_MinimumSizeHintDefault(void* ptr)
{
		return ({ QSize tmpValue = static_cast<QSvgWidget*>(ptr)->QSvgWidget::minimumSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QSvgWidget_InputMethodQueryDefault(void* ptr, long long query)
{
		return new QVariant(static_cast<QSvgWidget*>(ptr)->QSvgWidget::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

char QSvgWidget_HasHeightForWidthDefault(void* ptr)
{
		return static_cast<QSvgWidget*>(ptr)->QSvgWidget::hasHeightForWidth();
}

int QSvgWidget_HeightForWidthDefault(void* ptr, int w)
{
		return static_cast<QSvgWidget*>(ptr)->QSvgWidget::heightForWidth(w);
}

int QSvgWidget_MetricDefault(void* ptr, long long m)
{
		return static_cast<QSvgWidget*>(ptr)->QSvgWidget::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
}

char QSvgWidget_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QSvgWidget*>(ptr)->QSvgWidget::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QSvgWidget_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QSvgWidget*>(ptr)->QSvgWidget::childEvent(static_cast<QChildEvent*>(event));
}

void QSvgWidget_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QSvgWidget*>(ptr)->QSvgWidget::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSvgWidget_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QSvgWidget*>(ptr)->QSvgWidget::customEvent(static_cast<QEvent*>(event));
}

void QSvgWidget_DeleteLaterDefault(void* ptr)
{
		static_cast<QSvgWidget*>(ptr)->QSvgWidget::deleteLater();
}

void QSvgWidget_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QSvgWidget*>(ptr)->QSvgWidget::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSvgWidget_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QSvgWidget*>(ptr)->QSvgWidget::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QSvgWidget_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QSvgWidget*>(ptr)->QSvgWidget::metaObject());
}

