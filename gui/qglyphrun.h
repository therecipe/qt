#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QGlyphRun_NewQGlyphRun();
QtObjectPtr QGlyphRun_NewQGlyphRun2(QtObjectPtr other);
void QGlyphRun_Clear(QtObjectPtr ptr);
int QGlyphRun_Flags(QtObjectPtr ptr);
int QGlyphRun_IsEmpty(QtObjectPtr ptr);
int QGlyphRun_IsRightToLeft(QtObjectPtr ptr);
int QGlyphRun_Overline(QtObjectPtr ptr);
void QGlyphRun_SetBoundingRect(QtObjectPtr ptr, QtObjectPtr boundingRect);
void QGlyphRun_SetFlag(QtObjectPtr ptr, int flag, int enabled);
void QGlyphRun_SetFlags(QtObjectPtr ptr, int flags);
void QGlyphRun_SetOverline(QtObjectPtr ptr, int overline);
void QGlyphRun_SetRawFont(QtObjectPtr ptr, QtObjectPtr rawFont);
void QGlyphRun_SetRightToLeft(QtObjectPtr ptr, int rightToLeft);
void QGlyphRun_SetStrikeOut(QtObjectPtr ptr, int strikeOut);
void QGlyphRun_SetUnderline(QtObjectPtr ptr, int underline);
int QGlyphRun_StrikeOut(QtObjectPtr ptr);
void QGlyphRun_Swap(QtObjectPtr ptr, QtObjectPtr other);
int QGlyphRun_Underline(QtObjectPtr ptr);
void QGlyphRun_DestroyQGlyphRun(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif