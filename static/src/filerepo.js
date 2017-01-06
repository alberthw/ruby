function getReleaseFiles() {
    var result = [];
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
        var url ="/burnhostimage"; 
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
        if (f.Isdownloaded != true){
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
        //        var Column = FixedDataTable.Column;
        //        var Table = FixedDataTable.Table;
        //        var Cell = FixedDataTable.Cell;
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
                    width={150}
                    />
                <Column
                    header={<Cell>Download Status</Cell>}
                    cell={props => (
                        <Cell {...props}>
                            {rows[props.rowIndex].Isdownloaded.toString()}

                        </Cell>
                    )}
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
            5000
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
                <div hidden>
                    <a>Release Files:</a>
                    <input type="button" value="Sync" onClick={this.handleSyncButtonClick} />
                </div>

                <FileTable files={this.state.data}></FileTable>
            </div>
        );
    }
}

ReactDOM.render(
    <FileRepo />,
    document.getElementById("filerepo"));