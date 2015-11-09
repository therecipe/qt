#include "qtablewidgetitem.h"
#include <QVariant>
#include <QUrl>
#include <QDataStream>
#include <QBrush>
#include <QTableWidget>
#include <QString>
#include <QIcon>
#include <QFont>
#include <QSize>
#include <QModelIndex>
#include <QTableWidgetItem>
#include "_cgo_export.h"

class MyQTableWidgetItem: public QTableWidgetItem {
public:
};

void QTableWidgetItem_SetFlags(void* ptr, int flags){
	static_cast<QTableWidgetItem*>(ptr)->setFlags(static_cast<Qt::ItemFlag>(flags));
}

void* QTableWidgetItem_NewQTableWidgetItem3(void* icon, char* text, int ty){
	return new QTableWidgetItem(*static_cast<QIcon*>(icon), QString(text), ty);
}

void* QTableWidgetItem_NewQTableWidgetItem2(char* text, int ty){
	return new QTableWidgetItem(QString(text), ty);
}

void* QTableWidgetItem_NewQTableWidgetItem4(void* other){
	return new QTableWidgetItem(*static_cast<QTableWidgetItem*>(other));
}

void* QTableWidgetItem_NewQTableWidgetItem(int ty){
	return new QTableWidgetItem(ty);
}

void* QTableWidgetItem_Background(void* ptr){
	return new QBrush(static_cast<QTableWidgetItem*>(ptr)->background());
}

int QTableWidgetItem_CheckState(void* ptr){
	return static_cast<QTableWidgetItem*>(ptr)->checkState();
}

void* QTableWidgetItem_Clone(void* ptr){
	return static_cast<QTableWidgetItem*>(ptr)->clone();
}

int QTableWidgetItem_Column(void* ptr){
	return static_cast<QTableWidgetItem*>(ptr)->column();
}

void* QTableWidgetItem_Data(void* ptr, int role){
	return new QVariant(static_cast<QTableWidgetItem*>(ptr)->data(role));
}

int QTableWidgetItem_Flags(void* ptr){
	return static_cast<QTableWidgetItem*>(ptr)->flags();
}

void* QTableWidgetItem_Foreground(void* ptr){
	return new QBrush(static_cast<QTableWidgetItem*>(ptr)->foreground());
}

int QTableWidgetItem_IsSelected(void* ptr){
	return static_cast<QTableWidgetItem*>(ptr)->isSelected();
}

void QTableWidgetItem_Read(void* ptr, void* in){
	static_cast<QTableWidgetItem*>(ptr)->read(*static_cast<QDataStream*>(in));
}

int QTableWidgetItem_Row(void* ptr){
	return static_cast<QTableWidgetItem*>(ptr)->row();
}

void QTableWidgetItem_SetBackground(void* ptr, void* brush){
	static_cast<QTableWidgetItem*>(ptr)->setBackground(*static_cast<QBrush*>(brush));
}

void QTableWidgetItem_SetCheckState(void* ptr, int state){
	static_cast<QTableWidgetItem*>(ptr)->setCheckState(static_cast<Qt::CheckState>(state));
}

void QTableWidgetItem_SetData(void* ptr, int role, void* value){
	static_cast<QTableWidgetItem*>(ptr)->setData(role, *static_cast<QVariant*>(value));
}

void QTableWidgetItem_SetFont(void* ptr, void* font){
	static_cast<QTableWidgetItem*>(ptr)->setFont(*static_cast<QFont*>(font));
}

void QTableWidgetItem_SetForeground(void* ptr, void* brush){
	static_cast<QTableWidgetItem*>(ptr)->setForeground(*static_cast<QBrush*>(brush));
}

void QTableWidgetItem_SetIcon(void* ptr, void* icon){
	static_cast<QTableWidgetItem*>(ptr)->setIcon(*static_cast<QIcon*>(icon));
}

void QTableWidgetItem_SetSelected(void* ptr, int sele){
	static_cast<QTableWidgetItem*>(ptr)->setSelected(sele != 0);
}

void QTableWidgetItem_SetSizeHint(void* ptr, void* size){
	static_cast<QTableWidgetItem*>(ptr)->setSizeHint(*static_cast<QSize*>(size));
}

void QTableWidgetItem_SetStatusTip(void* ptr, char* statusTip){
	static_cast<QTableWidgetItem*>(ptr)->setStatusTip(QString(statusTip));
}

void QTableWidgetItem_SetText(void* ptr, char* text){
	static_cast<QTableWidgetItem*>(ptr)->setText(QString(text));
}

void QTableWidgetItem_SetTextAlignment(void* ptr, int alignment){
	static_cast<QTableWidgetItem*>(ptr)->setTextAlignment(alignment);
}

void QTableWidgetItem_SetToolTip(void* ptr, char* toolTip){
	static_cast<QTableWidgetItem*>(ptr)->setToolTip(QString(toolTip));
}

void QTableWidgetItem_SetWhatsThis(void* ptr, char* whatsThis){
	static_cast<QTableWidgetItem*>(ptr)->setWhatsThis(QString(whatsThis));
}

char* QTableWidgetItem_StatusTip(void* ptr){
	return static_cast<QTableWidgetItem*>(ptr)->statusTip().toUtf8().data();
}

void* QTableWidgetItem_TableWidget(void* ptr){
	return static_cast<QTableWidgetItem*>(ptr)->tableWidget();
}

char* QTableWidgetItem_Text(void* ptr){
	return static_cast<QTableWidgetItem*>(ptr)->text().toUtf8().data();
}

int QTableWidgetItem_TextAlignment(void* ptr){
	return static_cast<QTableWidgetItem*>(ptr)->textAlignment();
}

char* QTableWidgetItem_ToolTip(void* ptr){
	return static_cast<QTableWidgetItem*>(ptr)->toolTip().toUtf8().data();
}

int QTableWidgetItem_Type(void* ptr){
	return static_cast<QTableWidgetItem*>(ptr)->type();
}

char* QTableWidgetItem_WhatsThis(void* ptr){
	return static_cast<QTableWidgetItem*>(ptr)->whatsThis().toUtf8().data();
}

void QTableWidgetItem_Write(void* ptr, void* out){
	static_cast<QTableWidgetItem*>(ptr)->write(*static_cast<QDataStream*>(out));
}

void QTableWidgetItem_DestroyQTableWidgetItem(void* ptr){
	static_cast<QTableWidgetItem*>(ptr)->~QTableWidgetItem();
}

