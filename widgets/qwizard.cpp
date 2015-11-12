#include "qwizard.h"
#include <QString>
#include <QWidget>
#include <QWizardPage>
#include <QObject>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QPixmap>
#include <QAbstractButton>
#include <QMetaObject>
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

int QWizard_CurrentId(void* ptr){
	return static_cast<QWizard*>(ptr)->currentId();
}

int QWizard_HasVisitedPage(void* ptr, int id){
	return static_cast<QWizard*>(ptr)->hasVisitedPage(id);
}

int QWizard_Options(void* ptr){
	return static_cast<QWizard*>(ptr)->options();
}

void* QWizard_Page(void* ptr, int id){
	return static_cast<QWizard*>(ptr)->page(id);
}

void QWizard_SetOptions(void* ptr, int options){
	static_cast<QWizard*>(ptr)->setOptions(static_cast<QWizard::WizardOption>(options));
}

void QWizard_SetPage(void* ptr, int id, void* page){
	static_cast<QWizard*>(ptr)->setPage(id, static_cast<QWizardPage*>(page));
}

void QWizard_SetStartId(void* ptr, int id){
	static_cast<QWizard*>(ptr)->setStartId(id);
}

void QWizard_SetSubTitleFormat(void* ptr, int format){
	static_cast<QWizard*>(ptr)->setSubTitleFormat(static_cast<Qt::TextFormat>(format));
}

void QWizard_SetTitleFormat(void* ptr, int format){
	static_cast<QWizard*>(ptr)->setTitleFormat(static_cast<Qt::TextFormat>(format));
}

void QWizard_SetWizardStyle(void* ptr, int style){
	static_cast<QWizard*>(ptr)->setWizardStyle(static_cast<QWizard::WizardStyle>(style));
}

int QWizard_StartId(void* ptr){
	return static_cast<QWizard*>(ptr)->startId();
}

int QWizard_SubTitleFormat(void* ptr){
	return static_cast<QWizard*>(ptr)->subTitleFormat();
}

int QWizard_TitleFormat(void* ptr){
	return static_cast<QWizard*>(ptr)->titleFormat();
}

int QWizard_WizardStyle(void* ptr){
	return static_cast<QWizard*>(ptr)->wizardStyle();
}

void* QWizard_NewQWizard(void* parent, int flags){
	return new QWizard(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(flags));
}

int QWizard_AddPage(void* ptr, void* page){
	return static_cast<QWizard*>(ptr)->addPage(static_cast<QWizardPage*>(page));
}

void QWizard_Back(void* ptr){
	QMetaObject::invokeMethod(static_cast<QWizard*>(ptr), "back");
}

void* QWizard_Button(void* ptr, int which){
	return static_cast<QWizard*>(ptr)->button(static_cast<QWizard::WizardButton>(which));
}

char* QWizard_ButtonText(void* ptr, int which){
	return static_cast<QWizard*>(ptr)->buttonText(static_cast<QWizard::WizardButton>(which)).toUtf8().data();
}

void QWizard_ConnectCurrentIdChanged(void* ptr){
	QObject::connect(static_cast<QWizard*>(ptr), static_cast<void (QWizard::*)(int)>(&QWizard::currentIdChanged), static_cast<MyQWizard*>(ptr), static_cast<void (MyQWizard::*)(int)>(&MyQWizard::Signal_CurrentIdChanged));;
}

void QWizard_DisconnectCurrentIdChanged(void* ptr){
	QObject::disconnect(static_cast<QWizard*>(ptr), static_cast<void (QWizard::*)(int)>(&QWizard::currentIdChanged), static_cast<MyQWizard*>(ptr), static_cast<void (MyQWizard::*)(int)>(&MyQWizard::Signal_CurrentIdChanged));;
}

void* QWizard_CurrentPage(void* ptr){
	return static_cast<QWizard*>(ptr)->currentPage();
}

void QWizard_ConnectCustomButtonClicked(void* ptr){
	QObject::connect(static_cast<QWizard*>(ptr), static_cast<void (QWizard::*)(int)>(&QWizard::customButtonClicked), static_cast<MyQWizard*>(ptr), static_cast<void (MyQWizard::*)(int)>(&MyQWizard::Signal_CustomButtonClicked));;
}

void QWizard_DisconnectCustomButtonClicked(void* ptr){
	QObject::disconnect(static_cast<QWizard*>(ptr), static_cast<void (QWizard::*)(int)>(&QWizard::customButtonClicked), static_cast<MyQWizard*>(ptr), static_cast<void (MyQWizard::*)(int)>(&MyQWizard::Signal_CustomButtonClicked));;
}

void* QWizard_Field(void* ptr, char* name){
	return new QVariant(static_cast<QWizard*>(ptr)->field(QString(name)));
}

void QWizard_ConnectHelpRequested(void* ptr){
	QObject::connect(static_cast<QWizard*>(ptr), static_cast<void (QWizard::*)()>(&QWizard::helpRequested), static_cast<MyQWizard*>(ptr), static_cast<void (MyQWizard::*)()>(&MyQWizard::Signal_HelpRequested));;
}

void QWizard_DisconnectHelpRequested(void* ptr){
	QObject::disconnect(static_cast<QWizard*>(ptr), static_cast<void (QWizard::*)()>(&QWizard::helpRequested), static_cast<MyQWizard*>(ptr), static_cast<void (MyQWizard::*)()>(&MyQWizard::Signal_HelpRequested));;
}

void QWizard_Next(void* ptr){
	QMetaObject::invokeMethod(static_cast<QWizard*>(ptr), "next");
}

int QWizard_NextId(void* ptr){
	return static_cast<QWizard*>(ptr)->nextId();
}

void QWizard_ConnectPageAdded(void* ptr){
	QObject::connect(static_cast<QWizard*>(ptr), static_cast<void (QWizard::*)(int)>(&QWizard::pageAdded), static_cast<MyQWizard*>(ptr), static_cast<void (MyQWizard::*)(int)>(&MyQWizard::Signal_PageAdded));;
}

void QWizard_DisconnectPageAdded(void* ptr){
	QObject::disconnect(static_cast<QWizard*>(ptr), static_cast<void (QWizard::*)(int)>(&QWizard::pageAdded), static_cast<MyQWizard*>(ptr), static_cast<void (MyQWizard::*)(int)>(&MyQWizard::Signal_PageAdded));;
}

void QWizard_ConnectPageRemoved(void* ptr){
	QObject::connect(static_cast<QWizard*>(ptr), static_cast<void (QWizard::*)(int)>(&QWizard::pageRemoved), static_cast<MyQWizard*>(ptr), static_cast<void (MyQWizard::*)(int)>(&MyQWizard::Signal_PageRemoved));;
}

void QWizard_DisconnectPageRemoved(void* ptr){
	QObject::disconnect(static_cast<QWizard*>(ptr), static_cast<void (QWizard::*)(int)>(&QWizard::pageRemoved), static_cast<MyQWizard*>(ptr), static_cast<void (MyQWizard::*)(int)>(&MyQWizard::Signal_PageRemoved));;
}

void QWizard_RemovePage(void* ptr, int id){
	static_cast<QWizard*>(ptr)->removePage(id);
}

void QWizard_Restart(void* ptr){
	QMetaObject::invokeMethod(static_cast<QWizard*>(ptr), "restart");
}

void QWizard_SetButton(void* ptr, int which, void* button){
	static_cast<QWizard*>(ptr)->setButton(static_cast<QWizard::WizardButton>(which), static_cast<QAbstractButton*>(button));
}

void QWizard_SetButtonText(void* ptr, int which, char* text){
	static_cast<QWizard*>(ptr)->setButtonText(static_cast<QWizard::WizardButton>(which), QString(text));
}

void QWizard_SetDefaultProperty(void* ptr, char* className, char* property, char* changedSignal){
	static_cast<QWizard*>(ptr)->setDefaultProperty(const_cast<const char*>(className), const_cast<const char*>(property), const_cast<const char*>(changedSignal));
}

void QWizard_SetField(void* ptr, char* name, void* value){
	static_cast<QWizard*>(ptr)->setField(QString(name), *static_cast<QVariant*>(value));
}

void QWizard_SetOption(void* ptr, int option, int on){
	static_cast<QWizard*>(ptr)->setOption(static_cast<QWizard::WizardOption>(option), on != 0);
}

void QWizard_SetPixmap(void* ptr, int which, void* pixmap){
	static_cast<QWizard*>(ptr)->setPixmap(static_cast<QWizard::WizardPixmap>(which), *static_cast<QPixmap*>(pixmap));
}

void QWizard_SetSideWidget(void* ptr, void* widget){
	static_cast<QWizard*>(ptr)->setSideWidget(static_cast<QWidget*>(widget));
}

void QWizard_SetVisible(void* ptr, int visible){
	static_cast<QWizard*>(ptr)->setVisible(visible != 0);
}

void* QWizard_SideWidget(void* ptr){
	return static_cast<QWizard*>(ptr)->sideWidget();
}

int QWizard_TestOption(void* ptr, int option){
	return static_cast<QWizard*>(ptr)->testOption(static_cast<QWizard::WizardOption>(option));
}

int QWizard_ValidateCurrentPage(void* ptr){
	return static_cast<QWizard*>(ptr)->validateCurrentPage();
}

void QWizard_DestroyQWizard(void* ptr){
	static_cast<QWizard*>(ptr)->~QWizard();
}

