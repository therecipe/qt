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
	void releaseResources() { if (!callbackQQuickFramebufferObjectReleaseResources(this->objectName().toUtf8().data())) { QQuickFramebufferObject::releaseResources(); }; };
	void Signal_TextureFollowsItemSizeChanged(bool v) { callbackQQuickFramebufferObjectTextureFollowsItemSizeChanged(this->objectName().toUtf8().data(), v); };
protected:
	void classBegin() { if (!callbackQQuickFramebufferObjectClassBegin(this->objectName().toUtf8().data())) { QQuickFramebufferObject::classBegin(); }; };
	void componentComplete() { if (!callbackQQuickFramebufferObjectComponentComplete(this->objectName().toUtf8().data())) { QQuickFramebufferObject::componentComplete(); }; };
	void dragEnterEvent(QDragEnterEvent * event) { if (!callbackQQuickFramebufferObjectDragEnterEvent(this->objectName().toUtf8().data(), event)) { QQuickFramebufferObject::dragEnterEvent(event); }; };
	void dragLeaveEvent(QDragLeaveEvent * event) { if (!callbackQQuickFramebufferObjectDragLeaveEvent(this->objectName().toUtf8().data(), event)) { QQuickFramebufferObject::dragLeaveEvent(event); }; };
	void dragMoveEvent(QDragMoveEvent * event) { if (!callbackQQuickFramebufferObjectDragMoveEvent(this->objectName().toUtf8().data(), event)) { QQuickFramebufferObject::dragMoveEvent(event); }; };
	void dropEvent(QDropEvent * event) { if (!callbackQQuickFramebufferObjectDropEvent(this->objectName().toUtf8().data(), event)) { QQuickFramebufferObject::dropEvent(event); }; };
	void focusInEvent(QFocusEvent * event) { if (!callbackQQuickFramebufferObjectFocusInEvent(this->objectName().toUtf8().data(), event)) { QQuickFramebufferObject::focusInEvent(event); }; };
	void focusOutEvent(QFocusEvent * event) { if (!callbackQQuickFramebufferObjectFocusOutEvent(this->objectName().toUtf8().data(), event)) { QQuickFramebufferObject::focusOutEvent(event); }; };
	void hoverEnterEvent(QHoverEvent * event) { if (!callbackQQuickFramebufferObjectHoverEnterEvent(this->objectName().toUtf8().data(), event)) { QQuickFramebufferObject::hoverEnterEvent(event); }; };
	void hoverLeaveEvent(QHoverEvent * event) { if (!callbackQQuickFramebufferObjectHoverLeaveEvent(this->objectName().toUtf8().data(), event)) { QQuickFramebufferObject::hoverLeaveEvent(event); }; };
	void hoverMoveEvent(QHoverEvent * event) { if (!callbackQQuickFramebufferObjectHoverMoveEvent(this->objectName().toUtf8().data(), event)) { QQuickFramebufferObject::hoverMoveEvent(event); }; };
	void inputMethodEvent(QInputMethodEvent * event) { if (!callbackQQuickFramebufferObjectInputMethodEvent(this->objectName().toUtf8().data(), event)) { QQuickFramebufferObject::inputMethodEvent(event); }; };
	void keyPressEvent(QKeyEvent * event) { if (!callbackQQuickFramebufferObjectKeyPressEvent(this->objectName().toUtf8().data(), event)) { QQuickFramebufferObject::keyPressEvent(event); }; };
	void keyReleaseEvent(QKeyEvent * event) { if (!callbackQQuickFramebufferObjectKeyReleaseEvent(this->objectName().toUtf8().data(), event)) { QQuickFramebufferObject::keyReleaseEvent(event); }; };
	void mouseDoubleClickEvent(QMouseEvent * event) { if (!callbackQQuickFramebufferObjectMouseDoubleClickEvent(this->objectName().toUtf8().data(), event)) { QQuickFramebufferObject::mouseDoubleClickEvent(event); }; };
	void mouseMoveEvent(QMouseEvent * event) { if (!callbackQQuickFramebufferObjectMouseMoveEvent(this->objectName().toUtf8().data(), event)) { QQuickFramebufferObject::mouseMoveEvent(event); }; };
	void mousePressEvent(QMouseEvent * event) { if (!callbackQQuickFramebufferObjectMousePressEvent(this->objectName().toUtf8().data(), event)) { QQuickFramebufferObject::mousePressEvent(event); }; };
	void mouseReleaseEvent(QMouseEvent * event) { if (!callbackQQuickFramebufferObjectMouseReleaseEvent(this->objectName().toUtf8().data(), event)) { QQuickFramebufferObject::mouseReleaseEvent(event); }; };
	void mouseUngrabEvent() { if (!callbackQQuickFramebufferObjectMouseUngrabEvent(this->objectName().toUtf8().data())) { QQuickFramebufferObject::mouseUngrabEvent(); }; };
	void touchEvent(QTouchEvent * event) { if (!callbackQQuickFramebufferObjectTouchEvent(this->objectName().toUtf8().data(), event)) { QQuickFramebufferObject::touchEvent(event); }; };
	void touchUngrabEvent() { if (!callbackQQuickFramebufferObjectTouchUngrabEvent(this->objectName().toUtf8().data())) { QQuickFramebufferObject::touchUngrabEvent(); }; };
	void updatePolish() { if (!callbackQQuickFramebufferObjectUpdatePolish(this->objectName().toUtf8().data())) { QQuickFramebufferObject::updatePolish(); }; };
	void wheelEvent(QWheelEvent * event) { if (!callbackQQuickFramebufferObjectWheelEvent(this->objectName().toUtf8().data(), event)) { QQuickFramebufferObject::wheelEvent(event); }; };
	void timerEvent(QTimerEvent * event) { if (!callbackQQuickFramebufferObjectTimerEvent(this->objectName().toUtf8().data(), event)) { QQuickFramebufferObject::timerEvent(event); }; };
	void childEvent(QChildEvent * event) { if (!callbackQQuickFramebufferObjectChildEvent(this->objectName().toUtf8().data(), event)) { QQuickFramebufferObject::childEvent(event); }; };
	void customEvent(QEvent * event) { if (!callbackQQuickFramebufferObjectCustomEvent(this->objectName().toUtf8().data(), event)) { QQuickFramebufferObject::customEvent(event); }; };
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
	static_cast<QQuickFramebufferObject*>(ptr)->releaseResources();
}

void QQuickFramebufferObject_ConnectTextureFollowsItemSizeChanged(void* ptr){
	QObject::connect(static_cast<QQuickFramebufferObject*>(ptr), static_cast<void (QQuickFramebufferObject::*)(bool)>(&QQuickFramebufferObject::textureFollowsItemSizeChanged), static_cast<MyQQuickFramebufferObject*>(ptr), static_cast<void (MyQQuickFramebufferObject::*)(bool)>(&MyQQuickFramebufferObject::Signal_TextureFollowsItemSizeChanged));;
}

void QQuickFramebufferObject_DisconnectTextureFollowsItemSizeChanged(void* ptr){
	QObject::disconnect(static_cast<QQuickFramebufferObject*>(ptr), static_cast<void (QQuickFramebufferObject::*)(bool)>(&QQuickFramebufferObject::textureFollowsItemSizeChanged), static_cast<MyQQuickFramebufferObject*>(ptr), static_cast<void (MyQQuickFramebufferObject::*)(bool)>(&MyQQuickFramebufferObject::Signal_TextureFollowsItemSizeChanged));;
}

void* QQuickFramebufferObject_TextureProvider(void* ptr){
	return static_cast<QQuickFramebufferObject*>(ptr)->textureProvider();
}

class MyQQuickImageProvider: public QQuickImageProvider {
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	MyQQuickImageProvider(ImageType type, Flags flags) : QQuickImageProvider(type, flags) {};
protected:
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
	void Signal_WindowChanged(QQuickWindow * window) { callbackQQuickItemWindowChanged(this->objectName().toUtf8().data(), window); };
protected:
	void classBegin() { if (!callbackQQuickItemClassBegin(this->objectName().toUtf8().data())) { QQuickItem::classBegin(); }; };
	void componentComplete() { if (!callbackQQuickItemComponentComplete(this->objectName().toUtf8().data())) { QQuickItem::componentComplete(); }; };
	void dragEnterEvent(QDragEnterEvent * event) { if (!callbackQQuickItemDragEnterEvent(this->objectName().toUtf8().data(), event)) { QQuickItem::dragEnterEvent(event); }; };
	void dragLeaveEvent(QDragLeaveEvent * event) { if (!callbackQQuickItemDragLeaveEvent(this->objectName().toUtf8().data(), event)) { QQuickItem::dragLeaveEvent(event); }; };
	void dragMoveEvent(QDragMoveEvent * event) { if (!callbackQQuickItemDragMoveEvent(this->objectName().toUtf8().data(), event)) { QQuickItem::dragMoveEvent(event); }; };
	void dropEvent(QDropEvent * event) { if (!callbackQQuickItemDropEvent(this->objectName().toUtf8().data(), event)) { QQuickItem::dropEvent(event); }; };
	void focusInEvent(QFocusEvent * event) { if (!callbackQQuickItemFocusInEvent(this->objectName().toUtf8().data(), event)) { QQuickItem::focusInEvent(event); }; };
	void focusOutEvent(QFocusEvent * event) { if (!callbackQQuickItemFocusOutEvent(this->objectName().toUtf8().data(), event)) { QQuickItem::focusOutEvent(event); }; };
	void hoverEnterEvent(QHoverEvent * event) { if (!callbackQQuickItemHoverEnterEvent(this->objectName().toUtf8().data(), event)) { QQuickItem::hoverEnterEvent(event); }; };
	void hoverLeaveEvent(QHoverEvent * event) { if (!callbackQQuickItemHoverLeaveEvent(this->objectName().toUtf8().data(), event)) { QQuickItem::hoverLeaveEvent(event); }; };
	void hoverMoveEvent(QHoverEvent * event) { if (!callbackQQuickItemHoverMoveEvent(this->objectName().toUtf8().data(), event)) { QQuickItem::hoverMoveEvent(event); }; };
	void inputMethodEvent(QInputMethodEvent * event) { if (!callbackQQuickItemInputMethodEvent(this->objectName().toUtf8().data(), event)) { QQuickItem::inputMethodEvent(event); }; };
	void keyPressEvent(QKeyEvent * event) { if (!callbackQQuickItemKeyPressEvent(this->objectName().toUtf8().data(), event)) { QQuickItem::keyPressEvent(event); }; };
	void keyReleaseEvent(QKeyEvent * event) { if (!callbackQQuickItemKeyReleaseEvent(this->objectName().toUtf8().data(), event)) { QQuickItem::keyReleaseEvent(event); }; };
	void mouseDoubleClickEvent(QMouseEvent * event) { if (!callbackQQuickItemMouseDoubleClickEvent(this->objectName().toUtf8().data(), event)) { QQuickItem::mouseDoubleClickEvent(event); }; };
	void mouseMoveEvent(QMouseEvent * event) { if (!callbackQQuickItemMouseMoveEvent(this->objectName().toUtf8().data(), event)) { QQuickItem::mouseMoveEvent(event); }; };
	void mousePressEvent(QMouseEvent * event) { if (!callbackQQuickItemMousePressEvent(this->objectName().toUtf8().data(), event)) { QQuickItem::mousePressEvent(event); }; };
	void mouseReleaseEvent(QMouseEvent * event) { if (!callbackQQuickItemMouseReleaseEvent(this->objectName().toUtf8().data(), event)) { QQuickItem::mouseReleaseEvent(event); }; };
	void mouseUngrabEvent() { if (!callbackQQuickItemMouseUngrabEvent(this->objectName().toUtf8().data())) { QQuickItem::mouseUngrabEvent(); }; };
	void releaseResources() { if (!callbackQQuickItemReleaseResources(this->objectName().toUtf8().data())) { QQuickItem::releaseResources(); }; };
	void touchEvent(QTouchEvent * event) { if (!callbackQQuickItemTouchEvent(this->objectName().toUtf8().data(), event)) { QQuickItem::touchEvent(event); }; };
	void touchUngrabEvent() { if (!callbackQQuickItemTouchUngrabEvent(this->objectName().toUtf8().data())) { QQuickItem::touchUngrabEvent(); }; };
	void updatePolish() { if (!callbackQQuickItemUpdatePolish(this->objectName().toUtf8().data())) { QQuickItem::updatePolish(); }; };
	void wheelEvent(QWheelEvent * event) { if (!callbackQQuickItemWheelEvent(this->objectName().toUtf8().data(), event)) { QQuickItem::wheelEvent(event); }; };
	void timerEvent(QTimerEvent * event) { if (!callbackQQuickItemTimerEvent(this->objectName().toUtf8().data(), event)) { QQuickItem::timerEvent(event); }; };
	void childEvent(QChildEvent * event) { if (!callbackQQuickItemChildEvent(this->objectName().toUtf8().data(), event)) { QQuickItem::childEvent(event); }; };
	void customEvent(QEvent * event) { if (!callbackQQuickItemCustomEvent(this->objectName().toUtf8().data(), event)) { QQuickItem::customEvent(event); }; };
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

int QQuickItem_Contains(void* ptr, void* point){
	return static_cast<QQuickItem*>(ptr)->contains(*static_cast<QPointF*>(point));
}

int QQuickItem_FiltersChildMouseEvents(void* ptr){
	return static_cast<QQuickItem*>(ptr)->filtersChildMouseEvents();
}

int QQuickItem_Flags(void* ptr){
	return static_cast<QQuickItem*>(ptr)->flags();
}

void QQuickItem_ForceActiveFocus(void* ptr){
	static_cast<QQuickItem*>(ptr)->forceActiveFocus();
}

void QQuickItem_ForceActiveFocus2(void* ptr, int reason){
	static_cast<QQuickItem*>(ptr)->forceActiveFocus(static_cast<Qt::FocusReason>(reason));
}

void QQuickItem_GrabMouse(void* ptr){
	static_cast<QQuickItem*>(ptr)->grabMouse();
}

double QQuickItem_ImplicitWidth(void* ptr){
	return static_cast<double>(static_cast<QQuickItem*>(ptr)->implicitWidth());
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

void* QQuickItem_NextItemInFocusChain(void* ptr, int forward){
	return static_cast<QQuickItem*>(ptr)->nextItemInFocusChain(forward != 0);
}

void QQuickItem_Polish(void* ptr){
	static_cast<QQuickItem*>(ptr)->polish();
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

void* QQuickItem_Window(void* ptr){
	return static_cast<QQuickItem*>(ptr)->window();
}

void QQuickItem_ConnectWindowChanged(void* ptr){
	QObject::connect(static_cast<QQuickItem*>(ptr), static_cast<void (QQuickItem::*)(QQuickWindow *)>(&QQuickItem::windowChanged), static_cast<MyQQuickItem*>(ptr), static_cast<void (MyQQuickItem::*)(QQuickWindow *)>(&MyQQuickItem::Signal_WindowChanged));;
}

void QQuickItem_DisconnectWindowChanged(void* ptr){
	QObject::disconnect(static_cast<QQuickItem*>(ptr), static_cast<void (QQuickItem::*)(QQuickWindow *)>(&QQuickItem::windowChanged), static_cast<MyQQuickItem*>(ptr), static_cast<void (MyQQuickItem::*)(QQuickWindow *)>(&MyQQuickItem::Signal_WindowChanged));;
}

void QQuickItem_DestroyQQuickItem(void* ptr){
	static_cast<QQuickItem*>(ptr)->~QQuickItem();
}

class MyQQuickItemGrabResult: public QQuickItemGrabResult {
public:
	void Signal_Ready() { callbackQQuickItemGrabResultReady(this->objectName().toUtf8().data()); };
protected:
	void timerEvent(QTimerEvent * event) { if (!callbackQQuickItemGrabResultTimerEvent(this->objectName().toUtf8().data(), event)) { QQuickItemGrabResult::timerEvent(event); }; };
	void childEvent(QChildEvent * event) { if (!callbackQQuickItemGrabResultChildEvent(this->objectName().toUtf8().data(), event)) { QQuickItemGrabResult::childEvent(event); }; };
	void customEvent(QEvent * event) { if (!callbackQQuickItemGrabResultCustomEvent(this->objectName().toUtf8().data(), event)) { QQuickItemGrabResult::customEvent(event); }; };
};

void* QQuickItemGrabResult_Url(void* ptr){
	return new QUrl(static_cast<QQuickItemGrabResult*>(ptr)->url());
}

void QQuickItemGrabResult_ConnectReady(void* ptr){
	QObject::connect(static_cast<QQuickItemGrabResult*>(ptr), static_cast<void (QQuickItemGrabResult::*)()>(&QQuickItemGrabResult::ready), static_cast<MyQQuickItemGrabResult*>(ptr), static_cast<void (MyQQuickItemGrabResult::*)()>(&MyQQuickItemGrabResult::Signal_Ready));;
}

void QQuickItemGrabResult_DisconnectReady(void* ptr){
	QObject::disconnect(static_cast<QQuickItemGrabResult*>(ptr), static_cast<void (QQuickItemGrabResult::*)()>(&QQuickItemGrabResult::ready), static_cast<MyQQuickItemGrabResult*>(ptr), static_cast<void (MyQQuickItemGrabResult::*)()>(&MyQQuickItemGrabResult::Signal_Ready));;
}

int QQuickItemGrabResult_SaveToFile(void* ptr, char* fileName){
	return static_cast<QQuickItemGrabResult*>(ptr)->saveToFile(QString(fileName));
}

class MyQQuickPaintedItem: public QQuickPaintedItem {
public:
	void Signal_ContentsScaleChanged() { callbackQQuickPaintedItemContentsScaleChanged(this->objectName().toUtf8().data()); };
	void Signal_ContentsSizeChanged() { callbackQQuickPaintedItemContentsSizeChanged(this->objectName().toUtf8().data()); };
	void Signal_FillColorChanged() { callbackQQuickPaintedItemFillColorChanged(this->objectName().toUtf8().data()); };
	void Signal_RenderTargetChanged() { callbackQQuickPaintedItemRenderTargetChanged(this->objectName().toUtf8().data()); };
protected:
	void releaseResources() { if (!callbackQQuickPaintedItemReleaseResources(this->objectName().toUtf8().data())) { QQuickPaintedItem::releaseResources(); }; };
	void classBegin() { if (!callbackQQuickPaintedItemClassBegin(this->objectName().toUtf8().data())) { QQuickPaintedItem::classBegin(); }; };
	void componentComplete() { if (!callbackQQuickPaintedItemComponentComplete(this->objectName().toUtf8().data())) { QQuickPaintedItem::componentComplete(); }; };
	void dragEnterEvent(QDragEnterEvent * event) { if (!callbackQQuickPaintedItemDragEnterEvent(this->objectName().toUtf8().data(), event)) { QQuickPaintedItem::dragEnterEvent(event); }; };
	void dragLeaveEvent(QDragLeaveEvent * event) { if (!callbackQQuickPaintedItemDragLeaveEvent(this->objectName().toUtf8().data(), event)) { QQuickPaintedItem::dragLeaveEvent(event); }; };
	void dragMoveEvent(QDragMoveEvent * event) { if (!callbackQQuickPaintedItemDragMoveEvent(this->objectName().toUtf8().data(), event)) { QQuickPaintedItem::dragMoveEvent(event); }; };
	void dropEvent(QDropEvent * event) { if (!callbackQQuickPaintedItemDropEvent(this->objectName().toUtf8().data(), event)) { QQuickPaintedItem::dropEvent(event); }; };
	void focusInEvent(QFocusEvent * event) { if (!callbackQQuickPaintedItemFocusInEvent(this->objectName().toUtf8().data(), event)) { QQuickPaintedItem::focusInEvent(event); }; };
	void focusOutEvent(QFocusEvent * event) { if (!callbackQQuickPaintedItemFocusOutEvent(this->objectName().toUtf8().data(), event)) { QQuickPaintedItem::focusOutEvent(event); }; };
	void hoverEnterEvent(QHoverEvent * event) { if (!callbackQQuickPaintedItemHoverEnterEvent(this->objectName().toUtf8().data(), event)) { QQuickPaintedItem::hoverEnterEvent(event); }; };
	void hoverLeaveEvent(QHoverEvent * event) { if (!callbackQQuickPaintedItemHoverLeaveEvent(this->objectName().toUtf8().data(), event)) { QQuickPaintedItem::hoverLeaveEvent(event); }; };
	void hoverMoveEvent(QHoverEvent * event) { if (!callbackQQuickPaintedItemHoverMoveEvent(this->objectName().toUtf8().data(), event)) { QQuickPaintedItem::hoverMoveEvent(event); }; };
	void inputMethodEvent(QInputMethodEvent * event) { if (!callbackQQuickPaintedItemInputMethodEvent(this->objectName().toUtf8().data(), event)) { QQuickPaintedItem::inputMethodEvent(event); }; };
	void keyPressEvent(QKeyEvent * event) { if (!callbackQQuickPaintedItemKeyPressEvent(this->objectName().toUtf8().data(), event)) { QQuickPaintedItem::keyPressEvent(event); }; };
	void keyReleaseEvent(QKeyEvent * event) { if (!callbackQQuickPaintedItemKeyReleaseEvent(this->objectName().toUtf8().data(), event)) { QQuickPaintedItem::keyReleaseEvent(event); }; };
	void mouseDoubleClickEvent(QMouseEvent * event) { if (!callbackQQuickPaintedItemMouseDoubleClickEvent(this->objectName().toUtf8().data(), event)) { QQuickPaintedItem::mouseDoubleClickEvent(event); }; };
	void mouseMoveEvent(QMouseEvent * event) { if (!callbackQQuickPaintedItemMouseMoveEvent(this->objectName().toUtf8().data(), event)) { QQuickPaintedItem::mouseMoveEvent(event); }; };
	void mousePressEvent(QMouseEvent * event) { if (!callbackQQuickPaintedItemMousePressEvent(this->objectName().toUtf8().data(), event)) { QQuickPaintedItem::mousePressEvent(event); }; };
	void mouseReleaseEvent(QMouseEvent * event) { if (!callbackQQuickPaintedItemMouseReleaseEvent(this->objectName().toUtf8().data(), event)) { QQuickPaintedItem::mouseReleaseEvent(event); }; };
	void mouseUngrabEvent() { if (!callbackQQuickPaintedItemMouseUngrabEvent(this->objectName().toUtf8().data())) { QQuickPaintedItem::mouseUngrabEvent(); }; };
	void touchEvent(QTouchEvent * event) { if (!callbackQQuickPaintedItemTouchEvent(this->objectName().toUtf8().data(), event)) { QQuickPaintedItem::touchEvent(event); }; };
	void touchUngrabEvent() { if (!callbackQQuickPaintedItemTouchUngrabEvent(this->objectName().toUtf8().data())) { QQuickPaintedItem::touchUngrabEvent(); }; };
	void updatePolish() { if (!callbackQQuickPaintedItemUpdatePolish(this->objectName().toUtf8().data())) { QQuickPaintedItem::updatePolish(); }; };
	void wheelEvent(QWheelEvent * event) { if (!callbackQQuickPaintedItemWheelEvent(this->objectName().toUtf8().data(), event)) { QQuickPaintedItem::wheelEvent(event); }; };
	void timerEvent(QTimerEvent * event) { if (!callbackQQuickPaintedItemTimerEvent(this->objectName().toUtf8().data(), event)) { QQuickPaintedItem::timerEvent(event); }; };
	void childEvent(QChildEvent * event) { if (!callbackQQuickPaintedItemChildEvent(this->objectName().toUtf8().data(), event)) { QQuickPaintedItem::childEvent(event); }; };
	void customEvent(QEvent * event) { if (!callbackQQuickPaintedItemCustomEvent(this->objectName().toUtf8().data(), event)) { QQuickPaintedItem::customEvent(event); }; };
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

void QQuickPaintedItem_ConnectContentsScaleChanged(void* ptr){
	QObject::connect(static_cast<QQuickPaintedItem*>(ptr), static_cast<void (QQuickPaintedItem::*)()>(&QQuickPaintedItem::contentsScaleChanged), static_cast<MyQQuickPaintedItem*>(ptr), static_cast<void (MyQQuickPaintedItem::*)()>(&MyQQuickPaintedItem::Signal_ContentsScaleChanged));;
}

void QQuickPaintedItem_DisconnectContentsScaleChanged(void* ptr){
	QObject::disconnect(static_cast<QQuickPaintedItem*>(ptr), static_cast<void (QQuickPaintedItem::*)()>(&QQuickPaintedItem::contentsScaleChanged), static_cast<MyQQuickPaintedItem*>(ptr), static_cast<void (MyQQuickPaintedItem::*)()>(&MyQQuickPaintedItem::Signal_ContentsScaleChanged));;
}

void QQuickPaintedItem_ConnectContentsSizeChanged(void* ptr){
	QObject::connect(static_cast<QQuickPaintedItem*>(ptr), static_cast<void (QQuickPaintedItem::*)()>(&QQuickPaintedItem::contentsSizeChanged), static_cast<MyQQuickPaintedItem*>(ptr), static_cast<void (MyQQuickPaintedItem::*)()>(&MyQQuickPaintedItem::Signal_ContentsSizeChanged));;
}

void QQuickPaintedItem_DisconnectContentsSizeChanged(void* ptr){
	QObject::disconnect(static_cast<QQuickPaintedItem*>(ptr), static_cast<void (QQuickPaintedItem::*)()>(&QQuickPaintedItem::contentsSizeChanged), static_cast<MyQQuickPaintedItem*>(ptr), static_cast<void (MyQQuickPaintedItem::*)()>(&MyQQuickPaintedItem::Signal_ContentsSizeChanged));;
}

void QQuickPaintedItem_ConnectFillColorChanged(void* ptr){
	QObject::connect(static_cast<QQuickPaintedItem*>(ptr), static_cast<void (QQuickPaintedItem::*)()>(&QQuickPaintedItem::fillColorChanged), static_cast<MyQQuickPaintedItem*>(ptr), static_cast<void (MyQQuickPaintedItem::*)()>(&MyQQuickPaintedItem::Signal_FillColorChanged));;
}

void QQuickPaintedItem_DisconnectFillColorChanged(void* ptr){
	QObject::disconnect(static_cast<QQuickPaintedItem*>(ptr), static_cast<void (QQuickPaintedItem::*)()>(&QQuickPaintedItem::fillColorChanged), static_cast<MyQQuickPaintedItem*>(ptr), static_cast<void (MyQQuickPaintedItem::*)()>(&MyQQuickPaintedItem::Signal_FillColorChanged));;
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

void QQuickPaintedItem_ConnectRenderTargetChanged(void* ptr){
	QObject::connect(static_cast<QQuickPaintedItem*>(ptr), static_cast<void (QQuickPaintedItem::*)()>(&QQuickPaintedItem::renderTargetChanged), static_cast<MyQQuickPaintedItem*>(ptr), static_cast<void (MyQQuickPaintedItem::*)()>(&MyQQuickPaintedItem::Signal_RenderTargetChanged));;
}

void QQuickPaintedItem_DisconnectRenderTargetChanged(void* ptr){
	QObject::disconnect(static_cast<QQuickPaintedItem*>(ptr), static_cast<void (QQuickPaintedItem::*)()>(&QQuickPaintedItem::renderTargetChanged), static_cast<MyQQuickPaintedItem*>(ptr), static_cast<void (MyQQuickPaintedItem::*)()>(&MyQQuickPaintedItem::Signal_RenderTargetChanged));;
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

class MyQQuickRenderControl: public QQuickRenderControl {
public:
	MyQQuickRenderControl(QObject *parent) : QQuickRenderControl(parent) {};
	void Signal_RenderRequested() { callbackQQuickRenderControlRenderRequested(this->objectName().toUtf8().data()); };
	void Signal_SceneChanged() { callbackQQuickRenderControlSceneChanged(this->objectName().toUtf8().data()); };
protected:
	void timerEvent(QTimerEvent * event) { if (!callbackQQuickRenderControlTimerEvent(this->objectName().toUtf8().data(), event)) { QQuickRenderControl::timerEvent(event); }; };
	void childEvent(QChildEvent * event) { if (!callbackQQuickRenderControlChildEvent(this->objectName().toUtf8().data(), event)) { QQuickRenderControl::childEvent(event); }; };
	void customEvent(QEvent * event) { if (!callbackQQuickRenderControlCustomEvent(this->objectName().toUtf8().data(), event)) { QQuickRenderControl::customEvent(event); }; };
};

void* QQuickRenderControl_NewQQuickRenderControl(void* parent){
	return new MyQQuickRenderControl(static_cast<QObject*>(parent));
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

int QQuickRenderControl_Sync(void* ptr){
	return static_cast<QQuickRenderControl*>(ptr)->sync();
}

void QQuickRenderControl_DestroyQQuickRenderControl(void* ptr){
	static_cast<QQuickRenderControl*>(ptr)->~QQuickRenderControl();
}

void* QQuickTextDocument_NewQQuickTextDocument(void* parent){
	return new QQuickTextDocument(static_cast<QQuickItem*>(parent));
}

void* QQuickTextDocument_TextDocument(void* ptr){
	return static_cast<QQuickTextDocument*>(ptr)->textDocument();
}

class MyQQuickTextureFactory: public QQuickTextureFactory {
public:
protected:
	void timerEvent(QTimerEvent * event) { if (!callbackQQuickTextureFactoryTimerEvent(this->objectName().toUtf8().data(), event)) { QQuickTextureFactory::timerEvent(event); }; };
	void childEvent(QChildEvent * event) { if (!callbackQQuickTextureFactoryChildEvent(this->objectName().toUtf8().data(), event)) { QQuickTextureFactory::childEvent(event); }; };
	void customEvent(QEvent * event) { if (!callbackQQuickTextureFactoryCustomEvent(this->objectName().toUtf8().data(), event)) { QQuickTextureFactory::customEvent(event); }; };
};

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

class MyQQuickView: public QQuickView {
public:
	MyQQuickView(QQmlEngine *engine, QWindow *parent) : QQuickView(engine, parent) {};
	MyQQuickView(QWindow *parent) : QQuickView(parent) {};
	MyQQuickView(const QUrl &source, QWindow *parent) : QQuickView(source, parent) {};
	void Signal_StatusChanged(QQuickView::Status status) { callbackQQuickViewStatusChanged(this->objectName().toUtf8().data(), status); };
protected:
	void keyPressEvent(QKeyEvent * e) { if (!callbackQQuickViewKeyPressEvent(this->objectName().toUtf8().data(), e)) { QQuickView::keyPressEvent(e); }; };
	void keyReleaseEvent(QKeyEvent * e) { if (!callbackQQuickViewKeyReleaseEvent(this->objectName().toUtf8().data(), e)) { QQuickView::keyReleaseEvent(e); }; };
	void mouseMoveEvent(QMouseEvent * e) { if (!callbackQQuickViewMouseMoveEvent(this->objectName().toUtf8().data(), e)) { QQuickView::mouseMoveEvent(e); }; };
	void mousePressEvent(QMouseEvent * e) { if (!callbackQQuickViewMousePressEvent(this->objectName().toUtf8().data(), e)) { QQuickView::mousePressEvent(e); }; };
	void mouseReleaseEvent(QMouseEvent * e) { if (!callbackQQuickViewMouseReleaseEvent(this->objectName().toUtf8().data(), e)) { QQuickView::mouseReleaseEvent(e); }; };
	void exposeEvent(QExposeEvent * v) { if (!callbackQQuickViewExposeEvent(this->objectName().toUtf8().data(), v)) { QQuickView::exposeEvent(v); }; };
	void focusInEvent(QFocusEvent * ev) { if (!callbackQQuickViewFocusInEvent(this->objectName().toUtf8().data(), ev)) { QQuickView::focusInEvent(ev); }; };
	void focusOutEvent(QFocusEvent * ev) { if (!callbackQQuickViewFocusOutEvent(this->objectName().toUtf8().data(), ev)) { QQuickView::focusOutEvent(ev); }; };
	void hideEvent(QHideEvent * v) { if (!callbackQQuickViewHideEvent(this->objectName().toUtf8().data(), v)) { QQuickView::hideEvent(v); }; };
	void mouseDoubleClickEvent(QMouseEvent * event) { if (!callbackQQuickViewMouseDoubleClickEvent(this->objectName().toUtf8().data(), event)) { QQuickView::mouseDoubleClickEvent(event); }; };
	void resizeEvent(QResizeEvent * ev) { if (!callbackQQuickViewResizeEvent(this->objectName().toUtf8().data(), ev)) { QQuickView::resizeEvent(ev); }; };
	void showEvent(QShowEvent * v) { if (!callbackQQuickViewShowEvent(this->objectName().toUtf8().data(), v)) { QQuickView::showEvent(v); }; };
	void wheelEvent(QWheelEvent * event) { if (!callbackQQuickViewWheelEvent(this->objectName().toUtf8().data(), event)) { QQuickView::wheelEvent(event); }; };
	void moveEvent(QMoveEvent * ev) { if (!callbackQQuickViewMoveEvent(this->objectName().toUtf8().data(), ev)) { QQuickView::moveEvent(ev); }; };
	void tabletEvent(QTabletEvent * ev) { if (!callbackQQuickViewTabletEvent(this->objectName().toUtf8().data(), ev)) { QQuickView::tabletEvent(ev); }; };
	void touchEvent(QTouchEvent * ev) { if (!callbackQQuickViewTouchEvent(this->objectName().toUtf8().data(), ev)) { QQuickView::touchEvent(ev); }; };
	void timerEvent(QTimerEvent * event) { if (!callbackQQuickViewTimerEvent(this->objectName().toUtf8().data(), event)) { QQuickView::timerEvent(event); }; };
	void childEvent(QChildEvent * event) { if (!callbackQQuickViewChildEvent(this->objectName().toUtf8().data(), event)) { QQuickView::childEvent(event); }; };
	void customEvent(QEvent * event) { if (!callbackQQuickViewCustomEvent(this->objectName().toUtf8().data(), event)) { QQuickView::customEvent(event); }; };
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

void QQuickView_DestroyQQuickView(void* ptr){
	static_cast<QQuickView*>(ptr)->~QQuickView();
}

class MyQQuickWidget: public QQuickWidget {
public:
	MyQQuickWidget(QQmlEngine *engine, QWidget *parent) : QQuickWidget(engine, parent) {};
	MyQQuickWidget(QWidget *parent) : QQuickWidget(parent) {};
	MyQQuickWidget(const QUrl &source, QWidget *parent) : QQuickWidget(source, parent) {};
	void Signal_SceneGraphError(QQuickWindow::SceneGraphError error, const QString & message) { callbackQQuickWidgetSceneGraphError(this->objectName().toUtf8().data(), error, message.toUtf8().data()); };
	void Signal_StatusChanged(QQuickWidget::Status status) { callbackQQuickWidgetStatusChanged(this->objectName().toUtf8().data(), status); };
	void setVisible(bool visible) { if (!callbackQQuickWidgetSetVisible(this->objectName().toUtf8().data(), visible)) { QQuickWidget::setVisible(visible); }; };
protected:
	void focusInEvent(QFocusEvent * event) { if (!callbackQQuickWidgetFocusInEvent(this->objectName().toUtf8().data(), event)) { QQuickWidget::focusInEvent(event); }; };
	void focusOutEvent(QFocusEvent * event) { if (!callbackQQuickWidgetFocusOutEvent(this->objectName().toUtf8().data(), event)) { QQuickWidget::focusOutEvent(event); }; };
	void dragEnterEvent(QDragEnterEvent * e) { if (!callbackQQuickWidgetDragEnterEvent(this->objectName().toUtf8().data(), e)) { QQuickWidget::dragEnterEvent(e); }; };
	void dragLeaveEvent(QDragLeaveEvent * e) { if (!callbackQQuickWidgetDragLeaveEvent(this->objectName().toUtf8().data(), e)) { QQuickWidget::dragLeaveEvent(e); }; };
	void dragMoveEvent(QDragMoveEvent * e) { if (!callbackQQuickWidgetDragMoveEvent(this->objectName().toUtf8().data(), e)) { QQuickWidget::dragMoveEvent(e); }; };
	void dropEvent(QDropEvent * e) { if (!callbackQQuickWidgetDropEvent(this->objectName().toUtf8().data(), e)) { QQuickWidget::dropEvent(e); }; };
	void hideEvent(QHideEvent * v) { if (!callbackQQuickWidgetHideEvent(this->objectName().toUtf8().data(), v)) { QQuickWidget::hideEvent(v); }; };
	void keyPressEvent(QKeyEvent * e) { if (!callbackQQuickWidgetKeyPressEvent(this->objectName().toUtf8().data(), e)) { QQuickWidget::keyPressEvent(e); }; };
	void keyReleaseEvent(QKeyEvent * e) { if (!callbackQQuickWidgetKeyReleaseEvent(this->objectName().toUtf8().data(), e)) { QQuickWidget::keyReleaseEvent(e); }; };
	void mouseDoubleClickEvent(QMouseEvent * e) { if (!callbackQQuickWidgetMouseDoubleClickEvent(this->objectName().toUtf8().data(), e)) { QQuickWidget::mouseDoubleClickEvent(e); }; };
	void mouseMoveEvent(QMouseEvent * e) { if (!callbackQQuickWidgetMouseMoveEvent(this->objectName().toUtf8().data(), e)) { QQuickWidget::mouseMoveEvent(e); }; };
	void mousePressEvent(QMouseEvent * e) { if (!callbackQQuickWidgetMousePressEvent(this->objectName().toUtf8().data(), e)) { QQuickWidget::mousePressEvent(e); }; };
	void mouseReleaseEvent(QMouseEvent * e) { if (!callbackQQuickWidgetMouseReleaseEvent(this->objectName().toUtf8().data(), e)) { QQuickWidget::mouseReleaseEvent(e); }; };
	void showEvent(QShowEvent * v) { if (!callbackQQuickWidgetShowEvent(this->objectName().toUtf8().data(), v)) { QQuickWidget::showEvent(v); }; };
	void wheelEvent(QWheelEvent * e) { if (!callbackQQuickWidgetWheelEvent(this->objectName().toUtf8().data(), e)) { QQuickWidget::wheelEvent(e); }; };
	void actionEvent(QActionEvent * event) { if (!callbackQQuickWidgetActionEvent(this->objectName().toUtf8().data(), event)) { QQuickWidget::actionEvent(event); }; };
	void enterEvent(QEvent * event) { if (!callbackQQuickWidgetEnterEvent(this->objectName().toUtf8().data(), event)) { QQuickWidget::enterEvent(event); }; };
	void leaveEvent(QEvent * event) { if (!callbackQQuickWidgetLeaveEvent(this->objectName().toUtf8().data(), event)) { QQuickWidget::leaveEvent(event); }; };
	void moveEvent(QMoveEvent * event) { if (!callbackQQuickWidgetMoveEvent(this->objectName().toUtf8().data(), event)) { QQuickWidget::moveEvent(event); }; };
	void paintEvent(QPaintEvent * event) { if (!callbackQQuickWidgetPaintEvent(this->objectName().toUtf8().data(), event)) { QQuickWidget::paintEvent(event); }; };
	void changeEvent(QEvent * event) { if (!callbackQQuickWidgetChangeEvent(this->objectName().toUtf8().data(), event)) { QQuickWidget::changeEvent(event); }; };
	void closeEvent(QCloseEvent * event) { if (!callbackQQuickWidgetCloseEvent(this->objectName().toUtf8().data(), event)) { QQuickWidget::closeEvent(event); }; };
	void contextMenuEvent(QContextMenuEvent * event) { if (!callbackQQuickWidgetContextMenuEvent(this->objectName().toUtf8().data(), event)) { QQuickWidget::contextMenuEvent(event); }; };
	void initPainter(QPainter * painter) const { if (!callbackQQuickWidgetInitPainter(this->objectName().toUtf8().data(), painter)) { QQuickWidget::initPainter(painter); }; };
	void inputMethodEvent(QInputMethodEvent * event) { if (!callbackQQuickWidgetInputMethodEvent(this->objectName().toUtf8().data(), event)) { QQuickWidget::inputMethodEvent(event); }; };
	void resizeEvent(QResizeEvent * event) { if (!callbackQQuickWidgetResizeEvent(this->objectName().toUtf8().data(), event)) { QQuickWidget::resizeEvent(event); }; };
	void tabletEvent(QTabletEvent * event) { if (!callbackQQuickWidgetTabletEvent(this->objectName().toUtf8().data(), event)) { QQuickWidget::tabletEvent(event); }; };
	void timerEvent(QTimerEvent * event) { if (!callbackQQuickWidgetTimerEvent(this->objectName().toUtf8().data(), event)) { QQuickWidget::timerEvent(event); }; };
	void childEvent(QChildEvent * event) { if (!callbackQQuickWidgetChildEvent(this->objectName().toUtf8().data(), event)) { QQuickWidget::childEvent(event); }; };
	void customEvent(QEvent * event) { if (!callbackQQuickWidgetCustomEvent(this->objectName().toUtf8().data(), event)) { QQuickWidget::customEvent(event); }; };
};

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

void* QQuickWidget_Engine(void* ptr){
	return static_cast<QQuickWidget*>(ptr)->engine();
}

void* QQuickWidget_InitialSize(void* ptr){
	return new QSize(static_cast<QSize>(static_cast<QQuickWidget*>(ptr)->initialSize()).width(), static_cast<QSize>(static_cast<QQuickWidget*>(ptr)->initialSize()).height());
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

void QQuickWidget_SetClearColor(void* ptr, void* color){
	static_cast<QQuickWidget*>(ptr)->setClearColor(*static_cast<QColor*>(color));
}

void QQuickWidget_SetFormat(void* ptr, void* format){
	static_cast<QQuickWidget*>(ptr)->setFormat(*static_cast<QSurfaceFormat*>(format));
}

void QQuickWidget_SetSource(void* ptr, void* url){
	QMetaObject::invokeMethod(static_cast<QQuickWidget*>(ptr), "setSource", Q_ARG(QUrl, *static_cast<QUrl*>(url)));
}

void QQuickWidget_ConnectStatusChanged(void* ptr){
	QObject::connect(static_cast<QQuickWidget*>(ptr), static_cast<void (QQuickWidget::*)(QQuickWidget::Status)>(&QQuickWidget::statusChanged), static_cast<MyQQuickWidget*>(ptr), static_cast<void (MyQQuickWidget::*)(QQuickWidget::Status)>(&MyQQuickWidget::Signal_StatusChanged));;
}

void QQuickWidget_DisconnectStatusChanged(void* ptr){
	QObject::disconnect(static_cast<QQuickWidget*>(ptr), static_cast<void (QQuickWidget::*)(QQuickWidget::Status)>(&QQuickWidget::statusChanged), static_cast<MyQQuickWidget*>(ptr), static_cast<void (MyQQuickWidget::*)(QQuickWidget::Status)>(&MyQQuickWidget::Signal_StatusChanged));;
}

void* QQuickWidget_Source(void* ptr){
	return new QUrl(static_cast<QQuickWidget*>(ptr)->source());
}

void QQuickWidget_DestroyQQuickWidget(void* ptr){
	static_cast<QQuickWidget*>(ptr)->~QQuickWidget();
}

class MyQQuickWindow: public QQuickWindow {
public:
	MyQQuickWindow(QWindow *parent) : QQuickWindow(parent) {};
	void Signal_ActiveFocusItemChanged() { callbackQQuickWindowActiveFocusItemChanged(this->objectName().toUtf8().data()); };
	void Signal_AfterAnimating() { callbackQQuickWindowAfterAnimating(this->objectName().toUtf8().data()); };
	void Signal_AfterRendering() { callbackQQuickWindowAfterRendering(this->objectName().toUtf8().data()); };
	void Signal_AfterSynchronizing() { callbackQQuickWindowAfterSynchronizing(this->objectName().toUtf8().data()); };
	void Signal_BeforeRendering() { callbackQQuickWindowBeforeRendering(this->objectName().toUtf8().data()); };
	void Signal_BeforeSynchronizing() { callbackQQuickWindowBeforeSynchronizing(this->objectName().toUtf8().data()); };
	void Signal_ColorChanged(const QColor & v) { callbackQQuickWindowColorChanged(this->objectName().toUtf8().data(), new QColor(v)); };
	void Signal_FrameSwapped() { callbackQQuickWindowFrameSwapped(this->objectName().toUtf8().data()); };
	void Signal_OpenglContextCreated(QOpenGLContext * context) { callbackQQuickWindowOpenglContextCreated(this->objectName().toUtf8().data(), context); };
	void Signal_SceneGraphAboutToStop() { callbackQQuickWindowSceneGraphAboutToStop(this->objectName().toUtf8().data()); };
	void Signal_SceneGraphError(QQuickWindow::SceneGraphError error, const QString & message) { callbackQQuickWindowSceneGraphError(this->objectName().toUtf8().data(), error, message.toUtf8().data()); };
	void Signal_SceneGraphInitialized() { callbackQQuickWindowSceneGraphInitialized(this->objectName().toUtf8().data()); };
	void Signal_SceneGraphInvalidated() { callbackQQuickWindowSceneGraphInvalidated(this->objectName().toUtf8().data()); };
protected:
	void exposeEvent(QExposeEvent * v) { if (!callbackQQuickWindowExposeEvent(this->objectName().toUtf8().data(), v)) { QQuickWindow::exposeEvent(v); }; };
	void focusInEvent(QFocusEvent * ev) { if (!callbackQQuickWindowFocusInEvent(this->objectName().toUtf8().data(), ev)) { QQuickWindow::focusInEvent(ev); }; };
	void focusOutEvent(QFocusEvent * ev) { if (!callbackQQuickWindowFocusOutEvent(this->objectName().toUtf8().data(), ev)) { QQuickWindow::focusOutEvent(ev); }; };
	void hideEvent(QHideEvent * v) { if (!callbackQQuickWindowHideEvent(this->objectName().toUtf8().data(), v)) { QQuickWindow::hideEvent(v); }; };
	void keyPressEvent(QKeyEvent * e) { if (!callbackQQuickWindowKeyPressEvent(this->objectName().toUtf8().data(), e)) { QQuickWindow::keyPressEvent(e); }; };
	void keyReleaseEvent(QKeyEvent * e) { if (!callbackQQuickWindowKeyReleaseEvent(this->objectName().toUtf8().data(), e)) { QQuickWindow::keyReleaseEvent(e); }; };
	void mouseDoubleClickEvent(QMouseEvent * event) { if (!callbackQQuickWindowMouseDoubleClickEvent(this->objectName().toUtf8().data(), event)) { QQuickWindow::mouseDoubleClickEvent(event); }; };
	void mouseMoveEvent(QMouseEvent * event) { if (!callbackQQuickWindowMouseMoveEvent(this->objectName().toUtf8().data(), event)) { QQuickWindow::mouseMoveEvent(event); }; };
	void mousePressEvent(QMouseEvent * event) { if (!callbackQQuickWindowMousePressEvent(this->objectName().toUtf8().data(), event)) { QQuickWindow::mousePressEvent(event); }; };
	void mouseReleaseEvent(QMouseEvent * event) { if (!callbackQQuickWindowMouseReleaseEvent(this->objectName().toUtf8().data(), event)) { QQuickWindow::mouseReleaseEvent(event); }; };
	void resizeEvent(QResizeEvent * ev) { if (!callbackQQuickWindowResizeEvent(this->objectName().toUtf8().data(), ev)) { QQuickWindow::resizeEvent(ev); }; };
	void showEvent(QShowEvent * v) { if (!callbackQQuickWindowShowEvent(this->objectName().toUtf8().data(), v)) { QQuickWindow::showEvent(v); }; };
	void wheelEvent(QWheelEvent * event) { if (!callbackQQuickWindowWheelEvent(this->objectName().toUtf8().data(), event)) { QQuickWindow::wheelEvent(event); }; };
	void moveEvent(QMoveEvent * ev) { if (!callbackQQuickWindowMoveEvent(this->objectName().toUtf8().data(), ev)) { QQuickWindow::moveEvent(ev); }; };
	void tabletEvent(QTabletEvent * ev) { if (!callbackQQuickWindowTabletEvent(this->objectName().toUtf8().data(), ev)) { QQuickWindow::tabletEvent(ev); }; };
	void touchEvent(QTouchEvent * ev) { if (!callbackQQuickWindowTouchEvent(this->objectName().toUtf8().data(), ev)) { QQuickWindow::touchEvent(ev); }; };
	void timerEvent(QTimerEvent * event) { if (!callbackQQuickWindowTimerEvent(this->objectName().toUtf8().data(), event)) { QQuickWindow::timerEvent(event); }; };
	void childEvent(QChildEvent * event) { if (!callbackQQuickWindowChildEvent(this->objectName().toUtf8().data(), event)) { QQuickWindow::childEvent(event); }; };
	void customEvent(QEvent * event) { if (!callbackQQuickWindowCustomEvent(this->objectName().toUtf8().data(), event)) { QQuickWindow::customEvent(event); }; };
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

void QQuickWindow_ConnectAfterAnimating(void* ptr){
	QObject::connect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::afterAnimating), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_AfterAnimating));;
}

void QQuickWindow_DisconnectAfterAnimating(void* ptr){
	QObject::disconnect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::afterAnimating), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_AfterAnimating));;
}

void QQuickWindow_ConnectAfterRendering(void* ptr){
	QObject::connect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::afterRendering), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_AfterRendering));;
}

void QQuickWindow_DisconnectAfterRendering(void* ptr){
	QObject::disconnect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::afterRendering), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_AfterRendering));;
}

void QQuickWindow_ConnectAfterSynchronizing(void* ptr){
	QObject::connect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::afterSynchronizing), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_AfterSynchronizing));;
}

void QQuickWindow_DisconnectAfterSynchronizing(void* ptr){
	QObject::disconnect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::afterSynchronizing), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_AfterSynchronizing));;
}

void QQuickWindow_ConnectBeforeRendering(void* ptr){
	QObject::connect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::beforeRendering), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_BeforeRendering));;
}

void QQuickWindow_DisconnectBeforeRendering(void* ptr){
	QObject::disconnect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::beforeRendering), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_BeforeRendering));;
}

void QQuickWindow_ConnectBeforeSynchronizing(void* ptr){
	QObject::connect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::beforeSynchronizing), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_BeforeSynchronizing));;
}

void QQuickWindow_DisconnectBeforeSynchronizing(void* ptr){
	QObject::disconnect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::beforeSynchronizing), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_BeforeSynchronizing));;
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

void* QQuickWindow_CreateTextureFromImage2(void* ptr, void* image){
	return static_cast<QQuickWindow*>(ptr)->createTextureFromImage(*static_cast<QImage*>(image));
}

void* QQuickWindow_CreateTextureFromImage(void* ptr, void* image, int options){
	return static_cast<QQuickWindow*>(ptr)->createTextureFromImage(*static_cast<QImage*>(image), static_cast<QQuickWindow::CreateTextureOption>(options));
}

double QQuickWindow_EffectiveDevicePixelRatio(void* ptr){
	return static_cast<double>(static_cast<QQuickWindow*>(ptr)->effectiveDevicePixelRatio());
}

void QQuickWindow_ConnectFrameSwapped(void* ptr){
	QObject::connect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::frameSwapped), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_FrameSwapped));;
}

void QQuickWindow_DisconnectFrameSwapped(void* ptr){
	QObject::disconnect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::frameSwapped), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_FrameSwapped));;
}

int QQuickWindow_QQuickWindow_HasDefaultAlphaBuffer(){
	return QQuickWindow::hasDefaultAlphaBuffer();
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

void* QQuickWindow_MouseGrabberItem(void* ptr){
	return static_cast<QQuickWindow*>(ptr)->mouseGrabberItem();
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

void QQuickWindow_ConnectSceneGraphAboutToStop(void* ptr){
	QObject::connect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::sceneGraphAboutToStop), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_SceneGraphAboutToStop));;
}

void QQuickWindow_DisconnectSceneGraphAboutToStop(void* ptr){
	QObject::disconnect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::sceneGraphAboutToStop), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_SceneGraphAboutToStop));;
}

void QQuickWindow_ConnectSceneGraphError(void* ptr){
	QObject::connect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)(QQuickWindow::SceneGraphError, const QString &)>(&QQuickWindow::sceneGraphError), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)(QQuickWindow::SceneGraphError, const QString &)>(&MyQQuickWindow::Signal_SceneGraphError));;
}

void QQuickWindow_DisconnectSceneGraphError(void* ptr){
	QObject::disconnect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)(QQuickWindow::SceneGraphError, const QString &)>(&QQuickWindow::sceneGraphError), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)(QQuickWindow::SceneGraphError, const QString &)>(&MyQQuickWindow::Signal_SceneGraphError));;
}

void QQuickWindow_ConnectSceneGraphInitialized(void* ptr){
	QObject::connect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::sceneGraphInitialized), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_SceneGraphInitialized));;
}

void QQuickWindow_DisconnectSceneGraphInitialized(void* ptr){
	QObject::disconnect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::sceneGraphInitialized), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_SceneGraphInitialized));;
}

void QQuickWindow_ConnectSceneGraphInvalidated(void* ptr){
	QObject::connect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::sceneGraphInvalidated), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_SceneGraphInvalidated));;
}

void QQuickWindow_DisconnectSceneGraphInvalidated(void* ptr){
	QObject::disconnect(static_cast<QQuickWindow*>(ptr), static_cast<void (QQuickWindow::*)()>(&QQuickWindow::sceneGraphInvalidated), static_cast<MyQQuickWindow*>(ptr), static_cast<void (MyQQuickWindow::*)()>(&MyQQuickWindow::Signal_SceneGraphInvalidated));;
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

void QQuickWindow_Update(void* ptr){
	QMetaObject::invokeMethod(static_cast<QQuickWindow*>(ptr), "update");
}

void QQuickWindow_DestroyQQuickWindow(void* ptr){
	static_cast<QQuickWindow*>(ptr)->~QQuickWindow();
}

class MyQSGAbstractRenderer: public QSGAbstractRenderer {
public:
	void Signal_SceneGraphChanged() { callbackQSGAbstractRendererSceneGraphChanged(this->objectName().toUtf8().data()); };
protected:
	void timerEvent(QTimerEvent * event) { if (!callbackQSGAbstractRendererTimerEvent(this->objectName().toUtf8().data(), event)) { QSGAbstractRenderer::timerEvent(event); }; };
	void childEvent(QChildEvent * event) { if (!callbackQSGAbstractRendererChildEvent(this->objectName().toUtf8().data(), event)) { QSGAbstractRenderer::childEvent(event); }; };
	void customEvent(QEvent * event) { if (!callbackQSGAbstractRendererCustomEvent(this->objectName().toUtf8().data(), event)) { QSGAbstractRenderer::customEvent(event); }; };
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

void* QSGClipNode_NewQSGClipNode(){
	return new QSGClipNode();
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

int QSGDynamicTexture_UpdateTexture(void* ptr){
	return static_cast<QSGDynamicTexture*>(ptr)->updateTexture();
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
protected:
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

class MyQSGMaterial: public QSGMaterial {
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
protected:
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
	void activate() { if (!callbackQSGMaterialShaderActivate(this->objectNameAbs().toUtf8().data())) { QSGMaterialShader::activate(); }; };
	void deactivate() { if (!callbackQSGMaterialShaderDeactivate(this->objectNameAbs().toUtf8().data())) { QSGMaterialShader::deactivate(); }; };
protected:
	void compile() { if (!callbackQSGMaterialShaderCompile(this->objectNameAbs().toUtf8().data())) { QSGMaterialShader::compile(); }; };
	void initialize() { if (!callbackQSGMaterialShaderInitialize(this->objectNameAbs().toUtf8().data())) { QSGMaterialShader::initialize(); }; };
};

void QSGMaterialShader_Activate(void* ptr){
	static_cast<QSGMaterialShader*>(ptr)->activate();
}

void QSGMaterialShader_Deactivate(void* ptr){
	static_cast<QSGMaterialShader*>(ptr)->deactivate();
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
	void preprocess() { if (!callbackQSGNodePreprocess(this->objectNameAbs().toUtf8().data())) { QSGNode::preprocess(); }; };
protected:
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
	static_cast<QSGNode*>(ptr)->preprocess();
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

void QSGSimpleRectNode_SetColor(void* ptr, void* color){
	static_cast<QSGSimpleRectNode*>(ptr)->setColor(*static_cast<QColor*>(color));
}

void QSGSimpleRectNode_SetRect(void* ptr, void* rect){
	static_cast<QSGSimpleRectNode*>(ptr)->setRect(*static_cast<QRectF*>(rect));
}

void QSGSimpleRectNode_SetRect2(void* ptr, double x, double y, double w, double h){
	static_cast<QSGSimpleRectNode*>(ptr)->setRect(static_cast<double>(x), static_cast<double>(y), static_cast<double>(w), static_cast<double>(h));
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

void* QSGSimpleTextureNode_Texture(void* ptr){
	return static_cast<QSGSimpleTextureNode*>(ptr)->texture();
}

int QSGSimpleTextureNode_TextureCoordinatesTransform(void* ptr){
	return static_cast<QSGSimpleTextureNode*>(ptr)->textureCoordinatesTransform();
}

void QSGSimpleTextureNode_DestroyQSGSimpleTextureNode(void* ptr){
	static_cast<QSGSimpleTextureNode*>(ptr)->~QSGSimpleTextureNode();
}

class MyQSGTexture: public QSGTexture {
public:
protected:
	void timerEvent(QTimerEvent * event) { if (!callbackQSGTextureTimerEvent(this->objectName().toUtf8().data(), event)) { QSGTexture::timerEvent(event); }; };
	void childEvent(QChildEvent * event) { if (!callbackQSGTextureChildEvent(this->objectName().toUtf8().data(), event)) { QSGTexture::childEvent(event); }; };
	void customEvent(QEvent * event) { if (!callbackQSGTextureCustomEvent(this->objectName().toUtf8().data(), event)) { QSGTexture::customEvent(event); }; };
};

void QSGTexture_Bind(void* ptr){
	static_cast<QSGTexture*>(ptr)->bind();
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

class MyQSGTextureProvider: public QSGTextureProvider {
public:
	void Signal_TextureChanged() { callbackQSGTextureProviderTextureChanged(this->objectName().toUtf8().data()); };
protected:
	void timerEvent(QTimerEvent * event) { if (!callbackQSGTextureProviderTimerEvent(this->objectName().toUtf8().data(), event)) { QSGTextureProvider::timerEvent(event); }; };
	void childEvent(QChildEvent * event) { if (!callbackQSGTextureProviderChildEvent(this->objectName().toUtf8().data(), event)) { QSGTextureProvider::childEvent(event); }; };
	void customEvent(QEvent * event) { if (!callbackQSGTextureProviderCustomEvent(this->objectName().toUtf8().data(), event)) { QSGTextureProvider::customEvent(event); }; };
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

void* QSGTransformNode_NewQSGTransformNode(){
	return new QSGTransformNode();
}

void QSGTransformNode_SetMatrix(void* ptr, void* matrix){
	static_cast<QSGTransformNode*>(ptr)->setMatrix(*static_cast<QMatrix4x4*>(matrix));
}

void QSGTransformNode_DestroyQSGTransformNode(void* ptr){
	static_cast<QSGTransformNode*>(ptr)->~QSGTransformNode();
}

