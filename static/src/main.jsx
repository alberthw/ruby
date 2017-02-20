import React from "react";
import ReactDOM from "react-dom";
import RubyMenu from "./menu.jsx";
import DeviceConfig from "./deviceconfig.jsx"
import DeviceLog from "./log.jsx";
import UpgradePage from "./upgradepage.jsx";

// ReactDOM.render(<DeviceConfiguration />,
// document.getElementById("configuration"));

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
        <DeviceConfig />, configDIV);
}
