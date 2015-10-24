#include "qwindow.h"
#include <QUrl>
#include <QObject>
#include <QCursor>
#include <QRegion>
#include <QMetaObject>
#include <QPoint>
#include <QString>
#include <QIcon>
#include <QScreen>
#include <QSurface>
#include <QSurfaceFormat>
#include <QSize>
#include <QVariant>
#include <QModelIndex>
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

int QWindow_ContentOrientation(QtObjectPtr ptr){
	return static_cast<QWindow*>(ptr)->contentOrientation();
}

int QWindow_Flags(QtObjectPtr ptr){
	return static_cast<QWindow*>(ptr)->flags();
}

int QWindow_IsVisible(QtObjectPtr ptr){
	return static_cast<QWindow*>(ptr)->isVisible();
}

int QWindow_Modality(QtObjectPtr ptr){
	return static_cast<QWindow*>(ptr)->modality();
}

void QWindow_ReportContentOrientationChange(QtObjectPtr ptr, int orientation){
	static_cast<QWindow*>(ptr)->reportContentOrientationChange(static_cast<Qt::ScreenOrientation>(orientation));
}

void QWindow_SetFlags(QtObjectPtr ptr, int flags){
	static_cast<QWindow*>(ptr)->setFlags(static_cast<Qt::WindowType>(flags));
}

void QWindow_SetHeight(QtObjectPtr ptr, int arg){
	QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "setHeight", Q_ARG(int, arg));
}

void QWindow_SetMaximumHeight(QtObjectPtr ptr, int h){
	QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "setMaximumHeight", Q_ARG(int, h));
}

void QWindow_SetMaximumWidth(QtObjectPtr ptr, int w){
	QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "setMaximumWidth", Q_ARG(int, w));
}

void QWindow_SetMinimumHeight(QtObjectPtr ptr, int h){
	QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "setMinimumHeight", Q_ARG(int, h));
}

void QWindow_SetMinimumWidth(QtObjectPtr ptr, int w){
	QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "setMinimumWidth", Q_ARG(int, w));
}

void QWindow_SetModality(QtObjectPtr ptr, int modality){
	static_cast<QWindow*>(ptr)->setModality(static_cast<Qt::WindowModality>(modality));
}

void QWindow_SetTitle(QtObjectPtr ptr, char* v){
	QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "setTitle", Q_ARG(QString, QString(v)));
}

void QWindow_SetVisibility(QtObjectPtr ptr, int v){
	static_cast<QWindow*>(ptr)->setVisibility(static_cast<QWindow::Visibility>(v));
}

void QWindow_SetVisible(QtObjectPtr ptr, int visible){
	QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "setVisible", Q_ARG(bool, visible != 0));
}

void QWindow_SetWidth(QtObjectPtr ptr, int arg){
	QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "setWidth", Q_ARG(int, arg));
}

void QWindow_SetX(QtObjectPtr ptr, int arg){
	QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "setX", Q_ARG(int, arg));
}

void QWindow_SetY(QtObjectPtr ptr, int arg){
	QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "setY", Q_ARG(int, arg));
}

char* QWindow_Title(QtObjectPtr ptr){
	return static_cast<QWindow*>(ptr)->title().toUtf8().data();
}

int QWindow_Visibility(QtObjectPtr ptr){
	return static_cast<QWindow*>(ptr)->visibility();
}

QtObjectPtr QWindow_NewQWindow(QtObjectPtr targetScreen){
	return new QWindow(static_cast<QScreen*>(targetScreen));
}

QtObjectPtr QWindow_NewQWindow2(QtObjectPtr parent){
	return new QWindow(static_cast<QWindow*>(parent));
}

void QWindow_ConnectActiveChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)()>(&QWindow::activeChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)()>(&MyQWindow::Signal_ActiveChanged));;
}

void QWindow_DisconnectActiveChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)()>(&QWindow::activeChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)()>(&MyQWindow::Signal_ActiveChanged));;
}

void QWindow_Alert(QtObjectPtr ptr, int msec){
	QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "alert", Q_ARG(int, msec));
}

int QWindow_Close(QtObjectPtr ptr){
	return QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "close");
}

void QWindow_ConnectContentOrientationChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(Qt::ScreenOrientation)>(&QWindow::contentOrientationChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(Qt::ScreenOrientation)>(&MyQWindow::Signal_ContentOrientationChanged));;
}

void QWindow_DisconnectContentOrientationChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(Qt::ScreenOrientation)>(&QWindow::contentOrientationChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(Qt::ScreenOrientation)>(&MyQWindow::Signal_ContentOrientationChanged));;
}

void QWindow_Create(QtObjectPtr ptr){
	static_cast<QWindow*>(ptr)->create();
}

void QWindow_Destroy(QtObjectPtr ptr){
	static_cast<QWindow*>(ptr)->destroy();
}

char* QWindow_FilePath(QtObjectPtr ptr){
	return static_cast<QWindow*>(ptr)->filePath().toUtf8().data();
}

QtObjectPtr QWindow_FocusObject(QtObjectPtr ptr){
	return static_cast<QWindow*>(ptr)->focusObject();
}

void QWindow_ConnectFocusObjectChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(QObject *)>(&QWindow::focusObjectChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(QObject *)>(&MyQWindow::Signal_FocusObjectChanged));;
}

void QWindow_DisconnectFocusObjectChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(QObject *)>(&QWindow::focusObjectChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(QObject *)>(&MyQWindow::Signal_FocusObjectChanged));;
}

int QWindow_Height(QtObjectPtr ptr){
	return static_cast<QWindow*>(ptr)->height();
}

void QWindow_ConnectHeightChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(int)>(&QWindow::heightChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(int)>(&MyQWindow::Signal_HeightChanged));;
}

void QWindow_DisconnectHeightChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(int)>(&QWindow::heightChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(int)>(&MyQWindow::Signal_HeightChanged));;
}

void QWindow_Hide(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "hide");
}

int QWindow_IsActive(QtObjectPtr ptr){
	return static_cast<QWindow*>(ptr)->isActive();
}

int QWindow_IsAncestorOf(QtObjectPtr ptr, QtObjectPtr child, int mode){
	return static_cast<QWindow*>(ptr)->isAncestorOf(static_cast<QWindow*>(child), static_cast<QWindow::AncestorMode>(mode));
}

int QWindow_IsExposed(QtObjectPtr ptr){
	return static_cast<QWindow*>(ptr)->isExposed();
}

int QWindow_IsModal(QtObjectPtr ptr){
	return static_cast<QWindow*>(ptr)->isModal();
}

int QWindow_IsTopLevel(QtObjectPtr ptr){
	return static_cast<QWindow*>(ptr)->isTopLevel();
}

void QWindow_Lower(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "lower");
}

int QWindow_MaximumHeight(QtObjectPtr ptr){
	return static_cast<QWindow*>(ptr)->maximumHeight();
}

void QWindow_ConnectMaximumHeightChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(int)>(&QWindow::maximumHeightChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(int)>(&MyQWindow::Signal_MaximumHeightChanged));;
}

void QWindow_DisconnectMaximumHeightChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(int)>(&QWindow::maximumHeightChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(int)>(&MyQWindow::Signal_MaximumHeightChanged));;
}

int QWindow_MaximumWidth(QtObjectPtr ptr){
	return static_cast<QWindow*>(ptr)->maximumWidth();
}

void QWindow_ConnectMaximumWidthChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(int)>(&QWindow::maximumWidthChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(int)>(&MyQWindow::Signal_MaximumWidthChanged));;
}

void QWindow_DisconnectMaximumWidthChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(int)>(&QWindow::maximumWidthChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(int)>(&MyQWindow::Signal_MaximumWidthChanged));;
}

int QWindow_MinimumHeight(QtObjectPtr ptr){
	return static_cast<QWindow*>(ptr)->minimumHeight();
}

void QWindow_ConnectMinimumHeightChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(int)>(&QWindow::minimumHeightChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(int)>(&MyQWindow::Signal_MinimumHeightChanged));;
}

void QWindow_DisconnectMinimumHeightChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(int)>(&QWindow::minimumHeightChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(int)>(&MyQWindow::Signal_MinimumHeightChanged));;
}

int QWindow_MinimumWidth(QtObjectPtr ptr){
	return static_cast<QWindow*>(ptr)->minimumWidth();
}

void QWindow_ConnectMinimumWidthChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(int)>(&QWindow::minimumWidthChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(int)>(&MyQWindow::Signal_MinimumWidthChanged));;
}

void QWindow_DisconnectMinimumWidthChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(int)>(&QWindow::minimumWidthChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(int)>(&MyQWindow::Signal_MinimumWidthChanged));;
}

void QWindow_ConnectModalityChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(Qt::WindowModality)>(&QWindow::modalityChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(Qt::WindowModality)>(&MyQWindow::Signal_ModalityChanged));;
}

void QWindow_DisconnectModalityChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(Qt::WindowModality)>(&QWindow::modalityChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(Qt::WindowModality)>(&MyQWindow::Signal_ModalityChanged));;
}

QtObjectPtr QWindow_Parent(QtObjectPtr ptr){
	return static_cast<QWindow*>(ptr)->parent();
}

void QWindow_Raise(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "raise");
}

void QWindow_RequestActivate(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "requestActivate");
}

void QWindow_RequestUpdate(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "requestUpdate");
}

void QWindow_Resize(QtObjectPtr ptr, QtObjectPtr newSize){
	static_cast<QWindow*>(ptr)->resize(*static_cast<QSize*>(newSize));
}

void QWindow_Resize2(QtObjectPtr ptr, int w, int h){
	static_cast<QWindow*>(ptr)->resize(w, h);
}

QtObjectPtr QWindow_Screen(QtObjectPtr ptr){
	return static_cast<QWindow*>(ptr)->screen();
}

void QWindow_ConnectScreenChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(QScreen *)>(&QWindow::screenChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(QScreen *)>(&MyQWindow::Signal_ScreenChanged));;
}

void QWindow_DisconnectScreenChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(QScreen *)>(&QWindow::screenChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(QScreen *)>(&MyQWindow::Signal_ScreenChanged));;
}

void QWindow_SetBaseSize(QtObjectPtr ptr, QtObjectPtr size){
	static_cast<QWindow*>(ptr)->setBaseSize(*static_cast<QSize*>(size));
}

void QWindow_SetCursor(QtObjectPtr ptr, QtObjectPtr cursor){
	static_cast<QWindow*>(ptr)->setCursor(*static_cast<QCursor*>(cursor));
}

void QWindow_SetFilePath(QtObjectPtr ptr, char* filePath){
	static_cast<QWindow*>(ptr)->setFilePath(QString(filePath));
}

void QWindow_SetFormat(QtObjectPtr ptr, QtObjectPtr format){
	static_cast<QWindow*>(ptr)->setFormat(*static_cast<QSurfaceFormat*>(format));
}

void QWindow_SetFramePosition(QtObjectPtr ptr, QtObjectPtr point){
	static_cast<QWindow*>(ptr)->setFramePosition(*static_cast<QPoint*>(point));
}

void QWindow_SetGeometry2(QtObjectPtr ptr, QtObjectPtr rect){
	static_cast<QWindow*>(ptr)->setGeometry(*static_cast<QRect*>(rect));
}

void QWindow_SetGeometry(QtObjectPtr ptr, int posx, int posy, int w, int h){
	static_cast<QWindow*>(ptr)->setGeometry(posx, posy, w, h);
}

void QWindow_SetIcon(QtObjectPtr ptr, QtObjectPtr icon){
	static_cast<QWindow*>(ptr)->setIcon(*static_cast<QIcon*>(icon));
}

int QWindow_SetKeyboardGrabEnabled(QtObjectPtr ptr, int grab){
	return static_cast<QWindow*>(ptr)->setKeyboardGrabEnabled(grab != 0);
}

void QWindow_SetMask(QtObjectPtr ptr, QtObjectPtr region){
	static_cast<QWindow*>(ptr)->setMask(*static_cast<QRegion*>(region));
}

void QWindow_SetMaximumSize(QtObjectPtr ptr, QtObjectPtr size){
	static_cast<QWindow*>(ptr)->setMaximumSize(*static_cast<QSize*>(size));
}

void QWindow_SetMinimumSize(QtObjectPtr ptr, QtObjectPtr size){
	static_cast<QWindow*>(ptr)->setMinimumSize(*static_cast<QSize*>(size));
}

int QWindow_SetMouseGrabEnabled(QtObjectPtr ptr, int grab){
	return static_cast<QWindow*>(ptr)->setMouseGrabEnabled(grab != 0);
}

void QWindow_SetParent(QtObjectPtr ptr, QtObjectPtr parent){
	static_cast<QWindow*>(ptr)->setParent(static_cast<QWindow*>(parent));
}

void QWindow_SetPosition(QtObjectPtr ptr, QtObjectPtr pt){
	static_cast<QWindow*>(ptr)->setPosition(*static_cast<QPoint*>(pt));
}

void QWindow_SetPosition2(QtObjectPtr ptr, int posx, int posy){
	static_cast<QWindow*>(ptr)->setPosition(posx, posy);
}

void QWindow_SetScreen(QtObjectPtr ptr, QtObjectPtr newScreen){
	static_cast<QWindow*>(ptr)->setScreen(static_cast<QScreen*>(newScreen));
}

void QWindow_SetSizeIncrement(QtObjectPtr ptr, QtObjectPtr size){
	static_cast<QWindow*>(ptr)->setSizeIncrement(*static_cast<QSize*>(size));
}

void QWindow_SetSurfaceType(QtObjectPtr ptr, int surfaceType){
	static_cast<QWindow*>(ptr)->setSurfaceType(static_cast<QSurface::SurfaceType>(surfaceType));
}

void QWindow_SetTransientParent(QtObjectPtr ptr, QtObjectPtr parent){
	static_cast<QWindow*>(ptr)->setTransientParent(static_cast<QWindow*>(parent));
}

void QWindow_SetWindowState(QtObjectPtr ptr, int state){
	static_cast<QWindow*>(ptr)->setWindowState(static_cast<Qt::WindowState>(state));
}

void QWindow_Show(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "show");
}

void QWindow_ShowFullScreen(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "showFullScreen");
}

void QWindow_ShowMaximized(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "showMaximized");
}

void QWindow_ShowMinimized(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "showMinimized");
}

void QWindow_ShowNormal(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "showNormal");
}

int QWindow_SurfaceType(QtObjectPtr ptr){
	return static_cast<QWindow*>(ptr)->surfaceType();
}

QtObjectPtr QWindow_TransientParent(QtObjectPtr ptr){
	return static_cast<QWindow*>(ptr)->transientParent();
}

int QWindow_Type(QtObjectPtr ptr){
	return static_cast<QWindow*>(ptr)->type();
}

void QWindow_UnsetCursor(QtObjectPtr ptr){
	static_cast<QWindow*>(ptr)->unsetCursor();
}

void QWindow_ConnectVisibilityChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(QWindow::Visibility)>(&QWindow::visibilityChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(QWindow::Visibility)>(&MyQWindow::Signal_VisibilityChanged));;
}

void QWindow_DisconnectVisibilityChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(QWindow::Visibility)>(&QWindow::visibilityChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(QWindow::Visibility)>(&MyQWindow::Signal_VisibilityChanged));;
}

void QWindow_ConnectVisibleChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(bool)>(&QWindow::visibleChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(bool)>(&MyQWindow::Signal_VisibleChanged));;
}

void QWindow_DisconnectVisibleChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(bool)>(&QWindow::visibleChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(bool)>(&MyQWindow::Signal_VisibleChanged));;
}

int QWindow_Width(QtObjectPtr ptr){
	return static_cast<QWindow*>(ptr)->width();
}

void QWindow_ConnectWidthChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(int)>(&QWindow::widthChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(int)>(&MyQWindow::Signal_WidthChanged));;
}

void QWindow_DisconnectWidthChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(int)>(&QWindow::widthChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(int)>(&MyQWindow::Signal_WidthChanged));;
}

int QWindow_WindowState(QtObjectPtr ptr){
	return static_cast<QWindow*>(ptr)->windowState();
}

void QWindow_ConnectWindowStateChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(Qt::WindowState)>(&QWindow::windowStateChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(Qt::WindowState)>(&MyQWindow::Signal_WindowStateChanged));;
}

void QWindow_DisconnectWindowStateChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(Qt::WindowState)>(&QWindow::windowStateChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(Qt::WindowState)>(&MyQWindow::Signal_WindowStateChanged));;
}

void QWindow_ConnectWindowTitleChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(const QString &)>(&QWindow::windowTitleChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(const QString &)>(&MyQWindow::Signal_WindowTitleChanged));;
}

void QWindow_DisconnectWindowTitleChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(const QString &)>(&QWindow::windowTitleChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(const QString &)>(&MyQWindow::Signal_WindowTitleChanged));;
}

int QWindow_X(QtObjectPtr ptr){
	return static_cast<QWindow*>(ptr)->x();
}

void QWindow_ConnectXChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(int)>(&QWindow::xChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(int)>(&MyQWindow::Signal_XChanged));;
}

void QWindow_DisconnectXChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(int)>(&QWindow::xChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(int)>(&MyQWindow::Signal_XChanged));;
}

int QWindow_Y(QtObjectPtr ptr){
	return static_cast<QWindow*>(ptr)->y();
}

void QWindow_ConnectYChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(int)>(&QWindow::yChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(int)>(&MyQWindow::Signal_YChanged));;
}

void QWindow_DisconnectYChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(int)>(&QWindow::yChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(int)>(&MyQWindow::Signal_YChanged));;
}

void QWindow_DestroyQWindow(QtObjectPtr ptr){
	static_cast<QWindow*>(ptr)->~QWindow();
}

