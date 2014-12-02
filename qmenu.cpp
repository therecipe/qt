#include "qmenu.h"
#include <QMenu>
#include "cgoexport.h"

//MyQMenu
class MyQMenu: public QMenu {
Q_OBJECT
public:
void Signal_AboutToHide() { callbackQMenu(this, QString("aboutToHide").toUtf8().data()); };
void Signal_AboutToShow() { callbackQMenu(this, QString("aboutToShow").toUtf8().data()); };
void Signal_Hovered() { callbackQMenu(this, QString("hovered").toUtf8().data()); };
void Signal_Triggered() { callbackQMenu(this, QString("triggered").toUtf8().data()); };

signals:

};
#include "qmenu.moc"


//Public Functions
QtObjectPtr QMenu_New_QWidget(QtObjectPtr parent)
{
	return new QMenu(((QWidget*)(parent)));
}

QtObjectPtr QMenu_New_String_QWidget(char* title, QtObjectPtr parent)
{
	return new QMenu(QString(title), ((QWidget*)(parent)));
}

void QMenu_Destroy(QtObjectPtr ptr)
{
	((QMenu*)(ptr))->~QMenu();
}

QtObjectPtr QMenu_ActiveAction(QtObjectPtr ptr)
{
	return ((QMenu*)(ptr))->activeAction();
}

QtObjectPtr QMenu_AddAction_String(QtObjectPtr ptr, char* text)
{
	return ((QMenu*)(ptr))->addAction(QString(text));
}

void QMenu_AddAction_QAction(QtObjectPtr ptr, QtObjectPtr action)
{
	((QMenu*)(ptr))->addAction(((QAction*)(action)));
}

QtObjectPtr QMenu_AddMenu_QMenu(QtObjectPtr ptr, QtObjectPtr menu)
{
	return ((QMenu*)(ptr))->addMenu(((QMenu*)(menu)));
}

QtObjectPtr QMenu_AddMenu_String(QtObjectPtr ptr, char* title)
{
	return ((QMenu*)(ptr))->addMenu(QString(title));
}

QtObjectPtr QMenu_AddSection_String(QtObjectPtr ptr, char* text)
{
	return ((QMenu*)(ptr))->addSection(QString(text));
}

QtObjectPtr QMenu_AddSeparator(QtObjectPtr ptr)
{
	return ((QMenu*)(ptr))->addSeparator();
}

void QMenu_Clear(QtObjectPtr ptr)
{
	((QMenu*)(ptr))->clear();
}

QtObjectPtr QMenu_DefaultAction(QtObjectPtr ptr)
{
	return ((QMenu*)(ptr))->defaultAction();
}

QtObjectPtr QMenu_Exec(QtObjectPtr ptr)
{
	return ((QMenu*)(ptr))->exec();
}

void QMenu_HideTearOffMenu(QtObjectPtr ptr)
{
	((QMenu*)(ptr))->hideTearOffMenu();
}

QtObjectPtr QMenu_InsertMenu_QAction_QMenu(QtObjectPtr ptr, QtObjectPtr before, QtObjectPtr menu)
{
	return ((QMenu*)(ptr))->insertMenu(((QAction*)(before)), ((QMenu*)(menu)));
}

QtObjectPtr QMenu_InsertSection_QAction_String(QtObjectPtr ptr, QtObjectPtr before, char* text)
{
	return ((QMenu*)(ptr))->insertSection(((QAction*)(before)), QString(text));
}

QtObjectPtr QMenu_InsertSeparator_QAction(QtObjectPtr ptr, QtObjectPtr before)
{
	return ((QMenu*)(ptr))->insertSeparator(((QAction*)(before)));
}

int QMenu_IsEmpty(QtObjectPtr ptr)
{
	return ((QMenu*)(ptr))->isEmpty();
}

int QMenu_IsTearOffEnabled(QtObjectPtr ptr)
{
	return ((QMenu*)(ptr))->isTearOffEnabled();
}

int QMenu_IsTearOffMenuVisible(QtObjectPtr ptr)
{
	return ((QMenu*)(ptr))->isTearOffMenuVisible();
}

QtObjectPtr QMenu_MenuAction(QtObjectPtr ptr)
{
	return ((QMenu*)(ptr))->menuAction();
}

int QMenu_SeparatorsCollapsible(QtObjectPtr ptr)
{
	return ((QMenu*)(ptr))->separatorsCollapsible();
}

void QMenu_SetActiveAction_QAction(QtObjectPtr ptr, QtObjectPtr act)
{
	((QMenu*)(ptr))->setActiveAction(((QAction*)(act)));
}

void QMenu_SetDefaultAction_QAction(QtObjectPtr ptr, QtObjectPtr act)
{
	((QMenu*)(ptr))->setDefaultAction(((QAction*)(act)));
}

void QMenu_SetSeparatorsCollapsible_Bool(QtObjectPtr ptr, int collapse)
{
	((QMenu*)(ptr))->setSeparatorsCollapsible(collapse != 0);
}

void QMenu_SetTearOffEnabled_Bool(QtObjectPtr ptr, int tearOffEnabled)
{
	((QMenu*)(ptr))->setTearOffEnabled(tearOffEnabled != 0);
}

void QMenu_SetTitle_String(QtObjectPtr ptr, char* title)
{
	((QMenu*)(ptr))->setTitle(QString(title));
}

void QMenu_SetToolTipsVisible_Bool(QtObjectPtr ptr, int visible)
{
	((QMenu*)(ptr))->setToolTipsVisible(visible != 0);
}

char* QMenu_Title(QtObjectPtr ptr)
{
	return ((QMenu*)(ptr))->title().toUtf8().data();
}

int QMenu_ToolTipsVisible(QtObjectPtr ptr)
{
	return ((QMenu*)(ptr))->toolTipsVisible();
}

//Signals
void QMenu_ConnectSignalAboutToHide(QtObjectPtr ptr)
{
	QObject::connect(((QMenu*)(ptr)), &QMenu::aboutToHide, ((MyQMenu*)(ptr)), &MyQMenu::Signal_AboutToHide, Qt::QueuedConnection);
}

void QMenu_DisconnectSignalAboutToHide(QtObjectPtr ptr)
{
	QObject::disconnect(((QMenu*)(ptr)), &QMenu::aboutToHide, ((MyQMenu*)(ptr)), &MyQMenu::Signal_AboutToHide);
}

void QMenu_ConnectSignalAboutToShow(QtObjectPtr ptr)
{
	QObject::connect(((QMenu*)(ptr)), &QMenu::aboutToShow, ((MyQMenu*)(ptr)), &MyQMenu::Signal_AboutToShow, Qt::QueuedConnection);
}

void QMenu_DisconnectSignalAboutToShow(QtObjectPtr ptr)
{
	QObject::disconnect(((QMenu*)(ptr)), &QMenu::aboutToShow, ((MyQMenu*)(ptr)), &MyQMenu::Signal_AboutToShow);
}

void QMenu_ConnectSignalHovered(QtObjectPtr ptr)
{
	QObject::connect(((QMenu*)(ptr)), &QMenu::hovered, ((MyQMenu*)(ptr)), &MyQMenu::Signal_Hovered, Qt::QueuedConnection);
}

void QMenu_DisconnectSignalHovered(QtObjectPtr ptr)
{
	QObject::disconnect(((QMenu*)(ptr)), &QMenu::hovered, ((MyQMenu*)(ptr)), &MyQMenu::Signal_Hovered);
}

void QMenu_ConnectSignalTriggered(QtObjectPtr ptr)
{
	QObject::connect(((QMenu*)(ptr)), &QMenu::triggered, ((MyQMenu*)(ptr)), &MyQMenu::Signal_Triggered, Qt::QueuedConnection);
}

void QMenu_DisconnectSignalTriggered(QtObjectPtr ptr)
{
	QObject::disconnect(((QMenu*)(ptr)), &QMenu::triggered, ((MyQMenu*)(ptr)), &MyQMenu::Signal_Triggered);
}

