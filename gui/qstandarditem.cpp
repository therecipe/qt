#include "qstandarditem.h"
#include <QModelIndex>
#include <QIcon>
#include <QSize>
#include <QDataStream>
#include <QFont>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QBrush>
#include <QStandardItem>
#include "_cgo_export.h"

class MyQStandardItem: public QStandardItem {
public:
};

QtObjectPtr QStandardItem_NewQStandardItem(){
	return new QStandardItem();
}

QtObjectPtr QStandardItem_NewQStandardItem3(QtObjectPtr icon, char* text){
	return new QStandardItem(*static_cast<QIcon*>(icon), QString(text));
}

QtObjectPtr QStandardItem_NewQStandardItem2(char* text){
	return new QStandardItem(QString(text));
}

QtObjectPtr QStandardItem_NewQStandardItem4(int rows, int columns){
	return new QStandardItem(rows, columns);
}

char* QStandardItem_AccessibleDescription(QtObjectPtr ptr){
	return static_cast<QStandardItem*>(ptr)->accessibleDescription().toUtf8().data();
}

char* QStandardItem_AccessibleText(QtObjectPtr ptr){
	return static_cast<QStandardItem*>(ptr)->accessibleText().toUtf8().data();
}

void QStandardItem_AppendRow2(QtObjectPtr ptr, QtObjectPtr item){
	static_cast<QStandardItem*>(ptr)->appendRow(static_cast<QStandardItem*>(item));
}

int QStandardItem_CheckState(QtObjectPtr ptr){
	return static_cast<QStandardItem*>(ptr)->checkState();
}

QtObjectPtr QStandardItem_Child(QtObjectPtr ptr, int row, int column){
	return static_cast<QStandardItem*>(ptr)->child(row, column);
}

QtObjectPtr QStandardItem_Clone(QtObjectPtr ptr){
	return static_cast<QStandardItem*>(ptr)->clone();
}

int QStandardItem_Column(QtObjectPtr ptr){
	return static_cast<QStandardItem*>(ptr)->column();
}

int QStandardItem_ColumnCount(QtObjectPtr ptr){
	return static_cast<QStandardItem*>(ptr)->columnCount();
}

char* QStandardItem_Data(QtObjectPtr ptr, int role){
	return static_cast<QStandardItem*>(ptr)->data(role).toString().toUtf8().data();
}

int QStandardItem_Flags(QtObjectPtr ptr){
	return static_cast<QStandardItem*>(ptr)->flags();
}

int QStandardItem_HasChildren(QtObjectPtr ptr){
	return static_cast<QStandardItem*>(ptr)->hasChildren();
}

QtObjectPtr QStandardItem_Index(QtObjectPtr ptr){
	return static_cast<QStandardItem*>(ptr)->index().internalPointer();
}

void QStandardItem_InsertColumns(QtObjectPtr ptr, int column, int count){
	static_cast<QStandardItem*>(ptr)->insertColumns(column, count);
}

void QStandardItem_InsertRow2(QtObjectPtr ptr, int row, QtObjectPtr item){
	static_cast<QStandardItem*>(ptr)->insertRow(row, static_cast<QStandardItem*>(item));
}

void QStandardItem_InsertRows2(QtObjectPtr ptr, int row, int count){
	static_cast<QStandardItem*>(ptr)->insertRows(row, count);
}

int QStandardItem_IsCheckable(QtObjectPtr ptr){
	return static_cast<QStandardItem*>(ptr)->isCheckable();
}

int QStandardItem_IsDragEnabled(QtObjectPtr ptr){
	return static_cast<QStandardItem*>(ptr)->isDragEnabled();
}

int QStandardItem_IsDropEnabled(QtObjectPtr ptr){
	return static_cast<QStandardItem*>(ptr)->isDropEnabled();
}

int QStandardItem_IsEditable(QtObjectPtr ptr){
	return static_cast<QStandardItem*>(ptr)->isEditable();
}

int QStandardItem_IsEnabled(QtObjectPtr ptr){
	return static_cast<QStandardItem*>(ptr)->isEnabled();
}

int QStandardItem_IsSelectable(QtObjectPtr ptr){
	return static_cast<QStandardItem*>(ptr)->isSelectable();
}

int QStandardItem_IsTristate(QtObjectPtr ptr){
	return static_cast<QStandardItem*>(ptr)->isTristate();
}

QtObjectPtr QStandardItem_Model(QtObjectPtr ptr){
	return static_cast<QStandardItem*>(ptr)->model();
}

QtObjectPtr QStandardItem_Parent(QtObjectPtr ptr){
	return static_cast<QStandardItem*>(ptr)->parent();
}

void QStandardItem_Read(QtObjectPtr ptr, QtObjectPtr in){
	static_cast<QStandardItem*>(ptr)->read(*static_cast<QDataStream*>(in));
}

void QStandardItem_RemoveColumn(QtObjectPtr ptr, int column){
	static_cast<QStandardItem*>(ptr)->removeColumn(column);
}

void QStandardItem_RemoveColumns(QtObjectPtr ptr, int column, int count){
	static_cast<QStandardItem*>(ptr)->removeColumns(column, count);
}

void QStandardItem_RemoveRow(QtObjectPtr ptr, int row){
	static_cast<QStandardItem*>(ptr)->removeRow(row);
}

void QStandardItem_RemoveRows(QtObjectPtr ptr, int row, int count){
	static_cast<QStandardItem*>(ptr)->removeRows(row, count);
}

int QStandardItem_Row(QtObjectPtr ptr){
	return static_cast<QStandardItem*>(ptr)->row();
}

int QStandardItem_RowCount(QtObjectPtr ptr){
	return static_cast<QStandardItem*>(ptr)->rowCount();
}

void QStandardItem_SetAccessibleDescription(QtObjectPtr ptr, char* accessibleDescription){
	static_cast<QStandardItem*>(ptr)->setAccessibleDescription(QString(accessibleDescription));
}

void QStandardItem_SetAccessibleText(QtObjectPtr ptr, char* accessibleText){
	static_cast<QStandardItem*>(ptr)->setAccessibleText(QString(accessibleText));
}

void QStandardItem_SetBackground(QtObjectPtr ptr, QtObjectPtr brush){
	static_cast<QStandardItem*>(ptr)->setBackground(*static_cast<QBrush*>(brush));
}

void QStandardItem_SetCheckState(QtObjectPtr ptr, int state){
	static_cast<QStandardItem*>(ptr)->setCheckState(static_cast<Qt::CheckState>(state));
}

void QStandardItem_SetCheckable(QtObjectPtr ptr, int checkable){
	static_cast<QStandardItem*>(ptr)->setCheckable(checkable != 0);
}

void QStandardItem_SetChild2(QtObjectPtr ptr, int row, QtObjectPtr item){
	static_cast<QStandardItem*>(ptr)->setChild(row, static_cast<QStandardItem*>(item));
}

void QStandardItem_SetChild(QtObjectPtr ptr, int row, int column, QtObjectPtr item){
	static_cast<QStandardItem*>(ptr)->setChild(row, column, static_cast<QStandardItem*>(item));
}

void QStandardItem_SetColumnCount(QtObjectPtr ptr, int columns){
	static_cast<QStandardItem*>(ptr)->setColumnCount(columns);
}

void QStandardItem_SetData(QtObjectPtr ptr, char* value, int role){
	static_cast<QStandardItem*>(ptr)->setData(QVariant(value), role);
}

void QStandardItem_SetDragEnabled(QtObjectPtr ptr, int dragEnabled){
	static_cast<QStandardItem*>(ptr)->setDragEnabled(dragEnabled != 0);
}

void QStandardItem_SetDropEnabled(QtObjectPtr ptr, int dropEnabled){
	static_cast<QStandardItem*>(ptr)->setDropEnabled(dropEnabled != 0);
}

void QStandardItem_SetEditable(QtObjectPtr ptr, int editable){
	static_cast<QStandardItem*>(ptr)->setEditable(editable != 0);
}

void QStandardItem_SetEnabled(QtObjectPtr ptr, int enabled){
	static_cast<QStandardItem*>(ptr)->setEnabled(enabled != 0);
}

void QStandardItem_SetFlags(QtObjectPtr ptr, int flags){
	static_cast<QStandardItem*>(ptr)->setFlags(static_cast<Qt::ItemFlag>(flags));
}

void QStandardItem_SetFont(QtObjectPtr ptr, QtObjectPtr font){
	static_cast<QStandardItem*>(ptr)->setFont(*static_cast<QFont*>(font));
}

void QStandardItem_SetForeground(QtObjectPtr ptr, QtObjectPtr brush){
	static_cast<QStandardItem*>(ptr)->setForeground(*static_cast<QBrush*>(brush));
}

void QStandardItem_SetIcon(QtObjectPtr ptr, QtObjectPtr icon){
	static_cast<QStandardItem*>(ptr)->setIcon(*static_cast<QIcon*>(icon));
}

void QStandardItem_SetRowCount(QtObjectPtr ptr, int rows){
	static_cast<QStandardItem*>(ptr)->setRowCount(rows);
}

void QStandardItem_SetSelectable(QtObjectPtr ptr, int selectable){
	static_cast<QStandardItem*>(ptr)->setSelectable(selectable != 0);
}

void QStandardItem_SetSizeHint(QtObjectPtr ptr, QtObjectPtr size){
	static_cast<QStandardItem*>(ptr)->setSizeHint(*static_cast<QSize*>(size));
}

void QStandardItem_SetStatusTip(QtObjectPtr ptr, char* statusTip){
	static_cast<QStandardItem*>(ptr)->setStatusTip(QString(statusTip));
}

void QStandardItem_SetText(QtObjectPtr ptr, char* text){
	static_cast<QStandardItem*>(ptr)->setText(QString(text));
}

void QStandardItem_SetTextAlignment(QtObjectPtr ptr, int alignment){
	static_cast<QStandardItem*>(ptr)->setTextAlignment(static_cast<Qt::AlignmentFlag>(alignment));
}

void QStandardItem_SetToolTip(QtObjectPtr ptr, char* toolTip){
	static_cast<QStandardItem*>(ptr)->setToolTip(QString(toolTip));
}

void QStandardItem_SetTristate(QtObjectPtr ptr, int tristate){
	static_cast<QStandardItem*>(ptr)->setTristate(tristate != 0);
}

void QStandardItem_SetWhatsThis(QtObjectPtr ptr, char* whatsThis){
	static_cast<QStandardItem*>(ptr)->setWhatsThis(QString(whatsThis));
}

void QStandardItem_SortChildren(QtObjectPtr ptr, int column, int order){
	static_cast<QStandardItem*>(ptr)->sortChildren(column, static_cast<Qt::SortOrder>(order));
}

char* QStandardItem_StatusTip(QtObjectPtr ptr){
	return static_cast<QStandardItem*>(ptr)->statusTip().toUtf8().data();
}

QtObjectPtr QStandardItem_TakeChild(QtObjectPtr ptr, int row, int column){
	return static_cast<QStandardItem*>(ptr)->takeChild(row, column);
}

char* QStandardItem_Text(QtObjectPtr ptr){
	return static_cast<QStandardItem*>(ptr)->text().toUtf8().data();
}

int QStandardItem_TextAlignment(QtObjectPtr ptr){
	return static_cast<QStandardItem*>(ptr)->textAlignment();
}

char* QStandardItem_ToolTip(QtObjectPtr ptr){
	return static_cast<QStandardItem*>(ptr)->toolTip().toUtf8().data();
}

int QStandardItem_Type(QtObjectPtr ptr){
	return static_cast<QStandardItem*>(ptr)->type();
}

char* QStandardItem_WhatsThis(QtObjectPtr ptr){
	return static_cast<QStandardItem*>(ptr)->whatsThis().toUtf8().data();
}

void QStandardItem_Write(QtObjectPtr ptr, QtObjectPtr out){
	static_cast<QStandardItem*>(ptr)->write(*static_cast<QDataStream*>(out));
}

void QStandardItem_DestroyQStandardItem(QtObjectPtr ptr){
	static_cast<QStandardItem*>(ptr)->~QStandardItem();
}

