#include "qlabel.h"
#include <QMetaObject>
#include <QModelIndex>
#include <QWidget>
#include <QMovie>
#include <QPixmap>
#include <QPicture>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QLabel>
#include "_cgo_export.h"

class MyQLabel: public QLabel {
public:
void Signal_LinkActivated(const QString & link){callbackQLabelLinkActivated(this->objectName().toUtf8().data(), link.toUtf8().data());};
void Signal_LinkHovered(const QString & link){callbackQLabelLinkHovered(this->objectName().toUtf8().data(), link.toUtf8().data());};
};

int QLabel_Alignment(QtObjectPtr ptr){
	return static_cast<QLabel*>(ptr)->alignment();
}

int QLabel_HasScaledContents(QtObjectPtr ptr){
	return static_cast<QLabel*>(ptr)->hasScaledContents();
}

int QLabel_HasSelectedText(QtObjectPtr ptr){
	return static_cast<QLabel*>(ptr)->hasSelectedText();
}

int QLabel_Indent(QtObjectPtr ptr){
	return static_cast<QLabel*>(ptr)->indent();
}

int QLabel_Margin(QtObjectPtr ptr){
	return static_cast<QLabel*>(ptr)->margin();
}

int QLabel_OpenExternalLinks(QtObjectPtr ptr){
	return static_cast<QLabel*>(ptr)->openExternalLinks();
}

QtObjectPtr QLabel_Pixmap(QtObjectPtr ptr){
	return const_cast<QPixmap*>(static_cast<QLabel*>(ptr)->pixmap());
}

char* QLabel_SelectedText(QtObjectPtr ptr){
	return static_cast<QLabel*>(ptr)->selectedText().toUtf8().data();
}

void QLabel_SetAlignment(QtObjectPtr ptr, int v){
	static_cast<QLabel*>(ptr)->setAlignment(static_cast<Qt::AlignmentFlag>(v));
}

void QLabel_SetIndent(QtObjectPtr ptr, int v){
	static_cast<QLabel*>(ptr)->setIndent(v);
}

void QLabel_SetMargin(QtObjectPtr ptr, int v){
	static_cast<QLabel*>(ptr)->setMargin(v);
}

void QLabel_SetOpenExternalLinks(QtObjectPtr ptr, int open){
	static_cast<QLabel*>(ptr)->setOpenExternalLinks(open != 0);
}

void QLabel_SetPixmap(QtObjectPtr ptr, QtObjectPtr v){
	QMetaObject::invokeMethod(static_cast<QLabel*>(ptr), "setPixmap", Q_ARG(QPixmap, *static_cast<QPixmap*>(v)));
}

void QLabel_SetScaledContents(QtObjectPtr ptr, int v){
	static_cast<QLabel*>(ptr)->setScaledContents(v != 0);
}

void QLabel_SetText(QtObjectPtr ptr, char* v){
	QMetaObject::invokeMethod(static_cast<QLabel*>(ptr), "setText", Q_ARG(QString, QString(v)));
}

void QLabel_SetTextFormat(QtObjectPtr ptr, int v){
	static_cast<QLabel*>(ptr)->setTextFormat(static_cast<Qt::TextFormat>(v));
}

void QLabel_SetTextInteractionFlags(QtObjectPtr ptr, int flags){
	static_cast<QLabel*>(ptr)->setTextInteractionFlags(static_cast<Qt::TextInteractionFlag>(flags));
}

void QLabel_SetWordWrap(QtObjectPtr ptr, int on){
	static_cast<QLabel*>(ptr)->setWordWrap(on != 0);
}

char* QLabel_Text(QtObjectPtr ptr){
	return static_cast<QLabel*>(ptr)->text().toUtf8().data();
}

int QLabel_TextFormat(QtObjectPtr ptr){
	return static_cast<QLabel*>(ptr)->textFormat();
}

int QLabel_TextInteractionFlags(QtObjectPtr ptr){
	return static_cast<QLabel*>(ptr)->textInteractionFlags();
}

int QLabel_WordWrap(QtObjectPtr ptr){
	return static_cast<QLabel*>(ptr)->wordWrap();
}

QtObjectPtr QLabel_NewQLabel(QtObjectPtr parent, int f){
	return new QLabel(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(f));
}

QtObjectPtr QLabel_NewQLabel2(char* text, QtObjectPtr parent, int f){
	return new QLabel(QString(text), static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(f));
}

QtObjectPtr QLabel_Buddy(QtObjectPtr ptr){
	return static_cast<QLabel*>(ptr)->buddy();
}

void QLabel_Clear(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QLabel*>(ptr), "clear");
}

int QLabel_HeightForWidth(QtObjectPtr ptr, int w){
	return static_cast<QLabel*>(ptr)->heightForWidth(w);
}

void QLabel_ConnectLinkActivated(QtObjectPtr ptr){
	QObject::connect(static_cast<QLabel*>(ptr), static_cast<void (QLabel::*)(const QString &)>(&QLabel::linkActivated), static_cast<MyQLabel*>(ptr), static_cast<void (MyQLabel::*)(const QString &)>(&MyQLabel::Signal_LinkActivated));;
}

void QLabel_DisconnectLinkActivated(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QLabel*>(ptr), static_cast<void (QLabel::*)(const QString &)>(&QLabel::linkActivated), static_cast<MyQLabel*>(ptr), static_cast<void (MyQLabel::*)(const QString &)>(&MyQLabel::Signal_LinkActivated));;
}

void QLabel_ConnectLinkHovered(QtObjectPtr ptr){
	QObject::connect(static_cast<QLabel*>(ptr), static_cast<void (QLabel::*)(const QString &)>(&QLabel::linkHovered), static_cast<MyQLabel*>(ptr), static_cast<void (MyQLabel::*)(const QString &)>(&MyQLabel::Signal_LinkHovered));;
}

void QLabel_DisconnectLinkHovered(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QLabel*>(ptr), static_cast<void (QLabel::*)(const QString &)>(&QLabel::linkHovered), static_cast<MyQLabel*>(ptr), static_cast<void (MyQLabel::*)(const QString &)>(&MyQLabel::Signal_LinkHovered));;
}

QtObjectPtr QLabel_Movie(QtObjectPtr ptr){
	return static_cast<QLabel*>(ptr)->movie();
}

QtObjectPtr QLabel_Picture(QtObjectPtr ptr){
	return const_cast<QPicture*>(static_cast<QLabel*>(ptr)->picture());
}

int QLabel_SelectionStart(QtObjectPtr ptr){
	return static_cast<QLabel*>(ptr)->selectionStart();
}

void QLabel_SetBuddy(QtObjectPtr ptr, QtObjectPtr buddy){
	static_cast<QLabel*>(ptr)->setBuddy(static_cast<QWidget*>(buddy));
}

void QLabel_SetMovie(QtObjectPtr ptr, QtObjectPtr movie){
	QMetaObject::invokeMethod(static_cast<QLabel*>(ptr), "setMovie", Q_ARG(QMovie*, static_cast<QMovie*>(movie)));
}

void QLabel_SetNum(QtObjectPtr ptr, int num){
	QMetaObject::invokeMethod(static_cast<QLabel*>(ptr), "setNum", Q_ARG(int, num));
}

void QLabel_SetPicture(QtObjectPtr ptr, QtObjectPtr picture){
	QMetaObject::invokeMethod(static_cast<QLabel*>(ptr), "setPicture", Q_ARG(QPicture, *static_cast<QPicture*>(picture)));
}

void QLabel_SetSelection(QtObjectPtr ptr, int start, int length){
	static_cast<QLabel*>(ptr)->setSelection(start, length);
}

void QLabel_DestroyQLabel(QtObjectPtr ptr){
	static_cast<QLabel*>(ptr)->~QLabel();
}

