#include "qwindow.h"
#include <QSurface>
#include <QMetaObject>
#include <QObject>
#include <QPoint>
#include <QUrl>
#include <QCursor>
#include <QRegion>
#include <QIcon>
#include <QSize>
#include <QSurfaceFormat>
#include <QString>
#include <QVariant>
#include <QModelIndex>
#include <QScreen>
#include <QRect>
#include <QWindow>
#include "_cgo_export.h"

class MyQWindow: public QWindow {
public:
void Signal_ActiveChanged(){callbackQWindowActiveChanged(this->objectName().toUtf8().data());};
void Signal_ContentOrientationChanged(Qt::ScreenOrientation orientation){callbackQWindowContentOrientationChanged(this->objectName().toUtf8().data(), orientation);};
void Signal_FocusObjectChanged(QObject * object){callbackQWindowFocusObjectChanged(this->objectName().toUtf8().data(), object);};
void Signal_HeightChanged(int arg){callbackQWindowHeightChanged(this->objectName().toUtf8().data(), arg);};
void Signal_MaximumHeightChanged(int arg){callbackQWindowMaximumHeightChanged(this->objectName().toUtf8().data(), arg);};
void Signal_MaximumWidthChanged(int arg){callbackQWindowMaximumWidthChanged(this->objectName().toUtf8().data(), arg);};
void Signal_MinimumHeightChanged(int arg){callbackQWindowMinimumHeightChanged(this->objectName().toUtf8().data(), arg);};
void Signal_MinimumWidthChanged(int arg){callbackQWindowMinimumWidthChanged(this->objectName().toUtf8().data(), arg);};
void Signal_ModalityChanged(Qt::WindowModality modality){callbackQWindowModalityChanged(this->objectName().toUtf8().data(), modality);};
void Signal_ScreenChanged(QScreen * screen){callbackQWindowScreenChanged(this->objectName().toUtf8().data(), screen);};
void Signal_VisibilityChanged(QWindow::Visibility visibility){callbackQWindowVisibilityChanged(this->objectName().toUtf8().data(), visibility);};
void Signal_VisibleChanged(bool arg){callbackQWindowVisibleChanged(this->objectName().toUtf8().data(), arg);};
void Signal_WidthChanged(int arg){callbackQWindowWidthChanged(this->objectName().toUtf8().data(), arg);};
void Signal_WindowStateChanged(Qt::WindowState windowState){callbackQWindowWindowStateChanged(this->objectName().toUtf8().data(), windowState);};
void Signal_WindowTitleChanged(const QString & title){callbackQWindowWindowTitleChanged(this->objectName().toUtf8().data(), title.toUtf8().data());};
void Signal_XChanged(int arg){callbackQWindowXChanged(this->objectName().toUtf8().data(), arg);};
void Signal_YChanged(int arg){callbackQWindowYChanged(this->objectName().toUtf8().data(), arg);};
};

int QWindow_ContentOrientation(void* ptr){
	return static_cast<QWindow*>(ptr)->contentOrientation();
}

int QWindow_Flags(void* ptr){
	return static_cast<QWindow*>(ptr)->flags();
}

int QWindow_IsVisible(void* ptr){
	return static_cast<QWindow*>(ptr)->isVisible();
}

int QWindow_Modality(void* ptr){
	return static_cast<QWindow*>(ptr)->modality();
}

double QWindow_Opacity(void* ptr){
	return static_cast<double>(static_cast<QWindow*>(ptr)->opacity());
}

void QWindow_ReportContentOrientationChange(void* ptr, int orientation){
	static_cast<QWindow*>(ptr)->reportContentOrientationChange(static_cast<Qt::ScreenOrientation>(orientation));
}

void QWindow_SetFlags(void* ptr, int flags){
	static_cast<QWindow*>(ptr)->setFlags(static_cast<Qt::WindowType>(flags));
}

void QWindow_SetHeight(void* ptr, int arg){
	QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "setHeight", Q_ARG(int, arg));
}

void QWindow_SetMaximumHeight(void* ptr, int h){
	QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "setMaximumHeight", Q_ARG(int, h));
}

void QWindow_SetMaximumWidth(void* ptr, int w){
	QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "setMaximumWidth", Q_ARG(int, w));
}

void QWindow_SetMinimumHeight(void* ptr, int h){
	QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "setMinimumHeight", Q_ARG(int, h));
}

void QWindow_SetMinimumWidth(void* ptr, int w){
	QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "setMinimumWidth", Q_ARG(int, w));
}

void QWindow_SetModality(void* ptr, int modality){
	static_cast<QWindow*>(ptr)->setModality(static_cast<Qt::WindowModality>(modality));
}

void QWindow_SetOpacity(void* ptr, double level){
	static_cast<QWindow*>(ptr)->setOpacity(static_cast<qreal>(level));
}

void QWindow_SetTitle(void* ptr, char* v){
	QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "setTitle", Q_ARG(QString, QString(v)));
}

void QWindow_SetVisibility(void* ptr, int v){
	static_cast<QWindow*>(ptr)->setVisibility(static_cast<QWindow::Visibility>(v));
}

void QWindow_SetVisible(void* ptr, int visible){
	QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "setVisible", Q_ARG(bool, visible != 0));
}

void QWindow_SetWidth(void* ptr, int arg){
	QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "setWidth", Q_ARG(int, arg));
}

void QWindow_SetX(void* ptr, int arg){
	QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "setX", Q_ARG(int, arg));
}

void QWindow_SetY(void* ptr, int arg){
	QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "setY", Q_ARG(int, arg));
}

char* QWindow_Title(void* ptr){
	return static_cast<QWindow*>(ptr)->title().toUtf8().data();
}

int QWindow_Visibility(void* ptr){
	return static_cast<QWindow*>(ptr)->visibility();
}

void* QWindow_NewQWindow(void* targetScreen){
	return new QWindow(static_cast<QScreen*>(targetScreen));
}

void* QWindow_NewQWindow2(void* parent){
	return new QWindow(static_cast<QWindow*>(parent));
}

void QWindow_ConnectActiveChanged(void* ptr){
	QObject::connect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)()>(&QWindow::activeChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)()>(&MyQWindow::Signal_ActiveChanged));;
}

void QWindow_DisconnectActiveChanged(void* ptr){
	QObject::disconnect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)()>(&QWindow::activeChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)()>(&MyQWindow::Signal_ActiveChanged));;
}

void QWindow_Alert(void* ptr, int msec){
	QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "alert", Q_ARG(int, msec));
}

int QWindow_Close(void* ptr){
	return QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "close");
}

void QWindow_ConnectContentOrientationChanged(void* ptr){
	QObject::connect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(Qt::ScreenOrientation)>(&QWindow::contentOrientationChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(Qt::ScreenOrientation)>(&MyQWindow::Signal_ContentOrientationChanged));;
}

void QWindow_DisconnectContentOrientationChanged(void* ptr){
	QObject::disconnect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(Qt::ScreenOrientation)>(&QWindow::contentOrientationChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(Qt::ScreenOrientation)>(&MyQWindow::Signal_ContentOrientationChanged));;
}

void QWindow_Create(void* ptr){
	static_cast<QWindow*>(ptr)->create();
}

void QWindow_Destroy(void* ptr){
	static_cast<QWindow*>(ptr)->destroy();
}

double QWindow_DevicePixelRatio(void* ptr){
	return static_cast<double>(static_cast<QWindow*>(ptr)->devicePixelRatio());
}

char* QWindow_FilePath(void* ptr){
	return static_cast<QWindow*>(ptr)->filePath().toUtf8().data();
}

void* QWindow_FocusObject(void* ptr){
	return static_cast<QWindow*>(ptr)->focusObject();
}

void QWindow_ConnectFocusObjectChanged(void* ptr){
	QObject::connect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(QObject *)>(&QWindow::focusObjectChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(QObject *)>(&MyQWindow::Signal_FocusObjectChanged));;
}

void QWindow_DisconnectFocusObjectChanged(void* ptr){
	QObject::disconnect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(QObject *)>(&QWindow::focusObjectChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(QObject *)>(&MyQWindow::Signal_FocusObjectChanged));;
}

int QWindow_Height(void* ptr){
	return static_cast<QWindow*>(ptr)->height();
}

void QWindow_ConnectHeightChanged(void* ptr){
	QObject::connect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(int)>(&QWindow::heightChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(int)>(&MyQWindow::Signal_HeightChanged));;
}

void QWindow_DisconnectHeightChanged(void* ptr){
	QObject::disconnect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(int)>(&QWindow::heightChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(int)>(&MyQWindow::Signal_HeightChanged));;
}

void QWindow_Hide(void* ptr){
	QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "hide");
}

int QWindow_IsActive(void* ptr){
	return static_cast<QWindow*>(ptr)->isActive();
}

int QWindow_IsAncestorOf(void* ptr, void* child, int mode){
	return static_cast<QWindow*>(ptr)->isAncestorOf(static_cast<QWindow*>(child), static_cast<QWindow::AncestorMode>(mode));
}

int QWindow_IsExposed(void* ptr){
	return static_cast<QWindow*>(ptr)->isExposed();
}

int QWindow_IsModal(void* ptr){
	return static_cast<QWindow*>(ptr)->isModal();
}

int QWindow_IsTopLevel(void* ptr){
	return static_cast<QWindow*>(ptr)->isTopLevel();
}

void QWindow_Lower(void* ptr){
	QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "lower");
}

void* QWindow_Mask(void* ptr){
	return new QRegion(static_cast<QWindow*>(ptr)->mask());
}

int QWindow_MaximumHeight(void* ptr){
	return static_cast<QWindow*>(ptr)->maximumHeight();
}

void QWindow_ConnectMaximumHeightChanged(void* ptr){
	QObject::connect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(int)>(&QWindow::maximumHeightChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(int)>(&MyQWindow::Signal_MaximumHeightChanged));;
}

void QWindow_DisconnectMaximumHeightChanged(void* ptr){
	QObject::disconnect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(int)>(&QWindow::maximumHeightChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(int)>(&MyQWindow::Signal_MaximumHeightChanged));;
}

int QWindow_MaximumWidth(void* ptr){
	return static_cast<QWindow*>(ptr)->maximumWidth();
}

void QWindow_ConnectMaximumWidthChanged(void* ptr){
	QObject::connect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(int)>(&QWindow::maximumWidthChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(int)>(&MyQWindow::Signal_MaximumWidthChanged));;
}

void QWindow_DisconnectMaximumWidthChanged(void* ptr){
	QObject::disconnect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(int)>(&QWindow::maximumWidthChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(int)>(&MyQWindow::Signal_MaximumWidthChanged));;
}

int QWindow_MinimumHeight(void* ptr){
	return static_cast<QWindow*>(ptr)->minimumHeight();
}

void QWindow_ConnectMinimumHeightChanged(void* ptr){
	QObject::connect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(int)>(&QWindow::minimumHeightChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(int)>(&MyQWindow::Signal_MinimumHeightChanged));;
}

void QWindow_DisconnectMinimumHeightChanged(void* ptr){
	QObject::disconnect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(int)>(&QWindow::minimumHeightChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(int)>(&MyQWindow::Signal_MinimumHeightChanged));;
}

int QWindow_MinimumWidth(void* ptr){
	return static_cast<QWindow*>(ptr)->minimumWidth();
}

void QWindow_ConnectMinimumWidthChanged(void* ptr){
	QObject::connect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(int)>(&QWindow::minimumWidthChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(int)>(&MyQWindow::Signal_MinimumWidthChanged));;
}

void QWindow_DisconnectMinimumWidthChanged(void* ptr){
	QObject::disconnect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(int)>(&QWindow::minimumWidthChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(int)>(&MyQWindow::Signal_MinimumWidthChanged));;
}

void QWindow_ConnectModalityChanged(void* ptr){
	QObject::connect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(Qt::WindowModality)>(&QWindow::modalityChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(Qt::WindowModality)>(&MyQWindow::Signal_ModalityChanged));;
}

void QWindow_DisconnectModalityChanged(void* ptr){
	QObject::disconnect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(Qt::WindowModality)>(&QWindow::modalityChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(Qt::WindowModality)>(&MyQWindow::Signal_ModalityChanged));;
}

void* QWindow_Parent(void* ptr){
	return static_cast<QWindow*>(ptr)->parent();
}

void QWindow_Raise(void* ptr){
	QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "raise");
}

void QWindow_RequestActivate(void* ptr){
	QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "requestActivate");
}

void QWindow_RequestUpdate(void* ptr){
	QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "requestUpdate");
}

void QWindow_Resize(void* ptr, void* newSize){
	static_cast<QWindow*>(ptr)->resize(*static_cast<QSize*>(newSize));
}

void QWindow_Resize2(void* ptr, int w, int h){
	static_cast<QWindow*>(ptr)->resize(w, h);
}

void* QWindow_Screen(void* ptr){
	return static_cast<QWindow*>(ptr)->screen();
}

void QWindow_ConnectScreenChanged(void* ptr){
	QObject::connect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(QScreen *)>(&QWindow::screenChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(QScreen *)>(&MyQWindow::Signal_ScreenChanged));;
}

void QWindow_DisconnectScreenChanged(void* ptr){
	QObject::disconnect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(QScreen *)>(&QWindow::screenChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(QScreen *)>(&MyQWindow::Signal_ScreenChanged));;
}

void QWindow_SetBaseSize(void* ptr, void* size){
	static_cast<QWindow*>(ptr)->setBaseSize(*static_cast<QSize*>(size));
}

void QWindow_SetCursor(void* ptr, void* cursor){
	static_cast<QWindow*>(ptr)->setCursor(*static_cast<QCursor*>(cursor));
}

void QWindow_SetFilePath(void* ptr, char* filePath){
	static_cast<QWindow*>(ptr)->setFilePath(QString(filePath));
}

void QWindow_SetFormat(void* ptr, void* format){
	static_cast<QWindow*>(ptr)->setFormat(*static_cast<QSurfaceFormat*>(format));
}

void QWindow_SetFramePosition(void* ptr, void* point){
	static_cast<QWindow*>(ptr)->setFramePosition(*static_cast<QPoint*>(point));
}

void QWindow_SetGeometry2(void* ptr, void* rect){
	static_cast<QWindow*>(ptr)->setGeometry(*static_cast<QRect*>(rect));
}

void QWindow_SetGeometry(void* ptr, int posx, int posy, int w, int h){
	static_cast<QWindow*>(ptr)->setGeometry(posx, posy, w, h);
}

void QWindow_SetIcon(void* ptr, void* icon){
	static_cast<QWindow*>(ptr)->setIcon(*static_cast<QIcon*>(icon));
}

int QWindow_SetKeyboardGrabEnabled(void* ptr, int grab){
	return static_cast<QWindow*>(ptr)->setKeyboardGrabEnabled(grab != 0);
}

void QWindow_SetMask(void* ptr, void* region){
	static_cast<QWindow*>(ptr)->setMask(*static_cast<QRegion*>(region));
}

void QWindow_SetMaximumSize(void* ptr, void* size){
	static_cast<QWindow*>(ptr)->setMaximumSize(*static_cast<QSize*>(size));
}

void QWindow_SetMinimumSize(void* ptr, void* size){
	static_cast<QWindow*>(ptr)->setMinimumSize(*static_cast<QSize*>(size));
}

int QWindow_SetMouseGrabEnabled(void* ptr, int grab){
	return static_cast<QWindow*>(ptr)->setMouseGrabEnabled(grab != 0);
}

void QWindow_SetParent(void* ptr, void* parent){
	static_cast<QWindow*>(ptr)->setParent(static_cast<QWindow*>(parent));
}

void QWindow_SetPosition(void* ptr, void* pt){
	static_cast<QWindow*>(ptr)->setPosition(*static_cast<QPoint*>(pt));
}

void QWindow_SetPosition2(void* ptr, int posx, int posy){
	static_cast<QWindow*>(ptr)->setPosition(posx, posy);
}

void QWindow_SetScreen(void* ptr, void* newScreen){
	static_cast<QWindow*>(ptr)->setScreen(static_cast<QScreen*>(newScreen));
}

void QWindow_SetSizeIncrement(void* ptr, void* size){
	static_cast<QWindow*>(ptr)->setSizeIncrement(*static_cast<QSize*>(size));
}

void QWindow_SetSurfaceType(void* ptr, int surfaceType){
	static_cast<QWindow*>(ptr)->setSurfaceType(static_cast<QSurface::SurfaceType>(surfaceType));
}

void QWindow_SetTransientParent(void* ptr, void* parent){
	static_cast<QWindow*>(ptr)->setTransientParent(static_cast<QWindow*>(parent));
}

void QWindow_SetWindowState(void* ptr, int state){
	static_cast<QWindow*>(ptr)->setWindowState(static_cast<Qt::WindowState>(state));
}

void QWindow_Show(void* ptr){
	QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "show");
}

void QWindow_ShowFullScreen(void* ptr){
	QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "showFullScreen");
}

void QWindow_ShowMaximized(void* ptr){
	QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "showMaximized");
}

void QWindow_ShowMinimized(void* ptr){
	QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "showMinimized");
}

void QWindow_ShowNormal(void* ptr){
	QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "showNormal");
}

int QWindow_SurfaceType(void* ptr){
	return static_cast<QWindow*>(ptr)->surfaceType();
}

void* QWindow_TransientParent(void* ptr){
	return static_cast<QWindow*>(ptr)->transientParent();
}

int QWindow_Type(void* ptr){
	return static_cast<QWindow*>(ptr)->type();
}

void QWindow_UnsetCursor(void* ptr){
	static_cast<QWindow*>(ptr)->unsetCursor();
}

void QWindow_ConnectVisibilityChanged(void* ptr){
	QObject::connect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(QWindow::Visibility)>(&QWindow::visibilityChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(QWindow::Visibility)>(&MyQWindow::Signal_VisibilityChanged));;
}

void QWindow_DisconnectVisibilityChanged(void* ptr){
	QObject::disconnect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(QWindow::Visibility)>(&QWindow::visibilityChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(QWindow::Visibility)>(&MyQWindow::Signal_VisibilityChanged));;
}

void QWindow_ConnectVisibleChanged(void* ptr){
	QObject::connect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(bool)>(&QWindow::visibleChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(bool)>(&MyQWindow::Signal_VisibleChanged));;
}

void QWindow_DisconnectVisibleChanged(void* ptr){
	QObject::disconnect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(bool)>(&QWindow::visibleChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(bool)>(&MyQWindow::Signal_VisibleChanged));;
}

int QWindow_Width(void* ptr){
	return static_cast<QWindow*>(ptr)->width();
}

void QWindow_ConnectWidthChanged(void* ptr){
	QObject::connect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(int)>(&QWindow::widthChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(int)>(&MyQWindow::Signal_WidthChanged));;
}

void QWindow_DisconnectWidthChanged(void* ptr){
	QObject::disconnect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(int)>(&QWindow::widthChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(int)>(&MyQWindow::Signal_WidthChanged));;
}

int QWindow_WindowState(void* ptr){
	return static_cast<QWindow*>(ptr)->windowState();
}

void QWindow_ConnectWindowStateChanged(void* ptr){
	QObject::connect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(Qt::WindowState)>(&QWindow::windowStateChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(Qt::WindowState)>(&MyQWindow::Signal_WindowStateChanged));;
}

void QWindow_DisconnectWindowStateChanged(void* ptr){
	QObject::disconnect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(Qt::WindowState)>(&QWindow::windowStateChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(Qt::WindowState)>(&MyQWindow::Signal_WindowStateChanged));;
}

void QWindow_ConnectWindowTitleChanged(void* ptr){
	QObject::connect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(const QString &)>(&QWindow::windowTitleChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(const QString &)>(&MyQWindow::Signal_WindowTitleChanged));;
}

void QWindow_DisconnectWindowTitleChanged(void* ptr){
	QObject::disconnect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(const QString &)>(&QWindow::windowTitleChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(const QString &)>(&MyQWindow::Signal_WindowTitleChanged));;
}

int QWindow_X(void* ptr){
	return static_cast<QWindow*>(ptr)->x();
}

void QWindow_ConnectXChanged(void* ptr){
	QObject::connect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(int)>(&QWindow::xChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(int)>(&MyQWindow::Signal_XChanged));;
}

void QWindow_DisconnectXChanged(void* ptr){
	QObject::disconnect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(int)>(&QWindow::xChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(int)>(&MyQWindow::Signal_XChanged));;
}

int QWindow_Y(void* ptr){
	return static_cast<QWindow*>(ptr)->y();
}

void QWindow_ConnectYChanged(void* ptr){
	QObject::connect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(int)>(&QWindow::yChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(int)>(&MyQWindow::Signal_YChanged));;
}

void QWindow_DisconnectYChanged(void* ptr){
	QObject::disconnect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(int)>(&QWindow::yChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(int)>(&MyQWindow::Signal_YChanged));;
}

void QWindow_DestroyQWindow(void* ptr){
	static_cast<QWindow*>(ptr)->~QWindow();
}

