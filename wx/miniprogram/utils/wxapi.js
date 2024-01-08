"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.getUserProfile = exports.getUserInfo = exports.getSetting = exports.formatTime = void 0;
exports.formatTime = (date) => {
    const year = date.getFullYear();
    const month = date.getMonth() + 1;
    const day = date.getDate();
    const hour = date.getHours();
    const minute = date.getMinutes();
    const second = date.getSeconds();
    return ([year, month, day].map(formatNumber).join('/') +
        ' ' +
        [hour, minute, second].map(formatNumber).join(':'));
};
const formatNumber = (n) => {
    const s = n.toString();
    return s[1] ? s : '0' + s;
};
function getSetting() {
    return new Promise((resolve, reject) => {
        wx.getSetting({
            //success:res=>resolve(res),
            //res=>resolve(res)这是个函数，resolve本身也是个函数，res=>resolve(res)意思与直接resolve相同，reject同理
            success: resolve,
            fail: reject,
        });
    });
}
exports.getSetting = getSetting;
function getUserInfo() {
    return new Promise((resolve, reject) => {
        wx.getUserInfo({
            success: resolve,
            fail: reject,
        });
    });
}
exports.getUserInfo = getUserInfo;
function getUserProfile() {
    return new Promise((resolve, reject) => {
        wx.getUserProfile({
            desc: '展示用户信息',
            success: resolve,
            fail: reject,
        });
    });
}
exports.getUserProfile = getUserProfile;
