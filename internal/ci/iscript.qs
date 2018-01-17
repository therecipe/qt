function Controller()
{
  installer.wizardPageInsertionRequested.connect(function(widget, page)
  {
    installer.removeWizardPage(installer.components()[0], "WorkspaceWidget");
  })

  installer.autoRejectMessageBoxes();
  installer.installationFinished.connect(function()
  {
    gui.clickButton(buttons.NextButton);
  })
}

Controller.prototype.WelcomePageCallback = function()
{
  gui.clickButton(buttons.NextButton, 30000);
}

Controller.prototype.CredentialsPageCallback = function()
{
  gui.clickButton(buttons.NextButton);
}

Controller.prototype.IntroductionPageCallback = function()
{
  gui.clickButton(buttons.NextButton);
}

Controller.prototype.TargetDirectoryPageCallback = function()
{
  gui.clickButton(buttons.NextButton);
}

Controller.prototype.ComponentSelectionPageCallback = function()
{
  if (installer.value("DARWIN") == "true")
  {
    gui.currentPageWidget().selectComponent("qt.qt5.5100.clang_64");
  } 

  if (installer.value("IOS") == "true")
  {
    gui.currentPageWidget().selectComponent("qt.qt5.5100.ios");
  } 

  if (installer.value("WINDOWS") == "true")
  {
    gui.currentPageWidget().selectComponent("qt.qt5.5100.win32_mingw53");
  }

  if (installer.value("LINUX") == "true")
  {
    gui.currentPageWidget().selectComponent("qt.qt5.5100.gcc_64");
  }

  gui.currentPageWidget().selectComponent("qt.qt5.5100.qtvirtualkeyboard");
  gui.currentPageWidget().selectComponent("qt.qt5.5100.qtscript");
  gui.currentPageWidget().selectComponent("qt.qt5.5100.qtremoteobjects");
  gui.currentPageWidget().selectComponent("qt.qt5.5100.qtpurchasing");
  gui.currentPageWidget().selectComponent("qt.qt5.5100.qtnetworkauth");
  gui.currentPageWidget().selectComponent("qt.qt5.5100.qtdatavis3d");
  gui.currentPageWidget().selectComponent("qt.qt5.5100.qtcharts");
  gui.currentPageWidget().selectComponent("qt.qt5.5100.qtwebengine");
  gui.currentPageWidget().selectComponent("qt.qt5.5100.qtwebglplugin");

  gui.currentPageWidget().selectComponent("qt.qt5.5100.android_armv7");
  gui.currentPageWidget().selectComponent("qt.qt5.5100.android_x86");

  gui.clickButton(buttons.NextButton);
}

Controller.prototype.LicenseAgreementPageCallback = function()
{
  gui.currentPageWidget().AcceptLicenseRadioButton.setChecked(true);
  gui.clickButton(buttons.NextButton);
}

Controller.prototype.StartMenuDirectoryPageCallback = function()
{
  gui.clickButton(buttons.NextButton);
}

Controller.prototype.ReadyForInstallationPageCallback = function()
{
  gui.clickButton(buttons.NextButton);
}

Controller.prototype.FinishedPageCallback = function()
{
  var checkBox = gui.currentPageWidget().LaunchQtCreatorCheckBoxForm
  if (checkBox && checkBox.launchQtCreatorCheckBox)
  {
    checkBox.launchQtCreatorCheckBox.checked = false;
  }
  gui.clickButton(buttons.FinishButton);
}
