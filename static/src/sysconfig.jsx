import React from 'react';
import {Card, Table} from 'antd';
import $ from "jquery";

class SysConfigTable extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            data: [],
            loading: false
        };
        this.handleTableChange = this
            .handleTableChange
            .bind(this);
    }

    handleTableChange(pagination, filters, sorter) {
        console.log("pagination : ", pagination);
        console.log("filters :", filters);
        console.log("sorter : ", sorter);
    }

    render() {
                const columns = [
            {
                title: '#',
                key: "Items",
                dataIndex: 'Items',
            }, {
                title: 'Current',
                key: "Current",
                dataIndex: 'Current'
            }, {
                title: 'Last Known',
                key: "LastKnown",
                dataIndex: 'LastKnown'
            }
        ];
        return (
            <Table
                dataSource={this.state.data}
                columns={columns}
                loading={this.state.loading}
                onChange={this.handleTableChange}></Table>
        );
    }
}

export default class SystemConfiguration extends React.Component {
    constructor(props) {
        super(props);

        this.handleDeviceNameChange = this
            .handleDeviceNameChange
            .bind(this);
        this.handleSystemVersionChange = this
            .handleSystemVersionChange
            .bind(this);
        this.handleDeviceSKUChange = this
            .handleDeviceSKUChange
            .bind(this);
        this.handleSerialNumberChange = this
            .handleSerialNumberChange
            .bind(this);
        this.handleSoftwareBuildChange = this
            .handleSoftwareBuildChange
            .bind(this);
        this.handlePartNumberChange = this
            .handlePartNumberChange
            .bind(this);
        this.handleHardwareVersionChange = this
            .handleHardwareVersionChange
            .bind(this);

        this.handleUpdateButtonClick = this
            .handleUpdateButtonClick
            .bind(this);

        this.state = {
            current: {
                id: "",
                deviceName: "",
                systemVersion: "",
                deviceSKU: "",
                serialNumber: "",
                softwareBuild: "",
                partNumber: "",
                hardwareVersion: ""
            },
            lastKnown: {
                id: "",
                deviceName: "",
                systemVersion: "",
                deviceSKU: "",
                serialNumber: "",
                softwareBuild: "",
                partNumber: "",
                hardwareVersion: ""
            }
        };
    }
    handleDeviceNameChange(e) {
        var c = this.state.current;
        c.deviceName = e.target.value;
        this.setState({current: c});
    }
    handleSystemVersionChange(e) {
        var c = this.state.current;
        c.systemVersion = e.target.value;
        this.setState({current: c});
    }

    handleDeviceSKUChange(e) {
        var c = this.state.current;
        c.deviceSKU = e.target.value;
        this.setState({current: c});
    }

    handleSerialNumberChange(e) {
        var c = this.state.current;
        c.serialNumber = e.target.value;
        this.setState({current: c});
    }

    handleSoftwareBuildChange(e) {
        var c = this.state.current;
        c.softwareBuild = e.target.value;
        this.setState({current: c});
    }

    handlePartNumberChange(e) {
        var c = this.state.current;
        c.partNumber = e.target.value;
        this.setState({current: c});

    }

    handleHardwareVersionChange(e) {
        var c = this.state.current;
        c.hardwareVersion = e.target.value;
        this.setState({current: c});
    }

    handleUpdateButtonClick(e) {
        var url = "/setsysconfig";
        console.log("state:", this.state.current);
        $.ajax({
            url: url,
            dataType: "json",
            type: "POST",
            cache: false,
            async: false,
            data: this.state.current,
            success: function (data) {
                console.log(data);
                //result = data;
            },
            error: function (xhr, status, err) {
                console.error(url, status, err.toString());
            }
        });

    }

    getSysConfig(block) {
        var url = "/getsysconfig";
        var result = null;

        $.ajax({
            url: url,
            dataType: "json",
            type: "POST",
            cache: false,
            async: false,
            data: {
                block: block
            },
            success: function (data) {
                //            console.log(data);
                result = data;

            },
            error: function (xhr, status, err) {
                console.error(url, status, err.toString());
            }
        });
        return result;
    }

    componentDidMount() {
        var currentData = this.getSysConfig(0);
        var lastKnownData = this.getSysConfig(2);
        console.log("system current: ", currentData);
        console.log("system last known: ", lastKnownData);
        this.setState({
            current: {
                id: currentData.ID,
                deviceName: currentData.DeviceName,
                systemVersion: currentData.SystemVersion,
                deviceSKU: currentData.DeviceSKU,
                serialNumber: currentData.SerialNumber,
                softwareBuild: currentData.SoftwareBuild,
                partNumber: currentData.PartNumber,
                hardwareVersion: currentData.HardwareVersion
            },
            lastKnown: {
                id: lastKnownData.ID,
                deviceName: lastKnownData.DeviceName,
                systemVersion: lastKnownData.SystemVersion,
                deviceSKU: lastKnownData.DeviceSKU,
                serialNumber: lastKnownData.SerialNumber,
                softwareBuild: lastKnownData.SoftwareBuild,
                partNumber: lastKnownData.PartNumber,
                hardwareVersion: lastKnownData.HardwareVersion
            }
        });
    }

    render() {
        return (
            <Card title="System Configuration">
                <SysConfigTable/>
            </Card>
        );
    }

    /*
    render() {
        return (
            <div className="panel panel-default">
                <div className="panel-heading">System Configuration</div>
                <div className="panel-body">
                    <table className="table table-bordered table-hover">
                        <thead>
                            <tr>
                                <th>#</th>
                                <th>Current</th>
                                <th>Last Known</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr>
                                <th scope="row">Device Name</th>
                                <td><input type="text" value={this.state.current.deviceName} className="form-control" onChange={this.handleDeviceNameChange}></input></td>
                                <td>{this.state.lastKnown.deviceName}</td>
                            </tr>
                            <tr>
                                <th scope="row">Serial Number</th>
                                <td><input type="text" value={this.state.current.serialNumber} className="form-control" onChange={this.handleSerialNumberChange}></input></td>
                                <td>{this.state.lastKnown.serialNumber}</td>
                            </tr>
                            <tr>
                                <th scope="row">System Version</th>
                                <td><input type="text" value={this.state.current.systemVersion} className="form-control" onChange={this.handleSystemVersionChange}></input></td>
                                <td>{this.state.lastKnown.systemVersion}</td>
                            </tr>
                            <tr>
                                <th scope="row">Device SKU</th>
                                <td><input type="text" value={this.state.current.deviceSKU} className="form-control" onChange={this.handleDeviceSKUChange}></input></td>
                                <td>{this.state.lastKnown.deviceSKU}</td>
                            </tr>
                            <tr>
                                <th scope="row">Software Build</th>
                                <td><input type="text" value={this.state.current.softwareBuild} className="form-control" onChange={this.handleSoftwareBuildChange}></input></td>
                                <td>{this.state.lastKnown.softwareBuild}</td>
                            </tr>
                            <tr>
                                <th scope="row">Part Number</th>
                                <td><input type="text" value={this.state.current.partNumber} className="form-control" onChange={this.handlePartNumberChange}></input></td>
                                <td>{this.state.lastKnown.partNumber}</td>
                            </tr>
                            <tr>
                                <th scope="row">Hardware Version</th>
                                <td><input type="text" value={this.state.current.hardwareVersion} className="form-control" onChange={this.handleHardwareVersionChange}></input></td>
                                <td>{this.state.lastKnown.hardwareVersion}</td>
                            </tr>

                        </tbody>
                    </table>
                </div>
                <div className="panel-footer">
                    <input type="button" value="Edit" onClick={this.handleUpdateButtonClick}></input>
                </div>
            </div>
        );
    }
    */
}