<!doctype html>
<html>
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="initial-scale=1.0">
  <title>index</title>
  <link rel="stylesheet" href="../static/css/standardize.css">
  <link rel="stylesheet" href="../static/css/index-grid.css">
  <link rel="stylesheet" href="../static/css/index.css">
<script src="../static/js/jquery-min.js" type="text/javascript"></script>

</head>
<body class="body index clearfix">
  <div class="element element-1"></div>
  <div class="element element-2" id="scrollTextArea" name="scrollTextArea">抽奖啦</div>
  <button id="start" class="_button _button-1" type="button" onclick="start()">开 始</button>
  <button id="stop" class="_button _button-2" type="button" onclick="stop()">停 止</button>

  <div id="winprize" class="element element-3" >
	<div class="element element-4">~中奖名单~</div>
	<hr>
	<div id="luckPersonList" ></div>
</div>
	
<script type="text/javascript">
<!--

//所有参与抽奖的人员数据
var allPersonList = {{.QueuePersonList}} 

//预设中奖人数
var setLuckTotal = {{.SetLuckPersonNumber}}

var num = allPersonList.length 
var sysTimer 
var randomNumber

//滚动显示列表
function scroll(){ 
    randomNumber = GetRandomNumber(0,num);
	document.getElementById("scrollTextArea").innerHTML =  allPersonList[randomNumber].name; 
}

//启动 
function start(){ 
	clearInterval(sysTimer); 
	sysTimer = setInterval('scroll()',20); //越小随机变换速度越快
} 

//停止
function stop(){ 
	clearInterval(sysTimer); 
	//document.getElementById("luckperson").innerHTML=document.getElementById("scrollTextArea").innerText;
	var person = {
		"serial" : allPersonList[randomNumber].serial,
		"name" : allPersonList[randomNumber].name
	}
	$.ajax({
		url:"/push-person",
		type:"POST",
		dataType:"json",
		contentType:"application/json; charset=utf-8",
		data:JSON.stringify(person),
		success: function(result){
			if (result == -1 ){
					alert("记录中奖名单失败");
			}else{
				allPersonList = null;
				allPersonList = result.TQueuePersonList;
				num = allPersonList.length;
				if(result.LPersonList.length <= setLuckTotal){
					var addHtml="" ;
					for (var i = 0; i<result.LPersonList.length;i++){
						addHtml += "<div>"+(i+1)+"."+result.LPersonList[i].name+"</div>";
					}
					document.getElementById("luckPersonList").innerHTML = addHtml;
				}
				if(result.LPersonList.length >= setLuckTotal){//这里取等于，是防止start在第 setLuckTotal+1 次被点击
					$("#start").prop('onclick',null).off('click');
					$("#stop").prop('onclick',null).off('click');
					alert("抽奖已结束");
				}
			}
		}
	});
} 

//min到max之间随机数
function GetRandomNumber(min,max){ 
	return parseInt(Math.random()*(max-min+1)); 
}

// -->
</script>
 
</body>
</html>