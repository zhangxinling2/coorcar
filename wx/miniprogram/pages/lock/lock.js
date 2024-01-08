"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const routing_1 = require("../../utils/routing");
// pages/lock/lock.ts
const shareLocationKey = "share_location";
const uploadUrl = "120.25.124.86:8081/avatar";
Page({
    /**
     * 页面的初始数据
     */
    data: {
        shareLocation: false,
        avatarURL: '',
    },
    onChooseAvatar(e) {
        this.setData({
            avatarURL: e.detail.avatarUrl,
        });
        // wx.uploadFile({
        //   filePath:e.detail,
        //   name:"avatarImg",
        //   url:uploadUrl,
        //   success:res=>{
        //     console.log(res)
        //   },
        // })
    },
    onShareSwitch(e) {
        wx.setStorageSync(shareLocationKey, e.detail.value);
    },
    onUnlockTap() {
        wx.getLocation({
            type: 'gcj02',
            success: loc => {
                const tripID = 'trip456';
                //TODO: 开锁，传值头像和经纬度
                wx.showLoading({
                    title: '开锁中',
                    mask: true,
                });
                setTimeout(() => {
                    wx.redirectTo({
                        url: routing_1.routing.driving({
                            trip_id: tripID
                        }),
                        complete: () => wx.hideLoading(),
                    });
                }, 2000);
            },
            fail: () => wx.showToast({
                icon: 'none',
                title: "请前往设置页授权位置信息"
            })
        });
    },
    /**
     * 生命周期函数--监听页面加载
     */
    onLoad(opt) {
        const o = opt;
        console.log('unlocking car', o.car_id);
        this.setData({
            shareLocation: wx.getStorageSync(shareLocationKey) || false
        });
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
