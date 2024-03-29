import { IAppOption } from "../../appoption"
import { TripService } from "../../service/proto_gen/trip"
import { routing } from "../../utils/routing"

// pages/lock/lock.ts
const shareLocationKey="share_location"

const uploadUrl="120.25.124.86:8081/avatar"
Page({

  /**
   * 页面的初始数据
   */
  carId:'',
  data: {
    shareLocation:false,
    avatarURL:'',
  },

  onGetUserInfo(e:any){
    const userInfo:WechatMiniprogram.UserInfo=e.detail.userInfo
    getApp<IAppOption>().resolveUserInfo(userInfo)    
  },
  onShareSwitch(e:any){
    wx.setStorageSync(shareLocationKey,e.detail.value)
  },
  onChooseAvatar(e:any) {
      wx.setStorageSync(avatarUrlKey,e.detail.avatarUrl)
    this.setData({
      avatarURL:e.detail.avatarUrl ,
    })
    // wx.uploadFile({
    //   filePath:e.detail,
    //   name:"avatarImg",
    //   url:uploadUrl,
    //   success:res=>{
    //     console.log(res)
    //   },
    // })
  },
  onUnlockTap(){
    wx.getLocation({
      type:'gcj02',
      success:async loc=>{
        const trip=await  TripService.CreateTrip({
          start:{
              longitude:loc.longitude,
              latitude:loc.latitude,
          },
          carId:this.carId
        })
        console.log(trip)
        //TODO: 开锁，传值头像和经纬度
        wx.showLoading({
          title:'开锁中',
          mask:true,
        })
        setTimeout(() => {
          
          wx.redirectTo({
            url:routing.driving({
              trip_id: trip.id!,
            }),
            complete:()=>wx.hideLoading(),
          })
        },2000);
      },
      fail:()=>wx.showToast({
        icon:'none',
        title:"请前往设置页授权位置信息"
      })
    })
    
  },

  /**
   * 生命周期函数--监听页面加载
   */
  async onLoad(opt:Record<'car_id',string>) {
      
    const o:routing.LockOpts=opt
    this.carId=o.car_id
    //const userInfo=await getApp<IAppOption>().globalData.userInfo
     
    this.setData({
        avatarURL:wx.getStorageSync(avatarUrlKey)||'',
      shareLocation:wx.getStorageSync(shareLocationKey)||false
    })
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
})