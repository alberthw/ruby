



class SoftwareConfiguration extends React.Component {
    constructor(props) {
        super(props);
    }

    render() {
        return (
            <div className="panel panel-default">
                <div className="panel-heading">Software Configuration</div>
                <div className="panel-body">
                    <div className="rows container-fluid">
                        <div className="col-xs-4">
                            <SoftwareComponent name="Host Boot Loader" type="0" />
                        </div>
                        <div className="col-xs-4">
                            <SoftwareComponent name="Host Application" type="1" />
                        </div>
                        <div className="col-xs-4">
                            <SoftwareComponent name="DSP Application" type="2" />
                        </div>
                    </div>
                </div>
            </div>
        );
    }
}

class SoftwareComponent extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            current: {
                id: "",
                name: "",
                type: "",
                partNumber: "",
                version: "",
                imageCRC: "",
            },
            lastKnown: {
                id: "",
                name: "",
                type: "",
                partNumber: "",
                version: "",
                imageCRC: "",
            },
        };

    }

    getSoftwareConfig(block) {
        var url = "/getswconfig";
        var result = null;

        $.ajax({
            url: url,
            dataType: "json",
            type: "POST",
            cache: false,
            async: false,
            data: {
                "type": this.props.type,
                "block": block,
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
        var currentdata = this.getSoftwareConfig(0);
        var lastKnowndata = this.getSoftwareConfig(2);
        console.log(this.props.name, " current:", currentdata);
        console.log(this.props.name, " last known:", lastKnowndata);
        this.setState({
            current: {
                id: currentdata.ID,
                name: currentdata.Name,
                type: this.props.type,
                partNumber: currentdata.PartNumber,
                version: currentdata.Version,
                imageCRC: currentdata.ImageCRC,
            },
            lastKnown: {
                id: lastKnowndata.ID,
                name: lastKnowndata.Name,
                type: this.props.type,
                partNumber: lastKnowndata.PartNumber,
                version: lastKnowndata.Version,
                imageCRC: lastKnowndata.ImageCRC,
            },

        });
    }

    render() {
        return (
            <div className="panel panel-default">
                <div className="panel-heading">{this.props.name}</div>
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
                                <td>{this.state.current.partNumber}</td>
                                <td>{this.state.lastKnown.partNumber}</td>
                            </tr>
                            <tr>
                                <th scope="row">Version</th>
                                <td>{this.state.current.version}</td>
                                <td>{this.state.lastKnown.version}</td>
                            </tr>
                            <tr>
                                <th scope="row">CRC</th>
                                <td>{this.state.current.imageCRC}</td>
                                <td>{this.state.lastKnown.imageCRC}</td>
                            </tr>
                        </tbody>
                    </table>

                </div>

            </div >
        );
    }
}

export default class DeviceConfiguration extends React.Component {

    constructor(props) {
        super(props);

        this.handleValidateButtonClick = this.handleValidateButtonClick.bind(this);
        this.handleUpdateButtonClick = this.handleUpdateButtonClick.bind(this);

        this.state = {
            isConfigValidated: false,
        };

    }

    handleValidateButtonClick(e) {
        var url = "/validateconfig";
        $.ajax({
            url: url,
            dataType: "json",
            type: "GET",
            cache: false,
            async: false,
            success: function (data) {
                console.log(data);
                //result = data;
            },
            error: function (xhr, status, err) {
                console.error(url, status, err.toString());
            }
        });
    }

    handleUpdateButtonClick(e){
        this.getVersions();
        this.getLastKnownVersions();
        window.location.reload();
    }

    getConfigValidateStatus() {
        var url = "/config";
        $.ajax({
            url: url,
            dataType: "json",
            cache: false,
            async: false,
            success: function (data) {
                console.log("config:", data);
                this.setState({
                    isConfigValidated: data.IsConfigValidated,
                })
            }.bind(this),
            error: function (xhr, status, err) {
                console.error(url, status, err.toString());
            }.bind(this)
        });

    }

    getVersions() {
        var url = "/getversion";
        var result = null;

        $.ajax({
            url: url,
            dataType: "json",
            type: "GET",
            cache: false,
            async: false,
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

    getLastKnownVersions() {
        var url = "/getlastknownversion";
        var result = null;

        $.ajax({
            url: url,
            dataType: "json",
            type: "GET",
            cache: false,
            async: false,
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

    componentWillMount() {
        this.getConfigValidateStatus();
    }


    render() {
        return (
            <div className="panel panel-default panel-primary">
                <div className="panel-heading">Device Configuration</div>
                <div className="panel-body">
                    <div className="container-fluid">
                        <div className="panel panel-default">
                            <div className="panel-body">
                                <p className="navbar-text">Configuration Validation Status:</p>
                                <p className="navbar-text">{this.state.isConfigValidated.toString()}</p>
                                <input type="button" className="btn btn-default" value="Validate" onClick={this.handleValidateButtonClick}></input>
                                <input type="button" className="btn btn-default" value="Update" onClick={this.handleUpdateButtonClick}></input>
                            </div>
                        </div>
                    </div>
                    <div className="clearfix"></div>
                    <div className="rows container-fluid">
                        <div className="col-xs-6">
                            <SystemConfiguration />
                        </div>
                        <div className="col-xs-6">
                            <HardwareConfiguration />
                        </div>
                    </div>
                    <div>
                        <SoftwareConfiguration />
                    </div>
                </div>
            </div >
        );
    }

}

