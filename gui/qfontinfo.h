#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QFontInfo_NewQFontInfo(QtObjectPtr font);
QtObjectPtr QFontInfo_NewQFontInfo2(QtObjectPtr fi);
int QFontInfo_Bold(QtObjectPtr ptr);
int QFontInfo_ExactMatch(QtObjectPtr ptr);
char* QFontInfo_Family(QtObjectPtr ptr);
int QFontInfo_FixedPitch(QtObjectPtr ptr);
int QFontInfo_Italic(QtObjectPtr ptr);
int QFontInfo_PixelSize(QtObjectPtr ptr);
int QFontInfo_PointSize(QtObjectPtr ptr);
int QFontInfo_Style(QtObjectPtr ptr);
char* QFontInfo_StyleName(QtObjectPtr ptr);
void QFontInfo_Swap(QtObjectPtr ptr, QtObjectPtr other);
int QFontInfo_StyleHint(QtObjectPtr ptr);
int QFontInfo_Weight(QtObjectPtr ptr);
void QFontInfo_DestroyQFontInfo(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif