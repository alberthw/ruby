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
                <Menu.Item>Ruby Client</Menu.Item>
                <Menu.Item key="upgrade">
                    <a href="/">Upgrade</a>
                </Menu.Item>
                <Menu.Item key="configuration">
                    <a href="/device">Configuration</a>
                </Menu.Item>
                <Menu.Item key="log">
                    <a href="/log">Log</a>
                </Menu.Item>
            </Menu>

        );
    }
}
