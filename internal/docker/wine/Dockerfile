FROM therecipe/qt:wine_base as base

RUN QT=qt-opensource-windows-x86-5.11.1.exe && curl -sL --retry 10 --retry-delay 10 -O https://download.qt.io/official_releases/qt/5.11/5.11.1/$QT \
	&& QT_QPA_PLATFORM=minimal xvfb-run wine Z:\\$QT --script=C:\\gopath\\src\\github.com\\therecipe\\qt\\internal\\ci\\iscript.qs WINDOWS=true && rm -f $QT
RUN find $HOME/.wine/drive_c/Qt/Qt5.11.1/5.11.1/mingw53_32 -type f -name "*d.a" ! -name "*3d.a" ! -name "*pad.a" -delete
RUN find $HOME/.wine/drive_c/Qt/Qt5.11.1/5.11.1/mingw53_32 -type f -name "*d.dll" ! -name "3*d.dll" ! -name "*pad.dll" -delete
RUN find $HOME/.wine/drive_c/Qt/Qt5.11.1/Docs -type f ! -name "*.index" -delete


FROM ubuntu:16.04
LABEL maintainer therecipe

ENV USER user
ENV HOME /home/$USER
ENV GOPATH C:\\gopath
ENV WINEDEBUG -all
ENV WINEPATH C:\\Qt\\Qt5.11.1\\Tools\\mingw530_32\\bin;C:\\go\\bin

COPY --from=base $HOME/.wine/drive_c/go $HOME/.wine/drive_c/go
COPY --from=base $HOME/.wine/drive_c/gopath/bin $HOME/.wine/drive_c/gopath/bin
COPY --from=base $HOME/.wine/drive_c/gopath/src/github.com/therecipe/qt $HOME/.wine/drive_c/gopath/src/github.com/therecipe/qt
COPY --from=base $HOME/.wine/drive_c/Qt/Qt5.11.1/5.11.1/mingw53_32 $HOME/.wine/drive_c/Qt/Qt5.11.1/5.11.1/mingw53_32
COPY --from=base $HOME/.wine/drive_c/Qt/Qt5.11.1/Docs $HOME/.wine/drive_c/Qt/Qt5.11.1/Docs
COPY --from=base $HOME/.wine/drive_c/Qt/Qt5.11.1/Tools/mingw530_32 $HOME/.wine/drive_c/Qt/Qt5.11.1/Tools/mingw530_32

RUN apt-get -qq update && apt-get --no-install-recommends -qq -y install ca-certificates curl software-properties-common apt-transport-https && apt-get -qq clean
RUN dpkg --add-architecture i386
RUN curl -sL --retry 10 --retry-delay 10 https://dl.winehq.org/wine-builds/Release.key | apt-key add -
RUN apt-add-repository https://dl.winehq.org/wine-builds/ubuntu/
RUN apt-add-repository 'deb https://dl.winehq.org/wine-builds/ubuntu/ xenial main'
RUN apt-get -qq update && apt-get --no-install-recommends -qq -y install winehq-staging && apt-get -qq clean
RUN WINEDLLOVERRIDES="mscoree,mshtml=" wineboot && wineserver -w

RUN echo 'REGEDIT4\n\n\
[HKEY_CURRENT_USER\Environment]\n\
"QT_DOCKER"="true"' > new.reg && regedit Z:\\new.reg && wineserver -w

RUN wine $GOPATH\\bin\\qtsetup prep
RUN wine $GOPATH\\bin\\qtsetup check
RUN wine $GOPATH\\bin\\qtsetup generate
RUN cd $HOME/.wine/drive_c/gopath/src/github.com/therecipe/qt/internal/examples/widgets/line_edits && wine qtdeploy build windows && rm -rf ./deploy

RUN echo "#!/bin/bash\nexport GOPATH=\$(winepath -0 -w \$(echo \$GOPATH | sed -e 's/:/ /g') | sed -e 's/Z:/;Z:/2g')" > /usr/bin/qtpath && chmod +x /usr/bin/qtpath
RUN echo '#!/bin/bash\nsource qtpath\nwine qtdeploy "$@"' > /usr/bin/qtdeploy && chmod +x /usr/bin/qtdeploy
RUN echo '#!/bin/bash\nsource qtpath\nwine qtminimal "$@"' > /usr/bin/qtminimal && chmod +x /usr/bin/qtminimal
RUN echo '#!/bin/bash\nsource qtpath\nwine qtmoc "$@"' > /usr/bin/qtmoc && chmod +x /usr/bin/qtmoc
RUN echo '#!/bin/bash\nsource qtpath\nwine qtrcc "$@"' > /usr/bin/qtrcc && chmod +x /usr/bin/qtrcc
RUN ln -s $HOME/.wine/drive_c/gopath $HOME/work
