#define protected public

#include "quick.h"
#include "_cgo_export.h"

#include <QAction>
#include <QActionEvent>
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
#include <QMetaObject>
#include <QMouseEvent>
#include <QMoveEvent>
#include <QObject>
#include <QOpenGLContext>
#include <QOpenGLFramebufferObject>
#include <QPaintEvent>
#include <QPainter>
#include <QPixmap>
#include <QPoint>
#include <QPointF>
#include <QQmlEngine>
#include <QQmlImageProviderBase>
#include <QQuickFramebufferObject>
#include <QQuickImageProvider>
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
#include <QSGNode>
#include <QSGOpacityNode>
#include <QSGOpaqueTextureMaterial>
#include <QSGSimpleRectNode>
#include <QSGSimpleTextureNode>
#include <QSGTexture>
#include <QSGTextureProvider>
#include <QSGTransformNode>
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

class MyQQuickFramebufferObject: public QQuickFramebufferObject {
public:
	void releaseResources() { callbackQQuickFramebufferObjectReleaseResources(this, this->objectName().toUtf8().data()); };
	void Signal_TextureFollowsItemSizeChanged(bool v) { callbackQQuickFramebufferObjectTextureFollowsItemSizeChanged(this, this->objectName().toUtf8().data(), v); };
	void classBegin() { callbackQQuickFramebufferObjectClassBegin(this, this->objectName().toUtf8().data()); };
	void componentComplete() { callbackQQuickFramebufferObjectComponentComplete(this, this->objectName().toUtf8().data()); };
	void dragEnterEvent(QDragEnterEvent * event) { callbackQQuickFramebufferObjectDragEnterEvent(this, this->objectName().toUtf8().data(), event); };
	void dragLeaveEvent(QDragLeaveEvent * event) { callbackQQuickFramebufferObjectDragLeaveEvent(this, this->objectName().toUtf8().data(), event); };
	void dragMoveEvent(QDragMoveEvent * event) { callbackQQuickFramebufferObjectDragMoveEvent(this, this->objectName().toUtf8().data(), event); };
	void dropEvent(QDropEvent * event) { callbackQQuickFramebufferObjectDropEvent(this, this->objectName().toUtf8().data(), event); };
	void focusInEvent(QFocusEvent * event) { callbackQQuickFramebufferObjectFocusInEvent(this, this->objectName().toUtf8().data(), event); };
	void focusOutEvent(QFocusEvent * event) { callbackQQuickFramebufferObjectFocusOutEvent(this, this->objectName().toUtf8().data(), event); };
	void geometryChanged(const QRectF & newGeometry, const QRectF & oldGeometry) { callbackQQuickFramebufferObjectGeometryChanged(this, this->objectName().toUtf8().data(), new QRectF(static_cast<QRectF>(newGeometry).x(), static_cast<QRectF>(newGeometry).y(), static_cast<QRectF>(newGeometry).width(), static_cast<QRectF>(newGeometry).height()), new QRectF(static_cast<QRectF>(oldGeometry).x(), static_cast<QRectF>(oldGeometry).y(), static_cast<QRectF>(oldGeometry).width(), static_cast<QRectF>(oldGeometry).height())); };
	void hoverEnterEvent(QHoverEvent * event) { callbackQQuickFramebufferObjectHoverEnterEvent(this, this->objectName().toUtf8().data(), event); };
	void hoverLeaveEvent(QHoverEvent * event) { callbackQQuickFramebufferObjectHoverLeaveEvent(this, this->objectName().toUtf8().data(), event); };
	void hoverMoveEvent(QHoverEvent * event) { callbackQQuickFramebufferObjectHoverMoveEvent(this, this->objectName().toUtf8().data(), event); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackQQuickFramebufferObjectInputMethodEvent(this, this->objectName().toUtf8().data(), event); };
	void keyPressEvent(QKeyEvent * event) { callbackQQuickFramebufferObjectKeyPressEvent(this, this->objectName().toUtf8().data(), event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQQuickFramebufferObjectKeyReleaseEvent(this, this->objectName().toUtf8().data(), event); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQQuickFramebufferObjectMouseDoubleClickEvent(this, this->objectName().toUtf8().data(), event); };
	void mouseMoveEvent(QMouseEvent * event) { callbackQQuickFramebufferObjectMouseMoveEvent(this, this->objectName().toUtf8().data(), event); };
	void mousePressEvent(QMouseEvent * event) { callbackQQuickFramebufferObjectMousePressEvent(this, this->objectName().toUtf8().data(), event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackQQuickFramebufferObjectMouseReleaseEvent(this, this->objectName().toUtf8().data(), event); };
	void mouseUngrabEvent() { callbackQQuickFramebufferObjectMouseUngrabEvent(this, this->objectName().toUtf8().data()); };
	void touchEvent(QTouchEvent * event) { callbackQQuickFramebufferObjectTouchEvent(this, this->objectName().toUtf8().data(), event); };
	void touchUngrabEvent() { callbackQQuickFramebufferObjectTouchUngrabEvent(this, this->objectName().toUtf8().data()); };
	void updatePolish() { callbackQQuickFramebufferObjectUpdatePolish(this, this->objectName().toUtf8().data()); };
	void wheelEvent(QWheelEvent * event) { callbackQQuickFramebufferObjectWheelEvent(this, this->objectName().toUtf8().data(), event); };
	void timerEvent(QTimerEvent * event) { callbackQQuickFramebufferObjectTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQQuickFramebufferObjectChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQQuickFramebufferObjectCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void QQuickFramebufferObject_SetTextureFollowsItemSize(void* ptr, int follows){
	static_cast<QQuickFramebufferObject*>(ptr)->setTextureFollowsItemSize(follows != 0);
}

int QQuickFramebufferObject_TextureFollowsItemSize(void* ptr){
	return static_cast<QQuickFramebufferObject*>(ptr)->textureFollowsItemSize();
}

int QQuickFramebufferObject_IsTextureProvider(void* ptr){
	return static_cast<QQuickFramebufferObject*>(ptr)->isTextureProvider();
}

void QQuickFramebufferObject_ReleaseResources(void* ptr){
	static_cast<MyQQuickFramebufferObject*>(ptr)->releaseResources();
}

void QQuickFramebufferObject_ReleaseResourcesDefault(void* ptr){
	static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::releaseResources();
}

void QQuickFramebufferObject_ConnectTextureFollowsItemSizeChanged(void* ptr){
	QObject::connect(static_cast<QQuickFramebufferObject*>(ptr), static_cast<void (QQuickFramebufferObject::*)(bool)>(&QQuickFramebufferObject::textureFollowsItemSizeChanged), static_cast<MyQQuickFramebufferObject*>(ptr), static_cast<void (MyQQuickFramebufferObject::*)(bool)>(&MyQQuickFramebufferObject::Signal_TextureFollowsItemSizeChanged));;
}

void QQuickFramebufferObject_DisconnectTextureFollowsItemSizeChanged(void* ptr){
	QObject::disconnect(static_cast<QQuickFramebufferObject*>(ptr), static_cast<void (QQuickFramebufferObject::*)(bool)>(&QQuickFramebufferObject::textureFollowsItemSizeChanged), static_cast<MyQQuickFramebufferObject*>(ptr), static_cast<void (MyQQuickFramebufferObject::*)(bool)>(&MyQQuickFramebufferObject::Signal_TextureFollowsItemSizeChanged));;
}

void QQuickFramebufferObject_TextureFollowsItemSizeChanged(void* ptr, int v){
	static_cast<QQuickFramebufferObject*>(ptr)->textureFollowsItemSizeChanged(v != 0);
}

void* QQuickFramebufferObject_TextureProvider(void* ptr){
	return static_cast<QQuickFramebufferObject*>(ptr)->textureProvider();
}

void QQuickFramebufferObject_ClassBegin(void* ptr){
	static_cast<MyQQuickFramebufferObject*>(ptr)->classBegin();
}

void QQuickFramebufferObject_ClassBeginDefault(void* ptr){
	static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::classBegin();
}

void QQuickFramebufferObject_ComponentComplete(void* ptr){
	static_cast<MyQQuickFramebufferObject*>(ptr)->componentComplete();
}

void QQuickFramebufferObject_ComponentCompleteDefault(void* ptr){
	static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::componentComplete();
}

void QQuickFramebufferObject_DragEnterEvent(void* ptr, void* event){
	static_cast<MyQQuickFramebufferObject*>(ptr)->dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QQuickFramebufferObject_DragEnterEventDefault(void* ptr, void* event){
	static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QQuickFramebufferObject_DragLeaveEvent(void* ptr, void* event){
	static_cast<MyQQuickFramebufferObject*>(ptr)->dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QQuickFramebufferObject_DragLeaveEventDefault(void* ptr, void* event){
	static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QQuickFramebufferObject_DragMoveEvent(void* ptr, void* event){
	static_cast<MyQQuickFramebufferObject*>(ptr)->dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QQuickFramebufferObject_DragMoveEventDefault(void* ptr, void* event){
	static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QQuickFramebufferObject_DropEvent(void* ptr, void* event){
	static_cast<MyQQuickFramebufferObject*>(ptr)->dropEvent(static_cast<QDropEvent*>(event));
}

void QQuickFramebufferObject_DropEventDefault(void* ptr, void* event){
	static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::dropEvent(static_cast<QDropEvent*>(event));
}

void QQuickFramebufferObject_FocusInEvent(void* ptr, void* event){
	static_cast<MyQQuickFramebufferObject*>(ptr)->focusInEvent(static_cast<QFocusEvent*>(event));
}

void QQuickFramebufferObject_FocusInEventDefault(void* ptr, void* event){
	static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::focusInEvent(static_cast<QFocusEvent*>(event));
}

void QQuickFramebufferObject_FocusOutEvent(void* ptr, void* event){
	static_cast<MyQQuickFramebufferObject*>(ptr)->focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QQuickFramebufferObject_FocusOutEventDefault(void* ptr, void* event){
	static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QQuickFramebufferObject_GeometryChanged(void* ptr, void* newGeometry, void* oldGeometry){
	static_cast<MyQQuickFramebufferObject*>(ptr)->geometryChanged(*static_cast<QRectF*>(newGeometry), *static_cast<QRectF*>(oldGeometry));
}

void QQuickFramebufferObject_GeometryChangedDefault(void* ptr, void* newGeometry, void* oldGeometry){
	static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::geometryChanged(*static_cast<QRectF*>(newGeometry), *static_cast<QRectF*>(oldGeometry));
}

void QQuickFramebufferObject_HoverEnterEvent(void* ptr, void* event){
	static_cast<MyQQuickFramebufferObject*>(ptr)->hoverEnterEvent(static_cast<QHoverEvent*>(event));
}

void QQuickFramebufferObject_HoverEnterEventDefault(void* ptr, void* event){
	static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::hoverEnterEvent(static_cast<QHoverEvent*>(event));
}

void QQuickFramebufferObject_HoverLeaveEvent(void* ptr, void* event){
	static_cast<MyQQuickFramebufferObject*>(ptr)->hoverLeaveEvent(static_cast<QHoverEvent*>(event));
}

void QQuickFramebufferObject_HoverLeaveEventDefault(void* ptr, void* event){
	static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::hoverLeaveEvent(static_cast<QHoverEvent*>(event));
}

void QQuickFramebufferObject_HoverMoveEvent(void* ptr, void* event){
	static_cast<MyQQuickFramebufferObject*>(ptr)->hoverMoveEvent(static_cast<QHoverEvent*>(event));
}

void QQuickFramebufferObject_HoverMoveEventDefault(void* ptr, void* event){
	static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::hoverMoveEvent(static_cast<QHoverEvent*>(event));
}

void QQuickFramebufferObject_InputMethodEvent(void* ptr, void* event){
	static_cast<MyQQuickFramebufferObject*>(ptr)->inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void QQuickFramebufferObject_InputMethodEventDefault(void* ptr, void* event){
	static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void QQuickFramebufferObject_KeyPressEvent(void* ptr, void* event){
	static_cast<MyQQuickFramebufferObject*>(ptr)->keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QQuickFramebufferObject_KeyPressEventDefault(void* ptr, void* event){
	static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QQuickFramebufferObject_KeyReleaseEvent(void* ptr, void* event){
	static_cast<MyQQuickFramebufferObject*>(ptr)->keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QQuickFramebufferObject_KeyReleaseEventDefault(void* ptr, void* event){
	static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QQuickFramebufferObject_MouseDoubleClickEvent(void* ptr, void* event){
	static_cast<MyQQuickFramebufferObject*>(ptr)->mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QQuickFramebufferObject_MouseDoubleClickEventDefault(void* ptr, void* event){
	static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QQuickFramebufferObject_MouseMoveEvent(void* ptr, void* event){
	static_cast<MyQQuickFramebufferObject*>(ptr)->mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QQuickFramebufferObject_MouseMoveEventDefault(void* ptr, void* event){
	static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QQuickFramebufferObject_MousePressEvent(void* ptr, void* event){
	static_cast<MyQQuickFramebufferObject*>(ptr)->mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QQuickFramebufferObject_MousePressEventDefault(void* ptr, void* event){
	static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QQuickFramebufferObject_MouseReleaseEvent(void* ptr, void* event){
	static_cast<MyQQuickFramebufferObject*>(ptr)->mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void QQuickFramebufferObject_MouseReleaseEventDefault(void* ptr, void* event){
	static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void QQuickFramebufferObject_MouseUngrabEvent(void* ptr){
	static_cast<MyQQuickFramebufferObject*>(ptr)->mouseUngrabEvent();
}

void QQuickFramebufferObject_MouseUngrabEventDefault(void* ptr){
	static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::mouseUngrabEvent();
}

void QQuickFramebufferObject_TouchEvent(void* ptr, void* event){
	static_cast<MyQQuickFramebufferObject*>(ptr)->touchEvent(static_cast<QTouchEvent*>(event));
}

void QQuickFramebufferObject_TouchEventDefault(void* ptr, void* event){
	static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::touchEvent(static_cast<QTouchEvent*>(event));
}

void QQuickFramebufferObject_TouchUngrabEvent(void* ptr){
	static_cast<MyQQuickFramebufferObject*>(ptr)->touchUngrabEvent();
}

void QQuickFramebufferObject_TouchUngrabEventDefault(void* ptr){
	static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::touchUngrabEvent();
}

void QQuickFramebufferObject_UpdatePolish(void* ptr){
	static_cast<MyQQuickFramebufferObject*>(ptr)->updatePolish();
}

void QQuickFramebufferObject_UpdatePolishDefault(void* ptr){
	static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::updatePolish();
}

void QQuickFramebufferObject_WheelEvent(void* ptr, void* event){
	static_cast<MyQQuickFramebufferObject*>(ptr)->wheelEvent(static_cast<QWheelEvent*>(event));
}

void QQuickFramebufferObject_WheelEventDefault(void* ptr, void* event){
	static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::wheelEvent(static_cast<QWheelEvent*>(event));
}

void QQuickFramebufferObject_TimerEvent(void* ptr, void* event){
	static_cast<MyQQuickFramebufferObject*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QQuickFramebufferObject_TimerEventDefault(void* ptr, void* event){
	static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::timerEvent(static_cast<QTimerEvent*>(event));
}

void QQuickFramebufferObject_ChildEvent(void* ptr, void* event){
	static_cast<MyQQuickFramebufferObject*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QQuickFramebufferObject_ChildEventDefault(void* ptr, void* event){
	static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::childEvent(static_cast<QChildEvent*>(event));
}

void QQuickFramebufferObject_CustomEvent(void* ptr, void* event){
	static_cast<MyQQuickFramebufferObject*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QQuickFramebufferObject_CustomEventDefault(void* ptr, void* event){
	static_cast<QQuickFramebufferObject*>(ptr)->QQuickFramebufferObject::customEvent(static_cast<QEvent*>(event));
}

class MyQQuickImageProvider: public QQuickImageProvider {
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	MyQQuickImageProvider(ImageType type, Flags flags) : QQuickImageProvider(type, flags) {};
};

void* QQuickImageProvider_NewQQuickImageProvider(int ty, int flags){
	return new MyQQuickImageProvider(static_cast<QQmlImageProviderBase::ImageType>(ty), static_cast<QQmlImageProviderBase::Flag>(flags));
}

int QQuickImageProvider_Flags(void* ptr){
	return static_cast<QQuickImageProvider*>(ptr)->flags();
}

int QQuickImageProvider_ImageType(void* ptr){
	return static_cast<QQuickImageProvider*>(ptr)->imageType();
}

void* QQuickImageProvider_RequestImage(void* ptr, char* id, void* size, void* requestedSize){
	return new QImage(static_cast<QQuickImageProvider*>(ptr)->requestImage(QString(id), static_cast<QSize*>(size), *static_cast<QSize*>(requestedSize)));
}

void* QQuickImageProvider_RequestPixmap(void* ptr, char* id, void* size, void* requestedSize){
	return new QPixmap(static_cast<QQuickImageProvider*>(ptr)->requestPixmap(QString(id), static_cast<QSize*>(size), *static_cast<QSize*>(requestedSize)));
}

void QQuickImageProvider_DestroyQQuickImageProvider(void* ptr){
	static_cast<QQuickImageProvider*>(ptr)->~QQuickImageProvider();
}

char* QQuickImageProvider_ObjectNameAbs(void* ptr){
	if (dynamic_cast<MyQQuickImageProvider*>(static_cast<QQuickImageProvider*>(ptr))) {
		return static_cast<MyQQuickImageProvider*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QQuickImageProvider_BASE").toUtf8().data();
}

void QQuickImageProvider_SetObjectNameAbs(void* ptr, char* name){
	if (dynamic_cast<MyQQuickImageProvider*>(static_cast<QQuickImageProvider*>(ptr))) {
		static_cast<MyQQuickImageProvider*>(ptr)->setObjectNameAbs(QString(name));
	}
}

class MyQQuickItem: public QQuickItem {
public:
	MyQQuickItem(QQuickItem *parent) : QQuickItem(parent) {};
	void classBegin() { callbackQQuickItemClassBegin(this, this->objectName().toUtf8().data()); };
	void componentComplete() { callbackQQuickItemComponentComplete(this, this->objectName().toUtf8().data()); };
	void dragEnterEvent(QDragEnterEvent * event) { callbackQQuickItemDragEnterEvent(this, this->objectName().toUtf8().data(), event); };
	void dragLeaveEvent(QDragLeaveEvent * event) { callbackQQuickItemDragLeaveEvent(this, this->objectName().toUtf8().data(), event); };
	void dragMoveEvent(QDragMoveEvent * event) { callbackQQuickItemDragMoveEvent(this, this->objectName().toUtf8().data(), event); };
	void dropEvent(QDropEvent * event) { callbackQQuickItemDropEvent(this, this->objectName().toUtf8().data(), event); };
	void focusInEvent(QFocusEvent * event) { callbackQQuickItemFocusInEvent(this, this->objectName().toUtf8().data(), event); };
	void focusOutEvent(QFocusEvent * event) { callbackQQuickItemFocusOutEvent(this, this->objectName().toUtf8().data(), event); };
	void geometryChanged(const QRectF & newGeometry, const QRectF & oldGeometry) { callbackQQuickItemGeometryChanged(this, this->objectName().toUtf8().data(), new QRectF(static_cast<QRectF>(newGeometry).x(), static_cast<QRectF>(newGeometry).y(), static_cast<QRectF>(newGeometry).width(), static_cast<QRectF>(newGeometry).height()), new QRectF(static_cast<QRectF>(oldGeometry).x(), static_cast<QRectF>(oldGeometry).y(), static_cast<QRectF>(oldGeometry).width(), static_cast<QRectF>(oldGeometry).height())); };
	void hoverEnterEvent(QHoverEvent * event) { callbackQQuickItemHoverEnterEvent(this, this->objectName().toUtf8().data(), event); };
	void hoverLeaveEvent(QHoverEvent * event) { callbackQQuickItemHoverLeaveEvent(this, this->objectName().toUtf8().data(), event); };
	void hoverMoveEvent(QHoverEvent * event) { callbackQQuickItemHoverMoveEvent(this, this->objectName().toUtf8().data(), event); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackQQuickItemInputMethodEvent(this, this->objectName().toUtf8().data(), event); };
	void keyPressEvent(QKeyEvent * event) { callbackQQuickItemKeyPressEvent(this, this->objectName().toUtf8().data(), event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQQuickItemKeyReleaseEvent(this, this->objectName().toUtf8().data(), event); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQQuickItemMouseDoubleClickEvent(this, this->objectName().toUtf8().data(), event); };
	void mouseMoveEvent(QMouseEvent * event) { callbackQQuickItemMouseMoveEvent(this, this->objectName().toUtf8().data(), event); };
	void mousePressEvent(QMouseEvent * event) { callbackQQuickItemMousePressEvent(this, this->objectName().toUtf8().data(), event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackQQuickItemMouseReleaseEvent(this, this->objectName().toUtf8().data(), event); };
	void mouseUngrabEvent() { callbackQQuickItemMouseUngrabEvent(this, this->objectName().toUtf8().data()); };
	void releaseResources() { callbackQQuickItemReleaseResources(this, this->objectName().toUtf8().data()); };
	void touchEvent(QTouchEvent * event) { callbackQQuickItemTouchEvent(this, this->objectName().toUtf8().data(), event); };
	void touchUngrabEvent() { callbackQQuickItemTouchUngrabEvent(this, this->objectName().toUtf8().data()); };
	void updatePolish() { callbackQQuickItemUpdatePolish(this, this->objectName().toUtf8().data()); };
	void wheelEvent(QWheelEvent * event) { callbackQQuickItemWheelEvent(this, this->objectName().toUtf8().data(), event); };
	void Signal_WindowChanged(QQuickWindow * window) { callbackQQuickItemWindowChanged(this, this->objectName().toUtf8().data(), window); };
	void timerEvent(QTimerEvent * event) { callbackQQuickItemTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQQuickItemChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQQuickItemCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void* QQuickItem_NewQQuickItem(void* parent){
	return new MyQQuickItem(static_cast<QQuickItem*>(parent));
}

int QQuickItem_ActiveFocusOnTab(void* ptr){
	return static_cast<QQuickItem*>(ptr)->activeFocusOnTab();
}

int QQuickItem_Antialiasing(void* ptr){
	return static_cast<QQuickItem*>(ptr)->antialiasing();
}

double QQuickItem_BaselineOffset(void* ptr){
	return static_cast<double>(static_cast<QQuickItem*>(ptr)->baselineOffset());
}

void* QQuickItem_ChildrenRect(void* ptr){
	return new QRectF(static_cast<QRectF>(static_cast<QQuickItem*>(ptr)->childrenRect()).x(), static_cast<QRectF>(static_cast<QQuickItem*>(ptr)->childrenRect()).y(), static_cast<QRectF>(static_cast<QQuickItem*>(ptr)->childrenRect()).width(), static_cast<QRectF>(static_cast<QQuickItem*>(ptr)->childrenRect()).height());
}

int QQuickItem_Clip(void* ptr){
	return static_cast<QQuickItem*>(ptr)->clip();
}

int QQuickItem_HasActiveFocus(void* ptr){
	return static_cast<QQuickItem*>(ptr)->hasActiveFocus();
}

int QQuickItem_HasFocus(void* ptr){
	return static_cast<QQuickItem*>(ptr)->hasFocus();
}

double QQuickItem_Height(void* ptr){
	return static_cast<double>(static_cast<QQuickItem*>(ptr)->height());
}

double QQuickItem_ImplicitHeight(void* ptr){
	return static_cast<double>(static_cast<QQuickItem*>(ptr)->implicitHeight());
}

int QQuickItem_IsEnabled(void* ptr){
	return static_cast<QQuickItem*>(ptr)->isEnabled();
}

int QQuickItem_IsTextureProvider(void* ptr){
	return static_cast<QQuickItem*>(ptr)->isTextureProvider();
}

int QQuickItem_IsVisible(void* ptr){
	return static_cast<QQuickItem*>(ptr)->isVisible();
}

double QQuickItem_Opacity(void* ptr){
	return static_cast<double>(static_cast<QQuickItem*>(ptr)->opacity());
}

void* QQuickItem_ParentItem(void* ptr){
	return static_cast<QQuickItem*>(ptr)->parentItem();
}

void QQuickItem_ResetAntialiasing(void* ptr){
	static_cast<QQuickItem*>(ptr)->resetAntialiasing();
}

void QQuickItem_ResetHeight(void* ptr){
	static_cast<QQuickItem*>(ptr)->resetHeight();
}

void QQuickItem_ResetWidth(void* ptr){
	static_cast<QQuickItem*>(ptr)->resetWidth();
}

double QQuickItem_Rotation(void* ptr){
	return static_cast<double>(static_cast<QQuickItem*>(ptr)->rotation());
}

double QQuickItem_Scale(void* ptr){
	return static_cast<double>(static_cast<QQuickItem*>(ptr)->scale());
}

void QQuickItem_SetActiveFocusOnTab(void* ptr, int v){
	static_cast<QQuickItem*>(ptr)->setActiveFocusOnTab(v != 0);
}

void QQuickItem_SetAntialiasing(void* ptr, int v){
	static_cast<QQuickItem*>(ptr)->setAntialiasing(v != 0);
}

void QQuickItem_SetBaselineOffset(void* ptr, double v){
	static_cast<QQuickItem*>(ptr)->setBaselineOffset(static_cast<double>(v));
}

void QQuickItem_SetClip(void* ptr, int v){
	static_cast<QQuickItem*>(ptr)->setClip(v != 0);
}

void QQuickItem_SetEnabled(void* ptr, int v){
	static_cast<QQuickItem*>(ptr)->setEnabled(v != 0);
}

void QQuickItem_SetFocus(void* ptr, int v){
	static_cast<QQuickItem*>(ptr)->setFocus(v != 0);
}

void QQuickItem_SetFocus2(void* ptr, int focus, int reason){
	static_cast<QQuickItem*>(ptr)->setFocus(focus != 0, static_cast<Qt::FocusReason>(reason));
}

void QQuickItem_SetHeight(void* ptr, double v){
	static_cast<QQuickItem*>(ptr)->setHeight(static_cast<double>(v));
}

void QQuickItem_SetImplicitHeight(void* ptr, double v){
	static_cast<QQuickItem*>(ptr)->setImplicitHeight(static_cast<double>(v));
}

void QQuickItem_SetImplicitWidth(void* ptr, double v){
	static_cast<QQuickItem*>(ptr)->setImplicitWidth(static_cast<double>(v));
}

void QQuickItem_SetOpacity(void* ptr, double v){
	static_cast<QQuickItem*>(ptr)->setOpacity(static_cast<double>(v));
}

void QQuickItem_SetParentItem(void* ptr, void* parent){
	static_cast<QQuickItem*>(ptr)->setParentItem(static_cast<QQuickItem*>(parent));
}

void QQuickItem_SetRotation(void* ptr, double v){
	static_cast<QQuickItem*>(ptr)->setRotation(static_cast<double>(v));
}

void QQuickItem_SetScale(void* ptr, double v){
	static_cast<QQuickItem*>(ptr)->setScale(static_cast<double>(v));
}

void QQuickItem_SetSmooth(void* ptr, int v){
	static_cast<QQuickItem*>(ptr)->setSmooth(v != 0);
}

void QQuickItem_SetState(void* ptr, char* v){
	static_cast<QQuickItem*>(ptr)->setState(QString(v));
}

void QQuickItem_SetTransformOrigin(void* ptr, int v){
	static_cast<QQuickItem*>(ptr)->setTransformOrigin(static_cast<QQuickItem::TransformOrigin>(v));
}

void QQuickItem_SetVisible(void* ptr, int v){
	static_cast<QQuickItem*>(ptr)->setVisible(v != 0);
}

void QQuickItem_SetWidth(void* ptr, double v){
	static_cast<QQuickItem*>(ptr)->setWidth(static_cast<double>(v));
}

void QQuickItem_SetX(void* ptr, double v){
	static_cast<QQuickItem*>(ptr)->setX(static_cast<double>(v));
}

void QQuickItem_SetY(void* ptr, double v){
	static_cast<QQuickItem*>(ptr)->setY(static_cast<double>(v));
}

void QQuickItem_SetZ(void* ptr, double v){
	static_cast<QQuickItem*>(ptr)->setZ(static_cast<double>(v));
}

int QQuickItem_Smooth(void* ptr){
	return static_cast<QQuickItem*>(ptr)->smooth();
}

char* QQuickItem_State(void* ptr){
	return static_cast<QQuickItem*>(ptr)->state().toUtf8().data();
}

void* QQuickItem_TextureProvider(void* ptr){
	return static_cast<QQuickItem*>(ptr)->textureProvider();
}

int QQuickItem_TransformOrigin(void* ptr){
	return static_cast<QQuickItem*>(ptr)->transformOrigin();
}

double QQuickItem_Width(void* ptr){
	return static_cast<double>(static_cast<QQuickItem*>(ptr)->width());
}

double QQuickItem_X(void* ptr){
	return static_cast<double>(static_cast<QQuickItem*>(ptr)->x());
}

double QQuickItem_Y(void* ptr){
	return static_cast<double>(static_cast<QQuickItem*>(ptr)->y());
}

double QQuickItem_Z(void* ptr){
	return static_cast<double>(static_cast<QQuickItem*>(ptr)->z());
}

int QQuickItem_AcceptHoverEvents(void* ptr){
	return static_cast<QQuickItem*>(ptr)->acceptHoverEvents();
}

int QQuickItem_AcceptedMouseButtons(void* ptr){
	return static_cast<QQuickItem*>(ptr)->acceptedMouseButtons();
}

void* QQuickItem_ChildAt(void* ptr, double x, double y){
	return static_cast<QQuickItem*>(ptr)->childAt(static_cast<double>(x), static_cast<double>(y));
}

int QQuickItem_ChildMouseEventFilter(void* ptr, void* item, void* event){
	return static_cast<QQuickItem*>(ptr)->childMouseEventFilter(static_cast<QQuickItem*>(item), static_cast<QEvent*>(event));
}

void QQuickItem_ClassBegin(void* ptr){
	static_cast<MyQQuickItem*>(ptr)->classBegin();
}

void QQuickItem_ClassBeginDefault(void* ptr){
	static_cast<QQuickItem*>(ptr)->QQuickItem::classBegin();
}

void QQuickItem_ComponentComplete(void* ptr){
	static_cast<MyQQuickItem*>(ptr)->componentComplete();
}

void QQuickItem_ComponentCompleteDefault(void* ptr){
	static_cast<QQuickItem*>(ptr)->QQuickItem::componentComplete();
}

int QQuickItem_Contains(void* ptr, void* point){
	return static_cast<QQuickItem*>(ptr)->contains(*static_cast<QPointF*>(point));
}

void* QQuickItem_Cursor(void* ptr){
	return new QCursor(static_cast<QQuickItem*>(ptr)->cursor());
}

void QQuickItem_DragEnterEvent(void* ptr, void* event){
	static_cast<MyQQuickItem*>(ptr)->dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QQuickItem_DragEnterEventDefault(void* ptr, void* event){
	static_cast<QQuickItem*>(ptr)->QQuickItem::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QQuickItem_DragLeaveEvent(void* ptr, void* event){
	static_cast<MyQQuickItem*>(ptr)->dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QQuickItem_DragLeaveEventDefault(void* ptr, void* event){
	static_cast<QQuickItem*>(ptr)->QQuickItem::dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QQuickItem_DragMoveEvent(void* ptr, void* event){
	static_cast<MyQQuickItem*>(ptr)->dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QQuickItem_DragMoveEventDefault(void* ptr, void* event){
	static_cast<QQuickItem*>(ptr)->QQuickItem::dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QQuickItem_DropEvent(void* ptr, void* event){
	static_cast<MyQQuickItem*>(ptr)->dropEvent(static_cast<QDropEvent*>(event));
}

void QQuickItem_DropEventDefault(void* ptr, void* event){
	static_cast<QQuickItem*>(ptr)->QQuickItem::dropEvent(static_cast<QDropEvent*>(event));
}

int QQuickItem_Event(void* ptr, void* ev){
	return static_cast<QQuickItem*>(ptr)->event(static_cast<QEvent*>(ev));
}

int QQuickItem_FiltersChildMouseEvents(void* ptr){
	return static_cast<QQuickItem*>(ptr)->filtersChildMouseEvents();
}

int QQuickItem_Flags(void* ptr){
	return static_cast<QQuickItem*>(ptr)->flags();
}

void QQuickItem_FocusInEvent(void* ptr, void* event){
	static_cast<MyQQuickItem*>(ptr)->focusInEvent(static_cast<QFocusEvent*>(event));
}

void QQuickItem_FocusInEventDefault(void* ptr, void* event){
	static_cast<QQuickItem*>(ptr)->QQuickItem::focusInEvent(static_cast<QFocusEvent*>(event));
}

void QQuickItem_FocusOutEvent(void* ptr, void* event){
	static_cast<MyQQuickItem*>(ptr)->focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QQuickItem_FocusOutEventDefault(void* ptr, void* event){
	static_cast<QQuickItem*>(ptr)->QQuickItem::focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QQuickItem_ForceActiveFocus(void* ptr){
	static_cast<QQuickItem*>(ptr)->forceActiveFocus();
}

void QQuickItem_ForceActiveFocus2(void* ptr, int reason){
	static_cast<QQuickItem*>(ptr)->forceActiveFocus(static_cast<Qt::FocusReason>(reason));
}

void QQuickItem_GeometryChanged(void* ptr, void* newGeometry, void* oldGeometry){
	static_cast<MyQQuickItem*>(ptr)->geometryChanged(*static_cast<QRectF*>(newGeometry), *static_cast<QRectF*>(oldGeometry));
}

void QQuickItem_GeometryChangedDefault(void* ptr, void* newGeometry, void* oldGeometry){
	static_cast<QQuickItem*>(ptr)->QQuickItem::geometryChanged(*static_cast<QRectF*>(newGeometry), *static_cast<QRectF*>(oldGeometry));
}

void QQuickItem_GrabMouse(void* ptr){
	static_cast<QQuickItem*>(ptr)->grabMouse();
}

void QQuickItem_HoverEnterEvent(void* ptr, void* event){
	static_cast<MyQQuickItem*>(ptr)->hoverEnterEvent(static_cast<QHoverEvent*>(event));
}

void QQuickItem_HoverEnterEventDefault(void* ptr, void* event){
	static_cast<QQuickItem*>(ptr)->QQuickItem::hoverEnterEvent(static_cast<QHoverEvent*>(event));
}

void QQuickItem_HoverLeaveEvent(void* ptr, void* event){
	static_cast<MyQQuickItem*>(ptr)->hoverLeaveEvent(static_cast<QHoverEvent*>(event));
}

void QQuickItem_HoverLeaveEventDefault(void* ptr, void* event){
	static_cast<QQuickItem*>(ptr)->QQuickItem::hoverLeaveEvent(static_cast<QHoverEvent*>(event));
}

void QQuickItem_HoverMoveEvent(void* ptr, void* event){
	static_cast<MyQQuickItem*>(ptr)->hoverMoveEvent(static_cast<QHoverEvent*>(event));
}

void QQuickItem_HoverMoveEventDefault(void* ptr, void* event){
	static_cast<QQuickItem*>(ptr)->QQuickItem::hoverMoveEvent(static_cast<QHoverEvent*>(event));
}

double QQuickItem_ImplicitWidth(void* ptr){
	return static_cast<double>(static_cast<QQuickItem*>(ptr)->implicitWidth());
}

void QQuickItem_InputMethodEvent(void* ptr, void* event){
	static_cast<MyQQuickItem*>(ptr)->inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void QQuickItem_InputMethodEventDefault(void* ptr, void* event){
	static_cast<QQuickItem*>(ptr)->QQuickItem::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void* QQuickItem_InputMethodQuery(void* ptr, int query){
	return new QVariant(static_cast<QQuickItem*>(ptr)->inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

int QQuickItem_IsFocusScope(void* ptr){
	return static_cast<QQuickItem*>(ptr)->isFocusScope();
}

int QQuickItem_KeepMouseGrab(void* ptr){
	return static_cast<QQuickItem*>(ptr)->keepMouseGrab();
}

int QQuickItem_KeepTouchGrab(void* ptr){
	return static_cast<QQuickItem*>(ptr)->keepTouchGrab();
}

void QQuickItem_KeyPressEvent(void* ptr, void* event){
	static_cast<MyQQuickItem*>(ptr)->keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QQuickItem_KeyPressEventDefault(void* ptr, void* event){
	static_cast<QQuickItem*>(ptr)->QQuickItem::keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QQuickItem_KeyReleaseEvent(void* ptr, void* event){
	static_cast<MyQQuickItem*>(ptr)->keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QQuickItem_KeyReleaseEventDefault(void* ptr, void* event){
	static_cast<QQuickItem*>(ptr)->QQuickItem::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void* QQuickItem_MapFromItem(void* ptr, void* item, void* point){
	return new QPointF(static_cast<QPointF>(static_cast<QQuickItem*>(ptr)->mapFromItem(static_cast<QQuickItem*>(item), *static_cast<QPointF*>(point))).x(), static_cast<QPointF>(static_cast<QQuickItem*>(ptr)->mapFromItem(static_cast<QQuickItem*>(item), *static_cast<QPointF*>(point))).y());
}

void* QQuickItem_MapFromScene(void* ptr, void* point){
	return new QPointF(static_cast<QPointF>(static_cast<QQuickItem*>(ptr)->mapFromScene(*static_cast<QPointF*>(point))).x(), static_cast<QPointF>(static_cast<QQuickItem*>(ptr)->mapFromScene(*static_cast<QPointF*>(point))).y());
}

void* QQuickItem_MapRectFromItem(void* ptr, void* item, void* rect){
	return new QRectF(static_cast<QRectF>(static_cast<QQuickItem*>(ptr)->mapRectFromItem(static_cast<QQuickItem*>(item), *static_cast<QRectF*>(rect))).x(), static_cast<QRectF>(static_cast<QQuickItem*>(ptr)->mapRectFromItem(static_cast<QQuickItem*>(item), *static_cast<QRectF*>(rect))).y(), static_cast<QRectF>(static_cast<QQuickItem*>(ptr)->mapRectFromItem(static_cast<QQuickItem*>(item), *static_cast<QRectF*>(rect))).width(), static_cast<QRectF>(static_cast<QQuickItem*>(ptr)->mapRectFromItem(static_cast<QQuickItem*>(item), *static_cast<QRectF*>(rect))).height());
}

void* QQuickItem_MapRectFromScene(void* ptr, void* rect){
	return new QRectF(static_cast<QRectF>(static_cast<QQuickItem*>(ptr)->mapRectFromScene(*static_cast<QRectF*>(rect))).x(), static_cast<QRectF>(static_cast<QQuickItem*>(ptr)->mapRectFromScene(*static_cast<QRectF*>(rect))).y(), static_cast<QRectF>(static_cast<QQuickItem*>(ptr)->mapRectFromScene(*static_cast<QRectF*>(rect))).width(), static_cast<QRectF>(static_cast<QQuickItem*>(ptr)->mapRectFromScene(*static_cast<QRectF*>(rect))).height());
}

void* QQuickItem_MapRectToItem(void* ptr, void* item, void* rect){
	return new QRectF(static_cast<QRectF>(static_cast<QQuickItem*>(ptr)->mapRectToItem(static_cast<QQuickItem*>(item), *static_cast<QRectF*>(rect))).x(), static_cast<QRectF>(static_cast<QQuickItem*>(ptr)->mapRectToItem(static_cast<QQuickItem*>(item), *static_cast<QRectF*>(rect))).y(), static_cast<QRectF>(static_cast<QQuickItem*>(ptr)->mapRectToItem(static_cast<QQuickItem*>(item), *static_cast<QRectF*>(rect))).width(), static_cast<QRectF>(static_cast<QQuickItem*>(ptr)->mapRectToItem(static_cast<QQuickItem*>(item), *static_cast<QRectF*>(rect))).height());
}

void* QQuickItem_MapRectToScene(void* ptr, void* rect){
	return new QRectF(static_cast<QRectF>(static_cast<QQuickItem*>(ptr)->mapRectToScene(*static_cast<QRectF*>(rect))).x(), static_cast<QRectF>(static_cast<QQuickItem*>(ptr)->mapRectToScene(*static_cast<QRectF*>(rect))).y(), static_cast<QRectF>(static_cast<QQuickItem*>(ptr)->mapRectToScene(*static_cast<QRectF*>(rect))).width(), static_cast<QRectF>(static_cast<QQuickItem*>(ptr)->mapRectToScene(*static_cast<QRectF*>(rect))).height());
}

void* QQuickItem_MapToItem(void* ptr, void* item, void* point){
	return new QPointF(static_cast<QPointF>(static_cast<QQuickItem*>(ptr)->mapToItem(static_cast<QQuickItem*>(item), *static_cast<QPointF*>(point))).x(), static_cast<QPointF>(static_cast<QQuickItem*>(ptr)->mapToItem(static_cast<QQuickItem*>(item), *static_cast<QPointF*>(point))).y());
}

void* QQuickItem_MapToScene(void* ptr, void* point){
	return new QPointF(static_cast<QPointF>(static_cast<QQuickItem*>(ptr)->mapToScene(*static_cast<QPointF*>(point))).x(), static_cast<QPointF>(static_cast<QQuickItem*>(ptr)->mapToScene(*static_cast<QPointF*>(point))).y());
}

void QQuickItem_MouseDoubleClickEvent(void* ptr, void* event){
	static_cast<MyQQuickItem*>(ptr)->mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QQuickItem_MouseDoubleClickEventDefault(void* ptr, void* event){
	static_cast<QQuickItem*>(ptr)->QQuickItem::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QQuickItem_MouseMoveEvent(void* ptr, void* event){
	static_cast<MyQQuickItem*>(ptr)->mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QQuickItem_MouseMoveEventDefault(void* ptr, void* event){
	static_cast<QQuickItem*>(ptr)->QQuickItem::mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QQuickItem_MousePressEvent(void* ptr, void* event){
	static_cast<MyQQuickItem*>(ptr)->mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QQuickItem_MousePressEventDefault(void* ptr, void* event){
	static_cast<QQuickItem*>(ptr)->QQuickItem::mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QQuickItem_MouseReleaseEvent(void* ptr, void* event){
	static_cast<MyQQuickItem*>(ptr)->mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void QQuickItem_MouseReleaseEventDefault(void* ptr, void* event){
	static_cast<QQuickItem*>(ptr)->QQuickItem::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void QQuickItem_MouseUngrabEvent(void* ptr){
	static_cast<MyQQuickItem*>(ptr)->mouseUngrabEvent();
}

void QQuickItem_MouseUngrabEventDefault(void* ptr){
	static_cast<QQuickItem*>(ptr)->QQuickItem::mouseUngrabEvent();
}

void* QQuickItem_NextItemInFocusChain(void* ptr, int forward){
	return static_cast<QQuickItem*>(ptr)->nextItemInFocusChain(forward != 0);
}

void QQuickItem_Polish(void* ptr){
	static_cast<QQuickItem*>(ptr)->polish();
}

void QQuickItem_ReleaseResources(void* ptr){
	static_cast<MyQQuickItem*>(ptr)->releaseResources();
}

void QQuickItem_ReleaseResourcesDefault(void* ptr){
	static_cast<QQuickItem*>(ptr)->QQuickItem::releaseResources();
}

void* QQuickItem_ScopedFocusItem(void* ptr){
	return static_cast<QQuickItem*>(ptr)->scopedFocusItem();
}

void QQuickItem_SetAcceptHoverEvents(void* ptr, int enabled){
	static_cast<QQuickItem*>(ptr)->setAcceptHoverEvents(enabled != 0);
}

void QQuickItem_SetAcceptedMouseButtons(void* ptr, int buttons){
	static_cast<QQuickItem*>(ptr)->setAcceptedMouseButtons(static_cast<Qt::MouseButton>(buttons));
}

void QQuickItem_SetCursor(void* ptr, void* cursor){
	static_cast<QQuickItem*>(ptr)->setCursor(*static_cast<QCursor*>(cursor));
}

void QQuickItem_SetFiltersChildMouseEvents(void* ptr, int filter){
	static_cast<QQuickItem*>(ptr)->setFiltersChildMouseEvents(filter != 0);
}

void QQuickItem_SetFlag(void* ptr, int flag, int enabled){
	static_cast<QQuickItem*>(ptr)->setFlag(static_cast<QQuickItem::Flag>(flag), enabled != 0);
}

void QQuickItem_SetFlags(void* ptr, int flags){
	static_cast<QQuickItem*>(ptr)->setFlags(static_cast<QQuickItem::Flag>(flags));
}

void QQuickItem_SetKeepMouseGrab(void* ptr, int keep){
	static_cast<QQuickItem*>(ptr)->setKeepMouseGrab(keep != 0);
}

void QQuickItem_SetKeepTouchGrab(void* ptr, int keep){
	static_cast<QQuickItem*>(ptr)->setKeepTouchGrab(keep != 0);
}

void QQuickItem_StackAfter(void* ptr, void* sibling){
	static_cast<QQuickItem*>(ptr)->stackAfter(static_cast<QQuickItem*>(sibling));
}

void QQuickItem_StackBefore(void* ptr, void* sibling){
	static_cast<QQuickItem*>(ptr)->stackBefore(static_cast<QQuickItem*>(sibling));
}

void QQuickItem_TouchEvent(void* ptr, void* event){
	static_cast<MyQQuickItem*>(ptr)->touchEvent(static_cast<QTouchEvent*>(event));
}

void QQuickItem_TouchEventDefault(void* ptr, void* event){
	static_cast<QQuickItem*>(ptr)->QQuickItem::touchEvent(static_cast<QTouchEvent*>(event));
}

void QQuickItem_TouchUngrabEvent(void* ptr){
	static_cast<MyQQuickItem*>(ptr)->touchUngrabEvent();
}

void QQuickItem_TouchUngrabEventDefault(void* ptr){
	static_cast<QQuickItem*>(ptr)->QQuickItem::touchUngrabEvent();
}

void QQuickItem_UngrabMouse(void* ptr){
	static_cast<QQuickItem*>(ptr)->ungrabMouse();
}

void QQuickItem_UngrabTouchPoints(void* ptr){
	static_cast<QQuickItem*>(ptr)->ungrabTouchPoints();
}

void QQuickItem_UnsetCursor(void* ptr){
	static_cast<QQuickItem*>(ptr)->unsetCursor();
}

void QQuickItem_Update(void* ptr){
	QMetaObject::invokeMethod(static_cast<QQuickItem*>(ptr), "update");
}

void QQuickItem_UpdatePolish(void* ptr){
	static_cast<MyQQuickItem*>(ptr)->updatePolish();
}

void QQuickItem_UpdatePolishDefault(void* ptr){
	static_cast<QQuickItem*>(ptr)->QQuickItem::updatePolish();
}

void QQuickItem_WheelEvent(void* ptr, void* event){
	static_cast<MyQQuickItem*>(ptr)->wheelEvent(static_cast<QWheelEvent*>(event));
}

void QQuickItem_WheelEventDefault(void* ptr, void* event){
	static_cast<QQuickItem*>(ptr)->QQuickItem::wheelEvent(static_cast<QWheelEvent*>(event));
}

void* QQuickItem_Window(void* ptr){
	return static_cast<QQuickItem*>(ptr)->window();
}

void QQuickItem_ConnectWindowChanged(void* ptr){
	QObject::connect(static_cast<QQuickItem*>(ptr), static_cast<void (QQuickItem::*)(QQuickWindow *)>(&QQuickItem::windowChanged), static_cast<MyQQuickItem*>(ptr), static_cast<void (MyQQuickItem::*)(QQuickWindow *)>(&MyQQuickItem::Signal_WindowChanged));;
}

void QQuickItem_DisconnectWindowChanged(void* ptr){
	QObject::disconnect(static_cast<QQuickItem*>(ptr), static_cast<void (QQuickItem::*)(QQuickWindow *)>(&QQuickItem::windowChanged), static_cast<MyQQuickItem*>(ptr), static_cast<void (MyQQuickItem::*)(QQuickWindow *)>(&MyQQuickItem::Signal_WindowChanged));;
}

void QQuickItem_WindowChanged(void* ptr, void* window){
	static_cast<QQuickItem*>(ptr)->windowChanged(static_cast<QQuickWindow*>(window));
}

void QQuickItem_DestroyQQuickItem(void* ptr){
	static_cast<QQuickItem*>(ptr)->~QQuickItem();
}

void QQuickItem_TimerEvent(void* ptr, void* event){
	static_cast<MyQQuickItem*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QQuickItem_TimerEventDefault(void* ptr, void* event){
	static_cast<QQuickItem*>(ptr)->QQuickItem::timerEvent(static_cast<QTimerEvent*>(event));
}

void QQuickItem_ChildEvent(void* ptr, void* event){
	static_cast<MyQQuickItem*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QQuickItem_ChildEventDefault(void* ptr, void* event){
	static_cast<QQuickItem*>(ptr)->QQuickItem::childEvent(static_cast<QChildEvent*>(event));
}

void QQuickItem_CustomEvent(void* ptr, void* event){
	static_cast<MyQQuickItem*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QQuickItem_CustomEventDefault(void* ptr, void* event){
	static_cast<QQuickItem*>(ptr)->QQuickItem::customEvent(static_cast<QEvent*>(event));
}

class MyQQuickItemGrabResult: public QQuickItemGrabResult {
public:
	void Signal_Ready() { callbackQQuickItemGrabResultReady(this, this->objectName().toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQQuickItemGrabResultTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQQuickItemGrabResultChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQQuickItemGrabResultCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void* QQuickItemGrabResult_Image(void* ptr){
	return new QImage(static_cast<QQuickItemGrabResult*>(ptr)->image());
}

void* QQuickItemGrabResult_Url(void* ptr){
	return new QUrl(static_cast<QQuickItemGrabResult*>(ptr)->url());
}

void QQuickItemGrabResult_ConnectReady(void* ptr){
	QObject::connect(static_cast<QQuickItemGrabResult*>(ptr), static_cast<void (QQuickItemGrabResult::*)()>(&QQuickItemGrabResult::ready), static_cast<MyQQuickItemGrabResult*>(ptr), static_cast<void (MyQQuickItemGrabResult::*)()>(&MyQQuickItemGrabResult::Signal_Ready));;
}

void QQuickItemGrabResult_DisconnectReady(void* ptr){
	QObject::disconnect(static_cast<QQuickItemGrabResult*>(ptr), static_cast<void (QQuickItemGrabResult::*)()>(&QQuickItemGrabResult::ready), static_cast<MyQQuickItemGrabResult*>(ptr), static_cast<void (MyQQuickItemGrabResult::*)()>(&MyQQuickItemGrabResult::Signal_Ready));;
}

void QQuickItemGrabResult_Ready(void* ptr){
	static_cast<QQuickItemGrabResult*>(ptr)->ready();
}

int QQuickItemGrabResult_SaveToFile(void* ptr, char* fileName){
	return static_cast<QQuickItemGrabResult*>(ptr)->saveToFile(QString(fileName));
}

void QQuickItemGrabResult_TimerEvent(void* ptr, void* event){
	static_cast<MyQQuickItemGrabResult*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QQuickItemGrabResult_TimerEventDefault(void* ptr, void* event){
	static_cast<QQuickItemGrabResult*>(ptr)->QQuickItemGrabResult::timerEvent(static_cast<QTimerEvent*>(event));
}

void QQuickItemGrabResult_ChildEvent(void* ptr, void* event){
	static_cast<MyQQuickItemGrabResult*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QQuickItemGrabResult_ChildEventDefault(void* ptr, void* event){
	static_cast<QQuickItemGrabResult*>(ptr)->QQuickItemGrabResult::childEvent(static_cast<QChildEvent*>(event));
}

void QQuickItemGrabResult_CustomEvent(void* ptr, void* event){
	static_cast<MyQQuickItemGrabResult*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QQuickItemGrabResult_CustomEventDefault(void* ptr, void* event){
	static_cast<QQuickItemGrabResult*>(ptr)->QQuickItemGrabResult::customEvent(static_cast<QEvent*>(event));
}

class MyQQuickPaintedItem: public QQuickPaintedItem {
public:
	void Signal_ContentsScaleChanged() { callbackQQuickPaintedItemContentsScaleChanged(this, this->objectName().toUtf8().data()); };
	void Signal_ContentsSizeChanged() { callbackQQuickPaintedItemContentsSizeChanged(this, this->objectName().toUtf8().data()); };
	void Signal_FillColorChanged() { callbackQQuickPaintedItemFillColorChanged(this, this->objectName().toUtf8().data()); };
	void releaseResources() { callbackQQuickPaintedItemReleaseResources(this, this->objectName().toUtf8().data()); };
	void Signal_RenderTargetChanged() { callbackQQuickPaintedItemRenderTargetChanged(this, this->objectName().toUtf8().data()); };
	void classBegin() { callbackQQuickPaintedItemClassBegin(this, this->objectName().toUtf8().data()); };
	void componentComplete() { callbackQQuickPaintedItemComponentComplete(this, this->objectName().toUtf8().data()); };
	void dragEnterEvent(QDragEnterEvent * event) { callbackQQuickPaintedItemDragEnterEvent(this, this->objectName().toUtf8().data(), event); };
	void dragLeaveEvent(QDragLeaveEvent * event) { callbackQQuickPaintedItemDragLeaveEvent(this, this->objectName().toUtf8().data(), event); };
	void dragMoveEvent(QDragMoveEvent * event) { callbackQQuickPaintedItemDragMoveEvent(this, this->objectName().toUtf8().data(), event); };
	void dropEvent(QDropEvent * event) { callbackQQuickPaintedItemDropEvent(this, this->objectName().toUtf8().data(), event); };
	void focusInEvent(QFocusEvent * event) { callbackQQuickPaintedItemFocusInEvent(this, this->objectName().toUtf8().data(), event); };
	void focusOutEvent(QFocusEvent * event) { callbackQQuickPaintedItemFocusOutEvent(this, this->objectName().toUtf8().data(), event); };
	void geometryChanged(const QRectF & newGeometry, const QRectF & oldGeometry) { callbackQQuickPaintedItemGeometryChanged(this, this->objectName().toUtf8().data(), new QRectF(static_cast<QRectF>(newGeometry).x(), static_cast<QRectF>(newGeometry).y(), static_cast<QRectF>(newGeometry).width(), static_cast<QRectF>(newGeometry).height()), new QRectF(static_cast<QRectF>(oldGeometry).x(), static_cast<QRectF>(oldGeometry).y(), static_cast<QRectF>(oldGeometry).width(), static_cast<QRectF>(oldGeometry).height())); };
	void hoverEnterEvent(QHoverEvent * event) { callbackQQuickPaintedItemHoverEnterEvent(this, this->objectName().toUtf8().data(), event); };
	void hoverLeaveEvent(QHoverEvent * event) { callbackQQuickPaintedItemHoverLeaveEvent(this, this->objectName().toUtf8().data(), event); };
	void hoverMoveEvent(QHoverEvent * event) { callbackQQuickPaintedItemHoverMoveEvent(this, this->objectName().toUtf8().data(), event); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackQQuickPaintedItemInputMethodEvent(this, this->objectName().toUtf8().data(), event); };
	void keyPressEvent(QKeyEvent * event) { callbackQQuickPaintedItemKeyPressEvent(this, this->objectName().toUtf8().data(), event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQQuickPaintedItemKeyReleaseEvent(this, this->objectName().toUtf8().data(), event); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQQuickPaintedItemMouseDoubleClickEvent(this, this->objectName().toUtf8().data(), event); };
	void mouseMoveEvent(QMouseEvent * event) { callbackQQuickPaintedItemMouseMoveEvent(this, this->objectName().toUtf8().data(), event); };
	void mousePressEvent(QMouseEvent * event) { callbackQQuickPaintedItemMousePressEvent(this, this->objectName().toUtf8().data(), event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackQQuickPaintedItemMouseReleaseEvent(this, this->objectName().toUtf8().data(), event); };
	void mouseUngrabEvent() { callbackQQuickPaintedItemMouseUngrabEvent(this, this->objectName().toUtf8().data()); };
	void touchEvent(QTouchEvent * event) { callbackQQuickPaintedItemTouchEvent(this, this->objectName().toUtf8().data(), event); };
	void touchUngrabEvent() { callbackQQuickPaintedItemTouchUngrabEvent(this, this->objectName().toUtf8().data()); };
	void updatePolish() { callbackQQuickPaintedItemUpdatePolish(this, this->objectName().toUtf8().data()); };
	void wheelEvent(QWheelEvent * event) { callbackQQuickPaintedItemWheelEvent(this, this->objectName().toUtf8().data(), event); };
	void timerEvent(QTimerEvent * event) { callbackQQuickPaintedItemTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQQuickPaintedItemChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQQuickPaintedItemCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

double QQuickPaintedItem_ContentsScale(void* ptr){
	return static_cast<double>(static_cast<QQuickPaintedItem*>(ptr)->contentsScale());
}

void* QQuickPaintedItem_ContentsSize(void* ptr){
	return new QSize(static_cast<QSize>(static_cast<QQuickPaintedItem*>(ptr)->contentsSize()).width(), static_cast<QSize>(static_cast<QQuickPaintedItem*>(ptr)->contentsSize()).height());
}

void* QQuickPaintedItem_FillColor(void* ptr){
	return new QColor(static_cast<QQuickPaintedItem*>(ptr)->fillColor());
}

int QQuickPaintedItem_RenderTarget(void* ptr){
	return static_cast<QQuickPaintedItem*>(ptr)->renderTarget();
}

void QQuickPaintedItem_SetContentsScale(void* ptr, double v){
	static_cast<QQuickPaintedItem*>(ptr)->setContentsScale(static_cast<double>(v));
}

void QQuickPaintedItem_SetContentsSize(void* ptr, void* v){
	static_cast<QQuickPaintedItem*>(ptr)->setContentsSize(*static_cast<QSize*>(v));
}

void QQuickPaintedItem_SetFillColor(void* ptr, void* v){
	static_cast<QQuickPaintedItem*>(ptr)->setFillColor(*static_cast<QColor*>(v));
}

void QQuickPaintedItem_SetRenderTarget(void* ptr, int target){
	static_cast<QQuickPaintedItem*>(ptr)->setRenderTarget(static_cast<QQuickPaintedItem::RenderTarget>(target));
}

int QQuickPaintedItem_Antialiasing(void* ptr){
	return static_cast<QQuickPaintedItem*>(ptr)->antialiasing();
}

void* QQuickPaintedItem_ContentsBoundingRect(void* ptr){
	return new QRectF(static_cast<QRectF>(static_cast<QQuickPaintedItem*>(ptr)->contentsBoundingRect()).x(), static_cast<QRectF>(static_cast<QQuickPaintedItem*>(ptr)->contentsBoundingRect()).y(), static_cast<QRectF>(static_cast<QQuickPaintedItem*>(ptr)->contentsBoundingRect()).width(), static_cast<QRectF>(static_cast<QQuickPaintedItem*>(ptr)->contentsBoundingRect()).height());
}

void QQuickPaintedItem_ConnectContentsScaleChanged(void* ptr){
	QObject::connect(static_cast<QQuickPaintedItem*>(ptr), static_cast<void (QQuickPaintedItem::*)()>(&QQuickPaintedItem::contentsScaleChanged), static_cast<MyQQuickPaintedItem*>(ptr), static_cast<void (MyQQuickPaintedItem::*)()>(&MyQQuickPaintedItem::Signal_ContentsScaleChanged));;
}

void QQuickPaintedItem_DisconnectContentsScaleChanged(void* ptr){
	QObject::disconnect(static_cast<QQuickPaintedItem*>(ptr), static_cast<void (QQuickPaintedItem::*)()>(&QQuickPaintedItem::contentsScaleChanged), static_cast<MyQQuickPaintedItem*>(ptr), static_cast<void (MyQQuickPaintedItem::*)()>(&MyQQuickPaintedItem::Signal_ContentsScaleChanged));;
}

void QQuickPaintedItem_ContentsScaleChanged(void* ptr){
	static_cast<QQuickPaintedItem*>(ptr)->contentsScaleChanged();
}

void QQuickPaintedItem_ConnectContentsSizeChanged(void* ptr){
	QObject::connect(static_cast<QQuickPaintedItem*>(ptr), static_cast<void (QQuickPaintedItem::*)()>(&QQuickPaintedItem::contentsSizeChanged), static_cast<MyQQuickPaintedItem*>(ptr), static_cast<void (MyQQuickPaintedItem::*)()>(&MyQQuickPaintedItem::Signal_ContentsSizeChanged));;
}

void QQuickPaintedItem_DisconnectContentsSizeChanged(void* ptr){
	QObject::disconnect(static_cast<QQuickPaintedItem*>(ptr), static_cast<void (QQuickPaintedItem::*)()>(&QQuickPaintedItem::contentsSizeChanged), static_cast<MyQQuickPaintedItem*>(ptr), static_cast<void (MyQQuickPaintedItem::*)()>(&MyQQuickPaintedItem::Signal_ContentsSizeChanged));;
}

void QQuickPaintedItem_ContentsSizeChanged(void* ptr){
	static_cast<QQuickPaintedItem*>(ptr)->contentsSizeChanged();
}

void QQuickPaintedItem_ConnectFillColorChanged(void* ptr){
	QObject::connect(static_cast<QQuickPaintedItem*>(ptr), static_cast<void (QQuickPaintedItem::*)()>(&QQuickPaintedItem::fillColorChanged), static_cast<MyQQuickPaintedItem*>(ptr), static_cast<void (MyQQuickPaintedItem::*)()>(&MyQQuickPaintedItem::Signal_FillColorChanged));;
}

void QQuickPaintedItem_DisconnectFillColorChanged(void* ptr){
	QObject::disconnect(static_cast<QQuickPaintedItem*>(ptr), static_cast<void (QQuickPaintedItem::*)()>(&QQuickPaintedItem::fillColorChanged), static_cast<MyQQuickPaintedItem*>(ptr), static_cast<void (MyQQuickPaintedItem::*)()>(&MyQQuickPaintedItem::Signal_FillColorChanged));;
}

void QQuickPaintedItem_FillColorChanged(void* ptr){
	static_cast<QQuickPaintedItem*>(ptr)->fillColorChanged();
}

int QQuickPaintedItem_IsTextureProvider(void* ptr){
	return static_cast<QQuickPaintedItem*>(ptr)->isTextureProvider();
}

int QQuickPaintedItem_Mipmap(void* ptr){
	return static_cast<QQuickPaintedItem*>(ptr)->mipmap();
}

int QQuickPaintedItem_OpaquePainting(void* ptr){
	return static_cast<QQuickPaintedItem*>(ptr)->opaquePainting();
}

void QQuickPaintedItem_Paint(void* ptr, void* painter){
	static_cast<QQuickPaintedItem*>(ptr)->paint(static_cast<QPainter*>(painter));
}

int QQuickPaintedItem_PerformanceHints(void* ptr){
	return static_cast<QQuickPaintedItem*>(ptr)->performanceHints();
}

void QQuickPaintedItem_ReleaseResources(void* ptr){
	static_cast<MyQQuickPaintedItem*>(ptr)->releaseResources();
}

void QQuickPaintedItem_ReleaseResourcesDefault(void* ptr){
	static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::releaseResources();
}

void QQuickPaintedItem_ConnectRenderTargetChanged(void* ptr){
	QObject::connect(static_cast<QQuickPaintedItem*>(ptr), static_cast<void (QQuickPaintedItem::*)()>(&QQuickPaintedItem::renderTargetChanged), static_cast<MyQQuickPaintedItem*>(ptr), static_cast<void (MyQQuickPaintedItem::*)()>(&MyQQuickPaintedItem::Signal_RenderTargetChanged));;
}

void QQuickPaintedItem_DisconnectRenderTargetChanged(void* ptr){
	QObject::disconnect(static_cast<QQuickPaintedItem*>(ptr), static_cast<void (QQuickPaintedItem::*)()>(&QQuickPaintedItem::renderTargetChanged), static_cast<MyQQuickPaintedItem*>(ptr), static_cast<void (MyQQuickPaintedItem::*)()>(&MyQQuickPaintedItem::Signal_RenderTargetChanged));;
}

void QQuickPaintedItem_RenderTargetChanged(void* ptr){
	static_cast<QQuickPaintedItem*>(ptr)->renderTargetChanged();
}

void QQuickPaintedItem_ResetContentsSize(void* ptr){
	static_cast<QQuickPaintedItem*>(ptr)->resetContentsSize();
}

void QQuickPaintedItem_SetAntialiasing(void* ptr, int enable){
	static_cast<QQuickPaintedItem*>(ptr)->setAntialiasing(enable != 0);
}

void QQuickPaintedItem_SetMipmap(void* ptr, int enable){
	static_cast<QQuickPaintedItem*>(ptr)->setMipmap(enable != 0);
}

void QQuickPaintedItem_SetOpaquePainting(void* ptr, int opaque){
	static_cast<QQuickPaintedItem*>(ptr)->setOpaquePainting(opaque != 0);
}

void QQuickPaintedItem_SetPerformanceHint(void* ptr, int hint, int enabled){
	static_cast<QQuickPaintedItem*>(ptr)->setPerformanceHint(static_cast<QQuickPaintedItem::PerformanceHint>(hint), enabled != 0);
}

void QQuickPaintedItem_SetPerformanceHints(void* ptr, int hints){
	static_cast<QQuickPaintedItem*>(ptr)->setPerformanceHints(static_cast<QQuickPaintedItem::PerformanceHint>(hints));
}

void* QQuickPaintedItem_TextureProvider(void* ptr){
	return static_cast<QQuickPaintedItem*>(ptr)->textureProvider();
}

void QQuickPaintedItem_Update(void* ptr, void* rect){
	static_cast<QQuickPaintedItem*>(ptr)->update(*static_cast<QRect*>(rect));
}

void QQuickPaintedItem_DestroyQQuickPaintedItem(void* ptr){
	static_cast<QQuickPaintedItem*>(ptr)->~QQuickPaintedItem();
}

void QQuickPaintedItem_ClassBegin(void* ptr){
	static_cast<MyQQuickPaintedItem*>(ptr)->classBegin();
}

void QQuickPaintedItem_ClassBeginDefault(void* ptr){
	static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::classBegin();
}

void QQuickPaintedItem_ComponentComplete(void* ptr){
	static_cast<MyQQuickPaintedItem*>(ptr)->componentComplete();
}

void QQuickPaintedItem_ComponentCompleteDefault(void* ptr){
	static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::componentComplete();
}

void QQuickPaintedItem_DragEnterEvent(void* ptr, void* event){
	static_cast<MyQQuickPaintedItem*>(ptr)->dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QQuickPaintedItem_DragEnterEventDefault(void* ptr, void* event){
	static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QQuickPaintedItem_DragLeaveEvent(void* ptr, void* event){
	static_cast<MyQQuickPaintedItem*>(ptr)->dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QQuickPaintedItem_DragLeaveEventDefault(void* ptr, void* event){
	static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QQuickPaintedItem_DragMoveEvent(void* ptr, void* event){
	static_cast<MyQQuickPaintedItem*>(ptr)->dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QQuickPaintedItem_DragMoveEventDefault(void* ptr, void* event){
	static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QQuickPaintedItem_DropEvent(void* ptr, void* event){
	static_cast<MyQQuickPaintedItem*>(ptr)->dropEvent(static_cast<QDropEvent*>(event));
}

void QQuickPaintedItem_DropEventDefault(void* ptr, void* event){
	static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::dropEvent(static_cast<QDropEvent*>(event));
}

void QQuickPaintedItem_FocusInEvent(void* ptr, void* event){
	static_cast<MyQQuickPaintedItem*>(ptr)->focusInEvent(static_cast<QFocusEvent*>(event));
}

void QQuickPaintedItem_FocusInEventDefault(void* ptr, void* event){
	static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::focusInEvent(static_cast<QFocusEvent*>(event));
}

void QQuickPaintedItem_FocusOutEvent(void* ptr, void* event){
	static_cast<MyQQuickPaintedItem*>(ptr)->focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QQuickPaintedItem_FocusOutEventDefault(void* ptr, void* event){
	static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QQuickPaintedItem_GeometryChanged(void* ptr, void* newGeometry, void* oldGeometry){
	static_cast<MyQQuickPaintedItem*>(ptr)->geometryChanged(*static_cast<QRectF*>(newGeometry), *static_cast<QRectF*>(oldGeometry));
}

void QQuickPaintedItem_GeometryChangedDefault(void* ptr, void* newGeometry, void* oldGeometry){
	static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::geometryChanged(*static_cast<QRectF*>(newGeometry), *static_cast<QRectF*>(oldGeometry));
}

void QQuickPaintedItem_HoverEnterEvent(void* ptr, void* event){
	static_cast<MyQQuickPaintedItem*>(ptr)->hoverEnterEvent(static_cast<QHoverEvent*>(event));
}

void QQuickPaintedItem_HoverEnterEventDefault(void* ptr, void* event){
	static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::hoverEnterEvent(static_cast<QHoverEvent*>(event));
}

void QQuickPaintedItem_HoverLeaveEvent(void* ptr, void* event){
	static_cast<MyQQuickPaintedItem*>(ptr)->hoverLeaveEvent(static_cast<QHoverEvent*>(event));
}

void QQuickPaintedItem_HoverLeaveEventDefault(void* ptr, void* event){
	static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::hoverLeaveEvent(static_cast<QHoverEvent*>(event));
}

void QQuickPaintedItem_HoverMoveEvent(void* ptr, void* event){
	static_cast<MyQQuickPaintedItem*>(ptr)->hoverMoveEvent(static_cast<QHoverEvent*>(event));
}

void QQuickPaintedItem_HoverMoveEventDefault(void* ptr, void* event){
	static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::hoverMoveEvent(static_cast<QHoverEvent*>(event));
}

void QQuickPaintedItem_InputMethodEvent(void* ptr, void* event){
	static_cast<MyQQuickPaintedItem*>(ptr)->inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void QQuickPaintedItem_InputMethodEventDefault(void* ptr, void* event){
	static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void QQuickPaintedItem_KeyPressEvent(void* ptr, void* event){
	static_cast<MyQQuickPaintedItem*>(ptr)->keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QQuickPaintedItem_KeyPressEventDefault(void* ptr, void* event){
	static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QQuickPaintedItem_KeyReleaseEvent(void* ptr, void* event){
	static_cast<MyQQuickPaintedItem*>(ptr)->keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QQuickPaintedItem_KeyReleaseEventDefault(void* ptr, void* event){
	static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QQuickPaintedItem_MouseDoubleClickEvent(void* ptr, void* event){
	static_cast<MyQQuickPaintedItem*>(ptr)->mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QQuickPaintedItem_MouseDoubleClickEventDefault(void* ptr, void* event){
	static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QQuickPaintedItem_MouseMoveEvent(void* ptr, void* event){
	static_cast<MyQQuickPaintedItem*>(ptr)->mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QQuickPaintedItem_MouseMoveEventDefault(void* ptr, void* event){
	static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QQuickPaintedItem_MousePressEvent(void* ptr, void* event){
	static_cast<MyQQuickPaintedItem*>(ptr)->mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QQuickPaintedItem_MousePressEventDefault(void* ptr, void* event){
	static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QQuickPaintedItem_MouseReleaseEvent(void* ptr, void* event){
	static_cast<MyQQuickPaintedItem*>(ptr)->mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void QQuickPaintedItem_MouseReleaseEventDefault(void* ptr, void* event){
	static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void QQuickPaintedItem_MouseUngrabEvent(void* ptr){
	static_cast<MyQQuickPaintedItem*>(ptr)->mouseUngrabEvent();
}

void QQuickPaintedItem_MouseUngrabEventDefault(void* ptr){
	static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::mouseUngrabEvent();
}

void QQuickPaintedItem_TouchEvent(void* ptr, void* event){
	static_cast<MyQQuickPaintedItem*>(ptr)->touchEvent(static_cast<QTouchEvent*>(event));
}

void QQuickPaintedItem_TouchEventDefault(void* ptr, void* event){
	static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::touchEvent(static_cast<QTouchEvent*>(event));
}

void QQuickPaintedItem_TouchUngrabEvent(void* ptr){
	static_cast<MyQQuickPaintedItem*>(ptr)->touchUngrabEvent();
}

void QQuickPaintedItem_TouchUngrabEventDefault(void* ptr){
	static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::touchUngrabEvent();
}

void QQuickPaintedItem_UpdatePolish(void* ptr){
	static_cast<MyQQuickPaintedItem*>(ptr)->updatePolish();
}

void QQuickPaintedItem_UpdatePolishDefault(void* ptr){
	static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::updatePolish();
}

void QQuickPaintedItem_WheelEvent(void* ptr, void* event){
	static_cast<MyQQuickPaintedItem*>(ptr)->wheelEvent(static_cast<QWheelEvent*>(event));
}

void QQuickPaintedItem_WheelEventDefault(void* ptr, void* event){
	static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::wheelEvent(static_cast<QWheelEvent*>(event));
}

void QQuickPaintedItem_TimerEvent(void* ptr, void* event){
	static_cast<MyQQuickPaintedItem*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QQuickPaintedItem_TimerEventDefault(void* ptr, void* event){
	static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::timerEvent(static_cast<QTimerEvent*>(event));
}

void QQuickPaintedItem_ChildEvent(void* ptr, void* event){
	static_cast<MyQQuickPaintedItem*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QQuickPaintedItem_ChildEventDefault(void* ptr, void* event){
	static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::childEvent(static_cast<QChildEvent*>(event));
}

void QQuickPaintedItem_CustomEvent(void* ptr, void* event){
	static_cast<MyQQuickPaintedItem*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QQuickPaintedItem_CustomEventDefault(void* ptr, void* event){
	static_cast<QQuickPaintedItem*>(ptr)->QQuickPaintedItem::customEvent(static_cast<QEvent*>(event));
}

class MyQQuickRenderControl: public QQuickRenderControl {
public:
	MyQQuickRenderControl(QObject *parent) : QQuickRenderControl(parent) {};
	void Signal_RenderRequested() { callbackQQuickRenderControlRenderRequested(this, this->objectName().toUtf8().data()); };
	void Signal_SceneChanged() { callbackQQuickRenderControlSceneChanged(this, this->objectName().toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQQuickRenderControlTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQQuickRenderControlChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQQuickRenderControlCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void* QQuickRenderControl_NewQQuickRenderControl(void* parent){
	return new MyQQuickRenderControl(static_cast<QObject*>(parent));
}

void* QQuickRenderControl_Grab(void* ptr){
	return new QImage(static_cast<QQuickRenderControl*>(ptr)->grab());
}

void QQuickRenderControl_Initialize(void* ptr, void* gl){
	static_cast<QQuickRenderControl*>(ptr)->initialize(static_cast<QOpenGLContext*>(gl));
}

void QQuickRenderControl_Invalidate(void* ptr){
	static_cast<QQuickRenderControl*>(ptr)->invalidate();
}

void QQuickRenderControl_PolishItems(void* ptr){
	static_cast<QQuickRenderControl*>(ptr)->polishItems();
}

void QQuickRenderControl_PrepareThread(void* ptr, void* targetThread){
	static_cast<QQuickRenderControl*>(ptr)->prepareThread(static_cast<QThread*>(targetThread));
}

void QQuickRenderControl_Render(void* ptr){
	static_cast<QQuickRenderControl*>(ptr)->render();
}

void QQuickRenderControl_ConnectRenderRequested(void* ptr){
	QObject::connect(static_cast<QQuickRenderControl*>(ptr), static_cast<void (QQuickRenderControl::*)()>(&QQuickRenderControl::renderRequested), static_cast<MyQQuickRenderControl*>(ptr), static_cast<void (MyQQuickRenderControl::*)()>(&MyQQuickRenderControl::Signal_RenderRequested));;
}

void QQuickRenderControl_DisconnectRenderRequested(void* ptr){
	QObject::disconnect(static_cast<QQuickRenderControl*>(ptr), static_cast<void (QQuickRenderControl::*)()>(&QQuickRenderControl::renderRequested), static_cast<MyQQuickRenderControl*>(ptr), static_cast<void (MyQQuickRenderControl::*)()>(&MyQQuickRenderControl::Signal_RenderRequested));;
}

void QQuickRenderControl_RenderRequested(void* ptr){
	static_cast<QQuickRenderControl*>(ptr)->renderRequested();
}

void* QQuickRenderControl_RenderWindow(void* ptr, void* offset){
	return static_cast<QQuickRenderControl*>(ptr)->renderWindow(static_cast<QPoint*>(offset));
}

void* QQuickRenderControl_QQuickRenderControl_RenderWindowFor(void* win, void* offset){
	return QQuickRenderControl::renderWindowFor(static_cast<QQuickWindow*>(win), static_cast<QPoint*>(offset));
}

void QQuickRenderControl_ConnectSceneChanged(void* ptr){
	QObject::connect(static_cast<QQuickRenderControl*>(ptr), static_cast<void (QQuickRenderControl::*)()>(&QQuickRenderControl::sceneChanged), static_cast<MyQQuickRenderControl*>(ptr), static_cast<void (MyQQuickRenderControl::*)()>(&MyQQuickRenderControl::Signal_SceneChanged));;
}

void QQuickRenderControl_DisconnectSceneChanged(void* ptr){
	QObject::disconnect(static_cast<QQuickRenderControl*>(ptr), static_cast<void (QQuickRenderControl::*)()>(&QQuickRenderControl::sceneChanged), static_cast<MyQQuickRenderControl*>(ptr), static_cast<void (MyQQuickRenderControl::*)()>(&MyQQuickRenderControl::Signal_SceneChanged));;
}

void QQuickRenderControl_SceneChanged(void* ptr){
	static_cast<QQuickRenderControl*>(ptr)->sceneChanged();
}

int QQuickRenderControl_Sync(void* ptr){
	return static_cast<QQuickRenderControl*>(ptr)->sync();
}

void QQuickRenderControl_DestroyQQuickRenderControl(void* ptr){
	static_cast<QQuickRenderControl*>(ptr)->~QQuickRenderControl();
}

void QQuickRenderControl_TimerEvent(void* ptr, void* event){
	static_cast<MyQQuickRenderControl*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QQuickRenderControl_TimerEventDefault(void* ptr, void* event){
	static_cast<QQuickRenderControl*>(ptr)->QQuickRenderControl::timerEvent(static_cast<QTimerEvent*>(event));
}

void QQuickRenderControl_ChildEvent(void* ptr, void* event){
	static_cast<MyQQuickRenderControl*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QQuickRenderControl_ChildEventDefault(void* ptr, void* event){
	static_cast<QQuickRenderControl*>(ptr)->QQuickRenderControl::childEvent(static_cast<QChildEvent*>(event));
}

void QQuickRenderControl_CustomEvent(void* ptr, void* event){
	static_cast<MyQQuickRenderControl*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QQuickRenderControl_CustomEventDefault(void* ptr, void* event){
	static_cast<QQuickRenderControl*>(ptr)->QQuickRenderControl::customEvent(static_cast<QEvent*>(event));
}

void* QQuickTextDocument_NewQQuickTextDocument(void* parent){
	return new QQuickTextDocument(static_cast<QQuickItem*>(parent));
}

void* QQuickTextDocument_TextDocument(void* ptr){
	return static_cast<QQuickTextDocument*>(ptr)->textDocument();
}

void QQuickTextDocument_TimerEvent(void* ptr, void* event){
	static_cast<QQuickTextDocument*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QQuickTextDocument_TimerEventDefault(void* ptr, void* event){
	static_cast<QQuickTextDocument*>(ptr)->QQuickTextDocument::timerEvent(static_cast<QTimerEvent*>(event));
}

void QQuickTextDocument_ChildEvent(void* ptr, void* event){
	static_cast<QQuickTextDocument*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QQuickTextDocument_ChildEventDefault(void* ptr, void* event){
	static_cast<QQuickTextDocument*>(ptr)->QQuickTextDocument::childEvent(static_cast<QChildEvent*>(event));
}

void QQuickTextDocument_CustomEvent(void* ptr, void* event){
	static_cast<QQuickTextDocument*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QQuickTextDocument_CustomEventDefault(void* ptr, void* event){
	static_cast<QQuickTextDocument*>(ptr)->QQuickTextDocument::customEvent(static_cast<QEvent*>(event));
}

class MyQQuickTextureFactory: public QQuickTextureFactory {
public:
	void timerEvent(QTimerEvent * event) { callbackQQuickTextureFactoryTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQQuickTextureFactoryChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQQuickTextureFactoryCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void* QQuickTextureFactory_Image(void* ptr){
	return new QImage(static_cast<QQuickTextureFactory*>(ptr)->image());
}

void* QQuickTextureFactory_CreateTexture(void* ptr, void* window){
	return static_cast<QQuickTextureFactory*>(ptr)->createTexture(static_cast<QQuickWindow*>(window));
}

int QQuickTextureFactory_TextureByteCount(void* ptr){
	return static_cast<QQuickTextureFactory*>(ptr)->textureByteCount();
}

void* QQuickTextureFactory_TextureSize(void* ptr){
	return new QSize(static_cast<QSize>(static_cast<QQuickTextureFactory*>(ptr)->textureSize()).width(), static_cast<QSize>(static_cast<QQuickTextureFactory*>(ptr)->textureSize()).height());
}

void QQuickTextureFactory_DestroyQQuickTextureFactory(void* ptr){
	static_cast<QQuickTextureFactory*>(ptr)->~QQuickTextureFactory();
}

void QQuickTextureFactory_TimerEvent(void* ptr, void* event){
	static_cast<MyQQuickTextureFactory*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QQuickTextureFactory_TimerEventDefault(void* ptr, void* event){
	static_cast<QQuickTextureFactory*>(ptr)->QQuickTextureFactory::timerEvent(static_cast<QTimerEvent*>(event));
}

void QQuickTextureFactory_ChildEvent(void* ptr, void* event){
	static_cast<MyQQuickTextureFactory*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QQuickTextureFactory_ChildEventDefault(void* ptr, void* event){
	static_cast<QQuickTextureFactory*>(ptr)->QQuickTextureFactory::childEvent(static_cast<QChildEvent*>(event));
}

void QQuickTextureFactory_CustomEvent(void* ptr, void* event){
	static_cast<MyQQuickTextureFactory*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QQuickTextureFactory_CustomEventDefault(void* ptr, void* event){
	static_cast<QQuickTextureFactory*>(ptr)->QQuickTextureFactory::customEvent(static_cast<QEvent*>(event));
}

class MyQQuickView: public QQuickView {
public:
	MyQQuickView(QQmlEngine *engine, QWindow *parent) : QQuickView(engine, parent) {};
	MyQQuickView(QWindow *parent) : QQuickView(parent) {};
	MyQQuickView(const QUrl &source, QWindow *parent) : QQuickView(source, parent) {};
	void keyPressEvent(QKeyEvent * e) { callbackQQuickViewKeyPressEvent(this, this->objectName().toUtf8().data(), e); };
	void keyReleaseEvent(QKeyEvent * e) { callbackQQuickViewKeyReleaseEvent(this, this->objectName().toUtf8().data(), e); };
	void mouseMoveEvent(QMouseEvent * e) { callbackQQuickViewMouseMoveEvent(this, this->objectName().toUtf8().data(), e); };
	void mousePressEvent(QMouseEvent * e) { callbackQQuickViewMousePressEvent(this, this->objectName().toUtf8().data(), e); };
	void mouseReleaseEvent(QMouseEvent * e) { callbackQQuickViewMouseReleaseEvent(this, this->objectName().toUtf8().data(), e); };
	void Signal_StatusChanged(QQuickView::Status status) { callbackQQuickViewStatusChanged(this, this->objectName().toUtf8().data(), status); };
	void exposeEvent(QExposeEvent * v) { callbackQQuickViewExposeEvent(this, this->objectName().toUtf8().data(), v); };
	void focusInEvent(QFocusEvent * ev) { callbackQQuickViewFocusInEvent(this, this->objectName().toUtf8().data(), ev); };
	void focusOutEvent(QFocusEvent * ev) { callbackQQuickViewFocusOutEvent(this, this->objectName().toUtf8().data(), ev); };
	void hideEvent(QHideEvent * v) { callbackQQuickViewHideEvent(this, this->objectName().toUtf8().data(), v); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQQuickViewMouseDoubleClickEvent(this, this->objectName().toUtf8().data(), event); };
	void resizeEvent(QResizeEvent * ev) { callbackQQuickViewResizeEvent(this, this->objectName().toUtf8().data(), ev); };
	void showEvent(QShowEvent * v) { callbackQQuickViewShowEvent(this, this->objectName().toUtf8().data(), v); };
	void wheelEvent(QWheelEvent * event) { callbackQQuickViewWheelEvent(this, this->objectName().toUtf8().data(), event); };
	void moveEvent(QMoveEvent * ev) { callbackQQuickViewMoveEvent(this, this->objectName().toUtf8().data(), ev); };
	void tabletEvent(QTabletEvent * ev) { callbackQQuickViewTabletEvent(this, this->objectName().toUtf8().data(), ev); };
	void touchEvent(QTouchEvent * ev) { callbackQQuickViewTouchEvent(this, this->objectName().toUtf8().data(), ev); };
	void timerEvent(QTimerEvent * event) { callbackQQuickViewTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQQuickViewChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQQuickViewCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

int QQuickView_ResizeMode(void* ptr){
	return static_cast<QQuickView*>(ptr)->resizeMode();
}

void QQuickView_SetResizeMode(void* ptr, int v){
	static_cast<QQuickView*>(ptr)->setResizeMode(static_cast<QQuickView::ResizeMode>(v));
}

int QQuickView_Status(void* ptr){
	return static_cast<QQuickView*>(ptr)->status();
}

void* QQuickView_NewQQuickView2(void* engine, void* parent){
	return new MyQQuickView(static_cast<QQmlEngine*>(engine), static_cast<QWindow*>(parent));
}

void* QQuickView_NewQQuickView(void* parent){
	return new MyQQuickView(static_cast<QWindow*>(parent));
}

void* QQuickView_NewQQuickView3(void* source, void* parent){
	return new MyQQuickView(*static_cast<QUrl*>(source), static_cast<QWindow*>(parent));
}

void* QQuickView_Engine(void* ptr){
	return static_cast<QQuickView*>(ptr)->engine();
}

void* QQuickView_InitialSize(void* ptr){
	return new QSize(static_cast<QSize>(static_cast<QQuickView*>(ptr)->initialSize()).width(), static_cast<QSize>(static_cast<QQuickView*>(ptr)->initialSize()).height());
}

void QQuickView_KeyPressEvent(void* ptr, void* e){
	static_cast<MyQQuickView*>(ptr)->keyPressEvent(static_cast<QKeyEvent*>(e));
}

void QQuickView_KeyPressEventDefault(void* ptr, void* e){
	static_cast<QQuickView*>(ptr)->QQuickView::keyPressEvent(static_cast<QKeyEvent*>(e));
}

void QQuickView_KeyReleaseEvent(void* ptr, void* e){
	static_cast<MyQQuickView*>(ptr)->keyReleaseEvent(static_cast<QKeyEvent*>(e));
}

void QQuickView_KeyReleaseEventDefault(void* ptr, void* e){
	static_cast<QQuickView*>(ptr)->QQuickView::keyReleaseEvent(static_cast<QKeyEvent*>(e));
}

void QQuickView_MouseMoveEvent(void* ptr, void* e){
	static_cast<MyQQuickView*>(ptr)->mouseMoveEvent(static_cast<QMouseEvent*>(e));
}

void QQuickView_MouseMoveEventDefault(void* ptr, void* e){
	static_cast<QQuickView*>(ptr)->QQuickView::mouseMoveEvent(static_cast<QMouseEvent*>(e));
}

void QQuickView_MousePressEvent(void* ptr, void* e){
	static_cast<MyQQuickView*>(ptr)->mousePressEvent(static_cast<QMouseEvent*>(e));
}

void QQuickView_MousePressEventDefault(void* ptr, void* e){
	static_cast<QQuickView*>(ptr)->QQuickView::mousePressEvent(static_cast<QMouseEvent*>(e));
}

void QQuickView_MouseReleaseEvent(void* ptr, void* e){
	static_cast<MyQQuickView*>(ptr)->mouseReleaseEvent(static_cast<QMouseEvent*>(e));
}

void QQuickView_MouseReleaseEventDefault(void* ptr, void* e){
	static_cast<QQuickView*>(ptr)->QQuickView::mouseReleaseEvent(static_cast<QMouseEvent*>(e));
}

void* QQuickView_RootContext(void* ptr){
	return static_cast<QQuickView*>(ptr)->rootContext();
}

void* QQuickView_RootObject(void* ptr){
	return static_cast<QQuickView*>(ptr)->rootObject();
}

void QQuickView_SetSource(void* ptr, void* url){
	QMetaObject::invokeMethod(static_cast<QQuickView*>(ptr), "setSource", Q_ARG(QUrl, *static_cast<QUrl*>(url)));
}

void* QQuickView_Source(void* ptr){
	return new QUrl(static_cast<QQuickView*>(ptr)->source());
}

void QQuickView_ConnectStatusChanged(void* ptr){
	QObject::connect(static_cast<QQuickView*>(ptr), static_cast<void (QQuickView::*)(QQuickView::Status)>(&QQuickView::statusChanged), static_cast<MyQQuickView*>(ptr), static_cast<void (MyQQuickView::*)(QQuickView::Status)>(&MyQQuickView::Signal_StatusChanged));;
}

void QQuickView_DisconnectStatusChanged(void* ptr){
	QObject::disconnect(static_cast<QQuickView*>(ptr), static_cast<void (QQuickView::*)(QQuickView::Status)>(&QQuickView::statusChanged), static_cast<MyQQuickView*>(ptr), static_cast<void (MyQQuickView::*)(QQuickView::Status)>(&MyQQuickView::Signal_StatusChanged));;
}

void QQuickView_StatusChanged(void* ptr, int status){
	static_cast<QQuickView*>(ptr)->statusChanged(static_cast<QQuickView::Status>(status));
}

void QQuickView_DestroyQQuickView(void* ptr){
	static_cast<QQuickView*>(ptr)->~QQuickView();
}

void QQuickView_ExposeEvent(void* ptr, void* v){
	static_cast<MyQQuickView*>(ptr)->exposeEvent(static_cast<QExposeEvent*>(v));
}

void QQuickView_ExposeEventDefault(void* ptr, void* v){
	static_cast<QQuickView*>(ptr)->QQuickView::exposeEvent(static_cast<QExposeEvent*>(v));
}

void QQuickView_FocusInEvent(void* ptr, void* ev){
	static_cast<MyQQuickView*>(ptr)->focusInEvent(static_cast<QFocusEvent*>(ev));
}

void QQuickView_FocusInEventDefault(void* ptr, void* ev){
	static_cast<QQuickView*>(ptr)->QQuickView::focusInEvent(static_cast<QFocusEvent*>(ev));
}

void QQuickView_FocusOutEvent(void* ptr, void* ev){
	static_cast<MyQQuickView*>(ptr)->focusOutEvent(static_cast<QFocusEvent*>(ev));
}

void QQuickView_FocusOutEventDefault(void* ptr, void* ev){
	static_cast<QQuickView*>(ptr)->QQuickView::focusOutEvent(static_cast<QFocusEvent*>(ev));
}

void QQuickView_HideEvent(void* ptr, void* v){
	static_cast<MyQQuickView*>(ptr)->hideEvent(static_cast<QHideEvent*>(v));
}

void QQuickView_HideEventDefault(void* ptr, void* v){
	static_cast<QQuickView*>(ptr)->QQuickView::hideEvent(static_cast<QHideEvent*>(v));
}

void QQuickView_MouseDoubleClickEvent(void* ptr, void* event){
	static_cast<MyQQuickView*>(ptr)->mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QQuickView_MouseDoubleClickEventDefault(void* ptr, void* event){
	static_cast<QQuickView*>(ptr)->QQuickView::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QQuickView_ResizeEvent(void* ptr, void* ev){
	static_cast<MyQQuickView*>(ptr)->resizeEvent(static_cast<QResizeEvent*>(ev));
}

void QQuickView_ResizeEventDefault(void* ptr, void* ev){
	static_cast<QQuickView*>(ptr)->QQuickView::resizeEvent(static_cast<QResizeEvent*>(ev));
}

void QQuickView_ShowEvent(void* ptr, void* v){
	static_cast<MyQQuickView*>(ptr)->showEvent(static_cast<QShowEvent*>(v));
}

void QQuickView_ShowEventDefault(void* ptr, void* v){
	static_cast<QQuickView*>(ptr)->QQuickView::showEvent(static_cast<QShowEvent*>(v));
}

void QQuickView_WheelEvent(void* ptr, void* event){
	static_cast<MyQQuickView*>(ptr)->wheelEvent(static_cast<QWheelEvent*>(event));
}

void QQuickView_WheelEventDefault(void* ptr, void* event){
	static_cast<QQuickView*>(ptr)->QQuickView::wheelEvent(static_cast<QWheelEvent*>(event));
}

void QQuickView_MoveEvent(void* ptr, void* ev){
	static_cast<MyQQuickView*>(ptr)->moveEvent(static_cast<QMoveEvent*>(ev));
}

void QQuickView_MoveEventDefault(void* ptr, void* ev){
	static_cast<QQuickView*>(ptr)->QQuickView::moveEvent(static_cast<QMoveEvent*>(ev));
}

void QQuickView_TabletEvent(void* ptr, void* ev){
	static_cast<MyQQuickView*>(ptr)->tabletEvent(static_cast<QTabletEvent*>(ev));
}

void QQuickView_TabletEventDefault(void* ptr, void* ev){
	static_cast<QQuickView*>(ptr)->QQuickView::tabletEvent(static_cast<QTabletEvent*>(ev));
}

void QQuickView_TouchEvent(void* ptr, void* ev){
	static_cast<MyQQuickView*>(ptr)->touchEvent(static_cast<QTouchEvent*>(ev));
}

void QQuickView_TouchEventDefault(void* ptr, void* ev){
	static_cast<QQuickView*>(ptr)->QQuickView::touchEvent(static_cast<QTouchEvent*>(ev));
}

void QQuickView_TimerEvent(void* ptr, void* event){
	static_cast<MyQQuickView*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QQuickView_TimerEventDefault(void* ptr, void* event){
	static_cast<QQuickView*>(ptr)->QQuickView::timerEvent(static_cast<QTimerEvent*>(event));
}

void QQuickView_ChildEvent(void* ptr, void* event){
	static_cast<MyQQuickView*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QQuickView_ChildEventDefault(void* ptr, void* event){
	static_cast<QQuickView*>(ptr)->QQuickView::childEvent(static_cast<QChildEvent*>(event));
}

void QQuickView_CustomEvent(void* ptr, void* event){
	static_cast<MyQQuickView*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QQuickView_CustomEventDefault(void* ptr, void* event){
	static_cast<QQuickView*>(ptr)->QQuickView::customEvent(static_cast<QEvent*>(event));
}

class MyQQuickWidget: public QQuickWidget {
public:
	MyQQuickWidget(QQmlEngine *engine, QWidget *parent) : QQuickWidget(engine, parent) {};
	MyQQuickWidget(QWidget *parent) : QQuickWidget(parent) {};
	MyQQuickWidget(const QUrl &source, QWidget *parent) : QQuickWidget(source, parent) {};
	void focusInEvent(QFocusEvent * event) { callbackQQuickWidgetFocusInEvent(this, this->objectName().toUtf8().data(), event); };
	void focusOutEvent(QFocusEvent * event) { callbackQQuickWidgetFocusOutEvent(this, this->objectName().toUtf8().data(), event); };
	void dragEnterEvent(QDragEnterEvent * e) { callbackQQuickWidgetDragEnterEvent(this, this->objectName().toUtf8().data(), e); };
	void dragLeaveEvent(QDragLeaveEvent * e) { callbackQQuickWidgetDragLeaveEvent(this, this->objectName().toUtf8().data(), e); };
	void dragMoveEvent(QDragMoveEvent * e) { callbackQQuickWidgetDragMoveEvent(this, this->objectName().toUtf8().data(), e); };
	void dropEvent(QDropEvent * e) { callbackQQuickWidgetDropEvent(this, this->objectName().toUtf8().data(), e); };
	void hideEvent(QHideEvent * v) { callbackQQuickWidgetHideEvent(this, this->objectName().toUtf8().data(), v); };
	void keyPressEvent(QKeyEvent * e) { callbackQQuickWidgetKeyPressEvent(this, this->objectName().toUtf8().data(), e); };
	void keyReleaseEvent(QKeyEvent * e) { callbackQQuickWidgetKeyReleaseEvent(this, this->objectName().toUtf8().data(), e); };
	void mouseDoubleClickEvent(QMouseEvent * e) { callbackQQuickWidgetMouseDoubleClickEvent(this, this->objectName().toUtf8().data(), e); };
	void mouseMoveEvent(QMouseEvent * e) { callbackQQuickWidgetMouseMoveEvent(this, this->objectName().toUtf8().data(), e); };
	void mousePressEvent(QMouseEvent * e) { callbackQQuickWidgetMousePressEvent(this, this->objectName().toUtf8().data(), e); };
	void mouseReleaseEvent(QMouseEvent * e) { callbackQQuickWidgetMouseReleaseEvent(this, this->objectName().toUtf8().data(), e); };
	void Signal_SceneGraphError(QQuickWindow::SceneGraphError error, const QString & message) { callbackQQuickWidgetSceneGraphError(this, this->objectName().toUtf8().data(), error, message.toUtf8().data()); };
	void showEvent(QShowEvent * v) { callbackQQuickWidgetShowEvent(this, this->objectName().toUtf8().data(), v); };
	void Signal_StatusChanged(QQuickWidget::Status status) { callbackQQuickWidgetStatusChanged(this, this->objectName().toUtf8().data(), status); };
	void wheelEvent(QWheelEvent * e) { callbackQQuickWidgetWheelEvent(this, this->objectName().toUtf8().data(), e); };
	void actionEvent(QActionEvent * event) { callbackQQuickWidgetActionEvent(this, this->objectName().toUtf8().data(), event); };
	void enterEvent(QEvent * event) { callbackQQuickWidgetEnterEvent(this, this->objectName().toUtf8().data(), event); };
	void leaveEvent(QEvent * event) { callbackQQuickWidgetLeaveEvent(this, this->objectName().toUtf8().data(), event); };
	void moveEvent(QMoveEvent * event) { callbackQQuickWidgetMoveEvent(this, this->objectName().toUtf8().data(), event); };
	void paintEvent(QPaintEvent * event) { callbackQQuickWidgetPaintEvent(this, this->objectName().toUtf8().data(), event); };
	void setVisible(bool visible) { if (!callbackQQuickWidgetSetVisible(this, this->objectName().toUtf8().data(), visible)) { QQuickWidget::setVisible(visible); }; };
	void changeEvent(QEvent * event) { callbackQQuickWidgetChangeEvent(this, this->objectName().toUtf8().data(), event); };
	void closeEvent(QCloseEvent * event) { callbackQQuickWidgetCloseEvent(this, this->objectName().toUtf8().data(), event); };
	void contextMenuEvent(QContextMenuEvent * event) { callbackQQuickWidgetContextMenuEvent(this, this->objectName().toUtf8().data(), event); };
	void initPainter(QPainter * painter) const { callbackQQuickWidgetInitPainter(const_cast<MyQQuickWidget*>(this), this->objectName().toUtf8().data(), painter); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackQQuickWidgetInputMethodEvent(this, this->objectName().toUtf8().data(), event); };
	void resizeEvent(QResizeEvent * event) { callbackQQuickWidgetResizeEvent(this, this->objectName().toUtf8().data(), event); };
	void tabletEvent(QTabletEvent * event) { callbackQQuickWidgetTabletEvent(this, this->objectName().toUtf8().data(), event); };
	void timerEvent(QTimerEvent * event) { callbackQQuickWidgetTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQQuickWidgetChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQQuickWidgetCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void QQuickWidget_FocusInEvent(void* ptr, void* event){
	static_cast<MyQQuickWidget*>(ptr)->focusInEvent(static_cast<QFocusEvent*>(event));
}

void QQuickWidget_FocusInEventDefault(void* ptr, void* event){
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::focusInEvent(static_cast<QFocusEvent*>(event));
}

void QQuickWidget_FocusOutEvent(void* ptr, void* event){
	static_cast<MyQQuickWidget*>(ptr)->focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QQuickWidget_FocusOutEventDefault(void* ptr, void* event){
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::focusOutEvent(static_cast<QFocusEvent*>(event));
}

int QQuickWidget_ResizeMode(void* ptr){
	return static_cast<QQuickWidget*>(ptr)->resizeMode();
}

void QQuickWidget_SetResizeMode(void* ptr, int v){
	static_cast<QQuickWidget*>(ptr)->setResizeMode(static_cast<QQuickWidget::ResizeMode>(v));
}

int QQuickWidget_Status(void* ptr){
	return static_cast<QQuickWidget*>(ptr)->status();
}

void* QQuickWidget_NewQQuickWidget2(void* engine, void* parent){
	return new MyQQuickWidget(static_cast<QQmlEngine*>(engine), static_cast<QWidget*>(parent));
}

void* QQuickWidget_NewQQuickWidget(void* parent){
	return new MyQQuickWidget(static_cast<QWidget*>(parent));
}

void* QQuickWidget_NewQQuickWidget3(void* source, void* parent){
	return new MyQQuickWidget(*static_cast<QUrl*>(source), static_cast<QWidget*>(parent));
}

void QQuickWidget_DragEnterEvent(void* ptr, void* e){
	static_cast<MyQQuickWidget*>(ptr)->dragEnterEvent(static_cast<QDragEnterEvent*>(e));
}

void QQuickWidget_DragEnterEventDefault(void* ptr, void* e){
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::dragEnterEvent(static_cast<QDragEnterEvent*>(e));
}

void QQuickWidget_DragLeaveEvent(void* ptr, void* e){
	static_cast<MyQQuickWidget*>(ptr)->dragLeaveEvent(static_cast<QDragLeaveEvent*>(e));
}

void QQuickWidget_DragLeaveEventDefault(void* ptr, void* e){
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::dragLeaveEvent(static_cast<QDragLeaveEvent*>(e));
}

void QQuickWidget_DragMoveEvent(void* ptr, void* e){
	static_cast<MyQQuickWidget*>(ptr)->dragMoveEvent(static_cast<QDragMoveEvent*>(e));
}

void QQuickWidget_DragMoveEventDefault(void* ptr, void* e){
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::dragMoveEvent(static_cast<QDragMoveEvent*>(e));
}

void QQuickWidget_DropEvent(void* ptr, void* e){
	static_cast<MyQQuickWidget*>(ptr)->dropEvent(static_cast<QDropEvent*>(e));
}

void QQuickWidget_DropEventDefault(void* ptr, void* e){
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::dropEvent(static_cast<QDropEvent*>(e));
}

void* QQuickWidget_Engine(void* ptr){
	return static_cast<QQuickWidget*>(ptr)->engine();
}

int QQuickWidget_Event(void* ptr, void* e){
	return static_cast<QQuickWidget*>(ptr)->event(static_cast<QEvent*>(e));
}

void* QQuickWidget_Format(void* ptr){
	return new QSurfaceFormat(static_cast<QQuickWidget*>(ptr)->format());
}

void* QQuickWidget_GrabFramebuffer(void* ptr){
	return new QImage(static_cast<QQuickWidget*>(ptr)->grabFramebuffer());
}

void QQuickWidget_HideEvent(void* ptr, void* v){
	static_cast<MyQQuickWidget*>(ptr)->hideEvent(static_cast<QHideEvent*>(v));
}

void QQuickWidget_HideEventDefault(void* ptr, void* v){
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::hideEvent(static_cast<QHideEvent*>(v));
}

void* QQuickWidget_InitialSize(void* ptr){
	return new QSize(static_cast<QSize>(static_cast<QQuickWidget*>(ptr)->initialSize()).width(), static_cast<QSize>(static_cast<QQuickWidget*>(ptr)->initialSize()).height());
}

void QQuickWidget_KeyPressEvent(void* ptr, void* e){
	static_cast<MyQQuickWidget*>(ptr)->keyPressEvent(static_cast<QKeyEvent*>(e));
}

void QQuickWidget_KeyPressEventDefault(void* ptr, void* e){
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::keyPressEvent(static_cast<QKeyEvent*>(e));
}

void QQuickWidget_KeyReleaseEvent(void* ptr, void* e){
	static_cast<MyQQuickWidget*>(ptr)->keyReleaseEvent(static_cast<QKeyEvent*>(e));
}

void QQuickWidget_KeyReleaseEventDefault(void* ptr, void* e){
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::keyReleaseEvent(static_cast<QKeyEvent*>(e));
}

void QQuickWidget_MouseDoubleClickEvent(void* ptr, void* e){
	static_cast<MyQQuickWidget*>(ptr)->mouseDoubleClickEvent(static_cast<QMouseEvent*>(e));
}

void QQuickWidget_MouseDoubleClickEventDefault(void* ptr, void* e){
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::mouseDoubleClickEvent(static_cast<QMouseEvent*>(e));
}

void QQuickWidget_MouseMoveEvent(void* ptr, void* e){
	static_cast<MyQQuickWidget*>(ptr)->mouseMoveEvent(static_cast<QMouseEvent*>(e));
}

void QQuickWidget_MouseMoveEventDefault(void* ptr, void* e){
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::mouseMoveEvent(static_cast<QMouseEvent*>(e));
}

void QQuickWidget_MousePressEvent(void* ptr, void* e){
	static_cast<MyQQuickWidget*>(ptr)->mousePressEvent(static_cast<QMouseEvent*>(e));
}

void QQuickWidget_MousePressEventDefault(void* ptr, void* e){
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::mousePressEvent(static_cast<QMouseEvent*>(e));
}

void QQuickWidget_MouseReleaseEvent(void* ptr, void* e){
	static_cast<MyQQuickWidget*>(ptr)->mouseReleaseEvent(static_cast<QMouseEvent*>(e));
}

void QQuickWidget_MouseReleaseEventDefault(void* ptr, void* e){
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::mouseReleaseEvent(static_cast<QMouseEvent*>(e));
}

void* QQuickWidget_QuickWindow(void* ptr){
	return static_cast<QQuickWidget*>(ptr)->quickWindow();
}

void* QQuickWidget_RootContext(void* ptr){
	return static_cast<QQuickWidget*>(ptr)->rootContext();
}

void* QQuickWidget_RootObject(void* ptr){
	return static_cast<QQuickWidget*>(ptr)->rootObject();
}

void QQuickWidget_ConnectSceneGraphError(void* ptr){
	QObject::connect(static_cast<QQuickWidget*>(ptr), static_cast<void (QQuickWidget::*)(QQuickWindow::SceneGraphError, const QString &)>(&QQuickWidget::sceneGraphError), static_cast<MyQQuickWidget*>(ptr), static_cast<void (MyQQuickWidget::*)(QQuickWindow::SceneGraphError, const QString &)>(&MyQQuickWidget::Signal_SceneGraphError));;
}

void QQuickWidget_DisconnectSceneGraphError(void* ptr){
	QObject::disconnect(static_cast<QQuickWidget*>(ptr), static_cast<void (QQuickWidget::*)(QQuickWindow::SceneGraphError, const QString &)>(&QQuickWidget::sceneGraphError), static_cast<MyQQuickWidget*>(ptr), static_cast<void (MyQQuickWidget::*)(QQuickWindow::SceneGraphError, const QString &)>(&MyQQuickWidget::Signal_SceneGraphError));;
}

void QQuickWidget_SceneGraphError(void* ptr, int error, char* message){
	static_cast<QQuickWidget*>(ptr)->sceneGraphError(static_cast<QQuickWindow::SceneGraphError>(error), QString(message));
}

void QQuickWidget_SetClearColor(void* ptr, void* color){
	static_cast<QQuickWidget*>(ptr)->setClearColor(*static_cast<QColor*>(color));
}

void QQuickWidget_SetFormat(void* ptr, void* format){
	static_cast<QQuickWidget*>(ptr)->setFormat(*static_cast<QSurfaceFormat*>(format));
}

void QQuickWidget_SetSource(void* ptr, void* url){
	QMetaObject::invokeMethod(static_cast<QQuickWidget*>(ptr), "setSource", Q_ARG(QUrl, *static_cast<QUrl*>(url)));
}

void QQuickWidget_ShowEvent(void* ptr, void* v){
	static_cast<MyQQuickWidget*>(ptr)->showEvent(static_cast<QShowEvent*>(v));
}

void QQuickWidget_ShowEventDefault(void* ptr, void* v){
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::showEvent(static_cast<QShowEvent*>(v));
}

void QQuickWidget_ConnectStatusChanged(void* ptr){
	QObject::connect(static_cast<QQuickWidget*>(ptr), static_cast<void (QQuickWidget::*)(QQuickWidget::Status)>(&QQuickWidget::statusChanged), static_cast<MyQQuickWidget*>(ptr), static_cast<void (MyQQuickWidget::*)(QQuickWidget::Status)>(&MyQQuickWidget::Signal_StatusChanged));;
}

void QQuickWidget_DisconnectStatusChanged(void* ptr){
	QObject::disconnect(static_cast<QQuickWidget*>(ptr), static_cast<void (QQuickWidget::*)(QQuickWidget::Status)>(&QQuickWidget::statusChanged), static_cast<MyQQuickWidget*>(ptr), static_cast<void (MyQQuickWidget::*)(QQuickWidget::Status)>(&MyQQuickWidget::Signal_StatusChanged));;
}

void QQuickWidget_StatusChanged(void* ptr, int status){
	static_cast<QQuickWidget*>(ptr)->statusChanged(static_cast<QQuickWidget::Status>(status));
}

void* QQuickWidget_Source(void* ptr){
	return new QUrl(static_cast<QQuickWidget*>(ptr)->source());
}

void QQuickWidget_WheelEvent(void* ptr, void* e){
	static_cast<MyQQuickWidget*>(ptr)->wheelEvent(static_cast<QWheelEvent*>(e));
}

void QQuickWidget_WheelEventDefault(void* ptr, void* e){
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::wheelEvent(static_cast<QWheelEvent*>(e));
}

void QQuickWidget_DestroyQQuickWidget(void* ptr){
	static_cast<QQuickWidget*>(ptr)->~QQuickWidget();
}

void QQuickWidget_ActionEvent(void* ptr, void* event){
	static_cast<MyQQuickWidget*>(ptr)->actionEvent(static_cast<QActionEvent*>(event));
}

void QQuickWidget_ActionEventDefault(void* ptr, void* event){
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::actionEvent(static_cast<QActionEvent*>(event));
}

void QQuickWidget_EnterEvent(void* ptr, void* event){
	static_cast<MyQQuickWidget*>(ptr)->enterEvent(static_cast<QEvent*>(event));
}

void QQuickWidget_EnterEventDefault(void* ptr, void* event){
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::enterEvent(static_cast<QEvent*>(event));
}

void QQuickWidget_LeaveEvent(void* ptr, void* event){
	static_cast<MyQQuickWidget*>(ptr)->leaveEvent(static_cast<QEvent*>(event));
}

void QQuickWidget_LeaveEventDefault(void* ptr, void* event){
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::leaveEvent(static_cast<QEvent*>(event));
}

void QQuickWidget_MoveEvent(void* ptr, void* event){
	static_cast<MyQQuickWidget*>(ptr)->moveEvent(static_cast<QMoveEvent*>(event));
}

void QQuickWidget_MoveEventDefault(void* ptr, void* event){
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::moveEvent(static_cast<QMoveEvent*>(event));
}

void QQuickWidget_PaintEvent(void* ptr, void* event){
	static_cast<MyQQuickWidget*>(ptr)->paintEvent(static_cast<QPaintEvent*>(event));
}

void QQuickWidget_PaintEventDefault(void* ptr, void* event){
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::paintEvent(static_cast<QPaintEvent*>(event));
}

void QQuickWidget_SetVisible(void* ptr, int visible){
	QMetaObject::invokeMethod(static_cast<MyQQuickWidget*>(ptr), "setVisible", Q_ARG(bool, visible != 0));
}

void QQuickWidget_SetVisibleDefault(void* ptr, int visible){
	QMetaObject::invokeMethod(static_cast<QQuickWidget*>(ptr), "setVisible", Q_ARG(bool, visible != 0));
}

void QQuickWidget_ChangeEvent(void* ptr, void* event){
	static_cast<MyQQuickWidget*>(ptr)->changeEvent(static_cast<QEvent*>(event));
}

void QQuickWidget_ChangeEventDefault(void* ptr, void* event){
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::changeEvent(static_cast<QEvent*>(event));
}

void QQuickWidget_CloseEvent(void* ptr, void* event){
	static_cast<MyQQuickWidget*>(ptr)->closeEvent(static_cast<QCloseEvent*>(event));
}

void QQuickWidget_CloseEventDefault(void* ptr, void* event){
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::closeEvent(static_cast<QCloseEvent*>(event));
}

void QQuickWidget_ContextMenuEvent(void* ptr, void* event){
	static_cast<MyQQuickWidget*>(ptr)->contextMenuEvent(static_cast<QContextMenuEvent*>(event));
}

void QQuickWidget_ContextMenuEventDefault(void* ptr, void* event){
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
}

void QQuickWidget_InitPainter(void* ptr, void* painter){
	static_cast<MyQQuickWidget*>(ptr)->initPainter(static_cast<QPainter*>(painter));
}

void QQuickWidget_InitPainterDefault(void* ptr, void* painter){
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::initPainter(static_cast<QPainter*>(painter));
}

void QQuickWidget_InputMethodEvent(void* ptr, void* event){
	static_cast<MyQQuickWidget*>(ptr)->inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void QQuickWidget_InputMethodEventDefault(void* ptr, void* event){
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void QQuickWidget_ResizeEvent(void* ptr, void* event){
	static_cast<MyQQuickWidget*>(ptr)->resizeEvent(static_cast<QResizeEvent*>(event));
}

void QQuickWidget_ResizeEventDefault(void* ptr, void* event){
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::resizeEvent(static_cast<QResizeEvent*>(event));
}

void QQuickWidget_TabletEvent(void* ptr, void* event){
	static_cast<MyQQuickWidget*>(ptr)->tabletEvent(static_cast<QTabletEvent*>(event));
}

void QQuickWidget_TabletEventDefault(void* ptr, void* event){
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::tabletEvent(static_cast<QTabletEvent*>(event));
}

void QQuickWidget_TimerEvent(void* ptr, void* event){
	static_cast<MyQQuickWidget*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QQuickWidget_TimerEventDefault(void* ptr, void* event){
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::timerEvent(static_cast<QTimerEvent*>(event));
}

void QQuickWidget_ChildEvent(void* ptr, void* event){
	static_cast<MyQQuickWidget*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QQuickWidget_ChildEventDefault(void* ptr, void* event){
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::childEvent(static_cast<QChildEvent*>(event));
}

void QQuickWidget_CustomEvent(void* ptr, void* event){
	static_cast<MyQQuickWidget*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QQuickWidget_CustomEventDefault(void* ptr, void* event){
	static_cast<QQuickWidget*>(ptr)->QQuickWidget::customEvent(static_cast<QEvent*>(event));
}

class MyQQuickWindow: public QQuickWindow {
public:
	MyQQuickWindow(QWindow *parent) : QQuickWindow(parent) {};
	void Signal_ActiveFocusItemChanged() { callbackQQuickWindowActiveFocusItemChanged(this, this->objectName().toUtf8().data()); };
	void Signal_AfterAnimating() { callbackQQuickWindowAfterAnimating(this, this->objectName().toUtf8().data()); };
	void Signal_AfterRendering() { callbackQQuickWindowAfterRendering(this, this->objectName().toUtf8().data()); };
	void Signal_AfterSynchronizing() { callbackQQuickWindowAfterSynchronizing(this, this->objectName().toUtf8().data()); };
	void Signal_BeforeRendering() { callbackQQuickWindowBeforeRendering(this, this->objectName().toUtf8().data()); };
	void Signal_BeforeSynchronizing() { callbackQQuickWindowBeforeSynchronizing(this, this->objectName().toUtf8().data()); };
	void Signal_ColorChanged(const QColor & v) { callbackQQuickWindowColorChanged(this, this->objectName().toUtf8().data(), new QColor(v)); };
	void exposeEvent(QExposeEvent * v) { callbackQQuickWindowExposeEvent(this, this->objectName().toUtf8().data(), v); };
	void focusInEvent(QFocusEvent * ev) { callbackQQuickWindowFocusInEvent(this, this->objectName().toUtf8().data(), ev); };
	void focusOutEvent(QFocusEvent * ev) { callbackQQuickWindowFocusOutEvent(this, this->objectName().toUtf8().data(), ev); };
	void Signal_FrameSwapped() { callbackQQuickWindowFrameSwapped(this, this->objectName().toUtf8().data()); };
	void hideEvent(QHideEvent * v) { callbackQQuickWindowHideEvent(this, this->objectName().toUtf8().data(), v); };
	void keyPressEvent(QKeyEvent * e) { callbackQQuickWindowKeyPressEvent(this, this->objectName().toUtf8().data(), e); };
	void keyReleaseEvent(QKeyEvent * e) { callbackQQuickWindowKeyReleaseEvent(this, this->objectName().toUtf8().data(), e); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQQuickWindowMouseDoubleClickEvent(this, this->objectName().toUtf8().data(), event); };
	void mouseMoveEvent(QMouseEvent * event) { callbackQQuickWindowMouseMoveEvent(this, this->objectName().toUtf8().data(), event); };
	void mousePressEvent(QMouseEvent * event) { callbackQQuickWindowMousePressEvent(this, this->objectName().toUtf8().data(), event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackQQuickWindowMouseReleaseEvent(this, this->objectName().toUtf8().data(), event); };
	void Signal_OpenglContextCreated(QOpenGLContext * context) { callbackQQuickWindowOpenglContextCreated(this, this->objectName().toUtf8().data(), context); };
	void resizeEvent(QResizeEvent * ev) { callbackQQuickWindowResizeEvent(this, this->objectName().toUtf8().data(), ev); };
	void Signal_SceneGraphAboutToStop() { callbackQQuickWindowSceneGraphAboutToStop(this, this->objectName().toUtf8().data()); };
	void Signal_SceneGraphError(QQuickWindow::SceneGraphError error, const QString & message) { callbackQQuickWindowSceneGraphError(this, this->objectName().toUtf8().data(), error, message.toUtf8().data()); };
	void Signal_SceneGraphInitialized() { callbackQQuickWindowSceneGraphInitialized(this, this->objectName().toUtf8().data()); };
	void Signal_SceneGraphInvalidated() { callbackQQuickWindowSceneGraphInvalidated(this, this->objectName().toUtf8().data()); };
	void showEvent(QShowEvent * v) { callbackQQuickWindowShowEvent(this, this->objectName().toUtf8().data(), v); };
	void wheelEvent(QWheelEvent * event) { callbackQQuickWindowWheelEvent(this, this->objectName().toUtf8().data(), event); };
	void moveEvent(QMoveEvent * ev) { callbackQQuickWindowMoveEvent(this, this->objectName().toUtf8().data(), ev); };
	void tabletEvent(QTabletEvent * ev) { callbackQQuickWindowTabletEvent(this, this->objectName().toUtf8().data(), ev); };
	void touchEvent(QTouchEvent * ev) { callbackQQuickWindowTouchEvent(this, this->objectName().toUtf8().data(), ev); };
	void timerEvent(QTimerEvent * event) { callbackQQuickWindowTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQQuickWindowChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQQuickWindowCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void* QQuickWindow_ActiveFocusItem(void* ptr){
	return static_cast<QQuickWindow*>(ptr)->activeFocusItem();
}

void* QQuickWindow_Color(void* ptr){
	return new QColor(static_cast<QQuickWindow*>(ptr)->color());
}

void* QQuickWindow_ContentItem(void* ptr){
	return static_cast<QQuickWindow*>(ptr)->contentItem();
}

void QQuickWindow_SetColor(void* ptr, void* color){
	static_cast<QQuickWindow*>(ptr)->setColor(*static_cast<QColor*>(color));
}

void* QQuickWindow_NewQQuickWindow(void* parent){
	return new MyQQuickWindow(static_cast<QWindow*>(parent));
}

void* QQuickWindow_AccessibleRoot(void* ptr){
	return static_cast<QQuickWindow*>(ptr)->accessibleRoot();
}

void QQuickWindow_ConnectActiveFocusItemChanged(void* ptr){
	QObject::connect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::activeFocusItemChanged), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_ActiveFocusItemChanged));;
}

void QQuickWindow_DisconnectActiveFocusItemChanged(void* ptr){
	QObject::disconnect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::activeFocusItemChanged), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_ActiveFocusItemChanged));;
}

void QQuickWindow_ActiveFocusItemChanged(void* ptr){
	static_cast<QQuickWindow*>(ptr)->activeFocusItemChanged();
}

void QQuickWindow_ConnectAfterAnimating(void* ptr){
	QObject::connect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::afterAnimating), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_AfterAnimating));;
}

void QQuickWindow_DisconnectAfterAnimating(void* ptr){
	QObject::disconnect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::afterAnimating), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_AfterAnimating));;
}

void QQuickWindow_AfterAnimating(void* ptr){
	static_cast<QQuickWindow*>(ptr)->afterAnimating();
}

void QQuickWindow_ConnectAfterRendering(void* ptr){
	QObject::connect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::afterRendering), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_AfterRendering));;
}

void QQuickWindow_DisconnectAfterRendering(void* ptr){
	QObject::disconnect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::afterRendering), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_AfterRendering));;
}

void QQuickWindow_AfterRendering(void* ptr){
	static_cast<QQuickWindow*>(ptr)->afterRendering();
}

void QQuickWindow_ConnectAfterSynchronizing(void* ptr){
	QObject::connect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::afterSynchronizing), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_AfterSynchronizing));;
}

void QQuickWindow_DisconnectAfterSynchronizing(void* ptr){
	QObject::disconnect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::afterSynchronizing), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_AfterSynchronizing));;
}

void QQuickWindow_AfterSynchronizing(void* ptr){
	static_cast<QQuickWindow*>(ptr)->afterSynchronizing();
}

void QQuickWindow_ConnectBeforeRendering(void* ptr){
	QObject::connect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::beforeRendering), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_BeforeRendering));;
}

void QQuickWindow_DisconnectBeforeRendering(void* ptr){
	QObject::disconnect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::beforeRendering), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_BeforeRendering));;
}

void QQuickWindow_BeforeRendering(void* ptr){
	static_cast<QQuickWindow*>(ptr)->beforeRendering();
}

void QQuickWindow_ConnectBeforeSynchronizing(void* ptr){
	QObject::connect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::beforeSynchronizing), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_BeforeSynchronizing));;
}

void QQuickWindow_DisconnectBeforeSynchronizing(void* ptr){
	QObject::disconnect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::beforeSynchronizing), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_BeforeSynchronizing));;
}

void QQuickWindow_BeforeSynchronizing(void* ptr){
	static_cast<QQuickWindow*>(ptr)->beforeSynchronizing();
}

int QQuickWindow_ClearBeforeRendering(void* ptr){
	return static_cast<QQuickWindow*>(ptr)->clearBeforeRendering();
}

void QQuickWindow_ConnectColorChanged(void* ptr){
	QObject::connect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)(const QColor &)>(&QQuickWindow::colorChanged), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)(const QColor &)>(&MyQQuickWindow::Signal_ColorChanged));;
}

void QQuickWindow_DisconnectColorChanged(void* ptr){
	QObject::disconnect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)(const QColor &)>(&QQuickWindow::colorChanged), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)(const QColor &)>(&MyQQuickWindow::Signal_ColorChanged));;
}

void QQuickWindow_ColorChanged(void* ptr, void* v){
	static_cast<QQuickWindow*>(ptr)->colorChanged(*static_cast<QColor*>(v));
}

void* QQuickWindow_CreateTextureFromImage2(void* ptr, void* image){
	return static_cast<QQuickWindow*>(ptr)->createTextureFromImage(*static_cast<QImage*>(image));
}

void* QQuickWindow_CreateTextureFromImage(void* ptr, void* image, int options){
	return static_cast<QQuickWindow*>(ptr)->createTextureFromImage(*static_cast<QImage*>(image), static_cast<QQuickWindow::CreateTextureOption>(options));
}

double QQuickWindow_EffectiveDevicePixelRatio(void* ptr){
	return static_cast<double>(static_cast<QQuickWindow*>(ptr)->effectiveDevicePixelRatio());
}

int QQuickWindow_Event(void* ptr, void* e){
	return static_cast<QQuickWindow*>(ptr)->event(static_cast<QEvent*>(e));
}

void QQuickWindow_ExposeEvent(void* ptr, void* v){
	static_cast<MyQQuickWindow*>(ptr)->exposeEvent(static_cast<QExposeEvent*>(v));
}

void QQuickWindow_ExposeEventDefault(void* ptr, void* v){
	static_cast<QQuickWindow*>(ptr)->QQuickWindow::exposeEvent(static_cast<QExposeEvent*>(v));
}

void QQuickWindow_FocusInEvent(void* ptr, void* ev){
	static_cast<MyQQuickWindow*>(ptr)->focusInEvent(static_cast<QFocusEvent*>(ev));
}

void QQuickWindow_FocusInEventDefault(void* ptr, void* ev){
	static_cast<QQuickWindow*>(ptr)->QQuickWindow::focusInEvent(static_cast<QFocusEvent*>(ev));
}

void QQuickWindow_FocusOutEvent(void* ptr, void* ev){
	static_cast<MyQQuickWindow*>(ptr)->focusOutEvent(static_cast<QFocusEvent*>(ev));
}

void QQuickWindow_FocusOutEventDefault(void* ptr, void* ev){
	static_cast<QQuickWindow*>(ptr)->QQuickWindow::focusOutEvent(static_cast<QFocusEvent*>(ev));
}

void QQuickWindow_ConnectFrameSwapped(void* ptr){
	QObject::connect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::frameSwapped), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_FrameSwapped));;
}

void QQuickWindow_DisconnectFrameSwapped(void* ptr){
	QObject::disconnect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::frameSwapped), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_FrameSwapped));;
}

void QQuickWindow_FrameSwapped(void* ptr){
	static_cast<QQuickWindow*>(ptr)->frameSwapped();
}

void* QQuickWindow_GrabWindow(void* ptr){
	return new QImage(static_cast<QQuickWindow*>(ptr)->grabWindow());
}

int QQuickWindow_QQuickWindow_HasDefaultAlphaBuffer(){
	return QQuickWindow::hasDefaultAlphaBuffer();
}

void QQuickWindow_HideEvent(void* ptr, void* v){
	static_cast<MyQQuickWindow*>(ptr)->hideEvent(static_cast<QHideEvent*>(v));
}

void QQuickWindow_HideEventDefault(void* ptr, void* v){
	static_cast<QQuickWindow*>(ptr)->QQuickWindow::hideEvent(static_cast<QHideEvent*>(v));
}

void* QQuickWindow_IncubationController(void* ptr){
	return static_cast<QQuickWindow*>(ptr)->incubationController();
}

int QQuickWindow_IsPersistentOpenGLContext(void* ptr){
	return static_cast<QQuickWindow*>(ptr)->isPersistentOpenGLContext();
}

int QQuickWindow_IsPersistentSceneGraph(void* ptr){
	return static_cast<QQuickWindow*>(ptr)->isPersistentSceneGraph();
}

int QQuickWindow_IsSceneGraphInitialized(void* ptr){
	return static_cast<QQuickWindow*>(ptr)->isSceneGraphInitialized();
}

void QQuickWindow_KeyPressEvent(void* ptr, void* e){
	static_cast<MyQQuickWindow*>(ptr)->keyPressEvent(static_cast<QKeyEvent*>(e));
}

void QQuickWindow_KeyPressEventDefault(void* ptr, void* e){
	static_cast<QQuickWindow*>(ptr)->QQuickWindow::keyPressEvent(static_cast<QKeyEvent*>(e));
}

void QQuickWindow_KeyReleaseEvent(void* ptr, void* e){
	static_cast<MyQQuickWindow*>(ptr)->keyReleaseEvent(static_cast<QKeyEvent*>(e));
}

void QQuickWindow_KeyReleaseEventDefault(void* ptr, void* e){
	static_cast<QQuickWindow*>(ptr)->QQuickWindow::keyReleaseEvent(static_cast<QKeyEvent*>(e));
}

void QQuickWindow_MouseDoubleClickEvent(void* ptr, void* event){
	static_cast<MyQQuickWindow*>(ptr)->mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QQuickWindow_MouseDoubleClickEventDefault(void* ptr, void* event){
	static_cast<QQuickWindow*>(ptr)->QQuickWindow::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void* QQuickWindow_MouseGrabberItem(void* ptr){
	return static_cast<QQuickWindow*>(ptr)->mouseGrabberItem();
}

void QQuickWindow_MouseMoveEvent(void* ptr, void* event){
	static_cast<MyQQuickWindow*>(ptr)->mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QQuickWindow_MouseMoveEventDefault(void* ptr, void* event){
	static_cast<QQuickWindow*>(ptr)->QQuickWindow::mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QQuickWindow_MousePressEvent(void* ptr, void* event){
	static_cast<MyQQuickWindow*>(ptr)->mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QQuickWindow_MousePressEventDefault(void* ptr, void* event){
	static_cast<QQuickWindow*>(ptr)->QQuickWindow::mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QQuickWindow_MouseReleaseEvent(void* ptr, void* event){
	static_cast<MyQQuickWindow*>(ptr)->mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void QQuickWindow_MouseReleaseEventDefault(void* ptr, void* event){
	static_cast<QQuickWindow*>(ptr)->QQuickWindow::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void* QQuickWindow_OpenglContext(void* ptr){
	return static_cast<QQuickWindow*>(ptr)->openglContext();
}

void QQuickWindow_ConnectOpenglContextCreated(void* ptr){
	QObject::connect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)(QOpenGLContext *)>(&QQuickWindow::openglContextCreated), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)(QOpenGLContext *)>(&MyQQuickWindow::Signal_OpenglContextCreated));;
}

void QQuickWindow_DisconnectOpenglContextCreated(void* ptr){
	QObject::disconnect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)(QOpenGLContext *)>(&QQuickWindow::openglContextCreated), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)(QOpenGLContext *)>(&MyQQuickWindow::Signal_OpenglContextCreated));;
}

void QQuickWindow_OpenglContextCreated(void* ptr, void* context){
	static_cast<QQuickWindow*>(ptr)->openglContextCreated(static_cast<QOpenGLContext*>(context));
}

void QQuickWindow_ReleaseResources(void* ptr){
	QMetaObject::invokeMethod(static_cast<QQuickWindow*>(ptr), "releaseResources");
}

void* QQuickWindow_RenderTarget(void* ptr){
	return static_cast<QQuickWindow*>(ptr)->renderTarget();
}

void* QQuickWindow_RenderTargetSize(void* ptr){
	return new QSize(static_cast<QSize>(static_cast<QQuickWindow*>(ptr)->renderTargetSize()).width(), static_cast<QSize>(static_cast<QQuickWindow*>(ptr)->renderTargetSize()).height());
}

void QQuickWindow_ResetOpenGLState(void* ptr){
	static_cast<QQuickWindow*>(ptr)->resetOpenGLState();
}

void QQuickWindow_ResizeEvent(void* ptr, void* ev){
	static_cast<MyQQuickWindow*>(ptr)->resizeEvent(static_cast<QResizeEvent*>(ev));
}

void QQuickWindow_ResizeEventDefault(void* ptr, void* ev){
	static_cast<QQuickWindow*>(ptr)->QQuickWindow::resizeEvent(static_cast<QResizeEvent*>(ev));
}

void QQuickWindow_ConnectSceneGraphAboutToStop(void* ptr){
	QObject::connect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::sceneGraphAboutToStop), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_SceneGraphAboutToStop));;
}

void QQuickWindow_DisconnectSceneGraphAboutToStop(void* ptr){
	QObject::disconnect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::sceneGraphAboutToStop), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_SceneGraphAboutToStop));;
}

void QQuickWindow_SceneGraphAboutToStop(void* ptr){
	static_cast<QQuickWindow*>(ptr)->sceneGraphAboutToStop();
}

void QQuickWindow_ConnectSceneGraphError(void* ptr){
	QObject::connect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)(QQuickWindow::SceneGraphError, const QString &)>(&QQuickWindow::sceneGraphError), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)(QQuickWindow::SceneGraphError, const QString &)>(&MyQQuickWindow::Signal_SceneGraphError));;
}

void QQuickWindow_DisconnectSceneGraphError(void* ptr){
	QObject::disconnect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)(QQuickWindow::SceneGraphError, const QString &)>(&QQuickWindow::sceneGraphError), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)(QQuickWindow::SceneGraphError, const QString &)>(&MyQQuickWindow::Signal_SceneGraphError));;
}

void QQuickWindow_SceneGraphError(void* ptr, int error, char* message){
	static_cast<QQuickWindow*>(ptr)->sceneGraphError(static_cast<QQuickWindow::SceneGraphError>(error), QString(message));
}

void QQuickWindow_ConnectSceneGraphInitialized(void* ptr){
	QObject::connect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::sceneGraphInitialized), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_SceneGraphInitialized));;
}

void QQuickWindow_DisconnectSceneGraphInitialized(void* ptr){
	QObject::disconnect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::sceneGraphInitialized), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_SceneGraphInitialized));;
}

void QQuickWindow_SceneGraphInitialized(void* ptr){
	static_cast<QQuickWindow*>(ptr)->sceneGraphInitialized();
}

void QQuickWindow_ConnectSceneGraphInvalidated(void* ptr){
	QObject::connect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::sceneGraphInvalidated), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_SceneGraphInvalidated));;
}

void QQuickWindow_DisconnectSceneGraphInvalidated(void* ptr){
	QObject::disconnect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::sceneGraphInvalidated), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_SceneGraphInvalidated));;
}

void QQuickWindow_SceneGraphInvalidated(void* ptr){
	static_cast<QQuickWindow*>(ptr)->sceneGraphInvalidated();
}

void QQuickWindow_ScheduleRenderJob(void* ptr, void* job, int stage){
	static_cast<QQuickWindow*>(ptr)->scheduleRenderJob(static_cast<QRunnable*>(job), static_cast<QQuickWindow::RenderStage>(stage));
}

int QQuickWindow_SendEvent(void* ptr, void* item, void* e){
	return static_cast<QQuickWindow*>(ptr)->sendEvent(static_cast<QQuickItem*>(item), static_cast<QEvent*>(e));
}

void QQuickWindow_SetClearBeforeRendering(void* ptr, int enabled){
	static_cast<QQuickWindow*>(ptr)->setClearBeforeRendering(enabled != 0);
}

void QQuickWindow_QQuickWindow_SetDefaultAlphaBuffer(int useAlpha){
	QQuickWindow::setDefaultAlphaBuffer(useAlpha != 0);
}

void QQuickWindow_SetPersistentOpenGLContext(void* ptr, int persistent){
	static_cast<QQuickWindow*>(ptr)->setPersistentOpenGLContext(persistent != 0);
}

void QQuickWindow_SetPersistentSceneGraph(void* ptr, int persistent){
	static_cast<QQuickWindow*>(ptr)->setPersistentSceneGraph(persistent != 0);
}

void QQuickWindow_SetRenderTarget(void* ptr, void* fbo){
	static_cast<QQuickWindow*>(ptr)->setRenderTarget(static_cast<QOpenGLFramebufferObject*>(fbo));
}

void QQuickWindow_ShowEvent(void* ptr, void* v){
	static_cast<MyQQuickWindow*>(ptr)->showEvent(static_cast<QShowEvent*>(v));
}

void QQuickWindow_ShowEventDefault(void* ptr, void* v){
	static_cast<QQuickWindow*>(ptr)->QQuickWindow::showEvent(static_cast<QShowEvent*>(v));
}

void QQuickWindow_Update(void* ptr){
	QMetaObject::invokeMethod(static_cast<QQuickWindow*>(ptr), "update");
}

void QQuickWindow_WheelEvent(void* ptr, void* event){
	static_cast<MyQQuickWindow*>(ptr)->wheelEvent(static_cast<QWheelEvent*>(event));
}

void QQuickWindow_WheelEventDefault(void* ptr, void* event){
	static_cast<QQuickWindow*>(ptr)->QQuickWindow::wheelEvent(static_cast<QWheelEvent*>(event));
}

void QQuickWindow_DestroyQQuickWindow(void* ptr){
	static_cast<QQuickWindow*>(ptr)->~QQuickWindow();
}

void QQuickWindow_MoveEvent(void* ptr, void* ev){
	static_cast<MyQQuickWindow*>(ptr)->moveEvent(static_cast<QMoveEvent*>(ev));
}

void QQuickWindow_MoveEventDefault(void* ptr, void* ev){
	static_cast<QQuickWindow*>(ptr)->QQuickWindow::moveEvent(static_cast<QMoveEvent*>(ev));
}

void QQuickWindow_TabletEvent(void* ptr, void* ev){
	static_cast<MyQQuickWindow*>(ptr)->tabletEvent(static_cast<QTabletEvent*>(ev));
}

void QQuickWindow_TabletEventDefault(void* ptr, void* ev){
	static_cast<QQuickWindow*>(ptr)->QQuickWindow::tabletEvent(static_cast<QTabletEvent*>(ev));
}

void QQuickWindow_TouchEvent(void* ptr, void* ev){
	static_cast<MyQQuickWindow*>(ptr)->touchEvent(static_cast<QTouchEvent*>(ev));
}

void QQuickWindow_TouchEventDefault(void* ptr, void* ev){
	static_cast<QQuickWindow*>(ptr)->QQuickWindow::touchEvent(static_cast<QTouchEvent*>(ev));
}

void QQuickWindow_TimerEvent(void* ptr, void* event){
	static_cast<MyQQuickWindow*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QQuickWindow_TimerEventDefault(void* ptr, void* event){
	static_cast<QQuickWindow*>(ptr)->QQuickWindow::timerEvent(static_cast<QTimerEvent*>(event));
}

void QQuickWindow_ChildEvent(void* ptr, void* event){
	static_cast<MyQQuickWindow*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QQuickWindow_ChildEventDefault(void* ptr, void* event){
	static_cast<QQuickWindow*>(ptr)->QQuickWindow::childEvent(static_cast<QChildEvent*>(event));
}

void QQuickWindow_CustomEvent(void* ptr, void* event){
	static_cast<MyQQuickWindow*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QQuickWindow_CustomEventDefault(void* ptr, void* event){
	static_cast<QQuickWindow*>(ptr)->QQuickWindow::customEvent(static_cast<QEvent*>(event));
}

class MyQSGAbstractRenderer: public QSGAbstractRenderer {
public:
	void Signal_SceneGraphChanged() { callbackQSGAbstractRendererSceneGraphChanged(this, this->objectName().toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQSGAbstractRendererTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQSGAbstractRendererChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQSGAbstractRendererCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void* QSGAbstractRenderer_ClearColor(void* ptr){
	return new QColor(static_cast<QSGAbstractRenderer*>(ptr)->clearColor());
}

int QSGAbstractRenderer_ClearMode(void* ptr){
	return static_cast<QSGAbstractRenderer*>(ptr)->clearMode();
}

void* QSGAbstractRenderer_DeviceRect(void* ptr){
	return new QRect(static_cast<QRect>(static_cast<QSGAbstractRenderer*>(ptr)->deviceRect()).x(), static_cast<QRect>(static_cast<QSGAbstractRenderer*>(ptr)->deviceRect()).y(), static_cast<QRect>(static_cast<QSGAbstractRenderer*>(ptr)->deviceRect()).width(), static_cast<QRect>(static_cast<QSGAbstractRenderer*>(ptr)->deviceRect()).height());
}

void QSGAbstractRenderer_ConnectSceneGraphChanged(void* ptr){
	QObject::connect(static_cast<QSGAbstractRenderer*>(ptr), static_cast<void (QSGAbstractRenderer::*)()>(&QSGAbstractRenderer::sceneGraphChanged), static_cast<MyQSGAbstractRenderer*>(ptr), static_cast<void (MyQSGAbstractRenderer::*)()>(&MyQSGAbstractRenderer::Signal_SceneGraphChanged));;
}

void QSGAbstractRenderer_DisconnectSceneGraphChanged(void* ptr){
	QObject::disconnect(static_cast<QSGAbstractRenderer*>(ptr), static_cast<void (QSGAbstractRenderer::*)()>(&QSGAbstractRenderer::sceneGraphChanged), static_cast<MyQSGAbstractRenderer*>(ptr), static_cast<void (MyQSGAbstractRenderer::*)()>(&MyQSGAbstractRenderer::Signal_SceneGraphChanged));;
}

void QSGAbstractRenderer_SceneGraphChanged(void* ptr){
	static_cast<QSGAbstractRenderer*>(ptr)->sceneGraphChanged();
}

void QSGAbstractRenderer_SetClearColor(void* ptr, void* color){
	static_cast<QSGAbstractRenderer*>(ptr)->setClearColor(*static_cast<QColor*>(color));
}

void QSGAbstractRenderer_SetClearMode(void* ptr, int mode){
	static_cast<QSGAbstractRenderer*>(ptr)->setClearMode(static_cast<QSGAbstractRenderer::ClearModeBit>(mode));
}

void QSGAbstractRenderer_SetDeviceRect(void* ptr, void* rect){
	static_cast<QSGAbstractRenderer*>(ptr)->setDeviceRect(*static_cast<QRect*>(rect));
}

void QSGAbstractRenderer_SetDeviceRect2(void* ptr, void* size){
	static_cast<QSGAbstractRenderer*>(ptr)->setDeviceRect(*static_cast<QSize*>(size));
}

void QSGAbstractRenderer_SetProjectionMatrix(void* ptr, void* matrix){
	static_cast<QSGAbstractRenderer*>(ptr)->setProjectionMatrix(*static_cast<QMatrix4x4*>(matrix));
}

void QSGAbstractRenderer_SetProjectionMatrixToRect(void* ptr, void* rect){
	static_cast<QSGAbstractRenderer*>(ptr)->setProjectionMatrixToRect(*static_cast<QRectF*>(rect));
}

void QSGAbstractRenderer_SetViewportRect(void* ptr, void* rect){
	static_cast<QSGAbstractRenderer*>(ptr)->setViewportRect(*static_cast<QRect*>(rect));
}

void QSGAbstractRenderer_SetViewportRect2(void* ptr, void* size){
	static_cast<QSGAbstractRenderer*>(ptr)->setViewportRect(*static_cast<QSize*>(size));
}

void* QSGAbstractRenderer_ViewportRect(void* ptr){
	return new QRect(static_cast<QRect>(static_cast<QSGAbstractRenderer*>(ptr)->viewportRect()).x(), static_cast<QRect>(static_cast<QSGAbstractRenderer*>(ptr)->viewportRect()).y(), static_cast<QRect>(static_cast<QSGAbstractRenderer*>(ptr)->viewportRect()).width(), static_cast<QRect>(static_cast<QSGAbstractRenderer*>(ptr)->viewportRect()).height());
}

void QSGAbstractRenderer_TimerEvent(void* ptr, void* event){
	static_cast<MyQSGAbstractRenderer*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QSGAbstractRenderer_TimerEventDefault(void* ptr, void* event){
	static_cast<QSGAbstractRenderer*>(ptr)->QSGAbstractRenderer::timerEvent(static_cast<QTimerEvent*>(event));
}

void QSGAbstractRenderer_ChildEvent(void* ptr, void* event){
	static_cast<MyQSGAbstractRenderer*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QSGAbstractRenderer_ChildEventDefault(void* ptr, void* event){
	static_cast<QSGAbstractRenderer*>(ptr)->QSGAbstractRenderer::childEvent(static_cast<QChildEvent*>(event));
}

void QSGAbstractRenderer_CustomEvent(void* ptr, void* event){
	static_cast<MyQSGAbstractRenderer*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QSGAbstractRenderer_CustomEventDefault(void* ptr, void* event){
	static_cast<QSGAbstractRenderer*>(ptr)->QSGAbstractRenderer::customEvent(static_cast<QEvent*>(event));
}

void* QSGBasicGeometryNode_Geometry2(void* ptr){
	return static_cast<QSGBasicGeometryNode*>(ptr)->geometry();
}

void* QSGBasicGeometryNode_Geometry(void* ptr){
	return const_cast<QSGGeometry*>(static_cast<QSGBasicGeometryNode*>(ptr)->geometry());
}

void QSGBasicGeometryNode_SetGeometry(void* ptr, void* geometry){
	static_cast<QSGBasicGeometryNode*>(ptr)->setGeometry(static_cast<QSGGeometry*>(geometry));
}

void QSGBasicGeometryNode_DestroyQSGBasicGeometryNode(void* ptr){
	static_cast<QSGBasicGeometryNode*>(ptr)->~QSGBasicGeometryNode();
}

void QSGBasicGeometryNode_Preprocess(void* ptr){
	static_cast<QSGBasicGeometryNode*>(ptr)->preprocess();
}

void QSGBasicGeometryNode_PreprocessDefault(void* ptr){
	static_cast<QSGBasicGeometryNode*>(ptr)->QSGBasicGeometryNode::preprocess();
}

void* QSGClipNode_NewQSGClipNode(){
	return new QSGClipNode();
}

void* QSGClipNode_ClipRect(void* ptr){
	return new QRectF(static_cast<QRectF>(static_cast<QSGClipNode*>(ptr)->clipRect()).x(), static_cast<QRectF>(static_cast<QSGClipNode*>(ptr)->clipRect()).y(), static_cast<QRectF>(static_cast<QSGClipNode*>(ptr)->clipRect()).width(), static_cast<QRectF>(static_cast<QSGClipNode*>(ptr)->clipRect()).height());
}

int QSGClipNode_IsRectangular(void* ptr){
	return static_cast<QSGClipNode*>(ptr)->isRectangular();
}

void QSGClipNode_SetClipRect(void* ptr, void* rect){
	static_cast<QSGClipNode*>(ptr)->setClipRect(*static_cast<QRectF*>(rect));
}

void QSGClipNode_SetIsRectangular(void* ptr, int rectHint){
	static_cast<QSGClipNode*>(ptr)->setIsRectangular(rectHint != 0);
}

void QSGClipNode_DestroyQSGClipNode(void* ptr){
	static_cast<QSGClipNode*>(ptr)->~QSGClipNode();
}

void QSGClipNode_Preprocess(void* ptr){
	static_cast<QSGClipNode*>(ptr)->preprocess();
}

void QSGClipNode_PreprocessDefault(void* ptr){
	static_cast<QSGClipNode*>(ptr)->QSGClipNode::preprocess();
}

int QSGDynamicTexture_UpdateTexture(void* ptr){
	return static_cast<QSGDynamicTexture*>(ptr)->updateTexture();
}

void QSGDynamicTexture_TimerEvent(void* ptr, void* event){
	static_cast<QSGDynamicTexture*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QSGDynamicTexture_TimerEventDefault(void* ptr, void* event){
	static_cast<QSGDynamicTexture*>(ptr)->QSGDynamicTexture::timerEvent(static_cast<QTimerEvent*>(event));
}

void QSGDynamicTexture_ChildEvent(void* ptr, void* event){
	static_cast<QSGDynamicTexture*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QSGDynamicTexture_ChildEventDefault(void* ptr, void* event){
	static_cast<QSGDynamicTexture*>(ptr)->QSGDynamicTexture::childEvent(static_cast<QChildEvent*>(event));
}

void QSGDynamicTexture_CustomEvent(void* ptr, void* event){
	static_cast<QSGDynamicTexture*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QSGDynamicTexture_CustomEventDefault(void* ptr, void* event){
	static_cast<QSGDynamicTexture*>(ptr)->QSGDynamicTexture::customEvent(static_cast<QEvent*>(event));
}

void* QSGEngine_NewQSGEngine(void* parent){
	return new QSGEngine(static_cast<QObject*>(parent));
}

void* QSGEngine_CreateRenderer(void* ptr){
	return static_cast<QSGEngine*>(ptr)->createRenderer();
}

void* QSGEngine_CreateTextureFromImage(void* ptr, void* image, int options){
	return static_cast<QSGEngine*>(ptr)->createTextureFromImage(*static_cast<QImage*>(image), static_cast<QSGEngine::CreateTextureOption>(options));
}

void QSGEngine_Initialize(void* ptr, void* context){
	static_cast<QSGEngine*>(ptr)->initialize(static_cast<QOpenGLContext*>(context));
}

void QSGEngine_Invalidate(void* ptr){
	static_cast<QSGEngine*>(ptr)->invalidate();
}

void QSGEngine_DestroyQSGEngine(void* ptr){
	static_cast<QSGEngine*>(ptr)->~QSGEngine();
}

void QSGEngine_TimerEvent(void* ptr, void* event){
	static_cast<QSGEngine*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QSGEngine_TimerEventDefault(void* ptr, void* event){
	static_cast<QSGEngine*>(ptr)->QSGEngine::timerEvent(static_cast<QTimerEvent*>(event));
}

void QSGEngine_ChildEvent(void* ptr, void* event){
	static_cast<QSGEngine*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QSGEngine_ChildEventDefault(void* ptr, void* event){
	static_cast<QSGEngine*>(ptr)->QSGEngine::childEvent(static_cast<QChildEvent*>(event));
}

void QSGEngine_CustomEvent(void* ptr, void* event){
	static_cast<QSGEngine*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QSGEngine_CustomEventDefault(void* ptr, void* event){
	static_cast<QSGEngine*>(ptr)->QSGEngine::customEvent(static_cast<QEvent*>(event));
}

void* QSGFlatColorMaterial_Color(void* ptr){
	return new QColor(static_cast<QSGFlatColorMaterial*>(ptr)->color());
}

void QSGFlatColorMaterial_SetColor(void* ptr, void* color){
	static_cast<QSGFlatColorMaterial*>(ptr)->setColor(*static_cast<QColor*>(color));
}

class MyQSGGeometry: public QSGGeometry {
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
};

void QSGGeometry_Allocate(void* ptr, int vertexCount, int indexCount){
	static_cast<QSGGeometry*>(ptr)->allocate(vertexCount, indexCount);
}

int QSGGeometry_AttributeCount(void* ptr){
	return static_cast<QSGGeometry*>(ptr)->attributeCount();
}

int QSGGeometry_IndexCount(void* ptr){
	return static_cast<QSGGeometry*>(ptr)->indexCount();
}

void* QSGGeometry_IndexData(void* ptr){
	return static_cast<QSGGeometry*>(ptr)->indexData();
}

void* QSGGeometry_IndexData2(void* ptr){
	return const_cast<void*>(static_cast<QSGGeometry*>(ptr)->indexData());
}

int QSGGeometry_IndexDataPattern(void* ptr){
	return static_cast<QSGGeometry*>(ptr)->indexDataPattern();
}

int QSGGeometry_IndexType(void* ptr){
	return static_cast<QSGGeometry*>(ptr)->indexType();
}

void QSGGeometry_MarkIndexDataDirty(void* ptr){
	static_cast<QSGGeometry*>(ptr)->markIndexDataDirty();
}

void QSGGeometry_MarkVertexDataDirty(void* ptr){
	static_cast<QSGGeometry*>(ptr)->markVertexDataDirty();
}

void QSGGeometry_SetIndexDataPattern(void* ptr, int p){
	static_cast<QSGGeometry*>(ptr)->setIndexDataPattern(static_cast<QSGGeometry::DataPattern>(p));
}

void QSGGeometry_SetVertexDataPattern(void* ptr, int p){
	static_cast<QSGGeometry*>(ptr)->setVertexDataPattern(static_cast<QSGGeometry::DataPattern>(p));
}

int QSGGeometry_SizeOfIndex(void* ptr){
	return static_cast<QSGGeometry*>(ptr)->sizeOfIndex();
}

int QSGGeometry_SizeOfVertex(void* ptr){
	return static_cast<QSGGeometry*>(ptr)->sizeOfVertex();
}

void QSGGeometry_QSGGeometry_UpdateRectGeometry(void* g, void* rect){
	QSGGeometry::updateRectGeometry(static_cast<QSGGeometry*>(g), *static_cast<QRectF*>(rect));
}

void QSGGeometry_QSGGeometry_UpdateTexturedRectGeometry(void* g, void* rect, void* textureRect){
	QSGGeometry::updateTexturedRectGeometry(static_cast<QSGGeometry*>(g), *static_cast<QRectF*>(rect), *static_cast<QRectF*>(textureRect));
}

int QSGGeometry_VertexCount(void* ptr){
	return static_cast<QSGGeometry*>(ptr)->vertexCount();
}

void* QSGGeometry_VertexData(void* ptr){
	return static_cast<QSGGeometry*>(ptr)->vertexData();
}

void* QSGGeometry_VertexData2(void* ptr){
	return const_cast<void*>(static_cast<QSGGeometry*>(ptr)->vertexData());
}

int QSGGeometry_VertexDataPattern(void* ptr){
	return static_cast<QSGGeometry*>(ptr)->vertexDataPattern();
}

void QSGGeometry_DestroyQSGGeometry(void* ptr){
	static_cast<QSGGeometry*>(ptr)->~QSGGeometry();
}

char* QSGGeometry_ObjectNameAbs(void* ptr){
	if (dynamic_cast<MyQSGGeometry*>(static_cast<QSGGeometry*>(ptr))) {
		return static_cast<MyQSGGeometry*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QSGGeometry_BASE").toUtf8().data();
}

void QSGGeometry_SetObjectNameAbs(void* ptr, char* name){
	if (dynamic_cast<MyQSGGeometry*>(static_cast<QSGGeometry*>(ptr))) {
		static_cast<MyQSGGeometry*>(ptr)->setObjectNameAbs(QString(name));
	}
}

void* QSGGeometryNode_NewQSGGeometryNode(){
	return new QSGGeometryNode();
}

void* QSGGeometryNode_Material(void* ptr){
	return static_cast<QSGGeometryNode*>(ptr)->material();
}

void* QSGGeometryNode_OpaqueMaterial(void* ptr){
	return static_cast<QSGGeometryNode*>(ptr)->opaqueMaterial();
}

void QSGGeometryNode_SetMaterial(void* ptr, void* material){
	static_cast<QSGGeometryNode*>(ptr)->setMaterial(static_cast<QSGMaterial*>(material));
}

void QSGGeometryNode_SetOpaqueMaterial(void* ptr, void* material){
	static_cast<QSGGeometryNode*>(ptr)->setOpaqueMaterial(static_cast<QSGMaterial*>(material));
}

void QSGGeometryNode_DestroyQSGGeometryNode(void* ptr){
	static_cast<QSGGeometryNode*>(ptr)->~QSGGeometryNode();
}

void QSGGeometryNode_Preprocess(void* ptr){
	static_cast<QSGGeometryNode*>(ptr)->preprocess();
}

void QSGGeometryNode_PreprocessDefault(void* ptr){
	static_cast<QSGGeometryNode*>(ptr)->QSGGeometryNode::preprocess();
}

class MyQSGMaterial: public QSGMaterial {
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
};

int QSGMaterial_Compare(void* ptr, void* other){
	return static_cast<QSGMaterial*>(ptr)->compare(static_cast<QSGMaterial*>(other));
}

void* QSGMaterial_CreateShader(void* ptr){
	return static_cast<QSGMaterial*>(ptr)->createShader();
}

int QSGMaterial_Flags(void* ptr){
	return static_cast<QSGMaterial*>(ptr)->flags();
}

void QSGMaterial_SetFlag(void* ptr, int flags, int on){
	static_cast<QSGMaterial*>(ptr)->setFlag(static_cast<QSGMaterial::Flag>(flags), on != 0);
}

void* QSGMaterial_Type(void* ptr){
	return static_cast<QSGMaterial*>(ptr)->type();
}

char* QSGMaterial_ObjectNameAbs(void* ptr){
	if (dynamic_cast<MyQSGMaterial*>(static_cast<QSGMaterial*>(ptr))) {
		return static_cast<MyQSGMaterial*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QSGMaterial_BASE").toUtf8().data();
}

void QSGMaterial_SetObjectNameAbs(void* ptr, char* name){
	if (dynamic_cast<MyQSGMaterial*>(static_cast<QSGMaterial*>(ptr))) {
		static_cast<MyQSGMaterial*>(ptr)->setObjectNameAbs(QString(name));
	}
}

class MyQSGMaterialShader: public QSGMaterialShader {
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	void activate() { callbackQSGMaterialShaderActivate(this, this->objectNameAbs().toUtf8().data()); };
	void compile() { callbackQSGMaterialShaderCompile(this, this->objectNameAbs().toUtf8().data()); };
	void deactivate() { callbackQSGMaterialShaderDeactivate(this, this->objectNameAbs().toUtf8().data()); };
	void initialize() { callbackQSGMaterialShaderInitialize(this, this->objectNameAbs().toUtf8().data()); };
};

void QSGMaterialShader_Activate(void* ptr){
	static_cast<MyQSGMaterialShader*>(ptr)->activate();
}

void QSGMaterialShader_ActivateDefault(void* ptr){
	static_cast<QSGMaterialShader*>(ptr)->QSGMaterialShader::activate();
}

void QSGMaterialShader_Compile(void* ptr){
	static_cast<MyQSGMaterialShader*>(ptr)->compile();
}

void QSGMaterialShader_CompileDefault(void* ptr){
	static_cast<QSGMaterialShader*>(ptr)->QSGMaterialShader::compile();
}

void QSGMaterialShader_Deactivate(void* ptr){
	static_cast<MyQSGMaterialShader*>(ptr)->deactivate();
}

void QSGMaterialShader_DeactivateDefault(void* ptr){
	static_cast<QSGMaterialShader*>(ptr)->QSGMaterialShader::deactivate();
}

void QSGMaterialShader_Initialize(void* ptr){
	static_cast<MyQSGMaterialShader*>(ptr)->initialize();
}

void QSGMaterialShader_InitializeDefault(void* ptr){
	static_cast<QSGMaterialShader*>(ptr)->QSGMaterialShader::initialize();
}

void* QSGMaterialShader_Program(void* ptr){
	return static_cast<QSGMaterialShader*>(ptr)->program();
}

char* QSGMaterialShader_ObjectNameAbs(void* ptr){
	if (dynamic_cast<MyQSGMaterialShader*>(static_cast<QSGMaterialShader*>(ptr))) {
		return static_cast<MyQSGMaterialShader*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QSGMaterialShader_BASE").toUtf8().data();
}

void QSGMaterialShader_SetObjectNameAbs(void* ptr, char* name){
	if (dynamic_cast<MyQSGMaterialShader*>(static_cast<QSGMaterialShader*>(ptr))) {
		static_cast<MyQSGMaterialShader*>(ptr)->setObjectNameAbs(QString(name));
	}
}

class MyQSGNode: public QSGNode {
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	MyQSGNode() : QSGNode() {};
	void preprocess() { callbackQSGNodePreprocess(this, this->objectNameAbs().toUtf8().data()); };
};

void* QSGNode_ChildAtIndex(void* ptr, int i){
	return static_cast<QSGNode*>(ptr)->childAtIndex(i);
}

int QSGNode_ChildCount(void* ptr){
	return static_cast<QSGNode*>(ptr)->childCount();
}

void* QSGNode_NewQSGNode(){
	return new MyQSGNode();
}

void QSGNode_AppendChildNode(void* ptr, void* node){
	static_cast<QSGNode*>(ptr)->appendChildNode(static_cast<QSGNode*>(node));
}

void* QSGNode_FirstChild(void* ptr){
	return static_cast<QSGNode*>(ptr)->firstChild();
}

int QSGNode_Flags(void* ptr){
	return static_cast<QSGNode*>(ptr)->flags();
}

void QSGNode_InsertChildNodeAfter(void* ptr, void* node, void* after){
	static_cast<QSGNode*>(ptr)->insertChildNodeAfter(static_cast<QSGNode*>(node), static_cast<QSGNode*>(after));
}

void QSGNode_InsertChildNodeBefore(void* ptr, void* node, void* before){
	static_cast<QSGNode*>(ptr)->insertChildNodeBefore(static_cast<QSGNode*>(node), static_cast<QSGNode*>(before));
}

int QSGNode_IsSubtreeBlocked(void* ptr){
	return static_cast<QSGNode*>(ptr)->isSubtreeBlocked();
}

void* QSGNode_LastChild(void* ptr){
	return static_cast<QSGNode*>(ptr)->lastChild();
}

void QSGNode_MarkDirty(void* ptr, int bits){
	static_cast<QSGNode*>(ptr)->markDirty(static_cast<QSGNode::DirtyStateBit>(bits));
}

void* QSGNode_NextSibling(void* ptr){
	return static_cast<QSGNode*>(ptr)->nextSibling();
}

void* QSGNode_Parent(void* ptr){
	return static_cast<QSGNode*>(ptr)->parent();
}

void QSGNode_PrependChildNode(void* ptr, void* node){
	static_cast<QSGNode*>(ptr)->prependChildNode(static_cast<QSGNode*>(node));
}

void QSGNode_Preprocess(void* ptr){
	static_cast<MyQSGNode*>(ptr)->preprocess();
}

void QSGNode_PreprocessDefault(void* ptr){
	static_cast<QSGNode*>(ptr)->QSGNode::preprocess();
}

void* QSGNode_PreviousSibling(void* ptr){
	return static_cast<QSGNode*>(ptr)->previousSibling();
}

void QSGNode_RemoveAllChildNodes(void* ptr){
	static_cast<QSGNode*>(ptr)->removeAllChildNodes();
}

void QSGNode_RemoveChildNode(void* ptr, void* node){
	static_cast<QSGNode*>(ptr)->removeChildNode(static_cast<QSGNode*>(node));
}

void QSGNode_SetFlag(void* ptr, int f, int enabled){
	static_cast<QSGNode*>(ptr)->setFlag(static_cast<QSGNode::Flag>(f), enabled != 0);
}

void QSGNode_SetFlags(void* ptr, int f, int enabled){
	static_cast<QSGNode*>(ptr)->setFlags(static_cast<QSGNode::Flag>(f), enabled != 0);
}

int QSGNode_Type(void* ptr){
	return static_cast<QSGNode*>(ptr)->type();
}

void QSGNode_DestroyQSGNode(void* ptr){
	static_cast<QSGNode*>(ptr)->~QSGNode();
}

char* QSGNode_ObjectNameAbs(void* ptr){
	if (dynamic_cast<MyQSGNode*>(static_cast<QSGNode*>(ptr))) {
		return static_cast<MyQSGNode*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QSGNode_BASE").toUtf8().data();
}

void QSGNode_SetObjectNameAbs(void* ptr, char* name){
	if (dynamic_cast<MyQSGNode*>(static_cast<QSGNode*>(ptr))) {
		static_cast<MyQSGNode*>(ptr)->setObjectNameAbs(QString(name));
	}
}

void* QSGOpacityNode_NewQSGOpacityNode(){
	return new QSGOpacityNode();
}

double QSGOpacityNode_Opacity(void* ptr){
	return static_cast<double>(static_cast<QSGOpacityNode*>(ptr)->opacity());
}

void QSGOpacityNode_SetOpacity(void* ptr, double opacity){
	static_cast<QSGOpacityNode*>(ptr)->setOpacity(static_cast<double>(opacity));
}

void QSGOpacityNode_DestroyQSGOpacityNode(void* ptr){
	static_cast<QSGOpacityNode*>(ptr)->~QSGOpacityNode();
}

void QSGOpacityNode_Preprocess(void* ptr){
	static_cast<QSGOpacityNode*>(ptr)->preprocess();
}

void QSGOpacityNode_PreprocessDefault(void* ptr){
	static_cast<QSGOpacityNode*>(ptr)->QSGOpacityNode::preprocess();
}

int QSGOpaqueTextureMaterial_Filtering(void* ptr){
	return static_cast<QSGOpaqueTextureMaterial*>(ptr)->filtering();
}

int QSGOpaqueTextureMaterial_HorizontalWrapMode(void* ptr){
	return static_cast<QSGOpaqueTextureMaterial*>(ptr)->horizontalWrapMode();
}

int QSGOpaqueTextureMaterial_MipmapFiltering(void* ptr){
	return static_cast<QSGOpaqueTextureMaterial*>(ptr)->mipmapFiltering();
}

void QSGOpaqueTextureMaterial_SetFiltering(void* ptr, int filtering){
	static_cast<QSGOpaqueTextureMaterial*>(ptr)->setFiltering(static_cast<QSGTexture::Filtering>(filtering));
}

void QSGOpaqueTextureMaterial_SetHorizontalWrapMode(void* ptr, int mode){
	static_cast<QSGOpaqueTextureMaterial*>(ptr)->setHorizontalWrapMode(static_cast<QSGTexture::WrapMode>(mode));
}

void QSGOpaqueTextureMaterial_SetMipmapFiltering(void* ptr, int filtering){
	static_cast<QSGOpaqueTextureMaterial*>(ptr)->setMipmapFiltering(static_cast<QSGTexture::Filtering>(filtering));
}

void QSGOpaqueTextureMaterial_SetTexture(void* ptr, void* texture){
	static_cast<QSGOpaqueTextureMaterial*>(ptr)->setTexture(static_cast<QSGTexture*>(texture));
}

void QSGOpaqueTextureMaterial_SetVerticalWrapMode(void* ptr, int mode){
	static_cast<QSGOpaqueTextureMaterial*>(ptr)->setVerticalWrapMode(static_cast<QSGTexture::WrapMode>(mode));
}

void* QSGOpaqueTextureMaterial_Texture(void* ptr){
	return static_cast<QSGOpaqueTextureMaterial*>(ptr)->texture();
}

int QSGOpaqueTextureMaterial_VerticalWrapMode(void* ptr){
	return static_cast<QSGOpaqueTextureMaterial*>(ptr)->verticalWrapMode();
}

void* QSGSimpleRectNode_NewQSGSimpleRectNode2(){
	return new QSGSimpleRectNode();
}

void* QSGSimpleRectNode_NewQSGSimpleRectNode(void* rect, void* color){
	return new QSGSimpleRectNode(*static_cast<QRectF*>(rect), *static_cast<QColor*>(color));
}

void* QSGSimpleRectNode_Color(void* ptr){
	return new QColor(static_cast<QSGSimpleRectNode*>(ptr)->color());
}

void* QSGSimpleRectNode_Rect(void* ptr){
	return new QRectF(static_cast<QRectF>(static_cast<QSGSimpleRectNode*>(ptr)->rect()).x(), static_cast<QRectF>(static_cast<QSGSimpleRectNode*>(ptr)->rect()).y(), static_cast<QRectF>(static_cast<QSGSimpleRectNode*>(ptr)->rect()).width(), static_cast<QRectF>(static_cast<QSGSimpleRectNode*>(ptr)->rect()).height());
}

void QSGSimpleRectNode_SetColor(void* ptr, void* color){
	static_cast<QSGSimpleRectNode*>(ptr)->setColor(*static_cast<QColor*>(color));
}

void QSGSimpleRectNode_SetRect(void* ptr, void* rect){
	static_cast<QSGSimpleRectNode*>(ptr)->setRect(*static_cast<QRectF*>(rect));
}

void QSGSimpleRectNode_SetRect2(void* ptr, double x, double y, double w, double h){
	static_cast<QSGSimpleRectNode*>(ptr)->setRect(static_cast<double>(x), static_cast<double>(y), static_cast<double>(w), static_cast<double>(h));
}

void QSGSimpleRectNode_Preprocess(void* ptr){
	static_cast<QSGSimpleRectNode*>(ptr)->preprocess();
}

void QSGSimpleRectNode_PreprocessDefault(void* ptr){
	static_cast<QSGSimpleRectNode*>(ptr)->QSGSimpleRectNode::preprocess();
}

void* QSGSimpleTextureNode_NewQSGSimpleTextureNode(){
	return new QSGSimpleTextureNode();
}

int QSGSimpleTextureNode_Filtering(void* ptr){
	return static_cast<QSGSimpleTextureNode*>(ptr)->filtering();
}

int QSGSimpleTextureNode_OwnsTexture(void* ptr){
	return static_cast<QSGSimpleTextureNode*>(ptr)->ownsTexture();
}

void* QSGSimpleTextureNode_Rect(void* ptr){
	return new QRectF(static_cast<QRectF>(static_cast<QSGSimpleTextureNode*>(ptr)->rect()).x(), static_cast<QRectF>(static_cast<QSGSimpleTextureNode*>(ptr)->rect()).y(), static_cast<QRectF>(static_cast<QSGSimpleTextureNode*>(ptr)->rect()).width(), static_cast<QRectF>(static_cast<QSGSimpleTextureNode*>(ptr)->rect()).height());
}

void QSGSimpleTextureNode_SetFiltering(void* ptr, int filtering){
	static_cast<QSGSimpleTextureNode*>(ptr)->setFiltering(static_cast<QSGTexture::Filtering>(filtering));
}

void QSGSimpleTextureNode_SetOwnsTexture(void* ptr, int owns){
	static_cast<QSGSimpleTextureNode*>(ptr)->setOwnsTexture(owns != 0);
}

void QSGSimpleTextureNode_SetRect(void* ptr, void* r){
	static_cast<QSGSimpleTextureNode*>(ptr)->setRect(*static_cast<QRectF*>(r));
}

void QSGSimpleTextureNode_SetRect2(void* ptr, double x, double y, double w, double h){
	static_cast<QSGSimpleTextureNode*>(ptr)->setRect(static_cast<double>(x), static_cast<double>(y), static_cast<double>(w), static_cast<double>(h));
}

void QSGSimpleTextureNode_SetSourceRect(void* ptr, void* r){
	static_cast<QSGSimpleTextureNode*>(ptr)->setSourceRect(*static_cast<QRectF*>(r));
}

void QSGSimpleTextureNode_SetSourceRect2(void* ptr, double x, double y, double w, double h){
	static_cast<QSGSimpleTextureNode*>(ptr)->setSourceRect(static_cast<double>(x), static_cast<double>(y), static_cast<double>(w), static_cast<double>(h));
}

void QSGSimpleTextureNode_SetTexture(void* ptr, void* texture){
	static_cast<QSGSimpleTextureNode*>(ptr)->setTexture(static_cast<QSGTexture*>(texture));
}

void QSGSimpleTextureNode_SetTextureCoordinatesTransform(void* ptr, int mode){
	static_cast<QSGSimpleTextureNode*>(ptr)->setTextureCoordinatesTransform(static_cast<QSGSimpleTextureNode::TextureCoordinatesTransformFlag>(mode));
}

void* QSGSimpleTextureNode_SourceRect(void* ptr){
	return new QRectF(static_cast<QRectF>(static_cast<QSGSimpleTextureNode*>(ptr)->sourceRect()).x(), static_cast<QRectF>(static_cast<QSGSimpleTextureNode*>(ptr)->sourceRect()).y(), static_cast<QRectF>(static_cast<QSGSimpleTextureNode*>(ptr)->sourceRect()).width(), static_cast<QRectF>(static_cast<QSGSimpleTextureNode*>(ptr)->sourceRect()).height());
}

void* QSGSimpleTextureNode_Texture(void* ptr){
	return static_cast<QSGSimpleTextureNode*>(ptr)->texture();
}

int QSGSimpleTextureNode_TextureCoordinatesTransform(void* ptr){
	return static_cast<QSGSimpleTextureNode*>(ptr)->textureCoordinatesTransform();
}

void QSGSimpleTextureNode_DestroyQSGSimpleTextureNode(void* ptr){
	static_cast<QSGSimpleTextureNode*>(ptr)->~QSGSimpleTextureNode();
}

void QSGSimpleTextureNode_Preprocess(void* ptr){
	static_cast<QSGSimpleTextureNode*>(ptr)->preprocess();
}

void QSGSimpleTextureNode_PreprocessDefault(void* ptr){
	static_cast<QSGSimpleTextureNode*>(ptr)->QSGSimpleTextureNode::preprocess();
}

class MyQSGTexture: public QSGTexture {
public:
	void timerEvent(QTimerEvent * event) { callbackQSGTextureTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQSGTextureChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQSGTextureCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void QSGTexture_Bind(void* ptr){
	static_cast<QSGTexture*>(ptr)->bind();
}

void* QSGTexture_ConvertToNormalizedSourceRect(void* ptr, void* rect){
	return new QRectF(static_cast<QRectF>(static_cast<QSGTexture*>(ptr)->convertToNormalizedSourceRect(*static_cast<QRectF*>(rect))).x(), static_cast<QRectF>(static_cast<QSGTexture*>(ptr)->convertToNormalizedSourceRect(*static_cast<QRectF*>(rect))).y(), static_cast<QRectF>(static_cast<QSGTexture*>(ptr)->convertToNormalizedSourceRect(*static_cast<QRectF*>(rect))).width(), static_cast<QRectF>(static_cast<QSGTexture*>(ptr)->convertToNormalizedSourceRect(*static_cast<QRectF*>(rect))).height());
}

int QSGTexture_Filtering(void* ptr){
	return static_cast<QSGTexture*>(ptr)->filtering();
}

int QSGTexture_HasAlphaChannel(void* ptr){
	return static_cast<QSGTexture*>(ptr)->hasAlphaChannel();
}

int QSGTexture_HasMipmaps(void* ptr){
	return static_cast<QSGTexture*>(ptr)->hasMipmaps();
}

int QSGTexture_HorizontalWrapMode(void* ptr){
	return static_cast<QSGTexture*>(ptr)->horizontalWrapMode();
}

int QSGTexture_IsAtlasTexture(void* ptr){
	return static_cast<QSGTexture*>(ptr)->isAtlasTexture();
}

int QSGTexture_MipmapFiltering(void* ptr){
	return static_cast<QSGTexture*>(ptr)->mipmapFiltering();
}

void* QSGTexture_NormalizedTextureSubRect(void* ptr){
	return new QRectF(static_cast<QRectF>(static_cast<QSGTexture*>(ptr)->normalizedTextureSubRect()).x(), static_cast<QRectF>(static_cast<QSGTexture*>(ptr)->normalizedTextureSubRect()).y(), static_cast<QRectF>(static_cast<QSGTexture*>(ptr)->normalizedTextureSubRect()).width(), static_cast<QRectF>(static_cast<QSGTexture*>(ptr)->normalizedTextureSubRect()).height());
}

void* QSGTexture_RemovedFromAtlas(void* ptr){
	return static_cast<QSGTexture*>(ptr)->removedFromAtlas();
}

void QSGTexture_SetFiltering(void* ptr, int filter){
	static_cast<QSGTexture*>(ptr)->setFiltering(static_cast<QSGTexture::Filtering>(filter));
}

void QSGTexture_SetHorizontalWrapMode(void* ptr, int hwrap){
	static_cast<QSGTexture*>(ptr)->setHorizontalWrapMode(static_cast<QSGTexture::WrapMode>(hwrap));
}

void QSGTexture_SetMipmapFiltering(void* ptr, int filter){
	static_cast<QSGTexture*>(ptr)->setMipmapFiltering(static_cast<QSGTexture::Filtering>(filter));
}

void QSGTexture_SetVerticalWrapMode(void* ptr, int vwrap){
	static_cast<QSGTexture*>(ptr)->setVerticalWrapMode(static_cast<QSGTexture::WrapMode>(vwrap));
}

int QSGTexture_TextureId(void* ptr){
	return static_cast<QSGTexture*>(ptr)->textureId();
}

void* QSGTexture_TextureSize(void* ptr){
	return new QSize(static_cast<QSize>(static_cast<QSGTexture*>(ptr)->textureSize()).width(), static_cast<QSize>(static_cast<QSGTexture*>(ptr)->textureSize()).height());
}

void QSGTexture_UpdateBindOptions(void* ptr, int force){
	static_cast<QSGTexture*>(ptr)->updateBindOptions(force != 0);
}

int QSGTexture_VerticalWrapMode(void* ptr){
	return static_cast<QSGTexture*>(ptr)->verticalWrapMode();
}

void QSGTexture_DestroyQSGTexture(void* ptr){
	static_cast<QSGTexture*>(ptr)->~QSGTexture();
}

void QSGTexture_TimerEvent(void* ptr, void* event){
	static_cast<MyQSGTexture*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QSGTexture_TimerEventDefault(void* ptr, void* event){
	static_cast<QSGTexture*>(ptr)->QSGTexture::timerEvent(static_cast<QTimerEvent*>(event));
}

void QSGTexture_ChildEvent(void* ptr, void* event){
	static_cast<MyQSGTexture*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QSGTexture_ChildEventDefault(void* ptr, void* event){
	static_cast<QSGTexture*>(ptr)->QSGTexture::childEvent(static_cast<QChildEvent*>(event));
}

void QSGTexture_CustomEvent(void* ptr, void* event){
	static_cast<MyQSGTexture*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QSGTexture_CustomEventDefault(void* ptr, void* event){
	static_cast<QSGTexture*>(ptr)->QSGTexture::customEvent(static_cast<QEvent*>(event));
}

class MyQSGTextureProvider: public QSGTextureProvider {
public:
	void Signal_TextureChanged() { callbackQSGTextureProviderTextureChanged(this, this->objectName().toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQSGTextureProviderTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQSGTextureProviderChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQSGTextureProviderCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void* QSGTextureProvider_Texture(void* ptr){
	return static_cast<QSGTextureProvider*>(ptr)->texture();
}

void QSGTextureProvider_ConnectTextureChanged(void* ptr){
	QObject::connect(static_cast<QSGTextureProvider*>(ptr), static_cast<void (QSGTextureProvider::*)()>(&QSGTextureProvider::textureChanged), static_cast<MyQSGTextureProvider*>(ptr), static_cast<void (MyQSGTextureProvider::*)()>(&MyQSGTextureProvider::Signal_TextureChanged));;
}

void QSGTextureProvider_DisconnectTextureChanged(void* ptr){
	QObject::disconnect(static_cast<QSGTextureProvider*>(ptr), static_cast<void (QSGTextureProvider::*)()>(&QSGTextureProvider::textureChanged), static_cast<MyQSGTextureProvider*>(ptr), static_cast<void (MyQSGTextureProvider::*)()>(&MyQSGTextureProvider::Signal_TextureChanged));;
}

void QSGTextureProvider_TextureChanged(void* ptr){
	static_cast<QSGTextureProvider*>(ptr)->textureChanged();
}

void QSGTextureProvider_TimerEvent(void* ptr, void* event){
	static_cast<MyQSGTextureProvider*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QSGTextureProvider_TimerEventDefault(void* ptr, void* event){
	static_cast<QSGTextureProvider*>(ptr)->QSGTextureProvider::timerEvent(static_cast<QTimerEvent*>(event));
}

void QSGTextureProvider_ChildEvent(void* ptr, void* event){
	static_cast<MyQSGTextureProvider*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QSGTextureProvider_ChildEventDefault(void* ptr, void* event){
	static_cast<QSGTextureProvider*>(ptr)->QSGTextureProvider::childEvent(static_cast<QChildEvent*>(event));
}

void QSGTextureProvider_CustomEvent(void* ptr, void* event){
	static_cast<MyQSGTextureProvider*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QSGTextureProvider_CustomEventDefault(void* ptr, void* event){
	static_cast<QSGTextureProvider*>(ptr)->QSGTextureProvider::customEvent(static_cast<QEvent*>(event));
}

void* QSGTransformNode_NewQSGTransformNode(){
	return new QSGTransformNode();
}

void QSGTransformNode_SetMatrix(void* ptr, void* matrix){
	static_cast<QSGTransformNode*>(ptr)->setMatrix(*static_cast<QMatrix4x4*>(matrix));
}

void QSGTransformNode_DestroyQSGTransformNode(void* ptr){
	static_cast<QSGTransformNode*>(ptr)->~QSGTransformNode();
}

void QSGTransformNode_Preprocess(void* ptr){
	static_cast<QSGTransformNode*>(ptr)->preprocess();
}

void QSGTransformNode_PreprocessDefault(void* ptr){
	static_cast<QSGTransformNode*>(ptr)->QSGTransformNode::preprocess();
}

