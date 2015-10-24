#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QPicture_IsNull(QtObjectPtr ptr);
int QPicture_Load2(QtObjectPtr ptr, QtObjectPtr dev, char* format);
int QPicture_Load(QtObjectPtr ptr, char* fileName, char* format);
int QPicture_Play(QtObjectPtr ptr, QtObjectPtr painter);
int QPicture_Save2(QtObjectPtr ptr, QtObjectPtr dev, char* format);
int QPicture_Save(QtObjectPtr ptr, char* fileName, char* format);
void QPicture_SetBoundingRect(QtObjectPtr ptr, QtObjectPtr r);
void QPicture_Swap(QtObjectPtr ptr, QtObjectPtr other);
void QPicture_DestroyQPicture(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif