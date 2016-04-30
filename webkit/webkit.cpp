#define protected public
#define private public

#include "webkit.h"
#include "_cgo_export.h"

#include <QAction>
#include <QActionEvent>
#include <QByteArray>
#include <QChildEvent>
#include <QCloseEvent>
#include <QContextMenuEvent>
#include <QDate>
#include <QDateTime>
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
#include <QGraphicsSceneMoveEvent>
#include <QGraphicsSceneResizeEvent>
#include <QGraphicsSceneWheelEvent>
#include <QGraphicsWebView>
#include <QHideEvent>
#include <QIcon>
#include <QInputMethod>
#include <QInputMethodEvent>
#include <QKeyEvent>
#include <QMetaMethod>
#include <QMetaObject>
#include <QMouseEvent>
#include <QMoveEvent>
#include <QNetworkAccessManager>
#include <QNetworkReply>
#include <QNetworkRequest>
#include <QObject>
#include <QPaintDevice>
#include <QPaintEngine>
#include <QPaintEvent>
#include <QPainter>
#include <QPainterPath>
#include <QPalette>
#include <QPixmap>
#include <QPoint>
#include <QPointF>
#include <QPrinter>
#include <QRect>
#include <QRectF>
#include <QRegion>
#include <QResizeEvent>
#include <QShowEvent>
#include <QSize>
#include <QSizeF>
#include <QString>
#include <QStringList>
#include <QStyle>
#include <QStyleOption>
#include <QStyleOptionGraphicsItem>
#include <QTabletEvent>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>
#include <QUrl>
#include <QVariant>
#include <QWebDatabase>
#include <QWebElement>
#include <QWebElementCollection>
#include <QWebFrame>
#include <QWebHistory>
#include <QWebHistoryInterface>
#include <QWebHistoryItem>
#include <QWebHitTestResult>
#include <QWebInspector>
#include <QWebPage>
#include <QWebPluginFactory>
#include <QWebSecurityOrigin>
#include <QWebSettings>
#include <QWebView>
#include <QWheelEvent>
#include <QWidget>

class MyQGraphicsWebView: public QGraphicsWebView
{
public:
	MyQGraphicsWebView(QGraphicsItem *parent) : QGraphicsWebView(parent) {};
	void back() { callbackQGraphicsWebView_Back(this, this->objectName().toUtf8().data()); };
	void contextMenuEvent(QGraphicsSceneContextMenuEvent * ev) { callbackQGraphicsWebView_ContextMenuEvent(this, this->objectName().toUtf8().data(), ev); };
	void dragEnterEvent(QGraphicsSceneDragDropEvent * ev) { callbackQGraphicsWebView_DragEnterEvent(this, this->objectName().toUtf8().data(), ev); };
	void dragLeaveEvent(QGraphicsSceneDragDropEvent * ev) { callbackQGraphicsWebView_DragLeaveEvent(this, this->objectName().toUtf8().data(), ev); };
	void dragMoveEvent(QGraphicsSceneDragDropEvent * ev) { callbackQGraphicsWebView_DragMoveEvent(this, this->objectName().toUtf8().data(), ev); };
	void dropEvent(QGraphicsSceneDragDropEvent * ev) { callbackQGraphicsWebView_DropEvent(this, this->objectName().toUtf8().data(), ev); };
	bool event(QEvent * event) { return callbackQGraphicsWebView_Event(this, this->objectName().toUtf8().data(), event) != 0; };
	void focusInEvent(QFocusEvent * ev) { callbackQGraphicsWebView_FocusInEvent(this, this->objectName().toUtf8().data(), ev); };
	bool focusNextPrevChild(bool next) { return callbackQGraphicsWebView_FocusNextPrevChild(this, this->objectName().toUtf8().data(), next) != 0; };
	void focusOutEvent(QFocusEvent * ev) { callbackQGraphicsWebView_FocusOutEvent(this, this->objectName().toUtf8().data(), ev); };
	void forward() { callbackQGraphicsWebView_Forward(this, this->objectName().toUtf8().data()); };
	void hoverLeaveEvent(QGraphicsSceneHoverEvent * ev) { callbackQGraphicsWebView_HoverLeaveEvent(this, this->objectName().toUtf8().data(), ev); };
	void hoverMoveEvent(QGraphicsSceneHoverEvent * ev) { callbackQGraphicsWebView_HoverMoveEvent(this, this->objectName().toUtf8().data(), ev); };
	void Signal_IconChanged() { callbackQGraphicsWebView_IconChanged(this, this->objectName().toUtf8().data()); };
	void inputMethodEvent(QInputMethodEvent * ev) { callbackQGraphicsWebView_InputMethodEvent(this, this->objectName().toUtf8().data(), ev); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackQGraphicsWebView_InputMethodQuery(const_cast<MyQGraphicsWebView*>(this), this->objectName().toUtf8().data(), query)); };
	QVariant itemChange(QGraphicsItem::GraphicsItemChange change, const QVariant & value) { return *static_cast<QVariant*>(callbackQGraphicsWebView_ItemChange(this, this->objectName().toUtf8().data(), change, new QVariant(value))); };
	void keyPressEvent(QKeyEvent * ev) { callbackQGraphicsWebView_KeyPressEvent(this, this->objectName().toUtf8().data(), ev); };
	void keyReleaseEvent(QKeyEvent * ev) { callbackQGraphicsWebView_KeyReleaseEvent(this, this->objectName().toUtf8().data(), ev); };
	void Signal_LinkClicked(const QUrl & url) { callbackQGraphicsWebView_LinkClicked(this, this->objectName().toUtf8().data(), new QUrl(url)); };
	void Signal_LoadFinished(bool ok) { callbackQGraphicsWebView_LoadFinished(this, this->objectName().toUtf8().data(), ok); };
	void Signal_LoadProgress(int progress) { callbackQGraphicsWebView_LoadProgress(this, this->objectName().toUtf8().data(), progress); };
	void Signal_LoadStarted() { callbackQGraphicsWebView_LoadStarted(this, this->objectName().toUtf8().data()); };
	void mouseDoubleClickEvent(QGraphicsSceneMouseEvent * ev) { callbackQGraphicsWebView_MouseDoubleClickEvent(this, this->objectName().toUtf8().data(), ev); };
	void mouseMoveEvent(QGraphicsSceneMouseEvent * ev) { callbackQGraphicsWebView_MouseMoveEvent(this, this->objectName().toUtf8().data(), ev); };
	void mousePressEvent(QGraphicsSceneMouseEvent * ev) { callbackQGraphicsWebView_MousePressEvent(this, this->objectName().toUtf8().data(), ev); };
	void mouseReleaseEvent(QGraphicsSceneMouseEvent * ev) { callbackQGraphicsWebView_MouseReleaseEvent(this, this->objectName().toUtf8().data(), ev); };
	void paint(QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget) { callbackQGraphicsWebView_Paint(this, this->objectName().toUtf8().data(), painter, const_cast<QStyleOptionGraphicsItem*>(option), widget); };
	void reload() { callbackQGraphicsWebView_Reload(this, this->objectName().toUtf8().data()); };
	bool sceneEvent(QEvent * event) { return callbackQGraphicsWebView_SceneEvent(this, this->objectName().toUtf8().data(), event) != 0; };
	void setGeometry(const QRectF & rect) { callbackQGraphicsWebView_SetGeometry(this, this->objectName().toUtf8().data(), new QRectF(static_cast<QRectF>(rect).x(), static_cast<QRectF>(rect).y(), static_cast<QRectF>(rect).width(), static_cast<QRectF>(rect).height())); };
	QSizeF sizeHint(Qt::SizeHint which, const QSizeF & constraint) const { return *static_cast<QSizeF*>(callbackQGraphicsWebView_SizeHint(const_cast<MyQGraphicsWebView*>(this), this->objectName().toUtf8().data(), which, new QSizeF(static_cast<QSizeF>(constraint).width(), static_cast<QSizeF>(constraint).height()))); };
	void Signal_StatusBarMessage(const QString & text) { callbackQGraphicsWebView_StatusBarMessage(this, this->objectName().toUtf8().data(), text.toUtf8().data()); };
	void stop() { callbackQGraphicsWebView_Stop(this, this->objectName().toUtf8().data()); };
	void Signal_TitleChanged(const QString & vqs) { callbackQGraphicsWebView_TitleChanged(this, this->objectName().toUtf8().data(), vqs.toUtf8().data()); };
	void updateGeometry() { callbackQGraphicsWebView_UpdateGeometry(this, this->objectName().toUtf8().data()); };
	void Signal_UrlChanged(const QUrl & vqu) { callbackQGraphicsWebView_UrlChanged(this, this->objectName().toUtf8().data(), new QUrl(vqu)); };
	void wheelEvent(QGraphicsSceneWheelEvent * ev) { callbackQGraphicsWebView_WheelEvent(this, this->objectName().toUtf8().data(), ev); };
	QRectF boundingRect() const { return *static_cast<QRectF*>(callbackQGraphicsWebView_BoundingRect(const_cast<MyQGraphicsWebView*>(this), this->objectName().toUtf8().data())); };
	void changeEvent(QEvent * event) { callbackQGraphicsWebView_ChangeEvent(this, this->objectName().toUtf8().data(), event); };
	bool close() { return callbackQGraphicsWebView_Close(this, this->objectName().toUtf8().data()) != 0; };
	void closeEvent(QCloseEvent * event) { callbackQGraphicsWebView_CloseEvent(this, this->objectName().toUtf8().data(), event); };
	void getContentsMargins(qreal * left, qreal * top, qreal * right, qreal * bottom) const { callbackQGraphicsWebView_GetContentsMargins(const_cast<MyQGraphicsWebView*>(this), this->objectName().toUtf8().data(), *static_cast<double*>(left), *static_cast<double*>(top), *static_cast<double*>(right), *static_cast<double*>(bottom)); };
	void grabKeyboardEvent(QEvent * event) { callbackQGraphicsWebView_GrabKeyboardEvent(this, this->objectName().toUtf8().data(), event); };
	void grabMouseEvent(QEvent * event) { callbackQGraphicsWebView_GrabMouseEvent(this, this->objectName().toUtf8().data(), event); };
	void hideEvent(QHideEvent * event) { callbackQGraphicsWebView_HideEvent(this, this->objectName().toUtf8().data(), event); };
	void initStyleOption(QStyleOption * option) const { callbackQGraphicsWebView_InitStyleOption(const_cast<MyQGraphicsWebView*>(this), this->objectName().toUtf8().data(), option); };
	void moveEvent(QGraphicsSceneMoveEvent * event) { callbackQGraphicsWebView_MoveEvent(this, this->objectName().toUtf8().data(), event); };
	void paintWindowFrame(QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget) { callbackQGraphicsWebView_PaintWindowFrame(this, this->objectName().toUtf8().data(), painter, const_cast<QStyleOptionGraphicsItem*>(option), widget); };
	void polishEvent() { callbackQGraphicsWebView_PolishEvent(this, this->objectName().toUtf8().data()); };
	void resizeEvent(QGraphicsSceneResizeEvent * event) { callbackQGraphicsWebView_ResizeEvent(this, this->objectName().toUtf8().data(), event); };
	QPainterPath shape() const { return *static_cast<QPainterPath*>(callbackQGraphicsWebView_Shape(const_cast<MyQGraphicsWebView*>(this), this->objectName().toUtf8().data())); };
	void showEvent(QShowEvent * event) { callbackQGraphicsWebView_ShowEvent(this, this->objectName().toUtf8().data(), event); };
	int type() const { return callbackQGraphicsWebView_Type(const_cast<MyQGraphicsWebView*>(this), this->objectName().toUtf8().data()); };
	void ungrabKeyboardEvent(QEvent * event) { callbackQGraphicsWebView_UngrabKeyboardEvent(this, this->objectName().toUtf8().data(), event); };
	void ungrabMouseEvent(QEvent * event) { callbackQGraphicsWebView_UngrabMouseEvent(this, this->objectName().toUtf8().data(), event); };
	bool windowFrameEvent(QEvent * event) { return callbackQGraphicsWebView_WindowFrameEvent(this, this->objectName().toUtf8().data(), event) != 0; };
	Qt::WindowFrameSection windowFrameSectionAt(const QPointF & pos) const { return static_cast<Qt::WindowFrameSection>(callbackQGraphicsWebView_WindowFrameSectionAt(const_cast<MyQGraphicsWebView*>(this), this->objectName().toUtf8().data(), new QPointF(static_cast<QPointF>(pos).x(), static_cast<QPointF>(pos).y()))); };
	void updateMicroFocus() { callbackQGraphicsWebView_UpdateMicroFocus(this, this->objectName().toUtf8().data()); };
	void advance(int phase) { callbackQGraphicsWebView_Advance(this, this->objectName().toUtf8().data(), phase); };
	bool collidesWithItem(const QGraphicsItem * other, Qt::ItemSelectionMode mode) const { return callbackQGraphicsWebView_CollidesWithItem(const_cast<MyQGraphicsWebView*>(this), this->objectName().toUtf8().data(), const_cast<QGraphicsItem*>(other), mode) != 0; };
	bool collidesWithPath(const QPainterPath & path, Qt::ItemSelectionMode mode) const { return callbackQGraphicsWebView_CollidesWithPath(const_cast<MyQGraphicsWebView*>(this), this->objectName().toUtf8().data(), new QPainterPath(path), mode) != 0; };
	bool contains(const QPointF & point) const { return callbackQGraphicsWebView_Contains(const_cast<MyQGraphicsWebView*>(this), this->objectName().toUtf8().data(), new QPointF(static_cast<QPointF>(point).x(), static_cast<QPointF>(point).y())) != 0; };
	void hoverEnterEvent(QGraphicsSceneHoverEvent * event) { callbackQGraphicsWebView_HoverEnterEvent(this, this->objectName().toUtf8().data(), event); };
	bool isObscuredBy(const QGraphicsItem * item) const { return callbackQGraphicsWebView_IsObscuredBy(const_cast<MyQGraphicsWebView*>(this), this->objectName().toUtf8().data(), const_cast<QGraphicsItem*>(item)) != 0; };
	QPainterPath opaqueArea() const { return *static_cast<QPainterPath*>(callbackQGraphicsWebView_OpaqueArea(const_cast<MyQGraphicsWebView*>(this), this->objectName().toUtf8().data())); };
	bool sceneEventFilter(QGraphicsItem * watched, QEvent * event) { return callbackQGraphicsWebView_SceneEventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	void timerEvent(QTimerEvent * event) { callbackQGraphicsWebView_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQGraphicsWebView_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQGraphicsWebView_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQGraphicsWebView_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQGraphicsWebView_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQGraphicsWebView_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQGraphicsWebView_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQGraphicsWebView_MetaObject(const_cast<MyQGraphicsWebView*>(this), this->objectName().toUtf8().data())); };
};

void* QGraphicsWebView_Icon(void* ptr)
{
	return new QIcon(static_cast<QGraphicsWebView*>(ptr)->icon());
}

int QGraphicsWebView_IsModified(void* ptr)
{
	return static_cast<QGraphicsWebView*>(ptr)->isModified();
}

int QGraphicsWebView_IsTiledBackingStoreFrozen(void* ptr)
{
	return static_cast<QGraphicsWebView*>(ptr)->isTiledBackingStoreFrozen();
}

void QGraphicsWebView_Load2(void* ptr, void* request, int operation, char* body)
{
	static_cast<QGraphicsWebView*>(ptr)->load(*static_cast<QNetworkRequest*>(request), static_cast<QNetworkAccessManager::Operation>(operation), QByteArray(body));
}

int QGraphicsWebView_ResizesToContents(void* ptr)
{
	return static_cast<QGraphicsWebView*>(ptr)->resizesToContents();
}

void QGraphicsWebView_SetResizesToContents(void* ptr, int enabled)
{
	static_cast<QGraphicsWebView*>(ptr)->setResizesToContents(enabled != 0);
}

void QGraphicsWebView_SetTiledBackingStoreFrozen(void* ptr, int frozen)
{
	static_cast<QGraphicsWebView*>(ptr)->setTiledBackingStoreFrozen(frozen != 0);
}

void QGraphicsWebView_SetUrl(void* ptr, void* vqu)
{
	static_cast<QGraphicsWebView*>(ptr)->setUrl(*static_cast<QUrl*>(vqu));
}

void QGraphicsWebView_SetZoomFactor(void* ptr, double vqr)
{
	static_cast<QGraphicsWebView*>(ptr)->setZoomFactor(static_cast<double>(vqr));
}

char* QGraphicsWebView_Title(void* ptr)
{
	return static_cast<QGraphicsWebView*>(ptr)->title().toUtf8().data();
}

void* QGraphicsWebView_Url(void* ptr)
{
	return new QUrl(static_cast<QGraphicsWebView*>(ptr)->url());
}

double QGraphicsWebView_ZoomFactor(void* ptr)
{
	return static_cast<double>(static_cast<QGraphicsWebView*>(ptr)->zoomFactor());
}

void* QGraphicsWebView_NewQGraphicsWebView(void* parent)
{
	return new MyQGraphicsWebView(static_cast<QGraphicsItem*>(parent));
}

void QGraphicsWebView_Back(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QGraphicsWebView*>(ptr), "back");
}

void QGraphicsWebView_ContextMenuEvent(void* ptr, void* ev)
{
	static_cast<QGraphicsWebView*>(ptr)->contextMenuEvent(static_cast<QGraphicsSceneContextMenuEvent*>(ev));
}

void QGraphicsWebView_ContextMenuEventDefault(void* ptr, void* ev)
{
	static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::contextMenuEvent(static_cast<QGraphicsSceneContextMenuEvent*>(ev));
}

void QGraphicsWebView_DragEnterEvent(void* ptr, void* ev)
{
	static_cast<QGraphicsWebView*>(ptr)->dragEnterEvent(static_cast<QGraphicsSceneDragDropEvent*>(ev));
}

void QGraphicsWebView_DragEnterEventDefault(void* ptr, void* ev)
{
	static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::dragEnterEvent(static_cast<QGraphicsSceneDragDropEvent*>(ev));
}

void QGraphicsWebView_DragLeaveEvent(void* ptr, void* ev)
{
	static_cast<QGraphicsWebView*>(ptr)->dragLeaveEvent(static_cast<QGraphicsSceneDragDropEvent*>(ev));
}

void QGraphicsWebView_DragLeaveEventDefault(void* ptr, void* ev)
{
	static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::dragLeaveEvent(static_cast<QGraphicsSceneDragDropEvent*>(ev));
}

void QGraphicsWebView_DragMoveEvent(void* ptr, void* ev)
{
	static_cast<QGraphicsWebView*>(ptr)->dragMoveEvent(static_cast<QGraphicsSceneDragDropEvent*>(ev));
}

void QGraphicsWebView_DragMoveEventDefault(void* ptr, void* ev)
{
	static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::dragMoveEvent(static_cast<QGraphicsSceneDragDropEvent*>(ev));
}

void QGraphicsWebView_DropEvent(void* ptr, void* ev)
{
	static_cast<QGraphicsWebView*>(ptr)->dropEvent(static_cast<QGraphicsSceneDragDropEvent*>(ev));
}

void QGraphicsWebView_DropEventDefault(void* ptr, void* ev)
{
	static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::dropEvent(static_cast<QGraphicsSceneDragDropEvent*>(ev));
}

int QGraphicsWebView_Event(void* ptr, void* event)
{
	return static_cast<QGraphicsWebView*>(ptr)->event(static_cast<QEvent*>(event));
}

int QGraphicsWebView_EventDefault(void* ptr, void* event)
{
	return static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::event(static_cast<QEvent*>(event));
}

int QGraphicsWebView_FindText(void* ptr, char* subString, int options)
{
	return static_cast<QGraphicsWebView*>(ptr)->findText(QString(subString), static_cast<QWebPage::FindFlag>(options));
}

void QGraphicsWebView_FocusInEvent(void* ptr, void* ev)
{
	static_cast<QGraphicsWebView*>(ptr)->focusInEvent(static_cast<QFocusEvent*>(ev));
}

void QGraphicsWebView_FocusInEventDefault(void* ptr, void* ev)
{
	static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::focusInEvent(static_cast<QFocusEvent*>(ev));
}

int QGraphicsWebView_FocusNextPrevChild(void* ptr, int next)
{
	return static_cast<QGraphicsWebView*>(ptr)->focusNextPrevChild(next != 0);
}

int QGraphicsWebView_FocusNextPrevChildDefault(void* ptr, int next)
{
	return static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::focusNextPrevChild(next != 0);
}

void QGraphicsWebView_FocusOutEvent(void* ptr, void* ev)
{
	static_cast<QGraphicsWebView*>(ptr)->focusOutEvent(static_cast<QFocusEvent*>(ev));
}

void QGraphicsWebView_FocusOutEventDefault(void* ptr, void* ev)
{
	static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::focusOutEvent(static_cast<QFocusEvent*>(ev));
}

void QGraphicsWebView_Forward(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QGraphicsWebView*>(ptr), "forward");
}

void* QGraphicsWebView_History(void* ptr)
{
	return static_cast<QGraphicsWebView*>(ptr)->history();
}

void QGraphicsWebView_HoverLeaveEvent(void* ptr, void* ev)
{
	static_cast<QGraphicsWebView*>(ptr)->hoverLeaveEvent(static_cast<QGraphicsSceneHoverEvent*>(ev));
}

void QGraphicsWebView_HoverLeaveEventDefault(void* ptr, void* ev)
{
	static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::hoverLeaveEvent(static_cast<QGraphicsSceneHoverEvent*>(ev));
}

void QGraphicsWebView_HoverMoveEvent(void* ptr, void* ev)
{
	static_cast<QGraphicsWebView*>(ptr)->hoverMoveEvent(static_cast<QGraphicsSceneHoverEvent*>(ev));
}

void QGraphicsWebView_HoverMoveEventDefault(void* ptr, void* ev)
{
	static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::hoverMoveEvent(static_cast<QGraphicsSceneHoverEvent*>(ev));
}

void QGraphicsWebView_ConnectIconChanged(void* ptr)
{
	QObject::connect(static_cast<QGraphicsWebView*>(ptr), static_cast<void (QGraphicsWebView::*)()>(&QGraphicsWebView::iconChanged), static_cast<MyQGraphicsWebView*>(ptr), static_cast<void (MyQGraphicsWebView::*)()>(&MyQGraphicsWebView::Signal_IconChanged));
}

void QGraphicsWebView_DisconnectIconChanged(void* ptr)
{
	QObject::disconnect(static_cast<QGraphicsWebView*>(ptr), static_cast<void (QGraphicsWebView::*)()>(&QGraphicsWebView::iconChanged), static_cast<MyQGraphicsWebView*>(ptr), static_cast<void (MyQGraphicsWebView::*)()>(&MyQGraphicsWebView::Signal_IconChanged));
}

void QGraphicsWebView_IconChanged(void* ptr)
{
	static_cast<QGraphicsWebView*>(ptr)->iconChanged();
}

void QGraphicsWebView_InputMethodEvent(void* ptr, void* ev)
{
	static_cast<QGraphicsWebView*>(ptr)->inputMethodEvent(static_cast<QInputMethodEvent*>(ev));
}

void QGraphicsWebView_InputMethodEventDefault(void* ptr, void* ev)
{
	static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::inputMethodEvent(static_cast<QInputMethodEvent*>(ev));
}

void* QGraphicsWebView_InputMethodQuery(void* ptr, int query)
{
	return new QVariant(static_cast<QGraphicsWebView*>(ptr)->inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

void* QGraphicsWebView_InputMethodQueryDefault(void* ptr, int query)
{
	return new QVariant(static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

void* QGraphicsWebView_ItemChange(void* ptr, int change, void* value)
{
	return new QVariant(static_cast<QGraphicsWebView*>(ptr)->itemChange(static_cast<QGraphicsItem::GraphicsItemChange>(change), *static_cast<QVariant*>(value)));
}

void* QGraphicsWebView_ItemChangeDefault(void* ptr, int change, void* value)
{
	return new QVariant(static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::itemChange(static_cast<QGraphicsItem::GraphicsItemChange>(change), *static_cast<QVariant*>(value)));
}

void QGraphicsWebView_KeyPressEvent(void* ptr, void* ev)
{
	static_cast<QGraphicsWebView*>(ptr)->keyPressEvent(static_cast<QKeyEvent*>(ev));
}

void QGraphicsWebView_KeyPressEventDefault(void* ptr, void* ev)
{
	static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::keyPressEvent(static_cast<QKeyEvent*>(ev));
}

void QGraphicsWebView_KeyReleaseEvent(void* ptr, void* ev)
{
	static_cast<QGraphicsWebView*>(ptr)->keyReleaseEvent(static_cast<QKeyEvent*>(ev));
}

void QGraphicsWebView_KeyReleaseEventDefault(void* ptr, void* ev)
{
	static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::keyReleaseEvent(static_cast<QKeyEvent*>(ev));
}

void QGraphicsWebView_ConnectLinkClicked(void* ptr)
{
	QObject::connect(static_cast<QGraphicsWebView*>(ptr), static_cast<void (QGraphicsWebView::*)(const QUrl &)>(&QGraphicsWebView::linkClicked), static_cast<MyQGraphicsWebView*>(ptr), static_cast<void (MyQGraphicsWebView::*)(const QUrl &)>(&MyQGraphicsWebView::Signal_LinkClicked));
}

void QGraphicsWebView_DisconnectLinkClicked(void* ptr)
{
	QObject::disconnect(static_cast<QGraphicsWebView*>(ptr), static_cast<void (QGraphicsWebView::*)(const QUrl &)>(&QGraphicsWebView::linkClicked), static_cast<MyQGraphicsWebView*>(ptr), static_cast<void (MyQGraphicsWebView::*)(const QUrl &)>(&MyQGraphicsWebView::Signal_LinkClicked));
}

void QGraphicsWebView_LinkClicked(void* ptr, void* url)
{
	static_cast<QGraphicsWebView*>(ptr)->linkClicked(*static_cast<QUrl*>(url));
}

void QGraphicsWebView_Load(void* ptr, void* url)
{
	static_cast<QGraphicsWebView*>(ptr)->load(*static_cast<QUrl*>(url));
}

void QGraphicsWebView_ConnectLoadFinished(void* ptr)
{
	QObject::connect(static_cast<QGraphicsWebView*>(ptr), static_cast<void (QGraphicsWebView::*)(bool)>(&QGraphicsWebView::loadFinished), static_cast<MyQGraphicsWebView*>(ptr), static_cast<void (MyQGraphicsWebView::*)(bool)>(&MyQGraphicsWebView::Signal_LoadFinished));
}

void QGraphicsWebView_DisconnectLoadFinished(void* ptr)
{
	QObject::disconnect(static_cast<QGraphicsWebView*>(ptr), static_cast<void (QGraphicsWebView::*)(bool)>(&QGraphicsWebView::loadFinished), static_cast<MyQGraphicsWebView*>(ptr), static_cast<void (MyQGraphicsWebView::*)(bool)>(&MyQGraphicsWebView::Signal_LoadFinished));
}

void QGraphicsWebView_LoadFinished(void* ptr, int ok)
{
	static_cast<QGraphicsWebView*>(ptr)->loadFinished(ok != 0);
}

void QGraphicsWebView_ConnectLoadProgress(void* ptr)
{
	QObject::connect(static_cast<QGraphicsWebView*>(ptr), static_cast<void (QGraphicsWebView::*)(int)>(&QGraphicsWebView::loadProgress), static_cast<MyQGraphicsWebView*>(ptr), static_cast<void (MyQGraphicsWebView::*)(int)>(&MyQGraphicsWebView::Signal_LoadProgress));
}

void QGraphicsWebView_DisconnectLoadProgress(void* ptr)
{
	QObject::disconnect(static_cast<QGraphicsWebView*>(ptr), static_cast<void (QGraphicsWebView::*)(int)>(&QGraphicsWebView::loadProgress), static_cast<MyQGraphicsWebView*>(ptr), static_cast<void (MyQGraphicsWebView::*)(int)>(&MyQGraphicsWebView::Signal_LoadProgress));
}

void QGraphicsWebView_LoadProgress(void* ptr, int progress)
{
	static_cast<QGraphicsWebView*>(ptr)->loadProgress(progress);
}

void QGraphicsWebView_ConnectLoadStarted(void* ptr)
{
	QObject::connect(static_cast<QGraphicsWebView*>(ptr), static_cast<void (QGraphicsWebView::*)()>(&QGraphicsWebView::loadStarted), static_cast<MyQGraphicsWebView*>(ptr), static_cast<void (MyQGraphicsWebView::*)()>(&MyQGraphicsWebView::Signal_LoadStarted));
}

void QGraphicsWebView_DisconnectLoadStarted(void* ptr)
{
	QObject::disconnect(static_cast<QGraphicsWebView*>(ptr), static_cast<void (QGraphicsWebView::*)()>(&QGraphicsWebView::loadStarted), static_cast<MyQGraphicsWebView*>(ptr), static_cast<void (MyQGraphicsWebView::*)()>(&MyQGraphicsWebView::Signal_LoadStarted));
}

void QGraphicsWebView_LoadStarted(void* ptr)
{
	static_cast<QGraphicsWebView*>(ptr)->loadStarted();
}

void QGraphicsWebView_MouseDoubleClickEvent(void* ptr, void* ev)
{
	static_cast<QGraphicsWebView*>(ptr)->mouseDoubleClickEvent(static_cast<QGraphicsSceneMouseEvent*>(ev));
}

void QGraphicsWebView_MouseDoubleClickEventDefault(void* ptr, void* ev)
{
	static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::mouseDoubleClickEvent(static_cast<QGraphicsSceneMouseEvent*>(ev));
}

void QGraphicsWebView_MouseMoveEvent(void* ptr, void* ev)
{
	static_cast<QGraphicsWebView*>(ptr)->mouseMoveEvent(static_cast<QGraphicsSceneMouseEvent*>(ev));
}

void QGraphicsWebView_MouseMoveEventDefault(void* ptr, void* ev)
{
	static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::mouseMoveEvent(static_cast<QGraphicsSceneMouseEvent*>(ev));
}

void QGraphicsWebView_MousePressEvent(void* ptr, void* ev)
{
	static_cast<QGraphicsWebView*>(ptr)->mousePressEvent(static_cast<QGraphicsSceneMouseEvent*>(ev));
}

void QGraphicsWebView_MousePressEventDefault(void* ptr, void* ev)
{
	static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::mousePressEvent(static_cast<QGraphicsSceneMouseEvent*>(ev));
}

void QGraphicsWebView_MouseReleaseEvent(void* ptr, void* ev)
{
	static_cast<QGraphicsWebView*>(ptr)->mouseReleaseEvent(static_cast<QGraphicsSceneMouseEvent*>(ev));
}

void QGraphicsWebView_MouseReleaseEventDefault(void* ptr, void* ev)
{
	static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::mouseReleaseEvent(static_cast<QGraphicsSceneMouseEvent*>(ev));
}

void* QGraphicsWebView_Page(void* ptr)
{
	return static_cast<QGraphicsWebView*>(ptr)->page();
}

void* QGraphicsWebView_PageAction(void* ptr, int action)
{
	return static_cast<QGraphicsWebView*>(ptr)->pageAction(static_cast<QWebPage::WebAction>(action));
}

void QGraphicsWebView_Paint(void* ptr, void* painter, void* option, void* widget)
{
	static_cast<QGraphicsWebView*>(ptr)->paint(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
}

void QGraphicsWebView_PaintDefault(void* ptr, void* painter, void* option, void* widget)
{
	static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::paint(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
}

void QGraphicsWebView_Reload(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QGraphicsWebView*>(ptr), "reload");
}

int QGraphicsWebView_RenderHints(void* ptr)
{
	return static_cast<QGraphicsWebView*>(ptr)->renderHints();
}

int QGraphicsWebView_SceneEvent(void* ptr, void* event)
{
	return static_cast<QGraphicsWebView*>(ptr)->sceneEvent(static_cast<QEvent*>(event));
}

int QGraphicsWebView_SceneEventDefault(void* ptr, void* event)
{
	return static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::sceneEvent(static_cast<QEvent*>(event));
}

void QGraphicsWebView_SetContent(void* ptr, char* data, char* mimeType, void* baseUrl)
{
	static_cast<QGraphicsWebView*>(ptr)->setContent(QByteArray(data), QString(mimeType), *static_cast<QUrl*>(baseUrl));
}

void QGraphicsWebView_SetGeometry(void* ptr, void* rect)
{
	static_cast<QGraphicsWebView*>(ptr)->setGeometry(*static_cast<QRectF*>(rect));
}

void QGraphicsWebView_SetGeometryDefault(void* ptr, void* rect)
{
	static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::setGeometry(*static_cast<QRectF*>(rect));
}

void QGraphicsWebView_SetHtml(void* ptr, char* html, void* baseUrl)
{
	static_cast<QGraphicsWebView*>(ptr)->setHtml(QString(html), *static_cast<QUrl*>(baseUrl));
}

void QGraphicsWebView_SetPage(void* ptr, void* page)
{
	static_cast<QGraphicsWebView*>(ptr)->setPage(static_cast<QWebPage*>(page));
}

void QGraphicsWebView_SetRenderHint(void* ptr, int hint, int enabled)
{
	static_cast<QGraphicsWebView*>(ptr)->setRenderHint(static_cast<QPainter::RenderHint>(hint), enabled != 0);
}

void QGraphicsWebView_SetRenderHints(void* ptr, int hints)
{
	static_cast<QGraphicsWebView*>(ptr)->setRenderHints(static_cast<QPainter::RenderHint>(hints));
}

void* QGraphicsWebView_Settings(void* ptr)
{
	return static_cast<QGraphicsWebView*>(ptr)->settings();
}

void* QGraphicsWebView_SizeHint(void* ptr, int which, void* constraint)
{
	return new QSizeF(static_cast<QSizeF>(static_cast<QGraphicsWebView*>(ptr)->sizeHint(static_cast<Qt::SizeHint>(which), *static_cast<QSizeF*>(constraint))).width(), static_cast<QSizeF>(static_cast<QGraphicsWebView*>(ptr)->sizeHint(static_cast<Qt::SizeHint>(which), *static_cast<QSizeF*>(constraint))).height());
}

void* QGraphicsWebView_SizeHintDefault(void* ptr, int which, void* constraint)
{
	return new QSizeF(static_cast<QSizeF>(static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::sizeHint(static_cast<Qt::SizeHint>(which), *static_cast<QSizeF*>(constraint))).width(), static_cast<QSizeF>(static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::sizeHint(static_cast<Qt::SizeHint>(which), *static_cast<QSizeF*>(constraint))).height());
}

void QGraphicsWebView_ConnectStatusBarMessage(void* ptr)
{
	QObject::connect(static_cast<QGraphicsWebView*>(ptr), static_cast<void (QGraphicsWebView::*)(const QString &)>(&QGraphicsWebView::statusBarMessage), static_cast<MyQGraphicsWebView*>(ptr), static_cast<void (MyQGraphicsWebView::*)(const QString &)>(&MyQGraphicsWebView::Signal_StatusBarMessage));
}

void QGraphicsWebView_DisconnectStatusBarMessage(void* ptr)
{
	QObject::disconnect(static_cast<QGraphicsWebView*>(ptr), static_cast<void (QGraphicsWebView::*)(const QString &)>(&QGraphicsWebView::statusBarMessage), static_cast<MyQGraphicsWebView*>(ptr), static_cast<void (MyQGraphicsWebView::*)(const QString &)>(&MyQGraphicsWebView::Signal_StatusBarMessage));
}

void QGraphicsWebView_StatusBarMessage(void* ptr, char* text)
{
	static_cast<QGraphicsWebView*>(ptr)->statusBarMessage(QString(text));
}

void QGraphicsWebView_Stop(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QGraphicsWebView*>(ptr), "stop");
}

void QGraphicsWebView_ConnectTitleChanged(void* ptr)
{
	QObject::connect(static_cast<QGraphicsWebView*>(ptr), static_cast<void (QGraphicsWebView::*)(const QString &)>(&QGraphicsWebView::titleChanged), static_cast<MyQGraphicsWebView*>(ptr), static_cast<void (MyQGraphicsWebView::*)(const QString &)>(&MyQGraphicsWebView::Signal_TitleChanged));
}

void QGraphicsWebView_DisconnectTitleChanged(void* ptr)
{
	QObject::disconnect(static_cast<QGraphicsWebView*>(ptr), static_cast<void (QGraphicsWebView::*)(const QString &)>(&QGraphicsWebView::titleChanged), static_cast<MyQGraphicsWebView*>(ptr), static_cast<void (MyQGraphicsWebView::*)(const QString &)>(&MyQGraphicsWebView::Signal_TitleChanged));
}

void QGraphicsWebView_TitleChanged(void* ptr, char* vqs)
{
	static_cast<QGraphicsWebView*>(ptr)->titleChanged(QString(vqs));
}

void QGraphicsWebView_TriggerPageAction(void* ptr, int action, int checked)
{
	static_cast<QGraphicsWebView*>(ptr)->triggerPageAction(static_cast<QWebPage::WebAction>(action), checked != 0);
}

void QGraphicsWebView_UpdateGeometry(void* ptr)
{
	static_cast<QGraphicsWebView*>(ptr)->updateGeometry();
}

void QGraphicsWebView_UpdateGeometryDefault(void* ptr)
{
	static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::updateGeometry();
}

void QGraphicsWebView_ConnectUrlChanged(void* ptr)
{
	QObject::connect(static_cast<QGraphicsWebView*>(ptr), static_cast<void (QGraphicsWebView::*)(const QUrl &)>(&QGraphicsWebView::urlChanged), static_cast<MyQGraphicsWebView*>(ptr), static_cast<void (MyQGraphicsWebView::*)(const QUrl &)>(&MyQGraphicsWebView::Signal_UrlChanged));
}

void QGraphicsWebView_DisconnectUrlChanged(void* ptr)
{
	QObject::disconnect(static_cast<QGraphicsWebView*>(ptr), static_cast<void (QGraphicsWebView::*)(const QUrl &)>(&QGraphicsWebView::urlChanged), static_cast<MyQGraphicsWebView*>(ptr), static_cast<void (MyQGraphicsWebView::*)(const QUrl &)>(&MyQGraphicsWebView::Signal_UrlChanged));
}

void QGraphicsWebView_UrlChanged(void* ptr, void* vqu)
{
	static_cast<QGraphicsWebView*>(ptr)->urlChanged(*static_cast<QUrl*>(vqu));
}

void QGraphicsWebView_WheelEvent(void* ptr, void* ev)
{
	static_cast<QGraphicsWebView*>(ptr)->wheelEvent(static_cast<QGraphicsSceneWheelEvent*>(ev));
}

void QGraphicsWebView_WheelEventDefault(void* ptr, void* ev)
{
	static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::wheelEvent(static_cast<QGraphicsSceneWheelEvent*>(ev));
}

void QGraphicsWebView_DestroyQGraphicsWebView(void* ptr)
{
	static_cast<QGraphicsWebView*>(ptr)->~QGraphicsWebView();
}

void* QGraphicsWebView_BoundingRect(void* ptr)
{
	return new QRectF(static_cast<QRectF>(static_cast<QGraphicsWebView*>(ptr)->boundingRect()).x(), static_cast<QRectF>(static_cast<QGraphicsWebView*>(ptr)->boundingRect()).y(), static_cast<QRectF>(static_cast<QGraphicsWebView*>(ptr)->boundingRect()).width(), static_cast<QRectF>(static_cast<QGraphicsWebView*>(ptr)->boundingRect()).height());
}

void* QGraphicsWebView_BoundingRectDefault(void* ptr)
{
	return new QRectF(static_cast<QRectF>(static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::boundingRect()).x(), static_cast<QRectF>(static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::boundingRect()).y(), static_cast<QRectF>(static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::boundingRect()).width(), static_cast<QRectF>(static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::boundingRect()).height());
}

void QGraphicsWebView_ChangeEvent(void* ptr, void* event)
{
	static_cast<QGraphicsWebView*>(ptr)->changeEvent(static_cast<QEvent*>(event));
}

void QGraphicsWebView_ChangeEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::changeEvent(static_cast<QEvent*>(event));
}

int QGraphicsWebView_Close(void* ptr)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QGraphicsWebView*>(ptr), "close", Q_RETURN_ARG(bool, returnArg));
	return returnArg;
}

int QGraphicsWebView_CloseDefault(void* ptr)
{
	return static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::close();
}

void QGraphicsWebView_CloseEvent(void* ptr, void* event)
{
	static_cast<QGraphicsWebView*>(ptr)->closeEvent(static_cast<QCloseEvent*>(event));
}

void QGraphicsWebView_CloseEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::closeEvent(static_cast<QCloseEvent*>(event));
}





void QGraphicsWebView_GrabKeyboardEvent(void* ptr, void* event)
{
	static_cast<QGraphicsWebView*>(ptr)->grabKeyboardEvent(static_cast<QEvent*>(event));
}

void QGraphicsWebView_GrabKeyboardEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::grabKeyboardEvent(static_cast<QEvent*>(event));
}

void QGraphicsWebView_GrabMouseEvent(void* ptr, void* event)
{
	static_cast<QGraphicsWebView*>(ptr)->grabMouseEvent(static_cast<QEvent*>(event));
}

void QGraphicsWebView_GrabMouseEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::grabMouseEvent(static_cast<QEvent*>(event));
}

void QGraphicsWebView_HideEvent(void* ptr, void* event)
{
	static_cast<QGraphicsWebView*>(ptr)->hideEvent(static_cast<QHideEvent*>(event));
}

void QGraphicsWebView_HideEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::hideEvent(static_cast<QHideEvent*>(event));
}

void QGraphicsWebView_InitStyleOption(void* ptr, void* option)
{
	static_cast<QGraphicsWebView*>(ptr)->initStyleOption(static_cast<QStyleOption*>(option));
}

void QGraphicsWebView_InitStyleOptionDefault(void* ptr, void* option)
{
	static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::initStyleOption(static_cast<QStyleOption*>(option));
}

void QGraphicsWebView_MoveEvent(void* ptr, void* event)
{
	static_cast<QGraphicsWebView*>(ptr)->moveEvent(static_cast<QGraphicsSceneMoveEvent*>(event));
}

void QGraphicsWebView_MoveEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::moveEvent(static_cast<QGraphicsSceneMoveEvent*>(event));
}

void QGraphicsWebView_PaintWindowFrame(void* ptr, void* painter, void* option, void* widget)
{
	static_cast<QGraphicsWebView*>(ptr)->paintWindowFrame(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
}

void QGraphicsWebView_PaintWindowFrameDefault(void* ptr, void* painter, void* option, void* widget)
{
	static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::paintWindowFrame(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
}

void QGraphicsWebView_PolishEvent(void* ptr)
{
	static_cast<QGraphicsWebView*>(ptr)->polishEvent();
}

void QGraphicsWebView_PolishEventDefault(void* ptr)
{
	static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::polishEvent();
}

void QGraphicsWebView_ResizeEvent(void* ptr, void* event)
{
	static_cast<QGraphicsWebView*>(ptr)->resizeEvent(static_cast<QGraphicsSceneResizeEvent*>(event));
}

void QGraphicsWebView_ResizeEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::resizeEvent(static_cast<QGraphicsSceneResizeEvent*>(event));
}

void* QGraphicsWebView_Shape(void* ptr)
{
	return new QPainterPath(static_cast<QGraphicsWebView*>(ptr)->shape());
}

void* QGraphicsWebView_ShapeDefault(void* ptr)
{
	return new QPainterPath(static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::shape());
}

void QGraphicsWebView_ShowEvent(void* ptr, void* event)
{
	static_cast<QGraphicsWebView*>(ptr)->showEvent(static_cast<QShowEvent*>(event));
}

void QGraphicsWebView_ShowEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::showEvent(static_cast<QShowEvent*>(event));
}

int QGraphicsWebView_Type(void* ptr)
{
	return static_cast<QGraphicsWebView*>(ptr)->type();
}

int QGraphicsWebView_TypeDefault(void* ptr)
{
	return static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::type();
}

void QGraphicsWebView_UngrabKeyboardEvent(void* ptr, void* event)
{
	static_cast<QGraphicsWebView*>(ptr)->ungrabKeyboardEvent(static_cast<QEvent*>(event));
}

void QGraphicsWebView_UngrabKeyboardEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::ungrabKeyboardEvent(static_cast<QEvent*>(event));
}

void QGraphicsWebView_UngrabMouseEvent(void* ptr, void* event)
{
	static_cast<QGraphicsWebView*>(ptr)->ungrabMouseEvent(static_cast<QEvent*>(event));
}

void QGraphicsWebView_UngrabMouseEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::ungrabMouseEvent(static_cast<QEvent*>(event));
}

int QGraphicsWebView_WindowFrameEvent(void* ptr, void* event)
{
	return static_cast<QGraphicsWebView*>(ptr)->windowFrameEvent(static_cast<QEvent*>(event));
}

int QGraphicsWebView_WindowFrameEventDefault(void* ptr, void* event)
{
	return static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::windowFrameEvent(static_cast<QEvent*>(event));
}

int QGraphicsWebView_WindowFrameSectionAt(void* ptr, void* pos)
{
	return static_cast<QGraphicsWebView*>(ptr)->windowFrameSectionAt(*static_cast<QPointF*>(pos));
}

int QGraphicsWebView_WindowFrameSectionAtDefault(void* ptr, void* pos)
{
	return static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::windowFrameSectionAt(*static_cast<QPointF*>(pos));
}

void QGraphicsWebView_UpdateMicroFocus(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QGraphicsWebView*>(ptr), "updateMicroFocus");
}

void QGraphicsWebView_UpdateMicroFocusDefault(void* ptr)
{
	static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::updateMicroFocus();
}

void QGraphicsWebView_Advance(void* ptr, int phase)
{
	static_cast<QGraphicsWebView*>(ptr)->advance(phase);
}

void QGraphicsWebView_AdvanceDefault(void* ptr, int phase)
{
	static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::advance(phase);
}

int QGraphicsWebView_CollidesWithItem(void* ptr, void* other, int mode)
{
	return static_cast<QGraphicsWebView*>(ptr)->collidesWithItem(static_cast<QGraphicsItem*>(other), static_cast<Qt::ItemSelectionMode>(mode));
}

int QGraphicsWebView_CollidesWithItemDefault(void* ptr, void* other, int mode)
{
	return static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::collidesWithItem(static_cast<QGraphicsItem*>(other), static_cast<Qt::ItemSelectionMode>(mode));
}

int QGraphicsWebView_CollidesWithPath(void* ptr, void* path, int mode)
{
	return static_cast<QGraphicsWebView*>(ptr)->collidesWithPath(*static_cast<QPainterPath*>(path), static_cast<Qt::ItemSelectionMode>(mode));
}

int QGraphicsWebView_CollidesWithPathDefault(void* ptr, void* path, int mode)
{
	return static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::collidesWithPath(*static_cast<QPainterPath*>(path), static_cast<Qt::ItemSelectionMode>(mode));
}

int QGraphicsWebView_Contains(void* ptr, void* point)
{
	return static_cast<QGraphicsWebView*>(ptr)->contains(*static_cast<QPointF*>(point));
}

int QGraphicsWebView_ContainsDefault(void* ptr, void* point)
{
	return static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::contains(*static_cast<QPointF*>(point));
}

void QGraphicsWebView_HoverEnterEvent(void* ptr, void* event)
{
	static_cast<QGraphicsWebView*>(ptr)->hoverEnterEvent(static_cast<QGraphicsSceneHoverEvent*>(event));
}

void QGraphicsWebView_HoverEnterEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::hoverEnterEvent(static_cast<QGraphicsSceneHoverEvent*>(event));
}

int QGraphicsWebView_IsObscuredBy(void* ptr, void* item)
{
	return static_cast<QGraphicsWebView*>(ptr)->isObscuredBy(static_cast<QGraphicsItem*>(item));
}

int QGraphicsWebView_IsObscuredByDefault(void* ptr, void* item)
{
	return static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::isObscuredBy(static_cast<QGraphicsItem*>(item));
}

void* QGraphicsWebView_OpaqueArea(void* ptr)
{
	return new QPainterPath(static_cast<QGraphicsWebView*>(ptr)->opaqueArea());
}

void* QGraphicsWebView_OpaqueAreaDefault(void* ptr)
{
	return new QPainterPath(static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::opaqueArea());
}

int QGraphicsWebView_SceneEventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QGraphicsWebView*>(ptr)->sceneEventFilter(static_cast<QGraphicsItem*>(watched), static_cast<QEvent*>(event));
}

int QGraphicsWebView_SceneEventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::sceneEventFilter(static_cast<QGraphicsItem*>(watched), static_cast<QEvent*>(event));
}

void QGraphicsWebView_TimerEvent(void* ptr, void* event)
{
	static_cast<QGraphicsWebView*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QGraphicsWebView_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::timerEvent(static_cast<QTimerEvent*>(event));
}

void QGraphicsWebView_ChildEvent(void* ptr, void* event)
{
	static_cast<QGraphicsWebView*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QGraphicsWebView_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::childEvent(static_cast<QChildEvent*>(event));
}

void QGraphicsWebView_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QGraphicsWebView*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGraphicsWebView_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGraphicsWebView_CustomEvent(void* ptr, void* event)
{
	static_cast<QGraphicsWebView*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QGraphicsWebView_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::customEvent(static_cast<QEvent*>(event));
}

void QGraphicsWebView_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QGraphicsWebView*>(ptr), "deleteLater");
}

void QGraphicsWebView_DeleteLaterDefault(void* ptr)
{
	static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::deleteLater();
}

void QGraphicsWebView_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QGraphicsWebView*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGraphicsWebView_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QGraphicsWebView_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QGraphicsWebView*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QGraphicsWebView_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QGraphicsWebView_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QGraphicsWebView*>(ptr)->metaObject());
}

void* QGraphicsWebView_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::metaObject());
}

void* QWebDatabase_NewQWebDatabase(void* other)
{
	return new QWebDatabase(*static_cast<QWebDatabase*>(other));
}

char* QWebDatabase_DisplayName(void* ptr)
{
	return static_cast<QWebDatabase*>(ptr)->displayName().toUtf8().data();
}

long long QWebDatabase_ExpectedSize(void* ptr)
{
	return static_cast<long long>(static_cast<QWebDatabase*>(ptr)->expectedSize());
}

char* QWebDatabase_FileName(void* ptr)
{
	return static_cast<QWebDatabase*>(ptr)->fileName().toUtf8().data();
}

char* QWebDatabase_Name(void* ptr)
{
	return static_cast<QWebDatabase*>(ptr)->name().toUtf8().data();
}

void* QWebDatabase_Origin(void* ptr)
{
	return new QWebSecurityOrigin(static_cast<QWebDatabase*>(ptr)->origin());
}

void QWebDatabase_QWebDatabase_RemoveAllDatabases()
{
	QWebDatabase::removeAllDatabases();
}

void QWebDatabase_QWebDatabase_RemoveDatabase(void* db)
{
	QWebDatabase::removeDatabase(*static_cast<QWebDatabase*>(db));
}

long long QWebDatabase_Size(void* ptr)
{
	return static_cast<long long>(static_cast<QWebDatabase*>(ptr)->size());
}

void QWebDatabase_DestroyQWebDatabase(void* ptr)
{
	static_cast<QWebDatabase*>(ptr)->~QWebDatabase();
}

void* QWebElement_NewQWebElement()
{
	return new QWebElement();
}

void* QWebElement_NewQWebElement2(void* other)
{
	return new QWebElement(*static_cast<QWebElement*>(other));
}

void QWebElement_AddClass(void* ptr, char* name)
{
	static_cast<QWebElement*>(ptr)->addClass(QString(name));
}

void QWebElement_AppendInside(void* ptr, char* markup)
{
	static_cast<QWebElement*>(ptr)->appendInside(QString(markup));
}

void QWebElement_AppendInside2(void* ptr, void* element)
{
	static_cast<QWebElement*>(ptr)->appendInside(*static_cast<QWebElement*>(element));
}

void QWebElement_AppendOutside(void* ptr, char* markup)
{
	static_cast<QWebElement*>(ptr)->appendOutside(QString(markup));
}

void QWebElement_AppendOutside2(void* ptr, void* element)
{
	static_cast<QWebElement*>(ptr)->appendOutside(*static_cast<QWebElement*>(element));
}

char* QWebElement_Attribute(void* ptr, char* name, char* defaultValue)
{
	return static_cast<QWebElement*>(ptr)->attribute(QString(name), QString(defaultValue)).toUtf8().data();
}

char* QWebElement_AttributeNS(void* ptr, char* namespaceUri, char* name, char* defaultValue)
{
	return static_cast<QWebElement*>(ptr)->attributeNS(QString(namespaceUri), QString(name), QString(defaultValue)).toUtf8().data();
}

char* QWebElement_AttributeNames(void* ptr, char* namespaceUri)
{
	return static_cast<QWebElement*>(ptr)->attributeNames(QString(namespaceUri)).join("|").toUtf8().data();
}

char* QWebElement_Classes(void* ptr)
{
	return static_cast<QWebElement*>(ptr)->classes().join("|").toUtf8().data();
}

void* QWebElement_Clone(void* ptr)
{
	return new QWebElement(static_cast<QWebElement*>(ptr)->clone());
}

void* QWebElement_Document(void* ptr)
{
	return new QWebElement(static_cast<QWebElement*>(ptr)->document());
}

void QWebElement_EncloseContentsWith2(void* ptr, char* markup)
{
	static_cast<QWebElement*>(ptr)->encloseContentsWith(QString(markup));
}

void QWebElement_EncloseContentsWith(void* ptr, void* element)
{
	static_cast<QWebElement*>(ptr)->encloseContentsWith(*static_cast<QWebElement*>(element));
}

void QWebElement_EncloseWith(void* ptr, char* markup)
{
	static_cast<QWebElement*>(ptr)->encloseWith(QString(markup));
}

void QWebElement_EncloseWith2(void* ptr, void* element)
{
	static_cast<QWebElement*>(ptr)->encloseWith(*static_cast<QWebElement*>(element));
}

void* QWebElement_EvaluateJavaScript(void* ptr, char* scriptSource)
{
	return new QVariant(static_cast<QWebElement*>(ptr)->evaluateJavaScript(QString(scriptSource)));
}

void* QWebElement_FindAll(void* ptr, char* selectorQuery)
{
	return new QWebElementCollection(static_cast<QWebElement*>(ptr)->findAll(QString(selectorQuery)));
}

void* QWebElement_FindFirst(void* ptr, char* selectorQuery)
{
	return new QWebElement(static_cast<QWebElement*>(ptr)->findFirst(QString(selectorQuery)));
}

void* QWebElement_FirstChild(void* ptr)
{
	return new QWebElement(static_cast<QWebElement*>(ptr)->firstChild());
}

void* QWebElement_Geometry(void* ptr)
{
	return new QRect(static_cast<QRect>(static_cast<QWebElement*>(ptr)->geometry()).x(), static_cast<QRect>(static_cast<QWebElement*>(ptr)->geometry()).y(), static_cast<QRect>(static_cast<QWebElement*>(ptr)->geometry()).width(), static_cast<QRect>(static_cast<QWebElement*>(ptr)->geometry()).height());
}

int QWebElement_HasAttribute(void* ptr, char* name)
{
	return static_cast<QWebElement*>(ptr)->hasAttribute(QString(name));
}

int QWebElement_HasAttributeNS(void* ptr, char* namespaceUri, char* name)
{
	return static_cast<QWebElement*>(ptr)->hasAttributeNS(QString(namespaceUri), QString(name));
}

int QWebElement_HasAttributes(void* ptr)
{
	return static_cast<QWebElement*>(ptr)->hasAttributes();
}

int QWebElement_HasClass(void* ptr, char* name)
{
	return static_cast<QWebElement*>(ptr)->hasClass(QString(name));
}

int QWebElement_HasFocus(void* ptr)
{
	return static_cast<QWebElement*>(ptr)->hasFocus();
}

int QWebElement_IsNull(void* ptr)
{
	return static_cast<QWebElement*>(ptr)->isNull();
}

void* QWebElement_LastChild(void* ptr)
{
	return new QWebElement(static_cast<QWebElement*>(ptr)->lastChild());
}

char* QWebElement_LocalName(void* ptr)
{
	return static_cast<QWebElement*>(ptr)->localName().toUtf8().data();
}

char* QWebElement_NamespaceUri(void* ptr)
{
	return static_cast<QWebElement*>(ptr)->namespaceUri().toUtf8().data();
}

void* QWebElement_NextSibling(void* ptr)
{
	return new QWebElement(static_cast<QWebElement*>(ptr)->nextSibling());
}

void* QWebElement_Parent(void* ptr)
{
	return new QWebElement(static_cast<QWebElement*>(ptr)->parent());
}

char* QWebElement_Prefix(void* ptr)
{
	return static_cast<QWebElement*>(ptr)->prefix().toUtf8().data();
}

void QWebElement_PrependInside(void* ptr, char* markup)
{
	static_cast<QWebElement*>(ptr)->prependInside(QString(markup));
}

void QWebElement_PrependInside2(void* ptr, void* element)
{
	static_cast<QWebElement*>(ptr)->prependInside(*static_cast<QWebElement*>(element));
}

void QWebElement_PrependOutside(void* ptr, char* markup)
{
	static_cast<QWebElement*>(ptr)->prependOutside(QString(markup));
}

void QWebElement_PrependOutside2(void* ptr, void* element)
{
	static_cast<QWebElement*>(ptr)->prependOutside(*static_cast<QWebElement*>(element));
}

void* QWebElement_PreviousSibling(void* ptr)
{
	return new QWebElement(static_cast<QWebElement*>(ptr)->previousSibling());
}

void QWebElement_RemoveAllChildren(void* ptr)
{
	static_cast<QWebElement*>(ptr)->removeAllChildren();
}

void QWebElement_RemoveAttribute(void* ptr, char* name)
{
	static_cast<QWebElement*>(ptr)->removeAttribute(QString(name));
}

void QWebElement_RemoveAttributeNS(void* ptr, char* namespaceUri, char* name)
{
	static_cast<QWebElement*>(ptr)->removeAttributeNS(QString(namespaceUri), QString(name));
}

void QWebElement_RemoveClass(void* ptr, char* name)
{
	static_cast<QWebElement*>(ptr)->removeClass(QString(name));
}

void QWebElement_RemoveFromDocument(void* ptr)
{
	static_cast<QWebElement*>(ptr)->removeFromDocument();
}

void QWebElement_Render(void* ptr, void* painter)
{
	static_cast<QWebElement*>(ptr)->render(static_cast<QPainter*>(painter));
}

void QWebElement_Render2(void* ptr, void* painter, void* clip)
{
	static_cast<QWebElement*>(ptr)->render(static_cast<QPainter*>(painter), *static_cast<QRect*>(clip));
}

void QWebElement_Replace(void* ptr, char* markup)
{
	static_cast<QWebElement*>(ptr)->replace(QString(markup));
}

void QWebElement_Replace2(void* ptr, void* element)
{
	static_cast<QWebElement*>(ptr)->replace(*static_cast<QWebElement*>(element));
}

void QWebElement_SetAttribute(void* ptr, char* name, char* value)
{
	static_cast<QWebElement*>(ptr)->setAttribute(QString(name), QString(value));
}

void QWebElement_SetAttributeNS(void* ptr, char* namespaceUri, char* name, char* value)
{
	static_cast<QWebElement*>(ptr)->setAttributeNS(QString(namespaceUri), QString(name), QString(value));
}

void QWebElement_SetFocus(void* ptr)
{
	static_cast<QWebElement*>(ptr)->setFocus();
}

void QWebElement_SetInnerXml(void* ptr, char* markup)
{
	static_cast<QWebElement*>(ptr)->setInnerXml(QString(markup));
}

void QWebElement_SetOuterXml(void* ptr, char* markup)
{
	static_cast<QWebElement*>(ptr)->setOuterXml(QString(markup));
}

void QWebElement_SetPlainText(void* ptr, char* text)
{
	static_cast<QWebElement*>(ptr)->setPlainText(QString(text));
}

void QWebElement_SetStyleProperty(void* ptr, char* name, char* value)
{
	static_cast<QWebElement*>(ptr)->setStyleProperty(QString(name), QString(value));
}

char* QWebElement_StyleProperty(void* ptr, char* name, int strategy)
{
	return static_cast<QWebElement*>(ptr)->styleProperty(QString(name), static_cast<QWebElement::StyleResolveStrategy>(strategy)).toUtf8().data();
}

char* QWebElement_TagName(void* ptr)
{
	return static_cast<QWebElement*>(ptr)->tagName().toUtf8().data();
}

void* QWebElement_TakeFromDocument(void* ptr)
{
	return new QWebElement(static_cast<QWebElement*>(ptr)->takeFromDocument());
}

char* QWebElement_ToInnerXml(void* ptr)
{
	return static_cast<QWebElement*>(ptr)->toInnerXml().toUtf8().data();
}

char* QWebElement_ToOuterXml(void* ptr)
{
	return static_cast<QWebElement*>(ptr)->toOuterXml().toUtf8().data();
}

char* QWebElement_ToPlainText(void* ptr)
{
	return static_cast<QWebElement*>(ptr)->toPlainText().toUtf8().data();
}

void QWebElement_ToggleClass(void* ptr, char* name)
{
	static_cast<QWebElement*>(ptr)->toggleClass(QString(name));
}

void* QWebElement_WebFrame(void* ptr)
{
	return static_cast<QWebElement*>(ptr)->webFrame();
}

void QWebElement_DestroyQWebElement(void* ptr)
{
	static_cast<QWebElement*>(ptr)->~QWebElement();
}

void* QWebElementCollection_NewQWebElementCollection()
{
	return new QWebElementCollection();
}

void* QWebElementCollection_NewQWebElementCollection2(void* contextElement, char* query)
{
	return new QWebElementCollection(*static_cast<QWebElement*>(contextElement), QString(query));
}

void* QWebElementCollection_NewQWebElementCollection3(void* other)
{
	return new QWebElementCollection(*static_cast<QWebElementCollection*>(other));
}

void QWebElementCollection_Append(void* ptr, void* other)
{
	static_cast<QWebElementCollection*>(ptr)->append(*static_cast<QWebElementCollection*>(other));
}

void* QWebElementCollection_At(void* ptr, int i)
{
	return new QWebElement(static_cast<QWebElementCollection*>(ptr)->at(i));
}

int QWebElementCollection_Count(void* ptr)
{
	return static_cast<QWebElementCollection*>(ptr)->count();
}

void* QWebElementCollection_First(void* ptr)
{
	return new QWebElement(static_cast<QWebElementCollection*>(ptr)->first());
}

void* QWebElementCollection_Last(void* ptr)
{
	return new QWebElement(static_cast<QWebElementCollection*>(ptr)->last());
}

void QWebElementCollection_DestroyQWebElementCollection(void* ptr)
{
	static_cast<QWebElementCollection*>(ptr)->~QWebElementCollection();
}

class MyQWebFrame: public QWebFrame
{
public:
	void Signal_ContentsSizeChanged(const QSize & size) { callbackQWebFrame_ContentsSizeChanged(this, this->objectName().toUtf8().data(), new QSize(static_cast<QSize>(size).width(), static_cast<QSize>(size).height())); };
	QVariant evaluateJavaScript(const QString & scriptSource) { return *static_cast<QVariant*>(callbackQWebFrame_EvaluateJavaScript(this, this->objectName().toUtf8().data(), scriptSource.toUtf8().data())); };
	bool event(QEvent * e) { return callbackQWebFrame_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	void Signal_IconChanged() { callbackQWebFrame_IconChanged(this, this->objectName().toUtf8().data()); };
	void Signal_InitialLayoutCompleted() { callbackQWebFrame_InitialLayoutCompleted(this, this->objectName().toUtf8().data()); };
	void Signal_JavaScriptWindowObjectCleared() { callbackQWebFrame_JavaScriptWindowObjectCleared(this, this->objectName().toUtf8().data()); };
	void Signal_LoadFinished(bool ok) { callbackQWebFrame_LoadFinished(this, this->objectName().toUtf8().data(), ok); };
	void Signal_LoadStarted() { callbackQWebFrame_LoadStarted(this, this->objectName().toUtf8().data()); };
	void Signal_PageChanged() { callbackQWebFrame_PageChanged(this, this->objectName().toUtf8().data()); };
	void print(QPrinter * printer) const { callbackQWebFrame_Print(const_cast<MyQWebFrame*>(this), this->objectName().toUtf8().data(), printer); };
	void Signal_TitleChanged(const QString & title) { callbackQWebFrame_TitleChanged(this, this->objectName().toUtf8().data(), title.toUtf8().data()); };
	void Signal_UrlChanged(const QUrl & url) { callbackQWebFrame_UrlChanged(this, this->objectName().toUtf8().data(), new QUrl(url)); };
	void timerEvent(QTimerEvent * event) { callbackQWebFrame_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQWebFrame_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQWebFrame_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQWebFrame_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQWebFrame_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQWebFrame_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQWebFrame_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQWebFrame_MetaObject(const_cast<MyQWebFrame*>(this), this->objectName().toUtf8().data())); };
};

void QWebFrame_AddToJavaScriptWindowObject(void* ptr, char* name, void* object, int own)
{
	static_cast<QWebFrame*>(ptr)->addToJavaScriptWindowObject(QString(name), static_cast<QObject*>(object), static_cast<QWebFrame::ValueOwnership>(own));
}

void* QWebFrame_BaseUrl(void* ptr)
{
	return new QUrl(static_cast<QWebFrame*>(ptr)->baseUrl());
}

void* QWebFrame_ContentsSize(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QWebFrame*>(ptr)->contentsSize()).width(), static_cast<QSize>(static_cast<QWebFrame*>(ptr)->contentsSize()).height());
}

int QWebFrame_HasFocus(void* ptr)
{
	return static_cast<QWebFrame*>(ptr)->hasFocus();
}

void* QWebFrame_Icon(void* ptr)
{
	return new QIcon(static_cast<QWebFrame*>(ptr)->icon());
}

void* QWebFrame_RequestedUrl(void* ptr)
{
	return new QUrl(static_cast<QWebFrame*>(ptr)->requestedUrl());
}

void* QWebFrame_ScrollPosition(void* ptr)
{
	return new QPoint(static_cast<QPoint>(static_cast<QWebFrame*>(ptr)->scrollPosition()).x(), static_cast<QPoint>(static_cast<QWebFrame*>(ptr)->scrollPosition()).y());
}

void QWebFrame_SetScrollPosition(void* ptr, void* pos)
{
	static_cast<QWebFrame*>(ptr)->setScrollPosition(*static_cast<QPoint*>(pos));
}

void QWebFrame_SetUrl(void* ptr, void* url)
{
	static_cast<QWebFrame*>(ptr)->setUrl(*static_cast<QUrl*>(url));
}

void QWebFrame_SetZoomFactor(void* ptr, double factor)
{
	static_cast<QWebFrame*>(ptr)->setZoomFactor(static_cast<double>(factor));
}

char* QWebFrame_Title(void* ptr)
{
	return static_cast<QWebFrame*>(ptr)->title().toUtf8().data();
}

void* QWebFrame_Url(void* ptr)
{
	return new QUrl(static_cast<QWebFrame*>(ptr)->url());
}

double QWebFrame_ZoomFactor(void* ptr)
{
	return static_cast<double>(static_cast<QWebFrame*>(ptr)->zoomFactor());
}

void QWebFrame_ConnectContentsSizeChanged(void* ptr)
{
	QObject::connect(static_cast<QWebFrame*>(ptr), static_cast<void (QWebFrame::*)(const QSize &)>(&QWebFrame::contentsSizeChanged), static_cast<MyQWebFrame*>(ptr), static_cast<void (MyQWebFrame::*)(const QSize &)>(&MyQWebFrame::Signal_ContentsSizeChanged));
}

void QWebFrame_DisconnectContentsSizeChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWebFrame*>(ptr), static_cast<void (QWebFrame::*)(const QSize &)>(&QWebFrame::contentsSizeChanged), static_cast<MyQWebFrame*>(ptr), static_cast<void (MyQWebFrame::*)(const QSize &)>(&MyQWebFrame::Signal_ContentsSizeChanged));
}

void QWebFrame_ContentsSizeChanged(void* ptr, void* size)
{
	static_cast<QWebFrame*>(ptr)->contentsSizeChanged(*static_cast<QSize*>(size));
}

void* QWebFrame_DocumentElement(void* ptr)
{
	return new QWebElement(static_cast<QWebFrame*>(ptr)->documentElement());
}

void* QWebFrame_EvaluateJavaScript(void* ptr, char* scriptSource)
{
	QVariant returnArg;
	QMetaObject::invokeMethod(static_cast<QWebFrame*>(ptr), "evaluateJavaScript", Q_RETURN_ARG(QVariant, returnArg), Q_ARG(QString, QString(scriptSource)));
	return new QVariant(returnArg);
}

int QWebFrame_Event(void* ptr, void* e)
{
	return static_cast<QWebFrame*>(ptr)->event(static_cast<QEvent*>(e));
}

int QWebFrame_EventDefault(void* ptr, void* e)
{
	return static_cast<QWebFrame*>(ptr)->QWebFrame::event(static_cast<QEvent*>(e));
}

void* QWebFrame_FindAllElements(void* ptr, char* selectorQuery)
{
	return new QWebElementCollection(static_cast<QWebFrame*>(ptr)->findAllElements(QString(selectorQuery)));
}

void* QWebFrame_FindFirstElement(void* ptr, char* selectorQuery)
{
	return new QWebElement(static_cast<QWebFrame*>(ptr)->findFirstElement(QString(selectorQuery)));
}

char* QWebFrame_FrameName(void* ptr)
{
	return static_cast<QWebFrame*>(ptr)->frameName().toUtf8().data();
}

void* QWebFrame_Geometry(void* ptr)
{
	return new QRect(static_cast<QRect>(static_cast<QWebFrame*>(ptr)->geometry()).x(), static_cast<QRect>(static_cast<QWebFrame*>(ptr)->geometry()).y(), static_cast<QRect>(static_cast<QWebFrame*>(ptr)->geometry()).width(), static_cast<QRect>(static_cast<QWebFrame*>(ptr)->geometry()).height());
}

void* QWebFrame_HitTestContent(void* ptr, void* pos)
{
	return new QWebHitTestResult(static_cast<QWebFrame*>(ptr)->hitTestContent(*static_cast<QPoint*>(pos)));
}

void QWebFrame_ConnectIconChanged(void* ptr)
{
	QObject::connect(static_cast<QWebFrame*>(ptr), static_cast<void (QWebFrame::*)()>(&QWebFrame::iconChanged), static_cast<MyQWebFrame*>(ptr), static_cast<void (MyQWebFrame::*)()>(&MyQWebFrame::Signal_IconChanged));
}

void QWebFrame_DisconnectIconChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWebFrame*>(ptr), static_cast<void (QWebFrame::*)()>(&QWebFrame::iconChanged), static_cast<MyQWebFrame*>(ptr), static_cast<void (MyQWebFrame::*)()>(&MyQWebFrame::Signal_IconChanged));
}

void QWebFrame_IconChanged(void* ptr)
{
	static_cast<QWebFrame*>(ptr)->iconChanged();
}

void QWebFrame_ConnectInitialLayoutCompleted(void* ptr)
{
	QObject::connect(static_cast<QWebFrame*>(ptr), static_cast<void (QWebFrame::*)()>(&QWebFrame::initialLayoutCompleted), static_cast<MyQWebFrame*>(ptr), static_cast<void (MyQWebFrame::*)()>(&MyQWebFrame::Signal_InitialLayoutCompleted));
}

void QWebFrame_DisconnectInitialLayoutCompleted(void* ptr)
{
	QObject::disconnect(static_cast<QWebFrame*>(ptr), static_cast<void (QWebFrame::*)()>(&QWebFrame::initialLayoutCompleted), static_cast<MyQWebFrame*>(ptr), static_cast<void (MyQWebFrame::*)()>(&MyQWebFrame::Signal_InitialLayoutCompleted));
}

void QWebFrame_InitialLayoutCompleted(void* ptr)
{
	static_cast<QWebFrame*>(ptr)->initialLayoutCompleted();
}

void QWebFrame_ConnectJavaScriptWindowObjectCleared(void* ptr)
{
	QObject::connect(static_cast<QWebFrame*>(ptr), static_cast<void (QWebFrame::*)()>(&QWebFrame::javaScriptWindowObjectCleared), static_cast<MyQWebFrame*>(ptr), static_cast<void (MyQWebFrame::*)()>(&MyQWebFrame::Signal_JavaScriptWindowObjectCleared));
}

void QWebFrame_DisconnectJavaScriptWindowObjectCleared(void* ptr)
{
	QObject::disconnect(static_cast<QWebFrame*>(ptr), static_cast<void (QWebFrame::*)()>(&QWebFrame::javaScriptWindowObjectCleared), static_cast<MyQWebFrame*>(ptr), static_cast<void (MyQWebFrame::*)()>(&MyQWebFrame::Signal_JavaScriptWindowObjectCleared));
}

void QWebFrame_JavaScriptWindowObjectCleared(void* ptr)
{
	static_cast<QWebFrame*>(ptr)->javaScriptWindowObjectCleared();
}

void QWebFrame_Load2(void* ptr, void* req, int operation, char* body)
{
	static_cast<QWebFrame*>(ptr)->load(*static_cast<QNetworkRequest*>(req), static_cast<QNetworkAccessManager::Operation>(operation), QByteArray(body));
}

void QWebFrame_Load(void* ptr, void* url)
{
	static_cast<QWebFrame*>(ptr)->load(*static_cast<QUrl*>(url));
}

void QWebFrame_ConnectLoadFinished(void* ptr)
{
	QObject::connect(static_cast<QWebFrame*>(ptr), static_cast<void (QWebFrame::*)(bool)>(&QWebFrame::loadFinished), static_cast<MyQWebFrame*>(ptr), static_cast<void (MyQWebFrame::*)(bool)>(&MyQWebFrame::Signal_LoadFinished));
}

void QWebFrame_DisconnectLoadFinished(void* ptr)
{
	QObject::disconnect(static_cast<QWebFrame*>(ptr), static_cast<void (QWebFrame::*)(bool)>(&QWebFrame::loadFinished), static_cast<MyQWebFrame*>(ptr), static_cast<void (MyQWebFrame::*)(bool)>(&MyQWebFrame::Signal_LoadFinished));
}

void QWebFrame_LoadFinished(void* ptr, int ok)
{
	static_cast<QWebFrame*>(ptr)->loadFinished(ok != 0);
}

void QWebFrame_ConnectLoadStarted(void* ptr)
{
	QObject::connect(static_cast<QWebFrame*>(ptr), static_cast<void (QWebFrame::*)()>(&QWebFrame::loadStarted), static_cast<MyQWebFrame*>(ptr), static_cast<void (MyQWebFrame::*)()>(&MyQWebFrame::Signal_LoadStarted));
}

void QWebFrame_DisconnectLoadStarted(void* ptr)
{
	QObject::disconnect(static_cast<QWebFrame*>(ptr), static_cast<void (QWebFrame::*)()>(&QWebFrame::loadStarted), static_cast<MyQWebFrame*>(ptr), static_cast<void (MyQWebFrame::*)()>(&MyQWebFrame::Signal_LoadStarted));
}

void QWebFrame_LoadStarted(void* ptr)
{
	static_cast<QWebFrame*>(ptr)->loadStarted();
}

void* QWebFrame_Page(void* ptr)
{
	return static_cast<QWebFrame*>(ptr)->page();
}

void QWebFrame_ConnectPageChanged(void* ptr)
{
	QObject::connect(static_cast<QWebFrame*>(ptr), static_cast<void (QWebFrame::*)()>(&QWebFrame::pageChanged), static_cast<MyQWebFrame*>(ptr), static_cast<void (MyQWebFrame::*)()>(&MyQWebFrame::Signal_PageChanged));
}

void QWebFrame_DisconnectPageChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWebFrame*>(ptr), static_cast<void (QWebFrame::*)()>(&QWebFrame::pageChanged), static_cast<MyQWebFrame*>(ptr), static_cast<void (MyQWebFrame::*)()>(&MyQWebFrame::Signal_PageChanged));
}

void QWebFrame_PageChanged(void* ptr)
{
	static_cast<QWebFrame*>(ptr)->pageChanged();
}

void* QWebFrame_ParentFrame(void* ptr)
{
	return static_cast<QWebFrame*>(ptr)->parentFrame();
}

void* QWebFrame_Pos(void* ptr)
{
	return new QPoint(static_cast<QPoint>(static_cast<QWebFrame*>(ptr)->pos()).x(), static_cast<QPoint>(static_cast<QWebFrame*>(ptr)->pos()).y());
}

void QWebFrame_Print(void* ptr, void* printer)
{
	QMetaObject::invokeMethod(static_cast<QWebFrame*>(ptr), "print", Q_ARG(QPrinter*, static_cast<QPrinter*>(printer)));
}

void QWebFrame_Render2(void* ptr, void* painter, int layer, void* clip)
{
	static_cast<QWebFrame*>(ptr)->render(static_cast<QPainter*>(painter), static_cast<QWebFrame::RenderLayer>(layer), *static_cast<QRegion*>(clip));
}

void QWebFrame_Render(void* ptr, void* painter, void* clip)
{
	static_cast<QWebFrame*>(ptr)->render(static_cast<QPainter*>(painter), *static_cast<QRegion*>(clip));
}

void QWebFrame_Scroll(void* ptr, int dx, int dy)
{
	static_cast<QWebFrame*>(ptr)->scroll(dx, dy);
}

void* QWebFrame_ScrollBarGeometry(void* ptr, int orientation)
{
	return new QRect(static_cast<QRect>(static_cast<QWebFrame*>(ptr)->scrollBarGeometry(static_cast<Qt::Orientation>(orientation))).x(), static_cast<QRect>(static_cast<QWebFrame*>(ptr)->scrollBarGeometry(static_cast<Qt::Orientation>(orientation))).y(), static_cast<QRect>(static_cast<QWebFrame*>(ptr)->scrollBarGeometry(static_cast<Qt::Orientation>(orientation))).width(), static_cast<QRect>(static_cast<QWebFrame*>(ptr)->scrollBarGeometry(static_cast<Qt::Orientation>(orientation))).height());
}

int QWebFrame_ScrollBarMaximum(void* ptr, int orientation)
{
	return static_cast<QWebFrame*>(ptr)->scrollBarMaximum(static_cast<Qt::Orientation>(orientation));
}

int QWebFrame_ScrollBarMinimum(void* ptr, int orientation)
{
	return static_cast<QWebFrame*>(ptr)->scrollBarMinimum(static_cast<Qt::Orientation>(orientation));
}

int QWebFrame_ScrollBarPolicy(void* ptr, int orientation)
{
	return static_cast<QWebFrame*>(ptr)->scrollBarPolicy(static_cast<Qt::Orientation>(orientation));
}

int QWebFrame_ScrollBarValue(void* ptr, int orientation)
{
	return static_cast<QWebFrame*>(ptr)->scrollBarValue(static_cast<Qt::Orientation>(orientation));
}

void QWebFrame_ScrollToAnchor(void* ptr, char* anchor)
{
	static_cast<QWebFrame*>(ptr)->scrollToAnchor(QString(anchor));
}

void* QWebFrame_SecurityOrigin(void* ptr)
{
	return new QWebSecurityOrigin(static_cast<QWebFrame*>(ptr)->securityOrigin());
}

void QWebFrame_SetContent(void* ptr, char* data, char* mimeType, void* baseUrl)
{
	static_cast<QWebFrame*>(ptr)->setContent(QByteArray(data), QString(mimeType), *static_cast<QUrl*>(baseUrl));
}

void QWebFrame_SetFocus(void* ptr)
{
	static_cast<QWebFrame*>(ptr)->setFocus();
}

void QWebFrame_SetHtml(void* ptr, char* html, void* baseUrl)
{
	static_cast<QWebFrame*>(ptr)->setHtml(QString(html), *static_cast<QUrl*>(baseUrl));
}

void QWebFrame_SetScrollBarPolicy(void* ptr, int orientation, int policy)
{
	static_cast<QWebFrame*>(ptr)->setScrollBarPolicy(static_cast<Qt::Orientation>(orientation), static_cast<Qt::ScrollBarPolicy>(policy));
}

void QWebFrame_SetScrollBarValue(void* ptr, int orientation, int value)
{
	static_cast<QWebFrame*>(ptr)->setScrollBarValue(static_cast<Qt::Orientation>(orientation), value);
}

void QWebFrame_SetTextSizeMultiplier(void* ptr, double factor)
{
	static_cast<QWebFrame*>(ptr)->setTextSizeMultiplier(static_cast<double>(factor));
}

double QWebFrame_TextSizeMultiplier(void* ptr)
{
	return static_cast<double>(static_cast<QWebFrame*>(ptr)->textSizeMultiplier());
}

void QWebFrame_ConnectTitleChanged(void* ptr)
{
	QObject::connect(static_cast<QWebFrame*>(ptr), static_cast<void (QWebFrame::*)(const QString &)>(&QWebFrame::titleChanged), static_cast<MyQWebFrame*>(ptr), static_cast<void (MyQWebFrame::*)(const QString &)>(&MyQWebFrame::Signal_TitleChanged));
}

void QWebFrame_DisconnectTitleChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWebFrame*>(ptr), static_cast<void (QWebFrame::*)(const QString &)>(&QWebFrame::titleChanged), static_cast<MyQWebFrame*>(ptr), static_cast<void (MyQWebFrame::*)(const QString &)>(&MyQWebFrame::Signal_TitleChanged));
}

void QWebFrame_TitleChanged(void* ptr, char* title)
{
	static_cast<QWebFrame*>(ptr)->titleChanged(QString(title));
}

char* QWebFrame_ToHtml(void* ptr)
{
	return static_cast<QWebFrame*>(ptr)->toHtml().toUtf8().data();
}

char* QWebFrame_ToPlainText(void* ptr)
{
	return static_cast<QWebFrame*>(ptr)->toPlainText().toUtf8().data();
}

void QWebFrame_ConnectUrlChanged(void* ptr)
{
	QObject::connect(static_cast<QWebFrame*>(ptr), static_cast<void (QWebFrame::*)(const QUrl &)>(&QWebFrame::urlChanged), static_cast<MyQWebFrame*>(ptr), static_cast<void (MyQWebFrame::*)(const QUrl &)>(&MyQWebFrame::Signal_UrlChanged));
}

void QWebFrame_DisconnectUrlChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWebFrame*>(ptr), static_cast<void (QWebFrame::*)(const QUrl &)>(&QWebFrame::urlChanged), static_cast<MyQWebFrame*>(ptr), static_cast<void (MyQWebFrame::*)(const QUrl &)>(&MyQWebFrame::Signal_UrlChanged));
}

void QWebFrame_UrlChanged(void* ptr, void* url)
{
	static_cast<QWebFrame*>(ptr)->urlChanged(*static_cast<QUrl*>(url));
}

void QWebFrame_TimerEvent(void* ptr, void* event)
{
	static_cast<QWebFrame*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QWebFrame_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QWebFrame*>(ptr)->QWebFrame::timerEvent(static_cast<QTimerEvent*>(event));
}

void QWebFrame_ChildEvent(void* ptr, void* event)
{
	static_cast<QWebFrame*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QWebFrame_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QWebFrame*>(ptr)->QWebFrame::childEvent(static_cast<QChildEvent*>(event));
}

void QWebFrame_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QWebFrame*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWebFrame_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QWebFrame*>(ptr)->QWebFrame::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWebFrame_CustomEvent(void* ptr, void* event)
{
	static_cast<QWebFrame*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QWebFrame_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QWebFrame*>(ptr)->QWebFrame::customEvent(static_cast<QEvent*>(event));
}

void QWebFrame_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QWebFrame*>(ptr), "deleteLater");
}

void QWebFrame_DeleteLaterDefault(void* ptr)
{
	static_cast<QWebFrame*>(ptr)->QWebFrame::deleteLater();
}

void QWebFrame_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QWebFrame*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWebFrame_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QWebFrame*>(ptr)->QWebFrame::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QWebFrame_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QWebFrame*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QWebFrame_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QWebFrame*>(ptr)->QWebFrame::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QWebFrame_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QWebFrame*>(ptr)->metaObject());
}

void* QWebFrame_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QWebFrame*>(ptr)->QWebFrame::metaObject());
}

void QWebHistory_Back(void* ptr)
{
	static_cast<QWebHistory*>(ptr)->back();
}

void* QWebHistory_BackItem(void* ptr)
{
	return new QWebHistoryItem(static_cast<QWebHistory*>(ptr)->backItem());
}

int QWebHistory_CanGoBack(void* ptr)
{
	return static_cast<QWebHistory*>(ptr)->canGoBack();
}

int QWebHistory_CanGoForward(void* ptr)
{
	return static_cast<QWebHistory*>(ptr)->canGoForward();
}

void QWebHistory_Clear(void* ptr)
{
	static_cast<QWebHistory*>(ptr)->clear();
}

int QWebHistory_Count(void* ptr)
{
	return static_cast<QWebHistory*>(ptr)->count();
}

void* QWebHistory_CurrentItem(void* ptr)
{
	return new QWebHistoryItem(static_cast<QWebHistory*>(ptr)->currentItem());
}

int QWebHistory_CurrentItemIndex(void* ptr)
{
	return static_cast<QWebHistory*>(ptr)->currentItemIndex();
}

void QWebHistory_Forward(void* ptr)
{
	static_cast<QWebHistory*>(ptr)->forward();
}

void* QWebHistory_ForwardItem(void* ptr)
{
	return new QWebHistoryItem(static_cast<QWebHistory*>(ptr)->forwardItem());
}

void QWebHistory_GoToItem(void* ptr, void* item)
{
	static_cast<QWebHistory*>(ptr)->goToItem(*static_cast<QWebHistoryItem*>(item));
}

void* QWebHistory_ItemAt(void* ptr, int i)
{
	return new QWebHistoryItem(static_cast<QWebHistory*>(ptr)->itemAt(i));
}

int QWebHistory_MaximumItemCount(void* ptr)
{
	return static_cast<QWebHistory*>(ptr)->maximumItemCount();
}

void QWebHistory_SetMaximumItemCount(void* ptr, int count)
{
	static_cast<QWebHistory*>(ptr)->setMaximumItemCount(count);
}

class MyQWebHistoryInterface: public QWebHistoryInterface
{
public:
	MyQWebHistoryInterface(QObject *parent) : QWebHistoryInterface(parent) {};
	void addHistoryEntry(const QString & url) { callbackQWebHistoryInterface_AddHistoryEntry(this, this->objectName().toUtf8().data(), url.toUtf8().data()); };
	bool historyContains(const QString & url) const { return callbackQWebHistoryInterface_HistoryContains(const_cast<MyQWebHistoryInterface*>(this), this->objectName().toUtf8().data(), url.toUtf8().data()) != 0; };
	void timerEvent(QTimerEvent * event) { callbackQWebHistoryInterface_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQWebHistoryInterface_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQWebHistoryInterface_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQWebHistoryInterface_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQWebHistoryInterface_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQWebHistoryInterface_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQWebHistoryInterface_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQWebHistoryInterface_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQWebHistoryInterface_MetaObject(const_cast<MyQWebHistoryInterface*>(this), this->objectName().toUtf8().data())); };
};

void* QWebHistoryInterface_NewQWebHistoryInterface(void* parent)
{
	return new MyQWebHistoryInterface(static_cast<QObject*>(parent));
}

void QWebHistoryInterface_AddHistoryEntry(void* ptr, char* url)
{
	static_cast<QWebHistoryInterface*>(ptr)->addHistoryEntry(QString(url));
}

void* QWebHistoryInterface_QWebHistoryInterface_DefaultInterface()
{
	return QWebHistoryInterface::defaultInterface();
}

int QWebHistoryInterface_HistoryContains(void* ptr, char* url)
{
	return static_cast<QWebHistoryInterface*>(ptr)->historyContains(QString(url));
}

void QWebHistoryInterface_QWebHistoryInterface_SetDefaultInterface(void* defaultInterface)
{
	QWebHistoryInterface::setDefaultInterface(static_cast<QWebHistoryInterface*>(defaultInterface));
}

void QWebHistoryInterface_DestroyQWebHistoryInterface(void* ptr)
{
	static_cast<QWebHistoryInterface*>(ptr)->~QWebHistoryInterface();
}

void QWebHistoryInterface_TimerEvent(void* ptr, void* event)
{
	static_cast<QWebHistoryInterface*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QWebHistoryInterface_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QWebHistoryInterface*>(ptr)->QWebHistoryInterface::timerEvent(static_cast<QTimerEvent*>(event));
}

void QWebHistoryInterface_ChildEvent(void* ptr, void* event)
{
	static_cast<QWebHistoryInterface*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QWebHistoryInterface_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QWebHistoryInterface*>(ptr)->QWebHistoryInterface::childEvent(static_cast<QChildEvent*>(event));
}

void QWebHistoryInterface_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QWebHistoryInterface*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWebHistoryInterface_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QWebHistoryInterface*>(ptr)->QWebHistoryInterface::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWebHistoryInterface_CustomEvent(void* ptr, void* event)
{
	static_cast<QWebHistoryInterface*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QWebHistoryInterface_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QWebHistoryInterface*>(ptr)->QWebHistoryInterface::customEvent(static_cast<QEvent*>(event));
}

void QWebHistoryInterface_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QWebHistoryInterface*>(ptr), "deleteLater");
}

void QWebHistoryInterface_DeleteLaterDefault(void* ptr)
{
	static_cast<QWebHistoryInterface*>(ptr)->QWebHistoryInterface::deleteLater();
}

void QWebHistoryInterface_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QWebHistoryInterface*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWebHistoryInterface_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QWebHistoryInterface*>(ptr)->QWebHistoryInterface::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QWebHistoryInterface_Event(void* ptr, void* e)
{
	return static_cast<QWebHistoryInterface*>(ptr)->event(static_cast<QEvent*>(e));
}

int QWebHistoryInterface_EventDefault(void* ptr, void* e)
{
	return static_cast<QWebHistoryInterface*>(ptr)->QWebHistoryInterface::event(static_cast<QEvent*>(e));
}

int QWebHistoryInterface_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QWebHistoryInterface*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QWebHistoryInterface_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QWebHistoryInterface*>(ptr)->QWebHistoryInterface::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QWebHistoryInterface_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QWebHistoryInterface*>(ptr)->metaObject());
}

void* QWebHistoryInterface_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QWebHistoryInterface*>(ptr)->QWebHistoryInterface::metaObject());
}

void* QWebHistoryItem_NewQWebHistoryItem(void* other)
{
	return new QWebHistoryItem(*static_cast<QWebHistoryItem*>(other));
}

void* QWebHistoryItem_Icon(void* ptr)
{
	return new QIcon(static_cast<QWebHistoryItem*>(ptr)->icon());
}

int QWebHistoryItem_IsValid(void* ptr)
{
	return static_cast<QWebHistoryItem*>(ptr)->isValid();
}

void* QWebHistoryItem_LastVisited(void* ptr)
{
	return new QDateTime(static_cast<QWebHistoryItem*>(ptr)->lastVisited());
}

void* QWebHistoryItem_OriginalUrl(void* ptr)
{
	return new QUrl(static_cast<QWebHistoryItem*>(ptr)->originalUrl());
}

void QWebHistoryItem_SetUserData(void* ptr, void* userData)
{
	static_cast<QWebHistoryItem*>(ptr)->setUserData(*static_cast<QVariant*>(userData));
}

char* QWebHistoryItem_Title(void* ptr)
{
	return static_cast<QWebHistoryItem*>(ptr)->title().toUtf8().data();
}

void* QWebHistoryItem_Url(void* ptr)
{
	return new QUrl(static_cast<QWebHistoryItem*>(ptr)->url());
}

void* QWebHistoryItem_UserData(void* ptr)
{
	return new QVariant(static_cast<QWebHistoryItem*>(ptr)->userData());
}

void QWebHistoryItem_DestroyQWebHistoryItem(void* ptr)
{
	static_cast<QWebHistoryItem*>(ptr)->~QWebHistoryItem();
}

void* QWebHitTestResult_NewQWebHitTestResult()
{
	return new QWebHitTestResult();
}

void* QWebHitTestResult_NewQWebHitTestResult2(void* other)
{
	return new QWebHitTestResult(*static_cast<QWebHitTestResult*>(other));
}

char* QWebHitTestResult_AlternateText(void* ptr)
{
	return static_cast<QWebHitTestResult*>(ptr)->alternateText().toUtf8().data();
}

void* QWebHitTestResult_BoundingRect(void* ptr)
{
	return new QRect(static_cast<QRect>(static_cast<QWebHitTestResult*>(ptr)->boundingRect()).x(), static_cast<QRect>(static_cast<QWebHitTestResult*>(ptr)->boundingRect()).y(), static_cast<QRect>(static_cast<QWebHitTestResult*>(ptr)->boundingRect()).width(), static_cast<QRect>(static_cast<QWebHitTestResult*>(ptr)->boundingRect()).height());
}

void* QWebHitTestResult_Element(void* ptr)
{
	return new QWebElement(static_cast<QWebHitTestResult*>(ptr)->element());
}

void* QWebHitTestResult_EnclosingBlockElement(void* ptr)
{
	return new QWebElement(static_cast<QWebHitTestResult*>(ptr)->enclosingBlockElement());
}

void* QWebHitTestResult_Frame(void* ptr)
{
	return static_cast<QWebHitTestResult*>(ptr)->frame();
}

void* QWebHitTestResult_ImageUrl(void* ptr)
{
	return new QUrl(static_cast<QWebHitTestResult*>(ptr)->imageUrl());
}

int QWebHitTestResult_IsContentEditable(void* ptr)
{
	return static_cast<QWebHitTestResult*>(ptr)->isContentEditable();
}

int QWebHitTestResult_IsContentSelected(void* ptr)
{
	return static_cast<QWebHitTestResult*>(ptr)->isContentSelected();
}

int QWebHitTestResult_IsNull(void* ptr)
{
	return static_cast<QWebHitTestResult*>(ptr)->isNull();
}

void* QWebHitTestResult_LinkElement(void* ptr)
{
	return new QWebElement(static_cast<QWebHitTestResult*>(ptr)->linkElement());
}

void* QWebHitTestResult_LinkTargetFrame(void* ptr)
{
	return static_cast<QWebHitTestResult*>(ptr)->linkTargetFrame();
}

char* QWebHitTestResult_LinkText(void* ptr)
{
	return static_cast<QWebHitTestResult*>(ptr)->linkText().toUtf8().data();
}

char* QWebHitTestResult_LinkTitleString(void* ptr)
{
	return static_cast<QWebHitTestResult*>(ptr)->linkTitleString().toUtf8().data();
}

void* QWebHitTestResult_LinkUrl(void* ptr)
{
	return new QUrl(static_cast<QWebHitTestResult*>(ptr)->linkUrl());
}

void* QWebHitTestResult_MediaUrl(void* ptr)
{
	return new QUrl(static_cast<QWebHitTestResult*>(ptr)->mediaUrl());
}

void* QWebHitTestResult_Pixmap(void* ptr)
{
	return new QPixmap(static_cast<QWebHitTestResult*>(ptr)->pixmap());
}

void* QWebHitTestResult_Pos(void* ptr)
{
	return new QPoint(static_cast<QPoint>(static_cast<QWebHitTestResult*>(ptr)->pos()).x(), static_cast<QPoint>(static_cast<QWebHitTestResult*>(ptr)->pos()).y());
}

char* QWebHitTestResult_Title(void* ptr)
{
	return static_cast<QWebHitTestResult*>(ptr)->title().toUtf8().data();
}

void QWebHitTestResult_DestroyQWebHitTestResult(void* ptr)
{
	static_cast<QWebHitTestResult*>(ptr)->~QWebHitTestResult();
}

class MyQWebInspector: public QWebInspector
{
public:
	MyQWebInspector(QWidget *parent) : QWebInspector(parent) {};
	void closeEvent(QCloseEvent * event) { callbackQWebInspector_CloseEvent(this, this->objectName().toUtf8().data(), event); };
	bool event(QEvent * ev) { return callbackQWebInspector_Event(this, this->objectName().toUtf8().data(), ev) != 0; };
	void hideEvent(QHideEvent * event) { callbackQWebInspector_HideEvent(this, this->objectName().toUtf8().data(), event); };
	void resizeEvent(QResizeEvent * event) { callbackQWebInspector_ResizeEvent(this, this->objectName().toUtf8().data(), event); };
	void showEvent(QShowEvent * event) { callbackQWebInspector_ShowEvent(this, this->objectName().toUtf8().data(), event); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQWebInspector_SizeHint(const_cast<MyQWebInspector*>(this), this->objectName().toUtf8().data())); };
	void actionEvent(QActionEvent * event) { callbackQWebInspector_ActionEvent(this, this->objectName().toUtf8().data(), event); };
	void dragEnterEvent(QDragEnterEvent * event) { callbackQWebInspector_DragEnterEvent(this, this->objectName().toUtf8().data(), event); };
	void dragLeaveEvent(QDragLeaveEvent * event) { callbackQWebInspector_DragLeaveEvent(this, this->objectName().toUtf8().data(), event); };
	void dragMoveEvent(QDragMoveEvent * event) { callbackQWebInspector_DragMoveEvent(this, this->objectName().toUtf8().data(), event); };
	void dropEvent(QDropEvent * event) { callbackQWebInspector_DropEvent(this, this->objectName().toUtf8().data(), event); };
	void enterEvent(QEvent * event) { callbackQWebInspector_EnterEvent(this, this->objectName().toUtf8().data(), event); };
	void focusInEvent(QFocusEvent * event) { callbackQWebInspector_FocusInEvent(this, this->objectName().toUtf8().data(), event); };
	void focusOutEvent(QFocusEvent * event) { callbackQWebInspector_FocusOutEvent(this, this->objectName().toUtf8().data(), event); };
	void leaveEvent(QEvent * event) { callbackQWebInspector_LeaveEvent(this, this->objectName().toUtf8().data(), event); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackQWebInspector_Metric(const_cast<MyQWebInspector*>(this), this->objectName().toUtf8().data(), m); };
	QSize minimumSizeHint() const { return *static_cast<QSize*>(callbackQWebInspector_MinimumSizeHint(const_cast<MyQWebInspector*>(this), this->objectName().toUtf8().data())); };
	void moveEvent(QMoveEvent * event) { callbackQWebInspector_MoveEvent(this, this->objectName().toUtf8().data(), event); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQWebInspector_PaintEngine(const_cast<MyQWebInspector*>(this), this->objectName().toUtf8().data())); };
	void paintEvent(QPaintEvent * event) { callbackQWebInspector_PaintEvent(this, this->objectName().toUtf8().data(), event); };
	void setEnabled(bool vbo) { callbackQWebInspector_SetEnabled(this, this->objectName().toUtf8().data(), vbo); };
	void setStyleSheet(const QString & styleSheet) { callbackQWebInspector_SetStyleSheet(this, this->objectName().toUtf8().data(), styleSheet.toUtf8().data()); };
	void setVisible(bool visible) { callbackQWebInspector_SetVisible(this, this->objectName().toUtf8().data(), visible); };
	void setWindowModified(bool vbo) { callbackQWebInspector_SetWindowModified(this, this->objectName().toUtf8().data(), vbo); };
	void setWindowTitle(const QString & vqs) { callbackQWebInspector_SetWindowTitle(this, this->objectName().toUtf8().data(), vqs.toUtf8().data()); };
	void changeEvent(QEvent * event) { callbackQWebInspector_ChangeEvent(this, this->objectName().toUtf8().data(), event); };
	bool close() { return callbackQWebInspector_Close(this, this->objectName().toUtf8().data()) != 0; };
	void contextMenuEvent(QContextMenuEvent * event) { callbackQWebInspector_ContextMenuEvent(this, this->objectName().toUtf8().data(), event); };
	bool focusNextPrevChild(bool next) { return callbackQWebInspector_FocusNextPrevChild(this, this->objectName().toUtf8().data(), next) != 0; };
	bool hasHeightForWidth() const { return callbackQWebInspector_HasHeightForWidth(const_cast<MyQWebInspector*>(this), this->objectName().toUtf8().data()) != 0; };
	int heightForWidth(int w) const { return callbackQWebInspector_HeightForWidth(const_cast<MyQWebInspector*>(this), this->objectName().toUtf8().data(), w); };
	void hide() { callbackQWebInspector_Hide(this, this->objectName().toUtf8().data()); };
	void initPainter(QPainter * painter) const { callbackQWebInspector_InitPainter(const_cast<MyQWebInspector*>(this), this->objectName().toUtf8().data(), painter); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackQWebInspector_InputMethodEvent(this, this->objectName().toUtf8().data(), event); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackQWebInspector_InputMethodQuery(const_cast<MyQWebInspector*>(this), this->objectName().toUtf8().data(), query)); };
	void keyPressEvent(QKeyEvent * event) { callbackQWebInspector_KeyPressEvent(this, this->objectName().toUtf8().data(), event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQWebInspector_KeyReleaseEvent(this, this->objectName().toUtf8().data(), event); };
	void lower() { callbackQWebInspector_Lower(this, this->objectName().toUtf8().data()); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQWebInspector_MouseDoubleClickEvent(this, this->objectName().toUtf8().data(), event); };
	void mouseMoveEvent(QMouseEvent * event) { callbackQWebInspector_MouseMoveEvent(this, this->objectName().toUtf8().data(), event); };
	void mousePressEvent(QMouseEvent * event) { callbackQWebInspector_MousePressEvent(this, this->objectName().toUtf8().data(), event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackQWebInspector_MouseReleaseEvent(this, this->objectName().toUtf8().data(), event); };
	bool nativeEvent(const QByteArray & eventType, void * message, long * result) { return callbackQWebInspector_NativeEvent(this, this->objectName().toUtf8().data(), QString(eventType).toUtf8().data(), message, *result) != 0; };
	void raise() { callbackQWebInspector_Raise(this, this->objectName().toUtf8().data()); };
	void repaint() { callbackQWebInspector_Repaint(this, this->objectName().toUtf8().data()); };
	void setDisabled(bool disable) { callbackQWebInspector_SetDisabled(this, this->objectName().toUtf8().data(), disable); };
	void setFocus() { callbackQWebInspector_SetFocus2(this, this->objectName().toUtf8().data()); };
	void setHidden(bool hidden) { callbackQWebInspector_SetHidden(this, this->objectName().toUtf8().data(), hidden); };
	void show() { callbackQWebInspector_Show(this, this->objectName().toUtf8().data()); };
	void showFullScreen() { callbackQWebInspector_ShowFullScreen(this, this->objectName().toUtf8().data()); };
	void showMaximized() { callbackQWebInspector_ShowMaximized(this, this->objectName().toUtf8().data()); };
	void showMinimized() { callbackQWebInspector_ShowMinimized(this, this->objectName().toUtf8().data()); };
	void showNormal() { callbackQWebInspector_ShowNormal(this, this->objectName().toUtf8().data()); };
	void tabletEvent(QTabletEvent * event) { callbackQWebInspector_TabletEvent(this, this->objectName().toUtf8().data(), event); };
	void update() { callbackQWebInspector_Update(this, this->objectName().toUtf8().data()); };
	void updateMicroFocus() { callbackQWebInspector_UpdateMicroFocus(this, this->objectName().toUtf8().data()); };
	void wheelEvent(QWheelEvent * event) { callbackQWebInspector_WheelEvent(this, this->objectName().toUtf8().data(), event); };
	void timerEvent(QTimerEvent * event) { callbackQWebInspector_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQWebInspector_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQWebInspector_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQWebInspector_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQWebInspector_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQWebInspector_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQWebInspector_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQWebInspector_MetaObject(const_cast<MyQWebInspector*>(this), this->objectName().toUtf8().data())); };
};

void* QWebInspector_NewQWebInspector(void* parent)
{
	return new MyQWebInspector(static_cast<QWidget*>(parent));
}

void QWebInspector_CloseEvent(void* ptr, void* event)
{
	static_cast<QWebInspector*>(ptr)->closeEvent(static_cast<QCloseEvent*>(event));
}

void QWebInspector_CloseEventDefault(void* ptr, void* event)
{
	static_cast<QWebInspector*>(ptr)->QWebInspector::closeEvent(static_cast<QCloseEvent*>(event));
}

int QWebInspector_Event(void* ptr, void* ev)
{
	return static_cast<QWebInspector*>(ptr)->event(static_cast<QEvent*>(ev));
}

int QWebInspector_EventDefault(void* ptr, void* ev)
{
	return static_cast<QWebInspector*>(ptr)->QWebInspector::event(static_cast<QEvent*>(ev));
}

void QWebInspector_HideEvent(void* ptr, void* event)
{
	static_cast<QWebInspector*>(ptr)->hideEvent(static_cast<QHideEvent*>(event));
}

void QWebInspector_HideEventDefault(void* ptr, void* event)
{
	static_cast<QWebInspector*>(ptr)->QWebInspector::hideEvent(static_cast<QHideEvent*>(event));
}

void* QWebInspector_Page(void* ptr)
{
	return static_cast<QWebInspector*>(ptr)->page();
}

void QWebInspector_ResizeEvent(void* ptr, void* event)
{
	static_cast<QWebInspector*>(ptr)->resizeEvent(static_cast<QResizeEvent*>(event));
}

void QWebInspector_ResizeEventDefault(void* ptr, void* event)
{
	static_cast<QWebInspector*>(ptr)->QWebInspector::resizeEvent(static_cast<QResizeEvent*>(event));
}

void QWebInspector_SetPage(void* ptr, void* page)
{
	static_cast<QWebInspector*>(ptr)->setPage(static_cast<QWebPage*>(page));
}

void QWebInspector_ShowEvent(void* ptr, void* event)
{
	static_cast<QWebInspector*>(ptr)->showEvent(static_cast<QShowEvent*>(event));
}

void QWebInspector_ShowEventDefault(void* ptr, void* event)
{
	static_cast<QWebInspector*>(ptr)->QWebInspector::showEvent(static_cast<QShowEvent*>(event));
}

void* QWebInspector_SizeHint(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QWebInspector*>(ptr)->sizeHint()).width(), static_cast<QSize>(static_cast<QWebInspector*>(ptr)->sizeHint()).height());
}

void* QWebInspector_SizeHintDefault(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QWebInspector*>(ptr)->QWebInspector::sizeHint()).width(), static_cast<QSize>(static_cast<QWebInspector*>(ptr)->QWebInspector::sizeHint()).height());
}

void QWebInspector_DestroyQWebInspector(void* ptr)
{
	static_cast<QWebInspector*>(ptr)->~QWebInspector();
}

void QWebInspector_ActionEvent(void* ptr, void* event)
{
	static_cast<QWebInspector*>(ptr)->actionEvent(static_cast<QActionEvent*>(event));
}

void QWebInspector_ActionEventDefault(void* ptr, void* event)
{
	static_cast<QWebInspector*>(ptr)->QWebInspector::actionEvent(static_cast<QActionEvent*>(event));
}

void QWebInspector_DragEnterEvent(void* ptr, void* event)
{
	static_cast<QWebInspector*>(ptr)->dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QWebInspector_DragEnterEventDefault(void* ptr, void* event)
{
	static_cast<QWebInspector*>(ptr)->QWebInspector::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QWebInspector_DragLeaveEvent(void* ptr, void* event)
{
	static_cast<QWebInspector*>(ptr)->dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QWebInspector_DragLeaveEventDefault(void* ptr, void* event)
{
	static_cast<QWebInspector*>(ptr)->QWebInspector::dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QWebInspector_DragMoveEvent(void* ptr, void* event)
{
	static_cast<QWebInspector*>(ptr)->dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QWebInspector_DragMoveEventDefault(void* ptr, void* event)
{
	static_cast<QWebInspector*>(ptr)->QWebInspector::dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QWebInspector_DropEvent(void* ptr, void* event)
{
	static_cast<QWebInspector*>(ptr)->dropEvent(static_cast<QDropEvent*>(event));
}

void QWebInspector_DropEventDefault(void* ptr, void* event)
{
	static_cast<QWebInspector*>(ptr)->QWebInspector::dropEvent(static_cast<QDropEvent*>(event));
}

void QWebInspector_EnterEvent(void* ptr, void* event)
{
	static_cast<QWebInspector*>(ptr)->enterEvent(static_cast<QEvent*>(event));
}

void QWebInspector_EnterEventDefault(void* ptr, void* event)
{
	static_cast<QWebInspector*>(ptr)->QWebInspector::enterEvent(static_cast<QEvent*>(event));
}

void QWebInspector_FocusInEvent(void* ptr, void* event)
{
	static_cast<QWebInspector*>(ptr)->focusInEvent(static_cast<QFocusEvent*>(event));
}

void QWebInspector_FocusInEventDefault(void* ptr, void* event)
{
	static_cast<QWebInspector*>(ptr)->QWebInspector::focusInEvent(static_cast<QFocusEvent*>(event));
}

void QWebInspector_FocusOutEvent(void* ptr, void* event)
{
	static_cast<QWebInspector*>(ptr)->focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QWebInspector_FocusOutEventDefault(void* ptr, void* event)
{
	static_cast<QWebInspector*>(ptr)->QWebInspector::focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QWebInspector_LeaveEvent(void* ptr, void* event)
{
	static_cast<QWebInspector*>(ptr)->leaveEvent(static_cast<QEvent*>(event));
}

void QWebInspector_LeaveEventDefault(void* ptr, void* event)
{
	static_cast<QWebInspector*>(ptr)->QWebInspector::leaveEvent(static_cast<QEvent*>(event));
}

int QWebInspector_Metric(void* ptr, int m)
{
	return static_cast<QWebInspector*>(ptr)->metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
}

int QWebInspector_MetricDefault(void* ptr, int m)
{
	return static_cast<QWebInspector*>(ptr)->QWebInspector::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
}

void* QWebInspector_MinimumSizeHint(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QWebInspector*>(ptr)->minimumSizeHint()).width(), static_cast<QSize>(static_cast<QWebInspector*>(ptr)->minimumSizeHint()).height());
}

void* QWebInspector_MinimumSizeHintDefault(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QWebInspector*>(ptr)->QWebInspector::minimumSizeHint()).width(), static_cast<QSize>(static_cast<QWebInspector*>(ptr)->QWebInspector::minimumSizeHint()).height());
}

void QWebInspector_MoveEvent(void* ptr, void* event)
{
	static_cast<QWebInspector*>(ptr)->moveEvent(static_cast<QMoveEvent*>(event));
}

void QWebInspector_MoveEventDefault(void* ptr, void* event)
{
	static_cast<QWebInspector*>(ptr)->QWebInspector::moveEvent(static_cast<QMoveEvent*>(event));
}

void* QWebInspector_PaintEngine(void* ptr)
{
	return static_cast<QWebInspector*>(ptr)->paintEngine();
}

void* QWebInspector_PaintEngineDefault(void* ptr)
{
	return static_cast<QWebInspector*>(ptr)->QWebInspector::paintEngine();
}

void QWebInspector_PaintEvent(void* ptr, void* event)
{
	static_cast<QWebInspector*>(ptr)->paintEvent(static_cast<QPaintEvent*>(event));
}

void QWebInspector_PaintEventDefault(void* ptr, void* event)
{
	static_cast<QWebInspector*>(ptr)->QWebInspector::paintEvent(static_cast<QPaintEvent*>(event));
}

void QWebInspector_SetEnabled(void* ptr, int vbo)
{
	QMetaObject::invokeMethod(static_cast<QWebInspector*>(ptr), "setEnabled", Q_ARG(bool, vbo != 0));
}

void QWebInspector_SetEnabledDefault(void* ptr, int vbo)
{
	static_cast<QWebInspector*>(ptr)->QWebInspector::setEnabled(vbo != 0);
}

void QWebInspector_SetStyleSheet(void* ptr, char* styleSheet)
{
	QMetaObject::invokeMethod(static_cast<QWebInspector*>(ptr), "setStyleSheet", Q_ARG(QString, QString(styleSheet)));
}

void QWebInspector_SetStyleSheetDefault(void* ptr, char* styleSheet)
{
	static_cast<QWebInspector*>(ptr)->QWebInspector::setStyleSheet(QString(styleSheet));
}

void QWebInspector_SetVisible(void* ptr, int visible)
{
	QMetaObject::invokeMethod(static_cast<QWebInspector*>(ptr), "setVisible", Q_ARG(bool, visible != 0));
}

void QWebInspector_SetVisibleDefault(void* ptr, int visible)
{
	static_cast<QWebInspector*>(ptr)->QWebInspector::setVisible(visible != 0);
}

void QWebInspector_SetWindowModified(void* ptr, int vbo)
{
	QMetaObject::invokeMethod(static_cast<QWebInspector*>(ptr), "setWindowModified", Q_ARG(bool, vbo != 0));
}

void QWebInspector_SetWindowModifiedDefault(void* ptr, int vbo)
{
	static_cast<QWebInspector*>(ptr)->QWebInspector::setWindowModified(vbo != 0);
}

void QWebInspector_SetWindowTitle(void* ptr, char* vqs)
{
	QMetaObject::invokeMethod(static_cast<QWebInspector*>(ptr), "setWindowTitle", Q_ARG(QString, QString(vqs)));
}

void QWebInspector_SetWindowTitleDefault(void* ptr, char* vqs)
{
	static_cast<QWebInspector*>(ptr)->QWebInspector::setWindowTitle(QString(vqs));
}

void QWebInspector_ChangeEvent(void* ptr, void* event)
{
	static_cast<QWebInspector*>(ptr)->changeEvent(static_cast<QEvent*>(event));
}

void QWebInspector_ChangeEventDefault(void* ptr, void* event)
{
	static_cast<QWebInspector*>(ptr)->QWebInspector::changeEvent(static_cast<QEvent*>(event));
}

int QWebInspector_Close(void* ptr)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QWebInspector*>(ptr), "close", Q_RETURN_ARG(bool, returnArg));
	return returnArg;
}

int QWebInspector_CloseDefault(void* ptr)
{
	return static_cast<QWebInspector*>(ptr)->QWebInspector::close();
}

void QWebInspector_ContextMenuEvent(void* ptr, void* event)
{
	static_cast<QWebInspector*>(ptr)->contextMenuEvent(static_cast<QContextMenuEvent*>(event));
}

void QWebInspector_ContextMenuEventDefault(void* ptr, void* event)
{
	static_cast<QWebInspector*>(ptr)->QWebInspector::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
}

int QWebInspector_FocusNextPrevChild(void* ptr, int next)
{
	return static_cast<QWebInspector*>(ptr)->focusNextPrevChild(next != 0);
}

int QWebInspector_FocusNextPrevChildDefault(void* ptr, int next)
{
	return static_cast<QWebInspector*>(ptr)->QWebInspector::focusNextPrevChild(next != 0);
}

int QWebInspector_HasHeightForWidth(void* ptr)
{
	return static_cast<QWebInspector*>(ptr)->hasHeightForWidth();
}

int QWebInspector_HasHeightForWidthDefault(void* ptr)
{
	return static_cast<QWebInspector*>(ptr)->QWebInspector::hasHeightForWidth();
}

int QWebInspector_HeightForWidth(void* ptr, int w)
{
	return static_cast<QWebInspector*>(ptr)->heightForWidth(w);
}

int QWebInspector_HeightForWidthDefault(void* ptr, int w)
{
	return static_cast<QWebInspector*>(ptr)->QWebInspector::heightForWidth(w);
}

void QWebInspector_Hide(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QWebInspector*>(ptr), "hide");
}

void QWebInspector_HideDefault(void* ptr)
{
	static_cast<QWebInspector*>(ptr)->QWebInspector::hide();
}

void QWebInspector_InitPainter(void* ptr, void* painter)
{
	static_cast<QWebInspector*>(ptr)->initPainter(static_cast<QPainter*>(painter));
}

void QWebInspector_InitPainterDefault(void* ptr, void* painter)
{
	static_cast<QWebInspector*>(ptr)->QWebInspector::initPainter(static_cast<QPainter*>(painter));
}

void QWebInspector_InputMethodEvent(void* ptr, void* event)
{
	static_cast<QWebInspector*>(ptr)->inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void QWebInspector_InputMethodEventDefault(void* ptr, void* event)
{
	static_cast<QWebInspector*>(ptr)->QWebInspector::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void* QWebInspector_InputMethodQuery(void* ptr, int query)
{
	return new QVariant(static_cast<QWebInspector*>(ptr)->inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

void* QWebInspector_InputMethodQueryDefault(void* ptr, int query)
{
	return new QVariant(static_cast<QWebInspector*>(ptr)->QWebInspector::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

void QWebInspector_KeyPressEvent(void* ptr, void* event)
{
	static_cast<QWebInspector*>(ptr)->keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QWebInspector_KeyPressEventDefault(void* ptr, void* event)
{
	static_cast<QWebInspector*>(ptr)->QWebInspector::keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QWebInspector_KeyReleaseEvent(void* ptr, void* event)
{
	static_cast<QWebInspector*>(ptr)->keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QWebInspector_KeyReleaseEventDefault(void* ptr, void* event)
{
	static_cast<QWebInspector*>(ptr)->QWebInspector::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QWebInspector_Lower(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QWebInspector*>(ptr), "lower");
}

void QWebInspector_LowerDefault(void* ptr)
{
	static_cast<QWebInspector*>(ptr)->QWebInspector::lower();
}

void QWebInspector_MouseDoubleClickEvent(void* ptr, void* event)
{
	static_cast<QWebInspector*>(ptr)->mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QWebInspector_MouseDoubleClickEventDefault(void* ptr, void* event)
{
	static_cast<QWebInspector*>(ptr)->QWebInspector::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QWebInspector_MouseMoveEvent(void* ptr, void* event)
{
	static_cast<QWebInspector*>(ptr)->mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QWebInspector_MouseMoveEventDefault(void* ptr, void* event)
{
	static_cast<QWebInspector*>(ptr)->QWebInspector::mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QWebInspector_MousePressEvent(void* ptr, void* event)
{
	static_cast<QWebInspector*>(ptr)->mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QWebInspector_MousePressEventDefault(void* ptr, void* event)
{
	static_cast<QWebInspector*>(ptr)->QWebInspector::mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QWebInspector_MouseReleaseEvent(void* ptr, void* event)
{
	static_cast<QWebInspector*>(ptr)->mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void QWebInspector_MouseReleaseEventDefault(void* ptr, void* event)
{
	static_cast<QWebInspector*>(ptr)->QWebInspector::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

int QWebInspector_NativeEvent(void* ptr, char* eventType, void* message, long result)
{
	return static_cast<QWebInspector*>(ptr)->nativeEvent(QByteArray(eventType), message, &result);
}

int QWebInspector_NativeEventDefault(void* ptr, char* eventType, void* message, long result)
{
	return static_cast<QWebInspector*>(ptr)->QWebInspector::nativeEvent(QByteArray(eventType), message, &result);
}

void QWebInspector_Raise(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QWebInspector*>(ptr), "raise");
}

void QWebInspector_RaiseDefault(void* ptr)
{
	static_cast<QWebInspector*>(ptr)->QWebInspector::raise();
}

void QWebInspector_Repaint(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QWebInspector*>(ptr), "repaint");
}

void QWebInspector_RepaintDefault(void* ptr)
{
	static_cast<QWebInspector*>(ptr)->QWebInspector::repaint();
}

void QWebInspector_SetDisabled(void* ptr, int disable)
{
	QMetaObject::invokeMethod(static_cast<QWebInspector*>(ptr), "setDisabled", Q_ARG(bool, disable != 0));
}

void QWebInspector_SetDisabledDefault(void* ptr, int disable)
{
	static_cast<QWebInspector*>(ptr)->QWebInspector::setDisabled(disable != 0);
}

void QWebInspector_SetFocus2(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QWebInspector*>(ptr), "setFocus");
}

void QWebInspector_SetFocus2Default(void* ptr)
{
	static_cast<QWebInspector*>(ptr)->QWebInspector::setFocus();
}

void QWebInspector_SetHidden(void* ptr, int hidden)
{
	QMetaObject::invokeMethod(static_cast<QWebInspector*>(ptr), "setHidden", Q_ARG(bool, hidden != 0));
}

void QWebInspector_SetHiddenDefault(void* ptr, int hidden)
{
	static_cast<QWebInspector*>(ptr)->QWebInspector::setHidden(hidden != 0);
}

void QWebInspector_Show(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QWebInspector*>(ptr), "show");
}

void QWebInspector_ShowDefault(void* ptr)
{
	static_cast<QWebInspector*>(ptr)->QWebInspector::show();
}

void QWebInspector_ShowFullScreen(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QWebInspector*>(ptr), "showFullScreen");
}

void QWebInspector_ShowFullScreenDefault(void* ptr)
{
	static_cast<QWebInspector*>(ptr)->QWebInspector::showFullScreen();
}

void QWebInspector_ShowMaximized(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QWebInspector*>(ptr), "showMaximized");
}

void QWebInspector_ShowMaximizedDefault(void* ptr)
{
	static_cast<QWebInspector*>(ptr)->QWebInspector::showMaximized();
}

void QWebInspector_ShowMinimized(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QWebInspector*>(ptr), "showMinimized");
}

void QWebInspector_ShowMinimizedDefault(void* ptr)
{
	static_cast<QWebInspector*>(ptr)->QWebInspector::showMinimized();
}

void QWebInspector_ShowNormal(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QWebInspector*>(ptr), "showNormal");
}

void QWebInspector_ShowNormalDefault(void* ptr)
{
	static_cast<QWebInspector*>(ptr)->QWebInspector::showNormal();
}

void QWebInspector_TabletEvent(void* ptr, void* event)
{
	static_cast<QWebInspector*>(ptr)->tabletEvent(static_cast<QTabletEvent*>(event));
}

void QWebInspector_TabletEventDefault(void* ptr, void* event)
{
	static_cast<QWebInspector*>(ptr)->QWebInspector::tabletEvent(static_cast<QTabletEvent*>(event));
}

void QWebInspector_Update(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QWebInspector*>(ptr), "update");
}

void QWebInspector_UpdateDefault(void* ptr)
{
	static_cast<QWebInspector*>(ptr)->QWebInspector::update();
}

void QWebInspector_UpdateMicroFocus(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QWebInspector*>(ptr), "updateMicroFocus");
}

void QWebInspector_UpdateMicroFocusDefault(void* ptr)
{
	static_cast<QWebInspector*>(ptr)->QWebInspector::updateMicroFocus();
}

void QWebInspector_WheelEvent(void* ptr, void* event)
{
	static_cast<QWebInspector*>(ptr)->wheelEvent(static_cast<QWheelEvent*>(event));
}

void QWebInspector_WheelEventDefault(void* ptr, void* event)
{
	static_cast<QWebInspector*>(ptr)->QWebInspector::wheelEvent(static_cast<QWheelEvent*>(event));
}

void QWebInspector_TimerEvent(void* ptr, void* event)
{
	static_cast<QWebInspector*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QWebInspector_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QWebInspector*>(ptr)->QWebInspector::timerEvent(static_cast<QTimerEvent*>(event));
}

void QWebInspector_ChildEvent(void* ptr, void* event)
{
	static_cast<QWebInspector*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QWebInspector_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QWebInspector*>(ptr)->QWebInspector::childEvent(static_cast<QChildEvent*>(event));
}

void QWebInspector_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QWebInspector*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWebInspector_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QWebInspector*>(ptr)->QWebInspector::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWebInspector_CustomEvent(void* ptr, void* event)
{
	static_cast<QWebInspector*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QWebInspector_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QWebInspector*>(ptr)->QWebInspector::customEvent(static_cast<QEvent*>(event));
}

void QWebInspector_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QWebInspector*>(ptr), "deleteLater");
}

void QWebInspector_DeleteLaterDefault(void* ptr)
{
	static_cast<QWebInspector*>(ptr)->QWebInspector::deleteLater();
}

void QWebInspector_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QWebInspector*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWebInspector_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QWebInspector*>(ptr)->QWebInspector::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QWebInspector_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QWebInspector*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QWebInspector_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QWebInspector*>(ptr)->QWebInspector::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QWebInspector_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QWebInspector*>(ptr)->metaObject());
}

void* QWebInspector_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QWebInspector*>(ptr)->QWebInspector::metaObject());
}

class MyQWebPage: public QWebPage
{
public:
	MyQWebPage(QObject *parent) : QWebPage(parent) {};
	bool acceptNavigationRequest(QWebFrame * frame, const QNetworkRequest & request, QWebPage::NavigationType ty) { return callbackQWebPage_AcceptNavigationRequest(this, this->objectName().toUtf8().data(), frame, new QNetworkRequest(request), ty) != 0; };
	bool shouldInterruptJavaScript() { return callbackQWebPage_ShouldInterruptJavaScript(this, this->objectName().toUtf8().data()) != 0; };
	QString chooseFile(QWebFrame * parentFrame, const QString & suggestedFile) { return QString(callbackQWebPage_ChooseFile(this, this->objectName().toUtf8().data(), parentFrame, suggestedFile.toUtf8().data())); };
	void Signal_ContentsChanged() { callbackQWebPage_ContentsChanged(this, this->objectName().toUtf8().data()); };
	QObject * createPlugin(const QString & classid, const QUrl & url, const QStringList & paramNames, const QStringList & paramValues) { return static_cast<QObject*>(callbackQWebPage_CreatePlugin(this, this->objectName().toUtf8().data(), classid.toUtf8().data(), new QUrl(url), paramNames.join("|").toUtf8().data(), paramValues.join("|").toUtf8().data())); };
	QWebPage * createWindow(QWebPage::WebWindowType ty) { return static_cast<QWebPage*>(callbackQWebPage_CreateWindow(this, this->objectName().toUtf8().data(), ty)); };
	void Signal_DatabaseQuotaExceeded(QWebFrame * frame, QString databaseName) { callbackQWebPage_DatabaseQuotaExceeded(this, this->objectName().toUtf8().data(), frame, databaseName.toUtf8().data()); };
	void Signal_DownloadRequested(const QNetworkRequest & request) { callbackQWebPage_DownloadRequested(this, this->objectName().toUtf8().data(), new QNetworkRequest(request)); };
	bool event(QEvent * ev) { return callbackQWebPage_Event(this, this->objectName().toUtf8().data(), ev) != 0; };
	void Signal_FeaturePermissionRequestCanceled(QWebFrame * frame, QWebPage::Feature feature) { callbackQWebPage_FeaturePermissionRequestCanceled(this, this->objectName().toUtf8().data(), frame, feature); };
	void Signal_FeaturePermissionRequested(QWebFrame * frame, QWebPage::Feature feature) { callbackQWebPage_FeaturePermissionRequested(this, this->objectName().toUtf8().data(), frame, feature); };
	void Signal_FrameCreated(QWebFrame * frame) { callbackQWebPage_FrameCreated(this, this->objectName().toUtf8().data(), frame); };
	void Signal_GeometryChangeRequested(const QRect & geom) { callbackQWebPage_GeometryChangeRequested(this, this->objectName().toUtf8().data(), new QRect(static_cast<QRect>(geom).x(), static_cast<QRect>(geom).y(), static_cast<QRect>(geom).width(), static_cast<QRect>(geom).height())); };
	void javaScriptAlert(QWebFrame * frame, const QString & msg) { callbackQWebPage_JavaScriptAlert(this, this->objectName().toUtf8().data(), frame, msg.toUtf8().data()); };
	bool javaScriptConfirm(QWebFrame * frame, const QString & msg) { return callbackQWebPage_JavaScriptConfirm(this, this->objectName().toUtf8().data(), frame, msg.toUtf8().data()) != 0; };
	void javaScriptConsoleMessage(const QString & message, int lineNumber, const QString & sourceID) { callbackQWebPage_JavaScriptConsoleMessage(this, this->objectName().toUtf8().data(), message.toUtf8().data(), lineNumber, sourceID.toUtf8().data()); };
	bool javaScriptPrompt(QWebFrame * frame, const QString & msg, const QString & defaultValue, QString * result) { return callbackQWebPage_JavaScriptPrompt(this, this->objectName().toUtf8().data(), frame, msg.toUtf8().data(), defaultValue.toUtf8().data(), result->toUtf8().data()) != 0; };
	void Signal_LinkClicked(const QUrl & url) { callbackQWebPage_LinkClicked(this, this->objectName().toUtf8().data(), new QUrl(url)); };
	void Signal_LinkHovered(const QString & link, const QString & title, const QString & textContent) { callbackQWebPage_LinkHovered(this, this->objectName().toUtf8().data(), link.toUtf8().data(), title.toUtf8().data(), textContent.toUtf8().data()); };
	void Signal_LoadFinished(bool ok) { callbackQWebPage_LoadFinished(this, this->objectName().toUtf8().data(), ok); };
	void Signal_LoadProgress(int progress) { callbackQWebPage_LoadProgress(this, this->objectName().toUtf8().data(), progress); };
	void Signal_LoadStarted() { callbackQWebPage_LoadStarted(this, this->objectName().toUtf8().data()); };
	void Signal_MenuBarVisibilityChangeRequested(bool visible) { callbackQWebPage_MenuBarVisibilityChangeRequested(this, this->objectName().toUtf8().data(), visible); };
	void Signal_MicroFocusChanged() { callbackQWebPage_MicroFocusChanged(this, this->objectName().toUtf8().data()); };
	void Signal_PrintRequested(QWebFrame * frame) { callbackQWebPage_PrintRequested(this, this->objectName().toUtf8().data(), frame); };
	void Signal_RepaintRequested(const QRect & dirtyRect) { callbackQWebPage_RepaintRequested(this, this->objectName().toUtf8().data(), new QRect(static_cast<QRect>(dirtyRect).x(), static_cast<QRect>(dirtyRect).y(), static_cast<QRect>(dirtyRect).width(), static_cast<QRect>(dirtyRect).height())); };
	void Signal_RestoreFrameStateRequested(QWebFrame * frame) { callbackQWebPage_RestoreFrameStateRequested(this, this->objectName().toUtf8().data(), frame); };
	void Signal_SaveFrameStateRequested(QWebFrame * frame, QWebHistoryItem * item) { callbackQWebPage_SaveFrameStateRequested(this, this->objectName().toUtf8().data(), frame, item); };
	void Signal_ScrollRequested(int dx, int dy, const QRect & rectToScroll) { callbackQWebPage_ScrollRequested(this, this->objectName().toUtf8().data(), dx, dy, new QRect(static_cast<QRect>(rectToScroll).x(), static_cast<QRect>(rectToScroll).y(), static_cast<QRect>(rectToScroll).width(), static_cast<QRect>(rectToScroll).height())); };
	void Signal_SelectionChanged() { callbackQWebPage_SelectionChanged(this, this->objectName().toUtf8().data()); };
	void Signal_StatusBarMessage(const QString & text) { callbackQWebPage_StatusBarMessage(this, this->objectName().toUtf8().data(), text.toUtf8().data()); };
	void Signal_StatusBarVisibilityChangeRequested(bool visible) { callbackQWebPage_StatusBarVisibilityChangeRequested(this, this->objectName().toUtf8().data(), visible); };
	bool supportsExtension(QWebPage::Extension extension) const { return callbackQWebPage_SupportsExtension(const_cast<MyQWebPage*>(this), this->objectName().toUtf8().data(), extension) != 0; };
	void Signal_ToolBarVisibilityChangeRequested(bool visible) { callbackQWebPage_ToolBarVisibilityChangeRequested(this, this->objectName().toUtf8().data(), visible); };
	void triggerAction(QWebPage::WebAction action, bool checked) { callbackQWebPage_TriggerAction(this, this->objectName().toUtf8().data(), action, checked); };
	void Signal_UnsupportedContent(QNetworkReply * reply) { callbackQWebPage_UnsupportedContent(this, this->objectName().toUtf8().data(), reply); };
	QString userAgentForUrl(const QUrl & url) const { return QString(callbackQWebPage_UserAgentForUrl(const_cast<MyQWebPage*>(this), this->objectName().toUtf8().data(), new QUrl(url))); };
	void Signal_ViewportChangeRequested() { callbackQWebPage_ViewportChangeRequested(this, this->objectName().toUtf8().data()); };
	void Signal_WindowCloseRequested() { callbackQWebPage_WindowCloseRequested(this, this->objectName().toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQWebPage_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQWebPage_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQWebPage_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQWebPage_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQWebPage_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQWebPage_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQWebPage_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQWebPage_MetaObject(const_cast<MyQWebPage*>(this), this->objectName().toUtf8().data())); };
};

int QWebPage_AcceptNavigationRequest(void* ptr, void* frame, void* request, int ty)
{
	return static_cast<QWebPage*>(ptr)->acceptNavigationRequest(static_cast<QWebFrame*>(frame), *static_cast<QNetworkRequest*>(request), static_cast<QWebPage::NavigationType>(ty));
}

int QWebPage_AcceptNavigationRequestDefault(void* ptr, void* frame, void* request, int ty)
{
	return static_cast<QWebPage*>(ptr)->QWebPage::acceptNavigationRequest(static_cast<QWebFrame*>(frame), *static_cast<QNetworkRequest*>(request), static_cast<QWebPage::NavigationType>(ty));
}

int QWebPage_ForwardUnsupportedContent(void* ptr)
{
	return static_cast<QWebPage*>(ptr)->forwardUnsupportedContent();
}

int QWebPage_HasSelection(void* ptr)
{
	return static_cast<QWebPage*>(ptr)->hasSelection();
}

int QWebPage_IsContentEditable(void* ptr)
{
	return static_cast<QWebPage*>(ptr)->isContentEditable();
}

int QWebPage_IsModified(void* ptr)
{
	return static_cast<QWebPage*>(ptr)->isModified();
}

int QWebPage_LinkDelegationPolicy(void* ptr)
{
	return static_cast<QWebPage*>(ptr)->linkDelegationPolicy();
}

void* QWebPage_Palette(void* ptr)
{
	return new QPalette(static_cast<QWebPage*>(ptr)->palette());
}

void* QWebPage_PreferredContentsSize(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QWebPage*>(ptr)->preferredContentsSize()).width(), static_cast<QSize>(static_cast<QWebPage*>(ptr)->preferredContentsSize()).height());
}

char* QWebPage_SelectedHtml(void* ptr)
{
	return static_cast<QWebPage*>(ptr)->selectedHtml().toUtf8().data();
}

char* QWebPage_SelectedText(void* ptr)
{
	return static_cast<QWebPage*>(ptr)->selectedText().toUtf8().data();
}

void QWebPage_SetActualVisibleContentRect(void* ptr, void* rect)
{
	static_cast<QWebPage*>(ptr)->setActualVisibleContentRect(*static_cast<QRect*>(rect));
}

void QWebPage_SetContentEditable(void* ptr, int editable)
{
	static_cast<QWebPage*>(ptr)->setContentEditable(editable != 0);
}

void QWebPage_SetFeaturePermission(void* ptr, void* frame, int feature, int policy)
{
	static_cast<QWebPage*>(ptr)->setFeaturePermission(static_cast<QWebFrame*>(frame), static_cast<QWebPage::Feature>(feature), static_cast<QWebPage::PermissionPolicy>(policy));
}

void QWebPage_SetForwardUnsupportedContent(void* ptr, int forward)
{
	static_cast<QWebPage*>(ptr)->setForwardUnsupportedContent(forward != 0);
}

void QWebPage_SetLinkDelegationPolicy(void* ptr, int policy)
{
	static_cast<QWebPage*>(ptr)->setLinkDelegationPolicy(static_cast<QWebPage::LinkDelegationPolicy>(policy));
}

void QWebPage_SetPalette(void* ptr, void* palette)
{
	static_cast<QWebPage*>(ptr)->setPalette(*static_cast<QPalette*>(palette));
}

void QWebPage_SetPreferredContentsSize(void* ptr, void* size)
{
	static_cast<QWebPage*>(ptr)->setPreferredContentsSize(*static_cast<QSize*>(size));
}

void QWebPage_SetViewportSize(void* ptr, void* size)
{
	static_cast<QWebPage*>(ptr)->setViewportSize(*static_cast<QSize*>(size));
}

void QWebPage_SetVisibilityState(void* ptr, int vvi)
{
	static_cast<QWebPage*>(ptr)->setVisibilityState(static_cast<QWebPage::VisibilityState>(vvi));
}

int QWebPage_ShouldInterruptJavaScript(void* ptr)
{
	return static_cast<QWebPage*>(ptr)->shouldInterruptJavaScript();
}

int QWebPage_ShouldInterruptJavaScriptDefault(void* ptr)
{
	return static_cast<QWebPage*>(ptr)->QWebPage::shouldInterruptJavaScript();
}

void* QWebPage_ViewportSize(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QWebPage*>(ptr)->viewportSize()).width(), static_cast<QSize>(static_cast<QWebPage*>(ptr)->viewportSize()).height());
}

int QWebPage_VisibilityState(void* ptr)
{
	return static_cast<QWebPage*>(ptr)->visibilityState();
}

void* QWebPage_NewQWebPage(void* parent)
{
	return new MyQWebPage(static_cast<QObject*>(parent));
}

void* QWebPage_Action(void* ptr, int action)
{
	return static_cast<QWebPage*>(ptr)->action(static_cast<QWebPage::WebAction>(action));
}

char* QWebPage_ChooseFile(void* ptr, void* parentFrame, char* suggestedFile)
{
	return static_cast<QWebPage*>(ptr)->chooseFile(static_cast<QWebFrame*>(parentFrame), QString(suggestedFile)).toUtf8().data();
}

char* QWebPage_ChooseFileDefault(void* ptr, void* parentFrame, char* suggestedFile)
{
	return static_cast<QWebPage*>(ptr)->QWebPage::chooseFile(static_cast<QWebFrame*>(parentFrame), QString(suggestedFile)).toUtf8().data();
}

void QWebPage_ConnectContentsChanged(void* ptr)
{
	QObject::connect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)()>(&QWebPage::contentsChanged), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)()>(&MyQWebPage::Signal_ContentsChanged));
}

void QWebPage_DisconnectContentsChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)()>(&QWebPage::contentsChanged), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)()>(&MyQWebPage::Signal_ContentsChanged));
}

void QWebPage_ContentsChanged(void* ptr)
{
	static_cast<QWebPage*>(ptr)->contentsChanged();
}

void* QWebPage_CreatePlugin(void* ptr, char* classid, void* url, char* paramNames, char* paramValues)
{
	return static_cast<QWebPage*>(ptr)->createPlugin(QString(classid), *static_cast<QUrl*>(url), QString(paramNames).split("|", QString::SkipEmptyParts), QString(paramValues).split("|", QString::SkipEmptyParts));
}

void* QWebPage_CreatePluginDefault(void* ptr, char* classid, void* url, char* paramNames, char* paramValues)
{
	return static_cast<QWebPage*>(ptr)->QWebPage::createPlugin(QString(classid), *static_cast<QUrl*>(url), QString(paramNames).split("|", QString::SkipEmptyParts), QString(paramValues).split("|", QString::SkipEmptyParts));
}

void* QWebPage_CreateStandardContextMenu(void* ptr)
{
	return static_cast<QWebPage*>(ptr)->createStandardContextMenu();
}

void* QWebPage_CreateWindow(void* ptr, int ty)
{
	return static_cast<QWebPage*>(ptr)->createWindow(static_cast<QWebPage::WebWindowType>(ty));
}

void* QWebPage_CreateWindowDefault(void* ptr, int ty)
{
	return static_cast<QWebPage*>(ptr)->QWebPage::createWindow(static_cast<QWebPage::WebWindowType>(ty));
}

void* QWebPage_CurrentFrame(void* ptr)
{
	return static_cast<QWebPage*>(ptr)->currentFrame();
}

void QWebPage_ConnectDatabaseQuotaExceeded(void* ptr)
{
	QObject::connect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)(QWebFrame *, QString)>(&QWebPage::databaseQuotaExceeded), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)(QWebFrame *, QString)>(&MyQWebPage::Signal_DatabaseQuotaExceeded));
}

void QWebPage_DisconnectDatabaseQuotaExceeded(void* ptr)
{
	QObject::disconnect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)(QWebFrame *, QString)>(&QWebPage::databaseQuotaExceeded), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)(QWebFrame *, QString)>(&MyQWebPage::Signal_DatabaseQuotaExceeded));
}

void QWebPage_DatabaseQuotaExceeded(void* ptr, void* frame, char* databaseName)
{
	static_cast<QWebPage*>(ptr)->databaseQuotaExceeded(static_cast<QWebFrame*>(frame), QString(databaseName));
}

void QWebPage_ConnectDownloadRequested(void* ptr)
{
	QObject::connect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)(const QNetworkRequest &)>(&QWebPage::downloadRequested), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)(const QNetworkRequest &)>(&MyQWebPage::Signal_DownloadRequested));
}

void QWebPage_DisconnectDownloadRequested(void* ptr)
{
	QObject::disconnect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)(const QNetworkRequest &)>(&QWebPage::downloadRequested), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)(const QNetworkRequest &)>(&MyQWebPage::Signal_DownloadRequested));
}

void QWebPage_DownloadRequested(void* ptr, void* request)
{
	static_cast<QWebPage*>(ptr)->downloadRequested(*static_cast<QNetworkRequest*>(request));
}

int QWebPage_Event(void* ptr, void* ev)
{
	return static_cast<QWebPage*>(ptr)->event(static_cast<QEvent*>(ev));
}

int QWebPage_EventDefault(void* ptr, void* ev)
{
	return static_cast<QWebPage*>(ptr)->QWebPage::event(static_cast<QEvent*>(ev));
}

void QWebPage_ConnectFeaturePermissionRequestCanceled(void* ptr)
{
	QObject::connect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)(QWebFrame *, QWebPage::Feature)>(&QWebPage::featurePermissionRequestCanceled), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)(QWebFrame *, QWebPage::Feature)>(&MyQWebPage::Signal_FeaturePermissionRequestCanceled));
}

void QWebPage_DisconnectFeaturePermissionRequestCanceled(void* ptr)
{
	QObject::disconnect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)(QWebFrame *, QWebPage::Feature)>(&QWebPage::featurePermissionRequestCanceled), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)(QWebFrame *, QWebPage::Feature)>(&MyQWebPage::Signal_FeaturePermissionRequestCanceled));
}

void QWebPage_FeaturePermissionRequestCanceled(void* ptr, void* frame, int feature)
{
	static_cast<QWebPage*>(ptr)->featurePermissionRequestCanceled(static_cast<QWebFrame*>(frame), static_cast<QWebPage::Feature>(feature));
}

void QWebPage_ConnectFeaturePermissionRequested(void* ptr)
{
	QObject::connect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)(QWebFrame *, QWebPage::Feature)>(&QWebPage::featurePermissionRequested), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)(QWebFrame *, QWebPage::Feature)>(&MyQWebPage::Signal_FeaturePermissionRequested));
}

void QWebPage_DisconnectFeaturePermissionRequested(void* ptr)
{
	QObject::disconnect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)(QWebFrame *, QWebPage::Feature)>(&QWebPage::featurePermissionRequested), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)(QWebFrame *, QWebPage::Feature)>(&MyQWebPage::Signal_FeaturePermissionRequested));
}

void QWebPage_FeaturePermissionRequested(void* ptr, void* frame, int feature)
{
	static_cast<QWebPage*>(ptr)->featurePermissionRequested(static_cast<QWebFrame*>(frame), static_cast<QWebPage::Feature>(feature));
}

int QWebPage_FindText(void* ptr, char* subString, int options)
{
	return static_cast<QWebPage*>(ptr)->findText(QString(subString), static_cast<QWebPage::FindFlag>(options));
}

int QWebPage_FocusNextPrevChild(void* ptr, int next)
{
	return static_cast<QWebPage*>(ptr)->focusNextPrevChild(next != 0);
}

void* QWebPage_FrameAt(void* ptr, void* pos)
{
	return static_cast<QWebPage*>(ptr)->frameAt(*static_cast<QPoint*>(pos));
}

void QWebPage_ConnectFrameCreated(void* ptr)
{
	QObject::connect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)(QWebFrame *)>(&QWebPage::frameCreated), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)(QWebFrame *)>(&MyQWebPage::Signal_FrameCreated));
}

void QWebPage_DisconnectFrameCreated(void* ptr)
{
	QObject::disconnect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)(QWebFrame *)>(&QWebPage::frameCreated), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)(QWebFrame *)>(&MyQWebPage::Signal_FrameCreated));
}

void QWebPage_FrameCreated(void* ptr, void* frame)
{
	static_cast<QWebPage*>(ptr)->frameCreated(static_cast<QWebFrame*>(frame));
}

void QWebPage_ConnectGeometryChangeRequested(void* ptr)
{
	QObject::connect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)(const QRect &)>(&QWebPage::geometryChangeRequested), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)(const QRect &)>(&MyQWebPage::Signal_GeometryChangeRequested));
}

void QWebPage_DisconnectGeometryChangeRequested(void* ptr)
{
	QObject::disconnect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)(const QRect &)>(&QWebPage::geometryChangeRequested), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)(const QRect &)>(&MyQWebPage::Signal_GeometryChangeRequested));
}

void QWebPage_GeometryChangeRequested(void* ptr, void* geom)
{
	static_cast<QWebPage*>(ptr)->geometryChangeRequested(*static_cast<QRect*>(geom));
}

void* QWebPage_History(void* ptr)
{
	return static_cast<QWebPage*>(ptr)->history();
}

void* QWebPage_InputMethodQuery(void* ptr, int property)
{
	return new QVariant(static_cast<QWebPage*>(ptr)->inputMethodQuery(static_cast<Qt::InputMethodQuery>(property)));
}

void QWebPage_JavaScriptAlert(void* ptr, void* frame, char* msg)
{
	static_cast<QWebPage*>(ptr)->javaScriptAlert(static_cast<QWebFrame*>(frame), QString(msg));
}

void QWebPage_JavaScriptAlertDefault(void* ptr, void* frame, char* msg)
{
	static_cast<QWebPage*>(ptr)->QWebPage::javaScriptAlert(static_cast<QWebFrame*>(frame), QString(msg));
}

int QWebPage_JavaScriptConfirm(void* ptr, void* frame, char* msg)
{
	return static_cast<QWebPage*>(ptr)->javaScriptConfirm(static_cast<QWebFrame*>(frame), QString(msg));
}

int QWebPage_JavaScriptConfirmDefault(void* ptr, void* frame, char* msg)
{
	return static_cast<QWebPage*>(ptr)->QWebPage::javaScriptConfirm(static_cast<QWebFrame*>(frame), QString(msg));
}

void QWebPage_JavaScriptConsoleMessage(void* ptr, char* message, int lineNumber, char* sourceID)
{
	static_cast<QWebPage*>(ptr)->javaScriptConsoleMessage(QString(message), lineNumber, QString(sourceID));
}

void QWebPage_JavaScriptConsoleMessageDefault(void* ptr, char* message, int lineNumber, char* sourceID)
{
	static_cast<QWebPage*>(ptr)->QWebPage::javaScriptConsoleMessage(QString(message), lineNumber, QString(sourceID));
}

int QWebPage_JavaScriptPrompt(void* ptr, void* frame, char* msg, char* defaultValue, char* result)
{
	return static_cast<QWebPage*>(ptr)->javaScriptPrompt(static_cast<QWebFrame*>(frame), QString(msg), QString(defaultValue), new QString(result));
}

int QWebPage_JavaScriptPromptDefault(void* ptr, void* frame, char* msg, char* defaultValue, char* result)
{
	return static_cast<QWebPage*>(ptr)->QWebPage::javaScriptPrompt(static_cast<QWebFrame*>(frame), QString(msg), QString(defaultValue), new QString(result));
}

void QWebPage_ConnectLinkClicked(void* ptr)
{
	QObject::connect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)(const QUrl &)>(&QWebPage::linkClicked), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)(const QUrl &)>(&MyQWebPage::Signal_LinkClicked));
}

void QWebPage_DisconnectLinkClicked(void* ptr)
{
	QObject::disconnect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)(const QUrl &)>(&QWebPage::linkClicked), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)(const QUrl &)>(&MyQWebPage::Signal_LinkClicked));
}

void QWebPage_LinkClicked(void* ptr, void* url)
{
	static_cast<QWebPage*>(ptr)->linkClicked(*static_cast<QUrl*>(url));
}

void QWebPage_ConnectLinkHovered(void* ptr)
{
	QObject::connect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)(const QString &, const QString &, const QString &)>(&QWebPage::linkHovered), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)(const QString &, const QString &, const QString &)>(&MyQWebPage::Signal_LinkHovered));
}

void QWebPage_DisconnectLinkHovered(void* ptr)
{
	QObject::disconnect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)(const QString &, const QString &, const QString &)>(&QWebPage::linkHovered), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)(const QString &, const QString &, const QString &)>(&MyQWebPage::Signal_LinkHovered));
}

void QWebPage_LinkHovered(void* ptr, char* link, char* title, char* textContent)
{
	static_cast<QWebPage*>(ptr)->linkHovered(QString(link), QString(title), QString(textContent));
}

void QWebPage_ConnectLoadFinished(void* ptr)
{
	QObject::connect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)(bool)>(&QWebPage::loadFinished), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)(bool)>(&MyQWebPage::Signal_LoadFinished));
}

void QWebPage_DisconnectLoadFinished(void* ptr)
{
	QObject::disconnect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)(bool)>(&QWebPage::loadFinished), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)(bool)>(&MyQWebPage::Signal_LoadFinished));
}

void QWebPage_LoadFinished(void* ptr, int ok)
{
	static_cast<QWebPage*>(ptr)->loadFinished(ok != 0);
}

void QWebPage_ConnectLoadProgress(void* ptr)
{
	QObject::connect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)(int)>(&QWebPage::loadProgress), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)(int)>(&MyQWebPage::Signal_LoadProgress));
}

void QWebPage_DisconnectLoadProgress(void* ptr)
{
	QObject::disconnect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)(int)>(&QWebPage::loadProgress), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)(int)>(&MyQWebPage::Signal_LoadProgress));
}

void QWebPage_LoadProgress(void* ptr, int progress)
{
	static_cast<QWebPage*>(ptr)->loadProgress(progress);
}

void QWebPage_ConnectLoadStarted(void* ptr)
{
	QObject::connect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)()>(&QWebPage::loadStarted), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)()>(&MyQWebPage::Signal_LoadStarted));
}

void QWebPage_DisconnectLoadStarted(void* ptr)
{
	QObject::disconnect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)()>(&QWebPage::loadStarted), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)()>(&MyQWebPage::Signal_LoadStarted));
}

void QWebPage_LoadStarted(void* ptr)
{
	static_cast<QWebPage*>(ptr)->loadStarted();
}

void* QWebPage_MainFrame(void* ptr)
{
	return static_cast<QWebPage*>(ptr)->mainFrame();
}

void QWebPage_ConnectMenuBarVisibilityChangeRequested(void* ptr)
{
	QObject::connect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)(bool)>(&QWebPage::menuBarVisibilityChangeRequested), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)(bool)>(&MyQWebPage::Signal_MenuBarVisibilityChangeRequested));
}

void QWebPage_DisconnectMenuBarVisibilityChangeRequested(void* ptr)
{
	QObject::disconnect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)(bool)>(&QWebPage::menuBarVisibilityChangeRequested), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)(bool)>(&MyQWebPage::Signal_MenuBarVisibilityChangeRequested));
}

void QWebPage_MenuBarVisibilityChangeRequested(void* ptr, int visible)
{
	static_cast<QWebPage*>(ptr)->menuBarVisibilityChangeRequested(visible != 0);
}

void QWebPage_ConnectMicroFocusChanged(void* ptr)
{
	QObject::connect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)()>(&QWebPage::microFocusChanged), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)()>(&MyQWebPage::Signal_MicroFocusChanged));
}

void QWebPage_DisconnectMicroFocusChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)()>(&QWebPage::microFocusChanged), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)()>(&MyQWebPage::Signal_MicroFocusChanged));
}

void QWebPage_MicroFocusChanged(void* ptr)
{
	static_cast<QWebPage*>(ptr)->microFocusChanged();
}

void* QWebPage_NetworkAccessManager(void* ptr)
{
	return static_cast<QWebPage*>(ptr)->networkAccessManager();
}

void* QWebPage_PluginFactory(void* ptr)
{
	return static_cast<QWebPage*>(ptr)->pluginFactory();
}

void QWebPage_ConnectPrintRequested(void* ptr)
{
	QObject::connect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)(QWebFrame *)>(&QWebPage::printRequested), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)(QWebFrame *)>(&MyQWebPage::Signal_PrintRequested));
}

void QWebPage_DisconnectPrintRequested(void* ptr)
{
	QObject::disconnect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)(QWebFrame *)>(&QWebPage::printRequested), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)(QWebFrame *)>(&MyQWebPage::Signal_PrintRequested));
}

void QWebPage_PrintRequested(void* ptr, void* frame)
{
	static_cast<QWebPage*>(ptr)->printRequested(static_cast<QWebFrame*>(frame));
}

void QWebPage_ConnectRepaintRequested(void* ptr)
{
	QObject::connect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)(const QRect &)>(&QWebPage::repaintRequested), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)(const QRect &)>(&MyQWebPage::Signal_RepaintRequested));
}

void QWebPage_DisconnectRepaintRequested(void* ptr)
{
	QObject::disconnect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)(const QRect &)>(&QWebPage::repaintRequested), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)(const QRect &)>(&MyQWebPage::Signal_RepaintRequested));
}

void QWebPage_RepaintRequested(void* ptr, void* dirtyRect)
{
	static_cast<QWebPage*>(ptr)->repaintRequested(*static_cast<QRect*>(dirtyRect));
}

void QWebPage_ConnectRestoreFrameStateRequested(void* ptr)
{
	QObject::connect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)(QWebFrame *)>(&QWebPage::restoreFrameStateRequested), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)(QWebFrame *)>(&MyQWebPage::Signal_RestoreFrameStateRequested));
}

void QWebPage_DisconnectRestoreFrameStateRequested(void* ptr)
{
	QObject::disconnect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)(QWebFrame *)>(&QWebPage::restoreFrameStateRequested), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)(QWebFrame *)>(&MyQWebPage::Signal_RestoreFrameStateRequested));
}

void QWebPage_RestoreFrameStateRequested(void* ptr, void* frame)
{
	static_cast<QWebPage*>(ptr)->restoreFrameStateRequested(static_cast<QWebFrame*>(frame));
}

void QWebPage_ConnectSaveFrameStateRequested(void* ptr)
{
	QObject::connect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)(QWebFrame *, QWebHistoryItem *)>(&QWebPage::saveFrameStateRequested), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)(QWebFrame *, QWebHistoryItem *)>(&MyQWebPage::Signal_SaveFrameStateRequested));
}

void QWebPage_DisconnectSaveFrameStateRequested(void* ptr)
{
	QObject::disconnect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)(QWebFrame *, QWebHistoryItem *)>(&QWebPage::saveFrameStateRequested), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)(QWebFrame *, QWebHistoryItem *)>(&MyQWebPage::Signal_SaveFrameStateRequested));
}

void QWebPage_SaveFrameStateRequested(void* ptr, void* frame, void* item)
{
	static_cast<QWebPage*>(ptr)->saveFrameStateRequested(static_cast<QWebFrame*>(frame), static_cast<QWebHistoryItem*>(item));
}

void QWebPage_ConnectScrollRequested(void* ptr)
{
	QObject::connect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)(int, int, const QRect &)>(&QWebPage::scrollRequested), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)(int, int, const QRect &)>(&MyQWebPage::Signal_ScrollRequested));
}

void QWebPage_DisconnectScrollRequested(void* ptr)
{
	QObject::disconnect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)(int, int, const QRect &)>(&QWebPage::scrollRequested), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)(int, int, const QRect &)>(&MyQWebPage::Signal_ScrollRequested));
}

void QWebPage_ScrollRequested(void* ptr, int dx, int dy, void* rectToScroll)
{
	static_cast<QWebPage*>(ptr)->scrollRequested(dx, dy, *static_cast<QRect*>(rectToScroll));
}

void QWebPage_ConnectSelectionChanged(void* ptr)
{
	QObject::connect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)()>(&QWebPage::selectionChanged), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)()>(&MyQWebPage::Signal_SelectionChanged));
}

void QWebPage_DisconnectSelectionChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)()>(&QWebPage::selectionChanged), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)()>(&MyQWebPage::Signal_SelectionChanged));
}

void QWebPage_SelectionChanged(void* ptr)
{
	static_cast<QWebPage*>(ptr)->selectionChanged();
}

void QWebPage_SetNetworkAccessManager(void* ptr, void* manager)
{
	static_cast<QWebPage*>(ptr)->setNetworkAccessManager(static_cast<QNetworkAccessManager*>(manager));
}

void QWebPage_SetPluginFactory(void* ptr, void* factory)
{
	static_cast<QWebPage*>(ptr)->setPluginFactory(static_cast<QWebPluginFactory*>(factory));
}

void QWebPage_SetView(void* ptr, void* view)
{
	static_cast<QWebPage*>(ptr)->setView(static_cast<QWidget*>(view));
}

void* QWebPage_Settings(void* ptr)
{
	return static_cast<QWebPage*>(ptr)->settings();
}

void QWebPage_ConnectStatusBarMessage(void* ptr)
{
	QObject::connect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)(const QString &)>(&QWebPage::statusBarMessage), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)(const QString &)>(&MyQWebPage::Signal_StatusBarMessage));
}

void QWebPage_DisconnectStatusBarMessage(void* ptr)
{
	QObject::disconnect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)(const QString &)>(&QWebPage::statusBarMessage), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)(const QString &)>(&MyQWebPage::Signal_StatusBarMessage));
}

void QWebPage_StatusBarMessage(void* ptr, char* text)
{
	static_cast<QWebPage*>(ptr)->statusBarMessage(QString(text));
}

void QWebPage_ConnectStatusBarVisibilityChangeRequested(void* ptr)
{
	QObject::connect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)(bool)>(&QWebPage::statusBarVisibilityChangeRequested), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)(bool)>(&MyQWebPage::Signal_StatusBarVisibilityChangeRequested));
}

void QWebPage_DisconnectStatusBarVisibilityChangeRequested(void* ptr)
{
	QObject::disconnect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)(bool)>(&QWebPage::statusBarVisibilityChangeRequested), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)(bool)>(&MyQWebPage::Signal_StatusBarVisibilityChangeRequested));
}

void QWebPage_StatusBarVisibilityChangeRequested(void* ptr, int visible)
{
	static_cast<QWebPage*>(ptr)->statusBarVisibilityChangeRequested(visible != 0);
}

char* QWebPage_SupportedContentTypes(void* ptr)
{
	return static_cast<QWebPage*>(ptr)->supportedContentTypes().join("|").toUtf8().data();
}

int QWebPage_SupportsContentType(void* ptr, char* mimeType)
{
	return static_cast<QWebPage*>(ptr)->supportsContentType(QString(mimeType));
}

int QWebPage_SupportsExtension(void* ptr, int extension)
{
	return static_cast<QWebPage*>(ptr)->supportsExtension(static_cast<QWebPage::Extension>(extension));
}

int QWebPage_SupportsExtensionDefault(void* ptr, int extension)
{
	return static_cast<QWebPage*>(ptr)->QWebPage::supportsExtension(static_cast<QWebPage::Extension>(extension));
}

int QWebPage_SwallowContextMenuEvent(void* ptr, void* event)
{
	return static_cast<QWebPage*>(ptr)->swallowContextMenuEvent(static_cast<QContextMenuEvent*>(event));
}

void QWebPage_ConnectToolBarVisibilityChangeRequested(void* ptr)
{
	QObject::connect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)(bool)>(&QWebPage::toolBarVisibilityChangeRequested), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)(bool)>(&MyQWebPage::Signal_ToolBarVisibilityChangeRequested));
}

void QWebPage_DisconnectToolBarVisibilityChangeRequested(void* ptr)
{
	QObject::disconnect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)(bool)>(&QWebPage::toolBarVisibilityChangeRequested), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)(bool)>(&MyQWebPage::Signal_ToolBarVisibilityChangeRequested));
}

void QWebPage_ToolBarVisibilityChangeRequested(void* ptr, int visible)
{
	static_cast<QWebPage*>(ptr)->toolBarVisibilityChangeRequested(visible != 0);
}

void QWebPage_TriggerAction(void* ptr, int action, int checked)
{
	static_cast<QWebPage*>(ptr)->triggerAction(static_cast<QWebPage::WebAction>(action), checked != 0);
}

void QWebPage_TriggerActionDefault(void* ptr, int action, int checked)
{
	static_cast<QWebPage*>(ptr)->QWebPage::triggerAction(static_cast<QWebPage::WebAction>(action), checked != 0);
}

void* QWebPage_UndoStack(void* ptr)
{
	return static_cast<QWebPage*>(ptr)->undoStack();
}

void QWebPage_ConnectUnsupportedContent(void* ptr)
{
	QObject::connect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)(QNetworkReply *)>(&QWebPage::unsupportedContent), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)(QNetworkReply *)>(&MyQWebPage::Signal_UnsupportedContent));
}

void QWebPage_DisconnectUnsupportedContent(void* ptr)
{
	QObject::disconnect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)(QNetworkReply *)>(&QWebPage::unsupportedContent), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)(QNetworkReply *)>(&MyQWebPage::Signal_UnsupportedContent));
}

void QWebPage_UnsupportedContent(void* ptr, void* reply)
{
	static_cast<QWebPage*>(ptr)->unsupportedContent(static_cast<QNetworkReply*>(reply));
}

void QWebPage_UpdatePositionDependentActions(void* ptr, void* pos)
{
	static_cast<QWebPage*>(ptr)->updatePositionDependentActions(*static_cast<QPoint*>(pos));
}

char* QWebPage_UserAgentForUrl(void* ptr, void* url)
{
	return static_cast<QWebPage*>(ptr)->userAgentForUrl(*static_cast<QUrl*>(url)).toUtf8().data();
}

char* QWebPage_UserAgentForUrlDefault(void* ptr, void* url)
{
	return static_cast<QWebPage*>(ptr)->QWebPage::userAgentForUrl(*static_cast<QUrl*>(url)).toUtf8().data();
}

void* QWebPage_View(void* ptr)
{
	return static_cast<QWebPage*>(ptr)->view();
}

void QWebPage_ConnectViewportChangeRequested(void* ptr)
{
	QObject::connect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)()>(&QWebPage::viewportChangeRequested), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)()>(&MyQWebPage::Signal_ViewportChangeRequested));
}

void QWebPage_DisconnectViewportChangeRequested(void* ptr)
{
	QObject::disconnect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)()>(&QWebPage::viewportChangeRequested), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)()>(&MyQWebPage::Signal_ViewportChangeRequested));
}

void QWebPage_ViewportChangeRequested(void* ptr)
{
	static_cast<QWebPage*>(ptr)->viewportChangeRequested();
}

void QWebPage_ConnectWindowCloseRequested(void* ptr)
{
	QObject::connect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)()>(&QWebPage::windowCloseRequested), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)()>(&MyQWebPage::Signal_WindowCloseRequested));
}

void QWebPage_DisconnectWindowCloseRequested(void* ptr)
{
	QObject::disconnect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)()>(&QWebPage::windowCloseRequested), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)()>(&MyQWebPage::Signal_WindowCloseRequested));
}

void QWebPage_WindowCloseRequested(void* ptr)
{
	static_cast<QWebPage*>(ptr)->windowCloseRequested();
}

void QWebPage_DestroyQWebPage(void* ptr)
{
	static_cast<QWebPage*>(ptr)->~QWebPage();
}

void QWebPage_TimerEvent(void* ptr, void* event)
{
	static_cast<QWebPage*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QWebPage_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QWebPage*>(ptr)->QWebPage::timerEvent(static_cast<QTimerEvent*>(event));
}

void QWebPage_ChildEvent(void* ptr, void* event)
{
	static_cast<QWebPage*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QWebPage_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QWebPage*>(ptr)->QWebPage::childEvent(static_cast<QChildEvent*>(event));
}

void QWebPage_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QWebPage*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWebPage_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QWebPage*>(ptr)->QWebPage::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWebPage_CustomEvent(void* ptr, void* event)
{
	static_cast<QWebPage*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QWebPage_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QWebPage*>(ptr)->QWebPage::customEvent(static_cast<QEvent*>(event));
}

void QWebPage_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QWebPage*>(ptr), "deleteLater");
}

void QWebPage_DeleteLaterDefault(void* ptr)
{
	static_cast<QWebPage*>(ptr)->QWebPage::deleteLater();
}

void QWebPage_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QWebPage*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWebPage_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QWebPage*>(ptr)->QWebPage::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QWebPage_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QWebPage*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QWebPage_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QWebPage*>(ptr)->QWebPage::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QWebPage_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QWebPage*>(ptr)->metaObject());
}

void* QWebPage_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QWebPage*>(ptr)->QWebPage::metaObject());
}

class MyQWebPluginFactory: public QWebPluginFactory
{
public:
	QObject * create(const QString & mimeType, const QUrl & url, const QStringList & argumentNames, const QStringList & argumentValues) const { return static_cast<QObject*>(callbackQWebPluginFactory_Create(const_cast<MyQWebPluginFactory*>(this), this->objectName().toUtf8().data(), mimeType.toUtf8().data(), new QUrl(url), argumentNames.join("|").toUtf8().data(), argumentValues.join("|").toUtf8().data())); };
	void refreshPlugins() { callbackQWebPluginFactory_RefreshPlugins(this, this->objectName().toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQWebPluginFactory_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQWebPluginFactory_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQWebPluginFactory_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQWebPluginFactory_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQWebPluginFactory_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQWebPluginFactory_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQWebPluginFactory_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQWebPluginFactory_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQWebPluginFactory_MetaObject(const_cast<MyQWebPluginFactory*>(this), this->objectName().toUtf8().data())); };
};

void* QWebPluginFactory_Create(void* ptr, char* mimeType, void* url, char* argumentNames, char* argumentValues)
{
	return static_cast<QWebPluginFactory*>(ptr)->create(QString(mimeType), *static_cast<QUrl*>(url), QString(argumentNames).split("|", QString::SkipEmptyParts), QString(argumentValues).split("|", QString::SkipEmptyParts));
}

void QWebPluginFactory_RefreshPlugins(void* ptr)
{
	static_cast<QWebPluginFactory*>(ptr)->refreshPlugins();
}

void QWebPluginFactory_RefreshPluginsDefault(void* ptr)
{
	static_cast<QWebPluginFactory*>(ptr)->QWebPluginFactory::refreshPlugins();
}

void QWebPluginFactory_DestroyQWebPluginFactory(void* ptr)
{
	static_cast<QWebPluginFactory*>(ptr)->~QWebPluginFactory();
}

void QWebPluginFactory_TimerEvent(void* ptr, void* event)
{
	static_cast<QWebPluginFactory*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QWebPluginFactory_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QWebPluginFactory*>(ptr)->QWebPluginFactory::timerEvent(static_cast<QTimerEvent*>(event));
}

void QWebPluginFactory_ChildEvent(void* ptr, void* event)
{
	static_cast<QWebPluginFactory*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QWebPluginFactory_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QWebPluginFactory*>(ptr)->QWebPluginFactory::childEvent(static_cast<QChildEvent*>(event));
}

void QWebPluginFactory_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QWebPluginFactory*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWebPluginFactory_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QWebPluginFactory*>(ptr)->QWebPluginFactory::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWebPluginFactory_CustomEvent(void* ptr, void* event)
{
	static_cast<QWebPluginFactory*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QWebPluginFactory_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QWebPluginFactory*>(ptr)->QWebPluginFactory::customEvent(static_cast<QEvent*>(event));
}

void QWebPluginFactory_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QWebPluginFactory*>(ptr), "deleteLater");
}

void QWebPluginFactory_DeleteLaterDefault(void* ptr)
{
	static_cast<QWebPluginFactory*>(ptr)->QWebPluginFactory::deleteLater();
}

void QWebPluginFactory_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QWebPluginFactory*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWebPluginFactory_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QWebPluginFactory*>(ptr)->QWebPluginFactory::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QWebPluginFactory_Event(void* ptr, void* e)
{
	return static_cast<QWebPluginFactory*>(ptr)->event(static_cast<QEvent*>(e));
}

int QWebPluginFactory_EventDefault(void* ptr, void* e)
{
	return static_cast<QWebPluginFactory*>(ptr)->QWebPluginFactory::event(static_cast<QEvent*>(e));
}

int QWebPluginFactory_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QWebPluginFactory*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QWebPluginFactory_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QWebPluginFactory*>(ptr)->QWebPluginFactory::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QWebPluginFactory_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QWebPluginFactory*>(ptr)->metaObject());
}

void* QWebPluginFactory_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QWebPluginFactory*>(ptr)->QWebPluginFactory::metaObject());
}

void QWebSecurityOrigin_SetApplicationCacheQuota(void* ptr, long long quota)
{
	static_cast<QWebSecurityOrigin*>(ptr)->setApplicationCacheQuota(static_cast<long long>(quota));
}

void* QWebSecurityOrigin_NewQWebSecurityOrigin(void* url)
{
	return new QWebSecurityOrigin(*static_cast<QUrl*>(url));
}

void* QWebSecurityOrigin_NewQWebSecurityOrigin2(void* other)
{
	return new QWebSecurityOrigin(*static_cast<QWebSecurityOrigin*>(other));
}

void QWebSecurityOrigin_AddAccessWhitelistEntry(void* ptr, char* scheme, char* host, int subdomainSetting)
{
	static_cast<QWebSecurityOrigin*>(ptr)->addAccessWhitelistEntry(QString(scheme), QString(host), static_cast<QWebSecurityOrigin::SubdomainSetting>(subdomainSetting));
}

void QWebSecurityOrigin_QWebSecurityOrigin_AddLocalScheme(char* scheme)
{
	QWebSecurityOrigin::addLocalScheme(QString(scheme));
}

long long QWebSecurityOrigin_DatabaseQuota(void* ptr)
{
	return static_cast<long long>(static_cast<QWebSecurityOrigin*>(ptr)->databaseQuota());
}

long long QWebSecurityOrigin_DatabaseUsage(void* ptr)
{
	return static_cast<long long>(static_cast<QWebSecurityOrigin*>(ptr)->databaseUsage());
}

char* QWebSecurityOrigin_Host(void* ptr)
{
	return static_cast<QWebSecurityOrigin*>(ptr)->host().toUtf8().data();
}

char* QWebSecurityOrigin_QWebSecurityOrigin_LocalSchemes()
{
	return QWebSecurityOrigin::localSchemes().join("|").toUtf8().data();
}

int QWebSecurityOrigin_Port(void* ptr)
{
	return static_cast<QWebSecurityOrigin*>(ptr)->port();
}

void QWebSecurityOrigin_RemoveAccessWhitelistEntry(void* ptr, char* scheme, char* host, int subdomainSetting)
{
	static_cast<QWebSecurityOrigin*>(ptr)->removeAccessWhitelistEntry(QString(scheme), QString(host), static_cast<QWebSecurityOrigin::SubdomainSetting>(subdomainSetting));
}

void QWebSecurityOrigin_QWebSecurityOrigin_RemoveLocalScheme(char* scheme)
{
	QWebSecurityOrigin::removeLocalScheme(QString(scheme));
}

char* QWebSecurityOrigin_Scheme(void* ptr)
{
	return static_cast<QWebSecurityOrigin*>(ptr)->scheme().toUtf8().data();
}

void QWebSecurityOrigin_SetDatabaseQuota(void* ptr, long long quota)
{
	static_cast<QWebSecurityOrigin*>(ptr)->setDatabaseQuota(static_cast<long long>(quota));
}

void QWebSecurityOrigin_DestroyQWebSecurityOrigin(void* ptr)
{
	static_cast<QWebSecurityOrigin*>(ptr)->~QWebSecurityOrigin();
}

int QWebSettings_LocalContentCanAccessRemoteUrls_Type()
{
	return QWebSettings::LocalContentCanAccessRemoteUrls;
}

int QWebSettings_DnsPrefetchEnabled_Type()
{
	return QWebSettings::DnsPrefetchEnabled;
}

int QWebSettings_XSSAuditingEnabled_Type()
{
	return QWebSettings::XSSAuditingEnabled;
}

int QWebSettings_AcceleratedCompositingEnabled_Type()
{
	return QWebSettings::AcceleratedCompositingEnabled;
}

int QWebSettings_SpatialNavigationEnabled_Type()
{
	return QWebSettings::SpatialNavigationEnabled;
}

int QWebSettings_LocalContentCanAccessFileUrls_Type()
{
	return QWebSettings::LocalContentCanAccessFileUrls;
}

int QWebSettings_TiledBackingStoreEnabled_Type()
{
	return QWebSettings::TiledBackingStoreEnabled;
}

int QWebSettings_FrameFlatteningEnabled_Type()
{
	return QWebSettings::FrameFlatteningEnabled;
}

int QWebSettings_SiteSpecificQuirksEnabled_Type()
{
	return QWebSettings::SiteSpecificQuirksEnabled;
}

int QWebSettings_JavascriptCanCloseWindows_Type()
{
	return QWebSettings::JavascriptCanCloseWindows;
}

int QWebSettings_WebGLEnabled_Type()
{
	return QWebSettings::WebGLEnabled;
}

int QWebSettings_CSSRegionsEnabled_Type()
{
	return QWebSettings::CSSRegionsEnabled;
}

int QWebSettings_HyperlinkAuditingEnabled_Type()
{
	return QWebSettings::HyperlinkAuditingEnabled;
}

int QWebSettings_CSSGridLayoutEnabled_Type()
{
	return QWebSettings::CSSGridLayoutEnabled;
}

int QWebSettings_ScrollAnimatorEnabled_Type()
{
	return QWebSettings::ScrollAnimatorEnabled;
}

int QWebSettings_CaretBrowsingEnabled_Type()
{
	return QWebSettings::CaretBrowsingEnabled;
}

int QWebSettings_NotificationsEnabled_Type()
{
	return QWebSettings::NotificationsEnabled;
}

int QWebSettings_WebAudioEnabled_Type()
{
	return QWebSettings::WebAudioEnabled;
}

int QWebSettings_Accelerated2dCanvasEnabled_Type()
{
	return QWebSettings::Accelerated2dCanvasEnabled;
}

void QWebSettings_ResetAttribute(void* ptr, int attribute)
{
	static_cast<QWebSettings*>(ptr)->resetAttribute(static_cast<QWebSettings::WebAttribute>(attribute));
}

void QWebSettings_SetAttribute(void* ptr, int attribute, int on)
{
	static_cast<QWebSettings*>(ptr)->setAttribute(static_cast<QWebSettings::WebAttribute>(attribute), on != 0);
}

int QWebSettings_TestAttribute(void* ptr, int attribute)
{
	return static_cast<QWebSettings*>(ptr)->testAttribute(static_cast<QWebSettings::WebAttribute>(attribute));
}

void QWebSettings_QWebSettings_ClearIconDatabase()
{
	QWebSettings::clearIconDatabase();
}

void QWebSettings_QWebSettings_ClearMemoryCaches()
{
	QWebSettings::clearMemoryCaches();
}

char* QWebSettings_CssMediaType(void* ptr)
{
	return static_cast<QWebSettings*>(ptr)->cssMediaType().toUtf8().data();
}

char* QWebSettings_DefaultTextEncoding(void* ptr)
{
	return static_cast<QWebSettings*>(ptr)->defaultTextEncoding().toUtf8().data();
}

void QWebSettings_QWebSettings_EnablePersistentStorage(char* path)
{
	QWebSettings::enablePersistentStorage(QString(path));
}

char* QWebSettings_FontFamily(void* ptr, int which)
{
	return static_cast<QWebSettings*>(ptr)->fontFamily(static_cast<QWebSettings::FontFamily>(which)).toUtf8().data();
}

int QWebSettings_FontSize(void* ptr, int ty)
{
	return static_cast<QWebSettings*>(ptr)->fontSize(static_cast<QWebSettings::FontSize>(ty));
}

void* QWebSettings_QWebSettings_GlobalSettings()
{
	return QWebSettings::globalSettings();
}

char* QWebSettings_QWebSettings_IconDatabasePath()
{
	return QWebSettings::iconDatabasePath().toUtf8().data();
}

void* QWebSettings_QWebSettings_IconForUrl(void* url)
{
	return new QIcon(QWebSettings::iconForUrl(*static_cast<QUrl*>(url)));
}

char* QWebSettings_LocalStoragePath(void* ptr)
{
	return static_cast<QWebSettings*>(ptr)->localStoragePath().toUtf8().data();
}

int QWebSettings_QWebSettings_MaximumPagesInCache()
{
	return QWebSettings::maximumPagesInCache();
}

long long QWebSettings_QWebSettings_OfflineStorageDefaultQuota()
{
	return static_cast<long long>(QWebSettings::offlineStorageDefaultQuota());
}

char* QWebSettings_QWebSettings_OfflineStoragePath()
{
	return QWebSettings::offlineStoragePath().toUtf8().data();
}

char* QWebSettings_QWebSettings_OfflineWebApplicationCachePath()
{
	return QWebSettings::offlineWebApplicationCachePath().toUtf8().data();
}

long long QWebSettings_QWebSettings_OfflineWebApplicationCacheQuota()
{
	return static_cast<long long>(QWebSettings::offlineWebApplicationCacheQuota());
}

void QWebSettings_ResetFontFamily(void* ptr, int which)
{
	static_cast<QWebSettings*>(ptr)->resetFontFamily(static_cast<QWebSettings::FontFamily>(which));
}

void QWebSettings_ResetFontSize(void* ptr, int ty)
{
	static_cast<QWebSettings*>(ptr)->resetFontSize(static_cast<QWebSettings::FontSize>(ty));
}

void QWebSettings_SetCSSMediaType(void* ptr, char* ty)
{
	static_cast<QWebSettings*>(ptr)->setCSSMediaType(QString(ty));
}

void QWebSettings_SetDefaultTextEncoding(void* ptr, char* encoding)
{
	static_cast<QWebSettings*>(ptr)->setDefaultTextEncoding(QString(encoding));
}

void QWebSettings_SetFontFamily(void* ptr, int which, char* family)
{
	static_cast<QWebSettings*>(ptr)->setFontFamily(static_cast<QWebSettings::FontFamily>(which), QString(family));
}

void QWebSettings_SetFontSize(void* ptr, int ty, int size)
{
	static_cast<QWebSettings*>(ptr)->setFontSize(static_cast<QWebSettings::FontSize>(ty), size);
}

void QWebSettings_QWebSettings_SetIconDatabasePath(char* path)
{
	QWebSettings::setIconDatabasePath(QString(path));
}

void QWebSettings_SetLocalStoragePath(void* ptr, char* path)
{
	static_cast<QWebSettings*>(ptr)->setLocalStoragePath(QString(path));
}

void QWebSettings_QWebSettings_SetMaximumPagesInCache(int pages)
{
	QWebSettings::setMaximumPagesInCache(pages);
}

void QWebSettings_QWebSettings_SetObjectCacheCapacities(int cacheMinDeadCapacity, int cacheMaxDead, int totalCapacity)
{
	QWebSettings::setObjectCacheCapacities(cacheMinDeadCapacity, cacheMaxDead, totalCapacity);
}

void QWebSettings_QWebSettings_SetOfflineStorageDefaultQuota(long long maximumSize)
{
	QWebSettings::setOfflineStorageDefaultQuota(static_cast<long long>(maximumSize));
}

void QWebSettings_QWebSettings_SetOfflineStoragePath(char* path)
{
	QWebSettings::setOfflineStoragePath(QString(path));
}

void QWebSettings_QWebSettings_SetOfflineWebApplicationCachePath(char* path)
{
	QWebSettings::setOfflineWebApplicationCachePath(QString(path));
}

void QWebSettings_QWebSettings_SetOfflineWebApplicationCacheQuota(long long maximumSize)
{
	QWebSettings::setOfflineWebApplicationCacheQuota(static_cast<long long>(maximumSize));
}

void QWebSettings_SetThirdPartyCookiePolicy(void* ptr, int policy)
{
	static_cast<QWebSettings*>(ptr)->setThirdPartyCookiePolicy(static_cast<QWebSettings::ThirdPartyCookiePolicy>(policy));
}

void QWebSettings_SetUserStyleSheetUrl(void* ptr, void* location)
{
	static_cast<QWebSettings*>(ptr)->setUserStyleSheetUrl(*static_cast<QUrl*>(location));
}

void QWebSettings_QWebSettings_SetWebGraphic(int ty, void* graphic)
{
	QWebSettings::setWebGraphic(static_cast<QWebSettings::WebGraphic>(ty), *static_cast<QPixmap*>(graphic));
}

int QWebSettings_ThirdPartyCookiePolicy(void* ptr)
{
	return static_cast<QWebSettings*>(ptr)->thirdPartyCookiePolicy();
}

void* QWebSettings_UserStyleSheetUrl(void* ptr)
{
	return new QUrl(static_cast<QWebSettings*>(ptr)->userStyleSheetUrl());
}

void* QWebSettings_QWebSettings_WebGraphic(int ty)
{
	return new QPixmap(QWebSettings::webGraphic(static_cast<QWebSettings::WebGraphic>(ty)));
}

class MyQWebView: public QWebView
{
public:
	MyQWebView(QWidget *parent) : QWebView(parent) {};
	void back() { callbackQWebView_Back(this, this->objectName().toUtf8().data()); };
	void changeEvent(QEvent * e) { callbackQWebView_ChangeEvent(this, this->objectName().toUtf8().data(), e); };
	void contextMenuEvent(QContextMenuEvent * ev) { callbackQWebView_ContextMenuEvent(this, this->objectName().toUtf8().data(), ev); };
	QWebView * createWindow(QWebPage::WebWindowType ty) { return static_cast<QWebView*>(callbackQWebView_CreateWindow(this, this->objectName().toUtf8().data(), ty)); };
	void dragEnterEvent(QDragEnterEvent * ev) { callbackQWebView_DragEnterEvent(this, this->objectName().toUtf8().data(), ev); };
	void dragLeaveEvent(QDragLeaveEvent * ev) { callbackQWebView_DragLeaveEvent(this, this->objectName().toUtf8().data(), ev); };
	void dragMoveEvent(QDragMoveEvent * ev) { callbackQWebView_DragMoveEvent(this, this->objectName().toUtf8().data(), ev); };
	void dropEvent(QDropEvent * ev) { callbackQWebView_DropEvent(this, this->objectName().toUtf8().data(), ev); };
	bool event(QEvent * e) { return callbackQWebView_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	void focusInEvent(QFocusEvent * ev) { callbackQWebView_FocusInEvent(this, this->objectName().toUtf8().data(), ev); };
	bool focusNextPrevChild(bool next) { return callbackQWebView_FocusNextPrevChild(this, this->objectName().toUtf8().data(), next) != 0; };
	void focusOutEvent(QFocusEvent * ev) { callbackQWebView_FocusOutEvent(this, this->objectName().toUtf8().data(), ev); };
	void forward() { callbackQWebView_Forward(this, this->objectName().toUtf8().data()); };
	void Signal_IconChanged() { callbackQWebView_IconChanged(this, this->objectName().toUtf8().data()); };
	void inputMethodEvent(QInputMethodEvent * e) { callbackQWebView_InputMethodEvent(this, this->objectName().toUtf8().data(), e); };
	QVariant inputMethodQuery(Qt::InputMethodQuery property) const { return *static_cast<QVariant*>(callbackQWebView_InputMethodQuery(const_cast<MyQWebView*>(this), this->objectName().toUtf8().data(), property)); };
	void keyPressEvent(QKeyEvent * ev) { callbackQWebView_KeyPressEvent(this, this->objectName().toUtf8().data(), ev); };
	void keyReleaseEvent(QKeyEvent * ev) { callbackQWebView_KeyReleaseEvent(this, this->objectName().toUtf8().data(), ev); };
	void Signal_LinkClicked(const QUrl & url) { callbackQWebView_LinkClicked(this, this->objectName().toUtf8().data(), new QUrl(url)); };
	void Signal_LoadFinished(bool ok) { callbackQWebView_LoadFinished(this, this->objectName().toUtf8().data(), ok); };
	void Signal_LoadProgress(int progress) { callbackQWebView_LoadProgress(this, this->objectName().toUtf8().data(), progress); };
	void Signal_LoadStarted() { callbackQWebView_LoadStarted(this, this->objectName().toUtf8().data()); };
	void mouseDoubleClickEvent(QMouseEvent * ev) { callbackQWebView_MouseDoubleClickEvent(this, this->objectName().toUtf8().data(), ev); };
	void mouseMoveEvent(QMouseEvent * ev) { callbackQWebView_MouseMoveEvent(this, this->objectName().toUtf8().data(), ev); };
	void mousePressEvent(QMouseEvent * ev) { callbackQWebView_MousePressEvent(this, this->objectName().toUtf8().data(), ev); };
	void mouseReleaseEvent(QMouseEvent * ev) { callbackQWebView_MouseReleaseEvent(this, this->objectName().toUtf8().data(), ev); };
	void paintEvent(QPaintEvent * ev) { callbackQWebView_PaintEvent(this, this->objectName().toUtf8().data(), ev); };
	void print(QPrinter * printer) const { callbackQWebView_Print(const_cast<MyQWebView*>(this), this->objectName().toUtf8().data(), printer); };
	void reload() { callbackQWebView_Reload(this, this->objectName().toUtf8().data()); };
	void resizeEvent(QResizeEvent * e) { callbackQWebView_ResizeEvent(this, this->objectName().toUtf8().data(), e); };
	void Signal_SelectionChanged() { callbackQWebView_SelectionChanged(this, this->objectName().toUtf8().data()); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQWebView_SizeHint(const_cast<MyQWebView*>(this), this->objectName().toUtf8().data())); };
	void Signal_StatusBarMessage(const QString & text) { callbackQWebView_StatusBarMessage(this, this->objectName().toUtf8().data(), text.toUtf8().data()); };
	void stop() { callbackQWebView_Stop(this, this->objectName().toUtf8().data()); };
	void Signal_TitleChanged(const QString & title) { callbackQWebView_TitleChanged(this, this->objectName().toUtf8().data(), title.toUtf8().data()); };
	void Signal_UrlChanged(const QUrl & url) { callbackQWebView_UrlChanged(this, this->objectName().toUtf8().data(), new QUrl(url)); };
	void wheelEvent(QWheelEvent * ev) { callbackQWebView_WheelEvent(this, this->objectName().toUtf8().data(), ev); };
	void actionEvent(QActionEvent * event) { callbackQWebView_ActionEvent(this, this->objectName().toUtf8().data(), event); };
	void enterEvent(QEvent * event) { callbackQWebView_EnterEvent(this, this->objectName().toUtf8().data(), event); };
	void hideEvent(QHideEvent * event) { callbackQWebView_HideEvent(this, this->objectName().toUtf8().data(), event); };
	void leaveEvent(QEvent * event) { callbackQWebView_LeaveEvent(this, this->objectName().toUtf8().data(), event); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackQWebView_Metric(const_cast<MyQWebView*>(this), this->objectName().toUtf8().data(), m); };
	QSize minimumSizeHint() const { return *static_cast<QSize*>(callbackQWebView_MinimumSizeHint(const_cast<MyQWebView*>(this), this->objectName().toUtf8().data())); };
	void moveEvent(QMoveEvent * event) { callbackQWebView_MoveEvent(this, this->objectName().toUtf8().data(), event); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQWebView_PaintEngine(const_cast<MyQWebView*>(this), this->objectName().toUtf8().data())); };
	void setEnabled(bool vbo) { callbackQWebView_SetEnabled(this, this->objectName().toUtf8().data(), vbo); };
	void setStyleSheet(const QString & styleSheet) { callbackQWebView_SetStyleSheet(this, this->objectName().toUtf8().data(), styleSheet.toUtf8().data()); };
	void setVisible(bool visible) { callbackQWebView_SetVisible(this, this->objectName().toUtf8().data(), visible); };
	void setWindowModified(bool vbo) { callbackQWebView_SetWindowModified(this, this->objectName().toUtf8().data(), vbo); };
	void setWindowTitle(const QString & vqs) { callbackQWebView_SetWindowTitle(this, this->objectName().toUtf8().data(), vqs.toUtf8().data()); };
	void showEvent(QShowEvent * event) { callbackQWebView_ShowEvent(this, this->objectName().toUtf8().data(), event); };
	bool close() { return callbackQWebView_Close(this, this->objectName().toUtf8().data()) != 0; };
	void closeEvent(QCloseEvent * event) { callbackQWebView_CloseEvent(this, this->objectName().toUtf8().data(), event); };
	bool hasHeightForWidth() const { return callbackQWebView_HasHeightForWidth(const_cast<MyQWebView*>(this), this->objectName().toUtf8().data()) != 0; };
	int heightForWidth(int w) const { return callbackQWebView_HeightForWidth(const_cast<MyQWebView*>(this), this->objectName().toUtf8().data(), w); };
	void hide() { callbackQWebView_Hide(this, this->objectName().toUtf8().data()); };
	void initPainter(QPainter * painter) const { callbackQWebView_InitPainter(const_cast<MyQWebView*>(this), this->objectName().toUtf8().data(), painter); };
	void lower() { callbackQWebView_Lower(this, this->objectName().toUtf8().data()); };
	bool nativeEvent(const QByteArray & eventType, void * message, long * result) { return callbackQWebView_NativeEvent(this, this->objectName().toUtf8().data(), QString(eventType).toUtf8().data(), message, *result) != 0; };
	void raise() { callbackQWebView_Raise(this, this->objectName().toUtf8().data()); };
	void repaint() { callbackQWebView_Repaint(this, this->objectName().toUtf8().data()); };
	void setDisabled(bool disable) { callbackQWebView_SetDisabled(this, this->objectName().toUtf8().data(), disable); };
	void setFocus() { callbackQWebView_SetFocus2(this, this->objectName().toUtf8().data()); };
	void setHidden(bool hidden) { callbackQWebView_SetHidden(this, this->objectName().toUtf8().data(), hidden); };
	void show() { callbackQWebView_Show(this, this->objectName().toUtf8().data()); };
	void showFullScreen() { callbackQWebView_ShowFullScreen(this, this->objectName().toUtf8().data()); };
	void showMaximized() { callbackQWebView_ShowMaximized(this, this->objectName().toUtf8().data()); };
	void showMinimized() { callbackQWebView_ShowMinimized(this, this->objectName().toUtf8().data()); };
	void showNormal() { callbackQWebView_ShowNormal(this, this->objectName().toUtf8().data()); };
	void tabletEvent(QTabletEvent * event) { callbackQWebView_TabletEvent(this, this->objectName().toUtf8().data(), event); };
	void update() { callbackQWebView_Update(this, this->objectName().toUtf8().data()); };
	void updateMicroFocus() { callbackQWebView_UpdateMicroFocus(this, this->objectName().toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQWebView_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQWebView_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQWebView_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQWebView_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQWebView_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQWebView_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQWebView_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQWebView_MetaObject(const_cast<MyQWebView*>(this), this->objectName().toUtf8().data())); };
};

int QWebView_HasSelection(void* ptr)
{
	return static_cast<QWebView*>(ptr)->hasSelection();
}

void* QWebView_Icon(void* ptr)
{
	return new QIcon(static_cast<QWebView*>(ptr)->icon());
}

int QWebView_IsModified(void* ptr)
{
	return static_cast<QWebView*>(ptr)->isModified();
}

void QWebView_Load2(void* ptr, void* request, int operation, char* body)
{
	static_cast<QWebView*>(ptr)->load(*static_cast<QNetworkRequest*>(request), static_cast<QNetworkAccessManager::Operation>(operation), QByteArray(body));
}

char* QWebView_SelectedHtml(void* ptr)
{
	return static_cast<QWebView*>(ptr)->selectedHtml().toUtf8().data();
}

char* QWebView_SelectedText(void* ptr)
{
	return static_cast<QWebView*>(ptr)->selectedText().toUtf8().data();
}

void QWebView_SetUrl(void* ptr, void* url)
{
	static_cast<QWebView*>(ptr)->setUrl(*static_cast<QUrl*>(url));
}

void QWebView_SetZoomFactor(void* ptr, double factor)
{
	static_cast<QWebView*>(ptr)->setZoomFactor(static_cast<double>(factor));
}

char* QWebView_Title(void* ptr)
{
	return static_cast<QWebView*>(ptr)->title().toUtf8().data();
}

void* QWebView_Url(void* ptr)
{
	return new QUrl(static_cast<QWebView*>(ptr)->url());
}

double QWebView_ZoomFactor(void* ptr)
{
	return static_cast<double>(static_cast<QWebView*>(ptr)->zoomFactor());
}

void* QWebView_NewQWebView(void* parent)
{
	return new MyQWebView(static_cast<QWidget*>(parent));
}

void QWebView_Back(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QWebView*>(ptr), "back");
}

void QWebView_ChangeEvent(void* ptr, void* e)
{
	static_cast<QWebView*>(ptr)->changeEvent(static_cast<QEvent*>(e));
}

void QWebView_ChangeEventDefault(void* ptr, void* e)
{
	static_cast<QWebView*>(ptr)->QWebView::changeEvent(static_cast<QEvent*>(e));
}

void QWebView_ContextMenuEvent(void* ptr, void* ev)
{
	static_cast<QWebView*>(ptr)->contextMenuEvent(static_cast<QContextMenuEvent*>(ev));
}

void QWebView_ContextMenuEventDefault(void* ptr, void* ev)
{
	static_cast<QWebView*>(ptr)->QWebView::contextMenuEvent(static_cast<QContextMenuEvent*>(ev));
}

void* QWebView_CreateWindow(void* ptr, int ty)
{
	return static_cast<QWebView*>(ptr)->createWindow(static_cast<QWebPage::WebWindowType>(ty));
}

void* QWebView_CreateWindowDefault(void* ptr, int ty)
{
	return static_cast<QWebView*>(ptr)->QWebView::createWindow(static_cast<QWebPage::WebWindowType>(ty));
}

void QWebView_DragEnterEvent(void* ptr, void* ev)
{
	static_cast<QWebView*>(ptr)->dragEnterEvent(static_cast<QDragEnterEvent*>(ev));
}

void QWebView_DragEnterEventDefault(void* ptr, void* ev)
{
	static_cast<QWebView*>(ptr)->QWebView::dragEnterEvent(static_cast<QDragEnterEvent*>(ev));
}

void QWebView_DragLeaveEvent(void* ptr, void* ev)
{
	static_cast<QWebView*>(ptr)->dragLeaveEvent(static_cast<QDragLeaveEvent*>(ev));
}

void QWebView_DragLeaveEventDefault(void* ptr, void* ev)
{
	static_cast<QWebView*>(ptr)->QWebView::dragLeaveEvent(static_cast<QDragLeaveEvent*>(ev));
}

void QWebView_DragMoveEvent(void* ptr, void* ev)
{
	static_cast<QWebView*>(ptr)->dragMoveEvent(static_cast<QDragMoveEvent*>(ev));
}

void QWebView_DragMoveEventDefault(void* ptr, void* ev)
{
	static_cast<QWebView*>(ptr)->QWebView::dragMoveEvent(static_cast<QDragMoveEvent*>(ev));
}

void QWebView_DropEvent(void* ptr, void* ev)
{
	static_cast<QWebView*>(ptr)->dropEvent(static_cast<QDropEvent*>(ev));
}

void QWebView_DropEventDefault(void* ptr, void* ev)
{
	static_cast<QWebView*>(ptr)->QWebView::dropEvent(static_cast<QDropEvent*>(ev));
}

int QWebView_Event(void* ptr, void* e)
{
	return static_cast<QWebView*>(ptr)->event(static_cast<QEvent*>(e));
}

int QWebView_EventDefault(void* ptr, void* e)
{
	return static_cast<QWebView*>(ptr)->QWebView::event(static_cast<QEvent*>(e));
}

int QWebView_FindText(void* ptr, char* subString, int options)
{
	return static_cast<QWebView*>(ptr)->findText(QString(subString), static_cast<QWebPage::FindFlag>(options));
}

void QWebView_FocusInEvent(void* ptr, void* ev)
{
	static_cast<QWebView*>(ptr)->focusInEvent(static_cast<QFocusEvent*>(ev));
}

void QWebView_FocusInEventDefault(void* ptr, void* ev)
{
	static_cast<QWebView*>(ptr)->QWebView::focusInEvent(static_cast<QFocusEvent*>(ev));
}

int QWebView_FocusNextPrevChild(void* ptr, int next)
{
	return static_cast<QWebView*>(ptr)->focusNextPrevChild(next != 0);
}

int QWebView_FocusNextPrevChildDefault(void* ptr, int next)
{
	return static_cast<QWebView*>(ptr)->QWebView::focusNextPrevChild(next != 0);
}

void QWebView_FocusOutEvent(void* ptr, void* ev)
{
	static_cast<QWebView*>(ptr)->focusOutEvent(static_cast<QFocusEvent*>(ev));
}

void QWebView_FocusOutEventDefault(void* ptr, void* ev)
{
	static_cast<QWebView*>(ptr)->QWebView::focusOutEvent(static_cast<QFocusEvent*>(ev));
}

void QWebView_Forward(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QWebView*>(ptr), "forward");
}

void* QWebView_History(void* ptr)
{
	return static_cast<QWebView*>(ptr)->history();
}

void QWebView_ConnectIconChanged(void* ptr)
{
	QObject::connect(static_cast<QWebView*>(ptr), static_cast<void (QWebView::*)()>(&QWebView::iconChanged), static_cast<MyQWebView*>(ptr), static_cast<void (MyQWebView::*)()>(&MyQWebView::Signal_IconChanged));
}

void QWebView_DisconnectIconChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWebView*>(ptr), static_cast<void (QWebView::*)()>(&QWebView::iconChanged), static_cast<MyQWebView*>(ptr), static_cast<void (MyQWebView::*)()>(&MyQWebView::Signal_IconChanged));
}

void QWebView_IconChanged(void* ptr)
{
	static_cast<QWebView*>(ptr)->iconChanged();
}

void QWebView_InputMethodEvent(void* ptr, void* e)
{
	static_cast<QWebView*>(ptr)->inputMethodEvent(static_cast<QInputMethodEvent*>(e));
}

void QWebView_InputMethodEventDefault(void* ptr, void* e)
{
	static_cast<QWebView*>(ptr)->QWebView::inputMethodEvent(static_cast<QInputMethodEvent*>(e));
}

void* QWebView_InputMethodQuery(void* ptr, int property)
{
	return new QVariant(static_cast<QWebView*>(ptr)->inputMethodQuery(static_cast<Qt::InputMethodQuery>(property)));
}

void* QWebView_InputMethodQueryDefault(void* ptr, int property)
{
	return new QVariant(static_cast<QWebView*>(ptr)->QWebView::inputMethodQuery(static_cast<Qt::InputMethodQuery>(property)));
}

void QWebView_KeyPressEvent(void* ptr, void* ev)
{
	static_cast<QWebView*>(ptr)->keyPressEvent(static_cast<QKeyEvent*>(ev));
}

void QWebView_KeyPressEventDefault(void* ptr, void* ev)
{
	static_cast<QWebView*>(ptr)->QWebView::keyPressEvent(static_cast<QKeyEvent*>(ev));
}

void QWebView_KeyReleaseEvent(void* ptr, void* ev)
{
	static_cast<QWebView*>(ptr)->keyReleaseEvent(static_cast<QKeyEvent*>(ev));
}

void QWebView_KeyReleaseEventDefault(void* ptr, void* ev)
{
	static_cast<QWebView*>(ptr)->QWebView::keyReleaseEvent(static_cast<QKeyEvent*>(ev));
}

void QWebView_ConnectLinkClicked(void* ptr)
{
	QObject::connect(static_cast<QWebView*>(ptr), static_cast<void (QWebView::*)(const QUrl &)>(&QWebView::linkClicked), static_cast<MyQWebView*>(ptr), static_cast<void (MyQWebView::*)(const QUrl &)>(&MyQWebView::Signal_LinkClicked));
}

void QWebView_DisconnectLinkClicked(void* ptr)
{
	QObject::disconnect(static_cast<QWebView*>(ptr), static_cast<void (QWebView::*)(const QUrl &)>(&QWebView::linkClicked), static_cast<MyQWebView*>(ptr), static_cast<void (MyQWebView::*)(const QUrl &)>(&MyQWebView::Signal_LinkClicked));
}

void QWebView_LinkClicked(void* ptr, void* url)
{
	static_cast<QWebView*>(ptr)->linkClicked(*static_cast<QUrl*>(url));
}

void QWebView_Load(void* ptr, void* url)
{
	static_cast<QWebView*>(ptr)->load(*static_cast<QUrl*>(url));
}

void QWebView_ConnectLoadFinished(void* ptr)
{
	QObject::connect(static_cast<QWebView*>(ptr), static_cast<void (QWebView::*)(bool)>(&QWebView::loadFinished), static_cast<MyQWebView*>(ptr), static_cast<void (MyQWebView::*)(bool)>(&MyQWebView::Signal_LoadFinished));
}

void QWebView_DisconnectLoadFinished(void* ptr)
{
	QObject::disconnect(static_cast<QWebView*>(ptr), static_cast<void (QWebView::*)(bool)>(&QWebView::loadFinished), static_cast<MyQWebView*>(ptr), static_cast<void (MyQWebView::*)(bool)>(&MyQWebView::Signal_LoadFinished));
}

void QWebView_LoadFinished(void* ptr, int ok)
{
	static_cast<QWebView*>(ptr)->loadFinished(ok != 0);
}

void QWebView_ConnectLoadProgress(void* ptr)
{
	QObject::connect(static_cast<QWebView*>(ptr), static_cast<void (QWebView::*)(int)>(&QWebView::loadProgress), static_cast<MyQWebView*>(ptr), static_cast<void (MyQWebView::*)(int)>(&MyQWebView::Signal_LoadProgress));
}

void QWebView_DisconnectLoadProgress(void* ptr)
{
	QObject::disconnect(static_cast<QWebView*>(ptr), static_cast<void (QWebView::*)(int)>(&QWebView::loadProgress), static_cast<MyQWebView*>(ptr), static_cast<void (MyQWebView::*)(int)>(&MyQWebView::Signal_LoadProgress));
}

void QWebView_LoadProgress(void* ptr, int progress)
{
	static_cast<QWebView*>(ptr)->loadProgress(progress);
}

void QWebView_ConnectLoadStarted(void* ptr)
{
	QObject::connect(static_cast<QWebView*>(ptr), static_cast<void (QWebView::*)()>(&QWebView::loadStarted), static_cast<MyQWebView*>(ptr), static_cast<void (MyQWebView::*)()>(&MyQWebView::Signal_LoadStarted));
}

void QWebView_DisconnectLoadStarted(void* ptr)
{
	QObject::disconnect(static_cast<QWebView*>(ptr), static_cast<void (QWebView::*)()>(&QWebView::loadStarted), static_cast<MyQWebView*>(ptr), static_cast<void (MyQWebView::*)()>(&MyQWebView::Signal_LoadStarted));
}

void QWebView_LoadStarted(void* ptr)
{
	static_cast<QWebView*>(ptr)->loadStarted();
}

void QWebView_MouseDoubleClickEvent(void* ptr, void* ev)
{
	static_cast<QWebView*>(ptr)->mouseDoubleClickEvent(static_cast<QMouseEvent*>(ev));
}

void QWebView_MouseDoubleClickEventDefault(void* ptr, void* ev)
{
	static_cast<QWebView*>(ptr)->QWebView::mouseDoubleClickEvent(static_cast<QMouseEvent*>(ev));
}

void QWebView_MouseMoveEvent(void* ptr, void* ev)
{
	static_cast<QWebView*>(ptr)->mouseMoveEvent(static_cast<QMouseEvent*>(ev));
}

void QWebView_MouseMoveEventDefault(void* ptr, void* ev)
{
	static_cast<QWebView*>(ptr)->QWebView::mouseMoveEvent(static_cast<QMouseEvent*>(ev));
}

void QWebView_MousePressEvent(void* ptr, void* ev)
{
	static_cast<QWebView*>(ptr)->mousePressEvent(static_cast<QMouseEvent*>(ev));
}

void QWebView_MousePressEventDefault(void* ptr, void* ev)
{
	static_cast<QWebView*>(ptr)->QWebView::mousePressEvent(static_cast<QMouseEvent*>(ev));
}

void QWebView_MouseReleaseEvent(void* ptr, void* ev)
{
	static_cast<QWebView*>(ptr)->mouseReleaseEvent(static_cast<QMouseEvent*>(ev));
}

void QWebView_MouseReleaseEventDefault(void* ptr, void* ev)
{
	static_cast<QWebView*>(ptr)->QWebView::mouseReleaseEvent(static_cast<QMouseEvent*>(ev));
}

void* QWebView_Page(void* ptr)
{
	return static_cast<QWebView*>(ptr)->page();
}

void* QWebView_PageAction(void* ptr, int action)
{
	return static_cast<QWebView*>(ptr)->pageAction(static_cast<QWebPage::WebAction>(action));
}

void QWebView_PaintEvent(void* ptr, void* ev)
{
	static_cast<QWebView*>(ptr)->paintEvent(static_cast<QPaintEvent*>(ev));
}

void QWebView_PaintEventDefault(void* ptr, void* ev)
{
	static_cast<QWebView*>(ptr)->QWebView::paintEvent(static_cast<QPaintEvent*>(ev));
}

void QWebView_Print(void* ptr, void* printer)
{
	QMetaObject::invokeMethod(static_cast<QWebView*>(ptr), "print", Q_ARG(QPrinter*, static_cast<QPrinter*>(printer)));
}

void QWebView_Reload(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QWebView*>(ptr), "reload");
}

int QWebView_RenderHints(void* ptr)
{
	return static_cast<QWebView*>(ptr)->renderHints();
}

void QWebView_ResizeEvent(void* ptr, void* e)
{
	static_cast<QWebView*>(ptr)->resizeEvent(static_cast<QResizeEvent*>(e));
}

void QWebView_ResizeEventDefault(void* ptr, void* e)
{
	static_cast<QWebView*>(ptr)->QWebView::resizeEvent(static_cast<QResizeEvent*>(e));
}

void QWebView_ConnectSelectionChanged(void* ptr)
{
	QObject::connect(static_cast<QWebView*>(ptr), static_cast<void (QWebView::*)()>(&QWebView::selectionChanged), static_cast<MyQWebView*>(ptr), static_cast<void (MyQWebView::*)()>(&MyQWebView::Signal_SelectionChanged));
}

void QWebView_DisconnectSelectionChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWebView*>(ptr), static_cast<void (QWebView::*)()>(&QWebView::selectionChanged), static_cast<MyQWebView*>(ptr), static_cast<void (MyQWebView::*)()>(&MyQWebView::Signal_SelectionChanged));
}

void QWebView_SelectionChanged(void* ptr)
{
	static_cast<QWebView*>(ptr)->selectionChanged();
}

void QWebView_SetContent(void* ptr, char* data, char* mimeType, void* baseUrl)
{
	static_cast<QWebView*>(ptr)->setContent(QByteArray(data), QString(mimeType), *static_cast<QUrl*>(baseUrl));
}

void QWebView_SetHtml(void* ptr, char* html, void* baseUrl)
{
	static_cast<QWebView*>(ptr)->setHtml(QString(html), *static_cast<QUrl*>(baseUrl));
}

void QWebView_SetPage(void* ptr, void* page)
{
	static_cast<QWebView*>(ptr)->setPage(static_cast<QWebPage*>(page));
}

void QWebView_SetRenderHint(void* ptr, int hint, int enabled)
{
	static_cast<QWebView*>(ptr)->setRenderHint(static_cast<QPainter::RenderHint>(hint), enabled != 0);
}

void QWebView_SetRenderHints(void* ptr, int hints)
{
	static_cast<QWebView*>(ptr)->setRenderHints(static_cast<QPainter::RenderHint>(hints));
}

void QWebView_SetTextSizeMultiplier(void* ptr, double factor)
{
	static_cast<QWebView*>(ptr)->setTextSizeMultiplier(static_cast<double>(factor));
}

void* QWebView_Settings(void* ptr)
{
	return static_cast<QWebView*>(ptr)->settings();
}

void* QWebView_SizeHint(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QWebView*>(ptr)->sizeHint()).width(), static_cast<QSize>(static_cast<QWebView*>(ptr)->sizeHint()).height());
}

void* QWebView_SizeHintDefault(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QWebView*>(ptr)->QWebView::sizeHint()).width(), static_cast<QSize>(static_cast<QWebView*>(ptr)->QWebView::sizeHint()).height());
}

void QWebView_ConnectStatusBarMessage(void* ptr)
{
	QObject::connect(static_cast<QWebView*>(ptr), static_cast<void (QWebView::*)(const QString &)>(&QWebView::statusBarMessage), static_cast<MyQWebView*>(ptr), static_cast<void (MyQWebView::*)(const QString &)>(&MyQWebView::Signal_StatusBarMessage));
}

void QWebView_DisconnectStatusBarMessage(void* ptr)
{
	QObject::disconnect(static_cast<QWebView*>(ptr), static_cast<void (QWebView::*)(const QString &)>(&QWebView::statusBarMessage), static_cast<MyQWebView*>(ptr), static_cast<void (MyQWebView::*)(const QString &)>(&MyQWebView::Signal_StatusBarMessage));
}

void QWebView_StatusBarMessage(void* ptr, char* text)
{
	static_cast<QWebView*>(ptr)->statusBarMessage(QString(text));
}

void QWebView_Stop(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QWebView*>(ptr), "stop");
}

double QWebView_TextSizeMultiplier(void* ptr)
{
	return static_cast<double>(static_cast<QWebView*>(ptr)->textSizeMultiplier());
}

void QWebView_ConnectTitleChanged(void* ptr)
{
	QObject::connect(static_cast<QWebView*>(ptr), static_cast<void (QWebView::*)(const QString &)>(&QWebView::titleChanged), static_cast<MyQWebView*>(ptr), static_cast<void (MyQWebView::*)(const QString &)>(&MyQWebView::Signal_TitleChanged));
}

void QWebView_DisconnectTitleChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWebView*>(ptr), static_cast<void (QWebView::*)(const QString &)>(&QWebView::titleChanged), static_cast<MyQWebView*>(ptr), static_cast<void (MyQWebView::*)(const QString &)>(&MyQWebView::Signal_TitleChanged));
}

void QWebView_TitleChanged(void* ptr, char* title)
{
	static_cast<QWebView*>(ptr)->titleChanged(QString(title));
}

void QWebView_TriggerPageAction(void* ptr, int action, int checked)
{
	static_cast<QWebView*>(ptr)->triggerPageAction(static_cast<QWebPage::WebAction>(action), checked != 0);
}

void QWebView_ConnectUrlChanged(void* ptr)
{
	QObject::connect(static_cast<QWebView*>(ptr), static_cast<void (QWebView::*)(const QUrl &)>(&QWebView::urlChanged), static_cast<MyQWebView*>(ptr), static_cast<void (MyQWebView::*)(const QUrl &)>(&MyQWebView::Signal_UrlChanged));
}

void QWebView_DisconnectUrlChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWebView*>(ptr), static_cast<void (QWebView::*)(const QUrl &)>(&QWebView::urlChanged), static_cast<MyQWebView*>(ptr), static_cast<void (MyQWebView::*)(const QUrl &)>(&MyQWebView::Signal_UrlChanged));
}

void QWebView_UrlChanged(void* ptr, void* url)
{
	static_cast<QWebView*>(ptr)->urlChanged(*static_cast<QUrl*>(url));
}

void QWebView_WheelEvent(void* ptr, void* ev)
{
	static_cast<QWebView*>(ptr)->wheelEvent(static_cast<QWheelEvent*>(ev));
}

void QWebView_WheelEventDefault(void* ptr, void* ev)
{
	static_cast<QWebView*>(ptr)->QWebView::wheelEvent(static_cast<QWheelEvent*>(ev));
}

void QWebView_DestroyQWebView(void* ptr)
{
	static_cast<QWebView*>(ptr)->~QWebView();
}

void QWebView_ActionEvent(void* ptr, void* event)
{
	static_cast<QWebView*>(ptr)->actionEvent(static_cast<QActionEvent*>(event));
}

void QWebView_ActionEventDefault(void* ptr, void* event)
{
	static_cast<QWebView*>(ptr)->QWebView::actionEvent(static_cast<QActionEvent*>(event));
}

void QWebView_EnterEvent(void* ptr, void* event)
{
	static_cast<QWebView*>(ptr)->enterEvent(static_cast<QEvent*>(event));
}

void QWebView_EnterEventDefault(void* ptr, void* event)
{
	static_cast<QWebView*>(ptr)->QWebView::enterEvent(static_cast<QEvent*>(event));
}

void QWebView_HideEvent(void* ptr, void* event)
{
	static_cast<QWebView*>(ptr)->hideEvent(static_cast<QHideEvent*>(event));
}

void QWebView_HideEventDefault(void* ptr, void* event)
{
	static_cast<QWebView*>(ptr)->QWebView::hideEvent(static_cast<QHideEvent*>(event));
}

void QWebView_LeaveEvent(void* ptr, void* event)
{
	static_cast<QWebView*>(ptr)->leaveEvent(static_cast<QEvent*>(event));
}

void QWebView_LeaveEventDefault(void* ptr, void* event)
{
	static_cast<QWebView*>(ptr)->QWebView::leaveEvent(static_cast<QEvent*>(event));
}

int QWebView_Metric(void* ptr, int m)
{
	return static_cast<QWebView*>(ptr)->metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
}

int QWebView_MetricDefault(void* ptr, int m)
{
	return static_cast<QWebView*>(ptr)->QWebView::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
}

void* QWebView_MinimumSizeHint(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QWebView*>(ptr)->minimumSizeHint()).width(), static_cast<QSize>(static_cast<QWebView*>(ptr)->minimumSizeHint()).height());
}

void* QWebView_MinimumSizeHintDefault(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QWebView*>(ptr)->QWebView::minimumSizeHint()).width(), static_cast<QSize>(static_cast<QWebView*>(ptr)->QWebView::minimumSizeHint()).height());
}

void QWebView_MoveEvent(void* ptr, void* event)
{
	static_cast<QWebView*>(ptr)->moveEvent(static_cast<QMoveEvent*>(event));
}

void QWebView_MoveEventDefault(void* ptr, void* event)
{
	static_cast<QWebView*>(ptr)->QWebView::moveEvent(static_cast<QMoveEvent*>(event));
}

void* QWebView_PaintEngine(void* ptr)
{
	return static_cast<QWebView*>(ptr)->paintEngine();
}

void* QWebView_PaintEngineDefault(void* ptr)
{
	return static_cast<QWebView*>(ptr)->QWebView::paintEngine();
}

void QWebView_SetEnabled(void* ptr, int vbo)
{
	QMetaObject::invokeMethod(static_cast<QWebView*>(ptr), "setEnabled", Q_ARG(bool, vbo != 0));
}

void QWebView_SetEnabledDefault(void* ptr, int vbo)
{
	static_cast<QWebView*>(ptr)->QWebView::setEnabled(vbo != 0);
}

void QWebView_SetStyleSheet(void* ptr, char* styleSheet)
{
	QMetaObject::invokeMethod(static_cast<QWebView*>(ptr), "setStyleSheet", Q_ARG(QString, QString(styleSheet)));
}

void QWebView_SetStyleSheetDefault(void* ptr, char* styleSheet)
{
	static_cast<QWebView*>(ptr)->QWebView::setStyleSheet(QString(styleSheet));
}

void QWebView_SetVisible(void* ptr, int visible)
{
	QMetaObject::invokeMethod(static_cast<QWebView*>(ptr), "setVisible", Q_ARG(bool, visible != 0));
}

void QWebView_SetVisibleDefault(void* ptr, int visible)
{
	static_cast<QWebView*>(ptr)->QWebView::setVisible(visible != 0);
}

void QWebView_SetWindowModified(void* ptr, int vbo)
{
	QMetaObject::invokeMethod(static_cast<QWebView*>(ptr), "setWindowModified", Q_ARG(bool, vbo != 0));
}

void QWebView_SetWindowModifiedDefault(void* ptr, int vbo)
{
	static_cast<QWebView*>(ptr)->QWebView::setWindowModified(vbo != 0);
}

void QWebView_SetWindowTitle(void* ptr, char* vqs)
{
	QMetaObject::invokeMethod(static_cast<QWebView*>(ptr), "setWindowTitle", Q_ARG(QString, QString(vqs)));
}

void QWebView_SetWindowTitleDefault(void* ptr, char* vqs)
{
	static_cast<QWebView*>(ptr)->QWebView::setWindowTitle(QString(vqs));
}

void QWebView_ShowEvent(void* ptr, void* event)
{
	static_cast<QWebView*>(ptr)->showEvent(static_cast<QShowEvent*>(event));
}

void QWebView_ShowEventDefault(void* ptr, void* event)
{
	static_cast<QWebView*>(ptr)->QWebView::showEvent(static_cast<QShowEvent*>(event));
}

int QWebView_Close(void* ptr)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QWebView*>(ptr), "close", Q_RETURN_ARG(bool, returnArg));
	return returnArg;
}

int QWebView_CloseDefault(void* ptr)
{
	return static_cast<QWebView*>(ptr)->QWebView::close();
}

void QWebView_CloseEvent(void* ptr, void* event)
{
	static_cast<QWebView*>(ptr)->closeEvent(static_cast<QCloseEvent*>(event));
}

void QWebView_CloseEventDefault(void* ptr, void* event)
{
	static_cast<QWebView*>(ptr)->QWebView::closeEvent(static_cast<QCloseEvent*>(event));
}

int QWebView_HasHeightForWidth(void* ptr)
{
	return static_cast<QWebView*>(ptr)->hasHeightForWidth();
}

int QWebView_HasHeightForWidthDefault(void* ptr)
{
	return static_cast<QWebView*>(ptr)->QWebView::hasHeightForWidth();
}

int QWebView_HeightForWidth(void* ptr, int w)
{
	return static_cast<QWebView*>(ptr)->heightForWidth(w);
}

int QWebView_HeightForWidthDefault(void* ptr, int w)
{
	return static_cast<QWebView*>(ptr)->QWebView::heightForWidth(w);
}

void QWebView_Hide(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QWebView*>(ptr), "hide");
}

void QWebView_HideDefault(void* ptr)
{
	static_cast<QWebView*>(ptr)->QWebView::hide();
}

void QWebView_InitPainter(void* ptr, void* painter)
{
	static_cast<QWebView*>(ptr)->initPainter(static_cast<QPainter*>(painter));
}

void QWebView_InitPainterDefault(void* ptr, void* painter)
{
	static_cast<QWebView*>(ptr)->QWebView::initPainter(static_cast<QPainter*>(painter));
}

void QWebView_Lower(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QWebView*>(ptr), "lower");
}

void QWebView_LowerDefault(void* ptr)
{
	static_cast<QWebView*>(ptr)->QWebView::lower();
}

int QWebView_NativeEvent(void* ptr, char* eventType, void* message, long result)
{
	return static_cast<QWebView*>(ptr)->nativeEvent(QByteArray(eventType), message, &result);
}

int QWebView_NativeEventDefault(void* ptr, char* eventType, void* message, long result)
{
	return static_cast<QWebView*>(ptr)->QWebView::nativeEvent(QByteArray(eventType), message, &result);
}

void QWebView_Raise(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QWebView*>(ptr), "raise");
}

void QWebView_RaiseDefault(void* ptr)
{
	static_cast<QWebView*>(ptr)->QWebView::raise();
}

void QWebView_Repaint(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QWebView*>(ptr), "repaint");
}

void QWebView_RepaintDefault(void* ptr)
{
	static_cast<QWebView*>(ptr)->QWebView::repaint();
}

void QWebView_SetDisabled(void* ptr, int disable)
{
	QMetaObject::invokeMethod(static_cast<QWebView*>(ptr), "setDisabled", Q_ARG(bool, disable != 0));
}

void QWebView_SetDisabledDefault(void* ptr, int disable)
{
	static_cast<QWebView*>(ptr)->QWebView::setDisabled(disable != 0);
}

void QWebView_SetFocus2(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QWebView*>(ptr), "setFocus");
}

void QWebView_SetFocus2Default(void* ptr)
{
	static_cast<QWebView*>(ptr)->QWebView::setFocus();
}

void QWebView_SetHidden(void* ptr, int hidden)
{
	QMetaObject::invokeMethod(static_cast<QWebView*>(ptr), "setHidden", Q_ARG(bool, hidden != 0));
}

void QWebView_SetHiddenDefault(void* ptr, int hidden)
{
	static_cast<QWebView*>(ptr)->QWebView::setHidden(hidden != 0);
}

void QWebView_Show(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QWebView*>(ptr), "show");
}

void QWebView_ShowDefault(void* ptr)
{
	static_cast<QWebView*>(ptr)->QWebView::show();
}

void QWebView_ShowFullScreen(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QWebView*>(ptr), "showFullScreen");
}

void QWebView_ShowFullScreenDefault(void* ptr)
{
	static_cast<QWebView*>(ptr)->QWebView::showFullScreen();
}

void QWebView_ShowMaximized(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QWebView*>(ptr), "showMaximized");
}

void QWebView_ShowMaximizedDefault(void* ptr)
{
	static_cast<QWebView*>(ptr)->QWebView::showMaximized();
}

void QWebView_ShowMinimized(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QWebView*>(ptr), "showMinimized");
}

void QWebView_ShowMinimizedDefault(void* ptr)
{
	static_cast<QWebView*>(ptr)->QWebView::showMinimized();
}

void QWebView_ShowNormal(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QWebView*>(ptr), "showNormal");
}

void QWebView_ShowNormalDefault(void* ptr)
{
	static_cast<QWebView*>(ptr)->QWebView::showNormal();
}

void QWebView_TabletEvent(void* ptr, void* event)
{
	static_cast<QWebView*>(ptr)->tabletEvent(static_cast<QTabletEvent*>(event));
}

void QWebView_TabletEventDefault(void* ptr, void* event)
{
	static_cast<QWebView*>(ptr)->QWebView::tabletEvent(static_cast<QTabletEvent*>(event));
}

void QWebView_Update(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QWebView*>(ptr), "update");
}

void QWebView_UpdateDefault(void* ptr)
{
	static_cast<QWebView*>(ptr)->QWebView::update();
}

void QWebView_UpdateMicroFocus(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QWebView*>(ptr), "updateMicroFocus");
}

void QWebView_UpdateMicroFocusDefault(void* ptr)
{
	static_cast<QWebView*>(ptr)->QWebView::updateMicroFocus();
}

void QWebView_TimerEvent(void* ptr, void* event)
{
	static_cast<QWebView*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QWebView_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QWebView*>(ptr)->QWebView::timerEvent(static_cast<QTimerEvent*>(event));
}

void QWebView_ChildEvent(void* ptr, void* event)
{
	static_cast<QWebView*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QWebView_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QWebView*>(ptr)->QWebView::childEvent(static_cast<QChildEvent*>(event));
}

void QWebView_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QWebView*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWebView_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QWebView*>(ptr)->QWebView::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWebView_CustomEvent(void* ptr, void* event)
{
	static_cast<QWebView*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QWebView_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QWebView*>(ptr)->QWebView::customEvent(static_cast<QEvent*>(event));
}

void QWebView_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QWebView*>(ptr), "deleteLater");
}

void QWebView_DeleteLaterDefault(void* ptr)
{
	static_cast<QWebView*>(ptr)->QWebView::deleteLater();
}

void QWebView_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QWebView*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWebView_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QWebView*>(ptr)->QWebView::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QWebView_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QWebView*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QWebView_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QWebView*>(ptr)->QWebView::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QWebView_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QWebView*>(ptr)->metaObject());
}

void* QWebView_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QWebView*>(ptr)->QWebView::metaObject());
}

