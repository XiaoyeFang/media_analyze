package operation

import (
	"fmt"
	"pure-media/protos"
	"testing"
)

func TestUrlQueue(t *testing.T) {
	Durl := []string{
		"https://ss0.baidu.com/6ONWsjip0QIZ8tyhnq/it/u=2216102322,3587941637&fm=173&app=25&f=JPEG?w=218&h=146&s=B8C5A14C5FE19F6C14DFED01030070C9",
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1539671813993&di=080422caa5b56f29ae9b3c9dfbda2bf0&imgtype=0&src=http%3A%2F%2Fimg.zcool.cn%2Fcommunity%2F0139ef571adfda32f875a3998b57b7.gif",
		"https://apkpure.com/shadow-of-death-dark-knight-stickman-fighting/com.Zonmob.ShadowofDeath.FightingGames",
	}
	fmt.Println((Durl))
	select {}
}

func TestAsynUrlQueue(t *testing.T) {
	Durl := []string{
		"https://ss0.baidu.com/6ONWsjip0QIZ8tyhnq/it/u=2216102322,3587941637&fm=173&app=25&f=JPEG?w=218&h=146&s=B8C5A14C5FE19F6C14DFED01030070C9",
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1539671813993&di=080422caa5b56f29ae9b3c9dfbda2bf0&imgtype=0&src=http%3A%2F%2Fimg.zcool.cn%2Fcommunity%2F0139ef571adfda32f875a3998b57b7.gif",
		"https://image.winudf.com/v2/image/Y29tLm5jc29mdC5haW9uLmxlZ2lvbnMub2Yud2FyX2Jhbm5lcl8xNTM3NDc5MzcxXzAxMA/banner.jpg?w=850&fakeurl=1&type=.jpg",
		"https://ss0.baidu.com/6ONWsjip0QIZ8tyhnq/it/u=2216102322,3587941637&fm=173&app=25&f=JPEG?w=218&h=146&s=B8C5A14C5FE19F6C14DFED01030070C9",
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1539671813993&di=080422caa5b56f29ae9b3c9dfbda2bf0&imgtype=0&src=http%3A%2F%2Fimg.zcool.cn%2Fcommunity%2F0139ef571adfda32f875a3998b57b7.gif",
		"https://image.winudf.com/v2/image/Y29tLm5jc29mdC5haW9uLmxlZ2lvbnMub2Yud2FyX2Jhbm5lcl8xNTM3NDc5MzcxXzAxMA/banner.jpg?w=850&fakeurl=1&type=.jpg",
		"https://ss0.baidu.com/6ONWsjip0QIZ8tyhnq/it/u=2216102322,3587941637&fm=173&app=25&f=JPEG?w=218&h=146&s=B8C5A14C5FE19F6C14DFED01030070C9",
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1539671813993&di=080422caa5b56f29ae9b3c9dfbda2bf0&imgtype=0&src=http%3A%2F%2Fimg.zcool.cn%2Fcommunity%2F0139ef571adfda32f875a3998b57b7.gif",
		"https://image.winudf.com/v2/image/Y29tLm5jc29mdC5haW9uLmxlZ2lvbnMub2Yud2FyX2Jhbm5lcl8xNTM3NDc5MzcxXzAxMA/banner.jpg?w=850&fakeurl=1&type=.jpg",
		"https://ss0.baidu.com/6ONWsjip0QIZ8tyhnq/it/u=2216102322,3587941637&fm=173&app=25&f=JPEG?w=218&h=146&s=B8C5A14C5FE19F6C14DFED01030070C9",
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1539671813993&di=080422caa5b56f29ae9b3c9dfbda2bf0&imgtype=0&src=http%3A%2F%2Fimg.zcool.cn%2Fcommunity%2F0139ef571adfda32f875a3998b57b7.gif",
		"https://image.winudf.com/v2/image/Y29tLm5jc29mdC5haW9uLmxlZ2lvbnMub2Yud2FyX2Jhbm5lcl8xNTM3NDc5MzcxXzAxMA/banner.jpg?w=850&fakeurl=1&type=.jpg",
		"https://ss0.baidu.com/6ONWsjip0QIZ8tyhnq/it/u=2216102322,3587941637&fm=173&app=25&f=JPEG?w=218&h=146&s=B8C5A14C5FE19F6C14DFED01030070C9",
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1539671813993&di=080422caa5b56f29ae9b3c9dfbda2bf0&imgtype=0&src=http%3A%2F%2Fimg.zcool.cn%2Fcommunity%2F0139ef571adfda32f875a3998b57b7.gif",
		"https://image.winudf.com/v2/image/Y29tLm5jc29mdC5haW9uLmxlZ2lvbnMub2Yud2FyX2Jhbm5lcl8xNTM3NDc5MzcxXzAxMA/banner.jpg?w=850&fakeurl=1&type=.jpg",
		"https://ss0.baidu.com/6ONWsjip0QIZ8tyhnq/it/u=2216102322,3587941637&fm=173&app=25&f=JPEG?w=218&h=146&s=B8C5A14C5FE19F6C14DFED01030070C9",
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1539671813993&di=080422caa5b56f29ae9b3c9dfbda2bf0&imgtype=0&src=http%3A%2F%2Fimg.zcool.cn%2Fcommunity%2F0139ef571adfda32f875a3998b57b7.gif",
		"https://image.winudf.com/v2/image/Y29tLm5jc29mdC5haW9uLmxlZ2lvbnMub2Yud2FyX2Jhbm5lcl8xNTM3NDc5MzcxXzAxMA/banner.jpg?w=850&fakeurl=1&type=.jpg",
		"https://ss0.baidu.com/6ONWsjip0QIZ8tyhnq/it/u=2216102322,3587941637&fm=173&app=25&f=JPEG?w=218&h=146&s=B8C5A14C5FE19F6C14DFED01030070C9",
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1539671813993&di=080422caa5b56f29ae9b3c9dfbda2bf0&imgtype=0&src=http%3A%2F%2Fimg.zcool.cn%2Fcommunity%2F0139ef571adfda32f875a3998b57b7.gif",
		"https://image.winudf.com/v2/image/Y29tLm5jc29mdC5haW9uLmxlZ2lvbnMub2Yud2FyX2Jhbm5lcl8xNTM3NDc5MzcxXzAxMA/banner.jpg?w=850&fakeurl=1&type=.jpg",
		"https://ss0.baidu.com/6ONWsjip0QIZ8tyhnq/it/u=2216102322,3587941637&fm=173&app=25&f=JPEG?w=218&h=146&s=B8C5A14C5FE19F6C14DFED01030070C9",
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1539671813993&di=080422caa5b56f29ae9b3c9dfbda2bf0&imgtype=0&src=http%3A%2F%2Fimg.zcool.cn%2Fcommunity%2F0139ef571adfda32f875a3998b57b7.gif",
		"https://image.winudf.com/v2/image/Y29tLm5jc29mdC5haW9uLmxlZ2lvbnMub2Yud2FyX2Jhbm5lcl8xNTM3NDc5MzcxXzAxMA/banner.jpg?w=850&fakeurl=1&type=.jpg",
		"https://ss0.baidu.com/6ONWsjip0QIZ8tyhnq/it/u=2216102322,3587941637&fm=173&app=25&f=JPEG?w=218&h=146&s=B8C5A14C5FE19F6C14DFED01030070C9",
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1539671813993&di=080422caa5b56f29ae9b3c9dfbda2bf0&imgtype=0&src=http%3A%2F%2Fimg.zcool.cn%2Fcommunity%2F0139ef571adfda32f875a3998b57b7.gif",
		"https://image.winudf.com/v2/image/Y29tLm5jc29mdC5haW9uLmxlZ2lvbnMub2Yud2FyX2Jhbm5lcl8xNTM3NDc5MzcxXzAxMA/banner.jpg?w=850&fakeurl=1&type=.jpg",
		"https://ss0.baidu.com/6ONWsjip0QIZ8tyhnq/it/u=2216102322,3587941637&fm=173&app=25&f=JPEG?w=218&h=146&s=B8C5A14C5FE19F6C14DFED01030070C9",
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1539671813993&di=080422caa5b56f29ae9b3c9dfbda2bf0&imgtype=0&src=http%3A%2F%2Fimg.zcool.cn%2Fcommunity%2F0139ef571adfda32f875a3998b57b7.gif",
		"https://image.winudf.com/v2/image/Y29tLm5jc29mdC5haW9uLmxlZ2lvbnMub2Yud2FyX2Jhbm5lcl8xNTM3NDc5MzcxXzAxMA/banner.jpg?w=850&fakeurl=1&type=.jpg",
		"https://ss0.baidu.com/6ONWsjip0QIZ8tyhnq/it/u=2216102322,3587941637&fm=173&app=25&f=JPEG?w=218&h=146&s=B8C5A14C5FE19F6C14DFED01030070C9",
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1539671813993&di=080422caa5b56f29ae9b3c9dfbda2bf0&imgtype=0&src=http%3A%2F%2Fimg.zcool.cn%2Fcommunity%2F0139ef571adfda32f875a3998b57b7.gif",
		"https://image.winudf.com/v2/image/Y29tLm5jc29mdC5haW9uLmxlZ2lvbnMub2Yud2FyX2Jhbm5lcl8xNTM3NDc5MzcxXzAxMA/banner.jpg?w=850&fakeurl=1&type=.jpg",
		"https://ss0.baidu.com/6ONWsjip0QIZ8tyhnq/it/u=2216102322,3587941637&fm=173&app=25&f=JPEG?w=218&h=146&s=B8C5A14C5FE19F6C14DFED01030070C9",
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1539671813993&di=080422caa5b56f29ae9b3c9dfbda2bf0&imgtype=0&src=http%3A%2F%2Fimg.zcool.cn%2Fcommunity%2F0139ef571adfda32f875a3998b57b7.gif",
		"https://image.winudf.com/v2/image/Y29tLm5jc29mdC5haW9uLmxlZ2lvbnMub2Yud2FyX2Jhbm5lcl8xNTM3NDc5MzcxXzAxMA/banner.jpg?w=850&fakeurl=1&type=.jpg",
		"https://ss0.baidu.com/6ONWsjip0QIZ8tyhnq/it/u=2216102322,3587941637&fm=173&app=25&f=JPEG?w=218&h=146&s=B8C5A14C5FE19F6C14DFED01030070C9",
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1539671813993&di=080422caa5b56f29ae9b3c9dfbda2bf0&imgtype=0&src=http%3A%2F%2Fimg.zcool.cn%2Fcommunity%2F0139ef571adfda32f875a3998b57b7.gif",
		"https://image.winudf.com/v2/image/Y29tLm5jc29mdC5haW9uLmxlZ2lvbnMub2Yud2FyX2Jhbm5lcl8xNTM3NDc5MzcxXzAxMA/banner.jpg?w=850&fakeurl=1&type=.jpg",
		"https://ss0.baidu.com/6ONWsjip0QIZ8tyhnq/it/u=2216102322,3587941637&fm=173&app=25&f=JPEG?w=218&h=146&s=B8C5A14C5FE19F6C14DFED01030070C9",
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1539671813993&di=080422caa5b56f29ae9b3c9dfbda2bf0&imgtype=0&src=http%3A%2F%2Fimg.zcool.cn%2Fcommunity%2F0139ef571adfda32f875a3998b57b7.gif",
		"https://image.winudf.com/v2/image/Y29tLm5jc29mdC5haW9uLmxlZ2lvbnMub2Yud2FyX2Jhbm5lcl8xNTM3NDc5MzcxXzAxMA/banner.jpg?w=850&fakeurl=1&type=.jpg",
		"https://ss0.baidu.com/6ONWsjip0QIZ8tyhnq/it/u=2216102322,3587941637&fm=173&app=25&f=JPEG?w=218&h=146&s=B8C5A14C5FE19F6C14DFED01030070C9",
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1539671813993&di=080422caa5b56f29ae9b3c9dfbda2bf0&imgtype=0&src=http%3A%2F%2Fimg.zcool.cn%2Fcommunity%2F0139ef571adfda32f875a3998b57b7.gif",
		"https://image.winudf.com/v2/image/Y29tLm5jc29mdC5haW9uLmxlZ2lvbnMub2Yud2FyX2Jhbm5lcl8xNTM3NDc5MzcxXzAxMA/banner.jpg?w=850&fakeurl=1&type=.jpg",
		"https://ss0.baidu.com/6ONWsjip0QIZ8tyhnq/it/u=2216102322,3587941637&fm=173&app=25&f=JPEG?w=218&h=146&s=B8C5A14C5FE19F6C14DFED01030070C9",
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1539671813993&di=080422caa5b56f29ae9b3c9dfbda2bf0&imgtype=0&src=http%3A%2F%2Fimg.zcool.cn%2Fcommunity%2F0139ef571adfda32f875a3998b57b7.gif",
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1539671813993&di=080422caa5b56f29ae9b3c9dfbda2bf0&imgtype=0&src=http%3A%2F%2Fimg.zcool.cn%2Fcommunity%2F0139ef571adfda32f875a3998b57b7.gif",
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1539671813993&di=080422caa5b56f29ae9b3c9dfbda2bf0&imgtype=0&src=http%3A%2F%2Fimg.zcool.cn%2Fcommunity%2F0139ef571adfda32f875a3998b57b7.gif",
		"https://image.winudf.com/v2/image/Y29tLm5jc29mdC5haW9uLmxlZ2lvbnMub2Yud2FyX2Jhbm5lcl8xNTM3NDc5MzcxXzAxMA/banner.jpg?w=850&fakeurl=1&type=.jpg",
		"http://up.36992.com/pic_source/5a/f0/19/5af01952e75f06cdeb5dedccfba77c70.jpg",
		"http://up.36992.com/pic_source/5a/f0/19/5af01952e75f06cdeb5dedccfba77c70.jpg",
		"http://up.36992.com/pic_source/5a/f0/19/5af01952e75f06cdeb5dedccfba77c70.jpg",
		"http://up.36992.com/pic_source/5a/f0/19/5af01952e75f06cdeb5dedccfba77c70.jpg",
		"http://up.36992.com/pic_source/5a/f0/19/5af01952e75f06cdeb5dedccfba77c70.jpg",
		"http://up.36992.com/pic_source/5a/f0/19/5af01952e75f06cdeb5dedccfba77c70.jpg",
		"http://up.36992.com/pic_source/5a/f0/19/5af01952e75f06cdeb5dedccfba77c70.jpg",
	}
	t.Log(len(Durl))
	reply, err := AsynUrlQueue(Durl)
	if err != nil {
		t.Error(err)
	}
	t.Log("reply = ", reply)

}

func TestUploadUrlQueue(t *testing.T) {
	reqs := &protos.ListUploadFileRequest{}
	Durl := []string{
		"https://ss0.baidu.com/6ONWsjip0QIZ8tyhnq/it/u=2216102322,3587941637&fm=173&app=25&f=JPEG?w=218&h=146&s=B8C5A14C5FE19F6C14DFED01030070C9",
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1539671813993&di=080422caa5b56f29ae9b3c9dfbda2bf0&imgtype=0&src=http%3A%2F%2Fimg.zcool.cn%2Fcommunity%2F0139ef571adfda32f875a3998b57b7.gif",
		"https://image.winudf.com/v2/image/Y29tLm5jc29mdC5haW9uLmxlZ2lvbnMub2Yud2FyX2Jhbm5lcl8xNTM3NDc5MzcxXzAxMA/banner.jpg?w=850&fakeurl=1&type=.jpg",
		"https://ss0.baidu.com/6ONWsjip0QIZ8tyhnq/it/u=2216102322,3587941637&fm=173&app=25&f=JPEG?w=218&h=146&s=B8C5A14C5FE19F6C14DFED01030070C9",
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1539671813993&di=080422caa5b56f29ae9b3c9dfbda2bf0&imgtype=0&src=http%3A%2F%2Fimg.zcool.cn%2Fcommunity%2F0139ef571adfda32f875a3998b57b7.gif",
		"https://image.winudf.com/v2/image/Y29tLm5jc29mdC5haW9uLmxlZ2lvbnMub2Yud2FyX2Jhbm5lcl8xNTM3NDc5MzcxXzAxMA/banner.jpg?w=850&fakeurl=1&type=.jpg",
		"https://ss0.baidu.com/6ONWsjip0QIZ8tyhnq/it/u=2216102322,3587941637&fm=173&app=25&f=JPEG?w=218&h=146&s=B8C5A14C5FE19F6C14DFED01030070C9",
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1539671813993&di=080422caa5b56f29ae9b3c9dfbda2bf0&imgtype=0&src=http%3A%2F%2Fimg.zcool.cn%2Fcommunity%2F0139ef571adfda32f875a3998b57b7.gif",
		"https://image.winudf.com/v2/image/Y29tLm5jc29mdC5haW9uLmxlZ2lvbnMub2Yud2FyX2Jhbm5lcl8xNTM3NDc5MzcxXzAxMA/banner.jpg?w=850&fakeurl=1&type=.jpg",
		"https://ss0.baidu.com/6ONWsjip0QIZ8tyhnq/it/u=2216102322,3587941637&fm=173&app=25&f=JPEG?w=218&h=146&s=B8C5A14C5FE19F6C14DFED01030070C9",
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1539671813993&di=080422caa5b56f29ae9b3c9dfbda2bf0&imgtype=0&src=http%3A%2F%2Fimg.zcool.cn%2Fcommunity%2F0139ef571adfda32f875a3998b57b7.gif",
		"https://image.winudf.com/v2/image/Y29tLm5jc29mdC5haW9uLmxlZ2lvbnMub2Yud2FyX2Jhbm5lcl8xNTM3NDc5MzcxXzAxMA/banner.jpg?w=850&fakeurl=1&type=.jpg",
		"https://ss0.baidu.com/6ONWsjip0QIZ8tyhnq/it/u=2216102322,3587941637&fm=173&app=25&f=JPEG?w=218&h=146&s=B8C5A14C5FE19F6C14DFED01030070C9",
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1539671813993&di=080422caa5b56f29ae9b3c9dfbda2bf0&imgtype=0&src=http%3A%2F%2Fimg.zcool.cn%2Fcommunity%2F0139ef571adfda32f875a3998b57b7.gif",
		"https://image.winudf.com/v2/image/Y29tLm5jc29mdC5haW9uLmxlZ2lvbnMub2Yud2FyX2Jhbm5lcl8xNTM3NDc5MzcxXzAxMA/banner.jpg?w=850&fakeurl=1&type=.jpg",
		"https://ss0.baidu.com/6ONWsjip0QIZ8tyhnq/it/u=2216102322,3587941637&fm=173&app=25&f=JPEG?w=218&h=146&s=B8C5A14C5FE19F6C14DFED01030070C9",
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1539671813993&di=080422caa5b56f29ae9b3c9dfbda2bf0&imgtype=0&src=http%3A%2F%2Fimg.zcool.cn%2Fcommunity%2F0139ef571adfda32f875a3998b57b7.gif",
		"https://image.winudf.com/v2/image/Y29tLm5jc29mdC5haW9uLmxlZ2lvbnMub2Yud2FyX2Jhbm5lcl8xNTM3NDc5MzcxXzAxMA/banner.jpg?w=850&fakeurl=1&type=.jpg",
		"https://ss0.baidu.com/6ONWsjip0QIZ8tyhnq/it/u=2216102322,3587941637&fm=173&app=25&f=JPEG?w=218&h=146&s=B8C5A14C5FE19F6C14DFED01030070C9",
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1539671813993&di=080422caa5b56f29ae9b3c9dfbda2bf0&imgtype=0&src=http%3A%2F%2Fimg.zcool.cn%2Fcommunity%2F0139ef571adfda32f875a3998b57b7.gif",
		"https://image.winudf.com/v2/image/Y29tLm5jc29mdC5haW9uLmxlZ2lvbnMub2Yud2FyX2Jhbm5lcl8xNTM3NDc5MzcxXzAxMA/banner.jpg?w=850&fakeurl=1&type=.jpg",
		"https://ss0.baidu.com/6ONWsjip0QIZ8tyhnq/it/u=2216102322,3587941637&fm=173&app=25&f=JPEG?w=218&h=146&s=B8C5A14C5FE19F6C14DFED01030070C9",
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1539671813993&di=080422caa5b56f29ae9b3c9dfbda2bf0&imgtype=0&src=http%3A%2F%2Fimg.zcool.cn%2Fcommunity%2F0139ef571adfda32f875a3998b57b7.gif",
		"https://image.winudf.com/v2/image/Y29tLm5jc29mdC5haW9uLmxlZ2lvbnMub2Yud2FyX2Jhbm5lcl8xNTM3NDc5MzcxXzAxMA/banner.jpg?w=850&fakeurl=1&type=.jpg",
		"https://ss0.baidu.com/6ONWsjip0QIZ8tyhnq/it/u=2216102322,3587941637&fm=173&app=25&f=JPEG?w=218&h=146&s=B8C5A14C5FE19F6C14DFED01030070C9",
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1539671813993&di=080422caa5b56f29ae9b3c9dfbda2bf0&imgtype=0&src=http%3A%2F%2Fimg.zcool.cn%2Fcommunity%2F0139ef571adfda32f875a3998b57b7.gif",
		"https://image.winudf.com/v2/image/Y29tLm5jc29mdC5haW9uLmxlZ2lvbnMub2Yud2FyX2Jhbm5lcl8xNTM3NDc5MzcxXzAxMA/banner.jpg?w=850&fakeurl=1&type=.jpg",
		"https://ss0.baidu.com/6ONWsjip0QIZ8tyhnq/it/u=2216102322,3587941637&fm=173&app=25&f=JPEG?w=218&h=146&s=B8C5A14C5FE19F6C14DFED01030070C9",
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1539671813993&di=080422caa5b56f29ae9b3c9dfbda2bf0&imgtype=0&src=http%3A%2F%2Fimg.zcool.cn%2Fcommunity%2F0139ef571adfda32f875a3998b57b7.gif",
		"https://image.winudf.com/v2/image/Y29tLm5jc29mdC5haW9uLmxlZ2lvbnMub2Yud2FyX2Jhbm5lcl8xNTM3NDc5MzcxXzAxMA/banner.jpg?w=850&fakeurl=1&type=.jpg",
		"https://ss0.baidu.com/6ONWsjip0QIZ8tyhnq/it/u=2216102322,3587941637&fm=173&app=25&f=JPEG?w=218&h=146&s=B8C5A14C5FE19F6C14DFED01030070C9",
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1539671813993&di=080422caa5b56f29ae9b3c9dfbda2bf0&imgtype=0&src=http%3A%2F%2Fimg.zcool.cn%2Fcommunity%2F0139ef571adfda32f875a3998b57b7.gif",
		"https://image.winudf.com/v2/image/Y29tLm5jc29mdC5haW9uLmxlZ2lvbnMub2Yud2FyX2Jhbm5lcl8xNTM3NDc5MzcxXzAxMA/banner.jpg?w=850&fakeurl=1&type=.jpg",
		"https://ss0.baidu.com/6ONWsjip0QIZ8tyhnq/it/u=2216102322,3587941637&fm=173&app=25&f=JPEG?w=218&h=146&s=B8C5A14C5FE19F6C14DFED01030070C9",
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1539671813993&di=080422caa5b56f29ae9b3c9dfbda2bf0&imgtype=0&src=http%3A%2F%2Fimg.zcool.cn%2Fcommunity%2F0139ef571adfda32f875a3998b57b7.gif",
		"https://image.winudf.com/v2/image/Y29tLm5jc29mdC5haW9uLmxlZ2lvbnMub2Yud2FyX2Jhbm5lcl8xNTM3NDc5MzcxXzAxMA/banner.jpg?w=850&fakeurl=1&type=.jpg",
		"https://ss0.baidu.com/6ONWsjip0QIZ8tyhnq/it/u=2216102322,3587941637&fm=173&app=25&f=JPEG?w=218&h=146&s=B8C5A14C5FE19F6C14DFED01030070C9",
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1539671813993&di=080422caa5b56f29ae9b3c9dfbda2bf0&imgtype=0&src=http%3A%2F%2Fimg.zcool.cn%2Fcommunity%2F0139ef571adfda32f875a3998b57b7.gif",
		"https://image.winudf.com/v2/image/Y29tLm5jc29mdC5haW9uLmxlZ2lvbnMub2Yud2FyX2Jhbm5lcl8xNTM3NDc5MzcxXzAxMA/banner.jpg?w=850&fakeurl=1&type=.jpg",
		"https://ss0.baidu.com/6ONWsjip0QIZ8tyhnq/it/u=2216102322,3587941637&fm=173&app=25&f=JPEG?w=218&h=146&s=B8C5A14C5FE19F6C14DFED01030070C9",
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1539671813993&di=080422caa5b56f29ae9b3c9dfbda2bf0&imgtype=0&src=http%3A%2F%2Fimg.zcool.cn%2Fcommunity%2F0139ef571adfda32f875a3998b57b7.gif",
		"https://image.winudf.com/v2/image/Y29tLm5jc29mdC5haW9uLmxlZ2lvbnMub2Yud2FyX2Jhbm5lcl8xNTM3NDc5MzcxXzAxMA/banner.jpg?w=850&fakeurl=1&type=.jpg",
		"https://ss0.baidu.com/6ONWsjip0QIZ8tyhnq/it/u=2216102322,3587941637&fm=173&app=25&f=JPEG?w=218&h=146&s=B8C5A14C5FE19F6C14DFED01030070C9",
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1539671813993&di=080422caa5b56f29ae9b3c9dfbda2bf0&imgtype=0&src=http%3A%2F%2Fimg.zcool.cn%2Fcommunity%2F0139ef571adfda32f875a3998b57b7.gif",
		"https://image.winudf.com/v2/image/Y29tLm5jc29mdC5haW9uLmxlZ2lvbnMub2Yud2FyX2Jhbm5lcl8xNTM3NDc5MzcxXzAxMA/banner.jpg?w=850&fakeurl=1&type=.jpg",
		"https://ss0.baidu.com/6ONWsjip0QIZ8tyhnq/it/u=2216102322,3587941637&fm=173&app=25&f=JPEG?w=218&h=146&s=B8C5A14C5FE19F6C14DFED01030070C9",
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1539671813993&di=080422caa5b56f29ae9b3c9dfbda2bf0&imgtype=0&src=http%3A%2F%2Fimg.zcool.cn%2Fcommunity%2F0139ef571adfda32f875a3998b57b7.gif",
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1539671813993&di=080422caa5b56f29ae9b3c9dfbda2bf0&imgtype=0&src=http%3A%2F%2Fimg.zcool.cn%2Fcommunity%2F0139ef571adfda32f875a3998b57b7.gif",
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1539671813993&di=080422caa5b56f29ae9b3c9dfbda2bf0&imgtype=0&src=http%3A%2F%2Fimg.zcool.cn%2Fcommunity%2F0139ef571adfda32f875a3998b57b7.gif",
		"https://image.winudf.com/v2/image/Y29tLm5jc29mdC5haW9uLmxlZ2lvbnMub2Yud2FyX2Jhbm5lcl8xNTM3NDc5MzcxXzAxMA/banner.jpg?w=850&fakeurl=1&type=.jpg",
	}
	for _, v := range Durl {
		req := &protos.UploadFileRequest{
			FileUrl:  v,
			FileType: "image",
			Tags:     []string{"ssssss", "eeeeeeee"},
			Bucket:   "puremedia",
			Prefix:   "pure",
		}
		reqs.UpRequest = append(reqs.UpRequest, req)
	}

	UploadUrlQueue(reqs)
}
