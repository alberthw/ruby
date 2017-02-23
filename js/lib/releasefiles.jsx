import React from "react";
import {Table, Button, Icon, DatePicker} from "antd";
import $ from "jquery";

class DownloadFile extends React.Component {
    constructor(props) {
        super(props);
        this.handleDownloadButtonClick = this
            .handleDownloadButtonClick
            .bind(this);
    }

    handleDownloadButtonClick(e) {
        var result = this.downloadReleaseFile(this.props.file);
        //      alert(result);
    }

    downloadReleaseFile(fileinfo) {
        console.log("file id :", fileinfo);

        var url = "/downloadfile";
        var result = null;

        $.ajax({
            url: url,
            dataType: "json",
            cache: false,
            async: false,
            type: "POST",
            data: {
                id: fileinfo.ID,
                filepath: fileinfo.RemotePath
            },
            success: function (data) {
                result = data;
            },
            error: function (xhr, status, err) {
                console.error(url, status, err.toString());
            }
        });

        return result;
    }

    render() {
        return (
            <Button onClick={this.handleDownloadButtonClick}>Download</Button>
        );
    }
}

class BurnImage extends React.Component {
    constructor(props) {
        super(props);
        this.handleBurnImageButtonClick = this
            .handleBurnImageButtonClick
            .bind(this);
    }

    burnImage(file) {
        var url = "/burnhostimage";
        $.ajax({
            url: url,
            dataType: "json",
            cache: false,
            async: false,
            type: "POST",
            data: {
                filetype: file.FileType,
                filepath: file.LocalPath
            },
            success: function (data) {
                //          alert(data);
            },
            error: function (xhr, status, err) {
                console.error(url, status, err.toString());
            }
        });
    }

    handleBurnImageButtonClick(e) {
        var f = this.props.file;
        if (f.IsDownloaded != true) {
            alert("download the hex image first.")
            return;
        }

        this.burnImage(f);
    }

    render() {
        return (
            <Button onClick={this.handleBurnImageButtonClick}>Burn</Button>
        );
    }
}

export default class ReleaseFilesTable extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            data: [],
            searchDate: '',
            pagination: {
                showSizeChanger: true,
                pageSize: 10
            },
            filters: {},
            sorter: {},
            loading: false
        };
        this.handleTableChange = this
            .handleTableChange
            .bind(this);
        this.handleDateFilterChange = this
            .handleDateFilterChange
            .bind(this);
        this.handleSearchButtonClick = this
            .handleSearchButtonClick
            .bind(this);
    }

    handleSearchButtonClick(e) {
        this.getReleaseFile({
            searchDate: this.state.searchDate,
            pageSize: this.state.pagination.pageSize,
            page: this.state.pagination.current,
            sortField: this.state.sorter.field,
            sortOrder: this.state.sorter.order,
            filters: this.state.filters
        });
    }

    handleTableChange(pagination, filters, sorter) {
        //       console.log("pagination : ", pagination);        console.log("filters :
        // ", filters);       console.log("sorter : ", sorter);

        const pager = this.state.pagination;
        pager.current = pagination.current;
        pager.pageSize = pagination.pageSize;
        this.setState({pagination: pager, filters: filters, sorter: sorter});
        this.getReleaseFile({
            searchDate: this.state.searchDate,
            pageSize: this.state.pageSize,
            page: this.state.current,
            sortField: sorter.field,
            sortOrder: sorter.order,
            ...filters
        });
    }

    getReleaseFile(params) {
        console.log("params : ", params);
        this.setState({loading: true});
        var url = "/getfilerepo";
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
                this.setState({loading: false});
                console.error(url, status, err.toString());
            }.bind(this)
        });
    }

    componentDidMount() {
        this.getReleaseFile({pageSize: 10});
    }

    handleDateFilterChange(dateString, date) {
        //      console.log("select date : '", date, "'");
        this.setState({searchDate: date});
    }

    render() {
        const columns = [
            {
                title: 'File Type',
                key: "FileType",
                dataIndex: 'FileType',
                render: (text) => {
                    var result = "";
                    switch (text) {
                        case 1:
                            result = "Host Application";
                            break;
                        case 2:
                            result = "Host Boot Loader";
                            break;
                        case 3:
                            result = "DSP Application";
                            break;
                    }
                    return result;

                },
                filters: [
                    {
                        text: "Host Application",
                        value: 1
                    }, {
                        text: "Host Boot Loader",
                        value: 2
                    }, {
                        text: "DSP Application",
                        value: 3
                    }
                ]
            }, {
                title: 'File Name',
                key: "FileName",
                dataIndex: 'FileName'
            }, {
                title: 'Build Number',
                key: "BuildNumber",
                dataIndex: 'BuildNumber'
            }, {
                title: 'File Size(KB)',
                key: "FileSize",
                dataIndex: 'FileSize'
            }, {
                title: 'CRC',
                key: "CRC",
                dataIndex: 'CRC'
            }, {
                title: 'Download Status',
                key: "IsDownloaded",
                dataIndex: 'IsDownloaded',
                sorter: true,
                render: (text) => text
                    ? "true"
                    : "false"
            }, {
                title: 'Download',
                key: "download",
                dataIndex: 'download',
                render: (text, record) => (<DownloadFile file={record}/>)

            }, {
                title: 'Upgrade',
                key: "upgrade",
                dataIndex: 'upgrade',
                render: (text, record) => (<BurnImage file={record}/>)
            }
        ];
        return (
            <div>
                Select Date :
                <DatePicker format="MM/DD/YYYY" onChange={this.handleDateFilterChange}/>
                <Button icon="search" onClick={this.handleSearchButtonClick}>Search</Button>
                <Table
                    dataSource={this.state.data}
                    columns={columns}
                    pagination={this.state.pagination}
                    loading={this.state.loading}
                    onChange={this.handleTableChange}
                    bordered></Table>
            </div>
        );
    }
}
