import React from 'react';
import {Card, Table, Button} from 'antd';
import $ from "jquery";
import EditableCell from "./editablecell.jsx";

class HwConfigTable extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            loading: false,
            currentID: 0,
            lastKnownID: 0,
            dataSource: [
                {
                    key: 0,
                    ItemName: "Part Number",
                    Current: "",
                    LastKnown: ""
                }, {
                    key: 1,
                    ItemName: "Revision",
                    Current: "",
                    LastKnown: ""
                }, {
                    key: 2,
                    ItemName: "Serial Number",
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
        let url = "/sethwconfig";
        var data = {
            id: this.state.currentID,
            partNumber: this.state.dataSource[0].Current,
            revision: this.state.dataSource[1].Current,
            serialNumber: this.state.dataSource[2].Current
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

    getHwConfig(block) {
        var url = "/gethwconfig";
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
        var currentData = this.getHwConfig(0);
        console.log("hw current: ", currentData);

        if (currentData != null) {
            const dataSource = this.state.dataSource;
            this.setState({currentID: currentData.ID});
            dataSource[0].Current = currentData.PartNumber;
            dataSource[1].Current = currentData.Revision;
            dataSource[2].Current = currentData.SerialNumber;
            this.setState({dataSource});
        }
        var lastKnownData = this.getHwConfig(2);
        console.log("hw last known: ", lastKnownData);
        if (lastKnownData != null) {
            const dataSource = this.state.dataSource;
            this.setState({lastKnownID: lastKnownData.ID});
            dataSource[0].LastKnown = lastKnownData.PartNumber;
            dataSource[1].LastKnown = lastKnownData.Revision;
            dataSource[2].LastKnown = lastKnownData.SerialNumber;
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

export default class HardwareConfiguration extends React.Component {
    render() {
        return (
            <Card title="Hardware Configuration">
                <HwConfigTable/>
            </Card>
        );
    }

}