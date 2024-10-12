// -- gbk --

function goTonr(){
    this.headers = {
        "accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
        "accept-language": "zh-CN,zh;q=0.9",
        "cache-control": "no-cache",
        "pragma": "no-cache",
        "priority": "u=0, i",
        "referer": "https://travel.qunar.com/search/gonglue/22-shanghai-299878/hot_heat/3.htm",
        "^sec-ch-ua": "^\\^Google",
        "sec-ch-ua-mobile": "?0",
        "^sec-ch-ua-platform": "^\\^Windows^^^",
        "sec-fetch-dest": "document",
        "sec-fetch-mode": "navigate",
        "sec-fetch-site": "same-origin",
        "sec-fetch-user": "?1",
        "upgrade-insecure-requests": "1",
        "user-agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36"
    }
    this.getPageInfo = function (url){
        console.log("开始异步Promise请求",url)
        var d = new Promise(function(){
            req = cyhttp.get(url,{
                "headers": this.headers,
                "timeout": 30,
                // "proxies":"http://127.0.0.1:8888"
            })
            cyout(req.text)
            // cbb_a传递到then的第一个函数 || cbb_b传递到then的第二个函数
            cbb_a(req.text)
        })
        d.then(function (text){
            parseHTML = etree.HTML(text)
            title = parseHTML.xpath('//ul[@class="b_strategy_list "]/li//h2')
            urlList = parseHTML.xpath('//ul[@class="b_strategy_list "]/li//h2/a/@href')
            var b = fs.open("./a.csv",{"ms":"a"})
            for (var i=0 ; i<title.length; i++){
                b.write( [title[i].xpath('.//text()').join(" "), "https://travel.qunar.com"+urlList[i]].join(",")+"\n")
                console.log(title[i].xpath('.//text()').join(" "))

            }


            b.close()
            console.log("结束异步请求")
        })

    }


}

var b = fs.open("./a.csv",{"ms":"w"})
b.write(["标题","链接"].join(",")+"\n")
b.close()
var ff = new goTonr()
for (var i =1; i<10; i++){
    st = Date.now()
    ff.getPageInfo("https://travel.qunar.com/search/gonglue/22-shanghai-299878/hot_heat/"+i+".htm")
    // 等待前面异步操作结束
    wait()
    cyout("请求时间", Date.now() - st, "ms")
    debugger
    // 休眠10000ms
    Date.sleep(10000)
}
