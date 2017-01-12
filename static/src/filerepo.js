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
        console.log("file id :", fileinfo.Id, ", download file : ", fileinfo.Remotepath);
        var url = "/downloadfile";
        var result = null;
        $.ajax({
            url: url,
            dataType: "json",
            cache: false,
            async: false,
            type: "POST",
            data: {
                id: fileinfo.Id,
                filepath: fileinfo.Remotepath
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
                filetype: file.Filetype,
                filepath: file.Filepath
            },
            success: function (data) {
                alert(data);
            },
            error: function (xhr, status, err) {
                console.error(url, status, err.toString());
            }
        });
    }

    handleBurnImageButtonClick(e) {
        var f = this.props.file;
        if (f.Isdownloaded != true) {
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
                    cell={<TextCell col="Filename" />}

                    width={300}
                    />
                <Column
                    header={<Cell>CRC</Cell>}
                    cell={<TextCell col="Crc" />}
                    width={100}
                    />
                <Column
                    header={<Cell>File Size(KB)</Cell>}
                    cell={<TextCell col="Filesize" />}
                    width={100}
                    />
                <Column
                    header={<Cell>Build Number</Cell>}
                    cell={<TextCell col="Buildnumber" />}
                    width={150}
                    />
                <Column
                    header={<Cell>Download Status</Cell>}
                    cell={<TextCell col="Isdownloaded" />}
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
        console.log("date 1 :" + e.target.value);
        console.log("date 2 : " + $("#datepicker").val());
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
            <div>
                <a>Filter:</a>
                <input type="text" id="datepicker" placeholder="MM/DD/YYYY" onChange={this.handleDateChange}></input>
                <input type="button" value="Search" onClick={this.handleSearchDateClick} ></input>
            </div>
        );
    }
}


class FileRepo extends React.Component {
    constructor(props) {
        super(props);

        this.handleSyncButtonClick = this.handleSyncButtonClick.bind(this);
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
            () => this.handleSyncButtonClick(this.state.filter),
            5000
        );
    }

    componentWillUnmount() {
        clearInterval(this.timerID);
    }

    handleSyncButtonClick(e) {
        var data = this.getReleaseFiles(this.state.filter);
        this.setState({
            data: data
        });
        //       console.log(data);
    }

    handleFilterSearch(e) {
        this.handleSyncButtonClick();
    }

    handleFilterChange(e) {
        this.setState({
            filter: {
                date: e,
            }
        });
        this.handleSyncButtonClick();
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
                <div hidden>
                    <a>Release Files:</a>
                    <input type="button" value="Sync" onClick={this.handleSyncButtonClick} />
                </div>
                <TableFilter onChange={this.handleFilterChange} onSearch={this.handleFilterSearch}></TableFilter>

                <FileTable files={this.state.data}></FileTable>
            </div>
        );
    }
}

ReactDOM.render(
    <FileRepo />,
    document.getElementById("filerepo"));