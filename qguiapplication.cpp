#include "qguiapplication.h"
#include <QGuiApplication>
#include "cgoexport.h"

//MyQGuiApplication
class MyQGuiApplication: public QGuiApplication {
Q_OBJECT
public:
void Signal_ApplicationStateChanged() { callbackQGuiApplication(this, QString("applicationStateChanged").toUtf8().data()); };
void Signal_FocusObjectChanged() { callbackQGuiApplication(this, QString("focusObjectChanged").toUtf8().data()); };
void Signal_FontDatabaseChanged() { callbackQGuiApplication(this, QString("fontDatabaseChanged").toUtf8().data()); };
void Signal_LastWindowClosed() { callbackQGuiApplication(this, QString("lastWindowClosed").toUtf8().data()); };

signals:

};
#include "qguiapplication.moc"


//Public Functions
QtObjectPtr QGuiApplication_New_Int_String(int argc, char* argv)
{
	return new QGuiApplication(argc, ((char**)(argv)));
}

void QGuiApplication_Destroy(QtObjectPtr ptr)
{
	((QGuiApplication*)(ptr))->~QGuiApplication();
}

int QGuiApplication_IsSavingSession(QtObjectPtr ptr)
{
	return ((QGuiApplication*)(ptr))->isSavingSession();
}

int QGuiApplication_IsSessionRestored(QtObjectPtr ptr)
{
	return ((QGuiApplication*)(ptr))->isSessionRestored();
}

char* QGuiApplication_SessionId(QtObjectPtr ptr)
{
	return ((QGuiApplication*)(ptr))->sessionId().toUtf8().data();
}

char* QGuiApplication_SessionKey(QtObjectPtr ptr)
{
	return ((QGuiApplication*)(ptr))->sessionKey().toUtf8().data();
}

//Signals
void QGuiApplication_ConnectSignalApplicationStateChanged(QtObjectPtr ptr)
{
	QObject::connect(((QGuiApplication*)(ptr)), &QGuiApplication::applicationStateChanged, ((MyQGuiApplication*)(ptr)), &MyQGuiApplication::Signal_ApplicationStateChanged, Qt::QueuedConnection);
}

void QGuiApplication_DisconnectSignalApplicationStateChanged(QtObjectPtr ptr)
{
	QObject::disconnect(((QGuiApplication*)(ptr)), &QGuiApplication::applicationStateChanged, ((MyQGuiApplication*)(ptr)), &MyQGuiApplication::Signal_ApplicationStateChanged);
}

void QGuiApplication_ConnectSignalFocusObjectChanged(QtObjectPtr ptr)
{
	QObject::connect(((QGuiApplication*)(ptr)), &QGuiApplication::focusObjectChanged, ((MyQGuiApplication*)(ptr)), &MyQGuiApplication::Signal_FocusObjectChanged, Qt::QueuedConnection);
}

void QGuiApplication_DisconnectSignalFocusObjectChanged(QtObjectPtr ptr)
{
	QObject::disconnect(((QGuiApplication*)(ptr)), &QGuiApplication::focusObjectChanged, ((MyQGuiApplication*)(ptr)), &MyQGuiApplication::Signal_FocusObjectChanged);
}

void QGuiApplication_ConnectSignalFontDatabaseChanged(QtObjectPtr ptr)
{
	QObject::connect(((QGuiApplication*)(ptr)), &QGuiApplication::fontDatabaseChanged, ((MyQGuiApplication*)(ptr)), &MyQGuiApplication::Signal_FontDatabaseChanged, Qt::QueuedConnection);
}

void QGuiApplication_DisconnectSignalFontDatabaseChanged(QtObjectPtr ptr)
{
	QObject::disconnect(((QGuiApplication*)(ptr)), &QGuiApplication::fontDatabaseChanged, ((MyQGuiApplication*)(ptr)), &MyQGuiApplication::Signal_FontDatabaseChanged);
}

void QGuiApplication_ConnectSignalLastWindowClosed(QtObjectPtr ptr)
{
	QObject::connect(((QGuiApplication*)(ptr)), &QGuiApplication::lastWindowClosed, ((MyQGuiApplication*)(ptr)), &MyQGuiApplication::Signal_LastWindowClosed, Qt::QueuedConnection);
}

void QGuiApplication_DisconnectSignalLastWindowClosed(QtObjectPtr ptr)
{
	QObject::disconnect(((QGuiApplication*)(ptr)), &QGuiApplication::lastWindowClosed, ((MyQGuiApplication*)(ptr)), &MyQGuiApplication::Signal_LastWindowClosed);
}

//Static Public Members
char* QGuiApplication_ApplicationDisplayName()
{
	return QGuiApplication::applicationDisplayName().toUtf8().data();
}

int QGuiApplication_ApplicationState()
{
	return QGuiApplication::applicationState();
}

int QGuiApplication_DesktopSettingsAware()
{
	return QGuiApplication::desktopSettingsAware();
}

QtObjectPtr QGuiApplication_FocusObject()
{
	return QGuiApplication::focusObject();
}

int QGuiApplication_IsLeftToRight()
{
	return QGuiApplication::isLeftToRight();
}

int QGuiApplication_IsRightToLeft()
{
	return QGuiApplication::isRightToLeft();
}

int QGuiApplication_KeyboardModifiers()
{
	return QGuiApplication::keyboardModifiers();
}

int QGuiApplication_LayoutDirection()
{
	return QGuiApplication::layoutDirection();
}

int QGuiApplication_MouseButtons()
{
	return QGuiApplication::mouseButtons();
}

char* QGuiApplication_PlatformName()
{
	return QGuiApplication::platformName().toUtf8().data();
}

int QGuiApplication_QueryKeyboardModifiers()
{
	return QGuiApplication::queryKeyboardModifiers();
}

int QGuiApplication_QuitOnLastWindowClosed()
{
	return QGuiApplication::quitOnLastWindowClosed();
}

void QGuiApplication_SetApplicationDisplayName_String(char* name)
{
	QGuiApplication::setApplicationDisplayName(QString(name));
}

void QGuiApplication_SetDesktopSettingsAware_Bool(int on)
{
	QGuiApplication::setDesktopSettingsAware(on != 0);
}

void QGuiApplication_SetLayoutDirection_LayoutDirection(int direction)
{
	QGuiApplication::setLayoutDirection(((Qt::LayoutDirection)(direction)));
}

void QGuiApplication_SetQuitOnLastWindowClosed_Bool(int quit)
{
	QGuiApplication::setQuitOnLastWindowClosed(quit != 0);
}

