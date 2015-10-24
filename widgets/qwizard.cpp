#include "qwizard.h"
#include <QObject>
#include <QAbstractButton>
#include <QWizardPage>
#include <QWidget>
#include <QString>
#include <QUrl>
#include <QMetaObject>
#include <QPixmap>
#include <QVariant>
#include <QModelIndex>
#include <QWizard>
#include "_cgo_export.h"

class MyQWizard: public QWizard {
public:
void Signal_CurrentIdChanged(int id){callbackQWizardCurrentIdChanged(this->objectName().toUtf8().data(), id);};
void Signal_CustomButtonClicked(int which){callbackQWizardCustomButtonClicked(this->objectName().toUtf8().data(), which);};
void Signal_HelpRequested(){callbackQWizardHelpRequested(this->objectName().toUtf8().data());};
void Signal_PageAdded(int id){callbackQWizardPageAdded(this->objectName().toUtf8().data(), id);};
void Signal_PageRemoved(int id){callbackQWizardPageRemoved(this->objectName().toUtf8().data(), id);};
};

int QWizard_CurrentId(QtObjectPtr ptr){
	return static_cast<QWizard*>(ptr)->currentId();
}

int QWizard_HasVisitedPage(QtObjectPtr ptr, int id){
	return static_cast<QWizard*>(ptr)->hasVisitedPage(id);
}

int QWizard_Options(QtObjectPtr ptr){
	return static_cast<QWizard*>(ptr)->options();
}

QtObjectPtr QWizard_Page(QtObjectPtr ptr, int id){
	return static_cast<QWizard*>(ptr)->page(id);
}

void QWizard_SetOptions(QtObjectPtr ptr, int options){
	static_cast<QWizard*>(ptr)->setOptions(static_cast<QWizard::WizardOption>(options));
}

void QWizard_SetPage(QtObjectPtr ptr, int id, QtObjectPtr page){
	static_cast<QWizard*>(ptr)->setPage(id, static_cast<QWizardPage*>(page));
}

void QWizard_SetStartId(QtObjectPtr ptr, int id){
	static_cast<QWizard*>(ptr)->setStartId(id);
}

void QWizard_SetSubTitleFormat(QtObjectPtr ptr, int format){
	static_cast<QWizard*>(ptr)->setSubTitleFormat(static_cast<Qt::TextFormat>(format));
}

void QWizard_SetTitleFormat(QtObjectPtr ptr, int format){
	static_cast<QWizard*>(ptr)->setTitleFormat(static_cast<Qt::TextFormat>(format));
}

void QWizard_SetWizardStyle(QtObjectPtr ptr, int style){
	static_cast<QWizard*>(ptr)->setWizardStyle(static_cast<QWizard::WizardStyle>(style));
}

int QWizard_StartId(QtObjectPtr ptr){
	return static_cast<QWizard*>(ptr)->startId();
}

int QWizard_SubTitleFormat(QtObjectPtr ptr){
	return static_cast<QWizard*>(ptr)->subTitleFormat();
}

int QWizard_TitleFormat(QtObjectPtr ptr){
	return static_cast<QWizard*>(ptr)->titleFormat();
}

int QWizard_WizardStyle(QtObjectPtr ptr){
	return static_cast<QWizard*>(ptr)->wizardStyle();
}

QtObjectPtr QWizard_NewQWizard(QtObjectPtr parent, int flags){
	return new QWizard(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(flags));
}

int QWizard_AddPage(QtObjectPtr ptr, QtObjectPtr page){
	return static_cast<QWizard*>(ptr)->addPage(static_cast<QWizardPage*>(page));
}

void QWizard_Back(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QWizard*>(ptr), "back");
}

QtObjectPtr QWizard_Button(QtObjectPtr ptr, int which){
	return static_cast<QWizard*>(ptr)->button(static_cast<QWizard::WizardButton>(which));
}

char* QWizard_ButtonText(QtObjectPtr ptr, int which){
	return static_cast<QWizard*>(ptr)->buttonText(static_cast<QWizard::WizardButton>(which)).toUtf8().data();
}

void QWizard_ConnectCurrentIdChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QWizard*>(ptr), static_cast<void (QWizard::*)(int)>(&QWizard::currentIdChanged), static_cast<MyQWizard*>(ptr), static_cast<void (MyQWizard::*)(int)>(&MyQWizard::Signal_CurrentIdChanged));;
}

void QWizard_DisconnectCurrentIdChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QWizard*>(ptr), static_cast<void (QWizard::*)(int)>(&QWizard::currentIdChanged), static_cast<MyQWizard*>(ptr), static_cast<void (MyQWizard::*)(int)>(&MyQWizard::Signal_CurrentIdChanged));;
}

QtObjectPtr QWizard_CurrentPage(QtObjectPtr ptr){
	return static_cast<QWizard*>(ptr)->currentPage();
}

void QWizard_ConnectCustomButtonClicked(QtObjectPtr ptr){
	QObject::connect(static_cast<QWizard*>(ptr), static_cast<void (QWizard::*)(int)>(&QWizard::customButtonClicked), static_cast<MyQWizard*>(ptr), static_cast<void (MyQWizard::*)(int)>(&MyQWizard::Signal_CustomButtonClicked));;
}

void QWizard_DisconnectCustomButtonClicked(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QWizard*>(ptr), static_cast<void (QWizard::*)(int)>(&QWizard::customButtonClicked), static_cast<MyQWizard*>(ptr), static_cast<void (MyQWizard::*)(int)>(&MyQWizard::Signal_CustomButtonClicked));;
}

char* QWizard_Field(QtObjectPtr ptr, char* name){
	return static_cast<QWizard*>(ptr)->field(QString(name)).toString().toUtf8().data();
}

void QWizard_ConnectHelpRequested(QtObjectPtr ptr){
	QObject::connect(static_cast<QWizard*>(ptr), static_cast<void (QWizard::*)()>(&QWizard::helpRequested), static_cast<MyQWizard*>(ptr), static_cast<void (MyQWizard::*)()>(&MyQWizard::Signal_HelpRequested));;
}

void QWizard_DisconnectHelpRequested(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QWizard*>(ptr), static_cast<void (QWizard::*)()>(&QWizard::helpRequested), static_cast<MyQWizard*>(ptr), static_cast<void (MyQWizard::*)()>(&MyQWizard::Signal_HelpRequested));;
}

void QWizard_Next(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QWizard*>(ptr), "next");
}

int QWizard_NextId(QtObjectPtr ptr){
	return static_cast<QWizard*>(ptr)->nextId();
}

void QWizard_ConnectPageAdded(QtObjectPtr ptr){
	QObject::connect(static_cast<QWizard*>(ptr), static_cast<void (QWizard::*)(int)>(&QWizard::pageAdded), static_cast<MyQWizard*>(ptr), static_cast<void (MyQWizard::*)(int)>(&MyQWizard::Signal_PageAdded));;
}

void QWizard_DisconnectPageAdded(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QWizard*>(ptr), static_cast<void (QWizard::*)(int)>(&QWizard::pageAdded), static_cast<MyQWizard*>(ptr), static_cast<void (MyQWizard::*)(int)>(&MyQWizard::Signal_PageAdded));;
}

void QWizard_ConnectPageRemoved(QtObjectPtr ptr){
	QObject::connect(static_cast<QWizard*>(ptr), static_cast<void (QWizard::*)(int)>(&QWizard::pageRemoved), static_cast<MyQWizard*>(ptr), static_cast<void (MyQWizard::*)(int)>(&MyQWizard::Signal_PageRemoved));;
}

void QWizard_DisconnectPageRemoved(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QWizard*>(ptr), static_cast<void (QWizard::*)(int)>(&QWizard::pageRemoved), static_cast<MyQWizard*>(ptr), static_cast<void (MyQWizard::*)(int)>(&MyQWizard::Signal_PageRemoved));;
}

void QWizard_RemovePage(QtObjectPtr ptr, int id){
	static_cast<QWizard*>(ptr)->removePage(id);
}

void QWizard_Restart(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QWizard*>(ptr), "restart");
}

void QWizard_SetButton(QtObjectPtr ptr, int which, QtObjectPtr button){
	static_cast<QWizard*>(ptr)->setButton(static_cast<QWizard::WizardButton>(which), static_cast<QAbstractButton*>(button));
}

void QWizard_SetButtonText(QtObjectPtr ptr, int which, char* text){
	static_cast<QWizard*>(ptr)->setButtonText(static_cast<QWizard::WizardButton>(which), QString(text));
}

void QWizard_SetDefaultProperty(QtObjectPtr ptr, char* className, char* property, char* changedSignal){
	static_cast<QWizard*>(ptr)->setDefaultProperty(const_cast<const char*>(className), const_cast<const char*>(property), const_cast<const char*>(changedSignal));
}

void QWizard_SetField(QtObjectPtr ptr, char* name, char* value){
	static_cast<QWizard*>(ptr)->setField(QString(name), QVariant(value));
}

void QWizard_SetOption(QtObjectPtr ptr, int option, int on){
	static_cast<QWizard*>(ptr)->setOption(static_cast<QWizard::WizardOption>(option), on != 0);
}

void QWizard_SetPixmap(QtObjectPtr ptr, int which, QtObjectPtr pixmap){
	static_cast<QWizard*>(ptr)->setPixmap(static_cast<QWizard::WizardPixmap>(which), *static_cast<QPixmap*>(pixmap));
}

void QWizard_SetSideWidget(QtObjectPtr ptr, QtObjectPtr widget){
	static_cast<QWizard*>(ptr)->setSideWidget(static_cast<QWidget*>(widget));
}

void QWizard_SetVisible(QtObjectPtr ptr, int visible){
	static_cast<QWizard*>(ptr)->setVisible(visible != 0);
}

QtObjectPtr QWizard_SideWidget(QtObjectPtr ptr){
	return static_cast<QWizard*>(ptr)->sideWidget();
}

int QWizard_TestOption(QtObjectPtr ptr, int option){
	return static_cast<QWizard*>(ptr)->testOption(static_cast<QWizard::WizardOption>(option));
}

int QWizard_ValidateCurrentPage(QtObjectPtr ptr){
	return static_cast<QWizard*>(ptr)->validateCurrentPage();
}

void QWizard_DestroyQWizard(QtObjectPtr ptr){
	static_cast<QWizard*>(ptr)->~QWizard();
}

