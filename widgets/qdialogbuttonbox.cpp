#include "qdialogbuttonbox.h"
#include <QString>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QAbstractButton>
#include <QWidget>
#include <QVariant>
#include <QDial>
#include <QDialog>
#include <QDialogButtonBox>
#include "_cgo_export.h"

class MyQDialogButtonBox: public QDialogButtonBox {
public:
void Signal_Accepted(){callbackQDialogButtonBoxAccepted(this->objectName().toUtf8().data());};
void Signal_Clicked(QAbstractButton * button){callbackQDialogButtonBoxClicked(this->objectName().toUtf8().data(), button);};
void Signal_HelpRequested(){callbackQDialogButtonBoxHelpRequested(this->objectName().toUtf8().data());};
void Signal_Rejected(){callbackQDialogButtonBoxRejected(this->objectName().toUtf8().data());};
};

int QDialogButtonBox_CenterButtons(QtObjectPtr ptr){
	return static_cast<QDialogButtonBox*>(ptr)->centerButtons();
}

int QDialogButtonBox_Orientation(QtObjectPtr ptr){
	return static_cast<QDialogButtonBox*>(ptr)->orientation();
}

void QDialogButtonBox_SetCenterButtons(QtObjectPtr ptr, int center){
	static_cast<QDialogButtonBox*>(ptr)->setCenterButtons(center != 0);
}

void QDialogButtonBox_SetOrientation(QtObjectPtr ptr, int orientation){
	static_cast<QDialogButtonBox*>(ptr)->setOrientation(static_cast<Qt::Orientation>(orientation));
}

void QDialogButtonBox_SetStandardButtons(QtObjectPtr ptr, int buttons){
	static_cast<QDialogButtonBox*>(ptr)->setStandardButtons(static_cast<QDialogButtonBox::StandardButton>(buttons));
}

int QDialogButtonBox_StandardButtons(QtObjectPtr ptr){
	return static_cast<QDialogButtonBox*>(ptr)->standardButtons();
}

QtObjectPtr QDialogButtonBox_NewQDialogButtonBox(QtObjectPtr parent){
	return new QDialogButtonBox(static_cast<QWidget*>(parent));
}

QtObjectPtr QDialogButtonBox_NewQDialogButtonBox2(int orientation, QtObjectPtr parent){
	return new QDialogButtonBox(static_cast<Qt::Orientation>(orientation), static_cast<QWidget*>(parent));
}

QtObjectPtr QDialogButtonBox_NewQDialogButtonBox3(int buttons, QtObjectPtr parent){
	return new QDialogButtonBox(static_cast<QDialogButtonBox::StandardButton>(buttons), static_cast<QWidget*>(parent));
}

QtObjectPtr QDialogButtonBox_NewQDialogButtonBox4(int buttons, int orientation, QtObjectPtr parent){
	return new QDialogButtonBox(static_cast<QDialogButtonBox::StandardButton>(buttons), static_cast<Qt::Orientation>(orientation), static_cast<QWidget*>(parent));
}

void QDialogButtonBox_ConnectAccepted(QtObjectPtr ptr){
	QObject::connect(static_cast<QDialogButtonBox*>(ptr), static_cast<void (QDialogButtonBox::*)()>(&QDialogButtonBox::accepted), static_cast<MyQDialogButtonBox*>(ptr), static_cast<void (MyQDialogButtonBox::*)()>(&MyQDialogButtonBox::Signal_Accepted));;
}

void QDialogButtonBox_DisconnectAccepted(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QDialogButtonBox*>(ptr), static_cast<void (QDialogButtonBox::*)()>(&QDialogButtonBox::accepted), static_cast<MyQDialogButtonBox*>(ptr), static_cast<void (MyQDialogButtonBox::*)()>(&MyQDialogButtonBox::Signal_Accepted));;
}

QtObjectPtr QDialogButtonBox_AddButton3(QtObjectPtr ptr, int button){
	return static_cast<QDialogButtonBox*>(ptr)->addButton(static_cast<QDialogButtonBox::StandardButton>(button));
}

QtObjectPtr QDialogButtonBox_AddButton2(QtObjectPtr ptr, char* text, int role){
	return static_cast<QDialogButtonBox*>(ptr)->addButton(QString(text), static_cast<QDialogButtonBox::ButtonRole>(role));
}

void QDialogButtonBox_AddButton(QtObjectPtr ptr, QtObjectPtr button, int role){
	static_cast<QDialogButtonBox*>(ptr)->addButton(static_cast<QAbstractButton*>(button), static_cast<QDialogButtonBox::ButtonRole>(role));
}

QtObjectPtr QDialogButtonBox_Button(QtObjectPtr ptr, int which){
	return static_cast<QDialogButtonBox*>(ptr)->button(static_cast<QDialogButtonBox::StandardButton>(which));
}

int QDialogButtonBox_ButtonRole(QtObjectPtr ptr, QtObjectPtr button){
	return static_cast<QDialogButtonBox*>(ptr)->buttonRole(static_cast<QAbstractButton*>(button));
}

void QDialogButtonBox_Clear(QtObjectPtr ptr){
	static_cast<QDialogButtonBox*>(ptr)->clear();
}

void QDialogButtonBox_ConnectClicked(QtObjectPtr ptr){
	QObject::connect(static_cast<QDialogButtonBox*>(ptr), static_cast<void (QDialogButtonBox::*)(QAbstractButton *)>(&QDialogButtonBox::clicked), static_cast<MyQDialogButtonBox*>(ptr), static_cast<void (MyQDialogButtonBox::*)(QAbstractButton *)>(&MyQDialogButtonBox::Signal_Clicked));;
}

void QDialogButtonBox_DisconnectClicked(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QDialogButtonBox*>(ptr), static_cast<void (QDialogButtonBox::*)(QAbstractButton *)>(&QDialogButtonBox::clicked), static_cast<MyQDialogButtonBox*>(ptr), static_cast<void (MyQDialogButtonBox::*)(QAbstractButton *)>(&MyQDialogButtonBox::Signal_Clicked));;
}

void QDialogButtonBox_ConnectHelpRequested(QtObjectPtr ptr){
	QObject::connect(static_cast<QDialogButtonBox*>(ptr), static_cast<void (QDialogButtonBox::*)()>(&QDialogButtonBox::helpRequested), static_cast<MyQDialogButtonBox*>(ptr), static_cast<void (MyQDialogButtonBox::*)()>(&MyQDialogButtonBox::Signal_HelpRequested));;
}

void QDialogButtonBox_DisconnectHelpRequested(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QDialogButtonBox*>(ptr), static_cast<void (QDialogButtonBox::*)()>(&QDialogButtonBox::helpRequested), static_cast<MyQDialogButtonBox*>(ptr), static_cast<void (MyQDialogButtonBox::*)()>(&MyQDialogButtonBox::Signal_HelpRequested));;
}

void QDialogButtonBox_ConnectRejected(QtObjectPtr ptr){
	QObject::connect(static_cast<QDialogButtonBox*>(ptr), static_cast<void (QDialogButtonBox::*)()>(&QDialogButtonBox::rejected), static_cast<MyQDialogButtonBox*>(ptr), static_cast<void (MyQDialogButtonBox::*)()>(&MyQDialogButtonBox::Signal_Rejected));;
}

void QDialogButtonBox_DisconnectRejected(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QDialogButtonBox*>(ptr), static_cast<void (QDialogButtonBox::*)()>(&QDialogButtonBox::rejected), static_cast<MyQDialogButtonBox*>(ptr), static_cast<void (MyQDialogButtonBox::*)()>(&MyQDialogButtonBox::Signal_Rejected));;
}

void QDialogButtonBox_RemoveButton(QtObjectPtr ptr, QtObjectPtr button){
	static_cast<QDialogButtonBox*>(ptr)->removeButton(static_cast<QAbstractButton*>(button));
}

int QDialogButtonBox_StandardButton(QtObjectPtr ptr, QtObjectPtr button){
	return static_cast<QDialogButtonBox*>(ptr)->standardButton(static_cast<QAbstractButton*>(button));
}

void QDialogButtonBox_DestroyQDialogButtonBox(QtObjectPtr ptr){
	static_cast<QDialogButtonBox*>(ptr)->~QDialogButtonBox();
}

