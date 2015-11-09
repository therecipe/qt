#ifdef __cplusplus
extern "C" {
#endif

int QFont_Times_Type();
int QFont_Courier_Type();
int QFont_OldEnglish_Type();
int QFont_System_Type();
int QFont_AnyStyle_Type();
int QFont_Cursive_Type();
int QFont_Monospace_Type();
int QFont_Fantasy_Type();
char* QFont_DefaultFamily(void* ptr);
char* QFont_LastResortFamily(void* ptr);
char* QFont_LastResortFont(void* ptr);
void* QFont_NewQFont();
void* QFont_NewQFont4(void* font);
void* QFont_NewQFont3(void* font, void* pd);
void* QFont_NewQFont2(char* family, int pointSize, int weight, int italic);
int QFont_Bold(void* ptr);
int QFont_Capitalization(void* ptr);
int QFont_ExactMatch(void* ptr);
char* QFont_Family(void* ptr);
int QFont_FixedPitch(void* ptr);
int QFont_FromString(void* ptr, char* descrip);
int QFont_HintingPreference(void* ptr);
void QFont_QFont_InsertSubstitution(char* familyName, char* substituteName);
void QFont_QFont_InsertSubstitutions(char* familyName, char* substituteNames);
int QFont_IsCopyOf(void* ptr, void* f);
int QFont_Italic(void* ptr);
int QFont_Kerning(void* ptr);
char* QFont_Key(void* ptr);
double QFont_LetterSpacing(void* ptr);
int QFont_LetterSpacingType(void* ptr);
int QFont_Overline(void* ptr);
int QFont_PixelSize(void* ptr);
int QFont_PointSize(void* ptr);
double QFont_PointSizeF(void* ptr);
void QFont_QFont_RemoveSubstitutions(char* familyName);
void QFont_SetBold(void* ptr, int enable);
void QFont_SetCapitalization(void* ptr, int caps);
void QFont_SetFamily(void* ptr, char* family);
void QFont_SetFixedPitch(void* ptr, int enable);
void QFont_SetHintingPreference(void* ptr, int hintingPreference);
void QFont_SetItalic(void* ptr, int enable);
void QFont_SetKerning(void* ptr, int enable);
void QFont_SetLetterSpacing(void* ptr, int ty, double spacing);
void QFont_SetOverline(void* ptr, int enable);
void QFont_SetPixelSize(void* ptr, int pixelSize);
void QFont_SetPointSize(void* ptr, int pointSize);
void QFont_SetPointSizeF(void* ptr, double pointSize);
void QFont_SetStretch(void* ptr, int factor);
void QFont_SetStrikeOut(void* ptr, int enable);
void QFont_SetStyle(void* ptr, int style);
void QFont_SetStyleHint(void* ptr, int hint, int strategy);
void QFont_SetStyleName(void* ptr, char* styleName);
void QFont_SetStyleStrategy(void* ptr, int s);
void QFont_SetUnderline(void* ptr, int enable);
void QFont_SetWeight(void* ptr, int weight);
void QFont_SetWordSpacing(void* ptr, double spacing);
int QFont_Stretch(void* ptr);
int QFont_StrikeOut(void* ptr);
int QFont_Style(void* ptr);
int QFont_StyleHint(void* ptr);
char* QFont_StyleName(void* ptr);
int QFont_StyleStrategy(void* ptr);
char* QFont_QFont_Substitute(char* familyName);
char* QFont_QFont_Substitutes(char* familyName);
char* QFont_QFont_Substitutions();
void QFont_Swap(void* ptr, void* other);
char* QFont_ToString(void* ptr);
int QFont_Underline(void* ptr);
int QFont_Weight(void* ptr);
double QFont_WordSpacing(void* ptr);
void QFont_DestroyQFont(void* ptr);

#ifdef __cplusplus
}
#endif