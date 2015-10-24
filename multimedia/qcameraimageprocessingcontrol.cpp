#include "qcameraimageprocessingcontrol.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QCamera>
#include <QCameraImageProcessing>
#include <QCameraImageProcessingControl>
#include "_cgo_export.h"

class MyQCameraImageProcessingControl: public QCameraImageProcessingControl {
public:
};

int QCameraImageProcessingControl_IsParameterSupported(QtObjectPtr ptr, int parameter){
	return static_cast<QCameraImageProcessingControl*>(ptr)->isParameterSupported(static_cast<QCameraImageProcessingControl::ProcessingParameter>(parameter));
}

int QCameraImageProcessingControl_IsParameterValueSupported(QtObjectPtr ptr, int parameter, char* value){
	return static_cast<QCameraImageProcessingControl*>(ptr)->isParameterValueSupported(static_cast<QCameraImageProcessingControl::ProcessingParameter>(parameter), QVariant(value));
}

char* QCameraImageProcessingControl_Parameter(QtObjectPtr ptr, int parameter){
	return static_cast<QCameraImageProcessingControl*>(ptr)->parameter(static_cast<QCameraImageProcessingControl::ProcessingParameter>(parameter)).toString().toUtf8().data();
}

void QCameraImageProcessingControl_SetParameter(QtObjectPtr ptr, int parameter, char* value){
	static_cast<QCameraImageProcessingControl*>(ptr)->setParameter(static_cast<QCameraImageProcessingControl::ProcessingParameter>(parameter), QVariant(value));
}

void QCameraImageProcessingControl_DestroyQCameraImageProcessingControl(QtObjectPtr ptr){
	static_cast<QCameraImageProcessingControl*>(ptr)->~QCameraImageProcessingControl();
}

