#include "qlistwidgetitem.h"
#include <QList>
#include <QUrl>
#include <QDataStream>
#include <QBrush>
#include <QSize>
#include <QIcon>
#include <QListWidget>
#include <QString>
#include <QVariant>
#include <QModelIndex>
#include <QList>
#include <QFont>
#include <QListWidgetItem>
#include "_cgo_export.h"

class MyQListWidgetItem: public QListWidgetItem {
public:
};

void* QListWidgetItem_NewQListWidgetItem(void* parent, int ty){
	return new QListWidgetItem(static_cast<QListWidget*>(parent), ty);
}

void* QListWidgetItem_NewQListWidgetItem3(void* icon, char* text, void* parent, int ty){
	return new QListWidgetItem(*static_cast<QIcon*>(icon), QString(text), static_cast<QListWidget*>(parent), ty);
}

void* QListWidgetItem_NewQListWidgetItem2(char* text, void* parent, int ty){
	return new QListWidgetItem(QString(text), static_cast<QListWidget*>(parent), ty);
}

void QListWidgetItem_SetFlags(void* ptr, int flags){
	static_cast<QListWidgetItem*>(ptr)->setFlags(static_cast<Qt::ItemFlag>(flags));
}

void* QListWidgetItem_NewQListWidgetItem4(void* other){
	return new QListWidgetItem(*static_cast<QListWidgetItem*>(other));
}

void* QListWidgetItem_Background(void* ptr){
	return new QBrush(static_cast<QListWidgetItem*>(ptr)->background());
}

int QListWidgetItem_CheckState(void* ptr){
	return static_cast<QListWidgetItem*>(ptr)->checkState();
}

void* QListWidgetItem_Clone(void* ptr){
	return static_cast<QListWidgetItem*>(ptr)->clone();
}

void* QListWidgetItem_Data(void* ptr, int role){
	return new QVariant(static_cast<QListWidgetItem*>(ptr)->data(role));
}

int QListWidgetItem_Flags(void* ptr){
	return static_cast<QListWidgetItem*>(ptr)->flags();
}

void* QListWidgetItem_Foreground(void* ptr){
	return new QBrush(static_cast<QListWidgetItem*>(ptr)->foreground());
}

int QListWidgetItem_IsHidden(void* ptr){
	return static_cast<QListWidgetItem*>(ptr)->isHidden();
}

int QListWidgetItem_IsSelected(void* ptr){
	return static_cast<QListWidgetItem*>(ptr)->isSelected();
}

void* QListWidgetItem_ListWidget(void* ptr){
	return static_cast<QListWidgetItem*>(ptr)->listWidget();
}

void QListWidgetItem_Read(void* ptr, void* in){
	static_cast<QListWidgetItem*>(ptr)->read(*static_cast<QDataStream*>(in));
}

void QListWidgetItem_SetBackground(void* ptr, void* brush){
	static_cast<QListWidgetItem*>(ptr)->setBackground(*static_cast<QBrush*>(brush));
}

void QListWidgetItem_SetCheckState(void* ptr, int state){
	static_cast<QListWidgetItem*>(ptr)->setCheckState(static_cast<Qt::CheckState>(state));
}

void QListWidgetItem_SetData(void* ptr, int role, void* value){
	static_cast<QListWidgetItem*>(ptr)->setData(role, *static_cast<QVariant*>(value));
}

void QListWidgetItem_SetFont(void* ptr, void* font){
	static_cast<QListWidgetItem*>(ptr)->setFont(*static_cast<QFont*>(font));
}

void QListWidgetItem_SetForeground(void* ptr, void* brush){
	static_cast<QListWidgetItem*>(ptr)->setForeground(*static_cast<QBrush*>(brush));
}

void QListWidgetItem_SetHidden(void* ptr, int hide){
	static_cast<QListWidgetItem*>(ptr)->setHidden(hide != 0);
}

void QListWidgetItem_SetIcon(void* ptr, void* icon){
	static_cast<QListWidgetItem*>(ptr)->setIcon(*static_cast<QIcon*>(icon));
}

void QListWidgetItem_SetSelected(void* ptr, int sele){
	static_cast<QListWidgetItem*>(ptr)->setSelected(sele != 0);
}

void QListWidgetItem_SetSizeHint(void* ptr, void* size){
	static_cast<QListWidgetItem*>(ptr)->setSizeHint(*static_cast<QSize*>(size));
}

void QListWidgetItem_SetStatusTip(void* ptr, char* statusTip){
	static_cast<QListWidgetItem*>(ptr)->setStatusTip(QString(statusTip));
}

void QListWidgetItem_SetText(void* ptr, char* text){
	static_cast<QListWidgetItem*>(ptr)->setText(QString(text));
}

void QListWidgetItem_SetTextAlignment(void* ptr, int alignment){
	static_cast<QListWidgetItem*>(ptr)->setTextAlignment(alignment);
}

void QListWidgetItem_SetToolTip(void* ptr, char* toolTip){
	static_cast<QListWidgetItem*>(ptr)->setToolTip(QString(toolTip));
}

void QListWidgetItem_SetWhatsThis(void* ptr, char* whatsThis){
	static_cast<QListWidgetItem*>(ptr)->setWhatsThis(QString(whatsThis));
}

char* QListWidgetItem_StatusTip(void* ptr){
	return static_cast<QListWidgetItem*>(ptr)->statusTip().toUtf8().data();
}

char* QListWidgetItem_Text(void* ptr){
	return static_cast<QListWidgetItem*>(ptr)->text().toUtf8().data();
}

int QListWidgetItem_TextAlignment(void* ptr){
	return static_cast<QListWidgetItem*>(ptr)->textAlignment();
}

char* QListWidgetItem_ToolTip(void* ptr){
	return static_cast<QListWidgetItem*>(ptr)->toolTip().toUtf8().data();
}

int QListWidgetItem_Type(void* ptr){
	return static_cast<QListWidgetItem*>(ptr)->type();
}

char* QListWidgetItem_WhatsThis(void* ptr){
	return static_cast<QListWidgetItem*>(ptr)->whatsThis().toUtf8().data();
}

void QListWidgetItem_Write(void* ptr, void* out){
	static_cast<QListWidgetItem*>(ptr)->write(*static_cast<QDataStream*>(out));
}

void QListWidgetItem_DestroyQListWidgetItem(void* ptr){
	static_cast<QListWidgetItem*>(ptr)->~QListWidgetItem();
}

