#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QPolygonF_NewQPolygonF6(QtObjectPtr polygon);
QtObjectPtr QPolygonF_NewQPolygonF5(QtObjectPtr rectangle);
int QPolygonF_ContainsPoint(QtObjectPtr ptr, QtObjectPtr point, int fillRule);
QtObjectPtr QPolygonF_NewQPolygonF();
QtObjectPtr QPolygonF_NewQPolygonF3(QtObjectPtr polygon);
QtObjectPtr QPolygonF_NewQPolygonF2(int size);
int QPolygonF_IsClosed(QtObjectPtr ptr);
void QPolygonF_Swap(QtObjectPtr ptr, QtObjectPtr other);
void QPolygonF_Translate(QtObjectPtr ptr, QtObjectPtr offset);
void QPolygonF_DestroyQPolygonF(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif