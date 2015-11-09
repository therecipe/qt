#include "qlabel.h"
#include <QVariant>
#include <QPicture>
#include <QMovie>
#include <QWidget>
#include <QPixmap>
#include <QString>
#include <QUrl>
#include <QModelIndex>
#include <QMetaObject>
#include <QObject>
#include <QLabel>
#include "_cgo_export.h"

class MyQLabel: public QLabel {
public:
void Signal_LinkActivated(const QString & link){callbackQLabelLinkActivated(this->objectName().toUtf8().data(), link.toUtf8().data());};
void Signal_LinkHovered(const QString & link){callbackQLabelLinkHovered(this->objectName().toUtf8().data(), link.toUtf8().data());};
};

int QLabel_Alignment(void* ptr){
	return static_cast<QLabel*>(ptr)->alignment();
}

int QLabel_HasScaledContents(void* ptr){
	return static_cast<QLabel*>(ptr)->hasScaledContents();
}

int QLabel_HasSelectedText(void* ptr){
	return static_cast<QLabel*>(ptr)->hasSelectedText();
}

int QLabel_Indent(void* ptr){
	return static_cast<QLabel*>(ptr)->indent();
}

int QLabel_Margin(void* ptr){
	return static_cast<QLabel*>(ptr)->margin();
}

int QLabel_OpenExternalLinks(void* ptr){
	return static_cast<QLabel*>(ptr)->openExternalLinks();
}

void* QLabel_Pixmap(void* ptr){
	return const_cast<QPixmap*>(static_cast<QLabel*>(ptr)->pixmap());
}

char* QLabel_SelectedText(void* ptr){
	return static_cast<QLabel*>(ptr)->selectedText().toUtf8().data();
}

void QLabel_SetAlignment(void* ptr, int v){
	static_cast<QLabel*>(ptr)->setAlignment(static_cast<Qt::AlignmentFlag>(v));
}

void QLabel_SetIndent(void* ptr, int v){
	static_cast<QLabel*>(ptr)->setIndent(v);
}

void QLabel_SetMargin(void* ptr, int v){
	static_cast<QLabel*>(ptr)->setMargin(v);
}

void QLabel_SetOpenExternalLinks(void* ptr, int open){
	static_cast<QLabel*>(ptr)->setOpenExternalLinks(open != 0);
}

void QLabel_SetPixmap(void* ptr, void* v){
	QMetaObject::invokeMethod(static_cast<QLabel*>(ptr), "setPixmap", Q_ARG(QPixmap, *static_cast<QPixmap*>(v)));
}

void QLabel_SetScaledContents(void* ptr, int v){
	static_cast<QLabel*>(ptr)->setScaledContents(v != 0);
}

void QLabel_SetText(void* ptr, char* v){
	QMetaObject::invokeMethod(static_cast<QLabel*>(ptr), "setText", Q_ARG(QString, QString(v)));
}

void QLabel_SetTextFormat(void* ptr, int v){
	static_cast<QLabel*>(ptr)->setTextFormat(static_cast<Qt::TextFormat>(v));
}

void QLabel_SetTextInteractionFlags(void* ptr, int flags){
	static_cast<QLabel*>(ptr)->setTextInteractionFlags(static_cast<Qt::TextInteractionFlag>(flags));
}

void QLabel_SetWordWrap(void* ptr, int on){
	static_cast<QLabel*>(ptr)->setWordWrap(on != 0);
}

char* QLabel_Text(void* ptr){
	return static_cast<QLabel*>(ptr)->text().toUtf8().data();
}

int QLabel_TextFormat(void* ptr){
	return static_cast<QLabel*>(ptr)->textFormat();
}

int QLabel_TextInteractionFlags(void* ptr){
	return static_cast<QLabel*>(ptr)->textInteractionFlags();
}

int QLabel_WordWrap(void* ptr){
	return static_cast<QLabel*>(ptr)->wordWrap();
}

void* QLabel_NewQLabel(void* parent, int f){
	return new QLabel(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(f));
}

void* QLabel_NewQLabel2(char* text, void* parent, int f){
	return new QLabel(QString(text), static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(f));
}

void* QLabel_Buddy(void* ptr){
	return static_cast<QLabel*>(ptr)->buddy();
}

void QLabel_Clear(void* ptr){
	QMetaObject::invokeMethod(static_cast<QLabel*>(ptr), "clear");
}

int QLabel_HeightForWidth(void* ptr, int w){
	return static_cast<QLabel*>(ptr)->heightForWidth(w);
}

void QLabel_ConnectLinkActivated(void* ptr){
	QObject::connect(static_cast<QLabel*>(ptr), static_cast<void (QLabel::*)(const QString &)>(&QLabel::linkActivated), static_cast<MyQLabel*>(ptr), static_cast<void (MyQLabel::*)(const QString &)>(&MyQLabel::Signal_LinkActivated));;
}

void QLabel_DisconnectLinkActivated(void* ptr){
	QObject::disconnect(static_cast<QLabel*>(ptr), static_cast<void (QLabel::*)(const QString &)>(&QLabel::linkActivated), static_cast<MyQLabel*>(ptr), static_cast<void (MyQLabel::*)(const QString &)>(&MyQLabel::Signal_LinkActivated));;
}

void QLabel_ConnectLinkHovered(void* ptr){
	QObject::connect(static_cast<QLabel*>(ptr), static_cast<void (QLabel::*)(const QString &)>(&QLabel::linkHovered), static_cast<MyQLabel*>(ptr), static_cast<void (MyQLabel::*)(const QString &)>(&MyQLabel::Signal_LinkHovered));;
}

void QLabel_DisconnectLinkHovered(void* ptr){
	QObject::disconnect(static_cast<QLabel*>(ptr), static_cast<void (QLabel::*)(const QString &)>(&QLabel::linkHovered), static_cast<MyQLabel*>(ptr), static_cast<void (MyQLabel::*)(const QString &)>(&MyQLabel::Signal_LinkHovered));;
}

void* QLabel_Movie(void* ptr){
	return static_cast<QLabel*>(ptr)->movie();
}

void* QLabel_Picture(void* ptr){
	return const_cast<QPicture*>(static_cast<QLabel*>(ptr)->picture());
}

int QLabel_SelectionStart(void* ptr){
	return static_cast<QLabel*>(ptr)->selectionStart();
}

void QLabel_SetBuddy(void* ptr, void* buddy){
	static_cast<QLabel*>(ptr)->setBuddy(static_cast<QWidget*>(buddy));
}

void QLabel_SetMovie(void* ptr, void* movie){
	QMetaObject::invokeMethod(static_cast<QLabel*>(ptr), "setMovie", Q_ARG(QMovie*, static_cast<QMovie*>(movie)));
}

void QLabel_SetNum(void* ptr, int num){
	QMetaObject::invokeMethod(static_cast<QLabel*>(ptr), "setNum", Q_ARG(int, num));
}

void QLabel_SetPicture(void* ptr, void* picture){
	QMetaObject::invokeMethod(static_cast<QLabel*>(ptr), "setPicture", Q_ARG(QPicture, *static_cast<QPicture*>(picture)));
}

void QLabel_SetSelection(void* ptr, int start, int length){
	static_cast<QLabel*>(ptr)->setSelection(start, length);
}

void QLabel_DestroyQLabel(void* ptr){
	static_cast<QLabel*>(ptr)->~QLabel();
}

