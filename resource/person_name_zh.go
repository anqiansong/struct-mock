package resource

import (
	"strings"
	"time"
)

var (
	personNameZh = []string{
		"赵", "钱", "孙", "李", "周", "吴", "郑", "王", "冯", "陈", "褚", "卫", "蒋", "沈", "韩", "杨", "朱", "秦", "尤", "许", "何", "吕", "施", "张",
		"孔", "曹", "严", "华", "金", "魏", "陶", "姜", "戚", "谢", "邹", "喻", "柏", "水", "窦", "章", "云", "苏", "潘", "葛", "奚", "范", "彭", "郎",
		"鲁", "韦", "昌", "马", "苗", "凤", "花", "方", "俞", "任", "袁", "柳", "酆", "鲍", "史", "唐", "费", "廉", "岑", "薛", "雷", "贺", "倪", "汤",
		"滕", "殷", "罗", "毕", "郝", "邬", "安", "常", "乐", "于", "时", "傅", "皮", "卞", "齐", "康", "伍", "余", "元", "卜", "顾", "孟", "平", "黄",
		"和", "穆", "萧", "尹", "姚", "邵", "湛", "汪", "祁", "毛", "禹", "狄", "米", "贝", "明", "臧", "计", "伏", "成", "戴", "谈", "宋", "茅", "庞",
		"熊", "纪", "舒", "屈", "项", "祝", "董", "梁", "杜", "阮", "蓝", "闵", "席", "季", "麻", "强", "贾", "路", "娄", "危", "江", "童", "颜", "郭",
		"梅", "盛", "林", "刁", "钟", "徐", "邱", "骆", "高", "夏", "蔡", "田", "樊", "胡", "凌", "霍", "虞", "万", "支", "柯", "昝", "管", "卢", "莫",
		"经", "房", "裘", "缪", "干", "解", "应", "宗", "丁", "宣", "贲", "邓", "郁", "单", "杭", "洪", "包", "诸", "左", "石", "崔", "吉", "钮", "龚",
		"程", "嵇", "邢", "滑", "裴", "陆", "荣", "翁", "荀", "羊", "於", "惠", "甄", "曲", "家", "封", "芮", "羿", "储", "靳", "汲", "邴", "糜", "松",
		"井", "段", "富", "巫", "乌", "焦", "巴", "弓", "牧", "隗", "山", "谷", "车", "侯", "宓", "蓬", "全", "郗", "班", "仰", "秋", "仲", "伊", "宫",
		"宁", "仇", "栾", "暴", "甘", "钭", "厉", "戎", "祖", "武", "符", "刘", "景", "詹", "束", "龙", "叶", "幸", "司", "韶", "郜", "黎", "蓟", "薄",
		"印", "宿", "白", "怀", "蒲", "邰", "从", "鄂", "索", "咸", "籍", "赖", "卓", "蔺", "屠", "蒙", "池", "乔", "阴", "鬱", "胥", "能", "苍", "双",
		"闻", "莘", "党", "翟", "谭", "贡", "劳", "逄", "姬", "申", "扶", "堵", "冉", "宰", "郦", "雍", "卻", "璩", "桑", "桂", "濮", "牛", "寿", "通",
		"边", "扈", "燕", "冀", "郏", "浦", "尚", "农", "温", "别", "庄", "晏", "柴", "瞿", "阎", "充", "慕", "连", "茹", "习", "宦", "艾", "鱼", "容",
		"向", "古", "易", "慎", "戈", "廖", "庾", "终", "暨", "居", "衡", "步", "都", "耿", "满", "弘", "匡", "国", "文", "寇", "广", "禄", "阙", "东",
		"欧", "殳", "沃", "利", "蔚", "越", "夔", "隆", "师", "巩", "厍", "聂", "晁", "勾", "敖", "融", "冷", "訾", "辛", "阚", "那", "简", "饶", "空",
		"曾", "毋", "沙", "乜", "养", "鞠", "须", "丰", "巢", "关", "蒯", "相", "查", "后", "荆", "红", "游", "竺", "权", "逯", "盖", "益", "桓", "公",
		"万", "俟", "司", "马", "上", "官", "欧", "阳", "夏", "侯", "诸", "葛", "闻", "人", "东", "方", "赫", "连", "皇", "甫", "尉", "迟", "公", "羊",
		"澹", "台", "公", "冶", "宗", "政", "濮", "阳", "淳", "于", "单", "于", "太", "叔", "申", "屠", "公", "孙", "仲", "孙", "轩", "辕", "令", "狐",
		"钟", "离", "宇", "文", "长", "孙", "慕", "容", "鲜", "于", "闾", "丘", "司", "徒", "司", "空", "丌", "官", "司", "寇", "仉", "督", "子", "车",
		"颛", "孙", "端", "木", "巫", "马", "公", "西", "漆", "雕", "乐", "正", "壤", "驷", "公", "良", "拓", "跋", "夹", "谷", "宰", "父", "谷", "梁",
		"晋", "楚", "闫", "法", "汝", "鄢", "涂", "钦", "段", "干", "百", "里", "东", "郭", "南", "门", "呼", "延", "归", "海", "羊", "舌", "微", "生",
		"岳", "帅", "缑", "亢", "况", "郈", "有", "琴", "梁", "丘", "左", "丘", "东", "门", "西", "门", "商", "牟", "佘", "佴", "伯", "赏", "南", "宫",
		"墨", "哈", "谯", "笪", "年", "爱", "阳", "佟", "第", "五", "言", "福",
	}
	personLastNameZh = []string{
		"宪", "晔", "熙", "靓", "阳", "乐", "时", "万", "亮", "胜", "姣", "淳", "蓝", "光", "璧", "桓", "虎", "旻", "庭", "红", "宗", "嘉", "福", "渊",
		"浩", "方", "贞", "凤", "冠", "勋", "满", "蒙", "玟", "亭", "毓", "恒", "真", "晴", "泉", "新", "碧", "羚", "晋", "勇", "序", "维", "桦", "纯",
		"齐", "爽", "泽", "润", "奎", "营", "展", "洁", "毅", "元", "其", "黎", "申", "锐", "优", "钢", "瑾", "卓", "喜", "滨", "佚", "菱", "家", "贤",
		"昊", "飞", "升", "金", "石", "孝", "俪", "鸿", "尧", "奇", "菡", "超", "诚", "涛", "洺", "先", "逸", "议", "朗", "沛", "业", "善", "谊", "静",
		"春", "可", "衡", "略", "溢", "琬", "进", "姗", "沣", "充", "意", "登", "芷", "雄", "泰", "年", "聪", "淑", "文", "培", "夏", "琪", "秀", "震",
		"儿", "丰", "莹", "歌", "舜", "朦", "平", "琼", "苑", "远", "洲", "世", "南", "凰", "银", "西", "朔", "才", "妍", "玫", "良", "跃", "权", "朝",
		"标", "河", "融", "珍", "淼", "婧", "谚", "安", "亿", "予", "凡", "娇", "仪", "坚", "基", "颖", "珠", "映", "乾", "钗", "炎", "乃", "珊", "宏",
		"广", "菏", "松", "娜", "策", "星", "湾", "娥", "诗", "吉", "会", "小", "贝", "夕", "兰", "晗", "科", "畅", "韶", "聆", "沅", "敬", "杨", "水",
		"凯", "翔", "继", "冉", "力", "鹏", "宇", "朋", "京", "同", "锡", "至", "紫", "瑶", "保", "祺", "心", "韵", "晨", "育", "纪", "怡", "馥", "美",
		"威", "强", "玉", "丹", "玮", "选", "备", "睻", "生", "和", "革", "沫", "崊", "竹", "苓", "磊", "丫", "雯", "婷", "士", "谨", "瑞", "卫", "岳",
		"含", "瀚", "蕾", "发", "荷", "霏", "海", "化", "迎", "勤", "民", "铁", "连", "巩", "婵", "耘", "宸", "骅", "佑", "君", "天", "霖", "溓", "沁",
		"初", "之", "传", "礼", "曦", "向", "霄", "有", "正", "妹", "康", "少", "田", "伶", "旭", "柳", "山", "来", "骁", "怿", "圣", "源", "云", "炳",
		"根", "茜", "清", "惠", "双", "娴", "容", "梅", "壮", "征", "靖", "津", "芬", "武", "致", "怀", "儒", "照", "昌", "菁", "锦", "影", "起", "秋",
		"娣", "永", "启", "昭", "颜", "雪", "信", "玲", "佩", "显", "巍", "巧", "琳", "倩", "妙", "军", "誉", "锋", "渝", "书", "柏", "镇", "定", "璟",
		"苹", "筠", "茹", "义", "漫", "庆", "夫", "增", "迅", "婉", "莺", "成", "敖", "棠", "环", "理", "羽", "帆", "识", "英", "以", "莲", "芝", "森",
		"潮", "常", "婕", "功", "洋", "雷", "晟", "艺", "孜", "嫣", "辰", "章", "群", "欣", "汉", "垒", "柔", "达", "利", "钰", "瑷", "骄", "荣", "斌",
		"娅", "雅", "开", "百", "泓", "道", "劲", "弘", "璇", "纶", "枝", "喆", "佳", "爱", "渺", "彪", "萍", "恩", "菲", "川", "孟", "树", "铭", "鲁",
		"赏", "贵", "行", "雨", "越", "品", "易", "蕊", "咏", "钦", "灵", "纲", "言", "隆", "益", "辉", "冬", "逦", "宜", "轩", "波", "蓉", "榕", "兵",
		"珮", "北", "栋", "克", "花", "蓓", "筱", "林", "鸣", "奕", "曼", "汝", "尚", "娆", "琴", "虹", "茵", "允", "龙", "亨", "敏", "耀", "然", "灿",
		"城", "风", "滢", "全", "欢", "珑", "蔓", "翰", "姬", "丞", "一", "日", "承", "智", "哲", "潞", "景", "焕", "盈", "凝", "宝", "延", "龄", "叶",
		"莉", "纨", "青", "露", "慧", "任", "忻", "祖", "桐", "珂", "廷", "捷", "聚", "好", "峰", "悦", "明", "牧", "多", "潇", "富", "蝶", "男", "嫒",
		"微", "汶", "溪", "偌", "皓", "励", "泊", "枫", "丽", "凌", "麟", "滔", "禾", "岚", "桂", "艳", "霭", "萌", "玹", "钧", "鑫", "献", "澜", "卿",
		"伦", "友", "久", "函", "黛", "邦", "驰", "甜", "杰", "宣", "豪", "月", "宁", "蔚", "仑", "飘", "济", "苒", "兆", "高", "臻", "楠", "骏", "韬",
		"帅", "卉", "涓", "村", "乔", "思", "志", "彤", "姿", "迪", "沂", "园", "翠", "珺", "仲", "腾", "彰", "谱", "焘", "芹", "绮", "立", "赫", "轮",
		"政", "芸", "瑗", "顺", "刚", "伊", "祥", "崇", "瑜", "运", "鹤", "仁", "涵", "洪", "语", "恬", "扬", "钊", "香", "棋", "馨", "忆", "丛", "燕",
		"睿", "材", "晶", "茂", "忠", "江", "雁", "坤", "娟", "芳", "薇", "晖", "为", "东", "国", "昕", "诒", "岩", "寒", "眉", "振", "彩", "琛", "彬",
		"琦", "与", "望", "昱", "梦", "妮", "堂", "渲", "加", "谋", "伟", "伯", "剑", "寿", "谦", "建", "臣", "希", "麒", "子", "烟", "琰", "萱", "博",
		"采", "霜", "声", "霞", "莎", "州", "大", "楚", "如", "晓", "劭", "璐", "航", "霆", "罡", "亚", "盛", "德", "淇", "若", "长", "炜", "茗", "亦",
		"舒", "俊", "笑", "琸", "华", "学", "裕", "媛", "素", "韦", "非", "名", "梁", "中", "彦", "湘", "朵", "愉", "珏", "健", "妤", "固", "梓", "菊",
		"绍", "依", "冰", "兴", "厚",
	}
)

func RandPersonNameZh() string {
	nameLength := int64(len(personNameZh))
	lastNameLength := int64(len(personLastNameZh))
	if nameLength%2 == 0 {
		nameLength -= 1
	}
	if lastNameLength%2 == 0 {
		lastNameLength -= 1
	}
	unixNano := time.Now().UnixNano()
	builder := new(strings.Builder)
	builder.WriteString(personNameZh[unixNano%nameLength])
	builder.WriteString(personLastNameZh[unixNano%lastNameLength])
	if unixNano%2 == 0 {
		unixNano -= 1
		builder.WriteString(personLastNameZh[unixNano%lastNameLength])
	}
	return builder.String()
}