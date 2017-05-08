<!DOCTYPE html>
<html lang="zh-cn">
  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <!-- The above 3 meta tags *must* come first in the head; any other head content must come *after* these tags -->
    <meta name="description" content="">
    <meta name="author" content="">
    <link rel="icon" href="/static/img/logo.JPG">

    <title>致简生活</title>

    <!-- Bootstrap core CSS -->
    <link href="/static/css/bootstrap.min.css" rel="stylesheet">

    <!-- Custom styles for this template -->
    <link href="/static/css/cover.css" rel="stylesheet">

    <!-- Just for debugging purposes. Don't actually copy these 2 lines! -->
    <!--[if lt IE 9]><script src="../../assets/js/ie8-responsive-file-warning.js"></script><![endif]-->
    <script src="/static/js/ie-emulation-modes-warning.js"></script>

    <!-- HTML5 shim and Respond.js for IE8 support of HTML5 elements and media queries -->
    <!--[if lt IE 9]>
      <script src="https://oss.maxcdn.com/html5shiv/3.7.2/html5shiv.min.js"></script>
      <script src="https://oss.maxcdn.com/respond/1.4.2/respond.min.js"></script>
    <![endif]-->
  </head>

  <body>

    <div class="site-wrapper">

      <div class="site-wrapper-inner">

        <div class="cover-container">

          <div class="inner cover">
            <h1 class="logo">致简生活</h1>
            <img class="wxmp" src="/static/img/qrcode_for_gh_ec9c7cdf9ce4_430.jpg"></img>
            <img class="weapp" src="/static/img/gh_8741695dd43d_430.jpg"></img>
            <p class="lead">您可以通过公众号或者小程序查询天气，快递等，为您提供更多更精彩的生活服务！</p>
            <p class="lead">
              <a href="https://github.com/ToJaneLife" class="btn btn-lg btn-default btn-outline" target="_self">项目地址</a>
            </p>

          </div>

          <div class="mastfoot">
            <div class="inner">
              <div class="author">
                网址:
                <a href="http://{{.Website}}">{{.Website}}</a> /
                联系我:
                <a class="email" href="mailto:{{.Email}}">{{.Email}}</a>
              </div>
            </div>
          </div>

        </div>

      </div>

    </div>

    <!-- Bootstrap core JavaScript
    ================================================== -->
    <!-- Placed at the end of the document so the pages load faster -->
    <script src="/static/js/jquery-2.2.3.min.js"></script>
    <script src="/static/js/bootstrap.min.js"></script>
    <!-- IE10 viewport hack for Surface/desktop Windows 8 bug -->
    <script src="/static/js/ie10-viewport-bug-workaround.js"></script>
    <script>
        var imageNum=1;
        $("#changeback").click(function(){
            imageNum++;
            if(imageNum==4){
                imageNum=1;
            }
            console.log(imageNum);
            switch(imageNum){
                case 3:
                $("body,html").css({
                    "backgroundColor":"#333","background-image":"none"}
                    );
                event.preventDefault();
                break;
                case 2:
                $("body,html").css("background-image","url('/static/img/back001.jpg')");
                event.preventDefault();
                break;
                case 1:
                $("body,html").css("background-image","url('/static/img/back002.jpg')");
                event.preventDefault();
                break;
            }
        })

    </script>
  </body>
</html>
