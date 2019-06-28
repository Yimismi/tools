initAnalytics();

function initAnalytics()
{
    (function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
        (i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
        m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
    })(window,document,'script','//www.google-analytics.com/analytics.js','ga');

    ga('create', 'UA-86578-22', 'auto');
    ga('send', 'pageview');
}
$(function()
{
    const exampleSql =
`CREATE TABLE pet (
    name VARCHAR(20), 
    owner VARCHAR(20),
    species VARCHAR(20), 
    sex CHAR(1), 
    birth DATE, 
    death DATE
);
`;
    const toolName = "sql2go";
    const emptyInputMsg = "Paste Sql here";
    const emptyOutputMsg = "Go will appear here";
    const formattedEmptyInputMsg = '<span style="color: #777;">'+emptyInputMsg+'</span>';
    const formattedEmptyOutputMsg = '<span style="color: #777;">'+emptyOutputMsg+'</span>';

    function doConversion()
    {
        var input = $('#input').text().trim();
        if (!input || input == emptyInputMsg)
        {
            $('#output').html(formattedEmptyOutputMsg);
            return;
        }

        let output = getGo(input);

        if (!!output.error) {
            console.log(output.error);
            $("#output").html("<p style='color: #BC6060'>error: " + output.error + "</p>");
            return;
        }
        output = output.output;

        if (output) {
            var coloredOutput = hljs.highlight("go", output);
            var disText = "";
            if (!coloredOutput) {
                disTest = output
            } else {
                disText = coloredOutput.value
            }
            $('#output').html(disText);
        }
    }

    // Hides placeholder text
    $('#input').on('focus', function()
    {
        var val = $(this).text();
        if (!val)
        {
            $(this).html(formattedEmptyInputMsg);
            $('#output').html(formattedEmptyOutputMsg);
        }
        else if (val == emptyInputMsg)
            $(this).html("");
    });

    // Shows placeholder text
    $('#input').on('blur', function()
    {
        var val = $(this).text();
        if (!val)
        {
            $(this).html(formattedEmptyInputMsg);
            $('#output').html(formattedEmptyOutputMsg);
        }
    }).blur();

    // If tab is pressed, insert a tab instead of focusing on next element
    $('#input').keydown(function(e)
    {
        if (e.keyCode == 9)
        {
            document.execCommand('insertHTML', false, '&#009'); // insert tab
            e.preventDefault(); // don't go to next element
        }
    });

    // Automatically do the conversion on paste or change
    $('#input').keyup(function()
    {
        doConversion();
    });

    // Also do conversion when inlining preference changes
    $('#inline').change(function()
    {
        doConversion();
    })

    // Highlights the output for the user
    $('#output').click(function()
    {
        if (document.selection)
        {
            var range = document.body.createTextRange();
            range.moveToElementText(this);
            range.select();
        }
        else if (window.getSelection)
        {
            var range = document.createRange();
            range.selectNode(this);
            var sel = window.getSelection();
            sel.removeAllRanges(); // required as of Chrome 60: https://www.chromestatus.com/features/6680566019653632
            sel.addRange(range);
        }
    });

    // Fill in sample JSON if the user wants to see an example
    $('#sample1').click(function()
    {
        $('#input').text(exampleSql).keyup();
    });

    function xhrRequest() {
        var http;
        var activeX = ['MSXML2.XMLHTTP.3.0', 'MSXML2.XMLHTTP', 'Microsoft.XMLHTTP'];

        try {
            http = new XMLHttpRequest();
        } catch (e) {
            for (var i = 0; i < activeX.length; ++i) {
                try {
                    http = new ActiveXObject(activeX[i]);
                    break;
                } catch (e) { }
            }
        } finally {
            return http;
        }
    }
    function getGo(src, args) {
        if (!src) {
            return ""
        }
        xhr = xhrRequest();
        if (!!xhr) {
            xhr = new XMLHttpRequest();
            xhr.open("POST", getUrlRelativePath(), false);
            xhr.setRequestHeader('content-type', 'application/json');
            reqBody = JSON.stringify(getRequest(src, args))
            xhr.send(reqBody);
            return JSON.parse(xhr.responseText)
        } else {
            return {}.error = "xhr error";
        }

    }

    function getRequest(src, args) {
        a = {};
        a.src = src;
        Object.assign(a, args);
        return a;
    }

    function getUrlRelativePath()
    {
        var url = document.location.toString();
        var arrUrl = url.split("//");

        var start = arrUrl[1].indexOf("/");
        var relUrl = arrUrl[1].substring(start);//stop省略，截取从start开始到结尾的所有字符

        if(relUrl.indexOf("?") != -1){
            relUrl = relUrl.split("?")[0];
        }
        return relUrl;
    }

});