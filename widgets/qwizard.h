#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QWizard_CurrentId(QtObjectPtr ptr);
int QWizard_HasVisitedPage(QtObjectPtr ptr, int id);
int QWizard_Options(QtObjectPtr ptr);
QtObjectPtr QWizard_Page(QtObjectPtr ptr, int id);
void QWizard_SetOptions(QtObjectPtr ptr, int options);
void QWizard_SetPage(QtObjectPtr ptr, int id, QtObjectPtr page);
void QWizard_SetStartId(QtObjectPtr ptr, int id);
void QWizard_SetSubTitleFormat(QtObjectPtr ptr, int format);
void QWizard_SetTitleFormat(QtObjectPtr ptr, int format);
void QWizard_SetWizardStyle(QtObjectPtr ptr, int style);
int QWizard_StartId(QtObjectPtr ptr);
int QWizard_SubTitleFormat(QtObjectPtr ptr);
int QWizard_TitleFormat(QtObjectPtr ptr);
int QWizard_WizardStyle(QtObjectPtr ptr);
QtObjectPtr QWizard_NewQWizard(QtObjectPtr parent, int flags);
int QWizard_AddPage(QtObjectPtr ptr, QtObjectPtr page);
void QWizard_Back(QtObjectPtr ptr);
QtObjectPtr QWizard_Button(QtObjectPtr ptr, int which);
char* QWizard_ButtonText(QtObjectPtr ptr, int which);
void QWizard_ConnectCurrentIdChanged(QtObjectPtr ptr);
void QWizard_DisconnectCurrentIdChanged(QtObjectPtr ptr);
QtObjectPtr QWizard_CurrentPage(QtObjectPtr ptr);
void QWizard_ConnectCustomButtonClicked(QtObjectPtr ptr);
void QWizard_DisconnectCustomButtonClicked(QtObjectPtr ptr);
char* QWizard_Field(QtObjectPtr ptr, char* name);
void QWizard_ConnectHelpRequested(QtObjectPtr ptr);
void QWizard_DisconnectHelpRequested(QtObjectPtr ptr);
void QWizard_Next(QtObjectPtr ptr);
int QWizard_NextId(QtObjectPtr ptr);
void QWizard_ConnectPageAdded(QtObjectPtr ptr);
void QWizard_DisconnectPageAdded(QtObjectPtr ptr);
void QWizard_ConnectPageRemoved(QtObjectPtr ptr);
void QWizard_DisconnectPageRemoved(QtObjectPtr ptr);
void QWizard_RemovePage(QtObjectPtr ptr, int id);
void QWizard_Restart(QtObjectPtr ptr);
void QWizard_SetButton(QtObjectPtr ptr, int which, QtObjectPtr button);
void QWizard_SetButtonText(QtObjectPtr ptr, int which, char* text);
void QWizard_SetDefaultProperty(QtObjectPtr ptr, char* className, char* property, char* changedSignal);
void QWizard_SetField(QtObjectPtr ptr, char* name, char* value);
void QWizard_SetOption(QtObjectPtr ptr, int option, int on);
void QWizard_SetPixmap(QtObjectPtr ptr, int which, QtObjectPtr pixmap);
void QWizard_SetSideWidget(QtObjectPtr ptr, QtObjectPtr widget);
void QWizard_SetVisible(QtObjectPtr ptr, int visible);
QtObjectPtr QWizard_SideWidget(QtObjectPtr ptr);
int QWizard_TestOption(QtObjectPtr ptr, int option);
int QWizard_ValidateCurrentPage(QtObjectPtr ptr);
void QWizard_DestroyQWizard(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif