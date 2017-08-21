// +build !minimal

#define protected public
#define private public

#include "webkit.h"
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
#include <QDate>
#include <QDateTime>
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
#include <QGraphicsSceneMoveEvent>
#include <QGraphicsSceneResizeEvent>
#include <QGraphicsSceneWheelEvent>
#include <QGraphicsTransform>
#include <QGraphicsWebView>
#include <QGraphicsWidget>
#include <QHideEvent>
#include <QIcon>
#include <QInputMethod>
#include <QInputMethodEvent>
#include <QKeyEvent>
#include <QLayout>
#include <QList>
#include <QMap>
#include <QMediaPlaylist>
#include <QMediaRecorder>
#include <QMetaMethod>
#include <QMetaObject>
#include <QMouseEvent>
#include <QMoveEvent>
#include <QMultiMap>
#include <QNetworkAccessManager>
#include <QNetworkReply>
#include <QNetworkRequest>
#include <QObject>
#include <QOffscreenSurface>
#include <QPaintDevice>
#include <QPaintDeviceWindow>
#include <QPaintEngine>
#include <QPaintEvent>
#include <QPainter>
#include <QPainterPath>
#include <QPalette>
#include <QPdfWriter>
#include <QPixmap>
#include <QPoint>
#include <QPointF>
#include <QPrinter>
#include <QQuickItem>
#include <QRadioData>
#include <QRect>
#include <QRectF>
#include <QRegion>
#include <QResizeEvent>
#include <QShowEvent>
#include <QSignalSpy>
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
#include <QWindow>

class MyQGraphicsWebView: public QGraphicsWebView
{
public:
	MyQGraphicsWebView(QGraphicsItem *parent = Q_NULLPTR) : QGraphicsWebView(parent) {QGraphicsWebView_QGraphicsWebView_QRegisterMetaType();};
	QVariant itemChange(QGraphicsItem::GraphicsItemChange change, const QVariant & value) { return *static_cast<QVariant*>(callbackQGraphicsWebView_ItemChange(this, change, const_cast<QVariant*>(&value))); };
	bool event(QEvent * event) { return callbackQGraphicsWebView_Event(this, event) != 0; };
	bool focusNextPrevChild(bool next) { return callbackQGraphicsWebView_FocusNextPrevChild(this, next) != 0; };
	bool sceneEvent(QEvent * event) { return callbackQGraphicsWebView_SceneEvent(this, event) != 0; };
	void back() { callbackQGraphicsWebView_Back(this); };
	void contextMenuEvent(QGraphicsSceneContextMenuEvent * ev) { callbackQGraphicsWebView_ContextMenuEvent(this, ev); };
	void dragEnterEvent(QGraphicsSceneDragDropEvent * ev) { callbackQGraphicsWebView_DragEnterEvent(this, ev); };
	void dragLeaveEvent(QGraphicsSceneDragDropEvent * ev) { callbackQGraphicsWebView_DragLeaveEvent(this, ev); };
	void dragMoveEvent(QGraphicsSceneDragDropEvent * ev) { callbackQGraphicsWebView_DragMoveEvent(this, ev); };
	void dropEvent(QGraphicsSceneDragDropEvent * ev) { callbackQGraphicsWebView_DropEvent(this, ev); };
	void focusInEvent(QFocusEvent * ev) { callbackQGraphicsWebView_FocusInEvent(this, ev); };
	void focusOutEvent(QFocusEvent * ev) { callbackQGraphicsWebView_FocusOutEvent(this, ev); };
	void forward() { callbackQGraphicsWebView_Forward(this); };
	void hoverLeaveEvent(QGraphicsSceneHoverEvent * ev) { callbackQGraphicsWebView_HoverLeaveEvent(this, ev); };
	void hoverMoveEvent(QGraphicsSceneHoverEvent * ev) { callbackQGraphicsWebView_HoverMoveEvent(this, ev); };
	void Signal_IconChanged() { callbackQGraphicsWebView_IconChanged(this); };
	void inputMethodEvent(QInputMethodEvent * ev) { callbackQGraphicsWebView_InputMethodEvent(this, ev); };
	void keyPressEvent(QKeyEvent * ev) { callbackQGraphicsWebView_KeyPressEvent(this, ev); };
	void keyReleaseEvent(QKeyEvent * ev) { callbackQGraphicsWebView_KeyReleaseEvent(this, ev); };
	void Signal_LinkClicked(const QUrl & url) { callbackQGraphicsWebView_LinkClicked(this, const_cast<QUrl*>(&url)); };
	void Signal_LoadFinished(bool ok) { callbackQGraphicsWebView_LoadFinished(this, ok); };
	void Signal_LoadProgress(int progress) { callbackQGraphicsWebView_LoadProgress(this, progress); };
	void Signal_LoadStarted() { callbackQGraphicsWebView_LoadStarted(this); };
	void mouseDoubleClickEvent(QGraphicsSceneMouseEvent * ev) { callbackQGraphicsWebView_MouseDoubleClickEvent(this, ev); };
	void mouseMoveEvent(QGraphicsSceneMouseEvent * ev) { callbackQGraphicsWebView_MouseMoveEvent(this, ev); };
	void mousePressEvent(QGraphicsSceneMouseEvent * ev) { callbackQGraphicsWebView_MousePressEvent(this, ev); };
	void mouseReleaseEvent(QGraphicsSceneMouseEvent * ev) { callbackQGraphicsWebView_MouseReleaseEvent(this, ev); };
	void paint(QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget) { callbackQGraphicsWebView_Paint(this, painter, const_cast<QStyleOptionGraphicsItem*>(option), widget); };
	void reload() { callbackQGraphicsWebView_Reload(this); };
	void Signal_StatusBarMessage(const QString & text) { QByteArray t372ea0 = text.toUtf8(); QtWebKit_PackedString textPacked = { const_cast<char*>(t372ea0.prepend("WHITESPACE").constData()+10), t372ea0.size()-10 };callbackQGraphicsWebView_StatusBarMessage(this, textPacked); };
	void stop() { callbackQGraphicsWebView_Stop(this); };
	void Signal_TitleChanged(const QString & vqs) { QByteArray tda39a3 = vqs.toUtf8(); QtWebKit_PackedString vqsPacked = { const_cast<char*>(tda39a3.prepend("WHITESPACE").constData()+10), tda39a3.size()-10 };callbackQGraphicsWebView_TitleChanged(this, vqsPacked); };
	void updateGeometry() { callbackQGraphicsWebView_UpdateGeometry(this); };
	void Signal_UrlChanged(const QUrl & vqu) { callbackQGraphicsWebView_UrlChanged(this, const_cast<QUrl*>(&vqu)); };
	void wheelEvent(QGraphicsSceneWheelEvent * ev) { callbackQGraphicsWebView_WheelEvent(this, ev); };
	QSizeF sizeHint(Qt::SizeHint which, const QSizeF & constraint) const { return *static_cast<QSizeF*>(callbackQGraphicsWebView_SizeHint(const_cast<void*>(static_cast<const void*>(this)), which, const_cast<QSizeF*>(&constraint))); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackQGraphicsWebView_InputMethodQuery(const_cast<void*>(static_cast<const void*>(this)), query)); };
	bool close() { return callbackQGraphicsWebView_Close(this) != 0; };
	bool windowFrameEvent(QEvent * event) { return callbackQGraphicsWebView_WindowFrameEvent(this, event) != 0; };
	void changeEvent(QEvent * event) { callbackQGraphicsWebView_ChangeEvent(this, event); };
	void closeEvent(QCloseEvent * event) { callbackQGraphicsWebView_CloseEvent(this, event); };
	void Signal_GeometryChanged() { callbackQGraphicsWebView_GeometryChanged(this); };
	void grabKeyboardEvent(QEvent * event) { callbackQGraphicsWebView_GrabKeyboardEvent(this, event); };
	void grabMouseEvent(QEvent * event) { callbackQGraphicsWebView_GrabMouseEvent(this, event); };
	void hideEvent(QHideEvent * event) { callbackQGraphicsWebView_HideEvent(this, event); };
	void moveEvent(QGraphicsSceneMoveEvent * event) { callbackQGraphicsWebView_MoveEvent(this, event); };
	void paintWindowFrame(QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget) { callbackQGraphicsWebView_PaintWindowFrame(this, painter, const_cast<QStyleOptionGraphicsItem*>(option), widget); };
	void polishEvent() { callbackQGraphicsWebView_PolishEvent(this); };
	void resizeEvent(QGraphicsSceneResizeEvent * event) { callbackQGraphicsWebView_ResizeEvent(this, event); };
	void showEvent(QShowEvent * event) { callbackQGraphicsWebView_ShowEvent(this, event); };
	void ungrabKeyboardEvent(QEvent * event) { callbackQGraphicsWebView_UngrabKeyboardEvent(this, event); };
	void ungrabMouseEvent(QEvent * event) { callbackQGraphicsWebView_UngrabMouseEvent(this, event); };
	QPainterPath shape() const { return *static_cast<QPainterPath*>(callbackQGraphicsWebView_Shape(const_cast<void*>(static_cast<const void*>(this)))); };
	QRectF boundingRect() const { return *static_cast<QRectF*>(callbackQGraphicsWebView_BoundingRect(const_cast<void*>(static_cast<const void*>(this)))); };
	Qt::WindowFrameSection windowFrameSectionAt(const QPointF & pos) const { return static_cast<Qt::WindowFrameSection>(callbackQGraphicsWebView_WindowFrameSectionAt(const_cast<void*>(static_cast<const void*>(this)), const_cast<QPointF*>(&pos))); };
	int type() const { return callbackQGraphicsWebView_Type(const_cast<void*>(static_cast<const void*>(this))); };
	void getContentsMargins(qreal * left, qreal * top, qreal * right, qreal * bottom) const { callbackQGraphicsWebView_GetContentsMargins(const_cast<void*>(static_cast<const void*>(this)), *left, *top, *right, *bottom); };
	void initStyleOption(QStyleOption * option) const { callbackQGraphicsWebView_InitStyleOption(const_cast<void*>(static_cast<const void*>(this)), option); };
	void Signal_EnabledChanged() { callbackQGraphicsWebView_EnabledChanged(this); };
	void Signal_OpacityChanged() { callbackQGraphicsWebView_OpacityChanged(this); };
	void Signal_ParentChanged() { callbackQGraphicsWebView_ParentChanged(this); };
	void Signal_RotationChanged() { callbackQGraphicsWebView_RotationChanged(this); };
	void Signal_ScaleChanged() { callbackQGraphicsWebView_ScaleChanged(this); };
	void updateMicroFocus() { callbackQGraphicsWebView_UpdateMicroFocus(this); };
	void Signal_VisibleChanged() { callbackQGraphicsWebView_VisibleChanged(this); };
	void Signal_XChanged() { callbackQGraphicsWebView_XChanged(this); };
	void Signal_YChanged() { callbackQGraphicsWebView_YChanged(this); };
	void Signal_ZChanged() { callbackQGraphicsWebView_ZChanged(this); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQGraphicsWebView_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQGraphicsWebView_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQGraphicsWebView_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQGraphicsWebView_CustomEvent(this, event); };
	void deleteLater() { callbackQGraphicsWebView_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQGraphicsWebView_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQGraphicsWebView_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtWebKit_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQGraphicsWebView_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQGraphicsWebView_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQGraphicsWebView_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	bool sceneEventFilter(QGraphicsItem * watched, QEvent * event) { return callbackQGraphicsWebView_SceneEventFilter(this, watched, event) != 0; };
	void advance(int phase) { callbackQGraphicsWebView_Advance(this, phase); };
	void hoverEnterEvent(QGraphicsSceneHoverEvent * event) { callbackQGraphicsWebView_HoverEnterEvent(this, event); };
	QPainterPath opaqueArea() const { return *static_cast<QPainterPath*>(callbackQGraphicsWebView_OpaqueArea(const_cast<void*>(static_cast<const void*>(this)))); };
	bool collidesWithItem(const QGraphicsItem * other, Qt::ItemSelectionMode mode) const { return callbackQGraphicsWebView_CollidesWithItem(const_cast<void*>(static_cast<const void*>(this)), const_cast<QGraphicsItem*>(other), mode) != 0; };
	bool collidesWithPath(const QPainterPath & path, Qt::ItemSelectionMode mode) const { return callbackQGraphicsWebView_CollidesWithPath(const_cast<void*>(static_cast<const void*>(this)), const_cast<QPainterPath*>(&path), mode) != 0; };
	bool contains(const QPointF & point) const { return callbackQGraphicsWebView_Contains(const_cast<void*>(static_cast<const void*>(this)), const_cast<QPointF*>(&point)) != 0; };
	bool isObscuredBy(const QGraphicsItem * item) const { return callbackQGraphicsWebView_IsObscuredBy(const_cast<void*>(static_cast<const void*>(this)), const_cast<QGraphicsItem*>(item)) != 0; };
};

Q_DECLARE_METATYPE(MyQGraphicsWebView*)

int QGraphicsWebView_QGraphicsWebView_QRegisterMetaType(){qRegisterMetaType<QGraphicsWebView*>(); return qRegisterMetaType<MyQGraphicsWebView*>();}

void* QGraphicsWebView_NewQGraphicsWebView(void* parent)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQGraphicsWebView(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQGraphicsWebView(static_cast<QGraphicsWidget*>(parent));
	} else {
		return new MyQGraphicsWebView(static_cast<QGraphicsItem*>(parent));
	}
}

void* QGraphicsWebView_ItemChangeDefault(void* ptr, long long change, void* value)
{
		return new QVariant(static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::itemChange(static_cast<QGraphicsItem::GraphicsItemChange>(change), *static_cast<QVariant*>(value)));
}

char QGraphicsWebView_EventDefault(void* ptr, void* event)
{
		return static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::event(static_cast<QEvent*>(event));
}

char QGraphicsWebView_FindText(void* ptr, struct QtWebKit_PackedString subString, long long options)
{
	return static_cast<QGraphicsWebView*>(ptr)->findText(QString::fromUtf8(subString.data, subString.len), static_cast<QWebPage::FindFlag>(options));
}

char QGraphicsWebView_FocusNextPrevChildDefault(void* ptr, char next)
{
		return static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::focusNextPrevChild(next != 0);
}

char QGraphicsWebView_SceneEventDefault(void* ptr, void* event)
{
		return static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::sceneEvent(static_cast<QEvent*>(event));
}

void QGraphicsWebView_Back(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QGraphicsWebView*>(ptr), "back");
}

void QGraphicsWebView_BackDefault(void* ptr)
{
		static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::back();
}

void QGraphicsWebView_ContextMenuEventDefault(void* ptr, void* ev)
{
		static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::contextMenuEvent(static_cast<QGraphicsSceneContextMenuEvent*>(ev));
}

void QGraphicsWebView_DragEnterEventDefault(void* ptr, void* ev)
{
		static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::dragEnterEvent(static_cast<QGraphicsSceneDragDropEvent*>(ev));
}

void QGraphicsWebView_DragLeaveEventDefault(void* ptr, void* ev)
{
		static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::dragLeaveEvent(static_cast<QGraphicsSceneDragDropEvent*>(ev));
}

void QGraphicsWebView_DragMoveEventDefault(void* ptr, void* ev)
{
		static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::dragMoveEvent(static_cast<QGraphicsSceneDragDropEvent*>(ev));
}

void QGraphicsWebView_DropEventDefault(void* ptr, void* ev)
{
		static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::dropEvent(static_cast<QGraphicsSceneDragDropEvent*>(ev));
}

void QGraphicsWebView_FocusInEventDefault(void* ptr, void* ev)
{
		static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::focusInEvent(static_cast<QFocusEvent*>(ev));
}

void QGraphicsWebView_FocusOutEventDefault(void* ptr, void* ev)
{
		static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::focusOutEvent(static_cast<QFocusEvent*>(ev));
}

void QGraphicsWebView_Forward(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QGraphicsWebView*>(ptr), "forward");
}

void QGraphicsWebView_ForwardDefault(void* ptr)
{
		static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::forward();
}

void QGraphicsWebView_HoverLeaveEventDefault(void* ptr, void* ev)
{
		static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::hoverLeaveEvent(static_cast<QGraphicsSceneHoverEvent*>(ev));
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

void QGraphicsWebView_InputMethodEventDefault(void* ptr, void* ev)
{
		static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::inputMethodEvent(static_cast<QInputMethodEvent*>(ev));
}

void QGraphicsWebView_KeyPressEventDefault(void* ptr, void* ev)
{
		static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::keyPressEvent(static_cast<QKeyEvent*>(ev));
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

void QGraphicsWebView_Load2(void* ptr, void* request, long long operation, void* body)
{
	static_cast<QGraphicsWebView*>(ptr)->load(*static_cast<QNetworkRequest*>(request), static_cast<QNetworkAccessManager::Operation>(operation), *static_cast<QByteArray*>(body));
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

void QGraphicsWebView_LoadFinished(void* ptr, char ok)
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

void QGraphicsWebView_MouseDoubleClickEventDefault(void* ptr, void* ev)
{
		static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::mouseDoubleClickEvent(static_cast<QGraphicsSceneMouseEvent*>(ev));
}

void QGraphicsWebView_MouseMoveEventDefault(void* ptr, void* ev)
{
		static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::mouseMoveEvent(static_cast<QGraphicsSceneMouseEvent*>(ev));
}

void QGraphicsWebView_MousePressEventDefault(void* ptr, void* ev)
{
		static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::mousePressEvent(static_cast<QGraphicsSceneMouseEvent*>(ev));
}

void QGraphicsWebView_MouseReleaseEventDefault(void* ptr, void* ev)
{
		static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::mouseReleaseEvent(static_cast<QGraphicsSceneMouseEvent*>(ev));
}

void QGraphicsWebView_PaintDefault(void* ptr, void* painter, void* option, void* widget)
{
		static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::paint(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
}

void QGraphicsWebView_Reload(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QGraphicsWebView*>(ptr), "reload");
}

void QGraphicsWebView_ReloadDefault(void* ptr)
{
		static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::reload();
}

void QGraphicsWebView_SetContent(void* ptr, void* data, struct QtWebKit_PackedString mimeType, void* baseUrl)
{
	static_cast<QGraphicsWebView*>(ptr)->setContent(*static_cast<QByteArray*>(data), QString::fromUtf8(mimeType.data, mimeType.len), *static_cast<QUrl*>(baseUrl));
}

void QGraphicsWebView_SetGeometryDefault(void* ptr, void* rect)
{
		static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::setGeometry(*static_cast<QRectF*>(rect));
}

void QGraphicsWebView_SetHtml(void* ptr, struct QtWebKit_PackedString html, void* baseUrl)
{
	static_cast<QGraphicsWebView*>(ptr)->setHtml(QString::fromUtf8(html.data, html.len), *static_cast<QUrl*>(baseUrl));
}

void QGraphicsWebView_SetPage(void* ptr, void* page)
{
	static_cast<QGraphicsWebView*>(ptr)->setPage(static_cast<QWebPage*>(page));
}

void QGraphicsWebView_SetRenderHint(void* ptr, long long hint, char enabled)
{
	static_cast<QGraphicsWebView*>(ptr)->setRenderHint(static_cast<QPainter::RenderHint>(hint), enabled != 0);
}

void QGraphicsWebView_SetRenderHints(void* ptr, long long hints)
{
	static_cast<QGraphicsWebView*>(ptr)->setRenderHints(static_cast<QPainter::RenderHint>(hints));
}

void QGraphicsWebView_SetResizesToContents(void* ptr, char enabled)
{
	static_cast<QGraphicsWebView*>(ptr)->setResizesToContents(enabled != 0);
}

void QGraphicsWebView_SetTiledBackingStoreFrozen(void* ptr, char frozen)
{
	static_cast<QGraphicsWebView*>(ptr)->setTiledBackingStoreFrozen(frozen != 0);
}

void QGraphicsWebView_SetUrl(void* ptr, void* vqu)
{
	static_cast<QGraphicsWebView*>(ptr)->setUrl(*static_cast<QUrl*>(vqu));
}

void QGraphicsWebView_SetZoomFactor(void* ptr, double vqr)
{
	static_cast<QGraphicsWebView*>(ptr)->setZoomFactor(vqr);
}

void QGraphicsWebView_ConnectStatusBarMessage(void* ptr)
{
	QObject::connect(static_cast<QGraphicsWebView*>(ptr), static_cast<void (QGraphicsWebView::*)(const QString &)>(&QGraphicsWebView::statusBarMessage), static_cast<MyQGraphicsWebView*>(ptr), static_cast<void (MyQGraphicsWebView::*)(const QString &)>(&MyQGraphicsWebView::Signal_StatusBarMessage));
}

void QGraphicsWebView_DisconnectStatusBarMessage(void* ptr)
{
	QObject::disconnect(static_cast<QGraphicsWebView*>(ptr), static_cast<void (QGraphicsWebView::*)(const QString &)>(&QGraphicsWebView::statusBarMessage), static_cast<MyQGraphicsWebView*>(ptr), static_cast<void (MyQGraphicsWebView::*)(const QString &)>(&MyQGraphicsWebView::Signal_StatusBarMessage));
}

void QGraphicsWebView_StatusBarMessage(void* ptr, struct QtWebKit_PackedString text)
{
	static_cast<QGraphicsWebView*>(ptr)->statusBarMessage(QString::fromUtf8(text.data, text.len));
}

void QGraphicsWebView_Stop(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QGraphicsWebView*>(ptr), "stop");
}

void QGraphicsWebView_StopDefault(void* ptr)
{
		static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::stop();
}

void QGraphicsWebView_ConnectTitleChanged(void* ptr)
{
	QObject::connect(static_cast<QGraphicsWebView*>(ptr), static_cast<void (QGraphicsWebView::*)(const QString &)>(&QGraphicsWebView::titleChanged), static_cast<MyQGraphicsWebView*>(ptr), static_cast<void (MyQGraphicsWebView::*)(const QString &)>(&MyQGraphicsWebView::Signal_TitleChanged));
}

void QGraphicsWebView_DisconnectTitleChanged(void* ptr)
{
	QObject::disconnect(static_cast<QGraphicsWebView*>(ptr), static_cast<void (QGraphicsWebView::*)(const QString &)>(&QGraphicsWebView::titleChanged), static_cast<MyQGraphicsWebView*>(ptr), static_cast<void (MyQGraphicsWebView::*)(const QString &)>(&MyQGraphicsWebView::Signal_TitleChanged));
}

void QGraphicsWebView_TitleChanged(void* ptr, struct QtWebKit_PackedString vqs)
{
	static_cast<QGraphicsWebView*>(ptr)->titleChanged(QString::fromUtf8(vqs.data, vqs.len));
}

void QGraphicsWebView_TriggerPageAction(void* ptr, long long action, char checked)
{
	static_cast<QGraphicsWebView*>(ptr)->triggerPageAction(static_cast<QWebPage::WebAction>(action), checked != 0);
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

void QGraphicsWebView_WheelEventDefault(void* ptr, void* ev)
{
		static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::wheelEvent(static_cast<QGraphicsSceneWheelEvent*>(ev));
}

void QGraphicsWebView_DestroyQGraphicsWebView(void* ptr)
{
	static_cast<QGraphicsWebView*>(ptr)->~QGraphicsWebView();
}

void* QGraphicsWebView_PageAction(void* ptr, long long action)
{
	return static_cast<QGraphicsWebView*>(ptr)->pageAction(static_cast<QWebPage::WebAction>(action));
}

void* QGraphicsWebView_Icon(void* ptr)
{
	return new QIcon(static_cast<QGraphicsWebView*>(ptr)->icon());
}

long long QGraphicsWebView_RenderHints(void* ptr)
{
	return static_cast<QGraphicsWebView*>(ptr)->renderHints();
}

void* QGraphicsWebView_SizeHintDefault(void* ptr, long long which, void* constraint)
{
		return ({ QSizeF tmpValue = static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::sizeHint(static_cast<Qt::SizeHint>(which), *static_cast<QSizeF*>(constraint)); new QSizeF(tmpValue.width(), tmpValue.height()); });
}

struct QtWebKit_PackedString QGraphicsWebView_Title(void* ptr)
{
	return ({ QByteArray t3650e0 = static_cast<QGraphicsWebView*>(ptr)->title().toUtf8(); QtWebKit_PackedString { const_cast<char*>(t3650e0.prepend("WHITESPACE").constData()+10), t3650e0.size()-10 }; });
}

void* QGraphicsWebView_Url(void* ptr)
{
	return new QUrl(static_cast<QGraphicsWebView*>(ptr)->url());
}

void* QGraphicsWebView_InputMethodQueryDefault(void* ptr, long long query)
{
		return new QVariant(static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

void* QGraphicsWebView_History(void* ptr)
{
	return static_cast<QGraphicsWebView*>(ptr)->history();
}

void* QGraphicsWebView_Page(void* ptr)
{
	return static_cast<QGraphicsWebView*>(ptr)->page();
}

void* QGraphicsWebView_Settings(void* ptr)
{
	return static_cast<QGraphicsWebView*>(ptr)->settings();
}

char QGraphicsWebView_IsModified(void* ptr)
{
	return static_cast<QGraphicsWebView*>(ptr)->isModified();
}

char QGraphicsWebView_IsTiledBackingStoreFrozen(void* ptr)
{
	return static_cast<QGraphicsWebView*>(ptr)->isTiledBackingStoreFrozen();
}

char QGraphicsWebView_ResizesToContents(void* ptr)
{
	return static_cast<QGraphicsWebView*>(ptr)->resizesToContents();
}

double QGraphicsWebView_ZoomFactor(void* ptr)
{
	return static_cast<QGraphicsWebView*>(ptr)->zoomFactor();
}

void* QGraphicsWebView___addActions_actions_atList(void* ptr, int i)
{
	return const_cast<QAction*>(static_cast<QList<QAction *>*>(ptr)->at(i));
}

void QGraphicsWebView___addActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QGraphicsWebView___addActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>;
}

void* QGraphicsWebView___insertActions_actions_atList(void* ptr, int i)
{
	return const_cast<QAction*>(static_cast<QList<QAction *>*>(ptr)->at(i));
}

void QGraphicsWebView___insertActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QGraphicsWebView___insertActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>;
}

void* QGraphicsWebView___actions_atList(void* ptr, int i)
{
	return const_cast<QAction*>(static_cast<QList<QAction *>*>(ptr)->at(i));
}

void QGraphicsWebView___actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QGraphicsWebView___actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>;
}

void* QGraphicsWebView___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QGraphicsWebView___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QGraphicsWebView___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QGraphicsWebView___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QGraphicsWebView___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGraphicsWebView___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QGraphicsWebView___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QGraphicsWebView___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGraphicsWebView___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QGraphicsWebView___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QGraphicsWebView___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGraphicsWebView___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QGraphicsWebView___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QGraphicsWebView___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGraphicsWebView___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

void* QGraphicsWebView___setTransformations_transformations_atList(void* ptr, int i)
{
	return const_cast<QGraphicsTransform*>(static_cast<QList<QGraphicsTransform *>*>(ptr)->at(i));
}

void QGraphicsWebView___setTransformations_transformations_setList(void* ptr, void* i)
{
	static_cast<QList<QGraphicsTransform *>*>(ptr)->append(static_cast<QGraphicsTransform*>(i));
}

void* QGraphicsWebView___setTransformations_transformations_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QGraphicsTransform *>;
}

void* QGraphicsWebView___childItems_atList(void* ptr, int i)
{
	return const_cast<QGraphicsItem*>(static_cast<QList<QGraphicsItem *>*>(ptr)->at(i));
}

void QGraphicsWebView___childItems_setList(void* ptr, void* i)
{
	static_cast<QList<QGraphicsItem *>*>(ptr)->append(static_cast<QGraphicsItem*>(i));
}

void* QGraphicsWebView___childItems_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QGraphicsItem *>;
}

void* QGraphicsWebView___collidingItems_atList(void* ptr, int i)
{
	return const_cast<QGraphicsItem*>(static_cast<QList<QGraphicsItem *>*>(ptr)->at(i));
}

void QGraphicsWebView___collidingItems_setList(void* ptr, void* i)
{
	static_cast<QList<QGraphicsItem *>*>(ptr)->append(static_cast<QGraphicsItem*>(i));
}

void* QGraphicsWebView___collidingItems_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QGraphicsItem *>;
}

void* QGraphicsWebView___transformations_atList(void* ptr, int i)
{
	return const_cast<QGraphicsTransform*>(static_cast<QList<QGraphicsTransform *>*>(ptr)->at(i));
}

void QGraphicsWebView___transformations_setList(void* ptr, void* i)
{
	static_cast<QList<QGraphicsTransform *>*>(ptr)->append(static_cast<QGraphicsTransform*>(i));
}

void* QGraphicsWebView___transformations_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QGraphicsTransform *>;
}

char QGraphicsWebView_CloseDefault(void* ptr)
{
		return static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::close();
}

char QGraphicsWebView_WindowFrameEventDefault(void* ptr, void* event)
{
		return static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::windowFrameEvent(static_cast<QEvent*>(event));
}

void QGraphicsWebView_ChangeEventDefault(void* ptr, void* event)
{
		static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::changeEvent(static_cast<QEvent*>(event));
}

void QGraphicsWebView_CloseEventDefault(void* ptr, void* event)
{
		static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::closeEvent(static_cast<QCloseEvent*>(event));
}

void QGraphicsWebView_GrabKeyboardEventDefault(void* ptr, void* event)
{
		static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::grabKeyboardEvent(static_cast<QEvent*>(event));
}

void QGraphicsWebView_GrabMouseEventDefault(void* ptr, void* event)
{
		static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::grabMouseEvent(static_cast<QEvent*>(event));
}

void QGraphicsWebView_HideEventDefault(void* ptr, void* event)
{
		static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::hideEvent(static_cast<QHideEvent*>(event));
}

void QGraphicsWebView_MoveEventDefault(void* ptr, void* event)
{
		static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::moveEvent(static_cast<QGraphicsSceneMoveEvent*>(event));
}

void QGraphicsWebView_PaintWindowFrameDefault(void* ptr, void* painter, void* option, void* widget)
{
		static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::paintWindowFrame(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
}

void QGraphicsWebView_PolishEventDefault(void* ptr)
{
		static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::polishEvent();
}

void QGraphicsWebView_ResizeEventDefault(void* ptr, void* event)
{
		static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::resizeEvent(static_cast<QGraphicsSceneResizeEvent*>(event));
}

void QGraphicsWebView_ShowEventDefault(void* ptr, void* event)
{
		static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::showEvent(static_cast<QShowEvent*>(event));
}

void QGraphicsWebView_UngrabKeyboardEventDefault(void* ptr, void* event)
{
		static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::ungrabKeyboardEvent(static_cast<QEvent*>(event));
}

void QGraphicsWebView_UngrabMouseEventDefault(void* ptr, void* event)
{
		static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::ungrabMouseEvent(static_cast<QEvent*>(event));
}

void* QGraphicsWebView_ShapeDefault(void* ptr)
{
		return new QPainterPath(static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::shape());
}

void* QGraphicsWebView_BoundingRectDefault(void* ptr)
{
		return ({ QRectF tmpValue = static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::boundingRect(); new QRectF(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

long long QGraphicsWebView_WindowFrameSectionAtDefault(void* ptr, void* pos)
{
		return static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::windowFrameSectionAt(*static_cast<QPointF*>(pos));
}

int QGraphicsWebView_TypeDefault(void* ptr)
{
		return static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::type();
}

void QGraphicsWebView_GetContentsMarginsDefault(void* ptr, double left, double top, double right, double bottom)
{
		static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::getContentsMargins(&left, &top, &right, &bottom);
}

void QGraphicsWebView_InitStyleOptionDefault(void* ptr, void* option)
{
		static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::initStyleOption(static_cast<QStyleOption*>(option));
}

void QGraphicsWebView_UpdateMicroFocusDefault(void* ptr)
{
		static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::updateMicroFocus();
}

char QGraphicsWebView_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QGraphicsWebView_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::childEvent(static_cast<QChildEvent*>(event));
}

void QGraphicsWebView_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGraphicsWebView_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::customEvent(static_cast<QEvent*>(event));
}

void QGraphicsWebView_DeleteLaterDefault(void* ptr)
{
		static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::deleteLater();
}

void QGraphicsWebView_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGraphicsWebView_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QGraphicsWebView_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::metaObject());
}

char QGraphicsWebView_SceneEventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::sceneEventFilter(static_cast<QGraphicsItem*>(watched), static_cast<QEvent*>(event));
}

void QGraphicsWebView_AdvanceDefault(void* ptr, int phase)
{
		static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::advance(phase);
}

void QGraphicsWebView_HoverEnterEventDefault(void* ptr, void* event)
{
		static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::hoverEnterEvent(static_cast<QGraphicsSceneHoverEvent*>(event));
}

void* QGraphicsWebView_OpaqueAreaDefault(void* ptr)
{
		return new QPainterPath(static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::opaqueArea());
}

char QGraphicsWebView_CollidesWithItemDefault(void* ptr, void* other, long long mode)
{
		return static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::collidesWithItem(static_cast<QGraphicsItem*>(other), static_cast<Qt::ItemSelectionMode>(mode));
}

char QGraphicsWebView_CollidesWithPathDefault(void* ptr, void* path, long long mode)
{
		return static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::collidesWithPath(*static_cast<QPainterPath*>(path), static_cast<Qt::ItemSelectionMode>(mode));
}

char QGraphicsWebView_ContainsDefault(void* ptr, void* point)
{
		return static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::contains(*static_cast<QPointF*>(point));
}

char QGraphicsWebView_IsObscuredByDefault(void* ptr, void* item)
{
		return static_cast<QGraphicsWebView*>(ptr)->QGraphicsWebView::isObscuredBy(static_cast<QGraphicsItem*>(item));
}

void* QWebDatabase_NewQWebDatabase(void* other)
{
	return new QWebDatabase(*static_cast<QWebDatabase*>(other));
}

void QWebDatabase_QWebDatabase_RemoveAllDatabases()
{
	QWebDatabase::removeAllDatabases();
}

void QWebDatabase_QWebDatabase_RemoveDatabase(void* db)
{
	QWebDatabase::removeDatabase(*static_cast<QWebDatabase*>(db));
}

void QWebDatabase_DestroyQWebDatabase(void* ptr)
{
	static_cast<QWebDatabase*>(ptr)->~QWebDatabase();
}

struct QtWebKit_PackedString QWebDatabase_DisplayName(void* ptr)
{
	return ({ QByteArray t70b35e = static_cast<QWebDatabase*>(ptr)->displayName().toUtf8(); QtWebKit_PackedString { const_cast<char*>(t70b35e.prepend("WHITESPACE").constData()+10), t70b35e.size()-10 }; });
}

struct QtWebKit_PackedString QWebDatabase_FileName(void* ptr)
{
	return ({ QByteArray t6a9a5c = static_cast<QWebDatabase*>(ptr)->fileName().toUtf8(); QtWebKit_PackedString { const_cast<char*>(t6a9a5c.prepend("WHITESPACE").constData()+10), t6a9a5c.size()-10 }; });
}

struct QtWebKit_PackedString QWebDatabase_Name(void* ptr)
{
	return ({ QByteArray t1fba05 = static_cast<QWebDatabase*>(ptr)->name().toUtf8(); QtWebKit_PackedString { const_cast<char*>(t1fba05.prepend("WHITESPACE").constData()+10), t1fba05.size()-10 }; });
}

void* QWebDatabase_Origin(void* ptr)
{
	return new QWebSecurityOrigin(static_cast<QWebDatabase*>(ptr)->origin());
}

long long QWebDatabase_ExpectedSize(void* ptr)
{
	return static_cast<QWebDatabase*>(ptr)->expectedSize();
}

long long QWebDatabase_Size(void* ptr)
{
	return static_cast<QWebDatabase*>(ptr)->size();
}

void* QWebElement_EvaluateJavaScript(void* ptr, struct QtWebKit_PackedString scriptSource)
{
	return new QVariant(static_cast<QWebElement*>(ptr)->evaluateJavaScript(QString::fromUtf8(scriptSource.data, scriptSource.len)));
}

void* QWebElement_TakeFromDocument(void* ptr)
{
	return new QWebElement(static_cast<QWebElement*>(ptr)->takeFromDocument());
}

void* QWebElement_NewQWebElement()
{
	return new QWebElement();
}

void* QWebElement_NewQWebElement2(void* other)
{
	return new QWebElement(*static_cast<QWebElement*>(other));
}

void QWebElement_AddClass(void* ptr, struct QtWebKit_PackedString name)
{
	static_cast<QWebElement*>(ptr)->addClass(QString::fromUtf8(name.data, name.len));
}

void QWebElement_AppendInside(void* ptr, struct QtWebKit_PackedString markup)
{
	static_cast<QWebElement*>(ptr)->appendInside(QString::fromUtf8(markup.data, markup.len));
}

void QWebElement_AppendInside2(void* ptr, void* element)
{
	static_cast<QWebElement*>(ptr)->appendInside(*static_cast<QWebElement*>(element));
}

void QWebElement_AppendOutside(void* ptr, struct QtWebKit_PackedString markup)
{
	static_cast<QWebElement*>(ptr)->appendOutside(QString::fromUtf8(markup.data, markup.len));
}

void QWebElement_AppendOutside2(void* ptr, void* element)
{
	static_cast<QWebElement*>(ptr)->appendOutside(*static_cast<QWebElement*>(element));
}

void QWebElement_EncloseContentsWith2(void* ptr, struct QtWebKit_PackedString markup)
{
	static_cast<QWebElement*>(ptr)->encloseContentsWith(QString::fromUtf8(markup.data, markup.len));
}

void QWebElement_EncloseContentsWith(void* ptr, void* element)
{
	static_cast<QWebElement*>(ptr)->encloseContentsWith(*static_cast<QWebElement*>(element));
}

void QWebElement_EncloseWith(void* ptr, struct QtWebKit_PackedString markup)
{
	static_cast<QWebElement*>(ptr)->encloseWith(QString::fromUtf8(markup.data, markup.len));
}

void QWebElement_EncloseWith2(void* ptr, void* element)
{
	static_cast<QWebElement*>(ptr)->encloseWith(*static_cast<QWebElement*>(element));
}

void QWebElement_PrependInside(void* ptr, struct QtWebKit_PackedString markup)
{
	static_cast<QWebElement*>(ptr)->prependInside(QString::fromUtf8(markup.data, markup.len));
}

void QWebElement_PrependInside2(void* ptr, void* element)
{
	static_cast<QWebElement*>(ptr)->prependInside(*static_cast<QWebElement*>(element));
}

void QWebElement_PrependOutside(void* ptr, struct QtWebKit_PackedString markup)
{
	static_cast<QWebElement*>(ptr)->prependOutside(QString::fromUtf8(markup.data, markup.len));
}

void QWebElement_PrependOutside2(void* ptr, void* element)
{
	static_cast<QWebElement*>(ptr)->prependOutside(*static_cast<QWebElement*>(element));
}

void QWebElement_RemoveAllChildren(void* ptr)
{
	static_cast<QWebElement*>(ptr)->removeAllChildren();
}

void QWebElement_RemoveAttribute(void* ptr, struct QtWebKit_PackedString name)
{
	static_cast<QWebElement*>(ptr)->removeAttribute(QString::fromUtf8(name.data, name.len));
}

void QWebElement_RemoveAttributeNS(void* ptr, struct QtWebKit_PackedString namespaceUri, struct QtWebKit_PackedString name)
{
	static_cast<QWebElement*>(ptr)->removeAttributeNS(QString::fromUtf8(namespaceUri.data, namespaceUri.len), QString::fromUtf8(name.data, name.len));
}

void QWebElement_RemoveClass(void* ptr, struct QtWebKit_PackedString name)
{
	static_cast<QWebElement*>(ptr)->removeClass(QString::fromUtf8(name.data, name.len));
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

void QWebElement_Replace(void* ptr, struct QtWebKit_PackedString markup)
{
	static_cast<QWebElement*>(ptr)->replace(QString::fromUtf8(markup.data, markup.len));
}

void QWebElement_Replace2(void* ptr, void* element)
{
	static_cast<QWebElement*>(ptr)->replace(*static_cast<QWebElement*>(element));
}

void QWebElement_SetAttribute(void* ptr, struct QtWebKit_PackedString name, struct QtWebKit_PackedString value)
{
	static_cast<QWebElement*>(ptr)->setAttribute(QString::fromUtf8(name.data, name.len), QString::fromUtf8(value.data, value.len));
}

void QWebElement_SetAttributeNS(void* ptr, struct QtWebKit_PackedString namespaceUri, struct QtWebKit_PackedString name, struct QtWebKit_PackedString value)
{
	static_cast<QWebElement*>(ptr)->setAttributeNS(QString::fromUtf8(namespaceUri.data, namespaceUri.len), QString::fromUtf8(name.data, name.len), QString::fromUtf8(value.data, value.len));
}

void QWebElement_SetFocus(void* ptr)
{
	static_cast<QWebElement*>(ptr)->setFocus();
}

void QWebElement_SetInnerXml(void* ptr, struct QtWebKit_PackedString markup)
{
	static_cast<QWebElement*>(ptr)->setInnerXml(QString::fromUtf8(markup.data, markup.len));
}

void QWebElement_SetOuterXml(void* ptr, struct QtWebKit_PackedString markup)
{
	static_cast<QWebElement*>(ptr)->setOuterXml(QString::fromUtf8(markup.data, markup.len));
}

void QWebElement_SetPlainText(void* ptr, struct QtWebKit_PackedString text)
{
	static_cast<QWebElement*>(ptr)->setPlainText(QString::fromUtf8(text.data, text.len));
}

void QWebElement_SetStyleProperty(void* ptr, struct QtWebKit_PackedString name, struct QtWebKit_PackedString value)
{
	static_cast<QWebElement*>(ptr)->setStyleProperty(QString::fromUtf8(name.data, name.len), QString::fromUtf8(value.data, value.len));
}

void QWebElement_ToggleClass(void* ptr, struct QtWebKit_PackedString name)
{
	static_cast<QWebElement*>(ptr)->toggleClass(QString::fromUtf8(name.data, name.len));
}

void QWebElement_DestroyQWebElement(void* ptr)
{
	static_cast<QWebElement*>(ptr)->~QWebElement();
}

void* QWebElement_Geometry(void* ptr)
{
	return ({ QRect tmpValue = static_cast<QWebElement*>(ptr)->geometry(); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

struct QtWebKit_PackedString QWebElement_Attribute(void* ptr, struct QtWebKit_PackedString name, struct QtWebKit_PackedString defaultValue)
{
	return ({ QByteArray t91d24c = static_cast<QWebElement*>(ptr)->attribute(QString::fromUtf8(name.data, name.len), QString::fromUtf8(defaultValue.data, defaultValue.len)).toUtf8(); QtWebKit_PackedString { const_cast<char*>(t91d24c.prepend("WHITESPACE").constData()+10), t91d24c.size()-10 }; });
}

struct QtWebKit_PackedString QWebElement_AttributeNS(void* ptr, struct QtWebKit_PackedString namespaceUri, struct QtWebKit_PackedString name, struct QtWebKit_PackedString defaultValue)
{
	return ({ QByteArray te7d3da = static_cast<QWebElement*>(ptr)->attributeNS(QString::fromUtf8(namespaceUri.data, namespaceUri.len), QString::fromUtf8(name.data, name.len), QString::fromUtf8(defaultValue.data, defaultValue.len)).toUtf8(); QtWebKit_PackedString { const_cast<char*>(te7d3da.prepend("WHITESPACE").constData()+10), te7d3da.size()-10 }; });
}

struct QtWebKit_PackedString QWebElement_LocalName(void* ptr)
{
	return ({ QByteArray t62e192 = static_cast<QWebElement*>(ptr)->localName().toUtf8(); QtWebKit_PackedString { const_cast<char*>(t62e192.prepend("WHITESPACE").constData()+10), t62e192.size()-10 }; });
}

struct QtWebKit_PackedString QWebElement_NamespaceUri(void* ptr)
{
	return ({ QByteArray t6196a3 = static_cast<QWebElement*>(ptr)->namespaceUri().toUtf8(); QtWebKit_PackedString { const_cast<char*>(t6196a3.prepend("WHITESPACE").constData()+10), t6196a3.size()-10 }; });
}

struct QtWebKit_PackedString QWebElement_Prefix(void* ptr)
{
	return ({ QByteArray t3b50c5 = static_cast<QWebElement*>(ptr)->prefix().toUtf8(); QtWebKit_PackedString { const_cast<char*>(t3b50c5.prepend("WHITESPACE").constData()+10), t3b50c5.size()-10 }; });
}

struct QtWebKit_PackedString QWebElement_StyleProperty(void* ptr, struct QtWebKit_PackedString name, long long strategy)
{
	return ({ QByteArray t787493 = static_cast<QWebElement*>(ptr)->styleProperty(QString::fromUtf8(name.data, name.len), static_cast<QWebElement::StyleResolveStrategy>(strategy)).toUtf8(); QtWebKit_PackedString { const_cast<char*>(t787493.prepend("WHITESPACE").constData()+10), t787493.size()-10 }; });
}

struct QtWebKit_PackedString QWebElement_TagName(void* ptr)
{
	return ({ QByteArray t85055f = static_cast<QWebElement*>(ptr)->tagName().toUtf8(); QtWebKit_PackedString { const_cast<char*>(t85055f.prepend("WHITESPACE").constData()+10), t85055f.size()-10 }; });
}

struct QtWebKit_PackedString QWebElement_ToInnerXml(void* ptr)
{
	return ({ QByteArray tbed539 = static_cast<QWebElement*>(ptr)->toInnerXml().toUtf8(); QtWebKit_PackedString { const_cast<char*>(tbed539.prepend("WHITESPACE").constData()+10), tbed539.size()-10 }; });
}

struct QtWebKit_PackedString QWebElement_ToOuterXml(void* ptr)
{
	return ({ QByteArray td1ea25 = static_cast<QWebElement*>(ptr)->toOuterXml().toUtf8(); QtWebKit_PackedString { const_cast<char*>(td1ea25.prepend("WHITESPACE").constData()+10), td1ea25.size()-10 }; });
}

struct QtWebKit_PackedString QWebElement_ToPlainText(void* ptr)
{
	return ({ QByteArray tf7c87e = static_cast<QWebElement*>(ptr)->toPlainText().toUtf8(); QtWebKit_PackedString { const_cast<char*>(tf7c87e.prepend("WHITESPACE").constData()+10), tf7c87e.size()-10 }; });
}

struct QtWebKit_PackedString QWebElement_AttributeNames(void* ptr, struct QtWebKit_PackedString namespaceUri)
{
	return ({ QByteArray t08a8f8 = static_cast<QWebElement*>(ptr)->attributeNames(QString::fromUtf8(namespaceUri.data, namespaceUri.len)).join("|").toUtf8(); QtWebKit_PackedString { const_cast<char*>(t08a8f8.prepend("WHITESPACE").constData()+10), t08a8f8.size()-10 }; });
}

struct QtWebKit_PackedString QWebElement_Classes(void* ptr)
{
	return ({ QByteArray t46b13e = static_cast<QWebElement*>(ptr)->classes().join("|").toUtf8(); QtWebKit_PackedString { const_cast<char*>(t46b13e.prepend("WHITESPACE").constData()+10), t46b13e.size()-10 }; });
}

void* QWebElement_Clone(void* ptr)
{
	return new QWebElement(static_cast<QWebElement*>(ptr)->clone());
}

void* QWebElement_Document(void* ptr)
{
	return new QWebElement(static_cast<QWebElement*>(ptr)->document());
}

void* QWebElement_FindFirst(void* ptr, struct QtWebKit_PackedString selectorQuery)
{
	return new QWebElement(static_cast<QWebElement*>(ptr)->findFirst(QString::fromUtf8(selectorQuery.data, selectorQuery.len)));
}

void* QWebElement_FirstChild(void* ptr)
{
	return new QWebElement(static_cast<QWebElement*>(ptr)->firstChild());
}

void* QWebElement_LastChild(void* ptr)
{
	return new QWebElement(static_cast<QWebElement*>(ptr)->lastChild());
}

void* QWebElement_NextSibling(void* ptr)
{
	return new QWebElement(static_cast<QWebElement*>(ptr)->nextSibling());
}

void* QWebElement_Parent(void* ptr)
{
	return new QWebElement(static_cast<QWebElement*>(ptr)->parent());
}

void* QWebElement_PreviousSibling(void* ptr)
{
	return new QWebElement(static_cast<QWebElement*>(ptr)->previousSibling());
}

void* QWebElement_FindAll(void* ptr, struct QtWebKit_PackedString selectorQuery)
{
	return new QWebElementCollection(static_cast<QWebElement*>(ptr)->findAll(QString::fromUtf8(selectorQuery.data, selectorQuery.len)));
}

void* QWebElement_WebFrame(void* ptr)
{
	return static_cast<QWebElement*>(ptr)->webFrame();
}

char QWebElement_HasAttribute(void* ptr, struct QtWebKit_PackedString name)
{
	return static_cast<QWebElement*>(ptr)->hasAttribute(QString::fromUtf8(name.data, name.len));
}

char QWebElement_HasAttributeNS(void* ptr, struct QtWebKit_PackedString namespaceUri, struct QtWebKit_PackedString name)
{
	return static_cast<QWebElement*>(ptr)->hasAttributeNS(QString::fromUtf8(namespaceUri.data, namespaceUri.len), QString::fromUtf8(name.data, name.len));
}

char QWebElement_HasAttributes(void* ptr)
{
	return static_cast<QWebElement*>(ptr)->hasAttributes();
}

char QWebElement_HasClass(void* ptr, struct QtWebKit_PackedString name)
{
	return static_cast<QWebElement*>(ptr)->hasClass(QString::fromUtf8(name.data, name.len));
}

char QWebElement_HasFocus(void* ptr)
{
	return static_cast<QWebElement*>(ptr)->hasFocus();
}

char QWebElement_IsNull(void* ptr)
{
	return static_cast<QWebElement*>(ptr)->isNull();
}

void* QWebElementCollection_NewQWebElementCollection()
{
	return new QWebElementCollection();
}

void* QWebElementCollection_NewQWebElementCollection2(void* contextElement, struct QtWebKit_PackedString query)
{
	return new QWebElementCollection(*static_cast<QWebElement*>(contextElement), QString::fromUtf8(query.data, query.len));
}

void* QWebElementCollection_NewQWebElementCollection3(void* other)
{
	return new QWebElementCollection(*static_cast<QWebElementCollection*>(other));
}

void QWebElementCollection_Append(void* ptr, void* other)
{
	static_cast<QWebElementCollection*>(ptr)->append(*static_cast<QWebElementCollection*>(other));
}

void QWebElementCollection_DestroyQWebElementCollection(void* ptr)
{
	static_cast<QWebElementCollection*>(ptr)->~QWebElementCollection();
}

struct QtWebKit_PackedList QWebElementCollection_ToList(void* ptr)
{
	return ({ QList<QWebElement>* tmpValue = new QList<QWebElement>(static_cast<QWebElementCollection*>(ptr)->toList()); QtWebKit_PackedList { tmpValue, tmpValue->size() }; });
}

void* QWebElementCollection_At(void* ptr, int i)
{
	return new QWebElement(static_cast<QWebElementCollection*>(ptr)->at(i));
}

void* QWebElementCollection_First(void* ptr)
{
	return new QWebElement(static_cast<QWebElementCollection*>(ptr)->first());
}

void* QWebElementCollection_Last(void* ptr)
{
	return new QWebElement(static_cast<QWebElementCollection*>(ptr)->last());
}

int QWebElementCollection_Count(void* ptr)
{
	return static_cast<QWebElementCollection*>(ptr)->count();
}

void* QWebElementCollection___toList_atList(void* ptr, int i)
{
	return new QWebElement(static_cast<QList<QWebElement>*>(ptr)->at(i));
}

void QWebElementCollection___toList_setList(void* ptr, void* i)
{
	static_cast<QList<QWebElement>*>(ptr)->append(*static_cast<QWebElement*>(i));
}

void* QWebElementCollection___toList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QWebElement>;
}

class MyQWebFrame: public QWebFrame
{
public:
	bool event(QEvent * e) { return callbackQWebFrame_Event(this, e) != 0; };
	QVariant evaluateJavaScript(const QString & scriptSource) { QByteArray tf42f6a = scriptSource.toUtf8(); QtWebKit_PackedString scriptSourcePacked = { const_cast<char*>(tf42f6a.prepend("WHITESPACE").constData()+10), tf42f6a.size()-10 };return *static_cast<QVariant*>(callbackQWebFrame_EvaluateJavaScript(this, scriptSourcePacked)); };
	void Signal_ContentsSizeChanged(const QSize & size) { callbackQWebFrame_ContentsSizeChanged(this, const_cast<QSize*>(&size)); };
	void Signal_IconChanged() { callbackQWebFrame_IconChanged(this); };
	void Signal_InitialLayoutCompleted() { callbackQWebFrame_InitialLayoutCompleted(this); };
	void Signal_JavaScriptWindowObjectCleared() { callbackQWebFrame_JavaScriptWindowObjectCleared(this); };
	void Signal_LoadFinished(bool ok) { callbackQWebFrame_LoadFinished(this, ok); };
	void Signal_LoadStarted() { callbackQWebFrame_LoadStarted(this); };
	void Signal_PageChanged() { callbackQWebFrame_PageChanged(this); };
	void Signal_TitleChanged(const QString & title) { QByteArray t3c6de1 = title.toUtf8(); QtWebKit_PackedString titlePacked = { const_cast<char*>(t3c6de1.prepend("WHITESPACE").constData()+10), t3c6de1.size()-10 };callbackQWebFrame_TitleChanged(this, titlePacked); };
	void Signal_UrlChanged(const QUrl & url) { callbackQWebFrame_UrlChanged(this, const_cast<QUrl*>(&url)); };
	void print(QPrinter * printer) const { callbackQWebFrame_Print(const_cast<void*>(static_cast<const void*>(this)), printer); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQWebFrame_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQWebFrame_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQWebFrame_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQWebFrame_CustomEvent(this, event); };
	void deleteLater() { callbackQWebFrame_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQWebFrame_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQWebFrame_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtWebKit_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQWebFrame_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQWebFrame_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQWebFrame_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQWebFrame*)

int QWebFrame_QWebFrame_QRegisterMetaType(){qRegisterMetaType<QWebFrame*>(); return qRegisterMetaType<MyQWebFrame*>();}

char QWebFrame_EventDefault(void* ptr, void* e)
{
		return static_cast<QWebFrame*>(ptr)->QWebFrame::event(static_cast<QEvent*>(e));
}

void QWebFrame_AddToJavaScriptWindowObject(void* ptr, struct QtWebKit_PackedString name, void* object, long long own)
{
	static_cast<QWebFrame*>(ptr)->addToJavaScriptWindowObject(QString::fromUtf8(name.data, name.len), static_cast<QObject*>(object), static_cast<QWebFrame::ValueOwnership>(own));
}

void* QWebFrame_EvaluateJavaScript(void* ptr, struct QtWebKit_PackedString scriptSource)
{
	QVariant returnArg;
	QMetaObject::invokeMethod(static_cast<QWebFrame*>(ptr), "evaluateJavaScript", Q_RETURN_ARG(QVariant, returnArg), Q_ARG(const QString, QString::fromUtf8(scriptSource.data, scriptSource.len)));
	return new QVariant(returnArg);
}

void* QWebFrame_EvaluateJavaScriptDefault(void* ptr, struct QtWebKit_PackedString scriptSource)
{
		return new QVariant(static_cast<QWebFrame*>(ptr)->QWebFrame::evaluateJavaScript(QString::fromUtf8(scriptSource.data, scriptSource.len)));
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

void QWebFrame_Load2(void* ptr, void* req, long long operation, void* body)
{
	static_cast<QWebFrame*>(ptr)->load(*static_cast<QNetworkRequest*>(req), static_cast<QNetworkAccessManager::Operation>(operation), *static_cast<QByteArray*>(body));
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

void QWebFrame_LoadFinished(void* ptr, char ok)
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

void QWebFrame_Render2(void* ptr, void* painter, long long layer, void* clip)
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

void QWebFrame_ScrollToAnchor(void* ptr, struct QtWebKit_PackedString anchor)
{
	static_cast<QWebFrame*>(ptr)->scrollToAnchor(QString::fromUtf8(anchor.data, anchor.len));
}

void QWebFrame_SetContent(void* ptr, void* data, struct QtWebKit_PackedString mimeType, void* baseUrl)
{
	static_cast<QWebFrame*>(ptr)->setContent(*static_cast<QByteArray*>(data), QString::fromUtf8(mimeType.data, mimeType.len), *static_cast<QUrl*>(baseUrl));
}

void QWebFrame_SetFocus(void* ptr)
{
	static_cast<QWebFrame*>(ptr)->setFocus();
}

void QWebFrame_SetHtml(void* ptr, struct QtWebKit_PackedString html, void* baseUrl)
{
	static_cast<QWebFrame*>(ptr)->setHtml(QString::fromUtf8(html.data, html.len), *static_cast<QUrl*>(baseUrl));
}

void QWebFrame_SetScrollBarPolicy(void* ptr, long long orientation, long long policy)
{
	static_cast<QWebFrame*>(ptr)->setScrollBarPolicy(static_cast<Qt::Orientation>(orientation), static_cast<Qt::ScrollBarPolicy>(policy));
}

void QWebFrame_SetScrollBarValue(void* ptr, long long orientation, int value)
{
	static_cast<QWebFrame*>(ptr)->setScrollBarValue(static_cast<Qt::Orientation>(orientation), value);
}

void QWebFrame_SetScrollPosition(void* ptr, void* pos)
{
	static_cast<QWebFrame*>(ptr)->setScrollPosition(*static_cast<QPoint*>(pos));
}

void QWebFrame_SetTextSizeMultiplier(void* ptr, double factor)
{
	static_cast<QWebFrame*>(ptr)->setTextSizeMultiplier(factor);
}

void QWebFrame_SetUrl(void* ptr, void* url)
{
	static_cast<QWebFrame*>(ptr)->setUrl(*static_cast<QUrl*>(url));
}

void QWebFrame_SetZoomFactor(void* ptr, double factor)
{
	static_cast<QWebFrame*>(ptr)->setZoomFactor(factor);
}

void QWebFrame_ConnectTitleChanged(void* ptr)
{
	QObject::connect(static_cast<QWebFrame*>(ptr), static_cast<void (QWebFrame::*)(const QString &)>(&QWebFrame::titleChanged), static_cast<MyQWebFrame*>(ptr), static_cast<void (MyQWebFrame::*)(const QString &)>(&MyQWebFrame::Signal_TitleChanged));
}

void QWebFrame_DisconnectTitleChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWebFrame*>(ptr), static_cast<void (QWebFrame::*)(const QString &)>(&QWebFrame::titleChanged), static_cast<MyQWebFrame*>(ptr), static_cast<void (MyQWebFrame::*)(const QString &)>(&MyQWebFrame::Signal_TitleChanged));
}

void QWebFrame_TitleChanged(void* ptr, struct QtWebKit_PackedString title)
{
	static_cast<QWebFrame*>(ptr)->titleChanged(QString::fromUtf8(title.data, title.len));
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

void* QWebFrame_Icon(void* ptr)
{
	return new QIcon(static_cast<QWebFrame*>(ptr)->icon());
}

struct QtWebKit_PackedList QWebFrame_ChildFrames(void* ptr)
{
	return ({ QList<QWebFrame *>* tmpValue = new QList<QWebFrame *>(static_cast<QWebFrame*>(ptr)->childFrames()); QtWebKit_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtWebKit_PackedList QWebFrame_MetaData(void* ptr)
{
	return ({ QMultiMap<QString, QString>* tmpValue = new QMultiMap<QString, QString>(static_cast<QWebFrame*>(ptr)->metaData()); QtWebKit_PackedList { tmpValue, tmpValue->size() }; });
}

void* QWebFrame_Pos(void* ptr)
{
	return ({ QPoint tmpValue = static_cast<QWebFrame*>(ptr)->pos(); new QPoint(tmpValue.x(), tmpValue.y()); });
}

void* QWebFrame_ScrollPosition(void* ptr)
{
	return ({ QPoint tmpValue = static_cast<QWebFrame*>(ptr)->scrollPosition(); new QPoint(tmpValue.x(), tmpValue.y()); });
}

void* QWebFrame_Geometry(void* ptr)
{
	return ({ QRect tmpValue = static_cast<QWebFrame*>(ptr)->geometry(); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void* QWebFrame_ScrollBarGeometry(void* ptr, long long orientation)
{
	return ({ QRect tmpValue = static_cast<QWebFrame*>(ptr)->scrollBarGeometry(static_cast<Qt::Orientation>(orientation)); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void* QWebFrame_ContentsSize(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QWebFrame*>(ptr)->contentsSize(); new QSize(tmpValue.width(), tmpValue.height()); });
}

struct QtWebKit_PackedString QWebFrame_FrameName(void* ptr)
{
	return ({ QByteArray t8af5ca = static_cast<QWebFrame*>(ptr)->frameName().toUtf8(); QtWebKit_PackedString { const_cast<char*>(t8af5ca.prepend("WHITESPACE").constData()+10), t8af5ca.size()-10 }; });
}

struct QtWebKit_PackedString QWebFrame_Title(void* ptr)
{
	return ({ QByteArray tfd7064 = static_cast<QWebFrame*>(ptr)->title().toUtf8(); QtWebKit_PackedString { const_cast<char*>(tfd7064.prepend("WHITESPACE").constData()+10), tfd7064.size()-10 }; });
}

struct QtWebKit_PackedString QWebFrame_ToHtml(void* ptr)
{
	return ({ QByteArray t52635b = static_cast<QWebFrame*>(ptr)->toHtml().toUtf8(); QtWebKit_PackedString { const_cast<char*>(t52635b.prepend("WHITESPACE").constData()+10), t52635b.size()-10 }; });
}

struct QtWebKit_PackedString QWebFrame_ToPlainText(void* ptr)
{
	return ({ QByteArray tfa866a = static_cast<QWebFrame*>(ptr)->toPlainText().toUtf8(); QtWebKit_PackedString { const_cast<char*>(tfa866a.prepend("WHITESPACE").constData()+10), tfa866a.size()-10 }; });
}

void* QWebFrame_BaseUrl(void* ptr)
{
	return new QUrl(static_cast<QWebFrame*>(ptr)->baseUrl());
}

void* QWebFrame_RequestedUrl(void* ptr)
{
	return new QUrl(static_cast<QWebFrame*>(ptr)->requestedUrl());
}

void* QWebFrame_Url(void* ptr)
{
	return new QUrl(static_cast<QWebFrame*>(ptr)->url());
}

void* QWebFrame_DocumentElement(void* ptr)
{
	return new QWebElement(static_cast<QWebFrame*>(ptr)->documentElement());
}

void* QWebFrame_FindFirstElement(void* ptr, struct QtWebKit_PackedString selectorQuery)
{
	return new QWebElement(static_cast<QWebFrame*>(ptr)->findFirstElement(QString::fromUtf8(selectorQuery.data, selectorQuery.len)));
}

void* QWebFrame_OwnerElement(void* ptr)
{
	return new QWebElement(static_cast<QWebFrame*>(ptr)->ownerElement());
}

void* QWebFrame_FindAllElements(void* ptr, struct QtWebKit_PackedString selectorQuery)
{
	return new QWebElementCollection(static_cast<QWebFrame*>(ptr)->findAllElements(QString::fromUtf8(selectorQuery.data, selectorQuery.len)));
}

void* QWebFrame_ParentFrame(void* ptr)
{
	return static_cast<QWebFrame*>(ptr)->parentFrame();
}

void* QWebFrame_HitTestContent(void* ptr, void* pos)
{
	return new QWebHitTestResult(static_cast<QWebFrame*>(ptr)->hitTestContent(*static_cast<QPoint*>(pos)));
}

void* QWebFrame_Page(void* ptr)
{
	return static_cast<QWebFrame*>(ptr)->page();
}

void* QWebFrame_SecurityOrigin(void* ptr)
{
	return new QWebSecurityOrigin(static_cast<QWebFrame*>(ptr)->securityOrigin());
}

long long QWebFrame_ScrollBarPolicy(void* ptr, long long orientation)
{
	return static_cast<QWebFrame*>(ptr)->scrollBarPolicy(static_cast<Qt::Orientation>(orientation));
}

char QWebFrame_HasFocus(void* ptr)
{
	return static_cast<QWebFrame*>(ptr)->hasFocus();
}

int QWebFrame_ScrollBarMaximum(void* ptr, long long orientation)
{
	return static_cast<QWebFrame*>(ptr)->scrollBarMaximum(static_cast<Qt::Orientation>(orientation));
}

int QWebFrame_ScrollBarMinimum(void* ptr, long long orientation)
{
	return static_cast<QWebFrame*>(ptr)->scrollBarMinimum(static_cast<Qt::Orientation>(orientation));
}

int QWebFrame_ScrollBarValue(void* ptr, long long orientation)
{
	return static_cast<QWebFrame*>(ptr)->scrollBarValue(static_cast<Qt::Orientation>(orientation));
}

double QWebFrame_TextSizeMultiplier(void* ptr)
{
	return static_cast<QWebFrame*>(ptr)->textSizeMultiplier();
}

double QWebFrame_ZoomFactor(void* ptr)
{
	return static_cast<QWebFrame*>(ptr)->zoomFactor();
}

void QWebFrame_Print(void* ptr, void* printer)
{
	QMetaObject::invokeMethod(static_cast<QWebFrame*>(ptr), "print", Q_ARG(QPrinter*, static_cast<QPrinter*>(printer)));
}

void QWebFrame_PrintDefault(void* ptr, void* printer)
{
		static_cast<QWebFrame*>(ptr)->QWebFrame::print(static_cast<QPrinter*>(printer));
}

void* QWebFrame___childFrames_atList(void* ptr, int i)
{
	return const_cast<QWebFrame*>(static_cast<QList<QWebFrame *>*>(ptr)->at(i));
}

void QWebFrame___childFrames_setList(void* ptr, void* i)
{
	static_cast<QList<QWebFrame *>*>(ptr)->append(static_cast<QWebFrame*>(i));
}

void* QWebFrame___childFrames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QWebFrame *>;
}

struct QtWebKit_PackedString QWebFrame___metaData_atList(void* ptr, struct QtWebKit_PackedString i)
{
	return ({ QByteArray t25b5fd = static_cast<QMultiMap<QString, QString>*>(ptr)->value(QString::fromUtf8(i.data, i.len)).toUtf8(); QtWebKit_PackedString { const_cast<char*>(t25b5fd.prepend("WHITESPACE").constData()+10), t25b5fd.size()-10 }; });
}

void QWebFrame___metaData_setList(void* ptr, struct QtWebKit_PackedString key, struct QtWebKit_PackedString i)
{
	static_cast<QMultiMap<QString, QString>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), QString::fromUtf8(i.data, i.len));
}

void* QWebFrame___metaData_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMultiMap<QString, QString>;
}

struct QtWebKit_PackedList QWebFrame___metaData_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue = new QList<QString>(static_cast<QMultiMap<QString, QString>*>(ptr)->keys()); QtWebKit_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtWebKit_PackedString QWebFrame_____metaData_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray t29def6 = static_cast<QList<QString>*>(ptr)->at(i).toUtf8(); QtWebKit_PackedString { const_cast<char*>(t29def6.prepend("WHITESPACE").constData()+10), t29def6.size()-10 }; });
}

void QWebFrame_____metaData_keyList_setList(void* ptr, struct QtWebKit_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QWebFrame_____metaData_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>;
}

void* QWebFrame___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QWebFrame___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QWebFrame___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QWebFrame___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QWebFrame___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebFrame___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QWebFrame___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QWebFrame___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebFrame___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QWebFrame___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QWebFrame___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebFrame___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QWebFrame___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QWebFrame___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebFrame___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

char QWebFrame_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QWebFrame*>(ptr)->QWebFrame::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QWebFrame_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QWebFrame*>(ptr)->QWebFrame::childEvent(static_cast<QChildEvent*>(event));
}

void QWebFrame_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QWebFrame*>(ptr)->QWebFrame::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWebFrame_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QWebFrame*>(ptr)->QWebFrame::customEvent(static_cast<QEvent*>(event));
}

void QWebFrame_DeleteLaterDefault(void* ptr)
{
		static_cast<QWebFrame*>(ptr)->QWebFrame::deleteLater();
}

void QWebFrame_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QWebFrame*>(ptr)->QWebFrame::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWebFrame_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QWebFrame*>(ptr)->QWebFrame::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QWebFrame_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QWebFrame*>(ptr)->QWebFrame::metaObject());
}

void QWebHistory_Back(void* ptr)
{
	static_cast<QWebHistory*>(ptr)->back();
}

void QWebHistory_Clear(void* ptr)
{
	static_cast<QWebHistory*>(ptr)->clear();
}

void QWebHistory_Forward(void* ptr)
{
	static_cast<QWebHistory*>(ptr)->forward();
}

void QWebHistory_GoToItem(void* ptr, void* item)
{
	static_cast<QWebHistory*>(ptr)->goToItem(*static_cast<QWebHistoryItem*>(item));
}

void QWebHistory_SetMaximumItemCount(void* ptr, int count)
{
	static_cast<QWebHistory*>(ptr)->setMaximumItemCount(count);
}

struct QtWebKit_PackedList QWebHistory_BackItems(void* ptr, int maxItems)
{
	return ({ QList<QWebHistoryItem>* tmpValue = new QList<QWebHistoryItem>(static_cast<QWebHistory*>(ptr)->backItems(maxItems)); QtWebKit_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtWebKit_PackedList QWebHistory_ForwardItems(void* ptr, int maxItems)
{
	return ({ QList<QWebHistoryItem>* tmpValue = new QList<QWebHistoryItem>(static_cast<QWebHistory*>(ptr)->forwardItems(maxItems)); QtWebKit_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtWebKit_PackedList QWebHistory_Items(void* ptr)
{
	return ({ QList<QWebHistoryItem>* tmpValue = new QList<QWebHistoryItem>(static_cast<QWebHistory*>(ptr)->items()); QtWebKit_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtWebKit_PackedList QWebHistory_ToMap(void* ptr)
{
	return ({ QMap<QString, QVariant>* tmpValue = new QMap<QString, QVariant>(static_cast<QWebHistory*>(ptr)->toMap()); QtWebKit_PackedList { tmpValue, tmpValue->size() }; });
}

void* QWebHistory_BackItem(void* ptr)
{
	return new QWebHistoryItem(static_cast<QWebHistory*>(ptr)->backItem());
}

void* QWebHistory_CurrentItem(void* ptr)
{
	return new QWebHistoryItem(static_cast<QWebHistory*>(ptr)->currentItem());
}

void* QWebHistory_ForwardItem(void* ptr)
{
	return new QWebHistoryItem(static_cast<QWebHistory*>(ptr)->forwardItem());
}

void* QWebHistory_ItemAt(void* ptr, int i)
{
	return new QWebHistoryItem(static_cast<QWebHistory*>(ptr)->itemAt(i));
}

char QWebHistory_CanGoBack(void* ptr)
{
	return static_cast<QWebHistory*>(ptr)->canGoBack();
}

char QWebHistory_CanGoForward(void* ptr)
{
	return static_cast<QWebHistory*>(ptr)->canGoForward();
}

int QWebHistory_Count(void* ptr)
{
	return static_cast<QWebHistory*>(ptr)->count();
}

int QWebHistory_CurrentItemIndex(void* ptr)
{
	return static_cast<QWebHistory*>(ptr)->currentItemIndex();
}

int QWebHistory_MaximumItemCount(void* ptr)
{
	return static_cast<QWebHistory*>(ptr)->maximumItemCount();
}

void* QWebHistory___loadFromMap_map_atList(void* ptr, struct QtWebKit_PackedString i)
{
	return new QVariant(static_cast<QMap<QString, QVariant>*>(ptr)->value(QString::fromUtf8(i.data, i.len)));
}

void QWebHistory___loadFromMap_map_setList(void* ptr, struct QtWebKit_PackedString key, void* i)
{
	static_cast<QMap<QString, QVariant>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(i));
}

void* QWebHistory___loadFromMap_map_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<QString, QVariant>;
}

struct QtWebKit_PackedList QWebHistory___loadFromMap_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue = new QList<QString>(static_cast<QMap<QString, QVariant>*>(ptr)->keys()); QtWebKit_PackedList { tmpValue, tmpValue->size() }; });
}

void* QWebHistory___backItems_atList(void* ptr, int i)
{
	return new QWebHistoryItem(static_cast<QList<QWebHistoryItem>*>(ptr)->at(i));
}

void QWebHistory___backItems_setList(void* ptr, void* i)
{
	static_cast<QList<QWebHistoryItem>*>(ptr)->append(*static_cast<QWebHistoryItem*>(i));
}

void* QWebHistory___backItems_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QWebHistoryItem>;
}

void* QWebHistory___forwardItems_atList(void* ptr, int i)
{
	return new QWebHistoryItem(static_cast<QList<QWebHistoryItem>*>(ptr)->at(i));
}

void QWebHistory___forwardItems_setList(void* ptr, void* i)
{
	static_cast<QList<QWebHistoryItem>*>(ptr)->append(*static_cast<QWebHistoryItem*>(i));
}

void* QWebHistory___forwardItems_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QWebHistoryItem>;
}

void* QWebHistory___items_atList(void* ptr, int i)
{
	return new QWebHistoryItem(static_cast<QList<QWebHistoryItem>*>(ptr)->at(i));
}

void QWebHistory___items_setList(void* ptr, void* i)
{
	static_cast<QList<QWebHistoryItem>*>(ptr)->append(*static_cast<QWebHistoryItem*>(i));
}

void* QWebHistory___items_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QWebHistoryItem>;
}

void* QWebHistory___toMap_atList(void* ptr, struct QtWebKit_PackedString i)
{
	return new QVariant(static_cast<QMap<QString, QVariant>*>(ptr)->value(QString::fromUtf8(i.data, i.len)));
}

void QWebHistory___toMap_setList(void* ptr, struct QtWebKit_PackedString key, void* i)
{
	static_cast<QMap<QString, QVariant>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(i));
}

void* QWebHistory___toMap_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<QString, QVariant>;
}

struct QtWebKit_PackedList QWebHistory___toMap_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue = new QList<QString>(static_cast<QMap<QString, QVariant>*>(ptr)->keys()); QtWebKit_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtWebKit_PackedString QWebHistory_____loadFromMap_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray t29def6 = static_cast<QList<QString>*>(ptr)->at(i).toUtf8(); QtWebKit_PackedString { const_cast<char*>(t29def6.prepend("WHITESPACE").constData()+10), t29def6.size()-10 }; });
}

void QWebHistory_____loadFromMap_keyList_setList(void* ptr, struct QtWebKit_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QWebHistory_____loadFromMap_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>;
}

struct QtWebKit_PackedString QWebHistory_____toMap_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray t29def6 = static_cast<QList<QString>*>(ptr)->at(i).toUtf8(); QtWebKit_PackedString { const_cast<char*>(t29def6.prepend("WHITESPACE").constData()+10), t29def6.size()-10 }; });
}

void QWebHistory_____toMap_keyList_setList(void* ptr, struct QtWebKit_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QWebHistory_____toMap_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>;
}

class MyQWebHistoryInterface: public QWebHistoryInterface
{
public:
	MyQWebHistoryInterface(QObject *parent = Q_NULLPTR) : QWebHistoryInterface(parent) {QWebHistoryInterface_QWebHistoryInterface_QRegisterMetaType();};
	void addHistoryEntry(const QString & url) { QByteArray t817363 = url.toUtf8(); QtWebKit_PackedString urlPacked = { const_cast<char*>(t817363.prepend("WHITESPACE").constData()+10), t817363.size()-10 };callbackQWebHistoryInterface_AddHistoryEntry(this, urlPacked); };
	bool historyContains(const QString & url) const { QByteArray t817363 = url.toUtf8(); QtWebKit_PackedString urlPacked = { const_cast<char*>(t817363.prepend("WHITESPACE").constData()+10), t817363.size()-10 };return callbackQWebHistoryInterface_HistoryContains(const_cast<void*>(static_cast<const void*>(this)), urlPacked) != 0; };
	bool event(QEvent * e) { return callbackQWebHistoryInterface_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQWebHistoryInterface_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQWebHistoryInterface_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQWebHistoryInterface_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQWebHistoryInterface_CustomEvent(this, event); };
	void deleteLater() { callbackQWebHistoryInterface_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQWebHistoryInterface_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQWebHistoryInterface_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtWebKit_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQWebHistoryInterface_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQWebHistoryInterface_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQWebHistoryInterface_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQWebHistoryInterface*)

int QWebHistoryInterface_QWebHistoryInterface_QRegisterMetaType(){qRegisterMetaType<QWebHistoryInterface*>(); return qRegisterMetaType<MyQWebHistoryInterface*>();}

void* QWebHistoryInterface_QWebHistoryInterface_DefaultInterface()
{
	return QWebHistoryInterface::defaultInterface();
}

void* QWebHistoryInterface_NewQWebHistoryInterface(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQWebHistoryInterface(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQWebHistoryInterface(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQWebHistoryInterface(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQWebHistoryInterface(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQWebHistoryInterface(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQWebHistoryInterface(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQWebHistoryInterface(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQWebHistoryInterface(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQWebHistoryInterface(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQWebHistoryInterface(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQWebHistoryInterface(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQWebHistoryInterface(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQWebHistoryInterface(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQWebHistoryInterface(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQWebHistoryInterface(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQWebHistoryInterface(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQWebHistoryInterface(static_cast<QWindow*>(parent));
	} else {
		return new MyQWebHistoryInterface(static_cast<QObject*>(parent));
	}
}

void QWebHistoryInterface_AddHistoryEntry(void* ptr, struct QtWebKit_PackedString url)
{
	static_cast<QWebHistoryInterface*>(ptr)->addHistoryEntry(QString::fromUtf8(url.data, url.len));
}

void QWebHistoryInterface_QWebHistoryInterface_SetDefaultInterface(void* defaultInterface)
{
	QWebHistoryInterface::setDefaultInterface(static_cast<QWebHistoryInterface*>(defaultInterface));
}

void QWebHistoryInterface_DestroyQWebHistoryInterface(void* ptr)
{
	static_cast<QWebHistoryInterface*>(ptr)->~QWebHistoryInterface();
}

char QWebHistoryInterface_HistoryContains(void* ptr, struct QtWebKit_PackedString url)
{
	return static_cast<QWebHistoryInterface*>(ptr)->historyContains(QString::fromUtf8(url.data, url.len));
}

void* QWebHistoryInterface___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QWebHistoryInterface___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QWebHistoryInterface___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QWebHistoryInterface___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QWebHistoryInterface___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebHistoryInterface___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QWebHistoryInterface___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QWebHistoryInterface___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebHistoryInterface___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QWebHistoryInterface___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QWebHistoryInterface___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebHistoryInterface___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QWebHistoryInterface___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QWebHistoryInterface___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebHistoryInterface___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

char QWebHistoryInterface_EventDefault(void* ptr, void* e)
{
		return static_cast<QWebHistoryInterface*>(ptr)->QWebHistoryInterface::event(static_cast<QEvent*>(e));
}

char QWebHistoryInterface_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QWebHistoryInterface*>(ptr)->QWebHistoryInterface::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QWebHistoryInterface_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QWebHistoryInterface*>(ptr)->QWebHistoryInterface::childEvent(static_cast<QChildEvent*>(event));
}

void QWebHistoryInterface_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QWebHistoryInterface*>(ptr)->QWebHistoryInterface::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWebHistoryInterface_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QWebHistoryInterface*>(ptr)->QWebHistoryInterface::customEvent(static_cast<QEvent*>(event));
}

void QWebHistoryInterface_DeleteLaterDefault(void* ptr)
{
		static_cast<QWebHistoryInterface*>(ptr)->QWebHistoryInterface::deleteLater();
}

void QWebHistoryInterface_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QWebHistoryInterface*>(ptr)->QWebHistoryInterface::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWebHistoryInterface_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QWebHistoryInterface*>(ptr)->QWebHistoryInterface::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QWebHistoryInterface_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QWebHistoryInterface*>(ptr)->QWebHistoryInterface::metaObject());
}

void* QWebHistoryItem_NewQWebHistoryItem(void* other)
{
	return new QWebHistoryItem(*static_cast<QWebHistoryItem*>(other));
}

void QWebHistoryItem_SetUserData(void* ptr, void* userData)
{
	static_cast<QWebHistoryItem*>(ptr)->setUserData(*static_cast<QVariant*>(userData));
}

void QWebHistoryItem_DestroyQWebHistoryItem(void* ptr)
{
	static_cast<QWebHistoryItem*>(ptr)->~QWebHistoryItem();
}

void* QWebHistoryItem_LastVisited(void* ptr)
{
	return new QDateTime(static_cast<QWebHistoryItem*>(ptr)->lastVisited());
}

void* QWebHistoryItem_Icon(void* ptr)
{
	return new QIcon(static_cast<QWebHistoryItem*>(ptr)->icon());
}

struct QtWebKit_PackedString QWebHistoryItem_Title(void* ptr)
{
	return ({ QByteArray t984c4e = static_cast<QWebHistoryItem*>(ptr)->title().toUtf8(); QtWebKit_PackedString { const_cast<char*>(t984c4e.prepend("WHITESPACE").constData()+10), t984c4e.size()-10 }; });
}

void* QWebHistoryItem_OriginalUrl(void* ptr)
{
	return new QUrl(static_cast<QWebHistoryItem*>(ptr)->originalUrl());
}

void* QWebHistoryItem_Url(void* ptr)
{
	return new QUrl(static_cast<QWebHistoryItem*>(ptr)->url());
}

void* QWebHistoryItem_UserData(void* ptr)
{
	return new QVariant(static_cast<QWebHistoryItem*>(ptr)->userData());
}

struct QtWebKit_PackedList QWebHistoryItem_ToMap(void* ptr)
{
	return ({ QMap<QString, QVariant>* tmpValue = new QMap<QString, QVariant>(static_cast<QWebHistoryItem*>(ptr)->toMap()); QtWebKit_PackedList { tmpValue, tmpValue->size() }; });
}

char QWebHistoryItem_IsValid(void* ptr)
{
	return static_cast<QWebHistoryItem*>(ptr)->isValid();
}

void* QWebHistoryItem___loadFromMap_map_atList(void* ptr, struct QtWebKit_PackedString i)
{
	return new QVariant(static_cast<QMap<QString, QVariant>*>(ptr)->value(QString::fromUtf8(i.data, i.len)));
}

void QWebHistoryItem___loadFromMap_map_setList(void* ptr, struct QtWebKit_PackedString key, void* i)
{
	static_cast<QMap<QString, QVariant>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(i));
}

void* QWebHistoryItem___loadFromMap_map_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<QString, QVariant>;
}

struct QtWebKit_PackedList QWebHistoryItem___loadFromMap_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue = new QList<QString>(static_cast<QMap<QString, QVariant>*>(ptr)->keys()); QtWebKit_PackedList { tmpValue, tmpValue->size() }; });
}

void* QWebHistoryItem___toMap_atList(void* ptr, struct QtWebKit_PackedString i)
{
	return new QVariant(static_cast<QMap<QString, QVariant>*>(ptr)->value(QString::fromUtf8(i.data, i.len)));
}

void QWebHistoryItem___toMap_setList(void* ptr, struct QtWebKit_PackedString key, void* i)
{
	static_cast<QMap<QString, QVariant>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(i));
}

void* QWebHistoryItem___toMap_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<QString, QVariant>;
}

struct QtWebKit_PackedList QWebHistoryItem___toMap_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue = new QList<QString>(static_cast<QMap<QString, QVariant>*>(ptr)->keys()); QtWebKit_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtWebKit_PackedString QWebHistoryItem_____loadFromMap_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray t29def6 = static_cast<QList<QString>*>(ptr)->at(i).toUtf8(); QtWebKit_PackedString { const_cast<char*>(t29def6.prepend("WHITESPACE").constData()+10), t29def6.size()-10 }; });
}

void QWebHistoryItem_____loadFromMap_keyList_setList(void* ptr, struct QtWebKit_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QWebHistoryItem_____loadFromMap_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>;
}

struct QtWebKit_PackedString QWebHistoryItem_____toMap_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray t29def6 = static_cast<QList<QString>*>(ptr)->at(i).toUtf8(); QtWebKit_PackedString { const_cast<char*>(t29def6.prepend("WHITESPACE").constData()+10), t29def6.size()-10 }; });
}

void QWebHistoryItem_____toMap_keyList_setList(void* ptr, struct QtWebKit_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QWebHistoryItem_____toMap_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>;
}

void* QWebHitTestResult_NewQWebHitTestResult()
{
	return new QWebHitTestResult();
}

void* QWebHitTestResult_NewQWebHitTestResult2(void* other)
{
	return new QWebHitTestResult(*static_cast<QWebHitTestResult*>(other));
}

void QWebHitTestResult_DestroyQWebHitTestResult(void* ptr)
{
	static_cast<QWebHitTestResult*>(ptr)->~QWebHitTestResult();
}

void* QWebHitTestResult_Pixmap(void* ptr)
{
	return new QPixmap(static_cast<QWebHitTestResult*>(ptr)->pixmap());
}

void* QWebHitTestResult_Pos(void* ptr)
{
	return ({ QPoint tmpValue = static_cast<QWebHitTestResult*>(ptr)->pos(); new QPoint(tmpValue.x(), tmpValue.y()); });
}

void* QWebHitTestResult_BoundingRect(void* ptr)
{
	return ({ QRect tmpValue = static_cast<QWebHitTestResult*>(ptr)->boundingRect(); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

struct QtWebKit_PackedString QWebHitTestResult_AlternateText(void* ptr)
{
	return ({ QByteArray t891f2b = static_cast<QWebHitTestResult*>(ptr)->alternateText().toUtf8(); QtWebKit_PackedString { const_cast<char*>(t891f2b.prepend("WHITESPACE").constData()+10), t891f2b.size()-10 }; });
}

struct QtWebKit_PackedString QWebHitTestResult_LinkText(void* ptr)
{
	return ({ QByteArray t01b6fa = static_cast<QWebHitTestResult*>(ptr)->linkText().toUtf8(); QtWebKit_PackedString { const_cast<char*>(t01b6fa.prepend("WHITESPACE").constData()+10), t01b6fa.size()-10 }; });
}

struct QtWebKit_PackedString QWebHitTestResult_LinkTitleString(void* ptr)
{
	return ({ QByteArray tf0511d = static_cast<QWebHitTestResult*>(ptr)->linkTitleString().toUtf8(); QtWebKit_PackedString { const_cast<char*>(tf0511d.prepend("WHITESPACE").constData()+10), tf0511d.size()-10 }; });
}

struct QtWebKit_PackedString QWebHitTestResult_Title(void* ptr)
{
	return ({ QByteArray t4a2a4b = static_cast<QWebHitTestResult*>(ptr)->title().toUtf8(); QtWebKit_PackedString { const_cast<char*>(t4a2a4b.prepend("WHITESPACE").constData()+10), t4a2a4b.size()-10 }; });
}

void* QWebHitTestResult_ImageUrl(void* ptr)
{
	return new QUrl(static_cast<QWebHitTestResult*>(ptr)->imageUrl());
}

void* QWebHitTestResult_LinkUrl(void* ptr)
{
	return new QUrl(static_cast<QWebHitTestResult*>(ptr)->linkUrl());
}

void* QWebHitTestResult_MediaUrl(void* ptr)
{
	return new QUrl(static_cast<QWebHitTestResult*>(ptr)->mediaUrl());
}

void* QWebHitTestResult_Element(void* ptr)
{
	return new QWebElement(static_cast<QWebHitTestResult*>(ptr)->element());
}

void* QWebHitTestResult_EnclosingBlockElement(void* ptr)
{
	return new QWebElement(static_cast<QWebHitTestResult*>(ptr)->enclosingBlockElement());
}

void* QWebHitTestResult_LinkElement(void* ptr)
{
	return new QWebElement(static_cast<QWebHitTestResult*>(ptr)->linkElement());
}

void* QWebHitTestResult_Frame(void* ptr)
{
	return static_cast<QWebHitTestResult*>(ptr)->frame();
}

void* QWebHitTestResult_LinkTargetFrame(void* ptr)
{
	return static_cast<QWebHitTestResult*>(ptr)->linkTargetFrame();
}

char QWebHitTestResult_IsContentEditable(void* ptr)
{
	return static_cast<QWebHitTestResult*>(ptr)->isContentEditable();
}

char QWebHitTestResult_IsContentSelected(void* ptr)
{
	return static_cast<QWebHitTestResult*>(ptr)->isContentSelected();
}

char QWebHitTestResult_IsNull(void* ptr)
{
	return static_cast<QWebHitTestResult*>(ptr)->isNull();
}

class MyQWebInspector: public QWebInspector
{
public:
	MyQWebInspector(QWidget *parent = Q_NULLPTR) : QWebInspector(parent) {QWebInspector_QWebInspector_QRegisterMetaType();};
	bool event(QEvent * ev) { return callbackQWebInspector_Event(this, ev) != 0; };
	void closeEvent(QCloseEvent * event) { callbackQWebInspector_CloseEvent(this, event); };
	void hideEvent(QHideEvent * event) { callbackQWebInspector_HideEvent(this, event); };
	void resizeEvent(QResizeEvent * event) { callbackQWebInspector_ResizeEvent(this, event); };
	void showEvent(QShowEvent * event) { callbackQWebInspector_ShowEvent(this, event); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQWebInspector_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	bool close() { return callbackQWebInspector_Close(this) != 0; };
	bool focusNextPrevChild(bool next) { return callbackQWebInspector_FocusNextPrevChild(this, next) != 0; };
	bool nativeEvent(const QByteArray & eventType, void * message, long * result) { return callbackQWebInspector_NativeEvent(this, const_cast<QByteArray*>(&eventType), message, *result) != 0; };
	void actionEvent(QActionEvent * event) { callbackQWebInspector_ActionEvent(this, event); };
	void changeEvent(QEvent * event) { callbackQWebInspector_ChangeEvent(this, event); };
	void contextMenuEvent(QContextMenuEvent * event) { callbackQWebInspector_ContextMenuEvent(this, event); };
	void Signal_CustomContextMenuRequested(const QPoint & pos) { callbackQWebInspector_CustomContextMenuRequested(this, const_cast<QPoint*>(&pos)); };
	void dragEnterEvent(QDragEnterEvent * event) { callbackQWebInspector_DragEnterEvent(this, event); };
	void dragLeaveEvent(QDragLeaveEvent * event) { callbackQWebInspector_DragLeaveEvent(this, event); };
	void dragMoveEvent(QDragMoveEvent * event) { callbackQWebInspector_DragMoveEvent(this, event); };
	void dropEvent(QDropEvent * event) { callbackQWebInspector_DropEvent(this, event); };
	void enterEvent(QEvent * event) { callbackQWebInspector_EnterEvent(this, event); };
	void focusInEvent(QFocusEvent * event) { callbackQWebInspector_FocusInEvent(this, event); };
	void focusOutEvent(QFocusEvent * event) { callbackQWebInspector_FocusOutEvent(this, event); };
	void hide() { callbackQWebInspector_Hide(this); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackQWebInspector_InputMethodEvent(this, event); };
	void keyPressEvent(QKeyEvent * event) { callbackQWebInspector_KeyPressEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQWebInspector_KeyReleaseEvent(this, event); };
	void leaveEvent(QEvent * event) { callbackQWebInspector_LeaveEvent(this, event); };
	void lower() { callbackQWebInspector_Lower(this); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQWebInspector_MouseDoubleClickEvent(this, event); };
	void mouseMoveEvent(QMouseEvent * event) { callbackQWebInspector_MouseMoveEvent(this, event); };
	void mousePressEvent(QMouseEvent * event) { callbackQWebInspector_MousePressEvent(this, event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackQWebInspector_MouseReleaseEvent(this, event); };
	void moveEvent(QMoveEvent * event) { callbackQWebInspector_MoveEvent(this, event); };
	void paintEvent(QPaintEvent * event) { callbackQWebInspector_PaintEvent(this, event); };
	void raise() { callbackQWebInspector_Raise(this); };
	void repaint() { callbackQWebInspector_Repaint(this); };
	void setDisabled(bool disable) { callbackQWebInspector_SetDisabled(this, disable); };
	void setEnabled(bool vbo) { callbackQWebInspector_SetEnabled(this, vbo); };
	void setFocus() { callbackQWebInspector_SetFocus2(this); };
	void setHidden(bool hidden) { callbackQWebInspector_SetHidden(this, hidden); };
	void setStyleSheet(const QString & styleSheet) { QByteArray t728ae7 = styleSheet.toUtf8(); QtWebKit_PackedString styleSheetPacked = { const_cast<char*>(t728ae7.prepend("WHITESPACE").constData()+10), t728ae7.size()-10 };callbackQWebInspector_SetStyleSheet(this, styleSheetPacked); };
	void setVisible(bool visible) { callbackQWebInspector_SetVisible(this, visible); };
	void setWindowModified(bool vbo) { callbackQWebInspector_SetWindowModified(this, vbo); };
	void setWindowTitle(const QString & vqs) { QByteArray tda39a3 = vqs.toUtf8(); QtWebKit_PackedString vqsPacked = { const_cast<char*>(tda39a3.prepend("WHITESPACE").constData()+10), tda39a3.size()-10 };callbackQWebInspector_SetWindowTitle(this, vqsPacked); };
	void show() { callbackQWebInspector_Show(this); };
	void showFullScreen() { callbackQWebInspector_ShowFullScreen(this); };
	void showMaximized() { callbackQWebInspector_ShowMaximized(this); };
	void showMinimized() { callbackQWebInspector_ShowMinimized(this); };
	void showNormal() { callbackQWebInspector_ShowNormal(this); };
	void tabletEvent(QTabletEvent * event) { callbackQWebInspector_TabletEvent(this, event); };
	void update() { callbackQWebInspector_Update(this); };
	void updateMicroFocus() { callbackQWebInspector_UpdateMicroFocus(this); };
	void wheelEvent(QWheelEvent * event) { callbackQWebInspector_WheelEvent(this, event); };
	void Signal_WindowIconChanged(const QIcon & icon) { callbackQWebInspector_WindowIconChanged(this, const_cast<QIcon*>(&icon)); };
	void Signal_WindowTitleChanged(const QString & title) { QByteArray t3c6de1 = title.toUtf8(); QtWebKit_PackedString titlePacked = { const_cast<char*>(t3c6de1.prepend("WHITESPACE").constData()+10), t3c6de1.size()-10 };callbackQWebInspector_WindowTitleChanged(this, titlePacked); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQWebInspector_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize minimumSizeHint() const { return *static_cast<QSize*>(callbackQWebInspector_MinimumSizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackQWebInspector_InputMethodQuery(const_cast<void*>(static_cast<const void*>(this)), query)); };
	bool hasHeightForWidth() const { return callbackQWebInspector_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int heightForWidth(int w) const { return callbackQWebInspector_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackQWebInspector_Metric(const_cast<void*>(static_cast<const void*>(this)), m); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQWebInspector_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQWebInspector_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQWebInspector_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQWebInspector_CustomEvent(this, event); };
	void deleteLater() { callbackQWebInspector_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQWebInspector_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQWebInspector_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtWebKit_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQWebInspector_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQWebInspector_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQWebInspector_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQWebInspector*)

int QWebInspector_QWebInspector_QRegisterMetaType(){qRegisterMetaType<QWebInspector*>(); return qRegisterMetaType<MyQWebInspector*>();}

void* QWebInspector_NewQWebInspector(void* parent)
{
		return new MyQWebInspector(static_cast<QWidget*>(parent));
}

char QWebInspector_EventDefault(void* ptr, void* ev)
{
		return static_cast<QWebInspector*>(ptr)->QWebInspector::event(static_cast<QEvent*>(ev));
}

void QWebInspector_CloseEventDefault(void* ptr, void* event)
{
		static_cast<QWebInspector*>(ptr)->QWebInspector::closeEvent(static_cast<QCloseEvent*>(event));
}

void QWebInspector_HideEventDefault(void* ptr, void* event)
{
		static_cast<QWebInspector*>(ptr)->QWebInspector::hideEvent(static_cast<QHideEvent*>(event));
}

void QWebInspector_ResizeEventDefault(void* ptr, void* event)
{
		static_cast<QWebInspector*>(ptr)->QWebInspector::resizeEvent(static_cast<QResizeEvent*>(event));
}

void QWebInspector_SetPage(void* ptr, void* page)
{
	static_cast<QWebInspector*>(ptr)->setPage(static_cast<QWebPage*>(page));
}

void QWebInspector_ShowEventDefault(void* ptr, void* event)
{
		static_cast<QWebInspector*>(ptr)->QWebInspector::showEvent(static_cast<QShowEvent*>(event));
}

void QWebInspector_DestroyQWebInspector(void* ptr)
{
	static_cast<QWebInspector*>(ptr)->~QWebInspector();
}

void* QWebInspector_SizeHintDefault(void* ptr)
{
		return ({ QSize tmpValue = static_cast<QWebInspector*>(ptr)->QWebInspector::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QWebInspector_Page(void* ptr)
{
	return static_cast<QWebInspector*>(ptr)->page();
}

void* QWebInspector___addActions_actions_atList(void* ptr, int i)
{
	return const_cast<QAction*>(static_cast<QList<QAction *>*>(ptr)->at(i));
}

void QWebInspector___addActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QWebInspector___addActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>;
}

void* QWebInspector___insertActions_actions_atList(void* ptr, int i)
{
	return const_cast<QAction*>(static_cast<QList<QAction *>*>(ptr)->at(i));
}

void QWebInspector___insertActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QWebInspector___insertActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>;
}

void* QWebInspector___actions_atList(void* ptr, int i)
{
	return const_cast<QAction*>(static_cast<QList<QAction *>*>(ptr)->at(i));
}

void QWebInspector___actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QWebInspector___actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>;
}

void* QWebInspector___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QWebInspector___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QWebInspector___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QWebInspector___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QWebInspector___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebInspector___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QWebInspector___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QWebInspector___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebInspector___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QWebInspector___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QWebInspector___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebInspector___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QWebInspector___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QWebInspector___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebInspector___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

char QWebInspector_CloseDefault(void* ptr)
{
		return static_cast<QWebInspector*>(ptr)->QWebInspector::close();
}

char QWebInspector_FocusNextPrevChildDefault(void* ptr, char next)
{
		return static_cast<QWebInspector*>(ptr)->QWebInspector::focusNextPrevChild(next != 0);
}

char QWebInspector_NativeEventDefault(void* ptr, void* eventType, void* message, long result)
{
		return static_cast<QWebInspector*>(ptr)->QWebInspector::nativeEvent(*static_cast<QByteArray*>(eventType), message, &result);
}

void QWebInspector_ActionEventDefault(void* ptr, void* event)
{
		static_cast<QWebInspector*>(ptr)->QWebInspector::actionEvent(static_cast<QActionEvent*>(event));
}

void QWebInspector_ChangeEventDefault(void* ptr, void* event)
{
		static_cast<QWebInspector*>(ptr)->QWebInspector::changeEvent(static_cast<QEvent*>(event));
}

void QWebInspector_ContextMenuEventDefault(void* ptr, void* event)
{
		static_cast<QWebInspector*>(ptr)->QWebInspector::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
}

void QWebInspector_DragEnterEventDefault(void* ptr, void* event)
{
		static_cast<QWebInspector*>(ptr)->QWebInspector::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QWebInspector_DragLeaveEventDefault(void* ptr, void* event)
{
		static_cast<QWebInspector*>(ptr)->QWebInspector::dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QWebInspector_DragMoveEventDefault(void* ptr, void* event)
{
		static_cast<QWebInspector*>(ptr)->QWebInspector::dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QWebInspector_DropEventDefault(void* ptr, void* event)
{
		static_cast<QWebInspector*>(ptr)->QWebInspector::dropEvent(static_cast<QDropEvent*>(event));
}

void QWebInspector_EnterEventDefault(void* ptr, void* event)
{
		static_cast<QWebInspector*>(ptr)->QWebInspector::enterEvent(static_cast<QEvent*>(event));
}

void QWebInspector_FocusInEventDefault(void* ptr, void* event)
{
		static_cast<QWebInspector*>(ptr)->QWebInspector::focusInEvent(static_cast<QFocusEvent*>(event));
}

void QWebInspector_FocusOutEventDefault(void* ptr, void* event)
{
		static_cast<QWebInspector*>(ptr)->QWebInspector::focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QWebInspector_HideDefault(void* ptr)
{
		static_cast<QWebInspector*>(ptr)->QWebInspector::hide();
}

void QWebInspector_InputMethodEventDefault(void* ptr, void* event)
{
		static_cast<QWebInspector*>(ptr)->QWebInspector::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void QWebInspector_KeyPressEventDefault(void* ptr, void* event)
{
		static_cast<QWebInspector*>(ptr)->QWebInspector::keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QWebInspector_KeyReleaseEventDefault(void* ptr, void* event)
{
		static_cast<QWebInspector*>(ptr)->QWebInspector::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QWebInspector_LeaveEventDefault(void* ptr, void* event)
{
		static_cast<QWebInspector*>(ptr)->QWebInspector::leaveEvent(static_cast<QEvent*>(event));
}

void QWebInspector_LowerDefault(void* ptr)
{
		static_cast<QWebInspector*>(ptr)->QWebInspector::lower();
}

void QWebInspector_MouseDoubleClickEventDefault(void* ptr, void* event)
{
		static_cast<QWebInspector*>(ptr)->QWebInspector::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QWebInspector_MouseMoveEventDefault(void* ptr, void* event)
{
		static_cast<QWebInspector*>(ptr)->QWebInspector::mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QWebInspector_MousePressEventDefault(void* ptr, void* event)
{
		static_cast<QWebInspector*>(ptr)->QWebInspector::mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QWebInspector_MouseReleaseEventDefault(void* ptr, void* event)
{
		static_cast<QWebInspector*>(ptr)->QWebInspector::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void QWebInspector_MoveEventDefault(void* ptr, void* event)
{
		static_cast<QWebInspector*>(ptr)->QWebInspector::moveEvent(static_cast<QMoveEvent*>(event));
}

void QWebInspector_PaintEventDefault(void* ptr, void* event)
{
		static_cast<QWebInspector*>(ptr)->QWebInspector::paintEvent(static_cast<QPaintEvent*>(event));
}

void QWebInspector_RaiseDefault(void* ptr)
{
		static_cast<QWebInspector*>(ptr)->QWebInspector::raise();
}

void QWebInspector_RepaintDefault(void* ptr)
{
		static_cast<QWebInspector*>(ptr)->QWebInspector::repaint();
}

void QWebInspector_SetDisabledDefault(void* ptr, char disable)
{
		static_cast<QWebInspector*>(ptr)->QWebInspector::setDisabled(disable != 0);
}

void QWebInspector_SetEnabledDefault(void* ptr, char vbo)
{
		static_cast<QWebInspector*>(ptr)->QWebInspector::setEnabled(vbo != 0);
}

void QWebInspector_SetFocus2Default(void* ptr)
{
		static_cast<QWebInspector*>(ptr)->QWebInspector::setFocus();
}

void QWebInspector_SetHiddenDefault(void* ptr, char hidden)
{
		static_cast<QWebInspector*>(ptr)->QWebInspector::setHidden(hidden != 0);
}

void QWebInspector_SetStyleSheetDefault(void* ptr, struct QtWebKit_PackedString styleSheet)
{
		static_cast<QWebInspector*>(ptr)->QWebInspector::setStyleSheet(QString::fromUtf8(styleSheet.data, styleSheet.len));
}

void QWebInspector_SetVisibleDefault(void* ptr, char visible)
{
		static_cast<QWebInspector*>(ptr)->QWebInspector::setVisible(visible != 0);
}

void QWebInspector_SetWindowModifiedDefault(void* ptr, char vbo)
{
		static_cast<QWebInspector*>(ptr)->QWebInspector::setWindowModified(vbo != 0);
}

void QWebInspector_SetWindowTitleDefault(void* ptr, struct QtWebKit_PackedString vqs)
{
		static_cast<QWebInspector*>(ptr)->QWebInspector::setWindowTitle(QString::fromUtf8(vqs.data, vqs.len));
}

void QWebInspector_ShowDefault(void* ptr)
{
		static_cast<QWebInspector*>(ptr)->QWebInspector::show();
}

void QWebInspector_ShowFullScreenDefault(void* ptr)
{
		static_cast<QWebInspector*>(ptr)->QWebInspector::showFullScreen();
}

void QWebInspector_ShowMaximizedDefault(void* ptr)
{
		static_cast<QWebInspector*>(ptr)->QWebInspector::showMaximized();
}

void QWebInspector_ShowMinimizedDefault(void* ptr)
{
		static_cast<QWebInspector*>(ptr)->QWebInspector::showMinimized();
}

void QWebInspector_ShowNormalDefault(void* ptr)
{
		static_cast<QWebInspector*>(ptr)->QWebInspector::showNormal();
}

void QWebInspector_TabletEventDefault(void* ptr, void* event)
{
		static_cast<QWebInspector*>(ptr)->QWebInspector::tabletEvent(static_cast<QTabletEvent*>(event));
}

void QWebInspector_UpdateDefault(void* ptr)
{
		static_cast<QWebInspector*>(ptr)->QWebInspector::update();
}

void QWebInspector_UpdateMicroFocusDefault(void* ptr)
{
		static_cast<QWebInspector*>(ptr)->QWebInspector::updateMicroFocus();
}

void QWebInspector_WheelEventDefault(void* ptr, void* event)
{
		static_cast<QWebInspector*>(ptr)->QWebInspector::wheelEvent(static_cast<QWheelEvent*>(event));
}

void* QWebInspector_PaintEngineDefault(void* ptr)
{
		return static_cast<QWebInspector*>(ptr)->QWebInspector::paintEngine();
}

void* QWebInspector_MinimumSizeHintDefault(void* ptr)
{
		return ({ QSize tmpValue = static_cast<QWebInspector*>(ptr)->QWebInspector::minimumSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QWebInspector_InputMethodQueryDefault(void* ptr, long long query)
{
		return new QVariant(static_cast<QWebInspector*>(ptr)->QWebInspector::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

char QWebInspector_HasHeightForWidthDefault(void* ptr)
{
		return static_cast<QWebInspector*>(ptr)->QWebInspector::hasHeightForWidth();
}

int QWebInspector_HeightForWidthDefault(void* ptr, int w)
{
		return static_cast<QWebInspector*>(ptr)->QWebInspector::heightForWidth(w);
}

int QWebInspector_MetricDefault(void* ptr, long long m)
{
		return static_cast<QWebInspector*>(ptr)->QWebInspector::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
}

char QWebInspector_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QWebInspector*>(ptr)->QWebInspector::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QWebInspector_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QWebInspector*>(ptr)->QWebInspector::childEvent(static_cast<QChildEvent*>(event));
}

void QWebInspector_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QWebInspector*>(ptr)->QWebInspector::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWebInspector_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QWebInspector*>(ptr)->QWebInspector::customEvent(static_cast<QEvent*>(event));
}

void QWebInspector_DeleteLaterDefault(void* ptr)
{
		static_cast<QWebInspector*>(ptr)->QWebInspector::deleteLater();
}

void QWebInspector_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QWebInspector*>(ptr)->QWebInspector::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWebInspector_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QWebInspector*>(ptr)->QWebInspector::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QWebInspector_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QWebInspector*>(ptr)->QWebInspector::metaObject());
}

class MyQWebPage: public QWebPage
{
public:
	MyQWebPage(QObject *parent = Q_NULLPTR) : QWebPage(parent) {QWebPage_QWebPage_QRegisterMetaType();};
	QObject * createPlugin(const QString & classid, const QUrl & url, const QStringList & paramNames, const QStringList & paramValues) { QByteArray tecec06 = classid.toUtf8(); QtWebKit_PackedString classidPacked = { const_cast<char*>(tecec06.prepend("WHITESPACE").constData()+10), tecec06.size()-10 };QByteArray te78583 = paramNames.join("|").toUtf8(); QtWebKit_PackedString paramNamesPacked = { const_cast<char*>(te78583.prepend("WHITESPACE").constData()+10), te78583.size()-10 };QByteArray t90393c = paramValues.join("|").toUtf8(); QtWebKit_PackedString paramValuesPacked = { const_cast<char*>(t90393c.prepend("WHITESPACE").constData()+10), t90393c.size()-10 };return static_cast<QObject*>(callbackQWebPage_CreatePlugin(this, classidPacked, const_cast<QUrl*>(&url), paramNamesPacked, paramValuesPacked)); };
	QString chooseFile(QWebFrame * parentFrame, const QString & suggestedFile) { QByteArray tf87690 = suggestedFile.toUtf8(); QtWebKit_PackedString suggestedFilePacked = { const_cast<char*>(tf87690.prepend("WHITESPACE").constData()+10), tf87690.size()-10 };return ({ QtWebKit_PackedString tempVal = callbackQWebPage_ChooseFile(this, parentFrame, suggestedFilePacked); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	QWebPage * createWindow(QWebPage::WebWindowType ty) { return static_cast<QWebPage*>(callbackQWebPage_CreateWindow(this, ty)); };
	bool acceptNavigationRequest(QWebFrame * frame, const QNetworkRequest & request, QWebPage::NavigationType ty) { return callbackQWebPage_AcceptNavigationRequest(this, frame, const_cast<QNetworkRequest*>(&request), ty) != 0; };
	bool event(QEvent * ev) { return callbackQWebPage_Event(this, ev) != 0; };
	bool javaScriptConfirm(QWebFrame * frame, const QString & msg) { QByteArray t19f34e = msg.toUtf8(); QtWebKit_PackedString msgPacked = { const_cast<char*>(t19f34e.prepend("WHITESPACE").constData()+10), t19f34e.size()-10 };return callbackQWebPage_JavaScriptConfirm(this, frame, msgPacked) != 0; };
	bool javaScriptPrompt(QWebFrame * frame, const QString & msg, const QString & defaultValue, QString * result) { QByteArray t19f34e = msg.toUtf8(); QtWebKit_PackedString msgPacked = { const_cast<char*>(t19f34e.prepend("WHITESPACE").constData()+10), t19f34e.size()-10 };QByteArray te940d2 = defaultValue.toUtf8(); QtWebKit_PackedString defaultValuePacked = { const_cast<char*>(te940d2.prepend("WHITESPACE").constData()+10), te940d2.size()-10 };QByteArray t37a530 = result->toUtf8(); QtWebKit_PackedString resultPacked = { const_cast<char*>(t37a530.prepend("WHITESPACE").constData()+10), t37a530.size()-10 };return callbackQWebPage_JavaScriptPrompt(this, frame, msgPacked, defaultValuePacked, resultPacked) != 0; };
	bool shouldInterruptJavaScript() { return callbackQWebPage_ShouldInterruptJavaScript(this) != 0; };
	void Signal_ApplicationCacheQuotaExceeded(QWebSecurityOrigin * origin, quint64 defaultOriginQuota, quint64 totalSpaceNeeded) { callbackQWebPage_ApplicationCacheQuotaExceeded(this, origin, defaultOriginQuota, totalSpaceNeeded); };
	void Signal_ConsoleMessageReceived(QWebPage::MessageSource source, QWebPage::MessageLevel level, const QString & message, int lineNumber, const QString & sourceID) { QByteArray t6f9b9a = message.toUtf8(); QtWebKit_PackedString messagePacked = { const_cast<char*>(t6f9b9a.prepend("WHITESPACE").constData()+10), t6f9b9a.size()-10 };QByteArray tf767e3 = sourceID.toUtf8(); QtWebKit_PackedString sourceIDPacked = { const_cast<char*>(tf767e3.prepend("WHITESPACE").constData()+10), tf767e3.size()-10 };callbackQWebPage_ConsoleMessageReceived(this, source, level, messagePacked, lineNumber, sourceIDPacked); };
	void Signal_ContentsChanged() { callbackQWebPage_ContentsChanged(this); };
	void Signal_DatabaseQuotaExceeded(QWebFrame * frame, QString databaseName) { QByteArray t3266b3 = databaseName.toUtf8(); QtWebKit_PackedString databaseNamePacked = { const_cast<char*>(t3266b3.prepend("WHITESPACE").constData()+10), t3266b3.size()-10 };callbackQWebPage_DatabaseQuotaExceeded(this, frame, databaseNamePacked); };
	void Signal_DownloadRequested(const QNetworkRequest & request) { callbackQWebPage_DownloadRequested(this, const_cast<QNetworkRequest*>(&request)); };
	void Signal_FeaturePermissionRequestCanceled(QWebFrame * frame, QWebPage::Feature feature) { callbackQWebPage_FeaturePermissionRequestCanceled(this, frame, feature); };
	void Signal_FeaturePermissionRequested(QWebFrame * frame, QWebPage::Feature feature) { callbackQWebPage_FeaturePermissionRequested(this, frame, feature); };
	void Signal_FocusedElementChanged(const QWebElement & element) { callbackQWebPage_FocusedElementChanged(this, const_cast<QWebElement*>(&element)); };
	void Signal_FrameCreated(QWebFrame * frame) { callbackQWebPage_FrameCreated(this, frame); };
	void Signal_GeometryChangeRequested(const QRect & geom) { callbackQWebPage_GeometryChangeRequested(this, const_cast<QRect*>(&geom)); };
	void javaScriptAlert(QWebFrame * frame, const QString & msg) { QByteArray t19f34e = msg.toUtf8(); QtWebKit_PackedString msgPacked = { const_cast<char*>(t19f34e.prepend("WHITESPACE").constData()+10), t19f34e.size()-10 };callbackQWebPage_JavaScriptAlert(this, frame, msgPacked); };
	void javaScriptConsoleMessage(const QString & message, int lineNumber, const QString & sourceID) { QByteArray t6f9b9a = message.toUtf8(); QtWebKit_PackedString messagePacked = { const_cast<char*>(t6f9b9a.prepend("WHITESPACE").constData()+10), t6f9b9a.size()-10 };QByteArray tf767e3 = sourceID.toUtf8(); QtWebKit_PackedString sourceIDPacked = { const_cast<char*>(tf767e3.prepend("WHITESPACE").constData()+10), tf767e3.size()-10 };callbackQWebPage_JavaScriptConsoleMessage(this, messagePacked, lineNumber, sourceIDPacked); };
	void Signal_LinkClicked(const QUrl & url) { callbackQWebPage_LinkClicked(this, const_cast<QUrl*>(&url)); };
	void Signal_LinkHovered(const QString & link, const QString & title, const QString & textContent) { QByteArray t4f0aa5 = link.toUtf8(); QtWebKit_PackedString linkPacked = { const_cast<char*>(t4f0aa5.prepend("WHITESPACE").constData()+10), t4f0aa5.size()-10 };QByteArray t3c6de1 = title.toUtf8(); QtWebKit_PackedString titlePacked = { const_cast<char*>(t3c6de1.prepend("WHITESPACE").constData()+10), t3c6de1.size()-10 };QByteArray t2ec889 = textContent.toUtf8(); QtWebKit_PackedString textContentPacked = { const_cast<char*>(t2ec889.prepend("WHITESPACE").constData()+10), t2ec889.size()-10 };callbackQWebPage_LinkHovered(this, linkPacked, titlePacked, textContentPacked); };
	void Signal_LoadFinished(bool ok) { callbackQWebPage_LoadFinished(this, ok); };
	void Signal_LoadProgress(int progress) { callbackQWebPage_LoadProgress(this, progress); };
	void Signal_LoadStarted() { callbackQWebPage_LoadStarted(this); };
	void Signal_MenuBarVisibilityChangeRequested(bool visible) { callbackQWebPage_MenuBarVisibilityChangeRequested(this, visible); };
	void Signal_MicroFocusChanged() { callbackQWebPage_MicroFocusChanged(this); };
	void Signal_PrintRequested(QWebFrame * frame) { callbackQWebPage_PrintRequested(this, frame); };
	void Signal_RecentlyAudibleChanged(bool recentlyAudible) { callbackQWebPage_RecentlyAudibleChanged(this, recentlyAudible); };
	void Signal_RepaintRequested(const QRect & dirtyRect) { callbackQWebPage_RepaintRequested(this, const_cast<QRect*>(&dirtyRect)); };
	void Signal_RestoreFrameStateRequested(QWebFrame * frame) { callbackQWebPage_RestoreFrameStateRequested(this, frame); };
	void Signal_SaveFrameStateRequested(QWebFrame * frame, QWebHistoryItem * item) { callbackQWebPage_SaveFrameStateRequested(this, frame, item); };
	void Signal_ScrollRequested(int dx, int dy, const QRect & rectToScroll) { callbackQWebPage_ScrollRequested(this, dx, dy, const_cast<QRect*>(&rectToScroll)); };
	void Signal_SelectionChanged() { callbackQWebPage_SelectionChanged(this); };
	void Signal_StatusBarMessage(const QString & text) { QByteArray t372ea0 = text.toUtf8(); QtWebKit_PackedString textPacked = { const_cast<char*>(t372ea0.prepend("WHITESPACE").constData()+10), t372ea0.size()-10 };callbackQWebPage_StatusBarMessage(this, textPacked); };
	void Signal_StatusBarVisibilityChangeRequested(bool visible) { callbackQWebPage_StatusBarVisibilityChangeRequested(this, visible); };
	void Signal_ToolBarVisibilityChangeRequested(bool visible) { callbackQWebPage_ToolBarVisibilityChangeRequested(this, visible); };
	void triggerAction(QWebPage::WebAction action, bool checked) { callbackQWebPage_TriggerAction(this, action, checked); };
	void Signal_UnsupportedContent(QNetworkReply * reply) { callbackQWebPage_UnsupportedContent(this, reply); };
	void Signal_ViewportChangeRequested() { callbackQWebPage_ViewportChangeRequested(this); };
	void Signal_WindowCloseRequested() { callbackQWebPage_WindowCloseRequested(this); };
	QString userAgentForUrl(const QUrl & url) const { return ({ QtWebKit_PackedString tempVal = callbackQWebPage_UserAgentForUrl(const_cast<void*>(static_cast<const void*>(this)), const_cast<QUrl*>(&url)); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	bool supportsExtension(QWebPage::Extension extension) const { return callbackQWebPage_SupportsExtension(const_cast<void*>(static_cast<const void*>(this)), extension) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQWebPage_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQWebPage_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQWebPage_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQWebPage_CustomEvent(this, event); };
	void deleteLater() { callbackQWebPage_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQWebPage_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQWebPage_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtWebKit_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQWebPage_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQWebPage_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQWebPage_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQWebPage*)

int QWebPage_QWebPage_QRegisterMetaType(){qRegisterMetaType<QWebPage*>(); return qRegisterMetaType<MyQWebPage*>();}

void* QWebPage_CreateStandardContextMenu(void* ptr)
{
	return static_cast<QWebPage*>(ptr)->createStandardContextMenu();
}

void* QWebPage_CreatePlugin(void* ptr, struct QtWebKit_PackedString classid, void* url, struct QtWebKit_PackedString paramNames, struct QtWebKit_PackedString paramValues)
{
	return static_cast<QWebPage*>(ptr)->createPlugin(QString::fromUtf8(classid.data, classid.len), *static_cast<QUrl*>(url), QString::fromUtf8(paramNames.data, paramNames.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(paramValues.data, paramValues.len).split("|", QString::SkipEmptyParts));
}

void* QWebPage_CreatePluginDefault(void* ptr, struct QtWebKit_PackedString classid, void* url, struct QtWebKit_PackedString paramNames, struct QtWebKit_PackedString paramValues)
{
		return static_cast<QWebPage*>(ptr)->QWebPage::createPlugin(QString::fromUtf8(classid.data, classid.len), *static_cast<QUrl*>(url), QString::fromUtf8(paramNames.data, paramNames.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(paramValues.data, paramValues.len).split("|", QString::SkipEmptyParts));
}

struct QtWebKit_PackedString QWebPage_ChooseFile(void* ptr, void* parentFrame, struct QtWebKit_PackedString suggestedFile)
{
	return ({ QByteArray t1a6e0e = static_cast<QWebPage*>(ptr)->chooseFile(static_cast<QWebFrame*>(parentFrame), QString::fromUtf8(suggestedFile.data, suggestedFile.len)).toUtf8(); QtWebKit_PackedString { const_cast<char*>(t1a6e0e.prepend("WHITESPACE").constData()+10), t1a6e0e.size()-10 }; });
}

struct QtWebKit_PackedString QWebPage_ChooseFileDefault(void* ptr, void* parentFrame, struct QtWebKit_PackedString suggestedFile)
{
		return ({ QByteArray t468d42 = static_cast<QWebPage*>(ptr)->QWebPage::chooseFile(static_cast<QWebFrame*>(parentFrame), QString::fromUtf8(suggestedFile.data, suggestedFile.len)).toUtf8(); QtWebKit_PackedString { const_cast<char*>(t468d42.prepend("WHITESPACE").constData()+10), t468d42.size()-10 }; });
}

void* QWebPage_CreateWindow(void* ptr, long long ty)
{
	return static_cast<QWebPage*>(ptr)->createWindow(static_cast<QWebPage::WebWindowType>(ty));
}

void* QWebPage_CreateWindowDefault(void* ptr, long long ty)
{
		return static_cast<QWebPage*>(ptr)->QWebPage::createWindow(static_cast<QWebPage::WebWindowType>(ty));
}

void* QWebPage_NewQWebPage(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQWebPage(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQWebPage(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQWebPage(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQWebPage(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQWebPage(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQWebPage(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQWebPage(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQWebPage(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQWebPage(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQWebPage(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQWebPage(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQWebPage(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQWebPage(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQWebPage(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQWebPage(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQWebPage(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQWebPage(static_cast<QWindow*>(parent));
	} else {
		return new MyQWebPage(static_cast<QObject*>(parent));
	}
}

char QWebPage_AcceptNavigationRequest(void* ptr, void* frame, void* request, long long ty)
{
	return static_cast<QWebPage*>(ptr)->acceptNavigationRequest(static_cast<QWebFrame*>(frame), *static_cast<QNetworkRequest*>(request), static_cast<QWebPage::NavigationType>(ty));
}

char QWebPage_AcceptNavigationRequestDefault(void* ptr, void* frame, void* request, long long ty)
{
		return static_cast<QWebPage*>(ptr)->QWebPage::acceptNavigationRequest(static_cast<QWebFrame*>(frame), *static_cast<QNetworkRequest*>(request), static_cast<QWebPage::NavigationType>(ty));
}

char QWebPage_EventDefault(void* ptr, void* ev)
{
		return static_cast<QWebPage*>(ptr)->QWebPage::event(static_cast<QEvent*>(ev));
}

char QWebPage_FindText(void* ptr, struct QtWebKit_PackedString subString, long long options)
{
	return static_cast<QWebPage*>(ptr)->findText(QString::fromUtf8(subString.data, subString.len), static_cast<QWebPage::FindFlag>(options));
}

char QWebPage_FocusNextPrevChild(void* ptr, char next)
{
	return static_cast<QWebPage*>(ptr)->focusNextPrevChild(next != 0);
}

char QWebPage_JavaScriptConfirm(void* ptr, void* frame, struct QtWebKit_PackedString msg)
{
	return static_cast<QWebPage*>(ptr)->javaScriptConfirm(static_cast<QWebFrame*>(frame), QString::fromUtf8(msg.data, msg.len));
}

char QWebPage_JavaScriptConfirmDefault(void* ptr, void* frame, struct QtWebKit_PackedString msg)
{
		return static_cast<QWebPage*>(ptr)->QWebPage::javaScriptConfirm(static_cast<QWebFrame*>(frame), QString::fromUtf8(msg.data, msg.len));
}

char QWebPage_JavaScriptPrompt(void* ptr, void* frame, struct QtWebKit_PackedString msg, struct QtWebKit_PackedString defaultValue, struct QtWebKit_PackedString result)
{
	return static_cast<QWebPage*>(ptr)->javaScriptPrompt(static_cast<QWebFrame*>(frame), QString::fromUtf8(msg.data, msg.len), QString::fromUtf8(defaultValue.data, defaultValue.len), new QString(QString::fromUtf8(result.data, result.len)));
}

char QWebPage_JavaScriptPromptDefault(void* ptr, void* frame, struct QtWebKit_PackedString msg, struct QtWebKit_PackedString defaultValue, struct QtWebKit_PackedString result)
{
		return static_cast<QWebPage*>(ptr)->QWebPage::javaScriptPrompt(static_cast<QWebFrame*>(frame), QString::fromUtf8(msg.data, msg.len), QString::fromUtf8(defaultValue.data, defaultValue.len), new QString(QString::fromUtf8(result.data, result.len)));
}

char QWebPage_ShouldInterruptJavaScript(void* ptr)
{
	return static_cast<QWebPage*>(ptr)->shouldInterruptJavaScript();
}

char QWebPage_ShouldInterruptJavaScriptDefault(void* ptr)
{
		return static_cast<QWebPage*>(ptr)->QWebPage::shouldInterruptJavaScript();
}

char QWebPage_SwallowContextMenuEvent(void* ptr, void* event)
{
	return static_cast<QWebPage*>(ptr)->swallowContextMenuEvent(static_cast<QContextMenuEvent*>(event));
}

void QWebPage_ConnectApplicationCacheQuotaExceeded(void* ptr)
{
	QObject::connect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)(QWebSecurityOrigin *, quint64, quint64)>(&QWebPage::applicationCacheQuotaExceeded), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)(QWebSecurityOrigin *, quint64, quint64)>(&MyQWebPage::Signal_ApplicationCacheQuotaExceeded));
}

void QWebPage_DisconnectApplicationCacheQuotaExceeded(void* ptr)
{
	QObject::disconnect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)(QWebSecurityOrigin *, quint64, quint64)>(&QWebPage::applicationCacheQuotaExceeded), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)(QWebSecurityOrigin *, quint64, quint64)>(&MyQWebPage::Signal_ApplicationCacheQuotaExceeded));
}

void QWebPage_ApplicationCacheQuotaExceeded(void* ptr, void* origin, unsigned long long defaultOriginQuota, unsigned long long totalSpaceNeeded)
{
	static_cast<QWebPage*>(ptr)->applicationCacheQuotaExceeded(static_cast<QWebSecurityOrigin*>(origin), defaultOriginQuota, totalSpaceNeeded);
}

void QWebPage_ConnectConsoleMessageReceived(void* ptr)
{
	QObject::connect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)(QWebPage::MessageSource, QWebPage::MessageLevel, const QString &, int, const QString &)>(&QWebPage::consoleMessageReceived), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)(QWebPage::MessageSource, QWebPage::MessageLevel, const QString &, int, const QString &)>(&MyQWebPage::Signal_ConsoleMessageReceived));
}

void QWebPage_DisconnectConsoleMessageReceived(void* ptr)
{
	QObject::disconnect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)(QWebPage::MessageSource, QWebPage::MessageLevel, const QString &, int, const QString &)>(&QWebPage::consoleMessageReceived), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)(QWebPage::MessageSource, QWebPage::MessageLevel, const QString &, int, const QString &)>(&MyQWebPage::Signal_ConsoleMessageReceived));
}

void QWebPage_ConsoleMessageReceived(void* ptr, long long source, long long level, struct QtWebKit_PackedString message, int lineNumber, struct QtWebKit_PackedString sourceID)
{
	static_cast<QWebPage*>(ptr)->consoleMessageReceived(static_cast<QWebPage::MessageSource>(source), static_cast<QWebPage::MessageLevel>(level), QString::fromUtf8(message.data, message.len), lineNumber, QString::fromUtf8(sourceID.data, sourceID.len));
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

void QWebPage_ConnectDatabaseQuotaExceeded(void* ptr)
{
	QObject::connect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)(QWebFrame *, QString)>(&QWebPage::databaseQuotaExceeded), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)(QWebFrame *, QString)>(&MyQWebPage::Signal_DatabaseQuotaExceeded));
}

void QWebPage_DisconnectDatabaseQuotaExceeded(void* ptr)
{
	QObject::disconnect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)(QWebFrame *, QString)>(&QWebPage::databaseQuotaExceeded), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)(QWebFrame *, QString)>(&MyQWebPage::Signal_DatabaseQuotaExceeded));
}

void QWebPage_DatabaseQuotaExceeded(void* ptr, void* frame, struct QtWebKit_PackedString databaseName)
{
	static_cast<QWebPage*>(ptr)->databaseQuotaExceeded(static_cast<QWebFrame*>(frame), QString::fromUtf8(databaseName.data, databaseName.len));
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

void QWebPage_ConnectFeaturePermissionRequestCanceled(void* ptr)
{
	QObject::connect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)(QWebFrame *, QWebPage::Feature)>(&QWebPage::featurePermissionRequestCanceled), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)(QWebFrame *, QWebPage::Feature)>(&MyQWebPage::Signal_FeaturePermissionRequestCanceled));
}

void QWebPage_DisconnectFeaturePermissionRequestCanceled(void* ptr)
{
	QObject::disconnect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)(QWebFrame *, QWebPage::Feature)>(&QWebPage::featurePermissionRequestCanceled), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)(QWebFrame *, QWebPage::Feature)>(&MyQWebPage::Signal_FeaturePermissionRequestCanceled));
}

void QWebPage_FeaturePermissionRequestCanceled(void* ptr, void* frame, long long feature)
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

void QWebPage_FeaturePermissionRequested(void* ptr, void* frame, long long feature)
{
	static_cast<QWebPage*>(ptr)->featurePermissionRequested(static_cast<QWebFrame*>(frame), static_cast<QWebPage::Feature>(feature));
}

void QWebPage_ConnectFocusedElementChanged(void* ptr)
{
	QObject::connect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)(const QWebElement &)>(&QWebPage::focusedElementChanged), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)(const QWebElement &)>(&MyQWebPage::Signal_FocusedElementChanged));
}

void QWebPage_DisconnectFocusedElementChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)(const QWebElement &)>(&QWebPage::focusedElementChanged), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)(const QWebElement &)>(&MyQWebPage::Signal_FocusedElementChanged));
}

void QWebPage_FocusedElementChanged(void* ptr, void* element)
{
	static_cast<QWebPage*>(ptr)->focusedElementChanged(*static_cast<QWebElement*>(element));
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

void QWebPage_JavaScriptAlert(void* ptr, void* frame, struct QtWebKit_PackedString msg)
{
	static_cast<QWebPage*>(ptr)->javaScriptAlert(static_cast<QWebFrame*>(frame), QString::fromUtf8(msg.data, msg.len));
}

void QWebPage_JavaScriptAlertDefault(void* ptr, void* frame, struct QtWebKit_PackedString msg)
{
		static_cast<QWebPage*>(ptr)->QWebPage::javaScriptAlert(static_cast<QWebFrame*>(frame), QString::fromUtf8(msg.data, msg.len));
}

void QWebPage_JavaScriptConsoleMessage(void* ptr, struct QtWebKit_PackedString message, int lineNumber, struct QtWebKit_PackedString sourceID)
{
	static_cast<QWebPage*>(ptr)->javaScriptConsoleMessage(QString::fromUtf8(message.data, message.len), lineNumber, QString::fromUtf8(sourceID.data, sourceID.len));
}

void QWebPage_JavaScriptConsoleMessageDefault(void* ptr, struct QtWebKit_PackedString message, int lineNumber, struct QtWebKit_PackedString sourceID)
{
		static_cast<QWebPage*>(ptr)->QWebPage::javaScriptConsoleMessage(QString::fromUtf8(message.data, message.len), lineNumber, QString::fromUtf8(sourceID.data, sourceID.len));
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

void QWebPage_LinkHovered(void* ptr, struct QtWebKit_PackedString link, struct QtWebKit_PackedString title, struct QtWebKit_PackedString textContent)
{
	static_cast<QWebPage*>(ptr)->linkHovered(QString::fromUtf8(link.data, link.len), QString::fromUtf8(title.data, title.len), QString::fromUtf8(textContent.data, textContent.len));
}

void QWebPage_ConnectLoadFinished(void* ptr)
{
	QObject::connect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)(bool)>(&QWebPage::loadFinished), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)(bool)>(&MyQWebPage::Signal_LoadFinished));
}

void QWebPage_DisconnectLoadFinished(void* ptr)
{
	QObject::disconnect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)(bool)>(&QWebPage::loadFinished), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)(bool)>(&MyQWebPage::Signal_LoadFinished));
}

void QWebPage_LoadFinished(void* ptr, char ok)
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

void QWebPage_ConnectMenuBarVisibilityChangeRequested(void* ptr)
{
	QObject::connect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)(bool)>(&QWebPage::menuBarVisibilityChangeRequested), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)(bool)>(&MyQWebPage::Signal_MenuBarVisibilityChangeRequested));
}

void QWebPage_DisconnectMenuBarVisibilityChangeRequested(void* ptr)
{
	QObject::disconnect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)(bool)>(&QWebPage::menuBarVisibilityChangeRequested), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)(bool)>(&MyQWebPage::Signal_MenuBarVisibilityChangeRequested));
}

void QWebPage_MenuBarVisibilityChangeRequested(void* ptr, char visible)
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

void QWebPage_ConnectRecentlyAudibleChanged(void* ptr)
{
	QObject::connect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)(bool)>(&QWebPage::recentlyAudibleChanged), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)(bool)>(&MyQWebPage::Signal_RecentlyAudibleChanged));
}

void QWebPage_DisconnectRecentlyAudibleChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)(bool)>(&QWebPage::recentlyAudibleChanged), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)(bool)>(&MyQWebPage::Signal_RecentlyAudibleChanged));
}

void QWebPage_RecentlyAudibleChanged(void* ptr, char recentlyAudible)
{
	static_cast<QWebPage*>(ptr)->recentlyAudibleChanged(recentlyAudible != 0);
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

void QWebPage_SetContentEditable(void* ptr, char editable)
{
	static_cast<QWebPage*>(ptr)->setContentEditable(editable != 0);
}

void QWebPage_SetFeaturePermission(void* ptr, void* frame, long long feature, long long policy)
{
	static_cast<QWebPage*>(ptr)->setFeaturePermission(static_cast<QWebFrame*>(frame), static_cast<QWebPage::Feature>(feature), static_cast<QWebPage::PermissionPolicy>(policy));
}

void QWebPage_SetForwardUnsupportedContent(void* ptr, char forward)
{
	static_cast<QWebPage*>(ptr)->setForwardUnsupportedContent(forward != 0);
}

void QWebPage_SetLinkDelegationPolicy(void* ptr, long long policy)
{
	static_cast<QWebPage*>(ptr)->setLinkDelegationPolicy(static_cast<QWebPage::LinkDelegationPolicy>(policy));
}

void QWebPage_SetNetworkAccessManager(void* ptr, void* manager)
{
	static_cast<QWebPage*>(ptr)->setNetworkAccessManager(static_cast<QNetworkAccessManager*>(manager));
}

void QWebPage_SetPalette(void* ptr, void* palette)
{
	static_cast<QWebPage*>(ptr)->setPalette(*static_cast<QPalette*>(palette));
}

void QWebPage_SetPluginFactory(void* ptr, void* factory)
{
	static_cast<QWebPage*>(ptr)->setPluginFactory(static_cast<QWebPluginFactory*>(factory));
}

void QWebPage_SetView(void* ptr, void* view)
{
	static_cast<QWebPage*>(ptr)->setView(static_cast<QWidget*>(view));
}

void QWebPage_SetVisibilityState(void* ptr, long long vvi)
{
	static_cast<QWebPage*>(ptr)->setVisibilityState(static_cast<QWebPage::VisibilityState>(vvi));
}

void QWebPage_ConnectStatusBarMessage(void* ptr)
{
	QObject::connect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)(const QString &)>(&QWebPage::statusBarMessage), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)(const QString &)>(&MyQWebPage::Signal_StatusBarMessage));
}

void QWebPage_DisconnectStatusBarMessage(void* ptr)
{
	QObject::disconnect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)(const QString &)>(&QWebPage::statusBarMessage), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)(const QString &)>(&MyQWebPage::Signal_StatusBarMessage));
}

void QWebPage_StatusBarMessage(void* ptr, struct QtWebKit_PackedString text)
{
	static_cast<QWebPage*>(ptr)->statusBarMessage(QString::fromUtf8(text.data, text.len));
}

void QWebPage_ConnectStatusBarVisibilityChangeRequested(void* ptr)
{
	QObject::connect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)(bool)>(&QWebPage::statusBarVisibilityChangeRequested), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)(bool)>(&MyQWebPage::Signal_StatusBarVisibilityChangeRequested));
}

void QWebPage_DisconnectStatusBarVisibilityChangeRequested(void* ptr)
{
	QObject::disconnect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)(bool)>(&QWebPage::statusBarVisibilityChangeRequested), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)(bool)>(&MyQWebPage::Signal_StatusBarVisibilityChangeRequested));
}

void QWebPage_StatusBarVisibilityChangeRequested(void* ptr, char visible)
{
	static_cast<QWebPage*>(ptr)->statusBarVisibilityChangeRequested(visible != 0);
}

void QWebPage_ConnectToolBarVisibilityChangeRequested(void* ptr)
{
	QObject::connect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)(bool)>(&QWebPage::toolBarVisibilityChangeRequested), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)(bool)>(&MyQWebPage::Signal_ToolBarVisibilityChangeRequested));
}

void QWebPage_DisconnectToolBarVisibilityChangeRequested(void* ptr)
{
	QObject::disconnect(static_cast<QWebPage*>(ptr), static_cast<void (QWebPage::*)(bool)>(&QWebPage::toolBarVisibilityChangeRequested), static_cast<MyQWebPage*>(ptr), static_cast<void (MyQWebPage::*)(bool)>(&MyQWebPage::Signal_ToolBarVisibilityChangeRequested));
}

void QWebPage_ToolBarVisibilityChangeRequested(void* ptr, char visible)
{
	static_cast<QWebPage*>(ptr)->toolBarVisibilityChangeRequested(visible != 0);
}

void QWebPage_TriggerAction(void* ptr, long long action, char checked)
{
	static_cast<QWebPage*>(ptr)->triggerAction(static_cast<QWebPage::WebAction>(action), checked != 0);
}

void QWebPage_TriggerActionDefault(void* ptr, long long action, char checked)
{
		static_cast<QWebPage*>(ptr)->QWebPage::triggerAction(static_cast<QWebPage::WebAction>(action), checked != 0);
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

long long QWebPage_LinkDelegationPolicy(void* ptr)
{
	return static_cast<QWebPage*>(ptr)->linkDelegationPolicy();
}

void* QWebPage_Action(void* ptr, long long action)
{
	return static_cast<QWebPage*>(ptr)->action(static_cast<QWebPage::WebAction>(action));
}

void* QWebPage_CustomAction(void* ptr, int action)
{
	return static_cast<QWebPage*>(ptr)->customAction(action);
}

void* QWebPage_NetworkAccessManager(void* ptr)
{
	return static_cast<QWebPage*>(ptr)->networkAccessManager();
}

void* QWebPage_Palette(void* ptr)
{
	return new QPalette(static_cast<QWebPage*>(ptr)->palette());
}

void* QWebPage_PreferredContentsSize(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QWebPage*>(ptr)->preferredContentsSize(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QWebPage_ViewportSize(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QWebPage*>(ptr)->viewportSize(); new QSize(tmpValue.width(), tmpValue.height()); });
}

struct QtWebKit_PackedString QWebPage_SelectedHtml(void* ptr)
{
	return ({ QByteArray t204f0c = static_cast<QWebPage*>(ptr)->selectedHtml().toUtf8(); QtWebKit_PackedString { const_cast<char*>(t204f0c.prepend("WHITESPACE").constData()+10), t204f0c.size()-10 }; });
}

struct QtWebKit_PackedString QWebPage_SelectedText(void* ptr)
{
	return ({ QByteArray tde1ecb = static_cast<QWebPage*>(ptr)->selectedText().toUtf8(); QtWebKit_PackedString { const_cast<char*>(tde1ecb.prepend("WHITESPACE").constData()+10), tde1ecb.size()-10 }; });
}

struct QtWebKit_PackedString QWebPage_UserAgentForUrl(void* ptr, void* url)
{
	return ({ QByteArray t3bd945 = static_cast<QWebPage*>(ptr)->userAgentForUrl(*static_cast<QUrl*>(url)).toUtf8(); QtWebKit_PackedString { const_cast<char*>(t3bd945.prepend("WHITESPACE").constData()+10), t3bd945.size()-10 }; });
}

struct QtWebKit_PackedString QWebPage_UserAgentForUrlDefault(void* ptr, void* url)
{
		return ({ QByteArray t25b0ae = static_cast<QWebPage*>(ptr)->QWebPage::userAgentForUrl(*static_cast<QUrl*>(url)).toUtf8(); QtWebKit_PackedString { const_cast<char*>(t25b0ae.prepend("WHITESPACE").constData()+10), t25b0ae.size()-10 }; });
}

struct QtWebKit_PackedString QWebPage_SupportedContentTypes(void* ptr)
{
	return ({ QByteArray tec3baf = static_cast<QWebPage*>(ptr)->supportedContentTypes().join("|").toUtf8(); QtWebKit_PackedString { const_cast<char*>(tec3baf.prepend("WHITESPACE").constData()+10), tec3baf.size()-10 }; });
}

void* QWebPage_UndoStack(void* ptr)
{
	return static_cast<QWebPage*>(ptr)->undoStack();
}

void* QWebPage_InputMethodQuery(void* ptr, long long property)
{
	return new QVariant(static_cast<QWebPage*>(ptr)->inputMethodQuery(static_cast<Qt::InputMethodQuery>(property)));
}

void* QWebPage_CurrentFrame(void* ptr)
{
	return static_cast<QWebPage*>(ptr)->currentFrame();
}

void* QWebPage_FrameAt(void* ptr, void* pos)
{
	return static_cast<QWebPage*>(ptr)->frameAt(*static_cast<QPoint*>(pos));
}

void* QWebPage_MainFrame(void* ptr)
{
	return static_cast<QWebPage*>(ptr)->mainFrame();
}

void* QWebPage_History(void* ptr)
{
	return static_cast<QWebPage*>(ptr)->history();
}

void* QWebPage_PluginFactory(void* ptr)
{
	return static_cast<QWebPage*>(ptr)->pluginFactory();
}

void* QWebPage_Settings(void* ptr)
{
	return static_cast<QWebPage*>(ptr)->settings();
}

void* QWebPage_View(void* ptr)
{
	return static_cast<QWebPage*>(ptr)->view();
}

long long QWebPage_VisibilityState(void* ptr)
{
	return static_cast<QWebPage*>(ptr)->visibilityState();
}

char QWebPage_ForwardUnsupportedContent(void* ptr)
{
	return static_cast<QWebPage*>(ptr)->forwardUnsupportedContent();
}

char QWebPage_HasSelection(void* ptr)
{
	return static_cast<QWebPage*>(ptr)->hasSelection();
}

char QWebPage_IsContentEditable(void* ptr)
{
	return static_cast<QWebPage*>(ptr)->isContentEditable();
}

char QWebPage_IsModified(void* ptr)
{
	return static_cast<QWebPage*>(ptr)->isModified();
}

char QWebPage_RecentlyAudible(void* ptr)
{
	return static_cast<QWebPage*>(ptr)->recentlyAudible();
}

char QWebPage_SupportsContentType(void* ptr, struct QtWebKit_PackedString mimeType)
{
	return static_cast<QWebPage*>(ptr)->supportsContentType(QString::fromUtf8(mimeType.data, mimeType.len));
}

char QWebPage_SupportsExtension(void* ptr, long long extension)
{
	return static_cast<QWebPage*>(ptr)->supportsExtension(static_cast<QWebPage::Extension>(extension));
}

char QWebPage_SupportsExtensionDefault(void* ptr, long long extension)
{
		return static_cast<QWebPage*>(ptr)->QWebPage::supportsExtension(static_cast<QWebPage::Extension>(extension));
}

unsigned long long QWebPage_BytesReceived(void* ptr)
{
	return static_cast<QWebPage*>(ptr)->bytesReceived();
}

unsigned long long QWebPage_TotalBytes(void* ptr)
{
	return static_cast<QWebPage*>(ptr)->totalBytes();
}

void QWebPage_SetActualVisibleContentRect(void* ptr, void* rect)
{
	static_cast<QWebPage*>(ptr)->setActualVisibleContentRect(*static_cast<QRect*>(rect));
}

void QWebPage_SetPreferredContentsSize(void* ptr, void* size)
{
	static_cast<QWebPage*>(ptr)->setPreferredContentsSize(*static_cast<QSize*>(size));
}

void QWebPage_SetViewportSize(void* ptr, void* size)
{
	static_cast<QWebPage*>(ptr)->setViewportSize(*static_cast<QSize*>(size));
}

void* QWebPage___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QWebPage___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QWebPage___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QWebPage___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QWebPage___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebPage___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QWebPage___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QWebPage___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebPage___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QWebPage___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QWebPage___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebPage___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QWebPage___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QWebPage___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebPage___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

char QWebPage_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QWebPage*>(ptr)->QWebPage::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QWebPage_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QWebPage*>(ptr)->QWebPage::childEvent(static_cast<QChildEvent*>(event));
}

void QWebPage_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QWebPage*>(ptr)->QWebPage::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWebPage_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QWebPage*>(ptr)->QWebPage::customEvent(static_cast<QEvent*>(event));
}

void QWebPage_DeleteLaterDefault(void* ptr)
{
		static_cast<QWebPage*>(ptr)->QWebPage::deleteLater();
}

void QWebPage_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QWebPage*>(ptr)->QWebPage::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWebPage_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QWebPage*>(ptr)->QWebPage::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QWebPage_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QWebPage*>(ptr)->QWebPage::metaObject());
}

class MyQWebPluginFactory: public QWebPluginFactory
{
public:
	void refreshPlugins() { callbackQWebPluginFactory_RefreshPlugins(this); };
	 ~MyQWebPluginFactory() { callbackQWebPluginFactory_DestroyQWebPluginFactory(this); };
	QObject * create(const QString & mimeType, const QUrl & url, const QStringList & argumentNames, const QStringList & argumentValues) const { QByteArray t3313b8 = mimeType.toUtf8(); QtWebKit_PackedString mimeTypePacked = { const_cast<char*>(t3313b8.prepend("WHITESPACE").constData()+10), t3313b8.size()-10 };QByteArray t92dffc = argumentNames.join("|").toUtf8(); QtWebKit_PackedString argumentNamesPacked = { const_cast<char*>(t92dffc.prepend("WHITESPACE").constData()+10), t92dffc.size()-10 };QByteArray t6a576b = argumentValues.join("|").toUtf8(); QtWebKit_PackedString argumentValuesPacked = { const_cast<char*>(t6a576b.prepend("WHITESPACE").constData()+10), t6a576b.size()-10 };return static_cast<QObject*>(callbackQWebPluginFactory_Create(const_cast<void*>(static_cast<const void*>(this)), mimeTypePacked, const_cast<QUrl*>(&url), argumentNamesPacked, argumentValuesPacked)); };
	bool event(QEvent * e) { return callbackQWebPluginFactory_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQWebPluginFactory_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQWebPluginFactory_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQWebPluginFactory_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQWebPluginFactory_CustomEvent(this, event); };
	void deleteLater() { callbackQWebPluginFactory_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQWebPluginFactory_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQWebPluginFactory_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtWebKit_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQWebPluginFactory_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQWebPluginFactory_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQWebPluginFactory_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQWebPluginFactory*)

int QWebPluginFactory_QWebPluginFactory_QRegisterMetaType(){qRegisterMetaType<QWebPluginFactory*>(); return qRegisterMetaType<MyQWebPluginFactory*>();}

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

void QWebPluginFactory_DestroyQWebPluginFactoryDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QWebPluginFactory_Create(void* ptr, struct QtWebKit_PackedString mimeType, void* url, struct QtWebKit_PackedString argumentNames, struct QtWebKit_PackedString argumentValues)
{
	return static_cast<QWebPluginFactory*>(ptr)->create(QString::fromUtf8(mimeType.data, mimeType.len), *static_cast<QUrl*>(url), QString::fromUtf8(argumentNames.data, argumentNames.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(argumentValues.data, argumentValues.len).split("|", QString::SkipEmptyParts));
}

void* QWebPluginFactory___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QWebPluginFactory___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QWebPluginFactory___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QWebPluginFactory___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QWebPluginFactory___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebPluginFactory___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QWebPluginFactory___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QWebPluginFactory___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebPluginFactory___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QWebPluginFactory___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QWebPluginFactory___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebPluginFactory___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QWebPluginFactory___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QWebPluginFactory___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebPluginFactory___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

char QWebPluginFactory_EventDefault(void* ptr, void* e)
{
		return static_cast<QWebPluginFactory*>(ptr)->QWebPluginFactory::event(static_cast<QEvent*>(e));
}

char QWebPluginFactory_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QWebPluginFactory*>(ptr)->QWebPluginFactory::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QWebPluginFactory_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QWebPluginFactory*>(ptr)->QWebPluginFactory::childEvent(static_cast<QChildEvent*>(event));
}

void QWebPluginFactory_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QWebPluginFactory*>(ptr)->QWebPluginFactory::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWebPluginFactory_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QWebPluginFactory*>(ptr)->QWebPluginFactory::customEvent(static_cast<QEvent*>(event));
}

void QWebPluginFactory_DeleteLaterDefault(void* ptr)
{
		static_cast<QWebPluginFactory*>(ptr)->QWebPluginFactory::deleteLater();
}

void QWebPluginFactory_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QWebPluginFactory*>(ptr)->QWebPluginFactory::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWebPluginFactory_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QWebPluginFactory*>(ptr)->QWebPluginFactory::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QWebPluginFactory_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QWebPluginFactory*>(ptr)->QWebPluginFactory::metaObject());
}

struct QtWebKit_PackedList QWebSecurityOrigin_QWebSecurityOrigin_AllOrigins()
{
	return ({ QList<QWebSecurityOrigin>* tmpValue = new QList<QWebSecurityOrigin>(QWebSecurityOrigin::allOrigins()); QtWebKit_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtWebKit_PackedString QWebSecurityOrigin_QWebSecurityOrigin_LocalSchemes()
{
	return ({ QByteArray t24a558 = QWebSecurityOrigin::localSchemes().join("|").toUtf8(); QtWebKit_PackedString { const_cast<char*>(t24a558.prepend("WHITESPACE").constData()+10), t24a558.size()-10 }; });
}

void* QWebSecurityOrigin_NewQWebSecurityOrigin(void* url)
{
	return new QWebSecurityOrigin(*static_cast<QUrl*>(url));
}

void* QWebSecurityOrigin_NewQWebSecurityOrigin2(void* other)
{
	return new QWebSecurityOrigin(*static_cast<QWebSecurityOrigin*>(other));
}

void QWebSecurityOrigin_AddAccessWhitelistEntry(void* ptr, struct QtWebKit_PackedString scheme, struct QtWebKit_PackedString host, long long subdomainSetting)
{
	static_cast<QWebSecurityOrigin*>(ptr)->addAccessWhitelistEntry(QString::fromUtf8(scheme.data, scheme.len), QString::fromUtf8(host.data, host.len), static_cast<QWebSecurityOrigin::SubdomainSetting>(subdomainSetting));
}

void QWebSecurityOrigin_QWebSecurityOrigin_AddLocalScheme(struct QtWebKit_PackedString scheme)
{
	QWebSecurityOrigin::addLocalScheme(QString::fromUtf8(scheme.data, scheme.len));
}

void QWebSecurityOrigin_RemoveAccessWhitelistEntry(void* ptr, struct QtWebKit_PackedString scheme, struct QtWebKit_PackedString host, long long subdomainSetting)
{
	static_cast<QWebSecurityOrigin*>(ptr)->removeAccessWhitelistEntry(QString::fromUtf8(scheme.data, scheme.len), QString::fromUtf8(host.data, host.len), static_cast<QWebSecurityOrigin::SubdomainSetting>(subdomainSetting));
}

void QWebSecurityOrigin_QWebSecurityOrigin_RemoveLocalScheme(struct QtWebKit_PackedString scheme)
{
	QWebSecurityOrigin::removeLocalScheme(QString::fromUtf8(scheme.data, scheme.len));
}

void QWebSecurityOrigin_SetApplicationCacheQuota(void* ptr, long long quota)
{
	static_cast<QWebSecurityOrigin*>(ptr)->setApplicationCacheQuota(quota);
}

void QWebSecurityOrigin_SetDatabaseQuota(void* ptr, long long quota)
{
	static_cast<QWebSecurityOrigin*>(ptr)->setDatabaseQuota(quota);
}

void QWebSecurityOrigin_DestroyQWebSecurityOrigin(void* ptr)
{
	static_cast<QWebSecurityOrigin*>(ptr)->~QWebSecurityOrigin();
}

struct QtWebKit_PackedList QWebSecurityOrigin_Databases(void* ptr)
{
	return ({ QList<QWebDatabase>* tmpValue = new QList<QWebDatabase>(static_cast<QWebSecurityOrigin*>(ptr)->databases()); QtWebKit_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtWebKit_PackedString QWebSecurityOrigin_Host(void* ptr)
{
	return ({ QByteArray t9cde96 = static_cast<QWebSecurityOrigin*>(ptr)->host().toUtf8(); QtWebKit_PackedString { const_cast<char*>(t9cde96.prepend("WHITESPACE").constData()+10), t9cde96.size()-10 }; });
}

struct QtWebKit_PackedString QWebSecurityOrigin_Scheme(void* ptr)
{
	return ({ QByteArray t3c5acb = static_cast<QWebSecurityOrigin*>(ptr)->scheme().toUtf8(); QtWebKit_PackedString { const_cast<char*>(t3c5acb.prepend("WHITESPACE").constData()+10), t3c5acb.size()-10 }; });
}

int QWebSecurityOrigin_Port(void* ptr)
{
	return static_cast<QWebSecurityOrigin*>(ptr)->port();
}

long long QWebSecurityOrigin_DatabaseQuota(void* ptr)
{
	return static_cast<QWebSecurityOrigin*>(ptr)->databaseQuota();
}

long long QWebSecurityOrigin_DatabaseUsage(void* ptr)
{
	return static_cast<QWebSecurityOrigin*>(ptr)->databaseUsage();
}

void* QWebSecurityOrigin___allOrigins_atList(void* ptr, int i)
{
	return new QWebSecurityOrigin(static_cast<QList<QWebSecurityOrigin>*>(ptr)->at(i));
}

void QWebSecurityOrigin___allOrigins_setList(void* ptr, void* i)
{
	static_cast<QList<QWebSecurityOrigin>*>(ptr)->append(*static_cast<QWebSecurityOrigin*>(i));
}

void* QWebSecurityOrigin___allOrigins_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QWebSecurityOrigin>;
}

void* QWebSecurityOrigin___databases_atList(void* ptr, int i)
{
	return new QWebDatabase(static_cast<QList<QWebDatabase>*>(ptr)->at(i));
}

void QWebSecurityOrigin___databases_setList(void* ptr, void* i)
{
	static_cast<QList<QWebDatabase>*>(ptr)->append(*static_cast<QWebDatabase*>(i));
}

void* QWebSecurityOrigin___databases_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QWebDatabase>;
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

int QWebSettings_MediaSourceEnabled_Type()
{
	return QWebSettings::MediaSourceEnabled;
}

int QWebSettings_MediaEnabled_Type()
{
	return QWebSettings::MediaEnabled;
}

int QWebSettings_WebSecurityEnabled_Type()
{
	return QWebSettings::WebSecurityEnabled;
}

int QWebSettings_FullScreenSupportEnabled_Type()
{
	return QWebSettings::FullScreenSupportEnabled;
}

void* QWebSettings_QWebSettings_IconForUrl(void* url)
{
	return new QIcon(QWebSettings::iconForUrl(*static_cast<QUrl*>(url)));
}

void* QWebSettings_QWebSettings_WebGraphic(long long ty)
{
	return new QPixmap(QWebSettings::webGraphic(static_cast<QWebSettings::WebGraphic>(ty)));
}

struct QtWebKit_PackedString QWebSettings_QWebSettings_IconDatabasePath()
{
	return ({ QByteArray t3ba260 = QWebSettings::iconDatabasePath().toUtf8(); QtWebKit_PackedString { const_cast<char*>(t3ba260.prepend("WHITESPACE").constData()+10), t3ba260.size()-10 }; });
}

struct QtWebKit_PackedString QWebSettings_QWebSettings_OfflineStoragePath()
{
	return ({ QByteArray t2a2b6b = QWebSettings::offlineStoragePath().toUtf8(); QtWebKit_PackedString { const_cast<char*>(t2a2b6b.prepend("WHITESPACE").constData()+10), t2a2b6b.size()-10 }; });
}

struct QtWebKit_PackedString QWebSettings_QWebSettings_OfflineWebApplicationCachePath()
{
	return ({ QByteArray t4b10cf = QWebSettings::offlineWebApplicationCachePath().toUtf8(); QtWebKit_PackedString { const_cast<char*>(t4b10cf.prepend("WHITESPACE").constData()+10), t4b10cf.size()-10 }; });
}

struct QtWebKit_PackedString QWebSettings_QWebSettings_PluginSearchPaths()
{
	return ({ QByteArray tb6c464 = QWebSettings::pluginSearchPaths().join("|").toUtf8(); QtWebKit_PackedString { const_cast<char*>(tb6c464.prepend("WHITESPACE").constData()+10), tb6c464.size()-10 }; });
}

void* QWebSettings_QWebSettings_GlobalSettings()
{
	return QWebSettings::globalSettings();
}

int QWebSettings_QWebSettings_MaximumPagesInCache()
{
	return QWebSettings::maximumPagesInCache();
}

long long QWebSettings_QWebSettings_OfflineStorageDefaultQuota()
{
	return QWebSettings::offlineStorageDefaultQuota();
}

long long QWebSettings_QWebSettings_OfflineWebApplicationCacheQuota()
{
	return QWebSettings::offlineWebApplicationCacheQuota();
}

void QWebSettings_QWebSettings_ClearIconDatabase()
{
	QWebSettings::clearIconDatabase();
}

void QWebSettings_QWebSettings_ClearMemoryCaches()
{
	QWebSettings::clearMemoryCaches();
}

void QWebSettings_QWebSettings_EnablePersistentStorage(struct QtWebKit_PackedString path)
{
	QWebSettings::enablePersistentStorage(QString::fromUtf8(path.data, path.len));
}

void QWebSettings_ResetAttribute(void* ptr, long long attribute)
{
	static_cast<QWebSettings*>(ptr)->resetAttribute(static_cast<QWebSettings::WebAttribute>(attribute));
}

void QWebSettings_ResetFontFamily(void* ptr, long long which)
{
	static_cast<QWebSettings*>(ptr)->resetFontFamily(static_cast<QWebSettings::FontFamily>(which));
}

void QWebSettings_ResetFontSize(void* ptr, long long ty)
{
	static_cast<QWebSettings*>(ptr)->resetFontSize(static_cast<QWebSettings::FontSize>(ty));
}

void QWebSettings_SetAttribute(void* ptr, long long attribute, char on)
{
	static_cast<QWebSettings*>(ptr)->setAttribute(static_cast<QWebSettings::WebAttribute>(attribute), on != 0);
}

void QWebSettings_SetCSSMediaType(void* ptr, struct QtWebKit_PackedString ty)
{
	static_cast<QWebSettings*>(ptr)->setCSSMediaType(QString::fromUtf8(ty.data, ty.len));
}

void QWebSettings_SetDefaultTextEncoding(void* ptr, struct QtWebKit_PackedString encoding)
{
	static_cast<QWebSettings*>(ptr)->setDefaultTextEncoding(QString::fromUtf8(encoding.data, encoding.len));
}

void QWebSettings_SetFontFamily(void* ptr, long long which, struct QtWebKit_PackedString family)
{
	static_cast<QWebSettings*>(ptr)->setFontFamily(static_cast<QWebSettings::FontFamily>(which), QString::fromUtf8(family.data, family.len));
}

void QWebSettings_SetFontSize(void* ptr, long long ty, int size)
{
	static_cast<QWebSettings*>(ptr)->setFontSize(static_cast<QWebSettings::FontSize>(ty), size);
}

void QWebSettings_QWebSettings_SetIconDatabasePath(struct QtWebKit_PackedString path)
{
	QWebSettings::setIconDatabasePath(QString::fromUtf8(path.data, path.len));
}

void QWebSettings_SetLocalStoragePath(void* ptr, struct QtWebKit_PackedString path)
{
	static_cast<QWebSettings*>(ptr)->setLocalStoragePath(QString::fromUtf8(path.data, path.len));
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
	QWebSettings::setOfflineStorageDefaultQuota(maximumSize);
}

void QWebSettings_QWebSettings_SetOfflineStoragePath(struct QtWebKit_PackedString path)
{
	QWebSettings::setOfflineStoragePath(QString::fromUtf8(path.data, path.len));
}

void QWebSettings_QWebSettings_SetOfflineWebApplicationCachePath(struct QtWebKit_PackedString path)
{
	QWebSettings::setOfflineWebApplicationCachePath(QString::fromUtf8(path.data, path.len));
}

void QWebSettings_QWebSettings_SetOfflineWebApplicationCacheQuota(long long maximumSize)
{
	QWebSettings::setOfflineWebApplicationCacheQuota(maximumSize);
}

void QWebSettings_QWebSettings_SetPluginSearchPaths(struct QtWebKit_PackedString paths)
{
	QWebSettings::setPluginSearchPaths(QString::fromUtf8(paths.data, paths.len).split("|", QString::SkipEmptyParts));
}

void QWebSettings_SetThirdPartyCookiePolicy(void* ptr, long long policy)
{
	static_cast<QWebSettings*>(ptr)->setThirdPartyCookiePolicy(static_cast<QWebSettings::ThirdPartyCookiePolicy>(policy));
}

void QWebSettings_SetUserStyleSheetUrl(void* ptr, void* location)
{
	static_cast<QWebSettings*>(ptr)->setUserStyleSheetUrl(*static_cast<QUrl*>(location));
}

void QWebSettings_QWebSettings_SetWebGraphic(long long ty, void* graphic)
{
	QWebSettings::setWebGraphic(static_cast<QWebSettings::WebGraphic>(ty), *static_cast<QPixmap*>(graphic));
}

struct QtWebKit_PackedString QWebSettings_CssMediaType(void* ptr)
{
	return ({ QByteArray tdd08cc = static_cast<QWebSettings*>(ptr)->cssMediaType().toUtf8(); QtWebKit_PackedString { const_cast<char*>(tdd08cc.prepend("WHITESPACE").constData()+10), tdd08cc.size()-10 }; });
}

struct QtWebKit_PackedString QWebSettings_DefaultTextEncoding(void* ptr)
{
	return ({ QByteArray t935a67 = static_cast<QWebSettings*>(ptr)->defaultTextEncoding().toUtf8(); QtWebKit_PackedString { const_cast<char*>(t935a67.prepend("WHITESPACE").constData()+10), t935a67.size()-10 }; });
}

struct QtWebKit_PackedString QWebSettings_FontFamily(void* ptr, long long which)
{
	return ({ QByteArray tc25ac0 = static_cast<QWebSettings*>(ptr)->fontFamily(static_cast<QWebSettings::FontFamily>(which)).toUtf8(); QtWebKit_PackedString { const_cast<char*>(tc25ac0.prepend("WHITESPACE").constData()+10), tc25ac0.size()-10 }; });
}

struct QtWebKit_PackedString QWebSettings_LocalStoragePath(void* ptr)
{
	return ({ QByteArray t2bc14b = static_cast<QWebSettings*>(ptr)->localStoragePath().toUtf8(); QtWebKit_PackedString { const_cast<char*>(t2bc14b.prepend("WHITESPACE").constData()+10), t2bc14b.size()-10 }; });
}

void* QWebSettings_UserStyleSheetUrl(void* ptr)
{
	return new QUrl(static_cast<QWebSettings*>(ptr)->userStyleSheetUrl());
}

long long QWebSettings_ThirdPartyCookiePolicy(void* ptr)
{
	return static_cast<QWebSettings*>(ptr)->thirdPartyCookiePolicy();
}

char QWebSettings_TestAttribute(void* ptr, long long attribute)
{
	return static_cast<QWebSettings*>(ptr)->testAttribute(static_cast<QWebSettings::WebAttribute>(attribute));
}

int QWebSettings_FontSize(void* ptr, long long ty)
{
	return static_cast<QWebSettings*>(ptr)->fontSize(static_cast<QWebSettings::FontSize>(ty));
}

class MyQWebView: public QWebView
{
public:
	MyQWebView(QWidget *parent = Q_NULLPTR) : QWebView(parent) {QWebView_QWebView_QRegisterMetaType();};
	QWebView * createWindow(QWebPage::WebWindowType ty) { return static_cast<QWebView*>(callbackQWebView_CreateWindow(this, ty)); };
	bool event(QEvent * e) { return callbackQWebView_Event(this, e) != 0; };
	bool focusNextPrevChild(bool next) { return callbackQWebView_FocusNextPrevChild(this, next) != 0; };
	void back() { callbackQWebView_Back(this); };
	void changeEvent(QEvent * e) { callbackQWebView_ChangeEvent(this, e); };
	void contextMenuEvent(QContextMenuEvent * ev) { callbackQWebView_ContextMenuEvent(this, ev); };
	void dragEnterEvent(QDragEnterEvent * ev) { callbackQWebView_DragEnterEvent(this, ev); };
	void dragLeaveEvent(QDragLeaveEvent * ev) { callbackQWebView_DragLeaveEvent(this, ev); };
	void dragMoveEvent(QDragMoveEvent * ev) { callbackQWebView_DragMoveEvent(this, ev); };
	void dropEvent(QDropEvent * ev) { callbackQWebView_DropEvent(this, ev); };
	void focusInEvent(QFocusEvent * ev) { callbackQWebView_FocusInEvent(this, ev); };
	void focusOutEvent(QFocusEvent * ev) { callbackQWebView_FocusOutEvent(this, ev); };
	void forward() { callbackQWebView_Forward(this); };
	void Signal_IconChanged() { callbackQWebView_IconChanged(this); };
	void inputMethodEvent(QInputMethodEvent * e) { callbackQWebView_InputMethodEvent(this, e); };
	void keyPressEvent(QKeyEvent * ev) { callbackQWebView_KeyPressEvent(this, ev); };
	void keyReleaseEvent(QKeyEvent * ev) { callbackQWebView_KeyReleaseEvent(this, ev); };
	void Signal_LinkClicked(const QUrl & url) { callbackQWebView_LinkClicked(this, const_cast<QUrl*>(&url)); };
	void Signal_LoadFinished(bool ok) { callbackQWebView_LoadFinished(this, ok); };
	void Signal_LoadProgress(int progress) { callbackQWebView_LoadProgress(this, progress); };
	void Signal_LoadStarted() { callbackQWebView_LoadStarted(this); };
	void mouseDoubleClickEvent(QMouseEvent * ev) { callbackQWebView_MouseDoubleClickEvent(this, ev); };
	void mouseMoveEvent(QMouseEvent * ev) { callbackQWebView_MouseMoveEvent(this, ev); };
	void mousePressEvent(QMouseEvent * ev) { callbackQWebView_MousePressEvent(this, ev); };
	void mouseReleaseEvent(QMouseEvent * ev) { callbackQWebView_MouseReleaseEvent(this, ev); };
	void paintEvent(QPaintEvent * ev) { callbackQWebView_PaintEvent(this, ev); };
	void reload() { callbackQWebView_Reload(this); };
	void resizeEvent(QResizeEvent * e) { callbackQWebView_ResizeEvent(this, e); };
	void Signal_SelectionChanged() { callbackQWebView_SelectionChanged(this); };
	void Signal_StatusBarMessage(const QString & text) { QByteArray t372ea0 = text.toUtf8(); QtWebKit_PackedString textPacked = { const_cast<char*>(t372ea0.prepend("WHITESPACE").constData()+10), t372ea0.size()-10 };callbackQWebView_StatusBarMessage(this, textPacked); };
	void stop() { callbackQWebView_Stop(this); };
	void Signal_TitleChanged(const QString & title) { QByteArray t3c6de1 = title.toUtf8(); QtWebKit_PackedString titlePacked = { const_cast<char*>(t3c6de1.prepend("WHITESPACE").constData()+10), t3c6de1.size()-10 };callbackQWebView_TitleChanged(this, titlePacked); };
	void Signal_UrlChanged(const QUrl & url) { callbackQWebView_UrlChanged(this, const_cast<QUrl*>(&url)); };
	void wheelEvent(QWheelEvent * ev) { callbackQWebView_WheelEvent(this, ev); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQWebView_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	QVariant inputMethodQuery(Qt::InputMethodQuery property) const { return *static_cast<QVariant*>(callbackQWebView_InputMethodQuery(const_cast<void*>(static_cast<const void*>(this)), property)); };
	void print(QPrinter * printer) const { callbackQWebView_Print(const_cast<void*>(static_cast<const void*>(this)), printer); };
	bool close() { return callbackQWebView_Close(this) != 0; };
	bool nativeEvent(const QByteArray & eventType, void * message, long * result) { return callbackQWebView_NativeEvent(this, const_cast<QByteArray*>(&eventType), message, *result) != 0; };
	void actionEvent(QActionEvent * event) { callbackQWebView_ActionEvent(this, event); };
	void closeEvent(QCloseEvent * event) { callbackQWebView_CloseEvent(this, event); };
	void Signal_CustomContextMenuRequested(const QPoint & pos) { callbackQWebView_CustomContextMenuRequested(this, const_cast<QPoint*>(&pos)); };
	void enterEvent(QEvent * event) { callbackQWebView_EnterEvent(this, event); };
	void hide() { callbackQWebView_Hide(this); };
	void hideEvent(QHideEvent * event) { callbackQWebView_HideEvent(this, event); };
	void leaveEvent(QEvent * event) { callbackQWebView_LeaveEvent(this, event); };
	void lower() { callbackQWebView_Lower(this); };
	void moveEvent(QMoveEvent * event) { callbackQWebView_MoveEvent(this, event); };
	void raise() { callbackQWebView_Raise(this); };
	void repaint() { callbackQWebView_Repaint(this); };
	void setDisabled(bool disable) { callbackQWebView_SetDisabled(this, disable); };
	void setEnabled(bool vbo) { callbackQWebView_SetEnabled(this, vbo); };
	void setFocus() { callbackQWebView_SetFocus2(this); };
	void setHidden(bool hidden) { callbackQWebView_SetHidden(this, hidden); };
	void setStyleSheet(const QString & styleSheet) { QByteArray t728ae7 = styleSheet.toUtf8(); QtWebKit_PackedString styleSheetPacked = { const_cast<char*>(t728ae7.prepend("WHITESPACE").constData()+10), t728ae7.size()-10 };callbackQWebView_SetStyleSheet(this, styleSheetPacked); };
	void setVisible(bool visible) { callbackQWebView_SetVisible(this, visible); };
	void setWindowModified(bool vbo) { callbackQWebView_SetWindowModified(this, vbo); };
	void setWindowTitle(const QString & vqs) { QByteArray tda39a3 = vqs.toUtf8(); QtWebKit_PackedString vqsPacked = { const_cast<char*>(tda39a3.prepend("WHITESPACE").constData()+10), tda39a3.size()-10 };callbackQWebView_SetWindowTitle(this, vqsPacked); };
	void show() { callbackQWebView_Show(this); };
	void showEvent(QShowEvent * event) { callbackQWebView_ShowEvent(this, event); };
	void showFullScreen() { callbackQWebView_ShowFullScreen(this); };
	void showMaximized() { callbackQWebView_ShowMaximized(this); };
	void showMinimized() { callbackQWebView_ShowMinimized(this); };
	void showNormal() { callbackQWebView_ShowNormal(this); };
	void tabletEvent(QTabletEvent * event) { callbackQWebView_TabletEvent(this, event); };
	void update() { callbackQWebView_Update(this); };
	void updateMicroFocus() { callbackQWebView_UpdateMicroFocus(this); };
	void Signal_WindowIconChanged(const QIcon & icon) { callbackQWebView_WindowIconChanged(this, const_cast<QIcon*>(&icon)); };
	void Signal_WindowTitleChanged(const QString & title) { QByteArray t3c6de1 = title.toUtf8(); QtWebKit_PackedString titlePacked = { const_cast<char*>(t3c6de1.prepend("WHITESPACE").constData()+10), t3c6de1.size()-10 };callbackQWebView_WindowTitleChanged(this, titlePacked); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQWebView_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize minimumSizeHint() const { return *static_cast<QSize*>(callbackQWebView_MinimumSizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	bool hasHeightForWidth() const { return callbackQWebView_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int heightForWidth(int w) const { return callbackQWebView_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackQWebView_Metric(const_cast<void*>(static_cast<const void*>(this)), m); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQWebView_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQWebView_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQWebView_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQWebView_CustomEvent(this, event); };
	void deleteLater() { callbackQWebView_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQWebView_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQWebView_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtWebKit_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQWebView_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQWebView_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQWebView_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQWebView*)

int QWebView_QWebView_QRegisterMetaType(){qRegisterMetaType<QWebView*>(); return qRegisterMetaType<MyQWebView*>();}

void* QWebView_CreateWindow(void* ptr, long long ty)
{
	return static_cast<QWebView*>(ptr)->createWindow(static_cast<QWebPage::WebWindowType>(ty));
}

void* QWebView_CreateWindowDefault(void* ptr, long long ty)
{
		return static_cast<QWebView*>(ptr)->QWebView::createWindow(static_cast<QWebPage::WebWindowType>(ty));
}

void* QWebView_NewQWebView(void* parent)
{
		return new MyQWebView(static_cast<QWidget*>(parent));
}

char QWebView_EventDefault(void* ptr, void* e)
{
		return static_cast<QWebView*>(ptr)->QWebView::event(static_cast<QEvent*>(e));
}

char QWebView_FindText(void* ptr, struct QtWebKit_PackedString subString, long long options)
{
	return static_cast<QWebView*>(ptr)->findText(QString::fromUtf8(subString.data, subString.len), static_cast<QWebPage::FindFlag>(options));
}

char QWebView_FocusNextPrevChildDefault(void* ptr, char next)
{
		return static_cast<QWebView*>(ptr)->QWebView::focusNextPrevChild(next != 0);
}

void QWebView_Back(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QWebView*>(ptr), "back");
}

void QWebView_BackDefault(void* ptr)
{
		static_cast<QWebView*>(ptr)->QWebView::back();
}

void QWebView_ChangeEventDefault(void* ptr, void* e)
{
		static_cast<QWebView*>(ptr)->QWebView::changeEvent(static_cast<QEvent*>(e));
}

void QWebView_ContextMenuEventDefault(void* ptr, void* ev)
{
		static_cast<QWebView*>(ptr)->QWebView::contextMenuEvent(static_cast<QContextMenuEvent*>(ev));
}

void QWebView_DragEnterEventDefault(void* ptr, void* ev)
{
		static_cast<QWebView*>(ptr)->QWebView::dragEnterEvent(static_cast<QDragEnterEvent*>(ev));
}

void QWebView_DragLeaveEventDefault(void* ptr, void* ev)
{
		static_cast<QWebView*>(ptr)->QWebView::dragLeaveEvent(static_cast<QDragLeaveEvent*>(ev));
}

void QWebView_DragMoveEventDefault(void* ptr, void* ev)
{
		static_cast<QWebView*>(ptr)->QWebView::dragMoveEvent(static_cast<QDragMoveEvent*>(ev));
}

void QWebView_DropEventDefault(void* ptr, void* ev)
{
		static_cast<QWebView*>(ptr)->QWebView::dropEvent(static_cast<QDropEvent*>(ev));
}

void QWebView_FocusInEventDefault(void* ptr, void* ev)
{
		static_cast<QWebView*>(ptr)->QWebView::focusInEvent(static_cast<QFocusEvent*>(ev));
}

void QWebView_FocusOutEventDefault(void* ptr, void* ev)
{
		static_cast<QWebView*>(ptr)->QWebView::focusOutEvent(static_cast<QFocusEvent*>(ev));
}

void QWebView_Forward(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QWebView*>(ptr), "forward");
}

void QWebView_ForwardDefault(void* ptr)
{
		static_cast<QWebView*>(ptr)->QWebView::forward();
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

void QWebView_InputMethodEventDefault(void* ptr, void* e)
{
		static_cast<QWebView*>(ptr)->QWebView::inputMethodEvent(static_cast<QInputMethodEvent*>(e));
}

void QWebView_KeyPressEventDefault(void* ptr, void* ev)
{
		static_cast<QWebView*>(ptr)->QWebView::keyPressEvent(static_cast<QKeyEvent*>(ev));
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

void QWebView_Load2(void* ptr, void* request, long long operation, void* body)
{
	static_cast<QWebView*>(ptr)->load(*static_cast<QNetworkRequest*>(request), static_cast<QNetworkAccessManager::Operation>(operation), *static_cast<QByteArray*>(body));
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

void QWebView_LoadFinished(void* ptr, char ok)
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

void QWebView_MouseDoubleClickEventDefault(void* ptr, void* ev)
{
		static_cast<QWebView*>(ptr)->QWebView::mouseDoubleClickEvent(static_cast<QMouseEvent*>(ev));
}

void QWebView_MouseMoveEventDefault(void* ptr, void* ev)
{
		static_cast<QWebView*>(ptr)->QWebView::mouseMoveEvent(static_cast<QMouseEvent*>(ev));
}

void QWebView_MousePressEventDefault(void* ptr, void* ev)
{
		static_cast<QWebView*>(ptr)->QWebView::mousePressEvent(static_cast<QMouseEvent*>(ev));
}

void QWebView_MouseReleaseEventDefault(void* ptr, void* ev)
{
		static_cast<QWebView*>(ptr)->QWebView::mouseReleaseEvent(static_cast<QMouseEvent*>(ev));
}

void QWebView_PaintEventDefault(void* ptr, void* ev)
{
		static_cast<QWebView*>(ptr)->QWebView::paintEvent(static_cast<QPaintEvent*>(ev));
}

void QWebView_Reload(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QWebView*>(ptr), "reload");
}

void QWebView_ReloadDefault(void* ptr)
{
		static_cast<QWebView*>(ptr)->QWebView::reload();
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

void QWebView_SetContent(void* ptr, void* data, struct QtWebKit_PackedString mimeType, void* baseUrl)
{
	static_cast<QWebView*>(ptr)->setContent(*static_cast<QByteArray*>(data), QString::fromUtf8(mimeType.data, mimeType.len), *static_cast<QUrl*>(baseUrl));
}

void QWebView_SetHtml(void* ptr, struct QtWebKit_PackedString html, void* baseUrl)
{
	static_cast<QWebView*>(ptr)->setHtml(QString::fromUtf8(html.data, html.len), *static_cast<QUrl*>(baseUrl));
}

void QWebView_SetPage(void* ptr, void* page)
{
	static_cast<QWebView*>(ptr)->setPage(static_cast<QWebPage*>(page));
}

void QWebView_SetRenderHint(void* ptr, long long hint, char enabled)
{
	static_cast<QWebView*>(ptr)->setRenderHint(static_cast<QPainter::RenderHint>(hint), enabled != 0);
}

void QWebView_SetRenderHints(void* ptr, long long hints)
{
	static_cast<QWebView*>(ptr)->setRenderHints(static_cast<QPainter::RenderHint>(hints));
}

void QWebView_SetTextSizeMultiplier(void* ptr, double factor)
{
	static_cast<QWebView*>(ptr)->setTextSizeMultiplier(factor);
}

void QWebView_SetUrl(void* ptr, void* url)
{
	static_cast<QWebView*>(ptr)->setUrl(*static_cast<QUrl*>(url));
}

void QWebView_SetZoomFactor(void* ptr, double factor)
{
	static_cast<QWebView*>(ptr)->setZoomFactor(factor);
}

void QWebView_ConnectStatusBarMessage(void* ptr)
{
	QObject::connect(static_cast<QWebView*>(ptr), static_cast<void (QWebView::*)(const QString &)>(&QWebView::statusBarMessage), static_cast<MyQWebView*>(ptr), static_cast<void (MyQWebView::*)(const QString &)>(&MyQWebView::Signal_StatusBarMessage));
}

void QWebView_DisconnectStatusBarMessage(void* ptr)
{
	QObject::disconnect(static_cast<QWebView*>(ptr), static_cast<void (QWebView::*)(const QString &)>(&QWebView::statusBarMessage), static_cast<MyQWebView*>(ptr), static_cast<void (MyQWebView::*)(const QString &)>(&MyQWebView::Signal_StatusBarMessage));
}

void QWebView_StatusBarMessage(void* ptr, struct QtWebKit_PackedString text)
{
	static_cast<QWebView*>(ptr)->statusBarMessage(QString::fromUtf8(text.data, text.len));
}

void QWebView_Stop(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QWebView*>(ptr), "stop");
}

void QWebView_StopDefault(void* ptr)
{
		static_cast<QWebView*>(ptr)->QWebView::stop();
}

void QWebView_ConnectTitleChanged(void* ptr)
{
	QObject::connect(static_cast<QWebView*>(ptr), static_cast<void (QWebView::*)(const QString &)>(&QWebView::titleChanged), static_cast<MyQWebView*>(ptr), static_cast<void (MyQWebView::*)(const QString &)>(&MyQWebView::Signal_TitleChanged));
}

void QWebView_DisconnectTitleChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWebView*>(ptr), static_cast<void (QWebView::*)(const QString &)>(&QWebView::titleChanged), static_cast<MyQWebView*>(ptr), static_cast<void (MyQWebView::*)(const QString &)>(&MyQWebView::Signal_TitleChanged));
}

void QWebView_TitleChanged(void* ptr, struct QtWebKit_PackedString title)
{
	static_cast<QWebView*>(ptr)->titleChanged(QString::fromUtf8(title.data, title.len));
}

void QWebView_TriggerPageAction(void* ptr, long long action, char checked)
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

void QWebView_WheelEventDefault(void* ptr, void* ev)
{
		static_cast<QWebView*>(ptr)->QWebView::wheelEvent(static_cast<QWheelEvent*>(ev));
}

void QWebView_DestroyQWebView(void* ptr)
{
	static_cast<QWebView*>(ptr)->~QWebView();
}

void* QWebView_PageAction(void* ptr, long long action)
{
	return static_cast<QWebView*>(ptr)->pageAction(static_cast<QWebPage::WebAction>(action));
}

void* QWebView_Icon(void* ptr)
{
	return new QIcon(static_cast<QWebView*>(ptr)->icon());
}

long long QWebView_RenderHints(void* ptr)
{
	return static_cast<QWebView*>(ptr)->renderHints();
}

void* QWebView_SizeHintDefault(void* ptr)
{
		return ({ QSize tmpValue = static_cast<QWebView*>(ptr)->QWebView::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

struct QtWebKit_PackedString QWebView_SelectedHtml(void* ptr)
{
	return ({ QByteArray tff0b31 = static_cast<QWebView*>(ptr)->selectedHtml().toUtf8(); QtWebKit_PackedString { const_cast<char*>(tff0b31.prepend("WHITESPACE").constData()+10), tff0b31.size()-10 }; });
}

struct QtWebKit_PackedString QWebView_SelectedText(void* ptr)
{
	return ({ QByteArray t8fa677 = static_cast<QWebView*>(ptr)->selectedText().toUtf8(); QtWebKit_PackedString { const_cast<char*>(t8fa677.prepend("WHITESPACE").constData()+10), t8fa677.size()-10 }; });
}

struct QtWebKit_PackedString QWebView_Title(void* ptr)
{
	return ({ QByteArray te51897 = static_cast<QWebView*>(ptr)->title().toUtf8(); QtWebKit_PackedString { const_cast<char*>(te51897.prepend("WHITESPACE").constData()+10), te51897.size()-10 }; });
}

void* QWebView_Url(void* ptr)
{
	return new QUrl(static_cast<QWebView*>(ptr)->url());
}

void* QWebView_InputMethodQueryDefault(void* ptr, long long property)
{
		return new QVariant(static_cast<QWebView*>(ptr)->QWebView::inputMethodQuery(static_cast<Qt::InputMethodQuery>(property)));
}

void* QWebView_History(void* ptr)
{
	return static_cast<QWebView*>(ptr)->history();
}

void* QWebView_Page(void* ptr)
{
	return static_cast<QWebView*>(ptr)->page();
}

void* QWebView_Settings(void* ptr)
{
	return static_cast<QWebView*>(ptr)->settings();
}

char QWebView_HasSelection(void* ptr)
{
	return static_cast<QWebView*>(ptr)->hasSelection();
}

char QWebView_IsModified(void* ptr)
{
	return static_cast<QWebView*>(ptr)->isModified();
}

double QWebView_TextSizeMultiplier(void* ptr)
{
	return static_cast<QWebView*>(ptr)->textSizeMultiplier();
}

double QWebView_ZoomFactor(void* ptr)
{
	return static_cast<QWebView*>(ptr)->zoomFactor();
}

void QWebView_Print(void* ptr, void* printer)
{
	QMetaObject::invokeMethod(static_cast<QWebView*>(ptr), "print", Q_ARG(QPrinter*, static_cast<QPrinter*>(printer)));
}

void QWebView_PrintDefault(void* ptr, void* printer)
{
		static_cast<QWebView*>(ptr)->QWebView::print(static_cast<QPrinter*>(printer));
}

void* QWebView___addActions_actions_atList(void* ptr, int i)
{
	return const_cast<QAction*>(static_cast<QList<QAction *>*>(ptr)->at(i));
}

void QWebView___addActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QWebView___addActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>;
}

void* QWebView___insertActions_actions_atList(void* ptr, int i)
{
	return const_cast<QAction*>(static_cast<QList<QAction *>*>(ptr)->at(i));
}

void QWebView___insertActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QWebView___insertActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>;
}

void* QWebView___actions_atList(void* ptr, int i)
{
	return const_cast<QAction*>(static_cast<QList<QAction *>*>(ptr)->at(i));
}

void QWebView___actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QWebView___actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>;
}

void* QWebView___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QWebView___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QWebView___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QWebView___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QWebView___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebView___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QWebView___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QWebView___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebView___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QWebView___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QWebView___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebView___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QWebView___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QWebView___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWebView___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

char QWebView_CloseDefault(void* ptr)
{
		return static_cast<QWebView*>(ptr)->QWebView::close();
}

char QWebView_NativeEventDefault(void* ptr, void* eventType, void* message, long result)
{
		return static_cast<QWebView*>(ptr)->QWebView::nativeEvent(*static_cast<QByteArray*>(eventType), message, &result);
}

void QWebView_ActionEventDefault(void* ptr, void* event)
{
		static_cast<QWebView*>(ptr)->QWebView::actionEvent(static_cast<QActionEvent*>(event));
}

void QWebView_CloseEventDefault(void* ptr, void* event)
{
		static_cast<QWebView*>(ptr)->QWebView::closeEvent(static_cast<QCloseEvent*>(event));
}

void QWebView_EnterEventDefault(void* ptr, void* event)
{
		static_cast<QWebView*>(ptr)->QWebView::enterEvent(static_cast<QEvent*>(event));
}

void QWebView_HideDefault(void* ptr)
{
		static_cast<QWebView*>(ptr)->QWebView::hide();
}

void QWebView_HideEventDefault(void* ptr, void* event)
{
		static_cast<QWebView*>(ptr)->QWebView::hideEvent(static_cast<QHideEvent*>(event));
}

void QWebView_LeaveEventDefault(void* ptr, void* event)
{
		static_cast<QWebView*>(ptr)->QWebView::leaveEvent(static_cast<QEvent*>(event));
}

void QWebView_LowerDefault(void* ptr)
{
		static_cast<QWebView*>(ptr)->QWebView::lower();
}

void QWebView_MoveEventDefault(void* ptr, void* event)
{
		static_cast<QWebView*>(ptr)->QWebView::moveEvent(static_cast<QMoveEvent*>(event));
}

void QWebView_RaiseDefault(void* ptr)
{
		static_cast<QWebView*>(ptr)->QWebView::raise();
}

void QWebView_RepaintDefault(void* ptr)
{
		static_cast<QWebView*>(ptr)->QWebView::repaint();
}

void QWebView_SetDisabledDefault(void* ptr, char disable)
{
		static_cast<QWebView*>(ptr)->QWebView::setDisabled(disable != 0);
}

void QWebView_SetEnabledDefault(void* ptr, char vbo)
{
		static_cast<QWebView*>(ptr)->QWebView::setEnabled(vbo != 0);
}

void QWebView_SetFocus2Default(void* ptr)
{
		static_cast<QWebView*>(ptr)->QWebView::setFocus();
}

void QWebView_SetHiddenDefault(void* ptr, char hidden)
{
		static_cast<QWebView*>(ptr)->QWebView::setHidden(hidden != 0);
}

void QWebView_SetStyleSheetDefault(void* ptr, struct QtWebKit_PackedString styleSheet)
{
		static_cast<QWebView*>(ptr)->QWebView::setStyleSheet(QString::fromUtf8(styleSheet.data, styleSheet.len));
}

void QWebView_SetVisibleDefault(void* ptr, char visible)
{
		static_cast<QWebView*>(ptr)->QWebView::setVisible(visible != 0);
}

void QWebView_SetWindowModifiedDefault(void* ptr, char vbo)
{
		static_cast<QWebView*>(ptr)->QWebView::setWindowModified(vbo != 0);
}

void QWebView_SetWindowTitleDefault(void* ptr, struct QtWebKit_PackedString vqs)
{
		static_cast<QWebView*>(ptr)->QWebView::setWindowTitle(QString::fromUtf8(vqs.data, vqs.len));
}

void QWebView_ShowDefault(void* ptr)
{
		static_cast<QWebView*>(ptr)->QWebView::show();
}

void QWebView_ShowEventDefault(void* ptr, void* event)
{
		static_cast<QWebView*>(ptr)->QWebView::showEvent(static_cast<QShowEvent*>(event));
}

void QWebView_ShowFullScreenDefault(void* ptr)
{
		static_cast<QWebView*>(ptr)->QWebView::showFullScreen();
}

void QWebView_ShowMaximizedDefault(void* ptr)
{
		static_cast<QWebView*>(ptr)->QWebView::showMaximized();
}

void QWebView_ShowMinimizedDefault(void* ptr)
{
		static_cast<QWebView*>(ptr)->QWebView::showMinimized();
}

void QWebView_ShowNormalDefault(void* ptr)
{
		static_cast<QWebView*>(ptr)->QWebView::showNormal();
}

void QWebView_TabletEventDefault(void* ptr, void* event)
{
		static_cast<QWebView*>(ptr)->QWebView::tabletEvent(static_cast<QTabletEvent*>(event));
}

void QWebView_UpdateDefault(void* ptr)
{
		static_cast<QWebView*>(ptr)->QWebView::update();
}

void QWebView_UpdateMicroFocusDefault(void* ptr)
{
		static_cast<QWebView*>(ptr)->QWebView::updateMicroFocus();
}

void* QWebView_PaintEngineDefault(void* ptr)
{
		return static_cast<QWebView*>(ptr)->QWebView::paintEngine();
}

void* QWebView_MinimumSizeHintDefault(void* ptr)
{
		return ({ QSize tmpValue = static_cast<QWebView*>(ptr)->QWebView::minimumSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

char QWebView_HasHeightForWidthDefault(void* ptr)
{
		return static_cast<QWebView*>(ptr)->QWebView::hasHeightForWidth();
}

int QWebView_HeightForWidthDefault(void* ptr, int w)
{
		return static_cast<QWebView*>(ptr)->QWebView::heightForWidth(w);
}

int QWebView_MetricDefault(void* ptr, long long m)
{
		return static_cast<QWebView*>(ptr)->QWebView::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
}

char QWebView_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QWebView*>(ptr)->QWebView::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QWebView_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QWebView*>(ptr)->QWebView::childEvent(static_cast<QChildEvent*>(event));
}

void QWebView_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QWebView*>(ptr)->QWebView::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWebView_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QWebView*>(ptr)->QWebView::customEvent(static_cast<QEvent*>(event));
}

void QWebView_DeleteLaterDefault(void* ptr)
{
		static_cast<QWebView*>(ptr)->QWebView::deleteLater();
}

void QWebView_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QWebView*>(ptr)->QWebView::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWebView_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QWebView*>(ptr)->QWebView::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QWebView_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QWebView*>(ptr)->QWebView::metaObject());
}

