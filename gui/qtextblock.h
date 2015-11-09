#ifdef __cplusplus
extern "C" {
#endif

int QTextBlock_IsValid(void* ptr);
void* QTextBlock_NewQTextBlock(void* other);
int QTextBlock_BlockFormatIndex(void* ptr);
int QTextBlock_CharFormatIndex(void* ptr);
void QTextBlock_ClearLayout(void* ptr);
int QTextBlock_Contains(void* ptr, int position);
int QTextBlock_BlockNumber(void* ptr);
void* QTextBlock_Document(void* ptr);
int QTextBlock_FirstLineNumber(void* ptr);
int QTextBlock_IsVisible(void* ptr);
void* QTextBlock_Layout(void* ptr);
int QTextBlock_Length(void* ptr);
int QTextBlock_LineCount(void* ptr);
int QTextBlock_Position(void* ptr);
int QTextBlock_Revision(void* ptr);
void QTextBlock_SetLineCount(void* ptr, int count);
void QTextBlock_SetRevision(void* ptr, int rev);
void QTextBlock_SetUserData(void* ptr, void* data);
void QTextBlock_SetUserState(void* ptr, int state);
void QTextBlock_SetVisible(void* ptr, int visible);
char* QTextBlock_Text(void* ptr);
int QTextBlock_TextDirection(void* ptr);
void* QTextBlock_TextList(void* ptr);
void* QTextBlock_UserData(void* ptr);
int QTextBlock_UserState(void* ptr);

#ifdef __cplusplus
}
#endif