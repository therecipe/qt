#ifdef __cplusplus
extern "C" {
#endif

void* QFontInfo_NewQFontInfo(void* font);
void* QFontInfo_NewQFontInfo2(void* fi);
int QFontInfo_Bold(void* ptr);
int QFontInfo_ExactMatch(void* ptr);
char* QFontInfo_Family(void* ptr);
int QFontInfo_FixedPitch(void* ptr);
int QFontInfo_Italic(void* ptr);
int QFontInfo_PixelSize(void* ptr);
int QFontInfo_PointSize(void* ptr);
double QFontInfo_PointSizeF(void* ptr);
int QFontInfo_Style(void* ptr);
char* QFontInfo_StyleName(void* ptr);
void QFontInfo_Swap(void* ptr, void* other);
int QFontInfo_StyleHint(void* ptr);
int QFontInfo_Weight(void* ptr);
void QFontInfo_DestroyQFontInfo(void* ptr);

#ifdef __cplusplus
}
#endif