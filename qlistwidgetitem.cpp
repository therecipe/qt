#include "qlistwidgetitem.h"
#include <QListWidgetItem>
#include "cgoexport.h"



//Public Functions
QtObjectPtr QListWidgetItem_New_QListWidget_Int(QtObjectPtr parent, int typ)
{
	return new QListWidgetItem(((QListWidget*)(parent)), typ);
}

QtObjectPtr QListWidgetItem_New_String_QListWidget_Int(char* text, QtObjectPtr parent, int typ)
{
	return new QListWidgetItem(QString(text), ((QListWidget*)(parent)), typ);
}

void QListWidgetItem_Destroy(QtObjectPtr ptr)
{
	((QListWidgetItem*)(ptr))->~QListWidgetItem();
}

int QListWidgetItem_CheckState(QtObjectPtr ptr)
{
	return ((QListWidgetItem*)(ptr))->checkState();
}

QtObjectPtr QListWidgetItem_Clone(QtObjectPtr ptr)
{
	return ((QListWidgetItem*)(ptr))->clone();
}

int QListWidgetItem_Flags(QtObjectPtr ptr)
{
	return ((QListWidgetItem*)(ptr))->flags();
}

int QListWidgetItem_IsHidden(QtObjectPtr ptr)
{
	return ((QListWidgetItem*)(ptr))->isHidden();
}

int QListWidgetItem_IsSelected(QtObjectPtr ptr)
{
	return ((QListWidgetItem*)(ptr))->isSelected();
}

QtObjectPtr QListWidgetItem_ListWidget(QtObjectPtr ptr)
{
	return ((QListWidgetItem*)(ptr))->listWidget();
}

void QListWidgetItem_SetCheckState_CheckState(QtObjectPtr ptr, int state)
{
	((QListWidgetItem*)(ptr))->setCheckState(((Qt::CheckState)(state)));
}

void QListWidgetItem_SetFlags_ItemFlag(QtObjectPtr ptr, int flags)
{
	((QListWidgetItem*)(ptr))->setFlags(((Qt::ItemFlag)(flags)));
}

void QListWidgetItem_SetHidden_Bool(QtObjectPtr ptr, int hide)
{
	((QListWidgetItem*)(ptr))->setHidden(hide != 0);
}

void QListWidgetItem_SetSelected_Bool(QtObjectPtr ptr, int selected)
{
	((QListWidgetItem*)(ptr))->setSelected(selected != 0);
}

void QListWidgetItem_SetStatusTip_String(QtObjectPtr ptr, char* statusTip)
{
	((QListWidgetItem*)(ptr))->setStatusTip(QString(statusTip));
}

void QListWidgetItem_SetText_String(QtObjectPtr ptr, char* text)
{
	((QListWidgetItem*)(ptr))->setText(QString(text));
}

void QListWidgetItem_SetTextAlignment_Int(QtObjectPtr ptr, int alignment)
{
	((QListWidgetItem*)(ptr))->setTextAlignment(alignment);
}

void QListWidgetItem_SetToolTip_String(QtObjectPtr ptr, char* toolTip)
{
	((QListWidgetItem*)(ptr))->setToolTip(QString(toolTip));
}

void QListWidgetItem_SetWhatsThis_String(QtObjectPtr ptr, char* whatsThis)
{
	((QListWidgetItem*)(ptr))->setWhatsThis(QString(whatsThis));
}

char* QListWidgetItem_StatusTip(QtObjectPtr ptr)
{
	return ((QListWidgetItem*)(ptr))->statusTip().toUtf8().data();
}

char* QListWidgetItem_Text(QtObjectPtr ptr)
{
	return ((QListWidgetItem*)(ptr))->text().toUtf8().data();
}

int QListWidgetItem_TextAlignment(QtObjectPtr ptr)
{
	return ((QListWidgetItem*)(ptr))->textAlignment();
}

char* QListWidgetItem_ToolTip(QtObjectPtr ptr)
{
	return ((QListWidgetItem*)(ptr))->toolTip().toUtf8().data();
}

int QListWidgetItem_Type(QtObjectPtr ptr)
{
	return ((QListWidgetItem*)(ptr))->type();
}

char* QListWidgetItem_WhatsThis(QtObjectPtr ptr)
{
	return ((QListWidgetItem*)(ptr))->whatsThis().toUtf8().data();
}

