"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
// app.ts
App({
    globalData: {},
    onLaunch() {
        wx.request({
            url: "http://localhost:8080/trip/trip123",
            method: "GET",
            success: console.log,
            fail: console.error,
        });
        // 展示本地存储能力
        const logs = wx.getStorageSync('logs') || [];
        logs.unshift(Date.now());
        wx.setStorageSync('logs', logs);
        // 登录
        wx.login({
            success: res => {
                console.log(res.code);
                // 发送 res.code 到后台换取 openId, sessionKey, unionId
            },
        });
    },
});
