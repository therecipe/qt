#include "qguiapplication.h"
#include <QObject>
#include <QByteArray>
#include <QCursor>
#include <QVariant>
#include <QUrl>
#include <QWindow>
#include <QPoint>
#include <QString>
#include <QIcon>
#include <QScreen>
#include <QEvent>
#include <QPalette>
#include <QList>
#include <QModelIndex>
#include <QFont>
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

int QGuiApplication_IsSavingSession(void* ptr){
	return static_cast<QGuiApplication*>(ptr)->isSavingSession();
}

int QGuiApplication_IsSessionRestored(void* ptr){
	return static_cast<QGuiApplication*>(ptr)->isSessionRestored();
}

int QGuiApplication_QGuiApplication_LayoutDirection(){
	return QGuiApplication::layoutDirection();
}

void* QGuiApplication_QGuiApplication_OverrideCursor(){
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

char* QGuiApplication_SessionId(void* ptr){
	return static_cast<QGuiApplication*>(ptr)->sessionId().toUtf8().data();
}

char* QGuiApplication_SessionKey(void* ptr){
	return static_cast<QGuiApplication*>(ptr)->sessionKey().toUtf8().data();
}

void QGuiApplication_QGuiApplication_SetApplicationDisplayName(char* name){
	QGuiApplication::setApplicationDisplayName(QString(name));
}

void QGuiApplication_QGuiApplication_SetLayoutDirection(int direction){
	QGuiApplication::setLayoutDirection(static_cast<Qt::LayoutDirection>(direction));
}

void QGuiApplication_QGuiApplication_SetOverrideCursor(void* cursor){
	QGuiApplication::setOverrideCursor(*static_cast<QCursor*>(cursor));
}

void QGuiApplication_QGuiApplication_SetQuitOnLastWindowClosed(int quit){
	QGuiApplication::setQuitOnLastWindowClosed(quit != 0);
}

void QGuiApplication_QGuiApplication_SetWindowIcon(void* icon){
	QGuiApplication::setWindowIcon(*static_cast<QIcon*>(icon));
}

void* QGuiApplication_NewQGuiApplication(int argc, char* argv){
	QList<QByteArray> aList = QByteArray(argv).split('|');
	char *argvs[argc];
	static int argcs = argc;
	for (int i = 0; i < argc; i++)
		argvs[i] = aList[i].data();

	return new QGuiApplication(argcs, argvs);
}

void QGuiApplication_ConnectApplicationStateChanged(void* ptr){
	QObject::connect(static_cast<QGuiApplication*>(ptr), static_cast<void (QGuiApplication::*)(Qt::ApplicationState)>(&QGuiApplication::applicationStateChanged), static_cast<MyQGuiApplication*>(ptr), static_cast<void (MyQGuiApplication::*)(Qt::ApplicationState)>(&MyQGuiApplication::Signal_ApplicationStateChanged));;
}

void QGuiApplication_DisconnectApplicationStateChanged(void* ptr){
	QObject::disconnect(static_cast<QGuiApplication*>(ptr), static_cast<void (QGuiApplication::*)(Qt::ApplicationState)>(&QGuiApplication::applicationStateChanged), static_cast<MyQGuiApplication*>(ptr), static_cast<void (MyQGuiApplication::*)(Qt::ApplicationState)>(&MyQGuiApplication::Signal_ApplicationStateChanged));;
}

void QGuiApplication_QGuiApplication_ChangeOverrideCursor(void* cursor){
	QGuiApplication::changeOverrideCursor(*static_cast<QCursor*>(cursor));
}

void* QGuiApplication_QGuiApplication_Clipboard(){
	return QGuiApplication::clipboard();
}

int QGuiApplication_QGuiApplication_DesktopSettingsAware(){
	return QGuiApplication::desktopSettingsAware();
}

double QGuiApplication_DevicePixelRatio(void* ptr){
	return static_cast<double>(static_cast<QGuiApplication*>(ptr)->devicePixelRatio());
}

int QGuiApplication_QGuiApplication_Exec(){
	return QGuiApplication::exec();
}

void* QGuiApplication_QGuiApplication_FocusObject(){
	return QGuiApplication::focusObject();
}

void QGuiApplication_ConnectFocusObjectChanged(void* ptr){
	QObject::connect(static_cast<QGuiApplication*>(ptr), static_cast<void (QGuiApplication::*)(QObject *)>(&QGuiApplication::focusObjectChanged), static_cast<MyQGuiApplication*>(ptr), static_cast<void (MyQGuiApplication::*)(QObject *)>(&MyQGuiApplication::Signal_FocusObjectChanged));;
}

void QGuiApplication_DisconnectFocusObjectChanged(void* ptr){
	QObject::disconnect(static_cast<QGuiApplication*>(ptr), static_cast<void (QGuiApplication::*)(QObject *)>(&QGuiApplication::focusObjectChanged), static_cast<MyQGuiApplication*>(ptr), static_cast<void (MyQGuiApplication::*)(QObject *)>(&MyQGuiApplication::Signal_FocusObjectChanged));;
}

void* QGuiApplication_QGuiApplication_FocusWindow(){
	return QGuiApplication::focusWindow();
}

void QGuiApplication_ConnectFocusWindowChanged(void* ptr){
	QObject::connect(static_cast<QGuiApplication*>(ptr), static_cast<void (QGuiApplication::*)(QWindow *)>(&QGuiApplication::focusWindowChanged), static_cast<MyQGuiApplication*>(ptr), static_cast<void (MyQGuiApplication::*)(QWindow *)>(&MyQGuiApplication::Signal_FocusWindowChanged));;
}

void QGuiApplication_DisconnectFocusWindowChanged(void* ptr){
	QObject::disconnect(static_cast<QGuiApplication*>(ptr), static_cast<void (QGuiApplication::*)(QWindow *)>(&QGuiApplication::focusWindowChanged), static_cast<MyQGuiApplication*>(ptr), static_cast<void (MyQGuiApplication::*)(QWindow *)>(&MyQGuiApplication::Signal_FocusWindowChanged));;
}

void QGuiApplication_ConnectFontDatabaseChanged(void* ptr){
	QObject::connect(static_cast<QGuiApplication*>(ptr), static_cast<void (QGuiApplication::*)()>(&QGuiApplication::fontDatabaseChanged), static_cast<MyQGuiApplication*>(ptr), static_cast<void (MyQGuiApplication::*)()>(&MyQGuiApplication::Signal_FontDatabaseChanged));;
}

void QGuiApplication_DisconnectFontDatabaseChanged(void* ptr){
	QObject::disconnect(static_cast<QGuiApplication*>(ptr), static_cast<void (QGuiApplication::*)()>(&QGuiApplication::fontDatabaseChanged), static_cast<MyQGuiApplication*>(ptr), static_cast<void (MyQGuiApplication::*)()>(&MyQGuiApplication::Signal_FontDatabaseChanged));;
}

void* QGuiApplication_QGuiApplication_InputMethod(){
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

void QGuiApplication_ConnectLastWindowClosed(void* ptr){
	QObject::connect(static_cast<QGuiApplication*>(ptr), static_cast<void (QGuiApplication::*)()>(&QGuiApplication::lastWindowClosed), static_cast<MyQGuiApplication*>(ptr), static_cast<void (MyQGuiApplication::*)()>(&MyQGuiApplication::Signal_LastWindowClosed));;
}

void QGuiApplication_DisconnectLastWindowClosed(void* ptr){
	QObject::disconnect(static_cast<QGuiApplication*>(ptr), static_cast<void (QGuiApplication::*)()>(&QGuiApplication::lastWindowClosed), static_cast<MyQGuiApplication*>(ptr), static_cast<void (MyQGuiApplication::*)()>(&MyQGuiApplication::Signal_LastWindowClosed));;
}

void QGuiApplication_ConnectLayoutDirectionChanged(void* ptr){
	QObject::connect(static_cast<QGuiApplication*>(ptr), static_cast<void (QGuiApplication::*)(Qt::LayoutDirection)>(&QGuiApplication::layoutDirectionChanged), static_cast<MyQGuiApplication*>(ptr), static_cast<void (MyQGuiApplication::*)(Qt::LayoutDirection)>(&MyQGuiApplication::Signal_LayoutDirectionChanged));;
}

void QGuiApplication_DisconnectLayoutDirectionChanged(void* ptr){
	QObject::disconnect(static_cast<QGuiApplication*>(ptr), static_cast<void (QGuiApplication::*)(Qt::LayoutDirection)>(&QGuiApplication::layoutDirectionChanged), static_cast<MyQGuiApplication*>(ptr), static_cast<void (MyQGuiApplication::*)(Qt::LayoutDirection)>(&MyQGuiApplication::Signal_LayoutDirectionChanged));;
}

void* QGuiApplication_QGuiApplication_ModalWindow(){
	return QGuiApplication::modalWindow();
}

int QGuiApplication_QGuiApplication_MouseButtons(){
	return QGuiApplication::mouseButtons();
}

int QGuiApplication_Notify(void* ptr, void* object, void* event){
	return static_cast<QGuiApplication*>(ptr)->notify(static_cast<QObject*>(object), static_cast<QEvent*>(event));
}

void* QGuiApplication_QGuiApplication_PrimaryScreen(){
	return QGuiApplication::primaryScreen();
}

void QGuiApplication_ConnectScreenAdded(void* ptr){
	QObject::connect(static_cast<QGuiApplication*>(ptr), static_cast<void (QGuiApplication::*)(QScreen *)>(&QGuiApplication::screenAdded), static_cast<MyQGuiApplication*>(ptr), static_cast<void (MyQGuiApplication::*)(QScreen *)>(&MyQGuiApplication::Signal_ScreenAdded));;
}

void QGuiApplication_DisconnectScreenAdded(void* ptr){
	QObject::disconnect(static_cast<QGuiApplication*>(ptr), static_cast<void (QGuiApplication::*)(QScreen *)>(&QGuiApplication::screenAdded), static_cast<MyQGuiApplication*>(ptr), static_cast<void (MyQGuiApplication::*)(QScreen *)>(&MyQGuiApplication::Signal_ScreenAdded));;
}

void QGuiApplication_ConnectScreenRemoved(void* ptr){
	QObject::connect(static_cast<QGuiApplication*>(ptr), static_cast<void (QGuiApplication::*)(QScreen *)>(&QGuiApplication::screenRemoved), static_cast<MyQGuiApplication*>(ptr), static_cast<void (MyQGuiApplication::*)(QScreen *)>(&MyQGuiApplication::Signal_ScreenRemoved));;
}

void QGuiApplication_DisconnectScreenRemoved(void* ptr){
	QObject::disconnect(static_cast<QGuiApplication*>(ptr), static_cast<void (QGuiApplication::*)(QScreen *)>(&QGuiApplication::screenRemoved), static_cast<MyQGuiApplication*>(ptr), static_cast<void (MyQGuiApplication::*)(QScreen *)>(&MyQGuiApplication::Signal_ScreenRemoved));;
}

void QGuiApplication_QGuiApplication_SetDesktopSettingsAware(int on){
	QGuiApplication::setDesktopSettingsAware(on != 0);
}

void QGuiApplication_QGuiApplication_SetFont(void* font){
	QGuiApplication::setFont(*static_cast<QFont*>(font));
}

void QGuiApplication_QGuiApplication_SetPalette(void* pal){
	QGuiApplication::setPalette(*static_cast<QPalette*>(pal));
}

void* QGuiApplication_QGuiApplication_StyleHints(){
	return QGuiApplication::styleHints();
}

void QGuiApplication_QGuiApplication_Sync(){
	QGuiApplication::sync();
}

void* QGuiApplication_QGuiApplication_TopLevelAt(void* pos){
	return QGuiApplication::topLevelAt(*static_cast<QPoint*>(pos));
}

void QGuiApplication_DestroyQGuiApplication(void* ptr){
	static_cast<QGuiApplication*>(ptr)->~QGuiApplication();
}

