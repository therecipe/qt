#ifdef __cplusplus
extern "C" {
#endif

void QWizardPage_SetSubTitle(void* ptr, char* subTitle);
void QWizardPage_SetTitle(void* ptr, char* title);
char* QWizardPage_SubTitle(void* ptr);
char* QWizardPage_Title(void* ptr);
void* QWizardPage_NewQWizardPage(void* parent);
char* QWizardPage_ButtonText(void* ptr, int which);
void QWizardPage_CleanupPage(void* ptr);
void QWizardPage_ConnectCompleteChanged(void* ptr);
void QWizardPage_DisconnectCompleteChanged(void* ptr);
void QWizardPage_InitializePage(void* ptr);
int QWizardPage_IsCommitPage(void* ptr);
int QWizardPage_IsComplete(void* ptr);
int QWizardPage_IsFinalPage(void* ptr);
int QWizardPage_NextId(void* ptr);
void QWizardPage_SetButtonText(void* ptr, int which, char* text);
void QWizardPage_SetCommitPage(void* ptr, int commitPage);
void QWizardPage_SetFinalPage(void* ptr, int finalPage);
void QWizardPage_SetPixmap(void* ptr, int which, void* pixmap);
int QWizardPage_ValidatePage(void* ptr);
void QWizardPage_DestroyQWizardPage(void* ptr);

#ifdef __cplusplus
}
#endif