#include "qquickitem.h"
#include <QObject>
#include <QCursor>
#include <QPointF>
#include <QQuickWindow>
#include <QMetaObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QPoint>
#include <QQuickItem>
#include "_cgo_export.h"

class MyQQuickItem: public QQuickItem {
public:
void Signal_WindowChanged(QQuickWindow * window){callbackQQuickItemWindowChanged(this->objectName().toUtf8().data(), window);};
};

QtObjectPtr QQuickItem_NewQQuickItem(QtObjectPtr parent){
	return new QQuickItem(static_cast<QQuickItem*>(parent));
}

int QQuickItem_ActiveFocusOnTab(QtObjectPtr ptr){
	return static_cast<QQuickItem*>(ptr)->activeFocusOnTab();
}

int QQuickItem_Antialiasing(QtObjectPtr ptr){
	return static_cast<QQuickItem*>(ptr)->antialiasing();
}

int QQuickItem_Clip(QtObjectPtr ptr){
	return static_cast<QQuickItem*>(ptr)->clip();
}

int QQuickItem_HasActiveFocus(QtObjectPtr ptr){
	return static_cast<QQuickItem*>(ptr)->hasActiveFocus();
}

int QQuickItem_HasFocus(QtObjectPtr ptr){
	return static_cast<QQuickItem*>(ptr)->hasFocus();
}

int QQuickItem_IsEnabled(QtObjectPtr ptr){
	return static_cast<QQuickItem*>(ptr)->isEnabled();
}

int QQuickItem_IsTextureProvider(QtObjectPtr ptr){
	return static_cast<QQuickItem*>(ptr)->isTextureProvider();
}

int QQuickItem_IsVisible(QtObjectPtr ptr){
	return static_cast<QQuickItem*>(ptr)->isVisible();
}

QtObjectPtr QQuickItem_ParentItem(QtObjectPtr ptr){
	return static_cast<QQuickItem*>(ptr)->parentItem();
}

void QQuickItem_ResetAntialiasing(QtObjectPtr ptr){
	static_cast<QQuickItem*>(ptr)->resetAntialiasing();
}

void QQuickItem_ResetHeight(QtObjectPtr ptr){
	static_cast<QQuickItem*>(ptr)->resetHeight();
}

void QQuickItem_ResetWidth(QtObjectPtr ptr){
	static_cast<QQuickItem*>(ptr)->resetWidth();
}

void QQuickItem_SetActiveFocusOnTab(QtObjectPtr ptr, int v){
	static_cast<QQuickItem*>(ptr)->setActiveFocusOnTab(v != 0);
}

void QQuickItem_SetAntialiasing(QtObjectPtr ptr, int v){
	static_cast<QQuickItem*>(ptr)->setAntialiasing(v != 0);
}

void QQuickItem_SetClip(QtObjectPtr ptr, int v){
	static_cast<QQuickItem*>(ptr)->setClip(v != 0);
}

void QQuickItem_SetEnabled(QtObjectPtr ptr, int v){
	static_cast<QQuickItem*>(ptr)->setEnabled(v != 0);
}

void QQuickItem_SetFocus(QtObjectPtr ptr, int v){
	static_cast<QQuickItem*>(ptr)->setFocus(v != 0);
}

void QQuickItem_SetFocus2(QtObjectPtr ptr, int focus, int reason){
	static_cast<QQuickItem*>(ptr)->setFocus(focus != 0, static_cast<Qt::FocusReason>(reason));
}

void QQuickItem_SetParentItem(QtObjectPtr ptr, QtObjectPtr parent){
	static_cast<QQuickItem*>(ptr)->setParentItem(static_cast<QQuickItem*>(parent));
}

void QQuickItem_SetSmooth(QtObjectPtr ptr, int v){
	static_cast<QQuickItem*>(ptr)->setSmooth(v != 0);
}

void QQuickItem_SetState(QtObjectPtr ptr, char* v){
	static_cast<QQuickItem*>(ptr)->setState(QString(v));
}

void QQuickItem_SetTransformOrigin(QtObjectPtr ptr, int v){
	static_cast<QQuickItem*>(ptr)->setTransformOrigin(static_cast<QQuickItem::TransformOrigin>(v));
}

void QQuickItem_SetVisible(QtObjectPtr ptr, int v){
	static_cast<QQuickItem*>(ptr)->setVisible(v != 0);
}

int QQuickItem_Smooth(QtObjectPtr ptr){
	return static_cast<QQuickItem*>(ptr)->smooth();
}

QtObjectPtr QQuickItem_TextureProvider(QtObjectPtr ptr){
	return static_cast<QQuickItem*>(ptr)->textureProvider();
}

int QQuickItem_TransformOrigin(QtObjectPtr ptr){
	return static_cast<QQuickItem*>(ptr)->transformOrigin();
}

int QQuickItem_AcceptHoverEvents(QtObjectPtr ptr){
	return static_cast<QQuickItem*>(ptr)->acceptHoverEvents();
}

int QQuickItem_AcceptedMouseButtons(QtObjectPtr ptr){
	return static_cast<QQuickItem*>(ptr)->acceptedMouseButtons();
}

int QQuickItem_Contains(QtObjectPtr ptr, QtObjectPtr point){
	return static_cast<QQuickItem*>(ptr)->contains(*static_cast<QPointF*>(point));
}

int QQuickItem_FiltersChildMouseEvents(QtObjectPtr ptr){
	return static_cast<QQuickItem*>(ptr)->filtersChildMouseEvents();
}

int QQuickItem_Flags(QtObjectPtr ptr){
	return static_cast<QQuickItem*>(ptr)->flags();
}

void QQuickItem_ForceActiveFocus(QtObjectPtr ptr){
	static_cast<QQuickItem*>(ptr)->forceActiveFocus();
}

void QQuickItem_ForceActiveFocus2(QtObjectPtr ptr, int reason){
	static_cast<QQuickItem*>(ptr)->forceActiveFocus(static_cast<Qt::FocusReason>(reason));
}

void QQuickItem_GrabMouse(QtObjectPtr ptr){
	static_cast<QQuickItem*>(ptr)->grabMouse();
}

char* QQuickItem_InputMethodQuery(QtObjectPtr ptr, int query){
	return static_cast<QQuickItem*>(ptr)->inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)).toString().toUtf8().data();
}

int QQuickItem_IsFocusScope(QtObjectPtr ptr){
	return static_cast<QQuickItem*>(ptr)->isFocusScope();
}

int QQuickItem_KeepMouseGrab(QtObjectPtr ptr){
	return static_cast<QQuickItem*>(ptr)->keepMouseGrab();
}

int QQuickItem_KeepTouchGrab(QtObjectPtr ptr){
	return static_cast<QQuickItem*>(ptr)->keepTouchGrab();
}

QtObjectPtr QQuickItem_NextItemInFocusChain(QtObjectPtr ptr, int forward){
	return static_cast<QQuickItem*>(ptr)->nextItemInFocusChain(forward != 0);
}

void QQuickItem_Polish(QtObjectPtr ptr){
	static_cast<QQuickItem*>(ptr)->polish();
}

QtObjectPtr QQuickItem_ScopedFocusItem(QtObjectPtr ptr){
	return static_cast<QQuickItem*>(ptr)->scopedFocusItem();
}

void QQuickItem_SetAcceptHoverEvents(QtObjectPtr ptr, int enabled){
	static_cast<QQuickItem*>(ptr)->setAcceptHoverEvents(enabled != 0);
}

void QQuickItem_SetAcceptedMouseButtons(QtObjectPtr ptr, int buttons){
	static_cast<QQuickItem*>(ptr)->setAcceptedMouseButtons(static_cast<Qt::MouseButton>(buttons));
}

void QQuickItem_SetCursor(QtObjectPtr ptr, QtObjectPtr cursor){
	static_cast<QQuickItem*>(ptr)->setCursor(*static_cast<QCursor*>(cursor));
}

void QQuickItem_SetFiltersChildMouseEvents(QtObjectPtr ptr, int filter){
	static_cast<QQuickItem*>(ptr)->setFiltersChildMouseEvents(filter != 0);
}

void QQuickItem_SetFlag(QtObjectPtr ptr, int flag, int enabled){
	static_cast<QQuickItem*>(ptr)->setFlag(static_cast<QQuickItem::Flag>(flag), enabled != 0);
}

void QQuickItem_SetFlags(QtObjectPtr ptr, int flags){
	static_cast<QQuickItem*>(ptr)->setFlags(static_cast<QQuickItem::Flag>(flags));
}

void QQuickItem_SetKeepMouseGrab(QtObjectPtr ptr, int keep){
	static_cast<QQuickItem*>(ptr)->setKeepMouseGrab(keep != 0);
}

void QQuickItem_SetKeepTouchGrab(QtObjectPtr ptr, int keep){
	static_cast<QQuickItem*>(ptr)->setKeepTouchGrab(keep != 0);
}

void QQuickItem_StackAfter(QtObjectPtr ptr, QtObjectPtr sibling){
	static_cast<QQuickItem*>(ptr)->stackAfter(static_cast<QQuickItem*>(sibling));
}

void QQuickItem_StackBefore(QtObjectPtr ptr, QtObjectPtr sibling){
	static_cast<QQuickItem*>(ptr)->stackBefore(static_cast<QQuickItem*>(sibling));
}

void QQuickItem_UngrabMouse(QtObjectPtr ptr){
	static_cast<QQuickItem*>(ptr)->ungrabMouse();
}

void QQuickItem_UngrabTouchPoints(QtObjectPtr ptr){
	static_cast<QQuickItem*>(ptr)->ungrabTouchPoints();
}

void QQuickItem_UnsetCursor(QtObjectPtr ptr){
	static_cast<QQuickItem*>(ptr)->unsetCursor();
}

void QQuickItem_Update(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QQuickItem*>(ptr), "update");
}

QtObjectPtr QQuickItem_Window(QtObjectPtr ptr){
	return static_cast<QQuickItem*>(ptr)->window();
}

void QQuickItem_ConnectWindowChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QQuickItem*>(ptr), static_cast<void (QQuickItem::*)(QQuickWindow *)>(&QQuickItem::windowChanged), static_cast<MyQQuickItem*>(ptr), static_cast<void (MyQQuickItem::*)(QQuickWindow *)>(&MyQQuickItem::Signal_WindowChanged));;
}

void QQuickItem_DisconnectWindowChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QQuickItem*>(ptr), static_cast<void (QQuickItem::*)(QQuickWindow *)>(&QQuickItem::windowChanged), static_cast<MyQQuickItem*>(ptr), static_cast<void (MyQQuickItem::*)(QQuickWindow *)>(&MyQQuickItem::Signal_WindowChanged));;
}

void QQuickItem_DestroyQQuickItem(QtObjectPtr ptr){
	static_cast<QQuickItem*>(ptr)->~QQuickItem();
}

