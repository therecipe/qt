#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QPolygon_NewQPolygon5(QtObjectPtr rectangle, int closed);
int QPolygon_ContainsPoint(QtObjectPtr ptr, QtObjectPtr point, int fillRule);
void QPolygon_PutPoints3(QtObjectPtr ptr, int index, int nPoints, QtObjectPtr fromPolygon, int fromIndex);
QtObjectPtr QPolygon_NewQPolygon();
QtObjectPtr QPolygon_NewQPolygon3(QtObjectPtr polygon);
QtObjectPtr QPolygon_NewQPolygon2(int size);
void QPolygon_Point(QtObjectPtr ptr, int index, int x, int y);
void QPolygon_SetPoint2(QtObjectPtr ptr, int index, QtObjectPtr point);
void QPolygon_SetPoint(QtObjectPtr ptr, int index, int x, int y);
void QPolygon_SetPoints(QtObjectPtr ptr, int nPoints, int points);
void QPolygon_Swap(QtObjectPtr ptr, QtObjectPtr other);
void QPolygon_Translate2(QtObjectPtr ptr, QtObjectPtr offset);
void QPolygon_Translate(QtObjectPtr ptr, int dx, int dy);
void QPolygon_DestroyQPolygon(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif