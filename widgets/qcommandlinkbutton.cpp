#include "qcommandlinkbutton.h"
#include <QUrl>
#include <QModelIndex>
#include <QWidget>
#include <QString>
#include <QVariant>
#include <QCommandLinkButton>
#include "_cgo_export.h"

class MyQCommandLinkButton: public QCommandLinkButton {
public:
};

char* QCommandLinkButton_Description(QtObjectPtr ptr){
	return static_cast<QCommandLinkButton*>(ptr)->description().toUtf8().data();
}

void QCommandLinkButton_SetDescription(QtObjectPtr ptr, char* description){
	static_cast<QCommandLinkButton*>(ptr)->setDescription(QString(description));
}

QtObjectPtr QCommandLinkButton_NewQCommandLinkButton(QtObjectPtr parent){
	return new QCommandLinkButton(static_cast<QWidget*>(parent));
}

QtObjectPtr QCommandLinkButton_NewQCommandLinkButton2(char* text, QtObjectPtr parent){
	return new QCommandLinkButton(QString(text), static_cast<QWidget*>(parent));
}

QtObjectPtr QCommandLinkButton_NewQCommandLinkButton3(char* text, char* description, QtObjectPtr parent){
	return new QCommandLinkButton(QString(text), QString(description), static_cast<QWidget*>(parent));
}

void QCommandLinkButton_DestroyQCommandLinkButton(QtObjectPtr ptr){
	static_cast<QCommandLinkButton*>(ptr)->~QCommandLinkButton();
}

