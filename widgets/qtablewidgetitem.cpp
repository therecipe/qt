#include "qtablewidgetitem.h"
#include <QString>
#include <QUrl>
#include <QModelIndex>
#include <QFont>
#include <QSize>
#include <QVariant>
#include <QDataStream>
#include <QTableWidget>
#include <QBrush>
#include <QIcon>
#include <QTableWidgetItem>
#include "_cgo_export.h"

class MyQTableWidgetItem: public QTableWidgetItem {
public:
};

void QTableWidgetItem_SetFlags(QtObjectPtr ptr, int flags){
	static_cast<QTableWidgetItem*>(ptr)->setFlags(static_cast<Qt::ItemFlag>(flags));
}

QtObjectPtr QTableWidgetItem_NewQTableWidgetItem3(QtObjectPtr icon, char* text, int ty){
	return new QTableWidgetItem(*static_cast<QIcon*>(icon), QString(text), ty);
}

QtObjectPtr QTableWidgetItem_NewQTableWidgetItem2(char* text, int ty){
	return new QTableWidgetItem(QString(text), ty);
}

QtObjectPtr QTableWidgetItem_NewQTableWidgetItem4(QtObjectPtr other){
	return new QTableWidgetItem(*static_cast<QTableWidgetItem*>(other));
}

QtObjectPtr QTableWidgetItem_NewQTableWidgetItem(int ty){
	return new QTableWidgetItem(ty);
}

int QTableWidgetItem_CheckState(QtObjectPtr ptr){
	return static_cast<QTableWidgetItem*>(ptr)->checkState();
}

QtObjectPtr QTableWidgetItem_Clone(QtObjectPtr ptr){
	return static_cast<QTableWidgetItem*>(ptr)->clone();
}

int QTableWidgetItem_Column(QtObjectPtr ptr){
	return static_cast<QTableWidgetItem*>(ptr)->column();
}

char* QTableWidgetItem_Data(QtObjectPtr ptr, int role){
	return static_cast<QTableWidgetItem*>(ptr)->data(role).toString().toUtf8().data();
}

int QTableWidgetItem_Flags(QtObjectPtr ptr){
	return static_cast<QTableWidgetItem*>(ptr)->flags();
}

int QTableWidgetItem_IsSelected(QtObjectPtr ptr){
	return static_cast<QTableWidgetItem*>(ptr)->isSelected();
}

void QTableWidgetItem_Read(QtObjectPtr ptr, QtObjectPtr in){
	static_cast<QTableWidgetItem*>(ptr)->read(*static_cast<QDataStream*>(in));
}

int QTableWidgetItem_Row(QtObjectPtr ptr){
	return static_cast<QTableWidgetItem*>(ptr)->row();
}

void QTableWidgetItem_SetBackground(QtObjectPtr ptr, QtObjectPtr brush){
	static_cast<QTableWidgetItem*>(ptr)->setBackground(*static_cast<QBrush*>(brush));
}

void QTableWidgetItem_SetCheckState(QtObjectPtr ptr, int state){
	static_cast<QTableWidgetItem*>(ptr)->setCheckState(static_cast<Qt::CheckState>(state));
}

void QTableWidgetItem_SetData(QtObjectPtr ptr, int role, char* value){
	static_cast<QTableWidgetItem*>(ptr)->setData(role, QVariant(value));
}

void QTableWidgetItem_SetFont(QtObjectPtr ptr, QtObjectPtr font){
	static_cast<QTableWidgetItem*>(ptr)->setFont(*static_cast<QFont*>(font));
}

void QTableWidgetItem_SetForeground(QtObjectPtr ptr, QtObjectPtr brush){
	static_cast<QTableWidgetItem*>(ptr)->setForeground(*static_cast<QBrush*>(brush));
}

void QTableWidgetItem_SetIcon(QtObjectPtr ptr, QtObjectPtr icon){
	static_cast<QTableWidgetItem*>(ptr)->setIcon(*static_cast<QIcon*>(icon));
}

void QTableWidgetItem_SetSelected(QtObjectPtr ptr, int sele){
	static_cast<QTableWidgetItem*>(ptr)->setSelected(sele != 0);
}

void QTableWidgetItem_SetSizeHint(QtObjectPtr ptr, QtObjectPtr size){
	static_cast<QTableWidgetItem*>(ptr)->setSizeHint(*static_cast<QSize*>(size));
}

void QTableWidgetItem_SetStatusTip(QtObjectPtr ptr, char* statusTip){
	static_cast<QTableWidgetItem*>(ptr)->setStatusTip(QString(statusTip));
}

void QTableWidgetItem_SetText(QtObjectPtr ptr, char* text){
	static_cast<QTableWidgetItem*>(ptr)->setText(QString(text));
}

void QTableWidgetItem_SetTextAlignment(QtObjectPtr ptr, int alignment){
	static_cast<QTableWidgetItem*>(ptr)->setTextAlignment(alignment);
}

void QTableWidgetItem_SetToolTip(QtObjectPtr ptr, char* toolTip){
	static_cast<QTableWidgetItem*>(ptr)->setToolTip(QString(toolTip));
}

void QTableWidgetItem_SetWhatsThis(QtObjectPtr ptr, char* whatsThis){
	static_cast<QTableWidgetItem*>(ptr)->setWhatsThis(QString(whatsThis));
}

char* QTableWidgetItem_StatusTip(QtObjectPtr ptr){
	return static_cast<QTableWidgetItem*>(ptr)->statusTip().toUtf8().data();
}

QtObjectPtr QTableWidgetItem_TableWidget(QtObjectPtr ptr){
	return static_cast<QTableWidgetItem*>(ptr)->tableWidget();
}

char* QTableWidgetItem_Text(QtObjectPtr ptr){
	return static_cast<QTableWidgetItem*>(ptr)->text().toUtf8().data();
}

int QTableWidgetItem_TextAlignment(QtObjectPtr ptr){
	return static_cast<QTableWidgetItem*>(ptr)->textAlignment();
}

char* QTableWidgetItem_ToolTip(QtObjectPtr ptr){
	return static_cast<QTableWidgetItem*>(ptr)->toolTip().toUtf8().data();
}

int QTableWidgetItem_Type(QtObjectPtr ptr){
	return static_cast<QTableWidgetItem*>(ptr)->type();
}

char* QTableWidgetItem_WhatsThis(QtObjectPtr ptr){
	return static_cast<QTableWidgetItem*>(ptr)->whatsThis().toUtf8().data();
}

void QTableWidgetItem_Write(QtObjectPtr ptr, QtObjectPtr out){
	static_cast<QTableWidgetItem*>(ptr)->write(*static_cast<QDataStream*>(out));
}

void QTableWidgetItem_DestroyQTableWidgetItem(QtObjectPtr ptr){
	static_cast<QTableWidgetItem*>(ptr)->~QTableWidgetItem();
}

