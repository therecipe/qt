// +build !minimal

#define protected public
#define private public

#include "quick.h"
#include "_cgo_export.h"

#include <QAction>
#include <QActionEvent>
#include <QByteArray>
#include <QCamera>
#include <QCameraImageCapture>
#include <QChildEvent>
#include <QCloseEvent>
#include <QColor>
#include <QContextMenuEvent>
#include <QCursor>
#include <QDBusPendingCall>
#include <QDBusPendingCallWatcher>
#include <QDrag>
#include <QDragEnterEvent>
#include <QDragLeaveEvent>
#include <QDragMoveEvent>
#include <QDropEvent>
#include <QEvent>
#include <QExposeEvent>
#include <QExtensionFactory>
#include <QExtensionManager>
#include <QFocusEvent>
#include <QGraphicsObject>
#include <QGraphicsWidget>
#include <QHideEvent>
#include <QHoverEvent>
#include <QIcon>
#include <QImage>
#include <QInputMethod>
#include <QInputMethodEvent>
#include <QKeyEvent>
#include <QLayout>
#include <QList>
#include <QMatrix4x4>
#include <QMediaPlaylist>
#include <QMediaRecorder>
#include <QMetaMethod>
#include <QMetaObject>
#include <QMouseEvent>
#include <QMoveEvent>
#include <QObject>
#include <QOffscreenSurface>
#include <QOpenGLContext>
#include <QOpenGLFramebufferObject>
#include <QOpenGLShader>
#include <QPaintDevice>
#include <QPaintDeviceWindow>
#include <QPaintEngine>
#include <QPaintEvent>
#include <QPainter>
#include <QPdfWriter>
#include <QPixmap>
#include <QPoint>
#include <QPointF>
#include <QQmlEngine>
#include <QQmlError>
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
#include <QRadioData>
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
#include <QSGImageNode>
#include <QSGMaterial>
#include <QSGMaterialShader>
#include <QSGMaterialType>
#include <QSGNode>
#include <QSGOpacityNode>
#include <QSGOpaqueTextureMaterial>
#include <QSGRectangleNode>
#include <QSGRenderNode>
#include <QSGRendererInterface>
#include <QSGTexture>
#include <QSGTextureMaterial>
#include <QSGTextureProvider>
#include <QSGTransformNode>
#include <QSGVertexColorMaterial>
#include <QScreen>
#include <QShowEvent>
#include <QSignalSpy>
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
#include <QVector>
#include <QWheelEvent>
#include <QWidget>
#include <QWindow>

class MyQQuickAsyncImageProvider: public QQuickAsyncImageProvider
{
public:
	MyQQuickAsyncImageProvider() : QQuickAsyncImageProvider() {};
	QQuickImageResponse * requestImageResponse(const QString & id, const QSize & requestedSize) { QByteArray t87ea5d = id.toUtf8(); QtQuick_PackedString idPacked = { const_cast<char*>(t87ea5d.prepend("WHITESPACE").constData()+10), t87ea5d.size()-10 };return static_cast<QQuickImageResponse*>(callbackQQuickAsyncImageProvider_RequestImageResponse(this, idPacked, const_cast<QSize*>(&requestedSize))); };
	 ~MyQQuickAsyncImageProvider() { callbackQQuickAsyncImageProvider_DestroyQQuickAsyncImageProvider(this); };
	QImage requestImage(const QString & id, QSize * size, const QSize & requestedSize) { QByteArray t87ea5d = id.toUtf8(); QtQuick_PackedString idPacked = { const_cast<char*>(t87ea5d.prepend("WHITESPACE").constData()+10), t87ea5d.size()-10 };return *static_cast<QImage*>(callbackQQuickImageProvider_RequestImage(this, idPacked, size, const_cast<QSize*>(&requestedSize))); };
	QPixmap requestPixmap(const QString & id, QSize * size, const QSize & requestedSize) { QByteArray t87ea5d = id.toUtf8(); QtQuick_PackedString idPacked = { const_cast<char*>(t87ea5d.prepend("WHITESPACE").constData()+10), t87ea5d.size()-10 };return *static_cast<QPixmap*>(callbackQQuickImageProvider_RequestPixmap(this, idPacked, size, const_cast<QSize*>(&requestedSize))); };
	QQuickTextureFactory * requestTexture(const QString & id, QSize * size, const QSize & requestedSize) { QByteArray t87ea5d = id.toUtf8(); QtQuick_PackedString idPacked = { const_cast<char*>(t87ea5d.prepend("WHITESPACE").constData()+10), t87ea5d.size()-10 };return static_cast<QQuickTextureFactory*>(callbackQQuickImageProvider_RequestTexture(this, idPacked, size, const_cast<QSize*>(&requestedSize))); };
	Flags flags() const { return static_cast<QQmlImageProviderBase::Flag>(callbackQQuickImageProvider_Flags(const_cast<void*>(static_cast<const void*>(this)))); };
	ImageType imageType() const { return static_cast<QQmlImageProviderBase::ImageType>(callbackQQuickImageProvider_ImageType(const_cast<void*>(static_cast<const void*>(this)))); };
};

void* QQuickAsyncImageProvider_NewQQuickAsyncImageProvider()
{
	return new MyQQuickAsyncImageProvider();
}

void* QQuickAsyncImageProvider_RequestImageResponse(void* ptr, struct QtQuick_PackedString id, void* requestedSize)
{
	return static_cast<QQuickAsyncImageProvider*>(ptr)->requestImageResponse(QString::fromUtf8(id.data, id.len), *static_cast<QSize*>(requestedSize));
}

void QQuickAsyncImageProvider_DestroyQQuickAsyncImageProvider(void* ptr)
{
	static_cast<QQuickAsyncImageProvider*>(ptr)->~QQuickAsyncImageProvider();
}

void QQuickAsyncImageProvider_DestroyQQuickAsyncImageProviderDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

class MyQQuickFramebufferObject: public QQuickFramebufferObject
{
public:
	void Signal_MirrorVerticallyChanged(bool vbo) { callbackQQuickFramebufferObject_MirrorVerticallyChanged(this, vbo); };
	void releaseResources() { callbackQQuickItem_ReleaseResources(this); };
	void Signal_TextureFollowsItemSizeChanged(bool vbo) { callbackQQuickFramebufferObject_TextureFollowsItemSizeChanged(this, vbo); };
	QSGTextureProvider * textureProvider() const { return static_cast<QSGTextureProvider*>(callbackQQuickItem_TextureProvider(const_cast<void*>(static_cast<const void*>(this)))); };
	bool isTextureProvider() const { return callbackQQuickItem_IsTextureProvider(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	
	bool childMouseEventFilter(QQuickItem * item, QEvent * event) { return callbackQQuickItem_ChildMouseEventFilter(this, item, event) != 0; };
	bool event(QEvent * ev) { return callbackQQuickItem_Event(this, ev) != 0; };
	void classBegin() { callbackQQuickItem_ClassBegin(this); };
	void componentComplete() { callbackQQuickItem_ComponentComplete(this); };
	void dragEnterEvent(QDragEnterEvent * event) { callbackQQuickItem_DragEnterEvent(this, event); };
	void dragLeaveEvent(QDragLeaveEvent * event) { callbackQQuickItem_DragLeaveEvent(this, event); };
	void dragMoveEvent(QDragMoveEvent * event) { callbackQQuickItem_DragMoveEvent(this, event); };
	void dropEvent(QDropEvent * event) { callbackQQuickItem_DropEvent(this, event); };
	void focusInEvent(QFocusEvent * event) { callbackQQuickItem_FocusInEvent(this, event); };
	void focusOutEvent(QFocusEvent * event) { callbackQQuickItem_FocusOutEvent(this, event); };
	void geometryChanged(const QRectF & newGeometry, const QRectF & oldGeometry) { callbackQQuickItem_GeometryChanged(this, const_cast<QRectF*>(&newGeometry), const_cast<QRectF*>(&oldGeometry)); };
	void hoverEnterEvent(QHoverEvent * event) { callbackQQuickItem_HoverEnterEvent(this, event); };
	void hoverLeaveEvent(QHoverEvent * event) { callbackQQuickItem_HoverLeaveEvent(this, event); };
	void hoverMoveEvent(QHoverEvent * event) { callbackQQuickItem_HoverMoveEvent(this, event); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackQQuickItem_InputMethodEvent(this, event); };
	
	void keyPressEvent(QKeyEvent * event) { callbackQQuickItem_KeyPressEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQQuickItem_KeyReleaseEvent(this, event); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQQuickItem_MouseDoubleClickEvent(this, event); };
	void mouseMoveEvent(QMouseEvent * event) { callbackQQuickItem_MouseMoveEvent(this, event); };
	void mousePressEvent(QMouseEvent * event) { callbackQQuickItem_MousePressEvent(this, event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackQQuickItem_MouseReleaseEvent(this, event); };
	void mouseUngrabEvent() { callbackQQuickItem_MouseUngrabEvent(this); };
	void touchEvent(QTouchEvent * event) { callbackQQuickItem_TouchEvent(this, event); };
	void touchUngrabEvent() { callbackQQuickItem_TouchUngrabEvent(this); };
	void update() { callbackQQuickItem_Update(this); };
	void updatePolish() { callbackQQuickItem_UpdatePolish(this); };
	void wheelEvent(QWheelEvent * event) { callbackQQuickItem_WheelEvent(this, event); };
	void Signal_WindowChanged(QQuickWindow * window) { callbackQQuickItem_WindowChanged(this, window); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackQQuickItem_InputMethodQuery(const_cast<void*>(static_cast<const void*>(this)), query)); };
	bool contains(const QPointF & point) const { return callbackQQuickItem_Contains(const_cast<void*>(static_cast<const void*>(this)), const_cast<QPointF*>(&point)) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQQuickItem_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQQuickItem_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQQuickItem_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQQuickItem_CustomEvent(this, event); };
	void deleteLater() { callbackQQuickItem_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQQuickItem_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQQuickItem_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtQuick_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQQuickItem_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQQuickItem_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQQuickItem_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQQuickFramebufferObject*)

int QQuickFramebufferObject_QQuickFramebufferObject_QRegisterMetaType(){qRegisterMetaType<QQuickFramebufferObject*>(); return qRegisterMetaType<MyQQuickFramebufferObject*>();}

void QQuickFramebufferObject_ConnectMirrorVerticallyChanged(void* ptr)
{
	QObject::connect(static_cast<QQuickFramebufferObject*>(ptr), static_cast<void (QQuickFramebufferObject::*)(bool)>(&QQuickFramebufferObject::mirrorVerticallyChanged), static_cast<MyQQuickFramebufferObject*>(ptr), static_cast<void (MyQQuickFramebufferObject::*)(bool)>(&MyQQuickFramebufferObject::Signal_MirrorVerticallyChanged));
}

void QQuickFramebufferObject_DisconnectMirrorVerticallyChanged(void* ptr)
{
	QObject::disconnect(static_cast<QQuickFramebufferObject*>(ptr), static_cast<void (QQuickFramebufferObject::*)(bool)>(&QQuickFramebufferObject::mirrorVerticallyChanged), static_cast<MyQQuickFramebufferObject*>(ptr), static_cast<void (MyQQuickFramebufferObject::*)(bool)>(&MyQQuickFramebufferObject::Signal_MirrorVerticallyChanged));
}

void QQuickFramebufferObject_MirrorVerticallyChanged(void* ptr, char vbo)
{
	static_cast<QQuickFramebufferObject*>(ptr)->mirrorVerticallyChanged(vbo != 0);
}

void QQuickFramebufferObject_SetMirrorVertically(void* ptr, char enable)
{
	static_cast<QQuickFramebufferObject*>(ptr)->setMirrorVertically(enable != 0);
}

void QQuickFramebufferObject_SetTextureFollowsItemSize(void* ptr, char follows)
{
	static_cast<QQuickFramebufferObject*>(ptr)->setTextureFollowsItemSize(follows != 0);
}

void QQuickFramebufferObject_ConnectTextureFollowsItemSizeChanged(void* ptr)
{
	QObject::connect(static_cast<QQuickFramebufferObject*>(ptr), static_cast<void (QQuickFramebufferObject::*)(bool)>(&QQuickFramebufferObject::textureFollowsItemSizeChanged), static_cast<MyQQuickFramebufferObject*>(ptr), static_cast<void (MyQQuickFramebufferObject::*)(bool)>(&MyQQuickFramebufferObject::Signal_TextureFollowsItemSizeChanged));
}

void QQuickFramebufferObject_DisconnectTextureFollowsItemSizeChanged(void* ptr)
{
	QObject::disconnect(static_cast<QQuickFramebufferObject*>(ptr), static_cast<void (QQuickFramebufferObject::*)(bool)>(&QQuickFramebufferObject::textureFollowsItemSizeChanged), static_cast<MyQQuickFramebufferObject*>(ptr), static_cast<void (MyQQuickFramebufferObject::*)(bool)>(&MyQQuickFramebufferObject::Signal_TextureFollowsItemSizeChanged));
}

void QQuickFramebufferObject_TextureFollowsItemSizeChanged(void* ptr, char vbo)
{
	static_cast<QQuickFramebufferObject*>(ptr)->textureFollowsItemSizeChanged(vbo != 0);
}

char QQuickFramebufferObject_MirrorVertically(void* ptr)
{
	return static_cast<QQuickFramebufferObject*>(ptr)->mirrorVertically();
}

char QQuickFramebufferObject_TextureFollowsItemSize(void* ptr)
{
	return static_cast<QQuickFramebufferObject*>(ptr)->textureFollowsItemSize();
}

class MyQQuickImageProvider: public QQuickImageProvider
{
public:
	MyQQuickImageProvider(ImageType ty, Flags flags = Flags()) : QQuickImageProvider(ty, flags) {};
	QImage requestImage(const QString & id, QSize * size, const QSize & requestedSize) { QByteArray t87ea5d = id.toUtf8(); QtQuick_PackedString idPacked = { const_cast<char*>(t87ea5d.prepend("WHITESPACE").constData()+10), t87ea5d.size()-10 };return *static_cast<QImage*>(callbackQQuickImageProvider_RequestImage(this, idPacked, size, const_cast<QSize*>(&requestedSize))); };
	QPixmap requestPixmap(const QString & id, QSize * size, const QSize & requestedSize) { QByteArray t87ea5d = id.toUtf8(); QtQuick_PackedString idPacked = { const_cast<char*>(t87ea5d.prepend("WHITESPACE").constData()+10), t87ea5d.size()-10 };return *static_cast<QPixmap*>(callbackQQuickImageProvider_RequestPixmap(this, idPacked, size, const_cast<QSize*>(&requestedSize))); };
	QQuickTextureFactory * requestTexture(const QString & id, QSize * size, const QSize & requestedSize) { QByteArray t87ea5d = id.toUtf8(); QtQuick_PackedString idPacked = { const_cast<char*>(t87ea5d.prepend("WHITESPACE").constData()+10), t87ea5d.size()-10 };return static_cast<QQuickTextureFactory*>(callbackQQuickImageProvider_RequestTexture(this, idPacked, size, const_cast<QSize*>(&requestedSize))); };
	 ~MyQQuickImageProvider() { callbackQQuickImageProvider_DestroyQQuickImageProvider(this); };
	Flags flags() const { return static_cast<QQmlImageProviderBase::Flag>(callbackQQuickImageProvider_Flags(const_cast<void*>(static_cast<const void*>(this)))); };
	ImageType imageType() const { return static_cast<QQmlImageProviderBase::ImageType>(callbackQQuickImageProvider_ImageType(const_cast<void*>(static_cast<const void*>(this)))); };
};

void* QQuickImageProvider_RequestImage(void* ptr, struct QtQuick_PackedString id, void* size, void* requestedSize)
{
	return new QImage(static_cast<QQuickImageProvider*>(ptr)->requestImage(QString::fromUtf8(id.data, id.len), static_cast<QSize*>(size), *static_cast<QSize*>(requestedSize)));
}

void* QQuickImageProvider_RequestImageDefault(void* ptr, struct QtQuick_PackedString id, void* size, void* requestedSize)
{
	if (dynamic_cast<QQuickAsyncImageProvider*>(static_cast<QQuickImageProvider*>(ptr))) {
		return new QImage(static_cast<QQuickAsyncImageProvider*>(ptr)->QQuickAsyncImageProvider::requestImage(QString::fromUtf8(id.data, id.len), static_cast<QSize*>(size), *static_cast<QSize*>(requestedSize)));
	} else {
		return new QImage(static_cast<QQuickImageProvider*>(ptr)->QQuickImageProvider::requestImage(QString::fromUtf8(id.data, id.len), static_cast<QSize*>(size), *static_cast<QSize*>(requestedSize)));
	}
}

void* QQuickImageProvider_RequestPixmap(void* ptr, struct QtQuick_PackedString id, void* size, void* requestedSize)
{
	return new QPixmap(static_cast<QQuickImageProvider*>(ptr)->requestPixmap(QString::fromUtf8(id.data, id.len), static_cast<QSize*>(size), *static_cast<QSize*>(requestedSize)));
}

void* QQuickImageProvider_RequestPixmapDefault(void* ptr, struct QtQuick_PackedString id, void* size, void* requestedSize)
{
	if (dynamic_cast<QQuickAsyncImageProvider*>(static_cast<QQuickImageProvider*>(ptr))) {
		return new QPixmap(static_cast<QQuickAsyncImageProvider*>(ptr)->QQuickAsyncImageProvider::requestPixmap(QString::fromUtf8(id.data, id.len), static_cast<QSize*>(size), *static_cast<QSize*>(requestedSize)));
	} else {
		return new QPixmap(static_cast<QQuickImageProvider*>(ptr)->QQuickImageProvider::requestPixmap(QString::fromUtf8(id.data, id.len), static_cast<QSize*>(size), *static_cast<QSize*>(requestedSize)));
	}
}

void* QQuickImageProvider_NewQQuickImageProvider(long long ty, long long flags)
{
	return new MyQQuickImageProvider(static_cast<QQmlImageProviderBase::ImageType>(ty), static_cast<QQmlImageProviderBase::Flag>(flags));
}

void* QQuickImageProvider_RequestTexture(void* ptr, struct QtQuick_PackedString id, void* size, void* requestedSize)
{
	return static_cast<QQuickImageProvider*>(ptr)->requestTexture(QString::fromUtf8(id.data, id.len), static_cast<QSize*>(size), *static_cast<QSize*>(requestedSize));
}

void* QQuickImageProvider_RequestTextureDefault(void* ptr, struct QtQuick_PackedString id, void* size, void* requestedSize)
{
	if (dynamic_cast<QQuickAsyncImageProvider*>(static_cast<QQuickImageProvider*>(ptr))) {
		return static_cast<QQuickAsyncImageProvider*>(ptr)->QQuickAsyncImageProvider::requestTexture(QString::fromUtf8(id.data, id.len), static_cast<QSize*>(size), *static_cast<QSize*>(requestedSize));
	} else {
		return static_cast<QQuickImageProvider*>(ptr)->QQuickImageProvider::requestTexture(QString::fromUtf8(id.data, id.len), static_cast<QSize*>(size), *static_cast<QSize*>(requestedSize));
	}
}

void QQuickImageProvider_DestroyQQuickImageProvider(void* ptr)
{
	static_cast<QQuickImageProvider*>(ptr)->~QQuickImageProvider();
}

void QQuickImageProvider_DestroyQQuickImageProviderDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

long long QQuickImageProvider_Flags(void* ptr)
{
	return static_cast<QQuickImageProvider*>(ptr)->flags();
}

long long QQuickImageProvider_FlagsDefault(void* ptr)
{
	if (dynamic_cast<QQuickAsyncImageProvider*>(static_cast<QQuickImageProvider*>(ptr))) {
		return static_cast<QQuickAsyncImageProvider*>(ptr)->QQuickAsyncImageProvider::flags();
	} else {
		return static_cast<QQuickImageProvider*>(ptr)->QQuickImageProvider::flags();
	}
}

long long QQuickImageProvider_ImageType(void* ptr)
{
	return static_cast<QQuickImageProvider*>(ptr)->imageType();
}

long long QQuickImageProvider_ImageTypeDefault(void* ptr)
{
	if (dynamic_cast<QQuickAsyncImageProvider*>(static_cast<QQuickImageProvider*>(ptr))) {
		return static_cast<QQuickAsyncImageProvider*>(ptr)->QQuickAsyncImageProvider::imageType();
	} else {
		return static_cast<QQuickImageProvider*>(ptr)->QQuickImageProvider::imageType();
	}
}

class MyQQuickImageResponse: public QQuickImageResponse
{
public:
	MyQQuickImageResponse() : QQuickImageResponse() {QQuickImageResponse_QQuickImageResponse_QRegisterMetaType();};
	void cancel() { callbackQQuickImageResponse_Cancel(this); };
	void Signal_Finished() { callbackQQuickImageResponse_Finished(this); };
	 ~MyQQuickImageResponse() { callbackQQuickImageResponse_DestroyQQuickImageResponse(this); };
	QQuickTextureFactory * textureFactory() const { return static_cast<QQuickTextureFactory*>(callbackQQuickImageResponse_TextureFactory(const_cast<void*>(static_cast<const void*>(this)))); };
	QString errorString() const { return ({ QtQuick_PackedString tempVal = callbackQQuickImageResponse_ErrorString(const_cast<void*>(static_cast<const void*>(this))); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	bool event(QEvent * e) { return callbackQQuickImageResponse_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQQuickImageResponse_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQQuickImageResponse_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQQuickImageResponse_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQQuickImageResponse_CustomEvent(this, event); };
	void deleteLater() { callbackQQuickImageResponse_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQQuickImageResponse_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQQuickImageResponse_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtQuick_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQQuickImageResponse_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQQuickImageResponse_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQQuickImageResponse_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQQuickImageResponse*)

int QQuickImageResponse_QQuickImageResponse_QRegisterMetaType(){qRegisterMetaType<QQuickImageResponse*>(); return qRegisterMetaType<MyQQuickImageResponse*>();}

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

void QQuickImageResponse_DestroyQQuickImageResponse(void* ptr)
{
	static_cast<QQuickImageResponse*>(ptr)->~QQuickImageResponse();
}

void QQuickImageResponse_DestroyQQuickImageResponseDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QQuickImageResponse_TextureFactory(void* ptr)
{
	return static_cast<QQuickImageResponse*>(ptr)->textureFactory();
}

struct QtQuick_PackedString QQuickImageResponse_ErrorString(void* ptr)
{
	return ({ QByteArray tfd721f = static_cast<QQuickImageResponse*>(ptr)->errorString().toUtf8(); QtQuick_PackedString { const_cast<char*>(tfd721f.prepend("WHITESPACE").constData()+10), tfd721f.size()-10 }; });
}

struct QtQuick_PackedString QQuickImageResponse_ErrorStringDefault(void* ptr)
{
		return ({ QByteArray t94ceb3 = static_cast<QQuickImageResponse*>(ptr)->QQuickImageResponse::errorString().toUtf8(); QtQuick_PackedString { const_cast<char*>(t94ceb3.prepend("WHITESPACE").constData()+10), t94ceb3.size()-10 }; });
}

void* QQuickImageResponse___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QQuickImageResponse___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QQuickImageResponse___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QQuickImageResponse___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QQuickImageResponse___findChildren_setList2(void* ptr, void* i)
{
	if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QQuickItem*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QQuickImageResponse___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QQuickImageResponse___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QQuickImageResponse___findChildren_setList3(void* ptr, void* i)
{
	if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QQuickItem*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QQuickImageResponse___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QQuickImageResponse___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QQuickImageResponse___findChildren_setList(void* ptr, void* i)
{
	if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QQuickItem*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QQuickImageResponse___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QQuickImageResponse___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QQuickImageResponse___children_setList(void* ptr, void* i)
{
	if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QQuickItem*>(i));
	} else {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QQuickImageResponse___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

char QQuickImageResponse_EventDefault(void* ptr, void* e)
{
		return static_cast<QQuickImageResponse*>(ptr)->QQuickImageResponse::event(static_cast<QEvent*>(e));
}

char QQuickImageResponse_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(watched))) {
		return static_cast<QQuickImageResponse*>(ptr)->QQuickImageResponse::eventFilter(static_cast<QQuickItem*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QQuickImageResponse*>(ptr)->QQuickImageResponse::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

void QQuickImageResponse_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QQuickImageResponse*>(ptr)->QQuickImageResponse::childEvent(static_cast<QChildEvent*>(event));
}

void QQuickImageResponse_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QQuickImageResponse*>(ptr)->QQuickImageResponse::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQuickImageResponse_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QQuickImageResponse*>(ptr)->QQuickImageResponse::customEvent(static_cast<QEvent*>(event));
}

void QQuickImageResponse_DeleteLaterDefault(void* ptr)
{
		static_cast<QQuickImageResponse*>(ptr)->QQuickImageResponse::deleteLater();
}

void QQuickImageResponse_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QQuickImageResponse*>(ptr)->QQuickImageResponse::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQuickImageResponse_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QQuickImageResponse*>(ptr)->QQuickImageResponse::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QQuickImageResponse_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QQuickImageResponse*>(ptr)->QQuickImageResponse::metaObject());
}

class MyQQuickItem: public QQuickItem
{
public:
	MyQQuickItem(QQuickItem *parent = Q_NULLPTR) : QQuickItem(parent) {QQuickItem_QQuickItem_QRegisterMetaType();};
	bool childMouseEventFilter(QQuickItem * item, QEvent * event) { return callbackQQuickItem_ChildMouseEventFilter(this, item, event) != 0; };
	bool event(QEvent * ev) { return callbackQQuickItem_Event(this, ev) != 0; };
	void classBegin() { callbackQQuickItem_ClassBegin(this); };
	void componentComplete() { callbackQQuickItem_ComponentComplete(this); };
	void dragEnterEvent(QDragEnterEvent * event) { callbackQQuickItem_DragEnterEvent(this, event); };
	void dragLeaveEvent(QDragLeaveEvent * event) { callbackQQuickItem_DragLeaveEvent(this, event); };
	void dragMoveEvent(QDragMoveEvent * event) { callbackQQuickItem_DragMoveEvent(this, event); };
	void dropEvent(QDropEvent * event) { callbackQQuickItem_DropEvent(this, event); };
	void focusInEvent(QFocusEvent * event) { callbackQQuickItem_FocusInEvent(this, event); };
	void focusOutEvent(QFocusEvent * event) { callbackQQuickItem_FocusOutEvent(this, event); };
	void geometryChanged(const QRectF & newGeometry, const QRectF & oldGeometry) { callbackQQuickItem_GeometryChanged(this, const_cast<QRectF*>(&newGeometry), const_cast<QRectF*>(&oldGeometry)); };
	void hoverEnterEvent(QHoverEvent * event) { callbackQQuickItem_HoverEnterEvent(this, event); };
	void hoverLeaveEvent(QHoverEvent * event) { callbackQQuickItem_HoverLeaveEvent(this, event); };
	void hoverMoveEvent(QHoverEvent * event) { callbackQQuickItem_HoverMoveEvent(this, event); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackQQuickItem_InputMethodEvent(this, event); };
	void keyPressEvent(QKeyEvent * event) { callbackQQuickItem_KeyPressEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQQuickItem_KeyReleaseEvent(this, event); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQQuickItem_MouseDoubleClickEvent(this, event); };
	void mouseMoveEvent(QMouseEvent * event) { callbackQQuickItem_MouseMoveEvent(this, event); };
	void mousePressEvent(QMouseEvent * event) { callbackQQuickItem_MousePressEvent(this, event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackQQuickItem_MouseReleaseEvent(this, event); };
	void mouseUngrabEvent() { callbackQQuickItem_MouseUngrabEvent(this); };
	void releaseResources() { callbackQQuickItem_ReleaseResources(this); };
	void touchEvent(QTouchEvent * event) { callbackQQuickItem_TouchEvent(this, event); };
	void touchUngrabEvent() { callbackQQuickItem_TouchUngrabEvent(this); };
	void update() { callbackQQuickItem_Update(this); };
	void updatePolish() { callbackQQuickItem_UpdatePolish(this); };
	void wheelEvent(QWheelEvent * event) { callbackQQuickItem_WheelEvent(this, event); };
	void Signal_WindowChanged(QQuickWindow * window) { callbackQQuickItem_WindowChanged(this, window); };
	 ~MyQQuickItem() { callbackQQuickItem_DestroyQQuickItem(this); };
	QSGTextureProvider * textureProvider() const { return static_cast<QSGTextureProvider*>(callbackQQuickItem_TextureProvider(const_cast<void*>(static_cast<const void*>(this)))); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackQQuickItem_InputMethodQuery(const_cast<void*>(static_cast<const void*>(this)), query)); };
	bool contains(const QPointF & point) const { return callbackQQuickItem_Contains(const_cast<void*>(static_cast<const void*>(this)), const_cast<QPointF*>(&point)) != 0; };
	bool isTextureProvider() const { return callbackQQuickItem_IsTextureProvider(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQQuickItem_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQQuickItem_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQQuickItem_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQQuickItem_CustomEvent(this, event); };
	void deleteLater() { callbackQQuickItem_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQQuickItem_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQQuickItem_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtQuick_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQQuickItem_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQQuickItem_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQQuickItem_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQQuickItem*)

int QQuickItem_QQuickItem_QRegisterMetaType(){qRegisterMetaType<QQuickItem*>(); return qRegisterMetaType<MyQQuickItem*>();}

void* QQuickItem_NextItemInFocusChain(void* ptr, char forward)
{
		return static_cast<QQuickItem*>(ptr)->nextItemInFocusChain(forward != 0);
}

void* QQuickItem_NewQQuickItem(void* parent)
{
	return new MyQQuickItem(static_cast<QQuickItem*>(parent));
}

void* QQuickItem_ChildrenRect(void* ptr)
{
		return ({ QRectF tmpValue = static_cast<QQuickItem*>(ptr)->childrenRect(); new QRectF(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

char QQuickItem_ChildMouseEventFilter(void* ptr, void* item, void* event)
{
		return static_cast<QQuickItem*>(ptr)->childMouseEventFilter(static_cast<QQuickItem*>(item), static_cast<QEvent*>(event));
}

char QQuickItem_ChildMouseEventFilterDefault(void* ptr, void* item, void* event)
{
	if (dynamic_cast<QQuickPaintedItem*>(static_cast<QObject*>(ptr))) {
		return static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::childMouseEventFilter(static_cast<QQuickItem*>(item), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QQuickFramebufferObject*>(static_cast<QObject*>(ptr))) {
		return static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::childMouseEventFilter(static_cast<QQuickItem*>(item), static_cast<QEvent*>(event));
	} else {
		return static_cast<QQuickItem*>(ptr)->QQuickItem::childMouseEventFilter(static_cast<QQuickItem*>(item), static_cast<QEvent*>(event));
	}
}

char QQuickItem_Event(void* ptr, void* ev)
{
		return static_cast<QQuickItem*>(ptr)->event(static_cast<QEvent*>(ev));
}

char QQuickItem_EventDefault(void* ptr, void* ev)
{
	if (dynamic_cast<QQuickPaintedItem*>(static_cast<QObject*>(ptr))) {
		return static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::event(static_cast<QEvent*>(ev));
	} else if (dynamic_cast<QQuickFramebufferObject*>(static_cast<QObject*>(ptr))) {
		return static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::event(static_cast<QEvent*>(ev));
	} else {
		return static_cast<QQuickItem*>(ptr)->QQuickItem::event(static_cast<QEvent*>(ev));
	}
}

void QQuickItem_ClassBegin(void* ptr)
{
		static_cast<QQuickItem*>(ptr)->classBegin();
}

void QQuickItem_ClassBeginDefault(void* ptr)
{
	if (dynamic_cast<QQuickPaintedItem*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::classBegin();
	} else if (dynamic_cast<QQuickFramebufferObject*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::classBegin();
	} else {
		static_cast<QQuickItem*>(ptr)->QQuickItem::classBegin();
	}
}

void QQuickItem_ComponentComplete(void* ptr)
{
		static_cast<QQuickItem*>(ptr)->componentComplete();
}

void QQuickItem_ComponentCompleteDefault(void* ptr)
{
	if (dynamic_cast<QQuickPaintedItem*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::componentComplete();
	} else if (dynamic_cast<QQuickFramebufferObject*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::componentComplete();
	} else {
		static_cast<QQuickItem*>(ptr)->QQuickItem::componentComplete();
	}
}

void QQuickItem_DragEnterEvent(void* ptr, void* event)
{
		static_cast<QQuickItem*>(ptr)->dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QQuickItem_DragEnterEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QQuickPaintedItem*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
	} else if (dynamic_cast<QQuickFramebufferObject*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
	} else {
		static_cast<QQuickItem*>(ptr)->QQuickItem::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
	}
}

void QQuickItem_DragLeaveEvent(void* ptr, void* event)
{
		static_cast<QQuickItem*>(ptr)->dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QQuickItem_DragLeaveEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QQuickPaintedItem*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
	} else if (dynamic_cast<QQuickFramebufferObject*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
	} else {
		static_cast<QQuickItem*>(ptr)->QQuickItem::dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
	}
}

void QQuickItem_DragMoveEvent(void* ptr, void* event)
{
		static_cast<QQuickItem*>(ptr)->dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QQuickItem_DragMoveEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QQuickPaintedItem*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::dragMoveEvent(static_cast<QDragMoveEvent*>(event));
	} else if (dynamic_cast<QQuickFramebufferObject*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::dragMoveEvent(static_cast<QDragMoveEvent*>(event));
	} else {
		static_cast<QQuickItem*>(ptr)->QQuickItem::dragMoveEvent(static_cast<QDragMoveEvent*>(event));
	}
}

void QQuickItem_DropEvent(void* ptr, void* event)
{
		static_cast<QQuickItem*>(ptr)->dropEvent(static_cast<QDropEvent*>(event));
}

void QQuickItem_DropEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QQuickPaintedItem*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::dropEvent(static_cast<QDropEvent*>(event));
	} else if (dynamic_cast<QQuickFramebufferObject*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::dropEvent(static_cast<QDropEvent*>(event));
	} else {
		static_cast<QQuickItem*>(ptr)->QQuickItem::dropEvent(static_cast<QDropEvent*>(event));
	}
}

void QQuickItem_FocusInEvent(void* ptr, void* event)
{
		static_cast<QQuickItem*>(ptr)->focusInEvent(static_cast<QFocusEvent*>(event));
}

void QQuickItem_FocusInEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QQuickPaintedItem*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::focusInEvent(static_cast<QFocusEvent*>(event));
	} else if (dynamic_cast<QQuickFramebufferObject*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::focusInEvent(static_cast<QFocusEvent*>(event));
	} else {
		static_cast<QQuickItem*>(ptr)->QQuickItem::focusInEvent(static_cast<QFocusEvent*>(event));
	}
}

void QQuickItem_FocusOutEvent(void* ptr, void* event)
{
		static_cast<QQuickItem*>(ptr)->focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QQuickItem_FocusOutEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QQuickPaintedItem*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::focusOutEvent(static_cast<QFocusEvent*>(event));
	} else if (dynamic_cast<QQuickFramebufferObject*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::focusOutEvent(static_cast<QFocusEvent*>(event));
	} else {
		static_cast<QQuickItem*>(ptr)->QQuickItem::focusOutEvent(static_cast<QFocusEvent*>(event));
	}
}

void QQuickItem_ForceActiveFocus(void* ptr)
{
		static_cast<QQuickItem*>(ptr)->forceActiveFocus();
}

void QQuickItem_ForceActiveFocus2(void* ptr, long long reason)
{
		static_cast<QQuickItem*>(ptr)->forceActiveFocus(static_cast<Qt::FocusReason>(reason));
}

void QQuickItem_GeometryChanged(void* ptr, void* newGeometry, void* oldGeometry)
{
		static_cast<QQuickItem*>(ptr)->geometryChanged(*static_cast<QRectF*>(newGeometry), *static_cast<QRectF*>(oldGeometry));
}

void QQuickItem_GeometryChangedDefault(void* ptr, void* newGeometry, void* oldGeometry)
{
	if (dynamic_cast<QQuickPaintedItem*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::geometryChanged(*static_cast<QRectF*>(newGeometry), *static_cast<QRectF*>(oldGeometry));
	} else if (dynamic_cast<QQuickFramebufferObject*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::geometryChanged(*static_cast<QRectF*>(newGeometry), *static_cast<QRectF*>(oldGeometry));
	} else {
		static_cast<QQuickItem*>(ptr)->QQuickItem::geometryChanged(*static_cast<QRectF*>(newGeometry), *static_cast<QRectF*>(oldGeometry));
	}
}

void QQuickItem_GrabMouse(void* ptr)
{
		static_cast<QQuickItem*>(ptr)->grabMouse();
}

void QQuickItem_GrabTouchPoints(void* ptr, void* ids)
{
		static_cast<QQuickItem*>(ptr)->grabTouchPoints(*static_cast<QVector<int>*>(ids));
}

void QQuickItem_HoverEnterEvent(void* ptr, void* event)
{
		static_cast<QQuickItem*>(ptr)->hoverEnterEvent(static_cast<QHoverEvent*>(event));
}

void QQuickItem_HoverEnterEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QQuickPaintedItem*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::hoverEnterEvent(static_cast<QHoverEvent*>(event));
	} else if (dynamic_cast<QQuickFramebufferObject*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::hoverEnterEvent(static_cast<QHoverEvent*>(event));
	} else {
		static_cast<QQuickItem*>(ptr)->QQuickItem::hoverEnterEvent(static_cast<QHoverEvent*>(event));
	}
}

void QQuickItem_HoverLeaveEvent(void* ptr, void* event)
{
		static_cast<QQuickItem*>(ptr)->hoverLeaveEvent(static_cast<QHoverEvent*>(event));
}

void QQuickItem_HoverLeaveEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QQuickPaintedItem*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::hoverLeaveEvent(static_cast<QHoverEvent*>(event));
	} else if (dynamic_cast<QQuickFramebufferObject*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::hoverLeaveEvent(static_cast<QHoverEvent*>(event));
	} else {
		static_cast<QQuickItem*>(ptr)->QQuickItem::hoverLeaveEvent(static_cast<QHoverEvent*>(event));
	}
}

void QQuickItem_HoverMoveEvent(void* ptr, void* event)
{
		static_cast<QQuickItem*>(ptr)->hoverMoveEvent(static_cast<QHoverEvent*>(event));
}

void QQuickItem_HoverMoveEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QQuickPaintedItem*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::hoverMoveEvent(static_cast<QHoverEvent*>(event));
	} else if (dynamic_cast<QQuickFramebufferObject*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::hoverMoveEvent(static_cast<QHoverEvent*>(event));
	} else {
		static_cast<QQuickItem*>(ptr)->QQuickItem::hoverMoveEvent(static_cast<QHoverEvent*>(event));
	}
}

void QQuickItem_InputMethodEvent(void* ptr, void* event)
{
		static_cast<QQuickItem*>(ptr)->inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void QQuickItem_InputMethodEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QQuickPaintedItem*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
	} else if (dynamic_cast<QQuickFramebufferObject*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
	} else {
		static_cast<QQuickItem*>(ptr)->QQuickItem::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
	}
}

void QQuickItem_KeyPressEvent(void* ptr, void* event)
{
		static_cast<QQuickItem*>(ptr)->keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QQuickItem_KeyPressEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QQuickPaintedItem*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::keyPressEvent(static_cast<QKeyEvent*>(event));
	} else if (dynamic_cast<QQuickFramebufferObject*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::keyPressEvent(static_cast<QKeyEvent*>(event));
	} else {
		static_cast<QQuickItem*>(ptr)->QQuickItem::keyPressEvent(static_cast<QKeyEvent*>(event));
	}
}

void QQuickItem_KeyReleaseEvent(void* ptr, void* event)
{
		static_cast<QQuickItem*>(ptr)->keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QQuickItem_KeyReleaseEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QQuickPaintedItem*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::keyReleaseEvent(static_cast<QKeyEvent*>(event));
	} else if (dynamic_cast<QQuickFramebufferObject*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::keyReleaseEvent(static_cast<QKeyEvent*>(event));
	} else {
		static_cast<QQuickItem*>(ptr)->QQuickItem::keyReleaseEvent(static_cast<QKeyEvent*>(event));
	}
}

void QQuickItem_MouseDoubleClickEvent(void* ptr, void* event)
{
		static_cast<QQuickItem*>(ptr)->mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QQuickItem_MouseDoubleClickEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QQuickPaintedItem*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QQuickFramebufferObject*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
	} else {
		static_cast<QQuickItem*>(ptr)->QQuickItem::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
	}
}

void QQuickItem_MouseMoveEvent(void* ptr, void* event)
{
		static_cast<QQuickItem*>(ptr)->mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QQuickItem_MouseMoveEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QQuickPaintedItem*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::mouseMoveEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QQuickFramebufferObject*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::mouseMoveEvent(static_cast<QMouseEvent*>(event));
	} else {
		static_cast<QQuickItem*>(ptr)->QQuickItem::mouseMoveEvent(static_cast<QMouseEvent*>(event));
	}
}

void QQuickItem_MousePressEvent(void* ptr, void* event)
{
		static_cast<QQuickItem*>(ptr)->mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QQuickItem_MousePressEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QQuickPaintedItem*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::mousePressEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QQuickFramebufferObject*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::mousePressEvent(static_cast<QMouseEvent*>(event));
	} else {
		static_cast<QQuickItem*>(ptr)->QQuickItem::mousePressEvent(static_cast<QMouseEvent*>(event));
	}
}

void QQuickItem_MouseReleaseEvent(void* ptr, void* event)
{
		static_cast<QQuickItem*>(ptr)->mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void QQuickItem_MouseReleaseEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QQuickPaintedItem*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QQuickFramebufferObject*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
	} else {
		static_cast<QQuickItem*>(ptr)->QQuickItem::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
	}
}

void QQuickItem_MouseUngrabEvent(void* ptr)
{
		static_cast<QQuickItem*>(ptr)->mouseUngrabEvent();
}

void QQuickItem_MouseUngrabEventDefault(void* ptr)
{
	if (dynamic_cast<QQuickPaintedItem*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::mouseUngrabEvent();
	} else if (dynamic_cast<QQuickFramebufferObject*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::mouseUngrabEvent();
	} else {
		static_cast<QQuickItem*>(ptr)->QQuickItem::mouseUngrabEvent();
	}
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
	if (dynamic_cast<QQuickPaintedItem*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::releaseResources();
	} else if (dynamic_cast<QQuickFramebufferObject*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::releaseResources();
	} else {
		static_cast<QQuickItem*>(ptr)->QQuickItem::releaseResources();
	}
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

void QQuickItem_SetAcceptHoverEvents(void* ptr, char enabled)
{
		static_cast<QQuickItem*>(ptr)->setAcceptHoverEvents(enabled != 0);
}

void QQuickItem_SetAcceptedMouseButtons(void* ptr, long long buttons)
{
		static_cast<QQuickItem*>(ptr)->setAcceptedMouseButtons(static_cast<Qt::MouseButton>(buttons));
}

void QQuickItem_SetActiveFocusOnTab(void* ptr, char vbo)
{
		static_cast<QQuickItem*>(ptr)->setActiveFocusOnTab(vbo != 0);
}

void QQuickItem_SetAntialiasing(void* ptr, char vbo)
{
		static_cast<QQuickItem*>(ptr)->setAntialiasing(vbo != 0);
}

void QQuickItem_SetBaselineOffset(void* ptr, double vqr)
{
		static_cast<QQuickItem*>(ptr)->setBaselineOffset(vqr);
}

void QQuickItem_SetClip(void* ptr, char vbo)
{
		static_cast<QQuickItem*>(ptr)->setClip(vbo != 0);
}

void QQuickItem_SetCursor(void* ptr, void* cursor)
{
		static_cast<QQuickItem*>(ptr)->setCursor(*static_cast<QCursor*>(cursor));
}

void QQuickItem_SetEnabled(void* ptr, char vbo)
{
		static_cast<QQuickItem*>(ptr)->setEnabled(vbo != 0);
}

void QQuickItem_SetFiltersChildMouseEvents(void* ptr, char filter)
{
		static_cast<QQuickItem*>(ptr)->setFiltersChildMouseEvents(filter != 0);
}

void QQuickItem_SetFlag(void* ptr, long long flag, char enabled)
{
		static_cast<QQuickItem*>(ptr)->setFlag(static_cast<QQuickItem::Flag>(flag), enabled != 0);
}

void QQuickItem_SetFlags(void* ptr, long long flags)
{
		static_cast<QQuickItem*>(ptr)->setFlags(static_cast<QQuickItem::Flag>(flags));
}

void QQuickItem_SetFocus(void* ptr, char vbo)
{
		static_cast<QQuickItem*>(ptr)->setFocus(vbo != 0);
}

void QQuickItem_SetFocus2(void* ptr, char focus, long long reason)
{
		static_cast<QQuickItem*>(ptr)->setFocus(focus != 0, static_cast<Qt::FocusReason>(reason));
}

void QQuickItem_SetHeight(void* ptr, double vqr)
{
		static_cast<QQuickItem*>(ptr)->setHeight(vqr);
}

void QQuickItem_SetImplicitHeight(void* ptr, double vqr)
{
		static_cast<QQuickItem*>(ptr)->setImplicitHeight(vqr);
}

void QQuickItem_SetImplicitWidth(void* ptr, double vqr)
{
		static_cast<QQuickItem*>(ptr)->setImplicitWidth(vqr);
}

void QQuickItem_SetKeepMouseGrab(void* ptr, char keep)
{
		static_cast<QQuickItem*>(ptr)->setKeepMouseGrab(keep != 0);
}

void QQuickItem_SetKeepTouchGrab(void* ptr, char keep)
{
		static_cast<QQuickItem*>(ptr)->setKeepTouchGrab(keep != 0);
}

void QQuickItem_SetOpacity(void* ptr, double vqr)
{
		static_cast<QQuickItem*>(ptr)->setOpacity(vqr);
}

void QQuickItem_SetParentItem(void* ptr, void* parent)
{
		static_cast<QQuickItem*>(ptr)->setParentItem(static_cast<QQuickItem*>(parent));
}

void QQuickItem_SetRotation(void* ptr, double vqr)
{
		static_cast<QQuickItem*>(ptr)->setRotation(vqr);
}

void QQuickItem_SetScale(void* ptr, double vqr)
{
		static_cast<QQuickItem*>(ptr)->setScale(vqr);
}

void QQuickItem_SetSmooth(void* ptr, char vbo)
{
		static_cast<QQuickItem*>(ptr)->setSmooth(vbo != 0);
}

void QQuickItem_SetState(void* ptr, struct QtQuick_PackedString vqs)
{
		static_cast<QQuickItem*>(ptr)->setState(QString::fromUtf8(vqs.data, vqs.len));
}

void QQuickItem_SetTransformOrigin(void* ptr, long long vtr)
{
		static_cast<QQuickItem*>(ptr)->setTransformOrigin(static_cast<QQuickItem::TransformOrigin>(vtr));
}

void QQuickItem_SetVisible(void* ptr, char vbo)
{
		static_cast<QQuickItem*>(ptr)->setVisible(vbo != 0);
}

void QQuickItem_SetWidth(void* ptr, double vqr)
{
		static_cast<QQuickItem*>(ptr)->setWidth(vqr);
}

void QQuickItem_SetX(void* ptr, double vqr)
{
		static_cast<QQuickItem*>(ptr)->setX(vqr);
}

void QQuickItem_SetY(void* ptr, double vqr)
{
		static_cast<QQuickItem*>(ptr)->setY(vqr);
}

void QQuickItem_SetZ(void* ptr, double vqr)
{
		static_cast<QQuickItem*>(ptr)->setZ(vqr);
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
	if (dynamic_cast<QQuickPaintedItem*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::touchEvent(static_cast<QTouchEvent*>(event));
	} else if (dynamic_cast<QQuickFramebufferObject*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::touchEvent(static_cast<QTouchEvent*>(event));
	} else {
		static_cast<QQuickItem*>(ptr)->QQuickItem::touchEvent(static_cast<QTouchEvent*>(event));
	}
}

void QQuickItem_TouchUngrabEvent(void* ptr)
{
		static_cast<QQuickItem*>(ptr)->touchUngrabEvent();
}

void QQuickItem_TouchUngrabEventDefault(void* ptr)
{
	if (dynamic_cast<QQuickPaintedItem*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::touchUngrabEvent();
	} else if (dynamic_cast<QQuickFramebufferObject*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::touchUngrabEvent();
	} else {
		static_cast<QQuickItem*>(ptr)->QQuickItem::touchUngrabEvent();
	}
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

void QQuickItem_UpdateDefault(void* ptr)
{
	if (dynamic_cast<QQuickPaintedItem*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::update();
	} else if (dynamic_cast<QQuickFramebufferObject*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::update();
	} else {
		static_cast<QQuickItem*>(ptr)->QQuickItem::update();
	}
}

void QQuickItem_UpdateInputMethod(void* ptr, long long queries)
{
		static_cast<QQuickItem*>(ptr)->updateInputMethod(static_cast<Qt::InputMethodQuery>(queries));
}

void QQuickItem_UpdatePolish(void* ptr)
{
		static_cast<QQuickItem*>(ptr)->updatePolish();
}

void QQuickItem_UpdatePolishDefault(void* ptr)
{
	if (dynamic_cast<QQuickPaintedItem*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::updatePolish();
	} else if (dynamic_cast<QQuickFramebufferObject*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::updatePolish();
	} else {
		static_cast<QQuickItem*>(ptr)->QQuickItem::updatePolish();
	}
}

void QQuickItem_WheelEvent(void* ptr, void* event)
{
		static_cast<QQuickItem*>(ptr)->wheelEvent(static_cast<QWheelEvent*>(event));
}

void QQuickItem_WheelEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QQuickPaintedItem*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::wheelEvent(static_cast<QWheelEvent*>(event));
	} else if (dynamic_cast<QQuickFramebufferObject*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::wheelEvent(static_cast<QWheelEvent*>(event));
	} else {
		static_cast<QQuickItem*>(ptr)->QQuickItem::wheelEvent(static_cast<QWheelEvent*>(event));
	}
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

void QQuickItem_DestroyQQuickItemDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

long long QQuickItem_Flags(void* ptr)
{
		return static_cast<QQuickItem*>(ptr)->flags();
}

void* QQuickItem_Cursor(void* ptr)
{
		return new QCursor(static_cast<QQuickItem*>(ptr)->cursor());
}

struct QtQuick_PackedList QQuickItem_ChildItems(void* ptr)
{
		return ({ QList<QQuickItem *>* tmpValue = new QList<QQuickItem *>(static_cast<QQuickItem*>(ptr)->childItems()); QtQuick_PackedList { tmpValue, tmpValue->size() }; });
}

void* QQuickItem_MapFromGlobal(void* ptr, void* point)
{
		return ({ QPointF tmpValue = static_cast<QQuickItem*>(ptr)->mapFromGlobal(*static_cast<QPointF*>(point)); new QPointF(tmpValue.x(), tmpValue.y()); });
}

void* QQuickItem_MapFromItem(void* ptr, void* item, void* point)
{
		return ({ QPointF tmpValue = static_cast<QQuickItem*>(ptr)->mapFromItem(static_cast<QQuickItem*>(item), *static_cast<QPointF*>(point)); new QPointF(tmpValue.x(), tmpValue.y()); });
}

void* QQuickItem_MapFromScene(void* ptr, void* point)
{
		return ({ QPointF tmpValue = static_cast<QQuickItem*>(ptr)->mapFromScene(*static_cast<QPointF*>(point)); new QPointF(tmpValue.x(), tmpValue.y()); });
}

void* QQuickItem_MapToGlobal(void* ptr, void* point)
{
		return ({ QPointF tmpValue = static_cast<QQuickItem*>(ptr)->mapToGlobal(*static_cast<QPointF*>(point)); new QPointF(tmpValue.x(), tmpValue.y()); });
}

void* QQuickItem_MapToItem(void* ptr, void* item, void* point)
{
		return ({ QPointF tmpValue = static_cast<QQuickItem*>(ptr)->mapToItem(static_cast<QQuickItem*>(item), *static_cast<QPointF*>(point)); new QPointF(tmpValue.x(), tmpValue.y()); });
}

void* QQuickItem_MapToScene(void* ptr, void* point)
{
		return ({ QPointF tmpValue = static_cast<QQuickItem*>(ptr)->mapToScene(*static_cast<QPointF*>(point)); new QPointF(tmpValue.x(), tmpValue.y()); });
}

void* QQuickItem_ChildAt(void* ptr, double x, double y)
{
		return static_cast<QQuickItem*>(ptr)->childAt(x, y);
}

void* QQuickItem_ParentItem(void* ptr)
{
		return static_cast<QQuickItem*>(ptr)->parentItem();
}

void* QQuickItem_ScopedFocusItem(void* ptr)
{
		return static_cast<QQuickItem*>(ptr)->scopedFocusItem();
}

void* QQuickItem_Window(void* ptr)
{
		return static_cast<QQuickItem*>(ptr)->window();
}

void* QQuickItem_MapRectFromItem(void* ptr, void* item, void* rect)
{
		return ({ QRectF tmpValue = static_cast<QQuickItem*>(ptr)->mapRectFromItem(static_cast<QQuickItem*>(item), *static_cast<QRectF*>(rect)); new QRectF(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void* QQuickItem_MapRectFromScene(void* ptr, void* rect)
{
		return ({ QRectF tmpValue = static_cast<QQuickItem*>(ptr)->mapRectFromScene(*static_cast<QRectF*>(rect)); new QRectF(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void* QQuickItem_MapRectToItem(void* ptr, void* item, void* rect)
{
		return ({ QRectF tmpValue = static_cast<QQuickItem*>(ptr)->mapRectToItem(static_cast<QQuickItem*>(item), *static_cast<QRectF*>(rect)); new QRectF(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void* QQuickItem_MapRectToScene(void* ptr, void* rect)
{
		return ({ QRectF tmpValue = static_cast<QQuickItem*>(ptr)->mapRectToScene(*static_cast<QRectF*>(rect)); new QRectF(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void* QQuickItem_TextureProvider(void* ptr)
{
		return static_cast<QQuickItem*>(ptr)->textureProvider();
}

void* QQuickItem_TextureProviderDefault(void* ptr)
{
	if (dynamic_cast<QQuickPaintedItem*>(static_cast<QObject*>(ptr))) {
		return static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::textureProvider();
	} else if (dynamic_cast<QQuickFramebufferObject*>(static_cast<QObject*>(ptr))) {
		return static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::textureProvider();
	} else {
		return static_cast<QQuickItem*>(ptr)->QQuickItem::textureProvider();
	}
}

struct QtQuick_PackedString QQuickItem_State(void* ptr)
{
		return ({ QByteArray t803f42 = static_cast<QQuickItem*>(ptr)->state().toUtf8(); QtQuick_PackedString { const_cast<char*>(t803f42.prepend("WHITESPACE").constData()+10), t803f42.size()-10 }; });
}

void* QQuickItem_InputMethodQuery(void* ptr, long long query)
{
		return new QVariant(static_cast<QQuickItem*>(ptr)->inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

void* QQuickItem_InputMethodQueryDefault(void* ptr, long long query)
{
	if (dynamic_cast<QQuickPaintedItem*>(static_cast<QObject*>(ptr))) {
		return new QVariant(static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
	} else if (dynamic_cast<QQuickFramebufferObject*>(static_cast<QObject*>(ptr))) {
		return new QVariant(static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
	} else {
		return new QVariant(static_cast<QQuickItem*>(ptr)->QQuickItem::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
	}
}

long long QQuickItem_AcceptedMouseButtons(void* ptr)
{
		return static_cast<QQuickItem*>(ptr)->acceptedMouseButtons();
}

long long QQuickItem_TransformOrigin(void* ptr)
{
		return static_cast<QQuickItem*>(ptr)->transformOrigin();
}

char QQuickItem_AcceptHoverEvents(void* ptr)
{
		return static_cast<QQuickItem*>(ptr)->acceptHoverEvents();
}

char QQuickItem_ActiveFocusOnTab(void* ptr)
{
		return static_cast<QQuickItem*>(ptr)->activeFocusOnTab();
}

char QQuickItem_Antialiasing(void* ptr)
{
		return static_cast<QQuickItem*>(ptr)->antialiasing();
}

char QQuickItem_Clip(void* ptr)
{
		return static_cast<QQuickItem*>(ptr)->clip();
}

char QQuickItem_Contains(void* ptr, void* point)
{
		return static_cast<QQuickItem*>(ptr)->contains(*static_cast<QPointF*>(point));
}

char QQuickItem_ContainsDefault(void* ptr, void* point)
{
	if (dynamic_cast<QQuickPaintedItem*>(static_cast<QObject*>(ptr))) {
		return static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::contains(*static_cast<QPointF*>(point));
	} else if (dynamic_cast<QQuickFramebufferObject*>(static_cast<QObject*>(ptr))) {
		return static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::contains(*static_cast<QPointF*>(point));
	} else {
		return static_cast<QQuickItem*>(ptr)->QQuickItem::contains(*static_cast<QPointF*>(point));
	}
}

char QQuickItem_FiltersChildMouseEvents(void* ptr)
{
		return static_cast<QQuickItem*>(ptr)->filtersChildMouseEvents();
}

char QQuickItem_HasActiveFocus(void* ptr)
{
		return static_cast<QQuickItem*>(ptr)->hasActiveFocus();
}

char QQuickItem_HasFocus(void* ptr)
{
		return static_cast<QQuickItem*>(ptr)->hasFocus();
}

char QQuickItem_HeightValid(void* ptr)
{
		return static_cast<QQuickItem*>(ptr)->heightValid();
}

char QQuickItem_IsAncestorOf(void* ptr, void* child)
{
		return static_cast<QQuickItem*>(ptr)->isAncestorOf(static_cast<QQuickItem*>(child));
}

char QQuickItem_IsComponentComplete(void* ptr)
{
		return static_cast<QQuickItem*>(ptr)->isComponentComplete();
}

char QQuickItem_IsEnabled(void* ptr)
{
		return static_cast<QQuickItem*>(ptr)->isEnabled();
}

char QQuickItem_IsFocusScope(void* ptr)
{
		return static_cast<QQuickItem*>(ptr)->isFocusScope();
}

char QQuickItem_IsTextureProvider(void* ptr)
{
		return static_cast<QQuickItem*>(ptr)->isTextureProvider();
}

char QQuickItem_IsTextureProviderDefault(void* ptr)
{
	if (dynamic_cast<QQuickPaintedItem*>(static_cast<QObject*>(ptr))) {
		return static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::isTextureProvider();
	} else if (dynamic_cast<QQuickFramebufferObject*>(static_cast<QObject*>(ptr))) {
		return static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::isTextureProvider();
	} else {
		return static_cast<QQuickItem*>(ptr)->QQuickItem::isTextureProvider();
	}
}

char QQuickItem_IsVisible(void* ptr)
{
		return static_cast<QQuickItem*>(ptr)->isVisible();
}

char QQuickItem_KeepMouseGrab(void* ptr)
{
		return static_cast<QQuickItem*>(ptr)->keepMouseGrab();
}

char QQuickItem_KeepTouchGrab(void* ptr)
{
		return static_cast<QQuickItem*>(ptr)->keepTouchGrab();
}

char QQuickItem_Smooth(void* ptr)
{
		return static_cast<QQuickItem*>(ptr)->smooth();
}

char QQuickItem_WidthValid(void* ptr)
{
		return static_cast<QQuickItem*>(ptr)->widthValid();
}

double QQuickItem_BaselineOffset(void* ptr)
{
		return static_cast<QQuickItem*>(ptr)->baselineOffset();
}

double QQuickItem_Height(void* ptr)
{
		return static_cast<QQuickItem*>(ptr)->height();
}

double QQuickItem_ImplicitHeight(void* ptr)
{
		return static_cast<QQuickItem*>(ptr)->implicitHeight();
}

double QQuickItem_ImplicitWidth(void* ptr)
{
		return static_cast<QQuickItem*>(ptr)->implicitWidth();
}

double QQuickItem_Opacity(void* ptr)
{
		return static_cast<QQuickItem*>(ptr)->opacity();
}

double QQuickItem_Rotation(void* ptr)
{
		return static_cast<QQuickItem*>(ptr)->rotation();
}

double QQuickItem_Scale(void* ptr)
{
		return static_cast<QQuickItem*>(ptr)->scale();
}

double QQuickItem_Width(void* ptr)
{
		return static_cast<QQuickItem*>(ptr)->width();
}

double QQuickItem_X(void* ptr)
{
		return static_cast<QQuickItem*>(ptr)->x();
}

double QQuickItem_Y(void* ptr)
{
		return static_cast<QQuickItem*>(ptr)->y();
}

double QQuickItem_Z(void* ptr)
{
		return static_cast<QQuickItem*>(ptr)->z();
}

int QQuickItem___grabTouchPoints_ids_atList(void* ptr, int i)
{
		return static_cast<QVector<int>*>(ptr)->at(i);
}

void QQuickItem___grabTouchPoints_ids_setList(void* ptr, int i)
{
		static_cast<QVector<int>*>(ptr)->append(i);
}

void* QQuickItem___grabTouchPoints_ids_newList(void* ptr)
{
	Q_UNUSED(ptr);
		return new QVector<int>;
}

void* QQuickItem___childItems_atList(void* ptr, int i)
{
		return const_cast<QQuickItem*>(static_cast<QList<QQuickItem *>*>(ptr)->at(i));
}

void QQuickItem___childItems_setList(void* ptr, void* i)
{
		static_cast<QList<QQuickItem *>*>(ptr)->append(static_cast<QQuickItem*>(i));
}

void* QQuickItem___childItems_newList(void* ptr)
{
	Q_UNUSED(ptr);
		return new QList<QQuickItem *>;
}

void* QQuickItem___dynamicPropertyNames_atList(void* ptr, int i)
{
		return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QQuickItem___dynamicPropertyNames_setList(void* ptr, void* i)
{
		static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QQuickItem___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
		return new QList<QByteArray>;
}

void* QQuickItem___findChildren_atList2(void* ptr, int i)
{
		return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QQuickItem___findChildren_setList2(void* ptr, void* i)
{
	if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QQuickItem*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QQuickItem___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
		return new QList<QObject*>;
}

void* QQuickItem___findChildren_atList3(void* ptr, int i)
{
		return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QQuickItem___findChildren_setList3(void* ptr, void* i)
{
	if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QQuickItem*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QQuickItem___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
		return new QList<QObject*>;
}

void* QQuickItem___findChildren_atList(void* ptr, int i)
{
		return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QQuickItem___findChildren_setList(void* ptr, void* i)
{
	if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QQuickItem*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QQuickItem___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
		return new QList<QObject*>;
}

void* QQuickItem___children_atList(void* ptr, int i)
{
		return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QQuickItem___children_setList(void* ptr, void* i)
{
	if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QQuickItem*>(i));
	} else {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QQuickItem___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
		return new QList<QObject *>;
}

char QQuickItem_EventFilter(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(watched))) {
		return static_cast<QQuickItem*>(ptr)->eventFilter(static_cast<QQuickItem*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QQuickItem*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

char QQuickItem_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QQuickPaintedItem*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(watched))) {
			return static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::eventFilter(static_cast<QQuickItem*>(watched), static_cast<QEvent*>(event));
		} else {
			return static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
		}
	} else if (dynamic_cast<QQuickFramebufferObject*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(watched))) {
			return static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::eventFilter(static_cast<QQuickItem*>(watched), static_cast<QEvent*>(event));
		} else {
			return static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
		}
	} else {
		if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(watched))) {
			return static_cast<QQuickItem*>(ptr)->QQuickItem::eventFilter(static_cast<QQuickItem*>(watched), static_cast<QEvent*>(event));
		} else {
			return static_cast<QQuickItem*>(ptr)->QQuickItem::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
		}
	}
}

void QQuickItem_ChildEvent(void* ptr, void* event)
{
		static_cast<QQuickItem*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QQuickItem_ChildEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QQuickPaintedItem*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QQuickFramebufferObject*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::childEvent(static_cast<QChildEvent*>(event));
	} else {
		static_cast<QQuickItem*>(ptr)->QQuickItem::childEvent(static_cast<QChildEvent*>(event));
	}
}

void QQuickItem_ConnectNotify(void* ptr, void* sign)
{
		static_cast<QQuickItem*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQuickItem_ConnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QQuickPaintedItem*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QQuickFramebufferObject*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QQuickItem*>(ptr)->QQuickItem::connectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

void QQuickItem_CustomEvent(void* ptr, void* event)
{
		static_cast<QQuickItem*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QQuickItem_CustomEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QQuickPaintedItem*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QQuickFramebufferObject*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::customEvent(static_cast<QEvent*>(event));
	} else {
		static_cast<QQuickItem*>(ptr)->QQuickItem::customEvent(static_cast<QEvent*>(event));
	}
}

void QQuickItem_DeleteLater(void* ptr)
{
		QMetaObject::invokeMethod(static_cast<QQuickItem*>(ptr), "deleteLater");
}

void QQuickItem_DeleteLaterDefault(void* ptr)
{
	if (dynamic_cast<QQuickPaintedItem*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::deleteLater();
	} else if (dynamic_cast<QQuickFramebufferObject*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::deleteLater();
	} else {
		static_cast<QQuickItem*>(ptr)->QQuickItem::deleteLater();
	}
}

void QQuickItem_DisconnectNotify(void* ptr, void* sign)
{
		static_cast<QQuickItem*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQuickItem_DisconnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QQuickPaintedItem*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QQuickFramebufferObject*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QQuickItem*>(ptr)->QQuickItem::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

void QQuickItem_TimerEvent(void* ptr, void* event)
{
		static_cast<QQuickItem*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QQuickItem_TimerEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QQuickPaintedItem*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QQuickFramebufferObject*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::timerEvent(static_cast<QTimerEvent*>(event));
	} else {
		static_cast<QQuickItem*>(ptr)->QQuickItem::timerEvent(static_cast<QTimerEvent*>(event));
	}
}

void* QQuickItem_MetaObject(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QQuickItem*>(ptr)->metaObject());
}

void* QQuickItem_MetaObjectDefault(void* ptr)
{
	if (dynamic_cast<QQuickPaintedItem*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::metaObject());
	} else if (dynamic_cast<QQuickFramebufferObject*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::metaObject());
	} else {
		return const_cast<QMetaObject*>(static_cast<QQuickItem*>(ptr)->QQuickItem::metaObject());
	}
}

class MyQQuickItemGrabResult: public QQuickItemGrabResult
{
public:
	void Signal_Ready() { callbackQQuickItemGrabResult_Ready(this); };
	bool event(QEvent * e) { return callbackQQuickItemGrabResult_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQQuickItemGrabResult_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQQuickItemGrabResult_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQQuickItemGrabResult_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQQuickItemGrabResult_CustomEvent(this, event); };
	void deleteLater() { callbackQQuickItemGrabResult_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQQuickItemGrabResult_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQQuickItemGrabResult_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtQuick_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQQuickItemGrabResult_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQQuickItemGrabResult_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQQuickItemGrabResult_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQQuickItemGrabResult*)

int QQuickItemGrabResult_QQuickItemGrabResult_QRegisterMetaType(){qRegisterMetaType<QQuickItemGrabResult*>(); return qRegisterMetaType<MyQQuickItemGrabResult*>();}

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

void* QQuickItemGrabResult_Image(void* ptr)
{
	return new QImage(static_cast<QQuickItemGrabResult*>(ptr)->image());
}

void* QQuickItemGrabResult_Url(void* ptr)
{
	return new QUrl(static_cast<QQuickItemGrabResult*>(ptr)->url());
}

char QQuickItemGrabResult_SaveToFile(void* ptr, struct QtQuick_PackedString fileName)
{
	return static_cast<QQuickItemGrabResult*>(ptr)->saveToFile(QString::fromUtf8(fileName.data, fileName.len));
}

void* QQuickItemGrabResult___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QQuickItemGrabResult___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QQuickItemGrabResult___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QQuickItemGrabResult___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QQuickItemGrabResult___findChildren_setList2(void* ptr, void* i)
{
	if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QQuickItem*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QQuickItemGrabResult___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QQuickItemGrabResult___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QQuickItemGrabResult___findChildren_setList3(void* ptr, void* i)
{
	if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QQuickItem*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QQuickItemGrabResult___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QQuickItemGrabResult___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QQuickItemGrabResult___findChildren_setList(void* ptr, void* i)
{
	if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QQuickItem*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QQuickItemGrabResult___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QQuickItemGrabResult___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QQuickItemGrabResult___children_setList(void* ptr, void* i)
{
	if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QQuickItem*>(i));
	} else {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QQuickItemGrabResult___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

char QQuickItemGrabResult_EventDefault(void* ptr, void* e)
{
		return static_cast<QQuickItemGrabResult*>(ptr)->QQuickItemGrabResult::event(static_cast<QEvent*>(e));
}

char QQuickItemGrabResult_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(watched))) {
		return static_cast<QQuickItemGrabResult*>(ptr)->QQuickItemGrabResult::eventFilter(static_cast<QQuickItem*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QQuickItemGrabResult*>(ptr)->QQuickItemGrabResult::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

void QQuickItemGrabResult_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QQuickItemGrabResult*>(ptr)->QQuickItemGrabResult::childEvent(static_cast<QChildEvent*>(event));
}

void QQuickItemGrabResult_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QQuickItemGrabResult*>(ptr)->QQuickItemGrabResult::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQuickItemGrabResult_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QQuickItemGrabResult*>(ptr)->QQuickItemGrabResult::customEvent(static_cast<QEvent*>(event));
}

void QQuickItemGrabResult_DeleteLaterDefault(void* ptr)
{
		static_cast<QQuickItemGrabResult*>(ptr)->QQuickItemGrabResult::deleteLater();
}

void QQuickItemGrabResult_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QQuickItemGrabResult*>(ptr)->QQuickItemGrabResult::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQuickItemGrabResult_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QQuickItemGrabResult*>(ptr)->QQuickItemGrabResult::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QQuickItemGrabResult_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QQuickItemGrabResult*>(ptr)->QQuickItemGrabResult::metaObject());
}

class MyQQuickPaintedItem: public QQuickPaintedItem
{
public:
	MyQQuickPaintedItem(QQuickItem *parent = Q_NULLPTR) : QQuickPaintedItem(parent) {QQuickPaintedItem_QQuickPaintedItem_QRegisterMetaType();};
	void Signal_ContentsScaleChanged() { callbackQQuickPaintedItem_ContentsScaleChanged(this); };
	void Signal_ContentsSizeChanged() { callbackQQuickPaintedItem_ContentsSizeChanged(this); };
	void Signal_FillColorChanged() { callbackQQuickPaintedItem_FillColorChanged(this); };
	void paint(QPainter * painter) { callbackQQuickPaintedItem_Paint(this, painter); };
	void releaseResources() { callbackQQuickItem_ReleaseResources(this); };
	void Signal_RenderTargetChanged() { callbackQQuickPaintedItem_RenderTargetChanged(this); };
	void Signal_TextureSizeChanged() { callbackQQuickPaintedItem_TextureSizeChanged(this); };
	 ~MyQQuickPaintedItem() { callbackQQuickPaintedItem_DestroyQQuickPaintedItem(this); };
	QSGTextureProvider * textureProvider() const { return static_cast<QSGTextureProvider*>(callbackQQuickItem_TextureProvider(const_cast<void*>(static_cast<const void*>(this)))); };
	bool isTextureProvider() const { return callbackQQuickItem_IsTextureProvider(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	bool childMouseEventFilter(QQuickItem * item, QEvent * event) { return callbackQQuickItem_ChildMouseEventFilter(this, item, event) != 0; };
	bool event(QEvent * ev) { return callbackQQuickItem_Event(this, ev) != 0; };
	void classBegin() { callbackQQuickItem_ClassBegin(this); };
	void componentComplete() { callbackQQuickItem_ComponentComplete(this); };
	void dragEnterEvent(QDragEnterEvent * event) { callbackQQuickItem_DragEnterEvent(this, event); };
	void dragLeaveEvent(QDragLeaveEvent * event) { callbackQQuickItem_DragLeaveEvent(this, event); };
	void dragMoveEvent(QDragMoveEvent * event) { callbackQQuickItem_DragMoveEvent(this, event); };
	void dropEvent(QDropEvent * event) { callbackQQuickItem_DropEvent(this, event); };
	void focusInEvent(QFocusEvent * event) { callbackQQuickItem_FocusInEvent(this, event); };
	void focusOutEvent(QFocusEvent * event) { callbackQQuickItem_FocusOutEvent(this, event); };
	void geometryChanged(const QRectF & newGeometry, const QRectF & oldGeometry) { callbackQQuickItem_GeometryChanged(this, const_cast<QRectF*>(&newGeometry), const_cast<QRectF*>(&oldGeometry)); };
	void hoverEnterEvent(QHoverEvent * event) { callbackQQuickItem_HoverEnterEvent(this, event); };
	void hoverLeaveEvent(QHoverEvent * event) { callbackQQuickItem_HoverLeaveEvent(this, event); };
	void hoverMoveEvent(QHoverEvent * event) { callbackQQuickItem_HoverMoveEvent(this, event); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackQQuickItem_InputMethodEvent(this, event); };
	void keyPressEvent(QKeyEvent * event) { callbackQQuickItem_KeyPressEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQQuickItem_KeyReleaseEvent(this, event); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQQuickItem_MouseDoubleClickEvent(this, event); };
	void mouseMoveEvent(QMouseEvent * event) { callbackQQuickItem_MouseMoveEvent(this, event); };
	void mousePressEvent(QMouseEvent * event) { callbackQQuickItem_MousePressEvent(this, event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackQQuickItem_MouseReleaseEvent(this, event); };
	void mouseUngrabEvent() { callbackQQuickItem_MouseUngrabEvent(this); };
	void touchEvent(QTouchEvent * event) { callbackQQuickItem_TouchEvent(this, event); };
	void touchUngrabEvent() { callbackQQuickItem_TouchUngrabEvent(this); };
	void updatePolish() { callbackQQuickItem_UpdatePolish(this); };
	void wheelEvent(QWheelEvent * event) { callbackQQuickItem_WheelEvent(this, event); };
	void Signal_WindowChanged(QQuickWindow * window) { callbackQQuickItem_WindowChanged(this, window); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackQQuickItem_InputMethodQuery(const_cast<void*>(static_cast<const void*>(this)), query)); };
	bool contains(const QPointF & point) const { return callbackQQuickItem_Contains(const_cast<void*>(static_cast<const void*>(this)), const_cast<QPointF*>(&point)) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQQuickItem_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQQuickItem_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQQuickItem_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQQuickItem_CustomEvent(this, event); };
	void deleteLater() { callbackQQuickItem_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQQuickItem_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQQuickItem_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtQuick_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQQuickItem_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQQuickItem_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQQuickItem_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQQuickPaintedItem*)

int QQuickPaintedItem_QQuickPaintedItem_QRegisterMetaType(){qRegisterMetaType<QQuickPaintedItem*>(); return qRegisterMetaType<MyQQuickPaintedItem*>();}

void* QQuickPaintedItem_NewQQuickPaintedItem(void* parent)
{
		return new MyQQuickPaintedItem(static_cast<QQuickItem*>(parent));
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

void QQuickPaintedItem_Paint(void* ptr, void* painter)
{
	static_cast<QQuickPaintedItem*>(ptr)->paint(static_cast<QPainter*>(painter));
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

void QQuickPaintedItem_SetContentsScale(void* ptr, double vqr)
{
	static_cast<QQuickPaintedItem*>(ptr)->setContentsScale(vqr);
}

void QQuickPaintedItem_SetContentsSize(void* ptr, void* vqs)
{
	static_cast<QQuickPaintedItem*>(ptr)->setContentsSize(*static_cast<QSize*>(vqs));
}

void QQuickPaintedItem_SetFillColor(void* ptr, void* vqc)
{
	static_cast<QQuickPaintedItem*>(ptr)->setFillColor(*static_cast<QColor*>(vqc));
}

void QQuickPaintedItem_SetMipmap(void* ptr, char enable)
{
	static_cast<QQuickPaintedItem*>(ptr)->setMipmap(enable != 0);
}

void QQuickPaintedItem_SetOpaquePainting(void* ptr, char opaque)
{
	static_cast<QQuickPaintedItem*>(ptr)->setOpaquePainting(opaque != 0);
}

void QQuickPaintedItem_SetPerformanceHint(void* ptr, long long hint, char enabled)
{
	static_cast<QQuickPaintedItem*>(ptr)->setPerformanceHint(static_cast<QQuickPaintedItem::PerformanceHint>(hint), enabled != 0);
}

void QQuickPaintedItem_SetPerformanceHints(void* ptr, long long hints)
{
	static_cast<QQuickPaintedItem*>(ptr)->setPerformanceHints(static_cast<QQuickPaintedItem::PerformanceHint>(hints));
}

void QQuickPaintedItem_SetRenderTarget(void* ptr, long long target)
{
	static_cast<QQuickPaintedItem*>(ptr)->setRenderTarget(static_cast<QQuickPaintedItem::RenderTarget>(target));
}

void QQuickPaintedItem_SetTextureSize(void* ptr, void* size)
{
	static_cast<QQuickPaintedItem*>(ptr)->setTextureSize(*static_cast<QSize*>(size));
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

void QQuickPaintedItem_DestroyQQuickPaintedItemDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

long long QQuickPaintedItem_PerformanceHints(void* ptr)
{
	return static_cast<QQuickPaintedItem*>(ptr)->performanceHints();
}

void* QQuickPaintedItem_FillColor(void* ptr)
{
	return new QColor(static_cast<QQuickPaintedItem*>(ptr)->fillColor());
}

void* QQuickPaintedItem_ContentsSize(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QQuickPaintedItem*>(ptr)->contentsSize(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QQuickPaintedItem_TextureSize(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QQuickPaintedItem*>(ptr)->textureSize(); new QSize(tmpValue.width(), tmpValue.height()); });
}

long long QQuickPaintedItem_RenderTarget(void* ptr)
{
	return static_cast<QQuickPaintedItem*>(ptr)->renderTarget();
}

char QQuickPaintedItem_Mipmap(void* ptr)
{
	return static_cast<QQuickPaintedItem*>(ptr)->mipmap();
}

char QQuickPaintedItem_OpaquePainting(void* ptr)
{
	return static_cast<QQuickPaintedItem*>(ptr)->opaquePainting();
}

double QQuickPaintedItem_ContentsScale(void* ptr)
{
	return static_cast<QQuickPaintedItem*>(ptr)->contentsScale();
}

class MyQQuickRenderControl: public QQuickRenderControl
{
public:
	MyQQuickRenderControl(QObject *parent = Q_NULLPTR) : QQuickRenderControl(parent) {QQuickRenderControl_QQuickRenderControl_QRegisterMetaType();};
	QWindow * renderWindow(QPoint * offset) { return static_cast<QWindow*>(callbackQQuickRenderControl_RenderWindow(this, offset)); };
	void Signal_RenderRequested() { callbackQQuickRenderControl_RenderRequested(this); };
	void Signal_SceneChanged() { callbackQQuickRenderControl_SceneChanged(this); };
	bool event(QEvent * e) { return callbackQQuickRenderControl_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQQuickRenderControl_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQQuickRenderControl_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQQuickRenderControl_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQQuickRenderControl_CustomEvent(this, event); };
	void deleteLater() { callbackQQuickRenderControl_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQQuickRenderControl_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQQuickRenderControl_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtQuick_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQQuickRenderControl_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQQuickRenderControl_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQQuickRenderControl_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQQuickRenderControl*)

int QQuickRenderControl_QQuickRenderControl_QRegisterMetaType(){qRegisterMetaType<QQuickRenderControl*>(); return qRegisterMetaType<MyQQuickRenderControl*>();}

void* QQuickRenderControl_Grab(void* ptr)
{
	return new QImage(static_cast<QQuickRenderControl*>(ptr)->grab());
}

void* QQuickRenderControl_NewQQuickRenderControl(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQQuickRenderControl(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQQuickRenderControl(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQQuickRenderControl(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQQuickRenderControl(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQQuickRenderControl(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQQuickRenderControl(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQQuickRenderControl(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQQuickRenderControl(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQQuickRenderControl(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQQuickRenderControl(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQQuickRenderControl(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQQuickRenderControl(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQQuickRenderControl(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQQuickRenderControl(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQQuickRenderControl(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQQuickRenderControl(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQQuickRenderControl(static_cast<QWindow*>(parent));
	} else {
		return new MyQQuickRenderControl(static_cast<QObject*>(parent));
	}
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

char QQuickRenderControl_Sync(void* ptr)
{
	return static_cast<QQuickRenderControl*>(ptr)->sync();
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

void QQuickRenderControl_DestroyQQuickRenderControl(void* ptr)
{
	static_cast<QQuickRenderControl*>(ptr)->~QQuickRenderControl();
}

void* QQuickRenderControl___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QQuickRenderControl___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QQuickRenderControl___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QQuickRenderControl___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QQuickRenderControl___findChildren_setList2(void* ptr, void* i)
{
	if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QQuickItem*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QQuickRenderControl___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QQuickRenderControl___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QQuickRenderControl___findChildren_setList3(void* ptr, void* i)
{
	if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QQuickItem*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QQuickRenderControl___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QQuickRenderControl___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QQuickRenderControl___findChildren_setList(void* ptr, void* i)
{
	if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QQuickItem*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QQuickRenderControl___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QQuickRenderControl___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QQuickRenderControl___children_setList(void* ptr, void* i)
{
	if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QQuickItem*>(i));
	} else {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QQuickRenderControl___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

char QQuickRenderControl_EventDefault(void* ptr, void* e)
{
		return static_cast<QQuickRenderControl*>(ptr)->QQuickRenderControl::event(static_cast<QEvent*>(e));
}

char QQuickRenderControl_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(watched))) {
		return static_cast<QQuickRenderControl*>(ptr)->QQuickRenderControl::eventFilter(static_cast<QQuickItem*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QQuickRenderControl*>(ptr)->QQuickRenderControl::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

void QQuickRenderControl_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QQuickRenderControl*>(ptr)->QQuickRenderControl::childEvent(static_cast<QChildEvent*>(event));
}

void QQuickRenderControl_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QQuickRenderControl*>(ptr)->QQuickRenderControl::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQuickRenderControl_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QQuickRenderControl*>(ptr)->QQuickRenderControl::customEvent(static_cast<QEvent*>(event));
}

void QQuickRenderControl_DeleteLaterDefault(void* ptr)
{
		static_cast<QQuickRenderControl*>(ptr)->QQuickRenderControl::deleteLater();
}

void QQuickRenderControl_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QQuickRenderControl*>(ptr)->QQuickRenderControl::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQuickRenderControl_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QQuickRenderControl*>(ptr)->QQuickRenderControl::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QQuickRenderControl_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QQuickRenderControl*>(ptr)->QQuickRenderControl::metaObject());
}

class MyQQuickTextDocument: public QQuickTextDocument
{
public:
	MyQQuickTextDocument(QQuickItem *parent) : QQuickTextDocument(parent) {QQuickTextDocument_QQuickTextDocument_QRegisterMetaType();};
	bool event(QEvent * e) { return callbackQQuickTextDocument_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQQuickTextDocument_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQQuickTextDocument_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQQuickTextDocument_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQQuickTextDocument_CustomEvent(this, event); };
	void deleteLater() { callbackQQuickTextDocument_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQQuickTextDocument_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQQuickTextDocument_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtQuick_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQQuickTextDocument_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQQuickTextDocument_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQQuickTextDocument_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQQuickTextDocument*)

int QQuickTextDocument_QQuickTextDocument_QRegisterMetaType(){qRegisterMetaType<QQuickTextDocument*>(); return qRegisterMetaType<MyQQuickTextDocument*>();}

void* QQuickTextDocument_NewQQuickTextDocument(void* parent)
{
		return new MyQQuickTextDocument(static_cast<QQuickItem*>(parent));
}

void* QQuickTextDocument_TextDocument(void* ptr)
{
	return static_cast<QQuickTextDocument*>(ptr)->textDocument();
}

void* QQuickTextDocument___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QQuickTextDocument___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QQuickTextDocument___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QQuickTextDocument___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QQuickTextDocument___findChildren_setList2(void* ptr, void* i)
{
	if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QQuickItem*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QQuickTextDocument___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QQuickTextDocument___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QQuickTextDocument___findChildren_setList3(void* ptr, void* i)
{
	if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QQuickItem*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QQuickTextDocument___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QQuickTextDocument___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QQuickTextDocument___findChildren_setList(void* ptr, void* i)
{
	if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QQuickItem*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QQuickTextDocument___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QQuickTextDocument___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QQuickTextDocument___children_setList(void* ptr, void* i)
{
	if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QQuickItem*>(i));
	} else {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QQuickTextDocument___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

char QQuickTextDocument_EventDefault(void* ptr, void* e)
{
		return static_cast<QQuickTextDocument*>(ptr)->QQuickTextDocument::event(static_cast<QEvent*>(e));
}

char QQuickTextDocument_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(watched))) {
		return static_cast<QQuickTextDocument*>(ptr)->QQuickTextDocument::eventFilter(static_cast<QQuickItem*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QQuickTextDocument*>(ptr)->QQuickTextDocument::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

void QQuickTextDocument_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QQuickTextDocument*>(ptr)->QQuickTextDocument::childEvent(static_cast<QChildEvent*>(event));
}

void QQuickTextDocument_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QQuickTextDocument*>(ptr)->QQuickTextDocument::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQuickTextDocument_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QQuickTextDocument*>(ptr)->QQuickTextDocument::customEvent(static_cast<QEvent*>(event));
}

void QQuickTextDocument_DeleteLaterDefault(void* ptr)
{
		static_cast<QQuickTextDocument*>(ptr)->QQuickTextDocument::deleteLater();
}

void QQuickTextDocument_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QQuickTextDocument*>(ptr)->QQuickTextDocument::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQuickTextDocument_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QQuickTextDocument*>(ptr)->QQuickTextDocument::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QQuickTextDocument_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QQuickTextDocument*>(ptr)->QQuickTextDocument::metaObject());
}

class MyQQuickTextureFactory: public QQuickTextureFactory
{
public:
	MyQQuickTextureFactory() : QQuickTextureFactory() {QQuickTextureFactory_QQuickTextureFactory_QRegisterMetaType();};
	 ~MyQQuickTextureFactory() { callbackQQuickTextureFactory_DestroyQQuickTextureFactory(this); };
	QImage image() const { return *static_cast<QImage*>(callbackQQuickTextureFactory_Image(const_cast<void*>(static_cast<const void*>(this)))); };
	QSGTexture * createTexture(QQuickWindow * window) const { return static_cast<QSGTexture*>(callbackQQuickTextureFactory_CreateTexture(const_cast<void*>(static_cast<const void*>(this)), window)); };
	QSize textureSize() const { return *static_cast<QSize*>(callbackQQuickTextureFactory_TextureSize(const_cast<void*>(static_cast<const void*>(this)))); };
	int textureByteCount() const { return callbackQQuickTextureFactory_TextureByteCount(const_cast<void*>(static_cast<const void*>(this))); };
	bool event(QEvent * e) { return callbackQQuickTextureFactory_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQQuickTextureFactory_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQQuickTextureFactory_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQQuickTextureFactory_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQQuickTextureFactory_CustomEvent(this, event); };
	void deleteLater() { callbackQQuickTextureFactory_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQQuickTextureFactory_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQQuickTextureFactory_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtQuick_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQQuickTextureFactory_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQQuickTextureFactory_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQQuickTextureFactory_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQQuickTextureFactory*)

int QQuickTextureFactory_QQuickTextureFactory_QRegisterMetaType(){qRegisterMetaType<QQuickTextureFactory*>(); return qRegisterMetaType<MyQQuickTextureFactory*>();}

void* QQuickTextureFactory_QQuickTextureFactory_TextureFactoryForImage(void* image)
{
	return QQuickTextureFactory::textureFactoryForImage(*static_cast<QImage*>(image));
}

void* QQuickTextureFactory_NewQQuickTextureFactory()
{
	return new MyQQuickTextureFactory();
}

void QQuickTextureFactory_DestroyQQuickTextureFactory(void* ptr)
{
	static_cast<QQuickTextureFactory*>(ptr)->~QQuickTextureFactory();
}

void QQuickTextureFactory_DestroyQQuickTextureFactoryDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QQuickTextureFactory_Image(void* ptr)
{
	return new QImage(static_cast<QQuickTextureFactory*>(ptr)->image());
}

void* QQuickTextureFactory_ImageDefault(void* ptr)
{
		return new QImage(static_cast<QQuickTextureFactory*>(ptr)->QQuickTextureFactory::image());
}

void* QQuickTextureFactory_CreateTexture(void* ptr, void* window)
{
	return static_cast<QQuickTextureFactory*>(ptr)->createTexture(static_cast<QQuickWindow*>(window));
}

void* QQuickTextureFactory_TextureSize(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QQuickTextureFactory*>(ptr)->textureSize(); new QSize(tmpValue.width(), tmpValue.height()); });
}

int QQuickTextureFactory_TextureByteCount(void* ptr)
{
	return static_cast<QQuickTextureFactory*>(ptr)->textureByteCount();
}

void* QQuickTextureFactory___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QQuickTextureFactory___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QQuickTextureFactory___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QQuickTextureFactory___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QQuickTextureFactory___findChildren_setList2(void* ptr, void* i)
{
	if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QQuickItem*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QQuickTextureFactory___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QQuickTextureFactory___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QQuickTextureFactory___findChildren_setList3(void* ptr, void* i)
{
	if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QQuickItem*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QQuickTextureFactory___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QQuickTextureFactory___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QQuickTextureFactory___findChildren_setList(void* ptr, void* i)
{
	if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QQuickItem*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QQuickTextureFactory___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QQuickTextureFactory___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QQuickTextureFactory___children_setList(void* ptr, void* i)
{
	if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QQuickItem*>(i));
	} else {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QQuickTextureFactory___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

char QQuickTextureFactory_EventDefault(void* ptr, void* e)
{
		return static_cast<QQuickTextureFactory*>(ptr)->QQuickTextureFactory::event(static_cast<QEvent*>(e));
}

char QQuickTextureFactory_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(watched))) {
		return static_cast<QQuickTextureFactory*>(ptr)->QQuickTextureFactory::eventFilter(static_cast<QQuickItem*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QQuickTextureFactory*>(ptr)->QQuickTextureFactory::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

void QQuickTextureFactory_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QQuickTextureFactory*>(ptr)->QQuickTextureFactory::childEvent(static_cast<QChildEvent*>(event));
}

void QQuickTextureFactory_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QQuickTextureFactory*>(ptr)->QQuickTextureFactory::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQuickTextureFactory_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QQuickTextureFactory*>(ptr)->QQuickTextureFactory::customEvent(static_cast<QEvent*>(event));
}

void QQuickTextureFactory_DeleteLaterDefault(void* ptr)
{
		static_cast<QQuickTextureFactory*>(ptr)->QQuickTextureFactory::deleteLater();
}

void QQuickTextureFactory_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QQuickTextureFactory*>(ptr)->QQuickTextureFactory::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQuickTextureFactory_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QQuickTextureFactory*>(ptr)->QQuickTextureFactory::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QQuickTextureFactory_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QQuickTextureFactory*>(ptr)->QQuickTextureFactory::metaObject());
}

class MyQQuickView: public QQuickView
{
public:
	MyQQuickView(QQmlEngine *engine, QWindow *parent) : QQuickView(engine, parent) {QQuickView_QQuickView_QRegisterMetaType();};
	MyQQuickView(QWindow *parent = Q_NULLPTR) : QQuickView(parent) {QQuickView_QQuickView_QRegisterMetaType();};
	MyQQuickView(const QUrl &source, QWindow *parent = Q_NULLPTR) : QQuickView(source, parent) {QQuickView_QQuickView_QRegisterMetaType();};
	void setSource(const QUrl & url) { callbackQQuickView_SetSource(this, const_cast<QUrl*>(&url)); };
	void Signal_StatusChanged(QQuickView::Status status) { callbackQQuickView_StatusChanged(this, status); };
	 ~MyQQuickView() { callbackQQuickView_DestroyQQuickView(this); };
	bool event(QEvent * e) { return callbackQQuickWindow_Event(this, e) != 0; };
	void Signal_ActiveFocusItemChanged() { callbackQQuickWindow_ActiveFocusItemChanged(this); };
	void Signal_AfterAnimating() { callbackQQuickWindow_AfterAnimating(this); };
	void Signal_AfterRendering() { callbackQQuickWindow_AfterRendering(this); };
	void Signal_AfterSynchronizing() { callbackQQuickWindow_AfterSynchronizing(this); };
	void Signal_BeforeRendering() { callbackQQuickWindow_BeforeRendering(this); };
	void Signal_BeforeSynchronizing() { callbackQQuickWindow_BeforeSynchronizing(this); };
	void Signal_ColorChanged(const QColor & vqc) { callbackQQuickWindow_ColorChanged(this, const_cast<QColor*>(&vqc)); };
	void exposeEvent(QExposeEvent * vqe) { callbackQQuickWindow_ExposeEvent(this, vqe); };
	void focusInEvent(QFocusEvent * ev) { callbackQQuickWindow_FocusInEvent(this, ev); };
	void focusOutEvent(QFocusEvent * ev) { callbackQQuickWindow_FocusOutEvent(this, ev); };
	void Signal_FrameSwapped() { callbackQQuickWindow_FrameSwapped(this); };
	void hideEvent(QHideEvent * vqh) { callbackQQuickWindow_HideEvent(this, vqh); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQQuickWindow_MouseDoubleClickEvent(this, event); };
	void Signal_OpenglContextCreated(QOpenGLContext * context) { callbackQQuickWindow_OpenglContextCreated(this, context); };
	void releaseResources() { callbackQQuickWindow_ReleaseResources(this); };
	void resizeEvent(QResizeEvent * ev) { callbackQQuickWindow_ResizeEvent(this, ev); };
	void Signal_SceneGraphAboutToStop() { callbackQQuickWindow_SceneGraphAboutToStop(this); };
	void Signal_SceneGraphError(QQuickWindow::SceneGraphError error, const QString & message) { QByteArray t6f9b9a = message.toUtf8(); QtQuick_PackedString messagePacked = { const_cast<char*>(t6f9b9a.prepend("WHITESPACE").constData()+10), t6f9b9a.size()-10 };callbackQQuickWindow_SceneGraphError(this, error, messagePacked); };
	void Signal_SceneGraphInitialized() { callbackQQuickWindow_SceneGraphInitialized(this); };
	void Signal_SceneGraphInvalidated() { callbackQQuickWindow_SceneGraphInvalidated(this); };
	void showEvent(QShowEvent * vqs) { callbackQQuickWindow_ShowEvent(this, vqs); };
	void update() { callbackQQuickWindow_Update(this); };
	void wheelEvent(QWheelEvent * event) { callbackQQuickWindow_WheelEvent(this, event); };
	bool close() { return callbackQQuickWindow_Close(this) != 0; };
	bool nativeEvent(const QByteArray & eventType, void * message, long * result) { return callbackQQuickWindow_NativeEvent(this, const_cast<QByteArray*>(&eventType), message, *result) != 0; };
	void Signal_ActiveChanged() { callbackQQuickWindow_ActiveChanged(this); };
	void alert(int msec) { callbackQQuickWindow_Alert(this, msec); };
	void Signal_ContentOrientationChanged(Qt::ScreenOrientation orientation) { callbackQQuickWindow_ContentOrientationChanged(this, orientation); };
	void Signal_FocusObjectChanged(QObject * object) { callbackQQuickWindow_FocusObjectChanged(this, object); };
	void Signal_HeightChanged(int arg) { callbackQQuickWindow_HeightChanged(this, arg); };
	void hide() { callbackQQuickWindow_Hide(this); };
	void lower() { callbackQQuickWindow_Lower(this); };
	void Signal_MaximumHeightChanged(int arg) { callbackQQuickWindow_MaximumHeightChanged(this, arg); };
	void Signal_MaximumWidthChanged(int arg) { callbackQQuickWindow_MaximumWidthChanged(this, arg); };
	void Signal_MinimumHeightChanged(int arg) { callbackQQuickWindow_MinimumHeightChanged(this, arg); };
	void Signal_MinimumWidthChanged(int arg) { callbackQQuickWindow_MinimumWidthChanged(this, arg); };
	void Signal_ModalityChanged(Qt::WindowModality modality) { callbackQQuickWindow_ModalityChanged(this, modality); };
	void moveEvent(QMoveEvent * ev) { callbackQQuickWindow_MoveEvent(this, ev); };
	void Signal_OpacityChanged(qreal opacity) { callbackQQuickWindow_OpacityChanged(this, opacity); };
	void raise() { callbackQQuickWindow_Raise(this); };
	void requestActivate() { callbackQQuickWindow_RequestActivate(this); };
	void requestUpdate() { callbackQQuickWindow_RequestUpdate(this); };
	void Signal_ScreenChanged(QScreen * screen) { callbackQQuickWindow_ScreenChanged(this, screen); };
	void setHeight(int arg) { callbackQQuickWindow_SetHeight(this, arg); };
	void setMaximumHeight(int h) { callbackQQuickWindow_SetMaximumHeight(this, h); };
	void setMaximumWidth(int w) { callbackQQuickWindow_SetMaximumWidth(this, w); };
	void setMinimumHeight(int h) { callbackQQuickWindow_SetMinimumHeight(this, h); };
	void setMinimumWidth(int w) { callbackQQuickWindow_SetMinimumWidth(this, w); };
	void setTitle(const QString & vqs) { QByteArray tda39a3 = vqs.toUtf8(); QtQuick_PackedString vqsPacked = { const_cast<char*>(tda39a3.prepend("WHITESPACE").constData()+10), tda39a3.size()-10 };callbackQQuickWindow_SetTitle(this, vqsPacked); };
	void setVisible(bool visible) { callbackQQuickWindow_SetVisible(this, visible); };
	void setWidth(int arg) { callbackQQuickWindow_SetWidth(this, arg); };
	void setX(int arg) { callbackQQuickWindow_SetX(this, arg); };
	void setY(int arg) { callbackQQuickWindow_SetY(this, arg); };
	void show() { callbackQQuickWindow_Show(this); };
	void showFullScreen() { callbackQQuickWindow_ShowFullScreen(this); };
	void showMaximized() { callbackQQuickWindow_ShowMaximized(this); };
	void showMinimized() { callbackQQuickWindow_ShowMinimized(this); };
	void showNormal() { callbackQQuickWindow_ShowNormal(this); };
	void tabletEvent(QTabletEvent * ev) { callbackQQuickWindow_TabletEvent(this, ev); };
	void touchEvent(QTouchEvent * ev) { callbackQQuickWindow_TouchEvent(this, ev); };
	void Signal_VisibilityChanged(QWindow::Visibility visibility) { callbackQQuickWindow_VisibilityChanged(this, visibility); };
	void Signal_VisibleChanged(bool arg) { callbackQQuickWindow_VisibleChanged(this, arg); };
	void Signal_WidthChanged(int arg) { callbackQQuickWindow_WidthChanged(this, arg); };
	void Signal_WindowStateChanged(Qt::WindowState windowState) { callbackQQuickWindow_WindowStateChanged(this, windowState); };
	void Signal_WindowTitleChanged(const QString & title) { QByteArray t3c6de1 = title.toUtf8(); QtQuick_PackedString titlePacked = { const_cast<char*>(t3c6de1.prepend("WHITESPACE").constData()+10), t3c6de1.size()-10 };callbackQQuickWindow_WindowTitleChanged(this, titlePacked); };
	void Signal_XChanged(int arg) { callbackQQuickWindow_XChanged(this, arg); };
	void Signal_YChanged(int arg) { callbackQQuickWindow_YChanged(this, arg); };
	QObject * focusObject() const { return static_cast<QObject*>(callbackQQuickWindow_FocusObject(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize size() const { return *static_cast<QSize*>(callbackQQuickWindow_Size(const_cast<void*>(static_cast<const void*>(this)))); };
	QSurfaceFormat format() const { return *static_cast<QSurfaceFormat*>(callbackQQuickWindow_Format(const_cast<void*>(static_cast<const void*>(this)))); };
	SurfaceType surfaceType() const { return static_cast<QSurface::SurfaceType>(callbackQQuickWindow_SurfaceType(const_cast<void*>(static_cast<const void*>(this)))); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQQuickWindow_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQQuickWindow_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQQuickWindow_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQQuickWindow_CustomEvent(this, event); };
	void deleteLater() { callbackQQuickWindow_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQQuickWindow_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQQuickWindow_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtQuick_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQQuickWindow_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQQuickWindow_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQQuickWindow_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQQuickView*)

int QQuickView_QQuickView_QRegisterMetaType(){qRegisterMetaType<QQuickView*>(); return qRegisterMetaType<MyQQuickView*>();}

void* QQuickView_NewQQuickView2(void* engine, void* parent)
{
	if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQQuickView(static_cast<QQmlEngine*>(engine), static_cast<QPaintDeviceWindow*>(parent));
	} else {
		return new MyQQuickView(static_cast<QQmlEngine*>(engine), static_cast<QWindow*>(parent));
	}
}

void* QQuickView_NewQQuickView(void* parent)
{
	if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQQuickView(static_cast<QPaintDeviceWindow*>(parent));
	} else {
		return new MyQQuickView(static_cast<QWindow*>(parent));
	}
}

void* QQuickView_NewQQuickView3(void* source, void* parent)
{
	if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQQuickView(*static_cast<QUrl*>(source), static_cast<QPaintDeviceWindow*>(parent));
	} else {
		return new MyQQuickView(*static_cast<QUrl*>(source), static_cast<QWindow*>(parent));
	}
}

void QQuickView_SetResizeMode(void* ptr, long long vre)
{
	static_cast<QQuickView*>(ptr)->setResizeMode(static_cast<QQuickView::ResizeMode>(vre));
}

void QQuickView_SetSource(void* ptr, void* url)
{
	QMetaObject::invokeMethod(static_cast<QQuickView*>(ptr), "setSource", Q_ARG(const QUrl, *static_cast<QUrl*>(url)));
}

void QQuickView_SetSourceDefault(void* ptr, void* url)
{
		static_cast<QQuickView*>(ptr)->QQuickView::setSource(*static_cast<QUrl*>(url));
}

void QQuickView_ConnectStatusChanged(void* ptr)
{
	qRegisterMetaType<QQuickView::Status>();
	QObject::connect(static_cast<QQuickView*>(ptr), static_cast<void (QQuickView::*)(QQuickView::Status)>(&QQuickView::statusChanged), static_cast<MyQQuickView*>(ptr), static_cast<void (MyQQuickView::*)(QQuickView::Status)>(&MyQQuickView::Signal_StatusChanged));
}

void QQuickView_DisconnectStatusChanged(void* ptr)
{
	QObject::disconnect(static_cast<QQuickView*>(ptr), static_cast<void (QQuickView::*)(QQuickView::Status)>(&QQuickView::statusChanged), static_cast<MyQQuickView*>(ptr), static_cast<void (MyQQuickView::*)(QQuickView::Status)>(&MyQQuickView::Signal_StatusChanged));
}

void QQuickView_StatusChanged(void* ptr, long long status)
{
	static_cast<QQuickView*>(ptr)->statusChanged(static_cast<QQuickView::Status>(status));
}

void QQuickView_DestroyQQuickView(void* ptr)
{
	static_cast<QQuickView*>(ptr)->~QQuickView();
}

void QQuickView_DestroyQQuickViewDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

struct QtQuick_PackedList QQuickView_Errors(void* ptr)
{
	return ({ QList<QQmlError>* tmpValue = new QList<QQmlError>(static_cast<QQuickView*>(ptr)->errors()); QtQuick_PackedList { tmpValue, tmpValue->size() }; });
}

void* QQuickView_RootContext(void* ptr)
{
	return static_cast<QQuickView*>(ptr)->rootContext();
}

void* QQuickView_Engine(void* ptr)
{
	return static_cast<QQuickView*>(ptr)->engine();
}

void* QQuickView_RootObject(void* ptr)
{
	return static_cast<QQuickView*>(ptr)->rootObject();
}

void* QQuickView_InitialSize(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QQuickView*>(ptr)->initialSize(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QQuickView_Source(void* ptr)
{
	return new QUrl(static_cast<QQuickView*>(ptr)->source());
}

long long QQuickView_ResizeMode(void* ptr)
{
	return static_cast<QQuickView*>(ptr)->resizeMode();
}

long long QQuickView_Status(void* ptr)
{
	return static_cast<QQuickView*>(ptr)->status();
}

void* QQuickView___errors_atList(void* ptr, int i)
{
	return new QQmlError(static_cast<QList<QQmlError>*>(ptr)->at(i));
}

void QQuickView___errors_setList(void* ptr, void* i)
{
	static_cast<QList<QQmlError>*>(ptr)->append(*static_cast<QQmlError*>(i));
}

void* QQuickView___errors_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QQmlError>;
}

class MyQQuickWidget: public QQuickWidget
{
public:
	MyQQuickWidget(QQmlEngine *engine, QWidget *parent) : QQuickWidget(engine, parent) {QQuickWidget_QQuickWidget_QRegisterMetaType();};
	MyQQuickWidget(QWidget *parent = Q_NULLPTR) : QQuickWidget(parent) {QQuickWidget_QQuickWidget_QRegisterMetaType();};
	MyQQuickWidget(const QUrl &source, QWidget *parent = Q_NULLPTR) : QQuickWidget(source, parent) {QQuickWidget_QQuickWidget_QRegisterMetaType();};
	bool event(QEvent * e) { return callbackQQuickWidget_Event(this, e) != 0; };
	void dragEnterEvent(QDragEnterEvent * e) { callbackQQuickWidget_DragEnterEvent(this, e); };
	void dragLeaveEvent(QDragLeaveEvent * e) { callbackQQuickWidget_DragLeaveEvent(this, e); };
	void dragMoveEvent(QDragMoveEvent * e) { callbackQQuickWidget_DragMoveEvent(this, e); };
	void dropEvent(QDropEvent * e) { callbackQQuickWidget_DropEvent(this, e); };
	void focusInEvent(QFocusEvent * event) { callbackQQuickWidget_FocusInEvent(this, event); };
	void focusOutEvent(QFocusEvent * event) { callbackQQuickWidget_FocusOutEvent(this, event); };
	void hideEvent(QHideEvent * vqh) { callbackQQuickWidget_HideEvent(this, vqh); };
	void keyPressEvent(QKeyEvent * e) { callbackQQuickWidget_KeyPressEvent(this, e); };
	void keyReleaseEvent(QKeyEvent * e) { callbackQQuickWidget_KeyReleaseEvent(this, e); };
	void mouseDoubleClickEvent(QMouseEvent * e) { callbackQQuickWidget_MouseDoubleClickEvent(this, e); };
	void mouseMoveEvent(QMouseEvent * e) { callbackQQuickWidget_MouseMoveEvent(this, e); };
	void mousePressEvent(QMouseEvent * e) { callbackQQuickWidget_MousePressEvent(this, e); };
	void mouseReleaseEvent(QMouseEvent * e) { callbackQQuickWidget_MouseReleaseEvent(this, e); };
	void paintEvent(QPaintEvent * event) { callbackQQuickWidget_PaintEvent(this, event); };
	void Signal_SceneGraphError(QQuickWindow::SceneGraphError error, const QString & message) { QByteArray t6f9b9a = message.toUtf8(); QtQuick_PackedString messagePacked = { const_cast<char*>(t6f9b9a.prepend("WHITESPACE").constData()+10), t6f9b9a.size()-10 };callbackQQuickWidget_SceneGraphError(this, error, messagePacked); };
	void setSource(const QUrl & url) { callbackQQuickWidget_SetSource(this, const_cast<QUrl*>(&url)); };
	void showEvent(QShowEvent * vqs) { callbackQQuickWidget_ShowEvent(this, vqs); };
	void Signal_StatusChanged(QQuickWidget::Status status) { callbackQQuickWidget_StatusChanged(this, status); };
	void wheelEvent(QWheelEvent * e) { callbackQQuickWidget_WheelEvent(this, e); };
	 ~MyQQuickWidget() { callbackQQuickWidget_DestroyQQuickWidget(this); };
	bool close() { return callbackQQuickWidget_Close(this) != 0; };
	bool focusNextPrevChild(bool next) { return callbackQQuickWidget_FocusNextPrevChild(this, next) != 0; };
	bool nativeEvent(const QByteArray & eventType, void * message, long * result) { return callbackQQuickWidget_NativeEvent(this, const_cast<QByteArray*>(&eventType), message, *result) != 0; };
	void actionEvent(QActionEvent * event) { callbackQQuickWidget_ActionEvent(this, event); };
	void changeEvent(QEvent * event) { callbackQQuickWidget_ChangeEvent(this, event); };
	void closeEvent(QCloseEvent * event) { callbackQQuickWidget_CloseEvent(this, event); };
	void contextMenuEvent(QContextMenuEvent * event) { callbackQQuickWidget_ContextMenuEvent(this, event); };
	void Signal_CustomContextMenuRequested(const QPoint & pos) { callbackQQuickWidget_CustomContextMenuRequested(this, const_cast<QPoint*>(&pos)); };
	void enterEvent(QEvent * event) { callbackQQuickWidget_EnterEvent(this, event); };
	void hide() { callbackQQuickWidget_Hide(this); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackQQuickWidget_InputMethodEvent(this, event); };
	void leaveEvent(QEvent * event) { callbackQQuickWidget_LeaveEvent(this, event); };
	void lower() { callbackQQuickWidget_Lower(this); };
	void moveEvent(QMoveEvent * event) { callbackQQuickWidget_MoveEvent(this, event); };
	void raise() { callbackQQuickWidget_Raise(this); };
	void repaint() { callbackQQuickWidget_Repaint(this); };
	void resizeEvent(QResizeEvent * event) { callbackQQuickWidget_ResizeEvent(this, event); };
	void setDisabled(bool disable) { callbackQQuickWidget_SetDisabled(this, disable); };
	void setEnabled(bool vbo) { callbackQQuickWidget_SetEnabled(this, vbo); };
	void setFocus() { callbackQQuickWidget_SetFocus2(this); };
	void setHidden(bool hidden) { callbackQQuickWidget_SetHidden(this, hidden); };
	void setStyleSheet(const QString & styleSheet) { QByteArray t728ae7 = styleSheet.toUtf8(); QtQuick_PackedString styleSheetPacked = { const_cast<char*>(t728ae7.prepend("WHITESPACE").constData()+10), t728ae7.size()-10 };callbackQQuickWidget_SetStyleSheet(this, styleSheetPacked); };
	void setVisible(bool visible) { callbackQQuickWidget_SetVisible(this, visible); };
	void setWindowModified(bool vbo) { callbackQQuickWidget_SetWindowModified(this, vbo); };
	void setWindowTitle(const QString & vqs) { QByteArray tda39a3 = vqs.toUtf8(); QtQuick_PackedString vqsPacked = { const_cast<char*>(tda39a3.prepend("WHITESPACE").constData()+10), tda39a3.size()-10 };callbackQQuickWidget_SetWindowTitle(this, vqsPacked); };
	void show() { callbackQQuickWidget_Show(this); };
	void showFullScreen() { callbackQQuickWidget_ShowFullScreen(this); };
	void showMaximized() { callbackQQuickWidget_ShowMaximized(this); };
	void showMinimized() { callbackQQuickWidget_ShowMinimized(this); };
	void showNormal() { callbackQQuickWidget_ShowNormal(this); };
	void tabletEvent(QTabletEvent * event) { callbackQQuickWidget_TabletEvent(this, event); };
	void update() { callbackQQuickWidget_Update(this); };
	void updateMicroFocus() { callbackQQuickWidget_UpdateMicroFocus(this); };
	void Signal_WindowIconChanged(const QIcon & icon) { callbackQQuickWidget_WindowIconChanged(this, const_cast<QIcon*>(&icon)); };
	void Signal_WindowTitleChanged(const QString & title) { QByteArray t3c6de1 = title.toUtf8(); QtQuick_PackedString titlePacked = { const_cast<char*>(t3c6de1.prepend("WHITESPACE").constData()+10), t3c6de1.size()-10 };callbackQQuickWidget_WindowTitleChanged(this, titlePacked); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQQuickWidget_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize minimumSizeHint() const { return *static_cast<QSize*>(callbackQQuickWidget_MinimumSizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQQuickWidget_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackQQuickWidget_InputMethodQuery(const_cast<void*>(static_cast<const void*>(this)), query)); };
	bool hasHeightForWidth() const { return callbackQQuickWidget_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int heightForWidth(int w) const { return callbackQQuickWidget_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackQQuickWidget_Metric(const_cast<void*>(static_cast<const void*>(this)), m); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQQuickWidget_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQQuickWidget_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQQuickWidget_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQQuickWidget_CustomEvent(this, event); };
	void deleteLater() { callbackQQuickWidget_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQQuickWidget_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQQuickWidget_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtQuick_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQQuickWidget_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQQuickWidget_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQQuickWidget_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQQuickWidget*)

int QQuickWidget_QQuickWidget_QRegisterMetaType(){qRegisterMetaType<QQuickWidget*>(); return qRegisterMetaType<MyQQuickWidget*>();}

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

char QQuickWidget_EventDefault(void* ptr, void* e)
{
		return static_cast<QQuickWidget*>(ptr)->QQuickWidget::event(static_cast<QEvent*>(e));
}

void QQuickWidget_DragEnterEventDefault(void* ptr, void* e)
{
		static_cast<QQuickWidget*>(ptr)->QQuickWidget::dragEnterEvent(static_cast<QDragEnterEvent*>(e));
}

void QQuickWidget_DragLeaveEventDefault(void* ptr, void* e)
{
		static_cast<QQuickWidget*>(ptr)->QQuickWidget::dragLeaveEvent(static_cast<QDragLeaveEvent*>(e));
}

void QQuickWidget_DragMoveEventDefault(void* ptr, void* e)
{
		static_cast<QQuickWidget*>(ptr)->QQuickWidget::dragMoveEvent(static_cast<QDragMoveEvent*>(e));
}

void QQuickWidget_DropEventDefault(void* ptr, void* e)
{
		static_cast<QQuickWidget*>(ptr)->QQuickWidget::dropEvent(static_cast<QDropEvent*>(e));
}

void QQuickWidget_FocusInEventDefault(void* ptr, void* event)
{
		static_cast<QQuickWidget*>(ptr)->QQuickWidget::focusInEvent(static_cast<QFocusEvent*>(event));
}

void QQuickWidget_FocusOutEventDefault(void* ptr, void* event)
{
		static_cast<QQuickWidget*>(ptr)->QQuickWidget::focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QQuickWidget_HideEventDefault(void* ptr, void* vqh)
{
		static_cast<QQuickWidget*>(ptr)->QQuickWidget::hideEvent(static_cast<QHideEvent*>(vqh));
}

void QQuickWidget_KeyPressEventDefault(void* ptr, void* e)
{
		static_cast<QQuickWidget*>(ptr)->QQuickWidget::keyPressEvent(static_cast<QKeyEvent*>(e));
}

void QQuickWidget_KeyReleaseEventDefault(void* ptr, void* e)
{
		static_cast<QQuickWidget*>(ptr)->QQuickWidget::keyReleaseEvent(static_cast<QKeyEvent*>(e));
}

void QQuickWidget_MouseDoubleClickEventDefault(void* ptr, void* e)
{
		static_cast<QQuickWidget*>(ptr)->QQuickWidget::mouseDoubleClickEvent(static_cast<QMouseEvent*>(e));
}

void QQuickWidget_MouseMoveEventDefault(void* ptr, void* e)
{
		static_cast<QQuickWidget*>(ptr)->QQuickWidget::mouseMoveEvent(static_cast<QMouseEvent*>(e));
}

void QQuickWidget_MousePressEventDefault(void* ptr, void* e)
{
		static_cast<QQuickWidget*>(ptr)->QQuickWidget::mousePressEvent(static_cast<QMouseEvent*>(e));
}

void QQuickWidget_MouseReleaseEventDefault(void* ptr, void* e)
{
		static_cast<QQuickWidget*>(ptr)->QQuickWidget::mouseReleaseEvent(static_cast<QMouseEvent*>(e));
}

void QQuickWidget_PaintEventDefault(void* ptr, void* event)
{
		static_cast<QQuickWidget*>(ptr)->QQuickWidget::paintEvent(static_cast<QPaintEvent*>(event));
}

void QQuickWidget_ConnectSceneGraphError(void* ptr)
{
	qRegisterMetaType<QQuickWindow::SceneGraphError>();
	QObject::connect(static_cast<QQuickWidget*>(ptr), static_cast<void (QQuickWidget::*)(QQuickWindow::SceneGraphError, const QString &)>(&QQuickWidget::sceneGraphError), static_cast<MyQQuickWidget*>(ptr), static_cast<void (MyQQuickWidget::*)(QQuickWindow::SceneGraphError, const QString &)>(&MyQQuickWidget::Signal_SceneGraphError));
}

void QQuickWidget_DisconnectSceneGraphError(void* ptr)
{
	QObject::disconnect(static_cast<QQuickWidget*>(ptr), static_cast<void (QQuickWidget::*)(QQuickWindow::SceneGraphError, const QString &)>(&QQuickWidget::sceneGraphError), static_cast<MyQQuickWidget*>(ptr), static_cast<void (MyQQuickWidget::*)(QQuickWindow::SceneGraphError, const QString &)>(&MyQQuickWidget::Signal_SceneGraphError));
}

void QQuickWidget_SceneGraphError(void* ptr, long long error, struct QtQuick_PackedString message)
{
	static_cast<QQuickWidget*>(ptr)->sceneGraphError(static_cast<QQuickWindow::SceneGraphError>(error), QString::fromUtf8(message.data, message.len));
}

void QQuickWidget_SetClearColor(void* ptr, void* color)
{
	static_cast<QQuickWidget*>(ptr)->setClearColor(*static_cast<QColor*>(color));
}

void QQuickWidget_SetFormat(void* ptr, void* format)
{
	static_cast<QQuickWidget*>(ptr)->setFormat(*static_cast<QSurfaceFormat*>(format));
}

void QQuickWidget_SetResizeMode(void* ptr, long long vre)
{
	static_cast<QQuickWidget*>(ptr)->setResizeMode(static_cast<QQuickWidget::ResizeMode>(vre));
}

void QQuickWidget_SetSource(void* ptr, void* url)
{
	QMetaObject::invokeMethod(static_cast<QQuickWidget*>(ptr), "setSource", Q_ARG(const QUrl, *static_cast<QUrl*>(url)));
}

void QQuickWidget_SetSourceDefault(void* ptr, void* url)
{
		static_cast<QQuickWidget*>(ptr)->QQuickWidget::setSource(*static_cast<QUrl*>(url));
}

void QQuickWidget_ShowEventDefault(void* ptr, void* vqs)
{
		static_cast<QQuickWidget*>(ptr)->QQuickWidget::showEvent(static_cast<QShowEvent*>(vqs));
}

void QQuickWidget_ConnectStatusChanged(void* ptr)
{
	qRegisterMetaType<QQuickWidget::Status>();
	QObject::connect(static_cast<QQuickWidget*>(ptr), static_cast<void (QQuickWidget::*)(QQuickWidget::Status)>(&QQuickWidget::statusChanged), static_cast<MyQQuickWidget*>(ptr), static_cast<void (MyQQuickWidget::*)(QQuickWidget::Status)>(&MyQQuickWidget::Signal_StatusChanged));
}

void QQuickWidget_DisconnectStatusChanged(void* ptr)
{
	QObject::disconnect(static_cast<QQuickWidget*>(ptr), static_cast<void (QQuickWidget::*)(QQuickWidget::Status)>(&QQuickWidget::statusChanged), static_cast<MyQQuickWidget*>(ptr), static_cast<void (MyQQuickWidget::*)(QQuickWidget::Status)>(&MyQQuickWidget::Signal_StatusChanged));
}

void QQuickWidget_StatusChanged(void* ptr, long long status)
{
	static_cast<QQuickWidget*>(ptr)->statusChanged(static_cast<QQuickWidget::Status>(status));
}

void QQuickWidget_WheelEventDefault(void* ptr, void* e)
{
		static_cast<QQuickWidget*>(ptr)->QQuickWidget::wheelEvent(static_cast<QWheelEvent*>(e));
}

void QQuickWidget_DestroyQQuickWidget(void* ptr)
{
	static_cast<QQuickWidget*>(ptr)->~QQuickWidget();
}

void QQuickWidget_DestroyQQuickWidgetDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QQuickWidget_GrabFramebuffer(void* ptr)
{
	return new QImage(static_cast<QQuickWidget*>(ptr)->grabFramebuffer());
}

struct QtQuick_PackedList QQuickWidget_Errors(void* ptr)
{
	return ({ QList<QQmlError>* tmpValue = new QList<QQmlError>(static_cast<QQuickWidget*>(ptr)->errors()); QtQuick_PackedList { tmpValue, tmpValue->size() }; });
}

void* QQuickWidget_RootContext(void* ptr)
{
	return static_cast<QQuickWidget*>(ptr)->rootContext();
}

void* QQuickWidget_Engine(void* ptr)
{
	return static_cast<QQuickWidget*>(ptr)->engine();
}

void* QQuickWidget_RootObject(void* ptr)
{
	return static_cast<QQuickWidget*>(ptr)->rootObject();
}

void* QQuickWidget_QuickWindow(void* ptr)
{
	return static_cast<QQuickWidget*>(ptr)->quickWindow();
}

void* QQuickWidget_InitialSize(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QQuickWidget*>(ptr)->initialSize(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QQuickWidget_Format(void* ptr)
{
	return new QSurfaceFormat(static_cast<QQuickWidget*>(ptr)->format());
}

void* QQuickWidget_Source(void* ptr)
{
	return new QUrl(static_cast<QQuickWidget*>(ptr)->source());
}

long long QQuickWidget_ResizeMode(void* ptr)
{
	return static_cast<QQuickWidget*>(ptr)->resizeMode();
}

long long QQuickWidget_Status(void* ptr)
{
	return static_cast<QQuickWidget*>(ptr)->status();
}

void* QQuickWidget___errors_atList(void* ptr, int i)
{
	return new QQmlError(static_cast<QList<QQmlError>*>(ptr)->at(i));
}

void QQuickWidget___errors_setList(void* ptr, void* i)
{
	static_cast<QList<QQmlError>*>(ptr)->append(*static_cast<QQmlError*>(i));
}

void* QQuickWidget___errors_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QQmlError>;
}

void* QQuickWidget___addActions_actions_atList(void* ptr, int i)
{
	return const_cast<QAction*>(static_cast<QList<QAction *>*>(ptr)->at(i));
}

void QQuickWidget___addActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QQuickWidget___addActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>;
}

void* QQuickWidget___insertActions_actions_atList(void* ptr, int i)
{
	return const_cast<QAction*>(static_cast<QList<QAction *>*>(ptr)->at(i));
}

void QQuickWidget___insertActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QQuickWidget___insertActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>;
}

void* QQuickWidget___actions_atList(void* ptr, int i)
{
	return const_cast<QAction*>(static_cast<QList<QAction *>*>(ptr)->at(i));
}

void QQuickWidget___actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QQuickWidget___actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>;
}

void* QQuickWidget___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QQuickWidget___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QQuickWidget___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QQuickWidget___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QQuickWidget___findChildren_setList2(void* ptr, void* i)
{
	if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QQuickItem*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QQuickWidget___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QQuickWidget___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QQuickWidget___findChildren_setList3(void* ptr, void* i)
{
	if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QQuickItem*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QQuickWidget___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QQuickWidget___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QQuickWidget___findChildren_setList(void* ptr, void* i)
{
	if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QQuickItem*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QQuickWidget___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QQuickWidget___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QQuickWidget___children_setList(void* ptr, void* i)
{
	if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QQuickItem*>(i));
	} else {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QQuickWidget___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

char QQuickWidget_CloseDefault(void* ptr)
{
		return static_cast<QQuickWidget*>(ptr)->QQuickWidget::close();
}

char QQuickWidget_FocusNextPrevChildDefault(void* ptr, char next)
{
		return static_cast<QQuickWidget*>(ptr)->QQuickWidget::focusNextPrevChild(next != 0);
}

char QQuickWidget_NativeEventDefault(void* ptr, void* eventType, void* message, long result)
{
		return static_cast<QQuickWidget*>(ptr)->QQuickWidget::nativeEvent(*static_cast<QByteArray*>(eventType), message, &result);
}

void QQuickWidget_ActionEventDefault(void* ptr, void* event)
{
		static_cast<QQuickWidget*>(ptr)->QQuickWidget::actionEvent(static_cast<QActionEvent*>(event));
}

void QQuickWidget_ChangeEventDefault(void* ptr, void* event)
{
		static_cast<QQuickWidget*>(ptr)->QQuickWidget::changeEvent(static_cast<QEvent*>(event));
}

void QQuickWidget_CloseEventDefault(void* ptr, void* event)
{
		static_cast<QQuickWidget*>(ptr)->QQuickWidget::closeEvent(static_cast<QCloseEvent*>(event));
}

void QQuickWidget_ContextMenuEventDefault(void* ptr, void* event)
{
		static_cast<QQuickWidget*>(ptr)->QQuickWidget::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
}

void QQuickWidget_EnterEventDefault(void* ptr, void* event)
{
		static_cast<QQuickWidget*>(ptr)->QQuickWidget::enterEvent(static_cast<QEvent*>(event));
}

void QQuickWidget_HideDefault(void* ptr)
{
		static_cast<QQuickWidget*>(ptr)->QQuickWidget::hide();
}

void QQuickWidget_InputMethodEventDefault(void* ptr, void* event)
{
		static_cast<QQuickWidget*>(ptr)->QQuickWidget::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void QQuickWidget_LeaveEventDefault(void* ptr, void* event)
{
		static_cast<QQuickWidget*>(ptr)->QQuickWidget::leaveEvent(static_cast<QEvent*>(event));
}

void QQuickWidget_LowerDefault(void* ptr)
{
		static_cast<QQuickWidget*>(ptr)->QQuickWidget::lower();
}

void QQuickWidget_MoveEventDefault(void* ptr, void* event)
{
		static_cast<QQuickWidget*>(ptr)->QQuickWidget::moveEvent(static_cast<QMoveEvent*>(event));
}

void QQuickWidget_RaiseDefault(void* ptr)
{
		static_cast<QQuickWidget*>(ptr)->QQuickWidget::raise();
}

void QQuickWidget_RepaintDefault(void* ptr)
{
		static_cast<QQuickWidget*>(ptr)->QQuickWidget::repaint();
}

void QQuickWidget_ResizeEventDefault(void* ptr, void* event)
{
		static_cast<QQuickWidget*>(ptr)->QQuickWidget::resizeEvent(static_cast<QResizeEvent*>(event));
}

void QQuickWidget_SetDisabledDefault(void* ptr, char disable)
{
		static_cast<QQuickWidget*>(ptr)->QQuickWidget::setDisabled(disable != 0);
}

void QQuickWidget_SetEnabledDefault(void* ptr, char vbo)
{
		static_cast<QQuickWidget*>(ptr)->QQuickWidget::setEnabled(vbo != 0);
}

void QQuickWidget_SetFocus2Default(void* ptr)
{
		static_cast<QQuickWidget*>(ptr)->QQuickWidget::setFocus();
}

void QQuickWidget_SetHiddenDefault(void* ptr, char hidden)
{
		static_cast<QQuickWidget*>(ptr)->QQuickWidget::setHidden(hidden != 0);
}

void QQuickWidget_SetStyleSheetDefault(void* ptr, struct QtQuick_PackedString styleSheet)
{
		static_cast<QQuickWidget*>(ptr)->QQuickWidget::setStyleSheet(QString::fromUtf8(styleSheet.data, styleSheet.len));
}

void QQuickWidget_SetVisibleDefault(void* ptr, char visible)
{
		static_cast<QQuickWidget*>(ptr)->QQuickWidget::setVisible(visible != 0);
}

void QQuickWidget_SetWindowModifiedDefault(void* ptr, char vbo)
{
		static_cast<QQuickWidget*>(ptr)->QQuickWidget::setWindowModified(vbo != 0);
}

void QQuickWidget_SetWindowTitleDefault(void* ptr, struct QtQuick_PackedString vqs)
{
		static_cast<QQuickWidget*>(ptr)->QQuickWidget::setWindowTitle(QString::fromUtf8(vqs.data, vqs.len));
}

void QQuickWidget_ShowDefault(void* ptr)
{
		static_cast<QQuickWidget*>(ptr)->QQuickWidget::show();
}

void QQuickWidget_ShowFullScreenDefault(void* ptr)
{
		static_cast<QQuickWidget*>(ptr)->QQuickWidget::showFullScreen();
}

void QQuickWidget_ShowMaximizedDefault(void* ptr)
{
		static_cast<QQuickWidget*>(ptr)->QQuickWidget::showMaximized();
}

void QQuickWidget_ShowMinimizedDefault(void* ptr)
{
		static_cast<QQuickWidget*>(ptr)->QQuickWidget::showMinimized();
}

void QQuickWidget_ShowNormalDefault(void* ptr)
{
		static_cast<QQuickWidget*>(ptr)->QQuickWidget::showNormal();
}

void QQuickWidget_TabletEventDefault(void* ptr, void* event)
{
		static_cast<QQuickWidget*>(ptr)->QQuickWidget::tabletEvent(static_cast<QTabletEvent*>(event));
}

void QQuickWidget_UpdateDefault(void* ptr)
{
		static_cast<QQuickWidget*>(ptr)->QQuickWidget::update();
}

void QQuickWidget_UpdateMicroFocusDefault(void* ptr)
{
		static_cast<QQuickWidget*>(ptr)->QQuickWidget::updateMicroFocus();
}

void* QQuickWidget_PaintEngineDefault(void* ptr)
{
		return static_cast<QQuickWidget*>(ptr)->QQuickWidget::paintEngine();
}

void* QQuickWidget_MinimumSizeHintDefault(void* ptr)
{
		return ({ QSize tmpValue = static_cast<QQuickWidget*>(ptr)->QQuickWidget::minimumSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QQuickWidget_SizeHintDefault(void* ptr)
{
		return ({ QSize tmpValue = static_cast<QQuickWidget*>(ptr)->QQuickWidget::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QQuickWidget_InputMethodQueryDefault(void* ptr, long long query)
{
		return new QVariant(static_cast<QQuickWidget*>(ptr)->QQuickWidget::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

char QQuickWidget_HasHeightForWidthDefault(void* ptr)
{
		return static_cast<QQuickWidget*>(ptr)->QQuickWidget::hasHeightForWidth();
}

int QQuickWidget_HeightForWidthDefault(void* ptr, int w)
{
		return static_cast<QQuickWidget*>(ptr)->QQuickWidget::heightForWidth(w);
}

int QQuickWidget_MetricDefault(void* ptr, long long m)
{
		return static_cast<QQuickWidget*>(ptr)->QQuickWidget::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
}

char QQuickWidget_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(watched))) {
		return static_cast<QQuickWidget*>(ptr)->QQuickWidget::eventFilter(static_cast<QQuickItem*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QQuickWidget*>(ptr)->QQuickWidget::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

void QQuickWidget_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QQuickWidget*>(ptr)->QQuickWidget::childEvent(static_cast<QChildEvent*>(event));
}

void QQuickWidget_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QQuickWidget*>(ptr)->QQuickWidget::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQuickWidget_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QQuickWidget*>(ptr)->QQuickWidget::customEvent(static_cast<QEvent*>(event));
}

void QQuickWidget_DeleteLaterDefault(void* ptr)
{
		static_cast<QQuickWidget*>(ptr)->QQuickWidget::deleteLater();
}

void QQuickWidget_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QQuickWidget*>(ptr)->QQuickWidget::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQuickWidget_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QQuickWidget*>(ptr)->QQuickWidget::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QQuickWidget_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QQuickWidget*>(ptr)->QQuickWidget::metaObject());
}

class MyQQuickWindow: public QQuickWindow
{
public:
	MyQQuickWindow(QWindow *parent = Q_NULLPTR) : QQuickWindow(parent) {QQuickWindow_QQuickWindow_QRegisterMetaType();};
	bool event(QEvent * e) { return callbackQQuickWindow_Event(this, e) != 0; };
	void Signal_ActiveFocusItemChanged() { callbackQQuickWindow_ActiveFocusItemChanged(this); };
	void Signal_AfterAnimating() { callbackQQuickWindow_AfterAnimating(this); };
	void Signal_AfterRendering() { callbackQQuickWindow_AfterRendering(this); };
	void Signal_AfterSynchronizing() { callbackQQuickWindow_AfterSynchronizing(this); };
	void Signal_BeforeRendering() { callbackQQuickWindow_BeforeRendering(this); };
	void Signal_BeforeSynchronizing() { callbackQQuickWindow_BeforeSynchronizing(this); };
	void Signal_ColorChanged(const QColor & vqc) { callbackQQuickWindow_ColorChanged(this, const_cast<QColor*>(&vqc)); };
	void exposeEvent(QExposeEvent * vqe) { callbackQQuickWindow_ExposeEvent(this, vqe); };
	void focusInEvent(QFocusEvent * ev) { callbackQQuickWindow_FocusInEvent(this, ev); };
	void focusOutEvent(QFocusEvent * ev) { callbackQQuickWindow_FocusOutEvent(this, ev); };
	void Signal_FrameSwapped() { callbackQQuickWindow_FrameSwapped(this); };
	void hideEvent(QHideEvent * vqh) { callbackQQuickWindow_HideEvent(this, vqh); };
	void keyPressEvent(QKeyEvent * e) { callbackQQuickWindow_KeyPressEvent(this, e); };
	void keyReleaseEvent(QKeyEvent * e) { callbackQQuickWindow_KeyReleaseEvent(this, e); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQQuickWindow_MouseDoubleClickEvent(this, event); };
	void mouseMoveEvent(QMouseEvent * event) { callbackQQuickWindow_MouseMoveEvent(this, event); };
	void mousePressEvent(QMouseEvent * event) { callbackQQuickWindow_MousePressEvent(this, event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackQQuickWindow_MouseReleaseEvent(this, event); };
	void Signal_OpenglContextCreated(QOpenGLContext * context) { callbackQQuickWindow_OpenglContextCreated(this, context); };
	void releaseResources() { callbackQQuickWindow_ReleaseResources(this); };
	void resizeEvent(QResizeEvent * ev) { callbackQQuickWindow_ResizeEvent(this, ev); };
	void Signal_SceneGraphAboutToStop() { callbackQQuickWindow_SceneGraphAboutToStop(this); };
	void Signal_SceneGraphError(QQuickWindow::SceneGraphError error, const QString & message) { QByteArray t6f9b9a = message.toUtf8(); QtQuick_PackedString messagePacked = { const_cast<char*>(t6f9b9a.prepend("WHITESPACE").constData()+10), t6f9b9a.size()-10 };callbackQQuickWindow_SceneGraphError(this, error, messagePacked); };
	void Signal_SceneGraphInitialized() { callbackQQuickWindow_SceneGraphInitialized(this); };
	void Signal_SceneGraphInvalidated() { callbackQQuickWindow_SceneGraphInvalidated(this); };
	void showEvent(QShowEvent * vqs) { callbackQQuickWindow_ShowEvent(this, vqs); };
	void update() { callbackQQuickWindow_Update(this); };
	void wheelEvent(QWheelEvent * event) { callbackQQuickWindow_WheelEvent(this, event); };
	 ~MyQQuickWindow() { callbackQQuickWindow_DestroyQQuickWindow(this); };
	bool close() { return callbackQQuickWindow_Close(this) != 0; };
	bool nativeEvent(const QByteArray & eventType, void * message, long * result) { return callbackQQuickWindow_NativeEvent(this, const_cast<QByteArray*>(&eventType), message, *result) != 0; };
	void Signal_ActiveChanged() { callbackQQuickWindow_ActiveChanged(this); };
	void alert(int msec) { callbackQQuickWindow_Alert(this, msec); };
	void Signal_ContentOrientationChanged(Qt::ScreenOrientation orientation) { callbackQQuickWindow_ContentOrientationChanged(this, orientation); };
	void Signal_FocusObjectChanged(QObject * object) { callbackQQuickWindow_FocusObjectChanged(this, object); };
	void Signal_HeightChanged(int arg) { callbackQQuickWindow_HeightChanged(this, arg); };
	void hide() { callbackQQuickWindow_Hide(this); };
	void lower() { callbackQQuickWindow_Lower(this); };
	void Signal_MaximumHeightChanged(int arg) { callbackQQuickWindow_MaximumHeightChanged(this, arg); };
	void Signal_MaximumWidthChanged(int arg) { callbackQQuickWindow_MaximumWidthChanged(this, arg); };
	void Signal_MinimumHeightChanged(int arg) { callbackQQuickWindow_MinimumHeightChanged(this, arg); };
	void Signal_MinimumWidthChanged(int arg) { callbackQQuickWindow_MinimumWidthChanged(this, arg); };
	void Signal_ModalityChanged(Qt::WindowModality modality) { callbackQQuickWindow_ModalityChanged(this, modality); };
	void moveEvent(QMoveEvent * ev) { callbackQQuickWindow_MoveEvent(this, ev); };
	void Signal_OpacityChanged(qreal opacity) { callbackQQuickWindow_OpacityChanged(this, opacity); };
	void raise() { callbackQQuickWindow_Raise(this); };
	void requestActivate() { callbackQQuickWindow_RequestActivate(this); };
	void requestUpdate() { callbackQQuickWindow_RequestUpdate(this); };
	void Signal_ScreenChanged(QScreen * screen) { callbackQQuickWindow_ScreenChanged(this, screen); };
	void setHeight(int arg) { callbackQQuickWindow_SetHeight(this, arg); };
	void setMaximumHeight(int h) { callbackQQuickWindow_SetMaximumHeight(this, h); };
	void setMaximumWidth(int w) { callbackQQuickWindow_SetMaximumWidth(this, w); };
	void setMinimumHeight(int h) { callbackQQuickWindow_SetMinimumHeight(this, h); };
	void setMinimumWidth(int w) { callbackQQuickWindow_SetMinimumWidth(this, w); };
	void setTitle(const QString & vqs) { QByteArray tda39a3 = vqs.toUtf8(); QtQuick_PackedString vqsPacked = { const_cast<char*>(tda39a3.prepend("WHITESPACE").constData()+10), tda39a3.size()-10 };callbackQQuickWindow_SetTitle(this, vqsPacked); };
	void setVisible(bool visible) { callbackQQuickWindow_SetVisible(this, visible); };
	void setWidth(int arg) { callbackQQuickWindow_SetWidth(this, arg); };
	void setX(int arg) { callbackQQuickWindow_SetX(this, arg); };
	void setY(int arg) { callbackQQuickWindow_SetY(this, arg); };
	void show() { callbackQQuickWindow_Show(this); };
	void showFullScreen() { callbackQQuickWindow_ShowFullScreen(this); };
	void showMaximized() { callbackQQuickWindow_ShowMaximized(this); };
	void showMinimized() { callbackQQuickWindow_ShowMinimized(this); };
	void showNormal() { callbackQQuickWindow_ShowNormal(this); };
	void tabletEvent(QTabletEvent * ev) { callbackQQuickWindow_TabletEvent(this, ev); };
	void touchEvent(QTouchEvent * ev) { callbackQQuickWindow_TouchEvent(this, ev); };
	void Signal_VisibilityChanged(QWindow::Visibility visibility) { callbackQQuickWindow_VisibilityChanged(this, visibility); };
	void Signal_VisibleChanged(bool arg) { callbackQQuickWindow_VisibleChanged(this, arg); };
	void Signal_WidthChanged(int arg) { callbackQQuickWindow_WidthChanged(this, arg); };
	void Signal_WindowStateChanged(Qt::WindowState windowState) { callbackQQuickWindow_WindowStateChanged(this, windowState); };
	void Signal_WindowTitleChanged(const QString & title) { QByteArray t3c6de1 = title.toUtf8(); QtQuick_PackedString titlePacked = { const_cast<char*>(t3c6de1.prepend("WHITESPACE").constData()+10), t3c6de1.size()-10 };callbackQQuickWindow_WindowTitleChanged(this, titlePacked); };
	void Signal_XChanged(int arg) { callbackQQuickWindow_XChanged(this, arg); };
	void Signal_YChanged(int arg) { callbackQQuickWindow_YChanged(this, arg); };
	QObject * focusObject() const { return static_cast<QObject*>(callbackQQuickWindow_FocusObject(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize size() const { return *static_cast<QSize*>(callbackQQuickWindow_Size(const_cast<void*>(static_cast<const void*>(this)))); };
	QSurfaceFormat format() const { return *static_cast<QSurfaceFormat*>(callbackQQuickWindow_Format(const_cast<void*>(static_cast<const void*>(this)))); };
	SurfaceType surfaceType() const { return static_cast<QSurface::SurfaceType>(callbackQQuickWindow_SurfaceType(const_cast<void*>(static_cast<const void*>(this)))); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQQuickWindow_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQQuickWindow_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQQuickWindow_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQQuickWindow_CustomEvent(this, event); };
	void deleteLater() { callbackQQuickWindow_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQQuickWindow_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQQuickWindow_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtQuick_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQQuickWindow_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQQuickWindow_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQQuickWindow_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQQuickWindow*)

int QQuickWindow_QQuickWindow_QRegisterMetaType(){qRegisterMetaType<QQuickWindow*>(); return qRegisterMetaType<MyQQuickWindow*>();}

void* QQuickWindow_GrabWindow(void* ptr)
{
	return new QImage(static_cast<QQuickWindow*>(ptr)->grabWindow());
}

void* QQuickWindow_NewQQuickWindow(void* parent)
{
	if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQQuickWindow(static_cast<QPaintDeviceWindow*>(parent));
	} else {
		return new MyQQuickWindow(static_cast<QWindow*>(parent));
	}
}

struct QtQuick_PackedString QQuickWindow_QQuickWindow_SceneGraphBackend()
{
	return ({ QByteArray t3cc258 = QQuickWindow::sceneGraphBackend().toUtf8(); QtQuick_PackedString { const_cast<char*>(t3cc258.prepend("WHITESPACE").constData()+10), t3cc258.size()-10 }; });
}

char QQuickWindow_EventDefault(void* ptr, void* e)
{
	if (dynamic_cast<QQuickView*>(static_cast<QObject*>(ptr))) {
		return static_cast<QQuickView*>(ptr)->QQuickView::event(static_cast<QEvent*>(e));
	} else {
		return static_cast<QQuickWindow*>(ptr)->QQuickWindow::event(static_cast<QEvent*>(e));
	}
}

char QQuickWindow_QQuickWindow_HasDefaultAlphaBuffer()
{
	return QQuickWindow::hasDefaultAlphaBuffer();
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

void QQuickWindow_ExposeEventDefault(void* ptr, void* vqe)
{
	if (dynamic_cast<QQuickView*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickView*>(ptr)->QQuickView::exposeEvent(static_cast<QExposeEvent*>(vqe));
	} else {
		static_cast<QQuickWindow*>(ptr)->QQuickWindow::exposeEvent(static_cast<QExposeEvent*>(vqe));
	}
}

void QQuickWindow_FocusInEventDefault(void* ptr, void* ev)
{
	if (dynamic_cast<QQuickView*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickView*>(ptr)->QQuickView::focusInEvent(static_cast<QFocusEvent*>(ev));
	} else {
		static_cast<QQuickWindow*>(ptr)->QQuickWindow::focusInEvent(static_cast<QFocusEvent*>(ev));
	}
}

void QQuickWindow_FocusOutEventDefault(void* ptr, void* ev)
{
	if (dynamic_cast<QQuickView*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickView*>(ptr)->QQuickView::focusOutEvent(static_cast<QFocusEvent*>(ev));
	} else {
		static_cast<QQuickWindow*>(ptr)->QQuickWindow::focusOutEvent(static_cast<QFocusEvent*>(ev));
	}
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

void QQuickWindow_HideEventDefault(void* ptr, void* vqh)
{
	if (dynamic_cast<QQuickView*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickView*>(ptr)->QQuickView::hideEvent(static_cast<QHideEvent*>(vqh));
	} else {
		static_cast<QQuickWindow*>(ptr)->QQuickWindow::hideEvent(static_cast<QHideEvent*>(vqh));
	}
}

void QQuickWindow_KeyPressEventDefault(void* ptr, void* e)
{
	if (dynamic_cast<QQuickView*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickView*>(ptr)->QQuickView::keyPressEvent(static_cast<QKeyEvent*>(e));
	} else {
		static_cast<QQuickWindow*>(ptr)->QQuickWindow::keyPressEvent(static_cast<QKeyEvent*>(e));
	}
}

void QQuickWindow_KeyReleaseEventDefault(void* ptr, void* e)
{
	if (dynamic_cast<QQuickView*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickView*>(ptr)->QQuickView::keyReleaseEvent(static_cast<QKeyEvent*>(e));
	} else {
		static_cast<QQuickWindow*>(ptr)->QQuickWindow::keyReleaseEvent(static_cast<QKeyEvent*>(e));
	}
}

void QQuickWindow_MouseDoubleClickEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QQuickView*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickView*>(ptr)->QQuickView::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
	} else {
		static_cast<QQuickWindow*>(ptr)->QQuickWindow::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
	}
}

void QQuickWindow_MouseMoveEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QQuickView*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickView*>(ptr)->QQuickView::mouseMoveEvent(static_cast<QMouseEvent*>(event));
	} else {
		static_cast<QQuickWindow*>(ptr)->QQuickWindow::mouseMoveEvent(static_cast<QMouseEvent*>(event));
	}
}

void QQuickWindow_MousePressEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QQuickView*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickView*>(ptr)->QQuickView::mousePressEvent(static_cast<QMouseEvent*>(event));
	} else {
		static_cast<QQuickWindow*>(ptr)->QQuickWindow::mousePressEvent(static_cast<QMouseEvent*>(event));
	}
}

void QQuickWindow_MouseReleaseEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QQuickView*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickView*>(ptr)->QQuickView::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
	} else {
		static_cast<QQuickWindow*>(ptr)->QQuickWindow::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
	}
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

void QQuickWindow_ReleaseResourcesDefault(void* ptr)
{
	if (dynamic_cast<QQuickView*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickView*>(ptr)->QQuickView::releaseResources();
	} else {
		static_cast<QQuickWindow*>(ptr)->QQuickWindow::releaseResources();
	}
}

void QQuickWindow_ResetOpenGLState(void* ptr)
{
	static_cast<QQuickWindow*>(ptr)->resetOpenGLState();
}

void QQuickWindow_ResizeEventDefault(void* ptr, void* ev)
{
	if (dynamic_cast<QQuickView*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickView*>(ptr)->QQuickView::resizeEvent(static_cast<QResizeEvent*>(ev));
	} else {
		static_cast<QQuickWindow*>(ptr)->QQuickWindow::resizeEvent(static_cast<QResizeEvent*>(ev));
	}
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
	qRegisterMetaType<QQuickWindow::SceneGraphError>();
	QObject::connect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)(QQuickWindow::SceneGraphError, const QString &)>(&QQuickWindow::sceneGraphError), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)(QQuickWindow::SceneGraphError, const QString &)>(&MyQQuickWindow::Signal_SceneGraphError));
}

void QQuickWindow_DisconnectSceneGraphError(void* ptr)
{
	QObject::disconnect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)(QQuickWindow::SceneGraphError, const QString &)>(&QQuickWindow::sceneGraphError), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)(QQuickWindow::SceneGraphError, const QString &)>(&MyQQuickWindow::Signal_SceneGraphError));
}

void QQuickWindow_SceneGraphError(void* ptr, long long error, struct QtQuick_PackedString message)
{
	static_cast<QQuickWindow*>(ptr)->sceneGraphError(static_cast<QQuickWindow::SceneGraphError>(error), QString::fromUtf8(message.data, message.len));
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

void QQuickWindow_ScheduleRenderJob(void* ptr, void* job, long long stage)
{
	static_cast<QQuickWindow*>(ptr)->scheduleRenderJob(static_cast<QRunnable*>(job), static_cast<QQuickWindow::RenderStage>(stage));
}

void QQuickWindow_SetClearBeforeRendering(void* ptr, char enabled)
{
	static_cast<QQuickWindow*>(ptr)->setClearBeforeRendering(enabled != 0);
}

void QQuickWindow_SetColor(void* ptr, void* color)
{
	static_cast<QQuickWindow*>(ptr)->setColor(*static_cast<QColor*>(color));
}

void QQuickWindow_QQuickWindow_SetDefaultAlphaBuffer(char useAlpha)
{
	QQuickWindow::setDefaultAlphaBuffer(useAlpha != 0);
}

void QQuickWindow_SetPersistentOpenGLContext(void* ptr, char persistent)
{
	static_cast<QQuickWindow*>(ptr)->setPersistentOpenGLContext(persistent != 0);
}

void QQuickWindow_SetPersistentSceneGraph(void* ptr, char persistent)
{
	static_cast<QQuickWindow*>(ptr)->setPersistentSceneGraph(persistent != 0);
}

void QQuickWindow_SetRenderTarget(void* ptr, void* fbo)
{
	static_cast<QQuickWindow*>(ptr)->setRenderTarget(static_cast<QOpenGLFramebufferObject*>(fbo));
}

void QQuickWindow_SetRenderTarget2(void* ptr, unsigned int fboId, void* size)
{
	static_cast<QQuickWindow*>(ptr)->setRenderTarget(fboId, *static_cast<QSize*>(size));
}

void QQuickWindow_QQuickWindow_SetSceneGraphBackend(long long api)
{
	QQuickWindow::setSceneGraphBackend(static_cast<QSGRendererInterface::GraphicsApi>(api));
}

void QQuickWindow_QQuickWindow_SetSceneGraphBackend2(struct QtQuick_PackedString backend)
{
	QQuickWindow::setSceneGraphBackend(QString::fromUtf8(backend.data, backend.len));
}

void QQuickWindow_ShowEventDefault(void* ptr, void* vqs)
{
	if (dynamic_cast<QQuickView*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickView*>(ptr)->QQuickView::showEvent(static_cast<QShowEvent*>(vqs));
	} else {
		static_cast<QQuickWindow*>(ptr)->QQuickWindow::showEvent(static_cast<QShowEvent*>(vqs));
	}
}

void QQuickWindow_Update(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QQuickWindow*>(ptr), "update");
}

void QQuickWindow_UpdateDefault(void* ptr)
{
	if (dynamic_cast<QQuickView*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickView*>(ptr)->QQuickView::update();
	} else {
		static_cast<QQuickWindow*>(ptr)->QQuickWindow::update();
	}
}

void QQuickWindow_WheelEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QQuickView*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickView*>(ptr)->QQuickView::wheelEvent(static_cast<QWheelEvent*>(event));
	} else {
		static_cast<QQuickWindow*>(ptr)->QQuickWindow::wheelEvent(static_cast<QWheelEvent*>(event));
	}
}

void QQuickWindow_DestroyQQuickWindow(void* ptr)
{
	static_cast<QQuickWindow*>(ptr)->~QQuickWindow();
}

void QQuickWindow_DestroyQQuickWindowDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QQuickWindow_AccessibleRoot(void* ptr)
{
	return static_cast<QQuickWindow*>(ptr)->accessibleRoot();
}

void* QQuickWindow_Color(void* ptr)
{
	return new QColor(static_cast<QQuickWindow*>(ptr)->color());
}

void* QQuickWindow_OpenglContext(void* ptr)
{
	return static_cast<QQuickWindow*>(ptr)->openglContext();
}

void* QQuickWindow_RenderTarget(void* ptr)
{
	return static_cast<QQuickWindow*>(ptr)->renderTarget();
}

void* QQuickWindow_IncubationController(void* ptr)
{
	return static_cast<QQuickWindow*>(ptr)->incubationController();
}

void* QQuickWindow_ActiveFocusItem(void* ptr)
{
	return static_cast<QQuickWindow*>(ptr)->activeFocusItem();
}

void* QQuickWindow_ContentItem(void* ptr)
{
	return static_cast<QQuickWindow*>(ptr)->contentItem();
}

void* QQuickWindow_MouseGrabberItem(void* ptr)
{
	return static_cast<QQuickWindow*>(ptr)->mouseGrabberItem();
}

void* QQuickWindow_CreateImageNode(void* ptr)
{
	return static_cast<QQuickWindow*>(ptr)->createImageNode();
}

void* QQuickWindow_CreateRectangleNode(void* ptr)
{
	return static_cast<QQuickWindow*>(ptr)->createRectangleNode();
}

void* QQuickWindow_RendererInterface(void* ptr)
{
	return static_cast<QQuickWindow*>(ptr)->rendererInterface();
}

void* QQuickWindow_CreateTextureFromId(void* ptr, unsigned int id, void* size, long long options)
{
	return static_cast<QQuickWindow*>(ptr)->createTextureFromId(id, *static_cast<QSize*>(size), static_cast<QQuickWindow::CreateTextureOption>(options));
}

void* QQuickWindow_CreateTextureFromImage2(void* ptr, void* image)
{
	return static_cast<QQuickWindow*>(ptr)->createTextureFromImage(*static_cast<QImage*>(image));
}

void* QQuickWindow_CreateTextureFromImage(void* ptr, void* image, long long options)
{
	return static_cast<QQuickWindow*>(ptr)->createTextureFromImage(*static_cast<QImage*>(image), static_cast<QQuickWindow::CreateTextureOption>(options));
}

void* QQuickWindow_RenderTargetSize(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QQuickWindow*>(ptr)->renderTargetSize(); new QSize(tmpValue.width(), tmpValue.height()); });
}

char QQuickWindow_ClearBeforeRendering(void* ptr)
{
	return static_cast<QQuickWindow*>(ptr)->clearBeforeRendering();
}

char QQuickWindow_IsPersistentOpenGLContext(void* ptr)
{
	return static_cast<QQuickWindow*>(ptr)->isPersistentOpenGLContext();
}

char QQuickWindow_IsPersistentSceneGraph(void* ptr)
{
	return static_cast<QQuickWindow*>(ptr)->isPersistentSceneGraph();
}

char QQuickWindow_IsSceneGraphInitialized(void* ptr)
{
	return static_cast<QQuickWindow*>(ptr)->isSceneGraphInitialized();
}

double QQuickWindow_EffectiveDevicePixelRatio(void* ptr)
{
	return static_cast<QQuickWindow*>(ptr)->effectiveDevicePixelRatio();
}

unsigned int QQuickWindow_RenderTargetId(void* ptr)
{
	return static_cast<QQuickWindow*>(ptr)->renderTargetId();
}

void* QQuickWindow___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QQuickWindow___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QQuickWindow___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QQuickWindow___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QQuickWindow___findChildren_setList2(void* ptr, void* i)
{
	if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QQuickItem*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QQuickWindow___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QQuickWindow___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QQuickWindow___findChildren_setList3(void* ptr, void* i)
{
	if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QQuickItem*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QQuickWindow___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QQuickWindow___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QQuickWindow___findChildren_setList(void* ptr, void* i)
{
	if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QQuickItem*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QQuickWindow___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QQuickWindow___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QQuickWindow___children_setList(void* ptr, void* i)
{
	if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QQuickItem*>(i));
	} else {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QQuickWindow___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

char QQuickWindow_CloseDefault(void* ptr)
{
	if (dynamic_cast<QQuickView*>(static_cast<QObject*>(ptr))) {
		return static_cast<QQuickView*>(ptr)->QQuickView::close();
	} else {
		return static_cast<QQuickWindow*>(ptr)->QQuickWindow::close();
	}
}

char QQuickWindow_NativeEventDefault(void* ptr, void* eventType, void* message, long result)
{
	if (dynamic_cast<QQuickView*>(static_cast<QObject*>(ptr))) {
		return static_cast<QQuickView*>(ptr)->QQuickView::nativeEvent(*static_cast<QByteArray*>(eventType), message, &result);
	} else {
		return static_cast<QQuickWindow*>(ptr)->QQuickWindow::nativeEvent(*static_cast<QByteArray*>(eventType), message, &result);
	}
}

void QQuickWindow_AlertDefault(void* ptr, int msec)
{
	if (dynamic_cast<QQuickView*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickView*>(ptr)->QQuickView::alert(msec);
	} else {
		static_cast<QQuickWindow*>(ptr)->QQuickWindow::alert(msec);
	}
}

void QQuickWindow_HideDefault(void* ptr)
{
	if (dynamic_cast<QQuickView*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickView*>(ptr)->QQuickView::hide();
	} else {
		static_cast<QQuickWindow*>(ptr)->QQuickWindow::hide();
	}
}

void QQuickWindow_LowerDefault(void* ptr)
{
	if (dynamic_cast<QQuickView*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickView*>(ptr)->QQuickView::lower();
	} else {
		static_cast<QQuickWindow*>(ptr)->QQuickWindow::lower();
	}
}

void QQuickWindow_MoveEventDefault(void* ptr, void* ev)
{
	if (dynamic_cast<QQuickView*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickView*>(ptr)->QQuickView::moveEvent(static_cast<QMoveEvent*>(ev));
	} else {
		static_cast<QQuickWindow*>(ptr)->QQuickWindow::moveEvent(static_cast<QMoveEvent*>(ev));
	}
}

void QQuickWindow_RaiseDefault(void* ptr)
{
	if (dynamic_cast<QQuickView*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickView*>(ptr)->QQuickView::raise();
	} else {
		static_cast<QQuickWindow*>(ptr)->QQuickWindow::raise();
	}
}

void QQuickWindow_RequestActivateDefault(void* ptr)
{
	if (dynamic_cast<QQuickView*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickView*>(ptr)->QQuickView::requestActivate();
	} else {
		static_cast<QQuickWindow*>(ptr)->QQuickWindow::requestActivate();
	}
}

void QQuickWindow_RequestUpdateDefault(void* ptr)
{
	if (dynamic_cast<QQuickView*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickView*>(ptr)->QQuickView::requestUpdate();
	} else {
		static_cast<QQuickWindow*>(ptr)->QQuickWindow::requestUpdate();
	}
}

void QQuickWindow_SetHeightDefault(void* ptr, int arg)
{
	if (dynamic_cast<QQuickView*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickView*>(ptr)->QQuickView::setHeight(arg);
	} else {
		static_cast<QQuickWindow*>(ptr)->QQuickWindow::setHeight(arg);
	}
}

void QQuickWindow_SetMaximumHeightDefault(void* ptr, int h)
{
	if (dynamic_cast<QQuickView*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickView*>(ptr)->QQuickView::setMaximumHeight(h);
	} else {
		static_cast<QQuickWindow*>(ptr)->QQuickWindow::setMaximumHeight(h);
	}
}

void QQuickWindow_SetMaximumWidthDefault(void* ptr, int w)
{
	if (dynamic_cast<QQuickView*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickView*>(ptr)->QQuickView::setMaximumWidth(w);
	} else {
		static_cast<QQuickWindow*>(ptr)->QQuickWindow::setMaximumWidth(w);
	}
}

void QQuickWindow_SetMinimumHeightDefault(void* ptr, int h)
{
	if (dynamic_cast<QQuickView*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickView*>(ptr)->QQuickView::setMinimumHeight(h);
	} else {
		static_cast<QQuickWindow*>(ptr)->QQuickWindow::setMinimumHeight(h);
	}
}

void QQuickWindow_SetMinimumWidthDefault(void* ptr, int w)
{
	if (dynamic_cast<QQuickView*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickView*>(ptr)->QQuickView::setMinimumWidth(w);
	} else {
		static_cast<QQuickWindow*>(ptr)->QQuickWindow::setMinimumWidth(w);
	}
}

void QQuickWindow_SetTitleDefault(void* ptr, struct QtQuick_PackedString vqs)
{
	if (dynamic_cast<QQuickView*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickView*>(ptr)->QQuickView::setTitle(QString::fromUtf8(vqs.data, vqs.len));
	} else {
		static_cast<QQuickWindow*>(ptr)->QQuickWindow::setTitle(QString::fromUtf8(vqs.data, vqs.len));
	}
}

void QQuickWindow_SetVisibleDefault(void* ptr, char visible)
{
	if (dynamic_cast<QQuickView*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickView*>(ptr)->QQuickView::setVisible(visible != 0);
	} else {
		static_cast<QQuickWindow*>(ptr)->QQuickWindow::setVisible(visible != 0);
	}
}

void QQuickWindow_SetWidthDefault(void* ptr, int arg)
{
	if (dynamic_cast<QQuickView*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickView*>(ptr)->QQuickView::setWidth(arg);
	} else {
		static_cast<QQuickWindow*>(ptr)->QQuickWindow::setWidth(arg);
	}
}

void QQuickWindow_SetXDefault(void* ptr, int arg)
{
	if (dynamic_cast<QQuickView*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickView*>(ptr)->QQuickView::setX(arg);
	} else {
		static_cast<QQuickWindow*>(ptr)->QQuickWindow::setX(arg);
	}
}

void QQuickWindow_SetYDefault(void* ptr, int arg)
{
	if (dynamic_cast<QQuickView*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickView*>(ptr)->QQuickView::setY(arg);
	} else {
		static_cast<QQuickWindow*>(ptr)->QQuickWindow::setY(arg);
	}
}

void QQuickWindow_ShowDefault(void* ptr)
{
	if (dynamic_cast<QQuickView*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickView*>(ptr)->QQuickView::show();
	} else {
		static_cast<QQuickWindow*>(ptr)->QQuickWindow::show();
	}
}

void QQuickWindow_ShowFullScreenDefault(void* ptr)
{
	if (dynamic_cast<QQuickView*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickView*>(ptr)->QQuickView::showFullScreen();
	} else {
		static_cast<QQuickWindow*>(ptr)->QQuickWindow::showFullScreen();
	}
}

void QQuickWindow_ShowMaximizedDefault(void* ptr)
{
	if (dynamic_cast<QQuickView*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickView*>(ptr)->QQuickView::showMaximized();
	} else {
		static_cast<QQuickWindow*>(ptr)->QQuickWindow::showMaximized();
	}
}

void QQuickWindow_ShowMinimizedDefault(void* ptr)
{
	if (dynamic_cast<QQuickView*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickView*>(ptr)->QQuickView::showMinimized();
	} else {
		static_cast<QQuickWindow*>(ptr)->QQuickWindow::showMinimized();
	}
}

void QQuickWindow_ShowNormalDefault(void* ptr)
{
	if (dynamic_cast<QQuickView*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickView*>(ptr)->QQuickView::showNormal();
	} else {
		static_cast<QQuickWindow*>(ptr)->QQuickWindow::showNormal();
	}
}

void QQuickWindow_TabletEventDefault(void* ptr, void* ev)
{
	if (dynamic_cast<QQuickView*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickView*>(ptr)->QQuickView::tabletEvent(static_cast<QTabletEvent*>(ev));
	} else {
		static_cast<QQuickWindow*>(ptr)->QQuickWindow::tabletEvent(static_cast<QTabletEvent*>(ev));
	}
}

void QQuickWindow_TouchEventDefault(void* ptr, void* ev)
{
	if (dynamic_cast<QQuickView*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickView*>(ptr)->QQuickView::touchEvent(static_cast<QTouchEvent*>(ev));
	} else {
		static_cast<QQuickWindow*>(ptr)->QQuickWindow::touchEvent(static_cast<QTouchEvent*>(ev));
	}
}

void* QQuickWindow_FocusObjectDefault(void* ptr)
{
	if (dynamic_cast<QQuickView*>(static_cast<QObject*>(ptr))) {
		return static_cast<QQuickView*>(ptr)->QQuickView::focusObject();
	} else {
		return static_cast<QQuickWindow*>(ptr)->QQuickWindow::focusObject();
	}
}

void* QQuickWindow_SizeDefault(void* ptr)
{
	if (dynamic_cast<QQuickView*>(static_cast<QObject*>(ptr))) {
		return ({ QSize tmpValue = static_cast<QQuickView*>(ptr)->QQuickView::size(); new QSize(tmpValue.width(), tmpValue.height()); });
	} else {
		return ({ QSize tmpValue = static_cast<QQuickWindow*>(ptr)->QQuickWindow::size(); new QSize(tmpValue.width(), tmpValue.height()); });
	}
}

void* QQuickWindow_FormatDefault(void* ptr)
{
	if (dynamic_cast<QQuickView*>(static_cast<QObject*>(ptr))) {
		return new QSurfaceFormat(static_cast<QQuickView*>(ptr)->QQuickView::format());
	} else {
		return new QSurfaceFormat(static_cast<QQuickWindow*>(ptr)->QQuickWindow::format());
	}
}

long long QQuickWindow_SurfaceTypeDefault(void* ptr)
{
	if (dynamic_cast<QQuickView*>(static_cast<QObject*>(ptr))) {
		return static_cast<QQuickView*>(ptr)->QQuickView::surfaceType();
	} else {
		return static_cast<QQuickWindow*>(ptr)->QQuickWindow::surfaceType();
	}
}

char QQuickWindow_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QQuickView*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(watched))) {
			return static_cast<QQuickView*>(ptr)->QQuickView::eventFilter(static_cast<QQuickItem*>(watched), static_cast<QEvent*>(event));
		} else {
			return static_cast<QQuickView*>(ptr)->QQuickView::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
		}
	} else {
		if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(watched))) {
			return static_cast<QQuickWindow*>(ptr)->QQuickWindow::eventFilter(static_cast<QQuickItem*>(watched), static_cast<QEvent*>(event));
		} else {
			return static_cast<QQuickWindow*>(ptr)->QQuickWindow::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
		}
	}
}

void QQuickWindow_ChildEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QQuickView*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickView*>(ptr)->QQuickView::childEvent(static_cast<QChildEvent*>(event));
	} else {
		static_cast<QQuickWindow*>(ptr)->QQuickWindow::childEvent(static_cast<QChildEvent*>(event));
	}
}

void QQuickWindow_ConnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QQuickView*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickView*>(ptr)->QQuickView::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QQuickWindow*>(ptr)->QQuickWindow::connectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

void QQuickWindow_CustomEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QQuickView*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickView*>(ptr)->QQuickView::customEvent(static_cast<QEvent*>(event));
	} else {
		static_cast<QQuickWindow*>(ptr)->QQuickWindow::customEvent(static_cast<QEvent*>(event));
	}
}

void QQuickWindow_DeleteLaterDefault(void* ptr)
{
	if (dynamic_cast<QQuickView*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickView*>(ptr)->QQuickView::deleteLater();
	} else {
		static_cast<QQuickWindow*>(ptr)->QQuickWindow::deleteLater();
	}
}

void QQuickWindow_DisconnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QQuickView*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickView*>(ptr)->QQuickView::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QQuickWindow*>(ptr)->QQuickWindow::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

void QQuickWindow_TimerEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QQuickView*>(static_cast<QObject*>(ptr))) {
		static_cast<QQuickView*>(ptr)->QQuickView::timerEvent(static_cast<QTimerEvent*>(event));
	} else {
		static_cast<QQuickWindow*>(ptr)->QQuickWindow::timerEvent(static_cast<QTimerEvent*>(event));
	}
}

void* QQuickWindow_MetaObjectDefault(void* ptr)
{
	if (dynamic_cast<QQuickView*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QQuickView*>(ptr)->QQuickView::metaObject());
	} else {
		return const_cast<QMetaObject*>(static_cast<QQuickWindow*>(ptr)->QQuickWindow::metaObject());
	}
}

class MyQSGAbstractRenderer: public QSGAbstractRenderer
{
public:
	void renderScene(GLuint fboId) { callbackQSGAbstractRenderer_RenderScene(this, fboId); };
	void Signal_SceneGraphChanged() { callbackQSGAbstractRenderer_SceneGraphChanged(this); };
	bool event(QEvent * e) { return callbackQSGAbstractRenderer_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSGAbstractRenderer_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQSGAbstractRenderer_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSGAbstractRenderer_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQSGAbstractRenderer_CustomEvent(this, event); };
	void deleteLater() { callbackQSGAbstractRenderer_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQSGAbstractRenderer_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSGAbstractRenderer_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtQuick_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQSGAbstractRenderer_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQSGAbstractRenderer_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSGAbstractRenderer_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQSGAbstractRenderer*)

int QSGAbstractRenderer_QSGAbstractRenderer_QRegisterMetaType(){qRegisterMetaType<QSGAbstractRenderer*>(); return qRegisterMetaType<MyQSGAbstractRenderer*>();}

void QSGAbstractRenderer_RenderScene(void* ptr, unsigned int fboId)
{
	static_cast<QSGAbstractRenderer*>(ptr)->renderScene(fboId);
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

void QSGAbstractRenderer_SetClearMode(void* ptr, long long mode)
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

long long QSGAbstractRenderer_ClearMode(void* ptr)
{
	return static_cast<QSGAbstractRenderer*>(ptr)->clearMode();
}

void* QSGAbstractRenderer_ClearColor(void* ptr)
{
	return new QColor(static_cast<QSGAbstractRenderer*>(ptr)->clearColor());
}

void* QSGAbstractRenderer_ProjectionMatrix(void* ptr)
{
	return new QMatrix4x4(static_cast<QSGAbstractRenderer*>(ptr)->projectionMatrix());
}

void* QSGAbstractRenderer_DeviceRect(void* ptr)
{
	return ({ QRect tmpValue = static_cast<QSGAbstractRenderer*>(ptr)->deviceRect(); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void* QSGAbstractRenderer_ViewportRect(void* ptr)
{
	return ({ QRect tmpValue = static_cast<QSGAbstractRenderer*>(ptr)->viewportRect(); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void* QSGAbstractRenderer___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QSGAbstractRenderer___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QSGAbstractRenderer___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QSGAbstractRenderer___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QSGAbstractRenderer___findChildren_setList2(void* ptr, void* i)
{
	if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QQuickItem*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QSGAbstractRenderer___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QSGAbstractRenderer___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QSGAbstractRenderer___findChildren_setList3(void* ptr, void* i)
{
	if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QQuickItem*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QSGAbstractRenderer___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QSGAbstractRenderer___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QSGAbstractRenderer___findChildren_setList(void* ptr, void* i)
{
	if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QQuickItem*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QSGAbstractRenderer___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QSGAbstractRenderer___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QSGAbstractRenderer___children_setList(void* ptr, void* i)
{
	if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QQuickItem*>(i));
	} else {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QSGAbstractRenderer___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

char QSGAbstractRenderer_EventDefault(void* ptr, void* e)
{
		return static_cast<QSGAbstractRenderer*>(ptr)->QSGAbstractRenderer::event(static_cast<QEvent*>(e));
}

char QSGAbstractRenderer_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(watched))) {
		return static_cast<QSGAbstractRenderer*>(ptr)->QSGAbstractRenderer::eventFilter(static_cast<QQuickItem*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QSGAbstractRenderer*>(ptr)->QSGAbstractRenderer::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

void QSGAbstractRenderer_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QSGAbstractRenderer*>(ptr)->QSGAbstractRenderer::childEvent(static_cast<QChildEvent*>(event));
}

void QSGAbstractRenderer_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QSGAbstractRenderer*>(ptr)->QSGAbstractRenderer::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSGAbstractRenderer_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QSGAbstractRenderer*>(ptr)->QSGAbstractRenderer::customEvent(static_cast<QEvent*>(event));
}

void QSGAbstractRenderer_DeleteLaterDefault(void* ptr)
{
		static_cast<QSGAbstractRenderer*>(ptr)->QSGAbstractRenderer::deleteLater();
}

void QSGAbstractRenderer_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QSGAbstractRenderer*>(ptr)->QSGAbstractRenderer::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSGAbstractRenderer_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QSGAbstractRenderer*>(ptr)->QSGAbstractRenderer::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QSGAbstractRenderer_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QSGAbstractRenderer*>(ptr)->QSGAbstractRenderer::metaObject());
}

class MyQSGBasicGeometryNode: public QSGBasicGeometryNode
{
public:
	void preprocess() { callbackQSGNode_Preprocess(this); };
	bool isSubtreeBlocked() const { return callbackQSGNode_IsSubtreeBlocked(const_cast<void*>(static_cast<const void*>(this))) != 0; };
};

void* QSGBasicGeometryNode_Geometry2(void* ptr)
{
	return static_cast<QSGBasicGeometryNode*>(ptr)->geometry();
}

void QSGBasicGeometryNode_SetGeometry(void* ptr, void* geometry)
{
	static_cast<QSGBasicGeometryNode*>(ptr)->setGeometry(static_cast<QSGGeometry*>(geometry));
}

void QSGBasicGeometryNode_DestroyQSGBasicGeometryNode(void* ptr)
{
	static_cast<QSGBasicGeometryNode*>(ptr)->~QSGBasicGeometryNode();
}

void* QSGBasicGeometryNode_Geometry(void* ptr)
{
	return const_cast<QSGGeometry*>(static_cast<QSGBasicGeometryNode*>(ptr)->geometry());
}

class MyQSGClipNode: public QSGClipNode
{
public:
	MyQSGClipNode() : QSGClipNode() {};
	void preprocess() { callbackQSGNode_Preprocess(this); };
	bool isSubtreeBlocked() const { return callbackQSGNode_IsSubtreeBlocked(const_cast<void*>(static_cast<const void*>(this))) != 0; };
};

void* QSGClipNode_NewQSGClipNode()
{
	return new MyQSGClipNode();
}

void QSGClipNode_SetClipRect(void* ptr, void* rect)
{
	static_cast<QSGClipNode*>(ptr)->setClipRect(*static_cast<QRectF*>(rect));
}

void QSGClipNode_SetIsRectangular(void* ptr, char rectHint)
{
	static_cast<QSGClipNode*>(ptr)->setIsRectangular(rectHint != 0);
}

void QSGClipNode_DestroyQSGClipNode(void* ptr)
{
	static_cast<QSGClipNode*>(ptr)->~QSGClipNode();
}

void* QSGClipNode_ClipRect(void* ptr)
{
	return ({ QRectF tmpValue = static_cast<QSGClipNode*>(ptr)->clipRect(); new QRectF(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

char QSGClipNode_IsRectangular(void* ptr)
{
	return static_cast<QSGClipNode*>(ptr)->isRectangular();
}

class MyQSGDynamicTexture: public QSGDynamicTexture
{
public:
	bool updateTexture() { return callbackQSGDynamicTexture_UpdateTexture(this) != 0; };
	void bind() { callbackQSGDynamicTexture_Bind(this); };
	QRectF normalizedTextureSubRect() const { return *static_cast<QRectF*>(callbackQSGTexture_NormalizedTextureSubRect(const_cast<void*>(static_cast<const void*>(this)))); };
	QSGTexture * removedFromAtlas() const { return static_cast<QSGTexture*>(callbackQSGTexture_RemovedFromAtlas(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize textureSize() const { return *static_cast<QSize*>(callbackQSGDynamicTexture_TextureSize(const_cast<void*>(static_cast<const void*>(this)))); };
	bool hasAlphaChannel() const { return callbackQSGDynamicTexture_HasAlphaChannel(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	bool hasMipmaps() const { return callbackQSGDynamicTexture_HasMipmaps(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	bool isAtlasTexture() const { return callbackQSGTexture_IsAtlasTexture(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int textureId() const { return callbackQSGDynamicTexture_TextureId(const_cast<void*>(static_cast<const void*>(this))); };
	bool event(QEvent * e) { return callbackQSGTexture_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSGTexture_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQSGTexture_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSGTexture_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQSGTexture_CustomEvent(this, event); };
	void deleteLater() { callbackQSGTexture_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQSGTexture_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSGTexture_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtQuick_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQSGTexture_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQSGTexture_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSGTexture_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQSGDynamicTexture*)

int QSGDynamicTexture_QSGDynamicTexture_QRegisterMetaType(){qRegisterMetaType<QSGDynamicTexture*>(); return qRegisterMetaType<MyQSGDynamicTexture*>();}

char QSGDynamicTexture_UpdateTexture(void* ptr)
{
	return static_cast<QSGDynamicTexture*>(ptr)->updateTexture();
}

void QSGDynamicTexture_Bind(void* ptr)
{
	static_cast<QSGDynamicTexture*>(ptr)->bind();
}

void QSGDynamicTexture_BindDefault(void* ptr)
{
	Q_UNUSED(ptr);
	
}

void* QSGDynamicTexture_TextureSize(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QSGDynamicTexture*>(ptr)->textureSize(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QSGDynamicTexture_TextureSizeDefault(void* ptr)
{
	Q_UNUSED(ptr);
	
}

char QSGDynamicTexture_HasAlphaChannel(void* ptr)
{
	return static_cast<QSGDynamicTexture*>(ptr)->hasAlphaChannel();
}

char QSGDynamicTexture_HasAlphaChannelDefault(void* ptr)
{
	Q_UNUSED(ptr);
	
}

char QSGDynamicTexture_HasMipmaps(void* ptr)
{
	return static_cast<QSGDynamicTexture*>(ptr)->hasMipmaps();
}

char QSGDynamicTexture_HasMipmapsDefault(void* ptr)
{
	Q_UNUSED(ptr);
	
}

int QSGDynamicTexture_TextureId(void* ptr)
{
	return static_cast<QSGDynamicTexture*>(ptr)->textureId();
}

int QSGDynamicTexture_TextureIdDefault(void* ptr)
{
	Q_UNUSED(ptr);
	
}

class MyQSGEngine: public QSGEngine
{
public:
	MyQSGEngine(QObject *parent = Q_NULLPTR) : QSGEngine(parent) {QSGEngine_QSGEngine_QRegisterMetaType();};
	bool event(QEvent * e) { return callbackQSGEngine_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSGEngine_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQSGEngine_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSGEngine_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQSGEngine_CustomEvent(this, event); };
	void deleteLater() { callbackQSGEngine_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQSGEngine_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSGEngine_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtQuick_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQSGEngine_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQSGEngine_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSGEngine_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQSGEngine*)

int QSGEngine_QSGEngine_QRegisterMetaType(){qRegisterMetaType<QSGEngine*>(); return qRegisterMetaType<MyQSGEngine*>();}

void* QSGEngine_NewQSGEngine(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQSGEngine(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQSGEngine(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQSGEngine(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQSGEngine(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQSGEngine(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQSGEngine(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQSGEngine(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQSGEngine(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQSGEngine(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQSGEngine(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQSGEngine(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQSGEngine(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQSGEngine(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQSGEngine(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQSGEngine(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQSGEngine(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQSGEngine(static_cast<QWindow*>(parent));
	} else {
		return new MyQSGEngine(static_cast<QObject*>(parent));
	}
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

void* QSGEngine_CreateRenderer(void* ptr)
{
	return static_cast<QSGEngine*>(ptr)->createRenderer();
}

void* QSGEngine_CreateImageNode(void* ptr)
{
	return static_cast<QSGEngine*>(ptr)->createImageNode();
}

void* QSGEngine_CreateRectangleNode(void* ptr)
{
	return static_cast<QSGEngine*>(ptr)->createRectangleNode();
}

void* QSGEngine_RendererInterface(void* ptr)
{
	return static_cast<QSGEngine*>(ptr)->rendererInterface();
}

void* QSGEngine_CreateTextureFromId(void* ptr, unsigned int id, void* size, long long options)
{
	return static_cast<QSGEngine*>(ptr)->createTextureFromId(id, *static_cast<QSize*>(size), static_cast<QSGEngine::CreateTextureOption>(options));
}

void* QSGEngine_CreateTextureFromImage(void* ptr, void* image, long long options)
{
	return static_cast<QSGEngine*>(ptr)->createTextureFromImage(*static_cast<QImage*>(image), static_cast<QSGEngine::CreateTextureOption>(options));
}

void* QSGEngine___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QSGEngine___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QSGEngine___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QSGEngine___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QSGEngine___findChildren_setList2(void* ptr, void* i)
{
	if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QQuickItem*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QSGEngine___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QSGEngine___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QSGEngine___findChildren_setList3(void* ptr, void* i)
{
	if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QQuickItem*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QSGEngine___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QSGEngine___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QSGEngine___findChildren_setList(void* ptr, void* i)
{
	if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QQuickItem*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QSGEngine___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QSGEngine___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QSGEngine___children_setList(void* ptr, void* i)
{
	if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QQuickItem*>(i));
	} else {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QSGEngine___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

char QSGEngine_EventDefault(void* ptr, void* e)
{
		return static_cast<QSGEngine*>(ptr)->QSGEngine::event(static_cast<QEvent*>(e));
}

char QSGEngine_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(watched))) {
		return static_cast<QSGEngine*>(ptr)->QSGEngine::eventFilter(static_cast<QQuickItem*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QSGEngine*>(ptr)->QSGEngine::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

void QSGEngine_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QSGEngine*>(ptr)->QSGEngine::childEvent(static_cast<QChildEvent*>(event));
}

void QSGEngine_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QSGEngine*>(ptr)->QSGEngine::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSGEngine_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QSGEngine*>(ptr)->QSGEngine::customEvent(static_cast<QEvent*>(event));
}

void QSGEngine_DeleteLaterDefault(void* ptr)
{
		static_cast<QSGEngine*>(ptr)->QSGEngine::deleteLater();
}

void QSGEngine_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QSGEngine*>(ptr)->QSGEngine::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSGEngine_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QSGEngine*>(ptr)->QSGEngine::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QSGEngine_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QSGEngine*>(ptr)->QSGEngine::metaObject());
}

class MyQSGFlatColorMaterial: public QSGFlatColorMaterial
{
public:
	MyQSGFlatColorMaterial() : QSGFlatColorMaterial() {};
	QSGMaterialShader * createShader() const { return static_cast<QSGMaterialShader*>(callbackQSGFlatColorMaterial_CreateShader(const_cast<void*>(static_cast<const void*>(this)))); };
	QSGMaterialType * type() const { return static_cast<QSGMaterialType*>(callbackQSGFlatColorMaterial_Type(const_cast<void*>(static_cast<const void*>(this)))); };
	int compare(const QSGMaterial * other) const { return callbackQSGMaterial_Compare(const_cast<void*>(static_cast<const void*>(this)), const_cast<QSGMaterial*>(other)); };
};

void* QSGFlatColorMaterial_NewQSGFlatColorMaterial()
{
	return new MyQSGFlatColorMaterial();
}

void QSGFlatColorMaterial_SetColor(void* ptr, void* color)
{
	static_cast<QSGFlatColorMaterial*>(ptr)->setColor(*static_cast<QColor*>(color));
}

void* QSGFlatColorMaterial_Color(void* ptr)
{
	return const_cast<QColor*>(&static_cast<QSGFlatColorMaterial*>(ptr)->color());
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
	 ~MyQSGGeometry() { callbackQSGGeometry_DestroyQSGGeometry(this); };
};

unsigned short QSGGeometry_IndexDataAsUShort(void* ptr)
{
	return *static_cast<QSGGeometry*>(ptr)->indexDataAsUShort();
}

unsigned int QSGGeometry_IndexDataAsUInt(void* ptr)
{
	return *static_cast<QSGGeometry*>(ptr)->indexDataAsUInt();
}

void* QSGGeometry_IndexData(void* ptr)
{
	return static_cast<QSGGeometry*>(ptr)->indexData();
}

void* QSGGeometry_VertexData(void* ptr)
{
	return static_cast<QSGGeometry*>(ptr)->vertexData();
}

void QSGGeometry_Allocate(void* ptr, int vertexCount, int indexCount)
{
	static_cast<QSGGeometry*>(ptr)->allocate(vertexCount, indexCount);
}

void QSGGeometry_MarkIndexDataDirty(void* ptr)
{
	static_cast<QSGGeometry*>(ptr)->markIndexDataDirty();
}

void QSGGeometry_MarkVertexDataDirty(void* ptr)
{
	static_cast<QSGGeometry*>(ptr)->markVertexDataDirty();
}

void QSGGeometry_SetDrawingMode(void* ptr, unsigned int mode)
{
	static_cast<QSGGeometry*>(ptr)->setDrawingMode(mode);
}

void QSGGeometry_SetIndexDataPattern(void* ptr, long long p)
{
	static_cast<QSGGeometry*>(ptr)->setIndexDataPattern(static_cast<QSGGeometry::DataPattern>(p));
}

void QSGGeometry_SetLineWidth(void* ptr, float width)
{
	static_cast<QSGGeometry*>(ptr)->setLineWidth(width);
}

void QSGGeometry_SetVertexDataPattern(void* ptr, long long p)
{
	static_cast<QSGGeometry*>(ptr)->setVertexDataPattern(static_cast<QSGGeometry::DataPattern>(p));
}

void QSGGeometry_QSGGeometry_UpdateColoredRectGeometry(void* g, void* rect)
{
	QSGGeometry::updateColoredRectGeometry(static_cast<QSGGeometry*>(g), *static_cast<QRectF*>(rect));
}

void QSGGeometry_QSGGeometry_UpdateRectGeometry(void* g, void* rect)
{
	QSGGeometry::updateRectGeometry(static_cast<QSGGeometry*>(g), *static_cast<QRectF*>(rect));
}

void QSGGeometry_QSGGeometry_UpdateTexturedRectGeometry(void* g, void* rect, void* textureRect)
{
	QSGGeometry::updateTexturedRectGeometry(static_cast<QSGGeometry*>(g), *static_cast<QRectF*>(rect), *static_cast<QRectF*>(textureRect));
}

void QSGGeometry_DestroyQSGGeometry(void* ptr)
{
	static_cast<QSGGeometry*>(ptr)->~QSGGeometry();
}

void QSGGeometry_DestroyQSGGeometryDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

long long QSGGeometry_IndexDataPattern(void* ptr)
{
	return static_cast<QSGGeometry*>(ptr)->indexDataPattern();
}

long long QSGGeometry_VertexDataPattern(void* ptr)
{
	return static_cast<QSGGeometry*>(ptr)->vertexDataPattern();
}

unsigned short QSGGeometry_IndexDataAsUShort2(void* ptr)
{
	return *static_cast<QSGGeometry*>(ptr)->indexDataAsUShort();
}

unsigned int QSGGeometry_IndexDataAsUInt2(void* ptr)
{
	return *static_cast<QSGGeometry*>(ptr)->indexDataAsUInt();
}

void* QSGGeometry_IndexData2(void* ptr)
{
	return const_cast<void*>(static_cast<QSGGeometry*>(ptr)->indexData());
}

void* QSGGeometry_VertexData2(void* ptr)
{
	return const_cast<void*>(static_cast<QSGGeometry*>(ptr)->vertexData());
}

float QSGGeometry_LineWidth(void* ptr)
{
	return static_cast<QSGGeometry*>(ptr)->lineWidth();
}

int QSGGeometry_AttributeCount(void* ptr)
{
	return static_cast<QSGGeometry*>(ptr)->attributeCount();
}

int QSGGeometry_IndexCount(void* ptr)
{
	return static_cast<QSGGeometry*>(ptr)->indexCount();
}

int QSGGeometry_IndexType(void* ptr)
{
	return static_cast<QSGGeometry*>(ptr)->indexType();
}

int QSGGeometry_SizeOfIndex(void* ptr)
{
	return static_cast<QSGGeometry*>(ptr)->sizeOfIndex();
}

int QSGGeometry_SizeOfVertex(void* ptr)
{
	return static_cast<QSGGeometry*>(ptr)->sizeOfVertex();
}

int QSGGeometry_VertexCount(void* ptr)
{
	return static_cast<QSGGeometry*>(ptr)->vertexCount();
}

unsigned int QSGGeometry_DrawingMode(void* ptr)
{
	return static_cast<QSGGeometry*>(ptr)->drawingMode();
}

class MyQSGGeometryNode: public QSGGeometryNode
{
public:
	MyQSGGeometryNode() : QSGGeometryNode() {};
	void preprocess() { callbackQSGNode_Preprocess(this); };
	bool isSubtreeBlocked() const { return callbackQSGNode_IsSubtreeBlocked(const_cast<void*>(static_cast<const void*>(this))) != 0; };
};

void* QSGGeometryNode_NewQSGGeometryNode()
{
	return new MyQSGGeometryNode();
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

void* QSGGeometryNode_Material(void* ptr)
{
	return static_cast<QSGGeometryNode*>(ptr)->material();
}

void* QSGGeometryNode_OpaqueMaterial(void* ptr)
{
	return static_cast<QSGGeometryNode*>(ptr)->opaqueMaterial();
}

class MyQSGImageNode: public QSGImageNode
{
public:
	void setFiltering(QSGTexture::Filtering filtering) { callbackQSGImageNode_SetFiltering(this, filtering); };
	void setMipmapFiltering(QSGTexture::Filtering filtering) { callbackQSGImageNode_SetMipmapFiltering(this, filtering); };
	void setOwnsTexture(bool owns) { callbackQSGImageNode_SetOwnsTexture(this, owns); };
	void setRect(const QRectF & rect) { callbackQSGImageNode_SetRect(this, const_cast<QRectF*>(&rect)); };
	void setSourceRect(const QRectF & rect) { callbackQSGImageNode_SetSourceRect(this, const_cast<QRectF*>(&rect)); };
	void setTexture(QSGTexture * texture) { callbackQSGImageNode_SetTexture(this, texture); };
	void setTextureCoordinatesTransform(QSGImageNode::TextureCoordinatesTransformMode mode) { callbackQSGImageNode_SetTextureCoordinatesTransform(this, mode); };
	 ~MyQSGImageNode() { callbackQSGImageNode_DestroyQSGImageNode(this); };
	QRectF rect() const { return *static_cast<QRectF*>(callbackQSGImageNode_Rect(const_cast<void*>(static_cast<const void*>(this)))); };
	QRectF sourceRect() const { return *static_cast<QRectF*>(callbackQSGImageNode_SourceRect(const_cast<void*>(static_cast<const void*>(this)))); };
	QSGTexture * texture() const { return static_cast<QSGTexture*>(callbackQSGImageNode_Texture(const_cast<void*>(static_cast<const void*>(this)))); };
	QSGTexture::Filtering filtering() const { return static_cast<QSGTexture::Filtering>(callbackQSGImageNode_Filtering(const_cast<void*>(static_cast<const void*>(this)))); };
	QSGTexture::Filtering mipmapFiltering() const { return static_cast<QSGTexture::Filtering>(callbackQSGImageNode_MipmapFiltering(const_cast<void*>(static_cast<const void*>(this)))); };
	TextureCoordinatesTransformMode textureCoordinatesTransform() const { return static_cast<QSGImageNode::TextureCoordinatesTransformFlag>(callbackQSGImageNode_TextureCoordinatesTransform(const_cast<void*>(static_cast<const void*>(this)))); };
	bool ownsTexture() const { return callbackQSGImageNode_OwnsTexture(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	void preprocess() { callbackQSGNode_Preprocess(this); };
	bool isSubtreeBlocked() const { return callbackQSGNode_IsSubtreeBlocked(const_cast<void*>(static_cast<const void*>(this))) != 0; };
};

void QSGImageNode_QSGImageNode_RebuildGeometry(void* g, void* texture, void* rect, void* sourceRect, long long texCoordMode)
{
	QSGImageNode::rebuildGeometry(static_cast<QSGGeometry*>(g), static_cast<QSGTexture*>(texture), *static_cast<QRectF*>(rect), *static_cast<QRectF*>(sourceRect), static_cast<QSGImageNode::TextureCoordinatesTransformFlag>(texCoordMode));
}

void QSGImageNode_SetFiltering(void* ptr, long long filtering)
{
	static_cast<QSGImageNode*>(ptr)->setFiltering(static_cast<QSGTexture::Filtering>(filtering));
}

void QSGImageNode_SetMipmapFiltering(void* ptr, long long filtering)
{
	static_cast<QSGImageNode*>(ptr)->setMipmapFiltering(static_cast<QSGTexture::Filtering>(filtering));
}

void QSGImageNode_SetOwnsTexture(void* ptr, char owns)
{
	static_cast<QSGImageNode*>(ptr)->setOwnsTexture(owns != 0);
}

void QSGImageNode_SetRect(void* ptr, void* rect)
{
	static_cast<QSGImageNode*>(ptr)->setRect(*static_cast<QRectF*>(rect));
}

void QSGImageNode_SetRect2(void* ptr, double x, double y, double w, double h)
{
	static_cast<QSGImageNode*>(ptr)->setRect(x, y, w, h);
}

void QSGImageNode_SetSourceRect(void* ptr, void* rect)
{
	static_cast<QSGImageNode*>(ptr)->setSourceRect(*static_cast<QRectF*>(rect));
}

void QSGImageNode_SetSourceRect2(void* ptr, double x, double y, double w, double h)
{
	static_cast<QSGImageNode*>(ptr)->setSourceRect(x, y, w, h);
}

void QSGImageNode_SetTexture(void* ptr, void* texture)
{
	static_cast<QSGImageNode*>(ptr)->setTexture(static_cast<QSGTexture*>(texture));
}

void QSGImageNode_SetTextureCoordinatesTransform(void* ptr, long long mode)
{
	static_cast<QSGImageNode*>(ptr)->setTextureCoordinatesTransform(static_cast<QSGImageNode::TextureCoordinatesTransformFlag>(mode));
}

void QSGImageNode_DestroyQSGImageNode(void* ptr)
{
	static_cast<QSGImageNode*>(ptr)->~QSGImageNode();
}

void QSGImageNode_DestroyQSGImageNodeDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QSGImageNode_Rect(void* ptr)
{
	return ({ QRectF tmpValue = static_cast<QSGImageNode*>(ptr)->rect(); new QRectF(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void* QSGImageNode_SourceRect(void* ptr)
{
	return ({ QRectF tmpValue = static_cast<QSGImageNode*>(ptr)->sourceRect(); new QRectF(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void* QSGImageNode_Texture(void* ptr)
{
	return static_cast<QSGImageNode*>(ptr)->texture();
}

long long QSGImageNode_Filtering(void* ptr)
{
	return static_cast<QSGImageNode*>(ptr)->filtering();
}

long long QSGImageNode_MipmapFiltering(void* ptr)
{
	return static_cast<QSGImageNode*>(ptr)->mipmapFiltering();
}

long long QSGImageNode_TextureCoordinatesTransform(void* ptr)
{
	return static_cast<QSGImageNode*>(ptr)->textureCoordinatesTransform();
}

char QSGImageNode_OwnsTexture(void* ptr)
{
	return static_cast<QSGImageNode*>(ptr)->ownsTexture();
}

class MyQSGMaterial: public QSGMaterial
{
public:
	QSGMaterialShader * createShader() const { return static_cast<QSGMaterialShader*>(callbackQSGMaterial_CreateShader(const_cast<void*>(static_cast<const void*>(this)))); };
	QSGMaterialType * type() const { return static_cast<QSGMaterialType*>(callbackQSGMaterial_Type(const_cast<void*>(static_cast<const void*>(this)))); };
	int compare(const QSGMaterial * other) const { return callbackQSGMaterial_Compare(const_cast<void*>(static_cast<const void*>(this)), const_cast<QSGMaterial*>(other)); };
};

void QSGMaterial_SetFlag(void* ptr, long long flags, char on)
{
	static_cast<QSGMaterial*>(ptr)->setFlag(static_cast<QSGMaterial::Flag>(flags), on != 0);
}

long long QSGMaterial_Flags(void* ptr)
{
	return static_cast<QSGMaterial*>(ptr)->flags();
}

void* QSGMaterial_CreateShader(void* ptr)
{
	return static_cast<QSGMaterial*>(ptr)->createShader();
}

void* QSGMaterial_Type(void* ptr)
{
	return static_cast<QSGMaterial*>(ptr)->type();
}

int QSGMaterial_Compare(void* ptr, void* other)
{
	return static_cast<QSGMaterial*>(ptr)->compare(static_cast<QSGMaterial*>(other));
}

int QSGMaterial_CompareDefault(void* ptr, void* other)
{
	if (dynamic_cast<QSGVertexColorMaterial*>(static_cast<QSGMaterial*>(ptr))) {
		return static_cast<QSGVertexColorMaterial*>(ptr)->QSGVertexColorMaterial::compare(static_cast<QSGMaterial*>(other));
	} else if (dynamic_cast<QSGTextureMaterial*>(static_cast<QSGMaterial*>(ptr))) {
		return static_cast<QSGTextureMaterial*>(ptr)->QSGTextureMaterial::compare(static_cast<QSGMaterial*>(other));
	} else if (dynamic_cast<QSGOpaqueTextureMaterial*>(static_cast<QSGMaterial*>(ptr))) {
		return static_cast<QSGOpaqueTextureMaterial*>(ptr)->QSGOpaqueTextureMaterial::compare(static_cast<QSGMaterial*>(other));
	} else if (dynamic_cast<QSGFlatColorMaterial*>(static_cast<QSGMaterial*>(ptr))) {
		return static_cast<QSGFlatColorMaterial*>(ptr)->QSGFlatColorMaterial::compare(static_cast<QSGMaterial*>(other));
	} else {
		return static_cast<QSGMaterial*>(ptr)->QSGMaterial::compare(static_cast<QSGMaterial*>(other));
	}
}

class MyQSGMaterialShader: public QSGMaterialShader
{
public:
	const char * fragmentShader() const { return const_cast<const char*>(callbackQSGMaterialShader_FragmentShader(const_cast<void*>(static_cast<const void*>(this)))); };
	const char * vertexShader() const { return const_cast<const char*>(callbackQSGMaterialShader_VertexShader(const_cast<void*>(static_cast<const void*>(this)))); };
	void activate() { callbackQSGMaterialShader_Activate(this); };
	void compile() { callbackQSGMaterialShader_Compile(this); };
	void deactivate() { callbackQSGMaterialShader_Deactivate(this); };
	void initialize() { callbackQSGMaterialShader_Initialize(this); };
};

struct QtQuick_PackedString QSGMaterialShader_FragmentShader(void* ptr)
{
	return QtQuick_PackedString { const_cast<char*>(static_cast<QSGMaterialShader*>(ptr)->fragmentShader()), -1 };
}

struct QtQuick_PackedString QSGMaterialShader_FragmentShaderDefault(void* ptr)
{
		return QtQuick_PackedString { const_cast<char*>(static_cast<QSGMaterialShader*>(ptr)->QSGMaterialShader::fragmentShader()), -1 };
}

struct QtQuick_PackedString QSGMaterialShader_VertexShader(void* ptr)
{
	return QtQuick_PackedString { const_cast<char*>(static_cast<QSGMaterialShader*>(ptr)->vertexShader()), -1 };
}

struct QtQuick_PackedString QSGMaterialShader_VertexShaderDefault(void* ptr)
{
		return QtQuick_PackedString { const_cast<char*>(static_cast<QSGMaterialShader*>(ptr)->QSGMaterialShader::vertexShader()), -1 };
}

void* QSGMaterialShader_Program(void* ptr)
{
	return static_cast<QSGMaterialShader*>(ptr)->program();
}

void QSGMaterialShader_Activate(void* ptr)
{
	static_cast<QSGMaterialShader*>(ptr)->activate();
}

void QSGMaterialShader_ActivateDefault(void* ptr)
{
		static_cast<QSGMaterialShader*>(ptr)->QSGMaterialShader::activate();
}

void QSGMaterialShader_Compile(void* ptr)
{
	static_cast<QSGMaterialShader*>(ptr)->compile();
}

void QSGMaterialShader_CompileDefault(void* ptr)
{
		static_cast<QSGMaterialShader*>(ptr)->QSGMaterialShader::compile();
}

void QSGMaterialShader_Deactivate(void* ptr)
{
	static_cast<QSGMaterialShader*>(ptr)->deactivate();
}

void QSGMaterialShader_DeactivateDefault(void* ptr)
{
		static_cast<QSGMaterialShader*>(ptr)->QSGMaterialShader::deactivate();
}

void QSGMaterialShader_Initialize(void* ptr)
{
	static_cast<QSGMaterialShader*>(ptr)->initialize();
}

void QSGMaterialShader_InitializeDefault(void* ptr)
{
		static_cast<QSGMaterialShader*>(ptr)->QSGMaterialShader::initialize();
}

void QSGMaterialShader_SetShaderSourceFile(void* ptr, long long ty, struct QtQuick_PackedString sourceFile)
{
	static_cast<QSGMaterialShader*>(ptr)->setShaderSourceFile(static_cast<QOpenGLShader::ShaderTypeBit>(ty), QString::fromUtf8(sourceFile.data, sourceFile.len));
}

void QSGMaterialShader_SetShaderSourceFiles(void* ptr, long long ty, struct QtQuick_PackedString sourceFiles)
{
	static_cast<QSGMaterialShader*>(ptr)->setShaderSourceFiles(static_cast<QOpenGLShader::ShaderTypeBit>(ty), QString::fromUtf8(sourceFiles.data, sourceFiles.len).split("|", QString::SkipEmptyParts));
}

class MyQSGNode: public QSGNode
{
public:
	MyQSGNode() : QSGNode() {};
	void preprocess() { callbackQSGNode_Preprocess(this); };
	 ~MyQSGNode() { callbackQSGNode_DestroyQSGNode(this); };
	bool isSubtreeBlocked() const { return callbackQSGNode_IsSubtreeBlocked(const_cast<void*>(static_cast<const void*>(this))) != 0; };
};

void* QSGNode_NewQSGNode()
{
	return new MyQSGNode();
}

void QSGNode_AppendChildNode(void* ptr, void* node)
{
	static_cast<QSGNode*>(ptr)->appendChildNode(static_cast<QSGNode*>(node));
}

void QSGNode_InsertChildNodeAfter(void* ptr, void* node, void* after)
{
	static_cast<QSGNode*>(ptr)->insertChildNodeAfter(static_cast<QSGNode*>(node), static_cast<QSGNode*>(after));
}

void QSGNode_InsertChildNodeBefore(void* ptr, void* node, void* before)
{
	static_cast<QSGNode*>(ptr)->insertChildNodeBefore(static_cast<QSGNode*>(node), static_cast<QSGNode*>(before));
}

void QSGNode_MarkDirty(void* ptr, long long bits)
{
	static_cast<QSGNode*>(ptr)->markDirty(static_cast<QSGNode::DirtyStateBit>(bits));
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
	if (dynamic_cast<QSGRenderNode*>(static_cast<QSGNode*>(ptr))) {
		static_cast<QSGRenderNode*>(ptr)->QSGRenderNode::preprocess();
	} else if (dynamic_cast<QSGTransformNode*>(static_cast<QSGNode*>(ptr))) {
		static_cast<QSGTransformNode*>(ptr)->QSGTransformNode::preprocess();
	} else if (dynamic_cast<QSGOpacityNode*>(static_cast<QSGNode*>(ptr))) {
		static_cast<QSGOpacityNode*>(ptr)->QSGOpacityNode::preprocess();
	} else if (dynamic_cast<QSGRectangleNode*>(static_cast<QSGNode*>(ptr))) {
		static_cast<QSGRectangleNode*>(ptr)->QSGRectangleNode::preprocess();
	} else if (dynamic_cast<QSGImageNode*>(static_cast<QSGNode*>(ptr))) {
		static_cast<QSGImageNode*>(ptr)->QSGImageNode::preprocess();
	} else if (dynamic_cast<QSGGeometryNode*>(static_cast<QSGNode*>(ptr))) {
		static_cast<QSGGeometryNode*>(ptr)->QSGGeometryNode::preprocess();
	} else if (dynamic_cast<QSGClipNode*>(static_cast<QSGNode*>(ptr))) {
		static_cast<QSGClipNode*>(ptr)->QSGClipNode::preprocess();
	} else if (dynamic_cast<QSGBasicGeometryNode*>(static_cast<QSGNode*>(ptr))) {
		static_cast<QSGBasicGeometryNode*>(ptr)->QSGBasicGeometryNode::preprocess();
	} else {
		static_cast<QSGNode*>(ptr)->QSGNode::preprocess();
	}
}

void QSGNode_RemoveAllChildNodes(void* ptr)
{
	static_cast<QSGNode*>(ptr)->removeAllChildNodes();
}

void QSGNode_RemoveChildNode(void* ptr, void* node)
{
	static_cast<QSGNode*>(ptr)->removeChildNode(static_cast<QSGNode*>(node));
}

void QSGNode_SetFlag(void* ptr, long long fo, char enabled)
{
	static_cast<QSGNode*>(ptr)->setFlag(static_cast<QSGNode::Flag>(fo), enabled != 0);
}

void QSGNode_SetFlags(void* ptr, long long fo, char enabled)
{
	static_cast<QSGNode*>(ptr)->setFlags(static_cast<QSGNode::Flag>(fo), enabled != 0);
}

void QSGNode_DestroyQSGNode(void* ptr)
{
	static_cast<QSGNode*>(ptr)->~QSGNode();
}

void QSGNode_DestroyQSGNodeDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

long long QSGNode_Flags(void* ptr)
{
	return static_cast<QSGNode*>(ptr)->flags();
}

long long QSGNode_Type(void* ptr)
{
	return static_cast<QSGNode*>(ptr)->type();
}

void* QSGNode_ChildAtIndex(void* ptr, int i)
{
	return static_cast<QSGNode*>(ptr)->childAtIndex(i);
}

void* QSGNode_FirstChild(void* ptr)
{
	return static_cast<QSGNode*>(ptr)->firstChild();
}

void* QSGNode_LastChild(void* ptr)
{
	return static_cast<QSGNode*>(ptr)->lastChild();
}

void* QSGNode_NextSibling(void* ptr)
{
	return static_cast<QSGNode*>(ptr)->nextSibling();
}

void* QSGNode_Parent(void* ptr)
{
	return static_cast<QSGNode*>(ptr)->parent();
}

void* QSGNode_PreviousSibling(void* ptr)
{
	return static_cast<QSGNode*>(ptr)->previousSibling();
}

char QSGNode_IsSubtreeBlocked(void* ptr)
{
	return static_cast<QSGNode*>(ptr)->isSubtreeBlocked();
}

char QSGNode_IsSubtreeBlockedDefault(void* ptr)
{
	if (dynamic_cast<QSGRenderNode*>(static_cast<QSGNode*>(ptr))) {
		return static_cast<QSGRenderNode*>(ptr)->QSGRenderNode::isSubtreeBlocked();
	} else if (dynamic_cast<QSGTransformNode*>(static_cast<QSGNode*>(ptr))) {
		return static_cast<QSGTransformNode*>(ptr)->QSGTransformNode::isSubtreeBlocked();
	} else if (dynamic_cast<QSGOpacityNode*>(static_cast<QSGNode*>(ptr))) {
		return static_cast<QSGOpacityNode*>(ptr)->QSGOpacityNode::isSubtreeBlocked();
	} else if (dynamic_cast<QSGRectangleNode*>(static_cast<QSGNode*>(ptr))) {
		return static_cast<QSGRectangleNode*>(ptr)->QSGRectangleNode::isSubtreeBlocked();
	} else if (dynamic_cast<QSGImageNode*>(static_cast<QSGNode*>(ptr))) {
		return static_cast<QSGImageNode*>(ptr)->QSGImageNode::isSubtreeBlocked();
	} else if (dynamic_cast<QSGGeometryNode*>(static_cast<QSGNode*>(ptr))) {
		return static_cast<QSGGeometryNode*>(ptr)->QSGGeometryNode::isSubtreeBlocked();
	} else if (dynamic_cast<QSGClipNode*>(static_cast<QSGNode*>(ptr))) {
		return static_cast<QSGClipNode*>(ptr)->QSGClipNode::isSubtreeBlocked();
	} else if (dynamic_cast<QSGBasicGeometryNode*>(static_cast<QSGNode*>(ptr))) {
		return static_cast<QSGBasicGeometryNode*>(ptr)->QSGBasicGeometryNode::isSubtreeBlocked();
	} else {
		return static_cast<QSGNode*>(ptr)->QSGNode::isSubtreeBlocked();
	}
}

int QSGNode_ChildCount(void* ptr)
{
	return static_cast<QSGNode*>(ptr)->childCount();
}

class MyQSGOpacityNode: public QSGOpacityNode
{
public:
	MyQSGOpacityNode() : QSGOpacityNode() {};
	void preprocess() { callbackQSGNode_Preprocess(this); };
	bool isSubtreeBlocked() const { return callbackQSGNode_IsSubtreeBlocked(const_cast<void*>(static_cast<const void*>(this))) != 0; };
};

void* QSGOpacityNode_NewQSGOpacityNode()
{
	return new MyQSGOpacityNode();
}

void QSGOpacityNode_SetOpacity(void* ptr, double opacity)
{
	static_cast<QSGOpacityNode*>(ptr)->setOpacity(opacity);
}

void QSGOpacityNode_DestroyQSGOpacityNode(void* ptr)
{
	static_cast<QSGOpacityNode*>(ptr)->~QSGOpacityNode();
}

double QSGOpacityNode_Opacity(void* ptr)
{
	return static_cast<QSGOpacityNode*>(ptr)->opacity();
}

class MyQSGOpaqueTextureMaterial: public QSGOpaqueTextureMaterial
{
public:
	MyQSGOpaqueTextureMaterial() : QSGOpaqueTextureMaterial() {};
	QSGMaterialShader * createShader() const { return static_cast<QSGMaterialShader*>(callbackQSGOpaqueTextureMaterial_CreateShader(const_cast<void*>(static_cast<const void*>(this)))); };
	QSGMaterialType * type() const { return static_cast<QSGMaterialType*>(callbackQSGOpaqueTextureMaterial_Type(const_cast<void*>(static_cast<const void*>(this)))); };
	int compare(const QSGMaterial * other) const { return callbackQSGMaterial_Compare(const_cast<void*>(static_cast<const void*>(this)), const_cast<QSGMaterial*>(other)); };
};

void* QSGOpaqueTextureMaterial_NewQSGOpaqueTextureMaterial()
{
	return new MyQSGOpaqueTextureMaterial();
}

void QSGOpaqueTextureMaterial_SetAnisotropyLevel(void* ptr, long long level)
{
	static_cast<QSGOpaqueTextureMaterial*>(ptr)->setAnisotropyLevel(static_cast<QSGTexture::AnisotropyLevel>(level));
}

void QSGOpaqueTextureMaterial_SetFiltering(void* ptr, long long filtering)
{
	static_cast<QSGOpaqueTextureMaterial*>(ptr)->setFiltering(static_cast<QSGTexture::Filtering>(filtering));
}

void QSGOpaqueTextureMaterial_SetHorizontalWrapMode(void* ptr, long long mode)
{
	static_cast<QSGOpaqueTextureMaterial*>(ptr)->setHorizontalWrapMode(static_cast<QSGTexture::WrapMode>(mode));
}

void QSGOpaqueTextureMaterial_SetMipmapFiltering(void* ptr, long long filtering)
{
	static_cast<QSGOpaqueTextureMaterial*>(ptr)->setMipmapFiltering(static_cast<QSGTexture::Filtering>(filtering));
}

void QSGOpaqueTextureMaterial_SetTexture(void* ptr, void* texture)
{
	static_cast<QSGOpaqueTextureMaterial*>(ptr)->setTexture(static_cast<QSGTexture*>(texture));
}

void QSGOpaqueTextureMaterial_SetVerticalWrapMode(void* ptr, long long mode)
{
	static_cast<QSGOpaqueTextureMaterial*>(ptr)->setVerticalWrapMode(static_cast<QSGTexture::WrapMode>(mode));
}

void* QSGOpaqueTextureMaterial_Texture(void* ptr)
{
	return static_cast<QSGOpaqueTextureMaterial*>(ptr)->texture();
}

long long QSGOpaqueTextureMaterial_AnisotropyLevel(void* ptr)
{
	return static_cast<QSGOpaqueTextureMaterial*>(ptr)->anisotropyLevel();
}

long long QSGOpaqueTextureMaterial_Filtering(void* ptr)
{
	return static_cast<QSGOpaqueTextureMaterial*>(ptr)->filtering();
}

long long QSGOpaqueTextureMaterial_MipmapFiltering(void* ptr)
{
	return static_cast<QSGOpaqueTextureMaterial*>(ptr)->mipmapFiltering();
}

long long QSGOpaqueTextureMaterial_HorizontalWrapMode(void* ptr)
{
	return static_cast<QSGOpaqueTextureMaterial*>(ptr)->horizontalWrapMode();
}

long long QSGOpaqueTextureMaterial_VerticalWrapMode(void* ptr)
{
	return static_cast<QSGOpaqueTextureMaterial*>(ptr)->verticalWrapMode();
}

void* QSGOpaqueTextureMaterial_M_texture(void* ptr)
{
	return static_cast<QSGOpaqueTextureMaterial*>(ptr)->m_texture;
}

void QSGOpaqueTextureMaterial_SetM_texture(void* ptr, void* vqs)
{
	static_cast<QSGOpaqueTextureMaterial*>(ptr)->m_texture = static_cast<QSGTexture*>(vqs);
}

void* QSGOpaqueTextureMaterial_CreateShader(void* ptr)
{
	return static_cast<QSGOpaqueTextureMaterial*>(ptr)->createShader();
}

void* QSGOpaqueTextureMaterial_CreateShaderDefault(void* ptr)
{
	if (dynamic_cast<QSGTextureMaterial*>(static_cast<QSGOpaqueTextureMaterial*>(ptr))) {
		return static_cast<QSGTextureMaterial*>(ptr)->QSGTextureMaterial::createShader();
	} else {
		return static_cast<QSGOpaqueTextureMaterial*>(ptr)->QSGOpaqueTextureMaterial::createShader();
	}
}

void* QSGOpaqueTextureMaterial_Type(void* ptr)
{
	return static_cast<QSGOpaqueTextureMaterial*>(ptr)->type();
}

void* QSGOpaqueTextureMaterial_TypeDefault(void* ptr)
{
	if (dynamic_cast<QSGTextureMaterial*>(static_cast<QSGOpaqueTextureMaterial*>(ptr))) {
		return static_cast<QSGTextureMaterial*>(ptr)->QSGTextureMaterial::type();
	} else {
		return static_cast<QSGOpaqueTextureMaterial*>(ptr)->QSGOpaqueTextureMaterial::type();
	}
}

class MyQSGRectangleNode: public QSGRectangleNode
{
public:
	void setColor(const QColor & color) { callbackQSGRectangleNode_SetColor(this, const_cast<QColor*>(&color)); };
	void setRect(const QRectF & rect) { callbackQSGRectangleNode_SetRect(this, const_cast<QRectF*>(&rect)); };
	 ~MyQSGRectangleNode() { callbackQSGRectangleNode_DestroyQSGRectangleNode(this); };
	QColor color() const { return *static_cast<QColor*>(callbackQSGRectangleNode_Color(const_cast<void*>(static_cast<const void*>(this)))); };
	QRectF rect() const { return *static_cast<QRectF*>(callbackQSGRectangleNode_Rect(const_cast<void*>(static_cast<const void*>(this)))); };
	void preprocess() { callbackQSGNode_Preprocess(this); };
	bool isSubtreeBlocked() const { return callbackQSGNode_IsSubtreeBlocked(const_cast<void*>(static_cast<const void*>(this))) != 0; };
};

void QSGRectangleNode_SetColor(void* ptr, void* color)
{
	static_cast<QSGRectangleNode*>(ptr)->setColor(*static_cast<QColor*>(color));
}

void QSGRectangleNode_SetRect(void* ptr, void* rect)
{
	static_cast<QSGRectangleNode*>(ptr)->setRect(*static_cast<QRectF*>(rect));
}

void QSGRectangleNode_SetRect2(void* ptr, double x, double y, double w, double h)
{
	static_cast<QSGRectangleNode*>(ptr)->setRect(x, y, w, h);
}

void QSGRectangleNode_DestroyQSGRectangleNode(void* ptr)
{
	static_cast<QSGRectangleNode*>(ptr)->~QSGRectangleNode();
}

void QSGRectangleNode_DestroyQSGRectangleNodeDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QSGRectangleNode_Color(void* ptr)
{
	return new QColor(static_cast<QSGRectangleNode*>(ptr)->color());
}

void* QSGRectangleNode_Rect(void* ptr)
{
	return ({ QRectF tmpValue = static_cast<QSGRectangleNode*>(ptr)->rect(); new QRectF(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

class MyQSGRenderNode: public QSGRenderNode
{
public:
	void releaseResources() { callbackQSGRenderNode_ReleaseResources(this); };
	QRectF rect() const { return *static_cast<QRectF*>(callbackQSGRenderNode_Rect(const_cast<void*>(static_cast<const void*>(this)))); };
	RenderingFlags flags() const { return static_cast<QSGRenderNode::RenderingFlag>(callbackQSGRenderNode_Flags(const_cast<void*>(static_cast<const void*>(this)))); };
	StateFlags changedStates() const { return static_cast<QSGRenderNode::StateFlag>(callbackQSGRenderNode_ChangedStates(const_cast<void*>(static_cast<const void*>(this)))); };
	void preprocess() { callbackQSGNode_Preprocess(this); };
	bool isSubtreeBlocked() const { return callbackQSGNode_IsSubtreeBlocked(const_cast<void*>(static_cast<const void*>(this))) != 0; };
};

void QSGRenderNode_ReleaseResources(void* ptr)
{
	static_cast<QSGRenderNode*>(ptr)->releaseResources();
}

void QSGRenderNode_ReleaseResourcesDefault(void* ptr)
{
		static_cast<QSGRenderNode*>(ptr)->QSGRenderNode::releaseResources();
}

void QSGRenderNode_DestroyQSGRenderNode(void* ptr)
{
	static_cast<QSGRenderNode*>(ptr)->~QSGRenderNode();
}

void* QSGRenderNode_Rect(void* ptr)
{
	return ({ QRectF tmpValue = static_cast<QSGRenderNode*>(ptr)->rect(); new QRectF(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void* QSGRenderNode_RectDefault(void* ptr)
{
		return ({ QRectF tmpValue = static_cast<QSGRenderNode*>(ptr)->QSGRenderNode::rect(); new QRectF(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

long long QSGRenderNode_Flags(void* ptr)
{
	return static_cast<QSGRenderNode*>(ptr)->flags();
}

long long QSGRenderNode_FlagsDefault(void* ptr)
{
		return static_cast<QSGRenderNode*>(ptr)->QSGRenderNode::flags();
}

long long QSGRenderNode_ChangedStates(void* ptr)
{
	return static_cast<QSGRenderNode*>(ptr)->changedStates();
}

long long QSGRenderNode_ChangedStatesDefault(void* ptr)
{
		return static_cast<QSGRenderNode*>(ptr)->QSGRenderNode::changedStates();
}

void* QSGRenderNode_Matrix(void* ptr)
{
	return const_cast<QMatrix4x4*>(static_cast<QSGRenderNode*>(ptr)->matrix());
}

void* QSGRenderNode_ClipList(void* ptr)
{
	return const_cast<QSGClipNode*>(static_cast<QSGRenderNode*>(ptr)->clipList());
}

double QSGRenderNode_InheritedOpacity(void* ptr)
{
	return static_cast<QSGRenderNode*>(ptr)->inheritedOpacity();
}

class MyQSGRendererInterface: public QSGRendererInterface
{
public:
	 ~MyQSGRendererInterface() { callbackQSGRendererInterface_DestroyQSGRendererInterface(this); };
	GraphicsApi graphicsApi() const { return static_cast<QSGRendererInterface::GraphicsApi>(callbackQSGRendererInterface_GraphicsApi(const_cast<void*>(static_cast<const void*>(this)))); };
	ShaderCompilationTypes shaderCompilationType() const { return static_cast<QSGRendererInterface::ShaderCompilationType>(callbackQSGRendererInterface_ShaderCompilationType(const_cast<void*>(static_cast<const void*>(this)))); };
	ShaderSourceTypes shaderSourceType() const { return static_cast<QSGRendererInterface::ShaderSourceType>(callbackQSGRendererInterface_ShaderSourceType(const_cast<void*>(static_cast<const void*>(this)))); };
	ShaderType shaderType() const { return static_cast<QSGRendererInterface::ShaderType>(callbackQSGRendererInterface_ShaderType(const_cast<void*>(static_cast<const void*>(this)))); };
	void * getResource(QQuickWindow * window, QSGRendererInterface::Resource resource) const { return callbackQSGRendererInterface_GetResource(const_cast<void*>(static_cast<const void*>(this)), window, resource); };
	void * getResource(QQuickWindow * window, const char * resource) const { QtQuick_PackedString resourcePacked = { const_cast<char*>(resource), -1 };return callbackQSGRendererInterface_GetResource2(const_cast<void*>(static_cast<const void*>(this)), window, resourcePacked); };
};

void QSGRendererInterface_DestroyQSGRendererInterface(void* ptr)
{
	static_cast<QSGRendererInterface*>(ptr)->~QSGRendererInterface();
}

void QSGRendererInterface_DestroyQSGRendererInterfaceDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

long long QSGRendererInterface_GraphicsApi(void* ptr)
{
	return static_cast<QSGRendererInterface*>(ptr)->graphicsApi();
}

long long QSGRendererInterface_ShaderCompilationType(void* ptr)
{
	return static_cast<QSGRendererInterface*>(ptr)->shaderCompilationType();
}

long long QSGRendererInterface_ShaderSourceType(void* ptr)
{
	return static_cast<QSGRendererInterface*>(ptr)->shaderSourceType();
}

long long QSGRendererInterface_ShaderType(void* ptr)
{
	return static_cast<QSGRendererInterface*>(ptr)->shaderType();
}

void* QSGRendererInterface_GetResource(void* ptr, void* window, long long resource)
{
	return static_cast<QSGRendererInterface*>(ptr)->getResource(static_cast<QQuickWindow*>(window), static_cast<QSGRendererInterface::Resource>(resource));
}

void* QSGRendererInterface_GetResourceDefault(void* ptr, void* window, long long resource)
{
		return static_cast<QSGRendererInterface*>(ptr)->QSGRendererInterface::getResource(static_cast<QQuickWindow*>(window), static_cast<QSGRendererInterface::Resource>(resource));
}

void* QSGRendererInterface_GetResource2(void* ptr, void* window, char* resource)
{
	return static_cast<QSGRendererInterface*>(ptr)->getResource(static_cast<QQuickWindow*>(window), const_cast<const char*>(resource));
}

void* QSGRendererInterface_GetResource2Default(void* ptr, void* window, char* resource)
{
		return static_cast<QSGRendererInterface*>(ptr)->QSGRendererInterface::getResource(static_cast<QQuickWindow*>(window), const_cast<const char*>(resource));
}

class MyQSGTexture: public QSGTexture
{
public:
	MyQSGTexture() : QSGTexture() {QSGTexture_QSGTexture_QRegisterMetaType();};
	void bind() { callbackQSGTexture_Bind(this); };
	QRectF normalizedTextureSubRect() const { return *static_cast<QRectF*>(callbackQSGTexture_NormalizedTextureSubRect(const_cast<void*>(static_cast<const void*>(this)))); };
	QSGTexture * removedFromAtlas() const { return static_cast<QSGTexture*>(callbackQSGTexture_RemovedFromAtlas(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize textureSize() const { return *static_cast<QSize*>(callbackQSGTexture_TextureSize(const_cast<void*>(static_cast<const void*>(this)))); };
	bool hasAlphaChannel() const { return callbackQSGTexture_HasAlphaChannel(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	bool hasMipmaps() const { return callbackQSGTexture_HasMipmaps(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	bool isAtlasTexture() const { return callbackQSGTexture_IsAtlasTexture(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int textureId() const { return callbackQSGTexture_TextureId(const_cast<void*>(static_cast<const void*>(this))); };
	bool event(QEvent * e) { return callbackQSGTexture_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSGTexture_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQSGTexture_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSGTexture_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQSGTexture_CustomEvent(this, event); };
	void deleteLater() { callbackQSGTexture_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQSGTexture_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSGTexture_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtQuick_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQSGTexture_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQSGTexture_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSGTexture_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQSGTexture*)

int QSGTexture_QSGTexture_QRegisterMetaType(){qRegisterMetaType<QSGTexture*>(); return qRegisterMetaType<MyQSGTexture*>();}

void* QSGTexture_NewQSGTexture()
{
	return new MyQSGTexture();
}

void QSGTexture_Bind(void* ptr)
{
	static_cast<QSGTexture*>(ptr)->bind();
}

void QSGTexture_SetAnisotropyLevel(void* ptr, long long level)
{
	static_cast<QSGTexture*>(ptr)->setAnisotropyLevel(static_cast<QSGTexture::AnisotropyLevel>(level));
}

void QSGTexture_SetFiltering(void* ptr, long long filter)
{
	static_cast<QSGTexture*>(ptr)->setFiltering(static_cast<QSGTexture::Filtering>(filter));
}

void QSGTexture_SetHorizontalWrapMode(void* ptr, long long hwrap)
{
	static_cast<QSGTexture*>(ptr)->setHorizontalWrapMode(static_cast<QSGTexture::WrapMode>(hwrap));
}

void QSGTexture_SetMipmapFiltering(void* ptr, long long filter)
{
	static_cast<QSGTexture*>(ptr)->setMipmapFiltering(static_cast<QSGTexture::Filtering>(filter));
}

void QSGTexture_SetVerticalWrapMode(void* ptr, long long vwrap)
{
	static_cast<QSGTexture*>(ptr)->setVerticalWrapMode(static_cast<QSGTexture::WrapMode>(vwrap));
}

void QSGTexture_UpdateBindOptions(void* ptr, char force)
{
	static_cast<QSGTexture*>(ptr)->updateBindOptions(force != 0);
}

void QSGTexture_DestroyQSGTexture(void* ptr)
{
	static_cast<QSGTexture*>(ptr)->~QSGTexture();
}

void* QSGTexture_ConvertToNormalizedSourceRect(void* ptr, void* rect)
{
	return ({ QRectF tmpValue = static_cast<QSGTexture*>(ptr)->convertToNormalizedSourceRect(*static_cast<QRectF*>(rect)); new QRectF(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void* QSGTexture_NormalizedTextureSubRect(void* ptr)
{
	return ({ QRectF tmpValue = static_cast<QSGTexture*>(ptr)->normalizedTextureSubRect(); new QRectF(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void* QSGTexture_NormalizedTextureSubRectDefault(void* ptr)
{
	if (dynamic_cast<QSGDynamicTexture*>(static_cast<QObject*>(ptr))) {
		return ({ QRectF tmpValue = static_cast<QSGDynamicTexture*>(ptr)->QSGDynamicTexture::normalizedTextureSubRect(); new QRectF(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
	} else {
		return ({ QRectF tmpValue = static_cast<QSGTexture*>(ptr)->QSGTexture::normalizedTextureSubRect(); new QRectF(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
	}
}

void* QSGTexture_RemovedFromAtlas(void* ptr)
{
	return static_cast<QSGTexture*>(ptr)->removedFromAtlas();
}

void* QSGTexture_RemovedFromAtlasDefault(void* ptr)
{
	if (dynamic_cast<QSGDynamicTexture*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSGDynamicTexture*>(ptr)->QSGDynamicTexture::removedFromAtlas();
	} else {
		return static_cast<QSGTexture*>(ptr)->QSGTexture::removedFromAtlas();
	}
}

long long QSGTexture_AnisotropyLevel(void* ptr)
{
	return static_cast<QSGTexture*>(ptr)->anisotropyLevel();
}

long long QSGTexture_Filtering(void* ptr)
{
	return static_cast<QSGTexture*>(ptr)->filtering();
}

long long QSGTexture_MipmapFiltering(void* ptr)
{
	return static_cast<QSGTexture*>(ptr)->mipmapFiltering();
}

long long QSGTexture_HorizontalWrapMode(void* ptr)
{
	return static_cast<QSGTexture*>(ptr)->horizontalWrapMode();
}

long long QSGTexture_VerticalWrapMode(void* ptr)
{
	return static_cast<QSGTexture*>(ptr)->verticalWrapMode();
}

void* QSGTexture_TextureSize(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QSGTexture*>(ptr)->textureSize(); new QSize(tmpValue.width(), tmpValue.height()); });
}

char QSGTexture_HasAlphaChannel(void* ptr)
{
	return static_cast<QSGTexture*>(ptr)->hasAlphaChannel();
}

char QSGTexture_HasMipmaps(void* ptr)
{
	return static_cast<QSGTexture*>(ptr)->hasMipmaps();
}

char QSGTexture_IsAtlasTexture(void* ptr)
{
	return static_cast<QSGTexture*>(ptr)->isAtlasTexture();
}

char QSGTexture_IsAtlasTextureDefault(void* ptr)
{
	if (dynamic_cast<QSGDynamicTexture*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSGDynamicTexture*>(ptr)->QSGDynamicTexture::isAtlasTexture();
	} else {
		return static_cast<QSGTexture*>(ptr)->QSGTexture::isAtlasTexture();
	}
}

int QSGTexture_TextureId(void* ptr)
{
	return static_cast<QSGTexture*>(ptr)->textureId();
}

void* QSGTexture___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QSGTexture___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QSGTexture___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QSGTexture___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QSGTexture___findChildren_setList2(void* ptr, void* i)
{
	if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QQuickItem*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QSGTexture___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QSGTexture___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QSGTexture___findChildren_setList3(void* ptr, void* i)
{
	if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QQuickItem*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QSGTexture___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QSGTexture___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QSGTexture___findChildren_setList(void* ptr, void* i)
{
	if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QQuickItem*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QSGTexture___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QSGTexture___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QSGTexture___children_setList(void* ptr, void* i)
{
	if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QQuickItem*>(i));
	} else {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QSGTexture___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

char QSGTexture_EventDefault(void* ptr, void* e)
{
	if (dynamic_cast<QSGDynamicTexture*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSGDynamicTexture*>(ptr)->QSGDynamicTexture::event(static_cast<QEvent*>(e));
	} else {
		return static_cast<QSGTexture*>(ptr)->QSGTexture::event(static_cast<QEvent*>(e));
	}
}

char QSGTexture_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QSGDynamicTexture*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(watched))) {
			return static_cast<QSGDynamicTexture*>(ptr)->QSGDynamicTexture::eventFilter(static_cast<QQuickItem*>(watched), static_cast<QEvent*>(event));
		} else {
			return static_cast<QSGDynamicTexture*>(ptr)->QSGDynamicTexture::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
		}
	} else {
		if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(watched))) {
			return static_cast<QSGTexture*>(ptr)->QSGTexture::eventFilter(static_cast<QQuickItem*>(watched), static_cast<QEvent*>(event));
		} else {
			return static_cast<QSGTexture*>(ptr)->QSGTexture::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
		}
	}
}

void QSGTexture_ChildEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QSGDynamicTexture*>(static_cast<QObject*>(ptr))) {
		static_cast<QSGDynamicTexture*>(ptr)->QSGDynamicTexture::childEvent(static_cast<QChildEvent*>(event));
	} else {
		static_cast<QSGTexture*>(ptr)->QSGTexture::childEvent(static_cast<QChildEvent*>(event));
	}
}

void QSGTexture_ConnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QSGDynamicTexture*>(static_cast<QObject*>(ptr))) {
		static_cast<QSGDynamicTexture*>(ptr)->QSGDynamicTexture::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QSGTexture*>(ptr)->QSGTexture::connectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

void QSGTexture_CustomEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QSGDynamicTexture*>(static_cast<QObject*>(ptr))) {
		static_cast<QSGDynamicTexture*>(ptr)->QSGDynamicTexture::customEvent(static_cast<QEvent*>(event));
	} else {
		static_cast<QSGTexture*>(ptr)->QSGTexture::customEvent(static_cast<QEvent*>(event));
	}
}

void QSGTexture_DeleteLaterDefault(void* ptr)
{
	if (dynamic_cast<QSGDynamicTexture*>(static_cast<QObject*>(ptr))) {
		static_cast<QSGDynamicTexture*>(ptr)->QSGDynamicTexture::deleteLater();
	} else {
		static_cast<QSGTexture*>(ptr)->QSGTexture::deleteLater();
	}
}

void QSGTexture_DisconnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QSGDynamicTexture*>(static_cast<QObject*>(ptr))) {
		static_cast<QSGDynamicTexture*>(ptr)->QSGDynamicTexture::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QSGTexture*>(ptr)->QSGTexture::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

void QSGTexture_TimerEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QSGDynamicTexture*>(static_cast<QObject*>(ptr))) {
		static_cast<QSGDynamicTexture*>(ptr)->QSGDynamicTexture::timerEvent(static_cast<QTimerEvent*>(event));
	} else {
		static_cast<QSGTexture*>(ptr)->QSGTexture::timerEvent(static_cast<QTimerEvent*>(event));
	}
}

void* QSGTexture_MetaObjectDefault(void* ptr)
{
	if (dynamic_cast<QSGDynamicTexture*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QSGDynamicTexture*>(ptr)->QSGDynamicTexture::metaObject());
	} else {
		return const_cast<QMetaObject*>(static_cast<QSGTexture*>(ptr)->QSGTexture::metaObject());
	}
}

class MyQSGTextureMaterial: public QSGTextureMaterial
{
public:
	QSGMaterialShader * createShader() const { return static_cast<QSGMaterialShader*>(callbackQSGOpaqueTextureMaterial_CreateShader(const_cast<void*>(static_cast<const void*>(this)))); };
	QSGMaterialType * type() const { return static_cast<QSGMaterialType*>(callbackQSGOpaqueTextureMaterial_Type(const_cast<void*>(static_cast<const void*>(this)))); };
	int compare(const QSGMaterial * other) const { return callbackQSGMaterial_Compare(const_cast<void*>(static_cast<const void*>(this)), const_cast<QSGMaterial*>(other)); };
};

class MyQSGTextureProvider: public QSGTextureProvider
{
public:
	void Signal_TextureChanged() { callbackQSGTextureProvider_TextureChanged(this); };
	QSGTexture * texture() const { return static_cast<QSGTexture*>(callbackQSGTextureProvider_Texture(const_cast<void*>(static_cast<const void*>(this)))); };
	bool event(QEvent * e) { return callbackQSGTextureProvider_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSGTextureProvider_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQSGTextureProvider_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSGTextureProvider_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQSGTextureProvider_CustomEvent(this, event); };
	void deleteLater() { callbackQSGTextureProvider_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQSGTextureProvider_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSGTextureProvider_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtQuick_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQSGTextureProvider_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQSGTextureProvider_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSGTextureProvider_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQSGTextureProvider*)

int QSGTextureProvider_QSGTextureProvider_QRegisterMetaType(){qRegisterMetaType<QSGTextureProvider*>(); return qRegisterMetaType<MyQSGTextureProvider*>();}

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

void* QSGTextureProvider_Texture(void* ptr)
{
	return static_cast<QSGTextureProvider*>(ptr)->texture();
}

void* QSGTextureProvider___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QSGTextureProvider___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QSGTextureProvider___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QSGTextureProvider___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QSGTextureProvider___findChildren_setList2(void* ptr, void* i)
{
	if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QQuickItem*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QSGTextureProvider___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QSGTextureProvider___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QSGTextureProvider___findChildren_setList3(void* ptr, void* i)
{
	if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QQuickItem*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QSGTextureProvider___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QSGTextureProvider___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QSGTextureProvider___findChildren_setList(void* ptr, void* i)
{
	if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QQuickItem*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QSGTextureProvider___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QSGTextureProvider___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QSGTextureProvider___children_setList(void* ptr, void* i)
{
	if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QQuickItem*>(i));
	} else {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QSGTextureProvider___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

char QSGTextureProvider_EventDefault(void* ptr, void* e)
{
		return static_cast<QSGTextureProvider*>(ptr)->QSGTextureProvider::event(static_cast<QEvent*>(e));
}

char QSGTextureProvider_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(watched))) {
		return static_cast<QSGTextureProvider*>(ptr)->QSGTextureProvider::eventFilter(static_cast<QQuickItem*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QSGTextureProvider*>(ptr)->QSGTextureProvider::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

void QSGTextureProvider_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QSGTextureProvider*>(ptr)->QSGTextureProvider::childEvent(static_cast<QChildEvent*>(event));
}

void QSGTextureProvider_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QSGTextureProvider*>(ptr)->QSGTextureProvider::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSGTextureProvider_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QSGTextureProvider*>(ptr)->QSGTextureProvider::customEvent(static_cast<QEvent*>(event));
}

void QSGTextureProvider_DeleteLaterDefault(void* ptr)
{
		static_cast<QSGTextureProvider*>(ptr)->QSGTextureProvider::deleteLater();
}

void QSGTextureProvider_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QSGTextureProvider*>(ptr)->QSGTextureProvider::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSGTextureProvider_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QSGTextureProvider*>(ptr)->QSGTextureProvider::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QSGTextureProvider_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QSGTextureProvider*>(ptr)->QSGTextureProvider::metaObject());
}

class MyQSGTransformNode: public QSGTransformNode
{
public:
	MyQSGTransformNode() : QSGTransformNode() {};
	void preprocess() { callbackQSGNode_Preprocess(this); };
	bool isSubtreeBlocked() const { return callbackQSGNode_IsSubtreeBlocked(const_cast<void*>(static_cast<const void*>(this))) != 0; };
};

void* QSGTransformNode_NewQSGTransformNode()
{
	return new MyQSGTransformNode();
}

void QSGTransformNode_SetMatrix(void* ptr, void* matrix)
{
	static_cast<QSGTransformNode*>(ptr)->setMatrix(*static_cast<QMatrix4x4*>(matrix));
}

void QSGTransformNode_DestroyQSGTransformNode(void* ptr)
{
	static_cast<QSGTransformNode*>(ptr)->~QSGTransformNode();
}

void* QSGTransformNode_Matrix(void* ptr)
{
	return const_cast<QMatrix4x4*>(&static_cast<QSGTransformNode*>(ptr)->matrix());
}

class MyQSGVertexColorMaterial: public QSGVertexColorMaterial
{
public:
	MyQSGVertexColorMaterial() : QSGVertexColorMaterial() {};
	QSGMaterialShader * createShader() const { return static_cast<QSGMaterialShader*>(callbackQSGVertexColorMaterial_CreateShader(const_cast<void*>(static_cast<const void*>(this)))); };
	QSGMaterialType * type() const { return static_cast<QSGMaterialType*>(callbackQSGVertexColorMaterial_Type(const_cast<void*>(static_cast<const void*>(this)))); };
	int compare(const QSGMaterial * other) const { return callbackQSGMaterial_Compare(const_cast<void*>(static_cast<const void*>(this)), const_cast<QSGMaterial*>(other)); };
};

void* QSGVertexColorMaterial_NewQSGVertexColorMaterial()
{
	return new MyQSGVertexColorMaterial();
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

