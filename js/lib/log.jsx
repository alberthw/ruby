import React from "react";
import {Table} from "antd";
import $ from "jquery";

export default class DeviceLog extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            data: [],
            pagination: {
                showSizeChanger: true
            },
            loading: false
        };

        this.handleTableChange = this
            .handleTableChange
            .bind(this);
    }

    handleTableChange(pagination, filters, sorter) {
        const pager = this.state.pagination;
        pager.current = pagination.current;
        this.setState({pagination: pager});
        this.getDeviceLog({pageSize: pagination.pageSize, page: pagination.current, sortField: sorter.field, sortOrder: sorter.order, ...filters});
    }

    getDeviceLog(params) {
        console.log("params : ", params);
        this.setState({loading: true});
        var url = "/getdevicelog";
        $.ajax({
            url: url,
            dataType: "json",
            cache: false,
            async: false,
            type: "POST",
            data: params,
            success: function (result) {
                console.log("result :", result);
                const pagination = this.state.pagination;
                pagination.total = result.Total;
                this.setState({loading: false, data: result.Data, pagination});
            }.bind(this),
            error: function (xhr, status, err) {
                console.error(url, status, err.toString());
            }.bind(this)
        });
    }

    componentDidMount() {
        this.getDeviceLog({pageSize: 10});
    }

    render() {
        const columns = [
            {
                title: 'log Type',
                key: "LogType",
                dataIndex: 'LogType',
                render: (text) => text
                    ? "Event"
                    : "Error",
                filters: [
                    {
                        text: "Error",
                        value: 0
                    }, {
                        text: "Event",
                        value: 1
                    }
                ]
            }, {
                title: 'Content',
                dataIndex: 'Content',
                key: "Content"
            }, {
                title: 'Time',
                dataIndex: 'Created',
                key: "Created",
                sorter: true
            }
        ];

        return (
            <div>
                <Table
                    dataSource={this.state.data}
                    columns={columns}
                    pagination={this.state.pagination}
                    loading={this.state.loading}
                    onChange={this.handleTableChange}/>
            </div>
        );
    }
}
