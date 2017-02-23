import React from 'react';
import {Row, Col, Card, Table} from 'antd';
import $ from "jquery";

export default class SoftwareConfiguration extends React.Component {
    constructor(props) {
        super(props);
    }
    render() {
        return (
            <Card title="Software Configuration">
                <Row gutter={8}>
                    <Col span="8">
                        <SoftwareComponent name="Host Boot Loader" type="0"/>
                    </Col>
                    <Col span="8">
                        <SoftwareComponent name="Host Application" type="1"/>
                    </Col>
                    <Col span="8">
                        <SoftwareComponent name="DSP Application" type="2"/>
                    </Col>
                </Row>
            </Card>
        );
    }
}

class SwConfigTable extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            loading: false,
            currentID: 0,
            lastKnownID: 0,
            name: "",
            type: 0,
            dataSource: [
                {
                    key: 0,
                    ItemName: "Part Number",
                    Current: "",
                    LastKnown: ""
                }, {
                    key: 1,
                    ItemName: "Version",
                    Current: "",
                    LastKnown: ""
                }, {
                    key: 2,
                    ItemName: "CRC",
                    Current: "",
                    LastKnown: ""
                }
            ]
        };

        this.handleTableChange = this
            .handleTableChange
            .bind(this);
    }

    getSwConfig(block) {
        var url = "/getswconfig";
        var result = null;

        $.ajax({
            url: url,
            dataType: "json",
            type: "POST",
            cache: false,
            async: false,
            data: {
                "block": block,
                "type": this.props.type
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
        var currentData = this.getSwConfig(0);
        console.log("sw current: ", currentData);

        this.setState({type: this.props.type});

        if (currentData != null) {
            const dataSource = this.state.dataSource;
            this.setState({currentID: currentData.ID});
            dataSource[0].Current = currentData.PartNumber;
            dataSource[1].Current = currentData.Version;
            dataSource[2].Current = currentData.ImageCRC;
            this.setState({dataSource});
        }
        var lastKnownData = this.getSwConfig(2);
        console.log("sw last known: ", lastKnownData);
        if (lastKnownData != null) {
            const dataSource = this.state.dataSource;
            this.setState({lastKnownID: lastKnownData.ID});
            dataSource[0].LastKnown = lastKnownData.PartNumber;
            dataSource[1].LastKnown = lastKnownData.Version;
            dataSource[2].LastKnown = lastKnownData.ImageCRC;
            this.setState({dataSource});
        }
    }
    render() {
        const pagination = false;

        const columns = [
            {
                title: "#",
                key: "ItemName",
                dataIndex: 'ItemName',
                render: text => <b>{text}</b>
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
                pagination={pagination}
                dataSource={this.state.dataSource}
                columns={columns}
                loading={this.state.loading}
                bordered></Table >

        );
    }
}

class SoftwareComponent extends React.Component {
    render() {
        return (
            <Card title={this.props.name}>
                <SwConfigTable type={this.props.type}/>
            </Card>
        );
    }
}
