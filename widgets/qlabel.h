#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QLabel_Alignment(QtObjectPtr ptr);
int QLabel_HasScaledContents(QtObjectPtr ptr);
int QLabel_HasSelectedText(QtObjectPtr ptr);
int QLabel_Indent(QtObjectPtr ptr);
int QLabel_Margin(QtObjectPtr ptr);
int QLabel_OpenExternalLinks(QtObjectPtr ptr);
QtObjectPtr QLabel_Pixmap(QtObjectPtr ptr);
char* QLabel_SelectedText(QtObjectPtr ptr);
void QLabel_SetAlignment(QtObjectPtr ptr, int v);
void QLabel_SetIndent(QtObjectPtr ptr, int v);
void QLabel_SetMargin(QtObjectPtr ptr, int v);
void QLabel_SetOpenExternalLinks(QtObjectPtr ptr, int open);
void QLabel_SetPixmap(QtObjectPtr ptr, QtObjectPtr v);
void QLabel_SetScaledContents(QtObjectPtr ptr, int v);
void QLabel_SetText(QtObjectPtr ptr, char* v);
void QLabel_SetTextFormat(QtObjectPtr ptr, int v);
void QLabel_SetTextInteractionFlags(QtObjectPtr ptr, int flags);
void QLabel_SetWordWrap(QtObjectPtr ptr, int on);
char* QLabel_Text(QtObjectPtr ptr);
int QLabel_TextFormat(QtObjectPtr ptr);
int QLabel_TextInteractionFlags(QtObjectPtr ptr);
int QLabel_WordWrap(QtObjectPtr ptr);
QtObjectPtr QLabel_NewQLabel(QtObjectPtr parent, int f);
QtObjectPtr QLabel_NewQLabel2(char* text, QtObjectPtr parent, int f);
QtObjectPtr QLabel_Buddy(QtObjectPtr ptr);
void QLabel_Clear(QtObjectPtr ptr);
int QLabel_HeightForWidth(QtObjectPtr ptr, int w);
void QLabel_ConnectLinkActivated(QtObjectPtr ptr);
void QLabel_DisconnectLinkActivated(QtObjectPtr ptr);
void QLabel_ConnectLinkHovered(QtObjectPtr ptr);
void QLabel_DisconnectLinkHovered(QtObjectPtr ptr);
QtObjectPtr QLabel_Movie(QtObjectPtr ptr);
QtObjectPtr QLabel_Picture(QtObjectPtr ptr);
int QLabel_SelectionStart(QtObjectPtr ptr);
void QLabel_SetBuddy(QtObjectPtr ptr, QtObjectPtr buddy);
void QLabel_SetMovie(QtObjectPtr ptr, QtObjectPtr movie);
void QLabel_SetNum(QtObjectPtr ptr, int num);
void QLabel_SetPicture(QtObjectPtr ptr, QtObjectPtr picture);
void QLabel_SetSelection(QtObjectPtr ptr, int start, int length);
void QLabel_DestroyQLabel(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif