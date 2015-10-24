#include "qbuttongroup.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QAbstractButton>
#include <QObject>
#include <QButtonGroup>
#include "_cgo_export.h"

class MyQButtonGroup: public QButtonGroup {
public:
void Signal_ButtonClicked(QAbstractButton * button){callbackQButtonGroupButtonClicked(this->objectName().toUtf8().data(), button);};
void Signal_ButtonPressed(QAbstractButton * button){callbackQButtonGroupButtonPressed(this->objectName().toUtf8().data(), button);};
void Signal_ButtonReleased(QAbstractButton * button){callbackQButtonGroupButtonReleased(this->objectName().toUtf8().data(), button);};
void Signal_ButtonToggled(QAbstractButton * button, bool checked){callbackQButtonGroupButtonToggled(this->objectName().toUtf8().data(), button, checked);};
};

QtObjectPtr QButtonGroup_NewQButtonGroup(QtObjectPtr parent){
	return new QButtonGroup(static_cast<QObject*>(parent));
}

void QButtonGroup_AddButton(QtObjectPtr ptr, QtObjectPtr button, int id){
	static_cast<QButtonGroup*>(ptr)->addButton(static_cast<QAbstractButton*>(button), id);
}

QtObjectPtr QButtonGroup_Button(QtObjectPtr ptr, int id){
	return static_cast<QButtonGroup*>(ptr)->button(id);
}

QtObjectPtr QButtonGroup_CheckedButton(QtObjectPtr ptr){
	return static_cast<QButtonGroup*>(ptr)->checkedButton();
}

int QButtonGroup_CheckedId(QtObjectPtr ptr){
	return static_cast<QButtonGroup*>(ptr)->checkedId();
}

int QButtonGroup_Exclusive(QtObjectPtr ptr){
	return static_cast<QButtonGroup*>(ptr)->exclusive();
}

int QButtonGroup_Id(QtObjectPtr ptr, QtObjectPtr button){
	return static_cast<QButtonGroup*>(ptr)->id(static_cast<QAbstractButton*>(button));
}

void QButtonGroup_RemoveButton(QtObjectPtr ptr, QtObjectPtr button){
	static_cast<QButtonGroup*>(ptr)->removeButton(static_cast<QAbstractButton*>(button));
}

void QButtonGroup_SetExclusive(QtObjectPtr ptr, int v){
	static_cast<QButtonGroup*>(ptr)->setExclusive(v != 0);
}

void QButtonGroup_SetId(QtObjectPtr ptr, QtObjectPtr button, int id){
	static_cast<QButtonGroup*>(ptr)->setId(static_cast<QAbstractButton*>(button), id);
}

void QButtonGroup_DestroyQButtonGroup(QtObjectPtr ptr){
	static_cast<QButtonGroup*>(ptr)->~QButtonGroup();
}

void QButtonGroup_ConnectButtonClicked(QtObjectPtr ptr){
	QObject::connect(static_cast<QButtonGroup*>(ptr), static_cast<void (QButtonGroup::*)(QAbstractButton *)>(&QButtonGroup::buttonClicked), static_cast<MyQButtonGroup*>(ptr), static_cast<void (MyQButtonGroup::*)(QAbstractButton *)>(&MyQButtonGroup::Signal_ButtonClicked));;
}

void QButtonGroup_DisconnectButtonClicked(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QButtonGroup*>(ptr), static_cast<void (QButtonGroup::*)(QAbstractButton *)>(&QButtonGroup::buttonClicked), static_cast<MyQButtonGroup*>(ptr), static_cast<void (MyQButtonGroup::*)(QAbstractButton *)>(&MyQButtonGroup::Signal_ButtonClicked));;
}

void QButtonGroup_ConnectButtonPressed(QtObjectPtr ptr){
	QObject::connect(static_cast<QButtonGroup*>(ptr), static_cast<void (QButtonGroup::*)(QAbstractButton *)>(&QButtonGroup::buttonPressed), static_cast<MyQButtonGroup*>(ptr), static_cast<void (MyQButtonGroup::*)(QAbstractButton *)>(&MyQButtonGroup::Signal_ButtonPressed));;
}

void QButtonGroup_DisconnectButtonPressed(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QButtonGroup*>(ptr), static_cast<void (QButtonGroup::*)(QAbstractButton *)>(&QButtonGroup::buttonPressed), static_cast<MyQButtonGroup*>(ptr), static_cast<void (MyQButtonGroup::*)(QAbstractButton *)>(&MyQButtonGroup::Signal_ButtonPressed));;
}

void QButtonGroup_ConnectButtonReleased(QtObjectPtr ptr){
	QObject::connect(static_cast<QButtonGroup*>(ptr), static_cast<void (QButtonGroup::*)(QAbstractButton *)>(&QButtonGroup::buttonReleased), static_cast<MyQButtonGroup*>(ptr), static_cast<void (MyQButtonGroup::*)(QAbstractButton *)>(&MyQButtonGroup::Signal_ButtonReleased));;
}

void QButtonGroup_DisconnectButtonReleased(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QButtonGroup*>(ptr), static_cast<void (QButtonGroup::*)(QAbstractButton *)>(&QButtonGroup::buttonReleased), static_cast<MyQButtonGroup*>(ptr), static_cast<void (MyQButtonGroup::*)(QAbstractButton *)>(&MyQButtonGroup::Signal_ButtonReleased));;
}

void QButtonGroup_ConnectButtonToggled(QtObjectPtr ptr){
	QObject::connect(static_cast<QButtonGroup*>(ptr), static_cast<void (QButtonGroup::*)(QAbstractButton *, bool)>(&QButtonGroup::buttonToggled), static_cast<MyQButtonGroup*>(ptr), static_cast<void (MyQButtonGroup::*)(QAbstractButton *, bool)>(&MyQButtonGroup::Signal_ButtonToggled));;
}

void QButtonGroup_DisconnectButtonToggled(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QButtonGroup*>(ptr), static_cast<void (QButtonGroup::*)(QAbstractButton *, bool)>(&QButtonGroup::buttonToggled), static_cast<MyQButtonGroup*>(ptr), static_cast<void (MyQButtonGroup::*)(QAbstractButton *, bool)>(&MyQButtonGroup::Signal_ButtonToggled));;
}

