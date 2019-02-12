define({ "api": [
  {
    "type": "post",
    "url": "/api/v1/signature",
    "title": "CreateJSSDKConfig",
    "version": "1.0.0",
    "name": "CreateJSSDKConfig",
    "group": "Index",
    "filename": "../src/controller/index.go",
    "groupTitle": "Index",
    "description": "<p>获取冰岩在线服务号的 JSSDKConfig (POST 版本), 以下接口仍然可用：https://weixin.bingyan-tech.hustonline.net/service/resources/signature</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "url",
            "description": "<p>分享的网站链接</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n  \"url\": String\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "status",
            "defaultValue": "200",
            "description": "<p>状态码</p>"
          },
          {
            "group": "Success 200",
            "type": "Object",
            "optional": false,
            "field": "data",
            "description": "<p>正确返回数据</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 OK\n{\n  \"appId\": String,\n  \"nonce_str\": String,\n  \"signature\": String,\n  \"timestamp\": String,\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "type": "Number",
            "optional": false,
            "field": "status",
            "description": "<p>状态码</p>"
          },
          {
            "group": "Error 4xx",
            "type": "String",
            "optional": false,
            "field": "err_msg",
            "description": "<p>错误信息</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 401 Unauthorized\n{\n  \"status\": 401,\n  \"err_msg\": \"Unauthorized\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "post",
    "url": "/api/v1/qrcode",
    "title": "CreateQrcode",
    "version": "1.0.0",
    "name": "CreateQrcode",
    "group": "Index",
    "filename": "../src/controller/index.go",
    "groupTitle": "Index",
    "description": "<p>获取冰岩在线服务号的 二维码，具体请看微信公众号文档</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "expire_seconds",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "action_name",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "Object",
            "optional": false,
            "field": "action_info",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "action_info.scene_id",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "action_info.scene_str",
            "description": ""
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "     {\n       \"expire_seconds\": Number,\n       \"action_name\": String,\n       \"action_info\": {\n           \"scene\": {\n\t              \"scene_str\": \"test\",\n               \"scene_id\": 123,\n             },\n         },\n     }",
          "type": "json"
        }
      ]
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "status",
            "defaultValue": "200",
            "description": "<p>状态码</p>"
          },
          {
            "group": "Success 200",
            "type": "Object",
            "optional": false,
            "field": "data",
            "description": "<p>正确返回数据</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 OK\n{\n  \"ticket\": String, // 获取的二维码ticket，凭借此ticket可以在有效时间内换取二维码。\n  \"expire_seconds\": Number, // 该二维码有效时间，以秒为单位。 最大不超过2592000（即30天）。\n  \"url\": String, // 二维码图片解析后的地址，开发者可根据该地址自行生成需要的二维码图片\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "type": "Number",
            "optional": false,
            "field": "status",
            "description": "<p>状态码</p>"
          },
          {
            "group": "Error 4xx",
            "type": "String",
            "optional": false,
            "field": "err_msg",
            "description": "<p>错误信息</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 401 Unauthorized\n{\n  \"status\": 401,\n  \"err_msg\": \"Unauthorized\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "get",
    "url": "/api/v1/access-token",
    "title": "GetAccessToken",
    "version": "1.0.0",
    "name": "GetAccessToken",
    "group": "Index",
    "filename": "../src/controller/index.go",
    "groupTitle": "Index",
    "description": "<p>获取冰岩在线服务号的access_token, 以下接口仍然可用：https://weixin.bingyan-tech.hustonline.net/service/resources/AccessToken</p>",
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "status",
            "defaultValue": "200",
            "description": "<p>状态码</p>"
          },
          {
            "group": "Success 200",
            "type": "Object",
            "optional": false,
            "field": "data",
            "description": "<p>正确返回数据</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 OK\n{\n  \"access_token\": String,\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "type": "Number",
            "optional": false,
            "field": "status",
            "description": "<p>状态码</p>"
          },
          {
            "group": "Error 4xx",
            "type": "String",
            "optional": false,
            "field": "err_msg",
            "description": "<p>错误信息</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 401 Unauthorized\n{\n  \"status\": 401,\n  \"err_msg\": \"Unauthorized\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "get",
    "url": "/api/v1/js-api-ticket",
    "title": "GetJSApiTicket",
    "version": "1.0.0",
    "name": "GetJSApiTicket",
    "group": "Index",
    "filename": "../src/controller/index.go",
    "groupTitle": "Index",
    "description": "<p>获取冰岩在线服务号的 jsapi_ticket, 以下接口仍然可用：https://weixin.bingyan-tech.hustonline.net/service/resources/JsApiTicket</p>",
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "status",
            "defaultValue": "200",
            "description": "<p>状态码</p>"
          },
          {
            "group": "Success 200",
            "type": "Object",
            "optional": false,
            "field": "data",
            "description": "<p>正确返回数据</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 OK\n{\n  \"jsapi_ticket\": String,\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "type": "Number",
            "optional": false,
            "field": "status",
            "description": "<p>状态码</p>"
          },
          {
            "group": "Error 4xx",
            "type": "String",
            "optional": false,
            "field": "err_msg",
            "description": "<p>错误信息</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 401 Unauthorized\n{\n  \"status\": 401,\n  \"err_msg\": \"Unauthorized\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "get",
    "url": "/api/v1/signature",
    "title": "GetJSSDKConfig",
    "version": "1.0.0",
    "name": "GetJSSDKConfig",
    "group": "Index",
    "filename": "../src/controller/index.go",
    "groupTitle": "Index",
    "description": "<p>获取冰岩在线服务号的 JSSDKConfig, 以下接口仍然可用：https://weixin.bingyan-tech.hustonline.net/service/resources/signature</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "url",
            "description": "<p>分享的网站链接</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n  \"url\": String\n}",
          "type": "query"
        }
      ]
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "status",
            "defaultValue": "200",
            "description": "<p>状态码</p>"
          },
          {
            "group": "Success 200",
            "type": "Object",
            "optional": false,
            "field": "data",
            "description": "<p>正确返回数据</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 OK\n{\n  \"appId\": String,\n  \"nonce_str\": String,\n  \"signature\": String,\n  \"timestamp\": String,\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "type": "Number",
            "optional": false,
            "field": "status",
            "description": "<p>状态码</p>"
          },
          {
            "group": "Error 4xx",
            "type": "String",
            "optional": false,
            "field": "err_msg",
            "description": "<p>错误信息</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 401 Unauthorized\n{\n  \"status\": 401,\n  \"err_msg\": \"Unauthorized\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "post",
    "url": "/api/v1/msg/template/list/action/send",
    "title": "SendTemplate",
    "version": "1.0.0",
    "name": "SendTemplate",
    "group": "Message",
    "filename": "../src/controller/msg.go",
    "groupTitle": "Message",
    "description": "<p>发送模板消息</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "[]Object",
            "optional": false,
            "field": "templates",
            "description": "<p>模板数组</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "touser",
            "description": "<p>接收者unionid</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "template_id",
            "description": "<p>模板ID</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "url",
            "description": "<p>模板跳转链接（海外帐号没有跳转能力）</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "miniprogram",
            "description": "<p>跳小程序所需数据，不需跳小程序可不用传该数据</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "data",
            "description": "<p>模板数据</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n  \"templates\": [{\n      \"touser\":\"Unionid\",\n      \"template_id\":\"ngqIpbwh8bUfcSsECmogfXcV14J0tQlEpBO27izEYtY\",\n      \"url\":\"http://weixin.qq.com/download\",\n      \"miniprogram\":{\n          \"appid\":\"xiaochengxuappid12345\", // 所需跳转到的小程序appid（该小程序appid必须与发模板消息的公众号是绑定关联关系，暂不支持小游戏）\n          \"pagepath\":\"index?foo=bar\" // 所需跳转到小程序的具体页面路径，支持带参数,（示例index?foo=bar），暂不支持小游戏\n        },\n      \"data\":{\n             \"keyword1\":{\n                 \"value\":\"巧克力\",\n                 \"color\":\"#173177\" // 模板内容字体颜色，不填默认为黑色\n             },\n             \"keyword2\": {\n                 \"value\":\"39.8元\",\n                 \"color\":\"#173177\"\n             },\n             \"keyword3\": {\n                 \"value\":\"2014年9月22日\",\n                 \"color\":\"#173177\"\n             },\n      }\n  }]\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "status",
            "defaultValue": "200",
            "description": "<p>状态码</p>"
          },
          {
            "group": "Success 200",
            "type": "Object",
            "optional": false,
            "field": "data",
            "description": "<p>正确返回数据</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 OK\n{\n  \"status\": 200,\n  \"data\": [{\n      \"errcode\": Number,\n      \"errmsg\": String,\n      \"msgid\": Number,\n    }]\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "type": "Number",
            "optional": false,
            "field": "status",
            "description": "<p>状态码</p>"
          },
          {
            "group": "Error 4xx",
            "type": "String",
            "optional": false,
            "field": "err_msg",
            "description": "<p>错误信息</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 401 Unauthorized\n{\n  \"status\": 401,\n  \"err_msg\": \"Unauthorized\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "get",
    "url": "/api/v1/user/status/follow",
    "title": "GetUserFollowStatus",
    "version": "1.0.0",
    "name": "GetUserFollowStatus",
    "group": "User",
    "filename": "../src/controller/user.go",
    "groupTitle": "User",
    "description": "<p>获取是否关注了服务号</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "unionid",
            "description": "<p>unionid</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n  \"unionid\":\"Unionid\",\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "status",
            "defaultValue": "200",
            "description": "<p>状态码</p>"
          },
          {
            "group": "Success 200",
            "type": "Object",
            "optional": false,
            "field": "data",
            "description": "<p>正确返回数据</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 OK\n{\n   \"is_follow\": Boolean, // 是否关注了服务号\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "type": "Number",
            "optional": false,
            "field": "status",
            "description": "<p>状态码</p>"
          },
          {
            "group": "Error 4xx",
            "type": "String",
            "optional": false,
            "field": "err_msg",
            "description": "<p>错误信息</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 401 Unauthorized\n{\n  \"status\": 401,\n  \"err_msg\": \"Unauthorized\"\n}",
          "type": "json"
        }
      ]
    }
  }
] });
