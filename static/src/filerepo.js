class DownloadFile extends React.Component {
    constructor(props) {
        super(props);
        this.handleDownloadButtonClick = this.handleDownloadButtonClick.bind(this);
    }

    handleDownloadButtonClick(e) {
        var result = this.downloadReleaseFile(this.props.file);
        //      alert(result);
    }

    downloadReleaseFile(fileinfo) {
        console.log("file id :", fileinfo.ID, ", download file : ", fileinfo.RemotePath);
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
            <input type="button" value="Download" onClick={this.handleDownloadButtonClick}></input>
        );
    }
}

class BurnImage extends React.Component {
    constructor(props) {
        super(props);
        this.handleBurnImageButtonClick = this.handleBurnImageButtonClick.bind(this);
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
            <input type="button" value="Burn" onClick={this.handleBurnImageButtonClick}></input>
        );
    }
}

class FileTable extends React.Component {
    constructor(props) {
        super(props);
    }
    render() {
        const rows = this.props.files;

        const {Table, Column, Cell} = FixedDataTable;

        const TextCell = ({rowIndex, col, ...props}) => (
            <Cell {...props}>
                {rows[rowIndex][col].toString()}
            </Cell>
        );

        return (
            <Table
                rowsCount={rows.length}
                rowHeight={35}
                headerHeight={50}
                width={1000}
                height={600}
                {...this.props}>
                <Column
                    header={<Cell>File Name</Cell>}
                    cell={<TextCell col="FileName" />}

                    width={300}
                    />
                <Column
                    header={<Cell>CRC</Cell>}
                    cell={<TextCell col="CRC" />}
                    width={100}
                    />
                <Column
                    header={<Cell>File Size(KB)</Cell>}
                    cell={<TextCell col="FileSize" />}
                    width={100}
                    />
                <Column
                    header={<Cell>Build Number</Cell>}
                    cell={<TextCell col="BuildNumber" />}
                    width={150}
                    />
                <Column
                    header={<Cell>Download Status</Cell>}
                    cell={<TextCell col="IsDownloaded" />}
                    width={150}
                    />
                <Column
                    header={<Cell></Cell>}
                    cell={props => (
                        <Cell {...props}>
                            <DownloadFile file={rows[props.rowIndex]} />
                        </Cell>
                    )}
                    width={100}
                    />
                <Column
                    header={<Cell></Cell>}
                    cell={props => (
                        <Cell {...props}>
                            <BurnImage file={rows[props.rowIndex]} />
                        </Cell>
                    )}
                    width={100}
                    />
            </Table>
        );
    }
}

class TableFilter extends React.Component {
    constructor(props) {
        super(props);

        this.handleSearchDateClick = this.handleSearchDateClick.bind(this);
        this.handleDateChange = this.handleDateChange.bind(this);
    }

    handleSearchDateClick(e) {
        this.props.onSearch($("#datepicker").val());
    }

    handleDateChange(e) {
        this.props.onChange($("#datepicker").val());
    }

    componentDidMount() {
        var props = this.props;
        $("#datepicker").datepicker({
            onSelect: function (text) {
                props.onChange(text);
            }
        });
    }

    render() {
        return (
            <div className="input-group col-md-4">
                <span className="input-group-addon">Filter:</span>
                <input type="text" className="form-control" id="datepicker" placeholder="MM/DD/YYYY" onChange={this.handleDateChange}></input>
                <span className="input-group-btn">
                    <button type="button" className="btn btn-default" onClick={this.handleSearchDateClick}>Search</button>
                </span>
            </div>
        );
    }
}

class FileUpload extends React.Component{
    
}


class FileRepo extends React.Component {
    constructor(props) {
        super(props);

        this.handleFilterChange = this.handleFilterChange.bind(this);
        this.handleFilterSearch = this.handleFilterSearch.bind(this);

        this.state = {
            filter: {
                date: ""
            },
            data: this.getReleaseFiles(null),

        };
    }

    componentDidMount() {
        this.timerID = setInterval(
            () => this.syncData(this.state.filter),
            5000
        );
    }

    componentWillUnmount() {
        clearInterval(this.timerID);
    }

    syncData(e) {
        var data = this.getReleaseFiles(this.state.filter);
        this.setState({
            data: data
        });
        //       console.log(data);
    }

    handleFilterSearch(e) {
        this.syncData();
    }

    handleFilterChange(e) {
        this.setState({
            filter: {
                date: e,
            }
        });
        this.syncData();
    }

    getReleaseFiles(filter) {
        console.log("filter", filter);
        var result = [];
        var url = "/getfilerepo";
        var dt = "";
        if (filter != null) {
            dt = filter.date;
        }
        $.ajax({
            url: url,
            dataType: "json",
            cache: false,
            async: false,
            data: {
                "date": dt,
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
            <div>
                <TableFilter onChange={this.handleFilterChange} onSearch={this.handleFilterSearch}></TableFilter>
                <FileTable files={this.state.data}></FileTable>
            </div>
        );
    }
}

ReactDOM.render(
    <FileRepo />,
    document.getElementById("filerepo"));