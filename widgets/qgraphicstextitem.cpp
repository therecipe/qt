#include "qgraphicstextitem.h"
#include <QUrl>
#include <QStyleOptionGraphicsItem>
#include <QObject>
#include <QString>
#include <QModelIndex>
#include <QWidget>
#include <QStyleOption>
#include <QStyle>
#include <QColor>
#include <QPainter>
#include <QPoint>
#include <QTextDocument>
#include <QVariant>
#include <QFont>
#include <QGraphicsItem>
#include <QPointF>
#include <QTextCursor>
#include <QGraphicsTextItem>
#include "_cgo_export.h"

class MyQGraphicsTextItem: public QGraphicsTextItem {
public:
void Signal_LinkActivated(const QString & link){callbackQGraphicsTextItemLinkActivated(this->objectName().toUtf8().data(), link.toUtf8().data());};
void Signal_LinkHovered(const QString & link){callbackQGraphicsTextItemLinkHovered(this->objectName().toUtf8().data(), link.toUtf8().data());};
};

int QGraphicsTextItem_OpenExternalLinks(void* ptr){
	return static_cast<QGraphicsTextItem*>(ptr)->openExternalLinks();
}

void QGraphicsTextItem_SetOpenExternalLinks(void* ptr, int open){
	static_cast<QGraphicsTextItem*>(ptr)->setOpenExternalLinks(open != 0);
}

void QGraphicsTextItem_SetTextCursor(void* ptr, void* cursor){
	static_cast<QGraphicsTextItem*>(ptr)->setTextCursor(*static_cast<QTextCursor*>(cursor));
}

void QGraphicsTextItem_AdjustSize(void* ptr){
	static_cast<QGraphicsTextItem*>(ptr)->adjustSize();
}

int QGraphicsTextItem_Contains(void* ptr, void* point){
	return static_cast<QGraphicsTextItem*>(ptr)->contains(*static_cast<QPointF*>(point));
}

void* QGraphicsTextItem_DefaultTextColor(void* ptr){
	return new QColor(static_cast<QGraphicsTextItem*>(ptr)->defaultTextColor());
}

void* QGraphicsTextItem_Document(void* ptr){
	return static_cast<QGraphicsTextItem*>(ptr)->document();
}

int QGraphicsTextItem_IsObscuredBy(void* ptr, void* item){
	return static_cast<QGraphicsTextItem*>(ptr)->isObscuredBy(static_cast<QGraphicsItem*>(item));
}

void QGraphicsTextItem_ConnectLinkActivated(void* ptr){
	QObject::connect(static_cast<QGraphicsTextItem*>(ptr), static_cast<void (QGraphicsTextItem::*)(const QString &)>(&QGraphicsTextItem::linkActivated), static_cast<MyQGraphicsTextItem*>(ptr), static_cast<void (MyQGraphicsTextItem::*)(const QString &)>(&MyQGraphicsTextItem::Signal_LinkActivated));;
}

void QGraphicsTextItem_DisconnectLinkActivated(void* ptr){
	QObject::disconnect(static_cast<QGraphicsTextItem*>(ptr), static_cast<void (QGraphicsTextItem::*)(const QString &)>(&QGraphicsTextItem::linkActivated), static_cast<MyQGraphicsTextItem*>(ptr), static_cast<void (MyQGraphicsTextItem::*)(const QString &)>(&MyQGraphicsTextItem::Signal_LinkActivated));;
}

void QGraphicsTextItem_ConnectLinkHovered(void* ptr){
	QObject::connect(static_cast<QGraphicsTextItem*>(ptr), static_cast<void (QGraphicsTextItem::*)(const QString &)>(&QGraphicsTextItem::linkHovered), static_cast<MyQGraphicsTextItem*>(ptr), static_cast<void (MyQGraphicsTextItem::*)(const QString &)>(&MyQGraphicsTextItem::Signal_LinkHovered));;
}

void QGraphicsTextItem_DisconnectLinkHovered(void* ptr){
	QObject::disconnect(static_cast<QGraphicsTextItem*>(ptr), static_cast<void (QGraphicsTextItem::*)(const QString &)>(&QGraphicsTextItem::linkHovered), static_cast<MyQGraphicsTextItem*>(ptr), static_cast<void (MyQGraphicsTextItem::*)(const QString &)>(&MyQGraphicsTextItem::Signal_LinkHovered));;
}

void QGraphicsTextItem_Paint(void* ptr, void* painter, void* option, void* widget){
	static_cast<QGraphicsTextItem*>(ptr)->paint(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
}

void QGraphicsTextItem_SetDefaultTextColor(void* ptr, void* col){
	static_cast<QGraphicsTextItem*>(ptr)->setDefaultTextColor(*static_cast<QColor*>(col));
}

void QGraphicsTextItem_SetDocument(void* ptr, void* document){
	static_cast<QGraphicsTextItem*>(ptr)->setDocument(static_cast<QTextDocument*>(document));
}

void QGraphicsTextItem_SetFont(void* ptr, void* font){
	static_cast<QGraphicsTextItem*>(ptr)->setFont(*static_cast<QFont*>(font));
}

void QGraphicsTextItem_SetHtml(void* ptr, char* text){
	static_cast<QGraphicsTextItem*>(ptr)->setHtml(QString(text));
}

void QGraphicsTextItem_SetPlainText(void* ptr, char* text){
	static_cast<QGraphicsTextItem*>(ptr)->setPlainText(QString(text));
}

void QGraphicsTextItem_SetTabChangesFocus(void* ptr, int b){
	static_cast<QGraphicsTextItem*>(ptr)->setTabChangesFocus(b != 0);
}

void QGraphicsTextItem_SetTextInteractionFlags(void* ptr, int flags){
	static_cast<QGraphicsTextItem*>(ptr)->setTextInteractionFlags(static_cast<Qt::TextInteractionFlag>(flags));
}

void QGraphicsTextItem_SetTextWidth(void* ptr, double width){
	static_cast<QGraphicsTextItem*>(ptr)->setTextWidth(static_cast<qreal>(width));
}

int QGraphicsTextItem_TabChangesFocus(void* ptr){
	return static_cast<QGraphicsTextItem*>(ptr)->tabChangesFocus();
}

int QGraphicsTextItem_TextInteractionFlags(void* ptr){
	return static_cast<QGraphicsTextItem*>(ptr)->textInteractionFlags();
}

double QGraphicsTextItem_TextWidth(void* ptr){
	return static_cast<double>(static_cast<QGraphicsTextItem*>(ptr)->textWidth());
}

char* QGraphicsTextItem_ToHtml(void* ptr){
	return static_cast<QGraphicsTextItem*>(ptr)->toHtml().toUtf8().data();
}

char* QGraphicsTextItem_ToPlainText(void* ptr){
	return static_cast<QGraphicsTextItem*>(ptr)->toPlainText().toUtf8().data();
}

int QGraphicsTextItem_Type(void* ptr){
	return static_cast<QGraphicsTextItem*>(ptr)->type();
}

void QGraphicsTextItem_DestroyQGraphicsTextItem(void* ptr){
	static_cast<QGraphicsTextItem*>(ptr)->~QGraphicsTextItem();
}

