#include "qbuttongroup.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QAbstractButton>
#include <QButtonGroup>
#include "_cgo_export.h"

class MyQButtonGroup: public QButtonGroup {
public:
void Signal_ButtonClicked(QAbstractButton * button){callbackQButtonGroupButtonClicked(this->objectName().toUtf8().data(), button);};
void Signal_ButtonPressed(QAbstractButton * button){callbackQButtonGroupButtonPressed(this->objectName().toUtf8().data(), button);};
void Signal_ButtonReleased(QAbstractButton * button){callbackQButtonGroupButtonReleased(this->objectName().toUtf8().data(), button);};
void Signal_ButtonToggled(QAbstractButton * button, bool checked){callbackQButtonGroupButtonToggled(this->objectName().toUtf8().data(), button, checked);};
};

void* QButtonGroup_NewQButtonGroup(void* parent){
	return new QButtonGroup(static_cast<QObject*>(parent));
}

void QButtonGroup_AddButton(void* ptr, void* button, int id){
	static_cast<QButtonGroup*>(ptr)->addButton(static_cast<QAbstractButton*>(button), id);
}

void* QButtonGroup_Button(void* ptr, int id){
	return static_cast<QButtonGroup*>(ptr)->button(id);
}

void* QButtonGroup_CheckedButton(void* ptr){
	return static_cast<QButtonGroup*>(ptr)->checkedButton();
}

int QButtonGroup_CheckedId(void* ptr){
	return static_cast<QButtonGroup*>(ptr)->checkedId();
}

int QButtonGroup_Exclusive(void* ptr){
	return static_cast<QButtonGroup*>(ptr)->exclusive();
}

int QButtonGroup_Id(void* ptr, void* button){
	return static_cast<QButtonGroup*>(ptr)->id(static_cast<QAbstractButton*>(button));
}

void QButtonGroup_RemoveButton(void* ptr, void* button){
	static_cast<QButtonGroup*>(ptr)->removeButton(static_cast<QAbstractButton*>(button));
}

void QButtonGroup_SetExclusive(void* ptr, int v){
	static_cast<QButtonGroup*>(ptr)->setExclusive(v != 0);
}

void QButtonGroup_SetId(void* ptr, void* button, int id){
	static_cast<QButtonGroup*>(ptr)->setId(static_cast<QAbstractButton*>(button), id);
}

void QButtonGroup_DestroyQButtonGroup(void* ptr){
	static_cast<QButtonGroup*>(ptr)->~QButtonGroup();
}

void QButtonGroup_ConnectButtonClicked(void* ptr){
	QObject::connect(static_cast<QButtonGroup*>(ptr), static_cast<void (QButtonGroup::*)(QAbstractButton *)>(&QButtonGroup::buttonClicked), static_cast<MyQButtonGroup*>(ptr), static_cast<void (MyQButtonGroup::*)(QAbstractButton *)>(&MyQButtonGroup::Signal_ButtonClicked));;
}

void QButtonGroup_DisconnectButtonClicked(void* ptr){
	QObject::disconnect(static_cast<QButtonGroup*>(ptr), static_cast<void (QButtonGroup::*)(QAbstractButton *)>(&QButtonGroup::buttonClicked), static_cast<MyQButtonGroup*>(ptr), static_cast<void (MyQButtonGroup::*)(QAbstractButton *)>(&MyQButtonGroup::Signal_ButtonClicked));;
}

void QButtonGroup_ConnectButtonPressed(void* ptr){
	QObject::connect(static_cast<QButtonGroup*>(ptr), static_cast<void (QButtonGroup::*)(QAbstractButton *)>(&QButtonGroup::buttonPressed), static_cast<MyQButtonGroup*>(ptr), static_cast<void (MyQButtonGroup::*)(QAbstractButton *)>(&MyQButtonGroup::Signal_ButtonPressed));;
}

void QButtonGroup_DisconnectButtonPressed(void* ptr){
	QObject::disconnect(static_cast<QButtonGroup*>(ptr), static_cast<void (QButtonGroup::*)(QAbstractButton *)>(&QButtonGroup::buttonPressed), static_cast<MyQButtonGroup*>(ptr), static_cast<void (MyQButtonGroup::*)(QAbstractButton *)>(&MyQButtonGroup::Signal_ButtonPressed));;
}

void QButtonGroup_ConnectButtonReleased(void* ptr){
	QObject::connect(static_cast<QButtonGroup*>(ptr), static_cast<void (QButtonGroup::*)(QAbstractButton *)>(&QButtonGroup::buttonReleased), static_cast<MyQButtonGroup*>(ptr), static_cast<void (MyQButtonGroup::*)(QAbstractButton *)>(&MyQButtonGroup::Signal_ButtonReleased));;
}

void QButtonGroup_DisconnectButtonReleased(void* ptr){
	QObject::disconnect(static_cast<QButtonGroup*>(ptr), static_cast<void (QButtonGroup::*)(QAbstractButton *)>(&QButtonGroup::buttonReleased), static_cast<MyQButtonGroup*>(ptr), static_cast<void (MyQButtonGroup::*)(QAbstractButton *)>(&MyQButtonGroup::Signal_ButtonReleased));;
}

void QButtonGroup_ConnectButtonToggled(void* ptr){
	QObject::connect(static_cast<QButtonGroup*>(ptr), static_cast<void (QButtonGroup::*)(QAbstractButton *, bool)>(&QButtonGroup::buttonToggled), static_cast<MyQButtonGroup*>(ptr), static_cast<void (MyQButtonGroup::*)(QAbstractButton *, bool)>(&MyQButtonGroup::Signal_ButtonToggled));;
}

void QButtonGroup_DisconnectButtonToggled(void* ptr){
	QObject::disconnect(static_cast<QButtonGroup*>(ptr), static_cast<void (QButtonGroup::*)(QAbstractButton *, bool)>(&QButtonGroup::buttonToggled), static_cast<MyQButtonGroup*>(ptr), static_cast<void (MyQButtonGroup::*)(QAbstractButton *, bool)>(&MyQButtonGroup::Signal_ButtonToggled));;
}

