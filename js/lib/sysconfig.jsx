import React from 'react';
import {Card, Table, Button} from 'antd';
import $ from "jquery";
import EditableCell from "./editablecell";

class SysConfigTable extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            loading: false,
            currentID: 0,
            lastKnownID: 0,
            dataSource: [
                {
                    key: 0,
                    ItemName: "Device Name",
                    Current: "",
                    LastKnown: ""
                }, {
                    key: 1,
                    ItemName: "System Version",
                    Current: "",
                    LastKnown: ""
                }, {
                    key: 2,
                    ItemName: "Device SKU",
                    Current: "",
                    LastKnown: ""
                }, {
                    key: 3,
                    ItemName: "Serial Number",
                    Current: "",
                    LastKnown: ""
                }, {
                    key: 4,
                    ItemName: "Software Build",
                    Current: "",
                    LastKnown: ""
                }, {
                    key: 5,
                    ItemName: "Part Number",
                    Current: "",
                    LastKnown: ""
                }, {
                    key: 6,
                    ItemName: "Hardware Version",
                    Current: "",
                    LastKnown: ""
                }
            ]
        };

        this.handleTableChange = this
            .handleTableChange
            .bind(this);

        this.handleUpdateButtonClick = this
            .handleUpdateButtonClick
            .bind(this);

    }

    handleUpdateButtonClick(e) {
        let url = "/setsysconfig";
        var data = {
            id: this.state.currentID,
            deviceName: this.state.dataSource[0].Current,
            systemVersion: this.state.dataSource[1].Current,
            deviceSKU: this.state.dataSource[2].Current,
            serialNumber: this.state.dataSource[3].Current,
            softwareBuild: this.state.dataSource[4].Current,
            partNumber: this.state.dataSource[5].Current,
            hardwareVersion: this.state.dataSource[6].Current
        };
        $.ajax({
            url: url,
            dataType: "json",
            type: "POST",
            cache: false,
            async: false,
            data: data,
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

    handleTableChange(pagination, filters, sorter) {
        console.log("pagination : ", pagination);
        console.log("filters :", filters);
        console.log("sorter : ", sorter);
    }

    componentWillMount() {
        var currentData = this.getSysConfig(0);
        console.log("system current: ", currentData);

        if (currentData != null) {
            const dataSource = this.state.dataSource;
            this.setState({currentID: currentData.ID});
            dataSource[0].Current = currentData.DeviceName;
            dataSource[1].Current = currentData.SystemVersion;
            dataSource[2].Current = currentData.DeviceSKU;
            dataSource[3].Current = currentData.SerialNumber;
            dataSource[4].Current = currentData.SoftwareBuild;
            dataSource[5].Current = currentData.PartNumber;
            dataSource[6].Current = currentData.HardwareVersion;
            this.setState({dataSource});
        }
        var lastKnownData = this.getSysConfig(2);
        console.log("system last known: ", lastKnownData);
        if (lastKnownData != null) {
            const dataSource = this.state.dataSource;
            this.setState({lastKnownID: lastKnownData.ID});
            dataSource[0].LastKnown = lastKnownData.DeviceName;
            dataSource[1].LastKnown = lastKnownData.SystemVersion;
            dataSource[2].LastKnown = lastKnownData.DeviceSKU;
            dataSource[3].LastKnown = lastKnownData.SerialNumber;
            dataSource[4].LastKnown = lastKnownData.SoftwareBuild;
            dataSource[5].LastKnown = lastKnownData.PartNumber;
            dataSource[6].LastKnown = lastKnownData.HardwareVersion;
            this.setState({dataSource});
        }
    }

    onCellChange = (index, key) => {
        return (value) => {
            const dataSource = this.state.dataSource;
            //           console.log("dataSource:", dataSource);
            dataSource[index][key] = value;
            this.setState({dataSource});
        }
    }

    render() {
        const pagination = false;

        const columns = [
            {
                title: '#',
                key: "ItemName",
                dataIndex: 'ItemName',
                render: text => <b>{text}</b>
            }, {
                title: 'Current',
                key: "Current",
                dataIndex: 'Current',
                render: (text, record, index) => (<EditableCell value={text} onChange={this.onCellChange(index, 'Current')}/>)
            }, {
                title: 'Last Known',
                key: "LastKnown",
                dataIndex: 'LastKnown'
            }
        ];
        return (
            <div>
                <Table
                    pagination={pagination}
                    dataSource={this.state.dataSource}
                    columns={columns}
                    loading={this.state.loading}
                    onChange={this.handleTableChange}
                    bordered></Table >
                <Button onClick={this.handleUpdateButtonClick}>Update</Button>
            </div>

        );
    }
}

export default class SystemConfiguration extends React.Component {
    render() {
        return (
            <Card title="System Configuration">
                <SysConfigTable/>
            </Card>
        );
    }
}