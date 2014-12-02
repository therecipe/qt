#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

//Public Functions
QtObjectPtr QLabel_New_QWidget_WindowType(QtObjectPtr parent, int f);
QtObjectPtr QLabel_New_String_QWidget_WindowType(char* text, QtObjectPtr parent, int f);
void QLabel_Destroy(QtObjectPtr ptr);
int QLabel_Alignment(QtObjectPtr ptr);
QtObjectPtr QLabel_Buddy(QtObjectPtr ptr);
int QLabel_HasScaledContents(QtObjectPtr ptr);
int QLabel_HasSelectedText(QtObjectPtr ptr);
int QLabel_Indent(QtObjectPtr ptr);
int QLabel_Margin(QtObjectPtr ptr);
int QLabel_OpenExternalLinks(QtObjectPtr ptr);
char* QLabel_SelectedText(QtObjectPtr ptr);
int QLabel_SelectionStart(QtObjectPtr ptr);
void QLabel_SetAlignment_AlignmentFlag(QtObjectPtr ptr, int alignment);
void QLabel_SetBuddy_QWidget(QtObjectPtr ptr, QtObjectPtr buddy);
void QLabel_SetIndent_Int(QtObjectPtr ptr, int indent);
void QLabel_SetMargin_Int(QtObjectPtr ptr, int margin);
void QLabel_SetOpenExternalLinks_Bool(QtObjectPtr ptr, int open);
void QLabel_SetScaledContents_Bool(QtObjectPtr ptr, int scaledContents);
void QLabel_SetSelection_Int_Int(QtObjectPtr ptr, int start, int length);
void QLabel_SetTextFormat_TextFormat(QtObjectPtr ptr, int textFormat);
void QLabel_SetTextInteractionFlags_TextInteractionFlag(QtObjectPtr ptr, int flags);
void QLabel_SetWordWrap_Bool(QtObjectPtr ptr, int on);
char* QLabel_Text(QtObjectPtr ptr);
int QLabel_TextFormat(QtObjectPtr ptr);
int QLabel_TextInteractionFlags(QtObjectPtr ptr);
int QLabel_WordWrap(QtObjectPtr ptr);
//Public Slots
void QLabel_ConnectSlotClear(QtObjectPtr ptr);
void QLabel_DisconnectSlotClear(QtObjectPtr ptr);
void QLabel_Clear(QtObjectPtr ptr);
void QLabel_ConnectSlotSetText(QtObjectPtr ptr);
void QLabel_DisconnectSlotSetText(QtObjectPtr ptr);
void QLabel_SetText_String(QtObjectPtr ptr, char* text);
//Signals
void QLabel_ConnectSignalLinkActivated(QtObjectPtr ptr);
void QLabel_DisconnectSignalLinkActivated(QtObjectPtr ptr);
void QLabel_ConnectSignalLinkHovered(QtObjectPtr ptr);
void QLabel_DisconnectSignalLinkHovered(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif
