#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

//Public Functions
QtObjectPtr QRadioButton_New_QWidget(QtObjectPtr parent);
QtObjectPtr QRadioButton_New_String_QWidget(char* text, QtObjectPtr parent);
void QRadioButton_Destroy(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif
