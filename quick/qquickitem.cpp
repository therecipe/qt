#include "qquickitem.h"
#include <QUrl>
#include <QQuickWindow>
#include <QObject>
#include <QPoint>
#include <QCursor>
#include <QMetaObject>
#include <QPointF>
#include <QString>
#include <QVariant>
#include <QModelIndex>
#include <QQuickItem>
#include "_cgo_export.h"

class MyQQuickItem: public QQuickItem {
public:
void Signal_WindowChanged(QQuickWindow * window){callbackQQuickItemWindowChanged(this->objectName().toUtf8().data(), window);};
};

void* QQuickItem_NewQQuickItem(void* parent){
	return new QQuickItem(static_cast<QQuickItem*>(parent));
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
	static_cast<QQuickItem*>(ptr)->setBaselineOffset(static_cast<qreal>(v));
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
	static_cast<QQuickItem*>(ptr)->setHeight(static_cast<qreal>(v));
}

void QQuickItem_SetImplicitHeight(void* ptr, double v){
	static_cast<QQuickItem*>(ptr)->setImplicitHeight(static_cast<qreal>(v));
}

void QQuickItem_SetImplicitWidth(void* ptr, double v){
	static_cast<QQuickItem*>(ptr)->setImplicitWidth(static_cast<qreal>(v));
}

void QQuickItem_SetOpacity(void* ptr, double v){
	static_cast<QQuickItem*>(ptr)->setOpacity(static_cast<qreal>(v));
}

void QQuickItem_SetParentItem(void* ptr, void* parent){
	static_cast<QQuickItem*>(ptr)->setParentItem(static_cast<QQuickItem*>(parent));
}

void QQuickItem_SetRotation(void* ptr, double v){
	static_cast<QQuickItem*>(ptr)->setRotation(static_cast<qreal>(v));
}

void QQuickItem_SetScale(void* ptr, double v){
	static_cast<QQuickItem*>(ptr)->setScale(static_cast<qreal>(v));
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
	static_cast<QQuickItem*>(ptr)->setWidth(static_cast<qreal>(v));
}

void QQuickItem_SetX(void* ptr, double v){
	static_cast<QQuickItem*>(ptr)->setX(static_cast<qreal>(v));
}

void QQuickItem_SetY(void* ptr, double v){
	static_cast<QQuickItem*>(ptr)->setY(static_cast<qreal>(v));
}

void QQuickItem_SetZ(void* ptr, double v){
	static_cast<QQuickItem*>(ptr)->setZ(static_cast<qreal>(v));
}

int QQuickItem_Smooth(void* ptr){
	return static_cast<QQuickItem*>(ptr)->smooth();
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
	return static_cast<QQuickItem*>(ptr)->childAt(static_cast<qreal>(x), static_cast<qreal>(y));
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

