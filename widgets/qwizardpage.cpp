#include "qwizardpage.h"
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QWizard>
#include <QPixmap>
#include <QWidget>
#include <QString>
#include <QVariant>
#include <QWizardPage>
#include "_cgo_export.h"

class MyQWizardPage: public QWizardPage {
public:
void Signal_CompleteChanged(){callbackQWizardPageCompleteChanged(this->objectName().toUtf8().data());};
};

void QWizardPage_SetSubTitle(QtObjectPtr ptr, char* subTitle){
	static_cast<QWizardPage*>(ptr)->setSubTitle(QString(subTitle));
}

void QWizardPage_SetTitle(QtObjectPtr ptr, char* title){
	static_cast<QWizardPage*>(ptr)->setTitle(QString(title));
}

char* QWizardPage_SubTitle(QtObjectPtr ptr){
	return static_cast<QWizardPage*>(ptr)->subTitle().toUtf8().data();
}

char* QWizardPage_Title(QtObjectPtr ptr){
	return static_cast<QWizardPage*>(ptr)->title().toUtf8().data();
}

QtObjectPtr QWizardPage_NewQWizardPage(QtObjectPtr parent){
	return new QWizardPage(static_cast<QWidget*>(parent));
}

char* QWizardPage_ButtonText(QtObjectPtr ptr, int which){
	return static_cast<QWizardPage*>(ptr)->buttonText(static_cast<QWizard::WizardButton>(which)).toUtf8().data();
}

void QWizardPage_CleanupPage(QtObjectPtr ptr){
	static_cast<QWizardPage*>(ptr)->cleanupPage();
}

void QWizardPage_ConnectCompleteChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QWizardPage*>(ptr), static_cast<void (QWizardPage::*)()>(&QWizardPage::completeChanged), static_cast<MyQWizardPage*>(ptr), static_cast<void (MyQWizardPage::*)()>(&MyQWizardPage::Signal_CompleteChanged));;
}

void QWizardPage_DisconnectCompleteChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QWizardPage*>(ptr), static_cast<void (QWizardPage::*)()>(&QWizardPage::completeChanged), static_cast<MyQWizardPage*>(ptr), static_cast<void (MyQWizardPage::*)()>(&MyQWizardPage::Signal_CompleteChanged));;
}

void QWizardPage_InitializePage(QtObjectPtr ptr){
	static_cast<QWizardPage*>(ptr)->initializePage();
}

int QWizardPage_IsCommitPage(QtObjectPtr ptr){
	return static_cast<QWizardPage*>(ptr)->isCommitPage();
}

int QWizardPage_IsComplete(QtObjectPtr ptr){
	return static_cast<QWizardPage*>(ptr)->isComplete();
}

int QWizardPage_IsFinalPage(QtObjectPtr ptr){
	return static_cast<QWizardPage*>(ptr)->isFinalPage();
}

int QWizardPage_NextId(QtObjectPtr ptr){
	return static_cast<QWizardPage*>(ptr)->nextId();
}

void QWizardPage_SetButtonText(QtObjectPtr ptr, int which, char* text){
	static_cast<QWizardPage*>(ptr)->setButtonText(static_cast<QWizard::WizardButton>(which), QString(text));
}

void QWizardPage_SetCommitPage(QtObjectPtr ptr, int commitPage){
	static_cast<QWizardPage*>(ptr)->setCommitPage(commitPage != 0);
}

void QWizardPage_SetFinalPage(QtObjectPtr ptr, int finalPage){
	static_cast<QWizardPage*>(ptr)->setFinalPage(finalPage != 0);
}

void QWizardPage_SetPixmap(QtObjectPtr ptr, int which, QtObjectPtr pixmap){
	static_cast<QWizardPage*>(ptr)->setPixmap(static_cast<QWizard::WizardPixmap>(which), *static_cast<QPixmap*>(pixmap));
}

int QWizardPage_ValidatePage(QtObjectPtr ptr){
	return static_cast<QWizardPage*>(ptr)->validatePage();
}

void QWizardPage_DestroyQWizardPage(QtObjectPtr ptr){
	static_cast<QWizardPage*>(ptr)->~QWizardPage();
}

