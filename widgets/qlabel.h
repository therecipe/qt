#ifdef __cplusplus
extern "C" {
#endif

int QLabel_Alignment(void* ptr);
int QLabel_HasScaledContents(void* ptr);
int QLabel_HasSelectedText(void* ptr);
int QLabel_Indent(void* ptr);
int QLabel_Margin(void* ptr);
int QLabel_OpenExternalLinks(void* ptr);
void* QLabel_Pixmap(void* ptr);
char* QLabel_SelectedText(void* ptr);
void QLabel_SetAlignment(void* ptr, int v);
void QLabel_SetIndent(void* ptr, int v);
void QLabel_SetMargin(void* ptr, int v);
void QLabel_SetOpenExternalLinks(void* ptr, int open);
void QLabel_SetPixmap(void* ptr, void* v);
void QLabel_SetScaledContents(void* ptr, int v);
void QLabel_SetText(void* ptr, char* v);
void QLabel_SetTextFormat(void* ptr, int v);
void QLabel_SetTextInteractionFlags(void* ptr, int flags);
void QLabel_SetWordWrap(void* ptr, int on);
char* QLabel_Text(void* ptr);
int QLabel_TextFormat(void* ptr);
int QLabel_TextInteractionFlags(void* ptr);
int QLabel_WordWrap(void* ptr);
void* QLabel_NewQLabel(void* parent, int f);
void* QLabel_NewQLabel2(char* text, void* parent, int f);
void* QLabel_Buddy(void* ptr);
void QLabel_Clear(void* ptr);
int QLabel_HeightForWidth(void* ptr, int w);
void QLabel_ConnectLinkActivated(void* ptr);
void QLabel_DisconnectLinkActivated(void* ptr);
void QLabel_ConnectLinkHovered(void* ptr);
void QLabel_DisconnectLinkHovered(void* ptr);
void* QLabel_Movie(void* ptr);
void* QLabel_Picture(void* ptr);
int QLabel_SelectionStart(void* ptr);
void QLabel_SetBuddy(void* ptr, void* buddy);
void QLabel_SetMovie(void* ptr, void* movie);
void QLabel_SetNum(void* ptr, int num);
void QLabel_SetPicture(void* ptr, void* picture);
void QLabel_SetSelection(void* ptr, int start, int length);
void QLabel_DestroyQLabel(void* ptr);

#ifdef __cplusplus
}
#endif