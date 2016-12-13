function getReleaseFiles() {
    var result = null;
    var url = "/getfilerepo";
    $.ajax({
        url: url,
        dataType: "json",
        cache: false,
        async: false,
        success: function (data) {
            
            result = data;
        },
        error: function (xhr, status, err) {
            console.error(url, status, err.toString());
        }
    });
    return result;
}

function downloadReleaseFile(fileinfo) {
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
            id:fileinfo.Id,
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

class DownloadFile extends React.Component {
    constructor(props) {
        super(props);
        this.handleDownloadButtonClick = this.handleDownloadButtonClick.bind(this);
    }

    handleDownloadButtonClick(e) {
        var result = downloadReleaseFile(this.props.file);
        //      alert(result);
    }

    render() {
        return (
            <input type="button" value="Download" onClick={this.handleDownloadButtonClick}></input>
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
        //        var Column = FixedDataTable.Column;
        //        var Table = FixedDataTable.Table;
        //        var Cell = FixedDataTable.Cell;
        return (
            <Table
                rowsCount={rows.length}
                rowHeight={50}
                headerHeight={50}
                width={1000}
                height={1000}>
                <Column
                    header={<Cell>File Name</Cell>}
                    cell={props => (
                        <Cell {...props}>
                            {rows[props.rowIndex].Filename}
                        </Cell>
                    )}
                    width={300}
                    />
                <Column
                    header={<Cell>CRC</Cell>}
                    cell={props => (
                        <Cell {...props}>
                            {rows[props.rowIndex].Crc}
                        </Cell>
                    )}
                    width={100}
                    />
                    <Column
                    header={<Cell>File Size(KB)</Cell>}
                    cell={props => (
                        <Cell {...props}>
                            {rows[props.rowIndex].Filesize}
                        </Cell>
                    )}
                    width={100}
                    />
                <Column
                    header={<Cell>Build Number</Cell>}
                    cell={props => (
                        <Cell {...props}>
                            {rows[props.rowIndex].Buildnumber}
                        </Cell>
                    )}
                    width={200}
                    />
                <Column
                    header={<Cell>Download Status</Cell>}
                    cell={props => (
                        <Cell {...props}>
                            {rows[props.rowIndex].Isdownloaded.toString()}

                        </Cell>
                    )}
                    width={200}
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
            </Table>
        );
    }
}


class FileRepo extends React.Component {
    constructor(props) {
        super(props);

        this.handleSyncButtonClick = this.handleSyncButtonClick.bind(this);

        this.state = {
            data: getReleaseFiles()
        };
    }

    componentDidMount() {
        this.timerID = setInterval(
            () => this.handleSyncButtonClick(),
            2000
        );
    }

    componentWillUnmount() {
        clearInterval(this.timerID);
    }

    handleSyncButtonClick(e) {
        var data = getReleaseFiles();
        this.setState({
            data: data
        });
 //       console.log(data);
    }

    render() {
        return (
            <div>
                <a>Release Files:</a>
                <input type="button" value="Sync" onClick={this.handleSyncButtonClick} />
                <FileTable files={this.state.data}></FileTable>
            </div>
        );
    }
}

ReactDOM.render(
    <FileRepo />,
    document.getElementById("filerepo"));