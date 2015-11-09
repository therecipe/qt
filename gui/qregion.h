#ifdef __cplusplus
extern "C" {
#endif

void* QRegion_NewQRegion();
void* QRegion_NewQRegion5(void* bm);
void* QRegion_NewQRegion3(void* a, int fillRule);
void* QRegion_NewQRegion6(void* r, int t);
void* QRegion_NewQRegion4(void* r);
int QRegion_Contains(void* ptr, void* p);
int QRegion_Contains2(void* ptr, void* r);
void* QRegion_Intersected2(void* ptr, void* rect);
void* QRegion_Intersected(void* ptr, void* r);
int QRegion_Intersects2(void* ptr, void* rect);
int QRegion_IsEmpty(void* ptr);
int QRegion_IsNull(void* ptr);
int QRegion_RectCount(void* ptr);
void QRegion_SetRects(void* ptr, void* rects, int number);
void* QRegion_Subtracted(void* ptr, void* r);
void QRegion_Translate(void* ptr, int dx, int dy);
void* QRegion_United2(void* ptr, void* rect);
void* QRegion_United(void* ptr, void* r);
void* QRegion_Xored(void* ptr, void* r);
void* QRegion_NewQRegion2(int x, int y, int w, int h, int t);
int QRegion_Intersects(void* ptr, void* region);
void QRegion_Swap(void* ptr, void* other);
void QRegion_Translate2(void* ptr, void* point);
void* QRegion_Translated2(void* ptr, void* p);
void* QRegion_Translated(void* ptr, int dx, int dy);

#ifdef __cplusplus
}
#endif