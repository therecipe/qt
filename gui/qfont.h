#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QFont_Times_Type();
int QFont_Courier_Type();
int QFont_OldEnglish_Type();
int QFont_System_Type();
int QFont_AnyStyle_Type();
int QFont_Cursive_Type();
int QFont_Monospace_Type();
int QFont_Fantasy_Type();
char* QFont_DefaultFamily(QtObjectPtr ptr);
char* QFont_LastResortFamily(QtObjectPtr ptr);
char* QFont_LastResortFont(QtObjectPtr ptr);
QtObjectPtr QFont_NewQFont();
QtObjectPtr QFont_NewQFont4(QtObjectPtr font);
QtObjectPtr QFont_NewQFont3(QtObjectPtr font, QtObjectPtr pd);
QtObjectPtr QFont_NewQFont2(char* family, int pointSize, int weight, int italic);
int QFont_Bold(QtObjectPtr ptr);
int QFont_Capitalization(QtObjectPtr ptr);
int QFont_ExactMatch(QtObjectPtr ptr);
char* QFont_Family(QtObjectPtr ptr);
int QFont_FixedPitch(QtObjectPtr ptr);
int QFont_FromString(QtObjectPtr ptr, char* descrip);
int QFont_HintingPreference(QtObjectPtr ptr);
void QFont_QFont_InsertSubstitution(char* familyName, char* substituteName);
void QFont_QFont_InsertSubstitutions(char* familyName, char* substituteNames);
int QFont_IsCopyOf(QtObjectPtr ptr, QtObjectPtr f);
int QFont_Italic(QtObjectPtr ptr);
int QFont_Kerning(QtObjectPtr ptr);
char* QFont_Key(QtObjectPtr ptr);
int QFont_LetterSpacingType(QtObjectPtr ptr);
int QFont_Overline(QtObjectPtr ptr);
int QFont_PixelSize(QtObjectPtr ptr);
int QFont_PointSize(QtObjectPtr ptr);
void QFont_QFont_RemoveSubstitutions(char* familyName);
void QFont_SetBold(QtObjectPtr ptr, int enable);
void QFont_SetCapitalization(QtObjectPtr ptr, int caps);
void QFont_SetFamily(QtObjectPtr ptr, char* family);
void QFont_SetFixedPitch(QtObjectPtr ptr, int enable);
void QFont_SetHintingPreference(QtObjectPtr ptr, int hintingPreference);
void QFont_SetItalic(QtObjectPtr ptr, int enable);
void QFont_SetKerning(QtObjectPtr ptr, int enable);
void QFont_SetOverline(QtObjectPtr ptr, int enable);
void QFont_SetPixelSize(QtObjectPtr ptr, int pixelSize);
void QFont_SetPointSize(QtObjectPtr ptr, int pointSize);
void QFont_SetStretch(QtObjectPtr ptr, int factor);
void QFont_SetStrikeOut(QtObjectPtr ptr, int enable);
void QFont_SetStyle(QtObjectPtr ptr, int style);
void QFont_SetStyleHint(QtObjectPtr ptr, int hint, int strategy);
void QFont_SetStyleName(QtObjectPtr ptr, char* styleName);
void QFont_SetStyleStrategy(QtObjectPtr ptr, int s);
void QFont_SetUnderline(QtObjectPtr ptr, int enable);
void QFont_SetWeight(QtObjectPtr ptr, int weight);
int QFont_Stretch(QtObjectPtr ptr);
int QFont_StrikeOut(QtObjectPtr ptr);
int QFont_Style(QtObjectPtr ptr);
int QFont_StyleHint(QtObjectPtr ptr);
char* QFont_StyleName(QtObjectPtr ptr);
int QFont_StyleStrategy(QtObjectPtr ptr);
char* QFont_QFont_Substitute(char* familyName);
char* QFont_QFont_Substitutes(char* familyName);
char* QFont_QFont_Substitutions();
void QFont_Swap(QtObjectPtr ptr, QtObjectPtr other);
char* QFont_ToString(QtObjectPtr ptr);
int QFont_Underline(QtObjectPtr ptr);
int QFont_Weight(QtObjectPtr ptr);
void QFont_DestroyQFont(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif