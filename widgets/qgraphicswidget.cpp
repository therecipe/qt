#include "qgraphicswidget.h"
#include <QObject>
#include <QPalette>
#include <QSize>
#include <QString>
#include <QStyleOptionGraphicsItem>
#include <QWidget>
#include <QSizeF>
#include <QGraphicsLayout>
#include <QFont>
#include <QStyle>
#include <QMetaObject>
#include <QUrl>
#include <QModelIndex>
#include <QRect>
#include <QRectF>
#include <QVariant>
#include <QPainter>
#include <QKeySequence>
#include <QStyleOption>
#include <QAction>
#include <QGraphicsItem>
#include <QGraphicsWidget>
#include "_cgo_export.h"

class MyQGraphicsWidget: public QGraphicsWidget {
public:
void Signal_GeometryChanged(){callbackQGraphicsWidgetGeometryChanged(this->objectName().toUtf8().data());};
};

int QGraphicsWidget_AutoFillBackground(void* ptr){
	return static_cast<QGraphicsWidget*>(ptr)->autoFillBackground();
}

int QGraphicsWidget_FocusPolicy(void* ptr){
	return static_cast<QGraphicsWidget*>(ptr)->focusPolicy();
}

int QGraphicsWidget_LayoutDirection(void* ptr){
	return static_cast<QGraphicsWidget*>(ptr)->layoutDirection();
}

void QGraphicsWidget_Resize(void* ptr, void* size){
	static_cast<QGraphicsWidget*>(ptr)->resize(*static_cast<QSizeF*>(size));
}

void QGraphicsWidget_SetAutoFillBackground(void* ptr, int enabled){
	static_cast<QGraphicsWidget*>(ptr)->setAutoFillBackground(enabled != 0);
}

void QGraphicsWidget_SetFocusPolicy(void* ptr, int policy){
	static_cast<QGraphicsWidget*>(ptr)->setFocusPolicy(static_cast<Qt::FocusPolicy>(policy));
}

void QGraphicsWidget_SetFont(void* ptr, void* font){
	static_cast<QGraphicsWidget*>(ptr)->setFont(*static_cast<QFont*>(font));
}

void QGraphicsWidget_SetGeometry(void* ptr, void* rect){
	static_cast<QGraphicsWidget*>(ptr)->setGeometry(*static_cast<QRectF*>(rect));
}

void QGraphicsWidget_SetLayout(void* ptr, void* layout){
	static_cast<QGraphicsWidget*>(ptr)->setLayout(static_cast<QGraphicsLayout*>(layout));
}

void QGraphicsWidget_SetLayoutDirection(void* ptr, int direction){
	static_cast<QGraphicsWidget*>(ptr)->setLayoutDirection(static_cast<Qt::LayoutDirection>(direction));
}

void QGraphicsWidget_SetPalette(void* ptr, void* palette){
	static_cast<QGraphicsWidget*>(ptr)->setPalette(*static_cast<QPalette*>(palette));
}

void QGraphicsWidget_SetWindowFlags(void* ptr, int wFlags){
	static_cast<QGraphicsWidget*>(ptr)->setWindowFlags(static_cast<Qt::WindowType>(wFlags));
}

void QGraphicsWidget_SetWindowTitle(void* ptr, char* title){
	static_cast<QGraphicsWidget*>(ptr)->setWindowTitle(QString(title));
}

void QGraphicsWidget_UnsetLayoutDirection(void* ptr){
	static_cast<QGraphicsWidget*>(ptr)->unsetLayoutDirection();
}

int QGraphicsWidget_WindowFlags(void* ptr){
	return static_cast<QGraphicsWidget*>(ptr)->windowFlags();
}

char* QGraphicsWidget_WindowTitle(void* ptr){
	return static_cast<QGraphicsWidget*>(ptr)->windowTitle().toUtf8().data();
}

void* QGraphicsWidget_NewQGraphicsWidget(void* parent, int wFlags){
	return new QGraphicsWidget(static_cast<QGraphicsItem*>(parent), static_cast<Qt::WindowType>(wFlags));
}

void QGraphicsWidget_AddAction(void* ptr, void* action){
	static_cast<QGraphicsWidget*>(ptr)->addAction(static_cast<QAction*>(action));
}

void QGraphicsWidget_AdjustSize(void* ptr){
	static_cast<QGraphicsWidget*>(ptr)->adjustSize();
}

int QGraphicsWidget_Close(void* ptr){
	return QMetaObject::invokeMethod(static_cast<QGraphicsWidget*>(ptr), "close");
}

void* QGraphicsWidget_FocusWidget(void* ptr){
	return static_cast<QGraphicsWidget*>(ptr)->focusWidget();
}

void QGraphicsWidget_ConnectGeometryChanged(void* ptr){
	QObject::connect(static_cast<QGraphicsWidget*>(ptr), static_cast<void (QGraphicsWidget::*)()>(&QGraphicsWidget::geometryChanged), static_cast<MyQGraphicsWidget*>(ptr), static_cast<void (MyQGraphicsWidget::*)()>(&MyQGraphicsWidget::Signal_GeometryChanged));;
}

void QGraphicsWidget_DisconnectGeometryChanged(void* ptr){
	QObject::disconnect(static_cast<QGraphicsWidget*>(ptr), static_cast<void (QGraphicsWidget::*)()>(&QGraphicsWidget::geometryChanged), static_cast<MyQGraphicsWidget*>(ptr), static_cast<void (MyQGraphicsWidget::*)()>(&MyQGraphicsWidget::Signal_GeometryChanged));;
}

int QGraphicsWidget_GrabShortcut(void* ptr, void* sequence, int context){
	return static_cast<QGraphicsWidget*>(ptr)->grabShortcut(*static_cast<QKeySequence*>(sequence), static_cast<Qt::ShortcutContext>(context));
}

void QGraphicsWidget_InsertAction(void* ptr, void* before, void* action){
	static_cast<QGraphicsWidget*>(ptr)->insertAction(static_cast<QAction*>(before), static_cast<QAction*>(action));
}

int QGraphicsWidget_IsActiveWindow(void* ptr){
	return static_cast<QGraphicsWidget*>(ptr)->isActiveWindow();
}

void* QGraphicsWidget_Layout(void* ptr){
	return static_cast<QGraphicsWidget*>(ptr)->layout();
}

void QGraphicsWidget_Paint(void* ptr, void* painter, void* option, void* widget){
	static_cast<QGraphicsWidget*>(ptr)->paint(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
}

void QGraphicsWidget_PaintWindowFrame(void* ptr, void* painter, void* option, void* widget){
	static_cast<QGraphicsWidget*>(ptr)->paintWindowFrame(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
}

void QGraphicsWidget_ReleaseShortcut(void* ptr, int id){
	static_cast<QGraphicsWidget*>(ptr)->releaseShortcut(id);
}

void QGraphicsWidget_RemoveAction(void* ptr, void* action){
	static_cast<QGraphicsWidget*>(ptr)->removeAction(static_cast<QAction*>(action));
}

void QGraphicsWidget_Resize2(void* ptr, double w, double h){
	static_cast<QGraphicsWidget*>(ptr)->resize(static_cast<qreal>(w), static_cast<qreal>(h));
}

void QGraphicsWidget_SetAttribute(void* ptr, int attribute, int on){
	static_cast<QGraphicsWidget*>(ptr)->setAttribute(static_cast<Qt::WidgetAttribute>(attribute), on != 0);
}

void QGraphicsWidget_SetContentsMargins(void* ptr, double left, double top, double right, double bottom){
	static_cast<QGraphicsWidget*>(ptr)->setContentsMargins(static_cast<qreal>(left), static_cast<qreal>(top), static_cast<qreal>(right), static_cast<qreal>(bottom));
}

void QGraphicsWidget_SetGeometry2(void* ptr, double x, double y, double w, double h){
	static_cast<QGraphicsWidget*>(ptr)->setGeometry(static_cast<qreal>(x), static_cast<qreal>(y), static_cast<qreal>(w), static_cast<qreal>(h));
}

void QGraphicsWidget_SetShortcutAutoRepeat(void* ptr, int id, int enabled){
	static_cast<QGraphicsWidget*>(ptr)->setShortcutAutoRepeat(id, enabled != 0);
}

void QGraphicsWidget_SetShortcutEnabled(void* ptr, int id, int enabled){
	static_cast<QGraphicsWidget*>(ptr)->setShortcutEnabled(id, enabled != 0);
}

void QGraphicsWidget_SetStyle(void* ptr, void* style){
	static_cast<QGraphicsWidget*>(ptr)->setStyle(static_cast<QStyle*>(style));
}

void QGraphicsWidget_QGraphicsWidget_SetTabOrder(void* first, void* second){
	QGraphicsWidget::setTabOrder(static_cast<QGraphicsWidget*>(first), static_cast<QGraphicsWidget*>(second));
}

void QGraphicsWidget_SetWindowFrameMargins(void* ptr, double left, double top, double right, double bottom){
	static_cast<QGraphicsWidget*>(ptr)->setWindowFrameMargins(static_cast<qreal>(left), static_cast<qreal>(top), static_cast<qreal>(right), static_cast<qreal>(bottom));
}

void* QGraphicsWidget_Style(void* ptr){
	return static_cast<QGraphicsWidget*>(ptr)->style();
}

int QGraphicsWidget_TestAttribute(void* ptr, int attribute){
	return static_cast<QGraphicsWidget*>(ptr)->testAttribute(static_cast<Qt::WidgetAttribute>(attribute));
}

int QGraphicsWidget_Type(void* ptr){
	return static_cast<QGraphicsWidget*>(ptr)->type();
}

void QGraphicsWidget_UnsetWindowFrameMargins(void* ptr){
	static_cast<QGraphicsWidget*>(ptr)->unsetWindowFrameMargins();
}

int QGraphicsWidget_WindowType(void* ptr){
	return static_cast<QGraphicsWidget*>(ptr)->windowType();
}

void QGraphicsWidget_DestroyQGraphicsWidget(void* ptr){
	static_cast<QGraphicsWidget*>(ptr)->~QGraphicsWidget();
}

