#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

char* QMediaContainerControl_ContainerDescription(QtObjectPtr ptr, char* format);
char* QMediaContainerControl_ContainerFormat(QtObjectPtr ptr);
void QMediaContainerControl_SetContainerFormat(QtObjectPtr ptr, char* format);
char* QMediaContainerControl_SupportedContainers(QtObjectPtr ptr);
void QMediaContainerControl_DestroyQMediaContainerControl(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif