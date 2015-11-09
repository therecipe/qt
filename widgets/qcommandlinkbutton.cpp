#include "qcommandlinkbutton.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QWidget>
#include <QCommandLinkButton>
#include "_cgo_export.h"

class MyQCommandLinkButton: public QCommandLinkButton {
public:
};

char* QCommandLinkButton_Description(void* ptr){
	return static_cast<QCommandLinkButton*>(ptr)->description().toUtf8().data();
}

void QCommandLinkButton_SetDescription(void* ptr, char* description){
	static_cast<QCommandLinkButton*>(ptr)->setDescription(QString(description));
}

void* QCommandLinkButton_NewQCommandLinkButton(void* parent){
	return new QCommandLinkButton(static_cast<QWidget*>(parent));
}

void* QCommandLinkButton_NewQCommandLinkButton2(char* text, void* parent){
	return new QCommandLinkButton(QString(text), static_cast<QWidget*>(parent));
}

void* QCommandLinkButton_NewQCommandLinkButton3(char* text, char* description, void* parent){
	return new QCommandLinkButton(QString(text), QString(description), static_cast<QWidget*>(parent));
}

void QCommandLinkButton_DestroyQCommandLinkButton(void* ptr){
	static_cast<QCommandLinkButton*>(ptr)->~QCommandLinkButton();
}

