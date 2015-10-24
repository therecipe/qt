#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QLineF_Intersect(QtObjectPtr ptr, QtObjectPtr line, QtObjectPtr intersectionPoint);
QtObjectPtr QLineF_NewQLineF();
QtObjectPtr QLineF_NewQLineF4(QtObjectPtr line);
QtObjectPtr QLineF_NewQLineF2(QtObjectPtr p1, QtObjectPtr p2);
int QLineF_IsNull(QtObjectPtr ptr);
void QLineF_SetP1(QtObjectPtr ptr, QtObjectPtr p1);
void QLineF_SetP2(QtObjectPtr ptr, QtObjectPtr p2);
void QLineF_SetPoints(QtObjectPtr ptr, QtObjectPtr p1, QtObjectPtr p2);
void QLineF_Translate(QtObjectPtr ptr, QtObjectPtr offset);

#ifdef __cplusplus
}
#endif