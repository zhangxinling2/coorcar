"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const routing_1 = require("../../utils/routing");
const shareLocationKey = "share_location";
const uploadUrl = "120.25.124.86:8081/avatar";
Page({
    data: {
        shareLocation: false,
        avatarURL: '',
    },
    onChooseAvatar(e) {
        this.setData({
            avatarURL: e.detail.avatarUrl,
        });
    },
    onShareSwitch(e) {
        wx.setStorageSync(shareLocationKey, e.detail.value);
    },
    onUnlockTap() {
        wx.getLocation({
            type: 'gcj02',
            success: loc => {
                const tripID = 'trip456';
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
    onLoad(opt) {
        const o = opt;
        console.log('unlocking car', o.car_id);
        this.setData({
            shareLocation: wx.getStorageSync(shareLocationKey) || false
        });
    },
    onReady() {
    },
    onShow() {
    },
    onHide() {
    },
    onUnload() {
    },
    onPullDownRefresh() {
    },
    onReachBottom() {
    },
    onShareAppMessage() {
    }
});
//# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoibG9jay5qcyIsInNvdXJjZVJvb3QiOiIiLCJzb3VyY2VzIjpbImxvY2sudHMiXSwibmFtZXMiOltdLCJtYXBwaW5ncyI6Ijs7QUFBQSxpREFBNkM7QUFHN0MsTUFBTSxnQkFBZ0IsR0FBQyxnQkFBZ0IsQ0FBQTtBQUN2QyxNQUFNLFNBQVMsR0FBQywyQkFBMkIsQ0FBQTtBQUMzQyxJQUFJLENBQUM7SUFLSCxJQUFJLEVBQUU7UUFDSixhQUFhLEVBQUMsS0FBSztRQUNuQixTQUFTLEVBQUMsRUFBRTtLQUNiO0lBQ0QsY0FBYyxDQUFDLENBQUs7UUFDbEIsSUFBSSxDQUFDLE9BQU8sQ0FBQztZQUNYLFNBQVMsRUFBQyxDQUFDLENBQUMsTUFBTSxDQUFDLFNBQVM7U0FDN0IsQ0FBQyxDQUFBO0lBU0osQ0FBQztJQUNELGFBQWEsQ0FBQyxDQUFLO1FBQ2pCLEVBQUUsQ0FBQyxjQUFjLENBQUMsZ0JBQWdCLEVBQUMsQ0FBQyxDQUFDLE1BQU0sQ0FBQyxLQUFLLENBQUMsQ0FBQTtJQUNwRCxDQUFDO0lBQ0QsV0FBVztRQUNULEVBQUUsQ0FBQyxXQUFXLENBQUM7WUFDYixJQUFJLEVBQUMsT0FBTztZQUNaLE9BQU8sRUFBQyxHQUFHLENBQUEsRUFBRTtnQkFDWCxNQUFNLE1BQU0sR0FBQyxTQUFTLENBQUE7Z0JBRXRCLEVBQUUsQ0FBQyxXQUFXLENBQUM7b0JBQ2IsS0FBSyxFQUFDLEtBQUs7b0JBQ1gsSUFBSSxFQUFDLElBQUk7aUJBQ1YsQ0FBQyxDQUFBO2dCQUNGLFVBQVUsQ0FBQyxHQUFHLEVBQUU7b0JBRWQsRUFBRSxDQUFDLFVBQVUsQ0FBQzt3QkFDWixHQUFHLEVBQUMsaUJBQU8sQ0FBQyxPQUFPLENBQUM7NEJBQ2xCLE9BQU8sRUFBRSxNQUFNO3lCQUNoQixDQUFDO3dCQUNGLFFBQVEsRUFBQyxHQUFFLEVBQUUsQ0FBQSxFQUFFLENBQUMsV0FBVyxFQUFFO3FCQUM5QixDQUFDLENBQUE7Z0JBQ0osQ0FBQyxFQUFDLElBQUksQ0FBQyxDQUFDO1lBQ1YsQ0FBQztZQUNELElBQUksRUFBQyxHQUFFLEVBQUUsQ0FBQSxFQUFFLENBQUMsU0FBUyxDQUFDO2dCQUNwQixJQUFJLEVBQUMsTUFBTTtnQkFDWCxLQUFLLEVBQUMsY0FBYzthQUNyQixDQUFDO1NBQ0gsQ0FBQyxDQUFBO0lBRUosQ0FBQztJQUtELE1BQU0sQ0FBQyxHQUEyQjtRQUNoQyxNQUFNLENBQUMsR0FBa0IsR0FBRyxDQUFBO1FBQzVCLE9BQU8sQ0FBQyxHQUFHLENBQUMsZUFBZSxFQUFDLENBQUMsQ0FBQyxNQUFNLENBQUMsQ0FBQTtRQUNyQyxJQUFJLENBQUMsT0FBTyxDQUFDO1lBQ1gsYUFBYSxFQUFDLEVBQUUsQ0FBQyxjQUFjLENBQUMsZ0JBQWdCLENBQUMsSUFBRSxLQUFLO1NBQ3pELENBQUMsQ0FBQTtJQUNKLENBQUM7SUFLRCxPQUFPO0lBRVAsQ0FBQztJQUtELE1BQU07SUFFTixDQUFDO0lBS0QsTUFBTTtJQUVOLENBQUM7SUFLRCxRQUFRO0lBRVIsQ0FBQztJQUtELGlCQUFpQjtJQUVqQixDQUFDO0lBS0QsYUFBYTtJQUViLENBQUM7SUFLRCxpQkFBaUI7SUFFakIsQ0FBQztDQUNGLENBQUMsQ0FBQSIsInNvdXJjZXNDb250ZW50IjpbImltcG9ydCB7IHJvdXRpbmcgfSBmcm9tIFwiLi4vLi4vdXRpbHMvcm91dGluZ1wiXHJcblxyXG4vLyBwYWdlcy9sb2NrL2xvY2sudHNcclxuY29uc3Qgc2hhcmVMb2NhdGlvbktleT1cInNoYXJlX2xvY2F0aW9uXCJcclxuY29uc3QgdXBsb2FkVXJsPVwiMTIwLjI1LjEyNC44Njo4MDgxL2F2YXRhclwiXHJcblBhZ2Uoe1xyXG5cclxuICAvKipcclxuICAgKiDpobXpnaLnmoTliJ3lp4vmlbDmja5cclxuICAgKi9cclxuICBkYXRhOiB7XHJcbiAgICBzaGFyZUxvY2F0aW9uOmZhbHNlLFxyXG4gICAgYXZhdGFyVVJMOicnLFxyXG4gIH0sXHJcbiAgb25DaG9vc2VBdmF0YXIoZTphbnkpIHtcclxuICAgIHRoaXMuc2V0RGF0YSh7XHJcbiAgICAgIGF2YXRhclVSTDplLmRldGFpbC5hdmF0YXJVcmwgLFxyXG4gICAgfSlcclxuICAgIC8vIHd4LnVwbG9hZEZpbGUoe1xyXG4gICAgLy8gICBmaWxlUGF0aDplLmRldGFpbCxcclxuICAgIC8vICAgbmFtZTpcImF2YXRhckltZ1wiLFxyXG4gICAgLy8gICB1cmw6dXBsb2FkVXJsLFxyXG4gICAgLy8gICBzdWNjZXNzOnJlcz0+e1xyXG4gICAgLy8gICAgIGNvbnNvbGUubG9nKHJlcylcclxuICAgIC8vICAgfSxcclxuICAgIC8vIH0pXHJcbiAgfSxcclxuICBvblNoYXJlU3dpdGNoKGU6YW55KXtcclxuICAgIHd4LnNldFN0b3JhZ2VTeW5jKHNoYXJlTG9jYXRpb25LZXksZS5kZXRhaWwudmFsdWUpXHJcbiAgfSxcclxuICBvblVubG9ja1RhcCgpe1xyXG4gICAgd3guZ2V0TG9jYXRpb24oe1xyXG4gICAgICB0eXBlOidnY2owMicsXHJcbiAgICAgIHN1Y2Nlc3M6bG9jPT57XHJcbiAgICAgICAgY29uc3QgdHJpcElEPSd0cmlwNDU2J1xyXG4gICAgICAgIC8vVE9ETzog5byA6ZSB77yM5Lyg5YC85aS05YOP5ZKM57uP57qs5bqmXHJcbiAgICAgICAgd3guc2hvd0xvYWRpbmcoe1xyXG4gICAgICAgICAgdGl0bGU6J+W8gOmUgeS4rScsXHJcbiAgICAgICAgICBtYXNrOnRydWUsXHJcbiAgICAgICAgfSlcclxuICAgICAgICBzZXRUaW1lb3V0KCgpID0+IHtcclxuICAgICAgICAgIFxyXG4gICAgICAgICAgd3gucmVkaXJlY3RUbyh7XHJcbiAgICAgICAgICAgIHVybDpyb3V0aW5nLmRyaXZpbmcoe1xyXG4gICAgICAgICAgICAgIHRyaXBfaWQ6IHRyaXBJRFxyXG4gICAgICAgICAgICB9KSxcclxuICAgICAgICAgICAgY29tcGxldGU6KCk9Pnd4LmhpZGVMb2FkaW5nKCksXHJcbiAgICAgICAgICB9KVxyXG4gICAgICAgIH0sMjAwMCk7XHJcbiAgICAgIH0sXHJcbiAgICAgIGZhaWw6KCk9Pnd4LnNob3dUb2FzdCh7XHJcbiAgICAgICAgaWNvbjonbm9uZScsXHJcbiAgICAgICAgdGl0bGU6XCLor7fliY3lvoDorr7nva7pobXmjojmnYPkvY3nva7kv6Hmga9cIlxyXG4gICAgICB9KVxyXG4gICAgfSlcclxuICAgIFxyXG4gIH0sXHJcblxyXG4gIC8qKlxyXG4gICAqIOeUn+WRveWRqOacn+WHveaVsC0t55uR5ZCs6aG16Z2i5Yqg6L29XHJcbiAgICovXHJcbiAgb25Mb2FkKG9wdDpSZWNvcmQ8J2Nhcl9pZCcsc3RyaW5nPikge1xyXG4gICAgY29uc3Qgbzpyb3V0aW5nLkxvY2tPcHRzPW9wdFxyXG4gICAgY29uc29sZS5sb2coJ3VubG9ja2luZyBjYXInLG8uY2FyX2lkKVxyXG4gICAgdGhpcy5zZXREYXRhKHtcclxuICAgICAgc2hhcmVMb2NhdGlvbjp3eC5nZXRTdG9yYWdlU3luYyhzaGFyZUxvY2F0aW9uS2V5KXx8ZmFsc2VcclxuICAgIH0pXHJcbiAgfSxcclxuXHJcbiAgLyoqXHJcbiAgICog55Sf5ZG95ZGo5pyf5Ye95pWwLS3nm5HlkKzpobXpnaLliJ3mrKHmuLLmn5PlrozmiJBcclxuICAgKi9cclxuICBvblJlYWR5KCkge1xyXG5cclxuICB9LFxyXG5cclxuICAvKipcclxuICAgKiDnlJ/lkb3lkajmnJ/lh73mlbAtLeebkeWQrOmhtemdouaYvuekulxyXG4gICAqL1xyXG4gIG9uU2hvdygpIHtcclxuXHJcbiAgfSxcclxuXHJcbiAgLyoqXHJcbiAgICog55Sf5ZG95ZGo5pyf5Ye95pWwLS3nm5HlkKzpobXpnaLpmpDol49cclxuICAgKi9cclxuICBvbkhpZGUoKSB7XHJcblxyXG4gIH0sXHJcblxyXG4gIC8qKlxyXG4gICAqIOeUn+WRveWRqOacn+WHveaVsC0t55uR5ZCs6aG16Z2i5Y246L29XHJcbiAgICovXHJcbiAgb25VbmxvYWQoKSB7XHJcblxyXG4gIH0sXHJcblxyXG4gIC8qKlxyXG4gICAqIOmhtemdouebuOWFs+S6i+S7tuWkhOeQhuWHveaVsC0t55uR5ZCs55So5oi35LiL5ouJ5Yqo5L2cXHJcbiAgICovXHJcbiAgb25QdWxsRG93blJlZnJlc2goKSB7XHJcblxyXG4gIH0sXHJcblxyXG4gIC8qKlxyXG4gICAqIOmhtemdouS4iuaLieinpuW6leS6i+S7tueahOWkhOeQhuWHveaVsFxyXG4gICAqL1xyXG4gIG9uUmVhY2hCb3R0b20oKSB7XHJcblxyXG4gIH0sXHJcblxyXG4gIC8qKlxyXG4gICAqIOeUqOaIt+eCueWHu+WPs+S4iuinkuWIhuS6q1xyXG4gICAqL1xyXG4gIG9uU2hhcmVBcHBNZXNzYWdlKCkge1xyXG5cclxuICB9XHJcbn0pIl19