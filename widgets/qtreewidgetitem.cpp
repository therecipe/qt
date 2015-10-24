#include "qtreewidgetitem.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QSize>
#include <QIcon>
#include <QTreeWidget>
#include <QModelIndex>
#include <QBrush>
#include <QDataStream>
#include <QFont>
#include <QTreeWidgetItem>
#include "_cgo_export.h"

class MyQTreeWidgetItem: public QTreeWidgetItem {
public:
};

QtObjectPtr QTreeWidgetItem_NewQTreeWidgetItem5(QtObjectPtr parent, QtObjectPtr preceding, int ty){
	return new QTreeWidgetItem(static_cast<QTreeWidget*>(parent), static_cast<QTreeWidgetItem*>(preceding), ty);
}

QtObjectPtr QTreeWidgetItem_NewQTreeWidgetItem4(QtObjectPtr parent, char* strin, int ty){
	return new QTreeWidgetItem(static_cast<QTreeWidget*>(parent), QString(strin).split("|", QString::SkipEmptyParts), ty);
}

QtObjectPtr QTreeWidgetItem_NewQTreeWidgetItem3(QtObjectPtr parent, int ty){
	return new QTreeWidgetItem(static_cast<QTreeWidget*>(parent), ty);
}

QtObjectPtr QTreeWidgetItem_NewQTreeWidgetItem8(QtObjectPtr parent, QtObjectPtr preceding, int ty){
	return new QTreeWidgetItem(static_cast<QTreeWidgetItem*>(parent), static_cast<QTreeWidgetItem*>(preceding), ty);
}

int QTreeWidgetItem_Flags(QtObjectPtr ptr){
	return static_cast<QTreeWidgetItem*>(ptr)->flags();
}

void QTreeWidgetItem_SetFlags(QtObjectPtr ptr, int flags){
	static_cast<QTreeWidgetItem*>(ptr)->setFlags(static_cast<Qt::ItemFlag>(flags));
}

QtObjectPtr QTreeWidgetItem_NewQTreeWidgetItem7(QtObjectPtr parent, char* strin, int ty){
	return new QTreeWidgetItem(static_cast<QTreeWidgetItem*>(parent), QString(strin).split("|", QString::SkipEmptyParts), ty);
}

QtObjectPtr QTreeWidgetItem_NewQTreeWidgetItem6(QtObjectPtr parent, int ty){
	return new QTreeWidgetItem(static_cast<QTreeWidgetItem*>(parent), ty);
}

QtObjectPtr QTreeWidgetItem_NewQTreeWidgetItem2(char* strin, int ty){
	return new QTreeWidgetItem(QString(strin).split("|", QString::SkipEmptyParts), ty);
}

QtObjectPtr QTreeWidgetItem_NewQTreeWidgetItem9(QtObjectPtr other){
	return new QTreeWidgetItem(*static_cast<QTreeWidgetItem*>(other));
}

QtObjectPtr QTreeWidgetItem_NewQTreeWidgetItem(int ty){
	return new QTreeWidgetItem(ty);
}

void QTreeWidgetItem_AddChild(QtObjectPtr ptr, QtObjectPtr child){
	static_cast<QTreeWidgetItem*>(ptr)->addChild(static_cast<QTreeWidgetItem*>(child));
}

int QTreeWidgetItem_CheckState(QtObjectPtr ptr, int column){
	return static_cast<QTreeWidgetItem*>(ptr)->checkState(column);
}

QtObjectPtr QTreeWidgetItem_Child(QtObjectPtr ptr, int index){
	return static_cast<QTreeWidgetItem*>(ptr)->child(index);
}

int QTreeWidgetItem_ChildCount(QtObjectPtr ptr){
	return static_cast<QTreeWidgetItem*>(ptr)->childCount();
}

int QTreeWidgetItem_ChildIndicatorPolicy(QtObjectPtr ptr){
	return static_cast<QTreeWidgetItem*>(ptr)->childIndicatorPolicy();
}

int QTreeWidgetItem_ColumnCount(QtObjectPtr ptr){
	return static_cast<QTreeWidgetItem*>(ptr)->columnCount();
}

char* QTreeWidgetItem_Data(QtObjectPtr ptr, int column, int role){
	return static_cast<QTreeWidgetItem*>(ptr)->data(column, role).toString().toUtf8().data();
}

QtObjectPtr QTreeWidgetItem_Clone(QtObjectPtr ptr){
	return static_cast<QTreeWidgetItem*>(ptr)->clone();
}

int QTreeWidgetItem_IndexOfChild(QtObjectPtr ptr, QtObjectPtr child){
	return static_cast<QTreeWidgetItem*>(ptr)->indexOfChild(static_cast<QTreeWidgetItem*>(child));
}

void QTreeWidgetItem_InsertChild(QtObjectPtr ptr, int index, QtObjectPtr child){
	static_cast<QTreeWidgetItem*>(ptr)->insertChild(index, static_cast<QTreeWidgetItem*>(child));
}

int QTreeWidgetItem_IsDisabled(QtObjectPtr ptr){
	return static_cast<QTreeWidgetItem*>(ptr)->isDisabled();
}

int QTreeWidgetItem_IsExpanded(QtObjectPtr ptr){
	return static_cast<QTreeWidgetItem*>(ptr)->isExpanded();
}

int QTreeWidgetItem_IsFirstColumnSpanned(QtObjectPtr ptr){
	return static_cast<QTreeWidgetItem*>(ptr)->isFirstColumnSpanned();
}

int QTreeWidgetItem_IsHidden(QtObjectPtr ptr){
	return static_cast<QTreeWidgetItem*>(ptr)->isHidden();
}

int QTreeWidgetItem_IsSelected(QtObjectPtr ptr){
	return static_cast<QTreeWidgetItem*>(ptr)->isSelected();
}

QtObjectPtr QTreeWidgetItem_Parent(QtObjectPtr ptr){
	return static_cast<QTreeWidgetItem*>(ptr)->parent();
}

void QTreeWidgetItem_Read(QtObjectPtr ptr, QtObjectPtr in){
	static_cast<QTreeWidgetItem*>(ptr)->read(*static_cast<QDataStream*>(in));
}

void QTreeWidgetItem_RemoveChild(QtObjectPtr ptr, QtObjectPtr child){
	static_cast<QTreeWidgetItem*>(ptr)->removeChild(static_cast<QTreeWidgetItem*>(child));
}

void QTreeWidgetItem_SetBackground(QtObjectPtr ptr, int column, QtObjectPtr brush){
	static_cast<QTreeWidgetItem*>(ptr)->setBackground(column, *static_cast<QBrush*>(brush));
}

void QTreeWidgetItem_SetCheckState(QtObjectPtr ptr, int column, int state){
	static_cast<QTreeWidgetItem*>(ptr)->setCheckState(column, static_cast<Qt::CheckState>(state));
}

void QTreeWidgetItem_SetChildIndicatorPolicy(QtObjectPtr ptr, int policy){
	static_cast<QTreeWidgetItem*>(ptr)->setChildIndicatorPolicy(static_cast<QTreeWidgetItem::ChildIndicatorPolicy>(policy));
}

void QTreeWidgetItem_SetData(QtObjectPtr ptr, int column, int role, char* value){
	static_cast<QTreeWidgetItem*>(ptr)->setData(column, role, QVariant(value));
}

void QTreeWidgetItem_SetDisabled(QtObjectPtr ptr, int disabled){
	static_cast<QTreeWidgetItem*>(ptr)->setDisabled(disabled != 0);
}

void QTreeWidgetItem_SetExpanded(QtObjectPtr ptr, int expand){
	static_cast<QTreeWidgetItem*>(ptr)->setExpanded(expand != 0);
}

void QTreeWidgetItem_SetFirstColumnSpanned(QtObjectPtr ptr, int span){
	static_cast<QTreeWidgetItem*>(ptr)->setFirstColumnSpanned(span != 0);
}

void QTreeWidgetItem_SetFont(QtObjectPtr ptr, int column, QtObjectPtr font){
	static_cast<QTreeWidgetItem*>(ptr)->setFont(column, *static_cast<QFont*>(font));
}

void QTreeWidgetItem_SetForeground(QtObjectPtr ptr, int column, QtObjectPtr brush){
	static_cast<QTreeWidgetItem*>(ptr)->setForeground(column, *static_cast<QBrush*>(brush));
}

void QTreeWidgetItem_SetHidden(QtObjectPtr ptr, int hide){
	static_cast<QTreeWidgetItem*>(ptr)->setHidden(hide != 0);
}

void QTreeWidgetItem_SetIcon(QtObjectPtr ptr, int column, QtObjectPtr icon){
	static_cast<QTreeWidgetItem*>(ptr)->setIcon(column, *static_cast<QIcon*>(icon));
}

void QTreeWidgetItem_SetSelected(QtObjectPtr ptr, int sele){
	static_cast<QTreeWidgetItem*>(ptr)->setSelected(sele != 0);
}

void QTreeWidgetItem_SetSizeHint(QtObjectPtr ptr, int column, QtObjectPtr size){
	static_cast<QTreeWidgetItem*>(ptr)->setSizeHint(column, *static_cast<QSize*>(size));
}

void QTreeWidgetItem_SetStatusTip(QtObjectPtr ptr, int column, char* statusTip){
	static_cast<QTreeWidgetItem*>(ptr)->setStatusTip(column, QString(statusTip));
}

void QTreeWidgetItem_SetText(QtObjectPtr ptr, int column, char* text){
	static_cast<QTreeWidgetItem*>(ptr)->setText(column, QString(text));
}

void QTreeWidgetItem_SetTextAlignment(QtObjectPtr ptr, int column, int alignment){
	static_cast<QTreeWidgetItem*>(ptr)->setTextAlignment(column, alignment);
}

void QTreeWidgetItem_SetToolTip(QtObjectPtr ptr, int column, char* toolTip){
	static_cast<QTreeWidgetItem*>(ptr)->setToolTip(column, QString(toolTip));
}

void QTreeWidgetItem_SetWhatsThis(QtObjectPtr ptr, int column, char* whatsThis){
	static_cast<QTreeWidgetItem*>(ptr)->setWhatsThis(column, QString(whatsThis));
}

void QTreeWidgetItem_SortChildren(QtObjectPtr ptr, int column, int order){
	static_cast<QTreeWidgetItem*>(ptr)->sortChildren(column, static_cast<Qt::SortOrder>(order));
}

char* QTreeWidgetItem_StatusTip(QtObjectPtr ptr, int column){
	return static_cast<QTreeWidgetItem*>(ptr)->statusTip(column).toUtf8().data();
}

QtObjectPtr QTreeWidgetItem_TakeChild(QtObjectPtr ptr, int index){
	return static_cast<QTreeWidgetItem*>(ptr)->takeChild(index);
}

char* QTreeWidgetItem_Text(QtObjectPtr ptr, int column){
	return static_cast<QTreeWidgetItem*>(ptr)->text(column).toUtf8().data();
}

int QTreeWidgetItem_TextAlignment(QtObjectPtr ptr, int column){
	return static_cast<QTreeWidgetItem*>(ptr)->textAlignment(column);
}

char* QTreeWidgetItem_ToolTip(QtObjectPtr ptr, int column){
	return static_cast<QTreeWidgetItem*>(ptr)->toolTip(column).toUtf8().data();
}

QtObjectPtr QTreeWidgetItem_TreeWidget(QtObjectPtr ptr){
	return static_cast<QTreeWidgetItem*>(ptr)->treeWidget();
}

int QTreeWidgetItem_Type(QtObjectPtr ptr){
	return static_cast<QTreeWidgetItem*>(ptr)->type();
}

char* QTreeWidgetItem_WhatsThis(QtObjectPtr ptr, int column){
	return static_cast<QTreeWidgetItem*>(ptr)->whatsThis(column).toUtf8().data();
}

void QTreeWidgetItem_Write(QtObjectPtr ptr, QtObjectPtr out){
	static_cast<QTreeWidgetItem*>(ptr)->write(*static_cast<QDataStream*>(out));
}

void QTreeWidgetItem_DestroyQTreeWidgetItem(QtObjectPtr ptr){
	static_cast<QTreeWidgetItem*>(ptr)->~QTreeWidgetItem();
}

