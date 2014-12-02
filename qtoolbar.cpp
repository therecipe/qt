#include "qtoolbar.h"
#include <QToolBar>
#include "cgoexport.h"

//MyQToolBar
class MyQToolBar: public QToolBar {
Q_OBJECT
public:
void Signal_ActionTriggered() { callbackQToolBar(this, QString("actionTriggered").toUtf8().data()); };
void Signal_AllowedAreasChanged() { callbackQToolBar(this, QString("allowedAreasChanged").toUtf8().data()); };
void Signal_MovableChanged() { callbackQToolBar(this, QString("movableChanged").toUtf8().data()); };
void Signal_OrientationChanged() { callbackQToolBar(this, QString("orientationChanged").toUtf8().data()); };
void Signal_ToolButtonStyleChanged() { callbackQToolBar(this, QString("toolButtonStyleChanged").toUtf8().data()); };
void Signal_TopLevelChanged() { callbackQToolBar(this, QString("topLevelChanged").toUtf8().data()); };
void Signal_VisibilityChanged() { callbackQToolBar(this, QString("visibilityChanged").toUtf8().data()); };

signals:

};
#include "qtoolbar.moc"

#include <QAction>

//Public Functions
QtObjectPtr QToolBar_New_String_QWidget(char* title, QtObjectPtr parent)
{
	return new QToolBar(QString(title), ((QWidget*)(parent)));
}

QtObjectPtr QToolBar_New_QWidget(QtObjectPtr parent)
{
	return new QToolBar(((QWidget*)(parent)));
}

void QToolBar_Destroy(QtObjectPtr ptr)
{
	((QToolBar*)(ptr))->~QToolBar();
}

int QToolBar_AllowedAreas(QtObjectPtr ptr)
{
	return ((QToolBar*)(ptr))->allowedAreas();
}

void QToolBar_Clear(QtObjectPtr ptr)
{
	((QToolBar*)(ptr))->clear();
}

int QToolBar_IsAreaAllowed_ToolBarArea(QtObjectPtr ptr, int area)
{
	return ((QToolBar*)(ptr))->isAreaAllowed(((Qt::ToolBarArea)(area)));
}

int QToolBar_IsFloatable(QtObjectPtr ptr)
{
	return ((QToolBar*)(ptr))->isFloatable();
}

int QToolBar_IsFloating(QtObjectPtr ptr)
{
	return ((QToolBar*)(ptr))->isFloating();
}

int QToolBar_IsMovable(QtObjectPtr ptr)
{
	return ((QToolBar*)(ptr))->isMovable();
}

int QToolBar_Orientation(QtObjectPtr ptr)
{
	return ((QToolBar*)(ptr))->orientation();
}

void QToolBar_SetAllowedAreas_ToolBarArea(QtObjectPtr ptr, int areas)
{
	((QToolBar*)(ptr))->setAllowedAreas(((Qt::ToolBarArea)(areas)));
}

void QToolBar_SetFloatable_Bool(QtObjectPtr ptr, int floatable)
{
	((QToolBar*)(ptr))->setFloatable(floatable != 0);
}

void QToolBar_SetMovable_Bool(QtObjectPtr ptr, int movable)
{
	((QToolBar*)(ptr))->setMovable(movable != 0);
}

void QToolBar_SetOrientation_Orientation(QtObjectPtr ptr, int orientation)
{
	((QToolBar*)(ptr))->setOrientation(((Qt::Orientation)(orientation)));
}

int QToolBar_ToolButtonStyle(QtObjectPtr ptr)
{
	return ((QToolBar*)(ptr))->toolButtonStyle();
}

//Signals
void QToolBar_ConnectSignalActionTriggered(QtObjectPtr ptr)
{
	QObject::connect(((QToolBar*)(ptr)), &QToolBar::actionTriggered, ((MyQToolBar*)(ptr)), &MyQToolBar::Signal_ActionTriggered, Qt::QueuedConnection);
}

void QToolBar_DisconnectSignalActionTriggered(QtObjectPtr ptr)
{
	QObject::disconnect(((QToolBar*)(ptr)), &QToolBar::actionTriggered, ((MyQToolBar*)(ptr)), &MyQToolBar::Signal_ActionTriggered);
}

void QToolBar_ConnectSignalAllowedAreasChanged(QtObjectPtr ptr)
{
	QObject::connect(((QToolBar*)(ptr)), &QToolBar::allowedAreasChanged, ((MyQToolBar*)(ptr)), &MyQToolBar::Signal_AllowedAreasChanged, Qt::QueuedConnection);
}

void QToolBar_DisconnectSignalAllowedAreasChanged(QtObjectPtr ptr)
{
	QObject::disconnect(((QToolBar*)(ptr)), &QToolBar::allowedAreasChanged, ((MyQToolBar*)(ptr)), &MyQToolBar::Signal_AllowedAreasChanged);
}

void QToolBar_ConnectSignalMovableChanged(QtObjectPtr ptr)
{
	QObject::connect(((QToolBar*)(ptr)), &QToolBar::movableChanged, ((MyQToolBar*)(ptr)), &MyQToolBar::Signal_MovableChanged, Qt::QueuedConnection);
}

void QToolBar_DisconnectSignalMovableChanged(QtObjectPtr ptr)
{
	QObject::disconnect(((QToolBar*)(ptr)), &QToolBar::movableChanged, ((MyQToolBar*)(ptr)), &MyQToolBar::Signal_MovableChanged);
}

void QToolBar_ConnectSignalOrientationChanged(QtObjectPtr ptr)
{
	QObject::connect(((QToolBar*)(ptr)), &QToolBar::orientationChanged, ((MyQToolBar*)(ptr)), &MyQToolBar::Signal_OrientationChanged, Qt::QueuedConnection);
}

void QToolBar_DisconnectSignalOrientationChanged(QtObjectPtr ptr)
{
	QObject::disconnect(((QToolBar*)(ptr)), &QToolBar::orientationChanged, ((MyQToolBar*)(ptr)), &MyQToolBar::Signal_OrientationChanged);
}

void QToolBar_ConnectSignalToolButtonStyleChanged(QtObjectPtr ptr)
{
	QObject::connect(((QToolBar*)(ptr)), &QToolBar::toolButtonStyleChanged, ((MyQToolBar*)(ptr)), &MyQToolBar::Signal_ToolButtonStyleChanged, Qt::QueuedConnection);
}

void QToolBar_DisconnectSignalToolButtonStyleChanged(QtObjectPtr ptr)
{
	QObject::disconnect(((QToolBar*)(ptr)), &QToolBar::toolButtonStyleChanged, ((MyQToolBar*)(ptr)), &MyQToolBar::Signal_ToolButtonStyleChanged);
}

void QToolBar_ConnectSignalTopLevelChanged(QtObjectPtr ptr)
{
	QObject::connect(((QToolBar*)(ptr)), &QToolBar::topLevelChanged, ((MyQToolBar*)(ptr)), &MyQToolBar::Signal_TopLevelChanged, Qt::QueuedConnection);
}

void QToolBar_DisconnectSignalTopLevelChanged(QtObjectPtr ptr)
{
	QObject::disconnect(((QToolBar*)(ptr)), &QToolBar::topLevelChanged, ((MyQToolBar*)(ptr)), &MyQToolBar::Signal_TopLevelChanged);
}

void QToolBar_ConnectSignalVisibilityChanged(QtObjectPtr ptr)
{
	QObject::connect(((QToolBar*)(ptr)), &QToolBar::visibilityChanged, ((MyQToolBar*)(ptr)), &MyQToolBar::Signal_VisibilityChanged, Qt::QueuedConnection);
}

void QToolBar_DisconnectSignalVisibilityChanged(QtObjectPtr ptr)
{
	QObject::disconnect(((QToolBar*)(ptr)), &QToolBar::visibilityChanged, ((MyQToolBar*)(ptr)), &MyQToolBar::Signal_VisibilityChanged);
}

