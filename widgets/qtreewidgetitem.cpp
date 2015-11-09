#include "qtreewidgetitem.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QTreeWidget>
#include <QIcon>
#include <QDataStream>
#include <QSize>
#include <QString>
#include <QFont>
#include <QBrush>
#include <QTreeWidgetItem>
#include "_cgo_export.h"

class MyQTreeWidgetItem: public QTreeWidgetItem {
public:
};

void* QTreeWidgetItem_NewQTreeWidgetItem5(void* parent, void* preceding, int ty){
	return new QTreeWidgetItem(static_cast<QTreeWidget*>(parent), static_cast<QTreeWidgetItem*>(preceding), ty);
}

void* QTreeWidgetItem_NewQTreeWidgetItem4(void* parent, char* strin, int ty){
	return new QTreeWidgetItem(static_cast<QTreeWidget*>(parent), QString(strin).split("|", QString::SkipEmptyParts), ty);
}

void* QTreeWidgetItem_NewQTreeWidgetItem3(void* parent, int ty){
	return new QTreeWidgetItem(static_cast<QTreeWidget*>(parent), ty);
}

void* QTreeWidgetItem_NewQTreeWidgetItem8(void* parent, void* preceding, int ty){
	return new QTreeWidgetItem(static_cast<QTreeWidgetItem*>(parent), static_cast<QTreeWidgetItem*>(preceding), ty);
}

int QTreeWidgetItem_Flags(void* ptr){
	return static_cast<QTreeWidgetItem*>(ptr)->flags();
}

void QTreeWidgetItem_SetFlags(void* ptr, int flags){
	static_cast<QTreeWidgetItem*>(ptr)->setFlags(static_cast<Qt::ItemFlag>(flags));
}

void* QTreeWidgetItem_NewQTreeWidgetItem7(void* parent, char* strin, int ty){
	return new QTreeWidgetItem(static_cast<QTreeWidgetItem*>(parent), QString(strin).split("|", QString::SkipEmptyParts), ty);
}

void* QTreeWidgetItem_NewQTreeWidgetItem6(void* parent, int ty){
	return new QTreeWidgetItem(static_cast<QTreeWidgetItem*>(parent), ty);
}

void* QTreeWidgetItem_NewQTreeWidgetItem2(char* strin, int ty){
	return new QTreeWidgetItem(QString(strin).split("|", QString::SkipEmptyParts), ty);
}

void* QTreeWidgetItem_NewQTreeWidgetItem9(void* other){
	return new QTreeWidgetItem(*static_cast<QTreeWidgetItem*>(other));
}

void* QTreeWidgetItem_NewQTreeWidgetItem(int ty){
	return new QTreeWidgetItem(ty);
}

void QTreeWidgetItem_AddChild(void* ptr, void* child){
	static_cast<QTreeWidgetItem*>(ptr)->addChild(static_cast<QTreeWidgetItem*>(child));
}

void* QTreeWidgetItem_Background(void* ptr, int column){
	return new QBrush(static_cast<QTreeWidgetItem*>(ptr)->background(column));
}

int QTreeWidgetItem_CheckState(void* ptr, int column){
	return static_cast<QTreeWidgetItem*>(ptr)->checkState(column);
}

void* QTreeWidgetItem_Child(void* ptr, int index){
	return static_cast<QTreeWidgetItem*>(ptr)->child(index);
}

int QTreeWidgetItem_ChildCount(void* ptr){
	return static_cast<QTreeWidgetItem*>(ptr)->childCount();
}

int QTreeWidgetItem_ChildIndicatorPolicy(void* ptr){
	return static_cast<QTreeWidgetItem*>(ptr)->childIndicatorPolicy();
}

int QTreeWidgetItem_ColumnCount(void* ptr){
	return static_cast<QTreeWidgetItem*>(ptr)->columnCount();
}

void* QTreeWidgetItem_Data(void* ptr, int column, int role){
	return new QVariant(static_cast<QTreeWidgetItem*>(ptr)->data(column, role));
}

void* QTreeWidgetItem_Clone(void* ptr){
	return static_cast<QTreeWidgetItem*>(ptr)->clone();
}

void* QTreeWidgetItem_Foreground(void* ptr, int column){
	return new QBrush(static_cast<QTreeWidgetItem*>(ptr)->foreground(column));
}

int QTreeWidgetItem_IndexOfChild(void* ptr, void* child){
	return static_cast<QTreeWidgetItem*>(ptr)->indexOfChild(static_cast<QTreeWidgetItem*>(child));
}

void QTreeWidgetItem_InsertChild(void* ptr, int index, void* child){
	static_cast<QTreeWidgetItem*>(ptr)->insertChild(index, static_cast<QTreeWidgetItem*>(child));
}

int QTreeWidgetItem_IsDisabled(void* ptr){
	return static_cast<QTreeWidgetItem*>(ptr)->isDisabled();
}

int QTreeWidgetItem_IsExpanded(void* ptr){
	return static_cast<QTreeWidgetItem*>(ptr)->isExpanded();
}

int QTreeWidgetItem_IsFirstColumnSpanned(void* ptr){
	return static_cast<QTreeWidgetItem*>(ptr)->isFirstColumnSpanned();
}

int QTreeWidgetItem_IsHidden(void* ptr){
	return static_cast<QTreeWidgetItem*>(ptr)->isHidden();
}

int QTreeWidgetItem_IsSelected(void* ptr){
	return static_cast<QTreeWidgetItem*>(ptr)->isSelected();
}

void* QTreeWidgetItem_Parent(void* ptr){
	return static_cast<QTreeWidgetItem*>(ptr)->parent();
}

void QTreeWidgetItem_Read(void* ptr, void* in){
	static_cast<QTreeWidgetItem*>(ptr)->read(*static_cast<QDataStream*>(in));
}

void QTreeWidgetItem_RemoveChild(void* ptr, void* child){
	static_cast<QTreeWidgetItem*>(ptr)->removeChild(static_cast<QTreeWidgetItem*>(child));
}

void QTreeWidgetItem_SetBackground(void* ptr, int column, void* brush){
	static_cast<QTreeWidgetItem*>(ptr)->setBackground(column, *static_cast<QBrush*>(brush));
}

void QTreeWidgetItem_SetCheckState(void* ptr, int column, int state){
	static_cast<QTreeWidgetItem*>(ptr)->setCheckState(column, static_cast<Qt::CheckState>(state));
}

void QTreeWidgetItem_SetChildIndicatorPolicy(void* ptr, int policy){
	static_cast<QTreeWidgetItem*>(ptr)->setChildIndicatorPolicy(static_cast<QTreeWidgetItem::ChildIndicatorPolicy>(policy));
}

void QTreeWidgetItem_SetData(void* ptr, int column, int role, void* value){
	static_cast<QTreeWidgetItem*>(ptr)->setData(column, role, *static_cast<QVariant*>(value));
}

void QTreeWidgetItem_SetDisabled(void* ptr, int disabled){
	static_cast<QTreeWidgetItem*>(ptr)->setDisabled(disabled != 0);
}

void QTreeWidgetItem_SetExpanded(void* ptr, int expand){
	static_cast<QTreeWidgetItem*>(ptr)->setExpanded(expand != 0);
}

void QTreeWidgetItem_SetFirstColumnSpanned(void* ptr, int span){
	static_cast<QTreeWidgetItem*>(ptr)->setFirstColumnSpanned(span != 0);
}

void QTreeWidgetItem_SetFont(void* ptr, int column, void* font){
	static_cast<QTreeWidgetItem*>(ptr)->setFont(column, *static_cast<QFont*>(font));
}

void QTreeWidgetItem_SetForeground(void* ptr, int column, void* brush){
	static_cast<QTreeWidgetItem*>(ptr)->setForeground(column, *static_cast<QBrush*>(brush));
}

void QTreeWidgetItem_SetHidden(void* ptr, int hide){
	static_cast<QTreeWidgetItem*>(ptr)->setHidden(hide != 0);
}

void QTreeWidgetItem_SetIcon(void* ptr, int column, void* icon){
	static_cast<QTreeWidgetItem*>(ptr)->setIcon(column, *static_cast<QIcon*>(icon));
}

void QTreeWidgetItem_SetSelected(void* ptr, int sele){
	static_cast<QTreeWidgetItem*>(ptr)->setSelected(sele != 0);
}

void QTreeWidgetItem_SetSizeHint(void* ptr, int column, void* size){
	static_cast<QTreeWidgetItem*>(ptr)->setSizeHint(column, *static_cast<QSize*>(size));
}

void QTreeWidgetItem_SetStatusTip(void* ptr, int column, char* statusTip){
	static_cast<QTreeWidgetItem*>(ptr)->setStatusTip(column, QString(statusTip));
}

void QTreeWidgetItem_SetText(void* ptr, int column, char* text){
	static_cast<QTreeWidgetItem*>(ptr)->setText(column, QString(text));
}

void QTreeWidgetItem_SetTextAlignment(void* ptr, int column, int alignment){
	static_cast<QTreeWidgetItem*>(ptr)->setTextAlignment(column, alignment);
}

void QTreeWidgetItem_SetToolTip(void* ptr, int column, char* toolTip){
	static_cast<QTreeWidgetItem*>(ptr)->setToolTip(column, QString(toolTip));
}

void QTreeWidgetItem_SetWhatsThis(void* ptr, int column, char* whatsThis){
	static_cast<QTreeWidgetItem*>(ptr)->setWhatsThis(column, QString(whatsThis));
}

void QTreeWidgetItem_SortChildren(void* ptr, int column, int order){
	static_cast<QTreeWidgetItem*>(ptr)->sortChildren(column, static_cast<Qt::SortOrder>(order));
}

char* QTreeWidgetItem_StatusTip(void* ptr, int column){
	return static_cast<QTreeWidgetItem*>(ptr)->statusTip(column).toUtf8().data();
}

void* QTreeWidgetItem_TakeChild(void* ptr, int index){
	return static_cast<QTreeWidgetItem*>(ptr)->takeChild(index);
}

char* QTreeWidgetItem_Text(void* ptr, int column){
	return static_cast<QTreeWidgetItem*>(ptr)->text(column).toUtf8().data();
}

int QTreeWidgetItem_TextAlignment(void* ptr, int column){
	return static_cast<QTreeWidgetItem*>(ptr)->textAlignment(column);
}

char* QTreeWidgetItem_ToolTip(void* ptr, int column){
	return static_cast<QTreeWidgetItem*>(ptr)->toolTip(column).toUtf8().data();
}

void* QTreeWidgetItem_TreeWidget(void* ptr){
	return static_cast<QTreeWidgetItem*>(ptr)->treeWidget();
}

int QTreeWidgetItem_Type(void* ptr){
	return static_cast<QTreeWidgetItem*>(ptr)->type();
}

char* QTreeWidgetItem_WhatsThis(void* ptr, int column){
	return static_cast<QTreeWidgetItem*>(ptr)->whatsThis(column).toUtf8().data();
}

void QTreeWidgetItem_Write(void* ptr, void* out){
	static_cast<QTreeWidgetItem*>(ptr)->write(*static_cast<QDataStream*>(out));
}

void QTreeWidgetItem_DestroyQTreeWidgetItem(void* ptr){
	static_cast<QTreeWidgetItem*>(ptr)->~QTreeWidgetItem();
}

