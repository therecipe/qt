#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QPoint_NewQPoint();
QtObjectPtr QPoint_NewQPoint2(int xpos, int ypos);
int QPoint_QPoint_DotProduct(QtObjectPtr p1, QtObjectPtr p2);
int QPoint_IsNull(QtObjectPtr ptr);
int QPoint_ManhattanLength(QtObjectPtr ptr);
int QPoint_Rx(QtObjectPtr ptr);
int QPoint_Ry(QtObjectPtr ptr);
void QPoint_SetX(QtObjectPtr ptr, int x);
void QPoint_SetY(QtObjectPtr ptr, int y);
int QPoint_X(QtObjectPtr ptr);
int QPoint_Y(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif