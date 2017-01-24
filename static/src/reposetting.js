function TestRemoteServerConnection(url) {
    var result = false;
    $.ajax({
        url: url,
        dataType: "json",
        cache: false,
        async: false,
        success: function (data) {
            console.log(data);
            if (data == "200") {
                result = true;
            }
        },
        error: function (xhr, status, err) {
            console.error(url, status, err.toString());
        }
    });
    return result;
}

function GetRepositorySetting() {
    var result = null;
    $.ajax({
        url: "/reposetting",
        dataType: "json",
        //       cache: false,
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

function UpdateRepositorySetting(id, ip, status) {
    var result = false;
    $.ajax({
        url: "/reposetting",
        dataType: "json",
        type: "POST",
        async: false,
        data: {
            Id: id,
            Remoteserver: ip,
            Isconnected: status
        },
        success: function (data) {
            //          console.log(data);
            result = true;
        },
        error: function (xhr, status, err) {
            console.error(url, status, err.toString());
        }
    });

    return result;
}

class RepositorySetting extends React.Component {

    constructor(props) {
        super(props);

        var result = GetRepositorySetting();

        this.state = {
            Id: result.Id,
            Remoteserver: result.Remoteserver,
            Isconnected: result.Isconnected
        };

        this.handleRemoteServerIpChange = this
            .handleRemoteServerIpChange
            .bind(this);
        this.handleIsConnectedChange = this
            .handleIsConnectedChange
            .bind(this);
        this.handleCheckClick = this
            .handleCheckClick
            .bind(this);

    }

    handleRemoteServerIpChange(e) {
        this.setState({ Remoteserver: e.target.value });
    }

    handleIsConnectedChange(e) {
        this.setState({ Isconnected: e.target.value });
    }

    handleCheckClick(e) {
        UpdateRepositorySetting(this.state.Id, this.state.Remoteserver, isConnected);
        var url = "/testremoteserver";
        var isConnected = TestRemoteServerConnection(url);
        console.log("server connection :", isConnected);
        this.setState({
            Isconnected: isConnected
        });
        UpdateRepositorySetting(this.state.Id, this.state.Remoteserver, isConnected);
    }

    render() {
        const remoteServer = this.state.Remoteserver;
        const isConnected = this.state.Isconnected;
        return (
            <div className="container-fluid">
                <div className="input-group">
                    <span className="input-group-addon">Remote File Server:</span>
                    <input type="text" className="form-control" value={remoteServer} onChange={this.handleRemoteServerIpChange}></input>
                    <span className="input-group-addon">Connected:</span>
                    <input type="text" className="form-control" value={isConnected} onChange={this.handleIsConnectedChange} readOnly></input>
                    <span className="input-group-btn">
                        <button type="button" className="btn btn-default" onClick={this.handleCheckClick}>Check</button>
                    </span>
                </div>
            </div>

        );
    }

}

ReactDOM.render(<RepositorySetting />, document.getElementById("reposetting"));