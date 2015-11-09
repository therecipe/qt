#ifdef __cplusplus
extern "C" {
#endif

void* QRegExp_NewQRegExp();
void* QRegExp_NewQRegExp3(void* rx);
void* QRegExp_NewQRegExp2(char* pattern, int cs, int syntax);
char* QRegExp_Cap(void* ptr, int nth);
char* QRegExp_ErrorString(void* ptr);
int QRegExp_CaptureCount(void* ptr);
char* QRegExp_CapturedTexts(void* ptr);
int QRegExp_CaseSensitivity(void* ptr);
char* QRegExp_QRegExp_Escape(char* str);
int QRegExp_ExactMatch(void* ptr, char* str);
int QRegExp_IndexIn(void* ptr, char* str, int offset, int caretMode);
int QRegExp_IsEmpty(void* ptr);
int QRegExp_IsMinimal(void* ptr);
int QRegExp_IsValid(void* ptr);
int QRegExp_LastIndexIn(void* ptr, char* str, int offset, int caretMode);
int QRegExp_MatchedLength(void* ptr);
char* QRegExp_Pattern(void* ptr);
int QRegExp_PatternSyntax(void* ptr);
int QRegExp_Pos(void* ptr, int nth);
void QRegExp_SetCaseSensitivity(void* ptr, int cs);
void QRegExp_SetMinimal(void* ptr, int minimal);
void QRegExp_SetPattern(void* ptr, char* pattern);
void QRegExp_SetPatternSyntax(void* ptr, int syntax);
void QRegExp_Swap(void* ptr, void* other);
void QRegExp_DestroyQRegExp(void* ptr);

#ifdef __cplusplus
}
#endif