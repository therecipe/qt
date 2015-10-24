#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QRegExp_NewQRegExp();
QtObjectPtr QRegExp_NewQRegExp3(QtObjectPtr rx);
QtObjectPtr QRegExp_NewQRegExp2(char* pattern, int cs, int syntax);
char* QRegExp_Cap(QtObjectPtr ptr, int nth);
char* QRegExp_ErrorString(QtObjectPtr ptr);
int QRegExp_CaptureCount(QtObjectPtr ptr);
char* QRegExp_CapturedTexts(QtObjectPtr ptr);
int QRegExp_CaseSensitivity(QtObjectPtr ptr);
char* QRegExp_QRegExp_Escape(char* str);
int QRegExp_ExactMatch(QtObjectPtr ptr, char* str);
int QRegExp_IndexIn(QtObjectPtr ptr, char* str, int offset, int caretMode);
int QRegExp_IsEmpty(QtObjectPtr ptr);
int QRegExp_IsMinimal(QtObjectPtr ptr);
int QRegExp_IsValid(QtObjectPtr ptr);
int QRegExp_LastIndexIn(QtObjectPtr ptr, char* str, int offset, int caretMode);
int QRegExp_MatchedLength(QtObjectPtr ptr);
char* QRegExp_Pattern(QtObjectPtr ptr);
int QRegExp_PatternSyntax(QtObjectPtr ptr);
int QRegExp_Pos(QtObjectPtr ptr, int nth);
void QRegExp_SetCaseSensitivity(QtObjectPtr ptr, int cs);
void QRegExp_SetMinimal(QtObjectPtr ptr, int minimal);
void QRegExp_SetPattern(QtObjectPtr ptr, char* pattern);
void QRegExp_SetPatternSyntax(QtObjectPtr ptr, int syntax);
void QRegExp_Swap(QtObjectPtr ptr, QtObjectPtr other);
void QRegExp_DestroyQRegExp(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif