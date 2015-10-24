#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QRegularExpression_NewQRegularExpression();
QtObjectPtr QRegularExpression_NewQRegularExpression3(QtObjectPtr re);
QtObjectPtr QRegularExpression_NewQRegularExpression2(char* pattern, int options);
int QRegularExpression_CaptureCount(QtObjectPtr ptr);
char* QRegularExpression_ErrorString(QtObjectPtr ptr);
char* QRegularExpression_QRegularExpression_Escape(char* str);
int QRegularExpression_IsValid(QtObjectPtr ptr);
char* QRegularExpression_NamedCaptureGroups(QtObjectPtr ptr);
void QRegularExpression_Optimize(QtObjectPtr ptr);
char* QRegularExpression_Pattern(QtObjectPtr ptr);
int QRegularExpression_PatternErrorOffset(QtObjectPtr ptr);
int QRegularExpression_PatternOptions(QtObjectPtr ptr);
void QRegularExpression_SetPattern(QtObjectPtr ptr, char* pattern);
void QRegularExpression_SetPatternOptions(QtObjectPtr ptr, int options);
void QRegularExpression_Swap(QtObjectPtr ptr, QtObjectPtr other);
void QRegularExpression_DestroyQRegularExpression(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif