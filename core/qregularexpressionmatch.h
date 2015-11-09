#ifdef __cplusplus
extern "C" {
#endif

void* QRegularExpressionMatch_NewQRegularExpressionMatch();
void* QRegularExpressionMatch_NewQRegularExpressionMatch2(void* match);
char* QRegularExpressionMatch_Captured2(void* ptr, char* name);
char* QRegularExpressionMatch_Captured(void* ptr, int nth);
int QRegularExpressionMatch_CapturedEnd2(void* ptr, char* name);
int QRegularExpressionMatch_CapturedEnd(void* ptr, int nth);
int QRegularExpressionMatch_CapturedLength2(void* ptr, char* name);
int QRegularExpressionMatch_CapturedLength(void* ptr, int nth);
void* QRegularExpressionMatch_CapturedRef2(void* ptr, char* name);
void* QRegularExpressionMatch_CapturedRef(void* ptr, int nth);
int QRegularExpressionMatch_CapturedStart2(void* ptr, char* name);
int QRegularExpressionMatch_CapturedStart(void* ptr, int nth);
char* QRegularExpressionMatch_CapturedTexts(void* ptr);
int QRegularExpressionMatch_HasMatch(void* ptr);
int QRegularExpressionMatch_HasPartialMatch(void* ptr);
int QRegularExpressionMatch_IsValid(void* ptr);
int QRegularExpressionMatch_LastCapturedIndex(void* ptr);
int QRegularExpressionMatch_MatchOptions(void* ptr);
int QRegularExpressionMatch_MatchType(void* ptr);
void* QRegularExpressionMatch_RegularExpression(void* ptr);
void QRegularExpressionMatch_Swap(void* ptr, void* other);
void QRegularExpressionMatch_DestroyQRegularExpressionMatch(void* ptr);

#ifdef __cplusplus
}
#endif