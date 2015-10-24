#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QScrollBar_NewQScrollBar(QtObjectPtr parent);
QtObjectPtr QScrollBar_NewQScrollBar2(int orientation, QtObjectPtr parent);
int QScrollBar_Event(QtObjectPtr ptr, QtObjectPtr event);
void QScrollBar_DestroyQScrollBar(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif