import React from 'react';
import {Menu} from "antd";

export default class RubyMenu extends React.Component {
    constructor(props) {
        super(props);
    }
    render() {
        const MenuItemGroup = Menu.ItemGroup;
        return (
            <Menu mode="horizontal">
                <Menu.Item>
                    <b>Ruby Client</b>
                </Menu.Item>
                <Menu.Item key="upgrade">
                    <a href="/">Upgrade</a>
                </Menu.Item>
                <Menu.Item key="configuration">
                    <a href="/device">Configuration</a>
                </Menu.Item>
                <Menu.Item key="log">
                    <a href="/log">Log</a>
                </Menu.Item>
                <Menu.Item key="calibration">
                    <a href="/calibration">Calibration</a>
                </Menu.Item>
                <Menu.Item key="command">
                    <a href="/command">Command</a>
                </Menu.Item>
            </Menu>
        );
    }
}
