#include "qgraphicswidget.h"
#include <QStyleOptionGraphicsItem>
#include <QStyle>
#include <QVariant>
#include <QModelIndex>
#include <QSize>
#include <QRectF>
#include <QStyleOption>
#include <QSizeF>
#include <QMetaObject>
#include <QPainter>
#include <QPalette>
#include <QGraphicsLayout>
#include <QUrl>
#include <QAction>
#include <QRect>
#include <QGraphicsItem>
#include <QString>
#include <QObject>
#include <QFont>
#include <QWidget>
#include <QKeySequence>
#include <QGraphicsWidget>
#include "_cgo_export.h"

class MyQGraphicsWidget: public QGraphicsWidget {
public:
void Signal_GeometryChanged(){callbackQGraphicsWidgetGeometryChanged(this->objectName().toUtf8().data());};
};

int QGraphicsWidget_AutoFillBackground(QtObjectPtr ptr){
	return static_cast<QGraphicsWidget*>(ptr)->autoFillBackground();
}

int QGraphicsWidget_FocusPolicy(QtObjectPtr ptr){
	return static_cast<QGraphicsWidget*>(ptr)->focusPolicy();
}

int QGraphicsWidget_LayoutDirection(QtObjectPtr ptr){
	return static_cast<QGraphicsWidget*>(ptr)->layoutDirection();
}

void QGraphicsWidget_Resize(QtObjectPtr ptr, QtObjectPtr size){
	static_cast<QGraphicsWidget*>(ptr)->resize(*static_cast<QSizeF*>(size));
}

void QGraphicsWidget_SetAutoFillBackground(QtObjectPtr ptr, int enabled){
	static_cast<QGraphicsWidget*>(ptr)->setAutoFillBackground(enabled != 0);
}

void QGraphicsWidget_SetFocusPolicy(QtObjectPtr ptr, int policy){
	static_cast<QGraphicsWidget*>(ptr)->setFocusPolicy(static_cast<Qt::FocusPolicy>(policy));
}

void QGraphicsWidget_SetFont(QtObjectPtr ptr, QtObjectPtr font){
	static_cast<QGraphicsWidget*>(ptr)->setFont(*static_cast<QFont*>(font));
}

void QGraphicsWidget_SetGeometry(QtObjectPtr ptr, QtObjectPtr rect){
	static_cast<QGraphicsWidget*>(ptr)->setGeometry(*static_cast<QRectF*>(rect));
}

void QGraphicsWidget_SetLayout(QtObjectPtr ptr, QtObjectPtr layout){
	static_cast<QGraphicsWidget*>(ptr)->setLayout(static_cast<QGraphicsLayout*>(layout));
}

void QGraphicsWidget_SetLayoutDirection(QtObjectPtr ptr, int direction){
	static_cast<QGraphicsWidget*>(ptr)->setLayoutDirection(static_cast<Qt::LayoutDirection>(direction));
}

void QGraphicsWidget_SetPalette(QtObjectPtr ptr, QtObjectPtr palette){
	static_cast<QGraphicsWidget*>(ptr)->setPalette(*static_cast<QPalette*>(palette));
}

void QGraphicsWidget_SetWindowFlags(QtObjectPtr ptr, int wFlags){
	static_cast<QGraphicsWidget*>(ptr)->setWindowFlags(static_cast<Qt::WindowType>(wFlags));
}

void QGraphicsWidget_SetWindowTitle(QtObjectPtr ptr, char* title){
	static_cast<QGraphicsWidget*>(ptr)->setWindowTitle(QString(title));
}

void QGraphicsWidget_UnsetLayoutDirection(QtObjectPtr ptr){
	static_cast<QGraphicsWidget*>(ptr)->unsetLayoutDirection();
}

int QGraphicsWidget_WindowFlags(QtObjectPtr ptr){
	return static_cast<QGraphicsWidget*>(ptr)->windowFlags();
}

char* QGraphicsWidget_WindowTitle(QtObjectPtr ptr){
	return static_cast<QGraphicsWidget*>(ptr)->windowTitle().toUtf8().data();
}

QtObjectPtr QGraphicsWidget_NewQGraphicsWidget(QtObjectPtr parent, int wFlags){
	return new QGraphicsWidget(static_cast<QGraphicsItem*>(parent), static_cast<Qt::WindowType>(wFlags));
}

void QGraphicsWidget_AddAction(QtObjectPtr ptr, QtObjectPtr action){
	static_cast<QGraphicsWidget*>(ptr)->addAction(static_cast<QAction*>(action));
}

void QGraphicsWidget_AdjustSize(QtObjectPtr ptr){
	static_cast<QGraphicsWidget*>(ptr)->adjustSize();
}

int QGraphicsWidget_Close(QtObjectPtr ptr){
	return QMetaObject::invokeMethod(static_cast<QGraphicsWidget*>(ptr), "close");
}

QtObjectPtr QGraphicsWidget_FocusWidget(QtObjectPtr ptr){
	return static_cast<QGraphicsWidget*>(ptr)->focusWidget();
}

void QGraphicsWidget_ConnectGeometryChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QGraphicsWidget*>(ptr), static_cast<void (QGraphicsWidget::*)()>(&QGraphicsWidget::geometryChanged), static_cast<MyQGraphicsWidget*>(ptr), static_cast<void (MyQGraphicsWidget::*)()>(&MyQGraphicsWidget::Signal_GeometryChanged));;
}

void QGraphicsWidget_DisconnectGeometryChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QGraphicsWidget*>(ptr), static_cast<void (QGraphicsWidget::*)()>(&QGraphicsWidget::geometryChanged), static_cast<MyQGraphicsWidget*>(ptr), static_cast<void (MyQGraphicsWidget::*)()>(&MyQGraphicsWidget::Signal_GeometryChanged));;
}

int QGraphicsWidget_GrabShortcut(QtObjectPtr ptr, QtObjectPtr sequence, int context){
	return static_cast<QGraphicsWidget*>(ptr)->grabShortcut(*static_cast<QKeySequence*>(sequence), static_cast<Qt::ShortcutContext>(context));
}

void QGraphicsWidget_InsertAction(QtObjectPtr ptr, QtObjectPtr before, QtObjectPtr action){
	static_cast<QGraphicsWidget*>(ptr)->insertAction(static_cast<QAction*>(before), static_cast<QAction*>(action));
}

int QGraphicsWidget_IsActiveWindow(QtObjectPtr ptr){
	return static_cast<QGraphicsWidget*>(ptr)->isActiveWindow();
}

QtObjectPtr QGraphicsWidget_Layout(QtObjectPtr ptr){
	return static_cast<QGraphicsWidget*>(ptr)->layout();
}

void QGraphicsWidget_Paint(QtObjectPtr ptr, QtObjectPtr painter, QtObjectPtr option, QtObjectPtr widget){
	static_cast<QGraphicsWidget*>(ptr)->paint(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
}

void QGraphicsWidget_PaintWindowFrame(QtObjectPtr ptr, QtObjectPtr painter, QtObjectPtr option, QtObjectPtr widget){
	static_cast<QGraphicsWidget*>(ptr)->paintWindowFrame(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
}

void QGraphicsWidget_ReleaseShortcut(QtObjectPtr ptr, int id){
	static_cast<QGraphicsWidget*>(ptr)->releaseShortcut(id);
}

void QGraphicsWidget_RemoveAction(QtObjectPtr ptr, QtObjectPtr action){
	static_cast<QGraphicsWidget*>(ptr)->removeAction(static_cast<QAction*>(action));
}

void QGraphicsWidget_SetAttribute(QtObjectPtr ptr, int attribute, int on){
	static_cast<QGraphicsWidget*>(ptr)->setAttribute(static_cast<Qt::WidgetAttribute>(attribute), on != 0);
}

void QGraphicsWidget_SetShortcutAutoRepeat(QtObjectPtr ptr, int id, int enabled){
	static_cast<QGraphicsWidget*>(ptr)->setShortcutAutoRepeat(id, enabled != 0);
}

void QGraphicsWidget_SetShortcutEnabled(QtObjectPtr ptr, int id, int enabled){
	static_cast<QGraphicsWidget*>(ptr)->setShortcutEnabled(id, enabled != 0);
}

void QGraphicsWidget_SetStyle(QtObjectPtr ptr, QtObjectPtr style){
	static_cast<QGraphicsWidget*>(ptr)->setStyle(static_cast<QStyle*>(style));
}

void QGraphicsWidget_QGraphicsWidget_SetTabOrder(QtObjectPtr first, QtObjectPtr second){
	QGraphicsWidget::setTabOrder(static_cast<QGraphicsWidget*>(first), static_cast<QGraphicsWidget*>(second));
}

QtObjectPtr QGraphicsWidget_Style(QtObjectPtr ptr){
	return static_cast<QGraphicsWidget*>(ptr)->style();
}

int QGraphicsWidget_TestAttribute(QtObjectPtr ptr, int attribute){
	return static_cast<QGraphicsWidget*>(ptr)->testAttribute(static_cast<Qt::WidgetAttribute>(attribute));
}

int QGraphicsWidget_Type(QtObjectPtr ptr){
	return static_cast<QGraphicsWidget*>(ptr)->type();
}

void QGraphicsWidget_UnsetWindowFrameMargins(QtObjectPtr ptr){
	static_cast<QGraphicsWidget*>(ptr)->unsetWindowFrameMargins();
}

int QGraphicsWidget_WindowType(QtObjectPtr ptr){
	return static_cast<QGraphicsWidget*>(ptr)->windowType();
}

void QGraphicsWidget_DestroyQGraphicsWidget(QtObjectPtr ptr){
	static_cast<QGraphicsWidget*>(ptr)->~QGraphicsWidget();
}

