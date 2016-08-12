
var configItem = React.createClass({
    render:function(){
        return
    }
});
var RubyConfig = React.createClass({
    getInitialState: function () {
        return {
            SerialLine: "",
            SerialSpeed: null,
            updated: Date.now()
        };
    },
    handleSerialLineChange: function (e) {
        this.setState({
            SerialLine: e.target.value
        });
    },
    handleSerialSpeedChange: function (e) {
        this.setState({
            SerialSpeed: e.target.value
        });
    },
    render: function () {
        return (
            <form>
                Serial Line: <input type="text" id="serialline" placeholder="Serial Line" value={this.state.SerialLine} onChange={this.handleSerialLineChange}></input>
                Serial Speed: <input type="text" id="serialspeed" placeholder="Serial Speed" value={this.state.SerialSpeed} onChange={this.handleSerialSpeedChange}></input>
                <input type="submit" id="submitconfig" value="Submit"></input>
            </form>
        );
    }

});

ReactDOM.render(<RubyConfig />, document.getElementById("rubyconfig"));



/*
var Comment = React.createClass({
    render: function () {
        var md = new Remarkable();
        return (
            <div className="comment">
                <h2 className="commentAuthor">
                    {this.props.author}
                </h2>
                <span dangerouslySetInnerHTML={this.rawMarkup() } ></span>
            </div>
        );
    },
    rawMarkup: function () {
        var md = new Remarkable();
        var rawMarkup = md.render(this.props.children.toString());
        return { __html: rawMarkup };
    }
});

var data = [
    { id: 1, author: "Pete Hunt", text: "This is one comment." },
    { id: 2, author: "Jordan Walke", text: "This is *another* comment." }
];
var url = "/config";

var CommentList = React.createClass({
    render: function () {
        var commentNodes = this.props.data.map(function (comment) {
            return (
                <Comment author={comment.Serialline} key={comment.Id}>
                    {comment.Serialspeed}
                </Comment>
            );
        });
        return (
            <div className="commentList">
                {commentNodes}
            </div>

        );
    }
});

var CommentForm = React.createClass({
    render: function () {
        return (
            <div className="commentForm">
                Hello, world~I am a CommentForm.
            </div>
        );
    }

});

var CommentBox = React.createClass({
    getInitialState: function () {
        return { data: [] };
    },
    loadCommentsFromServer: function () {
        $.ajax({
            url: this.props.url,
            dataType: 'json',
            cache: false,
            success: function (data) {
                this.setState({ data: data });
            }.bind(this),
            error: function (xhr, status, err) {
                console.error(this.props.url, status, err.toString());
            }.bind(this)

        });
    },
    handleCommentSubmit: function (comment) {
        var comments = this.state.data;
        comments.index = Date.now();
        var newComments = comments.concat([comment]);
        this.setState({data:newComments});
        $.ajax({
            url: this.props.url,
            dataType: 'json',
            type: 'POST',
            data: comment,
            success: function (data) {
                this.setState({ data: data });
            }.bind(this),
            error: function (xhr, status, err) {
                this.setState({data:comments});
                console.error(this.props.url, status, err.toString());
            }.bind(this)
        });
    },
    componentDidMount: function () {
        this.loadCommentsFromServer();
        setInterval(this.loadCommentsFromServer, this.props.pollInterval);

    },
    render: function () {
        return (
            <div className="commentBox">
                <h1>Comments</h1>
                <CommentList data = {this.state.data} />
                <CommentForm onCommentSubmit={this.handleCommentSubmit} />
            </div>
        );

    }
});

var CommentForm = React.createClass({
    getInitialState: function () {
        return { Serialline: '', Serialspeed: '' };
    },
    handleSerialLineChange: function (e) {
        this.setState({ Serialline: e.target.value });
    },
    handleSerialSpeedChange: function (e) {
        this.setState({ Serialspeed: e.target.value });
    },
    handleSubmit: function (e) {
        e.preventDefault();
        var line = this.state.Serialline.trim();
        var speed = this.state.Serialspeed.trim();
        if (!line || !speed) {
            return;
        }
        this.props.onCommentSubmit({ Serialline: line, Serialspeed: speed })
        this.setState({ Serialline: "", Serialspeed: "" });
    },
    render: function () {
        return (
            <form className="commentForm" onSubmit={this.handleSubmit}>
                <input type="text" placeholder="Serial line" value={this.state.Serialline} onChange={this.handleSerialLineChange}/>
                <input type="text" placeholder="Serial speed" value={this.state.Serialspeed} onChange={this.handleSerialSpeedChange}/>
                <input type="submit" value="Post" />
            </form>
        );
    }
});

ReactDOM.render(<CommentBox data={data} url={url} pollInterval={20000}/>, document.getElementById('example5'));

*/