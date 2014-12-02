#include "qtoolbox.h"
#include <QToolBox>
#include "cgoexport.h"

//MyQToolBox
class MyQToolBox: public QToolBox {
Q_OBJECT
public:
void Signal_CurrentChanged() { callbackQToolBox(this, QString("currentChanged").toUtf8().data()); };

signals:
void Slot_SetCurrentIndex(int index);

};
#include "qtoolbox.moc"


//Public Functions
QtObjectPtr QToolBox_New_QWidget_WindowType(QtObjectPtr parent, int f)
{
	return new QToolBox(((QWidget*)(parent)), ((Qt::WindowType)(f)));
}

void QToolBox_Destroy(QtObjectPtr ptr)
{
	((QToolBox*)(ptr))->~QToolBox();
}

int QToolBox_AddItem_QWidget_String(QtObjectPtr ptr, QtObjectPtr w, char* text)
{
	return ((QToolBox*)(ptr))->addItem(((QWidget*)(w)), QString(text));
}

int QToolBox_Count(QtObjectPtr ptr)
{
	return ((QToolBox*)(ptr))->count();
}

int QToolBox_CurrentIndex(QtObjectPtr ptr)
{
	return ((QToolBox*)(ptr))->currentIndex();
}

QtObjectPtr QToolBox_CurrentWidget(QtObjectPtr ptr)
{
	return ((QToolBox*)(ptr))->currentWidget();
}

int QToolBox_IndexOf_QWidget(QtObjectPtr ptr, QtObjectPtr widget)
{
	return ((QToolBox*)(ptr))->indexOf(((QWidget*)(widget)));
}

int QToolBox_InsertItem_Int_QWidget_String(QtObjectPtr ptr, int index, QtObjectPtr widget, char* text)
{
	return ((QToolBox*)(ptr))->insertItem(index, ((QWidget*)(widget)), QString(text));
}

int QToolBox_IsItemEnabled_Int(QtObjectPtr ptr, int index)
{
	return ((QToolBox*)(ptr))->isItemEnabled(index);
}

char* QToolBox_ItemText_Int(QtObjectPtr ptr, int index)
{
	return ((QToolBox*)(ptr))->itemText(index).toUtf8().data();
}

char* QToolBox_ItemToolTip_Int(QtObjectPtr ptr, int index)
{
	return ((QToolBox*)(ptr))->itemToolTip(index).toUtf8().data();
}

void QToolBox_RemoveItem_Int(QtObjectPtr ptr, int index)
{
	((QToolBox*)(ptr))->removeItem(index);
}

void QToolBox_SetItemEnabled_Int_Bool(QtObjectPtr ptr, int index, int enabled)
{
	((QToolBox*)(ptr))->setItemEnabled(index, enabled != 0);
}

void QToolBox_SetItemText_Int_String(QtObjectPtr ptr, int index, char* text)
{
	((QToolBox*)(ptr))->setItemText(index, QString(text));
}

void QToolBox_SetItemToolTip_Int_String(QtObjectPtr ptr, int index, char* toolTip)
{
	((QToolBox*)(ptr))->setItemToolTip(index, QString(toolTip));
}

QtObjectPtr QToolBox_Widget_Int(QtObjectPtr ptr, int index)
{
	return ((QToolBox*)(ptr))->widget(index);
}

//Public Slots
void QToolBox_ConnectSlotSetCurrentIndex(QtObjectPtr ptr)
{
	QObject::connect(((MyQToolBox*)(ptr)), &MyQToolBox::Slot_SetCurrentIndex, ((QToolBox*)(ptr)), &QToolBox::setCurrentIndex, Qt::QueuedConnection);
}

void QToolBox_DisconnectSlotSetCurrentIndex(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQToolBox*)(ptr)), &MyQToolBox::Slot_SetCurrentIndex, ((QToolBox*)(ptr)), &QToolBox::setCurrentIndex);
}

void QToolBox_SetCurrentIndex_Int(QtObjectPtr ptr, int index)
{
	((MyQToolBox*)(ptr))->Slot_SetCurrentIndex(index);
}

//Signals
void QToolBox_ConnectSignalCurrentChanged(QtObjectPtr ptr)
{
	QObject::connect(((QToolBox*)(ptr)), &QToolBox::currentChanged, ((MyQToolBox*)(ptr)), &MyQToolBox::Signal_CurrentChanged, Qt::QueuedConnection);
}

void QToolBox_DisconnectSignalCurrentChanged(QtObjectPtr ptr)
{
	QObject::disconnect(((QToolBox*)(ptr)), &QToolBox::currentChanged, ((MyQToolBox*)(ptr)), &MyQToolBox::Signal_CurrentChanged);
}

