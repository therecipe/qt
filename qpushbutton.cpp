#include "qpushbutton.h"
#include <QPushButton>
#include "cgoexport.h"

//MyQPushButton
class MyQPushButton: public QPushButton {
Q_OBJECT
public:

signals:
void Slot_ShowMenu();

};
#include "qpushbutton.moc"


//Public Functions
QtObjectPtr QPushButton_New_QWidget(QtObjectPtr parent)
{
	return new QPushButton(((QWidget*)(parent)));
}

QtObjectPtr QPushButton_New_String_QWidget(char* text, QtObjectPtr parent)
{
	return new QPushButton(QString(text), ((QWidget*)(parent)));
}

void QPushButton_Destroy(QtObjectPtr ptr)
{
	((QPushButton*)(ptr))->~QPushButton();
}

int QPushButton_AutoDefault(QtObjectPtr ptr)
{
	return ((QPushButton*)(ptr))->autoDefault();
}

int QPushButton_IsDefault(QtObjectPtr ptr)
{
	return ((QPushButton*)(ptr))->isDefault();
}

int QPushButton_IsFlat(QtObjectPtr ptr)
{
	return ((QPushButton*)(ptr))->isFlat();
}

QtObjectPtr QPushButton_Menu(QtObjectPtr ptr)
{
	return ((QPushButton*)(ptr))->menu();
}

void QPushButton_SetAutoDefault_Bool(QtObjectPtr ptr, int autoDefault)
{
	((QPushButton*)(ptr))->setAutoDefault(autoDefault != 0);
}

void QPushButton_SetDefault_Bool(QtObjectPtr ptr, int defaul)
{
	((QPushButton*)(ptr))->setDefault(defaul != 0);
}

void QPushButton_SetFlat_Bool(QtObjectPtr ptr, int flat)
{
	((QPushButton*)(ptr))->setFlat(flat != 0);
}

void QPushButton_SetMenu_QMenu(QtObjectPtr ptr, QtObjectPtr menu)
{
	((QPushButton*)(ptr))->setMenu(((QMenu*)(menu)));
}

//Public Slots
void QPushButton_ConnectSlotShowMenu(QtObjectPtr ptr)
{
	QObject::connect(((MyQPushButton*)(ptr)), &MyQPushButton::Slot_ShowMenu, ((QPushButton*)(ptr)), &QPushButton::showMenu, Qt::QueuedConnection);
}

void QPushButton_DisconnectSlotShowMenu(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQPushButton*)(ptr)), &MyQPushButton::Slot_ShowMenu, ((QPushButton*)(ptr)), &QPushButton::showMenu);
}

void QPushButton_ShowMenu(QtObjectPtr ptr)
{
	((MyQPushButton*)(ptr))->Slot_ShowMenu();
}

