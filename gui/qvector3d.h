#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QVector3D_NewQVector3D();
QtObjectPtr QVector3D_NewQVector3D4(QtObjectPtr point);
QtObjectPtr QVector3D_NewQVector3D5(QtObjectPtr point);
QtObjectPtr QVector3D_NewQVector3D6(QtObjectPtr vector);
QtObjectPtr QVector3D_NewQVector3D8(QtObjectPtr vector);
int QVector3D_IsNull(QtObjectPtr ptr);
void QVector3D_Normalize(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif