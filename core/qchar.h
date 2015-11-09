#ifdef __cplusplus
extern "C" {
#endif

void* QChar_NewQChar();
void* QChar_NewQChar8(void* ch);
void* QChar_NewQChar7(int ch);
void* QChar_NewQChar9(char* ch);
void* QChar_NewQChar6(int code);
int QChar_Category(void* ptr);
int QChar_QChar_CurrentUnicodeVersion();
char* QChar_Decomposition(void* ptr);
int QChar_DecompositionTag(void* ptr);
int QChar_DigitValue(void* ptr);
int QChar_Direction(void* ptr);
int QChar_HasMirrored(void* ptr);
int QChar_IsDigit(void* ptr);
int QChar_IsHighSurrogate(void* ptr);
int QChar_IsLetter(void* ptr);
int QChar_IsLetterOrNumber(void* ptr);
int QChar_IsLower(void* ptr);
int QChar_IsLowSurrogate(void* ptr);
int QChar_IsMark(void* ptr);
int QChar_IsNonCharacter(void* ptr);
int QChar_IsNull(void* ptr);
int QChar_IsNumber(void* ptr);
int QChar_IsPrint(void* ptr);
int QChar_IsPunct(void* ptr);
int QChar_IsSpace(void* ptr);
int QChar_IsSurrogate(void* ptr);
int QChar_IsSymbol(void* ptr);
int QChar_IsTitleCase(void* ptr);
int QChar_IsUpper(void* ptr);
int QChar_JoiningType(void* ptr);
int QChar_Script(void* ptr);
int QChar_UnicodeVersion(void* ptr);

#ifdef __cplusplus
}
#endif