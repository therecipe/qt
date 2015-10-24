#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QRegion_NewQRegion();
QtObjectPtr QRegion_NewQRegion5(QtObjectPtr bm);
QtObjectPtr QRegion_NewQRegion3(QtObjectPtr a, int fillRule);
QtObjectPtr QRegion_NewQRegion6(QtObjectPtr r, int t);
QtObjectPtr QRegion_NewQRegion4(QtObjectPtr r);
int QRegion_Contains(QtObjectPtr ptr, QtObjectPtr p);
int QRegion_Contains2(QtObjectPtr ptr, QtObjectPtr r);
int QRegion_Intersects2(QtObjectPtr ptr, QtObjectPtr rect);
int QRegion_IsEmpty(QtObjectPtr ptr);
int QRegion_IsNull(QtObjectPtr ptr);
int QRegion_RectCount(QtObjectPtr ptr);
void QRegion_SetRects(QtObjectPtr ptr, QtObjectPtr rects, int number);
void QRegion_Translate(QtObjectPtr ptr, int dx, int dy);
QtObjectPtr QRegion_NewQRegion2(int x, int y, int w, int h, int t);
int QRegion_Intersects(QtObjectPtr ptr, QtObjectPtr region);
void QRegion_Swap(QtObjectPtr ptr, QtObjectPtr other);
void QRegion_Translate2(QtObjectPtr ptr, QtObjectPtr point);

#ifdef __cplusplus
}
#endif