#ifdef __cplusplus
extern "C" {
#endif

void* QRegularExpression_NewQRegularExpression();
void* QRegularExpression_NewQRegularExpression3(void* re);
void* QRegularExpression_NewQRegularExpression2(char* pattern, int options);
int QRegularExpression_CaptureCount(void* ptr);
char* QRegularExpression_ErrorString(void* ptr);
char* QRegularExpression_QRegularExpression_Escape(char* str);
void* QRegularExpression_GlobalMatch(void* ptr, char* subject, int offset, int matchType, int matchOptions);
void* QRegularExpression_GlobalMatch2(void* ptr, void* subjectRef, int offset, int matchType, int matchOptions);
int QRegularExpression_IsValid(void* ptr);
void* QRegularExpression_Match(void* ptr, char* subject, int offset, int matchType, int matchOptions);
void* QRegularExpression_Match2(void* ptr, void* subjectRef, int offset, int matchType, int matchOptions);
char* QRegularExpression_NamedCaptureGroups(void* ptr);
void QRegularExpression_Optimize(void* ptr);
char* QRegularExpression_Pattern(void* ptr);
int QRegularExpression_PatternErrorOffset(void* ptr);
int QRegularExpression_PatternOptions(void* ptr);
void QRegularExpression_SetPattern(void* ptr, char* pattern);
void QRegularExpression_SetPatternOptions(void* ptr, int options);
void QRegularExpression_Swap(void* ptr, void* other);
void QRegularExpression_DestroyQRegularExpression(void* ptr);

#ifdef __cplusplus
}
#endif