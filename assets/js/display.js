function parseReq(paras) {
    var url = window.location.href
    var paraString = url.substring(url.indexOf("?") + 1, url.length).split("&");
    var paraObj = {}
    for (i = 0; j = paraString[i]; i++) {
        paraObj[j.substring(0, j.indexOf("=")).toLowerCase()] = j.substring(j.indexOf("=") + 1, j.length);
    }
    var returnValue = paraObj[paras.toLowerCase()];
    if (typeof (returnValue) == "undefined") {
        return "";
    } else {
        return returnValue;
    }
}

function ajax(keyword) { //AJAX request
    $.ajax({
        url: "http://localhost:8080/search?this_id=" + keyword,
        dataType: "json",
        success: function (response) {
            this_data = response.data
            console.log(this_data);
            if (this_data.length === 0) {
                alert("Error retrieving search results, please refresh the page");
            } else {
                showResults(this_data);
            }
        },
        error: function () {
            alert("Error retrieving search results, please refresh the page");
        },
        beforeSend: function (r) {
            r.setRequestHeader('Access-Control-Allow-Origin', 'http://localhost:8080');
            r.setRequestHeader('Access-Control-Allow-Methods', 'GET');
            r.setRequestHeader('Access-Control-Allow-Headers', 'Content-Type, Authorization');

        }
    });
}
function showResults(this_data) {
    // this_data 格式 [{content: xxx, title: xxx}]
    $("#code-block").append("<div><span class='title'></span><br><span class='content' style='white-space:break-spaces;'></span><br></div>");
    var content = this_data[0]._source.content;
    $(".content").html(content);
    // var total_len = this_data.length
    // for (var i = 0; i < total_len; i++) {
    //     $(".code-block").append("<div class='result-list result-" + i + "'>" + "<span class='result-title title-" + i + "'></span>" + "<br>" + "<span style='white-space:break-spaces;' class='result-snippet snippet-" + i + "'></span>" + "<br>" + "<span class='result-metadata metadata-" + i + "'></span>" + "<br>" +"<span class='result-delete delete-"+i+"'></span>"+ "</div>");
    // }
    //
    // for (var m = 0; m < total_len; m++) {
    //     var title = this_data[m]._source.title;
    //     // var content = this_data[m]._source.content.replace(/\n/g,"<br/>");
    //     var content = this_data[m]._source.content;
    //     var _id = this_data[m]._id;
    //     // var timestamp = _this_data.query.search[m].timestamp;
    //     timestamp = new Date(_id * 1000);
    //     //"Wed Aug 27 2014 00:27:15 GMT+0100 (WAT)";
    //     console.log(timestamp);
    //     var datestring = ("0" + timestamp.getDate()).slice(-2) + "-" + ("0" + (timestamp.getMonth() + 1)).slice(-2) + "-" +
    //         timestamp.getFullYear() + " " + ("0" + timestamp.getHours()).slice(-2) + ":" + ("0" + timestamp.getMinutes()).slice(-2);
    //
    //     $(".title-" + m).html("<a href='http://localhost:8080/search?this_id=" + _id + "' target='_blank'>" + title + "</a>");
    //     $(".delete-" + m).html("<a href='http://localhost:8080/_delete?id=" + _id + "' target='_blank'>" + '删除' + "</a>");
    //     $(".snippet-" + m).html(content);
    //     // $(".metadata-" + m).html((this_data.query.search[m].size/1000).toFixed(0) + "kb (" + this_data.query.search[m].wordcount + " words) - " + timestamp);
    //     $(".metadata-" + m).html("10 kb (" + "1000 words) - " + datestring);
    // }
}

$(document).ready(function () {
    var _id = parseReq('id')
    console.log(_id)
    ajax(_id)
});
