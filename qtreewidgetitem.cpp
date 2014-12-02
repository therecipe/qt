#include "qtreewidgetitem.h"
#include <QTreeWidgetItem>
#include "cgoexport.h"



//Public Functions
QtObjectPtr QTreeWidgetItem_New_Int(int typ)
{
	return new QTreeWidgetItem(typ);
}

QtObjectPtr QTreeWidgetItem_New_QTreeWidget_Int(QtObjectPtr parent, int typ)
{
	return new QTreeWidgetItem(((QTreeWidget*)(parent)), typ);
}

QtObjectPtr QTreeWidgetItem_New_QTreeWidget_QTreeWidgetItem_Int(QtObjectPtr parent, QtObjectPtr preceding, int typ)
{
	return new QTreeWidgetItem(((QTreeWidget*)(parent)), ((QTreeWidgetItem*)(preceding)), typ);
}

QtObjectPtr QTreeWidgetItem_New_QTreeWidgetItem_Int(QtObjectPtr parent, int typ)
{
	return new QTreeWidgetItem(((QTreeWidgetItem*)(parent)), typ);
}

QtObjectPtr QTreeWidgetItem_New_QTreeWidgetItem_QTreeWidgetItem_Int(QtObjectPtr parent, QtObjectPtr preceding, int typ)
{
	return new QTreeWidgetItem(((QTreeWidgetItem*)(parent)), ((QTreeWidgetItem*)(preceding)), typ);
}

QtObjectPtr QTreeWidgetItem_New_QTreeWidgetItem(QtObjectPtr other)
{
	return new QTreeWidgetItem(((QTreeWidgetItem*)(other)));
}

void QTreeWidgetItem_Destroy(QtObjectPtr ptr)
{
	((QTreeWidgetItem*)(ptr))->~QTreeWidgetItem();
}

void QTreeWidgetItem_AddChild_QTreeWidgetItem(QtObjectPtr ptr, QtObjectPtr child)
{
	((QTreeWidgetItem*)(ptr))->addChild(((QTreeWidgetItem*)(child)));
}

int QTreeWidgetItem_CheckState_Int(QtObjectPtr ptr, int column)
{
	return ((QTreeWidgetItem*)(ptr))->checkState(column);
}

QtObjectPtr QTreeWidgetItem_Child_Int(QtObjectPtr ptr, int index)
{
	return ((QTreeWidgetItem*)(ptr))->child(index);
}

int QTreeWidgetItem_ChildCount(QtObjectPtr ptr)
{
	return ((QTreeWidgetItem*)(ptr))->childCount();
}

QtObjectPtr QTreeWidgetItem_Clone(QtObjectPtr ptr)
{
	return ((QTreeWidgetItem*)(ptr))->clone();
}

int QTreeWidgetItem_ColumnCount(QtObjectPtr ptr)
{
	return ((QTreeWidgetItem*)(ptr))->columnCount();
}

int QTreeWidgetItem_Flags(QtObjectPtr ptr)
{
	return ((QTreeWidgetItem*)(ptr))->flags();
}

int QTreeWidgetItem_IndexOfChild_QTreeWidgetItem(QtObjectPtr ptr, QtObjectPtr child)
{
	return ((QTreeWidgetItem*)(ptr))->indexOfChild(((QTreeWidgetItem*)(child)));
}

void QTreeWidgetItem_InsertChild_Int_QTreeWidgetItem(QtObjectPtr ptr, int index, QtObjectPtr child)
{
	((QTreeWidgetItem*)(ptr))->insertChild(index, ((QTreeWidgetItem*)(child)));
}

int QTreeWidgetItem_IsDisabled(QtObjectPtr ptr)
{
	return ((QTreeWidgetItem*)(ptr))->isDisabled();
}

int QTreeWidgetItem_IsExpanded(QtObjectPtr ptr)
{
	return ((QTreeWidgetItem*)(ptr))->isExpanded();
}

int QTreeWidgetItem_IsFirstColumnSpanned(QtObjectPtr ptr)
{
	return ((QTreeWidgetItem*)(ptr))->isFirstColumnSpanned();
}

int QTreeWidgetItem_IsHidden(QtObjectPtr ptr)
{
	return ((QTreeWidgetItem*)(ptr))->isHidden();
}

int QTreeWidgetItem_IsSelected(QtObjectPtr ptr)
{
	return ((QTreeWidgetItem*)(ptr))->isSelected();
}

QtObjectPtr QTreeWidgetItem_Parent(QtObjectPtr ptr)
{
	return ((QTreeWidgetItem*)(ptr))->parent();
}

void QTreeWidgetItem_RemoveChild_QTreeWidgetItem(QtObjectPtr ptr, QtObjectPtr child)
{
	((QTreeWidgetItem*)(ptr))->removeChild(((QTreeWidgetItem*)(child)));
}

void QTreeWidgetItem_SetBackground_Int_QBrush(QtObjectPtr ptr, int column, QtObjectPtr brush)
{
	((QTreeWidgetItem*)(ptr))->setBackground(column, *((QBrush*)(brush)));
}

void QTreeWidgetItem_SetCheckState_Int_CheckState(QtObjectPtr ptr, int column, int state)
{
	((QTreeWidgetItem*)(ptr))->setCheckState(column, ((Qt::CheckState)(state)));
}

void QTreeWidgetItem_SetDisabled_Bool(QtObjectPtr ptr, int disabled)
{
	((QTreeWidgetItem*)(ptr))->setDisabled(disabled != 0);
}

void QTreeWidgetItem_SetExpanded_Bool(QtObjectPtr ptr, int expand)
{
	((QTreeWidgetItem*)(ptr))->setExpanded(expand != 0);
}

void QTreeWidgetItem_SetFirstColumnSpanned_Bool(QtObjectPtr ptr, int span)
{
	((QTreeWidgetItem*)(ptr))->setFirstColumnSpanned(span != 0);
}

void QTreeWidgetItem_SetFlags_ItemFlag(QtObjectPtr ptr, int flags)
{
	((QTreeWidgetItem*)(ptr))->setFlags(((Qt::ItemFlag)(flags)));
}

void QTreeWidgetItem_SetForeground_Int_QBrush(QtObjectPtr ptr, int column, QtObjectPtr brush)
{
	((QTreeWidgetItem*)(ptr))->setForeground(column, *((QBrush*)(brush)));
}

void QTreeWidgetItem_SetHidden_Bool(QtObjectPtr ptr, int hide)
{
	((QTreeWidgetItem*)(ptr))->setHidden(hide != 0);
}

void QTreeWidgetItem_SetSelected_Bool(QtObjectPtr ptr, int selected)
{
	((QTreeWidgetItem*)(ptr))->setSelected(selected != 0);
}

void QTreeWidgetItem_SetStatusTip_Int_String(QtObjectPtr ptr, int column, char* statusTip)
{
	((QTreeWidgetItem*)(ptr))->setStatusTip(column, QString(statusTip));
}

void QTreeWidgetItem_SetText_Int_String(QtObjectPtr ptr, int column, char* text)
{
	((QTreeWidgetItem*)(ptr))->setText(column, QString(text));
}

void QTreeWidgetItem_SetTextAlignment_Int_Int(QtObjectPtr ptr, int column, int alignment)
{
	((QTreeWidgetItem*)(ptr))->setTextAlignment(column, alignment);
}

void QTreeWidgetItem_SetToolTip_Int_String(QtObjectPtr ptr, int column, char* toolTip)
{
	((QTreeWidgetItem*)(ptr))->setToolTip(column, QString(toolTip));
}

void QTreeWidgetItem_SetWhatsThis_Int_String(QtObjectPtr ptr, int column, char* whatsThis)
{
	((QTreeWidgetItem*)(ptr))->setWhatsThis(column, QString(whatsThis));
}

void QTreeWidgetItem_SortChildren_Int_SortOrder(QtObjectPtr ptr, int column, int order)
{
	((QTreeWidgetItem*)(ptr))->sortChildren(column, ((Qt::SortOrder)(order)));
}

char* QTreeWidgetItem_StatusTip_Int(QtObjectPtr ptr, int column)
{
	return ((QTreeWidgetItem*)(ptr))->statusTip(column).toUtf8().data();
}

QtObjectPtr QTreeWidgetItem_TakeChild_Int(QtObjectPtr ptr, int index)
{
	return ((QTreeWidgetItem*)(ptr))->takeChild(index);
}

char* QTreeWidgetItem_Text_Int(QtObjectPtr ptr, int column)
{
	return ((QTreeWidgetItem*)(ptr))->text(column).toUtf8().data();
}

int QTreeWidgetItem_TextAlignment_Int(QtObjectPtr ptr, int column)
{
	return ((QTreeWidgetItem*)(ptr))->textAlignment(column);
}

char* QTreeWidgetItem_ToolTip_Int(QtObjectPtr ptr, int column)
{
	return ((QTreeWidgetItem*)(ptr))->toolTip(column).toUtf8().data();
}

QtObjectPtr QTreeWidgetItem_TreeWidget(QtObjectPtr ptr)
{
	return ((QTreeWidgetItem*)(ptr))->treeWidget();
}

int QTreeWidgetItem_Type(QtObjectPtr ptr)
{
	return ((QTreeWidgetItem*)(ptr))->type();
}

char* QTreeWidgetItem_WhatsThis_Int(QtObjectPtr ptr, int column)
{
	return ((QTreeWidgetItem*)(ptr))->whatsThis(column).toUtf8().data();
}

