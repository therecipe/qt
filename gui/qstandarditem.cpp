#include "qstandarditem.h"
#include <QUrl>
#include <QIcon>
#include <QFont>
#include <QBrush>
#include <QString>
#include <QVariant>
#include <QModelIndex>
#include <QSize>
#include <QDataStream>
#include <QStandardItem>
#include "_cgo_export.h"

class MyQStandardItem: public QStandardItem {
public:
};

void* QStandardItem_NewQStandardItem(){
	return new QStandardItem();
}

void* QStandardItem_NewQStandardItem3(void* icon, char* text){
	return new QStandardItem(*static_cast<QIcon*>(icon), QString(text));
}

void* QStandardItem_NewQStandardItem2(char* text){
	return new QStandardItem(QString(text));
}

void* QStandardItem_NewQStandardItem4(int rows, int columns){
	return new QStandardItem(rows, columns);
}

char* QStandardItem_AccessibleDescription(void* ptr){
	return static_cast<QStandardItem*>(ptr)->accessibleDescription().toUtf8().data();
}

char* QStandardItem_AccessibleText(void* ptr){
	return static_cast<QStandardItem*>(ptr)->accessibleText().toUtf8().data();
}

void QStandardItem_AppendRow2(void* ptr, void* item){
	static_cast<QStandardItem*>(ptr)->appendRow(static_cast<QStandardItem*>(item));
}

void* QStandardItem_Background(void* ptr){
	return new QBrush(static_cast<QStandardItem*>(ptr)->background());
}

int QStandardItem_CheckState(void* ptr){
	return static_cast<QStandardItem*>(ptr)->checkState();
}

void* QStandardItem_Child(void* ptr, int row, int column){
	return static_cast<QStandardItem*>(ptr)->child(row, column);
}

void* QStandardItem_Clone(void* ptr){
	return static_cast<QStandardItem*>(ptr)->clone();
}

int QStandardItem_Column(void* ptr){
	return static_cast<QStandardItem*>(ptr)->column();
}

int QStandardItem_ColumnCount(void* ptr){
	return static_cast<QStandardItem*>(ptr)->columnCount();
}

void* QStandardItem_Data(void* ptr, int role){
	return new QVariant(static_cast<QStandardItem*>(ptr)->data(role));
}

int QStandardItem_Flags(void* ptr){
	return static_cast<QStandardItem*>(ptr)->flags();
}

void* QStandardItem_Foreground(void* ptr){
	return new QBrush(static_cast<QStandardItem*>(ptr)->foreground());
}

int QStandardItem_HasChildren(void* ptr){
	return static_cast<QStandardItem*>(ptr)->hasChildren();
}

void* QStandardItem_Index(void* ptr){
	return static_cast<QStandardItem*>(ptr)->index().internalPointer();
}

void QStandardItem_InsertColumns(void* ptr, int column, int count){
	static_cast<QStandardItem*>(ptr)->insertColumns(column, count);
}

void QStandardItem_InsertRow2(void* ptr, int row, void* item){
	static_cast<QStandardItem*>(ptr)->insertRow(row, static_cast<QStandardItem*>(item));
}

void QStandardItem_InsertRows2(void* ptr, int row, int count){
	static_cast<QStandardItem*>(ptr)->insertRows(row, count);
}

int QStandardItem_IsCheckable(void* ptr){
	return static_cast<QStandardItem*>(ptr)->isCheckable();
}

int QStandardItem_IsDragEnabled(void* ptr){
	return static_cast<QStandardItem*>(ptr)->isDragEnabled();
}

int QStandardItem_IsDropEnabled(void* ptr){
	return static_cast<QStandardItem*>(ptr)->isDropEnabled();
}

int QStandardItem_IsEditable(void* ptr){
	return static_cast<QStandardItem*>(ptr)->isEditable();
}

int QStandardItem_IsEnabled(void* ptr){
	return static_cast<QStandardItem*>(ptr)->isEnabled();
}

int QStandardItem_IsSelectable(void* ptr){
	return static_cast<QStandardItem*>(ptr)->isSelectable();
}

int QStandardItem_IsTristate(void* ptr){
	return static_cast<QStandardItem*>(ptr)->isTristate();
}

void* QStandardItem_Model(void* ptr){
	return static_cast<QStandardItem*>(ptr)->model();
}

void* QStandardItem_Parent(void* ptr){
	return static_cast<QStandardItem*>(ptr)->parent();
}

void QStandardItem_Read(void* ptr, void* in){
	static_cast<QStandardItem*>(ptr)->read(*static_cast<QDataStream*>(in));
}

void QStandardItem_RemoveColumn(void* ptr, int column){
	static_cast<QStandardItem*>(ptr)->removeColumn(column);
}

void QStandardItem_RemoveColumns(void* ptr, int column, int count){
	static_cast<QStandardItem*>(ptr)->removeColumns(column, count);
}

void QStandardItem_RemoveRow(void* ptr, int row){
	static_cast<QStandardItem*>(ptr)->removeRow(row);
}

void QStandardItem_RemoveRows(void* ptr, int row, int count){
	static_cast<QStandardItem*>(ptr)->removeRows(row, count);
}

int QStandardItem_Row(void* ptr){
	return static_cast<QStandardItem*>(ptr)->row();
}

int QStandardItem_RowCount(void* ptr){
	return static_cast<QStandardItem*>(ptr)->rowCount();
}

void QStandardItem_SetAccessibleDescription(void* ptr, char* accessibleDescription){
	static_cast<QStandardItem*>(ptr)->setAccessibleDescription(QString(accessibleDescription));
}

void QStandardItem_SetAccessibleText(void* ptr, char* accessibleText){
	static_cast<QStandardItem*>(ptr)->setAccessibleText(QString(accessibleText));
}

void QStandardItem_SetBackground(void* ptr, void* brush){
	static_cast<QStandardItem*>(ptr)->setBackground(*static_cast<QBrush*>(brush));
}

void QStandardItem_SetCheckState(void* ptr, int state){
	static_cast<QStandardItem*>(ptr)->setCheckState(static_cast<Qt::CheckState>(state));
}

void QStandardItem_SetCheckable(void* ptr, int checkable){
	static_cast<QStandardItem*>(ptr)->setCheckable(checkable != 0);
}

void QStandardItem_SetChild2(void* ptr, int row, void* item){
	static_cast<QStandardItem*>(ptr)->setChild(row, static_cast<QStandardItem*>(item));
}

void QStandardItem_SetChild(void* ptr, int row, int column, void* item){
	static_cast<QStandardItem*>(ptr)->setChild(row, column, static_cast<QStandardItem*>(item));
}

void QStandardItem_SetColumnCount(void* ptr, int columns){
	static_cast<QStandardItem*>(ptr)->setColumnCount(columns);
}

void QStandardItem_SetData(void* ptr, void* value, int role){
	static_cast<QStandardItem*>(ptr)->setData(*static_cast<QVariant*>(value), role);
}

void QStandardItem_SetDragEnabled(void* ptr, int dragEnabled){
	static_cast<QStandardItem*>(ptr)->setDragEnabled(dragEnabled != 0);
}

void QStandardItem_SetDropEnabled(void* ptr, int dropEnabled){
	static_cast<QStandardItem*>(ptr)->setDropEnabled(dropEnabled != 0);
}

void QStandardItem_SetEditable(void* ptr, int editable){
	static_cast<QStandardItem*>(ptr)->setEditable(editable != 0);
}

void QStandardItem_SetEnabled(void* ptr, int enabled){
	static_cast<QStandardItem*>(ptr)->setEnabled(enabled != 0);
}

void QStandardItem_SetFlags(void* ptr, int flags){
	static_cast<QStandardItem*>(ptr)->setFlags(static_cast<Qt::ItemFlag>(flags));
}

void QStandardItem_SetFont(void* ptr, void* font){
	static_cast<QStandardItem*>(ptr)->setFont(*static_cast<QFont*>(font));
}

void QStandardItem_SetForeground(void* ptr, void* brush){
	static_cast<QStandardItem*>(ptr)->setForeground(*static_cast<QBrush*>(brush));
}

void QStandardItem_SetIcon(void* ptr, void* icon){
	static_cast<QStandardItem*>(ptr)->setIcon(*static_cast<QIcon*>(icon));
}

void QStandardItem_SetRowCount(void* ptr, int rows){
	static_cast<QStandardItem*>(ptr)->setRowCount(rows);
}

void QStandardItem_SetSelectable(void* ptr, int selectable){
	static_cast<QStandardItem*>(ptr)->setSelectable(selectable != 0);
}

void QStandardItem_SetSizeHint(void* ptr, void* size){
	static_cast<QStandardItem*>(ptr)->setSizeHint(*static_cast<QSize*>(size));
}

void QStandardItem_SetStatusTip(void* ptr, char* statusTip){
	static_cast<QStandardItem*>(ptr)->setStatusTip(QString(statusTip));
}

void QStandardItem_SetText(void* ptr, char* text){
	static_cast<QStandardItem*>(ptr)->setText(QString(text));
}

void QStandardItem_SetTextAlignment(void* ptr, int alignment){
	static_cast<QStandardItem*>(ptr)->setTextAlignment(static_cast<Qt::AlignmentFlag>(alignment));
}

void QStandardItem_SetToolTip(void* ptr, char* toolTip){
	static_cast<QStandardItem*>(ptr)->setToolTip(QString(toolTip));
}

void QStandardItem_SetTristate(void* ptr, int tristate){
	static_cast<QStandardItem*>(ptr)->setTristate(tristate != 0);
}

void QStandardItem_SetWhatsThis(void* ptr, char* whatsThis){
	static_cast<QStandardItem*>(ptr)->setWhatsThis(QString(whatsThis));
}

void QStandardItem_SortChildren(void* ptr, int column, int order){
	static_cast<QStandardItem*>(ptr)->sortChildren(column, static_cast<Qt::SortOrder>(order));
}

char* QStandardItem_StatusTip(void* ptr){
	return static_cast<QStandardItem*>(ptr)->statusTip().toUtf8().data();
}

void* QStandardItem_TakeChild(void* ptr, int row, int column){
	return static_cast<QStandardItem*>(ptr)->takeChild(row, column);
}

char* QStandardItem_Text(void* ptr){
	return static_cast<QStandardItem*>(ptr)->text().toUtf8().data();
}

int QStandardItem_TextAlignment(void* ptr){
	return static_cast<QStandardItem*>(ptr)->textAlignment();
}

char* QStandardItem_ToolTip(void* ptr){
	return static_cast<QStandardItem*>(ptr)->toolTip().toUtf8().data();
}

int QStandardItem_Type(void* ptr){
	return static_cast<QStandardItem*>(ptr)->type();
}

char* QStandardItem_WhatsThis(void* ptr){
	return static_cast<QStandardItem*>(ptr)->whatsThis().toUtf8().data();
}

void QStandardItem_Write(void* ptr, void* out){
	static_cast<QStandardItem*>(ptr)->write(*static_cast<QDataStream*>(out));
}

void QStandardItem_DestroyQStandardItem(void* ptr){
	static_cast<QStandardItem*>(ptr)->~QStandardItem();
}

