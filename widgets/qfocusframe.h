#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QFocusFrame_NewQFocusFrame(QtObjectPtr parent);
void QFocusFrame_SetWidget(QtObjectPtr ptr, QtObjectPtr widget);
QtObjectPtr QFocusFrame_Widget(QtObjectPtr ptr);
void QFocusFrame_DestroyQFocusFrame(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif