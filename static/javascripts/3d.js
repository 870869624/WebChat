var hijiki="https://unpkg.com/live2d-widget-model-hijiki@1.0.5/assets/hijiki.model.json";
//白猫 
var tororo="https://unpkg.com/live2d-widget-model-tororo@1.0.5/assets/tororo.model.json";
//狗
var wanko="https://unpkg.com/live2d-widget-model-wanko@1.0.5/assets/wanko.model.json";
//人物
var koharu="https://unpkg.com/live2d-widget-model-koharu@1.0.5/assets/koharu.model.json";
var shizuku="https://unpkg.com/live2d-widget-model-shizuku@1.0.5/assets/shizuku.model.json";//默认形象
var z16="https://unpkg.com/live2d-widget-model-z16@1.0.5/assets/z16.model.json";
var chitose="https://unpkg.com/live2d-widget-model-chitose@1.0.5/assets/chitose.model.json";
var haruto="https://unpkg.com/live2d-widget-model-haruto@1.0.5/assets/haruto.model.json";
var hibiki="https://unpkg.com/live2d-widget-model-hibiki@1.0.5/assets/hibiki.model.json";
var izumi="https://unpkg.com/live2d-widget-model-izumi@1.0.5/assets/izumi.model.json";
var miku="https://unpkg.com/live2d-widget-model-miku@1.0.5/assets/miku.model.json";
var nico="https://unpkg.com/live2d-widget-model-nico@1.0.5/assets/nico.model.json";
var nipsilon="https://unpkg.com/live2d-widget-model-nipsilon@1.0.5/assets/nipsilon.model.json";
var nito="https://unpkg.com/live2d-widget-model-nito@1.0.5/assets/nito.model.json";
var tsumiki="https://unpkg.com/live2d-widget-model-tsumiki@1.0.5/assets/tsumiki.model.json";
var unitychan="https://unpkg.com/live2d-widget-model-unitychan@1.0.5/assets/unitychan.model.json";
setTimeout(() => {
    L2Dwidget.init({
        "model": {
            jsonPath:hijiki,//形象
            "scale": 1
        },
        "display": {
            "position": "right",//展示位置
            "width": 150,//看板娘宽
            "height": 225,//看板娘高
            "hOffset": 0,
            "vOffset": -20
        },
        "mobile": {
            "show": true,
            "scale": 0.5
        },
        "react": {
            "opacityDefault": 1,
            "opacityOnHover": 0.2
        }
    });
}, 1000)