#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QVector4D_NewQVector4D();
QtObjectPtr QVector4D_NewQVector4D4(QtObjectPtr point);
QtObjectPtr QVector4D_NewQVector4D5(QtObjectPtr point);
QtObjectPtr QVector4D_NewQVector4D6(QtObjectPtr vector);
QtObjectPtr QVector4D_NewQVector4D8(QtObjectPtr vector);
int QVector4D_IsNull(QtObjectPtr ptr);
void QVector4D_Normalize(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif