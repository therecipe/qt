#define protected public
#define private public

#include "quick.h"
#include "_cgo_export.h"

#include <QAction>
#include <QActionEvent>
#include <QByteArray>
#include <QChildEvent>
#include <QCloseEvent>
#include <QColor>
#include <QContextMenuEvent>
#include <QCursor>
#include <QDrag>
#include <QDragEnterEvent>
#include <QDragLeaveEvent>
#include <QDragMoveEvent>
#include <QDropEvent>
#include <QEvent>
#include <QExposeEvent>
#include <QFocusEvent>
#include <QHideEvent>
#include <QHoverEvent>
#include <QImage>
#include <QInputMethod>
#include <QInputMethodEvent>
#include <QKeyEvent>
#include <QMatrix4x4>
#include <QMetaMethod>
#include <QMetaObject>
#include <QMouseEvent>
#include <QMoveEvent>
#include <QObject>
#include <QOpenGLContext>
#include <QOpenGLFramebufferObject>
#include <QOpenGLShader>
#include <QPaintEvent>
#include <QPainter>
#include <QPoint>
#include <QPointF>
#include <QQmlEngine>
#include <QQmlImageProviderBase>
#include <QQuickAsyncImageProvider>
#include <QQuickFramebufferObject>
#include <QQuickImageProvider>
#include <QQuickImageResponse>
#include <QQuickItem>
#include <QQuickItemGrabResult>
#include <QQuickPaintedItem>
#include <QQuickRenderControl>
#include <QQuickTextDocument>
#include <QQuickTextureFactory>
#include <QQuickView>
#include <QQuickWidget>
#include <QQuickWindow>
#include <QRect>
#include <QRectF>
#include <QResizeEvent>
#include <QRunnable>
#include <QSGAbstractRenderer>
#include <QSGBasicGeometryNode>
#include <QSGClipNode>
#include <QSGDynamicTexture>
#include <QSGEngine>
#include <QSGFlatColorMaterial>
#include <QSGGeometry>
#include <QSGGeometryNode>
#include <QSGMaterial>
#include <QSGMaterialShader>
#include <QSGMaterialType>
#include <QSGNode>
#include <QSGOpacityNode>
#include <QSGOpaqueTextureMaterial>
#include <QSGSimpleRectNode>
#include <QSGSimpleTextureNode>
#include <QSGTexture>
#include <QSGTextureMaterial>
#include <QSGTextureProvider>
#include <QSGTransformNode>
#include <QSGVertexColorMaterial>
#include <QShowEvent>
#include <QSize>
#include <QString>
#include <QSurface>
#include <QSurfaceFormat>
#include <QTabletEvent>
#include <QThread>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>
#include <QTouchEvent>
#include <QUrl>
#include <QVariant>
#include <QWheelEvent>
#include <QWidget>
#include <QWindow>

class MyQQuickAsyncImageProvider: public QQuickAsyncImageProvider
{
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	MyQQuickAsyncImageProvider() : QQuickAsyncImageProvider() {};
	QQuickImageResponse * requestImageResponse(const QString & id, const QSize & requestedSize) { return static_cast<QQuickImageResponse*>(callbackQQuickAsyncImageProvider_RequestImageResponse(this, this->objectNameAbs().toUtf8().data(), id.toUtf8().data(), new QSize(static_cast<QSize>(requestedSize).width(), static_cast<QSize>(requestedSize).height()))); };
};

void* QQuickAsyncImageProvider_NewQQuickAsyncImageProvider()
{
	return new MyQQuickAsyncImageProvider();
}

void QQuickAsyncImageProvider_DestroyQQuickAsyncImageProvider(void* ptr)
{
	static_cast<QQuickAsyncImageProvider*>(ptr)->~QQuickAsyncImageProvider();
}

void* QQuickAsyncImageProvider_RequestImageResponse(void* ptr, char* id, void* requestedSize)
{
	return static_cast<QQuickAsyncImageProvider*>(ptr)->requestImageResponse(QString(id), *static_cast<QSize*>(requestedSize));
}

char* QQuickAsyncImageProvider_ObjectNameAbs(void* ptr)
{
	if (dynamic_cast<MyQQuickAsyncImageProvider*>(static_cast<QQuickAsyncImageProvider*>(ptr))) {
		return static_cast<MyQQuickAsyncImageProvider*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QQuickAsyncImageProvider_BASE").toUtf8().data();
}

void QQuickAsyncImageProvider_SetObjectNameAbs(void* ptr, char* name)
{
	if (dynamic_cast<MyQQuickAsyncImageProvider*>(static_cast<QQuickAsyncImageProvider*>(ptr))) {
		static_cast<MyQQuickAsyncImageProvider*>(ptr)->setObjectNameAbs(QString(name));
	}
}

class MyQQuickFramebufferObject: public QQuickFramebufferObject
{
public:
	bool isTextureProvider() const { return callbackQQuickFramebufferObject_IsTextureProvider(const_cast<MyQQuickFramebufferObject*>(this), this->objectName().toUtf8().data()) != 0; };
	void Signal_MirrorVerticallyChanged(bool vbo) { callbackQQuickFramebufferObject_MirrorVerticallyChanged(this, this->objectName().toUtf8().data(), vbo); };
	void releaseResources() { callbackQQuickFramebufferObject_ReleaseResources(this, this->objectName().toUtf8().data()); };
	void Signal_TextureFollowsItemSizeChanged(bool vbo) { callbackQQuickFramebufferObject_TextureFollowsItemSizeChanged(this, this->objectName().toUtf8().data(), vbo); };
	QSGTextureProvider * textureProvider() const { return static_cast<QSGTextureProvider*>(callbackQQuickFramebufferObject_TextureProvider(const_cast<MyQQuickFramebufferObject*>(this), this->objectName().toUtf8().data())); };
	bool childMouseEventFilter(QQuickItem * item, QEvent * event) { return callbackQQuickFramebufferObject_ChildMouseEventFilter(this, this->objectName().toUtf8().data(), item, event) != 0; };
	bool contains(const QPointF & point) const { return callbackQQuickFramebufferObject_Contains(const_cast<MyQQuickFramebufferObject*>(this), this->objectName().toUtf8().data(), new QPointF(static_cast<QPointF>(point).x(), static_cast<QPointF>(point).y())) != 0; };
	void dragEnterEvent(QDragEnterEvent * event) { callbackQQuickFramebufferObject_DragEnterEvent(this, this->objectName().toUtf8().data(), event); };
	void dragLeaveEvent(QDragLeaveEvent * event) { callbackQQuickFramebufferObject_DragLeaveEvent(this, this->objectName().toUtf8().data(), event); };
	void dragMoveEvent(QDragMoveEvent * event) { callbackQQuickFramebufferObject_DragMoveEvent(this, this->objectName().toUtf8().data(), event); };
	void dropEvent(QDropEvent * event) { callbackQQuickFramebufferObject_DropEvent(this, this->objectName().toUtf8().data(), event); };
	void focusInEvent(QFocusEvent * event) { callbackQQuickFramebufferObject_FocusInEvent(this, this->objectName().toUtf8().data(), event); };
	void focusOutEvent(QFocusEvent * event) { callbackQQuickFramebufferObject_FocusOutEvent(this, this->objectName().toUtf8().data(), event); };
	void geometryChanged(const QRectF & newGeometry, const QRectF & oldGeometry) { callbackQQuickFramebufferObject_GeometryChanged(this, this->objectName().toUtf8().data(), new QRectF(static_cast<QRectF>(newGeometry).x(), static_cast<QRectF>(newGeometry).y(), static_cast<QRectF>(newGeometry).width(), static_cast<QRectF>(newGeometry).height()), new QRectF(static_cast<QRectF>(oldGeometry).x(), static_cast<QRectF>(oldGeometry).y(), static_cast<QRectF>(oldGeometry).width(), static_cast<QRectF>(oldGeometry).height())); };
	void hoverEnterEvent(QHoverEvent * event) { callbackQQuickFramebufferObject_HoverEnterEvent(this, this->objectName().toUtf8().data(), event); };
	void hoverLeaveEvent(QHoverEvent * event) { callbackQQuickFramebufferObject_HoverLeaveEvent(this, this->objectName().toUtf8().data(), event); };
	void hoverMoveEvent(QHoverEvent * event) { callbackQQuickFramebufferObject_HoverMoveEvent(this, this->objectName().toUtf8().data(), event); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackQQuickFramebufferObject_InputMethodEvent(this, this->objectName().toUtf8().data(), event); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackQQuickFramebufferObject_InputMethodQuery(const_cast<MyQQuickFramebufferObject*>(this), this->objectName().toUtf8().data(), query)); };
	void keyPressEvent(QKeyEvent * event) { callbackQQuickFramebufferObject_KeyPressEvent(this, this->objectName().toUtf8().data(), event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQQuickFramebufferObject_KeyReleaseEvent(this, this->objectName().toUtf8().data(), event); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQQuickFramebufferObject_MouseDoubleClickEvent(this, this->objectName().toUtf8().data(), event); };
	void mouseMoveEvent(QMouseEvent * event) { callbackQQuickFramebufferObject_MouseMoveEvent(this, this->objectName().toUtf8().data(), event); };
	void mousePressEvent(QMouseEvent * event) { callbackQQuickFramebufferObject_MousePressEvent(this, this->objectName().toUtf8().data(), event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackQQuickFramebufferObject_MouseReleaseEvent(this, this->objectName().toUtf8().data(), event); };
	void mouseUngrabEvent() { callbackQQuickFramebufferObject_MouseUngrabEvent(this, this->objectName().toUtf8().data()); };
	void touchEvent(QTouchEvent * event) { callbackQQuickFramebufferObject_TouchEvent(this, this->objectName().toUtf8().data(), event); };
	void touchUngrabEvent() { callbackQQuickFramebufferObject_TouchUngrabEvent(this, this->objectName().toUtf8().data()); };
	void update() { callbackQQuickFramebufferObject_Update(this, this->objectName().toUtf8().data()); };
	void updatePolish() { callbackQQuickFramebufferObject_UpdatePolish(this, this->objectName().toUtf8().data()); };
	void wheelEvent(QWheelEvent * event) { callbackQQuickFramebufferObject_WheelEvent(this, this->objectName().toUtf8().data(), event); };
	void timerEvent(QTimerEvent * event) { callbackQQuickFramebufferObject_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQQuickFramebufferObject_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQQuickFramebufferObject_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQQuickFramebufferObject_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQQuickFramebufferObject_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQQuickFramebufferObject_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQQuickFramebufferObject_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQQuickFramebufferObject_MetaObject(const_cast<MyQQuickFramebufferObject*>(this), this->objectName().toUtf8().data())); };
};

int QQuickFramebufferObject_MirrorVertically(void* ptr)
{
	return static_cast<QQuickFramebufferObject*>(ptr)->mirrorVertically();
}

void QQuickFramebufferObject_SetMirrorVertically(void* ptr, int enable)
{
	static_cast<QQuickFramebufferObject*>(ptr)->setMirrorVertically(enable != 0);
}

void QQuickFramebufferObject_SetTextureFollowsItemSize(void* ptr, int follows)
{
	static_cast<QQuickFramebufferObject*>(ptr)->setTextureFollowsItemSize(follows != 0);
}

int QQuickFramebufferObject_TextureFollowsItemSize(void* ptr)
{
	return static_cast<QQuickFramebufferObject*>(ptr)->textureFollowsItemSize();
}

int QQuickFramebufferObject_IsTextureProvider(void* ptr)
{
	return static_cast<QQuickFramebufferObject*>(ptr)->isTextureProvider();
}

int QQuickFramebufferObject_IsTextureProviderDefault(void* ptr)
{
	return static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::isTextureProvider();
}

void QQuickFramebufferObject_ConnectMirrorVerticallyChanged(void* ptr)
{
	QObject::connect(static_cast<QQuickFramebufferObject*>(ptr), static_cast<void (QQuickFramebufferObject::*)(bool)>(&QQuickFramebufferObject::mirrorVerticallyChanged), static_cast<MyQQuickFramebufferObject*>(ptr), static_cast<void (MyQQuickFramebufferObject::*)(bool)>(&MyQQuickFramebufferObject::Signal_MirrorVerticallyChanged));
}

void QQuickFramebufferObject_DisconnectMirrorVerticallyChanged(void* ptr)
{
	QObject::disconnect(static_cast<QQuickFramebufferObject*>(ptr), static_cast<void (QQuickFramebufferObject::*)(bool)>(&QQuickFramebufferObject::mirrorVerticallyChanged), static_cast<MyQQuickFramebufferObject*>(ptr), static_cast<void (MyQQuickFramebufferObject::*)(bool)>(&MyQQuickFramebufferObject::Signal_MirrorVerticallyChanged));
}

void QQuickFramebufferObject_MirrorVerticallyChanged(void* ptr, int vbo)
{
	static_cast<QQuickFramebufferObject*>(ptr)->mirrorVerticallyChanged(vbo != 0);
}

void QQuickFramebufferObject_ReleaseResources(void* ptr)
{
	static_cast<QQuickFramebufferObject*>(ptr)->releaseResources();
}

void QQuickFramebufferObject_ReleaseResourcesDefault(void* ptr)
{
	static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::releaseResources();
}

void QQuickFramebufferObject_ConnectTextureFollowsItemSizeChanged(void* ptr)
{
	QObject::connect(static_cast<QQuickFramebufferObject*>(ptr), static_cast<void (QQuickFramebufferObject::*)(bool)>(&QQuickFramebufferObject::textureFollowsItemSizeChanged), static_cast<MyQQuickFramebufferObject*>(ptr), static_cast<void (MyQQuickFramebufferObject::*)(bool)>(&MyQQuickFramebufferObject::Signal_TextureFollowsItemSizeChanged));
}

void QQuickFramebufferObject_DisconnectTextureFollowsItemSizeChanged(void* ptr)
{
	QObject::disconnect(static_cast<QQuickFramebufferObject*>(ptr), static_cast<void (QQuickFramebufferObject::*)(bool)>(&QQuickFramebufferObject::textureFollowsItemSizeChanged), static_cast<MyQQuickFramebufferObject*>(ptr), static_cast<void (MyQQuickFramebufferObject::*)(bool)>(&MyQQuickFramebufferObject::Signal_TextureFollowsItemSizeChanged));
}

void QQuickFramebufferObject_TextureFollowsItemSizeChanged(void* ptr, int vbo)
{
	static_cast<QQuickFramebufferObject*>(ptr)->textureFollowsItemSizeChanged(vbo != 0);
}

void* QQuickFramebufferObject_TextureProvider(void* ptr)
{
	return static_cast<QQuickFramebufferObject*>(ptr)->textureProvider();
}

void* QQuickFramebufferObject_TextureProviderDefault(void* ptr)
{
	return static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::textureProvider();
}

int QQuickFramebufferObject_ChildMouseEventFilter(void* ptr, void* item, void* event)
{
	return static_cast<QQuickFramebufferObject*>(ptr)->childMouseEventFilter(static_cast<QQuickItem*>(item), static_cast<QEvent*>(event));
}

int QQuickFramebufferObject_ChildMouseEventFilterDefault(void* ptr, void* item, void* event)
{
	return static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::childMouseEventFilter(static_cast<QQuickItem*>(item), static_cast<QEvent*>(event));
}

int QQuickFramebufferObject_Contains(void* ptr, void* point)
{
	return static_cast<QQuickFramebufferObject*>(ptr)->contains(*static_cast<QPointF*>(point));
}

int QQuickFramebufferObject_ContainsDefault(void* ptr, void* point)
{
	return static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::contains(*static_cast<QPointF*>(point));
}

void QQuickFramebufferObject_DragEnterEvent(void* ptr, void* event)
{
	static_cast<QQuickFramebufferObject*>(ptr)->dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QQuickFramebufferObject_DragEnterEventDefault(void* ptr, void* event)
{
	static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QQuickFramebufferObject_DragLeaveEvent(void* ptr, void* event)
{
	static_cast<QQuickFramebufferObject*>(ptr)->dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QQuickFramebufferObject_DragLeaveEventDefault(void* ptr, void* event)
{
	static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QQuickFramebufferObject_DragMoveEvent(void* ptr, void* event)
{
	static_cast<QQuickFramebufferObject*>(ptr)->dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QQuickFramebufferObject_DragMoveEventDefault(void* ptr, void* event)
{
	static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QQuickFramebufferObject_DropEvent(void* ptr, void* event)
{
	static_cast<QQuickFramebufferObject*>(ptr)->dropEvent(static_cast<QDropEvent*>(event));
}

void QQuickFramebufferObject_DropEventDefault(void* ptr, void* event)
{
	static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::dropEvent(static_cast<QDropEvent*>(event));
}

void QQuickFramebufferObject_FocusInEvent(void* ptr, void* event)
{
	static_cast<QQuickFramebufferObject*>(ptr)->focusInEvent(static_cast<QFocusEvent*>(event));
}

void QQuickFramebufferObject_FocusInEventDefault(void* ptr, void* event)
{
	static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::focusInEvent(static_cast<QFocusEvent*>(event));
}

void QQuickFramebufferObject_FocusOutEvent(void* ptr, void* event)
{
	static_cast<QQuickFramebufferObject*>(ptr)->focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QQuickFramebufferObject_FocusOutEventDefault(void* ptr, void* event)
{
	static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QQuickFramebufferObject_GeometryChanged(void* ptr, void* newGeometry, void* oldGeometry)
{
	static_cast<QQuickFramebufferObject*>(ptr)->geometryChanged(*static_cast<QRectF*>(newGeometry), *static_cast<QRectF*>(oldGeometry));
}

void QQuickFramebufferObject_GeometryChangedDefault(void* ptr, void* newGeometry, void* oldGeometry)
{
	static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::geometryChanged(*static_cast<QRectF*>(newGeometry), *static_cast<QRectF*>(oldGeometry));
}

void QQuickFramebufferObject_HoverEnterEvent(void* ptr, void* event)
{
	static_cast<QQuickFramebufferObject*>(ptr)->hoverEnterEvent(static_cast<QHoverEvent*>(event));
}

void QQuickFramebufferObject_HoverEnterEventDefault(void* ptr, void* event)
{
	static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::hoverEnterEvent(static_cast<QHoverEvent*>(event));
}

void QQuickFramebufferObject_HoverLeaveEvent(void* ptr, void* event)
{
	static_cast<QQuickFramebufferObject*>(ptr)->hoverLeaveEvent(static_cast<QHoverEvent*>(event));
}

void QQuickFramebufferObject_HoverLeaveEventDefault(void* ptr, void* event)
{
	static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::hoverLeaveEvent(static_cast<QHoverEvent*>(event));
}

void QQuickFramebufferObject_HoverMoveEvent(void* ptr, void* event)
{
	static_cast<QQuickFramebufferObject*>(ptr)->hoverMoveEvent(static_cast<QHoverEvent*>(event));
}

void QQuickFramebufferObject_HoverMoveEventDefault(void* ptr, void* event)
{
	static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::hoverMoveEvent(static_cast<QHoverEvent*>(event));
}

void QQuickFramebufferObject_InputMethodEvent(void* ptr, void* event)
{
	static_cast<QQuickFramebufferObject*>(ptr)->inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void QQuickFramebufferObject_InputMethodEventDefault(void* ptr, void* event)
{
	static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void* QQuickFramebufferObject_InputMethodQuery(void* ptr, int query)
{
	return new QVariant(static_cast<QQuickFramebufferObject*>(ptr)->inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

void* QQuickFramebufferObject_InputMethodQueryDefault(void* ptr, int query)
{
	return new QVariant(static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

void QQuickFramebufferObject_KeyPressEvent(void* ptr, void* event)
{
	static_cast<QQuickFramebufferObject*>(ptr)->keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QQuickFramebufferObject_KeyPressEventDefault(void* ptr, void* event)
{
	static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QQuickFramebufferObject_KeyReleaseEvent(void* ptr, void* event)
{
	static_cast<QQuickFramebufferObject*>(ptr)->keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QQuickFramebufferObject_KeyReleaseEventDefault(void* ptr, void* event)
{
	static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QQuickFramebufferObject_MouseDoubleClickEvent(void* ptr, void* event)
{
	static_cast<QQuickFramebufferObject*>(ptr)->mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QQuickFramebufferObject_MouseDoubleClickEventDefault(void* ptr, void* event)
{
	static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QQuickFramebufferObject_MouseMoveEvent(void* ptr, void* event)
{
	static_cast<QQuickFramebufferObject*>(ptr)->mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QQuickFramebufferObject_MouseMoveEventDefault(void* ptr, void* event)
{
	static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QQuickFramebufferObject_MousePressEvent(void* ptr, void* event)
{
	static_cast<QQuickFramebufferObject*>(ptr)->mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QQuickFramebufferObject_MousePressEventDefault(void* ptr, void* event)
{
	static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QQuickFramebufferObject_MouseReleaseEvent(void* ptr, void* event)
{
	static_cast<QQuickFramebufferObject*>(ptr)->mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void QQuickFramebufferObject_MouseReleaseEventDefault(void* ptr, void* event)
{
	static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void QQuickFramebufferObject_MouseUngrabEvent(void* ptr)
{
	static_cast<QQuickFramebufferObject*>(ptr)->mouseUngrabEvent();
}

void QQuickFramebufferObject_MouseUngrabEventDefault(void* ptr)
{
	static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::mouseUngrabEvent();
}

void QQuickFramebufferObject_TouchEvent(void* ptr, void* event)
{
	static_cast<QQuickFramebufferObject*>(ptr)->touchEvent(static_cast<QTouchEvent*>(event));
}

void QQuickFramebufferObject_TouchEventDefault(void* ptr, void* event)
{
	static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::touchEvent(static_cast<QTouchEvent*>(event));
}

void QQuickFramebufferObject_TouchUngrabEvent(void* ptr)
{
	static_cast<QQuickFramebufferObject*>(ptr)->touchUngrabEvent();
}

void QQuickFramebufferObject_TouchUngrabEventDefault(void* ptr)
{
	static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::touchUngrabEvent();
}

void QQuickFramebufferObject_Update(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QQuickFramebufferObject*>(ptr), "update");
}

void QQuickFramebufferObject_UpdateDefault(void* ptr)
{
	static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::update();
}

void QQuickFramebufferObject_UpdatePolish(void* ptr)
{
	static_cast<QQuickFramebufferObject*>(ptr)->updatePolish();
}

void QQuickFramebufferObject_UpdatePolishDefault(void* ptr)
{
	static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::updatePolish();
}

void QQuickFramebufferObject_WheelEvent(void* ptr, void* event)
{
	static_cast<QQuickFramebufferObject*>(ptr)->wheelEvent(static_cast<QWheelEvent*>(event));
}

void QQuickFramebufferObject_WheelEventDefault(void* ptr, void* event)
{
	static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::wheelEvent(static_cast<QWheelEvent*>(event));
}

void QQuickFramebufferObject_TimerEvent(void* ptr, void* event)
{
	static_cast<QQuickFramebufferObject*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QQuickFramebufferObject_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::timerEvent(static_cast<QTimerEvent*>(event));
}

void QQuickFramebufferObject_ChildEvent(void* ptr, void* event)
{
	static_cast<QQuickFramebufferObject*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QQuickFramebufferObject_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::childEvent(static_cast<QChildEvent*>(event));
}

void QQuickFramebufferObject_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QQuickFramebufferObject*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQuickFramebufferObject_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQuickFramebufferObject_CustomEvent(void* ptr, void* event)
{
	static_cast<QQuickFramebufferObject*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QQuickFramebufferObject_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::customEvent(static_cast<QEvent*>(event));
}

void QQuickFramebufferObject_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QQuickFramebufferObject*>(ptr), "deleteLater");
}

void QQuickFramebufferObject_DeleteLaterDefault(void* ptr)
{
	static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::deleteLater();
}

void QQuickFramebufferObject_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QQuickFramebufferObject*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQuickFramebufferObject_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QQuickFramebufferObject_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QQuickFramebufferObject*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QQuickFramebufferObject_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QQuickFramebufferObject_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QQuickFramebufferObject*>(ptr)->metaObject());
}

void* QQuickFramebufferObject_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::metaObject());
}

class MyQQuickImageProvider: public QQuickImageProvider
{
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	MyQQuickImageProvider(ImageType type, Flags flags) : QQuickImageProvider(type, flags) {};
};

void* QQuickImageProvider_NewQQuickImageProvider(int ty, int flags)
{
	return new MyQQuickImageProvider(static_cast<QQmlImageProviderBase::ImageType>(ty), static_cast<QQmlImageProviderBase::Flag>(flags));
}

int QQuickImageProvider_Flags(void* ptr)
{
	return static_cast<QQuickImageProvider*>(ptr)->flags();
}

int QQuickImageProvider_ImageType(void* ptr)
{
	return static_cast<QQuickImageProvider*>(ptr)->imageType();
}

void QQuickImageProvider_DestroyQQuickImageProvider(void* ptr)
{
	static_cast<QQuickImageProvider*>(ptr)->~QQuickImageProvider();
}

char* QQuickImageProvider_ObjectNameAbs(void* ptr)
{
	if (dynamic_cast<MyQQuickImageProvider*>(static_cast<QQuickImageProvider*>(ptr))) {
		return static_cast<MyQQuickImageProvider*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QQuickImageProvider_BASE").toUtf8().data();
}

void QQuickImageProvider_SetObjectNameAbs(void* ptr, char* name)
{
	if (dynamic_cast<MyQQuickImageProvider*>(static_cast<QQuickImageProvider*>(ptr))) {
		static_cast<MyQQuickImageProvider*>(ptr)->setObjectNameAbs(QString(name));
	}
}

class MyQQuickImageResponse: public QQuickImageResponse
{
public:
	MyQQuickImageResponse() : QQuickImageResponse() {};
	void cancel() { callbackQQuickImageResponse_Cancel(this, this->objectName().toUtf8().data()); };
	QString errorString() const { return QString(callbackQQuickImageResponse_ErrorString(const_cast<MyQQuickImageResponse*>(this), this->objectName().toUtf8().data())); };
	void Signal_Finished() { callbackQQuickImageResponse_Finished(this, this->objectName().toUtf8().data()); };
	QQuickTextureFactory * textureFactory() const { return static_cast<QQuickTextureFactory*>(callbackQQuickImageResponse_TextureFactory(const_cast<MyQQuickImageResponse*>(this), this->objectName().toUtf8().data())); };
	void timerEvent(QTimerEvent * event) { callbackQQuickImageResponse_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQQuickImageResponse_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQQuickImageResponse_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQQuickImageResponse_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQQuickImageResponse_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQQuickImageResponse_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQQuickImageResponse_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQQuickImageResponse_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQQuickImageResponse_MetaObject(const_cast<MyQQuickImageResponse*>(this), this->objectName().toUtf8().data())); };
};

void* QQuickImageResponse_NewQQuickImageResponse()
{
	return new MyQQuickImageResponse();
}

void QQuickImageResponse_Cancel(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QQuickImageResponse*>(ptr), "cancel");
}

void QQuickImageResponse_CancelDefault(void* ptr)
{
	static_cast<QQuickImageResponse*>(ptr)->QQuickImageResponse::cancel();
}

char* QQuickImageResponse_ErrorString(void* ptr)
{
	return static_cast<QQuickImageResponse*>(ptr)->errorString().toUtf8().data();
}

char* QQuickImageResponse_ErrorStringDefault(void* ptr)
{
	return static_cast<QQuickImageResponse*>(ptr)->QQuickImageResponse::errorString().toUtf8().data();
}

void QQuickImageResponse_ConnectFinished(void* ptr)
{
	QObject::connect(static_cast<QQuickImageResponse*>(ptr), static_cast<void (QQuickImageResponse::*)()>(&QQuickImageResponse::finished), static_cast<MyQQuickImageResponse*>(ptr), static_cast<void (MyQQuickImageResponse::*)()>(&MyQQuickImageResponse::Signal_Finished));
}

void QQuickImageResponse_DisconnectFinished(void* ptr)
{
	QObject::disconnect(static_cast<QQuickImageResponse*>(ptr), static_cast<void (QQuickImageResponse::*)()>(&QQuickImageResponse::finished), static_cast<MyQQuickImageResponse*>(ptr), static_cast<void (MyQQuickImageResponse::*)()>(&MyQQuickImageResponse::Signal_Finished));
}

void QQuickImageResponse_Finished(void* ptr)
{
	static_cast<QQuickImageResponse*>(ptr)->finished();
}

void* QQuickImageResponse_TextureFactory(void* ptr)
{
	return static_cast<QQuickImageResponse*>(ptr)->textureFactory();
}

void QQuickImageResponse_DestroyQQuickImageResponse(void* ptr)
{
	static_cast<QQuickImageResponse*>(ptr)->~QQuickImageResponse();
}

void QQuickImageResponse_TimerEvent(void* ptr, void* event)
{
	static_cast<QQuickImageResponse*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QQuickImageResponse_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QQuickImageResponse*>(ptr)->QQuickImageResponse::timerEvent(static_cast<QTimerEvent*>(event));
}

void QQuickImageResponse_ChildEvent(void* ptr, void* event)
{
	static_cast<QQuickImageResponse*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QQuickImageResponse_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QQuickImageResponse*>(ptr)->QQuickImageResponse::childEvent(static_cast<QChildEvent*>(event));
}

void QQuickImageResponse_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QQuickImageResponse*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQuickImageResponse_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QQuickImageResponse*>(ptr)->QQuickImageResponse::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQuickImageResponse_CustomEvent(void* ptr, void* event)
{
	static_cast<QQuickImageResponse*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QQuickImageResponse_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QQuickImageResponse*>(ptr)->QQuickImageResponse::customEvent(static_cast<QEvent*>(event));
}

void QQuickImageResponse_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QQuickImageResponse*>(ptr), "deleteLater");
}

void QQuickImageResponse_DeleteLaterDefault(void* ptr)
{
	static_cast<QQuickImageResponse*>(ptr)->QQuickImageResponse::deleteLater();
}

void QQuickImageResponse_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QQuickImageResponse*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQuickImageResponse_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QQuickImageResponse*>(ptr)->QQuickImageResponse::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QQuickImageResponse_Event(void* ptr, void* e)
{
	return static_cast<QQuickImageResponse*>(ptr)->event(static_cast<QEvent*>(e));
}

int QQuickImageResponse_EventDefault(void* ptr, void* e)
{
	return static_cast<QQuickImageResponse*>(ptr)->QQuickImageResponse::event(static_cast<QEvent*>(e));
}

int QQuickImageResponse_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QQuickImageResponse*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QQuickImageResponse_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QQuickImageResponse*>(ptr)->QQuickImageResponse::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QQuickImageResponse_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QQuickImageResponse*>(ptr)->metaObject());
}

void* QQuickImageResponse_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QQuickImageResponse*>(ptr)->QQuickImageResponse::metaObject());
}

class MyQQuickItem: public QQuickItem
{
public:
	MyQQuickItem(QQuickItem *parent) : QQuickItem(parent) {};
	bool isTextureProvider() const { return callbackQQuickItem_IsTextureProvider(const_cast<MyQQuickItem*>(this), this->objectName().toUtf8().data()) != 0; };
	QSGTextureProvider * textureProvider() const { return static_cast<QSGTextureProvider*>(callbackQQuickItem_TextureProvider(const_cast<MyQQuickItem*>(this), this->objectName().toUtf8().data())); };
	bool childMouseEventFilter(QQuickItem * item, QEvent * event) { return callbackQQuickItem_ChildMouseEventFilter(this, this->objectName().toUtf8().data(), item, event) != 0; };
	bool contains(const QPointF & point) const { return callbackQQuickItem_Contains(const_cast<MyQQuickItem*>(this), this->objectName().toUtf8().data(), new QPointF(static_cast<QPointF>(point).x(), static_cast<QPointF>(point).y())) != 0; };
	void dragEnterEvent(QDragEnterEvent * event) { callbackQQuickItem_DragEnterEvent(this, this->objectName().toUtf8().data(), event); };
	void dragLeaveEvent(QDragLeaveEvent * event) { callbackQQuickItem_DragLeaveEvent(this, this->objectName().toUtf8().data(), event); };
	void dragMoveEvent(QDragMoveEvent * event) { callbackQQuickItem_DragMoveEvent(this, this->objectName().toUtf8().data(), event); };
	void dropEvent(QDropEvent * event) { callbackQQuickItem_DropEvent(this, this->objectName().toUtf8().data(), event); };
	void focusInEvent(QFocusEvent * event) { callbackQQuickItem_FocusInEvent(this, this->objectName().toUtf8().data(), event); };
	void focusOutEvent(QFocusEvent * event) { callbackQQuickItem_FocusOutEvent(this, this->objectName().toUtf8().data(), event); };
	void geometryChanged(const QRectF & newGeometry, const QRectF & oldGeometry) { callbackQQuickItem_GeometryChanged(this, this->objectName().toUtf8().data(), new QRectF(static_cast<QRectF>(newGeometry).x(), static_cast<QRectF>(newGeometry).y(), static_cast<QRectF>(newGeometry).width(), static_cast<QRectF>(newGeometry).height()), new QRectF(static_cast<QRectF>(oldGeometry).x(), static_cast<QRectF>(oldGeometry).y(), static_cast<QRectF>(oldGeometry).width(), static_cast<QRectF>(oldGeometry).height())); };
	void hoverEnterEvent(QHoverEvent * event) { callbackQQuickItem_HoverEnterEvent(this, this->objectName().toUtf8().data(), event); };
	void hoverLeaveEvent(QHoverEvent * event) { callbackQQuickItem_HoverLeaveEvent(this, this->objectName().toUtf8().data(), event); };
	void hoverMoveEvent(QHoverEvent * event) { callbackQQuickItem_HoverMoveEvent(this, this->objectName().toUtf8().data(), event); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackQQuickItem_InputMethodEvent(this, this->objectName().toUtf8().data(), event); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackQQuickItem_InputMethodQuery(const_cast<MyQQuickItem*>(this), this->objectName().toUtf8().data(), query)); };
	void keyPressEvent(QKeyEvent * event) { callbackQQuickItem_KeyPressEvent(this, this->objectName().toUtf8().data(), event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQQuickItem_KeyReleaseEvent(this, this->objectName().toUtf8().data(), event); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQQuickItem_MouseDoubleClickEvent(this, this->objectName().toUtf8().data(), event); };
	void mouseMoveEvent(QMouseEvent * event) { callbackQQuickItem_MouseMoveEvent(this, this->objectName().toUtf8().data(), event); };
	void mousePressEvent(QMouseEvent * event) { callbackQQuickItem_MousePressEvent(this, this->objectName().toUtf8().data(), event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackQQuickItem_MouseReleaseEvent(this, this->objectName().toUtf8().data(), event); };
	void mouseUngrabEvent() { callbackQQuickItem_MouseUngrabEvent(this, this->objectName().toUtf8().data()); };
	void releaseResources() { callbackQQuickItem_ReleaseResources(this, this->objectName().toUtf8().data()); };
	void touchEvent(QTouchEvent * event) { callbackQQuickItem_TouchEvent(this, this->objectName().toUtf8().data(), event); };
	void touchUngrabEvent() { callbackQQuickItem_TouchUngrabEvent(this, this->objectName().toUtf8().data()); };
	void update() { callbackQQuickItem_Update(this, this->objectName().toUtf8().data()); };
	void updatePolish() { callbackQQuickItem_UpdatePolish(this, this->objectName().toUtf8().data()); };
	void wheelEvent(QWheelEvent * event) { callbackQQuickItem_WheelEvent(this, this->objectName().toUtf8().data(), event); };
	void Signal_WindowChanged(QQuickWindow * window) { callbackQQuickItem_WindowChanged(this, this->objectName().toUtf8().data(), window); };
	void timerEvent(QTimerEvent * event) { callbackQQuickItem_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQQuickItem_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQQuickItem_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQQuickItem_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQQuickItem_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQQuickItem_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQQuickItem_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQQuickItem_MetaObject(const_cast<MyQQuickItem*>(this), this->objectName().toUtf8().data())); };
};

void* QQuickItem_NewQQuickItem(void* parent)
{
	return new MyQQuickItem(static_cast<QQuickItem*>(parent));
}

int QQuickItem_ActiveFocusOnTab(void* ptr)
{
	return static_cast<QQuickItem*>(ptr)->activeFocusOnTab();
}

int QQuickItem_Antialiasing(void* ptr)
{
	return static_cast<QQuickItem*>(ptr)->antialiasing();
}

double QQuickItem_BaselineOffset(void* ptr)
{
	return static_cast<double>(static_cast<QQuickItem*>(ptr)->baselineOffset());
}

void* QQuickItem_ChildrenRect(void* ptr)
{
	return new QRectF(static_cast<QRectF>(static_cast<QQuickItem*>(ptr)->childrenRect()).x(), static_cast<QRectF>(static_cast<QQuickItem*>(ptr)->childrenRect()).y(), static_cast<QRectF>(static_cast<QQuickItem*>(ptr)->childrenRect()).width(), static_cast<QRectF>(static_cast<QQuickItem*>(ptr)->childrenRect()).height());
}

int QQuickItem_Clip(void* ptr)
{
	return static_cast<QQuickItem*>(ptr)->clip();
}

int QQuickItem_HasActiveFocus(void* ptr)
{
	return static_cast<QQuickItem*>(ptr)->hasActiveFocus();
}

int QQuickItem_HasFocus(void* ptr)
{
	return static_cast<QQuickItem*>(ptr)->hasFocus();
}

double QQuickItem_Height(void* ptr)
{
	return static_cast<double>(static_cast<QQuickItem*>(ptr)->height());
}

double QQuickItem_ImplicitHeight(void* ptr)
{
	return static_cast<double>(static_cast<QQuickItem*>(ptr)->implicitHeight());
}

int QQuickItem_IsEnabled(void* ptr)
{
	return static_cast<QQuickItem*>(ptr)->isEnabled();
}

int QQuickItem_IsTextureProvider(void* ptr)
{
	return static_cast<QQuickItem*>(ptr)->isTextureProvider();
}

int QQuickItem_IsTextureProviderDefault(void* ptr)
{
	return static_cast<QQuickItem*>(ptr)->QQuickItem::isTextureProvider();
}

int QQuickItem_IsVisible(void* ptr)
{
	return static_cast<QQuickItem*>(ptr)->isVisible();
}

double QQuickItem_Opacity(void* ptr)
{
	return static_cast<double>(static_cast<QQuickItem*>(ptr)->opacity());
}

void* QQuickItem_ParentItem(void* ptr)
{
	return static_cast<QQuickItem*>(ptr)->parentItem();
}

void QQuickItem_ResetAntialiasing(void* ptr)
{
	static_cast<QQuickItem*>(ptr)->resetAntialiasing();
}

void QQuickItem_ResetHeight(void* ptr)
{
	static_cast<QQuickItem*>(ptr)->resetHeight();
}

void QQuickItem_ResetWidth(void* ptr)
{
	static_cast<QQuickItem*>(ptr)->resetWidth();
}

double QQuickItem_Rotation(void* ptr)
{
	return static_cast<double>(static_cast<QQuickItem*>(ptr)->rotation());
}

double QQuickItem_Scale(void* ptr)
{
	return static_cast<double>(static_cast<QQuickItem*>(ptr)->scale());
}

void QQuickItem_SetActiveFocusOnTab(void* ptr, int vbo)
{
	static_cast<QQuickItem*>(ptr)->setActiveFocusOnTab(vbo != 0);
}

void QQuickItem_SetAntialiasing(void* ptr, int vbo)
{
	static_cast<QQuickItem*>(ptr)->setAntialiasing(vbo != 0);
}

void QQuickItem_SetBaselineOffset(void* ptr, double vqr)
{
	static_cast<QQuickItem*>(ptr)->setBaselineOffset(static_cast<double>(vqr));
}

void QQuickItem_SetClip(void* ptr, int vbo)
{
	static_cast<QQuickItem*>(ptr)->setClip(vbo != 0);
}

void QQuickItem_SetEnabled(void* ptr, int vbo)
{
	static_cast<QQuickItem*>(ptr)->setEnabled(vbo != 0);
}

void QQuickItem_SetFocus(void* ptr, int vbo)
{
	static_cast<QQuickItem*>(ptr)->setFocus(vbo != 0);
}

void QQuickItem_SetFocus2(void* ptr, int focus, int reason)
{
	static_cast<QQuickItem*>(ptr)->setFocus(focus != 0, static_cast<Qt::FocusReason>(reason));
}

void QQuickItem_SetHeight(void* ptr, double vqr)
{
	static_cast<QQuickItem*>(ptr)->setHeight(static_cast<double>(vqr));
}

void QQuickItem_SetImplicitHeight(void* ptr, double vqr)
{
	static_cast<QQuickItem*>(ptr)->setImplicitHeight(static_cast<double>(vqr));
}

void QQuickItem_SetImplicitWidth(void* ptr, double vqr)
{
	static_cast<QQuickItem*>(ptr)->setImplicitWidth(static_cast<double>(vqr));
}

void QQuickItem_SetOpacity(void* ptr, double vqr)
{
	static_cast<QQuickItem*>(ptr)->setOpacity(static_cast<double>(vqr));
}

void QQuickItem_SetParentItem(void* ptr, void* parent)
{
	static_cast<QQuickItem*>(ptr)->setParentItem(static_cast<QQuickItem*>(parent));
}

void QQuickItem_SetRotation(void* ptr, double vqr)
{
	static_cast<QQuickItem*>(ptr)->setRotation(static_cast<double>(vqr));
}

void QQuickItem_SetScale(void* ptr, double vqr)
{
	static_cast<QQuickItem*>(ptr)->setScale(static_cast<double>(vqr));
}

void QQuickItem_SetSmooth(void* ptr, int vbo)
{
	static_cast<QQuickItem*>(ptr)->setSmooth(vbo != 0);
}

void QQuickItem_SetState(void* ptr, char* vqs)
{
	static_cast<QQuickItem*>(ptr)->setState(QString(vqs));
}

void QQuickItem_SetTransformOrigin(void* ptr, int vtr)
{
	static_cast<QQuickItem*>(ptr)->setTransformOrigin(static_cast<QQuickItem::TransformOrigin>(vtr));
}

void QQuickItem_SetVisible(void* ptr, int vbo)
{
	static_cast<QQuickItem*>(ptr)->setVisible(vbo != 0);
}

void QQuickItem_SetWidth(void* ptr, double vqr)
{
	static_cast<QQuickItem*>(ptr)->setWidth(static_cast<double>(vqr));
}

void QQuickItem_SetX(void* ptr, double vqr)
{
	static_cast<QQuickItem*>(ptr)->setX(static_cast<double>(vqr));
}

void QQuickItem_SetY(void* ptr, double vqr)
{
	static_cast<QQuickItem*>(ptr)->setY(static_cast<double>(vqr));
}

void QQuickItem_SetZ(void* ptr, double vqr)
{
	static_cast<QQuickItem*>(ptr)->setZ(static_cast<double>(vqr));
}

int QQuickItem_Smooth(void* ptr)
{
	return static_cast<QQuickItem*>(ptr)->smooth();
}

char* QQuickItem_State(void* ptr)
{
	return static_cast<QQuickItem*>(ptr)->state().toUtf8().data();
}

void* QQuickItem_TextureProvider(void* ptr)
{
	return static_cast<QQuickItem*>(ptr)->textureProvider();
}

void* QQuickItem_TextureProviderDefault(void* ptr)
{
	return static_cast<QQuickItem*>(ptr)->QQuickItem::textureProvider();
}

int QQuickItem_TransformOrigin(void* ptr)
{
	return static_cast<QQuickItem*>(ptr)->transformOrigin();
}

double QQuickItem_Width(void* ptr)
{
	return static_cast<double>(static_cast<QQuickItem*>(ptr)->width());
}

double QQuickItem_X(void* ptr)
{
	return static_cast<double>(static_cast<QQuickItem*>(ptr)->x());
}

double QQuickItem_Y(void* ptr)
{
	return static_cast<double>(static_cast<QQuickItem*>(ptr)->y());
}

double QQuickItem_Z(void* ptr)
{
	return static_cast<double>(static_cast<QQuickItem*>(ptr)->z());
}

int QQuickItem_AcceptHoverEvents(void* ptr)
{
	return static_cast<QQuickItem*>(ptr)->acceptHoverEvents();
}

int QQuickItem_AcceptedMouseButtons(void* ptr)
{
	return static_cast<QQuickItem*>(ptr)->acceptedMouseButtons();
}

void* QQuickItem_ChildAt(void* ptr, double x, double y)
{
	return static_cast<QQuickItem*>(ptr)->childAt(static_cast<double>(x), static_cast<double>(y));
}

int QQuickItem_ChildMouseEventFilter(void* ptr, void* item, void* event)
{
	return static_cast<QQuickItem*>(ptr)->childMouseEventFilter(static_cast<QQuickItem*>(item), static_cast<QEvent*>(event));
}

int QQuickItem_ChildMouseEventFilterDefault(void* ptr, void* item, void* event)
{
	return static_cast<QQuickItem*>(ptr)->QQuickItem::childMouseEventFilter(static_cast<QQuickItem*>(item), static_cast<QEvent*>(event));
}

void QQuickItem_ClassBegin(void* ptr)
{
	static_cast<QQuickItem*>(ptr)->classBegin();
}

void QQuickItem_ComponentComplete(void* ptr)
{
	static_cast<QQuickItem*>(ptr)->componentComplete();
}

int QQuickItem_Contains(void* ptr, void* point)
{
	return static_cast<QQuickItem*>(ptr)->contains(*static_cast<QPointF*>(point));
}

int QQuickItem_ContainsDefault(void* ptr, void* point)
{
	return static_cast<QQuickItem*>(ptr)->QQuickItem::contains(*static_cast<QPointF*>(point));
}

void* QQuickItem_Cursor(void* ptr)
{
	return new QCursor(static_cast<QQuickItem*>(ptr)->cursor());
}

void QQuickItem_DragEnterEvent(void* ptr, void* event)
{
	static_cast<QQuickItem*>(ptr)->dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QQuickItem_DragEnterEventDefault(void* ptr, void* event)
{
	static_cast<QQuickItem*>(ptr)->QQuickItem::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QQuickItem_DragLeaveEvent(void* ptr, void* event)
{
	static_cast<QQuickItem*>(ptr)->dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QQuickItem_DragLeaveEventDefault(void* ptr, void* event)
{
	static_cast<QQuickItem*>(ptr)->QQuickItem::dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QQuickItem_DragMoveEvent(void* ptr, void* event)
{
	static_cast<QQuickItem*>(ptr)->dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QQuickItem_DragMoveEventDefault(void* ptr, void* event)
{
	static_cast<QQuickItem*>(ptr)->QQuickItem::dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QQuickItem_DropEvent(void* ptr, void* event)
{
	static_cast<QQuickItem*>(ptr)->dropEvent(static_cast<QDropEvent*>(event));
}

void QQuickItem_DropEventDefault(void* ptr, void* event)
{
	static_cast<QQuickItem*>(ptr)->QQuickItem::dropEvent(static_cast<QDropEvent*>(event));
}

int QQuickItem_Event(void* ptr, void* ev)
{
	return static_cast<QQuickItem*>(ptr)->event(static_cast<QEvent*>(ev));
}

int QQuickItem_FiltersChildMouseEvents(void* ptr)
{
	return static_cast<QQuickItem*>(ptr)->filtersChildMouseEvents();
}

int QQuickItem_Flags(void* ptr)
{
	return static_cast<QQuickItem*>(ptr)->flags();
}

void QQuickItem_FocusInEvent(void* ptr, void* event)
{
	static_cast<QQuickItem*>(ptr)->focusInEvent(static_cast<QFocusEvent*>(event));
}

void QQuickItem_FocusInEventDefault(void* ptr, void* event)
{
	static_cast<QQuickItem*>(ptr)->QQuickItem::focusInEvent(static_cast<QFocusEvent*>(event));
}

void QQuickItem_FocusOutEvent(void* ptr, void* event)
{
	static_cast<QQuickItem*>(ptr)->focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QQuickItem_FocusOutEventDefault(void* ptr, void* event)
{
	static_cast<QQuickItem*>(ptr)->QQuickItem::focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QQuickItem_ForceActiveFocus(void* ptr)
{
	static_cast<QQuickItem*>(ptr)->forceActiveFocus();
}

void QQuickItem_ForceActiveFocus2(void* ptr, int reason)
{
	static_cast<QQuickItem*>(ptr)->forceActiveFocus(static_cast<Qt::FocusReason>(reason));
}

void QQuickItem_GeometryChanged(void* ptr, void* newGeometry, void* oldGeometry)
{
	static_cast<QQuickItem*>(ptr)->geometryChanged(*static_cast<QRectF*>(newGeometry), *static_cast<QRectF*>(oldGeometry));
}

void QQuickItem_GeometryChangedDefault(void* ptr, void* newGeometry, void* oldGeometry)
{
	static_cast<QQuickItem*>(ptr)->QQuickItem::geometryChanged(*static_cast<QRectF*>(newGeometry), *static_cast<QRectF*>(oldGeometry));
}

void QQuickItem_GrabMouse(void* ptr)
{
	static_cast<QQuickItem*>(ptr)->grabMouse();
}

int QQuickItem_HeightValid(void* ptr)
{
	return static_cast<QQuickItem*>(ptr)->heightValid();
}

void QQuickItem_HoverEnterEvent(void* ptr, void* event)
{
	static_cast<QQuickItem*>(ptr)->hoverEnterEvent(static_cast<QHoverEvent*>(event));
}

void QQuickItem_HoverEnterEventDefault(void* ptr, void* event)
{
	static_cast<QQuickItem*>(ptr)->QQuickItem::hoverEnterEvent(static_cast<QHoverEvent*>(event));
}

void QQuickItem_HoverLeaveEvent(void* ptr, void* event)
{
	static_cast<QQuickItem*>(ptr)->hoverLeaveEvent(static_cast<QHoverEvent*>(event));
}

void QQuickItem_HoverLeaveEventDefault(void* ptr, void* event)
{
	static_cast<QQuickItem*>(ptr)->QQuickItem::hoverLeaveEvent(static_cast<QHoverEvent*>(event));
}

void QQuickItem_HoverMoveEvent(void* ptr, void* event)
{
	static_cast<QQuickItem*>(ptr)->hoverMoveEvent(static_cast<QHoverEvent*>(event));
}

void QQuickItem_HoverMoveEventDefault(void* ptr, void* event)
{
	static_cast<QQuickItem*>(ptr)->QQuickItem::hoverMoveEvent(static_cast<QHoverEvent*>(event));
}

double QQuickItem_ImplicitWidth(void* ptr)
{
	return static_cast<double>(static_cast<QQuickItem*>(ptr)->implicitWidth());
}

void QQuickItem_InputMethodEvent(void* ptr, void* event)
{
	static_cast<QQuickItem*>(ptr)->inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void QQuickItem_InputMethodEventDefault(void* ptr, void* event)
{
	static_cast<QQuickItem*>(ptr)->QQuickItem::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void* QQuickItem_InputMethodQuery(void* ptr, int query)
{
	return new QVariant(static_cast<QQuickItem*>(ptr)->inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

void* QQuickItem_InputMethodQueryDefault(void* ptr, int query)
{
	return new QVariant(static_cast<QQuickItem*>(ptr)->QQuickItem::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

int QQuickItem_IsComponentComplete(void* ptr)
{
	return static_cast<QQuickItem*>(ptr)->isComponentComplete();
}

int QQuickItem_IsFocusScope(void* ptr)
{
	return static_cast<QQuickItem*>(ptr)->isFocusScope();
}

int QQuickItem_KeepMouseGrab(void* ptr)
{
	return static_cast<QQuickItem*>(ptr)->keepMouseGrab();
}

int QQuickItem_KeepTouchGrab(void* ptr)
{
	return static_cast<QQuickItem*>(ptr)->keepTouchGrab();
}

void QQuickItem_KeyPressEvent(void* ptr, void* event)
{
	static_cast<QQuickItem*>(ptr)->keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QQuickItem_KeyPressEventDefault(void* ptr, void* event)
{
	static_cast<QQuickItem*>(ptr)->QQuickItem::keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QQuickItem_KeyReleaseEvent(void* ptr, void* event)
{
	static_cast<QQuickItem*>(ptr)->keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QQuickItem_KeyReleaseEventDefault(void* ptr, void* event)
{
	static_cast<QQuickItem*>(ptr)->QQuickItem::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void* QQuickItem_MapFromItem(void* ptr, void* item, void* point)
{
	return new QPointF(static_cast<QPointF>(static_cast<QQuickItem*>(ptr)->mapFromItem(static_cast<QQuickItem*>(item), *static_cast<QPointF*>(point))).x(), static_cast<QPointF>(static_cast<QQuickItem*>(ptr)->mapFromItem(static_cast<QQuickItem*>(item), *static_cast<QPointF*>(point))).y());
}

void* QQuickItem_MapFromScene(void* ptr, void* point)
{
	return new QPointF(static_cast<QPointF>(static_cast<QQuickItem*>(ptr)->mapFromScene(*static_cast<QPointF*>(point))).x(), static_cast<QPointF>(static_cast<QQuickItem*>(ptr)->mapFromScene(*static_cast<QPointF*>(point))).y());
}

void* QQuickItem_MapRectFromItem(void* ptr, void* item, void* rect)
{
	return new QRectF(static_cast<QRectF>(static_cast<QQuickItem*>(ptr)->mapRectFromItem(static_cast<QQuickItem*>(item), *static_cast<QRectF*>(rect))).x(), static_cast<QRectF>(static_cast<QQuickItem*>(ptr)->mapRectFromItem(static_cast<QQuickItem*>(item), *static_cast<QRectF*>(rect))).y(), static_cast<QRectF>(static_cast<QQuickItem*>(ptr)->mapRectFromItem(static_cast<QQuickItem*>(item), *static_cast<QRectF*>(rect))).width(), static_cast<QRectF>(static_cast<QQuickItem*>(ptr)->mapRectFromItem(static_cast<QQuickItem*>(item), *static_cast<QRectF*>(rect))).height());
}

void* QQuickItem_MapRectFromScene(void* ptr, void* rect)
{
	return new QRectF(static_cast<QRectF>(static_cast<QQuickItem*>(ptr)->mapRectFromScene(*static_cast<QRectF*>(rect))).x(), static_cast<QRectF>(static_cast<QQuickItem*>(ptr)->mapRectFromScene(*static_cast<QRectF*>(rect))).y(), static_cast<QRectF>(static_cast<QQuickItem*>(ptr)->mapRectFromScene(*static_cast<QRectF*>(rect))).width(), static_cast<QRectF>(static_cast<QQuickItem*>(ptr)->mapRectFromScene(*static_cast<QRectF*>(rect))).height());
}

void* QQuickItem_MapRectToItem(void* ptr, void* item, void* rect)
{
	return new QRectF(static_cast<QRectF>(static_cast<QQuickItem*>(ptr)->mapRectToItem(static_cast<QQuickItem*>(item), *static_cast<QRectF*>(rect))).x(), static_cast<QRectF>(static_cast<QQuickItem*>(ptr)->mapRectToItem(static_cast<QQuickItem*>(item), *static_cast<QRectF*>(rect))).y(), static_cast<QRectF>(static_cast<QQuickItem*>(ptr)->mapRectToItem(static_cast<QQuickItem*>(item), *static_cast<QRectF*>(rect))).width(), static_cast<QRectF>(static_cast<QQuickItem*>(ptr)->mapRectToItem(static_cast<QQuickItem*>(item), *static_cast<QRectF*>(rect))).height());
}

void* QQuickItem_MapRectToScene(void* ptr, void* rect)
{
	return new QRectF(static_cast<QRectF>(static_cast<QQuickItem*>(ptr)->mapRectToScene(*static_cast<QRectF*>(rect))).x(), static_cast<QRectF>(static_cast<QQuickItem*>(ptr)->mapRectToScene(*static_cast<QRectF*>(rect))).y(), static_cast<QRectF>(static_cast<QQuickItem*>(ptr)->mapRectToScene(*static_cast<QRectF*>(rect))).width(), static_cast<QRectF>(static_cast<QQuickItem*>(ptr)->mapRectToScene(*static_cast<QRectF*>(rect))).height());
}

void* QQuickItem_MapToItem(void* ptr, void* item, void* point)
{
	return new QPointF(static_cast<QPointF>(static_cast<QQuickItem*>(ptr)->mapToItem(static_cast<QQuickItem*>(item), *static_cast<QPointF*>(point))).x(), static_cast<QPointF>(static_cast<QQuickItem*>(ptr)->mapToItem(static_cast<QQuickItem*>(item), *static_cast<QPointF*>(point))).y());
}

void* QQuickItem_MapToScene(void* ptr, void* point)
{
	return new QPointF(static_cast<QPointF>(static_cast<QQuickItem*>(ptr)->mapToScene(*static_cast<QPointF*>(point))).x(), static_cast<QPointF>(static_cast<QQuickItem*>(ptr)->mapToScene(*static_cast<QPointF*>(point))).y());
}

void QQuickItem_MouseDoubleClickEvent(void* ptr, void* event)
{
	static_cast<QQuickItem*>(ptr)->mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QQuickItem_MouseDoubleClickEventDefault(void* ptr, void* event)
{
	static_cast<QQuickItem*>(ptr)->QQuickItem::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QQuickItem_MouseMoveEvent(void* ptr, void* event)
{
	static_cast<QQuickItem*>(ptr)->mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QQuickItem_MouseMoveEventDefault(void* ptr, void* event)
{
	static_cast<QQuickItem*>(ptr)->QQuickItem::mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QQuickItem_MousePressEvent(void* ptr, void* event)
{
	static_cast<QQuickItem*>(ptr)->mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QQuickItem_MousePressEventDefault(void* ptr, void* event)
{
	static_cast<QQuickItem*>(ptr)->QQuickItem::mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QQuickItem_MouseReleaseEvent(void* ptr, void* event)
{
	static_cast<QQuickItem*>(ptr)->mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void QQuickItem_MouseReleaseEventDefault(void* ptr, void* event)
{
	static_cast<QQuickItem*>(ptr)->QQuickItem::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void QQuickItem_MouseUngrabEvent(void* ptr)
{
	static_cast<QQuickItem*>(ptr)->mouseUngrabEvent();
}

void QQuickItem_MouseUngrabEventDefault(void* ptr)
{
	static_cast<QQuickItem*>(ptr)->QQuickItem::mouseUngrabEvent();
}

void* QQuickItem_NextItemInFocusChain(void* ptr, int forward)
{
	return static_cast<QQuickItem*>(ptr)->nextItemInFocusChain(forward != 0);
}

void QQuickItem_Polish(void* ptr)
{
	static_cast<QQuickItem*>(ptr)->polish();
}

void QQuickItem_ReleaseResources(void* ptr)
{
	static_cast<QQuickItem*>(ptr)->releaseResources();
}

void QQuickItem_ReleaseResourcesDefault(void* ptr)
{
	static_cast<QQuickItem*>(ptr)->QQuickItem::releaseResources();
}

void* QQuickItem_ScopedFocusItem(void* ptr)
{
	return static_cast<QQuickItem*>(ptr)->scopedFocusItem();
}

void QQuickItem_SetAcceptHoverEvents(void* ptr, int enabled)
{
	static_cast<QQuickItem*>(ptr)->setAcceptHoverEvents(enabled != 0);
}

void QQuickItem_SetAcceptedMouseButtons(void* ptr, int buttons)
{
	static_cast<QQuickItem*>(ptr)->setAcceptedMouseButtons(static_cast<Qt::MouseButton>(buttons));
}

void QQuickItem_SetCursor(void* ptr, void* cursor)
{
	static_cast<QQuickItem*>(ptr)->setCursor(*static_cast<QCursor*>(cursor));
}

void QQuickItem_SetFiltersChildMouseEvents(void* ptr, int filter)
{
	static_cast<QQuickItem*>(ptr)->setFiltersChildMouseEvents(filter != 0);
}

void QQuickItem_SetFlag(void* ptr, int flag, int enabled)
{
	static_cast<QQuickItem*>(ptr)->setFlag(static_cast<QQuickItem::Flag>(flag), enabled != 0);
}

void QQuickItem_SetFlags(void* ptr, int flags)
{
	static_cast<QQuickItem*>(ptr)->setFlags(static_cast<QQuickItem::Flag>(flags));
}

void QQuickItem_SetKeepMouseGrab(void* ptr, int keep)
{
	static_cast<QQuickItem*>(ptr)->setKeepMouseGrab(keep != 0);
}

void QQuickItem_SetKeepTouchGrab(void* ptr, int keep)
{
	static_cast<QQuickItem*>(ptr)->setKeepTouchGrab(keep != 0);
}

void QQuickItem_StackAfter(void* ptr, void* sibling)
{
	static_cast<QQuickItem*>(ptr)->stackAfter(static_cast<QQuickItem*>(sibling));
}

void QQuickItem_StackBefore(void* ptr, void* sibling)
{
	static_cast<QQuickItem*>(ptr)->stackBefore(static_cast<QQuickItem*>(sibling));
}

void QQuickItem_TouchEvent(void* ptr, void* event)
{
	static_cast<QQuickItem*>(ptr)->touchEvent(static_cast<QTouchEvent*>(event));
}

void QQuickItem_TouchEventDefault(void* ptr, void* event)
{
	static_cast<QQuickItem*>(ptr)->QQuickItem::touchEvent(static_cast<QTouchEvent*>(event));
}

void QQuickItem_TouchUngrabEvent(void* ptr)
{
	static_cast<QQuickItem*>(ptr)->touchUngrabEvent();
}

void QQuickItem_TouchUngrabEventDefault(void* ptr)
{
	static_cast<QQuickItem*>(ptr)->QQuickItem::touchUngrabEvent();
}

void QQuickItem_UngrabMouse(void* ptr)
{
	static_cast<QQuickItem*>(ptr)->ungrabMouse();
}

void QQuickItem_UngrabTouchPoints(void* ptr)
{
	static_cast<QQuickItem*>(ptr)->ungrabTouchPoints();
}

void QQuickItem_UnsetCursor(void* ptr)
{
	static_cast<QQuickItem*>(ptr)->unsetCursor();
}

void QQuickItem_Update(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QQuickItem*>(ptr), "update");
}

void QQuickItem_UpdateInputMethod(void* ptr, int queries)
{
	static_cast<QQuickItem*>(ptr)->updateInputMethod(static_cast<Qt::InputMethodQuery>(queries));
}

void QQuickItem_UpdatePolish(void* ptr)
{
	static_cast<QQuickItem*>(ptr)->updatePolish();
}

void QQuickItem_UpdatePolishDefault(void* ptr)
{
	static_cast<QQuickItem*>(ptr)->QQuickItem::updatePolish();
}

void QQuickItem_WheelEvent(void* ptr, void* event)
{
	static_cast<QQuickItem*>(ptr)->wheelEvent(static_cast<QWheelEvent*>(event));
}

void QQuickItem_WheelEventDefault(void* ptr, void* event)
{
	static_cast<QQuickItem*>(ptr)->QQuickItem::wheelEvent(static_cast<QWheelEvent*>(event));
}

int QQuickItem_WidthValid(void* ptr)
{
	return static_cast<QQuickItem*>(ptr)->widthValid();
}

void* QQuickItem_Window(void* ptr)
{
	return static_cast<QQuickItem*>(ptr)->window();
}

void QQuickItem_ConnectWindowChanged(void* ptr)
{
	QObject::connect(static_cast<QQuickItem*>(ptr), static_cast<void (QQuickItem::*)(QQuickWindow *)>(&QQuickItem::windowChanged), static_cast<MyQQuickItem*>(ptr), static_cast<void (MyQQuickItem::*)(QQuickWindow *)>(&MyQQuickItem::Signal_WindowChanged));
}

void QQuickItem_DisconnectWindowChanged(void* ptr)
{
	QObject::disconnect(static_cast<QQuickItem*>(ptr), static_cast<void (QQuickItem::*)(QQuickWindow *)>(&QQuickItem::windowChanged), static_cast<MyQQuickItem*>(ptr), static_cast<void (MyQQuickItem::*)(QQuickWindow *)>(&MyQQuickItem::Signal_WindowChanged));
}

void QQuickItem_WindowChanged(void* ptr, void* window)
{
	static_cast<QQuickItem*>(ptr)->windowChanged(static_cast<QQuickWindow*>(window));
}

void QQuickItem_DestroyQQuickItem(void* ptr)
{
	static_cast<QQuickItem*>(ptr)->~QQuickItem();
}

void QQuickItem_TimerEvent(void* ptr, void* event)
{
	static_cast<QQuickItem*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QQuickItem_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QQuickItem*>(ptr)->QQuickItem::timerEvent(static_cast<QTimerEvent*>(event));
}

void QQuickItem_ChildEvent(void* ptr, void* event)
{
	static_cast<QQuickItem*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QQuickItem_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QQuickItem*>(ptr)->QQuickItem::childEvent(static_cast<QChildEvent*>(event));
}

void QQuickItem_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QQuickItem*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQuickItem_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QQuickItem*>(ptr)->QQuickItem::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQuickItem_CustomEvent(void* ptr, void* event)
{
	static_cast<QQuickItem*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QQuickItem_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QQuickItem*>(ptr)->QQuickItem::customEvent(static_cast<QEvent*>(event));
}

void QQuickItem_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QQuickItem*>(ptr), "deleteLater");
}

void QQuickItem_DeleteLaterDefault(void* ptr)
{
	static_cast<QQuickItem*>(ptr)->QQuickItem::deleteLater();
}

void QQuickItem_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QQuickItem*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQuickItem_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QQuickItem*>(ptr)->QQuickItem::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QQuickItem_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QQuickItem*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QQuickItem_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QQuickItem*>(ptr)->QQuickItem::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QQuickItem_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QQuickItem*>(ptr)->metaObject());
}

void* QQuickItem_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QQuickItem*>(ptr)->QQuickItem::metaObject());
}

class MyQQuickItemGrabResult: public QQuickItemGrabResult
{
public:
	void Signal_Ready() { callbackQQuickItemGrabResult_Ready(this, this->objectName().toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQQuickItemGrabResult_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQQuickItemGrabResult_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQQuickItemGrabResult_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQQuickItemGrabResult_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQQuickItemGrabResult_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQQuickItemGrabResult_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQQuickItemGrabResult_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQQuickItemGrabResult_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQQuickItemGrabResult_MetaObject(const_cast<MyQQuickItemGrabResult*>(this), this->objectName().toUtf8().data())); };
};

void* QQuickItemGrabResult_Image(void* ptr)
{
	return new QImage(static_cast<QQuickItemGrabResult*>(ptr)->image());
}

void* QQuickItemGrabResult_Url(void* ptr)
{
	return new QUrl(static_cast<QQuickItemGrabResult*>(ptr)->url());
}

void QQuickItemGrabResult_ConnectReady(void* ptr)
{
	QObject::connect(static_cast<QQuickItemGrabResult*>(ptr), static_cast<void (QQuickItemGrabResult::*)()>(&QQuickItemGrabResult::ready), static_cast<MyQQuickItemGrabResult*>(ptr), static_cast<void (MyQQuickItemGrabResult::*)()>(&MyQQuickItemGrabResult::Signal_Ready));
}

void QQuickItemGrabResult_DisconnectReady(void* ptr)
{
	QObject::disconnect(static_cast<QQuickItemGrabResult*>(ptr), static_cast<void (QQuickItemGrabResult::*)()>(&QQuickItemGrabResult::ready), static_cast<MyQQuickItemGrabResult*>(ptr), static_cast<void (MyQQuickItemGrabResult::*)()>(&MyQQuickItemGrabResult::Signal_Ready));
}

void QQuickItemGrabResult_Ready(void* ptr)
{
	static_cast<QQuickItemGrabResult*>(ptr)->ready();
}

int QQuickItemGrabResult_SaveToFile(void* ptr, char* fileName)
{
	return static_cast<QQuickItemGrabResult*>(ptr)->saveToFile(QString(fileName));
}

void QQuickItemGrabResult_TimerEvent(void* ptr, void* event)
{
	static_cast<QQuickItemGrabResult*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QQuickItemGrabResult_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QQuickItemGrabResult*>(ptr)->QQuickItemGrabResult::timerEvent(static_cast<QTimerEvent*>(event));
}

void QQuickItemGrabResult_ChildEvent(void* ptr, void* event)
{
	static_cast<QQuickItemGrabResult*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QQuickItemGrabResult_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QQuickItemGrabResult*>(ptr)->QQuickItemGrabResult::childEvent(static_cast<QChildEvent*>(event));
}

void QQuickItemGrabResult_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QQuickItemGrabResult*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQuickItemGrabResult_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QQuickItemGrabResult*>(ptr)->QQuickItemGrabResult::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQuickItemGrabResult_CustomEvent(void* ptr, void* event)
{
	static_cast<QQuickItemGrabResult*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QQuickItemGrabResult_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QQuickItemGrabResult*>(ptr)->QQuickItemGrabResult::customEvent(static_cast<QEvent*>(event));
}

void QQuickItemGrabResult_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QQuickItemGrabResult*>(ptr), "deleteLater");
}

void QQuickItemGrabResult_DeleteLaterDefault(void* ptr)
{
	static_cast<QQuickItemGrabResult*>(ptr)->QQuickItemGrabResult::deleteLater();
}

void QQuickItemGrabResult_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QQuickItemGrabResult*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQuickItemGrabResult_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QQuickItemGrabResult*>(ptr)->QQuickItemGrabResult::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QQuickItemGrabResult_Event(void* ptr, void* e)
{
	return static_cast<QQuickItemGrabResult*>(ptr)->event(static_cast<QEvent*>(e));
}

int QQuickItemGrabResult_EventDefault(void* ptr, void* e)
{
	return static_cast<QQuickItemGrabResult*>(ptr)->QQuickItemGrabResult::event(static_cast<QEvent*>(e));
}

int QQuickItemGrabResult_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QQuickItemGrabResult*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QQuickItemGrabResult_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QQuickItemGrabResult*>(ptr)->QQuickItemGrabResult::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QQuickItemGrabResult_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QQuickItemGrabResult*>(ptr)->metaObject());
}

void* QQuickItemGrabResult_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QQuickItemGrabResult*>(ptr)->QQuickItemGrabResult::metaObject());
}

class MyQQuickPaintedItem: public QQuickPaintedItem
{
public:
	MyQQuickPaintedItem(QQuickItem *parent) : QQuickPaintedItem(parent) {};
	void Signal_ContentsScaleChanged() { callbackQQuickPaintedItem_ContentsScaleChanged(this, this->objectName().toUtf8().data()); };
	void Signal_ContentsSizeChanged() { callbackQQuickPaintedItem_ContentsSizeChanged(this, this->objectName().toUtf8().data()); };
	void Signal_FillColorChanged() { callbackQQuickPaintedItem_FillColorChanged(this, this->objectName().toUtf8().data()); };
	bool isTextureProvider() const { return callbackQQuickPaintedItem_IsTextureProvider(const_cast<MyQQuickPaintedItem*>(this), this->objectName().toUtf8().data()) != 0; };
	void paint(QPainter * painter) { callbackQQuickPaintedItem_Paint(this, this->objectName().toUtf8().data(), painter); };
	void releaseResources() { callbackQQuickPaintedItem_ReleaseResources(this, this->objectName().toUtf8().data()); };
	void Signal_RenderTargetChanged() { callbackQQuickPaintedItem_RenderTargetChanged(this, this->objectName().toUtf8().data()); };
	QSGTextureProvider * textureProvider() const { return static_cast<QSGTextureProvider*>(callbackQQuickPaintedItem_TextureProvider(const_cast<MyQQuickPaintedItem*>(this), this->objectName().toUtf8().data())); };
	void Signal_TextureSizeChanged() { callbackQQuickPaintedItem_TextureSizeChanged(this, this->objectName().toUtf8().data()); };
	bool childMouseEventFilter(QQuickItem * item, QEvent * event) { return callbackQQuickPaintedItem_ChildMouseEventFilter(this, this->objectName().toUtf8().data(), item, event) != 0; };
	bool contains(const QPointF & point) const { return callbackQQuickPaintedItem_Contains(const_cast<MyQQuickPaintedItem*>(this), this->objectName().toUtf8().data(), new QPointF(static_cast<QPointF>(point).x(), static_cast<QPointF>(point).y())) != 0; };
	void dragEnterEvent(QDragEnterEvent * event) { callbackQQuickPaintedItem_DragEnterEvent(this, this->objectName().toUtf8().data(), event); };
	void dragLeaveEvent(QDragLeaveEvent * event) { callbackQQuickPaintedItem_DragLeaveEvent(this, this->objectName().toUtf8().data(), event); };
	void dragMoveEvent(QDragMoveEvent * event) { callbackQQuickPaintedItem_DragMoveEvent(this, this->objectName().toUtf8().data(), event); };
	void dropEvent(QDropEvent * event) { callbackQQuickPaintedItem_DropEvent(this, this->objectName().toUtf8().data(), event); };
	void focusInEvent(QFocusEvent * event) { callbackQQuickPaintedItem_FocusInEvent(this, this->objectName().toUtf8().data(), event); };
	void focusOutEvent(QFocusEvent * event) { callbackQQuickPaintedItem_FocusOutEvent(this, this->objectName().toUtf8().data(), event); };
	void geometryChanged(const QRectF & newGeometry, const QRectF & oldGeometry) { callbackQQuickPaintedItem_GeometryChanged(this, this->objectName().toUtf8().data(), new QRectF(static_cast<QRectF>(newGeometry).x(), static_cast<QRectF>(newGeometry).y(), static_cast<QRectF>(newGeometry).width(), static_cast<QRectF>(newGeometry).height()), new QRectF(static_cast<QRectF>(oldGeometry).x(), static_cast<QRectF>(oldGeometry).y(), static_cast<QRectF>(oldGeometry).width(), static_cast<QRectF>(oldGeometry).height())); };
	void hoverEnterEvent(QHoverEvent * event) { callbackQQuickPaintedItem_HoverEnterEvent(this, this->objectName().toUtf8().data(), event); };
	void hoverLeaveEvent(QHoverEvent * event) { callbackQQuickPaintedItem_HoverLeaveEvent(this, this->objectName().toUtf8().data(), event); };
	void hoverMoveEvent(QHoverEvent * event) { callbackQQuickPaintedItem_HoverMoveEvent(this, this->objectName().toUtf8().data(), event); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackQQuickPaintedItem_InputMethodEvent(this, this->objectName().toUtf8().data(), event); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackQQuickPaintedItem_InputMethodQuery(const_cast<MyQQuickPaintedItem*>(this), this->objectName().toUtf8().data(), query)); };
	void keyPressEvent(QKeyEvent * event) { callbackQQuickPaintedItem_KeyPressEvent(this, this->objectName().toUtf8().data(), event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQQuickPaintedItem_KeyReleaseEvent(this, this->objectName().toUtf8().data(), event); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQQuickPaintedItem_MouseDoubleClickEvent(this, this->objectName().toUtf8().data(), event); };
	void mouseMoveEvent(QMouseEvent * event) { callbackQQuickPaintedItem_MouseMoveEvent(this, this->objectName().toUtf8().data(), event); };
	void mousePressEvent(QMouseEvent * event) { callbackQQuickPaintedItem_MousePressEvent(this, this->objectName().toUtf8().data(), event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackQQuickPaintedItem_MouseReleaseEvent(this, this->objectName().toUtf8().data(), event); };
	void mouseUngrabEvent() { callbackQQuickPaintedItem_MouseUngrabEvent(this, this->objectName().toUtf8().data()); };
	void touchEvent(QTouchEvent * event) { callbackQQuickPaintedItem_TouchEvent(this, this->objectName().toUtf8().data(), event); };
	void touchUngrabEvent() { callbackQQuickPaintedItem_TouchUngrabEvent(this, this->objectName().toUtf8().data()); };
	void updatePolish() { callbackQQuickPaintedItem_UpdatePolish(this, this->objectName().toUtf8().data()); };
	void wheelEvent(QWheelEvent * event) { callbackQQuickPaintedItem_WheelEvent(this, this->objectName().toUtf8().data(), event); };
	void timerEvent(QTimerEvent * event) { callbackQQuickPaintedItem_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQQuickPaintedItem_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQQuickPaintedItem_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQQuickPaintedItem_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQQuickPaintedItem_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQQuickPaintedItem_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQQuickPaintedItem_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQQuickPaintedItem_MetaObject(const_cast<MyQQuickPaintedItem*>(this), this->objectName().toUtf8().data())); };
};

double QQuickPaintedItem_ContentsScale(void* ptr)
{
	return static_cast<double>(static_cast<QQuickPaintedItem*>(ptr)->contentsScale());
}

void* QQuickPaintedItem_ContentsSize(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QQuickPaintedItem*>(ptr)->contentsSize()).width(), static_cast<QSize>(static_cast<QQuickPaintedItem*>(ptr)->contentsSize()).height());
}

void* QQuickPaintedItem_FillColor(void* ptr)
{
	return new QColor(static_cast<QQuickPaintedItem*>(ptr)->fillColor());
}

int QQuickPaintedItem_RenderTarget(void* ptr)
{
	return static_cast<QQuickPaintedItem*>(ptr)->renderTarget();
}

void QQuickPaintedItem_SetContentsScale(void* ptr, double vqr)
{
	static_cast<QQuickPaintedItem*>(ptr)->setContentsScale(static_cast<double>(vqr));
}

void QQuickPaintedItem_SetContentsSize(void* ptr, void* vqs)
{
	static_cast<QQuickPaintedItem*>(ptr)->setContentsSize(*static_cast<QSize*>(vqs));
}

void QQuickPaintedItem_SetFillColor(void* ptr, void* vqc)
{
	static_cast<QQuickPaintedItem*>(ptr)->setFillColor(*static_cast<QColor*>(vqc));
}

void QQuickPaintedItem_SetRenderTarget(void* ptr, int target)
{
	static_cast<QQuickPaintedItem*>(ptr)->setRenderTarget(static_cast<QQuickPaintedItem::RenderTarget>(target));
}

void QQuickPaintedItem_SetTextureSize(void* ptr, void* size)
{
	static_cast<QQuickPaintedItem*>(ptr)->setTextureSize(*static_cast<QSize*>(size));
}

void* QQuickPaintedItem_TextureSize(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QQuickPaintedItem*>(ptr)->textureSize()).width(), static_cast<QSize>(static_cast<QQuickPaintedItem*>(ptr)->textureSize()).height());
}

void* QQuickPaintedItem_NewQQuickPaintedItem(void* parent)
{
	return new MyQQuickPaintedItem(static_cast<QQuickItem*>(parent));
}

int QQuickPaintedItem_Antialiasing(void* ptr)
{
	return static_cast<QQuickPaintedItem*>(ptr)->antialiasing();
}

void QQuickPaintedItem_ConnectContentsScaleChanged(void* ptr)
{
	QObject::connect(static_cast<QQuickPaintedItem*>(ptr), static_cast<void (QQuickPaintedItem::*)()>(&QQuickPaintedItem::contentsScaleChanged), static_cast<MyQQuickPaintedItem*>(ptr), static_cast<void (MyQQuickPaintedItem::*)()>(&MyQQuickPaintedItem::Signal_ContentsScaleChanged));
}

void QQuickPaintedItem_DisconnectContentsScaleChanged(void* ptr)
{
	QObject::disconnect(static_cast<QQuickPaintedItem*>(ptr), static_cast<void (QQuickPaintedItem::*)()>(&QQuickPaintedItem::contentsScaleChanged), static_cast<MyQQuickPaintedItem*>(ptr), static_cast<void (MyQQuickPaintedItem::*)()>(&MyQQuickPaintedItem::Signal_ContentsScaleChanged));
}

void QQuickPaintedItem_ContentsScaleChanged(void* ptr)
{
	static_cast<QQuickPaintedItem*>(ptr)->contentsScaleChanged();
}

void QQuickPaintedItem_ConnectContentsSizeChanged(void* ptr)
{
	QObject::connect(static_cast<QQuickPaintedItem*>(ptr), static_cast<void (QQuickPaintedItem::*)()>(&QQuickPaintedItem::contentsSizeChanged), static_cast<MyQQuickPaintedItem*>(ptr), static_cast<void (MyQQuickPaintedItem::*)()>(&MyQQuickPaintedItem::Signal_ContentsSizeChanged));
}

void QQuickPaintedItem_DisconnectContentsSizeChanged(void* ptr)
{
	QObject::disconnect(static_cast<QQuickPaintedItem*>(ptr), static_cast<void (QQuickPaintedItem::*)()>(&QQuickPaintedItem::contentsSizeChanged), static_cast<MyQQuickPaintedItem*>(ptr), static_cast<void (MyQQuickPaintedItem::*)()>(&MyQQuickPaintedItem::Signal_ContentsSizeChanged));
}

void QQuickPaintedItem_ContentsSizeChanged(void* ptr)
{
	static_cast<QQuickPaintedItem*>(ptr)->contentsSizeChanged();
}

void QQuickPaintedItem_ConnectFillColorChanged(void* ptr)
{
	QObject::connect(static_cast<QQuickPaintedItem*>(ptr), static_cast<void (QQuickPaintedItem::*)()>(&QQuickPaintedItem::fillColorChanged), static_cast<MyQQuickPaintedItem*>(ptr), static_cast<void (MyQQuickPaintedItem::*)()>(&MyQQuickPaintedItem::Signal_FillColorChanged));
}

void QQuickPaintedItem_DisconnectFillColorChanged(void* ptr)
{
	QObject::disconnect(static_cast<QQuickPaintedItem*>(ptr), static_cast<void (QQuickPaintedItem::*)()>(&QQuickPaintedItem::fillColorChanged), static_cast<MyQQuickPaintedItem*>(ptr), static_cast<void (MyQQuickPaintedItem::*)()>(&MyQQuickPaintedItem::Signal_FillColorChanged));
}

void QQuickPaintedItem_FillColorChanged(void* ptr)
{
	static_cast<QQuickPaintedItem*>(ptr)->fillColorChanged();
}

int QQuickPaintedItem_IsTextureProvider(void* ptr)
{
	return static_cast<QQuickPaintedItem*>(ptr)->isTextureProvider();
}

int QQuickPaintedItem_IsTextureProviderDefault(void* ptr)
{
	return static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::isTextureProvider();
}

int QQuickPaintedItem_Mipmap(void* ptr)
{
	return static_cast<QQuickPaintedItem*>(ptr)->mipmap();
}

int QQuickPaintedItem_OpaquePainting(void* ptr)
{
	return static_cast<QQuickPaintedItem*>(ptr)->opaquePainting();
}

void QQuickPaintedItem_Paint(void* ptr, void* painter)
{
	static_cast<QQuickPaintedItem*>(ptr)->paint(static_cast<QPainter*>(painter));
}

int QQuickPaintedItem_PerformanceHints(void* ptr)
{
	return static_cast<QQuickPaintedItem*>(ptr)->performanceHints();
}

void QQuickPaintedItem_ReleaseResources(void* ptr)
{
	static_cast<QQuickPaintedItem*>(ptr)->releaseResources();
}

void QQuickPaintedItem_ReleaseResourcesDefault(void* ptr)
{
	static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::releaseResources();
}

void QQuickPaintedItem_ConnectRenderTargetChanged(void* ptr)
{
	QObject::connect(static_cast<QQuickPaintedItem*>(ptr), static_cast<void (QQuickPaintedItem::*)()>(&QQuickPaintedItem::renderTargetChanged), static_cast<MyQQuickPaintedItem*>(ptr), static_cast<void (MyQQuickPaintedItem::*)()>(&MyQQuickPaintedItem::Signal_RenderTargetChanged));
}

void QQuickPaintedItem_DisconnectRenderTargetChanged(void* ptr)
{
	QObject::disconnect(static_cast<QQuickPaintedItem*>(ptr), static_cast<void (QQuickPaintedItem::*)()>(&QQuickPaintedItem::renderTargetChanged), static_cast<MyQQuickPaintedItem*>(ptr), static_cast<void (MyQQuickPaintedItem::*)()>(&MyQQuickPaintedItem::Signal_RenderTargetChanged));
}

void QQuickPaintedItem_RenderTargetChanged(void* ptr)
{
	static_cast<QQuickPaintedItem*>(ptr)->renderTargetChanged();
}

void QQuickPaintedItem_SetAntialiasing(void* ptr, int enable)
{
	static_cast<QQuickPaintedItem*>(ptr)->setAntialiasing(enable != 0);
}

void QQuickPaintedItem_SetMipmap(void* ptr, int enable)
{
	static_cast<QQuickPaintedItem*>(ptr)->setMipmap(enable != 0);
}

void QQuickPaintedItem_SetOpaquePainting(void* ptr, int opaque)
{
	static_cast<QQuickPaintedItem*>(ptr)->setOpaquePainting(opaque != 0);
}

void QQuickPaintedItem_SetPerformanceHint(void* ptr, int hint, int enabled)
{
	static_cast<QQuickPaintedItem*>(ptr)->setPerformanceHint(static_cast<QQuickPaintedItem::PerformanceHint>(hint), enabled != 0);
}

void QQuickPaintedItem_SetPerformanceHints(void* ptr, int hints)
{
	static_cast<QQuickPaintedItem*>(ptr)->setPerformanceHints(static_cast<QQuickPaintedItem::PerformanceHint>(hints));
}

void* QQuickPaintedItem_TextureProvider(void* ptr)
{
	return static_cast<QQuickPaintedItem*>(ptr)->textureProvider();
}

void* QQuickPaintedItem_TextureProviderDefault(void* ptr)
{
	return static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::textureProvider();
}

void QQuickPaintedItem_ConnectTextureSizeChanged(void* ptr)
{
	QObject::connect(static_cast<QQuickPaintedItem*>(ptr), static_cast<void (QQuickPaintedItem::*)()>(&QQuickPaintedItem::textureSizeChanged), static_cast<MyQQuickPaintedItem*>(ptr), static_cast<void (MyQQuickPaintedItem::*)()>(&MyQQuickPaintedItem::Signal_TextureSizeChanged));
}

void QQuickPaintedItem_DisconnectTextureSizeChanged(void* ptr)
{
	QObject::disconnect(static_cast<QQuickPaintedItem*>(ptr), static_cast<void (QQuickPaintedItem::*)()>(&QQuickPaintedItem::textureSizeChanged), static_cast<MyQQuickPaintedItem*>(ptr), static_cast<void (MyQQuickPaintedItem::*)()>(&MyQQuickPaintedItem::Signal_TextureSizeChanged));
}

void QQuickPaintedItem_TextureSizeChanged(void* ptr)
{
	static_cast<QQuickPaintedItem*>(ptr)->textureSizeChanged();
}

void QQuickPaintedItem_Update(void* ptr, void* rect)
{
	static_cast<QQuickPaintedItem*>(ptr)->update(*static_cast<QRect*>(rect));
}

void QQuickPaintedItem_DestroyQQuickPaintedItem(void* ptr)
{
	static_cast<QQuickPaintedItem*>(ptr)->~QQuickPaintedItem();
}

int QQuickPaintedItem_ChildMouseEventFilter(void* ptr, void* item, void* event)
{
	return static_cast<QQuickPaintedItem*>(ptr)->childMouseEventFilter(static_cast<QQuickItem*>(item), static_cast<QEvent*>(event));
}

int QQuickPaintedItem_ChildMouseEventFilterDefault(void* ptr, void* item, void* event)
{
	return static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::childMouseEventFilter(static_cast<QQuickItem*>(item), static_cast<QEvent*>(event));
}

int QQuickPaintedItem_Contains(void* ptr, void* point)
{
	return static_cast<QQuickPaintedItem*>(ptr)->contains(*static_cast<QPointF*>(point));
}

int QQuickPaintedItem_ContainsDefault(void* ptr, void* point)
{
	return static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::contains(*static_cast<QPointF*>(point));
}

void QQuickPaintedItem_DragEnterEvent(void* ptr, void* event)
{
	static_cast<QQuickPaintedItem*>(ptr)->dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QQuickPaintedItem_DragEnterEventDefault(void* ptr, void* event)
{
	static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QQuickPaintedItem_DragLeaveEvent(void* ptr, void* event)
{
	static_cast<QQuickPaintedItem*>(ptr)->dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QQuickPaintedItem_DragLeaveEventDefault(void* ptr, void* event)
{
	static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QQuickPaintedItem_DragMoveEvent(void* ptr, void* event)
{
	static_cast<QQuickPaintedItem*>(ptr)->dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QQuickPaintedItem_DragMoveEventDefault(void* ptr, void* event)
{
	static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QQuickPaintedItem_DropEvent(void* ptr, void* event)
{
	static_cast<QQuickPaintedItem*>(ptr)->dropEvent(static_cast<QDropEvent*>(event));
}

void QQuickPaintedItem_DropEventDefault(void* ptr, void* event)
{
	static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::dropEvent(static_cast<QDropEvent*>(event));
}

void QQuickPaintedItem_FocusInEvent(void* ptr, void* event)
{
	static_cast<QQuickPaintedItem*>(ptr)->focusInEvent(static_cast<QFocusEvent*>(event));
}

void QQuickPaintedItem_FocusInEventDefault(void* ptr, void* event)
{
	static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::focusInEvent(static_cast<QFocusEvent*>(event));
}

void QQuickPaintedItem_FocusOutEvent(void* ptr, void* event)
{
	static_cast<QQuickPaintedItem*>(ptr)->focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QQuickPaintedItem_FocusOutEventDefault(void* ptr, void* event)
{
	static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QQuickPaintedItem_GeometryChanged(void* ptr, void* newGeometry, void* oldGeometry)
{
	static_cast<QQuickPaintedItem*>(ptr)->geometryChanged(*static_cast<QRectF*>(newGeometry), *static_cast<QRectF*>(oldGeometry));
}

void QQuickPaintedItem_GeometryChangedDefault(void* ptr, void* newGeometry, void* oldGeometry)
{
	static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::geometryChanged(*static_cast<QRectF*>(newGeometry), *static_cast<QRectF*>(oldGeometry));
}

void QQuickPaintedItem_HoverEnterEvent(void* ptr, void* event)
{
	static_cast<QQuickPaintedItem*>(ptr)->hoverEnterEvent(static_cast<QHoverEvent*>(event));
}

void QQuickPaintedItem_HoverEnterEventDefault(void* ptr, void* event)
{
	static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::hoverEnterEvent(static_cast<QHoverEvent*>(event));
}

void QQuickPaintedItem_HoverLeaveEvent(void* ptr, void* event)
{
	static_cast<QQuickPaintedItem*>(ptr)->hoverLeaveEvent(static_cast<QHoverEvent*>(event));
}

void QQuickPaintedItem_HoverLeaveEventDefault(void* ptr, void* event)
{
	static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::hoverLeaveEvent(static_cast<QHoverEvent*>(event));
}

void QQuickPaintedItem_HoverMoveEvent(void* ptr, void* event)
{
	static_cast<QQuickPaintedItem*>(ptr)->hoverMoveEvent(static_cast<QHoverEvent*>(event));
}

void QQuickPaintedItem_HoverMoveEventDefault(void* ptr, void* event)
{
	static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::hoverMoveEvent(static_cast<QHoverEvent*>(event));
}

void QQuickPaintedItem_InputMethodEvent(void* ptr, void* event)
{
	static_cast<QQuickPaintedItem*>(ptr)->inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void QQuickPaintedItem_InputMethodEventDefault(void* ptr, void* event)
{
	static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void* QQuickPaintedItem_InputMethodQuery(void* ptr, int query)
{
	return new QVariant(static_cast<QQuickPaintedItem*>(ptr)->inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

void* QQuickPaintedItem_InputMethodQueryDefault(void* ptr, int query)
{
	return new QVariant(static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

void QQuickPaintedItem_KeyPressEvent(void* ptr, void* event)
{
	static_cast<QQuickPaintedItem*>(ptr)->keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QQuickPaintedItem_KeyPressEventDefault(void* ptr, void* event)
{
	static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QQuickPaintedItem_KeyReleaseEvent(void* ptr, void* event)
{
	static_cast<QQuickPaintedItem*>(ptr)->keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QQuickPaintedItem_KeyReleaseEventDefault(void* ptr, void* event)
{
	static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QQuickPaintedItem_MouseDoubleClickEvent(void* ptr, void* event)
{
	static_cast<QQuickPaintedItem*>(ptr)->mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QQuickPaintedItem_MouseDoubleClickEventDefault(void* ptr, void* event)
{
	static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QQuickPaintedItem_MouseMoveEvent(void* ptr, void* event)
{
	static_cast<QQuickPaintedItem*>(ptr)->mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QQuickPaintedItem_MouseMoveEventDefault(void* ptr, void* event)
{
	static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QQuickPaintedItem_MousePressEvent(void* ptr, void* event)
{
	static_cast<QQuickPaintedItem*>(ptr)->mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QQuickPaintedItem_MousePressEventDefault(void* ptr, void* event)
{
	static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QQuickPaintedItem_MouseReleaseEvent(void* ptr, void* event)
{
	static_cast<QQuickPaintedItem*>(ptr)->mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void QQuickPaintedItem_MouseReleaseEventDefault(void* ptr, void* event)
{
	static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void QQuickPaintedItem_MouseUngrabEvent(void* ptr)
{
	static_cast<QQuickPaintedItem*>(ptr)->mouseUngrabEvent();
}

void QQuickPaintedItem_MouseUngrabEventDefault(void* ptr)
{
	static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::mouseUngrabEvent();
}

void QQuickPaintedItem_TouchEvent(void* ptr, void* event)
{
	static_cast<QQuickPaintedItem*>(ptr)->touchEvent(static_cast<QTouchEvent*>(event));
}

void QQuickPaintedItem_TouchEventDefault(void* ptr, void* event)
{
	static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::touchEvent(static_cast<QTouchEvent*>(event));
}

void QQuickPaintedItem_TouchUngrabEvent(void* ptr)
{
	static_cast<QQuickPaintedItem*>(ptr)->touchUngrabEvent();
}

void QQuickPaintedItem_TouchUngrabEventDefault(void* ptr)
{
	static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::touchUngrabEvent();
}

void QQuickPaintedItem_UpdatePolish(void* ptr)
{
	static_cast<QQuickPaintedItem*>(ptr)->updatePolish();
}

void QQuickPaintedItem_UpdatePolishDefault(void* ptr)
{
	static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::updatePolish();
}

void QQuickPaintedItem_WheelEvent(void* ptr, void* event)
{
	static_cast<QQuickPaintedItem*>(ptr)->wheelEvent(static_cast<QWheelEvent*>(event));
}

void QQuickPaintedItem_WheelEventDefault(void* ptr, void* event)
{
	static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::wheelEvent(static_cast<QWheelEvent*>(event));
}

void QQuickPaintedItem_TimerEvent(void* ptr, void* event)
{
	static_cast<QQuickPaintedItem*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QQuickPaintedItem_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::timerEvent(static_cast<QTimerEvent*>(event));
}

void QQuickPaintedItem_ChildEvent(void* ptr, void* event)
{
	static_cast<QQuickPaintedItem*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QQuickPaintedItem_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::childEvent(static_cast<QChildEvent*>(event));
}

void QQuickPaintedItem_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QQuickPaintedItem*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQuickPaintedItem_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQuickPaintedItem_CustomEvent(void* ptr, void* event)
{
	static_cast<QQuickPaintedItem*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QQuickPaintedItem_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::customEvent(static_cast<QEvent*>(event));
}

void QQuickPaintedItem_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QQuickPaintedItem*>(ptr), "deleteLater");
}

void QQuickPaintedItem_DeleteLaterDefault(void* ptr)
{
	static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::deleteLater();
}

void QQuickPaintedItem_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QQuickPaintedItem*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQuickPaintedItem_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QQuickPaintedItem_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QQuickPaintedItem*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QQuickPaintedItem_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QQuickPaintedItem_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QQuickPaintedItem*>(ptr)->metaObject());
}

void* QQuickPaintedItem_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::metaObject());
}

class MyQQuickRenderControl: public QQuickRenderControl
{
public:
	MyQQuickRenderControl(QObject *parent) : QQuickRenderControl(parent) {};
	void Signal_RenderRequested() { callbackQQuickRenderControl_RenderRequested(this, this->objectName().toUtf8().data()); };
	QWindow * renderWindow(QPoint * offset) { return static_cast<QWindow*>(callbackQQuickRenderControl_RenderWindow(this, this->objectName().toUtf8().data(), offset)); };
	void Signal_SceneChanged() { callbackQQuickRenderControl_SceneChanged(this, this->objectName().toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQQuickRenderControl_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQQuickRenderControl_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQQuickRenderControl_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQQuickRenderControl_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQQuickRenderControl_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQQuickRenderControl_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQQuickRenderControl_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQQuickRenderControl_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQQuickRenderControl_MetaObject(const_cast<MyQQuickRenderControl*>(this), this->objectName().toUtf8().data())); };
};

void* QQuickRenderControl_NewQQuickRenderControl(void* parent)
{
	return new MyQQuickRenderControl(static_cast<QObject*>(parent));
}

void* QQuickRenderControl_Grab(void* ptr)
{
	return new QImage(static_cast<QQuickRenderControl*>(ptr)->grab());
}

void QQuickRenderControl_Initialize(void* ptr, void* gl)
{
	static_cast<QQuickRenderControl*>(ptr)->initialize(static_cast<QOpenGLContext*>(gl));
}

void QQuickRenderControl_Invalidate(void* ptr)
{
	static_cast<QQuickRenderControl*>(ptr)->invalidate();
}

void QQuickRenderControl_PolishItems(void* ptr)
{
	static_cast<QQuickRenderControl*>(ptr)->polishItems();
}

void QQuickRenderControl_PrepareThread(void* ptr, void* targetThread)
{
	static_cast<QQuickRenderControl*>(ptr)->prepareThread(static_cast<QThread*>(targetThread));
}

void QQuickRenderControl_Render(void* ptr)
{
	static_cast<QQuickRenderControl*>(ptr)->render();
}

void QQuickRenderControl_ConnectRenderRequested(void* ptr)
{
	QObject::connect(static_cast<QQuickRenderControl*>(ptr), static_cast<void (QQuickRenderControl::*)()>(&QQuickRenderControl::renderRequested), static_cast<MyQQuickRenderControl*>(ptr), static_cast<void (MyQQuickRenderControl::*)()>(&MyQQuickRenderControl::Signal_RenderRequested));
}

void QQuickRenderControl_DisconnectRenderRequested(void* ptr)
{
	QObject::disconnect(static_cast<QQuickRenderControl*>(ptr), static_cast<void (QQuickRenderControl::*)()>(&QQuickRenderControl::renderRequested), static_cast<MyQQuickRenderControl*>(ptr), static_cast<void (MyQQuickRenderControl::*)()>(&MyQQuickRenderControl::Signal_RenderRequested));
}

void QQuickRenderControl_RenderRequested(void* ptr)
{
	static_cast<QQuickRenderControl*>(ptr)->renderRequested();
}

void* QQuickRenderControl_RenderWindow(void* ptr, void* offset)
{
	return static_cast<QQuickRenderControl*>(ptr)->renderWindow(static_cast<QPoint*>(offset));
}

void* QQuickRenderControl_RenderWindowDefault(void* ptr, void* offset)
{
	return static_cast<QQuickRenderControl*>(ptr)->QQuickRenderControl::renderWindow(static_cast<QPoint*>(offset));
}

void* QQuickRenderControl_QQuickRenderControl_RenderWindowFor(void* win, void* offset)
{
	return QQuickRenderControl::renderWindowFor(static_cast<QQuickWindow*>(win), static_cast<QPoint*>(offset));
}

void QQuickRenderControl_ConnectSceneChanged(void* ptr)
{
	QObject::connect(static_cast<QQuickRenderControl*>(ptr), static_cast<void (QQuickRenderControl::*)()>(&QQuickRenderControl::sceneChanged), static_cast<MyQQuickRenderControl*>(ptr), static_cast<void (MyQQuickRenderControl::*)()>(&MyQQuickRenderControl::Signal_SceneChanged));
}

void QQuickRenderControl_DisconnectSceneChanged(void* ptr)
{
	QObject::disconnect(static_cast<QQuickRenderControl*>(ptr), static_cast<void (QQuickRenderControl::*)()>(&QQuickRenderControl::sceneChanged), static_cast<MyQQuickRenderControl*>(ptr), static_cast<void (MyQQuickRenderControl::*)()>(&MyQQuickRenderControl::Signal_SceneChanged));
}

void QQuickRenderControl_SceneChanged(void* ptr)
{
	static_cast<QQuickRenderControl*>(ptr)->sceneChanged();
}

int QQuickRenderControl_Sync(void* ptr)
{
	return static_cast<QQuickRenderControl*>(ptr)->sync();
}

void QQuickRenderControl_DestroyQQuickRenderControl(void* ptr)
{
	static_cast<QQuickRenderControl*>(ptr)->~QQuickRenderControl();
}

void QQuickRenderControl_TimerEvent(void* ptr, void* event)
{
	static_cast<QQuickRenderControl*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QQuickRenderControl_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QQuickRenderControl*>(ptr)->QQuickRenderControl::timerEvent(static_cast<QTimerEvent*>(event));
}

void QQuickRenderControl_ChildEvent(void* ptr, void* event)
{
	static_cast<QQuickRenderControl*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QQuickRenderControl_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QQuickRenderControl*>(ptr)->QQuickRenderControl::childEvent(static_cast<QChildEvent*>(event));
}

void QQuickRenderControl_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QQuickRenderControl*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQuickRenderControl_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QQuickRenderControl*>(ptr)->QQuickRenderControl::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQuickRenderControl_CustomEvent(void* ptr, void* event)
{
	static_cast<QQuickRenderControl*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QQuickRenderControl_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QQuickRenderControl*>(ptr)->QQuickRenderControl::customEvent(static_cast<QEvent*>(event));
}

void QQuickRenderControl_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QQuickRenderControl*>(ptr), "deleteLater");
}

void QQuickRenderControl_DeleteLaterDefault(void* ptr)
{
	static_cast<QQuickRenderControl*>(ptr)->QQuickRenderControl::deleteLater();
}

void QQuickRenderControl_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QQuickRenderControl*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQuickRenderControl_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QQuickRenderControl*>(ptr)->QQuickRenderControl::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QQuickRenderControl_Event(void* ptr, void* e)
{
	return static_cast<QQuickRenderControl*>(ptr)->event(static_cast<QEvent*>(e));
}

int QQuickRenderControl_EventDefault(void* ptr, void* e)
{
	return static_cast<QQuickRenderControl*>(ptr)->QQuickRenderControl::event(static_cast<QEvent*>(e));
}

int QQuickRenderControl_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QQuickRenderControl*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QQuickRenderControl_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QQuickRenderControl*>(ptr)->QQuickRenderControl::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QQuickRenderControl_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QQuickRenderControl*>(ptr)->metaObject());
}

void* QQuickRenderControl_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QQuickRenderControl*>(ptr)->QQuickRenderControl::metaObject());
}

void* QQuickTextDocument_NewQQuickTextDocument(void* parent)
{
	return new QQuickTextDocument(static_cast<QQuickItem*>(parent));
}

void* QQuickTextDocument_TextDocument(void* ptr)
{
	return static_cast<QQuickTextDocument*>(ptr)->textDocument();
}

void QQuickTextDocument_TimerEvent(void* ptr, void* event)
{
	static_cast<QQuickTextDocument*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QQuickTextDocument_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QQuickTextDocument*>(ptr)->QQuickTextDocument::timerEvent(static_cast<QTimerEvent*>(event));
}

void QQuickTextDocument_ChildEvent(void* ptr, void* event)
{
	static_cast<QQuickTextDocument*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QQuickTextDocument_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QQuickTextDocument*>(ptr)->QQuickTextDocument::childEvent(static_cast<QChildEvent*>(event));
}

void QQuickTextDocument_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QQuickTextDocument*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQuickTextDocument_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QQuickTextDocument*>(ptr)->QQuickTextDocument::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQuickTextDocument_CustomEvent(void* ptr, void* event)
{
	static_cast<QQuickTextDocument*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QQuickTextDocument_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QQuickTextDocument*>(ptr)->QQuickTextDocument::customEvent(static_cast<QEvent*>(event));
}

void QQuickTextDocument_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QQuickTextDocument*>(ptr), "deleteLater");
}

void QQuickTextDocument_DeleteLaterDefault(void* ptr)
{
	static_cast<QQuickTextDocument*>(ptr)->QQuickTextDocument::deleteLater();
}

void QQuickTextDocument_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QQuickTextDocument*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQuickTextDocument_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QQuickTextDocument*>(ptr)->QQuickTextDocument::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QQuickTextDocument_Event(void* ptr, void* e)
{
	return static_cast<QQuickTextDocument*>(ptr)->event(static_cast<QEvent*>(e));
}

int QQuickTextDocument_EventDefault(void* ptr, void* e)
{
	return static_cast<QQuickTextDocument*>(ptr)->QQuickTextDocument::event(static_cast<QEvent*>(e));
}

int QQuickTextDocument_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QQuickTextDocument*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QQuickTextDocument_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QQuickTextDocument*>(ptr)->QQuickTextDocument::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QQuickTextDocument_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QQuickTextDocument*>(ptr)->metaObject());
}

void* QQuickTextDocument_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QQuickTextDocument*>(ptr)->QQuickTextDocument::metaObject());
}

class MyQQuickTextureFactory: public QQuickTextureFactory
{
public:
	MyQQuickTextureFactory() : QQuickTextureFactory() {};
	QImage image() const { return *static_cast<QImage*>(callbackQQuickTextureFactory_Image(const_cast<MyQQuickTextureFactory*>(this), this->objectName().toUtf8().data())); };
	QSGTexture * createTexture(QQuickWindow * window) const { return static_cast<QSGTexture*>(callbackQQuickTextureFactory_CreateTexture(const_cast<MyQQuickTextureFactory*>(this), this->objectName().toUtf8().data(), window)); };
	int textureByteCount() const { return callbackQQuickTextureFactory_TextureByteCount(const_cast<MyQQuickTextureFactory*>(this), this->objectName().toUtf8().data()); };
	QSize textureSize() const { return *static_cast<QSize*>(callbackQQuickTextureFactory_TextureSize(const_cast<MyQQuickTextureFactory*>(this), this->objectName().toUtf8().data())); };
	void timerEvent(QTimerEvent * event) { callbackQQuickTextureFactory_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQQuickTextureFactory_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQQuickTextureFactory_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQQuickTextureFactory_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQQuickTextureFactory_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQQuickTextureFactory_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQQuickTextureFactory_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQQuickTextureFactory_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQQuickTextureFactory_MetaObject(const_cast<MyQQuickTextureFactory*>(this), this->objectName().toUtf8().data())); };
};

void* QQuickTextureFactory_Image(void* ptr)
{
	return new QImage(static_cast<QQuickTextureFactory*>(ptr)->image());
}

void* QQuickTextureFactory_ImageDefault(void* ptr)
{
	return new QImage(static_cast<QQuickTextureFactory*>(ptr)->QQuickTextureFactory::image());
}

void* QQuickTextureFactory_NewQQuickTextureFactory()
{
	return new MyQQuickTextureFactory();
}

void* QQuickTextureFactory_CreateTexture(void* ptr, void* window)
{
	return static_cast<QQuickTextureFactory*>(ptr)->createTexture(static_cast<QQuickWindow*>(window));
}

int QQuickTextureFactory_TextureByteCount(void* ptr)
{
	return static_cast<QQuickTextureFactory*>(ptr)->textureByteCount();
}

void* QQuickTextureFactory_QQuickTextureFactory_TextureFactoryForImage(void* image)
{
	return QQuickTextureFactory::textureFactoryForImage(*static_cast<QImage*>(image));
}

void* QQuickTextureFactory_TextureSize(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QQuickTextureFactory*>(ptr)->textureSize()).width(), static_cast<QSize>(static_cast<QQuickTextureFactory*>(ptr)->textureSize()).height());
}

void QQuickTextureFactory_DestroyQQuickTextureFactory(void* ptr)
{
	static_cast<QQuickTextureFactory*>(ptr)->~QQuickTextureFactory();
}

void QQuickTextureFactory_TimerEvent(void* ptr, void* event)
{
	static_cast<QQuickTextureFactory*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QQuickTextureFactory_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QQuickTextureFactory*>(ptr)->QQuickTextureFactory::timerEvent(static_cast<QTimerEvent*>(event));
}

void QQuickTextureFactory_ChildEvent(void* ptr, void* event)
{
	static_cast<QQuickTextureFactory*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QQuickTextureFactory_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QQuickTextureFactory*>(ptr)->QQuickTextureFactory::childEvent(static_cast<QChildEvent*>(event));
}

void QQuickTextureFactory_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QQuickTextureFactory*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQuickTextureFactory_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QQuickTextureFactory*>(ptr)->QQuickTextureFactory::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQuickTextureFactory_CustomEvent(void* ptr, void* event)
{
	static_cast<QQuickTextureFactory*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QQuickTextureFactory_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QQuickTextureFactory*>(ptr)->QQuickTextureFactory::customEvent(static_cast<QEvent*>(event));
}

void QQuickTextureFactory_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QQuickTextureFactory*>(ptr), "deleteLater");
}

void QQuickTextureFactory_DeleteLaterDefault(void* ptr)
{
	static_cast<QQuickTextureFactory*>(ptr)->QQuickTextureFactory::deleteLater();
}

void QQuickTextureFactory_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QQuickTextureFactory*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQuickTextureFactory_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QQuickTextureFactory*>(ptr)->QQuickTextureFactory::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QQuickTextureFactory_Event(void* ptr, void* e)
{
	return static_cast<QQuickTextureFactory*>(ptr)->event(static_cast<QEvent*>(e));
}

int QQuickTextureFactory_EventDefault(void* ptr, void* e)
{
	return static_cast<QQuickTextureFactory*>(ptr)->QQuickTextureFactory::event(static_cast<QEvent*>(e));
}

int QQuickTextureFactory_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QQuickTextureFactory*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QQuickTextureFactory_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QQuickTextureFactory*>(ptr)->QQuickTextureFactory::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QQuickTextureFactory_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QQuickTextureFactory*>(ptr)->metaObject());
}

void* QQuickTextureFactory_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QQuickTextureFactory*>(ptr)->QQuickTextureFactory::metaObject());
}

class MyQQuickView: public QQuickView
{
public:
	MyQQuickView(QQmlEngine *engine, QWindow *parent) : QQuickView(engine, parent) {};
	MyQQuickView(QWindow *parent) : QQuickView(parent) {};
	MyQQuickView(const QUrl &source, QWindow *parent) : QQuickView(source, parent) {};
	void setSource(const QUrl & url) { callbackQQuickView_SetSource(this, this->objectName().toUtf8().data(), new QUrl(url)); };
	void Signal_StatusChanged(QQuickView::Status status) { callbackQQuickView_StatusChanged(this, this->objectName().toUtf8().data(), status); };
	void releaseResources() { callbackQQuickView_ReleaseResources(this, this->objectName().toUtf8().data()); };
	void update() { callbackQQuickView_Update(this, this->objectName().toUtf8().data()); };
	void setHeight(int arg) { callbackQQuickView_SetHeight(this, this->objectName().toUtf8().data(), arg); };
	void setMaximumHeight(int h) { callbackQQuickView_SetMaximumHeight(this, this->objectName().toUtf8().data(), h); };
	void setMaximumWidth(int w) { callbackQQuickView_SetMaximumWidth(this, this->objectName().toUtf8().data(), w); };
	void setMinimumHeight(int h) { callbackQQuickView_SetMinimumHeight(this, this->objectName().toUtf8().data(), h); };
	void setMinimumWidth(int w) { callbackQQuickView_SetMinimumWidth(this, this->objectName().toUtf8().data(), w); };
	void setTitle(const QString & vqs) { callbackQQuickView_SetTitle(this, this->objectName().toUtf8().data(), vqs.toUtf8().data()); };
	void setVisible(bool visible) { callbackQQuickView_SetVisible(this, this->objectName().toUtf8().data(), visible); };
	void setWidth(int arg) { callbackQQuickView_SetWidth(this, this->objectName().toUtf8().data(), arg); };
	void setX(int arg) { callbackQQuickView_SetX(this, this->objectName().toUtf8().data(), arg); };
	void setY(int arg) { callbackQQuickView_SetY(this, this->objectName().toUtf8().data(), arg); };
	void alert(int msec) { callbackQQuickView_Alert(this, this->objectName().toUtf8().data(), msec); };
	bool close() { return callbackQQuickView_Close(this, this->objectName().toUtf8().data()) != 0; };
	QObject * focusObject() const { return static_cast<QObject*>(callbackQQuickView_FocusObject(const_cast<MyQQuickView*>(this), this->objectName().toUtf8().data())); };
	QSurfaceFormat format() const { return *static_cast<QSurfaceFormat*>(callbackQQuickView_Format(const_cast<MyQQuickView*>(this), this->objectName().toUtf8().data())); };
	void hide() { callbackQQuickView_Hide(this, this->objectName().toUtf8().data()); };
	void lower() { callbackQQuickView_Lower(this, this->objectName().toUtf8().data()); };
	void moveEvent(QMoveEvent * ev) { callbackQQuickView_MoveEvent(this, this->objectName().toUtf8().data(), ev); };
	bool nativeEvent(const QByteArray & eventType, void * message, long * result) { return callbackQQuickView_NativeEvent(this, this->objectName().toUtf8().data(), QString(eventType).toUtf8().data(), message, *result) != 0; };
	void raise() { callbackQQuickView_Raise(this, this->objectName().toUtf8().data()); };
	void requestActivate() { callbackQQuickView_RequestActivate(this, this->objectName().toUtf8().data()); };
	void requestUpdate() { callbackQQuickView_RequestUpdate(this, this->objectName().toUtf8().data()); };
	void show() { callbackQQuickView_Show(this, this->objectName().toUtf8().data()); };
	void showFullScreen() { callbackQQuickView_ShowFullScreen(this, this->objectName().toUtf8().data()); };
	void showMaximized() { callbackQQuickView_ShowMaximized(this, this->objectName().toUtf8().data()); };
	void showMinimized() { callbackQQuickView_ShowMinimized(this, this->objectName().toUtf8().data()); };
	void showNormal() { callbackQQuickView_ShowNormal(this, this->objectName().toUtf8().data()); };
	QSize size() const { return *static_cast<QSize*>(callbackQQuickView_Size(const_cast<MyQQuickView*>(this), this->objectName().toUtf8().data())); };
	SurfaceType surfaceType() const { return static_cast<QSurface::SurfaceType>(callbackQQuickView_SurfaceType(const_cast<MyQQuickView*>(this), this->objectName().toUtf8().data())); };
	void tabletEvent(QTabletEvent * ev) { callbackQQuickView_TabletEvent(this, this->objectName().toUtf8().data(), ev); };
	void touchEvent(QTouchEvent * ev) { callbackQQuickView_TouchEvent(this, this->objectName().toUtf8().data(), ev); };
	void timerEvent(QTimerEvent * event) { callbackQQuickView_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQQuickView_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQQuickView_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQQuickView_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQQuickView_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQQuickView_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQQuickView_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQQuickView_MetaObject(const_cast<MyQQuickView*>(this), this->objectName().toUtf8().data())); };
};

int QQuickView_ResizeMode(void* ptr)
{
	return static_cast<QQuickView*>(ptr)->resizeMode();
}

void QQuickView_SetResizeMode(void* ptr, int vre)
{
	static_cast<QQuickView*>(ptr)->setResizeMode(static_cast<QQuickView::ResizeMode>(vre));
}

int QQuickView_Status(void* ptr)
{
	return static_cast<QQuickView*>(ptr)->status();
}

void* QQuickView_NewQQuickView2(void* engine, void* parent)
{
	return new MyQQuickView(static_cast<QQmlEngine*>(engine), static_cast<QWindow*>(parent));
}

void* QQuickView_NewQQuickView(void* parent)
{
	return new MyQQuickView(static_cast<QWindow*>(parent));
}

void* QQuickView_NewQQuickView3(void* source, void* parent)
{
	return new MyQQuickView(*static_cast<QUrl*>(source), static_cast<QWindow*>(parent));
}

void* QQuickView_Engine(void* ptr)
{
	return static_cast<QQuickView*>(ptr)->engine();
}

void* QQuickView_InitialSize(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QQuickView*>(ptr)->initialSize()).width(), static_cast<QSize>(static_cast<QQuickView*>(ptr)->initialSize()).height());
}

void QQuickView_KeyPressEvent(void* ptr, void* e)
{
	static_cast<QQuickView*>(ptr)->keyPressEvent(static_cast<QKeyEvent*>(e));
}

void QQuickView_KeyReleaseEvent(void* ptr, void* e)
{
	static_cast<QQuickView*>(ptr)->keyReleaseEvent(static_cast<QKeyEvent*>(e));
}

void QQuickView_MouseMoveEvent(void* ptr, void* e)
{
	static_cast<QQuickView*>(ptr)->mouseMoveEvent(static_cast<QMouseEvent*>(e));
}

void QQuickView_MousePressEvent(void* ptr, void* e)
{
	static_cast<QQuickView*>(ptr)->mousePressEvent(static_cast<QMouseEvent*>(e));
}

void QQuickView_MouseReleaseEvent(void* ptr, void* e)
{
	static_cast<QQuickView*>(ptr)->mouseReleaseEvent(static_cast<QMouseEvent*>(e));
}

void* QQuickView_RootContext(void* ptr)
{
	return static_cast<QQuickView*>(ptr)->rootContext();
}

void* QQuickView_RootObject(void* ptr)
{
	return static_cast<QQuickView*>(ptr)->rootObject();
}

void QQuickView_SetSource(void* ptr, void* url)
{
	QMetaObject::invokeMethod(static_cast<QQuickView*>(ptr), "setSource", Q_ARG(QUrl, *static_cast<QUrl*>(url)));
}

void* QQuickView_Source(void* ptr)
{
	return new QUrl(static_cast<QQuickView*>(ptr)->source());
}

void QQuickView_ConnectStatusChanged(void* ptr)
{
	QObject::connect(static_cast<QQuickView*>(ptr), static_cast<void (QQuickView::*)(QQuickView::Status)>(&QQuickView::statusChanged), static_cast<MyQQuickView*>(ptr), static_cast<void (MyQQuickView::*)(QQuickView::Status)>(&MyQQuickView::Signal_StatusChanged));
}

void QQuickView_DisconnectStatusChanged(void* ptr)
{
	QObject::disconnect(static_cast<QQuickView*>(ptr), static_cast<void (QQuickView::*)(QQuickView::Status)>(&QQuickView::statusChanged), static_cast<MyQQuickView*>(ptr), static_cast<void (MyQQuickView::*)(QQuickView::Status)>(&MyQQuickView::Signal_StatusChanged));
}

void QQuickView_StatusChanged(void* ptr, int status)
{
	static_cast<QQuickView*>(ptr)->statusChanged(static_cast<QQuickView::Status>(status));
}

void QQuickView_DestroyQQuickView(void* ptr)
{
	static_cast<QQuickView*>(ptr)->~QQuickView();
}

void QQuickView_ReleaseResources(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QQuickView*>(ptr), "releaseResources");
}

void QQuickView_ReleaseResourcesDefault(void* ptr)
{
	static_cast<QQuickView*>(ptr)->QQuickView::releaseResources();
}

void QQuickView_Update(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QQuickView*>(ptr), "update");
}

void QQuickView_UpdateDefault(void* ptr)
{
	static_cast<QQuickView*>(ptr)->QQuickView::update();
}

void QQuickView_SetHeight(void* ptr, int arg)
{
	QMetaObject::invokeMethod(static_cast<QQuickView*>(ptr), "setHeight", Q_ARG(int, arg));
}

void QQuickView_SetHeightDefault(void* ptr, int arg)
{
	static_cast<QQuickView*>(ptr)->QQuickView::setHeight(arg);
}

void QQuickView_SetMaximumHeight(void* ptr, int h)
{
	QMetaObject::invokeMethod(static_cast<QQuickView*>(ptr), "setMaximumHeight", Q_ARG(int, h));
}

void QQuickView_SetMaximumHeightDefault(void* ptr, int h)
{
	static_cast<QQuickView*>(ptr)->QQuickView::setMaximumHeight(h);
}

void QQuickView_SetMaximumWidth(void* ptr, int w)
{
	QMetaObject::invokeMethod(static_cast<QQuickView*>(ptr), "setMaximumWidth", Q_ARG(int, w));
}

void QQuickView_SetMaximumWidthDefault(void* ptr, int w)
{
	static_cast<QQuickView*>(ptr)->QQuickView::setMaximumWidth(w);
}

void QQuickView_SetMinimumHeight(void* ptr, int h)
{
	QMetaObject::invokeMethod(static_cast<QQuickView*>(ptr), "setMinimumHeight", Q_ARG(int, h));
}

void QQuickView_SetMinimumHeightDefault(void* ptr, int h)
{
	static_cast<QQuickView*>(ptr)->QQuickView::setMinimumHeight(h);
}

void QQuickView_SetMinimumWidth(void* ptr, int w)
{
	QMetaObject::invokeMethod(static_cast<QQuickView*>(ptr), "setMinimumWidth", Q_ARG(int, w));
}

void QQuickView_SetMinimumWidthDefault(void* ptr, int w)
{
	static_cast<QQuickView*>(ptr)->QQuickView::setMinimumWidth(w);
}

void QQuickView_SetTitle(void* ptr, char* vqs)
{
	QMetaObject::invokeMethod(static_cast<QQuickView*>(ptr), "setTitle", Q_ARG(QString, QString(vqs)));
}

void QQuickView_SetTitleDefault(void* ptr, char* vqs)
{
	static_cast<QQuickView*>(ptr)->QQuickView::setTitle(QString(vqs));
}

void QQuickView_SetVisible(void* ptr, int visible)
{
	QMetaObject::invokeMethod(static_cast<QQuickView*>(ptr), "setVisible", Q_ARG(bool, visible != 0));
}

void QQuickView_SetVisibleDefault(void* ptr, int visible)
{
	static_cast<QQuickView*>(ptr)->QQuickView::setVisible(visible != 0);
}

void QQuickView_SetWidth(void* ptr, int arg)
{
	QMetaObject::invokeMethod(static_cast<QQuickView*>(ptr), "setWidth", Q_ARG(int, arg));
}

void QQuickView_SetWidthDefault(void* ptr, int arg)
{
	static_cast<QQuickView*>(ptr)->QQuickView::setWidth(arg);
}

void QQuickView_SetX(void* ptr, int arg)
{
	QMetaObject::invokeMethod(static_cast<QQuickView*>(ptr), "setX", Q_ARG(int, arg));
}

void QQuickView_SetXDefault(void* ptr, int arg)
{
	static_cast<QQuickView*>(ptr)->QQuickView::setX(arg);
}

void QQuickView_SetY(void* ptr, int arg)
{
	QMetaObject::invokeMethod(static_cast<QQuickView*>(ptr), "setY", Q_ARG(int, arg));
}

void QQuickView_SetYDefault(void* ptr, int arg)
{
	static_cast<QQuickView*>(ptr)->QQuickView::setY(arg);
}

void QQuickView_Alert(void* ptr, int msec)
{
	QMetaObject::invokeMethod(static_cast<QQuickView*>(ptr), "alert", Q_ARG(int, msec));
}

void QQuickView_AlertDefault(void* ptr, int msec)
{
	static_cast<QQuickView*>(ptr)->QQuickView::alert(msec);
}

int QQuickView_Close(void* ptr)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QQuickView*>(ptr), "close", Q_RETURN_ARG(bool, returnArg));
	return returnArg;
}

int QQuickView_CloseDefault(void* ptr)
{
	return static_cast<QQuickView*>(ptr)->QQuickView::close();
}

void* QQuickView_FocusObject(void* ptr)
{
	return static_cast<QQuickView*>(ptr)->focusObject();
}

void* QQuickView_FocusObjectDefault(void* ptr)
{
	return static_cast<QQuickView*>(ptr)->QQuickView::focusObject();
}

void* QQuickView_Format(void* ptr)
{
	return new QSurfaceFormat(static_cast<QQuickView*>(ptr)->format());
}

void* QQuickView_FormatDefault(void* ptr)
{
	return new QSurfaceFormat(static_cast<QQuickView*>(ptr)->QQuickView::format());
}

void QQuickView_Hide(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QQuickView*>(ptr), "hide");
}

void QQuickView_HideDefault(void* ptr)
{
	static_cast<QQuickView*>(ptr)->QQuickView::hide();
}

void QQuickView_Lower(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QQuickView*>(ptr), "lower");
}

void QQuickView_LowerDefault(void* ptr)
{
	static_cast<QQuickView*>(ptr)->QQuickView::lower();
}

void QQuickView_MoveEvent(void* ptr, void* ev)
{
	static_cast<QQuickView*>(ptr)->moveEvent(static_cast<QMoveEvent*>(ev));
}

void QQuickView_MoveEventDefault(void* ptr, void* ev)
{
	static_cast<QQuickView*>(ptr)->QQuickView::moveEvent(static_cast<QMoveEvent*>(ev));
}

int QQuickView_NativeEvent(void* ptr, char* eventType, void* message, long result)
{
	return static_cast<QQuickView*>(ptr)->nativeEvent(QByteArray(eventType), message, &result);
}

int QQuickView_NativeEventDefault(void* ptr, char* eventType, void* message, long result)
{
	return static_cast<QQuickView*>(ptr)->QQuickView::nativeEvent(QByteArray(eventType), message, &result);
}

void QQuickView_Raise(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QQuickView*>(ptr), "raise");
}

void QQuickView_RaiseDefault(void* ptr)
{
	static_cast<QQuickView*>(ptr)->QQuickView::raise();
}

void QQuickView_RequestActivate(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QQuickView*>(ptr), "requestActivate");
}

void QQuickView_RequestActivateDefault(void* ptr)
{
	static_cast<QQuickView*>(ptr)->QQuickView::requestActivate();
}

void QQuickView_RequestUpdate(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QQuickView*>(ptr), "requestUpdate");
}

void QQuickView_RequestUpdateDefault(void* ptr)
{
	static_cast<QQuickView*>(ptr)->QQuickView::requestUpdate();
}

void QQuickView_Show(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QQuickView*>(ptr), "show");
}

void QQuickView_ShowDefault(void* ptr)
{
	static_cast<QQuickView*>(ptr)->QQuickView::show();
}

void QQuickView_ShowFullScreen(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QQuickView*>(ptr), "showFullScreen");
}

void QQuickView_ShowFullScreenDefault(void* ptr)
{
	static_cast<QQuickView*>(ptr)->QQuickView::showFullScreen();
}

void QQuickView_ShowMaximized(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QQuickView*>(ptr), "showMaximized");
}

void QQuickView_ShowMaximizedDefault(void* ptr)
{
	static_cast<QQuickView*>(ptr)->QQuickView::showMaximized();
}

void QQuickView_ShowMinimized(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QQuickView*>(ptr), "showMinimized");
}

void QQuickView_ShowMinimizedDefault(void* ptr)
{
	static_cast<QQuickView*>(ptr)->QQuickView::showMinimized();
}

void QQuickView_ShowNormal(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QQuickView*>(ptr), "showNormal");
}

void QQuickView_ShowNormalDefault(void* ptr)
{
	static_cast<QQuickView*>(ptr)->QQuickView::showNormal();
}

void* QQuickView_Size(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QQuickView*>(ptr)->size()).width(), static_cast<QSize>(static_cast<QQuickView*>(ptr)->size()).height());
}

void* QQuickView_SizeDefault(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QQuickView*>(ptr)->QQuickView::size()).width(), static_cast<QSize>(static_cast<QQuickView*>(ptr)->QQuickView::size()).height());
}

int QQuickView_SurfaceType(void* ptr)
{
	return static_cast<QQuickView*>(ptr)->surfaceType();
}

int QQuickView_SurfaceTypeDefault(void* ptr)
{
	return static_cast<QQuickView*>(ptr)->QQuickView::surfaceType();
}

void QQuickView_TabletEvent(void* ptr, void* ev)
{
	static_cast<QQuickView*>(ptr)->tabletEvent(static_cast<QTabletEvent*>(ev));
}

void QQuickView_TabletEventDefault(void* ptr, void* ev)
{
	static_cast<QQuickView*>(ptr)->QQuickView::tabletEvent(static_cast<QTabletEvent*>(ev));
}

void QQuickView_TouchEvent(void* ptr, void* ev)
{
	static_cast<QQuickView*>(ptr)->touchEvent(static_cast<QTouchEvent*>(ev));
}

void QQuickView_TouchEventDefault(void* ptr, void* ev)
{
	static_cast<QQuickView*>(ptr)->QQuickView::touchEvent(static_cast<QTouchEvent*>(ev));
}

void QQuickView_TimerEvent(void* ptr, void* event)
{
	static_cast<QQuickView*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QQuickView_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QQuickView*>(ptr)->QQuickView::timerEvent(static_cast<QTimerEvent*>(event));
}

void QQuickView_ChildEvent(void* ptr, void* event)
{
	static_cast<QQuickView*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QQuickView_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QQuickView*>(ptr)->QQuickView::childEvent(static_cast<QChildEvent*>(event));
}

void QQuickView_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QQuickView*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQuickView_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QQuickView*>(ptr)->QQuickView::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQuickView_CustomEvent(void* ptr, void* event)
{
	static_cast<QQuickView*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QQuickView_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QQuickView*>(ptr)->QQuickView::customEvent(static_cast<QEvent*>(event));
}

void QQuickView_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QQuickView*>(ptr), "deleteLater");
}

void QQuickView_DeleteLaterDefault(void* ptr)
{
	static_cast<QQuickView*>(ptr)->QQuickView::deleteLater();
}

void QQuickView_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QQuickView*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQuickView_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QQuickView*>(ptr)->QQuickView::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QQuickView_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QQuickView*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QQuickView_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QQuickView*>(ptr)->QQuickView::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QQuickView_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QQuickView*>(ptr)->metaObject());
}

void* QQuickView_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QQuickView*>(ptr)->QQuickView::metaObject());
}

class MyQQuickWidget: public QQuickWidget
{
public:
	MyQQuickWidget(QQmlEngine *engine, QWidget *parent) : QQuickWidget(engine, parent) {};
	MyQQuickWidget(QWidget *parent) : QQuickWidget(parent) {};
	MyQQuickWidget(const QUrl &source, QWidget *parent) : QQuickWidget(source, parent) {};
	void dragEnterEvent(QDragEnterEvent * e) { callbackQQuickWidget_DragEnterEvent(this, this->objectName().toUtf8().data(), e); };
	void dragLeaveEvent(QDragLeaveEvent * e) { callbackQQuickWidget_DragLeaveEvent(this, this->objectName().toUtf8().data(), e); };
	void dragMoveEvent(QDragMoveEvent * e) { callbackQQuickWidget_DragMoveEvent(this, this->objectName().toUtf8().data(), e); };
	void dropEvent(QDropEvent * e) { callbackQQuickWidget_DropEvent(this, this->objectName().toUtf8().data(), e); };
	void focusInEvent(QFocusEvent * event) { callbackQQuickWidget_FocusInEvent(this, this->objectName().toUtf8().data(), event); };
	void focusOutEvent(QFocusEvent * event) { callbackQQuickWidget_FocusOutEvent(this, this->objectName().toUtf8().data(), event); };
	void hideEvent(QHideEvent * vqh) { callbackQQuickWidget_HideEvent(this, this->objectName().toUtf8().data(), vqh); };
	void keyPressEvent(QKeyEvent * e) { callbackQQuickWidget_KeyPressEvent(this, this->objectName().toUtf8().data(), e); };
	void keyReleaseEvent(QKeyEvent * e) { callbackQQuickWidget_KeyReleaseEvent(this, this->objectName().toUtf8().data(), e); };
	void mouseDoubleClickEvent(QMouseEvent * e) { callbackQQuickWidget_MouseDoubleClickEvent(this, this->objectName().toUtf8().data(), e); };
	void mouseMoveEvent(QMouseEvent * e) { callbackQQuickWidget_MouseMoveEvent(this, this->objectName().toUtf8().data(), e); };
	void mousePressEvent(QMouseEvent * e) { callbackQQuickWidget_MousePressEvent(this, this->objectName().toUtf8().data(), e); };
	void mouseReleaseEvent(QMouseEvent * e) { callbackQQuickWidget_MouseReleaseEvent(this, this->objectName().toUtf8().data(), e); };
	void Signal_SceneGraphError(QQuickWindow::SceneGraphError error, const QString & message) { callbackQQuickWidget_SceneGraphError(this, this->objectName().toUtf8().data(), error, message.toUtf8().data()); };
	void setSource(const QUrl & url) { callbackQQuickWidget_SetSource(this, this->objectName().toUtf8().data(), new QUrl(url)); };
	void showEvent(QShowEvent * vqs) { callbackQQuickWidget_ShowEvent(this, this->objectName().toUtf8().data(), vqs); };
	void Signal_StatusChanged(QQuickWidget::Status status) { callbackQQuickWidget_StatusChanged(this, this->objectName().toUtf8().data(), status); };
	void wheelEvent(QWheelEvent * e) { callbackQQuickWidget_WheelEvent(this, this->objectName().toUtf8().data(), e); };
	void actionEvent(QActionEvent * event) { callbackQQuickWidget_ActionEvent(this, this->objectName().toUtf8().data(), event); };
	void enterEvent(QEvent * event) { callbackQQuickWidget_EnterEvent(this, this->objectName().toUtf8().data(), event); };
	void leaveEvent(QEvent * event) { callbackQQuickWidget_LeaveEvent(this, this->objectName().toUtf8().data(), event); };
	QSize minimumSizeHint() const { return *static_cast<QSize*>(callbackQQuickWidget_MinimumSizeHint(const_cast<MyQQuickWidget*>(this), this->objectName().toUtf8().data())); };
	void moveEvent(QMoveEvent * event) { callbackQQuickWidget_MoveEvent(this, this->objectName().toUtf8().data(), event); };
	void paintEvent(QPaintEvent * event) { callbackQQuickWidget_PaintEvent(this, this->objectName().toUtf8().data(), event); };
	void setEnabled(bool vbo) { callbackQQuickWidget_SetEnabled(this, this->objectName().toUtf8().data(), vbo); };
	void setStyleSheet(const QString & styleSheet) { callbackQQuickWidget_SetStyleSheet(this, this->objectName().toUtf8().data(), styleSheet.toUtf8().data()); };
	void setVisible(bool visible) { callbackQQuickWidget_SetVisible(this, this->objectName().toUtf8().data(), visible); };
	void setWindowModified(bool vbo) { callbackQQuickWidget_SetWindowModified(this, this->objectName().toUtf8().data(), vbo); };
	void setWindowTitle(const QString & vqs) { callbackQQuickWidget_SetWindowTitle(this, this->objectName().toUtf8().data(), vqs.toUtf8().data()); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQQuickWidget_SizeHint(const_cast<MyQQuickWidget*>(this), this->objectName().toUtf8().data())); };
	void changeEvent(QEvent * event) { callbackQQuickWidget_ChangeEvent(this, this->objectName().toUtf8().data(), event); };
	bool close() { return callbackQQuickWidget_Close(this, this->objectName().toUtf8().data()) != 0; };
	void closeEvent(QCloseEvent * event) { callbackQQuickWidget_CloseEvent(this, this->objectName().toUtf8().data(), event); };
	void contextMenuEvent(QContextMenuEvent * event) { callbackQQuickWidget_ContextMenuEvent(this, this->objectName().toUtf8().data(), event); };
	bool focusNextPrevChild(bool next) { return callbackQQuickWidget_FocusNextPrevChild(this, this->objectName().toUtf8().data(), next) != 0; };
	bool hasHeightForWidth() const { return callbackQQuickWidget_HasHeightForWidth(const_cast<MyQQuickWidget*>(this), this->objectName().toUtf8().data()) != 0; };
	int heightForWidth(int w) const { return callbackQQuickWidget_HeightForWidth(const_cast<MyQQuickWidget*>(this), this->objectName().toUtf8().data(), w); };
	void hide() { callbackQQuickWidget_Hide(this, this->objectName().toUtf8().data()); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackQQuickWidget_InputMethodEvent(this, this->objectName().toUtf8().data(), event); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackQQuickWidget_InputMethodQuery(const_cast<MyQQuickWidget*>(this), this->objectName().toUtf8().data(), query)); };
	void lower() { callbackQQuickWidget_Lower(this, this->objectName().toUtf8().data()); };
	bool nativeEvent(const QByteArray & eventType, void * message, long * result) { return callbackQQuickWidget_NativeEvent(this, this->objectName().toUtf8().data(), QString(eventType).toUtf8().data(), message, *result) != 0; };
	void raise() { callbackQQuickWidget_Raise(this, this->objectName().toUtf8().data()); };
	void repaint() { callbackQQuickWidget_Repaint(this, this->objectName().toUtf8().data()); };
	void resizeEvent(QResizeEvent * event) { callbackQQuickWidget_ResizeEvent(this, this->objectName().toUtf8().data(), event); };
	void setDisabled(bool disable) { callbackQQuickWidget_SetDisabled(this, this->objectName().toUtf8().data(), disable); };
	void setFocus() { callbackQQuickWidget_SetFocus2(this, this->objectName().toUtf8().data()); };
	void setHidden(bool hidden) { callbackQQuickWidget_SetHidden(this, this->objectName().toUtf8().data(), hidden); };
	void show() { callbackQQuickWidget_Show(this, this->objectName().toUtf8().data()); };
	void showFullScreen() { callbackQQuickWidget_ShowFullScreen(this, this->objectName().toUtf8().data()); };
	void showMaximized() { callbackQQuickWidget_ShowMaximized(this, this->objectName().toUtf8().data()); };
	void showMinimized() { callbackQQuickWidget_ShowMinimized(this, this->objectName().toUtf8().data()); };
	void showNormal() { callbackQQuickWidget_ShowNormal(this, this->objectName().toUtf8().data()); };
	void tabletEvent(QTabletEvent * event) { callbackQQuickWidget_TabletEvent(this, this->objectName().toUtf8().data(), event); };
	void update() { callbackQQuickWidget_Update(this, this->objectName().toUtf8().data()); };
	void updateMicroFocus() { callbackQQuickWidget_UpdateMicroFocus(this, this->objectName().toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQQuickWidget_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQQuickWidget_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQQuickWidget_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQQuickWidget_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQQuickWidget_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQQuickWidget_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQQuickWidget_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQQuickWidget_MetaObject(const_cast<MyQQuickWidget*>(this), this->objectName().toUtf8().data())); };
};

int QQuickWidget_ResizeMode(void* ptr)
{
	return static_cast<QQuickWidget*>(ptr)->resizeMode();
}

void QQuickWidget_SetResizeMode(void* ptr, int vre)
{
	static_cast<QQuickWidget*>(ptr)->setResizeMode(static_cast<QQuickWidget::ResizeMode>(vre));
}

int QQuickWidget_Status(void* ptr)
{
	return static_cast<QQuickWidget*>(ptr)->status();
}

void* QQuickWidget_NewQQuickWidget2(void* engine, void* parent)
{
	return new MyQQuickWidget(static_cast<QQmlEngine*>(engine), static_cast<QWidget*>(parent));
}

void* QQuickWidget_NewQQuickWidget(void* parent)
{
	return new MyQQuickWidget(static_cast<QWidget*>(parent));
}

void* QQuickWidget_NewQQuickWidget3(void* source, void* parent)
{
	return new MyQQuickWidget(*static_cast<QUrl*>(source), static_cast<QWidget*>(parent));
}

void QQuickWidget_DragEnterEvent(void* ptr, void* e)
{
	static_cast<QQuickWidget*>(ptr)->dragEnterEvent(static_cast<QDragEnterEvent*>(e));
}

void QQuickWidget_DragEnterEventDefault(void* ptr, void* e)
{
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::dragEnterEvent(static_cast<QDragEnterEvent*>(e));
}

void QQuickWidget_DragLeaveEvent(void* ptr, void* e)
{
	static_cast<QQuickWidget*>(ptr)->dragLeaveEvent(static_cast<QDragLeaveEvent*>(e));
}

void QQuickWidget_DragLeaveEventDefault(void* ptr, void* e)
{
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::dragLeaveEvent(static_cast<QDragLeaveEvent*>(e));
}

void QQuickWidget_DragMoveEvent(void* ptr, void* e)
{
	static_cast<QQuickWidget*>(ptr)->dragMoveEvent(static_cast<QDragMoveEvent*>(e));
}

void QQuickWidget_DragMoveEventDefault(void* ptr, void* e)
{
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::dragMoveEvent(static_cast<QDragMoveEvent*>(e));
}

void QQuickWidget_DropEvent(void* ptr, void* e)
{
	static_cast<QQuickWidget*>(ptr)->dropEvent(static_cast<QDropEvent*>(e));
}

void QQuickWidget_DropEventDefault(void* ptr, void* e)
{
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::dropEvent(static_cast<QDropEvent*>(e));
}

void* QQuickWidget_Engine(void* ptr)
{
	return static_cast<QQuickWidget*>(ptr)->engine();
}

int QQuickWidget_Event(void* ptr, void* e)
{
	return static_cast<QQuickWidget*>(ptr)->event(static_cast<QEvent*>(e));
}

void QQuickWidget_FocusInEvent(void* ptr, void* event)
{
	static_cast<QQuickWidget*>(ptr)->focusInEvent(static_cast<QFocusEvent*>(event));
}

void QQuickWidget_FocusInEventDefault(void* ptr, void* event)
{
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::focusInEvent(static_cast<QFocusEvent*>(event));
}

void QQuickWidget_FocusOutEvent(void* ptr, void* event)
{
	static_cast<QQuickWidget*>(ptr)->focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QQuickWidget_FocusOutEventDefault(void* ptr, void* event)
{
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::focusOutEvent(static_cast<QFocusEvent*>(event));
}

void* QQuickWidget_Format(void* ptr)
{
	return new QSurfaceFormat(static_cast<QQuickWidget*>(ptr)->format());
}

void* QQuickWidget_GrabFramebuffer(void* ptr)
{
	return new QImage(static_cast<QQuickWidget*>(ptr)->grabFramebuffer());
}

void QQuickWidget_HideEvent(void* ptr, void* vqh)
{
	static_cast<QQuickWidget*>(ptr)->hideEvent(static_cast<QHideEvent*>(vqh));
}

void QQuickWidget_HideEventDefault(void* ptr, void* vqh)
{
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::hideEvent(static_cast<QHideEvent*>(vqh));
}

void* QQuickWidget_InitialSize(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QQuickWidget*>(ptr)->initialSize()).width(), static_cast<QSize>(static_cast<QQuickWidget*>(ptr)->initialSize()).height());
}

void QQuickWidget_KeyPressEvent(void* ptr, void* e)
{
	static_cast<QQuickWidget*>(ptr)->keyPressEvent(static_cast<QKeyEvent*>(e));
}

void QQuickWidget_KeyPressEventDefault(void* ptr, void* e)
{
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::keyPressEvent(static_cast<QKeyEvent*>(e));
}

void QQuickWidget_KeyReleaseEvent(void* ptr, void* e)
{
	static_cast<QQuickWidget*>(ptr)->keyReleaseEvent(static_cast<QKeyEvent*>(e));
}

void QQuickWidget_KeyReleaseEventDefault(void* ptr, void* e)
{
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::keyReleaseEvent(static_cast<QKeyEvent*>(e));
}

void QQuickWidget_MouseDoubleClickEvent(void* ptr, void* e)
{
	static_cast<QQuickWidget*>(ptr)->mouseDoubleClickEvent(static_cast<QMouseEvent*>(e));
}

void QQuickWidget_MouseDoubleClickEventDefault(void* ptr, void* e)
{
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::mouseDoubleClickEvent(static_cast<QMouseEvent*>(e));
}

void QQuickWidget_MouseMoveEvent(void* ptr, void* e)
{
	static_cast<QQuickWidget*>(ptr)->mouseMoveEvent(static_cast<QMouseEvent*>(e));
}

void QQuickWidget_MouseMoveEventDefault(void* ptr, void* e)
{
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::mouseMoveEvent(static_cast<QMouseEvent*>(e));
}

void QQuickWidget_MousePressEvent(void* ptr, void* e)
{
	static_cast<QQuickWidget*>(ptr)->mousePressEvent(static_cast<QMouseEvent*>(e));
}

void QQuickWidget_MousePressEventDefault(void* ptr, void* e)
{
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::mousePressEvent(static_cast<QMouseEvent*>(e));
}

void QQuickWidget_MouseReleaseEvent(void* ptr, void* e)
{
	static_cast<QQuickWidget*>(ptr)->mouseReleaseEvent(static_cast<QMouseEvent*>(e));
}

void QQuickWidget_MouseReleaseEventDefault(void* ptr, void* e)
{
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::mouseReleaseEvent(static_cast<QMouseEvent*>(e));
}

void* QQuickWidget_QuickWindow(void* ptr)
{
	return static_cast<QQuickWidget*>(ptr)->quickWindow();
}

void* QQuickWidget_RootContext(void* ptr)
{
	return static_cast<QQuickWidget*>(ptr)->rootContext();
}

void* QQuickWidget_RootObject(void* ptr)
{
	return static_cast<QQuickWidget*>(ptr)->rootObject();
}

void QQuickWidget_ConnectSceneGraphError(void* ptr)
{
	QObject::connect(static_cast<QQuickWidget*>(ptr), static_cast<void (QQuickWidget::*)(QQuickWindow::SceneGraphError, const QString &)>(&QQuickWidget::sceneGraphError), static_cast<MyQQuickWidget*>(ptr), static_cast<void (MyQQuickWidget::*)(QQuickWindow::SceneGraphError, const QString &)>(&MyQQuickWidget::Signal_SceneGraphError));
}

void QQuickWidget_DisconnectSceneGraphError(void* ptr)
{
	QObject::disconnect(static_cast<QQuickWidget*>(ptr), static_cast<void (QQuickWidget::*)(QQuickWindow::SceneGraphError, const QString &)>(&QQuickWidget::sceneGraphError), static_cast<MyQQuickWidget*>(ptr), static_cast<void (MyQQuickWidget::*)(QQuickWindow::SceneGraphError, const QString &)>(&MyQQuickWidget::Signal_SceneGraphError));
}

void QQuickWidget_SceneGraphError(void* ptr, int error, char* message)
{
	static_cast<QQuickWidget*>(ptr)->sceneGraphError(static_cast<QQuickWindow::SceneGraphError>(error), QString(message));
}

void QQuickWidget_SetClearColor(void* ptr, void* color)
{
	static_cast<QQuickWidget*>(ptr)->setClearColor(*static_cast<QColor*>(color));
}

void QQuickWidget_SetFormat(void* ptr, void* format)
{
	static_cast<QQuickWidget*>(ptr)->setFormat(*static_cast<QSurfaceFormat*>(format));
}

void QQuickWidget_SetSource(void* ptr, void* url)
{
	QMetaObject::invokeMethod(static_cast<QQuickWidget*>(ptr), "setSource", Q_ARG(QUrl, *static_cast<QUrl*>(url)));
}

void QQuickWidget_ShowEvent(void* ptr, void* vqs)
{
	static_cast<QQuickWidget*>(ptr)->showEvent(static_cast<QShowEvent*>(vqs));
}

void QQuickWidget_ShowEventDefault(void* ptr, void* vqs)
{
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::showEvent(static_cast<QShowEvent*>(vqs));
}

void QQuickWidget_ConnectStatusChanged(void* ptr)
{
	QObject::connect(static_cast<QQuickWidget*>(ptr), static_cast<void (QQuickWidget::*)(QQuickWidget::Status)>(&QQuickWidget::statusChanged), static_cast<MyQQuickWidget*>(ptr), static_cast<void (MyQQuickWidget::*)(QQuickWidget::Status)>(&MyQQuickWidget::Signal_StatusChanged));
}

void QQuickWidget_DisconnectStatusChanged(void* ptr)
{
	QObject::disconnect(static_cast<QQuickWidget*>(ptr), static_cast<void (QQuickWidget::*)(QQuickWidget::Status)>(&QQuickWidget::statusChanged), static_cast<MyQQuickWidget*>(ptr), static_cast<void (MyQQuickWidget::*)(QQuickWidget::Status)>(&MyQQuickWidget::Signal_StatusChanged));
}

void QQuickWidget_StatusChanged(void* ptr, int status)
{
	static_cast<QQuickWidget*>(ptr)->statusChanged(static_cast<QQuickWidget::Status>(status));
}

void* QQuickWidget_Source(void* ptr)
{
	return new QUrl(static_cast<QQuickWidget*>(ptr)->source());
}

void QQuickWidget_WheelEvent(void* ptr, void* e)
{
	static_cast<QQuickWidget*>(ptr)->wheelEvent(static_cast<QWheelEvent*>(e));
}

void QQuickWidget_WheelEventDefault(void* ptr, void* e)
{
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::wheelEvent(static_cast<QWheelEvent*>(e));
}

void QQuickWidget_DestroyQQuickWidget(void* ptr)
{
	static_cast<QQuickWidget*>(ptr)->~QQuickWidget();
}

void QQuickWidget_ActionEvent(void* ptr, void* event)
{
	static_cast<QQuickWidget*>(ptr)->actionEvent(static_cast<QActionEvent*>(event));
}

void QQuickWidget_ActionEventDefault(void* ptr, void* event)
{
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::actionEvent(static_cast<QActionEvent*>(event));
}

void QQuickWidget_EnterEvent(void* ptr, void* event)
{
	static_cast<QQuickWidget*>(ptr)->enterEvent(static_cast<QEvent*>(event));
}

void QQuickWidget_EnterEventDefault(void* ptr, void* event)
{
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::enterEvent(static_cast<QEvent*>(event));
}

void QQuickWidget_LeaveEvent(void* ptr, void* event)
{
	static_cast<QQuickWidget*>(ptr)->leaveEvent(static_cast<QEvent*>(event));
}

void QQuickWidget_LeaveEventDefault(void* ptr, void* event)
{
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::leaveEvent(static_cast<QEvent*>(event));
}

void* QQuickWidget_MinimumSizeHint(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QQuickWidget*>(ptr)->minimumSizeHint()).width(), static_cast<QSize>(static_cast<QQuickWidget*>(ptr)->minimumSizeHint()).height());
}

void* QQuickWidget_MinimumSizeHintDefault(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QQuickWidget*>(ptr)->QQuickWidget::minimumSizeHint()).width(), static_cast<QSize>(static_cast<QQuickWidget*>(ptr)->QQuickWidget::minimumSizeHint()).height());
}

void QQuickWidget_MoveEvent(void* ptr, void* event)
{
	static_cast<QQuickWidget*>(ptr)->moveEvent(static_cast<QMoveEvent*>(event));
}

void QQuickWidget_MoveEventDefault(void* ptr, void* event)
{
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::moveEvent(static_cast<QMoveEvent*>(event));
}

void QQuickWidget_PaintEvent(void* ptr, void* event)
{
	static_cast<QQuickWidget*>(ptr)->paintEvent(static_cast<QPaintEvent*>(event));
}

void QQuickWidget_PaintEventDefault(void* ptr, void* event)
{
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::paintEvent(static_cast<QPaintEvent*>(event));
}

void QQuickWidget_SetEnabled(void* ptr, int vbo)
{
	QMetaObject::invokeMethod(static_cast<QQuickWidget*>(ptr), "setEnabled", Q_ARG(bool, vbo != 0));
}

void QQuickWidget_SetEnabledDefault(void* ptr, int vbo)
{
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::setEnabled(vbo != 0);
}

void QQuickWidget_SetStyleSheet(void* ptr, char* styleSheet)
{
	QMetaObject::invokeMethod(static_cast<QQuickWidget*>(ptr), "setStyleSheet", Q_ARG(QString, QString(styleSheet)));
}

void QQuickWidget_SetStyleSheetDefault(void* ptr, char* styleSheet)
{
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::setStyleSheet(QString(styleSheet));
}

void QQuickWidget_SetVisible(void* ptr, int visible)
{
	QMetaObject::invokeMethod(static_cast<QQuickWidget*>(ptr), "setVisible", Q_ARG(bool, visible != 0));
}

void QQuickWidget_SetVisibleDefault(void* ptr, int visible)
{
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::setVisible(visible != 0);
}

void QQuickWidget_SetWindowModified(void* ptr, int vbo)
{
	QMetaObject::invokeMethod(static_cast<QQuickWidget*>(ptr), "setWindowModified", Q_ARG(bool, vbo != 0));
}

void QQuickWidget_SetWindowModifiedDefault(void* ptr, int vbo)
{
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::setWindowModified(vbo != 0);
}

void QQuickWidget_SetWindowTitle(void* ptr, char* vqs)
{
	QMetaObject::invokeMethod(static_cast<QQuickWidget*>(ptr), "setWindowTitle", Q_ARG(QString, QString(vqs)));
}

void QQuickWidget_SetWindowTitleDefault(void* ptr, char* vqs)
{
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::setWindowTitle(QString(vqs));
}

void* QQuickWidget_SizeHint(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QQuickWidget*>(ptr)->sizeHint()).width(), static_cast<QSize>(static_cast<QQuickWidget*>(ptr)->sizeHint()).height());
}

void* QQuickWidget_SizeHintDefault(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QQuickWidget*>(ptr)->QQuickWidget::sizeHint()).width(), static_cast<QSize>(static_cast<QQuickWidget*>(ptr)->QQuickWidget::sizeHint()).height());
}

void QQuickWidget_ChangeEvent(void* ptr, void* event)
{
	static_cast<QQuickWidget*>(ptr)->changeEvent(static_cast<QEvent*>(event));
}

void QQuickWidget_ChangeEventDefault(void* ptr, void* event)
{
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::changeEvent(static_cast<QEvent*>(event));
}

int QQuickWidget_Close(void* ptr)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QQuickWidget*>(ptr), "close", Q_RETURN_ARG(bool, returnArg));
	return returnArg;
}

int QQuickWidget_CloseDefault(void* ptr)
{
	return static_cast<QQuickWidget*>(ptr)->QQuickWidget::close();
}

void QQuickWidget_CloseEvent(void* ptr, void* event)
{
	static_cast<QQuickWidget*>(ptr)->closeEvent(static_cast<QCloseEvent*>(event));
}

void QQuickWidget_CloseEventDefault(void* ptr, void* event)
{
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::closeEvent(static_cast<QCloseEvent*>(event));
}

void QQuickWidget_ContextMenuEvent(void* ptr, void* event)
{
	static_cast<QQuickWidget*>(ptr)->contextMenuEvent(static_cast<QContextMenuEvent*>(event));
}

void QQuickWidget_ContextMenuEventDefault(void* ptr, void* event)
{
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
}

int QQuickWidget_FocusNextPrevChild(void* ptr, int next)
{
	return static_cast<QQuickWidget*>(ptr)->focusNextPrevChild(next != 0);
}

int QQuickWidget_FocusNextPrevChildDefault(void* ptr, int next)
{
	return static_cast<QQuickWidget*>(ptr)->QQuickWidget::focusNextPrevChild(next != 0);
}

int QQuickWidget_HasHeightForWidth(void* ptr)
{
	return static_cast<QQuickWidget*>(ptr)->hasHeightForWidth();
}

int QQuickWidget_HasHeightForWidthDefault(void* ptr)
{
	return static_cast<QQuickWidget*>(ptr)->QQuickWidget::hasHeightForWidth();
}

int QQuickWidget_HeightForWidth(void* ptr, int w)
{
	return static_cast<QQuickWidget*>(ptr)->heightForWidth(w);
}

int QQuickWidget_HeightForWidthDefault(void* ptr, int w)
{
	return static_cast<QQuickWidget*>(ptr)->QQuickWidget::heightForWidth(w);
}

void QQuickWidget_Hide(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QQuickWidget*>(ptr), "hide");
}

void QQuickWidget_HideDefault(void* ptr)
{
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::hide();
}

void QQuickWidget_InputMethodEvent(void* ptr, void* event)
{
	static_cast<QQuickWidget*>(ptr)->inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void QQuickWidget_InputMethodEventDefault(void* ptr, void* event)
{
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void* QQuickWidget_InputMethodQuery(void* ptr, int query)
{
	return new QVariant(static_cast<QQuickWidget*>(ptr)->inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

void* QQuickWidget_InputMethodQueryDefault(void* ptr, int query)
{
	return new QVariant(static_cast<QQuickWidget*>(ptr)->QQuickWidget::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

void QQuickWidget_Lower(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QQuickWidget*>(ptr), "lower");
}

void QQuickWidget_LowerDefault(void* ptr)
{
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::lower();
}

int QQuickWidget_NativeEvent(void* ptr, char* eventType, void* message, long result)
{
	return static_cast<QQuickWidget*>(ptr)->nativeEvent(QByteArray(eventType), message, &result);
}

int QQuickWidget_NativeEventDefault(void* ptr, char* eventType, void* message, long result)
{
	return static_cast<QQuickWidget*>(ptr)->QQuickWidget::nativeEvent(QByteArray(eventType), message, &result);
}

void QQuickWidget_Raise(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QQuickWidget*>(ptr), "raise");
}

void QQuickWidget_RaiseDefault(void* ptr)
{
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::raise();
}

void QQuickWidget_Repaint(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QQuickWidget*>(ptr), "repaint");
}

void QQuickWidget_RepaintDefault(void* ptr)
{
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::repaint();
}

void QQuickWidget_ResizeEvent(void* ptr, void* event)
{
	static_cast<QQuickWidget*>(ptr)->resizeEvent(static_cast<QResizeEvent*>(event));
}

void QQuickWidget_ResizeEventDefault(void* ptr, void* event)
{
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::resizeEvent(static_cast<QResizeEvent*>(event));
}

void QQuickWidget_SetDisabled(void* ptr, int disable)
{
	QMetaObject::invokeMethod(static_cast<QQuickWidget*>(ptr), "setDisabled", Q_ARG(bool, disable != 0));
}

void QQuickWidget_SetDisabledDefault(void* ptr, int disable)
{
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::setDisabled(disable != 0);
}

void QQuickWidget_SetFocus2(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QQuickWidget*>(ptr), "setFocus");
}

void QQuickWidget_SetFocus2Default(void* ptr)
{
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::setFocus();
}

void QQuickWidget_SetHidden(void* ptr, int hidden)
{
	QMetaObject::invokeMethod(static_cast<QQuickWidget*>(ptr), "setHidden", Q_ARG(bool, hidden != 0));
}

void QQuickWidget_SetHiddenDefault(void* ptr, int hidden)
{
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::setHidden(hidden != 0);
}

void QQuickWidget_Show(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QQuickWidget*>(ptr), "show");
}

void QQuickWidget_ShowDefault(void* ptr)
{
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::show();
}

void QQuickWidget_ShowFullScreen(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QQuickWidget*>(ptr), "showFullScreen");
}

void QQuickWidget_ShowFullScreenDefault(void* ptr)
{
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::showFullScreen();
}

void QQuickWidget_ShowMaximized(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QQuickWidget*>(ptr), "showMaximized");
}

void QQuickWidget_ShowMaximizedDefault(void* ptr)
{
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::showMaximized();
}

void QQuickWidget_ShowMinimized(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QQuickWidget*>(ptr), "showMinimized");
}

void QQuickWidget_ShowMinimizedDefault(void* ptr)
{
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::showMinimized();
}

void QQuickWidget_ShowNormal(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QQuickWidget*>(ptr), "showNormal");
}

void QQuickWidget_ShowNormalDefault(void* ptr)
{
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::showNormal();
}

void QQuickWidget_TabletEvent(void* ptr, void* event)
{
	static_cast<QQuickWidget*>(ptr)->tabletEvent(static_cast<QTabletEvent*>(event));
}

void QQuickWidget_TabletEventDefault(void* ptr, void* event)
{
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::tabletEvent(static_cast<QTabletEvent*>(event));
}

void QQuickWidget_Update(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QQuickWidget*>(ptr), "update");
}

void QQuickWidget_UpdateDefault(void* ptr)
{
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::update();
}

void QQuickWidget_UpdateMicroFocus(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QQuickWidget*>(ptr), "updateMicroFocus");
}

void QQuickWidget_UpdateMicroFocusDefault(void* ptr)
{
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::updateMicroFocus();
}

void QQuickWidget_TimerEvent(void* ptr, void* event)
{
	static_cast<QQuickWidget*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QQuickWidget_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::timerEvent(static_cast<QTimerEvent*>(event));
}

void QQuickWidget_ChildEvent(void* ptr, void* event)
{
	static_cast<QQuickWidget*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QQuickWidget_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::childEvent(static_cast<QChildEvent*>(event));
}

void QQuickWidget_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QQuickWidget*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQuickWidget_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQuickWidget_CustomEvent(void* ptr, void* event)
{
	static_cast<QQuickWidget*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QQuickWidget_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::customEvent(static_cast<QEvent*>(event));
}

void QQuickWidget_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QQuickWidget*>(ptr), "deleteLater");
}

void QQuickWidget_DeleteLaterDefault(void* ptr)
{
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::deleteLater();
}

void QQuickWidget_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QQuickWidget*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQuickWidget_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QQuickWidget_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QQuickWidget*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QQuickWidget_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QQuickWidget*>(ptr)->QQuickWidget::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QQuickWidget_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QQuickWidget*>(ptr)->metaObject());
}

void* QQuickWidget_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QQuickWidget*>(ptr)->QQuickWidget::metaObject());
}

class MyQQuickWindow: public QQuickWindow
{
public:
	MyQQuickWindow(QWindow *parent) : QQuickWindow(parent) {};
	void Signal_ActiveFocusItemChanged() { callbackQQuickWindow_ActiveFocusItemChanged(this, this->objectName().toUtf8().data()); };
	void Signal_AfterAnimating() { callbackQQuickWindow_AfterAnimating(this, this->objectName().toUtf8().data()); };
	void Signal_AfterRendering() { callbackQQuickWindow_AfterRendering(this, this->objectName().toUtf8().data()); };
	void Signal_AfterSynchronizing() { callbackQQuickWindow_AfterSynchronizing(this, this->objectName().toUtf8().data()); };
	void Signal_BeforeRendering() { callbackQQuickWindow_BeforeRendering(this, this->objectName().toUtf8().data()); };
	void Signal_BeforeSynchronizing() { callbackQQuickWindow_BeforeSynchronizing(this, this->objectName().toUtf8().data()); };
	void Signal_ColorChanged(const QColor & vqc) { callbackQQuickWindow_ColorChanged(this, this->objectName().toUtf8().data(), new QColor(vqc)); };
	void Signal_FrameSwapped() { callbackQQuickWindow_FrameSwapped(this, this->objectName().toUtf8().data()); };
	void Signal_OpenglContextCreated(QOpenGLContext * context) { callbackQQuickWindow_OpenglContextCreated(this, this->objectName().toUtf8().data(), context); };
	void releaseResources() { callbackQQuickWindow_ReleaseResources(this, this->objectName().toUtf8().data()); };
	void Signal_SceneGraphAboutToStop() { callbackQQuickWindow_SceneGraphAboutToStop(this, this->objectName().toUtf8().data()); };
	void Signal_SceneGraphError(QQuickWindow::SceneGraphError error, const QString & message) { callbackQQuickWindow_SceneGraphError(this, this->objectName().toUtf8().data(), error, message.toUtf8().data()); };
	void Signal_SceneGraphInitialized() { callbackQQuickWindow_SceneGraphInitialized(this, this->objectName().toUtf8().data()); };
	void Signal_SceneGraphInvalidated() { callbackQQuickWindow_SceneGraphInvalidated(this, this->objectName().toUtf8().data()); };
	void update() { callbackQQuickWindow_Update(this, this->objectName().toUtf8().data()); };
	void setHeight(int arg) { callbackQQuickWindow_SetHeight(this, this->objectName().toUtf8().data(), arg); };
	void setMaximumHeight(int h) { callbackQQuickWindow_SetMaximumHeight(this, this->objectName().toUtf8().data(), h); };
	void setMaximumWidth(int w) { callbackQQuickWindow_SetMaximumWidth(this, this->objectName().toUtf8().data(), w); };
	void setMinimumHeight(int h) { callbackQQuickWindow_SetMinimumHeight(this, this->objectName().toUtf8().data(), h); };
	void setMinimumWidth(int w) { callbackQQuickWindow_SetMinimumWidth(this, this->objectName().toUtf8().data(), w); };
	void setTitle(const QString & vqs) { callbackQQuickWindow_SetTitle(this, this->objectName().toUtf8().data(), vqs.toUtf8().data()); };
	void setVisible(bool visible) { callbackQQuickWindow_SetVisible(this, this->objectName().toUtf8().data(), visible); };
	void setWidth(int arg) { callbackQQuickWindow_SetWidth(this, this->objectName().toUtf8().data(), arg); };
	void setX(int arg) { callbackQQuickWindow_SetX(this, this->objectName().toUtf8().data(), arg); };
	void setY(int arg) { callbackQQuickWindow_SetY(this, this->objectName().toUtf8().data(), arg); };
	void alert(int msec) { callbackQQuickWindow_Alert(this, this->objectName().toUtf8().data(), msec); };
	bool close() { return callbackQQuickWindow_Close(this, this->objectName().toUtf8().data()) != 0; };
	QObject * focusObject() const { return static_cast<QObject*>(callbackQQuickWindow_FocusObject(const_cast<MyQQuickWindow*>(this), this->objectName().toUtf8().data())); };
	QSurfaceFormat format() const { return *static_cast<QSurfaceFormat*>(callbackQQuickWindow_Format(const_cast<MyQQuickWindow*>(this), this->objectName().toUtf8().data())); };
	void hide() { callbackQQuickWindow_Hide(this, this->objectName().toUtf8().data()); };
	void lower() { callbackQQuickWindow_Lower(this, this->objectName().toUtf8().data()); };
	void moveEvent(QMoveEvent * ev) { callbackQQuickWindow_MoveEvent(this, this->objectName().toUtf8().data(), ev); };
	bool nativeEvent(const QByteArray & eventType, void * message, long * result) { return callbackQQuickWindow_NativeEvent(this, this->objectName().toUtf8().data(), QString(eventType).toUtf8().data(), message, *result) != 0; };
	void raise() { callbackQQuickWindow_Raise(this, this->objectName().toUtf8().data()); };
	void requestActivate() { callbackQQuickWindow_RequestActivate(this, this->objectName().toUtf8().data()); };
	void requestUpdate() { callbackQQuickWindow_RequestUpdate(this, this->objectName().toUtf8().data()); };
	void show() { callbackQQuickWindow_Show(this, this->objectName().toUtf8().data()); };
	void showFullScreen() { callbackQQuickWindow_ShowFullScreen(this, this->objectName().toUtf8().data()); };
	void showMaximized() { callbackQQuickWindow_ShowMaximized(this, this->objectName().toUtf8().data()); };
	void showMinimized() { callbackQQuickWindow_ShowMinimized(this, this->objectName().toUtf8().data()); };
	void showNormal() { callbackQQuickWindow_ShowNormal(this, this->objectName().toUtf8().data()); };
	QSize size() const { return *static_cast<QSize*>(callbackQQuickWindow_Size(const_cast<MyQQuickWindow*>(this), this->objectName().toUtf8().data())); };
	SurfaceType surfaceType() const { return static_cast<QSurface::SurfaceType>(callbackQQuickWindow_SurfaceType(const_cast<MyQQuickWindow*>(this), this->objectName().toUtf8().data())); };
	void tabletEvent(QTabletEvent * ev) { callbackQQuickWindow_TabletEvent(this, this->objectName().toUtf8().data(), ev); };
	void touchEvent(QTouchEvent * ev) { callbackQQuickWindow_TouchEvent(this, this->objectName().toUtf8().data(), ev); };
	void timerEvent(QTimerEvent * event) { callbackQQuickWindow_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQQuickWindow_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQQuickWindow_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQQuickWindow_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQQuickWindow_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQQuickWindow_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQQuickWindow_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQQuickWindow_MetaObject(const_cast<MyQQuickWindow*>(this), this->objectName().toUtf8().data())); };
};

void* QQuickWindow_ActiveFocusItem(void* ptr)
{
	return static_cast<QQuickWindow*>(ptr)->activeFocusItem();
}

void* QQuickWindow_Color(void* ptr)
{
	return new QColor(static_cast<QQuickWindow*>(ptr)->color());
}

void* QQuickWindow_ContentItem(void* ptr)
{
	return static_cast<QQuickWindow*>(ptr)->contentItem();
}

void QQuickWindow_SetColor(void* ptr, void* color)
{
	static_cast<QQuickWindow*>(ptr)->setColor(*static_cast<QColor*>(color));
}

void* QQuickWindow_NewQQuickWindow(void* parent)
{
	return new MyQQuickWindow(static_cast<QWindow*>(parent));
}

void* QQuickWindow_AccessibleRoot(void* ptr)
{
	return static_cast<QQuickWindow*>(ptr)->accessibleRoot();
}

void QQuickWindow_ConnectActiveFocusItemChanged(void* ptr)
{
	QObject::connect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::activeFocusItemChanged), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_ActiveFocusItemChanged));
}

void QQuickWindow_DisconnectActiveFocusItemChanged(void* ptr)
{
	QObject::disconnect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::activeFocusItemChanged), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_ActiveFocusItemChanged));
}

void QQuickWindow_ActiveFocusItemChanged(void* ptr)
{
	static_cast<QQuickWindow*>(ptr)->activeFocusItemChanged();
}

void QQuickWindow_ConnectAfterAnimating(void* ptr)
{
	QObject::connect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::afterAnimating), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_AfterAnimating));
}

void QQuickWindow_DisconnectAfterAnimating(void* ptr)
{
	QObject::disconnect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::afterAnimating), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_AfterAnimating));
}

void QQuickWindow_AfterAnimating(void* ptr)
{
	static_cast<QQuickWindow*>(ptr)->afterAnimating();
}

void QQuickWindow_ConnectAfterRendering(void* ptr)
{
	QObject::connect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::afterRendering), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_AfterRendering));
}

void QQuickWindow_DisconnectAfterRendering(void* ptr)
{
	QObject::disconnect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::afterRendering), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_AfterRendering));
}

void QQuickWindow_AfterRendering(void* ptr)
{
	static_cast<QQuickWindow*>(ptr)->afterRendering();
}

void QQuickWindow_ConnectAfterSynchronizing(void* ptr)
{
	QObject::connect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::afterSynchronizing), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_AfterSynchronizing));
}

void QQuickWindow_DisconnectAfterSynchronizing(void* ptr)
{
	QObject::disconnect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::afterSynchronizing), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_AfterSynchronizing));
}

void QQuickWindow_AfterSynchronizing(void* ptr)
{
	static_cast<QQuickWindow*>(ptr)->afterSynchronizing();
}

void QQuickWindow_ConnectBeforeRendering(void* ptr)
{
	QObject::connect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::beforeRendering), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_BeforeRendering));
}

void QQuickWindow_DisconnectBeforeRendering(void* ptr)
{
	QObject::disconnect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::beforeRendering), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_BeforeRendering));
}

void QQuickWindow_BeforeRendering(void* ptr)
{
	static_cast<QQuickWindow*>(ptr)->beforeRendering();
}

void QQuickWindow_ConnectBeforeSynchronizing(void* ptr)
{
	QObject::connect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::beforeSynchronizing), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_BeforeSynchronizing));
}

void QQuickWindow_DisconnectBeforeSynchronizing(void* ptr)
{
	QObject::disconnect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::beforeSynchronizing), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_BeforeSynchronizing));
}

void QQuickWindow_BeforeSynchronizing(void* ptr)
{
	static_cast<QQuickWindow*>(ptr)->beforeSynchronizing();
}

int QQuickWindow_ClearBeforeRendering(void* ptr)
{
	return static_cast<QQuickWindow*>(ptr)->clearBeforeRendering();
}

void QQuickWindow_ConnectColorChanged(void* ptr)
{
	QObject::connect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)(const QColor &)>(&QQuickWindow::colorChanged), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)(const QColor &)>(&MyQQuickWindow::Signal_ColorChanged));
}

void QQuickWindow_DisconnectColorChanged(void* ptr)
{
	QObject::disconnect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)(const QColor &)>(&QQuickWindow::colorChanged), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)(const QColor &)>(&MyQQuickWindow::Signal_ColorChanged));
}

void QQuickWindow_ColorChanged(void* ptr, void* vqc)
{
	static_cast<QQuickWindow*>(ptr)->colorChanged(*static_cast<QColor*>(vqc));
}

void* QQuickWindow_CreateTextureFromImage2(void* ptr, void* image)
{
	return static_cast<QQuickWindow*>(ptr)->createTextureFromImage(*static_cast<QImage*>(image));
}

void* QQuickWindow_CreateTextureFromImage(void* ptr, void* image, int options)
{
	return static_cast<QQuickWindow*>(ptr)->createTextureFromImage(*static_cast<QImage*>(image), static_cast<QQuickWindow::CreateTextureOption>(options));
}

double QQuickWindow_EffectiveDevicePixelRatio(void* ptr)
{
	return static_cast<double>(static_cast<QQuickWindow*>(ptr)->effectiveDevicePixelRatio());
}

int QQuickWindow_Event(void* ptr, void* e)
{
	return static_cast<QQuickWindow*>(ptr)->event(static_cast<QEvent*>(e));
}

void QQuickWindow_ExposeEvent(void* ptr, void* vqe)
{
	static_cast<QQuickWindow*>(ptr)->exposeEvent(static_cast<QExposeEvent*>(vqe));
}

void QQuickWindow_FocusInEvent(void* ptr, void* ev)
{
	static_cast<QQuickWindow*>(ptr)->focusInEvent(static_cast<QFocusEvent*>(ev));
}

void QQuickWindow_FocusOutEvent(void* ptr, void* ev)
{
	static_cast<QQuickWindow*>(ptr)->focusOutEvent(static_cast<QFocusEvent*>(ev));
}

void QQuickWindow_ConnectFrameSwapped(void* ptr)
{
	QObject::connect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::frameSwapped), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_FrameSwapped));
}

void QQuickWindow_DisconnectFrameSwapped(void* ptr)
{
	QObject::disconnect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::frameSwapped), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_FrameSwapped));
}

void QQuickWindow_FrameSwapped(void* ptr)
{
	static_cast<QQuickWindow*>(ptr)->frameSwapped();
}

void* QQuickWindow_GrabWindow(void* ptr)
{
	return new QImage(static_cast<QQuickWindow*>(ptr)->grabWindow());
}

int QQuickWindow_QQuickWindow_HasDefaultAlphaBuffer()
{
	return QQuickWindow::hasDefaultAlphaBuffer();
}

void QQuickWindow_HideEvent(void* ptr, void* vqh)
{
	static_cast<QQuickWindow*>(ptr)->hideEvent(static_cast<QHideEvent*>(vqh));
}

void* QQuickWindow_IncubationController(void* ptr)
{
	return static_cast<QQuickWindow*>(ptr)->incubationController();
}

int QQuickWindow_IsPersistentOpenGLContext(void* ptr)
{
	return static_cast<QQuickWindow*>(ptr)->isPersistentOpenGLContext();
}

int QQuickWindow_IsPersistentSceneGraph(void* ptr)
{
	return static_cast<QQuickWindow*>(ptr)->isPersistentSceneGraph();
}

int QQuickWindow_IsSceneGraphInitialized(void* ptr)
{
	return static_cast<QQuickWindow*>(ptr)->isSceneGraphInitialized();
}

void QQuickWindow_KeyPressEvent(void* ptr, void* e)
{
	static_cast<QQuickWindow*>(ptr)->keyPressEvent(static_cast<QKeyEvent*>(e));
}

void QQuickWindow_KeyReleaseEvent(void* ptr, void* e)
{
	static_cast<QQuickWindow*>(ptr)->keyReleaseEvent(static_cast<QKeyEvent*>(e));
}

void QQuickWindow_MouseDoubleClickEvent(void* ptr, void* event)
{
	static_cast<QQuickWindow*>(ptr)->mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void* QQuickWindow_MouseGrabberItem(void* ptr)
{
	return static_cast<QQuickWindow*>(ptr)->mouseGrabberItem();
}

void QQuickWindow_MouseMoveEvent(void* ptr, void* event)
{
	static_cast<QQuickWindow*>(ptr)->mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QQuickWindow_MousePressEvent(void* ptr, void* event)
{
	static_cast<QQuickWindow*>(ptr)->mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QQuickWindow_MouseReleaseEvent(void* ptr, void* event)
{
	static_cast<QQuickWindow*>(ptr)->mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void* QQuickWindow_OpenglContext(void* ptr)
{
	return static_cast<QQuickWindow*>(ptr)->openglContext();
}

void QQuickWindow_ConnectOpenglContextCreated(void* ptr)
{
	QObject::connect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)(QOpenGLContext *)>(&QQuickWindow::openglContextCreated), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)(QOpenGLContext *)>(&MyQQuickWindow::Signal_OpenglContextCreated));
}

void QQuickWindow_DisconnectOpenglContextCreated(void* ptr)
{
	QObject::disconnect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)(QOpenGLContext *)>(&QQuickWindow::openglContextCreated), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)(QOpenGLContext *)>(&MyQQuickWindow::Signal_OpenglContextCreated));
}

void QQuickWindow_OpenglContextCreated(void* ptr, void* context)
{
	static_cast<QQuickWindow*>(ptr)->openglContextCreated(static_cast<QOpenGLContext*>(context));
}

void QQuickWindow_ReleaseResources(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QQuickWindow*>(ptr), "releaseResources");
}

void* QQuickWindow_RenderTarget(void* ptr)
{
	return static_cast<QQuickWindow*>(ptr)->renderTarget();
}

void* QQuickWindow_RenderTargetSize(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QQuickWindow*>(ptr)->renderTargetSize()).width(), static_cast<QSize>(static_cast<QQuickWindow*>(ptr)->renderTargetSize()).height());
}

void QQuickWindow_ResetOpenGLState(void* ptr)
{
	static_cast<QQuickWindow*>(ptr)->resetOpenGLState();
}

void QQuickWindow_ResizeEvent(void* ptr, void* ev)
{
	static_cast<QQuickWindow*>(ptr)->resizeEvent(static_cast<QResizeEvent*>(ev));
}

void QQuickWindow_ConnectSceneGraphAboutToStop(void* ptr)
{
	QObject::connect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::sceneGraphAboutToStop), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_SceneGraphAboutToStop));
}

void QQuickWindow_DisconnectSceneGraphAboutToStop(void* ptr)
{
	QObject::disconnect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::sceneGraphAboutToStop), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_SceneGraphAboutToStop));
}

void QQuickWindow_SceneGraphAboutToStop(void* ptr)
{
	static_cast<QQuickWindow*>(ptr)->sceneGraphAboutToStop();
}

void QQuickWindow_ConnectSceneGraphError(void* ptr)
{
	QObject::connect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)(QQuickWindow::SceneGraphError, const QString &)>(&QQuickWindow::sceneGraphError), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)(QQuickWindow::SceneGraphError, const QString &)>(&MyQQuickWindow::Signal_SceneGraphError));
}

void QQuickWindow_DisconnectSceneGraphError(void* ptr)
{
	QObject::disconnect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)(QQuickWindow::SceneGraphError, const QString &)>(&QQuickWindow::sceneGraphError), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)(QQuickWindow::SceneGraphError, const QString &)>(&MyQQuickWindow::Signal_SceneGraphError));
}

void QQuickWindow_SceneGraphError(void* ptr, int error, char* message)
{
	static_cast<QQuickWindow*>(ptr)->sceneGraphError(static_cast<QQuickWindow::SceneGraphError>(error), QString(message));
}

void QQuickWindow_ConnectSceneGraphInitialized(void* ptr)
{
	QObject::connect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::sceneGraphInitialized), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_SceneGraphInitialized));
}

void QQuickWindow_DisconnectSceneGraphInitialized(void* ptr)
{
	QObject::disconnect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::sceneGraphInitialized), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_SceneGraphInitialized));
}

void QQuickWindow_SceneGraphInitialized(void* ptr)
{
	static_cast<QQuickWindow*>(ptr)->sceneGraphInitialized();
}

void QQuickWindow_ConnectSceneGraphInvalidated(void* ptr)
{
	QObject::connect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::sceneGraphInvalidated), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_SceneGraphInvalidated));
}

void QQuickWindow_DisconnectSceneGraphInvalidated(void* ptr)
{
	QObject::disconnect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::sceneGraphInvalidated), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_SceneGraphInvalidated));
}

void QQuickWindow_SceneGraphInvalidated(void* ptr)
{
	static_cast<QQuickWindow*>(ptr)->sceneGraphInvalidated();
}

void QQuickWindow_ScheduleRenderJob(void* ptr, void* job, int stage)
{
	static_cast<QQuickWindow*>(ptr)->scheduleRenderJob(static_cast<QRunnable*>(job), static_cast<QQuickWindow::RenderStage>(stage));
}

int QQuickWindow_SendEvent(void* ptr, void* item, void* e)
{
	return static_cast<QQuickWindow*>(ptr)->sendEvent(static_cast<QQuickItem*>(item), static_cast<QEvent*>(e));
}

void QQuickWindow_SetClearBeforeRendering(void* ptr, int enabled)
{
	static_cast<QQuickWindow*>(ptr)->setClearBeforeRendering(enabled != 0);
}

void QQuickWindow_QQuickWindow_SetDefaultAlphaBuffer(int useAlpha)
{
	QQuickWindow::setDefaultAlphaBuffer(useAlpha != 0);
}

void QQuickWindow_SetPersistentOpenGLContext(void* ptr, int persistent)
{
	static_cast<QQuickWindow*>(ptr)->setPersistentOpenGLContext(persistent != 0);
}

void QQuickWindow_SetPersistentSceneGraph(void* ptr, int persistent)
{
	static_cast<QQuickWindow*>(ptr)->setPersistentSceneGraph(persistent != 0);
}

void QQuickWindow_SetRenderTarget(void* ptr, void* fbo)
{
	static_cast<QQuickWindow*>(ptr)->setRenderTarget(static_cast<QOpenGLFramebufferObject*>(fbo));
}

void QQuickWindow_ShowEvent(void* ptr, void* vqs)
{
	static_cast<QQuickWindow*>(ptr)->showEvent(static_cast<QShowEvent*>(vqs));
}

void QQuickWindow_Update(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QQuickWindow*>(ptr), "update");
}

void QQuickWindow_WheelEvent(void* ptr, void* event)
{
	static_cast<QQuickWindow*>(ptr)->wheelEvent(static_cast<QWheelEvent*>(event));
}

void QQuickWindow_DestroyQQuickWindow(void* ptr)
{
	static_cast<QQuickWindow*>(ptr)->~QQuickWindow();
}

void QQuickWindow_SetHeight(void* ptr, int arg)
{
	QMetaObject::invokeMethod(static_cast<QQuickWindow*>(ptr), "setHeight", Q_ARG(int, arg));
}

void QQuickWindow_SetHeightDefault(void* ptr, int arg)
{
	static_cast<QQuickWindow*>(ptr)->QQuickWindow::setHeight(arg);
}

void QQuickWindow_SetMaximumHeight(void* ptr, int h)
{
	QMetaObject::invokeMethod(static_cast<QQuickWindow*>(ptr), "setMaximumHeight", Q_ARG(int, h));
}

void QQuickWindow_SetMaximumHeightDefault(void* ptr, int h)
{
	static_cast<QQuickWindow*>(ptr)->QQuickWindow::setMaximumHeight(h);
}

void QQuickWindow_SetMaximumWidth(void* ptr, int w)
{
	QMetaObject::invokeMethod(static_cast<QQuickWindow*>(ptr), "setMaximumWidth", Q_ARG(int, w));
}

void QQuickWindow_SetMaximumWidthDefault(void* ptr, int w)
{
	static_cast<QQuickWindow*>(ptr)->QQuickWindow::setMaximumWidth(w);
}

void QQuickWindow_SetMinimumHeight(void* ptr, int h)
{
	QMetaObject::invokeMethod(static_cast<QQuickWindow*>(ptr), "setMinimumHeight", Q_ARG(int, h));
}

void QQuickWindow_SetMinimumHeightDefault(void* ptr, int h)
{
	static_cast<QQuickWindow*>(ptr)->QQuickWindow::setMinimumHeight(h);
}

void QQuickWindow_SetMinimumWidth(void* ptr, int w)
{
	QMetaObject::invokeMethod(static_cast<QQuickWindow*>(ptr), "setMinimumWidth", Q_ARG(int, w));
}

void QQuickWindow_SetMinimumWidthDefault(void* ptr, int w)
{
	static_cast<QQuickWindow*>(ptr)->QQuickWindow::setMinimumWidth(w);
}

void QQuickWindow_SetTitle(void* ptr, char* vqs)
{
	QMetaObject::invokeMethod(static_cast<QQuickWindow*>(ptr), "setTitle", Q_ARG(QString, QString(vqs)));
}

void QQuickWindow_SetTitleDefault(void* ptr, char* vqs)
{
	static_cast<QQuickWindow*>(ptr)->QQuickWindow::setTitle(QString(vqs));
}

void QQuickWindow_SetVisible(void* ptr, int visible)
{
	QMetaObject::invokeMethod(static_cast<QQuickWindow*>(ptr), "setVisible", Q_ARG(bool, visible != 0));
}

void QQuickWindow_SetVisibleDefault(void* ptr, int visible)
{
	static_cast<QQuickWindow*>(ptr)->QQuickWindow::setVisible(visible != 0);
}

void QQuickWindow_SetWidth(void* ptr, int arg)
{
	QMetaObject::invokeMethod(static_cast<QQuickWindow*>(ptr), "setWidth", Q_ARG(int, arg));
}

void QQuickWindow_SetWidthDefault(void* ptr, int arg)
{
	static_cast<QQuickWindow*>(ptr)->QQuickWindow::setWidth(arg);
}

void QQuickWindow_SetX(void* ptr, int arg)
{
	QMetaObject::invokeMethod(static_cast<QQuickWindow*>(ptr), "setX", Q_ARG(int, arg));
}

void QQuickWindow_SetXDefault(void* ptr, int arg)
{
	static_cast<QQuickWindow*>(ptr)->QQuickWindow::setX(arg);
}

void QQuickWindow_SetY(void* ptr, int arg)
{
	QMetaObject::invokeMethod(static_cast<QQuickWindow*>(ptr), "setY", Q_ARG(int, arg));
}

void QQuickWindow_SetYDefault(void* ptr, int arg)
{
	static_cast<QQuickWindow*>(ptr)->QQuickWindow::setY(arg);
}

void QQuickWindow_Alert(void* ptr, int msec)
{
	QMetaObject::invokeMethod(static_cast<QQuickWindow*>(ptr), "alert", Q_ARG(int, msec));
}

void QQuickWindow_AlertDefault(void* ptr, int msec)
{
	static_cast<QQuickWindow*>(ptr)->QQuickWindow::alert(msec);
}

int QQuickWindow_Close(void* ptr)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QQuickWindow*>(ptr), "close", Q_RETURN_ARG(bool, returnArg));
	return returnArg;
}

int QQuickWindow_CloseDefault(void* ptr)
{
	return static_cast<QQuickWindow*>(ptr)->QQuickWindow::close();
}

void* QQuickWindow_FocusObject(void* ptr)
{
	return static_cast<QQuickWindow*>(ptr)->focusObject();
}

void* QQuickWindow_FocusObjectDefault(void* ptr)
{
	return static_cast<QQuickWindow*>(ptr)->QQuickWindow::focusObject();
}

void* QQuickWindow_Format(void* ptr)
{
	return new QSurfaceFormat(static_cast<QQuickWindow*>(ptr)->format());
}

void* QQuickWindow_FormatDefault(void* ptr)
{
	return new QSurfaceFormat(static_cast<QQuickWindow*>(ptr)->QQuickWindow::format());
}

void QQuickWindow_Hide(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QQuickWindow*>(ptr), "hide");
}

void QQuickWindow_HideDefault(void* ptr)
{
	static_cast<QQuickWindow*>(ptr)->QQuickWindow::hide();
}

void QQuickWindow_Lower(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QQuickWindow*>(ptr), "lower");
}

void QQuickWindow_LowerDefault(void* ptr)
{
	static_cast<QQuickWindow*>(ptr)->QQuickWindow::lower();
}

void QQuickWindow_MoveEvent(void* ptr, void* ev)
{
	static_cast<QQuickWindow*>(ptr)->moveEvent(static_cast<QMoveEvent*>(ev));
}

void QQuickWindow_MoveEventDefault(void* ptr, void* ev)
{
	static_cast<QQuickWindow*>(ptr)->QQuickWindow::moveEvent(static_cast<QMoveEvent*>(ev));
}

int QQuickWindow_NativeEvent(void* ptr, char* eventType, void* message, long result)
{
	return static_cast<QQuickWindow*>(ptr)->nativeEvent(QByteArray(eventType), message, &result);
}

int QQuickWindow_NativeEventDefault(void* ptr, char* eventType, void* message, long result)
{
	return static_cast<QQuickWindow*>(ptr)->QQuickWindow::nativeEvent(QByteArray(eventType), message, &result);
}

void QQuickWindow_Raise(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QQuickWindow*>(ptr), "raise");
}

void QQuickWindow_RaiseDefault(void* ptr)
{
	static_cast<QQuickWindow*>(ptr)->QQuickWindow::raise();
}

void QQuickWindow_RequestActivate(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QQuickWindow*>(ptr), "requestActivate");
}

void QQuickWindow_RequestActivateDefault(void* ptr)
{
	static_cast<QQuickWindow*>(ptr)->QQuickWindow::requestActivate();
}

void QQuickWindow_RequestUpdate(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QQuickWindow*>(ptr), "requestUpdate");
}

void QQuickWindow_RequestUpdateDefault(void* ptr)
{
	static_cast<QQuickWindow*>(ptr)->QQuickWindow::requestUpdate();
}

void QQuickWindow_Show(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QQuickWindow*>(ptr), "show");
}

void QQuickWindow_ShowDefault(void* ptr)
{
	static_cast<QQuickWindow*>(ptr)->QQuickWindow::show();
}

void QQuickWindow_ShowFullScreen(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QQuickWindow*>(ptr), "showFullScreen");
}

void QQuickWindow_ShowFullScreenDefault(void* ptr)
{
	static_cast<QQuickWindow*>(ptr)->QQuickWindow::showFullScreen();
}

void QQuickWindow_ShowMaximized(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QQuickWindow*>(ptr), "showMaximized");
}

void QQuickWindow_ShowMaximizedDefault(void* ptr)
{
	static_cast<QQuickWindow*>(ptr)->QQuickWindow::showMaximized();
}

void QQuickWindow_ShowMinimized(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QQuickWindow*>(ptr), "showMinimized");
}

void QQuickWindow_ShowMinimizedDefault(void* ptr)
{
	static_cast<QQuickWindow*>(ptr)->QQuickWindow::showMinimized();
}

void QQuickWindow_ShowNormal(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QQuickWindow*>(ptr), "showNormal");
}

void QQuickWindow_ShowNormalDefault(void* ptr)
{
	static_cast<QQuickWindow*>(ptr)->QQuickWindow::showNormal();
}

void* QQuickWindow_Size(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QQuickWindow*>(ptr)->size()).width(), static_cast<QSize>(static_cast<QQuickWindow*>(ptr)->size()).height());
}

void* QQuickWindow_SizeDefault(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QQuickWindow*>(ptr)->QQuickWindow::size()).width(), static_cast<QSize>(static_cast<QQuickWindow*>(ptr)->QQuickWindow::size()).height());
}

int QQuickWindow_SurfaceType(void* ptr)
{
	return static_cast<QQuickWindow*>(ptr)->surfaceType();
}

int QQuickWindow_SurfaceTypeDefault(void* ptr)
{
	return static_cast<QQuickWindow*>(ptr)->QQuickWindow::surfaceType();
}

void QQuickWindow_TabletEvent(void* ptr, void* ev)
{
	static_cast<QQuickWindow*>(ptr)->tabletEvent(static_cast<QTabletEvent*>(ev));
}

void QQuickWindow_TabletEventDefault(void* ptr, void* ev)
{
	static_cast<QQuickWindow*>(ptr)->QQuickWindow::tabletEvent(static_cast<QTabletEvent*>(ev));
}

void QQuickWindow_TouchEvent(void* ptr, void* ev)
{
	static_cast<QQuickWindow*>(ptr)->touchEvent(static_cast<QTouchEvent*>(ev));
}

void QQuickWindow_TouchEventDefault(void* ptr, void* ev)
{
	static_cast<QQuickWindow*>(ptr)->QQuickWindow::touchEvent(static_cast<QTouchEvent*>(ev));
}

void QQuickWindow_TimerEvent(void* ptr, void* event)
{
	static_cast<QQuickWindow*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QQuickWindow_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QQuickWindow*>(ptr)->QQuickWindow::timerEvent(static_cast<QTimerEvent*>(event));
}

void QQuickWindow_ChildEvent(void* ptr, void* event)
{
	static_cast<QQuickWindow*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QQuickWindow_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QQuickWindow*>(ptr)->QQuickWindow::childEvent(static_cast<QChildEvent*>(event));
}

void QQuickWindow_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QQuickWindow*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQuickWindow_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QQuickWindow*>(ptr)->QQuickWindow::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQuickWindow_CustomEvent(void* ptr, void* event)
{
	static_cast<QQuickWindow*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QQuickWindow_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QQuickWindow*>(ptr)->QQuickWindow::customEvent(static_cast<QEvent*>(event));
}

void QQuickWindow_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QQuickWindow*>(ptr), "deleteLater");
}

void QQuickWindow_DeleteLaterDefault(void* ptr)
{
	static_cast<QQuickWindow*>(ptr)->QQuickWindow::deleteLater();
}

void QQuickWindow_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QQuickWindow*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQuickWindow_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QQuickWindow*>(ptr)->QQuickWindow::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QQuickWindow_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QQuickWindow*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QQuickWindow_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QQuickWindow*>(ptr)->QQuickWindow::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QQuickWindow_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QQuickWindow*>(ptr)->metaObject());
}

void* QQuickWindow_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QQuickWindow*>(ptr)->QQuickWindow::metaObject());
}

class MyQSGAbstractRenderer: public QSGAbstractRenderer
{
public:
	void Signal_SceneGraphChanged() { callbackQSGAbstractRenderer_SceneGraphChanged(this, this->objectName().toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQSGAbstractRenderer_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQSGAbstractRenderer_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSGAbstractRenderer_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQSGAbstractRenderer_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQSGAbstractRenderer_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSGAbstractRenderer_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQSGAbstractRenderer_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSGAbstractRenderer_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSGAbstractRenderer_MetaObject(const_cast<MyQSGAbstractRenderer*>(this), this->objectName().toUtf8().data())); };
};

void* QSGAbstractRenderer_ClearColor(void* ptr)
{
	return new QColor(static_cast<QSGAbstractRenderer*>(ptr)->clearColor());
}

int QSGAbstractRenderer_ClearMode(void* ptr)
{
	return static_cast<QSGAbstractRenderer*>(ptr)->clearMode();
}

void* QSGAbstractRenderer_DeviceRect(void* ptr)
{
	return new QRect(static_cast<QRect>(static_cast<QSGAbstractRenderer*>(ptr)->deviceRect()).x(), static_cast<QRect>(static_cast<QSGAbstractRenderer*>(ptr)->deviceRect()).y(), static_cast<QRect>(static_cast<QSGAbstractRenderer*>(ptr)->deviceRect()).width(), static_cast<QRect>(static_cast<QSGAbstractRenderer*>(ptr)->deviceRect()).height());
}

void QSGAbstractRenderer_ConnectSceneGraphChanged(void* ptr)
{
	QObject::connect(static_cast<QSGAbstractRenderer*>(ptr), static_cast<void (QSGAbstractRenderer::*)()>(&QSGAbstractRenderer::sceneGraphChanged), static_cast<MyQSGAbstractRenderer*>(ptr), static_cast<void (MyQSGAbstractRenderer::*)()>(&MyQSGAbstractRenderer::Signal_SceneGraphChanged));
}

void QSGAbstractRenderer_DisconnectSceneGraphChanged(void* ptr)
{
	QObject::disconnect(static_cast<QSGAbstractRenderer*>(ptr), static_cast<void (QSGAbstractRenderer::*)()>(&QSGAbstractRenderer::sceneGraphChanged), static_cast<MyQSGAbstractRenderer*>(ptr), static_cast<void (MyQSGAbstractRenderer::*)()>(&MyQSGAbstractRenderer::Signal_SceneGraphChanged));
}

void QSGAbstractRenderer_SceneGraphChanged(void* ptr)
{
	static_cast<QSGAbstractRenderer*>(ptr)->sceneGraphChanged();
}

void QSGAbstractRenderer_SetClearColor(void* ptr, void* color)
{
	static_cast<QSGAbstractRenderer*>(ptr)->setClearColor(*static_cast<QColor*>(color));
}

void QSGAbstractRenderer_SetClearMode(void* ptr, int mode)
{
	static_cast<QSGAbstractRenderer*>(ptr)->setClearMode(static_cast<QSGAbstractRenderer::ClearModeBit>(mode));
}

void QSGAbstractRenderer_SetDeviceRect(void* ptr, void* rect)
{
	static_cast<QSGAbstractRenderer*>(ptr)->setDeviceRect(*static_cast<QRect*>(rect));
}

void QSGAbstractRenderer_SetDeviceRect2(void* ptr, void* size)
{
	static_cast<QSGAbstractRenderer*>(ptr)->setDeviceRect(*static_cast<QSize*>(size));
}

void QSGAbstractRenderer_SetProjectionMatrix(void* ptr, void* matrix)
{
	static_cast<QSGAbstractRenderer*>(ptr)->setProjectionMatrix(*static_cast<QMatrix4x4*>(matrix));
}

void QSGAbstractRenderer_SetProjectionMatrixToRect(void* ptr, void* rect)
{
	static_cast<QSGAbstractRenderer*>(ptr)->setProjectionMatrixToRect(*static_cast<QRectF*>(rect));
}

void QSGAbstractRenderer_SetViewportRect(void* ptr, void* rect)
{
	static_cast<QSGAbstractRenderer*>(ptr)->setViewportRect(*static_cast<QRect*>(rect));
}

void QSGAbstractRenderer_SetViewportRect2(void* ptr, void* size)
{
	static_cast<QSGAbstractRenderer*>(ptr)->setViewportRect(*static_cast<QSize*>(size));
}

void* QSGAbstractRenderer_ViewportRect(void* ptr)
{
	return new QRect(static_cast<QRect>(static_cast<QSGAbstractRenderer*>(ptr)->viewportRect()).x(), static_cast<QRect>(static_cast<QSGAbstractRenderer*>(ptr)->viewportRect()).y(), static_cast<QRect>(static_cast<QSGAbstractRenderer*>(ptr)->viewportRect()).width(), static_cast<QRect>(static_cast<QSGAbstractRenderer*>(ptr)->viewportRect()).height());
}

void QSGAbstractRenderer_TimerEvent(void* ptr, void* event)
{
	static_cast<QSGAbstractRenderer*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QSGAbstractRenderer_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QSGAbstractRenderer*>(ptr)->QSGAbstractRenderer::timerEvent(static_cast<QTimerEvent*>(event));
}

void QSGAbstractRenderer_ChildEvent(void* ptr, void* event)
{
	static_cast<QSGAbstractRenderer*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QSGAbstractRenderer_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QSGAbstractRenderer*>(ptr)->QSGAbstractRenderer::childEvent(static_cast<QChildEvent*>(event));
}

void QSGAbstractRenderer_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QSGAbstractRenderer*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSGAbstractRenderer_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QSGAbstractRenderer*>(ptr)->QSGAbstractRenderer::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSGAbstractRenderer_CustomEvent(void* ptr, void* event)
{
	static_cast<QSGAbstractRenderer*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QSGAbstractRenderer_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QSGAbstractRenderer*>(ptr)->QSGAbstractRenderer::customEvent(static_cast<QEvent*>(event));
}

void QSGAbstractRenderer_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QSGAbstractRenderer*>(ptr), "deleteLater");
}

void QSGAbstractRenderer_DeleteLaterDefault(void* ptr)
{
	static_cast<QSGAbstractRenderer*>(ptr)->QSGAbstractRenderer::deleteLater();
}

void QSGAbstractRenderer_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QSGAbstractRenderer*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSGAbstractRenderer_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QSGAbstractRenderer*>(ptr)->QSGAbstractRenderer::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QSGAbstractRenderer_Event(void* ptr, void* e)
{
	return static_cast<QSGAbstractRenderer*>(ptr)->event(static_cast<QEvent*>(e));
}

int QSGAbstractRenderer_EventDefault(void* ptr, void* e)
{
	return static_cast<QSGAbstractRenderer*>(ptr)->QSGAbstractRenderer::event(static_cast<QEvent*>(e));
}

int QSGAbstractRenderer_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QSGAbstractRenderer*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QSGAbstractRenderer_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QSGAbstractRenderer*>(ptr)->QSGAbstractRenderer::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QSGAbstractRenderer_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QSGAbstractRenderer*>(ptr)->metaObject());
}

void* QSGAbstractRenderer_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QSGAbstractRenderer*>(ptr)->QSGAbstractRenderer::metaObject());
}

void* QSGBasicGeometryNode_Geometry2(void* ptr)
{
	return static_cast<QSGBasicGeometryNode*>(ptr)->geometry();
}

void* QSGBasicGeometryNode_Geometry(void* ptr)
{
	return const_cast<QSGGeometry*>(static_cast<QSGBasicGeometryNode*>(ptr)->geometry());
}

void QSGBasicGeometryNode_SetGeometry(void* ptr, void* geometry)
{
	static_cast<QSGBasicGeometryNode*>(ptr)->setGeometry(static_cast<QSGGeometry*>(geometry));
}

void QSGBasicGeometryNode_DestroyQSGBasicGeometryNode(void* ptr)
{
	static_cast<QSGBasicGeometryNode*>(ptr)->~QSGBasicGeometryNode();
}

int QSGBasicGeometryNode_IsSubtreeBlocked(void* ptr)
{
	return static_cast<QSGBasicGeometryNode*>(ptr)->isSubtreeBlocked();
}

int QSGBasicGeometryNode_IsSubtreeBlockedDefault(void* ptr)
{
	return static_cast<QSGBasicGeometryNode*>(ptr)->QSGBasicGeometryNode::isSubtreeBlocked();
}

void QSGBasicGeometryNode_Preprocess(void* ptr)
{
	static_cast<QSGBasicGeometryNode*>(ptr)->preprocess();
}

void QSGBasicGeometryNode_PreprocessDefault(void* ptr)
{
	static_cast<QSGBasicGeometryNode*>(ptr)->QSGBasicGeometryNode::preprocess();
}

void* QSGClipNode_NewQSGClipNode()
{
	return new QSGClipNode();
}

void* QSGClipNode_ClipRect(void* ptr)
{
	return new QRectF(static_cast<QRectF>(static_cast<QSGClipNode*>(ptr)->clipRect()).x(), static_cast<QRectF>(static_cast<QSGClipNode*>(ptr)->clipRect()).y(), static_cast<QRectF>(static_cast<QSGClipNode*>(ptr)->clipRect()).width(), static_cast<QRectF>(static_cast<QSGClipNode*>(ptr)->clipRect()).height());
}

int QSGClipNode_IsRectangular(void* ptr)
{
	return static_cast<QSGClipNode*>(ptr)->isRectangular();
}

void QSGClipNode_SetClipRect(void* ptr, void* rect)
{
	static_cast<QSGClipNode*>(ptr)->setClipRect(*static_cast<QRectF*>(rect));
}

void QSGClipNode_SetIsRectangular(void* ptr, int rectHint)
{
	static_cast<QSGClipNode*>(ptr)->setIsRectangular(rectHint != 0);
}

void QSGClipNode_DestroyQSGClipNode(void* ptr)
{
	static_cast<QSGClipNode*>(ptr)->~QSGClipNode();
}

int QSGClipNode_IsSubtreeBlocked(void* ptr)
{
	return static_cast<QSGClipNode*>(ptr)->isSubtreeBlocked();
}

int QSGClipNode_IsSubtreeBlockedDefault(void* ptr)
{
	return static_cast<QSGClipNode*>(ptr)->QSGClipNode::isSubtreeBlocked();
}

void QSGClipNode_Preprocess(void* ptr)
{
	static_cast<QSGClipNode*>(ptr)->preprocess();
}

void QSGClipNode_PreprocessDefault(void* ptr)
{
	static_cast<QSGClipNode*>(ptr)->QSGClipNode::preprocess();
}

class MyQSGDynamicTexture: public QSGDynamicTexture
{
public:
	bool updateTexture() { return callbackQSGDynamicTexture_UpdateTexture(this, this->objectName().toUtf8().data()) != 0; };
	void bind() { callbackQSGDynamicTexture_Bind(this, this->objectName().toUtf8().data()); };
	bool hasAlphaChannel() const { return callbackQSGDynamicTexture_HasAlphaChannel(const_cast<MyQSGDynamicTexture*>(this), this->objectName().toUtf8().data()) != 0; };
	bool hasMipmaps() const { return callbackQSGDynamicTexture_HasMipmaps(const_cast<MyQSGDynamicTexture*>(this), this->objectName().toUtf8().data()) != 0; };
	bool isAtlasTexture() const { return callbackQSGDynamicTexture_IsAtlasTexture(const_cast<MyQSGDynamicTexture*>(this), this->objectName().toUtf8().data()) != 0; };
	QRectF normalizedTextureSubRect() const { return *static_cast<QRectF*>(callbackQSGDynamicTexture_NormalizedTextureSubRect(const_cast<MyQSGDynamicTexture*>(this), this->objectName().toUtf8().data())); };
	QSGTexture * removedFromAtlas() const { return static_cast<QSGTexture*>(callbackQSGDynamicTexture_RemovedFromAtlas(const_cast<MyQSGDynamicTexture*>(this), this->objectName().toUtf8().data())); };
	int textureId() const { return callbackQSGDynamicTexture_TextureId(const_cast<MyQSGDynamicTexture*>(this), this->objectName().toUtf8().data()); };
	QSize textureSize() const { return *static_cast<QSize*>(callbackQSGDynamicTexture_TextureSize(const_cast<MyQSGDynamicTexture*>(this), this->objectName().toUtf8().data())); };
	void timerEvent(QTimerEvent * event) { callbackQSGDynamicTexture_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQSGDynamicTexture_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSGDynamicTexture_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQSGDynamicTexture_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQSGDynamicTexture_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSGDynamicTexture_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQSGDynamicTexture_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSGDynamicTexture_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSGDynamicTexture_MetaObject(const_cast<MyQSGDynamicTexture*>(this), this->objectName().toUtf8().data())); };
};

int QSGDynamicTexture_UpdateTexture(void* ptr)
{
	return static_cast<QSGDynamicTexture*>(ptr)->updateTexture();
}

void QSGDynamicTexture_Bind(void* ptr)
{
	static_cast<QSGDynamicTexture*>(ptr)->bind();
}



int QSGDynamicTexture_HasAlphaChannel(void* ptr)
{
	return static_cast<QSGDynamicTexture*>(ptr)->hasAlphaChannel();
}



int QSGDynamicTexture_HasMipmaps(void* ptr)
{
	return static_cast<QSGDynamicTexture*>(ptr)->hasMipmaps();
}



int QSGDynamicTexture_IsAtlasTexture(void* ptr)
{
	return static_cast<QSGDynamicTexture*>(ptr)->isAtlasTexture();
}

int QSGDynamicTexture_IsAtlasTextureDefault(void* ptr)
{
	return static_cast<QSGDynamicTexture*>(ptr)->QSGDynamicTexture::isAtlasTexture();
}

void* QSGDynamicTexture_NormalizedTextureSubRect(void* ptr)
{
	return new QRectF(static_cast<QRectF>(static_cast<QSGDynamicTexture*>(ptr)->normalizedTextureSubRect()).x(), static_cast<QRectF>(static_cast<QSGDynamicTexture*>(ptr)->normalizedTextureSubRect()).y(), static_cast<QRectF>(static_cast<QSGDynamicTexture*>(ptr)->normalizedTextureSubRect()).width(), static_cast<QRectF>(static_cast<QSGDynamicTexture*>(ptr)->normalizedTextureSubRect()).height());
}

void* QSGDynamicTexture_NormalizedTextureSubRectDefault(void* ptr)
{
	return new QRectF(static_cast<QRectF>(static_cast<QSGDynamicTexture*>(ptr)->QSGDynamicTexture::normalizedTextureSubRect()).x(), static_cast<QRectF>(static_cast<QSGDynamicTexture*>(ptr)->QSGDynamicTexture::normalizedTextureSubRect()).y(), static_cast<QRectF>(static_cast<QSGDynamicTexture*>(ptr)->QSGDynamicTexture::normalizedTextureSubRect()).width(), static_cast<QRectF>(static_cast<QSGDynamicTexture*>(ptr)->QSGDynamicTexture::normalizedTextureSubRect()).height());
}

void* QSGDynamicTexture_RemovedFromAtlas(void* ptr)
{
	return static_cast<QSGDynamicTexture*>(ptr)->removedFromAtlas();
}

void* QSGDynamicTexture_RemovedFromAtlasDefault(void* ptr)
{
	return static_cast<QSGDynamicTexture*>(ptr)->QSGDynamicTexture::removedFromAtlas();
}

int QSGDynamicTexture_TextureId(void* ptr)
{
	return static_cast<QSGDynamicTexture*>(ptr)->textureId();
}



void* QSGDynamicTexture_TextureSize(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QSGDynamicTexture*>(ptr)->textureSize()).width(), static_cast<QSize>(static_cast<QSGDynamicTexture*>(ptr)->textureSize()).height());
}



void QSGDynamicTexture_TimerEvent(void* ptr, void* event)
{
	static_cast<QSGDynamicTexture*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QSGDynamicTexture_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QSGDynamicTexture*>(ptr)->QSGDynamicTexture::timerEvent(static_cast<QTimerEvent*>(event));
}

void QSGDynamicTexture_ChildEvent(void* ptr, void* event)
{
	static_cast<QSGDynamicTexture*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QSGDynamicTexture_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QSGDynamicTexture*>(ptr)->QSGDynamicTexture::childEvent(static_cast<QChildEvent*>(event));
}

void QSGDynamicTexture_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QSGDynamicTexture*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSGDynamicTexture_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QSGDynamicTexture*>(ptr)->QSGDynamicTexture::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSGDynamicTexture_CustomEvent(void* ptr, void* event)
{
	static_cast<QSGDynamicTexture*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QSGDynamicTexture_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QSGDynamicTexture*>(ptr)->QSGDynamicTexture::customEvent(static_cast<QEvent*>(event));
}

void QSGDynamicTexture_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QSGDynamicTexture*>(ptr), "deleteLater");
}

void QSGDynamicTexture_DeleteLaterDefault(void* ptr)
{
	static_cast<QSGDynamicTexture*>(ptr)->QSGDynamicTexture::deleteLater();
}

void QSGDynamicTexture_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QSGDynamicTexture*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSGDynamicTexture_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QSGDynamicTexture*>(ptr)->QSGDynamicTexture::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QSGDynamicTexture_Event(void* ptr, void* e)
{
	return static_cast<QSGDynamicTexture*>(ptr)->event(static_cast<QEvent*>(e));
}

int QSGDynamicTexture_EventDefault(void* ptr, void* e)
{
	return static_cast<QSGDynamicTexture*>(ptr)->QSGDynamicTexture::event(static_cast<QEvent*>(e));
}

int QSGDynamicTexture_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QSGDynamicTexture*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QSGDynamicTexture_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QSGDynamicTexture*>(ptr)->QSGDynamicTexture::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QSGDynamicTexture_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QSGDynamicTexture*>(ptr)->metaObject());
}

void* QSGDynamicTexture_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QSGDynamicTexture*>(ptr)->QSGDynamicTexture::metaObject());
}

void* QSGEngine_NewQSGEngine(void* parent)
{
	return new QSGEngine(static_cast<QObject*>(parent));
}

void* QSGEngine_CreateRenderer(void* ptr)
{
	return static_cast<QSGEngine*>(ptr)->createRenderer();
}

void* QSGEngine_CreateTextureFromImage(void* ptr, void* image, int options)
{
	return static_cast<QSGEngine*>(ptr)->createTextureFromImage(*static_cast<QImage*>(image), static_cast<QSGEngine::CreateTextureOption>(options));
}

void QSGEngine_Initialize(void* ptr, void* context)
{
	static_cast<QSGEngine*>(ptr)->initialize(static_cast<QOpenGLContext*>(context));
}

void QSGEngine_Invalidate(void* ptr)
{
	static_cast<QSGEngine*>(ptr)->invalidate();
}

void QSGEngine_DestroyQSGEngine(void* ptr)
{
	static_cast<QSGEngine*>(ptr)->~QSGEngine();
}

void QSGEngine_TimerEvent(void* ptr, void* event)
{
	static_cast<QSGEngine*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QSGEngine_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QSGEngine*>(ptr)->QSGEngine::timerEvent(static_cast<QTimerEvent*>(event));
}

void QSGEngine_ChildEvent(void* ptr, void* event)
{
	static_cast<QSGEngine*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QSGEngine_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QSGEngine*>(ptr)->QSGEngine::childEvent(static_cast<QChildEvent*>(event));
}

void QSGEngine_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QSGEngine*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSGEngine_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QSGEngine*>(ptr)->QSGEngine::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSGEngine_CustomEvent(void* ptr, void* event)
{
	static_cast<QSGEngine*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QSGEngine_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QSGEngine*>(ptr)->QSGEngine::customEvent(static_cast<QEvent*>(event));
}

void QSGEngine_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QSGEngine*>(ptr), "deleteLater");
}

void QSGEngine_DeleteLaterDefault(void* ptr)
{
	static_cast<QSGEngine*>(ptr)->QSGEngine::deleteLater();
}

void QSGEngine_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QSGEngine*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSGEngine_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QSGEngine*>(ptr)->QSGEngine::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QSGEngine_Event(void* ptr, void* e)
{
	return static_cast<QSGEngine*>(ptr)->event(static_cast<QEvent*>(e));
}

int QSGEngine_EventDefault(void* ptr, void* e)
{
	return static_cast<QSGEngine*>(ptr)->QSGEngine::event(static_cast<QEvent*>(e));
}

int QSGEngine_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QSGEngine*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QSGEngine_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QSGEngine*>(ptr)->QSGEngine::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QSGEngine_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QSGEngine*>(ptr)->metaObject());
}

void* QSGEngine_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QSGEngine*>(ptr)->QSGEngine::metaObject());
}

void* QSGFlatColorMaterial_NewQSGFlatColorMaterial()
{
	return new QSGFlatColorMaterial();
}

void* QSGFlatColorMaterial_Color(void* ptr)
{
	return new QColor(static_cast<QSGFlatColorMaterial*>(ptr)->color());
}

void QSGFlatColorMaterial_SetColor(void* ptr, void* color)
{
	static_cast<QSGFlatColorMaterial*>(ptr)->setColor(*static_cast<QColor*>(color));
}

int QSGFlatColorMaterial_Compare(void* ptr, void* other)
{
	return static_cast<QSGFlatColorMaterial*>(ptr)->compare(static_cast<QSGMaterial*>(other));
}

int QSGFlatColorMaterial_CompareDefault(void* ptr, void* other)
{
	return static_cast<QSGFlatColorMaterial*>(ptr)->QSGFlatColorMaterial::compare(static_cast<QSGMaterial*>(other));
}

void* QSGFlatColorMaterial_CreateShader(void* ptr)
{
	return static_cast<QSGFlatColorMaterial*>(ptr)->createShader();
}

void* QSGFlatColorMaterial_CreateShaderDefault(void* ptr)
{
	return static_cast<QSGFlatColorMaterial*>(ptr)->QSGFlatColorMaterial::createShader();
}

void* QSGFlatColorMaterial_Type(void* ptr)
{
	return static_cast<QSGFlatColorMaterial*>(ptr)->type();
}

void* QSGFlatColorMaterial_TypeDefault(void* ptr)
{
	return static_cast<QSGFlatColorMaterial*>(ptr)->QSGFlatColorMaterial::type();
}

class MyQSGGeometry: public QSGGeometry
{
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
};

void QSGGeometry_Allocate(void* ptr, int vertexCount, int indexCount)
{
	static_cast<QSGGeometry*>(ptr)->allocate(vertexCount, indexCount);
}

int QSGGeometry_AttributeCount(void* ptr)
{
	return static_cast<QSGGeometry*>(ptr)->attributeCount();
}

int QSGGeometry_IndexCount(void* ptr)
{
	return static_cast<QSGGeometry*>(ptr)->indexCount();
}

void* QSGGeometry_IndexData(void* ptr)
{
	return static_cast<QSGGeometry*>(ptr)->indexData();
}

void* QSGGeometry_IndexData2(void* ptr)
{
	return const_cast<void*>(static_cast<QSGGeometry*>(ptr)->indexData());
}

int QSGGeometry_IndexDataPattern(void* ptr)
{
	return static_cast<QSGGeometry*>(ptr)->indexDataPattern();
}

int QSGGeometry_IndexType(void* ptr)
{
	return static_cast<QSGGeometry*>(ptr)->indexType();
}

void QSGGeometry_MarkIndexDataDirty(void* ptr)
{
	static_cast<QSGGeometry*>(ptr)->markIndexDataDirty();
}

void QSGGeometry_MarkVertexDataDirty(void* ptr)
{
	static_cast<QSGGeometry*>(ptr)->markVertexDataDirty();
}

void QSGGeometry_SetIndexDataPattern(void* ptr, int p)
{
	static_cast<QSGGeometry*>(ptr)->setIndexDataPattern(static_cast<QSGGeometry::DataPattern>(p));
}

void QSGGeometry_SetVertexDataPattern(void* ptr, int p)
{
	static_cast<QSGGeometry*>(ptr)->setVertexDataPattern(static_cast<QSGGeometry::DataPattern>(p));
}

int QSGGeometry_SizeOfIndex(void* ptr)
{
	return static_cast<QSGGeometry*>(ptr)->sizeOfIndex();
}

int QSGGeometry_SizeOfVertex(void* ptr)
{
	return static_cast<QSGGeometry*>(ptr)->sizeOfVertex();
}

void QSGGeometry_QSGGeometry_UpdateRectGeometry(void* g, void* rect)
{
	QSGGeometry::updateRectGeometry(static_cast<QSGGeometry*>(g), *static_cast<QRectF*>(rect));
}

void QSGGeometry_QSGGeometry_UpdateTexturedRectGeometry(void* g, void* rect, void* textureRect)
{
	QSGGeometry::updateTexturedRectGeometry(static_cast<QSGGeometry*>(g), *static_cast<QRectF*>(rect), *static_cast<QRectF*>(textureRect));
}

int QSGGeometry_VertexCount(void* ptr)
{
	return static_cast<QSGGeometry*>(ptr)->vertexCount();
}

void* QSGGeometry_VertexData(void* ptr)
{
	return static_cast<QSGGeometry*>(ptr)->vertexData();
}

void* QSGGeometry_VertexData2(void* ptr)
{
	return const_cast<void*>(static_cast<QSGGeometry*>(ptr)->vertexData());
}

int QSGGeometry_VertexDataPattern(void* ptr)
{
	return static_cast<QSGGeometry*>(ptr)->vertexDataPattern();
}

void QSGGeometry_DestroyQSGGeometry(void* ptr)
{
	static_cast<QSGGeometry*>(ptr)->~QSGGeometry();
}

char* QSGGeometry_ObjectNameAbs(void* ptr)
{
	if (dynamic_cast<MyQSGGeometry*>(static_cast<QSGGeometry*>(ptr))) {
		return static_cast<MyQSGGeometry*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QSGGeometry_BASE").toUtf8().data();
}

void QSGGeometry_SetObjectNameAbs(void* ptr, char* name)
{
	if (dynamic_cast<MyQSGGeometry*>(static_cast<QSGGeometry*>(ptr))) {
		static_cast<MyQSGGeometry*>(ptr)->setObjectNameAbs(QString(name));
	}
}

void* QSGGeometryNode_NewQSGGeometryNode()
{
	return new QSGGeometryNode();
}

void* QSGGeometryNode_Material(void* ptr)
{
	return static_cast<QSGGeometryNode*>(ptr)->material();
}

void* QSGGeometryNode_OpaqueMaterial(void* ptr)
{
	return static_cast<QSGGeometryNode*>(ptr)->opaqueMaterial();
}

void QSGGeometryNode_SetMaterial(void* ptr, void* material)
{
	static_cast<QSGGeometryNode*>(ptr)->setMaterial(static_cast<QSGMaterial*>(material));
}

void QSGGeometryNode_SetOpaqueMaterial(void* ptr, void* material)
{
	static_cast<QSGGeometryNode*>(ptr)->setOpaqueMaterial(static_cast<QSGMaterial*>(material));
}

void QSGGeometryNode_DestroyQSGGeometryNode(void* ptr)
{
	static_cast<QSGGeometryNode*>(ptr)->~QSGGeometryNode();
}

int QSGGeometryNode_IsSubtreeBlocked(void* ptr)
{
	return static_cast<QSGGeometryNode*>(ptr)->isSubtreeBlocked();
}

int QSGGeometryNode_IsSubtreeBlockedDefault(void* ptr)
{
	return static_cast<QSGGeometryNode*>(ptr)->QSGGeometryNode::isSubtreeBlocked();
}

void QSGGeometryNode_Preprocess(void* ptr)
{
	static_cast<QSGGeometryNode*>(ptr)->preprocess();
}

void QSGGeometryNode_PreprocessDefault(void* ptr)
{
	static_cast<QSGGeometryNode*>(ptr)->QSGGeometryNode::preprocess();
}

class MyQSGMaterial: public QSGMaterial
{
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	int compare(const QSGMaterial * other) const { return callbackQSGMaterial_Compare(const_cast<MyQSGMaterial*>(this), this->objectNameAbs().toUtf8().data(), const_cast<QSGMaterial*>(other)); };
	QSGMaterialShader * createShader() const { return static_cast<QSGMaterialShader*>(callbackQSGMaterial_CreateShader(const_cast<MyQSGMaterial*>(this), this->objectNameAbs().toUtf8().data())); };
	QSGMaterialType * type() const { return static_cast<QSGMaterialType*>(callbackQSGMaterial_Type(const_cast<MyQSGMaterial*>(this), this->objectNameAbs().toUtf8().data())); };
};

int QSGMaterial_Compare(void* ptr, void* other)
{
	return static_cast<QSGMaterial*>(ptr)->compare(static_cast<QSGMaterial*>(other));
}

int QSGMaterial_CompareDefault(void* ptr, void* other)
{
	return static_cast<QSGMaterial*>(ptr)->QSGMaterial::compare(static_cast<QSGMaterial*>(other));
}

void* QSGMaterial_CreateShader(void* ptr)
{
	return static_cast<QSGMaterial*>(ptr)->createShader();
}

int QSGMaterial_Flags(void* ptr)
{
	return static_cast<QSGMaterial*>(ptr)->flags();
}

void QSGMaterial_SetFlag(void* ptr, int flags, int on)
{
	static_cast<QSGMaterial*>(ptr)->setFlag(static_cast<QSGMaterial::Flag>(flags), on != 0);
}

void* QSGMaterial_Type(void* ptr)
{
	return static_cast<QSGMaterial*>(ptr)->type();
}

char* QSGMaterial_ObjectNameAbs(void* ptr)
{
	if (dynamic_cast<MyQSGMaterial*>(static_cast<QSGMaterial*>(ptr))) {
		return static_cast<MyQSGMaterial*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QSGMaterial_BASE").toUtf8().data();
}

void QSGMaterial_SetObjectNameAbs(void* ptr, char* name)
{
	if (dynamic_cast<MyQSGMaterial*>(static_cast<QSGMaterial*>(ptr))) {
		static_cast<MyQSGMaterial*>(ptr)->setObjectNameAbs(QString(name));
	}
}

class MyQSGMaterialShader: public QSGMaterialShader
{
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	const char * fragmentShader() const { return const_cast<const char*>(callbackQSGMaterialShader_FragmentShader(const_cast<MyQSGMaterialShader*>(this), this->objectNameAbs().toUtf8().data())); };
	const char * vertexShader() const { return const_cast<const char*>(callbackQSGMaterialShader_VertexShader(const_cast<MyQSGMaterialShader*>(this), this->objectNameAbs().toUtf8().data())); };
	void activate() { callbackQSGMaterialShader_Activate(this, this->objectNameAbs().toUtf8().data()); };
	void deactivate() { callbackQSGMaterialShader_Deactivate(this, this->objectNameAbs().toUtf8().data()); };
	void compile() { callbackQSGMaterialShader_Compile(this, this->objectNameAbs().toUtf8().data()); };
	void initialize() { callbackQSGMaterialShader_Initialize(this, this->objectNameAbs().toUtf8().data()); };
};

char* QSGMaterialShader_FragmentShader(void* ptr)
{
	return QString(static_cast<QSGMaterialShader*>(ptr)->fragmentShader()).toUtf8().data();
}

char* QSGMaterialShader_FragmentShaderDefault(void* ptr)
{
	return QString(static_cast<QSGMaterialShader*>(ptr)->QSGMaterialShader::fragmentShader()).toUtf8().data();
}

char* QSGMaterialShader_VertexShader(void* ptr)
{
	return QString(static_cast<QSGMaterialShader*>(ptr)->vertexShader()).toUtf8().data();
}

char* QSGMaterialShader_VertexShaderDefault(void* ptr)
{
	return QString(static_cast<QSGMaterialShader*>(ptr)->QSGMaterialShader::vertexShader()).toUtf8().data();
}

void QSGMaterialShader_Activate(void* ptr)
{
	static_cast<QSGMaterialShader*>(ptr)->activate();
}

void QSGMaterialShader_ActivateDefault(void* ptr)
{
	static_cast<QSGMaterialShader*>(ptr)->QSGMaterialShader::activate();
}

void QSGMaterialShader_Deactivate(void* ptr)
{
	static_cast<QSGMaterialShader*>(ptr)->deactivate();
}

void QSGMaterialShader_DeactivateDefault(void* ptr)
{
	static_cast<QSGMaterialShader*>(ptr)->QSGMaterialShader::deactivate();
}

void* QSGMaterialShader_Program(void* ptr)
{
	return static_cast<QSGMaterialShader*>(ptr)->program();
}

void QSGMaterialShader_Compile(void* ptr)
{
	static_cast<QSGMaterialShader*>(ptr)->compile();
}

void QSGMaterialShader_CompileDefault(void* ptr)
{
	static_cast<QSGMaterialShader*>(ptr)->QSGMaterialShader::compile();
}

void QSGMaterialShader_Initialize(void* ptr)
{
	static_cast<QSGMaterialShader*>(ptr)->initialize();
}

void QSGMaterialShader_InitializeDefault(void* ptr)
{
	static_cast<QSGMaterialShader*>(ptr)->QSGMaterialShader::initialize();
}

void QSGMaterialShader_SetShaderSourceFile(void* ptr, int ty, char* sourceFile)
{
	static_cast<QSGMaterialShader*>(ptr)->setShaderSourceFile(static_cast<QOpenGLShader::ShaderTypeBit>(ty), QString(sourceFile));
}

void QSGMaterialShader_SetShaderSourceFiles(void* ptr, int ty, char* sourceFiles)
{
	static_cast<QSGMaterialShader*>(ptr)->setShaderSourceFiles(static_cast<QOpenGLShader::ShaderTypeBit>(ty), QString(sourceFiles).split("|", QString::SkipEmptyParts));
}

char* QSGMaterialShader_ObjectNameAbs(void* ptr)
{
	if (dynamic_cast<MyQSGMaterialShader*>(static_cast<QSGMaterialShader*>(ptr))) {
		return static_cast<MyQSGMaterialShader*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QSGMaterialShader_BASE").toUtf8().data();
}

void QSGMaterialShader_SetObjectNameAbs(void* ptr, char* name)
{
	if (dynamic_cast<MyQSGMaterialShader*>(static_cast<QSGMaterialShader*>(ptr))) {
		static_cast<MyQSGMaterialShader*>(ptr)->setObjectNameAbs(QString(name));
	}
}

class MyQSGNode: public QSGNode
{
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	MyQSGNode() : QSGNode() {};
	bool isSubtreeBlocked() const { return callbackQSGNode_IsSubtreeBlocked(const_cast<MyQSGNode*>(this), this->objectNameAbs().toUtf8().data()) != 0; };
	void preprocess() { callbackQSGNode_Preprocess(this, this->objectNameAbs().toUtf8().data()); };
};

void* QSGNode_ChildAtIndex(void* ptr, int i)
{
	return static_cast<QSGNode*>(ptr)->childAtIndex(i);
}

int QSGNode_ChildCount(void* ptr)
{
	return static_cast<QSGNode*>(ptr)->childCount();
}

void* QSGNode_NewQSGNode()
{
	return new MyQSGNode();
}

void QSGNode_AppendChildNode(void* ptr, void* node)
{
	static_cast<QSGNode*>(ptr)->appendChildNode(static_cast<QSGNode*>(node));
}

void* QSGNode_FirstChild(void* ptr)
{
	return static_cast<QSGNode*>(ptr)->firstChild();
}

int QSGNode_Flags(void* ptr)
{
	return static_cast<QSGNode*>(ptr)->flags();
}

void QSGNode_InsertChildNodeAfter(void* ptr, void* node, void* after)
{
	static_cast<QSGNode*>(ptr)->insertChildNodeAfter(static_cast<QSGNode*>(node), static_cast<QSGNode*>(after));
}

void QSGNode_InsertChildNodeBefore(void* ptr, void* node, void* before)
{
	static_cast<QSGNode*>(ptr)->insertChildNodeBefore(static_cast<QSGNode*>(node), static_cast<QSGNode*>(before));
}

int QSGNode_IsSubtreeBlocked(void* ptr)
{
	return static_cast<QSGNode*>(ptr)->isSubtreeBlocked();
}

int QSGNode_IsSubtreeBlockedDefault(void* ptr)
{
	return static_cast<QSGNode*>(ptr)->QSGNode::isSubtreeBlocked();
}

void* QSGNode_LastChild(void* ptr)
{
	return static_cast<QSGNode*>(ptr)->lastChild();
}

void QSGNode_MarkDirty(void* ptr, int bits)
{
	static_cast<QSGNode*>(ptr)->markDirty(static_cast<QSGNode::DirtyStateBit>(bits));
}

void* QSGNode_NextSibling(void* ptr)
{
	return static_cast<QSGNode*>(ptr)->nextSibling();
}

void* QSGNode_Parent(void* ptr)
{
	return static_cast<QSGNode*>(ptr)->parent();
}

void QSGNode_PrependChildNode(void* ptr, void* node)
{
	static_cast<QSGNode*>(ptr)->prependChildNode(static_cast<QSGNode*>(node));
}

void QSGNode_Preprocess(void* ptr)
{
	static_cast<QSGNode*>(ptr)->preprocess();
}

void QSGNode_PreprocessDefault(void* ptr)
{
	static_cast<QSGNode*>(ptr)->QSGNode::preprocess();
}

void* QSGNode_PreviousSibling(void* ptr)
{
	return static_cast<QSGNode*>(ptr)->previousSibling();
}

void QSGNode_RemoveAllChildNodes(void* ptr)
{
	static_cast<QSGNode*>(ptr)->removeAllChildNodes();
}

void QSGNode_RemoveChildNode(void* ptr, void* node)
{
	static_cast<QSGNode*>(ptr)->removeChildNode(static_cast<QSGNode*>(node));
}

void QSGNode_SetFlag(void* ptr, int f, int enabled)
{
	static_cast<QSGNode*>(ptr)->setFlag(static_cast<QSGNode::Flag>(f), enabled != 0);
}

void QSGNode_SetFlags(void* ptr, int f, int enabled)
{
	static_cast<QSGNode*>(ptr)->setFlags(static_cast<QSGNode::Flag>(f), enabled != 0);
}

int QSGNode_Type(void* ptr)
{
	return static_cast<QSGNode*>(ptr)->type();
}

void QSGNode_DestroyQSGNode(void* ptr)
{
	static_cast<QSGNode*>(ptr)->~QSGNode();
}

char* QSGNode_ObjectNameAbs(void* ptr)
{
	if (dynamic_cast<MyQSGNode*>(static_cast<QSGNode*>(ptr))) {
		return static_cast<MyQSGNode*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QSGNode_BASE").toUtf8().data();
}

void QSGNode_SetObjectNameAbs(void* ptr, char* name)
{
	if (dynamic_cast<MyQSGNode*>(static_cast<QSGNode*>(ptr))) {
		static_cast<MyQSGNode*>(ptr)->setObjectNameAbs(QString(name));
	}
}

void* QSGOpacityNode_NewQSGOpacityNode()
{
	return new QSGOpacityNode();
}

double QSGOpacityNode_Opacity(void* ptr)
{
	return static_cast<double>(static_cast<QSGOpacityNode*>(ptr)->opacity());
}

void QSGOpacityNode_SetOpacity(void* ptr, double opacity)
{
	static_cast<QSGOpacityNode*>(ptr)->setOpacity(static_cast<double>(opacity));
}

void QSGOpacityNode_DestroyQSGOpacityNode(void* ptr)
{
	static_cast<QSGOpacityNode*>(ptr)->~QSGOpacityNode();
}

int QSGOpacityNode_IsSubtreeBlocked(void* ptr)
{
	return static_cast<QSGOpacityNode*>(ptr)->isSubtreeBlocked();
}

int QSGOpacityNode_IsSubtreeBlockedDefault(void* ptr)
{
	return static_cast<QSGOpacityNode*>(ptr)->QSGOpacityNode::isSubtreeBlocked();
}

void QSGOpacityNode_Preprocess(void* ptr)
{
	static_cast<QSGOpacityNode*>(ptr)->preprocess();
}

void QSGOpacityNode_PreprocessDefault(void* ptr)
{
	static_cast<QSGOpacityNode*>(ptr)->QSGOpacityNode::preprocess();
}

void* QSGOpaqueTextureMaterial_NewQSGOpaqueTextureMaterial()
{
	return new QSGOpaqueTextureMaterial();
}

int QSGOpaqueTextureMaterial_Filtering(void* ptr)
{
	return static_cast<QSGOpaqueTextureMaterial*>(ptr)->filtering();
}

int QSGOpaqueTextureMaterial_HorizontalWrapMode(void* ptr)
{
	return static_cast<QSGOpaqueTextureMaterial*>(ptr)->horizontalWrapMode();
}

int QSGOpaqueTextureMaterial_MipmapFiltering(void* ptr)
{
	return static_cast<QSGOpaqueTextureMaterial*>(ptr)->mipmapFiltering();
}

void QSGOpaqueTextureMaterial_SetFiltering(void* ptr, int filtering)
{
	static_cast<QSGOpaqueTextureMaterial*>(ptr)->setFiltering(static_cast<QSGTexture::Filtering>(filtering));
}

void QSGOpaqueTextureMaterial_SetHorizontalWrapMode(void* ptr, int mode)
{
	static_cast<QSGOpaqueTextureMaterial*>(ptr)->setHorizontalWrapMode(static_cast<QSGTexture::WrapMode>(mode));
}

void QSGOpaqueTextureMaterial_SetMipmapFiltering(void* ptr, int filtering)
{
	static_cast<QSGOpaqueTextureMaterial*>(ptr)->setMipmapFiltering(static_cast<QSGTexture::Filtering>(filtering));
}

void QSGOpaqueTextureMaterial_SetTexture(void* ptr, void* texture)
{
	static_cast<QSGOpaqueTextureMaterial*>(ptr)->setTexture(static_cast<QSGTexture*>(texture));
}

void QSGOpaqueTextureMaterial_SetVerticalWrapMode(void* ptr, int mode)
{
	static_cast<QSGOpaqueTextureMaterial*>(ptr)->setVerticalWrapMode(static_cast<QSGTexture::WrapMode>(mode));
}

void* QSGOpaqueTextureMaterial_Texture(void* ptr)
{
	return static_cast<QSGOpaqueTextureMaterial*>(ptr)->texture();
}

int QSGOpaqueTextureMaterial_VerticalWrapMode(void* ptr)
{
	return static_cast<QSGOpaqueTextureMaterial*>(ptr)->verticalWrapMode();
}

int QSGOpaqueTextureMaterial_Compare(void* ptr, void* other)
{
	return static_cast<QSGOpaqueTextureMaterial*>(ptr)->compare(static_cast<QSGMaterial*>(other));
}

int QSGOpaqueTextureMaterial_CompareDefault(void* ptr, void* other)
{
	return static_cast<QSGOpaqueTextureMaterial*>(ptr)->QSGOpaqueTextureMaterial::compare(static_cast<QSGMaterial*>(other));
}

void* QSGOpaqueTextureMaterial_CreateShader(void* ptr)
{
	return static_cast<QSGOpaqueTextureMaterial*>(ptr)->createShader();
}

void* QSGOpaqueTextureMaterial_CreateShaderDefault(void* ptr)
{
	return static_cast<QSGOpaqueTextureMaterial*>(ptr)->QSGOpaqueTextureMaterial::createShader();
}

void* QSGOpaqueTextureMaterial_Type(void* ptr)
{
	return static_cast<QSGOpaqueTextureMaterial*>(ptr)->type();
}

void* QSGOpaqueTextureMaterial_TypeDefault(void* ptr)
{
	return static_cast<QSGOpaqueTextureMaterial*>(ptr)->QSGOpaqueTextureMaterial::type();
}

void* QSGSimpleRectNode_NewQSGSimpleRectNode2()
{
	return new QSGSimpleRectNode();
}

void* QSGSimpleRectNode_NewQSGSimpleRectNode(void* rect, void* color)
{
	return new QSGSimpleRectNode(*static_cast<QRectF*>(rect), *static_cast<QColor*>(color));
}

void* QSGSimpleRectNode_Color(void* ptr)
{
	return new QColor(static_cast<QSGSimpleRectNode*>(ptr)->color());
}

void* QSGSimpleRectNode_Rect(void* ptr)
{
	return new QRectF(static_cast<QRectF>(static_cast<QSGSimpleRectNode*>(ptr)->rect()).x(), static_cast<QRectF>(static_cast<QSGSimpleRectNode*>(ptr)->rect()).y(), static_cast<QRectF>(static_cast<QSGSimpleRectNode*>(ptr)->rect()).width(), static_cast<QRectF>(static_cast<QSGSimpleRectNode*>(ptr)->rect()).height());
}

void QSGSimpleRectNode_SetColor(void* ptr, void* color)
{
	static_cast<QSGSimpleRectNode*>(ptr)->setColor(*static_cast<QColor*>(color));
}

void QSGSimpleRectNode_SetRect(void* ptr, void* rect)
{
	static_cast<QSGSimpleRectNode*>(ptr)->setRect(*static_cast<QRectF*>(rect));
}

void QSGSimpleRectNode_SetRect2(void* ptr, double x, double y, double w, double h)
{
	static_cast<QSGSimpleRectNode*>(ptr)->setRect(static_cast<double>(x), static_cast<double>(y), static_cast<double>(w), static_cast<double>(h));
}

int QSGSimpleRectNode_IsSubtreeBlocked(void* ptr)
{
	return static_cast<QSGSimpleRectNode*>(ptr)->isSubtreeBlocked();
}

int QSGSimpleRectNode_IsSubtreeBlockedDefault(void* ptr)
{
	return static_cast<QSGSimpleRectNode*>(ptr)->QSGSimpleRectNode::isSubtreeBlocked();
}

void QSGSimpleRectNode_Preprocess(void* ptr)
{
	static_cast<QSGSimpleRectNode*>(ptr)->preprocess();
}

void QSGSimpleRectNode_PreprocessDefault(void* ptr)
{
	static_cast<QSGSimpleRectNode*>(ptr)->QSGSimpleRectNode::preprocess();
}

void* QSGSimpleTextureNode_NewQSGSimpleTextureNode()
{
	return new QSGSimpleTextureNode();
}

int QSGSimpleTextureNode_Filtering(void* ptr)
{
	return static_cast<QSGSimpleTextureNode*>(ptr)->filtering();
}

int QSGSimpleTextureNode_OwnsTexture(void* ptr)
{
	return static_cast<QSGSimpleTextureNode*>(ptr)->ownsTexture();
}

void* QSGSimpleTextureNode_Rect(void* ptr)
{
	return new QRectF(static_cast<QRectF>(static_cast<QSGSimpleTextureNode*>(ptr)->rect()).x(), static_cast<QRectF>(static_cast<QSGSimpleTextureNode*>(ptr)->rect()).y(), static_cast<QRectF>(static_cast<QSGSimpleTextureNode*>(ptr)->rect()).width(), static_cast<QRectF>(static_cast<QSGSimpleTextureNode*>(ptr)->rect()).height());
}

void QSGSimpleTextureNode_SetFiltering(void* ptr, int filtering)
{
	static_cast<QSGSimpleTextureNode*>(ptr)->setFiltering(static_cast<QSGTexture::Filtering>(filtering));
}

void QSGSimpleTextureNode_SetOwnsTexture(void* ptr, int owns)
{
	static_cast<QSGSimpleTextureNode*>(ptr)->setOwnsTexture(owns != 0);
}

void QSGSimpleTextureNode_SetRect(void* ptr, void* r)
{
	static_cast<QSGSimpleTextureNode*>(ptr)->setRect(*static_cast<QRectF*>(r));
}

void QSGSimpleTextureNode_SetRect2(void* ptr, double x, double y, double w, double h)
{
	static_cast<QSGSimpleTextureNode*>(ptr)->setRect(static_cast<double>(x), static_cast<double>(y), static_cast<double>(w), static_cast<double>(h));
}

void QSGSimpleTextureNode_SetSourceRect(void* ptr, void* r)
{
	static_cast<QSGSimpleTextureNode*>(ptr)->setSourceRect(*static_cast<QRectF*>(r));
}

void QSGSimpleTextureNode_SetSourceRect2(void* ptr, double x, double y, double w, double h)
{
	static_cast<QSGSimpleTextureNode*>(ptr)->setSourceRect(static_cast<double>(x), static_cast<double>(y), static_cast<double>(w), static_cast<double>(h));
}

void QSGSimpleTextureNode_SetTexture(void* ptr, void* texture)
{
	static_cast<QSGSimpleTextureNode*>(ptr)->setTexture(static_cast<QSGTexture*>(texture));
}

void QSGSimpleTextureNode_SetTextureCoordinatesTransform(void* ptr, int mode)
{
	static_cast<QSGSimpleTextureNode*>(ptr)->setTextureCoordinatesTransform(static_cast<QSGSimpleTextureNode::TextureCoordinatesTransformFlag>(mode));
}

void* QSGSimpleTextureNode_SourceRect(void* ptr)
{
	return new QRectF(static_cast<QRectF>(static_cast<QSGSimpleTextureNode*>(ptr)->sourceRect()).x(), static_cast<QRectF>(static_cast<QSGSimpleTextureNode*>(ptr)->sourceRect()).y(), static_cast<QRectF>(static_cast<QSGSimpleTextureNode*>(ptr)->sourceRect()).width(), static_cast<QRectF>(static_cast<QSGSimpleTextureNode*>(ptr)->sourceRect()).height());
}

void* QSGSimpleTextureNode_Texture(void* ptr)
{
	return static_cast<QSGSimpleTextureNode*>(ptr)->texture();
}

int QSGSimpleTextureNode_TextureCoordinatesTransform(void* ptr)
{
	return static_cast<QSGSimpleTextureNode*>(ptr)->textureCoordinatesTransform();
}

void QSGSimpleTextureNode_DestroyQSGSimpleTextureNode(void* ptr)
{
	static_cast<QSGSimpleTextureNode*>(ptr)->~QSGSimpleTextureNode();
}

int QSGSimpleTextureNode_IsSubtreeBlocked(void* ptr)
{
	return static_cast<QSGSimpleTextureNode*>(ptr)->isSubtreeBlocked();
}

int QSGSimpleTextureNode_IsSubtreeBlockedDefault(void* ptr)
{
	return static_cast<QSGSimpleTextureNode*>(ptr)->QSGSimpleTextureNode::isSubtreeBlocked();
}

void QSGSimpleTextureNode_Preprocess(void* ptr)
{
	static_cast<QSGSimpleTextureNode*>(ptr)->preprocess();
}

void QSGSimpleTextureNode_PreprocessDefault(void* ptr)
{
	static_cast<QSGSimpleTextureNode*>(ptr)->QSGSimpleTextureNode::preprocess();
}

class MyQSGTexture: public QSGTexture
{
public:
	MyQSGTexture() : QSGTexture() {};
	void bind() { callbackQSGTexture_Bind(this, this->objectName().toUtf8().data()); };
	bool hasAlphaChannel() const { return callbackQSGTexture_HasAlphaChannel(const_cast<MyQSGTexture*>(this), this->objectName().toUtf8().data()) != 0; };
	bool hasMipmaps() const { return callbackQSGTexture_HasMipmaps(const_cast<MyQSGTexture*>(this), this->objectName().toUtf8().data()) != 0; };
	bool isAtlasTexture() const { return callbackQSGTexture_IsAtlasTexture(const_cast<MyQSGTexture*>(this), this->objectName().toUtf8().data()) != 0; };
	QRectF normalizedTextureSubRect() const { return *static_cast<QRectF*>(callbackQSGTexture_NormalizedTextureSubRect(const_cast<MyQSGTexture*>(this), this->objectName().toUtf8().data())); };
	QSGTexture * removedFromAtlas() const { return static_cast<QSGTexture*>(callbackQSGTexture_RemovedFromAtlas(const_cast<MyQSGTexture*>(this), this->objectName().toUtf8().data())); };
	int textureId() const { return callbackQSGTexture_TextureId(const_cast<MyQSGTexture*>(this), this->objectName().toUtf8().data()); };
	QSize textureSize() const { return *static_cast<QSize*>(callbackQSGTexture_TextureSize(const_cast<MyQSGTexture*>(this), this->objectName().toUtf8().data())); };
	void timerEvent(QTimerEvent * event) { callbackQSGTexture_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQSGTexture_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSGTexture_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQSGTexture_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQSGTexture_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSGTexture_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQSGTexture_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSGTexture_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSGTexture_MetaObject(const_cast<MyQSGTexture*>(this), this->objectName().toUtf8().data())); };
};

void* QSGTexture_NewQSGTexture()
{
	return new MyQSGTexture();
}

void QSGTexture_Bind(void* ptr)
{
	static_cast<QSGTexture*>(ptr)->bind();
}

void* QSGTexture_ConvertToNormalizedSourceRect(void* ptr, void* rect)
{
	return new QRectF(static_cast<QRectF>(static_cast<QSGTexture*>(ptr)->convertToNormalizedSourceRect(*static_cast<QRectF*>(rect))).x(), static_cast<QRectF>(static_cast<QSGTexture*>(ptr)->convertToNormalizedSourceRect(*static_cast<QRectF*>(rect))).y(), static_cast<QRectF>(static_cast<QSGTexture*>(ptr)->convertToNormalizedSourceRect(*static_cast<QRectF*>(rect))).width(), static_cast<QRectF>(static_cast<QSGTexture*>(ptr)->convertToNormalizedSourceRect(*static_cast<QRectF*>(rect))).height());
}

int QSGTexture_Filtering(void* ptr)
{
	return static_cast<QSGTexture*>(ptr)->filtering();
}

int QSGTexture_HasAlphaChannel(void* ptr)
{
	return static_cast<QSGTexture*>(ptr)->hasAlphaChannel();
}

int QSGTexture_HasMipmaps(void* ptr)
{
	return static_cast<QSGTexture*>(ptr)->hasMipmaps();
}

int QSGTexture_HorizontalWrapMode(void* ptr)
{
	return static_cast<QSGTexture*>(ptr)->horizontalWrapMode();
}

int QSGTexture_IsAtlasTexture(void* ptr)
{
	return static_cast<QSGTexture*>(ptr)->isAtlasTexture();
}

int QSGTexture_IsAtlasTextureDefault(void* ptr)
{
	return static_cast<QSGTexture*>(ptr)->QSGTexture::isAtlasTexture();
}

int QSGTexture_MipmapFiltering(void* ptr)
{
	return static_cast<QSGTexture*>(ptr)->mipmapFiltering();
}

void* QSGTexture_NormalizedTextureSubRect(void* ptr)
{
	return new QRectF(static_cast<QRectF>(static_cast<QSGTexture*>(ptr)->normalizedTextureSubRect()).x(), static_cast<QRectF>(static_cast<QSGTexture*>(ptr)->normalizedTextureSubRect()).y(), static_cast<QRectF>(static_cast<QSGTexture*>(ptr)->normalizedTextureSubRect()).width(), static_cast<QRectF>(static_cast<QSGTexture*>(ptr)->normalizedTextureSubRect()).height());
}

void* QSGTexture_NormalizedTextureSubRectDefault(void* ptr)
{
	return new QRectF(static_cast<QRectF>(static_cast<QSGTexture*>(ptr)->QSGTexture::normalizedTextureSubRect()).x(), static_cast<QRectF>(static_cast<QSGTexture*>(ptr)->QSGTexture::normalizedTextureSubRect()).y(), static_cast<QRectF>(static_cast<QSGTexture*>(ptr)->QSGTexture::normalizedTextureSubRect()).width(), static_cast<QRectF>(static_cast<QSGTexture*>(ptr)->QSGTexture::normalizedTextureSubRect()).height());
}

void* QSGTexture_RemovedFromAtlas(void* ptr)
{
	return static_cast<QSGTexture*>(ptr)->removedFromAtlas();
}

void* QSGTexture_RemovedFromAtlasDefault(void* ptr)
{
	return static_cast<QSGTexture*>(ptr)->QSGTexture::removedFromAtlas();
}

void QSGTexture_SetFiltering(void* ptr, int filter)
{
	static_cast<QSGTexture*>(ptr)->setFiltering(static_cast<QSGTexture::Filtering>(filter));
}

void QSGTexture_SetHorizontalWrapMode(void* ptr, int hwrap)
{
	static_cast<QSGTexture*>(ptr)->setHorizontalWrapMode(static_cast<QSGTexture::WrapMode>(hwrap));
}

void QSGTexture_SetMipmapFiltering(void* ptr, int filter)
{
	static_cast<QSGTexture*>(ptr)->setMipmapFiltering(static_cast<QSGTexture::Filtering>(filter));
}

void QSGTexture_SetVerticalWrapMode(void* ptr, int vwrap)
{
	static_cast<QSGTexture*>(ptr)->setVerticalWrapMode(static_cast<QSGTexture::WrapMode>(vwrap));
}

int QSGTexture_TextureId(void* ptr)
{
	return static_cast<QSGTexture*>(ptr)->textureId();
}

void* QSGTexture_TextureSize(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QSGTexture*>(ptr)->textureSize()).width(), static_cast<QSize>(static_cast<QSGTexture*>(ptr)->textureSize()).height());
}

void QSGTexture_UpdateBindOptions(void* ptr, int force)
{
	static_cast<QSGTexture*>(ptr)->updateBindOptions(force != 0);
}

int QSGTexture_VerticalWrapMode(void* ptr)
{
	return static_cast<QSGTexture*>(ptr)->verticalWrapMode();
}

void QSGTexture_DestroyQSGTexture(void* ptr)
{
	static_cast<QSGTexture*>(ptr)->~QSGTexture();
}

void QSGTexture_TimerEvent(void* ptr, void* event)
{
	static_cast<QSGTexture*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QSGTexture_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QSGTexture*>(ptr)->QSGTexture::timerEvent(static_cast<QTimerEvent*>(event));
}

void QSGTexture_ChildEvent(void* ptr, void* event)
{
	static_cast<QSGTexture*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QSGTexture_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QSGTexture*>(ptr)->QSGTexture::childEvent(static_cast<QChildEvent*>(event));
}

void QSGTexture_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QSGTexture*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSGTexture_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QSGTexture*>(ptr)->QSGTexture::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSGTexture_CustomEvent(void* ptr, void* event)
{
	static_cast<QSGTexture*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QSGTexture_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QSGTexture*>(ptr)->QSGTexture::customEvent(static_cast<QEvent*>(event));
}

void QSGTexture_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QSGTexture*>(ptr), "deleteLater");
}

void QSGTexture_DeleteLaterDefault(void* ptr)
{
	static_cast<QSGTexture*>(ptr)->QSGTexture::deleteLater();
}

void QSGTexture_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QSGTexture*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSGTexture_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QSGTexture*>(ptr)->QSGTexture::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QSGTexture_Event(void* ptr, void* e)
{
	return static_cast<QSGTexture*>(ptr)->event(static_cast<QEvent*>(e));
}

int QSGTexture_EventDefault(void* ptr, void* e)
{
	return static_cast<QSGTexture*>(ptr)->QSGTexture::event(static_cast<QEvent*>(e));
}

int QSGTexture_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QSGTexture*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QSGTexture_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QSGTexture*>(ptr)->QSGTexture::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QSGTexture_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QSGTexture*>(ptr)->metaObject());
}

void* QSGTexture_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QSGTexture*>(ptr)->QSGTexture::metaObject());
}

int QSGTextureMaterial_Compare(void* ptr, void* other)
{
	return static_cast<QSGTextureMaterial*>(ptr)->compare(static_cast<QSGMaterial*>(other));
}

int QSGTextureMaterial_CompareDefault(void* ptr, void* other)
{
	return static_cast<QSGTextureMaterial*>(ptr)->QSGTextureMaterial::compare(static_cast<QSGMaterial*>(other));
}

void* QSGTextureMaterial_CreateShader(void* ptr)
{
	return static_cast<QSGTextureMaterial*>(ptr)->createShader();
}

void* QSGTextureMaterial_CreateShaderDefault(void* ptr)
{
	return static_cast<QSGTextureMaterial*>(ptr)->QSGTextureMaterial::createShader();
}

void* QSGTextureMaterial_Type(void* ptr)
{
	return static_cast<QSGTextureMaterial*>(ptr)->type();
}

void* QSGTextureMaterial_TypeDefault(void* ptr)
{
	return static_cast<QSGTextureMaterial*>(ptr)->QSGTextureMaterial::type();
}

class MyQSGTextureProvider: public QSGTextureProvider
{
public:
	QSGTexture * texture() const { return static_cast<QSGTexture*>(callbackQSGTextureProvider_Texture(const_cast<MyQSGTextureProvider*>(this), this->objectName().toUtf8().data())); };
	void Signal_TextureChanged() { callbackQSGTextureProvider_TextureChanged(this, this->objectName().toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQSGTextureProvider_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQSGTextureProvider_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSGTextureProvider_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQSGTextureProvider_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQSGTextureProvider_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSGTextureProvider_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQSGTextureProvider_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSGTextureProvider_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSGTextureProvider_MetaObject(const_cast<MyQSGTextureProvider*>(this), this->objectName().toUtf8().data())); };
};

void* QSGTextureProvider_Texture(void* ptr)
{
	return static_cast<QSGTextureProvider*>(ptr)->texture();
}

void QSGTextureProvider_ConnectTextureChanged(void* ptr)
{
	QObject::connect(static_cast<QSGTextureProvider*>(ptr), static_cast<void (QSGTextureProvider::*)()>(&QSGTextureProvider::textureChanged), static_cast<MyQSGTextureProvider*>(ptr), static_cast<void (MyQSGTextureProvider::*)()>(&MyQSGTextureProvider::Signal_TextureChanged));
}

void QSGTextureProvider_DisconnectTextureChanged(void* ptr)
{
	QObject::disconnect(static_cast<QSGTextureProvider*>(ptr), static_cast<void (QSGTextureProvider::*)()>(&QSGTextureProvider::textureChanged), static_cast<MyQSGTextureProvider*>(ptr), static_cast<void (MyQSGTextureProvider::*)()>(&MyQSGTextureProvider::Signal_TextureChanged));
}

void QSGTextureProvider_TextureChanged(void* ptr)
{
	static_cast<QSGTextureProvider*>(ptr)->textureChanged();
}

void QSGTextureProvider_TimerEvent(void* ptr, void* event)
{
	static_cast<QSGTextureProvider*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QSGTextureProvider_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QSGTextureProvider*>(ptr)->QSGTextureProvider::timerEvent(static_cast<QTimerEvent*>(event));
}

void QSGTextureProvider_ChildEvent(void* ptr, void* event)
{
	static_cast<QSGTextureProvider*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QSGTextureProvider_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QSGTextureProvider*>(ptr)->QSGTextureProvider::childEvent(static_cast<QChildEvent*>(event));
}

void QSGTextureProvider_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QSGTextureProvider*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSGTextureProvider_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QSGTextureProvider*>(ptr)->QSGTextureProvider::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSGTextureProvider_CustomEvent(void* ptr, void* event)
{
	static_cast<QSGTextureProvider*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QSGTextureProvider_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QSGTextureProvider*>(ptr)->QSGTextureProvider::customEvent(static_cast<QEvent*>(event));
}

void QSGTextureProvider_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QSGTextureProvider*>(ptr), "deleteLater");
}

void QSGTextureProvider_DeleteLaterDefault(void* ptr)
{
	static_cast<QSGTextureProvider*>(ptr)->QSGTextureProvider::deleteLater();
}

void QSGTextureProvider_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QSGTextureProvider*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSGTextureProvider_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QSGTextureProvider*>(ptr)->QSGTextureProvider::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QSGTextureProvider_Event(void* ptr, void* e)
{
	return static_cast<QSGTextureProvider*>(ptr)->event(static_cast<QEvent*>(e));
}

int QSGTextureProvider_EventDefault(void* ptr, void* e)
{
	return static_cast<QSGTextureProvider*>(ptr)->QSGTextureProvider::event(static_cast<QEvent*>(e));
}

int QSGTextureProvider_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QSGTextureProvider*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QSGTextureProvider_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QSGTextureProvider*>(ptr)->QSGTextureProvider::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QSGTextureProvider_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QSGTextureProvider*>(ptr)->metaObject());
}

void* QSGTextureProvider_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QSGTextureProvider*>(ptr)->QSGTextureProvider::metaObject());
}

void* QSGTransformNode_NewQSGTransformNode()
{
	return new QSGTransformNode();
}

void QSGTransformNode_SetMatrix(void* ptr, void* matrix)
{
	static_cast<QSGTransformNode*>(ptr)->setMatrix(*static_cast<QMatrix4x4*>(matrix));
}

void QSGTransformNode_DestroyQSGTransformNode(void* ptr)
{
	static_cast<QSGTransformNode*>(ptr)->~QSGTransformNode();
}

int QSGTransformNode_IsSubtreeBlocked(void* ptr)
{
	return static_cast<QSGTransformNode*>(ptr)->isSubtreeBlocked();
}

int QSGTransformNode_IsSubtreeBlockedDefault(void* ptr)
{
	return static_cast<QSGTransformNode*>(ptr)->QSGTransformNode::isSubtreeBlocked();
}

void QSGTransformNode_Preprocess(void* ptr)
{
	static_cast<QSGTransformNode*>(ptr)->preprocess();
}

void QSGTransformNode_PreprocessDefault(void* ptr)
{
	static_cast<QSGTransformNode*>(ptr)->QSGTransformNode::preprocess();
}

void* QSGVertexColorMaterial_NewQSGVertexColorMaterial()
{
	return new QSGVertexColorMaterial();
}

int QSGVertexColorMaterial_Compare(void* ptr, void* other)
{
	return static_cast<QSGVertexColorMaterial*>(ptr)->compare(static_cast<QSGMaterial*>(other));
}

int QSGVertexColorMaterial_CompareDefault(void* ptr, void* other)
{
	return static_cast<QSGVertexColorMaterial*>(ptr)->QSGVertexColorMaterial::compare(static_cast<QSGMaterial*>(other));
}

void* QSGVertexColorMaterial_CreateShader(void* ptr)
{
	return static_cast<QSGVertexColorMaterial*>(ptr)->createShader();
}

void* QSGVertexColorMaterial_CreateShaderDefault(void* ptr)
{
	return static_cast<QSGVertexColorMaterial*>(ptr)->QSGVertexColorMaterial::createShader();
}

void* QSGVertexColorMaterial_Type(void* ptr)
{
	return static_cast<QSGVertexColorMaterial*>(ptr)->type();
}

void* QSGVertexColorMaterial_TypeDefault(void* ptr)
{
	return static_cast<QSGVertexColorMaterial*>(ptr)->QSGVertexColorMaterial::type();
}

