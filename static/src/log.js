class DeviceLog extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            data: [],
            pagination: {
                showSizeChanger: true,
            },
            loading: false,
        };

        this.handleTableChange = this.handleTableChange.bind(this);
    }

    handleTableChange(pagination, filters, sorter) {
        const pager = this.state.pagination;
        pager.current = pagination.current;
        this.setState({
            pagination: pager,
        });
        this.getDeviceLog({
            pageSize: pagination.pageSize,
            page: pagination.current,
            sortField: sorter.field,
            sortOrder: sorter.order,
            ...filters,
        });
    }

    getDeviceLog(params) {
        console.log("params : ", params);
        this.setState({
            loading: true,
        });
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
                for (let i = 0; i < result.Data.length; i++) {
                    result.Data[i].key = result.Data[i].ID;
                    switch (result.Data[i].LogType) {
                        case 0:
                            result.Data[i].LogType = "error";
                            break;
                        case 1:
                            result.Data[i].LogType = "event";
                            break;
                    }
                }
                this.setState({
                    loading: false,
                    data: result.Data,
                    pagination,
                });
            }.bind(this),
            error: function (xhr, status, err) {
                console.error(url, status, err.toString());
            }.bind(this)
        });
    }

    componentDidMount() {
        this.getDeviceLog({
            pageSize: 10,
        });
    }

    render() {
        const {Button, Table} = antd;
        const columns = [{
            title: 'log Type',
            dataIndex: 'LogType',
            filters: [
                { text: "Error", value: 0 },
                { text: "Event", value: 1 },
            ],

        }, {
            title: 'Content',
            dataIndex: 'Content',
        }, {
            title: 'Time',
            dataIndex: 'Created',
            sorter: true,
        }];

        return (
            <div>
                <Table dataSource={this.state.data} columns={columns} pagination={this.state.pagination} loading={this.state.loading} onChange={this.handleTableChange} />
            </div>

        );
    }
}


ReactDOM.render(<DeviceLog />, document.getElementById("log"));