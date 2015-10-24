#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QSvgWidget_NewQSvgWidget(QtObjectPtr parent);
QtObjectPtr QSvgWidget_NewQSvgWidget2(char* file, QtObjectPtr parent);
void QSvgWidget_Load2(QtObjectPtr ptr, QtObjectPtr contents);
void QSvgWidget_Load(QtObjectPtr ptr, char* file);
QtObjectPtr QSvgWidget_Renderer(QtObjectPtr ptr);
void QSvgWidget_DestroyQSvgWidget(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif