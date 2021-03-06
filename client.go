package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	addr := "127.0.0.1:8000"
	cli, err := net.Dial("tcp", addr)
	if err != nil {
		log.Println("链接服务器失败!")
		log.Panicln(err)
	}
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	defer cli.Close()
	fmt.Println("connet success")
	go func() {
		said := []string{"国产都是垃圾", "算了吧。还煮的呢。", "八套房说的很对，国产都是垃圾，", "而且苹果现在确实是身份象征，没看明星都用吗", "一直用国产的", "国产的房子更垃圾，", "出个限量版，抢疯了", "斜眼笑", "没觉得有什么垃圾的", "诺基亚，苹果", "只用过这两个牌子", "别的手机我不会用", "安卓会卡，他们都这么说，我没感觉所以没觉得不好", "用过htc", "苹果出了个齐刘海。", "大家都骂丑", "反而之前用那个苹果home建，还没九宫格输入法用这难受，但是你用习惯了就不难受了，早上还有人说没home建用这难受呢", "然后争先模仿", "/笑哭", "xs出来了，那8会不会便宜了苹果系统确实好，不过现在安卓流畅度也不差", "三星怎么不知不觉的就没见人用过了", "你为什么不问X会不会便宜", "我以前用苹果的时候也觉得安卓再也不想用了", "蛋总想买8，不买x吗我是先买的安卓机", "那时候还是2.0系统   后面确实卡", "我一般都是低两个档次", "@瑶瑶", "看到XS我觉得我的手机太卡了，不能用。再看看余额，我觉得我的手机还能用五年。", "感觉安卓的你别买那千元机，正儿八经买个两三千的机子，也还好了", "真能用五年吗", "iphoe6是什么时候出的", "不过苹果许多人买了用好几年", "我周围还有三个人用", "[图片", "我们这便宜货基本一两年就换了", "还有6plus", "华为不诋毁苹果，难道会捧吗？", "然后用7的，和x的就没有那时候多了，主要还是贵", "笑哭", "6普拉斯很好", "现在都可以买，性价比高", "说明一般工薪阶层还停留在6的价格的时候", "现在对手机不讲究，能满足日常用就好", "手机是个人用品，有的人在意，愿意花两个月工资买，有的人无所谓，半个月工资买个就觉得可以了", "这看个人喜好", "我还是喜欢有home键的iphone", "没有刘海的iphone", "你们一般手机用几年", "苹果能用两年以上", "安卓用两年撑死了瞎扯", "我的小苹果用了3年多。", "我这锤子还准备再用2年的", "电池扛不住了。", "我是说我本人", "苹果电池不咋的", "才换的", "又快两年了。", "最怀念4s", "我的华为用到第二年就卡成🐶了", "我的4S", "我爸还在用", "好得很", "苹果除了电池不行，系统还能流畅运行", "安卓机确实", "用不久。", "果断又换回了🍎", "两年的安卓机", "你挂个电话得一分钟", "这样太费电话费", "我觉得两三千的国产机用两年很轻松啊", "哪有这么夸张啊", "可能我不用手机玩游戏，感觉还好", "没那么夸张", "而且华为也好三星也好两年以后手机都看起来廉价感十足，苹果只要不乱跌颜值一般都还在线😂", "我这pro1都两年多了", "还流畅得不得了", "我现在换手机就没用坏过的", "/笑哭", "卡成狗也叫坏", "一般都是电池不行", "不是内存小都没坏过", "电池不行两年还是可以的", "颜值在我这第一", "我每次换手机都是电池", "现在用手机我都不贴膜的", "@二狗蛋", "随身带充电线啊", "充电宝", "用苹果就得随身带。", "没办法。", "/擦汗", "苹果的电池有那么夸张么", "我这vivo有时候充一次电用两天", "屏幕不要太亮，很费电", "我用过3年", "电池不行了", "许多人估计对手机要求不高"}

		for {
			cont := []byte(said[rand.Intn(len(said))] + "\r")
			cli.Write(cont)
			time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
		}

	}()
	// 按行读取服务器的响应
	r := bufio.NewReader(cli)
	go func() {
		<-sigs
		os.Exit(500)
	}()
	for {
		str, err := r.ReadString('\r')
		if err == io.EOF {
			cli.Close()
		}
		fmt.Println(str)
	}

}
