#ifdef __cplusplus
extern "C" {
#endif

void* QGlyphRun_NewQGlyphRun();
void* QGlyphRun_NewQGlyphRun2(void* other);
void QGlyphRun_Clear(void* ptr);
int QGlyphRun_Flags(void* ptr);
int QGlyphRun_IsEmpty(void* ptr);
int QGlyphRun_IsRightToLeft(void* ptr);
int QGlyphRun_Overline(void* ptr);
void QGlyphRun_SetBoundingRect(void* ptr, void* boundingRect);
void QGlyphRun_SetFlag(void* ptr, int flag, int enabled);
void QGlyphRun_SetFlags(void* ptr, int flags);
void QGlyphRun_SetOverline(void* ptr, int overline);
void QGlyphRun_SetRawFont(void* ptr, void* rawFont);
void QGlyphRun_SetRightToLeft(void* ptr, int rightToLeft);
void QGlyphRun_SetStrikeOut(void* ptr, int strikeOut);
void QGlyphRun_SetUnderline(void* ptr, int underline);
int QGlyphRun_StrikeOut(void* ptr);
void QGlyphRun_Swap(void* ptr, void* other);
int QGlyphRun_Underline(void* ptr);
void QGlyphRun_DestroyQGlyphRun(void* ptr);

#ifdef __cplusplus
}
#endif