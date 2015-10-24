#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QSplitterHandle_NewQSplitterHandle(int orientation, QtObjectPtr parent);
int QSplitterHandle_OpaqueResize(QtObjectPtr ptr);
int QSplitterHandle_Orientation(QtObjectPtr ptr);
void QSplitterHandle_SetOrientation(QtObjectPtr ptr, int orientation);
QtObjectPtr QSplitterHandle_Splitter(QtObjectPtr ptr);
void QSplitterHandle_DestroyQSplitterHandle(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif