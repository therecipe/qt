#include "qtoolbutton.h"
#include <QToolButton>
#include "cgoexport.h"

//MyQToolButton
class MyQToolButton: public QToolButton {
Q_OBJECT
public:
void Signal_Triggered() { callbackQToolButton(this, QString("triggered").toUtf8().data()); };

signals:
void Slot_ShowMenu();

};
#include "qtoolbutton.moc"

#include <QAction>

//Public Functions
QtObjectPtr QToolButton_New_QWidget(QtObjectPtr parent)
{
	return new QToolButton(((QWidget*)(parent)));
}

void QToolButton_Destroy(QtObjectPtr ptr)
{
	((QToolButton*)(ptr))->~QToolButton();
}

int QToolButton_ArrowType(QtObjectPtr ptr)
{
	return ((QToolButton*)(ptr))->arrowType();
}

int QToolButton_AutoRaise(QtObjectPtr ptr)
{
	return ((QToolButton*)(ptr))->autoRaise();
}

QtObjectPtr QToolButton_Menu(QtObjectPtr ptr)
{
	return ((QToolButton*)(ptr))->menu();
}

void QToolButton_SetArrowType_ArrowType(QtObjectPtr ptr, int typ)
{
	((QToolButton*)(ptr))->setArrowType(((Qt::ArrowType)(typ)));
}

void QToolButton_SetAutoRaise_Bool(QtObjectPtr ptr, int enable)
{
	((QToolButton*)(ptr))->setAutoRaise(enable != 0);
}

void QToolButton_SetMenu_QMenu(QtObjectPtr ptr, QtObjectPtr menu)
{
	((QToolButton*)(ptr))->setMenu(((QMenu*)(menu)));
}

int QToolButton_ToolButtonStyle(QtObjectPtr ptr)
{
	return ((QToolButton*)(ptr))->toolButtonStyle();
}

//Public Slots
void QToolButton_ConnectSlotShowMenu(QtObjectPtr ptr)
{
	QObject::connect(((MyQToolButton*)(ptr)), &MyQToolButton::Slot_ShowMenu, ((QToolButton*)(ptr)), &QToolButton::showMenu, Qt::QueuedConnection);
}

void QToolButton_DisconnectSlotShowMenu(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQToolButton*)(ptr)), &MyQToolButton::Slot_ShowMenu, ((QToolButton*)(ptr)), &QToolButton::showMenu);
}

void QToolButton_ShowMenu(QtObjectPtr ptr)
{
	((MyQToolButton*)(ptr))->Slot_ShowMenu();
}

//Signals
void QToolButton_ConnectSignalTriggered(QtObjectPtr ptr)
{
	QObject::connect(((QToolButton*)(ptr)), &QToolButton::triggered, ((MyQToolButton*)(ptr)), &MyQToolButton::Signal_Triggered, Qt::QueuedConnection);
}

void QToolButton_DisconnectSignalTriggered(QtObjectPtr ptr)
{
	QObject::disconnect(((QToolButton*)(ptr)), &QToolButton::triggered, ((MyQToolButton*)(ptr)), &MyQToolButton::Signal_Triggered);
}

