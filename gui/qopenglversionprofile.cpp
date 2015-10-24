#include "qopenglversionprofile.h"
#include <QSurfaceFormat>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSurface>
#include <QOpenGLVersionProfile>
#include "_cgo_export.h"

class MyQOpenGLVersionProfile: public QOpenGLVersionProfile {
public:
};

QtObjectPtr QOpenGLVersionProfile_NewQOpenGLVersionProfile(){
	return new QOpenGLVersionProfile();
}

QtObjectPtr QOpenGLVersionProfile_NewQOpenGLVersionProfile3(QtObjectPtr other){
	return new QOpenGLVersionProfile(*static_cast<QOpenGLVersionProfile*>(other));
}

QtObjectPtr QOpenGLVersionProfile_NewQOpenGLVersionProfile2(QtObjectPtr format){
	return new QOpenGLVersionProfile(*static_cast<QSurfaceFormat*>(format));
}

int QOpenGLVersionProfile_HasProfiles(QtObjectPtr ptr){
	return static_cast<QOpenGLVersionProfile*>(ptr)->hasProfiles();
}

int QOpenGLVersionProfile_IsLegacyVersion(QtObjectPtr ptr){
	return static_cast<QOpenGLVersionProfile*>(ptr)->isLegacyVersion();
}

int QOpenGLVersionProfile_IsValid(QtObjectPtr ptr){
	return static_cast<QOpenGLVersionProfile*>(ptr)->isValid();
}

int QOpenGLVersionProfile_Profile(QtObjectPtr ptr){
	return static_cast<QOpenGLVersionProfile*>(ptr)->profile();
}

void QOpenGLVersionProfile_SetProfile(QtObjectPtr ptr, int profile){
	static_cast<QOpenGLVersionProfile*>(ptr)->setProfile(static_cast<QSurfaceFormat::OpenGLContextProfile>(profile));
}

void QOpenGLVersionProfile_SetVersion(QtObjectPtr ptr, int majorVersion, int minorVersion){
	static_cast<QOpenGLVersionProfile*>(ptr)->setVersion(majorVersion, minorVersion);
}

void QOpenGLVersionProfile_DestroyQOpenGLVersionProfile(QtObjectPtr ptr){
	static_cast<QOpenGLVersionProfile*>(ptr)->~QOpenGLVersionProfile();
}

