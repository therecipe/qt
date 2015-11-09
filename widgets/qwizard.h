#ifdef __cplusplus
extern "C" {
#endif

int QWizard_CurrentId(void* ptr);
int QWizard_HasVisitedPage(void* ptr, int id);
int QWizard_Options(void* ptr);
void* QWizard_Page(void* ptr, int id);
void QWizard_SetOptions(void* ptr, int options);
void QWizard_SetPage(void* ptr, int id, void* page);
void QWizard_SetStartId(void* ptr, int id);
void QWizard_SetSubTitleFormat(void* ptr, int format);
void QWizard_SetTitleFormat(void* ptr, int format);
void QWizard_SetWizardStyle(void* ptr, int style);
int QWizard_StartId(void* ptr);
int QWizard_SubTitleFormat(void* ptr);
int QWizard_TitleFormat(void* ptr);
int QWizard_WizardStyle(void* ptr);
void* QWizard_NewQWizard(void* parent, int flags);
int QWizard_AddPage(void* ptr, void* page);
void QWizard_Back(void* ptr);
void* QWizard_Button(void* ptr, int which);
char* QWizard_ButtonText(void* ptr, int which);
void QWizard_ConnectCurrentIdChanged(void* ptr);
void QWizard_DisconnectCurrentIdChanged(void* ptr);
void* QWizard_CurrentPage(void* ptr);
void QWizard_ConnectCustomButtonClicked(void* ptr);
void QWizard_DisconnectCustomButtonClicked(void* ptr);
void* QWizard_Field(void* ptr, char* name);
void QWizard_ConnectHelpRequested(void* ptr);
void QWizard_DisconnectHelpRequested(void* ptr);
void QWizard_Next(void* ptr);
int QWizard_NextId(void* ptr);
void QWizard_ConnectPageAdded(void* ptr);
void QWizard_DisconnectPageAdded(void* ptr);
void QWizard_ConnectPageRemoved(void* ptr);
void QWizard_DisconnectPageRemoved(void* ptr);
void QWizard_RemovePage(void* ptr, int id);
void QWizard_Restart(void* ptr);
void QWizard_SetButton(void* ptr, int which, void* button);
void QWizard_SetButtonText(void* ptr, int which, char* text);
void QWizard_SetDefaultProperty(void* ptr, char* className, char* property, char* changedSignal);
void QWizard_SetField(void* ptr, char* name, void* value);
void QWizard_SetOption(void* ptr, int option, int on);
void QWizard_SetPixmap(void* ptr, int which, void* pixmap);
void QWizard_SetSideWidget(void* ptr, void* widget);
void QWizard_SetVisible(void* ptr, int visible);
void* QWizard_SideWidget(void* ptr);
int QWizard_TestOption(void* ptr, int option);
int QWizard_ValidateCurrentPage(void* ptr);
void QWizard_DestroyQWizard(void* ptr);

#ifdef __cplusplus
}
#endif