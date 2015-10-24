#include "qguiapplication.h"
#include <QCursor>
#include <QString>
#include <QVariant>
#include <QModelIndex>
#include <QWindow>
#include <QFont>
#include <QPalette>
#include <QUrl>
#include <QEvent>
#include <QScreen>
#include <QObject>
#include <QPoint>
#include <QIcon>
#include <QGuiApplication>
#include "_cgo_export.h"

class MyQGuiApplication: public QGuiApplication {
public:
void Signal_ApplicationStateChanged(Qt::ApplicationState state){callbackQGuiApplicationApplicationStateChanged(this->objectName().toUtf8().data(), state);};
void Signal_FocusObjectChanged(QObject * focusObject){callbackQGuiApplicationFocusObjectChanged(this->objectName().toUtf8().data(), focusObject);};
void Signal_FocusWindowChanged(QWindow * focusWindow){callbackQGuiApplicationFocusWindowChanged(this->objectName().toUtf8().data(), focusWindow);};
void Signal_FontDatabaseChanged(){callbackQGuiApplicationFontDatabaseChanged(this->objectName().toUtf8().data());};
void Signal_LastWindowClosed(){callbackQGuiApplicationLastWindowClosed(this->objectName().toUtf8().data());};
void Signal_LayoutDirectionChanged(Qt::LayoutDirection direction){callbackQGuiApplicationLayoutDirectionChanged(this->objectName().toUtf8().data(), direction);};
void Signal_ScreenAdded(QScreen * screen){callbackQGuiApplicationScreenAdded(this->objectName().toUtf8().data(), screen);};
void Signal_ScreenRemoved(QScreen * screen){callbackQGuiApplicationScreenRemoved(this->objectName().toUtf8().data(), screen);};
};

char* QGuiApplication_QGuiApplication_ApplicationDisplayName(){
	return QGuiApplication::applicationDisplayName().toUtf8().data();
}

int QGuiApplication_QGuiApplication_ApplicationState(){
	return QGuiApplication::applicationState();
}

int QGuiApplication_IsSavingSession(QtObjectPtr ptr){
	return static_cast<QGuiApplication*>(ptr)->isSavingSession();
}

int QGuiApplication_IsSessionRestored(QtObjectPtr ptr){
	return static_cast<QGuiApplication*>(ptr)->isSessionRestored();
}

int QGuiApplication_QGuiApplication_LayoutDirection(){
	return QGuiApplication::layoutDirection();
}

QtObjectPtr QGuiApplication_QGuiApplication_OverrideCursor(){
	return QGuiApplication::overrideCursor();
}

char* QGuiApplication_QGuiApplication_PlatformName(){
	return QGuiApplication::platformName().toUtf8().data();
}

int QGuiApplication_QGuiApplication_QueryKeyboardModifiers(){
	return QGuiApplication::queryKeyboardModifiers();
}

int QGuiApplication_QGuiApplication_QuitOnLastWindowClosed(){
	return QGuiApplication::quitOnLastWindowClosed();
}

void QGuiApplication_QGuiApplication_RestoreOverrideCursor(){
	QGuiApplication::restoreOverrideCursor();
}

char* QGuiApplication_SessionId(QtObjectPtr ptr){
	return static_cast<QGuiApplication*>(ptr)->sessionId().toUtf8().data();
}

char* QGuiApplication_SessionKey(QtObjectPtr ptr){
	return static_cast<QGuiApplication*>(ptr)->sessionKey().toUtf8().data();
}

void QGuiApplication_QGuiApplication_SetApplicationDisplayName(char* name){
	QGuiApplication::setApplicationDisplayName(QString(name));
}

void QGuiApplication_QGuiApplication_SetLayoutDirection(int direction){
	QGuiApplication::setLayoutDirection(static_cast<Qt::LayoutDirection>(direction));
}

void QGuiApplication_QGuiApplication_SetOverrideCursor(QtObjectPtr cursor){
	QGuiApplication::setOverrideCursor(*static_cast<QCursor*>(cursor));
}

void QGuiApplication_QGuiApplication_SetQuitOnLastWindowClosed(int quit){
	QGuiApplication::setQuitOnLastWindowClosed(quit != 0);
}

void QGuiApplication_QGuiApplication_SetWindowIcon(QtObjectPtr icon){
	QGuiApplication::setWindowIcon(*static_cast<QIcon*>(icon));
}

QtObjectPtr QGuiApplication_NewQGuiApplication(int argc, char* argv){
	return new QGuiApplication(argc, &argv);
}

void QGuiApplication_ConnectApplicationStateChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QGuiApplication*>(ptr), static_cast<void (QGuiApplication::*)(Qt::ApplicationState)>(&QGuiApplication::applicationStateChanged), static_cast<MyQGuiApplication*>(ptr), static_cast<void (MyQGuiApplication::*)(Qt::ApplicationState)>(&MyQGuiApplication::Signal_ApplicationStateChanged));;
}

void QGuiApplication_DisconnectApplicationStateChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QGuiApplication*>(ptr), static_cast<void (QGuiApplication::*)(Qt::ApplicationState)>(&QGuiApplication::applicationStateChanged), static_cast<MyQGuiApplication*>(ptr), static_cast<void (MyQGuiApplication::*)(Qt::ApplicationState)>(&MyQGuiApplication::Signal_ApplicationStateChanged));;
}

void QGuiApplication_QGuiApplication_ChangeOverrideCursor(QtObjectPtr cursor){
	QGuiApplication::changeOverrideCursor(*static_cast<QCursor*>(cursor));
}

QtObjectPtr QGuiApplication_QGuiApplication_Clipboard(){
	return QGuiApplication::clipboard();
}

int QGuiApplication_QGuiApplication_DesktopSettingsAware(){
	return QGuiApplication::desktopSettingsAware();
}

int QGuiApplication_QGuiApplication_Exec(){
	return QGuiApplication::exec();
}

QtObjectPtr QGuiApplication_QGuiApplication_FocusObject(){
	return QGuiApplication::focusObject();
}

void QGuiApplication_ConnectFocusObjectChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QGuiApplication*>(ptr), static_cast<void (QGuiApplication::*)(QObject *)>(&QGuiApplication::focusObjectChanged), static_cast<MyQGuiApplication*>(ptr), static_cast<void (MyQGuiApplication::*)(QObject *)>(&MyQGuiApplication::Signal_FocusObjectChanged));;
}

void QGuiApplication_DisconnectFocusObjectChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QGuiApplication*>(ptr), static_cast<void (QGuiApplication::*)(QObject *)>(&QGuiApplication::focusObjectChanged), static_cast<MyQGuiApplication*>(ptr), static_cast<void (MyQGuiApplication::*)(QObject *)>(&MyQGuiApplication::Signal_FocusObjectChanged));;
}

QtObjectPtr QGuiApplication_QGuiApplication_FocusWindow(){
	return QGuiApplication::focusWindow();
}

void QGuiApplication_ConnectFocusWindowChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QGuiApplication*>(ptr), static_cast<void (QGuiApplication::*)(QWindow *)>(&QGuiApplication::focusWindowChanged), static_cast<MyQGuiApplication*>(ptr), static_cast<void (MyQGuiApplication::*)(QWindow *)>(&MyQGuiApplication::Signal_FocusWindowChanged));;
}

void QGuiApplication_DisconnectFocusWindowChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QGuiApplication*>(ptr), static_cast<void (QGuiApplication::*)(QWindow *)>(&QGuiApplication::focusWindowChanged), static_cast<MyQGuiApplication*>(ptr), static_cast<void (MyQGuiApplication::*)(QWindow *)>(&MyQGuiApplication::Signal_FocusWindowChanged));;
}

void QGuiApplication_ConnectFontDatabaseChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QGuiApplication*>(ptr), static_cast<void (QGuiApplication::*)()>(&QGuiApplication::fontDatabaseChanged), static_cast<MyQGuiApplication*>(ptr), static_cast<void (MyQGuiApplication::*)()>(&MyQGuiApplication::Signal_FontDatabaseChanged));;
}

void QGuiApplication_DisconnectFontDatabaseChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QGuiApplication*>(ptr), static_cast<void (QGuiApplication::*)()>(&QGuiApplication::fontDatabaseChanged), static_cast<MyQGuiApplication*>(ptr), static_cast<void (MyQGuiApplication::*)()>(&MyQGuiApplication::Signal_FontDatabaseChanged));;
}

QtObjectPtr QGuiApplication_QGuiApplication_InputMethod(){
	return QGuiApplication::inputMethod();
}

int QGuiApplication_QGuiApplication_IsLeftToRight(){
	return QGuiApplication::isLeftToRight();
}

int QGuiApplication_QGuiApplication_IsRightToLeft(){
	return QGuiApplication::isRightToLeft();
}

int QGuiApplication_QGuiApplication_KeyboardModifiers(){
	return QGuiApplication::keyboardModifiers();
}

void QGuiApplication_ConnectLastWindowClosed(QtObjectPtr ptr){
	QObject::connect(static_cast<QGuiApplication*>(ptr), static_cast<void (QGuiApplication::*)()>(&QGuiApplication::lastWindowClosed), static_cast<MyQGuiApplication*>(ptr), static_cast<void (MyQGuiApplication::*)()>(&MyQGuiApplication::Signal_LastWindowClosed));;
}

void QGuiApplication_DisconnectLastWindowClosed(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QGuiApplication*>(ptr), static_cast<void (QGuiApplication::*)()>(&QGuiApplication::lastWindowClosed), static_cast<MyQGuiApplication*>(ptr), static_cast<void (MyQGuiApplication::*)()>(&MyQGuiApplication::Signal_LastWindowClosed));;
}

void QGuiApplication_ConnectLayoutDirectionChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QGuiApplication*>(ptr), static_cast<void (QGuiApplication::*)(Qt::LayoutDirection)>(&QGuiApplication::layoutDirectionChanged), static_cast<MyQGuiApplication*>(ptr), static_cast<void (MyQGuiApplication::*)(Qt::LayoutDirection)>(&MyQGuiApplication::Signal_LayoutDirectionChanged));;
}

void QGuiApplication_DisconnectLayoutDirectionChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QGuiApplication*>(ptr), static_cast<void (QGuiApplication::*)(Qt::LayoutDirection)>(&QGuiApplication::layoutDirectionChanged), static_cast<MyQGuiApplication*>(ptr), static_cast<void (MyQGuiApplication::*)(Qt::LayoutDirection)>(&MyQGuiApplication::Signal_LayoutDirectionChanged));;
}

QtObjectPtr QGuiApplication_QGuiApplication_ModalWindow(){
	return QGuiApplication::modalWindow();
}

int QGuiApplication_QGuiApplication_MouseButtons(){
	return QGuiApplication::mouseButtons();
}

int QGuiApplication_Notify(QtObjectPtr ptr, QtObjectPtr object, QtObjectPtr event){
	return static_cast<QGuiApplication*>(ptr)->notify(static_cast<QObject*>(object), static_cast<QEvent*>(event));
}

QtObjectPtr QGuiApplication_QGuiApplication_PrimaryScreen(){
	return QGuiApplication::primaryScreen();
}

void QGuiApplication_ConnectScreenAdded(QtObjectPtr ptr){
	QObject::connect(static_cast<QGuiApplication*>(ptr), static_cast<void (QGuiApplication::*)(QScreen *)>(&QGuiApplication::screenAdded), static_cast<MyQGuiApplication*>(ptr), static_cast<void (MyQGuiApplication::*)(QScreen *)>(&MyQGuiApplication::Signal_ScreenAdded));;
}

void QGuiApplication_DisconnectScreenAdded(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QGuiApplication*>(ptr), static_cast<void (QGuiApplication::*)(QScreen *)>(&QGuiApplication::screenAdded), static_cast<MyQGuiApplication*>(ptr), static_cast<void (MyQGuiApplication::*)(QScreen *)>(&MyQGuiApplication::Signal_ScreenAdded));;
}

void QGuiApplication_ConnectScreenRemoved(QtObjectPtr ptr){
	QObject::connect(static_cast<QGuiApplication*>(ptr), static_cast<void (QGuiApplication::*)(QScreen *)>(&QGuiApplication::screenRemoved), static_cast<MyQGuiApplication*>(ptr), static_cast<void (MyQGuiApplication::*)(QScreen *)>(&MyQGuiApplication::Signal_ScreenRemoved));;
}

void QGuiApplication_DisconnectScreenRemoved(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QGuiApplication*>(ptr), static_cast<void (QGuiApplication::*)(QScreen *)>(&QGuiApplication::screenRemoved), static_cast<MyQGuiApplication*>(ptr), static_cast<void (MyQGuiApplication::*)(QScreen *)>(&MyQGuiApplication::Signal_ScreenRemoved));;
}

void QGuiApplication_QGuiApplication_SetDesktopSettingsAware(int on){
	QGuiApplication::setDesktopSettingsAware(on != 0);
}

void QGuiApplication_QGuiApplication_SetFont(QtObjectPtr font){
	QGuiApplication::setFont(*static_cast<QFont*>(font));
}

void QGuiApplication_QGuiApplication_SetPalette(QtObjectPtr pal){
	QGuiApplication::setPalette(*static_cast<QPalette*>(pal));
}

QtObjectPtr QGuiApplication_QGuiApplication_StyleHints(){
	return QGuiApplication::styleHints();
}

void QGuiApplication_QGuiApplication_Sync(){
	QGuiApplication::sync();
}

QtObjectPtr QGuiApplication_QGuiApplication_TopLevelAt(QtObjectPtr pos){
	return QGuiApplication::topLevelAt(*static_cast<QPoint*>(pos));
}

void QGuiApplication_DestroyQGuiApplication(QtObjectPtr ptr){
	static_cast<QGuiApplication*>(ptr)->~QGuiApplication();
}

