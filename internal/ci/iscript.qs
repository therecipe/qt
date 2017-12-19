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
  gui.clickButton(buttons.NextButton, 2000);
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
  if (installer.value("IOS") != "true")
  {
    gui.currentPageWidget().deselectComponent("qt.593.ios");
  }

  if (installer.value("WINDOWS") == "true")
  {
    gui.currentPageWidget().selectComponent("qt.593.win32_mingw53");
  }

  if (installer.value("LINUX") == "true")
  {
    gui.currentPageWidget().selectComponent("qt.593.gcc_64");
  }

  gui.currentPageWidget().selectComponent("qt.593.qtvirtualkeyboard");
  gui.currentPageWidget().selectComponent("qt.593.qtspeech");
  gui.currentPageWidget().selectComponent("qt.593.qtscript");
  gui.currentPageWidget().selectComponent("qt.593.qtremoteobjects");
  gui.currentPageWidget().selectComponent("qt.593.qtpurchasing");
  gui.currentPageWidget().selectComponent("qt.593.qtnetworkauth");
  gui.currentPageWidget().selectComponent("qt.593.qtdatavis3d");
  gui.currentPageWidget().selectComponent("qt.593.qtcharts");
  gui.currentPageWidget().selectComponent("qt.593.qtwebengine");

  gui.currentPageWidget().selectComponent("qt.593.android_armv7");
  gui.currentPageWidget().selectComponent("qt.593.android_x86");
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
