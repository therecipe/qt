#include "qlistwidgetitem.h"
#include <QList>
#include <QModelIndex>
#include <QList>
#include <QFont>
#include <QBrush>
#include <QVariant>
#include <QUrl>
#include <QIcon>
#include <QListWidget>
#include <QDataStream>
#include <QSize>
#include <QString>
#include <QListWidgetItem>
#include "_cgo_export.h"

class MyQListWidgetItem: public QListWidgetItem {
public:
};

QtObjectPtr QListWidgetItem_NewQListWidgetItem(QtObjectPtr parent, int ty){
	return new QListWidgetItem(static_cast<QListWidget*>(parent), ty);
}

QtObjectPtr QListWidgetItem_NewQListWidgetItem3(QtObjectPtr icon, char* text, QtObjectPtr parent, int ty){
	return new QListWidgetItem(*static_cast<QIcon*>(icon), QString(text), static_cast<QListWidget*>(parent), ty);
}

QtObjectPtr QListWidgetItem_NewQListWidgetItem2(char* text, QtObjectPtr parent, int ty){
	return new QListWidgetItem(QString(text), static_cast<QListWidget*>(parent), ty);
}

void QListWidgetItem_SetFlags(QtObjectPtr ptr, int flags){
	static_cast<QListWidgetItem*>(ptr)->setFlags(static_cast<Qt::ItemFlag>(flags));
}

QtObjectPtr QListWidgetItem_NewQListWidgetItem4(QtObjectPtr other){
	return new QListWidgetItem(*static_cast<QListWidgetItem*>(other));
}

int QListWidgetItem_CheckState(QtObjectPtr ptr){
	return static_cast<QListWidgetItem*>(ptr)->checkState();
}

QtObjectPtr QListWidgetItem_Clone(QtObjectPtr ptr){
	return static_cast<QListWidgetItem*>(ptr)->clone();
}

char* QListWidgetItem_Data(QtObjectPtr ptr, int role){
	return static_cast<QListWidgetItem*>(ptr)->data(role).toString().toUtf8().data();
}

int QListWidgetItem_Flags(QtObjectPtr ptr){
	return static_cast<QListWidgetItem*>(ptr)->flags();
}

int QListWidgetItem_IsHidden(QtObjectPtr ptr){
	return static_cast<QListWidgetItem*>(ptr)->isHidden();
}

int QListWidgetItem_IsSelected(QtObjectPtr ptr){
	return static_cast<QListWidgetItem*>(ptr)->isSelected();
}

QtObjectPtr QListWidgetItem_ListWidget(QtObjectPtr ptr){
	return static_cast<QListWidgetItem*>(ptr)->listWidget();
}

void QListWidgetItem_Read(QtObjectPtr ptr, QtObjectPtr in){
	static_cast<QListWidgetItem*>(ptr)->read(*static_cast<QDataStream*>(in));
}

void QListWidgetItem_SetBackground(QtObjectPtr ptr, QtObjectPtr brush){
	static_cast<QListWidgetItem*>(ptr)->setBackground(*static_cast<QBrush*>(brush));
}

void QListWidgetItem_SetCheckState(QtObjectPtr ptr, int state){
	static_cast<QListWidgetItem*>(ptr)->setCheckState(static_cast<Qt::CheckState>(state));
}

void QListWidgetItem_SetData(QtObjectPtr ptr, int role, char* value){
	static_cast<QListWidgetItem*>(ptr)->setData(role, QVariant(value));
}

void QListWidgetItem_SetFont(QtObjectPtr ptr, QtObjectPtr font){
	static_cast<QListWidgetItem*>(ptr)->setFont(*static_cast<QFont*>(font));
}

void QListWidgetItem_SetForeground(QtObjectPtr ptr, QtObjectPtr brush){
	static_cast<QListWidgetItem*>(ptr)->setForeground(*static_cast<QBrush*>(brush));
}

void QListWidgetItem_SetHidden(QtObjectPtr ptr, int hide){
	static_cast<QListWidgetItem*>(ptr)->setHidden(hide != 0);
}

void QListWidgetItem_SetIcon(QtObjectPtr ptr, QtObjectPtr icon){
	static_cast<QListWidgetItem*>(ptr)->setIcon(*static_cast<QIcon*>(icon));
}

void QListWidgetItem_SetSelected(QtObjectPtr ptr, int sele){
	static_cast<QListWidgetItem*>(ptr)->setSelected(sele != 0);
}

void QListWidgetItem_SetSizeHint(QtObjectPtr ptr, QtObjectPtr size){
	static_cast<QListWidgetItem*>(ptr)->setSizeHint(*static_cast<QSize*>(size));
}

void QListWidgetItem_SetStatusTip(QtObjectPtr ptr, char* statusTip){
	static_cast<QListWidgetItem*>(ptr)->setStatusTip(QString(statusTip));
}

void QListWidgetItem_SetText(QtObjectPtr ptr, char* text){
	static_cast<QListWidgetItem*>(ptr)->setText(QString(text));
}

void QListWidgetItem_SetTextAlignment(QtObjectPtr ptr, int alignment){
	static_cast<QListWidgetItem*>(ptr)->setTextAlignment(alignment);
}

void QListWidgetItem_SetToolTip(QtObjectPtr ptr, char* toolTip){
	static_cast<QListWidgetItem*>(ptr)->setToolTip(QString(toolTip));
}

void QListWidgetItem_SetWhatsThis(QtObjectPtr ptr, char* whatsThis){
	static_cast<QListWidgetItem*>(ptr)->setWhatsThis(QString(whatsThis));
}

char* QListWidgetItem_StatusTip(QtObjectPtr ptr){
	return static_cast<QListWidgetItem*>(ptr)->statusTip().toUtf8().data();
}

char* QListWidgetItem_Text(QtObjectPtr ptr){
	return static_cast<QListWidgetItem*>(ptr)->text().toUtf8().data();
}

int QListWidgetItem_TextAlignment(QtObjectPtr ptr){
	return static_cast<QListWidgetItem*>(ptr)->textAlignment();
}

char* QListWidgetItem_ToolTip(QtObjectPtr ptr){
	return static_cast<QListWidgetItem*>(ptr)->toolTip().toUtf8().data();
}

int QListWidgetItem_Type(QtObjectPtr ptr){
	return static_cast<QListWidgetItem*>(ptr)->type();
}

char* QListWidgetItem_WhatsThis(QtObjectPtr ptr){
	return static_cast<QListWidgetItem*>(ptr)->whatsThis().toUtf8().data();
}

void QListWidgetItem_Write(QtObjectPtr ptr, QtObjectPtr out){
	static_cast<QListWidgetItem*>(ptr)->write(*static_cast<QDataStream*>(out));
}

void QListWidgetItem_DestroyQListWidgetItem(QtObjectPtr ptr){
	static_cast<QListWidgetItem*>(ptr)->~QListWidgetItem();
}

