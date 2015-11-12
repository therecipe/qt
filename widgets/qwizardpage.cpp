#include "qwizardpage.h"
#include <QModelIndex>
#include <QPixmap>
#include <QWizard>
#include <QWidget>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QWizardPage>
#include "_cgo_export.h"

class MyQWizardPage: public QWizardPage {
public:
void Signal_CompleteChanged(){callbackQWizardPageCompleteChanged(this->objectName().toUtf8().data());};
};

void QWizardPage_SetSubTitle(void* ptr, char* subTitle){
	static_cast<QWizardPage*>(ptr)->setSubTitle(QString(subTitle));
}

void QWizardPage_SetTitle(void* ptr, char* title){
	static_cast<QWizardPage*>(ptr)->setTitle(QString(title));
}

char* QWizardPage_SubTitle(void* ptr){
	return static_cast<QWizardPage*>(ptr)->subTitle().toUtf8().data();
}

char* QWizardPage_Title(void* ptr){
	return static_cast<QWizardPage*>(ptr)->title().toUtf8().data();
}

void* QWizardPage_NewQWizardPage(void* parent){
	return new QWizardPage(static_cast<QWidget*>(parent));
}

char* QWizardPage_ButtonText(void* ptr, int which){
	return static_cast<QWizardPage*>(ptr)->buttonText(static_cast<QWizard::WizardButton>(which)).toUtf8().data();
}

void QWizardPage_CleanupPage(void* ptr){
	static_cast<QWizardPage*>(ptr)->cleanupPage();
}

void QWizardPage_ConnectCompleteChanged(void* ptr){
	QObject::connect(static_cast<QWizardPage*>(ptr), static_cast<void (QWizardPage::*)()>(&QWizardPage::completeChanged), static_cast<MyQWizardPage*>(ptr), static_cast<void (MyQWizardPage::*)()>(&MyQWizardPage::Signal_CompleteChanged));;
}

void QWizardPage_DisconnectCompleteChanged(void* ptr){
	QObject::disconnect(static_cast<QWizardPage*>(ptr), static_cast<void (QWizardPage::*)()>(&QWizardPage::completeChanged), static_cast<MyQWizardPage*>(ptr), static_cast<void (MyQWizardPage::*)()>(&MyQWizardPage::Signal_CompleteChanged));;
}

void QWizardPage_InitializePage(void* ptr){
	static_cast<QWizardPage*>(ptr)->initializePage();
}

int QWizardPage_IsCommitPage(void* ptr){
	return static_cast<QWizardPage*>(ptr)->isCommitPage();
}

int QWizardPage_IsComplete(void* ptr){
	return static_cast<QWizardPage*>(ptr)->isComplete();
}

int QWizardPage_IsFinalPage(void* ptr){
	return static_cast<QWizardPage*>(ptr)->isFinalPage();
}

int QWizardPage_NextId(void* ptr){
	return static_cast<QWizardPage*>(ptr)->nextId();
}

void QWizardPage_SetButtonText(void* ptr, int which, char* text){
	static_cast<QWizardPage*>(ptr)->setButtonText(static_cast<QWizard::WizardButton>(which), QString(text));
}

void QWizardPage_SetCommitPage(void* ptr, int commitPage){
	static_cast<QWizardPage*>(ptr)->setCommitPage(commitPage != 0);
}

void QWizardPage_SetFinalPage(void* ptr, int finalPage){
	static_cast<QWizardPage*>(ptr)->setFinalPage(finalPage != 0);
}

void QWizardPage_SetPixmap(void* ptr, int which, void* pixmap){
	static_cast<QWizardPage*>(ptr)->setPixmap(static_cast<QWizard::WizardPixmap>(which), *static_cast<QPixmap*>(pixmap));
}

int QWizardPage_ValidatePage(void* ptr){
	return static_cast<QWizardPage*>(ptr)->validatePage();
}

void QWizardPage_DestroyQWizardPage(void* ptr){
	static_cast<QWizardPage*>(ptr)->~QWizardPage();
}

