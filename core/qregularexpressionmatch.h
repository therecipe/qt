#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QRegularExpressionMatch_NewQRegularExpressionMatch();
QtObjectPtr QRegularExpressionMatch_NewQRegularExpressionMatch2(QtObjectPtr match);
char* QRegularExpressionMatch_Captured2(QtObjectPtr ptr, char* name);
char* QRegularExpressionMatch_Captured(QtObjectPtr ptr, int nth);
int QRegularExpressionMatch_CapturedEnd2(QtObjectPtr ptr, char* name);
int QRegularExpressionMatch_CapturedEnd(QtObjectPtr ptr, int nth);
int QRegularExpressionMatch_CapturedLength2(QtObjectPtr ptr, char* name);
int QRegularExpressionMatch_CapturedLength(QtObjectPtr ptr, int nth);
int QRegularExpressionMatch_CapturedStart2(QtObjectPtr ptr, char* name);
int QRegularExpressionMatch_CapturedStart(QtObjectPtr ptr, int nth);
char* QRegularExpressionMatch_CapturedTexts(QtObjectPtr ptr);
int QRegularExpressionMatch_HasMatch(QtObjectPtr ptr);
int QRegularExpressionMatch_HasPartialMatch(QtObjectPtr ptr);
int QRegularExpressionMatch_IsValid(QtObjectPtr ptr);
int QRegularExpressionMatch_LastCapturedIndex(QtObjectPtr ptr);
int QRegularExpressionMatch_MatchOptions(QtObjectPtr ptr);
int QRegularExpressionMatch_MatchType(QtObjectPtr ptr);
void QRegularExpressionMatch_Swap(QtObjectPtr ptr, QtObjectPtr other);
void QRegularExpressionMatch_DestroyQRegularExpressionMatch(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif