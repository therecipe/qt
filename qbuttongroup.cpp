#include "qbuttongroup.h"
#include <QButtonGroup>
#include "cgoexport.h"



//Public Functions
QtObjectPtr QButtonGroup_New_QObject(QtObjectPtr parent)
{
	return new QButtonGroup(((QObject*)(parent)));
}

void QButtonGroup_Destroy(QtObjectPtr ptr)
{
	((QButtonGroup*)(ptr))->~QButtonGroup();
}

void QButtonGroup_AddButton_QAbstractButton_Int(QtObjectPtr ptr, QtObjectPtr button, int id)
{
	((QButtonGroup*)(ptr))->addButton(((QAbstractButton*)(button)), id);
}

QtObjectPtr QButtonGroup_Button_Int(QtObjectPtr ptr, int id)
{
	return ((QButtonGroup*)(ptr))->button(id);
}

QtObjectPtr QButtonGroup_CheckedButton(QtObjectPtr ptr)
{
	return ((QButtonGroup*)(ptr))->checkedButton();
}

int QButtonGroup_CheckedId(QtObjectPtr ptr)
{
	return ((QButtonGroup*)(ptr))->checkedId();
}

int QButtonGroup_Exclusive(QtObjectPtr ptr)
{
	return ((QButtonGroup*)(ptr))->exclusive();
}

int QButtonGroup_Id_QAbstractButton(QtObjectPtr ptr, QtObjectPtr button)
{
	return ((QButtonGroup*)(ptr))->id(((QAbstractButton*)(button)));
}

void QButtonGroup_RemoveButton_QAbstractButton(QtObjectPtr ptr, QtObjectPtr button)
{
	((QButtonGroup*)(ptr))->removeButton(((QAbstractButton*)(button)));
}

void QButtonGroup_SetExclusive_Bool(QtObjectPtr ptr, int exclusive)
{
	((QButtonGroup*)(ptr))->setExclusive(exclusive != 0);
}

void QButtonGroup_SetId_QAbstractButton_Int(QtObjectPtr ptr, QtObjectPtr button, int id)
{
	((QButtonGroup*)(ptr))->setId(((QAbstractButton*)(button)), id);
}

