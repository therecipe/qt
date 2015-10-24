#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QLine_NewQLine();
QtObjectPtr QLine_NewQLine2(QtObjectPtr p1, QtObjectPtr p2);
QtObjectPtr QLine_NewQLine3(int x1, int y1, int x2, int y2);
int QLine_Dx(QtObjectPtr ptr);
int QLine_Dy(QtObjectPtr ptr);
int QLine_IsNull(QtObjectPtr ptr);
void QLine_SetLine(QtObjectPtr ptr, int x1, int y1, int x2, int y2);
void QLine_SetP1(QtObjectPtr ptr, QtObjectPtr p1);
void QLine_SetP2(QtObjectPtr ptr, QtObjectPtr p2);
void QLine_SetPoints(QtObjectPtr ptr, QtObjectPtr p1, QtObjectPtr p2);
void QLine_Translate(QtObjectPtr ptr, QtObjectPtr offset);
void QLine_Translate2(QtObjectPtr ptr, int dx, int dy);
int QLine_X1(QtObjectPtr ptr);
int QLine_X2(QtObjectPtr ptr);
int QLine_Y1(QtObjectPtr ptr);
int QLine_Y2(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif