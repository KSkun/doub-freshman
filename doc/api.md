完整响应格式：

```json
{
  "data": {
    "next": {
      "id": "",
      "title": "",
      "text": "",
      "option": null,
      "delay": 0
    },
    "selection": [
      {
        "id": "60c6a411074fa0ce2aae7e7f",
        "title": "卷王之王",
        "text": "yyh23已经对各种题型了如指掌，随便什么考试在他眼里不过是浮云，看着学院一个个加权不如他的同学，yyh23冷哼：\"我不是针对谁，我是说，在座的各位在卷这方面都是凡人\" 话毕，羽化登仙，腾空而去",
        "option": [
          "yyh23已然非人，活在了传说中"
        ],
        "delay": 0
      },
      {
        "id": "60c6a432074fa0ce2aae7e83",
        "title": "卷王之王1",
        "text": "yyh23已经对各种题型了如指掌，随便什么考试在他眼里不过是浮云，看着学院一个个加权不如他的同学，yyh23冷哼：\"我不是针对谁，我是说，在座的各位在卷这方面都是凡人\" 话毕，羽化登仙，腾空而去",
        "option": [
          "yyh23已然非人，活在了传说中"
        ],
        "delay": 0
      },
      {
        "id": "60c6a432074fa0ce2aae7e80",
        "title": "上课要迟到了1",
        "text": "一觉醒来，yyh23发现寝室里已经空无一人，一看手表，惊讶地发现上课要迟到了，时不待人，yyh23决定___去上课",
        "option": [
          "步行",
          "骑电瓶车",
          "坐校车"
        ],
        "delay": 0
      }
    ],
    "flag": [
      "加权95.0",
      "假期",
      "萌新",
      "test"
    ],
    "dead": false,
    "end": false,
    "result": "",
    "id": "60c6af14c3fc278347d77f97"
  },
  "message": "",
  "success": true
}
```

接口：

POST /api/v1/new

新建游戏

```json
{
  "name": "玩家名",
  "major": "专业"
}
```

```json
{
  "next": {
    "id": "",
    "title": "",
    "text": "",
    "option": null,
    "delay": 0
  },
  "selection": [
    {
      "id": "60c6a411074fa0ce2aae7e7f",
      "title": "卷王之王",
      "text": "yyh23已经对各种题型了如指掌，随便什么考试在他眼里不过是浮云，看着学院一个个加权不如他的同学，yyh23冷哼：\"我不是针对谁，我是说，在座的各位在卷这方面都是凡人\" 话毕，羽化登仙，腾空而去",
      "option": [
        "yyh23已然非人，活在了传说中"
      ],
      "delay": 0
    },
    {
      "id": "60c6a432074fa0ce2aae7e83",
      "title": "卷王之王1",
      "text": "yyh23已经对各种题型了如指掌，随便什么考试在他眼里不过是浮云，看着学院一个个加权不如他的同学，yyh23冷哼：\"我不是针对谁，我是说，在座的各位在卷这方面都是凡人\" 话毕，羽化登仙，腾空而去",
      "option": [
        "yyh23已然非人，活在了传说中"
      ],
      "delay": 0
    },
    {
      "id": "60c6a432074fa0ce2aae7e80",
      "title": "上课要迟到了1",
      "text": "一觉醒来，yyh23发现寝室里已经空无一人，一看手表，惊讶地发现上课要迟到了，时不待人，yyh23决定___去上课",
      "option": [
        "步行",
        "骑电瓶车",
        "坐校车"
      ],
      "delay": 0
    }
  ],
  "flag": [
    "加权95.0",
    "假期",
    "萌新",
    "test"
  ],
  "dead": false,
  "end": false,
  "result": "",
  "id": "60c6af14c3fc278347d77f97"
}
```

PUT /api/v1/stage

选择剧情点

```json
{
  "player": "60c6af14c3fc278347d77f97",
  "stage": "60c6a432074fa0ce2aae7e80"
}
```

```json
{
  "next": {
    "id": "",
    "title": "",
    "text": "",
    "option": null,
    "delay": 0
  },
  "selection": [
    {
      "id": "60c6a411074fa0ce2aae7e7f",
      "title": "卷王之王",
      "text": "yyh23已经对各种题型了如指掌，随便什么考试在他眼里不过是浮云，看着学院一个个加权不如他的同学，yyh23冷哼：\"我不是针对谁，我是说，在座的各位在卷这方面都是凡人\" 话毕，羽化登仙，腾空而去",
      "option": [
        "yyh23已然非人，活在了传说中"
      ],
      "delay": 0
    },
    {
      "id": "60c6a432074fa0ce2aae7e83",
      "title": "卷王之王1",
      "text": "yyh23已经对各种题型了如指掌，随便什么考试在他眼里不过是浮云，看着学院一个个加权不如他的同学，yyh23冷哼：\"我不是针对谁，我是说，在座的各位在卷这方面都是凡人\" 话毕，羽化登仙，腾空而去",
      "option": [
        "yyh23已然非人，活在了传说中"
      ],
      "delay": 0
    }
  ],
  "flag": [
    "加权95.0",
    "假期",
    "萌新",
    "test"
  ],
  "dead": false,
  "end": false,
  "result": "",
  "flag_diff": null
}
```

PUT /api/v1/option

选择选项

```json
{
  "player": "60c6af14c3fc278347d77f97",
  "option": 1
}
```

```json
{
  "next": {
    "id": "",
    "title": "",
    "text": "",
    "option": null,
    "delay": 0
  },
  "selection": [
    {
      "id": "60c6a411074fa0ce2aae7e7f",
      "title": "卷王之王",
      "text": "yyh23已经对各种题型了如指掌，随便什么考试在他眼里不过是浮云，看着学院一个个加权不如他的同学，yyh23冷哼：\"我不是针对谁，我是说，在座的各位在卷这方面都是凡人\" 话毕，羽化登仙，腾空而去",
      "option": [
        "yyh23已然非人，活在了传说中"
      ],
      "delay": 0
    },
    {
      "id": "60c6a432074fa0ce2aae7e83",
      "title": "卷王之王1",
      "text": "yyh23已经对各种题型了如指掌，随便什么考试在他眼里不过是浮云，看着学院一个个加权不如他的同学，yyh23冷哼：\"我不是针对谁，我是说，在座的各位在卷这方面都是凡人\" 话毕，羽化登仙，腾空而去",
      "option": [
        "yyh23已然非人，活在了传说中"
      ],
      "delay": 0
    },
    {
      "id": "60c6a411074fa0ce2aae7e7c",
      "title": "上课要迟到了",
      "text": "一觉醒来，yyh23发现寝室里已经空无一人，一看手表，惊讶地发现上课要迟到了，时不待人，yyh23决定___去上课",
      "option": [
        "步行",
        "骑电瓶车",
        "坐校车"
      ],
      "delay": 0
    }
  ],
  "flag": [
    "加权95.0",
    "萌新",
    "test",
    "在校"
  ],
  "dead": false,
  "end": false,
  "flag_diff": [
    {
      "type": "del",
      "flag": "假期",
      "value": 0
    },
    {
      "type": "add",
      "flag": "在校",
      "value": 0
    }
  ],
  "result": "yyh23在楼下充电桩找了半天电动车都没找到，终于yyh23意识到自己根本没买电动车，只能步行去教室,很显然地迟到了"
}
```

POST /api/v1/stage

导入剧情点

```json
{
    "stages":
        [
            {
                "_id":"stage_上课要迟到了",
                "title":"上课要迟到了1",
                "text":"一觉醒来，${name}发现寝室里已经空无一人，一看手表，惊讶地发现上课要迟到了，时不待人，${name}决定___去上课",
                "dead":false,
                "enter_cond":[
                    {
                        "flag":"假期",
                        "op":"",
                        "value":0
                    }
                ],
                "option":[
                    {
                        "text":"步行",
                        "success":{
                            "next":"stage_上课迟到",
                            "text":"在去教室的路上，${name}嫉妒不满室友居然没叫自己，退一步越想越气，直到走到教室门口，才发现教室里也是空无一人。${name}一言不发地离开了。"
                        },
                        "failed":{
                            "next":"stage_上课迟到",
                            "text":"${name}步行到教室，意料之中地迟到了"
                        },
                        "condition":[{
                                "flag":"概率标签",
                                "op":"prob",
                                "value":0.1
                            }
                        ]
                    },
                    {
                        "text":"骑电瓶车",
                        "success":{
                            "next":"stage_上课",
                            "text":"${name}骑着电驴风一般地驶向了教室，推门铃响稳如狗，哪怕早八也不愁"
                        },
                        "failed":{
                            "next":"stage_上课迟到",
                            "text":"${name}在楼下充电桩找了半天电动车都没找到，终于${name}意识到自己根本没买电动车，只能步行去教室,很显然地迟到了"
                        },
                        "condition":[{
                                "flag":"电瓶车",
                                "op":"",
                                "value":0
                            }
                        ]
                    },
                    {
                        "text":"坐校车",
                        "success":{
                            "next":"stage_上课要迟到了_坐校车",
                            "text":"${name}决定坐校车去教室"
                        }
                    }
                ],
                "tag":"上课要迟到了",

                "continue":false,
                "delay":0
            },
            {
                "_id":"stage_上课要迟到了_坐校车",
                "title":"上课要迟到了1",
                "text":"在坐校车去教室的过程中,${name}",
                "dead":false,
                "option":[
                    {
                        "text":"拿出手机刷朋友圈",
                        "success":{
                            "next":"stage_上课要迟到了_坐校车_刷手机",
                            "text":"在坐校车去教室的过程中,${name}拿出手机刷朋友圈"
                        }
                    },
                    {
                        "text":"眯上眼睛再睡会儿",
                        "success":{
                            "next":"stage_缺课挂科",
                            "text":"校车带着在学校里逛了一早上，司机师傅下班才发现叫醒${name}，${name}很不幸地被老师点到名，愤怒的老师决定挂${name}的科"
                        }
                    },
                    {
                        "text":"抖音外放",
                        "success":{
                            "next":"stage_上课",
                            "text":"土味的BGM瞬间充斥着整辆校车，${name}感觉如芒在背，周围一道道锐利的目光刺向${name}，在不安中校车终于到了目的地，${name}下车去上课了"
                        }
                    }
                ],
                "tag":"上课要迟到了",
                "continue":true,
                "delay":0
            },
            {
                "_id":"stage_上课要迟到了_坐校车_刷手机",
                "title":"上课要迟到了1",
                "text":"正当${name}津津有味地刷着手机时，余光偏见一道强光透过窗户袭来，一股凉意从脚底直冒头顶，${name}决定：",
                "dead":false,
                "option":[
                    {
                        "text":"立即跳车",
                        "success":{
                            "next":"stage_上课迟到",
                            "text":"${name}身手矫健的翻出了窗户，一个前滚翻稳当地停在了路面上。所有路上的行人都停了下来，目瞪口呆地看着他，有人拿出了手机开始拍照。校车很快消失在了视线里，${name}不得不步行上课"
                        },
                        "failed":{
                            "next":"",
                            "text":"${name}身体探出窗户，突然重心不稳，摔下了校车，强大的惯性下他重重地摔在地上，再也没起来",
                            "event":[
                                {
                                    "type":"death"
                                }
                            ]
                        },
                        "condition":[
                            {
                                "flag":"体格强魄",
                                "op":"",
                                "value":0
                            }
                        ]
                    },
                    {
                        "text":"接着看手机，手机真好看",
                        "success":{
                            "next":"stage_上课",
                            "text":"虽然手机很好看，但校车很快到了目的地，${name}恋恋不舍地收起手机，只能去上课"
                        },
                        "failed":{
                            "next":"",
                            "text":"pong的一声巨响，${name}还没反应过来发生了什么，身子就飞出校车，然后世界一片空白",
                            "event":[
                                {
                                    "type":"death"
                                }
                            ]
                        },
                        "condition":[
                            {
                                "flag":"概率标签",
                                "op":"prob",
                                "value":0.5
                            }
                        ]
                    },
                    {
                        "text":"立即拨打原作者电话，问诡异的光怎么整",
                        "success":{
                            "next":"stage_上课",
                            "text":"“大家别再催我了，标准答案没出来，我怎么知道我想表达什么？” ${name}不明所以，校车到站去上课了"
                        }
                    }
                ],
                "tag":"上课要迟到了",

                "continue":true,
                "delay":0
            },
            
            {
                "_id":"stage_卷王之王",
                "title":"卷王之王1",
                "text":"${name}已经对各种题型了如指掌，随便什么考试在他眼里不过是浮云，看着学院一个个加权不如他的同学，${name}冷哼：\"我不是针对谁，我是说，在座的各位在卷这方面都是凡人\" 话毕，羽化登仙，腾空而去",
                "dead":false,
                "enter_cond":[
                    {
                        "flag":"加权",
                        "op":"gte",
                        "value":95
                    }
                ],
                "option":[
                    {
                        "text":"${name}已然非人，活在了传说中",
                        "success":{
                            "next":"",
                            "text":"${name}已然非人，活在了传说中",
                            "event":[
                                {
                                    "type":"death"
                                }
                            ]
                        }
                    }
                ],
                "tag":"卷王之王",
                "continue":false,
                "delay":0
            }
        ]
    }
```

```json
null
```
