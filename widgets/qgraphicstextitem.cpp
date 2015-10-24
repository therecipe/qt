#include "qgraphicstextitem.h"
#include <QGraphicsItem>
#include <QStyleOptionGraphicsItem>
#include <QStyle>
#include <QObject>
#include <QWidget>
#include <QTextCursor>
#include <QString>
#include <QColor>
#include <QTextDocument>
#include <QStyleOption>
#include <QFont>
#include <QUrl>
#include <QModelIndex>
#include <QPoint>
#include <QPointF>
#include <QVariant>
#include <QPainter>
#include <QGraphicsTextItem>
#include "_cgo_export.h"

class MyQGraphicsTextItem: public QGraphicsTextItem {
public:
void Signal_LinkActivated(const QString & link){callbackQGraphicsTextItemLinkActivated(this->objectName().toUtf8().data(), link.toUtf8().data());};
void Signal_LinkHovered(const QString & link){callbackQGraphicsTextItemLinkHovered(this->objectName().toUtf8().data(), link.toUtf8().data());};
};

int QGraphicsTextItem_OpenExternalLinks(QtObjectPtr ptr){
	return static_cast<QGraphicsTextItem*>(ptr)->openExternalLinks();
}

void QGraphicsTextItem_SetOpenExternalLinks(QtObjectPtr ptr, int open){
	static_cast<QGraphicsTextItem*>(ptr)->setOpenExternalLinks(open != 0);
}

void QGraphicsTextItem_SetTextCursor(QtObjectPtr ptr, QtObjectPtr cursor){
	static_cast<QGraphicsTextItem*>(ptr)->setTextCursor(*static_cast<QTextCursor*>(cursor));
}

QtObjectPtr QGraphicsTextItem_NewQGraphicsTextItem(QtObjectPtr parent){
	return new QGraphicsTextItem(static_cast<QGraphicsItem*>(parent));
}

QtObjectPtr QGraphicsTextItem_NewQGraphicsTextItem2(char* text, QtObjectPtr parent){
	return new QGraphicsTextItem(QString(text), static_cast<QGraphicsItem*>(parent));
}

void QGraphicsTextItem_AdjustSize(QtObjectPtr ptr){
	static_cast<QGraphicsTextItem*>(ptr)->adjustSize();
}

int QGraphicsTextItem_Contains(QtObjectPtr ptr, QtObjectPtr point){
	return static_cast<QGraphicsTextItem*>(ptr)->contains(*static_cast<QPointF*>(point));
}

QtObjectPtr QGraphicsTextItem_Document(QtObjectPtr ptr){
	return static_cast<QGraphicsTextItem*>(ptr)->document();
}

int QGraphicsTextItem_IsObscuredBy(QtObjectPtr ptr, QtObjectPtr item){
	return static_cast<QGraphicsTextItem*>(ptr)->isObscuredBy(static_cast<QGraphicsItem*>(item));
}

void QGraphicsTextItem_ConnectLinkActivated(QtObjectPtr ptr){
	QObject::connect(static_cast<QGraphicsTextItem*>(ptr), static_cast<void (QGraphicsTextItem::*)(const QString &)>(&QGraphicsTextItem::linkActivated), static_cast<MyQGraphicsTextItem*>(ptr), static_cast<void (MyQGraphicsTextItem::*)(const QString &)>(&MyQGraphicsTextItem::Signal_LinkActivated));;
}

void QGraphicsTextItem_DisconnectLinkActivated(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QGraphicsTextItem*>(ptr), static_cast<void (QGraphicsTextItem::*)(const QString &)>(&QGraphicsTextItem::linkActivated), static_cast<MyQGraphicsTextItem*>(ptr), static_cast<void (MyQGraphicsTextItem::*)(const QString &)>(&MyQGraphicsTextItem::Signal_LinkActivated));;
}

void QGraphicsTextItem_ConnectLinkHovered(QtObjectPtr ptr){
	QObject::connect(static_cast<QGraphicsTextItem*>(ptr), static_cast<void (QGraphicsTextItem::*)(const QString &)>(&QGraphicsTextItem::linkHovered), static_cast<MyQGraphicsTextItem*>(ptr), static_cast<void (MyQGraphicsTextItem::*)(const QString &)>(&MyQGraphicsTextItem::Signal_LinkHovered));;
}

void QGraphicsTextItem_DisconnectLinkHovered(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QGraphicsTextItem*>(ptr), static_cast<void (QGraphicsTextItem::*)(const QString &)>(&QGraphicsTextItem::linkHovered), static_cast<MyQGraphicsTextItem*>(ptr), static_cast<void (MyQGraphicsTextItem::*)(const QString &)>(&MyQGraphicsTextItem::Signal_LinkHovered));;
}

void QGraphicsTextItem_Paint(QtObjectPtr ptr, QtObjectPtr painter, QtObjectPtr option, QtObjectPtr widget){
	static_cast<QGraphicsTextItem*>(ptr)->paint(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
}

void QGraphicsTextItem_SetDefaultTextColor(QtObjectPtr ptr, QtObjectPtr col){
	static_cast<QGraphicsTextItem*>(ptr)->setDefaultTextColor(*static_cast<QColor*>(col));
}

void QGraphicsTextItem_SetDocument(QtObjectPtr ptr, QtObjectPtr document){
	static_cast<QGraphicsTextItem*>(ptr)->setDocument(static_cast<QTextDocument*>(document));
}

void QGraphicsTextItem_SetFont(QtObjectPtr ptr, QtObjectPtr font){
	static_cast<QGraphicsTextItem*>(ptr)->setFont(*static_cast<QFont*>(font));
}

void QGraphicsTextItem_SetHtml(QtObjectPtr ptr, char* text){
	static_cast<QGraphicsTextItem*>(ptr)->setHtml(QString(text));
}

void QGraphicsTextItem_SetPlainText(QtObjectPtr ptr, char* text){
	static_cast<QGraphicsTextItem*>(ptr)->setPlainText(QString(text));
}

void QGraphicsTextItem_SetTabChangesFocus(QtObjectPtr ptr, int b){
	static_cast<QGraphicsTextItem*>(ptr)->setTabChangesFocus(b != 0);
}

void QGraphicsTextItem_SetTextInteractionFlags(QtObjectPtr ptr, int flags){
	static_cast<QGraphicsTextItem*>(ptr)->setTextInteractionFlags(static_cast<Qt::TextInteractionFlag>(flags));
}

int QGraphicsTextItem_TabChangesFocus(QtObjectPtr ptr){
	return static_cast<QGraphicsTextItem*>(ptr)->tabChangesFocus();
}

int QGraphicsTextItem_TextInteractionFlags(QtObjectPtr ptr){
	return static_cast<QGraphicsTextItem*>(ptr)->textInteractionFlags();
}

char* QGraphicsTextItem_ToHtml(QtObjectPtr ptr){
	return static_cast<QGraphicsTextItem*>(ptr)->toHtml().toUtf8().data();
}

char* QGraphicsTextItem_ToPlainText(QtObjectPtr ptr){
	return static_cast<QGraphicsTextItem*>(ptr)->toPlainText().toUtf8().data();
}

int QGraphicsTextItem_Type(QtObjectPtr ptr){
	return static_cast<QGraphicsTextItem*>(ptr)->type();
}

void QGraphicsTextItem_DestroyQGraphicsTextItem(QtObjectPtr ptr){
	static_cast<QGraphicsTextItem*>(ptr)->~QGraphicsTextItem();
}

