import React from "react";
import ReactDOM from "react-dom";
import RubyMenu from "./menu";
import DeviceConfig from "./deviceconfig";
import DeviceLog from "./log";
import UpgradePage from "./upgradepage";
import Calibration from "./calibration";
import Command from "./command";

var menuDIV = document.getElementById("rubyMenu");
console.log("menu DIV : ", menuDIV);
if (menuDIV != null) {
    ReactDOM.render(
        <RubyMenu/>, menuDIV);
}

var logDIV = document.getElementById("log");
console.log("log DIV : ", logDIV);
if (logDIV != null) {
    ReactDOM.render(
        <DeviceLog/>, logDIV);
}

var upgradeDIV = document.getElementById("upgrade");
console.log("upgrade DIV : ", upgradeDIV);
if (upgradeDIV != null) {
    ReactDOM.render(
        <UpgradePage/>, upgradeDIV);
}

var configDIV = document.getElementById("configuration");
console.log("config DIV : ", configDIV);
if (configDIV != null) {
    ReactDOM.render(
        <DeviceConfig/>, configDIV);
}

var calibrationDIV = document.getElementById("calibration");
console.log("calibration DIV : ", calibrationDIV);
if (calibrationDIV != null) {
    ReactDOM.render(
        <Calibration/>, calibrationDIV);
}

var commandDIV = document.getElementById("command");
console.log("command DIV : ", commandDIV);
if (commandDIV != null) {
    ReactDOM.render(
        <Command/>, commandDIV);
}