import React from 'react';
import {Card, Table} from 'antd';
import $ from "jquery";

class HwConfigTable extends React.Component {
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



export default class HardwareConfiguration extends React.Component {
    constructor(props) {
        super(props);

        this.handlePartNumberChange = this.handlePartNumberChange.bind(this);
        this.handleRevisionChange = this.handleRevisionChange.bind(this);
        this.handleSerialNumberChange = this.handleSerialNumberChange.bind(this);

        this.handleUpdateButtonClick = this.handleUpdateButtonClick.bind(this);

        this.state = {
            current: {
                id: "",
                name: "",
                partNumber: "",
                revision: "",
                serialNumber: "",
            },
            lastKnown: {
                id: "",
                name: "",
                partNumber: "",
                revision: "",
                serialNumber: "",
            },


        };
    }

    handlePartNumberChange(e) {
        var c = this.state.current;
        c.partNumber = e.target.value;
        this.setState({
            current: c,
        });
    }

    handleRevisionChange(e) {
        var c = this.state.current;
        c.revision = e.target.value;
        this.setState({
            current: c,
        });
    }

    handleSerialNumberChange(e) {
        var c = this.state.current;
        c.serialNumber = e.target.value;
        this.setState({
            current: c,
        });
    }


    handleUpdateButtonClick(e) {
        var url = "/sethwconfig";
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
    componentDidMount() {
        var currentData = this.getHwConfig(0);
        var lastKnownData = this.getHwConfig(2);
        console.log("hardware current: ", currentData);
        console.log("hardware last known: ", lastKnownData);
        this.setState({
            current: {
                id: currentData.Id,
                name: currentData.Name,
                partNumber: currentData.PartNumber,
                revision: currentData.Revision,
                serialNumber: currentData.SerialNumber,
            },
            lastKnown: {
                id: lastKnownData.Id,
                name: lastKnownData.Name,
                partNumber: lastKnownData.PartNumber,
                revision: lastKnownData.Revision,
                serialNumber: lastKnownData.SerialNumber,
            },

        });
    }

    render(){
        return (
            <Card title="Hardware Configuration">
                <HwConfigTable/>
            </Card>
        );
    }

/*
    render() {
        return (
            <div className="panel panel-default">
                <div className="panel-heading">Hardware Configuration</div>
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
                                <th scope="row">Part Number</th>
                                <td><input type="text" value={this.state.current.partNumber} className="form-control" onChange={this.handlePartNumberChange}></input></td>
                                <td>{this.state.lastKnown.partNumber}</td>
                            </tr>
                            <tr>
                                <th scope="row">Revision</th>
                                <td><input type="text" value={this.state.current.revision} className="form-control" onChange={this.handleRevisionChange}></input></td>
                                <td>{this.state.lastKnown.revision}</td>
                            </tr>
                            <tr>
                                <th scope="row">Serial Number</th>
                                <td><input type="text" value={this.state.current.serialNumber} className="form-control" onChange={this.handleSerialNumberChange}></input></td>
                                <td>{this.state.lastKnown.serialNumber}</td>
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