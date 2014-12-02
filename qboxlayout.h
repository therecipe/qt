#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

//Public Functions
void QBoxLayout_Destroy(QtObjectPtr ptr);
void QBoxLayout_AddLayout_QLayout_Int(QtObjectPtr ptr, QtObjectPtr layout, int stretch);
void QBoxLayout_AddSpacing_Int(QtObjectPtr ptr, int size);
void QBoxLayout_AddStretch_Int(QtObjectPtr ptr, int stretch);
void QBoxLayout_AddStrut_Int(QtObjectPtr ptr, int size);
void QBoxLayout_InsertLayout_Int_QLayout_Int(QtObjectPtr ptr, int index, QtObjectPtr layout, int stretch);
void QBoxLayout_InsertSpacing_Int_Int(QtObjectPtr ptr, int index, int size);
void QBoxLayout_InsertStretch_Int_Int(QtObjectPtr ptr, int index, int stretch);
void QBoxLayout_InsertWidget_Int_QWidget_Int_AlignmentFlag(QtObjectPtr ptr, int index, QtObjectPtr widget, int stretch, int alignment);
void QBoxLayout_SetStretch_Int_Int(QtObjectPtr ptr, int index, int stretch);
int QBoxLayout_SetStretchFactor_QWidget_Int(QtObjectPtr ptr, QtObjectPtr widget, int stretch);
int QBoxLayout_Stretch_Int(QtObjectPtr ptr, int index);

#ifdef __cplusplus
}
#endif
