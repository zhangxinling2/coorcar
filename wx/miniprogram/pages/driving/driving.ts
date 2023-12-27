// pages/driving/driving.ts
Page({

  /**
   * 页面的初始数据
   */
  data: {
    elapsed:'15:11:11',
    fee:'15',
    location:{
      latitude:31,
      longitude:120,
  },
  scal:12,
  
  },
  setUpLocationUpdate(){
    wx.startLocationUpdate({
      fail:console.error,
    })
    wx.onLocationChange(loc=>{
      this.setData({
        location:{
          latitude:loc.latitude,
          longitude:loc.longitude
        }
      })
    })
  },
  /**
   * 生命周期函数--监听页面加载
   */
  onLoad() {
    this.setUpLocationUpdate()
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
    wx.stopLocationUpdate()
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