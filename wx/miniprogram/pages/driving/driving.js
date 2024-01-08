"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
// pages/driving/driving.ts
const centPerSec = 0.7;
Page({
    /**
     * 页面的初始数据
     */
    data: {
        elapsed: '00:00:00',
        fee: '00',
        location: {
            latitude: 31,
            longitude: 120,
        },
        scal: 12,
    },
    setupLocationUpdate() {
        wx.startLocationUpdate({
            fail: console.error,
        });
        wx.onLocationChange(loc => {
            this.setData({
                location: {
                    latitude: loc.latitude,
                    longitude: loc.longitude
                }
            });
        });
    },
    formatDuration(sec) {
        const padString = (n) => {
            return n < 10 ? '0' + n.toFixed(0) : n.toFixed(0);
        };
        const h = Math.floor(sec / 3600);
        sec = sec - 3600 * h;
        const m = Math.floor(sec / 60);
        sec = sec - m * 60;
        const s = Math.floor(sec);
        return `${padString(h)}:${padString(m)}:${padString(s)}`;
    },
    formatFee(cents) {
        return (cents / 100).toFixed(2);
    },
    setupTimer() {
        let elapsedSec = 0;
        let cent = 0;
        setInterval(() => {
            elapsedSec++;
            cent++;
            this.setData({
                elapsed: this.formatDuration(elapsedSec),
                fee: this.formatFee(cent)
            });
        }, 1000);
    },
    /**
     * 生命周期函数--监听页面加载
     */
    onLoad(opt) {
        const o = opt;
        console.log('trip', o.trip_id);
        this.setupLocationUpdate();
        this.setupTimer();
    },
    /**
     * 生命周期函数--监听页面初次渲染完成
     */
    onReady() {
    },
    /**
     * 生命周期函数--监听页面显示
     */
    onShow() {
    },
    /**
     * 生命周期函数--监听页面隐藏
     */
    onHide() {
    },
    /**
     * 生命周期函数--监听页面卸载
     */
    onUnload() {
        wx.stopLocationUpdate();
    },
    /**
     * 页面相关事件处理函数--监听用户下拉动作
     */
    onPullDownRefresh() {
    },
    /**
     * 页面上拉触底事件的处理函数
     */
    onReachBottom() {
    },
    /**
     * 用户点击右上角分享
     */
    onShareAppMessage() {
    }
});
