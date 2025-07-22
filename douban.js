//https://movie.douban.com/j/search_tags?type=movie
let defaultMovieTags = [
  { name: "热门", raw: "热门" },
  { name: "最新", raw: "最新" },
  { name: "经典", raw: "经典" },
  { name: "豆瓣高分", raw: "豆瓣高分" },
  { name: "冷门佳片", raw: "冷门佳片" },
  { name: "华语", raw: "华语" },
  { name: "欧美", raw: "欧美" },
  { name: "韩国", raw: "韩国" },
  { name: "日本", raw: "日本" },
  { name: "动作", raw: "动作" },
  { name: "喜剧", raw: "喜剧" },
  { name: "爱情", raw: "爱情" },
  { name: "科幻", raw: "科幻" },
  { name: "悬疑", raw: "悬疑" },
  { name: "恐怖", raw: "恐怖" },
  { name: "成长", raw: "成长" },
];

//https://movie.douban.com/j/search_tags?type=tv
let defaultTvTags = [
  { name: "热门", raw: "热门" },
  { name: "美剧", raw: "美剧" },
  { name: "英剧", raw: "英剧" },
  { name: "韩剧", raw: "韩剧" },
  { name: "日剧", raw: "日剧" },
  { name: "国产剧", raw: "国产剧" },
  { name: "港剧", raw: "港剧" },
  { name: "日本动画", raw: "日本动画" },
  { name: "综艺", raw: "综艺" },
  { name: "纪录片", raw: "纪录片" },
];

/*
{
	"name": "豆瓣",
	"raw": "douban",
	"class": [{
		"name": "电影",
		"raw": "movie",
		"tags": [{
			"name": "热门",
			"raw": "热门"
		}]
	}]
}
*/
function Home() {
  const data = {
    name: "豆瓣",
    raw: "douban",
    class: [
      {
        name: "电影",
        raw: "movie",
        tags: defaultMovieTags,
      },
      {
        name: "电视剧",
        raw: "tv",
        tags: defaultTvTags,
      },
    ],
  };

  return JSON.stringify(data);
}

//mtype 类别 movie, tv
//tag 标签
//pageStart 开始页
//pageCount 一页的条数
/*
{
	"subjects": [{
		"id": "36217796",
		"rate": "6.6",
		"title": "梅根2.0",
		"cover": "封面"
	}]
}
*/
function Data(mtype, tag, pageStart, pageCount) {
  const url = `https://movie.douban.com/j/search_subjects?type=${mtype}&tag=${tag}&sort=recommend&page_limit=${pageCount}&page_start=${pageStart}`;
  const headers = {
    Referer: "https://movie.douban.com/",
    Accept: "application/json, text/plain, */*",
  };

  return Get(url, headers);
}

module.exports = {
  Home: Home,
  Data: Data,
};
