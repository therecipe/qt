#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QChar_NewQChar();
QtObjectPtr QChar_NewQChar8(QtObjectPtr ch);
QtObjectPtr QChar_NewQChar7(int ch);
QtObjectPtr QChar_NewQChar9(char* ch);
QtObjectPtr QChar_NewQChar6(int code);
int QChar_Category(QtObjectPtr ptr);
int QChar_QChar_CurrentUnicodeVersion();
char* QChar_Decomposition(QtObjectPtr ptr);
int QChar_DecompositionTag(QtObjectPtr ptr);
int QChar_DigitValue(QtObjectPtr ptr);
int QChar_Direction(QtObjectPtr ptr);
int QChar_HasMirrored(QtObjectPtr ptr);
int QChar_IsDigit(QtObjectPtr ptr);
int QChar_IsHighSurrogate(QtObjectPtr ptr);
int QChar_IsLetter(QtObjectPtr ptr);
int QChar_IsLetterOrNumber(QtObjectPtr ptr);
int QChar_IsLower(QtObjectPtr ptr);
int QChar_IsLowSurrogate(QtObjectPtr ptr);
int QChar_IsMark(QtObjectPtr ptr);
int QChar_IsNonCharacter(QtObjectPtr ptr);
int QChar_IsNull(QtObjectPtr ptr);
int QChar_IsNumber(QtObjectPtr ptr);
int QChar_IsPrint(QtObjectPtr ptr);
int QChar_IsPunct(QtObjectPtr ptr);
int QChar_IsSpace(QtObjectPtr ptr);
int QChar_IsSurrogate(QtObjectPtr ptr);
int QChar_IsSymbol(QtObjectPtr ptr);
int QChar_IsTitleCase(QtObjectPtr ptr);
int QChar_IsUpper(QtObjectPtr ptr);
int QChar_JoiningType(QtObjectPtr ptr);
int QChar_Script(QtObjectPtr ptr);
int QChar_UnicodeVersion(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif