class SystemConfiguration extends React.Component {
    constructor(props) {
        super(props);

        this.handleDeviceNameChange = this.handleDeviceNameChange.bind(this);
        this.handleSystemVersionChange = this.handleSystemVersionChange.bind(this);
        this.handleDeviceSKUChange = this.handleDeviceSKUChange.bind(this);
        this.handleSerialNumberChange = this.handleSerialNumberChange.bind(this);
        this.handleSoftwareBuildChange = this.handleSoftwareBuildChange.bind(this);
        this.handlePartNumberChange = this.handlePartNumberChange.bind(this);
        this.handleHardwareVersionChange = this.handleHardwareVersionChange.bind(this);


        this.handleUpdateButtonClick = this.handleUpdateButtonClick.bind(this);

        this.state = {
            current: {
                id: "",
                deviceName: "",
                systemVersion: "",
                deviceSKU: "",
                serialNumber: "",
                softwareBuild: "",
                partNumber: "",
                hardwareVersion: "",
            },
            lastKnown: {
                id: "",
                deviceName: "",
                systemVersion: "",
                deviceSKU: "",
                serialNumber: "",
                softwareBuild: "",
                partNumber: "",
                hardwareVersion: "",
            },
        };
    }
    handleDeviceNameChange(e) {
        var c = this.state.current;
        c.deviceName = e.target.value;
        this.setState({
            current: c,
        });
    }
    handleSystemVersionChange(e) {
        var c = this.state.current;
        c.systemVersion = e.target.value;
        this.setState({
            current: c,
        });
    }

    handleDeviceSKUChange(e) {
        var c = this.state.current;
        c.deviceSKU = e.target.value;
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

    handleSoftwareBuildChange(e) {
        var c = this.state.current;
        c.softwareBuild = e.target.value;
        this.setState({
            current: c,
        });
    }

    handlePartNumberChange(e) {
        var c = this.state.current;
        c.partNumber = e.target.value;
        this.setState({
            current: c,
        });

    }

    handleHardwareVersionChange(e) {
        var c = this.state.current;
        c.hardwareVersion = e.target.value;
        this.setState({
            current: c,
        });
    }


    handleUpdateButtonClick(e) {
        var url = "/setsysconfig";
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

    getSysConfig(block) {
        var url = "/getsysconfig";
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
        var currentData = this.getSysConfig(0);
        var lastKnownData = this.getSysConfig(2);
        console.log("system current: ", currentData);
        console.log("system last known: ", lastKnownData);
        this.setState({
            current: {
                id: currentData.ID,
                deviceName: currentData.DeviceName,
                systemVersion: currentData.SystemVersion,
                deviceSKU: currentData.DeviceSKU,
                serialNumber: currentData.SerialNumber,
                softwareBuild: currentData.SoftwareBuild,
                partNumber: currentData.PartNumber,
                hardwareVersion: currentData.HardwareVersion,
            },
            lastKnown: {
                id: lastKnownData.ID,
                deviceName: lastKnownData.DeviceName,
                systemVersion: lastKnownData.SystemVersion,
                deviceSKU: lastKnownData.DeviceSKU,
                serialNumber: lastKnownData.SerialNumber,
                softwareBuild: lastKnownData.SoftwareBuild,
                partNumber: lastKnownData.PartNumber,
                hardwareVersion: lastKnownData.HardwareVersion,
            },
        });
    }


    render() {
        return (
            <div className="panel panel-default">
                <div className="panel-heading">System Configuration</div>
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
                                <th scope="row">Device Name</th>
                                <td><input type="text" value={this.state.current.deviceName} className="form-control" onChange={this.handleDeviceNameChange}></input></td>
                                <td>{this.state.lastKnown.deviceName}</td>
                            </tr>
                            <tr>
                                <th scope="row">Serial Number</th>
                                <td><input type="text" value={this.state.current.serialNumber} className="form-control" onChange={this.handleSerialNumberChange}></input></td>
                                <td>{this.state.lastKnown.serialNumber}</td>
                            </tr>
                            <tr>
                                <th scope="row">System Version</th>
                                <td><input type="text" value={this.state.current.systemVersion} className="form-control" onChange={this.handleSystemVersionChange}></input></td>
                                <td>{this.state.lastKnown.systemVersion}</td>
                            </tr>
                            <tr>
                                <th scope="row">Device SKU</th>
                                <td><input type="text" value={this.state.current.deviceSKU} className="form-control" onChange={this.handleDeviceSKUChange}></input></td>
                                <td>{this.state.lastKnown.deviceSKU}</td>
                            </tr>
                            <tr>
                                <th scope="row">Software Build</th>
                                <td><input type="text" value={this.state.current.softwareBuild} className="form-control" onChange={this.handleSoftwareBuildChange}></input></td>
                                <td>{this.state.lastKnown.softwareBuild}</td>
                            </tr>
                            <tr>
                                <th scope="row">Part Number</th>
                                <td><input type="text" value={this.state.current.partNumber} className="form-control" onChange={this.handlePartNumberChange}></input></td>
                                <td>{this.state.lastKnown.partNumber}</td>
                            </tr>
                            <tr>
                                <th scope="row">Hardware Version</th>
                                <td><input type="text" value={this.state.current.hardwareVersion} className="form-control" onChange={this.handleHardwareVersionChange}></input></td>
                                <td>{this.state.lastKnown.hardwareVersion}</td>
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
}

class HardwareConfiguration extends React.Component {
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
}

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

class DeviceConfiguration extends React.Component {

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

ReactDOM.render(<DeviceConfiguration />, document.getElementById("configuration"));