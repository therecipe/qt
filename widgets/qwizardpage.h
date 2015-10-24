#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QWizardPage_SetSubTitle(QtObjectPtr ptr, char* subTitle);
void QWizardPage_SetTitle(QtObjectPtr ptr, char* title);
char* QWizardPage_SubTitle(QtObjectPtr ptr);
char* QWizardPage_Title(QtObjectPtr ptr);
QtObjectPtr QWizardPage_NewQWizardPage(QtObjectPtr parent);
char* QWizardPage_ButtonText(QtObjectPtr ptr, int which);
void QWizardPage_CleanupPage(QtObjectPtr ptr);
void QWizardPage_ConnectCompleteChanged(QtObjectPtr ptr);
void QWizardPage_DisconnectCompleteChanged(QtObjectPtr ptr);
void QWizardPage_InitializePage(QtObjectPtr ptr);
int QWizardPage_IsCommitPage(QtObjectPtr ptr);
int QWizardPage_IsComplete(QtObjectPtr ptr);
int QWizardPage_IsFinalPage(QtObjectPtr ptr);
int QWizardPage_NextId(QtObjectPtr ptr);
void QWizardPage_SetButtonText(QtObjectPtr ptr, int which, char* text);
void QWizardPage_SetCommitPage(QtObjectPtr ptr, int commitPage);
void QWizardPage_SetFinalPage(QtObjectPtr ptr, int finalPage);
void QWizardPage_SetPixmap(QtObjectPtr ptr, int which, QtObjectPtr pixmap);
int QWizardPage_ValidatePage(QtObjectPtr ptr);
void QWizardPage_DestroyQWizardPage(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif