#ifdef __cplusplus
extern "C" {
#endif

char* QMediaContainerControl_ContainerDescription(void* ptr, char* format);
char* QMediaContainerControl_ContainerFormat(void* ptr);
void QMediaContainerControl_SetContainerFormat(void* ptr, char* format);
char* QMediaContainerControl_SupportedContainers(void* ptr);
void QMediaContainerControl_DestroyQMediaContainerControl(void* ptr);

#ifdef __cplusplus
}
#endif