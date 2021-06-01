window.onload = function () {
    document.getElementById("wiki-search-input").focus();
};

function ajax(keyword) { //AJAX request
    $.ajax({
        url: "http://localhost:8080/search?keyword=" + keyword,
        dataType: "json",
        success: function (response) {
            this_data = response.data
            console.log(this_data);
            if (this_data.length === 0) {
                showError(keyword);
            } else {
                showResults(this_data);
            }
        },
        error: function () {
            alert("Error retrieving search results, please refresh the page");
        },
        beforeSend: function (r) {
            // r.setRequestHeader('Access-Control-Allow-Origin', 'http://localhost:8080');
            // r.setRequestHeader('Access-Control-Allow-Methods', 'GET');
            // r.setRequestHeader('Access-Control-Allow-Headers', 'Content-Type, Authorization');

            // r.setRequestHeader('Access-Control-Allow-Origin', '*');
            // r.setRequestHeader('Access-Control-Allow-Methods', 'GET,POST');
            // r.setRequestHeader('Access-Control-Allow-Headers', 'Content-Type');
        }
    });
}

function index(content, title, tag, lang) {
    $.ajax({
        // url: "http://localhost:8080/_index?content=" + content + '&title=' + title + '&tag=' + tag + '&lang=' + lang,
        url: "http://localhost:8080/_index",
        dataType: "json",
        type: "POST",
        data: {'content': content, 'title': title, 'tag': tag, 'lang': lang},
        success: function (response) {
            alert("添加成功");
            var lang = $("#submit-select").val('');
            var title = $("#submit-title").val('');
            var tag = $("#submit-tag").val('');
            var code = $("#submit-content").val('');
        },
        error: function () {
            alert("Error retrieving search results, please refresh the page");
        },
        beforeSend: function (r) {
            // r.setRequestHeader('Access-Control-Allow-Origin', 'http://localhost:8080');
            // r.setRequestHeader('Access-Control-Allow-Methods', 'GET');
            // r.setRequestHeader('Access-Control-Allow-Headers', 'Content-Type, Authorization');

        }
    });
}

function _delete(_id) {
    $.ajax({
        // url: "http://localhost:8080/_index?content=" + content + '&title=' + title + '&tag=' + tag + '&lang=' + lang,
        url: "http://localhost:8080/_delete?id="+_id,
        dataType: "json",
        type: "GET",
        success: function (response) {
            alert("删除成功");
            window.location.reload();
        },
        error: function () {
            alert("Error retrieving search results, please refresh the page");
        },
    });
}

function showResults(this_data) {
    // this_data 格式 [{content: xxx, title: xxx}]
    var total_len = this_data.length
    for (var i = 0; i < total_len; i++) {
        $(".display-results").append("<div class='result-list result-" + i + "'>" + "<span class='result-title title-" + i + "'></span>" + "<br>" + "<span style='white-space:break-spaces;' class='result-snippet snippet-" + i + "'></span>" + "<br>" + "<span class='result-metadata metadata-" + i + "'></span>" + "<br>" +"<button class='result-delete delete-"+i+"'>删除</button>"+ "</div>");
    }

    for (var m = 0; m < total_len; m++) {
        var title = this_data[m]._source.title;
        // var content = this_data[m]._source.content.replace(/\n/g,"<br/>");
        var content = this_data[m]._source.content;
        var _id = this_data[m]._id;
        // var timestamp = _this_data.query.search[m].timestamp;
        timestamp = new Date(_id * 1000);
        //"Wed Aug 27 2014 00:27:15 GMT+0100 (WAT)";
        console.log(timestamp);
        var datestring = ("0" + timestamp.getDate()).slice(-2) + "-" + ("0" + (timestamp.getMonth() + 1)).slice(-2) + "-" +
            timestamp.getFullYear() + " " + ("0" + timestamp.getHours()).slice(-2) + ":" + ("0" + timestamp.getMinutes()).slice(-2);

        $(".title-" + m).html("<a href='http://localhost:8080/assets/html/display.html?id=" + _id + "' target='_blank'>" + title + "</a>");
        // $(".delete-" + m).html("<a href='http://localhost:8080/_delete?id=" + _id + "' target='_blank'>" + '删除' + "</a>");
        $(".snippet-" + m).html(content);
        // $(".metadata-" + m).html((this_data.query.search[m].size/1000).toFixed(0) + "kb (" + this_data.query.search[m].wordcount + " words) - " + timestamp);
        $(".metadata-" + m).html("10 kb (" + "1000 words) - " + datestring);
        $(".result-delete").click(function (event) {
            _delete(_id)
        });
    }
}

function showError(keyword) {
    $(".display-results").append("<div class='error'> <p>Your search <span class='keyword'>" + keyword + "</span> did not match any documents.</p> <p>Suggestions:</p><li>Make sure that all words are spelled correctly.</li><li>Try different keywords.</li><li>Try more general keywords.</li></div> ");
}

$(".result-btn-wiki").click(function (event) {
    event.preventDefault();
    $(".display-results").html("");
    var keyword = $(".result-wiki-search-form-input").val();
    document.getElementById("result-wiki-search-form-input").blur();
    ajax(keyword);
});

$(".btn-wiki").click(function (event) {
    event.preventDefault();
    var keyword = $(".wiki-search-input").val();

    if (keyword !== "") {
        $(".result-wiki-search-form-input").val(keyword);
        $(".home").addClass('hidden');
        $(".result").removeClass('hidden');
        document.getElementById("wiki-search-input").blur();
        $(".wiki-search-input").val("");
        document.getElementById("result-wiki-search-form-input").blur();
        $(".display-results").html("");
        ajax(keyword);
    } else {
        alert("Enter a keyword into the search box");
    }

});

$(".btn-grid").click(function (event) {
    event.preventDefault();
    var lang = $("#submit-select").val();

    var title = $("#submit-title").val();
    var tag = $("#submit-tag").val();
    var code = $("#submit-content").val();

    if (lang === "") {
        alert("Enter a lang into the search box");
        return
    }
    if (title === "") {
        alert("Enter a title into the search box");
        return
    }
    if (tag === "") {
        alert("Enter a tag into the search box");
        return
    }
    if (code === "") {
        alert("Enter a content into the search box");
        return
    }
    console.log(lang)
    console.log(title)
    console.log(tag)
    console.log(code)
    index(code, title, tag, lang);
});
