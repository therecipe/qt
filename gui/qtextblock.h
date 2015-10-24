#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QTextBlock_IsValid(QtObjectPtr ptr);
QtObjectPtr QTextBlock_NewQTextBlock(QtObjectPtr other);
int QTextBlock_BlockFormatIndex(QtObjectPtr ptr);
int QTextBlock_CharFormatIndex(QtObjectPtr ptr);
void QTextBlock_ClearLayout(QtObjectPtr ptr);
int QTextBlock_Contains(QtObjectPtr ptr, int position);
int QTextBlock_BlockNumber(QtObjectPtr ptr);
QtObjectPtr QTextBlock_Document(QtObjectPtr ptr);
int QTextBlock_FirstLineNumber(QtObjectPtr ptr);
int QTextBlock_IsVisible(QtObjectPtr ptr);
QtObjectPtr QTextBlock_Layout(QtObjectPtr ptr);
int QTextBlock_Length(QtObjectPtr ptr);
int QTextBlock_LineCount(QtObjectPtr ptr);
int QTextBlock_Position(QtObjectPtr ptr);
int QTextBlock_Revision(QtObjectPtr ptr);
void QTextBlock_SetLineCount(QtObjectPtr ptr, int count);
void QTextBlock_SetRevision(QtObjectPtr ptr, int rev);
void QTextBlock_SetUserData(QtObjectPtr ptr, QtObjectPtr data);
void QTextBlock_SetUserState(QtObjectPtr ptr, int state);
void QTextBlock_SetVisible(QtObjectPtr ptr, int visible);
char* QTextBlock_Text(QtObjectPtr ptr);
int QTextBlock_TextDirection(QtObjectPtr ptr);
QtObjectPtr QTextBlock_TextList(QtObjectPtr ptr);
QtObjectPtr QTextBlock_UserData(QtObjectPtr ptr);
int QTextBlock_UserState(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif