#include "qdialogbuttonbox.h"
#include <QUrl>
#include <QModelIndex>
#include <QWidget>
#include <QString>
#include <QVariant>
#include <QObject>
#include <QDial>
#include <QAbstractButton>
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

int QDialogButtonBox_CenterButtons(void* ptr){
	return static_cast<QDialogButtonBox*>(ptr)->centerButtons();
}

int QDialogButtonBox_Orientation(void* ptr){
	return static_cast<QDialogButtonBox*>(ptr)->orientation();
}

void QDialogButtonBox_SetCenterButtons(void* ptr, int center){
	static_cast<QDialogButtonBox*>(ptr)->setCenterButtons(center != 0);
}

void QDialogButtonBox_SetOrientation(void* ptr, int orientation){
	static_cast<QDialogButtonBox*>(ptr)->setOrientation(static_cast<Qt::Orientation>(orientation));
}

void QDialogButtonBox_SetStandardButtons(void* ptr, int buttons){
	static_cast<QDialogButtonBox*>(ptr)->setStandardButtons(static_cast<QDialogButtonBox::StandardButton>(buttons));
}

int QDialogButtonBox_StandardButtons(void* ptr){
	return static_cast<QDialogButtonBox*>(ptr)->standardButtons();
}

void* QDialogButtonBox_NewQDialogButtonBox(void* parent){
	return new QDialogButtonBox(static_cast<QWidget*>(parent));
}

void* QDialogButtonBox_NewQDialogButtonBox2(int orientation, void* parent){
	return new QDialogButtonBox(static_cast<Qt::Orientation>(orientation), static_cast<QWidget*>(parent));
}

void* QDialogButtonBox_NewQDialogButtonBox3(int buttons, void* parent){
	return new QDialogButtonBox(static_cast<QDialogButtonBox::StandardButton>(buttons), static_cast<QWidget*>(parent));
}

void* QDialogButtonBox_NewQDialogButtonBox4(int buttons, int orientation, void* parent){
	return new QDialogButtonBox(static_cast<QDialogButtonBox::StandardButton>(buttons), static_cast<Qt::Orientation>(orientation), static_cast<QWidget*>(parent));
}

void QDialogButtonBox_ConnectAccepted(void* ptr){
	QObject::connect(static_cast<QDialogButtonBox*>(ptr), static_cast<void (QDialogButtonBox::*)()>(&QDialogButtonBox::accepted), static_cast<MyQDialogButtonBox*>(ptr), static_cast<void (MyQDialogButtonBox::*)()>(&MyQDialogButtonBox::Signal_Accepted));;
}

void QDialogButtonBox_DisconnectAccepted(void* ptr){
	QObject::disconnect(static_cast<QDialogButtonBox*>(ptr), static_cast<void (QDialogButtonBox::*)()>(&QDialogButtonBox::accepted), static_cast<MyQDialogButtonBox*>(ptr), static_cast<void (MyQDialogButtonBox::*)()>(&MyQDialogButtonBox::Signal_Accepted));;
}

void* QDialogButtonBox_AddButton3(void* ptr, int button){
	return static_cast<QDialogButtonBox*>(ptr)->addButton(static_cast<QDialogButtonBox::StandardButton>(button));
}

void* QDialogButtonBox_AddButton2(void* ptr, char* text, int role){
	return static_cast<QDialogButtonBox*>(ptr)->addButton(QString(text), static_cast<QDialogButtonBox::ButtonRole>(role));
}

void QDialogButtonBox_AddButton(void* ptr, void* button, int role){
	static_cast<QDialogButtonBox*>(ptr)->addButton(static_cast<QAbstractButton*>(button), static_cast<QDialogButtonBox::ButtonRole>(role));
}

void* QDialogButtonBox_Button(void* ptr, int which){
	return static_cast<QDialogButtonBox*>(ptr)->button(static_cast<QDialogButtonBox::StandardButton>(which));
}

int QDialogButtonBox_ButtonRole(void* ptr, void* button){
	return static_cast<QDialogButtonBox*>(ptr)->buttonRole(static_cast<QAbstractButton*>(button));
}

void QDialogButtonBox_Clear(void* ptr){
	static_cast<QDialogButtonBox*>(ptr)->clear();
}

void QDialogButtonBox_ConnectClicked(void* ptr){
	QObject::connect(static_cast<QDialogButtonBox*>(ptr), static_cast<void (QDialogButtonBox::*)(QAbstractButton *)>(&QDialogButtonBox::clicked), static_cast<MyQDialogButtonBox*>(ptr), static_cast<void (MyQDialogButtonBox::*)(QAbstractButton *)>(&MyQDialogButtonBox::Signal_Clicked));;
}

void QDialogButtonBox_DisconnectClicked(void* ptr){
	QObject::disconnect(static_cast<QDialogButtonBox*>(ptr), static_cast<void (QDialogButtonBox::*)(QAbstractButton *)>(&QDialogButtonBox::clicked), static_cast<MyQDialogButtonBox*>(ptr), static_cast<void (MyQDialogButtonBox::*)(QAbstractButton *)>(&MyQDialogButtonBox::Signal_Clicked));;
}

void QDialogButtonBox_ConnectHelpRequested(void* ptr){
	QObject::connect(static_cast<QDialogButtonBox*>(ptr), static_cast<void (QDialogButtonBox::*)()>(&QDialogButtonBox::helpRequested), static_cast<MyQDialogButtonBox*>(ptr), static_cast<void (MyQDialogButtonBox::*)()>(&MyQDialogButtonBox::Signal_HelpRequested));;
}

void QDialogButtonBox_DisconnectHelpRequested(void* ptr){
	QObject::disconnect(static_cast<QDialogButtonBox*>(ptr), static_cast<void (QDialogButtonBox::*)()>(&QDialogButtonBox::helpRequested), static_cast<MyQDialogButtonBox*>(ptr), static_cast<void (MyQDialogButtonBox::*)()>(&MyQDialogButtonBox::Signal_HelpRequested));;
}

void QDialogButtonBox_ConnectRejected(void* ptr){
	QObject::connect(static_cast<QDialogButtonBox*>(ptr), static_cast<void (QDialogButtonBox::*)()>(&QDialogButtonBox::rejected), static_cast<MyQDialogButtonBox*>(ptr), static_cast<void (MyQDialogButtonBox::*)()>(&MyQDialogButtonBox::Signal_Rejected));;
}

void QDialogButtonBox_DisconnectRejected(void* ptr){
	QObject::disconnect(static_cast<QDialogButtonBox*>(ptr), static_cast<void (QDialogButtonBox::*)()>(&QDialogButtonBox::rejected), static_cast<MyQDialogButtonBox*>(ptr), static_cast<void (MyQDialogButtonBox::*)()>(&MyQDialogButtonBox::Signal_Rejected));;
}

void QDialogButtonBox_RemoveButton(void* ptr, void* button){
	static_cast<QDialogButtonBox*>(ptr)->removeButton(static_cast<QAbstractButton*>(button));
}

int QDialogButtonBox_StandardButton(void* ptr, void* button){
	return static_cast<QDialogButtonBox*>(ptr)->standardButton(static_cast<QAbstractButton*>(button));
}

void QDialogButtonBox_DestroyQDialogButtonBox(void* ptr){
	static_cast<QDialogButtonBox*>(ptr)->~QDialogButtonBox();
}

