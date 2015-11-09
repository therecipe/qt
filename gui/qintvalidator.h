#ifdef __cplusplus
extern "C" {
#endif

void QIntValidator_SetBottom(void* ptr, int v);
void QIntValidator_SetTop(void* ptr, int v);
void* QIntValidator_NewQIntValidator(void* parent);
void* QIntValidator_NewQIntValidator2(int minimum, int maximum, void* parent);
int QIntValidator_Bottom(void* ptr);
void QIntValidator_SetRange(void* ptr, int bottom, int top);
int QIntValidator_Top(void* ptr);
void QIntValidator_DestroyQIntValidator(void* ptr);

#ifdef __cplusplus
}
#endif