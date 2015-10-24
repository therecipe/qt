#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QVector2D_NewQVector2D();
QtObjectPtr QVector2D_NewQVector2D4(QtObjectPtr point);
QtObjectPtr QVector2D_NewQVector2D5(QtObjectPtr point);
QtObjectPtr QVector2D_NewQVector2D6(QtObjectPtr vector);
QtObjectPtr QVector2D_NewQVector2D7(QtObjectPtr vector);
int QVector2D_IsNull(QtObjectPtr ptr);
void QVector2D_Normalize(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif